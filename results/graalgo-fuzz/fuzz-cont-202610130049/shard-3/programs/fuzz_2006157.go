package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 521
    var v3 float32 = 0.0
    var v4 uint = 454
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 bool = (!(((("a" + "a") + ("" + "\n\t")) + "") < ("foo bar" + "\n\t")))
            switch fi5 {
            case 221:
                var v8 int16 = -1
                _ = v8
            default:
                v9 := GenS1{F0: (fi6 + (-(fi5))), F1: 0, F2: v2}
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if ((1 > 42) && (255 != 255)) {
        if false {
            var v10 float32 = 0.1
            _ = v10
        } else {
            v3 -= (v3 / (-2.25 * v3))
            v11 := GenS1{F0: 1, F1: 42, F2: v2}
            _ = v11
        }
        v2 -= (v2 / ((18446744073709551615) | 1))
        v2 = (v2 + v2)
    } else {
        if false {
            fmt.Println(v3, v2, v4)
        } else {
            v12 := GenS2{F0: -128, F1: -43, F2: 42}
            _ = v12
        }
    }
    v4 += (^(v4))
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := v1[uint32(v4)]
        switch "foo bar" {
        case "foo bar":
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(fi13, v4, v3)
                _ = fi15
            }
        case "\n\t":
            v16 := GenS1{F0: (fi13 % ((fi13) | 1)), F1: 255, F2: 18446744073709551615}
            _ = v16
        default:
            var v17 int = fi13
            _ = v17
        }
        var v18 int = 2147483647
        _ = fi13
        _ = v14
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 int64 = -9223372036854775808
        _ = fi19
        _ = v20
    }
    if ((1.5 / 1e10) != (-2.25 * 3.14159)) {
        v21 := v1[2]
        delete(v1, (v21 << (((v4 * v4)) & 63)))
        _ = v21
    }
    v3 += (v3 - 3.14159)
    var v22 int8 = -128
    if true {
        fmt.Println(v2)
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        var v24 uint64 = (^(v2))
        _ = fi23
        _ = v24
    }
    switch "foo bar" {
    case "a":
        if (!false) {
            for fi25 := 0; fi25 < 4; fi25++ {
                v3 += (v3 / (1e10 * (v3 / v3)))
                if ((float64(1e10) * ((-2.25 - 0.0) / ((0.0 - (0.0 * (-2.25 - float64(0.0)))) * ((0.0 - float64(v3)) / 1.0)))) <= 0.0) {
                    var v26 uint16 = 255
                    switch "a" {
                    case "foo bar":
                        var v27 int32 = int32(v4)
                        _ = v27
                    case "\n\t":
                        v28 := []int{fi25, (fi25 & (fi25 * (fi25 << ((v4) & 63)))), fi25, fi25}
                        _ = v28
                    default:
                        v29 := GenS3{F0: 893}
                        _ = v29
                    }
                    _ = v26
                } else {
                    v3 = float32(((0.1 / ((1.5 * float64((((((0.1 * (v3 + (v3 - 0.1))) * (v3 / v3)) / 0.1) / 3.14159) + (3.14159 + float32(0.1))))) - float64((v3 / (float32(0.0) * (v3 * ((float32(float64(float32(((0.0 - 3.14159) * 1.5)))) + v3) * v3))))))) / ((0.1 / (1e10 - (0.0 / 1.5))) - 1.5)))
                    v30 := []int{-1}
                    v22 = v22
                    _ = v30
                }
                delete(v1, 705)
                _ = fi25
            }
        }
    case "foo bar":
        for fi31 := 0; fi31 < 14; fi31++ {
            switch (("hello" + "a") + "hello") {
            case "\n\t":
                var v32 int = 15
                _ = v32
            default:
                v33 := GenS1{F0: 67, F1: 127, F2: (^(v2))}
                _ = v33
            }
            for fi34 := 0; fi34 < 5; fi34++ {
                fmt.Println(v2)
                _ = fi34
            }
            _ = fi31
        }
    }
    switch "\n\t" {
    case "":
        v2 += (v2 % ((940) | 1))
    case "hello":
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 uint64 = 0
            fmt.Println(v2, v2, fi35)
            _ = fi35
            _ = v36
        }
    case "a":
        v3 = ((v3 * -2.25) + (((v3 * (v3 + v3)) / (v3 / v3)) * v3))
    default:
        v4 = v4
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        fmt.Println(fi37)
        v38 := GenS2{F0: v22, F1: (v22 + 127), F2: int16(fi37)}
        delete(v1, 392)
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 14; fi39++ {
        var v40 float32 = v3
        v41 := []int{(fi39 >> (((v4 & v4)) & 63)), -9223372036854775808, fi39}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
