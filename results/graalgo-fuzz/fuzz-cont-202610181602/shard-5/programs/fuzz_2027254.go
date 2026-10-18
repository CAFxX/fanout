package main

import (
    "fmt"
)


func main() {
    var v1 bool = true
    v2 := []int{32767, -721}
    var v3 int16 = 42
    var v4 uint16 = 2
    switch "hello" {
    case "foo bar":
        v4 -= (v4 * (v4 + 432))
    case "":
        v4 -= v4
    case "\n\t":
        var v5 uint16 = (v4 | 0)
        _ = v5
    }
    fmt.Println(v3, v4)
    var v6 uint16 = (v4 + (^(v4)))
    _ = len(v2)
    var v7 uint16 = v6
    fmt.Println(v3, v4, v7)
    var v8 int = 42
    v3 -= 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
