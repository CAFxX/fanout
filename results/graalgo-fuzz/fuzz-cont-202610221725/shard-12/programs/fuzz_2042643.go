package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint64; F1 int32; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int; F1 uint; F2 uint32 }
func genFn0() uint {
    v4 := []int{}
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            if (1 <= 127) {
                var v7 float64 = (0.1 + float64(-2.25))
                v7 += v7
                var v8 string = (("hello" + "\n\t") + ((("\n\t" + "a") + "a") + "\n\t"))
                _ = v7
                _ = v8
            } else {
                var v9 uint8 = 42
                _ = v9
            }
            var v10 int64 = -9223372036854775808
            v10 += (-(v10))
            _ = fi6
            _ = v10
        }
        switch ("" + "hello") {
        case "a":
            for fi11 := 0; fi11 < 2; fi11++ {
                switch "foo bar" {
                case "foo bar":
                    var v12 float64 = 0.0
                    _ = v12
                case "a":
                    var v13 int = fi5
                    _ = v13
                case "":
                    var v14 float64 = (1.5 + 1.5)
                    _ = v14
                default:
                    var v15 string = "hello"
                    _ = v15
                }
                var v16 uint64 = 716
                v16 = v16
                v16 += v16
                _ = fi11
                _ = v16
            }
        }
        var v17 int32 = -2147483648
        _ = fi5
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v4 = append(v4, 0)
        var v19 string = "hello"
        v20 := GenS1{F0: fi18}
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v4
    return 255
}
func genFn1(p0 int, p1 uint) int {
    for fi21 := 0; fi21 < 8; fi21++ {
        p1 += p1
        if (!(!(-590 != -1))) {
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS2{F0: 42, F1: 42}
                var v24 uint64 = 318
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            p1 = 42
            var v25 int32 = 1
            _ = v25
        }
        var v26 bool = (!(p0 != (fi21 | fi21)))
        _ = fi21
        _ = v26
    }
    p0 = p0
    switch "\n\t" {
    case "\n\t":
        switch ("a" + "a") {
        case "hello":
            v27 := GenS3{F0: uint64(p1), F1: -338, F2: -1}
            _ = v27
        case "":
            p1 += 0
        case "a":
            p0 += (p0 ^ (p0 & 999))
        default:
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int8 = 0
                switch "foo bar" {
                case "hello":
                    var v30 int64 = 9223372036854775807
                    _ = v30
                case "foo bar":
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v32 := GenS3{F0: 42, F1: 0, F2: -113}
                        var v33 uint = genFn0()
                        p0 -= p0
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                case "a":
                    switch "a" {
                    case "\n\t":
                        var v34 int8 = v29
                        _ = v34
                    }
                }
                fmt.Println(p0, fi28)
                _ = fi28
                _ = v29
            }
        }
    case "a":
        v35 := []int{(^(p0)), p0, 2147483647}
        _ = v35
    }
    _ = p0
    _ = p1
    return (p0 / ((p0) | 1))
}

