# Intermediate Shades of Go

![alt](https://confusedcoders.com/wp-content/uploads/2016/10/golang-1.jpg)

These are little tricky as to fix as your program behave abruptly at runtime. We can address such Gotchas with little investigation because these are directly visible at runtime in form of unexpected behavior, panic etc.

## variable-shadow

This is a very common trap even for experienced Go developers. It's easy to make and it could be hard to spot. You can make use of linters `vet` or [go-nyet](https://github.com/barakmich/go-nyet) `for more aggressive checks` to detect such errors in the program.

## zero-sized-variables

Shouldn’t two different variables have different addresses? Well, it's not the case with Go 😊, if you have zero-sized variables they might share the exact same address in memory.

## entry-in-nil-map

A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic; don't do that. Always initialize a map, using the built-in `make` function.

## unexpected-values-in-range

Another most common gotcha for the developer accustomed to a foreach loop in Java. In Go, the range clause is different, it generates two values: first value is the index and the second one is data.

## update-reference-item-in-range

The data values generated in the range clause are copies of the actual collection elements, these does not reference original item. Which means that updating the values will not change the original data. You should make use of index operator on collection to update original value.

## breaking-out-of-for-switch

The break statement in the switch moves out of switch and not for loop. If you want to move out of for loop, one option is return but this may not be a solution every time another alternative is to use label in break statement (similar to goto statements)

## breaking-out-of-for-select

In Go, select works similar to switch statement only difference is case statements should have a channel. The break statement in the select moves out of select and not for loop. If you want to move out of for loop, one option is return but this may not be a solution every time another alternative is to use label in break.

## data-race-closure-and-iteration-variable

This is yet another most common pitfall in Go, the iteration variables in for statements are reused across iterations. Which means that each closure created in the for loop will reference the same variable. You can solve this by creating a local variable in the loop or other solution could be to pass variable as a parameter in the closure function. We can detect this early in the program by using [Go Race Detector](https://blog.golang.org/race-detector).

## recovering-from-a-panic

The recover function can be used to catch/intercept a panic. Calling recover will do the trick only when it's done in a deferred function. The call to recover works only if it's called directly in your deferred function.

## case-statement-is-not-executing

You might think we need to add break statement in the switch to move out, well this is quite different in Go, we don’t need to add break in the case block, as a result fall-through does not work as expected. You can force this in the case blocks to fall through by using the fall-through statement at the end of each case block or rewrite switch statement to use expression lists in the case blocks.

## empty-structs-in-the-json-text-output

Only exported fields of a Go struct will be present in the JSON output.  You can specify the JSON field name explicitly with a `json: tag`.

## nil-is-not-equal-to-nil

In Go, an interface equals to another interface, only if the concrete value and dynamic type are both nil. The same applies to nil value. You can think of the interface value nil as typed, and nil without type doesn’t equal nil with type. If we convert nil to the correct type, the values are indeed equal.

## why-goroutine-execution-did-not-complete

In Go, application does not wait for all goroutines to complete before it exists. This is a most common mistake done by a developer in early days of learning Go. We can avoid this using WaitGroup a most common solution, it allows us to wait in the main for goroutines to complete. Another solution is to pass goroutine execution state using channel.

## working-with-closed-channel

Receiving from a closed channel is safe, whereas writing on closed channel throws a panic. Second value received from the channel, indicates that is there more moves to be received or not. This is a well-documented behavior, but it's not very intuitive for new Go developers who might expect the send behavior to be similar to the receive behavior. This is complex enough and needs quite a good amount of thought as this may be resolved with minor code change or may need change in design.

## invalid-memory-address-or-nil-pointer-dereference 

Go developers many times faces the issue of [dereferencing](https://lkumarjain.blogspot.com/2020/01/why-calling-method-on-nil-struct.html) of a `nil` pointer. An uninitialized pointer is nil, and you can’t easily follow the nil pointer. If `x is nil`, an attempt to evaluate `*x` will cause a run-time panic.

## the-trim-function 

The [Trim](https://golang.org/pkg/strings/#Trim), [Trim Left](https://golang.org/pkg/strings/#TrimLeft) and [Trim Right](https://golang.org/pkg/strings/#TrimRight) functions strip all Unicode code points contained in a cut-set. To strip a trailing string, you should make use of [Trim Suffix](https://golang.org/pkg/strings/#TrimSuffix) or [Trim Prefix](https://golang.org/pkg/strings/#TrimPrefix).
