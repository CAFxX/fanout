package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    var v1 int16 = -1
    var v2 int = 689
    var v3 bool = true
    v4 := make(map[string]int)
    v1 += -1
    switch 1 {
    case 823:
        var v5 uint32 = 2147483647
        _ = v5
    default:
        var v6 uint = 943
        _ = v6
    }
    v2 -= v2
    v1 = (v1 & 1)
    v7 := []int{v2, v2, v2}
    var v8 int = (v2 % ((v2) | 1))
    if len(v7) > 0 {
        v7[0] = 32768
    }
    var v9 uint64 = 18446744073709551615
    var v10 float32 = float32((3.14159 / 1.0))
    _ = len(v4)
    v11 := GenS1{F0: 891, F1: v8}
    var v12 int16 = v1
    var v13 int32 = 42
    switch ("a" + ((("a" + "foo bar") + ("a" + (("" + ("foo bar" + (("a" + "a") + "foo bar"))) + "\n\t"))) + "\n\t")) {
    case "hello":
        v13 = v13
    default:
        v2 -= (v2 + v8)
    }
    _ = len(v7)
    var v14 int32 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
