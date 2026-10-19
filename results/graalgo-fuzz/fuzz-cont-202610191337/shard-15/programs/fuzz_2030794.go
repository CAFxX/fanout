package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64; F1 uint }
func genFn0(p0 int16, p1 int32) uint16 {
    if false {
        var v6 int64 = 42
        _ = v6
    }
    switch "foo bar" {
    case "a":
        fmt.Println(p0)
    case "":
        v7 := GenS1{F0: p0, F1: 124}
        _ = v7
    case "\n\t":
        fmt.Println(p0)
    }
    v8 := []int{4294967295, 355, 64, -561}
    var v9 int64 = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 1.5
    v3 := make(map[uint32]int)
    var v4 uint = 65535
    var v5 int8 = 100
    v10 := []int{42, 9223372036854775807}
    v1 -= v1
    var v11 bool = (!false)
    v4 -= v4
    var v12 bool = (1 >= 42)
    if (((!(!v12)) && (162 == 197)) || (1e10 > 0.1)) {
        v4 += 640
        v13 := GenS1{F0: -896, F1: (v5 ^ v5)}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = v13
    } else {
        var v14 float32 = float32(v1)
        _ = v14
    }
    if v12 {
        v15 := GenS2{F0: 914, F1: "a"}
        _ = v15
    } else {
        var v16 float32 = 3.14159
        _ = v16
    }
    v17 := GenS3{F0: 1, F1: v4}
    fmt.Println(v4, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
