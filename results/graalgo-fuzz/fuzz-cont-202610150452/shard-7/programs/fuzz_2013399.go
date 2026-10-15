package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0(p0 int32) bool {
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 float64 = -2.25
        var v7 int16 = 42
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{}
    p0 += (p0 / (((p0 * p0)) | 1))
    _ = p0
    _ = v8
    return false
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 1
    var v3 int = 127
    var v4 uint8 = 2
    v3 -= (v3 | (v3 | 4294967295))
    switch (v3 ^ (v3 - v3)) {
    case 1:
        v9 := GenS1{F0: (3.14159 > 1e10)}
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 3; fi10++ {
            switch (v3 * v3) {
            case 32768:
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 float32 = 1.5
                    v3 += int(v2)
                    delete(v1, (v2 & v2))
                    var v13 bool = (!true)
                    if (false && (float32(0.0) != v12)) {
                        var v14 string = "\n\t"
                        v2 += (v2 / ((v2) | 1))
                        v15 := []int{}
                        _ = v14
                        _ = v15
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case -9223372036854775808:
                v3 = (^(fi10))
            default:
                v1[v2] = 733
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                _ = len(v1)
                _ = fi16
            }
            _ = fi10
        }
    default:
        v4 -= (v4 & uint8(v2))
    }
    v17 := GenS1{F0: (-1 >= 66)}
    v18 := []int{}
    v3 = v3
    v3 = 1
    var v19 uint = 2
    v19 -= 42
    v2 -= 4294967295
    for fi20 := 0; fi20 < 17; fi20++ {
        fmt.Println(v19, fi20, v4)
        _ = fi20
    }
    _ = len(v18)
    v19 -= (v19 / ((65535) | 1))
    if ((v3 != (v3 ^ v3)) && (!(!false))) {
        v4 += (^(v4))
        delete(v1, (v2 | 1))
        var v21 int64 = 9223372036854775807
        _ = v21
    } else {
        v22 := GenS1{F0: (-1 <= -128)}
        var v23 bool = true
        if v23 {
            v24 := GenS1{F0: (1 <= 203)}
            _ = v24
        }
        _ = v22
        _ = v23
    }
    v4 = (v4 / ((uint8(v19)) | 1))
    v25 := GenS2{F0: -1, F1: (true || (int8(v3) < 0)), F2: v19}
    v2 = v2
    var v26 float32 = 1.5
    v27 := GenS3{F0: 2, F1: -68}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
