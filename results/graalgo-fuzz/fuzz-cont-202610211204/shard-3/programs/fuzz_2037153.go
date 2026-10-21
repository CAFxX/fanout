package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 int; F2 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int64; F1 uint64; F2 int64 }
type GenS14 struct { F0 int64; F1 string; F2 uint16 }
type GenS15 struct { F0 uint64; F1 bool; F2 int32 }
type GenS16 struct { F0 uint8; F1 uint8; F2 uint }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 uint8; F1 string }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 bool; F1 uint32 }
type GenS21 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS22 struct { F0 bool; F1 uint8 }
type GenS23 struct { F0 bool; F1 int64; F2 int16 }
type GenS24 struct { F0 uint32; F1 int64; F2 uint }
type GenS25 struct { F0 int8 }
type GenS26 struct { F0 int; F1 int }
type GenS27 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS28 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS29 struct { F0 uint32; F1 int32 }
type GenS30 struct { F0 bool }
type GenS31 struct { F0 uint32 }
type GenS32 struct { F0 uint64; F1 int }
func genFn0(p0 int16) int16 {
    var v6 int32 = 2147483647
    v7 := GenS1{F0: (v6 & int32(p0))}
    v6 = (v6 * v6)
    _ = p0
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    var v3 int = 1
    v4 := make(map[int]uint32)
    v5 := []uint32{2147483647}
    v8 := GenS1{F0: 603}
    if ((1 != -212) || (!true)) {
        for fi9 := 0; fi9 < 16; fi9++ {
            v3 = v3
            v2 -= v2
            switch "a" {
            case "hello":
                var v10 int = (-(fi9))
                _ = v10
            default:
                switch fi9 {
                case 859:
                    v3 = 2147483647
                default:
                    for fi11 := 0; fi11 < 19; fi11++ {
                        fmt.Println(fi11)
                        var v12 int = v3
                        switch "\n\t" {
                        case "a":
                            switch "hello" {
                            case "hello":
                                _ = len(v4)
                            case "":
                                if true {
                                    v13 := GenS1{F0: 1}
                                    v5 = append(v5, 2)
                                    if false {
                                        v14 := GenS1{F0: -2147483648}
                                        v4[127] = 2
                                        v15 := GenS1{F0: -2147483648}
                                        _ = v14
                                        _ = v15
                                    } else {
                                        var v16 float64 = ((0.0 + ((((((0.0 - 3.14159) - ((0.0 * -2.25) / 1.0)) / ((0.1 - 1e10) + (1.5 - ((1e10 / 0.1) / -2.25)))) - (((3.14159 * 1e10) - float64(float32(1e10))) - (((-2.25 * (((-2.25 - (-2.25 * -2.25)) + 0.1) + 1e10)) + 0.0) - 1.5))) / 3.14159) * (((0.0 + (1.5 + 1e10)) - -2.25) - (0.0 * float64(0.1))))) + (3.14159 - 0.1))
                                        if len(v5) > 0 {
                                            v5[0] = 4294967295
                                        }
                                        _ = v16
                                    }
                                    _ = v13
                                } else {
                                    v4[-1] = 65535
                                    if false {
                                        v12 = (^(fi11))
                                        if false {
                                            if (!(80 == 0)) {
                                                v12 += fi9
                                            }
                                        }
                                    }
                                    if false {
                                        v3 = v3
                                    } else {
                                        if false {
                                            if false {
                                                var v17 uint32 = 0
                                                _ = v17
                                            }
                                        }
                                    }
                                }
                            case "foo bar":
                                var v18 int = fi9
                                _ = v18
                            }
                        case "":
                            v2 = 539
                            var v19 float32 = (((1e10 + float32(3.14159)) / -2.25) / 1e10)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v12 = -405
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = v19
                        default:
                            switch (v12 | 42) {
                            case -872:
                                switch "a" {
                                case "foo bar":
                                    v2 = 1
                                case "a":
                                    v20 := v1[(-(fi9))]
                                    _ = v20
                                default:
                                    var v21 int64 = 1
                                    _ = v21
                                }
                            case 585:
                                if true {
                                    fmt.Println(v3, fi9)
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v22 uint8 = 0
                                        _ = v22
                                    case "hello":
                                        if ((v12 & fi11) == v3) {
                                            v23 := v4[fi9]
                                            switch "foo bar" {
                                            case "hello":
                                                v2 -= (v2 - 42)
                                            case "":
                                                switch 159 {
                                                case 2147483647:
                                                    var v24 float32 = 3.14159
                                                    _ = v24
                                                case 127:
                                                    v25 := v4[619]
                                                    _ = v25
                                                case 32767:
                                                    v12 = (v12 << ((4294967295) & 63))
                                                default:
                                                    switch 4294967295 {
                                                    case -2:
                                                        var v26 float64 = (0.0 * 0.1)
                                                        _ = v26
                                                    case 2147483647:
                                                        switch v12 {
                                                        case 42:
                                                            if false {
                                                                if ((v23 + (v23 / ((v23) | 1))) < (^(v23))) {
                                                                    var v27 float32 = (3.14159 - 0.0)
                                                                    v28 := GenS2{F0: 22}
                                                                    v29 := GenS2{F0: 0}
                                                                    _ = v27
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                                v3 -= (fi9 / (((fi11 - (v3 << ((42) & 63)))) | 1))
                                                            }
                                                        }
                                                    }
                                                }
                                                var v30 uint8 = 127
                                                if ((v30 - v30) > v30) {
                                                    var v31 float64 = 1.5
                                                    _ = v31
                                                }
                                                _ = v30
                                            default:
                                                switch 2147483647 {
                                                case -1:
                                                    switch ("a" + "\n\t") {
                                                    case "a":
                                                        v23 = (v23 >> ((803) & 63))
                                                    }
                                                case 426:
                                                    var v32 int32 = 1
                                                    _ = v32
                                                case 32768:
                                                    var v33 int64 = 746
                                                    _ = v33
                                                default:
                                                    v12 -= -2147483648
                                                }
                                            }
                                            switch 99 {
                                            case 926:
                                                v23 -= (^(v23))
                                            case 0:
                                                fmt.Println(v12, v23, fi11)
                                            case 32768:
                                                var v34 float32 = 0.1
                                                _ = v34
                                            }
                                            _ = v23
                                        } else {
                                            fmt.Println(v12, v12)
                                        }
                                    }
                                }
                            default:
                                v35 := GenS3{F0: 9223372036854775807, F1: true, F2: ""}
                                _ = v35
                            }
                        }
                        v5 = append(v5, 0)
                        fmt.Println(fi11, v2)
                        var v36 int16 = -32768
                        _ = fi11
                        _ = v12
                        _ = v36
                    }
                }
            }
            _ = fi9
        }
        switch (v3 % (((v3 / ((v3) | 1))) | 1)) {
        case -9223372036854775808:
            if false {
                _ = len(v4)
                for fi37 := 0; fi37 < 5; fi37++ {
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v39 := v1[fi37]
                        _ = len(v5)
                        _ = fi38
                        _ = v39
                    }
                    _ = fi37
                }
            } else {
                var v40 uint32 = 848
                _ = v40
            }
        case 2147483647:
            for fi41 := 0; fi41 < 9; fi41++ {
                v2 -= 1
                _ = fi41
            }
        case 0:
            var v42 int64 = -951
            _ = v42
        default:
            if false {
                v43 := GenS2{F0: 1}
                for fi44 := 0; fi44 < 8; fi44++ {
                    v3 -= 32768
                    _ = fi44
                }
                v45 := GenS4{F0: true, F1: -1}
                _ = v43
                _ = v45
            }
        }
        if (((!(false && (0 != 18446744073709551615))) && false) && false) {
            if true {
                var v46 float64 = -2.25
                v47 := GenS5{F0: -306, F1: 43}
                v48 := GenS1{F0: 1}
                _ = v46
                _ = v47
                _ = v48
            }
            v2 += v2
        } else {
            var v49 bool = false
            var v50 int16 = -1
            _ = v49
            _ = v50
        }
    }
    switch "hello" {
    case "hello":
        v3 = (v3 % (((v3 * v3)) | 1))
    case "":
        if false {
            if (0 == -1) {
                var v51 uint = 42
                v51 = 349
                var v52 uint32 = 2
                _ = v51
                _ = v52
            }
            for fi53 := 0; fi53 < 16; fi53++ {
                v2 = 0
                if true {
                    var v54 string = "hello"
                    _ = v54
                }
                _ = fi53
            }
            var v55 int64 = -1
            _ = v55
        }
    default:
        v2 = (v2 | 122)
    }
    if (0 <= 0) {
        switch v3 {
        case 895:
            var v56 int64 = 1
            _ = v56
        case 4294967295:
            for fi57 := 0; fi57 < 2; fi57++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v58 := GenS5{F0: 978, F1: -574}
                for fi59 := 0; fi59 < 15; fi59++ {
                    v3 += v3
                    if (!true) {
                        v3 -= 1
                        if true {
                            var v60 bool = ((-1 > 0) || true)
                            fmt.Println(fi59, fi59)
                            v3 -= (fi57 % ((2147483647) | 1))
                            _ = v60
                        }
                        var v61 string = ""
                        var v62 uint16 = v2
                        _ = v61
                        _ = v62
                    } else {
                        var v63 float32 = ((0.1 / 1.0) - float32(3.14159))
                        _ = v63
                    }
                    switch "" {
                    case "hello":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "a":
                        v5 = append(v5, 4294967295)
                    case "\n\t":
                        switch "hello" {
                        case "a":
                            fmt.Println(fi59, v2)
                        case "hello":
                            if true {
                                if (-1 != 0) {
                                    if false {
                                        var v64 int64 = 9223372036854775807
                                        fmt.Println(v3)
                                        _ = v64
                                    }
                                }
                                v3 = -733
                                if false {
                                    v2 -= (^(v2))
                                    v65 := v1[fi57]
                                    v65 = v65
                                    _ = v65
                                }
                                var v66 uint64 = 2
                                _ = v66
                            } else {
                                if (v3 != (^(v3))) {
                                    fmt.Println(v2)
                                    switch (("\n\t" + "\n\t") + "foo bar") {
                                    case "":
                                        switch (("" + ("foo bar" + "")) + "hello") {
                                        case "foo bar":
                                            if ("\n\t" < "foo bar") {
                                                switch ("" + "\n\t") {
                                                case "foo bar":
                                                    v5 = append(v5, 65535)
                                                default:
                                                    var v67 uint64 = 1
                                                    _ = v67
                                                }
                                                _ = len(v5)
                                                v68 := v4[4294967295]
                                                _ = v68
                                            }
                                        }
                                    default:
                                        if false {
                                            v69 := GenS6{F0: int64(fi59)}
                                            _ = len(v5)
                                            _ = v69
                                        }
                                    }
                                } else {
                                    switch fi57 {
                                    case 128:
                                        var v70 string = ""
                                        _ = v70
                                    default:
                                        v3 += (^(fi57))
                                        fmt.Println(fi59, v3, v2)
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = 65535
                                }
                                var v71 string = "foo bar"
                                if true {
                                    switch 32767 {
                                    case 32767:
                                        if (!true) {
                                            switch fi57 {
                                            case -410:
                                                v2 += v2
                                            case -9223372036854775808:
                                                if false {
                                                    if true {
                                                        fmt.Println(fi59)
                                                        v72 := GenS6{F0: 215}
                                                        var v73 int32 = -303
                                                        v74 := GenS5{F0: (^(v73)), F1: 1000}
                                                        v2 = (v2 >> ((42) & 63))
                                                        _ = v72
                                                        _ = v73
                                                        _ = v74
                                                    } else {
                                                        var v75 int = (v3 << ((734) & 63))
                                                        _ = v75
                                                    }
                                                    _ = len(v4)
                                                    if (false && false) {
                                                        v71 = (v71 + ((v71 + v71) + "foo bar"))
                                                    } else {
                                                        switch (v71 + (v71 + ((v71 + ("" + ("a" + "hello"))) + v71))) {
                                                        case "a":
                                                            v3 += fi57
                                                        case "\n\t":
                                                            if true {
                                                                switch fi59 {
                                                                case -1:
                                                                    fmt.Println(fi59)
                                                                default:
                                                                    v76 := GenS7{F0: fi59, F1: fi57, F2: (2 != 2)}
                                                                    v71 = (v71 + v71)
                                                                    switch 0 {
                                                                    case 4294967295:
                                                                        var v77 uint64 = 18446744073709551615
                                                                        var v78 int32 = -750
                                                                        _ = v77
                                                                        _ = v78
                                                                    default:
                                                                        fmt.Println(fi59)
                                                                    }
                                                                    if false {
                                                                        switch "a" {
                                                                        case "hello":
                                                                            fmt.Println(v71)
                                                                        default:
                                                                            switch fi57 {
                                                                            case 1:
                                                                                fmt.Println(v71)
                                                                            case 32767:
                                                                                fmt.Println(fi57, v71)
                                                                            }
                                                                        }
                                                                        v79 := GenS8{F0: -9223372036854775808}
                                                                        _ = v79
                                                                    }
                                                                    v4[(fi57 ^ v3)] = 42
                                                                    _ = v76
                                                                }
                                                                v71 = ""
                                                                if true {
                                                                    v80 := v1[-1]
                                                                    if (65535 >= 0) {
                                                                        var v81 bool = (fi57 < 32767)
                                                                        _ = v81
                                                                    } else {
                                                                        v80 = v80
                                                                        if (1 > -228) {
                                                                            switch (-(v3)) {
                                                                            case 2147483647:
                                                                                if (((false || false) && true) && true) {
                                                                                    fmt.Println(v3, v71, v71)
                                                                                }
                                                                            case -1:
                                                                                if true {
                                                                                    v82 := v4[(fi59 >> ((18446744073709551615) & 63))]
                                                                                    switch v71 {
                                                                                    case "\n\t":
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            if true {
                                                                                                var v83 uint64 = 270
                                                                                                fmt.Println(v82, v2, fi57)
                                                                                                _ = v83
                                                                                            } else {
                                                                                                delete(v4, 9223372036854775807)
                                                                                                if (-2.25 < (0.1 + ((-2.25 / 1.5) * ((0.1 / 0.1) + (1.5 * float32(1e10)))))) {
                                                                                                    if true {
                                                                                                        v1[(^(fi57))] = v71
                                                                                                        v82 += v82
                                                                                                    } else {
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            v82 -= 4294967295
                                                                                                        case "\n\t":
                                                                                                            var v84 uint = 255
                                                                                                            _ = v84
                                                                                                        default:
                                                                                                            v85 := GenS8{F0: 1}
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        if false {
                                                                                                            if false {
                                                                                                                v82 = v82
                                                                                                                v86 := GenS4{F0: (454 <= 1000), F1: 717}
                                                                                                                _ = v86
                                                                                                            } else {
                                                                                                                v87 := GenS8{F0: -1}
                                                                                                                _ = v87
                                                                                                            }
                                                                                                            var v88 uint32 = (v82 - 2147483647)
                                                                                                            if false {
                                                                                                                var v89 int64 = 42
                                                                                                                var v90 int32 = 1
                                                                                                                v90 = (v90 - v90)
                                                                                                                _ = v89
                                                                                                                _ = v90
                                                                                                            }
                                                                                                            _ = len(v5)
                                                                                                            if len(v5) > 0 {
                                                                                                                v5[0] = v82
                                                                                                            }
                                                                                                            v82 = (v82 + (v88 / (((v88 * v82)) | 1)))
                                                                                                            var v91 float32 = ((-2.25 + 1e10) + (1.5 / ((1e10 + 0.1) - (3.14159 + -2.25))))
                                                                                                            v92 := GenS5{F0: 719, F1: -32768}
                                                                                                            v93 := GenS1{F0: 2147483647}
                                                                                                            switch fi59 {
                                                                                                            case -9223372036854775808:
                                                                                                                v2 += (v2 & v2)
                                                                                                            case 42:
                                                                                                                v94 := GenS9{F0: (v71 + v80), F1: (-871 == -192)}
                                                                                                                switch v80 {
                                                                                                                case "foo bar":
                                                                                                                    if (8 < 18446744073709551615) {
                                                                                                                        var v95 string = "\n\t"
                                                                                                                        _ = len(v5)
                                                                                                                        _ = v95
                                                                                                                    } else {
                                                                                                                        if (!false) {
                                                                                                                            v96 := GenS10{F0: 32, F1: (v80 + (((("a" + v80) + "hello") + (v71 + (((v80 + (("a" + "\n\t") + "a")) + (v80 + v71)) + v80))) + "\n\t"))}
                                                                                                                            _ = v96
                                                                                                                        } else {
                                                                                                                            v97 := GenS3{F0: -9223372036854775808, F1: false, F2: "hello"}
                                                                                                                            _ = v97
                                                                                                                        }
                                                                                                                        if false {
                                                                                                                            v3 -= (fi59 ^ fi59)
                                                                                                                        } else {
                                                                                                                            var v98 uint32 = v88
                                                                                                                            v82 += 982
                                                                                                                            switch fi57 {
                                                                                                                            case -1:
                                                                                                                                if true {
                                                                                                                                    var v99 bool = false
                                                                                                                                    if v99 {
                                                                                                                                        v100 := GenS11{F0: 2, F1: (fi57 | (v3 * fi57))}
                                                                                                                                        _ = v100
                                                                                                                                    }
                                                                                                                                    _ = v99
                                                                                                                                }
                                                                                                                            case 4294967295:
                                                                                                                                var v101 int64 = -1
                                                                                                                                _ = v101
                                                                                                                            case 127:
                                                                                                                                v102 := GenS12{F0: ((v80 + ("foo bar" + ("foo bar" + ((("" + (((v80 + v80) + v80) + v80)) + v71) + v71)))) + ((v71 + (v80 + "a")) + (v71 + v71)))}
                                                                                                                                if false {
                                                                                                                                    if (42 >= 0) {
                                                                                                                                        if true {
                                                                                                                                            v88 -= v82
                                                                                                                                            v3 = (fi59 | fi57)
                                                                                                                                            v103 := GenS8{F0: 422}
                                                                                                                                            _ = v103
                                                                                                                                        } else {
                                                                                                                                            var v104 string = v71
                                                                                                                                            var v105 int32 = 2147483647
                                                                                                                                            v106 := v4[fi59]
                                                                                                                                            _ = v104
                                                                                                                                            _ = v105
                                                                                                                                            _ = v106
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v98 = 276
                                                                                                                                        v5 = append(v5, v98)
                                                                                                                                        var v107 int = v3
                                                                                                                                        _ = v107
                                                                                                                                    }
                                                                                                                                    v108 := GenS5{F0: 2147483647, F1: 1}
                                                                                                                                    _ = v108
                                                                                                                                }
                                                                                                                                _ = v102
                                                                                                                            }
                                                                                                                            fmt.Println(v98, v2, v71)
                                                                                                                            if (fi57 >= -1) {
                                                                                                                                v109 := GenS11{F0: 1, F1: fi57}
                                                                                                                                fmt.Println(v82)
                                                                                                                                _ = len(v1)
                                                                                                                                var v110 int = fi59
                                                                                                                                v80 = ((v71 + ((("hello" + (v71 + v71)) + ((("hello" + ("\n\t" + (v80 + v71))) + v80) + (v71 + "hello"))) + v71)) + ((((v71 + v80) + v80) + "hello") + ((v80 + v80) + v71)))
                                                                                                                                _ = v109
                                                                                                                                _ = v110
                                                                                                                            } else {
                                                                                                                                if (v80 <= (("" + ((v71 + "") + (("\n\t" + (v80 + v80)) + v80))) + ("a" + "foo bar"))) {
                                                                                                                                    if len(v5) > 0 {
                                                                                                                                        v5[0] = v88
                                                                                                                                    }
                                                                                                                                    v82 = 766
                                                                                                                                } else {
                                                                                                                                    var v111 int = 42
                                                                                                                                    var v112 int8 = 0
                                                                                                                                    _ = v111
                                                                                                                                    _ = v112
                                                                                                                                }
                                                                                                                                if ((v88 != 42) || (!true)) {
                                                                                                                                    v82 -= (v82 % ((825) | 1))
                                                                                                                                    v2 = (v2 >> ((1) & 63))
                                                                                                                                } else {
                                                                                                                                    var v113 float32 = ((v91 + 3.14159) * -2.25)
                                                                                                                                    if false {
                                                                                                                                        if false {
                                                                                                                                            var v114 int8 = 1
                                                                                                                                            _ = v114
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        var v115 int = (fi57 - 2147483647)
                                                                                                                                        if (!(116 != 127)) {
                                                                                                                                            switch (fi59 ^ v3) {
                                                                                                                                            case 0:
                                                                                                                                                var v116 float64 = 1.5
                                                                                                                                                _ = v116
                                                                                                                                            case 9223372036854775807:
                                                                                                                                                fmt.Println(v91)
                                                                                                                                            case -2147483648:
                                                                                                                                                fmt.Println(v113)
                                                                                                                                            default:
                                                                                                                                                if (1 < 1) {
                                                                                                                                                    var v117 float32 = 1e10
                                                                                                                                                    v118 := GenS3{F0: 9223372036854775807, F1: (v115 > v115), F2: v80}
                                                                                                                                                    _ = v117
                                                                                                                                                    _ = v118
                                                                                                                                                } else {
                                                                                                                                                    v113 = (v113 - -2.25)
                                                                                                                                                    fmt.Println(fi57, fi57, v3)
                                                                                                                                                    switch v71 {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        var v119 string = ("hello" + "a")
                                                                                                                                                        _ = v119
                                                                                                                                                    default:
                                                                                                                                                        fmt.Println(v91)
                                                                                                                                                        var v120 uint = 42
                                                                                                                                                        _ = v120
                                                                                                                                                    }
                                                                                                                                                    var v121 int16 = -873
                                                                                                                                                    if (!(!false)) {
                                                                                                                                                        var v122 bool = (0 <= 171)
                                                                                                                                                        var v123 int16 = v121
                                                                                                                                                        fmt.Println(v82)
                                                                                                                                                        _ = v122
                                                                                                                                                        _ = v123
                                                                                                                                                    }
                                                                                                                                                    _ = v121
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            switch v3 {
                                                                                                                                            case 1:
                                                                                                                                                v113 -= float32(-2.25)
                                                                                                                                            default:
                                                                                                                                                if (true && false) {
                                                                                                                                                    fmt.Println(v3, v71, v113)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            if false {
                                                                                                                                                var v124 uint8 = uint8(v88)
                                                                                                                                                v125 := v1[(fi59 | int(v82))]
                                                                                                                                                var v126 int64 = -889
                                                                                                                                                v3 = -808
                                                                                                                                                _ = v124
                                                                                                                                                _ = v125
                                                                                                                                                _ = v126
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        fmt.Println(v115, v3)
                                                                                                                                        _ = v115
                                                                                                                                    }
                                                                                                                                    if true {
                                                                                                                                        v98 += v98
                                                                                                                                        v127 := GenS11{F0: 2, F1: 1}
                                                                                                                                        _ = v127
                                                                                                                                    } else {
                                                                                                                                        if (false && false) {
                                                                                                                                            var v128 bool = (((3.14159 + 0.0) - (0.0 + v113)) > v113)
                                                                                                                                            _ = v128
                                                                                                                                        } else {
                                                                                                                                            if (((float32(0.0) / 1.0) + v91) >= (0.1 - 1e10)) {
                                                                                                                                                if false {
                                                                                                                                                    var v129 uint16 = v2
                                                                                                                                                    var v130 int64 = 0
                                                                                                                                                    if (42 < 825) {
                                                                                                                                                        var v131 uint8 = 1
                                                                                                                                                        _ = v131
                                                                                                                                                    } else {
                                                                                                                                                        fmt.Println(fi59)
                                                                                                                                                    }
                                                                                                                                                    _ = v129
                                                                                                                                                    _ = v130
                                                                                                                                                } else {
                                                                                                                                                    if false {
                                                                                                                                                        v91 = v113
                                                                                                                                                        var v132 int64 = -1
                                                                                                                                                        _ = v132
                                                                                                                                                    }
                                                                                                                                                    v113 -= v91
                                                                                                                                                }
                                                                                                                                                if (false && true) {
                                                                                                                                                    switch fi59 {
                                                                                                                                                    case 129:
                                                                                                                                                        if false {
                                                                                                                                                            fmt.Println(v88, v82, v2)
                                                                                                                                                            v133 := GenS13{F0: -957, F1: 450, F2: -697}
                                                                                                                                                            v134 := v4[(fi57 / (((v3 | fi57)) | 1))]
                                                                                                                                                            _ = v133
                                                                                                                                                            _ = v134
                                                                                                                                                        } else {
                                                                                                                                                            fmt.Println(v113, v88, fi59)
                                                                                                                                                            switch (fi57 / ((fi59) | 1)) {
                                                                                                                                                            case 9223372036854775807:
                                                                                                                                                                v2 += 1
                                                                                                                                                            case -34:
                                                                                                                                                                v91 = -2.25
                                                                                                                                                                switch fi59 {
                                                                                                                                                                case 4294967295:
                                                                                                                                                                    fmt.Println(v71)
                                                                                                                                                                }
                                                                                                                                                            case -418:
                                                                                                                                                                v135 := GenS3{F0: -1, F1: ((-2.25 == v113) || false), F2: ((v80 + v80) + "hello")}
                                                                                                                                                                switch (fi57 | 1) {
                                                                                                                                                                case -426:
                                                                                                                                                                    _ = len(v4)
                                                                                                                                                                default:
                                                                                                                                                                    if (v71 <= ((((v80 + v71) + ((("hello" + v80) + (v71 + "")) + (v80 + v71))) + (v80 + (v80 + v80))) + v71)) {
                                                                                                                                                                        v82 += 286
                                                                                                                                                                    } else {
                                                                                                                                                                        switch v3 {
                                                                                                                                                                        case 1:
                                                                                                                                                                            var v136 uint64 = 4294967295
                                                                                                                                                                            _ = v136
                                                                                                                                                                        case 850:
                                                                                                                                                                            var v137 uint32 = 251
                                                                                                                                                                            _ = v137
                                                                                                                                                                        case 0:
                                                                                                                                                                            v4[(fi57 | -2147483648)] = 0
                                                                                                                                                                        }
                                                                                                                                                                        if (0 <= 265) {
                                                                                                                                                                            _ = len(v4)
                                                                                                                                                                            v1[(-(fi59))] = v80
                                                                                                                                                                            switch v3 {
                                                                                                                                                                            case -1:
                                                                                                                                                                                var v138 uint64 = 4294967295
                                                                                                                                                                                if (!false) {
                                                                                                                                                                                    var v139 uint32 = v98
                                                                                                                                                                                    var v140 float64 = 0.0
                                                                                                                                                                                    _ = v139
                                                                                                                                                                                    _ = v140
                                                                                                                                                                                } else {
                                                                                                                                                                                    v3 = 32767
                                                                                                                                                                                    v82 = 196
                                                                                                                                                                                }
                                                                                                                                                                                v138 = 18446744073709551615
                                                                                                                                                                                _ = v138
                                                                                                                                                                            case -878:
                                                                                                                                                                                fmt.Println(fi57, v98)
                                                                                                                                                                            case 0:
                                                                                                                                                                                v141 := GenS14{F0: -947, F1: "\n\t", F2: 931}
                                                                                                                                                                                _ = v141
                                                                                                                                                                            default:
                                                                                                                                                                                v142 := GenS15{F0: 990, F1: false, F2: 467}
                                                                                                                                                                                v80 = ("hello" + "hello")
                                                                                                                                                                                _ = v142
                                                                                                                                                                            }
                                                                                                                                                                            fmt.Println(v80)
                                                                                                                                                                            if (true && (!true)) {
                                                                                                                                                                                if (0 != 65535) {
                                                                                                                                                                                    var v143 uint8 = 128
                                                                                                                                                                                    fmt.Println(fi57, v80, fi59)
                                                                                                                                                                                    _ = v143
                                                                                                                                                                                } else {
                                                                                                                                                                                    fmt.Println(v2)
                                                                                                                                                                                    fmt.Println(v98)
                                                                                                                                                                                }
                                                                                                                                                                                _ = len(v5)
                                                                                                                                                                                fmt.Println(v82, v2)
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                        var v144 bool = false
                                                                                                                                                                        var v145 uint64 = 936
                                                                                                                                                                        if v144 {
                                                                                                                                                                            fmt.Println(v3, v88, v3)
                                                                                                                                                                            switch (-(fi57)) {
                                                                                                                                                                            case 128:
                                                                                                                                                                                v2 += v2
                                                                                                                                                                            }
                                                                                                                                                                            var v146 uint32 = (v88 + v88)
                                                                                                                                                                            _ = v146
                                                                                                                                                                        } else {
                                                                                                                                                                            if (1e10 == -2.25) {
                                                                                                                                                                                _ = len(v4)
                                                                                                                                                                                switch fi57 {
                                                                                                                                                                                case 32767:
                                                                                                                                                                                    var v147 int8 = 1
                                                                                                                                                                                    _ = v147
                                                                                                                                                                                case -9223372036854775808:
                                                                                                                                                                                    var v148 string = ("" + (v80 + "hello"))
                                                                                                                                                                                    _ = v148
                                                                                                                                                                                case 4294967295:
                                                                                                                                                                                    v82 -= v98
                                                                                                                                                                                default:
                                                                                                                                                                                    var v149 float64 = 0.1
                                                                                                                                                                                    _ = v149
                                                                                                                                                                                }
                                                                                                                                                                                v3 = -9223372036854775808
                                                                                                                                                                            } else {
                                                                                                                                                                                var v150 float64 = ((0.1 * 1e10) + (1e10 / (1e10 / 3.14159)))
                                                                                                                                                                                if v144 {
                                                                                                                                                                                    var v151 float64 = v150
                                                                                                                                                                                    if (!(42 >= (^(v2)))) {
                                                                                                                                                                                        var v152 int8 = 1
                                                                                                                                                                                        fmt.Println(v151)
                                                                                                                                                                                        _ = v152
                                                                                                                                                                                    }
                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                        v5[0] = 1
                                                                                                                                                                                    }
                                                                                                                                                                                    v153 := v4[v3]
                                                                                                                                                                                    if v144 {
                                                                                                                                                                                        v5 = append(v5, 2147483648)
                                                                                                                                                                                        var v154 int32 = 1
                                                                                                                                                                                        v88 = 271
                                                                                                                                                                                        var v155 int8 = 1
                                                                                                                                                                                        var v156 uint8 = 207
                                                                                                                                                                                        if (!true) {
                                                                                                                                                                                            v80 = "hello"
                                                                                                                                                                                            v157 := GenS7{F0: 2147483647, F1: int(v154), F2: v144}
                                                                                                                                                                                            _ = v157
                                                                                                                                                                                        }
                                                                                                                                                                                        var v158 uint64 = 0
                                                                                                                                                                                        if false {
                                                                                                                                                                                            _ = len(v5)
                                                                                                                                                                                            v159 := GenS13{F0: 1, F1: (v158 << ((582) & 63)), F2: 1}
                                                                                                                                                                                            fmt.Println(v154, v145)
                                                                                                                                                                                            _ = v159
                                                                                                                                                                                        }
                                                                                                                                                                                        fmt.Println(v2)
                                                                                                                                                                                        _ = v154
                                                                                                                                                                                        _ = v155
                                                                                                                                                                                        _ = v156
                                                                                                                                                                                        _ = v158
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v160 := GenS1{F0: -116}
                                                                                                                                                                                        _ = v160
                                                                                                                                                                                    }
                                                                                                                                                                                    _ = v151
                                                                                                                                                                                    _ = v153
                                                                                                                                                                                } else {
                                                                                                                                                                                    v150 = (float64(((1.5 - ((1e10 / (float32((((v150 / v150) - 1e10) * ((v150 / 1e10) + v150))) - v91)) - v91)) / v113)) - (1e10 + 1.5))
                                                                                                                                                                                    switch v3 {
                                                                                                                                                                                    case 802:
                                                                                                                                                                                        v3 = (fi57 >> ((849) & 63))
                                                                                                                                                                                    case 0:
                                                                                                                                                                                        switch (v80 + (v80 + v71)) {
                                                                                                                                                                                        case "foo bar":
                                                                                                                                                                                            v161 := GenS16{F0: 127, F1: 238, F2: 378}
                                                                                                                                                                                            var v162 int16 = 1
                                                                                                                                                                                            _ = v161
                                                                                                                                                                                            _ = v162
                                                                                                                                                                                        default:
                                                                                                                                                                                            fmt.Println(v144, v88, v71)
                                                                                                                                                                                        }
                                                                                                                                                                                    default:
                                                                                                                                                                                        fmt.Println(v150, fi59)
                                                                                                                                                                                    }
                                                                                                                                                                                    v163 := GenS9{F0: (v80 + (v71 + (("" + v80) + (v71 + (("hello" + v80) + v71))))), F1: ((v3 / ((fi57) | 1)) < fi57)}
                                                                                                                                                                                    _ = v163
                                                                                                                                                                                }
                                                                                                                                                                                _ = v150
                                                                                                                                                                            }
                                                                                                                                                                            if v144 {
                                                                                                                                                                                v164 := GenS16{F0: 0, F1: 0, F2: 255}
                                                                                                                                                                                _ = v164
                                                                                                                                                                            }
                                                                                                                                                                            fmt.Println(v88, v91, v2)
                                                                                                                                                                            fmt.Println(fi59, v2)
                                                                                                                                                                        }
                                                                                                                                                                        if v144 {
                                                                                                                                                                            var v165 int16 = -32768
                                                                                                                                                                            v82 = (^(v82))
                                                                                                                                                                            _ = v165
                                                                                                                                                                        } else {
                                                                                                                                                                            _ = len(v4)
                                                                                                                                                                        }
                                                                                                                                                                        v88 = v98
                                                                                                                                                                        v80 = v71
                                                                                                                                                                        var v166 string = v71
                                                                                                                                                                        var v167 uint = 670
                                                                                                                                                                        v144 = false
                                                                                                                                                                        _ = v144
                                                                                                                                                                        _ = v145
                                                                                                                                                                        _ = v166
                                                                                                                                                                        _ = v167
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                                v71 = "\n\t"
                                                                                                                                                                _ = v135
                                                                                                                                                            default:
                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            var v168 float64 = float64(v91)
                                                                                                                                                            _ = v168
                                                                                                                                                        }
                                                                                                                                                    case -1:
                                                                                                                                                        switch (^(v3)) {
                                                                                                                                                        case -9223372036854775808:
                                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                                            }
                                                                                                                                                        case 160:
                                                                                                                                                            fmt.Println(v82, v91)
                                                                                                                                                        }
                                                                                                                                                    case 2147483647:
                                                                                                                                                        var v169 uint8 = 2
                                                                                                                                                        _ = v169
                                                                                                                                                    }
                                                                                                                                                    v170 := v1[-9223372036854775808]
                                                                                                                                                    _ = v170
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        if (true && (!false)) {
                                                                                                                                            v171 := GenS16{F0: 122, F1: 1, F2: 855}
                                                                                                                                            _ = v171
                                                                                                                                        } else {
                                                                                                                                            var v172 int32 = 705
                                                                                                                                            v173 := GenS3{F0: -1, F1: ((705 <= 57) || (42 > -1)), F2: "hello"}
                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                            }
                                                                                                                                            _ = v172
                                                                                                                                            _ = v173
                                                                                                                                        }
                                                                                                                                        v71 = (("\n\t" + (("a" + "") + v71)) + "foo bar")
                                                                                                                                    }
                                                                                                                                    _ = v113
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v98
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v94
                                                                                                            }
                                                                                                            _ = v88
                                                                                                            _ = v91
                                                                                                            _ = v92
                                                                                                            _ = v93
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if false {
                                                                                                            var v174 int = (-(v3))
                                                                                                            _ = v174
                                                                                                        } else {
                                                                                                            v3 = -557
                                                                                                        }
                                                                                                        v82 += v82
                                                                                                    }
                                                                                                    var v175 uint = 42
                                                                                                    if (false && false) {
                                                                                                        fmt.Println(v175, v2)
                                                                                                        v5 = append(v5, 2)
                                                                                                        if true {
                                                                                                            v80 = "a"
                                                                                                            fmt.Println(v80)
                                                                                                            if true {
                                                                                                                if true {
                                                                                                                    if (-2.25 < 1e10) {
                                                                                                                        if true {
                                                                                                                            var v176 uint16 = (v2 + (^(v2)))
                                                                                                                            _ = len(v1)
                                                                                                                            _ = v176
                                                                                                                        } else {
                                                                                                                            if ((-32768 != 32767) && (false || (0.1 == 0.0))) {
                                                                                                                                v177 := GenS15{F0: 686, F1: true, F2: 1}
                                                                                                                                _ = v177
                                                                                                                            } else {
                                                                                                                                v178 := GenS12{F0: (v71 + v71)}
                                                                                                                                if true {
                                                                                                                                    fmt.Println(v82, v82, v82)
                                                                                                                                    v175 = (v175 | v175)
                                                                                                                                    fmt.Println(v80)
                                                                                                                                    switch (v3 - v3) {
                                                                                                                                    case 4294967295:
                                                                                                                                        var v179 int8 = 0
                                                                                                                                        _ = v179
                                                                                                                                    case -9223372036854775808:
                                                                                                                                        var v180 int = -8
                                                                                                                                        _ = v180
                                                                                                                                    default:
                                                                                                                                        var v181 int64 = 42
                                                                                                                                        _ = v181
                                                                                                                                    }
                                                                                                                                    var v182 bool = true
                                                                                                                                    if v182 {
                                                                                                                                        if v182 {
                                                                                                                                            fmt.Println(fi57, v82, fi57)
                                                                                                                                            if v182 {
                                                                                                                                                fmt.Println(v175, v175)
                                                                                                                                                v183 := GenS13{F0: -1, F1: 1, F2: -9223372036854775808}
                                                                                                                                                _ = v183
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    v184 := GenS17{F0: 2147483647}
                                                                                                                                    var v185 int32 = -661
                                                                                                                                    _ = v182
                                                                                                                                    _ = v184
                                                                                                                                    _ = v185
                                                                                                                                } else {
                                                                                                                                    var v186 uint64 = 219
                                                                                                                                    _ = v186
                                                                                                                                }
                                                                                                                                _ = v178
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v2 += 1
                                                                                                                        if true {
                                                                                                                            v187 := GenS5{F0: -176, F1: -1}
                                                                                                                            _ = len(v4)
                                                                                                                            v188 := GenS3{F0: 42, F1: false, F2: v80}
                                                                                                                            _ = v187
                                                                                                                            _ = v188
                                                                                                                        } else {
                                                                                                                            v82 = 950
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v189 uint8 = 42
                                                                                                                        _ = v189
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    fmt.Println(fi57)
                                                                                                                }
                                                                                                                switch (^(fi57)) {
                                                                                                                case 41:
                                                                                                                    v82 = v82
                                                                                                                case 42:
                                                                                                                    if true {
                                                                                                                        switch v71 {
                                                                                                                        case "hello":
                                                                                                                            fmt.Println(v82, fi57)
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v190 string = v80
                                                                                                                        v175 = 579
                                                                                                                        _ = v190
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v191 := GenS12{F0: v80}
                                                                                                                    v192 := GenS18{F0: 2, F1: v80}
                                                                                                                    _ = v191
                                                                                                                    _ = v192
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v82)
                                                                                                        switch (fi59 << ((0) & 63)) {
                                                                                                        case 0:
                                                                                                            v3 -= v3
                                                                                                        case -179:
                                                                                                            var v193 uint8 = 0
                                                                                                            _ = v193
                                                                                                        }
                                                                                                    }
                                                                                                    switch (v3 & (v3 & (fi59 | fi57))) {
                                                                                                    case -1:
                                                                                                        switch v80 {
                                                                                                        case "a":
                                                                                                            v194 := GenS19{F0: true}
                                                                                                            _ = v194
                                                                                                        }
                                                                                                        var v195 uint8 = 127
                                                                                                        _ = v195
                                                                                                    case 42:
                                                                                                        switch 4294967295 {
                                                                                                        case -2147483648:
                                                                                                            if (!true) {
                                                                                                                fmt.Println(fi59, v71, fi59)
                                                                                                            } else {
                                                                                                                if true {
                                                                                                                    fmt.Println(v175, v80, fi59)
                                                                                                                    v5 = append(v5, v82)
                                                                                                                    var v196 float64 = 0.0
                                                                                                                    _ = v196
                                                                                                                } else {
                                                                                                                    var v197 uint64 = 1
                                                                                                                    var v198 uint16 = v2
                                                                                                                    v82 = (^(v82))
                                                                                                                    _ = v197
                                                                                                                    _ = v198
                                                                                                                }
                                                                                                                if (v80 == (v80 + (v80 + "foo bar"))) {
                                                                                                                    if (v2 == (v2 * (v2 << (((v175 - 4294967295)) & 63)))) {
                                                                                                                        switch v71 {
                                                                                                                        case "\n\t":
                                                                                                                            fmt.Println(v2, fi59)
                                                                                                                        case "foo bar":
                                                                                                                            var v199 int = 32767
                                                                                                                            _ = v199
                                                                                                                        default:
                                                                                                                            if true {
                                                                                                                                var v200 float32 = float32((1e10 * 1e10))
                                                                                                                                v2 = (^(v2))
                                                                                                                                _ = len(v5)
                                                                                                                                _ = v200
                                                                                                                            } else {
                                                                                                                                if ((float32(-2.25) + (0.0 - -2.25)) >= -2.25) {
                                                                                                                                    var v201 int16 = -1
                                                                                                                                    _ = v201
                                                                                                                                } else {
                                                                                                                                    v202 := GenS4{F0: true, F1: 2147483647}
                                                                                                                                    var v203 string = v80
                                                                                                                                    v204 := GenS20{F0: true, F1: 1}
                                                                                                                                    _ = v202
                                                                                                                                    _ = v203
                                                                                                                                    _ = v204
                                                                                                                                }
                                                                                                                                v175 = v175
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v82 -= v82
                                                                                                                        v80 = v80
                                                                                                                    }
                                                                                                                    v3 = -394
                                                                                                                } else {
                                                                                                                    var v205 int8 = 0
                                                                                                                    v175 = v175
                                                                                                                    _ = v205
                                                                                                                }
                                                                                                                switch v71 {
                                                                                                                case "hello":
                                                                                                                    var v206 int16 = 0
                                                                                                                    _ = v206
                                                                                                                case "":
                                                                                                                    v207 := GenS16{F0: 0, F1: 42, F2: v175}
                                                                                                                    var v208 int16 = -1
                                                                                                                    var v209 float64 = (float64(0.1) * 3.14159)
                                                                                                                    v2 = v2
                                                                                                                    _ = v207
                                                                                                                    _ = v208
                                                                                                                    _ = v209
                                                                                                                case "foo bar":
                                                                                                                    var v210 float32 = (-2.25 / 1.0)
                                                                                                                    _ = v210
                                                                                                                default:
                                                                                                                    v5 = append(v5, (v82 >> ((v175) & 63)))
                                                                                                                }
                                                                                                            }
                                                                                                        case 951:
                                                                                                            var v211 uint32 = uint32(v175)
                                                                                                            _ = v211
                                                                                                        }
                                                                                                    default:
                                                                                                        var v212 uint32 = (v82 | 1)
                                                                                                        _ = v212
                                                                                                    }
                                                                                                    if (false && true) {
                                                                                                        v2 = v2
                                                                                                        var v213 int16 = 1
                                                                                                        _ = v213
                                                                                                    } else {
                                                                                                        v5 = append(v5, 169)
                                                                                                        if (false && false) {
                                                                                                            var v214 float32 = ((1e10 / 1.0) + ((1.5 - (1.5 * 1.5)) + 0.0))
                                                                                                            if true {
                                                                                                                var v215 uint64 = 0
                                                                                                                fmt.Println(v2, v82, fi57)
                                                                                                                _ = v215
                                                                                                            } else {
                                                                                                                v216 := GenS2{F0: 0}
                                                                                                                switch fi57 {
                                                                                                                case 9223372036854775807:
                                                                                                                    var v217 int8 = 0
                                                                                                                    _ = v217
                                                                                                                case -970:
                                                                                                                    var v218 int32 = 0
                                                                                                                    _ = v218
                                                                                                                case 0:
                                                                                                                    fmt.Println(v80, fi59, v2)
                                                                                                                }
                                                                                                                _ = v216
                                                                                                            }
                                                                                                            _ = v214
                                                                                                        } else {
                                                                                                            v219 := GenS1{F0: 904}
                                                                                                            var v220 uint16 = 42
                                                                                                            var v221 uint16 = v2
                                                                                                            _ = v219
                                                                                                            _ = v220
                                                                                                            _ = v221
                                                                                                        }
                                                                                                        v71 = ((("\n\t" + ((v71 + v71) + ("foo bar" + "\n\t"))) + ((v80 + v71) + (v71 + "foo bar"))) + (((("foo bar" + ("\n\t" + (v71 + "hello"))) + v80) + (v80 + v71)) + v71))
                                                                                                        if false {
                                                                                                            var v222 uint32 = (v82 >> ((v175) & 63))
                                                                                                            switch (v80 + v80) {
                                                                                                            case "hello":
                                                                                                                switch ((v80 + (v80 + ((v80 + (v71 + ("\n\t" + "\n\t"))) + "\n\t"))) + "hello") {
                                                                                                                case "\n\t":
                                                                                                                    v82 -= (v222 + 852)
                                                                                                                case "foo bar":
                                                                                                                    if ((float64(0.0) / 1e10) == (3.14159 + 1.5)) {
                                                                                                                        if true {
                                                                                                                            if false {
                                                                                                                                v175 = uint(v82)
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                v3 = (v3 >> ((v175) & 63))
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            fmt.Println(v71, v175, v222)
                                                                                                                            if ((!(328 >= 42)) && (false || (1e10 <= 0.1))) {
                                                                                                                                fmt.Println(fi57)
                                                                                                                                v175 += (v175 & v175)
                                                                                                                                fmt.Println(v3, v175)
                                                                                                                            } else {
                                                                                                                                fmt.Println(v80, v82, v3)
                                                                                                                                if false {
                                                                                                                                    _ = len(v5)
                                                                                                                                } else {
                                                                                                                                    if ((v2 == v2) || true) {
                                                                                                                                        v223 := GenS21{F0: ((v2 >> ((4294967295) & 63)) != 0), F1: 4294967295, F2: (^(v2))}
                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                        }
                                                                                                                                        fmt.Println(v222, v175)
                                                                                                                                        _ = v223
                                                                                                                                    } else {
                                                                                                                                        v71 = v80
                                                                                                                                        v4[fi59] = (v222 & (v222 - v82))
                                                                                                                                        var v224 uint8 = uint8(v175)
                                                                                                                                        fmt.Println(v222, v2, v2)
                                                                                                                                        if false {
                                                                                                                                            var v225 uint16 = (^(v2))
                                                                                                                                            v225 = v225
                                                                                                                                            _ = v225
                                                                                                                                        } else {
                                                                                                                                            var v226 float64 = float64((1e10 * ((0.0 / 1.5) / 3.14159)))
                                                                                                                                            v226 += (float64(0.1) - v226)
                                                                                                                                            _ = v226
                                                                                                                                        }
                                                                                                                                        v227 := GenS22{F0: false, F1: uint8(v175)}
                                                                                                                                        v228 := v1[(fi57 / ((v3) | 1))]
                                                                                                                                        _ = v224
                                                                                                                                        _ = v227
                                                                                                                                        _ = v228
                                                                                                                                    }
                                                                                                                                    v2 = 1
                                                                                                                                    v82 = (v82 ^ uint32(v2))
                                                                                                                                }
                                                                                                                                v82 = (v82 + 2147483648)
                                                                                                                                switch fi59 {
                                                                                                                                case 4294967295:
                                                                                                                                    v3 -= (fi59 >> ((v175) & 63))
                                                                                                                                case -1:
                                                                                                                                    v229 := v1[127]
                                                                                                                                    _ = v229
                                                                                                                                }
                                                                                                                                v2 -= v2
                                                                                                                            }
                                                                                                                            fmt.Println(fi59, v175, v175)
                                                                                                                        }
                                                                                                                        v230 := GenS2{F0: 127}
                                                                                                                        _ = v230
                                                                                                                    } else {
                                                                                                                        var v231 uint = 536
                                                                                                                        _ = v231
                                                                                                                    }
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v232 := GenS23{F0: false, F1: -564, F2: 42}
                                                                                                                _ = v232
                                                                                                            default:
                                                                                                                if (!(2 > 4294967295)) {
                                                                                                                    switch v80 {
                                                                                                                    case "\n\t":
                                                                                                                        v222 += (v82 | (v82 - (v222 | v82)))
                                                                                                                    case "a":
                                                                                                                        if (!(true && false)) {
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if (!("" >= (((("hello" + v71) + "a") + "hello") + ""))) {
                                                                                                                                if (!((fi57 >> (((^(v175))) & 63)) != (^(fi57)))) {
                                                                                                                                    v233 := GenS24{F0: (^(v222)), F1: -9223372036854775808, F2: (v175 + v175)}
                                                                                                                                    _ = v233
                                                                                                                                } else {
                                                                                                                                    v82 = (^(v222))
                                                                                                                                    var v234 uint = v175
                                                                                                                                    if false {
                                                                                                                                        if len(v5) > 0 {
                                                                                                                                            v5[0] = (v222 | (v82 + v222))
                                                                                                                                        }
                                                                                                                                        v2 += v2
                                                                                                                                        var v235 int16 = -731
                                                                                                                                        _ = v235
                                                                                                                                    } else {
                                                                                                                                        if (0 != 42) {
                                                                                                                                            var v236 int64 = 278
                                                                                                                                            _ = v236
                                                                                                                                        } else {
                                                                                                                                            v234 -= 6
                                                                                                                                        }
                                                                                                                                        fmt.Println(v82, v222)
                                                                                                                                    }
                                                                                                                                    _ = v234
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        fmt.Println(v2, v71, v71)
                                                                                                                    case "":
                                                                                                                        v237 := v1[fi59]
                                                                                                                        _ = v237
                                                                                                                    default:
                                                                                                                        v82 -= (v222 + (^(v222)))
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v238 := GenS5{F0: 0, F1: 1000}
                                                                                                                    _ = len(v4)
                                                                                                                    if false {
                                                                                                                        var v239 string = v71
                                                                                                                        v80 = (v80 + ((v71 + "foo bar") + v239))
                                                                                                                        v240 := GenS14{F0: 42, F1: ((v239 + ("a" + ((v80 + "") + ((v239 + v239) + (v239 + v71))))) + (v71 + "hello")), F2: (v2 + v2)}
                                                                                                                        _ = v239
                                                                                                                        _ = v240
                                                                                                                    }
                                                                                                                    _ = v238
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v222
                                                                                                        } else {
                                                                                                            v175 = 42
                                                                                                            var v241 int = fi59
                                                                                                            _ = v241
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v175
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v82
                                                                                } else {
                                                                                    v242 := GenS14{F0: -928, F1: "foo bar", F2: 255}
                                                                                    var v243 int64 = 0
                                                                                    fmt.Println(v243, v2, v3)
                                                                                    _ = v242
                                                                                    _ = v243
                                                                                }
                                                                            }
                                                                            fmt.Println(v2)
                                                                            v80 = v71
                                                                        } else {
                                                                            var v244 float64 = 1.5
                                                                            v245 := GenS25{F0: int8(v2)}
                                                                            fmt.Println(fi57)
                                                                            _ = v244
                                                                            _ = v245
                                                                        }
                                                                        v2 = v2
                                                                        var v246 uint16 = (v2 << ((0) & 63))
                                                                        _ = v246
                                                                    }
                                                                    if true {
                                                                        switch v71 {
                                                                        case "a":
                                                                            v247 := GenS3{F0: 0, F1: false, F2: ((v80 + v71) + "a")}
                                                                            _ = v247
                                                                        case "foo bar":
                                                                            v80 = (v71 + "")
                                                                        case "hello":
                                                                            var v248 int = fi57
                                                                            _ = v248
                                                                        default:
                                                                            v3 = (fi57 & fi59)
                                                                        }
                                                                        if true {
                                                                            if (false || (-1 > 1)) {
                                                                                if false {
                                                                                    if false {
                                                                                        if (false || ((1 != 2) && (true || false))) {
                                                                                            v80 = (("a" + "a") + v80)
                                                                                        } else {
                                                                                            switch 0 {
                                                                                            case 4294967295:
                                                                                                if false {
                                                                                                    fmt.Println(fi59, v80, v71)
                                                                                                    v71 = (((((v80 + ((v71 + v71) + "a")) + (v71 + (v71 + ""))) + "hello") + v71) + v71)
                                                                                                } else {
                                                                                                    var v249 float64 = 0.0
                                                                                                    var v250 uint16 = (v2 / ((1) | 1))
                                                                                                    fmt.Println(v71, fi59, v3)
                                                                                                    v5 = append(v5, 65535)
                                                                                                    _ = v249
                                                                                                    _ = v250
                                                                                                }
                                                                                            case 42:
                                                                                                switch "foo bar" {
                                                                                                case "":
                                                                                                    if (!(true && true)) {
                                                                                                        var v251 string = v80
                                                                                                        v252 := GenS2{F0: 154}
                                                                                                        _ = v251
                                                                                                        _ = v252
                                                                                                    }
                                                                                                default:
                                                                                                    var v253 float32 = float32(((0.0 * (-2.25 + (0.0 * (0.1 * 0.0)))) * 1e10))
                                                                                                    _ = v253
                                                                                                }
                                                                                            case 902:
                                                                                                switch (v80 + "") {
                                                                                                case "foo bar":
                                                                                                    v254 := GenS26{F0: v3, F1: (^(fi57))}
                                                                                                    _ = v254
                                                                                                case "hello":
                                                                                                    v71 = ""
                                                                                                case "a":
                                                                                                    var v255 float32 = float32(((1.5 / 1.0) + 3.14159))
                                                                                                    _ = v255
                                                                                                default:
                                                                                                    v256 := GenS27{F0: 0, F1: (1 < 42), F2: 42}
                                                                                                    _ = v256
                                                                                                }
                                                                                                var v257 int64 = -9223372036854775808
                                                                                                v71 = "a"
                                                                                                _ = v257
                                                                                            default:
                                                                                                if (!true) {
                                                                                                    var v258 uint32 = 65535
                                                                                                    _ = v258
                                                                                                } else {
                                                                                                    v3 = v3
                                                                                                }
                                                                                            }
                                                                                            if (-2147483648 != -142) {
                                                                                                v2 = v2
                                                                                            }
                                                                                        }
                                                                                        v80 = ("a" + v71)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                v80 = ""
                                                                            }
                                                                            if false {
                                                                                v3 = 798
                                                                                if true {
                                                                                    v5 = append(v5, 2147483647)
                                                                                } else {
                                                                                    fmt.Println(fi59, v3, fi57)
                                                                                    v71 = v71
                                                                                }
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = 211
                                                                            }
                                                                        }
                                                                        var v259 int32 = 42
                                                                        _ = v259
                                                                    } else {
                                                                        var v260 uint64 = 2
                                                                        if false {
                                                                            switch "\n\t" {
                                                                            case "hello":
                                                                                v3 = fi59
                                                                            case "foo bar":
                                                                                v71 = (("foo bar" + v71) + v80)
                                                                            default:
                                                                                var v261 int64 = -1
                                                                                _ = v261
                                                                            }
                                                                            v262 := GenS24{F0: 2147483648, F1: 108, F2: 388}
                                                                            _ = v262
                                                                        }
                                                                        var v263 uint8 = 2
                                                                        _ = v260
                                                                        _ = v263
                                                                    }
                                                                    _ = v80
                                                                } else {
                                                                    var v264 int = v3
                                                                    _ = v264
                                                                }
                                                                v3 = (fi59 ^ fi57)
                                                                fmt.Println(fi57, v71, v2)
                                                            }
                                                        case "":
                                                            v2 = (^(v2))
                                                        default:
                                                            var v265 uint = 18446744073709551615
                                                            _ = v265
                                                        }
                                                        v2 -= v2
                                                    }
                                                    fmt.Println(fi59)
                                                }
                                            }
                                            fmt.Println(v3, fi57)
                                            var v266 int8 = 1
                                            _ = v266
                                        } else {
                                            v5 = append(v5, 2147483647)
                                        }
                                    case 1:
                                        var v267 uint = 18446744073709551615
                                        _ = v267
                                    case 128:
                                        if true {
                                            v2 += v2
                                        } else {
                                            var v268 int8 = -1
                                            _ = v268
                                        }
                                    default:
                                        v269 := GenS9{F0: ("" + v71), F1: (((true || (v71 == "foo bar")) || ("a" >= "")) || false)}
                                        _ = v269
                                    }
                                    if true {
                                        if (true || true) {
                                            v2 = 1
                                        }
                                    } else {
                                        v71 = v71
                                    }
                                }
                                if true {
                                    v5 = append(v5, 1)
                                } else {
                                    v270 := GenS6{F0: 1}
                                    var v271 int16 = genFn0(32767)
                                    _ = v270
                                    _ = v271
                                }
                                _ = v71
                            }
                        default:
                            if (v2 == (^(v2))) {
                                fmt.Println(v3, v3)
                                v272 := GenS28{F0: 2147483647, F1: 0, F2: 1}
                                switch "foo bar" {
                                case "hello":
                                    v2 -= v2
                                case "":
                                    var v273 int32 = -1
                                    _ = v273
                                case "foo bar":
                                    var v274 uint64 = 42
                                    _ = v274
                                default:
                                    v2 -= (v2 ^ 1)
                                }
                                _ = v272
                            }
                        }
                    }
                    v3 = v3
                    var v275 int = (fi57 - (fi59 << ((65535) & 63)))
                    _ = fi59
                    _ = v275
                }
                _ = fi57
                _ = v58
            }
        default:
            switch "a" {
            case "a":
                for fi276 := 0; fi276 < 14; fi276++ {
                    var v277 float64 = -2.25
                    _ = fi276
                    _ = v277
                }
            default:
                v2 = v2
            }
        }
        if ((v2 * v2) != v2) {
            if (false || (48 > 1)) {
                if true {
                    v3 -= (-(v3))
                    for fi278 := 0; fi278 < 20; fi278++ {
                        v279 := GenS2{F0: 127}
                        _ = fi278
                        _ = v279
                    }
                } else {
                    v280 := GenS28{F0: 288, F1: 0, F2: 1}
                    _ = v280
                }
            } else {
                v281 := GenS29{F0: 47, F1: 1}
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                }
                _ = v281
            }
            var v282 bool = (!false)
            _ = v282
        }
        for fi283 := 0; fi283 < 9; fi283++ {
            fmt.Println(v2)
            if (0.0 <= (-2.25 + 0.1)) {
                if true {
                    v2 -= (v2 >> ((uint(v2)) & 63))
                    for fi284 := 0; fi284 < 20; fi284++ {
                        switch (v3 << ((42) & 63)) {
                        case -2147483648:
                            v2 = (v2 - v2)
                        default:
                            var v285 int = (fi284 / ((int(v2)) | 1))
                            _ = v285
                        }
                        fmt.Println(v3, v2)
                        var v286 bool = true
                        v3 = (fi284 % (((v3 | (fi283 & 42))) | 1))
                        var v287 uint = 648
                        var v288 int16 = -32768
                        var v289 float64 = (1.5 * 3.14159)
                        if (v2 == 42) {
                            if (true && v286) {
                                v290 := GenS14{F0: 414, F1: ("hello" + "hello"), F2: v2}
                                if false {
                                    v286 = (v286 && (!(42 >= 42)))
                                    v289 -= (((v289 / ((v289 + (float64(float32(1.5)) - (v289 + ((((1.5 + (v289 * (0.1 + v289))) * 3.14159) * 3.14159) - -2.25)))) + (v289 / v289))) - v289) * 0.0)
                                } else {
                                    fmt.Println(v2)
                                    v291 := v1[(v3 << (((v287 + v287)) & 63))]
                                    fmt.Println(v288)
                                    _ = v291
                                }
                                var v292 float64 = 0.0
                                _ = v290
                                _ = v292
                            }
                        } else {
                            v286 = (-2.25 < (0.1 / 0.1))
                        }
                        if (4294967295 >= 42) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v293 := GenS30{F0: (194 == 124)}
                            var v294 float64 = v289
                            fmt.Println(v289, v287)
                            v289 = float64((float32(float64(1e10)) / ((-2.25 * (3.14159 * 0.1)) / 0.1)))
                            _ = v293
                            _ = v294
                        } else {
                            fmt.Println(v286, v3)
                            var v295 float32 = (1e10 / (((0.0 - 3.14159) - 3.14159) - -2.25))
                            var v296 uint64 = 18446744073709551615
                            _ = v295
                            _ = v296
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v297 := GenS31{F0: 65535}
                        _ = fi284
                        _ = v286
                        _ = v287
                        _ = v288
                        _ = v289
                        _ = v297
                    }
                    v2 = 65535
                } else {
                    if (469 == 18446744073709551615) {
                        _ = len(v4)
                    }
                    var v298 uint = 436
                    if len(v5) > 0 {
                        v5[0] = 465
                    }
                    _ = v298
                }
                if false {
                    v299 := GenS13{F0: 838, F1: 0, F2: 498}
                    _ = v299
                }
            }
            _ = fi283
        }
    }
    var v300 int = v3
    var v301 uint64 = 668
    var v302 float32 = ((1.5 - 0.0) * 1.5)
    v303 := GenS32{F0: v301, F1: (v3 + (v3 ^ v3))}
    if len(v5) > 0 {
        v5[0] = 349
    }
    v302 += v302
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v300:", v300)
    fmt.Println("v301:", v301)
    fmt.Println("v302:", v302)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v300
    _ = v301
    _ = v302
    _ = v303
    _ = v4
    _ = v5
    _ = v8
}
