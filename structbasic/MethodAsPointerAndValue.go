package structbasic

import "fmt"

type List []int

func (lst *List) Append(val int) {
	*lst = append(*lst, val)
}

func (lst List) Len() int {
	return len(lst)
}

func (lst List) get(i int) int {
	return lst[i]
}

func TestMethodAsPointer() {
	lst := new(List)
	lst.Append(10)
	lst.Append(12)
	for i := 0; i < lst.Len(); i++ {
		fmt.Println((*lst)[i])
	}
}
