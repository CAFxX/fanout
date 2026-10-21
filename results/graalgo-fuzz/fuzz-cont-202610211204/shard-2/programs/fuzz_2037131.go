package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int64; F2 int64 }
type GenS4 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{42, 2147483647, 4294967295}
    var v3 uint = 506
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    switch "\n\t" {
    case "":
        var v7 float64 = (1e10 + (float64(((0.1 + 1e10) / 1.0)) * ((3.14159 + float64(float32(0.1))) * 0.0)))
        _ = v7
    case "foo bar":
        v3 = (^(v3))
    default:
        var v8 uint32 = 145
        _ = v8
    }
    v3 = v3
    if (42 != 1) {
        var v9 uint = v3
        _ = v9
    }
    fmt.Println(v3)
    v10 := GenS1{F0: 65535, F1: 671}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int64 = -1
        v13 := GenS2{F0: 0}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint64 = 0
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = v3
        _ = fi15
    }
    fmt.Println(v14, v14)
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v14)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi16
    }
    v17 := GenS3{F0: 127, F1: 1, F2: 1}
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 float32 = float32(0.1)
        if (!((-2.25 + v19) < v19)) {
            v20 := GenS1{F0: 1, F1: 2}
            var v21 int16 = 472
            _ = v20
            _ = v21
        } else {
            var v22 int8 = 42
            v23 := GenS4{F0: v3, F1: "\n\t"}
            var v24 int16 = 42
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi18
        _ = v19
    }
    if false {
        var v25 uint32 = 202
        v14 = v14
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
