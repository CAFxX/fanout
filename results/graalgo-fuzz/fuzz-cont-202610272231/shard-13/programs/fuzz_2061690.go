package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 int8; F1 int64; F2 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 int64; F1 int8 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint64; F1 int16; F2 int }
type GenS16 struct { F0 int32; F1 uint16; F2 int16 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 uint; F1 uint64 }
type GenS19 struct { F0 int8; F1 string; F2 uint64 }
type GenS20 struct { F0 uint16; F1 uint32; F2 string }
type GenS21 struct { F0 int; F1 string; F2 uint64 }
type GenS22 struct { F0 int32; F1 uint64; F2 bool }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 uint64; F1 uint64 }
type GenS25 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS26 struct { F0 uint64; F1 int16 }
func genFn0(p0 int64, p1 int8) int {
    var v6 string = "a"
    var v7 int64 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 2147483647
}
func genFn1(p0 string, p1 int16, p2 bool) uint8 {
    fmt.Println(p1, p1, p2)
    var v8 int64 = 0
    p2 = p2
    p0 = ""
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}

func main() {
    var v1 int8 = 100
    var v2 int = -740
    var v3 uint64 = 99
    v4 := make(map[int]uint32)
    var v5 uint = 115
    v3 += v3
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v5, v1, v1)
            v10 := GenS1{F0: 0, F1: 1}
            _ = fi9
            _ = v10
        }
        fmt.Println(v2)
    }
    var v11 uint = 1
    var v12 uint64 = 1
    v11 = (^(v11))
    v2 = v2
    switch "foo bar" {
    case "":
        delete(v4, -810)
    case "hello":
        for fi13 := 0; fi13 < 9; fi13++ {
            v5 -= v11
            switch "\n\t" {
            case "\n\t":
                v14 := GenS2{F0: (v11 % (((v11 & 255)) | 1))}
                _ = v14
            case "":
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 int16 = -1
                    if false {
                        if true {
                            switch (("hello" + ("\n\t" + "foo bar")) + "foo bar") {
                            case "a":
                                fmt.Println(v11, v16)
                            case "":
                                if false {
                                    if true {
                                        if (false && true) {
                                            switch "\n\t" {
                                            case "hello":
                                                var v17 uint32 = 155
                                                _ = v17
                                            case "":
                                                v3 = v3
                                            }
                                            if false {
                                                switch v2 {
                                                case 0:
                                                    v18 := GenS3{F0: uint(v3), F1: (v16 & (v16 & (v16 + (v16 - v16))))}
                                                    _ = v18
                                                case 2147483647:
                                                    if ((0.0 / 1.5) >= float64(float32(1e10))) {
                                                        v19 := GenS4{F0: 0, F1: 169}
                                                        switch "" {
                                                        case "hello":
                                                            fmt.Println(fi15, v11, v1)
                                                        default:
                                                            if false {
                                                                var v20 int = (fi15 % ((genFn0(0, v1)) | 1))
                                                                _ = v20
                                                            }
                                                        }
                                                        v12 += (^(v12))
                                                        _ = v19
                                                    } else {
                                                        fmt.Println(v2, fi13)
                                                        v16 = 32767
                                                        fmt.Println(v16, v11, v12)
                                                    }
                                                case -1:
                                                    var v21 float32 = 1e10
                                                    _ = v21
                                                }
                                            } else {
                                                var v22 int = v2
                                                _ = v22
                                            }
                                        } else {
                                            fmt.Println(v12)
                                            switch (v2 / (((fi15 >> (((v11 % ((v5) | 1))) & 63))) | 1)) {
                                            case 1:
                                                _ = len(v4)
                                            case 0:
                                                if (true && true) {
                                                    v1 = 0
                                                    v3 = (v3 - v12)
                                                    var v23 uint32 = 65535
                                                    _ = v23
                                                } else {
                                                    if (false || false) {
                                                        v12 = v3
                                                        var v24 int64 = 0
                                                        _ = v24
                                                    } else {
                                                        v16 += (^(v16))
                                                        v25 := []int{(v2 | fi15), v2, 1, -2147483648}
                                                        _ = v25
                                                    }
                                                }
                                            case 32767:
                                                var v26 bool = (((v11 / (((^(v5))) | 1)) != uint(v12)) && true)
                                                _ = v26
                                            default:
                                                var v27 uint = v11
                                                _ = v27
                                            }
                                            v28 := GenS5{F0: 616, F1: (v3 & v12), F2: -32768}
                                            _ = v28
                                        }
                                    } else {
                                        v29 := []int{v2, int(v3)}
                                        switch "a" {
                                        case "hello":
                                            var v30 int8 = 28
                                            _ = v30
                                        case "foo bar":
                                            if true {
                                                v4[(v2 + fi15)] = 723
                                                switch (("hello" + ((("\n\t" + "\n\t") + "") + "")) + "") {
                                                case "":
                                                    if (!(true || true)) {
                                                        var v31 float32 = ((0.1 - float32(0.1)) * float32(0.0))
                                                        _ = v31
                                                    }
                                                case "\n\t":
                                                    if true {
                                                        v11 += (v5 / ((2) | 1))
                                                        fmt.Println(v5, v1)
                                                        var v32 int64 = -9223372036854775808
                                                        _ = v32
                                                    }
                                                default:
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v29 = append(v29, (-(fi15)))
                                                    case "foo bar":
                                                        v5 += v5
                                                    default:
                                                        var v33 int64 = -229
                                                        _ = v33
                                                    }
                                                }
                                            } else {
                                                var v34 int = (fi13 * 1)
                                                fmt.Println(v12)
                                                _ = v34
                                            }
                                        case "a":
                                            v35 := GenS1{F0: 551, F1: 788}
                                            _ = v35
                                        }
                                        var v36 int64 = 0
                                        _ = v29
                                        _ = v36
                                    }
                                    var v37 uint32 = 0
                                    if false {
                                        fmt.Println(fi15, fi13, v11)
                                        fmt.Println(v16, v37, v37)
                                        v38 := []int{(fi15 * genFn0(0, v1)), (fi15 << (((^(v11))) & 63)), genFn0(1, int8(v5)), v2}
                                        if true {
                                            v39 := GenS6{F0: (true || false)}
                                            fmt.Println(v11, v2, fi13)
                                            fmt.Println(v3)
                                            _ = v39
                                        }
                                        v2 = genFn0(0, 42)
                                        _ = v38
                                    } else {
                                        if true {
                                            var v40 float32 = -2.25
                                            v11 = v5
                                            v1 += v1
                                            if false {
                                                fmt.Println(v3, v5)
                                            } else {
                                                var v41 int16 = v16
                                                v4[fi13] = (v37 ^ v37)
                                                switch ("a" + "foo bar") {
                                                case "a":
                                                    v42 := GenS7{F0: (("\n\t" + (("a" + "\n\t") + "hello")) + "foo bar"), F1: -114}
                                                    _ = v42
                                                case "":
                                                    v5 += 0
                                                case "foo bar":
                                                    switch 128 {
                                                    case 127:
                                                        v43 := GenS7{F0: "", F1: (v1 + 127)}
                                                        _ = v43
                                                    case -2147483648:
                                                        v44 := GenS7{F0: "", F1: -1}
                                                        _ = v44
                                                    }
                                                }
                                                _ = v41
                                            }
                                            _ = v40
                                        } else {
                                            var v45 float32 = (1e10 + (1.5 * 1e10))
                                            v45 += float32(0.0)
                                            v2 = fi13
                                            _ = v45
                                        }
                                        var v46 uint32 = (v37 + v37)
                                        _ = v46
                                    }
                                    _ = v37
                                }
                            default:
                                fmt.Println(fi13, fi15, v11)
                            }
                            v47 := GenS3{F0: v5, F1: int16(v11)}
                            _ = v47
                        }
                    }
                    v48 := GenS1{F0: 2, F1: -1}
                    _ = fi15
                    _ = v16
                    _ = v48
                }
            }
            v49 := GenS8{F0: (v1 / (((v1 ^ v1)) | 1)), F1: 9223372036854775807, F2: genFn0(-1, 42)}
            _ = fi13
            _ = v49
        }
    }
    fmt.Println(v11)
    var v50 uint8 = genFn1("foo bar", -229, (545 != uint16(v12)))
    for fi51 := 0; fi51 < 14; fi51++ {
        if (!true) {
            var v52 int8 = (v1 & (v1 ^ v1))
            if (65535 != uint32(v3)) {
                v53 := GenS6{F0: ((v11 / ((18446744073709551615) | 1)) > v5)}
                for fi54 := 0; fi54 < 15; fi54++ {
                    if false {
                        var v55 int32 = -461
                        _ = v55
                    } else {
                        var v56 int = v2
                        v57 := v4[(fi54 >> ((v11) & 63))]
                        v50 -= genFn1("hello", 905, (-239 <= 98))
                        if (42 == -2147483648) {
                            if (v57 < v57) {
                                v1 += v1
                                if (0 == v1) {
                                    if true {
                                        switch "" {
                                        case "foo bar":
                                            v58 := []int{v2, (v2 ^ genFn0(422, -1)), v56, -2147483648}
                                            _ = v58
                                        }
                                        v59 := GenS9{F0: int8(v5)}
                                        v60 := GenS10{F0: uint8(v11)}
                                        _ = v59
                                        _ = v60
                                    } else {
                                        if (("foo bar" + "foo bar") != "\n\t") {
                                            if false {
                                                var v61 float32 = 0.0
                                                fmt.Println(v56, fi51)
                                                _ = v61
                                            } else {
                                                v62 := GenS11{F0: (-(fi51))}
                                                _ = v62
                                            }
                                            v63 := GenS12{F0: 790, F1: 18446744073709551615}
                                            switch (("\n\t" + "") + "\n\t") {
                                            case "hello":
                                                switch ("hello" + "foo bar") {
                                                case "\n\t":
                                                    var v64 float64 = ((0.1 + (1e10 - float64(1.5))) + -2.25)
                                                    _ = v64
                                                case "a":
                                                    v65 := []int{int(v11), (-(fi54))}
                                                    _ = v65
                                                }
                                            case "a":
                                                v66 := []int{32767, (fi51 << ((v11) & 63))}
                                                _ = v66
                                            case "":
                                                v67 := []int{-1, (v2 | fi51), (fi54 - genFn0(-1, (v1 / ((v52) | 1))))}
                                                _ = v67
                                            default:
                                                var v68 float64 = 1e10
                                                _ = v68
                                            }
                                            _ = v63
                                        }
                                        var v69 uint = v11
                                        v3 = v12
                                        _ = v69
                                    }
                                    if true {
                                        var v70 uint64 = (v12 | v3)
                                        var v71 float32 = (-2.25 * ((0.0 * 3.14159) - (float32(0.1) * 3.14159)))
                                        v72 := GenS8{F0: 1, F1: -9223372036854775808, F2: (v2 / ((570) | 1))}
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                    v50 = 2
                                } else {
                                    v4[genFn0(1, -128)] = (v57 >> (((^(v11))) & 63))
                                    if true {
                                        v73 := GenS12{F0: -2147483648, F1: (^(v3))}
                                        v50 = 42
                                        v52 = v1
                                        _ = v73
                                    }
                                    v52 = 100
                                    var v74 uint32 = (v57 % ((165) | 1))
                                    _ = v74
                                }
                                v75 := GenS9{F0: (v52 / ((v1) | 1))}
                                _ = v75
                            }
                            v76 := []int{}
                            if false {
                                v77 := GenS13{F0: 831, F1: -127}
                                _ = v77
                            }
                            _ = v76
                        }
                        v78 := []int{(v2 ^ int(v11)), 1}
                        switch (fi54 % ((fi51) | 1)) {
                        case 2147483647:
                            fmt.Println(v52, v2, v12)
                        case 4294967295:
                            v3 += v12
                        default:
                            v50 -= (v50 - v50)
                        }
                        v79 := GenS14{F0: -2147483648}
                        var v80 int32 = -2147483648
                        if false {
                            if true {
                                var v81 uint8 = 232
                                v50 -= genFn1("hello", 181, ((v12 >= 42) && false))
                                v78 = append(v78, v56)
                                fmt.Println(v56, v12)
                                v82 := GenS15{F0: (v3 >> (((v5 - v5)) & 63)), F1: 1000, F2: (fi51 % ((v2) | 1))}
                                _ = v81
                                _ = v82
                            } else {
                                v83 := GenS16{F0: (v80 % (((^(v80))) | 1)), F1: 65535, F2: -1}
                                v78 = append(v78, int(v1))
                                fmt.Println(v5, v5, v57)
                                if false {
                                    v84 := GenS17{F0: 4294967295}
                                    var v85 uint64 = v12
                                    _ = v84
                                    _ = v85
                                }
                                _ = v83
                            }
                        } else {
                            delete(v4, genFn0(9223372036854775807, v1))
                        }
                        _ = v56
                        _ = v57
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    switch ("a" + "") {
                    case "\n\t":
                        var v86 bool = (!(false && true))
                        _ = v86
                    }
                    if false {
                        v11 = (^(v11))
                        v87 := []int{v2}
                        _ = v87
                    } else {
                        v1 += v52
                        switch "foo bar" {
                        case "":
                            v88 := GenS18{F0: 18446744073709551615, F1: v12}
                            _ = v88
                        case "a":
                            switch "hello" {
                            case "":
                                var v89 uint32 = 42
                                _ = v89
                            case "\n\t":
                                switch (((("hello" + "a") + (("a" + "hello") + "foo bar")) + "\n\t") + "") {
                                case "\n\t":
                                    switch genFn0(-574, v1) {
                                    case -1:
                                        v1 += v52
                                    default:
                                        var v90 int = (^(v2))
                                        _ = v90
                                    }
                                case "hello":
                                    fmt.Println(v52, fi51)
                                case "a":
                                    if (v5 >= v5) {
                                        v91 := []int{220}
                                        if (1.5 < 0.0) {
                                            var v92 int16 = 0
                                            var v93 int8 = v52
                                            var v94 string = ("" + ("" + ""))
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        }
                                        _ = v91
                                    } else {
                                        fmt.Println(v11, v12, v1)
                                        v95 := []int{(fi51 / (((fi54 / ((v2) | 1))) | 1)), fi51, v2}
                                        v95 = append(v95, genFn0(42, (v1 - -1)))
                                        var v96 int16 = -1
                                        fmt.Println(v1, fi54, fi54)
                                        _ = v95
                                        _ = v96
                                    }
                                }
                            case "foo bar":
                                if false {
                                    if false {
                                        v52 -= v1
                                        switch "\n\t" {
                                        case "a":
                                            v50 = genFn1("foo bar", 1000, ((!false) || (1 > 199)))
                                        case "foo bar":
                                            var v97 uint = (v5 / ((v11) | 1))
                                            if false {
                                                fmt.Println(v11, v3, v3)
                                                if true {
                                                    v97 = (v5 | v5)
                                                }
                                                var v98 string = "\n\t"
                                                _ = v98
                                            } else {
                                                if true {
                                                    v99 := GenS19{F0: v1, F1: "\n\t", F2: (v3 / (((v12 + 416)) | 1))}
                                                    v97 -= (v97 * 42)
                                                    _ = v99
                                                }
                                                v100 := GenS20{F0: 65535, F1: 4294967295, F2: "foo bar"}
                                                v101 := GenS21{F0: (fi51 / (((^(v2))) | 1)), F1: "", F2: (^(v12))}
                                                v102 := GenS14{F0: 0}
                                                v103 := GenS19{F0: v1, F1: "", F2: v12}
                                                v104 := []int{(fi54 % ((v2) | 1)), genFn0(-705, -96)}
                                                v104 = append(v104, (fi54 | (fi51 ^ v2)))
                                                if false {
                                                    var v105 int16 = 1
                                                    var v106 uint8 = v50
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    var v107 uint = v11
                                                    switch "foo bar" {
                                                    case "a":
                                                        if true {
                                                            var v108 uint = 255
                                                            v109 := GenS22{F0: 1, F1: v12, F2: ((("hello" == "\n\t") && (!(!(int(v12) >= fi51)))) || false)}
                                                            v110 := v4[v2]
                                                            _ = v108
                                                            _ = v109
                                                            _ = v110
                                                        }
                                                    }
                                                    fmt.Println(v97)
                                                    if ("foo bar" < ("" + (("\n\t" + "\n\t") + "foo bar"))) {
                                                        if (((1e10 * 0.1) - (3.14159 / 0.1)) >= 0.0) {
                                                            v111 := GenS23{F0: int8(v5)}
                                                            var v112 uint8 = (v50 & (v50 >> ((v11) & 63)))
                                                            var v113 int16 = -1
                                                            _ = v111
                                                            _ = v112
                                                            _ = v113
                                                        }
                                                        var v114 uint16 = 42
                                                        _ = len(v104)
                                                        _ = v114
                                                    } else {
                                                        if (0 < -2147483648) {
                                                            switch "a" {
                                                            case "a":
                                                                var v115 string = ("" + "foo bar")
                                                                v116 := GenS19{F0: v1, F1: v115, F2: v3}
                                                                _ = v115
                                                                _ = v116
                                                            case "\n\t":
                                                                var v117 int8 = -128
                                                                _ = v117
                                                            case "":
                                                                if true {
                                                                    v11 = v5
                                                                }
                                                            }
                                                            var v118 bool = false
                                                            v119 := GenS24{F0: v12, F1: 35}
                                                            _ = v118
                                                            _ = v119
                                                        }
                                                        v97 += (v107 + (v107 * 42))
                                                    }
                                                    var v120 int8 = (v1 ^ v1)
                                                    v5 -= v97
                                                    v2 = -2147483648
                                                    v104 = append(v104, (fi51 << (((v5 * (v5 | (v107 | v11)))) & 63)))
                                                    _ = v107
                                                    _ = v120
                                                }
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                                _ = v104
                                            }
                                            v50 -= uint8(v5)
                                            var v121 int32 = 1
                                            _ = v121
                                            _ = v97
                                        case "":
                                            v122 := GenS25{F0: 786, F1: uint32(v3), F2: 32767}
                                            _ = v122
                                        default:
                                            var v123 uint16 = 32767
                                            _ = v123
                                        }
                                    }
                                    v124 := GenS26{F0: v3, F1: 691}
                                    var v125 int16 = 0
                                    _ = v124
                                    _ = v125
                                } else {
                                    v126 := GenS21{F0: (fi54 % (((fi54 % ((genFn0(-558, v52)) | 1))) | 1)), F1: "\n\t", F2: v3}
                                    _ = v126
                                }
                            default:
                                v127 := []int{}
                                _ = v127
                            }
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 615:
                                fmt.Println(v3, v1, v1)
                            case -2147483648:
                                if false {
                                    v2 = genFn0(-1, v1)
                                    v50 = v50
                                    var v128 int16 = 447
                                    _ = v128
                                }
                            default:
                                var v129 float64 = ((float64((-2.25 + (1e10 / 1.5))) * -2.25) * -2.25)
                                _ = v129
                            }
                            v1 += (v52 & (v1 ^ (v1 >> ((v11) & 63))))
                            if true {
                                v1 -= 42
                            } else {
                                var v130 int16 = -334
                                v131 := []int{genFn0(-974, v52)}
                                var v132 uint16 = 0
                                _ = v130
                                _ = v131
                                _ = v132
                            }
                            var v133 uint8 = (v50 >> ((v11) & 63))
                            _ = v133
                        default:
                            var v134 float32 = 0.1
                            _ = v134
                        }
                        v3 = v12
                    }
                    _ = fi54
                }
                var v135 int16 = int16(v1)
                _ = v135
                _ = v53
            }
            _ = v52
        }
        _ = fi51
    }
    v11 -= v11
    var v136 uint16 = 1
    fmt.Println(v1)
    _ = len(v4)
    for fi137 := 0; fi137 < 10; fi137++ {
        var v138 int8 = v1
        v139 := GenS21{F0: -2147483648, F1: (("\n\t" + "a") + "hello"), F2: uint64(v50)}
        _ = fi137
        _ = v138
        _ = v139
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v12
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
}
