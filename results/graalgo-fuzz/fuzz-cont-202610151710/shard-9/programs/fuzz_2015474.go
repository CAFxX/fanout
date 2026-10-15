package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 bool; F1 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 bool; F1 bool; F2 int8 }
type GenS12 struct { F0 uint32; F1 uint32; F2 bool }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint16; F1 int8 }
type GenS17 struct { F0 uint }
func genFn0(p0 int16, p1 uint64, p2 uint) uint8 {
    p1 = p1
    var v7 int16 = (-(p0))
    var v8 int64 = -9223372036854775808
    var v9 bool = (false && (!false))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1(p0 int8, p1 uint16, p2 uint16) int32 {
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: (((("foo bar" + "") + "\n\t") + (("\n\t" + "hello") + ("hello" + "foo bar"))) + "a"), F1: 0, F2: -9223372036854775808}
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 848
    var v3 int8 = 100
    var v4 int = -378
    var v5 float32 = 1.5
    var v6 int = (^(v4))
    v5 -= v5
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS2{F0: 75}
            _ = len(v1)
            if true {
                if true {
                    if (v3 != (^(v3))) {
                        var v16 int16 = 42
                        v4 = -748
                        _ = v16
                    } else {
                        v5 = v5
                        var v17 int64 = 277
                        _ = v17
                    }
                }
                v3 -= v3
                if false {
                    v18 := GenS2{F0: 184}
                    var v19 int32 = genFn1(v3, uint16(v2), 735)
                    _ = v18
                    _ = v19
                }
            }
            if (!true) {
                v20 := GenS2{F0: 42}
                var v21 float64 = (0.1 * 1e10)
                v21 -= (0.0 + v21)
                _ = v20
                _ = v21
            }
            switch "\n\t" {
            case "":
                var v22 float64 = (0.1 + -2.25)
                _ = v22
            default:
                v6 += 32768
            }
            v2 += v2
            v23 := []int{(fi14 * (fi14 << ((44) & 63)))}
            if false {
                var v24 float64 = float64(v5)
                v25 := GenS3{F0: v6}
                v26 := GenS1{F0: ("" + "hello"), F1: -286, F2: 1}
                fmt.Println(fi14)
                v6 = fi13
                if (!(42 == 2)) {
                    switch "foo bar" {
                    case "hello":
                        if false {
                            if false {
                                fmt.Println(v4, v24, fi14)
                            }
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                        }
                        fmt.Println(v4, v2)
                    case "":
                        v4 = (fi14 ^ (^(v6)))
                    default:
                        var v27 uint32 = (v2 - v2)
                        _ = v27
                    }
                } else {
                    fmt.Println(v6, v5, v24)
                    if ((fi13 ^ 127) < 32767) {
                        if (v3 != -35) {
                            v28 := v1["foo bar"]
                            var v29 uint8 = 26
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS4{F0: (v2 & v2), F1: (v3 % ((v3) | 1))}
                            v4 = v4
                            switch v6 {
                            case 128:
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            case 42:
                                v31 := GenS1{F0: "a", F1: -772, F2: -1}
                                _ = v31
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v23) > 0 {
                                        v23[0] = 2147483647
                                    }
                                case "hello":
                                    v5 = float32(v24)
                                    v3 = 42
                                    v4 = v4
                                    var v32 float64 = v24
                                    fmt.Println(fi14, v2)
                                    if ((((3.14159 <= (3.14159 - (float64(v5) * (float64(v5) / 0.1)))) && (!true)) || false) && (189 > genFn0(1, 763, 0))) {
                                        v2 += v2
                                        _ = len(v1)
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                        if false {
                                            v4 += v4
                                        } else {
                                            delete(v1, "a")
                                        }
                                    }
                                    var v33 uint32 = v2
                                    if true {
                                        var v34 bool = true
                                        var v35 float32 = 1.5
                                        var v36 uint = 255
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v32
                                    _ = v33
                                case "foo bar":
                                    v1[("a" + ("a" + "\n\t"))] = ((0 < v3) && true)
                                default:
                                    fmt.Println(v2, v24, v5)
                                }
                            }
                            _ = v30
                        }
                        switch 0 {
                        case -450:
                            if (1 != 773) {
                                v4 -= 32767
                                switch ("a" + (("hello" + ("foo bar" + "hello")) + (("" + ((("\n\t" + "") + "\n\t") + ("\n\t" + (("\n\t" + ("foo bar" + ("\n\t" + ("" + "foo bar")))) + "hello")))) + ""))) {
                                case "a":
                                    v37 := GenS5{F0: 255, F1: 1, F2: v2}
                                    _ = v37
                                case "\n\t":
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    v38 := GenS6{F0: (v6 & 0), F1: "\n\t"}
                                    if false {
                                        v5 = v5
                                    }
                                    var v39 int64 = -622
                                    _ = v38
                                    _ = v39
                                }
                                var v40 int64 = 825
                                _ = v40
                            } else {
                                if true {
                                    v41 := GenS7{F0: true}
                                    v5 += v5
                                    v42 := GenS4{F0: (v2 & v2), F1: (v3 << ((0) & 63))}
                                    v1[""] = false
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 uint8 = 0
                                    fmt.Println(v4, fi14, v24)
                                    _ = v43
                                }
                                var v44 bool = (-2147483648 != -1)
                                _ = v44
                            }
                        case 32768:
                            fmt.Println(v6, v24, v6)
                        default:
                            var v45 uint16 = 167
                            _ = v45
                        }
                    } else {
                        fmt.Println(fi14, v4)
                        if (false || true) {
                            var v46 uint64 = 0
                            _ = v46
                        }
                        switch 1 {
                        case -9223372036854775808:
                            v47 := GenS8{F0: 42, F1: (fi13 - 32768)}
                            _ = v47
                        default:
                            v2 -= v2
                            fmt.Println(v24, fi14, v2)
                            var v48 float64 = ((-2.25 / (3.14159 * (0.1 * v24))) * v24)
                            _ = v48
                        }
                    }
                    switch (v6 / ((fi14) | 1)) {
                    case -2147483648:
                        var v49 uint64 = 42
                        _ = v49
                    case 204:
                        v50 := GenS1{F0: "", F1: 1000, F2: -960}
                        _ = v50
                    }
                }
                v51 := GenS9{F0: (2 == 1), F1: v2}
                var v52 int = 32768
                var v53 string = ""
                _ = v24
                _ = v25
                _ = v26
                _ = v51
                _ = v52
                _ = v53
            } else {
                v5 -= 0.0
                var v54 uint32 = (v2 ^ (v2 - v2))
                if (42 >= 0) {
                    _ = len(v1)
                    switch -557 {
                    case 566:
                        fmt.Println(v2, v4)
                    case -957:
                        fmt.Println(v4, v5, v3)
                    default:
                        if len(v23) > 0 {
                            v23[0] = v4
                        }
                    }
                }
                _ = v54
            }
            _ = fi14
            _ = v15
            _ = v23
        }
        _ = fi13
    }
    v55 := GenS10{F0: -1}
    if ((v3 % ((1) | 1)) < -1) {
        v56 := GenS11{F0: false, F1: false, F2: v3}
        var v57 int8 = v3
        var v58 string = ""
        _ = v56
        _ = v57
        _ = v58
    }
    var v59 float64 = (1.5 - 0.0)
    var v60 float64 = (3.14159 * v59)
    v61 := v1["a"]
    if v61 {
        switch "foo bar" {
        case "":
            if (v3 < (v3 * v3)) {
                for fi62 := 0; fi62 < 13; fi62++ {
                    switch (fi62 >> ((407) & 63)) {
                    case 175:
                        fmt.Println(v6, v59, fi62)
                    case 195:
                        var v63 uint32 = (v2 % ((1) | 1))
                        _ = v63
                    case 2147483647:
                        v64 := GenS9{F0: v61, F1: v2}
                        _ = v64
                    default:
                        v65 := GenS2{F0: 1}
                        _ = v65
                    }
                    _ = fi62
                }
                var v66 uint64 = 1
                if v61 {
                    switch v6 {
                    case 2147483647:
                        if ((v3 << ((2) & 63)) > (v3 % ((127) | 1))) {
                            for fi67 := 0; fi67 < 19; fi67++ {
                                v68 := v1["\n\t"]
                                v2 -= v2
                                _ = fi67
                                _ = v68
                            }
                        }
                    case 4294967295:
                        for fi69 := 0; fi69 < 18; fi69++ {
                            var v70 int16 = 1000
                            for fi71 := 0; fi71 < 12; fi71++ {
                                switch -1 {
                                case 2147483647:
                                    var v72 uint32 = (v2 << ((42) & 63))
                                    _ = v72
                                case -1:
                                    if (v61 || (v61 && (33 < int32(v2)))) {
                                        v59 = 0.1
                                        if false {
                                            switch "\n\t" {
                                            case "":
                                                var v73 int32 = 1
                                                _ = v73
                                            case "\n\t":
                                                var v74 int16 = (-(v70))
                                                _ = v74
                                            default:
                                                var v75 uint = 1
                                                _ = v75
                                            }
                                            delete(v1, "\n\t")
                                        } else {
                                            switch fi69 {
                                            case 359:
                                                v76 := []int{128}
                                                _ = v76
                                            case -301:
                                                var v77 bool = v61
                                                _ = v77
                                            case 9223372036854775807:
                                                var v78 bool = false
                                                _ = v78
                                            }
                                            var v79 int16 = v70
                                            switch (("hello" + "\n\t") + "hello") {
                                            case "hello":
                                                fmt.Println(v6)
                                            case "\n\t":
                                                if (!(v2 == v2)) {
                                                    v5 -= (-2.25 - 0.1)
                                                    var v80 int32 = -1
                                                    _ = v80
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if true {
                                                            _ = len(v1)
                                                            fmt.Println(v61, fi69, v5)
                                                        }
                                                    case "foo bar":
                                                        switch (("hello" + "hello") + "") {
                                                        case "":
                                                            var v81 uint64 = 2
                                                            _ = v81
                                                        case "hello":
                                                            var v82 uint = 18446744073709551615
                                                            _ = v82
                                                        }
                                                    case "":
                                                        if true {
                                                            switch (v6 ^ (fi71 ^ v6)) {
                                                            case 1:
                                                                if ("hello" < "\n\t") {
                                                                    var v83 int32 = 42
                                                                    _ = v83
                                                                }
                                                            case 4294967295:
                                                                v70 = v70
                                                            default:
                                                                v66 = v66
                                                            }
                                                        }
                                                    default:
                                                        v79 = (v70 ^ (v79 << ((1) & 63)))
                                                    }
                                                    switch v4 {
                                                    case 2147483647:
                                                        if (v61 || ("" > ("foo bar" + "a"))) {
                                                            v4 = (^(fi71))
                                                            fmt.Println(v70, v70)
                                                        }
                                                    case 32767:
                                                        switch (fi69 + v6) {
                                                        case 128:
                                                            var v84 uint8 = 128
                                                            _ = v84
                                                        case 42:
                                                            switch ("foo bar" + ((("\n\t" + "") + "foo bar") + "a")) {
                                                            case "a":
                                                                v85 := []int{(fi71 + (v6 | 32768)), (v6 % (((fi69 << ((18446744073709551615) & 63))) | 1))}
                                                                _ = v85
                                                            default:
                                                                var v86 int64 = 741
                                                                _ = v86
                                                            }
                                                        }
                                                    case 871:
                                                        var v87 uint32 = v2
                                                        _ = v87
                                                    }
                                                }
                                            default:
                                                switch (v6 & (v6 ^ v4)) {
                                                case 1:
                                                    var v88 uint32 = (v2 % ((0) | 1))
                                                    _ = v88
                                                case -1:
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        v89 := GenS12{F0: v2, F1: (v2 | v2), F2: false}
                                                        _ = v89
                                                    case "foo bar":
                                                        var v90 int16 = (v70 | 1)
                                                        _ = v90
                                                    case "hello":
                                                        var v91 uint = 0
                                                        _ = v91
                                                    default:
                                                        var v92 uint = 745
                                                        _ = v92
                                                    }
                                                    if false {
                                                        v93 := []int{32767, int(v66), 243}
                                                        _ = v93
                                                    } else {
                                                        v94 := GenS13{F0: (((true || true) && (-415 < 0)) && v61)}
                                                        var v95 uint16 = 632
                                                        var v96 uint32 = v2
                                                        v97 := GenS14{F0: -1}
                                                        _ = len(v1)
                                                        _ = v94
                                                        _ = v95
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    var v98 string = "foo bar"
                                                    _ = v98
                                                case 42:
                                                    v4 = -641
                                                    v3 = 1
                                                    var v99 uint64 = 199
                                                    _ = v99
                                                }
                                            }
                                            _ = v79
                                        }
                                    }
                                case 42:
                                    var v100 float64 = (float64(0.1) + v60)
                                    _ = v100
                                default:
                                    if (65535 > 0) {
                                        var v101 float32 = v5
                                        _ = v101
                                    }
                                }
                                if (true || (!true)) {
                                    var v102 uint64 = (v66 * 4294967295)
                                    _ = v102
                                } else {
                                    if (!v61) {
                                        var v103 int32 = -80
                                        _ = v103
                                    } else {
                                        var v104 int16 = v70
                                        v105 := GenS5{F0: (v104 | v70), F1: 2, F2: 1}
                                        _ = v104
                                        _ = v105
                                    }
                                    v106 := []int{fi69, fi69, (v6 + fi71), (v4 & (v4 * (fi69 * (v6 | v6))))}
                                    v107 := GenS4{F0: v2, F1: v3}
                                    fmt.Println(fi71, fi69, v59)
                                    var v108 int32 = 1
                                    _ = v106
                                    _ = v107
                                    _ = v108
                                }
                                _ = fi71
                            }
                            _ = fi69
                            _ = v70
                        }
                    default:
                        fmt.Println(v66, v61, v61)
                    }
                    switch (v4 ^ v4) {
                    case 32767:
                        fmt.Println(v2, v6)
                    case 0:
                        if (uint(v66) != 873) {
                            var v109 uint64 = 601
                            if (32 < 127) {
                                for fi110 := 0; fi110 < 4; fi110++ {
                                    var v111 float64 = (v59 / 0.1)
                                    _ = fi110
                                    _ = v111
                                }
                                if (255 < 97) {
                                    v59 = v60
                                }
                                switch v6 {
                                case -9223372036854775808:
                                    var v112 int8 = -128
                                    _ = v112
                                case 0:
                                    v113 := GenS9{F0: v61, F1: v2}
                                    _ = v113
                                }
                            }
                            var v114 int64 = -1
                            _ = v109
                            _ = v114
                        } else {
                            var v115 int16 = 42
                            _ = v115
                        }
                    }
                    var v116 uint16 = 42
                    _ = v116
                } else {
                    v117 := GenS5{F0: -339, F1: 672, F2: (v2 << ((18446744073709551615) & 63))}
                    _ = v117
                }
                _ = v66
            }
        }
        v60 = v60
    } else {
        if v61 {
            var v118 uint64 = 391
            _ = v118
        } else {
            v119 := []int{v4, v4}
            v120 := GenS15{F0: "hello"}
            switch (v4 | -1) {
            case 128:
                switch ("" + "a") {
                case "\n\t":
                    if false {
                        var v121 uint16 = 1
                        v6 -= (-(v4))
                        switch "" {
                        case "foo bar":
                            _ = len(v1)
                        case "\n\t":
                            v3 += (v3 >> ((453) & 63))
                        }
                        _ = v121
                    } else {
                        if v61 {
                            for fi122 := 0; fi122 < 15; fi122++ {
                                v59 -= v60
                                v123 := GenS16{F0: 42, F1: -122}
                                _ = fi122
                                _ = v123
                            }
                        }
                        var v124 int32 = -1
                        v3 -= int8(v124)
                        _ = v124
                    }
                case "a":
                    switch "foo bar" {
                    case "a":
                        v125 := GenS13{F0: (!(-432 < 1))}
                        _ = v125
                    }
                }
            case 32767:
                v3 = (v3 * v3)
            default:
                fmt.Println(v5, v59, v59)
            }
            _ = v119
            _ = v120
        }
        if (true && true) {
            if v61 {
                var v126 int16 = 853
                v127 := []int{v4, 2147483647, -1}
                _ = v126
                _ = v127
            } else {
                v2 += (^(v2))
                if ((uint64(v2) >= 42) && v61) {
                    fmt.Println(v2, v60, v4)
                    v128 := GenS10{F0: 2147483647}
                    _ = v128
                }
                if ((!(1 < 52)) && (0.0 == v59)) {
                    v6 += 127
                    v129 := GenS2{F0: 1}
                    _ = v129
                }
            }
        } else {
            for fi130 := 0; fi130 < 3; fi130++ {
                var v131 int = v4
                var v132 float32 = v5
                _ = fi130
                _ = v131
                _ = v132
            }
            v6 = (v6 * v6)
            var v133 string = ("hello" + "")
            _ = v133
        }
    }
    for fi134 := 0; fi134 < 11; fi134++ {
        v60 += 1.5
        switch v6 {
        case 1:
            for fi135 := 0; fi135 < 3; fi135++ {
                if (241 > 1) {
                    switch v4 {
                    case 530:
                        v60 += v60
                    case -262:
                        var v136 float32 = float32(float64((v5 - v5)))
                        _ = v136
                    case -2147483648:
                        var v137 int64 = -9223372036854775808
                        _ = v137
                    }
                } else {
                    v138 := GenS2{F0: 255}
                    var v139 uint32 = v2
                    _ = v138
                    _ = v139
                }
                v140 := []int{fi134, (fi134 * fi135), (v6 ^ fi134)}
                v141 := GenS17{F0: 175}
                _ = fi135
                _ = v140
                _ = v141
            }
        case -9223372036854775808:
            v6 -= (v4 | v6)
        }
        _ = fi134
    }
    v142 := []int{}
    var v143 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v142:", len(v142), v142)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v142
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
