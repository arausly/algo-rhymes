const repeatedString = require("./solution");

test("returns 2 valleys for the path entries", () => {

  expect(repeatedString('aba', 10)).toBe(7);
});

test("returns 1 valley for the path entries", () => {
  expect(repeatedString('a', 1000000000000)).toBe(1000000000000);
});
