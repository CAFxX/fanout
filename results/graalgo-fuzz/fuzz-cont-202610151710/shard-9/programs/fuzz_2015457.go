package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint64; F1 uint; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int; F1 int; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 uint32 }
type GenS12 struct { F0 int; F1 uint64; F2 int32 }
type GenS13 struct { F0 int32; F1 string; F2 uint32 }
type GenS14 struct { F0 uint16; F1 uint16 }
type GenS15 struct { F0 uint64; F1 uint64; F2 int16 }
type GenS16 struct { F0 uint8; F1 bool; F2 bool }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 uint16; F1 string }
type GenS19 struct { F0 uint8 }
type GenS20 struct { F0 string; F1 uint32; F2 uint8 }
type GenS21 struct { F0 int32; F1 string; F2 uint32 }
type GenS22 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS23 struct { F0 uint64 }
type GenS24 struct { F0 uint; F1 uint8 }
type GenS25 struct { F0 int32; F1 int16 }
type GenS26 struct { F0 uint8; F1 string }
type GenS27 struct { F0 uint; F1 uint8 }
type GenS28 struct { F0 int32; F1 uint16; F2 int64 }
func genFn0(p0 int64, p1 int16) bool {
    p1 = (p1 / (((p1 + 32767)) | 1))
    var v4 int8 = -128
    v4 = 42
    if true {
        if (2147483647 == 763) {
            if true {
                if (-106 != -52) {
                    v5 := []int{-9223372036854775808}
                    _ = v5
                } else {
                    if true {
                        v4 = v4
                    } else {
                        p1 -= p1
                    }
                    p1 = p1
                }
                var v6 uint16 = 32767
                switch "foo bar" {
                case "\n\t":
                    if true {
                        v7 := GenS1{F0: v6, F1: 2}
                        _ = v7
                    } else {
                        var v8 int8 = -128
                        for fi9 := 0; fi9 < 11; fi9++ {
                            switch (fi9 >> ((4294967295) & 63)) {
                            case 32768:
                                v4 -= (v4 ^ 42)
                            case -768:
                                var v10 int64 = p0
                                _ = v10
                            }
                            v11 := GenS2{F0: (v6 << ((730) & 63)), F1: -1}
                            _ = fi9
                            _ = v11
                        }
                        switch "foo bar" {
                        case "\n\t":
                            var v12 float32 = ((0.0 * float32(3.14159)) * float32(0.0))
                            _ = v12
                        case "foo bar":
                            if (0 >= (v6 / (((v6 - v6)) | 1))) {
                                var v13 uint64 = 42
                                if (42 == v8) {
                                    if true {
                                        v6 += (v6 - v6)
                                        var v14 uint = 1
                                        _ = v14
                                    } else {
                                        v15 := GenS2{F0: (v6 / ((255) | 1)), F1: v8}
                                        switch ("" + "\n\t") {
                                        case "foo bar":
                                            p0 -= 0
                                        case "hello":
                                            var v16 int8 = (v4 % (((v4 | -128)) | 1))
                                            _ = v16
                                        case "\n\t":
                                            p0 += p0
                                        }
                                        for fi17 := 0; fi17 < 11; fi17++ {
                                            fmt.Println(v6, v8, v13)
                                            v6 -= (v6 << ((1) & 63))
                                            _ = fi17
                                        }
                                        _ = v15
                                    }
                                    for fi18 := 0; fi18 < 6; fi18++ {
                                        switch fi18 {
                                        case 0:
                                            for fi19 := 0; fi19 < 20; fi19++ {
                                                var v20 int64 = (-(p0))
                                                _ = fi19
                                                _ = v20
                                            }
                                        case 185:
                                            v21 := GenS3{F0: "a", F1: v13}
                                            _ = v21
                                        }
                                        v22 := GenS1{F0: (v6 % ((65535) | 1)), F1: v13}
                                        v8 += 127
                                        _ = fi18
                                        _ = v22
                                    }
                                    v23 := []int{}
                                    _ = v23
                                }
                                _ = v13
                            } else {
                                var v24 uint32 = 2147483648
                                for fi25 := 0; fi25 < 18; fi25++ {
                                    v26 := GenS4{F0: 0, F1: 729, F2: (p1 - 1000)}
                                    switch (("\n\t" + (("a" + "foo bar") + "")) + ("hello" + "foo bar")) {
                                    case "":
                                        for fi27 := 0; fi27 < 10; fi27++ {
                                            if true {
                                                v28 := []int{fi27, fi25, fi25}
                                                v29 := GenS2{F0: v6, F1: (v8 / ((0) | 1))}
                                                _ = v28
                                                _ = v29
                                            }
                                            v4 = v8
                                            switch ("a" + ("foo bar" + ("" + ("a" + "\n\t")))) {
                                            case "\n\t":
                                                v30 := []int{fi25, (-(fi25)), fi27, -675}
                                                _ = v30
                                            case "a":
                                                var v31 int64 = int64(v6)
                                                _ = v31
                                            }
                                            v4 = (^(v4))
                                            _ = fi27
                                        }
                                    default:
                                        fmt.Println(v4, v8, p1)
                                    }
                                    _ = fi25
                                    _ = v26
                                }
                                _ = v24
                            }
                        default:
                            v32 := GenS5{F0: uint32(v6)}
                            _ = v32
                        }
                        _ = v8
                    }
                }
                _ = v6
            }
            fmt.Println(p0, p0)
        } else {
            var v33 int32 = 2147483647
            if true {
                if (((p1 << ((876) & 63)) < -107) || (0 >= 704)) {
                    var v34 bool = true
                    v35 := []int{4294967295, -1}
                    _ = v34
                    _ = v35
                }
            }
            _ = v33
        }
    } else {
        var v36 uint = 0
        p0 = (p0 * 9223372036854775807)
        for fi37 := 0; fi37 < 8; fi37++ {
            var v38 uint8 = 10
            switch "\n\t" {
            case "foo bar":
                var v39 int32 = 718
                _ = v39
            case "hello":
                v4 = (v4 & v4)
            default:
                v40 := []int{fi37, (fi37 * (fi37 >> ((1) & 63))), 0, 127}
                _ = v40
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                v42 := GenS5{F0: 4294967295}
                switch (fi41 | -2147483648) {
                case -221:
                    if (!(v36 != v36)) {
                        switch ("" + "foo bar") {
                        case "a":
                            v43 := []int{(-(fi37)), -454}
                            _ = v43
                        case "foo bar":
                            var v44 uint32 = 177
                            _ = v44
                        case "":
                            v45 := []int{(-(fi37))}
                            _ = v45
                        }
                        fmt.Println(fi37)
                    } else {
                        v46 := []int{(fi41 ^ (fi41 >> (((^(v36))) & 63))), fi41, (-(fi41))}
                        var v47 uint64 = 958
                        _ = v46
                        _ = v47
                    }
                case -465:
                    var v48 uint16 = 0
                    _ = v48
                case -252:
                    var v49 uint = 4294967295
                    _ = v49
                default:
                    switch fi41 {
                    case -240:
                        fmt.Println(p1)
                    case 353:
                        var v50 int64 = (p0 % (((p0 + (p0 * (p0 | (p0 / (((p0 + (p0 & (p0 ^ p0)))) | 1)))))) | 1))
                        v51 := GenS2{F0: 2, F1: (v4 + (v4 >> ((v36) & 63)))}
                        _ = v50
                        _ = v51
                    }
                }
                fmt.Println(fi41, v4)
                if true {
                    var v52 uint64 = 389
                    _ = v52
                } else {
                    v53 := GenS6{F0: (v36 ^ 376)}
                    var v54 int64 = p0
                    _ = v53
                    _ = v54
                }
                _ = fi41
                _ = v42
            }
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    for fi55 := 0; fi55 < 10; fi55++ {
        if true {
            if (4294967295 <= 524) {
                switch ("foo bar" + "hello") {
                case "a":
                    if (p1 < p1) {
                        if (42 < 100) {
                            p0 -= p0
                            var v56 uint8 = 2
                            if true {
                                p1 = (p1 | 1)
                            } else {
                                for fi57 := 0; fi57 < 16; fi57++ {
                                    if true {
                                        v56 -= 2
                                    }
                                    _ = fi57
                                }
                                switch ("\n\t" + ("hello" + ("a" + "\n\t"))) {
                                case "":
                                    v58 := GenS5{F0: 65535}
                                    _ = v58
                                case "foo bar":
                                    if false {
                                        switch fi55 {
                                        case 4294967295:
                                            for fi59 := 0; fi59 < 3; fi59++ {
                                                if ((false || true) || (631 == 1)) {
                                                    var v60 int16 = p1
                                                    _ = v60
                                                } else {
                                                    v61 := GenS7{F0: -738, F1: 563, F2: p1}
                                                    v62 := GenS8{F0: ((!(3.14159 > ((((1.5 * (0.0 + (3.14159 * -2.25))) - float32(1.5)) + (0.1 / float32(1.5))) / 1.0))) && false)}
                                                    switch fi59 {
                                                    case 128:
                                                        v63 := []int{-9223372036854775808, (fi55 | (^(fi59))), 9223372036854775807, -9223372036854775808}
                                                        _ = v63
                                                    case 42:
                                                        v64 := []int{(fi59 - fi59), (fi55 >> ((65535) & 63)), (fi59 | (-(fi59)))}
                                                        _ = v64
                                                    }
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v65 uint8 = v56
                                                _ = fi59
                                                _ = v65
                                            }
                                        case -9223372036854775808:
                                            var v66 int64 = -43
                                            _ = v66
                                        default:
                                            fmt.Println(v4, v4)
                                        }
                                    }
                                default:
                                    v67 := GenS9{F0: p1}
                                    _ = v67
                                }
                                v4 += v4
                            }
                            _ = v56
                        } else {
                            v68 := GenS10{F0: ((false || false) || (((p0 & int64(v4)) < (p0 * p0)) || true))}
                            if true {
                                v69 := []int{1, (fi55 & fi55), fi55}
                                _ = v69
                            } else {
                                if (-677 >= -288) {
                                    var v70 uint64 = 1
                                    var v71 int64 = p0
                                    if (42 >= 32767) {
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v72 string = "foo bar"
                                            _ = v72
                                        default:
                                            switch -2147483648 {
                                            case -9223372036854775808:
                                                v4 -= (v4 * v4)
                                            default:
                                                for fi73 := 0; fi73 < 17; fi73++ {
                                                    v74 := GenS11{F0: 438, F1: 4294967295}
                                                    fmt.Println(p0, v4, p1)
                                                    _ = fi73
                                                    _ = v74
                                                }
                                            }
                                        }
                                    } else {
                                        var v75 int32 = -19
                                        var v76 float64 = (float64(-2.25) * 0.0)
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = v70
                                    _ = v71
                                } else {
                                    for fi77 := 0; fi77 < 20; fi77++ {
                                        var v78 int = int(v4)
                                        _ = fi77
                                        _ = v78
                                    }
                                    if (false || (((false || (!false)) || false) && (42 > 4294967295))) {
                                        v79 := []int{fi55, fi55, (fi55 + (^(fi55))), fi55}
                                        v4 -= (-(v4))
                                        v4 -= (v4 ^ v4)
                                        _ = v79
                                    } else {
                                        fmt.Println(p1, v4)
                                    }
                                }
                                v80 := []int{(fi55 ^ fi55), fi55, (fi55 % (((^(fi55))) | 1)), fi55}
                                for fi81 := 0; fi81 < 6; fi81++ {
                                    var v82 float64 = 1.5
                                    var v83 uint8 = 1
                                    if len(v80) >= 2 {
                                        v80 = v80[1:len(v80)]
                                    }
                                    if (18446744073709551615 != 4294967295) {
                                        v82 = 3.14159
                                        var v84 uint64 = 582
                                        _ = v84
                                    } else {
                                        v82 = (v82 * 0.0)
                                        if len(v80) > 0 {
                                            v80[0] = 2147483647
                                        }
                                        switch "hello" {
                                        case "hello":
                                            v80 = append(v80, fi55)
                                        case "a":
                                            var v85 int32 = -878
                                            _ = v85
                                        default:
                                            if len(v80) >= 2 {
                                                v80 = v80[1:len(v80)]
                                            }
                                        }
                                        v86 := GenS1{F0: 0, F1: 2}
                                        if len(v80) >= 2 {
                                            v80 = v80[1:len(v80)]
                                        }
                                        _ = v86
                                    }
                                    var v87 uint32 = 4294967295
                                    if (((("a" + "\n\t") + "\n\t") + ("hello" + ("hello" + ("a" + "hello")))) > "") {
                                        p0 += (p0 * p0)
                                    } else {
                                        var v88 int16 = p1
                                        _ = v88
                                    }
                                    if (903 < 4294967295) {
                                        v89 := GenS12{F0: fi55, F1: 2, F2: 2147483647}
                                        fmt.Println(v87, v4, fi55)
                                        _ = v89
                                    } else {
                                        if len(v80) > 0 {
                                            v80[0] = -959
                                        }
                                    }
                                    p1 = 32767
                                    var v90 uint64 = 1
                                    _ = fi81
                                    _ = v82
                                    _ = v83
                                    _ = v87
                                    _ = v90
                                }
                                _ = v80
                            }
                            var v91 uint32 = 42
                            _ = v68
                            _ = v91
                        }
                        v92 := GenS11{F0: 2147483647, F1: 941}
                        _ = v92
                    } else {
                        switch "foo bar" {
                        case "hello":
                            v93 := []int{}
                            _ = v93
                        case "":
                            var v94 string = ""
                            _ = v94
                        case "a":
                            fmt.Println(p1)
                        }
                    }
                default:
                    for fi95 := 0; fi95 < 6; fi95++ {
                        v4 = v4
                        var v96 uint = 0
                        if true {
                            if (106 < 797) {
                                if true {
                                    var v97 uint = 236
                                    fmt.Println(v96)
                                    var v98 int = (fi55 & fi55)
                                    _ = v97
                                    _ = v98
                                } else {
                                    if true {
                                        if (((-2.25 / ((((float32((((0.1 * 0.0) * ((1.5 + 0.0) + 1e10)) * ((((1e10 + 3.14159) * (-2.25 - (0.1 + 3.14159))) - 0.1) * (3.14159 * ((0.1 * 1e10) - ((0.0 / 1.0) + float64(0.1))))))) - ((1.5 / 0.1) - float32(((0.1 / (-2.25 / 1.5)) - -2.25)))) + 0.1) - 1e10) - 3.14159)) - 0.1) <= (-2.25 / (-2.25 * 3.14159))) {
                                            p0 = p0
                                            v99 := []int{fi55, fi95, 0}
                                            _ = v99
                                        } else {
                                            v100 := GenS13{F0: -971, F1: "a", F2: 1}
                                            _ = v100
                                        }
                                        if (42 >= 128) {
                                            v96 = v96
                                            var v101 int8 = (v4 >> ((v96) & 63))
                                            if (127 == 2) {
                                                if false {
                                                    v102 := []int{(fi95 / (((fi55 - fi95)) | 1)), 127, fi95}
                                                    var v103 uint = (v96 >> (((v96 >> ((896) & 63))) & 63))
                                                    _ = v102
                                                    _ = v103
                                                }
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    if false {
                                                        var v104 float64 = 1e10
                                                        if (float64((1.5 - (3.14159 * 3.14159))) > 3.14159) {
                                                            v105 := GenS14{F0: 255, F1: 1}
                                                            v106 := GenS15{F0: 2, F1: 321, F2: (p1 % ((-32768) | 1))}
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            var v107 string = ""
                                                            if true {
                                                                v108 := GenS1{F0: 32767, F1: 4294967295}
                                                                v101 = (v101 >> ((2) & 63))
                                                                v109 := []int{9223372036854775807, (fi55 & (fi55 >> ((2) & 63))), (fi95 | fi95)}
                                                                _ = v108
                                                                _ = v109
                                                            }
                                                            _ = v107
                                                        }
                                                        _ = v104
                                                    }
                                                default:
                                                    var v110 int8 = (-(v4))
                                                    _ = v110
                                                }
                                            }
                                            _ = v101
                                        }
                                    }
                                    p1 -= -32768
                                }
                                switch "foo bar" {
                                case "a":
                                    var v111 uint64 = 42
                                    _ = v111
                                default:
                                    if (1.5 <= ((1e10 - 0.0) + 1.5)) {
                                        v112 := []int{}
                                        var v113 int32 = 961
                                        p1 = -32768
                                        _ = v112
                                        _ = v113
                                    } else {
                                        var v114 int8 = int8(p0)
                                        var v115 int8 = v114
                                        _ = v114
                                        _ = v115
                                    }
                                }
                                switch (-(fi55)) {
                                case -9223372036854775808:
                                    v4 = v4
                                case 217:
                                    var v116 int8 = (^(v4))
                                    _ = v116
                                case 2147483647:
                                    p0 += (^(p0))
                                }
                            } else {
                                var v117 uint = 1
                                _ = v117
                            }
                            p0 -= p0
                            v118 := GenS2{F0: 990, F1: v4}
                            _ = v118
                        } else {
                            if (false && (620 != 42)) {
                                p0 += p0
                                p1 += p1
                                v119 := []int{fi55, (fi95 | (^(fi95))), (fi95 % ((fi55) | 1))}
                                _ = v119
                            }
                            fmt.Println(v96, p1, fi55)
                            switch "foo bar" {
                            case "\n\t":
                                v120 := GenS1{F0: 1, F1: 2}
                                _ = v120
                            default:
                                var v121 bool = (!true)
                                _ = v121
                            }
                            fmt.Println(v96, p0)
                            switch "" {
                            case "\n\t":
                                fmt.Println(p1)
                            case "":
                                var v122 int32 = 462
                                _ = v122
                            }
                        }
                        _ = fi95
                        _ = v96
                    }
                }
                var v123 int = fi55
                _ = v123
            } else {
                if (true && (false && (18446744073709551615 <= 2))) {
                    if true {
                        var v124 uint32 = 2147483648
                        switch (-(fi55)) {
                        case 786:
                            if (v124 == v124) {
                                p1 = p1
                            }
                        case 32768:
                            p1 -= 0
                        case -793:
                            p1 -= p1
                        default:
                            if true {
                                var v125 float64 = 0.1
                                v125 -= v125
                                switch ("foo bar" + "a") {
                                case "":
                                    var v126 int16 = (p1 >> ((1) & 63))
                                    _ = v126
                                case "a":
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        v127 := GenS1{F0: 904, F1: 18446744073709551615}
                                        _ = v127
                                    case "a":
                                        v125 -= 1e10
                                    case "\n\t":
                                        switch "a" {
                                        case "hello":
                                            switch "a" {
                                            case "foo bar":
                                                var v128 int = (^(fi55))
                                                _ = v128
                                            }
                                        case "foo bar":
                                            for fi129 := 0; fi129 < 5; fi129++ {
                                                if false {
                                                    p0 = p0
                                                } else {
                                                    v130 := []int{fi129}
                                                    var v131 bool = false
                                                    v132 := GenS3{F0: "hello", F1: 42}
                                                    var v133 string = ((("hello" + "foo bar") + "hello") + ("\n\t" + ""))
                                                    if v131 {
                                                        if v131 {
                                                            var v134 int = (-(fi55))
                                                            if ((3.14159 < float32(v125)) && true) {
                                                                v135 := GenS16{F0: 255, F1: v131, F2: v131}
                                                                v124 -= 65535
                                                                if v131 {
                                                                    if (1 > 18446744073709551615) {
                                                                        if v131 {
                                                                            var v136 uint16 = 32767
                                                                            var v137 int16 = (p1 - (p1 + (p1 << ((0) & 63))))
                                                                            var v138 float64 = (v125 - -2.25)
                                                                            _ = v136
                                                                            _ = v137
                                                                            _ = v138
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v139 int32 = 0
                                                                    _ = v139
                                                                }
                                                                v140 := GenS5{F0: v124}
                                                                _ = v135
                                                                _ = v140
                                                            } else {
                                                                if v131 {
                                                                    if len(v130) > 0 {
                                                                        v130[0] = (fi55 * (fi129 >> ((42) & 63)))
                                                                    }
                                                                    fmt.Println(p0)
                                                                    v134 = fi129
                                                                } else {
                                                                    var v141 uint8 = 127
                                                                    switch (("\n\t" + (v133 + v133)) + (v133 + v133)) {
                                                                    case "hello":
                                                                        fmt.Println(v124, v134)
                                                                    case "":
                                                                        if len(v130) > 0 {
                                                                            v130[0] = (fi129 >> ((4294967295) & 63))
                                                                        }
                                                                    case "foo bar":
                                                                        var v142 uint8 = 1
                                                                        _ = v142
                                                                    default:
                                                                        if (fi129 > (fi129 - (fi129 & -320))) {
                                                                            v143 := GenS14{F0: 1, F1: 65535}
                                                                            if ((34 == uint16(v124)) || (1.5 <= 1.5)) {
                                                                                _ = len(v130)
                                                                            } else {
                                                                                p0 = (-(p0))
                                                                                v144 := GenS17{F0: 1}
                                                                                fmt.Println(v133, v131, v124)
                                                                                _ = v144
                                                                            }
                                                                            var v145 uint64 = 1
                                                                            _ = v143
                                                                            _ = v145
                                                                        } else {
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v146 string = v133
                                                                                    _ = v146
                                                                                case "\n\t":
                                                                                    v125 = float64(3.14159)
                                                                                default:
                                                                                    v147 := GenS16{F0: (v141 / ((v141) | 1)), F1: (0 != 2), F2: v131}
                                                                                    _ = v147
                                                                                }
                                                                            case "a":
                                                                                _ = len(v130)
                                                                            case "foo bar":
                                                                                if len(v130) > 0 {
                                                                                    v130[0] = -346
                                                                                }
                                                                                var v148 int64 = (-(p0))
                                                                                if (614 == 1) {
                                                                                    v130 = append(v130, fi55)
                                                                                    v149 := GenS18{F0: 65535, F1: v133}
                                                                                    _ = v149
                                                                                }
                                                                                _ = v148
                                                                            default:
                                                                                p1 = 42
                                                                            }
                                                                        }
                                                                    }
                                                                    v130 = append(v130, (v134 + (^(v134))))
                                                                    if (100 > -128) {
                                                                        p0 = p0
                                                                    } else {
                                                                        var v150 int32 = 2147483647
                                                                        _ = v150
                                                                    }
                                                                    _ = v141
                                                                }
                                                                var v151 int16 = 32767
                                                                _ = v151
                                                            }
                                                            _ = v134
                                                        }
                                                        fmt.Println(p0)
                                                        _ = len(v130)
                                                    }
                                                    fmt.Println(v133, fi55, v124)
                                                    _ = v130
                                                    _ = v131
                                                    _ = v132
                                                    _ = v133
                                                }
                                                if true {
                                                    v152 := []int{}
                                                    fmt.Println(v4)
                                                    _ = v152
                                                }
                                                v124 = 2147483648
                                                v4 += v4
                                                fmt.Println(p1, fi129, fi55)
                                                v4 += (v4 / ((v4) | 1))
                                                fmt.Println(fi55, p1, v124)
                                                switch fi55 {
                                                case 1:
                                                    var v153 int16 = p1
                                                    _ = v153
                                                case 4294967295:
                                                    var v154 uint = 42
                                                    _ = v154
                                                case 187:
                                                    if (1 <= -2147483648) {
                                                        var v155 uint32 = v124
                                                        _ = v155
                                                    } else {
                                                        switch "a" {
                                                        case "foo bar":
                                                            p1 = -405
                                                        case "a":
                                                            var v156 int64 = int64(p1)
                                                            _ = v156
                                                        case "hello":
                                                            var v157 string = "a"
                                                            _ = v157
                                                        default:
                                                            var v158 uint8 = 0
                                                            _ = v158
                                                        }
                                                        p0 = (p0 + 0)
                                                    }
                                                default:
                                                    if true {
                                                        if true {
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if (float32(v125) <= 1e10) {
                                                                    var v159 uint32 = (v124 | (v124 >> ((221) & 63)))
                                                                    switch (("hello" + "a") + "") {
                                                                    case "":
                                                                        if true {
                                                                            switch ("hello" + ("" + (("hello" + ("hello" + "foo bar")) + "a"))) {
                                                                            case "":
                                                                                switch -540 {
                                                                                case -398:
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        v160 := GenS19{F0: 176}
                                                                                        _ = v160
                                                                                    case "hello":
                                                                                        if false {
                                                                                            var v161 float32 = 1e10
                                                                                            fmt.Println(v125)
                                                                                            v124 = 2147483648
                                                                                            _ = v161
                                                                                        }
                                                                                        switch fi129 {
                                                                                        case 9223372036854775807:
                                                                                            var v162 uint = 343
                                                                                            _ = v162
                                                                                        default:
                                                                                            var v163 uint64 = 1
                                                                                            _ = v163
                                                                                        }
                                                                                        v164 := []int{fi55}
                                                                                        _ = v164
                                                                                    default:
                                                                                        fmt.Println(v4)
                                                                                    }
                                                                                default:
                                                                                    p1 += (p1 ^ p1)
                                                                                }
                                                                            case "hello":
                                                                                v165 := []int{32768, fi55}
                                                                                _ = v165
                                                                            }
                                                                            v124 = 2147483647
                                                                            v166 := GenS20{F0: "hello", F1: 1, F2: 0}
                                                                            _ = v166
                                                                        }
                                                                        if ((true || (!true)) && true) {
                                                                            switch (fi129 >> ((731) & 63)) {
                                                                            case 32767:
                                                                                fmt.Println(p1)
                                                                            case -2147483648:
                                                                                v125 -= float64((0.0 + 0.1))
                                                                            case -558:
                                                                                v167 := GenS17{F0: (p0 * (p0 ^ (p0 | p0)))}
                                                                                _ = v167
                                                                            default:
                                                                                p0 = (^(p0))
                                                                            }
                                                                            v168 := []int{-751, fi55}
                                                                            _ = v168
                                                                        }
                                                                        var v169 int16 = (p1 - p1)
                                                                        v170 := []int{fi129, (fi129 << ((18446744073709551615) & 63))}
                                                                        v125 += v125
                                                                        v125 += v125
                                                                        _ = v169
                                                                        _ = v170
                                                                    case "foo bar":
                                                                        v124 -= v124
                                                                    case "hello":
                                                                        v171 := []int{}
                                                                        _ = v171
                                                                    default:
                                                                        v124 -= 809
                                                                    }
                                                                    fmt.Println(v124, p0)
                                                                    switch (fi55 & fi55) {
                                                                    case -2147483648:
                                                                        v172 := GenS21{F0: 2147483647, F1: "a", F2: (v159 / ((v159) | 1))}
                                                                        _ = v172
                                                                    case -507:
                                                                        if ((true && (p1 > p1)) && (p0 <= 0)) {
                                                                            v173 := GenS16{F0: 2, F1: ((!(!(false || true))) || false), F2: true}
                                                                            fmt.Println(fi129)
                                                                            v174 := GenS22{F0: 42, F1: 0, F2: -100}
                                                                            if (-2.25 > 0.1) {
                                                                                v175 := []int{fi129}
                                                                                v176 := GenS18{F0: 614, F1: ("\n\t" + "a")}
                                                                                _ = v175
                                                                                _ = v176
                                                                            }
                                                                            v159 -= (v159 / (((v159 - v159)) | 1))
                                                                            if true {
                                                                                v177 := GenS8{F0: (("hello" >= "hello") || false)}
                                                                                _ = v177
                                                                            } else {
                                                                                v4 = 127
                                                                            }
                                                                            switch 4294967295 {
                                                                            case 2147483647:
                                                                                v178 := GenS23{F0: 904}
                                                                                _ = v178
                                                                            case 1:
                                                                                var v179 float64 = 3.14159
                                                                                var v180 uint32 = v124
                                                                                _ = v179
                                                                                _ = v180
                                                                            case 0:
                                                                                v125 = (0.1 * v125)
                                                                            }
                                                                            v4 = (v4 + 1)
                                                                            _ = v173
                                                                            _ = v174
                                                                        }
                                                                    case 128:
                                                                        p1 = p1
                                                                    }
                                                                    _ = v159
                                                                } else {
                                                                    var v181 float32 = -2.25
                                                                    if true {
                                                                        v4 -= v4
                                                                    }
                                                                    var v182 int8 = v4
                                                                    _ = v181
                                                                    _ = v182
                                                                }
                                                            case "\n\t":
                                                                v4 -= (v4 << ((0) & 63))
                                                            case "foo bar":
                                                                var v183 uint64 = 661
                                                                _ = v183
                                                            }
                                                            v184 := GenS24{F0: 0, F1: 197}
                                                            _ = v184
                                                        } else {
                                                            var v185 float32 = float32((0.0 - v125))
                                                            v125 += v125
                                                            _ = v185
                                                        }
                                                        var v186 int16 = (p1 >> ((2) & 63))
                                                        _ = v186
                                                    }
                                                }
                                                p1 = (p1 >> ((147) & 63))
                                                if (true && true) {
                                                    if false {
                                                        var v187 int = fi129
                                                        v188 := []int{(v187 / ((2147483647) | 1))}
                                                        _ = v187
                                                        _ = v188
                                                    }
                                                    var v189 int = fi129
                                                    if false {
                                                        v190 := GenS6{F0: 945}
                                                        if (0.0 <= float32(v125)) {
                                                            fmt.Println(v124, p0, fi129)
                                                            if (42 <= 4294967295) {
                                                                if true {
                                                                    v191 := []int{(fi129 ^ v189), (-(v189))}
                                                                    _ = v191
                                                                } else {
                                                                    v125 = v125
                                                                }
                                                            }
                                                        }
                                                        _ = v190
                                                    }
                                                    p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
                                                    var v192 uint32 = (v124 | v124)
                                                    _ = v189
                                                    _ = v192
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v193 := []int{}
                                                    _ = v193
                                                case "foo bar":
                                                    v194 := []int{(fi55 | -9223372036854775808), int(v4), fi129}
                                                    _ = v194
                                                case "hello":
                                                    v124 += (v124 ^ v124)
                                                    v195 := []int{(fi55 / (((fi129 & -9223372036854775808)) | 1)), 127}
                                                    v125 -= 0.1
                                                    _ = v195
                                                default:
                                                    v196 := GenS14{F0: 42, F1: 32767}
                                                    _ = v196
                                                }
                                                _ = fi129
                                            }
                                        case "\n\t":
                                            for fi197 := 0; fi197 < 15; fi197++ {
                                                fmt.Println(v125, v125)
                                                var v198 int64 = p0
                                                switch "\n\t" {
                                                case "hello":
                                                    var v199 uint32 = v124
                                                    _ = v199
                                                default:
                                                    if false {
                                                        p0 += 0
                                                        v198 += p0
                                                        p1 -= p1
                                                    }
                                                }
                                                _ = fi197
                                                _ = v198
                                            }
                                        default:
                                            p1 = (p1 - p1)
                                        }
                                    default:
                                        if true {
                                            v200 := []int{}
                                            fmt.Println(v124, p0, p0)
                                            v201 := GenS21{F0: -1, F1: ("" + "\n\t"), F2: (v124 >> ((307) & 63))}
                                            _ = v200
                                            _ = v201
                                        } else {
                                            var v202 uint64 = 1
                                            if (1e10 == (3.14159 * (float32((v125 / 3.14159)) + 1e10))) {
                                                v124 = v124
                                                var v203 bool = false
                                                _ = v203
                                            }
                                            _ = v202
                                        }
                                    }
                                }
                                _ = v125
                            } else {
                                switch ("" + (("foo bar" + "") + ("" + "a"))) {
                                case "hello":
                                    var v204 int64 = p0
                                    _ = v204
                                case "foo bar":
                                    v205 := GenS25{F0: -860, F1: (p1 - p1)}
                                    _ = v205
                                }
                            }
                        }
                        v206 := GenS8{F0: (false || (!(true || (true || (!((fi55 >= fi55) && false))))))}
                        _ = v124
                        _ = v206
                    } else {
                        v207 := GenS26{F0: 255, F1: "a"}
                        _ = v207
                    }
                } else {
                    for fi208 := 0; fi208 < 18; fi208++ {
                        p1 = 1000
                        p0 += p0
                        _ = fi208
                    }
                    v209 := GenS5{F0: 1}
                    for fi210 := 0; fi210 < 4; fi210++ {
                        if true {
                            switch "" {
                            case "hello":
                                var v211 uint16 = 255
                                _ = v211
                            case "":
                                v212 := []int{(fi55 - fi210), fi55, fi210, (fi55 - -2147483648)}
                                fmt.Println(p1)
                                _ = len(v212)
                                if false {
                                    switch (fi55 & (fi55 / (((-(fi210))) | 1))) {
                                    case 4294967295:
                                        var v213 uint16 = 255
                                        _ = v213
                                    case 111:
                                        v4 = v4
                                    case 42:
                                        var v214 int8 = (v4 >> ((1) & 63))
                                        _ = v214
                                    default:
                                        fmt.Println(v4, fi210)
                                    }
                                } else {
                                    fmt.Println(fi55, fi55)
                                }
                                switch fi55 {
                                case -9223372036854775808:
                                    var v215 bool = (((1e10 * 1.5) <= 0.1) || (-9223372036854775808 != fi55))
                                    fmt.Println(p0, v4)
                                    _ = v215
                                case 32767:
                                    p1 -= (p1 % ((p1) | 1))
                                case 1:
                                    v216 := GenS6{F0: 255}
                                    _ = v216
                                }
                                v4 = (-(v4))
                                var v217 int16 = 1000
                                _ = v212
                                _ = v217
                            case "foo bar":
                                fmt.Println(fi210)
                                v218 := []int{(fi55 % ((fi210) | 1)), -736, fi55, (-(fi210))}
                                _ = v218
                            }
                            v219 := GenS27{F0: 0, F1: 127}
                            var v220 string = ""
                            _ = v219
                            _ = v220
                        } else {
                            if (!((-2.25 + 0.0) < (((1.5 * -2.25) - 3.14159) + 0.1))) {
                                fmt.Println(p0, p1, p0)
                            }
                        }
                        _ = fi210
                    }
                    _ = v209
                }
                var v221 uint64 = 1
                _ = v221
            }
            v222 := []int{fi55, 243, fi55, fi55}
            _ = v222
        }
        var v223 uint32 = 4294967295
        if true {
            var v224 float64 = 0.1
            _ = v224
        } else {
            if (true && (42 <= 42)) {
                v225 := GenS28{F0: 42, F1: 32767, F2: p0}
                fmt.Println(p0, fi55)
                _ = v225
            }
            for fi226 := 0; fi226 < 13; fi226++ {
                var v227 float32 = 0.1
                _ = fi226
                _ = v227
            }
            var v228 int = (fi55 >> ((2) & 63))
            _ = v228
        }
        _ = fi55
        _ = v223
    }
    _ = p0
    _ = p1
    _ = v4
    return true
}

