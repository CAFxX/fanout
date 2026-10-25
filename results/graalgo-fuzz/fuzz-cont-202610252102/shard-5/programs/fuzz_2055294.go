package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 uint8 }
func genFn0(p0 int8, p1 uint16) int8 {
    var v5 string = "hello"
    var v6 uint64 = 0
    fmt.Println(v5, p0)
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 -= p0
        var v8 int16 = 32767
        v9 := []int{(fi7 & fi7), (fi7 % (((fi7 % ((fi7) | 1))) | 1)), 42, 955}
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        fmt.Println(v5, v6)
        fmt.Println(p0, v6)
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-(p0))
}
func genFn1(p0 uint8) int64 {
    p0 -= p0
    v10 := GenS1{F0: -1, F1: 4294967295, F2: 1}
    p0 -= p0
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v1 = 42
    if true {
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := []int{fi11, fi11, fi11, (fi11 - (fi11 * fi11))}
            for fi13 := 0; fi13 < 6; fi13++ {
                if (933 == 109) {
                    v14 := GenS1{F0: 42, F1: 1, F2: (-(fi13))}
                    v15 := GenS2{F0: 65535, F1: 158}
                    v1 += v1
                    _ = v14
                    _ = v15
                }
                v16 := GenS3{F0: 2, F1: 42, F2: -626}
                fmt.Println(v1, fi13, fi13)
                v1 += (^(v1))
                _ = fi13
                _ = v16
            }
            _ = fi11
            _ = v12
        }
        var v17 uint64 = 42
        _ = v17
    } else {
        v1 += v1
        var v18 uint16 = 32767
        for fi19 := 0; fi19 < 8; fi19++ {
            v1 -= v1
            var v20 float64 = -2.25
            v1 -= v1
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    fmt.Println(v1)
    var v21 uint8 = 128
    v22 := GenS4{F0: false, F1: true, F2: (^(v21))}
    for fi23 := 0; fi23 < 2; fi23++ {
        v24 := GenS1{F0: 1, F1: (v1 * 2147483648), F2: (fi23 - 128)}
        v1 = v1
        _ = fi23
        _ = v24
    }
    fmt.Println(v1, v21)
    fmt.Println(v21, v21)
    var v25 int = -903
    v1 = (v1 & (v1 + (v1 << ((65535) & 63))))
    var v26 uint8 = (v21 % (((v21 * v21)) | 1))
    var v27 uint64 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
