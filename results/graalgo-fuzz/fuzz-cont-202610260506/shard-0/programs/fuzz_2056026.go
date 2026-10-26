package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int }

func main() {
    var v1 int = 4294967295
    var v2 float32 = 3.14159
    v3 := make(map[uint32]uint32)
    v4 := []int{v1, -325, 42}
    var v5 int = (v1 + v1)
    v6 := []uint32{}
    var v7 float64 = float64((v2 - v2))
    var v8 uint32 = 237
    v1 += v5
    if (!(42 < -760)) {
        switch (-(v1)) {
        case -9223372036854775808:
            var v9 string = ("a" + "foo bar")
            _ = v9
        default:
            var v10 int8 = 104
            _ = v10
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            for fi12 := 0; fi12 < 7; fi12++ {
                v7 -= v7
                _ = fi12
            }
            _ = fi11
        }
        if (2147483647 != v5) {
            v13 := GenS1{F0: 883}
            if true {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v6 = append(v6, (^(v8)))
                    var v15 uint = 749
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi14
                    _ = v15
                }
                v16 := GenS1{F0: 255}
                if (-270 <= 2147483647) {
                    switch v5 {
                    case -2147483648:
                        fmt.Println(v2, v1, v7)
                    default:
                        if (0 <= 42) {
                            for fi17 := 0; fi17 < 16; fi17++ {
                                if (!true) {
                                    fmt.Println(v8, v8, fi17)
                                    for fi18 := 0; fi18 < 5; fi18++ {
                                        _ = len(v4)
                                        v19 := GenS1{F0: 4294967295}
                                        fmt.Println(v7, fi18)
                                        _ = len(v4)
                                        if ((!(!(v7 != v7))) && false) {
                                            if false {
                                                v7 = 1e10
                                                delete(v3, v8)
                                                switch ("a" + "a") {
                                                case "":
                                                    if len(v6) > 0 {
                                                        v6[0] = v8
                                                    }
                                                default:
                                                    var v20 uint64 = 14
                                                    _ = v20
                                                }
                                            }
                                            var v21 uint16 = 1
                                            _ = v21
                                        }
                                        _ = fi18
                                        _ = v19
                                    }
                                    v22 := GenS1{F0: 255}
                                    _ = v22
                                } else {
                                    var v23 bool = false
                                    _ = v23
                                }
                                _ = len(v3)
                                _ = fi17
                            }
                            var v24 int64 = 411
                            _ = v24
                        }
                    }
                    fmt.Println(v8)
                }
                _ = v16
            } else {
                if true {
                    v5 += (v1 % ((v5) | 1))
                    v25 := GenS1{F0: 566}
                    _ = v25
                } else {
                    v1 -= v1
                    if true {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            v27 := GenS1{F0: 2}
                            var v28 float32 = v2
                            v4 = append(v4, fi26)
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v7 -= v7
                        var v29 uint = 18446744073709551615
                        _ = v29
                    } else {
                        fmt.Println(v1, v8, v2)
                        switch (v1 ^ v1) {
                        case -2147483648:
                            if false {
                                if true {
                                    for fi30 := 0; fi30 < 7; fi30++ {
                                        for fi31 := 0; fi31 < 8; fi31++ {
                                            v7 = v7
                                            v32 := GenS2{F0: false}
                                            switch (fi30 ^ v1) {
                                            case 9223372036854775807:
                                                switch fi30 {
                                                case 127:
                                                    if ((^(v5)) != v1) {
                                                        if ((v5 ^ 903) != (v1 | fi31)) {
                                                            var v33 float32 = v2
                                                            _ = v33
                                                        } else {
                                                            if true {
                                                                v7 = (float64(0.1) - -2.25)
                                                                var v34 int = v1
                                                                _ = v34
                                                            } else {
                                                                v7 -= (v7 + v7)
                                                                v5 -= 127
                                                            }
                                                        }
                                                        if (0 == 1) {
                                                            if (false && (!(93 < 127))) {
                                                                var v35 uint = 1
                                                                _ = v35
                                                            } else {
                                                                v36 := GenS3{F0: v8}
                                                                _ = v36
                                                            }
                                                            var v37 string = "\n\t"
                                                            fmt.Println(v2)
                                                            switch (fi30 * fi31) {
                                                            case 4294967295:
                                                                if false {
                                                                    v37 = v37
                                                                    v38 := GenS4{F0: v37, F1: -2147483648, F2: 562}
                                                                    v39 := GenS1{F0: 2}
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                            case 2147483647:
                                                                if true {
                                                                    v8 += (v8 * 65535)
                                                                }
                                                                fmt.Println(v8, fi31, v37)
                                                                v40 := GenS2{F0: ((1 == 42) && true)}
                                                                var v41 uint16 = 42
                                                                _ = v40
                                                                _ = v41
                                                            case 872:
                                                                v42 := GenS5{F0: 1, F1: 0}
                                                                _ = v42
                                                            }
                                                            _ = v37
                                                        }
                                                    }
                                                case 0:
                                                    v2 -= 0.0
                                                case 42:
                                                    var v43 uint = 255
                                                    _ = v43
                                                }
                                            case 128:
                                                var v44 uint = 18446744073709551615
                                                _ = v44
                                            case -9223372036854775808:
                                                v45 := GenS3{F0: v8}
                                                _ = v45
                                            default:
                                                var v46 uint64 = 18446744073709551615
                                                _ = v46
                                            }
                                            _ = fi31
                                            _ = v32
                                        }
                                        var v47 int8 = 0
                                        if true {
                                            if (v47 >= 30) {
                                                if (0 <= 729) {
                                                    if (!true) {
                                                        if (-2147483648 < 42) {
                                                            var v48 int16 = 0
                                                            v5 = (^(fi30))
                                                            var v49 int32 = 2147483647
                                                            _ = v48
                                                            _ = v49
                                                        } else {
                                                            if ("hello" < (("hello" + ("\n\t" + (("" + "foo bar") + "hello"))) + "")) {
                                                                if (200 == 127) {
                                                                    v5 += (v5 & 4294967295)
                                                                } else {
                                                                    for fi50 := 0; fi50 < 2; fi50++ {
                                                                        var v51 uint = 2
                                                                        var v52 int64 = 42
                                                                        if (("foo bar" + "\n\t") <= "foo bar") {
                                                                            switch int(v51) {
                                                                            case 308:
                                                                                v7 -= float64(v2)
                                                                            case 790:
                                                                                var v53 float64 = v7
                                                                                _ = v53
                                                                            default:
                                                                                v1 = v1
                                                                            }
                                                                            if true {
                                                                                v1 = fi30
                                                                            } else {
                                                                                var v54 uint64 = 4294967295
                                                                                _ = v54
                                                                            }
                                                                            var v55 uint = v51
                                                                            _ = v55
                                                                        }
                                                                        v4 = append(v4, (^(v5)))
                                                                        v56 := GenS5{F0: 141, F1: 16}
                                                                        _ = fi50
                                                                        _ = v51
                                                                        _ = v52
                                                                        _ = v56
                                                                    }
                                                                }
                                                                var v57 uint16 = 516
                                                                v5 -= (fi30 * v1)
                                                                _ = v57
                                                            }
                                                            var v58 uint8 = 128
                                                            _ = v58
                                                        }
                                                    } else {
                                                        v59 := v3[(v8 << ((276) & 63))]
                                                        var v60 string = "\n\t"
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                    var v61 float64 = 3.14159
                                                    var v62 uint64 = 2
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    var v63 float32 = (((((1.5 / (v2 / (v2 * float32(v7)))) - float32(v7)) - (v2 + v2)) - float32(((v7 - v7) * float64(v2)))) / (0.0 / v2))
                                                    v6 = append(v6, (v8 >> ((2) & 63)))
                                                    v2 = v2
                                                    _ = v63
                                                }
                                                if false {
                                                    var v64 uint32 = (v8 | 2147483648)
                                                    v3[(v64 & v64)] = (v64 + v64)
                                                    _ = v64
                                                }
                                            }
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = v8
                                            }
                                        }
                                        _ = fi30
                                        _ = v47
                                    }
                                }
                                fmt.Println(v2, v1, v5)
                            }
                        case -941:
                            switch v1 {
                            case -1:
                                var v65 float32 = v2
                                _ = v65
                            case 716:
                                if (1 >= 699) {
                                    v66 := GenS3{F0: 1}
                                    _ = v66
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = len(v3)
                                    v7 -= (0.1 * v7)
                                }
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    switch "" {
                                    case "foo bar":
                                        if (!((1e10 / 1e10) >= v7)) {
                                            var v67 uint32 = v8
                                            v68 := GenS4{F0: "", F1: 863, F2: 1}
                                            v7 += v7
                                            _ = v67
                                            _ = v68
                                        }
                                    case "hello":
                                        var v69 uint32 = 0
                                        _ = v69
                                    case "\n\t":
                                        if true {
                                            var v70 int32 = 174
                                            for fi71 := 0; fi71 < 12; fi71++ {
                                                if true {
                                                    var v72 int = 127
                                                    _ = v72
                                                } else {
                                                    v73 := GenS2{F0: true}
                                                    v7 = (0.0 * v7)
                                                    _ = v73
                                                }
                                                _ = len(v3)
                                                v8 -= 2147483648
                                                _ = fi71
                                            }
                                            var v74 bool = ((352 == 4294967295) && (v8 == v8))
                                            _ = v70
                                            _ = v74
                                        }
                                    }
                                case "hello":
                                    if false {
                                        if (-1 == -9223372036854775808) {
                                            var v75 bool = true
                                            _ = v75
                                        } else {
                                            if (9223372036854775807 < 0) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v76 int16 = -1
                                                fmt.Println(v5)
                                                _ = v76
                                            }
                                            delete(v3, v8)
                                        }
                                        if false {
                                            v5 = v5
                                            fmt.Println(v5)
                                            fmt.Println(v7, v2)
                                        } else {
                                            _ = len(v3)
                                            v1 = (v1 + 254)
                                        }
                                    }
                                }
                            }
                        default:
                            if (true && true) {
                                v77 := GenS6{F0: -1}
                                v5 = (v1 * 32767)
                                fmt.Println(v7, v5, v2)
                                _ = v77
                            }
                        }
                        fmt.Println(v1, v2)
                    }
                }
            }
            _ = v13
        } else {
            v78 := GenS7{F0: "foo bar", F1: 100}
            for fi79 := 0; fi79 < 19; fi79++ {
                v80 := GenS8{F0: 42, F1: 0}
                _ = fi79
                _ = v80
            }
            switch "foo bar" {
            case "hello":
                var v81 float32 = 0.1
                _ = v81
            case "foo bar":
                if ((0.0 - float64(v2)) != v7) {
                    v2 = v2
                    var v82 int32 = 376
                    _ = v82
                } else {
                    for fi83 := 0; fi83 < 11; fi83++ {
                        v7 = v7
                        v1 -= 423
                        _ = fi83
                    }
                    var v84 bool = false
                    var v85 uint32 = (^(v8))
                    _ = v84
                    _ = v85
                }
            }
            _ = v78
        }
    }
    switch (v5 << ((200) & 63)) {
    case -740:
        var v86 int64 = 42
        _ = v86
    case -413:
        fmt.Println(v1, v5, v8)
    case -760:
        if (!("foo bar" >= ("" + ""))) {
            for fi87 := 0; fi87 < 3; fi87++ {
                v1 += (fi87 ^ v5)
                var v88 uint64 = 0
                _ = fi87
                _ = v88
            }
            v7 = 3.14159
            var v89 int64 = int64(v5)
            _ = v89
        } else {
            v5 -= 0
            v90 := GenS4{F0: "\n\t", F1: 0, F2: 1}
            _ = v90
        }
    }
    v2 = v2
    for fi91 := 0; fi91 < 10; fi91++ {
        for fi92 := 0; fi92 < 11; fi92++ {
            if ((!true) && (2 > uint64(v8))) {
                fmt.Println(v1, v8, fi92)
                var v93 uint64 = 1
                _ = v93
            }
            v2 = v2
            var v94 uint8 = 42
            _ = fi92
            _ = v94
        }
        _ = fi91
    }
    fmt.Println(v8)
    for fi95 := 0; fi95 < 11; fi95++ {
        v7 += 3.14159
        if true {
            var v96 int32 = -1
            v7 = v7
            _ = v96
        }
        _ = fi95
    }
    for fi97 := 0; fi97 < 19; fi97++ {
        for fi98 := 0; fi98 < 17; fi98++ {
            v99 := GenS7{F0: (("\n\t" + "a") + "a"), F1: -28}
            fmt.Println(v2)
            _ = fi98
            _ = v99
        }
        for fi100 := 0; fi100 < 11; fi100++ {
            fmt.Println(v8, v5)
            v101 := GenS5{F0: 107, F1: 4294967295}
            switch "" {
            case "a":
                fmt.Println(fi97, fi100, v2)
            }
            _ = fi100
            _ = v101
        }
        _ = fi97
    }
    for fi102 := 0; fi102 < 19; fi102++ {
        var v103 bool = (981 >= 18446744073709551615)
        for fi104 := 0; fi104 < 6; fi104++ {
            switch v5 {
            case -647:
                if v103 {
                    if ((true || (v103 || ((32767 >= 430) || (v103 || false)))) || false) {
                        v7 += float64(0.0)
                        v105 := GenS9{F0: -399}
                        if (768 <= 242) {
                            v5 = (v5 - (fi102 / (((v5 - fi102)) | 1)))
                            v6 = append(v6, v8)
                        } else {
                            v1 = -9223372036854775808
                            v5 -= (fi102 / ((fi104) | 1))
                        }
                        var v106 int64 = 1
                        var v107 uint16 = 2
                        _ = v105
                        _ = v106
                        _ = v107
                    }
                    if (!(!(120 > 0))) {
                        v108 := GenS2{F0: (fi104 <= (v5 * (fi102 << ((4294967295) & 63))))}
                        var v109 uint = 2
                        v103 = (true || false)
                        _ = v108
                        _ = v109
                    } else {
                        if (-729 <= 42) {
                            v2 -= v2
                        }
                        var v110 uint64 = 0
                        v111 := GenS10{F0: 578}
                        _ = v110
                        _ = v111
                    }
                } else {
                    v112 := GenS10{F0: (v5 / (((fi104 ^ (v1 + (v1 | fi104)))) | 1))}
                    _ = v112
                }
            case 42:
                var v113 int = 42
                _ = v113
            case 9223372036854775807:
                switch "foo bar" {
                case "hello":
                    delete(v3, (v8 * (v8 << ((4294967295) & 63))))
                case "":
                    if true {
                        if v103 {
                            fmt.Println(fi102, fi104, v1)
                        } else {
                            switch ("hello" + "foo bar") {
                            case "\n\t":
                                var v114 string = "foo bar"
                                _ = v114
                            case "foo bar":
                                if (!(1.5 > float64((((0.1 / ((1e10 - 1e10) / ((3.14159 / v2) - 0.0))) + 0.0) / (v2 - 1.5))))) {
                                    v8 = v8
                                }
                            case "":
                                fmt.Println(fi102, v8, fi102)
                            }
                            var v115 int16 = 0
                            _ = v115
                        }
                        fmt.Println(v2, v8)
                        switch "foo bar" {
                        case "":
                            switch (v1 & (fi104 << ((136) & 63))) {
                            case 32767:
                                fmt.Println(v8)
                            case 127:
                                v116 := GenS7{F0: ("a" + ""), F1: 42}
                                _ = v116
                            }
                        }
                        var v117 float32 = v2
                        var v118 int32 = -192
                        v119 := GenS10{F0: (fi102 & -29)}
                        _ = v117
                        _ = v118
                        _ = v119
                    } else {
                        switch (fi102 | v1) {
                        case 1:
                            v8 = (^(v8))
                        case 135:
                            if v103 {
                                v2 -= 0.0
                                v1 = v1
                            } else {
                                switch fi104 {
                                case 0:
                                    v103 = (float32(v7) == (v2 - 0.1))
                                }
                                var v120 uint64 = 42
                                _ = v120
                            }
                        case 128:
                            if v103 {
                                var v121 uint64 = 6
                                v103 = (-2.25 < v7)
                                _ = v121
                            }
                        default:
                            v1 = 42
                        }
                        var v122 int16 = -1
                        if ((false || v103) || ((0 > -128) || (!v103))) {
                            v2 -= 0.0
                        }
                        v7 = ((-2.25 / float64(((v2 - 0.0) + (v2 * -2.25)))) + v7)
                        _ = v122
                    }
                default:
                    var v123 int32 = 42
                    _ = v123
                }
            }
            fmt.Println(v5)
            _ = fi104
        }
        for fi124 := 0; fi124 < 3; fi124++ {
            fmt.Println(v7, fi102)
            v7 = v7
            var v125 int16 = 1000
            _ = fi124
            _ = v125
        }
        _ = fi102
        _ = v103
    }
    var v126 float32 = (-2.25 * (3.14159 - v2))
    fmt.Println("v1:", v1)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
