package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS8 struct { F0 int32 }
func genFn0() int32 {
    var v7 uint8 = 2
    var v8 uint64 = 0
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{fi9, (fi9 << ((487) & 63)), (-(fi9)), fi9}
        var v11 string = "hello"
        if false {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(v7, v7, v8)
            v12 := GenS1{F0: 759, F1: (^(v7)), F2: 0}
            _ = v12
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v13 int64 = -681
    v8 = (v8 & v8)
    _ = v13
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int64, p1 int64, p2 uint) int32 {
    v14 := GenS1{F0: 1, F1: 2, F2: 100}
    v15 := GenS1{F0: 615, F1: 128, F2: 42}
    switch 0 {
    case 9223372036854775807:
        v16 := []int{-180, 32768}
        _ = v16
    case 32768:
        p1 -= p0
    case -9223372036854775808:
        p1 = p1
    default:
        var v17 bool = true
        _ = v17
    }
    v18 := []int{}
    _ = len(v18)
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v18
    return genFn0()
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.0
    var v3 int64 = 9223372036854775807
    v4 := []int{32768, 42, 906}
    var v5 uint64 = 1
    v6 := []string{"hello", "", ("a" + "")}
    if (v2 != v2) {
        v19 := GenS1{F0: 1, F1: 0, F2: -126}
        switch 32767 {
        case 0:
            _ = len(v4)
        case 1:
            if false {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v5 = v5
                    _ = fi20
                }
                v3 = (v3 * (v3 >> ((0) & 63)))
                fmt.Println(v3, v2, v5)
            }
        case -1:
            v6 = append(v6, "a")
        }
        _ = v19
    } else {
        if ("a" == "a") {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int16 = 32767
                _ = fi21
                _ = v22
            }
        }
        if false {
            var v23 uint = 183
            switch "a" {
            case "hello":
                fmt.Println(v2)
            }
            _ = v23
        }
        switch "\n\t" {
        case "\n\t":
            if (!true) {
                for fi24 := 0; fi24 < 18; fi24++ {
                    v25 := GenS1{F0: 32767, F1: 127, F2: 88}
                    if ((322 == 42) && (0 == 127)) {
                        v4 = append(v4, (fi24 % ((fi24) | 1)))
                        var v26 uint = 971
                        v27 := v1[-1]
                        _ = v26
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 19; fi28++ {
                            fmt.Println(fi24, v5)
                            switch "" {
                            case "foo bar":
                                v5 = 1
                            case "hello":
                                v29 := GenS2{F0: "a"}
                                _ = v29
                            default:
                                v3 += v3
                                v30 := GenS3{F0: 255, F1: 42}
                                var v31 float32 = 0.1
                                _ = v30
                                _ = v31
                            }
                            _ = fi28
                        }
                        v5 = 0
                        var v32 uint16 = 42
                        _ = v32
                    }
                    _ = fi24
                    _ = v25
                }
            } else {
                var v33 float32 = 0.0
                _ = v33
            }
        case "a":
            var v34 int = 77
            _ = v34
        case "":
            for fi35 := 0; fi35 < 10; fi35++ {
                fmt.Println(v2, v3)
                fmt.Println(fi35)
                _ = fi35
            }
        default:
            var v36 int64 = (v3 % (((v3 % ((734) | 1))) | 1))
            _ = v36
        }
    }
    v37 := GenS1{F0: 1, F1: 127, F2: 100}
    var v38 float64 = 3.14159
    for fi39 := 0; fi39 < 13; fi39++ {
        switch (fi39 * (^(fi39))) {
        case 32767:
            for fi40 := 0; fi40 < 16; fi40++ {
                v3 = -9223372036854775808
                switch fi39 {
                case 128:
                    if (false || true) {
                        var v41 float64 = float64(0.0)
                        _ = len(v4)
                        fmt.Println(v41, v38)
                        if (32767 >= 2) {
                            v42 := GenS2{F0: "\n\t"}
                            _ = v42
                        }
                        v38 = float64((1.5 / -2.25))
                        _ = v41
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                case 127:
                    switch fi40 {
                    case 128:
                        switch ("hello" + ("hello" + "foo bar")) {
                        case "":
                            if ((false || false) && (false && false)) {
                                v38 += v38
                                if false {
                                    v43 := v1[(-(fi40))]
                                    var v44 uint8 = 255
                                    _ = v43
                                    _ = v44
                                } else {
                                    if false {
                                        v45 := GenS4{F0: 18446744073709551615}
                                        if len(v6) > 0 {
                                            v6[0] = "a"
                                        }
                                        _ = v45
                                    } else {
                                        switch (-(fi39)) {
                                        case 128:
                                            v46 := GenS3{F0: 1, F1: 142}
                                            _ = v46
                                        default:
                                            fmt.Println(v38, v5)
                                        }
                                        var v47 bool = true
                                        var v48 uint32 = 42
                                        _ = v47
                                        _ = v48
                                    }
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(fi39, v3)
                                    case "\n\t":
                                        var v49 int16 = 1
                                        _ = v49
                                    case "":
                                        _ = len(v6)
                                    default:
                                        var v50 int64 = (v3 ^ v3)
                                        _ = v50
                                    }
                                    var v51 int16 = 790
                                    _ = v51
                                }
                            } else {
                                v3 = v3
                                fmt.Println(v2, fi39, v2)
                                if true {
                                    v52 := v1[(-(fi40))]
                                    fmt.Println(v38, v3, v52)
                                    _ = v52
                                } else {
                                    v53 := GenS3{F0: 0, F1: 127}
                                    var v54 uint8 = 57
                                    switch (^(fi39)) {
                                    case -963:
                                        var v55 uint16 = 1
                                        _ = v55
                                    }
                                    v1[(fi39 + (fi40 % ((42) | 1)))] = (fi40 + fi39)
                                    var v56 string = ""
                                    _ = v53
                                    _ = v54
                                    _ = v56
                                }
                            }
                        case "foo bar":
                            var v57 string = "foo bar"
                            _ = v57
                        case "\n\t":
                            var v58 uint = 449
                            _ = v58
                        default:
                            var v59 uint8 = 255
                            _ = v59
                        }
                    case 0:
                        _ = len(v1)
                    case -333:
                        var v60 bool = (true || (!false))
                        _ = v60
                    }
                default:
                    var v61 uint = 4294967295
                    _ = v61
                }
                fmt.Println(v5, v2)
                _ = fi40
            }
        case -84:
            fmt.Println(v2, v2, v2)
        case -319:
            var v62 string = "foo bar"
            _ = v62
        default:
            var v63 int16 = -716
            _ = v63
        }
        var v64 float32 = ((((1.5 + -2.25) / 1.0) / 1.0) + 0.1)
        _ = fi39
        _ = v64
    }
    switch -2147483648 {
    case 32767:
        switch "hello" {
        case "a":
            var v65 float32 = (0.1 - 1e10)
            _ = v65
        }
    }
    switch "hello" {
    case "\n\t":
        if false {
            v66 := GenS1{F0: 0, F1: 1, F2: -128}
            _ = v66
        }
    default:
        switch 32767 {
        case 0:
            if true {
                fmt.Println(v2, v38)
            } else {
                for fi67 := 0; fi67 < 3; fi67++ {
                    var v68 int8 = 0
                    _ = fi67
                    _ = v68
                }
            }
        case 799:
            v4 = append(v4, -945)
        }
    }
    var v69 int32 = 42
    if false {
        var v70 uint32 = 2147483647
        _ = v70
    }
    switch -1 {
    case 307:
        v1[9223372036854775807] = 1
    case 0:
        for fi71 := 0; fi71 < 1; fi71++ {
            fmt.Println(v3, fi71, v69)
            switch (("hello" + ("a" + ("" + "foo bar"))) + "\n\t") {
            case "a":
                v72 := GenS1{F0: 255, F1: 255, F2: -41}
                _ = v72
            default:
                var v73 float32 = 3.14159
                _ = v73
            }
            v74 := GenS4{F0: 18446744073709551615}
            _ = fi71
            _ = v74
        }
    case 489:
        if ((false && true) || false) {
            switch 3 {
            case 32768:
                v2 += v38
            default:
                v1[-1] = -1
            }
        } else {
            switch "a" {
            case "hello":
                for fi75 := 0; fi75 < 6; fi75++ {
                    for fi76 := 0; fi76 < 20; fi76++ {
                        if false {
                            switch "\n\t" {
                            case "hello":
                                if len(v6) > 0 {
                                    v6[0] = ("" + (("\n\t" + "a") + "foo bar"))
                                }
                            default:
                                v77 := GenS2{F0: "hello"}
                                _ = v77
                            }
                            var v78 string = "foo bar"
                            v79 := GenS2{F0: ""}
                            v38 = v2
                            fmt.Println(v3)
                            _ = v78
                            _ = v79
                        } else {
                            if true {
                                var v80 uint32 = 65535
                                var v81 bool = ((v69 / ((v69) | 1)) <= 448)
                                _ = len(v6)
                                _ = v80
                                _ = v81
                            }
                            v82 := GenS5{F0: (v3 & v3)}
                            fmt.Println(v2, v3)
                            _ = v82
                        }
                        v83 := GenS2{F0: (("a" + "foo bar") + ("hello" + (("" + "\n\t") + ("hello" + "hello"))))}
                        _ = fi76
                        _ = v83
                    }
                    if true {
                        v2 -= 3.14159
                        var v84 uint32 = 83
                        _ = v84
                    } else {
                        var v85 uint16 = 32767
                        var v86 int32 = (v69 % (((v69 & v69)) | 1))
                        fmt.Println(v3, v86, v38)
                        _ = v85
                        _ = v86
                    }
                    _ = fi75
                }
            case "a":
                switch "" {
                case "\n\t":
                    if (3.14159 <= 3.14159) {
                        var v87 int16 = 1
                        switch ("a" + ("foo bar" + ("hello" + "hello"))) {
                        case "":
                            var v88 int16 = 1
                            _ = v88
                        case "a":
                            var v89 float64 = (v2 * (3.14159 / 0.1))
                            _ = v89
                        case "foo bar":
                            v87 += (v87 + (^(v87)))
                        }
                        _ = v87
                    } else {
                        v90 := GenS2{F0: "a"}
                        for fi91 := 0; fi91 < 3; fi91++ {
                            v5 -= 2
                            if (true && ((42 == -743) || true)) {
                                var v92 string = "\n\t"
                                _ = v92
                            } else {
                                if false {
                                    var v93 uint8 = 196
                                    _ = v93
                                }
                            }
                            _ = fi91
                        }
                        _ = v90
                    }
                case "hello":
                    if (!false) {
                        var v94 uint8 = 42
                        v69 = -2147483648
                        for fi95 := 0; fi95 < 3; fi95++ {
                            switch fi95 {
                            case 9223372036854775807:
                                v2 -= v2
                            }
                            v5 += v5
                            v94 -= (v94 << ((0) & 63))
                            _ = fi95
                        }
                        _ = v94
                    } else {
                        switch "hello" {
                        case "":
                            if true {
                                var v96 int64 = (v3 ^ (v3 + v3))
                                _ = v96
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = "hello"
                                }
                                var v97 bool = true
                                v98 := GenS4{F0: 153}
                                _ = v97
                                _ = v98
                            }
                        default:
                            var v99 uint32 = 2
                            _ = v99
                        }
                        if ((v2 - float64(0.1)) == v38) {
                            for fi100 := 0; fi100 < 18; fi100++ {
                                v101 := GenS6{F0: false, F1: (^(v5))}
                                v69 -= v69
                                switch "foo bar" {
                                case "a":
                                    fmt.Println(v2, v5, v3)
                                case "":
                                    for fi102 := 0; fi102 < 14; fi102++ {
                                        switch "\n\t" {
                                        case "\n\t":
                                            switch (fi100 | int(v5)) {
                                            case 128:
                                                var v103 uint64 = v5
                                                _ = v103
                                            case 4294967295:
                                                switch 32768 {
                                                case 4294967295:
                                                    v104 := GenS7{F0: 905, F1: v69, F2: v5}
                                                    _ = v104
                                                case 9223372036854775807:
                                                    v3 += (^(v3))
                                                case -18:
                                                    fmt.Println(fi100, v38)
                                                }
                                            }
                                        case "hello":
                                            if true {
                                                var v105 float32 = (1.5 + (1e10 - (0.1 - (3.14159 / (1e10 * 3.14159)))))
                                                if len(v4) > 0 {
                                                    v4[0] = (fi102 - (fi100 >> ((1) & 63)))
                                                }
                                                var v106 uint16 = 0
                                                _ = v105
                                                _ = v106
                                            } else {
                                                v107 := GenS2{F0: ""}
                                                var v108 int64 = 9223372036854775807
                                                v69 = -586
                                                _ = v107
                                                _ = v108
                                            }
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = (("hello" + ("" + ("a" + "\n\t"))) + (("hello" + "a") + (("\n\t" + ("" + "\n\t")) + "\n\t")))
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = "\n\t"
                                        }
                                        _ = fi102
                                    }
                                default:
                                    if true {
                                        for fi109 := 0; fi109 < 13; fi109++ {
                                            v110 := GenS5{F0: v3}
                                            var v111 int = (fi100 * (fi100 - (fi100 * -2147483648)))
                                            _ = fi109
                                            _ = v110
                                            _ = v111
                                        }
                                        switch fi100 {
                                        case -229:
                                            var v112 float64 = v2
                                            _ = v112
                                        case 42:
                                            v4 = append(v4, int(v3))
                                        case 4294967295:
                                            for fi113 := 0; fi113 < 14; fi113++ {
                                                v5 = (v5 / ((v5) | 1))
                                                _ = fi113
                                            }
                                        }
                                    } else {
                                        var v114 uint16 = 32767
                                        _ = v114
                                    }
                                }
                                _ = fi100
                                _ = v101
                            }
                            if (!(!true)) {
                                for fi115 := 0; fi115 < 6; fi115++ {
                                    v116 := GenS2{F0: "hello"}
                                    var v117 int64 = 1
                                    var v118 uint = 18446744073709551615
                                    _ = fi115
                                    _ = v116
                                    _ = v117
                                    _ = v118
                                }
                            }
                            v3 = v3
                        } else {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v3, v3)
                            case "\n\t":
                                var v119 uint64 = v5
                                _ = v119
                            default:
                                var v120 uint16 = 1
                                _ = v120
                            }
                        }
                        v121 := v1[-64]
                        _ = v121
                    }
                }
            }
            for fi122 := 0; fi122 < 9; fi122++ {
                var v123 uint = 18446744073709551615
                for fi124 := 0; fi124 < 1; fi124++ {
                    var v125 int = 127
                    v126 := GenS8{F0: (-(v69))}
                    var v127 float32 = -2.25
                    _ = fi124
                    _ = v125
                    _ = v126
                    _ = v127
                }
                _ = fi122
                _ = v123
            }
            var v128 float64 = v2
            _ = v128
        }
    }
    for fi129 := 0; fi129 < 9; fi129++ {
        var v130 uint = 1
        _ = len(v4)
        _ = fi129
        _ = v130
    }
    var v131 uint32 = 0
    fmt.Println(v5, v69, v69)
    v2 = (v38 - 0.1)
    v132 := GenS6{F0: (93 >= 251), F1: (v5 + (v5 / ((v5) | 1)))}
    v133 := GenS4{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v131
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
