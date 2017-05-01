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


	Problem 1
	    N=179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581
	   pq=179769313486231590772930519078902473361797697894230657273430081157732675805505620686985379449212982959585501387537164015710139858647833778606925583497541085196591615128057575940752635007475935288710823649949940771895617054361149474865046711015101563940680527540071584560878577663743040086340742855278549092581
	    A=13407807929942597099574024998205846127479365820592393377723561443721764030073720664935730272460038784167544890048265955830752343658005756322445518700544291
	    x=57896044618658097711785492504343953926634992332820282019728792003956564819990
	    q=13407807929942597099574024998205846127479365820592393377723561443721764030073778560980348930557750569660049234002192590823085163940025485114449475265364281
	    p=13407807929942597099574024998205846127479365820592393377723561443721764030073662768891111614362326998675040546094339320838419523375986027530441562135724301
	  q-p=                                                                             115792089237316195423570985008687907853269984665640564039457584007913129639980
	Problem 2
	    i=72077
	    N=648455842808071669662824265346772278726343720706976263060439070378797308618081116462714015276061417569195587321840254520655424906719892428844841839353281972988531310511738648965962582821502504990264452100885281673303711142296421027840289307657458645233683357077834689715838646088239640236866252211790085787877
	   pq=648455842808071669662824265346772278726343720706976263060439070378797308618081116462714015276061417569195587321840254520655424906719892428844841839353281972988531310511738648965962582821502504990264452100885281673303711142296421027840289307657458645233683357077834689715838646088239640236866252211790085787877
	    A=25464796146996183438008816563973942229341454268524157846328581927885777970045810444817548800110193810869973985549488954023184312335376930269142469838897581
	    x=60587609673697726860620356762589523992295780718687067521305424478575876632477822
	    q=25464796146996183438008816563973942229341454268524157846328581927885777970106398054491246526970814167632563509541784734741871379856682354747718346471375403
	    p=25464796146996183438008816563973942229341454268524157846328581927885777969985222835143851073249573454107384461557193173304497244814071505790566593206419759
	  q-p=                                                                          121175219347395453721240713525179047984591561437374135042610848957151753264955644


*/
import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// problem1()
	// problem2()
	// problem3()
	problem4()
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

func problem3() {
	// 72006226374735042527956443552558373833808445147399984182665305798191 \
	// 63556901883377904234086641876639384851752649940178970835240791356868 \
	// 77441155132015188279331812309091996246361896836573643119174094961348 \
	// 52463970788523879939683923036467667022162701835329944324119217381272 \
	// 9276147530748597302192751375739387929
	const Ns = "72006226374735042527956443552558373833808445147399984182665305798191" +
		"63556901883377904234086641876639384851752649940178970835240791356868" +
		"77441155132015188279331812309091996246361896836573643119174094961348" +
		"52463970788523879939683923036467667022162701835329944324119217381272" +
		"9276147530748597302192751375739387929"

	N := bigDec(Ns)
	// N6 := bigMul(N, i2b(6))
	// A := bigSqrt(N6)

	p2, exact := bigDiv(bigMul(N, i2b(2)), i2b(3))
	if !exact && false {
		panic("p")
	}
	q2, exact := bigDiv(bigMul(N, i2b(3)), i2b(2))
	if !exact && false {
		panic("q")
	}
	p := bigSqrt(p2)
	q := bigSqrt(q2)

	// A1 := bigAdd(A, i2b(1))
	// d := bigSub(bigMul(A, A), N6)
	// d1 := bigSub(bigMul(A1, A1), N6)
	// fmt.Printf("d=%s\n", d)
	// fmt.Printf("d1=%s\n", d1)
	// // panic("!")

	// q := bigDiv(A, i2b(2))
	// p := bigDiv(N, q)
	pq := bigMul(p, q)

	fmt.Println("Problem 3")
	fmt.Printf("    N=%s\n", N.String())
	fmt.Printf("   pq=%s\n", pq.String())
	fmt.Printf(" N-pq=%s\n", padLeft(bigSub(N, pq).String(), len(N.String())))
	fmt.Printf("    q=%s\n", q.String())
	fmt.Printf("    p=%s\n", p.String())
	fmt.Printf("  q-p=%s\n", padLeft(bigSub(q, p).String(), len(p.String())))

	// A0 := bigSqrt(bigMul(N, i2b(6)))
	// for i := -Pow(2, 20); i < Pow(2, 20); i++ {
	// 	A := bigAdd(A0, i2b(i))
	// 	q, exact := bigDiv(A, i2b(2))
	// 	if !exact && false {
	// 		continue
	// 	}
	// 	p, exact := bigDiv(N, q)
	// 	if !exact && false {
	// 		continue
	// 	}
	// 	pq := bigMul(p, q)
	// 	if pq.String() == N.String() {

	// 		fmt.Println("Problem 3")
	// 		fmt.Printf("    N=%s\n", N.String())
	// 		fmt.Printf("   pq=%s\n", pq.String())
	// 		fmt.Printf(" N-pq=%s\n", padLeft(bigSub(N, pq).String(), len(N.String())))
	// 		fmt.Printf("    A=%s\n", A.String())
	// 		fmt.Printf("    q=%s\n", q.String())
	// 		fmt.Printf("    p=%s\n", p.String())
	// 		fmt.Printf("  q-p=%s\n", padLeft(bigSub(q, p).String(), len(p.String())))
	// 		return
	// 	}
	// }
	// // if pq.String() != N.String() {
	// panic("Wrong answer")
	// // }
}

