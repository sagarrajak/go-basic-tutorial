package packagesexample

import (
	"fmt"
	"regexp"
	"sync"
)

type Info struct {
	mu  sync.Mutex
	Str string
}

func Update(info *Info) {
	info.mu.Lock()
	info.Str = ""
	info.mu.Unlock()
}

func RegExpressionExample() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18" // string to search
	pat := "[0-9]+.[0-9]+"                                      // pattern search in searchIn

	// f := func(s string) string {
	// 	v, _ := strconv.ParseFloat(s, 32)
	// 	return strconv.FormatFloat(v*2, 'f', 2, 32)
	// }
	// if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
	// 	fmt.Println("Match found!")
	// }
	// re, _ := regexp.Compile(pat)

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}

	reg, _ := regexp.Compile(pat)
	str := reg.ReplaceAllString(searchIn, "##.##")
	fmt.Println(str)

	// str := re.ReplaceAllString(searchIn, "##.#") // replace pat with "##.#"
	// fmt.Println(str)
	// // using a function
	// str2 := re.ReplaceAllStringFunc(searchIn, f)
	// fmt.Println(str2) // pattern search in searchIn
}
