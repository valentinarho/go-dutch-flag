# go-dutch-flag
Simple implementation of the Dutch national flag problem in Golang.

## The problem

_source: [Wikipedia](https://en.wikipedia.org/wiki/Dutch_national_flag_problem)_

The Dutch national flag problem is a computer science programming problem proposed by Edsger Dijkstra. The flag of the Netherlands consists of three colors: red, white and blue. Given balls of these three colors arranged randomly in a line (the actual number of balls does not matter), the task is to arrange them such that all balls of the same color are together and their collective color groups are in the correct order.

This problem can also be viewed in terms of rearranging elements of an array. Suppose each of the possible elements could be classified into exactly one of three categories (bottom, middle, and top). The problem is then to produce an array such that all "bottom" elements come before (have an index less than the index of) all "middle" elements, which come before all "top" elements.


## Usage

Run `main.go` for a usage example.

DutchFlag function takes in input the array to sort and the partitioning function.
This function should be a function from int to int and return, given an integer item
the corresponding partition (-1, 0, 1).

The first example in `main.go` uses the builtin partitioning function that works on
arrays containing only elements in {1,2,3}.

The second example implements a custom partitioning function that split an array in
even, zeros and odd items.

It is possible to tailor the code in order to handle different data types in the array.
It is only necessary to modify in `dutch.go`:

* line 5: the type of the `elem` parameter in partitioningFunction
* line 8: the type of the `a` parameter
* line 19: the type of the `a` parameter

And then, create your `main.go` and your partitioning function!
