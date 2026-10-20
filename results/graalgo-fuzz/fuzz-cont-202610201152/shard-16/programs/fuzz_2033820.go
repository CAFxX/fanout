package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint32; F2 int8 }
type GenS6 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS7 struct { F0 int32; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint64; F1 int }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 int16; F1 uint8; F2 bool }
type GenS12 struct { F0 bool; F1 uint8 }
type GenS13 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS14 struct { F0 int32; F1 int64; F2 int }
type GenS15 struct { F0 uint; F1 int8 }
type GenS16 struct { F0 uint64; F1 uint }
type GenS17 struct { F0 int8 }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint16 = 4
        switch fi7 {
        case 955:
            v8 += v8
        case 128:
            v9 := GenS1{F0: 25}
            _ = v9
        case 4294967295:
            v10 := GenS1{F0: 0}
            _ = v10
        default:
            switch "\n\t" {
            case "":
                var v11 int64 = 42
                _ = v11
            case "a":
                v12 := []int{int(v8), 703, (^(fi7)), (fi7 ^ 165)}
                _ = v12
            }
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS1{F0: 255}
    switch 32767 {
    case -2147483648:
        var v14 uint32 = 4294967295
        _ = v14
    default:
        var v15 uint = 458
        _ = v15
    }
    _ = v13
    return 621
}

