package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS10 struct { F0 uint64; F1 bool; F2 int }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 uint16; F1 bool }
type GenS14 struct { F0 bool }

func main() {
    var v1 int64 = -1
    var v2 int8 = 127
    var v3 int32 = 2147483647
    var v4 int64 = (-(v1))
    v1 += (v4 * v4)
    var v5 uint16 = 579
    var v6 bool = true
    for fi7 := 0; fi7 < 5; fi7++ {
        switch -9223372036854775808 {
        case 186:
            for fi8 := 0; fi8 < 16; fi8++ {
                switch fi8 {
                case 4294967295:
                    var v9 uint8 = 2
                    _ = v9
                case -599:
                    v10 := GenS1{F0: 29}
                    _ = v10
                case -428:
                    v4 += v1
                    v11 := GenS1{F0: 255}
                    var v12 uint16 = (^(v5))
                    switch ("hello" + ("hello" + "hello")) {
                    case "hello":
                        var v13 bool = (v6 && (42 >= 694))
                        if (3.14159 == -2.25) {
                            if ((fi7 / ((fi8) | 1)) <= fi7) {
                                v14 := []int{}
                                var v15 int8 = v2
                                _ = v14
                                _ = v15
                            }
                        }
                        _ = v13
                    default:
                        if (v4 == v1) {
                            fmt.Println(v6, v12)
                            if v6 {
                                v16 := GenS2{F0: v5}
                                switch (fi7 + (fi7 << ((65535) & 63))) {
                                case 0:
                                    v17 := []int{}
                                    _ = v17
                                case 127:
                                    v18 := GenS1{F0: 166}
                                    _ = v18
                                default:
                                    v5 = v12
                                }
                                _ = v16
                            }
                        } else {
                            fmt.Println(v12, v1, fi8)
                            v12 -= 647
                        }
                        v19 := []int{(fi8 % ((2147483647) | 1)), (fi8 | (fi7 & fi8)), 813}
                        var v20 float64 = (3.14159 - (0.0 - 1e10))
                        _ = v19
                        _ = v20
                    }
                    _ = v11
                    _ = v12
                }
                _ = fi8
            }
        case 4294967295:
            var v21 uint32 = 0
            _ = v21
        case -1:
            var v22 int8 = (v2 % (((v2 ^ (v2 >> ((65535) & 63)))) | 1))
            _ = v22
        }
        if (v5 < 2) {
            v6 = true
            var v23 int8 = (-(v2))
            if ((v3 ^ v3) != (^(v3))) {
                v1 = -1
            }
            _ = v23
        } else {
            var v24 int64 = (v4 % (((v4 | (^(v1)))) | 1))
            _ = v24
        }
        if ((0.1 / -2.25) == (-2.25 / ((1.5 / (3.14159 / -2.25)) * -2.25))) {
            switch 127 {
            case 32767:
                for fi25 := 0; fi25 < 13; fi25++ {
                    if false {
                        switch ("a" + ("\n\t" + "")) {
                        case "":
                            v26 := GenS1{F0: 2}
                            _ = v26
                        case "a":
                            v2 -= 79
                            var v27 int8 = v2
                            _ = v27
                        case "foo bar":
                            var v28 bool = true
                            _ = v28
                        }
                    }
                    if (4294967295 != 605) {
                        v29 := GenS3{F0: (v4 - (v4 & v1))}
                        if (("foo bar" > "foo bar") && (-21 != (v2 << ((1) & 63)))) {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    v30 := GenS4{F0: (v5 ^ v5), F1: v5, F2: v4}
                                    if (v6 || true) {
                                        var v31 int16 = 32767
                                        var v32 uint = 497
                                        fmt.Println(v6, fi25, v2)
                                        _ = v31
                                        _ = v32
                                    } else {
                                        if false {
                                            v33 := GenS5{F0: "a"}
                                            fmt.Println(fi7, fi25)
                                            _ = v33
                                        }
                                        if ("foo bar" > ("hello" + "foo bar")) {
                                            v34 := []int{fi25, -9223372036854775808, 992, (fi7 * fi7)}
                                            v2 = (v2 + int8(v4))
                                            v35 := GenS1{F0: 127}
                                            var v36 uint16 = 865
                                            var v37 float64 = (float64(float32(-2.25)) * float64(-2.25))
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        } else {
                                            v38 := []int{(fi7 | (fi7 / ((fi7) | 1)))}
                                            _ = v38
                                        }
                                    }
                                    var v39 bool = v6
                                    _ = v30
                                    _ = v39
                                }
                            default:
                                var v40 int64 = (^(v4))
                                _ = v40
                            }
                            fmt.Println(fi25, v6)
                            v41 := GenS1{F0: 91}
                            _ = v41
                        }
                        _ = v29
                    } else {
                        v42 := GenS2{F0: v5}
                        var v43 int32 = v3
                        _ = v42
                        _ = v43
                    }
                    v44 := []int{(-(fi7)), (-(fi25))}
                    _ = fi25
                    _ = v44
                }
            case -1:
                fmt.Println(v3)
            case -950:
                if (v6 && v6) {
                    var v45 int8 = (v2 >> ((4294967295) & 63))
                    _ = v45
                }
            }
            v3 = v3
        } else {
            for fi46 := 0; fi46 < 20; fi46++ {
                var v47 uint32 = 320
                v5 -= v5
                fmt.Println(v5, v47, v3)
                _ = fi46
                _ = v47
            }
            var v48 uint = 65535
            v2 = v2
            _ = v48
        }
        _ = fi7
    }
    v6 = (((v5 & (v5 ^ 855)) >= 603) || (!(v6 || v6)))
    var v49 uint64 = 1
    if (1e10 >= float32(0.0)) {
        v50 := GenS6{F0: 42, F1: 341}
        _ = v50
    }
    v51 := []int{-9223372036854775808, 1, 981}
    var v52 int32 = 327
    if ((-102 <= 127) && (((v6 || (int16(v49) < -32768)) || (((v1 != 554) && (!(v6 || ((v4 & v4) < v4)))) || (!v6))) && false)) {
        switch 2147483647 {
        case 679:
            switch 0 {
            case 2147483647:
                v3 = v52
            case 247:
                var v53 int8 = (v2 & (^(v2)))
                _ = v53
            }
        }
        switch 9223372036854775807 {
        case 260:
            for fi54 := 0; fi54 < 13; fi54++ {
                for fi55 := 0; fi55 < 6; fi55++ {
                    fmt.Println(v2, v6, v52)
                    var v56 uint32 = 278
                    _ = fi55
                    _ = v56
                }
                for fi57 := 0; fi57 < 19; fi57++ {
                    v6 = (!v6)
                    var v58 uint8 = 42
                    v6 = v6
                    _ = fi57
                    _ = v58
                }
                _ = fi54
            }
        case -241:
            switch 0 {
            case -2147483648:
                switch "hello" {
                case "hello":
                    v51 = append(v51, 0)
                case "foo bar":
                    if v6 {
                        fmt.Println(v52, v3)
                    }
                }
            case 32768:
                v1 -= v4
            }
        default:
            v59 := GenS7{F0: v5, F1: 32768}
            _ = v59
        }
        if (!(1e10 >= 1.5)) {
            switch 1 {
            case 1:
                var v60 uint16 = 2
                _ = v60
            case 32768:
                for fi61 := 0; fi61 < 7; fi61++ {
                    if (3.14159 > (3.14159 + -2.25)) {
                        var v62 uint8 = 0
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 17; fi63++ {
                            if len(v51) >= 2 {
                                v51 = v51[1:len(v51)]
                            }
                            v5 -= v5
                            _ = fi63
                        }
                    }
                    _ = fi61
                }
            case -831:
                for fi64 := 0; fi64 < 7; fi64++ {
                    v4 += v1
                    fmt.Println(v2, v5)
                    _ = fi64
                }
            }
            if (1 != 4294967295) {
                v6 = (v6 || v6)
                fmt.Println(v52, v3, v6)
                if false {
                    for fi65 := 0; fi65 < 15; fi65++ {
                        if ((^(v49)) == (^(v49))) {
                            if (!v6) {
                                if v6 {
                                    switch (fi65 & 4294967295) {
                                    case 9223372036854775807:
                                        var v66 uint8 = 53
                                        _ = v66
                                    default:
                                        for fi67 := 0; fi67 < 6; fi67++ {
                                            if len(v51) >= 2 {
                                                v51 = v51[1:len(v51)]
                                            }
                                            _ = fi67
                                        }
                                    }
                                    v3 = (v3 - (v3 >> ((961) & 63)))
                                } else {
                                    switch fi65 {
                                    case 122:
                                        for fi68 := 0; fi68 < 9; fi68++ {
                                            v51 = append(v51, (^(fi68)))
                                            fmt.Println(fi68)
                                            _ = fi68
                                        }
                                    case -709:
                                        v69 := GenS8{F0: false, F1: 871}
                                        _ = v69
                                    default:
                                        if ((537 <= (^(v5))) || (!(v6 || v6))) {
                                            v51 = append(v51, fi65)
                                            var v70 int16 = 32767
                                            if v6 {
                                                if false {
                                                    var v71 uint64 = v49
                                                    _ = v71
                                                }
                                            } else {
                                                for fi72 := 0; fi72 < 11; fi72++ {
                                                    var v73 string = "a"
                                                    var v74 float32 = 0.1
                                                    v75 := GenS9{F0: 42, F1: (v5 + v5), F2: 127}
                                                    v76 := GenS10{F0: 337, F1: false, F2: (fi65 & (fi72 << ((4294967295) & 63)))}
                                                    v77 := GenS4{F0: (v5 ^ v5), F1: (v5 * 1), F2: v4}
                                                    if ((v2 | v2) < (v2 & 42)) {
                                                        v74 = 0.1
                                                        if (!((fi72 % ((int(v4)) | 1)) <= (fi65 / (((fi72 << ((4294967295) & 63))) | 1)))) {
                                                            switch (fi72 + (-(fi72))) {
                                                            case 1:
                                                                v6 = v6
                                                            case 128:
                                                                switch (fi65 | fi65) {
                                                                case -605:
                                                                    if len(v51) >= 2 {
                                                                        v51 = v51[1:len(v51)]
                                                                    }
                                                                }
                                                            default:
                                                                v4 += v1
                                                                _ = len(v51)
                                                            }
                                                            v51 = append(v51, fi65)
                                                        } else {
                                                            var v78 int16 = 1000
                                                            _ = v78
                                                        }
                                                    } else {
                                                        switch (-(fi65)) {
                                                        case 127:
                                                            if len(v51) >= 2 {
                                                                v51 = v51[1:len(v51)]
                                                            }
                                                        case -2147483648:
                                                            if true {
                                                                v1 += v1
                                                            }
                                                        case -986:
                                                            if len(v51) >= 2 {
                                                                v51 = v51[1:len(v51)]
                                                            }
                                                        default:
                                                            v3 += (v52 % (((v52 % ((v3) | 1))) | 1))
                                                            fmt.Println(v4)
                                                            v79 := GenS11{F0: (fi65 - (-(fi72)))}
                                                            if len(v51) > 0 {
                                                                v51[0] = fi72
                                                            }
                                                            _ = len(v51)
                                                            var v80 float64 = ((1e10 + (1e10 + 3.14159)) * 0.1)
                                                            _ = v79
                                                            _ = v80
                                                        }
                                                        var v81 int64 = 923
                                                        _ = v81
                                                    }
                                                    _ = fi72
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                }
                                                fmt.Println(v6, v49, v70)
                                                var v82 int16 = v70
                                                _ = v82
                                            }
                                            _ = v70
                                        } else {
                                            _ = len(v51)
                                            if true {
                                                fmt.Println(v4)
                                                var v83 int = (-(fi65))
                                                if (true && (v6 && (1e10 > 1.5))) {
                                                    if v6 {
                                                        if v6 {
                                                            v84 := GenS12{F0: 0, F1: v1}
                                                            var v85 uint8 = 72
                                                            if (!((345 >= 1) || false)) {
                                                                v86 := GenS13{F0: v5, F1: v6}
                                                                for fi87 := 0; fi87 < 19; fi87++ {
                                                                    v4 += (v1 % (((v4 / (((-(v1))) | 1))) | 1))
                                                                    v83 += (fi87 >> ((42) & 63))
                                                                    var v88 bool = (((-2.25 - 0.1) + 1.5) != 0.1)
                                                                    _ = fi87
                                                                    _ = v88
                                                                }
                                                                fmt.Println(v5)
                                                                _ = v86
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                if v6 {
                                                                    if false {
                                                                        v89 := GenS1{F0: 0}
                                                                        v90 := GenS14{F0: v6}
                                                                        fmt.Println(v49)
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                    fmt.Println(v6)
                                                                    v5 += v5
                                                                } else {
                                                                    var v91 float64 = 1e10
                                                                    _ = v91
                                                                }
                                                            default:
                                                                fmt.Println(v1, v5)
                                                            }
                                                            if len(v51) >= 2 {
                                                                v51 = v51[1:len(v51)]
                                                            }
                                                        }
                                                        for fi92 := 0; fi92 < 8; fi92++ {
                                                            v93 := GenS1{F0: 229}
                                                            var v94 float32 = -2.25
                                                            _ = fi92
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    } else {
                                                        switch (v83 << ((18446744073709551615) & 63)) {
                                                        case 127:
                                                            var v95 uint8 = 219
                                                            _ = v95
                                                        default:
                                                            var v96 string = "foo bar"
                                                            _ = v96
                                                        }
                                                    }
                                                } else {
                                                    for fi97 := 0; fi97 < 18; fi97++ {
                                                        if (2147483648 >= 0) {
                                                            fmt.Println(v3, v4)
                                                            var v98 int64 = (v4 * 246)
                                                            _ = v98
                                                        }
                                                        _ = fi97
                                                    }
                                                    for fi99 := 0; fi99 < 5; fi99++ {
                                                        var v100 uint16 = v5
                                                        v52 = -2147483648
                                                        switch fi99 {
                                                        case 2147483647:
                                                            v6 = (472 < 1)
                                                        case -402:
                                                            var v101 uint64 = v49
                                                            _ = v101
                                                        case 1:
                                                            switch (fi99 >> ((0) & 63)) {
                                                            case 128:
                                                                v52 = (v52 / (((v3 << ((uint(v100)) & 63))) | 1))
                                                            case 4294967295:
                                                                if len(v51) > 0 {
                                                                    v51[0] = 127
                                                                }
                                                            }
                                                        }
                                                        _ = fi99
                                                        _ = v100
                                                    }
                                                    v83 = (fi65 % (((v83 >> ((2) & 63))) | 1))
                                                }
                                                _ = v83
                                            }
                                            for fi102 := 0; fi102 < 15; fi102++ {
                                                v103 := GenS6{F0: 0, F1: 4294967295}
                                                _ = fi102
                                                _ = v103
                                            }
                                        }
                                    }
                                    var v104 uint64 = uint64(v5)
                                    switch 4294967295 {
                                    case 127:
                                        fmt.Println(v3, v104)
                                    }
                                    _ = v104
                                }
                                var v105 float64 = -2.25
                                switch int(v3) {
                                case 128:
                                    if (-9223372036854775808 > v1) {
                                        switch "" {
                                        case "\n\t":
                                            for fi106 := 0; fi106 < 15; fi106++ {
                                                v51 = append(v51, (fi106 - (fi106 << ((42) & 63))))
                                                _ = fi106
                                            }
                                        }
                                    } else {
                                        v105 -= v105
                                        for fi107 := 0; fi107 < 18; fi107++ {
                                            var v108 uint16 = (v5 + (v5 * v5))
                                            if len(v51) > 0 {
                                                v51[0] = -416
                                            }
                                            fmt.Println(v1, v2, v49)
                                            switch ((("hello" + "a") + "\n\t") + "hello") {
                                            case "":
                                                var v109 int32 = (v3 - v3)
                                                _ = v109
                                            case "\n\t":
                                                if (!(!((42 > 952) || v6))) {
                                                    if (v6 || false) {
                                                        var v110 int = -632
                                                        _ = v110
                                                    }
                                                } else {
                                                    var v111 string = (("a" + "hello") + "hello")
                                                    v112 := GenS13{F0: 1, F1: ((((fi65 | fi65) == (fi107 * (fi107 << ((1) & 63)))) || v6) && ((690 > -734) || (v6 && v6)))}
                                                    _ = v111
                                                    _ = v112
                                                }
                                            default:
                                                if v6 {
                                                    switch (("\n\t" + ("" + "a")) + (("hello" + "foo bar") + ("a" + ("hello" + ("hello" + "\n\t"))))) {
                                                    case "a":
                                                        var v113 uint16 = v108
                                                        _ = v113
                                                    case "foo bar":
                                                        fmt.Println(v2, v49)
                                                    default:
                                                        var v114 bool = (v6 && true)
                                                        _ = v114
                                                    }
                                                    v115 := GenS2{F0: (v108 << ((80) & 63))}
                                                    v1 += v1
                                                    _ = v115
                                                } else {
                                                    var v116 string = ("foo bar" + ("" + (("" + "\n\t") + "foo bar")))
                                                    var v117 int = (fi65 + -797)
                                                    if v6 {
                                                        _ = len(v51)
                                                        v105 = (v105 - (float64(3.14159) / (1e10 + (((1.5 - (float64(1e10) * (3.14159 - v105))) / v105) * v105))))
                                                        v4 = (v1 + int64(v2))
                                                    }
                                                    v108 = (v5 + (v5 % (((v5 & (v5 << ((18446744073709551615) & 63)))) | 1)))
                                                    var v118 uint16 = v5
                                                    _ = v116
                                                    _ = v117
                                                    _ = v118
                                                }
                                            }
                                            _ = fi107
                                            _ = v108
                                        }
                                        v4 -= 0
                                    }
                                case 9223372036854775807:
                                    fmt.Println(v3, v3, v2)
                                default:
                                    switch 128 {
                                    case 128:
                                        var v119 string = ("a" + (("" + (("foo bar" + "foo bar") + "foo bar")) + ((("foo bar" + "foo bar") + "foo bar") + "")))
                                        _ = v119
                                    case 42:
                                        v4 += -9223372036854775808
                                    case 0:
                                        _ = len(v51)
                                    default:
                                        v1 = -9223372036854775808
                                    }
                                }
                                _ = v105
                            }
                            for fi120 := 0; fi120 < 5; fi120++ {
                                fmt.Println(v6)
                                _ = fi120
                            }
                            var v121 uint32 = 0
                            _ = v121
                        }
                        _ = fi65
                    }
                }
            }
            v3 = (v52 << ((4294967295) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
