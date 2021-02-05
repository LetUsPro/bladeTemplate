/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
const searchRange = function (nums, target) {
    const length = nums.length
    if (length === 0) return [-1, -1]
    const left = searchFirstPosition(nums, target)
    if (left === -1) return [-1, -1]

    const right = searchLastPosition(nums, target)
    return [left, right]
};

const searchFirstPosition = function (nums, target) {
    const length = nums.length
    let left = 0, right = length - 1, ret = -1
    while (left < right) {
        let mid = left + Math.floor((right - left) / 2)
        if (nums[mid] < target) {
            left = mid + 1
        } else {
            right = mid
        }
    }
    if (nums[left] === target) {
        ret = left
    }
    return ret
}

const searchLastPosition = function (nums, target) {
    const length = nums.length
    let left = 0, right = length - 1
    while (left < right) {
        let mid = left + Math.ceil((right - left) / 2)
        if (nums[mid] > target) {
            right = mid - 1
        } else {
            left = mid
        }
    }
    return left
}

console.log(searchRange([5, 7, 7, 8, 8, 10], 6))