package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int64; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 uint16; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint8; F2 uint }
type GenS10 struct { F0 bool; F1 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int32 }

func main() {
    var v1 int16 = 0
    var v2 float64 = (0.1 - float64(float32(3.14159)))
    var v3 int64 = 9223372036854775807
    var v4 int32 = 0
    var v5 float32 = ((-2.25 * ((0.0 / -2.25) - 1e10)) + 0.1)
    v6 := make(map[int]uint32)
    switch 127 {
    case 32768:
        fmt.Println(v2, v1)
    default:
        var v7 uint16 = 42
        _ = v7
    }
    v8 := []int{-1}
    var v9 uint64 = 0
    var v10 string = ("a" + "foo bar")
    v3 -= v3
    var v11 int32 = (v4 << ((255) & 63))
    if (!false) {
        var v12 int = -9223372036854775808
        var v13 string = v10
        fmt.Println(v4)
        _ = v12
        _ = v13
    }
    v8 = append(v8, 9223372036854775807)
    var v14 uint8 = 255
    v15 := GenS1{F0: -95}
    for fi16 := 0; fi16 < 6; fi16++ {
        _ = len(v6)
        _ = fi16
    }
    v17 := GenS2{F0: 217, F1: v3, F2: (!false)}
    v18 := GenS2{F0: 0, F1: (v3 ^ v3), F2: false}
    switch 4294967295 {
    case 42:
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v9, v4, v10)
            fmt.Println(v2)
            if len(v8) > 0 {
                v8[0] = -1
            }
            _ = fi19
        }
    case 274:
        v5 = 0.1
    case -2147483648:
        var v20 uint32 = 2147483647
        _ = v20
    default:
        v9 -= v9
    }
    var v21 int8 = 71
    for fi22 := 0; fi22 < 11; fi22++ {
        fmt.Println(v10, v14, v21)
        v23 := GenS2{F0: 610, F1: (v3 * 0), F2: (!((255 <= 0) || false))}
        _ = fi22
        _ = v23
    }
    switch ("foo bar" + (v10 + v10)) {
    case "":
        if (v21 == v21) {
            fmt.Println(v10, v3, v4)
            var v24 uint32 = 42
            _ = v24
        }
    }
    if true {
        _ = len(v6)
        for fi25 := 0; fi25 < 6; fi25++ {
            if (1 >= 0) {
                if true {
                    switch "" {
                    case "hello":
                        switch (v10 + ((("hello" + v10) + v10) + ("hello" + ""))) {
                        case "\n\t":
                            fmt.Println(v3, v10, v3)
                        default:
                            if false {
                                var v26 uint32 = 2
                                v11 -= -135
                                _ = v26
                            }
                        }
                    case "":
                        if false {
                            v14 -= v14
                        }
                    }
                    if (false || ((76 > 1) || (v3 != (^(v3))))) {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            var v28 string = ((v10 + "") + "foo bar")
                            var v29 uint64 = 18446744073709551615
                            var v30 uint16 = 1
                            _ = fi27
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        v11 -= (v4 + -549)
                    } else {
                        delete(v6, -1)
                        v5 = ((1e10 + 1e10) / v5)
                    }
                    var v31 uint32 = 0
                    _ = v31
                } else {
                    var v32 float64 = v2
                    for fi33 := 0; fi33 < 9; fi33++ {
                        switch (fi25 % ((fi25) | 1)) {
                        case 4294967295:
                            if (!(436 > 42)) {
                                if false {
                                    v9 = (v9 >> ((658) & 63))
                                    var v34 int16 = 1
                                    v14 = 42
                                    _ = v34
                                } else {
                                    v5 = v5
                                    switch (fi33 >> ((18446744073709551615) & 63)) {
                                    case 32768:
                                        delete(v6, (fi25 * fi25))
                                    case -2147483648:
                                        var v35 string = "\n\t"
                                        _ = v35
                                    case -1:
                                        if len(v8) > 0 {
                                            v8[0] = 9223372036854775807
                                        }
                                    }
                                    var v36 int32 = (v11 ^ v4)
                                    if (true || (1e10 <= float32((v2 + v32)))) {
                                        v10 = (v10 + v10)
                                        v37 := GenS1{F0: v21}
                                        _ = v37
                                    }
                                    var v38 int16 = (v1 % ((v1) | 1))
                                    v32 -= ((1.5 - v2) + v32)
                                    _ = v36
                                    _ = v38
                                }
                                _ = len(v6)
                                var v39 uint16 = 1
                                if false {
                                    if (!(42 != 18446744073709551615)) {
                                        var v40 uint8 = v14
                                        var v41 uint32 = 638
                                        _ = v40
                                        _ = v41
                                    }
                                } else {
                                    v42 := GenS2{F0: 42, F1: 9223372036854775807, F2: ((v10 + v10) != "")}
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                    switch ((v10 + v10) + ("\n\t" + "\n\t")) {
                                    case "foo bar":
                                        var v43 uint = 255
                                        _ = v43
                                    case "a":
                                        var v44 uint = uint(v39)
                                        switch "" {
                                        case "":
                                            v32 = 0.0
                                        case "\n\t":
                                            fmt.Println(v5)
                                            if (fi33 != -68) {
                                                if true {
                                                    if true {
                                                        v6[fi33] = 4294967295
                                                    } else {
                                                        v4 -= (v4 / (((v4 * int32(v14))) | 1))
                                                        var v45 uint32 = 2147483647
                                                        _ = v45
                                                    }
                                                    var v46 int16 = v1
                                                    _ = v46
                                                } else {
                                                    v9 = 0
                                                    if (!((!false) || true)) {
                                                        _ = len(v8)
                                                    } else {
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            var v47 uint = v44
                                                            _ = v47
                                                        default:
                                                            v48 := GenS3{F0: "\n\t"}
                                                            _ = v48
                                                        }
                                                        if ((v21 + (v21 + int8(v39))) <= v21) {
                                                            if len(v8) > 0 {
                                                                v8[0] = fi33
                                                            }
                                                        } else {
                                                            v1 -= (v1 - v1)
                                                        }
                                                        v3 += -216
                                                    }
                                                    if (false && (false && true)) {
                                                        if true {
                                                            var v49 uint32 = 501
                                                            if true {
                                                                var v50 float64 = (-2.25 - float64((1.5 - 0.1)))
                                                                switch ((v10 + (v10 + ("\n\t" + v10))) + (v10 + ("foo bar" + (v10 + "")))) {
                                                                case "\n\t":
                                                                    var v51 int64 = 42
                                                                    _ = v51
                                                                case "hello":
                                                                    fmt.Println(v32, v49, v1)
                                                                case "":
                                                                    v44 = (v44 % ((2) | 1))
                                                                }
                                                                _ = v50
                                                            }
                                                            switch "foo bar" {
                                                            case "a":
                                                                v1 -= v1
                                                            }
                                                            var v52 int16 = v1
                                                            _ = v49
                                                            _ = v52
                                                        }
                                                        var v53 uint8 = 128
                                                        v32 += 1.5
                                                        switch (v10 + "foo bar") {
                                                        case "a":
                                                            v14 += (v14 << ((v44) & 63))
                                                        case "":
                                                            if true {
                                                                v54 := GenS3{F0: v10}
                                                                v4 = v4
                                                                v44 += uint(v39)
                                                                _ = v54
                                                            }
                                                        case "foo bar":
                                                            v10 = v10
                                                            if len(v8) >= 2 {
                                                                v8 = v8[1:len(v8)]
                                                            }
                                                            if true {
                                                                var v55 int64 = v3
                                                                var v56 float32 = v5
                                                                if (v44 >= 1) {
                                                                    var v57 int = (fi33 + fi33)
                                                                    var v58 uint8 = v53
                                                                    fmt.Println(fi25)
                                                                    v59 := GenS4{F0: (v10 + v10)}
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                v32 -= v2
                                                                v10 = v10
                                                            }
                                                        default:
                                                            v8 = append(v8, -832)
                                                        }
                                                        _ = v53
                                                    }
                                                }
                                                switch ("" + v10) {
                                                case "":
                                                    if true {
                                                        switch (fi33 | (^(fi33))) {
                                                        case 319:
                                                            v39 = (^(v39))
                                                        case 1:
                                                            fmt.Println(v3, fi25)
                                                        case 2147483647:
                                                            switch (("" + v10) + (((v10 + "hello") + v10) + (((v10 + v10) + "a") + v10))) {
                                                            case "":
                                                                if len(v8) > 0 {
                                                                    v8[0] = fi33
                                                                }
                                                            default:
                                                                switch "hello" {
                                                                case "":
                                                                    v2 += (3.14159 * v32)
                                                                default:
                                                                    var v60 int8 = (v21 * v21)
                                                                    _ = v60
                                                                }
                                                                if (false && ((float64(float32(v2)) / v2) >= v32)) {
                                                                    if (!((v3 | v3) != v3)) {
                                                                        _ = len(v6)
                                                                        v5 -= v5
                                                                        var v61 float64 = v32
                                                                        _ = v61
                                                                    }
                                                                    v9 = v9
                                                                    v39 = (^(v39))
                                                                    switch v10 {
                                                                    case "hello":
                                                                        v44 += v44
                                                                        v4 += 42
                                                                    }
                                                                    var v62 uint32 = 4294967295
                                                                    _ = v62
                                                                } else {
                                                                    switch (fi25 | int(v21)) {
                                                                    case -390:
                                                                        _ = len(v6)
                                                                    case -922:
                                                                        v63 := v6[(fi33 + fi25)]
                                                                        _ = v63
                                                                    }
                                                                    v6[fi25] = 4294967295
                                                                }
                                                                v11 -= v11
                                                            }
                                                        }
                                                        v32 -= (v32 / float64(v5))
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                    } else {
                                                        switch v10 {
                                                        case "foo bar":
                                                            var v64 uint = v44
                                                            _ = v64
                                                        case "\n\t":
                                                            var v65 bool = (144 > (v9 ^ v9))
                                                            _ = v65
                                                        case "a":
                                                            switch (fi33 - (fi33 - (fi25 * fi25))) {
                                                            case -108:
                                                                if ((v11 + 2147483647) >= v11) {
                                                                    v14 += (^(v14))
                                                                } else {
                                                                    var v66 float64 = v2
                                                                    _ = v66
                                                                }
                                                            case 348:
                                                                _ = len(v8)
                                                            }
                                                        default:
                                                            if false {
                                                                v6[fi25] = 2147483648
                                                                var v67 float64 = (v32 / -2.25)
                                                                var v68 uint = v44
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                        if true {
                                                            v69 := GenS2{F0: 267, F1: v3, F2: false}
                                                            _ = v69
                                                        } else {
                                                            fmt.Println(fi25)
                                                            var v70 int = -902
                                                            _ = v70
                                                        }
                                                        v6[fi33] = 4294967295
                                                    }
                                                case "\n\t":
                                                    var v71 uint64 = (v9 / (((v9 + (^(v9)))) | 1))
                                                    _ = v71
                                                case "foo bar":
                                                    var v72 uint64 = (v9 - v9)
                                                    _ = v72
                                                }
                                            }
                                        case "a":
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        default:
                                            switch "a" {
                                            case "\n\t":
                                                if true {
                                                    var v73 string = ""
                                                    if false {
                                                        var v74 int16 = v1
                                                        fmt.Println(v74, v2)
                                                        _ = v74
                                                    } else {
                                                        fmt.Println(v3, v2)
                                                    }
                                                    v2 = 3.14159
                                                    _ = v73
                                                } else {
                                                    var v75 uint32 = 51
                                                    var v76 string = ("\n\t" + "foo bar")
                                                    v77 := GenS5{F0: 1000}
                                                    var v78 int = int(v75)
                                                    _ = v75
                                                    _ = v76
                                                    _ = v77
                                                    _ = v78
                                                }
                                            case "":
                                                switch fi25 {
                                                case 0:
                                                    var v79 string = v10
                                                    _ = v79
                                                default:
                                                    var v80 uint16 = uint16(v44)
                                                    _ = v80
                                                }
                                            }
                                        }
                                        _ = v44
                                    default:
                                        var v81 int32 = (v4 / ((v11) | 1))
                                        v10 = v10
                                        _ = v81
                                    }
                                    _ = v42
                                }
                                _ = v39
                            } else {
                                if true {
                                    v82 := GenS2{F0: 896, F1: (v3 % ((v3) | 1)), F2: false}
                                    v5 -= -2.25
                                    _ = v82
                                } else {
                                    if (v32 > (float64(v5) * (v2 / v2))) {
                                        if (65535 >= 493) {
                                            var v83 uint32 = 581
                                            switch v10 {
                                            case "":
                                                v2 -= v32
                                                if false {
                                                    var v84 int64 = v3
                                                    _ = v84
                                                } else {
                                                    var v85 int16 = -32768
                                                    _ = v85
                                                }
                                                v5 -= (v5 / ((v5 - v5) / ((1e10 / 1e10) - (0.0 + v5))))
                                            case "a":
                                                v86 := GenS6{F0: int(v1)}
                                                _ = v86
                                            case "foo bar":
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            default:
                                                fmt.Println(v83)
                                                switch ("" + ((v10 + "") + "hello")) {
                                                case "hello":
                                                    v87 := GenS7{F0: fi25}
                                                    _ = v87
                                                case "\n\t":
                                                    if len(v8) > 0 {
                                                        v8[0] = (fi33 % ((fi33) | 1))
                                                    }
                                                case "":
                                                    v88 := v6[1]
                                                    _ = v88
                                                default:
                                                    fmt.Println(fi33, v14, fi25)
                                                }
                                            }
                                            _ = v83
                                        }
                                        v89 := GenS8{F0: (true && false), F1: 1, F2: v21}
                                        _ = v89
                                    } else {
                                        fmt.Println(v4, v32, v21)
                                        if (false || false) {
                                            delete(v6, 679)
                                            v90 := GenS9{F0: 2, F1: (^(v14)), F2: 951}
                                            var v91 int16 = (v1 >> ((0) & 63))
                                            fmt.Println(v9)
                                            _ = v90
                                            _ = v91
                                        }
                                        if (!false) {
                                            v14 = v14
                                            switch int(v1) {
                                            case -508:
                                                switch v10 {
                                                case "\n\t":
                                                    v92 := GenS10{F0: true, F1: (^(v14))}
                                                    _ = v92
                                                case "a":
                                                    switch fi33 {
                                                    case 621:
                                                        switch (v10 + v10) {
                                                        case "a":
                                                            v93 := GenS1{F0: (v21 % (((v21 | v21)) | 1))}
                                                            _ = v93
                                                        case "\n\t":
                                                            var v94 int32 = 42
                                                            _ = v94
                                                        default:
                                                            v5 -= v5
                                                        }
                                                    case 32767:
                                                        var v95 uint = 4294967295
                                                        _ = v95
                                                    case -2147483648:
                                                        _ = len(v8)
                                                    }
                                                case "hello":
                                                    v10 = v10
                                                default:
                                                    fmt.Println(v32, fi33, v14)
                                                }
                                            case 4294967295:
                                                var v96 int32 = (v11 & (v11 & -605))
                                                v2 = v2
                                                _ = v96
                                            }
                                            var v97 int32 = v11
                                            _ = v97
                                        }
                                    }
                                    delete(v6, fi33)
                                }
                                fmt.Println(v32, v32, v32)
                                v6[-2147483648] = 2147483648
                            }
                        }
                        switch (("" + v10) + v10) {
                        case "a":
                            v4 -= (v11 - (v11 - -518))
                        }
                        switch ("foo bar" + "") {
                        case "a":
                            if false {
                                if true {
                                    fmt.Println(v4, v21)
                                    if (0 < 518) {
                                        v98 := GenS11{F0: (v1 & 32767)}
                                        fmt.Println(v1)
                                        var v99 float32 = -2.25
                                        v100 := GenS12{F0: 788}
                                        var v101 int16 = (v1 & v1)
                                        _ = v100
                                        _ = v101
                                        _ = v98
                                        _ = v99
                                    }
                                }
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                        case "\n\t":
                            fmt.Println(v2, fi33)
                        case "":
                            v102 := GenS2{F0: 4294967295, F1: (v3 % ((v3) | 1)), F2: true}
                            _ = v102
                        default:
                            switch ((v10 + v10) + "foo bar") {
                            case "hello":
                                v11 = (v4 ^ 42)
                            case "":
                                var v103 float32 = float32(v2)
                                _ = v103
                            case "a":
                                fmt.Println(v32, v4)
                            }
                        }
                        var v104 int64 = (v3 | (v3 + (v3 % (((^(v3))) | 1))))
                        var v105 string = ""
                        switch "" {
                        case "\n\t":
                            fmt.Println(v1, v4, v5)
                        case "hello":
                            var v106 uint = 126
                            _ = v106
                        case "":
                            v11 += 0
                        default:
                            var v107 int32 = v11
                            _ = v107
                        }
                        switch "\n\t" {
                        case "":
                            if (!(!true)) {
                                fmt.Println(fi25, v21, v11)
                                var v108 uint8 = (^(v14))
                                _ = v108
                            }
                        case "\n\t":
                            v10 = (v10 + ("a" + v105))
                        case "a":
                            v14 -= 42
                        default:
                            v3 += (^(v104))
                        }
                        if false {
                            var v109 string = v105
                            _ = v109
                        } else {
                            fmt.Println(v104)
                        }
                        switch int(v9) {
                        case -1:
                            v11 = v4
                        }
                        delete(v6, (-(fi25)))
                        _ = fi33
                        _ = v104
                        _ = v105
                    }
                    if true {
                        var v110 int8 = v21
                        var v111 uint64 = uint64(v14)
                        fmt.Println(v14, v11, v110)
                        _ = v110
                        _ = v111
                    }
                    _ = v32
                }
            } else {
                for fi112 := 0; fi112 < 19; fi112++ {
                    if ((v2 + float64((v5 + v5))) != 3.14159) {
                        switch -1 {
                        case 32768:
                            v113 := GenS8{F0: false, F1: 255, F2: (-(v21))}
                            _ = v113
                        case -1:
                            var v114 uint32 = 2147483647
                            v115 := GenS13{F0: (^(v4))}
                            v21 = v21
                            _ = v114
                            _ = v115
                        }
                        v11 -= (-(v4))
                    }
                    if (-1 == v21) {
                        switch -498 {
                        case 32767:
                            v116 := GenS4{F0: v10}
                            _ = v116
                        case -2147483648:
                            if true {
                                v11 += 939
                                v8 = append(v8, fi112)
                            } else {
                                v4 = v4
                                var v117 uint8 = (v14 << ((1) & 63))
                                if false {
                                    _ = len(v8)
                                    fmt.Println(v11)
                                    var v118 int16 = (v1 + (v1 ^ -1))
                                    _ = v118
                                } else {
                                    v11 += (v4 / ((v4) | 1))
                                }
                                _ = v117
                            }
                        default:
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        }
                    }
                    v9 = 915
                    v4 = v4
                    _ = fi112
                }
            }
            fmt.Println(v9, v14)
            _ = fi25
        }
        var v119 uint = 2
        _ = v119
    } else {
        v2 -= 0.0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
