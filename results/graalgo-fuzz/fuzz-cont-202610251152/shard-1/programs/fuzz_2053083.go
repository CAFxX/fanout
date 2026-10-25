package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint16; F2 string }
type GenS7 struct { F0 int; F1 bool; F2 int32 }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8 }
func genFn0() uint16 {
    var v7 int16 = 42
    if false {
        var v8 float32 = ((1e10 / ((0.1 / 1e10) / (float32(0.1) + 0.1))) + 1.5)
        v7 = (v7 + (v7 | (v7 & (^(v7)))))
        v9 := []int{650}
        _ = v8
        _ = v9
    }
    switch -2147483648 {
    case -263:
        v7 = (v7 % ((-32768) | 1))
    case -41:
        for fi10 := 0; fi10 < 1; fi10++ {
            if (-1 > 182) {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 string = "foo bar"
                    var v13 uint16 = 0
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            }
            var v14 float32 = 0.1
            switch "a" {
            case "\n\t":
                v14 -= (((v14 - float32(3.14159)) / -2.25) * 1e10)
            }
            _ = fi10
            _ = v14
        }
    case 32767:
        v15 := GenS1{F0: 173, F1: 42, F2: 2147483648}
        _ = v15
    default:
        v7 += (v7 >> ((0) & 63))
    }
    _ = v7
    return 0
}
func genFn1(p0 int32) int16 {
    v16 := GenS2{F0: 42}
    if (!false) {
        v17 := []int{}
        if true {
            fmt.Println(p0)
            var v18 bool = false
            _ = v18
        } else {
            if (0.1 < float64(1.5)) {
                if (3.14159 < (float32(1e10) - ((3.14159 * 0.0) / float32(((((1e10 / ((1.5 - 0.0) + (0.1 - 1.5))) / 1e10) + 1e10) * -2.25))))) {
                    fmt.Println(p0)
                    p0 -= (^(p0))
                } else {
                    p0 = (p0 ^ 0)
                }
            } else {
                var v19 int = -9223372036854775808
                v20 := GenS1{F0: (p0 >> ((65535) & 63)), F1: (p0 | p0), F2: 201}
                var v21 uint16 = 255
                _ = v19
                _ = v20
                _ = v21
            }
        }
        if (42 >= 42) {
            p0 = 42
            if true {
                var v22 float64 = 3.14159
                var v23 uint8 = 127
                _ = v22
                _ = v23
            }
        }
        _ = v17
    }
    switch (("hello" + ("\n\t" + ("a" + ""))) + "hello") {
    case "hello":
        switch -1 {
        case 9223372036854775807:
            if (722 < 1) {
                p0 += (p0 >> ((285) & 63))
            } else {
                p0 += (p0 | p0)
            }
        case 32768:
            p0 -= (p0 / (((p0 >> ((0) & 63))) | 1))
        case 433:
            if (2 < 50) {
                var v24 uint8 = 1
                v25 := []int{961, 128, 127}
                _ = v24
                _ = v25
            }
        default:
            v26 := []int{82, 972, 2147483647, 42}
            _ = v26
        }
    case "foo bar":
        p0 += (^(p0))
    default:
        var v27 float32 = -2.25
        _ = v27
    }
    _ = p0
    _ = v16
    return 32767
}

