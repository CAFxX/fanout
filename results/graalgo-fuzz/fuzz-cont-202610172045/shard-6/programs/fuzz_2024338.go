package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint8, p1 int16, p2 uint) int32 {
    p1 = (p1 | -32768)
    p0 += (p0 % ((173) | 1))
    var v5 int32 = 160
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        p2 -= p2
        if false {
            if false {
                for fi9 := 0; fi9 < 20; fi9++ {
                    p0 = (p0 * 0)
                    switch "\n\t" {
                    case "hello":
                        switch "\n\t" {
                        case "foo bar":
                            v10 := GenS1{F0: v5, F1: 18446744073709551615, F2: -67}
                            _ = v10
                        default:
                            v11 := []int{}
                            _ = v11
                        }
                    }
                    _ = fi9
                }
                v5 = (v5 & (v5 / ((v5) | 1)))
            }
            if true {
                switch (fi8 >> (((^(p2))) & 63)) {
                case -9223372036854775808:
                    p0 -= p0
                }
                var v12 int8 = -128
                _ = v12
            } else {
                v13 := []int{-53}
                v14 := GenS1{F0: 446, F1: 2, F2: 4}
                var v15 int64 = 260
                _ = v13
                _ = v14
                _ = v15
            }
            v16 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: -17}
            _ = v16
        }
        fmt.Println(v5)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return v5
}
func genFn1() int32 {
    var v17 uint16 = 255
    switch -578 {
    case 0:
        switch 9223372036854775807 {
        case 127:
            for fi18 := 0; fi18 < 13; fi18++ {
                v17 = v17
                v17 -= v17
                _ = fi18
            }
        default:
            if (true && (1e10 < (float64(3.14159) * 3.14159))) {
                v19 := []int{9223372036854775807, 20, 280, 32768}
                var v20 string = ("foo bar" + ("hello" + "\n\t"))
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := []int{}
                    v17 += v17
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: -221, F1: 810, F2: -31}
                v24 := []int{32767, 132}
                _ = v23
                _ = v24
            }
        }
    case 1:
        v25 := []int{0}
        _ = v25
    }
    v26 := GenS1{F0: 0, F1: 4294967295, F2: 127}
    var v27 uint64 = 4294967295
    _ = v17
    _ = v26
    _ = v27
    return -2147483648
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 128
    var v3 int16 = 1
    var v4 int8 = 100
    v2 = v2
    v3 = (v3 ^ -32768)
    v28 := GenS2{F0: 604}
    fmt.Println(v3, v2)
    fmt.Println(v4, v2, v3)
    var v29 int16 = -32768
    v30 := GenS2{F0: 473}
    v31 := []int{7}
    v3 = 1
    var v32 int64 = -295
    for fi33 := 0; fi33 < 16; fi33++ {
        v4 = v4
        fmt.Println(v32, v29)
        switch -9223372036854775808 {
        case 693:
            v32 = 0
        case -770:
            for fi34 := 0; fi34 < 9; fi34++ {
                v35 := GenS3{F0: v4}
                v31 = append(v31, -898)
                var v36 int8 = (-(v4))
                _ = fi34
                _ = v35
                _ = v36
            }
        case 32767:
            v1[(fi33 - -2147483648)] = 431
        }
        _ = fi33
    }
    if (!false) {
        v37 := GenS3{F0: int8(v2)}
        var v38 int = 42
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 9; fi39++ {
            v29 = int16(v32)
            var v40 int = -9223372036854775808
            if (4294967295 != 2) {
                fmt.Println(v2, v2)
                var v41 int = v40
                _ = v41
            } else {
                var v42 string = "\n\t"
                v29 = -1
                v43 := GenS1{F0: -882, F1: 507, F2: -128}
                _ = v42
                _ = v43
            }
            _ = fi39
            _ = v40
        }
    }
    var v44 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v44
}
