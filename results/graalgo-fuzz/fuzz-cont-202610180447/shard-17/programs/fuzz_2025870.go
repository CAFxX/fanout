package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []uint32{97, 408, 2147483648}
    var v3 uint16 = 0
    var v4 string = (("" + ("hello" + "foo bar")) + ("foo bar" + "\n\t"))
    var v5 int32 = -423
    var v6 uint64 = 1
    v6 = (^(v6))
    var v7 uint64 = 182
    v6 = (v7 + v7)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int32 = -209
        _ = fi8
        _ = v9
    }
    if (2147483647 < v5) {
        var v10 int32 = (v5 & (v5 + v5))
        for fi11 := 0; fi11 < 9; fi11++ {
            fmt.Println(v6)
            v3 -= (v3 << ((2) & 63))
            for fi12 := 0; fi12 < 6; fi12++ {
                if (42 < -84) {
                    switch fi12 {
                    case -1:
                        v2 = append(v2, 0)
                    case 32767:
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            var v13 int = fi11
                            if ((-(v1)) <= (v1 | (-(v1)))) {
                                var v14 uint64 = v6
                                v15 := GenS1{F0: false}
                                var v16 uint64 = (v14 ^ (^(v14)))
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                switch 32767 {
                                case -821:
                                    v10 = v5
                                case -2147483648:
                                    switch -527 {
                                    case 32768:
                                        var v17 uint8 = 255
                                        _ = v17
                                    case 1:
                                        v18 := GenS1{F0: (1 != 882)}
                                        _ = v18
                                    case 2147483647:
                                        v13 = v13
                                    }
                                }
                                fmt.Println(v3, v4)
                            }
                            _ = v13
                        }
                    default:
                        switch fi12 {
                        case 32767:
                            v19 := GenS1{F0: (!(!(301 == v1)))}
                            _ = v19
                        case 32768:
                            v20 := GenS1{F0: false}
                            _ = v20
                        case 128:
                            switch fi11 {
                            case 42:
                                v1 = v1
                            case -909:
                                v5 += 814
                            case -245:
                                v10 += v5
                            default:
                                fmt.Println(fi12)
                            }
                        }
                        switch (fi11 / ((fi11) | 1)) {
                        case -700:
                            var v21 uint32 = 2
                            _ = v21
                        case 4294967295:
                            var v22 float64 = 0.0
                            _ = v22
                        }
                    }
                } else {
                    v1 = (v1 ^ v1)
                    if false {
                        var v23 int = 2147483647
                        if true {
                            var v24 bool = false
                            _ = v24
                        }
                        _ = v23
                    }
                    switch v4 {
                    case "hello":
                        _ = len(v2)
                    case "":
                        var v25 float32 = 3.14159
                        _ = v25
                    default:
                        var v26 int64 = v1
                        _ = v26
                    }
                }
                if len(v2) > 0 {
                    v2[0] = 111
                }
                var v27 string = v4
                switch fi11 {
                case -385:
                    v2 = append(v2, 2147483648)
                case -9223372036854775808:
                    var v28 uint32 = 2147483648
                    _ = v28
                case 127:
                    if (-32768 < -1) {
                        switch "" {
                        case "foo bar":
                            var v29 bool = true
                            _ = v29
                        }
                    }
                }
                var v30 int32 = (v10 << ((255) & 63))
                _ = fi12
                _ = v27
                _ = v30
            }
            _ = fi11
        }
        _ = v10
    } else {
        v3 += 518
    }
    if true {
        if (0.1 < 1.5) {
            v31 := GenS1{F0: ("foo bar" > v4)}
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 10; fi32++ {
            v33 := GenS2{F0: v3, F1: (("hello" + ("\n\t" + ((("a" + "hello") + "foo bar") + (v4 + v4)))) + v4)}
            var v34 int64 = (v1 >> ((0) & 63))
            if true {
                if ((true && true) && ((!(v4 > v4)) || true)) {
                    var v35 uint32 = 567
                    for fi36 := 0; fi36 < 7; fi36++ {
                        if true {
                            var v37 uint = 0
                            var v38 uint = v37
                            v37 -= (v38 % (((v38 + (v38 * (v38 >> ((v38) & 63))))) | 1))
                            _ = v37
                            _ = v38
                        }
                        _ = fi36
                    }
                    _ = v35
                }
                var v39 int32 = (v5 + 514)
                _ = v39
            } else {
                if false {
                    if ((^(fi32)) == fi32) {
                        switch "\n\t" {
                        case "foo bar":
                            switch (v4 + (v4 + (v4 + ("\n\t" + "\n\t")))) {
                            case "hello":
                                v34 += 168
                            default:
                                switch "a" {
                                case "hello":
                                    switch (fi32 | fi32) {
                                    case 128:
                                        for fi40 := 0; fi40 < 3; fi40++ {
                                            var v41 int = (^(fi32))
                                            _ = fi40
                                            _ = v41
                                        }
                                    default:
                                        if true {
                                            v42 := GenS2{F0: v3, F1: v4}
                                            for fi43 := 0; fi43 < 3; fi43++ {
                                                v44 := GenS1{F0: (((float32((0.0 + 1.5)) * 3.14159) > ((((0.1 - ((1.5 - 1e10) * 0.1)) - (1.5 * ((1e10 * 3.14159) * 1e10))) / 1e10) - 0.0)) || (2 < 114))}
                                                if false {
                                                    v1 -= v1
                                                    if true {
                                                        var v45 uint8 = 39
                                                        var v46 float32 = (3.14159 * float32(0.0))
                                                        var v47 uint64 = v6
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    var v48 uint64 = v7
                                                    _ = v48
                                                } else {
                                                    _ = len(v2)
                                                }
                                                _ = fi43
                                                _ = v44
                                            }
                                            var v49 uint8 = uint8(v7)
                                            _ = v42
                                            _ = v49
                                        } else {
                                            fmt.Println(v34)
                                        }
                                    }
                                case "foo bar":
                                    var v50 uint16 = v3
                                    _ = v50
                                }
                            }
                        }
                        v34 -= (v1 * (v34 | (v34 * v34)))
                        if true {
                            var v51 int64 = -1
                            var v52 int32 = v5
                            _ = v51
                            _ = v52
                        }
                    }
                    var v53 uint32 = uint32(v3)
                    _ = v53
                }
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        for fi54 := 0; fi54 < 17; fi54++ {
            var v55 uint8 = 0
            v6 += v6
            v5 -= v5
            _ = fi54
            _ = v55
        }
        _ = len(v2)
    }
    v3 = v3
    fmt.Println(v4)
    switch -9223372036854775808 {
    case -9223372036854775808:
        for fi56 := 0; fi56 < 6; fi56++ {
            var v57 float64 = (float64((((3.14159 - 1.5) / 3.14159) + 1e10)) / 1.5)
            switch (("" + (("\n\t" + ("\n\t" + v4)) + ((v4 + "hello") + v4))) + v4) {
            case "a":
                var v58 int = fi56
                _ = v58
            }
            _ = fi56
            _ = v57
        }
    case 42:
        if len(v2) > 0 {
            v2[0] = 990
        }
    case -46:
        var v59 int64 = (v1 ^ v1)
        _ = v59
    default:
        switch (v4 + (v4 + v4)) {
        case "foo bar":
            v2 = append(v2, 65535)
        case "":
            if len(v2) > 0 {
                v2[0] = uint32(v6)
            }
        case "\n\t":
            fmt.Println(v5, v3, v1)
        default:
            _ = len(v2)
        }
    }
    if true {
        v60 := GenS1{F0: true}
        v6 += 18446744073709551615
        fmt.Println(v7, v7)
        _ = v60
    }
    for fi61 := 0; fi61 < 20; fi61++ {
        if true {
            for fi62 := 0; fi62 < 4; fi62++ {
                _ = len(v2)
                _ = fi62
            }
            for fi63 := 0; fi63 < 11; fi63++ {
                if false {
                    if false {
                        v6 = v7
                    }
                    v3 -= v3
                    v64 := GenS1{F0: (363 > 18446744073709551615)}
                    _ = v64
                } else {
                    v2 = append(v2, 51)
                    v65 := GenS3{F0: int32(v3)}
                    v3 -= (v3 / ((v3) | 1))
                    _ = v65
                }
                _ = fi63
            }
        }
        fmt.Println(v6, v5, v7)
        _ = fi61
    }
    for fi66 := 0; fi66 < 20; fi66++ {
        for fi67 := 0; fi67 < 20; fi67++ {
            v4 = (v4 + v4)
            fmt.Println(v4)
            var v68 string = ("" + "foo bar")
            _ = fi67
            _ = v68
        }
        var v69 uint64 = v6
        _ = fi66
        _ = v69
    }
    v3 = v3
    v70 := GenS4{F0: int16(v6), F1: v7}
    v71 := GenS5{F0: 4294967295, F1: 127}
    v3 -= (v3 + 65535)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v70
    _ = v71
}
