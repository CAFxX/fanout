package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }

func main() {
    var v1 int8 = 42
    var v2 string = "\n\t"
    v3 := make(map[uint32]bool)
    var v4 bool = true
    v5 := GenS1{F0: 2}
    switch "\n\t" {
    case "hello":
        v1 += -1
    case "foo bar":
        var v6 int64 = -1
        _ = v6
    default:
        v7 := []int{}
        _ = v7
    }
    v8 := GenS2{F0: -2147483648, F1: -9223372036854775808, F2: -1}
    v2 = "\n\t"
    var v9 bool = true
    fmt.Println(v4)
    var v10 float64 = float64(1e10)
    if v9 {
        v11 := GenS2{F0: -1, F1: 9223372036854775807, F2: -562}
        _ = v11
    }
    switch 4294967295 {
    case 128:
        v12 := []int{4294967295}
        _ = v12
    case 813:
        v1 = 105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
