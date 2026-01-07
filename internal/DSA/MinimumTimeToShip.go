package dsa

/*

You are managing a set of n warehouses, where the i-th warehouse contains piles[i] packages that need to be shipped.
You have h hours to complete all shipments before a strict deadline.

Each hour, a shipping truck can visit any one warehouse and ship up to k packages from it.
If the warehouse has fewer than k packages, the truck ships all of them and remains idle for the rest of that hour.
Your objective is to determine the minimum shipping capacity k (packages per hour) such that all packages from all warehouses can be shipped within h hours.

Return the smallest possible integer k that satisfies this condition.


Example 1:

Input: piles = [3,6,7,11], h = 8

Output: 4


Example 2:

Input: piles = [30,11,23,4,20], h = 5
Output: 30
*/

import "fmt"

func canDelivery(weights []int, h, capacity int) bool {
	currentTime := 0
	for _, weight := range weights {
		time := (weight / capacity)
		if weight%capacity != 0 {
			time = time + 1
		}

		if currentTime+time <= h {
			currentTime += time
		} else {
			return false
		}
	}
	return currentTime <= h
}

func RunMinimumTimeToShip() {

	maxVal := 0
	//piles := []int{3, 6, 7, 11}
	piles := []int{30, 11, 23, 4, 20}
	h := 5
	for _, pile := range piles {
		if pile > maxVal {
			maxVal = pile
		}
	}
	ans := maxVal
	for i := maxVal; i >= 0; {
		if canDelivery(piles, h, i) {
			ans = i
			i--
		} else {
			break
		}
	}
	fmt.Println(ans)

}
