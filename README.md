# CountInversions
Takes an integer slice, returns the slice sorted and the number of inversions in the slice.
An inversion is defined as all instances of slice[i] > slice[j] while i < j. The number of inversions is a measure of how unsorted the slice is, hence useful for e.g. finding numberical similarity between two ranked lists. The higher the count, the more unsorted the slice.
To work with arrays, simply supplement fixed sizes in the definitions of slices.
## Sample Input
```
slice = [1, 88, 54, 32, 16, 2, 9]
```

## Sample Output
```
[1 2 9 16 32 54 88], 14
```