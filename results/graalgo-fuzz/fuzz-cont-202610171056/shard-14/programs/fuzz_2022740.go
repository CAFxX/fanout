package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 bool }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int64; F2 string }

func main() {
    var v1 int32 = 1
    var v2 uint8 = 128
    var v3 int32 = (v1 ^ -2147483648)
    v4 := []uint32{}
    var v5 bool = false
    for fi6 := 0; fi6 < 11; fi6++ {
        v2 = 0
        var v7 int = fi6
        if (2 < 0) {
            v8 := GenS1{F0: v5, F1: v1, F2: (!((5 != 127) || (255 < 116)))}
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int16 = int16(v3)
                _ = fi9
                _ = v10
            }
            switch v7 {
            case -2147483648:
                if true {
                    v4 = append(v4, 1)
                } else {
                    if (!(!v5)) {
                        v7 = -2147483648
                        for fi11 := 0; fi11 < 20; fi11++ {
                            if len(v4) > 0 {
                                v4[0] = 65535
                            }
                            var v12 bool = (146 < 0)
                            _ = fi11
                            _ = v12
                        }
                        var v13 float64 = float64(0.1)
                        _ = v13
                    }
                    var v14 uint8 = (v2 >> ((4294967295) & 63))
                    _ = v14
                }
            case 32767:
                fmt.Println(v2)
            default:
                fmt.Println(v2, v5, v2)
            }
            _ = v8
        } else {
            fmt.Println(v3, v3, v7)
            v3 += (v1 >> ((62) & 63))
        }
        _ = fi6
        _ = v7
    }
    v3 = (-(v1))
    if (-32768 <= 180) {
        if len(v4) > 0 {
            v4[0] = 42
        }
        var v15 uint16 = 0
        _ = v15
    }
    switch 1 {
    case 0:
        fmt.Println(v3, v3, v2)
    default:
        v16 := GenS1{F0: ((-2.25 * 1.5) == 1e10), F1: (v3 + -2147483648), F2: (!((0.0 / 1e10) > ((-2.25 * float64(0.1)) / 3.14159)))}
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 16; fi17++ {
            switch "hello" {
            case "hello":
                v3 = (^(v1))
            case "a":
                v18 := GenS2{F0: 127, F1: 0, F2: int(v3)}
                _ = v18
            }
            _ = fi17
        }
        switch "hello" {
        case "a":
            v4 = append(v4, 2147483647)
        default:
            switch 558 {
            case 938:
                var v19 float64 = 1.5
                _ = v19
            case -679:
                var v20 float64 = 1e10
                _ = v20
            default:
                fmt.Println(v1)
            }
        }
    }
    if ((v5 || (!(!false))) && (-1 >= -740)) {
        v21 := GenS3{F0: 32767}
        _ = v21
    }
    v4 = append(v4, 2147483648)
    v3 = v1
    fmt.Println(v1, v5)
    v22 := GenS4{F0: (v5 || v5), F1: ((1.5 == 1e10) || v5)}
    var v23 uint = 271
    v24 := GenS4{F0: (!v5), F1: (!v5)}
    switch -454 {
    case 42:
        v25 := GenS5{F0: (v2 % ((128) | 1)), F1: 0}
        _ = v25
    case -829:
        v23 += (^(v23))
    case 322:
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v3, fi26)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v27 uint16 = 32767
            _ = fi26
            _ = v27
        }
    }
    v28 := GenS6{F0: (v2 ^ (v2 & v2)), F1: 42, F2: "\n\t"}
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
