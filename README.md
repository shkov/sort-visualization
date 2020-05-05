# sort-visualization

## What is it?

Visualization of Sorting Algorithms in Golang.
Sorts random shuffles of integers, with both speed and the number of items adapted to each algorithm's complexity.

## Getting started

#### 1. Installation

```shell script
$ go install github.com/shkov/sort-visualization/cmd/sort-visualization
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

#### 3. Sorting Algorithms

- [x] Bubble Sort
![](demo/bubble.gif)
- [x] Insertion Sort
![](demo/insertion.gif)
- [x] Selection Sort
![](demo/selection.gif)
- [ ] Merge Sort
- [ ] Heapsort
- [ ] Quick Sort
- [ ] Radix