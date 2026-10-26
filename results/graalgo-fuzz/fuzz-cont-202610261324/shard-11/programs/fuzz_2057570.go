package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string; F1 int8 }
type GenS13 struct { F0 int8; F1 uint }
type GenS14 struct { F0 int8; F1 int16; F2 uint16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 string; F1 int }
type GenS17 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS18 struct { F0 int8; F1 uint }
type GenS19 struct { F0 int32; F1 bool; F2 uint }
type GenS20 struct { F0 int }
type GenS21 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS22 struct { F0 uint32; F1 int16; F2 int32 }
type GenS23 struct { F0 uint; F1 string; F2 bool }
type GenS24 struct { F0 uint; F1 int8 }
func genFn0() int8 {
    var v4 uint8 = 255
    var v5 int64 = 9223372036854775807
    _ = v4
    _ = v5
    return -17
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint = 624
    var v3 uint16 = 255
    if (!(int16(v2) == -1)) {
        var v6 string = "a"
        _ = v6
    } else {
        var v7 uint32 = 42
        _ = v7
    }
    v8 := GenS1{F0: -128, F1: 42}
    if ((v2 >> (((v2 ^ (v2 & (v2 ^ (v2 >> ((65535) & 63)))))) & 63)) >= v2) {
        switch (("hello" + "hello") + "hello") {
        case "a":
            var v9 uint64 = 883
            _ = v9
        case "foo bar":
            switch "\n\t" {
            case "":
                v10 := GenS1{F0: genFn0(), F1: 174}
                _ = v10
            case "a":
                fmt.Println(v2)
            case "foo bar":
                fmt.Println(v3, v3)
            }
        default:
            var v11 int8 = genFn0()
            _ = v11
        }
    } else {
        v2 += v2
        var v12 int32 = 0
        for fi13 := 0; fi13 < 5; fi13++ {
            switch ("" + "hello") {
            case "a":
                if (!false) {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 int8 = -1
                        var v16 float32 = (((((3.14159 / ((float32(((1e10 / 1.0) - -2.25)) - 3.14159) + (-2.25 + (float32(0.0) / float32(1e10))))) * 3.14159) - (0.0 - (0.0 - (0.0 + (1.5 * 0.1))))) / (3.14159 / float32(1.5))) * 0.0)
                        v12 = (-(v12))
                        v2 -= v2
                        if false {
                            fmt.Println(v12)
                            if (!((v3 >> ((uint(v3)) & 63)) > 65535)) {
                                var v17 bool = (42 > 209)
                                _ = v17
                            }
                            _ = len(v1)
                        } else {
                            fmt.Println(fi14, v3, fi13)
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                var v18 float64 = 3.14159
                                _ = v18
                            } else {
                                fmt.Println(v12)
                                v12 = 180
                                v15 += -81
                            }
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if ((((0.1 + 3.14159) / 1.5) - (1e10 / 1e10)) == ((float64(float32(1e10)) - 1e10) * 1.5)) {
                        if false {
                            switch "foo bar" {
                            case "hello":
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                            case "foo bar":
                                var v19 int16 = 1
                                _ = v19
                            }
                            v12 = (v12 >> ((4294967295) & 63))
                        } else {
                            v1 = append(v1, 65535)
                            v20 := GenS2{F0: -80}
                            v21 := GenS3{F0: v2, F1: uint16(v2)}
                            _ = v20
                            _ = v21
                        }
                    } else {
                        for fi22 := 0; fi22 < 11; fi22++ {
                            if (false || false) {
                                if true {
                                    var v23 bool = true
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v23
                                }
                            }
                            _ = fi22
                        }
                        v3 -= (v3 ^ v3)
                    }
                } else {
                    if (fi13 < (fi13 * fi13)) {
                        var v24 uint8 = 67
                        fmt.Println(v2, fi13)
                        v1 = append(v1, 4294967295)
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            v26 := GenS4{F0: v2, F1: fi25}
                            if false {
                                var v27 uint16 = v3
                                _ = v27
                            }
                            _ = fi25
                            _ = v26
                        }
                        for fi28 := 0; fi28 < 17; fi28++ {
                            if true {
                                if true {
                                    if (6 <= (v12 << ((v2) & 63))) {
                                        v29 := GenS1{F0: 0, F1: -32768}
                                        v12 -= 42
                                        var v30 float32 = ((1e10 * -2.25) / float32(1e10))
                                        v2 = 550
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v31 int = fi13
                                        switch fi28 {
                                        case 32767:
                                            if false {
                                                v12 = v12
                                                var v32 float32 = (0.0 / ((3.14159 - 0.1) / 3.14159))
                                                v33 := GenS5{F0: 421}
                                                _ = v32
                                                _ = v33
                                            }
                                        case -1:
                                            fmt.Println(v12, v31, v3)
                                        case 128:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                        _ = v31
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 65535
                                    }
                                    if (627 > (v2 % (((v2 | v2)) | 1))) {
                                        if true {
                                            fmt.Println(v3, fi28)
                                        }
                                    }
                                }
                            } else {
                                fmt.Println(v2, v12, fi13)
                                if false {
                                    v3 = (v3 << ((v2) & 63))
                                }
                                v3 = (^(v3))
                                var v34 uint64 = 1
                                _ = v34
                            }
                            _ = fi28
                        }
                    }
                    v12 += 1
                    var v35 uint = 2
                    _ = v35
                }
            case "foo bar":
                v36 := GenS6{F0: v3, F1: 0}
                _ = v36
            }
            switch "hello" {
            case "a":
                if (0 != 73) {
                    switch (fi13 % ((fi13) | 1)) {
                    case 128:
                        for fi37 := 0; fi37 < 1; fi37++ {
                            var v38 uint32 = 638
                            v1 = append(v1, v38)
                            if ((v12 - v12) <= v12) {
                                _ = len(v1)
                                fmt.Println(v38, fi37)
                                var v39 uint16 = 1
                                switch -2147483648 {
                                case 1:
                                    _ = len(v1)
                                case 366:
                                    var v40 int16 = 32767
                                    _ = v40
                                case -2147483648:
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                default:
                                    v1 = append(v1, v38)
                                }
                                _ = v39
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v41 uint8 = 255
                                v42 := GenS7{F0: 2}
                                _ = v41
                                _ = v42
                            }
                            _ = fi37
                            _ = v38
                        }
                    case 42:
                        for fi43 := 0; fi43 < 13; fi43++ {
                            switch "a" {
                            case "a":
                                var v44 int64 = -9223372036854775808
                                _ = v44
                            default:
                                if false {
                                    switch "hello" {
                                    case "\n\t":
                                        v45 := GenS3{F0: (^(v2)), F1: 672}
                                        _ = v45
                                    case "a":
                                        v1 = append(v1, 2147483647)
                                    case "hello":
                                        fmt.Println(v3, fi43)
                                    default:
                                        v12 = -2147483648
                                    }
                                }
                            }
                            fmt.Println(fi13, fi13)
                            if true {
                                var v46 int8 = -128
                                v47 := GenS8{F0: v12}
                                if false {
                                    v48 := GenS9{F0: 1, F1: 65535}
                                    if true {
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v49 bool = false
                                            _ = v49
                                        case "":
                                            v50 := GenS5{F0: 851}
                                            _ = v50
                                        case "\n\t":
                                            var v51 int = fi13
                                            _ = v51
                                        default:
                                            v52 := GenS10{F0: -1}
                                            _ = v52
                                        }
                                    }
                                    _ = v48
                                } else {
                                    v3 = v3
                                    if ((^(v12)) == 42) {
                                        v53 := GenS9{F0: -1, F1: 65535}
                                        switch (((("foo bar" + ((("\n\t" + "foo bar") + "foo bar") + "\n\t")) + ("hello" + "hello")) + "") + "hello") {
                                        case "a":
                                            if false {
                                                v46 = v46
                                                var v54 int8 = (v46 % ((127) | 1))
                                                _ = v54
                                            }
                                        case "foo bar":
                                            v12 = 249
                                        case "":
                                            v55 := GenS3{F0: v2, F1: (v3 + (v3 + v3))}
                                            _ = v55
                                        }
                                        _ = v53
                                    }
                                    var v56 int32 = v12
                                    _ = v56
                                }
                                _ = v46
                                _ = v47
                            }
                            _ = fi43
                        }
                    }
                    for fi57 := 0; fi57 < 16; fi57++ {
                        if (-2.25 <= 3.14159) {
                            v3 -= 293
                            var v58 int8 = 100
                            switch "" {
                            case "foo bar":
                                if false {
                                    v58 += -58
                                    if ((1.5 <= 1e10) && false) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if true {
                                            v59 := GenS5{F0: 2}
                                            var v60 int64 = 214
                                            if (v58 < (v58 - -1)) {
                                                var v61 string = ""
                                                var v62 int64 = (v60 - v60)
                                                _ = v61
                                                _ = v62
                                            }
                                            _ = v59
                                            _ = v60
                                        } else {
                                            v58 = (v58 >> (((^(v2))) & 63))
                                            _ = len(v1)
                                            var v63 uint16 = (^(v3))
                                            _ = v63
                                        }
                                    } else {
                                        switch ("\n\t" + "foo bar") {
                                        case "foo bar":
                                            var v64 int32 = int32(fi57)
                                            _ = v64
                                        }
                                        fmt.Println(fi13, fi57, v3)
                                    }
                                }
                            }
                            _ = v58
                        } else {
                            var v65 uint8 = 255
                            fmt.Println(v65, v3, v2)
                            fmt.Println(v65)
                            _ = v65
                        }
                        fmt.Println(fi13, v2, v12)
                        if false {
                            var v66 bool = false
                            _ = v66
                        }
                        v2 -= (v2 & v2)
                        _ = fi57
                    }
                    if false {
                        if true {
                            if (-32768 <= -1) {
                                v67 := GenS1{F0: int8(v12), F1: -722}
                                v2 = (^(v2))
                                _ = v67
                            }
                            v12 = -2147483648
                            v2 = v2
                        }
                        if false {
                            for fi68 := 0; fi68 < 7; fi68++ {
                                v2 -= (^(v2))
                                _ = fi68
                            }
                        } else {
                            v2 += v2
                        }
                        v69 := GenS8{F0: (v12 << ((v2) & 63))}
                        _ = v69
                    } else {
                        for fi70 := 0; fi70 < 12; fi70++ {
                            _ = len(v1)
                            switch (-(fi70)) {
                            case 0:
                                if false {
                                    _ = len(v1)
                                }
                            }
                            _ = fi70
                        }
                        for fi71 := 0; fi71 < 10; fi71++ {
                            if false {
                                if (0.0 == (((-2.25 + 0.1) - 3.14159) + (((3.14159 * (-2.25 * (-2.25 / 1e10))) * ((1.5 / 1.5) * 0.0)) / 1.5))) {
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch 2147483647 {
                                        case 127:
                                            if (!((false || ((161 != 53) || false)) || false)) {
                                                if false {
                                                    var v72 uint = (v2 - (v2 % ((203) | 1)))
                                                    v73 := GenS7{F0: 549}
                                                    v2 += (v2 % ((v72) | 1))
                                                    if (!(2 == 0)) {
                                                        var v74 float64 = float64((((-2.25 - (-2.25 * float32(float64(0.0)))) + (1.5 + (((0.1 / 0.1) * (-2.25 / 1e10)) - -2.25))) / 1.5))
                                                        _ = v74
                                                    } else {
                                                        v72 = 2
                                                        fmt.Println(fi13)
                                                        fmt.Println(fi13, v2, fi13)
                                                    }
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    if false {
                                                        var v75 int = (-(fi71))
                                                        v2 += v2
                                                        _ = v75
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            _ = len(v1)
                                        }
                                    }
                                } else {
                                    if (-2.25 == 3.14159) {
                                        if ((!false) || (65535 < 437)) {
                                            v3 = 631
                                            switch (fi71 - fi13) {
                                            case 127:
                                                v76 := GenS11{F0: 1}
                                                _ = v76
                                            case 0:
                                                var v77 uint16 = 675
                                                _ = v77
                                            default:
                                                var v78 uint32 = 999
                                                _ = v78
                                            }
                                        }
                                    }
                                    var v79 string = "a"
                                    fmt.Println(fi13)
                                    _ = v79
                                }
                                v12 -= 672
                                var v80 string = "foo bar"
                                _ = v80
                            }
                            var v81 int32 = (^(v12))
                            var v82 uint8 = 1
                            _ = fi71
                            _ = v81
                            _ = v82
                        }
                        fmt.Println(v12)
                    }
                } else {
                    v83 := GenS12{F0: "", F1: 0}
                    _ = v83
                }
            case "\n\t":
                var v84 float32 = float32(1e10)
                _ = v84
            case "":
                v3 -= 32767
            }
            var v85 string = ((("foo bar" + "hello") + "foo bar") + "foo bar")
            _ = fi13
            _ = v85
        }
        _ = v12
    }
    v3 += (^(v3))
    fmt.Println(v3, v3)
    for fi86 := 0; fi86 < 7; fi86++ {
        for fi87 := 0; fi87 < 12; fi87++ {
            v3 -= v3
            switch (fi86 / ((fi86) | 1)) {
            case -2147483648:
                var v88 uint32 = 2
                _ = v88
            default:
                v3 = uint16(v2)
            }
            _ = fi87
        }
        var v89 bool = false
        if v89 {
            if (1 <= 118) {
                var v90 int64 = 42
                var v91 string = ""
                _ = v90
                _ = v91
            }
            var v92 float32 = 1.5
            switch (fi86 | 759) {
            case 761:
                v2 = v2
            }
            _ = v92
        } else {
            for fi93 := 0; fi93 < 5; fi93++ {
                if v89 {
                    v94 := GenS13{F0: genFn0(), F1: v2}
                    _ = v94
                } else {
                    fmt.Println(v89)
                    if (0 != 0) {
                        v95 := GenS2{F0: 100}
                        if v89 {
                            v96 := GenS8{F0: -2147483648}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch -537 {
                            case -214:
                                fmt.Println(v2)
                            case -2147483648:
                                var v97 string = "foo bar"
                                _ = v97
                            case 9223372036854775807:
                                switch (("a" + "a") + "foo bar") {
                                case "a":
                                    v98 := GenS1{F0: -16, F1: 42}
                                    _ = v98
                                default:
                                    var v99 float32 = 1e10
                                    _ = v99
                                }
                            }
                            _ = v96
                        } else {
                            if ((v3 + v3) != (v3 ^ 65535)) {
                                var v100 int = 4294967295
                                switch "\n\t" {
                                case "\n\t":
                                    v1 = append(v1, 2147483647)
                                case "hello":
                                    fmt.Println(v100, fi93)
                                }
                                _ = v100
                            } else {
                                v2 = (v2 + (^(v2)))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (908 == 18446744073709551615) {
                                    switch (-(fi93)) {
                                    case 524:
                                        v3 = 255
                                    case 128:
                                        var v101 int = (-(fi86))
                                        _ = v101
                                    }
                                    if (("" + "foo bar") <= "hello") {
                                        v102 := GenS6{F0: v3, F1: 4294967295}
                                        var v103 uint64 = 42
                                        if ((v103 | uint64(v2)) != v103) {
                                            switch ("foo bar" + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(fi86, v3)
                                            case "\n\t":
                                                v89 = (1 != -1)
                                            }
                                            v3 -= v3
                                        }
                                        var v104 uint = 2
                                        v104 -= (v2 | v2)
                                        _ = v102
                                        _ = v103
                                        _ = v104
                                    } else {
                                        v2 = 42
                                    }
                                    v2 = v2
                                } else {
                                    switch 369 {
                                    case 32767:
                                        if (!((("hello" + "a") != "hello") || (127 >= 1))) {
                                            var v105 uint = v2
                                            if (296 < 2147483648) {
                                                if true {
                                                    var v106 uint = (v2 ^ 281)
                                                    switch fi93 {
                                                    case 128:
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                    case -888:
                                                        v3 -= 255
                                                    default:
                                                        switch "" {
                                                        case "":
                                                            v106 += (v2 ^ v2)
                                                        case "hello":
                                                            v105 = (^(v105))
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = 819
                                                    }
                                                    _ = v106
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        if v89 {
                                                            v107 := GenS2{F0: 42}
                                                            var v108 uint = v2
                                                            _ = v107
                                                            _ = v108
                                                        } else {
                                                            v109 := GenS4{F0: uint(v3), F1: int(v2)}
                                                            _ = v109
                                                        }
                                                    case "a":
                                                        var v110 uint16 = (v3 - 0)
                                                        v110 = uint16(v2)
                                                        var v111 float32 = (-2.25 + -2.25)
                                                        _ = v110
                                                        _ = v111
                                                    case "\n\t":
                                                        _ = len(v1)
                                                    }
                                                    var v112 int64 = 1
                                                    _ = v112
                                                }
                                                fmt.Println(v105, v3)
                                                v89 = true
                                            } else {
                                                v2 += v2
                                                switch (((("a" + ((("\n\t" + "a") + ("foo bar" + "")) + ("" + ""))) + "hello") + ("\n\t" + (("a" + "hello") + "\n\t"))) + (("foo bar" + "foo bar") + "hello")) {
                                                case "foo bar":
                                                    var v113 int8 = genFn0()
                                                    _ = v113
                                                case "":
                                                    var v114 bool = (-175 > 0)
                                                    _ = v114
                                                case "a":
                                                    if (0.1 == 0.1) {
                                                        if (!v89) {
                                                            if (-2147483648 >= 0) {
                                                                var v115 uint8 = 1
                                                                v1 = append(v1, uint32(v2))
                                                                _ = len(v1)
                                                                _ = v115
                                                            } else {
                                                                switch (fi93 - fi93) {
                                                                case 438:
                                                                    if v89 {
                                                                        switch ("hello" + "\n\t") {
                                                                        case "hello":
                                                                            if v89 {
                                                                                var v116 bool = v89
                                                                                _ = v116
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(fi93)
                                                                        case "a":
                                                                            v89 = v89
                                                                        }
                                                                        v1 = append(v1, 2147483647)
                                                                    } else {
                                                                        if (false || (-102 != -1)) {
                                                                            v89 = v89
                                                                            var v117 int32 = -398
                                                                            _ = v117
                                                                        }
                                                                        v89 = v89
                                                                        var v118 bool = v89
                                                                        v3 = 1
                                                                        _ = v118
                                                                    }
                                                                case 0:
                                                                    _ = len(v1)
                                                                case 465:
                                                                    var v119 uint = 255
                                                                    _ = v119
                                                                default:
                                                                    v3 += 913
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = 555
                                                                }
                                                            }
                                                            if v89 {
                                                                v120 := GenS14{F0: -128, F1: 0, F2: 119}
                                                                _ = v120
                                                            }
                                                            var v121 float32 = 3.14159
                                                            _ = v121
                                                        } else {
                                                            var v122 uint64 = 2
                                                            if v89 {
                                                                v123 := GenS2{F0: 100}
                                                                v105 = (v105 | (v2 + (v2 % ((v105) | 1))))
                                                                var v124 float64 = 3.14159
                                                                v125 := GenS11{F0: genFn0()}
                                                                v105 += (v105 / ((v105) | 1))
                                                                _ = v123
                                                                _ = v124
                                                                _ = v125
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = uint32(v2)
                                                                }
                                                                v2 = (v2 - v2)
                                                                fmt.Println(fi93)
                                                            }
                                                            _ = v122
                                                        }
                                                        v89 = v89
                                                    } else {
                                                        v89 = v89
                                                        v2 = 65535
                                                        v126 := GenS15{F0: "foo bar"}
                                                        _ = v126
                                                    }
                                                default:
                                                    switch "" {
                                                    case "a":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if false {
                                                            var v127 uint64 = 0
                                                            _ = v127
                                                        } else {
                                                            if v89 {
                                                                var v128 bool = true
                                                                v2 -= v2
                                                                _ = v128
                                                            } else {
                                                                if (v3 == (v3 >> (((v105 % ((v105) | 1))) & 63))) {
                                                                    switch fi93 {
                                                                    case 127:
                                                                        switch ("a" + "\n\t") {
                                                                        case "foo bar":
                                                                            v3 += 2
                                                                        case "hello":
                                                                            var v129 int = fi93
                                                                            _ = v129
                                                                        default:
                                                                            var v130 uint32 = 2
                                                                            _ = v130
                                                                        }
                                                                    case -892:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 0
                                                                        }
                                                                    case 128:
                                                                        _ = len(v1)
                                                                    default:
                                                                        fmt.Println(v89, fi93, v3)
                                                                        v131 := GenS7{F0: 805}
                                                                        v105 += 18446744073709551615
                                                                        _ = v131
                                                                    }
                                                                }
                                                            }
                                                            switch (("" + (("hello" + (("" + (("" + ("\n\t" + "")) + "foo bar")) + (("a" + "") + ((("foo bar" + ((("\n\t" + "") + "a") + "hello")) + "foo bar") + "")))) + ("foo bar" + ("hello" + "a")))) + "") {
                                                            case "hello":
                                                                if false {
                                                                    v132 := GenS4{F0: 42, F1: (fi93 & fi93)}
                                                                    _ = v132
                                                                }
                                                            case "":
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case "a":
                                                                if v89 {
                                                                    _ = len(v1)
                                                                } else {
                                                                    var v133 int64 = 1
                                                                    var v134 int = (-(fi93))
                                                                    v135 := GenS13{F0: -1, F1: v105}
                                                                    var v136 float32 = (((0.0 * (0.1 - 0.1)) - (3.14159 * (float32((0.0 / 1e10)) - ((float32(0.1) - 0.1) * (0.0 - -2.25))))) + 1.5)
                                                                    _ = v133
                                                                    _ = v134
                                                                    _ = v135
                                                                    _ = v136
                                                                }
                                                            }
                                                            if v89 {
                                                                switch (fi93 ^ (fi93 | (fi93 / (((fi93 & fi86)) | 1)))) {
                                                                case 2147483647:
                                                                    var v137 float32 = (0.0 * 1e10)
                                                                    _ = v137
                                                                case 517:
                                                                    var v138 int32 = -503
                                                                    _ = v138
                                                                case 1:
                                                                    v139 := GenS16{F0: (("hello" + "") + "hello"), F1: (fi93 - (^(fi86)))}
                                                                    _ = v139
                                                                default:
                                                                    if (!v89) {
                                                                        if (!(!(v89 && true))) {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 915
                                                                            }
                                                                            if true {
                                                                                if v89 {
                                                                                    switch fi93 {
                                                                                    case -975:
                                                                                        if (-922 != 851) {
                                                                                            var v140 int = 32767
                                                                                            _ = v140
                                                                                        } else {
                                                                                            v105 += (v105 - 736)
                                                                                            var v141 int8 = genFn0()
                                                                                            var v142 bool = v89
                                                                                            _ = v141
                                                                                            _ = v142
                                                                                        }
                                                                                    case 32768:
                                                                                        v105 = v105
                                                                                    case 127:
                                                                                        fmt.Println(fi93)
                                                                                    default:
                                                                                        if v89 {
                                                                                            if false {
                                                                                                v143 := GenS7{F0: 18446744073709551615}
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    v2 -= v2
                                                                                                case "a":
                                                                                                    v89 = v89
                                                                                                case "":
                                                                                                    switch fi86 {
                                                                                                    case 1:
                                                                                                        var v144 uint64 = 4294967295
                                                                                                        _ = v144
                                                                                                    case -9223372036854775808:
                                                                                                        v3 += v3
                                                                                                    case 217:
                                                                                                        if true {
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v3)
                                                                                                _ = v143
                                                                                            } else {
                                                                                                switch (fi93 / ((fi86) | 1)) {
                                                                                                case -556:
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = 991
                                                                                                    }
                                                                                                case 525:
                                                                                                    var v145 int = (fi93 + 18)
                                                                                                    _ = v145
                                                                                                case 712:
                                                                                                    if v89 {
                                                                                                        switch (("a" + (("\n\t" + ("\n\t" + "a")) + (("" + "") + "a"))) + "a") {
                                                                                                        case "hello":
                                                                                                            if v89 {
                                                                                                                v105 -= v105
                                                                                                                var v146 uint64 = 827
                                                                                                                v89 = (v89 || v89)
                                                                                                                _ = v146
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            var v147 int64 = 9223372036854775807
                                                                                                            _ = v147
                                                                                                        case "a":
                                                                                                            var v148 uint64 = 42
                                                                                                            _ = v148
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v105 = 0
                                                                                                }
                                                                                                v149 := GenS17{F0: int16(fi93), F1: 0, F2: v3}
                                                                                                _ = v149
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v3 += v3
                                                                                }
                                                                                var v150 int16 = -649
                                                                                switch (fi93 | (fi86 & (fi86 ^ (fi86 / ((42) | 1))))) {
                                                                                case 127:
                                                                                    v3 += (v3 % ((v3) | 1))
                                                                                case 32768:
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    fmt.Println(fi93, v150, fi93)
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        switch ("\n\t" + "") {
                                                                                        case "foo bar":
                                                                                            switch "\n\t" {
                                                                                            case "\n\t":
                                                                                                v151 := GenS18{F0: int8(v3), F1: (v105 - 18446744073709551615)}
                                                                                                _ = v151
                                                                                            }
                                                                                        case "":
                                                                                            var v152 int64 = 9223372036854775807
                                                                                            _ = v152
                                                                                        case "hello":
                                                                                            if (!true) {
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 2
                                                                                        }
                                                                                    case "a":
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            v89 = true
                                                                                        case "hello":
                                                                                            fmt.Println(v105, v105)
                                                                                        }
                                                                                    }
                                                                                case 826:
                                                                                    v150 -= (-(v150))
                                                                                default:
                                                                                    v3 += v3
                                                                                }
                                                                                if (-20 != genFn0()) {
                                                                                    if (1 < 1) {
                                                                                        var v153 uint = 255
                                                                                        var v154 int32 = 0
                                                                                        var v155 int8 = 127
                                                                                        _ = v153
                                                                                        _ = v154
                                                                                        _ = v155
                                                                                    } else {
                                                                                        v89 = (v89 || (fi93 < (fi93 << ((v105) & 63))))
                                                                                        v2 = 2
                                                                                    }
                                                                                    if v89 {
                                                                                        v156 := GenS19{F0: int32(v150), F1: false, F2: (v2 % (((v2 - 42)) | 1))}
                                                                                        if v89 {
                                                                                            if v89 {
                                                                                                if v89 {
                                                                                                    v157 := GenS4{F0: v105, F1: fi93}
                                                                                                    var v158 uint32 = 2147483648
                                                                                                    _ = v157
                                                                                                    _ = v158
                                                                                                }
                                                                                            }
                                                                                            v89 = ((fi93 != (fi93 * int(v2))) || v89)
                                                                                            v105 = v2
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 704
                                                                                            }
                                                                                            var v159 int = fi93
                                                                                            var v160 uint32 = 1
                                                                                            _ = v159
                                                                                            _ = v160
                                                                                        } else {
                                                                                            switch (fi93 | 661) {
                                                                                            case 2147483647:
                                                                                                v161 := GenS17{F0: v150, F1: 201, F2: (v3 >> (((v105 << (((v105 % ((v105) | 1))) & 63))) & 63))}
                                                                                                _ = v161
                                                                                            case 469:
                                                                                                var v162 string = (((((("" + "") + "a") + ((("foo bar" + ("a" + ("" + "foo bar"))) + "a") + (("a" + "hello") + "\n\t"))) + "foo bar") + ("foo bar" + "hello")) + "\n\t")
                                                                                                _ = v162
                                                                                            default:
                                                                                                fmt.Println(fi93)
                                                                                            }
                                                                                        }
                                                                                        switch (("\n\t" + "") + "") {
                                                                                        case "\n\t":
                                                                                            var v163 uint8 = 128
                                                                                            _ = v163
                                                                                        }
                                                                                        _ = v156
                                                                                    } else {
                                                                                        fmt.Println(v2, v2)
                                                                                        var v164 string = "hello"
                                                                                        var v165 int8 = -1
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                    }
                                                                                } else {
                                                                                    v2 = v105
                                                                                }
                                                                                v166 := GenS7{F0: 4294967295}
                                                                                _ = v150
                                                                                _ = v166
                                                                            } else {
                                                                                var v167 int8 = 1
                                                                                _ = v167
                                                                            }
                                                                        }
                                                                        if (1 == 0) {
                                                                            v168 := GenS1{F0: 0, F1: 1000}
                                                                            v105 = (v2 * (v105 >> ((v2) & 63)))
                                                                            var v169 float64 = -2.25
                                                                            _ = v168
                                                                            _ = v169
                                                                        } else {
                                                                            var v170 uint16 = v3
                                                                            v171 := GenS10{F0: 950}
                                                                            var v172 uint32 = 42
                                                                            _ = v170
                                                                            _ = v171
                                                                            _ = v172
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 2
                                                                        }
                                                                        v1 = append(v1, 2147483648)
                                                                        var v173 uint = 4294967295
                                                                        if (!("a" <= "hello")) {
                                                                            var v174 int8 = int8(fi86)
                                                                            if ((v174 < (v174 << (((v2 & (v173 | (v2 & 42)))) & 63))) || v89) {
                                                                                if (0.0 <= float64(1.5)) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 510
                                                                                    }
                                                                                    var v175 uint8 = 128
                                                                                    _ = v175
                                                                                } else {
                                                                                    switch fi93 {
                                                                                    case -9223372036854775808:
                                                                                        var v176 float64 = float64(0.0)
                                                                                        _ = v176
                                                                                    case -703:
                                                                                        v105 = (v173 + (v2 - (v173 - v173)))
                                                                                    default:
                                                                                        v105 = (^(v2))
                                                                                    }
                                                                                    if v89 {
                                                                                        var v177 uint16 = (^(v3))
                                                                                        _ = v177
                                                                                    } else {
                                                                                        v178 := GenS2{F0: (v174 | v174)}
                                                                                        var v179 uint16 = (v3 & v3)
                                                                                        _ = v178
                                                                                        _ = v179
                                                                                    }
                                                                                    var v180 bool = false
                                                                                    _ = v180
                                                                                }
                                                                            } else {
                                                                                v174 -= v174
                                                                                if ((float32(3.14159) - float32((0.1 / (1e10 - float64(float32((3.14159 - 0.1))))))) > 0.1) {
                                                                                    if (!(0 != 361)) {
                                                                                        var v181 uint = v173
                                                                                        var v182 int16 = 0
                                                                                        fmt.Println(fi93, v2)
                                                                                        _ = v181
                                                                                        _ = v182
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                        v3 = (v3 & (v3 * 1))
                                                                                        v2 -= (v2 >> (((v105 << (((^(v173))) & 63))) & 63))
                                                                                    }
                                                                                    if (!((!((false && v89) || true)) || (2147483647 <= 1))) {
                                                                                        var v183 uint16 = v3
                                                                                        _ = v183
                                                                                    } else {
                                                                                        fmt.Println(v3, v105, v174)
                                                                                    }
                                                                                    var v184 uint16 = (v3 & v3)
                                                                                    v3 += 776
                                                                                    var v185 uint = v105
                                                                                    _ = v184
                                                                                    _ = v185
                                                                                } else {
                                                                                    var v186 int8 = v174
                                                                                    var v187 float64 = ((0.1 / -2.25) - 0.0)
                                                                                    fmt.Println(v105)
                                                                                    if (((v187 / v187) > -2.25) && (true || true)) {
                                                                                        var v188 int32 = -2147483648
                                                                                        var v189 int64 = 0
                                                                                        _ = v188
                                                                                        _ = v189
                                                                                    } else {
                                                                                        fmt.Println(fi93)
                                                                                        if (fi93 != fi86) {
                                                                                            switch "hello" {
                                                                                            case "a":
                                                                                                v190 := GenS3{F0: 18446744073709551615, F1: v3}
                                                                                                _ = v190
                                                                                            case "hello":
                                                                                                var v191 int = (fi86 << ((2) & 63))
                                                                                                _ = v191
                                                                                            case "foo bar":
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            default:
                                                                                                v1 = append(v1, 65535)
                                                                                            }
                                                                                            v192 := GenS2{F0: 1}
                                                                                            var v193 uint = v105
                                                                                            _ = v192
                                                                                            _ = v193
                                                                                        } else {
                                                                                            var v194 uint = 18446744073709551615
                                                                                            v174 = (v186 - v186)
                                                                                            v187 = v187
                                                                                            _ = v194
                                                                                        }
                                                                                    }
                                                                                    var v195 uint = (^(v2))
                                                                                    v196 := GenS9{F0: 0, F1: 526}
                                                                                    v197 := GenS15{F0: "\n\t"}
                                                                                    if v89 {
                                                                                        v198 := GenS20{F0: (fi86 % (((fi93 | fi86)) | 1))}
                                                                                        _ = len(v1)
                                                                                        var v199 uint64 = 0
                                                                                        _ = v198
                                                                                        _ = v199
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    if v89 {
                                                                                        var v200 uint16 = uint16(v2)
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            v173 -= (v105 >> (((v105 / ((v105) | 1))) & 63))
                                                                                        case "foo bar":
                                                                                            if (1 < 1) {
                                                                                                v195 -= 569
                                                                                                var v201 int32 = 1
                                                                                                _ = v201
                                                                                            }
                                                                                        }
                                                                                        if (true && (255 < v105)) {
                                                                                            _ = len(v1)
                                                                                            v202 := GenS21{F0: 42, F1: 32767, F2: 1000}
                                                                                            var v203 int = fi86
                                                                                            _ = v202
                                                                                            _ = v203
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        _ = v200
                                                                                    } else {
                                                                                        fmt.Println(v195, fi93)
                                                                                    }
                                                                                    v174 += (v174 | -3)
                                                                                    if false {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            v3 = v3
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v89, v105, v173)
                                                                                    _ = v186
                                                                                    _ = v187
                                                                                    _ = v195
                                                                                    _ = v196
                                                                                    _ = v197
                                                                                }
                                                                                if (42 <= 1) {
                                                                                    var v204 int64 = -9223372036854775808
                                                                                    if (42 < 0) {
                                                                                        if v89 {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                        }
                                                                                        var v205 int32 = -2147483648
                                                                                        fmt.Println(v204)
                                                                                        _ = v205
                                                                                    } else {
                                                                                        v206 := GenS22{F0: 864, F1: 217, F2: 42}
                                                                                        fmt.Println(v3, v2)
                                                                                        v207 := GenS15{F0: "\n\t"}
                                                                                        var v208 uint8 = 177
                                                                                        _ = v206
                                                                                        _ = v207
                                                                                        _ = v208
                                                                                    }
                                                                                    var v209 uint32 = 0
                                                                                    _ = v204
                                                                                    _ = v209
                                                                                }
                                                                            }
                                                                            fmt.Println(fi93, fi86, v174)
                                                                            _ = v174
                                                                        }
                                                                        v1 = append(v1, 2147483648)
                                                                        _ = v173
                                                                    } else {
                                                                        v2 += (v105 - v105)
                                                                        fmt.Println(v2, v2)
                                                                        v210 := GenS9{F0: -9223372036854775808, F1: 0}
                                                                        _ = v210
                                                                    }
                                                                    var v211 int32 = 816
                                                                    _ = v211
                                                                }
                                                                var v212 int64 = 1
                                                                _ = v212
                                                            } else {
                                                                var v213 bool = (!v89)
                                                                if (v213 && v89) {
                                                                    switch ("\n\t" + "a") {
                                                                    case "foo bar":
                                                                        if v89 {
                                                                            v3 += v3
                                                                        }
                                                                    }
                                                                    var v214 uint = (v2 >> (((v2 * (^(v105)))) & 63))
                                                                    fmt.Println(v89, v214)
                                                                    _ = v214
                                                                }
                                                                _ = v213
                                                            }
                                                            if true {
                                                                fmt.Println(fi93, v2, v2)
                                                                v105 = 179
                                                            } else {
                                                                if len(v1) > 0 {
                                                                    v1[0] = 732
                                                                }
                                                                if (!(uint32(v3) >= 0)) {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                } else {
                                                                    var v215 int32 = 42
                                                                    _ = v215
                                                                }
                                                            }
                                                            v216 := GenS20{F0: fi93}
                                                            _ = v216
                                                        }
                                                    case "":
                                                        fmt.Println(v2)
                                                    default:
                                                        switch "foo bar" {
                                                        case "a":
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        case "foo bar":
                                                            if (v89 && (false && v89)) {
                                                                if (!v89) {
                                                                    v3 = 0
                                                                }
                                                                v217 := GenS6{F0: 150, F1: 2147483648}
                                                                _ = v217
                                                            } else {
                                                                fmt.Println(fi93, fi86, fi93)
                                                                fmt.Println(v105)
                                                                v218 := GenS21{F0: 352, F1: v3, F2: 199}
                                                                _ = v218
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v105
                                        } else {
                                            v2 = (^(v2))
                                            v2 -= v2
                                        }
                                    case 2147483647:
                                        switch (-(fi93)) {
                                        case 42:
                                            fmt.Println(fi93, fi93, v2)
                                        case -422:
                                            switch (^(fi86)) {
                                            case 4294967295:
                                                if (!true) {
                                                    var v219 int8 = 23
                                                    var v220 string = "a"
                                                    _ = v219
                                                    _ = v220
                                                } else {
                                                    v1 = append(v1, 4294967295)
                                                }
                                            default:
                                                v221 := GenS23{F0: (v2 * (v2 ^ (^(v2)))), F1: ("hello" + ""), F2: (v89 || v89)}
                                                _ = v221
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        _ = v95
                    } else {
                        if v89 {
                            v2 += v2
                            v2 -= (v2 / ((2) | 1))
                            if true {
                                var v222 int64 = -1
                                var v223 int = (fi93 + -2147483648)
                                _ = v222
                                _ = v223
                            }
                            v89 = v89
                            v89 = v89
                            var v224 uint8 = 255
                            var v225 bool = (v224 == v224)
                            _ = v224
                            _ = v225
                        } else {
                            var v226 float64 = 1e10
                            if ((fi86 & (^(fi93))) == 936) {
                                v89 = (true || v89)
                                switch 128 {
                                case -9223372036854775808:
                                    v227 := GenS24{F0: (v2 << (((v2 + 4294967295)) & 63)), F1: 85}
                                    _ = v227
                                case 1:
                                    switch (fi93 >> ((326) & 63)) {
                                    case 32767:
                                        v3 = (^(v3))
                                    case 9223372036854775807:
                                        var v228 int64 = -487
                                        _ = v228
                                    case 76:
                                        v3 = (v3 + v3)
                                    default:
                                        v2 = (v2 + uint(v3))
                                    }
                                case 169:
                                    fmt.Println(fi86, fi93, v89)
                                }
                                switch (fi86 & (fi86 >> ((0) & 63))) {
                                case -9223372036854775808:
                                    var v229 uint8 = 42
                                    _ = v229
                                case 0:
                                    if true {
                                        _ = len(v1)
                                        var v230 bool = v89
                                        var v231 string = ((("foo bar" + (("hello" + "\n\t") + "")) + ("" + "hello")) + "")
                                        _ = v230
                                        _ = v231
                                    }
                                case 243:
                                    if v89 {
                                        var v232 bool = v89
                                        _ = len(v1)
                                        var v233 uint64 = 726
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                        v89 = v89
                                        _ = v232
                                        _ = v233
                                    } else {
                                        if ((v226 <= (float64((0.1 + (0.1 - -2.25))) / v226)) && v89) {
                                            var v234 int8 = 0
                                            _ = v234
                                        }
                                    }
                                default:
                                    var v235 uint16 = v3
                                    _ = v235
                                }
                            } else {
                                v226 = (1.5 - (-2.25 / 1.5))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v1 = append(v1, 0)
                            }
                            _ = v226
                        }
                        v3 = (v3 + v3)
                    }
                    fmt.Println(v3, fi86)
                }
                _ = fi93
            }
            v3 += v3
        }
        _ = fi86
        _ = v89
    }
    v3 += 289
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    var v236 uint16 = (v3 % (((v3 << ((v2) & 63))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v236:", v236)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v236
    _ = v3
    _ = v8
}
