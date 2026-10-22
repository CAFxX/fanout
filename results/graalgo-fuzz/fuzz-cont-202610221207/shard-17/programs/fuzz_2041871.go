package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint64; F2 uint }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int8; F1 int8 }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 string; F1 int8; F2 int16 }
type GenS11 struct { F0 int16; F1 uint32; F2 int }
type GenS12 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint64; F1 int64 }
type GenS16 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 2147483647
    v3 := make(map[string]int)
    var v4 int16 = 42
    var v5 int32 = (v2 | (v2 % (((v2 | (v2 + v2))) | 1)))
    v4 = v4
    fmt.Println(v5, v2)
    v6 := GenS1{F0: true}
    if (4294967295 < 18446744073709551615) {
        v2 -= v2
    }
    v4 = 1
    fmt.Println(v5, v5)
    if (true || false) {
        v7 := GenS1{F0: false}
        if true {
            for fi8 := 0; fi8 < 13; fi8++ {
                switch (fi8 - fi8) {
                case 128:
                    v9 := []int{}
                    _ = v9
                case 127:
                    v10 := []int{}
                    _ = v10
                }
                if false {
                    v11 := GenS1{F0: false}
                    if true {
                        v2 = int32(v4)
                        if ((v4 >> ((2) & 63)) != (^(v4))) {
                            switch (((("foo bar" + "hello") + ("\n\t" + "")) + ("a" + ("hello" + "a"))) + "hello") {
                            case "hello":
                                if true {
                                    var v12 int16 = 0
                                    v4 = 42
                                    fmt.Println(v4)
                                    _ = v12
                                } else {
                                    var v13 float32 = 0.1
                                    v13 = ((0.1 + ((((0.0 + (((float32((((-2.25 / 1e10) - 3.14159) * 1.5)) / (v13 + float32(((1e10 * 0.0) / 1.0)))) - ((-2.25 * (v13 / (0.0 - 3.14159))) / ((1e10 + 1.5) * -2.25))) - v13)) + ((v13 - ((v13 - (float32(1.5) - 0.1)) - 1.5)) * (float32(1e10) * v13))) + (v13 * 1e10)) * float32(1e10))) - v13)
                                    _ = v13
                                }
                            }
                            for fi14 := 0; fi14 < 5; fi14++ {
                                v15 := GenS1{F0: ((573 >= (v4 + (v4 + (v4 % ((v4) | 1))))) && true)}
                                _ = fi14
                                _ = v15
                            }
                            _ = len(v3)
                        } else {
                            if (-1 <= (-(fi8))) {
                                v5 += 1
                                v5 -= (v2 | (v5 % ((v5) | 1)))
                            } else {
                                for fi16 := 0; fi16 < 6; fi16++ {
                                    var v17 int64 = 0
                                    if true {
                                        var v18 float32 = 1.5
                                        v18 = v18
                                        v18 = v18
                                        v18 = v18
                                        if false {
                                            switch "hello" {
                                            case "a":
                                                v19 := GenS1{F0: true}
                                                _ = v19
                                            case "":
                                                if (!true) {
                                                    switch "a" {
                                                    case "":
                                                        var v20 uint64 = 4294967295
                                                        _ = v20
                                                    case "hello":
                                                        var v21 int = (-(fi16))
                                                        v21 += (fi16 << ((2) & 63))
                                                        _ = v21
                                                    case "foo bar":
                                                        switch ("hello" + "a") {
                                                        case "":
                                                            var v22 uint8 = 242
                                                            _ = v22
                                                        case "a":
                                                            fmt.Println(v17, v2, v18)
                                                        case "foo bar":
                                                            _ = len(v3)
                                                        }
                                                    }
                                                }
                                            case "\n\t":
                                                v23 := []int{fi16, (fi16 * (-(fi16)))}
                                                _ = v23
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                switch 1 {
                                                case 251:
                                                    v17 += v17
                                                    switch ("" + "") {
                                                    case "a":
                                                        var v24 int16 = v4
                                                        _ = v24
                                                    }
                                                    if (0 >= 1) {
                                                        if true {
                                                            if false {
                                                                v18 = (v18 - float32((1e10 / 1e10)))
                                                                v17 += v17
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    delete(v3, ((("\n\t" + "a") + "a") + ""))
                                                                case "":
                                                                    var v25 bool = false
                                                                    _ = v25
                                                                }
                                                            }
                                                        } else {
                                                            if (false || (v4 != 307)) {
                                                                v17 = (v17 - 9223372036854775807)
                                                                v2 -= 360
                                                                v5 = v5
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                fmt.Println(v2, v2)
                                                            case "foo bar":
                                                                var v26 uint64 = 0
                                                                _ = v26
                                                            case "hello":
                                                                fmt.Println(v18, v18)
                                                            default:
                                                                v27 := []int{(fi16 - (fi16 >> ((1) & 63))), fi16, 127}
                                                                _ = v27
                                                            }
                                                            var v28 int32 = (v5 % ((v2) | 1))
                                                            _ = v28
                                                        }
                                                        if false {
                                                            v18 -= (float32(0.0) * v18)
                                                        } else {
                                                            var v29 uint8 = 0
                                                            v30 := GenS2{F0: "foo bar"}
                                                            var v31 uint8 = v29
                                                            switch (fi8 - -86) {
                                                            case 32767:
                                                                v32 := GenS1{F0: true}
                                                                _ = v32
                                                            default:
                                                                if true {
                                                                    if true {
                                                                        switch 32767 {
                                                                        case 9223372036854775807:
                                                                            v5 = v2
                                                                        case -9223372036854775808:
                                                                            v29 -= v31
                                                                        default:
                                                                            if false {
                                                                                v3["\n\t"] = (fi8 >> ((42) & 63))
                                                                                if true {
                                                                                    var v33 float32 = (0.1 - 1.5)
                                                                                    fmt.Println(v31, v2)
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        delete(v1, (-(fi8)))
                                                                                    default:
                                                                                        v4 -= (^(v4))
                                                                                    }
                                                                                    _ = v33
                                                                                }
                                                                            } else {
                                                                                var v34 string = ("foo bar" + (("hello" + ((("foo bar" + "hello") + "foo bar") + "")) + (("a" + (((((((("a" + "") + "\n\t") + "\n\t") + "\n\t") + "a") + "foo bar") + ("foo bar" + ("a" + ""))) + ("" + "hello"))) + ("foo bar" + "foo bar"))))
                                                                                v1[127] = 4294967295
                                                                                _ = v34
                                                                            }
                                                                        }
                                                                    }
                                                                    v35 := GenS1{F0: true}
                                                                    var v36 int16 = v4
                                                                    v2 = v2
                                                                    fmt.Println(fi16, v2)
                                                                    var v37 float64 = (0.0 * (1e10 + 1.5))
                                                                    v38 := GenS2{F0: ((((("" + (("a" + ("" + ("foo bar" + "hello"))) + "a")) + "\n\t") + (("\n\t" + ("hello" + ("hello" + ((("" + "a") + "foo bar") + "")))) + "foo bar")) + ("hello" + "hello")) + "")}
                                                                    _ = v35
                                                                    _ = v36
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    if false {
                                                                        if true {
                                                                            if (int32(v31) < (v2 ^ v5)) {
                                                                                v2 -= (^(v5))
                                                                                v17 += v17
                                                                                fmt.Println(v5, v29, fi8)
                                                                            } else {
                                                                                var v39 int = fi8
                                                                                _ = v39
                                                                            }
                                                                            switch "hello" {
                                                                            case "a":
                                                                                switch (^(fi8)) {
                                                                                case 1:
                                                                                    v5 -= -1
                                                                                case 4294967295:
                                                                                    var v40 string = "a"
                                                                                    _ = v40
                                                                                case 53:
                                                                                    var v41 uint64 = 18446744073709551615
                                                                                    _ = v41
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(v31, v31)
                                                                            default:
                                                                                var v42 string = ("foo bar" + "hello")
                                                                                _ = v42
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v43 float32 = v18
                                                                        v43 = (1e10 / (v18 * -2.25))
                                                                        var v44 int = fi8
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                            }
                                                            var v45 float64 = 1.5
                                                            _ = v29
                                                            _ = v30
                                                            _ = v31
                                                            _ = v45
                                                        }
                                                    }
                                                    if (!true) {
                                                        var v46 uint64 = 682
                                                        if (!(v18 == v18)) {
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v47 float64 = 3.14159
                                                                v18 -= v18
                                                                switch (^(fi16)) {
                                                                case 128:
                                                                    var v48 string = "a"
                                                                    _ = v48
                                                                case -2147483648:
                                                                    if false {
                                                                        switch "a" {
                                                                        case "a":
                                                                            if true {
                                                                                var v49 int32 = (v2 & (v2 + v5))
                                                                                _ = v49
                                                                            }
                                                                        default:
                                                                            var v50 int64 = v17
                                                                            _ = v50
                                                                        }
                                                                        v5 = -1
                                                                        v47 += v47
                                                                    } else {
                                                                        if ((false && (!(false && (42 <= 255)))) || false) {
                                                                            var v51 uint16 = 0
                                                                            v46 += v46
                                                                            _ = v51
                                                                        } else {
                                                                            v52 := GenS3{F0: 888}
                                                                            _ = v52
                                                                        }
                                                                        var v53 bool = false
                                                                        _ = v53
                                                                    }
                                                                }
                                                                _ = v47
                                                            case "foo bar":
                                                                switch fi8 {
                                                                case 1:
                                                                    switch 516 {
                                                                    case 4294967295:
                                                                        switch fi16 {
                                                                        case 603:
                                                                            v2 -= (v2 - (v2 - v5))
                                                                        case 1:
                                                                            v17 += 9223372036854775807
                                                                        case 295:
                                                                            if (v46 != v46) {
                                                                                var v54 uint = 18446744073709551615
                                                                                var v55 int8 = 74
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        var v56 int16 = (-(v4))
                                                                        v57 := v3["a"]
                                                                        var v58 uint8 = 42
                                                                        _ = v56
                                                                        _ = v57
                                                                        _ = v58
                                                                    case 65:
                                                                        if false {
                                                                            v1[9223372036854775807] = 782
                                                                            v4 -= v4
                                                                            fmt.Println(v2, v17)
                                                                        } else {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                fmt.Println(v46)
                                                                            case "hello":
                                                                                v59 := GenS3{F0: 42}
                                                                                _ = v59
                                                                            case "":
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    var v60 uint8 = 247
                                                                                    var v61 string = "a"
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                default:
                                                                                    fmt.Println(v17, fi16)
                                                                                }
                                                                            default:
                                                                                v62 := GenS4{F0: ""}
                                                                                _ = v62
                                                                            }
                                                                            v63 := GenS3{F0: 0}
                                                                            _ = v63
                                                                        }
                                                                    case -2147483648:
                                                                        v64 := GenS5{F0: fi8, F1: v46, F2: 65535}
                                                                        _ = v64
                                                                    default:
                                                                        fmt.Println(v18)
                                                                    }
                                                                case 781:
                                                                    if false {
                                                                        var v65 int16 = (v4 << ((0) & 63))
                                                                        _ = v65
                                                                    } else {
                                                                        v17 = v17
                                                                        v66 := []int{fi8, fi8}
                                                                        if ("\n\t" > "hello") {
                                                                            var v67 int32 = -2147483648
                                                                            fmt.Println(v2, v67, fi8)
                                                                            _ = v67
                                                                        } else {
                                                                            v68 := GenS6{F0: 18446744073709551615, F1: 65535, F2: 4294967295}
                                                                            _ = v68
                                                                        }
                                                                        _ = v66
                                                                    }
                                                                case 2147483647:
                                                                    v4 -= (v4 ^ (v4 & v4))
                                                                }
                                                            default:
                                                                var v69 float64 = -2.25
                                                                _ = v69
                                                            }
                                                        }
                                                        _ = v46
                                                    } else {
                                                        v1[9223372036854775807] = 2147483648
                                                    }
                                                    if false {
                                                        var v70 int8 = 42
                                                        v71 := v1[fi8]
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                    if false {
                                                        _ = len(v3)
                                                    } else {
                                                        if (true || false) {
                                                            if (v4 == 42) {
                                                                v72 := GenS5{F0: (fi8 - 589), F1: 1, F2: 18446744073709551615}
                                                                _ = v72
                                                            }
                                                            v18 = v18
                                                            var v73 int8 = -122
                                                            _ = v73
                                                        }
                                                        fmt.Println(v18, v4, v2)
                                                        v4 += -32768
                                                    }
                                                case -792:
                                                    switch (fi16 << ((2) & 63)) {
                                                    case -1:
                                                        fmt.Println(v2, fi8)
                                                    case -9223372036854775808:
                                                        var v74 uint = 1
                                                        _ = v74
                                                    }
                                                }
                                            case "\n\t":
                                                v75 := []int{(fi8 / ((32768) | 1)), 0}
                                                _ = v75
                                            case "hello":
                                                v76 := GenS7{F0: 127, F1: 0}
                                                _ = v76
                                            }
                                            switch "a" {
                                            case "a":
                                                var v77 uint32 = 2147483647
                                                _ = v77
                                            case "foo bar":
                                                v4 -= v4
                                            case "hello":
                                                switch "a" {
                                                case "a":
                                                    v5 = -792
                                                case "":
                                                    fmt.Println(v18, v2, v2)
                                                default:
                                                    var v78 uint32 = 1
                                                    _ = v78
                                                }
                                            }
                                            v17 += (v17 << ((4294967295) & 63))
                                        }
                                        _ = v18
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                                if (100 != 0) {
                                    v79 := GenS2{F0: ("a" + ("a" + "foo bar"))}
                                    v80 := GenS8{F0: -1, F1: (v4 ^ 1000)}
                                    v3["foo bar"] = fi8
                                    _ = v79
                                    _ = v80
                                } else {
                                    for fi81 := 0; fi81 < 13; fi81++ {
                                        var v82 int32 = v5
                                        _ = fi81
                                        _ = v82
                                    }
                                }
                                switch (fi8 - fi8) {
                                case -2147483648:
                                    for fi83 := 0; fi83 < 19; fi83++ {
                                        if ((!(int8(v4) < -1)) && false) {
                                            if false {
                                                v84 := GenS6{F0: 222, F1: 478, F2: 42}
                                                if true {
                                                    if (false && false) {
                                                        v85 := GenS9{F0: true, F1: (fi83 - (fi83 + 32767))}
                                                        if true {
                                                            v86 := GenS10{F0: "a", F1: int8(v4), F2: v4}
                                                            v87 := GenS9{F0: true, F1: (fi83 >> ((4294967295) & 63))}
                                                            v5 -= 1
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            var v88 int8 = 23
                                                            v4 -= -1
                                                            v5 = (v2 % ((647) | 1))
                                                            _ = v88
                                                        }
                                                        _ = v85
                                                    }
                                                }
                                                _ = v84
                                            }
                                            var v89 uint64 = 18446744073709551615
                                            _ = v89
                                        } else {
                                            v2 = (v2 + -820)
                                            fmt.Println(v2)
                                        }
                                        _ = fi83
                                    }
                                }
                            }
                            for fi90 := 0; fi90 < 4; fi90++ {
                                v91 := GenS5{F0: fi8, F1: 378, F2: 4294967295}
                                if ((-107 <= 127) || ("hello" < ("hello" + "a"))) {
                                    fmt.Println(v4, v5)
                                    switch (fi8 - fi90) {
                                    case 2147483647:
                                        v92 := GenS11{F0: v4, F1: 2147483647, F2: fi8}
                                        _ = v92
                                    case 1:
                                        v4 = 32767
                                    case 127:
                                        v4 = (^(v4))
                                    default:
                                        v93 := GenS4{F0: "a"}
                                        _ = v93
                                    }
                                    v94 := GenS12{F0: 2, F1: 689, F2: v2}
                                    v95 := []int{(fi90 + fi8)}
                                    _ = v94
                                    _ = v95
                                }
                                var v96 bool = true
                                _ = fi90
                                _ = v91
                                _ = v96
                            }
                            var v97 int16 = (-(v4))
                            _ = v97
                        }
                    } else {
                        v1[fi8] = 544
                        v98 := GenS6{F0: 0, F1: 65535, F2: 2}
                        switch "a" {
                        case "":
                            v5 = 795
                        default:
                            v1[(fi8 % ((630) | 1))] = 4294967295
                        }
                        _ = v98
                    }
                    if (695 == 44) {
                        fmt.Println(fi8)
                        fmt.Println(fi8, v5, v4)
                    } else {
                        var v99 uint8 = 176
                        _ = v99
                    }
                    _ = v11
                }
                switch "foo bar" {
                case "hello":
                    switch "foo bar" {
                    case "":
                        for fi100 := 0; fi100 < 8; fi100++ {
                            v101 := GenS13{F0: -128}
                            fmt.Println(fi100, fi8)
                            _ = fi100
                            _ = v101
                        }
                    case "foo bar":
                        v5 += (^(v5))
                    case "hello":
                        switch ("a" + ("hello" + ("a" + "\n\t"))) {
                        case "a":
                            v2 = -168
                        case "hello":
                            v2 = (v5 | (v5 * (v5 % ((1) | 1))))
                        case "\n\t":
                            fmt.Println(v2)
                        default:
                            var v102 uint16 = 255
                            _ = v102
                        }
                    }
                case "a":
                    for fi103 := 0; fi103 < 8; fi103++ {
                        v4 = v4
                        v104 := GenS14{F0: 4294967295, F1: 42, F2: 2147483648}
                        var v105 int = fi103
                        _ = fi103
                        _ = v104
                        _ = v105
                    }
                default:
                    v5 = (v2 ^ (v2 | 696))
                }
                _ = fi8
            }
        }
        if false {
            v106 := []int{-1, 0, -9223372036854775808, 495}
            v107 := GenS13{F0: 39}
            _ = v106
            _ = v107
        }
        _ = v7
    } else {
        v108 := []int{-2147483648, 128}
        switch ("" + ("a" + "\n\t")) {
        case "\n\t":
            if false {
                var v109 int = 128
                if true {
                    v2 = (v5 << ((877) & 63))
                } else {
                    if true {
                        for fi110 := 0; fi110 < 2; fi110++ {
                            for fi111 := 0; fi111 < 20; fi111++ {
                                var v112 uint32 = 952
                                _ = fi111
                                _ = v112
                            }
                            _ = fi110
                        }
                        if (128 <= -1) {
                            var v113 float64 = ((0.0 / -2.25) - 0.1)
                            fmt.Println(v5, v113)
                            if false {
                                var v114 string = ("foo bar" + "\n\t")
                                _ = v114
                            }
                            _ = v113
                        }
                    }
                    v5 += 2147483647
                    if (605 == 2) {
                        v115 := GenS15{F0: 2, F1: 0}
                        _ = v115
                    }
                }
                var v116 int8 = -1
                _ = v109
                _ = v116
            }
        case "foo bar":
            if false {
                v2 = (v2 ^ v2)
                v5 = (^(v5))
            }
        }
        v4 -= v4
        _ = v108
    }
    var v117 int16 = v4
    v118 := []int{-313, 127, 924, 369}
    if len(v118) >= 2 {
        v118 = v118[1:len(v118)]
    }
    if false {
        v119 := GenS16{F0: 127}
        _ = v119
    } else {
        var v120 int16 = v117
        _ = v120
    }
    for fi121 := 0; fi121 < 4; fi121++ {
        v122 := GenS4{F0: "hello"}
        _ = fi121
        _ = v122
    }
    for fi123 := 0; fi123 < 10; fi123++ {
        var v124 int32 = 1
        delete(v3, "")
        _ = fi123
        _ = v124
    }
    for fi125 := 0; fi125 < 5; fi125++ {
        fmt.Println(v2, fi125, v2)
        v126 := v1[(^(fi125))]
        var v127 float64 = 3.14159
        _ = fi125
        _ = v126
        _ = v127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v117:", v117)
    fmt.Println("slicelen v118:", len(v118), v118)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v117
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
