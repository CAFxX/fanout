package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 int; F1 uint64; F2 int8 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 bool; F1 int; F2 bool }
type GenS8 struct { F0 int; F1 uint16; F2 int }
func genFn0(p0 int16, p1 uint16) int64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(p0, p0, p0)
        _ = fi4
    }
    p1 = 42
    fmt.Println(p1, p0)
    var v5 uint = 65535
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1() string {
    var v6 float32 = (0.1 / (float32((3.14159 * 1.5)) - 3.14159))
    if false {
        v6 -= v6
        switch 4294967295 {
        case 2147483647:
            var v7 int8 = 96
            _ = v7
        case 1:
            fmt.Println(v6)
        }
        var v8 uint = 255
        _ = v8
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            switch ((("a" + ("\n\t" + "hello")) + "foo bar") + ("" + ("hello" + "\n\t"))) {
            case "hello":
                if (!((fi9 + 128) >= (-(fi9)))) {
                    if (-128 < 127) {
                        var v10 int32 = 491
                        _ = v10
                    }
                    v6 -= (0.1 * float32(3.14159))
                    fmt.Println(v6, v6)
                } else {
                    var v11 uint16 = 243
                    var v12 int8 = -1
                    _ = v11
                    _ = v12
                }
            }
            _ = fi9
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 float32 = (float32(3.14159) + -2.25)
            v15 := []int{}
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: 127, F1: 42, F2: 75}
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := []int{fi17}
        _ = fi17
        _ = v18
    }
    _ = v6
    return "a"
}

