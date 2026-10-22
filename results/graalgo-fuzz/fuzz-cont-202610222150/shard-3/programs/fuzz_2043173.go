package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 int32; F1 uint64 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32; F1 string }
type GenS13 struct { F0 uint64; F1 string; F2 int32 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 14; fi5++ {
        if (true || true) {
            switch "" {
            case "hello":
                v6 := []int{(^(fi5)), fi5}
                _ = v6
            case "\n\t":
                v7 := GenS1{F0: 151, F1: (fi5 | (fi5 + (fi5 & (fi5 + fi5))))}
                _ = v7
            case "":
                v8 := GenS1{F0: 454, F1: -987}
                _ = v8
            default:
                v9 := GenS1{F0: 42, F1: (fi5 | fi5)}
                _ = v9
            }
            v10 := GenS2{F0: 42}
            _ = v10
        } else {
            fmt.Println(fi5)
            fmt.Println(fi5)
        }
        _ = fi5
    }
    if (3.14159 >= 1.5) {
        var v11 float64 = 1.5
        v12 := GenS1{F0: 905, F1: 127}
        _ = v11
        _ = v12
    } else {
        var v13 int8 = -128
        v14 := GenS3{F0: (v13 * 73), F1: (("hello" + (("\n\t" + "\n\t") + "\n\t")) + ("a" + "\n\t"))}
        for fi15 := 0; fi15 < 19; fi15++ {
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := []int{(fi16 * (fi15 - (fi16 % (((fi15 / (((fi16 >> ((1) & 63))) | 1))) | 1)))), (fi15 / ((128) | 1))}
                switch (fi16 ^ fi15) {
                case 32768:
                    var v18 float64 = (1e10 - 3.14159)
                    _ = v18
                case 32767:
                    v19 := GenS1{F0: 0, F1: fi15}
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
            v13 -= 42
            v13 -= -52
            _ = fi15
        }
        _ = v13
        _ = v14
    }
    return 622
}
func genFn1() int {
    switch "\n\t" {
    case "":
        var v20 uint32 = 1
        _ = v20
    case "a":
        var v21 int32 = 42
        _ = v21
    case "foo bar":
        var v22 int8 = -1
        _ = v22
    default:
        switch -1 {
        case -2147483648:
            if (0 >= 2) {
                var v23 int64 = 1
                var v24 string = (("foo bar" + "a") + "")
                _ = v23
                _ = v24
            }
        default:
            if true {
                var v25 uint = 255
                _ = v25
            }
        }
    }
    var v26 bool = (!false)
    fmt.Println(v26)
    _ = v26
    return 334
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = false
    v3 := []string{"a", "foo bar"}
    var v4 int64 = -9223372036854775808
    v4 += (v4 % ((v4) | 1))
    if v2 {
        v27 := GenS4{F0: 781}
        fmt.Println(v4, v1)
        var v28 int = 26
        _ = v27
        _ = v28
    }
    if (!((18446744073709551615 == genFn0()) || v2)) {
        v29 := GenS5{F0: 42}
        _ = v29
    }
    if v2 {
        if (997 >= genFn0()) {
            for fi30 := 0; fi30 < 2; fi30++ {
                v1 -= v1
                v31 := GenS5{F0: 868}
                _ = fi30
                _ = v31
            }
            var v32 float64 = 1e10
            _ = v32
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            if true {
                v2 = v2
            } else {
                var v34 uint64 = (v1 >> ((4294967295) & 63))
                _ = v34
            }
            v35 := GenS6{F0: 529, F1: v2}
            switch "\n\t" {
            case "hello":
                var v36 int = genFn1()
                _ = v36
            case "\n\t":
                v37 := GenS3{F0: int8(v1), F1: ("\n\t" + "foo bar")}
                _ = v37
            case "":
                var v38 int = fi33
                _ = v38
            default:
                v4 += (v4 * (v4 | v4))
            }
            _ = fi33
            _ = v35
        }
        v4 = 42
    }
    if true {
        v2 = (int(v1) > 1)
    } else {
        if v2 {
            var v39 int8 = -1
            v4 = (v4 | (-(v4)))
            _ = v39
        }
        if (((v4 - (v4 & (-(v4)))) <= v4) && (55 < 42)) {
            v1 = (^(v1))
            v4 -= (v4 << ((509) & 63))
            var v40 uint64 = v1
            _ = v40
        } else {
            for fi41 := 0; fi41 < 20; fi41++ {
                fmt.Println(v2, fi41, fi41)
                v4 = (v4 >> ((444) & 63))
                for fi42 := 0; fi42 < 1; fi42++ {
                    var v43 uint32 = 65535
                    switch "\n\t" {
                    case "\n\t":
                        if v2 {
                            if false {
                                v2 = (!true)
                                if false {
                                    v44 := GenS7{F0: (v1 - v1), F1: genFn1()}
                                    _ = v44
                                }
                            }
                            v45 := GenS8{F0: 65535, F1: 391}
                            _ = v45
                        }
                    }
                    if v2 {
                        fmt.Println(fi42, fi42)
                        v46 := GenS6{F0: 2147483647, F1: true}
                        var v47 uint32 = (v43 >> ((18446744073709551615) & 63))
                        _ = v46
                        _ = v47
                    } else {
                        var v48 int8 = -126
                        var v49 int16 = -324
                        v4 = (v4 + 1)
                        _ = v48
                        _ = v49
                    }
                    _ = fi42
                    _ = v43
                }
                _ = fi41
            }
            var v50 uint8 = 255
            _ = v50
        }
    }
    if (127 > -263) {
        for fi51 := 0; fi51 < 6; fi51++ {
            if (!true) {
                if (v2 || false) {
                    v4 -= v4
                    fmt.Println(v4, v1)
                    v52 := GenS6{F0: -336, F1: (1e10 == 3.14159)}
                    _ = v52
                }
                if len(v3) > 0 {
                    v3[0] = ""
                }
            } else {
                fmt.Println(fi51)
                fmt.Println(v2)
            }
            for fi53 := 0; fi53 < 7; fi53++ {
                fmt.Println(v4, fi51)
                if v2 {
                    v1 = 5
                    var v54 int8 = 42
                    _ = v54
                } else {
                    v3 = append(v3, "\n\t")
                    if false {
                        if v2 {
                            _ = len(v3)
                            v2 = false
                            switch fi51 {
                            case -9223372036854775808:
                                switch "a" {
                                case "":
                                    var v55 uint64 = (v1 / ((genFn0()) | 1))
                                    var v56 uint = 255
                                    if (0.1 == 0.1) {
                                        if false {
                                            fmt.Println(fi53, v2, fi53)
                                            v1 = v1
                                            switch "hello" {
                                            case "hello":
                                                v57 := GenS7{F0: v1, F1: fi51}
                                                _ = v57
                                            }
                                            fmt.Println(fi53, v55, fi53)
                                        } else {
                                            if (128 != fi53) {
                                                var v58 uint16 = 255
                                                fmt.Println(v2, v1)
                                                _ = v58
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if false {
                                            var v59 float64 = (1e10 - 1e10)
                                            _ = v59
                                        } else {
                                            var v60 uint32 = 65535
                                            v1 -= uint64(v56)
                                            switch fi51 {
                                            case -964:
                                                if (false && v2) {
                                                    var v61 int32 = -149
                                                    fmt.Println(fi53, fi51)
                                                    switch "a" {
                                                    case "":
                                                        var v62 uint32 = (v60 & 42)
                                                        _ = v62
                                                    case "foo bar":
                                                        v55 += genFn0()
                                                    default:
                                                        fmt.Println(fi51)
                                                    }
                                                    v4 += (v4 >> ((v56) & 63))
                                                    _ = len(v3)
                                                    _ = v61
                                                } else {
                                                    var v63 float32 = ((1.5 / 0.1) + (0.1 + -2.25))
                                                    var v64 uint64 = 4294967295
                                                    fmt.Println(v4, v64, v4)
                                                    _ = v63
                                                    _ = v64
                                                }
                                                v56 -= 42
                                                var v65 int64 = (v4 >> (((^(v56))) & 63))
                                                _ = v65
                                            default:
                                                v66 := GenS9{F0: -953, F1: 496}
                                                _ = v66
                                            }
                                            v4 = (-(v4))
                                            _ = v60
                                        }
                                    }
                                    var v67 string = "\n\t"
                                    v68 := GenS10{F0: (!v2), F1: ("\n\t" + v67)}
                                    v1 -= v1
                                    if v2 {
                                        var v69 int = (^(fi53))
                                        var v70 int8 = 91
                                        _ = v69
                                        _ = v70
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (v67 + "foo bar")
                                        }
                                        if v2 {
                                            fmt.Println(v55, v4, v56)
                                        }
                                        switch 381 {
                                        case 4294967295:
                                            v71 := GenS4{F0: 628}
                                            _ = v71
                                        case 1:
                                            v2 = (-794 > 1)
                                        }
                                    }
                                    switch (fi53 - fi53) {
                                    case 32767:
                                        if (v56 >= (^(v56))) {
                                            if v2 {
                                                v55 += v55
                                                v55 = genFn0()
                                                var v72 uint = (v56 - 4294967295)
                                                _ = v72
                                            } else {
                                                fmt.Println(v2, v55, fi51)
                                            }
                                        }
                                    case -986:
                                        if (false && ((3.14159 * -2.25) != 0.1)) {
                                            if true {
                                                var v73 uint16 = 151
                                                var v74 int64 = (v4 / ((1) | 1))
                                                _ = v73
                                                _ = v74
                                            } else {
                                                if v2 {
                                                    _ = len(v3)
                                                    switch (fi51 ^ fi51) {
                                                    case -1:
                                                        switch "a" {
                                                        case "hello":
                                                            var v75 uint64 = v1
                                                            _ = v75
                                                        default:
                                                            if len(v3) > 0 {
                                                                v3[0] = ("a" + v67)
                                                            }
                                                        }
                                                        _ = len(v3)
                                                    case -2147483648:
                                                        v76 := GenS10{F0: ((true || v2) && v2), F1: v67}
                                                        _ = v76
                                                    case 1:
                                                        v4 = v4
                                                    default:
                                                        fmt.Println(v4)
                                                    }
                                                } else {
                                                    var v77 int64 = v4
                                                    switch ("a" + v67) {
                                                    case "":
                                                        if (v67 != v67) {
                                                            switch (fi53 << (((v56 ^ 42)) & 63)) {
                                                            case 1:
                                                                switch v67 {
                                                                case "hello":
                                                                    if v2 {
                                                                        var v78 uint16 = 1
                                                                        v3 = append(v3, v67)
                                                                        _ = len(v3)
                                                                        _ = v78
                                                                    } else {
                                                                        fmt.Println(fi53, fi51)
                                                                    }
                                                                }
                                                            case 0:
                                                                v55 = genFn0()
                                                            case 629:
                                                                var v79 uint32 = 2
                                                                _ = v79
                                                            }
                                                        } else {
                                                            if false {
                                                                v67 = "a"
                                                            }
                                                            switch (fi51 & genFn1()) {
                                                            case -2147483648:
                                                                var v80 uint8 = 1
                                                                _ = v80
                                                            }
                                                        }
                                                    case "hello":
                                                        v2 = (genFn1() != (-(fi51)))
                                                    }
                                                    _ = v77
                                                }
                                                v81 := GenS11{F0: v1}
                                                _ = v81
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        } else {
                                            var v82 int = int(v1)
                                            _ = v82
                                        }
                                    case 32768:
                                        if len(v3) > 0 {
                                            v3[0] = v67
                                        }
                                    }
                                    var v83 uint8 = 0
                                    _ = v55
                                    _ = v56
                                    _ = v67
                                    _ = v68
                                    _ = v83
                                case "hello":
                                    var v84 uint8 = 51
                                    _ = v84
                                default:
                                    v85 := GenS1{F0: 42, F1: (-(fi51))}
                                    _ = v85
                                }
                            }
                            v4 = (v4 ^ (v4 | (v4 % (((v4 << ((1) & 63))) | 1))))
                            v4 = -437
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v1, fi53)
                        }
                    } else {
                        _ = len(v3)
                    }
                }
                var v86 int = 1
                _ = fi53
                _ = v86
            }
            fmt.Println(v2)
            _ = fi51
        }
        var v87 uint32 = 573
        var v88 int64 = v4
        _ = v87
        _ = v88
    } else {
        v3 = append(v3, "\n\t")
    }
    for fi89 := 0; fi89 < 8; fi89++ {
        var v90 uint64 = (v1 % (((v1 % (((v1 ^ (v1 ^ (v1 & (v1 ^ v1))))) | 1))) | 1))
        _ = fi89
        _ = v90
    }
    fmt.Println(v1, v4)
    if v2 {
        var v91 uint = 42
        fmt.Println(v4)
        _ = v91
    }
    var v92 uint16 = 0
    var v93 int8 = 0
    switch ("hello" + "") {
    case "":
        var v94 uint32 = 752
        _ = v94
    }
    var v95 float64 = 1e10
    v3 = append(v3, "foo bar")
    var v96 int64 = (v4 / ((v4) | 1))
    v2 = v2
    for fi97 := 0; fi97 < 2; fi97++ {
        if (!(!(-1 < 42))) {
            v3 = append(v3, ("foo bar" + "\n\t"))
            var v98 uint16 = v92
            for fi99 := 0; fi99 < 20; fi99++ {
                v2 = v2
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi99
            }
            _ = v98
        }
        v2 = (!v2)
        v100 := GenS9{F0: 42, F1: genFn0()}
        _ = fi97
        _ = v100
    }
    if (2 >= 65535) {
        v101 := GenS1{F0: 783, F1: 1}
        if (25 <= 0) {
            if (v2 && true) {
                v96 = 0
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v95 += 0.0
            } else {
                var v102 int32 = 2147483647
                var v103 uint16 = 165
                var v104 int8 = (v93 << ((189) & 63))
                _ = v102
                _ = v103
                _ = v104
            }
            var v105 uint32 = 2147483648
            fmt.Println(v95, v95, v105)
            _ = v105
        } else {
            fmt.Println(v2, v2)
            switch -1 {
            case -670:
                v95 += v95
            case 128:
                switch "a" {
                case "\n\t":
                    var v106 string = "a"
                    _ = v106
                case "a":
                    if ((v2 && v2) && true) {
                        var v107 uint16 = v92
                        for fi108 := 0; fi108 < 15; fi108++ {
                            if (true && v2) {
                                if v2 {
                                    if ("" == ("foo bar" + "a")) {
                                        for fi109 := 0; fi109 < 19; fi109++ {
                                            fmt.Println(v1)
                                            _ = fi109
                                        }
                                        var v110 int = (fi108 ^ (^(fi108)))
                                        _ = v110
                                    }
                                    v3 = append(v3, "\n\t")
                                    v1 -= 402
                                }
                                if v2 {
                                    var v111 uint8 = 52
                                    _ = v111
                                } else {
                                    var v112 float32 = (-2.25 / -2.25)
                                    for fi113 := 0; fi113 < 19; fi113++ {
                                        if (1 == 32767) {
                                            switch "a" {
                                            case "":
                                                var v114 uint16 = v107
                                                _ = v114
                                            case "hello":
                                                fmt.Println(v112, v107)
                                            case "foo bar":
                                                v93 = (v93 | (-(v93)))
                                            default:
                                                v115 := GenS4{F0: -32768}
                                                _ = v115
                                            }
                                        }
                                        v107 = 214
                                        v92 = (v107 & 616)
                                        _ = len(v3)
                                        v107 = (v107 - 32767)
                                        _ = fi113
                                    }
                                    v107 += v92
                                    _ = v112
                                }
                            } else {
                                for fi116 := 0; fi116 < 6; fi116++ {
                                    if v2 {
                                        v4 += v4
                                        v117 := GenS3{F0: (-(v93)), F1: "\n\t"}
                                        var v118 uint32 = 0
                                        if (v118 <= v118) {
                                            v119 := GenS11{F0: (v1 >> ((858) & 63))}
                                            _ = v119
                                        }
                                        _ = v117
                                        _ = v118
                                    } else {
                                        fmt.Println(fi108, v92)
                                    }
                                    fmt.Println(v96, v95)
                                    var v120 uint = 822
                                    var v121 string = ""
                                    _ = fi116
                                    _ = v120
                                    _ = v121
                                }
                                switch int(v4) {
                                case -867:
                                    if v2 {
                                        v1 -= (v1 >> ((2) & 63))
                                        if v2 {
                                            var v122 int32 = -2147483648
                                            _ = v122
                                        } else {
                                            v123 := GenS1{F0: -855, F1: fi108}
                                            v95 += (((v95 + v95) - 1e10) + 3.14159)
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                            _ = v123
                                        }
                                        v95 = v95
                                    }
                                case 32768:
                                    if (591 <= 42) {
                                        for fi124 := 0; fi124 < 6; fi124++ {
                                            fmt.Println(v93, v107, v96)
                                            _ = fi124
                                        }
                                        v125 := GenS4{F0: 32767}
                                        _ = v125
                                    } else {
                                        fmt.Println(v107, fi108, v92)
                                    }
                                case 538:
                                    for fi126 := 0; fi126 < 16; fi126++ {
                                        var v127 uint32 = 1
                                        if true {
                                            v93 -= (v93 / ((1) | 1))
                                            var v128 int16 = -960
                                            var v129 float64 = -2.25
                                            _ = v128
                                            _ = v129
                                        } else {
                                            switch (fi126 ^ fi126) {
                                            case 1:
                                                v93 = v93
                                            case 0:
                                                var v130 int8 = -128
                                                _ = v130
                                            default:
                                                v96 -= (-(v4))
                                            }
                                            var v131 uint64 = (v1 / ((genFn0()) | 1))
                                            v3 = append(v3, "hello")
                                            v127 -= 0
                                            _ = v131
                                        }
                                        v132 := GenS12{F0: -2147483648, F1: ("\n\t" + "\n\t")}
                                        _ = fi126
                                        _ = v127
                                        _ = v132
                                    }
                                default:
                                    v96 -= v96
                                }
                                v133 := GenS4{F0: 32767}
                                _ = v133
                            }
                            var v134 uint32 = 2
                            _ = fi108
                            _ = v134
                        }
                        _ = v107
                    }
                default:
                    var v135 float32 = 0.0
                    _ = v135
                }
            case 32767:
                var v136 int = -1
                _ = v136
            }
            var v137 uint16 = 32767
            _ = v137
        }
        if (60 <= v93) {
            v92 = v92
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = v101
    } else {
        for fi138 := 0; fi138 < 3; fi138++ {
            if (!(v2 || false)) {
                var v139 uint32 = 694
                _ = v139
            }
            _ = fi138
        }
        v3 = append(v3, "")
        if false {
            if v2 {
                var v140 string = "foo bar"
                v141 := GenS13{F0: 2, F1: "hello", F2: -1}
                _ = v140
                _ = v141
            } else {
                v4 = (v96 | (^(v4)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v142 := GenS10{F0: (((false || (!(!(!v2)))) || v2) && v2), F1: ((("hello" + (("" + "\n\t") + "a")) + "hello") + "foo bar")}
                _ = v142
            }
        } else {
            for fi143 := 0; fi143 < 20; fi143++ {
                v2 = v2
                for fi144 := 0; fi144 < 4; fi144++ {
                    var v145 float64 = float64((0.0 / -2.25))
                    var v146 int16 = 42
                    var v147 int = fi143
                    _ = fi144
                    _ = v145
                    _ = v146
                    _ = v147
                }
                _ = fi143
            }
            for fi148 := 0; fi148 < 14; fi148++ {
                var v149 int8 = (v93 / (((v93 ^ v93)) | 1))
                v96 += (v96 ^ v96)
                _ = fi148
                _ = v149
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v92
    _ = v93
    _ = v95
    _ = v96
}
