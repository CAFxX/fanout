package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 uint8; F2 bool }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 int32; F1 int32; F2 uint16 }
type GenS15 struct { F0 int16 }
func genFn0(p0 int64, p1 uint) int {
    if true {
        v6 := GenS1{F0: 702, F1: 875, F2: ("" + "hello")}
        fmt.Println(p0)
        _ = v6
    } else {
        var v7 int8 = -6
        p0 -= (^(p0))
        fmt.Println(v7)
        _ = v7
    }
    var v8 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 int) bool {
    p1 += (p1 % ((p1) | 1))
    if false {
        v9 := GenS2{F0: p1, F1: (p0 + (p0 % ((0) | 1))), F2: 127}
        p0 = (p0 ^ (-(p0)))
        _ = v9
    }
    v10 := []int{}
    p1 = 0
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint32 = 226
        v10 = append(v10, fi11)
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return true
}

func main() {
    var v1 int16 = -32768
    var v2 uint16 = 65535
    v3 := make(map[uint32]uint32)
    v4 := []uint32{1, 2, 290}
    var v5 float64 = 0.0
    if false {
        v2 = v2
        if (2 >= (v2 % (((v2 & (v2 << ((1) & 63)))) | 1))) {
            v1 -= (v1 ^ -746)
        }
    } else {
        fmt.Println(v1, v5)
        if (genFn0(-9223372036854775808, 18446744073709551615) == 9223372036854775807) {
            var v13 uint32 = 371
            _ = v13
        } else {
            var v14 uint16 = v2
            _ = v14
        }
        fmt.Println(v1, v2)
    }
    if (32767 < 127) {
        if (!true) {
            _ = len(v3)
        }
        var v15 uint16 = v2
        v16 := GenS3{F0: ("hello" + "hello"), F1: 32767}
        _ = v15
        _ = v16
    } else {
        var v17 uint16 = (^(v2))
        if ((831 < 349) && (!(42 <= 0))) {
            _ = len(v3)
            fmt.Println(v2, v17, v17)
            if ((0.1 + 1e10) < (float32((v5 - (0.0 - 0.1))) * 1e10)) {
                var v18 uint64 = 18446744073709551615
                for fi19 := 0; fi19 < 16; fi19++ {
                    var v20 int32 = 42
                    for fi21 := 0; fi21 < 14; fi21++ {
                        v22 := GenS1{F0: v18, F1: (v1 ^ int16(v20)), F2: ("foo bar" + "hello")}
                        switch (fi21 - genFn0(9223372036854775807, 394)) {
                        case 9223372036854775807:
                            var v23 int64 = -9223372036854775808
                            _ = v23
                        }
                        _ = fi21
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            } else {
                var v24 uint64 = 0
                v2 = 2
                _ = v24
            }
        }
        switch "" {
        case "foo bar":
            if ((1.5 + float32(v5)) != 3.14159) {
                var v25 float64 = v5
                _ = v25
            } else {
                switch "" {
                case "\n\t":
                    for fi26 := 0; fi26 < 20; fi26++ {
                        v27 := GenS4{F0: 9223372036854775807, F1: 0}
                        fmt.Println(v17, v1, v1)
                        fmt.Println(fi26)
                        _ = fi26
                        _ = v27
                    }
                case "hello":
                    v2 = (^(v2))
                case "":
                    var v28 int64 = 0
                    _ = v28
                default:
                    switch "a" {
                    case "":
                        var v29 uint8 = 2
                        _ = v29
                    case "foo bar":
                        var v30 int64 = -1
                        _ = v30
                    }
                }
                switch "a" {
                case "a":
                    var v31 uint8 = 255
                    _ = v31
                default:
                    for fi32 := 0; fi32 < 14; fi32++ {
                        v33 := GenS5{F0: 0, F1: 151, F2: 255}
                        for fi34 := 0; fi34 < 7; fi34++ {
                            v4 = append(v4, 2147483648)
                            if (1e10 <= 1e10) {
                                if len(v4) > 0 {
                                    v4[0] = 1
                                }
                            }
                            switch (fi34 * -26) {
                            case 1:
                                fmt.Println(v1)
                            default:
                                var v35 uint32 = 2
                                _ = v35
                            }
                            _ = fi34
                        }
                        switch (fi32 - (-(fi32))) {
                        case 127:
                            v5 = (v5 * (v5 * float64(((3.14159 * ((-2.25 / 1e10) + 0.1)) * 0.0))))
                        case 32767:
                            for fi36 := 0; fi36 < 4; fi36++ {
                                var v37 int64 = 9223372036854775807
                                var v38 string = "a"
                                v39 := GenS1{F0: 1, F1: v1, F2: ("a" + (v38 + (((v38 + v38) + v38) + v38)))}
                                var v40 uint64 = 2
                                _ = fi36
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                        case 370:
                            v5 = float64((0.0 - float32(1.5)))
                        default:
                            for fi41 := 0; fi41 < 11; fi41++ {
                                v1 = 0
                                v17 -= (v2 * v17)
                                _ = fi41
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                }
                switch "foo bar" {
                case "a":
                    v42 := v3[65535]
                    _ = v42
                case "\n\t":
                    v43 := GenS5{F0: 2, F1: 2, F2: 1}
                    _ = v43
                case "foo bar":
                    for fi44 := 0; fi44 < 15; fi44++ {
                        v45 := GenS6{F0: 42}
                        for fi46 := 0; fi46 < 16; fi46++ {
                            if false {
                                v47 := v3[0]
                                var v48 int64 = int64(v17)
                                _ = v47
                                _ = v48
                            }
                            switch int(v1) {
                            case 42:
                                if ((-128 > 0) && true) {
                                    v17 = (v2 << ((42) & 63))
                                }
                            default:
                                v2 = v2
                            }
                            var v49 string = ("hello" + ("hello" + "hello"))
                            _ = fi46
                            _ = v49
                        }
                        if (!(!false)) {
                            if true {
                                if true {
                                    var v50 uint32 = 2147483648
                                    _ = v50
                                } else {
                                    switch "hello" {
                                    case "\n\t":
                                        v2 += (^(v17))
                                    case "hello":
                                        v5 = v5
                                    case "foo bar":
                                        if false {
                                            for fi51 := 0; fi51 < 6; fi51++ {
                                                var v52 bool = (v5 >= (v5 * v5))
                                                switch (fi51 & (fi51 + fi51)) {
                                                case 750:
                                                    if v52 {
                                                        var v53 string = "\n\t"
                                                        _ = v53
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v17, v17, v52)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    var v54 uint = 42
                                                    _ = v54
                                                }
                                                _ = fi51
                                                _ = v52
                                            }
                                        } else {
                                            var v55 uint8 = 41
                                            for fi56 := 0; fi56 < 11; fi56++ {
                                                switch "foo bar" {
                                                case "a":
                                                    var v57 uint = 313
                                                    _ = v57
                                                case "":
                                                    if ((v1 >> ((567) & 63)) > 1) {
                                                        v1 += v1
                                                    }
                                                case "\n\t":
                                                    v5 = (float64(0.0) + v5)
                                                    if true {
                                                        _ = len(v4)
                                                        if len(v4) > 0 {
                                                            v4[0] = 2147483648
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v1, fi44)
                                                }
                                                switch (("" + "foo bar") + "a") {
                                                case "":
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v58 bool = true
                                                        _ = v58
                                                    case "\n\t":
                                                        _ = len(v4)
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi56, v5)
                                                case "hello":
                                                    if true {
                                                        var v59 string = "\n\t"
                                                        v2 = (v17 * v17)
                                                        switch 128 {
                                                        case 13:
                                                            var v60 float64 = -2.25
                                                            _ = v60
                                                        default:
                                                            if false {
                                                                if (!((v1 << ((1) & 63)) != (-(v1)))) {
                                                                    var v61 uint32 = 2147483648
                                                                    _ = v61
                                                                } else {
                                                                    v3[1] = 948
                                                                    v5 = v5
                                                                }
                                                                v59 = v59
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        switch ("" + ("hello" + v59)) {
                                                                        case "\n\t":
                                                                            var v62 uint32 = 2147483647
                                                                            _ = v62
                                                                        case "":
                                                                            v55 -= 128
                                                                        case "hello":
                                                                            v63 := GenS5{F0: uint(v55), F1: 0, F2: 42}
                                                                            _ = v63
                                                                        default:
                                                                            if false {
                                                                                v59 = v59
                                                                            } else {
                                                                                var v64 uint = 4294967295
                                                                                v65 := GenS6{F0: 613}
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                        }
                                                                        _ = len(v3)
                                                                    }
                                                                    fmt.Println(fi56)
                                                                    if (v1 < v1) {
                                                                        v66 := GenS7{F0: 65535}
                                                                        switch (-(fi44)) {
                                                                        case 341:
                                                                            _ = len(v3)
                                                                        case 9223372036854775807:
                                                                            v67 := GenS8{F0: 179, F1: 4294967295}
                                                                            _ = v67
                                                                        case 32767:
                                                                            var v68 uint8 = 176
                                                                            _ = v68
                                                                        default:
                                                                            fmt.Println(v2, v2, fi56)
                                                                        }
                                                                        _ = v66
                                                                    } else {
                                                                        switch v59 {
                                                                        case "foo bar":
                                                                            var v69 int16 = v1
                                                                            _ = v69
                                                                        }
                                                                        var v70 bool = (int64(v2) < -692)
                                                                        v17 = v17
                                                                        _ = v70
                                                                    }
                                                                    fmt.Println(v55, v55, v2)
                                                                } else {
                                                                    if (0 != 740) {
                                                                        v5 = (v5 / float64((-2.25 / float32((v5 * v5)))))
                                                                        v71 := GenS2{F0: (-(fi56)), F1: 42, F2: 76}
                                                                        var v72 int = fi56
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                    switch genFn0(0, 18446744073709551615) {
                                                                    case 32767:
                                                                        v17 -= v2
                                                                    default:
                                                                        var v73 bool = true
                                                                        _ = v73
                                                                    }
                                                                    fmt.Println(v5, v1)
                                                                }
                                                                var v74 bool = true
                                                                v75 := GenS9{F0: -92}
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                        }
                                                        _ = v59
                                                    } else {
                                                        v17 -= (^(v2))
                                                        var v76 float64 = v5
                                                        v77 := GenS5{F0: 592, F1: 174, F2: v55}
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                }
                                                fmt.Println(v17, v1, fi56)
                                                var v78 uint16 = 32767
                                                _ = fi56
                                                _ = v78
                                            }
                                            for fi79 := 0; fi79 < 2; fi79++ {
                                                v1 += v1
                                                _ = fi79
                                            }
                                            _ = v55
                                        }
                                    }
                                    switch "" {
                                    case "":
                                        var v80 int = (fi44 << ((4294967295) & 63))
                                        _ = v80
                                    case "a":
                                        v17 = (v17 >> ((0) & 63))
                                    case "\n\t":
                                        v5 += v5
                                    default:
                                        v2 -= 296
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = 774
                                    }
                                }
                                var v81 int8 = int8(fi44)
                                _ = v81
                            } else {
                                var v82 float64 = 0.0
                                v2 = (v2 >> ((255) & 63))
                                _ = v82
                            }
                        }
                        _ = fi44
                        _ = v45
                    }
                }
            }
        case "a":
            switch "" {
            case "":
                switch "" {
                case "":
                    for fi83 := 0; fi83 < 9; fi83++ {
                        switch -1 {
                        case 32768:
                            v84 := GenS3{F0: "hello", F1: (v1 % ((v1) | 1))}
                            _ = v84
                        case 32767:
                            if (!(false || (((-(fi83)) >= fi83) && true))) {
                                var v85 bool = true
                                fmt.Println(v2)
                                _ = v85
                            } else {
                                v4 = append(v4, 585)
                                v86 := v3[85]
                                _ = v86
                            }
                        case -2147483648:
                            fmt.Println(v2, v5, v1)
                        }
                        var v87 int16 = -207
                        _ = fi83
                        _ = v87
                    }
                }
            case "foo bar":
                v2 -= (v2 << ((42) & 63))
            case "hello":
                var v88 float64 = (3.14159 * 1.5)
                _ = v88
            default:
                v17 = (v17 & (v17 & v17))
            }
        case "\n\t":
            var v89 bool = false
            _ = v89
        default:
            for fi90 := 0; fi90 < 1; fi90++ {
                fmt.Println(v1, v1, fi90)
                v91 := GenS6{F0: 2}
                _ = fi90
                _ = v91
            }
        }
        _ = v17
    }
    var v92 int8 = 42
    if false {
        switch 32767 {
        case 127:
            var v93 uint32 = 1
            _ = v93
        default:
            for fi94 := 0; fi94 < 8; fi94++ {
                for fi95 := 0; fi95 < 8; fi95++ {
                    var v96 uint16 = 1
                    _ = fi95
                    _ = v96
                }
                var v97 uint8 = 0
                _ = fi94
                _ = v97
            }
        }
        v98 := v3[0]
        _ = v98
    } else {
        fmt.Println(v2, v92)
        v4 = append(v4, 42)
    }
    v99 := GenS10{F0: 2147483647, F1: 181, F2: true}
    if (!false) {
        if (!true) {
            var v100 float32 = -2.25
            var v101 int64 = -9223372036854775808
            if false {
                for fi102 := 0; fi102 < 1; fi102++ {
                    switch genFn0((v101 * (v101 >> ((359) & 63))), 1) {
                    case -135:
                        if true {
                            v2 -= (v2 & 255)
                        } else {
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                            if len(v4) > 0 {
                                v4[0] = 54
                            }
                        }
                    default:
                        fmt.Println(fi102)
                    }
                    var v103 uint8 = 255
                    if false {
                        v3[4294967295] = 227
                        if ((-(fi102)) != fi102) {
                            var v104 uint16 = v2
                            _ = v104
                        }
                        var v105 float64 = ((3.14159 / 1.0) - v5)
                        _ = v105
                    }
                    _ = fi102
                    _ = v103
                }
                _ = len(v4)
            }
            _ = v100
            _ = v101
        }
        fmt.Println(v92, v92, v2)
    }
    switch "" {
    case "":
        for fi106 := 0; fi106 < 4; fi106++ {
            v107 := GenS11{F0: v1, F1: 42}
            _ = len(v4)
            v108 := GenS7{F0: 0}
            _ = fi106
            _ = v107
            _ = v108
        }
    case "a":
        if false {
            v92 = (v92 / (((v92 << ((4294967295) & 63))) | 1))
            if ((-2.25 + 3.14159) <= 0.1) {
                v109 := v3[728]
                if true {
                    if false {
                        v5 = ((-2.25 - 0.1) - (3.14159 - -2.25))
                        if false {
                            v4 = append(v4, 112)
                            var v110 int8 = (v92 - v92)
                            _ = v110
                        } else {
                            if (((1000 == v1) || (false && true)) && true) {
                                if true {
                                    if true {
                                        var v111 uint32 = 911
                                        var v112 int32 = 563
                                        var v113 int8 = (v92 ^ v92)
                                        _ = v111
                                        _ = v112
                                        _ = v113
                                    }
                                }
                                for fi114 := 0; fi114 < 1; fi114++ {
                                    if (true && true) {
                                        if (true && true) {
                                            for fi115 := 0; fi115 < 4; fi115++ {
                                                v92 = (^(v92))
                                                if true {
                                                    if (!false) {
                                                        v116 := GenS4{F0: -9223372036854775808, F1: 1}
                                                        v5 += (v5 * v5)
                                                        _ = v116
                                                    } else {
                                                        var v117 uint16 = (v2 / ((v2) | 1))
                                                        _ = v117
                                                    }
                                                } else {
                                                    fmt.Println(fi114, fi114)
                                                }
                                                _ = len(v3)
                                                _ = fi115
                                            }
                                        } else {
                                            for fi118 := 0; fi118 < 7; fi118++ {
                                                switch (fi114 - 32768) {
                                                case -1:
                                                    fmt.Println(v5)
                                                }
                                                var v119 float64 = v5
                                                if true {
                                                    v120 := GenS4{F0: 0, F1: v92}
                                                    var v121 float64 = (v5 / float64((3.14159 * -2.25)))
                                                    _ = v120
                                                    _ = v121
                                                } else {
                                                    switch fi114 {
                                                    case 32767:
                                                        v122 := GenS12{F0: (v109 ^ 905), F1: 1, F2: 42}
                                                        _ = v122
                                                    case 127:
                                                        var v123 float32 = 3.14159
                                                        _ = v123
                                                    case 439:
                                                        v3[(v109 >> ((1) & 63))] = (v109 + (^(v109)))
                                                    default:
                                                        var v124 uint = 255
                                                        _ = v124
                                                    }
                                                }
                                                _ = fi118
                                                _ = v119
                                            }
                                        }
                                    } else {
                                        _ = len(v3)
                                        v125 := GenS8{F0: -32768, F1: 4294967295}
                                        _ = v125
                                    }
                                    var v126 uint16 = v2
                                    v127 := GenS1{F0: 537, F1: (v1 | (v1 | v1)), F2: "hello"}
                                    _ = fi114
                                    _ = v126
                                    _ = v127
                                }
                            }
                            for fi128 := 0; fi128 < 12; fi128++ {
                                v3[4294967295] = (v109 * (v109 * 1))
                                for fi129 := 0; fi129 < 16; fi129++ {
                                    v130 := GenS2{F0: -84, F1: v1, F2: (v92 ^ (v92 << ((42) & 63)))}
                                    var v131 bool = (v92 >= (v92 << ((2) & 63)))
                                    _ = fi129
                                    _ = v130
                                    _ = v131
                                }
                                _ = fi128
                            }
                        }
                        v132 := GenS8{F0: v1, F1: 2}
                        _ = v132
                    } else {
                        _ = len(v4)
                        v92 = (v92 << ((4294967295) & 63))
                    }
                }
                for fi133 := 0; fi133 < 15; fi133++ {
                    v134 := GenS3{F0: ("foo bar" + "a"), F1: v1}
                    delete(v3, (v109 << ((156) & 63)))
                    if (!true) {
                        for fi135 := 0; fi135 < 7; fi135++ {
                            var v136 uint8 = 127
                            var v137 int8 = v92
                            v2 += (^(v2))
                            _ = fi135
                            _ = v136
                            _ = v137
                        }
                    }
                    _ = fi133
                    _ = v134
                }
                _ = v109
            } else {
                fmt.Println(v5, v92)
            }
            var v138 float64 = -2.25
            _ = v138
        }
    case "\n\t":
        fmt.Println(v2)
    default:
        v5 -= v5
    }
    for fi139 := 0; fi139 < 10; fi139++ {
        switch ("" + (("" + ("\n\t" + ("a" + "\n\t"))) + "hello")) {
        case "hello":
            var v140 int64 = -1
            _ = v140
        case "":
            for fi141 := 0; fi141 < 16; fi141++ {
                var v142 uint8 = 0
                if ((3.14159 + 1e10) != -2.25) {
                    fmt.Println(v1, fi139, fi141)
                }
                _ = fi141
                _ = v142
            }
        case "\n\t":
            var v143 int64 = -1
            _ = v143
        }
        _ = fi139
    }
    if true {
        switch "a" {
        case "":
            v5 = float64(float32(v5))
        }
        v5 -= v5
        fmt.Println(v1, v92, v2)
    }
    var v144 int64 = 42
    if false {
        v2 = v2
        if ((((1e10 * -2.25) - 1.5) <= ((1.5 + 1e10) * 3.14159)) || true) {
            if (326 != uint32(v2)) {
                fmt.Println(v144)
                v2 += v2
                switch 127 {
                case 626:
                    if (3.14159 >= (3.14159 * v5)) {
                        v1 = (v1 / ((v1) | 1))
                        for fi145 := 0; fi145 < 18; fi145++ {
                            v2 += 42
                            v146 := GenS13{F0: 2, F1: ("\n\t" > ("" + "a"))}
                            _ = fi145
                            _ = v146
                        }
                        v92 += v92
                    } else {
                        v92 = 127
                        if len(v4) > 0 {
                            v4[0] = 395
                        }
                        v92 = (^(v92))
                    }
                case 128:
                    if true {
                        var v147 uint8 = 108
                        var v148 int8 = v92
                        _ = v147
                        _ = v148
                    }
                }
            } else {
                v1 -= v1
                v92 += (v92 + -115)
            }
            v4 = append(v4, 0)
            for fi149 := 0; fi149 < 2; fi149++ {
                for fi150 := 0; fi150 < 18; fi150++ {
                    v2 += v2
                    v151 := GenS14{F0: 0, F1: 47, F2: 0}
                    _ = fi150
                    _ = v151
                }
                v152 := v3[2]
                var v153 int8 = (v92 << ((583) & 63))
                _ = fi149
                _ = v152
                _ = v153
            }
        }
    } else {
        for fi154 := 0; fi154 < 20; fi154++ {
            for fi155 := 0; fi155 < 18; fi155++ {
                switch fi154 {
                case 9223372036854775807:
                    v156 := GenS15{F0: int16(v2)}
                    _ = v156
                case -1:
                    v144 -= v144
                case 128:
                    v92 = -1
                default:
                    var v157 int = genFn0(1, 0)
                    if (true || (!(false || false))) {
                        var v158 int8 = v92
                        var v159 string = ""
                        var v160 float32 = (-2.25 - (float32(0.0) + 3.14159))
                        _ = v158
                        _ = v159
                        _ = v160
                    }
                    v2 += (v2 & v2)
                    _ = v157
                }
                v144 = v144
                var v161 int32 = 42
                _ = fi155
                _ = v161
            }
            _ = fi154
        }
        var v162 string = "\n\t"
        for fi163 := 0; fi163 < 14; fi163++ {
            for fi164 := 0; fi164 < 16; fi164++ {
                var v165 uint32 = 59
                _ = fi164
                _ = v165
            }
            _ = fi163
        }
        _ = v162
    }
    var v166 int = genFn0(9223372036854775807, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v144:", v144)
    fmt.Println("v166:", v166)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v144
    _ = v166
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
    _ = v99
}
