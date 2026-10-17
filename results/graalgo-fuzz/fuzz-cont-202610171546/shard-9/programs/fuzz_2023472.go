package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 bool }
type GenS2 struct { F0 uint8; F1 int8; F2 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int; F2 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 uint16; F1 uint; F2 string }
type GenS15 struct { F0 string }
type GenS16 struct { F0 string; F1 string; F2 int }
type GenS17 struct { F0 bool; F1 int16; F2 int }
type GenS18 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS19 struct { F0 string }
func genFn0() uint {
    if false {
        var v7 uint8 = 255
        _ = v7
    }
    switch "a" {
    case "\n\t":
        if (!(-37 <= 128)) {
            v8 := []int{4294967295, 968, 127, 42}
            _ = v8
        } else {
            var v9 int32 = 2147483647
            _ = v9
        }
    case "hello":
        var v10 uint32 = 530
        _ = v10
    default:
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int16 = 32767
            _ = fi11
            _ = v12
        }
    }
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v13 uint32 = 42
        _ = v13
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            v14 := GenS1{F0: 1, F1: 1, F2: false}
            _ = v14
        case "a":
            v15 := []int{0, 128}
            _ = v15
        case "":
            if true {
                var v16 uint8 = 192
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float32 = 3.14159
                    v18 += v18
                    _ = fi17
                    _ = v18
                }
                _ = v16
            } else {
                v19 := GenS2{F0: 1, F1: 127, F2: 0}
                if (!(899 < 4294967295)) {
                    v20 := GenS1{F0: 127, F1: 4294967295, F2: true}
                    var v21 bool = false
                    var v22 float64 = 0.1
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    if (!(!true)) {
                        var v23 bool = true
                        if (4294967295 == 820) {
                            v23 = v23
                        }
                        var v24 uint64 = 42
                        _ = v23
                        _ = v24
                    }
                    var v25 uint16 = 0
                    v26 := GenS3{F0: 42, F1: 1}
                    _ = v25
                    _ = v26
                }
                _ = v19
            }
        }
    }
    return 42
}
func genFn1(p0 bool) uint32 {
    v27 := GenS4{F0: 255, F1: 2}
    for fi28 := 0; fi28 < 7; fi28++ {
        if (!p0) {
            if p0 {
                p0 = (2147483647 <= 1)
                v29 := GenS1{F0: 42, F1: 953, F2: p0}
                var v30 int8 = 42
                _ = v29
                _ = v30
            } else {
                for fi31 := 0; fi31 < 11; fi31++ {
                    p0 = ((-2.25 / 3.14159) > (1.5 + (0.0 - 3.14159)))
                    v32 := []int{fi28, (^(fi31)), (fi28 >> ((255) & 63))}
                    _ = fi31
                    _ = v32
                }
            }
            if (!(-335 > 0)) {
                var v33 int64 = -9223372036854775808
                var v34 float32 = 3.14159
                _ = v33
                _ = v34
            } else {
                var v35 uint = 0
                if (!(9223372036854775807 > 710)) {
                    var v36 bool = ((p0 && true) || p0)
                    var v37 string = "a"
                    _ = v36
                    _ = v37
                } else {
                    if p0 {
                        fmt.Println(fi28)
                        if p0 {
                            if p0 {
                                v35 = v35
                            }
                            if p0 {
                                v38 := GenS3{F0: 1, F1: 42}
                                p0 = (!(p0 || p0))
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    v40 := []int{fi28, (^(fi39)), 9223372036854775807, (fi39 ^ 1)}
                                    switch (fi28 * fi39) {
                                    case 9223372036854775807:
                                        var v41 int32 = -2147483648
                                        _ = v41
                                    }
                                    fmt.Println(fi28, fi39)
                                    if (false && ((fi39 ^ 32768) == fi39)) {
                                        p0 = (("\n\t" > "") || (2 != 2))
                                        var v42 uint = v35
                                        p0 = p0
                                        switch "hello" {
                                        case "hello":
                                            var v43 int8 = 1
                                            _ = v43
                                        default:
                                            p0 = p0
                                        }
                                        if ((^(v42)) > 274) {
                                            v44 := GenS5{F0: fi39, F1: (v42 << ((1) & 63))}
                                            switch fi39 {
                                            case -302:
                                                p0 = p0
                                            case -2147483648:
                                                v42 = 18
                                            default:
                                                v35 -= genFn0()
                                            }
                                            var v45 int32 = 1
                                            _ = v44
                                            _ = v45
                                        }
                                        _ = v42
                                    } else {
                                        switch "" {
                                        case "hello":
                                            if p0 {
                                                p0 = (65535 != (v35 % ((122) | 1)))
                                                p0 = p0
                                                v35 = (v35 / ((genFn0()) | 1))
                                            } else {
                                                if ((v35 > (^(v35))) || true) {
                                                    p0 = (!(p0 && ((true && p0) || ((-1 >= 354) && (((^(v35)) > (v35 + genFn0())) && (((("\n\t" + (("foo bar" + "a") + "a")) + (("hello" + "foo bar") + "a")) + "a") < "\n\t"))))))
                                                }
                                                if true {
                                                    var v46 int64 = 133
                                                    _ = v46
                                                }
                                            }
                                        case "\n\t":
                                            switch ("foo bar" + ("\n\t" + "hello")) {
                                            case "":
                                                v47 := GenS1{F0: 126, F1: 2, F2: p0}
                                                _ = v47
                                            default:
                                                if p0 {
                                                    if (42 > 65535) {
                                                        var v48 uint8 = 166
                                                        _ = v48
                                                    } else {
                                                        v49 := GenS5{F0: fi39, F1: (v35 % (((^(v35))) | 1))}
                                                        var v50 int16 = -32768
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    if len(v40) >= 2 {
                                                        v40 = v40[1:len(v40)]
                                                    }
                                                    fmt.Println(v35, fi28)
                                                }
                                                p0 = p0
                                            }
                                        default:
                                            var v51 uint64 = 18446744073709551615
                                            _ = v51
                                        }
                                    }
                                    _ = fi39
                                    _ = v40
                                }
                                _ = v38
                            }
                            fmt.Println(p0, fi28)
                        }
                        p0 = false
                    }
                }
                _ = v35
            }
        }
        for fi52 := 0; fi52 < 6; fi52++ {
            var v53 int = (fi52 % ((0) | 1))
            p0 = p0
            _ = fi52
            _ = v53
        }
        p0 = (p0 && (127 > 117))
        _ = fi28
    }
    for fi54 := 0; fi54 < 14; fi54++ {
        v55 := GenS4{F0: 4294967295, F1: 32767}
        _ = fi54
        _ = v55
    }
    v56 := GenS1{F0: 100, F1: 4294967295, F2: p0}
    _ = p0
    _ = v27
    _ = v56
    return 2
}

