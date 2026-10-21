package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int8 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int }

func main() {
    var v1 int64 = -1
    var v2 uint64 = 42
    v3 := []int{-868, 32767, 32767}
    var v4 int16 = 32767
    var v5 int = -9223372036854775808
    v6 := GenS1{F0: v5, F1: 42, F2: 112}
    for fi7 := 0; fi7 < 20; fi7++ {
        switch (v5 * -2147483648) {
        case 127:
            if ((1e10 * 1.5) > float32((1e10 * 1.5))) {
                if true {
                    var v8 uint = 4294967295
                    var v9 uint64 = v2
                    _ = v8
                    _ = v9
                }
                var v10 bool = false
                var v11 uint16 = 1
                _ = v10
                _ = v11
            } else {
                v2 -= v2
                if ((127 >= 42) && (2 != v2)) {
                    var v12 int8 = 0
                    _ = v12
                } else {
                    v3 = append(v3, (v5 >> ((65535) & 63)))
                }
                if (11 >= (v1 << ((56) & 63))) {
                    fmt.Println(v4, v1)
                    for fi13 := 0; fi13 < 8; fi13++ {
                        switch ("\n\t" + "foo bar") {
                        case "hello":
                            switch ("" + ("foo bar" + ("a" + (("hello" + "") + "a")))) {
                            case "":
                                v3 = append(v3, -810)
                            case "foo bar":
                                v4 += (v4 / (((v4 % ((506) | 1))) | 1))
                            }
                        case "foo bar":
                            v14 := GenS1{F0: -2147483648, F1: 1, F2: 42}
                            _ = v14
                        default:
                            if false {
                                fmt.Println(v2, fi13)
                                if len(v3) > 0 {
                                    v3[0] = fi13
                                }
                                var v15 uint = 169
                                _ = v15
                            }
                        }
                        if (false || (3.14159 == -2.25)) {
                            v16 := GenS1{F0: (v5 * fi13), F1: 255, F2: 127}
                            if true {
                                var v17 uint = 0
                                var v18 string = "a"
                                switch v18 {
                                case "\n\t":
                                    v17 -= (^(v17))
                                }
                                _ = v17
                                _ = v18
                            }
                            _ = v16
                        }
                        _ = fi13
                    }
                }
            }
        }
        var v19 uint64 = v2
        switch "" {
        case "hello":
            for fi20 := 0; fi20 < 15; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v21 := GenS1{F0: (fi7 / ((v5) | 1)), F1: 947, F2: 0}
                if len(v3) > 0 {
                    v3[0] = v5
                }
                _ = fi20
                _ = v21
            }
        case "foo bar":
            var v22 uint8 = 198
            _ = v22
        case "a":
            if false {
                var v23 string = (((("hello" + "a") + "\n\t") + ("\n\t" + "\n\t")) + "a")
                if ((!false) || true) {
                    if true {
                        var v24 uint8 = 255
                        _ = v24
                    }
                    v4 = v4
                }
                _ = v23
            } else {
                v2 = 0
                fmt.Println(fi7, fi7)
            }
        default:
            v2 = v2
        }
        _ = fi7
        _ = v19
    }
    v4 += 500
    v4 += (v4 & (v4 ^ (v4 << ((65535) & 63))))
    v4 = (v4 >> ((1) & 63))
    if (!false) {
        for fi25 := 0; fi25 < 15; fi25++ {
            v26 := GenS2{F0: 1, F1: -1}
            _ = fi25
            _ = v26
        }
        if (0 < 42) {
            var v27 float64 = float64(1.5)
            if true {
                fmt.Println(v4)
                fmt.Println(v27, v2, v4)
                var v28 int64 = 950
                _ = v28
            } else {
                if (65535 != 2147483647) {
                    v1 += v1
                }
                var v29 int8 = -54
                v3 = append(v3, (v5 >> ((0) & 63)))
                _ = v29
            }
            _ = v27
        } else {
            var v30 string = "foo bar"
            if (!(0.1 == -2.25)) {
                var v31 int16 = v4
                _ = v31
            }
            _ = v30
        }
    }
    fmt.Println(v1)
    v3 = append(v3, v5)
    _ = len(v3)
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    switch "a" {
    case "a":
        for fi32 := 0; fi32 < 18; fi32++ {
            switch (fi32 << ((70) & 63)) {
            case 32767:
                var v33 int16 = (^(v4))
                _ = v33
            }
            _ = fi32
        }
    default:
        for fi34 := 0; fi34 < 18; fi34++ {
            _ = len(v3)
            _ = fi34
        }
    }
    var v35 uint32 = 2147483647
    for fi36 := 0; fi36 < 14; fi36++ {
        for fi37 := 0; fi37 < 19; fi37++ {
            v1 += v1
            v1 -= 9223372036854775807
            var v38 uint16 = uint16(v2)
            _ = fi37
            _ = v38
        }
        _ = fi36
    }
    var v39 string = "foo bar"
    if (-891 == 1) {
        _ = len(v3)
        switch "foo bar" {
        case "foo bar":
            var v40 uint32 = (v35 + (^(v35)))
            _ = v40
        case "\n\t":
            if len(v3) > 0 {
                v3[0] = (v5 & 127)
            }
        default:
            v41 := GenS1{F0: v5, F1: 284, F2: -100}
            _ = v41
        }
        switch "" {
        case "\n\t":
            v42 := GenS2{F0: v35, F1: -1}
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        var v44 float64 = 0.1
        var v45 int32 = -2147483648
        switch v39 {
        case "foo bar":
            switch v5 {
            case 32767:
                fmt.Println(v44)
            default:
                v1 -= (v1 + (v1 * v1))
            }
        }
        _ = fi43
        _ = v44
        _ = v45
    }
    if (!(1 <= 65535)) {
        if (-2.25 >= 1e10) {
            for fi46 := 0; fi46 < 19; fi46++ {
                fmt.Println(v35)
                if true {
                    switch v5 {
                    case 4294967295:
                        v2 = (v2 - 4294967295)
                    case 9223372036854775807:
                        v35 = (v35 * (^(v35)))
                    }
                } else {
                    v47 := GenS2{F0: v35, F1: -2147483648}
                    for fi48 := 0; fi48 < 12; fi48++ {
                        switch fi48 {
                        case 32768:
                            v1 += (v1 >> ((2) & 63))
                        case -468:
                            switch v5 {
                            case 2147483647:
                                switch "\n\t" {
                                case "hello":
                                    v4 = v4
                                case "":
                                    v2 = (^(v2))
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    switch ((v39 + (v39 + "foo bar")) + "\n\t") {
                                    case "\n\t":
                                        var v49 uint32 = 2147483647
                                        _ = v49
                                    case "foo bar":
                                        switch (^(fi48)) {
                                        case 1:
                                            if true {
                                                v3 = append(v3, v5)
                                                var v50 bool = true
                                                var v51 float32 = 1e10
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v52 := GenS3{F0: uint8(v35), F1: -790, F2: (v4 * (v4 ^ (-(v4))))}
                                                _ = v52
                                            }
                                        case -1:
                                            var v53 uint8 = 127
                                            _ = v53
                                        case 42:
                                            v2 -= (v2 & v2)
                                        }
                                    }
                                }
                            default:
                                fmt.Println(fi46, v39, v35)
                            }
                        default:
                            if true {
                                v54 := GenS4{F0: 0, F1: (v35 << ((0) & 63))}
                                if (-16 >= 42) {
                                    fmt.Println(v39, v1)
                                }
                                v2 = (v2 + v2)
                                _ = v54
                            } else {
                                if true {
                                    v2 += (^(v2))
                                    var v55 int32 = 42
                                    _ = v55
                                }
                                if len(v3) > 0 {
                                    v3[0] = v5
                                }
                                fmt.Println(v35, fi46, v4)
                                switch (("" + v39) + v39) {
                                case "a":
                                    var v56 int32 = 1
                                    _ = v56
                                case "\n\t":
                                    v35 -= (v35 ^ v35)
                                default:
                                    switch 127 {
                                    case 93:
                                        if true {
                                            var v57 uint8 = 0
                                            if true {
                                                var v58 string = "hello"
                                                v59 := GenS1{F0: (fi48 >> ((18446744073709551615) & 63)), F1: 18446744073709551615, F2: -8}
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if (!true) {
                                                    var v60 int16 = (v4 & (^(v4)))
                                                    v35 += 4294967295
                                                    _ = v60
                                                }
                                                v3 = append(v3, (fi48 % ((-795) | 1)))
                                                fmt.Println(v35, v35, v57)
                                            }
                                            _ = v57
                                        } else {
                                            if ((!true) || false) {
                                                v5 = (fi46 + (v5 + -1))
                                                fmt.Println(v1, fi46, v4)
                                                switch v5 {
                                                case 32768:
                                                    if (!false) {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v35 = (v35 + (^(v35)))
                                                            var v61 string = ((v39 + "hello") + "hello")
                                                            _ = v61
                                                        case "a":
                                                            var v62 int = fi46
                                                            var v63 string = v39
                                                            _ = v62
                                                            _ = v63
                                                        default:
                                                            var v64 int8 = -1
                                                            _ = v64
                                                        }
                                                        v65 := GenS2{F0: v35, F1: -275}
                                                        var v66 uint64 = (v2 ^ (^(v2)))
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        if len(v3) > 0 {
                                                            v3[0] = fi46
                                                        }
                                                        v67 := GenS5{F0: v4, F1: -862}
                                                        _ = v67
                                                    }
                                                case 9223372036854775807:
                                                    v3 = append(v3, (fi46 & (v5 | fi46)))
                                                }
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = fi48
                                                }
                                                if true {
                                                    if false {
                                                        switch (v5 % ((int(v35)) | 1)) {
                                                        case -36:
                                                            fmt.Println(v1, v1, v2)
                                                        case 9223372036854775807:
                                                            v2 = (^(v2))
                                                        default:
                                                            var v68 bool = true
                                                            _ = v68
                                                        }
                                                        var v69 uint32 = 1
                                                        if false {
                                                            fmt.Println(fi48, v39)
                                                            v2 = (v2 * 4294967295)
                                                        } else {
                                                            v2 += v2
                                                            var v70 uint64 = (v2 << ((255) & 63))
                                                            switch fi46 {
                                                            case 691:
                                                                var v71 uint8 = 128
                                                                _ = v71
                                                            }
                                                            _ = v70
                                                        }
                                                        _ = v69
                                                    } else {
                                                        if true {
                                                            var v72 uint32 = (^(v35))
                                                            v73 := GenS6{F0: 0, F1: -115}
                                                            _ = v72
                                                            _ = v73
                                                        } else {
                                                            var v74 int = 127
                                                            var v75 int32 = -321
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        var v76 int32 = 382
                                                        v76 = 940
                                                        _ = v76
                                                    }
                                                    fmt.Println(v1, fi46)
                                                    v5 += (v5 / (((fi48 * 606)) | 1))
                                                }
                                                if (false || false) {
                                                    v5 = (v5 >> ((42) & 63))
                                                    if true {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        var v77 int32 = -813
                                                        v77 = (v77 | (v77 - v77))
                                                        _ = v77
                                                    } else {
                                                        if ((100 < -83) || ((^(v1)) == (v1 + (-(v1))))) {
                                                            if false {
                                                                var v78 uint8 = 149
                                                                fmt.Println(v78, v4)
                                                                var v79 int32 = -2147483648
                                                                var v80 uint16 = 255
                                                                v78 = (v78 >> ((18446744073709551615) & 63))
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                v81 := GenS7{F0: 1}
                                                                if true {
                                                                    v1 += -258
                                                                }
                                                                var v82 float64 = 1.5
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            v83 := GenS8{F0: v39, F1: (v5 << ((0) & 63))}
                                                            var v84 uint8 = 245
                                                            var v85 int = fi48
                                                            v86 := GenS8{F0: v39, F1: -153}
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        } else {
                                                            v35 -= (v35 * v35)
                                                        }
                                                        var v87 float32 = (0.0 / (0.1 + 3.14159))
                                                        _ = v87
                                                    }
                                                    v88 := GenS8{F0: (v39 + (v39 + ("" + "a"))), F1: v5}
                                                    switch v39 {
                                                    case "foo bar":
                                                        if (525 < 65535) {
                                                            var v89 float32 = 0.0
                                                            _ = v89
                                                        }
                                                    case "hello":
                                                        switch v39 {
                                                        case "\n\t":
                                                            v35 += (^(v35))
                                                        case "a":
                                                            if true {
                                                                var v90 uint16 = 32767
                                                                _ = v90
                                                            }
                                                        }
                                                    case "":
                                                        v39 = (v39 + (v39 + ("" + v39)))
                                                    }
                                                    v2 += 1
                                                    _ = v88
                                                }
                                            }
                                            var v91 int = -16
                                            var v92 uint8 = 255
                                            _ = v91
                                            _ = v92
                                        }
                                    }
                                }
                                var v93 uint32 = (v35 - uint32(v2))
                                _ = v93
                            }
                        }
                        _ = fi48
                    }
                    _ = v47
                }
                _ = len(v3)
                _ = fi46
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
