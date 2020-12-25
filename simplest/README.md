# Simplest Shades of Go

![alt](https://miro.medium.com/max/6250/1*oL0hjMQxldCYMKRb0a3yhg.png)

These are very simplest form of Gotchas as we can catch them at compile time. Compiler or linters (static analysis tools) complains about these mistakes and we can resolve them quickly. Some of them are:  

## unexpected semicolon or newline before `{` 

Unlike any other language, we are not allowed to place opening brace in the separate line. You can thank automatic semicolon injection.

## *** declared and not used 

Go does not allow you to have to unused variables in inside function, but is ok to have unused global variables.

## imported and not used 

Go does not allow to have unused imports, if you really want to import one for initialization of drivers or something else, you can make use of `_` as its package name to avoid compilation failure.

## non-declaration statement outside function body 

Go does not allow us to declare a variable using short variable declarations outside functions.

## no new variables on left side of `:=` 

Unlike any other language, go does not allow us to redeclare a variable in the standalone statement, but it is allowed in the multi-variable declarations where at least one new variable is present.

## non-name *** on left side of `:=` 

As per Rob Pike in [Issue #6842](https://github.com/golang/go/issues/6842) The `:=` notation is a shorthand for common cases. It's not meant to cover every possible declaration one may write.

## use of untyped `nil` 

If you don't specify the variable type the compiler will fail to compile your code because it can't identify the type of variable to define its zero-value as `nil`.

## invalid argument *** (type `map[string]int`) for `cap` 

Map capacity can be specified at the time of creation using `make(map[string]int, 50)`, but can't use cap function.

## cannot use `nil` as type `string` in assignment 

Unlike Java we can’t create a `nil` string in the Go.

## invalid operation: `s == nil` (mismatched types `string` and `nil`) 

As we can’t create nil string in Go, compilation will fail if we try to compare a string value with the `nil`

## cannot assign to `s[0]` 

Go strings are sequence (read-only slice) of bytes as a result these can’t be changed (immutable like Java). In case you want to change you should convert string to byte slice change and convert back to string.

## `mtx.Unlock` undefined (type *** has no field or method Unlock) 

Defining new type from an existing (non-interface) type does not inherit methods define on that struct. If you need these methods either use type embedding with anonymous name or initialize variable using actual struct and assign this to new type.

## cannot call pointer method on *** 

The operand must be addressable [composite literal](http://golang.org/ref/spec#Composite_literals). In Go not every variable is addressable though, a map entry cannot be addressed (as its address might change during map growth / shrink).

## syntax error: unexpected `++` 

Unlike Java, Go does not support prefix version of increment and decrement operation. Also Go does not allow use of these operators in the expression.

## fatal error: all goroutines are asleep - `deadlock!` 

Send and receive operations on a nil channel block forever and gives fatal error: all goroutines are asleep - deadlock! Inspite of a well-documented behaviour, this can surprise for new developers. To avoid such scenario, we should initialize a channel before using this.
