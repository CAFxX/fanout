package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 string }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }
type GenS4 struct { F0 uint64; F1 int64; F2 int64 }
type GenS5 struct { F0 int32; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 bool; F1 int16; F2 int }
type GenS11 struct { F0 int; F1 int16; F2 int }
type GenS12 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS13 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS14 struct { F0 int16; F1 int64 }
type GenS15 struct { F0 uint }
func genFn0() uint16 {
    var v6 int = -64
    var v7 uint16 = 307
    if (true && false) {
        v8 := []int{(v6 * 42), v6, v6}
        var v9 int16 = -1
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 841, F1: 660, F2: "hello"}
            fmt.Println(v7)
            _ = fi10
            _ = v11
        }
        fmt.Println(v6, v7)
    }
    _ = v6
    _ = v7
    return (v7 ^ 65535)
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    v4 := []string{}
    var v5 int8 = 0
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS1{F0: 42, F1: 2, F2: v1}
        if (2 <= 169) {
            if (!((v5 << ((4294967295) & 63)) != v5)) {
                _ = len(v2)
                if (!true) {
                    v3 -= 2
                    fmt.Println(v1)
                } else {
                    v1 = "hello"
                }
            } else {
                if true {
                    delete(v2, (fi12 / (((-(fi12))) | 1)))
                    fmt.Println(v5, v3)
                }
                switch (fi12 * (fi12 % (((fi12 & fi12)) | 1))) {
                case 0:
                    fmt.Println(v1, v1, v3)
                case 42:
                    var v14 bool = ((42 != 1) && (290 >= 2))
                    _ = v14
                }
                fmt.Println(v1, v1, v1)
            }
        }
        _ = fi12
        _ = v13
    }
    if (2147483647 >= 1) {
        _ = len(v4)
        var v15 int64 = -1
        _ = v15
    } else {
        v16 := GenS1{F0: 450, F1: 966, F2: ("foo bar" + v1)}
        switch (v1 + (v1 + "\n\t")) {
        case "hello":
            fmt.Println(v1, v1, v5)
        default:
            _ = len(v2)
        }
        _ = v16
    }
    switch (("a" + "foo bar") + "\n\t") {
    case "a":
        if false {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint = 905
                _ = fi17
                _ = v18
            }
            switch "" {
            case "":
                if (42 != 228) {
                    switch ((("" + v1) + (v1 + (v1 + v1))) + (v1 + v1)) {
                    case "foo bar":
                        var v19 float32 = (1.5 / -2.25)
                        _ = v19
                    case "":
                        v20 := GenS1{F0: -1, F1: 2, F2: (((("\n\t" + v1) + (((v1 + (v1 + v1)) + "foo bar") + v1)) + (v1 + v1)) + (((("a" + ("hello" + (v1 + (v1 + ("\n\t" + ("\n\t" + v1)))))) + v1) + "a") + v1))}
                        _ = v20
                    default:
                        v21 := GenS2{F0: 100, F1: v3}
                        _ = v21
                    }
                    fmt.Println(v5, v3, v3)
                } else {
                    for fi22 := 0; fi22 < 6; fi22++ {
                        v1 = v1
                        for fi23 := 0; fi23 < 14; fi23++ {
                            var v24 int16 = int16(fi22)
                            _ = len(v2)
                            _ = fi23
                            _ = v24
                        }
                        _ = fi22
                    }
                    if true {
                        v25 := GenS1{F0: -2147483648, F1: 2, F2: ""}
                        v26 := GenS3{F0: false, F1: 165, F2: -936}
                        _ = v25
                        _ = v26
                    } else {
                        switch v1 {
                        case "a":
                            if len(v4) > 0 {
                                v4[0] = (v1 + v1)
                            }
                        case "hello":
                            var v27 uint8 = 0
                            _ = v27
                        default:
                            for fi28 := 0; fi28 < 7; fi28++ {
                                if false {
                                    v4 = append(v4, "")
                                    v3 -= (v3 << ((2) & 63))
                                }
                                for fi29 := 0; fi29 < 12; fi29++ {
                                    v30 := GenS2{F0: v5, F1: v3}
                                    _ = fi29
                                    _ = v30
                                }
                                if true {
                                    v5 = (v5 % (((v5 | v5)) | 1))
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        if true {
                                            v1 = v1
                                            v32 := v2[2147483647]
                                            _ = v32
                                        }
                                        switch 42 {
                                        case 1:
                                            var v33 uint64 = (v3 - (^(v3)))
                                            _ = v33
                                        case -199:
                                            var v34 uint = 0
                                            _ = v34
                                        case 32768:
                                            fmt.Println(v1)
                                        default:
                                            v35 := GenS4{F0: (v3 ^ (v3 >> ((2) & 63))), F1: 89, F2: 42}
                                            _ = v35
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            var v36 int = fi31
                                            _ = v36
                                        case "foo bar":
                                            var v37 uint64 = v3
                                            _ = v37
                                        default:
                                            var v38 string = (((v1 + ("hello" + v1)) + v1) + v1)
                                            _ = v38
                                        }
                                        _ = fi31
                                    }
                                } else {
                                    var v39 int16 = 42
                                    v40 := GenS5{F0: -1, F1: fi28}
                                    _ = v39
                                    _ = v40
                                }
                                _ = fi28
                            }
                        }
                    }
                    var v41 uint16 = 255
                    _ = v41
                }
            default:
                v42 := GenS6{F0: ((v1 + "\n\t") >= "hello")}
                _ = v42
            }
            v3 = v3
        } else {
            v43 := GenS4{F0: (v3 ^ v3), F1: -998, F2: 1}
            if true {
                _ = len(v4)
                switch v1 {
                case "\n\t":
                    if false {
                        v44 := GenS3{F0: false, F1: 781, F2: -1}
                        _ = v44
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v3, v5, v5)
                    }
                case "foo bar":
                    v1 = v1
                case "a":
                    var v45 int32 = 2147483647
                    _ = v45
                }
            } else {
                switch v1 {
                case "hello":
                    if false {
                        fmt.Println(v5, v3, v1)
                        switch (v1 + v1) {
                        case "hello":
                            v46 := GenS2{F0: v5, F1: 0}
                            _ = v46
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = v1
                            }
                        case "a":
                            v47 := GenS5{F0: 0, F1: -517}
                            _ = v47
                        }
                    } else {
                        var v48 uint64 = 622
                        _ = v48
                    }
                }
                var v49 float64 = (0.0 + float64(float32(((1e10 + 1e10) - 0.0))))
                _ = v49
            }
            _ = v43
        }
    case "\n\t":
        v50 := v2[-2147483648]
        _ = v50
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    if false {
        switch ("" + v1) {
        case "a":
            if false {
                var v51 string = "a"
                for fi52 := 0; fi52 < 5; fi52++ {
                    for fi53 := 0; fi53 < 1; fi53++ {
                        if true {
                            if true {
                                if ((true && false) || true) {
                                    if (3.14159 < 0.0) {
                                        var v54 uint8 = 25
                                        _ = len(v2)
                                        v5 += (v5 - (v5 * 1))
                                        _ = v54
                                    }
                                }
                                v3 = (v3 | 0)
                                var v55 bool = true
                                switch fi52 {
                                case -482:
                                    var v56 int = (fi53 & fi52)
                                    _ = v56
                                }
                                v57 := v2[(fi53 * 9223372036854775807)]
                                _ = v55
                                _ = v57
                            } else {
                                var v58 uint16 = 593
                                fmt.Println(v58, v58)
                                v51 = ("foo bar" + ((v1 + v1) + v51))
                                _ = v58
                            }
                            fmt.Println(fi52, fi52)
                            v5 -= -1
                        } else {
                            var v59 uint8 = 138
                            var v60 int8 = (v5 >> ((2) & 63))
                            v2[(^(fi53))] = (^(fi53))
                            v61 := GenS6{F0: (-785 != 9223372036854775807)}
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        var v62 uint = 65535
                        _ = len(v4)
                        var v63 int = fi53
                        v63 = (fi52 | fi53)
                        _ = fi53
                        _ = v62
                        _ = v63
                    }
                    if (627 == 2147483648) {
                        var v64 int8 = 20
                        v65 := GenS2{F0: v5, F1: v3}
                        fmt.Println(v5, v51)
                        _ = v64
                        _ = v65
                    } else {
                        var v66 bool = false
                        _ = v66
                    }
                    v3 += 2
                    _ = fi52
                }
                v1 = ((v1 + v1) + v51)
                _ = v51
            }
        default:
            var v67 float32 = (float32(-2.25) * ((1e10 + 3.14159) - float32(0.0)))
            _ = v67
        }
        if false {
            if (18446744073709551615 < 752) {
                var v68 uint64 = 1
                v69 := v2[-815]
                _ = v68
                _ = v69
            }
            for fi70 := 0; fi70 < 8; fi70++ {
                for fi71 := 0; fi71 < 1; fi71++ {
                    var v72 string = v1
                    switch ((v72 + (("\n\t" + v1) + v72)) + (v72 + v72)) {
                    case "a":
                        var v73 int16 = 252
                        _ = v73
                    case "foo bar":
                        if ((v3 >> ((18446744073709551615) & 63)) < v3) {
                            v74 := GenS3{F0: ((genFn0() > 475) || false), F1: 1, F2: 137}
                            v4 = append(v4, v72)
                            v1 = (v72 + "hello")
                            var v75 int = fi70
                            fmt.Println(v5, fi70, v5)
                            _ = v74
                            _ = v75
                        } else {
                            var v76 uint16 = 2
                            v77 := GenS7{F0: 32767, F1: 383, F2: (v76 * (v76 / ((0) | 1)))}
                            v72 = "foo bar"
                            _ = v76
                            _ = v77
                        }
                    case "hello":
                        if false {
                            v78 := v2[fi71]
                            fmt.Println(v72, v3, v5)
                            v79 := GenS8{F0: 419}
                            v72 = (v72 + (v1 + "hello"))
                            _ = v78
                            _ = v79
                        }
                    default:
                        v3 -= (v3 * 0)
                    }
                    var v80 uint64 = 18446744073709551615
                    _ = fi71
                    _ = v72
                    _ = v80
                }
                _ = fi70
            }
        }
        if len(v4) > 0 {
            v4[0] = v1
        }
    } else {
        if false {
            switch 42 {
            case -2147483648:
                delete(v2, 2147483647)
            case 202:
                fmt.Println(v3)
            }
            fmt.Println(v3)
            switch 1 {
            case 2147483647:
                v3 = 331
            }
        }
        if false {
            var v81 bool = (0 == 42)
            if ((525 >= v3) && (v81 || ((645 <= 655) && true))) {
                for fi82 := 0; fi82 < 20; fi82++ {
                    for fi83 := 0; fi83 < 19; fi83++ {
                        switch "a" {
                        case "a":
                            v84 := GenS2{F0: 0, F1: (v3 / (((v3 << ((403) & 63))) | 1))}
                            _ = v84
                        case "\n\t":
                            _ = len(v2)
                        case "":
                            var v85 uint64 = (v3 + (^(v3)))
                            _ = v85
                        }
                        if ((!(0.0 <= float64(((-2.25 / float32((1.5 / -2.25))) - 1e10)))) || v81) {
                            var v86 string = "foo bar"
                            if (v3 >= 4294967295) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v1 = v86
                            }
                            _ = v86
                        }
                        if (0 != 651) {
                            if v81 {
                                var v87 uint = 0
                                var v88 float32 = 1.5
                                if (("\n\t" + v1) > v1) {
                                    if true {
                                        v89 := GenS9{F0: -1}
                                        _ = v89
                                    } else {
                                        var v90 uint = 18446744073709551615
                                        fmt.Println(v1)
                                        _ = v90
                                    }
                                    fmt.Println(fi82, v3, v88)
                                } else {
                                    var v91 float32 = float32(0.0)
                                    if false {
                                        fmt.Println(fi82)
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if ((v1 + v1) >= ("hello" + (v1 + v1))) {
                                            switch (v1 + v1) {
                                            case "a":
                                                v92 := GenS10{F0: v81, F1: 42, F2: fi82}
                                                v93 := GenS11{F0: fi82, F1: 1000, F2: fi83}
                                                _ = v92
                                                _ = v93
                                            case "":
                                                var v94 int64 = int64(v3)
                                                _ = v94
                                            case "hello":
                                                var v95 float64 = ((1e10 / 1e10) * float64(v88))
                                                _ = v95
                                            }
                                            var v96 uint32 = 318
                                            _ = v96
                                        } else {
                                            if false {
                                                var v97 string = v1
                                                _ = v97
                                            } else {
                                                if (0 <= -19) {
                                                    delete(v2, (fi83 | (fi82 - fi82)))
                                                    var v98 float64 = 1.5
                                                    v99 := GenS10{F0: v81, F1: -32768, F2: fi82}
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    fmt.Println(fi82, v87)
                                                    v100 := GenS12{F0: 0, F1: 32767, F2: (v3 & 1)}
                                                    _ = v100
                                                }
                                                v4 = append(v4, ("foo bar" + v1))
                                                v5 = v5
                                            }
                                        }
                                    }
                                    _ = v91
                                }
                                v101 := GenS6{F0: ((0 > 1) && (127 <= 42))}
                                switch (fi83 << (((v87 << (((v87 * (v87 * (v87 - v87)))) & 63))) & 63)) {
                                case -553:
                                    switch v1 {
                                    case "hello":
                                        if v81 {
                                            if (!((649 < 2147483648) || true)) {
                                                var v102 float32 = v88
                                                var v103 bool = v81
                                                _ = v102
                                                _ = v103
                                            } else {
                                                v87 = 983
                                                fmt.Println(fi82, v81)
                                            }
                                            v104 := v2[32767]
                                            _ = v104
                                        }
                                    case "":
                                        var v105 uint8 = 2
                                        _ = v105
                                    case "\n\t":
                                        v81 = (true || true)
                                    default:
                                        switch "hello" {
                                        case "a":
                                            delete(v2, 442)
                                        case "foo bar":
                                            v2[fi83] = fi82
                                        case "":
                                            var v106 int = (fi82 ^ (fi83 & (fi83 >> ((v87) & 63))))
                                            _ = v106
                                        default:
                                            v1 = ("" + ("hello" + v1))
                                        }
                                    }
                                    if v81 {
                                        v87 -= (v87 | v87)
                                    }
                                }
                                _ = v101
                                _ = v87
                                _ = v88
                            }
                        }
                        v107 := GenS13{F0: 272, F1: 251, F2: genFn0()}
                        _ = fi83
                        _ = v107
                    }
                    _ = fi82
                }
                fmt.Println(v5, v1, v81)
                var v108 uint16 = genFn0()
                _ = v108
            } else {
                v109 := GenS8{F0: 18446744073709551615}
                var v110 float32 = (1.5 - -2.25)
                for fi111 := 0; fi111 < 14; fi111++ {
                    _ = len(v4)
                    _ = fi111
                }
                _ = v109
                _ = v110
            }
            _ = v81
        }
    }
    for fi112 := 0; fi112 < 15; fi112++ {
        switch -9223372036854775808 {
        case -1:
            _ = len(v4)
        case 328:
            switch "foo bar" {
            case "a":
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
            default:
                if false {
                    var v113 float64 = (0.1 / 1.0)
                    switch fi112 {
                    case 127:
                        if true {
                            v114 := GenS12{F0: 438, F1: genFn0(), F2: v3}
                            switch ("foo bar" + (v1 + "\n\t")) {
                            case "hello":
                                if (false || false) {
                                    v113 -= ((0.1 - (0.0 * -2.25)) / float64(0.1))
                                }
                            case "\n\t":
                                v113 += 1e10
                            case "":
                                var v115 uint16 = 42
                                _ = v115
                            }
                            v1 = v1
                            _ = v114
                        }
                    case 2147483647:
                        delete(v2, (fi112 / (((^(fi112))) | 1)))
                    default:
                        v116 := GenS12{F0: 2147483647, F1: 248, F2: (v3 << ((641) & 63))}
                        _ = v116
                    }
                    for fi117 := 0; fi117 < 13; fi117++ {
                        v118 := v2[127]
                        _ = fi117
                        _ = v118
                    }
                    _ = v113
                }
            }
        case -927:
            var v119 int8 = v5
            _ = v119
        default:
            var v120 float64 = (0.1 + 1.5)
            _ = v120
        }
        _ = fi112
    }
    for fi121 := 0; fi121 < 18; fi121++ {
        switch (-(fi121)) {
        case 127:
            if true {
                for fi122 := 0; fi122 < 20; fi122++ {
                    v123 := GenS14{F0: 42, F1: 0}
                    _ = fi122
                    _ = v123
                }
                var v124 int8 = -128
                v124 -= 127
                _ = v124
            }
        case 32768:
            v1 = v1
        case 374:
            for fi125 := 0; fi125 < 10; fi125++ {
                v5 = (v5 & (v5 << ((2) & 63)))
                _ = fi125
            }
        default:
            if (true || false) {
                for fi126 := 0; fi126 < 18; fi126++ {
                    v2[(fi126 * (fi126 | int(v3)))] = fi126
                    _ = fi126
                }
            }
        }
        v1 = "a"
        var v127 uint64 = (v3 << ((21) & 63))
        _ = fi121
        _ = v127
    }
    v3 = (v3 << ((1) & 63))
    for fi128 := 0; fi128 < 10; fi128++ {
        v1 = "hello"
        var v129 int = fi128
        _ = fi128
        _ = v129
    }
    v1 = v1
    var v130 float32 = (0.1 + 0.0)
    for fi131 := 0; fi131 < 19; fi131++ {
        var v132 int64 = -799
        _ = fi131
        _ = v132
    }
    fmt.Println(v3)
    v1 = ((v1 + v1) + (v1 + v1))
    switch 42 {
    case 2147483647:
        fmt.Println(v1, v5, v1)
    case -741:
        var v133 uint32 = 644
        _ = v133
    }
    v134 := GenS4{F0: (v3 & v3), F1: -608, F2: 42}
    for fi135 := 0; fi135 < 11; fi135++ {
        v136 := GenS15{F0: 255}
        var v137 int = 9223372036854775807
        v5 = -112
        _ = fi135
        _ = v136
        _ = v137
    }
    for fi138 := 0; fi138 < 18; fi138++ {
        fmt.Println(fi138)
        _ = fi138
    }
    var v139 uint16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v139:", v139)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v130
    _ = v134
    _ = v139
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
