package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
func genFn0(p0 int8) int64 {
    v6 := GenS1{F0: 128, F1: p0, F2: -1}
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS2{F0: 625}
        p0 = p0
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    var v9 uint16 = 1
    v9 -= (v9 & (v9 % (((v9 / (((v9 % ((v9) | 1))) | 1))) | 1)))
    switch 1 {
    case 105:
        v10 := GenS2{F0: 2}
        _ = v10
    case 0:
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 uint = 4294967295
            v9 += 2
            var v13 float64 = float64((-2.25 * 0.1))
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = v6
    _ = v9
    return -1
}
func genFn1(p0 uint16) string {
    p0 = p0
    if (239 != 42) {
        p0 += p0
        p0 = (p0 ^ p0)
    } else {
        p0 -= (p0 >> ((269) & 63))
        var v14 string = "hello"
        p0 -= 553
        _ = v14
    }
    p0 = (p0 / (((p0 - 42)) | 1))
    p0 = p0
    _ = p0
    return "hello"
}

func main() {
    var v1 string = "hello"
    var v2 uint8 = 127
    var v3 float64 = (-2.25 * (1e10 - (1.5 * 1e10)))
    v4 := make(map[uint32]bool)
    var v5 uint32 = 42
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 bool = (false && false)
        _ = fi15
        _ = v16
    }
    var v17 uint = 2
    switch "foo bar" {
    case "hello":
        _ = len(v4)
    default:
        switch -2147483648 {
        case -1:
            v5 -= (v5 << (((v17 | uint(v2))) & 63))
        case -2147483648:
            for fi18 := 0; fi18 < 3; fi18++ {
                v19 := []int{(fi18 + fi18), (fi18 * -297)}
                v1 = v1
                _ = fi18
                _ = v19
            }
        default:
            fmt.Println(v1, v2, v2)
        }
    }
    v3 += (float64(-2.25) - (float64(float32(-2.25)) / v3))
    fmt.Println(v5, v17, v5)
    v4[(v5 * 135)] = (1 >= uint64(v5))
    for fi20 := 0; fi20 < 12; fi20++ {
        var v21 int16 = -588
        _ = fi20
        _ = v21
    }
    var v22 uint8 = v2
    fmt.Println(v1)
    if true {
        v1 = v1
        v4[(v5 + v5)] = true
    }
    var v23 float64 = (((v3 * -2.25) / 1.0) * 0.0)
    var v24 int16 = -398
    v22 += 250
    if false {
        switch v1 {
        case "":
            v25 := GenS1{F0: (v22 >> ((v17) & 63)), F1: 117, F2: 1}
            _ = v25
        default:
            var v26 float32 = 1.5
            _ = v26
        }
        var v27 uint64 = 909
        _ = v27
    } else {
        v28 := v4[(^(v5))]
        switch ("a" + v1) {
        case "hello":
            v5 = v5
        default:
            var v29 uint32 = (v5 << (((^(v17))) & 63))
            _ = v29
        }
        _ = v28
    }
    if ((!false) || false) {
        switch 951 {
        case -202:
            if (143 >= 0) {
                v30 := GenS2{F0: uint32(v2)}
                v31 := GenS3{F0: v17, F1: (v24 | 1000), F2: 0}
                v22 = (^(v2))
                _ = v30
                _ = v31
            }
        case 32768:
            v32 := v4[uint32(v2)]
            _ = v32
        }
    } else {
        var v33 float32 = float32(v23)
        _ = v33
    }
    v2 -= (^(v2))
    var v34 int32 = 278
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
