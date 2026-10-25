package main

import (
    "fmt"
)


func main() {
    var v1 bool = ((1.5 + -2.25) == 1e10)
    var v2 int8 = 1
    v3 := []string{("hello" + "foo bar"), "a", (("foo bar" + "") + ""), "hello"}
    var v4 float64 = ((((3.14159 * float64((-2.25 + 0.0))) + 0.1) * (1e10 + -2.25)) * ((0.1 + 0.1) * 3.14159))
    var v5 uint64 = 0
    var v6 float32 = ((float32((v4 - -2.25)) * 3.14159) - 1e10)
    fmt.Println(v2)
    var v7 float64 = (v4 - v4)
    v5 -= (v5 & (v5 % ((v5) | 1)))
    var v8 uint8 = 157
    if len(v3) > 0 {
        v3[0] = ("" + (("a" + "foo bar") + ""))
    }
    _ = len(v3)
    v1 = (v1 && true)
    if v1 {
        if (255 != v8) {
            fmt.Println(v8)
            fmt.Println(v1, v7, v5)
        }
    }
    var v9 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
