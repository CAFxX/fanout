package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint; F1 int8; F2 int32 }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int; F1 int8 }
type GenS8 struct { F0 int; F1 int64; F2 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool; F1 uint32 }
type GenS11 struct { F0 uint16; F1 int32 }
type GenS12 struct { F0 uint32; F1 bool }
type GenS13 struct { F0 int }
type GenS14 struct { F0 bool; F1 bool; F2 string }
type GenS15 struct { F0 uint; F1 int32 }
type GenS16 struct { F0 int32; F1 uint16 }
type GenS17 struct { F0 uint }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    if true {
        p0 += p0
    }
    v7 := GenS1{F0: int32(p0), F1: -15, F2: 42}
    fmt.Println(p0)
    switch "foo bar" {
    case "":
        v8 := []int{4294967295, 32768, 509, -9223372036854775808}
        _ = v8
    case "\n\t":
        var v9 uint64 = 621
        _ = v9
    default:
        v10 := []int{127, -82, 32768}
        _ = v10
    }
    _ = p0
    _ = v7
    return 2
}
func genFn1() string {
    var v11 int32 = 563
    v11 += v11
    switch -809 {
    case 32768:
        var v12 int8 = 127
        _ = v12
    }
    _ = v11
    return (("foo bar" + ("" + ("" + "hello"))) + "\n\t")
}

