package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint, p1 uint64) uint {
    switch -65 {
    case 322:
        var v4 int32 = -2147483648
        _ = v4
    default:
        var v5 float64 = 1.5
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{fi6}
        v8 := GenS1{F0: int64(p1)}
        v7 = append(v7, fi6)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    return p0
}
func genFn1(p0 uint, p1 int, p2 int) uint64 {
    v9 := []int{p1, (p2 - 127), p2}
    if false {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint32 = 42
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi10
            _ = v11
        }
        var v12 uint64 = 42
        _ = v12
    }
    if ((0.0 < (3.14159 * 1.5)) || (!(1000 == 1000))) {
        fmt.Println(p1, p1, p2)
        if true {
            var v13 int8 = -128
            if false {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v13 = v13
                    _ = fi14
                }
                if ((p0 >= p0) && (-788 > 911)) {
                    var v15 int32 = 391
                    _ = v15
                }
                if (false && false) {
                    v13 = v13
                    var v16 int32 = -1
                    _ = v16
                }
            }
            p2 = p2
            _ = v13
        }
    } else {
        var v17 bool = (2 == 42)
        _ = v17
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{65535, 2, 982, 42}
    var v2 int = 32767
    v3 := make(map[string]int)
    var v18 int16 = 98
    if (1 == 1) {
        switch "a" {
        case "foo bar":
            v19 := GenS2{F0: 127, F1: 514}
            _ = v19
        case "\n\t":
            var v20 uint64 = genFn1(0, 4294967295, -2147483648)
            _ = v20
        default:
            switch v2 {
            case 32767:
                v18 -= -1
            case -2147483648:
                var v21 int = 63
                _ = v21
            case 127:
                var v22 uint32 = 2
                _ = v22
            default:
                for fi23 := 0; fi23 < 14; fi23++ {
                    if (215 == 255) {
                        if false {
                            if true {
                                v18 += (v18 + v18)
                                var v24 string = "foo bar"
                                _ = v24
                            }
                            fmt.Println(v18, v18)
                        } else {
                            v18 = (v18 * v18)
                            v25 := GenS3{F0: 242}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            fmt.Println(v2, fi26)
                            _ = fi26
                        }
                    }
                    v2 = fi23
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi23
                }
            }
        }
        if ((float32(-2.25) == 3.14159) || false) {
            for fi27 := 0; fi27 < 1; fi27++ {
                v1 = append(v1, 257)
                _ = fi27
            }
            v18 -= v18
            v2 -= -1
        } else {
            for fi28 := 0; fi28 < 20; fi28++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = -2147483648
                _ = fi28
            }
            if (false || (!(1e10 < ((-2.25 - float32(3.14159)) - (1e10 + (-2.25 - (1.5 * 1e10))))))) {
                v29 := GenS2{F0: 1, F1: 42}
                _ = v29
            } else {
                _ = len(v3)
                fmt.Println(v2, v18)
                for fi30 := 0; fi30 < 15; fi30++ {
                    v3[("hello" + "foo bar")] = fi30
                    v1 = append(v1, 579)
                    v2 = 0
                    _ = fi30
                }
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                if true {
                    var v32 int8 = 42
                    var v33 float64 = -2.25
                    fmt.Println(v2)
                    _ = v32
                    _ = v33
                }
                delete(v3, ("hello" + "\n\t"))
                _ = fi31
            }
        }
        delete(v3, "")
    }
    if true {
        switch (v2 ^ v2) {
        case 127:
            if (-427 < 0) {
                switch "foo bar" {
                case "\n\t":
                    v34 := v3[("hello" + ("\n\t" + "a"))]
                    _ = v34
                }
            } else {
                var v35 uint32 = 2147483648
                var v36 int64 = 0
                _ = v35
                _ = v36
            }
        default:
            switch (v2 + v2) {
            case -1:
                switch "" {
                case "":
                    v2 += v2
                }
            default:
                v2 = 4294967295
            }
        }
        for fi37 := 0; fi37 < 11; fi37++ {
            fmt.Println(fi37)
            var v38 int16 = (v18 >> ((genFn0(948, 18446744073709551615)) & 63))
            _ = fi37
            _ = v38
        }
    } else {
        if (((v18 << ((953) & 63)) > 1) || (128 <= 128)) {
            var v39 bool = (!(4294967295 <= 4294967295))
            v18 = -1
            v39 = v39
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            delete(v3, "foo bar")
            var v41 int = v2
            if true {
                var v42 int16 = v18
                var v43 uint8 = 42
                _ = v42
                _ = v43
            } else {
                if false {
                    var v44 uint8 = 7
                    var v45 bool = true
                    var v46 uint64 = 0
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v18 -= (v18 / ((1) | 1))
                }
                _ = len(v1)
            }
            _ = fi40
            _ = v41
        }
    }
    v47 := GenS2{F0: 42, F1: 274}
    switch 0 {
    case 9223372036854775807:
        if ((false || false) || false) {
            v18 = (v18 & (v18 % ((-1) | 1)))
            for fi48 := 0; fi48 < 20; fi48++ {
                v49 := GenS1{F0: 0}
                for fi50 := 0; fi50 < 2; fi50++ {
                    if true {
                        switch (fi48 * (v2 & fi50)) {
                        case -1:
                            fmt.Println(v2)
                        case -2147483648:
                            v51 := v3["hello"]
                            _ = v51
                        }
                        v2 = fi50
                    }
                    _ = fi50
                }
                _ = fi48
                _ = v49
            }
        } else {
            for fi52 := 0; fi52 < 15; fi52++ {
                switch "foo bar" {
                case "":
                    v53 := GenS2{F0: 0, F1: 2}
                    _ = v53
                case "\n\t":
                    v2 = 0
                }
                _ = fi52
            }
        }
    default:
        v54 := v3[("a" + ((("hello" + "") + ("\n\t" + "\n\t")) + "\n\t"))]
        _ = v54
    }
    switch 32768 {
    case 4294967295:
        v2 -= 456
    case 128:
        var v55 uint = 803
        _ = v55
    default:
        var v56 uint64 = 42
        _ = v56
    }
    var v57 int16 = 377
    switch v2 {
    case 9223372036854775807:
        for fi58 := 0; fi58 < 13; fi58++ {
            if len(v1) > 0 {
                v1[0] = 922
            }
            var v59 int16 = 42
            _ = fi58
            _ = v59
        }
    case 128:
        v2 = v2
    default:
        for fi60 := 0; fi60 < 17; fi60++ {
            v18 = v57
            if false {
                if true {
                    fmt.Println(fi60, fi60)
                    for fi61 := 0; fi61 < 10; fi61++ {
                        switch fi60 {
                        case 32768:
                            v57 = (v18 * 42)
                        }
                        var v62 int16 = 1
                        v62 -= v18
                        if true {
                            switch "foo bar" {
                            case "hello":
                                v63 := GenS4{F0: 0}
                                _ = v63
                            default:
                                v64 := GenS2{F0: -9, F1: 0}
                                _ = v64
                            }
                            v2 = 42
                            v65 := GenS5{F0: 492}
                            _ = v65
                        }
                        _ = fi61
                        _ = v62
                    }
                    for fi66 := 0; fi66 < 5; fi66++ {
                        v2 -= fi66
                        v67 := GenS6{F0: 127}
                        if false {
                            v2 += fi60
                            fmt.Println(fi60)
                            switch "hello" {
                            case "":
                                v2 += fi66
                            case "\n\t":
                                var v68 uint16 = 255
                                _ = v68
                            default:
                                v69 := GenS7{F0: 202}
                                _ = v69
                            }
                            v2 = v2
                        }
                        v70 := GenS7{F0: 42}
                        _ = fi66
                        _ = v67
                        _ = v70
                    }
                } else {
                    if (2147483647 != -2147483648) {
                        v71 := GenS7{F0: 4294967295}
                        _ = v71
                    }
                }
                if false {
                    if (1 < 393) {
                        v2 = v2
                        switch 0 {
                        case 128:
                            v3["\n\t"] = v2
                        case 283:
                            var v72 uint16 = 32767
                            _ = v72
                        }
                        switch "\n\t" {
                        case "\n\t":
                            for fi73 := 0; fi73 < 18; fi73++ {
                                v74 := GenS1{F0: -9223372036854775808}
                                var v75 int8 = 1
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                        }
                    } else {
                        delete(v3, (("foo bar" + "hello") + (("a" + ("hello" + "hello")) + "")))
                    }
                    var v76 bool = true
                    if false {
                        var v77 int8 = 0
                        _ = v77
                    } else {
                        if true {
                            var v78 string = ("a" + "hello")
                            var v79 int32 = 595
                            v80 := GenS6{F0: 111}
                            _ = v78
                            _ = v79
                            _ = v80
                        } else {
                            switch "hello" {
                            case "hello":
                                fmt.Println(v18, v57)
                            }
                            v2 -= (v2 + (fi60 / ((v2) | 1)))
                        }
                        if (4294967295 > 156) {
                            _ = len(v1)
                            switch ("hello" + ((("a" + "hello") + "\n\t") + "a")) {
                            case "":
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        v2 = 987
                                    case "a":
                                        v2 = v2
                                    case "hello":
                                        for fi81 := 0; fi81 < 13; fi81++ {
                                            var v82 float32 = 0.0
                                            v83 := GenS7{F0: 42}
                                            _ = fi81
                                            _ = v82
                                            _ = v83
                                        }
                                    default:
                                        var v84 int = (^(v2))
                                        _ = v84
                                    }
                                } else {
                                    v85 := GenS1{F0: 42}
                                    v76 = (3.14159 == 0.0)
                                    switch "foo bar" {
                                    case "":
                                        switch (-(v2)) {
                                        case 9223372036854775807:
                                            var v86 int16 = int16(fi60)
                                            _ = v86
                                        }
                                    case "hello":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "\n\t":
                                        v3[("hello" + ("" + "hello"))] = -170
                                    }
                                    _ = v85
                                }
                            case "hello":
                                switch fi60 {
                                case 1:
                                    v76 = (0.0 != (((((1e10 / -2.25) - ((3.14159 + (0.1 * 1.5)) / 0.1)) * float64(float32(1.5))) - 0.0) - 3.14159))
                                }
                            }
                            v87 := GenS6{F0: 0}
                            _ = v87
                        }
                    }
                    _ = v76
                } else {
                    v18 = v18
                    if (-9223372036854775808 < int64(v57)) {
                        _ = len(v1)
                        for fi88 := 0; fi88 < 1; fi88++ {
                            var v89 uint16 = 67
                            var v90 uint8 = 14
                            v18 += (v57 * v18)
                            _ = fi88
                            _ = v89
                            _ = v90
                        }
                        v57 = v57
                    } else {
                        if false {
                            if (2 != 18446744073709551615) {
                                v2 = (fi60 / (((fi60 * -9223372036854775808)) | 1))
                                for fi91 := 0; fi91 < 14; fi91++ {
                                    var v92 uint8 = 127
                                    _ = fi91
                                    _ = v92
                                }
                                if false {
                                    if ((("foo bar" + "a") + ((("a" + "hello") + "\n\t") + "hello")) >= "foo bar") {
                                        if true {
                                            for fi93 := 0; fi93 < 4; fi93++ {
                                                if (((!(true || ((!((-828 == -319) && ((false && (((0.0 == 1e10) || (42 < 2147483647)) || true)) || ("\n\t" == (("hello" + "\n\t") + "a"))))) || false))) && ((2 == 312) && false)) || (415 >= 969)) {
                                                    if (!false) {
                                                        _ = len(v1)
                                                    }
                                                } else {
                                                    var v94 int32 = 42
                                                    fmt.Println(v57, v94, fi60)
                                                    _ = len(v3)
                                                    _ = v94
                                                }
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = fi93
                                            }
                                            v18 += -32768
                                        }
                                        v1 = append(v1, 868)
                                    } else {
                                        v95 := GenS1{F0: 863}
                                        _ = v95
                                    }
                                    v18 = 966
                                    v57 = 1000
                                }
                            }
                            for fi96 := 0; fi96 < 12; fi96++ {
                                v97 := GenS1{F0: -9223372036854775808}
                                _ = len(v1)
                                switch fi60 {
                                case 31:
                                    if (true && (1 >= 31)) {
                                        v98 := GenS6{F0: 145}
                                        _ = len(v1)
                                        var v99 uint16 = 940
                                        var v100 uint32 = 31
                                        _ = v100
                                        _ = v98
                                        _ = v99
                                    }
                                case -2147483648:
                                    v57 = v18
                                case 1:
                                    v2 = (fi60 % ((fi96) | 1))
                                    if true {
                                        v101 := GenS8{F0: 78}
                                        if (true && false) {
                                            var v102 uint = 319
                                            v18 = v57
                                            _ = v102
                                        }
                                        _ = v101
                                    } else {
                                        v57 = (^(v18))
                                        switch ("foo bar" + ("\n\t" + "foo bar")) {
                                        case "foo bar":
                                            var v103 float64 = 0.1
                                            _ = v103
                                        default:
                                            if false {
                                                if false {
                                                    v104 := GenS2{F0: 89, F1: 255}
                                                    _ = v104
                                                } else {
                                                    fmt.Println(v18, v2)
                                                }
                                                var v105 uint64 = 1
                                                _ = v105
                                            }
                                        }
                                        v106 := v3[("\n\t" + ("a" + ""))]
                                        _ = v106
                                    }
                                    v107 := GenS4{F0: 118}
                                    _ = v107
                                }
                                _ = fi96
                                _ = v97
                            }
                        } else {
                            var v108 int16 = (v18 * 1000)
                            _ = v108
                        }
                        v57 = v18
                        switch -9223372036854775808 {
                        case 104:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    }
                }
            }
            _ = fi60
        }
    }
    v1 = append(v1, 65535)
    _ = len(v3)
    if (false || false) {
        v109 := v3["hello"]
        _ = v109
    }
    if (false || (true && true)) {
        fmt.Println(v18, v57, v57)
        var v110 bool = ("" < "")
        _ = v110
    } else {
        if (v57 > (v18 & (-(v57)))) {
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
        switch (^(v2)) {
        case 32768:
            if (0 == 2) {
                switch v2 {
                case -472:
                    var v111 float64 = (1.5 - 1e10)
                    _ = v111
                case 128:
                    v112 := GenS9{F0: (-(v57)), F1: -1}
                    _ = v112
                case 42:
                    var v113 uint64 = 650
                    _ = v113
                }
            }
        case 32767:
            var v114 uint32 = 4294967295
            _ = v114
        case -725:
            for fi115 := 0; fi115 < 8; fi115++ {
                var v116 uint = 42
                _ = fi115
                _ = v116
            }
        default:
            v117 := GenS10{F0: (0 == v57), F1: genFn1(558, (-(v2)), v2)}
            _ = v117
        }
        if (((v2 * v2) == (-(v2))) && true) {
            v2 = -161
        }
    }
    var v118 uint32 = 4294967295
    v2 = v2
    var v119 bool = true
    v18 = (-(v57))
    if v119 {
        v3["\n\t"] = (^(v2))
    } else {
        _ = len(v1)
        if true {
            switch (-(v2)) {
            case -1:
                for fi120 := 0; fi120 < 14; fi120++ {
                    if (v119 || (!v119)) {
                        v57 -= (v57 | -1)
                        if (128 > v2) {
                            switch (-(fi120)) {
                            case 9223372036854775807:
                                fmt.Println(v118, v18)
                            case 0:
                                if v119 {
                                    var v121 uint = genFn0(genFn0(42, genFn1(genFn0(503, 0), (fi120 | 360), v2)), uint64(v118))
                                    _ = v121
                                } else {
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        var v122 int8 = 100
                                        _ = v122
                                    case "a":
                                        v18 = -281
                                    case "":
                                        v1 = append(v1, (v118 - (v118 + (v118 >> ((18446744073709551615) & 63)))))
                                    }
                                }
                            default:
                                v123 := GenS3{F0: 42}
                                _ = v123
                            }
                            v57 -= v18
                        }
                    }
                    _ = fi120
                }
            case 284:
                v3["hello"] = v2
            }
            v118 = (v118 + (v118 & v118))
            fmt.Println(v18, v119)
        }
        v18 += (v57 - (v57 ^ 0))
    }
    switch v2 {
    case 42:
        var v124 int = (-(v2))
        _ = v124
    case -1:
        for fi125 := 0; fi125 < 14; fi125++ {
            v126 := GenS7{F0: (v118 ^ (v118 * v118))}
            _ = fi125
            _ = v126
        }
    default:
        v1 = append(v1, (^(v118)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v57:", v57)
    _ = v1
    _ = v118
    _ = v119
    _ = v18
    _ = v2
    _ = v3
    _ = v47
    _ = v57
}
