package main

/*
	The following modulus N is a products of two primes p and q where |p−q|<2N^1/4. Find the smaller
	of the two factors

		Assume p <= q

		A := (p + q) / 2
		0 <= A−sqrt(N) < 1
	so 	A=ceil(sqrt(N))

		x := q - A
		N = p*q = (A - x) * (A + x)
		x = (A^2 - N)^1/2

	N = 17976931348623159077293051907890247336179769789423065727343008115 \
	77326758055056206869853794492129829595855013875371640157101398586 \
	47833778606925583497541085196591615128057575940752635007475935288 \
	71082364994994077189561705436114947486504671101510156394068052754 \
	0071584560878577663743040086340742855278549092581

	N =179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581
	pq=179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581
	A =13407807929942597099574024998205846127479365820592393377723561443721764030073720664935730272460038784167544890048265955830752343658005756322445518700544291
	x=57896044618658097711785492504343953926634992332820282019728792003956564819990
	q=13407807929942597099574024998205846127479365820592393377723561443721764030073778560980348930557750569660049234002192590823085163940025485114449475265364281
	p=13407807929942597099574024998205846127479365820592393377723561443721764030073662768891111614362326998675040546094339320838419523375986027530441562135724301
	q-p=                                                                           115792089237316195423570985008687907853269984665640564039457584007913129639980


*/
import (
	"fmt"
	"math/big"
)

func main() {
	problem1()
	problem2()
}

func problem1() {
	const Ns = "17976931348623159077293051907890247336179769789423065727343008115" +
		"77326758055056206869853794492129829595855013875371640157101398586" +
		"47833778606925583497541085196591615128057575940752635007475935288" +
		"71082364994994077189561705436114947486504671101510156394068052754" +
		"0071584560878577663743040086340742855278549092581"

	N := bigDec(Ns)
	A := bigSqrtCeil(N)
	x := bigSqrt(bigSub(bigMul(A, A), N))
	p := bigSub(A, x)
	q := bigAdd(A, x)
	pq := bigMul(p, q)

	fmt.Println("Problem 1")
	fmt.Printf("    N=%s\n", N.String())
	fmt.Printf("   pq=%s\n", pq.String())
	fmt.Printf("    A=%s\n", A.String())
	fmt.Printf("    x=%s\n", x.String())
	fmt.Printf("    q=%s\n", q.String())
	fmt.Printf("    p=%s\n", p.String())
	fmt.Printf("  q-p=%s\n", padLeft(bigSub(q, p).String(), len(p.String())))
	if pq.String() != N.String() {
		panic("Wrong answer")
	}
}

func problem2() {
	const Ns = "6484558428080716696628242653467722787263437207069762630604390703787" +
		"9730861808111646271401527606141756919558732184025452065542490671989" +
		"2428844841839353281972988531310511738648965962582821502504990264452" +
		"1008852816733037111422964210278402893076574586452336833570778346897" +
		"15838646088239640236866252211790085787877"

	fmt.Println("Problem 2")
	N := bigDec(Ns)
	A0 := bigSqrtCeil(N)
	var A, x, p, q, pq *big.Int
	found := false
	for i := 0; i < Pow(2, 20); i++ {
		A = bigAdd(A0, i2b(i))
		x = bigSqrt(bigSub(bigMul(A, A), N))
		p = bigSub(A, x)
		q = bigAdd(A, x)
		pq = bigMul(p, q)
		if pq.String() == N.String() {
			fmt.Printf("    i=%d\n", i)
			found = true
			break
		}
	}

	if !found {
		panic("Wrong answer")
	}

	fmt.Printf("    N=%s\n", N.String())
	fmt.Printf("   pq=%s\n", pq.String())
	fmt.Printf("    A=%s\n", A.String())
	fmt.Printf("    x=%s\n", x.String())
	fmt.Printf("    q=%s\n", q.String())
	fmt.Printf("    p=%s\n", p.String())
	fmt.Printf("  q-p=%s\n", padLeft(bigSub(q, p).String(), len(p.String())))
}

func padLeft(s string, n int) string {
	for len(s) < n {
		s = " " + s
	}
	return s
}

func bigDec(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}

func bigFloat(s string) *big.Float {
	i := new(big.Float)
	i.SetString(s)
	return i
}

// bigSqrt returns ⌊√x⌋
func bigSqrt(x *big.Int) *big.Int {
	i := new(big.Int)
	// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
	return i.Sqrt(x)
}

// bigSqrtCeil returns ceil(√x)
func bigSqrtCeil(x *big.Int) *big.Int {
	i := bigSqrt(x)
	if bigMul(i, i).String() == x.String() {
		return i
	}
	return bigAdd(i, i2b(1))
}

func f2b(x *big.Float) *big.Int {
	i := new(big.Int)
	floor, acc := x.Int(i)
	if acc != big.Exact {
		panic("Not an int")
	}
	return floor
}

func bigCeil(x *big.Float) *big.Int {
	i := new(big.Int)
	floor, acc := x.Int(i)
	if acc == big.Exact {
		return floor
	}
	return bigAdd(floor, i2b(1))
}

func bigExp(x, y, m *big.Int) *big.Int {
	i := new(big.Int)
	return i.Exp(x, y, m)
}

func bigExp3(x *big.Int, y int) *big.Int {
	i := new(big.Int)
	return i.Exp(x, i2b(y), nil)
}

func bigMul(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Mul(x, y)
}

func bigMod(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Mod(x, y)
}

func bigAdd(x, y *big.Int) *big.Int {
	i := new(big.Int)
	return i.Add(x, y)
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

func b2f(i *big.Int) *big.Float {
	x := new(big.Float)
	x.SetInt(i)
	return x
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
