package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = (3.14159 * (1.5 * 1e10))
    var v2 int = 32767
    var v3 uint16 = 349
    v4 := make(map[uint32]bool)
    var v5 int8 = 78
    v6 := GenS1{F0: (v2 << ((0) & 63))}
    v7 := GenS2{F0: (v5 & int8(v3)), F1: v5}
    switch "\n\t" {
    case "hello":
        _ = len(v4)
    case "":
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS3{F0: 2147483647, F1: 2}
            var v10 int32 = -90
            var v11 uint16 = 1
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    default:
        if false {
            v5 -= v5
        } else {
            v1 = (float32(3.14159) - (float32(-2.25) * v1))
            v2 -= (v2 - -659)
            v2 += v2
        }
    }
    v12 := []int{(^(v2))}
    var v13 float32 = v1
    for fi14 := 0; fi14 < 16; fi14++ {
        v2 = v2
        v4[65535] = ((true && ((v3 | v3) == v3)) && ((v5 << ((2) & 63)) < (v5 / (((^(v5))) | 1))))
        fmt.Println(v5, v1)
        _ = fi14
    }
    switch "" {
    case "":
        switch "hello" {
        case "a":
            switch v2 {
            case 1:
                v3 = 255
            case 574:
                for fi15 := 0; fi15 < 2; fi15++ {
                    if (!(false || false)) {
                        fmt.Println(v5)
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v2 += v2
                            v5 += (v5 << ((0) & 63))
                            v17 := GenS1{F0: -723}
                            _ = fi16
                            _ = v17
                        }
                        fmt.Println(v5)
                    } else {
                        v13 -= v13
                        switch "a" {
                        case "a":
                            switch -2147483648 {
                            case 128:
                                var v18 bool = (!true)
                                _ = v18
                            case 502:
                                v19 := GenS4{F0: 18446744073709551615, F1: 1000, F2: 18446744073709551615}
                                _ = v19
                            default:
                                for fi20 := 0; fi20 < 15; fi20++ {
                                    v21 := GenS2{F0: (v5 | v5), F1: v5}
                                    v3 = 2
                                    v5 = (v5 << ((42) & 63))
                                    fmt.Println(fi20, v13, v13)
                                    v22 := GenS3{F0: 683, F1: 1}
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                            }
                        default:
                            v5 -= v5
                        }
                    }
                    var v23 uint8 = 2
                    v13 = v13
                    _ = fi15
                    _ = v23
                }
            case 32767:
                fmt.Println(v2, v3)
            }
        case "hello":
            switch v2 {
            case 32768:
                var v24 bool = (255 == 42)
                _ = v24
            }
        }
    case "\n\t":
        switch int(v3) {
        case 32768:
            v4[4294967295] = (((2 < 127) && (("" + ("a" + ("a" + "foo bar"))) >= "\n\t")) || (!(32767 > 1000)))
        case 0:
            v3 = (v3 - v3)
        default:
            switch v2 {
            case 341:
                fmt.Println(v13, v5, v3)
            case -9223372036854775808:
                if len(v12) > 0 {
                    v12[0] = 245
                }
            }
        }
    }
    v25 := GenS3{F0: 2147483648, F1: 2}
    if (!true) {
        if ((uint32(v3) > 2) && ((false && true) || true)) {
            v5 = (v5 & v5)
        }
        v5 = 1
    } else {
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 string = (("foo bar" + "") + ((("a" + ((("a" + ("hello" + "hello")) + ("a" + "\n\t")) + "hello")) + "foo bar") + ""))
            _ = fi26
            _ = v27
        }
        var v28 uint = 606
        var v29 int8 = 62
        _ = v28
        _ = v29
    }
    v2 += v2
    if (v1 != (v13 / v1)) {
        v1 -= v1
    } else {
        var v30 uint8 = 1
        for fi31 := 0; fi31 < 13; fi31++ {
            var v32 uint = 0
            for fi33 := 0; fi33 < 12; fi33++ {
                v34 := GenS5{F0: 1, F1: 0}
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = v30
    }
    v13 = v13
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
