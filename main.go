package main

import (
	"bufio"
	"fmt"
	"main/internal/concurrency"
	_ "main/internal/pointers"
	"os"
	"sort"
	"strings"
)

func strings1() {
	s := "ఐ"
	// fmt.Println(s,len(s))
	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s)) // logical representation
	fmt.Printf("%8T %[1]v\n", []byte(s)) // physical representation
}

func getAverage() {
	var sum float64
	var n int

	for {
		var val float64
		_, err := fmt.Fscanln(os.Stdin, &val)
		if err != nil {
			break
		}
		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}
	fmt.Println("Avg:", sum/float64(n))

}

func f1() {
	a := 1
	b := 21.11
	fmt.Printf("%d\n", a)
	fmt.Printf("%T %[1]v", b)
}

func replace() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}
	old, new1 := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		// s := strings.Replace(scan.Text(),old,new1,-1)
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new1)

		fmt.Println(t)
	}
}

func countUniqueWords() {
	scan := bufio.NewScanner(os.Stdin)
	// make function creates some space for the map
	words := make(map[string]int)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}
	fmt.Println(len(words), "unique words")

	type kv struct {
		key string
		val int
	}
	var ss []kv
	for k, v := range words {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})
	for _, v := range ss[:5] {
		fmt.Println(v.key, "appeared", v.val, "times")
	}

}

func formatStringForInteger() {
	a, b := 31, 10
	fmt.Printf("%d %d\n", a, b)
	fmt.Printf("%x %x\n", a, b)
	fmt.Printf("%X %X\n", a, b)
	fmt.Printf("%#X %#X\n", a, b)
	fmt.Printf("%#x %#x\n", a, b)
	fmt.Printf("|%6d|%6d|\n", a, b)
	fmt.Printf("|%-6d|%-6d|\n", a, b) // left justified
	fmt.Printf("|%06d|%06d|\n", a, b)

	fmt.Println()
	s := []rune{'a', 'b', 'c'}
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#q\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%T\n", s)

}

func formattingMapsAndStrings() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(m)
	fmt.Printf("%v\n", m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#q\n", m)
	fmt.Printf("%q\n", m)
	fmt.Printf("%T\n", m)

	s := "hi Hello"
	s2 := []byte(s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%#q\n", s)
	fmt.Printf("%q\n", s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", s2)

}

// A descriptor would be a data structure that holds information of some data and a reference (pointer) to it.

// closures

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

func callingAClosure() {
	f, g := fib(), fib()
	fmt.Println(f(), g())
}

func main() {
	// getAverage()
	// strings1()
	// replace()
	// countUniqueWords()
	// formatStringForInteger()
	// formattingMapsAndStrings()
	// callingAClosure()
	// pointers.PrintEmail()
	// concurrency.ConcurrentPrint()
	// concurrency.PrintSumUsingChannels()
	concurrency.UpdateMapUsingMutex()
}