func main() {
    var v1 string = "a"
    var v2 float64 = float64(1.5)
    v3 := make(map[uint32]uint32)
    var v229 bool = true
    switch 990 {
    case -32:
        switch -956 {
        case 0:
            v229 = v229
        default:
            v2 += -2.25
        }
    case 32768:
        var v230 int32 = 1
        _ = v230
    }
    var v231 uint8 = 255
    for fi232 := 0; fi232 < 2; fi232++ {
        _ = len(v3)
        v229 = (114 >= 113)
        v231 = v231
        _ = fi232
    }
    _ = len(v3)
    if (!((v229 && v229) || (-5 <= 42))) {
        if v229 {
            var v233 string = ""
            _ = v233
        } else {
            if true {
                for fi234 := 0; fi234 < 2; fi234++ {
                    v235 := []int{fi234}
                    _ = fi234
                    _ = v235
                }
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(v1, v2)
                }
            }
            if (435 >= 0) {
                fmt.Println(v1, v2, v231)
            }
            var v236 string = "a"
            _ = v236
        }
        switch int(v231) {
        case -287:
            v237 := v3[2]
            _ = v237
        }
    }
    var v238 uint16 = 255
    v3[2147483648] = 621
    var v239 float64 = v2
    v3[uint32(v238)] = 2147483647
    v240 := GenS13{F0: -1, F1: v1, F2: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v229:", v229)
    fmt.Println("v231:", v231)
    fmt.Println("v238:", v238)
    fmt.Println("v239:", v239)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v229
    _ = v231
    _ = v238
    _ = v239
    _ = v240
    _ = v3
}
