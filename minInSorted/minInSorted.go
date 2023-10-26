package main

func findMin(nums []int) int {
	//fully rotated
	if nums[0] < nums[len(nums)/2] && nums[len(nums)/2] < nums[len(nums)-1] {
		return nums[0]
	}
	//half rotated
	if nums[len(nums)/2] < nums[len(nums)/2-1] {
		return nums[len(nums)/2]
	}

	//partially rotated
	lp := 0
	rp := len(nums) - 1

	//

	for lp <= rp {
		med := (lp + rp) / 2

		//if the number to the left is bigger, current is the biggest
		if med > 0 && nums[med-1] > nums[med] {
			return nums[med]
		}
		if med == len(nums)-1 {
			break
		}
		//if the number to the right is smaller, right is the smallest
		if nums[med+1] < nums[med] {
			return nums[med+1]
		}

		//if med > end then between med and end
		//if med < end then between beg and med

		if nums[med] > nums[rp] {
			lp = med + 1
		} else {
			rp = med - 1
		}
	}
	return 0
}

//			|
//0,1,2,4,5,6,7,8,9,10,11
//1,2,4,5,6,7,8,9,10,11,0
//2,4,5,6,7,8,9,10,11,0,1
//4,5,6,7,8,9,10,11,0,1,2
//5,6,7,8,9,10,11,0,1,2,4
//6,7,8,9,10,11,0,1,2,4,5
//7,8,9,10,11,0,1,2,4,5,6
//8,9,10,11,0,1,2,4,5,6,7
//9,10,11,0,1,2,4,5,6,7,8
//10,11,0,1,2,4,5,6,7,8,9
//11,0,1,2,4,5,6,7,8,9,10
//0,1,2,4,5,6,7,8,9,10,11,

func main() {
	println(findMin([]int{2, 4, 5, 6, 7, 8, 9, 10, 11, 0, 1}))
	println(findMin([]int{11, 0, 1, 2, 4, 5, 6, 7, 8, 9, 10}))
	println(findMin([]int{0, 1, 2, 4, 5, 6, 7, 8, 9, 10, 11}))
}
