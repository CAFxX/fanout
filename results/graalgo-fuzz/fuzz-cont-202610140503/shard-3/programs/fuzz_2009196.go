package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 uint8) uint32 {
    var v6 int16 = 32767
    fmt.Println(v6, p0)
    _ = p0
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 790
    var v3 uint16 = 2
    v4 := []int{176}
    var v5 string = "hello"
    _ = len(v4)
    v3 = (^(v3))
    var v7 float32 = ((1e10 * 0.1) - (float32(-2.25) / 1.5))
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint8 = 72
        _ = fi8
        _ = v9
    }
    if false {
        if ("\n\t" != v5) {
            var v10 string = v5
            v11 := GenS1{F0: v10, F1: true, F2: v10}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 8; fi12++ {
                switch fi12 {
                case 127:
                    _ = len(v4)
                case 4294967295:
                    v13 := v1[v2]
                    _ = v13
                default:
                    v7 -= (v7 + (v7 / v7))
                }
                var v14 int32 = 2147483647
                _ = fi12
                _ = v14
            }
            fmt.Println(v5)
            v15 := GenS2{F0: 9223372036854775807, F1: 31}
            _ = v15
        }
        _ = len(v1)
        for fi16 := 0; fi16 < 17; fi16++ {
            _ = len(v1)
            _ = fi16
        }
    }
    var v17 int8 = 127
    var v18 uint = 593
    v19 := GenS2{F0: 0, F1: 0}
    if (0 > uint8(v3)) {
        var v20 int32 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v21 uint64 = 4294967295
        _ = v20
        _ = v21
    } else {
        if (v7 >= ((0.0 / float32((1.5 * (0.1 / 1.0)))) * 0.1)) {
            switch v5 {
            case "":
                if len(v4) > 0 {
                    v4[0] = 128
                }
            case "\n\t":
                for fi22 := 0; fi22 < 1; fi22++ {
                    v17 = (v17 & v17)
                    _ = fi22
                }
            }
            var v23 uint64 = 784
            _ = v23
        } else {
            var v24 float32 = -2.25
            var v25 float32 = -2.25
            v26 := GenS3{F0: (v17 % (((v17 | v17)) | 1)), F1: -457, F2: 667}
            _ = v24
            _ = v25
            _ = v26
        }
        v27 := GenS1{F0: v5, F1: (!true), F2: ""}
        _ = v27
    }
    var v28 int8 = 42
    fmt.Println(v3, v5)
    fmt.Println(v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
