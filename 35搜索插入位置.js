/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
const searchInsert = function (nums, target) {
    if (nums.length === 0) return 0
    const length = nums.length
    if (nums[length - 1] < target) return length
    if (nums[0] >= target) return 0
    let left = 0, right = length, ret
    while (left <= right) {
        let mid = left + Math.ceil((right - left) / 2)
        if (nums[mid] === target) {
            ret = mid
            break
        }
        if (nums[mid] < target && nums[mid + 1] > target) {
            ret = mid + 1
            break
        }
        if (nums[mid] < target) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return ret
};

console.log(searchInsert([1, 3, 5, 6], 0))