package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 string }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 uint; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string; F1 int16; F2 int16 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int8; F2 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 int; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int16; F1 uint8 }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint }

func main() {
    var v1 int = 0
    v2 := []int{(v1 ^ -85), v1, (v1 + (v1 << ((0) & 63))), 128}
    var v3 int16 = 892
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]uint32)
    var v6 uint64 = 42
    _ = len(v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint32 = 4294967295
            _ = fi8
            _ = v9
        }
        var v10 float64 = 0.0
        _ = fi7
        _ = v10
    }
    switch ("" + "hello") {
    case "\n\t":
        v11 := GenS1{F0: 53, F1: -1, F2: "\n\t"}
        _ = v11
    case "foo bar":
        if (!(326 == 255)) {
            var v12 int8 = -118
            if false {
                v12 += v12
            } else {
                if (-2.25 <= -2.25) {
                    v2 = append(v2, v1)
                } else {
                    v13 := GenS2{F0: (v12 % (((v12 % (((v12 | v12)) | 1))) | 1)), F1: "\n\t"}
                    v12 = v12
                    var v14 float32 = 0.0
                    _ = v13
                    _ = v14
                }
                v12 += (v12 * (v12 >> ((65535) & 63)))
                var v15 uint16 = 1
                _ = v15
            }
            _ = v12
        } else {
            v1 = v1
            if ((-(v3)) < (v3 * -32768)) {
                switch v1 {
                case -101:
                    v6 = v6
                case 9223372036854775807:
                    if true {
                        delete(v5, 65535)
                        var v16 uint64 = (v6 % (((^(v6))) | 1))
                        _ = v16
                    }
                case 42:
                    var v17 float32 = -2.25
                    _ = v17
                default:
                    var v18 int8 = 0
                    _ = v18
                }
                if len(v2) > 0 {
                    v2[0] = (-(v1))
                }
            } else {
                v1 += (^(v1))
                v1 += (v1 * (v1 << ((255) & 63)))
                _ = len(v2)
            }
            var v19 int = (v1 / (((v1 % (((v1 >> ((4294967295) & 63))) | 1))) | 1))
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint32 = 2147483647
        _ = fi20
        _ = v21
    }
    var v22 uint8 = 128
    for fi23 := 0; fi23 < 7; fi23++ {
        v1 = 127
        v6 -= v6
        v22 = 1
        _ = fi23
    }
    var v24 uint8 = 128
    fmt.Println(v3, v1)
    var v25 float64 = 3.14159
    if (!false) {
        if (9223372036854775807 != -1) {
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v24)
                if (((uint64(v24) <= 18446744073709551615) && (false || (1.5 == (0.0 + 1.5)))) || (true && (v6 <= (^(v6))))) {
                    v22 = v22
                    v3 -= (^(v3))
                    fmt.Println(v6, v25)
                } else {
                    for fi27 := 0; fi27 < 9; fi27++ {
                        v1 = (fi27 / (((fi26 ^ (-(fi27)))) | 1))
                        switch (("\n\t" + "") + "") {
                        case "foo bar":
                            v1 = -823
                        default:
                            if false {
                                if false {
                                    var v28 float32 = -2.25
                                    if len(v2) > 0 {
                                        v2[0] = (^(v1))
                                    }
                                    switch 32768 {
                                    case 127:
                                        v29 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: "a"}
                                        switch "a" {
                                        case "\n\t":
                                            var v30 uint64 = (v6 + (^(v6)))
                                            _ = v30
                                        }
                                        var v31 float64 = (v25 + (v25 - 1.5))
                                        _ = v29
                                        _ = v31
                                    case 507:
                                        switch (-(v1)) {
                                        case 613:
                                            if (696 != 2147483647) {
                                                fmt.Println(fi26, v22, v28)
                                            }
                                        case 700:
                                            _ = len(v2)
                                        case 4294967295:
                                            switch "hello" {
                                            case "hello":
                                                var v32 float32 = v28
                                                _ = v32
                                            case "foo bar":
                                                if true {
                                                    v33 := GenS1{F0: -9223372036854775808, F1: int32(v22), F2: ""}
                                                    fmt.Println(fi27)
                                                    delete(v5, 42)
                                                    _ = v33
                                                } else {
                                                    v34 := GenS1{F0: -9223372036854775808, F1: 2147483647, F2: "a"}
                                                    v28 -= (0.1 - 1e10)
                                                    _ = v34
                                                }
                                            case "\n\t":
                                                switch ("\n\t" + "foo bar") {
                                                case "foo bar":
                                                    fmt.Println(v6, v3, fi27)
                                                case "\n\t":
                                                    v24 = v22
                                                }
                                                switch ((("a" + "a") + "hello") + ("a" + "foo bar")) {
                                                case "hello":
                                                    v2 = append(v2, (fi26 - (fi27 ^ (fi27 % (((-(fi26))) | 1)))))
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            default:
                                                v1 = int(v3)
                                            }
                                        default:
                                            if true {
                                                v35 := v4["foo bar"]
                                                _ = v35
                                            } else {
                                                switch fi27 {
                                                case 0:
                                                    switch 9223372036854775807 {
                                                    case 0:
                                                        fmt.Println(v1)
                                                    default:
                                                        switch "a" {
                                                        case "\n\t":
                                                            if (!(false || (false || (false && (!((!false) && true)))))) {
                                                                v6 = (v6 - (v6 % ((v6) | 1)))
                                                                fmt.Println(v1, v1, v3)
                                                                delete(v5, 756)
                                                                switch ("hello" + "a") {
                                                                case "":
                                                                    v36 := GenS1{F0: -1, F1: -661, F2: ""}
                                                                    _ = v36
                                                                default:
                                                                    v6 = v6
                                                                }
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v6 = 4294967295
                                                                case "\n\t":
                                                                    v37 := GenS2{F0: 120, F1: "foo bar"}
                                                                    _ = v37
                                                                default:
                                                                    v22 -= (v22 | 42)
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = -493
                                                                }
                                                            } else {
                                                                v28 = 3.14159
                                                            }
                                                        case "hello":
                                                            v38 := v4[("hello" + "\n\t")]
                                                            _ = v38
                                                        default:
                                                            var v39 uint32 = 65535
                                                            _ = v39
                                                        }
                                                    }
                                                case -2147483648:
                                                    v25 -= v25
                                                }
                                                var v40 float32 = 0.1
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = v40
                                            }
                                        }
                                    case 2147483647:
                                        v28 = 1e10
                                    default:
                                        v22 += (^(v22))
                                    }
                                    _ = v28
                                } else {
                                    fmt.Println(fi26)
                                    v25 -= ((-2.25 / 1.0) - float64(1.5))
                                }
                                if false {
                                    if ((v1 | (fi27 & (fi27 ^ 1))) < (-(fi26))) {
                                        var v41 float64 = 0.0
                                        _ = v41
                                    }
                                    fmt.Println(v24)
                                    switch (fi27 >> ((42) & 63)) {
                                    case 553:
                                        fmt.Println(v25)
                                    case 521:
                                        var v42 int64 = 6
                                        _ = v42
                                    case 2147483647:
                                        v43 := GenS1{F0: 9223372036854775807, F1: -2147483648, F2: ("" + ("hello" + "\n\t"))}
                                        _ = v43
                                    }
                                }
                            } else {
                                if (false || false) {
                                    var v44 int = (fi27 - (fi26 ^ 2147483647))
                                    fmt.Println(fi27, fi27)
                                    var v45 uint32 = 2147483647
                                    var v46 bool = false
                                    v1 = (fi27 - fi26)
                                    switch "a" {
                                    case "":
                                        switch "a" {
                                        case "\n\t":
                                            var v47 uint16 = 255
                                            _ = v47
                                        case "":
                                            v48 := GenS3{F0: 18446744073709551615, F1: 2147483647, F2: v45}
                                            _ = v48
                                        default:
                                            v22 = v24
                                        }
                                    case "\n\t":
                                        _ = len(v4)
                                    case "a":
                                        v24 = 1
                                    default:
                                        v44 = (fi27 ^ fi27)
                                    }
                                    v25 = v25
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                } else {
                                    v22 -= 1
                                    if false {
                                        v6 = uint64(v24)
                                        v24 = 97
                                    } else {
                                        if (v25 == float64(3.14159)) {
                                            switch (v1 ^ (fi27 << ((4294967295) & 63))) {
                                            case 9223372036854775807:
                                                if (((0.0 - 0.1) / 1e10) == 1e10) {
                                                    var v49 int = 1
                                                    if (v1 <= fi27) {
                                                        if true {
                                                            var v50 float32 = float32(v25)
                                                            _ = v50
                                                        }
                                                        if true {
                                                            v24 += v24
                                                            var v51 uint64 = v6
                                                            v49 -= v1
                                                            _ = v51
                                                        }
                                                        fmt.Println(v3)
                                                    } else {
                                                        v49 -= v1
                                                    }
                                                    _ = v49
                                                }
                                            case 42:
                                                v6 = (v6 >> ((65535) & 63))
                                                _ = len(v2)
                                            default:
                                                var v52 uint64 = v6
                                                _ = v52
                                            }
                                            v24 += (v22 ^ v24)
                                            var v53 int64 = -1
                                            _ = v53
                                        }
                                        var v54 int = (fi27 % ((36) | 1))
                                        v25 = v25
                                        _ = v54
                                    }
                                    v6 = (v6 >> ((18446744073709551615) & 63))
                                }
                                v55 := v4["hello"]
                                _ = v55
                            }
                            v3 += 42
                            v3 += int16(fi27)
                        }
                        _ = fi27
                    }
                    delete(v4, (((("foo bar" + "") + "a") + ("hello" + "")) + ""))
                    var v56 int32 = -207
                    _ = v56
                }
                v57 := v5[1]
                _ = fi26
                _ = v57
            }
            for fi58 := 0; fi58 < 8; fi58++ {
                var v59 int16 = v3
                v60 := GenS4{F0: "a", F1: (v3 % (((v3 - v59)) | 1)), F2: 0}
                switch "a" {
                case "hello":
                    v61 := GenS3{F0: 2, F1: 2147483647, F2: 2147483647}
                    _ = v61
                default:
                    delete(v5, uint32(v24))
                }
                _ = fi58
                _ = v59
                _ = v60
            }
            for fi62 := 0; fi62 < 17; fi62++ {
                if true {
                    v6 = (v6 / (((v6 / ((v6) | 1))) | 1))
                    v63 := GenS4{F0: (("\n\t" + "foo bar") + ("" + "hello")), F1: v3, F2: 955}
                    for fi64 := 0; fi64 < 8; fi64++ {
                        _ = len(v5)
                        v65 := GenS5{F0: 2}
                        _ = fi64
                        _ = v65
                    }
                    _ = v63
                } else {
                    var v66 uint8 = 37
                    var v67 int = (fi62 | fi62)
                    v22 = v24
                    _ = v66
                    _ = v67
                }
                _ = fi62
            }
        } else {
            if (true && (("" + (("hello" + "foo bar") + "a")) != "\n\t")) {
                for fi68 := 0; fi68 < 6; fi68++ {
                    if true {
                        fmt.Println(v3, v1)
                        v69 := GenS6{F0: -9223372036854775808, F1: 100, F2: 2}
                        _ = v69
                    } else {
                        for fi70 := 0; fi70 < 15; fi70++ {
                            v71 := GenS6{F0: 1, F1: 127, F2: 1}
                            var v72 string = (("a" + "a") + "")
                            if ((42 == 42) || true) {
                                v73 := GenS2{F0: -128, F1: "hello"}
                                _ = v73
                            } else {
                                v22 += v22
                                if len(v2) > 0 {
                                    v2[0] = (v1 % (((fi68 - (v1 / ((fi68) | 1)))) | 1))
                                }
                                fmt.Println(fi68, v3, v22)
                            }
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                        var v74 int = (fi68 ^ (fi68 * fi68))
                        var v75 float64 = v25
                        _ = v74
                        _ = v75
                    }
                    _ = fi68
                }
                for fi76 := 0; fi76 < 11; fi76++ {
                    v1 = -9223372036854775808
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v77 := GenS1{F0: -579, F1: 42, F2: "a"}
                        var v78 uint8 = (v24 % (((^(v24))) | 1))
                        _ = v77
                        _ = v78
                    }
                    switch "hello" {
                    case "hello":
                        v1 = (v1 / ((int(v22)) | 1))
                    case "\n\t":
                        v6 = 805
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi76
                }
            } else {
                switch "\n\t" {
                case "a":
                    var v79 float32 = 1.5
                    _ = v79
                case "":
                    v80 := v5[2147483647]
                    _ = v80
                default:
                    if (2 > 289) {
                        v81 := GenS7{F0: "foo bar"}
                        _ = v81
                    }
                }
            }
            if true {
                for fi82 := 0; fi82 < 4; fi82++ {
                    for fi83 := 0; fi83 < 15; fi83++ {
                        fmt.Println(v25)
                        v3 = v3
                        var v84 uint32 = 65535
                        _ = fi83
                        _ = v84
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v24 = v22
                    _ = fi82
                }
            }
            v25 = (float64(0.0) + v25)
        }
        v6 = (^(v6))
        if (!(42 <= int32(v3))) {
            if true {
                fmt.Println(v6, v6, v22)
                v25 = (1e10 + 3.14159)
            } else {
                v85 := GenS8{F0: (false || ((0.1 - float64(float32((v25 - v25)))) >= (3.14159 * float64(float32((v25 + v25)))))), F1: 0}
                _ = v85
            }
        }
    } else {
        v6 = (v6 * v6)
        for fi86 := 0; fi86 < 13; fi86++ {
            v87 := GenS9{F0: fi86, F1: 42, F2: 255}
            switch "" {
            case "a":
                fmt.Println(v25, v1)
            case "foo bar":
                v25 = v25
            default:
                _ = len(v2)
            }
            if len(v2) > 0 {
                v2[0] = (-(v1))
            }
            _ = fi86
            _ = v87
        }
        if len(v2) > 0 {
            v2[0] = v1
        }
    }
    if (v1 >= 9223372036854775807) {
        for fi88 := 0; fi88 < 19; fi88++ {
            v6 -= v6
            _ = fi88
        }
        fmt.Println(v1)
    }
    if (!false) {
        if (729 >= -76) {
            var v89 int64 = -1
            _ = v89
        }
        var v90 uint64 = v6
        _ = v90
    } else {
        if ((float64(-2.25) - (float64((1.5 + 1.5)) * 3.14159)) <= (v25 + v25)) {
            v6 = 0
        } else {
            v24 = v22
            delete(v5, 1)
        }
        var v91 uint32 = 509
        if false {
            var v92 uint = 848
            if (42 <= (^(v24))) {
                fmt.Println(v3, v6)
            } else {
                var v93 bool = ("a" < "a")
                if v93 {
                    var v94 string = "hello"
                    fmt.Println(v24, v24, v94)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v94
                } else {
                    if v93 {
                        v92 = uint(v24)
                        v4["\n\t"] = 65535
                    } else {
                        v95 := GenS1{F0: 42, F1: -936, F2: ""}
                        _ = len(v2)
                        _ = v95
                    }
                    var v96 string = (("a" + ("" + "foo bar")) + "\n\t")
                    if (127 <= 127) {
                        if (v93 && ((!true) || false)) {
                            fmt.Println(v6, v22)
                            var v97 uint8 = v22
                            for fi98 := 0; fi98 < 8; fi98++ {
                                var v99 uint64 = (v6 & 4294967295)
                                _ = fi98
                                _ = v99
                            }
                            _ = v97
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if (2 == 1) {
                            v100 := GenS10{F0: v91}
                            _ = v100
                        }
                        if false {
                            switch (v1 + v1) {
                            case 0:
                                switch "" {
                                case "foo bar":
                                    _ = len(v4)
                                case "\n\t":
                                    var v101 uint8 = (^(v24))
                                    _ = v101
                                case "":
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v24, v24, v92)
                                    case "\n\t":
                                        v24 -= (v22 / ((2) | 1))
                                    default:
                                        v93 = v93
                                    }
                                default:
                                    v25 = -2.25
                                }
                            case 127:
                                fmt.Println(v91)
                            default:
                                for fi102 := 0; fi102 < 17; fi102++ {
                                    v22 = 1
                                    if v93 {
                                        for fi103 := 0; fi103 < 3; fi103++ {
                                            var v104 uint16 = 185
                                            switch (v1 - int(v6)) {
                                            case 9223372036854775807:
                                                var v105 int32 = 42
                                                _ = v105
                                            case 4294967295:
                                                v92 = v92
                                                fmt.Println(v92, v91, v104)
                                            case 0:
                                                if (v93 || ((v1 >> (((v92 / ((42) | 1))) & 63)) == (v1 ^ (fi103 >> (((^(v92))) & 63))))) {
                                                    v92 = v92
                                                } else {
                                                    var v106 int32 = 1
                                                    _ = v106
                                                }
                                            }
                                            fmt.Println(fi103, v91, v92)
                                            _ = fi103
                                            _ = v104
                                        }
                                        if v93 {
                                            for fi107 := 0; fi107 < 11; fi107++ {
                                                switch ((v96 + ("\n\t" + (v96 + "a"))) + "foo bar") {
                                                case "\n\t":
                                                    v6 = 4294967295
                                                }
                                                if v93 {
                                                    v25 += ((v25 - (v25 / v25)) + v25)
                                                    if v93 {
                                                        v108 := GenS11{F0: -373, F1: (v24 * v22)}
                                                        var v109 uint = v92
                                                        v5[uint32(v6)] = 0
                                                        _ = v108
                                                        _ = v109
                                                    } else {
                                                        v110 := GenS12{F0: v3}
                                                        v111 := GenS11{F0: v3, F1: (v24 | (v24 / ((uint8(v91)) | 1)))}
                                                        if v93 {
                                                            var v112 int = fi107
                                                            _ = v112
                                                        }
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                } else {
                                                    fmt.Println(v24)
                                                }
                                                v113 := GenS5{F0: uint(v24)}
                                                _ = fi107
                                                _ = v113
                                            }
                                            v114 := GenS13{F0: v92}
                                            _ = v114
                                        } else {
                                            for fi115 := 0; fi115 < 16; fi115++ {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = fi115
                                            }
                                            v25 -= 0.0
                                            fmt.Println(v93, v93, v91)
                                        }
                                    }
                                    _ = fi102
                                }
                            }
                            v92 = 2
                        }
                    }
                    _ = v96
                }
                for fi116 := 0; fi116 < 14; fi116++ {
                    for fi117 := 0; fi117 < 6; fi117++ {
                        v5[(v91 << ((1) & 63))] = v91
                        v25 -= v25
                        _ = fi117
                    }
                    v6 += (v6 | v6)
                    var v118 int32 = 2147483647
                    _ = fi116
                    _ = v118
                }
                _ = v93
            }
            v22 -= v24
            _ = v92
        } else {
            v25 -= 1e10
            var v119 uint64 = (v6 - v6)
            if true {
                var v120 float64 = (((0.0 / (v25 + v25)) + float64(1.5)) - v25)
                v25 += v120
                v121 := GenS6{F0: 9223372036854775807, F1: 1, F2: v91}
                _ = v120
                _ = v121
            }
            _ = v119
        }
        _ = v91
    }
    if true {
        switch "foo bar" {
        case "":
            switch v1 {
            case 32767:
                v24 = 255
            }
        default:
            for fi122 := 0; fi122 < 15; fi122++ {
                fmt.Println(v6, v25)
                v123 := GenS9{F0: (v1 * int(v24)), F1: 803, F2: 1}
                _ = fi122
                _ = v123
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
