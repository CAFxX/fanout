package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool; F1 uint8; F2 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 bool; F1 int32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint16; F1 uint; F2 uint }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int32; F1 uint16 }
type GenS16 struct { F0 int32 }
func genFn0(p0 uint8, p1 int16) uint {
    for fi6 := 0; fi6 < 6; fi6++ {
        if ((-128 > 100) || (42 < -19)) {
            v7 := GenS1{F0: 32767, F1: (-128 <= 0), F2: 4294967295}
            _ = v7
        }
        v8 := []int{}
        v9 := GenS2{F0: -2147483648, F1: 0}
        _ = fi6
        _ = v8
        _ = v9
    }
    p0 = (^(p0))
    p0 = (p0 + 128)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    var v2 bool = false
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 255
    if (!(-9223372036854775808 <= 737)) {
        switch 32768 {
        case 42:
            if ((-706 < 1) || (117 == 2)) {
                switch ("foo bar" + "foo bar") {
                case "a":
                    for fi10 := 0; fi10 < 1; fi10++ {
                        switch fi10 {
                        case 761:
                            v11 := []int{859, fi10, 772}
                            _ = v11
                        }
                        fmt.Println(v5)
                        var v12 bool = v2
                        _ = fi10
                        _ = v12
                    }
                case "hello":
                    switch 128 {
                    case 32767:
                        v13 := GenS2{F0: 968, F1: 42}
                        _ = v13
                    case 0:
                        v1 = v1
                    }
                default:
                    for fi14 := 0; fi14 < 11; fi14++ {
                        for fi15 := 0; fi15 < 17; fi15++ {
                            if (!v2) {
                                var v16 string = ("foo bar" + ((("" + "hello") + "\n\t") + "foo bar"))
                                v17 := GenS2{F0: -2147483648, F1: 2147483648}
                                v1 -= (v5 ^ 0)
                                _ = v16
                                _ = v17
                            } else {
                                if v2 {
                                    v4[42] = true
                                    v18 := v4[122]
                                    var v19 int64 = 42
                                    _ = v18
                                    _ = v19
                                }
                                fmt.Println(v2)
                                v4[65535] = ((v2 || v2) || ((fi14 % (((fi15 - 42)) | 1)) < fi14))
                                if v2 {
                                    v20 := GenS3{F0: (^(v5))}
                                    var v21 float32 = -2.25
                                    _ = v20
                                    _ = v21
                                }
                            }
                            if v2 {
                                v2 = ((((0.1 / float64(1.5)) == 1.5) || (1 != 361)) && true)
                                if ("hello" != "") {
                                    v2 = v2
                                    v22 := v3[516]
                                    _ = v22
                                }
                                if (false && v2) {
                                    var v23 uint = 65535
                                    v23 = v23
                                    _ = v23
                                } else {
                                    var v24 float64 = 0.0
                                    _ = v24
                                }
                                var v25 bool = v2
                                _ = v25
                            }
                            _ = fi15
                        }
                        _ = fi14
                    }
                }
                var v26 int32 = 782
                _ = v26
            } else {
                var v27 uint64 = 187
                var v28 string = "foo bar"
                v2 = ((0.0 * 1e10) == 0.1)
                _ = v27
                _ = v28
            }
        case 780:
            if (-18 == 127) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    delete(v4, 42)
                case "hello":
                    v5 -= v5
                }
                v29 := GenS4{F0: "hello"}
                _ = v29
            } else {
                v30 := []int{-688, -2147483648, 0}
                switch 9223372036854775807 {
                case -48:
                    var v31 int = -2147483648
                    _ = v31
                default:
                    fmt.Println(v2, v2, v5)
                }
                _ = v30
            }
        case 0:
            v32 := []int{}
            _ = v32
        }
    } else {
        for fi33 := 0; fi33 < 1; fi33++ {
            fmt.Println(fi33, v2)
            var v34 uint64 = 42
            v35 := []int{}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    if (!(0 >= -816)) {
        if (float32(((0.0 * 0.0) / 1e10)) < (3.14159 - 1e10)) {
            v36 := GenS5{F0: 274, F1: -2147483648}
            for fi37 := 0; fi37 < 2; fi37++ {
                for fi38 := 0; fi38 < 13; fi38++ {
                    v2 = (1000 >= 660)
                    v39 := GenS4{F0: "foo bar"}
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
            _ = v36
        }
    }
    var v40 int8 = 85
    v1 += v5
    if v2 {
        v41 := GenS2{F0: 42, F1: 204}
        v42 := GenS6{F0: 42, F1: 42}
        v2 = v2
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 6; fi43++ {
        var v44 int64 = 684
        _ = fi43
        _ = v44
    }
    v45 := []int{32768, 42, 32768, 127}
    if len(v45) >= 2 {
        v45 = v45[1:len(v45)]
    }
    v46 := GenS6{F0: -1, F1: -676}
    var v47 uint64 = 2
    if v2 {
        var v48 uint32 = 65535
        if (v2 && true) {
            switch -954 {
            case 275:
                if v2 {
                    if (!(v2 || (false && v2))) {
                        if (255 <= 1) {
                            fmt.Println(v47, v1)
                            if (!v2) {
                                v47 += v47
                                var v49 int32 = 440
                                v50 := GenS7{F0: 42}
                                _ = v49
                                _ = v50
                            } else {
                                var v51 string = ""
                                if (!false) {
                                    var v52 uint = 42
                                    v53 := GenS8{F0: v2, F1: 212, F2: -32768}
                                    var v54 string = ("foo bar" + (v51 + v51))
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                } else {
                                    v55 := v3[v48]
                                    fmt.Println(v51, v40, v47)
                                    var v56 int = -1
                                    _ = v55
                                    _ = v56
                                }
                                var v57 uint8 = 128
                                _ = v51
                                _ = v57
                            }
                        }
                        switch 127 {
                        case 1:
                            _ = len(v45)
                        case 32768:
                            switch "a" {
                            case "hello":
                                _ = len(v3)
                            case "":
                                var v58 uint32 = v48
                                _ = v58
                            case "a":
                                v4[(v48 & v48)] = (!(!(65535 > 504)))
                            }
                        default:
                            v40 = (v40 | v40)
                        }
                    }
                    switch 2147483647 {
                    case 4294967295:
                        var v59 int = -1
                        _ = v59
                    }
                    for fi60 := 0; fi60 < 6; fi60++ {
                        for fi61 := 0; fi61 < 3; fi61++ {
                            if v2 {
                                fmt.Println(fi61)
                            } else {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                if v2 {
                                    fmt.Println(v48, v40, v2)
                                    var v62 uint = genFn0(1, int16(v47))
                                    _ = v62
                                } else {
                                    switch fi61 {
                                    case 9223372036854775807:
                                        var v63 uint8 = 127
                                        _ = v63
                                    case 32767:
                                        _ = len(v3)
                                    default:
                                        v64 := GenS9{F0: 0}
                                        _ = v64
                                    }
                                    v4[(^(v48))] = v2
                                }
                                v65 := GenS7{F0: -411}
                                if (true && (v40 > (v40 * v40))) {
                                    v40 = v40
                                } else {
                                    v48 += uint32(v47)
                                }
                                _ = v65
                            }
                            v66 := GenS10{F0: int16(v48), F1: 229}
                            _ = fi61
                            _ = v66
                        }
                        _ = fi60
                    }
                }
            }
            for fi67 := 0; fi67 < 14; fi67++ {
                if ((false && v2) || (false || (v2 && (!((v5 * v5) <= v1))))) {
                    for fi68 := 0; fi68 < 1; fi68++ {
                        if (42 > (^(v48))) {
                            v69 := GenS11{F0: false, F1: -2147483648}
                            v70 := GenS12{F0: 42}
                            var v71 int64 = -979
                            v5 = 255
                            if (v2 || v2) {
                                v47 = (v47 * 470)
                            } else {
                                var v72 bool = ("\n\t" < "foo bar")
                                _ = v72
                            }
                            _ = v69
                            _ = v70
                            _ = v71
                        } else {
                            v48 = v48
                            switch (fi67 << ((1) & 63)) {
                            case -448:
                                fmt.Println(fi67, v48, fi68)
                                var v73 float64 = 0.1
                                _ = v73
                            case 939:
                                var v74 float64 = 1e10
                                _ = v74
                            case -702:
                                if (-2147483648 != 586) {
                                    switch fi67 {
                                    case 812:
                                        var v75 int64 = 9223372036854775807
                                        if ((-2.25 * (3.14159 + 3.14159)) == (0.0 * float64((-2.25 / 3.14159)))) {
                                            if len(v45) >= 2 {
                                                v45 = v45[1:len(v45)]
                                            }
                                            if (1.5 <= 0.1) {
                                                v40 -= (v40 ^ 1)
                                                var v76 float64 = (3.14159 + -2.25)
                                                _ = v76
                                            } else {
                                                var v77 uint = 255
                                                v48 -= (v48 ^ 535)
                                                switch "hello" {
                                                case "foo bar":
                                                    var v78 uint8 = 140
                                                    _ = v78
                                                case "hello":
                                                    v48 = v48
                                                }
                                                if ((float64(0.0) > (((((float64(-2.25) - -2.25) + (-2.25 * (0.1 - -2.25))) - ((1.5 - -2.25) + float64(0.1))) * 0.1) + -2.25)) && false) {
                                                    var v79 int = (fi68 / ((fi68) | 1))
                                                    _ = v79
                                                } else {
                                                    switch fi68 {
                                                    case 2147483647:
                                                        v4[(v48 | (v48 ^ 969))] = true
                                                    default:
                                                        v5 += v5
                                                    }
                                                    switch 32768 {
                                                    case 400:
                                                        v80 := GenS2{F0: -46, F1: v48}
                                                        _ = v80
                                                    case -930:
                                                        v2 = false
                                                    case 42:
                                                        fmt.Println(v1, v48)
                                                        v81 := GenS7{F0: (v75 >> ((v77) & 63))}
                                                        _ = v81
                                                    }
                                                    var v82 int8 = (-(v40))
                                                    _ = v82
                                                }
                                                _ = v77
                                            }
                                            fmt.Println(fi68, v47)
                                        } else {
                                            var v83 uint64 = 42
                                            v75 += v75
                                            if v2 {
                                                v84 := GenS5{F0: genFn0(16, -1), F1: (fi67 + fi68)}
                                                delete(v3, v48)
                                                switch ("\n\t" + "hello") {
                                                case "a":
                                                    v4[(^(v48))] = true
                                                case "\n\t":
                                                    if false {
                                                        if (!v2) {
                                                            v5 = (v1 & (v5 & (v5 + v1)))
                                                        }
                                                        var v85 float32 = (1.5 * 1.5)
                                                        _ = v85
                                                    }
                                                default:
                                                    v2 = (!((343 == 42) || v2))
                                                }
                                                _ = v84
                                            }
                                            fmt.Println(fi67, v83, v75)
                                            v75 -= 9223372036854775807
                                            _ = v83
                                        }
                                        v48 += v48
                                        _ = v75
                                    case -395:
                                        switch ("foo bar" + "\n\t") {
                                        case "a":
                                            v86 := GenS13{F0: v5, F1: 2, F2: 2}
                                            _ = v86
                                        case "\n\t":
                                            var v87 bool = ("a" <= ("a" + ("hello" + "hello")))
                                            _ = v87
                                        case "hello":
                                            var v88 uint64 = v47
                                            _ = v88
                                        }
                                    case 127:
                                        v89 := GenS14{F0: -2147483648}
                                        _ = v89
                                    default:
                                        if (true || v2) {
                                            var v90 uint32 = (v48 << ((0) & 63))
                                            v47 = (v47 % (((v47 >> ((1000) & 63))) | 1))
                                            var v91 bool = (2 > 0)
                                            _ = v90
                                            _ = v91
                                        } else {
                                            v48 = v48
                                        }
                                    }
                                }
                            }
                        }
                        if (v2 && (0.1 != 0.1)) {
                            var v92 int = (fi68 * (fi67 & 592))
                            var v93 uint8 = 2
                            _ = v92
                            _ = v93
                        }
                        _ = fi68
                    }
                    if v2 {
                        var v94 uint = uint(v47)
                        switch "hello" {
                        case "a":
                            if ((9223372036854775807 == -731) || (!v2)) {
                                v2 = true
                            } else {
                                v1 -= 605
                                fmt.Println(v1)
                            }
                        }
                        _ = v94
                    } else {
                        var v95 float64 = 0.0
                        var v96 int32 = 2147483647
                        if ((65535 <= 589) || (uint8(v48) != 0)) {
                            var v97 float64 = -2.25
                            for fi98 := 0; fi98 < 8; fi98++ {
                                if v2 {
                                    if (v2 && (9223372036854775807 < -9223372036854775808)) {
                                        v95 = 0.1
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v96 = (^(v96))
                                    default:
                                        if (uint64(v5) != v47) {
                                            var v99 uint8 = 179
                                            _ = v99
                                        }
                                    }
                                    if false {
                                        fmt.Println(v48)
                                        _ = len(v45)
                                        if v2 {
                                            fmt.Println(v47, v47)
                                        } else {
                                            switch "hello" {
                                            case "foo bar":
                                                v47 += (^(v47))
                                            default:
                                                v48 = (v48 + v48)
                                            }
                                            fmt.Println(v96, v95)
                                            fmt.Println(v47, v2, v40)
                                        }
                                        var v100 int8 = v40
                                        _ = v100
                                    }
                                    if (uint32(v1) > (^(v48))) {
                                        var v101 bool = true
                                        v3[uint32(v5)] = true
                                        var v102 uint32 = v48
                                        _ = v101
                                        _ = v102
                                    }
                                    v1 = v5
                                } else {
                                    var v103 int32 = v96
                                    _ = v103
                                }
                                if len(v45) > 0 {
                                    v45[0] = (fi98 << ((292) & 63))
                                }
                                _ = len(v3)
                                _ = fi98
                            }
                            _ = len(v45)
                            _ = v97
                        }
                        _ = v95
                        _ = v96
                    }
                }
                _ = len(v45)
                v104 := GenS15{F0: 0, F1: (v1 / ((v5) | 1))}
                _ = fi67
                _ = v104
            }
            if (!(9223372036854775807 > 635)) {
                switch 62 {
                case -2147483648:
                    delete(v3, v48)
                case 128:
                    v105 := v3[uint32(v47)]
                    _ = v105
                case 9223372036854775807:
                    if (0 > -1) {
                        _ = len(v4)
                        switch 510 {
                        case 4294967295:
                            if v2 {
                                v106 := GenS6{F0: 151, F1: -1}
                                fmt.Println(v48)
                                _ = v106
                            } else {
                                _ = len(v45)
                                if true {
                                    v1 = v1
                                } else {
                                    var v107 int32 = 396
                                    if v2 {
                                        for fi108 := 0; fi108 < 6; fi108++ {
                                            v5 += v5
                                            var v109 string = "foo bar"
                                            var v110 string = (v109 + ((v109 + ((("foo bar" + v109) + "a") + v109)) + v109))
                                            _ = fi108
                                            _ = v109
                                            _ = v110
                                        }
                                        for fi111 := 0; fi111 < 2; fi111++ {
                                            v112 := GenS16{F0: v107}
                                            _ = fi111
                                            _ = v112
                                        }
                                    }
                                    for fi113 := 0; fi113 < 14; fi113++ {
                                        delete(v3, (v48 % ((v48) | 1)))
                                        _ = fi113
                                    }
                                    _ = v107
                                }
                            }
                        case -9223372036854775808:
                            var v114 string = "foo bar"
                            _ = v114
                        }
                    } else {
                        if (v2 || v2) {
                            v48 -= v48
                            for fi115 := 0; fi115 < 3; fi115++ {
                                if true {
                                    v1 -= 2
                                }
                                _ = fi115
                            }
                        } else {
                            if true {
                                v116 := GenS2{F0: 479, F1: 663}
                                var v117 uint = 192
                                fmt.Println(v5, v1, v5)
                                _ = v116
                                _ = v117
                            } else {
                                var v118 int = 451
                                v118 += (v118 | (^(v118)))
                                if (65535 >= 2) {
                                    v118 += v118
                                    v47 = (v47 & 18446744073709551615)
                                }
                                _ = v118
                            }
                            v2 = true
                        }
                    }
                default:
                    if (762 >= 42) {
                        if len(v45) > 0 {
                            v45[0] = 9223372036854775807
                        }
                        fmt.Println(v48, v40, v47)
                    }
                }
                for fi119 := 0; fi119 < 9; fi119++ {
                    v2 = (float64(float32(-2.25)) >= 1e10)
                    _ = fi119
                }
            }
        }
        if false {
            v120 := GenS9{F0: -1}
            v40 = v40
            _ = v120
        }
        _ = v48
    } else {
        if (-2.25 != 1e10) {
            v1 = 32767
            var v121 float64 = 3.14159
            if true {
                var v122 uint32 = 840
                _ = v122
            } else {
                if (127 > 0) {
                    var v123 string = ""
                    _ = v123
                } else {
                    fmt.Println(v40, v47, v47)
                }
                v47 -= (^(v47))
                if len(v45) > 0 {
                    v45[0] = 4294967295
                }
            }
            _ = v121
        }
        switch "\n\t" {
        case "hello":
            v124 := GenS11{F0: ("foo bar" == ""), F1: -2147483648}
            _ = v124
        case "foo bar":
            v1 = 1
        case "":
            for fi125 := 0; fi125 < 17; fi125++ {
                var v126 bool = (1 >= -709)
                v127 := GenS7{F0: int64(v47)}
                switch (fi125 - (fi125 - fi125)) {
                case -2147483648:
                    var v128 int = (-(fi125))
                    _ = v128
                }
                _ = fi125
                _ = v126
                _ = v127
            }
        }
        v129 := v3[2147483648]
        _ = v129
    }
    var v130 float64 = 1e10
    var v131 bool = (v2 && false)
    var v132 int16 = 1000
    for fi133 := 0; fi133 < 17; fi133++ {
        if true {
            v5 -= v1
        }
        var v134 uint8 = 1
        _ = fi133
        _ = v134
    }
    v1 = (v1 & v1)
    var v135 uint = 881
    var v136 int = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v130
    _ = v131
    _ = v132
    _ = v135
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
