const countingValleys = require("../solutions/solution");

test("returns 2 valleys for the path entries", () => {
  expect(countingValleys(12, "DDUUDDUDUUUD")).toBe(2);
});

test("returns 1 valley for the path entries", () => {
  expect(countingValleys(8, "UDDDUDUU")).toBe(1);
});