func main() {
    var v1 float32 = (0.0 + ((float32(1e10) + (float32(float64((3.14159 * 0.1))) * (0.0 / (float32((1e10 - 1e10)) + float32(-2.25))))) - -2.25))
    var v2 uint8 = 127
    var v3 int64 = -9223372036854775808
    var v4 int32 = -2147483648
    v5 := make(map[string]string)
    var v6 int64 = (-(v3))
    if false {
        var v13 int = -1
        var v14 int32 = v4
        v13 = 73
        _ = v13
        _ = v14
    }
    v4 = (v4 / ((0) | 1))
    for fi15 := 0; fi15 < 12; fi15++ {
        if (!false) {
            if ((-(v4)) <= v4) {
                var v16 string = "foo bar"
                _ = v16
            } else {
                v17 := GenS1{F0: 1, F1: 0, F2: 0}
                var v18 int64 = v6
                _ = v17
                _ = v18
            }
            switch "\n\t" {
            case "foo bar":
                switch (fi15 * int(v2)) {
                case 42:
                    v19 := GenS1{F0: v4, F1: 0, F2: 42}
                    _ = v19
                case 128:
                    v4 += 587
                case -20:
                    v20 := GenS2{F0: "", F1: 388}
                    _ = v20
                default:
                    if (2 == 1) {
                        var v21 float64 = 3.14159
                        _ = v21
                    }
                }
            case "\n\t":
                if true {
                    if (true && (!(-115 < 0))) {
                        v2 = v2
                        v22 := GenS1{F0: (v4 & (v4 % (((v4 / ((v4) | 1))) | 1))), F1: -1, F2: -1}
                        _ = v22
                    } else {
                        v23 := GenS3{F0: (1 >= 1), F1: 2147483647}
                        fmt.Println(v1, v4)
                        var v24 uint8 = 255
                        _ = v23
                        _ = v24
                    }
                } else {
                    var v25 int32 = 904
                    fmt.Println(v1, v4, v3)
                    _ = v25
                }
            case "":
                if false {
                    fmt.Println(v6)
                } else {
                    v2 += v2
                    if (genFn0(1) < genFn0(724)) {
                        for fi26 := 0; fi26 < 2; fi26++ {
                            var v27 int16 = -458
                            var v28 float32 = v1
                            v29 := GenS1{F0: (v4 / (((v4 & v4)) | 1)), F1: 100, F2: v27}
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 2; fi30++ {
                            var v31 uint32 = 1
                            var v32 int64 = v6
                            var v33 uint32 = (v31 % (((v31 + v31)) | 1))
                            _ = fi30
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                    } else {
                        v34 := GenS2{F0: "a", F1: 0}
                        v35 := GenS1{F0: v4, F1: 42, F2: 1}
                        _ = v34
                        _ = v35
                    }
                    if false {
                        for fi36 := 0; fi36 < 5; fi36++ {
                            v37 := GenS4{F0: 900, F1: 112, F2: (v4 - -2147483648)}
                            var v38 int64 = v3
                            fmt.Println(v6)
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        v39 := GenS5{F0: 18446744073709551615, F1: "foo bar"}
                        _ = v39
                    } else {
                        v1 = ((v1 + (float32((float64(v1) + (1e10 - (-2.25 + 0.0)))) / (1.5 + 0.1))) / (float32(1e10) / (v1 / v1)))
                    }
                }
            default:
                fmt.Println(fi15, fi15, v2)
            }
        } else {
            if true {
                var v40 int8 = -123
                v2 += 0
                for fi41 := 0; fi41 < 15; fi41++ {
                    var v42 int = fi41
                    v2 -= v2
                    _ = fi41
                    _ = v42
                }
                _ = v40
            }
        }
        fmt.Println(v1)
        if (true && (0.0 >= ((-2.25 * -2.25) * (-2.25 / (float64(1.5) + ((1e10 - (float64((float32((0.0 * (3.14159 - 0.0))) + 0.1)) * 0.1)) - float64(0.1))))))) {
            switch fi15 {
            case 9223372036854775807:
                v43 := v5[(("\n\t" + ("hello" + (("\n\t" + "") + ("hello" + "a")))) + ("" + "a"))]
                _ = v43
            case 2147483647:
                fmt.Println(v3)
            default:
                v44 := GenS6{F0: -24}
                _ = v44
            }
        }
        _ = fi15
    }
    if ((v4 + 42) <= v4) {
        if (true || (2147483647 < 32767)) {
            v1 -= float32(0.0)
            for fi45 := 0; fi45 < 2; fi45++ {
                v46 := []int{(fi45 >> ((2) & 63))}
                _ = fi45
                _ = v46
            }
        } else {
            v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
        }
        v47 := GenS2{F0: ("\n\t" + "\n\t"), F1: 792}
        _ = v47
    } else {
        fmt.Println(v4, v2, v1)
        v1 = 1e10
    }
    v6 -= 42
    v3 = v6
    switch "foo bar" {
    case "a":
        if (v6 >= (v3 - 42)) {
            v6 -= -9223372036854775808
            var v48 int64 = 9223372036854775807
            fmt.Println(v2, v2)
            _ = v48
        } else {
            delete(v5, "hello")
            switch 32767 {
            case 75:
                if true {
                    v49 := GenS7{F0: -1, F1: -128}
                    for fi50 := 0; fi50 < 17; fi50++ {
                        var v51 int16 = 1000
                        _ = fi50
                        _ = v51
                    }
                    _ = v49
                } else {
                    var v52 int32 = (v4 - v4)
                    if (!(v1 <= v1)) {
                        fmt.Println(v3, v1, v3)
                        var v53 float64 = (((((0.0 * 0.0) * (((3.14159 + 0.0) + (((float64(float32(1.5)) * ((-2.25 * 3.14159) / 3.14159)) + (0.1 + (((((((1.5 * 1e10) + -2.25) / -2.25) - 1e10) * (0.0 / 3.14159)) - 3.14159) * 0.1))) * (((((1.5 * (3.14159 - (float64(v1) / (float64(v1) / 3.14159)))) / 1.0) * 1e10) + ((((float64((0.1 - v1)) / ((-2.25 - 1.5) * (1e10 * ((1e10 / 1.0) / 0.1)))) * 0.0) * float64((v1 + -2.25))) * (1e10 + 1e10))) - (1.5 * ((0.1 + (0.1 * float64(0.1))) / (((1.5 * float64(v1)) - -2.25) / 1.5)))))) * 0.0)) - 0.0) - 1.5) * 0.0)
                        fmt.Println(v1, v6, v52)
                        _ = v53
                    }
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        fmt.Println(v6)
                    case "foo bar":
                        if (-33 >= 42) {
                            for fi54 := 0; fi54 < 14; fi54++ {
                                switch fi54 {
                                case -1:
                                    v55 := GenS5{F0: 0, F1: "foo bar"}
                                    _ = v55
                                case 9223372036854775807:
                                    if false {
                                        switch (fi54 >> ((897) & 63)) {
                                        case 78:
                                            v52 = v52
                                        case -55:
                                            if false {
                                                if (10 >= 127) {
                                                    for fi56 := 0; fi56 < 7; fi56++ {
                                                        v1 += float32(-2.25)
                                                        switch (("\n\t" + "\n\t") + "\n\t") {
                                                        case "hello":
                                                            if ((float64((3.14159 + (v1 * v1))) + (1.5 / float64(v1))) > 0.0) {
                                                                if true {
                                                                    fmt.Println(v4, v52)
                                                                } else {
                                                                    if (((v3 / (((v3 % ((-1) | 1))) | 1)) >= (v3 >> ((65535) & 63))) || (-128 <= 100)) {
                                                                        var v57 string = ("a" + "foo bar")
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 int16 = 1
                                                                        var v59 int8 = 0
                                                                        if false {
                                                                            v60 := GenS8{F0: (fi54 * fi54), F1: int64(v2), F2: v2}
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v61 float32 = -2.25
                                                                                    _ = v61
                                                                                } else {
                                                                                    if false {
                                                                                        var v62 uint8 = (v2 & genFn0(0))
                                                                                        fmt.Println(v2)
                                                                                        _ = v62
                                                                                    }
                                                                                    v3 -= 9223372036854775807
                                                                                    v63 := GenS9{F0: 255}
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                            fmt.Println(v52, v6, v1)
                                                                            _ = v60
                                                                        } else {
                                                                            v4 = v4
                                                                            switch -9223372036854775808 {
                                                                            case 0:
                                                                                v1 -= v1
                                                                            case -115:
                                                                                v4 = (-(v52))
                                                                            case 1:
                                                                                v4 = (v4 * (v4 + v4))
                                                                            }
                                                                            var v64 string = "hello"
                                                                            v65 := GenS10{F0: (v64 != (((v64 + v64) + "a") + v64)), F1: 2147483647}
                                                                            v58 = v58
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                }
                                                                v3 = -9223372036854775808
                                                            }
                                                        case "foo bar":
                                                            var v66 uint8 = v2
                                                            _ = v66
                                                        }
                                                        _ = fi56
                                                    }
                                                    v6 = -9223372036854775808
                                                } else {
                                                    var v67 int8 = -75
                                                    switch "" {
                                                    case "hello":
                                                        v68 := v5["a"]
                                                        _ = v68
                                                    case "\n\t":
                                                        switch fi54 {
                                                        case 647:
                                                            v69 := []int{}
                                                            _ = v69
                                                        default:
                                                            switch "" {
                                                            case "hello":
                                                                var v70 int16 = 32767
                                                                _ = v70
                                                            case "foo bar":
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v6 = 9223372036854775807
                                                                default:
                                                                    for fi71 := 0; fi71 < 9; fi71++ {
                                                                        var v72 int32 = v4
                                                                        var v73 uint8 = 42
                                                                        _ = fi71
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v67
                                                }
                                                v74 := GenS3{F0: ((v2 << ((290) & 63)) != (^(v2))), F1: 870}
                                                _ = v74
                                            }
                                        default:
                                            for fi75 := 0; fi75 < 10; fi75++ {
                                                var v76 int64 = 200
                                                if ((false && false) && true) {
                                                    v3 = v76
                                                    if false {
                                                        v77 := GenS11{F0: 1, F1: v4}
                                                        v78 := []int{fi75, 42, (^(fi75)), 13}
                                                        switch "hello" {
                                                        case "a":
                                                            var v79 int8 = -128
                                                            _ = v79
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    } else {
                                                        var v80 int32 = -2147483648
                                                        _ = v80
                                                    }
                                                    var v81 string = ""
                                                    _ = v81
                                                }
                                                v82 := v5["\n\t"]
                                                _ = fi75
                                                _ = v76
                                                _ = v82
                                            }
                                        }
                                        if true {
                                            v83 := []int{-1}
                                            _ = v83
                                        }
                                    } else {
                                        for fi84 := 0; fi84 < 3; fi84++ {
                                            v4 = (-(v4))
                                            _ = fi84
                                        }
                                        if false {
                                            if (false || true) {
                                                var v85 uint = 255
                                                for fi86 := 0; fi86 < 1; fi86++ {
                                                    var v87 uint8 = genFn0(2)
                                                    if true {
                                                        v6 -= v6
                                                        var v88 int64 = 42
                                                        var v89 bool = (!(2147483647 == 0))
                                                        v90 := GenS12{F0: 0, F1: (!v89)}
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    } else {
                                                        v5[("hello" + "\n\t")] = ("hello" + (("\n\t" + ("" + (("a" + (("hello" + "\n\t") + "hello")) + ""))) + "foo bar"))
                                                        v91 := GenS7{F0: (fi54 >> ((v85) & 63)), F1: -11}
                                                        if true {
                                                            fmt.Println(v6, v87, v87)
                                                            v92 := GenS7{F0: 32767, F1: -124}
                                                            _ = v92
                                                        } else {
                                                            v93 := GenS5{F0: 2, F1: "\n\t"}
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if (true && ((v87 - v2) >= v87)) {
                                                                    v94 := GenS2{F0: "", F1: 42}
                                                                    v87 = v2
                                                                    _ = v94
                                                                }
                                                            case "":
                                                                var v95 uint64 = 0
                                                                _ = v95
                                                            }
                                                            if ("foo bar" <= "hello") {
                                                                delete(v5, "")
                                                                var v96 float32 = 3.14159
                                                                _ = v96
                                                            }
                                                            _ = v93
                                                        }
                                                        _ = v91
                                                    }
                                                    if false {
                                                        v97 := GenS13{F0: (^(fi86))}
                                                        _ = v97
                                                    } else {
                                                        if (1 == -13) {
                                                            v98 := GenS8{F0: (fi86 & 32767), F1: int64(v85), F2: v2}
                                                            switch fi86 {
                                                            case 128:
                                                                var v99 int = fi54
                                                                _ = v99
                                                            }
                                                            v85 += v85
                                                            v87 = (^(v87))
                                                            switch "a" {
                                                            case "foo bar":
                                                                v52 = (v4 + (v52 / (((v52 ^ (v4 * 104))) | 1)))
                                                            }
                                                            _ = v98
                                                        } else {
                                                            var v100 int8 = 1
                                                            _ = v100
                                                        }
                                                        v101 := GenS14{F0: true, F1: false, F2: ""}
                                                        if ((false && ((v1 - 1.5) < v1)) && (false && false)) {
                                                            _ = len(v5)
                                                            var v102 int = (fi86 - fi54)
                                                            var v103 float64 = (1e10 - ((0.0 / 1e10) - 1.5))
                                                            v104 := v5[""]
                                                            _ = v102
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        _ = v101
                                                    }
                                                    _ = fi86
                                                    _ = v87
                                                }
                                                _ = v85
                                            }
                                            v1 = v1
                                            var v105 int = fi54
                                            _ = v105
                                        }
                                    }
                                case 127:
                                    for fi106 := 0; fi106 < 15; fi106++ {
                                        var v107 uint = 4294967295
                                        switch (fi54 << ((42) & 63)) {
                                        case 0:
                                            if (-2147483648 <= -2147483648) {
                                                v5["foo bar"] = ("foo bar" + ("a" + "foo bar"))
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v1)
                                                case "hello":
                                                    v108 := GenS15{F0: 0, F1: -234}
                                                    _ = v108
                                                }
                                                if false {
                                                    v1 += ((float32(3.14159) - (0.0 / (((float32(1.5) - v1) / ((v1 - (((v1 - 1.5) - v1) + (v1 - (((0.0 / 3.14159) / float32(1.5)) / ((v1 + 3.14159) - v1))))) / 1e10)) * v1))) / -2.25)
                                                    if true {
                                                        switch "" {
                                                        case "hello":
                                                            var v109 uint64 = 1
                                                            _ = v109
                                                        }
                                                        v110 := v5[("" + ("\n\t" + (("hello" + (("foo bar" + "") + "\n\t")) + "")))]
                                                        var v111 uint32 = 28
                                                        _ = v110
                                                        _ = v111
                                                    }
                                                    if true {
                                                        v107 -= (v107 >> (((v107 & (^(v107)))) & 63))
                                                        if (false || (v3 >= 9223372036854775807)) {
                                                            if false {
                                                                var v112 float64 = (3.14159 + 0.0)
                                                                v1 = (v1 / v1)
                                                                _ = v112
                                                            }
                                                        }
                                                        var v113 uint32 = 2147483647
                                                        _ = v113
                                                    } else {
                                                        fmt.Println(fi106, v4)
                                                    }
                                                }
                                                v4 -= (-(v52))
                                            }
                                        case 9223372036854775807:
                                            delete(v5, "")
                                        case -782:
                                            v6 = 950
                                        default:
                                            var v114 uint32 = 0
                                            _ = v114
                                        }
                                        v3 = v6
                                        var v115 string = "foo bar"
                                        v1 = v1
                                        _ = fi106
                                        _ = v107
                                        _ = v115
                                    }
                                default:
                                    v5["hello"] = (("\n\t" + "hello") + "")
                                }
                                _ = fi54
                            }
                        }
                    }
                    _ = v52
                }
            case 1:
                v116 := []int{-710, 32768}
                _ = v116
            default:
                v117 := GenS6{F0: -128}
                _ = v117
            }
        }
    case "foo bar":
        v118 := GenS16{F0: v4, F1: 2}
        _ = v118
    }
    v4 -= 0
    delete(v5, "")
    v119 := GenS13{F0: 127}
    var v120 string = "\n\t"
    var v121 int16 = 78
    var v122 int = -9223372036854775808
    v4 += (v4 | 1000)
    if false {
        if (0 >= -1) {
            if false {
                for fi123 := 0; fi123 < 9; fi123++ {
                    for fi124 := 0; fi124 < 14; fi124++ {
                        if false {
                            fmt.Println(v4, v3)
                            v121 += v121
                        } else {
                            v125 := []int{32767, (fi123 >> ((531) & 63))}
                            v1 -= v1
                            v121 += int16(fi123)
                            v121 = (v121 % ((int16(v3)) | 1))
                            _ = v125
                        }
                        switch v120 {
                        case "":
                            fmt.Println(v1)
                        case "a":
                            if true {
                                fmt.Println(v122)
                                v126 := GenS16{F0: -2147483648, F1: 178}
                                var v127 int = fi124
                                _ = v126
                                _ = v127
                            }
                        case "hello":
                            var v128 uint32 = 0
                            _ = v128
                        default:
                            var v129 int64 = -645
                            fmt.Println(v129)
                            switch 128 {
                            case 32768:
                                var v130 float32 = (v1 * (v1 * v1))
                                _ = v130
                            case 2147483647:
                                switch 0 {
                                case 9223372036854775807:
                                    fmt.Println(v122, v122)
                                case 128:
                                    v122 = 397
                                case 0:
                                    fmt.Println(v2)
                                }
                            default:
                                v121 += (v121 ^ v121)
                            }
                            _ = v129
                        }
                        _ = fi124
                    }
                    if false {
                        v131 := GenS17{F0: 0}
                        delete(v5, v120)
                        _ = v131
                    } else {
                        var v132 bool = true
                        var v133 bool = false
                        v122 += v122
                        _ = v132
                        _ = v133
                    }
                    var v134 uint64 = 650
                    _ = fi123
                    _ = v134
                }
                switch 0 {
                case 0:
                    delete(v5, ((v120 + "a") + "foo bar"))
                }
                var v135 int8 = 1
                _ = v135
            } else {
                switch (v122 ^ (v122 ^ v122)) {
                case 42:
                    var v136 uint32 = 42
                    _ = v136
                case -2147483648:
                    v137 := GenS16{F0: (v4 >> ((2) & 63)), F1: 2}
                    _ = v137
                default:
                    v121 = 0
                }
            }
            v121 += v121
        } else {
            for fi138 := 0; fi138 < 1; fi138++ {
                v122 += fi138
                var v139 string = (v120 + (((("foo bar" + "") + "") + (v120 + v120)) + (((v120 + v120) + "hello") + v120)))
                v6 = (v6 ^ v3)
                _ = fi138
                _ = v139
            }
            v121 -= v121
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v119
    _ = v120
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
