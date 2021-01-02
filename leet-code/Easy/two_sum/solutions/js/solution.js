/**
 *
 * @param {Array<number>} numbers
 * @param {number} target
 * @returns {Array<number>}
 */
const twoSum = (numbers, target) => {
  for (let i = 0; i < numbers.length - 1; i++) {
    for (let j = i + 1; j < numbers.length; j++) {
      let n1 = numbers[i],
        n2 = numbers[j];
      const sum = n1 + n2;
      if (sum === target) {
        return [i, j];
      }
    }
  }
};

module.exports = twoSum;
