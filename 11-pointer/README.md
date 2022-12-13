## Pass Argument by (Value | Reference)
arguments are passed by value in Go language. If you want to update something, you need to pass a reference.

taking the following code as an example, the zero function updates the value of the x variable, but because only the value of x is passed, the original value of x variable is not updated since there is no reference between the x in the zero function and the original x.
```go
func zero(x int) {
  x = 0
}

func main() {
  x := 5
  zero(x)
  fmt.Println(x) // x is still 5
}
```

By changing the zero function parameter to a pointer, the original x will be updated when you updated the reference.

```go
func zero(xPtr *int) {
  *xPtr = 0
}

func main() {
  x := 5
  zero(&x)
  fmt.Println(x) // x is 0
}
```

## The `&` and `*`
- `&`" use to obtain the address of a variable
- `*`: dereference a pointer variable



