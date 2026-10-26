package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int8; F1 int16 }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 uint64; F2 string }
type GenS14 struct { F0 int; F1 int32 }
type GenS15 struct { F0 uint16; F1 int32; F2 uint }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint16; F1 bool; F2 uint }
type GenS19 struct { F0 uint64 }
func genFn0(p0 uint) int16 {
    v6 := []int{-130, 127, 652}
    if false {
        p0 -= (^(p0))
        var v7 int8 = 0
        _ = v7
    } else {
        switch -1 {
        case 2147483647:
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: -32768}
                _ = len(v6)
                p0 = p0
                _ = fi8
                _ = v9
            }
        case 734:
            for fi10 := 0; fi10 < 4; fi10++ {
                switch "foo bar" {
                case "a":
                    v6 = append(v6, 32767)
                }
                var v11 uint = 385
                var v12 uint16 = 65535
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            p0 = (p0 >> ((p0) & 63))
        }
        p0 = 18446744073709551615
        if false {
            if (!(-968 != 4294967295)) {
                p0 = p0
            } else {
                p0 += p0
            }
            v13 := GenS2{F0: 218, F1: true, F2: 607}
            _ = len(v6)
            _ = v13
        } else {
            p0 -= (p0 + (p0 % ((p0) | 1)))
            p0 = 4294967295
            if (!false) {
                var v14 uint64 = 1
                var v15 uint = p0
                _ = v14
                _ = v15
            }
        }
    }
    _ = p0
    _ = v6
    return -32768
}
func genFn1() int64 {
    var v16 int = -214
    v17 := GenS2{F0: 127, F1: true, F2: 42}
    v18 := []int{(v16 << ((0) & 63))}
    if true {
        if false {
            var v19 uint32 = 4294967295
            var v20 uint = 1
            switch ("a" + "a") {
            case "":
                v21 := GenS3{F0: (false && false), F1: 255, F2: 2}
                _ = v21
            case "hello":
                v19 = (v19 * uint32(v20))
            default:
                fmt.Println(v19, v20, v20)
            }
            _ = v19
            _ = v20
        } else {
            v16 -= v16
            var v22 int8 = 100
            switch v16 {
            case 127:
                fmt.Println(v16, v22)
            default:
                for fi23 := 0; fi23 < 13; fi23++ {
                    v18 = append(v18, (v16 % ((1) | 1)))
                    if len(v18) >= 2 {
                        v18 = v18[1:len(v18)]
                    }
                    _ = fi23
                }
            }
            _ = v22
        }
    }
    var v24 float32 = 0.1
    _ = v16
    _ = v17
    _ = v18
    _ = v24
    return -305
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 bool = (false || false)
    var v5 bool = v4
    if v4 {
        if v4 {
            v25 := []int{32767}
            _ = v25
        } else {
            v26 := GenS4{F0: -1, F1: 1, F2: ("hello" + "\n\t")}
            _ = v26
        }
        v27 := GenS3{F0: (false && v3), F1: 140, F2: 1}
        var v28 int64 = -119
        _ = v27
        _ = v28
    } else {
        v29 := GenS5{F0: 928}
        for fi30 := 0; fi30 < 17; fi30++ {
            for fi31 := 0; fi31 < 13; fi31++ {
                v1 += (^(v1))
                if v3 {
                    fmt.Println(fi30)
                    var v32 int32 = -2147483648
                    _ = v32
                }
                delete(v2, 1)
                _ = fi31
            }
            _ = fi30
        }
        if true {
            var v33 string = "a"
            v33 = v33
            _ = v33
        }
        _ = v29
    }
    v34 := GenS2{F0: 178, F1: (!v4), F2: -144}
    v35 := []int{}
    if (!(((v5 || v3) && (v3 || (true && (1e10 < 0.1)))) && (4294967295 < 32767))) {
        v5 = (510 == 18446744073709551615)
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 int32 = -434
            fmt.Println(v37, v4)
            fmt.Println(v3, v5, v3)
            _ = fi36
            _ = v37
        }
    }
    v38 := GenS6{F0: v5, F1: 35, F2: 2147483647}
    v39 := GenS7{F0: -128, F1: genFn0(4294967295)}
    v4 = (v3 && v5)
    for fi40 := 0; fi40 < 9; fi40++ {
        v4 = v4
        _ = fi40
    }
    var v41 int32 = 51
    switch -2147483648 {
    case 2147483647:
        var v42 float32 = ((0.0 + 0.1) + 0.0)
        _ = v42
    }
    v43 := GenS8{F0: 9223372036854775807, F1: "hello"}
    _ = len(v35)
    for fi44 := 0; fi44 < 13; fi44++ {
        fmt.Println(v3)
        if v3 {
            var v45 bool = v3
            if (42 >= 18446744073709551615) {
                v4 = (!false)
                v41 += -1
            } else {
                switch (^(fi44)) {
                case -1:
                    var v46 int64 = -9223372036854775808
                    _ = v46
                default:
                    var v47 int64 = -1
                    _ = v47
                }
                var v48 uint32 = 42
                switch (fi44 % (((fi44 >> ((65535) & 63))) | 1)) {
                case 127:
                    v48 -= 2147483648
                case -526:
                    switch "a" {
                    case "foo bar":
                        v48 = (^(v48))
                    }
                case -9223372036854775808:
                    if ((!(v4 && (true || ((53 > 42) && (!v5))))) && (uint8(v48) >= 255)) {
                        for fi49 := 0; fi49 < 3; fi49++ {
                            fmt.Println(v4, v1)
                            v50 := GenS9{F0: 1}
                            _ = fi49
                            _ = v50
                        }
                    } else {
                        v51 := GenS10{F0: v48}
                        var v52 uint16 = 255
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                        _ = v51
                        _ = v52
                    }
                default:
                    if v45 {
                        _ = len(v2)
                        switch ("\n\t" + "hello") {
                        case "":
                            if false {
                                var v53 uint8 = 42
                                if v3 {
                                    if ((-840 == (v41 / ((int32(v48)) | 1))) && (v3 || (42 != 9223372036854775807))) {
                                        if (!(-416 > int64(v41))) {
                                            v41 += (v41 << ((4294967295) & 63))
                                            v35 = append(v35, 127)
                                        }
                                    }
                                    var v54 int32 = 1
                                    _ = v54
                                } else {
                                    if v45 {
                                        var v55 uint16 = 32767
                                        v56 := v2[42]
                                        _ = v55
                                        _ = v56
                                    }
                                    v48 -= (^(v48))
                                    switch (fi44 << ((977) & 63)) {
                                    case 294:
                                        fmt.Println(v4)
                                    case 32767:
                                        if len(v35) > 0 {
                                            v35[0] = int(v53)
                                        }
                                    }
                                }
                                for fi57 := 0; fi57 < 5; fi57++ {
                                    v45 = v3
                                    v58 := v2[v48]
                                    _ = fi57
                                    _ = v58
                                }
                                _ = v53
                            } else {
                                _ = len(v2)
                                var v59 int = (fi44 * fi44)
                                if (true && ((false || (1 > 0)) && v45)) {
                                    v45 = (!(4294967295 >= 42))
                                    v48 = v48
                                    for fi60 := 0; fi60 < 3; fi60++ {
                                        var v61 uint8 = 59
                                        var v62 uint = 65535
                                        fmt.Println(v62, v62)
                                        _ = fi60
                                        _ = v61
                                        _ = v62
                                    }
                                } else {
                                    switch v59 {
                                    case 127:
                                        v63 := GenS7{F0: 127, F1: v1}
                                        _ = v63
                                    case 2147483647:
                                        var v64 string = "\n\t"
                                        _ = v64
                                    }
                                }
                                _ = v59
                            }
                        case "foo bar":
                            fmt.Println(v41, v41)
                        }
                    }
                }
                _ = v48
            }
            for fi65 := 0; fi65 < 15; fi65++ {
                if (((1.5 - 1e10) / 1.0) >= (1e10 * 0.1)) {
                    delete(v2, 860)
                } else {
                    if (42 == 5) {
                        if (v45 && (42 != 466)) {
                            var v66 int8 = 100
                            if (("hello" > "foo bar") || v4) {
                                if (v4 || (!(!v45))) {
                                    var v67 int8 = -128
                                    switch "foo bar" {
                                    case "\n\t":
                                        v1 += (v1 / ((genFn0(0)) | 1))
                                    case "a":
                                        var v68 bool = v45
                                        _ = v68
                                    default:
                                        var v69 int8 = int8(fi44)
                                        _ = v69
                                    }
                                    if true {
                                        var v70 int64 = -9223372036854775808
                                        if (0.0 >= ((0.1 + float64(((1e10 / 1e10) + 3.14159))) * -2.25)) {
                                            v35 = append(v35, fi44)
                                            delete(v2, 42)
                                            v71 := GenS9{F0: 4294967295}
                                            _ = v71
                                        }
                                        if v4 {
                                            fmt.Println(v1)
                                        }
                                        _ = v70
                                    }
                                    if true {
                                        var v72 int16 = v1
                                        _ = len(v2)
                                        var v73 float32 = (3.14159 / 1.0)
                                        _ = v72
                                        _ = v73
                                    } else {
                                        v74 := GenS9{F0: 0}
                                        if v45 {
                                            var v75 bool = v4
                                            _ = v75
                                        } else {
                                            if (v4 && true) {
                                                v41 -= (v41 & v41)
                                                if (597 >= 42) {
                                                    var v76 int32 = v41
                                                    fmt.Println(v67, v76)
                                                    if v3 {
                                                        var v77 int8 = v66
                                                        _ = len(v2)
                                                        var v78 int64 = -311
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                    _ = v76
                                                }
                                                v41 = (v41 % ((v41) | 1))
                                                var v79 uint8 = 0
                                                _ = v79
                                            }
                                            var v80 float32 = 0.1
                                            var v81 int = fi44
                                            _ = v80
                                            _ = v81
                                        }
                                        _ = v74
                                    }
                                    _ = v67
                                } else {
                                    v4 = v3
                                    var v82 int16 = v1
                                    _ = v82
                                }
                                v1 = (-(v1))
                            } else {
                                if (v41 <= (v41 ^ (v41 & (v41 * v41)))) {
                                    v83 := v2[65535]
                                    v3 = v45
                                    v84 := GenS6{F0: v5, F1: 690, F2: 942}
                                    _ = v83
                                    _ = v84
                                } else {
                                    switch 32768 {
                                    case 127:
                                        if (255 == 255) {
                                            v1 -= (-(v1))
                                            fmt.Println(v5)
                                            v66 -= 0
                                            if (!v4) {
                                                v4 = (false || false)
                                            }
                                        } else {
                                            v35 = append(v35, fi65)
                                            v45 = v45
                                            if false {
                                                v85 := GenS6{F0: ((((v1 & (v1 & v1)) > -642) && (1 < 1)) && true), F1: 4294967295, F2: 4294967295}
                                                if v3 {
                                                    v86 := GenS1{F0: v1}
                                                    fmt.Println(v41)
                                                    if v5 {
                                                        var v87 int32 = v41
                                                        var v88 int32 = v41
                                                        fmt.Println(v3, v45)
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        var v89 int64 = 429
                                                        _ = v89
                                                    }
                                                    _ = v86
                                                } else {
                                                    v3 = ((2147483648 >= 356) || (!(fi44 == 1)))
                                                    if v45 {
                                                        v90 := GenS7{F0: int8(fi44), F1: v1}
                                                        _ = v90
                                                    }
                                                    var v91 bool = v45
                                                    switch fi65 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v5, v5, fi44)
                                                        var v92 uint32 = 2
                                                        _ = len(v35)
                                                        delete(v2, v92)
                                                        switch "hello" {
                                                        case "\n\t":
                                                            fmt.Println(v3)
                                                            v93 := GenS6{F0: (("a" + "") != "\n\t"), F1: 18446744073709551615, F2: 2}
                                                            v1 += (v1 | (v1 + genFn0(65535)))
                                                            _ = v93
                                                        case "foo bar":
                                                            v94 := GenS11{F0: v66}
                                                            if v4 {
                                                                v95 := GenS4{F0: -1, F1: 32767, F2: (((((("a" + ("a" + "hello")) + "") + ("foo bar" + ("foo bar" + "foo bar"))) + "hello") + ("" + "\n\t")) + "foo bar")}
                                                                if (254 == 42) {
                                                                    if ("" == "foo bar") {
                                                                        fmt.Println(v41, v41, v91)
                                                                    }
                                                                }
                                                                _ = v95
                                                            } else {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    switch (fi44 | fi44) {
                                                                    case 4294967295:
                                                                        v96 := GenS10{F0: 4294967295}
                                                                        _ = v96
                                                                    }
                                                                case "foo bar":
                                                                    var v97 float32 = ((3.14159 - 0.1) / 1.5)
                                                                    _ = v97
                                                                case "hello":
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        v98 := GenS6{F0: true, F1: 0, F2: (v92 & (v92 - v92))}
                                                                        _ = v98
                                                                    case "\n\t":
                                                                        switch fi44 {
                                                                        case -9223372036854775808:
                                                                            v5 = (18446744073709551615 == 419)
                                                                            v5 = false
                                                                            _ = len(v35)
                                                                        case 42:
                                                                            switch "a" {
                                                                            case "":
                                                                                var v99 int32 = 42
                                                                                _ = v99
                                                                            }
                                                                        case -744:
                                                                            var v100 uint8 = 128
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                default:
                                                                    fmt.Println(v1)
                                                                    _ = len(v35)
                                                                    var v101 uint8 = 234
                                                                    _ = v101
                                                                }
                                                                v2[1] = (fi44 << ((2) & 63))
                                                                var v102 int32 = (-(v41))
                                                                _ = v102
                                                            }
                                                            _ = v94
                                                        default:
                                                            fmt.Println(fi44)
                                                        }
                                                        _ = v92
                                                    case -682:
                                                        if len(v35) > 0 {
                                                            v35[0] = 127
                                                        }
                                                        v103 := GenS9{F0: 1}
                                                        _ = v103
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        var v104 int32 = v41
                                                        _ = v104
                                                    case "":
                                                        v45 = v45
                                                    case "a":
                                                        v4 = v45
                                                    }
                                                    _ = v91
                                                }
                                                _ = v85
                                            }
                                        }
                                    case 42:
                                        var v105 uint32 = 4294967295
                                        _ = v105
                                    }
                                    v106 := GenS12{F0: 18446744073709551615}
                                    v107 := GenS13{F0: true, F1: 989, F2: ("" + "a")}
                                    _ = v106
                                    _ = v107
                                }
                                v45 = (!(!(!((55 != 180) || true))))
                            }
                            _ = v66
                        }
                        var v108 float32 = 1e10
                        _ = len(v35)
                        fmt.Println(v3, v108, v5)
                        _ = len(v35)
                        v109 := GenS14{F0: (^(fi44)), F1: v41}
                        var v110 int = fi44
                        _ = v108
                        _ = v109
                        _ = v110
                    } else {
                        _ = len(v2)
                        v45 = (!((^(fi65)) < fi65))
                    }
                    v111 := GenS15{F0: 1, F1: (^(v41)), F2: 1}
                    v41 -= v41
                    switch ("hello" + ("foo bar" + ("\n\t" + ("foo bar" + ("hello" + ""))))) {
                    case "foo bar":
                        v45 = (!(0 < -1))
                    default:
                        if len(v35) >= 2 {
                            v35 = v35[1:len(v35)]
                        }
                    }
                    _ = v111
                }
                if true {
                    switch ("hello" + (("foo bar" + "\n\t") + "a")) {
                    case "foo bar":
                        fmt.Println(v5, v45, v1)
                    case "\n\t":
                        v112 := GenS16{F0: "hello"}
                        _ = v112
                    }
                    switch "hello" {
                    case "foo bar":
                        if v3 {
                            var v113 string = ("foo bar" + ("hello" + "\n\t"))
                            _ = v113
                        } else {
                            if len(v35) > 0 {
                                v35[0] = fi44
                            }
                        }
                    case "a":
                        if false {
                            var v114 float32 = (1.5 / 1.5)
                            var v115 float32 = 1.5
                            fmt.Println(v45, v3, v5)
                            _ = len(v2)
                            if (830 >= 166) {
                                v2[42] = 0
                            }
                            _ = v114
                            _ = v115
                        } else {
                            v45 = ((fi65 % ((fi44) | 1)) >= (fi44 << ((65535) & 63)))
                            v41 += v41
                        }
                    default:
                        v116 := GenS17{F0: (-(fi65))}
                        var v117 float32 = 0.1
                        _ = len(v2)
                        _ = v116
                        _ = v117
                    }
                    v2[42] = (fi65 ^ 471)
                    var v118 int32 = v41
                    var v119 uint16 = 32767
                    v120 := GenS11{F0: 100}
                    v5 = (128 >= 42)
                    _ = v118
                    _ = v119
                    _ = v120
                } else {
                    switch (fi44 | fi44) {
                    case 42:
                        v121 := GenS7{F0: 127, F1: v1}
                        _ = v121
                    case 9223372036854775807:
                        var v122 int16 = (v1 >> ((473) & 63))
                        _ = v122
                    }
                    fmt.Println(v45, v41)
                    switch fi44 {
                    case 4294967295:
                        v123 := GenS18{F0: 65535, F1: false, F2: 15}
                        _ = v123
                    case 127:
                        if false {
                            var v124 int64 = -558
                            if (42 >= 551) {
                                if (!false) {
                                    if ((572 < 2) && false) {
                                        v125 := GenS10{F0: 673}
                                        switch ("a" + ("\n\t" + "\n\t")) {
                                        case "foo bar":
                                            _ = len(v35)
                                        case "":
                                            var v126 int = 127
                                            _ = v126
                                        default:
                                            if (134 <= 1) {
                                                var v127 int64 = (v124 >> ((18446744073709551615) & 63))
                                                v127 = (v127 + v127)
                                                _ = v127
                                            } else {
                                                if v45 {
                                                    var v128 string = ("hello" + ("hello" + (("" + "foo bar") + "hello")))
                                                    var v129 int8 = -1
                                                    var v130 uint16 = 42
                                                    _ = v128
                                                    _ = v129
                                                    _ = v130
                                                } else {
                                                    delete(v2, 0)
                                                    v131 := GenS15{F0: 531, F1: v41, F2: 65535}
                                                    v41 = v41
                                                    _ = v131
                                                }
                                                switch "" {
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v41 -= v41
                                                    case "hello":
                                                        var v132 float32 = (0.1 - (float32(3.14159) - (-2.25 + -2.25)))
                                                        _ = v132
                                                    default:
                                                        v133 := GenS19{F0: 4294967295}
                                                        _ = v133
                                                    }
                                                case "foo bar":
                                                    var v134 uint16 = 42
                                                    _ = v134
                                                }
                                            }
                                        }
                                        if len(v35) > 0 {
                                            v35[0] = fi65
                                        }
                                        _ = v125
                                    }
                                    var v135 int = 32767
                                    _ = v135
                                }
                            }
                            _ = v124
                        } else {
                            var v136 uint64 = 4294967295
                            _ = len(v35)
                            switch (("foo bar" + "\n\t") + "a") {
                            case "":
                                var v137 uint8 = 1
                                _ = v137
                            case "\n\t":
                                v41 -= (v41 - (v41 + (v41 >> ((42) & 63))))
                            }
                            _ = v136
                        }
                    case 742:
                        v1 = (v1 << ((0) & 63))
                    }
                }
                fmt.Println(v1, v5)
                _ = fi65
            }
            _ = v45
        }
        _ = fi44
    }
    var v138 string = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v138:", v138)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v138
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
