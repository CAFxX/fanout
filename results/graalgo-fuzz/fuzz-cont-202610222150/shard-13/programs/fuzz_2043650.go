package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 string; F2 int64 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 uint16; F1 int64; F2 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int8; F1 uint16; F2 int8 }
type GenS12 struct { F0 int64; F1 bool; F2 bool }
type GenS13 struct { F0 bool; F1 string; F2 int64 }
type GenS14 struct { F0 uint16; F1 uint32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 int8 }
type GenS17 struct { F0 uint64 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint64; F1 uint64 }
type GenS20 struct { F0 int8; F1 uint64 }
type GenS21 struct { F0 int16; F1 int8 }
type GenS22 struct { F0 uint8 }
type GenS23 struct { F0 uint64; F1 int8; F2 uint }
type GenS24 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS25 struct { F0 int64; F1 uint }
func genFn0(p0 uint32) bool {
    var v5 uint = 1
    v5 += (v5 << (((v5 | (v5 * v5))) & 63))
    _ = p0
    _ = v5
    return true
}
func genFn1(p0 bool, p1 int16, p2 uint8) uint64 {
    p1 = (-(p1))
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{}
        p1 = -1
        switch "foo bar" {
        case "a":
            for fi8 := 0; fi8 < 16; fi8++ {
                if len(v7) > 0 {
                    v7[0] = fi6
                }
                _ = fi8
            }
        case "\n\t":
            var v9 uint16 = 0
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int32 = 359
    var v3 uint = 492
    var v4 uint64 = uint64(v3)
    if ((("foo bar" + "a") >= ("hello" + "")) && false) {
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            if true {
                if true {
                    var v10 uint16 = 255
                    var v11 uint32 = 2
                    v1[v11] = ("" + "foo bar")
                    _ = v10
                    _ = v11
                }
            }
        case "hello":
            v12 := []int{32768}
            _ = v12
        }
        if true {
            switch 9223372036854775807 {
            case -893:
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := []int{fi13, (fi13 / ((893) | 1)), 32768, fi13}
                    if (-32768 != 0) {
                        switch "" {
                        case "hello":
                            fmt.Println(v3, v4, fi13)
                        case "a":
                            v3 += (v3 * v3)
                        default:
                            var v15 uint64 = 667
                            _ = v15
                        }
                        var v16 uint = 4294967295
                        var v17 uint64 = v4
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                }
            case 1:
                if (((1e10 * ((1.5 / -2.25) / (1e10 + 0.1))) > 1e10) && ((1 <= 2147483648) && ((true && (v3 != v3)) || (!(434 >= 254))))) {
                    if (!false) {
                        v3 = 0
                        switch "a" {
                        case "\n\t":
                            switch "\n\t" {
                            case "a":
                                if (float32((1.5 / -2.25)) == 0.0) {
                                    var v18 uint = v3
                                    if (((("foo bar" + "a") + "hello") + ("" + "hello")) > (((("a" + ("" + "foo bar")) + "a") + (("a" + "hello") + "\n\t")) + "a")) {
                                        for fi19 := 0; fi19 < 1; fi19++ {
                                            for fi20 := 0; fi20 < 17; fi20++ {
                                                if (!(!true)) {
                                                    v21 := GenS1{F0: 42}
                                                    if true {
                                                        if true {
                                                            v22 := GenS2{F0: -128, F1: (("hello" + ("hello" + ("foo bar" + ""))) + "")}
                                                            v23 := []int{127, (fi19 + fi19), 2147483647, fi19}
                                                            _ = v22
                                                            _ = v23
                                                        } else {
                                                            fmt.Println(v4, v3)
                                                        }
                                                    } else {
                                                        if ((!false) && (!true)) {
                                                            v24 := GenS3{F0: 0}
                                                            fmt.Println(fi19)
                                                            fmt.Println(v3, v3, v2)
                                                            if (!true) {
                                                                var v25 uint16 = 734
                                                                if true {
                                                                    v25 += v25
                                                                    v26 := []int{fi19, (fi20 % ((fi19) | 1)), (fi19 * -792), fi20}
                                                                    var v27 int32 = v2
                                                                    _ = v26
                                                                    _ = v27
                                                                }
                                                                _ = v25
                                                            }
                                                            _ = v24
                                                        }
                                                        v28 := GenS4{F0: 42}
                                                        _ = v28
                                                    }
                                                    var v29 float64 = 0.1
                                                    _ = v21
                                                    _ = v29
                                                } else {
                                                    var v30 uint = (v3 ^ (^(v3)))
                                                    v31 := []int{fi20, (fi20 | (fi20 | -1)), (fi20 << ((v18) & 63)), 769}
                                                    switch (("" + (((("\n\t" + ("\n\t" + (("foo bar" + "a") + "a"))) + ("foo bar" + "foo bar")) + "a") + "")) + "hello") {
                                                    case "\n\t":
                                                        delete(v1, 42)
                                                    }
                                                    v1[65535] = "a"
                                                    _ = v30
                                                    _ = v31
                                                }
                                                _ = fi20
                                            }
                                            _ = fi19
                                        }
                                        v3 = (v3 ^ (^(v18)))
                                        var v32 uint = v18
                                        _ = v32
                                    } else {
                                        var v33 int32 = v2
                                        _ = v33
                                    }
                                    _ = v18
                                } else {
                                    fmt.Println(v3, v2, v4)
                                }
                            case "\n\t":
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    if (128 >= 128) {
                                        delete(v1, 65535)
                                    } else {
                                        v1[uint32(v4)] = "a"
                                        v4 = (v4 | v4)
                                    }
                                case -2147483648:
                                    var v34 int32 = -2147483648
                                    _ = v34
                                }
                            case "":
                                if true {
                                    fmt.Println(v3, v3, v3)
                                    for fi35 := 0; fi35 < 15; fi35++ {
                                        v36 := GenS5{F0: (v4 | v4), F1: "", F2: 9223372036854775807}
                                        v37 := GenS1{F0: 575}
                                        var v38 uint = (v3 >> (((v3 % (((v3 ^ v3)) | 1))) & 63))
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    switch "hello" {
                                    case "":
                                        v4 = (v4 ^ genFn1(false, -519, 128))
                                    case "foo bar":
                                        var v39 uint16 = 65535
                                        _ = v39
                                    }
                                }
                            }
                        case "foo bar":
                            v40 := GenS6{F0: 65535, F1: v3}
                            _ = v40
                        }
                    }
                } else {
                    v41 := []int{int(v3), 9223372036854775807, 0}
                    switch -974 {
                    case -253:
                        v42 := GenS1{F0: 42}
                        _ = v42
                    case -9223372036854775808:
                        var v43 uint16 = 65535
                        _ = v43
                    case -134:
                        v4 -= (v4 - (v4 << ((4294967295) & 63)))
                    default:
                        for fi44 := 0; fi44 < 15; fi44++ {
                            if true {
                                var v45 float32 = (-2.25 * (0.0 / -2.25))
                                v3 = (v3 / ((0) | 1))
                                var v46 float32 = 0.0
                                _ = v45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 13; fi47++ {
                                fmt.Println(fi47, fi47, v3)
                                v4 = (v4 | 4294967295)
                                _ = fi47
                            }
                            fmt.Println(v2, fi44)
                            _ = fi44
                        }
                    }
                    _ = v41
                }
            case 0:
                v3 += v3
            }
        }
        v48 := v1[2]
        _ = v48
    }
    if false {
        if ((v2 << ((v3) & 63)) != (-(v2))) {
            var v49 float32 = 1e10
            var v50 uint32 = 309
            _ = v49
            _ = v50
        }
        var v51 float32 = 1e10
        _ = v51
    }
    var v52 uint64 = (^(v4))
    fmt.Println(v4, v52)
    for fi53 := 0; fi53 < 15; fi53++ {
        v54 := GenS7{F0: 65535, F1: 0, F2: 1}
        if ((0 > 0) && (true && (float64(((0.1 + 3.14159) / float32(3.14159))) == (float64(0.0) / 1.5)))) {
            for fi55 := 0; fi55 < 2; fi55++ {
                v56 := GenS8{F0: true}
                _ = fi55
                _ = v56
            }
            fmt.Println(v3, fi53)
            fmt.Println(v52)
        }
        for fi57 := 0; fi57 < 15; fi57++ {
            var v58 uint16 = 252
            var v59 uint8 = 126
            switch "a" {
            case "a":
                v52 = (v52 / (((v52 / (((v4 + (v4 / (((v4 >> (((^(v3))) & 63))) | 1)))) | 1))) | 1))
            case "":
                var v60 int64 = -9223372036854775808
                _ = v60
            case "hello":
                fmt.Println(fi53, v4, v59)
            }
            _ = fi57
            _ = v58
            _ = v59
        }
        _ = fi53
        _ = v54
    }
    fmt.Println(v52, v52, v2)
    switch (("foo bar" + "\n\t") + (("hello" + "") + ((("" + "foo bar") + "hello") + "foo bar"))) {
    case "":
        v61 := GenS5{F0: 4294967295, F1: "\n\t", F2: 9223372036854775807}
        _ = v61
    case "\n\t":
        v62 := []int{32767, int(v52)}
        _ = v62
    case "hello":
        switch 0 {
        case 4294967295:
            if true {
                fmt.Println(v3, v3)
                v2 -= v2
                for fi63 := 0; fi63 < 11; fi63++ {
                    v64 := GenS5{F0: v52, F1: "", F2: 1}
                    v65 := GenS1{F0: 0}
                    v66 := GenS2{F0: 100, F1: "hello"}
                    _ = fi63
                    _ = v64
                    _ = v65
                    _ = v66
                }
            } else {
                v67 := GenS8{F0: (0 > 127)}
                if (!(42 >= 1000)) {
                    if (0 == 2147483648) {
                        v4 += (v52 * (v52 - v52))
                        if (!(((!false) || true) || true)) {
                            if (!false) {
                                fmt.Println(v3)
                                v68 := []int{42, 32768}
                                switch "hello" {
                                case "hello":
                                    v4 += genFn1((false || (42 != 0)), 1000, 42)
                                case "foo bar":
                                    fmt.Println(v4, v52, v52)
                                case "a":
                                    v2 -= (v2 * v2)
                                default:
                                    switch ("\n\t" + ("" + "a")) {
                                    case "hello":
                                        for fi69 := 0; fi69 < 16; fi69++ {
                                            var v70 int8 = -29
                                            _ = fi69
                                            _ = v70
                                        }
                                    case "\n\t":
                                        var v71 uint8 = 255
                                        _ = v71
                                    default:
                                        for fi72 := 0; fi72 < 5; fi72++ {
                                            for fi73 := 0; fi73 < 2; fi73++ {
                                                if true {
                                                    fmt.Println(v52, v52)
                                                    v52 = (v4 + 0)
                                                    v4 = 2
                                                    if false {
                                                        var v74 int16 = -496
                                                        v75 := GenS2{F0: 42, F1: ("foo bar" + "foo bar")}
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    var v76 uint32 = 42
                                                    _ = v76
                                                } else {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v77 int32 = (v2 & v2)
                                                        _ = v77
                                                    case "a":
                                                        v78 := GenS9{F0: (v3 * v3), F1: (^(v3))}
                                                        _ = v78
                                                    default:
                                                        switch 128 {
                                                        case 122:
                                                            if len(v68) >= 2 {
                                                                v68 = v68[1:len(v68)]
                                                            }
                                                        default:
                                                            v79 := v1[0]
                                                            _ = v79
                                                        }
                                                    }
                                                }
                                                _ = fi73
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                var v80 int32 = v2
                                                _ = v80
                                            default:
                                                if len(v68) > 0 {
                                                    v68[0] = fi72
                                                }
                                            }
                                            _ = fi72
                                        }
                                    }
                                }
                                _ = v68
                            } else {
                                fmt.Println(v4)
                                fmt.Println(v4, v2)
                            }
                            if false {
                                if true {
                                    var v81 uint64 = 42
                                    var v82 uint = v3
                                    _ = v81
                                    _ = v82
                                }
                                if true {
                                    if true {
                                        if false {
                                            fmt.Println(v3, v2, v4)
                                            for fi83 := 0; fi83 < 18; fi83++ {
                                                v3 -= v3
                                                v84 := GenS7{F0: 942, F1: 42, F2: fi83}
                                                if (true || true) {
                                                    v52 += (^(v4))
                                                    fmt.Println(v4)
                                                } else {
                                                    for fi85 := 0; fi85 < 9; fi85++ {
                                                        v1[42] = (("a" + (("hello" + "hello") + "\n\t")) + "foo bar")
                                                        v86 := GenS7{F0: 323, F1: 9223372036854775807, F2: (fi85 / (((fi85 / (((fi85 << (((v3 * v3)) & 63))) | 1))) | 1))}
                                                        var v87 int16 = -1
                                                        switch fi83 {
                                                        case 32767:
                                                            if false {
                                                                fmt.Println(fi83, fi85)
                                                                fmt.Println(v52)
                                                                v4 -= v52
                                                                v52 = v4
                                                            } else {
                                                                if false {
                                                                    if ((!((false || true) || false)) && true) {
                                                                        if (false || false) {
                                                                            v88 := GenS5{F0: (v52 + (v52 << (((v3 * 910)) & 63))), F1: ("" + ""), F2: -59}
                                                                            _ = v88
                                                                        } else {
                                                                            v2 += 0
                                                                        }
                                                                    }
                                                                    v89 := v1[uint32(v3)]
                                                                    _ = v89
                                                                } else {
                                                                    switch ("" + ("" + (("\n\t" + "") + "hello"))) {
                                                                    case "":
                                                                        v90 := []int{(fi85 & -2147483648), fi83}
                                                                        _ = v90
                                                                    case "a":
                                                                        if ((v3 % ((231) | 1)) > v3) {
                                                                            _ = len(v1)
                                                                            if false {
                                                                                v87 = (v87 % ((v87) | 1))
                                                                                v91 := []int{(fi83 ^ (fi83 & fi85))}
                                                                                _ = v91
                                                                            } else {
                                                                                switch 854 {
                                                                                case 983:
                                                                                    var v92 int8 = 1
                                                                                    _ = v92
                                                                                default:
                                                                                    var v93 uint16 = 42
                                                                                    _ = v93
                                                                                }
                                                                            }
                                                                            v94 := GenS10{F0: v87, F1: (^(v2))}
                                                                            var v95 string = "hello"
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                    case "hello":
                                                                        v96 := GenS3{F0: -902}
                                                                        _ = v96
                                                                    default:
                                                                        var v97 uint16 = 434
                                                                        _ = v97
                                                                    }
                                                                    if true {
                                                                        var v98 int16 = v87
                                                                        switch fi85 {
                                                                        case 127:
                                                                            var v99 uint32 = 1
                                                                            _ = v99
                                                                        case 32768:
                                                                            var v100 int64 = -1
                                                                            _ = v100
                                                                        }
                                                                        _ = v98
                                                                    } else {
                                                                        if true {
                                                                            if (-101 > (fi83 << ((v3) & 63))) {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        if false {
                                                                                            v101 := GenS11{F0: -1, F1: 0, F2: -1}
                                                                                            var v102 int16 = -32768
                                                                                            v2 += 865
                                                                                            _ = v101
                                                                                            _ = v102
                                                                                        } else {
                                                                                            switch (fi85 | 128) {
                                                                                            case 991:
                                                                                                fmt.Println(fi83)
                                                                                            case -433:
                                                                                                var v103 bool = false
                                                                                                _ = v103
                                                                                            default:
                                                                                                switch fi85 {
                                                                                                case -597:
                                                                                                    var v104 bool = ((^(v87)) != v87)
                                                                                                    _ = v104
                                                                                                case -583:
                                                                                                    var v105 string = "hello"
                                                                                                    _ = v105
                                                                                                }
                                                                                            }
                                                                                            v87 = (v87 | v87)
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    var v106 float32 = (-2.25 * ((0.1 + -2.25) / (1.5 - (3.14159 * 0.1))))
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        switch (("a" + "foo bar") + ("foo bar" + "foo bar")) {
                                                                                        case "hello":
                                                                                            var v107 float32 = 3.14159
                                                                                            _ = v107
                                                                                        case "\n\t":
                                                                                            v4 -= 618
                                                                                        default:
                                                                                            v108 := v1[2147483647]
                                                                                            _ = v108
                                                                                        }
                                                                                    }
                                                                                    var v109 uint8 = 2
                                                                                    _ = v106
                                                                                    _ = v109
                                                                                }
                                                                                if true {
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if false {
                                                                                            if true {
                                                                                                if (!false) {
                                                                                                    _ = len(v1)
                                                                                                    v110 := GenS2{F0: 0, F1: ""}
                                                                                                    _ = v110
                                                                                                } else {
                                                                                                    if false {
                                                                                                        v2 = (^(v2))
                                                                                                        if (0 != -594) {
                                                                                                            switch fi85 {
                                                                                                            case -2147483648:
                                                                                                                var v111 uint = (v3 / ((v3) | 1))
                                                                                                                _ = v111
                                                                                                            default:
                                                                                                                switch "foo bar" {
                                                                                                                case "foo bar":
                                                                                                                    switch (fi85 * (fi85 << ((v3) & 63))) {
                                                                                                                    case 4294967295:
                                                                                                                        v112 := []int{}
                                                                                                                        _ = v112
                                                                                                                    case -1:
                                                                                                                        v113 := GenS12{F0: 42, F1: ((fi83 / ((fi83) | 1)) == fi85), F2: (v52 <= genFn1((true || true), v87, 127))}
                                                                                                                        _ = v113
                                                                                                                    }
                                                                                                                default:
                                                                                                                    var v114 uint16 = 255
                                                                                                                    _ = v114
                                                                                                                }
                                                                                                            }
                                                                                                            var v115 uint16 = 2
                                                                                                            _ = v115
                                                                                                        } else {
                                                                                                            _ = len(v1)
                                                                                                            switch (fi83 & (fi83 & fi85)) {
                                                                                                            case 9223372036854775807:
                                                                                                                v116 := GenS13{F0: true, F1: ("hello" + ""), F2: -9223372036854775808}
                                                                                                                _ = v116
                                                                                                            default:
                                                                                                                var v117 int32 = int32(v3)
                                                                                                                _ = v117
                                                                                                            }
                                                                                                        }
                                                                                                        if true {
                                                                                                            fmt.Println(v3, v4)
                                                                                                            v118 := v1[1]
                                                                                                            var v119 int64 = 852
                                                                                                            _ = v118
                                                                                                            _ = v119
                                                                                                        } else {
                                                                                                            v120 := GenS6{F0: 32767, F1: (v3 / ((v3) | 1))}
                                                                                                            if (0 != -128) {
                                                                                                                v121 := GenS14{F0: 2, F1: 0}
                                                                                                                v122 := v1[0]
                                                                                                                v4 = (v52 | (v4 | (v4 / ((2) | 1))))
                                                                                                                _ = v121
                                                                                                                _ = v122
                                                                                                            } else {
                                                                                                                if (2147483647 <= 2147483648) {
                                                                                                                    if (176 <= 402) {
                                                                                                                        fmt.Println(fi83, v52)
                                                                                                                    }
                                                                                                                    v123 := []int{fi83}
                                                                                                                    if (true || (3.14159 <= 3.14159)) {
                                                                                                                        var v124 uint8 = 1
                                                                                                                        _ = v124
                                                                                                                    }
                                                                                                                    v125 := v1[42]
                                                                                                                    var v126 uint16 = 2
                                                                                                                    _ = v123
                                                                                                                    _ = v125
                                                                                                                    _ = v126
                                                                                                                } else {
                                                                                                                    var v127 uint16 = 42
                                                                                                                    v128 := GenS11{F0: 100, F1: v127, F2: int8(v52)}
                                                                                                                    v3 -= (v3 | (^(v3)))
                                                                                                                    _ = v127
                                                                                                                    _ = v128
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v120
                                                                                                        }
                                                                                                    } else {
                                                                                                        if (false && (true && true)) {
                                                                                                            var v129 uint16 = uint16(v4)
                                                                                                            var v130 uint64 = v4
                                                                                                            fmt.Println(v3)
                                                                                                            _ = v129
                                                                                                            _ = v130
                                                                                                        } else {
                                                                                                            switch (("hello" + "") + "a") {
                                                                                                            case "hello":
                                                                                                                switch "a" {
                                                                                                                case "":
                                                                                                                    switch fi85 {
                                                                                                                    case 236:
                                                                                                                        if false {
                                                                                                                            v2 += v2
                                                                                                                            var v131 bool = (1.5 == 0.1)
                                                                                                                            v132 := GenS14{F0: 2, F1: 2147483648}
                                                                                                                            _ = v131
                                                                                                                            _ = v132
                                                                                                                        } else {
                                                                                                                            v133 := []int{fi83, (-(fi83))}
                                                                                                                            _ = v133
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v134 := []int{(fi85 - fi85)}
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            switch "\n\t" {
                                                                                                            case "":
                                                                                                                v135 := GenS7{F0: 255, F1: 696, F2: fi83}
                                                                                                                _ = v135
                                                                                                            case "hello":
                                                                                                                v136 := []int{(fi83 | (fi83 | 2147483647)), 658, fi83, fi83}
                                                                                                                _ = v136
                                                                                                            }
                                                                                                            if (0.1 >= 0.0) {
                                                                                                                if true {
                                                                                                                    fmt.Println(v3, v87, v4)
                                                                                                                    fmt.Println(v4, v52)
                                                                                                                    var v137 int = -1
                                                                                                                    _ = v137
                                                                                                                } else {
                                                                                                                    v3 -= (v3 << (((v3 / ((uint(v52)) | 1))) & 63))
                                                                                                                    if true {
                                                                                                                        v2 = v2
                                                                                                                        v138 := GenS8{F0: (true || false)}
                                                                                                                        v139 := GenS11{F0: int8(v52), F1: 729, F2: 1}
                                                                                                                        var v140 uint = v3
                                                                                                                        v3 -= (v140 / ((v140) | 1))
                                                                                                                        _ = v138
                                                                                                                        _ = v139
                                                                                                                        _ = v140
                                                                                                                    }
                                                                                                                }
                                                                                                                v3 -= 255
                                                                                                                v141 := GenS15{F0: (false && true)}
                                                                                                                _ = v141
                                                                                                            }
                                                                                                            v3 -= 392
                                                                                                        }
                                                                                                        if ((false || false) && (242 == 103)) {
                                                                                                            fmt.Println(v4, v52)
                                                                                                            fmt.Println(v3, v2, v4)
                                                                                                        }
                                                                                                        switch fi85 {
                                                                                                        case 199:
                                                                                                            v87 = v87
                                                                                                        default:
                                                                                                            var v142 float32 = (3.14159 * 0.1)
                                                                                                            _ = v142
                                                                                                        }
                                                                                                    }
                                                                                                    switch -1 {
                                                                                                    case 1:
                                                                                                        v143 := GenS9{F0: (v3 << ((v3) & 63)), F1: 1}
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    if true {
                                                                                                        v144 := GenS8{F0: (fi83 <= (fi85 - fi83))}
                                                                                                        v52 = 2
                                                                                                        _ = v144
                                                                                                    } else {
                                                                                                        var v145 int32 = -26
                                                                                                        var v146 int = (^(fi83))
                                                                                                        var v147 uint32 = 1
                                                                                                        v148 := GenS16{F0: 127, F1: -1}
                                                                                                        v2 += (v2 / ((v145) | 1))
                                                                                                        if (9223372036854775807 != 9223372036854775807) {
                                                                                                            var v149 uint8 = 127
                                                                                                            v150 := GenS8{F0: true}
                                                                                                            _ = v149
                                                                                                            _ = v150
                                                                                                        }
                                                                                                        _ = v145
                                                                                                        _ = v146
                                                                                                        _ = v147
                                                                                                        _ = v148
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v151 string = (("" + "\n\t") + ((("\n\t" + "") + (("foo bar" + "foo bar") + ("\n\t" + ((("\n\t" + "\n\t") + "\n\t") + "foo bar")))) + "a"))
                                                                                                v152 := []int{}
                                                                                                _ = v151
                                                                                                _ = v152
                                                                                            }
                                                                                            var v153 int64 = 0
                                                                                            v154 := GenS17{F0: 18446744073709551615}
                                                                                            _ = v153
                                                                                            _ = v154
                                                                                        }
                                                                                    case "":
                                                                                        v155 := GenS18{F0: 100}
                                                                                        _ = v155
                                                                                    }
                                                                                    var v156 int32 = v2
                                                                                    v157 := GenS13{F0: (v3 > 255), F1: (("a" + "hello") + "a"), F2: -404}
                                                                                    _ = v156
                                                                                    _ = v157
                                                                                }
                                                                            } else {
                                                                                if ("foo bar" > (("foo bar" + "a") + ("\n\t" + (("hello" + "") + "")))) {
                                                                                    switch (fi85 ^ (fi85 ^ (fi83 - fi83))) {
                                                                                    case 9223372036854775807:
                                                                                        var v158 int64 = 0
                                                                                        fmt.Println(fi85, v2)
                                                                                        _ = v158
                                                                                    case 598:
                                                                                        v159 := v1[2]
                                                                                        _ = v159
                                                                                    case 252:
                                                                                        v87 -= (v87 / (((v87 << (((^(v3))) & 63))) | 1))
                                                                                    default:
                                                                                        if ((0 != 2147483647) || true) {
                                                                                            v87 = (-(v87))
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch fi85 {
                                                                                    case 99:
                                                                                        switch (fi85 - (fi83 << ((v3) & 63))) {
                                                                                        case -2147483648:
                                                                                            if false {
                                                                                                v160 := GenS18{F0: 100}
                                                                                                var v161 uint16 = 32767
                                                                                                _ = v160
                                                                                                _ = v161
                                                                                            }
                                                                                        case 128:
                                                                                            _ = len(v1)
                                                                                        case 0:
                                                                                            fmt.Println(v52, fi85)
                                                                                        default:
                                                                                            v87 = 32767
                                                                                        }
                                                                                    }
                                                                                    var v162 int16 = 1
                                                                                    var v163 int16 = (v162 / (((v87 & (^(v162)))) | 1))
                                                                                    v3 -= v3
                                                                                    v162 = 1000
                                                                                    v163 -= 363
                                                                                    switch (fi83 ^ (fi85 % (((fi85 % ((-9223372036854775808) | 1))) | 1))) {
                                                                                    case 127:
                                                                                        v164 := []int{fi85}
                                                                                        _ = v164
                                                                                    case 4294967295:
                                                                                        v165 := GenS16{F0: 42, F1: -80}
                                                                                        _ = v165
                                                                                    case 931:
                                                                                        delete(v1, 1)
                                                                                    default:
                                                                                        switch (^(fi85)) {
                                                                                        case -2147483648:
                                                                                            v166 := GenS19{F0: v4, F1: v52}
                                                                                            _ = v166
                                                                                        case -1:
                                                                                            var v167 uint32 = 1
                                                                                            _ = v167
                                                                                        case 42:
                                                                                            v2 = (^(v2))
                                                                                        default:
                                                                                            v168 := []int{(fi85 ^ fi83), fi85, fi85, -2147483648}
                                                                                            _ = v168
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        v3 += v3
                                                                                    }
                                                                                    _ = v162
                                                                                    _ = v163
                                                                                }
                                                                                v169 := []int{fi85, fi85, 32768}
                                                                                v170 := GenS20{F0: -87, F1: v52}
                                                                                _ = v169
                                                                                _ = v170
                                                                            }
                                                                            if ((489 > 65535) && true) {
                                                                                v87 = 1
                                                                                v171 := GenS21{F0: -815, F1: -128}
                                                                                _ = v171
                                                                            }
                                                                        } else {
                                                                            var v172 int64 = 9223372036854775807
                                                                            _ = v172
                                                                        }
                                                                        v2 = v2
                                                                        fmt.Println(v2, fi85)
                                                                        v173 := GenS22{F0: 42}
                                                                        fmt.Println(v2, fi85, v4)
                                                                        _ = v173
                                                                    }
                                                                    var v174 int64 = 160
                                                                    v175 := v1[2]
                                                                    _ = v174
                                                                    _ = v175
                                                                }
                                                            }
                                                        case -2147483648:
                                                            delete(v1, 0)
                                                        }
                                                        switch -365 {
                                                        case 127:
                                                            v176 := v1[2147483648]
                                                            _ = v176
                                                        case -2147483648:
                                                            v177 := []int{fi85, fi83, (fi85 & 0)}
                                                            _ = v177
                                                        default:
                                                            v178 := GenS23{F0: 701, F1: -128, F2: v3}
                                                            _ = v178
                                                        }
                                                        _ = fi85
                                                        _ = v86
                                                        _ = v87
                                                    }
                                                }
                                                _ = fi83
                                                _ = v84
                                            }
                                        }
                                        v179 := v1[42]
                                        var v180 string = v179
                                        _ = v179
                                        _ = v180
                                    }
                                    var v181 uint32 = 922
                                    v182 := GenS24{F0: 42, F1: (v181 / ((v181) | 1)), F2: v181}
                                    _ = v181
                                    _ = v182
                                } else {
                                    v1[uint32(v3)] = "\n\t"
                                }
                                for fi183 := 0; fi183 < 3; fi183++ {
                                    v52 = (v52 / ((v52) | 1))
                                    _ = fi183
                                }
                            }
                        } else {
                            var v184 uint64 = v52
                            var v185 uint16 = 42
                            switch ("foo bar" + ("\n\t" + "hello")) {
                            case "hello":
                                v52 -= 475
                            default:
                                var v186 int8 = -98
                                _ = v186
                            }
                            _ = v184
                            _ = v185
                        }
                    }
                    for fi187 := 0; fi187 < 2; fi187++ {
                        switch (fi187 << (((v3 & v3)) & 63)) {
                        case -368:
                            v188 := []int{}
                            _ = v188
                        case 127:
                            var v189 uint32 = 0
                            _ = v189
                        }
                        _ = fi187
                    }
                    fmt.Println(v2, v3, v3)
                } else {
                    var v190 string = "foo bar"
                    for fi191 := 0; fi191 < 16; fi191++ {
                        var v192 uint16 = 523
                        fmt.Println(v3)
                        _ = fi191
                        _ = v192
                    }
                    _ = v190
                }
                if false {
                    var v193 uint16 = 795
                    for fi194 := 0; fi194 < 3; fi194++ {
                        for fi195 := 0; fi195 < 17; fi195++ {
                            delete(v1, 2147483648)
                            v196 := GenS1{F0: 808}
                            _ = fi195
                            _ = v196
                        }
                        v197 := GenS7{F0: v193, F1: int64(v2), F2: (fi194 ^ (fi194 >> (((v3 / ((1) | 1))) & 63)))}
                        _ = fi194
                        _ = v197
                    }
                    v198 := []int{9223372036854775807, 32767, -269}
                    _ = v193
                    _ = v198
                } else {
                    v199 := GenS1{F0: 4294967295}
                    var v200 string = "hello"
                    _ = v199
                    _ = v200
                }
                _ = v67
            }
        default:
            v201 := GenS25{F0: 9223372036854775807, F1: 42}
            _ = v201
        }
    default:
        v202 := v1[276]
        _ = v202
    }
    fmt.Println(v2, v2, v3)
    fmt.Println(v2)
    var v203 uint64 = (v52 & v4)
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v203
    _ = v3
    _ = v4
    _ = v52
}
