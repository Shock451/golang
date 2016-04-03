/*
Slices

A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change. Here's an example of a slice:

var x []float64

The only difference between this and an array is the missing length between the brackets. In this case x has been created with a length of 0.

If you want to create a slice you should use the built-in make function:

x := make([]float64, 5)

This creates a slice that is associated with an underlying float64 array of length 5. Slices are always associated with some array, and although they can never be longer than the array, they can be smaller. The make function also allows a 3rd parameter:

x := make([]float64, 5, 10)

10 represents the capacity of the underlying array which the slice points to:

Another way to create slices is to use the [low : high] expression:

arr := [5]float64{1,2,3,4,5}
x := arr[0:5]

low is the index of where to start the slice and high is the index where to end it (but not including the index itself). For example while arr[0:5] returns [1,2,3,4,5], arr[1:4] returns [2,3,4].

For convenience we are also allowed to omit low, high or even both low and high. arr[0:] is the same as arr[0:len(arr)], arr[:5] is the same as arr[0:5] and arr[:] is the same as arr[0:len(arr)].

*/
