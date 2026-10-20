package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }

func main() {
    var v1 float64 = 0.0
    var v2 float64 = 1.5
    var v3 int32 = -1
    v4 := make(map[uint32]uint32)
    var v5 uint32 = 599
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int16 = 1
        for fi8 := 0; fi8 < 12; fi8++ {
            v1 = 0.0
            var v9 uint32 = 2147483648
            if (127 != 127) {
                switch fi8 {
                case 127:
                    var v10 int8 = 103
                    _ = v10
                case 2147483647:
                    if (!false) {
                        v1 -= v2
                        var v11 bool = true
                        fmt.Println(v11, v11, v5)
                        v7 = v7
                        v12 := GenS1{F0: 65535, F1: 143}
                        _ = v11
                        _ = v12
                    }
                case 609:
                    var v13 float64 = (float64(((float32((1e10 + (v2 * 0.1))) * (0.1 + (1e10 + 1e10))) * float32(0.1))) + 1.5)
                    _ = v13
                }
                var v14 float64 = float64(0.1)
                _ = v14
            }
            _ = fi8
            _ = v9
        }
        v7 = -32768
        _ = fi6
        _ = v7
    }
    var v15 int16 = 1000
    v16 := v4[(v5 * v5)]
    v17 := GenS2{F0: "", F1: 841}
    var v18 int32 = (v3 - (v3 & v3))
    v19 := GenS2{F0: "hello", F1: 42}
    v20 := v4[(v5 & (v5 - (v5 >> ((805) & 63))))]
    if true {
        var v21 uint8 = 0
        _ = v21
    } else {
        var v22 bool = true
        v23 := []int{2147483647}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v18, fi24)
        if false {
            fmt.Println(fi24)
            v15 -= 1
        } else {
            delete(v4, (v5 - (v20 | 2147483647)))
            v5 = 1
            var v25 uint = 410
            _ = v25
        }
        v20 = v20
        _ = fi24
    }
    v16 = v16
    v26 := v4[(v16 << ((1) & 63))]
    v15 -= 908
    if ("a" < ("hello" + "foo bar")) {
        fmt.Println(v1, v16)
        v27 := GenS3{F0: (v15 / (((v15 << ((419) & 63))) | 1)), F1: 0, F2: 1}
        _ = v27
    }
    v28 := []int{-1}
    fmt.Println(v26)
    for fi29 := 0; fi29 < 14; fi29++ {
        var v30 bool = ((42 < 190) || (!(("a" + "hello") == "")))
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
