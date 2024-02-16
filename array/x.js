var maxSubArray2 = function (nums) {
  let max = Math.min(-1);

  for (let i = 0; i < nums.length; i++) {
    for (let j = i; j <= nums.length; j++) {
      let subarr = nums.slice(i, j);

      let sum = 0;
      for (let k = 0; k < nums.length; k++) {
        sum = sum + subarr[k];
      }

      if (max < sum) {
        max = sum;
      }
    }
  }

  return max;
};
let arr = [-1];
let sum = maxSubArray2(arr);
console.log(sum, "max returns");
