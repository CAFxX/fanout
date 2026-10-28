package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32, p1 uint8) bool {
    if (true && true) {
        fmt.Println(p1, p0)
    } else {
        var v6 int32 = -2147483648
        _ = v6
    }
    v7 := GenS1{F0: uint32(p1)}
    var v8 uint16 = 1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (!true)
}
func genFn1(p0 uint64, p1 int8) int16 {
    p1 += p1
    v9 := GenS1{F0: 65535}
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(p1, fi10, p0)
        v11 := GenS2{F0: -1, F1: p1, F2: 128}
        v12 := GenS1{F0: 618}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    v3 := []int{32768, 1}
    v4 := []uint32{}
    var v5 int8 = -128
    v13 := GenS1{F0: 98}
    var v14 int = 111
    if (((42 == 42) || false) && false) {
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint16 = 65535
            v14 = (fi15 + (fi15 ^ -9223372036854775808))
            _ = fi15
            _ = v16
        }
        v5 -= v5
    }
    v5 -= v5
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint16 = 0
        v14 += fi17
        _ = fi17
        _ = v18
    }
    fmt.Println(v14)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2[65535] = 2147483648
    var v19 uint64 = 383
    var v20 int64 = -748
    _ = len(v1)
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 int = v14
        _ = fi21
        _ = v22
    }
    switch v14 {
    case 2147483647:
        v5 -= (^(v5))
    case 42:
        var v23 uint64 = (v19 & v19)
        _ = v23
    default:
        if (255 < 255) {
            v14 = (^(v14))
        } else {
            for fi24 := 0; fi24 < 8; fi24++ {
                v25 := GenS3{F0: 289}
                _ = fi24
                _ = v25
            }
            v20 -= v20
            var v26 int8 = (v5 >> ((18446744073709551615) & 63))
            _ = v26
        }
    }
    v27 := GenS2{F0: 1, F1: (v5 ^ v5), F2: 255}
    for fi28 := 0; fi28 < 6; fi28++ {
        if true {
            for fi29 := 0; fi29 < 17; fi29++ {
                if len(v4) > 0 {
                    v4[0] = 2
                }
                v19 = v19
                v30 := GenS4{F0: 362}
                _ = fi29
                _ = v30
            }
        } else {
            var v31 string = "a"
            for fi32 := 0; fi32 < 13; fi32++ {
                v20 -= 0
                if (0 <= 42) {
                    var v33 float32 = ((0.1 + 0.1) / (0.1 * (1e10 + 1.5)))
                    var v34 float64 = 3.14159
                    _ = v33
                    _ = v34
                }
                _ = len(v2)
                var v35 uint16 = uint16(v19)
                _ = fi32
                _ = v35
            }
            v36 := GenS5{F0: 65535, F1: 1}
            _ = v31
            _ = v36
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            v4 = append(v4, 2147483647)
            switch v14 {
            case 2147483647:
                var v38 uint8 = 255
                _ = v38
            case -413:
                var v39 uint16 = 2
                _ = v39
            default:
                var v40 int32 = 645
                _ = v40
            }
            v14 += (fi37 ^ (^(v14)))
            _ = fi37
        }
        v19 = 42
        _ = fi28
    }
    var v41 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
