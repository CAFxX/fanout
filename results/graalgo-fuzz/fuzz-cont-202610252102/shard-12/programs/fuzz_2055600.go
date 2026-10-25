package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 uint32; F2 string }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 uint16; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint16; F1 uint32 }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 bool; F1 uint16; F2 bool }
type GenS17 struct { F0 int64; F1 uint8; F2 int32 }
func genFn0(p0 uint) uint {
    p0 = (p0 ^ 1)
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: 255, F1: (("hello" + "") + "\n\t")}
            v7 := GenS2{F0: -32768}
            p0 = p0
            p0 = p0
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 string = "foo bar"
        switch "a" {
        case "hello":
            var v9 int16 = 32767
            _ = v9
        }
        _ = fi4
        _ = v8
    }
    if (!(9223372036854775807 == -1)) {
        fmt.Println(p0)
        fmt.Println(p0)
        if true {
            for fi10 := 0; fi10 < 19; fi10++ {
                for fi11 := 0; fi11 < 18; fi11++ {
                    if false {
                        if true {
                            var v12 uint64 = 2
                            _ = v12
                        }
                        fmt.Println(fi10, fi11, fi11)
                        p0 = p0
                        v13 := []int{(fi10 + (-(fi10))), fi11, 32767}
                        var v14 uint = (p0 >> ((65535) & 63))
                        _ = v13
                        _ = v14
                    } else {
                        var v15 uint16 = 42
                        v16 := []int{fi11, 42}
                        _ = v15
                        _ = v16
                    }
                    var v17 string = "\n\t"
                    p0 = p0
                    _ = fi11
                    _ = v17
                }
                v18 := []int{}
                if (false || (((float32((0.1 - 3.14159)) * 1e10) != 1e10) || false)) {
                    var v19 bool = false
                    if (!(!true)) {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                    } else {
                        if (127 <= -97) {
                            if ((((255 <= 1) || (-301 != -531)) && (!(fi10 > 59))) || ((^(p0)) < 512)) {
                                var v20 uint8 = 42
                                if (3.14159 >= 0.0) {
                                    var v21 uint64 = 42
                                    if v19 {
                                        for fi22 := 0; fi22 < 13; fi22++ {
                                            switch ("hello" + "\n\t") {
                                            case "hello":
                                                switch fi10 {
                                                case 4294967295:
                                                    switch (fi22 * 645) {
                                                    case -820:
                                                        v21 = (v21 % (((v21 + 18446744073709551615)) | 1))
                                                    case -990:
                                                        v19 = v19
                                                    }
                                                case -9223372036854775808:
                                                    var v23 uint = 42
                                                    _ = v23
                                                default:
                                                    switch "foo bar" {
                                                    case "":
                                                        var v24 bool = true
                                                        _ = v24
                                                    case "hello":
                                                        v19 = false
                                                    }
                                                }
                                            case "":
                                                var v25 int64 = -919
                                                _ = v25
                                            }
                                            _ = fi22
                                        }
                                        v18 = append(v18, (fi10 << ((p0) & 63)))
                                        var v26 uint8 = v20
                                        _ = v26
                                    }
                                    _ = v21
                                }
                                _ = v20
                            } else {
                                if true {
                                    v27 := GenS1{F0: 0, F1: "foo bar"}
                                    v28 := GenS2{F0: -224}
                                    var v29 int64 = -486
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                } else {
                                    if false {
                                        fmt.Println(v19, p0, fi10)
                                        p0 = 4294967295
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        for fi30 := 0; fi30 < 6; fi30++ {
                                            v19 = v19
                                            v31 := GenS1{F0: 238, F1: "foo bar"}
                                            switch (fi10 - fi10) {
                                            case -880:
                                                if len(v18) > 0 {
                                                    v18[0] = fi10
                                                }
                                            case 9223372036854775807:
                                                v32 := GenS3{F0: 1}
                                                _ = v32
                                            default:
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                            }
                                            v19 = (!true)
                                            _ = fi30
                                            _ = v31
                                        }
                                    }
                                    fmt.Println(v19, v19, p0)
                                }
                                for fi33 := 0; fi33 < 20; fi33++ {
                                    v19 = (!v19)
                                    _ = len(v18)
                                    var v34 int16 = 0
                                    _ = fi33
                                    _ = v34
                                }
                            }
                            if false {
                                v18 = append(v18, (fi10 / ((96) | 1)))
                                var v35 uint8 = 128
                                if true {
                                    for fi36 := 0; fi36 < 10; fi36++ {
                                        switch "a" {
                                        case "":
                                            if (4294967295 <= 2) {
                                                if false {
                                                    var v37 int8 = 101
                                                    _ = v37
                                                }
                                                v35 = (^(v35))
                                            } else {
                                                var v38 bool = ((1 <= 1) && v19)
                                                v39 := GenS4{F0: (fi36 >> (((p0 / ((420) | 1))) & 63)), F1: 2147483647, F2: ""}
                                                _ = v38
                                                _ = v39
                                            }
                                        case "a":
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            p0 -= p0
                                        case "\n\t":
                                            var v40 uint16 = 2
                                            _ = v40
                                        default:
                                            switch fi36 {
                                            case 1:
                                                var v41 int16 = 32767
                                                _ = v41
                                            case 601:
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                p0 = (^(p0))
                                                _ = len(v18)
                                                switch ("a" + ("\n\t" + "foo bar")) {
                                                case "":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if ((1 > 1000) && (fi36 <= -9223372036854775808)) {
                                                            fmt.Println(fi36, v19, v35)
                                                        }
                                                    case "":
                                                        if v19 {
                                                            var v42 int32 = -1
                                                            switch fi10 {
                                                            case 1:
                                                                var v43 int16 = 32767
                                                                _ = v43
                                                            case 25:
                                                                v19 = v19
                                                            default:
                                                                switch "a" {
                                                                case "a":
                                                                    p0 = (p0 % (((^(p0))) | 1))
                                                                case "foo bar":
                                                                    v44 := GenS5{F0: 374, F1: 182}
                                                                    _ = v44
                                                                default:
                                                                    var v45 float32 = 1e10
                                                                    _ = v45
                                                                }
                                                            }
                                                            if (((p0 * 718) > p0) || (fi10 == fi36)) {
                                                                switch (("" + (("foo bar" + "\n\t") + "foo bar")) + "hello") {
                                                                case "":
                                                                    v46 := GenS2{F0: 869}
                                                                    _ = v46
                                                                default:
                                                                    if len(v18) > 0 {
                                                                        v18[0] = (fi10 - (^(fi36)))
                                                                    }
                                                                }
                                                                var v47 uint = (p0 ^ uint(v35))
                                                                if (!false) {
                                                                    var v48 uint64 = 401
                                                                    _ = v48
                                                                }
                                                                if false {
                                                                    v35 -= 11
                                                                    if v19 {
                                                                        v18 = append(v18, -9223372036854775808)
                                                                        v42 = v42
                                                                        v19 = true
                                                                    } else {
                                                                        v49 := GenS3{F0: 1}
                                                                        _ = v49
                                                                    }
                                                                }
                                                                fmt.Println(v42, p0, p0)
                                                                if len(v18) > 0 {
                                                                    v18[0] = (fi36 / ((fi36) | 1))
                                                                }
                                                                _ = v47
                                                            }
                                                            v50 := GenS3{F0: 1000}
                                                            if (!(42 >= 1)) {
                                                                var v51 int8 = 0
                                                                fmt.Println(p0, fi36, v35)
                                                                var v52 int64 = 1
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            if len(v18) >= 2 {
                                                                v18 = v18[1:len(v18)]
                                                            }
                                                            _ = v42
                                                            _ = v50
                                                        } else {
                                                            var v53 float32 = 3.14159
                                                            _ = v53
                                                        }
                                                    case "foo bar":
                                                        v19 = (42 > 822)
                                                        var v54 uint = uint(v35)
                                                        _ = v54
                                                    }
                                                case "\n\t":
                                                    v55 := GenS6{F0: -33}
                                                    switch "\n\t" {
                                                    case "":
                                                        if v19 {
                                                            v35 -= 128
                                                        } else {
                                                            if (3.14159 == (0.1 + 0.1)) {
                                                                var v56 float64 = 0.1
                                                                var v57 int8 = -128
                                                                v58 := GenS3{F0: 42}
                                                                if (((("a" + ("" + (("\n\t" + "\n\t") + "hello"))) + "\n\t") + "\n\t") > ("a" + "")) {
                                                                    var v59 string = "foo bar"
                                                                    _ = v59
                                                                }
                                                                var v60 int16 = 36
                                                                v61 := GenS7{F0: 461, F1: 32767, F2: -1}
                                                                _ = v56
                                                                _ = v57
                                                                _ = v58
                                                                _ = v60
                                                                _ = v61
                                                            }
                                                        }
                                                    }
                                                    _ = v55
                                                case "foo bar":
                                                    if (v19 && v19) {
                                                        var v62 string = "\n\t"
                                                        var v63 uint = p0
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                            default:
                                                var v64 float64 = 3.14159
                                                _ = v64
                                            }
                                        }
                                        v65 := GenS8{F0: (p0 % ((65535) | 1))}
                                        _ = len(v18)
                                        _ = fi36
                                        _ = v65
                                    }
                                    var v66 uint8 = 255
                                    _ = v66
                                }
                                _ = v35
                            } else {
                                if v19 {
                                    _ = len(v18)
                                    for fi67 := 0; fi67 < 9; fi67++ {
                                        fmt.Println(v19)
                                        v19 = (!(true || (-745 <= 0)))
                                        v19 = (-2147483648 == (fi10 * 32768))
                                        fmt.Println(p0, fi67, v19)
                                        _ = fi67
                                    }
                                    switch 127 {
                                    case -1:
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                    }
                                } else {
                                    v68 := GenS3{F0: 1000}
                                    v19 = v19
                                    p0 = p0
                                    _ = v68
                                }
                            }
                        } else {
                            var v69 int64 = 42
                            for fi70 := 0; fi70 < 14; fi70++ {
                                switch fi70 {
                                case 9223372036854775807:
                                    v18 = append(v18, (fi10 + (fi70 % (((fi70 * (fi70 - (fi70 % ((127) | 1))))) | 1))))
                                case 589:
                                    var v71 uint8 = 25
                                    _ = v71
                                case 4294967295:
                                    if len(v18) > 0 {
                                        v18[0] = (fi70 & fi70)
                                    }
                                default:
                                    switch fi10 {
                                    case 128:
                                        fmt.Println(fi10, fi70, v19)
                                    case -2147483648:
                                        fmt.Println(v69)
                                    case -9223372036854775808:
                                        fmt.Println(fi10, v69)
                                    }
                                }
                                var v72 uint16 = 255
                                _ = fi70
                                _ = v72
                            }
                            var v73 uint = p0
                            _ = v69
                            _ = v73
                        }
                        v74 := GenS3{F0: 0}
                        v19 = (9223372036854775807 >= 216)
                        _ = v74
                    }
                    v19 = (1000 < 32767)
                    _ = v19
                } else {
                    v75 := GenS1{F0: 211, F1: "\n\t"}
                    _ = v75
                }
                _ = fi10
                _ = v18
            }
            fmt.Println(p0)
            var v76 uint = 402
            _ = v76
        } else {
            v77 := []int{32767, 919}
            _ = v77
        }
    } else {
        v78 := []int{36, 32767, -67}
        var v79 string = ("\n\t" + (("hello" + ("a" + "")) + "foo bar"))
        _ = v78
        _ = v79
    }
    if false {
        var v80 int64 = 0
        var v81 bool = true
        v82 := GenS9{F0: p0}
        _ = v80
        _ = v81
        _ = v82
    } else {
        var v83 uint32 = 2147483648
        _ = v83
    }
    if false {
        v84 := GenS10{F0: -62}
        switch -366 {
        case 128:
            switch -96 {
            case -38:
                fmt.Println(p0)
            case 997:
                if false {
                    if (!false) {
                        for fi85 := 0; fi85 < 1; fi85++ {
                            if ((479 > 658) && (2147483648 <= 2147483648)) {
                                v86 := []int{-107, (fi85 | 441)}
                                var v87 uint = 65535
                                _ = v86
                                _ = v87
                            } else {
                                var v88 int16 = -32768
                                _ = v88
                            }
                            p0 = (p0 * 1)
                            _ = fi85
                        }
                        switch (("a" + "a") + "a") {
                        case "\n\t":
                            for fi89 := 0; fi89 < 5; fi89++ {
                                p0 += (^(p0))
                                var v90 int16 = -32768
                                fmt.Println(v90, fi89, fi89)
                                _ = fi89
                                _ = v90
                            }
                        case "":
                            var v91 float64 = (0.1 + -2.25)
                            _ = v91
                        case "foo bar":
                            v92 := GenS8{F0: (p0 / (((p0 << (((p0 | (p0 % (((p0 - p0)) | 1)))) & 63))) | 1))}
                            _ = v92
                        default:
                            var v93 uint16 = 2
                            _ = v93
                        }
                    }
                }
            case 0:
                fmt.Println(p0)
            default:
                v94 := GenS8{F0: 255}
                _ = v94
            }
        default:
            var v95 uint64 = 42
            _ = v95
        }
        fmt.Println(p0)
        _ = v84
    }
    _ = p0
    return 65535
}

func main() {
    var v1 uint32 = 994
    var v2 float64 = ((float64(1.5) * -2.25) * 1.5)
    var v3 bool = (-2.25 != 0.0)
    v96 := []int{-179, 42, 1, 267}
    if false {
        v1 -= v1
        v97 := GenS11{F0: 9223372036854775807}
        var v98 int8 = -1
        _ = v97
        _ = v98
    } else {
        if len(v96) > 0 {
            v96[0] = -9223372036854775808
        }
        switch ("hello" + "") {
        case "\n\t":
            v3 = (v1 > (^(v1)))
        case "hello":
            var v99 float32 = 0.0
            _ = v99
        case "foo bar":
            for fi100 := 0; fi100 < 10; fi100++ {
                fmt.Println(fi100, v2, fi100)
                fmt.Println(v3, fi100, v1)
                _ = fi100
            }
        default:
            _ = len(v96)
        }
    }
    switch "foo bar" {
    case "a":
        for fi101 := 0; fi101 < 14; fi101++ {
            for fi102 := 0; fi102 < 10; fi102++ {
                if false {
                    switch ("" + "") {
                    case "a":
                        v96 = append(v96, -9223372036854775808)
                    case "":
                        if v3 {
                            v1 = (v1 % ((v1) | 1))
                            var v103 uint64 = 4294967295
                            _ = v103
                        } else {
                            switch "\n\t" {
                            case "hello":
                                var v104 int64 = -1
                                _ = v104
                            case "":
                                if len(v96) >= 2 {
                                    v96 = v96[1:len(v96)]
                                }
                                fmt.Println(fi102, v1, fi101)
                                if (-32768 <= 92) {
                                    var v105 float64 = (0.1 / -2.25)
                                    v106 := GenS3{F0: 1}
                                    var v107 uint64 = 546
                                    fmt.Println(v1)
                                    v96 = append(v96, fi102)
                                    if false {
                                        v2 = 1e10
                                        switch (fi102 % ((fi102) | 1)) {
                                        case -2147483648:
                                            v105 = v105
                                        case 2147483647:
                                            v1 += 318
                                        default:
                                            if ((!v3) || v3) {
                                                _ = len(v96)
                                                if v3 {
                                                    v96 = append(v96, (^(fi102)))
                                                } else {
                                                    v107 += 245
                                                    v105 -= (v2 / v105)
                                                    v3 = v3
                                                }
                                            }
                                        }
                                    }
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                } else {
                                    if v3 {
                                        v3 = (v1 <= 2147483647)
                                    }
                                    if len(v96) > 0 {
                                        v96[0] = (-(fi102))
                                    }
                                    fmt.Println(v2)
                                }
                            case "foo bar":
                                v1 = (v1 << ((4294967295) & 63))
                            }
                        }
                    case "\n\t":
                        _ = len(v96)
                    default:
                        var v108 uint8 = 0
                        _ = v108
                    }
                    v3 = (0 < 100)
                } else {
                    var v109 int32 = 723
                    v110 := GenS4{F0: fi101, F1: v1, F2: "hello"}
                    if v3 {
                        if ((0.0 * (1.5 - 1.5)) == (float64(0.0) + v2)) {
                            v96 = append(v96, -601)
                        } else {
                            var v111 float64 = ((0.0 / 1.0) * float64(-2.25))
                            _ = v111
                        }
                        v112 := GenS4{F0: (fi102 | (-(fi101))), F1: v1, F2: "a"}
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2)
                        default:
                            v1 = 957
                        }
                        _ = v112
                    } else {
                        v109 -= v109
                    }
                    _ = v109
                    _ = v110
                }
                var v113 float32 = 1e10
                _ = fi102
                _ = v113
            }
            _ = fi101
        }
    case "hello":
        v2 -= v2
    case "":
        v114 := GenS2{F0: 1}
        _ = v114
    default:
        for fi115 := 0; fi115 < 19; fi115++ {
            switch 0 {
            case -1:
                v1 -= (v1 + 2147483648)
            case -2147483648:
                v1 -= v1
            }
            _ = fi115
        }
    }
    v2 -= v2
    v116 := GenS7{F0: 42, F1: 2, F2: 1}
    v96 = append(v96, 32768)
    var v117 uint32 = v1
    for fi118 := 0; fi118 < 1; fi118++ {
        switch "\n\t" {
        case "foo bar":
            for fi119 := 0; fi119 < 10; fi119++ {
                v117 = (v1 % (((^(v117))) | 1))
                if len(v96) > 0 {
                    v96[0] = fi119
                }
                switch fi118 {
                case -287:
                    v120 := GenS9{F0: 18446744073709551615}
                    _ = v120
                case 2147483647:
                    v96 = append(v96, (fi118 - fi118))
                    if true {
                        v117 = (v1 - (v117 << ((65535) & 63)))
                        var v121 uint32 = v117
                        var v122 int64 = 519
                        _ = v121
                        _ = v122
                    }
                case 32768:
                    var v123 uint = 65535
                    _ = v123
                }
                _ = fi119
            }
        case "hello":
            var v124 int32 = -955
            _ = v124
        case "\n\t":
            if (!v3) {
                v2 = v2
                if len(v96) > 0 {
                    v96[0] = int(v117)
                }
            } else {
                var v125 int32 = -2147483648
                _ = v125
            }
        default:
            if (!v3) {
                _ = len(v96)
                v126 := GenS6{F0: 1}
                _ = v126
            }
        }
        v2 -= v2
        v117 += (v117 << ((2) & 63))
        _ = fi118
    }
    if (127 == 1) {
        v117 = v117
        fmt.Println(v1, v2)
    } else {
        switch "foo bar" {
        case "hello":
            v117 = (v117 << ((genFn0(255)) & 63))
        case "foo bar":
            if v3 {
                v96 = append(v96, 128)
            }
        case "\n\t":
            switch -333 {
            case 9223372036854775807:
                fmt.Println(v2)
            }
        default:
            v1 = 0
        }
    }
    fmt.Println(v3, v3, v117)
    v127 := GenS10{F0: -1}
    v117 = 634
    for fi128 := 0; fi128 < 12; fi128++ {
        v129 := GenS12{F0: 659, F1: 42}
        for fi130 := 0; fi130 < 16; fi130++ {
            if v3 {
                if v3 {
                    v2 -= v2
                    switch ("hello" + "") {
                    case "\n\t":
                        var v131 int64 = 1
                        _ = v131
                    default:
                        switch ("a" + (("hello" + "foo bar") + "")) {
                        case "hello":
                            _ = len(v96)
                        case "":
                            var v132 int64 = 42
                            _ = v132
                        case "\n\t":
                            var v133 string = "a"
                            _ = v133
                        }
                    }
                } else {
                    v134 := GenS13{F0: 1000}
                    if false {
                        var v135 uint8 = 115
                        switch "a" {
                        case "\n\t":
                            switch fi130 {
                            case 289:
                                if (-128 > 100) {
                                    if len(v96) > 0 {
                                        v96[0] = (fi128 << ((255) & 63))
                                    }
                                    v136 := GenS14{F0: 590, F1: 2}
                                    _ = v136
                                } else {
                                    if (0.1 != (((v2 + 0.0) * v2) * -2.25)) {
                                        var v137 uint64 = 695
                                        v135 = (^(v135))
                                        fmt.Println(v135, v3, fi130)
                                        v1 = (v117 & v117)
                                        _ = v137
                                    } else {
                                        v138 := GenS4{F0: (fi128 * fi130), F1: v117, F2: "hello"}
                                        _ = v138
                                    }
                                    if v3 {
                                        var v139 bool = (int8(v1) > -1)
                                        if (0 < 1) {
                                            v2 = (float64(3.14159) + (0.0 - v2))
                                            v1 = v1
                                            v96 = append(v96, (fi130 ^ (fi130 << ((42) & 63))))
                                        } else {
                                            v117 -= 4294967295
                                            var v140 int16 = 0
                                            var v141 int8 = 1
                                            _ = v140
                                            _ = v141
                                        }
                                        _ = v139
                                    }
                                    var v142 int16 = 42
                                    _ = v142
                                }
                            case 77:
                                if ((v117 << ((2) & 63)) == 1) {
                                    var v143 uint8 = v135
                                    _ = v143
                                }
                            }
                        case "":
                            v117 -= (^(v1))
                        }
                        v144 := GenS11{F0: 42}
                        _ = v135
                        _ = v144
                    } else {
                        if len(v96) > 0 {
                            v96[0] = (fi128 & (fi128 >> ((65535) & 63)))
                        }
                        var v145 uint32 = 4294967295
                        v1 = v145
                        _ = v145
                    }
                    _ = v134
                }
                fmt.Println(v1, v1, v3)
                var v146 uint32 = v117
                v117 = 2
                _ = v146
            } else {
                var v147 int = (fi128 - fi128)
                switch -9223372036854775808 {
                case -2147483648:
                    fmt.Println(v2)
                case 1:
                    if v3 {
                        v148 := GenS2{F0: int16(fi130)}
                        v3 = true
                        if v3 {
                            fmt.Println(v117, v2, v3)
                            if v3 {
                                v149 := GenS15{F0: 204}
                                v2 -= float64((0.1 / 1.5))
                                if len(v96) >= 2 {
                                    v96 = v96[1:len(v96)]
                                }
                                _ = v149
                            }
                        }
                        _ = v148
                    } else {
                        fmt.Println(v117, fi130)
                        var v150 uint32 = (v117 / ((v1) | 1))
                        _ = v150
                    }
                }
                v151 := GenS13{F0: 32767}
                _ = v147
                _ = v151
            }
            v1 += v117
            v1 += (^(v117))
            _ = fi130
        }
        _ = fi128
        _ = v129
    }
    switch -2147483648 {
    case 230:
        fmt.Println(v2)
    case 876:
        v1 = v117
    case -807:
        switch 870 {
        case 9223372036854775807:
            if false {
                fmt.Println(v3, v3)
                v1 += v117
            }
        }
    default:
        v1 -= v117
    }
    if (88 > 42) {
        v2 -= (float64(0.0) / 1.5)
        for fi152 := 0; fi152 < 14; fi152++ {
            var v153 uint64 = 221
            _ = fi152
            _ = v153
        }
        switch "foo bar" {
        case "foo bar":
            v2 -= v2
        case "a":
            var v154 uint64 = 18446744073709551615
            _ = v154
        case "\n\t":
            for fi155 := 0; fi155 < 10; fi155++ {
                if v3 {
                    var v156 string = (("\n\t" + "hello") + "")
                    v1 -= v1
                    _ = v156
                } else {
                    var v157 uint = 4294967295
                    for fi158 := 0; fi158 < 1; fi158++ {
                        v3 = v3
                        var v159 string = "\n\t"
                        v3 = ((v157 - 42) == v157)
                        var v160 int = fi155
                        _ = fi158
                        _ = v159
                        _ = v160
                    }
                    v161 := GenS15{F0: v157}
                    _ = v157
                    _ = v161
                }
                _ = fi155
            }
        default:
            v162 := GenS1{F0: 27, F1: "hello"}
            _ = v162
        }
    } else {
        var v163 int64 = 42
        _ = v163
    }
    if v3 {
        v117 = 2147483647
        for fi164 := 0; fi164 < 15; fi164++ {
            var v165 uint16 = 0
            switch 0 {
            case 42:
                switch "foo bar" {
                case "foo bar":
                    switch 42 {
                    case 32767:
                        _ = len(v96)
                    default:
                        var v166 float32 = 1.5
                        _ = v166
                    }
                case "a":
                    v167 := GenS16{F0: (1000 <= 255), F1: (v165 - uint16(v1)), F2: false}
                    _ = v167
                default:
                    var v168 bool = false
                    _ = v168
                }
            case -1:
                v169 := GenS3{F0: 0}
                _ = v169
            default:
                v1 += 2
            }
            if ("\n\t" >= "a") {
                var v170 int64 = 9223372036854775807
                _ = v170
            }
            _ = fi164
            _ = v165
        }
        switch 128 {
        case -379:
            for fi171 := 0; fi171 < 13; fi171++ {
                v1 += 286
                v172 := GenS17{F0: 42, F1: 42, F2: 2147483647}
                _ = fi171
                _ = v172
            }
        case 623:
            v96 = append(v96, 983)
        case 0:
            fmt.Println(v2, v1, v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v117:", v117)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v96:", len(v96), v96)
    _ = v1
    _ = v116
    _ = v117
    _ = v127
    _ = v2
    _ = v3
    _ = v96
}
