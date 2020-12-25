# Advanced Shades of Go

![alt](https://confusedcoders.com/wp-content/uploads/2016/10/golang.jpg)

These are the most complex Gotchas and needs good amount of digging to identify and address as they may not be directly visible in the program such as resource leak, blocking channel, deadlocks, routine leaks, infinite loop etc.

## iota-does-not-always-start-with-zero

At a first glance, it seems that iota identifier is an increment operator starts form zero and continues to increase from there, well this is not always true in Go. The value of iota is identified by an index operator for the current line in the constant declaration block, which means if the use of iota is not the first line in the constant declaration block the initial value may not be zero.

## array-function-arguments

As a Java developer, you may think, when you pass arrays to functions the functions reference the same memory location, so they can update the original data. But it's different in Go, arrays are values and does not share memory location, so when you pass arrays to functions. They get their own copy of the original array data. This can be a problem, if you are trying to update the array data.

## non-existing-map-keys

Map returns a zero-value of a type for non-existing keys. Checking for the appropriate zero value can be used to determine, if the map record exists or not, but it's not always reliable (e.g., what if you have a map of Boolean where the zero value is false). The most reliable way to know, if a given map record exists or not is to check the second value returned by the map access operation.

## hidden-data-in-slices

The new slice created out of an existing slice will reference the array of the original slice. This behavior can lead to unexpected memory usage, if your application allocates large slices and creates new slices from them to refer to small sections of the original data. This can be avoided by making a copy of the data needed from the temporary slice.

## append-can-cause-data-corruption-in-slice

Slices are backed by an array and capacity as the original. So, if you change an element in the slice, the original contents are modified as well. At some time adding data to one of the slices may result in a to a new array allocation, if the original array can't hold more data as a result other slices will point to the old array (with old data) whereas this can point to new instance of slice.

## unexpected-value-is-being-passed-in-differed-function

Arguments for a deferred function call are evaluated at the time of defer statement is evaluation and not at the time of function execution. You should make a use of pointer parameters to overcome this as pointer is saved at the time of defer statement is evaluation.

## wrong-use-of-defer-can-cause-resource-leak

You might use a defer in the code block and think that this will be called at the end of function to clean resources. This can eventually cause a resource leak if you are running a long for loop and calling defer in the code block to release resources. This can be avoided, by warping code block to function block or use direct clean-up instead of defer statement.

## why-execution-flow-is-blocked-for-channel

You might think the sender will not be blocked until your message is processed by the receiver. Depending on the machine where you are running the code, the receiver goroutine may or may not have enough time to process the message before the sender continues its execution.

## where-is-my-copy 

The copy function in the Go, copies minimum elements of the source to destination, to ensure correct coping, you should allocate sufficient destination slice. The copy number of elements copied. Alternatively, you can use append function in Go to copy array elements, make a note that, size (capacity) of the append slice will be larger than the length of slice.
