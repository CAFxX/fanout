package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 int8 = 100
    var v2 bool = false
    var v3 bool = (v2 && v2)
    v4 := GenS1{F0: -83, F1: 9223372036854775807, F2: 2}
    var v5 uint64 = 1
    v2 = false
    v6 := []int{4294967295, 1, 405}
    v5 = 488
    v7 := GenS1{F0: 0, F1: 9223372036854775807, F2: 4294967295}
    if len(v6) > 0 {
        v6[0] = -475
    }
    switch -2147483648 {
    case -1:
        var v8 int16 = -166
        _ = v8
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 498 {
    case 973:
        v9 := GenS2{F0: "a", F1: v2}
        _ = v9
    default:
        var v10 bool = false
        _ = v10
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
