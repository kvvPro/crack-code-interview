# crack-code-interview
Solve different tasks of code interview

## Two-sum
### Intuition
I guess that I need to use hashmap to avoid nested loops.

### Approach
First you need to add hashmap.
Second you need to decide what to store in this hashmap.
After that you will add a hashmap to store the difference between target and each element of List.
In main loop you need to iterate the List and try to find in hashmap the element of List. If it exists and not equal to itself (compare indices), you fave found the solution.

### Complexity
Time complexity:
O(n)O(n)O(n)

Space complexity:
O(n)O(n)O(n)