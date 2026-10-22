package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int64; F1 int64; F2 uint64 }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 int32; F1 uint }
type GenS12 struct { F0 uint64; F1 int8; F2 int64 }
type GenS13 struct { F0 int8; F1 int16; F2 int }
type GenS14 struct { F0 string }
type GenS15 struct { F0 bool; F1 string; F2 bool }
func genFn0(p0 string, p1 string) int {
    v5 := []int{0, 128, 2147483647, 128}
    var v6 uint16 = 255
    var v7 uint16 = (v6 & (v6 % ((v6) | 1)))
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 float64 = 3.14159
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint8, p1 int64) string {
    var v10 int16 = -32768
    var v11 float32 = ((((float32(0.0) / (0.1 * 3.14159)) + (3.14159 - -2.25)) + 3.14159) - 1e10)
    p1 -= int64(p0)
    switch -972 {
    case 32767:
        switch "\n\t" {
        case "":
            fmt.Println(v11)
        case "\n\t":
            fmt.Println(v10)
        case "hello":
            if false {
                var v12 uint64 = 859
                v13 := []int{42, genFn0("", ("" + (("" + (("foo bar" + "foo bar") + "")) + "foo bar")))}
                var v14 int64 = p1
                _ = v12
                _ = v13
                _ = v14
            } else {
                var v15 float64 = (3.14159 * float64(v11))
                if true {
                    v16 := GenS1{F0: 42, F1: ((1e10 - (-2.25 - (v15 + float64(v11)))) > v15)}
                    for fi17 := 0; fi17 < 14; fi17++ {
                        v18 := GenS2{F0: 100, F1: -1}
                        fmt.Println(v11)
                        _ = fi17
                        _ = v18
                    }
                    var v19 string = "a"
                    _ = v16
                    _ = v19
                }
                v11 = ((-2.25 * 1.5) / v11)
                _ = v15
            }
        default:
            v20 := []int{42, 128, 0}
            _ = v20
        }
    case -894:
        v11 -= v11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return "foo bar"
}

func main() {
    var v1 string = "hello"
    v2 := []uint32{2147483648, 755, 65535, 620}
    var v3 uint64 = 18446744073709551615
    var v4 string = v1
    if (1 < 170) {
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 15; fi22++ {
                var v23 float32 = -2.25
                v3 -= 42
                _ = fi22
                _ = v23
            }
            switch v4 {
            case "a":
                for fi24 := 0; fi24 < 19; fi24++ {
                    if (!(431 > 2147483647)) {
                        var v25 int = fi21
                        if true {
                            v26 := GenS3{F0: (0 <= 42), F1: 2147483648}
                            switch fi21 {
                            case 127:
                                if (!(1 > -1)) {
                                    var v27 uint = 65535
                                    if (42 < 1) {
                                        var v28 bool = true
                                        v27 = 1
                                        v29 := GenS2{F0: -1, F1: 100}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        switch v1 {
                                        case "foo bar":
                                            var v30 int = v25
                                            _ = v30
                                        default:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                        var v31 uint64 = (v3 >> ((2) & 63))
                                        v32 := GenS4{F0: 9223372036854775807, F1: -1, F2: v31}
                                        _ = v31
                                        _ = v32
                                    }
                                    fmt.Println(fi21, v27, fi21)
                                    _ = v27
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v33 bool = (!(v25 < 42))
                                    v1 = "\n\t"
                                    _ = v33
                                }
                            case 42:
                                if false {
                                    v4 = (v4 + "a")
                                }
                            case 714:
                                if len(v2) > 0 {
                                    v2[0] = 210
                                }
                            default:
                                fmt.Println(v1)
                            }
                            if false {
                                fmt.Println(v3, v3)
                                v4 = (((v1 + v1) + v1) + (("" + "\n\t") + (v1 + v1)))
                                v1 = (v1 + (((v1 + ("hello" + v1)) + v1) + ("hello" + "")))
                            } else {
                                var v34 int = (fi21 << ((255) & 63))
                                _ = len(v2)
                                fmt.Println(v34)
                                _ = v34
                            }
                            _ = v26
                        } else {
                            v4 = (v4 + v1)
                            v35 := GenS4{F0: 9223372036854775807, F1: 1, F2: v3}
                            _ = v35
                        }
                        _ = v25
                    } else {
                        var v36 bool = (960 > 292)
                        var v37 string = "\n\t"
                        switch ("" + v4) {
                        case "":
                            v4 = "a"
                        case "foo bar":
                            if (v36 || false) {
                                fmt.Println(v4, fi21, fi24)
                            } else {
                                if v36 {
                                    var v38 uint = 1
                                    var v39 int16 = int16(v38)
                                    _ = v38
                                    _ = v39
                                }
                                if v36 {
                                    if (65535 != 255) {
                                        if v36 {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            v1 = (v1 + v4)
                                            var v40 uint8 = 255
                                            v41 := GenS2{F0: 95, F1: 0}
                                            _ = v40
                                            _ = v41
                                        }
                                        var v42 float32 = 3.14159
                                        var v43 int64 = 863
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if (((float32(0.0) / 1.5) + 0.1) <= -2.25) {
                                            var v44 uint16 = 42
                                            if (v36 || (v36 && v36)) {
                                                if (true && (0 > 42)) {
                                                    v36 = (195 <= 255)
                                                    v2 = append(v2, 0)
                                                    if false {
                                                        if v36 {
                                                            if (1 > 0) {
                                                                if (9223372036854775807 >= -1) {
                                                                    v44 += (v44 % (((v44 / (((v44 | v44)) | 1))) | 1))
                                                                    v36 = (!((649 < 699) && ((fi21 & int(v44)) == fi21)))
                                                                    v44 += (v44 | 65535)
                                                                } else {
                                                                    v45 := GenS1{F0: 4294967295, F1: v36}
                                                                    _ = v45
                                                                }
                                                                if v36 {
                                                                    var v46 bool = (-9223372036854775808 == 9223372036854775807)
                                                                    _ = v46
                                                                } else {
                                                                    if (true || v36) {
                                                                        fmt.Println(v3, v44)
                                                                        _ = len(v2)
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = 475
                                                                        }
                                                                    }
                                                                    var v47 uint32 = 2
                                                                    _ = v47
                                                                }
                                                            } else {
                                                                v48 := GenS3{F0: (-32768 < 42), F1: 1}
                                                                var v49 uint8 = uint8(v3)
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            fmt.Println(v1, v36, v1)
                                                            var v50 int8 = -111
                                                            _ = v50
                                                        }
                                                        var v51 float32 = 3.14159
                                                        _ = v51
                                                    }
                                                }
                                                v4 = v1
                                                var v52 uint32 = 804
                                                if v36 {
                                                    var v53 float32 = 1e10
                                                    var v54 int32 = -197
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    v37 = (((v4 + v1) + (v37 + ("" + v4))) + (v1 + ("\n\t" + (v1 + v4))))
                                                    switch fi21 {
                                                    case 32768:
                                                        v36 = (255 != 42)
                                                    case 76:
                                                        switch v4 {
                                                        case "a":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v37)
                                                        }
                                                    default:
                                                        fmt.Println(fi21, v1)
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = (v52 | (v52 << ((65535) & 63)))
                                                    }
                                                }
                                                var v55 float64 = (0.0 / 1.0)
                                                v4 = "a"
                                                _ = v52
                                                _ = v55
                                            }
                                            var v56 uint8 = 154
                                            fmt.Println(v4, fi24, fi21)
                                            var v57 uint = uint(v3)
                                            _ = v44
                                            _ = v56
                                            _ = v57
                                        }
                                        if (128 < 1) {
                                            v37 = ""
                                            v3 = 1
                                            _ = len(v2)
                                            var v58 uint64 = v3
                                            v2 = append(v2, 702)
                                            _ = v58
                                        } else {
                                            v59 := GenS5{F0: 1, F1: 42}
                                            _ = v59
                                        }
                                        if (-2147483648 != 42) {
                                            if (!((1000 <= 42) && (!v36))) {
                                                fmt.Println(v3)
                                                if (2 < 42) {
                                                    v2 = append(v2, 1)
                                                    var v60 int64 = 674
                                                    v4 = "hello"
                                                    _ = v60
                                                }
                                            } else {
                                                if (!true) {
                                                    if (v36 || ((9223372036854775807 < -453) && (-772 > 1))) {
                                                        v37 = v4
                                                    }
                                                    v61 := GenS5{F0: 117, F1: 4294967295}
                                                    var v62 uint = 255
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    _ = len(v2)
                                                }
                                                switch "hello" {
                                                case "":
                                                    v36 = ((!(-157 < 613)) && (217 <= 108))
                                                case "\n\t":
                                                    v63 := GenS2{F0: 42, F1: 0}
                                                    _ = v63
                                                }
                                            }
                                            switch fi24 {
                                            case 0:
                                                v3 = (^(v3))
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                                v64 := GenS6{F0: 2}
                                                var v65 int64 = -9223372036854775808
                                                _ = v64
                                                _ = v65
                                            case 42:
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                            }
                                            v36 = v36
                                        } else {
                                            var v66 int8 = 1
                                            switch 32767 {
                                            case 42:
                                                var v67 int = 210
                                                _ = v67
                                            case 520:
                                                fmt.Println(v3, fi24, v36)
                                            }
                                            _ = v66
                                        }
                                    }
                                    var v68 uint8 = 1
                                    v68 -= v68
                                    _ = v68
                                } else {
                                    if v36 {
                                        var v69 int64 = -1
                                        switch "hello" {
                                        case "hello":
                                            var v70 float32 = -2.25
                                            _ = v70
                                        case "a":
                                            if v36 {
                                                var v71 bool = (v36 && (v36 || v36))
                                                switch (v1 + v4) {
                                                case "hello":
                                                    var v72 float64 = 0.0
                                                    _ = v72
                                                }
                                                _ = v71
                                            }
                                        case "foo bar":
                                            v73 := GenS7{F0: (v3 >> ((1) & 63)), F1: 100}
                                            _ = v73
                                        default:
                                            fmt.Println(v3, v69)
                                        }
                                        _ = v69
                                    }
                                    var v74 int32 = 0
                                    _ = v74
                                }
                            }
                        default:
                            _ = len(v2)
                        }
                        v75 := GenS3{F0: true, F1: uint32(v3)}
                        switch genFn0("foo bar", v1) {
                        case 4294967295:
                            if v36 {
                                if v36 {
                                    var v76 float64 = 1.5
                                    _ = v76
                                } else {
                                    v77 := GenS7{F0: v3, F1: 87}
                                    v78 := GenS8{F0: 42}
                                    _ = v77
                                    _ = v78
                                }
                            }
                        default:
                            if v36 {
                                if true {
                                    var v79 float64 = -2.25
                                    var v80 int64 = 431
                                    v36 = (182 <= 601)
                                    _ = v79
                                    _ = v80
                                } else {
                                    var v81 bool = ((-1 <= -555) || false)
                                    _ = v81
                                }
                                v37 = ("" + v37)
                            } else {
                                fmt.Println(v37, v3, v37)
                                v37 = v1
                                var v82 bool = true
                                _ = v82
                            }
                        }
                        _ = v36
                        _ = v37
                        _ = v75
                    }
                    _ = fi24
                }
            case "":
                if true {
                    if false {
                        switch fi21 {
                        case 2147483647:
                            for fi83 := 0; fi83 < 12; fi83++ {
                                var v84 int16 = 0
                                v3 -= v3
                                _ = fi83
                                _ = v84
                            }
                        case 0:
                            switch fi21 {
                            case 366:
                                switch (v1 + (v1 + v4)) {
                                case "":
                                    switch (fi21 >> ((255) & 63)) {
                                    case -993:
                                        var v85 uint32 = 497
                                        _ = v85
                                    default:
                                        var v86 uint32 = uint32(v3)
                                        _ = v86
                                    }
                                }
                            case 992:
                                _ = len(v2)
                            }
                        default:
                            v4 = v1
                        }
                        var v87 float64 = 1.5
                        v88 := GenS8{F0: -100}
                        _ = v87
                        _ = v88
                    }
                } else {
                    v3 = 0
                    var v89 int8 = -1
                    _ = v89
                }
            case "foo bar":
                if true {
                    var v90 uint64 = v3
                    var v91 float32 = 0.0
                    _ = v90
                    _ = v91
                } else {
                    for fi92 := 0; fi92 < 19; fi92++ {
                        v3 -= (v3 * 2)
                        var v93 string = "foo bar"
                        _ = fi92
                        _ = v93
                    }
                }
            default:
                switch fi21 {
                case -679:
                    v94 := GenS9{F0: 393, F1: -2147483648}
                    _ = v94
                default:
                    var v95 bool = (!(18446744073709551615 > 42))
                    _ = v95
                }
            }
            _ = fi21
        }
        for fi96 := 0; fi96 < 8; fi96++ {
            for fi97 := 0; fi97 < 19; fi97++ {
                if false {
                    var v98 int32 = 42
                    v4 = "hello"
                    var v99 int16 = 42
                    _ = v98
                    _ = v99
                } else {
                    v100 := GenS9{F0: fi97, F1: -1}
                    v4 = ("hello" + "foo bar")
                    if (42 < 42) {
                        if false {
                            fmt.Println(fi97)
                            if (65535 < 255) {
                                switch v1 {
                                case "\n\t":
                                    v4 = (v4 + "\n\t")
                                }
                                v101 := GenS10{F0: 42, F1: 190}
                                switch (fi97 >> ((uint(v3)) & 63)) {
                                case -9223372036854775808:
                                    _ = len(v2)
                                case 4294967295:
                                    var v102 uint64 = v3
                                    _ = v102
                                }
                                v103 := GenS4{F0: -466, F1: 9223372036854775807, F2: (^(v3))}
                                _ = v101
                                _ = v103
                            }
                        } else {
                            if (0.1 != 1e10) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v104 := GenS11{F0: -2147483648, F1: 2}
                                _ = v104
                            } else {
                                v1 = (((v4 + (v1 + ((v4 + v1) + ""))) + (v1 + v1)) + v1)
                                v3 = v3
                            }
                            v105 := GenS12{F0: v3, F1: -60, F2: -1}
                            switch 32768 {
                            case 0:
                                var v106 string = "a"
                                _ = v106
                            default:
                                fmt.Println(v1)
                            }
                            var v107 bool = false
                            if v107 {
                                fmt.Println(v1, fi96, v1)
                                if v107 {
                                    v1 = (v4 + "a")
                                    var v108 uint16 = 255
                                    _ = v108
                                } else {
                                    v109 := GenS7{F0: v3, F1: 127}
                                    _ = v109
                                }
                                var v110 int32 = 1
                                v111 := GenS1{F0: 65535, F1: (!v107)}
                                _ = v110
                                _ = v111
                            } else {
                                fmt.Println(v1, fi97)
                                v1 = (v1 + "hello")
                                var v112 bool = true
                                _ = v112
                            }
                            _ = v105
                            _ = v107
                        }
                    } else {
                        if ((32767 <= 0) || (v3 == v3)) {
                            var v113 bool = true
                            _ = len(v2)
                            var v114 uint8 = 255
                            _ = v113
                            _ = v114
                        } else {
                            v115 := GenS13{F0: 42, F1: 32767, F2: fi97}
                            var v116 int64 = -9223372036854775808
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v115
                            _ = v116
                        }
                        var v117 uint8 = 2
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v117
                    }
                    v4 = "\n\t"
                    v3 = (v3 ^ v3)
                    var v118 float32 = 3.14159
                    switch 566 {
                    case 42:
                        var v119 uint8 = 129
                        _ = v119
                    case -9223372036854775808:
                        v2 = append(v2, 1)
                    case 4294967295:
                        if false {
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            } else {
                                if ((false && ((-1 != 100) && (!true))) || true) {
                                    var v120 uint16 = 674
                                    _ = v120
                                } else {
                                    var v121 uint64 = v3
                                    fmt.Println(v121, v4)
                                    _ = v121
                                }
                            }
                            var v122 string = v1
                            _ = v122
                        } else {
                            if (true && (1 < int64(fi97))) {
                                var v123 string = ("\n\t" + v1)
                                _ = v123
                            } else {
                                v2 = append(v2, 210)
                            }
                        }
                    default:
                        var v124 string = v4
                        _ = v124
                    }
                    var v125 uint8 = 127
                    var v126 string = (v1 + "a")
                    if (223 == 2) {
                        v4 = (v126 + "\n\t")
                        var v127 int16 = -50
                        _ = v127
                    } else {
                        fmt.Println(v4)
                        v3 = (v3 >> ((18446744073709551615) & 63))
                        v128 := GenS14{F0: v1}
                        _ = v128
                    }
                    _ = v100
                    _ = v118
                    _ = v125
                    _ = v126
                }
                _ = fi97
            }
            if ((fi96 <= genFn0(("a" + ((v4 + "hello") + v1)), "")) && ((uint16(v3) > 42) || false)) {
                if len(v2) > 0 {
                    v2[0] = 541
                }
            } else {
                var v129 int16 = 32767
                _ = v129
            }
            _ = fi96
        }
    } else {
        if (((!false) || (false || (0 <= 1000))) || (-1 == -1)) {
            var v130 bool = true
            _ = v130
        } else {
            v131 := GenS2{F0: 26, F1: 127}
            _ = len(v2)
            if (127 != 246) {
                var v132 int8 = -82
                v132 -= (v132 - 0)
                v1 = "hello"
                _ = v132
            } else {
                fmt.Println(v1)
                switch 0 {
                case 32768:
                    for fi133 := 0; fi133 < 2; fi133++ {
                        fmt.Println(v4)
                        fmt.Println(v4, fi133)
                        v134 := GenS15{F0: ((4294967295 < 117) && false), F1: ((v4 + v1) + "foo bar"), F2: false}
                        _ = fi133
                        _ = v134
                    }
                default:
                    var v135 uint64 = (v3 - 1)
                    _ = v135
                }
            }
            _ = v131
        }
    }
    var v136 int8 = 75
    v137 := GenS3{F0: true, F1: 2147483648}
    v4 = ((((v4 + ("foo bar" + v1)) + v1) + "foo bar") + v4)
    fmt.Println(v1, v3)
    var v138 string = v4
    if (!(!(0.0 >= (((0.1 / ((float32(((-2.25 * -2.25) + 1e10)) * 0.1) / (-2.25 + (0.0 * 1.5)))) + 0.0) / 1.0)))) {
        if (!(true || false)) {
            for fi139 := 0; fi139 < 6; fi139++ {
                var v140 int8 = 100
                _ = fi139
                _ = v140
            }
            v3 = v3
        }
        var v141 uint32 = 4294967295
        _ = v141
    }
    var v142 uint = 65535
    v2 = append(v2, 65535)
    var v143 int16 = 0
    for fi144 := 0; fi144 < 3; fi144++ {
        v138 = "\n\t"
        switch (v138 + v1) {
        case "a":
            v2 = append(v2, 479)
        }
        _ = fi144
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi145 := 0; fi145 < 6; fi145++ {
        v142 = v142
        v1 = "hello"
        var v146 uint32 = 92
        _ = fi145
        _ = v146
    }
    v4 = ""
    fmt.Println("v1:", v1)
    fmt.Println("v136:", v136)
    fmt.Println("v138:", v138)
    fmt.Println("v142:", v142)
    fmt.Println("v143:", v143)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v136
    _ = v137
    _ = v138
    _ = v142
    _ = v143
    _ = v2
    _ = v3
    _ = v4
}
