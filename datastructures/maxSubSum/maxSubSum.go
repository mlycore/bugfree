package maxSubSum

func maxSubSum(list []int)int {
	var max, sum int
	for i:=0; i<len(list); i++ {
		for j:=i; j<len(list); j++ {
			sum = 0
			for k:=i; k<=j; k++ {
				sum += list[k]
			}	
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

