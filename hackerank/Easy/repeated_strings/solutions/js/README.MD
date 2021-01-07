# Question
[Link to hackerRank Question](https://www.hackerrank.com/challenges/repeated-string/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)

# Explanation
I actually found this question interesting and pretty straight forward, but I was also concious not to spend much time in going through the O(n) loop, i.e i wanted the algorithm to run faster than the normal solution that comes to mind.

That been said, I looked deeply and i noticed a pattern. I noticed I could get the actual answer while going through a minimal loop regardless of what N may be.

I divided N by the String length in order to get a Base. I just wanted to know how many times I have to repeat the string till I get to N then after that I wanted to know the remainder so that I can loop the remainder and know how many 'A's are in there. 

The summation of the above would give me my result and I didn't need to go through the whole loop if N was like a million or billion.