package maxSubSum

func maxSubSum_1(list []int)int {
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


func maxSubSum_2(list []int)int {
	var max, sum int
	for i:=0; i<len(list); i++ {
		sum = 0
		for j:=i; j<len(list); j++ {
			sum += list[j]			
			if max < sum {
				max = sum
			} 
		}
	}
	return max
}

// it's just an accident
func maxSubSum_3(list []int)int {
	var max, sum int
	for i:=0; i<len(list); i++ {
		sum += list[i]			
		if max < sum {
			max = sum
		}
	}
	return max
}
