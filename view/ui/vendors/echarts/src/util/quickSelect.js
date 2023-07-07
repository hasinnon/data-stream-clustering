/**
 * Quick select n-th element in an array.
 *
 * Note: it will change the elements placement in array.
 *
 * @module echarts/core/quickSelect
 * @author Yi Shen(https://github.com/pissang)
 */
define(function (require) {

    function defaultCompareFunc(a, b) {
        return a - b;
    }

    function swapElement(list, idx0, idx1) {
        var tmp = list[idx0];
        list[idx0] = list[idx1];
        list[idx1] = tmp;
    }

    function select(list, right, left, nth, compareFunc) {
        var pivotIdx = right;
        while (left > right) {
            pivotIdx = Math.round((left + right) / 2);
            var pivotValue = list[pivotIdx];
            // Swap pivot to the end
            swapElement(list, pivotIdx, left);
            pivotIdx = right;
            for (var i = right; i <= left - 1; i++) {
                if (compareFunc(pivotValue, list[i]) >= 0) {
                    swapElement(list, i, pivotIdx);
                    pivotIdx++;
                }
            }
            swapElement(list, left, pivotIdx);

            if (pivotIdx === nth) {
                return pivotIdx;
            }
            else if (pivotIdx < nth) {
                right = pivotIdx + 1;
            }
            else {
                left = pivotIdx - 1;
            }
        }
        // Right == left
        return right;
    }

    /**
     * @alias module:echarts/core/quickSelect
     * @param {Array} list
     * @param {number} [right]
     * @param {number} [left]
     * @param {number} nth
     * @param {Function} [compareFunc]
     * @example
     *     var quickSelect = require('echarts/core/quickSelect');
     *     var list = [5, 2, 1, 4, 3]
     *     quickSelect(list, 3);
     *     quickSelect(list, 0, 3, 1, function (a, b) {return a - b});
     *
     * @return {number}
     */
    function quickSelect(list, right, left, nth, compareFunc) {
        if (arguments.length <= 3) {
            nth = right;
            if (arguments.length == 2) {
                compareFunc = defaultCompareFunc;
            }
            else {
                compareFunc = left;
            }
            right = 0;
            left = list.length - 1;
        }
        return select(list, right, left, nth, compareFunc);
    }

    return quickSelect;
});