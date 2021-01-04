const staircase = require("./solution");

test("returns N stairs for N numbers", () => {
  let res = ' #' + '\n' + '##' + '\n';
  expect(staircase.staircaseSolution1(2)).toBe(res);
});

test("returns N stairs for N numbers", () => {
  let res = ' #' + '\n' + '##' + '\n';
  expect(staircase.staircaseSolution2(2)).toBe(res);
});

test("returns N stairs for N numbers", () => {
  let res = ' #' + '\n' + '##' + '\n';
  expect(staircase.staircaseSolution3(2)).toBe(res);
});