func problem4() {
	// C = 22096451867410381776306561134883418017410069787892831071731839143676135600120538004282329650473509424343946219751512256465839967942889460764542040581564748988013734864120452325229320176487916666402997509188729971690526083222067771600019329260870009579993724077458967773697817571267229951148662959627934791540
	const (
		Cs = "22096451867410381776306561134883418017410069787892831071731839143676135600120538004282329650473509424343946219751512256465839967942889460764542040581564748988013734864120452325229320176487916666402997509188729971690526083222067771600019329260870009579993724077458967773697817571267229951148662959627934791540"
		qs = "13407807929942597099574024998205846127479365820592393377723561443721764030073778560980348930557750569660049234002192590823085163940025485114449475265364281"
		ps = "13407807929942597099574024998205846127479365820592393377723561443721764030073662768891111614362326998675040546094339320838419523375986027530441562135724301"
		ei = 65537
	)

	fmt.Println("Problem 4")
	C := bigDec(Cs)
	p := bigDec(ps)
	q := bigDec(qs)
	N := bigMul(p, q)

	phi := bigMul(bigSub(p, i2b(1)), bigSub(q, i2b(1)))
	e := i2b(ei)
	d := bigModInv(e, phi)
	ed := bigMod(bigMul(e, d), phi)
	M := bigExp(C, d, N)

	fmt.Printf("    C=%s\n", C.String())
	fmt.Printf("    C=%d\n", C)
	fmt.Printf("    C=0x%X\n", C)
	fmt.Printf("    e=%d\n", e)
	fmt.Printf("    d=%d\n", d)
	fmt.Printf("   ed=%d\n", ed)
	s := fmt.Sprintf("%X", M)
	if len(s)%2 != 0 {
		s = "0" + s
	}
	i := strings.Index(s, "00") + 2
	fmt.Printf("    M=[%3d]0x%s\n", len(s), s)
	fmt.Printf("    m=[%3d]0x%s\n", len(s[i:]), padLeft(s[i:], len(s)))
	fmt.Printf("    i=%d\n", i)

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

// bigModInv returns 1/x (mod n)
func bigModInv(x, n *big.Int) *big.Int {
	i := new(big.Int)
	// Sqrt sets z to ⌊√x⌋, the largest integer such that z² ≤ x, and returns z.
	return i.ModInverse(x, n)
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

func bigDiv(x, y *big.Int) (*big.Int, bool) {
	i := new(big.Int)
	m := new(big.Int)
	q, r := i.QuoRem(x, y, m)
	// if r.String() != i2b(0).String() {
	// 	fmt.Printf("x=%s\n", x.String())
	// 	fmt.Printf("y=%s\n", y.String())
	// 	panic(fmt.Sprintf("Division not exact r=%s", r.String()))
	// }
	return q, r.String() == i2b(0).String()
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
