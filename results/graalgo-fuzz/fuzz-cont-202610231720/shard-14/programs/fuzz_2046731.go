package main

import (
    "fmt"
)


func main() {
    v1 := []string{"a", "hello", ("hello" + "hello")}
    var v2 bool = (0.0 != (3.14159 + 3.14159))
    var v3 int = 42
    var v4 string = "foo bar"
    fmt.Println(v3, v2)
    fmt.Println(v3, v4)
    v2 = v2
    var v5 uint = 1
    var v6 bool = (("hello" + v4) == ("hello" + "a"))
    var v7 uint = v5
    switch v4 {
    case "\n\t":
        v5 = v7
    }
    var v8 uint16 = 785
    v3 += 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
