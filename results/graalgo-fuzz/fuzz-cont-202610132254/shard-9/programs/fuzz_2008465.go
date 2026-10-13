package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
func genFn0(p0 int64, p1 uint16) int {
    v5 := []int{}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        if len(v5) > 0 {
            v5[0] = fi6
        }
        fmt.Println(fi6, p1)
        _ = fi6
    }
    p1 = (p1 % ((p1) | 1))
    if (false || true) {
        var v7 uint32 = 0
        if true {
            fmt.Println(v7, v7)
            v7 += v7
        } else {
            var v8 int16 = 0
            _ = v8
        }
        p0 = (p0 << ((255) & 63))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 32768
}

func main() {
    var v1 string = ((("hello" + "\n\t") + "foo bar") + "foo bar")
    v2 := []uint32{84, 1, 1}
    var v3 float32 = 0.1
    var v4 int8 = 127
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 int64 = 0
        var v11 int64 = (v10 | 42)
        v12 := GenS1{F0: v1, F1: 18446744073709551615}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    var v13 int16 = -492
    if len(v2) > 0 {
        v2[0] = 2
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(fi14, v3)
        if false {
            if true {
                v4 = v4
                v15 := GenS2{F0: -2147483648, F1: -9223372036854775808}
                _ = v15
            }
            _ = len(v2)
            v16 := GenS3{F0: 556, F1: v13}
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS4{F0: 0, F1: (v4 << ((286) & 63))}
            _ = len(v2)
            _ = len(v2)
            _ = fi17
            _ = v18
        }
        _ = fi14
    }
    v4 = (v4 << ((255) & 63))
    for fi19 := 0; fi19 < 9; fi19++ {
        v4 += (v4 << ((18446744073709551615) & 63))
        _ = fi19
    }
    switch -526 {
    case 127:
        v20 := GenS5{F0: 9223372036854775807, F1: 331, F2: (v4 + (v4 & (v4 + v4)))}
        _ = v20
    }
    var v21 uint8 = 2
    v1 = (v1 + v1)
    v4 = v4
    var v22 float64 = (-2.25 - (float64(v3) + float64((3.14159 - v3))))
    switch v1 {
    case "a":
        switch v1 {
        case "":
            var v23 float32 = (float32(v22) - v3)
            _ = v23
        case "hello":
            _ = len(v2)
        case "a":
            v22 -= ((((1.5 - 3.14159) / (v22 / 1.5)) * v22) * 1.5)
        default:
            var v24 bool = (1 < 255)
            _ = v24
        }
    case "hello":
        var v25 bool = true
        _ = v25
    case "foo bar":
        fmt.Println(v3, v13)
    default:
        for fi26 := 0; fi26 < 19; fi26++ {
            var v27 string = v1
            if (!false) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v13 -= -1
            }
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS2{F0: 579, F1: 9223372036854775807}
    v13 += 0
    if (float64(v3) == v22) {
        v13 = v13
        v1 = v1
    }
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v3
    _ = v4
}
