package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 uint32, p1 int64, p2 uint64) int8 {
    if false {
        p0 -= p0
    } else {
        if true {
            var v5 float64 = 1.5
            _ = v5
        }
    }
    var v6 int32 = -54
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 float32 = ((0.1 / (3.14159 - (1e10 + 1e10))) + 1.5)
            var v9 string = "a"
            _ = fi7
            _ = v8
            _ = v9
        }
    case "foo bar":
        var v10 uint8 = 42
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    return 127
}
func genFn1(p0 uint64, p1 int64) uint8 {
    v12 := GenS1{F0: 1, F1: 127, F2: ((^(p0)) <= p0)}
    p0 -= 42
    var v13 uint32 = uint32(p0)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 131
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 18446744073709551615
    var v3 string = ""
    var v4 uint = v2
    var v14 uint64 = 4294967295
    var v15 uint = (v4 / (((v2 * 4294967295)) | 1))
    v15 += (v2 & (v15 / ((582) | 1)))
    v3 = (v3 + "foo bar")
    fmt.Println(v14, v4, v2)
    _ = len(v1)
    v16 := []int{}
    var v17 uint32 = 122
    switch 206 {
    case 675:
        fmt.Println(v14, v17)
    }
    v3 = "a"
    fmt.Println(v4, v2, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
