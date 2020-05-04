# sort-visualization

## What is it?

Visualization of Sorting Algorithms in Golang.
Sorts random shuffles of integers, with both speed and the number of items adapted to each algorithm's complexity.

![](demo/example.gif)

## Getting started

#### 1. Installation

```shell script
$ go install github.com/shkov/sort-visualization
```

#### 2. Usage

```shell script
$ sort-visualization -refresh=1ms -algorithm=bubble
```

```shell script
$ sort-visualization --help 

Usage of sort-visualization:
  -algorithm string
    	sorting algorithm type (default "bubble")
  -refresh duration
    	bar chart refresh interval (default 5ms)
```