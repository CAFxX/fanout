package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint32 }
func genFn0(p0 int8, p1 int32) string {
    if false {
        v7 := GenS1{F0: 0, F1: 36, F2: 4294967295}
        _ = v7
    }
    var v8 float32 = 1.5
    switch "a" {
    case "":
        var v9 string = ""
        _ = v9
    case "a":
        var v10 uint8 = 60
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    return "foo bar"
}

func main() {
    var v1 int16 = -1
    var v2 float64 = (1.5 / 1.0)
    v3 := []uint32{497, 1, 1}
    v4 := make(map[int]int)
    v5 := make(map[int]uint32)
    v6 := []int{}
    fmt.Println(v1, v1)
    if (1 >= 153) {
        v2 -= v2
        v1 = v1
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: v1, F1: 18, F2: 4294967295}
            v13 := v4[fi11]
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        v1 += v1
    }
    var v14 uint32 = 2147483647
    fmt.Println(v14, v1)
    var v15 int8 = -85
    var v16 uint8 = 0
    if ((0.1 / v2) <= (v2 * (1e10 - v2))) {
        var v17 uint16 = 813
        _ = v17
    }
    v16 -= 127
    v14 -= 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
