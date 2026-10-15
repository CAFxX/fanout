package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint64; F1 string; F2 uint8 }
type GenS6 struct { F0 int64; F1 bool; F2 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int8; F2 int16 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint16; F1 uint32 }
type GenS13 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 string = ""
    for fi5 := 0; fi5 < 17; fi5++ {
        switch (fi5 - -1) {
        case -9223372036854775808:
            var v6 uint = 0
            _ = v6
        }
        v7 := GenS1{F0: 980, F1: 1}
        if true {
            v8 := []int{fi5, (fi5 ^ fi5), (fi5 - fi5), (^(fi5))}
            if (3.14159 < 1e10) {
                switch (v4 + "foo bar") {
                case "hello":
                    var v9 float64 = 1e10
                    _ = v9
                }
                var v10 string = v4
                if false {
                    if true {
                        for fi11 := 0; fi11 < 13; fi11++ {
                            if false {
                                var v12 int16 = 42
                                if true {
                                    v12 += v12
                                    var v13 int16 = (v12 << ((18446744073709551615) & 63))
                                    switch v10 {
                                    case "\n\t":
                                        var v14 uint16 = 165
                                        _ = v14
                                    case "":
                                        var v15 uint32 = 979
                                        _ = v15
                                    case "hello":
                                        fmt.Println(v12, fi5, fi5)
                                    default:
                                        fmt.Println(fi5, v4, v13)
                                    }
                                    _ = v13
                                }
                                _ = v12
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            _ = fi11
                        }
                        fmt.Println(v4, fi5)
                        v16 := GenS2{F0: 42, F1: -75}
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            fmt.Println(fi5, fi5)
                            _ = fi17
                        }
                        v18 := GenS2{F0: 2147483647, F1: -42}
                        var v19 uint64 = 4294967295
                        _ = v18
                        _ = v19
                    }
                    v20 := GenS2{F0: 424, F1: 127}
                    _ = v20
                }
                _ = v10
            } else {
                v4 = ("hello" + "")
                if (!(true || (!true))) {
                    switch ((v4 + v4) + v4) {
                    case "foo bar":
                        var v21 uint32 = 2147483647
                        _ = v21
                    case "\n\t":
                        if true {
                            delete(v1, "hello")
                            for fi22 := 0; fi22 < 15; fi22++ {
                                if false {
                                    fmt.Println(fi5)
                                }
                                v4 = ((("foo bar" + v4) + v4) + v4)
                                _ = fi22
                            }
                        }
                    case "":
                        var v23 uint16 = 615
                        _ = v23
                    default:
                        var v24 int8 = 127
                        _ = v24
                    }
                    v25 := GenS3{F0: 49, F1: 0, F2: 131}
                    var v26 uint16 = 1
                    _ = v25
                    _ = v26
                } else {
                    _ = len(v1)
                    _ = len(v8)
                    if false {
                        _ = len(v8)
                        for fi27 := 0; fi27 < 15; fi27++ {
                            var v28 uint16 = 255
                            var v29 uint16 = (v28 / ((42) | 1))
                            v29 -= (v29 + 1)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                        v4 = ""
                    }
                }
                for fi30 := 0; fi30 < 8; fi30++ {
                    v31 := GenS2{F0: 2147483647, F1: -53}
                    switch fi5 {
                    case 127:
                        if (!true) {
                            var v32 float64 = 0.0
                            if true {
                                v8 = append(v8, (fi30 << ((1) & 63)))
                                switch (^(fi30)) {
                                case -2147483648:
                                    var v33 int8 = 100
                                    _ = v33
                                case 1:
                                    v32 = 0.0
                                default:
                                    fmt.Println(v32)
                                }
                                _ = len(v8)
                            }
                            _ = v32
                        }
                    case -2147483648:
                        if false {
                            v34 := GenS2{F0: -1, F1: -118}
                            _ = v34
                        }
                        v35 := v3[(-(fi5))]
                        _ = v35
                    case -150:
                        v8 = append(v8, 1)
                    }
                    if (("foo bar" + "\n\t") <= ((v4 + "") + "a")) {
                        if (987 >= -403) {
                            var v36 bool = (42 >= (fi5 % ((fi30) | 1)))
                            fmt.Println(fi30, v36)
                            if v36 {
                                var v37 int16 = 42
                                _ = v37
                            }
                            _ = v36
                        } else {
                            v4 = v4
                            var v38 int64 = -9223372036854775808
                            var v39 string = "\n\t"
                            _ = v38
                            _ = v39
                        }
                    } else {
                        switch ("\n\t" + v4) {
                        case "hello":
                            var v40 int16 = 1
                            _ = v40
                        case "a":
                            v4 = "foo bar"
                        }
                    }
                    _ = fi30
                    _ = v31
                }
            }
            _ = v8
        } else {
            var v41 float32 = (-2.25 * (0.0 / float32(-2.25)))
            for fi42 := 0; fi42 < 1; fi42++ {
                var v43 int8 = 127
                if (true && (201 < 1)) {
                    var v44 uint32 = 428
                    v45 := []int{}
                    fmt.Println(v44, v43, fi5)
                    _ = v44
                    _ = v45
                }
                _ = fi42
                _ = v43
            }
            switch (fi5 >> ((65535) & 63)) {
            case 4294967295:
                v46 := []int{fi5, (fi5 | 682)}
                _ = v46
            default:
                v41 += v41
            }
            _ = v41
        }
        _ = fi5
        _ = v7
    }
    if (128 >= 1) {
        var v47 float32 = 0.1
        if true {
            v47 = (((float32(1e10) / ((3.14159 * (0.0 * v47)) - float32(-2.25))) + v47) + v47)
        } else {
            if true {
                fmt.Println(v4)
                if true {
                    if false {
                        v4 = (v4 + ((("hello" + v4) + (v4 + v4)) + v4))
                        var v48 string = "foo bar"
                        switch v4 {
                        case "hello":
                            if (true || ((100 >= 0) || false)) {
                                if true {
                                    v49 := []int{2147483647}
                                    _ = v49
                                }
                                var v50 uint32 = 65535
                                _ = v50
                            } else {
                                _ = len(v3)
                                if ((v47 - v47) == 1e10) {
                                    v47 += (v47 + (v47 - (3.14159 * 1.5)))
                                    v51 := GenS4{F0: 4294967295, F1: (0 != 0)}
                                    _ = v51
                                } else {
                                    if false {
                                        v47 = (0.1 * (1e10 + v47))
                                    }
                                }
                                v3[32768] = 4294967295
                            }
                        case "foo bar":
                            for fi52 := 0; fi52 < 17; fi52++ {
                                v53 := []int{-9223372036854775808, 128, (fi52 | 2147483647), fi52}
                                _ = fi52
                                _ = v53
                            }
                        }
                        _ = v48
                    }
                    var v54 uint16 = 0
                    _ = v54
                } else {
                    v55 := GenS1{F0: 475, F1: 649}
                    _ = v55
                }
                v4 = v4
            } else {
                if true {
                    fmt.Println(v4)
                } else {
                    v47 = v47
                    for fi56 := 0; fi56 < 9; fi56++ {
                        v57 := []int{(fi56 % (((fi56 % (((fi56 - (fi56 ^ 761))) | 1))) | 1)), fi56, 32768, -9223372036854775808}
                        v58 := GenS5{F0: 42, F1: v4, F2: 68}
                        var v59 float64 = 0.1
                        _ = fi56
                        _ = v57
                        _ = v58
                        _ = v59
                    }
                }
                if false {
                    if true {
                        v4 = "\n\t"
                        switch (v4 + (((v4 + "foo bar") + v4) + ((v4 + "hello") + ""))) {
                        case "a":
                            v4 = v4
                        case "hello":
                            if true {
                                delete(v2, 389)
                                switch v4 {
                                case "hello":
                                    if true {
                                        v60 := []int{}
                                        if false {
                                            v4 = (v4 + "hello")
                                            if false {
                                                v61 := GenS3{F0: -1, F1: 0, F2: 107}
                                                var v62 float64 = 1e10
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                        _ = v60
                                    }
                                }
                            } else {
                                v47 = ((((float32(1e10) + v47) + ((v47 + v47) - v47)) - (v47 + ((v47 + 3.14159) - v47))) - (3.14159 / ((-2.25 + v47) / (1e10 - ((1e10 * float32(0.0)) * (v47 * 0.1))))))
                                v63 := v3[-725]
                                v64 := GenS4{F0: 0, F1: false}
                                _ = v63
                                _ = v64
                            }
                        }
                    } else {
                        fmt.Println(v4, v4)
                        switch v4 {
                        case "":
                            v4 = v4
                        default:
                            var v65 int32 = 2147483647
                            _ = v65
                        }
                    }
                } else {
                    switch (v4 + (("" + "") + "foo bar")) {
                    case "hello":
                        v4 = ((v4 + v4) + ("hello" + "\n\t"))
                    case "foo bar":
                        for fi66 := 0; fi66 < 6; fi66++ {
                            for fi67 := 0; fi67 < 4; fi67++ {
                                var v68 int32 = 2147483647
                                switch fi66 {
                                case 32768:
                                    var v69 uint64 = 2
                                    _ = v69
                                }
                                _ = fi67
                                _ = v68
                            }
                            v70 := GenS6{F0: 0, F1: (-658 >= 42), F2: -9223372036854775808}
                            _ = fi66
                            _ = v70
                        }
                    case "":
                        v71 := GenS5{F0: 239, F1: v4, F2: 0}
                        _ = v71
                    default:
                        v47 = 0.0
                    }
                    var v72 string = v4
                    for fi73 := 0; fi73 < 16; fi73++ {
                        _ = len(v2)
                        var v74 uint16 = 2
                        _ = fi73
                        _ = v74
                    }
                    _ = v72
                }
                for fi75 := 0; fi75 < 14; fi75++ {
                    if false {
                        v4 = "\n\t"
                        for fi76 := 0; fi76 < 8; fi76++ {
                            var v77 uint16 = 725
                            _ = fi76
                            _ = v77
                        }
                        if (true && ("" < "")) {
                            if false {
                                if (((true || (9223372036854775807 <= -9223372036854775808)) && (false || (1 > 0))) || ("foo bar" < ((v4 + v4) + "foo bar"))) {
                                    var v78 uint = 4294967295
                                    for fi79 := 0; fi79 < 8; fi79++ {
                                        v3[fi75] = 767
                                        v80 := []int{fi75, -459, fi75}
                                        _ = fi79
                                        _ = v80
                                    }
                                    v47 = v47
                                    _ = v78
                                } else {
                                    for fi81 := 0; fi81 < 6; fi81++ {
                                        v82 := GenS7{F0: v4}
                                        var v83 uint16 = 42
                                        v84 := GenS1{F0: 2, F1: (v83 ^ 65535)}
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                        _ = v84
                                    }
                                    v85 := GenS2{F0: 627, F1: -65}
                                    _ = v85
                                }
                                v47 += 1.5
                            } else {
                                var v86 uint = 18446744073709551615
                                _ = v86
                            }
                            for fi87 := 0; fi87 < 4; fi87++ {
                                v4 = v4
                                v4 = ""
                                v88 := GenS8{F0: ("" + ("hello" + v4)), F1: 127, F2: -1}
                                _ = fi87
                                _ = v88
                            }
                        }
                    } else {
                        if false {
                            for fi89 := 0; fi89 < 1; fi89++ {
                                fmt.Println(fi89, v4)
                                var v90 uint8 = 255
                                if (v90 > 42) {
                                    v90 = (v90 / ((2) | 1))
                                    v47 = float32(-2.25)
                                } else {
                                    _ = len(v3)
                                    var v91 float64 = -2.25
                                    if (((!(100 >= 1)) && true) || ((v47 / ((v47 + (v47 - float32((0.1 * v91)))) / v47)) < v47)) {
                                        v47 += 0.0
                                    }
                                    switch "" {
                                    case "":
                                        if true {
                                            var v92 uint16 = 65535
                                            var v93 uint32 = 1
                                            switch fi75 {
                                            case -333:
                                                switch "\n\t" {
                                                case "a":
                                                    fmt.Println(v4, v4)
                                                }
                                            case 2147483647:
                                                switch fi75 {
                                                case 2147483647:
                                                    if ((4294967295 <= 42) && (42 == -1)) {
                                                        fmt.Println(fi75, fi75, fi75)
                                                        fmt.Println(fi89, v4)
                                                    }
                                                }
                                            default:
                                                v94 := GenS9{F0: ((v90 < 0) && ((!(int64(v90) > 552)) || ((fi89 >> ((uint(v90)) & 63)) < (fi89 << ((42) & 63))))), F1: 0}
                                                _ = v94
                                            }
                                            _ = v92
                                            _ = v93
                                        }
                                    }
                                    _ = v91
                                }
                                _ = fi89
                                _ = v90
                            }
                            switch v4 {
                            case "hello":
                                for fi95 := 0; fi95 < 17; fi95++ {
                                    if (!true) {
                                        switch (("foo bar" + "a") + v4) {
                                        case "":
                                            switch (fi95 >> ((730) & 63)) {
                                            case 114:
                                                v96 := GenS9{F0: true, F1: 352}
                                                _ = v96
                                            case 32768:
                                                if true {
                                                    if true {
                                                        if (1000 == -163) {
                                                            var v97 string = ("\n\t" + ((v4 + "") + "hello"))
                                                            var v98 int = fi75
                                                            v98 += fi75
                                                            delete(v1, "foo bar")
                                                            _ = v97
                                                            _ = v98
                                                        } else {
                                                            v99 := GenS7{F0: ("" + v4)}
                                                            _ = len(v2)
                                                            var v100 uint8 = 2
                                                            _ = v100
                                                            _ = v99
                                                        }
                                                    } else {
                                                        var v101 int32 = 0
                                                        v102 := GenS7{F0: "a"}
                                                        var v103 int32 = 1
                                                        _ = v101
                                                        _ = v102
                                                        _ = v103
                                                    }
                                                    delete(v2, 4294967295)
                                                } else {
                                                    v104 := GenS4{F0: 10, F1: (-23 > 100)}
                                                    if false {
                                                        if false {
                                                            switch v4 {
                                                            case "":
                                                                v105 := GenS2{F0: 0, F1: 1}
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    v106 := GenS6{F0: 1, F1: (255 < 4294967295), F2: (fi75 * 164)}
                                                                    _ = v106
                                                                case "a":
                                                                    switch 0 {
                                                                    case 1:
                                                                        fmt.Println(v4, v47)
                                                                        v47 = v47
                                                                    case -9223372036854775808:
                                                                        if false {
                                                                            v47 = ((v47 + 1e10) - v47)
                                                                            var v107 int64 = 0
                                                                            v47 = (-2.25 / 1.0)
                                                                            _ = v107
                                                                        } else {
                                                                            v108 := GenS2{F0: 42, F1: -1}
                                                                            v109 := GenS3{F0: 1, F1: 4294967295, F2: 255}
                                                                            _ = v108
                                                                            _ = v109
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v110 uint16 = 42
                                                                    _ = v110
                                                                }
                                                                v47 = v47
                                                                _ = v105
                                                            case "hello":
                                                                v111 := []int{fi95}
                                                                _ = v111
                                                            }
                                                            switch 9223372036854775807 {
                                                            case -9223372036854775808:
                                                                v112 := []int{}
                                                                _ = v112
                                                            }
                                                        } else {
                                                            var v113 int = -37
                                                            v47 = (float32(1.5) * v47)
                                                            _ = v113
                                                        }
                                                        v4 = v4
                                                        if true {
                                                            v47 = v47
                                                            var v114 int32 = -341
                                                            v47 = 1.5
                                                            _ = v114
                                                        } else {
                                                            var v115 uint16 = 42
                                                            _ = v115
                                                        }
                                                    } else {
                                                        fmt.Println(v47, v4, v4)
                                                        v4 = v4
                                                        v47 += (v47 + 0.1)
                                                    }
                                                    if (42 >= 777) {
                                                        if true {
                                                            switch v4 {
                                                            case "a":
                                                                v47 += v47
                                                            }
                                                            v116 := GenS1{F0: 255, F1: 42}
                                                            _ = v116
                                                        }
                                                    } else {
                                                        switch 283 {
                                                        case 1:
                                                            v117 := []int{-466, fi95}
                                                            _ = v117
                                                        }
                                                        var v118 int = 0
                                                        if true {
                                                            if (true && (false || true)) {
                                                                switch v4 {
                                                                case "\n\t":
                                                                    var v119 uint64 = 4294967295
                                                                    _ = v119
                                                                default:
                                                                    var v120 int8 = 0
                                                                    _ = v120
                                                                }
                                                                v4 = "\n\t"
                                                                if (0.1 > -2.25) {
                                                                    v121 := GenS8{F0: v4, F1: -101, F2: 32767}
                                                                    var v122 int8 = 127
                                                                    v123 := []int{127}
                                                                    _ = v121
                                                                    _ = v122
                                                                    _ = v123
                                                                }
                                                                v124 := GenS4{F0: 2, F1: (1 >= 140)}
                                                                if true {
                                                                    v4 = "a"
                                                                    fmt.Println(v118, v118)
                                                                    _ = len(v1)
                                                                } else {
                                                                    v125 := GenS2{F0: -1, F1: 1}
                                                                    var v126 int32 = -2147483648
                                                                    _ = v125
                                                                    _ = v126
                                                                }
                                                                _ = v124
                                                            } else {
                                                                switch v4 {
                                                                case "foo bar":
                                                                    var v127 int8 = -128
                                                                    _ = v127
                                                                case "hello":
                                                                    fmt.Println(fi95)
                                                                case "\n\t":
                                                                    v4 = (v4 + "a")
                                                                default:
                                                                    var v128 int32 = 2147483647
                                                                    _ = v128
                                                                }
                                                                v129 := GenS10{F0: (!false)}
                                                                _ = v129
                                                            }
                                                            v130 := []int{(^(v118)), fi75, fi95}
                                                            _ = v130
                                                        }
                                                        _ = v118
                                                    }
                                                    _ = v104
                                                }
                                            case 440:
                                                fmt.Println(v4, fi95, v4)
                                            }
                                        case "\n\t":
                                            var v131 int64 = 9223372036854775807
                                            _ = v131
                                        case "a":
                                            var v132 uint8 = 255
                                            _ = v132
                                        default:
                                            var v133 bool = true
                                            _ = v133
                                        }
                                        switch v4 {
                                        case "\n\t":
                                            v47 += ((float32(0.0) / 1.0) / (v47 * 0.0))
                                        case "":
                                            switch (v4 + v4) {
                                            case "hello":
                                                v134 := v3[(fi75 & (^(fi75)))]
                                                _ = v134
                                            case "foo bar":
                                                var v135 int32 = 1
                                                _ = v135
                                            }
                                        }
                                    } else {
                                        v47 = v47
                                        v47 = 1.5
                                        v136 := v3[32768]
                                        _ = v136
                                    }
                                    _ = fi95
                                }
                            default:
                                if false {
                                    v47 += (0.1 * v47)
                                    switch ("a" + (((("hello" + v4) + ((((v4 + "\n\t") + v4) + "\n\t") + v4)) + v4) + "")) {
                                    case "foo bar":
                                        _ = len(v3)
                                    case "\n\t":
                                        fmt.Println(fi75)
                                    default:
                                        if false {
                                            fmt.Println(fi75, fi75)
                                            var v137 float64 = (0.1 + (float64((v47 * 1.5)) * (float64(0.1) * -2.25)))
                                            var v138 int16 = 0
                                            _ = v137
                                            _ = v138
                                        }
                                    }
                                    v47 -= v47
                                }
                            }
                        }
                        v4 = (v4 + "a")
                        v139 := []int{(fi75 / ((fi75) | 1)), 128}
                        _ = v139
                    }
                    v47 -= 0.1
                    v47 = v47
                    _ = fi75
                }
            }
            var v140 float32 = (3.14159 - float32(3.14159))
            v141 := GenS11{F0: 127}
            _ = v140
            _ = v141
        }
        var v142 uint8 = 255
        _ = v142
        _ = v47
    } else {
        switch (v4 + ((v4 + v4) + "hello")) {
        case "a":
            if (9223372036854775807 != 127) {
                var v143 float64 = ((0.0 / (0.0 - (-2.25 * -2.25))) / 1e10)
                _ = v143
            }
        case "hello":
            v144 := GenS12{F0: 1, F1: 2}
            _ = v144
        }
        switch 42 {
        case 535:
            var v145 uint32 = 65535
            _ = v145
        case 32768:
            if true {
                v4 = (v4 + "")
                for fi146 := 0; fi146 < 20; fi146++ {
                    for fi147 := 0; fi147 < 9; fi147++ {
                        switch fi146 {
                        case 0:
                            v4 = (("foo bar" + v4) + v4)
                        default:
                            v148 := []int{(fi147 % (((^(fi147))) | 1))}
                            _ = v148
                        }
                        v4 = v4
                        var v149 int = (fi146 % ((fi147) | 1))
                        _ = fi147
                        _ = v149
                    }
                    var v150 float64 = (((1.5 / 1.0) + 0.0) * 1e10)
                    var v151 uint8 = 236
                    _ = fi146
                    _ = v150
                    _ = v151
                }
            } else {
                v152 := GenS10{F0: (false && true)}
                _ = v152
            }
        case -9223372036854775808:
            for fi153 := 0; fi153 < 11; fi153++ {
                for fi154 := 0; fi154 < 8; fi154++ {
                    v155 := []int{(fi154 >> ((4294967295) & 63))}
                    var v156 int16 = -32768
                    _ = fi154
                    _ = v155
                    _ = v156
                }
                _ = fi153
            }
        }
        if (1.5 >= -2.25) {
            for fi157 := 0; fi157 < 9; fi157++ {
                switch v4 {
                case "a":
                    v158 := v3[(fi157 - fi157)]
                    _ = v158
                case "":
                    v4 = (v4 + "\n\t")
                case "foo bar":
                    v159 := GenS9{F0: false, F1: 42}
                    _ = v159
                default:
                    v160 := []int{fi157}
                    _ = v160
                }
                v4 = "hello"
                switch fi157 {
                case -1:
                    v4 = v4
                case 660:
                    if true {
                        var v161 int16 = -853
                        switch v4 {
                        case "\n\t":
                            for fi162 := 0; fi162 < 4; fi162++ {
                                var v163 float64 = 1e10
                                var v164 uint8 = 1
                                _ = fi162
                                _ = v163
                                _ = v164
                            }
                        default:
                            fmt.Println(v4, fi157)
                        }
                        fmt.Println(v161)
                        _ = v161
                    } else {
                        _ = len(v1)
                        fmt.Println(fi157)
                    }
                case -166:
                    for fi165 := 0; fi165 < 16; fi165++ {
                        if (true || (239 < 1)) {
                            var v166 bool = true
                            var v167 uint = 483
                            _ = v166
                            _ = v167
                        } else {
                            switch v4 {
                            case "hello":
                                v168 := GenS11{F0: 255}
                                _ = v168
                            default:
                                var v169 int8 = -57
                                _ = v169
                            }
                        }
                        if (true || true) {
                            v4 = (v4 + (v4 + ((v4 + v4) + (v4 + "hello"))))
                        }
                        _ = fi165
                    }
                default:
                    switch fi157 {
                    case 128:
                        v4 = ""
                    case -972:
                        fmt.Println(v4, fi157)
                    default:
                        v4 = ""
                    }
                }
                _ = fi157
            }
            fmt.Println(v4)
        }
    }
    v4 = "foo bar"
    switch 0 {
    case 112:
        var v170 int8 = 0
        _ = v170
    case 127:
        fmt.Println(v4)
    }
    v4 = v4
    fmt.Println(v4)
    v171 := []int{858, 32768}
    fmt.Println(v4)
    for fi172 := 0; fi172 < 10; fi172++ {
        var v173 uint32 = 2147483647
        _ = fi172
        _ = v173
    }
    v4 = (v4 + ((v4 + v4) + ""))
    v174 := GenS13{F0: 0, F1: 0}
    var v175 uint32 = 2147483648
    var v176 bool = false
    var v177 int = 928
    var v178 int32 = 790
    v176 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v171:", len(v171), v171)
    fmt.Println("v175:", v175)
    fmt.Println("v176:", v176)
    fmt.Println("v177:", v177)
    fmt.Println("v178:", v178)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v171
    _ = v174
    _ = v175
    _ = v176
    _ = v177
    _ = v178
    _ = v2
    _ = v3
    _ = v4
}
