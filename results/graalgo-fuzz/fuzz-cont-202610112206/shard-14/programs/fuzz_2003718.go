package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 string; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 uint64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int8 }
func genFn0() bool {
    var v5 int16 = 42
    v6 := GenS1{F0: 0}
    var v7 uint = 65535
    v8 := []int{754, 42}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return ("hello" <= (("" + "a") + ("\n\t" + "hello")))
}

func main() {
    var v1 int32 = -343
    v2 := make(map[uint32]int)
    var v3 uint64 = 2
    var v4 uint16 = 42
    v9 := v2[42]
    v3 = 2
    v9 += (v9 / ((v9) | 1))
    v10 := []int{}
    v3 = (v3 >> ((18446744073709551615) & 63))
    var v11 uint8 = 0
    v12 := GenS1{F0: 2}
    v13 := GenS1{F0: v3}
    if false {
        v11 = uint8(v3)
        for fi14 := 0; fi14 < 15; fi14++ {
            v9 = 2147483647
            if len(v10) > 0 {
                v10[0] = (fi14 & fi14)
            }
            _ = fi14
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int32 = v1
            var v17 uint8 = v11
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v18 int16 = 0
    if (!true) {
        if false {
            switch "" {
            case "":
                var v19 string = "foo bar"
                _ = v19
            case "a":
                switch (v9 * v9) {
                case -1:
                    v11 = (v11 + 42)
                }
            case "hello":
                switch (v9 % (((v9 ^ -1)) | 1)) {
                case 913:
                    v20 := GenS2{F0: ""}
                    _ = v20
                case 0:
                    v21 := GenS3{F0: ((1.5 * (float32(1e10) / ((0.1 / ((float32((float64(-2.25) - 3.14159)) - float32(float64(3.14159))) - 0.1)) / 1.0))) <= -2.25)}
                    _ = v21
                default:
                    switch (v9 / (((v9 | (v9 >> ((42) & 63)))) | 1)) {
                    case 2147483647:
                        v9 -= v9
                    case 32767:
                        if true {
                            v22 := GenS2{F0: "hello"}
                            _ = v22
                        } else {
                            switch v9 {
                            case 9223372036854775807:
                                v11 -= v11
                            case -73:
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    v24 := GenS2{F0: ("hello" + (("\n\t" + "") + "\n\t"))}
                                    if (!("" <= "\n\t")) {
                                        fmt.Println(v4, v9, v11)
                                    } else {
                                        switch "\n\t" {
                                        case "\n\t":
                                            v11 = v11
                                        case "foo bar":
                                            v4 = (v4 | v4)
                                        default:
                                            v9 = (v9 + fi23)
                                        }
                                        for fi25 := 0; fi25 < 9; fi25++ {
                                            v9 = v9
                                            v1 -= 886
                                            _ = fi25
                                        }
                                    }
                                    if ((!(0 <= 9223372036854775807)) && (42 >= 2147483647)) {
                                        fmt.Println(v3)
                                    } else {
                                        v10 = append(v10, (fi23 / ((v9) | 1)))
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            case -974:
                                if len(v10) > 0 {
                                    v10[0] = 0
                                }
                            }
                            v3 += (^(v3))
                            fmt.Println(v3, v1)
                        }
                    case 313:
                        v4 = 2
                    default:
                        var v26 string = ((((("hello" + "") + "foo bar") + "\n\t") + "foo bar") + "\n\t")
                        _ = v26
                    }
                }
            }
            if len(v10) > 0 {
                v10[0] = (v9 * v9)
            }
            v3 = v3
        }
        v18 += (v18 & 0)
    } else {
        switch "foo bar" {
        case "a":
            if (v9 <= v9) {
                v3 = (v3 ^ v3)
                fmt.Println(v4, v3)
            } else {
                if false {
                    v27 := GenS1{F0: v3}
                    for fi28 := 0; fi28 < 15; fi28++ {
                        var v29 int8 = 100
                        if false {
                            for fi30 := 0; fi30 < 2; fi30++ {
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                delete(v2, 42)
                                _ = fi30
                            }
                            switch "" {
                            case "\n\t":
                                var v31 float32 = ((3.14159 - 0.0) - (0.1 - (0.0 + 0.1)))
                                _ = v31
                            default:
                                if (42 > v1) {
                                    var v32 int8 = v29
                                    _ = v32
                                }
                            }
                        } else {
                            _ = len(v10)
                        }
                        _ = fi28
                        _ = v29
                    }
                    if true {
                        _ = len(v10)
                        var v33 float32 = float32(-2.25)
                        _ = v33
                    }
                    _ = v27
                }
            }
        case "":
            fmt.Println(v9, v9, v11)
        case "hello":
            if false {
                v34 := GenS4{F0: 42, F1: ("" + "foo bar"), F2: -668}
                v35 := GenS5{F0: false}
                switch "hello" {
                case "foo bar":
                    fmt.Println(v11, v9, v9)
                case "":
                    switch (v9 ^ 2147483647) {
                    case 234:
                        v36 := GenS3{F0: (v9 == 9223372036854775807)}
                        _ = v36
                    case 9223372036854775807:
                        if true {
                            v18 += (v18 - v18)
                        }
                    case 746:
                        if false {
                            if (182 < 4294967295) {
                                var v37 int16 = (v18 << ((255) & 63))
                                v2[1] = (v9 << ((737) & 63))
                                if len(v10) >= 2 {
                                    v10 = v10[1:len(v10)]
                                }
                                _ = v37
                            } else {
                                v1 += v1
                                fmt.Println(v4)
                            }
                        } else {
                            var v38 uint32 = 4294967295
                            fmt.Println(v38, v38, v3)
                            if false {
                                switch ("\n\t" + "\n\t") {
                                case "foo bar":
                                    fmt.Println(v18, v38, v11)
                                case "\n\t":
                                    switch (v9 & v9) {
                                    case 0:
                                        for fi39 := 0; fi39 < 5; fi39++ {
                                            for fi40 := 0; fi40 < 6; fi40++ {
                                                switch (fi39 - 661) {
                                                case -9223372036854775808:
                                                    switch (v9 / (((v9 / (((v9 / ((fi39) | 1))) | 1))) | 1)) {
                                                    case 536:
                                                        v41 := GenS3{F0: false}
                                                        _ = v41
                                                    case 4294967295:
                                                        fmt.Println(fi39, v4)
                                                    case 415:
                                                        switch fi39 {
                                                        case -9223372036854775808:
                                                            if ((1.5 - (3.14159 / (((1.5 / -2.25) - float64(3.14159)) + 1e10))) < (float64(((3.14159 - float32(1.5)) / 1e10)) / (1e10 / 0.1))) {
                                                                var v42 float32 = 0.1
                                                                var v43 int8 = 1
                                                                _ = v42
                                                                _ = v43
                                                            } else {
                                                                var v44 int16 = int16(v38)
                                                                v9 += 4294967295
                                                                _ = v44
                                                            }
                                                        case 472:
                                                            switch v9 {
                                                            case -204:
                                                                v4 = 42
                                                            case -480:
                                                                v4 = 2
                                                            default:
                                                                fmt.Println(v11, v11, v38)
                                                            }
                                                        }
                                                        v45 := GenS2{F0: "foo bar"}
                                                        _ = v45
                                                    default:
                                                        v1 = v1
                                                    }
                                                }
                                                _ = fi40
                                            }
                                            v1 -= (-(v1))
                                            var v46 int16 = (v18 << ((2) & 63))
                                            _ = fi39
                                            _ = v46
                                        }
                                    }
                                case "a":
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                default:
                                    v9 = 32768
                                }
                            }
                            _ = v38
                        }
                    }
                case "a":
                    v3 += (v3 << ((4294967295) & 63))
                }
                _ = v34
                _ = v35
            }
        }
        switch "foo bar" {
        case "foo bar":
            v47 := GenS1{F0: (v3 + 857)}
            _ = v47
        default:
            if true {
                if false {
                    if ((uint16(v11) < v4) && false) {
                        v9 = (v9 * (v9 ^ 0))
                        for fi48 := 0; fi48 < 4; fi48++ {
                            if ("\n\t" >= "") {
                                var v49 uint64 = (v3 % ((v3) | 1))
                                _ = v49
                            }
                            v50 := GenS2{F0: ("hello" + ((("a" + ("a" + ((("foo bar" + "foo bar") + ("\n\t" + "")) + ""))) + ("foo bar" + "\n\t")) + ((("hello" + ("" + ("" + "a"))) + "a") + "foo bar")))}
                            var v51 uint16 = (v4 / ((v4) | 1))
                            _ = fi48
                            _ = v50
                            _ = v51
                        }
                    } else {
                        for fi52 := 0; fi52 < 10; fi52++ {
                            if true {
                                fmt.Println(v3, v1, v18)
                                if true {
                                    v53 := v2[535]
                                    v11 += v11
                                    for fi54 := 0; fi54 < 13; fi54++ {
                                        var v55 uint64 = v3
                                        var v56 uint32 = 707
                                        if false {
                                            switch -322 {
                                            case 0:
                                                var v57 int = 0
                                                _ = v57
                                            case 9223372036854775807:
                                                var v58 uint64 = 2
                                                _ = v58
                                            }
                                            v55 -= 612
                                            var v59 uint8 = (v11 ^ v11)
                                            if true {
                                                if len(v10) > 0 {
                                                    v10[0] = (fi54 & (v9 % ((-938) | 1)))
                                                }
                                                var v60 uint16 = (v4 ^ v4)
                                                _ = v60
                                            }
                                            _ = v59
                                        } else {
                                            if (!true) {
                                                v9 += 963
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                switch (v9 << ((353) & 63)) {
                                                case 2147483647:
                                                    if (true && (int8(v4) == -128)) {
                                                        v3 -= (v55 ^ 0)
                                                        _ = len(v10)
                                                        if true {
                                                            switch int(v18) {
                                                            case -2147483648:
                                                                fmt.Println(v9, v4, v18)
                                                            case 32767:
                                                                var v61 uint16 = v4
                                                                _ = v61
                                                            default:
                                                                fmt.Println(v4, v18, v55)
                                                            }
                                                            switch (("foo bar" + "a") + "hello") {
                                                            case "\n\t":
                                                                v55 += 923
                                                            case "foo bar":
                                                                fmt.Println(fi54)
                                                            case "a":
                                                                v11 = 79
                                                            }
                                                        }
                                                    } else {
                                                        var v62 uint8 = (v11 >> ((4294967295) & 63))
                                                        _ = v62
                                                    }
                                                default:
                                                    if false {
                                                        v56 -= v56
                                                        if true {
                                                            fmt.Println(v3, v55)
                                                            v63 := GenS2{F0: "foo bar"}
                                                            var v64 uint = 65535
                                                            if false {
                                                                if true {
                                                                    v53 = 9223372036854775807
                                                                } else {
                                                                    v4 = (v4 - 593)
                                                                }
                                                                v4 = v4
                                                            }
                                                            _ = v63
                                                            _ = v64
                                                        } else {
                                                            var v65 int16 = (v18 ^ (v18 + v18))
                                                            _ = v65
                                                        }
                                                        var v66 string = ("" + ("" + ("foo bar" + "\n\t")))
                                                        v67 := GenS4{F0: 42, F1: ("\n\t" + v66), F2: v1}
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                if (0 == 287) {
                                                    v18 += (v18 >> ((1) & 63))
                                                    if false {
                                                        switch ("a" + "") {
                                                        case "a":
                                                            switch (-(v9)) {
                                                            case 32768:
                                                                if (65535 <= 333) {
                                                                    v68 := GenS6{F0: (v56 << ((78) & 63)), F1: 0}
                                                                    if len(v10) >= 2 {
                                                                        v10 = v10[1:len(v10)]
                                                                    }
                                                                    var v69 uint = 42
                                                                    v70 := GenS6{F0: (^(v56)), F1: -2147483648}
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v55 = v3
                                                                    case "a":
                                                                        v71 := GenS4{F0: 1, F1: "", F2: 2147483647}
                                                                        _ = v71
                                                                    default:
                                                                        var v72 uint8 = v11
                                                                        _ = v72
                                                                    }
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    v9 += fi54
                                                                    if false {
                                                                        v73 := GenS6{F0: (v56 - v56), F1: 42}
                                                                        _ = v73
                                                                    } else {
                                                                        fmt.Println(fi54, fi52)
                                                                    }
                                                                    v74 := GenS7{F0: 0}
                                                                    var v75 float64 = -2.25
                                                                    if false {
                                                                        var v76 string = "foo bar"
                                                                        var v77 string = (v76 + v76)
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    var v78 string = "a"
                                                                    _ = v74
                                                                    _ = v75
                                                                    _ = v78
                                                                }
                                                            case 128:
                                                                fmt.Println(v11)
                                                            case 140:
                                                                _ = len(v10)
                                                            }
                                                        default:
                                                            fmt.Println(v4)
                                                        }
                                                        var v79 int8 = -66
                                                        v79 = (v79 << ((42) & 63))
                                                        _ = v79
                                                    }
                                                } else {
                                                    _ = len(v10)
                                                }
                                            }
                                            if (-1 >= 100) {
                                                if true {
                                                    v11 = 0
                                                    _ = len(v10)
                                                }
                                                v11 -= (v11 & v11)
                                            }
                                        }
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v53
                                }
                            } else {
                                switch (v9 ^ v9) {
                                case -2147483648:
                                    if (v3 <= 692) {
                                        switch (v9 % ((v9) | 1)) {
                                        case -9223372036854775808:
                                            v80 := GenS1{F0: (v3 ^ v3)}
                                            _ = v80
                                        case -177:
                                            for fi81 := 0; fi81 < 7; fi81++ {
                                                v82 := GenS8{F0: v3, F1: v3}
                                                switch ("hello" + "\n\t") {
                                                case "hello":
                                                    fmt.Println(v4)
                                                case "":
                                                    var v83 uint64 = v3
                                                    _ = v83
                                                }
                                                _ = fi81
                                                _ = v82
                                            }
                                        }
                                        v9 = int(v4)
                                    }
                                case -890:
                                    switch fi52 {
                                    case 127:
                                        v84 := GenS2{F0: ("a" + "\n\t")}
                                        _ = v84
                                    case -9223372036854775808:
                                        v3 = v3
                                    case 544:
                                        for fi85 := 0; fi85 < 4; fi85++ {
                                            if false {
                                                v3 -= 514
                                            }
                                            _ = fi85
                                        }
                                    default:
                                        v86 := GenS2{F0: "a"}
                                        _ = v86
                                    }
                                }
                                v1 -= 2147483647
                            }
                            _ = fi52
                        }
                        v18 = 42
                    }
                    fmt.Println(v9, v4, v1)
                    for fi87 := 0; fi87 < 8; fi87++ {
                        var v88 uint32 = 2
                        _ = fi87
                        _ = v88
                    }
                } else {
                    _ = len(v10)
                    for fi89 := 0; fi89 < 15; fi89++ {
                        var v90 uint = 2
                        _ = fi89
                        _ = v90
                    }
                }
                for fi91 := 0; fi91 < 17; fi91++ {
                    switch "hello" {
                    case "a":
                        v92 := GenS5{F0: false}
                        _ = v92
                    default:
                        v93 := GenS9{F0: (-(fi91))}
                        _ = v93
                    }
                    _ = fi91
                }
                var v94 uint8 = (^(v11))
                _ = v94
            } else {
                fmt.Println(v1, v3)
                if true {
                    fmt.Println(v9)
                    v11 -= (v11 / (((v11 ^ 0)) | 1))
                    for fi95 := 0; fi95 < 15; fi95++ {
                        v10 = append(v10, (v9 & (^(fi95))))
                        _ = fi95
                    }
                } else {
                    switch (^(v9)) {
                    case 238:
                        if (v9 != 32767) {
                            v96 := GenS10{F0: (1e10 != ((1e10 * (3.14159 + 0.1)) - ((0.0 * 0.0) / -2.25)))}
                            for fi97 := 0; fi97 < 9; fi97++ {
                                v18 = (v18 | v18)
                                _ = fi97
                            }
                            for fi98 := 0; fi98 < 7; fi98++ {
                                v99 := GenS6{F0: 1, F1: 1}
                                v100 := GenS2{F0: "a"}
                                switch (v9 >> ((513) & 63)) {
                                case 0:
                                    switch 40 {
                                    case 42:
                                        switch (fi98 | fi98) {
                                        case -9223372036854775808:
                                            if (true || (!false)) {
                                                var v101 float32 = ((-2.25 * 1e10) / (((-2.25 / 0.1) - ((0.1 / 0.1) / (1e10 / 1.0))) / 1.5))
                                                switch (^(fi98)) {
                                                case -944:
                                                    for fi102 := 0; fi102 < 4; fi102++ {
                                                        v9 = 127
                                                        v101 = v101
                                                        _ = fi102
                                                    }
                                                default:
                                                    switch fi98 {
                                                    case -1:
                                                        var v103 float32 = (v101 + v101)
                                                        _ = v103
                                                    case 128:
                                                        var v104 string = "\n\t"
                                                        _ = v104
                                                    default:
                                                        fmt.Println(v101, v1, v11)
                                                    }
                                                }
                                                if true {
                                                    var v105 int8 = -1
                                                    v11 = (v11 % ((255) | 1))
                                                    _ = v105
                                                }
                                                _ = v101
                                            } else {
                                                for fi106 := 0; fi106 < 18; fi106++ {
                                                    v9 -= (v9 ^ (v9 | 32768))
                                                    _ = fi106
                                                }
                                                v107 := GenS6{F0: 2147483648, F1: (fi98 ^ v9)}
                                                _ = v107
                                            }
                                        }
                                    case 32767:
                                        v108 := GenS8{F0: 0, F1: 1}
                                        _ = v108
                                    default:
                                        switch v9 {
                                        case 42:
                                            fmt.Println(v9)
                                        }
                                    }
                                default:
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v18)
                                    case "foo bar":
                                        v109 := GenS2{F0: ("a" + "")}
                                        _ = v109
                                    case "a":
                                        if true {
                                            for fi110 := 0; fi110 < 1; fi110++ {
                                                v111 := GenS1{F0: (v3 + v3)}
                                                var v112 float32 = (0.1 / 0.1)
                                                if false {
                                                    v3 -= v3
                                                } else {
                                                    if (false || ((!true) || false)) {
                                                        if (!(-1 >= 1)) {
                                                            v113 := v2[1]
                                                            _ = v113
                                                        } else {
                                                            var v114 uint = 670
                                                            _ = v114
                                                        }
                                                        var v115 uint = 151
                                                        _ = v115
                                                    }
                                                    if (false || true) {
                                                        if len(v10) > 0 {
                                                            v10[0] = (fi98 % ((9223372036854775807) | 1))
                                                        }
                                                        v10 = append(v10, (fi110 >> ((862) & 63)))
                                                    }
                                                }
                                                _ = fi110
                                                _ = v111
                                                _ = v112
                                            }
                                            switch (v9 & v9) {
                                            case 4294967295:
                                                var v116 int8 = 100
                                                _ = v116
                                            }
                                        }
                                    }
                                }
                                _ = fi98
                                _ = v100
                                _ = v99
                            }
                            _ = v96
                        }
                    case 127:
                        if false {
                            var v117 uint8 = v11
                            _ = v117
                        }
                    }
                    for fi118 := 0; fi118 < 19; fi118++ {
                        v119 := GenS10{F0: (!true)}
                        _ = len(v10)
                        _ = fi118
                        _ = v119
                    }
                }
                switch v9 {
                case 4294967295:
                    var v120 uint64 = (v3 | 798)
                    _ = v120
                case 32768:
                    v4 -= (v4 | 706)
                }
            }
        }
        switch -422 {
        case 4294967295:
            for fi121 := 0; fi121 < 14; fi121++ {
                switch 2147483647 {
                case 127:
                    fmt.Println(v18, v4)
                case -9223372036854775808:
                    _ = len(v10)
                case 2147483647:
                    if (true && true) {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        v122 := GenS8{F0: 18446744073709551615, F1: v3}
                        _ = v122
                    } else {
                        if (v3 <= v3) {
                            if len(v10) > 0 {
                                v10[0] = (^(fi121))
                            }
                        }
                        var v123 int = fi121
                        if false {
                            switch ("\n\t" + "") {
                            case "foo bar":
                                var v124 uint64 = (v3 | (v3 >> ((255) & 63)))
                                _ = v124
                            case "\n\t":
                                var v125 int16 = int16(v1)
                                _ = v125
                            }
                        } else {
                            v126 := GenS11{F0: 100}
                            _ = v126
                        }
                        _ = v123
                    }
                }
                _ = fi121
            }
        case -9223372036854775808:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
