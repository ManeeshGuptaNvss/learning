package main

import (
	"bufio"
	"fmt"
	"main/internal/concurrency"
	_ "main/internal/interfaces"
	_ "main/internal/pointers"
	"os"
	"sort"
	"strings"
)

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

// A descriptor would be a data structure that holds information of some data and a reference (pointer) to it.

/*
TODO: Try writing tests for all these functions
Learn what is the difference between RWMutex and Mutex
*/
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
	concurrency.WorkerPoolMain()
	// interfaces.ExampleMain()
}
