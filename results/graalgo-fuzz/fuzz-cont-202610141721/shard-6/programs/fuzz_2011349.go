package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS3 struct { F0 string; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int16; F1 int16 }
type GenS8 struct { F0 int32; F1 uint64; F2 int64 }
type GenS9 struct { F0 int; F1 uint32; F2 uint8 }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int16; F1 uint8 }
type GenS13 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 int64; F1 int64; F2 bool }
type GenS18 struct { F0 int8; F1 int; F2 bool }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]string)
    var v3 uint32 = 1
    v4 := make(map[int]int)
    var v5 int32 = 0
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(v1, fi6, v3)
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v3, v1, fi7)
            v1 = (v1 / ((int8(v3)) | 1))
            if true {
                v1 = v1
                if true {
                    var v8 uint16 = 2
                    v1 += v1
                    v9 := GenS1{F0: 749, F1: 77}
                    if (true && false) {
                        v8 = (v8 - 65535)
                    } else {
                        _ = len(v2)
                        if ((true || true) && true) {
                            v3 -= (v3 / ((v3) | 1))
                        } else {
                            v8 = (v8 << ((4294967295) & 63))
                        }
                        v10 := GenS1{F0: 42, F1: v1}
                        _ = v10
                    }
                    var v11 int64 = int64(fi6)
                    var v12 string = ""
                    v13 := GenS2{F0: 788, F1: 4, F2: v5}
                    _ = v11
                    _ = v12
                    _ = v13
                    _ = v8
                    _ = v9
                }
            }
            v1 += (v1 - 127)
            _ = fi7
        }
        _ = fi6
    }
    v14 := GenS2{F0: 317, F1: 255, F2: 2147483647}
    v1 -= (^(v1))
    v1 += -67
    switch "foo bar" {
    case "a":
        v1 += 127
    case "hello":
        switch (("\n\t" + "\n\t") + "a") {
        case "\n\t":
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(v5)
                var v16 bool = false
                _ = fi15
                _ = v16
            }
        case "a":
            fmt.Println(v1)
        default:
            switch 809 {
            case 32767:
                v17 := GenS2{F0: 4294967295, F1: 1, F2: v5}
                _ = v17
            case 4294967295:
                fmt.Println(v3, v3, v1)
            case 2147483647:
                for fi18 := 0; fi18 < 7; fi18++ {
                    if (((0.1 - -2.25) / (0.0 + (0.0 - (1.5 / ((float32(-2.25) + (float32(3.14159) / (float32(0.1) + 1.5))) - float32(float64(1.5))))))) < 0.1) {
                        v19 := GenS2{F0: 1, F1: 30, F2: v5}
                        _ = v19
                    }
                    _ = fi18
                }
            }
        }
    default:
        var v20 int16 = 50
        _ = v20
    }
    switch 127 {
    case 0:
        var v21 int16 = -62
        _ = v21
    case 127:
        var v22 int64 = 1
        _ = v22
    case 128:
        if false {
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := GenS3{F0: "foo bar", F1: fi23, F2: 42}
                for fi25 := 0; fi25 < 17; fi25++ {
                    var v26 uint8 = 127
                    if false {
                        v27 := GenS4{F0: v26, F1: 42}
                        fmt.Println(fi25, v1, fi25)
                        v28 := []int{int(v26)}
                        _ = v27
                        _ = v28
                    } else {
                        v3 = 1
                        v26 += (v26 & 1)
                    }
                    fmt.Println(v1)
                    _ = fi25
                    _ = v26
                }
                _ = fi23
                _ = v24
            }
        } else {
            for fi29 := 0; fi29 < 2; fi29++ {
                v3 = (v3 | (v3 >> ((95) & 63)))
                for fi30 := 0; fi30 < 1; fi30++ {
                    v5 = (-(v5))
                    v5 = v5
                    _ = fi30
                }
                v1 += v1
                _ = fi29
            }
        }
    }
    v31 := GenS5{F0: 4294967295, F1: 401, F2: 2}
    for fi32 := 0; fi32 < 3; fi32++ {
        v5 -= (v5 - v5)
        for fi33 := 0; fi33 < 11; fi33++ {
            v34 := []int{}
            if false {
                switch "\n\t" {
                case "foo bar":
                    if true {
                        v2[("" + ("" + ("a" + (("foo bar" + "") + "hello"))))] = "\n\t"
                        if len(v34) > 0 {
                            v34[0] = int(v3)
                        }
                        v1 += int8(v5)
                    } else {
                        if (42 <= 0) {
                            var v35 int64 = int64(v1)
                            if (("hello" + "hello") == "a") {
                                switch (fi32 & (fi33 & (fi32 / ((fi33) | 1)))) {
                                case 127:
                                    var v36 float32 = 1e10
                                    _ = v36
                                case 32768:
                                    if (v5 == v5) {
                                        var v37 uint64 = 4294967295
                                        v38 := GenS1{F0: -959, F1: 42}
                                        switch fi33 {
                                        case 9223372036854775807:
                                            v1 += -128
                                        case -633:
                                            v34 = append(v34, int(v5))
                                        default:
                                            v37 = 2
                                        }
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v35 = v35
                                        switch int(v35) {
                                        case 190:
                                            v35 = -1
                                        case 128:
                                            var v39 uint8 = 132
                                            var v40 int16 = int16(v5)
                                            switch fi32 {
                                            case 1:
                                                switch "" {
                                                case "\n\t":
                                                    v41 := GenS6{F0: v35, F1: v5}
                                                    _ = v41
                                                }
                                            case 32768:
                                                var v42 string = ""
                                                v43 := GenS6{F0: 0, F1: (-(v5))}
                                                v44 := GenS6{F0: v35, F1: v5}
                                                _ = v42
                                                _ = v43
                                                _ = v44
                                            default:
                                                v45 := GenS2{F0: 4294967295, F1: v39, F2: 722}
                                                _ = v45
                                            }
                                            _ = v39
                                            _ = v40
                                        case 0:
                                            if (v5 >= v5) {
                                                var v46 uint = 856
                                                _ = v46
                                            } else {
                                                var v47 int16 = 1
                                                if ((v47 ^ (v47 >> ((65535) & 63))) > (v47 & v47)) {
                                                    v48 := GenS7{F0: (v47 ^ (v47 * v47)), F1: -32768}
                                                    fmt.Println(v35, v35, v3)
                                                    _ = v48
                                                }
                                                var v49 int64 = -9223372036854775808
                                                _ = v47
                                                _ = v49
                                            }
                                        }
                                    }
                                case 4294967295:
                                    var v50 int64 = (v35 ^ (v35 - v35))
                                    _ = v50
                                }
                                switch "\n\t" {
                                case "hello":
                                    if len(v34) >= 2 {
                                        v34 = v34[1:len(v34)]
                                    }
                                case "a":
                                    _ = len(v34)
                                case "":
                                    if (!false) {
                                        switch fi32 {
                                        case 1:
                                            switch (("\n\t" + "foo bar") + "foo bar") {
                                            case "\n\t":
                                                fmt.Println(fi32, fi32)
                                            case "a":
                                                _ = len(v34)
                                            }
                                        case -2147483648:
                                            switch fi32 {
                                            case -2147483648:
                                                v1 += (^(v1))
                                            case 9223372036854775807:
                                                if (((v1 | (v1 ^ v1)) >= (v1 << ((uint(v3)) & 63))) || true) {
                                                    var v51 int16 = 332
                                                    _ = v51
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        switch fi32 {
                                                        case -372:
                                                            delete(v2, ("a" + (("\n\t" + "") + (("hello" + "") + ""))))
                                                        case -723:
                                                            var v52 float64 = 0.0
                                                            _ = v52
                                                        default:
                                                            switch (fi32 | fi32) {
                                                            case 42:
                                                                switch (fi32 >> ((65535) & 63)) {
                                                                case 2147483647:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v35 = v35
                                                                    case "a":
                                                                        if (false || (fi32 > (fi32 ^ (fi32 | -813)))) {
                                                                            if (float64(1e10) < (3.14159 * 1e10)) {
                                                                                v35 = (v35 / ((v35) | 1))
                                                                                var v53 int64 = v35
                                                                                _ = v53
                                                                            } else {
                                                                                if (42 < int16(fi33)) {
                                                                                    v54 := v4[fi32]
                                                                                    _ = v54
                                                                                } else {
                                                                                    switch (fi32 >> ((255) & 63)) {
                                                                                    case -2147483648:
                                                                                        fmt.Println(v1, fi33)
                                                                                    default:
                                                                                        if len(v34) > 0 {
                                                                                            v34[0] = (fi33 | 707)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(v1, v35)
                                                                                v55 := GenS1{F0: (v35 / ((v35) | 1)), F1: 1}
                                                                                _ = v55
                                                                            }
                                                                            var v56 int32 = (v5 % (((v5 ^ int32(v35))) | 1))
                                                                            v3 = 338
                                                                            _ = v56
                                                                        }
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v57 string = "\n\t"
                                                                    _ = v57
                                                                }
                                                            case 32768:
                                                                fmt.Println(v5, fi32, v3)
                                                            case -701:
                                                                v1 += 76
                                                            default:
                                                                v5 = v5
                                                            }
                                                            if len(v34) >= 2 {
                                                                v34 = v34[1:len(v34)]
                                                            }
                                                        }
                                                    case "a":
                                                        switch "a" {
                                                        case "\n\t":
                                                            if len(v34) > 0 {
                                                                v34[0] = fi33
                                                            }
                                                        case "foo bar":
                                                            if (0 >= 25) {
                                                                v35 = 0
                                                                v58 := GenS8{F0: (v5 >> ((0) & 63)), F1: 1, F2: -9223372036854775808}
                                                                v3 = v3
                                                                switch "a" {
                                                                case "hello":
                                                                    var v59 string = "foo bar"
                                                                    _ = v59
                                                                default:
                                                                    _ = len(v34)
                                                                }
                                                                _ = v58
                                                            }
                                                        case "a":
                                                            fmt.Println(v3)
                                                        default:
                                                            fmt.Println(v35, fi33, fi32)
                                                            v1 -= v1
                                                            if (true || true) {
                                                                if false {
                                                                    v60 := GenS9{F0: (fi33 << ((432) & 63)), F1: v3, F2: 0}
                                                                    v61 := GenS10{F0: v5, F1: -603}
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    v3 -= v3
                                                                }
                                                            }
                                                        }
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            if true {
                                                                if (1 == -32768) {
                                                                    fmt.Println(v3, fi33)
                                                                    switch fi32 {
                                                                    case 42:
                                                                        var v62 float32 = 0.1
                                                                        _ = v62
                                                                    case 2147483647:
                                                                        var v63 uint32 = v3
                                                                        _ = v63
                                                                    case 389:
                                                                        if (42 >= 35) {
                                                                            switch (fi32 ^ (fi32 % (((fi33 >> ((255) & 63))) | 1))) {
                                                                            case 410:
                                                                                fmt.Println(fi32, v5)
                                                                            case -455:
                                                                                fmt.Println(fi33, v5, v3)
                                                                            }
                                                                        } else {
                                                                            v3 = 2147483648
                                                                        }
                                                                    default:
                                                                        v64 := v4[fi33]
                                                                        _ = v64
                                                                    }
                                                                }
                                                                fmt.Println(fi32)
                                                                v1 += int8(v5)
                                                                var v65 int16 = 212
                                                                if false {
                                                                    if false {
                                                                        v3 += (v3 - v3)
                                                                    }
                                                                    v2["foo bar"] = ("hello" + "foo bar")
                                                                    v5 = (v5 + (v5 >> ((255) & 63)))
                                                                    fmt.Println(fi33, fi32, v1)
                                                                    v35 = (v35 << ((36) & 63))
                                                                } else {
                                                                    fmt.Println(v5, v1, v3)
                                                                }
                                                                _ = v65
                                                            } else {
                                                                fmt.Println(v3, fi33, v1)
                                                                if ((v1 != (v1 << ((78) & 63))) || false) {
                                                                    if false {
                                                                        v35 += v35
                                                                    } else {
                                                                        if (true && (false && true)) {
                                                                            if false {
                                                                                if (!(!false)) {
                                                                                    v35 = (v35 >> ((4294967295) & 63))
                                                                                    var v66 bool = false
                                                                                    _ = v66
                                                                                } else {
                                                                                    v67 := GenS11{F0: 863}
                                                                                    v68 := GenS12{F0: -1, F1: 127}
                                                                                    var v69 float32 = (float32((1e10 / 3.14159)) / -2.25)
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                    _ = v69
                                                                                }
                                                                                v70 := GenS10{F0: v5, F1: v5}
                                                                                v71 := v4[(^(fi32))]
                                                                                _ = v70
                                                                                _ = v71
                                                                            } else {
                                                                                if false {
                                                                                    if ((v1 * v1) != -1) {
                                                                                        fmt.Println(v1, fi32, v1)
                                                                                        switch (fi33 << ((872) & 63)) {
                                                                                        case 32767:
                                                                                            var v72 uint32 = v3
                                                                                            _ = v72
                                                                                        case 127:
                                                                                            if (2147483648 == (v3 % (((v3 ^ (v3 * 1))) | 1))) {
                                                                                                var v73 int16 = 1000
                                                                                                _ = v73
                                                                                            } else {
                                                                                                if false {
                                                                                                    v1 -= (v1 | 100)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v74 uint16 = 32767
                                                                                    _ = v74
                                                                                }
                                                                            }
                                                                            v35 += (v35 - -697)
                                                                        }
                                                                        fmt.Println(v3)
                                                                        fmt.Println(fi32, v35)
                                                                    }
                                                                }
                                                                v1 -= v1
                                                            }
                                                        case "hello":
                                                            if true {
                                                                v4[fi33] = 600
                                                                v75 := v4[(fi32 - (fi33 | (fi32 - fi33)))]
                                                                _ = v75
                                                            }
                                                        case "\n\t":
                                                            switch ("" + "hello") {
                                                            case "\n\t":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    v1 = (v1 ^ (v1 >> ((1) & 63)))
                                                                default:
                                                                    var v76 float32 = 0.1
                                                                    _ = v76
                                                                }
                                                            case "foo bar":
                                                                var v77 int16 = 42
                                                                _ = v77
                                                            case "a":
                                                                v78 := GenS13{F0: v35, F1: 42, F2: v3}
                                                                _ = v78
                                                            default:
                                                                switch -1 {
                                                                case 4294967295:
                                                                    v35 = v35
                                                                case -730:
                                                                    v79 := GenS12{F0: -335, F1: 90}
                                                                    _ = v79
                                                                case -1:
                                                                    v80 := GenS14{F0: 4294967295, F1: 255}
                                                                    _ = v80
                                                                default:
                                                                    var v81 string = ""
                                                                    _ = v81
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            v82 := GenS4{F0: 85, F1: 2147483647}
                                            _ = v82
                                        }
                                    } else {
                                        var v83 uint64 = 1
                                        _ = v83
                                    }
                                default:
                                    if (0 == 255) {
                                        var v84 uint = 65535
                                        _ = v84
                                    }
                                }
                            } else {
                                if true {
                                    var v85 uint = 1
                                    var v86 float32 = 0.1
                                    _ = v85
                                    _ = v86
                                } else {
                                    if (!false) {
                                        if len(v34) > 0 {
                                            v34[0] = (fi33 % (((fi32 | 42)) | 1))
                                        }
                                        v87 := GenS5{F0: -2147483648, F1: 2, F2: 65535}
                                        v1 += v1
                                        _ = v87
                                    } else {
                                        if false {
                                            if ("a" >= "\n\t") {
                                                fmt.Println(v35)
                                                fmt.Println(v35, v5)
                                            }
                                            var v88 uint64 = 582
                                            var v89 uint8 = 166
                                            _ = v88
                                            _ = v89
                                        } else {
                                            switch fi32 {
                                            case 127:
                                                if true {
                                                    v35 += (v35 - 9223372036854775807)
                                                    var v90 uint = 255
                                                    var v91 bool = false
                                                    v1 = 42
                                                    var v92 string = ((("\n\t" + "") + (("a" + "\n\t") + "")) + "a")
                                                    _ = v90
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    fmt.Println(fi32, fi33)
                                                }
                                            default:
                                                switch "a" {
                                                case "":
                                                    fmt.Println(v1)
                                                    v93 := GenS7{F0: 32767, F1: 0}
                                                    if true {
                                                        if (1.5 == -2.25) {
                                                            var v94 uint8 = 1
                                                            v95 := GenS8{F0: v5, F1: 1, F2: v35}
                                                            _ = v94
                                                            _ = v95
                                                        } else {
                                                            var v96 int8 = (v1 & v1)
                                                            if len(v34) >= 2 {
                                                                v34 = v34[1:len(v34)]
                                                            }
                                                            var v97 int = fi33
                                                            switch (fi33 % ((-2147483648) | 1)) {
                                                            case 127:
                                                                v1 = v1
                                                                if false {
                                                                    v98 := GenS5{F0: 872, F1: 2, F2: 65535}
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                    var v99 bool = true
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    var v100 float32 = -2.25
                                                                    _ = v100
                                                                }
                                                            case 937:
                                                                var v101 int = -939
                                                                _ = v101
                                                            case -290:
                                                                if (764 <= 2) {
                                                                    var v102 int64 = (-(v35))
                                                                    fmt.Println(v5, v96, fi33)
                                                                    _ = len(v4)
                                                                    switch (^(fi32)) {
                                                                    case 1:
                                                                        if (!(!false)) {
                                                                            fmt.Println(v97, v35, v96)
                                                                        }
                                                                    case 0:
                                                                        if (3.14159 >= ((1e10 * ((0.1 + 1e10) + 1.5)) * (float64(3.14159) * 1.5))) {
                                                                            var v103 float64 = float64(1.5)
                                                                            _ = v103
                                                                        } else {
                                                                            v97 = int(v3)
                                                                            var v104 int8 = v1
                                                                            _ = v104
                                                                        }
                                                                    default:
                                                                        v34 = append(v34, 2147483647)
                                                                    }
                                                                    v1 += 127
                                                                    _ = v102
                                                                } else {
                                                                    var v105 uint8 = 68
                                                                    if (!true) {
                                                                        var v106 uint8 = v105
                                                                        v1 += v96
                                                                        fmt.Println(v3, fi32, v106)
                                                                        _ = v106
                                                                    }
                                                                    _ = v105
                                                                }
                                                            default:
                                                                v107 := GenS3{F0: "\n\t", F1: -238, F2: 0}
                                                                _ = v107
                                                            }
                                                            var v108 int64 = -9223372036854775808
                                                            if len(v34) >= 2 {
                                                                v34 = v34[1:len(v34)]
                                                            }
                                                            var v109 uint32 = 417
                                                            if true {
                                                                var v110 int = 2147483647
                                                                v97 = 9223372036854775807
                                                                var v111 int = fi33
                                                                _ = v110
                                                                _ = v111
                                                            } else {
                                                                v3 -= 2
                                                                v1 = 54
                                                                switch "a" {
                                                                case "":
                                                                    fmt.Println(v35)
                                                                }
                                                                var v112 float64 = 3.14159
                                                                _ = v112
                                                            }
                                                            v109 = 42
                                                            _ = v108
                                                            _ = v109
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        v113 := GenS2{F0: 1, F1: 255, F2: (v5 - v5)}
                                                        _ = v113
                                                    }
                                                    v114 := GenS9{F0: fi33, F1: (v3 >> ((42) & 63)), F2: 107}
                                                    var v115 float64 = float64(1e10)
                                                    _ = v114
                                                    _ = v115
                                                    _ = v93
                                                case "\n\t":
                                                    var v116 uint64 = 1
                                                    _ = v116
                                                case "a":
                                                    var v117 uint = 4294967295
                                                    _ = v117
                                                default:
                                                    if false {
                                                        fmt.Println(fi32)
                                                        var v118 uint8 = 1
                                                        fmt.Println(v118)
                                                        _ = v118
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    var v119 uint32 = v3
                                    v120 := GenS14{F0: 2, F1: 1}
                                    _ = v119
                                    _ = v120
                                }
                                var v121 uint16 = 2
                                fmt.Println(v35, v5, v121)
                                v122 := GenS15{F0: false}
                                if false {
                                    v123 := GenS2{F0: 117, F1: 127, F2: (v5 << ((65535) & 63))}
                                    _ = v123
                                } else {
                                    v124 := GenS16{F0: -1}
                                    var v125 string = ("" + ("a" + ("\n\t" + (("hello" + "a") + "hello"))))
                                    _ = v124
                                    _ = v125
                                }
                                _ = v121
                                _ = v122
                            }
                            var v126 int64 = v35
                            v4[fi32] = int(v1)
                            fmt.Println(fi33)
                            _ = v126
                            _ = v35
                        } else {
                            delete(v2, "")
                        }
                    }
                default:
                    var v127 float32 = float32((float64(0.0) * float64(3.14159)))
                    _ = v127
                }
                var v128 uint64 = 2
                var v129 float64 = (0.0 + (0.0 + 1.5))
                var v130 float64 = v129
                v3 += (v3 << ((850) & 63))
                v131 := GenS4{F0: 220, F1: (fi32 >> ((1) & 63))}
                _ = v128
                _ = v129
                _ = v130
                _ = v131
            }
            v3 -= (^(v3))
            var v132 int64 = -9223372036854775808
            v3 += 1
            _ = fi33
            _ = v132
            _ = v34
        }
        v133 := []int{fi32, (fi32 / ((507) | 1))}
        _ = fi32
        _ = v133
    }
    v134 := []int{42}
    if len(v134) >= 2 {
        v134 = v134[1:len(v134)]
    }
    switch ("" + (("foo bar" + "\n\t") + ("hello" + ("hello" + "a")))) {
    case "":
        for fi135 := 0; fi135 < 20; fi135++ {
            var v136 int8 = v1
            _ = fi135
            _ = v136
        }
    case "a":
        v137 := GenS5{F0: -2147483648, F1: 733, F2: 42}
        _ = v137
    case "\n\t":
        v138 := GenS17{F0: -9223372036854775808, F1: 9223372036854775807, F2: true}
        _ = v138
    }
    v139 := GenS18{F0: 42, F1: int(v5), F2: (42 != v5)}
    for fi140 := 0; fi140 < 2; fi140++ {
        v3 = (v3 - (v3 | (v3 & 4294967295)))
        _ = fi140
    }
    v141 := GenS11{F0: -1}
    v1 = int8(v3)
    v1 += -128
    if len(v134) > 0 {
        v134[0] = -61
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v134:", len(v134), v134)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v134
    _ = v139
    _ = v14
    _ = v141
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
