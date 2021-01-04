/**
 *
 * @param {number} n  Number of stairs
 */

function staircaseSolution1(n) {
  let spaceNum = 0;
  let space = '';
  let harsh = '';
  let res = '';
  for(let i = 0; i < n; i++){
      spaceNum = n - i;
      for(let j = 1; j < spaceNum; j++){
          space = space + ' ';
      }
      harsh = harsh + '#';
      res += space + harsh + '\n'
      space = '';
    }
    return res;
};

/** 
 * Looking at the function above, you would notice the above
 * runs for O(n)square. I think we can make this even run better
 */ 

function staircaseSolution2(n) {
  let spaceNum = 0;
  let space = '';
  let res = '';
  let harsh = '';
  for(let i = 0; i < n; i++){
    spaceNum = (n - 1) - i;
    space = ' '.repeat(spaceNum);
    harsh = harsh + '#';
    // console.log(space + harsh);
    res += space + harsh + '\n';
    space = '';
  }
  return res;
};

/** 
 * The function above runs faster (O(n)), However, I think we can 
 * manage memory even more.
 */ 

function staircaseSolution3(n) {
  let stairs = "";
  let harsh = "#";
  for (let i = 1; i <= n; i++) {
    stairs += `${" ".repeat(n - i)}${harsh.repeat(i)}\n`;
  }
  return stairs;
};

/** 
 * The function above still runs fast (O(n)) and manages memory better IMO.
 */ 

 module.exports = {
  staircaseSolution1,
  staircaseSolution2,
  staircaseSolution3
 }