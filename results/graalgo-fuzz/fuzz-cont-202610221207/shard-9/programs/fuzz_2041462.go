package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint; F1 uint8; F2 uint }
type GenS8 struct { F0 bool; F1 uint16 }
type GenS9 struct { F0 string; F1 string }
type GenS10 struct { F0 uint8; F1 uint32; F2 string }
type GenS11 struct { F0 int; F1 int16; F2 int64 }
type GenS12 struct { F0 uint32; F1 string }
type GenS13 struct { F0 int64; F1 uint64 }
type GenS14 struct { F0 uint; F1 uint32; F2 bool }
type GenS15 struct { F0 uint64; F1 int32 }
type GenS16 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS17 struct { F0 int16; F1 uint; F2 uint16 }
type GenS18 struct { F0 uint }
func genFn0(p0 uint64) uint16 {
    p0 -= p0
    var v7 uint = 0
    v7 += v7
    var v8 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 int64, p1 int16) int32 {
    var v9 int64 = (p0 + -988)
    if (true && false) {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int8 = 42
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := []int{fi13, (^(fi13))}
                var v15 int8 = 47
                if (!(!((v15 | (v15 ^ v15)) > v15))) {
                    if ("" >= "\n\t") {
                        fmt.Println(fi12, v15)
                        switch (-(fi13)) {
                        case 0:
                            v9 -= p0
                        case 4294967295:
                            v15 = v15
                        default:
                            p1 = 32767
                        }
                        var v16 bool = false
                        _ = v16
                    }
                    switch fi12 {
                    case 255:
                        v15 = -128
                    }
                } else {
                    var v17 uint32 = 2147483648
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            if (2147483647 == 65535) {
                v9 -= p0
                var v18 uint = 18446744073709551615
                v19 := GenS1{F0: 18446744073709551615, F1: 0, F2: 1}
                _ = v18
                _ = v19
            } else {
                v20 := []int{4294967295}
                _ = v20
            }
            _ = fi12
        }
        var v21 float64 = (1.5 * ((3.14159 - 0.0) - (((3.14159 * 3.14159) - float64(1e10)) - (((0.0 / -2.25) / float64(1e10)) + 0.1))))
        _ = v21
    } else {
        var v22 uint = 1
        _ = v22
    }
    v9 = v9
    p1 = 32767
    var v23 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v23
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 2
    var v3 string = ("hello" + "\n\t")
    var v4 int32 = 0
    var v5 int64 = -197
    var v6 float32 = ((3.14159 - 0.0) / 3.14159)
    v2 += 42
    for fi24 := 0; fi24 < 3; fi24++ {
        fmt.Println(fi24)
        v1[(^(fi24))] = ("" + (v3 + v3))
        for fi25 := 0; fi25 < 5; fi25++ {
            switch (v3 + "foo bar") {
            case "\n\t":
                v3 = (v3 + v3)
            case "hello":
                if true {
                    v26 := GenS2{F0: -1}
                    var v27 uint16 = genFn0(42)
                    _ = v26
                    _ = v27
                }
            case "a":
                v3 = v3
            default:
                v28 := GenS2{F0: (v4 ^ (v4 - 2147483647))}
                _ = v28
            }
            _ = len(v1)
            if true {
                v6 = v6
                v4 += 2147483647
                if (v3 != (v3 + "a")) {
                    fmt.Println(fi24)
                    v3 = v3
                }
            } else {
                switch ((v3 + v3) + "a") {
                case "foo bar":
                    if (false || true) {
                        var v29 uint32 = 864
                        v2 = (v2 - genFn0(974))
                        _ = v29
                    }
                default:
                    var v30 uint = 2
                    _ = v30
                }
                switch v3 {
                case "foo bar":
                    if (0 > 4294967295) {
                        switch (fi24 & (^(fi24))) {
                        case -2147483648:
                            v5 = (v5 >> ((4294967295) & 63))
                            v31 := GenS2{F0: 2147483647}
                            v6 = -2.25
                            _ = v31
                        case 127:
                            v3 = "a"
                        }
                    }
                case "\n\t":
                    v32 := GenS3{F0: 524, F1: (!(!(false && ((!false) && ((v2 >> ((255) & 63)) != (v2 / (((v2 | v2)) | 1))))))), F2: "a"}
                    _ = v32
                }
            }
            _ = fi25
        }
        _ = fi24
    }
    var v33 int8 = 43
    var v34 string = ""
    var v35 int8 = (v33 & v33)
    v3 = v3
    v36 := GenS4{F0: 4294967295, F1: (v2 - (v2 * genFn0(42)))}
    var v37 int16 = 926
    v38 := GenS4{F0: 762, F1: v2}
    v5 -= v5
    v39 := GenS5{F0: 42}
    v40 := GenS1{F0: 18446744073709551615, F1: 560, F2: 98}
    v41 := GenS6{F0: "a", F1: int64(v33)}
    if ((4294967295 <= 42) || false) {
        v2 = 32767
    }
    v2 += genFn0(2)
    if false {
        var v42 int = 30
        _ = v42
    } else {
        switch v3 {
        case "":
            v43 := GenS7{F0: 0, F1: 2, F2: 30}
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        if false {
            if ((((1e10 + (0.1 / (1.5 / (-2.25 - float64(0.1))))) / 1.0) / 0.1) >= 1.5) {
                v35 += (v35 >> ((2) & 63))
                for fi45 := 0; fi45 < 6; fi45++ {
                    var v46 int64 = v5
                    var v47 uint8 = 144
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            } else {
                if (3.14159 < (float32(0.1) + v6)) {
                    if false {
                        var v48 uint16 = v2
                        for fi49 := 0; fi49 < 6; fi49++ {
                            if false {
                                v6 -= 1.5
                                v50 := []int{(^(fi44)), fi49, 128, (fi44 >> ((4294967295) & 63))}
                                if true {
                                    v34 = "\n\t"
                                    v5 = 482
                                }
                                v37 -= (v37 * (^(v37)))
                                var v51 uint32 = 4294967295
                                if (4294967295 > 1) {
                                    v1[1] = v34
                                    switch (-(fi44)) {
                                    case 4294967295:
                                        if (!true) {
                                            if (uint(v51) == 386) {
                                                var v52 bool = true
                                                _ = v52
                                            }
                                        }
                                    case -9223372036854775808:
                                        var v53 float64 = 1.5
                                        _ = v53
                                    default:
                                        if false {
                                            if (!true) {
                                                switch 42 {
                                                case -9223372036854775808:
                                                    switch (fi49 >> ((556) & 63)) {
                                                    case 32768:
                                                        if (false && (1e10 > 0.0)) {
                                                            var v54 string = v34
                                                            _ = v54
                                                        }
                                                    case 128:
                                                        v37 += (v37 ^ (v37 % (((v37 - (v37 * (v37 & (v37 | (v37 % ((v37) | 1))))))) | 1)))
                                                        v55 := GenS1{F0: 4294967295, F1: (^(v51)), F2: 68}
                                                        v56 := GenS1{F0: 4294967295, F1: 110, F2: 188}
                                                        v5 = (^(v5))
                                                        _ = v55
                                                        _ = v56
                                                    default:
                                                        if (((v37 >> ((606) & 63)) < (v37 ^ v37)) && ((0.1 * v6) <= (v6 * (v6 * 0.1)))) {
                                                            v57 := GenS6{F0: v3, F1: v5}
                                                            _ = len(v1)
                                                            _ = v57
                                                        } else {
                                                            var v58 int64 = -1
                                                            var v59 int = fi44
                                                            v60 := GenS8{F0: true, F1: 65535}
                                                            _ = v58
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                case 2147483647:
                                                    var v61 uint16 = v48
                                                    _ = v61
                                                case 127:
                                                    var v62 uint8 = 1
                                                    _ = v62
                                                }
                                                v63 := GenS9{F0: (v34 + v3), F1: (((v34 + (("\n\t" + ("" + ((v3 + v34) + v34))) + "hello")) + (v3 + v34)) + v3)}
                                                var v64 bool = (1 == 54)
                                                _ = v63
                                                _ = v64
                                            }
                                            v50 = append(v50, (fi49 | fi49))
                                            switch v3 {
                                            case "":
                                                delete(v1, (^(fi44)))
                                            case "a":
                                                v4 -= (v4 >> ((255) & 63))
                                            }
                                        }
                                    }
                                }
                                _ = v50
                                _ = v51
                            } else {
                                var v65 string = v3
                                var v66 bool = false
                                _ = v65
                                _ = v66
                            }
                            _ = fi49
                        }
                        _ = v48
                    } else {
                        v6 = ((0.1 + ((1.5 - float32(float64(((v6 + v6) - (v6 * v6))))) + -2.25)) - v6)
                        v67 := GenS10{F0: 1, F1: 4294967295, F2: (v3 + "hello")}
                        _ = v67
                    }
                } else {
                    delete(v1, (fi44 + 128))
                    var v68 float64 = float64(v6)
                    _ = v68
                }
                v69 := GenS8{F0: false, F1: genFn0(855)}
                var v70 uint = 283
                _ = v69
                _ = v70
            }
            for fi71 := 0; fi71 < 11; fi71++ {
                switch fi44 {
                case 42:
                    _ = len(v1)
                case 9223372036854775807:
                    fmt.Println(v4)
                case 2147483647:
                    var v72 float32 = 0.0
                    _ = v72
                }
                if true {
                    var v73 int32 = -622
                    if ((v37 ^ 42) < -692) {
                        v74 := []int{9223372036854775807, int(v2), fi44}
                        if (true && (false && true)) {
                            var v75 uint64 = 4294967295
                            _ = v75
                        }
                        v37 -= (v37 % (((v37 + v37)) | 1))
                        _ = v74
                    } else {
                        v1[fi71] = "hello"
                        switch -9223372036854775808 {
                        case 42:
                            v76 := GenS11{F0: (fi44 | 128), F1: 42, F2: -1}
                            _ = v76
                        }
                        delete(v1, fi44)
                        switch "" {
                        case "":
                            v6 += v6
                        case "\n\t":
                            switch "hello" {
                            case "\n\t":
                                switch 2147483647 {
                                case 42:
                                    v77 := []int{fi44, 32767}
                                    _ = v77
                                case 0:
                                    v78 := GenS12{F0: 0, F1: v34}
                                    _ = v78
                                }
                            case "hello":
                                v79 := GenS13{F0: v5, F1: 18446744073709551615}
                                _ = v79
                            default:
                                if true {
                                    switch v3 {
                                    case "":
                                        if (genFn1(v5, v37) != v4) {
                                            v80 := GenS6{F0: (v34 + "\n\t"), F1: v5}
                                            v81 := GenS14{F0: 0, F1: 2147483648, F2: (135 >= 480)}
                                            _ = v80
                                            _ = v81
                                        }
                                    }
                                    v5 = v5
                                    fmt.Println(v34, fi44)
                                } else {
                                    fmt.Println(v34, v3)
                                    if (42 > 127) {
                                        v1[807] = v34
                                    } else {
                                        v34 = (v34 + ("hello" + (v34 + v34)))
                                        v82 := GenS4{F0: 732, F1: genFn0(4294967295)}
                                        fmt.Println(fi71, v6, v2)
                                        _ = v82
                                    }
                                    v83 := []int{fi44}
                                    _ = v83
                                }
                            }
                        case "hello":
                            var v84 uint8 = 1
                            _ = v84
                        default:
                            if false {
                                v85 := []int{}
                                v3 = ((v3 + "hello") + v34)
                                var v86 string = ("a" + ("\n\t" + "foo bar"))
                                var v87 int16 = (v37 & v37)
                                if (false || false) {
                                    if (v35 > 100) {
                                        fmt.Println(v33, fi44, v73)
                                        v3 = "foo bar"
                                        switch (v86 + v86) {
                                        case "\n\t":
                                            switch -1 {
                                            case 0:
                                                var v88 float64 = 0.1
                                                _ = v88
                                            case -916:
                                                var v89 float32 = 0.0
                                                _ = v89
                                            case 1:
                                                v33 += v35
                                            }
                                        case "foo bar":
                                            switch 600 {
                                            case 283:
                                                if (42 > 21) {
                                                    v87 = v37
                                                    switch int(v2) {
                                                    case 4294967295:
                                                        fmt.Println(fi44, v86, v86)
                                                    default:
                                                        var v90 float64 = (-2.25 * 3.14159)
                                                        _ = v90
                                                    }
                                                }
                                                v73 += int32(fi44)
                                                switch v34 {
                                                case "":
                                                    v91 := GenS7{F0: 2, F1: 127, F2: 382}
                                                    _ = v91
                                                case "\n\t":
                                                    v73 = genFn1(v5, v37)
                                                case "a":
                                                    var v92 string = "foo bar"
                                                    _ = v92
                                                default:
                                                    if false {
                                                        if false {
                                                            if (0.1 >= 3.14159) {
                                                                var v93 bool = (!((!false) || true))
                                                                v37 = (-(v87))
                                                                fmt.Println(v34, v37, v3)
                                                                _ = v93
                                                            }
                                                            v94 := v1[2147483647]
                                                            v95 := GenS15{F0: 18446744073709551615, F1: v73}
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            var v96 int64 = (v5 << ((65535) & 63))
                                                            v97 := GenS6{F0: (v3 + v86), F1: -9223372036854775808}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        v98 := GenS16{F0: 42, F1: false, F2: 2147483647}
                                                        _ = v98
                                                    } else {
                                                        switch (fi44 * -2147483648) {
                                                        case 0:
                                                            fmt.Println(v73, v34)
                                                        case 9223372036854775807:
                                                            switch -50 {
                                                            case 9223372036854775807:
                                                                v99 := GenS17{F0: int16(v5), F1: 4294967295, F2: genFn0(1)}
                                                                _ = len(v1)
                                                                if (genFn1(-991, 1) != 1) {
                                                                    v33 -= (v33 & -1)
                                                                    v100 := GenS18{F0: 0}
                                                                    if (255 >= 18446744073709551615) {
                                                                        v2 = (v2 % (((^(v2))) | 1))
                                                                        var v101 int32 = (v4 << ((65535) & 63))
                                                                        if false {
                                                                            var v102 uint32 = 9
                                                                            fmt.Println(v35, v5, v3)
                                                                            if len(v85) >= 2 {
                                                                                v85 = v85[1:len(v85)]
                                                                            }
                                                                            _ = v102
                                                                        }
                                                                        _ = v101
                                                                    }
                                                                    v5 -= v5
                                                                    _ = v100
                                                                }
                                                                v37 = (v87 ^ (v87 ^ (v87 >> ((989) & 63))))
                                                                _ = v99
                                                            }
                                                            fmt.Println(v3, fi44)
                                                        default:
                                                            if len(v85) > 0 {
                                                                v85[0] = (fi44 << ((65535) & 63))
                                                            }
                                                        }
                                                        var v103 int16 = 32767
                                                        _ = v103
                                                    }
                                                }
                                                v1[fi44] = v34
                                                if (!false) {
                                                    v37 -= 32767
                                                }
                                            default:
                                                fmt.Println(v34, v73)
                                            }
                                        }
                                    }
                                } else {
                                    var v104 float32 = v6
                                    _ = v104
                                }
                                _ = v85
                                _ = v86
                                _ = v87
                            } else {
                                if false {
                                    var v105 uint8 = 42
                                    _ = len(v1)
                                    v33 = (v33 + 4)
                                    v106 := GenS17{F0: 32767, F1: 189, F2: v2}
                                    v107 := GenS7{F0: 0, F1: (v105 >> ((255) & 63)), F2: 2}
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                } else {
                                    var v108 bool = ((!false) || (!false))
                                    if v108 {
                                        if v108 {
                                            var v109 int = fi44
                                            fmt.Println(v37)
                                            if true {
                                                var v110 uint16 = v2
                                                fmt.Println(v4)
                                                if true {
                                                    v111 := v1[fi71]
                                                    if false {
                                                        v33 -= v33
                                                        v112 := GenS5{F0: 627}
                                                        _ = v112
                                                    } else {
                                                        var v113 bool = false
                                                        _ = v113
                                                    }
                                                    v35 += v35
                                                    var v114 uint64 = 0
                                                    _ = v111
                                                    _ = v114
                                                }
                                                _ = v110
                                            } else {
                                                var v115 bool = ((2 == 101) && true)
                                                switch fi44 {
                                                case 128:
                                                    v5 -= 950
                                                }
                                                var v116 string = v3
                                                _ = v115
                                                _ = v116
                                            }
                                            _ = v109
                                        } else {
                                            v73 -= -400
                                            switch fi71 {
                                            case -498:
                                                var v117 int64 = v5
                                                _ = v117
                                            case 32768:
                                                var v118 uint64 = 4294967295
                                                _ = v118
                                            }
                                        }
                                        v5 = (v5 << ((347) & 63))
                                    } else {
                                        if v108 {
                                            v119 := []int{fi71}
                                            _ = v119
                                        } else {
                                            v120 := []int{127, fi44, int(v33)}
                                            var v121 int32 = genFn1(v5, (v37 & 704))
                                            v122 := GenS8{F0: (false || true), F1: 436}
                                            _ = v120
                                            _ = v121
                                            _ = v122
                                        }
                                        v5 = 52
                                        v35 = v33
                                    }
                                    v73 = -688
                                    v123 := []int{(^(fi44)), (fi44 - fi44), 4294967295, fi44}
                                    _ = v108
                                    _ = v123
                                }
                                v4 = -652
                            }
                        }
                        var v124 int32 = (v73 + genFn1(v5, -1))
                        v3 = (v3 + "hello")
                        var v125 bool = true
                        _ = v124
                        _ = v125
                    }
                    _ = v73
                } else {
                    v126 := GenS3{F0: 173, F1: (true || false), F2: v34}
                    v4 = genFn1((v5 & v5), (v37 + v37))
                    v127 := GenS5{F0: 4294967295}
                    v2 = (v2 + (v2 / ((220) | 1)))
                    fmt.Println(v4, v34)
                    _ = v126
                    _ = v127
                }
                _ = fi71
            }
        }
        fmt.Println(v34, v3)
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
