package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 int; F2 int }
type GenS10 struct { F0 int64; F1 int8 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint }
func genFn0(p0 string) string {
    fmt.Println(p0)
    v6 := GenS1{F0: 1, F1: 42}
    v7 := GenS2{F0: -24, F1: 42, F2: 128}
    _ = p0
    _ = v6
    _ = v7
    return p0
}
func genFn1() uint16 {
    var v8 int16 = 32767
    v8 = (v8 | v8)
    if false {
        if (((-2.25 * 0.1) * 3.14159) != 1.5) {
            var v9 int = -1
            v10 := GenS2{F0: 1, F1: 219, F2: 127}
            _ = v10
            _ = v9
        } else {
            v11 := GenS2{F0: 68, F1: 2, F2: 127}
            _ = v11
        }
        if true {
            v12 := []int{9223372036854775807}
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 uint64 = 4294967295
                var v15 bool = true
                var v16 int64 = 9223372036854775807
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
            _ = v12
        }
        v17 := GenS1{F0: 0, F1: 18446744073709551615}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 4; fi18++ {
            switch "\n\t" {
            case "hello":
                v19 := GenS3{F0: ""}
                _ = v19
            case "foo bar":
                if (255 <= 937) {
                    if ((false || false) && (!true)) {
                        if (42 >= 1) {
                            switch ("hello" + (("" + "") + (("\n\t" + ("\n\t" + "hello")) + "\n\t"))) {
                            case "hello":
                                v8 = v8
                            case "":
                                fmt.Println(fi18)
                            }
                            switch fi18 {
                            case -2147483648:
                                var v20 uint16 = 42
                                _ = v20
                            }
                        }
                    }
                    v21 := []int{fi18, -53, 127, fi18}
                    _ = v21
                } else {
                    var v22 int = 32768
                    _ = v22
                }
            }
            v23 := GenS1{F0: 42, F1: 2}
            _ = fi18
            _ = v23
        }
        switch (("foo bar" + "foo bar") + ("hello" + ("" + ""))) {
        case "a":
            if (-929 >= -1) {
                fmt.Println(v8)
                fmt.Println(v8)
            } else {
                if (!true) {
                    fmt.Println(v8)
                    var v24 int8 = 0
                    _ = v24
                }
            }
        }
        if true {
            if ((0 <= 0) && (!false)) {
                for fi25 := 0; fi25 < 6; fi25++ {
                    v26 := []int{fi25, fi25}
                    if len(v26) > 0 {
                        v26[0] = fi25
                    }
                    if ("" != "") {
                        v27 := GenS2{F0: -128, F1: 0, F2: 255}
                        _ = v27
                    }
                    _ = fi25
                    _ = v26
                }
                v28 := GenS4{F0: 144, F1: 4294967295}
                if ("a" == (("" + ("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + ("a" + "")))) + ("" + "a"))) {
                    v8 = 1
                    var v29 float32 = float32((1e10 - -2.25))
                    if false {
                        var v30 int8 = 42
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            var v32 int64 = -9223372036854775808
                            v33 := GenS3{F0: "\n\t"}
                            fmt.Println(v32, v32)
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 1; fi34++ {
                            var v35 int64 = 322
                            _ = fi34
                            _ = v35
                        }
                    }
                    _ = v29
                } else {
                    fmt.Println(v8)
                    if false {
                        v36 := GenS5{F0: 4294967295, F1: 1, F2: 0}
                        _ = v36
                    } else {
                        v8 = -32768
                        switch ("" + "") {
                        case "foo bar":
                            var v37 uint16 = 458
                            _ = v37
                        case "":
                            v8 += (v8 >> ((932) & 63))
                        default:
                            v38 := GenS1{F0: 511, F1: 101}
                            _ = v38
                        }
                        var v39 int = -2147483648
                        _ = v39
                    }
                }
                _ = v28
            } else {
                var v40 uint32 = 42
                if true {
                    var v41 int64 = 42
                    _ = v41
                }
                if true {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        v8 = v8
                        _ = fi42
                    }
                    v40 += v40
                    fmt.Println(v8)
                } else {
                    for fi43 := 0; fi43 < 4; fi43++ {
                        if true {
                            v44 := []int{-583, fi43, (fi43 ^ (^(fi43)))}
                            if (2147483648 >= v40) {
                                for fi45 := 0; fi45 < 17; fi45++ {
                                    v46 := GenS4{F0: 0, F1: 2}
                                    if true {
                                        v40 += 510
                                        if (823 < uint64(v40)) {
                                            var v47 int8 = -1
                                            _ = v47
                                        }
                                        v40 = v40
                                        v48 := GenS5{F0: 219, F1: v40, F2: v40}
                                        var v49 int16 = (v8 + 32767)
                                        v8 = (v49 * int16(v40))
                                        _ = v48
                                        _ = v49
                                    } else {
                                        fmt.Println(v8, fi43, fi45)
                                        v50 := GenS5{F0: 828, F1: 65535, F2: (^(v40))}
                                        switch "a" {
                                        case "foo bar":
                                            v51 := GenS1{F0: 42, F1: 0}
                                            _ = v51
                                        default:
                                            var v52 int16 = (v8 - (v8 / ((174) | 1)))
                                            _ = v52
                                        }
                                        _ = v50
                                    }
                                    v8 = v8
                                    _ = fi45
                                    _ = v46
                                }
                                for fi53 := 0; fi53 < 19; fi53++ {
                                    var v54 string = "foo bar"
                                    _ = fi53
                                    _ = v54
                                }
                                _ = len(v44)
                            }
                            _ = v44
                        }
                        for fi55 := 0; fi55 < 8; fi55++ {
                            switch "hello" {
                            case "\n\t":
                                v56 := GenS6{F0: 0, F1: v40, F2: 42}
                                _ = v56
                            case "":
                                v40 += (^(v40))
                            default:
                                var v57 uint32 = v40
                                _ = v57
                            }
                            switch "a" {
                            case "":
                                if (!true) {
                                    v58 := []int{fi55, (^(fi55)), fi55}
                                    _ = v58
                                } else {
                                    if (float32(0.1) <= 1e10) {
                                        v59 := GenS2{F0: -1, F1: 2, F2: 127}
                                        if (!false) {
                                            if false {
                                                var v60 uint16 = 506
                                                var v61 uint16 = 338
                                                switch fi43 {
                                                case -1:
                                                    if true {
                                                        v8 = v8
                                                    } else {
                                                        fmt.Println(v61)
                                                    }
                                                default:
                                                    if true {
                                                        v40 += 704
                                                        v8 = (-(v8))
                                                        v40 = v40
                                                        var v62 uint16 = (v60 & v60)
                                                        v63 := []int{(fi55 | 1)}
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                v64 := []int{32768, (-(fi43)), 977, 42}
                                                v61 = 601
                                                var v65 int = -344
                                                var v66 int8 = -91
                                                v66 = v66
                                                _ = v60
                                                _ = v61
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                            if true {
                                                if (128 >= 2) {
                                                    var v67 uint64 = 0
                                                    v8 = v8
                                                    _ = v67
                                                }
                                                var v68 int8 = 127
                                                _ = v68
                                            } else {
                                                v69 := GenS6{F0: 0, F1: 644, F2: 255}
                                                if (2147483647 > 2) {
                                                    var v70 bool = (!(((4294967295 < 1) && true) || false))
                                                    v8 = (v8 ^ (v8 * (-(v8))))
                                                    _ = v70
                                                }
                                                _ = v69
                                            }
                                            var v71 bool = false
                                            _ = v71
                                        }
                                        _ = v59
                                    } else {
                                        var v72 int16 = (v8 & (v8 & v8))
                                        v73 := GenS7{F0: v40, F1: false}
                                        v8 += 917
                                        switch "a" {
                                        case "foo bar":
                                            v74 := []int{(fi55 - fi43), fi43, 134, (fi55 % ((2147483647) | 1))}
                                            _ = v74
                                        case "hello":
                                            var v75 uint64 = 42
                                            _ = v75
                                        case "":
                                            v76 := []int{}
                                            _ = v76
                                        }
                                        _ = v72
                                        _ = v73
                                    }
                                    v8 = (v8 | v8)
                                    var v77 bool = ((2 >= 383) && false)
                                    _ = v77
                                }
                            }
                            v40 += v40
                            _ = fi55
                        }
                        for fi78 := 0; fi78 < 10; fi78++ {
                            if true {
                                var v79 int16 = 145
                                _ = v79
                            } else {
                                v8 -= 1
                            }
                            v80 := []int{(fi43 | (fi78 ^ (fi43 | 1))), fi78, fi78, (fi43 - -1)}
                            fmt.Println(v40)
                            _ = fi78
                            _ = v80
                        }
                        _ = fi43
                    }
                    fmt.Println(v8, v40)
                    var v81 bool = (!false)
                    _ = v81
                }
                _ = v40
            }
            if (128 <= 127) {
                switch "hello" {
                case "a":
                    switch -206 {
                    case 128:
                        if true {
                            var v82 uint8 = 128
                            fmt.Println(v82, v8)
                            _ = v82
                        } else {
                            v8 = v8
                            if true {
                                v8 = (-(v8))
                                v83 := GenS8{F0: 692}
                                _ = v83
                            } else {
                                v84 := []int{9223372036854775807}
                                v8 = 1
                                v8 = v8
                                _ = v84
                            }
                            var v85 uint32 = 172
                            _ = v85
                        }
                    case 848:
                        for fi86 := 0; fi86 < 18; fi86++ {
                            var v87 int8 = -71
                            _ = fi86
                            _ = v87
                        }
                    }
                case "":
                    v88 := []int{42, 42, 1}
                    _ = v88
                case "hello":
                    var v89 int32 = -2147483648
                    _ = v89
                default:
                    v90 := GenS1{F0: 18446744073709551615, F1: 232}
                    _ = v90
                }
                var v91 uint64 = 168
                _ = v91
            }
            for fi92 := 0; fi92 < 9; fi92++ {
                v93 := GenS8{F0: 0}
                _ = fi92
                _ = v93
            }
        } else {
            v8 += v8
            fmt.Println(v8)
            for fi94 := 0; fi94 < 19; fi94++ {
                fmt.Println(v8, fi94)
                v95 := GenS8{F0: 1}
                _ = fi94
                _ = v95
            }
        }
    }
    for fi96 := 0; fi96 < 1; fi96++ {
        switch ("a" + ("foo bar" + ("foo bar" + ((("" + "a") + "foo bar") + "hello")))) {
        case "foo bar":
            v97 := []int{fi96, 4294967295}
            _ = v97
        }
        v8 = v8
        _ = fi96
    }
    _ = v8
    return 0
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{("hello" + ("foo bar" + "hello")), "", "a", ("" + "hello")}
    v3 := make(map[int]string)
    v4 := []int{32767, 0}
    v5 := make(map[int]int)
    var v98 float64 = ((((float64((v1 / v1)) * (0.1 / 0.1)) / 1e10) + 0.1) - -2.25)
    v1 -= v1
    var v99 uint16 = 42
    v1 = ((v1 - 0.0) * -2.25)
    var v100 string = "a"
    fmt.Println(v99, v100)
    for fi101 := 0; fi101 < 10; fi101++ {
        switch "a" {
        case "\n\t":
            v99 = 891
        case "a":
            for fi102 := 0; fi102 < 14; fi102++ {
                if true {
                    var v103 uint8 = 1
                    fmt.Println(fi102, fi101)
                    _ = v103
                } else {
                    switch (fi101 | fi102) {
                    case 1:
                        v104 := GenS9{F0: 751, F1: fi102, F2: (fi102 + fi101)}
                        if (0 <= 0) {
                            var v105 int64 = -9223372036854775808
                            _ = v105
                        }
                        _ = v104
                    }
                    var v106 int32 = 489
                    var v107 float32 = 1.5
                    _ = v106
                    _ = v107
                }
                var v108 uint16 = v99
                v108 = (v99 % (((v99 | 1)) | 1))
                _ = fi102
                _ = v108
            }
        default:
            v1 -= float32((v98 / v98))
        }
        if true {
            switch -405 {
            case 9223372036854775807:
                for fi109 := 0; fi109 < 1; fi109++ {
                    fmt.Println(v1, v99, fi101)
                    v5[fi101] = fi101
                    _ = fi109
                }
            case -1:
                v110 := GenS6{F0: 42, F1: 65535, F2: 18446744073709551615}
                _ = v110
            case 32767:
                if false {
                    v99 -= (^(v99))
                    v111 := GenS1{F0: 14, F1: 2}
                    v100 = ""
                    _ = v111
                } else {
                    var v112 bool = false
                    if (!false) {
                        v113 := GenS10{F0: -1, F1: 42}
                        fmt.Println(v100, v98, v99)
                        if v112 {
                            v1 = v1
                        }
                        _ = v113
                    }
                    _ = v112
                }
            default:
                switch fi101 {
                case -359:
                    v100 = v100
                case -9223372036854775808:
                    var v114 float32 = (0.1 / v1)
                    _ = v114
                case 12:
                    var v115 uint = 4294967295
                    _ = v115
                default:
                    var v116 uint32 = 2147483648
                    _ = v116
                }
            }
            for fi117 := 0; fi117 < 9; fi117++ {
                v1 += 3.14159
                switch v100 {
                case "hello":
                    v118 := v5[fi101]
                    if (v118 > fi101) {
                        var v119 uint = 0
                        if (false || (0 != 42)) {
                            var v120 int32 = 543
                            v99 += (v99 * (v99 & v99))
                            _ = v120
                        } else {
                            var v121 int = v118
                            fmt.Println(fi101, v118, v119)
                            _ = v121
                        }
                        v122 := GenS6{F0: 9223372036854775807, F1: 2147483648, F2: 18446744073709551615}
                        _ = v119
                        _ = v122
                    }
                    var v123 float32 = v1
                    _ = v118
                    _ = v123
                case "a":
                    var v124 int16 = 0
                    _ = v124
                default:
                    fmt.Println(fi117, v1, v98)
                }
                v125 := GenS11{F0: (v99 / ((v99) | 1))}
                if (!false) {
                    if ((1e10 < float64(v1)) && false) {
                        var v126 uint32 = 0
                        var v127 string = v100
                        if (42 > 758) {
                            v127 = (v100 + (("\n\t" + v100) + v127))
                        }
                        _ = v126
                        _ = v127
                    } else {
                        v98 = v98
                    }
                    if (genFn1() >= 333) {
                        v98 = 1e10
                        v128 := GenS12{F0: 767}
                        v5[825] = fi101
                        _ = v128
                    }
                    v98 -= (1e10 * ((3.14159 + (v98 / (-2.25 / v98))) / (3.14159 + v98)))
                } else {
                    v129 := GenS2{F0: int8(v99), F1: 1, F2: 225}
                    _ = v129
                }
                _ = fi117
                _ = v125
            }
        } else {
            var v130 uint16 = v99
            _ = len(v3)
            _ = v130
        }
        for fi131 := 0; fi131 < 20; fi131++ {
            fmt.Println(v1, v100)
            v98 = (-2.25 * v98)
            _ = fi131
        }
        _ = fi101
    }
    var v132 uint64 = 802
    v132 = v132
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v132:", v132)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v132
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v98
    _ = v99
}
