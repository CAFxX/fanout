package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint64 {
    var v5 int64 = -48
    var v6 float64 = 0.1
    v6 -= (float64(3.14159) - float64(((1.5 + ((0.0 + (1.5 + 0.1)) / 0.1)) - 0.0)))
    switch 168 {
    case -1:
        v5 -= (v5 + v5)
    case 128:
        fmt.Println(v6)
    default:
        v5 = -9223372036854775808
    }
    _ = v5
    _ = v6
    return 18446744073709551615
}
func genFn1() int32 {
    var v7 int64 = 1
    v8 := []int{-444, 400, 32768, 32768}
    _ = v7
    _ = v8
    return -340
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 int8 = 100
    var v4 float32 = float32(0.1)
    if ((833 < 255) && false) {
        v9 := GenS1{F0: 42}
        _ = v9
    }
    v10 := []int{32768, 939, 42, 1}
    v11 := GenS1{F0: 1}
    v12 := GenS1{F0: 32767}
    v13 := GenS1{F0: 791}
    fmt.Println(v3, v4)
    var v14 float32 = v4
    var v15 string = ("" + "\n\t")
    if true {
        v16 := GenS1{F0: 722}
        var v17 float64 = (0.1 + (1.5 / 0.1))
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
