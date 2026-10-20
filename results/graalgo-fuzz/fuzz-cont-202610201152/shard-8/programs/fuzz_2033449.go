package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 uint; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
type GenS4 struct { F0 bool; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64; F1 bool; F2 uint64 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 int64; F1 int64; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 uint64 }
type GenS11 struct { F0 uint; F1 int16; F2 int16 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint; F1 bool; F2 int16 }
type GenS14 struct { F0 uint8; F1 int32 }
func genFn0(p0 uint32) string {
    var v6 int8 = 0
    fmt.Println(p0, p0)
    var v7 string = "\n\t"
    _ = p0
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int8) int {
    switch ("foo bar" + "a") {
    case "hello":
        v8 := GenS1{F0: 42, F1: "\n\t", F2: 65535}
        _ = v8
    case "a":
        var v9 uint = 761
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 14; fi10++ {
            switch "foo bar" {
            case "":
                fmt.Println(fi10, fi10)
            case "a":
                switch "" {
                case "foo bar":
                    v11 := []int{fi10, (fi10 & fi10)}
                    _ = v11
                case "a":
                    fmt.Println(fi10)
                case "":
                    var v12 uint8 = 42
                    _ = v12
                default:
                    var v13 int = (-(fi10))
                    _ = v13
                }
            case "foo bar":
                v14 := GenS1{F0: 32767, F1: (("hello" + "hello") + "\n\t"), F2: 65535}
                _ = v14
            }
            var v15 uint16 = 32767
            v16 := []int{fi10, (fi10 - 908)}
            _ = fi10
            _ = v15
            _ = v16
        }
    }
    fmt.Println(p0)
    _ = p0
    return -704
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    var v5 uint8 = 1
    var v17 int8 = 127
    var v18 uint8 = v5
    switch -1 {
    case 128:
        var v19 bool = false
        _ = v19
    case 9223372036854775807:
        if (!true) {
            var v20 int32 = 42
            switch 2147483647 {
            case 9223372036854775807:
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS2{F0: 18446744073709551615, F1: ((((("hello" + "\n\t") + "") + "hello") + ("a" + "\n\t")) != "a"), F2: 908}
                    v23 := []int{}
                    var v24 float64 = ((1.5 + float64(3.14159)) / v3)
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            case 2147483647:
                switch -99 {
                case 4294967295:
                    v17 += v17
                case 42:
                    _ = len(v2)
                }
            }
            _ = v20
        } else {
            switch "foo bar" {
            case "a":
                delete(v1, 32768)
            case "foo bar":
                v3 += v3
            case "hello":
                var v25 float64 = 0.0
                _ = v25
            default:
                var v26 uint32 = 587
                _ = v26
            }
            var v27 int16 = -32768
            _ = v27
        }
    case -9223372036854775808:
        if (-2147483648 >= -802) {
            var v28 uint32 = 2147483648
            v29 := GenS2{F0: 889, F1: (1e10 >= v3), F2: 32767}
            _ = v28
            _ = v29
        } else {
            fmt.Println(v3, v17)
            for fi30 := 0; fi30 < 17; fi30++ {
                switch fi30 {
                case 753:
                    v18 -= v18
                }
                switch (fi30 ^ fi30) {
                case 765:
                    switch genFn1(v17) {
                    case 2147483647:
                        var v31 int32 = -1
                        _ = v31
                    case 32768:
                        v18 = (v18 * (v5 >> ((2) & 63)))
                    case -9223372036854775808:
                        v32 := []int{fi30, (fi30 & genFn1(v17)), 32768, -9223372036854775808}
                        _ = v32
                    }
                case 42:
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v17 = v17
                        switch ("\n\t" + "\n\t") {
                        case "hello":
                            switch ("hello" + "\n\t") {
                            case "hello":
                                v17 = -88
                            case "\n\t":
                                v34 := []int{32767, fi33}
                                _ = v34
                            default:
                                var v35 int = fi33
                                _ = v35
                            }
                        case "foo bar":
                            if false {
                                v18 = 128
                                var v36 uint64 = uint64(v18)
                                var v37 uint8 = 0
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS3{F0: 2, F1: (v5 << ((533) & 63)), F2: "\n\t"}
                                v39 := []int{int(v18), genFn1(-128)}
                                var v40 int16 = 397
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        }
                        v5 = (^(v5))
                        _ = fi33
                    }
                case 1:
                    v5 = (v18 * v18)
                default:
                    v41 := GenS2{F0: 1, F1: (!(-705 > -137)), F2: 904}
                    _ = v41
                }
                _ = fi30
            }
            for fi42 := 0; fi42 < 8; fi42++ {
                v5 += v18
                switch genFn1((v17 >> ((419) & 63))) {
                case -668:
                    switch ("" + ("hello" + "foo bar")) {
                    case "":
                        if ("hello" == "hello") {
                            v43 := GenS4{F0: true, F1: 2147483647, F2: 42}
                            v3 = 1e10
                            if (v17 > v17) {
                                v44 := []int{}
                                fmt.Println(v3)
                                _ = v44
                            }
                            _ = v43
                        }
                    case "foo bar":
                        if true {
                            for fi45 := 0; fi45 < 1; fi45++ {
                                var v46 uint = 65535
                                _ = fi45
                                _ = v46
                            }
                        }
                    case "hello":
                        if false {
                            v5 = v18
                            if true {
                                for fi47 := 0; fi47 < 10; fi47++ {
                                    v48 := GenS5{F0: 193}
                                    if (!false) {
                                        var v49 int32 = -140
                                        _ = v49
                                    } else {
                                        if true {
                                            v3 += v3
                                            v5 -= 128
                                            var v50 uint16 = 2
                                            _ = v50
                                        }
                                    }
                                    var v51 int32 = 731
                                    _ = fi47
                                    _ = v48
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 16; fi52++ {
                                    switch fi52 {
                                    case -1:
                                        switch (fi42 & genFn1(v17)) {
                                        case 9223372036854775807:
                                            _ = len(v2)
                                        case -420:
                                            v18 = v5
                                        }
                                    case 2147483647:
                                        v3 = 1.5
                                    case -66:
                                        v53 := GenS6{F0: -1, F1: true, F2: 2}
                                        _ = v53
                                    }
                                    _ = fi52
                                }
                                v5 += v18
                            } else {
                                var v54 string = "a"
                                v55 := GenS7{F0: (v54 + ((((v54 + (v54 + "hello")) + v54) + v54) + v54)), F1: 42}
                                switch "" {
                                case "foo bar":
                                    v18 -= v5
                                case "":
                                    v56 := GenS1{F0: 820, F1: ("" + ("" + v54)), F2: 0}
                                    _ = v56
                                case "hello":
                                    v3 = v3
                                }
                                _ = v54
                                _ = v55
                            }
                            if (724 > 2) {
                                var v57 float64 = ((v3 * (3.14159 + v3)) - v3)
                                if (539 <= 2147483648) {
                                    for fi58 := 0; fi58 < 9; fi58++ {
                                        var v59 float64 = ((0.1 / 1e10) - v57)
                                        switch "" {
                                        case "foo bar":
                                            if ((0 <= 0) || true) {
                                                switch "hello" {
                                                case "\n\t":
                                                    switch (^(fi58)) {
                                                    case -712:
                                                        switch 445 {
                                                        case 32768:
                                                            v60 := []int{(fi42 << ((18446744073709551615) & 63)), fi42, (fi42 % ((fi42) | 1))}
                                                            _ = v60
                                                        case -2147483648:
                                                            v2["a"] = "a"
                                                        }
                                                    case -1:
                                                        v3 -= ((0.1 - 0.0) / ((float64(0.0) * v3) / (float64(0.0) - v59)))
                                                    case -9223372036854775808:
                                                        v61 := []int{}
                                                        _ = v61
                                                    default:
                                                        v62 := []int{}
                                                        var v63 int16 = 42
                                                        _ = len(v62)
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                case "hello":
                                                    switch ("hello" + ("hello" + "")) {
                                                    case "a":
                                                        switch ("a" + "hello") {
                                                        case "":
                                                            if (2 > 65535) {
                                                                v17 = (v17 << ((31) & 63))
                                                                v64 := []int{fi58}
                                                                switch (("" + "hello") + "a") {
                                                                case "foo bar":
                                                                    if false {
                                                                        v57 = v59
                                                                        v17 = -100
                                                                    }
                                                                case "":
                                                                    v59 += v57
                                                                case "a":
                                                                    switch fi58 {
                                                                    case 42:
                                                                        switch 1 {
                                                                        case 0:
                                                                            var v65 int32 = -1
                                                                            _ = v65
                                                                        }
                                                                    case -77:
                                                                        v57 = v57
                                                                    case -88:
                                                                        var v66 float32 = -2.25
                                                                        _ = v66
                                                                    }
                                                                }
                                                                _ = v64
                                                            } else {
                                                                v3 = v3
                                                                switch fi58 {
                                                                case 1:
                                                                    if (false && false) {
                                                                        v57 += (v3 - v57)
                                                                        if (42 != int16(fi42)) {
                                                                            _ = len(v4)
                                                                            v59 -= v59
                                                                        }
                                                                    }
                                                                default:
                                                                    var v67 uint16 = 32767
                                                                    _ = v67
                                                                }
                                                                v68 := []int{fi58, (fi42 >> ((87) & 63))}
                                                                _ = v68
                                                            }
                                                        case "hello":
                                                            fmt.Println(v59)
                                                        case "a":
                                                            var v69 int64 = 0
                                                            _ = v69
                                                        }
                                                    case "\n\t":
                                                        fmt.Println(v5, v18)
                                                    case "hello":
                                                        if false {
                                                            v70 := []int{fi42, (fi42 / (((fi42 >> ((79) & 63))) | 1)), genFn1(v17)}
                                                            var v71 float32 = 1.5
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    default:
                                                        if (!(109 == 1)) {
                                                            v57 += 1.5
                                                        }
                                                    }
                                                case "a":
                                                    v72 := GenS8{F0: -468, F1: 0, F2: 0}
                                                    _ = v72
                                                default:
                                                    v73 := GenS7{F0: ("hello" + "\n\t"), F1: (^(v18))}
                                                    _ = v73
                                                }
                                            } else {
                                                v74 := []int{fi42}
                                                v59 += v57
                                                fmt.Println(v5)
                                                _ = v74
                                            }
                                        default:
                                            var v75 bool = (int(v17) < fi58)
                                            _ = v75
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                } else {
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        var v77 float64 = (v57 * 0.0)
                                        v77 = (v57 / -2.25)
                                        v78 := GenS8{F0: 42, F1: 0, F2: 4294967295}
                                        v79 := GenS9{F0: -9223372036854775808}
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                    }
                                    switch fi42 {
                                    case -206:
                                        var v80 float64 = (v57 + v57)
                                        _ = v80
                                    case -725:
                                        var v81 float64 = 1e10
                                        _ = v81
                                    default:
                                        var v82 float32 = ((0.0 / (float32(v3) * 1e10)) - 1.5)
                                        _ = v82
                                    }
                                    if ((v17 * (v17 | (v17 % ((v17) | 1)))) <= (v17 >> ((1) & 63))) {
                                        var v83 uint = 18446744073709551615
                                        switch genFn1((v17 / ((v17) | 1))) {
                                        case -1:
                                            for fi84 := 0; fi84 < 18; fi84++ {
                                                var v85 uint16 = 65535
                                                var v86 uint16 = v85
                                                _ = fi84
                                                _ = v85
                                                _ = v86
                                            }
                                        case 32768:
                                            for fi87 := 0; fi87 < 13; fi87++ {
                                                fmt.Println(fi42, v57, fi42)
                                                var v88 string = (("foo bar" + "a") + "hello")
                                                var v89 int16 = 0
                                                _ = fi87
                                                _ = v88
                                                _ = v89
                                            }
                                        case 127:
                                            v90 := GenS10{F0: v17, F1: 4294967295}
                                            _ = v90
                                        default:
                                            if ((^(v83)) <= v83) {
                                                switch fi42 {
                                                case 1:
                                                    v18 += 1
                                                default:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v91 := GenS8{F0: 42, F1: 1, F2: 135}
                                                        _ = v91
                                                    case "\n\t":
                                                        v92 := GenS6{F0: 1, F1: true, F2: 18446744073709551615}
                                                        _ = v92
                                                    }
                                                }
                                                fmt.Println(fi42, v3, v5)
                                                v17 = -71
                                            } else {
                                                v18 -= uint8(v83)
                                            }
                                        }
                                        _ = v83
                                    } else {
                                        v5 -= v5
                                        for fi93 := 0; fi93 < 18; fi93++ {
                                            var v94 uint32 = 1
                                            v95 := v4[(fi93 + fi93)]
                                            v96 := GenS7{F0: "foo bar", F1: v5}
                                            _ = fi93
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        }
                                        if ("\n\t" <= "foo bar") {
                                            if true {
                                                for fi97 := 0; fi97 < 7; fi97++ {
                                                    var v98 string = "foo bar"
                                                    _ = fi97
                                                    _ = v98
                                                }
                                            }
                                            fmt.Println(v18)
                                        }
                                    }
                                }
                                if (!true) {
                                    for fi99 := 0; fi99 < 4; fi99++ {
                                        if false {
                                            var v100 int32 = -2147483648
                                            if true {
                                                v18 -= (v5 | (v5 << ((1) & 63)))
                                            }
                                            _ = v100
                                        }
                                        v101 := []int{(fi42 + fi99), (fi42 >> ((0) & 63))}
                                        var v102 int64 = 931
                                        _ = fi99
                                        _ = v101
                                        _ = v102
                                    }
                                } else {
                                    if (0 <= -1) {
                                        v17 = v17
                                    }
                                    _ = len(v1)
                                }
                                _ = v57
                            }
                        }
                    }
                case 1:
                    v5 -= (v18 & v5)
                default:
                    if (v3 >= v3) {
                        v5 -= v18
                        v5 -= 255
                        if true {
                            var v103 int8 = v17
                            v3 = float64(1e10)
                            if (!(32767 != 0)) {
                                v3 -= (-2.25 + (v3 + 1e10))
                                fmt.Println(v3, v103, v3)
                            }
                            _ = v103
                        }
                    } else {
                        var v104 int = fi42
                        for fi105 := 0; fi105 < 8; fi105++ {
                            if (0 > v17) {
                                fmt.Println(v17, v18, v17)
                            }
                            _ = fi105
                        }
                        fmt.Println(v5, v3)
                        _ = v104
                    }
                }
                delete(v1, -239)
                _ = fi42
            }
        }
    default:
        var v106 int64 = -113
        _ = v106
    }
    v107 := GenS11{F0: 426, F1: 1000, F2: 1}
    for fi108 := 0; fi108 < 6; fi108++ {
        fmt.Println(v18, fi108)
        var v109 uint64 = 4294967295
        if false {
            v3 = v3
            fmt.Println(v109)
            for fi110 := 0; fi110 < 8; fi110++ {
                v2[("hello" + "foo bar")] = "hello"
                v111 := GenS12{F0: (v18 % ((42) | 1))}
                _ = fi110
                _ = v111
            }
        }
        _ = fi108
        _ = v109
    }
    v5 -= (v5 ^ (v5 & (v18 | 1)))
    var v112 uint16 = 32767
    switch (("\n\t" + "a") + "foo bar") {
    case "\n\t":
        fmt.Println(v112)
    case "hello":
        switch 2147483647 {
        case 32767:
            v112 = (v112 + 65535)
        }
    case "a":
        for fi113 := 0; fi113 < 4; fi113++ {
            fmt.Println(v112, v3)
            for fi114 := 0; fi114 < 1; fi114++ {
                var v115 uint8 = 1
                var v116 int32 = -2147483648
                _ = fi114
                _ = v115
                _ = v116
            }
            _ = fi113
        }
    }
    var v117 uint64 = 0
    if false {
        v118 := GenS2{F0: 42, F1: true, F2: (v112 + 0)}
        var v119 float64 = v3
        _ = v118
        _ = v119
    } else {
        switch ("\n\t" + (("hello" + "a") + "")) {
        case "\n\t":
            if (("a" + "hello") == ("a" + ("\n\t" + "a"))) {
                var v120 uint64 = v117
                v18 += uint8(v117)
                var v121 uint64 = v117
                _ = v120
                _ = v121
            } else {
                v17 -= (v17 & (v17 % ((-1) | 1)))
                fmt.Println(v18, v117, v5)
            }
        default:
            if false {
                v18 = 136
            }
        }
    }
    v3 -= float64(((-2.25 * float32(float64(0.1))) + ((float32(v3) / 1.5) - -2.25)))
    if (v112 > (v112 ^ 2)) {
        v122 := GenS7{F0: ("hello" + ("" + "\n\t")), F1: (v5 % ((42) | 1))}
        switch (("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + (("\n\t" + "") + (("hello" + "hello") + "foo bar"))))) {
        case "\n\t":
            v123 := GenS13{F0: 2, F1: (v3 <= float64((0.1 / 0.1))), F2: 42}
            _ = v123
        case "foo bar":
            var v124 string = ("hello" + "hello")
            _ = v124
        case "":
            if true {
                if (18446744073709551615 >= 18446744073709551615) {
                    v125 := []int{}
                    v126 := GenS6{F0: 800, F1: true, F2: (v117 % (((v117 - (v117 << ((2) & 63)))) | 1))}
                    _ = v125
                    _ = v126
                }
                for fi127 := 0; fi127 < 3; fi127++ {
                    var v128 int8 = 100
                    _ = fi127
                    _ = v128
                }
            } else {
                for fi129 := 0; fi129 < 7; fi129++ {
                    v130 := GenS6{F0: 654, F1: true, F2: (v117 * v117)}
                    _ = fi129
                    _ = v130
                }
                switch "foo bar" {
                case "":
                    var v131 bool = false
                    _ = v131
                case "\n\t":
                    v132 := []int{32767}
                    _ = v132
                case "a":
                    switch -1 {
                    case 32768:
                        for fi133 := 0; fi133 < 20; fi133++ {
                            if false {
                                for fi134 := 0; fi134 < 18; fi134++ {
                                    v135 := GenS1{F0: v112, F1: ("" + ("\n\t" + "foo bar")), F2: 313}
                                    v112 -= (v112 * 0)
                                    v112 += (v112 / ((32767) | 1))
                                    _ = fi134
                                    _ = v135
                                }
                                v17 += v17
                            }
                            if (0 < 1000) {
                                v1[fi133] = (false || false)
                            }
                            if (true || false) {
                                fmt.Println(v112)
                            } else {
                                if true {
                                    if (false || ((!false) || true)) {
                                        var v136 int = 42
                                        v137 := []int{(v136 - (-(v136)))}
                                        _ = v136
                                        _ = v137
                                    } else {
                                        v4[(fi133 >> ((42) & 63))] = (0 > -712)
                                        fmt.Println(v117)
                                        v117 += v117
                                    }
                                    if false {
                                        fmt.Println(v17, v18, v117)
                                        if (!(false && false)) {
                                            if true {
                                                fmt.Println(fi133, v117, v17)
                                            }
                                            if ((fi133 | (-(fi133))) >= fi133) {
                                                v138 := GenS12{F0: (v5 * v18)}
                                                switch (fi133 * int(v17)) {
                                                case -727:
                                                    fmt.Println(v18)
                                                default:
                                                    v139 := []int{genFn1(v17), genFn1(v17), genFn1(int8(v112)), genFn1(-32)}
                                                    _ = v139
                                                }
                                                _ = v138
                                            }
                                        } else {
                                            v18 -= 197
                                            v112 -= v112
                                            v17 = int8(v117)
                                        }
                                    } else {
                                        v5 += (v18 ^ (v5 << ((0) & 63)))
                                        var v140 int16 = -313
                                        var v141 float64 = v3
                                        _ = v140
                                        _ = v141
                                    }
                                }
                            }
                            _ = fi133
                        }
                    case -439:
                        var v142 float32 = 1e10
                        _ = v142
                    case 212:
                        for fi143 := 0; fi143 < 12; fi143++ {
                            v112 -= v112
                            switch "hello" {
                            case "a":
                                if true {
                                    var v144 int64 = 1
                                    if false {
                                        if false {
                                            if ((^(v117)) < v117) {
                                                fmt.Println(v17)
                                                var v145 uint32 = 452
                                                _ = v145
                                            } else {
                                                var v146 float64 = v3
                                                var v147 uint64 = 2
                                                _ = v146
                                                _ = v147
                                            }
                                        } else {
                                            v18 = 42
                                            v148 := []int{0}
                                            _ = v148
                                        }
                                    }
                                    _ = v144
                                }
                            case "":
                                fmt.Println(v5)
                            }
                            _ = fi143
                        }
                    }
                }
                v149 := GenS14{F0: v18, F1: 1}
                _ = v149
            }
        }
        _ = v122
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v117:", v117)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v107
    _ = v112
    _ = v117
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
