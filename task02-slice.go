package homework

func reverse(input []int64) (result []int64) {
	r:=make([]int64,len(input),len(input))
	for i:=len(input)-1;i>=0;i--{
		r[len(input)-i-1]=input[i]
	}
	return r
}
