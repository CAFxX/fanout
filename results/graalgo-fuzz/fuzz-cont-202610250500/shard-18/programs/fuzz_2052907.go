package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
func genFn0() uint8 {
    v7 := GenS1{F0: 0, F1: 128, F2: 127}
    var v8 int = -1
    var v9 int = (v8 ^ 790)
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1() int64 {
    var v10 float32 = 1.5
    fmt.Println(v10)
    v11 := GenS1{F0: 128, F1: 1, F2: 1}
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int = 519
    var v2 uint64 = 1
    var v3 float32 = ((((0.1 - 3.14159) / 1.0) + 1.5) + (0.0 - 1e10))
    var v4 string = ("foo bar" + ("" + ""))
    var v5 bool = true
    v6 := []int{(v1 - v1), v1}
    v1 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 = v1
    switch "\n\t" {
    case "\n\t":
        var v12 uint = 0
        _ = v12
    case "hello":
        var v13 float64 = 1.5
        _ = v13
    case "foo bar":
        v2 = v2
    default:
        var v14 uint = 18446744073709551615
        _ = v14
    }
    v2 = (v2 ^ 14)
    _ = len(v6)
    var v15 uint8 = 0
    _ = len(v6)
    var v16 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
