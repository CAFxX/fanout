package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int; F2 int16 }
type GenS8 struct { F0 uint8; F1 bool; F2 int8 }
type GenS9 struct { F0 int32; F1 int64 }
type GenS10 struct { F0 uint16; F1 int32; F2 int }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 int64 }
func genFn0() int16 {
    var v5 uint8 = 228
    v6 := GenS1{F0: 255, F1: 894}
    if false {
        fmt.Println(v5)
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                if (42 > 32767) {
                    fmt.Println(fi8, v5, fi8)
                }
                _ = fi8
            }
            var v9 uint8 = 2
            v10 := []int{}
            _ = fi7
            _ = v10
            _ = v9
        }
    }
    v5 = 222
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := []int{fi11, 32767, fi11, 128}
            _ = fi11
            _ = v12
        }
    case "\n\t":
        if true {
            if false {
                if false {
                    switch "" {
                    case "hello":
                        v13 := GenS2{F0: 324, F1: 0, F2: 215}
                        _ = v13
                    case "\n\t":
                        v5 += (v5 * v5)
                    default:
                        v14 := GenS3{F0: 1}
                        _ = v14
                    }
                    var v15 uint8 = v5
                    var v16 uint = 375
                    _ = v15
                    _ = v16
                }
                v5 -= (v5 * v5)
                v5 = 15
            }
            if (128 >= 128) {
                var v17 string = ("a" + ((((("hello" + "\n\t") + "hello") + "\n\t") + "foo bar") + ((("hello" + ("a" + "hello")) + "hello") + ("\n\t" + "foo bar"))))
                _ = v17
            }
        }
    case "foo bar":
        if (-2.25 < (3.14159 * 1e10)) {
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := []int{(fi18 << ((1) & 63)), (fi18 >> ((4294967295) & 63)), -231}
                _ = fi18
                _ = v19
            }
            switch "foo bar" {
            case "foo bar":
                switch 1 {
                case 0:
                    var v20 uint16 = 140
                    _ = v20
                }
            case "":
                for fi21 := 0; fi21 < 17; fi21++ {
                    v22 := GenS2{F0: uint16(v5), F1: 18446744073709551615, F2: 1}
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        if (!((v5 >> ((331) & 63)) >= v5)) {
                            for fi23 := 0; fi23 < 4; fi23++ {
                                v5 = (v5 >> ((448) & 63))
                                _ = fi23
                            }
                            for fi24 := 0; fi24 < 7; fi24++ {
                                switch "\n\t" {
                                case "\n\t":
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        if (42 == 681) {
                                            v25 := []int{fi21, (fi24 & fi21), 32767, 243}
                                            _ = v25
                                        }
                                    default:
                                        switch "a" {
                                        case "hello":
                                            if ((((1.5 * (0.1 / 3.14159)) + 0.0) * 3.14159) != float64((0.0 + ((1.5 - (0.0 * ((0.1 + 0.0) * (1e10 * 0.0)))) + 0.0)))) {
                                                v26 := []int{}
                                                var v27 uint = 994
                                                if (true || (4294967295 == 747)) {
                                                    switch ((("a" + "a") + "hello") + "foo bar") {
                                                    case "foo bar":
                                                        var v28 uint32 = 2147483647
                                                        _ = v28
                                                    case "\n\t":
                                                        v5 += v5
                                                    case "hello":
                                                        _ = len(v26)
                                                    default:
                                                        v5 -= v5
                                                    }
                                                    var v29 uint16 = 255
                                                    switch (fi24 * 2147483647) {
                                                    case 32768:
                                                        if false {
                                                            var v30 uint16 = (v29 | 1)
                                                            switch fi24 {
                                                            case -2147483648:
                                                                v31 := GenS3{F0: uint64(v30)}
                                                                var v32 int16 = -807
                                                                v5 = (v5 >> ((v27) & 63))
                                                                _ = v31
                                                                _ = v32
                                                            case 17:
                                                                switch (^(fi21)) {
                                                                case 128:
                                                                    if false {
                                                                        fmt.Println(v27, v27)
                                                                    } else {
                                                                        switch (("a" + "") + ("" + "foo bar")) {
                                                                        case "hello":
                                                                            if (1 < 905) {
                                                                                var v33 float64 = 1e10
                                                                                _ = v33
                                                                            } else {
                                                                                fmt.Println(v5, v5)
                                                                                var v34 uint16 = (v29 % ((v29) | 1))
                                                                                if (-2.25 == 0.1) {
                                                                                    fmt.Println(v30, fi24, fi24)
                                                                                } else {
                                                                                    if len(v26) > 0 {
                                                                                        v26[0] = fi24
                                                                                    }
                                                                                    if false {
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            v35 := GenS1{F0: v29, F1: (v29 * (v34 - 2))}
                                                                                            _ = v35
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v34
                                                                            }
                                                                            if true {
                                                                                v29 += v30
                                                                                if len(v26) >= 2 {
                                                                                    v26 = v26[1:len(v26)]
                                                                                }
                                                                                fmt.Println(fi21, fi24, v29)
                                                                                if false {
                                                                                    v36 := GenS3{F0: 4294967295}
                                                                                    v5 += v5
                                                                                    v37 := GenS4{F0: -441, F1: v29}
                                                                                    _ = v36
                                                                                    _ = v37
                                                                                } else {
                                                                                    var v38 int = 589
                                                                                    _ = v38
                                                                                }
                                                                                switch -1 {
                                                                                case 32767:
                                                                                    fmt.Println(v29, v5)
                                                                                case 2147483647:
                                                                                    v39 := GenS5{F0: -1, F1: (v30 & (^(v30))), F2: v30}
                                                                                    _ = v39
                                                                                case 251:
                                                                                    v26 = append(v26, fi21)
                                                                                }
                                                                            }
                                                                        case "foo bar":
                                                                            var v40 bool = true
                                                                            _ = v40
                                                                        default:
                                                                            if true {
                                                                                if true {
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        v26 = append(v26, -1)
                                                                                    case "foo bar":
                                                                                        fmt.Println(fi24)
                                                                                    default:
                                                                                        v41 := GenS5{F0: 1, F1: uint16(v5), F2: (v30 / (((v29 ^ v29)) | 1))}
                                                                                        _ = v41
                                                                                    }
                                                                                    fmt.Println(v5, fi21, v29)
                                                                                } else {
                                                                                    switch -474 {
                                                                                    case 32767:
                                                                                        fmt.Println(v29)
                                                                                    case 1:
                                                                                        var v42 string = "foo bar"
                                                                                        _ = v42
                                                                                    case 0:
                                                                                        v29 = v29
                                                                                    }
                                                                                    fmt.Println(v29, fi24, v30)
                                                                                }
                                                                                v27 = 2
                                                                            }
                                                                        }
                                                                        if len(v26) > 0 {
                                                                            v26[0] = (fi24 & (fi24 >> (((v27 / ((18446744073709551615) | 1))) & 63)))
                                                                        }
                                                                        v30 = v30
                                                                    }
                                                                case 1:
                                                                    if false {
                                                                        var v43 uint = (^(v27))
                                                                        v43 += v43
                                                                        if true {
                                                                            switch ("" + "a") {
                                                                            case "":
                                                                                v30 = 269
                                                                            case "foo bar":
                                                                                var v44 float64 = 3.14159
                                                                                v45 := GenS3{F0: 1}
                                                                                if true {
                                                                                    v46 := GenS4{F0: 0, F1: 42}
                                                                                    var v47 uint8 = (^(v5))
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                }
                                                                                v5 = 58
                                                                                switch (("hello" + "hello") + "hello") {
                                                                                case "foo bar":
                                                                                    switch ("foo bar" + "hello") {
                                                                                    case "":
                                                                                        fmt.Println(fi21, fi21)
                                                                                    case "foo bar":
                                                                                        var v48 uint64 = 967
                                                                                        _ = v48
                                                                                    }
                                                                                case "":
                                                                                    v49 := GenS1{F0: (v29 % ((915) | 1)), F1: uint16(v5)}
                                                                                    _ = v49
                                                                                case "hello":
                                                                                    fmt.Println(v5, v43, v44)
                                                                                    v26 = append(v26, fi21)
                                                                                }
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            v27 -= v43
                                                                            switch (fi24 * int(v5)) {
                                                                            case 127:
                                                                                if true {
                                                                                    v43 -= (v43 ^ 2)
                                                                                    var v50 int16 = 32767
                                                                                    _ = len(v26)
                                                                                    _ = v50
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v30 = uint16(v5)
                                                                                    case "\n\t":
                                                                                        v29 = v29
                                                                                    default:
                                                                                        v27 = 42
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        switch fi24 {
                                                                                        case 649:
                                                                                            var v51 string = "hello"
                                                                                            _ = v51
                                                                                        case 128:
                                                                                            if ((v5 ^ 255) != 0) {
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    if false {
                                                                                                        var v52 uint32 = 2
                                                                                                        if (!false) {
                                                                                                            _ = len(v26)
                                                                                                        }
                                                                                                        v52 -= v52
                                                                                                        _ = v52
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    if false {
                                                                                                        v53 := GenS3{F0: 4294967295}
                                                                                                        if false {
                                                                                                            var v54 int64 = -9223372036854775808
                                                                                                            var v55 int32 = -1
                                                                                                            switch "\n\t" {
                                                                                                            case "\n\t":
                                                                                                                if len(v26) >= 2 {
                                                                                                                    v26 = v26[1:len(v26)]
                                                                                                                }
                                                                                                            default:
                                                                                                                switch fi24 {
                                                                                                                case 981:
                                                                                                                    if len(v26) > 0 {
                                                                                                                        v26[0] = fi21
                                                                                                                    }
                                                                                                                case 0:
                                                                                                                    fmt.Println(v55, v27)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v54
                                                                                                            _ = v55
                                                                                                        }
                                                                                                        _ = v53
                                                                                                    }
                                                                                                default:
                                                                                                    var v56 uint16 = (^(v29))
                                                                                                    var v57 int = fi24
                                                                                                    _ = v56
                                                                                                    _ = v57
                                                                                                }
                                                                                                fmt.Println(fi24, fi21, v29)
                                                                                                switch (fi24 | (fi21 << ((uint(v30)) & 63))) {
                                                                                                case -2147483648:
                                                                                                    if (-657 < -982) {
                                                                                                        fmt.Println(v5, fi24, fi21)
                                                                                                        v29 += 65535
                                                                                                    } else {
                                                                                                        v30 = (v29 >> (((v27 & (v43 ^ 0))) & 63))
                                                                                                    }
                                                                                                case 502:
                                                                                                    fmt.Println(fi21, fi21)
                                                                                                    v5 -= 72
                                                                                                    if (-1 < -59) {
                                                                                                        fmt.Println(v43, fi21, v43)
                                                                                                        var v58 uint32 = 42
                                                                                                        _ = v58
                                                                                                    }
                                                                                                case 32767:
                                                                                                    v59 := GenS6{F0: 102, F1: (v43 + (v43 + (^(v43)))), F2: 147}
                                                                                                    _ = v59
                                                                                                default:
                                                                                                    switch ("a" + ("\n\t" + ("\n\t" + ("a" + ("foo bar" + "\n\t"))))) {
                                                                                                    case "foo bar":
                                                                                                        v5 = 0
                                                                                                    default:
                                                                                                        v30 += v29
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v27)
                                                                                            } else {
                                                                                                var v60 bool = (!(!((42 >= 2147483647) && (v5 == uint8(v30)))))
                                                                                                var v61 int8 = 22
                                                                                                _ = v60
                                                                                                _ = v61
                                                                                            }
                                                                                        case 4294967295:
                                                                                            v62 := GenS7{F0: 1, F1: fi21, F2: 0}
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v63 := GenS8{F0: uint8(v43), F1: false, F2: -128}
                                                                                _ = v63
                                                                            }
                                                                            switch 229 {
                                                                            case 1:
                                                                                v64 := GenS7{F0: 0, F1: fi21, F2: -299}
                                                                                _ = v64
                                                                            case 690:
                                                                                if true {
                                                                                    var v65 float64 = 1.5
                                                                                    switch "" {
                                                                                    case "":
                                                                                        if true {
                                                                                            v29 += v29
                                                                                        }
                                                                                        var v66 uint = (v27 % (((^(v27))) | 1))
                                                                                        _ = v66
                                                                                    }
                                                                                    v30 -= (v30 / (((^(v30))) | 1))
                                                                                    _ = v65
                                                                                } else {
                                                                                    v67 := GenS2{F0: uint16(v43), F1: uint64(v5), F2: v43}
                                                                                    _ = v67
                                                                                }
                                                                            case 450:
                                                                                switch (fi21 >> (((v43 * 517)) & 63)) {
                                                                                case -630:
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v68 float64 = float64(((0.0 - 1.5) / 0.1))
                                                                                        _ = v68
                                                                                    case "foo bar":
                                                                                        v43 = (v27 << (((^(v43))) & 63))
                                                                                    default:
                                                                                        v43 = (v27 << ((v27) & 63))
                                                                                    }
                                                                                case 2147483647:
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        switch ("hello" + "a") {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                var v69 int = (fi24 % ((fi24) | 1))
                                                                                                v70 := GenS8{F0: (v5 & (v5 % (((v5 / ((204) | 1))) | 1))), F1: false, F2: -1}
                                                                                                _ = v69
                                                                                                _ = v70
                                                                                            }
                                                                                        default:
                                                                                            if false {
                                                                                                v27 -= (v43 << ((v27) & 63))
                                                                                                v71 := GenS6{F0: 100, F1: v27, F2: (^(v5))}
                                                                                                var v72 int8 = -1
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            } else {
                                                                                                v73 := GenS3{F0: 4294967295}
                                                                                                v5 -= (v5 >> ((v27) & 63))
                                                                                                var v74 float32 = (1e10 / 0.1)
                                                                                                _ = v73
                                                                                                _ = v74
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v75 uint8 = uint8(v43)
                                                                                        _ = v75
                                                                                    default:
                                                                                        v26 = append(v26, fi21)
                                                                                    }
                                                                                default:
                                                                                    var v76 uint32 = 218
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            if ((-5 < 2147483647) || false) {
                                                                                if len(v26) >= 2 {
                                                                                    v26 = v26[1:len(v26)]
                                                                                }
                                                                                v77 := GenS3{F0: 42}
                                                                                v29 -= (v30 + (v29 | v29))
                                                                                _ = v77
                                                                            } else {
                                                                                v78 := GenS9{F0: 706, F1: 42}
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            if len(v26) >= 2 {
                                                                                v26 = v26[1:len(v26)]
                                                                            }
                                                                        }
                                                                        _ = v43
                                                                    }
                                                                case -2147483648:
                                                                    v29 = uint16(v27)
                                                                default:
                                                                    var v79 int = (-(fi21))
                                                                    _ = v79
                                                                }
                                                            default:
                                                                v26 = append(v26, fi21)
                                                            }
                                                            _ = v30
                                                        }
                                                    }
                                                    v80 := GenS6{F0: -1, F1: (^(v27)), F2: v5}
                                                    fmt.Println(fi21, fi24)
                                                    v27 = v27
                                                    _ = v29
                                                    _ = v80
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        v27 = (v27 + uint(v5))
                                                    }
                                                }
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                    }
                                case "foo bar":
                                    if false {
                                        if (!(v5 <= 128)) {
                                            var v81 bool = false
                                            v5 -= 0
                                            v82 := []int{fi24}
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v83 uint32 = 2147483648
                                            v83 -= v83
                                            v84 := []int{42, (fi24 % ((fi21) | 1))}
                                            v5 = (v5 << ((0) & 63))
                                            _ = v83
                                            _ = v84
                                        }
                                        v5 = (v5 >> ((0) & 63))
                                    }
                                case "a":
                                    fmt.Println(fi21)
                                default:
                                    if (1 != 4294967295) {
                                        var v85 uint = 65535
                                        v85 = (v85 ^ v85)
                                        v85 = (v85 + uint(v5))
                                        v86 := GenS10{F0: 1, F1: 2147483647, F2: 32767}
                                        _ = v85
                                        _ = v86
                                    }
                                }
                                var v87 string = "hello"
                                _ = fi24
                                _ = v87
                            }
                            var v88 bool = true
                            _ = v88
                        } else {
                            for fi89 := 0; fi89 < 1; fi89++ {
                                v90 := []int{fi89, (fi89 % (((fi21 % (((^(fi21))) | 1))) | 1)), fi21}
                                _ = fi89
                                _ = v90
                            }
                            v91 := GenS2{F0: 0, F1: 939, F2: 42}
                            v92 := []int{(-(fi21)), (-(fi21)), fi21}
                            _ = v91
                            _ = v92
                        }
                    case "a":
                        v93 := GenS2{F0: 0, F1: 171, F2: 255}
                        _ = v93
                    case "\n\t":
                        if true {
                            fmt.Println(v5)
                            switch (fi21 % (((fi21 * fi21)) | 1)) {
                            case 9223372036854775807:
                                switch "\n\t" {
                                case "hello":
                                    var v94 uint32 = 748
                                    _ = v94
                                case "":
                                    v5 = v5
                                default:
                                    v95 := GenS11{F0: uint64(v5), F1: 9223372036854775807}
                                    _ = v95
                                }
                            case -61:
                                v5 += v5
                            case -854:
                                v96 := []int{}
                                _ = v96
                            default:
                                v5 -= (v5 - v5)
                            }
                            v97 := []int{(fi21 % ((-2147483648) | 1)), (-(fi21)), (fi21 >> ((65535) & 63)), fi21}
                            _ = v97
                        }
                    }
                    if (42 < 100) {
                        var v98 uint = 938
                        var v99 int = fi21
                        _ = v98
                        _ = v99
                    }
                    _ = fi21
                    _ = v22
                }
            case "\n\t":
                if false {
                    for fi100 := 0; fi100 < 17; fi100++ {
                        v5 = 42
                        v5 -= (v5 / ((v5) | 1))
                        _ = fi100
                    }
                    if false {
                        for fi101 := 0; fi101 < 14; fi101++ {
                            for fi102 := 0; fi102 < 6; fi102++ {
                                v5 -= 255
                                switch fi102 {
                                case -88:
                                    fmt.Println(v5, v5, fi102)
                                case 32767:
                                    v103 := []int{32767, -1, -29, (fi101 ^ (^(fi102)))}
                                    _ = v103
                                case 1:
                                    v104 := GenS12{F0: 349}
                                    _ = v104
                                default:
                                    fmt.Println(fi101, v5, fi102)
                                }
                                _ = fi102
                            }
                            if true {
                                if (!true) {
                                    var v105 int32 = 127
                                    fmt.Println(v105, v105, fi101)
                                    v106 := []int{(fi101 | -769)}
                                    _ = v105
                                    _ = v106
                                } else {
                                    fmt.Println(fi101, fi101)
                                }
                                var v107 int64 = 769
                                _ = v107
                            } else {
                                v5 = v5
                            }
                            v5 += 128
                            _ = fi101
                        }
                        var v108 string = "foo bar"
                        _ = v108
                    } else {
                        v109 := []int{-9223372036854775808, -460}
                        fmt.Println(v5)
                        _ = v109
                    }
                    v110 := GenS2{F0: 711, F1: 42, F2: 906}
                    _ = v110
                } else {
                    if false {
                        if false {
                            var v111 uint16 = 1
                            v111 = (v111 + v111)
                            v112 := []int{int(v5), 32767}
                            _ = v111
                            _ = v112
                        }
                    }
                    fmt.Println(v5)
                }
            default:
                for fi113 := 0; fi113 < 18; fi113++ {
                    v114 := []int{fi113, (^(fi113))}
                    _ = fi113
                    _ = v114
                }
            }
        } else {
            fmt.Println(v5)
            v5 += (v5 - v5)
            var v115 bool = (!false)
            _ = v115
        }
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -1
    v3 := []string{"hello", "hello"}
    v4 := make(map[int]uint32)
    _ = len(v3)
    if false {
        if false {
            fmt.Println(v2)
            var v116 int16 = -914
            _ = len(v3)
            _ = v116
        } else {
            fmt.Println(v2)
        }
        v117 := GenS4{F0: -1, F1: 288}
        v2 = 42
        _ = v117
    }
    for fi118 := 0; fi118 < 6; fi118++ {
        _ = len(v4)
        if true {
            for fi119 := 0; fi119 < 17; fi119++ {
                if true {
                    if false {
                        fmt.Println(v2)
                    } else {
                        switch "hello" {
                        case "foo bar":
                            var v120 int8 = 42
                            _ = v120
                        }
                        switch "a" {
                        case "":
                            var v121 uint32 = 519
                            _ = v121
                        case "\n\t":
                            v2 += 42
                        case "hello":
                            if true {
                                var v122 uint32 = 1
                                fmt.Println(v122, v122)
                                _ = v122
                            }
                        }
                    }
                    fmt.Println(fi119)
                } else {
                    v2 = 834
                }
                _ = fi119
            }
            v2 -= (v2 + (v2 | (^(v2))))
            for fi123 := 0; fi123 < 12; fi123++ {
                v2 = (v2 - (-(v2)))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch fi123 {
                case 127:
                    delete(v4, (fi118 << ((485) & 63)))
                }
                _ = fi123
            }
        } else {
            for fi124 := 0; fi124 < 20; fi124++ {
                switch "foo bar" {
                case "\n\t":
                    v2 -= (v2 >> ((4294967295) & 63))
                }
                fmt.Println(fi124)
                _ = fi124
            }
            switch (("a" + "") + "") {
            case "foo bar":
                for fi125 := 0; fi125 < 4; fi125++ {
                    v126 := v1["hello"]
                    if false {
                        v2 += v2
                        v2 = v2
                        v127 := GenS13{F0: (("" + "") + "a"), F1: 9223372036854775807}
                        _ = v127
                    } else {
                        var v128 float32 = -2.25
                        var v129 string = ""
                        _ = len(v1)
                        _ = v128
                        _ = v129
                    }
                    fmt.Println(fi118, fi118, fi118)
                    _ = fi125
                    _ = v126
                }
            case "hello":
                var v130 int16 = 32767
                _ = v130
            }
            var v131 uint16 = 190
            _ = v131
        }
        v4[569] = 0
        _ = fi118
    }
    v2 -= -1
    v2 += 2147483647
    if len(v3) > 0 {
        v3[0] = ((("hello" + "\n\t") + "hello") + "")
    }
    var v132 string = ((("hello" + ("\n\t" + "\n\t")) + "foo bar") + "foo bar")
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v132
    _ = v2
    _ = v3
    _ = v4
}
