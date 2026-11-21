package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 int64; F1 int16; F2 uint16 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint8; F1 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int8; F1 int8 }
type GenS15 struct { F0 uint32; F1 uint32; F2 uint }
func genFn0(p0 uint, p1 int32) int16 {
    v7 := []int{}
    var v8 int8 = 20
    var v9 bool = (458 == -405)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1() uint16 {
    v10 := GenS1{F0: -755}
    var v11 int16 = 1000
    var v12 uint = 2
    _ = v10
    _ = v11
    _ = v12
    return 42
}

func main() {
    v1 := []int{9223372036854775807, 32768, 32768, 233}
    var v2 uint16 = 0
    var v3 int16 = 902
    v4 := []int{}
    var v5 string = "foo bar"
    v6 := make(map[string]string)
    var v13 uint32 = 4294967295
    var v14 uint8 = 127
    switch (("foo bar" + "\n\t") + (v5 + "foo bar")) {
    case "":
        if (3.14159 < -2.25) {
            for fi15 := 0; fi15 < 15; fi15++ {
                fmt.Println(fi15, v13, v2)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                switch fi15 {
                case 0:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32768:
                    fmt.Println(v3)
                }
                _ = fi15
            }
            var v16 uint = 65535
            _ = v16
        }
    case "hello":
        fmt.Println(v13)
    case "foo bar":
        v17 := GenS1{F0: 951}
        _ = v17
    default:
        var v18 int = 42
        _ = v18
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        var v20 int64 = 1
        if len(v1) > 0 {
            v1[0] = fi19
        }
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        v22 := GenS2{F0: v3, F1: 18446744073709551615}
        if false {
            v13 = v13
        } else {
            fmt.Println(v13, v5)
            v2 = v2
        }
        var v23 uint16 = (v2 % (((v2 - v2)) | 1))
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(v14)
        _ = fi24
    }
    v13 = v13
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch ((v5 + v5) + ((v5 + "foo bar") + v5)) {
    case "\n\t":
        var v25 float32 = 1.5
        _ = v25
    default:
        v26 := v6[v5]
        _ = v26
    }
    var v27 bool = (true || false)
    v28 := GenS3{F0: (-(v3))}
    switch ("a" + v5) {
    case "a":
        v27 = v27
    case "hello":
        if false {
            v29 := GenS3{F0: v3}
            switch "a" {
            case "hello":
                if true {
                    if v27 {
                        if (1.5 != 0.1) {
                            var v30 string = (((v5 + "\n\t") + ("foo bar" + v5)) + ("" + ("\n\t" + "hello")))
                            var v31 uint = 2
                            _ = v30
                            _ = v31
                        }
                        v32 := GenS1{F0: 0}
                        _ = v32
                    } else {
                        var v33 uint64 = 4294967295
                        v34 := GenS2{F0: genFn0(0, -742), F1: 65535}
                        _ = v33
                        _ = v34
                    }
                    var v35 int8 = 42
                    var v36 int64 = -619
                    _ = v35
                    _ = v36
                } else {
                    switch 4294967295 {
                    case 293:
                        fmt.Println(v2, v13, v3)
                    case 519:
                        if false {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                v38 := GenS3{F0: v3}
                                if (1e10 > float32((1.5 * 1.5))) {
                                    v39 := GenS2{F0: genFn0(1, 1), F1: 500}
                                    var v40 int = fi37
                                    v6["foo bar"] = v5
                                    _ = v39
                                    _ = v40
                                } else {
                                    for fi41 := 0; fi41 < 4; fi41++ {
                                        if (v27 && true) {
                                            v6[v5] = ((v5 + v5) + ("hello" + v5))
                                        } else {
                                            if v27 {
                                                var v42 uint32 = (v13 * (v13 + v13))
                                                _ = len(v6)
                                                _ = v42
                                            }
                                            var v43 float32 = 1.5
                                            v5 = (v5 + "")
                                            var v44 uint64 = 431
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = fi41
                                    }
                                    v27 = (42 != 2)
                                    switch (v5 + v5) {
                                    case "\n\t":
                                        for fi45 := 0; fi45 < 14; fi45++ {
                                            fmt.Println(v3, v5, v27)
                                            if true {
                                                if (533 <= 4294967295) {
                                                    if (0 == 570) {
                                                        _ = len(v6)
                                                        v13 += (^(v13))
                                                    }
                                                    if (v27 || v27) {
                                                        if true {
                                                            switch fi45 {
                                                            case 782:
                                                                v46 := GenS1{F0: 1}
                                                                _ = v46
                                                            case 214:
                                                                var v47 float32 = (1e10 / (1.5 * (3.14159 - (-2.25 - 0.0))))
                                                                _ = v47
                                                            case -2147483648:
                                                                if v27 {
                                                                    v48 := GenS2{F0: 42, F1: 2}
                                                                    v49 := GenS4{F0: 9223372036854775807, F1: (true || ("hello" < "a")), F2: v27}
                                                                    var v50 int32 = 2147483647
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            }
                                                            v6["\n\t"] = "foo bar"
                                                        }
                                                        var v51 uint = 4294967295
                                                        _ = v51
                                                    }
                                                }
                                                var v52 int8 = 42
                                                var v53 string = v5
                                                _ = v52
                                                _ = v53
                                            } else {
                                                switch (fi37 % (((fi45 - fi37)) | 1)) {
                                                case 2147483647:
                                                    v13 = v13
                                                case 4294967295:
                                                    var v54 bool = (2147483647 != 1)
                                                    _ = v54
                                                }
                                            }
                                            _ = fi45
                                        }
                                    case "foo bar":
                                        for fi55 := 0; fi55 < 15; fi55++ {
                                            fmt.Println(v5, v27, v27)
                                            if ((2 <= uint(v14)) || (127 < 1)) {
                                                v56 := GenS4{F0: -1, F1: v27, F2: false}
                                                _ = v56
                                            }
                                            v27 = false
                                            _ = fi55
                                        }
                                    }
                                }
                                v13 += v13
                                _ = fi37
                                _ = v38
                            }
                        } else {
                            switch "" {
                            case "a":
                                v5 = v5
                            }
                            delete(v6, ((v5 + v5) + "\n\t"))
                        }
                    }
                    v3 -= (-(v3))
                }
            case "a":
                v13 = v13
            default:
                for fi57 := 0; fi57 < 14; fi57++ {
                    if (v2 > (v2 ^ v2)) {
                        v2 = (^(v2))
                        switch "a" {
                        case "hello":
                            fmt.Println(v2, v14)
                        case "\n\t":
                            var v58 float64 = 1e10
                            _ = v58
                        }
                        switch ((v5 + "hello") + v5) {
                        case "a":
                            v59 := GenS5{F0: -9223372036854775808, F1: v3, F2: (^(v2))}
                            _ = v59
                        default:
                            if false {
                                if false {
                                    var v60 float32 = 3.14159
                                    fmt.Println(fi57, v60, v2)
                                    if (v13 == 2147483648) {
                                        var v61 uint = 0
                                        _ = v61
                                    }
                                    _ = v60
                                } else {
                                    var v62 int16 = (v3 >> ((255) & 63))
                                    v2 -= genFn1()
                                    v14 = v14
                                    _ = v62
                                }
                            } else {
                                switch -9223372036854775808 {
                                case -2147483648:
                                    v63 := GenS6{F0: fi57, F1: v5}
                                    _ = v63
                                case 32768:
                                    switch v5 {
                                    case "a":
                                        _ = len(v1)
                                    case "foo bar":
                                        var v64 uint = 877
                                        _ = v64
                                    default:
                                        for fi65 := 0; fi65 < 10; fi65++ {
                                            var v66 uint64 = 4294967295
                                            v4 = append(v4, (fi65 % (((fi57 | fi57)) | 1)))
                                            if ((v2 ^ (v2 % ((genFn1()) | 1))) > 899) {
                                                v2 = v2
                                                var v67 uint16 = (v2 | uint16(v13))
                                                _ = v67
                                            }
                                            _ = fi65
                                            _ = v66
                                        }
                                    }
                                case 0:
                                    for fi68 := 0; fi68 < 5; fi68++ {
                                        var v69 string = (v5 + (v5 + v5))
                                        _ = len(v4)
                                        if (!(15 < -9223372036854775808)) {
                                            v2 = (v2 & v2)
                                            var v70 uint = 4294967295
                                            if true {
                                                _ = len(v4)
                                                var v71 uint = (v70 & v70)
                                                if v27 {
                                                    if len(v4) > 0 {
                                                        v4[0] = fi57
                                                    }
                                                    v72 := GenS7{F0: v71}
                                                    if v27 {
                                                        var v73 int16 = genFn0(v70, -2147483648)
                                                        _ = v73
                                                    } else {
                                                        v70 = v70
                                                        var v74 float32 = 3.14159
                                                        v27 = v27
                                                        _ = v74
                                                    }
                                                    _ = v72
                                                }
                                                _ = v71
                                            } else {
                                                if v27 {
                                                    switch v5 {
                                                    case "\n\t":
                                                        var v75 float32 = 0.1
                                                        if false {
                                                            v2 -= v2
                                                            var v76 uint32 = 42
                                                            v77 := GenS8{F0: (v3 >> (((v70 % ((uint(v14)) | 1))) & 63)), F1: 42}
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            v78 := GenS2{F0: (v3 | -32768), F1: 65535}
                                                            if v27 {
                                                                if (v27 || v27) {
                                                                    if v27 {
                                                                        v75 = (v75 * (v75 / 1e10))
                                                                        var v79 string = v69
                                                                        _ = v79
                                                                    }
                                                                    var v80 uint8 = (v14 ^ (v14 + (v14 + (v14 << ((v70) & 63)))))
                                                                    _ = v80
                                                                } else {
                                                                    var v81 uint = (v70 & (^(v70)))
                                                                    var v82 int = fi57
                                                                    if (!(v27 || true)) {
                                                                        v70 = v81
                                                                        if (uint8(v2) >= (v14 | 193)) {
                                                                            v3 += v3
                                                                            if v27 {
                                                                                v6[v69] = v5
                                                                                var v83 int64 = 936
                                                                                fmt.Println(v83, v81, fi57)
                                                                                v84 := GenS5{F0: 0, F1: v3, F2: genFn1()}
                                                                                _ = v83
                                                                                _ = v84
                                                                            } else {
                                                                                fmt.Println(v81, v81)
                                                                                var v85 float32 = float32(1e10)
                                                                                if (!false) {
                                                                                    if v27 {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                    switch (v82 >> ((v81) & 63)) {
                                                                                    case -9223372036854775808:
                                                                                        switch (^(v82)) {
                                                                                        case -762:
                                                                                            v82 = fi57
                                                                                        }
                                                                                    }
                                                                                    var v86 uint = v81
                                                                                    if (1e10 <= 1.5) {
                                                                                        if true {
                                                                                            v85 = (1e10 - -2.25)
                                                                                            v13 -= (v13 % (((^(v13))) | 1))
                                                                                        } else {
                                                                                            fmt.Println(v3)
                                                                                            v6[v5] = v5
                                                                                            if ((v3 / ((genFn0((^(v81)), -2147483648)) | 1)) <= (v3 >> (((^(v86))) & 63))) {
                                                                                                var v87 uint = (v70 & (v70 ^ (v70 % ((v70) | 1))))
                                                                                                v5 = ((v5 + "\n\t") + (v69 + v69))
                                                                                                v82 = fi57
                                                                                                _ = v87
                                                                                            } else {
                                                                                                var v88 int = 2147483647
                                                                                                if (!(!(!(true && false)))) {
                                                                                                    var v89 uint64 = 1
                                                                                                    v90 := GenS9{F0: 909, F1: 0}
                                                                                                    var v91 uint = (v81 % ((v81) | 1))
                                                                                                    _ = v89
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                } else {
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = 213
                                                                                                    }
                                                                                                    v70 = (v86 * 18446744073709551615)
                                                                                                    if true {
                                                                                                        fmt.Println(v70, v81)
                                                                                                    }
                                                                                                }
                                                                                                var v92 int32 = 441
                                                                                                _ = v88
                                                                                                _ = v92
                                                                                            }
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        if v27 {
                                                                                            if true {
                                                                                                v75 += v75
                                                                                                fmt.Println(v14)
                                                                                                switch "foo bar" {
                                                                                                case "\n\t":
                                                                                                    var v93 int32 = -2147483648
                                                                                                    _ = v93
                                                                                                case "":
                                                                                                    switch "foo bar" {
                                                                                                    case "\n\t":
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    case "":
                                                                                                        switch fi57 {
                                                                                                        case -9223372036854775808:
                                                                                                            v27 = false
                                                                                                        default:
                                                                                                            fmt.Println(v14, v82)
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    var v94 int = (v82 << (((v70 / (((v86 / ((v86) | 1))) | 1))) & 63))
                                                                                                    _ = v94
                                                                                                }
                                                                                            } else {
                                                                                                v4 = append(v4, 32767)
                                                                                            }
                                                                                            if ((v75 + 3.14159) <= v85) {
                                                                                                if (42 >= -1) {
                                                                                                    v27 = v27
                                                                                                }
                                                                                                v86 = (^(v70))
                                                                                            }
                                                                                            v95 := GenS10{F0: v14, F1: int(v81)}
                                                                                            _ = v95
                                                                                        } else {
                                                                                            switch ((v5 + v5) + ("hello" + (v5 + v5))) {
                                                                                            case "foo bar":
                                                                                                var v96 int = fi57
                                                                                                _ = v96
                                                                                            case "hello":
                                                                                                var v97 int64 = 327
                                                                                                _ = v97
                                                                                            case "a":
                                                                                                if v27 {
                                                                                                    v86 -= (v70 << (((v81 << (((v70 >> ((v81) & 63))) & 63))) & 63))
                                                                                                } else {
                                                                                                    var v98 int8 = 1
                                                                                                    _ = v98
                                                                                                }
                                                                                            default:
                                                                                                v86 -= 18446744073709551615
                                                                                            }
                                                                                            var v99 uint8 = v14
                                                                                            if v27 {
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    var v100 int16 = v3
                                                                                                    _ = v100
                                                                                                case "":
                                                                                                    var v101 uint = (v81 % ((632) | 1))
                                                                                                    _ = v101
                                                                                                case "a":
                                                                                                    v86 = uint(v2)
                                                                                                default:
                                                                                                    v102 := GenS11{F0: "foo bar"}
                                                                                                    _ = v102
                                                                                                }
                                                                                                v4 = append(v4, fi57)
                                                                                                v103 := GenS5{F0: -9223372036854775808, F1: v3, F2: genFn1()}
                                                                                                _ = v103
                                                                                            }
                                                                                            var v104 uint32 = (^(v13))
                                                                                            var v105 uint16 = (^(v2))
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                            _ = v99
                                                                                        }
                                                                                        if (v13 != (v13 + (v13 + (v13 * (^(v13)))))) {
                                                                                            if v27 {
                                                                                                fmt.Println(v27)
                                                                                                delete(v6, v69)
                                                                                                v5 = v69
                                                                                            }
                                                                                        } else {
                                                                                            v81 -= (v86 | 4294967295)
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v86)
                                                                                    _ = v86
                                                                                }
                                                                                fmt.Println(v82)
                                                                                _ = v85
                                                                            }
                                                                            switch 128 {
                                                                            case 927:
                                                                                if v27 {
                                                                                    var v106 uint64 = 18446744073709551615
                                                                                    var v107 uint = v81
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                } else {
                                                                                    v108 := GenS12{F0: -1, F1: (v13 & (v13 | 390)), F2: 18446744073709551615}
                                                                                    _ = v108
                                                                                }
                                                                            case 42:
                                                                                v1 = append(v1, int(v3))
                                                                                v6[(v69 + (v5 + v5))] = v5
                                                                                if (-1 <= 42) {
                                                                                    v69 = ""
                                                                                    if false {
                                                                                        var v109 uint8 = v14
                                                                                        var v110 uint8 = v109
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    }
                                                                                    delete(v6, v5)
                                                                                } else {
                                                                                    var v111 uint64 = 2
                                                                                    v112 := GenS11{F0: "\n\t"}
                                                                                    _ = v111
                                                                                    _ = v112
                                                                                }
                                                                                if false {
                                                                                    var v113 bool = (!v27)
                                                                                    _ = v113
                                                                                }
                                                                            case 1:
                                                                                if true {
                                                                                    if ((v70 >> ((v70) & 63)) == (v70 - v81)) {
                                                                                        switch fi57 {
                                                                                        case -333:
                                                                                            v114 := GenS1{F0: 1}
                                                                                            _ = v114
                                                                                        case 509:
                                                                                            delete(v6, "a")
                                                                                        default:
                                                                                            if v27 {
                                                                                                v4 = append(v4, fi68)
                                                                                                var v115 float32 = v75
                                                                                                switch -1 {
                                                                                                case 320:
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        v116 := GenS3{F0: 32767}
                                                                                                        _ = v116
                                                                                                    default:
                                                                                                        v75 -= (v115 - float32(float64(float32(((0.0 + (1e10 / 1.0)) - 0.0)))))
                                                                                                        switch ("hello" + v69) {
                                                                                                        case "foo bar":
                                                                                                            var v117 string = v5
                                                                                                            _ = v117
                                                                                                        case "a":
                                                                                                            v3 += (v3 + -1)
                                                                                                        case "hello":
                                                                                                            v69 = (v5 + v5)
                                                                                                        }
                                                                                                        v14 = v14
                                                                                                    }
                                                                                                case 665:
                                                                                                    switch (("hello" + (v5 + v69)) + (((("\n\t" + ("" + "hello")) + ("foo bar" + v5)) + (v5 + v5)) + "\n\t")) {
                                                                                                    case "a":
                                                                                                        if (4294967295 > 521) {
                                                                                                            switch "foo bar" {
                                                                                                            case "\n\t":
                                                                                                                fmt.Println(v3, v115)
                                                                                                            case "hello":
                                                                                                                v118 := GenS5{F0: 9223372036854775807, F1: v3, F2: genFn1()}
                                                                                                                _ = v118
                                                                                                            case "a":
                                                                                                                var v119 bool = (!(-9223372036854775808 >= -1))
                                                                                                                _ = v119
                                                                                                            default:
                                                                                                                fmt.Println(v3, v13, v3)
                                                                                                            }
                                                                                                            switch ((v5 + "\n\t") + v5) {
                                                                                                            case "foo bar":
                                                                                                                v120 := GenS10{F0: 11, F1: (fi57 - (fi68 & (fi68 | (^(fi57)))))}
                                                                                                                var v121 int16 = (v3 + 0)
                                                                                                                if (65535 <= (v70 ^ 452)) {
                                                                                                                    switch "" {
                                                                                                                    case "":
                                                                                                                        switch 128 {
                                                                                                                        case 1:
                                                                                                                            v70 = (^(v70))
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v122 float32 = 3.14159
                                                                                                                        _ = v122
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v123 := GenS8{F0: genFn0(v70, 1), F1: 0}
                                                                                                                    v6[v5] = "\n\t"
                                                                                                                    _ = v123
                                                                                                                }
                                                                                                                _ = v120
                                                                                                                _ = v121
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v124 float64 = 0.0
                                                                                                            v125 := v6["foo bar"]
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            _ = v124
                                                                                                            _ = v125
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v3 += (-(v3))
                                                                                                    }
                                                                                                case 127:
                                                                                                    switch (fi68 >> ((153) & 63)) {
                                                                                                    case 42:
                                                                                                        fmt.Println(v81)
                                                                                                    case 1:
                                                                                                        v126 := GenS13{F0: (!true)}
                                                                                                        _ = v126
                                                                                                    default:
                                                                                                        v27 = (245 >= int32(v82))
                                                                                                    }
                                                                                                default:
                                                                                                    switch v69 {
                                                                                                    case "hello":
                                                                                                        fmt.Println(v69, fi68)
                                                                                                    }
                                                                                                }
                                                                                                _ = v115
                                                                                            } else {
                                                                                                v5 = v69
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                fmt.Println(v70)
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v5 = "foo bar"
                                                                                        if ((^(v14)) > (^(v14))) {
                                                                                            if v27 {
                                                                                                v13 = (v13 * (^(v13)))
                                                                                                v27 = ((1e10 + 0.0) != 3.14159)
                                                                                            }
                                                                                            var v127 bool = v27
                                                                                            v128 := GenS14{F0: 0, F1: 113}
                                                                                            switch ((v69 + v69) + v5) {
                                                                                            case "hello":
                                                                                                fmt.Println(v81, v3, fi68)
                                                                                            }
                                                                                            if (v13 != v13) {
                                                                                                fmt.Println(v75, fi68)
                                                                                                v129 := GenS5{F0: 759, F1: v3, F2: v2}
                                                                                                _ = v129
                                                                                            }
                                                                                            _ = v127
                                                                                            _ = v128
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v130 float32 = ((v75 - float32(0.0)) * float32(float64(((v75 * v75) / 1.0))))
                                                                                    _ = v130
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    var v131 int = (fi57 ^ (v82 << (((v81 << ((v81) & 63))) & 63)))
                                                                    _ = v131
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                switch ("\n\t" + v69) {
                                                                case "hello":
                                                                    v132 := GenS11{F0: ((("foo bar" + (v5 + "")) + "") + (v5 + (("\n\t" + v5) + "foo bar")))}
                                                                    _ = v132
                                                                case "a":
                                                                    fmt.Println(v2)
                                                                }
                                                                v133 := v6["a"]
                                                                _ = v133
                                                            } else {
                                                                var v134 uint32 = v13
                                                                _ = v134
                                                            }
                                                            _ = len(v6)
                                                            fmt.Println(v5, fi68)
                                                            if (9223372036854775807 == 251) {
                                                                switch 191 {
                                                                case -9223372036854775808:
                                                                    fmt.Println(v14, v75, v69)
                                                                }
                                                                v135 := GenS3{F0: (v3 - 1)}
                                                                switch (fi57 << (((v70 % ((v70) | 1))) & 63)) {
                                                                case 42:
                                                                    fmt.Println(v69, v69, v27)
                                                                case 457:
                                                                    if (-2.25 < 1.5) {
                                                                        switch (fi57 / ((fi57) | 1)) {
                                                                        case 226:
                                                                            v136 := GenS1{F0: 360}
                                                                            _ = v136
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        var v137 uint = (v70 & v70)
                                                                        _ = v137
                                                                    }
                                                                }
                                                                _ = v135
                                                            } else {
                                                                fmt.Println(v13)
                                                            }
                                                            _ = v78
                                                        }
                                                        _ = v75
                                                    }
                                                    v69 = v69
                                                }
                                                _ = len(v6)
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            _ = v70
                                        }
                                        fmt.Println(v13, v5)
                                        v138 := GenS9{F0: -2147483648, F1: 879}
                                        if (!v27) {
                                            var v139 uint8 = v14
                                            _ = len(v4)
                                            v140 := GenS15{F0: 65535, F1: v13, F2: 65535}
                                            if len(v1) > 0 {
                                                v1[0] = 1
                                            }
                                            if true {
                                                v141 := GenS10{F0: v14, F1: (fi57 % ((int(v139)) | 1))}
                                                _ = v141
                                            } else {
                                                v13 = v13
                                                fmt.Println(v27)
                                                if v27 {
                                                    v142 := GenS11{F0: v5}
                                                    var v143 uint32 = (v13 + (v13 >> ((966) & 63)))
                                                    var v144 int16 = (v3 ^ -1)
                                                    switch 104 {
                                                    case 194:
                                                        v145 := GenS15{F0: v13, F1: (v13 % ((v143) | 1)), F2: 612}
                                                        _ = v145
                                                    case -1:
                                                        switch fi57 {
                                                        case 32767:
                                                            var v146 float64 = 1e10
                                                            if false {
                                                                var v147 uint16 = v2
                                                                var v148 uint = uint(v139)
                                                                _ = v147
                                                                _ = v148
                                                            } else {
                                                                fmt.Println(v14, v144, v144)
                                                            }
                                                            _ = v146
                                                        case 174:
                                                            v13 -= 4294967295
                                                        }
                                                    default:
                                                        v149 := GenS3{F0: v3}
                                                        _ = v149
                                                    }
                                                    if true {
                                                        _ = len(v4)
                                                    }
                                                    switch (-(fi68)) {
                                                    case -9223372036854775808:
                                                        fmt.Println(v27, v14, fi57)
                                                    }
                                                    _ = v142
                                                    _ = v143
                                                    _ = v144
                                                } else {
                                                    var v150 int8 = -31
                                                    var v151 int8 = (v150 * v150)
                                                    var v152 int16 = (v3 << ((42) & 63))
                                                    _ = v150
                                                    _ = v151
                                                    _ = v152
                                                }
                                            }
                                            _ = v139
                                            _ = v140
                                        } else {
                                            var v153 int8 = -128
                                            v153 += v153
                                            _ = v153
                                        }
                                        _ = fi68
                                        _ = v138
                                        _ = v69
                                    }
                                default:
                                    v14 -= uint8(v2)
                                }
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    for fi154 := 0; fi154 < 2; fi154++ {
                                        v2 = (v2 & v2)
                                        v13 -= (v13 << ((566) & 63))
                                        _ = fi154
                                    }
                                    fmt.Println(v2, v13)
                                }
                            }
                        }
                    }
                    if v27 {
                        _ = len(v4)
                    }
                    _ = fi57
                }
            }
            _ = v29
        }
    }
    v155 := GenS4{F0: 74, F1: v27, F2: (!(65535 <= 568))}
    var v156 string = v5
    v3 = (v3 / ((v3) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v155
    _ = v156
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
