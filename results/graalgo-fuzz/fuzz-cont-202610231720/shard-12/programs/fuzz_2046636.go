package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 string; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 int16; F1 int16; F2 bool }
func genFn0(p0 int32) bool {
    var v7 int = 0
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 uint16 = 401
        fmt.Println(p0, v7, p0)
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 42
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{9223372036854775807, (v7 ^ (fi12 - (-(fi12)))), (fi10 % (((-(fi10))) | 1)), 503}
            _ = fi12
            _ = v13
        }
        if (false && true) {
            if (!false) {
                if true {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if false {
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                v15 := []int{v7, v7}
                                _ = v15
                            default:
                                v16 := []int{}
                                _ = v16
                            }
                        }
                        var v17 bool = true
                        var v18 int8 = 37
                        fmt.Println(v18)
                        v19 := GenS1{F0: 127}
                        _ = fi14
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                if true {
                    var v20 string = ("\n\t" + "foo bar")
                    var v21 float64 = float64((1e10 + 0.0))
                    _ = v20
                    _ = v21
                }
                var v22 int8 = 114
                _ = v22
            } else {
                if (4294967295 > 99) {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        if (589 <= 65535) {
                            v24 := GenS1{F0: 127}
                            _ = v24
                        } else {
                            v25 := []int{(fi23 * 2147483647), 42}
                            _ = v25
                        }
                        var v26 int64 = -98
                        v27 := GenS2{F0: 1}
                        _ = fi23
                        _ = v26
                        _ = v27
                    }
                } else {
                    if false {
                        v11 += (^(v11))
                        var v28 uint64 = 576
                        if false {
                            v29 := []int{(v7 >> ((390) & 63)), (^(fi10))}
                            fmt.Println(v7, v7, p0)
                            if false {
                                for fi30 := 0; fi30 < 17; fi30++ {
                                    var v31 uint16 = (v11 << ((42) & 63))
                                    var v32 int = 32768
                                    if ("hello" != (("foo bar" + ("foo bar" + "a")) + "a")) {
                                        var v33 bool = true
                                        var v34 int = 127
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v11 = (v31 & 32767)
                                    }
                                    var v35 uint32 = 4294967295
                                    var v36 uint32 = (v35 * (v35 << ((4294967295) & 63)))
                                    var v37 float32 = 1e10
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                v11 += 2
                                switch fi10 {
                                case 2147483647:
                                    var v38 int64 = 9223372036854775807
                                    _ = v38
                                case 4294967295:
                                    v39 := GenS3{F0: 84}
                                    _ = v39
                                }
                            }
                            _ = v29
                        } else {
                            v40 := GenS4{F0: 2}
                            if ("hello" != "\n\t") {
                                v41 := []int{2147483647, v7}
                                v7 -= (v7 * v7)
                                v11 = (v11 & (^(v11)))
                                _ = v41
                            }
                            _ = v40
                        }
                        _ = v28
                    }
                    switch "" {
                    case "\n\t":
                        p0 = p0
                    case "a":
                        if (131 < v7) {
                            v42 := []int{0, fi10}
                            var v43 uint64 = 301
                            _ = v42
                            _ = v43
                        } else {
                            fmt.Println(v7, fi10)
                            switch v7 {
                            case 4294967295:
                                switch fi10 {
                                case 127:
                                    v44 := GenS5{F0: p0}
                                    _ = v44
                                case -430:
                                    switch "hello" {
                                    case "\n\t":
                                        switch (v7 / (((fi10 | int(v11))) | 1)) {
                                        case 576:
                                            p0 -= (^(p0))
                                        }
                                    case "foo bar":
                                        var v45 int8 = 0
                                        _ = v45
                                    case "a":
                                        v7 -= fi10
                                    }
                                }
                            case -9223372036854775808:
                                var v46 uint = 4294967295
                                _ = v46
                            case 920:
                                var v47 string = ("a" + "a")
                                _ = v47
                            default:
                                p0 = p0
                            }
                            v7 = (fi10 + (v7 >> ((42) & 63)))
                        }
                    case "hello":
                        v48 := GenS3{F0: 727}
                        _ = v48
                    }
                }
            }
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v7
    return true
}

