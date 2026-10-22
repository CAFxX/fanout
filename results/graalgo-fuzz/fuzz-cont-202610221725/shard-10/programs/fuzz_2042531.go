package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 string; F1 uint16; F2 uint16 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 int8; F1 string; F2 bool }
type GenS13 struct { F0 uint64; F1 uint32 }
type GenS14 struct { F0 uint16; F1 int }
type GenS15 struct { F0 bool; F1 int8; F2 int }
func genFn0() string {
    v7 := GenS1{F0: 42, F1: 42}
    if true {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS1{F0: 65535, F1: 2}
                if true {
                    v11 := GenS2{F0: 2147483647, F1: fi9, F2: 604}
                    if true {
                        var v12 string = ((("hello" + ("\n\t" + "hello")) + ("a" + ("a" + "hello"))) + ("hello" + ("" + "foo bar")))
                        var v13 float32 = 1e10
                        var v14 uint32 = 2147483648
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        switch fi9 {
                        case 42:
                            if (1 < -978) {
                                v15 := GenS3{F0: 2147483647}
                                _ = v15
                            }
                        }
                    }
                    if (!false) {
                        v16 := GenS4{F0: -1, F1: 42, F2: 4294967295}
                        var v17 string = ("\n\t" + "")
                        _ = v16
                        _ = v17
                    }
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
            v18 := []int{fi8, fi8}
            switch "foo bar" {
            case "\n\t":
                var v19 uint64 = 97
                _ = v19
            default:
                v20 := GenS5{F0: 65535}
                _ = v20
            }
            _ = fi8
            _ = v18
        }
    } else {
        switch -9223372036854775808 {
        case 42:
            v21 := GenS6{F0: 1, F1: 2}
            _ = v21
        case -949:
            if false {
                for fi22 := 0; fi22 < 7; fi22++ {
                    if true {
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 float64 = -2.25
                            var v25 bool = false
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 uint = 255
                        fmt.Println(v26)
                        _ = v26
                    }
                    _ = fi22
                }
                var v27 uint32 = 2147483647
                for fi28 := 0; fi28 < 15; fi28++ {
                    for fi29 := 0; fi29 < 2; fi29++ {
                        v27 += (v27 << ((1) & 63))
                        v30 := []int{(fi29 + (fi28 % ((-1) | 1))), -842, (-(fi29)), 127}
                        _ = fi29
                        _ = v30
                    }
                    _ = fi28
                }
                _ = v27
            } else {
                if true {
                    var v31 float64 = (1.5 - 0.0)
                    if false {
                        if ((!true) && true) {
                            fmt.Println(v31)
                            v31 = v31
                        } else {
                            switch 8 {
                            case 1:
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    v31 += float64(1e10)
                                    _ = fi32
                                }
                            default:
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    for fi34 := 0; fi34 < 18; fi34++ {
                                        var v35 string = ("" + "a")
                                        _ = fi34
                                        _ = v35
                                    }
                                    _ = fi33
                                }
                            }
                            var v36 uint64 = 821
                            _ = v36
                        }
                        if (!(-128 <= 0)) {
                            v31 += v31
                            var v37 float32 = -2.25
                            _ = v37
                        } else {
                            var v38 uint8 = 42
                            var v39 string = "a"
                            _ = v38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 15; fi40++ {
                            var v41 uint16 = 42
                            _ = fi40
                            _ = v41
                        }
                    } else {
                        if (!(("\n\t" + "\n\t") < ("\n\t" + "\n\t"))) {
                            switch "hello" {
                            case "\n\t":
                                v42 := GenS4{F0: 0, F1: 291, F2: 2}
                                _ = v42
                            case "a":
                                var v43 int64 = -1
                                _ = v43
                            case "":
                                fmt.Println(v31)
                            default:
                                switch -629 {
                                case -9223372036854775808:
                                    if (!(2 > 42)) {
                                        var v44 int64 = 1
                                        _ = v44
                                    }
                                case 9223372036854775807:
                                    var v45 int64 = 9223372036854775807
                                    _ = v45
                                case 2147483647:
                                    var v46 uint16 = 428
                                    _ = v46
                                }
                            }
                            v31 = 0.0
                        }
                        var v47 int32 = 1
                        _ = v47
                    }
                    v48 := []int{-9223372036854775808}
                    _ = v31
                    _ = v48
                }
                var v49 uint64 = 77
                var v50 int16 = 607
                _ = v49
                _ = v50
            }
        case 128:
            var v51 bool = true
            _ = v51
        default:
            var v52 string = ((("a" + ("foo bar" + ("" + (((("hello" + ("foo bar" + "")) + "foo bar") + ("foo bar" + "")) + "a")))) + "foo bar") + "hello")
            _ = v52
        }
        if (76 >= 188) {
            for fi53 := 0; fi53 < 17; fi53++ {
                var v54 int8 = 42
                v55 := GenS1{F0: 0, F1: 98}
                _ = fi53
                _ = v54
                _ = v55
            }
            var v56 uint16 = 255
            _ = v56
        } else {
            var v57 int = -84
            for fi58 := 0; fi58 < 18; fi58++ {
                if true {
                    v59 := GenS7{F0: "a", F1: -79}
                    _ = v59
                } else {
                    if (1000 < -549) {
                        v60 := GenS8{F0: "", F1: 1, F2: 689}
                        if (-1 <= -222) {
                            if false {
                                v57 = v57
                            }
                        } else {
                            for fi61 := 0; fi61 < 1; fi61++ {
                                v57 = (^(fi61))
                                v57 = fi58
                                var v62 uint32 = 149
                                v57 = v57
                                v63 := []int{}
                                _ = fi61
                                _ = v62
                                _ = v63
                            }
                        }
                        switch ("hello" + ("hello" + "\n\t")) {
                        case "foo bar":
                            if (2 <= 888) {
                                v57 -= (fi58 - (^(fi58)))
                                fmt.Println(v57)
                                v64 := GenS6{F0: 3, F1: 2}
                                _ = v64
                            } else {
                                v65 := GenS5{F0: 65535}
                                _ = v65
                            }
                        case "\n\t":
                            if (!(92 > 42)) {
                                switch "foo bar" {
                                case "foo bar":
                                    switch 580 {
                                    case 127:
                                        v57 = -13
                                    case 32767:
                                        var v66 uint32 = 422
                                        _ = v66
                                    default:
                                        v67 := GenS6{F0: 42, F1: 247}
                                        _ = v67
                                    }
                                case "":
                                    for fi68 := 0; fi68 < 20; fi68++ {
                                        switch 42 {
                                        case 32768:
                                            v57 -= -696
                                            v57 = -2147483648
                                            switch ("a" + ("foo bar" + "")) {
                                            case "foo bar":
                                                switch ("\n\t" + ("" + "\n\t")) {
                                                case "\n\t":
                                                    switch "a" {
                                                    case "a":
                                                        var v69 int8 = 127
                                                        _ = v69
                                                    case "":
                                                        switch "" {
                                                        case "":
                                                            v70 := GenS9{F0: 183, F1: 244}
                                                            _ = v70
                                                        case "a":
                                                            v71 := []int{4294967295, (fi68 & fi68), fi58, (fi58 + -1)}
                                                            _ = v71
                                                        case "foo bar":
                                                            v57 = fi68
                                                            if (4294967295 <= 1) {
                                                                v57 -= 188
                                                                v72 := []int{(fi68 / ((-9223372036854775808) | 1))}
                                                                _ = v72
                                                            }
                                                        default:
                                                            var v73 float32 = 0.1
                                                            _ = v73
                                                        }
                                                        v57 -= fi68
                                                        v74 := []int{(fi58 | fi68), 32767, fi68}
                                                        _ = v74
                                                    default:
                                                        switch "a" {
                                                        case "\n\t":
                                                            if (false || true) {
                                                                var v75 int32 = 42
                                                                var v76 float32 = -2.25
                                                                v77 := GenS10{F0: 0}
                                                                v57 -= v57
                                                                v75 += v75
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                        case "hello":
                                                            v57 -= 931
                                                            v57 = 1
                                                            v78 := GenS11{F0: 4294967295, F1: (fi58 << ((42) & 63))}
                                                            _ = v78
                                                        case "foo bar":
                                                            v79 := []int{(fi68 / (((fi68 | 42)) | 1)), (-(fi58))}
                                                            _ = v79
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                switch fi58 {
                                                case 127:
                                                    v80 := GenS4{F0: 1, F1: -821, F2: 1}
                                                    _ = v80
                                                case 1:
                                                    var v81 float32 = 1.5
                                                    _ = v81
                                                case 971:
                                                    if false {
                                                        var v82 int16 = -258
                                                        _ = v82
                                                    }
                                                default:
                                                    var v83 int8 = 78
                                                    if true {
                                                        v83 += (-(v83))
                                                        switch 32768 {
                                                        case 42:
                                                            var v84 uint = 255
                                                            _ = v84
                                                        default:
                                                            var v85 uint16 = 255
                                                            _ = v85
                                                        }
                                                        v86 := []int{fi58, fi58}
                                                        _ = v86
                                                    }
                                                    v57 = -9223372036854775808
                                                    _ = v83
                                                }
                                            }
                                        case 128:
                                            if (42 >= 32767) {
                                                var v87 uint16 = 65535
                                                if false {
                                                    var v88 bool = false
                                                    v87 -= v87
                                                    var v89 bool = (3.14159 <= float32(((3.14159 - (((1.5 - (0.0 - 1e10)) * ((3.14159 - 3.14159) + -2.25)) - ((3.14159 + -2.25) + (0.0 + (-2.25 + float64(0.1)))))) - -2.25)))
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    v90 := GenS8{F0: "\n\t", F1: v87, F2: 32767}
                                                    _ = v90
                                                }
                                                _ = v87
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    v91 := []int{127, fi68, v57}
                                                    _ = v91
                                                }
                                            }
                                        case 2147483647:
                                            fmt.Println(fi68, v57)
                                        default:
                                            v92 := GenS11{F0: 708, F1: (-(v57))}
                                            _ = v92
                                        }
                                        v57 = v57
                                        v93 := []int{(fi58 ^ fi58), (fi68 * 65), (-(v57)), (fi58 >> ((712) & 63))}
                                        _ = fi68
                                        _ = v93
                                    }
                                }
                            }
                        case "hello":
                            v94 := GenS2{F0: 957, F1: (fi58 ^ v57), F2: 617}
                            _ = v94
                        }
                        _ = v60
                    } else {
                        var v95 uint8 = 127
                        v57 = 2147483647
                        _ = v95
                    }
                }
                _ = fi58
            }
            if false {
                v96 := []int{(v57 + 2147483647), (v57 / ((v57) | 1)), (-(v57))}
                v57 += v57
                _ = v96
            } else {
                for fi97 := 0; fi97 < 14; fi97++ {
                    switch 622 {
                    case -120:
                        fmt.Println(v57)
                    case 583:
                        v57 -= fi97
                    }
                    _ = fi97
                }
                var v98 float64 = ((1e10 - 1e10) / 1.0)
                v99 := GenS12{F0: 18, F1: "a", F2: false}
                _ = v98
                _ = v99
            }
            _ = v57
        }
        v100 := GenS13{F0: 1, F1: 2}
        _ = v100
    }
    for fi101 := 0; fi101 < 18; fi101++ {
        fmt.Println(fi101)
        for fi102 := 0; fi102 < 9; fi102++ {
            if (655 != 42) {
                fmt.Println(fi101, fi102)
                switch (fi101 / (((-(fi101))) | 1)) {
                case 1:
                    var v103 uint16 = 674
                    _ = v103
                case 42:
                    v104 := GenS8{F0: ("a" + ""), F1: 1, F2: 42}
                    _ = v104
                default:
                    var v105 float32 = (1.5 * float32(float64(float32(-2.25))))
                    _ = v105
                }
                v106 := []int{(fi102 >> ((18446744073709551615) & 63)), 42}
                _ = v106
            }
            if (!false) {
                var v107 float32 = 0.0
                _ = v107
            } else {
                v108 := GenS14{F0: 255, F1: fi102}
                var v109 uint64 = 637
                _ = v108
                _ = v109
            }
            var v110 int8 = 116
            _ = fi102
            _ = v110
        }
        var v111 string = ("foo bar" + "hello")
        _ = fi101
        _ = v111
    }
    v112 := GenS15{F0: false, F1: 100, F2: 128}
    var v113 float32 = (1.5 / -2.25)
    _ = v112
    _ = v113
    _ = v7
    return ("hello" + "foo bar")
}

func main() {
    var v1 float64 = 0.0
    var v2 uint32 = 42
    var v3 uint = 915
    v4 := make(map[uint32]bool)
    v5 := make(map[string]bool)
    v6 := []int{9223372036854775807, 262, 127}
    var v114 uint64 = 0
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi115 := 0; fi115 < 18; fi115++ {
        v2 = (v2 * (^(v2)))
        v116 := GenS9{F0: 124, F1: 42}
        _ = fi115
        _ = v116
    }
    var v117 int32 = 2147483647
    v2 = v2
    v117 = 573
    var v118 string = "foo bar"
    fmt.Println(v114)
    var v119 uint = (^(v3))
    var v120 uint = (^(v119))
    var v121 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v114
    _ = v117
    _ = v118
    _ = v119
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
