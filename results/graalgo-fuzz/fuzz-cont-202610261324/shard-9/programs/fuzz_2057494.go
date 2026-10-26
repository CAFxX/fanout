package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 bool = false
    var v2 uint32 = 21
    v3 := []uint32{}
    var v4 int32 = -2147483648
    v5 := make(map[int]int)
    v1 = (true || ((-(v4)) != (v4 & v4)))
    v2 -= 65535
    var v6 bool = ("hello" != "foo bar")
    var v7 string = "\n\t"
    v1 = v1
    v8 := GenS1{F0: -128}
    if false {
        var v9 uint = 42
        _ = v9
    } else {
        v6 = (!(v2 >= 4294967295))
    }
    if ((!v6) || false) {
        v1 = v1
        var v10 float64 = 0.0
        _ = v10
    } else {
        fmt.Println(v6)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
