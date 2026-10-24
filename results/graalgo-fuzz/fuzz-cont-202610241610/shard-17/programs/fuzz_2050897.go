package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 uint64; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS7 struct { F0 int64; F1 int32; F2 uint64 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 uint8 }
func genFn0(p0 uint32, p1 int32, p2 uint8) uint16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        p0 -= p0
        var v5 string = ((("foo bar" + "") + "foo bar") + "hello")
        _ = fi4
        _ = v5
    }
    p1 = (p1 ^ (p1 >> ((255) & 63)))
    var v6 uint32 = p0
    switch "" {
    case "\n\t":
        var v7 string = "a"
        _ = v7
    case "":
        v8 := GenS1{F0: 727, F1: 331}
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := []int{}
            var v11 int = (fi9 >> ((4294967295) & 63))
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        var v12 uint8 = p2
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}
func genFn1(p0 int8) int16 {
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: 18446744073709551615, F1: 4294967295}
        var v15 uint = 42
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint16 = genFn0(2, 0, 2)
    var v17 float32 = (0.1 - (0.0 + -2.25))
    v18 := []int{-1, -579, 830, 0}
    v17 += 3.14159
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return 32767
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    var v3 int32 = 42
    for fi19 := 0; fi19 < 16; fi19++ {
        if (v3 > v3) {
            v20 := []int{fi19, fi19, (^(fi19)), fi19}
            for fi21 := 0; fi21 < 20; fi21++ {
                v1 = ("" + ("foo bar" + "\n\t"))
                switch v1 {
                case "\n\t":
                    v22 := GenS1{F0: 1, F1: 2}
                    _ = v22
                case "foo bar":
                    if (!(fi19 < fi19)) {
                        switch v1 {
                        case "a":
                            fmt.Println(fi21, fi21, fi21)
                        case "":
                            var v23 int16 = 213
                            var v24 int16 = v23
                            v23 = v23
                            _ = v23
                            _ = v24
                        case "hello":
                            if true {
                                v25 := GenS1{F0: 0, F1: 179}
                                if true {
                                    v1 = ""
                                }
                                var v26 uint = 701
                                v27 := GenS1{F0: 891, F1: 2}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        default:
                            v28 := GenS2{F0: 1, F1: 42, F2: 255}
                            _ = v28
                        }
                        switch v1 {
                        case "foo bar":
                            v20 = append(v20, fi21)
                        case "a":
                            var v29 uint8 = 128
                            _ = v29
                        case "\n\t":
                            v3 = 1
                        }
                    }
                case "":
                    fmt.Println(v1)
                    if false {
                        switch -9223372036854775808 {
                        case -140:
                            var v30 int32 = v3
                            _ = v30
                        default:
                            if len(v20) > 0 {
                                v20[0] = fi19
                            }
                        }
                        v31 := GenS2{F0: 2, F1: 0, F2: 174}
                        if (!true) {
                            fmt.Println(fi21, fi21)
                            v3 -= v3
                            var v32 int32 = v3
                            switch fi21 {
                            case -353:
                                v32 -= (v32 >> ((503) & 63))
                            case -349:
                                fmt.Println(fi21, fi21)
                            }
                            _ = v32
                        }
                        _ = v31
                    }
                    if (!(553 >= 973)) {
                        fmt.Println(fi21, fi19)
                    }
                    switch (v1 + ("\n\t" + v1)) {
                    case "hello":
                        v3 = 242
                    case "foo bar":
                        v3 += -728
                    default:
                        v1 = "a"
                    }
                    fmt.Println(v1, fi21, v1)
                    switch v1 {
                    case "a":
                        var v33 int32 = (v3 - v3)
                        _ = v33
                    case "\n\t":
                        if (!true) {
                            var v34 uint = 4294967295
                            v3 = 42
                            v3 = (^(v3))
                            v34 -= v34
                            switch (fi21 - -2147483648) {
                            case -1:
                                var v35 uint8 = 127
                                _ = v35
                            }
                            v36 := GenS3{F0: 716, F1: 722, F2: 2147483647}
                            _ = v34
                            _ = v36
                        }
                    case "foo bar":
                        var v37 float64 = (1e10 + float64((((1e10 + 1.5) * 1e10) / 0.1)))
                        _ = v37
                    default:
                        fmt.Println(v1, fi19, v3)
                    }
                default:
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                }
                _ = fi21
            }
            _ = v20
        } else {
            v1 = v1
        }
        _ = fi19
    }
    fmt.Println(v3, v1)
    v3 = (v3 ^ (v3 | v3))
    v3 = (v3 >> ((18446744073709551615) & 63))
    var v38 float32 = (((((((-2.25 - 0.0) - 1.5) / (-2.25 - ((0.0 + float32(-2.25)) + 1e10))) - ((1e10 / (1e10 + 1e10)) / -2.25)) - (1e10 * (1e10 + ((3.14159 + 1.5) * (1e10 / -2.25))))) * 0.0) * float32(3.14159))
    switch -418 {
    case 32767:
        v39 := GenS2{F0: 65535, F1: 0, F2: 0}
        _ = v39
    case 1:
        var v40 string = v1
        _ = v40
    default:
        fmt.Println(v1, v1)
    }
    if (false && (9223372036854775807 <= 2147483647)) {
        if (!true) {
            v38 += (((v38 * v38) - 0.1) - v38)
        } else {
            fmt.Println(v1, v1)
            if (32767 != 1000) {
                switch v1 {
                case "foo bar":
                    if true {
                        v41 := []int{-323, 4294967295, 864, 2147483647}
                        switch 127 {
                        case 256:
                            for fi42 := 0; fi42 < 8; fi42++ {
                                v38 += 0.0
                                var v43 uint8 = 127
                                var v44 uint16 = 727
                                _ = fi42
                                _ = v43
                                _ = v44
                            }
                        }
                        _ = v41
                    } else {
                        for fi45 := 0; fi45 < 16; fi45++ {
                            var v46 int64 = 480
                            _ = fi45
                            _ = v46
                        }
                        var v47 bool = true
                        _ = v47
                    }
                default:
                    v48 := GenS3{F0: 42, F1: 941, F2: 856}
                    _ = v48
                }
            } else {
                if (true && (127 < 0)) {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        switch "hello" {
                        case "foo bar":
                            var v50 int64 = -231
                            _ = v50
                        default:
                            if false {
                                if (false && true) {
                                    fmt.Println(v3, v38, v3)
                                    if (!(!false)) {
                                        fmt.Println(v1)
                                        for fi51 := 0; fi51 < 7; fi51++ {
                                            fmt.Println(v38)
                                            _ = fi51
                                        }
                                        for fi52 := 0; fi52 < 1; fi52++ {
                                            v53 := GenS4{F0: 9}
                                            v38 = v38
                                            fmt.Println(v3, v38)
                                            v54 := []int{(fi52 - 860), (fi52 | (-(fi49)))}
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                }
                            } else {
                                for fi55 := 0; fi55 < 1; fi55++ {
                                    v56 := []int{(fi49 + (fi55 - fi49)), -2147483648}
                                    _ = fi55
                                    _ = v56
                                }
                            }
                        }
                        v38 += 0.0
                        for fi57 := 0; fi57 < 14; fi57++ {
                            var v58 string = "\n\t"
                            _ = fi57
                            _ = v58
                        }
                        _ = fi49
                    }
                    if false {
                        var v59 int16 = genFn1(100)
                        _ = v59
                    }
                    v60 := GenS1{F0: 1, F1: 929}
                    _ = v60
                }
                if (!true) {
                    v61 := []int{860, 128}
                    var v62 int16 = 1000
                    _ = v61
                    _ = v62
                }
            }
            var v63 int32 = (-(v3))
            _ = v63
        }
    } else {
        v1 = "\n\t"
        v64 := []int{383, 0, -773, 9223372036854775807}
        switch 128 {
        case 2147483647:
            var v65 uint32 = 746
            _ = v65
        case -1:
            v2[0] = true
        default:
            if (195 == 2) {
                for fi66 := 0; fi66 < 9; fi66++ {
                    for fi67 := 0; fi67 < 6; fi67++ {
                        if true {
                            v68 := GenS2{F0: 1, F1: fi66, F2: 792}
                            fmt.Println(v3)
                            switch (v1 + (v1 + v1)) {
                            case "foo bar":
                                v69 := GenS5{F0: -760, F1: 658, F2: (v3 >> ((0) & 63))}
                                _ = v69
                            case "a":
                                v70 := GenS1{F0: 4294967295, F1: 4294967295}
                                _ = v70
                            case "":
                                if false {
                                    var v71 int = int(v3)
                                    v72 := GenS6{F0: 68, F1: -9223372036854775808, F2: 657}
                                    _ = v71
                                    _ = v72
                                }
                            }
                            _ = v68
                        } else {
                            v1 = ""
                            switch (v1 + (v1 + "a")) {
                            case "":
                                switch fi67 {
                                case 127:
                                    v3 += v3
                                case 0:
                                    switch 127 {
                                    case -1:
                                        v38 += v38
                                    case 770:
                                        if (2147483647 < 2147483647) {
                                            v73 := GenS1{F0: 42, F1: 0}
                                            _ = v73
                                        } else {
                                            if true {
                                                v74 := v2[(^(fi67))]
                                                v38 += -2.25
                                                _ = v74
                                            } else {
                                                v3 -= (v3 - v3)
                                            }
                                        }
                                    case 32768:
                                        v38 -= float32(0.1)
                                    }
                                case -2147483648:
                                    fmt.Println(v1, fi66)
                                default:
                                    v38 += (1e10 + 1.5)
                                }
                            }
                        }
                        var v75 uint = 255
                        switch "" {
                        case "a":
                            v76 := GenS1{F0: 42, F1: 18446744073709551615}
                            _ = v76
                        default:
                            if false {
                                switch (fi66 ^ fi67) {
                                case -565:
                                    var v77 float64 = (1e10 * 0.0)
                                    _ = v77
                                case -767:
                                    var v78 uint32 = 2
                                    _ = v78
                                case 32767:
                                    var v79 float32 = ((v38 / ((((v38 / v38) / ((0.1 - -2.25) - v38)) - float32(0.1)) - (-2.25 * ((v38 / (((float32((0.0 + 0.1)) + (0.1 - v38)) * v38) + -2.25)) * (v38 + 0.1))))) - ((v38 / v38) + float32(1e10)))
                                    _ = v79
                                default:
                                    if len(v64) >= 2 {
                                        v64 = v64[1:len(v64)]
                                    }
                                }
                                if (int64(v75) >= -138) {
                                    switch (v1 + "\n\t") {
                                    case "foo bar":
                                        v80 := GenS2{F0: 0, F1: 127, F2: 255}
                                        _ = v80
                                    case "a":
                                        var v81 string = v1
                                        _ = v81
                                    case "\n\t":
                                        var v82 uint8 = 127
                                        _ = v82
                                    }
                                } else {
                                    if true {
                                        v38 = ((3.14159 - 1.5) - ((((0.0 + v38) * v38) / 1e10) + 3.14159))
                                        if false {
                                            if (-9223372036854775808 == (fi66 & (-(fi67)))) {
                                                var v83 float32 = v38
                                                _ = v83
                                            } else {
                                                var v84 uint8 = 110
                                                var v85 uint32 = 2147483647
                                                _ = v84
                                                _ = v85
                                            }
                                            var v86 uint64 = 2
                                            v86 += v86
                                            _ = v86
                                        }
                                        switch (fi67 * (fi66 | (-(fi67)))) {
                                        case 9223372036854775807:
                                            var v87 uint = v75
                                            _ = v87
                                        case 898:
                                            delete(v2, (fi67 * -992))
                                        default:
                                            var v88 uint32 = 687
                                            _ = v88
                                        }
                                    }
                                    v38 = 1e10
                                    fmt.Println(v1)
                                }
                                var v89 float64 = (3.14159 + (1e10 / 1.0))
                                _ = v89
                            }
                        }
                        fmt.Println(v75, fi67)
                        v38 -= (0.1 * float32((0.1 * 1e10)))
                        var v90 uint8 = uint8(v75)
                        if (v3 < (v3 & (v3 << ((v75) & 63)))) {
                            var v91 float32 = ((0.0 / float32(float64(((v38 * v38) + 1.5)))) - v38)
                            var v92 int = fi67
                            v93 := GenS6{F0: 0, F1: -1, F2: 42}
                            _ = v91
                            _ = v92
                            _ = v93
                        } else {
                            if (!false) {
                                v94 := GenS3{F0: 920, F1: 4294967295, F2: fi67}
                                var v95 int8 = 1
                                _ = v94
                                _ = v95
                            } else {
                                var v96 uint = v75
                                if (1 <= 9223372036854775807) {
                                    v97 := GenS7{F0: -1, F1: 1, F2: 1}
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v98 := GenS2{F0: 65535, F1: fi66, F2: v75}
                                        _ = v98
                                    case "\n\t":
                                        v64 = append(v64, int(v3))
                                    }
                                    switch "a" {
                                    case "a":
                                        v38 = (v38 * v38)
                                    case "hello":
                                        if (!true) {
                                            v75 = v96
                                            var v99 int16 = 1000
                                            _ = v99
                                        } else {
                                            v100 := GenS8{F0: 2147483647, F1: (v96 + v75)}
                                            _ = v100
                                        }
                                    default:
                                        if len(v64) >= 2 {
                                            v64 = v64[1:len(v64)]
                                        }
                                    }
                                    _ = v97
                                }
                                _ = v96
                            }
                            var v101 int8 = 100
                            if len(v64) >= 2 {
                                v64 = v64[1:len(v64)]
                            }
                            fmt.Println(fi66, v75, fi67)
                            switch "foo bar" {
                            case "hello":
                                if (-2.25 <= 1e10) {
                                    v38 += (1e10 / ((3.14159 - ((v38 - v38) + (float32(-2.25) - (v38 - (v38 / v38))))) / (3.14159 * v38)))
                                    fmt.Println(fi66)
                                    var v102 uint8 = v90
                                    fmt.Println(v75)
                                    _ = v102
                                } else {
                                    var v103 int8 = (v101 & -1)
                                    if true {
                                        var v104 uint = (v75 >> ((0) & 63))
                                        _ = v104
                                    } else {
                                        var v105 uint16 = 32767
                                        _ = v105
                                    }
                                    var v106 uint32 = 42
                                    _ = v103
                                    _ = v106
                                }
                            case "\n\t":
                                v90 = 0
                            default:
                                if false {
                                    v90 += (v90 / ((v90) | 1))
                                    v90 -= 127
                                } else {
                                    v75 = v75
                                }
                            }
                            _ = v101
                        }
                        _ = fi67
                        _ = v75
                        _ = v90
                    }
                    if false {
                        v38 += (((v38 * ((float32(3.14159) + 1.5) / (v38 - v38))) / (v38 - 1e10)) / v38)
                        if (v1 == v1) {
                            if true {
                                var v107 uint8 = 42
                                v38 -= (0.0 + -2.25)
                                _ = v107
                            }
                            v108 := GenS7{F0: -9223372036854775808, F1: 2147483647, F2: 393}
                            switch (fi66 - fi66) {
                            case 32767:
                                delete(v2, fi66)
                            default:
                                switch (fi66 - (fi66 ^ (fi66 ^ fi66))) {
                                case 2147483647:
                                    var v109 int8 = 0
                                    _ = v109
                                case -2147483648:
                                    if len(v64) >= 2 {
                                        v64 = v64[1:len(v64)]
                                    }
                                case -175:
                                    var v110 string = ("hello" + v1)
                                    _ = v110
                                default:
                                    var v111 int8 = -124
                                    _ = v111
                                }
                            }
                            _ = v108
                        }
                    } else {
                        var v112 float64 = 3.14159
                        var v113 float64 = (v112 + float64(1e10))
                        var v114 uint16 = 1
                        _ = v112
                        _ = v113
                        _ = v114
                    }
                    v115 := GenS9{F0: 0}
                    _ = fi66
                    _ = v115
                }
                fmt.Println(v3, v38, v3)
            }
        }
        _ = v64
    }
    v3 = 2147483647
    if false {
        var v116 float64 = ((((float64(v38) / 0.1) - float64(float32(-2.25))) - 3.14159) * 3.14159)
        v116 = ((0.0 - 1e10) - v116)
        for fi117 := 0; fi117 < 19; fi117++ {
            v116 = v116
            _ = fi117
        }
        _ = v116
    }
    if (false && (2 != 65535)) {
        v2[-586] = (127 != 0)
        fmt.Println(v38, v1, v1)
    }
    v1 = (v1 + (v1 + v1))
    if ((true || true) && (65535 >= 761)) {
        var v118 int32 = 42
        _ = v118
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
}