func main() {
    v1 := []uint32{671}
    v2 := []string{("hello" + ("a" + ("" + ("foo bar" + "hello"))))}
    var v3 float64 = -2.25
    v4 := []string{}
    var v5 uint = 998
    var v6 int = 42
    for fi28 := 0; fi28 < 7; fi28++ {
        if true {
            var v29 uint8 = 94
            _ = v29
        } else {
            for fi30 := 0; fi30 < 9; fi30++ {
                v31 := GenS3{F0: 2147483647}
                var v32 float64 = 1.5
                _ = fi30
                _ = v31
                _ = v32
            }
            v1 = append(v1, 42)
            if (!true) {
                v3 = (float64(-2.25) - v3)
                fmt.Println(v6, v5, fi28)
            }
        }
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        _ = fi28
    }
    _ = len(v4)
    v33 := GenS3{F0: 91}
    v6 = v6
    fmt.Println(v6, v6, v5)
    var v34 uint16 = 1
    v35 := GenS1{F0: 2147483647, F1: -1, F2: 65535}
    var v36 uint16 = uint16(v5)
    v6 -= v6
    for fi37 := 0; fi37 < 19; fi37++ {
        if true {
            switch fi37 {
            case 42:
                var v38 uint16 = (v34 / ((0) | 1))
                _ = v38
            case -9223372036854775808:
                v36 += 32767
            case 9223372036854775807:
                fmt.Println(v5, fi37)
            }
        }
        v5 -= v5
        v39 := GenS2{F0: (fi37 * 745)}
        _ = fi37
        _ = v39
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "\n\t")
    }
    var v40 int16 = 1
    v4 = append(v4, "\n\t")
    v5 -= v5
    for fi41 := 0; fi41 < 20; fi41++ {
        for fi42 := 0; fi42 < 4; fi42++ {
            _ = len(v4)
            _ = fi42
        }
        switch "hello" {
        case "hello":
            var v43 uint64 = uint64(v36)
            _ = v43
        }
        var v44 uint16 = (v36 | 0)
        _ = fi41
        _ = v44
    }
    fmt.Println(v3)
    if true {
        var v45 float32 = 0.0
        switch 0 {
        case 42:
            v5 = (^(v5))
        case 1:
            if (2 == genFn0()) {
                v5 -= (v5 | (v5 - 911))
                v40 -= v40
            }
        default:
            if (233 != 42) {
                var v46 int16 = v40
                _ = v46
            } else {
                fmt.Println(v36, v36)
                if ((v36 >> (((^(v5))) & 63)) != v36) {
                    for fi47 := 0; fi47 < 12; fi47++ {
                        if ((-265 >= 9223372036854775807) && (false || ((42 == 100) || true))) {
                            if ((2147483647 < -1) || (1 != uint8(v36))) {
                                v34 -= v34
                                var v48 int = -2147483648
                                if (!(-2147483648 < (-(v48)))) {
                                    var v49 int64 = -190
                                    switch (("foo bar" + (("hello" + ("" + "a")) + ("a" + ((("foo bar" + "a") + "hello") + "\n\t")))) + "") {
                                    case "hello":
                                        if ((2147483647 > 65535) || (-878 <= (v49 >> ((255) & 63)))) {
                                            v36 += (^(v34))
                                            for fi50 := 0; fi50 < 15; fi50++ {
                                                switch "a" {
                                                case "foo bar":
                                                    switch 42 {
                                                    case 0:
                                                        fmt.Println(fi47, fi50)
                                                    case 2147483647:
                                                        v49 -= v49
                                                    case 4294967295:
                                                        if len(v2) > 0 {
                                                            v2[0] = (("a" + (((("foo bar" + "a") + "") + ("foo bar" + "")) + "\n\t")) + "\n\t")
                                                        }
                                                    default:
                                                        fmt.Println(v5, v6, v49)
                                                    }
                                                case "a":
                                                    v3 += float64((((0.0 - ((v45 - v45) / v45)) * float32(0.0)) / (v45 + 1e10)))
                                                }
                                                _ = fi50
                                            }
                                        } else {
                                            var v51 bool = false
                                            _ = v51
                                        }
                                    default:
                                        var v52 uint32 = 42
                                        _ = v52
                                    }
                                    _ = v49
                                }
                                _ = v48
                            } else {
                                fmt.Println(v45, v6, fi47)
                                if ((138 != 211) || false) {
                                    if (2 > v36) {
                                        fmt.Println(v45, fi47)
                                        v5 = (^(v5))
                                    } else {
                                        var v53 int64 = -9223372036854775808
                                        v4 = append(v4, "\n\t")
                                        if true {
                                            if ((v5 * v5) < (v5 >> ((v5) & 63))) {
                                                if ((v40 >> (((v5 % ((v5) | 1))) & 63)) <= 1000) {
                                                    if (!(!false)) {
                                                        v34 = v34
                                                        var v54 float32 = ((v45 + (v45 - v45)) / 1.5)
                                                        v53 = 42
                                                        _ = v54
                                                    } else {
                                                        switch (fi47 / ((fi47) | 1)) {
                                                        case 9223372036854775807:
                                                            fmt.Println(fi47)
                                                        default:
                                                            var v55 uint = 42
                                                            _ = v55
                                                        }
                                                        v2 = append(v2, ("hello" + "\n\t"))
                                                    }
                                                    v56 := GenS1{F0: 2147483647, F1: -1, F2: 806}
                                                    var v57 uint8 = 145
                                                    _ = v56
                                                    _ = v57
                                                }
                                                if false {
                                                    switch 9223372036854775807 {
                                                    case -1:
                                                        var v58 uint = v5
                                                        _ = v58
                                                    case 484:
                                                        v6 -= v6
                                                    case 4294967295:
                                                        switch "" {
                                                        case "\n\t":
                                                            var v59 int32 = -1
                                                            _ = v59
                                                        case "":
                                                            fmt.Println(fi47, v53, fi47)
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = "foo bar"
                                                            }
                                                        default:
                                                            if (false && false) {
                                                                v45 -= (v45 / v45)
                                                                fmt.Println(v34, v36)
                                                                var v60 int16 = (v40 / ((v40) | 1))
                                                                _ = v60
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v53, v34, v36)
                                                    }
                                                }
                                                var v61 bool = true
                                                _ = v61
                                            } else {
                                                var v62 int16 = v40
                                                var v63 int64 = 1
                                                _ = v62
                                                _ = v63
                                            }
                                            switch (("" + "\n\t") + "") {
                                            case "\n\t":
                                                if false {
                                                    var v64 bool = false
                                                    _ = v64
                                                } else {
                                                    v3 = 3.14159
                                                    switch "" {
                                                    case "a":
                                                        fmt.Println(v3)
                                                    case "hello":
                                                        v65 := GenS1{F0: 0, F1: 0, F2: 1}
                                                        _ = v65
                                                    case "foo bar":
                                                        fmt.Println(v40, v34, v6)
                                                    default:
                                                        var v66 uint16 = v34
                                                        _ = v66
                                                    }
                                                    if true {
                                                        v34 += (^(v36))
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483648
                                                        }
                                                    } else {
                                                        var v67 int64 = 0
                                                        v68 := GenS3{F0: 2}
                                                        v45 = (-2.25 / (v45 / (v45 - v45)))
                                                        _ = v67
                                                        _ = v68
                                                    }
                                                }
                                            case "hello":
                                                v45 = v45
                                            default:
                                                for fi69 := 0; fi69 < 18; fi69++ {
                                                    v70 := GenS4{F0: (v5 & (v5 & (v5 / (((v5 % ((v5) | 1))) | 1)))), F1: 0}
                                                    var v71 uint16 = v36
                                                    v72 := GenS5{F0: -128}
                                                    _ = fi69
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                }
                                            }
                                            v53 -= 42
                                        }
                                        _ = v53
                                    }
                                    for fi73 := 0; fi73 < 15; fi73++ {
                                        v2 = append(v2, ((("foo bar" + "hello") + "a") + ("\n\t" + "hello")))
                                        v74 := GenS3{F0: 42}
                                        _ = fi73
                                        _ = v74
                                    }
                                }
                            }
                            v75 := GenS2{F0: 128}
                            _ = v75
                        }
                        switch "" {
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = 65535
                            }
                        default:
                            if (true || (((float32(v3) / (((0.1 + v45) / float32(v3)) / (0.0 - v45))) + v45) != float32(v3))) {
                                for fi76 := 0; fi76 < 14; fi76++ {
                                    v77 := GenS6{F0: 666, F1: v34, F2: "hello"}
                                    var v78 uint = v5
                                    v79 := GenS3{F0: 0}
                                    var v80 uint64 = 4294967295
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                        for fi81 := 0; fi81 < 4; fi81++ {
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            v6 -= (v6 - v6)
                            v82 := GenS5{F0: 0}
                            _ = fi81
                            _ = v82
                        }
                        _ = fi47
                    }
                    var v83 uint8 = 212
                    for fi84 := 0; fi84 < 8; fi84++ {
                        for fi85 := 0; fi85 < 6; fi85++ {
                            v86 := GenS1{F0: 1, F1: 42, F2: uint32(v36)}
                            if false {
                                var v87 uint32 = 42
                                if len(v1) > 0 {
                                    v1[0] = 2147483648
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = v87
                            } else {
                                if (!true) {
                                    switch "" {
                                    case "":
                                        fmt.Println(v45)
                                    case "hello":
                                        v5 = (v5 + 1)
                                    }
                                    v88 := GenS7{F0: (fi85 << ((v5) & 63)), F1: false, F2: -860}
                                    var v89 float32 = 1.5
                                    if (false || (v3 != v3)) {
                                        fmt.Println(v5)
                                        v5 = 10
                                        if (genFn1(-52) == (^(v40))) {
                                            var v90 uint64 = 42
                                            switch "foo bar" {
                                            case "":
                                                fmt.Println(v6, v40)
                                            case "a":
                                                if false {
                                                    if true {
                                                        var v91 uint64 = v90
                                                        var v92 string = ("hello" + "foo bar")
                                                        v3 = v3
                                                        _ = v91
                                                        _ = v92
                                                    } else {
                                                        if true {
                                                            var v93 float32 = (v89 * 0.1)
                                                            _ = v93
                                                        } else {
                                                            var v94 int16 = genFn1(-2147483648)
                                                            _ = v94
                                                        }
                                                    }
                                                    v36 -= (v36 / (((v34 - genFn0())) | 1))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v45 = ((v89 + v45) * 0.1)
                                                } else {
                                                    v36 = 42
                                                    switch "hello" {
                                                    case "":
                                                        var v95 float64 = (v3 - 1e10)
                                                        _ = v95
                                                    case "hello":
                                                        var v96 uint16 = 285
                                                        _ = v96
                                                    case "\n\t":
                                                        var v97 uint32 = uint32(v83)
                                                        _ = v97
                                                    }
                                                    var v98 uint = v5
                                                    fmt.Println(v90)
                                                    _ = v98
                                                }
                                            default:
                                                var v99 int64 = int64(v90)
                                                _ = v99
                                            }
                                            fmt.Println(v89)
                                            _ = v90
                                        }
                                        switch int(v36) {
                                        case -221:
                                            v6 += (-(fi85))
                                        case 1:
                                            _ = len(v1)
                                        }
                                    } else {
                                        var v100 uint8 = (^(v83))
                                        _ = v100
                                    }
                                    _ = v88
                                    _ = v89
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                    if true {
                                        switch ("foo bar" + "foo bar") {
                                        case "\n\t":
                                            var v101 int16 = v40
                                            _ = v101
                                        }
                                    }
                                    var v102 bool = true
                                    _ = v102
                                }
                                v5 = (v5 / (((v5 / ((v5) | 1))) | 1))
                                switch "" {
                                case "":
                                    v45 -= 1.5
                                    var v103 int64 = -9223372036854775808
                                    _ = v103
                                case "hello":
                                    var v104 int = v6
                                    v6 += (v104 * 852)
                                    if false {
                                        if (false || true) {
                                            if true {
                                                if true {
                                                    v36 = v36
                                                } else {
                                                    v34 = v34
                                                    switch "hello" {
                                                    case "a":
                                                        if len(v1) > 0 {
                                                            v1[0] = 0
                                                        }
                                                    case "hello":
                                                        fmt.Println(v45, v104, v3)
                                                    }
                                                    var v105 bool = true
                                                    _ = v105
                                                }
                                            }
                                            v45 += (0.1 * float32(float64(1.5)))
                                        }
                                        var v106 int8 = 42
                                        _ = v106
                                    } else {
                                        if false {
                                            v40 = v40
                                            var v107 float32 = 0.0
                                            _ = v107
                                        } else {
                                            var v108 int = v6
                                            v109 := GenS2{F0: 482}
                                            _ = v108
                                            _ = v109
                                        }
                                    }
                                    _ = v104
                                case "foo bar":
                                    v3 -= 3.14159
                                }
                            }
                            v3 += v3
                            _ = fi85
                            _ = v86
                        }
                        _ = fi84
                    }
                    _ = v83
                } else {
                    var v110 uint8 = 42
                    for fi111 := 0; fi111 < 20; fi111++ {
                        v112 := GenS8{F0: (v34 * uint16(v5)), F1: 0, F2: 2147483647}
                        v5 += v5
                        switch ("foo bar" + "hello") {
                        case "":
                            v113 := GenS9{F0: 4294967295}
                            _ = v113
                        case "a":
                            v45 += v45
                        case "\n\t":
                            v114 := GenS3{F0: 2}
                            _ = v114
                        }
                        _ = fi111
                        _ = v112
                    }
                    _ = v110
                }
                for fi115 := 0; fi115 < 3; fi115++ {
                    var v116 float32 = (float32(float64(0.0)) + float32(0.0))
                    _ = fi115
                    _ = v116
                }
            }
        }
        var v117 uint32 = 2
        _ = v117
        _ = v45
    } else {
        v118 := GenS7{F0: v6, F1: false, F2: 31}
        _ = v118
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            var v119 uint = uint(v34)
            _ = v119
        }
    } else {
        v120 := GenS10{F0: v6}
        if (false || true) {
            fmt.Println(v3, v6)
            for fi121 := 0; fi121 < 3; fi121++ {
                var v122 int64 = 9223372036854775807
                if (2147483647 >= 42) {
                    v5 += (v5 - v5)
                    switch (("a" + (("a" + "\n\t") + ((("" + "a") + (("foo bar" + "foo bar") + (("foo bar" + "\n\t") + "a"))) + "a"))) + "hello") {
                    case "hello":
                        for fi123 := 0; fi123 < 13; fi123++ {
                            if true {
                                v124 := GenS9{F0: 0}
                                _ = v124
                            } else {
                                v36 = 65535
                            }
                            _ = fi123
                        }
                    case "\n\t":
                        v4 = append(v4, "a")
                    }
                    v1 = append(v1, 2147483648)
                } else {
                    v125 := GenS3{F0: 42}
                    if false {
                        if false {
                            var v126 uint8 = 127
                            var v127 uint32 = 2
                            _ = v126
                            _ = v127
                        } else {
                            for fi128 := 0; fi128 < 8; fi128++ {
                                var v129 uint16 = genFn0()
                                _ = fi128
                                _ = v129
                            }
                            for fi130 := 0; fi130 < 20; fi130++ {
                                v4 = append(v4, ("a" + "\n\t"))
                                var v131 float32 = float32(0.1)
                                v132 := GenS11{F0: 1}
                                _ = fi130
                                _ = v131
                                _ = v132
                            }
                        }
                    }
                    _ = v125
                }
                _ = fi121
                _ = v122
            }
        }
        for fi133 := 0; fi133 < 13; fi133++ {
            var v134 bool = (((65535 < 2147483648) || (!false)) || true)
            switch v6 {
            case -2147483648:
                v135 := GenS12{F0: ""}
                _ = v135
            case 0:
                v136 := GenS13{F0: -128}
                _ = v136
            case -1:
                switch ("" + "") {
                case "":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "foo bar":
                    v137 := GenS4{F0: v5, F1: (v36 << (((v5 >> ((989) & 63))) & 63))}
                    _ = v137
                case "\n\t":
                    var v138 uint64 = 42
                    _ = v138
                default:
                    var v139 uint32 = 42
                    _ = v139
                }
            }
            var v140 int16 = v40
            _ = fi133
            _ = v134
            _ = v140
        }
        _ = v120
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
