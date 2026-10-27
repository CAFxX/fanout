package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]int)
    var v3 uint8 = 51
    var v4 int16 = -13
    var v5 int32 = 42
    var v6 float64 = ((((((1.5 * (1.5 + float64(1e10))) * 1e10) * (-2.25 / 0.1)) - 1e10) + ((-2.25 - -2.25) * 0.1)) / -2.25)
    v7 := GenS1{F0: -776, F1: 4294967295}
    v8 := GenS2{F0: 42, F1: -1}
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float64 = (-2.25 - v6)
        switch 0 {
        case 9223372036854775807:
            v6 = v10
        default:
            v11 := GenS1{F0: -999, F1: 999}
            _ = v11
        }
        var v12 uint32 = 694
        _ = fi9
        _ = v10
        _ = v12
    }
    if true {
        var v13 int32 = v5
        if (584 <= 818) {
            v3 = v3
            if false {
                v1[127] = (true || ((v4 < v4) || true))
            }
            v5 += (v13 * (v13 << ((787) & 63)))
        } else {
            v14 := []int{}
            _ = v14
        }
        _ = v13
    } else {
        var v15 bool = (float32(v6) > (1.5 / 1.5))
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int64 = -9223372036854775808
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 int64 = v17
            _ = fi18
            _ = v19
        }
        v20 := GenS1{F0: v17, F1: 607}
        _ = fi16
        _ = v17
        _ = v20
    }
    var v21 uint8 = v3
    v3 += (v21 / ((v21) | 1))
    fmt.Println(v21, v3, v4)
    for fi22 := 0; fi22 < 7; fi22++ {
        if (!false) {
            v4 += (v4 + (v4 / ((-1) | 1)))
        }
        for fi23 := 0; fi23 < 11; fi23++ {
            if true {
                var v24 uint8 = (v3 & (v21 % ((v3) | 1)))
                v25 := []int{fi23, fi23, 32768}
                if false {
                    v26 := GenS3{F0: -1, F1: -194}
                    _ = v26
                }
                v6 += float64(0.1)
                _ = v24
                _ = v25
            }
            _ = fi23
        }
        v27 := GenS4{F0: 350}
        _ = fi22
        _ = v27
    }
    v28 := GenS1{F0: 9223372036854775807, F1: 365}
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v3, v6)
        _ = fi29
    }
    v30 := GenS1{F0: -1, F1: 184}
    fmt.Println(v5)
    v31 := GenS3{F0: v4, F1: (v5 % ((v5) | 1))}
    v32 := GenS2{F0: (v5 << ((0) & 63)), F1: 4294967295}
    switch 32768 {
    case -493:
        var v33 int8 = -49
        _ = v33
    case 127:
        var v34 bool = (65535 == 42)
        _ = v34
    }
    var v35 int8 = 1
    if true {
        v36 := GenS5{F0: (-(v35)), F1: 42}
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
