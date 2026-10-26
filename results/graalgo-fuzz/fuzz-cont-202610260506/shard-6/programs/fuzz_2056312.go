package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
func genFn0(p0 bool, p1 int16, p2 string) bool {
    var v5 int32 = -539
    v6 := GenS1{F0: 550, F1: (-(v5))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 string, p1 uint8, p2 int8) uint16 {
    switch 128 {
    case 0:
        p1 = 1
    case -2147483648:
        var v7 uint16 = 1
        _ = v7
    default:
        var v8 uint = 2
        _ = v8
    }
    switch 0 {
    case -591:
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS2{F0: -9223372036854775808}
            if (false || true) {
                if (888 < 42) {
                    for fi11 := 0; fi11 < 9; fi11++ {
                        v12 := GenS2{F0: 9223372036854775807}
                        v13 := []int{-631}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    if true {
                        v14 := GenS3{F0: 362, F1: 4294967295, F2: -2147483648}
                        _ = v14
                    }
                    p0 = p0
                } else {
                    v15 := []int{-47}
                    var v16 int8 = (p2 ^ p2)
                    var v17 int8 = (v16 << ((977) & 63))
                    _ = v15
                    _ = v16
                    _ = v17
                }
                v18 := GenS4{F0: 1}
                _ = v18
            } else {
                p0 = p0
                p0 = "hello"
                fmt.Println(p0, fi9)
            }
            _ = fi9
            _ = v10
        }
    case 683:
        p2 -= (^(p2))
    default:
        if true {
            var v19 uint16 = 542
            v19 -= 65535
            p2 = (-(p2))
            _ = v19
        }
    }
    fmt.Println(p2)
    var v20 int8 = (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 783
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    var v4 int8 = -128
    var v21 uint = 2
    var v22 int = -9223372036854775808
    var v23 uint32 = 377
    var v24 uint32 = v23
    var v25 uint16 = 42
    v21 = (v21 ^ (v21 | v21))
    var v26 float64 = 0.0
    v25 = (v25 * v25)
    if (-304 <= 32767) {
        v4 -= (v4 * (v4 >> ((42) & 63)))
        v27 := []int{v22, v22}
        _ = v27
    } else {
        v23 += (v24 << ((310) & 63))
        v28 := []int{}
        _ = v28
    }
    var v29 int16 = 1
    if (!false) {
        for fi30 := 0; fi30 < 14; fi30++ {
            v29 += 1
            for fi31 := 0; fi31 < 1; fi31++ {
                if true {
                    v32 := GenS4{F0: (v21 * (v21 % (((v21 * v21)) | 1)))}
                    if (!(!(0 <= 42))) {
                        fmt.Println(v25, v4)
                    }
                    var v33 int32 = int32(v22)
                    _ = v32
                    _ = v33
                }
                _ = len(v2)
                var v34 float64 = (v26 * (v26 * (v26 * float64(1e10))))
                var v35 uint16 = v25
                v21 -= 875
                _ = fi31
                _ = v34
                _ = v35
            }
            _ = fi30
        }
        v36 := GenS5{F0: 315}
        _ = v36
    } else {
        delete(v3, "hello")
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
