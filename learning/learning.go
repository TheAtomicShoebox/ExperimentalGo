package learning

import (
	"fmt"
	"io"
	"iter"
	"log"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"time"
	"unicode/utf8"
)

/*
I will not apologize for how
ridiculously oversized this file is.
that would imply that I care, and I don't
*/
func Learning_main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithProg)
	if len(argsWithoutProg) > 0 {
		fmt.Println(argsWithoutProg)

		var intArgs []int
		for _, arg := range argsWithoutProg {
			i, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal(err)
			}
			intArgs = append(intArgs, i)
		}

		for idx, i := range intArgs {
			if idx+1 == len(intArgs) {
				fmt.Printf("%d = %d\n", i, sum(intArgs...))
			} else {
				fmt.Printf("%d + ", i)
			}
		}
		methodTesting()
	}

	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("pointer: ", &i)

	runeExercise()

	structExercise()

	enumExercise()

	genericsExercise()

	RangeOverIterators()
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func sum(a ...int) int {
	sum := 0
	for _, operand := range a {
		sum += operand
	}
	return sum
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type describer interface {
	describe() string
}

func genericsExercise() {
	var s = []string{"foo", "bar", "zoo"}
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	//types shown here explicitly
	_ = SlicesIndex[[]string, string](s, "zoo")

	list := List[int]{}
	list.Push(10)
	list.Push(22)
	list.Push(54)
	fmt.Println("list: ", list.AllElements())
}

func enumExercise() {
	ns := Transition(StateIdle)
	fmt.Println(ns)

	ns2 := Transition(ns)
	fmt.Println(ns2)
}

func structExercise() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{Name: "Alice", Age: 30})
	fmt.Println(Person{Name: "Fred"})
	fmt.Println(&Person{Name: "Ann", Age: 40})
	fmt.Println(NewPerson("Jon", 42))

	s := Person{Name: "Sean", Age: 50}
	fmt.Println(s.Name)

	sp := &s
	fmt.Println(sp.Age)
	sp.Age = 51
	fmt.Println(sp.Age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	c := circle{radius: 5}
	measure(r)
	measure(c)

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	var d describer = co
	fmt.Println("describer:", d.describe())
}

func runeExercise() {
	const s = "สวัสดี"

	fmt.Println("Len:", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

func methodTesting() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newSeq := intSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())

	fmt.Println(factorial(7))

	var fibonacci func(n int) int

	fibonacci = func(n int) int {
		if n < 2 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	fmt.Println(fibonacci(24))

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s => %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func FormatAndWrite(w io.Writer, format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	io.WriteString(w, s)
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	p := Person{Name: name, Age: age}
	return &p
}

type ServerState int

const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func Transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	Head, Tail *Element[T]
}

type Element[T any] struct {
	Next *Element[T]
	Val  T
}

func (lst *List[T]) Push(v T) {
	if lst.Tail == nil {
		lst.Head = &Element[T]{Val: v}
		lst.Tail = lst.Head
	} else {
		lst.Tail.Next = &Element[T]{Val: v}
		lst.Tail = lst.Tail.Next
	}
}

func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.Head; e != nil; e = e.Next {
		elems = append(elems, e.Val)
	}
	return elems
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.Head; e != nil; e = e.Next {
			if !yield(e.Val) {
				return
			}
		}
	}
}

func GenFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func RunArrays() {
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("length:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2D: ", twoD)
}

func RunSlices() {
	var s []string
	fmt.Println("uninitialized: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty: ", s, s == nil, len(s) == 0)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get: ", s)
	fmt.Println("set: ", s[2])
	fmt.Println("length: ", len(s[2]))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy: ", c)

	l := s[2:5]
	fmt.Println("slice: ", l)

	l = s[:5]
	fmt.Println("slice: ", l)

	t := []string{"g", "h", "i"}
	fmt.Println("declare: ", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D: ", twoD)
}

func RunMaps() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("length: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	clear(m)
	fmt.Println("map: ", m)

	_, isPresent := m["k2"]
	fmt.Println("k2 is present: ", isPresent)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func RangeOverIterators() {
	list := List[int]{}
	list.Push(10)
	list.Push(22)
	list.Push(54)

	for e := range list.All() {
		fmt.Println(e)
	}
	all := slices.Collect(list.All())
	fmt.Println("all:", all)

	for n := range GenFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}

func ForLoops() {
	i := 1
	for i <= 3 {
		fmt.Printf("i = %d\n", i)
		i += 1
	}

	for j := 1; j < 3; j++ {
		FormatAndWrite(os.Stdout, "j = %v\n", j)
	}

	for i := range 3 {
		FormatAndWrite(os.Stdout, "i = %v\n", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func IfElse() {
	if 7%2 == 0 {
		fmt.Printf("7 is even")
	} else {
		fmt.Printf("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("0 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func Switch() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("I don't recognize type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(time.Now())
}
