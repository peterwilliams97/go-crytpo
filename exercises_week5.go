package main

/*
	Your goal this week is to write a program to compute discrete log modulo a prime p. Let g be
	some element in ℤ∗p and suppose you are given h in ℤ∗p such that h=g^x where 1≤x≤2^40. Your goal
	is to find x. More precisely, the input to your program is p,g,h and the output is x.

	p = 134078079299425970995740249982058461274793658205923933 \
	    77723561443721764030073546976801874298166903427690031 \
	    858186486050853753882811946569946433649006084171
	g = 11717829880366207009516117596335367088558084999998952205 \
	    59997945906392949973658374667057217647146031292859482967 \
	    5428279466566527115212748467589894601965568
	h = 323947510405045044356526437872806578864909752095244 \
	    952783479245297198197614329255807385693795855318053 \
	    2878928001494706097394108577585732452307673444020333
*/
import (
	"fmt"
	"math/big"
	"time"
)

const (
	ps = "134078079299425970995740249982058461274793658205923933" +
		"77723561443721764030073546976801874298166903427690031" +
		"858186486050853753882811946569946433649006084171"
	gs = "11717829880366207009516117596335367088558084999998952205" +
		"59997945906392949973658374667057217647146031292859482967" +
		"5428279466566527115212748467589894601965568"
	hs = "323947510405045044356526437872806578864909752095244" +
		"952783479245297198197614329255807385693795855318053" +
		"2878928001494706097394108577585732452307673444020333"
)

func main() {
	p := bigDec(ps)
	g := bigDec(gs)
	h := bigDec(hs)
	B := Pow(2, 20)

	// p = i2b(7)
	// g = i2b(3)
	// h = i2b(5)
	// B = 4
	// for x := 0; x < int(p.Int64()); x++ {
	// 	h1 := bigExp(g, i2b(x), p)
	// 	h2 := Pow(int(g.Int64()), x) % int(p.Int64())
	// 	fmt.Printf("%d^%d %% %d = %d = %d\n", g.Int64(), x, p.Int64(), h1.Int64(), h2)
	// }
	// for x := 1; x < int(p.Int64()); x++ {
	// 	y := modInv(i2b(x), p)
	// 	z := bigMod(bigMul(i2b(x), y), p)
	// 	fmt.Printf("%d^-1 %% %d = %d z=%d\n", x, p.Int64(), y.Int64(), z.Int64())
	// 	if z.Int64() != 1 {
	// 		panic("Bad inverse")
	// 	}
	// }
	x := solve(p, g, h, B)
	h1 := bigExp(g, i2b(x), p)
	fmt.Printf("x=%d\n", x)
	fmt.Printf("%d^%d (mod %d) = %d\n", g.Int64(), x, p.Int64(), h1.Int64())
	if h.String() != h1.String() {
		panic("Wrong answer")
	}
}

func solve(p, g, h *big.Int, B int) int {

	fmt.Printf("p=%s\n", p)
	fmt.Printf("g=%s\n", g)
	fmt.Printf("h=%s\n", h)
	fmt.Printf("B=%d\n", B)

	t0 := time.Now()
	t := t0
	lhsX1 := map[string]int{}
	for x1 := 0; x1 < B; x1++ {
		// lhs = h/(g^x1) (mod p)
		b := bigExp(g, i2b(x1), p)
		b1 := modInv(b, p)
		a := bigMul(h, b1)
		lhs := bigMod(a, p)
		lhsX1[lhs.String()] = x1
		// fmt.Printf("%d %s\n", x1, lhs[x1])
		// if x1 > 10 {
		// 	panic("stop")
		// }
		// fmt.Printf("  x1=%d b=%d b1=%d a=%d lhs=%d\n", x1, b.Int64(), b1.Int64(),
		// 	a.Int64(), lhs.Int64())
		if time.Since(t) > 10*time.Second || x1 == B-1 {
			dt := time.Since(t0).Seconds()
			n := len(lhsX1)
			fmt.Printf("%.1f sec %d %.3f %.2f\n", dt, n, float64(n)/float64(dt), float64(n)/float64(B))
			t = time.Now()
		}
	}

	gB := bigExp(g, i2b(B), p)
	for x0 := 0; x0 < B; x0++ {
		rhs := bigExp(gB, i2b(x0), p)
		// if x0%(B/20) == 0 {
		// 	fmt.Printf("x0=%d v=%s\n", x0, rhs)
		// }
		x1, ok := lhsX1[rhs.String()]
		if ok && (x0 != 0 || x1 != 0) {
			x := x0*B + x1
			fmt.Printf("x0=0x%X\n", x0)
			fmt.Printf("x1=0x%X\n", x1)
			fmt.Printf("x =0x%X\n", x)
			fmt.Printf("v =%s\n", rhs)
			return x
		}
	}
	return -1
}

func bigDec(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}

func bigExp(x, y, m *big.Int) *big.Int {
	i := new(big.Int)
	return i.Exp(x, y, m)
}

func bigMul(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Mul(x, y)
}

func bigMod(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Mod(x, y)
}

func bigSub(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Sub(x, y)
}

func bigExp2(x, y int) *big.Int {
	i := new(big.Int)
	return i.Exp(i2b(x), i2b(y), nil)
}

// modInv returns 1 / x (mod p) where p is prime
func modInv(x, p *big.Int) *big.Int {
	// Fermat's little theorem
	return bigExp(x, bigSub(p, i2b(2)), p)
}

func i2b(x int) *big.Int {
	i := new(big.Int)
	i.SetInt64(int64(x))
	return i
}

// Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
