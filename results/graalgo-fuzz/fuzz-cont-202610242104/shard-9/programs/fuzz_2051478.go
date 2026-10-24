package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint; F2 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 int32; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint64; F2 string }
type GenS7 struct { F0 int; F1 int32; F2 uint }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 uint; F1 bool }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int8 }
func genFn0(p0 int16, p1 uint8) string {
    for fi5 := 0; fi5 < 2; fi5++ {
        p0 -= (p0 - p0)
        p1 += (p1 ^ 127)
        switch fi5 {
        case 127:
            var v6 int8 = 76
            _ = v6
        }
        _ = fi5
    }
    p1 += (p1 & 128)
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 uint = 255
        if (!true) {
            fmt.Println(p1, fi7, p1)
            for fi9 := 0; fi9 < 1; fi9++ {
                v10 := []int{1, 1, (^(fi9))}
                _ = fi9
                _ = v10
            }
        } else {
            v11 := GenS1{F0: -275, F1: (p0 / ((p0) | 1)), F2: 135}
            v12 := []int{}
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    switch 4294967295 {
    case 547:
        p1 = 2
    case 659:
        if false {
            for fi13 := 0; fi13 < 1; fi13++ {
                for fi14 := 0; fi14 < 7; fi14++ {
                    v15 := GenS1{F0: (p0 ^ 140), F1: -32768, F2: 0}
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    v17 := []int{(fi16 ^ (fi16 * -173)), (fi13 + int(p0)), (^(fi13)), fi13}
                    if true {
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                        var v18 int64 = 425
                        v19 := GenS1{F0: 59, F1: (p0 & (p0 + -32768)), F2: 2}
                        if false {
                            p0 = (p0 * p0)
                        } else {
                            if (127 >= 1) {
                                var v20 int64 = v18
                                if false {
                                    var v21 uint16 = 904
                                    var v22 string = ("\n\t" + "\n\t")
                                    _ = v21
                                    _ = v22
                                } else {
                                    fmt.Println(p1, v20)
                                }
                                _ = v20
                            }
                        }
                        if (0 > 100) {
                            if true {
                                if len(v17) > 0 {
                                    v17[0] = fi16
                                }
                                v18 += v18
                                var v23 int8 = -1
                                p1 -= (p1 & p1)
                                _ = len(v17)
                                _ = v23
                            }
                        }
                        var v24 uint32 = 753
                        _ = v18
                        _ = v19
                        _ = v24
                    } else {
                        if false {
                            v25 := GenS1{F0: p0, F1: 635, F2: 0}
                            p1 -= 1
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            p0 += 1
                            if len(v17) > 0 {
                                v17[0] = (fi16 % ((9223372036854775807) | 1))
                            }
                            if false {
                                if false {
                                    switch fi16 {
                                    case 32768:
                                        var v26 int8 = -128
                                        _ = v26
                                    case 9223372036854775807:
                                        v17 = append(v17, 1)
                                    case 4294967295:
                                        p1 -= p1
                                    default:
                                        v17 = append(v17, fi13)
                                    }
                                    var v27 int16 = p0
                                    if len(v17) > 0 {
                                        v17[0] = fi13
                                    }
                                    p1 = (p1 % ((p1) | 1))
                                    v28 := GenS2{F0: 0}
                                    switch fi16 {
                                    case 128:
                                        v29 := GenS3{F0: 1, F1: 42, F2: -1}
                                        _ = v29
                                    default:
                                        if (0 < 0) {
                                            v27 = 32767
                                        }
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                _ = len(v17)
                            } else {
                                p1 += 195
                                v30 := GenS4{F0: "\n\t"}
                                var v31 uint8 = 2
                                _ = v30
                                _ = v31
                            }
                            var v32 uint8 = p1
                            var v33 uint = 65535
                            _ = v25
                            _ = v32
                            _ = v33
                        } else {
                            fmt.Println(p1, p1, p0)
                            fmt.Println(p1, fi13, p1)
                        }
                        var v34 float32 = (1.5 * 1e10)
                        p1 = (p1 ^ 127)
                        _ = v34
                    }
                    _ = fi16
                    _ = v17
                }
                _ = fi13
            }
        }
    default:
        for fi35 := 0; fi35 < 8; fi35++ {
            switch fi35 {
            case -779:
                var v36 string = "a"
                _ = v36
            }
            p1 = p1
            for fi37 := 0; fi37 < 8; fi37++ {
                v38 := []int{(fi37 % ((0) | 1))}
                _ = fi37
                _ = v38
            }
            _ = fi35
        }
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = ((((0.1 - -2.25) + float64((((0.0 + 3.14159) / 1e10) * 1.5))) + float64(-2.25)) / ((0.1 / 1.0) / 3.14159))
    v3 := make(map[string]int)
    var v4 int64 = 0
    var v39 float64 = v2
    if (-1 != -785) {
        var v40 uint = 737
        var v41 float32 = (0.1 - 1.5)
        _ = v40
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        if true {
            v39 -= 0.1
        } else {
            v39 += 1e10
            v43 := []int{(^(fi42)), fi42, int(v4), (^(fi42))}
            v44 := GenS4{F0: "a"}
            _ = v43
            _ = v44
        }
        v45 := v1[""]
        delete(v1, "foo bar")
        _ = fi42
        _ = v45
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v2, v2)
    case "":
        v2 = v39
    default:
        v2 -= v2
    }
    v46 := []int{127}
    var v47 bool = (false && (32767 == -1))
    for fi48 := 0; fi48 < 17; fi48++ {
        delete(v1, ("\n\t" + "hello"))
        v2 += v39
        v49 := GenS1{F0: 632, F1: 1000, F2: 4294967295}
        _ = fi48
        _ = v49
    }
    switch "\n\t" {
    case "foo bar":
        for fi50 := 0; fi50 < 8; fi50++ {
            for fi51 := 0; fi51 < 18; fi51++ {
                var v52 uint16 = 862
                if (4294967295 < 686) {
                    v2 -= (v39 * float64(0.0))
                    v2 += (v2 / v2)
                    _ = len(v46)
                }
                var v53 int32 = -414
                v54 := GenS1{F0: 42, F1: 42, F2: 295}
                if (!(134 < -852)) {
                    v55 := GenS5{F0: v4, F1: v53, F2: v52}
                    v56 := GenS5{F0: v4, F1: v53, F2: v52}
                    var v57 int8 = 100
                    _ = v55
                    _ = v56
                    _ = v57
                }
                _ = fi51
                _ = v52
                _ = v53
                _ = v54
            }
            var v58 string = (("hello" + "a") + ("foo bar" + ("a" + ("a" + ("a" + ("" + ("a" + (((("foo bar" + "hello") + "\n\t") + ("hello" + "foo bar")) + "foo bar"))))))))
            fmt.Println(v4, v2, v47)
            _ = fi50
            _ = v58
        }
    }
    fmt.Println(v47)
    var v59 bool = false
    for fi60 := 0; fi60 < 4; fi60++ {
        for fi61 := 0; fi61 < 15; fi61++ {
            var v62 uint64 = 4294967295
            switch "" {
            case "hello":
                if len(v46) > 0 {
                    v46[0] = (fi60 << ((822) & 63))
                }
            }
            _ = fi61
            _ = v62
        }
        for fi63 := 0; fi63 < 12; fi63++ {
            if ("a" == "") {
                if v47 {
                    if ("\n\t" > "\n\t") {
                        if len(v46) >= 2 {
                            v46 = v46[1:len(v46)]
                        }
                        if v47 {
                            fmt.Println(v39, fi60)
                            if len(v46) >= 2 {
                                v46 = v46[1:len(v46)]
                            }
                            switch (^(fi63)) {
                            case -613:
                                if (v2 > v2) {
                                    switch fi63 {
                                    case -52:
                                        _ = len(v1)
                                    case -9223372036854775808:
                                        v4 += v4
                                    case 421:
                                        if v59 {
                                            v64 := GenS2{F0: 2}
                                            v2 -= ((-2.25 * float64(((0.0 - 1e10) + float32(v2)))) + ((((v2 + v39) * v39) * -2.25) + v39))
                                            _ = v64
                                        }
                                        var v65 int64 = 1
                                        fmt.Println(fi60, v39)
                                        _ = v65
                                    }
                                    v66 := GenS2{F0: 2}
                                    _ = v66
                                } else {
                                    v4 += (v4 ^ v4)
                                    v39 += v2
                                    v47 = false
                                }
                            case 390:
                                var v67 int16 = -32768
                                _ = v67
                            default:
                                var v68 uint16 = 65535
                                _ = v68
                            }
                        }
                        if len(v46) >= 2 {
                            v46 = v46[1:len(v46)]
                        }
                    }
                    v39 -= float64((float32((0.0 - (float64(1.5) / v2))) - 3.14159))
                    var v69 uint32 = 65535
                    _ = v69
                }
                if ((((("hello" + "foo bar") + ("hello" + ("\n\t" + (("foo bar" + "\n\t") + (("hello" + ("foo bar" + "foo bar")) + ((("hello" + "") + "foo bar") + "hello")))))) + "hello") + "") == "hello") {
                    v59 = false
                    v70 := GenS2{F0: 0}
                    var v71 uint32 = 4294967295
                    _ = v70
                    _ = v71
                }
                var v72 uint32 = 1
                _ = v72
            }
            var v73 int64 = v4
            v4 += v4
            _ = fi63
            _ = v73
        }
        _ = fi60
    }
    fmt.Println(v4)
    for fi74 := 0; fi74 < 8; fi74++ {
        delete(v3, "foo bar")
        _ = fi74
    }
    v59 = false
    for fi75 := 0; fi75 < 5; fi75++ {
        if (!(v47 && true)) {
            var v76 int64 = (v4 & 9223372036854775807)
            if len(v46) > 0 {
                v46[0] = fi75
            }
            _ = len(v3)
            _ = v76
        }
        switch "\n\t" {
        case "a":
            var v77 uint64 = 425
            _ = v77
        case "foo bar":
            fmt.Println(v2, v4, fi75)
        default:
            for fi78 := 0; fi78 < 7; fi78++ {
                v2 += v2
                switch (fi78 + fi78) {
                case 0:
                    var v79 int16 = -71
                    _ = v79
                default:
                    switch fi75 {
                    case 1:
                        var v80 int8 = 0
                        _ = v80
                    case -1:
                        v81 := GenS4{F0: "a"}
                        _ = v81
                    }
                }
                v39 = (v2 / 0.1)
                v2 += (1.5 - v39)
                var v82 uint = 45
                v83 := GenS1{F0: -32768, F1: -424, F2: 2}
                switch "foo bar" {
                case "hello":
                    _ = len(v3)
                }
                v84 := GenS6{F0: 32767, F1: 1, F2: "\n\t"}
                var v85 uint16 = 42
                v47 = false
                switch (-(fi75)) {
                case 674:
                    if false {
                        switch "foo bar" {
                        case "":
                            var v86 int32 = 2147483647
                            _ = v86
                        default:
                            v59 = false
                        }
                        fmt.Println(v39)
                        v87 := GenS4{F0: ("foo bar" + "hello")}
                        _ = v87
                    } else {
                        v88 := GenS7{F0: (fi75 & fi75), F1: 162, F2: (^(v82))}
                        if (!v47) {
                            switch "a" {
                            case "":
                                v2 += (0.0 * ((0.0 * v2) + (v39 + (((0.0 * ((v39 - v39) / v39)) + v2) - v39))))
                            }
                            v89 := GenS1{F0: -624, F1: -1, F2: 752}
                            v59 = ((285 < -747) || false)
                            switch "foo bar" {
                            case "foo bar":
                                if (!v47) {
                                    if v59 {
                                        v90 := GenS6{F0: -82, F1: 0, F2: "a"}
                                        if v47 {
                                            if (("" + "hello") > ((((((("a" + "hello") + ("a" + "a")) + ("a" + "\n\t")) + (("a" + "\n\t") + "\n\t")) + ("foo bar" + "")) + ("foo bar" + (("\n\t" + "") + "\n\t"))) + "hello")) {
                                                switch "" {
                                                case "a":
                                                    v91 := GenS1{F0: 42, F1: -156, F2: 632}
                                                    _ = v91
                                                }
                                                v92 := v3[("a" + "foo bar")]
                                                _ = v92
                                            } else {
                                                v93 := GenS8{F0: fi75, F1: (1e10 == (float32(v39) / (float32(v39) * 1e10)))}
                                                v2 -= (float64(float32(0.0)) + 1.5)
                                                v94 := GenS2{F0: 4294967295}
                                                if true {
                                                    var v95 float64 = v39
                                                    v4 = int64(fi75)
                                                    v47 = (!v47)
                                                    _ = v95
                                                } else {
                                                    var v96 float64 = 3.14159
                                                    _ = v96
                                                }
                                                var v97 string = "\n\t"
                                                _ = v93
                                                _ = v94
                                                _ = v97
                                            }
                                        }
                                        v98 := GenS9{F0: (v82 / ((v82) | 1)), F1: (-9223372036854775808 > (v4 % ((int64(fi78)) | 1)))}
                                        v99 := GenS2{F0: 692}
                                        switch (fi78 * fi78) {
                                        case -885:
                                            if len(v46) > 0 {
                                                v46[0] = fi75
                                            }
                                        case 9223372036854775807:
                                            fmt.Println(v59, v39, v82)
                                            var v100 uint32 = 65535
                                            _ = len(v1)
                                            _ = v100
                                        case 42:
                                            v85 = (^(v85))
                                        default:
                                            if v47 {
                                                var v101 uint64 = 4294967295
                                                _ = v101
                                            }
                                        }
                                        fmt.Println(fi78)
                                        switch "a" {
                                        case "foo bar":
                                            v39 += v39
                                        }
                                        fmt.Println(v85, v82)
                                        if len(v46) >= 2 {
                                            v46 = v46[1:len(v46)]
                                        }
                                        var v102 uint16 = 155
                                        _ = v102
                                        _ = v90
                                        _ = v98
                                        _ = v99
                                    } else {
                                        v103 := GenS9{F0: (v82 * 2), F1: false}
                                        if (v4 <= v4) {
                                            v2 = v2
                                            _ = len(v46)
                                        } else {
                                            v2 += (((v39 / v2) + ((v2 * v2) / 1.0)) / v39)
                                            v85 = v85
                                            if true {
                                                var v104 uint16 = (v85 / ((v85) | 1))
                                                v105 := v3["\n\t"]
                                                _ = v104
                                                _ = v105
                                            } else {
                                                v47 = v47
                                            }
                                            if (true && true) {
                                                var v106 uint = uint(v85)
                                                if ((0.0 * (1e10 / (((1.5 + (float32(v2) + 3.14159)) * -2.25) * 0.0))) > float32(v39)) {
                                                    v82 += v106
                                                    var v107 float32 = -2.25
                                                    _ = v107
                                                }
                                                _ = v106
                                            }
                                            if len(v46) >= 2 {
                                                v46 = v46[1:len(v46)]
                                            }
                                            if v47 {
                                                switch int(v85) {
                                                case -9223372036854775808:
                                                    v85 += v85
                                                case 2147483647:
                                                    var v108 bool = false
                                                    _ = v108
                                                case 1:
                                                    var v109 uint8 = 1
                                                    _ = v109
                                                }
                                            } else {
                                                if (52 <= 255) {
                                                    if (("hello" + ("hello" + "")) < "foo bar") {
                                                        v4 = v4
                                                    }
                                                    v46 = append(v46, (fi75 * 0))
                                                    fmt.Println(v85, fi78, fi75)
                                                }
                                                var v110 int8 = 42
                                                _ = v110
                                            }
                                        }
                                        var v111 bool = (v2 < v2)
                                        var v112 int = int(v4)
                                        _ = v103
                                        _ = v111
                                        _ = v112
                                    }
                                    if ((v47 && (2 == 18446744073709551615)) && (!v47)) {
                                        v39 -= float64((0.1 - (0.0 / 1.5)))
                                    } else {
                                        var v113 uint8 = 42
                                        if (false || (v47 && ((!(!(4294967295 != 966))) && (((0.0 - 0.0) / (float32((1e10 * ((float64(1e10) / v39) - v2))) - -2.25)) >= -2.25)))) {
                                            v114 := GenS10{F0: 1000, F1: 18446744073709551615}
                                            _ = v114
                                        } else {
                                            if len(v46) >= 2 {
                                                v46 = v46[1:len(v46)]
                                            }
                                            v115 := GenS6{F0: 32767, F1: 42, F2: ""}
                                            var v116 int8 = 0
                                            _ = v115
                                            _ = v116
                                        }
                                        fmt.Println(fi78, v59, v4)
                                        _ = v113
                                    }
                                    var v117 int16 = 0
                                    v118 := GenS6{F0: (v117 + v117), F1: 0, F2: "foo bar"}
                                    fmt.Println(v47)
                                    _ = v117
                                    _ = v118
                                } else {
                                    v85 = v85
                                    v82 = v82
                                }
                                var v119 float32 = 1e10
                                _ = v119
                            default:
                                var v120 int = fi75
                                _ = v120
                            }
                            v121 := GenS5{F0: int64(v85), F1: 1, F2: 65535}
                            v4 = v4
                            _ = v121
                            _ = v89
                        } else {
                            fmt.Println(v47, v85, v82)
                            fmt.Println(v2)
                            delete(v3, "a")
                            var v122 uint32 = 42
                            fmt.Println(v47)
                            _ = v122
                        }
                        fmt.Println(v82)
                        _ = v88
                    }
                case 1:
                    var v123 bool = ((v39 - 0.1) != v2)
                    _ = v123
                case 0:
                    var v124 int8 = 0
                    switch "hello" {
                    case "a":
                        if ((v59 && (v82 >= v82)) && true) {
                            v125 := GenS6{F0: int16(v4), F1: 0, F2: "a"}
                            _ = v125
                        } else {
                            var v126 string = "foo bar"
                            _ = v126
                        }
                    case "hello":
                        fmt.Println(v59)
                    }
                    if (-159 > (fi78 - fi78)) {
                        fmt.Println(v124, fi78, v85)
                        v127 := GenS6{F0: int16(v85), F1: 4294967295, F2: "hello"}
                        _ = v127
                    } else {
                        switch fi75 {
                        case 9223372036854775807:
                            var v128 int8 = v124
                            _ = v128
                        case -9223372036854775808:
                            v1[("hello" + "")] = int(v82)
                        default:
                            v129 := v3["foo bar"]
                            _ = v129
                        }
                        fmt.Println(v59, fi75, v39)
                    }
                    _ = v124
                default:
                    fmt.Println(v4, v85)
                }
                switch (("hello" + "") + ("hello" + "\n\t")) {
                case "a":
                    v39 = v2
                case "\n\t":
                    if (!(1000 > 408)) {
                        _ = len(v46)
                        fmt.Println(v47)
                        fmt.Println(v4, fi75, v4)
                    } else {
                        v46 = append(v46, (fi75 | -9223372036854775808))
                        var v130 bool = true
                        v131 := GenS11{F0: 71}
                        _ = v130
                        _ = v131
                    }
                }
                _ = fi78
                _ = v82
                _ = v83
                _ = v84
                _ = v85
            }
        }
        v4 = v4
        _ = fi75
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v59
}
