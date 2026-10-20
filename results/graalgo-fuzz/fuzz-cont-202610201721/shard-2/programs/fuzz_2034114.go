package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 string; F1 int8; F2 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint }
type GenS7 struct { F0 uint16; F1 int32 }

func main() {
    var v1 bool = false
    v2 := []int{-1, 9223372036854775807, 538}
    v3 := []uint32{}
    v4 := make(map[string]int)
    v5 := []uint32{4294967295}
    var v6 int16 = 42
    fmt.Println(v6)
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = 1000
        var v8 int64 = -9223372036854775808
        var v9 float64 = float64((1.5 / 1.5))
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 127 {
    case 1:
        if false {
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        } else {
            var v10 int64 = -9223372036854775808
            _ = v10
        }
    case -883:
        v11 := GenS1{F0: v6, F1: 2147483647}
        _ = v11
    case 2147483647:
        if (2147483647 <= 42) {
            switch 638 {
            case 32768:
                var v12 int64 = 67
                _ = v12
            case 128:
                v1 = (0.0 <= ((3.14159 * (-2.25 + (0.1 + (0.0 + -2.25)))) + (float64((-2.25 - ((0.1 + (((((1.5 + 1e10) * 1.5) * ((1e10 - -2.25) + (float32(0.0) + 0.1))) + (3.14159 / (1.5 + (0.1 - 0.1)))) + ((0.0 * 3.14159) * 0.1))) - 0.1))) / (1e10 - -2.25))))
            }
            if len(v5) > 0 {
                v5[0] = 0
            }
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                switch "hello" {
                case "\n\t":
                    fmt.Println(fi13, v1)
                case "":
                    if (2147483647 > 247) {
                        v4["hello"] = fi13
                        for fi14 := 0; fi14 < 5; fi14++ {
                            v6 = (v6 / ((v6) | 1))
                            if len(v3) > 0 {
                                v3[0] = 574
                            }
                            fmt.Println(v1)
                            _ = fi14
                        }
                        var v15 int = 622
                        _ = v15
                    }
                }
                _ = fi13
            }
            if true {
                for fi16 := 0; fi16 < 8; fi16++ {
                    var v17 int64 = 42
                    switch "foo bar" {
                    case "hello":
                        for fi18 := 0; fi18 < 5; fi18++ {
                            v4[(("a" + "\n\t") + "\n\t")] = (fi18 & (fi18 | (fi18 & fi18)))
                            if v1 {
                                v1 = v1
                            }
                            _ = fi18
                        }
                    case "":
                        var v19 int8 = -128
                        _ = v19
                    case "foo bar":
                        v6 = (v6 | 0)
                    default:
                        v20 := GenS2{F0: 2, F1: 694}
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
                switch "foo bar" {
                case "foo bar":
                    var v21 int = -68
                    _ = v21
                case "hello":
                    if v1 {
                        switch -2147483648 {
                        case 32767:
                            v1 = v1
                        case -1:
                            fmt.Println(v6)
                        default:
                            var v22 int16 = v6
                            _ = v22
                        }
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    } else {
                        if ((!v1) && false) {
                            v1 = (9223372036854775807 >= -423)
                            v23 := GenS1{F0: 666, F1: 1}
                            for fi24 := 0; fi24 < 7; fi24++ {
                                delete(v4, "foo bar")
                                if (((-2.25 - float32(0.1)) - float32(0.1)) < 1.5) {
                                    if (412 != 318) {
                                        fmt.Println(fi24, fi24, fi24)
                                        switch 2147483647 {
                                        case 1:
                                            for fi25 := 0; fi25 < 7; fi25++ {
                                                if (1e10 >= -2.25) {
                                                    v1 = (false && (!(1.5 > -2.25)))
                                                } else {
                                                    v26 := v4["a"]
                                                    switch (v26 / (((v26 % ((378) | 1))) | 1)) {
                                                    case 128:
                                                        v6 += v6
                                                    case -94:
                                                        var v27 bool = (1e10 < 0.1)
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                }
                                                _ = fi25
                                            }
                                        case 175:
                                            var v28 int = fi24
                                            _ = v28
                                        case 2147483647:
                                            var v29 uint = 796
                                            _ = v29
                                        }
                                        var v30 float64 = 0.1
                                        _ = v30
                                    }
                                    v6 = v6
                                } else {
                                    for fi31 := 0; fi31 < 17; fi31++ {
                                        v1 = (823 < 42)
                                        if (v1 && (((2 >= 255) && (2147483647 <= 2147483647)) && v1)) {
                                            var v32 bool = (true && false)
                                            var v33 uint16 = 65535
                                            if (217 <= 164) {
                                                var v34 float64 = ((-2.25 - 1e10) * 3.14159)
                                                v33 += 32767
                                                _ = v34
                                            } else {
                                                if v1 {
                                                    var v35 float64 = (1e10 * float64(1e10))
                                                    _ = v35
                                                } else {
                                                    switch ("hello" + ("foo bar" + "")) {
                                                    case "":
                                                        v36 := GenS2{F0: v33, F1: 1}
                                                        _ = v36
                                                    case "hello":
                                                        v37 := GenS1{F0: (^(v6)), F1: 0}
                                                        _ = v37
                                                    case "\n\t":
                                                        v4[("a" + "a")] = (fi24 << ((2) & 63))
                                                    }
                                                    fmt.Println(v33, v32, v6)
                                                }
                                                var v38 uint16 = v33
                                                _ = v38
                                            }
                                            _ = v32
                                            _ = v33
                                        } else {
                                            v1 = (231 == 127)
                                            var v39 uint64 = 200
                                            _ = v39
                                        }
                                        var v40 float32 = (0.1 + float32((3.14159 + (3.14159 / 1.5))))
                                        _ = fi31
                                        _ = v40
                                    }
                                }
                                if (!v1) {
                                    v41 := GenS2{F0: 0, F1: 357}
                                    _ = v41
                                } else {
                                    switch ("" + "foo bar") {
                                    case "\n\t":
                                        if v1 {
                                            for fi42 := 0; fi42 < 20; fi42++ {
                                                delete(v4, "hello")
                                                var v43 uint64 = 993
                                                v1 = (fi42 < 380)
                                                switch "" {
                                                case "a":
                                                    switch -803 {
                                                    case 2147483647:
                                                        var v44 uint64 = 1
                                                        _ = v44
                                                    case 0:
                                                        if (434 == 65535) {
                                                            if (!v1) {
                                                                _ = len(v4)
                                                            } else {
                                                                var v45 uint32 = 65535
                                                                v46 := GenS2{F0: 255, F1: 1}
                                                                if v1 {
                                                                    if v1 {
                                                                        v4[("" + "foo bar")] = 9223372036854775807
                                                                    } else {
                                                                        var v47 uint = 230
                                                                        _ = v47
                                                                    }
                                                                    var v48 bool = true
                                                                    v1 = ((true && (v1 && (!v1))) || (216 != 651))
                                                                    _ = v48
                                                                }
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        } else {
                                                            fmt.Println(v43, fi24, v1)
                                                            if false {
                                                                v49 := GenS1{F0: v6, F1: 4294967295}
                                                                if true {
                                                                    _ = len(v4)
                                                                }
                                                                var v50 uint32 = 2147483647
                                                                _ = v49
                                                                _ = v50
                                                            } else {
                                                                v51 := GenS1{F0: (v6 + 900), F1: 42}
                                                                v52 := GenS2{F0: 0, F1: 255}
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                        }
                                                    }
                                                    var v53 float32 = 3.14159
                                                    _ = v53
                                                case "foo bar":
                                                    v6 = v6
                                                case "\n\t":
                                                    switch (fi24 * fi24) {
                                                    case 2147483647:
                                                        fmt.Println(fi42)
                                                    case -2147483648:
                                                        v1 = ("a" <= "")
                                                    case 127:
                                                        v54 := GenS2{F0: 25, F1: 279}
                                                        _ = v54
                                                    }
                                                default:
                                                    if v1 {
                                                        var v55 float32 = float32(1e10)
                                                        _ = v55
                                                    } else {
                                                        switch -2147483648 {
                                                        case 659:
                                                            v56 := v4[""]
                                                            _ = v56
                                                        case 4294967295:
                                                            if false {
                                                                var v57 int16 = -1
                                                                _ = v57
                                                            }
                                                        default:
                                                            var v58 int = -1
                                                            _ = v58
                                                        }
                                                    }
                                                }
                                                v59 := GenS2{F0: 488, F1: 316}
                                                _ = fi42
                                                _ = v43
                                                _ = v59
                                            }
                                        }
                                    }
                                    var v60 string = "foo bar"
                                    for fi61 := 0; fi61 < 3; fi61++ {
                                        v60 = "hello"
                                        var v62 float32 = 1.5
                                        _ = fi61
                                        _ = v62
                                    }
                                    _ = v60
                                }
                                _ = fi24
                            }
                            _ = v23
                        }
                        if true {
                            for fi63 := 0; fi63 < 17; fi63++ {
                                v6 = 1
                                v6 += v6
                                _ = fi63
                            }
                            var v64 uint32 = 2
                            var v65 int = -720
                            _ = v64
                            _ = v65
                        }
                        if v1 {
                            v66 := v4[("\n\t" + ((("" + "foo bar") + "\n\t") + ((("foo bar" + "hello") + "hello") + ("foo bar" + "a"))))]
                            for fi67 := 0; fi67 < 19; fi67++ {
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v66, v66, v66)
                                default:
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v6, v1)
                                        case "":
                                            switch fi67 {
                                            case -751:
                                                fmt.Println(v1, v66)
                                            case 2147483647:
                                                switch 649 {
                                                case 666:
                                                    v4[("" + ((("\n\t" + "\n\t") + "") + "foo bar"))] = (^(v66))
                                                case 2147483647:
                                                    var v68 uint64 = 74
                                                    _ = v68
                                                }
                                            default:
                                                var v69 float64 = (1.5 / 1.0)
                                                _ = v69
                                            }
                                        default:
                                            var v70 uint64 = 2
                                            _ = v70
                                        }
                                    } else {
                                        var v71 string = "\n\t"
                                        _ = v71
                                    }
                                }
                                if true {
                                    v6 = 1000
                                    var v72 int = 128
                                    _ = v72
                                }
                                v73 := GenS3{F0: "", F1: 42}
                                _ = fi67
                                _ = v73
                            }
                            _ = v66
                        }
                    }
                }
                if v1 {
                    for fi74 := 0; fi74 < 5; fi74++ {
                        var v75 uint32 = 2
                        v76 := GenS3{F0: "hello", F1: 27}
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                    v77 := GenS2{F0: 698, F1: 0}
                    v6 += 0
                    _ = v77
                }
            } else {
                if v1 {
                    v1 = (0 < 331)
                }
                fmt.Println(v1, v6)
                v1 = false
            }
        }
    }
    switch 264 {
    case 4294967295:
        if len(v5) > 0 {
            v5[0] = 2
        }
    case -1:
        fmt.Println(v6, v6)
    default:
        fmt.Println(v6, v1)
    }
    var v78 uint = 42
    fmt.Println(v6)
    for fi79 := 0; fi79 < 19; fi79++ {
        v80 := GenS3{F0: ("\n\t" + "\n\t"), F1: 128}
        _ = fi79
        _ = v80
    }
    var v81 int8 = 63
    var v82 uint16 = 42
    v3 = append(v3, uint32(v82))
    for fi83 := 0; fi83 < 7; fi83++ {
        for fi84 := 0; fi84 < 18; fi84++ {
            v6 = v6
            var v85 int = (fi83 + (-(fi84)))
            v82 = v82
            _ = fi84
            _ = v85
        }
        _ = len(v4)
        var v86 int16 = (v6 + v6)
        _ = fi83
        _ = v86
    }
    switch "foo bar" {
    case "":
        for fi87 := 0; fi87 < 1; fi87++ {
            v88 := GenS2{F0: (v82 | v82), F1: (v78 & (v78 >> ((v78) & 63)))}
            _ = fi87
            _ = v88
        }
    case "foo bar":
        for fi89 := 0; fi89 < 5; fi89++ {
            var v90 int32 = 969
            var v91 uint32 = 913
            if false {
                var v92 float64 = float64(((0.1 - (0.1 - ((1.5 / 1.0) - 0.1))) * (3.14159 + (1e10 + 1.5))))
                if ((v92 + 3.14159) >= float64(((-2.25 + 0.0) + ((((0.1 + (3.14159 / (float32(((v92 - (1e10 - (1e10 + v92))) * 0.1)) + (3.14159 / float32(v92))))) * 0.0) + 0.0) * 3.14159)))) {
                    v2 = append(v2, fi89)
                    switch "a" {
                    case "hello":
                        v93 := GenS2{F0: v82, F1: (v78 % ((v78) | 1))}
                        _ = v93
                    default:
                        v94 := GenS3{F0: "", F1: 1}
                        _ = v94
                    }
                }
                _ = v92
            }
            _ = fi89
            _ = v90
            _ = v91
        }
    }
    v95 := GenS1{F0: (v6 & (v6 | v6)), F1: 864}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v78 += v78
    fmt.Println(v81)
    switch (("hello" + "") + "foo bar") {
    case "a":
        var v96 float32 = (((1e10 + float32(0.1)) + 1e10) * 0.1)
        _ = v96
    case "foo bar":
        switch 144 {
        case 127:
            fmt.Println(v1)
        case 309:
            fmt.Println(v1, v81, v78)
        case 2147483647:
            if (!(301 >= v78)) {
                if false {
                    var v97 uint = (v78 | v78)
                    _ = v97
                } else {
                    fmt.Println(v82, v82)
                    v82 = (v82 & (v82 & v82))
                }
                v98 := v4["a"]
                _ = v98
            } else {
                for fi99 := 0; fi99 < 16; fi99++ {
                    if ((("hello" + "a") + "foo bar") == "foo bar") {
                        switch ("hello" + ((("a" + ("\n\t" + ("\n\t" + "hello"))) + "hello") + "hello")) {
                        case "\n\t":
                            v78 += (v78 - v78)
                        case "foo bar":
                            switch -1 {
                            case 128:
                                for fi100 := 0; fi100 < 4; fi100++ {
                                    v6 = v6
                                    v78 -= 4294967295
                                    v101 := GenS3{F0: "a", F1: 0}
                                    _ = fi100
                                    _ = v101
                                }
                            case -969:
                                if v1 {
                                    if v1 {
                                        if false {
                                            v102 := GenS4{F0: "hello", F1: v81, F2: (v81 >> ((65535) & 63))}
                                            v1 = (!v1)
                                            _ = v102
                                        } else {
                                            var v103 float64 = 0.1
                                            v104 := v4["\n\t"]
                                            _ = v103
                                            _ = v104
                                        }
                                    } else {
                                        for fi105 := 0; fi105 < 1; fi105++ {
                                            if (("foo bar" + "") >= "a") {
                                                if v1 {
                                                    var v106 uint64 = 459
                                                    var v107 float32 = (((3.14159 / (((0.1 - 0.1) / 1e10) - 0.1)) - 0.0) + 1.5)
                                                    _ = v106
                                                    _ = v107
                                                }
                                                if (("\n\t" + "") != ("" + "foo bar")) {
                                                    v108 := GenS4{F0: ("hello" + "foo bar"), F1: 116, F2: (v81 >> ((4294967295) & 63))}
                                                    delete(v4, "a")
                                                    _ = v108
                                                }
                                                if (v1 || (!v1)) {
                                                    _ = len(v4)
                                                }
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = -326
                                            }
                                            _ = fi105
                                        }
                                        if (!(42 >= 18446744073709551615)) {
                                            fmt.Println(v82, v1, v82)
                                            for fi109 := 0; fi109 < 2; fi109++ {
                                                switch "foo bar" {
                                                case "":
                                                    v2 = append(v2, 128)
                                                case "\n\t":
                                                    var v110 uint32 = 501
                                                    _ = v110
                                                case "foo bar":
                                                    if len(v2) > 0 {
                                                        v2[0] = fi109
                                                    }
                                                default:
                                                    var v111 int16 = v6
                                                    fmt.Println(v1, v111)
                                                    _ = v111
                                                }
                                                var v112 uint8 = 0
                                                if v1 {
                                                    v112 -= 128
                                                } else {
                                                    v113 := GenS5{F0: 0}
                                                    fmt.Println(v81, v1, v112)
                                                    _ = v113
                                                }
                                                _ = fi109
                                                _ = v112
                                            }
                                        } else {
                                            v1 = v1
                                            if false {
                                                for fi114 := 0; fi114 < 4; fi114++ {
                                                    v115 := GenS6{F0: (v6 + (v6 ^ (v6 - (^(v6))))), F1: 140, F2: v78}
                                                    _ = fi114
                                                    _ = v115
                                                }
                                                if (!(v82 > v82)) {
                                                    var v116 uint = v78
                                                    var v117 int16 = (v6 | (v6 + (-(v6))))
                                                    fmt.Println(fi99, v117, v81)
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    if ((true || v1) && (!v1)) {
                                                        for fi118 := 0; fi118 < 6; fi118++ {
                                                            fmt.Println(fi118)
                                                            _ = fi118
                                                        }
                                                    }
                                                    var v119 bool = v1
                                                    v78 = 588
                                                    _ = v119
                                                }
                                                v4["hello"] = (fi99 ^ (fi99 >> ((v78) & 63)))
                                            }
                                        }
                                    }
                                    var v120 int8 = (v81 | -128)
                                    switch "\n\t" {
                                    case "a":
                                        v120 = -128
                                    case "foo bar":
                                        var v121 uint8 = 94
                                        _ = v121
                                    case "\n\t":
                                        var v122 string = "hello"
                                        _ = v122
                                    default:
                                        var v123 int32 = 475
                                        _ = v123
                                    }
                                    _ = v120
                                } else {
                                    var v124 uint = v78
                                    _ = v124
                                }
                            default:
                                v125 := GenS7{F0: v82, F1: -2147483648}
                                _ = v125
                            }
                        default:
                            var v126 int8 = v81
                            _ = v126
                        }
                    }
                    v1 = v1
                    if (!v1) {
                        v6 = (v6 | v6)
                        var v127 int = 1
                        _ = v127
                    }
                    _ = fi99
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v82 = (v82 % (((v82 - 65535)) | 1))
                case "foo bar":
                    v82 = (^(v82))
                case "a":
                    v128 := GenS1{F0: v6, F1: 466}
                    _ = v128
                }
                v78 -= (v78 * v78)
            }
        }
    case "\n\t":
        v129 := GenS5{F0: (v81 * v81)}
        _ = v129
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v78
    _ = v81
    _ = v82
    _ = v95
}
