package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
func genFn0(p0 int, p1 uint16, p2 uint8) int {
    var v5 uint32 = 2147483648
    var v6 uint64 = 35
    fmt.Println(v6, p2, v6)
    var v7 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 int8 = -128
    v4 := []int{2147483647}
    if true {
        v8 := GenS1{F0: 1, F1: "\n\t", F2: (-(v3))}
        _ = v8
    }
    var v9 int64 = 1
    v3 = 1
    if false {
        _ = len(v1)
    } else {
        var v10 float64 = 0.1
        v11 := GenS1{F0: 2147483647, F1: (("hello" + ("" + "\n\t")) + "foo bar"), F2: v3}
        _ = len(v4)
        _ = v10
        _ = v11
    }
    switch 9223372036854775807 {
    case 185:
        v3 = v3
    case 9223372036854775807:
        _ = len(v4)
    }
    v12 := v2["hello"]
    fmt.Println(v12)
    fmt.Println(v9, v9, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
