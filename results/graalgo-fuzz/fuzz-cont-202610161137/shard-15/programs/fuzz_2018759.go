package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 9223372036854775807
    var v3 uint32 = 776
    var v4 uint64 = 120
    v5 := GenS1{F0: 42, F1: v2}
    delete(v1, "\n\t")
    v6 := GenS1{F0: -1, F1: v2}
    v7 := []int{}
    v8 := GenS2{F0: 127}
    if true {
        var v9 string = "hello"
        _ = v9
    }
    var v10 int16 = -146
    if false {
        fmt.Println(v4, v10, v10)
        fmt.Println(v3, v3)
    } else {
        v11 := v1[((("\n\t" + "hello") + "a") + "\n\t")]
        _ = v11
    }
    switch 833 {
    case 4294967295:
        v10 = (v10 - 1)
    default:
        fmt.Println(v4, v3)
    }
    v7 = append(v7, -242)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
