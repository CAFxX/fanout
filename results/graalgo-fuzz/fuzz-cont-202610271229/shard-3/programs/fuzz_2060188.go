package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint; F1 int8; F2 int64 }
type GenS10 struct { F0 int; F1 int }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint8; F1 uint8; F2 bool }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 uint16) uint {
    v6 := GenS1{F0: 42, F1: 9223372036854775807}
    if (-1 >= 0) {
        fmt.Println(p0)
    } else {
        switch 42 {
        case 42:
            if false {
                p0 = (p0 | p0)
                v7 := GenS1{F0: 966, F1: 42}
                p0 += (p0 & p0)
                _ = v7
            } else {
                p0 = (p0 >> ((140) & 63))
                v8 := []int{217}
                p0 -= (p0 | (p0 - (p0 / ((p0) | 1))))
                _ = v8
            }
        case 128:
            v9 := GenS1{F0: 42, F1: 42}
            _ = v9
        }
        p0 -= (p0 + 255)
    }
    p0 += p0
    _ = p0
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = -2147483648
    var v3 uint16 = 255
    var v4 int = 128
    v5 := []int{946}
    var v10 int16 = -895
    var v11 float64 = (((3.14159 - 0.1) / 1.0) / 1.0)
    var v12 bool = (true || (v2 == (v2 & v2)))
    fmt.Println(v3)
    v13 := GenS2{F0: 4294967295, F1: 255, F2: 42}
    if false {
        if true {
            v3 = (^(v3))
        }
        var v14 bool = (0.0 <= 0.1)
        var v15 uint = 432
        _ = v14
        _ = v15
    }
    if v12 {
        fmt.Println(v12, v3)
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 float64 = (float64(float32((3.14159 - (v11 + v11)))) / ((float64(1e10) + 1.5) * v11))
            v12 = v12
            _ = fi16
            _ = v17
        }
        v5 = append(v5, (v4 << ((18446744073709551615) & 63)))
    } else {
        if len(v5) > 0 {
            v5[0] = v4
        }
        switch v4 {
        case 461:
            v18 := GenS1{F0: 610, F1: -9223372036854775808}
            _ = v18
        case -631:
            if v12 {
                var v19 uint8 = 255
                _ = v19
            }
        case -2147483648:
            if v12 {
                var v20 int32 = v2
                v12 = (v12 && v12)
                _ = v20
            } else {
                fmt.Println(v2, v3, v10)
            }
        }
    }
    if v12 {
        if v12 {
            switch v4 {
            case 4294967295:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 127:
                v21 := GenS1{F0: 819, F1: 225}
                _ = v21
            }
        } else {
            v22 := GenS3{F0: 1}
            v12 = (!(v12 || v12))
            var v23 int = (v4 % (((v4 >> ((4294967295) & 63))) | 1))
            _ = v22
            _ = v23
        }
    } else {
        var v24 int = v4
        _ = v24
    }
    if false {
        var v25 int8 = 42
        _ = v25
    }
    fmt.Println(v4, v3)
    v5 = append(v5, v4)
    var v26 uint = 165
    var v27 int = (^(v4))
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 uint8 = 42
        _ = fi28
        _ = v29
    }
    fmt.Println(v3, v11)
    if v12 {
        for fi30 := 0; fi30 < 4; fi30++ {
            if (0.0 != -2.25) {
                var v31 bool = v12
                _ = v31
            }
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 int64 = 0
                var v34 uint32 = uint32(v3)
                v4 = (fi30 % ((v27) | 1))
                v35 := GenS4{F0: 18446744073709551615, F1: 42}
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            _ = fi30
        }
        v1[422] = "hello"
        if (("" + (((("foo bar" + "") + "foo bar") + "") + "\n\t")) == ("a" + "\n\t")) {
            v36 := GenS5{F0: 0, F1: 65}
            for fi37 := 0; fi37 < 8; fi37++ {
                if v12 {
                    v38 := GenS6{F0: v2, F1: (^(v26))}
                    var v39 int32 = v2
                    _ = v38
                    _ = v39
                }
                _ = fi37
            }
            delete(v1, 65535)
            _ = v36
        }
    } else {
        if len(v5) > 0 {
            v5[0] = 128
        }
    }
    if v12 {
        var v40 uint32 = 696
        if (v3 == 2) {
            if len(v5) > 0 {
                v5[0] = -170
            }
            fmt.Println(v2, v10)
            if v12 {
                _ = len(v5)
            } else {
                var v41 uint32 = (v40 % ((v40) | 1))
                if (!(-612 > -1)) {
                    var v42 bool = false
                    var v43 float64 = float64((3.14159 + 3.14159))
                    v44 := GenS4{F0: 4294967295, F1: 4294967295}
                    _ = v42
                    _ = v43
                    _ = v44
                } else {
                    fmt.Println(v10, v10)
                    v1[v41] = ("hello" + "hello")
                    switch "\n\t" {
                    case "\n\t":
                        if ((v26 >= genFn0((v3 / ((v3) | 1)))) && v12) {
                            v45 := GenS1{F0: v41, F1: 9223372036854775807}
                            _ = v45
                        } else {
                            if ((v11 - (1.5 + ((0.1 + (1e10 + v11)) * v11))) != (v11 * float64((-2.25 + ((float32(1.5) + 3.14159) * 0.1))))) {
                                v46 := GenS7{F0: v27, F1: v10}
                                fmt.Println(v3, v41, v10)
                                var v47 int8 = 78
                                _ = v46
                                _ = v47
                            } else {
                                for fi48 := 0; fi48 < 5; fi48++ {
                                    v1[(v41 & (v41 + (v40 & v41)))] = ""
                                    var v49 string = "\n\t"
                                    _ = fi48
                                    _ = v49
                                }
                                var v50 uint = (v26 << (((v26 - v26)) & 63))
                                if v12 {
                                    if false {
                                        v26 = (v26 % ((v26) | 1))
                                    }
                                    v51 := GenS1{F0: v40, F1: 0}
                                    fmt.Println(v11)
                                    _ = v51
                                }
                                _ = v50
                            }
                            v52 := GenS8{F0: (v27 << ((v26) & 63)), F1: 32767, F2: -16}
                            _ = v52
                        }
                    case "foo bar":
                        if (!v12) {
                            fmt.Println(v3, v26)
                        } else {
                            if v12 {
                                fmt.Println(v3, v26, v3)
                            }
                            for fi53 := 0; fi53 < 17; fi53++ {
                                var v54 uint = (v26 * (v26 << (((v26 * (^(v26)))) & 63)))
                                var v55 uint = v26
                                for fi56 := 0; fi56 < 1; fi56++ {
                                    if v12 {
                                        var v57 int = fi53
                                        switch "" {
                                        case "a":
                                            v10 -= (^(v10))
                                        default:
                                            fmt.Println(v57, v3)
                                        }
                                        _ = v57
                                    } else {
                                        var v58 int32 = v2
                                        v4 = (-(v27))
                                        _ = v58
                                    }
                                    _ = fi56
                                }
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                            for fi59 := 0; fi59 < 17; fi59++ {
                                var v60 uint8 = 255
                                v61 := GenS6{F0: (v2 % (((^(v2))) | 1)), F1: genFn0(v3)}
                                for fi62 := 0; fi62 < 14; fi62++ {
                                    var v63 int16 = v10
                                    var v64 int32 = (^(v2))
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                }
                                _ = fi59
                                _ = v60
                                _ = v61
                            }
                        }
                    case "hello":
                        var v65 bool = true
                        _ = v65
                    }
                }
                _ = v41
            }
        } else {
            switch (("\n\t" + "") + "") {
            case "a":
                if v12 {
                    if (true && v12) {
                        for fi66 := 0; fi66 < 8; fi66++ {
                            var v67 int64 = -1
                            switch -9223372036854775808 {
                            case 1:
                                if v12 {
                                    v26 += v26
                                    v1[(v40 | v40)] = ("\n\t" + (("\n\t" + "\n\t") + ("foo bar" + "hello")))
                                    v2 -= v2
                                }
                            case 18:
                                v68 := GenS1{F0: v40, F1: v67}
                                _ = v68
                            case 4294967295:
                                v12 = ((v3 | (v3 & v3)) == (v3 - (^(v3))))
                            default:
                                fmt.Println(v4, v2)
                            }
                            _ = fi66
                            _ = v67
                        }
                        if (-106 < 72) {
                            v69 := GenS9{F0: (^(v26)), F1: 100, F2: 767}
                            for fi70 := 0; fi70 < 2; fi70++ {
                                for fi71 := 0; fi71 < 6; fi71++ {
                                    if ((v3 > (v3 - 42)) && (false || v12)) {
                                        var v72 int64 = 1
                                        _ = v72
                                    } else {
                                        v12 = v12
                                    }
                                    if (float64((float32(3.14159) + ((1e10 + (0.1 - -2.25)) * (-2.25 - 1e10)))) > 1e10) {
                                        var v73 uint16 = (v3 | 32767)
                                        if true {
                                            var v74 int8 = 0
                                            _ = len(v5)
                                            _ = len(v1)
                                            v75 := GenS2{F0: v40, F1: 0, F2: v26}
                                            v73 = v73
                                            if ((1e10 / 0.1) < 3.14159) {
                                                if v12 {
                                                    v11 = 3.14159
                                                    var v76 bool = true
                                                    _ = v76
                                                }
                                                if true {
                                                    var v77 float64 = (-2.25 / float64((-2.25 / (1e10 - 3.14159))))
                                                    if false {
                                                        if (!(!((int32(v73) >= (v2 << (((v26 % ((genFn0(255)) | 1))) & 63))) && false))) {
                                                            v78 := GenS10{F0: (fi71 % ((v4) | 1)), F1: fi70}
                                                            if len(v5) > 0 {
                                                                v5[0] = int(v2)
                                                            }
                                                            v1[v40] = ("foo bar" + "\n\t")
                                                            _ = v78
                                                        } else {
                                                            var v79 float32 = float32(v11)
                                                            _ = v79
                                                        }
                                                    } else {
                                                        var v80 uint32 = (v40 >> (((v26 * genFn0((v73 - v3)))) & 63))
                                                        var v81 int64 = 42
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    var v82 int64 = -42
                                                    _ = v77
                                                    _ = v82
                                                }
                                                var v83 float64 = v11
                                                if ((!v12) || (!v12)) {
                                                    var v84 uint64 = 956
                                                    var v85 int8 = (v74 * v74)
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = v83
                                            } else {
                                                var v86 uint8 = 127
                                                _ = v86
                                            }
                                            _ = v74
                                            _ = v75
                                        } else {
                                            v87 := v1[(v40 - v40)]
                                            v27 -= (v27 | 9223372036854775807)
                                            _ = v87
                                        }
                                        v73 -= 796
                                        switch int(v40) {
                                        case -1:
                                            switch "hello" {
                                            case "":
                                                var v88 int64 = -9223372036854775808
                                                _ = v88
                                            default:
                                                switch v4 {
                                                case -805:
                                                    var v89 int32 = v2
                                                    _ = v89
                                                case 628:
                                                    v2 -= -1
                                                case 4294967295:
                                                    if len(v5) > 0 {
                                                        v5[0] = (^(fi70))
                                                    }
                                                default:
                                                    v90 := v1[(v40 % ((341) | 1))]
                                                    v12 = ("foo bar" != (v90 + v90))
                                                    _ = v90
                                                }
                                            }
                                        case -9223372036854775808:
                                            v91 := GenS11{F0: -83}
                                            _ = v91
                                        }
                                        _ = v73
                                    } else {
                                        if (v12 || v12) {
                                            if (!((true && v12) || (v12 || (v12 || (-128 <= 100))))) {
                                                if v12 {
                                                    v40 -= v40
                                                    fmt.Println(v4)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v92 := GenS10{F0: int(v40), F1: (-(v4))}
                                                        _ = v92
                                                    case "hello":
                                                        switch v4 {
                                                        case -602:
                                                            fmt.Println(fi71, v4, v26)
                                                        case 32767:
                                                            v93 := GenS8{F0: v27, F1: uint16(v26), F2: 42}
                                                            _ = v93
                                                        case -726:
                                                            var v94 uint = (v26 >> ((v26) & 63))
                                                            _ = v94
                                                        }
                                                    case "":
                                                        var v95 uint8 = 92
                                                        _ = v95
                                                    }
                                                } else {
                                                    var v96 uint32 = (v40 & 2147483647)
                                                    _ = v96
                                                }
                                                v3 = uint16(v26)
                                            } else {
                                                v27 = fi70
                                                var v97 float64 = 1.5
                                                _ = v97
                                            }
                                            v98 := GenS12{F0: "\n\t"}
                                            v99 := GenS5{F0: 2, F1: 930}
                                            _ = v98
                                            _ = v99
                                        } else {
                                            v100 := GenS8{F0: fi71, F1: (v3 % (((v3 | (v3 / ((v3) | 1)))) | 1)), F2: -1}
                                            if (v12 && v12) {
                                                v26 = (v26 - uint(v40))
                                                if (!true) {
                                                    var v101 string = "a"
                                                    v102 := GenS13{F0: 147, F1: 0, F2: v12}
                                                    _ = v101
                                                    _ = v102
                                                }
                                                if (false && v12) {
                                                    if len(v5) > 0 {
                                                        v5[0] = fi70
                                                    }
                                                    v40 = (^(v40))
                                                } else {
                                                    if false {
                                                        var v103 uint16 = v3
                                                        _ = v103
                                                    }
                                                }
                                                var v104 int = v27
                                                _ = v104
                                            }
                                            if (v10 == v10) {
                                                fmt.Println(v2)
                                                _ = len(v1)
                                                var v105 float64 = float64((-2.25 + 0.1))
                                                _ = v105
                                            }
                                            _ = v100
                                        }
                                        v106 := GenS14{F0: int32(v26)}
                                        _ = v106
                                    }
                                    _ = fi71
                                }
                                v11 += (0.0 / (((((float64((1.5 * -2.25)) / (v11 - (v11 + v11))) / ((v11 / v11) - v11)) * (v11 / (float64((1.5 / ((float32(v11) / ((float32(1e10) + (((-2.25 - 0.1) * (1e10 / 1.5)) * (3.14159 - 3.14159))) + 3.14159)) / 1e10))) - ((v11 * v11) - v11)))) * -2.25) + float64(1e10)))
                                for fi107 := 0; fi107 < 4; fi107++ {
                                    v4 -= 4294967295
                                    v108 := GenS7{F0: (v27 >> (((v26 ^ v26)) & 63)), F1: (v10 ^ (-(v10)))}
                                    _ = fi107
                                    _ = v108
                                }
                                _ = fi70
                            }
                            _ = v69
                        } else {
                            fmt.Println(v40, v2, v2)
                            _ = len(v1)
                            if v12 {
                                switch v4 {
                                case 518:
                                    switch (v4 - (v4 & v27)) {
                                    case 767:
                                        v109 := GenS15{F0: ((v11 - (v11 + ((v11 / (((3.14159 / v11) * (v11 / v11)) + float64(-2.25))) + 0.1))) >= 1e10), F1: 85, F2: 65}
                                        _ = v109
                                    }
                                case 231:
                                    v110 := GenS15{F0: ((v3 | (v3 % ((v3) | 1))) != (v3 & (v3 % (((^(v3))) | 1)))), F1: 252, F2: 1}
                                    _ = v110
                                case 4294967295:
                                    switch v27 {
                                    case 32767:
                                        v11 = ((v11 + float64(float32(((v11 / v11) - (v11 / v11))))) + float64(0.1))
                                    }
                                }
                                v27 = (v27 << (((v26 | (v26 & v26))) & 63))
                                switch -933 {
                                case 921:
                                    _ = len(v5)
                                case 128:
                                    v5 = append(v5, v4)
                                case 1:
                                    if v12 {
                                        v111 := GenS11{F0: int8(v26)}
                                        switch "hello" {
                                        case "\n\t":
                                            var v112 bool = (false || (((9223372036854775807 >= -1) && v12) && v12))
                                            _ = v112
                                        case "a":
                                            _ = len(v5)
                                        case "foo bar":
                                            if v12 {
                                                for fi113 := 0; fi113 < 11; fi113++ {
                                                    v5 = append(v5, 2147483647)
                                                    var v114 float64 = v11
                                                    _ = fi113
                                                    _ = v114
                                                }
                                                if false {
                                                    var v115 int8 = -1
                                                    _ = v115
                                                }
                                            }
                                        default:
                                            var v116 bool = ((v40 & v40) >= 695)
                                            _ = v116
                                        }
                                        v5 = append(v5, (v4 >> (((v26 * 859)) & 63)))
                                        _ = v111
                                    }
                                }
                            }
                        }
                    }
                }
            case "hello":
                fmt.Println(v12, v3)
            case "\n\t":
                var v117 int = 2147483647
                _ = v117
            }
            var v118 int64 = 1
            var v119 bool = (!(-617 == (v10 * 1)))
            _ = v118
            _ = v119
        }
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
