package main

import "fmt"

func twoSum(nums []int, target int) []int {
    // Create a map to store the numbers and their indices
    numMap := make(map[int]int)

    // Iterate through the nums array
    for i, num := range nums {
        // Calculate the complement
        complement := target - num

        // Check if the complement exists in the map
        if index, ok := numMap[complement]; ok {
            // If it exists, return the current index and the index of the complement
            return []int{index, i}
        }

        // If it doesn't exist, add the current number and its index to the map
        numMap[num] = i
    }

    // Return an empty array if no solution is found
    return []int{}
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result)
}
