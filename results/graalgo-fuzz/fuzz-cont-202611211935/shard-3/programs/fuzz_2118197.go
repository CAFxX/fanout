package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint64; F1 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 int32 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int32; F1 uint32 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int8; F1 uint32 }
type GenS18 struct { F0 int16; F1 uint64 }
type GenS19 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS20 struct { F0 uint8; F1 uint8; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 127
    v3 := make(map[uint32]bool)
    var v4 int8 = v2
    v5 := []int{-370, 32767}
    switch -2147483648 {
    case -2147483648:
        for fi6 := 0; fi6 < 9; fi6++ {
            if (42 > 729) {
                if true {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        switch 4294967295 {
                        case 4294967295:
                            v2 = v4
                        }
                        var v8 int8 = v2
                        v9 := v1[42]
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    if (true || (false || true)) {
                        for fi10 := 0; fi10 < 4; fi10++ {
                            fmt.Println(fi6, fi10, v4)
                            v4 += (v2 | (v2 << ((91) & 63)))
                            v2 = (v4 >> ((42) & 63))
                            _ = fi10
                        }
                        v11 := GenS1{F0: 2, F1: v2, F2: fi6}
                        fmt.Println(v2, v4)
                        _ = v11
                    }
                } else {
                    v2 = (v2 - v4)
                }
                fmt.Println(v4)
            } else {
                if (2 == 42) {
                    v12 := GenS2{F0: 64, F1: (^(v4)), F2: (fi6 << ((867) & 63))}
                    _ = len(v5)
                    _ = v12
                }
                for fi13 := 0; fi13 < 12; fi13++ {
                    v14 := GenS1{F0: 255, F1: (^(v4)), F2: 0}
                    switch (fi13 | fi6) {
                    case 558:
                        v4 = v2
                    }
                    var v15 uint = 977
                    _ = fi13
                    _ = v14
                    _ = v15
                }
            }
            if false {
                v16 := GenS1{F0: 230, F1: (v2 + (-(v4))), F2: fi6}
                _ = v16
            } else {
                switch ("\n\t" + (("hello" + "foo bar") + "a")) {
                case "hello":
                    v2 = (v2 & v2)
                default:
                    var v17 float64 = (1e10 + (-2.25 - (1.5 + -2.25)))
                    _ = v17
                }
                v2 = v4
            }
            _ = fi6
        }
    case 32767:
        if true {
            v2 = 54
            if false {
                if (v2 <= (-(v4))) {
                    v18 := GenS2{F0: 0, F1: 12, F2: 393}
                    var v19 bool = (0.1 < 1.5)
                    _ = v18
                    _ = v19
                }
                v3[116] = true
            } else {
                for fi20 := 0; fi20 < 10; fi20++ {
                    _ = len(v5)
                    var v21 bool = false
                    _ = fi20
                    _ = v21
                }
                var v22 uint16 = 994
                _ = v22
            }
            switch 1 {
            case 9223372036854775807:
                if false {
                    if true {
                        v4 += (v4 / ((v4) | 1))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v23 bool = true
                        _ = v23
                    }
                    for fi24 := 0; fi24 < 5; fi24++ {
                        v25 := GenS3{F0: ("" + ("a" + "hello"))}
                        v3[909] = false
                        _ = fi24
                        _ = v25
                    }
                    v1[65535] = 2
                }
            default:
                v4 = 1
            }
        } else {
            switch ("a" + "a") {
            case "foo bar":
                var v26 int32 = 1
                _ = v26
            case "":
                v27 := GenS2{F0: 114, F1: 0, F2: 9223372036854775807}
                _ = v27
            case "hello":
                for fi28 := 0; fi28 < 13; fi28++ {
                    for fi29 := 0; fi29 < 13; fi29++ {
                        var v30 bool = false
                        switch ("hello" + ("" + "hello")) {
                        case "\n\t":
                            var v31 int8 = 100
                            _ = v31
                        case "a":
                            var v32 int8 = 100
                            _ = v32
                        }
                        _ = fi29
                        _ = v30
                    }
                    fmt.Println(fi28, v2)
                    var v33 int = 32768
                    _ = fi28
                    _ = v33
                }
            }
            if true {
                switch 4294967295 {
                case 0:
                    v4 -= v2
                case 99:
                    var v34 uint16 = 65535
                    _ = v34
                }
            } else {
                delete(v3, 4294967295)
                v2 -= v4
                v35 := GenS2{F0: 1, F1: -110, F2: 0}
                _ = v35
            }
            v2 = v2
        }
    case 65:
        v5 = append(v5, 4294967295)
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v2 = v4
    var v36 uint8 = 127
    for fi37 := 0; fi37 < 20; fi37++ {
        _ = len(v5)
        if true {
            v38 := GenS3{F0: (("a" + "foo bar") + "a")}
            if false {
                var v39 int16 = 32767
                switch fi37 {
                case -2147483648:
                    v39 = 32767
                case 0:
                    switch (fi37 * (fi37 % (((fi37 | fi37)) | 1))) {
                    case 0:
                        var v40 int8 = 18
                        _ = v40
                    case 42:
                        if (0 == v39) {
                            v2 = (^(v4))
                        }
                    case 128:
                        var v41 float32 = -2.25
                        _ = v41
                    default:
                        if false {
                            var v42 int64 = 0
                            var v43 int8 = -128
                            for fi44 := 0; fi44 < 10; fi44++ {
                                v4 = v2
                                _ = fi44
                            }
                            _ = v42
                            _ = v43
                        } else {
                            var v45 uint64 = 760
                            var v46 uint16 = 65535
                            if true {
                                for fi47 := 0; fi47 < 11; fi47++ {
                                    if (!false) {
                                        var v48 bool = ((v39 & (v39 * -32768)) <= -259)
                                        _ = v48
                                    } else {
                                        v46 -= (v46 + 317)
                                        fmt.Println(fi37, v2)
                                    }
                                    var v49 int = -409
                                    _ = fi47
                                    _ = v49
                                }
                                if (((-2.25 == (-2.25 + 1.5)) || ("" <= "\n\t")) || true) {
                                    var v50 uint32 = 42
                                    for fi51 := 0; fi51 < 8; fi51++ {
                                        if len(v5) > 0 {
                                            v5[0] = (fi37 - fi51)
                                        }
                                        if ((0.0 - ((3.14159 + 1e10) + (-2.25 * (1e10 / float32((0.1 / -2.25)))))) >= 3.14159) {
                                            if (!(v39 > v39)) {
                                                var v52 int32 = -2147483648
                                                fmt.Println(v39, v45)
                                                _ = v52
                                            }
                                            if (((int8(v46) == int8(v45)) && (!(!(788 > 9223372036854775807)))) && false) {
                                                if ((v50 * (^(v50))) > 0) {
                                                    _ = len(v3)
                                                } else {
                                                    v1[334] = (v50 / (((v50 >> ((65535) & 63))) | 1))
                                                }
                                                var v53 uint = 865
                                                switch (("a" + "\n\t") + (("foo bar" + "") + "")) {
                                                case "a":
                                                    v4 = v4
                                                default:
                                                    var v54 uint64 = v45
                                                    _ = v54
                                                }
                                                _ = v53
                                            }
                                        }
                                        _ = fi51
                                    }
                                    _ = v50
                                } else {
                                    if false {
                                        for fi55 := 0; fi55 < 12; fi55++ {
                                            _ = len(v5)
                                            v56 := GenS2{F0: 107, F1: (v4 | v4), F2: 877}
                                            v4 += (v4 << ((300) & 63))
                                            _ = fi55
                                            _ = v56
                                        }
                                        var v57 bool = false
                                        _ = v57
                                    } else {
                                        var v58 uint16 = v46
                                        for fi59 := 0; fi59 < 9; fi59++ {
                                            if (660 != v58) {
                                                v4 = v4
                                                switch (fi37 * (fi37 >> ((110) & 63))) {
                                                case 32767:
                                                    var v60 int8 = 0
                                                    _ = v60
                                                case 32768:
                                                    var v61 int8 = (v2 ^ 127)
                                                    _ = v61
                                                case -369:
                                                    if (!true) {
                                                        fmt.Println(v36)
                                                        var v62 uint64 = (v45 >> ((305) & 63))
                                                        switch "" {
                                                        case "hello":
                                                            if (false || (!(-804 >= (fi37 * fi59)))) {
                                                                var v63 float32 = 0.0
                                                                v62 -= 810
                                                                _ = v63
                                                            } else {
                                                                _ = len(v3)
                                                            }
                                                        }
                                                        _ = v62
                                                    }
                                                default:
                                                    if false {
                                                        v2 -= (v4 + v2)
                                                        var v64 uint32 = uint32(v45)
                                                        _ = v64
                                                    } else {
                                                        if true {
                                                            if (!(uint(v46) > 4294967295)) {
                                                                v46 = v58
                                                                if (uint64(v36) >= (v45 >> ((18446744073709551615) & 63))) {
                                                                    switch ("" + "") {
                                                                    case "a":
                                                                        fmt.Println(fi59, v58)
                                                                    case "hello":
                                                                        v2 -= (v2 << ((0) & 63))
                                                                    case "\n\t":
                                                                        v39 += v39
                                                                    }
                                                                    v36 = v36
                                                                    v65 := GenS1{F0: v36, F1: v2, F2: (fi59 * fi37)}
                                                                    _ = v65
                                                                } else {
                                                                    if false {
                                                                        switch (fi37 % (((fi59 * fi37)) | 1)) {
                                                                        case 214:
                                                                            delete(v3, 2147483648)
                                                                        case 4294967295:
                                                                            v58 += v46
                                                                        default:
                                                                            var v66 float32 = 3.14159
                                                                            _ = v66
                                                                        }
                                                                        v2 = -1
                                                                    }
                                                                }
                                                                v67 := GenS2{F0: (v36 + v36), F1: v2, F2: fi37}
                                                                fmt.Println(v45, v4, v39)
                                                                var v68 float64 = (0.1 / 1.5)
                                                                v69 := GenS4{F0: (^(v45)), F1: 42, F2: (v46 & (v58 << ((528) & 63)))}
                                                                v45 += v45
                                                                v70 := GenS2{F0: v36, F1: (v2 - v4), F2: fi59}
                                                                var v71 int16 = v39
                                                                fmt.Println(v4, v39, v2)
                                                                _ = v67
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            v72 := GenS1{F0: (v36 << ((685) & 63)), F1: v4, F2: fi59}
                                                            if (-455 < 42) {
                                                                if (!true) {
                                                                    if ("\n\t" != "\n\t") {
                                                                        var v73 int32 = -1
                                                                        v2 = v4
                                                                        switch 32767 {
                                                                        case 32768:
                                                                            v45 = 718
                                                                        case -2147483648:
                                                                            fmt.Println(v4, v2)
                                                                        case 2147483647:
                                                                            fmt.Println(v45, v58)
                                                                        default:
                                                                            var v74 int32 = (v73 % ((v73) | 1))
                                                                            _ = v74
                                                                        }
                                                                        delete(v1, 2147483648)
                                                                        _ = v73
                                                                    } else {
                                                                        v45 = (^(v45))
                                                                    }
                                                                    var v75 int = 0
                                                                    _ = v75
                                                                }
                                                                v76 := GenS5{F0: (^(v45))}
                                                                _ = v76
                                                            } else {
                                                                var v77 float32 = 3.14159
                                                                fmt.Println(v39, v58)
                                                                _ = v77
                                                            }
                                                            _ = v72
                                                        }
                                                    }
                                                }
                                            }
                                            if (4294967295 < 795) {
                                                v58 -= v58
                                            }
                                            var v78 uint8 = (v36 | (v36 >> ((4294967295) & 63)))
                                            _ = fi59
                                            _ = v78
                                        }
                                        _ = v58
                                    }
                                    v4 -= v4
                                    v79 := GenS1{F0: uint8(v45), F1: v2, F2: fi37}
                                    _ = v79
                                }
                                switch "hello" {
                                case "\n\t":
                                    if (false || true) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            for fi80 := 0; fi80 < 12; fi80++ {
                                                if (!(255 >= 599)) {
                                                    if (true || (v46 > (v46 | 854))) {
                                                        if len(v5) > 0 {
                                                            v5[0] = 32768
                                                        }
                                                    }
                                                    v39 += (v39 << ((65) & 63))
                                                    _ = len(v3)
                                                    v5 = append(v5, (fi80 * -427))
                                                } else {
                                                    switch int(v36) {
                                                    case 236:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if (!(false && (false && (!false)))) {
                                                            if (!(false && (0 > (v45 ^ (v45 + (v45 * v45)))))) {
                                                                v39 += (v39 >> ((2) & 63))
                                                            }
                                                            v45 += (^(v45))
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            _ = len(v3)
                                                            var v81 int8 = v2
                                                            v36 -= v36
                                                            _ = v81
                                                        }
                                                    case -812:
                                                        v82 := GenS6{F0: 0, F1: 2147483648}
                                                        _ = v82
                                                    case 128:
                                                        if (v46 == v46) {
                                                            fmt.Println(v39, v39)
                                                            var v83 float64 = (-2.25 - 0.0)
                                                            switch fi37 {
                                                            case 623:
                                                                switch ("hello" + ("" + ("\n\t" + "a"))) {
                                                                case "hello":
                                                                    var v84 float64 = (v83 * (v83 * v83))
                                                                    v2 -= v4
                                                                    var v85 int64 = -1
                                                                    _ = v84
                                                                    _ = v85
                                                                case "a":
                                                                    var v86 float32 = ((0.0 - -2.25) / -2.25)
                                                                    _ = v86
                                                                }
                                                            default:
                                                                switch fi80 {
                                                                case -67:
                                                                    v45 = v45
                                                                case 127:
                                                                    switch (fi80 | (fi37 & -383)) {
                                                                    case 841:
                                                                        v87 := GenS1{F0: v36, F1: (v4 % ((v2) | 1)), F2: int(v4)}
                                                                        _ = v87
                                                                    case 42:
                                                                        switch -696 {
                                                                        case 1:
                                                                            if ((874 != 2) && true) {
                                                                                v88 := GenS4{F0: 18446744073709551615, F1: 0, F2: v46}
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        v89 := GenS7{F0: v45, F1: 197}
                                                                        _ = v89
                                                                    default:
                                                                        if true {
                                                                            var v90 bool = (-958 < 0)
                                                                            _ = v90
                                                                        } else {
                                                                            v4 = (v2 - -99)
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                v91 := GenS5{F0: 527}
                                                                                _ = v91
                                                                            case "\n\t":
                                                                                v92 := GenS2{F0: v36, F1: (^(v2)), F2: (fi80 % ((2147483647) | 1))}
                                                                                _ = v92
                                                                            case "hello":
                                                                                v93 := GenS4{F0: (v45 * (^(v45))), F1: 9223372036854775807, F2: (v46 / (((v46 - v46)) | 1))}
                                                                                _ = v93
                                                                            default:
                                                                                if true {
                                                                                    if (int8(v45) == (v2 / ((v4) | 1))) {
                                                                                        var v94 string = "a"
                                                                                        var v95 uint32 = 4294967295
                                                                                        fmt.Println(v39)
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    }
                                                                                    var v96 int8 = int8(v39)
                                                                                    _ = v96
                                                                                } else {
                                                                                    if false {
                                                                                        fmt.Println(fi37)
                                                                                        var v97 uint8 = (v36 % ((v36) | 1))
                                                                                        _ = v97
                                                                                    } else {
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = (fi37 >> ((1) & 63))
                                                                                        }
                                                                                        switch fi80 {
                                                                                        case 4294967295:
                                                                                            delete(v3, 42)
                                                                                        case 115:
                                                                                            fmt.Println(v36, fi37, v2)
                                                                                        case 32768:
                                                                                            var v98 int64 = -1
                                                                                            v36 = (^(v36))
                                                                                            switch -2147483648 {
                                                                                            case 205:
                                                                                                v99 := GenS2{F0: (v36 / ((159) | 1)), F1: v2, F2: (fi37 + 42)}
                                                                                                _ = v99
                                                                                            case 127:
                                                                                                v100 := GenS2{F0: v36, F1: v2, F2: fi37}
                                                                                                _ = v100
                                                                                            }
                                                                                            _ = v98
                                                                                        default:
                                                                                            v39 += v39
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v39)
                                                                                    if false {
                                                                                        var v101 uint64 = (v45 ^ 1)
                                                                                        if false {
                                                                                            fmt.Println(v39, v101)
                                                                                        }
                                                                                        var v102 int64 = int64(v36)
                                                                                        _ = v101
                                                                                        _ = v102
                                                                                    }
                                                                                    v4 -= v2
                                                                                    v39 = int16(v46)
                                                                                    var v103 bool = (!false)
                                                                                    v104 := GenS7{F0: 4294967295, F1: v45}
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            switch (fi80 >> ((2) & 63)) {
                                                                            case 285:
                                                                                if true {
                                                                                    delete(v3, 4294967295)
                                                                                }
                                                                            case 32768:
                                                                                v36 = uint8(v45)
                                                                            case -1:
                                                                                v45 -= 42
                                                                            }
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (fi80 | 9223372036854775807)
                                                                            }
                                                                        }
                                                                    }
                                                                case 9223372036854775807:
                                                                    v83 = 3.14159
                                                                default:
                                                                    var v105 int = 0
                                                                    _ = v105
                                                                }
                                                            }
                                                            _ = v83
                                                        } else {
                                                            _ = len(v5)
                                                            var v106 int16 = v39
                                                            _ = v106
                                                        }
                                                    }
                                                    fmt.Println(v45)
                                                }
                                                v45 += (^(v45))
                                                var v107 uint16 = 2
                                                _ = fi80
                                                _ = v107
                                            }
                                        case "a":
                                            if false {
                                                var v108 bool = (fi37 < (fi37 | (^(fi37))))
                                                v109 := GenS8{F0: v45, F1: 990}
                                                if v108 {
                                                    v46 = (^(v46))
                                                }
                                                _ = v108
                                                _ = v109
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v45
                            _ = v46
                        }
                    }
                case -735:
                    var v110 int16 = (v39 | -1)
                    _ = v110
                }
                _ = v39
            } else {
                for fi111 := 0; fi111 < 14; fi111++ {
                    fmt.Println(fi111, v36)
                    _ = fi111
                }
                if (!(618 <= -1)) {
                    switch ("hello" + "") {
                    case "":
                        v2 += v2
                    }
                } else {
                    fmt.Println(v2, v2, fi37)
                    for fi112 := 0; fi112 < 17; fi112++ {
                        v113 := GenS9{F0: 963}
                        _ = fi112
                        _ = v113
                    }
                }
                if len(v5) > 0 {
                    v5[0] = 289
                }
            }
            v2 += 1
            _ = v38
        } else {
            if (367 < -1) {
                if false {
                    v114 := v3[0]
                    switch (fi37 + 9223372036854775807) {
                    case 588:
                        switch fi37 {
                        case -2147483648:
                            var v115 int = (fi37 % ((fi37) | 1))
                            _ = v115
                        case -337:
                            fmt.Println(v4)
                        default:
                            if v114 {
                                v116 := GenS10{F0: ("hello" < "foo bar")}
                                v4 = 6
                                _ = v116
                            } else {
                                switch fi37 {
                                case -1:
                                    v114 = (5 <= v36)
                                case 4294967295:
                                    fmt.Println(v4)
                                case -2147483648:
                                    switch fi37 {
                                    case 9223372036854775807:
                                        v4 = v4
                                    default:
                                        for fi117 := 0; fi117 < 2; fi117++ {
                                            switch (-(fi117)) {
                                            case -345:
                                                var v118 uint64 = 4294967295
                                                _ = v118
                                            case 594:
                                                var v119 float32 = -2.25
                                                _ = v119
                                            case 2147483647:
                                                v36 -= (^(v36))
                                                v120 := GenS11{F0: ("" + "foo bar")}
                                                _ = v120
                                            }
                                            _ = fi117
                                        }
                                    }
                                default:
                                    v4 += 1
                                }
                                if false {
                                    switch ("hello" + "\n\t") {
                                    case "hello":
                                        for fi121 := 0; fi121 < 14; fi121++ {
                                            v4 += 100
                                            var v122 uint8 = (^(v36))
                                            _ = fi121
                                            _ = v122
                                        }
                                    default:
                                        for fi123 := 0; fi123 < 20; fi123++ {
                                            if v114 {
                                                v124 := GenS12{F0: (^(fi123)), F1: 29}
                                                v36 -= (v36 >> ((1) & 63))
                                                _ = v124
                                            } else {
                                                v125 := GenS13{F0: 813}
                                                v126 := GenS14{F0: 1, F1: 0}
                                                var v127 uint8 = 2
                                                _ = v125
                                                _ = v126
                                                _ = v127
                                            }
                                            if v114 {
                                                v2 = v4
                                            }
                                            if v114 {
                                                var v128 int64 = 1
                                                var v129 int8 = 42
                                                _ = v128
                                                _ = v129
                                            } else {
                                                if true {
                                                    v2 = (v2 % (((v2 * v2)) | 1))
                                                    v130 := GenS11{F0: "foo bar"}
                                                    _ = v130
                                                }
                                                if (1 == 871) {
                                                    var v131 int = (fi123 ^ fi123)
                                                    _ = v131
                                                } else {
                                                    v3[0] = v114
                                                }
                                                if v114 {
                                                    var v132 float32 = (1.5 * (0.0 / -2.25))
                                                    v133 := GenS15{F0: 18446744073709551615}
                                                    _ = v132
                                                    _ = v133
                                                }
                                            }
                                            _ = fi123
                                        }
                                    }
                                }
                                v36 += v36
                            }
                        }
                    default:
                        v2 = (v4 & (v4 - (v4 + (v2 | (v2 + v2)))))
                    }
                    _ = v114
                } else {
                    v134 := GenS4{F0: uint64(v36), F1: -384, F2: 32767}
                    v135 := GenS14{F0: 1, F1: 65535}
                    if true {
                        for fi136 := 0; fi136 < 18; fi136++ {
                            delete(v1, 0)
                            v2 -= 42
                            switch "foo bar" {
                            case "foo bar":
                                var v137 uint32 = 989
                                _ = v137
                            case "":
                                _ = len(v5)
                            case "a":
                                var v138 uint = 42
                                v139 := GenS4{F0: 42, F1: 9223372036854775807, F2: 0}
                                _ = v138
                                _ = v139
                            }
                            _ = fi136
                        }
                    } else {
                        var v140 bool = (false && true)
                        var v141 uint64 = 148
                        delete(v1, 4294967295)
                        _ = v140
                        _ = v141
                    }
                    _ = v134
                    _ = v135
                }
                v5 = append(v5, 4294967295)
                v142 := GenS6{F0: 661, F1: 4294967295}
                _ = v142
            }
            if true {
                if (!(0 < 754)) {
                    v36 -= v36
                    v143 := v3[2147483647]
                    _ = v143
                } else {
                    if true {
                        v4 -= v2
                        v4 = 113
                        for fi144 := 0; fi144 < 8; fi144++ {
                            v36 += 222
                            if false {
                                switch 525 {
                                case -1:
                                    v145 := v3[4294967295]
                                    _ = v145
                                }
                                var v146 int16 = -32768
                                _ = v146
                            } else {
                                if (!false) {
                                    fmt.Println(v36, v4, fi144)
                                    _ = len(v3)
                                    var v147 int = (fi37 & fi37)
                                    v147 += 127
                                    var v148 int32 = 0
                                    _ = v147
                                    _ = v148
                                } else {
                                    v149 := GenS16{F0: 2147483647}
                                    switch (("" + "hello") + "a") {
                                    case "":
                                        if (18446744073709551615 != 65535) {
                                            fmt.Println(v4, fi37)
                                            switch ("" + "") {
                                            case "foo bar":
                                                v4 += v2
                                            default:
                                                v2 = (v4 - (v2 - 34))
                                            }
                                        } else {
                                            var v150 uint32 = 526
                                            _ = v150
                                        }
                                    case "\n\t":
                                        var v151 float32 = 0.1
                                        _ = v151
                                    case "foo bar":
                                        v152 := GenS17{F0: v4, F1: 65535}
                                        _ = v152
                                    }
                                    _ = v149
                                }
                            }
                            v36 += v36
                            _ = fi144
                        }
                    } else {
                        var v153 int32 = 42
                        v36 = v36
                        _ = v153
                    }
                    var v154 uint16 = 32767
                    switch fi37 {
                    case -1:
                        v155 := GenS18{F0: 1000, F1: 0}
                        _ = v155
                    }
                    _ = v154
                }
            } else {
                var v156 uint = 711
                _ = v156
            }
        }
        var v157 float64 = 3.14159
        _ = fi37
        _ = v157
    }
    fmt.Println(v4, v2, v4)
    for fi158 := 0; fi158 < 8; fi158++ {
        var v159 int8 = (v4 << ((1) & 63))
        if false {
            v160 := GenS8{F0: 18446744073709551615, F1: 65535}
            switch "foo bar" {
            case "hello":
                v1[2147483647] = 2147483648
            case "":
                var v161 int = (^(fi158))
                _ = v161
            }
            v4 += 1
            _ = v160
        }
        fmt.Println(fi158, v2, fi158)
        _ = fi158
        _ = v159
    }
    var v162 uint = 285
    if false {
        for fi163 := 0; fi163 < 5; fi163++ {
            v2 -= v2
            fmt.Println(v2)
            _ = fi163
        }
    } else {
        var v164 uint16 = 1
        v1[65535] = 590
        v2 = (v4 << ((18446744073709551615) & 63))
        _ = v164
    }
    v36 += v36
    v36 += v36
    v162 -= v162
    if false {
        switch "foo bar" {
        case "hello":
            v165 := GenS19{F0: 313, F1: 581, F2: 128}
            _ = v165
        case "\n\t":
            for fi166 := 0; fi166 < 3; fi166++ {
                v162 = (v162 * (v162 + v162))
                if false {
                    v167 := GenS1{F0: (^(v36)), F1: -128, F2: 32767}
                    _ = v167
                } else {
                    v2 = (-(v2))
                    if len(v5) > 0 {
                        v5[0] = (fi166 % ((fi166) | 1))
                    }
                    v2 = (^(v4))
                }
                _ = fi166
            }
        default:
            if (!(32767 != 32767)) {
                switch "" {
                case "a":
                    v168 := GenS20{F0: 244, F1: (v36 * (v36 - (v36 << ((v162) & 63)))), F2: 42}
                    _ = v168
                case "hello":
                    v36 -= (v36 - 184)
                case "\n\t":
                    delete(v3, 4294967295)
                }
                var v169 uint16 = 917
                _ = v169
            }
        }
    } else {
        var v170 uint32 = 2147483648
        if true {
            var v171 bool = true
            _ = v171
        }
        _ = v170
    }
    var v172 int = 32768
    if false {
        fmt.Println(v2, v2)
    } else {
        if (42 == 19) {
            for fi173 := 0; fi173 < 18; fi173++ {
                for fi174 := 0; fi174 < 18; fi174++ {
                    fmt.Println(v162, fi173)
                    v175 := GenS20{F0: uint8(v162), F1: (v36 & v36), F2: 936}
                    _ = fi174
                    _ = v175
                }
                fmt.Println(v4)
                if false {
                    if false {
                        v1[1] = 65535
                        fmt.Println(v36)
                        var v176 uint8 = v36
                        _ = v176
                    }
                }
                _ = fi173
            }
            var v177 float64 = (float64(-2.25) / (float64(1.5) - 3.14159))
            _ = v177
        }
        for fi178 := 0; fi178 < 3; fi178++ {
            v2 += v2
            _ = fi178
        }
        v179 := v1[2]
        _ = v179
    }
    v5 = append(v5, (v172 << ((v162) & 63)))
    if true {
        fmt.Println(v172, v36)
    } else {
        var v180 string = (("foo bar" + (("" + "foo bar") + "hello")) + ("\n\t" + ("" + "foo bar")))
        v2 = v4
        _ = v180
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v162:", v162)
    fmt.Println("v172:", v172)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v162
    _ = v172
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
