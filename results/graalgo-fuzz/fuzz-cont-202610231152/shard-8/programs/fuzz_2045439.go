package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0(p0 int64, p1 uint) int {
    v5 := GenS1{F0: 1, F1: 927}
    p0 = (^(p0))
    p0 = p0
    if (18446744073709551615 <= 18446744073709551615) {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int64 = p0
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 0, F1: 0}
        _ = v8
    }
    p1 -= 2
    _ = p0
    _ = p1
    _ = v5
    return 4294967295
}
func genFn1() int32 {
    if true {
        for fi9 := 0; fi9 < 2; fi9++ {
            if (-2147483648 < 1) {
                var v10 uint32 = 2147483647
                _ = v10
            } else {
                var v11 float32 = 0.0
                v11 = v11
                _ = v11
            }
            v12 := GenS1{F0: 128, F1: 65535}
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint8 = 203
                var v15 int = (fi9 / ((fi9) | 1))
                if true {
                    if false {
                        v14 -= v14
                    }
                    v15 = 9223372036854775807
                    var v16 uint16 = 889
                    _ = v16
                } else {
                    v14 = (v14 | (v14 ^ v14))
                    v17 := GenS1{F0: (v14 | (v14 ^ (v14 - v14))), F1: uint32(v14)}
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi9
            _ = v12
        }
        v18 := []int{4294967295}
        var v19 int64 = -1
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: false}
        _ = v20
    }
    var v21 int32 = 42
    v22 := []int{-2147483648, 1, 247}
    v21 = (v21 ^ (v21 % ((v21) | 1)))
    v23 := GenS3{F0: v21, F1: 234}
    _ = v21
    _ = v22
    _ = v23
    return (v21 + (-(v21)))
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"hello"}
    v3 := make(map[string]bool)
    var v4 float32 = 0.1
    var v24 uint16 = 32767
    for fi25 := 0; fi25 < 17; fi25++ {
        _ = len(v1)
        fmt.Println(v4, v4, v4)
        _ = fi25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        var v27 float32 = -2.25
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = fi26
        _ = v27
    }
    var v28 int32 = 122
    if false {
        fmt.Println(v24, v4)
        var v29 uint64 = 348
        _ = v29
    }
    if (128 <= 2) {
        if (!false) {
            v30 := GenS4{F0: 65535, F1: v24, F2: 42}
            v4 = v4
            var v31 int = 32767
            _ = v30
            _ = v31
        }
    } else {
        v24 -= (v24 % (((v24 % (((v24 ^ v24)) | 1))) | 1))
        var v32 uint = 4294967295
        _ = v32
    }
    if (376 >= 999) {
        var v33 int64 = 9223372036854775807
        _ = v33
    } else {
        switch -462 {
        case 128:
            var v34 int8 = 52
            _ = v34
        case -1:
            v4 -= (float32(float64(-2.25)) + (v4 + (float32(0.1) - float32(1.5))))
        }
        v24 -= 42
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 int8 = -100
            var v38 uint16 = (^(v24))
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
