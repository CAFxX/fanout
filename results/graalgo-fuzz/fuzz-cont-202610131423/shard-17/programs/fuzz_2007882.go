package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 string; F1 uint; F2 bool }
type GenS3 struct { F0 uint8; F1 bool; F2 int16 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 int16; F1 string; F2 string }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 uint32; F1 int16 }
type GenS14 struct { F0 bool; F1 int32 }
type GenS15 struct { F0 uint8; F1 uint8 }
type GenS16 struct { F0 uint }
type GenS17 struct { F0 bool }
func genFn0(p0 uint16, p1 int32, p2 bool) int {
    if ((478 <= 761) && p2) {
        var v7 int = 4294967295
        _ = v7
    }
    v8 := GenS1{F0: 2, F1: uint(p0)}
    switch "foo bar" {
    case "\n\t":
        switch "\n\t" {
        case "hello":
            p2 = p2
        case "foo bar":
            var v9 string = ("\n\t" + (("hello" + ((("hello" + "\n\t") + "foo bar") + "foo bar")) + "a"))
            _ = v9
        case "a":
            fmt.Println(p2)
        }
    case "":
        if p2 {
            fmt.Println(p0)
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := GenS2{F0: "hello", F1: 198, F2: p2}
                _ = fi10
                _ = v11
            }
        } else {
            var v12 int = 42
            var v13 int8 = 1
            switch "" {
            case "a":
                fmt.Println(p0)
            }
            _ = v12
            _ = v13
        }
    case "hello":
        var v14 uint32 = 4294967295
        _ = v14
    default:
        var v15 float32 = (((float32(1e10) - 0.0) * ((1.5 / float32(((1e10 - float64((0.1 * -2.25))) / 3.14159))) + 3.14159)) / 1e10)
        _ = v15
    }
    p0 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 226
}