func main() {
    var v1 int = 1
    v2 := []uint32{2, 874, 26, 2147483648}
    var v3 uint32 = 0
    var v19 int8 = 100
    if false {
        var v20 bool = true
        v2 = append(v2, 2147483647)
        _ = v20
    } else {
        for fi21 := 0; fi21 < 19; fi21++ {
            switch fi21 {
            case 9223372036854775807:
                switch v1 {
                case 4294967295:
                    if (-2.25 > 0.0) {
                        v3 = (^(v3))
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 int16 = 32767
                            var v24 int = fi21
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = len(v2)
                    }
                case -501:
                    var v25 int64 = 198
                    _ = v25
                case -334:
                    switch (^(fi21)) {
                    case -784:
                        v3 += v3
                    default:
                        v3 -= v3
                    }
                default:
                    if false {
                        if false {
                            for fi26 := 0; fi26 < 8; fi26++ {
                                fmt.Println(fi26, v19, v19)
                                var v27 int = fi26
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = len(v2)
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi28 := 0; fi28 < 2; fi28++ {
                                v3 -= (^(v3))
                                if (true || ("a" == "a")) {
                                    fmt.Println(fi28, fi21, fi28)
                                    var v29 uint64 = 42
                                    var v30 bool = (v3 <= 270)
                                    if (255 != 5) {
                                        fmt.Println(v1, v1)
                                        fmt.Println(v19, fi28, v1)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    if (((v29 | (^(v29))) != v29) || true) {
                                        if (2147483647 == 2) {
                                            v31 := GenS1{F0: 168, F1: 151, F2: -111}
                                            _ = v31
                                        } else {
                                            v29 = 2
                                            if (v30 && (v30 || v30)) {
                                                v2 = append(v2, v3)
                                                var v32 uint8 = 102
                                                var v33 float64 = ((0.0 + 0.0) * (1.5 + -2.25))
                                                _ = v32
                                                _ = v33
                                            }
                                            v34 := GenS1{F0: 42, F1: 255, F2: v19}
                                            _ = v34
                                        }
                                    } else {
                                        var v35 uint8 = 0
                                        _ = v35
                                    }
                                    v2 = append(v2, v3)
                                    _ = v29
                                    _ = v30
                                }
                                _ = fi28
                            }
                        }
                    }
                }
            case 509:
                var v36 uint8 = 42
                _ = v36
            default:
                for fi37 := 0; fi37 < 8; fi37++ {
                    var v38 float32 = 0.1
                    _ = fi37
                    _ = v38
                }
            }
            _ = fi21
        }
        if false {
            v39 := GenS1{F0: 0, F1: 1, F2: 1}
            for fi40 := 0; fi40 < 3; fi40++ {
                if (0 < 1) {
                    v3 = v3
                    v41 := GenS1{F0: 2, F1: 128, F2: (v19 << ((4294967295) & 63))}
                    _ = v41
                }
                for fi42 := 0; fi42 < 8; fi42++ {
                    _ = len(v2)
                    _ = len(v2)
                    fmt.Println(fi42, v19)
                    var v43 string = "hello"
                    if true {
                        var v44 uint32 = v3
                        var v45 uint16 = 2
                        switch v1 {
                        case 688:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 42:
                            v46 := GenS1{F0: 1, F1: 127, F2: (v19 >> ((4294967295) & 63))}
                            _ = v46
                        default:
                            v47 := GenS1{F0: 0, F1: 194, F2: int8(v3)}
                            v3 = (^(v3))
                            var v48 uint = 786
                            _ = v47
                            _ = v48
                        }
                        _ = v44
                        _ = v45
                    } else {
                        var v49 bool = (!(2147483647 <= 42))
                        if ((32767 > 42) || ((v43 + ((v43 + v43) + v43)) <= v43)) {
                            var v50 uint8 = 127
                            var v51 bool = v49
                            v3 -= v3
                            if v49 {
                                v51 = true
                            }
                            _ = v50
                            _ = v51
                        } else {
                            v1 -= (fi40 | 127)
                            switch v43 {
                            case "":
                                var v52 int16 = 321
                                _ = v52
                            case "\n\t":
                                if ((v19 / ((v19) | 1)) >= (v19 + v19)) {
                                    switch v43 {
                                    case "\n\t":
                                        v53 := GenS1{F0: 128, F1: 255, F2: -128}
                                        _ = v53
                                    case "hello":
                                        v2 = append(v2, 429)
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v3, v1)
                                        case "hello":
                                            v43 = "hello"
                                        case "a":
                                            var v54 uint = 42
                                            _ = v54
                                        default:
                                            var v55 bool = (-480 != 42)
                                            _ = v55
                                        }
                                        if false {
                                            v1 = (fi42 ^ (fi40 & v1))
                                            var v56 uint8 = 1
                                            _ = len(v2)
                                            _ = v56
                                        } else {
                                            v49 = (!v49)
                                        }
                                        if v49 {
                                            var v57 float64 = ((0.0 + (-2.25 - 1.5)) * 3.14159)
                                            v58 := GenS2{F0: 18446744073709551615}
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v59 := GenS1{F0: 1, F1: 255, F2: (v19 << ((4294967295) & 63))}
                                            _ = v59
                                        }
                                        v1 += 32767
                                    }
                                    var v60 uint64 = 42
                                    var v61 uint64 = 1
                                    _ = v60
                                    _ = v61
                                }
                            case "a":
                                _ = len(v2)
                            }
                        }
                        _ = v49
                    }
                    _ = fi42
                    _ = v43
                }
                _ = fi40
            }
            fmt.Println(v3)
            _ = v39
        } else {
            v1 = -643
            var v62 bool = (false || true)
            switch "hello" {
            case "\n\t":
                switch v1 {
                case -1:
                    switch 127 {
                    case 817:
                        v19 += (v19 / ((-128) | 1))
                    }
                case 42:
                    if len(v2) > 0 {
                        v2[0] = (v3 + v3)
                    }
                case 32768:
                    var v63 int = v1
                    _ = v63
                }
            }
            _ = v62
        }
    }
    fmt.Println(v1, v1, v19)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((^(v3)) > 0) {
        var v64 uint16 = 897
        _ = v64
    }
    v65 := GenS2{F0: 255}
    for fi66 := 0; fi66 < 18; fi66++ {
        if (false && true) {
            v19 = v19
        } else {
            if (!(42 >= 42)) {
                var v67 string = "\n\t"
                _ = v67
            } else {
                var v68 int16 = 42
                if false {
                    var v69 bool = (646 > 0)
                    if ((v3 / ((2) | 1)) <= v3) {
                        var v70 float32 = 3.14159
                        var v71 uint32 = v3
                        _ = v70
                        _ = v71
                    }
                    if v69 {
                        var v72 uint16 = 1
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v72
                    } else {
                        if (593 <= 530) {
                            if false {
                                for fi73 := 0; fi73 < 13; fi73++ {
                                    if (!v69) {
                                        var v74 int16 = (v68 << ((18446744073709551615) & 63))
                                        v68 -= -32768
                                        v2 = append(v2, (^(v3)))
                                        _ = v74
                                    } else {
                                        fmt.Println(fi73)
                                        var v75 uint16 = 255
                                        if false {
                                            var v76 int64 = -1
                                            _ = v76
                                        }
                                        switch 32768 {
                                        case -142:
                                            switch fi66 {
                                            case 128:
                                                v3 = v3
                                            default:
                                                v19 = v19
                                                var v77 int64 = genFn0(v68, 702)
                                                _ = v77
                                            }
                                        case 42:
                                            var v78 uint64 = 16
                                            _ = v78
                                        case 2147483647:
                                            fmt.Println(v75, v1)
                                        }
                                        var v79 float32 = float32((0.0 - float64(((((3.14159 - 0.0) / -2.25) - 1.5) * (0.0 / 1.5)))))
                                        v80 := GenS1{F0: 255, F1: 255, F2: (v19 % (((^(v19))) | 1))}
                                        var v81 int = v1
                                        _ = v75
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                    fmt.Println(fi73, v69, fi66)
                                    _ = len(v2)
                                    v19 -= v19
                                    _ = fi73
                                }
                                v82 := GenS2{F0: 2}
                                if (false || (2 < 664)) {
                                    var v83 uint32 = (v3 - 221)
                                    fmt.Println(v69)
                                    _ = v83
                                } else {
                                    var v84 float32 = 0.0
                                    for fi85 := 0; fi85 < 2; fi85++ {
                                        switch (v1 + (v1 << ((42) & 63))) {
                                        case -690:
                                            v84 += float32(1.5)
                                        case 128:
                                            v1 += fi66
                                        }
                                        v86 := GenS3{F0: (v68 ^ v68), F1: fi85, F2: (^(v68))}
                                        _ = fi85
                                        _ = v86
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        switch (fi66 + 2147483647) {
                                        case -20:
                                            var v87 float64 = float64(((v84 + 0.1) * v84))
                                            _ = v87
                                        case 127:
                                            if (0 < 785) {
                                                v88 := GenS4{F0: "", F1: (fi66 ^ (fi66 & 774))}
                                                if v69 {
                                                    var v89 float32 = v84
                                                    var v90 uint64 = 42
                                                    v91 := GenS4{F0: "\n\t", F1: (fi66 | -159)}
                                                    _ = v89
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    var v92 uint16 = 32767
                                                    var v93 string = "\n\t"
                                                    for fi94 := 0; fi94 < 12; fi94++ {
                                                        v1 += v1
                                                        fmt.Println(v84)
                                                        _ = fi94
                                                    }
                                                    _ = v92
                                                    _ = v93
                                                }
                                                v95 := GenS2{F0: 2}
                                                _ = v88
                                                _ = v95
                                            } else {
                                                var v96 int32 = -231
                                                if len(v2) > 0 {
                                                    v2[0] = v3
                                                }
                                                _ = v96
                                            }
                                        }
                                    }
                                    _ = v84
                                }
                                _ = v82
                            }
                            fmt.Println(v68, v19, v3)
                            v97 := GenS5{F0: -501, F1: 2, F2: (v19 >> ((2) & 63))}
                            _ = v97
                        } else {
                            v98 := GenS5{F0: fi66, F1: 1, F2: int8(v3)}
                            for fi99 := 0; fi99 < 1; fi99++ {
                                if (v68 < (v68 | v68)) {
                                    var v100 uint8 = 61
                                    v69 = (v69 || ((v68 % ((-151) | 1)) > v68))
                                    fmt.Println(v19, v68, v69)
                                    _ = v100
                                } else {
                                    if v69 {
                                        var v101 uint = 255
                                        _ = v101
                                    }
                                }
                                if (v3 == 2147483648) {
                                    v102 := GenS4{F0: "foo bar", F1: fi99}
                                    _ = v102
                                }
                                _ = fi99
                            }
                            if (!v69) {
                                var v103 float64 = ((-2.25 + 0.0) * (3.14159 / (((0.1 + 0.0) + 1.5) + -2.25)))
                                if v69 {
                                    var v104 uint8 = 127
                                    v105 := GenS5{F0: 127, F1: 1, F2: (v19 | v19)}
                                    v2 = append(v2, v3)
                                    _ = v104
                                    _ = v105
                                } else {
                                    if v69 {
                                        if (v69 && (!(0.0 == 3.14159))) {
                                            v3 = v3
                                            if (!false) {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch v1 {
                                                case 128:
                                                    switch ("a" + "a") {
                                                    case "a":
                                                        var v106 uint = 1
                                                        _ = v106
                                                    case "":
                                                        for fi107 := 0; fi107 < 6; fi107++ {
                                                            if len(v2) > 0 {
                                                                v2[0] = v3
                                                            }
                                                            var v108 uint8 = 0
                                                            _ = fi107
                                                            _ = v108
                                                        }
                                                    case "hello":
                                                        switch ("a" + "foo bar") {
                                                        case "a":
                                                            fmt.Println(v68)
                                                        case "hello":
                                                            if (v19 == v19) {
                                                                v2 = append(v2, 1)
                                                                fmt.Println(v103, v1)
                                                                var v109 int16 = -1
                                                                _ = v109
                                                            } else {
                                                                if true {
                                                                    v110 := GenS4{F0: "hello", F1: fi66}
                                                                    var v111 int = 127
                                                                    v112 := GenS1{F0: 1, F1: 127, F2: (v19 / ((v19) | 1))}
                                                                    _ = v110
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v113 := GenS1{F0: 2, F1: 127, F2: 100}
                                                            _ = v113
                                                        default:
                                                            v114 := GenS4{F0: ("foo bar" + "a"), F1: (fi66 | v1)}
                                                            _ = v114
                                                        }
                                                    default:
                                                        var v115 uint32 = (v3 & (v3 + 65535))
                                                        _ = v115
                                                    }
                                                default:
                                                    var v116 int16 = v68
                                                    _ = v116
                                                }
                                            }
                                        }
                                        fmt.Println(fi66, v69, v3)
                                    } else {
                                        var v117 int = fi66
                                        _ = v117
                                    }
                                }
                                fmt.Println(v1, v103)
                                _ = v103
                            }
                            _ = v98
                        }
                        if ((v19 / ((52) | 1)) < (v19 | (v19 * -108))) {
                            var v118 uint64 = 0
                            _ = v118
                        }
                    }
                    _ = v69
                }
                _ = v68
            }
            switch (fi66 - 4294967295) {
            case 0:
                for fi119 := 0; fi119 < 17; fi119++ {
                    var v120 int8 = (-(v19))
                    _ = fi119
                    _ = v120
                }
            default:
                for fi121 := 0; fi121 < 2; fi121++ {
                    var v122 int16 = 0
                    switch (-(fi121)) {
                    case 943:
                        v2 = append(v2, 4294967295)
                    }
                    var v123 uint64 = 0
                    v1 += (fi121 % ((42) | 1))
                    _ = fi121
                    _ = v122
                    _ = v123
                }
            }
        }
        _ = fi66
    }
    var v124 int64 = 1
    var v125 int16 = 42
    switch (v1 | v1) {
    case 171:
        v1 = v1
    case 32767:
        switch int(v124) {
        case 2147483647:
            var v126 uint = 18446744073709551615
            _ = v126
        case 42:
            var v127 int = v1
            _ = v127
        case 500:
            for fi128 := 0; fi128 < 12; fi128++ {
                fmt.Println(v3)
                v19 -= (v19 - (v19 * (v19 % ((-128) | 1))))
                if (1 >= 32767) {
                    v129 := GenS2{F0: 4294967295}
                    _ = len(v2)
                    _ = v129
                } else {
                    for fi130 := 0; fi130 < 16; fi130++ {
                        if len(v2) > 0 {
                            v2[0] = v3
                        }
                        _ = fi130
                    }
                    v131 := GenS6{F0: v1, F1: true}
                    v132 := GenS5{F0: fi128, F1: 4294967295, F2: v19}
                    _ = v131
                    _ = v132
                }
                _ = fi128
            }
        default:
            fmt.Println(v1)
        }
    case 4294967295:
        var v133 int8 = (-(v19))
        _ = v133
    }
    v2 = append(v2, v3)
    for fi134 := 0; fi134 < 11; fi134++ {
        v1 = v1
        fmt.Println(fi134)
        v19 = v19
        _ = fi134
    }
    var v135 uint = 63
    fmt.Println(v135, v135)
    if false {
        v136 := GenS7{F0: true, F1: (v1 / (((v1 * (v1 % (((v1 / ((32768) | 1))) | 1)))) | 1)), F2: false}
        var v137 int64 = v124
        if (!false) {
            var v138 string = ("foo bar" + "\n\t")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v19 += v19
            _ = v138
        }
        _ = v136
        _ = v137
    } else {
        for fi139 := 0; fi139 < 2; fi139++ {
            if false {
                for fi140 := 0; fi140 < 12; fi140++ {
                    var v141 uint16 = 42
                    if true {
                        _ = len(v2)
                    }
                    var v142 int16 = int16(v3)
                    _ = fi140
                    _ = v141
                    _ = v142
                }
                var v143 float32 = 1.5
                var v144 float32 = float32(0.1)
                _ = v143
                _ = v144
            }
            var v145 bool = (("" < ("\n\t" + "\n\t")) || true)
            _ = fi139
            _ = v145
        }
        v146 := GenS8{F0: -686, F1: 65535, F2: v1}
        _ = len(v2)
        _ = v146
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v147 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v135:", v135)
    fmt.Println("v147:", v147)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v124
    _ = v125
    _ = v135
    _ = v147
    _ = v19
    _ = v2
    _ = v3
    _ = v65
}
