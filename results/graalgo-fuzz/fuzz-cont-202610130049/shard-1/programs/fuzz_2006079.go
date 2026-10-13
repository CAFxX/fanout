package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 string; F2 uint16 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 string; F1 int32; F2 uint32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int16; F1 uint; F2 uint }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int64; F1 bool }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 int8; F1 int16 }
func genFn0(p0 int32, p1 int64, p2 uint8) int8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        switch int(p2) {
        case -612:
            v6 := GenS1{F0: 0, F1: 4294967295}
            _ = v6
        }
        _ = fi5
    }
    var v7 uint16 = 255
    p2 = (^(p2))
    switch 32767 {
    case 124:
        v8 := []int{-1, -858}
        _ = v8
    default:
        v9 := []int{918}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint = v1
    v3 := []uint32{}
    var v4 uint32 = 65535
    _ = len(v3)
    if (true && true) {
        var v10 uint64 = 51
        if true {
            var v11 float64 = 0.1
            for fi12 := 0; fi12 < 18; fi12++ {
                var v13 uint8 = 2
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = v10
    } else {
        if ((-9223372036854775808 >= 336) && false) {
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 string = ""
                v2 -= (^(v1))
                if (v4 == 234) {
                    if true {
                        switch fi14 {
                        case 953:
                            v16 := GenS2{F0: 215, F1: 585}
                            _ = v16
                        case 4294967295:
                            var v17 uint16 = 2
                            _ = v17
                        case 42:
                            var v18 int32 = 1
                            _ = v18
                        }
                        var v19 int32 = 243
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                        _ = v19
                    } else {
                        v20 := GenS2{F0: 127, F1: 0}
                        _ = v20
                    }
                } else {
                    var v21 int32 = 0
                    v21 = v21
                    _ = v21
                }
                _ = fi14
                _ = v15
            }
        }
        v1 = uint(v4)
        fmt.Println(v2)
    }
    v22 := GenS3{F0: 425}
    var v23 uint8 = 25
    for fi24 := 0; fi24 < 19; fi24++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v25 int8 = genFn0(2147483647, -20, v23)
        _ = fi24
        _ = v25
    }
    switch 9223372036854775807 {
    case -2147483648:
        var v26 int16 = 1000
        _ = v26
    case -9223372036854775808:
        v3 = append(v3, v4)
    case 4294967295:
        var v27 uint64 = 18446744073709551615
        _ = v27
    }
    var v28 int64 = -698
    v29 := GenS3{F0: 0}
    if (!(false || true)) {
        for fi30 := 0; fi30 < 12; fi30++ {
            if false {
                switch "a" {
                case "a":
                    for fi31 := 0; fi31 < 5; fi31++ {
                        if false {
                            switch "" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = v4
                                }
                            default:
                                v28 -= (v28 & int64(v2))
                            }
                            fmt.Println(v4, v1)
                            v28 += (v28 % ((v28) | 1))
                        }
                        if len(v3) > 0 {
                            v3[0] = (v4 & (v4 * v4))
                        }
                        _ = fi31
                    }
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = 805
                    }
                }
            } else {
                if true {
                    v32 := GenS3{F0: 18446744073709551615}
                    switch ("a" + ("foo bar" + "hello")) {
                    case "\n\t":
                        var v33 float32 = (1.5 + (0.1 * 0.1))
                        _ = v33
                    case "":
                        v1 = (v1 << ((65535) & 63))
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v32
                } else {
                    switch 9223372036854775807 {
                    case 128:
                        fmt.Println(v23)
                    case -9223372036854775808:
                        var v34 int32 = -2147483648
                        _ = v34
                    case -125:
                        for fi35 := 0; fi35 < 11; fi35++ {
                            if true {
                                v3 = append(v3, v4)
                            }
                            _ = fi35
                        }
                    }
                }
                switch ("foo bar" + ((("a" + "") + "foo bar") + "a")) {
                case "foo bar":
                    var v36 uint = (v1 - v1)
                    _ = v36
                case "a":
                    for fi37 := 0; fi37 < 18; fi37++ {
                        var v38 int64 = v28
                        v39 := GenS3{F0: 42}
                        v40 := GenS1{F0: 124, F1: 1}
                        _ = fi37
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                default:
                    if len(v3) > 0 {
                        v3[0] = v4
                    }
                }
            }
            var v41 bool = true
            if v41 {
                v42 := GenS4{F0: (fi30 + (fi30 ^ fi30)), F1: "", F2: 679}
                var v43 int16 = 159
                _ = v42
                _ = v43
            }
            _ = fi30
            _ = v41
        }
    } else {
        if (!(v4 < v4)) {
            switch -644 {
            case 2147483647:
                var v44 uint64 = 18446744073709551615
                _ = v44
            default:
                v23 -= v23
            }
            var v45 uint = (v2 << (((v2 << ((v2) & 63))) & 63))
            for fi46 := 0; fi46 < 18; fi46++ {
                switch (^(fi46)) {
                case 0:
                    for fi47 := 0; fi47 < 4; fi47++ {
                        var v48 int64 = v28
                        if true {
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                        } else {
                            v49 := GenS4{F0: fi46, F1: "\n\t", F2: 2}
                            _ = v49
                        }
                        v3 = append(v3, 301)
                        _ = fi47
                        _ = v48
                    }
                case -1:
                    if (3.14159 < (0.1 / float32((0.0 - -2.25)))) {
                        for fi50 := 0; fi50 < 14; fi50++ {
                            v28 = (v28 - (v28 ^ v28))
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi50
                        }
                    }
                }
                for fi51 := 0; fi51 < 11; fi51++ {
                    if true {
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v52 uint = (v2 & 4294967295)
                            v52 -= (v45 % (((v2 + v45)) | 1))
                            switch "foo bar" {
                            case "":
                                if ((-(fi46)) >= fi51) {
                                    v23 += (^(v23))
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    fmt.Println(fi51, v45)
                                    switch (fi51 + (-(fi51))) {
                                    case 32768:
                                        var v53 uint8 = v23
                                        _ = v53
                                    case -17:
                                        var v54 int64 = 9223372036854775807
                                        _ = v54
                                    case 42:
                                        if len(v3) > 0 {
                                            v3[0] = (v4 + v4)
                                        }
                                    default:
                                        var v55 int32 = 949
                                        _ = v55
                                    }
                                } else {
                                    v56 := GenS2{F0: uint8(v45), F1: 0}
                                    if len(v3) > 0 {
                                        v3[0] = uint32(v23)
                                    }
                                    v57 := GenS5{F0: true, F1: true}
                                    var v58 bool = (!false)
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                            case "a":
                                v23 -= v23
                            default:
                                if ((v4 * (v4 | (v4 % ((885) | 1)))) >= v4) {
                                    if (v4 >= (v4 + uint32(v45))) {
                                        v59 := GenS4{F0: (fi46 / (((fi46 | (fi46 + int(v28)))) | 1)), F1: "a", F2: 223}
                                        if len(v3) > 0 {
                                            v3[0] = v4
                                        }
                                        _ = len(v3)
                                        var v60 int32 = 0
                                        if (-2.25 != 0.0) {
                                            var v61 int32 = 504
                                            if true {
                                                v62 := GenS5{F0: false, F1: (0 >= 0)}
                                                _ = v62
                                            } else {
                                                if (212 != uint8(v4)) {
                                                    v61 -= 2147483647
                                                    var v63 uint16 = 255
                                                    _ = v63
                                                }
                                                if (((42 != 1) || (!false)) || (v4 < v4)) {
                                                    var v64 uint = (v2 >> ((790) & 63))
                                                    _ = v64
                                                }
                                            }
                                            v3 = append(v3, 65535)
                                            _ = v61
                                        } else {
                                            if false {
                                                v60 = (v60 * (v60 % ((v60) | 1)))
                                                fmt.Println(v60, v60, v45)
                                            } else {
                                                var v65 uint8 = 157
                                                _ = v65
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        var v66 int8 = 100
                                        v66 += (-(v66))
                                        _ = len(v3)
                                        fmt.Println(fi46, v4, v2)
                                        _ = v59
                                        _ = v60
                                        _ = v66
                                    }
                                    v67 := GenS6{F0: (fi46 + (-(fi46))), F1: (v23 - (v23 | 149))}
                                    var v68 float64 = ((float64(1e10) * 0.0) / 1.0)
                                    _ = v67
                                    _ = v68
                                } else {
                                    var v69 int8 = 1
                                    if (!(!false)) {
                                        v28 = -9223372036854775808
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        if ((v4 * v4) > (v4 + 2147483648)) {
                                            switch 933 {
                                            case 32767:
                                                if true {
                                                    v4 -= (v4 - 2)
                                                    v3 = append(v3, v4)
                                                    _ = len(v3)
                                                    if false {
                                                        v45 = v45
                                                        var v70 int16 = -348
                                                        if (fi46 > (fi51 + int(v69))) {
                                                            v70 -= v70
                                                        }
                                                        var v71 int16 = (v70 + int16(v4))
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                            case 1:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case -223:
                                                var v72 string = "a"
                                                _ = v72
                                            }
                                            v73 := GenS4{F0: (fi46 / (((fi51 >> ((v45) & 63))) | 1)), F1: ("hello" + "foo bar"), F2: 1}
                                            switch (fi51 << (((v45 & (v45 + (v2 - (v45 ^ (v1 << ((v2) & 63))))))) & 63)) {
                                            case 32767:
                                                if true {
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(v4))
                                                    }
                                                    v23 -= (v23 * 128)
                                                } else {
                                                    v74 := GenS6{F0: (^(fi46)), F1: 1}
                                                    if (((("a" + ("" + "hello")) + (("\n\t" + "") + "a")) + "a") < (("hello" + "\n\t") + "foo bar")) {
                                                        var v75 int64 = (v28 % ((int64(v4)) | 1))
                                                        fmt.Println(v52, v69)
                                                        v69 += v69
                                                        _ = v75
                                                    } else {
                                                        v76 := GenS7{F0: "", F1: -1, F2: uint32(v1)}
                                                        v77 := GenS5{F0: (-505 >= int(v2)), F1: (!false)}
                                                        if len(v3) > 0 {
                                                            v3[0] = 175
                                                        }
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    _ = v74
                                                }
                                            }
                                            _ = v73
                                        }
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v45 = v52
                                    }
                                    fmt.Println(v52)
                                    _ = v69
                                }
                            }
                            _ = v52
                        case "\n\t":
                            var v78 int32 = -1
                            _ = v78
                        case "a":
                            v79 := GenS8{F0: fi46, F1: (32767 > 1)}
                            _ = v79
                        }
                        if true {
                            v3 = append(v3, 2147483648)
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                            v28 += (v28 << ((255) & 63))
                        } else {
                            switch "a" {
                            case "\n\t":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            default:
                                if false {
                                    var v80 uint = 42
                                    switch "foo bar" {
                                    case "":
                                        _ = len(v3)
                                        fmt.Println(v45)
                                        fmt.Println(v2)
                                    case "hello":
                                        fmt.Println(v23, v1)
                                    case "foo bar":
                                        v81 := GenS8{F0: (fi51 % ((int(v23)) | 1)), F1: false}
                                        switch (fi51 & -2147483648) {
                                        case 2147483647:
                                            if false {
                                                if (uint32(v2) < 828) {
                                                    switch (("a" + ("hello" + "")) + "hello") {
                                                    case "\n\t":
                                                        if false {
                                                            switch (fi46 >> (((v80 + v45)) & 63)) {
                                                            case 939:
                                                                switch (-(fi46)) {
                                                                case -597:
                                                                    v3 = append(v3, (v4 >> (((v45 & v80)) & 63)))
                                                                    fmt.Println(v45, v4, fi51)
                                                                    v2 += v45
                                                                case -200:
                                                                    var v82 int8 = 100
                                                                    _ = v82
                                                                case 9223372036854775807:
                                                                    var v83 uint = 4294967295
                                                                    _ = v83
                                                                }
                                                            case 68:
                                                                switch ("foo bar" + ("\n\t" + "a")) {
                                                                case "hello":
                                                                    fmt.Println(fi46, v45)
                                                                case "a":
                                                                    if ((int16(v4) == -32768) || (!(65535 < 42))) {
                                                                        v84 := GenS9{F0: 2147483647}
                                                                        v85 := GenS4{F0: int(v4), F1: "foo bar", F2: 255}
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                default:
                                                                    v86 := GenS5{F0: true, F1: ((1 > 553) || (42 >= 0))}
                                                                    _ = v86
                                                                }
                                                            }
                                                            var v87 float64 = 1.5
                                                            v88 := GenS10{F0: v23}
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        var v89 float64 = 1.5
                                                        v80 -= 42
                                                        _ = v89
                                                    case "hello":
                                                        fmt.Println(v1, v4, fi51)
                                                    }
                                                    v28 += -1
                                                    var v90 float32 = 0.1
                                                    _ = v90
                                                } else {
                                                    if true {
                                                        fmt.Println(v23)
                                                        v91 := GenS6{F0: fi51, F1: v23}
                                                        _ = v91
                                                    }
                                                    switch "a" {
                                                    case "hello":
                                                        var v92 bool = true
                                                        _ = v92
                                                    case "\n\t":
                                                        fmt.Println(v23, fi46, fi46)
                                                    case "a":
                                                        v4 = 2147483648
                                                    default:
                                                        v93 := GenS5{F0: (0 < v23), F1: false}
                                                        _ = v93
                                                    }
                                                }
                                                v45 = (v45 ^ v80)
                                                v80 += (v1 >> ((255) & 63))
                                            }
                                        case 1:
                                            v28 += v28
                                        default:
                                            var v94 uint8 = (^(v23))
                                            _ = v94
                                        }
                                        _ = v81
                                    }
                                    _ = v80
                                } else {
                                    v95 := GenS11{F0: ((fi51 - (fi46 * int(v1))) < (fi51 | fi51)), F1: (^(fi51))}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v28 -= v28
                                    _ = v95
                                }
                            }
                        }
                        _ = len(v3)
                    } else {
                        var v96 int16 = 32767
                        _ = v96
                    }
                    _ = fi51
                }
                _ = fi46
            }
            _ = v45
        } else {
            v1 -= (^(v1))
            if true {
                v1 = (v2 - v1)
                if false {
                    v23 = (v23 + (^(v23)))
                } else {
                    if (!false) {
                        if ((!((0.1 * 0.0) >= 0.1)) && true) {
                            v1 -= (v1 * (^(v1)))
                            switch -9223372036854775808 {
                            case 2147483647:
                                switch 127 {
                                case -1:
                                    v97 := GenS12{F0: 832}
                                    _ = v97
                                case 32767:
                                    if ((v1 * v2) == (v1 - (v1 >> ((uint(v4)) & 63)))) {
                                        v98 := GenS13{F0: -1, F1: (^(v2)), F2: v1}
                                        var v99 int16 = -674
                                        _ = v98
                                        _ = v99
                                    } else {
                                        fmt.Println(v4, v1, v4)
                                        var v100 float64 = 0.1
                                        if false {
                                            switch "foo bar" {
                                            case "foo bar":
                                                fmt.Println(v100, v28, v28)
                                            case "\n\t":
                                                v4 = v4
                                            case "a":
                                                var v101 uint64 = 42
                                                _ = v101
                                            }
                                            v102 := GenS3{F0: 893}
                                            _ = v102
                                        }
                                        _ = v100
                                    }
                                default:
                                    var v103 int16 = 1000
                                    _ = v103
                                }
                            }
                            for fi104 := 0; fi104 < 1; fi104++ {
                                v105 := GenS14{F0: 2}
                                if len(v3) > 0 {
                                    v3[0] = uint32(v1)
                                }
                                v4 = v4
                                _ = fi104
                                _ = v105
                            }
                        }
                        for fi106 := 0; fi106 < 9; fi106++ {
                            switch 545 {
                            case 127:
                                if (42 <= -1) {
                                    var v107 int64 = (^(v28))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v28)
                                    _ = v107
                                } else {
                                    switch 32767 {
                                    case 308:
                                        for fi108 := 0; fi108 < 18; fi108++ {
                                            _ = len(v3)
                                            switch "" {
                                            case "":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "foo bar":
                                                v23 -= (^(v23))
                                            case "\n\t":
                                                switch "foo bar" {
                                                case "hello":
                                                    v3 = append(v3, v4)
                                                default:
                                                    v4 -= 2147483647
                                                }
                                            }
                                            v23 = 0
                                            _ = fi108
                                        }
                                    }
                                }
                            default:
                                v109 := GenS9{F0: 492}
                                _ = v109
                            }
                            v110 := GenS15{F0: (v23 | v23)}
                            _ = fi106
                            _ = v110
                        }
                        var v111 string = ""
                        _ = v111
                    }
                    v4 = uint32(v2)
                }
                v1 += v1
            } else {
                var v112 int32 = 171
                var v113 int = -676
                var v114 uint8 = (v23 / ((v23) | 1))
                _ = v112
                _ = v113
                _ = v114
            }
            for fi115 := 0; fi115 < 11; fi115++ {
                if ((0.0 - -2.25) < float32(1e10)) {
                    if len(v3) > 0 {
                        v3[0] = (^(v4))
                    }
                    switch "" {
                    case "":
                        var v116 int64 = (-(v28))
                        _ = v116
                    case "foo bar":
                        for fi117 := 0; fi117 < 3; fi117++ {
                            var v118 string = "foo bar"
                            var v119 uint8 = v23
                            var v120 uint64 = 42
                            v4 -= (^(v4))
                            v121 := GenS16{F0: (v28 % (((v28 >> ((42) & 63))) | 1)), F1: true}
                            var v122 uint64 = (^(v120))
                            _ = fi117
                            _ = v118
                            _ = v119
                            _ = v120
                            _ = v121
                            _ = v122
                        }
                    case "hello":
                        var v123 uint64 = 1
                        _ = v123
                    }
                    switch (("hello" + "") + ("" + ("foo bar" + "\n\t"))) {
                    case "a":
                        for fi124 := 0; fi124 < 19; fi124++ {
                            if true {
                                var v125 int8 = -1
                                _ = len(v3)
                                if true {
                                    v4 = (v4 + (v4 | (v4 & 2)))
                                    v126 := GenS17{F0: genFn0(-1, 0, v23), F1: v2}
                                    _ = v126
                                } else {
                                    v4 -= (v4 - (v4 - uint32(v23)))
                                    v23 = v23
                                    if (452 > v2) {
                                        switch fi115 {
                                        case 578:
                                            var v127 string = ("" + "a")
                                            _ = v127
                                        }
                                        v2 += v2
                                        v28 -= v28
                                        v4 += uint32(v23)
                                        v23 = 226
                                    } else {
                                        v1 -= (v1 & v1)
                                        v3 = append(v3, (^(v4)))
                                    }
                                }
                                _ = v125
                            }
                            _ = fi124
                        }
                    case "":
                        v128 := GenS6{F0: 2147483647, F1: (v23 + 205)}
                        _ = v128
                    case "foo bar":
                        for fi129 := 0; fi129 < 6; fi129++ {
                            if false {
                                v28 += v28
                                v130 := GenS1{F0: 0, F1: 2147483647}
                                _ = v130
                            } else {
                                v28 += v28
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = len(v3)
                                if len(v3) > 0 {
                                    v3[0] = 118
                                }
                            }
                            if len(v3) > 0 {
                                v3[0] = v4
                            }
                            _ = fi129
                        }
                    default:
                        for fi131 := 0; fi131 < 7; fi131++ {
                            var v132 int64 = 0
                            var v133 int32 = 42
                            _ = fi131
                            _ = v132
                            _ = v133
                        }
                    }
                }
                _ = fi115
            }
        }
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        v134 := GenS18{F0: -82, F1: 1}
        _ = v134
    case "\n\t":
        switch "a" {
        case "hello":
            v135 := GenS17{F0: -128, F1: v2}
            _ = v135
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
