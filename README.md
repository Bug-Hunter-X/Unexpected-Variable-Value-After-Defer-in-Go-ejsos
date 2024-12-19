# Unexpected Variable Value After Defer in Go

This example demonstrates a common misconception regarding Go's `defer` keyword.  Many developers assume that a `defer` function modifies variables in the scope it's called from *before* the function returns. This is incorrect; `defer`'s execution happens *after* the main function concludes.

The code shows a variable `i` modified both inside the `defer` and outside. The final value of `i` printed shows the deferred modification doesn't change the value at the point where `defer` was called.