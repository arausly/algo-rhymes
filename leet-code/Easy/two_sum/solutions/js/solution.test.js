const twoSum = require("./solution");

test("should return 0 and 1 for numbers provided", () => {
  const numbers = [2, 7, 11, 15];
  const target = 9;
  const indices = twoSum(numbers, target);
  const expected = [1, 0];
  expect(indices).toEqual(expect.arrayContaining(expected));
});

test("should return 0 and 1 for arrays with duplicate numbers", () => {
  const numbers = [3, 3, 3];
  const target = 6;
  const indices = twoSum(numbers, target);
  const expected = [0, 1];
  expect(indices).toEqual(expect.arrayContaining(expected));
});