func main() {
    var v1 uint16 = 598
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    v4 := []string{"hello", "foo bar"}
    var v5 int64 = -117
    v6 := make(map[uint32]uint32)
    if true {
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    } else {
        v2[42] = false
        delete(v2, 32768)
    }
    v16 := GenS2{F0: "\n\t", F1: 18446744073709551615, F2: true}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v17 uint16 = 42
    switch 2147483647 {
    case 132:
        fmt.Println(v1, v3)
    case 4294967295:
        var v18 int8 = 1
        _ = v18
    case -9223372036854775808:
        for fi19 := 0; fi19 < 1; fi19++ {
            v3 = -2.25
            fmt.Println(v1, v1)
            v20 := GenS3{F0: 1, F1: (true || ((96 > int8(v5)) && (402 >= 1))), F2: -487}
            _ = fi19
            _ = v20
        }
    }
    v21 := GenS2{F0: (("a" + ("a" + ("hello" + "hello"))) + ("foo bar" + "")), F1: 155, F2: (!((42 >= 0) || true))}
    switch -525 {
    case -1:
        for fi22 := 0; fi22 < 8; fi22++ {
            if false {
                switch ("foo bar" + ("" + "a")) {
                case "hello":
                    for fi23 := 0; fi23 < 8; fi23++ {
                        fmt.Println(fi22, v1, v17)
                        _ = fi23
                    }
                }
                switch genFn0((^(v1)), 0, false) {
                case 2147483647:
                    if true {
                        delete(v2, fi22)
                        for fi24 := 0; fi24 < 6; fi24++ {
                            if (((true && (343 >= 42)) && true) || (2 != 4294967295)) {
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = "hello"
                                    }
                                    switch 2147483647 {
                                    case 2147483647:
                                        v25 := GenS4{F0: 65535, F1: ""}
                                        _ = v25
                                    case 0:
                                        v3 = v3
                                        var v26 uint = 4294967295
                                        var v27 int64 = (v5 >> (((v26 % (((v26 << ((2) & 63))) | 1))) & 63))
                                        v3 = v3
                                        v1 = v17
                                        _ = v26
                                        _ = v27
                                    default:
                                        var v28 float64 = v3
                                        _ = v28
                                    }
                                } else {
                                    if true {
                                        fmt.Println(fi24)
                                    }
                                    var v29 int8 = -1
                                    v17 -= v1
                                    _ = v29
                                }
                                v30 := v6[1]
                                fmt.Println(v5, fi22, v1)
                                if false {
                                    v3 -= ((v3 - (v3 - float64(float32(0.0)))) - (((0.1 * float64((3.14159 + 0.1))) / 3.14159) / ((0.0 * v3) + (1.5 - v3))))
                                    v1 = 32767
                                    fmt.Println(v5, v30)
                                } else {
                                    fmt.Println(v30, fi22, fi24)
                                }
                                _ = v30
                            } else {
                                switch ((("\n\t" + "") + "hello") + "") {
                                case "a":
                                    var v31 uint64 = 238
                                    _ = v31
                                case "\n\t":
                                    if true {
                                        fmt.Println(fi22, v3, v3)
                                    }
                                case "hello":
                                    v3 = v3
                                default:
                                    var v32 uint16 = v17
                                    _ = v32
                                }
                                v33 := GenS1{F0: 254, F1: 255}
                                fmt.Println(fi22, v5, v17)
                                v17 += v1
                                _ = v33
                            }
                            _ = fi24
                        }
                    } else {
                        var v34 string = ("a" + "\n\t")
                        _ = v34
                    }
                case -9223372036854775808:
                    v35 := GenS5{F0: 11, F1: ("a" + ("foo bar" + ("foo bar" + ("\n\t" + "a")))), F2: ""}
                    _ = v35
                default:
                    for fi36 := 0; fi36 < 14; fi36++ {
                        if (v3 != float64(0.1)) {
                            v5 += v5
                        }
                        var v37 int16 = 32767
                        _ = fi36
                        _ = v37
                    }
                }
            }
            for fi38 := 0; fi38 < 11; fi38++ {
                switch fi22 {
                case 128:
                    var v39 uint = 768
                    _ = v39
                }
                v40 := GenS3{F0: 127, F1: (-1 == genFn0((v17 * v1), -1, (!(!true)))), F2: -654}
                _ = fi38
                _ = v40
            }
            _ = fi22
        }
    case 128:
        if (genFn0((v1 + (v1 | 0)), 226, ((4294967295 == 1) && true)) > 127) {
            v41 := v2[-883]
            fmt.Println(v1)
            v42 := GenS6{F0: -52, F1: 127}
            _ = v41
            _ = v42
        } else {
            if (111 > 127) {
                if true {
                    if false {
                        var v43 uint = 255
                        _ = len(v4)
                        var v44 int32 = -2147483648
                        _ = v43
                        _ = v44
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v3, v17, v3)
                        v17 += (v1 + 323)
                    }
                    v3 = v3
                } else {
                    for fi45 := 0; fi45 < 11; fi45++ {
                        if (v5 > (v5 & v5)) {
                            if false {
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    v47 := GenS7{F0: false, F1: 127}
                                    switch genFn0((^(v17)), -2147483648, (112 >= (v17 & v17))) {
                                    case 9223372036854775807:
                                        var v48 uint8 = 42
                                        _ = v48
                                    case 128:
                                        var v49 uint16 = (v1 / ((32767) | 1))
                                        _ = v49
                                    }
                                    _ = fi46
                                    _ = v47
                                }
                            } else {
                                if true {
                                    if (0.1 != 3.14159) {
                                        switch 128 {
                                        case 813:
                                            var v50 uint16 = v1
                                            _ = v50
                                        case 32767:
                                            var v51 uint16 = (v17 >> ((18446744073709551615) & 63))
                                            _ = v51
                                        }
                                        for fi52 := 0; fi52 < 13; fi52++ {
                                            if true {
                                                var v53 string = ""
                                                v1 = v1
                                                _ = v53
                                            } else {
                                                _ = len(v4)
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = "foo bar"
                                                    }
                                                } else {
                                                    fmt.Println(v1, v5, v3)
                                                    _ = len(v4)
                                                }
                                                var v54 float64 = ((-2.25 + (((-2.25 - v3) + v3) - (v3 / -2.25))) + v3)
                                                if ((2 <= 65535) || false) {
                                                    var v55 uint32 = 2147483647
                                                    _ = v55
                                                }
                                                v3 += (((v3 - v3) + v54) / -2.25)
                                                _ = v54
                                            }
                                            var v56 int32 = -667
                                            _ = fi52
                                            _ = v56
                                        }
                                    } else {
                                        for fi57 := 0; fi57 < 14; fi57++ {
                                            var v58 float32 = (-2.25 - float32(float64(0.1)))
                                            _ = fi57
                                            _ = v58
                                        }
                                        fmt.Println(fi45)
                                    }
                                    for fi59 := 0; fi59 < 1; fi59++ {
                                        switch "foo bar" {
                                        case "a":
                                            if ((true || true) && false) {
                                                if (1e10 <= 1.5) {
                                                    if false {
                                                        var v60 int8 = 42
                                                        _ = v60
                                                    }
                                                    fmt.Println(fi59)
                                                    v61 := GenS7{F0: (false || (true && true)), F1: 183}
                                                    if (380 > 0) {
                                                        v5 = v5
                                                        v5 -= (v5 + (v5 % (((v5 << ((4294967295) & 63))) | 1)))
                                                        if ((933 > 2) && ((1 < 6) && false)) {
                                                            switch "hello" {
                                                            case "":
                                                                var v62 int8 = -19
                                                                _ = v62
                                                            default:
                                                                var v63 int8 = 100
                                                                _ = v63
                                                            }
                                                            v3 = ((float64((0.0 / 1.5)) - float64(-2.25)) / float64((1e10 - (-2.25 * -2.25))))
                                                            v3 = (1.5 * v3)
                                                            fmt.Println(v3, v17)
                                                            var v64 string = ""
                                                            _ = v64
                                                        } else {
                                                            if true {
                                                                v1 = (^(v1))
                                                            }
                                                            switch ("foo bar" + "hello") {
                                                            case "foo bar":
                                                                var v65 int = fi45
                                                                _ = v65
                                                            case "hello":
                                                                var v66 uint32 = 4294967295
                                                                _ = v66
                                                            case "a":
                                                                switch fi45 {
                                                                case 127:
                                                                    fmt.Println(v5, fi59, fi45)
                                                                case -187:
                                                                    if (120 == 0) {
                                                                        v5 = 9223372036854775807
                                                                    }
                                                                }
                                                            default:
                                                                v67 := v2[(fi45 >> ((18446744073709551615) & 63))]
                                                                _ = v67
                                                            }
                                                            v68 := v6[2]
                                                            _ = v68
                                                        }
                                                    } else {
                                                        v3 = 0.0
                                                        v17 += v17
                                                    }
                                                    _ = v61
                                                }
                                                if false {
                                                    v69 := GenS7{F0: true, F1: 211}
                                                    _ = v69
                                                } else {
                                                    if (964 != 1) {
                                                        v3 = float64(1e10)
                                                    }
                                                    fmt.Println(v17, fi59, v5)
                                                    v70 := v6[2147483647]
                                                    fmt.Println(v5, v5)
                                                    var v71 int = (fi59 | (fi45 + 897))
                                                    _ = v70
                                                    _ = v71
                                                }
                                                if (((-2.25 / 0.1) <= 3.14159) || true) {
                                                    v72 := GenS8{F0: 42}
                                                    _ = v72
                                                } else {
                                                    var v73 uint = 0
                                                    _ = v73
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = (("hello" + ("foo bar" + "a")) + "\n\t")
                                                }
                                            }
                                        case "":
                                            var v74 bool = false
                                            _ = v74
                                        }
                                        var v75 uint32 = 559
                                        _ = fi59
                                        _ = v75
                                    }
                                }
                                if false {
                                    if (42 < (v5 | 42)) {
                                        var v76 uint8 = 173
                                        for fi77 := 0; fi77 < 16; fi77++ {
                                            fmt.Println(fi77)
                                            switch 4294967295 {
                                            case -9223372036854775808:
                                                fmt.Println(fi45, v3, v17)
                                            case 32767:
                                                switch (fi77 % ((genFn0(0, 42, true)) | 1)) {
                                                case 4294967295:
                                                    if false {
                                                        v2[(fi77 ^ 541)] = (0 != 42)
                                                    }
                                                default:
                                                    fmt.Println(v1, v3)
                                                }
                                            case 2147483647:
                                                switch -738 {
                                                case -2147483648:
                                                    switch "foo bar" {
                                                    case "a":
                                                        v78 := v2[fi45]
                                                        _ = v78
                                                    case "":
                                                        switch ("foo bar" + (("foo bar" + "\n\t") + "\n\t")) {
                                                        case "hello":
                                                            _ = len(v4)
                                                        case "\n\t":
                                                            if false {
                                                                fmt.Println(fi77, v1, fi77)
                                                            }
                                                        }
                                                    case "hello":
                                                        v5 += 1
                                                    default:
                                                        if true {
                                                            if len(v4) > 0 {
                                                                v4[0] = "hello"
                                                            }
                                                            switch fi77 {
                                                            case -549:
                                                                v76 -= v76
                                                            default:
                                                                var v79 int64 = v5
                                                                v6[4294967295] = 2
                                                                var v80 float64 = v3
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                            var v81 uint16 = uint16(v76)
                                                            _ = v81
                                                        }
                                                    }
                                                    fmt.Println(v3, v76)
                                                case 4294967295:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v5, fi77, v3)
                                                    case "":
                                                        if (((-2.25 * v3) - -2.25) < (-2.25 * float64(-2.25))) {
                                                            delete(v2, fi77)
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v82 int32 = 42
                                                            var v83 uint = 0
                                                            var v84 uint64 = 42
                                                            var v85 bool = false
                                                            _ = v82
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            var v86 int16 = 0
                                                            _ = v86
                                                        }
                                                    case "hello":
                                                        switch fi45 {
                                                        case -9223372036854775808:
                                                            var v87 uint8 = v76
                                                            _ = v87
                                                        default:
                                                            v76 = 127
                                                        }
                                                    default:
                                                        v1 -= (v17 / ((0) | 1))
                                                    }
                                                case -117:
                                                    var v88 int = -1
                                                    _ = v88
                                                default:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            default:
                                                var v89 uint64 = 42
                                                v89 += v89
                                                _ = v89
                                            }
                                            _ = fi77
                                        }
                                        v76 += uint8(v1)
                                        _ = v76
                                    } else {
                                        if (260 >= -2147483648) {
                                            for fi90 := 0; fi90 < 13; fi90++ {
                                                var v91 uint32 = 162
                                                switch "a" {
                                                case "hello":
                                                    v1 += 2
                                                case "\n\t":
                                                    v92 := GenS4{F0: 255, F1: "a"}
                                                    _ = v92
                                                case "":
                                                    var v93 uint32 = v91
                                                    _ = v93
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = "foo bar"
                                                }
                                                _ = fi90
                                                _ = v91
                                            }
                                            v6[564] = 4294967295
                                        } else {
                                            var v94 float32 = (1e10 - (0.0 / (0.1 / (float32(v3) + (float32(float64(0.1)) - float32((v3 / v3)))))))
                                            v95 := GenS3{F0: 211, F1: false, F2: -1}
                                            var v96 uint = 7
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                        }
                                        switch (^(fi45)) {
                                        case 128:
                                            v5 = (v5 - 641)
                                        case -2147483648:
                                            var v97 int64 = 909
                                            _ = v97
                                        case -1:
                                            v1 += 942
                                        }
                                        var v98 int32 = -1
                                        _ = v98
                                    }
                                    var v99 int8 = -91
                                    _ = len(v4)
                                    _ = v99
                                }
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v3 += (((v3 - float64(0.1)) * v3) * (3.14159 + 1e10))
                        _ = fi45
                    }
                    v6[290] = 725
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            }
            var v100 int = -18
            _ = v100
        }
    default:
        switch 9223372036854775807 {
        case 834:
            v5 += (v5 % ((v5) | 1))
        }
    }
    switch genFn0(v17, -1, (!(genFn0(v17, -2147483648, (42 == 18446744073709551615)) > genFn0(v1, 2147483647, true)))) {
    case -40:
        if (!((4294967295 <= -627) || ((!false) || false))) {
            v5 = 1
            v101 := GenS9{F0: 100}
            for fi102 := 0; fi102 < 12; fi102++ {
                switch "foo bar" {
                case "\n\t":
                    var v103 uint = 351
                    _ = v103
                case "a":
                    v17 -= 65535
                }
                for fi104 := 0; fi104 < 4; fi104++ {
                    fmt.Println(fi104, v5)
                    var v105 uint32 = 1
                    _ = fi104
                    _ = v105
                }
                _ = fi102
            }
            _ = v101
        }
    case 32767:
        if (true || ((true && (((v1 >= v17) && (!(1000 >= 404))) && ((2147483647 >= -9223372036854775808) && false))) || ((2147483647 != 1) || (0 > -9223372036854775808)))) {
            v17 -= v1
            v1 += v17
        }
    case 0:
        if false {
            for fi106 := 0; fi106 < 20; fi106++ {
                var v107 float32 = (-2.25 / float32(1.5))
                _ = fi106
                _ = v107
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
        }
    }
    v3 += (v3 - v3)
    for fi108 := 0; fi108 < 15; fi108++ {
        var v109 int8 = -128
        switch (fi108 - (fi108 % (((fi108 * (fi108 % ((32767) | 1)))) | 1))) {
        case 0:
            if (4294967295 <= 314) {
                if ((fi108 ^ fi108) > 32768) {
                    for fi110 := 0; fi110 < 16; fi110++ {
                        _ = len(v4)
                        v111 := GenS8{F0: 42}
                        if true {
                            v4 = append(v4, (("a" + "a") + "hello"))
                            switch (("a" + ("\n\t" + ("foo bar" + ("\n\t" + "hello")))) + (("" + ("" + "foo bar")) + "a")) {
                            case "\n\t":
                                var v112 uint8 = 127
                                _ = v112
                            }
                            switch "hello" {
                            case "":
                                var v113 uint = 2
                                _ = v113
                            case "a":
                                v109 -= v109
                            }
                        }
                        _ = fi110
                        _ = v111
                    }
                    var v114 int = genFn0(v17, 2147483647, (false || true))
                    var v115 int8 = (-(v109))
                    _ = v114
                    _ = v115
                }
                v116 := GenS10{F0: (v5 | v5)}
                _ = v116
            }
        case 18:
            if (!true) {
                for fi117 := 0; fi117 < 19; fi117++ {
                    _ = len(v6)
                    if true {
                        switch "" {
                        case "a":
                            var v118 uint8 = 57
                            _ = v118
                        default:
                            v119 := v6[0]
                            _ = v119
                        }
                        v109 += -13
                        v1 = v17
                        switch "foo bar" {
                        case "foo bar":
                            v120 := GenS10{F0: (^(v5))}
                            _ = v120
                        }
                    }
                    var v121 uint16 = v17
                    if (v5 >= -1) {
                        var v122 uint64 = 376
                        _ = v122
                    } else {
                        v123 := GenS4{F0: 18446744073709551615, F1: "hello"}
                        v124 := GenS10{F0: -9223372036854775808}
                        _ = v123
                        _ = v124
                    }
                    _ = fi117
                    _ = v121
                }
                _ = len(v4)
            } else {
                switch (fi108 * (fi108 * fi108)) {
                case -774:
                    for fi125 := 0; fi125 < 2; fi125++ {
                        var v126 int64 = -9223372036854775808
                        if ((-2.25 >= float64(3.14159)) || true) {
                            v127 := GenS11{F0: 127}
                            _ = v127
                        } else {
                            v128 := GenS11{F0: 127}
                            v129 := GenS12{F0: 2}
                            _ = v128
                            _ = v129
                        }
                        v17 -= v17
                        _ = fi125
                        _ = v126
                    }
                default:
                    var v130 int8 = v109
                    _ = v130
                }
            }
        default:
            var v131 int16 = int16(v109)
            _ = v131
        }
        if (1 < -2147483648) {
            v132 := GenS13{F0: 5, F1: -958}
            for fi133 := 0; fi133 < 6; fi133++ {
                switch "\n\t" {
                case "hello":
                    _ = len(v4)
                    if (128 >= 127) {
                        var v134 float64 = 1e10
                        _ = v134
                    }
                case "":
                    v2[(fi108 | genFn0(v1, 897, (!false)))] = false
                case "\n\t":
                    var v135 uint32 = 65535
                    _ = v135
                default:
                    fmt.Println(v3, v109)
                }
                v136 := GenS14{F0: true, F1: 42}
                _ = fi133
                _ = v136
            }
            _ = v132
        } else {
            fmt.Println(v5)
            switch 0 {
            case 32767:
                if (8 > 245) {
                    v3 = 0.0
                } else {
                    if true {
                        v3 = v3
                        v5 = 42
                    } else {
                        for fi137 := 0; fi137 < 19; fi137++ {
                            delete(v2, fi137)
                            if (false || (42 >= 1)) {
                                v138 := v6[435]
                                fmt.Println(v109, v138)
                                if false {
                                    switch fi137 {
                                    case 0:
                                        if ((v138 >> ((uint(v17)) & 63)) >= v138) {
                                            v5 -= v5
                                            v109 -= -1
                                        } else {
                                            switch genFn0(v17, 42, (!false)) {
                                            case 9223372036854775807:
                                                v2[(fi137 * fi108)] = true
                                            case 746:
                                                _ = len(v2)
                                            case -880:
                                                v3 = v3
                                            default:
                                                fmt.Println(v1, fi137)
                                                var v139 uint16 = 0
                                                fmt.Println(fi137, v3, v138)
                                                _ = v139
                                            }
                                            v5 = v5
                                            fmt.Println(v138)
                                        }
                                    case 42:
                                        var v140 uint16 = v1
                                        _ = v140
                                    case -9223372036854775808:
                                        var v141 uint16 = (v17 + 0)
                                        v6[(v138 - (v138 | (v138 & v138)))] = 2147483648
                                        _ = v141
                                    }
                                    if (!(843 <= 0)) {
                                        v142 := GenS15{F0: 127, F1: 255}
                                        _ = v142
                                    }
                                    var v143 string = ("a" + "a")
                                    v143 = v143
                                    var v144 bool = true
                                    var v145 uint = 2
                                    switch v143 {
                                    case "a":
                                        if (-1 >= 1000) {
                                            v5 -= (v5 >> ((v145) & 63))
                                            var v146 float64 = v3
                                            v5 -= int64(v138)
                                            _ = v146
                                        } else {
                                            if v144 {
                                                var v147 int16 = 1
                                                _ = v147
                                            } else {
                                                var v148 bool = v144
                                                _ = v148
                                            }
                                            switch fi108 {
                                            case 9223372036854775807:
                                                var v149 uint8 = uint8(v138)
                                                _ = v149
                                            case 127:
                                                fmt.Println(fi108, v5)
                                            default:
                                                if true {
                                                    var v150 int16 = -283
                                                    if v144 {
                                                        var v151 int8 = int8(v138)
                                                        v144 = (v144 && (!v144))
                                                        v3 = 3.14159
                                                        fmt.Println(v3)
                                                        _ = v151
                                                    }
                                                    _ = v150
                                                }
                                            }
                                            switch fi108 {
                                            case 42:
                                                v4 = append(v4, (("hello" + "a") + v143))
                                            }
                                        }
                                    case "foo bar":
                                        var v152 bool = (!(v144 && (v144 || true)))
                                        _ = v152
                                    }
                                    _ = v143
                                    _ = v144
                                    _ = v145
                                } else {
                                    var v153 string = ""
                                    v4 = append(v4, v153)
                                    var v154 float32 = 1e10
                                    _ = v153
                                    _ = v154
                                }
                                _ = v138
                            }
                            switch fi137 {
                            case -376:
                                v155 := GenS16{F0: 809}
                                _ = v155
                            case 496:
                                v109 += v109
                            case 127:
                                v156 := GenS5{F0: 1, F1: ("\n\t" + ("hello" + "hello")), F2: "hello"}
                                v157 := GenS17{F0: (-2.25 > (3.14159 / 3.14159))}
                                if true {
                                    v5 = -9223372036854775808
                                    fmt.Println(v1)
                                } else {
                                    v3 = 3.14159
                                    var v158 uint = 65535
                                    _ = v158
                                }
                                _ = v156
                                _ = v157
                            }
                            _ = fi137
                        }
                    }
                    var v159 int16 = -672
                    _ = v159
                }
            }
        }
        _ = fi108
        _ = v109
    }
    v17 = v1
    v17 += v1
    v3 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
