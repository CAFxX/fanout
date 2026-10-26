package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint64; F2 int64 }
func genFn0(p0 uint16, p1 int64) uint16 {
    v4 := GenS1{F0: "", F1: 2147483647, F2: 42}
    v5 := GenS1{F0: "\n\t", F1: -1, F2: 0}
    v6 := []int{1, 32767, -165, -574}
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    var v3 int = 127
    v8 := v2[(v3 / ((9223372036854775807) | 1))]
    v9 := GenS2{F0: 2, F1: 349, F2: 0}
    switch "hello" {
    case "foo bar":
        v10 := GenS2{F0: 645, F1: 42, F2: -1}
        _ = v10
    default:
        var v11 uint32 = 2147483647
        _ = v11
    }
    var v12 bool = (42 > 1000)
    v12 = true
    var v13 uint64 = 18446744073709551615
    fmt.Println(v8, v3, v3)
    var v14 string = ("" + "foo bar")
    if true {
        var v15 int16 = -98
        v12 = v12
        fmt.Println(v14)
        _ = v15
    } else {
        var v16 float32 = ((0.0 + 0.1) / (3.14159 - 0.0))
        v3 = v3
        _ = v16
    }
    v12 = (1e10 < 0.1)
    var v17 int = (v3 + 32768)
    v2[(v3 / ((v17) | 1))] = v12
    switch ((v14 + v14) + (v14 + v14)) {
    case "\n\t":
        v13 += v13
    default:
        v8 = (22 != 116)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
