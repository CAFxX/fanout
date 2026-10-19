package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 bool; F1 int64; F2 int16 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 int64 }

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 4294967295
    v3 := make(map[int]uint32)
    var v4 uint8 = 0
    v5 := make(map[uint32]uint32)
    var v6 float32 = 0.0
    if (!false) {
        var v7 uint8 = (v4 | (^(v4)))
        fmt.Println(v1)
        v1 = (v1 ^ v1)
        _ = v7
    }
    v5[1] = 0
    fmt.Println(v2)
    var v8 int16 = 1
    if true {
        for fi9 := 0; fi9 < 13; fi9++ {
            v6 -= 3.14159
            var v10 uint32 = 0
            var v11 uint8 = (^(v4))
            _ = fi9
            _ = v10
            _ = v11
        }
        switch "a" {
        case "":
            v4 -= 127
        case "a":
            v12 := GenS1{F0: "a", F1: 775}
            _ = v12
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                v6 = (float32(1e10) / v6)
                fmt.Println(v4, fi13, fi13)
                _ = fi13
            }
        }
    }
    fmt.Println(v1)
    v14 := GenS1{F0: (("hello" + ("foo bar" + "")) + ""), F1: 42}
    var v15 uint8 = v4
    for fi16 := 0; fi16 < 9; fi16++ {
        v4 = (v15 * v4)
        for fi17 := 0; fi17 < 1; fi17++ {
            v15 = (v15 / ((v4) | 1))
            var v18 bool = (!(-2 < 53))
            v2 -= (v2 - (v2 % ((v2) | 1)))
            v19 := GenS2{F0: 80, F1: -116}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi16
    }
    switch "" {
    case "":
        if false {
            if true {
                switch ("hello" + ("\n\t" + "\n\t")) {
                case "":
                    var v20 uint8 = (v15 + v4)
                    _ = v20
                case "foo bar":
                    for fi21 := 0; fi21 < 14; fi21++ {
                        v8 = v8
                        switch fi21 {
                        case 127:
                            v22 := GenS2{F0: 740, F1: -1}
                            _ = v22
                        case 376:
                            v23 := GenS2{F0: 2, F1: 0}
                            _ = v23
                        case 331:
                            v24 := []int{fi21, (^(fi21)), fi21, fi21}
                            _ = v24
                        default:
                            v25 := GenS2{F0: 928, F1: 127}
                            _ = v25
                        }
                        var v26 float64 = 1.5
                        _ = fi21
                        _ = v26
                    }
                }
                v15 = v4
            } else {
                for fi27 := 0; fi27 < 6; fi27++ {
                    fmt.Println(fi27, v8)
                    var v28 int32 = v1
                    for fi29 := 0; fi29 < 6; fi29++ {
                        v30 := v3[fi29]
                        v3[fi29] = 1
                        _ = fi29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
                fmt.Println(v15, v15, v4)
            }
            v31 := GenS2{F0: uint32(v2), F1: 127}
            if true {
                v2 = v2
                var v32 uint = 1
                for fi33 := 0; fi33 < 1; fi33++ {
                    v34 := GenS1{F0: ("foo bar" + ""), F1: 0}
                    _ = fi33
                    _ = v34
                }
                _ = v32
            }
            _ = v31
        } else {
            for fi35 := 0; fi35 < 17; fi35++ {
                fmt.Println(v8)
                var v36 uint = 255
                _ = fi35
                _ = v36
            }
            switch 55 {
            case 1:
                var v37 bool = (65535 <= 755)
                _ = v37
            case 42:
                for fi38 := 0; fi38 < 20; fi38++ {
                    v6 = float32((0.0 + 1e10))
                    v8 = (^(v8))
                    v39 := []int{-574}
                    _ = fi38
                    _ = v39
                }
            }
            switch ("foo bar" + "hello") {
            case "hello":
                if false {
                    v6 += ((v6 / v6) / ((float32(((1e10 / -2.25) + ((3.14159 - 1.5) / (1e10 - (3.14159 - 0.1))))) / v6) + (v6 + (v6 / (v6 + 0.0)))))
                    for fi40 := 0; fi40 < 15; fi40++ {
                        fmt.Println(v15)
                        var v41 int = (-(fi40))
                        v42 := GenS3{F0: v2}
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                    v4 = (v15 | 137)
                } else {
                    v43 := v3[128]
                    v6 = (((((v6 / (v6 - (v6 + v6))) + v6) / (v6 * (v6 + 1e10))) - v6) - v6)
                    v5[2] = v43
                    _ = v43
                }
            case "foo bar":
                var v44 int64 = -594
                _ = v44
            default:
                v45 := GenS1{F0: "hello", F1: 0}
                _ = v45
            }
        }
    case "a":
        for fi46 := 0; fi46 < 4; fi46++ {
            for fi47 := 0; fi47 < 7; fi47++ {
                v2 = v2
                if (v1 > -2147483648) {
                    v48 := []int{(fi47 * 1), fi47, (fi47 << ((0) & 63))}
                    _ = v48
                } else {
                    fmt.Println(v6)
                    v49 := GenS4{F0: -585}
                    var v50 float32 = float32(float64(v6))
                    _ = v49
                    _ = v50
                }
                var v51 int32 = (v1 / (((^(v1))) | 1))
                fmt.Println(fi47, v4)
                v6 -= v6
                _ = fi47
                _ = v51
            }
            v4 = v4
            v15 -= 0
            _ = fi46
        }
    case "foo bar":
        if false {
            var v52 int = -895
            _ = v52
        } else {
            if (!false) {
                var v53 uint64 = 2
                v6 += 0.1
                for fi54 := 0; fi54 < 8; fi54++ {
                    for fi55 := 0; fi55 < 14; fi55++ {
                        switch "hello" {
                        case "foo bar":
                            v56 := GenS5{F0: 1}
                            _ = v56
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                switch (fi55 % ((fi55) | 1)) {
                                case 1:
                                    var v57 uint16 = 1
                                    _ = v57
                                case 4294967295:
                                    v58 := GenS5{F0: 32767}
                                    _ = v58
                                case 911:
                                    v2 = (v53 % ((v2) | 1))
                                }
                            default:
                                var v59 int = 32767
                                _ = v59
                            }
                        case "":
                            switch (fi55 ^ 127) {
                            case 9223372036854775807:
                                v60 := GenS6{F0: (v8 / ((0) | 1))}
                                _ = v60
                            case -493:
                                switch (("foo bar" + "a") + (("a" + "a") + "a")) {
                                case "\n\t":
                                    var v61 uint16 = 42
                                    _ = v61
                                }
                            case 127:
                                switch (-(fi54)) {
                                case 827:
                                    var v62 float64 = 0.0
                                    _ = v62
                                case 284:
                                    v63 := GenS7{F0: (!(127 > 42)), F1: 1, F2: (v8 | (v8 << ((611) & 63)))}
                                    _ = v63
                                default:
                                    var v64 uint = 1
                                    _ = v64
                                }
                                var v65 float64 = 0.1
                                if ((0 <= 0) || (310 < 50)) {
                                    v3[(fi55 >> ((0) & 63))] = 2
                                    if (v8 <= (v8 | (v8 - -1))) {
                                        var v66 bool = true
                                        v67 := []int{fi54, (fi54 + (fi54 % ((fi55) | 1))), (fi55 << ((4294967295) & 63))}
                                        v4 = (v4 - v15)
                                        v15 += (v4 ^ (v4 % ((v4) | 1)))
                                        _ = v66
                                        _ = v67
                                    } else {
                                        var v68 bool = false
                                        v4 += (v15 & v15)
                                        if v68 {
                                            if (v8 >= v8) {
                                                v69 := GenS5{F0: 65535}
                                                if (0.1 <= (v6 / 1.0)) {
                                                    var v70 int8 = -109
                                                    var v71 string = "foo bar"
                                                    _ = v70
                                                    _ = v71
                                                }
                                                fmt.Println(v15, v65)
                                                _ = v69
                                            } else {
                                                delete(v3, fi55)
                                                var v72 uint8 = v15
                                                if (((v65 - v65) >= ((3.14159 / 3.14159) - ((((1e10 / v65) * (-2.25 + ((v65 * v65) + v65))) / float64(float32(v65))) * v65))) || v68) {
                                                    var v73 string = ""
                                                    fmt.Println(v2)
                                                    var v74 float32 = v6
                                                    var v75 uint16 = 0
                                                    v76 := v3[(fi54 | (fi54 | 32768))]
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    if (698 > fi54) {
                                                        if true {
                                                            v5[2147483648] = uint32(v53)
                                                            var v77 uint64 = 867
                                                            v6 -= float32((0.1 / (v65 * v65)))
                                                            _ = v77
                                                        }
                                                        v78 := GenS3{F0: (^(v2))}
                                                        if (false || v68) {
                                                            v53 = (v2 ^ 18446744073709551615)
                                                            fmt.Println(v2)
                                                        } else {
                                                            fmt.Println(fi54, v72)
                                                            v79 := GenS8{F0: v2, F1: v8}
                                                            _ = v79
                                                        }
                                                        _ = v78
                                                    } else {
                                                        v1 += v1
                                                    }
                                                }
                                                var v80 int8 = -1
                                                _ = v72
                                                _ = v80
                                            }
                                        }
                                        switch (fi55 ^ (fi54 | fi55)) {
                                        case 733:
                                            v81 := []int{(fi55 << ((155) & 63))}
                                            v15 = (v15 - v15)
                                            _ = v81
                                        case 4294967295:
                                            v82 := GenS3{F0: 2}
                                            _ = v82
                                        case -798:
                                            fmt.Println(v53)
                                        }
                                        _ = v68
                                    }
                                } else {
                                    v83 := GenS8{F0: (^(v53)), F1: (v8 + -716)}
                                    v84 := GenS6{F0: -1}
                                    _ = v83
                                    _ = v84
                                }
                                _ = v65
                            default:
                                if false {
                                    v1 = v1
                                    v85 := []int{(^(fi55))}
                                    fmt.Println(v1, fi54, v15)
                                    _ = v85
                                }
                            }
                            switch "hello" {
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    v4 -= v4
                                default:
                                    var v86 uint16 = 42
                                    _ = v86
                                }
                            case "a":
                                var v87 uint64 = 620
                                _ = v87
                            case "foo bar":
                                var v88 int = (fi54 ^ fi54)
                                _ = v88
                            }
                            var v89 int64 = -9223372036854775808
                            _ = v89
                        }
                        if true {
                            v90 := []int{fi54, (fi54 >> ((513) & 63)), fi54, (fi54 << ((652) & 63))}
                            fmt.Println(v15, v15, fi55)
                            var v91 uint16 = 255
                            v92 := GenS2{F0: 65535, F1: -40}
                            _ = v90
                            _ = v91
                            _ = v92
                        }
                        v93 := GenS2{F0: 347, F1: -128}
                        var v94 float32 = v6
                        v15 -= (v4 / ((126) | 1))
                        _ = fi55
                        _ = v93
                        _ = v94
                    }
                    _ = fi54
                }
                _ = v53
            }
            if true {
                if (float64(((((3.14159 - (((v6 / v6) * v6) - (-2.25 + ((0.0 + v6) / -2.25)))) * 1.5) * v6) * ((3.14159 - ((3.14159 + (v6 - (3.14159 + float32(((1e10 - -2.25) + 3.14159))))) * v6)) + (0.1 - v6)))) < (0.0 * 1.5)) {
                    switch 4294967295 {
                    case 916:
                        v15 = (v4 % (((v4 % ((v15) | 1))) | 1))
                    case 0:
                        var v95 float32 = (float32((0.1 + 1.5)) + v6)
                        _ = v95
                    case 128:
                        if false {
                            var v96 int64 = 0
                            v97 := GenS7{F0: (-91 < 42), F1: v96, F2: 1000}
                            _ = v96
                            _ = v97
                        } else {
                            switch -114 {
                            case -632:
                                if true {
                                    if ((^(v15)) != (v4 - v15)) {
                                        switch (("a" + "hello") + "a") {
                                        case "a":
                                            if (255 >= 0) {
                                                fmt.Println(v15, v8)
                                                switch -2147483648 {
                                                case 128:
                                                    if (42 > 255) {
                                                        v2 += v2
                                                    } else {
                                                        fmt.Println(v4, v15)
                                                        v98 := v5[389]
                                                        _ = len(v5)
                                                        _ = v98
                                                    }
                                                case -9223372036854775808:
                                                    v8 = int16(v2)
                                                default:
                                                    fmt.Println(v4)
                                                }
                                                v2 = (v2 << ((616) & 63))
                                            }
                                        case "\n\t":
                                            _ = len(v5)
                                        }
                                        fmt.Println(v6, v15, v1)
                                    }
                                } else {
                                    v99 := GenS2{F0: 2147483647, F1: 1}
                                    for fi100 := 0; fi100 < 18; fi100++ {
                                        var v101 int8 = 127
                                        v4 = (v15 | (v15 & (v15 / ((v4) | 1))))
                                        _ = fi100
                                        _ = v101
                                    }
                                    if false {
                                        switch -2147483648 {
                                        case 128:
                                            if ((false && true) && false) {
                                                var v102 uint8 = (v15 >> ((2) & 63))
                                                var v103 bool = true
                                                for fi104 := 0; fi104 < 6; fi104++ {
                                                    for fi105 := 0; fi105 < 14; fi105++ {
                                                        var v106 int32 = (v1 - v1)
                                                        fmt.Println(v1)
                                                        _ = fi105
                                                        _ = v106
                                                    }
                                                    v5[744] = 2147483648
                                                    _ = fi104
                                                }
                                                _ = v102
                                                _ = v103
                                            }
                                        case 1:
                                            v1 = v1
                                        }
                                    } else {
                                        if true {
                                            v1 = (v1 | (v1 - 660))
                                            v107 := []int{-2147483648, 2147483647}
                                            _ = v107
                                        } else {
                                            var v108 string = (("hello" + "foo bar") + "a")
                                            v109 := GenS2{F0: uint32(v4), F1: 18}
                                            var v110 uint = 902
                                            _ = v108
                                            _ = v109
                                            _ = v110
                                        }
                                    }
                                    _ = v99
                                }
                            case 2147483647:
                                var v111 string = "hello"
                                _ = v111
                            case 128:
                                switch 2147483647 {
                                case 42:
                                    var v112 string = ("foo bar" + "foo bar")
                                    _ = v112
                                case -1:
                                    v6 += ((3.14159 - (0.1 - (0.0 + 3.14159))) - v6)
                                }
                            default:
                                fmt.Println(v2)
                            }
                        }
                    default:
                        var v113 uint16 = 985
                        _ = v113
                    }
                    fmt.Println(v8, v2, v6)
                    v114 := GenS7{F0: (18446744073709551615 >= 401), F1: 1, F2: (v8 >> ((144) & 63))}
                    _ = v114
                } else {
                    v115 := GenS9{F0: (v8 << ((255) & 63)), F1: 870}
                    _ = v115
                }
            } else {
                v1 += 2147483647
                v116 := GenS10{F0: -1, F1: 42}
                v4 = 255
                _ = v116
            }
            for fi117 := 0; fi117 < 15; fi117++ {
                if (false || true) {
                    fmt.Println(fi117, v4, fi117)
                    v118 := GenS2{F0: 2147483648, F1: 1}
                    for fi119 := 0; fi119 < 2; fi119++ {
                        delete(v3, (fi119 - fi119))
                        v6 -= float32(float64(v6))
                        _ = fi119
                    }
                    _ = v118
                }
                delete(v5, 42)
                for fi120 := 0; fi120 < 16; fi120++ {
                    v15 += (v15 ^ 255)
                    v4 = 0
                    v121 := []int{(fi117 >> ((0) & 63)), 9223372036854775807}
                    _ = fi120
                    _ = v121
                }
                _ = fi117
            }
        }
    }
    v122 := GenS11{F0: 0}
    switch 9223372036854775807 {
    case 2147483647:
        var v123 bool = (4294967295 > -2147483648)
        _ = v123
    case 691:
        var v124 uint16 = 1
        _ = v124
    default:
        var v125 uint8 = 1
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v122
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
