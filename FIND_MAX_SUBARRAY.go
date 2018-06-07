package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	L, R, S := FIND_MAX_SUBARRAY(A, 0, len(A)-1)
	fmt.Printf("L: %d, R: %d, S: %d, %v\n", L, R, S, A[L:R+1])
}

// FIND-MAX-SUBARRAY(A, low, high)
//   if low == high
//     return (low, high, A[low])				// base case: only one element
//   mid = [(low + high) / 2] // round down
//   (left-low, left-high, left-sum) = FIND-MAX-SUBARRAY(A, low, mid)
//   (right-low, right-high, right-sum) = FIND-MAX-SUBARRAY(A, mid+1, high)
//   (cross-low, cross-high, cross-sum) = FIND-MAX-CROSSING-SUBARRAY(A, low, mid, high)
//   if left-sum >= right-sum and left-sum >= cross-sum
//     return (left-low, left-high, left-sum)
//   elseif right-sum >= right-sum and right-sum >= cross-sum
//     return (right-low, right-high, right-sum)
//   else
//     return (cross-low, cross-high, cross-sum)
func FIND_MAX_SUBARRAY(A []int, low, high int) (max_left, max_right, max_sum int) {
	if low == high {
		return low, high, A[low]
	}

	mid := (low + high) / 2
	left_low, left_high, left_sum := FIND_MAX_SUBARRAY(A, low, mid)
	right_low, right_high, right_sum := FIND_MAX_SUBARRAY(A, mid+1, high)
	cross_low, cross_high, cross_sum := FIND_MAX_CROSSING_SUBARRAY(A, low, mid, high)
	if left_sum >= right_sum && left_sum >= cross_sum {
		return left_low, left_high, left_sum
	} else if right_sum >= left_sum && right_sum >= cross_sum {
		return right_low, right_high, right_sum
	} else {
		return cross_low, cross_high, cross_sum
	}
}

// FIND-MAX-CROSSING-SUBARRAY(A, low, mid, high)
//
// left-sum = INFINITESIMAL
// sum = 0
// for i = mid downto low
//     sum = sum + A[i]
//     if sum > left-sum
//         left-sum = sum
//         max-left = i
// right-sum = INFINITESIMAL
// sum = 0
// for j = mid + 1 to high
//     sum = sum + A[j]
//     if sum > right-sum
//         right-sum = sum
//         max-right= j
// return (max-left, max-right, left-sum + right-sum)
func FIND_MAX_CROSSING_SUBARRAY(A []int, low, mid, high int) (max_low, max_high, max_sum int) {
	LEFT_SUM := math.MinInt64
	SUM := 0
	MAX_LEFT := mid
	for i := mid; i >= low; i-- {
		SUM = SUM + A[i]
		if SUM > LEFT_SUM {
			LEFT_SUM = SUM
			MAX_LEFT = i
		}
	}

	RIGHT_SUM := math.MinInt64
	SUM = 0
	MAX_RIGHT := mid + 1
	for j := mid + 1; j <= high; j++ {
		SUM = SUM + A[j]
		if SUM > RIGHT_SUM {
			RIGHT_SUM = SUM
			MAX_RIGHT = j
		}
	}

	return MAX_LEFT, MAX_RIGHT, LEFT_SUM + RIGHT_SUM
}
