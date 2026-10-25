package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 uint8 }
type GenS2 struct { F0 int32; F1 int16; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint; F1 uint; F2 uint8 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 bool; F2 uint16 }
type GenS10 struct { F0 int16; F1 uint64 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint, p1 uint16) int8 {
    var v7 float64 = float64(float32(-2.25))
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v7, p0, v7)
        _ = fi8
    }
    v9 := []int{-1, int(p1), 128, -2147483648}
    if (!false) {
        var v10 int64 = 0
        var v11 uint16 = (p1 ^ p1)
        _ = v10
        _ = v11
    }
    fmt.Println(v7, v7, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    var v6 uint32 = 2147483647
    delete(v3, "\n\t")
    v12 := GenS1{F0: 2, F1: -2147483648, F2: 127}
    v2[2] = -2147483648
    switch "" {
    case "a":
        v13 := GenS2{F0: -202, F1: -32768, F2: "foo bar"}
        _ = v13
    case "foo bar":
        v6 = v6
    default:
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint32 = (v6 % (((v6 | (^(v6)))) | 1))
            _ = fi14
            _ = v15
        }
    }
    switch "\n\t" {
    case "":
        switch ("\n\t" + "hello") {
        case "":
            var v16 float64 = 3.14159
            _ = v16
        case "hello":
            fmt.Println(v1)
        default:
            if (((54 <= 2) && false) && true) {
                for fi17 := 0; fi17 < 12; fi17++ {
                    for fi18 := 0; fi18 < 13; fi18++ {
                        _ = len(v5)
                        v2[v6] = fi17
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v5 = append(v5, 42)
                        v20 := GenS3{F0: 255}
                        _ = len(v5)
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                }
                v6 -= (^(v6))
            }
        }
    default:
        v21 := GenS2{F0: 2147483647, F1: -43, F2: "foo bar"}
        _ = v21
    }
    v22 := GenS4{F0: 0, F1: v6}
    for fi23 := 0; fi23 < 12; fi23++ {
        if false {
            if (3.14159 >= 0.0) {
                switch "hello" {
                case "hello":
                    switch (fi23 << ((65535) & 63)) {
                    case 127:
                        for fi24 := 0; fi24 < 11; fi24++ {
                            var v25 int64 = 0
                            v25 = -1
                            v26 := GenS5{F0: 14}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case 94:
                        for fi27 := 0; fi27 < 9; fi27++ {
                            switch "\n\t" {
                            case "\n\t":
                                if (!(!false)) {
                                    if false {
                                        fmt.Println(fi27)
                                        if len(v5) > 0 {
                                            v5[0] = 42
                                        }
                                    }
                                    v6 = v6
                                } else {
                                    v28 := GenS2{F0: -2147483648, F1: int16(fi27), F2: "\n\t"}
                                    if false {
                                        fmt.Println(fi23, fi23)
                                        v1 = v1
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 218
                                    }
                                    _ = v28
                                }
                            case "":
                                var v29 uint8 = 154
                                _ = v29
                            case "foo bar":
                                var v30 int32 = 42
                                _ = v30
                            default:
                                fmt.Println(fi23, fi23)
                            }
                            v1 -= v1
                            if (true && (int8(v6) > 87)) {
                                v31 := GenS2{F0: 0, F1: 1000, F2: ("hello" + ("hello" + "\n\t"))}
                                fmt.Println(v1, fi23)
                                fmt.Println(v6)
                                _ = v31
                            } else {
                                if (!(725 != (v6 - 896))) {
                                    switch (fi27 & (fi27 ^ 313)) {
                                    case 128:
                                        var v32 uint16 = 65535
                                        _ = v32
                                    default:
                                        fmt.Println(v1, v1)
                                    }
                                    var v33 int32 = 42
                                    if (-283 >= 32767) {
                                        var v34 uint64 = 241
                                        _ = v34
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v33
                                } else {
                                    v35 := GenS4{F0: 127, F1: (v6 >> ((721) & 63))}
                                    _ = v35
                                }
                            }
                            _ = fi27
                        }
                    case 4294967295:
                        v1 = v1
                    }
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    switch "\n\t" {
                    case "\n\t":
                        v1 = (v1 + (v1 / ((v1) | 1)))
                    case "":
                        if len(v5) > 0 {
                            v5[0] = (^(v6))
                        }
                    }
                    _ = fi36
                }
            }
        } else {
            var v37 float32 = 1e10
            for fi38 := 0; fi38 < 6; fi38++ {
                var v39 int64 = 0
                var v40 string = ""
                v39 += -613
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = v37
        }
        fmt.Println(v1, fi23)
        var v41 uint8 = 203
        _ = fi23
        _ = v41
    }
    v5 = append(v5, v6)
    var v42 float32 = 0.1
    for fi43 := 0; fi43 < 8; fi43++ {
        v1 -= (v1 << ((18446744073709551615) & 63))
        _ = fi43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        for fi45 := 0; fi45 < 17; fi45++ {
            v42 = (v42 + (float32(1e10) / (1e10 / v42)))
            v42 = v42
            _ = fi45
        }
        var v46 string = ""
        v46 = "a"
        _ = fi44
        _ = v46
    }
    var v47 string = "\n\t"
    for fi48 := 0; fi48 < 5; fi48++ {
        _ = len(v5)
        v42 -= float32(float64(v42))
        for fi49 := 0; fi49 < 7; fi49++ {
            _ = len(v3)
            switch (fi48 & 2147483647) {
            case 0:
                v50 := GenS5{F0: 1}
                _ = v50
            case -241:
                var v51 uint = 0
                _ = v51
            }
            if (("\n\t" + "") < v47) {
                if (172 == 470) {
                    var v52 int64 = int64(fi48)
                    v42 -= 0.1
                    fmt.Println(fi48, v6)
                    if false {
                        var v53 float32 = (v42 + (3.14159 / (((v42 - v42) * v42) / 1.0)))
                        v6 += v6
                        switch (fi48 * -753) {
                        case 9223372036854775807:
                            fmt.Println(v47, v1)
                        case 1:
                            var v54 int16 = 32767
                            _ = v54
                        case -9223372036854775808:
                            v6 += v6
                        default:
                            if false {
                                fmt.Println(v52, v52)
                                fmt.Println(fi48, v52)
                                switch (fi48 & (-(fi49))) {
                                case 4294967295:
                                    var v55 uint32 = v6
                                    _ = v55
                                case 841:
                                    switch fi49 {
                                    case 675:
                                        var v56 int16 = 0
                                        _ = v56
                                    case 42:
                                        v52 = (v52 | v52)
                                    default:
                                        switch "" {
                                        case "foo bar":
                                            var v57 string = "a"
                                            _ = v57
                                        default:
                                            v58 := GenS2{F0: 42, F1: 1000, F2: "foo bar"}
                                            _ = v58
                                        }
                                    }
                                case -2147483648:
                                    v6 += (v6 | 1)
                                default:
                                    v42 = v42
                                }
                            }
                        }
                        _ = v53
                    }
                    v42 = float32(float64(3.14159))
                    v42 = float32(float64(0.0))
                    _ = v52
                } else {
                    v1 = v1
                }
            } else {
                switch "\n\t" {
                case "a":
                    var v59 float64 = 3.14159
                    _ = v59
                case "":
                    v47 = (v47 + ("hello" + v47))
                default:
                    if true {
                        v6 -= 2
                        var v60 float64 = 1e10
                        _ = v60
                    }
                }
            }
            _ = fi49
        }
        _ = fi48
    }
    for fi61 := 0; fi61 < 10; fi61++ {
        if (!(((float32(0.0) * 0.0) * (v42 * ((float32((0.0 + (0.1 / (1e10 / (1e10 * 1.5))))) / ((v42 + ((v42 + v42) + 0.1)) + v42)) - 3.14159))) <= ((v42 + float32(1.5)) - 1.5))) {
            if ((v1 * 138) != 1) {
                switch "" {
                case "\n\t":
                    v62 := GenS1{F0: 127, F1: -1, F2: 255}
                    _ = v62
                }
                v63 := GenS4{F0: 2, F1: uint32(v1)}
                for fi64 := 0; fi64 < 14; fi64++ {
                    var v65 uint32 = (v6 | (v6 / ((uint32(v1)) | 1)))
                    var v66 float32 = v42
                    delete(v3, (v47 + "hello"))
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = v63
            }
            v67 := GenS6{F0: v6}
            _ = v67
        } else {
            var v68 int = fi61
            for fi69 := 0; fi69 < 8; fi69++ {
                fmt.Println(v42, v6)
                if false {
                    if true {
                        var v70 string = ("a" + (v47 + v47))
                        _ = v70
                    } else {
                        v42 += (float32((1.5 - 0.1)) + v42)
                    }
                    switch "" {
                    case "hello":
                        switch (v68 << ((529) & 63)) {
                        case 9223372036854775807:
                            if false {
                                if true {
                                    v47 = v47
                                    if (!(!true)) {
                                        fmt.Println(v68, v1)
                                        fmt.Println(v68)
                                        switch 32767 {
                                        case 4294967295:
                                            var v71 uint16 = 2
                                            _ = v71
                                        case 128:
                                            fmt.Println(v47, v42)
                                        case 42:
                                            var v72 float64 = ((0.0 / 3.14159) / (((0.1 + (3.14159 / 1e10)) * 3.14159) / float64(1.5)))
                                            _ = v72
                                        }
                                    }
                                    v47 = (v47 + v47)
                                } else {
                                    v47 = ("hello" + v47)
                                    var v73 uint16 = 42
                                    _ = v73
                                }
                                _ = len(v4)
                                _ = len(v5)
                            }
                        case 32767:
                            fmt.Println(v6, fi61, fi69)
                        }
                    default:
                        if (0 >= genFn0(210, 161)) {
                            var v74 int32 = 2147483647
                            _ = v74
                        }
                    }
                    if (!(!(-1 < 24))) {
                        var v75 string = (v47 + (v47 + v47))
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        fmt.Println(fi69)
                        _ = v75
                    } else {
                        v47 = v47
                        fmt.Println(fi69, v6)
                        switch (fi69 & fi69) {
                        case 1:
                            v76 := GenS4{F0: 163, F1: 42}
                            _ = v76
                        case -1:
                            var v77 bool = true
                            _ = v77
                        }
                    }
                } else {
                    if true {
                        if (2 == 255) {
                            var v78 string = "foo bar"
                            _ = v78
                        } else {
                            v79 := GenS3{F0: 0}
                            _ = v79
                        }
                        delete(v2, 1)
                        var v80 uint8 = 0
                        _ = v80
                    } else {
                        v1 -= v1
                    }
                    var v81 int32 = 2147483647
                    switch (-(fi69)) {
                    case -2147483648:
                        var v82 bool = true
                        _ = v82
                    case -227:
                        var v83 int16 = 32767
                        _ = v83
                    }
                    _ = v81
                }
                _ = fi69
            }
            _ = v68
        }
        for fi84 := 0; fi84 < 5; fi84++ {
            v85 := GenS7{F0: 608, F1: 255, F2: 127}
            if ((^(v1)) >= (v1 ^ 4294967295)) {
                fmt.Println(v1, v42)
            }
            var v86 int64 = -183
            var v87 uint = 2
            _ = fi84
            _ = v85
            _ = v86
            _ = v87
        }
        _ = fi61
    }
    delete(v3, "hello")
    if true {
        switch (((v47 + v47) + "\n\t") + v47) {
        case "hello":
            for fi88 := 0; fi88 < 12; fi88++ {
                v89 := v4[(v6 & uint32(v1))]
                v1 = (^(v1))
                _ = fi88
                _ = v89
            }
        case "foo bar":
            var v90 uint64 = 18446744073709551615
            _ = v90
        case "":
            if (v42 < v42) {
                v91 := GenS4{F0: 0, F1: v6}
                _ = v91
            } else {
                for fi92 := 0; fi92 < 16; fi92++ {
                    v93 := GenS8{F0: true}
                    if true {
                        fmt.Println(v47)
                        if false {
                            for fi94 := 0; fi94 < 10; fi94++ {
                                if (1 == 0) {
                                    v95 := GenS9{F0: v47, F1: (((((v47 + (v47 + "a")) + (v47 + "\n\t")) + "hello") == v47) || true), F2: 65535}
                                    fmt.Println(v42, v1)
                                    var v96 uint8 = 1
                                    _ = v95
                                    _ = v96
                                }
                                var v97 float64 = 1.5
                                var v98 uint8 = 42
                                _ = fi94
                                _ = v97
                                _ = v98
                            }
                            for fi99 := 0; fi99 < 6; fi99++ {
                                if ((v42 / v42) > v42) {
                                    var v100 int8 = -1
                                    v47 = v47
                                    v6 -= uint32(v1)
                                    var v101 int64 = 0
                                    _ = v100
                                    _ = v101
                                }
                                var v102 float64 = 0.0
                                _ = fi99
                                _ = v102
                            }
                            v103 := v2[uint32(v1)]
                            _ = v103
                        } else {
                            switch v47 {
                            case "":
                                v104 := GenS10{F0: 0, F1: v1}
                                _ = v104
                            case "foo bar":
                                if true {
                                    v1 = 603
                                    var v105 bool = false
                                    _ = v105
                                } else {
                                    for fi106 := 0; fi106 < 6; fi106++ {
                                        var v107 int64 = 149
                                        var v108 uint = 255
                                        var v109 int64 = 795
                                        _ = fi106
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    }
                                    if (583 > -1) {
                                        fmt.Println(fi92)
                                    }
                                }
                            default:
                                switch ((v47 + "foo bar") + "a") {
                                case "":
                                    v110 := GenS10{F0: 1000, F1: 1}
                                    _ = v110
                                case "hello":
                                    fmt.Println(v1, v42)
                                case "foo bar":
                                    v42 = v42
                                default:
                                    var v111 uint16 = uint16(v6)
                                    _ = v111
                                }
                            }
                            _ = len(v2)
                        }
                        if (v6 <= v6) {
                            if ((v1 * v1) < 4294967295) {
                                var v112 string = ""
                                _ = v112
                            } else {
                                var v113 int16 = 49
                                if (1.5 <= 0.1) {
                                    var v114 uint32 = (v6 * 65535)
                                    for fi115 := 0; fi115 < 15; fi115++ {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v1, fi92)
                                        _ = fi115
                                    }
                                    var v116 int64 = 42
                                    _ = v114
                                    _ = v116
                                } else {
                                    var v117 int64 = 42
                                    _ = v117
                                }
                                v118 := GenS9{F0: "foo bar", F1: true, F2: 1}
                                _ = v113
                                _ = v118
                            }
                            var v119 float64 = (3.14159 * 3.14159)
                            _ = v119
                        } else {
                            v47 = (v47 + v47)
                            for fi120 := 0; fi120 < 9; fi120++ {
                                v121 := GenS9{F0: ("foo bar" + "\n\t"), F1: ((!false) || true), F2: 0}
                                _ = fi120
                                _ = v121
                            }
                        }
                    }
                    var v122 int = (fi92 + fi92)
                    _ = fi92
                    _ = v122
                    _ = v93
                }
                fmt.Println(v42, v47, v6)
                var v123 int8 = genFn0(210, 1)
                _ = v123
            }
        }
    } else {
        var v124 float64 = ((-2.25 * -2.25) / (1.5 - 1e10))
        for fi125 := 0; fi125 < 2; fi125++ {
            _ = len(v5)
            _ = fi125
        }
        var v126 uint8 = 1
        _ = v124
        _ = v126
    }
    if (v1 < (^(v1))) {
        for fi127 := 0; fi127 < 10; fi127++ {
            for fi128 := 0; fi128 < 17; fi128++ {
                var v129 int64 = 0
                v47 = "hello"
                _ = fi128
                _ = v129
            }
            _ = fi127
        }
        fmt.Println(v47)
        for fi130 := 0; fi130 < 5; fi130++ {
            if true {
                if (2 > v1) {
                    v131 := GenS11{F0: 0}
                    _ = v131
                } else {
                    v6 -= (v6 << ((4294967295) & 63))
                    v1 = (^(v1))
                    v5 = append(v5, v6)
                }
            } else {
                if (-787 > (fi130 % ((fi130) | 1))) {
                    for fi132 := 0; fi132 < 14; fi132++ {
                        var v133 bool = (1 != 42)
                        var v134 uint8 = 255
                        _ = fi132
                        _ = v133
                        _ = v134
                    }
                }
                switch (v47 + v47) {
                case "":
                    if (-421 >= -32768) {
                        for fi135 := 0; fi135 < 13; fi135++ {
                            if (fi130 > (fi130 | (fi135 * fi135))) {
                                switch (v47 + v47) {
                                case "a":
                                    if ((false && ((-2.25 == (-2.25 * (-2.25 + (3.14159 / 3.14159)))) || (-314 > -1))) && (1 >= -1)) {
                                        var v136 uint32 = v6
                                        switch (fi135 % ((fi135) | 1)) {
                                        case 32768:
                                            v42 = (-2.25 * (v42 * (v42 - float32((float64(v42) / 0.1)))))
                                        default:
                                            switch 32768 {
                                            case 32768:
                                                v1 = v1
                                            case 9223372036854775807:
                                                if true {
                                                    if (18446744073709551615 < (v1 / (((^(v1))) | 1))) {
                                                        v47 = "hello"
                                                    }
                                                    var v137 int8 = -128
                                                    v138 := GenS12{F0: v47}
                                                    v6 = v6
                                                    switch 1 {
                                                    case -2:
                                                        if len(v5) > 0 {
                                                            v5[0] = v6
                                                        }
                                                    case 4294967295:
                                                        var v139 string = ((("" + (v47 + "a")) + "a") + v47)
                                                        _ = v139
                                                    case 2147483647:
                                                        if true {
                                                            v6 = v136
                                                            switch v47 {
                                                            case "hello":
                                                                v6 -= (v6 << ((0) & 63))
                                                            }
                                                            if (0 < 972) {
                                                                switch v47 {
                                                                case "a":
                                                                    v136 += 2147483647
                                                                }
                                                                var v140 uint64 = (v1 >> ((18446744073709551615) & 63))
                                                                var v141 float32 = v42
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                        }
                                                    }
                                                    _ = v137
                                                    _ = v138
                                                } else {
                                                    _ = len(v2)
                                                    fmt.Println(v6)
                                                    v136 -= (v136 >> ((255) & 63))
                                                    var v142 uint32 = (v6 >> ((uint(v1)) & 63))
                                                    _ = v142
                                                }
                                            }
                                            fmt.Println(v6, v47, v136)
                                            if len(v5) > 0 {
                                                v5[0] = (v136 * (v136 | (v136 & 720)))
                                            }
                                        }
                                        fmt.Println(fi130)
                                        _ = v136
                                    } else {
                                        delete(v3, v47)
                                        switch v47 {
                                        case "\n\t":
                                            if len(v5) > 0 {
                                                v5[0] = (v6 / ((v6) | 1))
                                            }
                                        case "hello":
                                            fmt.Println(v47, v47)
                                        case "foo bar":
                                            if len(v5) > 0 {
                                                v5[0] = 1
                                            }
                                        }
                                    }
                                case "":
                                    var v143 int8 = 45
                                    _ = v143
                                default:
                                    var v144 int64 = -1
                                    _ = v144
                                }
                                v47 = v47
                                v5 = append(v5, v6)
                            } else {
                                fmt.Println(fi130)
                            }
                            var v145 float64 = 0.1
                            v47 = (("\n\t" + (v47 + "a")) + v47)
                            _ = fi135
                            _ = v145
                        }
                        fmt.Println(v6, v47, v42)
                    } else {
                        switch 32767 {
                        case -9223372036854775808:
                            v6 = v6
                        case 0:
                            if (-128 >= genFn0(0, 0)) {
                                var v146 uint32 = v6
                                v147 := GenS9{F0: v47, F1: false, F2: 42}
                                v1 += (v1 | v1)
                                _ = v146
                                _ = v147
                            } else {
                                var v148 uint32 = 724
                                var v149 int = 9223372036854775807
                                v150 := GenS7{F0: uint(v6), F1: 1, F2: 0}
                                _ = v148
                                _ = v149
                                _ = v150
                            }
                        case -297:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        v151 := GenS13{F0: (^(v1)), F1: -831}
                        _ = v151
                    }
                case "hello":
                    var v152 uint = 18446744073709551615
                    _ = v152
                }
                v153 := GenS1{F0: 127, F1: -1, F2: 1}
                _ = v153
            }
            _ = fi130
        }
    }
    delete(v4, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v42
    _ = v47
    _ = v5
    _ = v6
}
