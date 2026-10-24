package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int32; F1 uint64; F2 bool }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64 }
func genFn0() int {
    v7 := GenS1{F0: 1, F1: -202}
    v8 := GenS1{F0: 473, F1: -93}
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 int32 = -333
        v10 -= v10
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int16 = 109
        _ = fi11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint16 = 865
            var v15 int32 = 870
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 uint32 = 2
        _ = v16
    } else {
        var v17 uint16 = 65535
        var v18 uint32 = 42
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 4; fi20++ {
                var v21 string = ""
                fmt.Println(fi19)
                _ = fi20
                _ = v21
            }
            if true {
                switch "a" {
                case "\n\t":
                    v18 = (v18 | 4294967295)
                default:
                    var v22 float32 = 1e10
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 float32 = 1.5
                    var v25 uint16 = v17
                    var v26 int = 127
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                v17 -= (v17 | 42)
            } else {
                fmt.Println(v18, v17, v18)
                var v27 uint32 = v18
                if true {
                    var v28 bool = (867 < 4294967295)
                    v29 := []int{fi19}
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                    _ = v28
                    _ = v29
                } else {
                    v27 -= 65535
                    v30 := GenS1{F0: 26, F1: fi19}
                    fmt.Println(v17, v27)
                    _ = v30
                }
                _ = v27
            }
            v31 := GenS1{F0: 630, F1: 310}
            _ = fi19
            _ = v31
        }
        _ = v17
        _ = v18
    }
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int = 42
    v2 := make(map[string]uint32)
    var v3 int64 = -1
    var v4 float64 = (3.14159 * 0.0)
    v5 := make(map[int]uint32)
    v6 := []string{"", "hello", "hello", ((("\n\t" + ("\n\t" + "a")) + "hello") + "a")}
    v1 = genFn0()
    if false {
        v3 = (v3 << ((255) & 63))
        fmt.Println(v1)
        if (4294967295 < 627) {
            fmt.Println(v3, v4, v3)
        } else {
            var v32 uint = 255
            _ = v32
        }
    }
    v1 -= v1
    var v33 uint8 = 0
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 int64 = v3
        var v36 uint64 = 285
        v36 = (v36 << ((2) & 63))
        _ = fi34
        _ = v35
        _ = v36
    }
    switch "a" {
    case "a":
        var v37 int64 = v3
        _ = v37
    }
    v38 := v5[-528]
    if (true || (("\n\t" + "a") >= ("a" + "hello"))) {
        if true {
            for fi39 := 0; fi39 < 5; fi39++ {
                v6 = append(v6, "")
                _ = fi39
            }
            fmt.Println(v3, v38)
        }
        if true {
            v40 := GenS1{F0: 18446744073709551615, F1: v1}
            _ = v40
        } else {
            v41 := GenS2{F0: 65535, F1: 1, F2: 214}
            _ = v41
        }
    }
    var v42 uint8 = 0
    var v43 uint = 255
    switch (v1 - v1) {
    case 2147483647:
        var v44 int32 = 759
        _ = v44
    default:
        var v45 int = (-(v1))
        _ = v45
    }
    v2[("hello" + "foo bar")] = 545
    switch (^(v1)) {
    case -2147483648:
        if (!(false && ((100 != 127) || true))) {
            v46 := GenS3{F0: 0, F1: false, F2: v42}
            if true {
                v3 = (v3 - v3)
                switch "foo bar" {
                case "":
                    v47 := GenS4{F0: -128, F1: 42, F2: 42}
                    _ = v47
                case "foo bar":
                    v48 := v5[(^(v1))]
                    _ = v48
                default:
                    v43 = 2
                }
            }
            v49 := GenS2{F0: (v43 >> (((^(v43))) & 63)), F1: 32767, F2: 42}
            _ = v46
            _ = v49
        } else {
            v50 := GenS2{F0: v43, F1: 42, F2: 4294967295}
            var v51 uint32 = v38
            v6 = append(v6, "a")
            _ = v50
            _ = v51
        }
    }
    v33 = (v42 - v42)
    switch (^(v1)) {
    case 2147483647:
        v5[32768] = 1
    case 37:
        if (true || true) {
            for fi52 := 0; fi52 < 13; fi52++ {
                if (!true) {
                    v33 = (v33 - (v33 - 128))
                    switch (v1 ^ int(v43)) {
                    case 128:
                        v53 := GenS4{F0: 9, F1: 1, F2: 2}
                        _ = v53
                    case -9223372036854775808:
                        switch "a" {
                        case "\n\t":
                            switch (v1 << ((uint(v38)) & 63)) {
                            case -441:
                                v42 += (^(v33))
                            case 42:
                                if (false && false) {
                                    for fi54 := 0; fi54 < 2; fi54++ {
                                        fmt.Println(v1, v33, fi54)
                                        v1 -= (fi52 - 55)
                                        if (true || false) {
                                            var v55 bool = false
                                            _ = v55
                                        } else {
                                            var v56 uint64 = 309
                                            var v57 int64 = (v3 & (v3 + 1))
                                            v6 = append(v6, "a")
                                            _ = v56
                                            _ = v57
                                        }
                                        var v58 int8 = -19
                                        _ = fi54
                                        _ = v58
                                    }
                                }
                            case 788:
                                fmt.Println(v42, v42)
                            }
                        case "hello":
                            if true {
                                v59 := GenS5{F0: 0}
                                var v60 int8 = 127
                                _ = v59
                                _ = v60
                            }
                        default:
                            for fi61 := 0; fi61 < 4; fi61++ {
                                v1 = (-(fi61))
                                var v62 int32 = int32(v38)
                                _ = fi61
                                _ = v62
                            }
                        }
                    case 127:
                        v6 = append(v6, ("" + "foo bar"))
                    default:
                        if (true || false) {
                            v63 := GenS4{F0: 24, F1: 116, F2: 772}
                            switch "foo bar" {
                            case "\n\t":
                                for fi64 := 0; fi64 < 10; fi64++ {
                                    v38 = (v38 ^ v38)
                                    var v65 float64 = 0.0
                                    _ = fi64
                                    _ = v65
                                }
                            default:
                                v66 := GenS3{F0: -1, F1: (32767 > 0), F2: 1}
                                _ = v66
                            }
                            _ = v63
                        } else {
                            var v67 uint32 = (v38 | v38)
                            if true {
                                for fi68 := 0; fi68 < 17; fi68++ {
                                    v42 -= (v33 >> (((v43 + v43)) & 63))
                                    v4 -= v4
                                    _ = fi68
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if ((-1 == 883) && false) {
                                    if ((!(false || false)) || (false && false)) {
                                        switch "foo bar" {
                                        case "\n\t":
                                            _ = len(v6)
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = "foo bar"
                                            }
                                        }
                                        fmt.Println(v4, v38, v33)
                                    }
                                    v67 -= (v38 % ((2) | 1))
                                }
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = ""
                                }
                                if ((v38 * 65535) > (v38 * (v38 - v38))) {
                                    v69 := GenS5{F0: 105}
                                    if false {
                                        v2["hello"] = (v38 | (v67 | uint32(v43)))
                                    }
                                    _ = v69
                                } else {
                                    switch fi52 {
                                    case -436:
                                        v33 = 42
                                    case 32767:
                                        for fi70 := 0; fi70 < 7; fi70++ {
                                            var v71 int = (fi52 - genFn0())
                                            _ = fi70
                                            _ = v71
                                        }
                                    }
                                    for fi72 := 0; fi72 < 16; fi72++ {
                                        switch "hello" {
                                        case "foo bar":
                                            v73 := GenS5{F0: 109}
                                            _ = v73
                                        case "hello":
                                            v33 = v42
                                        case "a":
                                            var v74 uint64 = 42
                                            _ = v74
                                        default:
                                            v33 = (v33 << ((v43) & 63))
                                        }
                                        if (v3 >= v3) {
                                            if ((^(v43)) != v43) {
                                                var v75 uint8 = v33
                                                v76 := GenS1{F0: (v43 ^ v43), F1: 2147483647}
                                                _ = v75
                                                _ = v76
                                            } else {
                                                v1 += -9223372036854775808
                                                fmt.Println(v38)
                                                var v77 int32 = 214
                                                _ = v77
                                            }
                                            v42 -= 18
                                            switch "a" {
                                            case "a":
                                                fmt.Println(v1, v1)
                                            default:
                                                var v78 int = v1
                                                _ = v78
                                            }
                                        }
                                        v1 -= fi72
                                        _ = fi72
                                    }
                                }
                                fmt.Println(v33, v1)
                            }
                            _ = v67
                        }
                    }
                } else {
                    var v79 uint64 = 2
                    switch "\n\t" {
                    case "hello":
                        for fi80 := 0; fi80 < 1; fi80++ {
                            v81 := GenS6{F0: (v38 / ((0) | 1))}
                            var v82 int32 = -1
                            v83 := GenS5{F0: -1}
                            _ = fi80
                            _ = v81
                            _ = v82
                            _ = v83
                        }
                    case "\n\t":
                        var v84 int64 = (v3 / (((v3 - (v3 << ((v43) & 63)))) | 1))
                        _ = v84
                    case "a":
                        for fi85 := 0; fi85 < 18; fi85++ {
                            switch (v1 << ((255) & 63)) {
                            case -9223372036854775808:
                                if true {
                                    v33 += (v33 + v42)
                                    var v86 uint32 = v38
                                    _ = v86
                                }
                            case 128:
                                v33 = 255
                                var v87 uint = (v43 % ((v43) | 1))
                                _ = v87
                            case 75:
                                var v88 string = "\n\t"
                                _ = v88
                            }
                            var v89 bool = true
                            v90 := GenS7{F0: 42}
                            _ = fi85
                            _ = v89
                            _ = v90
                        }
                    default:
                        if false {
                            for fi91 := 0; fi91 < 4; fi91++ {
                                if (42 > 2147483647) {
                                    if true {
                                        fmt.Println(v3)
                                        v2[(("a" + "foo bar") + "")] = v38
                                        v33 += v42
                                        v92 := GenS8{F0: int32(v79), F1: uint64(v42), F2: false}
                                        _ = v92
                                    } else {
                                        var v93 int8 = 100
                                        var v94 float64 = v4
                                        v95 := GenS3{F0: -577, F1: false, F2: v42}
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    }
                                    if true {
                                        var v96 uint8 = 42
                                        _ = v96
                                    }
                                    if false {
                                        v79 -= v79
                                    } else {
                                        v97 := GenS2{F0: (^(v43)), F1: 417, F2: 2}
                                        var v98 float64 = v4
                                        _ = v97
                                        _ = v98
                                    }
                                    if true {
                                        _ = len(v5)
                                    }
                                    v5[genFn0()] = (^(v38))
                                    v43 += (v43 + v43)
                                    v99 := GenS2{F0: v43, F1: 65535, F2: v79}
                                    v38 = (v38 - (v38 - uint32(v42)))
                                    _ = v99
                                }
                                v100 := GenS9{F0: v79}
                                var v101 int32 = -1
                                _ = fi91
                                _ = v100
                                _ = v101
                            }
                            for fi102 := 0; fi102 < 15; fi102++ {
                                fmt.Println(v38, v43)
                                v103 := GenS7{F0: 477}
                                switch v1 {
                                case 32768:
                                    if ((-2.25 > ((0.1 * 1e10) + 3.14159)) || (false || (((1000 <= 32767) || ((float64(float32((0.0 + (v4 / v4)))) < 1e10) && false)) && (v3 > v3)))) {
                                        v5[(fi52 - 32767)] = 61
                                        var v104 uint16 = 32767
                                        var v105 string = "\n\t"
                                        fmt.Println(fi102, v79)
                                        _ = v104
                                        _ = v105
                                    } else {
                                        v1 = fi52
                                        v2["foo bar"] = 4294967295
                                        v4 = (v4 * v4)
                                        var v106 int64 = (^(v3))
                                        _ = v106
                                    }
                                case 32767:
                                    var v107 uint = v43
                                    _ = v107
                                default:
                                    v38 = v38
                                }
                                _ = fi102
                                _ = v103
                            }
                            if true {
                                v108 := GenS7{F0: 180}
                                for fi109 := 0; fi109 < 4; fi109++ {
                                    fmt.Println(v33)
                                    v5[-750] = v38
                                    _ = fi109
                                }
                                _ = v108
                            } else {
                                if false {
                                    var v110 int32 = 1
                                    for fi111 := 0; fi111 < 11; fi111++ {
                                        var v112 int16 = 415
                                        _ = fi111
                                        _ = v112
                                    }
                                    if (0 >= 42) {
                                        var v113 uint8 = (v42 * v33)
                                        fmt.Println(v4, v42, v4)
                                        if (false || (!true)) {
                                            for fi114 := 0; fi114 < 17; fi114++ {
                                                v42 -= v33
                                                switch (fi52 / ((genFn0()) | 1)) {
                                                case 42:
                                                    v113 += (^(v33))
                                                }
                                                v115 := GenS10{F0: 644}
                                                var v116 string = ("a" + "\n\t")
                                                var v117 string = (v116 + v116)
                                                _ = fi114
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                            }
                                            var v118 uint8 = 225
                                            _ = v118
                                        }
                                        _ = v113
                                    }
                                    _ = v110
                                }
                                _ = len(v6)
                                switch 32768 {
                                case -1:
                                    var v119 string = (("hello" + "\n\t") + "a")
                                    _ = v119
                                case 0:
                                    v43 = (v43 >> ((v43) & 63))
                                }
                            }
                        }
                    }
                    _ = v79
                }
                var v120 string = ""
                _ = fi52
                _ = v120
            }
        }
    }
    var v121 uint64 = 0
    fmt.Println(v1)
    fmt.Println(v121, v43, v121)
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v33
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
