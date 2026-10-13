package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16; F1 int8; F2 int64 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string; F1 int8; F2 int64 }
type GenS12 struct { F0 int16; F1 int32; F2 string }
type GenS13 struct { F0 uint16; F1 uint }
type GenS14 struct { F0 int8 }

func main() {
    v1 := []int{32768, 32768, 127, 0}
    var v2 uint32 = 2147483648
    v3 := []int{478, -2147483648, -9223372036854775808}
    v4 := make(map[string]int)
    v5 := []int{}
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 13; fi7++ {
        v2 = 42
        v6 += (v6 << ((1) & 63))
        _ = fi7
    }
    v2 = uint32(v6)
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: false}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        if ((v6 / ((127) | 1)) == 255) {
            v6 = v6
            v11 := GenS1{F0: (32767 == -1)}
            _ = v11
        } else {
            switch 391 {
            case 32767:
                if len(v3) > 0 {
                    v3[0] = fi10
                }
            case 1:
                v4["hello"] = 4294967295
            case 655:
                if ((!true) && false) {
                    var v12 int = (fi10 + fi10)
                    switch fi10 {
                    case -9223372036854775808:
                        v12 = 141
                    case 127:
                        switch 987 {
                        case 9223372036854775807:
                            v12 += (v12 * 42)
                        case 42:
                            v6 += (v6 ^ (v6 - (v6 | 71)))
                        case 128:
                            for fi13 := 0; fi13 < 18; fi13++ {
                                var v14 int8 = -97
                                v12 = (fi13 & v12)
                                v2 += (^(v2))
                                _ = fi13
                                _ = v14
                            }
                        default:
                            var v15 float64 = 0.1
                            _ = v15
                        }
                    case 2147483647:
                        if true {
                            v16 := GenS1{F0: (false && (42 < 65535))}
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 13; fi17++ {
                                v18 := GenS1{F0: (!true)}
                                if len(v3) > 0 {
                                    v3[0] = v12
                                }
                                if (1 == 0) {
                                    v19 := GenS1{F0: true}
                                    v12 -= v12
                                    v6 = v6
                                    _ = v19
                                } else {
                                    if (232 >= 794) {
                                        v20 := GenS1{F0: false}
                                        _ = v20
                                    }
                                }
                                _ = len(v4)
                                _ = fi17
                                _ = v18
                            }
                        }
                    default:
                        v21 := GenS2{F0: 2147483647}
                        _ = v21
                    }
                    _ = v12
                } else {
                    v22 := GenS3{F0: 18446744073709551615, F1: v2}
                    v23 := GenS2{F0: -2147483648}
                    _ = v22
                    _ = v23
                }
            default:
                if (!false) {
                    var v24 float64 = (-2.25 / 1.0)
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        if (!false) {
                            _ = len(v5)
                            switch "" {
                            case "\n\t":
                                v2 -= v2
                            }
                        } else {
                            fmt.Println(v2)
                            v2 = (v2 / ((v2) | 1))
                            var v26 float32 = 0.1
                            _ = len(v4)
                            switch (fi10 | fi10) {
                            case 529:
                                if false {
                                    if false {
                                        var v27 uint = 42
                                        var v28 string = "foo bar"
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v29 := GenS3{F0: 1, F1: v2}
                                        if (!(!(292 > 18446744073709551615))) {
                                            v30 := GenS2{F0: 0}
                                            _ = v30
                                        } else {
                                            fmt.Println(v26)
                                        }
                                        switch "" {
                                        case "foo bar":
                                            switch ("a" + "a") {
                                            case "a":
                                                v26 = ((v26 / float32(3.14159)) * v26)
                                            default:
                                                var v31 uint64 = 4294967295
                                                _ = v31
                                            }
                                        }
                                        _ = v29
                                    }
                                    v26 -= v26
                                    v26 -= 3.14159
                                } else {
                                    v6 -= v6
                                    fmt.Println(fi10)
                                    v3 = append(v3, 450)
                                }
                                v32 := GenS4{F0: 100}
                                switch fi10 {
                                case 1:
                                    v2 += 2147483648
                                case 539:
                                    if false {
                                        var v33 float32 = v26
                                        fmt.Println(fi25, v2)
                                        _ = v33
                                    } else {
                                        if (1 >= 4294967295) {
                                            switch "hello" {
                                            case "a":
                                                v2 = (v2 % ((v2) | 1))
                                            case "hello":
                                                fmt.Println(v2, fi25)
                                            default:
                                                var v34 float64 = (((0.1 + 0.0) * 3.14159) - (3.14159 / (1.5 / 0.1)))
                                                _ = v34
                                            }
                                            v35 := GenS4{F0: -128}
                                            var v36 bool = (42 < 4294967295)
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                case 0:
                                    var v37 uint64 = 62
                                    _ = v37
                                }
                                _ = v32
                            default:
                                switch ("foo bar" + "a") {
                                case "":
                                    var v38 float32 = v26
                                    _ = v38
                                case "foo bar":
                                    v26 = (-2.25 * v26)
                                case "hello":
                                    var v39 uint64 = 4294967295
                                    _ = v39
                                }
                            }
                            _ = v26
                        }
                        switch (fi10 ^ 127) {
                        case 0:
                            if false {
                                if true {
                                    if false {
                                        v40 := GenS5{F0: -578, F1: 190, F2: 1}
                                        _ = v40
                                    }
                                } else {
                                    var v41 int64 = 9223372036854775807
                                    if true {
                                        _ = len(v1)
                                        if (42 > 1) {
                                            var v42 float64 = 3.14159
                                            if true {
                                                v43 := GenS6{F0: 104}
                                                v4[((("foo bar" + "hello") + "foo bar") + "foo bar")] = 94
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = v43
                                            }
                                            _ = v42
                                        } else {
                                            var v44 int32 = -2147483648
                                            _ = v44
                                        }
                                        _ = len(v5)
                                        var v45 float32 = 0.0
                                        _ = v45
                                    } else {
                                        v46 := GenS7{F0: 0, F1: 4, F2: v41}
                                        _ = v46
                                    }
                                    _ = v41
                                }
                                var v47 bool = (true && true)
                                switch ("hello" + "foo bar") {
                                case "\n\t":
                                    v2 += (^(v2))
                                case "foo bar":
                                    v6 -= v6
                                case "":
                                    var v48 int16 = -3
                                    _ = v48
                                default:
                                    v2 = (^(v2))
                                }
                                _ = v47
                            }
                        default:
                            v6 += 255
                        }
                        if (!false) {
                            switch (("foo bar" + "hello") + "a") {
                            case "hello":
                                v6 += 65
                            case "\n\t":
                                v4[("foo bar" + (("hello" + "\n\t") + "foo bar"))] = fi10
                            }
                            switch "\n\t" {
                            case "\n\t":
                                switch (fi10 - fi10) {
                                case 127:
                                    var v49 uint = 345
                                    _ = v49
                                case -194:
                                    fmt.Println(v2, v6, v2)
                                case 32768:
                                    var v50 int8 = int8(v2)
                                    _ = v50
                                default:
                                    var v51 float32 = 3.14159
                                    _ = v51
                                }
                            case "foo bar":
                                v52 := GenS8{F0: -1, F1: 1, F2: 1}
                                _ = v52
                            case "hello":
                                var v53 uint64 = 4294967295
                                _ = v53
                            default:
                                var v54 float32 = (1e10 * float32((float64((1.5 * 1e10)) / (1.5 / ((3.14159 / ((0.1 + float64(3.14159)) - 1e10)) * 1.5)))))
                                _ = v54
                            }
                            v6 += 69
                        } else {
                            var v55 float64 = 3.14159
                            if ("a" != ("hello" + "")) {
                                v56 := GenS4{F0: -47}
                                if (42 <= 2147483647) {
                                    switch -2147483648 {
                                    case 9223372036854775807:
                                        delete(v4, ((((("" + ("a" + "\n\t")) + "foo bar") + "foo bar") + "") + "\n\t"))
                                    case 1:
                                        v55 = float64((0.0 / 1e10))
                                    }
                                    v4[("a" + "foo bar")] = 4294967295
                                }
                                switch fi10 {
                                case -9223372036854775808:
                                    switch "" {
                                    case "":
                                        if (128 >= (fi25 << ((446) & 63))) {
                                            switch -108 {
                                            case 0:
                                                var v57 bool = true
                                                _ = v57
                                            case 9223372036854775807:
                                                var v58 float64 = 1e10
                                                switch "a" {
                                                case "hello":
                                                    delete(v4, "foo bar")
                                                case "\n\t":
                                                    v58 -= (float64((-2.25 - 1.5)) + v55)
                                                case "":
                                                    v55 = v58
                                                default:
                                                    var v59 uint32 = v2
                                                    _ = v59
                                                }
                                                v60 := GenS5{F0: 42, F1: 42, F2: (v6 / ((v6) | 1))}
                                                _ = v58
                                                _ = v60
                                            default:
                                                v4[(("hello" + "hello") + (("foo bar" + ("" + (("hello" + "hello") + "foo bar"))) + ("\n\t" + ("a" + ""))))] = (-(fi25))
                                            }
                                            if (1 < 613) {
                                                v4[(("foo bar" + "a") + ("hello" + ("\n\t" + "hello")))] = 348
                                                v6 = (^(v6))
                                                if (uint16(v6) > 42) {
                                                    v6 = (v6 - (^(v6)))
                                                    if false {
                                                        v2 = (v2 << ((uint(v2)) & 63))
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v61 := GenS8{F0: int64(fi25), F1: 790, F2: 0}
                                                        var v62 uint32 = v2
                                                        _ = len(v4)
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        v55 = float64(0.1)
                                                    }
                                                    v63 := GenS2{F0: 1}
                                                    _ = v63
                                                }
                                            }
                                            v2 += (v2 << ((708) & 63))
                                            var v64 int = (fi10 ^ (fi25 >> ((208) & 63)))
                                            if true {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = (^(fi25))
                                                }
                                                if false {
                                                    var v65 int32 = -2147483648
                                                    switch fi10 {
                                                    case -58:
                                                        if (v65 <= (v65 * v65)) {
                                                            v6 = 128
                                                            v2 -= (^(v2))
                                                        }
                                                        var v66 int = (fi25 & fi25)
                                                        v65 -= v65
                                                        _ = v66
                                                    case -643:
                                                        var v67 uint8 = (v6 * (v6 >> ((1) & 63)))
                                                        _ = v67
                                                    case 42:
                                                        v68 := GenS9{F0: 4294967295, F1: v6}
                                                        var v69 int8 = 127
                                                        fmt.Println(v2)
                                                        _ = v68
                                                        _ = v69
                                                    default:
                                                        v6 += v6
                                                    }
                                                    var v70 string = (("a" + (((("a" + (("\n\t" + "\n\t") + "a")) + "hello") + "foo bar") + "a")) + ("\n\t" + ("a" + "hello")))
                                                    _ = v65
                                                    _ = v70
                                                }
                                            } else {
                                                switch ((("a" + "\n\t") + "hello") + ("" + "a")) {
                                                case "":
                                                    var v71 uint64 = 357
                                                    _ = v71
                                                case "hello":
                                                    if true {
                                                        v6 -= (v6 % ((32) | 1))
                                                        fmt.Println(v64, fi10)
                                                        v55 -= 0.0
                                                    }
                                                default:
                                                    switch ("" + ("hello" + ("" + "hello"))) {
                                                    case "":
                                                        switch (^(fi10)) {
                                                        case -1:
                                                            var v72 uint8 = 249
                                                            _ = v72
                                                        case 128:
                                                            if true {
                                                                delete(v4, "hello")
                                                            } else {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            var v73 uint64 = 1
                                                            _ = v73
                                                        }
                                                    case "hello":
                                                        var v74 bool = true
                                                        _ = v74
                                                    case "foo bar":
                                                        fmt.Println(fi10)
                                                    default:
                                                        v4["a"] = v64
                                                    }
                                                }
                                            }
                                            _ = v64
                                        } else {
                                            var v75 int32 = 1
                                            _ = v75
                                        }
                                    }
                                case 32768:
                                    switch fi10 {
                                    case 42:
                                        var v76 int32 = int32(fi25)
                                        _ = v76
                                    case 943:
                                        var v77 int32 = 1
                                        _ = v77
                                    default:
                                        var v78 uint16 = 834
                                        _ = v78
                                    }
                                case -491:
                                    if true {
                                        var v79 float32 = float32(float64(1e10))
                                        if true {
                                            v2 = (^(v2))
                                            v80 := GenS10{F0: fi10}
                                            fmt.Println(fi25, v2)
                                            _ = v80
                                        }
                                        var v81 float64 = v55
                                        v2 = v2
                                        switch "\n\t" {
                                        case "":
                                            _ = len(v4)
                                        default:
                                            if (v79 >= ((float32((-2.25 - 1.5)) * v79) / v79)) {
                                                v79 = v79
                                                fmt.Println(v55, v81, v79)
                                                var v82 int32 = 42
                                                _ = v82
                                            } else {
                                                v81 = ((0.1 * v81) / 1.0)
                                            }
                                        }
                                        _ = v79
                                        _ = v81
                                    }
                                    v55 = 3.14159
                                    v2 = v2
                                    v2 += 2147483647
                                    var v83 float32 = (1.5 + ((3.14159 - (float32((1e10 + v55)) - float32(3.14159))) - 0.0))
                                    _ = v83
                                default:
                                    switch "\n\t" {
                                    case "hello":
                                        v84 := GenS9{F0: (v2 ^ (v2 + (v2 / ((v2) | 1)))), F1: 0}
                                        _ = v84
                                    }
                                }
                                _ = v56
                            }
                            var v85 int64 = 9223372036854775807
                            _ = v55
                            _ = v85
                        }
                        _ = fi25
                    }
                    for fi86 := 0; fi86 < 14; fi86++ {
                        var v87 float32 = (0.1 / 1.0)
                        var v88 int16 = -174
                        switch fi10 {
                        case -417:
                            var v89 uint64 = 42
                            _ = v89
                        default:
                            if false {
                                delete(v4, ("foo bar" + "\n\t"))
                            } else {
                                v90 := GenS7{F0: 65535, F1: -128, F2: 9223372036854775807}
                                v91 := GenS3{F0: 42, F1: v2}
                                var v92 uint8 = (^(v6))
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                        }
                        _ = fi86
                        _ = v87
                        _ = v88
                    }
                }
            }
            v6 = v6
        }
        for fi93 := 0; fi93 < 5; fi93++ {
            fmt.Println(fi93, v2, v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch (-(fi93)) {
            case 127:
                var v94 int8 = 0
                _ = v94
            case -773:
                v6 += v6
            case 128:
                if false {
                    v95 := GenS10{F0: (fi93 + fi10)}
                    _ = v95
                } else {
                    if (0.0 != (0.0 - (-2.25 - 3.14159))) {
                        v6 += (v6 & 5)
                        v6 = (^(v6))
                        v6 += (v6 % ((v6) | 1))
                    } else {
                        delete(v4, "hello")
                    }
                    var v96 float64 = 1.5
                    _ = v96
                }
            default:
                var v97 uint = 42
                _ = v97
            }
            if true {
                fmt.Println(fi10, v2)
                _ = len(v5)
                var v98 float64 = 3.14159
                _ = v98
            } else {
                var v99 float32 = 0.0
                if false {
                    v99 = v99
                    var v100 string = "foo bar"
                    _ = v100
                }
                _ = v99
            }
            if false {
                v101 := GenS5{F0: 325, F1: 1, F2: v6}
                if false {
                    switch fi93 {
                    case -933:
                        switch fi93 {
                        case 2147483647:
                            v102 := GenS2{F0: 2147483647}
                            _ = v102
                        case -615:
                            v103 := GenS1{F0: false}
                            _ = v103
                        case 227:
                            switch (fi93 >> ((255) & 63)) {
                            case 612:
                                var v104 int = 0
                                _ = v104
                            default:
                                delete(v4, ("foo bar" + ((("hello" + "a") + "a") + "hello")))
                            }
                        }
                    case 42:
                        var v105 int = fi10
                        _ = v105
                    default:
                        var v106 int8 = -128
                        _ = v106
                    }
                }
                v107 := GenS11{F0: ("hello" + "foo bar"), F1: 42, F2: -9223372036854775808}
                var v108 int = fi93
                var v109 int64 = -1
                _ = v101
                _ = v107
                _ = v108
                _ = v109
            } else {
                var v110 int = 1
                _ = v110
            }
            _ = fi93
        }
        fmt.Println(fi10, v2)
        _ = fi10
    }
    v6 = (v6 * 128)
    var v111 int = 0
    for fi112 := 0; fi112 < 16; fi112++ {
        if (841 <= 919) {
            for fi113 := 0; fi113 < 1; fi113++ {
                if (!((true || ((("hello" + ("foo bar" + "hello")) + "hello") > ("\n\t" + "\n\t"))) || (0 < 1))) {
                    var v114 int32 = int32(v6)
                    if (v114 != 768) {
                        if false {
                            var v115 bool = (false || false)
                            switch 0 {
                            case -643:
                                fmt.Println(v115, fi112, v111)
                            case -950:
                                v6 += (v6 + (v6 - (v6 << ((65535) & 63))))
                            }
                            _ = v115
                        }
                        var v116 bool = ((^(v2)) < 4294967295)
                        _ = v116
                    }
                    switch (v111 >> ((904) & 63)) {
                    case -2147483648:
                        v117 := GenS12{F0: -768, F1: (v114 ^ v114), F2: (("a" + "hello") + ("\n\t" + "foo bar"))}
                        _ = v117
                    case 128:
                        switch fi113 {
                        case 0:
                            v118 := GenS13{F0: 65535, F1: 1}
                            _ = v118
                        case -1:
                            var v119 int64 = 9223372036854775807
                            _ = v119
                        case -2147483648:
                            v6 = 189
                        }
                    default:
                        switch "a" {
                        case "a":
                            switch "a" {
                            case "hello":
                                if (-262 != 42) {
                                    v120 := GenS14{F0: -1}
                                    var v121 uint32 = (^(v2))
                                    _ = v120
                                    _ = v121
                                }
                                var v122 int64 = 1
                                _ = v122
                            }
                            if len(v3) > 0 {
                                v3[0] = (fi113 + (fi112 << ((uint(v2)) & 63)))
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v111 = v111
                            fmt.Println(v114, fi113, v111)
                        }
                    }
                    _ = v114
                } else {
                    var v123 uint16 = 2
                    _ = v123
                }
                switch fi112 {
                case -9223372036854775808:
                    v6 = 138
                }
                _ = fi113
            }
            fmt.Println(fi112, v6, v111)
        }
        _ = fi112
    }
    if (v2 == (^(v2))) {
        fmt.Println(v111, v6, v2)
    } else {
        var v124 int64 = 217
        var v125 int8 = 100
        _ = v124
        _ = v125
    }
    var v126 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v111:", v111)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v111
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
