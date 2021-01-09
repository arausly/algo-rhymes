/**
 *
 * @param {string} string of characters 
 * @param {number} n line of integers 
 */
const repeatedString = (s, n) => {
  let base = Math.floor(n / s.length);
  let remainder = n - (base * s.length);
  let res = 0
  for(let i = 0; i < s.length; i++){
      if(s[i] == 'a') res++;
  }
  res = res * base;
  for(let i = 0; i < remainder; i++){
      if(s[i] == 'a') res++;
  }
  return res;
};

module.exports = repeatedString;
