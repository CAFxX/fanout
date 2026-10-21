package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0(p0 int16, p1 int16, p2 uint) uint8 {
    var v5 uint8 = 234
    var v6 string = "foo bar"
    switch "a" {
    case "foo bar":
        v5 = v5
    case "a":
        switch -551 {
        case 4294967295:
            p1 = (^(p0))
        case 42:
            for fi7 := 0; fi7 < 5; fi7++ {
                var v8 float32 = (3.14159 - 3.14159)
                var v9 bool = false
                _ = fi7
                _ = v8
                _ = v9
            }
        case 128:
            v10 := []int{-1, 0}
            _ = v10
        }
    default:
        fmt.Println(v5, p1, v5)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint8, p1 uint, p2 int32) bool {
    v11 := GenS1{F0: 1, F1: 2147483648}
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return false
}

func main() {
    var v1 uint = 108
    var v2 int8 = 100
    var v3 float64 = 3.14159
    var v4 float32 = 0.0
    switch ((("foo bar" + "hello") + "foo bar") + ("a" + "foo bar")) {
    case "\n\t":
        var v12 string = "\n\t"
        _ = v12
    default:
        v2 = v2
    }
    var v13 uint8 = 61
    switch "" {
    case "\n\t":
        fmt.Println(v13)
    default:
        var v14 uint8 = v13
        _ = v14
    }
    fmt.Println(v1, v4)
    v15 := GenS1{F0: -963, F1: uint32(v1)}
    if true {
        v16 := []int{}
        _ = v16
    } else {
        var v17 bool = (!(100 < v2))
        for fi18 := 0; fi18 < 11; fi18++ {
            for fi19 := 0; fi19 < 10; fi19++ {
                if (!v17) {
                    v1 = 367
                    switch "foo bar" {
                    case "foo bar":
                        if (1 != 355) {
                            if (!(42 > 1)) {
                                fmt.Println(fi18, v1)
                                if (!(((fi18 ^ (^(fi19))) >= 9223372036854775807) || (0 < 2))) {
                                    var v20 bool = v17
                                    if v20 {
                                        var v21 uint16 = uint16(v13)
                                        v22 := GenS2{F0: 42}
                                        if v17 {
                                            if v20 {
                                                v23 := []int{32768}
                                                _ = v23
                                            } else {
                                                v2 = (^(v2))
                                                var v24 bool = (((-2.25 + v3) * (((float64(1e10) - (((v3 * v3) * v3) + -2.25)) - ((v3 + ((1e10 * (v3 - v3)) - v3)) / v3)) / v3)) > (v3 + (float64(v4) * v3)))
                                                switch "" {
                                                case "foo bar":
                                                    v25 := []int{fi19, int(v2)}
                                                    if false {
                                                        v1 += (v1 + (^(v1)))
                                                        v13 -= (v13 | (^(v13)))
                                                        switch 32768 {
                                                        case 1:
                                                            v21 -= 938
                                                        case 2147483647:
                                                            if (-2.25 < v3) {
                                                                if len(v25) > 0 {
                                                                    v25[0] = fi19
                                                                }
                                                                if false {
                                                                    var v26 float64 = 1e10
                                                                    fmt.Println(fi18)
                                                                    _ = v26
                                                                }
                                                                switch (fi19 * fi18) {
                                                                case 4294967295:
                                                                    var v27 uint8 = (v13 | v13)
                                                                    _ = v27
                                                                }
                                                            } else {
                                                                if v24 {
                                                                    var v28 uint = v1
                                                                    if true {
                                                                        v29 := GenS2{F0: -32768}
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            if len(v25) >= 2 {
                                                                                v25 = v25[1:len(v25)]
                                                                            }
                                                                        case "hello":
                                                                            var v30 string = "\n\t"
                                                                            _ = v30
                                                                        }
                                                                        switch fi19 {
                                                                        case 2147483647:
                                                                            v31 := GenS1{F0: 9223372036854775807, F1: 1}
                                                                            var v32 int64 = -1
                                                                            _ = v31
                                                                            _ = v32
                                                                        case -842:
                                                                            v33 := GenS3{F0: v21, F1: v13, F2: 2147483648}
                                                                            _ = v33
                                                                        }
                                                                        _ = v29
                                                                    } else {
                                                                        v13 += v13
                                                                        var v34 int32 = 1
                                                                        _ = v34
                                                                    }
                                                                    fmt.Println(v2, v2, fi18)
                                                                    _ = v28
                                                                }
                                                                v4 -= (v4 + float32(float64(((float32(v3) / v4) + 0.1))))
                                                                fmt.Println(fi18)
                                                                var v35 int16 = -32768
                                                                fmt.Println(v3, v35, v20)
                                                                _ = v35
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(fi18, v1)
                                                    }
                                                    v1 += uint(v13)
                                                    _ = v25
                                                case "a":
                                                    v36 := GenS3{F0: v21, F1: 127, F2: 4294967295}
                                                    _ = v36
                                                }
                                                var v37 int64 = 0
                                                _ = v24
                                                _ = v37
                                            }
                                            var v38 uint = uint(v21)
                                            switch (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "foo bar") {
                                            case "a":
                                                switch int(v1) {
                                                case 127:
                                                    switch ("foo bar" + "\n\t") {
                                                    case "hello":
                                                        v39 := GenS2{F0: -32768}
                                                        _ = v39
                                                    case "a":
                                                        v38 = (^(v38))
                                                    }
                                                }
                                            case "\n\t":
                                                v38 += v1
                                            default:
                                                if (-1 < 0) {
                                                    v17 = false
                                                    var v40 uint32 = 2
                                                    fmt.Println(v20)
                                                    _ = v40
                                                }
                                            }
                                            if v17 {
                                                if v17 {
                                                    v41 := []int{}
                                                    _ = v41
                                                }
                                                if false {
                                                    if true {
                                                        v42 := GenS1{F0: -751, F1: 1}
                                                        _ = v42
                                                    } else {
                                                        fmt.Println(v20, v3, v38)
                                                        switch ("\n\t" + (("a" + (("" + ("foo bar" + "\n\t")) + "a")) + (("hello" + "\n\t") + ""))) {
                                                        case "a":
                                                            v1 = 18446744073709551615
                                                            v43 := GenS4{F0: genFn0(1, 0, v38)}
                                                            _ = v43
                                                        case "":
                                                            v44 := GenS5{F0: ("" + "a")}
                                                            _ = v44
                                                        case "\n\t":
                                                            v21 -= (v21 << ((v38) & 63))
                                                        }
                                                        switch "" {
                                                        case "foo bar":
                                                            var v45 float64 = (((v3 - v3) * 0.1) + v3)
                                                            _ = v45
                                                        default:
                                                            v46 := GenS4{F0: 189}
                                                            _ = v46
                                                        }
                                                        v47 := GenS6{F0: v13}
                                                        _ = v47
                                                    }
                                                    if v17 {
                                                        fmt.Println(v21, fi18)
                                                        v48 := GenS4{F0: (^(v13))}
                                                        v49 := []int{128, 1}
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                    var v50 float64 = -2.25
                                                    v38 += v1
                                                    v51 := []int{fi18, (-(fi19)), 1}
                                                    var v52 int = fi19
                                                    fmt.Println(v21, fi19, v50)
                                                    _ = v50
                                                    _ = v51
                                                    _ = v52
                                                }
                                                v53 := []int{42}
                                                _ = v53
                                            } else {
                                                var v54 uint8 = 32
                                                if false {
                                                    if (1.5 >= (v4 - v4)) {
                                                        v21 -= (v21 | (v21 % ((108) | 1)))
                                                        var v55 float64 = ((v3 / (1.5 / v3)) / 3.14159)
                                                        _ = v55
                                                    } else {
                                                        v1 = (v38 % ((v38) | 1))
                                                    }
                                                    v20 = v20
                                                    switch "\n\t" {
                                                    case "a":
                                                        v56 := []int{(fi18 >> (((v38 ^ 760)) & 63)), fi18, fi19}
                                                        _ = v56
                                                    case "":
                                                        if ((v2 >> ((v1) & 63)) == (v2 & v2)) {
                                                            v57 := GenS4{F0: v54}
                                                            var v58 int32 = -175
                                                            v59 := GenS3{F0: (v21 - (v21 / ((v21) | 1))), F1: uint8(v21), F2: 2}
                                                            switch ("\n\t" + "\n\t") {
                                                            case "\n\t":
                                                                var v60 uint64 = 42
                                                                _ = v60
                                                            case "":
                                                                switch "a" {
                                                                case "a":
                                                                    if v20 {
                                                                        if ((v38 != 65535) && v20) {
                                                                            var v61 int16 = -1
                                                                            v62 := GenS2{F0: int16(fi19)}
                                                                            v58 = v58
                                                                            _ = v61
                                                                            _ = v62
                                                                        } else {
                                                                            if (255 != (^(v54))) {
                                                                                v63 := GenS2{F0: 32767}
                                                                                v58 = (v58 - 2147483647)
                                                                                _ = v63
                                                                            } else {
                                                                                var v64 float32 = v4
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    v65 := GenS3{F0: 161, F1: (v54 + 250), F2: 42}
                                                                                    _ = v65
                                                                                }
                                                                                if false {
                                                                                    v38 += uint(v13)
                                                                                    fmt.Println(v58, v38, v2)
                                                                                }
                                                                                _ = v64
                                                                            }
                                                                            var v66 uint8 = genFn0(67, 0, v1)
                                                                            v58 += v58
                                                                            _ = v66
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            if (1 > fi19) {
                                                                                var v67 uint = (^(v38))
                                                                                v68 := GenS2{F0: 32767}
                                                                                _ = v67
                                                                                _ = v68
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v69 uint = 203
                                                                        var v70 float64 = (1.5 - v3)
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                }
                                                            case "a":
                                                                if true {
                                                                    var v71 uint16 = (v21 ^ 65535)
                                                                    fmt.Println(v58)
                                                                    v4 = float32(v3)
                                                                    _ = v71
                                                                } else {
                                                                    if v17 {
                                                                        var v72 int32 = (-(v58))
                                                                        if (v38 >= v1) {
                                                                            v2 = (v2 & -1)
                                                                            if v17 {
                                                                                v58 += 0
                                                                                var v73 int64 = 645
                                                                                _ = v73
                                                                            }
                                                                            v21 += (v21 ^ uint16(v13))
                                                                        } else {
                                                                            var v74 uint64 = 0
                                                                            _ = v74
                                                                        }
                                                                        v75 := []int{fi18, (fi19 ^ (fi19 - fi18))}
                                                                        _ = v72
                                                                        _ = v75
                                                                    } else {
                                                                        v76 := GenS1{F0: int64(fi18), F1: 4294967295}
                                                                        var v77 float64 = (v3 * (v3 + v3))
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                }
                                                            default:
                                                                var v78 uint8 = 1
                                                                _ = v78
                                                            }
                                                            var v79 float64 = 0.0
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                            _ = v79
                                                        } else {
                                                            var v80 float64 = v3
                                                            if v20 {
                                                                v81 := GenS7{F0: -827, F1: 2147483648}
                                                                v82 := []int{(fi19 + -2147483648)}
                                                                var v83 float32 = (v4 + (((-2.25 + float32(v80)) + -2.25) - (v4 * v4)))
                                                                _ = v81
                                                                _ = v82
                                                                _ = v83
                                                            }
                                                            if v17 {
                                                                v3 += v3
                                                                var v84 int = fi19
                                                                _ = v84
                                                            } else {
                                                                v85 := []int{(-(fi18)), (fi18 | int(v21)), 32768, fi18}
                                                                var v86 uint32 = 2
                                                                if v20 {
                                                                    fmt.Println(fi18)
                                                                    v87 := GenS8{F0: 0}
                                                                    if len(v85) > 0 {
                                                                        v85[0] = (^(fi19))
                                                                    }
                                                                    _ = v87
                                                                }
                                                                switch 0 {
                                                                case 127:
                                                                    v3 = float64(v4)
                                                                case 0:
                                                                    var v88 int64 = 743
                                                                    _ = v88
                                                                case 76:
                                                                    v89 := GenS3{F0: v21, F1: (v13 >> ((2) & 63)), F2: (v86 ^ v86)}
                                                                    _ = v89
                                                                }
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            _ = v80
                                                        }
                                                    case "foo bar":
                                                        v90 := []int{fi19, 0}
                                                        _ = v90
                                                    }
                                                    var v91 bool = (-2.25 == ((0.1 / float64((1e10 + ((-2.25 / v4) * 0.1)))) - ((0.1 + (v3 - float64(-2.25))) / v3)))
                                                    _ = v91
                                                }
                                                _ = v54
                                            }
                                            if v20 {
                                                v92 := []int{}
                                                if len(v92) >= 2 {
                                                    v92 = v92[1:len(v92)]
                                                }
                                                _ = v92
                                            } else {
                                                var v93 int64 = 1
                                                switch ((("\n\t" + ("\n\t" + ("hello" + (("foo bar" + ("hello" + ("a" + (("a" + ("a" + "foo bar")) + ("\n\t" + ""))))) + ("" + "foo bar"))))) + "a") + "\n\t") {
                                                case "hello":
                                                    v94 := GenS8{F0: 42}
                                                    _ = v94
                                                default:
                                                    v38 += (v1 | v38)
                                                }
                                                v20 = false
                                                _ = v93
                                            }
                                            _ = v38
                                        }
                                        if true {
                                            switch fi19 {
                                            case 127:
                                                v95 := []int{(fi18 ^ (-(fi18)))}
                                                _ = v95
                                            case -2147483648:
                                                fmt.Println(v1)
                                            case -430:
                                                v96 := GenS5{F0: "a"}
                                                _ = v96
                                            default:
                                                v97 := []int{fi19}
                                                _ = v97
                                            }
                                            fmt.Println(fi19, v1)
                                            v21 = 0
                                            v98 := GenS1{F0: 1, F1: 2}
                                            fmt.Println(v1, v1)
                                            _ = v98
                                        }
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v99 int16 = 32767
                                        v100 := GenS3{F0: 1, F1: (v13 ^ 2), F2: 994}
                                        fmt.Println(v2, fi18)
                                        _ = v100
                                        _ = v99
                                    }
                                    switch 533 {
                                    case 989:
                                        v101 := GenS9{F0: v13, F1: genFn0(0, -601, (v1 >> ((v1) & 63)))}
                                        _ = v101
                                    case 4294967295:
                                        v102 := []int{(fi19 << ((42) & 63)), int(v2)}
                                        _ = v102
                                    default:
                                        fmt.Println(v3, v20)
                                    }
                                    _ = v20
                                }
                                v13 += 226
                            }
                            var v103 int32 = -935
                            fmt.Println(fi19, v2, v1)
                            _ = v103
                        } else {
                            v104 := []int{fi18}
                            _ = v104
                        }
                    case "hello":
                        v105 := GenS3{F0: 2, F1: (v13 / ((2) | 1)), F2: 2}
                        _ = v105
                    case "\n\t":
                        v1 -= v1
                    default:
                        v106 := []int{fi19, fi18}
                        _ = v106
                    }
                    if (((v3 + v3) / v3) > (((v3 / (float64(v4) + 0.1)) - v3) / float64(v4))) {
                        var v107 string = "foo bar"
                        var v108 int16 = -32768
                        _ = v107
                        _ = v108
                    } else {
                        v13 = 2
                        var v109 string = ("foo bar" + "a")
                        v110 := []int{(fi19 * fi19), fi18}
                        _ = v109
                        _ = v110
                    }
                }
                var v111 int16 = 42
                _ = fi19
                _ = v111
            }
            _ = fi18
        }
        _ = v17
    }
    v2 = v2
    if (42 > 907) {
        v1 = (v1 / ((v1) | 1))
        v112 := GenS4{F0: 1}
        _ = v112
    } else {
        for fi113 := 0; fi113 < 10; fi113++ {
            v114 := []int{(fi113 << (((v1 | v1)) & 63)), 642}
            fmt.Println(v3)
            _ = fi113
            _ = v114
        }
        if false {
            switch -598 {
            case -2147483648:
                if true {
                    for fi115 := 0; fi115 < 14; fi115++ {
                        if (0 <= 32767) {
                            if (fi115 < fi115) {
                                v3 = v3
                            }
                            var v116 int32 = -960
                            v1 += 0
                            _ = v116
                        } else {
                            v4 += (-2.25 - v4)
                            var v117 uint8 = 159
                            _ = v117
                        }
                        if true {
                            for fi118 := 0; fi118 < 19; fi118++ {
                                v1 = v1
                                var v119 int16 = -394
                                var v120 int16 = 1000
                                _ = fi118
                                _ = v119
                                _ = v120
                            }
                            v1 = (v1 & (v1 / ((v1) | 1)))
                            var v121 int = (fi115 - (fi115 / ((-570) | 1)))
                            _ = v121
                        } else {
                            v4 += v4
                            if (fi115 == fi115) {
                                for fi122 := 0; fi122 < 7; fi122++ {
                                    var v123 string = ("hello" + "")
                                    v3 = 3.14159
                                    _ = fi122
                                    _ = v123
                                }
                            } else {
                                fmt.Println(v1, v2, v2)
                                for fi124 := 0; fi124 < 4; fi124++ {
                                    v125 := []int{fi115, fi124, fi124}
                                    v125 = append(v125, int(v13))
                                    switch "hello" {
                                    case "a":
                                        var v126 uint64 = 18446744073709551615
                                        _ = v126
                                    case "":
                                        fmt.Println(v1)
                                        v125 = append(v125, fi124)
                                        switch "hello" {
                                        case "\n\t":
                                            if (1 != 0) {
                                                var v127 int16 = -262
                                                if (true || ((float64((v4 + v4)) + v3) < (1.5 * (v3 * v3)))) {
                                                    fmt.Println(v3)
                                                    v128 := GenS9{F0: v13, F1: (v13 >> ((2) & 63))}
                                                    _ = v128
                                                } else {
                                                    var v129 int16 = 1
                                                    v129 = 440
                                                    fmt.Println(fi115)
                                                    _ = v129
                                                }
                                                _ = v127
                                            } else {
                                                if false {
                                                    var v130 float32 = (v4 / -2.25)
                                                    v4 += 1.5
                                                    v13 = 42
                                                    _ = v130
                                                }
                                                if len(v125) >= 2 {
                                                    v125 = v125[1:len(v125)]
                                                }
                                                v131 := GenS10{F0: -215}
                                                _ = v131
                                            }
                                        case "":
                                            if len(v125) >= 2 {
                                                v125 = v125[1:len(v125)]
                                            }
                                        case "foo bar":
                                            v132 := GenS7{F0: -353, F1: uint32(v1)}
                                            _ = v132
                                        }
                                    }
                                    _ = fi124
                                    _ = v125
                                }
                            }
                        }
                        if false {
                            v133 := GenS3{F0: uint16(v13), F1: v13, F2: uint32(v13)}
                            _ = v133
                        }
                        _ = fi115
                    }
                    fmt.Println(v1, v2)
                } else {
                    v134 := []int{32768}
                    if false {
                        _ = len(v134)
                        if len(v134) >= 2 {
                            v134 = v134[1:len(v134)]
                        }
                    } else {
                        v13 = (^(v13))
                        v4 -= 0.1
                        v135 := GenS3{F0: 0, F1: v13, F2: 0}
                        _ = v135
                    }
                    _ = v134
                }
            case -492:
                if false {
                    var v136 uint = v1
                    if false {
                        v1 = v136
                    }
                    fmt.Println(v2)
                    _ = v136
                }
            default:
                if (-32768 == 800) {
                    v137 := GenS8{F0: 3}
                    v1 = 65535
                    for fi138 := 0; fi138 < 10; fi138++ {
                        var v139 int = 0
                        _ = fi138
                        _ = v139
                    }
                    _ = v137
                }
            }
        }
        var v140 int8 = (v2 + (v2 >> ((160) & 63)))
        _ = v140
    }
    if true {
        v141 := GenS11{F0: 5, F1: "hello", F2: 18446744073709551615}
        _ = v141
    } else {
        var v142 int32 = -892
        _ = v142
    }
    fmt.Println(v1, v4, v1)
    var v143 uint8 = (v13 + v13)
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v143
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
