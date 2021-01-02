# Question
 *[Link to hackerRank](https://www.hackerrank.com/challenges/counting-valleys/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup)


 ## Explanation
My solution here in three languages is pretty straightforward and follows a similar approach. The question requires a total count of the number of valleys created by the path given. I will walk you through the mental leaps I had to take to arrive at my linear time complexity solution. 


Firstly, in the context of this question, what is a valley? A valley is created when the hiker moves below sea-level and touches back to the sea-level. Imagine the sea-level to be at point 0 on the vertical rule scale. A movement towards path "D" should cause a drop to about -1 for 1 unit movement and vice-versa a movement towards path "U" should raise to about 1 for 1 unit movement. 


Having this mental picture made it relatively easy to associate these ideas with variables. The "depthTracker" variable for example was declared to act as a watchman to ensure that for every deep in sea-level there was a commensurate touch to sea-level else it would not be counted as a valley. 
 
