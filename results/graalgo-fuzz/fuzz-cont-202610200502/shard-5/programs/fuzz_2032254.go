package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 string; F2 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 bool; F2 uint }

func main() {
    var v1 int32 = 1
    var v2 int64 = 9223372036854775807
    v3 := []int{}
    if (42 <= 0) {
        _ = len(v3)
    }
    var v4 int8 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v5 int64 = int64(v4)
    var v6 float64 = 0.0
    var v7 string = "a"
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = (0 > 4294967295)
        if (((v6 / (v6 / v6)) / v6) == v6) {
            var v10 int8 = (v4 ^ v4)
            _ = v10
        }
        switch (fi8 << ((18446744073709551615) & 63)) {
        case 32768:
            switch (fi8 % ((fi8) | 1)) {
            case -9223372036854775808:
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 uint = 1
                    _ = len(v3)
                    _ = len(v3)
                    _ = fi11
                    _ = v12
                }
            case 1:
                if false {
                    fmt.Println(v5)
                }
            case 758:
                if (0 != 1) {
                    var v13 int32 = (v1 - (v1 | v1))
                    v13 -= v1
                    _ = v13
                }
            default:
                _ = len(v3)
            }
        case 4294967295:
            v14 := GenS1{F0: 0}
            _ = v14
        case -9223372036854775808:
            fmt.Println(fi8, v4, v9)
        default:
            _ = len(v3)
        }
        _ = fi8
        _ = v9
    }
    _ = len(v3)
    v15 := GenS1{F0: (-(v5))}
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 float32 = 1.5
        for fi18 := 0; fi18 < 13; fi18++ {
            v7 = "hello"
            _ = len(v3)
            v3 = append(v3, fi16)
            _ = fi18
        }
        v5 -= 821
        _ = fi16
        _ = v17
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v19 float32 = (-2.25 / 1e10)
    var v20 float64 = v6
    if false {
        if false {
            var v21 string = (v7 + v7)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v21
        } else {
            fmt.Println(v7, v19)
            v22 := GenS2{F0: (^(v1)), F1: "\n\t", F2: (v4 / (((^(v4))) | 1))}
            _ = v22
        }
        fmt.Println(v7)
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        v3 = append(v3, (fi23 * 4294967295))
        switch -1 {
        case 92:
            v24 := GenS3{F0: 2}
            _ = v24
        default:
            v25 := GenS4{F0: 4294967295}
            _ = v25
        }
        v26 := GenS5{F0: -32768, F1: ((1 != 0) || true), F2: 847}
        _ = fi23
        _ = v26
    }
    var v27 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
