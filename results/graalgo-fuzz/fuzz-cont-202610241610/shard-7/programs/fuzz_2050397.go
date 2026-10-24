package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint; F1 int64; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 bool; F2 int }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 string; F1 int64; F2 uint }
type GenS12 struct { F0 uint64; F1 int16; F2 uint }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16; F1 int32 }
type GenS15 struct { F0 string; F1 uint32 }
type GenS16 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS17 struct { F0 uint; F1 uint32 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 uint }

func main() {
    var v1 uint8 = 192
    var v2 float64 = -2.25
    var v3 uint32 = 2
    var v4 bool = false
    var v5 uint = 2
    var v6 int8 = -9
    var v7 float64 = v2
    if (false || v4) {
        fmt.Println(v4, v5)
        v2 += ((3.14159 + v7) + v2)
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{fi8, (fi8 << (((v5 / ((v5) | 1))) & 63))}
            for fi10 := 0; fi10 < 17; fi10++ {
                var v11 bool = (!(597 >= int(v1)))
                _ = fi10
                _ = v11
            }
            if (((("" + "\n\t") + "a") != (((("foo bar" + "\n\t") + (("\n\t" + "foo bar") + "")) + "foo bar") + "a")) && false) {
                _ = len(v9)
            } else {
                switch (fi8 ^ (fi8 >> ((v5) & 63))) {
                case 0:
                    fmt.Println(v6, v3)
                default:
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                }
            }
            _ = fi8
            _ = v9
        }
    } else {
        var v12 float64 = 1e10
        for fi13 := 0; fi13 < 4; fi13++ {
            switch ("" + ("a" + "foo bar")) {
            case "\n\t":
                v14 := []int{fi13}
                _ = v14
            case "a":
                v2 = v7
            default:
                var v15 float64 = float64(0.0)
                _ = v15
            }
            v16 := GenS1{F0: 0, F1: 128}
            _ = fi13
            _ = v16
        }
        _ = v12
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 int64 = 290
        switch fi17 {
        case 127:
            for fi19 := 0; fi19 < 17; fi19++ {
                if (!false) {
                    fmt.Println(v4)
                } else {
                    v20 := []int{fi17, fi19, fi19}
                    v21 := GenS2{F0: uint64(v3), F1: v18, F2: -32768}
                    var v22 int16 = 1
                    _ = v20
                    _ = v21
                    _ = v22
                }
                if ((false && ((v6 >> ((v5) & 63)) < 0)) || ((v6 * (-(v6))) < v6)) {
                    if (!((v18 * 308) > int64(v1))) {
                        var v23 int64 = v18
                        var v24 string = "a"
                        var v25 int32 = -1
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        v26 := GenS1{F0: -976, F1: 79}
                        _ = v26
                    }
                } else {
                    var v27 uint64 = 4294967295
                    v28 := GenS1{F0: -1, F1: (v1 & 1)}
                    v27 += (v27 + v27)
                    _ = v27
                    _ = v28
                }
                if ((18446744073709551615 < 0) || (!(!false))) {
                    if true {
                        fmt.Println(v7, v2)
                        switch (-(fi17)) {
                        case 2147483647:
                            if v4 {
                                fmt.Println(v7)
                                if (v1 >= v1) {
                                    var v29 int32 = -2147483648
                                    fmt.Println(v18, fi17)
                                    _ = v29
                                } else {
                                    v30 := GenS3{F0: (-(fi17)), F1: 2, F2: -32768}
                                    _ = v30
                                }
                                if true {
                                    var v31 int16 = 0
                                    _ = v31
                                } else {
                                    if v4 {
                                        if false {
                                            v2 += ((1.5 * v7) - v7)
                                            v6 += v6
                                            v6 += v6
                                            if ((0.0 * (-2.25 - ((1.5 / float32(v2)) / (-2.25 / (float32(float64(-2.25)) * (-2.25 / (((((-2.25 - 0.1) * 0.1) / 1.5) * 0.1) - -2.25))))))) >= 3.14159) {
                                                if (false && false) {
                                                    switch (fi19 / ((0) | 1)) {
                                                    case -9223372036854775808:
                                                        v5 -= v5
                                                    case 0:
                                                        var v32 int = (fi19 / (((fi17 | (fi19 << ((18446744073709551615) & 63)))) | 1))
                                                        _ = v32
                                                    case 127:
                                                        v33 := []int{fi17, (fi19 ^ fi17), fi17, 32767}
                                                        _ = v33
                                                    default:
                                                        if (-744 < (v18 / (((-(v18))) | 1))) {
                                                            v34 := []int{fi17, -381, (-(fi17))}
                                                            _ = len(v34)
                                                            v18 = v18
                                                            _ = v34
                                                        }
                                                    }
                                                }
                                                v4 = (v4 || ((v5 - (v5 * (v5 & (v5 & 936)))) == 2))
                                                v35 := GenS4{F0: (v5 % ((v5) | 1)), F1: int64(v3), F2: (v6 | 78)}
                                                _ = v35
                                            } else {
                                                var v36 float64 = (float64((float32((-2.25 * v7)) * (-2.25 / (float32(v7) - 3.14159)))) - (0.1 * v2))
                                                v37 := []int{fi19, 0}
                                                _ = v36
                                                _ = v37
                                            }
                                            v2 = (v7 + v2)
                                        }
                                        if v4 {
                                            fmt.Println(fi17, v3)
                                            v38 := GenS5{F0: 1000}
                                            if (v7 >= v2) {
                                                v2 = (v2 - (v2 - 1.5))
                                                switch ("a" + "\n\t") {
                                                case "a":
                                                    v1 -= (^(v1))
                                                case "hello":
                                                    if false {
                                                        v18 = (v18 >> (((v5 % (((^(v5))) | 1))) & 63))
                                                        v39 := GenS6{F0: 65535, F1: v4, F2: (fi19 & (fi17 & (-(fi17))))}
                                                        _ = v39
                                                    } else {
                                                        var v40 int64 = (v18 | (v18 << (((v5 / ((v5) | 1))) & 63)))
                                                        _ = v40
                                                    }
                                                default:
                                                    v41 := GenS6{F0: 42, F1: (790 != 1), F2: -204}
                                                    var v42 int32 = -715
                                                    _ = v41
                                                    _ = v42
                                                }
                                            }
                                            _ = v38
                                        } else {
                                            if v4 {
                                                if v4 {
                                                    if v4 {
                                                        v43 := GenS3{F0: fi17, F1: (v1 + 42), F2: 1}
                                                        switch fi17 {
                                                        case -1:
                                                            var v44 int8 = (v6 + v6)
                                                            _ = v44
                                                        case 42:
                                                            v5 = (v5 * 231)
                                                        default:
                                                            v45 := []int{(fi19 * (fi17 - int(v6))), fi19, fi17}
                                                            switch "hello" {
                                                            case "":
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v7 += (3.14159 + -2.25)
                                                                }
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                            _ = v45
                                                        }
                                                        v46 := GenS7{F0: 0, F1: v1}
                                                        if true {
                                                            var v47 int32 = -1
                                                            v18 -= 316
                                                            _ = v47
                                                        } else {
                                                            if false {
                                                                var v48 uint = v5
                                                                _ = v48
                                                            }
                                                            switch (fi19 + (fi17 % ((fi19) | 1))) {
                                                            case 581:
                                                                if (!((v4 && v4) && (("" + "") <= ""))) {
                                                                    v49 := []int{}
                                                                    _ = v49
                                                                } else {
                                                                    v50 := GenS8{F0: v1}
                                                                    _ = v50
                                                                }
                                                            case 9223372036854775807:
                                                                var v51 int64 = -1
                                                                v7 -= v2
                                                                v5 = (^(v5))
                                                                var v52 int16 = -1
                                                                if (0 >= 32767) {
                                                                    v53 := GenS9{F0: 4294967295, F1: 4294967295}
                                                                    v54 := []int{fi19, fi17, fi17}
                                                                    v51 = 9223372036854775807
                                                                    v51 -= (v18 + v18)
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    v55 := GenS9{F0: 0, F1: fi17}
                                                                    switch (fi17 << ((v5) & 63)) {
                                                                    case 32767:
                                                                        fmt.Println(v51, v1)
                                                                    }
                                                                    v5 = 255
                                                                    v3 -= 42
                                                                    _ = v55
                                                                }
                                                                _ = v51
                                                                _ = v52
                                                            case -550:
                                                                var v56 uint16 = 2
                                                                _ = v56
                                                            default:
                                                                var v57 float32 = 1.5
                                                                _ = v57
                                                            }
                                                        }
                                                        if ((v4 && (v3 >= (v3 >> ((296) & 63)))) || v4) {
                                                            v58 := GenS2{F0: 0, F1: -54, F2: -830}
                                                            var v59 string = ((("a" + "hello") + ("\n\t" + (("hello" + "a") + ("\n\t" + "\n\t")))) + (("a" + "foo bar") + "hello"))
                                                            v59 = v59
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        _ = v43
                                                        _ = v46
                                                    }
                                                }
                                                v7 = v2
                                            } else {
                                                var v60 int64 = v18
                                                var v61 uint8 = 83
                                                _ = v60
                                                _ = v61
                                            }
                                            if ((32767 == -32768) && v4) {
                                                switch (-(fi19)) {
                                                case -1:
                                                    if v4 {
                                                        var v62 uint64 = 4294967295
                                                        var v63 int64 = 0
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                case 61:
                                                    var v64 uint = (v5 >> ((28) & 63))
                                                    _ = v64
                                                case -245:
                                                    v1 += (^(v1))
                                                default:
                                                    var v65 uint = (v5 + (v5 & (v5 / (((^(v5))) | 1))))
                                                    _ = v65
                                                }
                                                v66 := GenS10{F0: "a", F1: int8(v5)}
                                                v67 := []int{fi19}
                                                _ = v66
                                                _ = v67
                                            } else {
                                                v68 := []int{fi19, (fi19 % (((fi19 / (((fi19 | (^(fi19)))) | 1))) | 1)), fi19, (fi17 ^ (fi17 + (fi19 << (((v5 & (^(v5)))) & 63))))}
                                                if v4 {
                                                    if false {
                                                        var v69 uint8 = 128
                                                        _ = v69
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi17)
                                                    v1 -= v1
                                                    v18 = -1
                                                    v70 := GenS11{F0: "a", F1: (v18 << (((v5 * (^(v5)))) & 63)), F2: (v5 + (v5 | 4294967295))}
                                                    _ = v70
                                                }
                                                v7 = v7
                                                v1 += v1
                                                _ = v68
                                            }
                                            var v71 int16 = -1
                                            v72 := GenS12{F0: 2, F1: (v71 >> ((v5) & 63)), F2: 18446744073709551615}
                                            fmt.Println(v3, v7)
                                            _ = v71
                                            _ = v72
                                        }
                                    }
                                    v7 = (1.5 * v7)
                                    var v73 int16 = -32768
                                    _ = v73
                                }
                            } else {
                                v74 := []int{-36, fi19, fi19, (fi19 % ((fi19) | 1))}
                                _ = v74
                            }
                        case 128:
                            v75 := []int{}
                            _ = v75
                        case -493:
                            var v76 uint = 1
                            _ = v76
                        default:
                            v7 = ((v2 - (v7 / v2)) * -2.25)
                        }
                        var v77 uint = (v5 - v5)
                        _ = v77
                    }
                }
                _ = fi19
            }
        case 521:
            if v4 {
                if (29 != v6) {
                    v6 += v6
                    v78 := []int{fi17, fi17, fi17, (fi17 >> (((^(v5))) & 63))}
                    _ = v78
                } else {
                    v79 := GenS3{F0: 128, F1: v1, F2: 803}
                    v80 := GenS1{F0: v18, F1: v1}
                    _ = v79
                    _ = v80
                }
                var v81 int32 = 42
                _ = v81
            } else {
                var v82 int32 = 0
                if (!((v18 / ((v18) | 1)) != v18)) {
                    v83 := GenS13{F0: ("a" + ("" + "foo bar"))}
                    fmt.Println(v7, v1)
                    _ = v83
                } else {
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v6)
                    default:
                        var v84 int16 = 0
                        _ = v84
                    }
                }
                _ = v82
            }
        default:
            v2 = (float64((-2.25 * (float32((v7 * float64(((1e10 / 1e10) * ((-2.25 / 1.5) * 0.0))))) + (1.5 - (3.14159 - float32(v2)))))) + ((v2 + 0.0) + (v2 / v2)))
        }
        _ = fi17
        _ = v18
    }
    var v85 int32 = -1
    if v4 {
        var v86 int32 = -1
        if v4 {
            for fi87 := 0; fi87 < 8; fi87++ {
                for fi88 := 0; fi88 < 17; fi88++ {
                    v89 := GenS10{F0: "hello", F1: v6}
                    _ = fi88
                    _ = v89
                }
                v3 = (v3 | 2)
                v90 := []int{}
                _ = fi87
                _ = v90
            }
            if v4 {
                if false {
                    v91 := GenS14{F0: -1, F1: 4}
                    _ = v91
                } else {
                    v92 := []int{0}
                    _ = v92
                }
                v6 += (^(v6))
                v4 = v4
            } else {
                if (v86 == v86) {
                    var v93 int64 = 1
                    var v94 int64 = v93
                    _ = v93
                    _ = v94
                }
                v6 = (v6 >> ((1) & 63))
                var v95 int32 = v86
                _ = v95
            }
        } else {
            v96 := []int{42, -645}
            _ = v96
        }
        for fi97 := 0; fi97 < 15; fi97++ {
            if v4 {
                var v98 float64 = -2.25
                fmt.Println(v86, v6)
                if ((v6 / (((-(v6))) | 1)) >= (v6 ^ (v6 ^ v6))) {
                    fmt.Println(v4, v2, v98)
                    switch -570 {
                    case -9223372036854775808:
                        v99 := GenS15{F0: "\n\t", F1: (v3 >> ((v5) & 63))}
                        _ = v99
                    default:
                        if (2 <= (^(v1))) {
                            var v100 uint64 = 4294967295
                            _ = v100
                        } else {
                            fmt.Println(fi97, v3)
                        }
                    }
                }
                _ = v98
            }
            v86 -= v86
            fmt.Println(v7, v2, v86)
            _ = fi97
        }
        _ = v86
    }
    v6 -= (v6 - -42)
    var v101 float32 = 1.5
    var v102 uint16 = 42
    if false {
        v103 := []int{int(v6)}
        v102 -= (v102 * v102)
        var v104 int = 169
        _ = v103
        _ = v104
    } else {
        for fi105 := 0; fi105 < 2; fi105++ {
            for fi106 := 0; fi106 < 20; fi106++ {
                var v107 uint64 = 18446744073709551615
                v108 := []int{}
                if ("a" > "\n\t") {
                    v85 = (v85 + (v85 + (v85 % ((-1) | 1))))
                    var v109 float64 = v2
                    _ = v109
                } else {
                    var v110 int8 = v6
                    _ = v110
                }
                _ = fi106
                _ = v107
                _ = v108
            }
            var v111 uint = 1
            v112 := GenS12{F0: 0, F1: 0, F2: (v111 >> ((v5) & 63))}
            _ = fi105
            _ = v111
            _ = v112
        }
        for fi113 := 0; fi113 < 16; fi113++ {
            for fi114 := 0; fi114 < 6; fi114++ {
                var v115 bool = v4
                var v116 uint32 = (v3 % ((v3) | 1))
                v5 += uint(v1)
                _ = fi114
                _ = v115
                _ = v116
            }
            for fi117 := 0; fi117 < 2; fi117++ {
                if true {
                    fmt.Println(v101, v7, v6)
                    v7 -= v2
                    v118 := []int{fi113}
                    _ = v118
                }
                var v119 float32 = (float32(0.0) / v101)
                v120 := []int{fi113, 127, (-(fi117)), -9223372036854775808}
                _ = fi117
                _ = v119
                _ = v120
            }
            v121 := GenS16{F0: 316, F1: uint32(v102), F2: (^(v1))}
            _ = fi113
            _ = v121
        }
        var v122 float32 = 1.5
        _ = v122
    }
    fmt.Println(v2, v5, v4)
    var v123 string = "hello"
    for fi124 := 0; fi124 < 7; fi124++ {
        fmt.Println(v2)
        switch v123 {
        case "\n\t":
            var v125 bool = true
            _ = v125
        case "foo bar":
            if true {
                v126 := GenS7{F0: uint64(v3), F1: 94}
                var v127 uint8 = (v1 >> (((v5 << ((285) & 63))) & 63))
                _ = v126
                _ = v127
            } else {
                v128 := []int{fi124, fi124, (fi124 + 9223372036854775807)}
                _ = v128
            }
        default:
            v129 := GenS15{F0: v123, F1: (v3 ^ v3)}
            _ = v129
        }
        v130 := []int{(fi124 - fi124), -591, fi124, -2147483648}
        _ = fi124
        _ = v130
    }
    var v131 int16 = -463
    if false {
        for fi132 := 0; fi132 < 15; fi132++ {
            fmt.Println(v2, fi132, v123)
            fmt.Println(v131, v6)
            if ((0 != -9223372036854775808) && (!v4)) {
                v133 := []int{int(v5)}
                _ = v133
            }
            _ = fi132
        }
        if v4 {
            var v134 uint32 = (v3 & (^(v3)))
            v135 := []int{}
            _ = v134
            _ = v135
        }
        v1 += 127
    } else {
        v5 -= (^(v5))
        var v136 string = v123
        for fi137 := 0; fi137 < 16; fi137++ {
            v136 = v123
            switch 0 {
            case -2147483648:
                switch fi137 {
                case -2147483648:
                    v138 := GenS17{F0: 65535, F1: (^(v3))}
                    _ = v138
                case -9223372036854775808:
                    v139 := GenS18{F0: ("foo bar" + ((v123 + (v136 + ("hello" + v136))) + (v123 + v123)))}
                    _ = v139
                case 4294967295:
                    if false {
                        if (!false) {
                            fmt.Println(v5, v3, fi137)
                            v140 := GenS19{F0: (v1 + v1)}
                            _ = v140
                        }
                    }
                default:
                    for fi141 := 0; fi141 < 10; fi141++ {
                        switch v136 {
                        case "foo bar":
                            v7 = (((v2 + 1e10) + float64(v101)) * 1e10)
                        case "a":
                            fmt.Println(v102, v6, v123)
                        }
                        _ = fi141
                    }
                }
            case 32768:
                var v142 uint32 = v3
                _ = v142
            case 2147483647:
                v85 -= -920
            default:
                var v143 int = (fi137 - (fi137 - fi137))
                _ = v143
            }
            _ = fi137
        }
        _ = v136
    }
    v144 := GenS20{F0: 353}
    for fi145 := 0; fi145 < 18; fi145++ {
        v4 = (int8(v5) > (v6 | (v6 + (-(v6)))))
        _ = fi145
    }
    var v146 bool = (2 <= 42)
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v123:", v123)
    fmt.Println("v131:", v131)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v101
    _ = v102
    _ = v123
    _ = v131
    _ = v144
    _ = v146
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v85
}
