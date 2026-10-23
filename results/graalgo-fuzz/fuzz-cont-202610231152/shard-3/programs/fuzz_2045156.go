package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint16; F1 string; F2 string }
func genFn0() uint64 {
    v5 := []int{127, -296, 32767, 494}
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int8 = 42
        for fi8 := 0; fi8 < 13; fi8++ {
            v7 += v7
            v7 -= 2
            fmt.Println(fi6)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    var v9 uint32 = 64
    v10 := GenS1{F0: true, F1: 42}
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int64 = -1
        var v13 uint32 = (^(v9))
        v14 := GenS1{F0: (!true), F1: 826}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v5
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2}
    var v2 int = 128
    v3 := make(map[int]string)
    var v4 uint16 = 2
    v1 = append(v1, 4294967295)
    if true {
        v15 := GenS2{F0: 305, F1: 0, F2: 32767}
        _ = v15
    } else {
        var v16 string = ""
        var v17 int8 = 42
        v17 = -105
        _ = v16
        _ = v17
    }
    if false {
        var v18 float32 = float32((0.0 * 1.5))
        _ = v18
    } else {
        fmt.Println(v2)
    }
    v4 = v4
    for fi19 := 0; fi19 < 4; fi19++ {
        var v20 int16 = 32767
        v20 += (v20 << ((1) & 63))
        _ = fi19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS3{F0: v4, F1: "\n\t", F2: "a"}
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v2, v4)
    v2 += v2
    _ = len(v3)
    delete(v3, v2)
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 float64 = float64(0.0)
        if true {
            v25 := GenS2{F0: 928, F1: genFn0(), F2: 1000}
            _ = v25
        }
        var v26 float32 = 0.0
        _ = fi23
        _ = v24
        _ = v26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
