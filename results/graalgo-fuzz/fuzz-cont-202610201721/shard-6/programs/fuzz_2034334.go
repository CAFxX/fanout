package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint16; F1 int64; F2 int16 }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int; F1 uint16 }
type GenS14 struct { F0 uint32; F1 uint8; F2 string }
type GenS15 struct { F0 bool; F1 int16; F2 uint32 }
type GenS16 struct { F0 int8; F1 int8 }
type GenS17 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 int }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 int8; F1 string }
type GenS22 struct { F0 uint32; F1 uint8 }
type GenS23 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS24 struct { F0 string; F1 uint32; F2 uint16 }
type GenS25 struct { F0 string; F1 uint64; F2 int64 }
type GenS26 struct { F0 uint; F1 int }
type GenS27 struct { F0 int; F1 uint32 }
type GenS28 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint8) string {
    p0 += (p0 + p0)
    p0 += (^(p0))
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        var v8 float32 = (((0.0 + (float32((1.5 + -2.25)) + 0.1)) - 0.1) * ((((0.0 - 0.0) * -2.25) - 1.5) - 1.5))
        _ = v8
    case "":
        v9 := GenS1{F0: ("a" + "a"), F1: 222}
        _ = v9
    case "\n\t":
        switch "foo bar" {
        case "hello":
            p0 = p0
        case "\n\t":
            fmt.Println(p0)
        default:
            var v10 float64 = (0.0 + 1.5)
            _ = v10
        }
    default:
        var v11 int64 = -1
        _ = v11
    }
    _ = p0
    return ("a" + ("hello" + "hello"))
}
func genFn1(p0 int64, p1 bool, p2 int) string {
    var v12 float64 = 1e10
    v13 := []int{(-(p2)), 2147483647, p2}
    switch (("a" + "\n\t") + "\n\t") {
    case "a":
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = -2.25
    var v3 string = ("hello" + ("hello" + (("\n\t" + "\n\t") + "\n\t")))
    v4 := make(map[uint32]int)
    var v5 bool = true
    v5 = false
    switch 128 {
    case 32768:
        fmt.Println(v5)
    default:
        var v14 int8 = 100
        _ = v14
    }
    if ((0.0 < (v2 / 3.14159)) && (v5 && (v3 <= v3))) {
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := []int{4294967295}
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 18; fi17++ {
            if v5 {
                var v18 int64 = -583
                fmt.Println(v2, v3, v5)
                v2 += float64(float32((v2 - 0.0)))
                _ = v18
            } else {
                v5 = v5
                switch "foo bar" {
                case "":
                    v5 = (v5 && (false || ("foo bar" != "")))
                case "a":
                    v19 := []int{883, fi17, (fi17 & (fi17 | (fi17 * fi17)))}
                    _ = v19
                case "foo bar":
                    v20 := GenS2{F0: 32767, F1: -168}
                    _ = v20
                default:
                    v2 = v2
                }
            }
            for fi21 := 0; fi21 < 20; fi21++ {
                v2 = 1e10
                if (42 > 2) {
                    if (!(101 == -2147483648)) {
                        v2 -= 1e10
                        v4[2147483648] = -9223372036854775808
                        if true {
                            v22 := GenS2{F0: -332, F1: 9223372036854775807}
                            var v23 int16 = 1000
                            v2 -= (v2 - 0.1)
                            switch fi17 {
                            case 530:
                                v4[660] = 128
                            }
                            var v24 float32 = float32((float64(-2.25) * v2))
                            _ = v22
                            _ = v23
                            _ = v24
                        } else {
                            switch (fi21 / ((fi21) | 1)) {
                            case 42:
                                v25 := []int{994, 0}
                                v26 := GenS1{F0: (v3 + v3), F1: fi21}
                                v2 = (v2 + v2)
                                switch "" {
                                case "\n\t":
                                    _ = len(v1)
                                case "":
                                    switch v3 {
                                    case "hello":
                                        var v27 string = (v3 + v3)
                                        _ = v27
                                    case "foo bar":
                                        if false {
                                            v28 := GenS3{F0: 65535, F1: 42, F2: 42}
                                            _ = v28
                                        }
                                    default:
                                        v29 := GenS4{F0: ((0 < 2147483647) || (1 != 2))}
                                        _ = v29
                                    }
                                default:
                                    var v30 int8 = -1
                                    _ = v30
                                }
                                fmt.Println(v2, v3, v5)
                                _ = v25
                                _ = v26
                            case -2147483648:
                                switch v3 {
                                case "\n\t":
                                    v31 := []int{fi17, (fi21 >> ((2) & 63)), fi17, (fi17 + (^(fi17)))}
                                    _ = v31
                                case "a":
                                    switch v3 {
                                    case "\n\t":
                                        v32 := GenS1{F0: "hello", F1: fi17}
                                        _ = v32
                                    case "a":
                                        if v5 {
                                            if (v5 || true) {
                                                fmt.Println(fi17)
                                                v33 := GenS5{F0: -144}
                                                var v34 uint = 255
                                                _ = v33
                                                _ = v34
                                            }
                                            switch v3 {
                                            case "a":
                                                v35 := []int{}
                                                _ = v35
                                            default:
                                                var v36 uint64 = 4294967295
                                                _ = v36
                                            }
                                            fmt.Println(fi21, v5, v5)
                                            var v37 uint32 = 2147483647
                                            switch fi17 {
                                            case 1:
                                                v4[(^(v37))] = fi17
                                            default:
                                                if v5 {
                                                    if (!v5) {
                                                        v37 -= 65535
                                                        v2 = v2
                                                    }
                                                    switch 42 {
                                                    case -9223372036854775808:
                                                        var v38 float64 = 0.0
                                                        _ = v38
                                                    case -1:
                                                        var v39 uint8 = 2
                                                        _ = v39
                                                    default:
                                                        var v40 bool = v5
                                                        _ = v40
                                                    }
                                                }
                                            }
                                            v41 := GenS6{F0: 32767}
                                            _ = v37
                                            _ = v41
                                        }
                                    default:
                                        v2 += v2
                                    }
                                }
                            case 1:
                                if true {
                                    v3 = "foo bar"
                                    if false {
                                        var v42 int32 = -1
                                        v43 := GenS6{F0: 32767}
                                        _ = v42
                                        _ = v43
                                    } else {
                                        if false {
                                            v3 = ("foo bar" + ("a" + ""))
                                            v2 = v2
                                        } else {
                                            v44 := GenS6{F0: 65535}
                                            _ = v44
                                        }
                                        if (!((-1 < 9223372036854775807) && (((161 < 128) && v5) || (1 <= 6)))) {
                                            v45 := v4[2147483648]
                                            v46 := GenS7{F0: 0}
                                            _ = v45
                                            _ = v46
                                        } else {
                                            v47 := v1[93]
                                            var v48 float32 = -2.25
                                            var v49 string = ("a" + v3)
                                            _ = v47
                                            _ = v48
                                            _ = v49
                                        }
                                        fmt.Println(fi21)
                                        fmt.Println(fi17, v3)
                                        v3 = v3
                                    }
                                    switch fi17 {
                                    case -610:
                                        var v50 int8 = -1
                                        _ = v50
                                    default:
                                        fmt.Println(v5, v2, fi21)
                                    }
                                }
                            }
                            v51 := GenS5{F0: 9223372036854775807}
                            _ = v51
                        }
                    }
                    v52 := []int{(fi17 | (^(fi21))), (fi17 ^ (fi17 - fi21))}
                    switch fi17 {
                    case 975:
                        delete(v1, 42)
                    case -9223372036854775808:
                        if v5 {
                            switch (-(fi21)) {
                            case 145:
                                if v5 {
                                    switch "foo bar" {
                                    case "hello":
                                        if (894 < 2147483648) {
                                            var v53 string = v3
                                            _ = v53
                                        }
                                    case "\n\t":
                                        if (!v5) {
                                            v52 = append(v52, 1)
                                            if (v3 == "hello") {
                                                v54 := GenS6{F0: 1}
                                                v2 -= v2
                                                var v55 uint16 = 42
                                                _ = v54
                                                _ = v55
                                            } else {
                                                switch (fi21 % ((fi21) | 1)) {
                                                case -7:
                                                    v2 = v2
                                                case -916:
                                                    v3 = (v3 + "\n\t")
                                                }
                                                v3 = (v3 + ("foo bar" + v3))
                                                v52 = append(v52, (fi21 - fi21))
                                            }
                                        }
                                    default:
                                        if (32768 != -9223372036854775808) {
                                            if v5 {
                                                if true {
                                                    v5 = v5
                                                    if (!false) {
                                                        if (false || (((-(fi17)) == 9223372036854775807) || v5)) {
                                                            v56 := GenS6{F0: 0}
                                                            _ = v56
                                                        } else {
                                                            fmt.Println(v5, v5, fi21)
                                                            v57 := GenS4{F0: (0.1 != 1.5)}
                                                            v3 = v3
                                                            _ = v57
                                                        }
                                                        if ("a" < ("a" + v3)) {
                                                            v1[0] = true
                                                        }
                                                    }
                                                    if (true && false) {
                                                        v58 := GenS8{F0: 42, F1: 42}
                                                        v59 := GenS6{F0: 42}
                                                        if v5 {
                                                            v2 = v2
                                                            var v60 int = fi21
                                                            var v61 uint64 = 192
                                                            _ = v60
                                                            _ = v61
                                                        } else {
                                                            switch (fi17 % (((fi21 >> ((39) & 63))) | 1)) {
                                                            case 4294967295:
                                                                if true {
                                                                    switch 128 {
                                                                    case 784:
                                                                        switch 9223372036854775807 {
                                                                        case 9223372036854775807:
                                                                            var v62 bool = false
                                                                            _ = v62
                                                                        }
                                                                    case -81:
                                                                        if false {
                                                                            if v5 {
                                                                                var v63 uint16 = 2
                                                                                _ = v63
                                                                            }
                                                                            if (v5 || v5) {
                                                                                v3 = v3
                                                                            } else {
                                                                                if false {
                                                                                    fmt.Println(v3, v3)
                                                                                    v52 = append(v52, 651)
                                                                                } else {
                                                                                    if (18446744073709551615 != 451) {
                                                                                        fmt.Println(v3, v3, fi17)
                                                                                    }
                                                                                    delete(v4, 2147483647)
                                                                                    v3 = v3
                                                                                    var v64 float64 = v2
                                                                                    v65 := GenS6{F0: 32767}
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                                if ((3.14159 * -2.25) == 1.5) {
                                                                                    var v66 bool = (1 != 255)
                                                                                    switch v3 {
                                                                                    case "hello":
                                                                                        var v67 float64 = v2
                                                                                        _ = v67
                                                                                    }
                                                                                    v3 = (((v3 + (v3 + "hello")) + (((v3 + v3) + ("" + (((v3 + v3) + v3) + (v3 + "")))) + v3)) + v3)
                                                                                    _ = v66
                                                                                }
                                                                            }
                                                                        }
                                                                    case -844:
                                                                        v3 = ("hello" + v3)
                                                                        if (957 < 707) {
                                                                            fmt.Println(v3, fi17, v2)
                                                                            if true {
                                                                                v3 = v3
                                                                                if len(v52) >= 2 {
                                                                                    v52 = v52[1:len(v52)]
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v68 int8 = 100
                                                                            if (1 >= 801) {
                                                                                v69 := GenS7{F0: v68}
                                                                                _ = v69
                                                                            }
                                                                            switch fi17 {
                                                                            case -2147483648:
                                                                                fmt.Println(fi21, fi21)
                                                                            }
                                                                            var v70 uint = 18446744073709551615
                                                                            v70 -= v70
                                                                            if (0 > 2147483648) {
                                                                                switch 128 {
                                                                                case 128:
                                                                                    v71 := GenS6{F0: 0}
                                                                                    _ = v71
                                                                                case -9223372036854775808:
                                                                                    var v72 bool = v5
                                                                                    _ = v72
                                                                                case -700:
                                                                                    var v73 float32 = (1.5 - (1.5 - (0.1 - 0.0)))
                                                                                    _ = v73
                                                                                default:
                                                                                    v68 = v68
                                                                                }
                                                                            } else {
                                                                                v3 = v3
                                                                            }
                                                                            _ = v68
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (fi21 < (fi17 % (((fi21 << ((301) & 63))) | 1))) {
                                                                        v74 := GenS3{F0: 18446744073709551615, F1: 2, F2: -1}
                                                                        _ = v74
                                                                    } else {
                                                                        var v75 int8 = -1
                                                                        v76 := GenS8{F0: (-(v75)), F1: -206}
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    switch v3 {
                                                                    case "hello":
                                                                        var v77 float64 = -2.25
                                                                        _ = v77
                                                                    case "foo bar":
                                                                        if (fi17 == 1) {
                                                                            var v78 int = (fi21 & fi17)
                                                                            _ = v78
                                                                        }
                                                                    case "a":
                                                                        v79 := GenS7{F0: 0}
                                                                        _ = v79
                                                                    default:
                                                                        v3 = v3
                                                                    }
                                                                }
                                                            case 0:
                                                                var v80 int8 = 1
                                                                _ = v80
                                                            case 32768:
                                                                v2 += ((((1e10 - v2) + v2) * v2) + v2)
                                                            default:
                                                                v81 := GenS9{F0: 2147483647}
                                                                _ = v81
                                                            }
                                                            fmt.Println(fi21)
                                                        }
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        fmt.Println(v5, v5)
                                                    }
                                                    var v82 uint64 = 18446744073709551615
                                                    _ = v82
                                                } else {
                                                    v2 = 3.14159
                                                    var v83 uint8 = 127
                                                    _ = v83
                                                }
                                                v2 = (1e10 / (1.5 * (float64(-2.25) / v2)))
                                            }
                                            var v84 float32 = 0.0
                                            if false {
                                                var v85 uint8 = 128
                                                var v86 bool = v5
                                                switch (fi17 % ((4) | 1)) {
                                                case 0:
                                                    var v87 int64 = 680
                                                    _ = v87
                                                case 127:
                                                    v86 = v86
                                                case 9223372036854775807:
                                                    var v88 uint32 = 695
                                                    _ = v88
                                                }
                                                if true {
                                                    if len(v52) > 0 {
                                                        v52[0] = fi21
                                                    }
                                                } else {
                                                    v1[850] = v5
                                                }
                                                if (1000 <= -1) {
                                                    if (42 >= 537) {
                                                        v89 := GenS10{F0: 1, F1: 9223372036854775807, F2: 160}
                                                        v86 = v5
                                                        _ = v89
                                                    } else {
                                                        v85 -= v85
                                                        var v90 float64 = v2
                                                        _ = v90
                                                    }
                                                }
                                                _ = v85
                                                _ = v86
                                            } else {
                                                v3 = v3
                                                var v91 int32 = 1
                                                _ = v91
                                            }
                                            _ = v84
                                        } else {
                                            v2 += v2
                                            var v92 uint8 = 42
                                            delete(v1, 677)
                                            _ = v92
                                        }
                                    }
                                    switch v3 {
                                    case "hello":
                                        delete(v1, 823)
                                        var v93 int32 = 350
                                        _ = v93
                                    case "":
                                        if (!false) {
                                            var v94 uint8 = 127
                                            v95 := GenS2{F0: 32767, F1: -9223372036854775808}
                                            _ = v94
                                            _ = v95
                                        }
                                    default:
                                        if (332 > 2) {
                                            if (1 != 2) {
                                                v5 = v5
                                                if (!(42 != 0)) {
                                                    if v5 {
                                                        var v96 uint16 = 0
                                                        var v97 uint8 = 42
                                                        v3 = v3
                                                        _ = v96
                                                        _ = v97
                                                    } else {
                                                        v3 = ""
                                                        v98 := GenS11{F0: 0, F1: 759}
                                                        v3 = "foo bar"
                                                        _ = v98
                                                    }
                                                    _ = len(v1)
                                                    if (1 >= 0) {
                                                        var v99 int8 = -1
                                                        var v100 uint16 = 1
                                                        v99 = (v99 ^ v99)
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                } else {
                                                    var v101 uint32 = 4294967295
                                                    fmt.Println(v101, v3, fi21)
                                                    _ = v101
                                                }
                                                var v102 uint = 4294967295
                                                if len(v52) > 0 {
                                                    v52[0] = (fi17 & fi21)
                                                }
                                                switch (v3 + v3) {
                                                case "":
                                                    if (!((^(v102)) >= v102)) {
                                                        if (42 != -900) {
                                                            if (true || v5) {
                                                                switch fi21 {
                                                                case -9223372036854775808:
                                                                    var v103 float64 = 1e10
                                                                    _ = v103
                                                                case 2147483647:
                                                                    var v104 bool = v5
                                                                    _ = v104
                                                                case -130:
                                                                    v2 = 1e10
                                                                default:
                                                                    switch fi17 {
                                                                    case 744:
                                                                        if len(v52) >= 2 {
                                                                            v52 = v52[1:len(v52)]
                                                                        }
                                                                    case -9223372036854775808:
                                                                        fmt.Println(fi21, v2)
                                                                    default:
                                                                        switch 42 {
                                                                        case 262:
                                                                            if v5 {
                                                                                var v105 uint64 = 712
                                                                                if ((^(v105)) == v105) {
                                                                                    v2 -= v2
                                                                                    switch fi17 {
                                                                                    case 0:
                                                                                        var v106 int = fi21
                                                                                        _ = v106
                                                                                    case -1:
                                                                                        v3 = v3
                                                                                    }
                                                                                    v107 := GenS12{F0: -1}
                                                                                    fmt.Println(v3)
                                                                                    _ = v107
                                                                                }
                                                                                _ = v105
                                                                            } else {
                                                                                if (true || (-32768 <= -517)) {
                                                                                    fmt.Println(v3, v102, v5)
                                                                                    fmt.Println(v3, v3, v102)
                                                                                    v102 = (v102 % ((1) | 1))
                                                                                }
                                                                                var v108 uint32 = 0
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                v5 = true
                                                            } else {
                                                                if v5 {
                                                                    if len(v52) >= 2 {
                                                                        v52 = v52[1:len(v52)]
                                                                    }
                                                                    _ = len(v4)
                                                                    v5 = (false && false)
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v109 uint16 = 65535
                                                        v110 := GenS13{F0: fi17, F1: v109}
                                                        if len(v52) > 0 {
                                                            v52[0] = -1
                                                        }
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                case "a":
                                                    delete(v1, 1)
                                                case "foo bar":
                                                    if v5 {
                                                        var v111 float64 = 0.0
                                                        _ = v111
                                                    } else {
                                                        v5 = (!v5)
                                                        switch v3 {
                                                        case "foo bar":
                                                            if false {
                                                                v52 = append(v52, (fi21 / ((fi21) | 1)))
                                                                v2 = 0.0
                                                                if (false && (98 != 2147483647)) {
                                                                    fmt.Println(v102, v5)
                                                                    if v5 {
                                                                        _ = len(v4)
                                                                        v102 = (v102 << (((v102 % ((v102) | 1))) & 63))
                                                                    }
                                                                    v112 := GenS14{F0: 325, F1: 42, F2: v3}
                                                                    var v113 int8 = -64
                                                                    v102 = (^(v102))
                                                                    _ = v112
                                                                    _ = v113
                                                                } else {
                                                                    v114 := GenS2{F0: 1000, F1: -815}
                                                                    _ = v114
                                                                }
                                                                if (-2.25 == 0.0) {
                                                                    if (!v5) {
                                                                        v115 := GenS15{F0: ((^(fi17)) == fi17), F1: 42, F2: 42}
                                                                        if v5 {
                                                                            v116 := GenS7{F0: -127}
                                                                            if (!(!false)) {
                                                                                var v117 uint64 = 2
                                                                                fmt.Println(v102, v102)
                                                                                _ = v117
                                                                            }
                                                                            v3 = v3
                                                                            var v118 int = fi17
                                                                            _ = v116
                                                                            _ = v118
                                                                        }
                                                                        _ = v115
                                                                    } else {
                                                                        v119 := GenS9{F0: 27}
                                                                        var v120 uint8 = 46
                                                                        _ = v119
                                                                        _ = v120
                                                                    }
                                                                }
                                                                if (fi21 == -928) {
                                                                    v121 := GenS4{F0: (!(1 != 255))}
                                                                    v5 = v5
                                                                    _ = v121
                                                                } else {
                                                                    var v122 bool = v5
                                                                    v122 = (!(-746 == 648))
                                                                    var v123 int8 = 42
                                                                    _ = v122
                                                                    _ = v123
                                                                }
                                                                var v124 int64 = 933
                                                                _ = v124
                                                            } else {
                                                                v52 = append(v52, (fi21 % (((fi21 % ((fi21) | 1))) | 1)))
                                                                var v125 uint8 = 128
                                                                var v126 string = "hello"
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                        case "":
                                                            v102 = (v102 / ((v102) | 1))
                                                        case "\n\t":
                                                            if (0.1 <= ((1.5 - v2) / v2)) {
                                                                fmt.Println(v3, v102, fi21)
                                                            } else {
                                                                fmt.Println(v102)
                                                                v127 := GenS16{F0: 1, F1: -68}
                                                                var v128 int64 = 647
                                                                _ = v127
                                                                _ = v128
                                                            }
                                                        default:
                                                            if len(v52) > 0 {
                                                                v52[0] = (^(fi21))
                                                            }
                                                        }
                                                        v2 += v2
                                                    }
                                                }
                                                if (128 < 1) {
                                                    if true {
                                                        v2 -= float64(float32((((float64(0.0) + float64((((float32((1.5 / 1.0)) + 0.1) * -2.25) * ((1.5 / (3.14159 - 0.0)) / (-2.25 - 0.0))))) + ((((((float64((1.5 + 1e10)) / 1.0) / ((v2 + 0.0) + (v2 - 0.0))) + v2) / (v2 * (v2 + 1.5))) + ((v2 / v2) + (1.5 - (0.1 - v2)))) + v2)) - (((float64(-2.25) + v2) - v2) - v2))))
                                                    }
                                                } else {
                                                    v129 := GenS17{F0: 735, F1: (1e10 <= 1e10), F2: 828}
                                                    var v130 string = ""
                                                    if v5 {
                                                        v2 += v2
                                                        fmt.Println(fi21, fi21, fi17)
                                                        if (v130 > v3) {
                                                            v102 -= v102
                                                            var v131 uint16 = 2
                                                            v132 := GenS18{F0: v5}
                                                            _ = v131
                                                            _ = v132
                                                        } else {
                                                            v130 = (v130 + v3)
                                                            v133 := GenS19{F0: (fi21 * (fi17 - (^(fi21))))}
                                                            var v134 uint = (^(v102))
                                                            fmt.Println(v3, v3)
                                                            _ = v133
                                                            _ = v134
                                                        }
                                                    }
                                                    _ = v129
                                                    _ = v130
                                                }
                                                delete(v4, 1)
                                                _ = v102
                                            }
                                            v3 = "hello"
                                            switch v3 {
                                            case "":
                                                v5 = ((fi17 != fi17) || false)
                                            case "\n\t":
                                                if false {
                                                    v135 := GenS20{F0: 242}
                                                    v1[42] = v5
                                                    var v136 bool = v5
                                                    _ = v135
                                                    _ = v136
                                                } else {
                                                    var v137 int64 = -313
                                                    var v138 string = ("a" + (v3 + ""))
                                                    switch (^(fi17)) {
                                                    case -2147483648:
                                                        var v139 uint = 448
                                                        _ = v139
                                                    }
                                                    _ = v137
                                                    _ = v138
                                                }
                                            case "a":
                                                v3 = v3
                                            default:
                                                var v140 uint16 = 103
                                                _ = v140
                                            }
                                        }
                                    }
                                } else {
                                    var v141 float32 = 0.1
                                    v142 := GenS21{F0: -1, F1: (((v3 + "foo bar") + "a") + ("hello" + v3))}
                                    _ = len(v1)
                                    _ = v141
                                    _ = v142
                                }
                            case 32768:
                                v2 = v2
                            default:
                                v3 = "foo bar"
                            }
                            switch v3 {
                            case "foo bar":
                                var v143 int8 = -51
                                _ = v143
                            default:
                                v144 := GenS1{F0: (v3 + "hello"), F1: fi21}
                                _ = v144
                            }
                            if v5 {
                                if len(v52) >= 2 {
                                    v52 = v52[1:len(v52)]
                                }
                            } else {
                                var v145 uint32 = 42
                                v1[(v145 + (v145 - v145))] = (((v5 && v5) || (!true)) || v5)
                                v146 := GenS22{F0: v145, F1: 128}
                                if ((v5 || v5) || (((float64(((3.14159 * 0.1) / ((0.1 / 3.14159) - float32(1e10)))) - -2.25) * -2.25) != v2)) {
                                    var v147 int64 = -967
                                    _ = v147
                                }
                                v145 += 569
                                _ = v145
                                _ = v146
                            }
                            v2 += v2
                            var v148 float64 = 0.0
                            if len(v52) > 0 {
                                v52[0] = fi21
                            }
                            v5 = false
                            v149 := GenS6{F0: 774}
                            v150 := GenS23{F0: 496, F1: 468, F2: -191}
                            switch "foo bar" {
                            case "":
                                v151 := GenS1{F0: v3, F1: fi17}
                                _ = v151
                            default:
                                v152 := GenS14{F0: 42, F1: 2, F2: v3}
                                _ = v152
                            }
                            _ = v148
                            _ = v149
                            _ = v150
                        }
                    case 32767:
                        v153 := GenS24{F0: v3, F1: 4294967295, F2: 486}
                        var v154 float32 = 1e10
                        _ = v153
                        _ = v154
                    }
                    var v155 uint = 908
                    switch fi21 {
                    case 1:
                        var v156 int = (fi21 - -2147483648)
                        _ = v156
                    case 128:
                        v4[0] = 0
                    }
                    fmt.Println(v3, v155)
                    v3 = v3
                    _ = v155
                    _ = v52
                } else {
                    fmt.Println(fi17, fi21, fi17)
                    if (0.0 == (((0.1 + 1e10) - 1e10) * -2.25)) {
                        switch fi17 {
                        case 4294967295:
                            var v157 bool = v5
                            _ = v157
                        case -9223372036854775808:
                            v158 := v4[2147483648]
                            _ = v158
                        }
                    }
                    if (42 != 42) {
                        v159 := []int{9223372036854775807, fi17}
                        _ = v159
                    } else {
                        if v5 {
                            var v160 bool = (99 != 162)
                            _ = v160
                        }
                    }
                }
                v161 := []int{(^(fi21))}
                _ = fi21
                _ = v161
            }
            switch (fi17 * fi17) {
            case -2147483648:
                v162 := v1[4294967295]
                _ = v162
            case 4294967295:
                delete(v4, 656)
            default:
                if false {
                    if (!true) {
                        for fi163 := 0; fi163 < 10; fi163++ {
                            var v164 bool = v5
                            v165 := GenS2{F0: -32768, F1: -1}
                            _ = fi163
                            _ = v164
                            _ = v165
                        }
                        v3 = (v3 + v3)
                        fmt.Println(fi17, v3)
                    } else {
                        if v5 {
                            var v166 int8 = 0
                            switch -698 {
                            case 0:
                                fmt.Println(fi17, fi17)
                            }
                            _ = v166
                        } else {
                            v5 = ((867 != 1) || v5)
                            for fi167 := 0; fi167 < 2; fi167++ {
                                if (-9223372036854775808 < (fi167 - fi17)) {
                                    if (1e10 <= (0.1 * (0.1 * 0.1))) {
                                        switch 32768 {
                                        case -2147483648:
                                            v5 = v5
                                        case 0:
                                            v2 -= (float64((-2.25 / 1.0)) + 1e10)
                                        case -9223372036854775808:
                                            v5 = v5
                                        default:
                                            v168 := []int{fi17, (fi167 & (-(fi17))), (fi167 / (((fi167 & fi17)) | 1))}
                                            _ = v168
                                        }
                                    } else {
                                        v169 := []int{}
                                        if false {
                                            var v170 int8 = 42
                                            var v171 uint64 = 42
                                            var v172 int = (fi17 >> ((18446744073709551615) & 63))
                                            _ = v170
                                            _ = v171
                                            _ = v172
                                        } else {
                                            if len(v169) >= 2 {
                                                v169 = v169[1:len(v169)]
                                            }
                                            _ = len(v4)
                                            v173 := GenS12{F0: 12}
                                            _ = v173
                                        }
                                        v2 = v2
                                        _ = v169
                                    }
                                } else {
                                    v174 := GenS25{F0: (v3 + v3), F1: 4294967295, F2: -454}
                                    _ = v174
                                }
                                v5 = (!v5)
                                _ = fi167
                            }
                            v2 -= v2
                        }
                        v175 := GenS26{F0: 42, F1: fi17}
                        _ = v175
                    }
                    var v176 int16 = 1000
                    _ = v176
                } else {
                    switch (fi17 + 162) {
                    case 32768:
                        var v177 float64 = v2
                        _ = v177
                    case 309:
                        v3 = (v3 + "a")
                    }
                }
            }
            _ = fi17
        }
        var v178 float32 = (0.1 - (((0.0 - 0.0) / 1.5) / 1.0))
        _ = v178
    }
    var v179 uint64 = 18446744073709551615
    v180 := []int{}
    switch "a" {
    case "a":
        fmt.Println(v3, v3, v3)
    case "foo bar":
        if ((v5 && v5) || ((2147483647 == -138) || (1 < 130))) {
            if len(v180) >= 2 {
                v180 = v180[1:len(v180)]
            }
            v2 -= ((float64(1e10) - (float64(-2.25) * v2)) / ((float64(0.1) + ((v2 / 1.5) + v2)) / v2))
            v4[2] = -9223372036854775808
        }
    case "hello":
        v2 = v2
    }
    switch v3 {
    case "":
        var v181 int32 = 0
        _ = v181
    case "hello":
        v179 = v179
    case "\n\t":
        var v182 string = ("\n\t" + "foo bar")
        _ = v182
    }
    v5 = true
    var v183 uint32 = 2147483647
    if v5 {
        var v184 uint8 = 42
        _ = v184
    }
    for fi185 := 0; fi185 < 12; fi185++ {
        for fi186 := 0; fi186 < 7; fi186++ {
            if v5 {
                var v187 uint = 2
                _ = v187
            } else {
                v5 = ("foo bar" != ("\n\t" + "foo bar"))
                v3 = v3
            }
            if true {
                switch ((v3 + "foo bar") + "a") {
                case "a":
                    v188 := GenS27{F0: fi185, F1: v183}
                    _ = v188
                case "":
                    if (-2147483648 >= -1) {
                        var v189 uint64 = 160
                        v179 -= 435
                        var v190 uint32 = 2147483648
                        _ = v189
                        _ = v190
                    } else {
                        var v191 uint = 42
                        _ = v191
                    }
                default:
                    v183 = (v183 * 881)
                }
            }
            fmt.Println(v5)
            _ = fi186
        }
        _ = fi185
    }
    v192 := GenS28{F0: 1000, F1: 128}
    for fi193 := 0; fi193 < 6; fi193++ {
        fmt.Println(fi193)
        _ = fi193
    }
    fmt.Println(v2, v3)
    if v5 {
        v179 -= (v179 & 846)
        for fi194 := 0; fi194 < 6; fi194++ {
            if (v179 <= (v179 - v179)) {
                for fi195 := 0; fi195 < 11; fi195++ {
                    fmt.Println(v5)
                    if v5 {
                        if true {
                            fmt.Println(fi195, fi194, fi194)
                            fmt.Println(fi194, v5)
                        }
                        var v196 int8 = 0
                        _ = v196
                    } else {
                        fmt.Println(fi195)
                    }
                    v1[(v183 >> ((1) & 63))] = v5
                    _ = fi195
                }
            }
            v180 = append(v180, (fi194 - -295))
            v2 = (v2 + v2)
            _ = fi194
        }
        var v197 uint32 = (v183 ^ v183)
        _ = v197
    } else {
        v5 = v5
    }
    v5 = false
    switch (v3 + ((v3 + "\n\t") + v3)) {
    case "\n\t":
        v2 = v2
    case "foo bar":
        v183 = (v183 << ((308) & 63))
    default:
        if false {
            var v198 bool = (49 == 96)
            _ = v198
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v179:", v179)
    fmt.Println("slicelen v180:", len(v180), v180)
    fmt.Println("v183:", v183)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v179
    _ = v180
    _ = v183
    _ = v192
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
