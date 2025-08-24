package stringexamples

import "fmt"

func strings1() {
	s := "ఐ"
	// fmt.Println(s,len(s))
	fmt.Printf("%8T %[1]v\n", s)
	fmt.Printf("%8T %[1]v\n", []rune(s)) // logical representation
	fmt.Printf("%8T %[1]v\n", []byte(s)) // physical representation
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
