package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int; F2 int8 }
type GenS10 struct { F0 int32; F1 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 int16; F2 string }
type GenS13 struct { F0 uint64; F1 uint32; F2 string }
type GenS14 struct { F0 uint64; F1 bool; F2 uint32 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 int32) uint32 {
    p0 = p0
    if false {
        v5 := GenS1{F0: 0}
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(fi6, p0)
            if true {
                p0 -= 1
                v7 := GenS2{F0: 18446744073709551615, F1: 0, F2: 32767}
                _ = v7
            }
            _ = fi6
        }
        v8 := GenS1{F0: 32768}
        _ = v5
        _ = v8
    }
    v9 := []int{}
    var v10 uint64 = 42
    for fi11 := 0; fi11 < 10; fi11++ {
        if (v10 > (v10 + v10)) {
            for fi12 := 0; fi12 < 12; fi12++ {
                if (!false) {
                    var v13 bool = (false || (1 != 1))
                    p0 += p0
                    _ = v13
                }
                _ = fi12
            }
            _ = len(v9)
            fmt.Println(p0, fi11)
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            switch "" {
            case "\n\t":
                var v15 int16 = 32767
                _ = v15
            case "a":
                p0 = 42
            default:
                if (161 == 2) {
                    fmt.Println(v10, fi11, p0)
                } else {
                    p0 -= p0
                    v16 := GenS3{F0: 955}
                    _ = v16
                }
            }
            var v17 int = fi11
            v9 = append(v9, (fi14 & fi11))
            if (true && (v10 < 22)) {
                if len(v9) > 0 {
                    v9[0] = fi14
                }
            }
            _ = fi14
            _ = v17
        }
        p0 -= p0
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v9
    return 4294967295
}
func genFn1(p0 int8, p1 string) int32 {
    if ("" <= "\n\t") {
        v18 := []int{42, 127, 9223372036854775807, 4294967295}
        _ = v18
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            p1 = p1
            for fi20 := 0; fi20 < 5; fi20++ {
                var v21 int16 = -1
                p1 = p1
                if (fi20 <= 42) {
                    var v22 float64 = ((3.14159 + 0.1) / (1e10 * (-2.25 / 1.0)))
                    var v23 uint16 = 42
                    p1 = p1
                    _ = v22
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        if (((!false) || (4294967295 == 979)) || (!false)) {
            v24 := GenS1{F0: 128}
            fmt.Println(p1)
            v25 := []int{}
            _ = v24
            _ = v25
        } else {
            p0 = (p0 | (^(p0)))
        }
        if false {
            fmt.Println(p0, p0)
        } else {
            switch "" {
            case "foo bar":
                switch 42 {
                case 32767:
                    v26 := GenS4{F0: 695, F1: 1}
                    _ = v26
                case 32768:
                    var v27 float32 = -2.25
                    _ = v27
                case 790:
                    v28 := GenS5{F0: ("a" + ""), F1: 34, F2: 432}
                    _ = v28
                default:
                    var v29 int8 = p0
                    _ = v29
                }
            default:
                p1 = "\n\t"
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                p1 = "\n\t"
                for fi31 := 0; fi31 < 16; fi31++ {
                    v32 := GenS2{F0: 867, F1: 42, F2: 42}
                    switch (fi31 % (((fi30 % ((fi30) | 1))) | 1)) {
                    case 9223372036854775807:
                        var v33 int = (fi30 << ((255) & 63))
                        var v34 uint16 = 1
                        _ = v33
                        _ = v34
                    case 127:
                        switch fi31 {
                        case 950:
                            v35 := GenS3{F0: 42}
                            _ = v35
                        default:
                            if false {
                                p0 += -44
                                fmt.Println(fi30, p0)
                            } else {
                                if true {
                                    p1 = (p1 + ((("foo bar" + p1) + ("foo bar" + ("a" + p1))) + (p1 + p1)))
                                    fmt.Println(p0)
                                }
                                switch fi30 {
                                case 4294967295:
                                    if (2 == 42) {
                                        var v36 float64 = float64(0.0)
                                        v37 := GenS6{F0: 325}
                                        var v38 string = p1
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                case 127:
                                    v39 := GenS7{F0: false}
                                    _ = v39
                                case 128:
                                    var v40 int32 = -1
                                    _ = v40
                                }
                            }
                        }
                    case 42:
                        p1 = (p1 + "a")
                    default:
                        switch int(p0) {
                        case 32768:
                            var v41 bool = (!(2147483648 > 65535))
                            _ = v41
                        case 128:
                            var v42 uint16 = 65535
                            _ = v42
                        case -2147483648:
                            switch 128 {
                            case 42:
                                if (!(-113 > -1)) {
                                    if (false || (false || false)) {
                                        v43 := []int{}
                                        _ = v43
                                    }
                                    switch (p1 + p1) {
                                    case "foo bar":
                                        p1 = p1
                                    }
                                    if (!((fi30 - fi30) == fi30)) {
                                        p1 = (("hello" + p1) + (p1 + p1))
                                        var v44 uint16 = 2
                                        _ = v44
                                    } else {
                                        if true {
                                            fmt.Println(fi30)
                                            fmt.Println(fi30, p1, p0)
                                        } else {
                                            if (-1 < -9223372036854775808) {
                                                var v45 int = (fi31 / (((^(fi30))) | 1))
                                                fmt.Println(p0, fi31, fi31)
                                                _ = v45
                                            }
                                            v46 := GenS8{F0: 868}
                                            _ = v46
                                        }
                                        p1 = "hello"
                                    }
                                } else {
                                    p0 += (p0 * (p0 << ((0) & 63)))
                                    p0 = (p0 >> ((18446744073709551615) & 63))
                                }
                            case -222:
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(p1, p0, p1)
                                case "a":
                                    v47 := GenS7{F0: true}
                                    _ = v47
                                default:
                                    if (!(243 <= 128)) {
                                        p0 += (p0 ^ -1)
                                        p0 += (p0 | -20)
                                    }
                                }
                            case -9223372036854775808:
                                v48 := []int{(^(fi30)), (fi30 >> ((1) & 63))}
                                _ = v48
                            }
                        default:
                            p1 = (p1 + p1)
                        }
                    }
                    v49 := GenS3{F0: 0}
                    _ = fi31
                    _ = v32
                    _ = v49
                }
                _ = fi30
            }
            if (42 <= 0) {
                v50 := []int{127, 128, -397, -9223372036854775808}
                _ = v50
            } else {
                p0 = 91
                p0 = p0
                for fi51 := 0; fi51 < 3; fi51++ {
                    var v52 float32 = 1.5
                    _ = fi51
                    _ = v52
                }
            }
        }
    }
    if (1 < 1) {
        var v53 int = 128
        if true {
            var v54 int64 = -9223372036854775808
            for fi55 := 0; fi55 < 10; fi55++ {
                var v56 int64 = (v54 & (v54 << ((1) & 63)))
                _ = fi55
                _ = v56
            }
            switch p1 {
            case "foo bar":
                fmt.Println(p0, v54)
            default:
                p1 = p1
            }
            _ = v54
        }
        v57 := GenS2{F0: 334, F1: 0, F2: 836}
        _ = v53
        _ = v57
    } else {
        p0 = (p0 + 42)
        v58 := []int{32768, 32768, 388, 1}
        p0 = 121
        _ = v58
    }
    _ = p0
    _ = p1
    return 841
}

func main() {
    var v1 int = 599
    var v2 uint64 = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    var v59 uint32 = 983
    v60 := GenS9{F0: (true || (65535 > 330)), F1: v1, F2: int8(v2)}
    if false {
        v1 -= (-(v1))
        switch 9223372036854775807 {
        case 4294967295:
            for fi61 := 0; fi61 < 10; fi61++ {
                v1 += (fi61 & fi61)
                var v62 string = ("" + ("a" + "foo bar"))
                v63 := GenS7{F0: (!(!(float32((float64(-2.25) * ((0.0 - 1e10) - 0.1))) >= 1.5)))}
                _ = fi61
                _ = v62
                _ = v63
            }
        case 32768:
            var v64 uint = 1
            _ = v64
        default:
            if (42 <= 950) {
                switch (((("a" + "") + ((("a" + "foo bar") + "hello") + (("hello" + ("foo bar" + ("foo bar" + "foo bar"))) + "foo bar"))) + "foo bar") + ("hello" + "foo bar")) {
                case "\n\t":
                    if (true || true) {
                        var v65 float32 = 1e10
                        _ = v65
                    } else {
                        v2 += v2
                    }
                case "foo bar":
                    fmt.Println(v59, v2, v1)
                default:
                    for fi66 := 0; fi66 < 9; fi66++ {
                        for fi67 := 0; fi67 < 19; fi67++ {
                            if ((false && true) || false) {
                                if true {
                                    var v68 float32 = ((-2.25 / ((0.0 + ((3.14159 + (0.0 + 3.14159)) - (3.14159 - 0.1))) * 3.14159)) * float32(((((1e10 + (0.0 + 1.5)) - 3.14159) - 1.5) + 3.14159)))
                                    _ = v68
                                }
                            }
                            var v69 uint64 = (v2 | 0)
                            v69 = v2
                            var v70 uint64 = 42
                            _ = fi67
                            _ = v69
                            _ = v70
                        }
                        _ = fi66
                    }
                }
            }
        }
    }
    v2 = 18446744073709551615
    if true {
        v71 := GenS10{F0: 1, F1: v59}
        v1 = 42
        var v72 int64 = 42
        _ = v71
        _ = v72
    }
    for fi73 := 0; fi73 < 7; fi73++ {
        if true {
            v2 -= v2
        }
        _ = fi73
    }
    for fi74 := 0; fi74 < 13; fi74++ {
        if ((v59 * v59) != v59) {
            for fi75 := 0; fi75 < 1; fi75++ {
                v76 := []int{fi75, fi74}
                if false {
                    fmt.Println(v59)
                } else {
                    fmt.Println(v59, v2)
                }
                v59 = v59
                _ = fi75
                _ = v76
            }
            v77 := v4["a"]
            var v78 int16 = 205
            _ = v77
            _ = v78
        }
        _ = fi74
    }
    switch v1 {
    case 42:
        v2 += (^(v2))
    }
    for fi79 := 0; fi79 < 15; fi79++ {
        if false {
            var v80 int64 = 238
            v2 = 1
            _ = v80
        }
        if true {
            v81 := GenS11{F0: ""}
            v82 := GenS12{F0: fi79, F1: 1000, F2: "a"}
            _ = v81
            _ = v82
        } else {
            v4[""] = v59
            for fi83 := 0; fi83 < 3; fi83++ {
                if false {
                    var v84 uint64 = v2
                    _ = v84
                }
                _ = fi83
            }
        }
        _ = fi79
    }
    switch ("\n\t" + ("" + "")) {
    case "hello":
        if (-2.25 == -2.25) {
            for fi85 := 0; fi85 < 1; fi85++ {
                v59 = 2147483648
                v86 := []int{v1, -2147483648}
                fmt.Println(fi85, fi85)
                _ = fi85
                _ = v86
            }
            var v87 int = (v1 | -9223372036854775808)
            _ = v87
        }
    case "foo bar":
        if true {
            v88 := GenS4{F0: 255, F1: -723}
            _ = v88
        }
    case "":
        for fi89 := 0; fi89 < 19; fi89++ {
            if (!true) {
                v90 := []int{fi89, fi89, 32767, (-(fi89))}
                if false {
                    v2 = (v2 % ((18446744073709551615) | 1))
                    if (float32((1e10 * 0.0)) == (float32(1e10) * (-2.25 / 1.0))) {
                        var v91 uint32 = v59
                        if false {
                            if true {
                                v92 := GenS2{F0: 4294967295, F1: 65535, F2: 255}
                                v1 += fi89
                                _ = v92
                            } else {
                                for fi93 := 0; fi93 < 13; fi93++ {
                                    if true {
                                        v94 := v3["hello"]
                                        var v95 uint8 = 200
                                        switch 1 {
                                        case -9223372036854775808:
                                            v94 -= 419
                                        case 215:
                                            var v96 float32 = 0.0
                                            delete(v4, "")
                                            v95 = v95
                                            _ = v96
                                        }
                                        _ = v94
                                        _ = v95
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            var v97 string = "foo bar"
                                            _ = v97
                                        default:
                                            var v98 int16 = int16(v2)
                                            _ = v98
                                        }
                                        v99 := GenS13{F0: (v2 - 42), F1: (v59 >> ((379) & 63)), F2: "\n\t"}
                                        _ = v99
                                    }
                                    v2 = v2
                                    if true {
                                        v59 = v91
                                        v100 := GenS10{F0: -547, F1: v91}
                                        _ = v100
                                    }
                                    _ = fi93
                                }
                                switch (("a" + "foo bar") + "\n\t") {
                                case "\n\t":
                                    for fi101 := 0; fi101 < 15; fi101++ {
                                        v102 := GenS14{F0: 272, F1: (!true), F2: (v59 - (^(v91)))}
                                        switch (fi89 << ((0) & 63)) {
                                        case 127:
                                            switch fi101 {
                                            case 32768:
                                                fmt.Println(fi89, v59, v59)
                                            }
                                        }
                                        var v103 int32 = 0
                                        _ = fi101
                                        _ = v102
                                        _ = v103
                                    }
                                case "":
                                    v2 = 0
                                case "foo bar":
                                    v90 = append(v90, int(v59))
                                default:
                                    v3[""] = genFn0(259)
                                }
                                for fi104 := 0; fi104 < 10; fi104++ {
                                    v4["foo bar"] = 2147483648
                                    switch (fi89 << ((255) & 63)) {
                                    case 619:
                                        switch v1 {
                                        case 32768:
                                            fmt.Println(fi104)
                                        default:
                                            var v105 int8 = 127
                                            _ = v105
                                        }
                                    case 0:
                                        if (0.1 > 1e10) {
                                            var v106 int32 = 1
                                            _ = v106
                                        } else {
                                            var v107 int8 = 100
                                            _ = v107
                                        }
                                    default:
                                        var v108 uint16 = 65535
                                        _ = v108
                                    }
                                    v2 -= v2
                                    _ = fi104
                                }
                            }
                            v2 -= v2
                        } else {
                            v2 = 1
                        }
                        if (!(false && (!(false || true)))) {
                            v109 := GenS9{F0: (false && ((v59 & v59) != v59)), F1: (fi89 & fi89), F2: 42}
                            if true {
                                v2 = v2
                                v1 += fi89
                                v59 = (v59 << ((255) & 63))
                            }
                            _ = v109
                        } else {
                            switch "" {
                            case "\n\t":
                                if true {
                                    if (!false) {
                                        if len(v90) >= 2 {
                                            v90 = v90[1:len(v90)]
                                        }
                                        v2 = v2
                                        v59 -= (v59 & (v59 + v91))
                                    } else {
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                v59 = genFn0(-82)
                                            } else {
                                                fmt.Println(v1)
                                            }
                                        default:
                                            if false {
                                                v91 = 2147483648
                                                switch 352 {
                                                case 32767:
                                                    fmt.Println(v1, v2, v59)
                                                }
                                                if (!false) {
                                                    if false {
                                                        for fi110 := 0; fi110 < 7; fi110++ {
                                                            if (false && true) {
                                                                if false {
                                                                    if (!false) {
                                                                        switch (-(fi89)) {
                                                                        case 9223372036854775807:
                                                                            var v111 uint8 = 174
                                                                            _ = v111
                                                                        }
                                                                        v59 = genFn0(genFn1(42, "hello"))
                                                                        var v112 int16 = 32767
                                                                        var v113 int8 = 28
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    v1 += (fi89 | (fi110 / ((fi89) | 1)))
                                                                    _ = len(v4)
                                                                } else {
                                                                    var v114 uint8 = 77
                                                                    _ = v114
                                                                }
                                                            } else {
                                                                fmt.Println(v91)
                                                                switch "\n\t" {
                                                                case "":
                                                                    v1 = fi89
                                                                }
                                                            }
                                                            v1 = -1
                                                            switch ("\n\t" + (("hello" + "") + "a")) {
                                                            case "foo bar":
                                                                v115 := GenS12{F0: fi89, F1: -32768, F2: ("hello" + "foo bar")}
                                                                v59 = (v59 - 65535)
                                                                _ = v115
                                                            case "hello":
                                                                var v116 bool = (true && false)
                                                                _ = v116
                                                            case "a":
                                                                if true {
                                                                    if len(v90) > 0 {
                                                                        v90[0] = (fi89 / ((-2147483648) | 1))
                                                                    }
                                                                }
                                                                fmt.Println(v1, v1, v91)
                                                            default:
                                                                v2 = 4294967295
                                                            }
                                                            _ = fi110
                                                        }
                                                        if (230 >= 30) {
                                                            fmt.Println(fi89, fi89)
                                                        } else {
                                                            v117 := v3[(("\n\t" + "\n\t") + (("\n\t" + ("\n\t" + "foo bar")) + "a"))]
                                                            _ = v117
                                                        }
                                                        if ("a" < "a") {
                                                            var v118 uint = 783
                                                            var v119 uint64 = v2
                                                            var v120 int = 990
                                                            _ = v118
                                                            _ = v119
                                                            _ = v120
                                                        }
                                                    }
                                                    var v121 uint16 = 42
                                                    var v122 bool = true
                                                    _ = v121
                                                    _ = v122
                                                } else {
                                                    v1 += (fi89 % ((v1) | 1))
                                                    v123 := GenS15{F0: 34}
                                                    for fi124 := 0; fi124 < 16; fi124++ {
                                                        v59 = (v91 << ((65535) & 63))
                                                        var v125 uint16 = 2
                                                        fmt.Println(v59)
                                                        if len(v90) >= 2 {
                                                            v90 = v90[1:len(v90)]
                                                        }
                                                        if (!(false || (v1 > fi89))) {
                                                            v125 = (v125 ^ (^(v125)))
                                                        }
                                                        _ = fi124
                                                        _ = v125
                                                    }
                                                    _ = v123
                                                }
                                            } else {
                                                v2 += 18446744073709551615
                                            }
                                        }
                                        for fi126 := 0; fi126 < 18; fi126++ {
                                            v127 := GenS16{F0: (true && false), F1: 4294967295, F2: 128}
                                            v128 := GenS7{F0: true}
                                            _ = fi126
                                            _ = v127
                                            _ = v128
                                        }
                                    }
                                    var v129 uint16 = 67
                                    for fi130 := 0; fi130 < 19; fi130++ {
                                        var v131 int16 = 95
                                        _ = fi130
                                        _ = v131
                                    }
                                    _ = v129
                                }
                            case "foo bar":
                                v1 = v1
                            default:
                                v132 := GenS17{F0: 743}
                                _ = v132
                            }
                            v133 := v3[("" + "hello")]
                            _ = v133
                        }
                        _ = v91
                    }
                } else {
                    if ((v1 - v1) < (-(fi89))) {
                        fmt.Println(v1)
                        v134 := GenS18{F0: 42, F1: v1, F2: (v59 + 4294967295)}
                        fmt.Println(v59, fi89)
                        _ = v134
                    } else {
                        for fi135 := 0; fi135 < 20; fi135++ {
                            var v136 float64 = ((float64(0.0) * 1.5) - 3.14159)
                            v137 := GenS2{F0: 255, F1: 65535, F2: 1}
                            var v138 uint = 249
                            delete(v3, ("hello" + "hello"))
                            _ = fi135
                            _ = v136
                            _ = v137
                            _ = v138
                        }
                    }
                    v139 := GenS11{F0: "foo bar"}
                    _ = v139
                }
                _ = v90
            } else {
                var v140 int64 = 9223372036854775807
                delete(v4, "")
                _ = v140
            }
            if (!(226 != 32767)) {
                var v141 uint8 = 42
                _ = v141
            }
            v2 = v2
            _ = fi89
        }
    default:
        var v142 uint32 = v59
        _ = v142
    }
    var v143 uint32 = (v59 >> ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v143:", v143)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    _ = v1
    _ = v143
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
}
