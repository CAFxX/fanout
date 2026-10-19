package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int16; F1 int32; F2 int }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32 }
func genFn0(p0 int) uint32 {
    switch "a" {
    case "":
        p0 = (-(p0))
    default:
        if true {
            for fi4 := 0; fi4 < 9; fi4++ {
                p0 -= (fi4 * fi4)
                _ = fi4
            }
            v5 := GenS1{F0: 281}
            var v6 float32 = 0.0
            _ = v5
            _ = v6
        } else {
            p0 = (p0 << ((142) & 63))
        }
    }
    if (!true) {
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 2; fi8++ {
                p0 = (p0 - fi7)
                _ = fi8
            }
            _ = fi7
        }
        var v9 string = ""
        _ = v9
    }
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 13; fi11++ {
        p0 = fi11
        switch 128 {
        case 9223372036854775807:
            for fi12 := 0; fi12 < 9; fi12++ {
                p0 -= -1
                p0 = (fi11 - p0)
                v13 := []int{(^(fi11)), -1, 32768}
                _ = fi12
                _ = v13
            }
        case -482:
            var v14 int = (fi11 % (((fi11 >> ((42) & 63))) | 1))
            _ = v14
        case 32767:
            if true {
                var v15 int8 = 0
                _ = v15
            }
        }
        switch "foo bar" {
        case "":
            v16 := GenS1{F0: -76}
            _ = v16
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]bool)
    v3 := []uint32{2147483648, 2147483647}
    if true {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        var v17 float32 = float32(float64(3.14159))
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v17)
                _ = fi19
            }
            if true {
                v17 += (v17 + v17)
                v17 -= v17
                var v20 int32 = 42
                _ = v20
            } else {
                var v21 int8 = -1
                if ((false || true) || true) {
                    _ = len(v3)
                    v17 = 1.5
                } else {
                    for fi22 := 0; fi22 < 14; fi22++ {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v1)
                        case "foo bar":
                            if (!((!true) || false)) {
                                _ = len(v3)
                                var v23 bool = (1 != 109)
                                var v24 int = fi18
                                var v25 bool = v23
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                        default:
                            if true {
                                var v26 uint32 = 985
                                _ = v26
                            }
                        }
                        var v27 float32 = (0.0 + v17)
                        var v28 int16 = 1000
                        if (!false) {
                            if false {
                                fmt.Println(v21, v21, v1)
                            } else {
                                if false {
                                    var v29 uint32 = 1
                                    switch (fi18 + fi22) {
                                    case 32768:
                                        v27 = 3.14159
                                    }
                                    var v30 uint32 = 42
                                    v29 += genFn0((fi22 >> ((255) & 63)))
                                    var v31 int8 = -91
                                    if true {
                                        var v32 float32 = 0.0
                                        v33 := GenS2{F0: 67, F1: fi22}
                                        _ = v32
                                        _ = v33
                                    } else {
                                        switch (fi18 >> ((2) & 63)) {
                                        case -273:
                                            if len(v3) > 0 {
                                                v3[0] = 119
                                            }
                                        case 32768:
                                            if ((!((v30 | v29) == 0)) || (!((v29 * v29) < v29))) {
                                                if true {
                                                    v30 = (v30 % ((genFn0((fi18 % ((133) | 1)))) | 1))
                                                } else {
                                                    if true {
                                                        var v34 int16 = 32767
                                                        v1 = v1
                                                        v17 -= v27
                                                        _ = v34
                                                    } else {
                                                        _ = len(v3)
                                                        fmt.Println(v17, v17, v17)
                                                    }
                                                }
                                                var v35 bool = (v17 < (3.14159 / (0.0 - 3.14159)))
                                                if (true && v35) {
                                                    if v35 {
                                                        delete(v2, fi18)
                                                        var v36 uint = 495
                                                        var v37 float32 = v27
                                                        if false {
                                                            v29 += v29
                                                        } else {
                                                            switch (("foo bar" + v1) + "foo bar") {
                                                            case "hello":
                                                                v38 := GenS2{F0: 85, F1: (^(fi22))}
                                                                _ = v38
                                                            case "":
                                                                v39 := GenS1{F0: int64(v31)}
                                                                _ = v39
                                                            default:
                                                                fmt.Println(v27, v36, v1)
                                                            }
                                                            var v40 int64 = -9223372036854775808
                                                            switch fi18 {
                                                            case -9223372036854775808:
                                                                v27 -= (((-2.25 + 0.0) * (v37 / v17)) / v27)
                                                            case -881:
                                                                v41 := GenS1{F0: v40}
                                                                var v42 int = -1
                                                                if len(v3) > 0 {
                                                                    v3[0] = v29
                                                                }
                                                                _ = v41
                                                                _ = v42
                                                            default:
                                                                v2[fi18] = v35
                                                            }
                                                            _ = v40
                                                        }
                                                        _ = v36
                                                        _ = v37
                                                    }
                                                    v43 := GenS1{F0: 28}
                                                    if (("a" > v1) && v35) {
                                                        var v44 string = ("foo bar" + v1)
                                                        _ = v44
                                                    }
                                                    _ = v43
                                                } else {
                                                    delete(v2, 128)
                                                }
                                                _ = v35
                                            } else {
                                                if true {
                                                    v45 := GenS1{F0: -543}
                                                    _ = v45
                                                } else {
                                                    var v46 bool = (!false)
                                                    v30 = v30
                                                    _ = v46
                                                }
                                                var v47 int = fi18
                                                _ = v47
                                            }
                                        case -532:
                                            v48 := GenS3{F0: 2, F1: 670, F2: 2}
                                            _ = v48
                                        }
                                        _ = len(v2)
                                        var v49 uint8 = 127
                                        _ = v49
                                    }
                                    v50 := GenS1{F0: 9223372036854775807}
                                    v51 := GenS2{F0: 127, F1: (fi22 * fi22)}
                                    var v52 string = v1
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    fmt.Println(v1)
                                    v1 = ("hello" + ("hello" + ""))
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                }
                            }
                            var v53 int64 = -9223372036854775808
                            var v54 int = (fi22 ^ -2147483648)
                            if (!false) {
                                var v55 uint16 = 42
                                var v56 uint64 = 0
                                _ = v55
                                _ = v56
                            } else {
                                fmt.Println(v21)
                            }
                            if true {
                                v57 := GenS3{F0: 0, F1: 18446744073709551615, F2: 4294967295}
                                _ = len(v3)
                                if (!(127 <= 1)) {
                                    v58 := GenS4{F0: true, F1: 110}
                                    _ = v58
                                } else {
                                    var v59 bool = (908 > 1)
                                    var v60 uint = 42
                                    _ = v59
                                    _ = v60
                                }
                                _ = v57
                            }
                            switch v54 {
                            case 127:
                                v54 = 236
                            case 722:
                                var v61 uint64 = 122
                                _ = v61
                            }
                            _ = v53
                            _ = v54
                        } else {
                            var v62 int = (^(fi22))
                            _ = v62
                        }
                        switch (fi22 + fi18) {
                        case 127:
                            switch "a" {
                            case "a":
                                switch (fi22 ^ -749) {
                                case -325:
                                    var v63 float64 = (0.0 + 1e10)
                                    _ = v63
                                case -1:
                                    var v64 bool = false
                                    _ = v64
                                default:
                                    switch fi22 {
                                    case 366:
                                        switch v1 {
                                        case "\n\t":
                                            if (true && false) {
                                                var v65 int8 = (v21 | v21)
                                                if true {
                                                    if false {
                                                        v66 := GenS2{F0: 2, F1: fi22}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v17 += 0.0
                                                        fmt.Println(v1)
                                                        _ = v66
                                                    } else {
                                                        v21 += v21
                                                        if false {
                                                            v27 = 0.1
                                                        }
                                                    }
                                                } else {
                                                    var v67 uint64 = 2
                                                    switch fi18 {
                                                    case -831:
                                                        if (9223372036854775807 >= -761) {
                                                            if (!false) {
                                                                fmt.Println(v28, v67)
                                                                var v68 string = (("foo bar" + v1) + ("" + (v1 + "foo bar")))
                                                                _ = v68
                                                            }
                                                            if false {
                                                                fmt.Println(v28, v65, v65)
                                                                var v69 uint64 = v67
                                                                var v70 int32 = 314
                                                                v3 = append(v3, 422)
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            fmt.Println(v65)
                                                        } else {
                                                            v2[fi18] = true
                                                        }
                                                    default:
                                                        if false {
                                                            fmt.Println(v65)
                                                            fmt.Println(v1, v67, v67)
                                                            fmt.Println(v27, v65)
                                                        }
                                                    }
                                                    if true {
                                                        delete(v2, (fi18 & fi18))
                                                        v27 -= v17
                                                    }
                                                    _ = v67
                                                }
                                                _ = v65
                                            }
                                        }
                                    case -1:
                                        if true {
                                            var v71 int8 = (^(v21))
                                            if (!((-(v28)) <= 1000)) {
                                                v72 := GenS5{F0: 18446744073709551615}
                                                if (true || (-2.25 < -2.25)) {
                                                    v71 -= v21
                                                } else {
                                                    fmt.Println(fi22, v17, v17)
                                                    v73 := GenS6{F0: 630, F1: v1}
                                                    v74 := GenS3{F0: 65535, F1: 4294967295, F2: 889}
                                                    _ = v73
                                                    _ = v74
                                                }
                                                _ = v72
                                            } else {
                                                v21 += v71
                                                v17 -= 1.5
                                                if true {
                                                    if false {
                                                        if false {
                                                            var v75 uint64 = 2
                                                            _ = v75
                                                        }
                                                    } else {
                                                        var v76 uint16 = 1
                                                        _ = v76
                                                    }
                                                } else {
                                                    switch v1 {
                                                    case "foo bar":
                                                        delete(v2, (fi18 ^ (fi22 >> ((472) & 63))))
                                                    case "hello":
                                                        fmt.Println(v17)
                                                    case "\n\t":
                                                        v21 -= v21
                                                    }
                                                }
                                            }
                                            var v77 int = fi18
                                            _ = v71
                                            _ = v77
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi22
                        _ = v27
                        _ = v28
                    }
                    var v78 uint32 = 0
                    _ = v78
                }
                switch (v1 + v1) {
                case "\n\t":
                    v1 = v1
                case "a":
                    v1 = v1
                case "":
                    if (-2147483648 <= -2147483648) {
                        v21 = (v21 + 46)
                        var v79 uint = 65535
                        _ = v79
                    }
                }
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    } else {
        for fi80 := 0; fi80 < 5; fi80++ {
            if (0.1 != -2.25) {
                var v81 bool = (false && (495 < 0))
                _ = v81
            } else {
                v2[(^(fi80))] = (-7 > -11)
            }
            v82 := GenS1{F0: 0}
            var v83 int16 = -297
            _ = fi80
            _ = v82
            _ = v83
        }
        switch v1 {
        case "\n\t":
            switch v1 {
            case "hello":
                if (9223372036854775807 > -1) {
                    var v84 uint8 = 42
                    _ = v84
                } else {
                    for fi85 := 0; fi85 < 16; fi85++ {
                        var v86 uint = 42
                        _ = fi85
                        _ = v86
                    }
                }
            case "foo bar":
                if (2 <= 703) {
                    if (!(genFn0(158) <= 2147483648)) {
                        switch 42 {
                        case 4294967295:
                            for fi87 := 0; fi87 < 18; fi87++ {
                                var v88 int64 = 1
                                _ = fi87
                                _ = v88
                            }
                        case 0:
                            delete(v2, -309)
                        }
                        switch 926 {
                        case -916:
                            v1 = v1
                        }
                        for fi89 := 0; fi89 < 11; fi89++ {
                            if false {
                                if (!((0.0 >= 3.14159) && (492 == 4294967295))) {
                                    v90 := GenS1{F0: 270}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v90
                                } else {
                                    v2[(fi89 * (fi89 * fi89))] = true
                                    switch 127 {
                                    case 4294967295:
                                        if ((fi89 % (((-(fi89))) | 1)) != fi89) {
                                            var v91 int = 128
                                            var v92 uint32 = 781
                                            _ = v91
                                            _ = v92
                                        }
                                    case 127:
                                        v2[-383] = (255 != 49)
                                    }
                                }
                                var v93 string = v1
                                _ = v93
                            } else {
                                v1 = ("a" + ((("" + (v1 + "foo bar")) + (v1 + (v1 + (v1 + (v1 + (v1 + v1)))))) + ("foo bar" + v1)))
                            }
                            _ = fi89
                        }
                    } else {
                        v1 = ""
                        for fi94 := 0; fi94 < 1; fi94++ {
                            if ((!(float32(3.14159) > 1e10)) || false) {
                                v95 := GenS7{F0: 1000, F1: 1, F2: 127}
                                _ = v95
                            } else {
                                var v96 uint16 = 143
                                _ = v96
                            }
                            var v97 uint32 = 1
                            _ = fi94
                            _ = v97
                        }
                        for fi98 := 0; fi98 < 11; fi98++ {
                            v3 = append(v3, 42)
                            v99 := GenS5{F0: 18446744073709551615}
                            _ = fi98
                            _ = v99
                        }
                    }
                } else {
                    v100 := GenS8{F0: 514, F1: 958}
                    v2[42] = (1 <= 675)
                    v1 = "hello"
                    _ = v100
                }
            default:
                for fi101 := 0; fi101 < 20; fi101++ {
                    fmt.Println(fi101)
                    _ = fi101
                }
            }
        }
    }
    if ("foo bar" <= (v1 + v1)) {
        fmt.Println(v1)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = "\n\t"
        for fi102 := 0; fi102 < 5; fi102++ {
            for fi103 := 0; fi103 < 15; fi103++ {
                v104 := GenS9{F0: -32768}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if false {
                    var v105 uint8 = 0
                    var v106 uint64 = 42
                    _ = v105
                    _ = v106
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi103
                _ = v104
            }
            var v107 int64 = 9223372036854775807
            var v108 bool = true
            _ = fi102
            _ = v107
            _ = v108
        }
    }
    v3 = append(v3, 149)
    switch -9223372036854775808 {
    case 523:
        var v109 bool = (1.5 != 1.5)
        _ = v109
    default:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    v110 := v2[-50]
    v111 := GenS9{F0: -32768}
    v112 := v2[343]
    switch v1 {
    case "\n\t":
        v1 = v1
    case "foo bar":
        for fi113 := 0; fi113 < 13; fi113++ {
            var v114 uint16 = 12
            var v115 uint64 = 667
            _ = fi113
            _ = v114
            _ = v115
        }
    case "hello":
        for fi116 := 0; fi116 < 19; fi116++ {
            v117 := GenS2{F0: 255, F1: (fi116 | -17)}
            if (1000 == -323) {
                _ = len(v3)
            } else {
                var v118 uint8 = 2
                _ = len(v2)
                var v119 int8 = 42
                _ = v118
                _ = v119
            }
            _ = fi116
            _ = v117
        }
    }
    switch 32767 {
    case 176:
        fmt.Println(v1, v112)
    case -1:
        v1 = (v1 + v1)
    case 1:
        var v120 uint32 = 2
        _ = v120
    }
    var v121 bool = (!(3.14159 < 3.14159))
    fmt.Println(v110, v112, v1)
    v121 = v112
    v3 = append(v3, 2)
    fmt.Println(v112, v110)
    var v122 int16 = 1000
    for fi123 := 0; fi123 < 7; fi123++ {
        v124 := GenS10{F0: 705}
        if len(v3) > 0 {
            v3[0] = 42
        }
        v1 = v1
        _ = fi123
        _ = v124
    }
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v112:", v112)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v110
    _ = v111
    _ = v112
    _ = v121
    _ = v122
    _ = v2
    _ = v3
}
