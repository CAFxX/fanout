package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16, p1 uint8) int32 {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{-1, (fi7 & (fi7 ^ fi7)), int(p0), fi7}
        switch fi7 {
        case 32768:
            fmt.Println(p1, fi7, fi7)
        default:
            v9 := GenS1{F0: p1}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(p1, p1)
    var v10 uint8 = 191
    var v11 uint8 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 858
}

func main() {
    var v1 uint16 = 42
    var v2 int32 = -2147483648
    v3 := make(map[string]int)
    v4 := []uint32{2147483647, uint32(v1)}
    var v5 int16 = 1
    v6 := make(map[uint32]int)
    delete(v3, "foo bar")
    var v12 uint = 65535
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := v6[926]
        var v15 uint32 = 4294967295
        _ = fi13
        _ = v14
        _ = v15
    }
    if (false && true) {
        var v16 int8 = 109
        _ = v16
    } else {
        fmt.Println(v12, v1, v12)
        v5 = 1
        if (true || false) {
            v17 := GenS1{F0: 0}
            _ = v17
        } else {
            v5 = (v5 | v5)
            _ = len(v3)
        }
    }
    var v18 uint32 = 42
    var v19 uint16 = (^(v1))
    _ = len(v4)
    var v20 uint32 = (v18 / ((42) | 1))
    var v21 float32 = (-2.25 - 3.14159)
    v1 -= v19
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 19; fi23++ {
            _ = len(v6)
            if true {
                v24 := GenS2{F0: 42}
                fmt.Println(v1, v21, v5)
                _ = v24
            }
            _ = fi23
        }
        _ = fi22
    }
    var v25 float32 = v21
    var v26 float32 = (v25 / float32(0.1))
    fmt.Println(v19)
    v26 = float32(-2.25)
    v20 -= (^(v20))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
