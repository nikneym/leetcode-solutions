/*
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/
 */

const assert = require("assert");

function rangeBitwiseAnd(left, right) {
  let res = left;
  for (let i = left; i < right; i++) {
    res &= i;
  }

  return res;
}

assert(rangeBitwiseAnd(5, 7) === 4);
assert(rangeBitwiseAnd(0, 0) === 0);
assert(rangeBitwiseAnd(1, 2147483647) === 0);
