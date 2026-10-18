package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0() int {
    var v5 int64 = 972
    fmt.Println(v5)
    v6 := GenS1{F0: -1, F1: 0, F2: -669}
    if true {
        v5 += (v5 | (v5 & v5))
    } else {
        var v7 uint32 = 4294967295
        _ = v7
    }
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 42
    v2 := []uint32{}
    v3 := make(map[int]string)
    var v4 uint16 = v1
    var v8 uint16 = 255
    var v9 float64 = 1.5
    if false {
        v8 = 65535
    } else {
        v8 = v4
        v4 -= v4
    }
    var v10 uint32 = 65535
    var v11 int = 2147483647
    v4 -= 32767
    v3[v11] = ""
    v10 = (v10 | v10)
    v4 -= (^(v1))
    switch "hello" {
    case "\n\t":
        var v12 string = "foo bar"
        _ = v12
    default:
        v13 := GenS2{F0: 1, F1: v10}
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
