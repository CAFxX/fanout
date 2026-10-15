package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64; F2 int16 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint64; F1 uint8 }
type GenS9 struct { F0 int32; F1 string }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS16 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS17 struct { F0 int; F1 int32 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int; F1 int8 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 bool; F1 int16 }
type GenS22 struct { F0 uint; F1 int16; F2 uint32 }
type GenS23 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 int32) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        switch "" {
        case "a":
            v5 := GenS1{F0: 42, F1: true, F2: 42}
            _ = v5
        case "foo bar":
            var v6 uint32 = 596
            _ = v6
        case "":
            v7 := []int{128, fi4, (fi4 * 2147483647)}
            _ = v7
        }
        fmt.Println(p2, p1, fi4)
        v8 := GenS2{F0: (p0 < p1)}
        _ = fi4
        _ = v8
    }
    p0 += p1
    var v9 uint16 = p0
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{(-(fi10)), fi10}
        var v12 int8 = 8
        v13 := GenS3{F0: 161, F1: 1, F2: -1}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    if false {
        v14 := GenS1{F0: 873, F1: (396 > 2147483647), F2: 42}
        if true {
            fmt.Println(p2, p1, p2)
            v15 := []int{}
            _ = v15
        } else {
            switch 2147483647 {
            case 552:
                p0 += (v9 + p0)
            case 0:
                fmt.Println(p0)
            case 32768:
                p2 += -2147483648
            default:
                v16 := []int{94, 1}
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 5; fi17++ {
            v9 = p0
            _ = fi17
        }
        _ = v14
    } else {
        var v18 uint32 = 301
        var v19 int = 42
        if (!(v19 <= v19)) {
            var v20 uint64 = 292
            var v21 int = v19
            if (("\n\t" < (((("foo bar" + "foo bar") + (("hello" + ("hello" + (("hello" + (("" + "") + ("\n\t" + (("\n\t" + ("hello" + "a")) + "a")))) + "a"))) + "")) + ("foo bar" + "hello")) + "foo bar")) || false) {
                fmt.Println(v18, v19)
                v22 := GenS4{F0: -128, F1: v20, F2: 1}
                if (v18 != (v18 | (v18 << ((0) & 63)))) {
                    fmt.Println(v20, v9)
                    for fi23 := 0; fi23 < 9; fi23++ {
                        fmt.Println(v18, v21)
                        switch 4294967295 {
                        case -1:
                            v20 += v20
                        case 4294967295:
                            if (1 != 733) {
                                var v24 uint8 = 128
                                if false {
                                    if (2 < 1) {
                                        v25 := []int{}
                                        if (!(0 != -136)) {
                                            switch "a" {
                                            case "":
                                                var v26 int16 = 541
                                                _ = v26
                                            case "hello":
                                                var v27 int64 = 877
                                                _ = v27
                                            }
                                            fmt.Println(p0, v19)
                                        } else {
                                            var v28 uint32 = v18
                                            _ = v28
                                        }
                                        var v29 int32 = (p2 | p2)
                                        _ = v25
                                        _ = v29
                                    } else {
                                        if (1 <= v20) {
                                            v24 = (^(v24))
                                        } else {
                                            v19 = 1
                                            for fi30 := 0; fi30 < 14; fi30++ {
                                                if false {
                                                    v31 := []int{(-(v21)), v21}
                                                    v32 := GenS3{F0: 4294967295, F1: 715, F2: 1000}
                                                    if ((0.0 * (0.0 + 1e10)) <= (1e10 * 0.0)) {
                                                        _ = len(v31)
                                                        switch (("foo bar" + "") + "foo bar") {
                                                        case "hello":
                                                            var v33 int64 = 70
                                                            _ = v33
                                                        case "":
                                                            switch v21 {
                                                            case -1:
                                                                fmt.Println(fi23)
                                                            }
                                                        case "a":
                                                            _ = len(v31)
                                                        default:
                                                            if len(v31) > 0 {
                                                                v31[0] = (v21 >> ((416) & 63))
                                                            }
                                                        }
                                                        if true {
                                                            if false {
                                                                p0 += (^(v9))
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v18 += v18
                                                                case "":
                                                                    fmt.Println(v9)
                                                                }
                                                                var v34 int = (v21 << ((255) & 63))
                                                                _ = v34
                                                            }
                                                        } else {
                                                            v35 := GenS4{F0: -128, F1: (v20 >> ((4294967295) & 63)), F2: 42}
                                                            var v36 uint8 = (v24 % ((v24) | 1))
                                                            var v37 uint = 18446744073709551615
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    } else {
                                                        v38 := GenS1{F0: 2147483648, F1: (!(42 > v20)), F2: 530}
                                                        _ = v38
                                                    }
                                                    v20 -= (v20 >> ((42) & 63))
                                                    var v39 uint16 = v9
                                                    _ = v31
                                                    _ = v32
                                                    _ = v39
                                                }
                                                v20 = (^(v20))
                                                if (!((p2 | p2) > p2)) {
                                                    var v40 int8 = 127
                                                    v19 = int(p2)
                                                    switch 127 {
                                                    case -2147483648:
                                                        p2 = (-(p2))
                                                    case 32768:
                                                        v21 = v19
                                                    case -41:
                                                        fmt.Println(v19, fi23, v18)
                                                    default:
                                                        v41 := []int{(fi30 - v21), int(p2), int(v24), (^(fi30))}
                                                        _ = v41
                                                    }
                                                    if true {
                                                        v42 := []int{(fi23 ^ (fi30 % (((v19 | 128)) | 1))), int(v20)}
                                                        v43 := GenS3{F0: 1, F1: -1, F2: 32767}
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        p1 = (p0 + (p1 + (p1 | p1)))
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        v44 := GenS2{F0: (v20 >= uint64(v24))}
                                                        _ = v44
                                                    default:
                                                        var v45 uint32 = v18
                                                        _ = v45
                                                    }
                                                    _ = v40
                                                }
                                                _ = fi30
                                            }
                                        }
                                        var v46 uint8 = v24
                                        _ = v46
                                    }
                                }
                                _ = v24
                            } else {
                                for fi47 := 0; fi47 < 9; fi47++ {
                                    p0 += (p0 - (p0 / (((v9 >> ((816) & 63))) | 1)))
                                    var v48 string = ("hello" + ("hello" + "\n\t"))
                                    if true {
                                        var v49 uint32 = (^(v18))
                                        switch v48 {
                                        case "\n\t":
                                            if true {
                                                v50 := []int{-712}
                                                p1 = p0
                                                p1 += v9
                                                _ = v50
                                            }
                                        case "hello":
                                            var v51 int8 = 127
                                            _ = v51
                                        }
                                        _ = v49
                                    } else {
                                        v52 := GenS4{F0: int8(fi47), F1: (v20 & (^(v20))), F2: 1}
                                        _ = v52
                                    }
                                    var v53 int16 = 0
                                    var v54 int64 = -102
                                    var v55 int16 = (v53 % ((1) | 1))
                                    _ = fi47
                                    _ = v48
                                    _ = v53
                                    _ = v54
                                    _ = v55
                                }
                                fmt.Println(v18, fi23)
                            }
                        default:
                            v56 := GenS5{F0: v18}
                            _ = v56
                        }
                        v9 += (p0 * (^(p0)))
                        _ = fi23
                    }
                    switch "a" {
                    case "":
                        for fi57 := 0; fi57 < 8; fi57++ {
                            v58 := []int{v19, int(p2), -9223372036854775808, int(v20)}
                            _ = fi57
                            _ = v58
                        }
                    default:
                        v59 := []int{v21, (v19 * (-(v19))), 128, (v19 + v19)}
                        _ = v59
                    }
                } else {
                    for fi60 := 0; fi60 < 6; fi60++ {
                        switch v21 {
                        case 4294967295:
                            v61 := GenS6{F0: true}
                            _ = v61
                        case 127:
                            if true {
                                fmt.Println(p0, v20)
                                var v62 int8 = 127
                                _ = v62
                            } else {
                                switch int(v18) {
                                case 42:
                                    var v63 string = "a"
                                    _ = v63
                                default:
                                    if ((!(!false)) && (127 != -34)) {
                                        switch (("" + ("a" + "hello")) + "hello") {
                                        case "hello":
                                            for fi64 := 0; fi64 < 11; fi64++ {
                                                v65 := []int{(v21 | (v21 & 9223372036854775807)), (fi60 / (((fi64 / (((v21 ^ (v21 + fi64))) | 1))) | 1))}
                                                fmt.Println(p2)
                                                _ = fi64
                                                _ = v65
                                            }
                                        default:
                                            fmt.Println(v18)
                                        }
                                        var v66 uint64 = 42
                                        fmt.Println(v18)
                                        _ = v66
                                    }
                                }
                                for fi67 := 0; fi67 < 15; fi67++ {
                                    fmt.Println(v20, p2)
                                    if ((!((1 > p0) && true)) || true) {
                                        v18 = v18
                                        v68 := []int{(^(fi67)), 4294967295, (^(v19)), (v19 / ((32767) | 1))}
                                        if false {
                                            if true {
                                                _ = len(v68)
                                                _ = len(v68)
                                            } else {
                                                var v69 uint64 = (v20 & v20)
                                                p0 = v9
                                                _ = v69
                                            }
                                            var v70 int8 = int8(p0)
                                            var v71 string = (("foo bar" + (("" + ("" + "foo bar")) + "hello")) + "foo bar")
                                            _ = v70
                                            _ = v71
                                        }
                                        switch (fi60 | -9223372036854775808) {
                                        case 127:
                                            fmt.Println(v19, p1, v21)
                                        default:
                                            if len(v68) > 0 {
                                                v68[0] = fi60
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        p2 = (p2 ^ int32(fi60))
                                        var v72 string = "foo bar"
                                        _ = v72
                                    }
                                    if true {
                                        if true {
                                            fmt.Println(p2, v18, v9)
                                            v73 := GenS4{F0: 1, F1: 783, F2: p1}
                                            _ = v73
                                        }
                                        if (true && (!(fi60 != (v21 ^ (-(fi67)))))) {
                                            var v74 uint8 = 127
                                            v75 := GenS1{F0: 2147483648, F1: true, F2: 0}
                                            v20 += (v20 * 2)
                                            _ = v74
                                            _ = v75
                                        }
                                        v76 := GenS4{F0: 28, F1: v20, F2: 98}
                                        var v77 uint16 = 42
                                        v78 := []int{int(p0), (fi60 >> ((2) & 63))}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                    _ = fi67
                                }
                                if true {
                                    switch v19 {
                                    case 128:
                                        fmt.Println(v20, p1, v18)
                                    case 9223372036854775807:
                                        switch 0 {
                                        case -1:
                                            var v79 uint32 = 972
                                            _ = v79
                                        case 310:
                                            fmt.Println(p1, v21, v21)
                                        default:
                                            if true {
                                                fmt.Println(fi60, p2, v18)
                                            } else {
                                                v80 := GenS3{F0: 42, F1: 9223372036854775807, F2: -454}
                                                var v81 uint32 = (v18 >> ((42) & 63))
                                                _ = v80
                                                _ = v81
                                            }
                                        }
                                    default:
                                        switch "a" {
                                        case "\n\t":
                                            fmt.Println(v20)
                                        case "a":
                                            for fi82 := 0; fi82 < 2; fi82++ {
                                                if true {
                                                    p2 = p2
                                                    var v83 int = v21
                                                    if false {
                                                        fmt.Println(v19, v19)
                                                    }
                                                    _ = v83
                                                }
                                                if (100 >= -128) {
                                                    fmt.Println(v9)
                                                }
                                                v20 = (v20 ^ 679)
                                                _ = fi82
                                            }
                                        case "":
                                            v84 := GenS7{F0: 781}
                                            _ = v84
                                        default:
                                            for fi85 := 0; fi85 < 8; fi85++ {
                                                var v86 bool = false
                                                switch (("\n\t" + "a") + ("foo bar" + "a")) {
                                                case "":
                                                    var v87 uint8 = 151
                                                    _ = v87
                                                case "foo bar":
                                                    if (p2 > 548) {
                                                        v88 := GenS8{F0: (v20 * (v20 % ((v20) | 1))), F1: 129}
                                                        if (0 != -611) {
                                                            var v89 uint64 = v20
                                                            _ = v89
                                                        } else {
                                                            var v90 bool = (57 <= 42)
                                                            _ = v90
                                                        }
                                                        var v91 int16 = int16(v9)
                                                        _ = v88
                                                        _ = v91
                                                    }
                                                default:
                                                    var v92 uint8 = 166
                                                    _ = v92
                                                }
                                                _ = fi85
                                                _ = v86
                                            }
                                        }
                                    }
                                    var v93 uint64 = v20
                                    _ = v93
                                }
                            }
                        case -9223372036854775808:
                            for fi94 := 0; fi94 < 7; fi94++ {
                                v95 := []int{-66, v19}
                                if false {
                                    switch ((("hello" + "hello") + "\n\t") + ("a" + "foo bar")) {
                                    case "hello":
                                        v95 = append(v95, (fi60 & fi94))
                                    }
                                    v18 = (v18 * (v18 - 4294967295))
                                }
                                fmt.Println(v18, fi94, v18)
                                v20 += v20
                                if len(v95) > 0 {
                                    v95[0] = (v21 >> ((0) & 63))
                                }
                                var v96 int32 = (p2 >> ((860) & 63))
                                _ = fi94
                                _ = v95
                                _ = v96
                            }
                        default:
                            fmt.Println(v19)
                        }
                        var v97 float32 = (1.5 - 3.14159)
                        v98 := GenS9{F0: (p2 ^ -557), F1: "foo bar"}
                        _ = fi60
                        _ = v97
                        _ = v98
                    }
                    for fi99 := 0; fi99 < 19; fi99++ {
                        v100 := GenS7{F0: 18446744073709551615}
                        _ = fi99
                        _ = v100
                    }
                    v101 := GenS10{F0: -97}
                    _ = v101
                }
                _ = v22
            } else {
                v102 := GenS5{F0: 65535}
                v103 := []int{}
                switch v19 {
                case 128:
                    var v104 int = (-(v21))
                    _ = v104
                case 2147483647:
                    switch 32767 {
                    case -348:
                        v105 := GenS11{F0: -9223372036854775808, F1: v20, F2: v18}
                        _ = v105
                    case -882:
                        v106 := GenS12{F0: (^(p2))}
                        _ = v106
                    }
                default:
                    v20 -= v20
                }
                _ = v102
                _ = v103
            }
            _ = v20
            _ = v21
        }
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}
func genFn1() uint64 {
    var v107 int = 32768
    if false {
        var v108 float32 = 1e10
        if true {
            for fi109 := 0; fi109 < 4; fi109++ {
                v110 := GenS13{F0: -1}
                for fi111 := 0; fi111 < 18; fi111++ {
                    switch v107 {
                    case -1:
                        var v112 int8 = -128
                        _ = v112
                    case 0:
                        fmt.Println(fi109, fi111)
                    case 4294967295:
                        var v113 int = 48
                        _ = v113
                    }
                    v114 := []int{(-(fi109))}
                    v107 = -1
                    _ = fi111
                    _ = v114
                }
                if (!((2147483647 >= 1) || (1.5 == 0.0))) {
                    switch "\n\t" {
                    case "\n\t":
                        if true {
                            var v115 uint16 = 255
                            _ = v115
                        } else {
                            if false {
                                fmt.Println(v107, fi109, v108)
                                v107 = (fi109 % ((v107) | 1))
                                fmt.Println(v108, v108, fi109)
                            } else {
                                for fi116 := 0; fi116 < 16; fi116++ {
                                    var v117 int16 = 0
                                    v108 = (float32((float64(1.5) - 0.0)) * v108)
                                    _ = fi116
                                    _ = v117
                                }
                            }
                            v118 := []int{(v107 & v107), -615}
                            _ = v118
                        }
                    case "":
                        if false {
                            for fi119 := 0; fi119 < 12; fi119++ {
                                v120 := []int{}
                                if false {
                                    var v121 uint32 = 2147483647
                                    _ = v121
                                }
                                var v122 int = fi109
                                _ = fi119
                                _ = v120
                                _ = v122
                            }
                            v108 = (v108 / (v108 + float32((1e10 / 3.14159))))
                            v108 = v108
                        } else {
                            fmt.Println(v108)
                        }
                    case "hello":
                        fmt.Println(fi109, fi109, fi109)
                    }
                }
                _ = fi109
                _ = v110
            }
            for fi123 := 0; fi123 < 20; fi123++ {
                v107 += fi123
                fmt.Println(v108, v108)
                switch v107 {
                case 128:
                    v108 = 0.1
                case 805:
                    v107 = 483
                }
                _ = fi123
            }
            if (v108 == 1.5) {
                if ((-18 != -1) && true) {
                    switch "" {
                    case "\n\t":
                        var v124 int32 = -2147483648
                        _ = v124
                    case "":
                        v108 = v108
                    }
                    v108 = (3.14159 / v108)
                    fmt.Println(v107)
                }
            }
        } else {
            v125 := []int{v107, v107, -447}
            _ = v125
        }
        _ = v108
    } else {
        v107 -= (v107 + v107)
        switch -47 {
        case 32768:
            v107 += v107
        case 541:
            var v126 string = ("foo bar" + (("a" + ("foo bar" + "a")) + "foo bar"))
            _ = v126
        case -490:
            for fi127 := 0; fi127 < 5; fi127++ {
                if (false || true) {
                    switch ("" + "hello") {
                    case "foo bar":
                        var v128 uint64 = 192
                        _ = v128
                    case "hello":
                        if false {
                            if false {
                                for fi129 := 0; fi129 < 6; fi129++ {
                                    if ((true && (true && ("hello" <= (("hello" + "") + ("hello" + "foo bar"))))) || false) {
                                        if true {
                                            v130 := []int{32768, (^(v107))}
                                            if false {
                                                var v131 bool = (!true)
                                                v132 := GenS10{F0: -1}
                                                _ = v131
                                                _ = v132
                                            }
                                            var v133 int16 = 1000
                                            _ = v130
                                            _ = v133
                                        } else {
                                            var v134 uint = 42
                                            v134 = 0
                                            switch "foo bar" {
                                            case "a":
                                                v134 = (v134 - (^(v134)))
                                            }
                                            _ = v134
                                        }
                                    }
                                    _ = fi129
                                }
                            }
                            v135 := []int{(fi127 ^ fi127)}
                            for fi136 := 0; fi136 < 17; fi136++ {
                                v135 = append(v135, v107)
                                v107 += (fi127 << ((862) & 63))
                                _ = fi136
                            }
                            _ = v135
                        }
                    case "a":
                        v107 = (fi127 % ((v107) | 1))
                    default:
                        for fi137 := 0; fi137 < 9; fi137++ {
                            if (!true) {
                                v107 = 599
                                v107 += (fi127 * 0)
                                var v138 string = ("hello" + "a")
                                _ = v138
                            } else {
                                var v139 string = "a"
                                _ = v139
                            }
                            v107 = -9223372036854775808
                            if (2 <= 65535) {
                                switch "" {
                                case "a":
                                    v140 := []int{}
                                    v107 -= -9223372036854775808
                                    var v141 uint32 = 2
                                    _ = v140
                                    _ = v141
                                case "":
                                    if (!false) {
                                        if false {
                                            if true {
                                                v107 = (fi137 % (((fi127 << ((65535) & 63))) | 1))
                                            } else {
                                                if (((v107 & (fi127 ^ fi137)) == v107) && (false && ("hello" != ((("\n\t" + "foo bar") + "\n\t") + "")))) {
                                                    if ((1 != 127) || (!(true && false))) {
                                                        v142 := []int{127, (fi137 * (fi127 & -2147483648))}
                                                        if (!true) {
                                                            v107 = (fi127 ^ (-(v107)))
                                                        }
                                                        _ = v142
                                                    } else {
                                                        v107 -= -1
                                                    }
                                                    var v143 int8 = -128
                                                    var v144 string = ("a" + "\n\t")
                                                    _ = v143
                                                    _ = v144
                                                }
                                            }
                                            v107 -= fi137
                                            if false {
                                                if true {
                                                    if false {
                                                        v107 = 42
                                                        v107 += -9223372036854775808
                                                    }
                                                } else {
                                                    v107 = v107
                                                    fmt.Println(fi127, v107)
                                                    if false {
                                                        v107 -= fi127
                                                    }
                                                }
                                            } else {
                                                v107 = 4294967295
                                                if ((0.1 != -2.25) && false) {
                                                    v145 := GenS3{F0: 4294967295, F1: -1, F2: -686}
                                                    v146 := GenS14{F0: 2}
                                                    _ = v145
                                                    _ = v146
                                                } else {
                                                    v147 := GenS15{F0: 65535, F1: 2, F2: 1}
                                                    var v148 uint8 = 132
                                                    _ = v147
                                                    _ = v148
                                                }
                                            }
                                            v149 := []int{(^(fi127))}
                                            _ = v149
                                        }
                                        v107 = fi127
                                        var v150 int32 = 2147483647
                                        _ = v150
                                    } else {
                                        fmt.Println(fi137, fi137, fi137)
                                        v107 = (fi127 ^ -483)
                                    }
                                case "\n\t":
                                    var v151 int = (fi127 * fi127)
                                    _ = v151
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                v107 = -449
                            case "foo bar":
                                v152 := []int{v107, (fi127 & fi137), fi137, (v107 & (fi127 << ((501) & 63)))}
                                _ = v152
                            case "a":
                                v153 := GenS16{F0: 1, F1: 65535, F2: -32768}
                                _ = v153
                            default:
                                var v154 int32 = 887
                                _ = v154
                            }
                            _ = fi137
                        }
                    }
                }
                for fi155 := 0; fi155 < 16; fi155++ {
                    switch "hello" {
                    case "\n\t":
                        var v156 bool = ((!true) && (!(1.5 != 0.1)))
                        _ = v156
                    case "hello":
                        v157 := GenS17{F0: 128, F1: -2147483648}
                        _ = v157
                    }
                    switch ("" + ("foo bar" + "hello")) {
                    case "hello":
                        fmt.Println(fi155, fi127)
                    case "\n\t":
                        fmt.Println(fi155, v107, fi127)
                    }
                    v158 := GenS18{F0: ""}
                    _ = fi155
                    _ = v158
                }
                v107 = fi127
                _ = fi127
            }
        default:
            var v159 float32 = float32(0.1)
            _ = v159
        }
        var v160 float32 = 0.1
        _ = v160
    }
    v161 := []int{850, v107}
    _ = v107
    _ = v161
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v162 int16 = -261
    var v163 uint8 = 145
    v164 := GenS13{F0: -1}
    v165 := v3[2147483647]
    if true {
        var v166 int64 = -1
        if false {
            var v167 int16 = v162
            _ = v167
        } else {
            var v168 float64 = 3.14159
            v169 := GenS16{F0: 1, F1: 65535, F2: 1000}
            v170 := v2["\n\t"]
            _ = v168
            _ = v169
            _ = v170
        }
        if false {
            switch 32768 {
            case 128:
                if (!false) {
                    if (v166 > v166) {
                        v163 -= (^(v163))
                        switch v165 {
                        case "a":
                            v171 := GenS19{F0: 32768, F1: 0}
                            _ = v171
                        case "foo bar":
                            switch v165 {
                            case "foo bar":
                                var v172 uint = 18446744073709551615
                                _ = v172
                            case "":
                                v173 := GenS16{F0: 1, F1: 2147483648, F2: v162}
                                _ = v173
                            case "a":
                                fmt.Println(v163)
                            }
                        default:
                            v163 += v163
                        }
                        var v174 string = ("\n\t" + v165)
                        _ = v174
                    }
                    v175 := GenS13{F0: (^(v166))}
                    fmt.Println(v163, v165, v162)
                    _ = v175
                } else {
                    var v176 uint = 1
                    _ = v176
                }
            case 370:
                switch v165 {
                case "\n\t":
                    if true {
                        var v177 uint8 = (^(v163))
                        var v178 float64 = (3.14159 / 1.0)
                        for fi179 := 0; fi179 < 2; fi179++ {
                            v180 := []int{fi179, (^(fi179)), fi179}
                            _ = fi179
                            _ = v180
                        }
                        _ = v177
                        _ = v178
                    } else {
                        var v181 int64 = v166
                        v1[v165] = true
                        for fi182 := 0; fi182 < 3; fi182++ {
                            fmt.Println(v163, fi182, fi182)
                            _ = fi182
                        }
                        _ = v181
                    }
                case "hello":
                    for fi183 := 0; fi183 < 2; fi183++ {
                        v165 = v165
                        if (42 <= v166) {
                            if false {
                                for fi184 := 0; fi184 < 4; fi184++ {
                                    var v185 bool = false
                                    if v185 {
                                        v186 := []int{fi183, (fi184 ^ fi184), 758, -893}
                                        switch 443 {
                                        case 1:
                                            if true {
                                                fmt.Println(fi183)
                                                var v187 uint8 = (v163 + 1)
                                                if ((!(v185 && (0 == 139))) && false) {
                                                    v165 = v165
                                                } else {
                                                    switch (fi183 | (-(fi183))) {
                                                    case -877:
                                                        var v188 float32 = float32(3.14159)
                                                        _ = v188
                                                    case -2147483648:
                                                        var v189 int16 = (-(v162))
                                                        _ = v189
                                                    default:
                                                        switch fi183 {
                                                        case 724:
                                                            if len(v186) > 0 {
                                                                v186[0] = fi184
                                                            }
                                                        case 0:
                                                            fmt.Println(v163, v162, fi184)
                                                        default:
                                                            switch ("hello" + "hello") {
                                                            case "\n\t":
                                                                if (!v185) {
                                                                    v166 += v166
                                                                    v190 := GenS20{F0: true}
                                                                    v187 -= (^(v187))
                                                                    v191 := GenS6{F0: (false || ((-(v166)) > (v166 | (v166 % ((v166) | 1)))))}
                                                                    _ = v190
                                                                    _ = v191
                                                                }
                                                            case "hello":
                                                                if v185 {
                                                                    if len(v186) >= 2 {
                                                                        v186 = v186[1:len(v186)]
                                                                    }
                                                                    v192 := GenS18{F0: "\n\t"}
                                                                    _ = v192
                                                                } else {
                                                                    var v193 float32 = 0.0
                                                                    var v194 uint8 = v187
                                                                    if v185 {
                                                                        var v195 bool = (218 != 18446744073709551615)
                                                                        _ = len(v186)
                                                                        _ = v195
                                                                    } else {
                                                                        var v196 uint32 = 42
                                                                        v187 += v163
                                                                        if v185 {
                                                                            if (127 >= 127) {
                                                                                var v197 bool = false
                                                                                _ = v197
                                                                            }
                                                                        }
                                                                        _ = v196
                                                                    }
                                                                    if (fi184 >= fi183) {
                                                                        var v198 uint32 = 0
                                                                        switch 712 {
                                                                        case 128:
                                                                            if (!(((float32(-2.25) + (v193 - ((-2.25 - float32(1e10)) / v193))) - 0.1) <= float32(0.1))) {
                                                                                v165 = ""
                                                                            } else {
                                                                                if v185 {
                                                                                    if (1 > 18446744073709551615) {
                                                                                        var v199 int64 = (v166 ^ v166)
                                                                                        v193 = v193
                                                                                        fmt.Println(v162, v198)
                                                                                        _ = len(v186)
                                                                                        _ = v199
                                                                                    }
                                                                                } else {
                                                                                    switch -2147483648 {
                                                                                    case -9223372036854775808:
                                                                                        var v200 int16 = -479
                                                                                        _ = v200
                                                                                    case -830:
                                                                                        v201 := GenS21{F0: (!v185), F1: int16(v198)}
                                                                                        _ = v201
                                                                                    }
                                                                                    v202 := GenS14{F0: 122}
                                                                                    v198 += v198
                                                                                    fmt.Println(v185)
                                                                                    _ = v202
                                                                                }
                                                                                if (v166 < (v166 * (-(v166)))) {
                                                                                    v198 = v198
                                                                                    if false {
                                                                                        if false {
                                                                                            v163 = v163
                                                                                            v203 := GenS10{F0: 42}
                                                                                            var v204 uint64 = 18446744073709551615
                                                                                            _ = v203
                                                                                            _ = v204
                                                                                        } else {
                                                                                            fmt.Println(fi183)
                                                                                        }
                                                                                        if ((v185 && ((v162 + v162) < v162)) && true) {
                                                                                            v205 := GenS7{F0: 42}
                                                                                            v206 := GenS22{F0: 65535, F1: v162, F2: (^(v198))}
                                                                                            _ = v205
                                                                                            _ = v206
                                                                                        }
                                                                                    } else {
                                                                                        var v207 uint64 = 2
                                                                                        _ = v207
                                                                                    }
                                                                                }
                                                                                v208 := GenS22{F0: 1, F1: (v162 / ((v162) | 1)), F2: v198}
                                                                                fmt.Println(fi184, fi183)
                                                                                switch v165 {
                                                                                case "":
                                                                                    var v209 bool = (101 == v194)
                                                                                    _ = v209
                                                                                }
                                                                                _ = v208
                                                                            }
                                                                        case 2147483647:
                                                                            _ = len(v2)
                                                                        default:
                                                                            var v210 float64 = float64(v193)
                                                                            _ = v210
                                                                        }
                                                                        if false {
                                                                            v194 -= v163
                                                                            var v211 string = v165
                                                                            if (true && v185) {
                                                                                var v212 int8 = 42
                                                                                fmt.Println(fi184, v185, v163)
                                                                                _ = v212
                                                                            } else {
                                                                                switch v211 {
                                                                                case "foo bar":
                                                                                    var v213 float32 = float32(1.5)
                                                                                    _ = v213
                                                                                }
                                                                            }
                                                                            _ = v211
                                                                        } else {
                                                                            v163 = (v163 >> ((226) & 63))
                                                                            fmt.Println(v198)
                                                                            v194 = v187
                                                                        }
                                                                        _ = v198
                                                                    }
                                                                    v162 = (v162 * -374)
                                                                    _ = v193
                                                                    _ = v194
                                                                }
                                                            default:
                                                                switch (v165 + (v165 + (v165 + (v165 + "")))) {
                                                                case "":
                                                                    v214 := GenS22{F0: 18446744073709551615, F1: (v162 + (v162 * (v162 ^ -1))), F2: 0}
                                                                    _ = v214
                                                                default:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if (v162 >= -1) {
                                                                            v166 = v166
                                                                            _ = len(v1)
                                                                            v163 += 42
                                                                        }
                                                                    case "a":
                                                                        if (978 == 113) {
                                                                            var v215 float64 = 3.14159
                                                                            fmt.Println(v163, fi183)
                                                                            v187 = v163
                                                                            _ = v215
                                                                        } else {
                                                                            var v216 uint64 = 900
                                                                            v217 := GenS8{F0: (v216 / ((42) | 1)), F1: (v187 ^ 0)}
                                                                            var v218 bool = (0 == 761)
                                                                            _ = v216
                                                                            _ = v217
                                                                            _ = v218
                                                                        }
                                                                    case "foo bar":
                                                                        if len(v186) > 0 {
                                                                            v186[0] = fi183
                                                                        }
                                                                    }
                                                                    switch v165 {
                                                                    case "a":
                                                                        fmt.Println(v162)
                                                                    }
                                                                    if v185 {
                                                                        v163 += (v163 ^ v187)
                                                                        fmt.Println(v185, v166, fi183)
                                                                    }
                                                                }
                                                                if false {
                                                                    v187 = v187
                                                                    var v219 int64 = v166
                                                                    v1[(v165 + "")] = v185
                                                                    _ = v219
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v162 = v162
                                                }
                                                v163 = (v163 & v163)
                                                _ = v187
                                            }
                                        case 0:
                                            switch "" {
                                            case "":
                                                if len(v186) >= 2 {
                                                    v186 = v186[1:len(v186)]
                                                }
                                            default:
                                                v185 = true
                                            }
                                        default:
                                            delete(v2, v165)
                                        }
                                        v163 += v163
                                        fmt.Println(v162, v185)
                                        if (v185 || (v166 <= (v166 - (v166 & int64(v163))))) {
                                            var v220 uint64 = genFn1()
                                            var v221 bool = v185
                                            delete(v1, v165)
                                            _ = v220
                                            _ = v221
                                        } else {
                                            if (1.5 != 1.5) {
                                                var v222 float32 = 3.14159
                                                if len(v186) >= 2 {
                                                    v186 = v186[1:len(v186)]
                                                }
                                                fmt.Println(v185)
                                                _ = v222
                                            }
                                        }
                                        _ = v186
                                    }
                                    var v223 int16 = (v162 >> ((42) & 63))
                                    _ = fi184
                                    _ = v185
                                    _ = v223
                                }
                            }
                            delete(v2, (v165 + (("" + (v165 + v165)) + v165)))
                        } else {
                            v163 -= v163
                            var v224 int = -1
                            fmt.Println(fi183, v166, v224)
                            _ = v224
                        }
                        _ = fi183
                    }
                case "a":
                    v165 = (("foo bar" + ((v165 + "foo bar") + v165)) + v165)
                }
            case 9223372036854775807:
                var v225 float32 = (-2.25 / 0.1)
                _ = v225
            }
            var v226 int64 = (v166 % ((-80) | 1))
            v227 := GenS5{F0: 0}
            _ = v226
            _ = v227
        } else {
            v163 = 2
            for fi228 := 0; fi228 < 20; fi228++ {
                fmt.Println(v166, v166, v166)
                fmt.Println(v166, v165, fi228)
                fmt.Println(v165)
                _ = fi228
            }
        }
        _ = v166
    } else {
        v229 := GenS6{F0: ((!(v165 != v165)) || false)}
        for fi230 := 0; fi230 < 2; fi230++ {
            if false {
                v165 = (v165 + ("\n\t" + v165))
                var v231 string = (((("\n\t" + v165) + (v165 + (v165 + (v165 + "\n\t")))) + v165) + v165)
                _ = v231
            }
            var v232 uint8 = v163
            _ = fi230
            _ = v232
        }
        v233 := GenS18{F0: ((("foo bar" + v165) + v165) + ((v165 + v165) + ("" + (v165 + ((((v165 + ((v165 + ("a" + "foo bar")) + v165)) + v165) + v165) + "\n\t")))))}
        _ = v229
        _ = v233
    }
    v165 = v165
    _ = len(v1)
    switch -931 {
    case 32767:
        for fi234 := 0; fi234 < 7; fi234++ {
            var v235 uint64 = 42
            _ = fi234
            _ = v235
        }
    }
    v162 -= v162
    v163 += 2
    var v236 uint = 255
    if false {
        var v237 int64 = -9223372036854775808
        v238 := GenS23{F0: 18446744073709551615, F1: 264}
        _ = v237
        _ = v238
    }
    var v239 int = -345
    fmt.Println(v236)
    switch -2147483648 {
    case -359:
        switch "\n\t" {
        case "":
            fmt.Println(v162, v165)
        case "a":
            v239 -= (v239 ^ 277)
        }
    default:
        v239 = v239
    }
    var v240 uint8 = 42
    for fi241 := 0; fi241 < 17; fi241++ {
        var v242 float32 = -2.25
        _ = fi241
        _ = v242
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v162:", v162)
    fmt.Println("v163:", v163)
    fmt.Println("v165:", v165)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v236:", v236)
    fmt.Println("v239:", v239)
    fmt.Println("v240:", v240)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v162
    _ = v163
    _ = v164
    _ = v165
    _ = v2
    _ = v236
    _ = v239
    _ = v240
    _ = v3
}
