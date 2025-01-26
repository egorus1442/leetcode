package main

func removeDuplicates(nums []int) int {
	indexInsert := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[indexInsert] {
			indexInsert++
			nums[indexInsert] = nums[i]
		}
	}
	return indexInsert + 1
}

func main() {

}
