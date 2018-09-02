/* Solve the k-th (k = N/2) max problem, check the performance when k is different.
 */

package kth 
import (
)

func Sort (dataset []int) []int {
	for i:=0; i<len(dataset); i++ {
		for j:=i+1; j<len(dataset); j++ {
			if dataset[i] < dataset[j] {
				dataset[i], dataset[j] = dataset[j], dataset[i]
			}		
		}	
	}			
	return dataset
} 

func Kth(k int, dataset []int) int {
	list := Sort(dataset)	
	return list[k - 1]	
}
