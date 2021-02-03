/**
 * Forward declaration of guess API.
 * @param {number} num   your guess
 * @param pick
 * @return                -1 if num is lower than the guess number
 *                         1 if num is higher than the guess number
 *                       otherwise return 0
 * var guess = function(num) {}
 */

const guess = function (num, pick = 3) {
    if (num === pick) return 0
    return num > pick ? -1 : 1
}

/**
 * @param {number} n
 * @return {number}
 */
const guessNumber = function (n) {
    let left = 1, right = n, ret
    while(left <= right) {
        let mid = left + Math.floor((right - left)/2)
        if (guess(mid) === 0) {
            ret = mid
            break
        }
        if(guess(mid) === -1) {
            right = mid -1
        } else {
            left = mid + 1
        }
    }
    return ret
};

console.log(guessNumber(3))