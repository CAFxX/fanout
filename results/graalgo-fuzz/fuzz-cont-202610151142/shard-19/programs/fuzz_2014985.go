package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 uint16, p1 int64) bool {
    p1 = (p1 ^ int64(p0))
    p1 = p1
    _ = p0
    _ = p1
    return ((!false) || true)
}

func main() {
    v1 := []int{-730}
    var v2 uint8 = 137
    var v3 uint64 = 2
    fmt.Println(v3, v3)
    v1 = append(v1, 127)
    v4 := GenS1{F0: v2, F1: -9223372036854775808}
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 bool = (3.14159 != -2.25)
            _ = fi6
            _ = v7
        }
        switch fi5 {
        case 9223372036854775807:
            v8 := GenS2{F0: int32(fi5), F1: false}
            _ = v8
        case -871:
            var v9 float64 = 0.0
            _ = v9
        case 215:
            for fi10 := 0; fi10 < 3; fi10++ {
                v2 = (^(v2))
                v11 := GenS1{F0: v2, F1: -9223372036854775808}
                if true {
                    if ((v3 - 42) > v3) {
                        var v12 uint16 = 633
                        var v13 uint32 = 4294967295
                        v13 = (v13 << ((255) & 63))
                        v14 := GenS1{F0: v2, F1: 9223372036854775807}
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v15 := GenS2{F0: 42, F1: (-861 <= -638)}
                    _ = v15
                } else {
                    v3 -= v3
                    fmt.Println(v3)
                    var v16 float64 = 0.1
                    _ = v16
                }
                _ = fi10
                _ = v11
            }
        }
        var v17 uint = 0
        _ = fi5
        _ = v17
    }
    v3 = (v3 + (v3 - v3))
    var v18 int = 1
    switch v18 {
    case -593:
        var v19 float32 = 1.5
        _ = v19
    case -36:
        v20 := GenS3{F0: 121, F1: 367, F2: ""}
        _ = v20
    case 9223372036854775807:
        var v21 uint8 = (v2 - 127)
        _ = v21
    default:
        if false {
            if false {
                for fi22 := 0; fi22 < 3; fi22++ {
                    var v23 int16 = 838
                    _ = fi22
                    _ = v23
                }
                v18 += (^(v18))
            }
        } else {
            v18 = (v18 >> ((uint(v2)) & 63))
            var v24 bool = false
            if len(v1) > 0 {
                v1[0] = v18
            }
            _ = v24
        }
    }
    var v25 bool = (false && (-1 >= 42))
    v1 = append(v1, (^(v18)))
    var v26 float32 = 1e10
    v26 = v26
    var v27 uint8 = v2
    var v28 int8 = 72
    var v29 uint = 18446744073709551615
    var v30 int64 = -9223372036854775808
    v18 = (v18 ^ (v18 % (((v18 * (v18 << ((v29) & 63)))) | 1)))
    v31 := GenS1{F0: (v27 >> (((v29 * 255)) & 63)), F1: int64(v18)}
    if true {
        var v32 int8 = v28
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
