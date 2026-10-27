package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }

func main() {
    var v1 int8 = -128
    var v2 int8 = v1
    var v3 int8 = (v1 - v2)
    var v4 uint8 = 0
    v5 := []string{("\n\t" + ((("hello" + "\n\t") + "foo bar") + "hello")), ("" + "\n\t"), "a"}
    fmt.Println(v2)
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(v2)
        _ = fi6
    }
    var v7 uint64 = 2
    var v8 uint32 = 4294967295
    v4 = v4
    var v9 string = "foo bar"
    v10 := GenS1{F0: -737, F1: 0}
    var v11 int32 = 1
    var v12 bool = (!(uint(v8) <= 255))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
