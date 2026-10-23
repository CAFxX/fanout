package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint }
type GenS6 struct { F0 int; F1 uint; F2 uint8 }
func genFn0() uint {
    switch ("foo bar" + "hello") {
    case "":
        switch 42 {
        case 32767:
            if ((false && true) || false) {
                for fi6 := 0; fi6 < 10; fi6++ {
                    var v7 uint64 = 2
                    v7 -= (^(v7))
                    _ = fi6
                    _ = v7
                }
            } else {
                for fi8 := 0; fi8 < 10; fi8++ {
                    v9 := GenS1{F0: 1, F1: 731}
                    _ = fi8
                    _ = v9
                }
                if false {
                    v10 := GenS1{F0: 255, F1: -9223372036854775808}
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 int = fi11
                        var v13 uint16 = 425
                        v14 := []int{fi11, fi11}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    var v15 float32 = 1e10
                    _ = v15
                }
            }
        case 661:
            var v16 uint16 = 2
            _ = v16
        case 128:
            v17 := []int{207}
            _ = v17
        default:
            v18 := GenS1{F0: 128, F1: -1}
            _ = v18
        }
    }
    v19 := []int{}
    v20 := GenS2{F0: -1, F1: 0, F2: 616}
    var v21 int16 = -267
    for fi22 := 0; fi22 < 8; fi22++ {
        if (((340 <= 812) && true) && (127 < 74)) {
            if (3.14159 == 0.1) {
                v19 = append(v19, 127)
                for fi23 := 0; fi23 < 15; fi23++ {
                    v19 = append(v19, 1)
                    v24 := GenS3{F0: int8(v21), F1: 862}
                    fmt.Println(fi22, fi23, fi22)
                    _ = fi23
                    _ = v24
                }
            }
            var v25 int64 = 141
            var v26 float64 = -2.25
            _ = v25
            _ = v26
        } else {
            v21 -= v21
        }
        _ = fi22
    }
    _ = v19
    _ = v20
    _ = v21
    return 255
}
func genFn1(p0 uint64, p1 int32, p2 uint16) bool {
    p0 = p0
    var v27 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return (true || true)
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{0, 1, 2, 4294967295}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]uint32)
    var v5 int = 42
    v5 -= v5
    v28 := GenS4{F0: 32767}
    for fi29 := 0; fi29 < 4; fi29++ {
        fmt.Println(fi29)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi29
    }
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 int8 = -1
        v5 = fi30
        _ = fi30
        _ = v31
    }
    fmt.Println(v5)
    var v32 int = 2147483647
    for fi33 := 0; fi33 < 13; fi33++ {
        fmt.Println(v32)
        var v34 int8 = 42
        _ = fi33
        _ = v34
    }
    v32 += (v5 & v5)
    v32 += 4294967295
    switch 1 {
    case 1:
        _ = len(v3)
    case 839:
        fmt.Println(v5)
    default:
        v35 := GenS5{F0: false, F1: 4294967295, F2: 4294967295}
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        v5 -= (v5 | (v32 / ((v32) | 1)))
        v37 := GenS6{F0: v32, F1: 2, F2: 255}
        _ = fi36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
