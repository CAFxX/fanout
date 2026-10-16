package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint = 65535
        v8 := []int{(fi6 - (fi6 & 128))}
        v7 = (v7 | v7)
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int32 = -432
    var v10 bool = true
    v9 = v9
    _ = v10
    _ = v9
    return 42
}
func genFn1() int8 {
    v11 := []int{-379}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 float64 = ((-2.25 + 0.1) / 1.0)
    _ = v11
    _ = v12
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 42
    v3 := []int{128}
    v4 := make(map[uint32]bool)
    v5 := []string{}
    var v13 int8 = 1
    v13 = genFn1()
    v14 := GenS1{F0: 2, F1: 32767}
    v2 = (^(v2))
    if len(v3) > 0 {
        v3[0] = 591
    }
    if true {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS1{F0: 737, F1: 0}
            _ = fi15
            _ = v16
        }
    } else {
        switch "a" {
        case "\n\t":
            if false {
                if (!false) {
                    v2 = v2
                    v17 := GenS2{F0: -1, F1: 42, F2: v2}
                    _ = v17
                }
                v2 = v2
            }
        default:
            var v18 bool = true
            _ = v18
        }
        v19 := GenS1{F0: 1, F1: 32767}
        v13 -= v13
        _ = v19
    }
    v1[-674] = 42
    v13 = 77
    if true {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS3{F0: 42, F1: int16(v2)}
            if false {
                if (2147483647 < -82) {
                    var v22 int32 = 417
                    for fi23 := 0; fi23 < 17; fi23++ {
                        switch "foo bar" {
                        case "":
                            var v24 int16 = 42
                            _ = v24
                        }
                        _ = fi23
                    }
                    _ = v22
                }
                var v25 uint8 = 128
                v2 -= v2
                _ = v25
            } else {
                switch fi20 {
                case 32768:
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v27 := v4[4294967295]
                        _ = fi26
                        _ = v27
                    }
                case 328:
                    switch (fi20 ^ fi20) {
                    case -147:
                        var v28 uint64 = 18446744073709551615
                        _ = v28
                    }
                case 599:
                    if (!(1 <= 0)) {
                        var v29 int64 = 0
                        v2 -= 65535
                        for fi30 := 0; fi30 < 15; fi30++ {
                            v5 = append(v5, "\n\t")
                            var v31 uint16 = v2
                            var v32 float64 = 3.14159
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        _ = v29
                    }
                default:
                    var v33 bool = false
                    _ = v33
                }
                v2 = (v2 / (((v2 * v2)) | 1))
                var v34 uint = 18446744073709551615
                _ = v34
            }
            var v35 bool = false
            _ = fi20
            _ = v21
            _ = v35
        }
        v36 := GenS4{F0: 333, F1: 32767, F2: 254}
        _ = v36
    }
    v2 += (v2 - v2)
    fmt.Println(v2, v13)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
