package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 int64 }
func genFn0() uint32 {
    var v4 int64 = 9223372036854775807
    var v5 int8 = 0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 string) uint64 {
    var v6 string = p0
    var v7 float64 = 0.0
    _ = p0
    _ = v6
    _ = v7
    return 872
}

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    var v3 string = "\n\t"
    var v8 uint16 = 255
    var v9 int16 = 32767
    switch v3 {
    case "foo bar":
        v3 = "hello"
    case "":
        for fi10 := 0; fi10 < 18; fi10++ {
            switch (v3 + v3) {
            case "":
                for fi11 := 0; fi11 < 20; fi11++ {
                    var v12 int = (-(fi10))
                    v3 = v3
                    _ = fi11
                    _ = v12
                }
            case "\n\t":
                v3 = ("\n\t" + "\n\t")
            case "a":
                var v13 uint64 = genFn1("hello")
                _ = v13
            }
            if true {
                v1 = append(v1, ((v3 + (v3 + (v3 + (v3 + v3)))) + v3))
                if len(v1) > 0 {
                    v1[0] = (v3 + (v3 + ""))
                }
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 uint64 = genFn1(("hello" + ""))
                _ = fi14
                _ = v15
            }
            _ = fi10
        }
    default:
        fmt.Println(v8, v3, v9)
    }
    v8 = 289
    var v16 uint64 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        v8 -= v8
    }
    v17 := GenS1{F0: 1, F1: -1}
    for fi18 := 0; fi18 < 18; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(fi19, v9)
            if true {
                v1 = append(v1, v3)
                switch v3 {
                case "\n\t":
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                case "hello":
                    switch (("" + v3) + "a") {
                    case "\n\t":
                        if false {
                            fmt.Println(v2, fi19, v8)
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v1) > 0 {
                                v1[0] = ("" + v3)
                            }
                            switch fi18 {
                            case 2147483647:
                                if len(v1) > 0 {
                                    v1[0] = ((("foo bar" + (("foo bar" + (v3 + (("" + v3) + ((v3 + "a") + v3)))) + "")) + v3) + "a")
                                }
                            }
                        }
                    case "hello":
                        v20 := GenS2{F0: 2, F1: 0}
                        _ = v20
                    }
                }
            } else {
                v8 -= 602
                if false {
                    switch (fi19 | fi18) {
                    case -1:
                        var v21 uint = 399
                        _ = v21
                    case 32767:
                        var v22 string = v3
                        _ = v22
                    case -163:
                        fmt.Println(v9, fi19)
                    default:
                        v1 = append(v1, (v3 + ((v3 + ("a" + "hello")) + "\n\t")))
                    }
                } else {
                    if false {
                        var v23 float64 = 1e10
                        switch (fi18 >> ((18446744073709551615) & 63)) {
                        case 4294967295:
                            if (!(((v3 + v3) + v3) != (v3 + (v3 + (v3 + v3))))) {
                                v9 -= v9
                            } else {
                                v24 := GenS1{F0: v8, F1: -128}
                                _ = v24
                            }
                        case -2147483648:
                            var v25 int8 = -128
                            _ = v25
                        case 9223372036854775807:
                            v3 = "a"
                        }
                        fmt.Println(v2)
                        _ = len(v1)
                        _ = v23
                    }
                    var v26 uint8 = 0
                    v2 -= 1.5
                    _ = v26
                }
                fmt.Println(fi18, fi19, fi19)
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi19
        }
        var v27 uint32 = 72
        v27 -= 1
        _ = fi18
        _ = v27
    }
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi28 := 0; fi28 < 19; fi28++ {
            fmt.Println(v8)
            if true {
                switch v3 {
                case "a":
                    if false {
                        var v29 float64 = v2
                        var v30 bool = true
                        for fi31 := 0; fi31 < 19; fi31++ {
                            _ = len(v1)
                            fmt.Println(v29, v30, v8)
                            fmt.Println(v8, fi31)
                            var v32 uint32 = 635
                            _ = len(v1)
                            v33 := GenS2{F0: 127, F1: 0}
                            v16 += v16
                            if v30 {
                                v30 = (v30 && v30)
                                if (-1 < fi28) {
                                    switch 9223372036854775807 {
                                    case 128:
                                        v34 := GenS1{F0: 2, F1: 1}
                                        _ = v34
                                    case 32768:
                                        switch (("a" + v3) + (v3 + "foo bar")) {
                                        case "\n\t":
                                            var v35 uint16 = (v8 * (^(v8)))
                                            _ = v35
                                        case "hello":
                                            var v36 int = (fi31 >> ((18446744073709551615) & 63))
                                            _ = v36
                                        default:
                                            switch fi31 {
                                            case 32767:
                                                switch fi31 {
                                                case 128:
                                                    v9 = v9
                                                case 921:
                                                    v16 -= v16
                                                }
                                                if (127 < 100) {
                                                    fmt.Println(fi28)
                                                    var v37 int32 = -175
                                                    if ((0.1 + -2.25) >= float32(1.5)) {
                                                        var v38 uint = 2
                                                        v9 = int16(v8)
                                                        v1 = append(v1, "hello")
                                                        _ = v38
                                                    }
                                                    _ = v37
                                                }
                                                v29 += ((0.1 + v2) + float64((1e10 * 0.1)))
                                            case 990:
                                                v8 = v8
                                            default:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        switch v3 {
                                        case "\n\t":
                                            var v39 float32 = (1e10 * 1.5)
                                            _ = v39
                                        case "foo bar":
                                            var v40 string = (("foo bar" + v3) + v3)
                                            var v41 uint32 = (v32 ^ (v32 % ((v32) | 1)))
                                            _ = v40
                                            _ = v41
                                        }
                                        v9 = (v9 + v9)
                                    case 127:
                                        fmt.Println(v8, v9)
                                    default:
                                        var v42 uint16 = 65535
                                        _ = v42
                                    }
                                    v9 = 42
                                } else {
                                    v32 += (^(v32))
                                    var v43 int32 = 1
                                    _ = v43
                                }
                            } else {
                                var v44 uint32 = (v32 * v32)
                                var v45 bool = ((v30 || (-1 < 42)) && false)
                                var v46 int32 = -2147483648
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            var v47 uint32 = (v32 | v32)
                            _ = fi31
                            _ = v32
                            _ = v33
                            _ = v47
                        }
                        _ = v29
                        _ = v30
                    }
                case "hello":
                    fmt.Println(v16, v3)
                case "":
                    v9 -= (v9 % (((v9 * v9)) | 1))
                default:
                    if false {
                        var v48 string = ((((v3 + "\n\t") + v3) + "a") + "\n\t")
                        _ = v48
                    } else {
                        if ((fi28 % (((-(fi28))) | 1)) > 32767) {
                            for fi49 := 0; fi49 < 4; fi49++ {
                                if true {
                                    fmt.Println(v3)
                                    var v50 uint64 = (v16 * (^(v16)))
                                    v51 := GenS1{F0: v8, F1: 24}
                                    _ = v50
                                    _ = v51
                                } else {
                                    v9 += (v9 / ((v9) | 1))
                                    fmt.Println(fi28, fi49, v3)
                                }
                                v16 = (v16 - 2)
                                _ = fi49
                            }
                            fmt.Println(v3, fi28, v2)
                        } else {
                            var v52 uint = 4294967295
                            _ = v52
                        }
                    }
                }
                if false {
                    if (true || (882 != 42)) {
                        if false {
                            if true {
                                fmt.Println(v3, v16, v2)
                                for fi53 := 0; fi53 < 14; fi53++ {
                                    var v54 uint64 = 939
                                    var v55 int32 = -377
                                    _ = fi53
                                    _ = v54
                                    _ = v55
                                }
                                v8 -= (v8 << ((0) & 63))
                            }
                            if true {
                                for fi56 := 0; fi56 < 2; fi56++ {
                                    switch (fi28 % (((fi28 / ((-1) | 1))) | 1)) {
                                    case 127:
                                        switch (v3 + (((v3 + (v3 + ("\n\t" + (((("" + v3) + (v3 + ("foo bar" + v3))) + "a") + v3)))) + v3) + v3)) {
                                        case "foo bar":
                                            if (((v16 >> ((42) & 63)) == 1) || false) {
                                                var v57 int64 = 9223372036854775807
                                                switch fi28 {
                                                case 0:
                                                    if true {
                                                        var v58 float32 = (float32((v2 - 1.5)) + 0.0)
                                                        _ = v58
                                                    } else {
                                                        v57 += 1
                                                        v57 = 909
                                                    }
                                                case 32767:
                                                    v59 := GenS1{F0: 255, F1: 1}
                                                    _ = v59
                                                case 671:
                                                    var v60 int16 = (v9 + (v9 ^ v9))
                                                    _ = v60
                                                }
                                                v16 += 1
                                                _ = len(v1)
                                                v57 -= (v57 % (((v57 - -9223372036854775808)) | 1))
                                                _ = v57
                                            } else {
                                                var v61 int8 = -86
                                                var v62 uint = 4294967295
                                                _ = len(v1)
                                                _ = v61
                                                _ = v62
                                            }
                                        case "hello":
                                            v3 = "hello"
                                        case "":
                                            var v63 float64 = -2.25
                                            _ = v63
                                        }
                                    case 952:
                                        var v64 int8 = -128
                                        _ = v64
                                    default:
                                        fmt.Println(v9, v16, v2)
                                    }
                                    _ = fi56
                                }
                            } else {
                                v65 := GenS2{F0: 1, F1: 42}
                                if (0 != 119) {
                                    v66 := GenS3{F0: (!(!false)), F1: v9, F2: 0}
                                    v67 := GenS3{F0: (-1 >= -42), F1: -697, F2: 1}
                                    var v68 int = fi28
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                } else {
                                    switch v3 {
                                    case "hello":
                                        if (-128 < 51) {
                                            var v69 uint = 18446744073709551615
                                            v16 -= v16
                                            if (false && ((v2 - v2) == v2)) {
                                                if ((!(0.1 >= 0.0)) || false) {
                                                    if (!true) {
                                                        var v70 int32 = int32(v9)
                                                        _ = v70
                                                    } else {
                                                        for fi71 := 0; fi71 < 11; fi71++ {
                                                            v3 = "foo bar"
                                                            v9 = v9
                                                            _ = fi71
                                                        }
                                                    }
                                                    v3 = v3
                                                    fmt.Println(v9)
                                                } else {
                                                    for fi72 := 0; fi72 < 16; fi72++ {
                                                        if (uint8(v8) >= 176) {
                                                            fmt.Println(v3)
                                                            var v73 uint32 = uint32(v16)
                                                            v3 = ((v3 + "foo bar") + v3)
                                                            _ = v73
                                                        } else {
                                                            v2 += v2
                                                            var v74 string = (v3 + "\n\t")
                                                            v9 -= (v9 & -333)
                                                            _ = v74
                                                        }
                                                        v16 += (v16 << ((v69) & 63))
                                                        var v75 float32 = 0.1
                                                        _ = fi72
                                                        _ = v75
                                                    }
                                                    v1 = append(v1, v3)
                                                }
                                                if true {
                                                    switch fi28 {
                                                    case 954:
                                                        v9 = (^(v9))
                                                    case -788:
                                                        fmt.Println(v8, v69)
                                                    case -2147483648:
                                                        if false {
                                                            var v76 int32 = 489
                                                            if true {
                                                                v2 = ((((-2.25 - v2) - (v2 - 0.1)) - v2) * 0.0)
                                                            } else {
                                                                v76 += v76
                                                                for fi77 := 0; fi77 < 10; fi77++ {
                                                                    var v78 string = ((v3 + v3) + "foo bar")
                                                                    fmt.Println(fi77, v2, fi77)
                                                                    _ = fi77
                                                                    _ = v78
                                                                }
                                                            }
                                                            _ = v76
                                                        }
                                                    }
                                                    v8 -= (v8 >> ((42) & 63))
                                                }
                                            } else {
                                                switch v3 {
                                                case "\n\t":
                                                    if (v16 < (v16 | (v16 ^ (v16 * 2)))) {
                                                        for fi79 := 0; fi79 < 11; fi79++ {
                                                            if true {
                                                                if true {
                                                                    v80 := GenS2{F0: 255, F1: int16(fi28)}
                                                                    var v81 int32 = 473
                                                                    var v82 string = v3
                                                                    _ = v80
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    var v83 uint32 = 2147483648
                                                                    if false {
                                                                        switch v3 {
                                                                        case "\n\t":
                                                                            var v84 uint = v69
                                                                            _ = v84
                                                                        case "":
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v3
                                                                            }
                                                                        case "foo bar":
                                                                            v1 = append(v1, "a")
                                                                        }
                                                                        switch v3 {
                                                                        case "hello":
                                                                            v1 = append(v1, ((v3 + "") + "hello"))
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        _ = len(v1)
                                                                    } else {
                                                                        if false {
                                                                            v85 := GenS2{F0: 1, F1: 576}
                                                                            var v86 uint32 = (v83 + v83)
                                                                            switch fi28 {
                                                                            case 0:
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    switch v3 {
                                                                                    case "":
                                                                                        v9 += (v9 >> ((65535) & 63))
                                                                                    case "hello":
                                                                                        v87 := GenS1{F0: 0, F1: 127}
                                                                                        _ = v87
                                                                                    default:
                                                                                        v88 := GenS1{F0: 2, F1: 1}
                                                                                        _ = v88
                                                                                    }
                                                                                default:
                                                                                    var v89 string = v3
                                                                                    _ = v89
                                                                                }
                                                                            case -9223372036854775808:
                                                                                var v90 int = -2147483648
                                                                                _ = v90
                                                                            case 127:
                                                                                v9 = (v9 << ((v69) & 63))
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                            _ = v85
                                                                            _ = v86
                                                                        } else {
                                                                            var v91 int64 = 0
                                                                            _ = v91
                                                                        }
                                                                        if ((v83 != v83) && false) {
                                                                            v92 := GenS1{F0: (v8 % ((v8) | 1)), F1: -1}
                                                                            v93 := GenS4{F0: (^(v8))}
                                                                            _ = v92
                                                                            _ = v93
                                                                        } else {
                                                                            switch (^(fi28)) {
                                                                            case 32767:
                                                                                v2 = v2
                                                                            case 696:
                                                                                v94 := GenS4{F0: 1}
                                                                                _ = v94
                                                                            default:
                                                                                fmt.Println(fi79, v3, v83)
                                                                            }
                                                                            v95 := GenS1{F0: v8, F1: 70}
                                                                            _ = len(v1)
                                                                            v96 := GenS5{F0: 2147483647}
                                                                            fmt.Println(v8)
                                                                            _ = v95
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    var v97 int64 = 780
                                                                    _ = v83
                                                                    _ = v97
                                                                }
                                                                v8 += v8
                                                            }
                                                            if (1 >= (^(v8))) {
                                                                v69 -= v69
                                                            }
                                                            if (true && true) {
                                                                _ = len(v1)
                                                                switch ((v3 + "\n\t") + v3) {
                                                                case "hello":
                                                                    var v98 int8 = 1
                                                                    _ = v98
                                                                case "a":
                                                                    if true {
                                                                        v9 -= v9
                                                                        if (!((v2 / (((1.5 / 1.0) - v2) + 3.14159)) >= v2)) {
                                                                            fmt.Println(v2, fi79, v69)
                                                                            if true {
                                                                                if true {
                                                                                    if true {
                                                                                        switch (fi79 << (((v69 >> ((v69) & 63))) & 63)) {
                                                                                        case 0:
                                                                                            var v99 string = ""
                                                                                            _ = v99
                                                                                        default:
                                                                                            v16 += (v16 / (((v16 + (^(v16)))) | 1))
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    switch (("\n\t" + v3) + (v3 + "")) {
                                                                                    case "hello":
                                                                                        v2 -= v2
                                                                                    case "foo bar":
                                                                                        var v100 int32 = 1
                                                                                        _ = v100
                                                                                    case "":
                                                                                        v101 := GenS6{F0: v16, F1: v3}
                                                                                        _ = v101
                                                                                    default:
                                                                                        _ = len(v1)
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v9, v3, fi79)
                                                                                    v1 = append(v1, "\n\t")
                                                                                }
                                                                                var v102 uint64 = (v16 & v16)
                                                                                _ = v102
                                                                            }
                                                                            v2 = v2
                                                                        }
                                                                    }
                                                                    v103 := GenS1{F0: v8, F1: int8(v8)}
                                                                    v8 = v8
                                                                    if ((float64(float32((v2 * 3.14159))) * v2) > (v2 * v2)) {
                                                                        var v104 uint = 990
                                                                        _ = v104
                                                                    }
                                                                    _ = v103
                                                                default:
                                                                    if (!(128 < 2)) {
                                                                        if false {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v105 uint8 = 0
                                                                            v106 := GenS7{F0: 489}
                                                                            _ = v105
                                                                            _ = v106
                                                                        } else {
                                                                            v69 += uint(v16)
                                                                            v107 := GenS7{F0: v9}
                                                                            _ = v107
                                                                        }
                                                                    }
                                                                    var v108 int8 = 1
                                                                    fmt.Println(v16)
                                                                    v109 := GenS7{F0: (-(v9))}
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                                _ = len(v1)
                                                            }
                                                            _ = fi79
                                                        }
                                                        if true {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v110 string = v3
                                                            var v111 int64 = -9223372036854775808
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            var v112 uint32 = 2
                                                            v2 = (v2 * v2)
                                                            _ = v112
                                                        }
                                                        var v113 int8 = 127
                                                        _ = v113
                                                    }
                                                }
                                                var v114 int8 = 94
                                                _ = v114
                                            }
                                            _ = v69
                                        }
                                    default:
                                        var v115 float64 = ((v2 / float64(1e10)) + (3.14159 / v2))
                                        _ = v115
                                    }
                                    v116 := GenS5{F0: genFn0()}
                                    _ = v116
                                }
                                _ = v65
                            }
                            if (1 > -355) {
                                v16 = genFn1(v3)
                                v3 = "\n\t"
                                for fi117 := 0; fi117 < 6; fi117++ {
                                    v118 := GenS3{F0: (!false), F1: -46, F2: 11}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi117
                                    _ = v118
                                }
                            }
                        }
                    } else {
                        for fi119 := 0; fi119 < 18; fi119++ {
                            if false {
                                v120 := GenS8{F0: (v8 % ((v8) | 1)), F1: 1}
                                var v121 uint32 = 42
                                _ = v120
                                _ = v121
                            }
                            fmt.Println(v9)
                            _ = fi119
                        }
                    }
                } else {
                    v1 = append(v1, ((v3 + "a") + (v3 + v3)))
                    switch (^(fi28)) {
                    case -2147483648:
                        if true {
                            switch fi28 {
                            case -990:
                                _ = len(v1)
                            default:
                                v2 += v2
                            }
                            for fi122 := 0; fi122 < 3; fi122++ {
                                var v123 int16 = v9
                                fmt.Println(v3)
                                v1 = append(v1, "")
                                fmt.Println(v2)
                                _ = fi122
                                _ = v123
                            }
                        }
                    default:
                        v1 = append(v1, v3)
                    }
                    v124 := GenS6{F0: 248, F1: (v3 + v3)}
                    _ = v124
                }
            } else {
                v1 = append(v1, v3)
                v2 = float64((0.1 / 3.14159))
                var v125 float32 = 1e10
                _ = v125
            }
            _ = len(v1)
            _ = fi28
        }
        for fi126 := 0; fi126 < 17; fi126++ {
            v16 = genFn1((v3 + ""))
            for fi127 := 0; fi127 < 2; fi127++ {
                var v128 uint64 = (v16 >> ((1) & 63))
                fmt.Println(v9)
                _ = fi127
                _ = v128
            }
            _ = fi126
        }
    }
    var v129 uint = 255
    var v130 uint16 = v8
    var v131 uint32 = 940
    v132 := GenS6{F0: (v16 - (^(v16))), F1: v3}
    if false {
        v3 = (("hello" + v3) + "hello")
    } else {
        var v133 int64 = 0
        for fi134 := 0; fi134 < 12; fi134++ {
            v131 -= v131
            _ = fi134
        }
        v135 := GenS2{F0: 63, F1: int16(v129)}
        _ = v133
        _ = v135
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v9, v131, v130)
    var v136 int32 = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v136:", v136)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v136
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
