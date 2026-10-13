package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint }
func genFn0(p0 int, p1 int8) uint64 {
    v6 := GenS1{F0: -1}
    var v7 int32 = 42
    v8 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 880
}
func genFn1(p0 uint32, p1 uint64) int8 {
    fmt.Println(p1, p1)
    v9 := GenS2{F0: -56, F1: int64(p1), F2: int16(p0)}
    _ = p0
    _ = p1
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 608
    var v2 string = "a"
    var v3 int64 = -1
    var v4 uint = 18446744073709551615
    var v5 uint16 = 42
    fmt.Println(v3, v4)
    for fi10 := 0; fi10 < 9; fi10++ {
        if (-145 == (fi10 - fi10)) {
            if false {
                switch -306 {
                case 1:
                    var v11 uint8 = uint8(v1)
                    _ = v11
                case 173:
                    v12 := []int{(fi10 ^ (fi10 ^ (fi10 * (fi10 >> ((uint(v1)) & 63))))), (fi10 << ((274) & 63)), 32767}
                    _ = v12
                case 127:
                    fmt.Println(v5, fi10)
                default:
                    var v13 float64 = 3.14159
                    _ = v13
                }
                v14 := []int{32767, fi10, (fi10 ^ fi10), (fi10 * fi10)}
                v5 = (v5 / (((v5 & v5)) | 1))
                _ = v14
            }
            for fi15 := 0; fi15 < 13; fi15++ {
                switch 127 {
                case 4294967295:
                    fmt.Println(v3)
                case 32767:
                    var v16 uint64 = (^(v1))
                    _ = v16
                case 81:
                    if (!((!(2 != (v4 - 1))) || (227 >= 0))) {
                        fmt.Println(fi10)
                        v4 = (^(v4))
                        v17 := []int{(^(fi15)), (fi15 % ((fi10) | 1)), 9223372036854775807}
                        _ = v17
                    }
                default:
                    var v18 string = ((v2 + (v2 + v2)) + "a")
                    _ = v18
                }
                _ = fi15
            }
            var v19 int64 = (v3 * (v3 / (((v3 & int64(v1))) | 1)))
            _ = v19
        }
        var v20 int8 = 0
        _ = fi10
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS2{F0: -128, F1: -847, F2: 258}
        var v23 bool = true
        v4 += (v4 << ((v4) & 63))
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint8 = 127
    v25 := GenS1{F0: genFn1(0, genFn0(128, 42))}
    v26 := GenS3{F0: -1, F1: (v24 + (v24 << (((v4 ^ (v4 >> ((v4) & 63)))) & 63))), F2: 4294967295}
    if false {
        v1 += uint64(v24)
    } else {
        v27 := GenS4{F0: false, F1: (v4 >> ((v4) & 63))}
        fmt.Println(v5, v5, v24)
        var v28 uint16 = 2
        _ = v27
        _ = v28
    }
    v4 -= (v4 % ((616) | 1))
    var v29 int = 128
    var v30 uint32 = 977
    v31 := []int{9223372036854775807, (v29 + v29), v29}
    var v32 int8 = genFn1((v30 ^ v30), genFn0((v29 >> (((^(v4))) & 63)), 68))
    var v33 int16 = 32767
    for fi34 := 0; fi34 < 19; fi34++ {
        fmt.Println(v1, v3, fi34)
        v29 = v29
        v2 = v2
        _ = fi34
    }
    var v35 float32 = (0.0 - -2.25)
    fmt.Println(v29, v4, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v4
    _ = v5
}
