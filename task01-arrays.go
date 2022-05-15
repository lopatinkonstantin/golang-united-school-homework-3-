package homework

func average(input [15]float32) (result float32) {
	var s float32
   	for i:=0; i<15; i++ {
   		s+=input[i]
   	}
   	return s/15
}
