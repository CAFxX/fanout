package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0(p0 bool, p1 uint64) int {
    v5 := GenS1{F0: -538, F1: -1, F2: 42}
    var v6 bool = (1.5 < -2.25)
    var v7 uint64 = 774
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 uint64, p1 uint16, p2 bool) int32 {
    switch 127 {
    case -608:
        switch 42 {
        case 32767:
            for fi8 := 0; fi8 < 16; fi8++ {
                v9 := GenS2{F0: -41}
                _ = fi8
                _ = v9
            }
        }
    case 9223372036854775807:
        v10 := GenS3{F0: 0, F1: 65535, F2: (p1 % ((1) | 1))}
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 bool = (((fi11 < (fi11 >> ((1) & 63))) || p2) && true)
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: -1, F1: 65535, F2: (p1 >> ((2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 183
}

func main() {
    var v1 int64 = 0
    v2 := []uint32{}
    var v3 uint16 = 255
    v4 := make(map[uint32]int)
    var v14 int32 = 42
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 = genFn1(42, 2, false)
    if true {
        v3 -= 42
    }
    if (32767 == 1) {
        v15 := GenS1{F0: genFn1(4294967295, v3, false), F1: 4294967295, F2: 105}
        _ = v15
    }
    if ((0.1 / 0.1) == 0.0) {
        v1 += v1
        switch -153 {
        case -36:
            fmt.Println(v1, v14, v3)
        case -596:
            v4[42] = 0
        case -547:
            v3 = v3
        default:
            fmt.Println(v14, v1)
        }
    }
    switch -192 {
    case -972:
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS4{F0: 42, F1: ((v1 << ((2) & 63)) > v1), F2: 42}
            _ = fi16
            _ = v17
        }
    default:
        v18 := GenS2{F0: -293}
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        switch fi19 {
        case 9223372036854775807:
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS3{F0: (-(fi20)), F1: 4294967295, F2: (v3 ^ 544)}
                _ = fi20
                _ = v21
            }
        case 200:
            fmt.Println(v14, v14)
        default:
            var v22 int8 = 127
            _ = v22
        }
        var v23 float64 = 0.0
        _ = fi19
        _ = v23
    }
    v14 -= (v14 >> ((255) & 63))
    v24 := GenS4{F0: 2, F1: true, F2: 128}
    switch "a" {
    case "foo bar":
        switch "foo bar" {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case "":
            v25 := v4[937]
            _ = v25
        }
    default:
        var v26 int32 = -313
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        fmt.Println(fi27, v1)
        var v28 uint64 = 4294967295
        _ = fi27
        _ = v28
    }
    v1 -= -1
    if (v14 > (-(v14))) {
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 int8 = -128
            _ = fi29
            _ = v30
        }
        v1 -= (^(v1))
    } else {
        for fi31 := 0; fi31 < 2; fi31++ {
            v32 := GenS4{F0: 42, F1: (float32(((1.5 + ((((-2.25 * 1e10) + (0.0 / 1.0)) - (0.0 / -2.25)) + (3.14159 - (-2.25 * 3.14159)))) + 0.0)) == -2.25), F2: 119}
            var v33 int32 = genFn1(1, 927, (!false))
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
