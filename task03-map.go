package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	r:=make([]string,len(input),len(input))
	s:=make([]int,len(input),len(input))
	j:=0
	for i:= range input {
		s[j]=i
		j++
	}
	sort.Ints(s)
	for i:=0; i<len(input); i++ {
		r[i]=input[s[i]]
	}
	return r
}
