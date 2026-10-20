package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int64; F2 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 string; F1 int64; F2 int8 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint16; F1 int8; F2 bool }
type GenS15 struct { F0 int32; F1 int; F2 int64 }
type GenS16 struct { F0 int; F1 uint16 }
type GenS17 struct { F0 int; F1 int8; F2 bool }
type GenS18 struct { F0 uint }
func genFn0() int64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 uint8 = 128
        var v6 float32 = (0.0 / 1e10)
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int32 = 1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.0
    v3 := []uint32{2147483648, 65535, 42}
    v8 := GenS1{F0: 2, F1: 1, F2: 4294967295}
    for fi9 := 0; fi9 < 3; fi9++ {
        fmt.Println(v2, fi9)
        _ = fi9
    }
    var v10 int = 635
    var v11 int64 = 504
    if true {
        for fi12 := 0; fi12 < 16; fi12++ {
            if true {
                for fi13 := 0; fi13 < 7; fi13++ {
                    if (!false) {
                        v14 := GenS1{F0: 0, F1: 255, F2: 595}
                        _ = v14
                    } else {
                        v10 -= fi12
                        var v15 int8 = 37
                        _ = v15
                    }
                    _ = fi13
                }
                v11 -= genFn0()
                for fi16 := 0; fi16 < 14; fi16++ {
                    if true {
                        fmt.Println(v11, v11)
                        v3 = append(v3, 0)
                        switch (fi16 ^ 9223372036854775807) {
                        case -2147483648:
                            v2 = 1e10
                        case -9223372036854775808:
                            v17 := GenS2{F0: fi12, F1: 32767}
                            _ = v17
                        case 9223372036854775807:
                            if false {
                                v2 = v2
                                var v18 int16 = 1000
                                _ = v18
                            }
                        }
                    }
                    v10 = fi16
                    v19 := GenS1{F0: 1, F1: 42, F2: 711}
                    _ = fi16
                    _ = v19
                }
            } else {
                v2 -= (v2 + 1e10)
                var v20 int64 = v11
                _ = v20
            }
            var v21 int16 = 0
            switch (-(fi12)) {
            case 32767:
                if ((false && false) && (42 != 572)) {
                    fmt.Println(v2, v2, v2)
                    v10 = 743
                    if (genFn0() < v11) {
                        var v22 float64 = 0.1
                        v11 -= (v11 / (((v11 + 9223372036854775807)) | 1))
                        _ = v22
                    } else {
                        if (1e10 == float32(0.0)) {
                            switch "" {
                            case "\n\t":
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    var v24 float32 = 3.14159
                                    if (!(true || (2 == 1))) {
                                        v10 = 127
                                        v25 := GenS3{F0: 65535, F1: (-765 == 2147483647), F2: 734}
                                        var v26 string = "\n\t"
                                        _ = v25
                                        _ = v26
                                    } else {
                                        v10 = 1
                                    }
                                    v27 := GenS4{F0: (v21 >> ((2) & 63))}
                                    _ = fi23
                                    _ = v24
                                    _ = v27
                                }
                            case "":
                                var v28 int32 = 0
                                _ = v28
                            }
                            v29 := v1[("\n\t" + "a")]
                            switch (-(v29)) {
                            case -1:
                                v30 := GenS1{F0: 0, F1: 42, F2: 2}
                                _ = v30
                            case 42:
                                if false {
                                    if (v11 > genFn0()) {
                                        var v31 float32 = float32(0.0)
                                        if false {
                                            v11 = v11
                                        }
                                        v32 := v1[""]
                                        _ = v31
                                        _ = v32
                                    } else {
                                        if false {
                                            v10 = 4294967295
                                            switch "foo bar" {
                                            case "a":
                                                v33 := GenS2{F0: (v10 / (((v10 & 4294967295)) | 1)), F1: v21}
                                                _ = v33
                                            default:
                                                v34 := GenS5{F0: 255, F1: 1, F2: fi12}
                                                _ = v34
                                            }
                                        } else {
                                            fmt.Println(v21, v2)
                                            for fi35 := 0; fi35 < 15; fi35++ {
                                                if len(v3) > 0 {
                                                    v3[0] = 145
                                                }
                                                switch (fi12 % ((v29) | 1)) {
                                                case 1:
                                                    v36 := GenS6{F0: 514}
                                                    _ = v36
                                                case -386:
                                                    _ = len(v3)
                                                case 0:
                                                    v21 = 1
                                                default:
                                                    var v37 int16 = 0
                                                    _ = v37
                                                }
                                                var v38 uint16 = 446
                                                _ = fi35
                                                _ = v38
                                            }
                                        }
                                        v39 := v1["foo bar"]
                                        var v40 float32 = ((v2 + v2) * 1e10)
                                        _ = v39
                                        _ = v40
                                    }
                                    v41 := GenS2{F0: fi12, F1: v21}
                                    _ = v41
                                } else {
                                    for fi42 := 0; fi42 < 9; fi42++ {
                                        v2 -= (v2 - v2)
                                        v11 = (v11 % (((v11 | v11)) | 1))
                                        if false {
                                            var v43 uint8 = 8
                                            fmt.Println(v11, v29, v2)
                                            _ = v43
                                        }
                                        _ = fi42
                                    }
                                    if (true || (32767 >= (fi12 / ((fi12) | 1)))) {
                                        var v44 uint32 = 65535
                                        var v45 uint = 583
                                        _ = len(v3)
                                        _ = v44
                                        _ = v45
                                    }
                                }
                            case 191:
                                v46 := GenS7{F0: (v11 / (((v11 & (v11 * v11))) | 1)), F1: 2}
                                _ = v46
                            }
                            _ = v29
                        }
                    }
                } else {
                    v11 = int64(v21)
                    switch (fi12 * fi12) {
                    case 32767:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 1:
                        var v47 int8 = 0
                        _ = v47
                    default:
                        switch "foo bar" {
                        case "a":
                            if false {
                                v10 -= (v10 / (((fi12 % ((v10) | 1))) | 1))
                                v11 = v11
                                var v48 bool = (35 != 255)
                                _ = v48
                            } else {
                                _ = len(v1)
                                switch "" {
                                case "hello":
                                    var v49 string = "\n\t"
                                    _ = v49
                                case "\n\t":
                                    switch ("hello" + "hello") {
                                    case "":
                                        if true {
                                            v50 := GenS8{F0: 42}
                                            v21 += v21
                                            v21 = 32767
                                            _ = v50
                                        }
                                    case "a":
                                        switch "\n\t" {
                                        case "foo bar":
                                            _ = len(v3)
                                        case "a":
                                            v51 := GenS2{F0: fi12, F1: (v21 / ((v21) | 1))}
                                            _ = v51
                                        default:
                                            switch (v10 & (^(v10))) {
                                            case 0:
                                                var v52 int16 = v21
                                                _ = v52
                                            case 992:
                                                if true {
                                                    v11 += (-(v11))
                                                    fmt.Println(v2, v2)
                                                    switch "a" {
                                                    case "":
                                                        v21 = 1000
                                                    default:
                                                        switch ("a" + "foo bar") {
                                                        case "":
                                                            var v53 string = (("\n\t" + "a") + "a")
                                                            _ = v53
                                                        }
                                                    }
                                                } else {
                                                    if (!true) {
                                                        var v54 uint16 = 454
                                                        fmt.Println(v11)
                                                        _ = v54
                                                    }
                                                    if true {
                                                        for fi55 := 0; fi55 < 6; fi55++ {
                                                            v2 = v2
                                                            var v56 string = "hello"
                                                            _ = fi55
                                                            _ = v56
                                                        }
                                                    } else {
                                                        v21 = (v21 & 32767)
                                                        var v57 uint = 65535
                                                        if false {
                                                            v2 = ((1.5 / float32((0.0 - 1.5))) - 0.1)
                                                            var v58 int16 = (-(v21))
                                                            _ = v58
                                                        } else {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                if (!(!("hello" <= "foo bar"))) {
                                                                    fmt.Println(v10)
                                                                }
                                                            default:
                                                                if false {
                                                                    if ((0 != 4294967295) || (!false)) {
                                                                        for fi59 := 0; fi59 < 9; fi59++ {
                                                                            _ = len(v1)
                                                                            switch "hello" {
                                                                            case "hello":
                                                                                if (true || (true || true)) {
                                                                                    v11 = -1
                                                                                    v10 = -1
                                                                                }
                                                                                v60 := GenS9{F0: 9223372036854775807, F1: (45 == -90)}
                                                                                if false {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                } else {
                                                                                    var v61 uint16 = 27
                                                                                    var v62 uint = v57
                                                                                    _ = v61
                                                                                    _ = v62
                                                                                }
                                                                                _ = v60
                                                                            default:
                                                                                v63 := GenS10{F0: 65535}
                                                                                _ = v63
                                                                            }
                                                                            v64 := GenS11{F0: (v57 << ((0) & 63)), F1: (fi12 ^ (^(v10)))}
                                                                            _ = fi59
                                                                            _ = v64
                                                                        }
                                                                        switch 9223372036854775807 {
                                                                        case 42:
                                                                            for fi65 := 0; fi65 < 8; fi65++ {
                                                                                if false {
                                                                                    v21 = -1
                                                                                }
                                                                                _ = fi65
                                                                            }
                                                                        default:
                                                                            if false {
                                                                                if (!true) {
                                                                                    fmt.Println(fi12, v2, fi12)
                                                                                    var v66 uint32 = 42
                                                                                    v57 -= v57
                                                                                    _ = v66
                                                                                }
                                                                                fmt.Println(v10, v10, v11)
                                                                                switch 128 {
                                                                                case 1:
                                                                                    if (0 > 164) {
                                                                                        v2 = float32(0.1)
                                                                                    } else {
                                                                                        var v67 bool = true
                                                                                        v68 := GenS12{F0: "hello", F1: v11, F2: 100}
                                                                                        _ = v67
                                                                                        _ = v68
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        var v69 int32 = 300
                                                                        _ = v69
                                                                    }
                                                                } else {
                                                                    for fi70 := 0; fi70 < 18; fi70++ {
                                                                        var v71 float64 = 0.1
                                                                        var v72 bool = ((("" + (("hello" + "\n\t") + "hello")) == "foo bar") && false)
                                                                        _ = fi70
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v57
                                                    }
                                                    v11 = v11
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        if (32767 >= 120) {
                                            if false {
                                                switch "a" {
                                                case "\n\t":
                                                    var v73 uint64 = 729
                                                    _ = v73
                                                }
                                                v11 -= -1
                                                fmt.Println(v10, v2, v2)
                                            }
                                            fmt.Println(v21, v11)
                                            var v74 int = v10
                                            _ = v74
                                        }
                                    default:
                                        var v75 uint8 = 249
                                        _ = v75
                                    }
                                case "foo bar":
                                    v11 = v11
                                default:
                                    v76 := GenS13{F0: 42}
                                    _ = v76
                                }
                            }
                        }
                    }
                }
            case -9223372036854775808:
                v77 := GenS14{F0: 1, F1: -76, F2: true}
                _ = v77
            case -371:
                for fi78 := 0; fi78 < 4; fi78++ {
                    v10 = (^(fi78))
                    _ = fi78
                }
            }
            _ = fi12
            _ = v21
        }
    } else {
        switch (v10 % (((v10 / (((v10 << ((42) & 63))) | 1))) | 1)) {
        case -543:
            v79 := GenS13{F0: 94}
            _ = v79
        case 0:
            _ = len(v1)
        }
    }
    for fi80 := 0; fi80 < 11; fi80++ {
        delete(v1, "hello")
        switch ("hello" + "foo bar") {
        case "foo bar":
            v81 := GenS7{F0: v11, F1: 497}
            _ = v81
        case "":
            var v82 bool = true
            _ = v82
        }
        _ = fi80
    }
    for fi83 := 0; fi83 < 13; fi83++ {
        if (127 != -1) {
            var v84 int = fi83
            if (!false) {
                fmt.Println(v11, v11)
                for fi85 := 0; fi85 < 18; fi85++ {
                    var v86 int32 = 523
                    _ = fi85
                    _ = v86
                }
            }
            var v87 uint32 = 2
            _ = v84
            _ = v87
        }
        _ = fi83
    }
    if false {
        if (true || false) {
            for fi88 := 0; fi88 < 17; fi88++ {
                switch 32768 {
                case 9223372036854775807:
                    v2 += v2
                case -1:
                    v1[("" + "")] = (^(v10))
                case 127:
                    for fi89 := 0; fi89 < 4; fi89++ {
                        v90 := GenS1{F0: 255, F1: 402, F2: 4294967295}
                        switch 4294967295 {
                        case 32767:
                            var v91 uint8 = 128
                            _ = v91
                        }
                        _ = fi89
                        _ = v90
                    }
                }
                switch fi88 {
                case 128:
                    for fi92 := 0; fi92 < 18; fi92++ {
                        v2 = float32(float64(float32(0.1)))
                        _ = len(v3)
                        if true {
                            v10 = -1
                        } else {
                            v2 = (((v2 - ((v2 + -2.25) / (1.5 - v2))) - float32(3.14159)) - 0.1)
                            var v93 float64 = ((3.14159 + 0.0) / 1e10)
                            fmt.Println(fi88, v93, v2)
                            _ = v93
                        }
                        _ = fi92
                    }
                }
                for fi94 := 0; fi94 < 7; fi94++ {
                    v11 -= genFn0()
                    var v95 uint16 = 133
                    v10 = fi88
                    _ = fi94
                    _ = v95
                }
                _ = fi88
            }
            fmt.Println(v10)
        }
        if (-164 >= 1) {
            var v96 int = v10
            if false {
                if (!(255 != 0)) {
                    var v97 int32 = -244
                    for fi98 := 0; fi98 < 16; fi98++ {
                        var v99 uint8 = 2
                        for fi100 := 0; fi100 < 12; fi100++ {
                            var v101 int16 = 279
                            var v102 int = 128
                            var v103 uint64 = 4294967295
                            if true {
                                var v104 float64 = 0.1
                                v96 = v10
                                v96 = 447
                                fmt.Println(v104)
                                fmt.Println(v11, v96, v99)
                                v101 -= v101
                                _ = v104
                            } else {
                                var v105 string = ""
                                var v106 float64 = float64(0.0)
                                _ = v105
                                _ = v106
                            }
                            _ = fi100
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                        _ = fi98
                        _ = v99
                    }
                    _ = v97
                } else {
                    _ = len(v3)
                    v10 -= (v96 ^ (^(v96)))
                    if true {
                        for fi107 := 0; fi107 < 6; fi107++ {
                            if ((1 != 512) || (!(!(!(false || (-32768 == -1)))))) {
                                var v108 int = (v10 ^ fi107)
                                v96 = fi107
                                if (4294967295 != 252) {
                                    fmt.Println(v2)
                                    var v109 uint8 = 255
                                    _ = v109
                                }
                                _ = v108
                            } else {
                                v110 := GenS1{F0: 64, F1: 4294967295, F2: 1}
                                _ = v110
                            }
                            if ((0 >= 2) && true) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v111 := GenS15{F0: 1, F1: -770, F2: -9223372036854775808}
                                var v112 uint = 65535
                                v113 := GenS7{F0: v11, F1: 714}
                                _ = v111
                                _ = v112
                                _ = v113
                            }
                            _ = fi107
                        }
                        v96 = (v10 * v10)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 2147483648
                        }
                        var v114 int8 = 0
                        _ = v114
                    }
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                var v115 float32 = v2
                _ = v115
            }
            _ = v96
        } else {
            _ = len(v1)
            v2 = v2
            var v116 int8 = 1
            _ = v116
        }
        if false {
            var v117 string = ("hello" + ("" + "hello"))
            _ = v117
        } else {
            var v118 int = 0
            v118 += v118
            _ = v118
        }
    }
    if (!false) {
        if ((!false) && ((128 >= v10) && ((((("" + "hello") + ("a" + ("\n\t" + "foo bar"))) + ("hello" + "hello")) == "foo bar") && (995 != 575)))) {
            var v119 float32 = v2
            _ = v119
        }
        v3 = append(v3, 0)
        switch "\n\t" {
        case "":
            var v120 uint8 = 1
            _ = v120
        default:
            for fi121 := 0; fi121 < 18; fi121++ {
                var v122 string = (((("a" + ("foo bar" + "\n\t")) + "foo bar") + ("foo bar" + ("foo bar" + "hello"))) + "")
                v123 := GenS13{F0: 0}
                if ((("a" >= "foo bar") && true) && true) {
                    switch -2147483648 {
                    case -835:
                        if true {
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            var v124 uint16 = 1
                            _ = v124
                        }
                    case -2147483648:
                        for fi125 := 0; fi125 < 14; fi125++ {
                            _ = len(v3)
                            _ = fi125
                        }
                    case 9223372036854775807:
                        var v126 float64 = 1e10
                        _ = v126
                    }
                    fmt.Println(fi121, v10)
                    for fi127 := 0; fi127 < 15; fi127++ {
                        var v128 uint = 65535
                        _ = fi127
                        _ = v128
                    }
                } else {
                    if true {
                        var v129 int = 887
                        v130 := v1[("\n\t" + (((v122 + ("a" + ("hello" + "\n\t"))) + v122) + ((("foo bar" + v122) + "\n\t") + ((v122 + "a") + ("hello" + ((v122 + (v122 + v122)) + (((v122 + ("hello" + v122)) + v122) + ((v122 + "") + "hello"))))))))]
                        v2 -= v2
                        _ = v129
                        _ = v130
                    }
                    var v131 uint8 = 255
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    _ = v131
                }
                _ = fi121
                _ = v122
                _ = v123
            }
        }
    } else {
        v10 -= 79
    }
    _ = len(v3)
    for fi132 := 0; fi132 < 14; fi132++ {
        var v133 uint32 = 2147483647
        v134 := GenS15{F0: -891, F1: (v10 - 4294967295), F2: v11}
        _ = fi132
        _ = v133
        _ = v134
    }
    switch "" {
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v135 uint16 = 0
        _ = v135
    default:
        var v136 uint64 = 18446744073709551615
        _ = v136
    }
    var v137 uint = 4294967295
    if false {
        var v138 int8 = -23
        var v139 uint32 = 504
        if (9223372036854775807 != 423) {
            switch (v10 - (v10 * v10)) {
            case 1:
                var v140 int16 = 1
                _ = v140
            default:
                if true {
                    var v141 float64 = 0.0
                    _ = v141
                }
            }
            v142 := GenS16{F0: 250, F1: 0}
            _ = v142
        } else {
            v2 -= 1e10
            var v143 float32 = v2
            var v144 int16 = int16(v139)
            _ = v143
            _ = v144
        }
        _ = v138
        _ = v139
    } else {
        var v145 uint64 = 4294967295
        var v146 uint16 = uint16(v145)
        delete(v1, "a")
        _ = v145
        _ = v146
    }
    switch 128 {
    case 4294967295:
        var v147 uint = 130
        _ = v147
    case 983:
        for fi148 := 0; fi148 < 16; fi148++ {
            v10 -= fi148
            for fi149 := 0; fi149 < 9; fi149++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch ("foo bar" + "") {
                case "":
                    if len(v3) > 0 {
                        v3[0] = 1
                    }
                default:
                    var v150 float32 = (v2 - v2)
                    _ = v150
                }
                _ = fi149
            }
            fmt.Println(v2, fi148, v11)
            _ = fi148
        }
    default:
        switch "\n\t" {
        case "foo bar":
            if (-1 < 42) {
                for fi151 := 0; fi151 < 14; fi151++ {
                    fmt.Println(v2)
                    if true {
                        delete(v1, "a")
                    } else {
                        if true {
                            for fi152 := 0; fi152 < 18; fi152++ {
                                if (!((!false) || ((v137 ^ (v137 << (((v137 | (v137 >> ((v137) & 63)))) & 63))) == 42))) {
                                    if false {
                                        v10 = (fi152 - v10)
                                        var v153 uint = v137
                                        _ = v153
                                    }
                                    v2 += (v2 * v2)
                                    var v154 int16 = 0
                                    _ = v154
                                } else {
                                    var v155 float32 = 1.5
                                    var v156 uint64 = 1
                                    _ = v155
                                    _ = v156
                                }
                                var v157 uint64 = 42
                                fmt.Println(fi152, v157, fi152)
                                _ = fi152
                                _ = v157
                            }
                            fmt.Println(fi151, v2, fi151)
                            v2 -= float32(3.14159)
                        }
                        v10 += fi151
                        var v158 string = ("hello" + "hello")
                        _ = v158
                    }
                    v3 = append(v3, 4294967295)
                    _ = fi151
                }
                var v159 int16 = -1
                _ = v159
            } else {
                var v160 uint64 = uint64(v137)
                fmt.Println(v10)
                v137 = 1
                _ = v160
            }
        case "hello":
            switch "" {
            case "hello":
                if (v137 == v137) {
                    var v161 bool = false
                    var v162 float64 = ((float64(3.14159) + 1e10) * 0.0)
                    v10 = v10
                    _ = v161
                    _ = v162
                } else {
                    fmt.Println(v2, v11)
                    v2 = 0.1
                    for fi163 := 0; fi163 < 18; fi163++ {
                        for fi164 := 0; fi164 < 3; fi164++ {
                            v165 := GenS17{F0: (fi163 / ((fi163) | 1)), F1: -128, F2: true}
                            _ = len(v3)
                            v166 := GenS1{F0: 234, F1: (v137 >> ((18446744073709551615) & 63)), F2: 14}
                            _ = fi164
                            _ = v165
                            _ = v166
                        }
                        if true {
                            if (!true) {
                                var v167 uint8 = 1
                                fmt.Println(fi163, v10, v2)
                                _ = v167
                            } else {
                                delete(v1, "foo bar")
                            }
                        } else {
                            v1["hello"] = (fi163 ^ (fi163 << ((v137) & 63)))
                            for fi168 := 0; fi168 < 7; fi168++ {
                                v137 -= 18446744073709551615
                                v11 = 9223372036854775807
                                v11 += (-(v11))
                                _ = fi168
                            }
                            v137 -= (v137 & (v137 << (((v137 + 42)) & 63)))
                        }
                        _ = len(v1)
                        _ = fi163
                    }
                }
            case "":
                var v169 int16 = -285
                _ = v169
            case "\n\t":
                v170 := GenS18{F0: (v137 / ((v137) | 1))}
                _ = v170
            default:
                switch (v10 % (((v10 << ((768) & 63))) | 1)) {
                case -627:
                    v171 := GenS17{F0: int(v137), F1: 1, F2: true}
                    _ = v171
                }
            }
        }
    }
    v172 := GenS11{F0: 42, F1: v10}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v137
    _ = v172
    _ = v2
    _ = v3
    _ = v8
}
