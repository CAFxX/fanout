package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 uint64; F2 uint }
type GenS8 struct { F0 uint; F1 uint; F2 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 int32; F1 int8; F2 string }
type GenS17 struct { F0 string; F1 uint8; F2 bool }
type GenS18 struct { F0 uint8; F1 uint }
func genFn0(p0 int16, p1 int16, p2 bool) int16 {
    var v6 uint32 = 65535
    if p2 {
        p2 = p2
        if p2 {
            var v7 bool = (-128 >= 42)
            var v8 uint8 = 255
            _ = v7
            _ = v8
        } else {
            if (p1 < -148) {
                fmt.Println(p1)
            } else {
                v9 := GenS1{F0: v6}
                _ = v9
            }
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{(fi10 / ((fi10) | 1)), fi10}
                _ = fi10
                _ = v11
            }
            p2 = (44 != 255)
        }
        var v12 float32 = 1e10
        _ = v12
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 19; fi14++ {
                p1 = p0
                _ = fi14
            }
            v15 := GenS2{F0: fi13, F1: v6, F2: (-2.25 < -2.25)}
            if false {
                fmt.Println(p1, v6, v6)
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS3{F0: 2147483647, F1: v6, F2: 319}
                    var v18 uint8 = 255
                    v19 := GenS2{F0: (fi13 - -1), F1: v6, F2: ((1 == 278) && (0.1 == 0.0))}
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            } else {
                var v20 string = "foo bar"
                p2 = false
                for fi21 := 0; fi21 < 6; fi21++ {
                    v6 = v6
                    switch (fi21 - (fi21 / ((fi13) | 1))) {
                    case 445:
                        v22 := []int{(fi13 >> ((4294967295) & 63)), (fi21 << ((255) & 63))}
                        _ = v22
                    case -2147483648:
                        p1 = (p1 >> ((255) & 63))
                        v23 := GenS1{F0: v6}
                        if false {
                            v24 := GenS1{F0: v6}
                            if p2 {
                                v25 := GenS4{F0: 2}
                                var v26 int64 = -1
                                v27 := []int{4294967295, (fi21 ^ fi21), -889, (^(fi21))}
                                switch fi21 {
                                case 128:
                                    p1 += p1
                                    v20 = v20
                                case 42:
                                    if len(v27) >= 2 {
                                        v27 = v27[1:len(v27)]
                                    }
                                case 32768:
                                    v28 := GenS5{F0: 65535, F1: 100}
                                    _ = v28
                                default:
                                    v29 := GenS6{F0: ""}
                                    _ = v29
                                }
                                switch fi21 {
                                case -359:
                                    v20 = "hello"
                                default:
                                    switch 4294967295 {
                                    case -911:
                                        v20 = ("foo bar" + "hello")
                                    case 42:
                                        v27 = append(v27, (fi13 << ((2) & 63)))
                                    case 1:
                                        var v30 uint8 = 225
                                        fmt.Println(p0, fi21, fi21)
                                        _ = v30
                                    }
                                }
                                _ = v25
                                _ = v26
                                _ = v27
                            } else {
                                switch (fi13 - (fi21 ^ (fi21 / ((fi13) | 1)))) {
                                case -53:
                                    p1 = p0
                                case -9223372036854775808:
                                    p2 = false
                                default:
                                    var v31 int16 = p0
                                    _ = v31
                                }
                            }
                            _ = v24
                        } else {
                            if ((531 > 1) || p2) {
                                v32 := []int{fi21, (^(fi21)), fi13, 989}
                                fmt.Println(fi21, fi13)
                                _ = v32
                            }
                            if (!(-2.25 < 1.5)) {
                                var v33 uint64 = 18446744073709551615
                                v34 := GenS7{F0: (!p2), F1: 4294967295, F2: 255}
                                if true {
                                    v35 := []int{fi13, (fi21 * (fi21 % (((fi13 << ((1) & 63))) | 1))), 1}
                                    var v36 float32 = -2.25
                                    var v37 string = v20
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                _ = v33
                                _ = v34
                            }
                            fmt.Println(v20, p0)
                            var v38 int = 4294967295
                            p1 = (p1 * (p1 >> ((932) & 63)))
                            _ = v38
                        }
                        switch (((v20 + "foo bar") + "foo bar") + (v20 + v20)) {
                        case "":
                            var v39 uint32 = (^(v6))
                            _ = v39
                        case "hello":
                            var v40 float64 = 3.14159
                            _ = v40
                        default:
                            v20 = v20
                        }
                        switch 127 {
                        case 42:
                            fmt.Println(fi21, p0, v20)
                        default:
                            var v41 int8 = -43
                            _ = v41
                        }
                        p1 = p1
                        _ = v23
                    case 127:
                        v42 := []int{(fi13 + fi13), 0, (fi13 / (((^(fi13))) | 1))}
                        _ = v42
                    default:
                        if p2 {
                            switch v20 {
                            case "foo bar":
                                p1 -= (^(p1))
                            case "":
                                var v43 uint64 = 0
                                _ = v43
                            case "\n\t":
                                v20 = "a"
                            }
                            var v44 uint64 = 1
                            _ = v44
                        } else {
                            var v45 int64 = 0
                            if (!p2) {
                                if p2 {
                                    v46 := GenS8{F0: 18446744073709551615, F1: 4294967295, F2: 0}
                                    _ = v46
                                }
                            }
                            var v47 bool = (!p2)
                            _ = v45
                            _ = v47
                        }
                    }
                    _ = fi21
                }
                _ = v20
            }
            _ = fi13
            _ = v15
        }
        for fi48 := 0; fi48 < 2; fi48++ {
            p0 -= p0
            _ = fi48
        }
        if false {
            var v49 int64 = 9223372036854775807
            _ = v49
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]uint32)
    v3 := []string{(("\n\t" + "\n\t") + "hello")}
    var v4 uint = 65535
    var v5 uint32 = 772
    if ((("\n\t" + "a") < "hello") && true) {
        v50 := GenS9{F0: v4}
        _ = v50
    } else {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v5)
        case "hello":
            v51 := GenS10{F0: 103, F1: -1}
            _ = v51
        }
        if (1 >= -818) {
            v52 := GenS3{F0: 481, F1: v5, F2: 0}
            _ = v52
        } else {
            switch -1 {
            case 32767:
                if (-1 > -1) {
                    if ((-1 != int(v4)) && ((false || false) || (280 <= 65535))) {
                        fmt.Println(v5, v5, v5)
                    } else {
                        v53 := GenS11{F0: 0}
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 10; fi54++ {
                        v55 := GenS12{F0: 32767, F1: 1}
                        for fi56 := 0; fi56 < 6; fi56++ {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            var v57 int8 = 0
                            v2[fi56] = 489
                            fmt.Println(v5, v57, fi54)
                            _ = fi56
                            _ = v57
                        }
                        _ = fi54
                        _ = v55
                    }
                    if (false || (!(18446744073709551615 <= 83))) {
                        v1 -= v1
                        fmt.Println(v4)
                    }
                }
            }
        }
        var v58 float64 = (0.0 + 0.0)
        _ = v58
    }
    var v59 int64 = -1
    var v60 float32 = float32(0.1)
    if true {
        switch "foo bar" {
        case "a":
            var v61 int64 = (v59 / ((v59) | 1))
            _ = v61
        case "":
            v62 := GenS9{F0: 1}
            _ = v62
        default:
            if (42 == 116) {
                v63 := GenS13{F0: (!true)}
                v60 = (3.14159 - v60)
                for fi64 := 0; fi64 < 11; fi64++ {
                    v4 = v4
                    v59 = v59
                    _ = fi64
                }
                _ = v63
            }
        }
        fmt.Println(v59, v1, v5)
    } else {
        v65 := v2[int(v5)]
        if (("hello" + (("hello" + "foo bar") + ("a" + ""))) > "foo bar") {
            if (!true) {
                var v66 float64 = 0.0
                _ = v66
            }
            for fi67 := 0; fi67 < 8; fi67++ {
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                if (genFn0(genFn0((v1 % ((-814) | 1)), (^(v1)), false), v1, true) <= (^(v1))) {
                    fmt.Println(v60, v4)
                    for fi68 := 0; fi68 < 12; fi68++ {
                        v69 := GenS8{F0: v4, F1: 42, F2: 0}
                        switch (-(fi68)) {
                        case 127:
                            var v70 string = "foo bar"
                            _ = v70
                        default:
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        }
                        _ = fi68
                        _ = v69
                    }
                }
                _ = fi67
            }
            for fi71 := 0; fi71 < 16; fi71++ {
                if false {
                    v72 := v2[-100]
                    var v73 int8 = -25
                    _ = v72
                    _ = v73
                }
                _ = fi71
            }
        }
        _ = v65
    }
    v1 += v1
    fmt.Println(v4, v5, v4)
    if (42 <= 1) {
        v74 := GenS14{F0: 42}
        var v75 bool = ("\n\t" == "foo bar")
        _ = v74
        _ = v75
    } else {
        var v76 string = ""
        v3 = append(v3, v76)
        _ = v76
    }
    v60 -= v60
    for fi77 := 0; fi77 < 19; fi77++ {
        fmt.Println(v59)
        if (!(127 == 0)) {
            v78 := GenS11{F0: (v59 ^ v59)}
            if ((-124 <= fi77) && true) {
                v5 -= v5
                switch "a" {
                case "a":
                    v3 = append(v3, "a")
                case "\n\t":
                    if (v59 > v59) {
                        v79 := GenS10{F0: 2, F1: 2147483647}
                        if ((!(true && (!false))) && false) {
                            var v80 string = ((("\n\t" + ("foo bar" + "foo bar")) + "") + "\n\t")
                            v4 += (v4 - (v4 >> ((18446744073709551615) & 63)))
                            var v81 uint16 = 424
                            _ = v80
                            _ = v81
                        } else {
                            var v82 int8 = 127
                            var v83 int8 = -128
                            v82 += 53
                            _ = v82
                            _ = v83
                        }
                        fmt.Println(v60, v1)
                        _ = v79
                    } else {
                        fmt.Println(v4)
                        for fi84 := 0; fi84 < 6; fi84++ {
                            switch ("" + "foo bar") {
                            case "\n\t":
                                if false {
                                    switch fi77 {
                                    case 247:
                                        var v85 uint32 = v5
                                        _ = v85
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if true {
                                        switch (fi77 << (((^(v4))) & 63)) {
                                        case 32767:
                                            switch fi84 {
                                            case 252:
                                                v4 = 191
                                            case -9223372036854775808:
                                                var v86 uint8 = 42
                                                _ = v86
                                            case 69:
                                                var v87 float64 = (1e10 + 1e10)
                                                _ = v87
                                            default:
                                                var v88 uint32 = 42
                                                _ = v88
                                            }
                                        case 128:
                                            v5 = 1
                                        case -2147483648:
                                            if true {
                                                switch 32768 {
                                                case 42:
                                                    var v89 int = (fi84 & fi77)
                                                    _ = v89
                                                }
                                                if (!((v4 - v4) >= 0)) {
                                                    if true {
                                                        switch "" {
                                                        case "foo bar":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            v1 = v1
                                                            v90 := GenS4{F0: 675}
                                                            v59 += -1
                                                            _ = v90
                                                        }
                                                        fmt.Println(fi84, v60, v5)
                                                        if false {
                                                            v60 = (float32((1e10 * (-2.25 / (-2.25 * 0.1)))) - v60)
                                                        }
                                                    }
                                                }
                                                switch fi77 {
                                                case 795:
                                                    fmt.Println(v4)
                                                case 9223372036854775807:
                                                    fmt.Println(v59, v5, v4)
                                                case 42:
                                                    v91 := GenS3{F0: v5, F1: (v5 * (^(v5))), F2: 227}
                                                    _ = v91
                                                default:
                                                    var v92 int = fi77
                                                    _ = v92
                                                }
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                } else {
                                                    if true {
                                                        v5 -= (v5 ^ 65535)
                                                    } else {
                                                        v93 := GenS15{F0: genFn0((v1 % ((int16(fi84)) | 1)), -32768, true), F1: genFn0(genFn0((v1 + v1), (v1 / ((42) | 1)), true), 42, false)}
                                                        if (0.1 <= -2.25) {
                                                            var v94 uint8 = 255
                                                            var v95 float64 = 3.14159
                                                            v3 = append(v3, "hello")
                                                            var v96 int8 = 1
                                                            _ = v94
                                                            _ = v95
                                                            _ = v96
                                                        } else {
                                                            v97 := GenS14{F0: 787}
                                                            switch int(v1) {
                                                            case -2147483648:
                                                                var v98 uint8 = 255
                                                                _ = v98
                                                            case -731:
                                                                var v99 uint16 = 524
                                                                _ = v99
                                                            case 9223372036854775807:
                                                                v100 := GenS16{F0: 2147483647, F1: 42, F2: "\n\t"}
                                                                v1 -= genFn0(42, 42, false)
                                                                var v101 string = "hello"
                                                                v60 += float32(1e10)
                                                                var v102 string = v101
                                                                _ = v100
                                                                _ = v101
                                                                _ = v102
                                                            default:
                                                                if true {
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                            if true {
                                                                v60 -= v60
                                                                var v103 int32 = -754
                                                                _ = v103
                                                            }
                                                            _ = v97
                                                        }
                                                        v104 := GenS17{F0: "foo bar", F1: 37, F2: (!true)}
                                                        _ = v104
                                                        _ = v93
                                                    }
                                                }
                                            }
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v105 int64 = v59
                                                _ = v105
                                            case "":
                                                v60 = (1.5 - 1e10)
                                            }
                                        }
                                    } else {
                                        if true {
                                            fmt.Println(v4, fi77, fi77)
                                            switch fi84 {
                                            case 2147483647:
                                                v59 += (v59 * (^(v59)))
                                                v1 += v1
                                            case 4294967295:
                                                var v106 int32 = -840
                                                _ = v106
                                            default:
                                                v3 = append(v3, "\n\t")
                                            }
                                        } else {
                                            var v107 uint16 = 255
                                            _ = v107
                                        }
                                        switch ("\n\t" + ("a" + "hello")) {
                                        case "hello":
                                            if false {
                                                v108 := GenS18{F0: 1, F1: 1}
                                                var v109 uint = uint(v5)
                                                if true {
                                                    if (fi84 > (fi77 & (fi77 - (fi77 * (fi84 + fi84))))) {
                                                        v5 = 623
                                                    }
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v110 int = fi84
                                                fmt.Println(v5)
                                                _ = v108
                                                _ = v109
                                                _ = v110
                                            } else {
                                                var v111 int16 = int16(v59)
                                                if len(v3) > 0 {
                                                    v3[0] = "\n\t"
                                                }
                                                v112 := GenS11{F0: 9223372036854775807}
                                                _ = v111
                                                _ = v112
                                            }
                                        }
                                        if true {
                                            v1 -= (^(v1))
                                            switch fi77 {
                                            case 1:
                                                if false {
                                                    var v113 bool = (688 >= 192)
                                                    v114 := GenS12{F0: (v1 * genFn0(-380, (v1 >> ((uint(v5)) & 63)), (4294967295 <= v4))), F1: -821}
                                                    _ = v113
                                                    _ = v114
                                                }
                                            default:
                                                var v115 uint64 = 834
                                                _ = v115
                                            }
                                        } else {
                                            var v116 int16 = (v1 >> (((v4 % (((v4 - (v4 >> ((v4) & 63)))) | 1))) & 63))
                                            _ = v116
                                        }
                                    }
                                }
                            case "foo bar":
                                var v117 int32 = 42
                                _ = v117
                            case "a":
                                fmt.Println(v5, v4)
                            default:
                                if (4294967295 > 2) {
                                    var v118 bool = (false && true)
                                    _ = v118
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                _ = len(v3)
                            case "foo bar":
                                if (-1 <= (^(v59))) {
                                    var v119 bool = false
                                    var v120 int16 = v1
                                    _ = v119
                                    _ = v120
                                } else {
                                    v1 = (v1 + genFn0(1000, (^(v1)), ((!false) || true)))
                                }
                            }
                            _ = fi84
                        }
                        switch -2147483648 {
                        case -2147483648:
                            v4 = (v4 % (((v4 * 332)) | 1))
                        case -1:
                            v121 := GenS1{F0: v5}
                            _ = v121
                        default:
                            if (int(v59) < 4294967295) {
                                for fi122 := 0; fi122 < 10; fi122++ {
                                    var v123 uint8 = 0
                                    if (42 == 0) {
                                        switch "" {
                                        case "":
                                            if (!(0.0 > (0.0 + 1.5))) {
                                                v1 = v1
                                                var v124 float64 = float64(-2.25)
                                                _ = v124
                                            }
                                        case "foo bar":
                                            v4 -= uint(v5)
                                        case "\n\t":
                                            v125 := GenS15{F0: (^(v1)), F1: (v1 << (((^(v4))) & 63))}
                                            _ = v125
                                        default:
                                            var v126 int32 = 0
                                            _ = v126
                                        }
                                        var v127 int = (-(fi77))
                                        v1 -= (v1 + (^(v1)))
                                        _ = v127
                                    } else {
                                        v60 -= v60
                                        fmt.Println(fi77, v123)
                                    }
                                    _ = fi122
                                    _ = v123
                                }
                                var v128 uint32 = (v5 & (v5 << (((^(v4))) & 63)))
                                var v129 bool = (-1 == 986)
                                _ = v128
                                _ = v129
                            } else {
                                v60 = (float32(((0.0 - 1.5) / 3.14159)) - 0.0)
                                for fi130 := 0; fi130 < 8; fi130++ {
                                    var v131 uint8 = 218
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi130
                                    _ = v131
                                }
                            }
                        }
                    }
                case "hello":
                    v132 := GenS1{F0: 2147483648}
                    _ = v132
                default:
                    v5 = (v5 / ((v5) | 1))
                }
                for fi133 := 0; fi133 < 19; fi133++ {
                    _ = len(v3)
                    var v134 float32 = v60
                    v4 = v4
                    _ = fi133
                    _ = v134
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v135 := GenS2{F0: 32767, F1: (v5 | v5), F2: (v59 > (^(v59)))}
                for fi136 := 0; fi136 < 10; fi136++ {
                    switch 497 {
                    case 128:
                        v137 := GenS18{F0: 2, F1: v4}
                        _ = v137
                    case 1:
                        var v138 uint = v4
                        _ = v138
                    case 0:
                        v4 = 18446744073709551615
                    default:
                        v59 = v59
                    }
                    var v139 uint64 = 722
                    _ = fi136
                    _ = v139
                }
                _ = v135
            }
            _ = v78
        } else {
            for fi140 := 0; fi140 < 3; fi140++ {
                fmt.Println(v1)
                if ((-2.25 * 1.5) < 0.1) {
                    var v141 uint64 = 219
                    var v142 int = (^(fi140))
                    var v143 bool = false
                    _ = v141
                    _ = v142
                    _ = v143
                }
                var v144 float64 = (1e10 - float64(v60))
                _ = fi140
                _ = v144
            }
            v145 := GenS5{F0: 385, F1: -1}
            _ = v145
        }
        _ = fi77
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
