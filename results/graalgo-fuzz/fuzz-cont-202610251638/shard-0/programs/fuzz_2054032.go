package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint32 }
type GenS9 struct { F0 uint32; F1 bool; F2 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64; F1 int64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 bool; F1 int8; F2 int32 }
type GenS15 struct { F0 int; F1 uint16 }
type GenS16 struct { F0 uint; F1 uint8 }
type GenS17 struct { F0 int64; F1 uint32; F2 int32 }
type GenS18 struct { F0 int32; F1 int32; F2 int16 }
type GenS19 struct { F0 uint64; F1 uint8; F2 int }
func genFn0() int {
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int16 = 0
            _ = fi5
            _ = v6
        }
        var v7 bool = (!(127 > fi4))
        v8 := GenS1{F0: 128, F1: -393}
        _ = fi4
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS2{F0: 492, F1: 0, F2: 255}
        var v11 float32 = 3.14159
        _ = fi9
        _ = v10
        _ = v11
    }
    if (1 < 304) {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 bool = (128 <= 2)
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 int8 = 1
                v16 := []int{}
                if true {
                    if (v13 && (int(v15) > (^(fi12)))) {
                        v13 = (0.0 != (float32(float64(-2.25)) * -2.25))
                    }
                    v17 := GenS1{F0: 146, F1: 241}
                    switch (fi12 << ((0) & 63)) {
                    case 128:
                        v13 = (!true)
                    case 2147483647:
                        switch (fi12 & 40) {
                        case -168:
                            v18 := GenS3{F0: ""}
                            if (65535 == 32767) {
                                if (0 == 789) {
                                    var v19 float32 = -2.25
                                    _ = v19
                                } else {
                                    _ = len(v16)
                                }
                                v20 := GenS3{F0: "foo bar"}
                                _ = v20
                            }
                            var v21 uint16 = 826
                            _ = v18
                            _ = v21
                        default:
                            v13 = (!(v13 || (4294967295 > 255)))
                        }
                    default:
                        switch fi12 {
                        case -701:
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        default:
                            v22 := GenS3{F0: ("hello" + ("foo bar" + ("foo bar" + "\n\t")))}
                            _ = v22
                        }
                    }
                    _ = v17
                }
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    } else {
        var v23 int16 = 1
        v24 := []int{960, -2147483648}
        _ = v23
        _ = v24
    }
    var v25 uint16 = 32767
    var v26 int64 = 9223372036854775807
    _ = v25
    _ = v26
    return 2147483647
}
func genFn1(p0 bool, p1 uint64) string {
    if (0 != 127) {
        fmt.Println(p1)
        v27 := []int{}
        var v28 int32 = 0
        _ = v27
        _ = v28
    } else {
        v29 := GenS1{F0: 192, F1: 9223372036854775807}
        fmt.Println(p0, p0)
        _ = v29
    }
    p0 = p0
    v30 := GenS3{F0: "a"}
    if false {
        p0 = false
        var v31 uint = 255
        for fi32 := 0; fi32 < 7; fi32++ {
            v33 := []int{genFn0(), (fi32 % ((fi32) | 1))}
            var v34 bool = p0
            for fi35 := 0; fi35 < 16; fi35++ {
                v36 := GenS1{F0: 255, F1: -9223372036854775808}
                if v34 {
                    var v37 uint32 = 1
                    var v38 float32 = 1.5
                    _ = v37
                    _ = v38
                } else {
                    var v39 uint32 = 77
                    v34 = ((0.1 >= 0.1) && ((-2147483648 <= 0) || false))
                    var v40 int8 = int8(v31)
                    _ = v39
                    _ = v40
                }
                p0 = v34
                _ = fi35
                _ = v36
            }
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = v31
    } else {
        v41 := GenS1{F0: 229, F1: 471}
        _ = v41
    }
    _ = p0
    _ = p1
    _ = v30
    return ""
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{("a" + ""), (v1 + v1)}
    var v3 int8 = 95
    for fi42 := 0; fi42 < 5; fi42++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch genFn0() {
        case 127:
            var v43 string = (((((v1 + (v1 + "a")) + ("foo bar" + ("a" + (v1 + v1)))) + v1) + "\n\t") + "foo bar")
            _ = v43
        case 9223372036854775807:
            for fi44 := 0; fi44 < 1; fi44++ {
                if (true || (42 >= 4294967295)) {
                    v3 = int8(fi42)
                } else {
                    var v45 uint = 255
                    _ = v45
                }
                _ = fi44
            }
        }
        _ = fi42
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        switch v1 {
        case "a":
            switch 330 {
            case 652:
                var v47 uint64 = 18446744073709551615
                _ = v47
            case 2147483647:
                for fi48 := 0; fi48 < 4; fi48++ {
                    v3 += 123
                    if (!(("a" + v1) > (v1 + v1))) {
                        if (!true) {
                            v1 = ("hello" + v1)
                            var v49 uint32 = 42
                            _ = v49
                        } else {
                            v50 := GenS2{F0: -2147483648, F1: 649, F2: 2}
                            _ = v50
                        }
                        var v51 int32 = 42
                        _ = v51
                    } else {
                        v1 = "foo bar"
                    }
                    _ = fi48
                }
            default:
                if false {
                    if (!(!(-387 >= -490))) {
                        var v52 float64 = -2.25
                        v1 = v1
                        for fi53 := 0; fi53 < 4; fi53++ {
                            switch (fi53 * 32768) {
                            case 128:
                                v1 = (v1 + "")
                            case 32768:
                                if ((822 <= 2147483647) && true) {
                                    var v54 int16 = 742
                                    _ = v54
                                } else {
                                    v3 = v3
                                    fmt.Println(v1, fi46)
                                }
                            case 289:
                                if (0 > 331) {
                                    var v55 int16 = 51
                                    _ = v55
                                }
                            }
                            fmt.Println(fi46)
                            fmt.Println(fi46)
                            _ = fi53
                        }
                        _ = v52
                    }
                    var v56 float64 = 1e10
                    _ = v56
                } else {
                    var v57 float32 = 0.0
                    _ = v57
                }
            }
        case "":
            _ = len(v2)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            if false {
                var v58 int64 = 0
                for fi59 := 0; fi59 < 2; fi59++ {
                    fmt.Println(fi59)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi59
                }
                _ = v58
            } else {
                if true {
                    if (true || ("\n\t" == "\n\t")) {
                        v3 = v3
                        fmt.Println(v1, v1, v3)
                    } else {
                        if ((2147483647 < 65535) && (42 >= 255)) {
                            v1 = ("" + v1)
                            v1 = v1
                            var v60 string = "\n\t"
                            _ = v60
                        }
                        if true {
                            for fi61 := 0; fi61 < 20; fi61++ {
                                v2 = append(v2, ((v1 + "a") + (v1 + "a")))
                                v1 = ""
                                switch 2147483647 {
                                case 42:
                                    var v62 uint = 42
                                    _ = v62
                                case 9223372036854775807:
                                    fmt.Println(v1, fi46)
                                default:
                                    if false {
                                        var v63 float32 = float32((1.5 / 0.1))
                                        fmt.Println(v1, fi46)
                                        _ = v63
                                    }
                                    switch "foo bar" {
                                    case "":
                                        v1 = v1
                                    default:
                                        v1 = v1
                                    }
                                    switch -604 {
                                    case 127:
                                        var v64 bool = (!true)
                                        _ = v64
                                    case -2147483648:
                                        switch (("a" + (v1 + ("a" + v1))) + ((v1 + v1) + ((v1 + (v1 + ("foo bar" + v1))) + "\n\t"))) {
                                        case "hello":
                                            var v65 int16 = 32767
                                            _ = v65
                                        case "foo bar":
                                            switch "foo bar" {
                                            case "a":
                                                _ = len(v2)
                                            case "foo bar":
                                                switch (fi46 << ((342) & 63)) {
                                                case 444:
                                                    v1 = "hello"
                                                default:
                                                    v2 = append(v2, ("foo bar" + "a"))
                                                }
                                            default:
                                                if true {
                                                    if (0 > 809) {
                                                        if false {
                                                            var v66 int8 = int8(fi61)
                                                            v3 = (^(v3))
                                                            v67 := GenS4{F0: -1}
                                                            v1 = (((v1 + v1) + v1) + "")
                                                            var v68 int8 = 0
                                                            var v69 float64 = 0.0
                                                            v70 := GenS2{F0: -2147483648, F1: 42, F2: 65535}
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                            _ = v70
                                                        } else {
                                                            var v71 uint16 = 32767
                                                            var v72 float64 = ((float64(float32((0.0 + (-2.25 / (3.14159 - float64(0.1)))))) * float64((0.0 * (((3.14159 + 1e10) + 3.14159) * -2.25)))) / ((((1.5 + 0.0) + 1.5) / 0.1) / -2.25))
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        v3 = (v3 & (v3 & -95))
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    v73 := GenS5{F0: 267, F1: -9223372036854775808}
                                                    switch (v1 + v1) {
                                                    case "hello":
                                                        if (((0 <= 802) && (-9223372036854775808 > 0)) || true) {
                                                            _ = len(v2)
                                                            v2 = append(v2, v1)
                                                            v1 = v1
                                                            v3 -= v3
                                                        } else {
                                                            var v74 uint32 = 65535
                                                            _ = v74
                                                        }
                                                        if false {
                                                            var v75 uint64 = 18446744073709551615
                                                            _ = v75
                                                        } else {
                                                            v76 := GenS3{F0: v1}
                                                            _ = v76
                                                        }
                                                    }
                                                    _ = v73
                                                } else {
                                                    v77 := GenS4{F0: 9223372036854775807}
                                                    v2 = append(v2, "hello")
                                                    v3 = 96
                                                    _ = v77
                                                }
                                            }
                                        case "a":
                                            var v78 uint = 1
                                            _ = v78
                                        }
                                    default:
                                        switch fi46 {
                                        case -9223372036854775808:
                                            var v79 uint64 = 880
                                            _ = v79
                                        case 556:
                                            v80 := GenS3{F0: v1}
                                            _ = v80
                                        }
                                    }
                                }
                                _ = fi61
                            }
                        }
                        if (198 != 2) {
                            v1 = v1
                            for fi81 := 0; fi81 < 12; fi81++ {
                                if false {
                                    v3 -= v3
                                    v1 = "foo bar"
                                    fmt.Println(fi81, v1, v1)
                                }
                                var v82 int32 = 1
                                _ = fi81
                                _ = v82
                            }
                        } else {
                            if false {
                                var v83 bool = (!true)
                                _ = v83
                            } else {
                                v84 := GenS1{F0: 215, F1: 42}
                                var v85 int8 = v3
                                _ = len(v2)
                                _ = v84
                                _ = v85
                            }
                            switch v1 {
                            case "foo bar":
                                for fi86 := 0; fi86 < 14; fi86++ {
                                    var v87 uint16 = 65535
                                    v1 = v1
                                    _ = fi86
                                    _ = v87
                                }
                            case "hello":
                                if (65535 >= 1) {
                                    var v88 int64 = -1
                                    _ = v88
                                } else {
                                    switch v1 {
                                    case "\n\t":
                                        switch genFn0() {
                                        case 32767:
                                            switch (fi46 << ((31) & 63)) {
                                            case 32767:
                                                if len(v2) > 0 {
                                                    v2[0] = v1
                                                }
                                            case 2147483647:
                                                v1 = "a"
                                            case 4294967295:
                                                fmt.Println(fi46, v1)
                                            }
                                        case -2147483648:
                                            switch "" {
                                            case "hello":
                                                var v89 int8 = v3
                                                _ = v89
                                            case "foo bar":
                                                if true {
                                                    for fi90 := 0; fi90 < 15; fi90++ {
                                                        if (339 > 42) {
                                                            v1 = v1
                                                        } else {
                                                            _ = len(v2)
                                                            switch (^(fi46)) {
                                                            case 32767:
                                                                v3 -= v3
                                                            case 794:
                                                                switch (v1 + ((("foo bar" + v1) + v1) + v1)) {
                                                                case "hello":
                                                                    switch (-(fi90)) {
                                                                    case -1:
                                                                        var v91 int8 = 127
                                                                        _ = v91
                                                                    }
                                                                case "":
                                                                    v92 := GenS6{F0: 652}
                                                                    _ = v92
                                                                default:
                                                                    if (127 == 0) {
                                                                        var v93 float32 = -2.25
                                                                        _ = v93
                                                                    }
                                                                    var v94 float64 = ((float64(-2.25) / float64(-2.25)) * float64(-2.25))
                                                                    if (!(239 > 2)) {
                                                                        v3 = (v3 / (((v3 << ((2) & 63))) | 1))
                                                                    }
                                                                    _ = v94
                                                                }
                                                                _ = len(v2)
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = v1
                                                            }
                                                        }
                                                        _ = fi90
                                                    }
                                                    v3 = (v3 ^ v3)
                                                    if false {
                                                        for fi95 := 0; fi95 < 14; fi95++ {
                                                            var v96 uint = 58
                                                            var v97 int16 = 1000
                                                            _ = fi95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        if (-32768 < 1000) {
                                                            v98 := GenS6{F0: 1}
                                                            var v99 int32 = 0
                                                            _ = v98
                                                            _ = v99
                                                        } else {
                                                            fmt.Println(v1, v1)
                                                            var v100 string = v1
                                                            var v101 bool = (true || false)
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                        if true {
                                                            for fi102 := 0; fi102 < 3; fi102++ {
                                                                v1 = ((v1 + ((v1 + v1) + v1)) + ("" + v1))
                                                                _ = fi102
                                                            }
                                                            switch (fi46 % ((fi46) | 1)) {
                                                            case 2147483647:
                                                                for fi103 := 0; fi103 < 7; fi103++ {
                                                                    fmt.Println(fi46, v1, v3)
                                                                    switch fi46 {
                                                                    case 452:
                                                                        v2 = append(v2, v1)
                                                                    case 2147483647:
                                                                        v1 = v1
                                                                    }
                                                                    var v104 int32 = 42
                                                                    _ = fi103
                                                                    _ = v104
                                                                }
                                                            }
                                                        } else {
                                                            var v105 int8 = v3
                                                            _ = v105
                                                        }
                                                    } else {
                                                        v3 += -1
                                                        if false {
                                                            var v106 uint8 = 2
                                                            var v107 float64 = ((((0.0 / float64(0.1)) - -2.25) - 0.0) * 0.0)
                                                            if true {
                                                                if false {
                                                                    if false {
                                                                        for fi108 := 0; fi108 < 19; fi108++ {
                                                                            fmt.Println(v106)
                                                                            if false {
                                                                                var v109 uint32 = 795
                                                                                fmt.Println(v109, v106)
                                                                                _ = v109
                                                                            }
                                                                            switch genFn0() {
                                                                            case 4294967295:
                                                                                v110 := GenS5{F0: 32767, F1: 9223372036854775807}
                                                                                _ = v110
                                                                            case 9223372036854775807:
                                                                                v107 = (float64(0.1) + 0.0)
                                                                            }
                                                                            _ = fi108
                                                                        }
                                                                    } else {
                                                                        if ((("foo bar" > (v1 + ((v1 + v1) + "\n\t"))) || true) || false) {
                                                                            v106 = (^(v106))
                                                                        }
                                                                        var v111 string = v1
                                                                        switch (fi46 << ((65535) & 63)) {
                                                                        case -2147483648:
                                                                            fmt.Println(fi46)
                                                                        case 128:
                                                                            v112 := GenS4{F0: -586}
                                                                            _ = v112
                                                                        case 1:
                                                                            if false {
                                                                                v113 := GenS1{F0: 22, F1: 1}
                                                                                var v114 int64 = -171
                                                                                _ = len(v2)
                                                                                _ = v113
                                                                                _ = v114
                                                                            } else {
                                                                                var v115 uint16 = 540
                                                                                for fi116 := 0; fi116 < 5; fi116++ {
                                                                                    v1 = v1
                                                                                    if (false && true) {
                                                                                        var v117 uint32 = 4294967295
                                                                                        v118 := GenS4{F0: 1}
                                                                                        _ = v117
                                                                                        _ = v118
                                                                                    }
                                                                                    v119 := GenS1{F0: (v106 | (v106 - uint8(v115))), F1: 819}
                                                                                    v3 -= v3
                                                                                    var v120 string = (v1 + "foo bar")
                                                                                    _ = fi116
                                                                                    _ = v119
                                                                                    _ = v120
                                                                                }
                                                                                fmt.Println(v106, v106, v1)
                                                                                _ = v115
                                                                            }
                                                                        default:
                                                                            var v121 bool = true
                                                                            _ = v121
                                                                        }
                                                                        _ = v111
                                                                    }
                                                                    for fi122 := 0; fi122 < 3; fi122++ {
                                                                        fmt.Println(v3, fi122, fi122)
                                                                        v123 := GenS7{F0: v1, F1: 255}
                                                                        switch (fi46 >> ((359) & 63)) {
                                                                        case 9223372036854775807:
                                                                            if (v107 < (v107 * v107)) {
                                                                                v3 = (v3 / (((v3 & int8(v106))) | 1))
                                                                                v107 -= (((0.1 * v107) - v107) * 1e10)
                                                                            } else {
                                                                                var v124 int16 = 32767
                                                                                _ = v124
                                                                            }
                                                                        case 127:
                                                                            var v125 int64 = -950
                                                                            _ = v125
                                                                        }
                                                                        if (v1 < "\n\t") {
                                                                            v107 = v107
                                                                            v1 = "\n\t"
                                                                            v2 = append(v2, v1)
                                                                        } else {
                                                                            v106 -= 0
                                                                        }
                                                                        _ = fi122
                                                                        _ = v123
                                                                    }
                                                                }
                                                            } else {
                                                                v106 = 128
                                                            }
                                                            _ = v106
                                                            _ = v107
                                                        } else {
                                                            v3 = (v3 << ((0) & 63))
                                                        }
                                                    }
                                                }
                                            default:
                                                for fi126 := 0; fi126 < 8; fi126++ {
                                                    if true {
                                                        if (65535 <= 473) {
                                                            switch (fi126 | 42) {
                                                            case 9223372036854775807:
                                                                v127 := GenS2{F0: 968, F1: 823, F2: 32767}
                                                                _ = v127
                                                            case -2147483648:
                                                                v1 = v1
                                                            case 42:
                                                                _ = len(v2)
                                                            default:
                                                                v128 := GenS3{F0: v1}
                                                                _ = v128
                                                            }
                                                            v2 = append(v2, (v1 + v1))
                                                            v3 = (-(v3))
                                                        }
                                                        var v129 uint64 = 938
                                                        v130 := GenS8{F0: 1000, F1: 2}
                                                        _ = len(v2)
                                                        _ = v129
                                                        _ = v130
                                                    } else {
                                                        if (!false) {
                                                            v3 += v3
                                                            var v131 uint64 = 188
                                                            if false {
                                                                v3 = 0
                                                                v132 := GenS2{F0: -2147483648, F1: 42, F2: 255}
                                                                if false {
                                                                    switch fi46 {
                                                                    case 4294967295:
                                                                        fmt.Println(v131)
                                                                    case -2147483648:
                                                                        v131 -= v131
                                                                    default:
                                                                        v133 := GenS2{F0: int32(v131), F1: 0, F2: 32767}
                                                                        _ = v133
                                                                    }
                                                                    fmt.Println(v3, v1, v131)
                                                                    v1 = "foo bar"
                                                                }
                                                                _ = v132
                                                            } else {
                                                                var v134 uint16 = 255
                                                                _ = v134
                                                            }
                                                            _ = v131
                                                        } else {
                                                            fmt.Println(v3, v1)
                                                        }
                                                        v1 = v1
                                                    }
                                                    _ = fi126
                                                }
                                            }
                                        case -52:
                                            if ((v1 <= ((((v1 + v1) + v1) + v1) + v1)) || false) {
                                                v1 = v1
                                                v3 = (v3 / (((-(v3))) | 1))
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                for fi135 := 0; fi135 < 4; fi135++ {
                                                    var v136 bool = (!(-2.25 <= 0.0))
                                                    fmt.Println(fi46, fi46)
                                                    switch "" {
                                                    case "":
                                                        var v137 string = ""
                                                        _ = v137
                                                    }
                                                    _ = fi135
                                                    _ = v136
                                                }
                                                v3 += (v3 % (((^(v3))) | 1))
                                            }
                                        }
                                    default:
                                        var v138 uint32 = 294
                                        _ = v138
                                    }
                                    var v139 uint32 = 0
                                    if (true || (genFn0() < (fi46 << ((666) & 63)))) {
                                        var v140 int = genFn0()
                                        if (!true) {
                                            fmt.Println(fi46, v140)
                                            if (1.5 >= 1e10) {
                                                v2 = append(v2, "foo bar")
                                                v1 = ((("" + (v1 + "")) + v1) + "foo bar")
                                            }
                                        } else {
                                            _ = len(v2)
                                            for fi141 := 0; fi141 < 4; fi141++ {
                                                v142 := GenS7{F0: ("\n\t" + (v1 + ((("" + v1) + "\n\t") + ((v1 + "") + "foo bar")))), F1: 42}
                                                _ = fi141
                                                _ = v142
                                            }
                                            var v143 uint32 = (v139 % (((v139 & 65535)) | 1))
                                            _ = v143
                                        }
                                        v3 -= v3
                                        _ = v140
                                    }
                                    _ = v139
                                }
                            default:
                                if (32767 <= 327) {
                                    if (204 <= 876) {
                                        if true {
                                            fmt.Println(fi46, v1)
                                            _ = len(v2)
                                            for fi144 := 0; fi144 < 4; fi144++ {
                                                v3 += 66
                                                var v145 float32 = (1e10 * 0.1)
                                                var v146 bool = true
                                                _ = fi144
                                                _ = v145
                                                _ = v146
                                            }
                                        }
                                        if false {
                                            v3 += (v3 ^ (v3 << ((255) & 63)))
                                        }
                                        if (128 > 247) {
                                            for fi147 := 0; fi147 < 16; fi147++ {
                                                var v148 uint8 = 2
                                                if (v148 == (^(v148))) {
                                                    _ = len(v2)
                                                    if ((^(v3)) != v3) {
                                                        v148 += v148
                                                        var v149 int64 = 1
                                                        switch fi147 {
                                                        case -108:
                                                            var v150 uint8 = 2
                                                            _ = v150
                                                        case -12:
                                                            switch "a" {
                                                            case "a":
                                                                var v151 int = (fi46 | (fi46 * (fi46 & genFn0())))
                                                                _ = v151
                                                            }
                                                        default:
                                                            if true {
                                                                switch (fi147 / (((fi147 / (((fi147 | 32767)) | 1))) | 1)) {
                                                                case 9223372036854775807:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case 2147483647:
                                                                    v3 += v3
                                                                }
                                                                if true {
                                                                    v2 = append(v2, (((v1 + ((v1 + "foo bar") + (((("\n\t" + "") + v1) + (("\n\t" + "\n\t") + (v1 + "a"))) + v1))) + v1) + ("" + "hello")))
                                                                }
                                                                var v152 uint16 = 42
                                                                _ = v152
                                                            }
                                                        }
                                                        _ = v149
                                                    }
                                                    if len(v2) > 0 {
                                                        v2[0] = v1
                                                    }
                                                } else {
                                                    switch fi46 {
                                                    case 415:
                                                        switch v1 {
                                                        case "\n\t":
                                                            if (((fi46 % ((fi46) | 1)) <= int(v3)) && false) {
                                                                var v153 uint8 = v148
                                                                fmt.Println(fi147)
                                                                var v154 int8 = -18
                                                                if true {
                                                                    var v155 int64 = 751
                                                                    var v156 uint32 = 2147483647
                                                                    _ = v155
                                                                    _ = v156
                                                                } else {
                                                                    v1 = ("a" + "foo bar")
                                                                    if false {
                                                                        var v157 float64 = -2.25
                                                                        v157 = ((float64(-2.25) + (3.14159 - (v157 * v157))) / 0.1)
                                                                        v2 = append(v2, ("foo bar" + v1))
                                                                        _ = v157
                                                                    }
                                                                }
                                                                var v158 string = (v1 + ("foo bar" + v1))
                                                                _ = v153
                                                                _ = v154
                                                                _ = v158
                                                            }
                                                        case "hello":
                                                            v1 = ("foo bar" + v1)
                                                        case "a":
                                                            var v159 float32 = 0.0
                                                            _ = v159
                                                        }
                                                    case 42:
                                                        v160 := GenS9{F0: 2147483648, F1: true, F2: 354}
                                                        _ = v160
                                                    default:
                                                        v1 = v1
                                                    }
                                                }
                                                _ = len(v2)
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if ((fi46 * fi147) < (fi46 & (fi147 << ((42) & 63)))) {
                                                        v3 -= (v3 * v3)
                                                        v1 = "foo bar"
                                                        var v161 int64 = 9223372036854775807
                                                        _ = v161
                                                    } else {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v3 += 29
                                                        v3 += v3
                                                    }
                                                    var v162 uint32 = 1
                                                    _ = v162
                                                } else {
                                                    var v163 bool = (1.5 <= 0.0)
                                                    if (genFn0() == fi147) {
                                                        v3 = int8(v148)
                                                        v163 = (2147483647 > -2147483648)
                                                        fmt.Println(v148, v3, v163)
                                                    }
                                                    fmt.Println(v1, v3, v1)
                                                    _ = v163
                                                }
                                                var v164 int16 = 42
                                                if true {
                                                    fmt.Println(v164, v1, v1)
                                                    v165 := GenS10{F0: fi46}
                                                    v1 = v1
                                                    switch (fi147 & genFn0()) {
                                                    case -244:
                                                        switch (fi46 >> ((43) & 63)) {
                                                        case 42:
                                                            var v166 int = (fi147 << ((65535) & 63))
                                                            _ = v166
                                                        default:
                                                            v3 -= (^(v3))
                                                        }
                                                    default:
                                                        fmt.Println(v3, v148)
                                                    }
                                                    _ = v165
                                                }
                                                _ = fi147
                                                _ = v148
                                                _ = v164
                                            }
                                            if true {
                                                if (!true) {
                                                    for fi167 := 0; fi167 < 15; fi167++ {
                                                        if true {
                                                            var v168 int8 = (v3 << ((4294967295) & 63))
                                                            if true {
                                                                var v169 int32 = 0
                                                                var v170 int = (fi167 + genFn0())
                                                                v170 = genFn0()
                                                                v170 += int(v3)
                                                                _ = v169
                                                                _ = v170
                                                            }
                                                            _ = v168
                                                        } else {
                                                            var v171 int32 = 218
                                                            v1 = ("foo bar" + v1)
                                                            _ = v171
                                                        }
                                                        v3 = (v3 + (v3 | v3))
                                                        var v172 uint8 = 2
                                                        _ = fi167
                                                        _ = v172
                                                    }
                                                    switch v1 {
                                                    case "\n\t":
                                                        var v173 uint16 = 32767
                                                        _ = v173
                                                    case "hello":
                                                        if true {
                                                            v174 := GenS11{F0: 42, F1: 42}
                                                            _ = v174
                                                        }
                                                    default:
                                                        v175 := GenS5{F0: 0, F1: -907}
                                                        _ = v175
                                                    }
                                                    if (3.14159 < 0.0) {
                                                        fmt.Println(v3, v1)
                                                        for fi176 := 0; fi176 < 19; fi176++ {
                                                            if (-1 < 1) {
                                                                v177 := GenS10{F0: fi46}
                                                                v178 := GenS8{F0: 32767, F1: 4294967295}
                                                                if false {
                                                                    switch v1 {
                                                                    case "hello":
                                                                        fmt.Println(fi46, fi176, fi176)
                                                                    }
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v1
                                                                    }
                                                                    var v179 uint = 42
                                                                    _ = v179
                                                                } else {
                                                                    fmt.Println(fi46, fi176)
                                                                    v3 = (v3 % ((v3) | 1))
                                                                }
                                                                _ = v177
                                                                _ = v178
                                                            } else {
                                                                v3 += v3
                                                            }
                                                            var v180 uint16 = 32767
                                                            if ((fi176 >= fi46) && (!true)) {
                                                                var v181 uint32 = 65535
                                                                if true {
                                                                    v182 := GenS9{F0: (v181 + (v181 ^ (^(v181)))), F1: false, F2: int16(v180)}
                                                                    fmt.Println(v181)
                                                                    _ = v182
                                                                }
                                                                _ = v181
                                                            }
                                                            if len(v2) > 0 {
                                                                v2[0] = v1
                                                            }
                                                            v2 = append(v2, ((v1 + "foo bar") + v1))
                                                            _ = fi176
                                                            _ = v180
                                                        }
                                                    } else {
                                                        var v183 uint16 = 122
                                                        if (42 <= 142) {
                                                            var v184 uint16 = (v183 & v183)
                                                            if (fi46 == fi46) {
                                                                for fi185 := 0; fi185 < 19; fi185++ {
                                                                    if false {
                                                                        v186 := GenS12{F0: false, F1: (true && true)}
                                                                        fmt.Println(fi185, v1, v184)
                                                                        _ = v186
                                                                    } else {
                                                                        v1 = v1
                                                                        if (v183 > 2) {
                                                                            fmt.Println(fi46, fi185, v3)
                                                                            var v187 int16 = 32767
                                                                            _ = v187
                                                                        } else {
                                                                            if true {
                                                                                v3 = v3
                                                                                v188 := GenS11{F0: -1, F1: -604}
                                                                                var v189 uint64 = 42
                                                                                v2 = append(v2, "foo bar")
                                                                                _ = v188
                                                                                _ = v189
                                                                            }
                                                                            v1 = (((v1 + "\n\t") + "hello") + "foo bar")
                                                                        }
                                                                        v183 = (v183 | (v184 % ((v184) | 1)))
                                                                        v190 := GenS5{F0: (v183 >> ((1) & 63)), F1: 42}
                                                                        if false {
                                                                            _ = len(v2)
                                                                            v191 := GenS4{F0: 589}
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            v192 := GenS5{F0: v183, F1: 970}
                                                                            _ = v191
                                                                            _ = v192
                                                                        }
                                                                        _ = v190
                                                                    }
                                                                    _ = fi185
                                                                }
                                                                v184 = v184
                                                            } else {
                                                                v3 += v3
                                                            }
                                                            v193 := GenS13{F0: (true && (!(false || ((false && false) && (((false && ((true || (true || (-2147483648 > 1))) && true)) || (312 > 1000)) || false))))), F1: 645}
                                                            _ = v184
                                                            _ = v193
                                                        }
                                                        switch genFn0() {
                                                        case -1:
                                                            for fi194 := 0; fi194 < 9; fi194++ {
                                                                fmt.Println(v3, v3, v3)
                                                                v1 = "foo bar"
                                                                if (1 > 2) {
                                                                    v195 := GenS9{F0: 2147483647, F1: false, F2: -1}
                                                                    switch (v1 + "\n\t") {
                                                                    case "foo bar":
                                                                        var v196 int16 = 32767
                                                                        _ = v196
                                                                    case "":
                                                                        v197 := GenS13{F0: false, F1: 118}
                                                                        _ = v197
                                                                    default:
                                                                        fmt.Println(v1, v3, v3)
                                                                    }
                                                                    _ = v195
                                                                }
                                                                v183 -= v183
                                                                _ = fi194
                                                            }
                                                        case -228:
                                                            fmt.Println(fi46, v183)
                                                        case 4294967295:
                                                            var v198 float32 = -2.25
                                                            _ = v198
                                                        }
                                                        _ = v183
                                                    }
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                v1 = v1
                                                v199 := GenS2{F0: 42, F1: 32767, F2: 287}
                                                switch fi46 {
                                                case 127:
                                                    if (-1 > v3) {
                                                        fmt.Println(v1)
                                                        for fi200 := 0; fi200 < 10; fi200++ {
                                                            v3 += 100
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            if false {
                                                                v3 = -98
                                                                switch v1 {
                                                                case "foo bar":
                                                                    if (true && (false || false)) {
                                                                        fmt.Println(v1)
                                                                        switch v1 {
                                                                        case "foo bar":
                                                                            var v201 string = ("\n\t" + "hello")
                                                                            _ = v201
                                                                        case "hello":
                                                                            fmt.Println(v3, v1)
                                                                        case "":
                                                                            v1 = v1
                                                                        }
                                                                        v2 = append(v2, "a")
                                                                    } else {
                                                                        if ((false && false) && false) {
                                                                            fmt.Println(fi46)
                                                                            var v202 int64 = -9223372036854775808
                                                                            v203 := GenS14{F0: false, F1: (^(v3)), F2: 928}
                                                                            _ = v202
                                                                            _ = v203
                                                                        }
                                                                        v3 -= v3
                                                                        var v204 uint64 = 4294967295
                                                                        _ = v204
                                                                    }
                                                                case "":
                                                                    var v205 uint8 = 127
                                                                    _ = v205
                                                                default:
                                                                    v3 -= (v3 % ((v3) | 1))
                                                                }
                                                                var v206 float64 = (0.1 + 3.14159)
                                                                _ = v206
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        if false {
                                                                            if (("" + (v1 + v1)) <= v1) {
                                                                                if (380 == 955) {
                                                                                    v3 = v3
                                                                                }
                                                                                if ((genFn0() == genFn0()) || (1 < 2)) {
                                                                                    v3 += (v3 << ((0) & 63))
                                                                                }
                                                                            } else {
                                                                                v3 += 127
                                                                                v207 := GenS9{F0: 2147483648, F1: (128 != 127), F2: 42}
                                                                                var v208 string = (v1 + (v1 + (v1 + (((v1 + ((("\n\t" + v1) + ((v1 + v1) + v1)) + v1)) + v1) + v1))))
                                                                                _ = v207
                                                                                _ = v208
                                                                            }
                                                                            v209 := GenS15{F0: -989, F1: 65535}
                                                                            if false {
                                                                                if (127 <= 0) {
                                                                                    var v210 int8 = (v3 + (v3 + v3))
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    v1 = (("foo bar" + v1) + v1)
                                                                                    v1 = (v1 + v1)
                                                                                    _ = v210
                                                                                } else {
                                                                                    v211 := GenS16{F0: 255, F1: 2}
                                                                                    v1 = ((v1 + "\n\t") + v1)
                                                                                    v212 := GenS17{F0: -849, F1: 4294967295, F2: -1}
                                                                                    _ = v211
                                                                                    _ = v212
                                                                                }
                                                                                _ = len(v2)
                                                                                var v213 uint32 = 0
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v213
                                                                            }
                                                                            _ = v209
                                                                        } else {
                                                                            var v214 string = "foo bar"
                                                                            _ = v214
                                                                        }
                                                                        v215 := GenS12{F0: (!true), F1: true}
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            fmt.Println(v3)
                                                                        }
                                                                        _ = v215
                                                                    }
                                                                } else {
                                                                    v216 := GenS18{F0: -560, F1: 2147483647, F2: 1}
                                                                    v217 := GenS12{F0: (!true), F1: (1 < 0)}
                                                                    _ = v216
                                                                    _ = v217
                                                                }
                                                                _ = len(v2)
                                                            }
                                                            _ = fi200
                                                        }
                                                        var v218 int64 = 1
                                                        _ = v218
                                                    }
                                                case -571:
                                                    if len(v2) > 0 {
                                                        v2[0] = v1
                                                    }
                                                default:
                                                    switch genFn0() {
                                                    case 1:
                                                        if (0 > 65535) {
                                                            if (!true) {
                                                                var v219 uint32 = 2147483647
                                                                _ = v219
                                                            } else {
                                                                v2 = append(v2, v1)
                                                                v1 = ((v1 + (v1 + "\n\t")) + ((("a" + (v1 + v1)) + "\n\t") + "a"))
                                                            }
                                                        }
                                                    case 2147483647:
                                                        for fi220 := 0; fi220 < 2; fi220++ {
                                                            var v221 float32 = 1.5
                                                            v3 += v3
                                                            _ = fi220
                                                            _ = v221
                                                        }
                                                    }
                                                }
                                                _ = v199
                                            }
                                            if (127 > 216) {
                                                v222 := GenS13{F0: (false || false), F1: 773}
                                                var v223 uint = 4294967295
                                                _ = v222
                                                _ = v223
                                            } else {
                                                fmt.Println(fi46, v1, v3)
                                            }
                                        } else {
                                            var v224 bool = (false || (false && (1 == (fi46 - -162))))
                                            var v225 int16 = -32768
                                            _ = v224
                                            _ = v225
                                        }
                                    }
                                    v1 = v1
                                    var v226 int64 = -644
                                    _ = v226
                                } else {
                                    var v227 float64 = (-2.25 / (-2.25 + 1.5))
                                    _ = v227
                                }
                            }
                            fmt.Println(fi46, v1, v1)
                        }
                    }
                } else {
                    if (!((!true) && false)) {
                        v228 := GenS11{F0: -1, F1: 579}
                        _ = v228
                    }
                    if len(v2) > 0 {
                        v2[0] = v1
                    }
                }
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
                for fi229 := 0; fi229 < 8; fi229++ {
                    if (false && (0 != -421)) {
                        var v230 float64 = (1.5 + (((1.5 + (0.0 + (((3.14159 + (-2.25 / 0.1)) - 0.1) / 3.14159))) + 1.5) / -2.25))
                        v1 = v1
                        _ = v230
                    }
                    if len(v2) > 0 {
                        v2[0] = "hello"
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi229
                }
            }
            var v231 int64 = -322
            _ = v231
        }
        _ = fi46
    }
    if false {
        v3 = -19
        var v232 uint = 255
        _ = v232
    }
    switch "a" {
    case "foo bar":
        if ((38 < 989) && false) {
            v2 = append(v2, ((((((v1 + ((v1 + "") + v1)) + (v1 + (v1 + v1))) + v1) + "hello") + "foo bar") + v1))
            v1 = v1
        } else {
            v1 = v1
            switch ((v1 + (v1 + "a")) + "foo bar") {
            case "a":
                for fi233 := 0; fi233 < 4; fi233++ {
                    var v234 int32 = 42
                    if (2 < 1) {
                        v1 = "foo bar"
                        var v235 int = (fi233 & (fi233 | fi233))
                        v235 = v235
                        _ = v235
                    } else {
                        for fi236 := 0; fi236 < 7; fi236++ {
                            var v237 uint16 = 255
                            _ = fi236
                            _ = v237
                        }
                        switch 128 {
                        case 2147483647:
                            var v238 int64 = -9223372036854775808
                            _ = v238
                        case 127:
                            for fi239 := 0; fi239 < 6; fi239++ {
                                var v240 uint16 = 222
                                var v241 float64 = -2.25
                                _ = fi239
                                _ = v240
                                _ = v241
                            }
                        }
                    }
                    v2 = append(v2, (v1 + v1))
                    _ = fi233
                    _ = v234
                }
            }
            v2 = append(v2, v1)
        }
    case "hello":
        v3 -= 0
    }
    v242 := GenS19{F0: 0, F1: 1, F2: genFn0()}
    if ((0 == -1) && (v1 != v1)) {
        if true {
            v1 = (("\n\t" + ("a" + v1)) + ((v1 + ("foo bar" + v1)) + v1))
            v243 := GenS9{F0: 153, F1: (true || false), F2: 0}
            switch 600 {
            case -2147483648:
                fmt.Println(v1)
            case -932:
                var v244 uint64 = 18446744073709551615
                _ = v244
            default:
                switch 32767 {
                case -601:
                    if ((!((v3 ^ (v3 & v3)) <= v3)) || true) {
                        var v245 bool = true
                        if v245 {
                            v3 = v3
                        }
                        _ = v245
                    }
                default:
                    v1 = v1
                }
            }
            _ = v243
        } else {
            if ((!(true && (false || (2 <= 101)))) || (true && true)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    }
    switch v1 {
    case "a":
        var v246 int16 = 32767
        _ = v246
    default:
        v2 = append(v2, v1)
    }
    v247 := GenS2{F0: 0, F1: 0, F2: 32767}
    v1 = "hello"
    v248 := GenS7{F0: "a", F1: 675}
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v242
    _ = v247
    _ = v248
    _ = v3
}
