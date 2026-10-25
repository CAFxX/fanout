package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int; F1 uint; F2 int8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint; F1 string; F2 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 uint32; F1 uint32; F2 int }
type GenS11 struct { F0 uint64; F1 int }
type GenS12 struct { F0 string; F1 uint; F2 uint }
type GenS13 struct { F0 int8; F1 uint16; F2 int64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int32; F1 bool; F2 int32 }
type GenS16 struct { F0 bool; F1 int8; F2 uint64 }
type GenS17 struct { F0 uint16 }
type GenS18 struct { F0 int64; F1 uint32 }
type GenS19 struct { F0 int64; F1 uint32 }
type GenS20 struct { F0 uint64 }
func genFn0() uint8 {
    var v5 int8 = 0
    v5 = (-(v5))
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{}
        var v8 uint64 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    fmt.Println(v5)
    if true {
        v9 := GenS1{F0: 819, F1: 914}
        if true {
            switch ("a" + "hello") {
            case "a":
                switch ("a" + "") {
                case "foo bar":
                    if ((!false) || false) {
                        if (!((3.14159 >= -2.25) && true)) {
                            var v10 int8 = 0
                            var v11 uint64 = 632
                            _ = v10
                            _ = v11
                        } else {
                            v5 = v5
                        }
                        for fi12 := 0; fi12 < 4; fi12++ {
                            switch (("hello" + "") + "a") {
                            case "foo bar":
                                fmt.Println(v5, fi12)
                            }
                            _ = fi12
                        }
                        v13 := GenS1{F0: 32767, F1: 1}
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 42, F1: 18446744073709551615}
                        _ = v14
                    }
                case "hello":
                    v5 = -1
                case "":
                    if true {
                        v15 := []int{-325, 32767, 127}
                        _ = v15
                    } else {
                        var v16 int8 = (v5 >> ((0) & 63))
                        var v17 uint = 2
                        _ = v16
                        _ = v17
                    }
                default:
                    switch 9223372036854775807 {
                    case 1:
                        v18 := GenS2{F0: 429, F1: 4294967295, F2: -11}
                        _ = v18
                    default:
                        if ((-39 < v5) && (((4294967295 != 248) && true) || ("\n\t" <= "\n\t"))) {
                            for fi19 := 0; fi19 < 20; fi19++ {
                                v20 := []int{(fi19 - fi19)}
                                _ = fi19
                                _ = v20
                            }
                            v21 := []int{}
                            _ = v21
                        }
                    }
                }
            case "foo bar":
                fmt.Println(v5)
            }
        } else {
            var v22 uint = 559
            v5 += -128
            _ = v22
        }
        _ = v9
    } else {
        v23 := []int{}
        var v24 int64 = -5
        if (-748 == -498) {
            switch (("\n\t" + "a") + "hello") {
            case "\n\t":
                switch -2147483648 {
                case -9223372036854775808:
                    switch ("\n\t" + ("a" + "foo bar")) {
                    case "a":
                        var v25 int = -860
                        _ = v25
                    default:
                        switch -9223372036854775808 {
                        case 293:
                            switch -651 {
                            case -1:
                                if true {
                                    if (128 > 128) {
                                        var v26 float64 = (1e10 + 0.1)
                                        for fi27 := 0; fi27 < 16; fi27++ {
                                            for fi28 := 0; fi28 < 11; fi28++ {
                                                var v29 uint = 0
                                                if (3.14159 == (float32(1.5) * 0.1)) {
                                                    if false {
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                        fmt.Println(v24, v26, v29)
                                                    }
                                                } else {
                                                    v30 := GenS3{F0: 42, F1: 100}
                                                    v24 -= 0
                                                    _ = v30
                                                }
                                                _ = fi28
                                                _ = v29
                                            }
                                            v5 = v5
                                            _ = fi27
                                        }
                                        _ = v26
                                    }
                                    for fi31 := 0; fi31 < 6; fi31++ {
                                        for fi32 := 0; fi32 < 8; fi32++ {
                                            v24 += v24
                                            _ = fi32
                                        }
                                        switch (-(fi31)) {
                                        case 127:
                                            var v33 string = ""
                                            _ = v33
                                        case 32767:
                                            v24 = 9223372036854775807
                                        case 32768:
                                            if (!((-2.25 - 1.5) > 1.5)) {
                                                switch "hello" {
                                                case "a":
                                                    fmt.Println(fi31, v5)
                                                case "":
                                                    fmt.Println(fi31)
                                                case "hello":
                                                    v34 := GenS4{F0: 65535, F1: ("" + ""), F2: (!(!false))}
                                                    _ = v34
                                                default:
                                                    if ((!(42 > 42)) || false) {
                                                        v23 = append(v23, (fi31 % (((fi31 / (((fi31 | (fi31 >> ((911) & 63)))) | 1))) | 1)))
                                                        if (false && ((1.5 + 1.5) != (0.0 / 0.1))) {
                                                            if ((128 == 38) || false) {
                                                                if ("foo bar" > ((("a" + (("\n\t" + "") + "\n\t")) + "a") + ("\n\t" + "hello"))) {
                                                                    fmt.Println(v24)
                                                                    for fi35 := 0; fi35 < 19; fi35++ {
                                                                        var v36 float64 = 1.5
                                                                        v24 += (v24 ^ 1)
                                                                        _ = fi35
                                                                        _ = v36
                                                                    }
                                                                    for fi37 := 0; fi37 < 7; fi37++ {
                                                                        fmt.Println(v5)
                                                                        var v38 string = ""
                                                                        _ = fi37
                                                                        _ = v38
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v39 int16 = 42
                                                                        _ = v39
                                                                    }
                                                                    for fi40 := 0; fi40 < 4; fi40++ {
                                                                        var v41 int8 = (v5 & 42)
                                                                        _ = fi40
                                                                        _ = v41
                                                                    }
                                                                }
                                                                v42 := GenS2{F0: fi31, F1: 65535, F2: (v5 * (v5 / ((v5) | 1)))}
                                                                _ = v42
                                                            } else {
                                                                v43 := GenS5{F0: 2}
                                                                _ = v43
                                                            }
                                                        } else {
                                                            for fi44 := 0; fi44 < 3; fi44++ {
                                                                if (1 == 2) {
                                                                    fmt.Println(fi31, fi31, fi44)
                                                                }
                                                                if true {
                                                                    v45 := GenS2{F0: fi31, F1: 92, F2: (v5 % ((v5) | 1))}
                                                                    _ = v45
                                                                }
                                                                var v46 uint8 = 2
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v46 += (v46 + 255)
                                                                case "hello":
                                                                    v47 := GenS2{F0: (fi44 / (((fi44 << ((65535) & 63))) | 1)), F1: 210, F2: (v5 << ((2) & 63))}
                                                                    _ = v47
                                                                case "a":
                                                                    if false {
                                                                        if (1e10 < 0.1) {
                                                                            if true {
                                                                                var v48 float32 = 0.0
                                                                                if len(v23) > 0 {
                                                                                    v23[0] = fi31
                                                                                }
                                                                                _ = v48
                                                                            } else {
                                                                                v24 = v24
                                                                                v49 := GenS6{F0: 2, F1: (true || true)}
                                                                                switch fi31 {
                                                                                case 1:
                                                                                    var v50 bool = (true || (false && (-2.25 < (0.1 + float64((((-2.25 * 3.14159) - 1e10) + -2.25))))))
                                                                                    _ = v50
                                                                                case -263:
                                                                                    v51 := GenS5{F0: 1}
                                                                                    _ = v51
                                                                                }
                                                                                _ = v49
                                                                            }
                                                                        } else {
                                                                            v52 := GenS7{F0: 2, F1: "foo bar"}
                                                                            fmt.Println(fi31)
                                                                            if (1 != 2) {
                                                                                v53 := GenS6{F0: 42, F1: true}
                                                                                var v54 uint16 = 2
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                if len(v23) > 0 {
                                                                                    v23[0] = (fi31 + -478)
                                                                                }
                                                                            }
                                                                            _ = v52
                                                                        }
                                                                        v46 -= (v46 & 255)
                                                                        if (42 != (fi31 % ((fi31) | 1))) {
                                                                            if (32767 > 32767) {
                                                                                if len(v23) >= 2 {
                                                                                    v23 = v23[1:len(v23)]
                                                                                }
                                                                            } else {
                                                                                v55 := GenS8{F0: ("foo bar" + "a")}
                                                                                v56 := GenS3{F0: v46, F1: v5}
                                                                                _ = v55
                                                                                _ = v56
                                                                            }
                                                                            _ = len(v23)
                                                                            if (!(-2.25 < 0.1)) {
                                                                                v23 = append(v23, 127)
                                                                            } else {
                                                                                if len(v23) >= 2 {
                                                                                    v23 = v23[1:len(v23)]
                                                                                }
                                                                                if true {
                                                                                    if true {
                                                                                        v57 := GenS9{F0: -1, F1: 65535}
                                                                                        var v58 uint = 255
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    }
                                                                                    if true {
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v59 float64 = (-2.25 - 0.0)
                                                                                            _ = v59
                                                                                        case "hello":
                                                                                            v46 = v46
                                                                                        case "foo bar":
                                                                                            v46 = v46
                                                                                        default:
                                                                                            if len(v23) > 0 {
                                                                                                v23[0] = (fi31 & fi31)
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        var v60 uint = 310
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            switch ("hello" + ("" + ((("\n\t" + "foo bar") + "") + "\n\t"))) {
                                                                                            case "hello":
                                                                                                fmt.Println(v5, v24, v60)
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v61 uint64 = 42
                                                                                            _ = v61
                                                                                        }
                                                                                        _ = v60
                                                                                    }
                                                                                    fmt.Println(fi44, v46)
                                                                                    if len(v23) >= 2 {
                                                                                        v23 = v23[1:len(v23)]
                                                                                    }
                                                                                    v23 = append(v23, fi31)
                                                                                }
                                                                            }
                                                                            if len(v23) >= 2 {
                                                                                v23 = v23[1:len(v23)]
                                                                            }
                                                                            v24 -= (v24 >> ((4294967295) & 63))
                                                                        } else {
                                                                            var v62 int64 = v24
                                                                            v46 = (v46 / (((v46 + (v46 << ((65535) & 63)))) | 1))
                                                                            _ = v62
                                                                        }
                                                                        v5 = v5
                                                                        var v63 int64 = v24
                                                                        var v64 uint32 = 2147483647
                                                                        _ = v63
                                                                        _ = v64
                                                                    } else {
                                                                        v24 -= -1
                                                                        v23 = append(v23, (fi44 ^ -9223372036854775808))
                                                                    }
                                                                default:
                                                                    var v65 uint16 = 2
                                                                    v23 = append(v23, (-(fi31)))
                                                                    _ = v65
                                                                }
                                                                var v66 bool = false
                                                                _ = fi44
                                                                _ = v46
                                                                _ = v66
                                                            }
                                                        }
                                                        v24 -= v24
                                                    }
                                                }
                                                for fi67 := 0; fi67 < 11; fi67++ {
                                                    v68 := GenS2{F0: fi67, F1: 1, F2: 27}
                                                    fmt.Println(v5)
                                                    _ = fi67
                                                    _ = v68
                                                }
                                                var v69 float32 = ((-2.25 / (0.0 - 1e10)) + (3.14159 + 0.0))
                                                _ = v69
                                            } else {
                                                if (42 < 1) {
                                                    v5 = v5
                                                    v5 += 1
                                                    _ = len(v23)
                                                } else {
                                                    var v70 int8 = 42
                                                    v70 = -1
                                                    fmt.Println(fi31, v5, v24)
                                                    _ = v70
                                                }
                                                v24 = v24
                                            }
                                        }
                                        _ = fi31
                                    }
                                }
                            default:
                                if ((2147483647 <= 1) || true) {
                                    if len(v23) > 0 {
                                        v23[0] = 0
                                    }
                                    fmt.Println(v24, v5)
                                    fmt.Println(v24, v24)
                                } else {
                                    var v71 uint8 = 1
                                    if false {
                                        v72 := GenS10{F0: 375, F1: 4294967295, F2: 0}
                                        _ = v72
                                    } else {
                                        v73 := GenS4{F0: 31, F1: ("" + ((("" + "") + ("hello" + "foo bar")) + "\n\t")), F2: true}
                                        _ = v73
                                    }
                                    switch -492 {
                                    case 32768:
                                        if (!false) {
                                            var v74 uint = 300
                                            if (false || false) {
                                                switch 534 {
                                                case 42:
                                                    var v75 int8 = -41
                                                    _ = v75
                                                default:
                                                    v23 = append(v23, -51)
                                                }
                                            }
                                            _ = v74
                                        }
                                    case 128:
                                        var v76 uint64 = 4294967295
                                        _ = v76
                                    default:
                                        for fi77 := 0; fi77 < 15; fi77++ {
                                            v24 += -9223372036854775808
                                            v78 := GenS11{F0: 42, F1: 807}
                                            for fi79 := 0; fi79 < 15; fi79++ {
                                                v5 = (v5 + v5)
                                                switch (fi79 * -2147483648) {
                                                case 2147483647:
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                case 42:
                                                    switch (fi79 + (fi77 ^ 916)) {
                                                    case 127:
                                                        fmt.Println(fi77, v5, fi77)
                                                    case 32767:
                                                        var v80 uint32 = 2147483647
                                                        _ = v80
                                                    case 32768:
                                                        var v81 uint16 = 755
                                                        _ = v81
                                                    }
                                                default:
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                }
                                                _ = fi79
                                            }
                                            _ = fi77
                                            _ = v78
                                        }
                                    }
                                    _ = v71
                                }
                            }
                        default:
                            if false {
                                fmt.Println(v5, v24)
                            } else {
                                if true {
                                    if true {
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                        if len(v23) >= 2 {
                                            v23 = v23[1:len(v23)]
                                        }
                                    }
                                }
                                if ((1 >= 220) || (255 >= 255)) {
                                    var v82 int16 = 717
                                    _ = v82
                                }
                                for fi83 := 0; fi83 < 18; fi83++ {
                                    fmt.Println(v24, v24)
                                    for fi84 := 0; fi84 < 14; fi84++ {
                                        if true {
                                            v85 := GenS12{F0: ("foo bar" + "\n\t"), F1: 42, F2: 4294967295}
                                            if false {
                                                switch 528 {
                                                case -76:
                                                    fmt.Println(v24)
                                                }
                                                _ = len(v23)
                                                if (!((438 > 4294967295) || true)) {
                                                    v86 := GenS4{F0: 255, F1: (("foo bar" + "hello") + "hello"), F2: (!(v5 != 0))}
                                                    v5 = (v5 / ((v5) | 1))
                                                    if (-2.25 < float32(((1e10 * 1e10) - 0.0))) {
                                                        if len(v23) >= 2 {
                                                            v23 = v23[1:len(v23)]
                                                        }
                                                    }
                                                    _ = v86
                                                }
                                                v24 = (v24 << ((18446744073709551615) & 63))
                                                v24 += (v24 * v24)
                                            } else {
                                                var v87 uint = 362
                                                v87 -= (v87 >> ((42) & 63))
                                                _ = v87
                                            }
                                            _ = v85
                                        }
                                        _ = fi84
                                    }
                                    v24 = (-(v24))
                                    _ = fi83
                                }
                            }
                        }
                    }
                default:
                    v88 := GenS9{F0: -9223372036854775808, F1: 4294967295}
                    _ = v88
                }
            default:
                for fi89 := 0; fi89 < 5; fi89++ {
                    var v90 uint32 = 381
                    _ = fi89
                    _ = v90
                }
            }
            v91 := GenS12{F0: "", F1: 65535, F2: 1}
            fmt.Println(v5, v24)
            _ = v91
        } else {
            switch "\n\t" {
            case "a":
                for fi92 := 0; fi92 < 9; fi92++ {
                    fmt.Println(v5, v5)
                    _ = fi92
                }
            case "hello":
                v24 += -1
            }
            var v93 float32 = 3.14159
            for fi94 := 0; fi94 < 1; fi94++ {
                v24 = v24
                _ = len(v23)
                _ = fi94
            }
            _ = v93
        }
        _ = v23
        _ = v24
    }
    _ = v5
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{42, 2, 2147483648}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v95 uint16 = 211
    var v96 int64 = -9223372036854775808
    var v97 uint32 = 709
    fmt.Println(v95, v97, v96)
    v95 -= v95
    v98 := GenS8{F0: (("\n\t" + ("" + "foo bar")) + "a")}
    switch "foo bar" {
    case "":
        v96 += (v96 << ((1) & 63))
    case "hello":
        delete(v4, 371)
    case "a":
        var v99 uint64 = 112
        _ = v99
    }
    if false {
        v95 = (v95 >> ((65535) & 63))
    }
    switch -9223372036854775808 {
    case 127:
        var v100 uint16 = (v95 & 0)
        _ = v100
    case 32768:
        var v101 int8 = 1
        _ = v101
    case 42:
        v102 := GenS7{F0: v95, F1: ""}
        _ = v102
    default:
        if false {
            fmt.Println(v97, v97, v97)
            v97 -= (^(v97))
            for fi103 := 0; fi103 < 16; fi103++ {
                var v104 uint32 = v97
                _ = fi103
                _ = v104
            }
        } else {
            if (false || true) {
                var v105 int64 = (v96 | v96)
                for fi106 := 0; fi106 < 3; fi106++ {
                    v107 := GenS10{F0: (v97 + 2147483647), F1: v97, F2: fi106}
                    _ = fi106
                    _ = v107
                }
                _ = v105
            } else {
                switch "\n\t" {
                case "":
                    var v108 uint8 = 127
                    _ = v108
                case "foo bar":
                    switch -773 {
                    case 9223372036854775807:
                        v109 := GenS13{F0: 1, F1: (v95 % (((^(v95))) | 1)), F2: v96}
                        _ = v109
                    case 42:
                        var v110 float32 = 0.0
                        _ = v110
                    case 533:
                        v97 += v97
                    default:
                        for fi111 := 0; fi111 < 3; fi111++ {
                            if true {
                                for fi112 := 0; fi112 < 15; fi112++ {
                                    var v113 uint = 2
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v113 = (v113 - v113)
                                    case "a":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "":
                                        v96 = (v96 + (v96 / ((v96) | 1)))
                                    default:
                                        var v114 int8 = 42
                                        _ = v114
                                    }
                                    fmt.Println(fi111, v95, v95)
                                    _ = fi112
                                    _ = v113
                                }
                                for fi115 := 0; fi115 < 2; fi115++ {
                                    if (true && (v97 >= (^(v97)))) {
                                        if (127 <= -1) {
                                            var v116 uint16 = v95
                                            _ = v116
                                        } else {
                                            v117 := GenS1{F0: fi115, F1: 1}
                                            if (!false) {
                                                v95 -= v95
                                            } else {
                                                v96 = v96
                                            }
                                            _ = v117
                                        }
                                        v118 := GenS9{F0: v96, F1: v97}
                                        if (int(v95) != fi111) {
                                            var v119 int32 = -558
                                            var v120 uint32 = (v97 / (((^(v97))) | 1))
                                            _ = v119
                                            _ = v120
                                        }
                                        _ = v118
                                    } else {
                                        var v121 float32 = 1.5
                                        fmt.Println(v97, v95)
                                        switch ("a" + "a") {
                                        case "a":
                                            v97 = (^(v97))
                                        case "hello":
                                            v122 := GenS9{F0: (-(v96)), F1: 65535}
                                            _ = v122
                                        case "foo bar":
                                            var v123 float32 = (1e10 * (v121 - (v121 - 0.0)))
                                            _ = v123
                                        default:
                                            switch (fi115 << ((0) & 63)) {
                                            case 4294967295:
                                                v124 := GenS4{F0: 1, F1: "hello", F2: false}
                                                _ = v124
                                            case 127:
                                                v125 := v1[(v97 ^ (v97 - (v97 & (v97 % ((v97) | 1)))))]
                                                _ = v125
                                            }
                                        }
                                        switch "hello" {
                                        case "":
                                            var v126 int32 = 2147483647
                                            _ = v126
                                        case "\n\t":
                                            v127 := GenS1{F0: (fi115 / (((fi115 | 32768)) | 1)), F1: 4294967295}
                                            _ = v127
                                        case "hello":
                                            var v128 int64 = 1
                                            _ = v128
                                        default:
                                            var v129 int = fi115
                                            _ = v129
                                        }
                                        _ = v121
                                    }
                                    if (genFn0() > 110) {
                                        var v130 float64 = 0.1
                                        v3[74] = false
                                        v131 := GenS14{F0: ""}
                                        var v132 float32 = 0.1
                                        if true {
                                            var v133 int8 = -1
                                            _ = v133
                                        }
                                        _ = v130
                                        _ = v131
                                        _ = v132
                                    } else {
                                        v1[(v97 & (v97 + (v97 / (((v97 >> ((200) & 63))) | 1))))] = v97
                                        if true {
                                            switch 833 {
                                            case 168:
                                                v96 = (v96 & (v96 % ((v96) | 1)))
                                            case 686:
                                                switch "a" {
                                                case "foo bar":
                                                    fmt.Println(v95, fi111, fi111)
                                                default:
                                                    v134 := GenS15{F0: 2147483647, F1: (v97 == (v97 * (v97 - (v97 ^ (v97 >> ((65535) & 63)))))), F2: -2147483648}
                                                    _ = v134
                                                }
                                            case 42:
                                                v135 := GenS16{F0: false, F1: 93, F2: 4294967295}
                                                _ = v135
                                            default:
                                                v136 := GenS17{F0: v95}
                                                _ = v136
                                            }
                                            if false {
                                                var v137 int = -9223372036854775808
                                                fmt.Println(v96, fi115, v96)
                                                _ = v137
                                            }
                                        } else {
                                            var v138 float64 = float64((1.5 * 1e10))
                                            var v139 uint64 = 753
                                            switch fi111 {
                                            case -496:
                                                var v140 int64 = (v96 >> ((65535) & 63))
                                                switch "foo bar" {
                                                case "":
                                                    if false {
                                                        var v141 int = (fi115 + (fi111 << ((65535) & 63)))
                                                        v141 = fi115
                                                        _ = v141
                                                    } else {
                                                        if false {
                                                            fmt.Println(v138)
                                                            v138 = v138
                                                            var v142 int = (^(fi115))
                                                            _ = v142
                                                        } else {
                                                            v97 += v97
                                                            v96 += (v140 >> ((255) & 63))
                                                        }
                                                        v97 -= (v97 | (v97 ^ v97))
                                                    }
                                                case "\n\t":
                                                    var v143 uint8 = 1
                                                    _ = v143
                                                case "foo bar":
                                                    delete(v4, v97)
                                                default:
                                                    var v144 int = (^(fi115))
                                                    if false {
                                                        if (float32(v138) >= 1.5) {
                                                            switch (("" + "") + "foo bar") {
                                                            case "a":
                                                                switch ("a" + "\n\t") {
                                                                case "foo bar":
                                                                    v96 = v140
                                                                }
                                                            case "hello":
                                                                delete(v4, (v97 | (v97 | v97)))
                                                            default:
                                                                var v145 int8 = 127
                                                                _ = v145
                                                            }
                                                            if false {
                                                                v146 := GenS8{F0: ""}
                                                                _ = v146
                                                            }
                                                        } else {
                                                            v139 = 912
                                                            switch "foo bar" {
                                                            case "hello":
                                                                if false {
                                                                    v97 = uint32(v95)
                                                                    v138 += ((v138 - v138) - (v138 - v138))
                                                                    v138 = v138
                                                                } else {
                                                                    var v147 int8 = -45
                                                                    _ = v147
                                                                }
                                                            case "\n\t":
                                                                v97 = 0
                                                            case "a":
                                                                var v148 uint16 = (^(v95))
                                                                _ = v148
                                                            }
                                                            var v149 uint8 = 128
                                                            _ = v149
                                                        }
                                                        if len(v2) > 0 {
                                                            v2[0] = (v97 / (((v97 / (((^(v97))) | 1))) | 1))
                                                        }
                                                        fmt.Println(v144, v140)
                                                    } else {
                                                        if (!false) {
                                                            v1[(v97 ^ v97)] = (v97 * v97)
                                                        }
                                                        if true {
                                                            v150 := GenS18{F0: v96, F1: v97}
                                                            var v151 float64 = v138
                                                            _ = v150
                                                            _ = v151
                                                        } else {
                                                            v140 -= v140
                                                            v97 = (v97 | (v97 + v97))
                                                            delete(v1, (v97 - (v97 / ((uint32(v95)) | 1))))
                                                            switch ("hello" + (("" + "foo bar") + ((("foo bar" + "\n\t") + "a") + "hello"))) {
                                                            case "hello":
                                                                if false {
                                                                    v97 -= uint32(v95)
                                                                    var v152 int = (fi111 << ((65535) & 63))
                                                                    _ = v152
                                                                }
                                                            case "\n\t":
                                                                v139 = (^(v139))
                                                            case "foo bar":
                                                                v97 = 115
                                                            default:
                                                                if false {
                                                                    if (true && true) {
                                                                        if false {
                                                                            fmt.Println(v144)
                                                                            _ = len(v3)
                                                                        } else {
                                                                            v97 = (v97 << ((18446744073709551615) & 63))
                                                                            fmt.Println(v97, v96)
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v153 uint16 = 842
                                                        _ = v153
                                                    }
                                                    var v154 uint = 798
                                                    _ = v144
                                                    _ = v154
                                                }
                                                _ = v140
                                            case 32767:
                                                if (2 < (v97 & 1)) {
                                                    if true {
                                                        _ = len(v2)
                                                    }
                                                }
                                            default:
                                                var v155 bool = true
                                                _ = v155
                                            }
                                            _ = v138
                                            _ = v139
                                        }
                                    }
                                    v95 = v95
                                    v95 -= 1
                                    _ = fi115
                                }
                            } else {
                                for fi156 := 0; fi156 < 20; fi156++ {
                                    switch ("foo bar" + "foo bar") {
                                    case "a":
                                        var v157 float64 = 3.14159
                                        _ = v157
                                    default:
                                        v95 = (v95 - (v95 / ((v95) | 1)))
                                    }
                                    var v158 uint32 = v97
                                    _ = fi156
                                    _ = v158
                                }
                                v159 := GenS19{F0: (v96 ^ (v96 << ((446) & 63))), F1: v97}
                                _ = v159
                            }
                            for fi160 := 0; fi160 < 4; fi160++ {
                                v2 = append(v2, v97)
                                v161 := v3[(v97 | 2147483648)]
                                v162 := GenS20{F0: 1}
                                _ = fi160
                                _ = v161
                                _ = v162
                            }
                            switch (("" + "a") + "foo bar") {
                            case "":
                                if len(v2) > 0 {
                                    v2[0] = (v97 % ((v97) | 1))
                                }
                            case "\n\t":
                                for fi163 := 0; fi163 < 12; fi163++ {
                                    v95 -= v95
                                    v164 := GenS2{F0: fi163, F1: 393, F2: 0}
                                    _ = fi163
                                    _ = v164
                                }
                            default:
                                v95 += v95
                            }
                            _ = fi111
                        }
                    }
                }
                var v165 bool = false
                _ = v165
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
