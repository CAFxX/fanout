package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int }
type GenS9 struct { F0 int; F1 uint64; F2 bool }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 uint }
func genFn0() int32 {
    if ((18446744073709551615 > 4294967295) || false) {
        var v4 uint64 = 0
        _ = v4
    }
    if (true && ((-615 != 1) || (42 > 0))) {
        switch "hello" {
        case "":
            var v5 uint64 = 18446744073709551615
            _ = v5
        case "\n\t":
            if (true || true) {
                var v6 bool = false
                v6 = (!v6)
                for fi7 := 0; fi7 < 16; fi7++ {
                    if (!v6) {
                        for fi8 := 0; fi8 < 11; fi8++ {
                            switch (fi7 * (fi7 % (((fi8 | (fi8 << ((1) & 63)))) | 1))) {
                            case 1:
                                v9 := []int{}
                                switch fi7 {
                                case 42:
                                    v10 := GenS1{F0: fi7, F1: "", F2: 100}
                                    _ = v10
                                }
                                var v11 int32 = 0
                                _ = v11
                                _ = v9
                            }
                            v12 := GenS1{F0: fi8, F1: "foo bar", F2: 1}
                            v6 = v6
                            _ = fi8
                            _ = v12
                        }
                    }
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 int16 = -322
                        var v15 string = "a"
                        v6 = true
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = fi7
                }
                _ = v6
            }
        case "hello":
            var v16 int8 = -106
            _ = v16
        default:
            var v17 uint = 42
            _ = v17
        }
    } else {
        v18 := []int{-745}
        _ = v18
    }
    return 669
}

