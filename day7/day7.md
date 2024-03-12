### 1. Container With Most Water

[code](max_area_test.go)

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i^th` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

Example 1:

![images](assets/question_11.jpg)

`Input`: height = [1,8,6,2,5,4,8,3,7]

`Output`: 49

`Explanation`: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Example 2:

`Input`: height = [1,1]
`Output`: 1

### Explanation Algorithm 
![images](assets/container.svg)

### 2. Max Number of K-Sum Pairs

You are given an integer array `nums` and an integer `k`.

In one operation, you can pick two numbers from the array whose sum equals `k` and remove them from the array.

Return the maximum number of operations you can perform on the array.

Example 1:

`Input`: nums = [1,2,3,4], k = 5

`Output`: 2

`Explanation`: Starting with nums = [1,2,3,4]:
- Remove numbers 1 and 4, then nums = [2,3]
- Remove numbers 2 and 3, then nums = []
  There are no more pairs that sum up to 5, hence a total of 2 operations.
  Example 2:

`Input`: nums = [3,1,3,4,3], k = 6

`Output`: 1

`Explanation`: Starting with nums = [3,1,3,4,3]:
- Remove the first two 3's, then nums = [1,4,3]
  There are no more pairs that sum up to 6, hence a total of 1 operation.

### Explanation Algorithm
![images](assets/max-operation.png)

### 3. Maximum Average Subarray I

[code](find_max_average_test.go)

You are given an integer array `nums` consisting of `n` elements, and an integer `k`.

Find a contiguous subarray whose **length is equal to** `k` that has the maximum average value and return this value. Any answer with a calculation error less than `10^-5` will be accepted.

Example 1:

`Input`: nums = [1,12,-5,-6,50,3], k = 4

`Output`: 12.75000

`Explanation`: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

Example 2:

`Input`: nums = [5], k = 1

`Output`: 5.00000
 
### Explanation Algorithm 
![images](assets/find-max-average.png)