func main() {
    v1 := []string{"a", "\n\t"}
    var v2 int32 = -516
    v3 := make(map[string]int)
    var v4 uint = 625
    var v5 float64 = 0.0
    v6 := []uint32{387, 293}
    if true {
        fmt.Println(v5)
        v57 := GenS3{F0: 4294967295, F1: 1}
        var v58 int64 = -157
        _ = v57
        _ = v58
    } else {
        for fi59 := 0; fi59 < 20; fi59++ {
            if (1 > 2147483647) {
                v5 = v5
                for fi60 := 0; fi60 < 1; fi60++ {
                    v5 += v5
                    var v61 uint32 = 799
                    var v62 uint16 = 1
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                if false {
                    var v63 uint64 = 1
                    _ = v63
                }
            } else {
                if (false || true) {
                    v1 = append(v1, ("a" + "a"))
                } else {
                    for fi64 := 0; fi64 < 11; fi64++ {
                        v65 := GenS6{F0: 127}
                        v2 = (^(v2))
                        _ = fi64
                        _ = v65
                    }
                    _ = len(v3)
                    v66 := GenS7{F0: 1, F1: 0}
                    _ = v66
                }
                if false {
                    v67 := GenS3{F0: 0, F1: v2}
                    _ = v67
                } else {
                    v68 := GenS8{F0: (v4 & v4), F1: "\n\t"}
                    _ = v68
                }
                if false {
                    if false {
                        var v69 float32 = (((0.1 + float32((1e10 / v5))) - 3.14159) + 1.5)
                        switch "" {
                        case "":
                            for fi70 := 0; fi70 < 17; fi70++ {
                                v69 -= ((v69 - (0.1 * 3.14159)) + (v69 / v69))
                                if true {
                                    var v71 uint8 = 255
                                    var v72 bool = false
                                    if (-908 > -807) {
                                        switch fi70 {
                                        case 32767:
                                            if true {
                                                v72 = ((621 != 0) || (-1 <= 404))
                                                fmt.Println(fi59, v4)
                                                switch fi70 {
                                                case 582:
                                                    v73 := GenS2{F0: v71, F1: 0, F2: (^(v4))}
                                                    switch "a" {
                                                    case "a":
                                                        fmt.Println(fi59, v72)
                                                    case "\n\t":
                                                        if (!("a" >= ("a" + ("hello" + ("a" + ("" + "\n\t")))))) {
                                                            if (1 > (^(fi59))) {
                                                                v69 = v69
                                                                fmt.Println(v71, v4)
                                                                v74 := v3[("a" + "")]
                                                                if true {
                                                                    v3[""] = fi70
                                                                    v69 = ((-2.25 / 0.1) * v69)
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v69 -= (v69 * 3.14159)
                                                                    }
                                                                } else {
                                                                    fmt.Println(v2, fi59, v72)
                                                                }
                                                                v72 = v72
                                                                _ = v74
                                                            } else {
                                                                v72 = (!true)
                                                                switch "" {
                                                                case "hello":
                                                                    if (v2 > v2) {
                                                                        if (!v72) {
                                                                            v75 := GenS2{F0: v71, F1: -30, F2: (v4 << ((v4) & 63))}
                                                                            switch ("a" + ("hello" + ("\n\t" + "\n\t"))) {
                                                                            case "hello":
                                                                                var v76 int8 = 100
                                                                                _ = v76
                                                                            default:
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ("a" + ("" + "a"))
                                                                                    }
                                                                                case "a":
                                                                                    var v77 uint8 = v71
                                                                                    _ = v77
                                                                                case "":
                                                                                    fmt.Println(v71)
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    var v78 uint16 = 2
                                                                                    _ = v78
                                                                                default:
                                                                                    v72 = (1000 <= 0)
                                                                                }
                                                                            }
                                                                            var v79 uint32 = 314
                                                                            var v80 uint32 = (v79 << (((v4 - v4)) & 63))
                                                                            _ = v75
                                                                            _ = v79
                                                                            _ = v80
                                                                        } else {
                                                                            _ = len(v1)
                                                                        }
                                                                        v81 := GenS9{F0: "\n\t"}
                                                                        _ = v81
                                                                    }
                                                                default:
                                                                    fmt.Println(v5, v5)
                                                                }
                                                                if (((("\n\t" + "foo bar") + "") >= ("" + ("\n\t" + ("\n\t" + "\n\t")))) && v72) {
                                                                    if (41 < 12) {
                                                                        v71 += (v71 + (v71 / ((v71) | 1)))
                                                                    } else {
                                                                        v3["a"] = -2147483648
                                                                        v82 := GenS2{F0: v71, F1: 89, F2: v4}
                                                                        if false {
                                                                            var v83 int8 = -128
                                                                            var v84 uint64 = 2
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            fmt.Println(v71)
                                                                            var v85 int32 = (v2 | v2)
                                                                            var v86 int16 = -32768
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                        v5 += (v5 - v5)
                                                                        _ = v82
                                                                    }
                                                                }
                                                                if (uint32(v4) > 836) {
                                                                    if (v69 >= (v69 * v69)) {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            if false {
                                                                                var v87 uint8 = (v71 / (((v71 << ((720) & 63))) | 1))
                                                                                var v88 int64 = -1
                                                                                v3[""] = fi59
                                                                                _ = v87
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                        if (v69 < ((float32((0.1 - (v5 / v5))) * ((v69 + 1e10) - (v69 * v69))) * 0.0)) {
                                                                            var v89 string = "\n\t"
                                                                            v5 = (v5 - 1.5)
                                                                            _ = v89
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v90 float64 = (((v5 * v5) + 3.14159) + v5)
                                                                    var v91 int = fi59
                                                                    if (!true) {
                                                                        fmt.Println(v4, v71, v72)
                                                                        var v92 int = fi59
                                                                        _ = v92
                                                                    }
                                                                    _ = v90
                                                                    _ = v91
                                                                }
                                                            }
                                                            var v93 int16 = 32767
                                                            var v94 int32 = -600
                                                            _ = v93
                                                            _ = v94
                                                        } else {
                                                            fmt.Println(fi59)
                                                            v95 := GenS5{F0: (fi70 | fi59), F1: 255}
                                                            v5 -= (1.5 - (-2.25 * 0.1))
                                                            if (1 < 0) {
                                                                v96 := GenS10{F0: 1000, F1: (fi70 >> ((1) & 63)), F2: 454}
                                                                v4 += (v4 & v4)
                                                                _ = v96
                                                            }
                                                            if v72 {
                                                                v69 = (v69 * 3.14159)
                                                                fmt.Println(v71)
                                                                var v97 int8 = 127
                                                                _ = v97
                                                            }
                                                            _ = v95
                                                        }
                                                    default:
                                                        v4 += 18446744073709551615
                                                    }
                                                    _ = v73
                                                case 2147483647:
                                                    v98 := v3[("" + "")]
                                                    _ = v98
                                                default:
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(fi70, v71)
                                                    }
                                                    var v99 uint8 = v71
                                                    v100 := GenS8{F0: 42, F1: "foo bar"}
                                                    _ = v100
                                                    _ = v99
                                                }
                                                v4 = 65535
                                                if (fi70 == (fi70 & fi59)) {
                                                    v101 := GenS5{F0: fi70, F1: v4}
                                                    _ = v101
                                                }
                                            } else {
                                                v5 = 1.5
                                                fmt.Println(v72, v72, v4)
                                                var v102 int = (fi70 ^ (fi59 * (fi70 + fi59)))
                                                _ = v102
                                            }
                                        case 514:
                                            v69 += v69
                                        default:
                                            v4 += (v4 & (v4 & v4))
                                            v2 = 42
                                            var v103 uint8 = (^(v71))
                                            _ = v103
                                        }
                                        v71 -= v71
                                        v4 = v4
                                    } else {
                                        var v104 int16 = 1000
                                        v104 = v104
                                        if (-87 >= 127) {
                                            fmt.Println(v2)
                                            var v105 string = ("" + "hello")
                                            v3["a"] = 9223372036854775807
                                            v106 := GenS11{F0: (v104 * v104)}
                                            if (707 < uint64(v71)) {
                                                if ((v71 - (v71 ^ 128)) > (v71 & (v71 + (^(v71))))) {
                                                    var v107 uint = (^(v4))
                                                    var v108 int64 = 42
                                                    if ((v104 ^ v104) <= v104) {
                                                        var v109 string = "a"
                                                        _ = v109
                                                    }
                                                    _ = v107
                                                    _ = v108
                                                } else {
                                                    v110 := GenS11{F0: v104}
                                                    _ = v110
                                                }
                                            }
                                            var v111 int = (-(fi70))
                                            _ = v105
                                            _ = v106
                                            _ = v111
                                        } else {
                                            var v112 int32 = v2
                                            if ((((v2 >= -694) && (820 <= -9223372036854775808)) && v72) && true) {
                                                fmt.Println(v69, v5, fi59)
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v113 int = fi70
                                                    _ = v113
                                                default:
                                                    v114 := GenS12{F0: 6}
                                                    _ = v114
                                                }
                                                fmt.Println(v5)
                                                var v115 int64 = 0
                                                _ = v115
                                            } else {
                                                v3["foo bar"] = (fi59 / ((int(v2)) | 1))
                                                _ = len(v6)
                                                if (!v72) {
                                                    var v116 uint8 = v71
                                                    _ = v116
                                                }
                                                var v117 uint32 = 42
                                                _ = v117
                                            }
                                            v71 = v71
                                            _ = v112
                                        }
                                        delete(v3, ("hello" + "a"))
                                        v118 := GenS13{F0: (fi70 % ((4294967295) | 1))}
                                        _ = v104
                                        _ = v118
                                    }
                                    v71 = 1
                                    _ = v71
                                    _ = v72
                                }
                                var v119 int16 = -521
                                _ = fi70
                                _ = v119
                            }
                        case "a":
                            for fi120 := 0; fi120 < 15; fi120++ {
                                fmt.Println(v5)
                                v4 -= v4
                                _ = fi120
                            }
                        }
                        switch "\n\t" {
                        case "a":
                            var v121 uint8 = 24
                            _ = v121
                        case "\n\t":
                            fmt.Println(v5, v5)
                        case "hello":
                            var v122 float32 = 1.5
                            _ = v122
                        default:
                            fmt.Println(v4, fi59, v5)
                        }
                        _ = v69
                    }
                    v4 += (v4 | (v4 % ((genFn0()) | 1)))
                    var v123 string = "\n\t"
                    _ = v123
                } else {
                    var v124 int16 = -1
                    switch fi59 {
                    case -16:
                        v2 += (v2 + (v2 >> (((v4 % (((v4 | 18446744073709551615)) | 1))) & 63)))
                    default:
                        v5 = v5
                    }
                    _ = v124
                }
            }
            _ = fi59
        }
        if ((false || false) || false) {
            for fi125 := 0; fi125 < 9; fi125++ {
                v6 = append(v6, genFn1(true))
                for fi126 := 0; fi126 < 8; fi126++ {
                    fmt.Println(v5, v2)
                    var v127 float32 = (1e10 + ((-2.25 - (float32(v5) / 1.5)) * (0.0 + (0.0 - 1e10))))
                    switch fi125 {
                    case 32767:
                        v127 += ((v127 * (v127 * float32(v5))) + v127)
                    }
                    v4 = v4
                    _ = fi126
                    _ = v127
                }
                _ = fi125
            }
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch 32768 {
    case -2147483648:
        v2 = (-(v2))
    case 32768:
        fmt.Println(v4, v4)
    case -564:
        var v128 uint16 = 0
        _ = v128
    }
    var v129 int16 = 1
    for fi130 := 0; fi130 < 4; fi130++ {
        if true {
            for fi131 := 0; fi131 < 6; fi131++ {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                var v132 int16 = (v129 & (v129 + 219))
                _ = fi131
                _ = v132
            }
            if (255 > 1) {
                v5 += (1.5 + 1e10)
            } else {
                delete(v3, "\n\t")
                v133 := GenS7{F0: 4294967295, F1: v2}
                fmt.Println(fi130, v129, v5)
                _ = v133
            }
        }
        for fi134 := 0; fi134 < 5; fi134++ {
            v129 = v129
            if (v2 >= (^(v2))) {
                var v135 float32 = 0.0
                var v136 float64 = (1e10 + v5)
                if (0 != 9223372036854775807) {
                    v136 += v5
                    switch "foo bar" {
                    case "\n\t":
                        v137 := v3["hello"]
                        _ = v137
                    case "":
                        switch "a" {
                        case "":
                            v5 = v5
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                    case "hello":
                        v2 -= (^(v2))
                    }
                } else {
                    if false {
                        v138 := GenS14{F0: 2, F1: genFn0(), F2: ("" + ((("a" + "foo bar") + "") + "a"))}
                        _ = v138
                    } else {
                        if false {
                            var v139 uint8 = 191
                            var v140 int16 = -32768
                            _ = v139
                            _ = v140
                        }
                    }
                }
                _ = v135
                _ = v136
            } else {
                fmt.Println(fi134, v2)
            }
            _ = fi134
        }
        fmt.Println(v129)
        _ = fi130
    }
    v141 := GenS15{F0: ""}
    v4 = (v4 + v4)
    if false {
        for fi142 := 0; fi142 < 1; fi142++ {
            v143 := GenS4{F0: v4, F1: 0}
            var v144 int32 = (v2 ^ v2)
            fmt.Println(v5, fi142, v2)
            _ = fi142
            _ = v143
            _ = v144
        }
        for fi145 := 0; fi145 < 6; fi145++ {
            if (false && false) {
                fmt.Println(v129)
                for fi146 := 0; fi146 < 2; fi146++ {
                    v129 += 827
                    var v147 uint64 = 2
                    switch (^(fi145)) {
                    case 127:
                        v129 += (v129 / ((v129) | 1))
                    default:
                        fmt.Println(fi145, fi146, v147)
                    }
                    _ = fi146
                    _ = v147
                }
            } else {
                for fi148 := 0; fi148 < 5; fi148++ {
                    v2 += 1
                    fmt.Println(fi145, v129)
                    _ = fi148
                }
                var v149 float64 = float64((0.1 / (0.1 / 1.5)))
                _ = v149
            }
            var v150 uint8 = 127
            v2 = (v2 + 1)
            _ = fi145
            _ = v150
        }
        switch "\n\t" {
        case "":
            var v151 float32 = 1.5
            _ = v151
        case "a":
            fmt.Println(v129)
        case "foo bar":
            for fi152 := 0; fi152 < 11; fi152++ {
                var v153 int64 = 0
                switch (^(fi152)) {
                case -9223372036854775808:
                    v5 = float64(0.0)
                case 127:
                    for fi154 := 0; fi154 < 14; fi154++ {
                        if (2 == 2) {
                            v155 := GenS16{F0: "", F1: ("a" + ("a" + "a")), F2: (fi152 & fi152)}
                            v3["hello"] = (fi154 << ((v4) & 63))
                            _ = v155
                        } else {
                            switch (fi154 & (fi152 << ((v4) & 63))) {
                            case 1:
                                var v156 uint8 = 2
                                _ = v156
                            case 40:
                                fmt.Println(v5)
                            case 4294967295:
                                if ((^(v4)) >= (v4 ^ v4)) {
                                    switch (fi152 << ((v4) & 63)) {
                                    case 9223372036854775807:
                                        v3["a"] = (fi152 + (fi154 / ((-1) | 1)))
                                    case 42:
                                        var v157 float64 = v5
                                        _ = v157
                                    case 4294967295:
                                        v129 += (v129 >> ((genFn0()) & 63))
                                    default:
                                        switch fi152 {
                                        case 128:
                                            switch (fi154 << ((v4) & 63)) {
                                            case -1:
                                                fmt.Println(v129)
                                            }
                                        default:
                                            if true {
                                                v5 = v5
                                                v158 := GenS2{F0: 127, F1: 117, F2: genFn0()}
                                                _ = v158
                                            } else {
                                                if ((v153 & (v153 + -493)) == v153) {
                                                    v153 -= v153
                                                }
                                                v129 += v129
                                                if true {
                                                    var v159 int64 = (v153 ^ (^(v153)))
                                                    _ = v159
                                                } else {
                                                    if (0 > 42) {
                                                        _ = len(v3)
                                                    } else {
                                                        switch ("\n\t" + ((("\n\t" + "hello") + "\n\t") + "a")) {
                                                        case "a":
                                                            switch ("a" + "foo bar") {
                                                            case "\n\t":
                                                                fmt.Println(v153, fi154)
                                                                v153 = 9223372036854775807
                                                            case "foo bar":
                                                                v2 += (v2 + (v2 << (((v4 / (((v4 & genFn0())) | 1))) & 63)))
                                                            case "":
                                                                fmt.Println(fi154, fi154)
                                                            default:
                                                                var v160 bool = (!(!true))
                                                                _ = v160
                                                            }
                                                        default:
                                                            fmt.Println(v2)
                                                        }
                                                        v161 := GenS17{F0: (!(!(!(113 >= 42)))), F1: v129, F2: 42}
                                                        v162 := v3[""]
                                                        _ = v161
                                                        _ = v162
                                                    }
                                                    switch "\n\t" {
                                                    case "a":
                                                        delete(v3, "a")
                                                        var v163 int = (fi152 / ((int(v2)) | 1))
                                                        fmt.Println(v5)
                                                        _ = v163
                                                    case "hello":
                                                        v164 := v3["foo bar"]
                                                        _ = v164
                                                    }
                                                    v2 = (-(v2))
                                                }
                                            }
                                        }
                                    }
                                    v165 := GenS18{F0: 65535, F1: 42, F2: -128}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v165
                                } else {
                                    switch ("hello" + (("hello" + "foo bar") + "\n\t")) {
                                    case "a":
                                        switch "" {
                                        case "a":
                                            v166 := GenS18{F0: 422, F1: 571, F2: 25}
                                            _ = v166
                                        case "foo bar":
                                            var v167 float64 = 0.0
                                            _ = v167
                                        case "\n\t":
                                            var v168 bool = (true || (((-2.25 + 0.0) > (3.14159 - (-2.25 + -2.25))) || (uint32(v4) > 4294967295)))
                                            _ = v168
                                        }
                                    case "foo bar":
                                        v129 += (^(v129))
                                    case "":
                                        switch "" {
                                        case "hello":
                                            fmt.Println(fi152)
                                        case "":
                                            switch (fi154 * fi152) {
                                            case 2147483647:
                                                var v169 float32 = -2.25
                                                _ = v169
                                            case 0:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v170 int64 = v153
                                                    _ = v170
                                                }
                                            default:
                                                var v171 uint16 = 32767
                                                _ = v171
                                            }
                                        case "\n\t":
                                            v172 := GenS7{F0: 18446744073709551615, F1: -33}
                                            _ = v172
                                        default:
                                            var v173 uint = (v4 << ((genFn0()) & 63))
                                            _ = v173
                                        }
                                    default:
                                        v2 = (^(v2))
                                    }
                                    var v174 float32 = ((0.0 + 1.5) - 1e10)
                                    _ = v174
                                }
                                var v175 uint8 = 137
                                _ = v175
                            }
                        }
                        fmt.Println(v4)
                        _ = fi154
                    }
                case 42:
                    var v176 bool = (!(!true))
                    _ = v176
                default:
                    _ = len(v3)
                }
                v177 := GenS17{F0: ((-128 < -43) || ((v153 - (v153 << (((v4 | v4)) & 63))) > -382)), F1: 32767, F2: 128}
                _ = fi152
                _ = v153
                _ = v177
            }
        }
    }
    v129 = v129
    var v178 uint = (v4 ^ (v4 >> ((40) & 63)))
    if true {
        for fi179 := 0; fi179 < 5; fi179++ {
            var v180 int8 = -79
            _ = fi179
            _ = v180
        }
        if (v2 >= v2) {
            v181 := GenS19{F0: ("hello" + ("" + "hello"))}
            if (128 >= 0) {
                _ = len(v6)
                var v182 int16 = v129
                _ = v182
            }
            v4 = genFn0()
            _ = v181
        }
    } else {
        fmt.Println(v178, v129, v5)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v129:", v129)
    fmt.Println("v178:", v178)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v129
    _ = v141
    _ = v178
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
