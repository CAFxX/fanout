package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int8 }
func genFn0(p0 uint16, p1 int8, p2 bool) uint16 {
    if false {
        switch ("hello" + ("hello" + "a")) {
        case "\n\t":
            var v5 uint32 = 1
            _ = v5
        case "hello":
            v6 := []int{32768}
            _ = v6
        }
    } else {
        v7 := []int{-93, -369, -328}
        if ("\n\t" >= ((("\n\t" + "a") + "a") + "\n\t")) {
            if (!((2147483647 > -2147483648) && ((-16 == p1) && p2))) {
                v8 := GenS1{F0: 42, F1: (p1 * p1)}
                for fi9 := 0; fi9 < 5; fi9++ {
                    var v10 int8 = 42
                    fmt.Println(fi9)
                    _ = fi9
                    _ = v10
                }
                _ = v8
            } else {
                var v11 float32 = float32((3.14159 / 1.5))
                var v12 uint16 = p0
                p0 -= v12
                _ = v11
                _ = v12
            }
            p2 = p2
        } else {
            var v13 int32 = 0
            if (p2 || (-1 > 417)) {
                fmt.Println(p2)
                var v14 uint = 0
                if p2 {
                    var v15 int = 2147483647
                    p0 = (p0 * (p0 >> ((0) & 63)))
                    _ = v15
                } else {
                    v16 := GenS1{F0: 225, F1: (p1 % ((p1) | 1))}
                    v17 := GenS1{F0: (v14 % (((v14 * v14)) | 1)), F1: p1}
                    for fi18 := 0; fi18 < 3; fi18++ {
                        switch ("hello" + "hello") {
                        case "a":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case "":
                            for fi19 := 0; fi19 < 19; fi19++ {
                                p2 = (42 != 2)
                                var v20 int8 = p1
                                _ = fi19
                                _ = v20
                            }
                        case "\n\t":
                            switch "a" {
                            case "a":
                                if (p2 || ((^(v13)) < (v13 >> (((^(v14))) & 63)))) {
                                    v14 = v14
                                    fmt.Println(p2)
                                    var v21 int16 = -475
                                    _ = v21
                                }
                            case "hello":
                                for fi22 := 0; fi22 < 11; fi22++ {
                                    switch 1 {
                                    case -522:
                                        if (v14 >= v14) {
                                            v7 = append(v7, (fi22 % (((fi22 >> ((2) & 63))) | 1)))
                                            fmt.Println(fi18, p2)
                                            var v23 uint64 = 18446744073709551615
                                            _ = v23
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                p0 = p0
                                                var v24 float64 = ((((1e10 / 1.0) * ((1.5 - (0.1 - ((float64(0.0) - 0.0) + (1e10 + 0.1)))) * (0.0 + (float64(1.5) / 0.1)))) - 1e10) * float64(-2.25))
                                                _ = v24
                                            case "\n\t":
                                                fmt.Println(v14)
                                            }
                                            var v25 uint16 = (p0 / (((p0 % ((p0) | 1))) | 1))
                                            v13 = v13
                                            _ = v25
                                        }
                                    case 32768:
                                        fmt.Println(v14)
                                    case -1:
                                        fmt.Println(p1)
                                        if ("a" <= "hello") {
                                            v14 += (v14 | v14)
                                            if p2 {
                                                var v26 int8 = p1
                                                _ = v26
                                            }
                                            var v27 int32 = (v13 >> ((606) & 63))
                                            _ = v27
                                        } else {
                                            v13 = (v13 / ((-527) | 1))
                                        }
                                        var v28 uint16 = (p0 - (p0 & 796))
                                        _ = v28
                                    }
                                    _ = fi22
                                }
                            default:
                                fmt.Println(p1)
                            }
                        default:
                            p2 = ((!true) && ((1.5 * -2.25) >= -2.25))
                        }
                        var v29 float32 = (1.5 * (3.14159 * (0.0 * 0.1)))
                        _ = fi18
                        _ = v29
                    }
                    _ = v16
                    _ = v17
                }
                _ = v14
            }
            _ = v13
        }
        _ = v7
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 5; fi31++ {
            p1 = (p1 & 72)
            _ = fi31
        }
        if (((p2 || (-1 != 1)) && ((1 <= 2) && (int64(fi30) == 1))) || (!p2)) {
            v32 := GenS1{F0: 42, F1: int8(p0)}
            _ = v32
        }
        v33 := []int{(fi30 << ((0) & 63)), fi30, fi30, 4294967295}
        _ = fi30
        _ = v33
    }
    if p2 {
        v34 := GenS2{F0: "hello", F1: 42}
        _ = v34
    }
    v35 := GenS2{F0: "foo bar", F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    return (^(p0))
}
func genFn1() bool {
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 uint16 = 2
        _ = fi36
        _ = v37
    }
    switch ("\n\t" + ("\n\t" + "")) {
    case "\n\t":
        v38 := GenS1{F0: 4294967295, F1: 42}
        _ = v38
    case "a":
        var v39 uint32 = 2147483648
        _ = v39
    case "foo bar":
        var v40 int = -9223372036854775808
        _ = v40
    default:
        var v41 bool = false
        _ = v41
    }
    v42 := []int{0}
    if len(v42) > 0 {
        v42[0] = -1
    }
    _ = v42
    return (4294967295 < 42)
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (200 <= 754)
    var v3 uint32 = 2147483648
    var v4 uint = 18446744073709551615
    switch ("\n\t" + "a") {
    case "":
        var v43 int = 42
        _ = v43
    case "hello":
        v44 := GenS2{F0: ((("a" + "a") + ("" + "a")) + "a"), F1: -386}
        _ = v44
    default:
        switch "\n\t" {
        case "hello":
            v45 := GenS1{F0: (v4 << (((^(v4))) & 63)), F1: -1}
            _ = v45
        case "a":
            if v2 {
                var v46 int8 = -128
                fmt.Println(v2, v2, v4)
                var v47 int64 = -1
                _ = v46
                _ = v47
            }
        default:
            for fi48 := 0; fi48 < 3; fi48++ {
                delete(v1, 32767)
                var v49 uint16 = 42
                _ = fi48
                _ = v49
            }
        }
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        for fi51 := 0; fi51 < 20; fi51++ {
            v52 := []int{}
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    switch -764 {
    case -417:
        if ((("a" + "a") + "a") == "\n\t") {
            if false {
                v53 := []int{}
                for fi54 := 0; fi54 < 10; fi54++ {
                    v55 := GenS3{F0: "", F1: 0}
                    for fi56 := 0; fi56 < 3; fi56++ {
                        v53 = append(v53, -361)
                        v57 := GenS1{F0: (^(v4)), F1: -128}
                        var v58 uint16 = 196
                        var v59 uint16 = v58
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v53
            } else {
                v3 = (v3 ^ v3)
            }
            var v60 int64 = 9223372036854775807
            _ = v60
        } else {
            if (uint32(v4) < 1) {
                switch ((((("a" + "foo bar") + "a") + ("" + "a")) + ("a" + "")) + "") {
                case "":
                    var v61 int8 = 42
                    _ = v61
                case "\n\t":
                    fmt.Println(v2, v4, v3)
                default:
                    if (403 > 476) {
                        if v2 {
                            switch ((("\n\t" + "") + "a") + "") {
                            case "hello":
                                for fi62 := 0; fi62 < 8; fi62++ {
                                    switch 613 {
                                    case 1:
                                        var v63 uint = 255
                                        _ = v63
                                    case 127:
                                        switch (fi62 | (^(fi62))) {
                                        case 2147483647:
                                            v64 := GenS3{F0: "foo bar", F1: 130}
                                            _ = v64
                                        case -9223372036854775808:
                                            fmt.Println(v2)
                                        case -650:
                                            var v65 int = (fi62 ^ fi62)
                                            _ = v65
                                        default:
                                            v3 += (v3 - (^(v3)))
                                        }
                                    case 128:
                                        for fi66 := 0; fi66 < 8; fi66++ {
                                            v67 := []int{(fi66 << ((v4) & 63))}
                                            _ = len(v1)
                                            _ = fi66
                                            _ = v67
                                        }
                                    default:
                                        switch -1 {
                                        case 987:
                                            for fi68 := 0; fi68 < 13; fi68++ {
                                                v69 := []int{fi62, 32768, (fi62 - (fi68 * -236))}
                                                switch "" {
                                                case "hello":
                                                    var v70 float64 = 0.0
                                                    _ = v70
                                                }
                                                v1[(^(fi68))] = 1
                                                var v71 uint64 = 42
                                                switch "a" {
                                                case "\n\t":
                                                    switch -9223372036854775808 {
                                                    case -330:
                                                        switch "hello" {
                                                        case "":
                                                            var v72 bool = (1e10 < (0.0 + 0.1))
                                                            _ = v72
                                                        default:
                                                            var v73 int8 = 100
                                                            _ = v73
                                                        }
                                                    }
                                                case "foo bar":
                                                    v4 = uint(v71)
                                                }
                                                _ = fi68
                                                _ = v69
                                                _ = v71
                                            }
                                        case 32767:
                                            delete(v1, 32767)
                                        default:
                                            _ = len(v1)
                                        }
                                    }
                                    v3 = v3
                                    _ = fi62
                                }
                            case "foo bar":
                                for fi74 := 0; fi74 < 16; fi74++ {
                                    var v75 bool = (uint64(v3) == 292)
                                    _ = fi74
                                    _ = v75
                                }
                            }
                            v2 = (444 >= 525)
                        }
                        v3 += 42
                        v76 := []int{32767, 128}
                        _ = v76
                    } else {
                        switch 128 {
                        case -281:
                            v2 = v2
                        default:
                            v4 = (v4 ^ v4)
                        }
                    }
                }
                var v77 uint = (v4 + (^(v4)))
                v77 += v4
                _ = v77
            } else {
                for fi78 := 0; fi78 < 16; fi78++ {
                    v79 := GenS4{F0: 42, F1: -2147483648}
                    var v80 uint64 = 1
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                v81 := GenS5{F0: 279, F1: 818, F2: 0}
                _ = v81
            }
        }
    }
    v2 = (-2.25 < 0.1)
    var v82 float64 = 0.0
    for fi83 := 0; fi83 < 14; fi83++ {
        if (1 >= 42) {
            v84 := GenS4{F0: 105, F1: (fi83 << ((v4) & 63))}
            for fi85 := 0; fi85 < 14; fi85++ {
                _ = len(v1)
                if (v82 != ((v82 * 1.5) + v82)) {
                    v2 = (1e10 != float32(v82))
                    var v86 uint64 = 18446744073709551615
                    v87 := GenS6{F0: 0, F1: 42}
                    _ = v86
                    _ = v87
                }
                v88 := GenS6{F0: -2147483648, F1: -886}
                _ = fi85
                _ = v88
            }
            _ = v84
        } else {
            if v2 {
                if true {
                    fmt.Println(v2, v82, v3)
                    var v89 bool = (348 >= -536)
                    if (18446744073709551615 == 799) {
                        v82 += 1.5
                    }
                    _ = v89
                } else {
                    if v2 {
                        var v90 string = ("foo bar" + "")
                        v91 := GenS4{F0: -128, F1: (fi83 / (((fi83 >> (((v4 - 73)) & 63))) | 1))}
                        var v92 bool = ((1 == 0) || (!((v3 != 2) && (1.5 <= (-2.25 + (float32(v82) - 3.14159))))))
                        _ = v90
                        _ = v91
                        _ = v92
                    }
                    for fi93 := 0; fi93 < 20; fi93++ {
                        if v2 {
                            v4 += v4
                            var v94 bool = (v2 || v2)
                            switch (fi93 & 127) {
                            case -1:
                                var v95 int8 = 100
                                _ = v95
                            case -199:
                                v94 = (!false)
                                switch fi93 {
                                case 509:
                                    var v96 int8 = 127
                                    _ = v96
                                }
                                var v97 uint32 = v3
                                _ = v97
                            case 0:
                                var v98 int64 = -9223372036854775808
                                _ = v98
                            default:
                                if (508 >= genFn0(0, 100, (42 < 42))) {
                                    fmt.Println(v3, v4)
                                    switch "" {
                                    case "hello":
                                        var v99 string = ("a" + ((((("hello" + "a") + ("\n\t" + ("foo bar" + "foo bar"))) + "foo bar") + "foo bar") + "hello"))
                                        _ = v99
                                    case "\n\t":
                                        var v100 int32 = 14
                                        _ = v100
                                    case "a":
                                        v4 += (v4 + 698)
                                    default:
                                        if v2 {
                                            delete(v1, 32768)
                                            var v101 int64 = 42
                                            _ = v101
                                        } else {
                                            if v2 {
                                                v3 += (v3 ^ 859)
                                                switch ("foo bar" + ("a" + "")) {
                                                case "foo bar":
                                                    v102 := GenS7{F0: (v4 % ((2) | 1)), F1: 1}
                                                    _ = v102
                                                case "hello":
                                                    if false {
                                                        v103 := v1[fi83]
                                                        var v104 int32 = 1
                                                        v103 += int(v104)
                                                        _ = v103
                                                        _ = v104
                                                    }
                                                default:
                                                    v105 := GenS8{F0: v3, F1: 42}
                                                    _ = v105
                                                }
                                                delete(v1, int(v3))
                                            }
                                            if ((-9223372036854775808 == (fi93 << ((718) & 63))) && false) {
                                                var v106 int64 = 714
                                                if v94 {
                                                    v107 := GenS5{F0: uint16(v4), F1: -1, F2: 0}
                                                    _ = v107
                                                } else {
                                                    v94 = true
                                                    v108 := GenS2{F0: "", F1: 83}
                                                    _ = v108
                                                }
                                                _ = v106
                                            } else {
                                                var v109 uint8 = 186
                                                v2 = v94
                                                _ = v109
                                            }
                                            v94 = ((691 > uint64(v3)) && (2147483647 >= 0))
                                        }
                                    }
                                    var v110 uint16 = 2
                                    _ = v110
                                }
                            }
                            _ = v94
                        } else {
                            if v2 {
                                v111 := []int{fi93, fi83, (fi93 ^ fi93), 4294967295}
                                _ = v111
                            }
                            switch fi83 {
                            case 538:
                                var v112 bool = false
                                _ = v112
                            case -9223372036854775808:
                                var v113 float64 = (0.0 * (((v82 + v82) + v82) * (0.1 / 1.0)))
                                _ = v113
                            default:
                                v2 = ((v3 | v3) < 50)
                                var v114 int16 = 0
                                if v2 {
                                    v115 := v1[-518]
                                    _ = v115
                                }
                                v3 = 4294967295
                                v116 := GenS4{F0: 42, F1: (^(fi93))}
                                v117 := []int{(fi93 * fi83), fi93}
                                _ = v114
                                _ = v116
                                _ = v117
                            }
                        }
                        v2 = v2
                        _ = fi93
                    }
                }
            }
            switch (("\n\t" + ("hello" + "hello")) + "\n\t") {
            case "":
                delete(v1, fi83)
            case "a":
                v1[(fi83 / ((fi83) | 1))] = (fi83 % (((fi83 & fi83)) | 1))
            case "foo bar":
                switch -1 {
                case 9223372036854775807:
                    for fi118 := 0; fi118 < 6; fi118++ {
                        var v119 int16 = 972
                        v4 = v4
                        v120 := []int{}
                        fmt.Println(fi118)
                        _ = fi118
                        _ = v119
                        _ = v120
                    }
                case 32767:
                    v121 := v1[fi83]
                    _ = v121
                case -9223372036854775808:
                    v122 := []int{(fi83 << ((v4) & 63))}
                    _ = v122
                default:
                    if false {
                        var v123 float32 = (0.0 / 1.0)
                        for fi124 := 0; fi124 < 17; fi124++ {
                            if (!true) {
                                fmt.Println(v4, fi83, fi124)
                                v3 = (v3 ^ (^(v3)))
                            } else {
                                v125 := []int{(fi83 - 127), 886}
                                v126 := GenS9{F0: -1}
                                var v127 int32 = 42
                                v2 = ((^(v127)) != (v127 + (v127 << ((255) & 63))))
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                            _ = fi124
                        }
                        _ = v123
                    } else {
                        for fi128 := 0; fi128 < 2; fi128++ {
                            switch 853 {
                            case 9223372036854775807:
                                v129 := GenS10{F0: 100}
                                _ = v129
                            case 128:
                                v130 := GenS4{F0: 120, F1: fi83}
                                _ = v130
                            }
                            _ = fi128
                        }
                        v4 -= v4
                        v131 := []int{}
                        _ = v131
                    }
                }
            default:
                delete(v1, 1)
            }
        }
        var v132 float32 = 0.0
        _ = fi83
        _ = v132
    }
    for fi133 := 0; fi133 < 9; fi133++ {
        fmt.Println(v2, v82)
        _ = fi133
    }
    v1[32767] = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
}
