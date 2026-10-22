package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint; F1 uint8 }
func genFn0(p0 uint32) int {
    var v7 int16 = -32768
    for fi8 := 0; fi8 < 18; fi8++ {
        p0 = p0
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(fi8, p0)
            var v10 int16 = (v7 + (v7 << ((255) & 63)))
            var v11 int64 = -9223372036854775808
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: -128, F1: false}
        _ = fi8
        _ = v12
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 2
    v2 := []int{357}
    v3 := make(map[string]bool)
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    var v6 uint8 = 128
    v1 = v1
    v1 = v1
    v13 := GenS2{F0: 32767, F1: genFn0(42), F2: v1}
    if ((!(793 == 0)) || true) {
        v14 := GenS2{F0: 32767, F1: genFn0((v1 / (((v1 - v1)) | 1))), F2: v1}
        v6 = 0
        fmt.Println(v6, v1)
        _ = v14
    }
    v1 = (v1 + (v1 * v1))
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            if (true || false) {
                if false {
                    _ = len(v2)
                }
            }
            _ = len(v3)
            _ = fi16
        }
        var v17 int8 = 50
        if true {
            var v18 int8 = v17
            var v19 float32 = 1.5
            if ((v18 <= v17) || (0 < 936)) {
                if (!(639 >= 2)) {
                    v20 := GenS3{F0: 92}
                    switch genFn0((v1 << ((18446744073709551615) & 63))) {
                    case -2147483648:
                        delete(v4, genFn0(v1))
                    default:
                        var v21 int64 = 0
                        _ = v21
                    }
                    v17 = v17
                    _ = v20
                } else {
                    v22 := GenS2{F0: 2, F1: (fi15 ^ 128), F2: v1}
                    v23 := GenS4{F0: 1, F1: 0}
                    _ = v22
                    _ = v23
                }
                v1 = v1
            } else {
                v24 := GenS5{F0: 5, F1: 67}
                _ = v24
            }
            _ = v18
            _ = v19
        }
        _ = fi15
        _ = v17
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            if len(v2) > 0 {
                v2[0] = (fi25 / ((fi26) | 1))
            }
            fmt.Println(fi25, v6, v6)
            _ = fi26
        }
        fmt.Println(v1)
        fmt.Println(v6, v6, v1)
        _ = fi25
    }
    fmt.Println(v6, v6)
    v27 := GenS2{F0: 994, F1: genFn0((v1 % ((v1) | 1))), F2: (v1 + v1)}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
