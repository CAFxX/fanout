package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() string {
    switch "\n\t" {
    case "a":
        var v6 int32 = -249
        _ = v6
    case "\n\t":
        var v7 int8 = -121
        _ = v7
    case "hello":
        var v8 int8 = -128
        _ = v8
    default:
        var v9 int = -1
        _ = v9
    }
    v10 := GenS1{F0: 42}
    var v11 int32 = 42
    fmt.Println(v11)
    _ = v10
    _ = v11
    return "a"
}
func genFn1() int8 {
    var v12 int32 = -184
    if (0 > 128) {
        var v13 string = ""
        v13 = v13
        var v14 float32 = 3.14159
        _ = v13
        _ = v14
    } else {
        v12 = -759
        v12 += 993
    }
    var v15 float64 = (-2.25 - 3.14159)
    var v16 uint = 2
    _ = v12
    _ = v15
    _ = v16
    return -56
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]int)
    var v3 uint16 = 65535
    var v4 uint64 = 4294967295
    v5 := []uint32{}
    var v17 uint8 = 127
    v17 = 128
    v4 -= (v4 & (v4 | v4))
    var v18 uint32 = 515
    fmt.Println(v17)
    var v19 int16 = 450
    v20 := GenS1{F0: -1}
    var v21 uint32 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
