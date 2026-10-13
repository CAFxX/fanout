package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 uint16; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 uint; F2 uint8 }
type GenS11 struct { F0 uint32; F1 uint8 }
type GenS12 struct { F0 uint64; F1 uint; F2 int16 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 float32 = (float32(float64(-2.25)) + (0.1 - (v1 / v1)))
    v4 := make(map[int]bool)
    var v5 int8 = -1
    var v6 int64 = -1
    var v7 float32 = v3
    for fi8 := 0; fi8 < 2; fi8++ {
        if true {
            if true {
                v5 = -128
                if false {
                    var v9 int = -9223372036854775808
                    _ = v9
                } else {
                    v10 := GenS1{F0: v5}
                    v2["hello"] = fi8
                    switch (fi8 >> ((42) & 63)) {
                    case 875:
                        fmt.Println(v1, v5)
                    case -1:
                        v11 := []int{fi8, (fi8 & (fi8 & (fi8 & fi8))), 127, fi8}
                        _ = v11
                    default:
                        v12 := GenS1{F0: 57}
                        _ = v12
                    }
                    _ = v10
                }
            } else {
                if (1 != 2) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        fmt.Println(fi8, v5, v5)
                        _ = fi13
                    }
                }
                var v14 string = ""
                _ = v14
            }
            v15 := v2[(("a" + (("foo bar" + "\n\t") + "hello")) + ((("" + "") + "hello") + "a"))]
            _ = v15
        }
        if true {
            fmt.Println(v3, v1)
            if (fi8 != (^(fi8))) {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 int64 = v6
                    switch ("a" + "hello") {
                    case "a":
                        v18 := GenS1{F0: v5}
                        _ = v18
                    }
                    _ = fi16
                    _ = v17
                }
                for fi19 := 0; fi19 < 7; fi19++ {
                    if true {
                        fmt.Println(v6)
                        var v20 int8 = (v5 | v5)
                        if (18446744073709551615 == 42) {
                            var v21 uint8 = 88
                            v1 = float32(0.0)
                            _ = v21
                        } else {
                            fmt.Println(fi8, v6, fi8)
                            if true {
                                var v22 int32 = -1
                                switch "hello" {
                                case "a":
                                    var v23 uint = 255
                                    _ = v23
                                default:
                                    var v24 int8 = (-(v5))
                                    _ = v24
                                }
                                v5 -= v20
                                _ = v22
                            }
                        }
                        _ = v20
                    } else {
                        v5 = v5
                        v25 := v2["foo bar"]
                        _ = v25
                    }
                    if false {
                        v26 := GenS1{F0: (v5 >> ((255) & 63))}
                        _ = v26
                    } else {
                        var v27 uint64 = 314
                        var v28 int64 = 42
                        if (-169 > 1000) {
                            fmt.Println(v3, v5)
                            if false {
                                v1 = (v7 * v3)
                                fmt.Println(v27)
                            }
                        }
                        v29 := GenS2{F0: 2, F1: 1}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    v30 := v4[fi8]
                    switch 0 {
                    case 0:
                        if v30 {
                            if v30 {
                                var v31 int64 = 1
                                v32 := v4[-2147483648]
                                if ((v6 - -767) == -9223372036854775808) {
                                    var v33 int16 = 1000
                                    if (v33 > 42) {
                                        var v34 uint64 = 1
                                        _ = v34
                                    }
                                    var v35 int32 = 1
                                    var v36 float64 = 1.5
                                    _ = v33
                                    _ = v35
                                    _ = v36
                                } else {
                                    if false {
                                        delete(v4, (fi8 % ((fi19) | 1)))
                                        switch fi8 {
                                        case 127:
                                            var v37 int16 = -32768
                                            _ = v37
                                        default:
                                            fmt.Println(v30, v5)
                                        }
                                    } else {
                                        v38 := GenS3{F0: (v5 * (v5 - (v5 - (v5 * v5))))}
                                        v6 += (v31 >> ((299) & 63))
                                        _ = v38
                                    }
                                    v1 -= ((-2.25 - 3.14159) - v7)
                                    var v39 int16 = 1000
                                    v30 = v32
                                    _ = v39
                                }
                                _ = v31
                                _ = v32
                            }
                            var v40 uint8 = 127
                            _ = v40
                        } else {
                            if true {
                                switch "foo bar" {
                                case "a":
                                    v30 = false
                                    v5 = int8(v6)
                                    _ = len(v2)
                                case "\n\t":
                                    var v41 int64 = v6
                                    _ = v41
                                }
                                v7 = ((v1 / (v7 + v3)) + 0.0)
                                v6 -= 0
                            } else {
                                if v30 {
                                    var v42 int16 = -760
                                    var v43 int8 = (v5 / (((v5 + (v5 & v5))) | 1))
                                    switch "hello" {
                                    case "foo bar":
                                        var v44 float32 = (3.14159 + v3)
                                        _ = v44
                                    default:
                                        v6 += v6
                                    }
                                    _ = v42
                                    _ = v43
                                }
                            }
                            var v45 uint64 = 369
                            _ = v45
                        }
                    case -9223372036854775808:
                        v46 := v2[""]
                        _ = v46
                    default:
                        if (1 != 2) {
                            v47 := []int{(fi19 ^ (fi8 ^ fi19))}
                            switch -2147483648 {
                            case -217:
                                if true {
                                    v48 := v4[-9223372036854775808]
                                    var v49 string = ""
                                    _ = v48
                                    _ = v49
                                }
                            }
                            fmt.Println(fi19, v30, v5)
                            _ = v47
                        } else {
                            var v50 float32 = v3
                            var v51 uint8 = 127
                            _ = v50
                            _ = v51
                        }
                    }
                    fmt.Println(v1, v3)
                    _ = fi19
                    _ = v30
                }
                v52 := GenS4{F0: ("foo bar" + ("foo bar" + "foo bar"))}
                _ = v52
            } else {
                var v53 uint8 = 185
                v1 = (v3 * float32(1e10))
                v3 = (3.14159 * -2.25)
                _ = v53
            }
        } else {
            for fi54 := 0; fi54 < 5; fi54++ {
                switch fi54 {
                case 32767:
                    var v55 uint64 = 1
                    _ = v55
                case 0:
                    v6 += v6
                }
                v56 := GenS2{F0: 362, F1: (v5 | (v5 & v5))}
                v5 = (v5 << ((0) & 63))
                _ = fi54
                _ = v56
            }
            if false {
                switch "" {
                case "a":
                    if (v6 < (v6 ^ v6)) {
                        v7 = v7
                        _ = len(v4)
                    } else {
                        v3 += (3.14159 + -2.25)
                    }
                case "hello":
                    if (fi8 >= (fi8 ^ (fi8 & -9223372036854775808))) {
                        switch fi8 {
                        case 42:
                            var v57 uint32 = 270
                            _ = v57
                        case 310:
                            v58 := GenS3{F0: (-(v5))}
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 4; fi59++ {
                            v60 := GenS1{F0: v5}
                            switch "foo bar" {
                            case "hello":
                                if (false || (!true)) {
                                    _ = len(v2)
                                }
                            }
                            _ = fi59
                            _ = v60
                        }
                        var v61 uint = 1
                        _ = v61
                    } else {
                        v6 = v6
                    }
                default:
                    v62 := []int{}
                    _ = v62
                }
            } else {
                for fi63 := 0; fi63 < 7; fi63++ {
                    if (true && (!false)) {
                        v64 := []int{fi8, (fi63 / ((32767) | 1))}
                        _ = v64
                    } else {
                        v4[(fi8 | fi63)] = (false && false)
                        var v65 int16 = 1
                        v6 -= (^(v6))
                        switch "hello" {
                        case "a":
                            v66 := GenS2{F0: 856, F1: v5}
                            _ = v66
                        case "foo bar":
                            v67 := GenS1{F0: v5}
                            _ = v67
                        case "\n\t":
                            if false {
                                v3 -= 1.5
                                var v68 int32 = -2147483648
                                v3 = v1
                                var v69 int8 = (^(v5))
                                v70 := GenS2{F0: 2147483648, F1: v69}
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        }
                        _ = v65
                    }
                    if false {
                        switch "\n\t" {
                        case "":
                            v71 := []int{fi8, (fi63 | fi63)}
                            _ = v71
                        }
                    }
                    v6 += -1
                    _ = fi63
                }
                if false {
                    var v72 string = ("foo bar" + ("" + "\n\t"))
                    _ = v72
                } else {
                    v73 := []int{fi8}
                    _ = v73
                }
            }
        }
        var v74 float64 = (-2.25 + 0.0)
        _ = fi8
        _ = v74
    }
    v75 := []int{-1, 2147483647, 9223372036854775807, 2147483647}
    if (!("a" == "hello")) {
        v3 = v1
        fmt.Println(v1)
        if (!(42 != 4294967295)) {
            var v76 int16 = 52
            _ = v76
        } else {
            v75 = append(v75, 127)
            if false {
                var v77 int64 = (v6 & (v6 + (-(v6))))
                if false {
                    v6 -= 532
                    fmt.Println(v1, v6, v77)
                }
                _ = v77
            }
        }
    }
    if ((true && (true && ((!false) || (((v6 - -941) >= v6) || false)))) || (109 <= 2)) {
        v7 = v1
        v3 = (1.5 - ((0.0 / v3) * v7))
        for fi78 := 0; fi78 < 9; fi78++ {
            if false {
                switch "a" {
                case "\n\t":
                    v6 = (v6 | v6)
                case "a":
                    var v79 int64 = v6
                    _ = v79
                }
                if (0 == 32767) {
                    v6 = -1
                    switch -2147483648 {
                    case 32768:
                        fmt.Println(v7, v7, v6)
                    }
                } else {
                    v1 = (0.1 - v7)
                }
            }
            _ = fi78
        }
    } else {
        for fi80 := 0; fi80 < 4; fi80++ {
            for fi81 := 0; fi81 < 14; fi81++ {
                v82 := GenS5{F0: (v6 & v6)}
                v7 = ((float32((3.14159 + 3.14159)) * 0.0) - ((1.5 * ((v7 * -2.25) / v3)) * (v1 * v7)))
                _ = fi81
                _ = v82
            }
            for fi83 := 0; fi83 < 13; fi83++ {
                var v84 uint8 = 2
                if (-37 >= 1) {
                    if true {
                        v85 := GenS3{F0: v5}
                        v86 := GenS6{F0: (255 >= v84)}
                        _ = v85
                        _ = v86
                    }
                    v87 := GenS7{F0: ((("\n\t" + "") + "hello") + "hello"), F1: v6}
                    _ = v87
                } else {
                    var v88 uint8 = (v84 & 24)
                    v89 := GenS8{F0: 32767, F1: 42}
                    _ = v88
                    _ = v89
                }
                fmt.Println(v3)
                switch fi80 {
                case 32767:
                    switch "hello" {
                    case "":
                        var v90 int = (-(fi80))
                        _ = v90
                    case "a":
                        v5 -= v5
                    case "\n\t":
                        if ((4294967295 == 621) || false) {
                            var v91 uint64 = 549
                            if len(v75) > 0 {
                                v75[0] = fi80
                            }
                            _ = v91
                        } else {
                            v92 := GenS9{F0: (!true)}
                            v93 := GenS7{F0: "foo bar", F1: (v6 - v6)}
                            if false {
                                _ = len(v75)
                                v7 -= 0.1
                                v1 += v3
                                fmt.Println(v84, v3)
                            }
                            _ = v92
                            _ = v93
                        }
                    default:
                        var v94 uint32 = 42
                        _ = v94
                    }
                case 42:
                    if true {
                        fmt.Println(v84, v6, v84)
                        v5 = -128
                        switch ("hello" + "") {
                        case "hello":
                            v3 += v1
                        case "a":
                            switch ("hello" + "foo bar") {
                            case "a":
                                if false {
                                    v84 += 128
                                    if false {
                                        var v95 int32 = 42
                                        v96 := GenS2{F0: 65535, F1: 127}
                                        v3 = (v7 / ((((0.1 - ((1.5 * 3.14159) / (v3 * v3))) + v7) - -2.25) / v3))
                                        _ = v95
                                        _ = v96
                                    }
                                    var v97 float32 = 0.0
                                    _ = v97
                                } else {
                                    var v98 string = "a"
                                    v99 := GenS5{F0: v6}
                                    v100 := GenS10{F0: 127, F1: 0, F2: v84}
                                    _ = v100
                                    _ = v98
                                    _ = v99
                                }
                            }
                        default:
                            if len(v75) >= 2 {
                                v75 = v75[1:len(v75)]
                            }
                        }
                        if len(v75) > 0 {
                            v75[0] = (-(fi80))
                        }
                        var v101 int16 = 316
                        v101 = v101
                        if false {
                            switch fi80 {
                            case 0:
                                var v102 float64 = (0.0 - (1.5 + ((0.0 * (-2.25 / 1.5)) - float64(0.1))))
                                _ = v102
                            case 684:
                                v5 = (v5 / (((-(v5))) | 1))
                            case 42:
                                switch fi80 {
                                case 32768:
                                    _ = len(v75)
                                    if (!true) {
                                        v7 -= 0.0
                                        if true {
                                            switch "hello" {
                                            case "":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v103 := GenS2{F0: 2, F1: v5}
                                                    _ = v103
                                                case "":
                                                    switch 0 {
                                                    case 32768:
                                                        if false {
                                                            if len(v75) >= 2 {
                                                                v75 = v75[1:len(v75)]
                                                            }
                                                        }
                                                    case 2147483647:
                                                        if true {
                                                            var v104 uint16 = 255
                                                            var v105 uint16 = v104
                                                            v106 := v2[("\n\t" + "hello")]
                                                            v107 := GenS11{F0: 2147483648, F1: 147}
                                                            _ = v104
                                                            _ = v105
                                                            _ = v106
                                                            _ = v107
                                                        }
                                                    default:
                                                        switch "" {
                                                        case "foo bar":
                                                            fmt.Println(v5)
                                                        case "hello":
                                                            fmt.Println(v84, v101, v101)
                                                        }
                                                    }
                                                }
                                            case "foo bar":
                                                v7 = (float32(((((1e10 / ((-2.25 + -2.25) * (0.1 * ((float64(v1) * float64(v3)) - 1.5)))) / 1.0) * 0.1) / (3.14159 / 1.5))) * v3)
                                            case "hello":
                                                switch fi80 {
                                                case 128:
                                                    var v108 int8 = (v5 % (((v5 & -63)) | 1))
                                                    _ = v108
                                                default:
                                                    switch fi80 {
                                                    case 42:
                                                        var v109 int32 = 908
                                                        _ = v109
                                                    default:
                                                        if (((v5 * v5) <= (v5 ^ 42)) && false) {
                                                            if len(v75) >= 2 {
                                                                v75 = v75[1:len(v75)]
                                                            }
                                                            v5 = (v5 - (v5 * 8))
                                                            if (!true) {
                                                                fmt.Println(v3, fi80)
                                                                fmt.Println(fi83, v3)
                                                                v110 := v4[(fi80 * fi80)]
                                                                _ = v110
                                                            }
                                                        } else {
                                                            v111 := GenS1{F0: v5}
                                                            var v112 int64 = v6
                                                            fmt.Println(v3)
                                                            _ = len(v4)
                                                            v113 := GenS12{F0: 4294967295, F1: 750, F2: v101}
                                                            _ = v111
                                                            _ = v112
                                                            _ = v113
                                                        }
                                                    }
                                                }
                                            }
                                            fmt.Println(fi83, v5, v101)
                                            var v114 string = "a"
                                            _ = v114
                                        }
                                    } else {
                                        fmt.Println(v3)
                                    }
                                case -9223372036854775808:
                                    fmt.Println(v84)
                                case 0:
                                    var v115 float64 = 1e10
                                    _ = v115
                                default:
                                    v7 = float32(0.1)
                                }
                            default:
                                v5 += v5
                            }
                            var v116 bool = (!(!(!true)))
                            v116 = ((1e10 + -2.25) <= v3)
                            _ = v116
                        }
                        _ = v101
                    }
                case 9223372036854775807:
                    fmt.Println(v1)
                }
                _ = fi83
                _ = v84
            }
            v5 = v5
            _ = fi80
        }
        v6 -= -9223372036854775808
    }
    var v117 uint64 = 4294967295
    if (-273 < (v6 / ((v6) | 1))) {
        v118 := GenS3{F0: (v5 - v5)}
        _ = v118
    }
    if true {
        fmt.Println(v1, v117)
    } else {
        v1 -= -2.25
        v2["hello"] = 32767
    }
    v119 := GenS7{F0: ("foo bar" + "hello"), F1: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v75:", len(v75), v75)
    _ = v1
    _ = v117
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v75
}
