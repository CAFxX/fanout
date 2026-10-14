package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 string; F1 bool; F2 uint }
type GenS9 struct { F0 int16; F1 int8; F2 int64 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS12 struct { F0 int64; F1 uint }
type GenS13 struct { F0 int8; F1 uint8 }
type GenS14 struct { F0 bool; F1 int8 }
func genFn0() uint32 {
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 uint8 = 163
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: (!true), F1: 9223372036854775807}
        switch ("hello" + "") {
        case "":
            if true {
                var v10 uint = 0
                for fi11 := 0; fi11 < 10; fi11++ {
                    v10 += 18446744073709551615
                    _ = fi11
                }
                _ = v10
            } else {
                var v12 float32 = 0.0
                if false {
                    v12 += v12
                    v13 := GenS1{F0: true, F1: -1}
                    _ = v13
                }
                v12 += (v12 * 1.5)
                _ = v12
            }
        case "hello":
            if true {
                var v14 string = (("foo bar" + "foo bar") + "hello")
                _ = v14
            }
        default:
            var v15 string = ("\n\t" + "hello")
            _ = v15
        }
        _ = v9
    } else {
        var v16 int = -251
        v16 += 32768
        v17 := GenS1{F0: (false || false), F1: 0}
        _ = v16
        _ = v17
    }
    var v18 float32 = (-2.25 * float32(float64(0.1)))
    v18 += v18
    if true {
        fmt.Println(v18)
        fmt.Println(v18)
    } else {
        var v19 uint8 = 255
        var v20 int8 = 127
        _ = v19
        _ = v20
    }
    _ = v18
    return 2147483647
}

