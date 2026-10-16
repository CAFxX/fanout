package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 bool }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 int16; F1 bool; F2 int }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int16; F1 uint64; F2 uint32 }
type GenS11 struct { F0 uint16; F1 int }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 uint64; F1 int8 }
type GenS14 struct { F0 string; F1 string }
type GenS15 struct { F0 int64; F1 int32 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int64; F1 int8; F2 uint8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int; F1 int64; F2 uint16 }
func genFn0(p0 int, p1 int8, p2 int32) string {
    var v7 uint = 255
    var v8 int64 = 1
    var v9 uint16 = 435
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{fi10}
        if true {
            var v12 bool = ((!(!((0.0 * 0.1) > (0.0 - float64(1e10))))) || ((-1 != 9) && true))
            switch -2147483648 {
            case 32767:
                var v13 bool = v12
                _ = v13
            case 9223372036854775807:
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 int32 = (p2 - p2)
                    _ = fi14
                    _ = v15
                }
            }
            _ = v12
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint16 = 101
            switch "foo bar" {
            case "hello":
                _ = len(v11)
            default:
                var v18 uint32 = 65535
                _ = v18
            }
            if true {
                switch fi10 {
                case -606:
                    fmt.Println(fi16)
                }
            }
            _ = fi16
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    p0 -= int(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return ""
}
func genFn1(p0 int32) bool {
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            if true {
                fmt.Println(p0, p0, fi19)
                switch (fi19 << ((42) & 63)) {
                case 9223372036854775807:
                    if (!false) {
                        if ((!false) && false) {
                            v21 := GenS1{F0: 0, F1: (p0 << ((65535) & 63)), F2: (!(-1 == 0))}
                            p0 = (p0 >> ((322) & 63))
                            v22 := GenS2{F0: 0, F1: 100, F2: ((p0 * p0) == (p0 - (p0 & p0)))}
                            _ = v21
                            _ = v22
                        } else {
                            p0 = 2147483647
                            p0 = p0
                        }
                        if false {
                            if false {
                                fmt.Println(fi19, fi20, fi20)
                                var v23 string = ("" + "hello")
                                _ = v23
                            }
                        } else {
                            switch ("hello" + "\n\t") {
                            case "":
                                var v24 uint32 = 1
                                v24 = (v24 ^ (v24 % (((v24 / ((v24) | 1))) | 1)))
                                var v25 int = 128
                                _ = v24
                                _ = v25
                            case "a":
                                p0 = (p0 - (p0 >> ((1) & 63)))
                            case "foo bar":
                                var v26 int32 = p0
                                _ = v26
                            default:
                                var v27 float32 = (0.0 + -2.25)
                                _ = v27
                            }
                        }
                        switch 0 {
                        case 32768:
                            if (((1.5 * ((0.0 - -2.25) * (1.5 + 3.14159))) + 3.14159) <= ((0.0 - 0.1) - float64((0.0 / (((0.1 + 1e10) / 1.5) + (0.0 - float32(1e10))))))) {
                                if true {
                                    fmt.Println(p0, fi20, p0)
                                    v28 := GenS1{F0: 0, F1: (p0 + (p0 * p0)), F2: (255 <= 65535)}
                                    v29 := GenS3{F0: 265, F1: fi19}
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 int32 = (p0 ^ p0)
                                    v31 := []int{870, 0}
                                    var v32 int8 = -1
                                    if true {
                                        if (46 > 255) {
                                            v32 -= v32
                                        } else {
                                            if (!(1.5 != float64(float32(1.5)))) {
                                                switch fi19 {
                                                case 320:
                                                    fmt.Println(fi20)
                                                }
                                            }
                                            v33 := GenS3{F0: 4294967295, F1: (fi19 * (fi19 % (((fi19 << ((42) & 63))) | 1)))}
                                            var v34 uint = 18446744073709551615
                                            _ = v33
                                            _ = v34
                                        }
                                    } else {
                                        v32 = 28
                                        var v35 uint = 127
                                        _ = v35
                                    }
                                    v36 := GenS1{F0: 1000, F1: (v30 | p0), F2: true}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                    _ = v36
                                }
                                p0 = (p0 ^ -1)
                                p0 = (p0 << ((18446744073709551615) & 63))
                                v37 := GenS4{F0: -1, F1: 989}
                                _ = v37
                            } else {
                                if false {
                                    var v38 float64 = (((float64((float32((1e10 + 0.0)) - 0.1)) + (float64(-2.25) + 3.14159)) / 1.0) / 1.0)
                                    if (!true) {
                                        var v39 int16 = 571
                                        v40 := GenS2{F0: -9223372036854775808, F1: -17, F2: (false && false)}
                                        v41 := []int{42, fi19}
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        if ((0.0 + ((0.0 - (0.1 * 3.14159)) / 0.1)) == 1e10) {
                                            p0 -= p0
                                            switch fi20 {
                                            case 862:
                                                if (478 != -797) {
                                                    v42 := GenS5{F0: true, F1: 1}
                                                    var v43 uint32 = 1
                                                    _ = v42
                                                    _ = v43
                                                }
                                            case 2147483647:
                                                v44 := []int{fi19, (fi20 * fi20), 42, -1}
                                                _ = v44
                                            default:
                                                var v45 uint32 = 231
                                                _ = v45
                                            }
                                            v38 += (v38 / (v38 / 3.14159))
                                        }
                                    }
                                    var v46 int64 = -9223372036854775808
                                    _ = v38
                                    _ = v46
                                }
                                if (255 <= 66) {
                                    p0 += (^(p0))
                                }
                            }
                        case -413:
                            p0 = p0
                        case -952:
                            v47 := GenS3{F0: 42, F1: (-(fi19))}
                            _ = v47
                        default:
                            v48 := []int{(fi20 << ((42) & 63)), fi19, 128, fi20}
                            if false {
                                p0 += p0
                                if false {
                                    p0 = (p0 >> ((2) & 63))
                                    switch fi20 {
                                    case 42:
                                        switch fi20 {
                                        case 32768:
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch (fi20 >> ((65535) & 63)) {
                                                case 1:
                                                    var v49 int = (fi20 & fi20)
                                                    _ = v49
                                                case 0:
                                                    var v50 float32 = float32(1e10)
                                                    _ = v50
                                                }
                                            }
                                        case 127:
                                            p0 -= (p0 * -442)
                                        }
                                    case -9223372036854775808:
                                        var v51 uint = 18446744073709551615
                                        _ = v51
                                    }
                                    var v52 uint = 255
                                    fmt.Println(fi20)
                                    var v53 uint16 = 65535
                                    if ((v52 * (v52 - v52)) <= (^(v52))) {
                                        switch "foo bar" {
                                        case "":
                                            if ((fi20 - (fi20 * fi20)) < fi19) {
                                                v53 -= 0
                                            }
                                        default:
                                            if (!(false && (100 > 0))) {
                                                var v54 int64 = 0
                                                _ = len(v48)
                                                var v55 int16 = 488
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    } else {
                                        v52 -= (^(v52))
                                        switch "\n\t" {
                                        case "a":
                                            fmt.Println(p0, v53)
                                        case "foo bar":
                                            var v56 float32 = 3.14159
                                            _ = v56
                                        case "\n\t":
                                            var v57 int32 = 0
                                            _ = v57
                                        default:
                                            var v58 int32 = p0
                                            _ = v58
                                        }
                                        if len(v48) > 0 {
                                            v48[0] = fi19
                                        }
                                    }
                                    v48 = append(v48, (fi19 % ((127) | 1)))
                                    v59 := GenS5{F0: (true && false), F1: uint64(v52)}
                                    _ = v52
                                    _ = v53
                                    _ = v59
                                }
                                switch (("hello" + "\n\t") + (("" + "hello") + "\n\t")) {
                                case "foo bar":
                                    var v60 float64 = ((0.0 / 1.5) * 0.0)
                                    _ = v60
                                case "\n\t":
                                    v61 := GenS1{F0: 42, F1: p0, F2: false}
                                    _ = v61
                                default:
                                    fmt.Println(fi19, p0, fi20)
                                }
                            }
                            if true {
                                if (255 >= 238) {
                                    if ((2147483648 >= 2147483648) || (false || (127 >= 42))) {
                                        var v62 uint16 = 0
                                        _ = v62
                                    } else {
                                        p0 = p0
                                    }
                                    if (true && true) {
                                        p0 -= p0
                                    } else {
                                        v63 := GenS6{F0: fi20}
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            v48 = append(v48, fi20)
                                        case "a":
                                            p0 = p0
                                        case "":
                                            fmt.Println(fi19)
                                        }
                                        if false {
                                            p0 = 42
                                            var v64 float32 = (-2.25 * 1e10)
                                            v64 = (v64 * (v64 * v64))
                                            _ = v64
                                        }
                                        _ = v63
                                    }
                                    var v65 int32 = -1
                                    _ = v65
                                } else {
                                    v66 := GenS7{F0: 370, F1: 2}
                                    fmt.Println(fi20, fi20, p0)
                                    _ = v66
                                }
                                switch fi19 {
                                case -381:
                                    p0 = (p0 & p0)
                                case 9223372036854775807:
                                    switch 0 {
                                    case 9223372036854775807:
                                        p0 -= (p0 ^ p0)
                                    case 32768:
                                        var v67 bool = false
                                        _ = v67
                                    }
                                    v48 = append(v48, 2147483647)
                                    _ = len(v48)
                                    p0 += (p0 + 2147483647)
                                case 2147483647:
                                    p0 -= (p0 % ((1) | 1))
                                default:
                                    fmt.Println(fi20, p0, fi20)
                                    if false {
                                        var v68 int = (fi19 ^ fi19)
                                        _ = v68
                                    }
                                }
                            }
                            _ = v48
                        }
                    }
                case 547:
                    switch ("\n\t" + ("" + "\n\t")) {
                    case "a":
                        v69 := []int{}
                        _ = v69
                    case "\n\t":
                        if true {
                            p0 += p0
                            p0 -= 1
                            v70 := []int{-298, (fi20 * (-(fi19)))}
                            if len(v70) > 0 {
                                v70[0] = fi19
                            }
                            fmt.Println(fi20)
                            p0 = (p0 - (p0 % ((970) | 1)))
                            _ = v70
                        }
                        fmt.Println(p0)
                        if true {
                            p0 -= (-(p0))
                        }
                        p0 = (p0 * (p0 % (((p0 ^ 0)) | 1)))
                        switch fi20 {
                        case 128:
                            switch fi20 {
                            case 127:
                                if false {
                                    p0 -= 1
                                    var v71 string = "a"
                                    _ = v71
                                } else {
                                    var v72 int32 = (^(p0))
                                    _ = v72
                                }
                            case 0:
                                p0 -= -1
                            }
                        case 4294967295:
                            var v73 int32 = 42
                            _ = v73
                        default:
                            p0 = (p0 << ((42) & 63))
                        }
                        switch (fi19 >> ((4294967295) & 63)) {
                        case 825:
                            switch "a" {
                            case "hello":
                                p0 = (^(p0))
                            }
                            v74 := GenS8{F0: 1, F1: false, F2: (^(fi20))}
                            v75 := GenS3{F0: 0, F1: (fi20 >> ((18446744073709551615) & 63))}
                            _ = v74
                            _ = v75
                        case 4294967295:
                            if true {
                                v76 := GenS9{F0: true, F1: fi20}
                                v77 := GenS8{F0: -1, F1: true, F2: (fi20 - (fi19 + fi20))}
                                v78 := []int{128, fi20}
                                _ = v76
                                _ = v77
                                _ = v78
                            } else {
                                v79 := []int{}
                                var v80 int = 262
                                _ = len(v79)
                                _ = v79
                                _ = v80
                            }
                        default:
                            if (1e10 < 1e10) {
                                if ((p0 < (p0 % (((p0 % ((p0) | 1))) | 1))) || false) {
                                    p0 = p0
                                } else {
                                    if true {
                                        p0 += p0
                                        var v81 bool = false
                                        v82 := []int{(fi20 * fi20), (^(fi20))}
                                        _ = v81
                                        _ = v82
                                    } else {
                                        p0 += (p0 >> ((42) & 63))
                                        var v83 int8 = -7
                                        _ = v83
                                    }
                                    if (0.1 == (1.5 - float64((1e10 * 3.14159)))) {
                                        v84 := GenS4{F0: 1, F1: 42}
                                        switch 646 {
                                        case 128:
                                            switch ("foo bar" + ("" + "")) {
                                            case "hello":
                                                v85 := GenS3{F0: 893, F1: (^(fi19))}
                                                _ = v85
                                            case "foo bar":
                                                var v86 uint = 299
                                                _ = v86
                                            }
                                        default:
                                            fmt.Println(fi19)
                                            if false {
                                                v87 := []int{(fi20 * fi20), (fi19 << ((4294967295) & 63)), (-(fi19))}
                                                var v88 int16 = 77
                                                v87 = append(v87, (fi19 - (fi20 >> ((42) & 63))))
                                                _ = v87
                                                _ = v88
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(p0, fi20)
                                                default:
                                                    p0 += (p0 | p0)
                                                }
                                                p0 = (p0 / (((^(p0))) | 1))
                                                var v89 int16 = 1
                                                _ = v89
                                            }
                                            p0 = (-(p0))
                                            if (("hello" + ("\n\t" + "a")) < (("\n\t" + "\n\t") + ("foo bar" + ("a" + "\n\t")))) {
                                                fmt.Println(fi19)
                                                v90 := []int{}
                                                if false {
                                                    if ((65535 == 0) && false) {
                                                        if (!false) {
                                                            v91 := GenS10{F0: -1, F1: 4294967295, F2: 180}
                                                            if true {
                                                                if true {
                                                                    fmt.Println(fi19, fi20)
                                                                    p0 -= -801
                                                                    v92 := GenS2{F0: 42, F1: 127, F2: false}
                                                                    p0 = 1
                                                                    _ = v92
                                                                } else {
                                                                    var v93 float64 = -2.25
                                                                    var v94 uint32 = 2147483648
                                                                    v95 := GenS11{F0: 32767, F1: (fi20 ^ (fi19 | fi19))}
                                                                    _ = v93
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                                if len(v90) > 0 {
                                                                    v90[0] = 1
                                                                }
                                                            }
                                                            _ = v91
                                                        }
                                                    } else {
                                                        switch fi20 {
                                                        case 42:
                                                            fmt.Println(fi20, fi19, fi19)
                                                        case 516:
                                                            var v96 int16 = 496
                                                            _ = v96
                                                        case 9223372036854775807:
                                                            p0 = p0
                                                        default:
                                                            fmt.Println(fi20, fi20, p0)
                                                        }
                                                    }
                                                }
                                                _ = v90
                                            }
                                            p0 -= (p0 % ((p0) | 1))
                                        }
                                        _ = v84
                                    }
                                    var v97 float64 = (1e10 + ((3.14159 / -2.25) * 0.0))
                                    p0 += -706
                                    _ = v97
                                }
                                var v98 uint64 = 18446744073709551615
                                var v99 float64 = float64((3.14159 * (1e10 - 1e10)))
                                _ = v98
                                _ = v99
                            } else {
                                if (65535 > 42) {
                                    switch ("a" + "\n\t") {
                                    case "":
                                        switch "hello" {
                                        case "hello":
                                            var v100 int8 = -1
                                            _ = v100
                                        case "":
                                            var v101 float32 = 1e10
                                            _ = v101
                                        case "\n\t":
                                            fmt.Println(fi19, p0, p0)
                                        }
                                    }
                                }
                                var v102 int8 = -128
                                switch fi20 {
                                case 1:
                                    fmt.Println(p0, fi19, fi19)
                                case 583:
                                    v102 -= (v102 << ((255) & 63))
                                case -2147483648:
                                    v102 = v102
                                default:
                                    var v103 float32 = -2.25
                                    _ = v103
                                }
                                _ = v102
                            }
                        }
                        p0 = p0
                        v104 := GenS9{F0: ((0.0 / 3.14159) <= (float32(float64(-2.25)) / 1.0)), F1: -56}
                        v105 := GenS12{F0: p0}
                        _ = v104
                        _ = v105
                    case "":
                        p0 = 0
                    default:
                        var v106 uint16 = 32767
                        _ = v106
                    }
                }
                switch (fi19 & fi20) {
                case 0:
                    var v107 string = "\n\t"
                    _ = v107
                case 868:
                    fmt.Println(fi19, fi20, fi19)
                case 32767:
                    v108 := GenS13{F0: 1, F1: 127}
                    _ = v108
                default:
                    v109 := GenS14{F0: "hello", F1: ((("foo bar" + "") + "a") + "foo bar")}
                    _ = v109
                }
                if false {
                    switch -753 {
                    case 32767:
                        v110 := GenS15{F0: -420, F1: (p0 >> ((543) & 63))}
                        _ = v110
                    case -9223372036854775808:
                        switch "a" {
                        case "\n\t":
                            switch (fi19 >> ((850) & 63)) {
                            case 127:
                                var v111 uint64 = 4294967295
                                p0 = p0
                                v112 := GenS16{F0: -126}
                                _ = v111
                                _ = v112
                            }
                        }
                    case 9223372036854775807:
                        switch (fi20 % (((fi19 % ((127) | 1))) | 1)) {
                        case -385:
                            fmt.Println(fi20, p0)
                        case -80:
                            switch fi19 {
                            case -2147483648:
                                var v113 uint32 = 876
                                _ = v113
                            case 128:
                                var v114 bool = (false || (-1 > 0))
                                _ = v114
                            default:
                                p0 -= (p0 % ((p0) | 1))
                            }
                        case -831:
                            v115 := []int{(fi19 - (fi19 / ((fi20) | 1))), (fi19 & 1)}
                            _ = v115
                        }
                    default:
                        v116 := []int{(fi19 ^ (^(fi19))), (fi20 << ((18446744073709551615) & 63)), (^(fi19))}
                        _ = v116
                    }
                    p0 += 42
                    var v117 float64 = (1.5 - (3.14159 * ((-2.25 / 3.14159) - 1.5)))
                    _ = v117
                } else {
                    v118 := GenS8{F0: -32768, F1: true, F2: fi20}
                    _ = v118
                }
                switch fi19 {
                case -1:
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        var v119 int64 = 9223372036854775807
                        _ = v119
                    case "a":
                        if false {
                            var v120 int16 = 42
                            v121 := GenS17{F0: 14}
                            p0 -= 1
                            _ = v120
                            _ = v121
                        }
                    }
                }
            } else {
                switch 9223372036854775807 {
                case -9223372036854775808:
                    v122 := GenS7{F0: 2, F1: 65535}
                    _ = v122
                case -667:
                    fmt.Println(p0, fi19)
                }
                p0 -= p0
                var v123 uint = 265
                switch -2147483648 {
                case 903:
                    v124 := GenS18{F0: 958, F1: 0, F2: 214}
                    _ = v124
                default:
                    p0 += -642
                }
                _ = v123
            }
            v125 := GenS5{F0: (-1 != -166), F1: 4294967295}
            _ = fi20
            _ = v125
        }
        _ = fi19
    }
    var v126 uint64 = 2
    for fi127 := 0; fi127 < 5; fi127++ {
        for fi128 := 0; fi128 < 3; fi128++ {
            v129 := GenS19{F0: 911}
            _ = fi128
            _ = v129
        }
        for fi130 := 0; fi130 < 6; fi130++ {
            var v131 bool = true
            var v132 string = "foo bar"
            _ = fi130
            _ = v131
            _ = v132
        }
        _ = fi127
    }
    switch 32768 {
    case 0:
        p0 = (-(p0))
    case 42:
        var v133 float64 = ((-2.25 - 0.0) - 1.5)
        _ = v133
    case 703:
        for fi134 := 0; fi134 < 2; fi134++ {
            v135 := GenS20{F0: 127, F1: -9223372036854775808, F2: 42}
            if true {
                var v136 float32 = float32(0.0)
                for fi137 := 0; fi137 < 14; fi137++ {
                    v126 -= (^(v126))
                    fmt.Println(fi134)
                    v136 = v136
                    _ = fi137
                }
                _ = v136
            } else {
                var v138 uint = 65535
                v126 = v126
                var v139 float32 = 1.5
                _ = v138
                _ = v139
            }
            _ = fi134
            _ = v135
        }
    }
    _ = p0
    _ = v126
    return (((-9223372036854775808 != -600) || false) && false)
}

func main() {
    var v1 uint = 42
    var v2 uint = v1
    v3 := make(map[uint32]string)
    var v4 int16 = 32767
    var v5 float32 = (0.1 * -2.25)
    v6 := []uint32{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if (!(float64((3.14159 * float32(1.5))) != -2.25)) {
        var v140 int = 847
        _ = v140
    } else {
        v4 = v4
        v3[2147483648] = ("hello" + ("" + ""))
    }
    var v141 float64 = (1.5 + ((((0.1 / float64(float32(3.14159))) + 1.5) / (-2.25 + 1.5)) * (3.14159 * 1e10)))
    v1 = v1
    v142 := GenS10{F0: (v4 ^ 32767), F1: 42, F2: 2147483648}
    v2 += v1
    var v143 uint64 = 158
    var v144 uint32 = 952
    fmt.Println("v1:", v1)
    fmt.Println("v141:", v141)
    fmt.Println("v143:", v143)
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v141
    _ = v142
    _ = v143
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
