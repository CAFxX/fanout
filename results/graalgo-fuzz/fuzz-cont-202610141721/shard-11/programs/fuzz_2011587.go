package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 uint16; F1 int8; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 int32; F1 uint; F2 uint64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int32 }
func genFn0(p0 uint8) uint {
    fmt.Println(p0)
    v4 := []int{-2147483648, 128, 127, 32768}
    _ = len(v4)
    _ = p0
    _ = v4
    return 908
}
func genFn1(p0 uint16, p1 int32, p2 int64) int64 {
    fmt.Println(p0, p0, p2)
    v5 := GenS1{F0: (-791 < int16(p2)), F1: 65535, F2: -62}
    v6 := []int{323, 42, -9223372036854775808, 0}
    var v7 uint = 2
    switch ("" + "a") {
    case "foo bar":
        fmt.Println(p1)
    default:
        switch -9223372036854775808 {
        case 0:
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := GenS2{F0: 2, F1: false}
                var v10 uint = v7
                switch fi8 {
                case -101:
                    v6 = append(v6, fi8)
                case -1:
                    if len(v6) > 0 {
                        v6[0] = (fi8 | (fi8 / ((-75) | 1)))
                    }
                case -486:
                    _ = len(v6)
                default:
                    if true {
                        switch ("" + "a") {
                        case "foo bar":
                            _ = len(v6)
                        case "a":
                            var v11 int16 = -1
                            _ = v11
                        case "":
                            switch (fi8 % ((fi8) | 1)) {
                            case 136:
                                fmt.Println(fi8, p1, p1)
                            default:
                                for fi12 := 0; fi12 < 14; fi12++ {
                                    v7 = 619
                                    _ = fi12
                                }
                            }
                        default:
                            var v13 int = 1
                            _ = v13
                        }
                    }
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        case 4294967295:
            for fi14 := 0; fi14 < 16; fi14++ {
                if ((p0 / ((2) | 1)) < (p0 & p0)) {
                    var v15 int64 = (^(p2))
                    _ = v15
                } else {
                    _ = len(v6)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi16 := 0; fi16 < 6; fi16++ {
                        p1 -= p1
                        var v17 int8 = -67
                        var v18 float32 = 0.1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
                switch fi14 {
                case -9223372036854775808:
                    if (true || (1e10 >= 3.14159)) {
                        v19 := GenS3{F0: uint16(v7), F1: p0}
                        var v20 uint8 = 171
                        for fi21 := 0; fi21 < 5; fi21++ {
                            if false {
                                p0 = p0
                            } else {
                                p2 += p2
                                p2 = -9223372036854775808
                                v22 := GenS3{F0: 65535, F1: (p0 >> (((v7 * v7)) & 63))}
                                _ = v22
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if (("hello" + "foo bar") != "a") {
                                _ = len(v6)
                                p1 -= (p1 - (^(p1)))
                                v6 = append(v6, 436)
                            }
                            var v23 int = 9223372036854775807
                            var v24 int64 = -1
                            switch "a" {
                            case "\n\t":
                                if ((1e10 / -2.25) > float32(1.5)) {
                                    var v25 int64 = (v24 << ((genFn0((v20 / ((v20) | 1)))) & 63))
                                    if false {
                                        var v26 int64 = (v25 % ((v25) | 1))
                                        _ = v26
                                    }
                                    _ = v25
                                }
                            }
                            _ = fi21
                            _ = v23
                            _ = v24
                        }
                        _ = v19
                        _ = v20
                    } else {
                        for fi27 := 0; fi27 < 12; fi27++ {
                            v28 := GenS4{F0: 2, F1: 127, F2: 255}
                            _ = fi27
                            _ = v28
                        }
                        _ = len(v6)
                        var v29 int16 = 42
                        _ = v29
                    }
                case 32768:
                    v7 += v7
                default:
                    var v30 float64 = 1e10
                    _ = v30
                }
                _ = fi14
            }
        default:
            v31 := GenS5{F0: 100}
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (-(p2))
}

func main() {
    var v1 string = ("\n\t" + "a")
    v2 := make(map[string]bool)
    var v3 uint16 = 255
    if ((0 == 0) && (9223372036854775807 > -9223372036854775808)) {
        var v32 float32 = (0.1 / (((((1e10 * (-2.25 - ((((0.0 - float32((((1e10 * (0.0 - -2.25)) - 0.1) - 0.1))) + 3.14159) - (0.1 * (0.1 * 1e10))) / (float32((float64(0.1) / 1.0)) / (float32(((1.5 * (0.0 / 1.5)) * (0.0 / (0.1 / (0.1 - float64((0.1 + 1.5))))))) - -2.25))))) + 0.0) / (1e10 + -2.25)) - float32((0.1 - (0.0 * ((0.1 / 0.1) * 1.5))))) / 0.1))
        v33 := GenS5{F0: 0}
        if false {
            v3 = v3
            v34 := GenS1{F0: false, F1: 2147483648, F2: 32767}
            _ = v34
        }
        _ = v32
        _ = v33
    } else {
        fmt.Println(v3, v1)
        switch -336 {
        case -1:
            for fi35 := 0; fi35 < 19; fi35++ {
                fmt.Println(v3, v1, v1)
                _ = fi35
            }
        case 127:
            switch 42 {
            case 62:
                var v36 bool = true
                _ = v36
            case -1:
                v1 = "foo bar"
            case -79:
                for fi37 := 0; fi37 < 13; fi37++ {
                    v38 := GenS2{F0: 1, F1: (4294967295 != 0)}
                    switch v1 {
                    case "foo bar":
                        switch 208 {
                        case -524:
                            for fi39 := 0; fi39 < 7; fi39++ {
                                if false {
                                    v1 = v1
                                    if (false || (!true)) {
                                        v1 = "a"
                                        var v40 int64 = 295
                                        v41 := GenS6{F0: 42, F1: -1}
                                        if false {
                                            var v42 int = 2147483647
                                            _ = v42
                                        }
                                        _ = v40
                                        _ = v41
                                    }
                                    if false {
                                        delete(v2, v1)
                                    } else {
                                        var v43 int8 = 0
                                        if false {
                                            if true {
                                                fmt.Println(v1, fi37, v1)
                                                v44 := []int{0, fi39, fi39, (fi39 + fi37)}
                                                fmt.Println(fi37, v1, fi39)
                                                _ = v44
                                            }
                                        }
                                        v3 += (v3 & v3)
                                        var v45 uint16 = 255
                                        _ = v43
                                        _ = v45
                                    }
                                } else {
                                    v46 := GenS4{F0: 4294967295, F1: 103, F2: 248}
                                    fmt.Println(fi37)
                                    if (3.14159 != 1.5) {
                                        v47 := []int{}
                                        _ = v47
                                    } else {
                                        v1 = "hello"
                                        var v48 uint8 = 2
                                        var v49 int16 = int16(v48)
                                        _ = v48
                                        _ = v49
                                    }
                                    if (!true) {
                                        if (-2147483648 <= 42) {
                                            switch fi39 {
                                            case -101:
                                                v50 := []int{(fi37 >> ((65535) & 63)), (fi39 >> ((1) & 63)), fi37}
                                                _ = v50
                                            case -1:
                                                v1 = v1
                                            }
                                            v51 := []int{(^(fi39)), (fi39 - 1), fi39, 128}
                                            _ = v51
                                        } else {
                                            var v52 float64 = 1e10
                                            v52 += float64(3.14159)
                                            var v53 int64 = -9223372036854775808
                                            _ = v52
                                            _ = v53
                                        }
                                        fmt.Println(fi39, fi39, v3)
                                        if false {
                                            v54 := GenS3{F0: 42, F1: (v3 << ((267) & 63))}
                                            _ = v54
                                        } else {
                                            var v55 int = 937
                                            fmt.Println(fi39)
                                            _ = v55
                                        }
                                    }
                                    _ = v46
                                }
                                var v56 int8 = 0
                                v57 := GenS7{F0: v3, F1: v56, F2: false}
                                _ = fi39
                                _ = v56
                                _ = v57
                            }
                        case 32767:
                            var v58 uint8 = 0
                            _ = v58
                        case 448:
                            switch v1 {
                            case "a":
                                for fi59 := 0; fi59 < 11; fi59++ {
                                    v60 := []int{42, (fi59 + fi37)}
                                    _ = fi59
                                    _ = v60
                                }
                            case "hello":
                                if true {
                                    for fi61 := 0; fi61 < 16; fi61++ {
                                        v3 = 32767
                                        v62 := []int{(fi37 * (fi37 % ((fi61) | 1))), (fi37 | (fi37 >> ((1) & 63)))}
                                        v63 := GenS8{F0: 0}
                                        _ = fi61
                                        _ = v62
                                        _ = v63
                                    }
                                    var v64 float64 = -2.25
                                    _ = v64
                                } else {
                                    var v65 uint64 = 42
                                    if ((42 >= 867) || false) {
                                        v1 = (v1 + v1)
                                    } else {
                                        switch -893 {
                                        case 0:
                                            switch "" {
                                            case "a":
                                                fmt.Println(fi37, v1)
                                            default:
                                                var v66 string = (v1 + v1)
                                                _ = v66
                                            }
                                        case 4294967295:
                                            v65 = v65
                                        case -1:
                                            v65 = uint64(v3)
                                        }
                                    }
                                    var v67 uint16 = (^(v3))
                                    _ = v65
                                    _ = v67
                                }
                            case "\n\t":
                                fmt.Println(fi37)
                            default:
                                var v68 string = v1
                                _ = v68
                            }
                        }
                    default:
                        fmt.Println(fi37, fi37)
                    }
                    if false {
                        v69 := GenS9{F0: 42, F1: 32767}
                        v1 = v1
                        _ = v69
                    }
                    _ = fi37
                    _ = v38
                }
            }
        case -2147483648:
            v70 := GenS10{F0: 2147483647, F1: uint(v3), F2: 710}
            _ = v70
        }
    }
    v2[(v1 + "\n\t")] = (false && (!(!false)))
    var v71 uint16 = (v3 / ((v3) | 1))
    switch v1 {
    case "":
        var v72 int32 = 0
        _ = v72
    case "hello":
        v71 += (^(v3))
    default:
        for fi73 := 0; fi73 < 8; fi73++ {
            v74 := GenS6{F0: 127, F1: -305}
            v75 := []int{fi73, fi73}
            for fi76 := 0; fi76 < 8; fi76++ {
                fmt.Println(v71, v1, v3)
                v1 = ""
                _ = fi76
            }
            _ = fi73
            _ = v74
            _ = v75
        }
    }
    v77 := GenS10{F0: -680, F1: 2, F2: 870}
    v1 = v1
    v71 -= (v3 ^ 383)
    var v78 int16 = 1000
    if false {
        v1 = "foo bar"
        for fi79 := 0; fi79 < 9; fi79++ {
            switch fi79 {
            case 127:
                var v80 uint = 55
                _ = v80
            }
            for fi81 := 0; fi81 < 4; fi81++ {
                var v82 string = "a"
                _ = fi81
                _ = v82
            }
            v83 := []int{2147483647, fi79, fi79, -599}
            _ = fi79
            _ = v83
        }
    }
    for fi84 := 0; fi84 < 8; fi84++ {
        for fi85 := 0; fi85 < 9; fi85++ {
            fmt.Println(fi85)
            v86 := []int{(fi84 + 1), fi85, (fi85 ^ (fi84 ^ (fi84 / ((1) | 1)))), fi85}
            _ = fi85
            _ = v86
        }
        if (int8(fi84) < -128) {
            var v87 int8 = 100
            var v88 string = ("\n\t" + "foo bar")
            var v89 bool = false
            _ = v87
            _ = v88
            _ = v89
        }
        _ = fi84
    }
    switch v1 {
    case "\n\t":
        var v90 uint16 = 574
        _ = v90
    case "":
        for fi91 := 0; fi91 < 19; fi91++ {
            var v92 uint8 = 2
            v93 := GenS11{F0: 2}
            _ = fi91
            _ = v92
            _ = v93
        }
    case "foo bar":
        switch v1 {
        case "foo bar":
            fmt.Println(v3)
        default:
            fmt.Println(v1, v78)
        }
    }
    v94 := []int{799}
    if (180 < 118) {
        var v95 int64 = 42
        fmt.Println(v1, v1, v78)
        _ = v95
    } else {
        for fi96 := 0; fi96 < 20; fi96++ {
            switch fi96 {
            case 799:
                v78 -= (v78 - v78)
            case 240:
                for fi97 := 0; fi97 < 10; fi97++ {
                    var v98 int8 = 1
                    v98 += v98
                    _ = fi97
                    _ = v98
                }
            case 0:
                if true {
                    var v99 float64 = 1.5
                    _ = v99
                } else {
                    if true {
                        if false {
                            fmt.Println(v71)
                            if ((1.5 - float64(1.5)) == 3.14159) {
                                v100 := GenS12{F0: 2147483647}
                                if (!true) {
                                    fmt.Println(fi96, fi96, v1)
                                    v1 = "foo bar"
                                    var v101 int8 = -51
                                    _ = v101
                                } else {
                                    if len(v94) > 0 {
                                        v94[0] = (^(fi96))
                                    }
                                }
                                var v102 int8 = -67
                                _ = v100
                                _ = v102
                            }
                        } else {
                            var v103 int = -865
                            var v104 string = "\n\t"
                            var v105 string = (("" + v104) + ((v1 + v1) + v104))
                            _ = v103
                            _ = v104
                            _ = v105
                        }
                        fmt.Println(v1)
                    }
                    if false {
                        var v106 int16 = 0
                        v1 = v1
                        for fi107 := 0; fi107 < 2; fi107++ {
                            v106 -= v106
                            var v108 int32 = 42
                            v109 := GenS8{F0: 128}
                            if (2147483647 <= 64) {
                                var v110 uint = 18446744073709551615
                                if true {
                                    var v111 bool = (!(v1 <= ""))
                                    _ = v111
                                }
                                _ = v110
                            } else {
                                v112 := GenS6{F0: 1, F1: -1}
                                var v113 uint32 = 370
                                var v114 bool = true
                                _ = v112
                                _ = v113
                                _ = v114
                            }
                            _ = fi107
                            _ = v108
                            _ = v109
                        }
                        _ = v106
                    }
                }
            }
            var v115 float32 = 1e10
            _ = fi96
            _ = v115
        }
    }
    if len(v94) > 0 {
        v94[0] = -49
    }
    v1 = (v1 + "foo bar")
    for fi116 := 0; fi116 < 17; fi116++ {
        _ = len(v2)
        _ = fi116
    }
    v117 := v2["foo bar"]
    var v118 int = -322
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v71:", v71)
    fmt.Println("v78:", v78)
    fmt.Println("slicelen v94:", len(v94), v94)
    _ = v1
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v71
    _ = v77
    _ = v78
    _ = v94
}
