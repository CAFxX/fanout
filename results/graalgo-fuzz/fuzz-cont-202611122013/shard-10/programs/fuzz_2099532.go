package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int64 = 1
    var v2 uint8 = 255
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 3; fi4++ {
        if true {
            v1 = (v1 << ((uint(v2)) & 63))
        }
        v3[42] = (!(858 >= 1))
        _ = fi4
    }
    fmt.Println(v1)
    v5 := []int{2147483647, 836, 865}
    var v6 uint8 = (^(v2))
    if false {
        var v7 int8 = 127
        for fi8 := 0; fi8 < 17; fi8++ {
            v5 = append(v5, (fi8 >> ((2) & 63)))
            var v9 float32 = -2.25
            _ = fi8
            _ = v9
        }
        if (32767 >= 638) {
            fmt.Println(v7, v1)
            v10 := GenS1{F0: -1, F1: 2}
            for fi11 := 0; fi11 < 2; fi11++ {
                v2 = v6
                v12 := GenS2{F0: -650}
                v2 = (v6 >> ((uint(v2)) & 63))
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            v6 = v2
            v13 := GenS1{F0: 919, F1: 510}
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14, v1)
                for fi15 := 0; fi15 < 1; fi15++ {
                    v16 := GenS1{F0: 32767, F1: 65535}
                    v7 = -128
                    _ = fi15
                    _ = v16
                }
                v7 += v7
                _ = fi14
            }
            _ = v13
        }
        _ = v7
    } else {
        var v17 int32 = 42
        switch -152 {
        case 4294967295:
            switch 127 {
            case 32767:
                delete(v3, 42)
            case -9223372036854775808:
                if true {
                    var v18 int16 = int16(v17)
                    v17 = (v17 >> ((0) & 63))
                    var v19 string = (("foo bar" + "") + ("a" + "a"))
                    _ = v18
                    _ = v19
                }
            default:
                var v20 uint32 = 0
                _ = v20
            }
        case -567:
            var v21 uint = 65535
            _ = v21
        case -222:
            fmt.Println(v2, v1)
        default:
            v3[1] = (2 == 1)
        }
        var v22 uint16 = 710
        _ = v17
        _ = v22
    }
    var v23 uint = 1
    switch "hello" {
    case "hello":
        if false {
            fmt.Println(v2, v23, v23)
            v1 += (^(v1))
        } else {
            v6 = (v2 * (^(v2)))
            v1 -= (v1 ^ v1)
            v23 -= 42
        }
    case "foo bar":
        if true {
            v2 += v6
        } else {
            v6 -= v6
        }
    default:
        switch 4294967295 {
        case 32768:
            var v24 uint32 = 2147483647
            _ = v24
        default:
            var v25 int32 = 1
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        if (!false) {
            v3[2] = ((-2.25 + float32(((-2.25 - float64((0.1 + 3.14159))) / 1e10))) >= float32(1e10))
        }
        fmt.Println(v23, fi26)
        _ = fi26
    }
    fmt.Println(v2, v1, v2)
    var v27 int32 = 2147483647
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch -457 {
    case 0:
        v27 += (v27 + 28)
    case -625:
        v2 -= v6
    case 709:
        switch "a" {
        case "":
            var v28 int32 = v27
            _ = v28
        }
    default:
        fmt.Println(v2, v27)
    }
    var v29 float64 = ((((1.5 * (float64(1e10) * ((1e10 / (((0.1 + float64(3.14159)) / float64(0.1)) * (0.1 / 1e10))) / (0.1 / 1.0)))) * ((0.1 / 1.0) + (0.1 - (0.0 - (0.1 / ((float64(float32(((1.5 / 1e10) / 1e10))) * 0.0) + 3.14159)))))) * ((float64(3.14159) - (1.5 + (0.1 - (3.14159 + 1.5)))) * ((float64(3.14159) + ((float64(-2.25) + ((((1.5 + 0.1) - 3.14159) + (1.5 * 0.1)) * 0.0)) + 1e10)) / (0.0 + 1e10)))) - 1.5)
    v5 = append(v5, 32768)
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS2{F0: (v27 - (v27 + 42))}
        var v32 float64 = ((1e10 / (0.0 * v29)) + float64(1e10))
        _ = fi30
        _ = v31
        _ = v32
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v33 int32 = (v27 * (v27 - (^(v27))))
    v34 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v5
    _ = v6
}