func main() {
    var v1 uint8 = 255
    var v2 float32 = 0.0
    v3 := make(map[int]uint32)
    if (-9223372036854775808 < -9223372036854775808) {
        switch 240 {
        case -212:
            v19 := []int{128, -1}
            _ = v19
        case 9223372036854775807:
            v20 := GenS1{F0: 128, F1: "a", F2: 38}
            _ = v20
        case 4294967295:
            v21 := GenS2{F0: 255}
            _ = v21
        default:
            v1 = v1
        }
        _ = len(v3)
        var v22 uint32 = 338
        _ = v22
    } else {
        fmt.Println(v2)
        v23 := GenS1{F0: 128, F1: ("a" + ("\n\t" + "\n\t")), F2: -1}
        _ = v23
    }
    if (!true) {
        for fi24 := 0; fi24 < 7; fi24++ {
            if (449 == 0) {
                if true {
                    fmt.Println(v2)
                    v1 = v1
                    var v25 uint = 0
                    _ = v25
                } else {
                    v1 += (v1 << ((uint(v1)) & 63))
                    var v26 uint16 = 255
                    var v27 uint16 = v26
                    _ = v26
                    _ = v27
                }
            }
            v2 -= (-2.25 + float32((float64(-2.25) / -2.25)))
            _ = fi24
        }
        fmt.Println(v2, v1)
    } else {
        if ((42 != -1) || (("a" + "\n\t") <= "")) {
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 uint = 509
                if false {
                    v1 = 2
                    var v30 uint32 = 207
                    var v31 uint16 = 65535
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 10; fi32++ {
                        v33 := GenS1{F0: (fi28 + 9223372036854775807), F1: (("hello" + (("hello" + "a") + "\n\t")) + "a"), F2: 19}
                        _ = fi32
                        _ = v33
                    }
                    for fi34 := 0; fi34 < 4; fi34++ {
                        if true {
                            if true {
                                switch 0 {
                                case 9223372036854775807:
                                    v35 := GenS1{F0: fi28, F1: "", F2: -1}
                                    if (!true) {
                                        v36 := GenS1{F0: fi34, F1: "\n\t", F2: 127}
                                        v37 := GenS3{F0: 2147483647}
                                        _ = v36
                                        _ = v37
                                    } else {
                                        if false {
                                            v38 := GenS3{F0: 783}
                                            fmt.Println(v2)
                                            if false {
                                                var v39 string = "hello"
                                                _ = v39
                                            } else {
                                                var v40 uint8 = 0
                                                _ = v40
                                            }
                                            _ = v38
                                        }
                                        v41 := GenS3{F0: 2147483648}
                                        v42 := GenS4{F0: (("hello" + ("\n\t" + "foo bar")) + ((("" + "") + ((("foo bar" + "foo bar") + (("" + "") + "hello")) + "\n\t")) + ("\n\t" + (("hello" + "") + "")))), F1: 0}
                                        _ = v41
                                        _ = v42
                                    }
                                    if (-2147483648 > -623) {
                                        v43 := GenS3{F0: 1}
                                        _ = v43
                                    }
                                    _ = v35
                                default:
                                    var v44 uint8 = (v1 % ((uint8(v29)) | 1))
                                    _ = v44
                                }
                            }
                            switch "hello" {
                            case "":
                                var v45 int8 = 1
                                _ = v45
                            case "\n\t":
                                if (!(449 == 255)) {
                                    v46 := []int{(^(fi34))}
                                    var v47 string = (("foo bar" + ("foo bar" + "")) + ("hello" + ("hello" + "\n\t")))
                                    var v48 uint16 = 42
                                    switch v47 {
                                    case "\n\t":
                                        v47 = (v47 + v47)
                                    case "foo bar":
                                        if (!(uint16(v1) == v48)) {
                                            switch "hello" {
                                            case "foo bar":
                                                if false {
                                                    v29 -= v29
                                                    switch (fi28 * (fi34 >> (((v29 + (v29 % ((822) | 1)))) & 63))) {
                                                    case 962:
                                                        v49 := GenS2{F0: 42}
                                                        _ = v49
                                                    case 42:
                                                        var v50 string = v47
                                                        var v51 uint32 = 4294967295
                                                        v1 = (^(v1))
                                                        _ = v50
                                                        _ = v51
                                                    default:
                                                        var v52 uint8 = (^(v1))
                                                        _ = v52
                                                    }
                                                } else {
                                                    fmt.Println(fi28, v48, v48)
                                                    var v53 uint64 = 42
                                                    v48 -= (v48 & v48)
                                                    v1 += (v1 - (^(v1)))
                                                    switch (fi34 * fi28) {
                                                    case 32767:
                                                        var v54 uint32 = uint32(v48)
                                                        _ = v54
                                                    case 475:
                                                        v29 = v29
                                                    }
                                                    v48 += v48
                                                    v55 := v3[fi34]
                                                    if true {
                                                        if false {
                                                            if false {
                                                                v47 = v47
                                                            }
                                                            var v56 int16 = 207
                                                            if ((v56 | -1) >= (v56 ^ (v56 << (((v29 / (((^(v29))) | 1))) & 63)))) {
                                                                v57 := GenS1{F0: fi34, F1: "a", F2: -1}
                                                                var v58 float64 = float64(float32(0.1))
                                                                if (!(false || (false || true))) {
                                                                    var v59 uint64 = 735
                                                                    v60 := GenS1{F0: fi28, F1: (v47 + ("hello" + v47)), F2: 42}
                                                                    delete(v3, fi34)
                                                                    var v61 bool = true
                                                                    if (!(!(!true))) {
                                                                        var v62 float32 = v2
                                                                        var v63 string = ""
                                                                        _ = v62
                                                                        _ = v63
                                                                    } else {
                                                                        delete(v3, fi34)
                                                                    }
                                                                    v3[(fi28 - fi34)] = v55
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                                _ = v57
                                                                _ = v58
                                                            } else {
                                                                switch 454 {
                                                                case 2147483647:
                                                                    var v64 int16 = (v56 % ((1) | 1))
                                                                    _ = v64
                                                                case 1:
                                                                    v65 := GenS2{F0: v29}
                                                                    _ = v65
                                                                case 4294967295:
                                                                    if true {
                                                                        v29 -= (v29 - v29)
                                                                    } else {
                                                                        switch 0 {
                                                                        case 519:
                                                                            var v66 float32 = (v2 / ((v2 - (float32(((1.5 / float64(v2)) - (-2.25 * -2.25))) - (1.5 - 0.0))) + (float32(3.14159) - ((((v2 * v2) - (v2 - (0.0 * v2))) + v2) * v2))))
                                                                            _ = v66
                                                                        case 280:
                                                                            if ((uint8(v55) >= (^(v1))) || (false && true)) {
                                                                                switch 4294967295 {
                                                                                case -1:
                                                                                    v67 := GenS5{F0: (v47 + (("\n\t" + v47) + (v47 + "hello")))}
                                                                                    _ = v67
                                                                                case 42:
                                                                                    v68 := v3[-9223372036854775808]
                                                                                    _ = v68
                                                                                default:
                                                                                    fmt.Println(v47)
                                                                                }
                                                                                var v69 uint = 18446744073709551615
                                                                                _ = v69
                                                                            } else {
                                                                                v29 += (v29 & uint(v1))
                                                                            }
                                                                        case 127:
                                                                            if true {
                                                                                _ = len(v46)
                                                                            } else {
                                                                                if (!false) {
                                                                                    v70 := GenS6{F0: (v48 % ((65535) | 1)), F1: (^(v1))}
                                                                                    v56 += 42
                                                                                    _ = v70
                                                                                }
                                                                            }
                                                                        }
                                                                        if (1 != (v56 | 0)) {
                                                                            var v71 bool = false
                                                                            var v72 int16 = 1
                                                                            var v73 float64 = -2.25
                                                                            v74 := GenS4{F0: v47, F1: 1}
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                            _ = v74
                                                                        } else {
                                                                            v75 := GenS1{F0: fi34, F1: ("hello" + v47), F2: 127}
                                                                            _ = v75
                                                                        }
                                                                        v55 -= (^(v55))
                                                                    }
                                                                default:
                                                                    var v76 int64 = 0
                                                                    _ = v76
                                                                }
                                                                v77 := GenS7{F0: 2147483648}
                                                                v29 = uint(v48)
                                                                _ = v77
                                                            }
                                                            v78 := GenS4{F0: ("" + (v47 + v47)), F1: (v55 * (v55 ^ 0))}
                                                            _ = v56
                                                            _ = v78
                                                        }
                                                    }
                                                    if false {
                                                        var v79 uint64 = (^(v53))
                                                        var v80 uint16 = 32767
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                    _ = v53
                                                    _ = v55
                                                }
                                            case "a":
                                                var v81 int = (fi34 | fi28)
                                                _ = v81
                                            case "hello":
                                                if (false || true) {
                                                    if (!true) {
                                                        var v82 float32 = (v2 + ((3.14159 * v2) + v2))
                                                        if (2 == 0) {
                                                            var v83 uint16 = (v48 + v48)
                                                            v84 := GenS6{F0: v83, F1: (v1 - v1)}
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        if false {
                                                            if ((1000 < -353) || (true || false)) {
                                                                v85 := GenS4{F0: v47, F1: 0}
                                                                var v86 float32 = v2
                                                                _ = v85
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS6{F0: (^(v48)), F1: 0}
                                                                v88 := GenS7{F0: 56}
                                                                v82 = (1.5 + v2)
                                                                v89 := GenS1{F0: (fi28 / ((2147483647) | 1)), F1: v47, F2: -126}
                                                                var v90 int = (fi28 + fi34)
                                                                _ = v87
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                        }
                                                        _ = v82
                                                    } else {
                                                        fmt.Println(v47)
                                                        switch (v47 + (v47 + ("foo bar" + "a"))) {
                                                        case "hello":
                                                            v1 = (v1 | v1)
                                                        case "\n\t":
                                                            if false {
                                                                v47 = v47
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(v47)
                                                                    }
                                                                    v91 := GenS1{F0: (^(fi34)), F1: ("hello" + "hello"), F2: 1}
                                                                    _ = v91
                                                                }
                                                            } else {
                                                                if false {
                                                                    fmt.Println(fi34)
                                                                    var v92 string = v47
                                                                    switch (fi34 ^ (fi28 ^ (^(fi28)))) {
                                                                    case 32767:
                                                                        delete(v3, fi34)
                                                                    case 0:
                                                                        var v93 uint32 = 0
                                                                        _ = v93
                                                                    default:
                                                                        fmt.Println(v1, v47, v47)
                                                                    }
                                                                    _ = v92
                                                                } else {
                                                                    fmt.Println(v29, v29, v47)
                                                                    v48 = v48
                                                                }
                                                            }
                                                        default:
                                                            v46 = append(v46, fi28)
                                                        }
                                                        v94 := GenS5{F0: v47}
                                                        _ = v94
                                                    }
                                                } else {
                                                    if (!((v2 != v2) && false)) {
                                                        var v95 uint8 = v1
                                                        fmt.Println(fi28, v48)
                                                        if len(v46) > 0 {
                                                            v46[0] = (fi34 * 1)
                                                        }
                                                        _ = v95
                                                    }
                                                }
                                            }
                                        } else {
                                            _ = len(v3)
                                            v96 := GenS5{F0: "hello"}
                                            _ = v96
                                        }
                                    }
                                    fmt.Println(fi28)
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                            case "foo bar":
                                switch fi28 {
                                case 42:
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v2)
                                        switch fi28 {
                                        case 4294967295:
                                            v97 := GenS8{F0: v1, F1: (fi28 >> (((v29 >> ((v29) & 63))) & 63))}
                                            _ = v97
                                        }
                                        if (-128 == -100) {
                                            v98 := GenS9{F0: fi28, F1: 42, F2: false}
                                            v3[fi34] = 2
                                            v99 := GenS10{F0: 2, F1: true}
                                            v2 = 1e10
                                            _ = v98
                                            _ = v99
                                        }
                                    case "foo bar":
                                        fmt.Println(v29, v1, v29)
                                    default:
                                        v100 := GenS9{F0: 0, F1: 201, F2: (1 >= -1)}
                                        _ = v100
                                    }
                                case -2147483648:
                                    v2 += 0.0
                                case -9223372036854775808:
                                    var v101 uint64 = 4294967295
                                    v2 += 1e10
                                    v1 = v1
                                    _ = v101
                                }
                            default:
                                v29 -= 0
                            }
                            v29 = uint(v1)
                        } else {
                            if (42 >= 0) {
                                switch 32768 {
                                case -709:
                                    v1 += 42
                                case 4294967295:
                                    v102 := []int{(fi34 + int(v1)), (fi34 / ((fi28) | 1))}
                                    _ = v102
                                case -56:
                                    v1 += 255
                                default:
                                    if false {
                                        var v103 int16 = -660
                                        v103 = 32767
                                        _ = v103
                                    } else {
                                        v104 := v3[(fi28 & fi28)]
                                        v1 = (v1 + 1)
                                        var v105 uint8 = v1
                                        _ = v104
                                        _ = v105
                                    }
                                }
                            }
                            if (-15 <= -128) {
                                v106 := GenS6{F0: 32767, F1: 1}
                                if true {
                                    var v107 uint8 = (v1 + v1)
                                    _ = v107
                                } else {
                                    if (0 <= 1) {
                                        if true {
                                            var v108 uint64 = 1
                                            v29 = 0
                                            _ = v108
                                        } else {
                                            var v109 uint16 = 416
                                            switch (fi28 | (fi34 | (fi28 - 591))) {
                                            case 161:
                                                var v110 uint32 = 0
                                                _ = v110
                                            case 42:
                                                fmt.Println(v1)
                                            }
                                            v29 = (v29 * v29)
                                            _ = v109
                                        }
                                        v1 = (^(v1))
                                    }
                                    var v111 bool = (!(-32768 != 218))
                                    _ = v111
                                }
                                v112 := GenS11{F0: (v29 ^ v29)}
                                _ = v106
                                _ = v112
                            } else {
                                var v113 bool = (false || ((fi34 % (((fi34 | (-(fi34)))) | 1)) <= (fi28 + 32767)))
                                _ = v113
                            }
                        }
                        _ = fi34
                    }
                }
                _ = fi28
                _ = v29
            }
            switch "hello" {
            case "hello":
                switch 2147483647 {
                case 680:
                    v3[127] = 0
                default:
                    v114 := []int{32767, 9223372036854775807}
                    _ = v114
                }
            case "foo bar":
                if (18446744073709551615 > 604) {
                    v2 = (((1e10 - v2) + v2) + v2)
                } else {
                    var v115 uint16 = 141
                    _ = v115
                }
            case "":
                switch "" {
                case "":
                    if (0 > 1000) {
                        v116 := []int{-1, -1}
                        var v117 float64 = 1e10
                        v1 -= (v1 << ((4294967295) & 63))
                        _ = v116
                        _ = v117
                    }
                case "\n\t":
                    var v118 string = "foo bar"
                    _ = v118
                default:
                    if (2147483647 > 42) {
                        fmt.Println(v2)
                        switch -2147483648 {
                        case -9223372036854775808:
                            v1 -= v1
                        default:
                            v119 := GenS7{F0: 567}
                            _ = v119
                        }
                        if (!false) {
                            var v120 int64 = -742
                            _ = v120
                        } else {
                            var v121 float32 = v2
                            _ = v121
                        }
                    } else {
                        var v122 uint64 = 42
                        fmt.Println(v122, v1)
                        _ = v122
                    }
                }
            }
        }
    }
    v123 := []int{-2147483648, 764, -5}
    if len(v123) > 0 {
        v123[0] = -1
    }
    for fi124 := 0; fi124 < 4; fi124++ {
        if (377 < 342) {
            _ = len(v3)
            v1 = (v1 / (((v1 & 127)) | 1))
        }
        var v125 int8 = -128
        _ = fi124
        _ = v125
    }
    switch "\n\t" {
    case "a":
        for fi126 := 0; fi126 < 4; fi126++ {
            for fi127 := 0; fi127 < 5; fi127++ {
                if (395 >= 32767) {
                    var v128 int32 = 2147483647
                    if false {
                        v123 = append(v123, fi127)
                    }
                    _ = v128
                }
                v2 = ((float32(1.5) * 1e10) / 1.5)
                _ = fi127
            }
            _ = fi126
        }
    case "":
        _ = len(v3)
    case "\n\t":
        v1 += (v1 >> ((65535) & 63))
    }
    v129 := v3[0]
    for fi130 := 0; fi130 < 2; fi130++ {
        var v131 bool = false
        var v132 uint64 = 2
        _ = fi130
        _ = v131
        _ = v132
    }
    var v133 string = ("" + "")
    var v134 int64 = -177
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v123:", len(v123), v123)
    fmt.Println("v129:", v129)
    fmt.Println("v133:", v133)
    fmt.Println("v134:", v134)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v123
    _ = v129
    _ = v133
    _ = v134
    _ = v2
    _ = v3
}
