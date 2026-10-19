package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint8; F2 int }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 int32; F1 uint32; F2 int16 }
type GenS8 struct { F0 uint32; F1 int8; F2 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int8; F1 uint32 }
type GenS13 struct { F0 bool; F1 bool; F2 int64 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 uint }
func genFn0(p0 bool, p1 uint32, p2 int32) int {
    fmt.Println(p2, p2)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{(^(fi5))}
        switch fi5 {
        case -766:
            p0 = p0
        case 0:
            for fi7 := 0; fi7 < 4; fi7++ {
                v8 := GenS1{F0: 100, F1: -545}
                switch (fi5 & 128) {
                case 32767:
                    if p0 {
                        p1 = (^(p1))
                        p0 = ("hello" == "foo bar")
                    }
                case 160:
                    fmt.Println(fi5)
                }
                v9 := GenS1{F0: 53, F1: 547}
                _ = fi7
                _ = v8
                _ = v9
            }
        case 32768:
            p0 = true
        }
        _ = fi5
        _ = v6
    }
    v10 := GenS2{F0: 34}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -800
}

func main() {
    var v1 uint64 = 0
    var v2 string = ""
    v3 := []uint32{1, 4294967295, 456, 65535}
    var v4 int16 = -471
    fmt.Println(v1)
    if (true && true) {
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 = append(v3, 4294967295)
            if (true || false) {
                switch fi11 {
                case -812:
                    var v12 uint32 = 956
                    _ = v12
                }
                v13 := GenS1{F0: 1, F1: -54}
                _ = v13
            } else {
                v1 -= v1
            }
            if true {
                if false {
                    switch (fi11 << ((4294967295) & 63)) {
                    case -9223372036854775808:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 854:
                        var v14 string = v2
                        _ = v14
                    case -835:
                        if false {
                            var v15 uint16 = 139
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v15
                        } else {
                            v16 := GenS2{F0: 4294967295}
                            _ = v16
                        }
                    default:
                        if (!(-2147483648 >= -2147483648)) {
                            var v17 uint32 = 99
                            _ = v17
                        } else {
                            v4 -= (v4 - -1)
                            v2 = v2
                            var v18 int32 = 2147483647
                            _ = v18
                        }
                    }
                }
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int32 = 1
                    if (false || false) {
                        if (574 > 42) {
                            fmt.Println(fi11, v2, v1)
                        }
                        fmt.Println(fi11, v2)
                        if (!true) {
                            if (!true) {
                                v1 = 2
                                v1 += v1
                                if ((v2 + ("" + v2)) < (v2 + (("foo bar" + ("a" + v2)) + (v2 + "hello")))) {
                                    var v21 int32 = -174
                                    _ = v21
                                }
                            } else {
                                if true {
                                    v22 := GenS2{F0: v1}
                                    switch (v2 + (((v2 + v2) + ((v2 + (v2 + "foo bar")) + v2)) + ((("" + (v2 + (v2 + (v2 + "")))) + v2) + v2))) {
                                    case "":
                                        fmt.Println(v4)
                                    case "a":
                                        if (false && ((!true) && true)) {
                                            if ((127 <= 42) || (-725 >= v20)) {
                                                var v23 int32 = 1
                                                switch "a" {
                                                case "a":
                                                    v24 := GenS3{F0: 6, F1: 255, F2: genFn0((!true), 2147483647, v23)}
                                                    _ = v24
                                                case "\n\t":
                                                    switch (((v2 + v2) + "a") + "a") {
                                                    case "":
                                                        switch (fi19 / (((fi19 | genFn0(false, 0, v20))) | 1)) {
                                                        case 128:
                                                            var v25 uint16 = 1
                                                            _ = v25
                                                        case -2147483648:
                                                            v23 += 537
                                                        case 4294967295:
                                                            fmt.Println(v2, fi11, v20)
                                                        }
                                                    case "\n\t":
                                                        var v26 int = 1
                                                        _ = v26
                                                    default:
                                                        fmt.Println(v1, v20, fi11)
                                                    }
                                                default:
                                                    v2 = ""
                                                }
                                                if false {
                                                    switch fi19 {
                                                    case -681:
                                                        switch v2 {
                                                        case "foo bar":
                                                            v4 = v4
                                                            v4 -= (v4 % ((v4) | 1))
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(fi19)
                                                    }
                                                    v27 := GenS2{F0: v1}
                                                    var v28 int8 = -121
                                                    var v29 uint = 65535
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                                _ = v23
                                            }
                                            if (!(true || (1e10 != (3.14159 - (0.0 - 0.0))))) {
                                                if false {
                                                    var v30 int16 = v4
                                                    var v31 int32 = v20
                                                    _ = v30
                                                    _ = v31
                                                }
                                            }
                                            v32 := GenS2{F0: v1}
                                            _ = v32
                                        }
                                    default:
                                        v33 := GenS1{F0: 100, F1: 42}
                                        if len(v3) > 0 {
                                            v3[0] = 2147483648
                                        }
                                        _ = v33
                                    }
                                    _ = v22
                                }
                            }
                        }
                    }
                    if ((v4 < (v4 | (v4 + v4))) || ((255 < 32767) && false)) {
                        v34 := GenS3{F0: -1, F1: 1, F2: -999}
                        _ = v34
                    } else {
                        fmt.Println(v4)
                        var v35 uint8 = 155
                        _ = v35
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                var v36 uint = 120
                var v37 float64 = 3.14159
                _ = v36
                _ = v37
            }
            _ = fi11
        }
        v2 = (v2 + "hello")
    } else {
        switch -1 {
        case 360:
            switch v2 {
            case "\n\t":
                if true {
                    if (!false) {
                        for fi38 := 0; fi38 < 1; fi38++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v39 int64 = 42
                            switch (fi38 | -2147483648) {
                            case 0:
                                for fi40 := 0; fi40 < 20; fi40++ {
                                    fmt.Println(fi38)
                                    _ = len(v3)
                                    if (1 >= v1) {
                                        v41 := GenS1{F0: 127, F1: (v39 & v39)}
                                        v42 := GenS4{F0: 255, F1: 113}
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi40
                                }
                            case 127:
                                v3 = append(v3, 42)
                            case -2147483648:
                                switch (^(fi38)) {
                                case 812:
                                    for fi43 := 0; fi43 < 16; fi43++ {
                                        if false {
                                            v44 := GenS2{F0: 553}
                                            fmt.Println(fi43, fi38, fi43)
                                            _ = v44
                                        }
                                        if false {
                                            var v45 uint32 = 4294967295
                                            _ = v45
                                        } else {
                                            v46 := GenS4{F0: 255, F1: 2}
                                            if (49 <= 83) {
                                                v47 := GenS5{F0: (!true), F1: 128}
                                                _ = v47
                                            } else {
                                                fmt.Println(v1, v1)
                                            }
                                            if (genFn0(false, 42, 1) > fi43) {
                                                var v48 int32 = 109
                                                if len(v3) > 0 {
                                                    v3[0] = 65535
                                                }
                                                _ = v48
                                            }
                                            _ = v46
                                        }
                                        var v49 int8 = 1
                                        _ = fi43
                                        _ = v49
                                    }
                                default:
                                    switch v2 {
                                    case "\n\t":
                                        v4 -= (v4 >> ((18446744073709551615) & 63))
                                    case "a":
                                        if (float64(3.14159) <= 1e10) {
                                            switch fi38 {
                                            case 42:
                                                for fi50 := 0; fi50 < 8; fi50++ {
                                                    var v51 float32 = 0.1
                                                    v2 = v2
                                                    var v52 float64 = 0.0
                                                    _ = fi50
                                                    _ = v51
                                                    _ = v52
                                                }
                                            case 32768:
                                                var v53 int16 = v4
                                                _ = v53
                                            default:
                                                v54 := GenS6{F0: "foo bar", F1: true}
                                                _ = v54
                                            }
                                            v1 = v1
                                        } else {
                                            if (-2.25 <= 0.1) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if (24 == 127) {
                                                    v55 := GenS6{F0: v2, F1: true}
                                                    _ = v55
                                                }
                                            }
                                            v56 := GenS2{F0: v1}
                                            _ = v56
                                        }
                                    case "":
                                        v57 := GenS2{F0: v1}
                                        _ = v57
                                    }
                                }
                            }
                            _ = fi38
                            _ = v39
                        }
                        v58 := GenS2{F0: 4294967295}
                        _ = v58
                    } else {
                        if true {
                            v59 := GenS5{F0: (false && ((121 != 1) || false)), F1: 127}
                            _ = v59
                        } else {
                            switch v2 {
                            case "a":
                                v3 = append(v3, 0)
                            default:
                                v60 := GenS2{F0: 1}
                                _ = v60
                            }
                            var v61 int = -9223372036854775808
                            _ = v61
                        }
                        v62 := GenS7{F0: 42, F1: 65535, F2: (v4 >> ((2) & 63))}
                        _ = v62
                    }
                    switch genFn0(false, 0, -60) {
                    case 42:
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                    case 32767:
                        var v63 float64 = (0.0 - -2.25)
                        _ = v63
                    case 127:
                        var v64 uint = 604
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 4; fi65++ {
                        if false {
                            var v66 float64 = (float64(1.5) + float64(1.5))
                            fmt.Println(v1, v4)
                            _ = v66
                        }
                        _ = fi65
                    }
                }
            case "a":
                if true {
                    v67 := GenS8{F0: 42, F1: 1, F2: 1}
                    if true {
                        v68 := GenS3{F0: 42, F1: 42, F2: genFn0((!(1 == -1)), 2147483647, 1)}
                        _ = v68
                    } else {
                        if (0 > 42) {
                            var v69 int64 = 9223372036854775807
                            v70 := GenS9{F0: 65535}
                            switch v2 {
                            case "\n\t":
                                v3 = append(v3, 42)
                            case "":
                                fmt.Println(v69, v4, v2)
                            case "a":
                                fmt.Println(v2)
                            }
                            _ = v69
                            _ = v70
                        }
                    }
                    _ = v67
                }
            default:
                if false {
                    switch ("foo bar" + "foo bar") {
                    case "hello":
                        var v71 uint32 = 2147483648
                        _ = v71
                    case "a":
                        v4 -= 1
                    default:
                        if ((v4 - v4) > v4) {
                            if (true && ((((((-2.25 / float64((3.14159 * 1e10))) + -2.25) - -2.25) - 0.1) / (((-2.25 - float64(3.14159)) + 1e10) / 0.1)) >= float64((-2.25 * 1.5)))) {
                                if false {
                                    var v72 bool = true
                                    for fi73 := 0; fi73 < 16; fi73++ {
                                        fmt.Println(fi73)
                                        _ = fi73
                                    }
                                    v2 = v2
                                    _ = v72
                                } else {
                                    v4 += (-(v4))
                                }
                                var v74 uint16 = 255
                                switch (("foo bar" + "hello") + v2) {
                                case "\n\t":
                                    var v75 uint = 4294967295
                                    _ = v75
                                default:
                                    var v76 int8 = 42
                                    _ = v76
                                }
                                _ = v74
                            }
                            _ = len(v3)
                        } else {
                            v4 -= (v4 >> ((1) & 63))
                            var v77 int32 = 0
                            _ = v77
                        }
                    }
                    for fi78 := 0; fi78 < 4; fi78++ {
                        switch 0 {
                        case 608:
                            for fi79 := 0; fi79 < 10; fi79++ {
                                var v80 string = (v2 + v2)
                                _ = fi79
                                _ = v80
                            }
                        case 1:
                            v3 = append(v3, 903)
                        case -2147483648:
                            fmt.Println(v4)
                        default:
                            v81 := GenS7{F0: -1, F1: 2, F2: v4}
                            _ = v81
                        }
                        _ = fi78
                    }
                } else {
                    v1 = (v1 | v1)
                }
            }
        case -860:
            if false {
                if (!(-9223372036854775808 < -980)) {
                    if (1.5 == (1.5 + (3.14159 + (float64(3.14159) - float64(1e10))))) {
                        switch 127 {
                        case -2147483648:
                            switch v2 {
                            case "foo bar":
                                v3 = append(v3, 1)
                            case "a":
                                v3 = append(v3, 4294967295)
                            }
                        default:
                            v82 := GenS10{F0: -72}
                            _ = v82
                        }
                    }
                    for fi83 := 0; fi83 < 1; fi83++ {
                        for fi84 := 0; fi84 < 14; fi84++ {
                            _ = len(v3)
                            _ = fi84
                        }
                        _ = fi83
                    }
                    if true {
                        fmt.Println(v1)
                        v2 = (v2 + v2)
                    }
                }
                v4 = 0
            }
        default:
            var v85 uint64 = 4294967295
            _ = v85
        }
        if (v1 > 210) {
            for fi86 := 0; fi86 < 10; fi86++ {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
                v87 := GenS11{F0: 32767}
                if false {
                    v1 -= v1
                    var v88 uint8 = 255
                    for fi89 := 0; fi89 < 8; fi89++ {
                        if (0 < (v88 - v88)) {
                            fmt.Println(v4)
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v90 := GenS12{F0: 42, F1: 25}
                                _ = v90
                            }
                        } else {
                            if true {
                                switch int(v88) {
                                case -2147483648:
                                    switch v2 {
                                    case "hello":
                                        if ((v4 >> ((42) & 63)) != 898) {
                                            var v91 string = (((v2 + v2) + v2) + "foo bar")
                                            v91 = v91
                                            if len(v3) > 0 {
                                                v3[0] = 1
                                            }
                                            _ = v91
                                        }
                                    case "\n\t":
                                        switch (fi89 << ((1) & 63)) {
                                        case 565:
                                            _ = len(v3)
                                        }
                                    }
                                case 128:
                                    var v92 float32 = float32((0.1 / (float64(float32(0.1)) / 0.1)))
                                    _ = v92
                                default:
                                    v1 -= (v1 + v1)
                                }
                                var v93 uint = 65535
                                v94 := GenS6{F0: v2, F1: (false || true)}
                                _ = v93
                                _ = v94
                            }
                            var v95 uint16 = 602
                            _ = v95
                        }
                        var v96 float32 = 1.5
                        _ = fi89
                        _ = v96
                    }
                    _ = v88
                } else {
                    switch "" {
                    case "":
                        if (42 > 1) {
                            var v97 uint8 = 0
                            v2 = v2
                            _ = v97
                        } else {
                            v98 := GenS2{F0: 18446744073709551615}
                            switch "" {
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = 4294967295
                                }
                            case "":
                                v1 = (v1 >> ((899) & 63))
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 42
                                }
                            }
                            _ = v98
                        }
                    case "foo bar":
                        v2 = v2
                    }
                    v99 := GenS13{F0: (true || (fi86 < fi86)), F1: (874 == 636), F2: 0}
                    var v100 uint = 4294967295
                    _ = v100
                    _ = v99
                }
                _ = fi86
                _ = v87
            }
        } else {
            var v101 uint64 = v1
            _ = len(v3)
            if (-1 == -1) {
                switch "\n\t" {
                case "a":
                    v1 -= v101
                }
                if false {
                    v4 -= -652
                    var v102 float64 = 1e10
                    _ = v102
                } else {
                    var v103 float64 = 0.1
                    switch -9223372036854775808 {
                    case 32767:
                        v4 += v4
                    case 9223372036854775807:
                        if true {
                            fmt.Println(v1)
                            fmt.Println(v1, v2)
                            for fi104 := 0; fi104 < 12; fi104++ {
                                var v105 uint8 = 114
                                fmt.Println(v101, v4)
                                v3 = append(v3, 4294967295)
                                _ = fi104
                                _ = v105
                            }
                        } else {
                            var v106 int32 = 1
                            v101 = v1
                            if false {
                                var v107 bool = (!(((float64(-2.25) + v103) + ((1e10 * v103) + 1e10)) >= v103))
                                var v108 uint64 = v1
                                for fi109 := 0; fi109 < 10; fi109++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v110 int8 = 108
                                    _ = fi109
                                    _ = v110
                                }
                                _ = v107
                                _ = v108
                            }
                            _ = v106
                        }
                    case 128:
                        v1 = v1
                    default:
                        _ = len(v3)
                    }
                    if ((float32(((float64(((1e10 / float32(float64(float32(0.1)))) * 1e10)) + 1e10) - float64(float32(v103)))) > 3.14159) && false) {
                        if (2 <= 664) {
                            v3 = append(v3, 2147483647)
                        }
                    } else {
                        for fi111 := 0; fi111 < 16; fi111++ {
                            var v112 bool = false
                            _ = fi111
                            _ = v112
                        }
                        var v113 uint64 = (^(v101))
                        for fi114 := 0; fi114 < 4; fi114++ {
                            _ = len(v3)
                            _ = fi114
                        }
                        _ = v113
                    }
                    _ = v103
                }
                fmt.Println(v101, v4)
            }
            _ = v101
        }
    }
    switch v2 {
    case "foo bar":
        switch -862 {
        case 0:
            switch v2 {
            case "":
                fmt.Println(v2, v4, v1)
            case "foo bar":
                for fi115 := 0; fi115 < 17; fi115++ {
                    v116 := GenS6{F0: v2, F1: ((fi115 * fi115) == genFn0((42 < 8), 2147483647, 0))}
                    for fi117 := 0; fi117 < 9; fi117++ {
                        var v118 int32 = -25
                        v119 := GenS14{F0: -942}
                        if (!true) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v120 := GenS12{F0: -127, F1: 811}
                            fmt.Println(fi115, fi115)
                            _ = v120
                        }
                        _ = fi117
                        _ = v118
                        _ = v119
                    }
                    for fi121 := 0; fi121 < 15; fi121++ {
                        switch v2 {
                        case "hello":
                            var v122 uint32 = 977
                            _ = v122
                        case "foo bar":
                            v1 -= v1
                        case "a":
                            var v123 int32 = 9
                            _ = v123
                        default:
                            v3 = append(v3, 460)
                        }
                        _ = fi121
                    }
                    _ = fi115
                    _ = v116
                }
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = uint32(v1)
                }
            }
        case 128:
            v4 -= (v4 / (((v4 * 320)) | 1))
        case -9223372036854775808:
            var v124 int8 = -96
            _ = v124
        default:
            v2 = (((v2 + ("foo bar" + v2)) + (v2 + v2)) + (v2 + v2))
        }
    case "":
        for fi125 := 0; fi125 < 13; fi125++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            var v126 int32 = 630
            _ = fi125
            _ = v126
        }
    case "a":
        fmt.Println(v4)
    }
    for fi127 := 0; fi127 < 9; fi127++ {
        switch fi127 {
        case 4294967295:
            _ = len(v3)
        case -1:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v3 = append(v3, 4294967295)
        }
        _ = fi127
    }
    v128 := GenS10{F0: -128}
    switch v2 {
    case "foo bar":
        for fi129 := 0; fi129 < 1; fi129++ {
            v130 := GenS15{F0: 18446744073709551615}
            _ = fi129
            _ = v130
        }
    }
    var v131 int16 = 1
    _ = len(v3)
    for fi132 := 0; fi132 < 10; fi132++ {
        switch (v2 + v2) {
        case "foo bar":
            var v133 string = ((v2 + v2) + ("hello" + ("a" + "")))
            _ = v133
        default:
            for fi134 := 0; fi134 < 12; fi134++ {
                var v135 bool = ("" == (v2 + v2))
                _ = fi134
                _ = v135
            }
        }
        if true {
            if (!(!false)) {
                var v136 int = (^(fi132))
                _ = v136
            }
        }
        _ = fi132
    }
    var v137 float64 = ((1e10 * 0.1) * 1.5)
    v138 := GenS15{F0: 20}
    fmt.Println("v1:", v1)
    fmt.Println("v131:", v131)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v128
    _ = v131
    _ = v137
    _ = v138
    _ = v2
    _ = v3
    _ = v4
}
