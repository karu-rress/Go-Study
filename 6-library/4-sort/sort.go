package main

import (
	"fmt"
	"sort"
)

func Sort() {
	a := []int{1, 3, 6, 2, 6}

	// Int, Float64, String
	sort.Sort(sort.IntSlice(a))
	fmt.Println(a)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)

	// Ints, Float64s, Strings
	sort.Ints(a)
	fmt.Println(a)
}

type Student struct { // 정렬할 데이터
	name  string
	score float32
}

type By func(s1, s2 *Student) bool

func (by By) Sort(students []Student) {
	sorter := &studentSorter{students: students, by: by}
	sort.Sort(sorter)
}

type studentSorter struct {
	students []Student
	by       func(s1, s2 *Student) bool
}

// sort.Interface 구현 메서드 3개
func (s *studentSorter) Len() int {
	return len(s.students)
}

func (s *studentSorter) Less(i, j int) bool {
	return s.by(&s.students[i], &s.students[j])
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func SortWithKey() {
	s := []Student{
		{"Maria", 39.5},
		{"Andrew", 78.2},
		{"John", 96.3},
	}

	name := func(p1, p2 *Student) bool {
		return p1.name < p2.name
	}
	score := func(p1, p2 *Student) bool {
		return p1.score < p2.score
	}
	reverseScore := func(p1, p2 *Student) bool {
		return !score(p1, p2)
	}

	// By(x) : 그냥 형변환 한 거 (생성자)
	By(name).Sort(s)
	fmt.Println(s)

	By(reverseScore).Sort(s)
	fmt.Println(s)
}

func main() {
	Sort()
	SortWithKey()
}
