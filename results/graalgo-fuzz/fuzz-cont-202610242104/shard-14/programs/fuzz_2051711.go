package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint64; F1 int8; F2 int8 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int32; F1 uint8; F2 uint64 }
func genFn0() uint8 {
    var v7 int32 = -1
    v8 := []int{32768, -694, -1}
    var v9 string = ("foo bar" + "\n\t")
    _ = v7
    _ = v8
    _ = v9
    return 181
}
func genFn1(p0 int8, p1 uint32, p2 int64) int64 {
    var v10 float32 = -2.25
    v11 := GenS1{F0: 42, F1: 76, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return p2
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]bool)
    var v4 float32 = 3.14159
    v5 := []string{"hello", "hello"}
    var v6 uint = 203
    fmt.Println(v6, v6, v2)
    v4 += v4
    fmt.Println(v2)
    var v12 float32 = v4
    v2 = v2
    var v13 string = ((("\n\t" + (("a" + (("" + "\n\t") + ("hello" + ("hello" + ((("\n\t" + ("hello" + ("" + ""))) + ("a" + "hello")) + "a"))))) + "a")) + "") + ("a" + "hello"))
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: (fi14 | (fi14 << (((v6 | v6)) & 63))), F1: 127, F2: 1}
        _ = fi14
        _ = v15
    }
    var v16 uint = 2
    delete(v1, ((v13 + v13) + ("\n\t" + v13)))
    v17 := GenS1{F0: 32767, F1: 42, F2: 711}
    for fi18 := 0; fi18 < 19; fi18++ {
        if (false && true) {
            fmt.Println(fi18, v16, v2)
        } else {
            v2 = 1
            var v19 float64 = (0.0 - (1.5 / 1e10))
            for fi20 := 0; fi20 < 18; fi20++ {
                switch (v13 + (v13 + (v13 + v13))) {
                case "hello":
                    v3[42] = true
                case "foo bar":
                    var v21 int16 = 42
                    _ = v21
                }
                if true {
                    if ((^(fi20)) == fi18) {
                        var v22 float64 = 1.5
                        _ = v22
                    }
                    _ = len(v3)
                    switch (-(fi20)) {
                    case 128:
                        var v23 int = fi18
                        var v24 int8 = -27
                        v13 = v13
                        _ = v23
                        _ = v24
                    case -776:
                        var v25 bool = (false || true)
                        _ = v25
                    default:
                        if (!(2 > 32767)) {
                            switch 32767 {
                            case 0:
                                v12 += (v12 * v12)
                                v19 = v19
                            }
                            v26 := v1[(v13 + "a")]
                            _ = v26
                        } else {
                            v27 := GenS1{F0: 0, F1: 1, F2: 282}
                            if false {
                                if true {
                                    var v28 int = fi20
                                    var v29 string = (v13 + v13)
                                    fmt.Println(fi20, v19, v4)
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(v4, v6, fi20)
                            } else {
                                if (false || true) {
                                    delete(v1, v13)
                                    fmt.Println(v16, v13, v4)
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    var v30 uint8 = 89
                                    _ = v30
                                } else {
                                    switch (("" + v13) + (((v13 + ("hello" + v13)) + "a") + v13)) {
                                    case "":
                                        var v31 int16 = 153
                                        _ = v31
                                    case "foo bar":
                                        var v32 int16 = 908
                                        _ = v32
                                    case "a":
                                        if (!true) {
                                            delete(v3, 0)
                                            v33 := GenS2{F0: 18446744073709551615, F1: 32767}
                                            var v34 uint8 = 128
                                            v16 += v6
                                            _ = v33
                                            _ = v34
                                        }
                                    default:
                                        switch v13 {
                                        case "\n\t":
                                            var v35 int8 = 1
                                            _ = v35
                                        case "foo bar":
                                            switch fi18 {
                                            case -1:
                                                v16 = 215
                                                v36 := GenS2{F0: (v2 % (((v2 >> ((v16) & 63))) | 1)), F1: 1}
                                                fmt.Println(fi18)
                                                _ = v36
                                            case 127:
                                                v12 -= v4
                                            case 32767:
                                                var v37 bool = (v13 > ("\n\t" + "a"))
                                                _ = v37
                                            }
                                        }
                                    }
                                    switch fi20 {
                                    case 32768:
                                        fmt.Println(v2)
                                        var v38 int = fi18
                                        v5 = append(v5, v13)
                                        _ = v38
                                    }
                                }
                                v2 += (v2 / (((v2 * (v2 | v2))) | 1))
                                switch (("hello" + ((("" + v13) + v13) + (v13 + v13))) + ("a" + (v13 + ""))) {
                                case "hello":
                                    if len(v5) > 0 {
                                        v5[0] = v13
                                    }
                                case "foo bar":
                                    v19 = 3.14159
                                }
                            }
                            _ = v27
                        }
                    }
                } else {
                    if false {
                        var v39 int = fi20
                        _ = v39
                    }
                    if len(v5) > 0 {
                        v5[0] = ("a" + "\n\t")
                    }
                    v40 := GenS2{F0: (v2 - 204), F1: 42}
                    _ = v40
                }
                var v41 uint8 = 0
                _ = fi20
                _ = v41
            }
            _ = v19
        }
        _ = fi18
    }
    var v42 int8 = -128
    switch 42 {
    case 9223372036854775807:
        v43 := GenS1{F0: -1, F1: 255, F2: 153}
        _ = v43
    case -9223372036854775808:
        for fi44 := 0; fi44 < 4; fi44++ {
            for fi45 := 0; fi45 < 10; fi45++ {
                v12 = ((1e10 * v12) * (v12 * 1e10))
                _ = fi45
            }
            _ = fi44
        }
    case 127:
        var v46 int16 = 32767
        _ = v46
    }
    v12 -= v4
    if false {
        var v47 string = (v13 + v13)
        var v48 uint = v6
        _ = v47
        _ = v48
    } else {
        switch 0 {
        case -9223372036854775808:
            if true {
                if false {
                    v16 += 42
                } else {
                    if true {
                        if false {
                            v49 := GenS2{F0: 0, F1: 32767}
                            fmt.Println(v6)
                            _ = v49
                        } else {
                            v5 = append(v5, v13)
                            fmt.Println(v13, v2, v12)
                        }
                    }
                    v12 -= (v12 - (-2.25 + float32(((0.0 - 1e10) + float64(0.0)))))
                }
                var v50 uint32 = 65535
                var v51 float64 = (1.5 / (0.1 * float64(1.5)))
                _ = v50
                _ = v51
            } else {
                var v52 int32 = 0
                for fi53 := 0; fi53 < 20; fi53++ {
                    for fi54 := 0; fi54 < 15; fi54++ {
                        v55 := GenS3{F0: v13}
                        delete(v1, (v13 + ((v13 + v13) + v13)))
                        var v56 int64 = -637
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    _ = fi53
                }
                _ = v52
            }
        default:
            var v57 int = 324
            _ = v57
        }
        for fi58 := 0; fi58 < 2; fi58++ {
            for fi59 := 0; fi59 < 12; fi59++ {
                if (v4 != v12) {
                    if (-9223372036854775808 <= 9223372036854775807) {
                        v60 := GenS4{F0: "foo bar", F1: (fi59 | (-(fi59)))}
                        if (v2 >= (v2 - (v2 ^ (v2 + (^(v2)))))) {
                            fmt.Println(v4, v16, v2)
                            if (!((v6 ^ 871) >= (v16 >> ((826) & 63)))) {
                                if false {
                                    v42 += 0
                                    fmt.Println(fi59, v6)
                                    var v61 float32 = v12
                                    if len(v5) > 0 {
                                        v5[0] = v13
                                    }
                                    v16 += (v16 & (v16 & 390))
                                    v62 := GenS2{F0: (v2 / ((v2) | 1)), F1: 1000}
                                    var v63 uint16 = 1
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                }
                                v64 := GenS1{F0: (-(fi59)), F1: 192, F2: 2147483647}
                                if true {
                                    v1[v13] = (fi58 ^ fi59)
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v65 bool = (true && (255 >= 2))
                                        _ = v65
                                    case "hello":
                                        v2 -= 247
                                    }
                                    v42 += (v42 - (^(v42)))
                                }
                                _ = v64
                            } else {
                                v6 = v6
                                v66 := GenS5{F0: 1}
                                if (!(25 <= v42)) {
                                    var v67 uint16 = 42
                                    switch 2147483647 {
                                    case 1:
                                        fmt.Println(fi58)
                                    default:
                                        var v68 uint8 = 2
                                        _ = v68
                                    }
                                    fmt.Println(v13, fi58)
                                    v16 += 2
                                    var v69 uint16 = (v67 ^ v67)
                                    _ = v67
                                    _ = v69
                                }
                                var v70 uint8 = 0
                                _ = v66
                                _ = v70
                            }
                            fmt.Println(v12)
                            if false {
                                delete(v1, v13)
                            } else {
                                var v71 float32 = (3.14159 * 1.5)
                                switch v13 {
                                case "\n\t":
                                    v72 := v1["hello"]
                                    _ = v72
                                default:
                                    delete(v3, 1)
                                }
                                if true {
                                    var v73 int64 = -1
                                    fmt.Println(v42, v73, v6)
                                    if (!(0.0 < -2.25)) {
                                        if false {
                                            fmt.Println(v6, v2)
                                            var v74 float64 = 1.5
                                            _ = v74
                                        }
                                        var v75 int32 = int32(fi58)
                                        if false {
                                            if true {
                                                fmt.Println(fi58, v13)
                                                switch v13 {
                                                case "":
                                                    var v76 int64 = v73
                                                    _ = v76
                                                case "a":
                                                    fmt.Println(v2, v12, v73)
                                                case "\n\t":
                                                    v77 := GenS2{F0: uint64(v6), F1: -1}
                                                    v78 := v1[("hello" + "hello")]
                                                    _ = v77
                                                    _ = v78
                                                }
                                                v3[4294967295] = (0.0 == (float64(v4) - 0.1))
                                            }
                                        } else {
                                            v42 += v42
                                        }
                                        var v79 int16 = 1000
                                        var v80 uint32 = 1
                                        _ = v75
                                        _ = v79
                                        _ = v80
                                    } else {
                                        var v81 uint = (v6 >> ((v16) & 63))
                                        v82 := GenS2{F0: 18446744073709551615, F1: 923}
                                        _ = v81
                                        _ = v82
                                    }
                                    switch v13 {
                                    case "a":
                                        var v83 uint32 = 830
                                        _ = v83
                                    case "hello":
                                        var v84 uint = (v6 ^ (v16 + v16))
                                        _ = v84
                                    case "foo bar":
                                        if (!false) {
                                            v85 := GenS5{F0: 617}
                                            v86 := GenS6{F0: v13, F1: 940}
                                            var v87 float64 = float64(v4)
                                            v87 = v87
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        } else {
                                            if (-32768 == 42) {
                                                v13 = v13
                                                var v88 float64 = (float64(v4) * (1.5 - (0.0 * -2.25)))
                                                delete(v3, 2147483647)
                                                _ = v88
                                            }
                                            fmt.Println(v2, fi58)
                                            if (-2147483648 == 851) {
                                                v89 := GenS7{F0: v2, F1: (v42 + v42), F2: -128}
                                                var v90 uint64 = v2
                                                var v91 int8 = 0
                                                _ = v89
                                                _ = v90
                                                _ = v91
                                            } else {
                                                v92 := GenS8{F0: 32767, F1: 0}
                                                switch fi58 {
                                                case 32768:
                                                    if false {
                                                        delete(v1, "hello")
                                                    } else {
                                                        _ = len(v5)
                                                        v13 = (((v13 + "a") + "foo bar") + v13)
                                                        if (!false) {
                                                            v93 := GenS2{F0: v2, F1: 1000}
                                                            v94 := GenS5{F0: 69}
                                                            if false {
                                                                switch v13 {
                                                                case "\n\t":
                                                                    var v95 int8 = (-(v42))
                                                                    _ = v95
                                                                case "a":
                                                                    _ = len(v5)
                                                                }
                                                                if (false && (!true)) {
                                                                    var v96 int32 = -226
                                                                    var v97 uint = 65535
                                                                    switch v13 {
                                                                    case "hello":
                                                                        var v98 int32 = v96
                                                                        _ = v98
                                                                    case "a":
                                                                        var v99 int64 = (v73 & (v73 & v73))
                                                                        _ = v99
                                                                    }
                                                                    _ = v96
                                                                    _ = v97
                                                                } else {
                                                                    if (!true) {
                                                                        v100 := GenS9{F0: 733, F1: 0}
                                                                        if false {
                                                                            if (0 >= 41) {
                                                                                var v101 int = (fi58 ^ fi59)
                                                                                var v102 int8 = v42
                                                                                v42 = v102
                                                                                _ = v101
                                                                                _ = v102
                                                                            }
                                                                        }
                                                                        var v103 int32 = 2147483647
                                                                        _ = v100
                                                                        _ = v103
                                                                    }
                                                                    var v104 uint64 = v2
                                                                    _ = v104
                                                                }
                                                            }
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if true {
                                                            var v105 uint32 = uint32(v16)
                                                            fmt.Println(fi58, v6)
                                                            if (v2 != v2) {
                                                                v105 -= (v105 ^ (^(v105)))
                                                            } else {
                                                                v3[819] = ((v105 % ((v105) | 1)) < v105)
                                                            }
                                                            var v106 int8 = 0
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            if true {
                                                                var v107 int16 = -32768
                                                                _ = v107
                                                            }
                                                        }
                                                        v73 -= -1
                                                        var v108 float32 = (v4 - v71)
                                                        switch fi59 {
                                                        case 769:
                                                            fmt.Println(v12, v6, fi59)
                                                        default:
                                                            v109 := GenS7{F0: uint64(v6), F1: v42, F2: v42}
                                                            _ = v109
                                                        }
                                                        if true {
                                                            fmt.Println(fi58)
                                                            var v110 int32 = int32(v2)
                                                            v3[2] = false
                                                            if len(v5) > 0 {
                                                                v5[0] = v13
                                                            }
                                                            _ = v110
                                                        }
                                                        _ = v108
                                                    }
                                                }
                                                switch (fi58 | fi59) {
                                                case -9223372036854775808:
                                                    v111 := GenS2{F0: (v2 ^ (^(v2))), F1: 1}
                                                    _ = v111
                                                case 127:
                                                    if ((v73 - (^(v73))) >= int64(v2)) {
                                                        v112 := GenS2{F0: v2, F1: 0}
                                                        var v113 int64 = v73
                                                        switch v13 {
                                                        case "\n\t":
                                                            v114 := GenS10{F0: (v42 & (v42 * int8(v2)))}
                                                            _ = v114
                                                        }
                                                        _ = v112
                                                        _ = v113
                                                    }
                                                default:
                                                    v115 := GenS11{F0: -1}
                                                    _ = v115
                                                }
                                                _ = v92
                                            }
                                        }
                                    default:
                                        var v116 uint8 = 84
                                        _ = v116
                                    }
                                    _ = v73
                                }
                                _ = v71
                            }
                            v117 := v3[42]
                            _ = v117
                        }
                        _ = v60
                    } else {
                        v118 := GenS6{F0: ("foo bar" + ("foo bar" + "foo bar")), F1: 1}
                        _ = v118
                    }
                    v119 := GenS12{F0: 2147483647, F1: 255, F2: v2}
                    v6 = v16
                    _ = v119
                }
                var v120 int32 = 313
                v120 = (v120 & (v120 * v120))
                _ = fi59
                _ = v120
            }
            v121 := GenS10{F0: -17}
            _ = fi58
            _ = v121
        }
        var v122 float32 = ((0.0 * v4) + (0.0 / v4))
        _ = v122
    }
    v2 = 33
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
