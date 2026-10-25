package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint16, p1 int, p2 int32) uint {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: (p0 % ((32767) | 1)), F1: (p2 - (p2 | (p2 + (p2 + p2)))), F2: -822}
            p1 = fi6
            if false {
                if (160 < 127) {
                    switch "a" {
                    case "\n\t":
                        v9 := GenS1{F0: 2, F1: p2, F2: 331}
                        _ = v9
                    case "foo bar":
                        p2 = 889
                    default:
                        var v10 uint32 = 4294967295
                        _ = v10
                    }
                    fmt.Println(fi6, p0)
                    var v11 float64 = 0.0
                    _ = v11
                } else {
                    p1 += (-(fi6))
                    var v12 bool = false
                    v13 := []int{(fi7 << ((1) & 63))}
                    _ = v12
                    _ = v13
                }
            } else {
                v14 := []int{}
                switch "" {
                case "hello":
                    var v15 int8 = 42
                    _ = v15
                case "foo bar":
                    p1 = 32768
                }
                if ((!true) || (p0 == p0)) {
                    p0 -= 0
                }
                _ = v14
            }
            _ = fi7
            _ = v8
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if (-1 <= -967) {
                var v17 uint64 = 678
                p2 = 0
                _ = v17
            }
            p2 += p2
            var v18 int8 = 1
            _ = fi16
            _ = v18
        }
        var v19 string = "\n\t"
        _ = fi6
        _ = v19
    }
    var v20 uint16 = p0
    var v21 uint = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 768
}
func genFn1(p0 string, p1 uint8, p2 uint8) uint32 {
    v22 := GenS2{F0: -128, F1: 0, F2: 65535}
    switch 128 {
    case 42:
        v23 := GenS2{F0: 90, F1: 2147483647, F2: 4294967295}
        _ = v23
    case -9223372036854775808:
        p0 = "a"
    default:
        var v24 bool = true
        _ = v24
    }
    fmt.Println(p2)
    if false {
        v25 := []int{4294967295, -403, -1}
        var v26 int32 = 42
        var v27 int8 = 94
        _ = v25
        _ = v26
        _ = v27
    }
    var v28 int = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v28
    return 745
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    var v5 uint = 195
    v1 += (v1 + v1)
    v1 -= (v1 + (v1 + uint8(v5)))
    var v29 float32 = -2.25
    v1 = (v1 % (((v1 ^ v1)) | 1))
    var v30 float32 = float32(0.0)
    if true {
        v31 := GenS3{F0: -277, F1: uint32(v1)}
        _ = v31
    }
    v5 -= v5
    v5 += 68
    for fi32 := 0; fi32 < 14; fi32++ {
        var v33 uint8 = (^(v1))
        var v34 int32 = 0
        var v35 int8 = -63
        _ = fi32
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := []int{-9223372036854775808, 42, 2147483647, -1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
}