func main() {
    var v1 float32 = 0.1
    var v2 uint8 = 0
    var v3 int16 = 32767
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]int)
    v5[(^(v4))] = (42 >= 65535)
    fmt.Println(v2, v2)
    if (65535 < 45) {
        fmt.Println(v3, v4)
        var v21 uint16 = 65535
        var v22 uint8 = (v2 << ((42) & 63))
        _ = v21
        _ = v22
    }
    delete(v5, 42)
    var v23 float32 = float32(0.0)
    v23 = v1
    v3 = v3
    switch 128 {
    case 277:
        switch 32768 {
        case 0:
            v2 += 2
        default:
            var v24 uint32 = (v4 / ((genFn0()) | 1))
            _ = v24
        }
    case 311:
        v23 -= v23
    }
    fmt.Println(v1)
    if true {
        switch "foo bar" {
        case "\n\t":
            for fi25 := 0; fi25 < 10; fi25++ {
                var v26 int = fi25
                if true {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 bool = (893 <= 0)
                        v1 = v1
                        _ = fi27
                        _ = v28
                    }
                    var v29 uint = 0
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v29 = (v29 << ((v29) & 63))
                        var v31 uint32 = (v4 & v4)
                        _ = fi30
                        _ = v31
                    }
                    _ = v29
                } else {
                    if true {
                        v23 += float32(3.14159)
                        if (873 >= 1) {
                            v32 := []int{(fi25 ^ -2147483648), (fi25 / (((v26 >> ((65535) & 63))) | 1)), v26}
                            if false {
                                var v33 int64 = -1
                                _ = v33
                            }
                            _ = v32
                        } else {
                            fmt.Println(v2)
                            v6[genFn0()] = v26
                            if (v3 >= v3) {
                                v34 := []int{}
                                var v35 uint64 = 720
                                _ = v34
                                _ = v35
                            }
                        }
                    }
                    v36 := v6[65535]
                    fmt.Println(v4, v26)
                    _ = v36
                }
                _ = fi25
                _ = v26
            }
        case "":
            switch "a" {
            case "":
                if (!(false && (65535 < 42))) {
                    v2 = 218
                    if (-339 > 42) {
                        _ = len(v5)
                    }
                }
            case "a":
                var v37 int16 = v3
                _ = v37
            default:
                if (-53 <= 40) {
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        for fi38 := 0; fi38 < 2; fi38++ {
                            if (3.14159 >= (v1 - (v23 + ((-2.25 + ((((v23 + ((v1 + v1) - 1e10)) / v23) - (v23 - float32(0.1))) + v1)) - v23)))) {
                                if false {
                                    if (((!(true && false)) || (false || true)) || (("\n\t" + "\n\t") == ("foo bar" + ""))) {
                                        v39 := GenS2{F0: v3, F1: 995, F2: (v3 % ((v3) | 1))}
                                        _ = v39
                                    } else {
                                        for fi40 := 0; fi40 < 15; fi40++ {
                                            v41 := []int{(-(fi38))}
                                            v2 += (v2 & 1)
                                            _ = fi40
                                            _ = v41
                                        }
                                    }
                                    v42 := GenS1{F0: (!true), F1: 730}
                                    var v43 uint32 = v4
                                    _ = v42
                                    _ = v43
                                } else {
                                    fmt.Println(v23, v1)
                                    if (!((65535 == 18446744073709551615) && false)) {
                                        if (0 < -305) {
                                            if true {
                                                for fi44 := 0; fi44 < 14; fi44++ {
                                                    var v45 int32 = 247
                                                    v6[589] = (fi44 ^ 864)
                                                    v2 += v2
                                                    v45 = v45
                                                    _ = fi44
                                                    _ = v45
                                                }
                                                var v46 float64 = 0.1
                                                if false {
                                                    fmt.Println(v46, v46, fi38)
                                                    var v47 int = fi38
                                                    _ = v47
                                                }
                                                _ = v46
                                            }
                                            switch "" {
                                            case "":
                                                fmt.Println(v3, v1, v4)
                                            case "\n\t":
                                                v48 := []int{(fi38 >> ((18446744073709551615) & 63))}
                                                _ = v48
                                            case "foo bar":
                                                var v49 int16 = (v3 - -954)
                                                _ = v49
                                            default:
                                                v4 -= genFn0()
                                            }
                                        }
                                        v3 = (v3 ^ (v3 % ((v3) | 1)))
                                    } else {
                                        var v50 int32 = 1
                                        for fi51 := 0; fi51 < 9; fi51++ {
                                            var v52 int32 = v50
                                            _ = fi51
                                            _ = v52
                                        }
                                        fmt.Println(v23)
                                        _ = v50
                                    }
                                }
                                for fi53 := 0; fi53 < 18; fi53++ {
                                    v54 := []int{(fi38 * (fi53 | (fi38 | int(v3)))), fi38}
                                    _ = fi53
                                    _ = v54
                                }
                            } else {
                                delete(v6, v4)
                                v1 = v1
                                v4 += uint32(v2)
                            }
                            var v55 string = "a"
                            v56 := []int{(fi38 | 32767), (fi38 * fi38), (fi38 ^ fi38)}
                            _ = fi38
                            _ = v55
                            _ = v56
                        }
                    default:
                        v57 := GenS1{F0: true, F1: -1}
                        _ = v57
                    }
                    v58 := GenS2{F0: (v3 >> ((18446744073709551615) & 63)), F1: 65535, F2: v3}
                    v59 := []int{42, -724, 399, 32767}
                    _ = v58
                    _ = v59
                } else {
                    switch 127 {
                    case 2147483647:
                        fmt.Println(v3, v4, v4)
                    case 32767:
                        v3 = v3
                    default:
                        v2 = (v2 + (v2 * v2))
                    }
                    var v60 bool = true
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v2 = (v2 + v2)
                        var v62 bool = (v60 && (!(-128 == 100)))
                        var v63 string = "\n\t"
                        _ = fi61
                        _ = v62
                        _ = v63
                    }
                    _ = v60
                }
            }
        case "hello":
            fmt.Println(v2)
        }
        if false {
            v4 += (^(v4))
            if (true && (2 != 42)) {
                for fi64 := 0; fi64 < 14; fi64++ {
                    switch (fi64 << ((1) & 63)) {
                    case 554:
                        if true {
                            var v65 int8 = -1
                            for fi66 := 0; fi66 < 19; fi66++ {
                                v23 += v23
                                fmt.Println(fi66)
                                delete(v6, (v4 + 2))
                                _ = fi66
                            }
                            if true {
                                if (("hello" + ("foo bar" + "")) > "a") {
                                    switch "\n\t" {
                                    case "a":
                                        v67 := GenS3{F0: 42, F1: (v4 + v4), F2: genFn0()}
                                        _ = v67
                                    }
                                    for fi68 := 0; fi68 < 3; fi68++ {
                                        if true {
                                            v1 = float32((3.14159 * float64(0.1)))
                                            if false {
                                                var v69 uint32 = v4
                                                _ = v69
                                            } else {
                                                v5[65535] = (v3 <= 0)
                                                v70 := GenS2{F0: (v3 & (v3 % ((v3) | 1))), F1: 319, F2: -1}
                                                switch (fi68 - (fi68 + int(v4))) {
                                                case 42:
                                                    v71 := GenS4{F0: 2, F1: 42}
                                                    _ = v71
                                                case -1:
                                                    var v72 uint16 = 194
                                                    _ = v72
                                                case 167:
                                                    _ = len(v5)
                                                }
                                                _ = v70
                                            }
                                        } else {
                                            fmt.Println(v2)
                                        }
                                        _ = fi68
                                    }
                                    v4 = genFn0()
                                }
                            } else {
                                if false {
                                    if false {
                                        if false {
                                            v73 := GenS3{F0: 42, F1: 4294967295, F2: 2147483647}
                                            for fi74 := 0; fi74 < 5; fi74++ {
                                                var v75 uint32 = (v4 ^ 1)
                                                v1 = v1
                                                var v76 uint = 750
                                                var v77 uint = (v76 + v76)
                                                _ = fi74
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            }
                                            v78 := []int{}
                                            _ = v73
                                            _ = v78
                                        } else {
                                            var v79 bool = ((v3 << ((1) & 63)) < v3)
                                            var v80 uint = 841
                                            if v79 {
                                                var v81 uint8 = uint8(v80)
                                                var v82 uint32 = v4
                                                v83 := []int{}
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            } else {
                                                var v84 string = (("\n\t" + "") + ("foo bar" + ((("foo bar" + "hello") + "foo bar") + "a")))
                                                v85 := GenS3{F0: -454, F1: v4, F2: genFn0()}
                                                var v86 string = "foo bar"
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            _ = v79
                                            _ = v80
                                        }
                                        for fi87 := 0; fi87 < 18; fi87++ {
                                            var v88 int8 = v65
                                            if true {
                                                v89 := []int{fi87, fi64, (-(fi64)), (fi87 ^ (fi87 % ((fi64) | 1)))}
                                                switch fi87 {
                                                case 146:
                                                    v89 = append(v89, 2147483647)
                                                case 0:
                                                    v3 = v3
                                                default:
                                                    v2 -= (v2 << ((42) & 63))
                                                }
                                                if true {
                                                    var v90 uint8 = (v2 + v2)
                                                    _ = v90
                                                }
                                                _ = v89
                                            } else {
                                                var v91 string = ("" + "foo bar")
                                                if true {
                                                    v92 := []int{(fi87 & fi64), 32767, (-(fi64)), -2147483648}
                                                    _ = v92
                                                } else {
                                                    var v93 string = ""
                                                    switch "hello" {
                                                    case "foo bar":
                                                        fmt.Println(v93, fi87)
                                                    }
                                                    v2 = v2
                                                    v88 += -1
                                                    if true {
                                                        var v94 uint16 = 2
                                                        if (v93 != v93) {
                                                            var v95 int8 = v65
                                                            _ = v95
                                                        }
                                                        _ = v94
                                                    }
                                                    switch (fi87 << ((317) & 63)) {
                                                    case 0:
                                                        if (!true) {
                                                            switch v91 {
                                                            case "hello":
                                                                v91 = "hello"
                                                            default:
                                                                fmt.Println(v91, fi64, v2)
                                                            }
                                                        } else {
                                                            v96 := GenS5{F0: (v2 * v2)}
                                                            if false {
                                                                v2 = (v2 >> ((0) & 63))
                                                                var v97 int32 = 0
                                                                v98 := GenS1{F0: (false && (v97 > v97)), F1: -689}
                                                                var v99 int8 = (-(v88))
                                                                _ = v97
                                                                _ = v98
                                                                _ = v99
                                                            } else {
                                                                var v100 string = "a"
                                                                v88 = v65
                                                                v101 := GenS6{F0: 255, F1: 2147483647}
                                                                _ = v100
                                                                _ = v101
                                                            }
                                                            v102 := GenS1{F0: false, F1: -494}
                                                            _ = v102
                                                            _ = v96
                                                        }
                                                    case 2147483647:
                                                        v6[genFn0()] = (fi64 + 4294967295)
                                                        switch fi87 {
                                                        case 127:
                                                            if true {
                                                                var v103 int16 = v3
                                                                var v104 int32 = 42
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                switch -1 {
                                                                case -78:
                                                                    v105 := GenS1{F0: ((v93 + v93) >= v93), F1: -9223372036854775808}
                                                                    _ = v105
                                                                default:
                                                                    switch v93 {
                                                                    case "hello":
                                                                        v106 := GenS5{F0: v2}
                                                                        _ = v106
                                                                    case "":
                                                                        var v107 uint32 = v4
                                                                        _ = v107
                                                                    }
                                                                }
                                                                var v108 uint8 = v2
                                                                fmt.Println(v65)
                                                                _ = v108
                                                            }
                                                        case 829:
                                                            if true {
                                                                var v109 int16 = -91
                                                                _ = v109
                                                            } else {
                                                                var v110 int32 = 1
                                                                v4 = v4
                                                                v91 = ("foo bar" + (("foo bar" + v93) + v91))
                                                                _ = v110
                                                            }
                                                        case -580:
                                                            fmt.Println(v3)
                                                        default:
                                                            if false {
                                                                v93 = v91
                                                                v3 = (v3 - (v3 % (((v3 << ((1) & 63))) | 1)))
                                                            }
                                                            v111 := []int{}
                                                            v3 = (v3 * v3)
                                                            var v112 uint = 42
                                                            _ = v111
                                                            _ = v112
                                                        }
                                                    default:
                                                        if false {
                                                            v113 := []int{fi64, fi87, -130}
                                                            fmt.Println(v3)
                                                            _ = v113
                                                        }
                                                    }
                                                    v4 -= genFn0()
                                                    var v114 int64 = -9223372036854775808
                                                    v114 = (^(v114))
                                                    _ = v114
                                                    _ = v93
                                                }
                                                if (115 <= 0) {
                                                    v23 = (-2.25 + 1.5)
                                                }
                                                _ = v91
                                            }
                                            v115 := GenS7{F0: v88}
                                            _ = fi87
                                            _ = v115
                                            _ = v88
                                        }
                                    }
                                    v116 := GenS5{F0: 255}
                                    if ((v3 + v3) <= v3) {
                                        for fi117 := 0; fi117 < 9; fi117++ {
                                            _ = len(v6)
                                            _ = fi117
                                        }
                                        if (!true) {
                                            if (236 == 42) {
                                                v118 := GenS8{F0: "foo bar", F1: (int32(v3) >= 42), F2: 18446744073709551615}
                                                _ = v118
                                            }
                                        }
                                        if true {
                                            switch ((("a" + "") + "a") + "foo bar") {
                                            case "foo bar":
                                                v23 -= (v23 * (float32(3.14159) + 0.1))
                                            case "\n\t":
                                                v119 := []int{}
                                                _ = v119
                                            case "a":
                                                fmt.Println(v23, v2)
                                            }
                                        } else {
                                            v1 = float32(-2.25)
                                        }
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            v5[v4] = (((v4 <= uint32(v2)) || true) && false)
                                        }
                                        fmt.Println(fi64, v23, fi64)
                                        for fi120 := 0; fi120 < 16; fi120++ {
                                            v2 -= (v2 | v2)
                                            var v121 string = "hello"
                                            _ = fi120
                                            _ = v121
                                        }
                                    }
                                    _ = v116
                                } else {
                                    fmt.Println(v3)
                                    if ((v65 << ((4294967295) & 63)) != v65) {
                                        v122 := []int{fi64, 128, fi64, fi64}
                                        v65 -= int8(fi64)
                                        _ = v122
                                    }
                                }
                            }
                            _ = v65
                        }
                    case -1:
                        for fi123 := 0; fi123 < 1; fi123++ {
                            v124 := GenS6{F0: 2, F1: int32(v2)}
                            _ = fi123
                            _ = v124
                        }
                    default:
                        if true {
                            for fi125 := 0; fi125 < 11; fi125++ {
                                fmt.Println(fi64, v1)
                                if false {
                                    var v126 uint8 = v2
                                    var v127 bool = (0 >= v3)
                                    v128 := []int{fi125, (fi64 % ((fi64) | 1)), (fi64 << ((727) & 63))}
                                    _ = v126
                                    _ = v127
                                    _ = v128
                                } else {
                                    var v129 float64 = ((0.0 / 1.0) * (3.14159 / 3.14159))
                                    switch "hello" {
                                    case "hello":
                                        var v130 bool = true
                                        _ = v130
                                    case "a":
                                        v131 := GenS7{F0: 0}
                                        _ = v131
                                    case "foo bar":
                                        v2 -= v2
                                    }
                                    var v132 float32 = 3.14159
                                    _ = v129
                                    _ = v132
                                }
                                if ("hello" >= "a") {
                                    var v133 int64 = 0
                                    if (genFn0() <= v4) {
                                        v134 := GenS1{F0: false, F1: (v133 >> ((0) & 63))}
                                        var v135 uint8 = (v2 * 128)
                                        v136 := GenS9{F0: (v3 / ((v3) | 1)), F1: 100, F2: v133}
                                        _ = v134
                                        _ = v135
                                        _ = v136
                                    }
                                    var v137 int16 = v3
                                    _ = v133
                                    _ = v137
                                } else {
                                    v138 := []int{}
                                    var v139 int32 = 663
                                    _ = v138
                                    _ = v139
                                }
                                _ = fi125
                            }
                            var v140 string = ("foo bar" + ("\n\t" + ""))
                            _ = v140
                        } else {
                            if false {
                                var v141 int8 = 100
                                _ = v141
                            }
                            if (((-(fi64)) != -2147483648) || true) {
                                v4 += 4294967295
                                var v142 uint = 255
                                _ = v142
                            }
                        }
                    }
                    v1 += ((v1 / (float32(0.0) * ((v23 + v1) - 0.0))) - (float32((1e10 / 1.0)) / ((float32(0.0) * v23) + 0.0)))
                    if true {
                        v143 := GenS10{F0: -246, F1: ("a" + "foo bar")}
                        _ = v143
                    }
                    _ = fi64
                }
                var v144 uint64 = 408
                var v145 uint64 = 557
                _ = v144
                _ = v145
            } else {
                if false {
                    _ = len(v6)
                } else {
                    if (-1 >= -1) {
                        fmt.Println(v4, v3)
                        var v146 float64 = 0.0
                        _ = v146
                    } else {
                        var v147 string = "\n\t"
                        _ = v147
                    }
                    var v148 int = 0
                    _ = v148
                }
                var v149 uint64 = 2
                v150 := v5[(v4 | 2)]
                _ = v149
                _ = v150
            }
        } else {
            fmt.Println(v3, v1)
            v151 := GenS11{F0: 2, F1: 949, F2: (v2 & v2)}
            _ = len(v6)
            _ = v151
        }
    }
    v152 := GenS12{F0: 9223372036854775807, F1: 599}
    v153 := GenS13{F0: 0, F1: (v2 % (((v2 + v2)) | 1))}
    if (!true) {
        var v154 uint64 = 1
        _ = v154
    }
    v155 := GenS14{F0: true, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v152
    _ = v153
    _ = v155
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
