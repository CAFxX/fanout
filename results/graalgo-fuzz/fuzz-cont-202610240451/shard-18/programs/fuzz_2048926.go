package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int32; F2 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 int16 }
type GenS7 struct { F0 int; F1 string; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 bool; F2 int32 }
type GenS10 struct { F0 uint8; F1 int64; F2 bool }
type GenS11 struct { F0 string; F1 int64; F2 bool }
type GenS12 struct { F0 int32; F1 string; F2 uint32 }
type GenS13 struct { F0 int; F1 uint8; F2 uint8 }
type GenS14 struct { F0 string; F1 int32 }
type GenS15 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS16 struct { F0 uint; F1 string }
type GenS17 struct { F0 uint64; F1 uint64 }
type GenS18 struct { F0 bool; F1 int; F2 uint16 }
type GenS19 struct { F0 int; F1 uint64 }
func genFn0(p0 uint64, p1 int) int {
    v4 := GenS1{F0: 1, F1: -1}
    fmt.Println(p0)
    v5 := []int{p1, p1, -841, 806}
    v6 := GenS1{F0: -1, F1: -128}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return p1
}
func genFn1(p0 int8, p1 int64) int8 {
    var v7 string = (("" + ("hello" + "\n\t")) + "a")
    v8 := []int{0, 32767}
    var v9 uint = 1
    var v10 int16 = 1000
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (p0 * 4)
}

