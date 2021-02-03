/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
const search = function (nums, target) {
    let left = 0, right = nums.length - 1;
    let ret = -1;
    while (left <= right) {
        const mid = left + Math.ceil((right - left) / 2);
        if (nums[mid] === target) {
            ret = mid
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

console.log(search([-1, 0, 3, 5, 9, 12], 9))