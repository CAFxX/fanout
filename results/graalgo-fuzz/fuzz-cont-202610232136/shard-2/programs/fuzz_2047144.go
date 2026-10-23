package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint; F2 uint }
type GenS5 struct { F0 bool; F1 int32 }
type GenS6 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int; F1 uint16; F2 uint32 }
type GenS11 struct { F0 int16; F1 string }
type GenS12 struct { F0 bool; F1 string; F2 string }
type GenS13 struct { F0 int16; F1 uint8 }
type GenS14 struct { F0 string; F1 uint16; F2 string }
type GenS15 struct { F0 int; F1 uint8 }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int }
type GenS19 struct { F0 bool; F1 int; F2 uint16 }
type GenS20 struct { F0 uint16; F1 uint32; F2 uint }
type GenS21 struct { F0 int8; F1 uint32 }
type GenS22 struct { F0 uint8; F1 uint }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 string }
func genFn0(p0 uint16, p1 bool) uint {
    v7 := GenS1{F0: 817}
    p0 -= 727
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 string) uint16 {
    v8 := GenS2{F0: -1, F1: 580}
    for fi9 := 0; fi9 < 3; fi9++ {
        switch p0 {
        case "hello":
            var v10 bool = true
            _ = v10
        default:
            var v11 float64 = 1e10
            _ = v11
        }
        if (-1 < -128) {
            p0 = p0
            p0 = p0
        }
        var v12 int = fi9
        _ = fi9
        _ = v12
    }
    var v13 int16 = 32767
    var v14 int32 = 1
    var v15 uint16 = 1
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    _ = v8
    return v15
}

