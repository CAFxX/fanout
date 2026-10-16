package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint32 }
type GenS6 struct { F0 int64; F1 int64; F2 bool }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 uint; F2 int32 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint8, p1 int8) bool {
    v5 := GenS1{F0: 579, F1: ((false && false) && true), F2: 2147483648}
    v6 := GenS1{F0: 42, F1: (42 <= 1), F2: 4294967295}
    p1 = int8(p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!true)
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 0.0
    v3 := make(map[string]int)
    var v4 int64 = 42
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float32 = 3.14159
        for fi9 := 0; fi9 < 6; fi9++ {
            v4 += v4
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    if (71 < 2147483647) {
        v10 := []int{9223372036854775807}
        var v11 int32 = 1
        switch 128 {
        case 646:
            var v12 int64 = -123
            _ = v12
        default:
            var v13 int32 = (v11 ^ (v11 - v11))
            _ = v13
        }
        _ = v10
        _ = v11
    } else {
        v2 -= float64((((1e10 - (1.5 / (1e10 + 1.5))) + float32(-2.25)) + 1.5))
        v1 -= (v1 >> ((566) & 63))
        for fi14 := 0; fi14 < 11; fi14++ {
            v1 = 1000
            switch fi14 {
            case 0:
                var v15 string = "hello"
                _ = v15
            case 127:
                switch 9223372036854775807 {
                case -742:
                    if ((v1 % (((v1 % ((-32768) | 1))) | 1)) >= (v1 % ((243) | 1))) {
                        switch (^(fi14)) {
                        case 1:
                            v16 := []int{(fi14 & (fi14 / (((fi14 % ((408) | 1))) | 1))), fi14, (fi14 + fi14), (fi14 * 32767)}
                            _ = v16
                        default:
                            switch -684 {
                            case 42:
                                switch fi14 {
                                case 9223372036854775807:
                                    v17 := GenS1{F0: 65535, F1: false, F2: 2}
                                    _ = v17
                                }
                            case 1:
                                for fi18 := 0; fi18 < 14; fi18++ {
                                    if (!false) {
                                        switch (fi18 - fi14) {
                                        case -563:
                                            fmt.Println(v4, v2)
                                        case -1:
                                            v19 := GenS1{F0: 42, F1: (1 != 42), F2: 0}
                                            _ = v19
                                        default:
                                            v20 := GenS2{F0: 967, F1: 18446744073709551615}
                                            var v21 float32 = 1.5
                                            fmt.Println(v21, fi14, v4)
                                            v2 = float64(float32((0.1 / float64(1.5))))
                                            fmt.Println(v2)
                                            _ = v20
                                            _ = v21
                                        }
                                    } else {
                                        v22 := []int{}
                                        var v23 uint64 = 0
                                        v22 = append(v22, fi18)
                                        _ = v22
                                        _ = v23
                                    }
                                    if false {
                                        var v24 bool = ((v4 << ((65535) & 63)) != (-(v4)))
                                        switch ("hello" + "") {
                                        case "hello":
                                            v25 := []int{(fi18 >> ((281) & 63))}
                                            _ = v25
                                        }
                                        _ = v24
                                    }
                                    var v26 uint16 = 255
                                    _ = fi18
                                    _ = v26
                                }
                            case -847:
                                v4 += (v4 & v4)
                            }
                        }
                        v1 = v1
                    } else {
                        var v27 int8 = 1
                        _ = v27
                    }
                case 12:
                    var v28 int = fi14
                    _ = v28
                case 8:
                    for fi29 := 0; fi29 < 3; fi29++ {
                        var v30 float64 = float64(((3.14159 * 0.1) / (1.5 / (1.5 - float32((v2 / 1.5))))))
                        if (-1 <= (fi29 | fi29)) {
                            switch "hello" {
                            case "\n\t":
                                v31 := GenS3{F0: -638, F1: 4294967295, F2: 2}
                                _ = v31
                            case "hello":
                                v2 -= v30
                            case "foo bar":
                                switch (-(fi14)) {
                                case 2147483647:
                                    v3[""] = (fi14 >> ((42) & 63))
                                case 634:
                                    var v32 int16 = (v1 % (((v1 - (v1 - (^(v1))))) | 1))
                                    _ = v32
                                }
                            }
                        } else {
                            v33 := GenS1{F0: 2, F1: false, F2: 65535}
                            var v34 int32 = 0
                            fmt.Println(v2, fi29)
                            if (!true) {
                                v35 := []int{fi14, 479, fi14}
                                _ = v35
                            }
                            var v36 string = ("a" + "hello")
                            if (1 != 65535) {
                                v4 += v4
                            } else {
                                switch 4294967295 {
                                case 9223372036854775807:
                                    v37 := GenS1{F0: 308, F1: ("hello" == "a"), F2: 4294967295}
                                    _ = v37
                                case -1:
                                    if (!(float32(v2) > 1.5)) {
                                        v38 := v3[v36]
                                        v1 = 32767
                                        _ = v38
                                    }
                                case 0:
                                    var v39 int8 = 0
                                    _ = v39
                                default:
                                    var v40 int16 = -32768
                                    _ = v40
                                }
                                var v41 string = (v36 + v36)
                                var v42 bool = (int(v34) != fi29)
                                _ = v41
                                _ = v42
                            }
                            _ = v33
                            _ = v34
                            _ = v36
                        }
                        _ = fi29
                        _ = v30
                    }
                default:
                    for fi43 := 0; fi43 < 15; fi43++ {
                        switch fi43 {
                        case 119:
                            v1 = 42
                        case 9223372036854775807:
                            var v44 uint8 = 214
                            _ = v44
                        }
                        _ = fi43
                    }
                }
            case 128:
                for fi45 := 0; fi45 < 9; fi45++ {
                    if (1.5 < 3.14159) {
                        if true {
                            var v46 float32 = 0.0
                            var v47 uint16 = 718
                            var v48 float32 = (-2.25 * v46)
                            _ = v46
                            _ = v47
                            _ = v48
                        } else {
                            v2 -= v2
                            var v49 int = (fi14 | (fi14 >> ((18446744073709551615) & 63)))
                            if (false && (255 >= 0)) {
                                if (42 == 498) {
                                    var v50 uint16 = 42
                                    var v51 uint8 = 255
                                    switch "a" {
                                    case "a":
                                        v52 := GenS4{F0: true}
                                        _ = v52
                                    case "foo bar":
                                        var v53 int32 = 42
                                        _ = v53
                                    case "hello":
                                        if false {
                                            switch int(v50) {
                                            case 4294967295:
                                                v54 := GenS2{F0: v1, F1: 296}
                                                _ = v54
                                            }
                                            if true {
                                                v55 := GenS5{F0: 622, F1: 928}
                                                _ = len(v3)
                                                if (!("\n\t" <= "")) {
                                                    v56 := GenS3{F0: v1, F1: 4294967295, F2: 255}
                                                    v51 += (^(v51))
                                                    _ = v56
                                                }
                                                _ = v55
                                            } else {
                                                var v57 uint64 = 4294967295
                                                var v58 uint32 = 188
                                                v2 = 1.5
                                                fmt.Println(v2, v4)
                                                _ = v57
                                                _ = v58
                                            }
                                            v59 := []int{(-(v49)), (v49 + 128)}
                                            _ = v59
                                        }
                                    }
                                    v60 := []int{fi45, (v49 ^ (fi45 >> ((4294967295) & 63)))}
                                    var v61 uint64 = 0
                                    _ = v50
                                    _ = v51
                                    _ = v60
                                    _ = v61
                                } else {
                                    var v62 float64 = float64(float32(float64(0.1)))
                                    _ = v62
                                }
                                v63 := []int{fi14}
                                _ = v63
                            } else {
                                fmt.Println(v1, v49, v49)
                                fmt.Println(v4)
                            }
                            var v64 float32 = -2.25
                            var v65 int16 = v1
                            switch (("" + "") + ("hello" + "\n\t")) {
                            case "hello":
                                fmt.Println(v65)
                            case "\n\t":
                                v1 = -814
                            default:
                                switch fi45 {
                                case 4294967295:
                                    if (-459 == (-(fi14))) {
                                        v3["foo bar"] = 32768
                                        var v66 float32 = ((float32(float64(v64)) - v64) * (0.0 / v64))
                                        v49 += 32767
                                        v64 = (v66 / -2.25)
                                        _ = v66
                                    } else {
                                        _ = len(v3)
                                        v49 += -253
                                        v49 = -40
                                    }
                                    v49 = fi45
                                case 1:
                                    var v67 int16 = (-(v65))
                                    _ = v67
                                default:
                                    v68 := GenS2{F0: (v1 + (v65 >> ((65535) & 63))), F1: 0}
                                    _ = v68
                                }
                            }
                            _ = v49
                            _ = v64
                            _ = v65
                        }
                        v2 = float64(-2.25)
                        fmt.Println(v1)
                    } else {
                        v69 := GenS2{F0: (v1 % ((1) | 1)), F1: 0}
                        _ = v69
                    }
                    _ = fi45
                }
            }
            var v70 string = "\n\t"
            _ = fi14
            _ = v70
        }
    }
    if (-1 == 42) {
        var v71 int = 0
        v72 := GenS1{F0: 237, F1: (v4 == v4), F2: 886}
        v71 = (v71 - (v71 / ((42) | 1)))
        _ = v71
        _ = v72
    }
    if false {
        if true {
            if false {
                fmt.Println(v4, v4)
            } else {
                var v73 int32 = 512
                var v74 uint = 95
                var v75 uint64 = 956
                _ = v73
                _ = v74
                _ = v75
            }
            switch 127 {
            case -393:
                v3[("\n\t" + "")] = 42
            case -420:
                for fi76 := 0; fi76 < 13; fi76++ {
                    if (0 == 42) {
                        v3["hello"] = fi76
                    } else {
                        v77 := GenS6{F0: -9223372036854775808, F1: v4, F2: true}
                        if (65535 <= 42) {
                            v78 := GenS6{F0: v4, F1: (v4 | v4), F2: (162 <= 4294967295)}
                            fmt.Println(v2, v2, v2)
                            _ = v78
                        } else {
                            for fi79 := 0; fi79 < 6; fi79++ {
                                v80 := []int{(^(fi76)), (fi76 * 681)}
                                if (!(((true && true) || (42 <= 43)) && false)) {
                                    fmt.Println(v4)
                                } else {
                                    fmt.Println(fi76, fi76)
                                    fmt.Println(v1, fi76, v4)
                                }
                                v3["a"] = fi79
                                _ = fi79
                                _ = v80
                            }
                            v81 := []int{fi76}
                            _ = v81
                        }
                        v4 -= v4
                        _ = v77
                    }
                    for fi82 := 0; fi82 < 6; fi82++ {
                        fmt.Println(v4)
                        fmt.Println(v2, v1, fi76)
                        if (false || false) {
                            v83 := GenS6{F0: v4, F1: (-(v4)), F2: false}
                            if false {
                                fmt.Println(fi82, fi76)
                            } else {
                                fmt.Println(v4, v1, fi82)
                                if ((float32(v2) * (((-2.25 / 3.14159) - (1e10 - float32(v2))) - ((-2.25 * float32(v2)) - float32((((((3.14159 * (v2 / v2)) - v2) / (1e10 + v2)) + 1.5) + 0.0))))) >= 3.14159) {
                                    v4 = (v4 % ((v4) | 1))
                                }
                                if false {
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v84 int64 = 1
                                        _ = v84
                                    case "":
                                        if true {
                                            v1 -= v1
                                            v85 := GenS4{F0: (("foo bar" + "a") != "\n\t")}
                                            v4 = (v4 - v4)
                                            _ = v85
                                        } else {
                                            var v86 uint16 = 255
                                            v86 = (v86 << ((255) & 63))
                                            _ = v86
                                        }
                                    default:
                                        v87 := v3[("\n\t" + ((("" + "") + ("a" + ("" + (((("foo bar" + "") + "foo bar") + "") + "hello")))) + "foo bar"))]
                                        _ = v87
                                    }
                                    v4 -= int64(fi82)
                                    v88 := []int{fi76, fi76}
                                    _ = v88
                                }
                            }
                            _ = v83
                        } else {
                            if (!false) {
                                v89 := GenS7{F0: 1, F1: 607}
                                v90 := GenS2{F0: v1, F1: 0}
                                _ = v89
                                _ = v90
                            }
                            v4 += (v4 * (v4 | (v4 >> ((4294967295) & 63))))
                            switch (fi76 ^ 32767) {
                            case 32767:
                                v91 := GenS8{F0: 255, F1: (fi76 & fi76)}
                                _ = v91
                            }
                        }
                        _ = fi82
                    }
                    if ((1 != 42) || false) {
                        v4 -= v4
                        for fi92 := 0; fi92 < 5; fi92++ {
                            v4 = 1
                            fmt.Println(fi76, v4, v4)
                            var v93 int = 847
                            _ = fi92
                            _ = v93
                        }
                        v4 += (^(v4))
                    } else {
                        v1 = v1
                        if false {
                            if false {
                                var v94 uint = 151
                                for fi95 := 0; fi95 < 12; fi95++ {
                                    var v96 int = (^(fi76))
                                    v97 := GenS3{F0: 42, F1: (v94 + v94), F2: 255}
                                    if (false || (0 == 32767)) {
                                        v1 = (v1 * v1)
                                        v98 := GenS9{F0: 2}
                                        v4 = (^(v4))
                                        _ = v98
                                    }
                                    _ = fi95
                                    _ = v96
                                    _ = v97
                                }
                                _ = v94
                            }
                            for fi99 := 0; fi99 < 12; fi99++ {
                                fmt.Println(v4)
                                var v100 float32 = (float32(v2) + -2.25)
                                if true {
                                    fmt.Println(v1)
                                }
                                _ = fi99
                                _ = v100
                            }
                        } else {
                            for fi101 := 0; fi101 < 10; fi101++ {
                                var v102 float64 = (v2 - v2)
                                if (2 >= 2) {
                                    if (!false) {
                                        switch ("\n\t" + "\n\t") {
                                        case "\n\t":
                                            fmt.Println(v2)
                                        case "":
                                            if true {
                                                fmt.Println(fi76, v2, v102)
                                                v103 := GenS9{F0: 0}
                                                _ = v103
                                            } else {
                                                var v104 int32 = 42
                                                v2 += (1e10 - ((v2 * 1e10) - -2.25))
                                                delete(v3, "hello")
                                                _ = v104
                                            }
                                        default:
                                            switch int(v4) {
                                            case 42:
                                                var v105 int16 = v1
                                                _ = v105
                                            case 175:
                                                var v106 int = 228
                                                _ = v106
                                            }
                                        }
                                        v1 = (v1 + (v1 * v1))
                                    } else {
                                        var v107 int = 127
                                        if true {
                                            v107 += fi76
                                            var v108 uint64 = 18446744073709551615
                                            v109 := GenS10{F0: 42}
                                            _ = v108
                                            _ = v109
                                        }
                                        var v110 string = "\n\t"
                                        _ = v107
                                        _ = v110
                                    }
                                } else {
                                    v111 := []int{fi101, (-(fi101)), fi76}
                                    var v112 int16 = (v1 & -633)
                                    v113 := GenS11{F0: v1, F1: 4294967295, F2: 2147483647}
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                }
                                v114 := GenS3{F0: 42, F1: 65535, F2: 479}
                                _ = fi101
                                _ = v102
                                _ = v114
                            }
                            var v115 float64 = v2
                            if (true && false) {
                                var v116 uint16 = 0
                                for fi117 := 0; fi117 < 1; fi117++ {
                                    v118 := GenS12{F0: (v4 % (((-(v4))) | 1)), F1: (v1 | 42)}
                                    _ = fi117
                                    _ = v118
                                }
                                _ = v116
                            }
                            _ = v115
                        }
                        fmt.Println(fi76, fi76, v1)
                    }
                    _ = fi76
                }
            }
        }
    } else {
        delete(v3, "\n\t")
        fmt.Println(v2, v2)
        switch -943 {
        case 32767:
            fmt.Println(v4, v4, v4)
        default:
            if false {
                switch "foo bar" {
                case "":
                    var v119 int16 = v1
                    _ = v119
                }
                var v120 float32 = (0.0 + 1e10)
                _ = v120
            } else {
                v121 := GenS13{F0: 2}
                for fi122 := 0; fi122 < 9; fi122++ {
                    var v123 uint16 = 1
                    _ = fi122
                    _ = v123
                }
                var v124 uint32 = 2147483647
                _ = v121
                _ = v124
            }
        }
    }
    var v125 int = 746
    for fi126 := 0; fi126 < 16; fi126++ {
        v2 += (0.0 - (v2 + (v2 * v2)))
        _ = fi126
    }
    v4 = (-(v4))
    v127 := []int{}
    var v128 int8 = -1
    if true {
        if (false || true) {
            v129 := GenS14{F0: 777, F1: 128}
            v130 := GenS11{F0: v1, F1: 65535, F2: -2147483648}
            if ((!false) && false) {
                v128 = v128
                for fi131 := 0; fi131 < 10; fi131++ {
                    switch "a" {
                    case "":
                        if (1 != 2) {
                            var v132 int32 = -482
                            var v133 bool = true
                            _ = v132
                            _ = v133
                        }
                    }
                    var v134 float64 = v2
                    _ = fi131
                    _ = v134
                }
            }
            _ = v129
            _ = v130
        } else {
            fmt.Println(v128, v4)
            var v135 int32 = -2147483648
            _ = v135
        }
        switch (v125 | v125) {
        case 128:
            if false {
                v125 += 0
                for fi136 := 0; fi136 < 16; fi136++ {
                    v1 += -996
                    for fi137 := 0; fi137 < 17; fi137++ {
                        if (!false) {
                            switch fi137 {
                            case 32767:
                                v138 := GenS7{F0: (^(v1)), F1: 762}
                                _ = v138
                            case -9223372036854775808:
                                v139 := GenS2{F0: 32767, F1: 2}
                                _ = v139
                            default:
                                _ = len(v127)
                            }
                            var v140 string = "foo bar"
                            delete(v3, (v140 + "\n\t"))
                            _ = v140
                        } else {
                            if true {
                                _ = len(v3)
                            }
                        }
                        _ = fi137
                    }
                    _ = fi136
                }
            }
        default:
            var v141 uint8 = 42
            _ = v141
        }
    }
    v1 = v1
    if len(v127) > 0 {
        v127[0] = (v125 | v125)
    }
    v4 = -9223372036854775808
    v2 += v2
    v4 = (v4 - 0)
    if len(v127) > 0 {
        v127[0] = (v125 - v125)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v127:", len(v127), v127)
    fmt.Println("v128:", v128)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v125
    _ = v127
    _ = v128
    _ = v2
    _ = v3
    _ = v4
}
