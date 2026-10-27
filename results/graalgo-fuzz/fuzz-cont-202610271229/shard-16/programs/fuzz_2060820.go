package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
func genFn0(p0 int32, p1 uint, p2 int) string {
    p1 += p1
    var v4 string = ("foo bar" + "")
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return v4
}
func genFn1(p0 uint16, p1 int32) uint8 {
    var v5 uint = 2
    var v6 uint8 = 134
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v6 - v6)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 32767
    v3 := make(map[int]uint32)
    v7 := GenS1{F0: "foo bar"}
    switch 9223372036854775807 {
    case 674:
        var v8 uint = 18446744073709551615
        _ = v8
    case 979:
        if true {
            var v9 float64 = ((float64(0.1) * 0.0) * (1.5 + 1e10))
            v9 -= 1e10
            switch -162 {
            case 128:
                v2 = (v2 & v2)
            case -9223372036854775808:
                var v10 int8 = -128
                _ = v10
            case 32767:
                v11 := GenS2{F0: (2 == 4294967295), F1: true, F2: -40}
                _ = v11
            }
            _ = v9
        }
    }
    v2 = -1
    var v12 uint32 = 1
    v2 = v2
    v12 += (v12 >> ((147) & 63))
    for fi13 := 0; fi13 < 19; fi13++ {
        v12 += v12
        var v14 uint16 = 1
        _ = fi13
        _ = v14
    }
    v15 := []int{1}
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int16 = v2
        _ = len(v15)
        var v18 int16 = v17
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v12, v12)
    var v19 int64 = -1
    var v20 int32 = -2147483648
    for fi21 := 0; fi21 < 9; fi21++ {
        fmt.Println(v12, v12, v20)
        v20 = (-(v20))
        _ = fi21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
}