func main() {
    v1 := []uint32{2147483648, 4294967295}
    v2 := make(map[string]string)
    var v3 int32 = 900
    var v4 uint = 2
    var v5 int64 = 9223372036854775807
    v6 := make(map[string]string)
    v49 := GenS5{F0: v3}
    if (1 < 42) {
        var v50 int = 32767
        _ = v50
    }
    for fi51 := 0; fi51 < 16; fi51++ {
        switch fi51 {
        case 32768:
            var v52 uint = 1
            _ = v52
        case -1:
            if false {
                v53 := GenS4{F0: uint32(v4)}
                if false {
                    switch 32768 {
                    case -1:
                        v5 = v5
                    case 2147483647:
                        for fi54 := 0; fi54 < 17; fi54++ {
                            delete(v6, "foo bar")
                            if true {
                                fmt.Println(fi51, v3, fi51)
                            }
                            _ = fi54
                        }
                    default:
                        var v55 uint = v4
                        _ = v55
                    }
                    _ = len(v1)
                }
                _ = v53
            } else {
                for fi56 := 0; fi56 < 9; fi56++ {
                    if len(v1) > 0 {
                        v1[0] = 989
                    }
                    switch 105 {
                    case 598:
                        var v57 int8 = int8(v4)
                        _ = v57
                    case 759:
                        if false {
                            switch (fi51 >> (((v4 >> (((v4 / ((v4) | 1))) & 63))) & 63)) {
                            case -226:
                                if true {
                                    var v58 uint8 = 1
                                    _ = v58
                                }
                            case -2147483648:
                                v59 := GenS4{F0: 1}
                                _ = v59
                            default:
                                v4 = (v4 | v4)
                            }
                            var v60 int32 = v3
                            _ = v60
                        } else {
                            if ((fi56 == (fi51 * (fi51 >> ((v4) & 63)))) && (!false)) {
                                var v61 bool = (255 >= 552)
                                if false {
                                    if (!true) {
                                        var v62 string = "foo bar"
                                        v3 = (v3 ^ (v3 * (v3 << ((v4) & 63))))
                                        _ = len(v1)
                                        _ = v62
                                    } else {
                                        var v63 int16 = -32768
                                        if (179 >= 922) {
                                            var v64 float64 = 0.0
                                            switch "\n\t" {
                                            case "a":
                                                switch (("foo bar" + ("hello" + (("hello" + "foo bar") + ((("a" + "hello") + ("a" + "")) + ("foo bar" + "a"))))) + (("a" + "a") + ("hello" + (("" + "\n\t") + (("" + (("" + "foo bar") + (("hello" + "a") + ("\n\t" + "hello")))) + "a"))))) {
                                                case "foo bar":
                                                    switch -1 {
                                                    case 42:
                                                        var v65 int = (^(fi56))
                                                        v3 = v3
                                                        fmt.Println(fi51, v63, fi51)
                                                        _ = v65
                                                    case 1:
                                                        switch ("\n\t" + "\n\t") {
                                                        case "hello":
                                                            v4 += v4
                                                        case "":
                                                            v4 = 255
                                                        default:
                                                            fmt.Println(v5)
                                                        }
                                                    case 32767:
                                                        var v66 int16 = (-(v63))
                                                        _ = v66
                                                    }
                                                case "\n\t":
                                                    if v61 {
                                                        _ = len(v1)
                                                        v2["hello"] = "foo bar"
                                                        var v67 uint = 0
                                                        _ = v67
                                                    } else {
                                                        var v68 int32 = 1
                                                        v2["foo bar"] = "hello"
                                                        v3 = 515
                                                        if (59 > 42) {
                                                            var v69 string = ""
                                                            _ = v69
                                                        }
                                                        v70 := GenS6{F0: (65535 > uint32(v4))}
                                                        var v71 uint32 = 65535
                                                        var v72 uint16 = 2
                                                        _ = v68
                                                        _ = v70
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                }
                                            default:
                                                var v73 int8 = 127
                                                _ = v73
                                            }
                                            if (v61 || v61) {
                                                var v74 float32 = (3.14159 * (-2.25 * (0.0 + 3.14159)))
                                                fmt.Println(v74, v4, fi51)
                                                _ = v74
                                            }
                                            var v75 uint = v4
                                            _ = v64
                                            _ = v75
                                        } else {
                                            var v76 float64 = (1e10 * 0.0)
                                            if ((("hello" + "hello") + "foo bar") <= (("a" + ("hello" + "hello")) + "\n\t")) {
                                                v77 := GenS7{F0: 4294967295}
                                                var v78 uint8 = 149
                                                var v79 uint64 = 2
                                                _ = v77
                                                _ = v78
                                                _ = v79
                                            }
                                            _ = v76
                                        }
                                        _ = v63
                                    }
                                    v61 = true
                                } else {
                                    var v80 string = "a"
                                    var v81 int64 = 9223372036854775807
                                    if false {
                                        v82 := GenS8{F0: 1}
                                        v61 = (42 < 42)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v82
                                    }
                                    v83 := GenS2{F0: 0}
                                    _ = v80
                                    _ = v81
                                    _ = v83
                                }
                                _ = v61
                            } else {
                                v3 = (v3 & v3)
                                var v84 int8 = 100
                                _ = v84
                            }
                            switch "\n\t" {
                            case "hello":
                                v1 = append(v1, 275)
                            }
                        }
                    }
                    _ = fi56
                }
                v85 := GenS9{F0: -1}
                _ = v85
            }
        case -930:
            var v86 uint16 = 42
            _ = v86
        }
        for fi87 := 0; fi87 < 8; fi87++ {
            _ = len(v1)
            if (false && (255 < 0)) {
                var v88 bool = (true && false)
                var v89 float32 = (1.5 * (float32(((0.1 * 1.5) + float64(3.14159))) * 3.14159))
                if v88 {
                    _ = len(v6)
                    v89 -= (v89 / 1e10)
                    v88 = v88
                    var v90 int32 = (v3 << (((v4 - (v4 - 274))) & 63))
                    _ = v90
                }
                _ = v88
                _ = v89
            } else {
                if true {
                    fmt.Println(v5)
                    v4 = v4
                    v4 = v4
                }
                v3 = v3
                v91 := v2[""]
                _ = v91
            }
            if true {
                if len(v1) > 0 {
                    v1[0] = 0
                }
                v5 += -588
                fmt.Println(fi87)
            }
            _ = fi87
        }
        switch ("" + ("\n\t" + "foo bar")) {
        case "foo bar":
            switch "foo bar" {
            case "hello":
                v4 = v4
            case "\n\t":
                var v92 int8 = -1
                _ = v92
            case "a":
                v5 -= -9223372036854775808
            }
        case "a":
            v93 := v6[("" + "foo bar")]
            _ = v93
        default:
            var v94 int64 = (^(v5))
            _ = v94
        }
        _ = fi51
    }
    delete(v2, "\n\t")
    fmt.Println(v4, v4)
    var v95 uint16 = 32767
    var v96 uint64 = 4294967295
    v5 += v5
    var v97 uint8 = 1
    switch ("\n\t" + "") {
    case "foo bar":
        v96 = 2
    case "":
        for fi98 := 0; fi98 < 19; fi98++ {
            var v99 int32 = 2147483647
            v95 = v95
            var v100 int16 = -303
            _ = fi98
            _ = v100
            _ = v99
        }
    }
    v101 := GenS3{F0: 2147483647}
    for fi102 := 0; fi102 < 6; fi102++ {
        fmt.Println(v97)
        v5 -= v5
        _ = fi102
    }
    v103 := GenS10{F0: 920, F1: "hello", F2: 1}
    if ((4294967295 >= 42) && (v3 != (v3 - v3))) {
        var v104 uint32 = 2147483648
        _ = v104
    } else {
        if (0.0 >= (1.5 - 3.14159)) {
            if true {
                v97 += v97
                var v105 int16 = 42
                _ = v105
            } else {
                v95 = v95
                v1 = append(v1, 42)
                v106 := GenS7{F0: v96}
                _ = v106
            }
            if len(v1) > 0 {
                v1[0] = 65535
            }
            fmt.Println(v95)
        } else {
            for fi107 := 0; fi107 < 17; fi107++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v108 := GenS11{F0: -224}
                _ = fi107
                _ = v108
            }
            switch ((((((("\n\t" + "a") + "\n\t") + (("" + "") + ("" + "a"))) + "\n\t") + "a") + "hello") + "foo bar") {
            case "":
                if (42 < v4) {
                    v109 := GenS7{F0: 2}
                    _ = v109
                }
            }
        }
        var v110 int = 32768
        if true {
            v96 = uint64(v95)
            var v111 uint16 = v95
            _ = v111
        }
        _ = v110
    }
    for fi112 := 0; fi112 < 12; fi112++ {
        for fi113 := 0; fi113 < 20; fi113++ {
            if false {
                if false {
                    switch "hello" {
                    case "\n\t":
                        v5 += v5
                    }
                    if false {
                        var v114 uint8 = (v97 * uint8(v95))
                        v6["hello"] = "foo bar"
                        var v115 uint32 = 42
                        var v116 int32 = (v3 + v3)
                        v95 = (v95 + v95)
                        _ = v114
                        _ = v115
                        _ = v116
                    } else {
                        if false {
                            v4 = (v4 % (((v4 << ((4294967295) & 63))) | 1))
                            var v117 bool = (uint64(v4) <= (v96 >> (((v4 ^ (^(v4)))) & 63)))
                            if v117 {
                                if len(v1) > 0 {
                                    v1[0] = 565
                                }
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if (-128 > 123) {
                                switch "foo bar" {
                                case "hello":
                                    v118 := GenS11{F0: v5}
                                    _ = v118
                                default:
                                    v97 = v97
                                }
                            } else {
                                v117 = false
                                v119 := GenS12{F0: ((!v117) && ((923 != 180) || false))}
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                _ = v119
                            }
                            _ = v117
                        }
                        var v120 uint = v4
                        _ = v120
                    }
                }
                v95 = v95
            } else {
                if true {
                    v121 := GenS13{F0: 0}
                    switch ("hello" + ("a" + "a")) {
                    case "":
                        v96 -= (v96 << ((42) & 63))
                    case "foo bar":
                        var v122 uint64 = (v96 - uint64(v97))
                        v3 += (v3 << ((42) & 63))
                        _ = v122
                    case "a":
                        switch "" {
                        case "":
                            v123 := GenS14{F0: 1}
                            _ = v123
                        case "hello":
                            _ = len(v1)
                        }
                    default:
                        v4 += uint(v95)
                    }
                    _ = v121
                } else {
                    v1 = append(v1, 121)
                }
                v97 = uint8(v95)
                v124 := GenS15{F0: 32767, F1: int16(v3), F2: false}
                _ = v124
            }
            var v125 int16 = 1
            v1 = append(v1, 42)
            _ = fi113
            _ = v125
        }
        var v126 float32 = 0.1
        if true {
            _ = len(v1)
            if ((v95 << (((v4 - 42)) & 63)) >= (v95 * (v95 * 27))) {
                var v127 int8 = 42
                v1 = append(v1, 65535)
                _ = v127
            }
            fmt.Println(v126)
        } else {
            var v128 float32 = (((((v126 - (float32(3.14159) - (-2.25 / v126))) + (v126 - 0.0)) + v126) / ((v126 * v126) + 0.0)) + v126)
            _ = v128
        }
        _ = fi112
        _ = v126
    }
    var v129 int64 = v5
    v95 -= (v95 % (((v95 * (v95 << ((v4) & 63)))) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v129:", v129)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v101
    _ = v103
    _ = v129
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v95
    _ = v96
    _ = v97
}
