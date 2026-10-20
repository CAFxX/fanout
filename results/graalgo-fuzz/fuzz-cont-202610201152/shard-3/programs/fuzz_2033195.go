package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int8 }
type GenS5 struct { F0 int64; F1 bool }

func main() {
    v1 := []int{}
    var v2 bool = (!false)
    var v3 uint16 = 42
    if (v2 && ((v3 << ((1) & 63)) < (v3 | 42))) {
        v4 := GenS1{F0: -892, F1: (v2 || v2)}
        _ = v4
    } else {
        v5 := GenS1{F0: -32768, F1: false}
        _ = v5
    }
    if ((float32((1e10 + (((-2.25 * (1.5 + (3.14159 + 1.5))) + -2.25) - 3.14159))) < (1e10 * (((((-2.25 * (0.0 * (3.14159 / 1.0))) - (float32((float64(0.0) / (1e10 * 0.1))) * -2.25)) - -2.25) / (1e10 / 1.5)) - -2.25))) && (v2 && (!(0.0 < 0.1)))) {
        var v6 int16 = 32767
        _ = v6
    }
    switch -9223372036854775808 {
    case 32767:
        for fi7 := 0; fi7 < 6; fi7++ {
            v1 = append(v1, fi7)
            if (v2 || (false || false)) {
                var v8 uint = 701
                _ = v8
            }
            _ = fi7
        }
    case -9223372036854775808:
        if (!(2 > 42)) {
            for fi9 := 0; fi9 < 16; fi9++ {
                v3 = v3
                _ = fi9
            }
            v10 := GenS2{F0: (v3 % (((^(v3))) | 1)), F1: 0}
            v11 := GenS3{F0: 33, F1: 165}
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 13; fi12++ {
            v2 = v2
            var v13 uint64 = 832
            _ = fi12
            _ = v13
        }
    }
    v3 += v3
    var v14 uint8 = 127
    var v15 float64 = (0.0 / float64((3.14159 - 0.1)))
    if v2 {
        fmt.Println(v3, v14)
        fmt.Println(v15, v14, v2)
    } else {
        switch 4294967295 {
        case -975:
            var v16 uint16 = 32767
            _ = v16
        case 32768:
            var v17 uint64 = 1
            _ = v17
        case 32767:
            v18 := GenS4{F0: 32767, F1: 1, F2: 42}
            _ = v18
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v19 := GenS5{F0: -1, F1: false}
        _ = v19
    }
    v15 += (v15 / 1.0)
    v15 -= 3.14159
    fmt.Println(v2)
    v1 = append(v1, 128)
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint = 42
        _ = fi20
        _ = v21
    }
    var v22 float64 = (0.1 * -2.25)
    if (!v2) {
        v2 = v2
        if v2 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        var v23 uint16 = (v3 & (^(v3)))
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
}
