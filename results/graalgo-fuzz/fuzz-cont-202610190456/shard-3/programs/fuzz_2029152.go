package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 bool; F2 bool }
type GenS8 struct { F0 string; F1 int32; F2 string }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint64; F1 uint; F2 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool; F1 int32 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint; F1 uint64; F2 uint8 }
func genFn0(p0 uint8) uint16 {
    v5 := []int{576, 128, -1}
    var v6 uint16 = 2
    switch ("" + "\n\t") {
    case "foo bar":
        if (65535 == 2) {
            v7 := GenS1{F0: uint32(v6)}
            _ = v7
        }
    }
    v8 := GenS1{F0: 65535}
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return v6
}

func main() {
    var v1 int16 = 847
    var v2 uint8 = 48
    v3 := make(map[uint32]string)
    var v4 int64 = -9223372036854775808
    switch 9223372036854775807 {
    case 127:
        if ("foo bar" < "") {
            switch "hello" {
            case "foo bar":
                var v9 bool = (42 <= 42)
                _ = v9
            }
            v10 := []int{-849, -814}
            _ = v10
        } else {
            if true {
                if true {
                    if (4294967295 < 1) {
                        v11 := GenS2{F0: ("a" + "\n\t"), F1: (161 > 32767)}
                        v2 = v2
                        var v12 bool = (true || true)
                        _ = v11
                        _ = v12
                    }
                }
                switch ("\n\t" + "hello") {
                case "hello":
                    if false {
                        v4 = int64(v2)
                    }
                default:
                    v13 := GenS3{F0: (v4 + v4), F1: v1}
                    _ = v13
                }
            }
            var v14 float64 = 0.0
            _ = v14
        }
    case 9223372036854775807:
        fmt.Println(v2)
    case 128:
        if (127 == 38) {
            for fi15 := 0; fi15 < 18; fi15++ {
                switch "hello" {
                case "hello":
                    var v16 bool = (2147483647 == 428)
                    _ = v16
                case "\n\t":
                    fmt.Println(v4, v2, v4)
                case "foo bar":
                    switch "a" {
                    case "\n\t":
                        v1 -= v1
                    case "foo bar":
                        var v17 uint8 = (v2 / ((v2) | 1))
                        _ = v17
                    }
                default:
                    if false {
                        v18 := []int{42, 42, int(v2), (fi15 % (((fi15 << ((0) & 63))) | 1))}
                        switch "" {
                        case "a":
                            var v19 uint = 18446744073709551615
                            _ = v19
                        case "\n\t":
                            v20 := GenS4{F0: v2, F1: "a"}
                            _ = v20
                        case "foo bar":
                            fmt.Println(v2, fi15)
                        default:
                            var v21 int16 = (v1 << ((2) & 63))
                            _ = v21
                        }
                        var v22 uint8 = (v2 << ((0) & 63))
                        _ = v18
                        _ = v22
                    } else {
                        if false {
                            v2 = v2
                        }
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v24 := GenS5{F0: 676}
                            if (1 < 980) {
                                switch (^(fi23)) {
                                case 310:
                                    var v25 int = fi23
                                    _ = v25
                                case -54:
                                    v26 := v3[2]
                                    fmt.Println(v2)
                                    fmt.Println(v1, v26, v2)
                                    v27 := []int{(fi15 - -1), fi23, -2147483648}
                                    _ = v26
                                    _ = v27
                                default:
                                    v4 = 402
                                }
                                var v28 int64 = (v4 - (v4 % (((-(v4))) | 1)))
                                _ = v28
                            }
                            v29 := GenS2{F0: ("" + "foo bar"), F1: true}
                            _ = fi23
                            _ = v24
                            _ = v29
                        }
                    }
                }
                var v30 int64 = (v4 ^ (^(v4)))
                _ = fi15
                _ = v30
            }
            if true {
                if (!true) {
                    if ((v4 >> ((0) & 63)) < 42) {
                        if (true || (1 >= -1)) {
                            v1 += v1
                            if ((!((false || (true && (!(4294967295 > 1)))) || ((!(32767 != 42)) || ((v4 << ((2) & 63)) > (v4 << ((853) & 63)))))) || false) {
                                v4 = v4
                                for fi31 := 0; fi31 < 12; fi31++ {
                                    v32 := GenS4{F0: 0, F1: ("foo bar" + ("hello" + ("a" + "foo bar")))}
                                    var v33 uint = 141
                                    _ = fi31
                                    _ = v32
                                    _ = v33
                                }
                                _ = len(v3)
                            }
                            v34 := GenS3{F0: (v4 - (v4 * v4)), F1: 1}
                            _ = v34
                        } else {
                            v2 = (v2 << ((998) & 63))
                        }
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v36 := GenS6{F0: true}
                            v1 = (v1 >> ((478) & 63))
                            var v37 int16 = v1
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        if (18446744073709551615 != 999) {
                            v2 += (v2 >> ((65535) & 63))
                            var v38 float32 = -2.25
                            _ = v38
                        }
                    }
                    switch 0 {
                    case -2147483648:
                        fmt.Println(v1)
                    case -922:
                        var v39 uint8 = (v2 - 1)
                        _ = v39
                    }
                    var v40 int16 = (v1 ^ int16(v4))
                    _ = v40
                }
                var v41 int64 = (v4 | 0)
                if (v4 != 42) {
                    v42 := GenS4{F0: 255, F1: ("\n\t" + ("hello" + ("foo bar" + "foo bar")))}
                    _ = v42
                } else {
                    if (1 <= (v1 % ((v1) | 1))) {
                        var v43 string = "a"
                        _ = v43
                    }
                    var v44 float32 = 3.14159
                    _ = v44
                }
                _ = v41
            } else {
                if true {
                    switch 9223372036854775807 {
                    case 127:
                        if false {
                            switch "hello" {
                            case "foo bar":
                                var v45 int16 = 0
                                _ = v45
                            case "":
                                switch "\n\t" {
                                case "\n\t":
                                    for fi46 := 0; fi46 < 7; fi46++ {
                                        switch "\n\t" {
                                        case "":
                                            var v47 uint32 = 2
                                            _ = v47
                                        case "a":
                                            v2 += v2
                                        case "\n\t":
                                            v48 := GenS7{F0: -2147483648, F1: (false || true), F2: true}
                                            _ = v48
                                        default:
                                            if (4294967295 > 4294967295) {
                                                if ((-(fi46)) != (fi46 >> ((503) & 63))) {
                                                    var v49 int16 = v1
                                                    _ = v49
                                                } else {
                                                    if false {
                                                        v50 := []int{int(v4), (fi46 >> ((0) & 63)), fi46}
                                                        v1 = (v1 - (v1 >> ((255) & 63)))
                                                        _ = v50
                                                    } else {
                                                        switch (fi46 * (fi46 * fi46)) {
                                                        case 0:
                                                            v51 := GenS1{F0: 42}
                                                            _ = v51
                                                        case 128:
                                                            v52 := []int{fi46, (fi46 - 32767)}
                                                            _ = v52
                                                        default:
                                                            for fi53 := 0; fi53 < 20; fi53++ {
                                                                v54 := GenS8{F0: ("a" + "hello"), F1: 0, F2: "hello"}
                                                                switch "" {
                                                                case "\n\t":
                                                                    v2 += v2
                                                                case "foo bar":
                                                                    switch (fi46 << ((65535) & 63)) {
                                                                    case 2147483647:
                                                                        var v55 int16 = v1
                                                                        _ = v55
                                                                    case 32768:
                                                                        if (v4 != 61) {
                                                                            v56 := []int{(fi53 % (((-(fi46))) | 1)), (fi46 % ((fi46) | 1)), (fi53 >> ((9) & 63))}
                                                                            _ = v56
                                                                        } else {
                                                                            fmt.Println(fi46)
                                                                            var v57 uint32 = 65535
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    v58 := GenS7{F0: 2147483647, F1: true, F2: true}
                                                                    _ = v58
                                                                default:
                                                                    v4 = (v4 << ((4294967295) & 63))
                                                                }
                                                                var v59 int32 = 0
                                                                _ = fi53
                                                                _ = v54
                                                                _ = v59
                                                            }
                                                        }
                                                    }
                                                }
                                                switch fi46 {
                                                case 32767:
                                                    v60 := GenS6{F0: true}
                                                    _ = v60
                                                case 0:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v61 := GenS3{F0: 9223372036854775807, F1: 1}
                                                        _ = v61
                                                    case "":
                                                        var v62 uint8 = v2
                                                        _ = v62
                                                    default:
                                                        for fi63 := 0; fi63 < 2; fi63++ {
                                                            switch ("a" + (("hello" + "") + ("\n\t" + (("hello" + ("foo bar" + ("a" + ""))) + ((("hello" + "hello") + "\n\t") + (("\n\t" + "a") + "")))))) {
                                                            case "":
                                                                v1 += v1
                                                            case "foo bar":
                                                                v4 = 1
                                                            case "hello":
                                                                fmt.Println(v1)
                                                            }
                                                            v4 = int64(fi63)
                                                            _ = fi63
                                                        }
                                                    }
                                                case -192:
                                                    switch (^(fi46)) {
                                                    case 998:
                                                        for fi64 := 0; fi64 < 2; fi64++ {
                                                            if (0 > 32767) {
                                                                switch fi46 {
                                                                case 0:
                                                                    v1 = v1
                                                                case 601:
                                                                    var v65 uint16 = 921
                                                                    _ = v65
                                                                case 32768:
                                                                    v66 := GenS7{F0: int32(v2), F1: (v1 > int16(v2)), F2: (true || true)}
                                                                    _ = v66
                                                                default:
                                                                    v1 -= (v1 * (v1 % (((v1 % ((v1) | 1))) | 1)))
                                                                }
                                                                v4 = v4
                                                                fmt.Println(fi46, fi46)
                                                            }
                                                            var v67 int32 = 42
                                                            _ = fi64
                                                            _ = v67
                                                        }
                                                    case -1:
                                                        for fi68 := 0; fi68 < 6; fi68++ {
                                                            v69 := GenS9{F0: 0}
                                                            _ = fi68
                                                            _ = v69
                                                        }
                                                    default:
                                                        v1 -= (^(v1))
                                                    }
                                                default:
                                                    v70 := GenS2{F0: "", F1: false}
                                                    _ = v70
                                                }
                                                fmt.Println(fi46, v2, v2)
                                            }
                                        }
                                        if true {
                                            v71 := []int{}
                                            v2 = (v2 | v2)
                                            _ = v71
                                        } else {
                                            switch (("\n\t" + "\n\t") + ("a" + "foo bar")) {
                                            case "hello":
                                                v72 := []int{(fi46 ^ fi46), -946, fi46, (fi46 ^ (fi46 >> ((2) & 63)))}
                                                _ = v72
                                            case "a":
                                                v73 := GenS3{F0: (^(v4)), F1: (v1 + -1)}
                                                _ = v73
                                            case "":
                                                for fi74 := 0; fi74 < 19; fi74++ {
                                                    v75 := []int{32767, fi46, fi46}
                                                    _ = fi74
                                                    _ = v75
                                                }
                                            default:
                                                switch (^(fi46)) {
                                                case 32768:
                                                    for fi76 := 0; fi76 < 10; fi76++ {
                                                        v1 = (v1 << ((1) & 63))
                                                        if false {
                                                            v4 = v4
                                                        }
                                                        _ = fi76
                                                    }
                                                default:
                                                    v77 := []int{(fi46 % (((fi46 | fi46)) | 1)), fi46}
                                                    _ = v77
                                                }
                                            }
                                        }
                                        var v78 uint8 = v2
                                        _ = fi46
                                        _ = v78
                                    }
                                case "foo bar":
                                    if (-2147483648 >= -133) {
                                        var v79 uint32 = 2
                                        for fi80 := 0; fi80 < 7; fi80++ {
                                            v81 := []int{fi80, (fi80 * (fi80 - (fi80 + 0))), (fi80 & fi80), 9223372036854775807}
                                            v82 := GenS1{F0: (v79 / ((42) | 1))}
                                            for fi83 := 0; fi83 < 13; fi83++ {
                                                fmt.Println(v4, fi80, fi83)
                                                _ = fi83
                                            }
                                            _ = fi80
                                            _ = v81
                                            _ = v82
                                        }
                                        _ = v79
                                    } else {
                                        if ((3.14159 * 3.14159) >= -2.25) {
                                            v84 := []int{127, 1, 127}
                                            if false {
                                                v4 = v4
                                                switch 497 {
                                                case 950:
                                                    v85 := GenS2{F0: "foo bar", F1: (906 <= 0)}
                                                    _ = v85
                                                case 4294967295:
                                                    v1 += (v1 | (v1 - v1))
                                                case 9223372036854775807:
                                                    if (26 < 0) {
                                                        for fi86 := 0; fi86 < 6; fi86++ {
                                                            _ = len(v3)
                                                            var v87 string = "foo bar"
                                                            _ = fi86
                                                            _ = v87
                                                        }
                                                        _ = len(v3)
                                                    }
                                                }
                                            }
                                            if (!("a" > "\n\t")) {
                                                for fi88 := 0; fi88 < 9; fi88++ {
                                                    var v89 float32 = 0.1
                                                    _ = fi88
                                                    _ = v89
                                                }
                                            } else {
                                                v90 := GenS1{F0: 65535}
                                                _ = v90
                                            }
                                            _ = v84
                                        }
                                    }
                                }
                            default:
                                if true {
                                    switch 9223372036854775807 {
                                    case 42:
                                        var v91 int32 = -2147483648
                                        _ = v91
                                    case 206:
                                        fmt.Println(v1)
                                    default:
                                        if false {
                                            if (-1 <= 42) {
                                                if ((!(((!true) || (v4 != -871)) || false)) && (2147483648 < 712)) {
                                                    var v92 int64 = (v4 | 1)
                                                    var v93 int8 = 42
                                                    var v94 uint16 = 42
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                }
                                            } else {
                                                for fi95 := 0; fi95 < 16; fi95++ {
                                                    v1 -= (^(v1))
                                                    switch (("" + "foo bar") + "\n\t") {
                                                    case "hello":
                                                        v96 := GenS7{F0: 42, F1: ((((("a" + ((((("a" + "\n\t") + "a") + ("foo bar" + "")) + "foo bar") + "")) + ((("\n\t" + "\n\t") + "") + (("hello" + "\n\t") + "hello"))) + "a") < "foo bar") && (!(!(65535 == 2147483648)))), F2: (!((v4 - -1) == v4))}
                                                        _ = v96
                                                    case "\n\t":
                                                        var v97 uint32 = 2147483647
                                                        _ = v97
                                                    default:
                                                        var v98 int32 = 338
                                                        _ = v98
                                                    }
                                                    _ = fi95
                                                }
                                            }
                                        } else {
                                            for fi99 := 0; fi99 < 4; fi99++ {
                                                fmt.Println(v2, v1)
                                                switch "" {
                                                case "":
                                                    for fi100 := 0; fi100 < 16; fi100++ {
                                                        v4 = v4
                                                        v4 -= v4
                                                        switch "" {
                                                        case "foo bar":
                                                            switch fi99 {
                                                            case 4294967295:
                                                                v101 := GenS10{F0: 0, F1: 255, F2: 0}
                                                                _ = v101
                                                            }
                                                        case "hello":
                                                            var v102 float64 = 1e10
                                                            _ = v102
                                                        }
                                                        _ = fi100
                                                    }
                                                }
                                                v103 := v3[65535]
                                                _ = fi99
                                                _ = v103
                                            }
                                            switch -209 {
                                            case 9223372036854775807:
                                                fmt.Println(v2, v4, v2)
                                            case 128:
                                                var v104 string = "a"
                                                _ = v104
                                            case 864:
                                                v2 = 2
                                            }
                                            for fi105 := 0; fi105 < 17; fi105++ {
                                                v2 = (^(v2))
                                                var v106 int16 = (v1 | v1)
                                                _ = fi105
                                                _ = v106
                                            }
                                        }
                                    }
                                    switch "\n\t" {
                                    case "":
                                        if (-1 > 1) {
                                            var v107 float32 = 3.14159
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            default:
                                                v2 = 42
                                            }
                                            _ = v107
                                        }
                                    case "foo bar":
                                        v4 = 9223372036854775807
                                    default:
                                        v2 = v2
                                    }
                                    switch 32767 {
                                    case -758:
                                        var v108 uint = 42
                                        _ = v108
                                    }
                                }
                            }
                            var v109 int = 774
                            for fi110 := 0; fi110 < 12; fi110++ {
                                v2 = (^(v2))
                                fmt.Println(v2)
                                _ = fi110
                            }
                            _ = v109
                        }
                    case 32768:
                        fmt.Println(v4)
                    default:
                        v2 = (v2 >> ((0) & 63))
                    }
                } else {
                    v1 -= v1
                    v2 = 255
                }
                if true {
                    for fi111 := 0; fi111 < 20; fi111++ {
                        var v112 uint32 = uint32(v2)
                        var v113 int64 = int64(v1)
                        _ = fi111
                        _ = v112
                        _ = v113
                    }
                    if false {
                        var v114 int8 = 42
                        if false {
                            v4 = v4
                            var v115 int = 127
                            switch "\n\t" {
                            case "a":
                                v116 := []int{(v115 * (v115 - v115)), (v115 << ((1) & 63)), 0}
                                _ = v116
                            default:
                                for fi117 := 0; fi117 < 10; fi117++ {
                                    var v118 uint16 = 0
                                    switch ("a" + (("" + "foo bar") + "a")) {
                                    case "a":
                                        var v119 float32 = (3.14159 / float32(1.5))
                                        _ = v119
                                    }
                                    _ = fi117
                                    _ = v118
                                }
                            }
                            _ = v115
                        }
                        var v120 uint64 = 4294967295
                        _ = v114
                        _ = v120
                    } else {
                        for fi121 := 0; fi121 < 15; fi121++ {
                            fmt.Println(fi121, fi121)
                            if ("" != "hello") {
                                fmt.Println(v2, fi121)
                                fmt.Println(fi121)
                                fmt.Println(v1)
                            }
                            switch ("\n\t" + ("foo bar" + ((((("\n\t" + "hello") + (("" + "") + "\n\t")) + "\n\t") + "\n\t") + ""))) {
                            case "foo bar":
                                v122 := []int{fi121, (fi121 | 0), fi121, (fi121 & (-(fi121)))}
                                _ = v122
                            case "\n\t":
                                if true {
                                    v4 = v4
                                }
                            case "a":
                                v123 := GenS11{F0: 106}
                                _ = v123
                            default:
                                var v124 int = 1
                                _ = v124
                            }
                            _ = fi121
                        }
                    }
                    v1 = (v1 << ((65535) & 63))
                } else {
                    for fi125 := 0; fi125 < 15; fi125++ {
                        var v126 int8 = -1
                        fmt.Println(fi125, v4, v126)
                        fmt.Println(v1, v2)
                        _ = fi125
                        _ = v126
                    }
                }
                var v127 bool = (false && (true && false))
                _ = v127
            }
            for fi128 := 0; fi128 < 13; fi128++ {
                v129 := GenS12{F0: (!(4294967295 < 42)), F1: -304}
                if (42 != 127) {
                    v3[42] = "hello"
                    fmt.Println(v2, fi128)
                    v130 := []int{(fi128 ^ int(v4)), 781, fi128, fi128}
                    _ = v130
                }
                switch (fi128 * (^(fi128))) {
                case 127:
                    v131 := GenS9{F0: -49}
                    _ = v131
                }
                _ = fi128
                _ = v129
            }
        }
    }
    if (!((!true) && true)) {
        if true {
            if true {
                if false {
                    v132 := GenS13{F0: 65535}
                    v1 = v1
                    v1 -= v1
                    _ = v132
                }
            }
            v133 := []int{-1, 32768, 127}
            _ = v133
        }
    }
    var v134 uint16 = 660
    v135 := []int{}
    for fi136 := 0; fi136 < 20; fi136++ {
        var v137 uint64 = 42
        _ = fi136
        _ = v137
    }
    switch -1 {
    case 9223372036854775807:
        var v138 int8 = int8(v134)
        _ = v138
    case 127:
        for fi139 := 0; fi139 < 3; fi139++ {
            var v140 uint16 = v134
            _ = fi139
            _ = v140
        }
    default:
        v1 = v1
    }
    fmt.Println(v1)
    var v141 string = ("" + "\n\t")
    switch (("\n\t" + (("foo bar" + v141) + v141)) + (v141 + "\n\t")) {
    case "":
        v2 += (v2 / (((v2 * (v2 | (v2 % ((127) | 1))))) | 1))
    case "\n\t":
        if true {
            v4 = 9223372036854775807
            v134 = v134
            v142 := GenS14{F0: 1, F1: 18446744073709551615, F2: 255}
            _ = v142
        } else {
            fmt.Println(v4, v4, v1)
        }
    case "a":
        var v143 float64 = (((float64(0.0) + 0.1) - (1.5 / 1.0)) + (1e10 + float64(0.1)))
        _ = v143
    }
    if false {
        if (((v1 != (v1 << ((0) & 63))) || ((!false) || true)) || (-2.25 == 1e10)) {
            var v144 int16 = (v1 / ((v1) | 1))
            v145 := GenS3{F0: v4, F1: (v1 >> ((0) & 63))}
            _ = v144
            _ = v145
        }
    }
    v146 := v3[42]
    for fi147 := 0; fi147 < 14; fi147++ {
        v146 = ("foo bar" + "a")
        var v148 int8 = 42
        _ = fi147
        _ = v148
    }
    var v149 uint16 = 42
    v134 += v134
    if (2 != 42) {
        v135 = append(v135, -1)
        var v150 uint64 = 42
        if (!true) {
            var v151 float32 = ((0.0 / 3.14159) + 3.14159)
            fmt.Println(v2, v134)
            _ = v151
        } else {
            var v152 uint16 = 2
            v3[42] = "\n\t"
            switch "hello" {
            case "":
                v141 = v146
            }
            _ = v152
        }
        _ = v150
    } else {
        v141 = ("\n\t" + "foo bar")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v135:", len(v135), v135)
    fmt.Println("v141:", v141)
    fmt.Println("v146:", v146)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v134
    _ = v135
    _ = v141
    _ = v146
    _ = v149
    _ = v2
    _ = v3
    _ = v4
}