func main() {
    var v1 uint32 = 2147483647
    var v2 bool = (!false)
    var v3 float64 = (0.1 + -2.25)
    var v4 float64 = 0.0
    var v5 float64 = (-2.25 + 1e10)
    var v6 bool = true
    if (42 >= 255) {
        fmt.Println(v4)
        if v6 {
            for fi16 := 0; fi16 < 19; fi16++ {
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch (fi16 & 1) {
                    case 1:
                        fmt.Println(v4, v5)
                    default:
                        switch fi16 {
                        case -1:
                            switch "\n\t" {
                            case "":
                                fmt.Println(fi17, v2, v4)
                            default:
                                v18 := []int{fi16, (fi16 << ((18446744073709551615) & 63))}
                                _ = v18
                            }
                        case 9223372036854775807:
                            fmt.Println(v3)
                        default:
                            v19 := GenS2{F0: "hello", F1: v2, F2: 118}
                            _ = v19
                        }
                    }
                    v20 := []int{}
                    fmt.Println(v6, v5, fi16)
                    _ = fi17
                    _ = v20
                }
                for fi21 := 0; fi21 < 3; fi21++ {
                    var v22 uint64 = 18446744073709551615
                    v5 -= v3
                    _ = fi21
                    _ = v22
                }
                v2 = true
                _ = fi16
            }
            v23 := []int{-383, 570, -771, -2147483648}
            _ = v23
        } else {
            var v24 float32 = 0.0
            var v25 uint32 = 2
            _ = v24
            _ = v25
        }
    }
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 int8 = -87
        switch fi26 {
        case 70:
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 float32 = (0.0 * 1e10)
                _ = fi28
                _ = v29
            }
        case 128:
            var v30 int8 = (-(v27))
            _ = v30
        case 4294967295:
            for fi31 := 0; fi31 < 5; fi31++ {
                v2 = true
                fmt.Println(v2, v5, v2)
                _ = fi31
            }
        default:
            switch "a" {
            case "\n\t":
                var v32 float32 = ((3.14159 - (0.1 * 0.0)) - float32(v3))
                _ = v32
            case "":
                fmt.Println(v5, v27, v2)
            case "foo bar":
                v33 := []int{int(v1), 9223372036854775807, fi26, (fi26 - 4294967295)}
                _ = v33
            }
        }
        _ = fi26
        _ = v27
    }
    for fi34 := 0; fi34 < 14; fi34++ {
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 string = "foo bar"
            fmt.Println(v36, v2, v1)
            fmt.Println(v2, fi35)
            _ = fi35
            _ = v36
        }
        fmt.Println(v6, v3, fi34)
        for fi37 := 0; fi37 < 13; fi37++ {
            switch (^(fi37)) {
            case 32768:
                v4 = v4
            case -729:
                v6 = v2
            default:
                v38 := GenS2{F0: "", F1: v6, F2: 1}
                _ = v38
            }
            if ((false && v2) || v6) {
                var v39 uint16 = 210
                _ = v39
            } else {
                fmt.Println(v5, v3, fi34)
                if false {
                    switch "hello" {
                    case "foo bar":
                        v2 = v6
                    case "\n\t":
                        var v40 string = "foo bar"
                        var v41 int32 = 339
                        v42 := []int{1, (fi37 + (fi34 << ((4294967295) & 63))), fi37, (fi34 | (fi37 << ((4294967295) & 63)))}
                        _ = v40
                        _ = v41
                        _ = v42
                    default:
                        v1 = v1
                    }
                }
            }
            var v43 float32 = 0.1
            _ = fi37
            _ = v43
        }
        _ = fi34
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        v45 := []int{fi44, fi44, fi44, (fi44 & 128)}
        v3 = (v3 - 3.14159)
        _ = fi44
        _ = v45
    }
    if v6 {
        v46 := GenS3{F0: 65535, F1: 1, F2: -1}
        for fi47 := 0; fi47 < 2; fi47++ {
            for fi48 := 0; fi48 < 11; fi48++ {
                switch "" {
                case "\n\t":
                    switch fi47 {
                    case -279:
                        switch fi47 {
                        case -505:
                            v49 := []int{}
                            _ = v49
                        case 1:
                            v1 = (v1 * (v1 & v1))
                        case 9223372036854775807:
                            var v50 int8 = 127
                            _ = v50
                        }
                    case -9223372036854775808:
                        var v51 uint32 = 49
                        _ = v51
                    case 2147483647:
                        switch (fi48 ^ 32767) {
                        case -415:
                            v52 := GenS1{F0: 255}
                            _ = v52
                        default:
                            fmt.Println(fi48, v3, v3)
                        }
                    default:
                        v6 = v2
                        if v6 {
                            v5 = 1e10
                            v53 := []int{}
                            fmt.Println(v6, v5, v6)
                            v6 = (!v2)
                            _ = v53
                        }
                        switch ("\n\t" + ("hello" + "a")) {
                        case "":
                            var v54 float32 = (0.1 * (3.14159 / (-2.25 / 1.0)))
                            _ = v54
                        }
                        switch "hello" {
                        case "\n\t":
                            v55 := GenS4{F0: -2147483648, F1: v1}
                            _ = v55
                        case "foo bar":
                            v56 := []int{}
                            _ = v56
                        default:
                            v57 := GenS2{F0: "a", F1: v2, F2: 0}
                            _ = v57
                        }
                        v58 := []int{fi48, (fi47 | (fi48 - (^(fi47)))), (fi47 | (fi47 + 32767)), 2147483647}
                        switch fi47 {
                        case 2147483647:
                            switch ("foo bar" + "foo bar") {
                            case "":
                                switch "a" {
                                case "":
                                    if len(v58) > 0 {
                                        v58[0] = (fi47 << ((584) & 63))
                                    }
                                    if false {
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v3, v6, v6)
                                        default:
                                            switch fi47 {
                                            case 1:
                                                if v6 {
                                                    v59 := GenS5{F0: -9223372036854775808, F1: (v1 % (((v1 ^ 4294967295)) | 1)), F2: 8}
                                                    switch ("\n\t" + "hello") {
                                                    case "a":
                                                        if v2 {
                                                            var v60 int = fi48
                                                            v1 = v1
                                                            var v61 int8 = 1
                                                            v62 := GenS6{F0: 18446744073709551615, F1: 9223372036854775807, F2: (v1 - v1)}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        }
                                                    default:
                                                        switch "" {
                                                        case "a":
                                                            switch fi48 {
                                                            case 9223372036854775807:
                                                                var v63 uint64 = 18446744073709551615
                                                                _ = v63
                                                            }
                                                        default:
                                                            v3 += v5
                                                        }
                                                    }
                                                    _ = len(v58)
                                                    _ = v59
                                                }
                                            case -9223372036854775808:
                                                var v64 uint32 = v1
                                                var v65 uint = 65535
                                                if len(v58) >= 2 {
                                                    v58 = v58[1:len(v58)]
                                                }
                                                _ = v64
                                                _ = v65
                                            default:
                                                if ((-(fi48)) > (-(fi48))) {
                                                    if true {
                                                        v5 += v4
                                                        v6 = (true || true)
                                                        switch "hello" {
                                                        case "hello":
                                                            var v66 bool = (!(99 == 127))
                                                            _ = v66
                                                        case "\n\t":
                                                            var v67 uint = 0
                                                            _ = v67
                                                        case "":
                                                            if (v2 && (-2147483648 > 0)) {
                                                                v58 = append(v58, (-(fi48)))
                                                                if ((-999 != 2147483647) && false) {
                                                                    v68 := GenS7{F0: 1, F1: 128, F2: 0}
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        v69 := GenS1{F0: 2}
                                                                        _ = v69
                                                                    default:
                                                                        var v70 float32 = -2.25
                                                                        _ = v70
                                                                    }
                                                                    switch (fi48 - fi48) {
                                                                    case 32768:
                                                                        v71 := GenS8{F0: 42, F1: fi48}
                                                                        if v6 {
                                                                            var v72 float64 = float64(((-2.25 - (((1e10 / 3.14159) + (1.5 / -2.25)) + float32(v3))) + -2.25))
                                                                            switch (fi48 % ((fi48) | 1)) {
                                                                            case 4294967295:
                                                                                v4 -= v5
                                                                            default:
                                                                                v3 = (float64(0.1) * 1.5)
                                                                            }
                                                                            v3 = 0.1
                                                                            _ = v72
                                                                        }
                                                                        if len(v58) > 0 {
                                                                            v58[0] = 32767
                                                                        }
                                                                        _ = v71
                                                                    case 70:
                                                                        if (2147483647 <= v1) {
                                                                            var v73 int64 = 1
                                                                            v3 += ((v3 / (0.0 + v3)) + 0.0)
                                                                            _ = v73
                                                                        }
                                                                    case -2147483648:
                                                                        if (v6 || true) {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v74 uint16 = 807
                                                                                _ = v74
                                                                            }
                                                                            v75 := GenS9{F0: 42, F1: 545}
                                                                            _ = v75
                                                                        }
                                                                    default:
                                                                        var v76 float64 = (0.1 + (float64(-2.25) + v4))
                                                                        _ = v76
                                                                    }
                                                                    _ = v68
                                                                }
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    v1 = (v1 >> ((0) & 63))
                                                    switch ((("foo bar" + ("foo bar" + "a")) + "\n\t") + "\n\t") {
                                                    case "foo bar":
                                                        var v77 uint = 65535
                                                        _ = v77
                                                    case "":
                                                        var v78 int8 = -46
                                                        _ = v78
                                                    case "hello":
                                                        v79 := GenS10{F0: 255, F1: 289}
                                                        _ = v79
                                                    }
                                                    switch 9223372036854775807 {
                                                    case 32768:
                                                        var v80 uint16 = 65535
                                                        _ = v80
                                                    }
                                                }
                                            }
                                        }
                                        var v81 uint16 = genFn0()
                                        v1 = (v1 << ((65535) & 63))
                                        var v82 uint = 255
                                        v82 -= (v82 & (v82 * v82))
                                        _ = v81
                                        _ = v82
                                    } else {
                                        v5 += v3
                                        var v83 int8 = 127
                                        fmt.Println(v6, v4)
                                        _ = v83
                                    }
                                    if len(v58) >= 2 {
                                        v58 = v58[1:len(v58)]
                                    }
                                }
                                v84 := GenS11{F0: 1, F1: 61, F2: true}
                                v58 = append(v58, (-(fi47)))
                                var v85 uint64 = 2
                                _ = v84
                                _ = v85
                            case "foo bar":
                                var v86 uint = 1
                                _ = v86
                            }
                        case -1:
                            if v2 {
                                v58 = append(v58, fi47)
                            } else {
                                v2 = (!(!v2))
                            }
                        }
                        _ = v58
                    }
                default:
                    if ((fi48 ^ -483) == (fi47 & fi48)) {
                        v2 = false
                        fmt.Println(v2)
                        var v87 int64 = 993
                        _ = v87
                    } else {
                        v5 = v3
                        v6 = true
                        v1 = 0
                        switch (fi48 | 470) {
                        case 9223372036854775807:
                            fmt.Println(v4, v2)
                        }
                        if v2 {
                            if v6 {
                                fmt.Println(v3, v2, v5)
                                v5 -= v3
                                v88 := GenS9{F0: int64(fi48), F1: 931}
                                _ = v88
                            }
                            v89 := []int{(fi47 % ((fi48) | 1)), -577, 590}
                            var v90 uint32 = 65535
                            _ = v89
                            _ = v90
                        }
                    }
                }
                var v91 uint16 = 255
                if v6 {
                    switch 2147483647 {
                    case -2147483648:
                        var v92 float64 = v3
                        _ = v92
                    default:
                        var v93 string = "hello"
                        _ = v93
                    }
                    v94 := GenS12{F0: v6, F1: 127}
                    v95 := []int{}
                    v96 := GenS13{F0: -128, F1: 127, F2: v91}
                    _ = v94
                    _ = v95
                    _ = v96
                }
                _ = fi48
                _ = v91
            }
            if v6 {
                var v97 uint16 = 65535
                _ = v97
            } else {
                switch "foo bar" {
                case "a":
                    v98 := []int{(fi47 >> ((1) & 63))}
                    _ = v98
                case "\n\t":
                    var v99 bool = v2
                    _ = v99
                case "hello":
                    var v100 int16 = -221
                    _ = v100
                }
                v101 := []int{(^(fi47)), (^(fi47))}
                _ = v101
            }
            _ = fi47
        }
        _ = v46
    }
    v102 := GenS14{F0: 45, F1: -9223372036854775808, F2: -399}
    v103 := GenS15{F0: 0, F1: int8(v1)}
    if v6 {
        v104 := GenS4{F0: -559, F1: v1}
        _ = v104
    }
    v105 := GenS6{F0: 42, F1: -9223372036854775808, F2: 1}
    fmt.Println(v4)
    var v106 uint32 = 0
    fmt.Println(v5, v2)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v6)
    }
    if true {
        var v107 int16 = 1
        for fi108 := 0; fi108 < 8; fi108++ {
            if v6 {
                v1 = (v106 * v106)
                v109 := GenS10{F0: 2, F1: v1}
                _ = v109
            }
            if (((0 >= 0) || (86 != 1)) && (!(2 >= 42))) {
                v110 := []int{(fi108 - 128), fi108, 595}
                var v111 uint32 = 2147483648
                _ = v110
                _ = v111
            }
            if ((42 == 42) && false) {
                v5 += -2.25
                v112 := GenS16{F0: 2, F1: 834}
                v107 += v107
                _ = v112
            } else {
                var v113 int = (fi108 + (fi108 ^ (fi108 / ((fi108) | 1))))
                var v114 int8 = 100
                if (float32((v3 * float64(0.1))) >= 1e10) {
                    var v115 int = (v113 % (((fi108 / (((-(v113))) | 1))) | 1))
                    if (0 > (v1 >> ((0) & 63))) {
                        v114 = (v114 % ((1) | 1))
                        v115 += -9223372036854775808
                        if v2 {
                            var v116 uint64 = 1
                            _ = v116
                        }
                    }
                    _ = v115
                }
                _ = v113
                _ = v114
            }
            _ = fi108
        }
        _ = v107
    }
    var v117 string = "a"
    v118 := GenS17{F0: 42}
    fmt.Println(v1, v2)
    for fi119 := 0; fi119 < 14; fi119++ {
        v106 -= v1
        fmt.Println(v106)
        var v120 float64 = (v4 - 1e10)
        _ = fi119
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v105
    _ = v106
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
