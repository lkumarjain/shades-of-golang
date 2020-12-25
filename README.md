# Shades of Go

Go is a simple and powerful language, but unlike any other language it also has few Gotchas… Some of these mistakes are natural traps for programmers, especially if you are accustomed to some other language thought process.

![alt](https://cdn.evilmartians.com/front/posts/errors-in-go-from-denial-to-acceptance/cover-a684519.png)

> Gotcha is a valid construct in a system, program or programming language that works as documented but is counter-intuitive and almost invites mistakes because it is both easy to invoke and unexpected or unreasonable in its outcome.
Source: [Wikipedia](https://en.wikipedia.org/wiki/Gotcha_(programming))

I am trying to list couple of `Gotchas` along with `alternatives` to address them, I came across while learning Golang. I would love to hear `feedback` and `learn from you` in case I have missed some of them and you have encountered these while development… I have divided these Gotchas in to multiple categories based on the complexity.  

## Simplest

These are very simplest form of Gotchas as we can catch them at compile time. Compiler or linters (static analysis tools) complains about these mistakes and we can resolve them quickly. Some of them are:  

![alt](https://miro.medium.com/max/6250/1*oL0hjMQxldCYMKRb0a3yhg.png)

- unexpected semicolon or newline before `{`
- *** declared and `not used`
- imported and `not used`
- non-declaration statement `outside function body`
- no new variables on left side of `:=`
- non-name *** on left side of `:=`
- use of untyped `nil`
- invalid argument *** (type `map[string]int)` for `cap`
- cannot use `nil` as type `string` in assignment
- invalid operation: `s == nil` (mismatched types `string` and `nil`)
- cannot assign to `s[0]`
- `mtx.Unlock` undefined (type *** has no field or method `Unlock`)
- cannot call `pointer method` on ***
- syntax error: unexpected `++`
- fatal error: all `goroutines are asleep` - deadlock!

## Intermediate

These are little tricky as to fix as your program behave abruptly at runtime. We can address such Gotchas with little investigation because these are directly visible at runtime in form of unexpected behavior, panic etc.

![alt](https://confusedcoders.com/wp-content/uploads/2016/10/golang-1.jpg)

- variable-shadow
- zero-sized-variables
- entry-in-nil-map
- unexpected-values-in-range
- update-reference-item-in-range
- breaking-out-of-for-switch
- breaking-out-of-for-select
- data-race-closure-and-iteration-variable
- recovering-from-a-panic
- case-statement-is-not-executing
- empty-structs-in-the-json-text-output
- nil-is-not-equal-to-nil
- why-goroutine-execution-did-not-complete
- working-with-closed-channel
- invalid-memory-address-or-nil-pointer-dereference
- the-trim-function

## Advanced

These are the most complex Gotchas and needs good amount of digging to identify and address as they may not be directly visible in the program such as resource leak, blocking channel, deadlocks, routine leaks, infinite loop etc.

![alt](https://confusedcoders.com/wp-content/uploads/2016/10/golang.jpg)

- iota-does-not-always-start-with-zero
- array-function-arguments
- non-existing-map-keys
- hidden-data-in-slices
- append-can-cause-data-corruption-in-slice
- unexpected-value-is-being-passed-in-differed-function
- wrong-use-of-defer-can-cause-resource-leak
- why-execution-flow-is-blocked-for-channel
- where-is-my-copy

