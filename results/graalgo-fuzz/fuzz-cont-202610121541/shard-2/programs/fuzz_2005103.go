package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 uint; F2 uint }
type GenS9 struct { F0 uint16; F1 int16; F2 bool }
func genFn0(p0 string, p1 bool) uint64 {
    v4 := GenS1{F0: 255, F1: 4294967295}
    for fi5 := 0; fi5 < 19; fi5++ {
        p1 = (!true)
        for fi6 := 0; fi6 < 7; fi6++ {
            fmt.Println(fi5)
            v7 := []int{}
            p0 = "\n\t"
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{32767, 0, 4294967295, 351}
    for fi9 := 0; fi9 < 11; fi9++ {
        v8 = append(v8, fi9)
        fmt.Println(fi9, fi9)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v8
    return 18446744073709551615
}
func genFn1(p0 int8) uint32 {
    var v10 int = 1
    p0 -= -1
    if false {
        switch ("a" + "hello") {
        case "":
            v11 := GenS1{F0: 254, F1: v10}
            _ = v11
        case "hello":
            v12 := GenS1{F0: 0, F1: v10}
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 1; fi13++ {
                switch fi13 {
                case -1:
                    var v14 string = (("hello" + "hello") + (("\n\t" + "a") + "foo bar"))
                    _ = v14
                case 128:
                    fmt.Println(fi13)
                case 1:
                    var v15 int32 = -1
                    _ = v15
                }
                _ = fi13
            }
        }
        var v16 float64 = -2.25
        _ = v16
    } else {
        v17 := []int{(v10 % ((v10) | 1))}
        p0 -= p0
        _ = v17
    }
    switch (v10 - v10) {
    case 4294967295:
        if true {
            p0 += (p0 << ((42) & 63))
        } else {
            v18 := []int{v10, v10, v10, v10}
            _ = v18
        }
    default:
        var v19 float64 = ((-2.25 - -2.25) * -2.25)
        _ = v19
    }
    _ = p0
    _ = v10
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 42
    var v3 uint32 = 273
    v3 -= v3
    if true {
        v20 := GenS1{F0: uint8(v3), F1: 32768}
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int32 = 0
            var v23 int = fi21
            if true {
                v24 := GenS1{F0: 2, F1: (v23 & fi21)}
                _ = v24
            } else {
                for fi25 := 0; fi25 < 20; fi25++ {
                    if (!true) {
                        if ((fi21 == (fi25 | (v23 << ((65535) & 63)))) && (!(((0.1 / 1.0) + 3.14159) < 0.1))) {
                            switch 32768 {
                            case 127:
                                v3 += v3
                            case -996:
                                if false {
                                    if false {
                                        fmt.Println(v3)
                                        switch (fi21 << ((4294967295) & 63)) {
                                        case 1:
                                            fmt.Println(v2, v2)
                                            var v26 uint32 = v3
                                            _ = v26
                                        }
                                        v27 := GenS2{F0: v3}
                                        if (1.5 < (0.0 + 0.1)) {
                                            v28 := GenS1{F0: 200, F1: (-(v23))}
                                            if false {
                                                fmt.Println(fi25, fi21, fi21)
                                                switch -2147483648 {
                                                case -473:
                                                    if (-152 >= -81) {
                                                        var v29 uint32 = 2147483647
                                                        _ = v29
                                                    }
                                                default:
                                                    v30 := GenS3{F0: (v22 & v22), F1: 42}
                                                    _ = v30
                                                }
                                            }
                                            fmt.Println(v3, v23, fi25)
                                            _ = v28
                                        }
                                        if false {
                                            v22 -= int32(v2)
                                        } else {
                                            v31 := GenS3{F0: (v22 + v22), F1: (fi21 << ((326) & 63))}
                                            var v32 int16 = -32768
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v27
                                    } else {
                                        if false {
                                            if (!true) {
                                                v33 := []int{128, (fi25 + 134), fi25}
                                                if false {
                                                    if (2 <= 18446744073709551615) {
                                                        fmt.Println(v23)
                                                    } else {
                                                        v34 := v1[915]
                                                        var v35 int8 = 100
                                                        var v36 uint16 = v2
                                                        if true {
                                                            if (772 == 9223372036854775807) {
                                                                if len(v33) >= 2 {
                                                                    v33 = v33[1:len(v33)]
                                                                }
                                                                var v37 float32 = 1e10
                                                                fmt.Println(v2, v2, v2)
                                                                _ = v37
                                                            }
                                                            v22 = (v22 % (((v22 * 2147483647)) | 1))
                                                        } else {
                                                            v38 := GenS3{F0: v22, F1: 4294967295}
                                                            v39 := GenS4{F0: "\n\t", F1: (v3 & v3)}
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                } else {
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v40 uint16 = (v2 | 42)
                                                        _ = v40
                                                    case "a":
                                                        if false {
                                                            var v41 float64 = (-2.25 / (3.14159 - (0.1 * (0.0 - (1.5 * -2.25)))))
                                                            _ = len(v1)
                                                            var v42 uint32 = 2
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    case "\n\t":
                                                        if false {
                                                            v33 = append(v33, 4294967295)
                                                            fmt.Println(fi21, fi21)
                                                            fmt.Println(fi25, fi25)
                                                            var v43 string = "hello"
                                                            v43 = "a"
                                                            _ = v43
                                                        } else {
                                                            v3 = 770
                                                            var v44 uint16 = 255
                                                            var v45 int16 = 1000
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    }
                                                    v46 := GenS4{F0: ("a" + "\n\t"), F1: v3}
                                                    _ = v46
                                                }
                                                v22 = (v22 / (((^(v22))) | 1))
                                                _ = v33
                                            }
                                            var v47 float64 = 1.5
                                            v48 := []int{fi21, v23, (fi21 | fi25)}
                                            var v49 string = "hello"
                                            if true {
                                                if true {
                                                    v47 = v47
                                                }
                                                v50 := v1[128]
                                                if false {
                                                    v51 := GenS5{F0: -1, F1: 2}
                                                    var v52 int32 = (v22 << ((4294967295) & 63))
                                                    var v53 string = v50
                                                    v54 := GenS1{F0: 1, F1: fi25}
                                                    _ = v51
                                                    _ = v52
                                                    _ = v53
                                                    _ = v54
                                                }
                                                _ = v50
                                            } else {
                                                _ = len(v48)
                                                fmt.Println(fi25, v3)
                                            }
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        } else {
                                            v3 = (^(v3))
                                        }
                                        v1[(-(fi21))] = "\n\t"
                                        v55 := GenS2{F0: 42}
                                        _ = v55
                                    }
                                    if (0 < 42) {
                                        if false {
                                            fmt.Println(v3, v22, fi25)
                                            var v56 bool = false
                                            if ("\n\t" >= "\n\t") {
                                                v57 := GenS3{F0: -1, F1: v23}
                                                v58 := []int{}
                                                if v56 {
                                                    switch (-(v23)) {
                                                    case -1:
                                                        var v59 uint8 = 2
                                                        _ = v59
                                                    case -895:
                                                        if v56 {
                                                            v3 -= (v3 * (v3 >> ((271) & 63)))
                                                            v60 := GenS1{F0: 2, F1: fi21}
                                                            fmt.Println(fi25)
                                                            _ = v60
                                                        }
                                                    default:
                                                        var v61 uint16 = (v2 % (((v2 | 769)) | 1))
                                                        _ = v61
                                                    }
                                                    if v56 {
                                                        v62 := GenS2{F0: (v3 + (v3 & v3))}
                                                        v58 = append(v58, fi25)
                                                        var v63 uint = 1
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        v58 = append(v58, fi21)
                                                        var v64 uint8 = 2
                                                        var v65 uint64 = 295
                                                        v64 = v64
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    delete(v1, (-(v23)))
                                                } else {
                                                    fmt.Println(v23, fi25)
                                                    if (!false) {
                                                        var v66 float32 = (float32((3.14159 - 1.5)) * 0.0)
                                                        delete(v1, (^(fi21)))
                                                        _ = v66
                                                    }
                                                }
                                                _ = v57
                                                _ = v58
                                            }
                                            _ = v56
                                        }
                                        v67 := []int{(fi25 * -377), fi21, fi25, (v23 + v23)}
                                        _ = v67
                                    } else {
                                        v3 -= (v3 | genFn1(-12))
                                        v1[(fi21 & -2147483648)] = "hello"
                                    }
                                } else {
                                    if false {
                                        v23 = fi25
                                        switch (v23 / (((fi21 >> ((1) & 63))) | 1)) {
                                        case 127:
                                            v1[fi21] = ""
                                        }
                                        var v68 string = "\n\t"
                                        _ = v68
                                    }
                                    var v69 uint32 = (v3 * 65535)
                                    if (v2 != v2) {
                                        var v70 int8 = 100
                                        _ = v70
                                    } else {
                                        v71 := []int{(^(v23)), 363, (v23 & int(v2)), 815}
                                        switch "hello" {
                                        case "hello":
                                            v23 = (v23 & (fi21 / (((v23 | fi25)) | 1)))
                                        case "a":
                                            v71 = append(v71, (fi21 | -2147483648))
                                        default:
                                            _ = len(v71)
                                        }
                                        v72 := v1[(-(v23))]
                                        _ = v71
                                        _ = v72
                                    }
                                    _ = v69
                                }
                            case 128:
                                v1[42] = ("\n\t" + "")
                            default:
                                var v73 float64 = ((0.0 * 1.5) + ((0.1 + -2.25) - ((((((0.0 * 0.0) + (3.14159 - 3.14159)) * (0.0 / 0.1)) - (((1e10 - 0.0) * (1.5 / 1e10)) / 1e10)) + float64(0.1)) / float64(1.5))))
                                _ = v73
                            }
                        } else {
                            v3 -= 0
                            fmt.Println(fi25, v23, v22)
                            var v74 uint = 160
                            if (1e10 != 1e10) {
                                var v75 float64 = -2.25
                                _ = v75
                            } else {
                                var v76 float32 = (float32(float64(-2.25)) - 0.1)
                                _ = v76
                            }
                            if (!(-207 >= 419)) {
                                switch v23 {
                                case 325:
                                    switch "a" {
                                    case "":
                                        if ("a" < "hello") {
                                            v77 := v1[-377]
                                            _ = v77
                                        } else {
                                            v78 := []int{-2147483648, (v23 % ((v23) | 1))}
                                            _ = v78
                                        }
                                    case "\n\t":
                                        v1[fi25] = "\n\t"
                                    case "foo bar":
                                        var v79 int = 1
                                        _ = v79
                                    default:
                                        switch fi25 {
                                        case 128:
                                            if true {
                                                var v80 uint8 = 194
                                                _ = v80
                                            } else {
                                                if false {
                                                    v3 -= 0
                                                    var v81 int8 = 127
                                                    _ = v81
                                                }
                                            }
                                        case -9223372036854775808:
                                            v82 := []int{(fi25 << ((v74) & 63)), v23, v23, 4294967295}
                                            v22 = -655
                                            _ = v82
                                        case 127:
                                            var v83 float64 = (3.14159 + ((1e10 / 0.1) - (1.5 / (3.14159 / 1.0))))
                                            _ = v83
                                        default:
                                            var v84 uint = (v74 % (((v74 + 4294967295)) | 1))
                                            _ = v84
                                        }
                                    }
                                    if ("\n\t" != ("a" + "foo bar")) {
                                        v85 := GenS6{F0: ""}
                                        if (!(!(v2 == v2))) {
                                            v86 := []int{v23, v23, 32767}
                                            _ = v86
                                        } else {
                                            v23 += fi21
                                        }
                                        v23 -= fi25
                                        _ = v85
                                    }
                                    if true {
                                        v87 := GenS4{F0: "", F1: 320}
                                        v23 += (fi21 << ((v74) & 63))
                                        v88 := []int{(fi25 ^ fi25), v23, v23, v23}
                                        _ = v87
                                        _ = v88
                                    } else {
                                        v22 = (^(v22))
                                        v1[-376] = "hello"
                                        var v89 float32 = 0.1
                                        _ = v89
                                    }
                                    if false {
                                        fmt.Println(v3)
                                        fmt.Println(v23, v74)
                                        fmt.Println(fi25, v3)
                                        v90 := v1[(-(v23))]
                                        var v91 float32 = 1e10
                                        v2 = v2
                                        if (true || false) {
                                            v92 := GenS3{F0: 1, F1: -1}
                                            v1[fi25] = v90
                                            _ = v92
                                        } else {
                                            v91 += v91
                                        }
                                        _ = v90
                                        _ = v91
                                    } else {
                                        if true {
                                            if true {
                                                v74 += 65535
                                                switch fi25 {
                                                case -443:
                                                    v93 := GenS3{F0: -877, F1: (fi25 % (((fi25 * 4294967295)) | 1))}
                                                    _ = v93
                                                }
                                            }
                                            var v94 uint = (v74 + 619)
                                            _ = v94
                                        } else {
                                            v95 := []int{}
                                            _ = v95
                                        }
                                    }
                                    if true {
                                        var v96 bool = false
                                        var v97 uint = 0
                                        _ = v96
                                        _ = v97
                                    } else {
                                        switch 309 {
                                        case 127:
                                            var v98 float64 = 1.5
                                            _ = v98
                                        case 1:
                                            switch (v23 >> (((v74 << ((v74) & 63))) & 63)) {
                                            case 0:
                                                v2 = (v2 + 32767)
                                            default:
                                                v23 += (v23 % (((fi21 >> ((v74) & 63))) | 1))
                                            }
                                        case 42:
                                            fmt.Println(v74)
                                        default:
                                            if true {
                                                var v99 float32 = (1.5 / float32(3.14159))
                                                v2 = v2
                                                fmt.Println(v74, v74)
                                                _ = v99
                                            }
                                        }
                                        if true {
                                            v3 -= (v3 >> ((v74) & 63))
                                        }
                                    }
                                    v2 += 94
                                    v74 = 0
                                    fmt.Println(v2, fi25)
                                case 42:
                                    v3 = 2147483648
                                case 433:
                                    if false {
                                        fmt.Println(v22, v22)
                                        fmt.Println(v3)
                                    } else {
                                        v22 = (v22 ^ v22)
                                    }
                                }
                                var v100 uint32 = 2147483647
                                if (!(("hello" + (("a" + ("hello" + "")) + "foo bar")) >= "\n\t")) {
                                    var v101 uint32 = genFn1(-128)
                                    _ = v101
                                }
                                _ = v100
                            } else {
                                var v102 uint32 = v3
                                v23 = v23
                                v102 -= genFn1(-1)
                                _ = v102
                            }
                            v103 := v1[741]
                            if true {
                                switch (v103 + (("foo bar" + v103) + v103)) {
                                case "":
                                    var v104 uint = v74
                                    _ = v104
                                case "foo bar":
                                    v105 := v1[1]
                                    _ = v105
                                }
                                var v106 float64 = 1e10
                                var v107 uint = (v74 / (((v74 >> ((4294967295) & 63))) | 1))
                                fmt.Println(v22, v74)
                                _ = v106
                                _ = v107
                            }
                            _ = v103
                            _ = v74
                        }
                    } else {
                        v3 += (v3 >> ((255) & 63))
                        if (false && true) {
                            v108 := GenS7{F0: v2}
                            _ = v108
                        }
                        fmt.Println(v3, v3)
                    }
                    _ = fi25
                }
                v109 := v1[v23]
                _ = v109
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = v20
    } else {
        v2 -= v2
        var v110 uint64 = genFn0(("hello" + (("hello" + "a") + "foo bar")), (0.1 == 0.1))
        if true {
            v110 = genFn0("", false)
            var v111 int = int(v2)
            var v112 float64 = (-2.25 + float64(0.0))
            _ = v111
            _ = v112
        } else {
            if false {
                switch -722 {
                case -9223372036854775808:
                    switch -1 {
                    case 313:
                        for fi113 := 0; fi113 < 4; fi113++ {
                            var v114 float64 = (0.1 * 0.0)
                            fmt.Println(v114, v2, v3)
                            _ = fi113
                            _ = v114
                        }
                    default:
                        fmt.Println(v3, v110)
                    }
                }
                if (1e10 == (float64(1e10) * 3.14159)) {
                    v3 -= (^(v3))
                }
            } else {
                for fi115 := 0; fi115 < 12; fi115++ {
                    var v116 int = fi115
                    _ = fi115
                    _ = v116
                }
                if true {
                    fmt.Println(v110, v3)
                    var v117 uint16 = (^(v2))
                    fmt.Println(v3, v3)
                    _ = v117
                }
                var v118 int16 = 0
                _ = v118
            }
        }
        _ = v110
    }
    v3 = (v3 % ((genFn1(-67)) | 1))
    fmt.Println(v2, v2)
    v119 := []int{2147483647, -673, 127}
    fmt.Println(v2)
    switch 860 {
    case 851:
        if false {
            for fi120 := 0; fi120 < 2; fi120++ {
                switch fi120 {
                case -9223372036854775808:
                    v2 += 65535
                case 688:
                    v2 = v2
                default:
                    var v121 uint64 = 1
                    _ = v121
                }
                _ = fi120
            }
            v119 = append(v119, 73)
        }
    case 1:
        for fi122 := 0; fi122 < 19; fi122++ {
            v3 = v3
            var v123 int16 = -32768
            if (-9223372036854775808 >= -9223372036854775808) {
                fmt.Println(v2)
                if false {
                    v124 := GenS8{F0: (236 != 42), F1: 71, F2: 545}
                    for fi125 := 0; fi125 < 8; fi125++ {
                        v123 += (v123 % ((v123) | 1))
                        _ = fi125
                    }
                    var v126 int64 = 42
                    _ = v124
                    _ = v126
                } else {
                    v3 = v3
                }
                switch "foo bar" {
                case "hello":
                    if (true && false) {
                        for fi127 := 0; fi127 < 18; fi127++ {
                            if (0 <= -128) {
                                v123 -= (v123 * -484)
                                switch "foo bar" {
                                case "\n\t":
                                    if (0 < uint64(v3)) {
                                        v128 := GenS5{F0: -864, F1: 18446744073709551615}
                                        var v129 uint16 = (v2 - v2)
                                        _ = v128
                                        _ = v129
                                    }
                                default:
                                    fmt.Println(v123)
                                }
                            }
                            _ = fi127
                        }
                        v3 = genFn1(-128)
                    }
                case "foo bar":
                    if false {
                        _ = len(v119)
                        if (true || true) {
                            var v130 uint = 1
                            v130 += v130
                            var v131 int16 = (v123 % ((-32768) | 1))
                            _ = v130
                            _ = v131
                        } else {
                            var v132 uint = 853
                            _ = v132
                        }
                    }
                default:
                    for fi133 := 0; fi133 < 20; fi133++ {
                        v134 := v1[(fi122 + (^(fi133)))]
                        var v135 uint32 = 1
                        _ = fi133
                        _ = v134
                        _ = v135
                    }
                }
            }
            _ = fi122
            _ = v123
        }
    case 32767:
        fmt.Println(v2, v2)
    default:
        delete(v1, -378)
    }
    v136 := GenS8{F0: (1 == -1), F1: 42, F2: 4294967295}
    v3 = v3
    for fi137 := 0; fi137 < 15; fi137++ {
        v138 := GenS4{F0: ((("" + "a") + ("\n\t" + "a")) + "\n\t"), F1: genFn1(-6)}
        var v139 int16 = 1000
        _ = fi137
        _ = v138
        _ = v139
    }
    if false {
        for fi140 := 0; fi140 < 19; fi140++ {
            fmt.Println(v2, v2, fi140)
            var v141 uint32 = v3
            if false {
                for fi142 := 0; fi142 < 1; fi142++ {
                    v143 := GenS3{F0: 1, F1: 170}
                    _ = fi142
                    _ = v143
                }
            } else {
                var v144 int16 = 42
                var v145 uint16 = v2
                switch -481 {
                case 127:
                    var v146 int = 852
                    _ = v146
                }
                _ = v144
                _ = v145
            }
            _ = fi140
            _ = v141
        }
        fmt.Println(v3, v2)
    } else {
        var v147 float32 = 1.5
        _ = v147
    }
    var v148 float64 = (1e10 * 3.14159)
    var v149 bool = (416 > 1)
    v119 = append(v119, -9223372036854775808)
    v2 = v2
    v150 := GenS6{F0: ""}
    v151 := GenS9{F0: v2, F1: 42, F2: false}
    var v152 uint8 = 83
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v119:", len(v119), v119)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v152:", v152)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v119
    _ = v136
    _ = v148
    _ = v149
    _ = v150
    _ = v151
    _ = v152
    _ = v2
    _ = v3
}
