package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 string; F1 uint32 }
type GenS8 struct { F0 int32; F1 uint; F2 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 int8 }
func genFn0() int16 {
    var v6 uint16 = 1
    v6 -= (v6 & v6)
    v6 = (v6 >> ((uint(v6)) & 63))
    switch "foo bar" {
    case "\n\t":
        var v7 string = "a"
        _ = v7
    default:
        v8 := GenS1{F0: ("hello" >= ("\n\t" + "a"))}
        _ = v8
    }
    _ = v6
    return 739
}
func genFn1(p0 string, p1 bool, p2 uint32) uint16 {
    var v9 int8 = -1
    v10 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 0
}

func main() {
    var v1 int = 4294967295
    var v2 int16 = 1000
    var v3 int16 = v2
    var v4 uint64 = 634
    v5 := make(map[string]int)
    v11 := GenS1{F0: (!true)}
    v5[(((((((((("a" + "a") + ((("foo bar" + "\n\t") + ((("" + ("" + "a")) + "\n\t") + "foo bar")) + (("\n\t" + "a") + (("a" + ("foo bar" + "foo bar")) + "foo bar")))) + ("hello" + "a")) + (("\n\t" + "foo bar") + "foo bar")) + "\n\t") + "foo bar") + "") + "a") + "foo bar") + (("" + ("\n\t" + "")) + "hello"))] = v1
    switch (v1 - (v1 & -2147483648)) {
    case 127:
        if false {
            v12 := []int{(^(v1)), v1}
            _ = v12
        } else {
            switch 4294967295 {
            case 556:
                for fi13 := 0; fi13 < 13; fi13++ {
                    if true {
                        var v14 int32 = -1
                        for fi15 := 0; fi15 < 9; fi15++ {
                            fmt.Println(fi15)
                            _ = fi15
                        }
                        fmt.Println(fi13)
                        _ = v14
                    } else {
                        var v16 float32 = float32(1.5)
                        v3 -= v2
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 4; fi17++ {
                        v18 := GenS2{F0: 648}
                        if false {
                            switch fi13 {
                            case -314:
                                var v19 uint = 1
                                _ = v19
                            case 1:
                                v20 := GenS2{F0: 4294967295}
                                _ = v20
                            }
                            if true {
                                v21 := GenS3{F0: 1}
                                _ = v21
                            }
                            var v22 uint32 = 2
                            _ = v22
                        } else {
                            fmt.Println(fi17, fi13, v2)
                        }
                        var v23 int = 127
                        v24 := []int{v23, v23}
                        _ = fi17
                        _ = v18
                        _ = v23
                        _ = v24
                    }
                    if ((42 > 255) || (false && (v3 <= (v2 >> ((65535) & 63))))) {
                        v25 := GenS3{F0: 18446744073709551615}
                        if (18446744073709551615 >= 18446744073709551615) {
                            fmt.Println(v3, v2, v4)
                            var v26 int16 = v3
                            v26 -= v26
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 12; fi27++ {
                            if (!(v4 > 42)) {
                                var v28 string = "foo bar"
                                var v29 uint = 4294967295
                                _ = v28
                                _ = v29
                            }
                            _ = fi27
                        }
                        _ = v25
                    }
                    _ = fi13
                }
            }
            if (v4 < (v4 % ((v4) | 1))) {
                if (!(!false)) {
                    v2 = (v2 + -744)
                    for fi30 := 0; fi30 < 18; fi30++ {
                        switch ("hello" + "hello") {
                        case "":
                            fmt.Println(v2)
                        default:
                            v31 := GenS1{F0: ((true || false) && false)}
                            _ = v31
                        }
                        v32 := []int{fi30, 225, -830}
                        _ = fi30
                        _ = v32
                    }
                    v33 := []int{4294967295, v1, v1, (v1 + (v1 | 833))}
                    _ = v33
                }
            } else {
                switch v1 {
                case 128:
                    for fi34 := 0; fi34 < 19; fi34++ {
                        if false {
                            v35 := GenS4{F0: 588, F1: 0, F2: (v4 ^ (v4 ^ (v4 % ((v4) | 1))))}
                            v5[""] = (v1 + (v1 % ((-9223372036854775808) | 1)))
                            var v36 uint16 = 32767
                            _ = v35
                            _ = v36
                        }
                        _ = fi34
                    }
                }
                if false {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        if false {
                            fmt.Println(v3)
                        } else {
                            var v38 string = "hello"
                            for fi39 := 0; fi39 < 5; fi39++ {
                                switch (fi37 * (^(v1))) {
                                case 2147483647:
                                    var v40 string = ("" + v38)
                                    _ = v40
                                }
                                fmt.Println(v1, v4, fi37)
                                v41 := GenS2{F0: 1}
                                _ = fi39
                                _ = v41
                            }
                            var v42 float64 = 1e10
                            _ = v38
                            _ = v42
                        }
                        v43 := GenS5{F0: (("a" + (((("" + ("hello" + "\n\t")) + ("" + ("foo bar" + ("hello" + "foo bar")))) + ("\n\t" + (("" + "hello") + "hello"))) + "foo bar")) + ("" + "")), F1: 9223372036854775807, F2: 1}
                        _ = fi37
                        _ = v43
                    }
                    for fi44 := 0; fi44 < 8; fi44++ {
                        switch (-(v1)) {
                        case 128:
                            v45 := GenS1{F0: (((true || false) && (false || ((!true) || (genFn0() != (v2 >> ((4294967295) & 63)))))) || false)}
                            _ = v45
                        case 32767:
                            switch "a" {
                            case "":
                                var v46 int8 = -1
                                _ = v46
                            case "a":
                                if (!false) {
                                    var v47 float32 = (0.0 - 0.0)
                                    _ = v47
                                }
                            case "foo bar":
                                v48 := GenS6{F0: "hello", F1: (^(v4))}
                                _ = v48
                            default:
                                if (!((!((-(v2)) < genFn0())) || false)) {
                                    if (true || (42 > 0)) {
                                        switch ("foo bar" + "a") {
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "a":
                                                v2 = 1
                                            case "":
                                                fmt.Println(v4, v3, v1)
                                            case "foo bar":
                                                for fi49 := 0; fi49 < 5; fi49++ {
                                                    var v50 int16 = v3
                                                    var v51 int64 = 42
                                                    _ = fi49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        }
                                        for fi52 := 0; fi52 < 14; fi52++ {
                                            v2 = genFn0()
                                            if ((!false) && (!false)) {
                                                v53 := []int{(fi52 + -1), (v1 & (fi44 >> ((255) & 63))), 1, fi52}
                                                var v54 uint8 = 2
                                                v55 := GenS7{F0: "\n\t", F1: 42}
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            v2 -= -32768
                                            _ = fi52
                                        }
                                    }
                                }
                            }
                        default:
                            var v56 uint = 18446744073709551615
                            _ = v56
                        }
                        _ = fi44
                    }
                    var v57 int16 = (v3 % (((v2 - v3)) | 1))
                    _ = v57
                }
            }
        }
    case 648:
        switch v1 {
        case -72:
            fmt.Println(v1, v4)
        case 9223372036854775807:
            v58 := []int{v1, (^(v1)), (v1 - (-(v1))), v1}
            _ = v58
        case -2147483648:
            if (-805 >= 2147483647) {
                for fi59 := 0; fi59 < 13; fi59++ {
                    for fi60 := 0; fi60 < 17; fi60++ {
                        v3 -= genFn0()
                        _ = fi60
                    }
                    _ = fi59
                }
            }
        }
    default:
        v1 = v1
    }
    if false {
        v61 := v5[(("" + ("hello" + "\n\t")) + "hello")]
        if (100 < 1) {
            v62 := []int{}
            _ = v62
        }
        for fi63 := 0; fi63 < 15; fi63++ {
            v64 := []int{v1, 1}
            if false {
                fmt.Println(fi63, v4)
                if false {
                    v1 = (v1 - (v1 - (v61 - (v61 & (v1 - fi63)))))
                }
                fmt.Println(v3, v3)
            } else {
                v3 = (v2 & v3)
                var v65 int8 = -11
                _ = v65
            }
            v1 = 32767
            _ = fi63
            _ = v64
        }
        _ = v61
    } else {
        v3 -= (^(v2))
    }
    for fi66 := 0; fi66 < 13; fi66++ {
        v2 = v3
        _ = fi66
    }
    v67 := GenS5{F0: "a", F1: 0, F2: 603}
    for fi68 := 0; fi68 < 5; fi68++ {
        var v69 int16 = (v2 ^ 32767)
        for fi70 := 0; fi70 < 17; fi70++ {
            if true {
                var v71 float64 = 1e10
                switch v1 {
                case -9223372036854775808:
                    v2 += v69
                case 972:
                    v2 = v3
                case 9223372036854775807:
                    _ = len(v5)
                default:
                    v69 = (v3 % ((v3) | 1))
                }
                _ = v71
            } else {
                v72 := GenS4{F0: 388, F1: 0, F2: 4294967295}
                var v73 float64 = -2.25
                v4 += v4
                _ = v72
                _ = v73
            }
            var v74 int8 = -94
            v4 = v4
            fmt.Println(fi68)
            if true {
                if false {
                    v75 := GenS5{F0: "hello", F1: 9223372036854775807, F2: 668}
                    _ = v75
                }
                switch -669 {
                case 42:
                    var v76 float64 = (float64(3.14159) * -2.25)
                    _ = v76
                case -1:
                    var v77 int = (fi70 ^ fi70)
                    _ = v77
                case 507:
                    v2 += 962
                default:
                    v78 := GenS2{F0: 0}
                    _ = v78
                }
            }
            v79 := GenS3{F0: 1}
            fmt.Println(v2, v1, fi70)
            v69 -= (^(v2))
            v80 := GenS8{F0: -2147483648, F1: 0, F2: 677}
            _ = fi70
            _ = v74
            _ = v79
            _ = v80
        }
        if (false || (!((-831 > -9223372036854775808) && (false || (488 == 0))))) {
            switch (("" + "\n\t") + "\n\t") {
            case "a":
                v4 += (^(v4))
            case "\n\t":
                for fi81 := 0; fi81 < 11; fi81++ {
                    v5[("foo bar" + ((("hello" + "\n\t") + "") + "foo bar"))] = v1
                    var v82 uint8 = 1
                    _ = fi81
                    _ = v82
                }
            }
        } else {
            if false {
                for fi83 := 0; fi83 < 6; fi83++ {
                    switch v1 {
                    case 4294967295:
                        switch "" {
                        case "foo bar":
                            v1 = (^(v1))
                        }
                    default:
                        var v84 uint32 = 65535
                        _ = v84
                    }
                    v4 = (^(v4))
                    _ = fi83
                }
                var v85 uint16 = 32767
                _ = v85
            }
            fmt.Println(v69, fi68, v1)
            if false {
                fmt.Println(v4)
                if true {
                    v86 := []int{(fi68 << ((42) & 63)), (fi68 / ((fi68) | 1)), fi68, (^(v1))}
                    v4 += v4
                    var v87 int32 = 2147483647
                    _ = v86
                    _ = v87
                }
                switch (v1 * 128) {
                case 128:
                    delete(v5, "")
                case 839:
                    var v88 uint16 = 2
                    _ = v88
                default:
                    if (!(true || false)) {
                        if (0.0 < float64((1.5 - (0.1 + -2.25)))) {
                            delete(v5, "hello")
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    fmt.Println(v3)
                                    v69 = (v69 >> ((4294967295) & 63))
                                    var v89 uint16 = 65535
                                    _ = v89
                                }
                            case "a":
                                var v90 int = (v1 << ((255) & 63))
                                _ = v90
                            case "\n\t":
                                var v91 int = v1
                                _ = v91
                            default:
                                var v92 int = 1
                                _ = v92
                            }
                        } else {
                            v3 += 1
                            for fi93 := 0; fi93 < 12; fi93++ {
                                v4 = v4
                                if false {
                                    v3 = v69
                                } else {
                                    if (!(!(!(-83 >= 127)))) {
                                        v94 := []int{}
                                        _ = v94
                                    }
                                }
                                _ = fi93
                            }
                        }
                        switch 127 {
                        case 920:
                            v95 := GenS2{F0: 2147483648}
                            _ = v95
                        case 128:
                            v96 := []int{fi68}
                            _ = v96
                        case 32768:
                            var v97 int = fi68
                            _ = v97
                        }
                    } else {
                        for fi98 := 0; fi98 < 13; fi98++ {
                            var v99 uint32 = 65535
                            v100 := GenS7{F0: ("\n\t" + "\n\t"), F1: v99}
                            v101 := GenS1{F0: true}
                            _ = fi98
                            _ = v100
                            _ = v101
                            _ = v99
                        }
                        if (((1.5 - (0.0 * 0.1)) * -2.25) >= 0.1) {
                            _ = len(v5)
                            v102 := GenS5{F0: (("hello" + "foo bar") + ""), F1: 0, F2: 0}
                            if true {
                                v103 := v5["hello"]
                                var v104 int16 = v3
                                _ = v103
                                _ = v104
                            } else {
                                v69 -= 623
                            }
                            _ = v102
                        } else {
                            for fi105 := 0; fi105 < 11; fi105++ {
                                fmt.Println(fi68, fi68, v1)
                                v69 += v2
                                if false {
                                    if false {
                                        if false {
                                            switch (fi105 ^ v1) {
                                            case 364:
                                                var v106 int8 = 100
                                                _ = v106
                                            case 42:
                                                v5[("\n\t" + "\n\t")] = (v1 | (v1 - (fi68 * (-(fi68)))))
                                            case 1:
                                                v1 -= fi68
                                            default:
                                                var v107 int16 = v69
                                                _ = v107
                                            }
                                            v2 -= v3
                                        } else {
                                            v108 := GenS3{F0: 1}
                                            _ = v108
                                        }
                                        v2 -= genFn0()
                                        fmt.Println(fi105, v4)
                                        var v109 int16 = (v3 % (((v69 >> ((824) & 63))) | 1))
                                        v3 -= v3
                                        _ = v109
                                    }
                                    v5[("\n\t" + "hello")] = fi68
                                } else {
                                    var v110 uint = 1
                                    var v111 int32 = -2147483648
                                    _ = v110
                                    _ = v111
                                }
                                fmt.Println(v2, fi105, v3)
                                _ = fi105
                            }
                            fmt.Println(v2)
                            var v112 bool = (255 <= 168)
                            _ = v112
                        }
                        for fi113 := 0; fi113 < 4; fi113++ {
                            var v114 uint = 255
                            var v115 int16 = (v2 - 1)
                            _ = fi113
                            _ = v114
                            _ = v115
                        }
                    }
                }
            } else {
                _ = len(v5)
                var v116 bool = ((!true) && true)
                switch ("foo bar" + "hello") {
                case "a":
                    v117 := GenS5{F0: "", F1: -9223372036854775808, F2: genFn1("\n\t", (-1 != genFn0()), 843)}
                    _ = v117
                }
                _ = v116
            }
        }
        _ = fi68
        _ = v69
    }
    var v118 uint8 = 183
    if true {
        v1 = (v1 >> ((65535) & 63))
    } else {
        var v119 float64 = (((1e10 - (3.14159 / 3.14159)) / 1.5) * ((1.5 + (3.14159 - 0.0)) / 0.1))
        for fi120 := 0; fi120 < 17; fi120++ {
            if true {
                v1 = (fi120 >> ((42) & 63))
                fmt.Println(v119, fi120)
                v2 = (v2 / (((v3 * 206)) | 1))
            }
            v118 += v118
            for fi121 := 0; fi121 < 2; fi121++ {
                fmt.Println(v1, v118)
                v122 := GenS6{F0: "\n\t", F1: v4}
                _ = fi121
                _ = v122
            }
            _ = fi120
        }
        _ = v119
    }
    var v123 uint32 = 65535
    for fi124 := 0; fi124 < 15; fi124++ {
        v3 -= (v3 & 32767)
        switch "a" {
        case "hello":
            v123 = v123
        case "a":
            switch "a" {
            case "\n\t":
                if (float32(-2.25) <= (((1e10 / 1.0) * (1e10 / 1e10)) / 3.14159)) {
                    var v125 uint16 = 2
                    var v126 uint16 = (^(v125))
                    _ = v125
                    _ = v126
                }
            case "a":
                var v127 uint64 = v4
                _ = v127
            case "hello":
                if ((0.1 + 1.5) < 3.14159) {
                    var v128 string = (("\n\t" + "") + "\n\t")
                    _ = v128
                } else {
                    v129 := GenS3{F0: uint64(v118)}
                    if (1 < -1) {
                        for fi130 := 0; fi130 < 4; fi130++ {
                            switch "" {
                            case "a":
                                if ((900 != 544) && true) {
                                    v131 := GenS9{F0: 255}
                                    v118 = (v118 % (((v118 >> ((440) & 63))) | 1))
                                    v123 -= v123
                                    switch fi124 {
                                    case 32767:
                                        if (!(42 < 254)) {
                                            var v132 int64 = -1
                                            v133 := GenS10{F0: 0}
                                            var v134 uint64 = 225
                                            _ = v132
                                            _ = v133
                                            _ = v134
                                        }
                                    case -1:
                                        v3 = genFn0()
                                    }
                                    _ = v131
                                }
                            case "\n\t":
                                switch 127 {
                                case 461:
                                    v135 := GenS11{F0: "foo bar"}
                                    _ = v135
                                default:
                                    v136 := []int{v1, fi130, (v1 % (((fi124 - (-(fi130)))) | 1)), (v1 % ((fi124) | 1))}
                                    _ = v136
                                }
                            case "":
                                switch (fi124 / ((-788) | 1)) {
                                case 1:
                                    v2 += (v3 * (v3 ^ v3))
                                }
                            }
                            _ = fi130
                        }
                    }
                    _ = v129
                }
            }
        default:
            switch (v1 | fi124) {
            case 0:
                switch ("hello" + "hello") {
                case "hello":
                    delete(v5, ("hello" + "foo bar"))
                case "":
                    var v137 string = ""
                    _ = v137
                case "\n\t":
                    v138 := GenS12{F0: -877, F1: 100}
                    _ = v138
                }
            case -353:
                switch ("hello" + "a") {
                case "foo bar":
                    v4 = (v4 + (v4 << ((2) & 63)))
                case "hello":
                    var v139 uint8 = 224
                    _ = v139
                case "a":
                    switch ("hello" + "hello") {
                    case "\n\t":
                        if (0 <= 9223372036854775807) {
                            for fi140 := 0; fi140 < 7; fi140++ {
                                if (!(true && ((!true) || false))) {
                                    v141 := GenS4{F0: v123, F1: 1, F2: v4}
                                    v1 = 0
                                    if (false && false) {
                                        switch 9223372036854775807 {
                                        case -1:
                                            v118 = 255
                                        case 127:
                                            v118 = (^(v118))
                                        case 17:
                                            switch "" {
                                            case "foo bar":
                                                v3 = genFn0()
                                            default:
                                                if false {
                                                    var v142 int32 = 1
                                                    var v143 uint16 = 65535
                                                    _ = v142
                                                    _ = v143
                                                }
                                            }
                                        }
                                    } else {
                                        v144 := GenS9{F0: 255}
                                        v145 := []int{(^(v1))}
                                        var v146 uint32 = 2
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                    }
                                    v118 += 127
                                    switch 9223372036854775807 {
                                    case 128:
                                        v118 -= v118
                                        var v147 float64 = 0.0
                                        _ = v147
                                    case 1:
                                        var v148 uint16 = 371
                                        _ = v148
                                    case 9223372036854775807:
                                        var v149 string = ("" + ("\n\t" + (("foo bar" + "\n\t") + "\n\t")))
                                        _ = v149
                                    default:
                                        v2 -= (^(v3))
                                    }
                                    _ = v141
                                }
                                var v150 uint16 = genFn1(("foo bar" + "hello"), false, 4294967295)
                                v2 += 42
                                _ = fi140
                                _ = v150
                            }
                            v151 := GenS9{F0: v118}
                            _ = v151
                        } else {
                            v123 += v123
                        }
                    case "hello":
                        var v152 float64 = (0.0 - 0.1)
                        _ = v152
                    default:
                        var v153 uint64 = 42
                        _ = v153
                    }
                }
            }
        }
        v4 += v4
        _ = fi124
    }
    switch (v1 - v1) {
    case -310:
        v3 -= (v3 / ((genFn0()) | 1))
    case -2147483648:
        v154 := []int{int(v118), v1, (-(v1)), (v1 * v1)}
        _ = v154
    case 4294967295:
        var v155 uint8 = (v118 << ((4294967295) & 63))
        _ = v155
    }
    var v156 uint8 = v118
    fmt.Println("v1:", v1)
    fmt.Println("v118:", v118)
    fmt.Println("v123:", v123)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v118
    _ = v123
    _ = v156
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
