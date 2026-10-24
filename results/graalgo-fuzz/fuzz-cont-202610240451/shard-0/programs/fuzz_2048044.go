package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 bool) uint {
    switch 32767 {
    case 776:
        if p0 {
            v5 := []int{-470}
            fmt.Println(p0)
            for fi6 := 0; fi6 < 3; fi6++ {
                var v7 uint64 = 315
                var v8 uint16 = 0
                var v9 uint8 = 1
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
            _ = v5
        } else {
            fmt.Println(p0)
            var v10 int32 = 2147483647
            fmt.Println(v10, p0)
            _ = v10
        }
    case -907:
        v11 := []int{-9223372036854775808}
        _ = v11
    }
    var v12 string = "foo bar"
    var v13 float32 = (float32((-2.25 - 1.5)) - 0.0)
    var v14 uint32 = 383
    v15 := []int{2147483647, -2147483648}
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 607
}
func genFn1() string {
    if true {
        if true {
            if true {
                v16 := GenS1{F0: "", F1: 4294967295, F2: 2}
                _ = v16
            }
        } else {
            v17 := []int{9223372036854775807, 32767, 1}
            for fi18 := 0; fi18 < 16; fi18++ {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                fmt.Println(fi18)
                var v19 float64 = float64(((((float32(3.14159) * 1.5) + (0.1 - (((1.5 - float32(((((-2.25 * 1e10) / 1e10) * 0.0) * float64(1.5)))) + 0.0) * 0.1))) / ((0.0 / (1.5 - 0.1)) - (1e10 + -2.25))) - 1e10))
                _ = fi18
                _ = v19
            }
            var v20 int16 = -1
            _ = v17
            _ = v20
        }
        var v21 uint32 = 2
        _ = v21
    }
    switch "foo bar" {
    case "foo bar":
        for fi22 := 0; fi22 < 12; fi22++ {
            fmt.Println(fi22)
            var v23 string = "\n\t"
            var v24 int = (fi22 % ((128) | 1))
            _ = fi22
            _ = v23
            _ = v24
        }
    case "hello":
        for fi25 := 0; fi25 < 14; fi25++ {
            if true {
                fmt.Println(fi25)
                if (101 >= 42) {
                    v26 := []int{fi25}
                    var v27 bool = false
                    if (1 > -32768) {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            v27 = true
                            if (0 != 0) {
                                var v29 uint16 = 1
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v27)
                                default:
                                    if false {
                                        var v30 uint64 = 1
                                        if ((("foo bar" + "\n\t") + "foo bar") > "a") {
                                            v27 = false
                                            if false {
                                                v27 = true
                                                var v31 uint32 = 2
                                                _ = v31
                                            } else {
                                                var v32 int32 = 2147483647
                                                var v33 uint16 = v29
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        _ = v30
                                    } else {
                                        var v34 float32 = float32((float64(0.1) - 0.0))
                                        v27 = ((1 != 128) && ((v29 << ((856) & 63)) > (v29 - 255)))
                                        _ = v34
                                    }
                                }
                                if len(v26) >= 2 {
                                    v26 = v26[1:len(v26)]
                                }
                                var v35 uint64 = 281
                                if v27 {
                                    var v36 uint32 = 65535
                                    if (((v36 + (v36 - 667)) == (v36 / ((v36) | 1))) || (true && (!(4294967295 >= 2)))) {
                                        var v37 string = "hello"
                                        if v27 {
                                            v37 = v37
                                            var v38 uint = genFn0((false || (v27 || (65535 >= 844))))
                                            _ = v38
                                        }
                                        _ = v37
                                    }
                                    _ = v36
                                }
                                v35 -= (v35 % (((v35 + (v35 ^ (v35 - v35)))) | 1))
                                _ = v29
                                _ = v35
                            }
                            switch fi28 {
                            case 0:
                                var v39 int = -841
                                _ = v39
                            }
                            _ = fi28
                        }
                    } else {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                    _ = v26
                    _ = v27
                }
                var v40 int8 = 0
                _ = v40
            }
            _ = fi25
        }
    default:
        v41 := GenS2{F0: 1}
        _ = v41
    }
    return ("" + "")
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]uint32)
    var v3 int8 = 42
    var v4 int = -924
    v42 := GenS3{F0: v4, F1: 1000}
    switch "" {
    case "foo bar":
        if false {
            var v43 uint = 42
            _ = v43
        } else {
            v3 += (^(v3))
            switch "foo bar" {
            case "a":
                var v44 uint16 = 408
                _ = v44
            case "\n\t":
                var v45 uint32 = 42
                _ = v45
            }
        }
    case "":
        _ = len(v2)
    }
    v46 := []int{v4, v4, 127, -2147483648}
    v47 := GenS2{F0: 4294967295}
    switch ("a" + (("" + "") + "\n\t")) {
    case "\n\t":
        v48 := GenS4{F0: -1, F1: 42, F2: 255}
        _ = v48
    case "":
        for fi49 := 0; fi49 < 14; fi49++ {
            v3 -= (v3 - 127)
            var v50 uint32 = 1
            _ = fi49
            _ = v50
        }
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v4 = v4
        }
    default:
        switch ("" + "a") {
        case "\n\t":
            var v51 uint8 = 237
            _ = v51
        }
    }
    if true {
        v3 -= v3
    } else {
        for fi52 := 0; fi52 < 18; fi52++ {
            for fi53 := 0; fi53 < 8; fi53++ {
                var v54 int = (v4 ^ fi53)
                if true {
                    var v55 int16 = 1
                    _ = v55
                } else {
                    var v56 int16 = -1
                    _ = v56
                }
                _ = fi53
                _ = v54
            }
            _ = fi52
        }
    }
    v57 := GenS5{F0: (v4 + (v4 | -1)), F1: v3}
    v58 := GenS5{F0: v4, F1: (^(v3))}
    for fi59 := 0; fi59 < 17; fi59++ {
        var v60 int = (v4 * fi59)
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 int32 = 1
            _ = fi61
            _ = v62
        }
        for fi63 := 0; fi63 < 17; fi63++ {
            v4 = 0
            _ = fi63
        }
        _ = fi59
        _ = v60
    }
    for fi64 := 0; fi64 < 20; fi64++ {
        if ((0.1 + 1.5) < -2.25) {
            v3 += -1
            switch 835 {
            case -494:
                var v65 uint32 = 2
                _ = v65
            case 32768:
                switch ("hello" + "foo bar") {
                case "a":
                    for fi66 := 0; fi66 < 9; fi66++ {
                        var v67 bool = true
                        v3 = -90
                        v4 = (v4 >> ((946) & 63))
                        _ = fi66
                        _ = v67
                    }
                case "foo bar":
                    v68 := GenS5{F0: v4, F1: (v3 % ((v3) | 1))}
                    _ = v68
                case "\n\t":
                    var v69 int64 = 42
                    _ = v69
                default:
                    v70 := GenS5{F0: (v4 % (((fi64 % ((127) | 1))) | 1)), F1: 54}
                    _ = v70
                }
            case 42:
                v4 -= (fi64 % ((fi64) | 1))
            default:
                if (((2147483647 <= 2147483647) || false) || true) {
                    if (147 <= 67) {
                        for fi71 := 0; fi71 < 12; fi71++ {
                            delete(v1, (fi71 ^ 2147483647))
                            v3 += v3
                            var v72 string = "a"
                            _ = fi71
                            _ = v72
                        }
                        if true {
                            for fi73 := 0; fi73 < 17; fi73++ {
                                switch (v4 % (((fi73 ^ fi64)) | 1)) {
                                case 128:
                                    v3 = (v3 ^ v3)
                                case 84:
                                    if true {
                                        var v74 int = -1
                                        var v75 int8 = (-(v3))
                                        _ = v74
                                        _ = v75
                                    }
                                case -950:
                                    switch ("hello" + "") {
                                    case "":
                                        var v76 float32 = (-2.25 / 1e10)
                                        _ = v76
                                    case "foo bar":
                                        var v77 uint16 = 594
                                        _ = v77
                                    }
                                }
                                _ = fi73
                            }
                        }
                    }
                    if false {
                        switch (("\n\t" + "hello") + "hello") {
                        case "foo bar":
                            var v78 int16 = 630
                            _ = v78
                        case "":
                            var v79 uint8 = 255
                            _ = v79
                        case "hello":
                            if (!true) {
                                v3 -= v3
                            } else {
                                v4 += fi64
                                for fi80 := 0; fi80 < 1; fi80++ {
                                    switch ("foo bar" + "hello") {
                                    case "":
                                        v81 := GenS5{F0: -1, F1: v3}
                                        _ = v81
                                    case "foo bar":
                                        switch v4 {
                                        case -523:
                                            if len(v46) > 0 {
                                                v46[0] = (^(fi80))
                                            }
                                        case -263:
                                            var v82 uint = 475
                                            _ = v82
                                        case 2147483647:
                                            if (1 >= 742) {
                                                var v83 int16 = 1
                                                v84 := GenS6{F0: 42}
                                                var v85 uint = 4294967295
                                                if false {
                                                    v86 := GenS7{F0: (18446744073709551615 <= 2), F1: 696}
                                                    var v87 float64 = (0.1 / 1.0)
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    v88 := GenS4{F0: -1, F1: 2147483647, F2: 343}
                                                    var v89 int = -675
                                                    if false {
                                                        delete(v2, "a")
                                                    }
                                                    _ = v88
                                                    _ = v89
                                                }
                                                _ = v83
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                        v4 = (fi64 * 32767)
                                        var v90 float64 = (0.0 - 3.14159)
                                        _ = v90
                                    default:
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            if len(v46) >= 2 {
                                                v46 = v46[1:len(v46)]
                                            }
                                            v91 := GenS3{F0: fi64, F1: -32768}
                                            _ = v91
                                        case "":
                                            var v92 int = fi64
                                            _ = v92
                                        case "a":
                                            switch fi80 {
                                            case 127:
                                                v3 = 0
                                            case 544:
                                                v93 := GenS6{F0: 2147483647}
                                                _ = v93
                                            case 238:
                                                v1[v4] = 4294967295
                                            }
                                            fmt.Println(fi80, v4, fi80)
                                        default:
                                            if (!false) {
                                                var v94 uint8 = 60
                                                fmt.Println(v4, v3, v94)
                                                _ = v94
                                            }
                                            if (!false) {
                                                _ = len(v46)
                                                if false {
                                                    v4 = 129
                                                    var v95 float32 = 0.0
                                                    v46 = append(v46, v4)
                                                    _ = v95
                                                } else {
                                                    fmt.Println(v3)
                                                }
                                                v4 = -2147483648
                                            } else {
                                                if (!(true && ((1 >= 349) || false))) {
                                                    var v96 int8 = v3
                                                    switch v4 {
                                                    case 4294967295:
                                                        v97 := GenS8{F0: 76, F1: 1}
                                                        _ = v97
                                                    default:
                                                        v98 := GenS5{F0: (v4 % ((fi80) | 1)), F1: v3}
                                                        _ = v98
                                                    }
                                                    v99 := GenS5{F0: v4, F1: v96}
                                                    _ = v96
                                                    _ = v99
                                                } else {
                                                    var v100 string = "hello"
                                                    var v101 int16 = 32767
                                                    fmt.Println(v4, v3)
                                                    _ = v100
                                                    _ = v101
                                                }
                                                fmt.Println(v4)
                                                v4 = v4
                                            }
                                            v102 := GenS9{F0: 0}
                                            _ = v102
                                        }
                                    }
                                    var v103 float64 = (0.1 - ((-2.25 / (((1e10 / ((float64(0.1) / 1.5) * (float64(3.14159) / 3.14159))) / ((((((float64(-2.25) + 0.1) - (0.1 - float64((0.0 * ((((3.14159 * (((-2.25 - -2.25) - 3.14159) - (1.5 * 1e10))) + (0.0 - 1e10)) - -2.25) * -2.25))))) - (1.5 * float64((1.5 - ((float32(1e10) - (1e10 + 1e10)) + 0.0))))) / 1.0) * 3.14159) * 1.5)) + float64(-2.25))) + 1e10))
                                    var v104 uint = 2
                                    _ = fi80
                                    _ = v103
                                    _ = v104
                                }
                                if (((0.1 / 1.0) * ((float32(1.5) * 3.14159) - 1e10)) == 0.1) {
                                    if true {
                                        var v105 int = fi64
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v105, v105, v105)
                                        default:
                                            var v106 uint16 = 65535
                                            _ = v106
                                        }
                                        _ = v105
                                    }
                                }
                            }
                        default:
                            v3 -= v3
                        }
                        if (-2147483648 != 0) {
                            v3 = v3
                        } else {
                            switch "hello" {
                            case "":
                                if false {
                                    if len(v46) >= 2 {
                                        v46 = v46[1:len(v46)]
                                    }
                                } else {
                                    v107 := v2[""]
                                    switch "" {
                                    case "a":
                                        _ = len(v46)
                                    default:
                                        if true {
                                            switch "\n\t" {
                                            case "hello":
                                                switch fi64 {
                                                case -672:
                                                    var v108 bool = false
                                                    _ = v108
                                                case 4294967295:
                                                    var v109 bool = true
                                                    _ = v109
                                                }
                                            case "a":
                                                if (42 < 18446744073709551615) {
                                                    if len(v46) > 0 {
                                                        v46[0] = (fi64 / (((v4 / (((fi64 * 9223372036854775807)) | 1))) | 1))
                                                    }
                                                } else {
                                                    var v110 uint64 = 0
                                                    var v111 float32 = -2.25
                                                    fmt.Println(v111, v110)
                                                    _ = v110
                                                    _ = v111
                                                }
                                            default:
                                                var v112 uint = 65535
                                                _ = v112
                                            }
                                            if false {
                                                fmt.Println(v107, v4)
                                                var v113 float32 = 1e10
                                                if true {
                                                    for fi114 := 0; fi114 < 3; fi114++ {
                                                        v115 := v2["a"]
                                                        v3 += v3
                                                        var v116 uint8 = 43
                                                        _ = fi114
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                }
                                                _ = v113
                                            } else {
                                                v4 += fi64
                                            }
                                        }
                                    }
                                    _ = v107
                                }
                            case "foo bar":
                                var v117 uint8 = 0
                                _ = v117
                            case "a":
                                var v118 uint8 = 255
                                _ = v118
                            }
                            var v119 uint8 = 155
                            _ = v119
                        }
                        fmt.Println(v3, v3)
                    } else {
                        var v120 int8 = (-(v3))
                        delete(v1, 415)
                        _ = v120
                    }
                } else {
                    v121 := GenS8{F0: -362, F1: 9223372036854775807}
                    fmt.Println(v4, v4, fi64)
                    _ = v121
                }
            }
            v3 = (v3 ^ -36)
        } else {
            for fi122 := 0; fi122 < 7; fi122++ {
                var v123 uint = genFn0((false || false))
                _ = fi122
                _ = v123
            }
        }
        for fi124 := 0; fi124 < 6; fi124++ {
            v125 := GenS10{F0: genFn0(true), F1: false, F2: fi64}
            _ = fi124
            _ = v125
        }
        v4 += (v4 % ((v4) | 1))
        _ = fi64
    }
    for fi126 := 0; fi126 < 4; fi126++ {
        v4 += (v4 ^ fi126)
        var v127 uint32 = 65535
        var v128 string = "foo bar"
        _ = fi126
        _ = v127
        _ = v128
    }
    var v129 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v129:", v129)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    _ = v1
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v57
    _ = v58
}