func main() {
    var v1 float64 = 3.14159
    v2 := []string{}
    var v3 int8 = -1
    var v4 int32 = 313
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]bool)
    switch ("a" + "foo bar") {
    case "hello":
        fmt.Println(v3, v1)
    default:
        if ((-9223372036854775808 < -652) || (((v4 % ((2147483647) | 1)) < 0) && false)) {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 float32 = float32((float64((float32(v1) * (0.0 - 0.0))) - v1))
                var v18 uint = 255
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            fmt.Println(v1)
            switch ("a" + "\n\t") {
            case "foo bar":
                var v19 uint32 = 357
                _ = v19
            case "":
                var v20 int8 = (v3 * -128)
                _ = v20
            default:
                v21 := GenS1{F0: 2}
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint = 18446744073709551615
        v24 := GenS2{F0: -9223372036854775808, F1: 2}
        _ = len(v6)
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int16 = 938
    if (false && (!false)) {
        if false {
            var v26 float32 = (1.5 * 0.0)
            _ = v26
        } else {
            fmt.Println(v3, v3, v3)
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 uint8 = 5
                v29 := GenS2{F0: -9223372036854775808, F1: 157}
                _ = fi27
                _ = v28
                _ = v29
            }
            var v30 uint32 = 42
            _ = v30
        }
        for fi31 := 0; fi31 < 20; fi31++ {
            var v32 uint8 = 0
            if (2 < 214) {
                for fi33 := 0; fi33 < 10; fi33++ {
                    v2 = append(v2, "\n\t")
                    v25 += (v25 >> ((42) & 63))
                    switch (fi33 - fi31) {
                    case -1:
                        v34 := GenS3{F0: 4294967295}
                        _ = v34
                    case 0:
                        v4 += (^(v4))
                    default:
                        var v35 uint64 = 95
                        _ = v35
                    }
                    if true {
                        v6[143] = ((9223372036854775807 == 42) || false)
                        v1 = v1
                        v36 := GenS2{F0: 301, F1: 255}
                        _ = v36
                    }
                    switch (fi33 << ((genFn0(255, true)) & 63)) {
                    case -9223372036854775808:
                        fmt.Println(v4, v4, v4)
                    case -851:
                        if ((^(fi31)) > fi31) {
                            if false {
                                var v37 uint32 = 42
                                switch "foo bar" {
                                case "a":
                                    var v38 float64 = (v1 / 3.14159)
                                    _ = v38
                                case "":
                                    var v39 int16 = v25
                                    _ = v39
                                case "\n\t":
                                    fmt.Println(v3, v37)
                                    v40 := GenS3{F0: (fi33 & fi33)}
                                    if false {
                                        switch fi33 {
                                        case 1:
                                            fmt.Println(v4, v3, v25)
                                        case 128:
                                            v41 := GenS4{F0: v4, F1: 4294967295, F2: 18446744073709551615}
                                            _ = v41
                                        default:
                                            var v42 uint16 = 65535
                                            _ = v42
                                        }
                                        switch -843 {
                                        case 4294967295:
                                            v2 = append(v2, ("" + "hello"))
                                        case -75:
                                            if len(v2) > 0 {
                                                v2[0] = ((("" + "\n\t") + ("foo bar" + "")) + "a")
                                            }
                                        case 127:
                                            var v43 int8 = -87
                                            _ = v43
                                        }
                                        fmt.Println(v32)
                                        v25 = v25
                                    } else {
                                        v3 = 1
                                        if len(v2) > 0 {
                                            v2[0] = ("a" + ("" + "a"))
                                        }
                                    }
                                    _ = v40
                                }
                                _ = v37
                            } else {
                                if (768 == 1) {
                                    v44 := GenS3{F0: (fi31 | 32767)}
                                    if ((-2.25 / float64(0.1)) < v1) {
                                        if false {
                                            v4 -= (v4 | v4)
                                            switch (fi31 / ((42) | 1)) {
                                            case -732:
                                                v3 += 100
                                            case 99:
                                                v45 := v6[2147483647]
                                                switch 1 {
                                                case -612:
                                                    fmt.Println(fi31, v3, v1)
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v45
                                            case 4294967295:
                                                v1 = v1
                                            }
                                            v46 := GenS5{F0: (v1 < 1e10), F1: (v4 * int32(v3))}
                                            _ = v46
                                        } else {
                                            delete(v6, 960)
                                            v3 = v3
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        if true {
                                            v47 := GenS5{F0: (!((v32 << ((316) & 63)) > 198)), F1: 1}
                                            _ = v47
                                        }
                                    } else {
                                        v25 = (^(v25))
                                    }
                                    _ = v44
                                } else {
                                    v4 -= (v4 / ((v4) | 1))
                                    v48 := GenS6{F0: 0, F1: v3, F2: 0}
                                    v3 -= (v3 / ((v3) | 1))
                                    _ = v48
                                }
                                delete(v5, 865)
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            }
                            v25 -= v25
                        }
                    default:
                        fmt.Println(v25)
                    }
                    _ = fi33
                }
                var v49 uint32 = 4
                _ = v49
            }
            _ = fi31
            _ = v32
        }
    }
    var v50 int32 = v4
    switch ("\n\t" + "hello") {
    case "foo bar":
        var v51 string = ("\n\t" + "foo bar")
        _ = v51
    case "a":
        var v52 bool = true
        _ = v52
    case "":
        if true {
            if false {
                if false {
                    fmt.Println(v25)
                    var v53 uint32 = 0
                    if true {
                        var v54 int64 = 42
                        _ = v54
                    } else {
                        for fi55 := 0; fi55 < 18; fi55++ {
                            if false {
                                switch ("hello" + "a") {
                                case "":
                                    for fi56 := 0; fi56 < 18; fi56++ {
                                        var v57 uint64 = 2
                                        var v58 uint64 = (^(v57))
                                        fmt.Println(v58, v50)
                                        _ = fi56
                                        _ = v57
                                        _ = v58
                                    }
                                case "hello":
                                    var v59 string = "a"
                                    _ = v59
                                }
                                var v60 uint8 = 1
                                if true {
                                    v3 = 0
                                    for fi61 := 0; fi61 < 8; fi61++ {
                                        var v62 int8 = v3
                                        v1 = v1
                                        switch fi61 {
                                        case -193:
                                            v60 = 0
                                            if false {
                                                if ((1e10 + float32(v1)) == (1e10 + 0.1)) {
                                                    _ = len(v5)
                                                } else {
                                                    var v63 uint64 = 953
                                                    if (!true) {
                                                        var v64 int = (fi61 - 128)
                                                        v65 := GenS4{F0: (v4 & 2147483647), F1: 203, F2: 42}
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    _ = len(v5)
                                                    _ = v63
                                                }
                                                switch 42 {
                                                case 4294967295:
                                                    _ = len(v2)
                                                default:
                                                    v66 := GenS7{F0: 255}
                                                    _ = v66
                                                }
                                            } else {
                                                v50 += -978
                                                v50 -= v4
                                            }
                                            if (v1 >= v1) {
                                                if (!true) {
                                                    if false {
                                                        if true {
                                                            var v67 uint16 = uint16(v53)
                                                            _ = v67
                                                        }
                                                        if true {
                                                            v25 = 330
                                                            v3 += v3
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v68 := GenS6{F0: 724, F1: v3, F2: 4294967295}
                                                            var v69 int64 = 0
                                                            switch "hello" {
                                                            case "hello":
                                                                var v70 int16 = (v25 >> ((3) & 63))
                                                                _ = v70
                                                            case "":
                                                                var v71 string = ""
                                                                _ = v71
                                                            case "\n\t":
                                                                v72 := v5[0]
                                                                _ = v72
                                                            default:
                                                                v50 = v4
                                                            }
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    }
                                                    fmt.Println(v4, v25, fi61)
                                                    if true {
                                                        var v73 int16 = (-(v25))
                                                        fmt.Println(v25, v53, v73)
                                                        _ = v73
                                                    } else {
                                                        switch -96 {
                                                        case 1:
                                                            v50 -= (v50 & (v4 - v50))
                                                        case 32768:
                                                            switch (("" + ("\n\t" + "foo bar")) + "foo bar") {
                                                            case "a":
                                                                var v74 int8 = v62
                                                                _ = v74
                                                            case "hello":
                                                                var v75 uint64 = 42
                                                                _ = v75
                                                            default:
                                                                if false {
                                                                    v3 -= (-(v62))
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v76 := GenS3{F0: (fi55 - (fi61 * (-(fi61))))}
                                                                        _ = v76
                                                                    default:
                                                                        v53 += (v53 & (v53 * 2147483648))
                                                                        var v77 uint64 = 2
                                                                        _ = v77
                                                                    }
                                                                    delete(v6, (v53 >> ((1) & 63)))
                                                                    fmt.Println(v1, v1, v4)
                                                                }
                                                            }
                                                        case -549:
                                                            v4 = (v4 & (^(v50)))
                                                        default:
                                                            if false {
                                                                _ = len(v5)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v78 bool = false
                                                    v1 = v1
                                                    v3 = -128
                                                    _ = v78
                                                }
                                            }
                                            _ = len(v2)
                                        case -1:
                                            fmt.Println(v62, v3)
                                        }
                                        _ = fi61
                                        _ = v62
                                    }
                                    var v79 bool = false
                                    _ = v79
                                } else {
                                    var v80 int16 = v25
                                    var v81 int16 = (v25 << ((uint(v60)) & 63))
                                    for fi82 := 0; fi82 < 18; fi82++ {
                                        _ = len(v6)
                                        var v83 int32 = (v4 >> ((4294967295) & 63))
                                        _ = len(v5)
                                        _ = fi82
                                        _ = v83
                                    }
                                    _ = v80
                                    _ = v81
                                }
                                _ = v60
                            }
                            v84 := GenS8{F0: 903, F1: (("foo bar" + "hello") + "hello")}
                            v25 += v25
                            _ = fi55
                            _ = v84
                        }
                    }
                    _ = v53
                }
            }
            if false {
                if ((1 == v25) || (3.14159 >= v1)) {
                    v3 -= 103
                }
                v85 := GenS9{F0: 32767, F1: false}
                for fi86 := 0; fi86 < 14; fi86++ {
                    v87 := v5[568]
                    _ = fi86
                    _ = v87
                }
                _ = v85
            } else {
                v1 = -2.25
                for fi88 := 0; fi88 < 12; fi88++ {
                    for fi89 := 0; fi89 < 19; fi89++ {
                        var v90 uint64 = 0
                        var v91 int32 = (v50 << ((18446744073709551615) & 63))
                        fmt.Println(v3)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    for fi92 := 0; fi92 < 7; fi92++ {
                        if false {
                            fmt.Println(fi92, fi88, v4)
                        } else {
                            v93 := GenS5{F0: false, F1: (v50 << ((18446744073709551615) & 63))}
                            _ = v93
                        }
                        _ = fi92
                    }
                    var v94 bool = (true || (true && (0 <= v3)))
                    _ = fi88
                    _ = v94
                }
                var v95 int8 = 42
                _ = v95
            }
            v1 -= (float64(1e10) - -2.25)
        }
    default:
        for fi96 := 0; fi96 < 10; fi96++ {
            if (-128 == (v3 << ((4294967295) & 63))) {
                for fi97 := 0; fi97 < 14; fi97++ {
                    v98 := GenS8{F0: 18446744073709551615, F1: "a"}
                    v3 = 127
                    v4 += 22
                    _ = fi97
                    _ = v98
                }
                v50 -= (^(v4))
            } else {
                for fi99 := 0; fi99 < 9; fi99++ {
                    if (!(3.14159 < float32(v1))) {
                        if false {
                            if false {
                                switch "\n\t" {
                                case "a":
                                    v1 += (v1 * (3.14159 * (1e10 - (((1e10 * (1e10 / 1e10)) + 3.14159) + v1))))
                                case "":
                                    if false {
                                        v2 = append(v2, ("a" + "\n\t"))
                                        fmt.Println(v1, v4, v3)
                                    }
                                }
                                switch "\n\t" {
                                case "":
                                    v1 += v1
                                case "hello":
                                    if (-643 <= -9223372036854775808) {
                                        switch "" {
                                        case "foo bar":
                                            switch 96 {
                                            case 128:
                                                v4 = v50
                                            case 42:
                                                var v100 int = (fi99 | (fi96 | fi99))
                                                _ = v100
                                            }
                                        case "hello":
                                            v3 = (v3 / ((v3) | 1))
                                        }
                                        if ((fi99 << ((4294967295) & 63)) != (fi99 & fi99)) {
                                            v50 = -2147483648
                                            v50 = 2147483647
                                            fmt.Println(v1, v50)
                                        } else {
                                            if false {
                                                switch (fi96 % (((fi96 - -9223372036854775808)) | 1)) {
                                                case 965:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 32768:
                                                    var v101 uint32 = 2147483647
                                                    _ = v101
                                                case 42:
                                                    var v102 uint64 = 2
                                                    _ = v102
                                                default:
                                                    var v103 uint32 = 262
                                                    _ = v103
                                                }
                                            }
                                            var v104 int = -1
                                            v105 := v5[1]
                                            _ = v104
                                            _ = v105
                                        }
                                        if true {
                                            switch "foo bar" {
                                            case "foo bar":
                                                if (127 < 100) {
                                                    switch (^(fi96)) {
                                                    case 17:
                                                        v106 := GenS10{F0: (fi99 * fi99), F1: 127, F2: 42}
                                                        _ = v106
                                                    default:
                                                        delete(v6, 448)
                                                    }
                                                } else {
                                                    v1 -= v1
                                                    fmt.Println(v4, v3)
                                                }
                                            case "a":
                                                if false {
                                                    var v107 uint32 = 1
                                                    switch "hello" {
                                                    case "a":
                                                        if (!true) {
                                                            if false {
                                                                var v108 float32 = ((0.0 / (((3.14159 + 0.0) * 0.1) - ((float32(v1) + ((1e10 / 1.0) + 1.5)) * 3.14159))) / (1e10 + (1.5 / (0.1 * 1.5))))
                                                                if true {
                                                                    if false {
                                                                        if (!(true && ((42 > v3) || false))) {
                                                                            _ = len(v2)
                                                                            var v109 uint64 = 819
                                                                            _ = v109
                                                                        } else {
                                                                            _ = len(v5)
                                                                            v4 += v50
                                                                        }
                                                                    } else {
                                                                        if (true || ((v107 ^ v107) != v107)) {
                                                                            v50 -= v50
                                                                            v107 = 595
                                                                            switch "a" {
                                                                            case "hello":
                                                                                v110 := GenS11{F0: v25, F1: ("a" + "\n\t")}
                                                                                _ = v110
                                                                            }
                                                                        }
                                                                        if true {
                                                                            if false {
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    var v111 uint64 = 4294967295
                                                                                    _ = v111
                                                                                }
                                                                                var v112 int = (fi96 - (fi99 / ((fi96) | 1)))
                                                                                if ((v4 + (v50 >> ((488) & 63))) <= v50) {
                                                                                    v1 += v1
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    var v113 int32 = -2147483648
                                                                                    _ = v113
                                                                                } else {
                                                                                    v114 := GenS12{F0: ((v1 + 0.1) == ((v1 * v1) + 0.1)), F1: ("\n\t" + "\n\t"), F2: ("hello" + "hello")}
                                                                                    _ = v114
                                                                                }
                                                                                _ = v112
                                                                            } else {
                                                                                if ((!false) || true) {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = ("hello" + "\n\t")
                                                                                    }
                                                                                }
                                                                                v1 = (v1 - 1.5)
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = "\n\t"
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v115 int32 = (v50 ^ int32(v107))
                                                                            _ = v115
                                                                        }
                                                                    }
                                                                    v4 = v4
                                                                    switch (fi96 ^ fi96) {
                                                                    case 4294967295:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            if false {
                                                                                var v116 bool = (true || (((float32(v1) - 1e10) / v108) > (((v108 * v108) * 0.0) + (v108 + v108))))
                                                                                v116 = false
                                                                                v25 += (^(v25))
                                                                                _ = v116
                                                                            }
                                                                        case "hello":
                                                                            v4 -= v50
                                                                        }
                                                                    case -2147483648:
                                                                        if false {
                                                                            v25 = v25
                                                                            if true {
                                                                                var v117 uint8 = 199
                                                                                v118 := GenS13{F0: (v25 / (((v25 & 0)) | 1)), F1: v117}
                                                                                _ = v117
                                                                                _ = v118
                                                                            }
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            switch fi96 {
                                                                            case 42:
                                                                                if true {
                                                                                    v25 = -193
                                                                                }
                                                                            case 32767:
                                                                                var v119 float32 = v108
                                                                                _ = v119
                                                                            case 1:
                                                                                var v120 float32 = 1.5
                                                                                _ = v120
                                                                            }
                                                                            if true {
                                                                                if ((v3 | (-(v3))) != (-(v3))) {
                                                                                    switch (fi96 << ((2) & 63)) {
                                                                                    case -119:
                                                                                        var v121 uint64 = 921
                                                                                        _ = v121
                                                                                    case 32767:
                                                                                        _ = len(v6)
                                                                                    case -518:
                                                                                        v2 = append(v2, "")
                                                                                    default:
                                                                                        switch fi96 {
                                                                                        case 32767:
                                                                                            var v122 int16 = v25
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v3 -= (v3 - v3)
                                                                                            _ = v122
                                                                                        }
                                                                                    }
                                                                                    switch fi96 {
                                                                                    case 182:
                                                                                        if (true || true) {
                                                                                            v123 := GenS14{F0: "a", F1: 32767, F2: "foo bar"}
                                                                                            v3 = (v3 * (v3 - (v3 >> ((genFn0(32767, (-1 > (v25 + v25)))) & 63))))
                                                                                            v124 := GenS3{F0: (fi96 ^ fi99)}
                                                                                            fmt.Println(v25, v50)
                                                                                            _ = v123
                                                                                            _ = v124
                                                                                        } else {
                                                                                            var v125 bool = true
                                                                                            var v126 int16 = v25
                                                                                            v3 -= (v3 % (((v3 % (((v3 >> ((uint(v107)) & 63))) | 1))) | 1))
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                        }
                                                                                    case 0:
                                                                                        var v127 uint16 = 298
                                                                                        _ = v127
                                                                                    case 9223372036854775807:
                                                                                        delete(v5, v107)
                                                                                    }
                                                                                    v128 := GenS9{F0: genFn1(("\n\t" + "a")), F1: true}
                                                                                    _ = v128
                                                                                } else {
                                                                                    var v129 uint8 = 147
                                                                                    _ = v129
                                                                                }
                                                                                fmt.Println(v107)
                                                                                switch fi96 {
                                                                                case 32768:
                                                                                    v108 += (1e10 * (v108 * v108))
                                                                                default:
                                                                                    var v130 int = 32767
                                                                                    _ = v130
                                                                                }
                                                                            } else {
                                                                                var v131 int64 = 1
                                                                                fmt.Println(v131, v3, fi96)
                                                                                _ = v131
                                                                            }
                                                                            v5[(^(v107))] = "a"
                                                                            v50 = v50
                                                                        } else {
                                                                            switch "" {
                                                                            case "hello":
                                                                                fmt.Println(v4, fi99, v1)
                                                                            case "a":
                                                                                v132 := GenS15{F0: int(v25), F1: 92}
                                                                                _ = v132
                                                                            }
                                                                            if ((!false) && false) {
                                                                                switch (fi96 >> ((genFn0(65535, false)) & 63)) {
                                                                                case 32768:
                                                                                    v133 := GenS11{F0: v25, F1: "foo bar"}
                                                                                    _ = v133
                                                                                case 9223372036854775807:
                                                                                    if false {
                                                                                        if (v107 == (v107 | v107)) {
                                                                                            if (false && false) {
                                                                                                v134 := GenS7{F0: 389}
                                                                                                _ = v134
                                                                                            }
                                                                                        } else {
                                                                                            v5[(^(v107))] = "hello"
                                                                                            _ = len(v5)
                                                                                            v5[428] = ""
                                                                                        }
                                                                                        if true {
                                                                                            v135 := GenS16{F0: false, F1: 77}
                                                                                            v136 := GenS17{F0: fi96}
                                                                                            _ = v135
                                                                                            _ = v136
                                                                                        }
                                                                                        v1 -= -2.25
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if true {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    var v137 float32 = v108
                                                                                    _ = v137
                                                                                }
                                                                                v1 += 3.14159
                                                                                if ((v25 <= (v25 * v25)) && true) {
                                                                                    v107 += v107
                                                                                    v108 = v108
                                                                                    v25 = v25
                                                                                } else {
                                                                                    if (217 < 0) {
                                                                                        var v138 uint = 255
                                                                                        _ = v138
                                                                                    } else {
                                                                                        v108 -= v108
                                                                                    }
                                                                                    v139 := GenS18{F0: (-(fi99))}
                                                                                    _ = v139
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v140 := v6[(^(v107))]
                                                                        switch fi99 {
                                                                        case -705:
                                                                            fmt.Println(v1, fi96, v108)
                                                                        case 32767:
                                                                            v108 = 3.14159
                                                                        case 766:
                                                                            if true {
                                                                                switch ("hello" + "\n\t") {
                                                                                case "foo bar":
                                                                                    v141 := GenS15{F0: (fi96 & 0), F1: 128}
                                                                                    _ = v141
                                                                                case "a":
                                                                                    v2 = append(v2, "\n\t")
                                                                                case "hello":
                                                                                    v1 -= (-2.25 * (1.5 * v1))
                                                                                }
                                                                                v25 -= (v25 / (((v25 - -175)) | 1))
                                                                                var v142 int8 = v3
                                                                                var v143 int16 = -32768
                                                                                _ = v142
                                                                                _ = v143
                                                                            }
                                                                        default:
                                                                            var v144 int32 = (v4 | 698)
                                                                            _ = v144
                                                                        }
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "hello"
                                                                        }
                                                                        _ = v140
                                                                    }
                                                                    v145 := GenS13{F0: v25, F1: 2}
                                                                    v25 = (v25 >> ((0) & 63))
                                                                    v146 := GenS19{F0: false, F1: (-(fi96)), F2: genFn1("\n\t")}
                                                                    _ = v145
                                                                    _ = v146
                                                                }
                                                                _ = v108
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "\n\t"
                                                                }
                                                                v25 -= (^(v25))
                                                            }
                                                            switch (fi96 ^ 0) {
                                                            case 1:
                                                                var v147 float32 = 3.14159
                                                                _ = v147
                                                            case -9223372036854775808:
                                                                v148 := GenS16{F0: (334 < 0), F1: fi99}
                                                                _ = v148
                                                            case 9223372036854775807:
                                                                fmt.Println(fi99, v3, v107)
                                                            }
                                                        } else {
                                                            v25 = v25
                                                            switch (fi99 | fi99) {
                                                            case -23:
                                                                v149 := GenS1{F0: 0}
                                                                _ = v149
                                                            case 524:
                                                                v4 += v4
                                                            case -1:
                                                                v50 -= v50
                                                            default:
                                                                v150 := GenS12{F0: (v3 > 1), F1: "a", F2: ("\n\t" + "")}
                                                                _ = v150
                                                            }
                                                        }
                                                    }
                                                    if (!(!((int64(v107) <= 0) && true))) {
                                                        v107 += v107
                                                        if (false || ((false || (255 <= 490)) || (!(671 > (fi96 << ((770) & 63)))))) {
                                                            var v151 uint = genFn0(1, (!((fi96 - fi99) <= (-(fi96)))))
                                                            switch (fi96 * int(v151)) {
                                                            case -286:
                                                                var v152 uint8 = 1
                                                                _ = v152
                                                            case 2147483647:
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    var v153 float64 = ((v1 + 1.5) / v1)
                                                                    _ = v153
                                                                case "\n\t":
                                                                    v25 += (v25 | int16(v50))
                                                                case "":
                                                                    var v154 string = "hello"
                                                                    _ = v154
                                                                default:
                                                                    var v155 uint16 = 255
                                                                    _ = v155
                                                                }
                                                            case 32767:
                                                                var v156 uint = (v151 >> (((^(v151))) & 63))
                                                                _ = v156
                                                            }
                                                            v107 += (^(v107))
                                                            _ = v151
                                                        } else {
                                                            v157 := GenS17{F0: fi99}
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if true {
                                                                    var v158 uint32 = (^(v107))
                                                                    _ = len(v6)
                                                                    fmt.Println(fi99)
                                                                    _ = v158
                                                                } else {
                                                                    v1 += (-2.25 - (0.0 - (float64((1e10 / 1.0)) * ((float64(3.14159) * (0.1 - v1)) - 0.0))))
                                                                }
                                                            case "\n\t":
                                                                v4 -= int32(fi99)
                                                            case "":
                                                                if false {
                                                                    var v159 int8 = v3
                                                                    switch -530 {
                                                                    case -605:
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "hello"
                                                                        }
                                                                    case 42:
                                                                        if true {
                                                                            fmt.Println(v50, fi96)
                                                                            var v160 float32 = (0.0 + 3.14159)
                                                                            var v161 float32 = (0.1 * v160)
                                                                            v160 = (v160 / v161)
                                                                            var v162 string = ("foo bar" + "hello")
                                                                            _ = v160
                                                                            _ = v161
                                                                            _ = v162
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if (false && (!((true && false) || false))) {
                                                                            var v163 int = (fi96 << ((794) & 63))
                                                                            fmt.Println(v159, v1, v3)
                                                                            v25 = (^(v25))
                                                                            var v164 uint64 = 1
                                                                            _ = len(v5)
                                                                            _ = v163
                                                                            _ = v164
                                                                        }
                                                                    default:
                                                                        switch "" {
                                                                        case "hello":
                                                                            switch (^(fi96)) {
                                                                            case 128:
                                                                                v1 -= v1
                                                                            case 284:
                                                                                switch "" {
                                                                                case "a":
                                                                                    if (false || true) {
                                                                                        switch 0 {
                                                                                        case 32767:
                                                                                            v159 = v159
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            v159 += 1
                                                                                        } else {
                                                                                            _ = len(v5)
                                                                                            if false {
                                                                                                v25 = v25
                                                                                                v165 := GenS13{F0: (v25 >> ((148) & 63)), F1: 1}
                                                                                                v5[v107] = "hello"
                                                                                                _ = v165
                                                                                            } else {
                                                                                                delete(v6, (^(v107)))
                                                                                                var v166 uint8 = 121
                                                                                                _ = v166
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    v25 -= (v25 + (v25 >> ((288) & 63)))
                                                                                }
                                                                            default:
                                                                                v1 += (v1 - (v1 * v1))
                                                                            }
                                                                        case "\n\t":
                                                                            if true {
                                                                                var v167 float64 = 3.14159
                                                                                switch fi99 {
                                                                                case 4294967295:
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        _ = len(v5)
                                                                                    case "\n\t":
                                                                                        if (!(-495 == 1)) {
                                                                                            v1 += 1e10
                                                                                            var v168 float64 = 1e10
                                                                                            _ = v168
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v169 uint16 = 42
                                                                                        _ = v169
                                                                                    }
                                                                                case 128:
                                                                                    _ = len(v6)
                                                                                }
                                                                                _ = v167
                                                                            } else {
                                                                                fmt.Println(v159)
                                                                                v170 := GenS7{F0: 781}
                                                                                _ = v170
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                v50 = (v4 - v4)
                                                                                var v171 int = 521
                                                                                _ = v171
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v159
                                                                } else {
                                                                    v172 := GenS20{F0: 639, F1: v107, F2: 0}
                                                                    v50 -= (v4 ^ (v4 & 42))
                                                                    v173 := GenS21{F0: (v3 * v3), F1: (^(v107))}
                                                                    _ = v172
                                                                    _ = v173
                                                                }
                                                            }
                                                            _ = v157
                                                        }
                                                    }
                                                    var v174 uint64 = 203
                                                    _ = v107
                                                    _ = v174
                                                } else {
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v175 int = (fi96 * 735)
                                                        _ = v175
                                                    default:
                                                        v50 -= 0
                                                    }
                                                    var v176 int32 = -2147483648
                                                    _ = v176
                                                }
                                            case "hello":
                                                _ = len(v5)
                                            }
                                            v2 = append(v2, "hello")
                                            delete(v6, 779)
                                        } else {
                                            var v177 float64 = v1
                                            var v178 float32 = (0.1 * -2.25)
                                            _ = v177
                                            _ = v178
                                        }
                                        if (65535 > 2) {
                                            v1 = ((v1 * ((v1 * v1) + v1)) + v1)
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                        } else {
                                            var v179 bool = true
                                            _ = v179
                                        }
                                    }
                                default:
                                    var v180 uint64 = 4294967295
                                    fmt.Println(fi99)
                                    if (!(!true)) {
                                        v25 = 60
                                        v181 := GenS4{F0: (v4 % ((42) | 1)), F1: genFn0(65535, false), F2: 4294967295}
                                        if true {
                                            var v182 int64 = -9223372036854775808
                                            _ = v182
                                        } else {
                                            v180 += (^(v180))
                                            var v183 uint32 = 49
                                            var v184 string = (((("foo bar" + ("\n\t" + "a")) + "") + "hello") + ("hello" + ("hello" + "foo bar")))
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = v183
                                            _ = v184
                                        }
                                        _ = v181
                                    } else {
                                        v185 := GenS22{F0: 2, F1: 46}
                                        fmt.Println(v1)
                                        _ = v185
                                    }
                                    v186 := GenS12{F0: (v180 != v180), F1: "", F2: ""}
                                    _ = v180
                                    _ = v186
                                }
                            }
                            v3 = (v3 / ((112) | 1))
                        } else {
                            var v187 int8 = -128
                            v187 += 42
                            v50 += (v4 | -1)
                            _ = v187
                        }
                        var v188 int16 = (v25 - v25)
                        var v189 uint32 = 11
                        _ = v188
                        _ = v189
                    }
                    _ = fi99
                }
                fmt.Println(v3, v1)
            }
            var v190 float32 = (-2.25 + 1e10)
            _ = fi96
            _ = v190
        }
    }
    v25 += v25
    var v191 bool = false
    if false {
        switch 32767 {
        case 128:
            v192 := GenS23{F0: 65535}
            _ = v192
        case -1:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            switch -29 {
            case -2147483648:
                var v193 uint8 = 0
                _ = v193
            default:
                switch 4294967295 {
                case 0:
                    for fi194 := 0; fi194 < 15; fi194++ {
                        switch fi194 {
                        case -551:
                            v25 = v25
                        case 128:
                            switch (fi194 + (fi194 << ((255) & 63))) {
                            case 127:
                                var v195 int16 = v25
                                _ = v195
                            case 420:
                                switch ("hello" + (("a" + ("hello" + "")) + ("" + ""))) {
                                case "":
                                    delete(v5, 165)
                                }
                            default:
                                switch 42 {
                                case -827:
                                    fmt.Println(v191, v25, v4)
                                }
                            }
                        }
                        v50 = (v50 | -2147483648)
                        v50 = 1
                        _ = fi194
                    }
                case 1:
                    switch 542 {
                    case 42:
                        v25 += int16(v3)
                    }
                }
            }
        }
    } else {
        var v196 float64 = v1
        _ = v196
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v197 uint8 = 42
    fmt.Println(v1, v197)
    v5[0] = "hello"
    if v191 {
        v198 := GenS24{F0: ("a" + ("a" + "foo bar"))}
        switch "hello" {
        case "hello":
            switch 9223372036854775807 {
            case 463:
                if (v191 || (1 == 0)) {
                    for fi199 := 0; fi199 < 10; fi199++ {
                        if v191 {
                            if (960 != -1) {
                                v197 += (v197 | (v197 - v197))
                            }
                            fmt.Println(v4, fi199, v4)
                        }
                        _ = fi199
                    }
                }
            case -2147483648:
                v191 = v191
            case -9223372036854775808:
                v200 := GenS17{F0: 1}
                _ = v200
            default:
                var v201 uint64 = 1
                _ = v201
            }
        case "":
            fmt.Println(v25, v197)
        default:
            for fi202 := 0; fi202 < 16; fi202++ {
                for fi203 := 0; fi203 < 5; fi203++ {
                    var v204 int16 = v25
                    _ = fi203
                    _ = v204
                }
                _ = fi202
            }
        }
        _ = v198
    } else {
        v205 := GenS18{F0: 1}
        if (!false) {
            v191 = (!v191)
            v206 := GenS9{F0: 128, F1: v191}
            _ = v206
        } else {
            fmt.Println(v50)
        }
        for fi207 := 0; fi207 < 11; fi207++ {
            if false {
                for fi208 := 0; fi208 < 17; fi208++ {
                    v50 += -2147483648
                    var v209 int64 = -1
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    var v210 float64 = 1e10
                    var v211 float64 = float64((1e10 - (((0.0 - float32(v210)) - (1.5 + 1.5)) / float32(v210))))
                    _ = fi208
                    _ = v209
                    _ = v210
                    _ = v211
                }
                v191 = (!v191)
            } else {
                if len(v2) > 0 {
                    v2[0] = ""
                }
                if (true || (2 > 932)) {
                    var v212 int = 32767
                    switch v212 {
                    case 9223372036854775807:
                        v212 = v212
                    }
                    _ = v212
                }
            }
            for fi213 := 0; fi213 < 16; fi213++ {
                if (2 <= 65535) {
                    fmt.Println(fi207, fi213, v25)
                    switch ("\n\t" + "\n\t") {
                    case "foo bar":
                        if true {
                            v214 := v6[1]
                            fmt.Println(v197, v191)
                            var v215 uint = 886
                            _ = v214
                            _ = v215
                        } else {
                            v216 := GenS7{F0: 1}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v6[65535] = v191
                            _ = v216
                        }
                    case "hello":
                        var v217 uint16 = 42
                        _ = v217
                    default:
                        _ = len(v6)
                    }
                    _ = len(v2)
                } else {
                    v197 = (v197 % (((v197 ^ v197)) | 1))
                    var v218 bool = (v191 && v191)
                    _ = len(v2)
                    _ = v218
                }
                var v219 uint32 = uint32(v197)
                _ = fi213
                _ = v219
            }
            _ = fi207
        }
        _ = v205
    }
    var v220 float64 = ((float64(-2.25) * v1) - v1)
    v3 += v3
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v191:", v191)
    fmt.Println("v197:", v197)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v220:", v220)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v191
    _ = v197
    _ = v2
    _ = v220
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
