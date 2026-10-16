package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int8 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 string }
func genFn0(p0 int, p1 uint16, p2 int) int16 {
    v7 := GenS1{F0: p1, F1: -128}
    p0 = 127
    var v8 float64 = (-2.25 + (1.5 + (1.5 / ((1.5 * (3.14159 - 1e10)) + 0.1))))
    var v9 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -465
}

func main() {
    var v1 uint64 = 402
    v2 := make(map[uint32]int)
    v3 := make(map[string]uint32)
    var v4 int64 = -531
    var v5 float64 = 3.14159
    var v6 uint64 = v1
    fmt.Println(v1, v4)
    v10 := []int{}
    if true {
        var v11 int = -2147483648
        _ = v11
    }
    var v12 int32 = -2147483648
    for fi13 := 0; fi13 < 9; fi13++ {
        if (!false) {
            var v14 uint8 = 2
            for fi15 := 0; fi15 < 1; fi15++ {
                switch 1 {
                case 0:
                    switch (fi15 + 32767) {
                    case -222:
                        v4 += (-(v4))
                    case -9223372036854775808:
                        if (-2.25 >= float32(v5)) {
                            var v16 int64 = (v4 >> ((65535) & 63))
                            _ = v16
                        }
                    case -1:
                        v5 = v5
                    }
                default:
                    if true {
                        switch "hello" {
                        case "foo bar":
                            _ = len(v2)
                        case "hello":
                            var v17 uint16 = 1
                            _ = v17
                        case "":
                            switch fi13 {
                            case 1:
                                var v18 bool = true
                                _ = v18
                            default:
                                v10 = append(v10, int(v1))
                            }
                        default:
                            if true {
                                var v19 int16 = genFn0(128, uint16(v14), fi13)
                                fmt.Println(fi15)
                                _ = v19
                            }
                        }
                    }
                }
                _ = fi15
            }
            _ = v14
        } else {
            if true {
                v5 = 0.0
                switch fi13 {
                case -500:
                    v10 = append(v10, (-(fi13)))
                case -2147483648:
                    if false {
                        for fi20 := 0; fi20 < 15; fi20++ {
                            fmt.Println(fi13)
                            v21 := GenS1{F0: 0, F1: -71}
                            v22 := GenS1{F0: 589, F1: 1}
                            fmt.Println(v12)
                            v5 = 0.0
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        var v23 int32 = v12
                        v1 += v1
                        _ = v23
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v24 := v3["foo bar"]
                            _ = v24
                        case "a":
                            v4 = (v4 + (v4 >> ((1) & 63)))
                        default:
                            for fi25 := 0; fi25 < 7; fi25++ {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                if (float64((-2.25 + ((1e10 - 0.0) - 0.0))) == 0.1) {
                                    var v26 uint = 278
                                    _ = v26
                                }
                                var v27 int32 = (v12 | (^(v12)))
                                v27 = -2147483648
                                fmt.Println(v27)
                                _ = fi25
                                _ = v27
                            }
                        }
                    }
                }
                if false {
                    switch (fi13 % (((^(fi13))) | 1)) {
                    case 32767:
                        var v28 uint = 49
                        _ = v28
                    default:
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v1 = (v1 * v6)
                            _ = len(v10)
                            _ = fi29
                        }
                    }
                    var v30 bool = true
                    _ = v30
                }
            } else {
                var v31 int64 = 1
                var v32 bool = (2 <= 2)
                v1 = (^(v6))
                _ = v31
                _ = v32
            }
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 int32 = (-(v12))
                v5 += float64(0.0)
                fmt.Println(v6, v4, fi33)
                _ = fi33
                _ = v34
            }
            if true {
                v6 -= (v6 | 4294967295)
                delete(v2, 1)
            }
        }
        v5 -= ((v5 - -2.25) + (v5 * (0.1 / v5)))
        _ = fi13
    }
    var v35 int32 = -2147483648
    v12 = (-(v35))
    for fi36 := 0; fi36 < 7; fi36++ {
        if (v35 == v35) {
            var v37 float64 = (3.14159 * v5)
            switch "" {
            case "foo bar":
                if (!(!((v12 + (-(v35))) == v12))) {
                    for fi38 := 0; fi38 < 2; fi38++ {
                        var v39 int8 = 42
                        var v40 int32 = -89
                        v12 = v35
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                } else {
                    var v41 int16 = genFn0((-(fi36)), 529, fi36)
                    var v42 int16 = 32767
                    _ = v41
                    _ = v42
                }
            case "hello":
                var v43 int32 = v12
                _ = v43
            default:
                v44 := GenS1{F0: 42, F1: 42}
                _ = v44
            }
            v45 := GenS2{F0: 42, F1: (-(v4)), F2: -22}
            _ = v37
            _ = v45
        }
        _ = fi36
    }
    fmt.Println(v5, v4, v12)
    v46 := GenS2{F0: v4, F1: (v4 | (v4 - v4)), F2: 1}
    switch 1 {
    case 32767:
        v12 -= v35
    case 127:
        if false {
            v47 := GenS3{F0: (((1.5 / 1.5) + 0.0) < ((3.14159 / (1.5 / 1.0)) * 1.5)), F1: 470, F2: 1}
            if (!false) {
                v35 = v12
                if false {
                    switch 42 {
                    case 2147483647:
                        v48 := GenS3{F0: (float64(1e10) > 1.5), F1: 4294967295, F2: 100}
                        _ = v48
                    case 249:
                        switch "" {
                        case "hello":
                            switch "hello" {
                            case "foo bar":
                                if (2 > 1) {
                                    for fi49 := 0; fi49 < 16; fi49++ {
                                        fmt.Println(v12, v1, v12)
                                        _ = fi49
                                    }
                                    v5 = v5
                                    if false {
                                        if ((709 <= v12) || (0.1 >= (-2.25 * 0.1))) {
                                            v6 -= (v1 ^ 0)
                                            for fi50 := 0; fi50 < 2; fi50++ {
                                                if len(v10) > 0 {
                                                    v10[0] = fi50
                                                }
                                                if (true || true) {
                                                    fmt.Println(v4, v35, v1)
                                                    v4 += v4
                                                    v12 = v35
                                                } else {
                                                    v6 -= v6
                                                    v3[("foo bar" + "\n\t")] = 65535
                                                    var v51 int64 = -633
                                                    _ = v51
                                                }
                                                _ = fi50
                                            }
                                            var v52 uint64 = (v6 * v6)
                                            _ = v52
                                        } else {
                                            var v53 uint64 = 900
                                            _ = v53
                                        }
                                        fmt.Println(v5)
                                    }
                                }
                            case "\n\t":
                                fmt.Println(v6, v6)
                            case "":
                                for fi54 := 0; fi54 < 3; fi54++ {
                                    for fi55 := 0; fi55 < 11; fi55++ {
                                        v4 -= v4
                                        if len(v10) > 0 {
                                            v10[0] = fi55
                                        }
                                        _ = fi55
                                    }
                                    var v56 int8 = 1
                                    var v57 int32 = (v35 | (v35 + -1))
                                    _ = fi54
                                    _ = v56
                                    _ = v57
                                }
                            default:
                                var v58 uint = 1
                                _ = v58
                            }
                        default:
                            v5 -= float64((1.5 + -2.25))
                        }
                    case 0:
                        v5 -= (((((v5 + (float64(0.1) - (3.14159 - float64((1.5 + ((3.14159 + -2.25) / (1.5 / 1.0))))))) + v5) / v5) + v5) * (1.5 / (3.14159 * float64(((3.14159 * 0.0) + -2.25)))))
                    default:
                        v59 := GenS3{F0: true, F1: 857, F2: 1}
                        _ = v59
                    }
                    switch "" {
                    case "foo bar":
                        var v60 uint64 = (^(v6))
                        _ = v60
                    case "hello":
                        v3[("hello" + "\n\t")] = 434
                    case "":
                        var v61 int16 = 42
                        _ = v61
                    default:
                        var v62 int16 = 1
                        _ = v62
                    }
                }
                if true {
                    v63 := GenS3{F0: ((v4 + (^(v4))) < v4), F1: 1, F2: 11}
                    _ = v63
                } else {
                    if ((v5 < (v5 * float64(1.5))) || false) {
                        if true {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                v65 := GenS4{F0: 42, F1: 42, F2: (v1 ^ (^(v6)))}
                                _ = fi64
                                _ = v65
                            }
                            if len(v10) > 0 {
                                v10[0] = -119
                            }
                        } else {
                            if (true && (true && (255 < 18446744073709551615))) {
                                if (18446744073709551615 < 18446744073709551615) {
                                    v5 = v5
                                    for fi66 := 0; fi66 < 17; fi66++ {
                                        fmt.Println(v5, v5)
                                        fmt.Println(v4, v4)
                                        v4 += -1
                                        _ = fi66
                                    }
                                    for fi67 := 0; fi67 < 13; fi67++ {
                                        for fi68 := 0; fi68 < 3; fi68++ {
                                            switch (fi67 & (fi67 | (fi67 | (fi68 / ((fi68) | 1))))) {
                                            case 127:
                                                var v69 float64 = v5
                                                _ = v69
                                            case 128:
                                                fmt.Println(v1)
                                            case -1:
                                                var v70 int8 = 0
                                                _ = v70
                                            }
                                            _ = fi68
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v71 := GenS2{F0: 9223372036854775807, F1: -1, F2: -128}
                                            _ = v71
                                        default:
                                            v5 += float64(1e10)
                                        }
                                        _ = fi67
                                    }
                                } else {
                                    fmt.Println(v5)
                                    fmt.Println(v5, v12)
                                    v35 = 1
                                }
                            }
                            if false {
                                var v72 uint = 0
                                var v73 uint16 = 32767
                                if true {
                                    _ = len(v10)
                                    delete(v2, 65535)
                                    for fi74 := 0; fi74 < 16; fi74++ {
                                        var v75 int32 = v35
                                        _ = fi74
                                        _ = v75
                                    }
                                } else {
                                    for fi76 := 0; fi76 < 14; fi76++ {
                                        v4 -= -9223372036854775808
                                        fmt.Println(v35)
                                        _ = fi76
                                    }
                                }
                                _ = v72
                                _ = v73
                            }
                        }
                        for fi77 := 0; fi77 < 17; fi77++ {
                            for fi78 := 0; fi78 < 20; fi78++ {
                                if ((float64(float32(0.0)) / v5) == v5) {
                                    v79 := GenS2{F0: v4, F1: v4, F2: 0}
                                    v35 = v35
                                    _ = v79
                                } else {
                                    v80 := GenS3{F0: (false && true), F1: 4294967295, F2: int8(v12)}
                                    if ((fi78 & (fi77 >> ((42) & 63))) != fi78) {
                                        if false {
                                            fmt.Println(v1, v4, fi77)
                                            v12 = v35
                                        } else {
                                            var v81 uint8 = 128
                                            switch "foo bar" {
                                            case "":
                                                v12 = v35
                                            }
                                            var v82 int = (fi77 + -263)
                                            _ = v81
                                            _ = v82
                                        }
                                    } else {
                                        switch -242 {
                                        case 9223372036854775807:
                                            v12 = (v12 >> ((763) & 63))
                                        case 128:
                                            switch (fi78 * fi78) {
                                            case 128:
                                                _ = len(v2)
                                            case -9223372036854775808:
                                                var v83 int32 = 42
                                                v84 := GenS3{F0: (1.5 < 0.1), F1: 65535, F2: -128}
                                                if false {
                                                    if (false && (!true)) {
                                                        switch (("foo bar" + (("" + "a") + "a")) + "foo bar") {
                                                        case "a":
                                                            delete(v2, 693)
                                                        case "hello":
                                                            if true {
                                                                if (v12 > v35) {
                                                                    v85 := GenS4{F0: 42, F1: 42, F2: 42}
                                                                    _ = len(v2)
                                                                    _ = v85
                                                                }
                                                                if len(v10) >= 2 {
                                                                    v10 = v10[1:len(v10)]
                                                                }
                                                            } else {
                                                                v83 = v35
                                                                var v86 int = fi78
                                                                if false {
                                                                    if (float64(1.5) == (v5 / v5)) {
                                                                        var v87 int16 = 115
                                                                        if len(v10) >= 2 {
                                                                            v10 = v10[1:len(v10)]
                                                                        }
                                                                        v88 := GenS2{F0: (^(v4)), F1: (v4 >> ((65535) & 63)), F2: 42}
                                                                        if len(v10) >= 2 {
                                                                            v10 = v10[1:len(v10)]
                                                                        }
                                                                        _ = len(v10)
                                                                        _ = v87
                                                                        _ = v88
                                                                    } else {
                                                                        v89 := GenS1{F0: 65535, F1: -128}
                                                                        v90 := GenS3{F0: false, F1: 65535, F2: 42}
                                                                        var v91 int32 = (v83 << ((255) & 63))
                                                                        _ = v89
                                                                        _ = v90
                                                                        _ = v91
                                                                    }
                                                                    if len(v10) > 0 {
                                                                        v10[0] = v86
                                                                    }
                                                                } else {
                                                                    var v92 int16 = -386
                                                                    if (573 > 42) {
                                                                        if (692 > 2147483648) {
                                                                            var v93 uint16 = 342
                                                                            var v94 bool = false
                                                                            var v95 float64 = 0.1
                                                                            _ = v93
                                                                            _ = v94
                                                                            _ = v95
                                                                        } else {
                                                                            var v96 uint64 = 2
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                    var v97 bool = (true || false)
                                                                    _ = v92
                                                                    _ = v97
                                                                }
                                                                v98 := GenS1{F0: 1, F1: 100}
                                                                v5 += float64((0.1 / 1.0))
                                                                _ = v86
                                                                _ = v98
                                                            }
                                                        case "\n\t":
                                                            v4 = (v4 - -9223372036854775808)
                                                        default:
                                                            v6 = v6
                                                        }
                                                        v99 := GenS4{F0: 567, F1: 100, F2: v1}
                                                        if (2 != 42) {
                                                            switch fi78 {
                                                            case -110:
                                                                var v100 int = 42
                                                                _ = v100
                                                            case 0:
                                                                _ = len(v3)
                                                            case 128:
                                                                v1 -= v6
                                                            }
                                                            if ((fi77 % (((-(fi77))) | 1)) > 32768) {
                                                                v101 := GenS5{F0: 73}
                                                                v102 := v3[""]
                                                                var v103 string = ""
                                                                _ = v101
                                                                _ = v102
                                                                _ = v103
                                                            } else {
                                                                var v104 int64 = 1
                                                                _ = v104
                                                            }
                                                            switch "a" {
                                                            case "hello":
                                                                v4 += v4
                                                            case "":
                                                                fmt.Println(v5)
                                                            case "foo bar":
                                                                switch fi78 {
                                                                case 9223372036854775807:
                                                                    v83 = (v83 ^ int32(v1))
                                                                default:
                                                                    var v105 string = "\n\t"
                                                                    _ = v105
                                                                }
                                                            }
                                                        } else {
                                                            if false {
                                                                var v106 uint64 = 2
                                                                var v107 float64 = float64((((float32(((v5 + (((float64(3.14159) * 3.14159) * (v5 + v5)) / ((v5 + ((float64(1.5) - ((v5 - (3.14159 - 3.14159)) + v5)) / ((v5 * v5) / v5))) / (v5 * 0.1)))) - v5)) * 1.5) + 3.14159) / float32((v5 - float64(-2.25)))))
                                                                if true {
                                                                    var v108 int64 = 1
                                                                    _ = v108
                                                                }
                                                                v3[""] = 2
                                                                _ = v106
                                                                _ = v107
                                                            }
                                                            if len(v10) > 0 {
                                                                v10[0] = (-(fi77))
                                                            }
                                                            v6 = v1
                                                        }
                                                        _ = v99
                                                    } else {
                                                        v109 := GenS6{F0: -79}
                                                        var v110 uint32 = 979
                                                        _ = v109
                                                        _ = v110
                                                    }
                                                    v111 := v2[0]
                                                    v112 := GenS5{F0: 2}
                                                    _ = v111
                                                    _ = v112
                                                }
                                                v113 := GenS3{F0: false, F1: 458, F2: 100}
                                                switch fi77 {
                                                case -9223372036854775808:
                                                    v1 -= v6
                                                case 695:
                                                    if true {
                                                        var v114 int8 = -85
                                                        _ = v114
                                                    } else {
                                                        fmt.Println(v4, fi77, v83)
                                                    }
                                                }
                                                _ = v113
                                                _ = v83
                                                _ = v84
                                            case -223:
                                                v1 += v1
                                            default:
                                                v115 := GenS3{F0: false, F1: 4294967295, F2: 0}
                                                _ = v115
                                            }
                                        case 32767:
                                            var v116 uint64 = 0
                                            _ = v116
                                        default:
                                            var v117 int64 = v4
                                            _ = v117
                                        }
                                        v118 := GenS1{F0: 334, F1: -70}
                                        _ = v118
                                    }
                                    var v119 string = "hello"
                                    _ = v119
                                    _ = v80
                                }
                                _ = fi78
                            }
                            var v120 int16 = 77
                            v1 = (v6 - (v1 >> ((4294967295) & 63)))
                            _ = fi77
                            _ = v120
                        }
                        v3["foo bar"] = 148
                    }
                    v5 += v5
                }
            }
            _ = v47
        }
    case 42:
        if (1e10 <= float64(0.1)) {
            v6 += v1
        }
    default:
        v121 := GenS7{F0: 936, F1: ""}
        _ = v121
    }
    v5 += v5
    v4 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
