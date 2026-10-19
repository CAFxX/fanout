package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64) uint8 {
    if true {
        p0 = 0
        fmt.Println(p0)
        if true {
            p0 = (p0 * 18446744073709551615)
        } else {
            v5 := []int{0, 127, -486}
            _ = v5
        }
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 uint64 = (^(p0))
        for fi8 := 0; fi8 < 18; fi8++ {
            switch fi8 {
            case -1:
                if true {
                    v7 = (^(p0))
                } else {
                    v9 := GenS1{F0: (fi8 >> ((0) & 63)), F1: 10}
                    _ = v9
                }
            case -122:
                v10 := []int{}
                _ = v10
            case 2147483647:
                p0 = (^(v7))
            }
            var v11 bool = false
            var v12 float32 = 3.14159
            v13 := GenS1{F0: 128, F1: 774}
            _ = fi8
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    v14 := []int{642, 42, 0}
    v14 = append(v14, -9223372036854775808)
    p0 = (p0 | (p0 % ((432) | 1)))
    _ = p0
    _ = v14
    return 128
}
func genFn1(p0 bool, p1 bool) int16 {
    if (p1 && (0 > 0)) {
        p1 = (1 > -89)
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                var v17 int32 = 0
                fmt.Println(v17)
                _ = fi16
                _ = v17
            }
            fmt.Println(p1, p1, fi15)
            _ = fi15
        }
        if (p1 || (32767 < 65535)) {
            if (79 < 1) {
                v18 := GenS2{F0: 1, F1: 960}
                var v19 float64 = 1.5
                fmt.Println(v19)
                _ = v18
                _ = v19
            }
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 uint = 42
                fmt.Println(p1, fi20, p1)
                _ = fi20
                _ = v21
            }
        } else {
            fmt.Println(p1, p0)
            v22 := []int{}
            var v23 float64 = -2.25
            _ = v22
            _ = v23
        }
    }
    v24 := GenS3{F0: 1, F1: 88, F2: 100}
    _ = p0
    _ = p1
    _ = v24
    return 1000
}

func main() {
    var v1 uint16 = 0
    v2 := []string{"\n\t", "hello", "a"}
    v3 := make(map[string]bool)
    var v4 uint8 = 5
    v3["\n\t"] = (463 > 18446744073709551615)
    v25 := GenS4{F0: 255, F1: 1}
    for fi26 := 0; fi26 < 2; fi26++ {
        v27 := GenS3{F0: 4294967295, F1: 2147483648, F2: 100}
        fmt.Println(fi26)
        _ = fi26
        _ = v27
    }
    var v28 uint8 = v4
    switch "" {
    case "":
        v4 = (v4 / ((genFn0(440)) | 1))
    case "foo bar":
        v29 := GenS2{F0: 65, F1: -9223372036854775808}
        _ = v29
    case "\n\t":
        _ = len(v3)
    default:
        v4 += v28
    }
    v4 += 127
    var v30 float32 = (0.1 * ((3.14159 * 3.14159) + (0.0 * 0.0)))
    var v31 int64 = -1
    v28 -= (v28 % ((1) | 1))
    fmt.Println(v4, v31, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v32 := GenS2{F0: 2, F1: 1}
    v33 := GenS3{F0: 4294967295, F1: 539, F2: -70}
    v1 += (v1 - (v1 * (v1 * 42)))
    v28 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
