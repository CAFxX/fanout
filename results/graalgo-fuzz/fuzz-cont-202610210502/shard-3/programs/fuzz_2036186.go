package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 int8 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8; F1 int }
type GenS13 struct { F0 int8 }
func genFn0(p0 bool, p1 int16, p2 string) uint {
    p2 = "a"
    v7 := []int{-9223372036854775808, -761, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{94, 2147483647, 65535}
    v2 := []uint32{}
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 980}
        for fi10 := 0; fi10 < 20; fi10++ {
            v2 = append(v2, 42)
            _ = fi10
        }
        v6 += (v6 % (((v6 + (v6 / (((^(v6))) | 1)))) | 1))
        _ = fi8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        v6 -= v6
    case "":
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int8 = 1
            var v13 string = "foo bar"
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 1}
    fmt.Println(v6)
    switch "" {
    case "hello":
        var v15 uint64 = (^(v6))
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS1{F0: 2}
            if false {
                if false {
                    for fi18 := 0; fi18 < 4; fi18++ {
                        v19 := GenS2{F0: 888}
                        _ = fi18
                        _ = v19
                    }
                    fmt.Println(fi16)
                    v6 = 1
                }
                if (9223372036854775807 < 787) {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v6 += 0
                        var v21 int64 = -9223372036854775808
                        _ = fi20
                        _ = v21
                    }
                    if false {
                        v22 := GenS1{F0: 149}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v6 = v6
                            _ = fi23
                        }
                        v6 = v6
                    }
                    for fi24 := 0; fi24 < 20; fi24++ {
                        var v25 uint = genFn0(true, 332, ("" + ""))
                        _ = fi24
                        _ = v25
                    }
                } else {
                    switch (fi16 * (fi16 % ((-728) | 1))) {
                    case 42:
                        v26 := GenS1{F0: genFn0((false || false), -32768, ("a" + ""))}
                        _ = v26
                    case 989:
                        for fi27 := 0; fi27 < 7; fi27++ {
                            switch (fi27 % ((fi27) | 1)) {
                            case 2147483647:
                                var v28 uint64 = 2
                                _ = v28
                            case 618:
                                fmt.Println(fi16, fi16, fi16)
                            default:
                                var v29 bool = (-1 == -1)
                                fmt.Println(fi27)
                                var v30 uint64 = v6
                                if v29 {
                                    if v29 {
                                        v29 = (42 != 32767)
                                        var v31 int32 = 1
                                        v6 -= (v6 | v30)
                                        v32 := GenS1{F0: 0}
                                        switch (fi27 + fi27) {
                                        case -855:
                                            switch "hello" {
                                            case "foo bar":
                                                switch (fi27 % ((fi16) | 1)) {
                                                case 128:
                                                    if false {
                                                        if (!true) {
                                                            var v33 string = ("\n\t" + "hello")
                                                            v34 := GenS2{F0: 1}
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        v29 = v29
                                                        var v35 uint16 = 42
                                                        _ = v35
                                                    } else {
                                                        v31 -= v31
                                                        var v36 float32 = 0.1
                                                        _ = v36
                                                    }
                                                    v31 -= (v31 >> ((100) & 63))
                                                default:
                                                    v37 := GenS3{F0: (("hello" + "") + "hello"), F1: 0}
                                                    _ = v37
                                                }
                                            case "hello":
                                                v38 := GenS3{F0: (("" + "foo bar") + ""), F1: -21}
                                                _ = v38
                                            case "a":
                                                v31 -= (v31 % (((v31 * v31)) | 1))
                                            }
                                        case 128:
                                            v39 := GenS4{F0: v31, F1: 42}
                                            _ = v39
                                        case 2147483647:
                                            var v40 int64 = -9223372036854775808
                                            _ = v40
                                        }
                                        _ = v31
                                        _ = v32
                                    }
                                    fmt.Println(v6)
                                    if (127 < 0) {
                                        fmt.Println(fi16, fi16)
                                        v4[4294967295] = (fi27 / (((^(fi16))) | 1))
                                    } else {
                                        v30 += v6
                                        v41 := v4[0]
                                        _ = v41
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    if v29 {
                                        var v42 int8 = -1
                                        v4[4294967295] = fi16
                                        _ = v42
                                    }
                                    v29 = (true || v29)
                                    if (!(!(673 == 2147483647))) {
                                        if true {
                                            var v43 uint32 = 2147483647
                                            v29 = ((-32768 >= -84) && false)
                                            v43 -= (v43 / ((1) | 1))
                                            _ = v43
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 269
                                        }
                                    } else {
                                        switch fi27 {
                                        case 2147483647:
                                            fmt.Println(v30, fi16, v6)
                                        case 1:
                                            v44 := GenS4{F0: 42, F1: 42}
                                            _ = v44
                                        default:
                                            if false {
                                                if v29 {
                                                    v30 = (v30 + (v6 * v30))
                                                    var v45 string = ("hello" + "")
                                                    v29 = v29
                                                    _ = v45
                                                }
                                            }
                                        }
                                    }
                                    v29 = true
                                }
                                v46 := v3[(("" + "a") + "foo bar")]
                                v46 = 42
                                if true {
                                    if (((44 != 127) && true) && (182 < 0)) {
                                        v46 = (v46 >> ((65535) & 63))
                                        if v29 {
                                            v2 = append(v2, v46)
                                            v30 = (v6 - v6)
                                            var v47 float64 = float64(float32((-2.25 - ((((0.0 - (3.14159 / (0.1 / -2.25))) * -2.25) - (0.1 * (0.1 - float64(-2.25)))) / 1.5))))
                                            _ = v47
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                v48 := GenS5{F0: 65535, F1: (v46 | (v46 >> ((uint(v6)) & 63))), F2: 0}
                                                _ = v48
                                            case "hello":
                                                v49 := GenS6{F0: 0}
                                                _ = v49
                                            default:
                                                fmt.Println(v29)
                                            }
                                            if v29 {
                                                if v29 {
                                                    var v50 int8 = -41
                                                    switch ("foo bar" + "\n\t") {
                                                    case "hello":
                                                        var v51 bool = (0 <= 255)
                                                        _ = v51
                                                    case "a":
                                                        var v52 float32 = (1e10 / ((3.14159 * 1.5) / (-2.25 + 1e10)))
                                                        _ = v52
                                                    default:
                                                        switch fi16 {
                                                        case -9223372036854775808:
                                                            switch (fi27 & fi16) {
                                                            case 128:
                                                                if false {
                                                                    v29 = (-1 != -1)
                                                                    v50 -= (v50 / (((v50 & (v50 + v50))) | 1))
                                                                }
                                                            default:
                                                                switch (fi27 | (-(fi27))) {
                                                                case -683:
                                                                    v6 = (v6 / ((v6) | 1))
                                                                case -1:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                case 0:
                                                                    v53 := GenS4{F0: 2147483647, F1: v50}
                                                                    _ = v53
                                                                default:
                                                                    switch fi16 {
                                                                    case 2147483647:
                                                                        v54 := GenS6{F0: 259}
                                                                        _ = v54
                                                                    default:
                                                                        v6 += 42
                                                                    }
                                                                    switch ("hello" + "hello") {
                                                                    case "":
                                                                        var v55 uint64 = v30
                                                                        _ = v55
                                                                    case "hello":
                                                                        switch (fi16 << ((670) & 63)) {
                                                                        case 944:
                                                                            v50 += v50
                                                                        case -1:
                                                                            if v29 {
                                                                                v6 = (^(v6))
                                                                                delete(v3, "hello")
                                                                                var v56 bool = v29
                                                                                _ = v56
                                                                            } else {
                                                                                var v57 int8 = 110
                                                                                var v58 uint32 = (v46 << ((42) & 63))
                                                                                _ = v57
                                                                                _ = v58
                                                                            }
                                                                        case -505:
                                                                            v29 = v29
                                                                        default:
                                                                            v6 += 18446744073709551615
                                                                        }
                                                                    case "\n\t":
                                                                        v46 = v46
                                                                    default:
                                                                        fmt.Println(v46)
                                                                    }
                                                                    fmt.Println(fi16)
                                                                }
                                                            }
                                                        case -962:
                                                            v59 := GenS5{F0: 1, F1: (^(v46)), F2: 42}
                                                            _ = v59
                                                        default:
                                                            if (!(genFn0(v29, 32767, "") <= 1)) {
                                                                if false {
                                                                    if false {
                                                                        delete(v3, "hello")
                                                                        v29 = ((65535 != 42) || v29)
                                                                    }
                                                                    v50 = (v50 % (((^(v50))) | 1))
                                                                    v29 = (255 == 0)
                                                                }
                                                                if (!(-909 != 748)) {
                                                                    v5[v46] = v46
                                                                }
                                                                var v60 float64 = -2.25
                                                                _ = v60
                                                            }
                                                        }
                                                    }
                                                    _ = len(v4)
                                                    _ = v50
                                                } else {
                                                    if v29 {
                                                        var v61 uint = 520
                                                        v46 = v46
                                                        _ = v61
                                                    } else {
                                                        var v62 int32 = -2147483648
                                                        fmt.Println(v46, v29, v6)
                                                        _ = v62
                                                    }
                                                }
                                                v63 := v5[v46]
                                                var v64 uint32 = v63
                                                _ = v63
                                                _ = v64
                                            } else {
                                                fmt.Println(v30, v46)
                                                if (((0 > -2147483648) || v29) && (1.5 == 0.0)) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v5)
                                                }
                                            }
                                        }
                                        var v65 float32 = 1.5
                                        _ = v65
                                    }
                                } else {
                                    v30 += (v30 << ((65535) & 63))
                                    v6 = v30
                                    if len(v2) > 0 {
                                        v2[0] = (v46 >> ((470) & 63))
                                    }
                                }
                                _ = v29
                                _ = v30
                                _ = v46
                            }
                            fmt.Println(fi27)
                            _ = fi27
                        }
                    case 2147483647:
                        fmt.Println(v6, v6)
                    }
                    v6 -= v6
                    var v66 int = 947
                    _ = v66
                }
                var v67 string = "foo bar"
                _ = v67
            } else {
                v6 -= (v6 & v6)
                switch ("" + "a") {
                case "a":
                    for fi68 := 0; fi68 < 7; fi68++ {
                        v1 = append(v1, 978)
                        _ = fi68
                    }
                case "hello":
                    v6 = (v6 % (((v6 - 0)) | 1))
                case "":
                    var v69 int32 = 0
                    _ = v69
                default:
                    if true {
                        v6 += (v6 ^ (^(v6)))
                        v6 -= 42
                        v1 = append(v1, 1)
                    } else {
                        for fi70 := 0; fi70 < 14; fi70++ {
                            fmt.Println(fi70, fi16, fi70)
                            _ = fi70
                        }
                    }
                }
                v71 := GenS7{F0: false}
                _ = v71
            }
            var v72 int8 = 127
            _ = fi16
            _ = v17
            _ = v72
        }
    default:
        switch 9223372036854775807 {
        case 128:
            v6 = 0
        case 127:
            var v73 bool = false
            _ = v73
        }
    }
    fmt.Println(v6)
    if false {
        if true {
            var v74 uint16 = 1
            _ = v74
        } else {
            var v75 int = 32767
            _ = v75
        }
        var v76 int = 4294967295
        v77 := GenS4{F0: -2147483648, F1: -128}
        _ = v76
        _ = v77
    } else {
        v6 = (v6 - (v6 / ((v6) | 1)))
    }
    var v78 bool = (false || (1 == 1))
    switch 127 {
    case -973:
        fmt.Println(v78, v6)
    case 9223372036854775807:
        for fi79 := 0; fi79 < 14; fi79++ {
            var v80 int16 = -358
            if false {
                var v81 bool = (v6 > (v6 / ((4294967295) | 1)))
                v82 := GenS8{F0: 692, F1: -14}
                _ = v81
                _ = v82
            } else {
                v5[1] = 2
                v78 = false
                var v83 float64 = (((1e10 / (1.5 + (0.0 / 1.0))) * 0.1) + ((1.5 / (0.1 * (-2.25 - (float64(0.1) - 1e10)))) / (1e10 * (((0.1 - ((float64(((1e10 + ((1.5 - (float32(3.14159) - ((0.0 * 1e10) * float32(1.5)))) / float32((0.1 - (1e10 / (1e10 / 3.14159)))))) + 0.0)) / -2.25) + (0.1 / 1.5))) - 1.5) + (-2.25 * (((1e10 / 1.5) + 1e10) - 1.5))))))
                _ = v83
            }
            _ = fi79
            _ = v80
        }
    case -84:
        for fi84 := 0; fi84 < 4; fi84++ {
            v85 := GenS9{F0: ("a" + "foo bar")}
            var v86 int8 = 42
            _ = fi84
            _ = v85
            _ = v86
        }
    }
    if (v78 && v78) {
        switch 183 {
        case -379:
            v78 = v78
        case 127:
            if (167 < -399) {
                v6 = (v6 - 0)
            } else {
                fmt.Println(v6, v6)
                for fi87 := 0; fi87 < 18; fi87++ {
                    v6 += 2
                    switch (("" + "hello") + "foo bar") {
                    case "\n\t":
                        var v88 float64 = -2.25
                        _ = v88
                    default:
                        var v89 bool = (235 == 2)
                        _ = v89
                    }
                    if (206 != 2) {
                        switch "\n\t" {
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        case "hello":
                            v3[("hello" + ("\n\t" + ("foo bar" + "")))] = 2147483648
                        case "":
                            var v90 int = fi87
                            _ = v90
                        }
                        v91 := GenS5{F0: 65535, F1: 2, F2: 676}
                        var v92 int32 = 42
                        _ = v91
                        _ = v92
                    }
                    _ = fi87
                }
            }
        default:
            switch ("\n\t" + "") {
            case "\n\t":
                if v78 {
                    for fi93 := 0; fi93 < 8; fi93++ {
                        for fi94 := 0; fi94 < 12; fi94++ {
                            if (!(1e10 <= 1.5)) {
                                _ = len(v4)
                                fmt.Println(v78, v6, fi94)
                                if false {
                                    v95 := GenS10{F0: (v6 << ((4294967295) & 63)), F1: v78, F2: 31}
                                    var v96 float32 = 1e10
                                    fmt.Println(v6, v6, v78)
                                    var v97 int16 = 1
                                    v98 := v5[2]
                                    fmt.Println(fi93, fi94)
                                    fmt.Println(fi93)
                                    _ = v95
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                }
                            }
                            _ = fi94
                        }
                        var v99 int16 = -1
                        v6 = (v6 >> ((42) & 63))
                        _ = fi93
                        _ = v99
                    }
                    if (true || (!v78)) {
                        v78 = (1.5 <= 1.5)
                        v6 = 261
                        if v78 {
                            v100 := v5[2]
                            var v101 string = "a"
                            fmt.Println(v101, v100)
                            _ = v100
                            _ = v101
                        } else {
                            v78 = v78
                        }
                    } else {
                        for fi102 := 0; fi102 < 10; fi102++ {
                            fmt.Println(fi102, fi102)
                            if v78 {
                                _ = len(v4)
                                var v103 int16 = 42
                                var v104 int = 4294967295
                                _ = v103
                                _ = v104
                            }
                            _ = fi102
                        }
                        v105 := GenS11{F0: 65535}
                        _ = v105
                    }
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        v106 := GenS6{F0: 9223372036854775807}
        if len(v2) > 0 {
            v2[0] = 65535
        }
        v107 := GenS12{F0: 127, F1: -774}
        _ = v106
        _ = v107
    }
    if (32767 >= 169) {
        if v78 {
            for fi108 := 0; fi108 < 12; fi108++ {
                fmt.Println(v78)
                _ = fi108
            }
            for fi109 := 0; fi109 < 16; fi109++ {
                for fi110 := 0; fi110 < 7; fi110++ {
                    fmt.Println(fi109, fi110)
                    switch (fi110 ^ 127) {
                    case 32768:
                        switch "" {
                        case "a":
                            var v111 string = (("a" + "\n\t") + "a")
                            _ = v111
                        case "hello":
                            v78 = true
                        default:
                            switch (fi110 & fi109) {
                            case 2147483647:
                                v112 := GenS12{F0: -128, F1: int(v6)}
                                _ = v112
                            default:
                                v78 = v78
                            }
                        }
                    case 9223372036854775807:
                        fmt.Println(v6)
                    }
                    if v78 {
                        fmt.Println(v78)
                        var v113 uint16 = 65535
                        _ = v113
                    } else {
                        var v114 int64 = 9223372036854775807
                        v114 += v114
                        v114 += (v114 - 1)
                        _ = v114
                    }
                    _ = fi110
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v115 int64 = -9223372036854775808
                _ = fi109
                _ = v115
            }
        } else {
            if ((18446744073709551615 != (v6 * (^(v6)))) && (false || v78)) {
                var v116 int16 = -800
                if (v116 == 32767) {
                    for fi117 := 0; fi117 < 16; fi117++ {
                        for fi118 := 0; fi118 < 14; fi118++ {
                            fmt.Println(v6)
                            fmt.Println(fi118, v116, v78)
                            _ = fi118
                        }
                        _ = fi117
                    }
                } else {
                    var v119 uint32 = 4294967295
                    _ = v119
                }
                v120 := GenS13{F0: -1}
                _ = v116
                _ = v120
            } else {
                var v121 int = 1
                _ = v121
            }
            v78 = v78
        }
        v122 := GenS6{F0: -762}
        _ = v122
    }
    var v123 int = 4294967295
    fmt.Println(v6, v123)
    var v124 uint = 4294967295
    var v125 uint64 = (v6 + v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v123
    _ = v124
    _ = v125
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
}
