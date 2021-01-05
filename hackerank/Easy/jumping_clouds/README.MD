# Question
*[Link to hackerRank](https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)



 ## Explanation
My solution here in golang shows three main thought trends. First, I try to generate an array with just the indices, then I find the minimum jumps. The second approach was to work with the given array "c" directly, rather than getting the indices. 

NOTE: the first two approaches are both of linear time-complexity.
``` O(n) + O(n) === O(2n) === O(n)```

The final approach drew learnings from the previous solutions. I realized that I could focus on the given array "c" (idea from the second solution) and particularly only on the zeros.  to explain further, I iterated through the array, and for every iteration, I tried to make the longest jump possible (2). if ```n + 2``` existed. I disregarded the next element and counted that as a valid jump. If ```n + 2``` didn't exist I will add the next element. 