func main() {
    var v1 uint64 = 974
    var v2 uint16 = 42
    v3 := make(map[int]string)
    var v11 uint64 = v1
    v1 = (v11 - v11)
    var v12 int64 = -9223372036854775808
    for fi13 := 0; fi13 < 7; fi13++ {
        if false {
            v14 := []int{fi13, -9223372036854775808, 0}
            v1 = (v11 << ((1) & 63))
            var v15 int32 = -2147483648
            _ = v14
            _ = v15
        } else {
            var v16 int16 = 1000
            v17 := GenS1{F0: v12, F1: -1}
            var v18 uint = 18446744073709551615
            _ = v16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            if ((v1 & (^(v11))) > 235) {
                var v20 float32 = 1.5
                _ = v20
            } else {
                switch (fi19 % (((^(fi19))) | 1)) {
                case 53:
                    v11 += 0
                case 128:
                    v21 := []int{genFn0(v1, (^(fi13))), -371, 128}
                    _ = v21
                default:
                    var v22 int64 = v12
                    if (!(275 <= v11)) {
                        v3[fi19] = "a"
                        if false {
                            var v23 int32 = -2147483648
                            _ = v23
                        } else {
                            v24 := GenS1{F0: (v22 % (((^(v12))) | 1)), F1: 110}
                            if true {
                                _ = len(v3)
                                v25 := GenS1{F0: 597, F1: 42}
                                _ = v25
                            } else {
                                var v26 int8 = 1
                                _ = v26
                            }
                            switch fi13 {
                            case 32768:
                                if ((v12 + (v12 - (v22 - (v12 >> ((42) & 63))))) <= 947) {
                                    if false {
                                        v2 += (v2 | (v2 * 42))
                                        switch "a" {
                                        case "a":
                                            fmt.Println(v1, v12, v22)
                                        case "hello":
                                            var v27 int16 = -32768
                                            _ = v27
                                        case "":
                                            var v28 float32 = 0.1
                                            _ = v28
                                        }
                                        switch "" {
                                        case "foo bar":
                                            v12 = (^(v12))
                                        case "":
                                            if ((v22 + (^(v22))) > (^(v22))) {
                                                v1 = (v11 + 1)
                                                var v29 uint64 = v1
                                                switch (("" + "\n\t") + "a") {
                                                case "hello":
                                                    v30 := GenS2{F0: (v1 >> ((65535) & 63))}
                                                    _ = v30
                                                case "a":
                                                    if false {
                                                        if false {
                                                            v3[fi19] = "\n\t"
                                                        } else {
                                                            var v31 float32 = 1.5
                                                            _ = v31
                                                        }
                                                        if (true && (!true)) {
                                                            var v32 int64 = v22
                                                            v2 = 32767
                                                            v33 := GenS3{F0: -24}
                                                            _ = v32
                                                            _ = v33
                                                        } else {
                                                            v34 := []int{}
                                                            if (2 >= (v2 % (((v2 - v2)) | 1))) {
                                                                if true {
                                                                    var v35 string = "\n\t"
                                                                    v36 := GenS4{F0: 445, F1: 0, F2: false}
                                                                    v37 := GenS5{F0: 18446744073709551615}
                                                                    _ = v35
                                                                    _ = v36
                                                                    _ = v37
                                                                } else {
                                                                    if (-2.25 != 0.0) {
                                                                        v2 += (v2 ^ v2)
                                                                        v38 := GenS4{F0: 723, F1: 205, F2: ((!true) || false)}
                                                                        switch (("\n\t" + "\n\t") + "hello") {
                                                                        case "":
                                                                            switch fi13 {
                                                                            case 526:
                                                                                if ((true && true) && (42 <= 0)) {
                                                                                    var v39 int8 = 100
                                                                                    switch ((("foo bar" + "a") + "hello") + "foo bar") {
                                                                                    case "a":
                                                                                        var v40 uint64 = v1
                                                                                        _ = v40
                                                                                    }
                                                                                    v12 -= 244
                                                                                    _ = v39
                                                                                } else {
                                                                                    var v41 int32 = 957
                                                                                    var v42 float32 = 1.5
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                        case "\n\t":
                                                                            fmt.Println(v12, fi19, v11)
                                                                        case "foo bar":
                                                                            fmt.Println(v2, v12)
                                                                        default:
                                                                            var v43 uint = 36
                                                                            _ = v43
                                                                        }
                                                                        _ = v38
                                                                    }
                                                                    var v44 int16 = 42
                                                                    if true {
                                                                        var v45 float32 = 1.5
                                                                        _ = v45
                                                                    }
                                                                    fmt.Println(v2)
                                                                    _ = v44
                                                                }
                                                                if true {
                                                                    v29 = (v29 / (((^(v1))) | 1))
                                                                    var v46 int8 = 3
                                                                    var v47 int = fi19
                                                                    _ = v46
                                                                    _ = v47
                                                                } else {
                                                                    _ = len(v34)
                                                                }
                                                            } else {
                                                                switch "a" {
                                                                case "hello":
                                                                    v11 = v11
                                                                case "":
                                                                    if true {
                                                                        switch "" {
                                                                        case "a":
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                v48 := GenS6{F0: 4294967295, F1: -1}
                                                                                _ = v48
                                                                            case "\n\t":
                                                                                var v49 string = "\n\t"
                                                                                _ = v49
                                                                            case "a":
                                                                                if (!false) {
                                                                                    if ((!(true || true)) || false) {
                                                                                        fmt.Println(v1, v11, v2)
                                                                                        if true {
                                                                                            v50 := GenS1{F0: v12, F1: 71}
                                                                                            v22 = v12
                                                                                            _ = v50
                                                                                        }
                                                                                    } else {
                                                                                        v2 += 0
                                                                                        v1 -= (v1 ^ (v1 & (v1 / ((0) | 1))))
                                                                                        v51 := GenS5{F0: 1}
                                                                                        if ((v22 * (^(v12))) < (v12 << ((1) & 63))) {
                                                                                            var v52 int8 = 0
                                                                                            switch fi13 {
                                                                                            case -1:
                                                                                                var v53 int32 = -2147483648
                                                                                                _ = v53
                                                                                            case 32768:
                                                                                                v11 = 4294967295
                                                                                            case -494:
                                                                                                v54 := GenS7{F0: fi19, F1: ("a" + "a"), F2: v12}
                                                                                                var v55 int32 = 0
                                                                                                _ = v54
                                                                                                _ = v55
                                                                                            default:
                                                                                                var v56 int = genFn0(2, int(v22))
                                                                                                _ = v56
                                                                                            }
                                                                                            _ = v52
                                                                                        }
                                                                                        v11 = (v29 - 18446744073709551615)
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            v57 := GenS2{F0: v29}
                                                                                            _ = v57
                                                                                        case "":
                                                                                            v29 -= 0
                                                                                        case "hello":
                                                                                            v29 = (v11 | v11)
                                                                                        default:
                                                                                            var v58 int16 = -84
                                                                                            _ = v58
                                                                                        }
                                                                                        var v59 uint8 = 127
                                                                                        _ = v51
                                                                                        _ = v59
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                var v60 uint32 = 0
                                                                                var v61 bool = false
                                                                                if (true && false) {
                                                                                    v1 += 2
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v62 int16 = 32767
                                                                                        _ = v62
                                                                                    case "hello":
                                                                                        var v63 int16 = 0
                                                                                        _ = v63
                                                                                    case "\n\t":
                                                                                        v64 := GenS1{F0: (v22 + (v22 - (v22 * v22))), F1: -25}
                                                                                        _ = v64
                                                                                    default:
                                                                                        if (true || true) {
                                                                                            if len(v34) >= 2 {
                                                                                                v34 = v34[1:len(v34)]
                                                                                            }
                                                                                            if len(v34) >= 2 {
                                                                                                v34 = v34[1:len(v34)]
                                                                                            }
                                                                                            if v61 {
                                                                                                var v65 int = (^(fi13))
                                                                                                var v66 int8 = -128
                                                                                                var v67 int = v65
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                            } else {
                                                                                                var v68 int64 = int64(v2)
                                                                                                _ = v68
                                                                                            }
                                                                                            v69 := GenS7{F0: genFn0(750, -145), F1: "\n\t", F2: v12}
                                                                                            _ = v69
                                                                                        } else {
                                                                                            var v70 bool = v61
                                                                                            _ = v70
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v60, fi13, v60)
                                                                                } else {
                                                                                    v61 = ((2147483647 < -234) && (v61 && ((859 > 4294967295) && (v61 || false))))
                                                                                    v12 = v12
                                                                                }
                                                                                _ = v60
                                                                                _ = v61
                                                                            }
                                                                        }
                                                                        v1 -= 426
                                                                    } else {
                                                                        v71 := GenS8{F0: v2}
                                                                        if true {
                                                                            v72 := GenS9{F0: 0, F1: (1 >= 42), F2: 1}
                                                                            _ = len(v3)
                                                                            _ = v72
                                                                        }
                                                                        fmt.Println(v1, v29)
                                                                        _ = v71
                                                                    }
                                                                case "foo bar":
                                                                    v29 += v29
                                                                default:
                                                                    v73 := GenS6{F0: (v11 + v1), F1: -1}
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                    _ = v73
                                                                }
                                                                v74 := GenS5{F0: 2}
                                                                var v75 int16 = -97
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            v34 = append(v34, 128)
                                                            _ = v34
                                                        }
                                                    }
                                                case "foo bar":
                                                    v29 -= v29
                                                    v2 = v2
                                                    fmt.Println(v12)
                                                    var v76 int = fi13
                                                    _ = v76
                                                }
                                                _ = v29
                                            }
                                            var v77 uint32 = 107
                                            _ = v77
                                        }
                                    } else {
                                        if true {
                                            var v78 uint = 1
                                            var v79 int16 = int16(v1)
                                            if ((fi13 ^ -456) == (fi13 * -2147483648)) {
                                                v80 := GenS10{F0: 112, F1: v22, F2: true}
                                                var v81 uint = v78
                                                _ = v80
                                                _ = v81
                                            } else {
                                                fmt.Println(v78, v12, fi19)
                                                v82 := v3[32767]
                                                var v83 uint8 = 42
                                                _ = v82
                                                _ = v83
                                            }
                                            v84 := v3[fi19]
                                            v22 += (v12 >> ((v78) & 63))
                                            v85 := GenS3{F0: v79}
                                            v84 = v84
                                            _ = v78
                                            _ = v79
                                            _ = v84
                                            _ = v85
                                        } else {
                                            if (-1 == 1) {
                                                switch "" {
                                                case "hello":
                                                    if true {
                                                        switch (fi19 * fi19) {
                                                        case 0:
                                                            fmt.Println(v2)
                                                        case 32767:
                                                            var v86 float32 = (0.1 * 3.14159)
                                                            _ = v86
                                                        default:
                                                            var v87 int64 = (-(v22))
                                                            var v88 float32 = 0.0
                                                            fmt.Println(v87, v87)
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        v11 = (v1 << ((18446744073709551615) & 63))
                                                    }
                                                case "a":
                                                    switch "a" {
                                                    case "hello":
                                                        v89 := GenS11{F0: ("" + "a"), F1: v22, F2: (!true)}
                                                        _ = v89
                                                    default:
                                                        v90 := GenS6{F0: v1, F1: 456}
                                                        _ = v90
                                                    }
                                                }
                                                var v91 uint8 = 174
                                                _ = v91
                                            }
                                            switch ("foo bar" + "\n\t") {
                                            case "\n\t":
                                                if false {
                                                    v92 := GenS8{F0: (v2 * (v2 & 32767))}
                                                    _ = len(v3)
                                                    v22 += (v12 >> ((uint(v1)) & 63))
                                                    _ = v92
                                                } else {
                                                    if (true && true) {
                                                        v93 := []int{(fi19 % (((-(fi13))) | 1)), 2147483647, fi19, fi13}
                                                        v94 := GenS7{F0: genFn0(v1, genFn0(1, 4294967295)), F1: ("" + (("a" + "a") + "")), F2: v12}
                                                        var v95 uint8 = uint8(v2)
                                                        v3[(fi13 / (((-(fi13))) | 1))] = "hello"
                                                        if (18446744073709551615 == 445) {
                                                            if (float32(3.14159) > 1e10) {
                                                                var v96 string = ("a" + "a")
                                                                _ = v96
                                                            }
                                                        } else {
                                                            v3[int(v95)] = ""
                                                            if false {
                                                                v2 = v2
                                                                if false {
                                                                    if false {
                                                                        v97 := GenS2{F0: (v1 & v11)}
                                                                        var v98 bool = true
                                                                        var v99 int8 = 100
                                                                        _ = v97
                                                                        _ = v98
                                                                        _ = v99
                                                                    }
                                                                    var v100 int16 = 32767
                                                                    _ = v100
                                                                }
                                                            } else {
                                                                switch (fi19 >> ((1000) & 63)) {
                                                                case 127:
                                                                    v2 = 0
                                                                case 1:
                                                                    var v101 float32 = float32((((0.1 - -2.25) / (3.14159 - (1.5 * 1e10))) / -2.25))
                                                                    _ = v101
                                                                case 0:
                                                                    if false {
                                                                        switch (fi13 + (fi13 | -1)) {
                                                                        case 891:
                                                                            v1 = (^(v11))
                                                                        case 42:
                                                                            var v102 int8 = 54
                                                                            fmt.Println(v12, v11)
                                                                            _ = v102
                                                                        case 2147483647:
                                                                            v2 += v2
                                                                        default:
                                                                            v22 = -9223372036854775808
                                                                        }
                                                                        v1 = (v1 % (((v11 + (v11 - v1))) | 1))
                                                                    }
                                                                default:
                                                                    var v103 int8 = -128
                                                                    _ = v103
                                                                }
                                                                fmt.Println(v12)
                                                                var v104 uint = uint(v2)
                                                                _ = v104
                                                            }
                                                            if (643 >= 42) {
                                                                v2 += (v2 - v2)
                                                            }
                                                        }
                                                        _ = v93
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        fmt.Println(fi13, v22)
                                                    }
                                                    fmt.Println(v11, fi19)
                                                }
                                            case "foo bar":
                                                v105 := GenS12{F0: -1, F1: ((("foo bar" + "hello") + "foo bar") + "a"), F2: 103}
                                                _ = v105
                                            case "":
                                                v12 += v22
                                            }
                                            if false {
                                                fmt.Println(v22)
                                                fmt.Println(fi13, v12)
                                                var v106 float32 = float32(1e10)
                                                _ = v106
                                            } else {
                                                var v107 bool = true
                                                v108 := []int{fi13, genFn0(v1, int(v12)), fi13, fi13}
                                                if v107 {
                                                    v1 = v1
                                                    v2 = v2
                                                } else {
                                                    v1 = v11
                                                    v11 += 748
                                                }
                                                _ = v107
                                                _ = v108
                                            }
                                            v12 = v12
                                        }
                                        if true {
                                            var v109 float32 = 0.0
                                            v110 := GenS13{F0: fi19, F1: 255, F2: 42}
                                            v109 = v109
                                            _ = v109
                                            _ = v110
                                        }
                                    }
                                    v1 += (v1 * 1)
                                    if true {
                                        v12 = v12
                                    }
                                } else {
                                    var v111 int64 = (v22 % ((42) | 1))
                                    _ = v111
                                }
                            case 0:
                                switch fi13 {
                                case 2147483647:
                                    var v112 uint64 = v1
                                    _ = v112
                                default:
                                    v2 = 255
                                }
                            case 255:
                                var v113 string = ""
                                _ = v113
                            default:
                                var v114 uint32 = 2
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        v115 := GenS8{F0: (v2 & v2)}
                                        _ = v115
                                    case "hello":
                                        var v116 string = "foo bar"
                                        _ = v116
                                    case "foo bar":
                                        v117 := GenS14{F0: "\n\t", F1: 1}
                                        _ = v117
                                    default:
                                        v22 += v12
                                    }
                                    fmt.Println(v22)
                                } else {
                                    v12 += 188
                                    var v118 int32 = -2147483648
                                    _ = v118
                                }
                                v2 += v2
                                _ = v114
                            }
                            _ = v24
                        }
                    }
                    v12 += v22
                    var v119 int32 = 0
                    if false {
                        if (255 > 209) {
                            if true {
                                v119 -= (v119 ^ v119)
                            } else {
                                v120 := GenS11{F0: "hello", F1: -563, F2: false}
                                var v121 int16 = int16(v119)
                                v11 = v1
                                _ = v120
                                _ = v121
                            }
                            var v122 uint = 65535
                            _ = v122
                        } else {
                            v123 := GenS1{F0: v22, F1: -128}
                            var v124 uint32 = 2147483648
                            var v125 string = "\n\t"
                            _ = v123
                            _ = v124
                            _ = v125
                        }
                    }
                    _ = v119
                    _ = v22
                }
                switch (fi19 % ((fi13) | 1)) {
                case 474:
                    v1 += 297
                case -825:
                    v126 := []int{fi13, genFn0((^(v1)), fi13)}
                    _ = v126
                case -919:
                    fmt.Println(v2, v2)
                default:
                    v1 = v1
                    v1 += (v11 >> ((608) & 63))
                }
            }
            _ = fi19
        }
        fmt.Println(fi13, v11)
        _ = fi13
    }
    var v127 uint32 = 65535
    v128 := GenS6{F0: (v11 | v1), F1: 32767}
    v127 = 65535
    var v129 uint16 = v2
    var v130 int32 = 2147483647
    if false {
        if ("a" >= "foo bar") {
            v1 -= v1
        }
        v131 := v3[4294967295]
        _ = v131
    } else {
        v132 := v3[750]
        switch 2147483647 {
        case -2147483648:
            var v133 uint32 = (v127 + v127)
            _ = v133
        case 320:
            for fi134 := 0; fi134 < 14; fi134++ {
                if (1e10 >= (-2.25 / 1e10)) {
                    v135 := v3[(-(fi134))]
                    for fi136 := 0; fi136 < 10; fi136++ {
                        var v137 uint = 2
                        if ((!false) && (42 == 2)) {
                            v130 = v130
                            switch fi136 {
                            case 42:
                                v1 = 68
                            case 32767:
                                switch (v135 + v132) {
                                case "hello":
                                    v3[int(v11)] = ("foo bar" + (v132 + v132))
                                    var v138 uint = v137
                                    switch (v132 + v135) {
                                    case "":
                                        var v139 bool = (318 <= (v129 - v2))
                                        _ = v139
                                    default:
                                        fmt.Println(v127)
                                    }
                                    _ = v138
                                }
                            case 4294967295:
                                v129 -= 344
                            }
                        }
                        v140 := []int{(fi134 ^ (fi136 * fi134)), 791, fi136, fi134}
                        v132 = "a"
                        _ = fi136
                        _ = v137
                        _ = v140
                    }
                    v141 := GenS1{F0: (-(v12)), F1: 34}
                    _ = v135
                    _ = v141
                }
                v12 -= v12
                fmt.Println(fi134)
                _ = fi134
            }
        case 2147483647:
            var v142 uint8 = 202
            _ = v142
        }
        switch 9223372036854775807 {
        case 0:
            v143 := GenS15{F0: v129, F1: 255, F2: 127}
            _ = v143
        case 9223372036854775807:
            v127 -= v127
        case 127:
            v144 := GenS12{F0: v130, F1: v132, F2: v127}
            _ = v144
        }
        _ = v132
    }
    switch int(v12) {
    case 32768:
        switch "" {
        case "":
            fmt.Println(v2, v1, v127)
        case "\n\t":
            v145 := []int{genFn0(v1, 546), 42, 128}
            _ = v145
        case "a":
            var v146 string = ("\n\t" + "a")
            _ = v146
        default:
            switch "hello" {
            case "\n\t":
                var v147 int64 = 42
                _ = v147
            default:
                v12 -= (v12 / (((-(v12))) | 1))
            }
        }
    case -9223372036854775808:
        switch -2147483648 {
        case 853:
            v148 := GenS4{F0: 42, F1: (v130 + v130), F2: (1 > uint16(v127))}
            _ = v148
        case 802:
            switch 452 {
            case 32767:
                v149 := []int{}
                _ = v149
            case 679:
                v127 = 2147483647
            }
        case 10:
            for fi150 := 0; fi150 < 20; fi150++ {
                var v151 uint64 = v1
                var v152 int32 = v130
                for fi153 := 0; fi153 < 11; fi153++ {
                    fmt.Println(fi153, v127, v2)
                    _ = fi153
                }
                _ = fi150
                _ = v151
                _ = v152
            }
        }
    case 128:
        if ((false || ((!false) || ((v11 + (v1 >> ((0) & 63))) != v1))) && (!("a" < "hello"))) {
            v11 = v11
        } else {
            v154 := GenS16{F0: 0, F1: "foo bar"}
            v155 := GenS10{F0: 255, F1: (^(v12)), F2: true}
            for fi156 := 0; fi156 < 11; fi156++ {
                v11 = uint64(v129)
                switch (^(fi156)) {
                case 0:
                    _ = len(v3)
                }
                v157 := GenS17{F0: v11, F1: v1}
                _ = fi156
                _ = v157
            }
            _ = v154
            _ = v155
        }
    default:
        var v158 float32 = (3.14159 - 0.0)
        _ = v158
    }
    v129 -= (v129 / (((v2 & (v129 - v2))) | 1))
    for fi159 := 0; fi159 < 16; fi159++ {
        v160 := v3[440]
        v12 -= (v12 - (v12 % (((v12 * v12)) | 1)))
        var v161 int = genFn0((v11 >> ((0) & 63)), genFn0(4294967295, fi159))
        _ = fi159
        _ = v160
        _ = v161
    }
    for fi162 := 0; fi162 < 16; fi162++ {
        switch ("\n\t" + "") {
        case "foo bar":
            v1 = (v1 * v11)
        default:
            v163 := GenS5{F0: 1}
            _ = v163
        }
        if true {
            for fi164 := 0; fi164 < 1; fi164++ {
                v130 = (v130 ^ v130)
                v165 := v3[(-(fi164))]
                v166 := GenS18{F0: ((v129 + (v2 % (((^(v2))) | 1))) == v129), F1: fi164, F2: (v129 - (^(v129)))}
                _ = fi164
                _ = v165
                _ = v166
            }
        } else {
            fmt.Println(v129, v2)
            var v167 int = fi162
            _ = v167
        }
        v168 := GenS19{F0: 2147483647, F1: v1}
        _ = fi162
        _ = v168
    }
    for fi169 := 0; fi169 < 11; fi169++ {
        var v170 uint32 = v127
        _ = fi169
        _ = v170
    }
    var v171 float32 = 0.1
    if false {
        switch "\n\t" {
        case "\n\t":
            v172 := GenS18{F0: false, F1: int(v127), F2: 2}
            _ = v172
        case "a":
            fmt.Println(v12)
        case "":
            fmt.Println(v129, v171)
        }
        v173 := []int{1}
        var v174 int32 = (v130 - v130)
        _ = v173
        _ = v174
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v127:", v127)
    fmt.Println("v129:", v129)
    fmt.Println("v130:", v130)
    fmt.Println("v171:", v171)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v127
    _ = v128
    _ = v129
    _ = v130
    _ = v171
    _ = v2
    _ = v3
}
