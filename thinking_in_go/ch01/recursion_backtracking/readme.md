# Recursion and Backtracking

# 2.1 Introduction
> In this chapter, we will look at one of the important topics, which will be used in almost every chapter, and also its relative "backtracking"

# 2.2 What is Recursion?
> Any function which calls itself is called recursive. A recursive method solves a problem by a copy of itself to work 
on a smaller problem. This is called the recursion step can result in many more such recursive calls.

It is important to ensure that the recursion terminates. Each time the function calls itself with a slightly simpler
version of the original problem. The sequence of smaller problems must eventually converge on the base case.

# 2.3 Why Recursion?

 Recursion is a useful technique borrowed from mathematics. Recursive code is generally shorter and easier to write than iterative code. Generally, loops are turned into recursive functions when they are compiled or interpreted.

Recursion is most useful for tasks that can be defined in terms of similar subtasks. For example, sort, search,
and traversal problems often have simple recursive solutions.

# 2.4 Format of a Recursive Function

A recursive function performs a task in part by calling itself to perform the subtasks. At some point, the function encounters a subtask that it can perform without calling itself. 

This case, where the function does not recur, is called the base case. The former
, where the function calls itself to perform a subtask, is referred to as the recursive
 case.
 
 # 2.6 Recursion versus Iteration
 
Recursion
 - Terminates when a base case is reached.
 - Each recursive call requires extra space on the stack frame(memory).
 - If we get infinite recursion, the program may run out of memory and result in stack
  overflow.
 - Solutions to some problems are easier to formulate recursively.
 
Iteration
- Terminates when a condition is proven to be false.
- Each iteration does not require extra space.
- An infinite loop could loop forever since there is no extra memory being created.
- Iterative solutions to a problem may not always be as obvious as a recursive solution.
 
# 2.10 What is Backtracking
> Backtracking is an improvement of the brute force approach. It systematically searches for a solution to a problem
  among all available options. 

What’s interesting about backtracking is that we back up only as far as needed to reach a previous decision point with an as-yet-unexplored alternative. 

In general, that will be at the most recent decision point. 

Eventually, more and more of these decision points will have been fully explored, and we will have to backtrack further and further.

If we backtrack all the way to our initial state and have explored all alternatives from there, we can conclude the particular problem is unsolvable. In such a case, we will have done all the work of the exhaustive recursion and known that there is no viable solution possible.

