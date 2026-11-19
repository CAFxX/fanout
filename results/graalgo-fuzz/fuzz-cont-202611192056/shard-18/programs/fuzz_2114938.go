package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }

func main() {
    var v1 string = "hello"
    var v2 bool = (394 > 0)
    var v3 int16 = 1
    var v4 int64 = 0
    fmt.Println(v4, v3)
    for fi5 := 0; fi5 < 10; fi5++ {
        v3 = v3
        var v6 int8 = 78
        v7 := GenS1{F0: v1, F1: (0.0 <= 0.1)}
        _ = fi5
        _ = v6
        _ = v7
    }
    if v2 {
        v8 := []int{557, 1, -1}
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    fmt.Println(v3, v2, v4)
    var v10 uint8 = 0
    v1 = "foo bar"
    switch v1 {
    case "foo bar":
        var v11 int8 = -28
        _ = v11
    case "\n\t":
        v12 := []int{1, -790}
        _ = v12
    case "":
        v13 := []int{978, 1, 42}
        _ = v13
    default:
        v14 := []int{-1, 127, 128, -675}
        _ = v14
    }
    var v15 int32 = 2147483647
    v10 += v10
    v16 := []int{-101, -1}
    switch "\n\t" {
    case "\n\t":
        v17 := GenS1{F0: (v1 + v1), F1: false}
        _ = v17
    default:
        v1 = v1
    }
    var v18 int16 = 746
    fmt.Println(v2, v10, v4)
    fmt.Println(v4, v3, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
