const binarySearch = function (arr, find) {
    let result, mid = Math.ceil(arr.length / 2)
    for (let i = 0; i < arr.length; i++) {
        if (arr[mid] > find) {
            mid = mid - 1
            continue
        }
        if (arr[mid] < find) {
            mid = mid + 1
            continue
        }
        if (arr[mid] === find) {
            result = mid
            break
        }
    }
    return result
}

console.time('start')
console.log(binarySearch([1, 2, 4, 6, 7, 10], 6))
console.timeEnd('start')