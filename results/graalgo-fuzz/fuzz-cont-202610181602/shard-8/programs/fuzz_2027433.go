package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
func genFn0(p0 bool, p1 int16, p2 int8) int8 {
    var v7 string = (("" + "hello") + "")
    for fi8 := 0; fi8 < 17; fi8++ {
        p2 += int8(fi8)
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p2 - (p2 / ((p2) | 1)))
}

func main() {
    var v1 uint8 = 42
    v2 := []int{1, 1}
    v3 := []uint32{4294967295, 65535, 2147483647, 2}
    var v4 uint = uint(v1)
    var v5 uint8 = v1
    var v6 uint32 = 225
    fmt.Println(v6)
    if true {
        var v10 int64 = -1
        var v11 uint64 = uint64(v1)
        v6 = (v6 * v6)
        _ = v10
        _ = v11
    }
    if (v1 >= v1) {
        switch "hello" {
        case "a":
            fmt.Println(v1, v4)
        }
        var v12 float64 = 0.1
        fmt.Println(v12, v4, v12)
        _ = v12
    }
    v1 = 42
    v6 += v6
    if len(v2) > 0 {
        v2[0] = -735
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint64 = 296
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int = (fi15 * 0)
        v4 = v4
        _ = fi15
        _ = v16
    }
    fmt.Println(v6, v5)
    v17 := GenS1{F0: 1000, F1: true, F2: v4}
    fmt.Println(v6)
    fmt.Println(v4, v1)
    v2 = append(v2, -918)
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 bool = (-7 >= -85)
        _ = fi18
        _ = v19
    }
    v3 = append(v3, v6)
    v4 = (v4 << (((v4 * (v4 | uint(v5)))) & 63))
    var v20 uint32 = (v6 / ((v6) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
