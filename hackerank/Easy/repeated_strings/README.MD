# Question
[Link to hackerRank Question](https://www.hackerrank.com/challenges/repeated-string/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)

# Explanation
Although this question is marked as easy, it can prove difficult if close attention to the input constraints is not considered in your algorithmic solution. I came up with three(3) different approaches shown only in my golang solution. For other languages, specifically javascript and python, I only showed the most optimal solution that passes all the test cases. 


The first approach had a quadratic (O(n^2)) time complexity. I decided that as long as the base string "s" is not up to the length "n", keep adding the characters (s) one after the other until it reached the target length (n)

see code snippet below: 

```
    r := s
    sArr = strings.Split(s, "")
	for s < n {
       for _, val := range sArr {
		     if len(s) == n  {
				 break
			 }
			 r += val
	   }
	}
```

This was a valid solution, however, given the input constraints (1 <= n <= 10 ^ 12) our current solution would timeout for really large values of n. This is because depending on the length of s and the value of n, I would have to loop at least twice to construct the repeated string.

My second approach was similar but nearly linear time complexity. 
see below code snippet:

```
	for int64(len(r)) < n {
			d := n - int64(len(r))
			if int(d) > len(s) {
				d = int64(len(s))
			}
			r += s[:d]
		}
```

This also is a valid solution but doesn't pass the time-limit set in some test cases. In this approach, I get the difference per iteration and append a slice of the base string with reference to the difference.

NOTE: a quick note here, for the two approaches above I focused on generating the repeated infinite string and then counting the occurrences of "a" in that repeated string.



My final and the optimal solution passing all possible test cases and totally disregards the approach of trying to generate the repeated string first. Instead, it counts the occurrences in the first string, then it multiplies it by the possible number of times the base string "s" will be repeated given n. And then adds any excess remainder if available.

see main code snippets below:

```
	    g := math.Floor(float64(n / l))
		m := n % l
           ...
           ...
           ...
        f *= int64(g)

// where m is the remainder
//       g is the number of times the length of the base string goes into the target "n"
//       f is the frequency of the letter "a"
```
