package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 string; F1 int32; F2 uint }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 int8; F2 int }
type GenS11 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS12 struct { F0 bool; F1 int64; F2 bool }
type GenS13 struct { F0 uint16; F1 bool }
type GenS14 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    var v4 float64 = 0.1
    if false {
        if (42 == 65535) {
            v5 := []int{}
            for fi6 := 0; fi6 < 15; fi6++ {
                var v7 uint8 = 1
                _ = fi6
                _ = v7
            }
            v4 = ((0.1 * float64(((3.14159 / 1e10) + (1e10 / (0.0 + (((-2.25 * (-2.25 / 1e10)) + 1e10) / ((((3.14159 / -2.25) * (1e10 * 0.1)) - (((3.14159 / -2.25) / 1.5) / 0.1)) * 3.14159))))))) / 3.14159)
            _ = v5
        }
    } else {
        v8 := []int{}
        if false {
            var v9 float32 = (((3.14159 - -2.25) + -2.25) / 1.0)
            if (127 > -1) {
                v10 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 2}
                fmt.Println(v4, v4)
                _ = v10
            }
            _ = v9
        } else {
            if (1 != 127) {
                var v11 uint16 = 691
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 uint32 = 2
                    _ = fi12
                    _ = v13
                }
                _ = v11
            }
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int32 = 558
            var v16 int64 = 1
            fmt.Println(fi14, fi14)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v8
    }
    v17 := GenS1{F0: -433, F1: 0, F2: 42}
    v18 := GenS1{F0: -691, F1: 42, F2: 2}
    if false {
        if (-2.25 > float32(v4)) {
            v4 += (float64((float32(1e10) * 3.14159)) * 0.1)
        }
        var v19 float32 = (0.0 - 0.0)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            v4 = 3.14159
            _ = fi20
        }
        v21 := GenS2{F0: -102, F1: 65535}
        _ = v21
    }
    fmt.Println(v4)
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 int64 = -9223372036854775808
        v24 := GenS3{F0: 0, F1: (fi22 * fi22)}
        _ = fi22
        _ = v23
        _ = v24
    }
    if (-9223372036854775808 == 128) {
        delete(v2, 65535)
        fmt.Println(v4)
        var v25 int = -440
        _ = v25
    } else {
        fmt.Println(v4)
        if ("a" != "a") {
            var v26 bool = true
            for fi27 := 0; fi27 < 19; fi27++ {
                v4 = (v4 / v4)
                v28 := []int{}
                for fi29 := 0; fi29 < 8; fi29++ {
                    delete(v1, 2)
                    v26 = v26
                    v30 := GenS1{F0: 1, F1: 42, F2: 42}
                    _ = fi29
                    _ = v30
                }
                _ = fi27
                _ = v28
            }
            v4 += v4
            _ = v26
        } else {
            var v31 uint32 = 42
            if (!false) {
                v31 = v31
                var v32 int16 = 699
                if false {
                    fmt.Println(v32, v32)
                }
                _ = v32
            }
            _ = v31
        }
    }
    fmt.Println(v4)
    v33 := GenS3{F0: 2, F1: 4294967295}
    switch ("hello" + "") {
    case "hello":
        v34 := GenS3{F0: 42, F1: -627}
        _ = v34
    case "a":
        var v35 int32 = 954
        _ = v35
    case "\n\t":
        v36 := []int{-2147483648, 2147483647}
        _ = v36
    default:
        v37 := []int{-520, 127, 874}
        _ = v37
    }
    var v38 uint = 870
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 string = (("" + ("" + "foo bar")) + "a")
        switch (fi39 + 32768) {
        case 42:
            _ = len(v2)
        case -2147483648:
            var v41 uint16 = 32767
            _ = v41
        case 2147483647:
            v40 = "a"
        }
        _ = fi39
        _ = v40
    }
    if true {
        var v42 uint64 = 1
        _ = v42
    }
    v43 := GenS4{F0: "a", F1: 180, F2: 65535}
    for fi44 := 0; fi44 < 6; fi44++ {
        fmt.Println(fi44, v38)
        _ = fi44
    }
    fmt.Println(v4, v38)
    if false {
        var v45 int32 = -2147483648
        _ = v45
    } else {
        if false {
            v46 := []int{828}
            fmt.Println(v38, v38)
            v47 := v1[447]
            _ = v46
            _ = v47
        }
        switch 4294967295 {
        case 9223372036854775807:
            var v48 float32 = (1.5 - 1e10)
            _ = v48
        case 42:
            for fi49 := 0; fi49 < 4; fi49++ {
                if false {
                    fmt.Println(v38)
                    fmt.Println(fi49, v4)
                } else {
                    delete(v2, 1)
                    v50 := GenS3{F0: 198, F1: fi49}
                    var v51 int16 = 42
                    _ = v50
                    _ = v51
                }
                v52 := GenS2{F0: (fi49 / ((-2147483648) | 1)), F1: 2147483647}
                fmt.Println(fi49, v4, fi49)
                _ = fi49
                _ = v52
            }
        default:
            if false {
                fmt.Println(v4, v4)
                var v53 int = 9223372036854775807
                _ = v53
            }
        }
        _ = len(v2)
    }
    if (!false) {
        switch 331 {
        case 340:
            var v54 int8 = -128
            _ = v54
        }
        if (18446744073709551615 >= 0) {
            fmt.Println(v38, v38)
            v55 := GenS5{F0: 204, F1: (("hello" + "\n\t") + "a")}
            if true {
                for fi56 := 0; fi56 < 17; fi56++ {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        if (255 < 1) {
                            switch "a" {
                            case "a":
                                if true {
                                    v58 := GenS4{F0: "a", F1: 2147483647, F2: (v38 / (((v38 + 940)) | 1))}
                                    if true {
                                        fmt.Println(fi57)
                                    } else {
                                        var v59 uint32 = 350
                                        _ = v59
                                    }
                                    _ = v58
                                } else {
                                    v4 = v4
                                    if true {
                                        v60 := []int{1}
                                        switch 127 {
                                        case 952:
                                            v38 += (v38 - v38)
                                        case 2147483647:
                                            v4 += ((((v4 + (v4 - float64(0.0))) / v4) - -2.25) - v4)
                                        case -1:
                                            fmt.Println(fi57, v4, fi56)
                                        }
                                        if (fi57 < (fi57 - 32768)) {
                                            switch fi57 {
                                            case 32768:
                                                if (1 >= 1) {
                                                    if true {
                                                        if false {
                                                            var v61 float32 = 1.5
                                                            var v62 float32 = (v61 - -2.25)
                                                            var v63 string = "\n\t"
                                                            fmt.Println(fi56, fi57, fi56)
                                                            _ = v61
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            fmt.Println(fi56)
                                                            v64 := GenS2{F0: fi57, F1: 42}
                                                            _ = v64
                                                        }
                                                        var v65 uint32 = 1
                                                        var v66 uint64 = 42
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                } else {
                                                    fmt.Println(fi56, v4, fi57)
                                                    var v67 int64 = 227
                                                    _ = v67
                                                }
                                            default:
                                                fmt.Println(v38, v38)
                                                v68 := GenS3{F0: 255, F1: 0}
                                                _ = v68
                                            }
                                            if (!(false || (true && (!true)))) {
                                                _ = len(v60)
                                                v69 := GenS2{F0: fi57, F1: 42}
                                                var v70 uint16 = 674
                                                _ = v69
                                                _ = v70
                                            } else {
                                                if true {
                                                    fmt.Println(v4, v4, v38)
                                                    v71 := GenS6{F0: v38}
                                                    if false {
                                                        if (true || (((v4 + (v4 / ((((float64(0.1) + v4) - 1e10) + -2.25) - float64(0.1)))) * 3.14159) >= -2.25)) {
                                                            var v72 uint8 = 114
                                                            var v73 uint64 = 42
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    } else {
                                                        v74 := GenS6{F0: (v38 - (v38 - 255))}
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v60 = append(v60, (fi57 & fi56))
                                                        case "a":
                                                            v75 := v2[2]
                                                            _ = v75
                                                        case "":
                                                            if false {
                                                                if false {
                                                                    v60 = append(v60, (fi57 - (fi56 % ((fi56) | 1))))
                                                                } else {
                                                                    v76 := GenS2{F0: fi56, F1: 2147483647}
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if len(v60) >= 2 {
                                                                            v60 = v60[1:len(v60)]
                                                                        }
                                                                    case "":
                                                                        v77 := GenS6{F0: (^(v38))}
                                                                        v38 += (v38 << ((777) & 63))
                                                                        _ = v77
                                                                    case "hello":
                                                                        v38 = v38
                                                                    default:
                                                                        v38 += 42
                                                                    }
                                                                    v38 -= v38
                                                                    _ = v76
                                                                }
                                                                v78 := GenS1{F0: 9223372036854775807, F1: 2, F2: 65535}
                                                                _ = v78
                                                            }
                                                        default:
                                                            var v79 uint = 255
                                                            fmt.Println(fi57, fi56)
                                                            var v80 uint = (v38 * v79)
                                                            v81 := GenS2{F0: (fi56 & 419), F1: 2147483647}
                                                            fmt.Println(v4)
                                                            switch "a" {
                                                            case "":
                                                                if false {
                                                                    v82 := GenS7{F0: 65535, F1: ""}
                                                                    _ = v82
                                                                }
                                                            case "foo bar":
                                                                var v83 uint64 = 18446744073709551615
                                                                _ = v83
                                                            case "hello":
                                                                v84 := GenS8{F0: "\n\t", F1: 1}
                                                                _ = v84
                                                            }
                                                            if (v4 <= float64(1e10)) {
                                                                v85 := GenS8{F0: "hello", F1: 4294967295}
                                                                if len(v60) >= 2 {
                                                                    v60 = v60[1:len(v60)]
                                                                }
                                                                v79 -= (v79 >> (((v38 / ((v38) | 1))) & 63))
                                                                _ = v85
                                                            } else {
                                                                if false {
                                                                    switch "" {
                                                                    case "":
                                                                        if (127 <= 79) {
                                                                            fmt.Println(v79)
                                                                            fmt.Println(v4, fi56, v79)
                                                                            var v86 bool = true
                                                                            _ = v86
                                                                        }
                                                                        var v87 float32 = (float32(((float64(-2.25) - ((1e10 / v4) + float64(1.5))) + v4)) * (3.14159 - float32(v4)))
                                                                        v38 -= (v79 % ((131) | 1))
                                                                        _ = v87
                                                                    }
                                                                    var v88 int32 = 1
                                                                    v89 := GenS9{F0: (v38 / ((65535) | 1))}
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                v4 -= v4
                                                                var v90 bool = (true || false)
                                                                _ = v90
                                                            }
                                                            if true {
                                                                fmt.Println(v80, fi56, v4)
                                                                var v91 string = "\n\t"
                                                                var v92 int8 = 42
                                                                fmt.Println(fi56, fi57)
                                                                _ = v91
                                                                _ = v92
                                                            }
                                                            _ = v79
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                        _ = v74
                                                    }
                                                    _ = v71
                                                }
                                                if len(v60) >= 2 {
                                                    v60 = v60[1:len(v60)]
                                                }
                                                fmt.Println(v38, v38)
                                                fmt.Println(v4, v38)
                                            }
                                            var v93 uint = 322
                                            _ = v93
                                        } else {
                                            var v94 uint32 = 4294967295
                                            _ = v94
                                        }
                                        switch "hello" {
                                        case "":
                                            var v95 int16 = int16(fi57)
                                            _ = v95
                                        case "\n\t":
                                            v1[4294967295] = false
                                        case "a":
                                            v38 -= v38
                                        default:
                                            if (2147483647 <= 1) {
                                                if false {
                                                    var v96 int64 = -9223372036854775808
                                                    v4 = v4
                                                    _ = v96
                                                }
                                                if false {
                                                    var v97 int8 = 79
                                                    v97 += (v97 % ((v97) | 1))
                                                    var v98 int = -38
                                                    _ = v97
                                                    _ = v98
                                                }
                                                if (68 < 529) {
                                                    if true {
                                                        v4 = v4
                                                        var v99 bool = ((-2.25 + 3.14159) < 0.0)
                                                        var v100 uint8 = 127
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                    if (57 > 127) {
                                                        v101 := GenS2{F0: fi56, F1: 1}
                                                        var v102 uint8 = 128
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v103 := GenS10{F0: 18446744073709551615, F1: -18, F2: (fi56 / ((-985) | 1))}
                                                            _ = v103
                                                        case "":
                                                            v38 -= (^(v38))
                                                        }
                                                    }
                                                    v4 = v4
                                                }
                                            }
                                        }
                                        _ = v60
                                    } else {
                                        v104 := v1[542]
                                        if v104 {
                                            v104 = v104
                                        } else {
                                            fmt.Println(fi57, v38, v38)
                                            if (42 <= 829) {
                                                v105 := GenS10{F0: v38, F1: 1, F2: (-(fi56))}
                                                v106 := []int{fi56, (fi56 ^ (fi57 >> ((913) & 63))), fi56, fi57}
                                                _ = v105
                                                _ = v106
                                            } else {
                                                fmt.Println(v4, fi57)
                                            }
                                            var v107 uint64 = 4294967295
                                            _ = v107
                                        }
                                        if (!(-1 >= int32(fi56))) {
                                            v4 += v4
                                            v104 = false
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                switch "" {
                                                case "hello":
                                                    if (true && (!v104)) {
                                                        delete(v3, "\n\t")
                                                    } else {
                                                        v108 := GenS11{F0: 42, F1: (v38 | v38), F2: 4294967295}
                                                        _ = v108
                                                    }
                                                case "":
                                                    v109 := GenS8{F0: "a", F1: 2147483647}
                                                    _ = v109
                                                case "foo bar":
                                                    var v110 string = "hello"
                                                    _ = v110
                                                default:
                                                    var v111 int = fi56
                                                    _ = v111
                                                }
                                            }
                                            v38 = v38
                                            var v112 uint8 = 2
                                            _ = v112
                                        }
                                        _ = v104
                                    }
                                    var v113 int = 128
                                    _ = v113
                                }
                            }
                            var v114 float64 = float64(0.0)
                            v4 = (float64(float32(((((-2.25 - (v4 + (v4 + 1e10))) * v114) / 3.14159) - v114))) / 1.0)
                            _ = v114
                        }
                        _ = fi57
                    }
                    if false {
                        v38 = (^(v38))
                    } else {
                        switch "hello" {
                        case "\n\t":
                            if false {
                                for fi115 := 0; fi115 < 6; fi115++ {
                                    switch 550 {
                                    case -173:
                                        var v116 int8 = -128
                                        _ = v116
                                    case -511:
                                        switch (fi115 << ((42) & 63)) {
                                        case 128:
                                            fmt.Println(fi115)
                                        case -2147483648:
                                            v117 := GenS2{F0: (fi56 >> ((v38) & 63)), F1: 2147483647}
                                            _ = v117
                                        case 165:
                                            if false {
                                                var v118 uint = (v38 & (v38 | (v38 * v38)))
                                                if false {
                                                    v119 := []int{32768, (fi56 / (((fi56 | 358)) | 1)), 0, (fi56 / (((fi115 % (((^(fi115))) | 1))) | 1))}
                                                    if false {
                                                        var v120 int16 = 42
                                                        var v121 int32 = 1
                                                        if true {
                                                            v121 -= (v121 & (-(v121)))
                                                            v122 := GenS12{F0: false, F1: -9223372036854775808, F2: false}
                                                            switch fi56 {
                                                            case 418:
                                                                v121 += -614
                                                            default:
                                                                v123 := GenS2{F0: fi56, F1: 65535}
                                                                _ = v123
                                                            }
                                                            _ = v122
                                                        }
                                                        _ = v120
                                                        _ = v121
                                                    }
                                                    _ = v119
                                                } else {
                                                    var v124 bool = false
                                                    v4 = (v4 * v4)
                                                    if v124 {
                                                        var v125 float64 = float64((0.1 - 1.5))
                                                        v38 = v118
                                                        v124 = (!(9223372036854775807 == (fi115 - 9223372036854775807)))
                                                        _ = v125
                                                    } else {
                                                        var v126 int32 = 42
                                                        _ = v126
                                                    }
                                                    _ = v124
                                                }
                                                if (42 <= -9223372036854775808) {
                                                    v127 := GenS13{F0: 523, F1: true}
                                                    if true {
                                                        v38 = 42
                                                    }
                                                    _ = v127
                                                }
                                                _ = v118
                                            }
                                        }
                                    }
                                    fmt.Println(fi56, v38)
                                    _ = fi115
                                }
                                v4 = (v4 - (float64(1e10) * 1.5))
                                var v128 string = "a"
                                _ = v128
                            } else {
                                var v129 int = 975
                                v130 := []int{(fi56 * v129), 128, fi56}
                                _ = v129
                                _ = v130
                            }
                        case "":
                            var v131 uint16 = 0
                            _ = v131
                        case "foo bar":
                            var v132 int32 = -206
                            _ = v132
                        }
                        v4 += (v4 + v4)
                        switch 1 {
                        case -991:
                            for fi133 := 0; fi133 < 14; fi133++ {
                                v134 := []int{fi133, fi56, fi133}
                                _ = fi133
                                _ = v134
                            }
                        case -2147483648:
                            v135 := GenS14{F0: 255}
                            _ = v135
                        }
                    }
                    _ = fi56
                }
                v4 -= -2.25
            }
            _ = v55
        } else {
            _ = len(v3)
            if true {
                if true {
                    for fi136 := 0; fi136 < 12; fi136++ {
                        var v137 float32 = 1e10
                        _ = fi136
                        _ = v137
                    }
                    var v138 int16 = -32768
                    for fi139 := 0; fi139 < 4; fi139++ {
                        fmt.Println(fi139, v38)
                        v138 += (v138 + (v138 << ((v38) & 63)))
                        _ = fi139
                    }
                    _ = v138
                } else {
                    v140 := []int{0, 32768, 0, 42}
                    v4 -= ((v4 + (v4 * float64((3.14159 - 1.5)))) * 0.1)
                    switch "foo bar" {
                    case "a":
                        var v141 int = 2147483647
                        _ = v141
                    case "hello":
                        delete(v3, "foo bar")
                    }
                    _ = v140
                }
                v38 -= 1
            } else {
                for fi142 := 0; fi142 < 1; fi142++ {
                    v2[1] = true
                    fmt.Println(v38)
                    var v143 int64 = -1
                    _ = fi142
                    _ = v143
                }
                v38 += v38
                v38 = 604
            }
            v4 -= (v4 * ((3.14159 * -2.25) * 1e10))
        }
        if true {
            var v144 uint64 = 481
            fmt.Println(v4, v4, v144)
            _ = v144
        }
    } else {
        var v145 int16 = 1
        _ = v145
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v43
}
