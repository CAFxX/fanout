package main

import (
    "fmt"
)

func genFn0() uint32 {
    if (false && false) {
        var v5 bool = true
        fmt.Println(v5)
        _ = v5
    }
    v6 := []int{-2147483648, 9223372036854775807}
    _ = v6
    return 954
}

func main() {
    var v1 int32 = 42
    var v2 float64 = -2.25
    var v3 int16 = 0
    var v4 int16 = -32768
    fmt.Println(v3, v2)
    fmt.Println(v2)
    v7 := []int{}
    var v8 bool = false
    var v9 bool = v8
    v9 = v8
    if len(v7) > 0 {
        v7[0] = 32768
    }
    v8 = (-2.25 <= 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