func main() {
    var v1 uint = 255
    v2 := []int{-2147483648, 32768, 936, 4294967295}
    var v3 uint16 = 502
    switch 9223372036854775807 {
    case -1:
        var v36 int64 = 0
        _ = v36
    case 4294967295:
        switch 9223372036854775807 {
        case 128:
            if true {
                v3 = (v3 + v3)
                fmt.Println(v1)
                if false {
                    if (0 < -1) {
                        switch 32767 {
                        case 410:
                            var v37 uint16 = v3
                            _ = v37
                        case 478:
                            if true {
                                switch "a" {
                                case "foo bar":
                                    v1 = (v1 / ((genFn0()) | 1))
                                case "":
                                    v38 := GenS1{F0: 997}
                                    _ = v38
                                case "hello":
                                    var v39 int = -1
                                    _ = v39
                                }
                                var v40 uint16 = v3
                                _ = v40
                            }
                        case 345:
                            fmt.Println(v1, v3)
                        }
                        switch ("" + ("a" + ((((("" + "\n\t") + (((("a" + "") + "") + "a") + "a")) + "a") + "a") + "hello"))) {
                        case "\n\t":
                            v1 = uint(v3)
                        case "":
                            v41 := GenS4{F0: 0}
                            _ = v41
                        }
                    }
                }
            }
        }
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        switch "foo bar" {
        case "foo bar":
            switch 2147483647 {
            case 127:
                var v43 uint8 = 127
                _ = v43
            default:
                v44 := GenS1{F0: fi42}
                _ = v44
            }
        case "hello":
            v1 += v1
        case "a":
            fmt.Println(fi42)
        default:
            for fi45 := 0; fi45 < 5; fi45++ {
                v3 = 0
                fmt.Println(fi42)
                _ = len(v2)
                _ = fi45
            }
        }
        fmt.Println(fi42, v1)
        _ = fi42
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (("hello" + "\n\t") + "a") {
        case "a":
            if true {
                if (false || (!(!true))) {
                    v46 := GenS5{F0: v3, F1: 0, F2: -101}
                    if true {
                        switch "foo bar" {
                        case "":
                            fmt.Println(v1, v1)
                        case "hello":
                            if (true && ((93 >= 65535) && false)) {
                                var v47 int8 = -1
                                if true {
                                    var v48 uint = v1
                                    if (0 > 100) {
                                        v1 -= (^(v48))
                                    } else {
                                        var v49 int = -703
                                        _ = v49
                                    }
                                    if false {
                                        switch -53 {
                                        case 42:
                                            v47 += (^(v47))
                                        case -443:
                                            if true {
                                                fmt.Println(v48)
                                                for fi50 := 0; fi50 < 12; fi50++ {
                                                    v47 += v47
                                                    _ = fi50
                                                }
                                                v48 -= (v48 & v48)
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case 4294967295:
                                            fmt.Println(v47)
                                        default:
                                            v1 += v48
                                        }
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            if len(v2) > 0 {
                                                v2[0] = (^(fi51))
                                            }
                                            if (!(42 > -1)) {
                                                var v52 float32 = ((3.14159 / ((0.1 / 1.0) * 1.5)) + 1e10)
                                                var v53 float32 = (v52 - float32(1e10))
                                                switch -2147483648 {
                                                case 127:
                                                    if len(v2) > 0 {
                                                        v2[0] = genFn1(-1, (v1 + v48))
                                                    }
                                                default:
                                                    if ("\n\t" <= "a") {
                                                        v54 := GenS4{F0: 75}
                                                        v52 = (float32((1e10 - (float64(0.1) + (0.1 + 0.0)))) + 1.5)
                                                        _ = v54
                                                    } else {
                                                        if (true || false) {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v55 := GenS4{F0: 0}
                                                            var v56 int = (fi51 / (((fi51 * (fi51 ^ -699))) | 1))
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    }
                                                }
                                                _ = v52
                                                _ = v53
                                            }
                                            switch "foo bar" {
                                            case "hello":
                                                var v57 float64 = -2.25
                                                _ = v57
                                            case "a":
                                                if false {
                                                    v47 -= v47
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi51 & (fi51 - (fi51 % (((fi51 + fi51)) | 1))))
                                                    }
                                                    if (true && false) {
                                                        v48 = v1
                                                        for fi58 := 0; fi58 < 4; fi58++ {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = fi58
                                                        }
                                                        for fi59 := 0; fi59 < 6; fi59++ {
                                                            var v60 int16 = 32767
                                                            _ = fi59
                                                            _ = v60
                                                        }
                                                    } else {
                                                        if false {
                                                            var v61 int16 = 1
                                                            _ = v61
                                                        } else {
                                                            switch ("\n\t" + ("\n\t" + "a")) {
                                                            case "hello":
                                                                v47 = (v47 | (v47 / (((v47 + 0)) | 1)))
                                                            default:
                                                                v47 = (-(v47))
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v3 -= v3
                                                    v3 += (v3 * (v3 >> (((v48 >> (((^(v48))) & 63))) & 63)))
                                                    _ = len(v2)
                                                }
                                            case "\n\t":
                                                var v62 int32 = -2147483648
                                                _ = v62
                                            default:
                                                switch (fi51 + genFn1(9223372036854775807, uint(v3))) {
                                                case 0:
                                                    v63 := GenS5{F0: v3, F1: 32767, F2: (^(v47))}
                                                    _ = v63
                                                case 2147483647:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 127:
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi51 ^ 128)
                                                    }
                                                }
                                            }
                                            _ = fi51
                                        }
                                    }
                                    _ = v48
                                }
                                switch 127 {
                                case -301:
                                    var v64 int16 = 32767
                                    _ = v64
                                case 32768:
                                    v3 = 0
                                default:
                                    _ = len(v2)
                                }
                                _ = v47
                            } else {
                                if false {
                                    v3 -= v3
                                }
                            }
                        }
                    } else {
                        for fi65 := 0; fi65 < 19; fi65++ {
                            var v66 uint16 = v3
                            fmt.Println(v3, v3, v66)
                            _ = fi65
                            _ = v66
                        }
                        v3 += v3
                        var v67 int64 = 1
                        _ = v67
                    }
                    v3 += v3
                    _ = v46
                } else {
                    for fi68 := 0; fi68 < 17; fi68++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v1 += (v1 + 151)
                        for fi69 := 0; fi69 < 19; fi69++ {
                            if ((127 != 1) && true) {
                                v70 := GenS6{F0: 42}
                                fmt.Println(fi68, fi69)
                                _ = v70
                            }
                            fmt.Println(v3, fi68)
                            _ = fi69
                        }
                        _ = fi68
                    }
                    v1 -= genFn0()
                    var v71 uint8 = 24
                    _ = v71
                }
                for fi72 := 0; fi72 < 8; fi72++ {
                    if true {
                        var v73 uint32 = 2147483647
                        var v74 int8 = int8(fi72)
                        _ = v73
                        _ = v74
                    } else {
                        for fi75 := 0; fi75 < 5; fi75++ {
                            v1 = 4294967295
                            var v76 uint = (v1 % ((genFn0()) | 1))
                            if len(v2) > 0 {
                                v2[0] = (fi72 / (((fi75 + 863)) | 1))
                            }
                            _ = fi75
                            _ = v76
                        }
                        var v77 bool = (639 < fi72)
                        _ = v77
                    }
                    for fi78 := 0; fi78 < 16; fi78++ {
                        if len(v2) > 0 {
                            v2[0] = fi72
                        }
                        switch (fi72 - -968) {
                        case -814:
                            fmt.Println(v3)
                        case -973:
                            v3 = v3
                        default:
                            var v79 int = fi78
                            _ = v79
                        }
                        _ = fi78
                    }
                    _ = fi72
                }
                fmt.Println(v3)
            } else {
                v1 = v1
                v2 = append(v2, 128)
            }
        default:
            v3 = 2
        }
        for fi80 := 0; fi80 < 19; fi80++ {
            fmt.Println(fi80)
            _ = fi80
        }
    }
    var v81 string = ""
    var v82 uint64 = 42
    fmt.Println(v3, v81, v81)
    switch -109 {
    case 127:
        for fi83 := 0; fi83 < 17; fi83++ {
            for fi84 := 0; fi84 < 7; fi84++ {
                v81 = (("foo bar" + ("hello" + v81)) + ((v81 + "hello") + "hello"))
                _ = fi84
            }
            _ = fi83
        }
    case 42:
        v85 := GenS5{F0: v3, F1: 42, F2: 0}
        _ = v85
    case 1:
        for fi86 := 0; fi86 < 7; fi86++ {
            for fi87 := 0; fi87 < 20; fi87++ {
                v2 = append(v2, 9223372036854775807)
                v1 += 865
                _ = fi87
            }
            v82 -= v82
            var v88 uint16 = v3
            _ = fi86
            _ = v88
        }
    default:
        for fi89 := 0; fi89 < 18; fi89++ {
            var v90 int64 = 403
            fmt.Println(v3, v90)
            fmt.Println(v82)
            _ = fi89
            _ = v90
        }
    }
    v2 = append(v2, 1)
    v2 = append(v2, 32767)
    for fi91 := 0; fi91 < 14; fi91++ {
        for fi92 := 0; fi92 < 14; fi92++ {
            var v93 bool = false
            v81 = "hello"
            v93 = v93
            _ = fi92
            _ = v93
        }
        _ = fi91
    }
    switch genFn1(-443, 7) {
    case 32768:
        switch v81 {
        case "":
            var v94 uint32 = 0
            _ = v94
        case "a":
            v95 := GenS1{F0: genFn1(2147483647, (v1 | genFn0()))}
            _ = v95
        default:
            var v96 string = ("\n\t" + "hello")
            _ = v96
        }
    case 272:
        var v97 uint8 = 235
        _ = v97
    case 128:
        if len(v2) > 0 {
            v2[0] = genFn1(205, 65535)
        }
    default:
        var v98 uint32 = 1
        _ = v98
    }
    for fi99 := 0; fi99 < 13; fi99++ {
        fmt.Println(v82)
        _ = fi99
    }
    if false {
        v3 += v3
    }
    v100 := GenS4{F0: 0}
    for fi101 := 0; fi101 < 20; fi101++ {
        switch v81 {
        case "foo bar":
            v81 = v81
        case "\n\t":
            for fi102 := 0; fi102 < 19; fi102++ {
                v103 := GenS5{F0: uint16(v82), F1: -1, F2: 1}
                _ = fi102
                _ = v103
            }
        case "a":
            for fi104 := 0; fi104 < 10; fi104++ {
                v3 = v3
                if true {
                    switch "hello" {
                    case "a":
                        v3 = (^(v3))
                    case "":
                        v105 := GenS6{F0: 161}
                        _ = v105
                    case "hello":
                        if (((-128 >= -128) || true) || (!((v3 <= (v3 % (((^(v3))) | 1))) && (genFn1((fi104 + 42), (v1 ^ (v1 % ((v1) | 1)))) > fi101)))) {
                            if len(v2) > 0 {
                                v2[0] = fi104
                            }
                        }
                    default:
                        if true {
                            v106 := GenS7{F0: (fi101 - (fi104 ^ fi104)), F1: v1, F2: 42}
                            fmt.Println(v1, v82)
                            _ = v106
                        }
                    }
                    _ = len(v2)
                }
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
                _ = fi104
            }
        default:
            v81 = ((((v81 + v81) + v81) + (((v81 + "\n\t") + v81) + ("foo bar" + (v81 + v81)))) + v81)
        }
        _ = fi101
    }
    var v107 bool = (!false)
    switch "\n\t" {
    case "":
        if len(v2) > 0 {
            v2[0] = genFn1(128, (v1 / ((v1) | 1)))
        }
    case "a":
        var v108 uint = genFn0()
        _ = v108
    case "foo bar":
        v1 -= (v1 * 2)
    default:
        v107 = v107
    }
    v3 -= 141
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v100
    _ = v107
    _ = v2
    _ = v3
    _ = v81
    _ = v82
}
