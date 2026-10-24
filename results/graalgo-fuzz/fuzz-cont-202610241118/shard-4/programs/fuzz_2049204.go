package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 bool; F1 int8; F2 uint64 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 bool; F1 bool; F2 bool }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint64; F1 uint32 }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 bool; F1 int32; F2 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint8; F1 int32; F2 int16 }
type GenS19 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 int32, p1 uint8) uint8 {
    v4 := []int{128}
    if true {
        fmt.Println(p1)
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        if (("" + "") >= "a") {
            if len(v4) > 0 {
                v4[0] = -2147483648
            }
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS1{F0: 42}
                _ = fi6
                _ = v7
            }
            switch (^(fi5)) {
            case -2147483648:
                if false {
                    var v8 float32 = 0.0
                    var v9 uint8 = p1
                    var v10 uint = 65535
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 3; fi11++ {
                        fmt.Println(fi11)
                        p1 += p1
                        if true {
                            switch "" {
                            case "\n\t":
                                p0 = (p0 >> ((2) & 63))
                            case "foo bar":
                                p0 -= (p0 & p0)
                            default:
                                var v12 int16 = 48
                                _ = v12
                            }
                        } else {
                            v13 := GenS2{F0: 257, F1: -680, F2: 0}
                            _ = v13
                        }
                        _ = fi11
                    }
                    v14 := GenS3{F0: 42, F1: (p1 / (((p1 & 138)) | 1))}
                    _ = v14
                }
            default:
                var v15 uint = 65535
                _ = v15
            }
        } else {
            v4 = append(v4, fi5)
            var v16 int = fi5
            _ = v16
        }
        _ = fi5
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v4
    return (p1 >> ((255) & 63))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{2, 42, 1, 2147483647}
    var v3 int64 = -508
    if (false && true) {
        v3 -= (v3 % ((-9223372036854775808) | 1))
    }
    var v17 int16 = 42
    v17 += (-(v17))
    if false {
        v18 := GenS1{F0: (v17 - v17)}
        var v19 int16 = (v17 ^ 1)
        var v20 int16 = v19
        _ = v18
        _ = v19
        _ = v20
    }
    switch ("" + "foo bar") {
    case "\n\t":
        var v21 uint16 = 255
        _ = v21
    case "hello":
        switch 32768 {
        case 1:
            v22 := GenS4{F0: true, F1: -71, F2: 2}
            _ = v22
        case 32768:
            fmt.Println(v3, v3)
        case -320:
            if (false || (v3 > (v3 | v3))) {
                v23 := GenS5{F0: 127, F1: 649}
                var v24 int8 = -127
                v25 := GenS6{F0: true, F1: true, F2: (150 == 65535)}
                _ = v23
                _ = v24
                _ = v25
            } else {
                delete(v1, -2147483648)
            }
        default:
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 float32 = 1e10
                var v28 int32 = 0
                _ = fi26
                _ = v27
                _ = v28
            }
        }
    case "a":
        for fi29 := 0; fi29 < 9; fi29++ {
            switch (fi29 ^ -1) {
            case 42:
                if true {
                    fmt.Println(fi29, v3, fi29)
                    v30 := GenS1{F0: -590}
                    _ = v30
                }
            default:
                v31 := GenS7{F0: 391, F1: 18446744073709551615}
                _ = v31
            }
            var v32 uint16 = 935
            var v33 uint32 = uint32(v32)
            _ = fi29
            _ = v32
            _ = v33
        }
    }
    _ = len(v1)
    _ = len(v2)
    for fi34 := 0; fi34 < 8; fi34++ {
        v3 = v3
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 = (v3 << ((0) & 63))
        _ = fi34
    }
    switch "\n\t" {
    case "\n\t":
        if false {
            var v35 uint8 = 128
            for fi36 := 0; fi36 < 17; fi36++ {
                v37 := GenS3{F0: 948, F1: (v35 ^ genFn0(-2147483648, (v35 / (((v35 | v35)) | 1))))}
                fmt.Println(v17, fi36)
                if false {
                    for fi38 := 0; fi38 < 2; fi38++ {
                        if (!true) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v3)
                        } else {
                            switch "foo bar" {
                            case "hello":
                                v3 -= (v3 % (((v3 >> ((2) & 63))) | 1))
                            case "a":
                                switch ("hello" + ("hello" + "hello")) {
                                case "":
                                    v35 += (v35 & genFn0(41, (v35 / ((v35) | 1))))
                                case "hello":
                                    v39 := GenS6{F0: true, F1: (243 > 18446744073709551615), F2: false}
                                    _ = v39
                                }
                            }
                            var v40 uint16 = 713
                            v40 = v40
                            _ = v40
                        }
                        _ = len(v2)
                        var v41 float64 = (1.5 / 3.14159)
                        _ = fi38
                        _ = v41
                    }
                    v42 := GenS4{F0: true, F1: 127, F2: 18446744073709551615}
                    _ = v42
                } else {
                    for fi43 := 0; fi43 < 7; fi43++ {
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        v35 += (v35 * 255)
                        if (!false) {
                            var v44 uint8 = 255
                            v44 += (^(v35))
                            var v45 uint8 = (v35 + 255)
                            v1[(fi36 ^ -1)] = 478
                            _ = v44
                            _ = v45
                        } else {
                            v46 := GenS8{F0: (2 <= 0), F1: 18446744073709551615}
                            var v47 float32 = float32(0.1)
                            v17 -= 1000
                            v48 := GenS1{F0: (v17 & (v17 ^ v17))}
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        v3 += -179
                        _ = fi43
                    }
                    v17 += (v17 - (-(v17)))
                    for fi49 := 0; fi49 < 6; fi49++ {
                        switch (fi49 * fi49) {
                        case 9223372036854775807:
                            switch "hello" {
                            case "hello":
                                v50 := GenS9{F0: "hello"}
                                _ = v50
                            case "":
                                delete(v1, -1)
                            }
                        }
                        _ = fi49
                    }
                }
                _ = fi36
                _ = v37
            }
            _ = v35
        }
    case "foo bar":
        if (!((554 >= 0) || (!(127 < 127)))) {
            v3 = v3
            v3 = -9223372036854775808
        }
    default:
        v51 := GenS10{F0: 2, F1: 2147483647, F2: -223}
        _ = v51
    }
    var v52 float32 = ((-2.25 / -2.25) / 1.5)
    switch 128 {
    case 0:
        fmt.Println(v17, v17, v52)
    case -994:
        for fi53 := 0; fi53 < 20; fi53++ {
            var v54 uint64 = 42
            v55 := GenS11{F0: 42, F1: "a"}
            var v56 float64 = 0.0
            _ = fi53
            _ = v54
            _ = v55
            _ = v56
        }
    case 952:
        for fi57 := 0; fi57 < 3; fi57++ {
            switch "\n\t" {
            case "a":
                v3 -= (v3 & int64(v17))
            default:
                if true {
                    var v58 uint64 = 1
                    v52 += 0.1
                    _ = v58
                } else {
                    var v59 uint64 = 597
                    var v60 uint64 = v59
                    _ = v59
                    _ = v60
                }
            }
            v61 := GenS12{F0: 866, F1: 0}
            var v62 int32 = 319
            _ = fi57
            _ = v61
            _ = v62
        }
    default:
        for fi63 := 0; fi63 < 1; fi63++ {
            if false {
                if (((1.5 - 1.5) - 1e10) >= 0.0) {
                    var v64 int16 = (v17 << ((2) & 63))
                    v65 := GenS13{F0: 2, F1: false}
                    _ = v64
                    _ = v65
                } else {
                    fmt.Println(v52, v17)
                    for fi66 := 0; fi66 < 13; fi66++ {
                        fmt.Println(v52, fi63)
                        _ = fi66
                    }
                    var v67 float32 = v52
                    _ = v67
                }
                v68 := GenS4{F0: (((!(0 == 328)) && ((v3 >> ((575) & 63)) > -9223372036854775808)) && ((0 > (v3 ^ (v3 / ((v3) | 1)))) || (32767 <= 65535))), F1: -1, F2: 18446744073709551615}
                for fi69 := 0; fi69 < 20; fi69++ {
                    switch "" {
                    case "foo bar":
                        v70 := v1[fi69]
                        _ = v70
                    case "hello":
                        var v71 int16 = (v17 & v17)
                        _ = len(v1)
                        var v72 int16 = (v17 >> ((500) & 63))
                        _ = v71
                        _ = v72
                    default:
                        v73 := GenS4{F0: (32767 >= 1), F1: 1, F2: 2}
                        _ = v73
                    }
                    switch ((("foo bar" + ("hello" + "")) + ("" + ("" + "\n\t"))) + "hello") {
                    case "foo bar":
                        if (42 > 127) {
                            switch ("a" + "a") {
                            case "foo bar":
                                fmt.Println(v52, fi69, fi63)
                                delete(v1, 1)
                            case "a":
                                fmt.Println(v17, v3, fi69)
                                fmt.Println(v52)
                            case "hello":
                                var v74 bool = false
                                _ = v74
                            default:
                                if true {
                                    if (431 != (-(v3))) {
                                        if (0 == 2) {
                                            if false {
                                                var v75 float64 = ((0.0 + (3.14159 + 1e10)) * 0.1)
                                                _ = v75
                                            } else {
                                                v3 += -1
                                            }
                                            var v76 bool = true
                                            v76 = true
                                            v77 := GenS14{F0: v76, F1: 0, F2: 0}
                                            fmt.Println(v76)
                                            _ = v76
                                            _ = v77
                                        }
                                        switch (-(fi69)) {
                                        case 32768:
                                            var v78 int16 = (v17 << ((65535) & 63))
                                            _ = v78
                                        case 127:
                                            if (v3 != 42) {
                                                v17 += 1
                                                var v79 int16 = (v17 << ((1) & 63))
                                                fmt.Println(v3, fi69)
                                                _ = v79
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(fi63, v52, v3)
                                            }
                                        case 32767:
                                            switch -2147483648 {
                                            case -9223372036854775808:
                                                delete(v1, -2147483648)
                                                v17 = (v17 * v17)
                                                if false {
                                                    var v80 uint16 = 457
                                                    if true {
                                                        v52 -= 3.14159
                                                        v3 = 42
                                                    }
                                                    v3 = v3
                                                    switch (fi63 ^ 32767) {
                                                    case 9223372036854775807:
                                                        fmt.Println(fi63, v52, v3)
                                                    case 128:
                                                        var v81 int64 = v3
                                                        _ = v81
                                                    case 735:
                                                        if false {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if false {
                                                                    var v82 uint64 = 42
                                                                    delete(v1, fi63)
                                                                    _ = v82
                                                                } else {
                                                                    if true {
                                                                        v52 = ((3.14159 / 3.14159) * 0.0)
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v83 := v1[(fi63 >> ((255) & 63))]
                                                                        _ = v83
                                                                    }
                                                                    v3 -= (-(v3))
                                                                    v52 += v52
                                                                }
                                                            }
                                                        } else {
                                                            if (1.5 >= -2.25) {
                                                                v17 += (-(v17))
                                                            } else {
                                                                v1[-2147483648] = 527
                                                                fmt.Println(v3, fi69)
                                                                if false {
                                                                    fmt.Println(v3)
                                                                    if ((2 == 0) && false) {
                                                                        var v84 int8 = -104
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v84
                                                                    }
                                                                    if false {
                                                                        if ((fi69 % ((fi69) | 1)) == fi69) {
                                                                            fmt.Println(fi69, v80, v80)
                                                                            v3 += -1
                                                                        } else {
                                                                            v3 = (v3 - v3)
                                                                            v85 := GenS12{F0: 18446744073709551615, F1: 219}
                                                                            if false {
                                                                                v3 += (v3 << ((2) & 63))
                                                                                _ = len(v1)
                                                                                if false {
                                                                                    v80 -= 1
                                                                                    fmt.Println(v17, v3, v80)
                                                                                }
                                                                                v17 += -1
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        v86 := GenS4{F0: (!false), F1: 42, F2: 18446744073709551615}
                                                                        _ = v86
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(fi69, v17, fi69)
                                                            if len(v2) > 0 {
                                                                v2[0] = 4294967295
                                                            }
                                                            fmt.Println(v3, fi69, v17)
                                                        }
                                                    }
                                                    var v87 uint32 = 1
                                                    _ = v80
                                                    _ = v87
                                                }
                                                var v88 float32 = (v52 / 1.0)
                                                _ = v88
                                            default:
                                                var v89 int = 4294967295
                                                _ = v89
                                            }
                                        default:
                                            var v90 float32 = -2.25
                                            _ = v90
                                        }
                                        var v91 int8 = int8(fi69)
                                        v52 += -2.25
                                        _ = v91
                                    }
                                }
                            }
                            if ((v17 >> ((14) & 63)) != 1000) {
                                switch (fi63 + (fi69 >> ((308) & 63))) {
                                case -2147483648:
                                    var v92 uint = 42
                                    _ = v92
                                case -45:
                                    v3 = (v3 % ((338) | 1))
                                default:
                                    v3 -= v3
                                }
                                switch ((("hello" + "hello") + ("" + "")) + (("\n\t" + "foo bar") + "\n\t")) {
                                case "":
                                    var v93 int8 = -1
                                    _ = v93
                                case "a":
                                    v94 := v1[(fi69 << ((4294967295) & 63))]
                                    v1[fi69] = v94
                                    var v95 int32 = 0
                                    _ = v94
                                    _ = v95
                                }
                            }
                            _ = len(v2)
                            fmt.Println(fi69, v3)
                            if (v3 != (v3 % ((-9223372036854775808) | 1))) {
                                v52 = v52
                                var v96 uint32 = 2147483647
                                var v97 bool = (false || false)
                                _ = v96
                                _ = v97
                            } else {
                                v98 := GenS11{F0: 2, F1: ("a" + ("foo bar" + "foo bar"))}
                                _ = v98
                            }
                        }
                    case "hello":
                        if ((v17 << ((255) & 63)) != (v17 * (v17 + v17))) {
                            var v99 uint16 = 65535
                            _ = v99
                        } else {
                            var v100 int32 = 42
                            if false {
                                if false {
                                    if true {
                                        var v101 float32 = (0.0 / 1.0)
                                        _ = v101
                                    } else {
                                        switch (fi63 | (fi69 << ((65535) & 63))) {
                                        case 4294967295:
                                            if (65535 > 999) {
                                                var v102 int32 = (v100 * 1)
                                                v103 := GenS13{F0: 159, F1: false}
                                                var v104 string = "foo bar"
                                                _ = v102
                                                _ = v103
                                                _ = v104
                                            }
                                        }
                                        if false {
                                            var v105 int16 = 32767
                                            var v106 uint64 = 1
                                            _ = v105
                                            _ = v106
                                        } else {
                                            switch (fi63 * (fi69 & (fi69 + fi63))) {
                                            case -48:
                                                if true {
                                                    v17 += (v17 << ((42) & 63))
                                                    var v107 uint = 0
                                                    if ((v100 < v100) || ((128 >= 246) || (!true))) {
                                                        v108 := GenS7{F0: 1, F1: 2}
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483648
                                                        }
                                                        _ = v108
                                                    } else {
                                                        var v109 bool = true
                                                        v100 = v100
                                                        _ = v109
                                                    }
                                                    _ = v107
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    if true {
                                        v110 := GenS6{F0: true, F1: (42 <= 2), F2: (3.14159 < 3.14159)}
                                        fmt.Println(v52)
                                        _ = v110
                                    } else {
                                        var v111 string = ("hello" + ("a" + (("foo bar" + (("\n\t" + "foo bar") + "foo bar")) + "")))
                                        if (v3 != v3) {
                                            v1[-2147483648] = 1
                                        }
                                        v112 := GenS15{F0: 0}
                                        _ = v111
                                        _ = v112
                                    }
                                    v100 = (v100 ^ (v100 * v100))
                                    var v113 int = (fi63 + (fi69 % ((32768) | 1)))
                                    _ = v113
                                }
                            }
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            _ = v100
                        }
                    }
                    _ = fi69
                }
                _ = v68
            } else {
                v52 = 0.1
                v114 := GenS16{F0: false}
                _ = v114
            }
            if false {
                for fi115 := 0; fi115 < 4; fi115++ {
                    var v116 int8 = -2
                    var v117 uint32 = 647
                    _ = fi115
                    _ = v116
                    _ = v117
                }
                fmt.Println(fi63, v52)
                for fi118 := 0; fi118 < 20; fi118++ {
                    v119 := GenS17{F0: 100}
                    v2 = append(v2, 4294967295)
                    v17 += -1
                    _ = fi118
                    _ = v119
                }
            } else {
                v120 := GenS11{F0: 18446744073709551615, F1: "foo bar"}
                var v121 uint16 = 1
                _ = v120
                _ = v121
            }
            _ = fi63
        }
    }
    for fi122 := 0; fi122 < 9; fi122++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi123 := 0; fi123 < 12; fi123++ {
            var v124 float32 = v52
            _ = fi123
            _ = v124
        }
        _ = fi122
    }
    v125 := GenS17{F0: -9}
    switch "" {
    case "\n\t":
        switch "" {
        case "a":
            v17 += v17
        case "":
            if (!true) {
                for fi126 := 0; fi126 < 4; fi126++ {
                    if false {
                        if (false || (((576 < 141) || (true && ((v17 << ((4294967295) & 63)) <= (^(v17))))) || true)) {
                            var v127 bool = (!((128 >= 43) || (1.5 > 0.1)))
                            var v128 uint32 = 4294967295
                            _ = v127
                            _ = v128
                        } else {
                            if (!(((float32(0.1) - (v52 + v52)) * v52) > float32((0.1 / (1.5 - 3.14159))))) {
                                v3 -= -772
                            }
                            switch "a" {
                            case "hello":
                                fmt.Println(v52, fi126, fi126)
                            }
                        }
                        for fi129 := 0; fi129 < 18; fi129++ {
                            v52 = (v52 + (v52 - v52))
                            _ = fi129
                        }
                    } else {
                        fmt.Println(v3)
                        fmt.Println(v3, v3, v3)
                        v130 := GenS18{F0: 31, F1: 0, F2: 0}
                        _ = v130
                    }
                    if len(v2) > 0 {
                        v2[0] = 408
                    }
                    _ = fi126
                }
                v131 := GenS19{F0: -198, F1: v3, F2: 900}
                v1[969] = 4294967295
                _ = v131
            }
        case "foo bar":
            v132 := GenS15{F0: 713}
            _ = v132
        }
    default:
        fmt.Println(v52, v52, v3)
    }
    if (-2.25 == (v52 * v52)) {
        v3 = 9223372036854775807
        var v133 int64 = (v3 - (v3 | v3))
        _ = v133
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v125
    _ = v17
    _ = v2
    _ = v3
    _ = v52
}
