package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32; F1 int32 }
type GenS12 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    var v1 uint8 = 128
    var v2 float64 = float64(0.0)
    var v3 uint32 = 42
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: false, F1: 127}
    fmt.Println(v3)
    switch (("hello" + ("a" + ("foo bar" + ("hello" + "a")))) + "\n\t") {
    case "hello":
        if false {
            if (v2 < v2) {
                v2 -= (3.14159 * v2)
                v2 -= ((v2 - (-2.25 * float64(3.14159))) * (float64(-2.25) - (1e10 / v2)))
            } else {
                v1 += v1
                for fi6 := 0; fi6 < 8; fi6++ {
                    for fi7 := 0; fi7 < 3; fi7++ {
                        switch "a" {
                        case "foo bar":
                            if (105 != uint64(v3)) {
                                if (-247 >= -1) {
                                    if (42 <= uint16(v3)) {
                                        v4["hello"] = 42
                                    } else {
                                        v1 += v1
                                        v8 := []int{}
                                        var v9 uint64 = 1
                                        _ = v8
                                        _ = v9
                                    }
                                    v3 = 2
                                    var v10 int = fi6
                                    _ = v10
                                } else {
                                    var v11 int8 = 1
                                    v1 -= (v1 + 219)
                                    if true {
                                        var v12 uint8 = v1
                                        v13 := []int{850, fi7}
                                        if ((!false) || (uint64(v3) != 18446744073709551615)) {
                                            fmt.Println(fi6, v1, v12)
                                            v1 -= (v12 * v1)
                                            if len(v13) >= 2 {
                                                v13 = v13[1:len(v13)]
                                            }
                                        }
                                        _ = v12
                                        _ = v13
                                    } else {
                                        if ((!true) && (0 > 42)) {
                                            v14 := GenS1{F0: true, F1: fi6}
                                            v15 := GenS2{F0: v11, F1: -1}
                                            _ = v14
                                            _ = v15
                                        }
                                        v2 = float64(((3.14159 - (float32(((-2.25 * 1e10) + 1.5)) - ((0.1 + 3.14159) - -2.25))) + (-2.25 / 3.14159)))
                                        switch fi7 {
                                        case 32768:
                                            if (true && (!(!true))) {
                                                v16 := GenS2{F0: v11, F1: 42}
                                                v1 += v1
                                                v1 = v1
                                                v17 := v4["foo bar"]
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v18 int16 = 32767
                                                    _ = v18
                                                case "\n\t":
                                                    var v19 string = ("foo bar" + ((("a" + "") + "\n\t") + ("a" + "\n\t")))
                                                    var v20 float32 = 1.5
                                                    v21 := GenS1{F0: ((!(!true)) || false), F1: (fi7 ^ (fi7 - 1))}
                                                    var v22 uint32 = (v3 * (v3 % (((v3 / (((v17 * v17)) | 1))) | 1)))
                                                    _ = v19
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                }
                                                _ = v16
                                                _ = v17
                                            }
                                        }
                                    }
                                    v3 += 2147483647
                                    if false {
                                        if (v2 > v2) {
                                            var v23 bool = true
                                            _ = v23
                                        } else {
                                            v2 = v2
                                            v24 := GenS1{F0: (42 == 976), F1: fi7}
                                            _ = v24
                                        }
                                    }
                                    if (!(true && (true || false))) {
                                        v3 = (v3 + v3)
                                    } else {
                                        fmt.Println(v3, fi7)
                                        var v25 uint64 = uint64(v1)
                                        _ = v25
                                    }
                                    _ = v11
                                }
                                var v26 int64 = 0
                                switch "a" {
                                case "":
                                    var v27 string = "\n\t"
                                    _ = v27
                                case "\n\t":
                                    var v28 uint = 42
                                    _ = v28
                                default:
                                    v2 -= v2
                                }
                                fmt.Println(v2, v2)
                                v1 += v1
                                _ = v26
                            }
                        case "a":
                            fmt.Println(fi7, fi7)
                        default:
                            v29 := []int{}
                            _ = v29
                        }
                        _ = fi7
                    }
                    _ = fi6
                }
            }
            v30 := GenS2{F0: 127, F1: -833}
            _ = v30
        } else {
            v3 = (v3 | (v3 ^ (v3 ^ (v3 ^ 0))))
        }
    case "foo bar":
        v2 = ((((v2 - (float64((float32(3.14159) / 1.5)) * v2)) / 0.1) * (v2 * (1.5 * v2))) - v2)
    default:
        delete(v4, "foo bar")
    }
    v2 -= ((v2 + (v2 + v2)) - v2)
    if false {
        v31 := v4[("hello" + "hello")]
        fmt.Println(v3, v3)
        _ = v31
    }
    var v32 int16 = -1
    if false {
        var v33 uint = 1
        for fi34 := 0; fi34 < 4; fi34++ {
            for fi35 := 0; fi35 < 8; fi35++ {
                _ = len(v4)
                if (!true) {
                    v36 := []int{899, 725, (fi35 + (fi34 - (fi35 * (fi34 >> ((v33) & 63))))), (fi35 % ((int(v32)) | 1))}
                    v2 -= 3.14159
                    var v37 float64 = -2.25
                    _ = v36
                    _ = v37
                }
                var v38 int8 = 1
                _ = fi35
                _ = v38
            }
            _ = fi34
        }
        _ = v33
    }
    switch "hello" {
    case "":
        if true {
            if true {
                v39 := GenS2{F0: 75, F1: -637}
                _ = v39
            } else {
                delete(v4, "hello")
                switch "\n\t" {
                case "\n\t":
                    if ((v32 | v32) != -852) {
                        if false {
                            if (true && (0 <= -9223372036854775808)) {
                                var v40 int8 = -88
                                var v41 uint = 2
                                var v42 uint = (^(v41))
                                _ = v40
                                _ = v41
                                _ = v42
                            } else {
                                v43 := GenS3{F0: 128, F1: 1}
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 17; fi44++ {
                                fmt.Println(v2, fi44)
                                var v45 uint64 = 4294967295
                                fmt.Println(v32, v2, v3)
                                _ = fi44
                                _ = v45
                            }
                            v46 := GenS4{F0: (^(v1)), F1: 1}
                            _ = v46
                        } else {
                            v47 := GenS5{F0: 1, F1: 0}
                            v48 := GenS6{F0: v1}
                            for fi49 := 0; fi49 < 16; fi49++ {
                                if false {
                                    if true {
                                        for fi50 := 0; fi50 < 1; fi50++ {
                                            var v51 uint8 = (v1 / ((140) | 1))
                                            _ = fi50
                                            _ = v51
                                        }
                                    }
                                } else {
                                    if true {
                                        fmt.Println(v32, v1)
                                        for fi52 := 0; fi52 < 15; fi52++ {
                                            if ((fi52 | (fi52 ^ -9223372036854775808)) <= 861) {
                                                var v53 int64 = 655
                                                switch 4294967295 {
                                                case -9223372036854775808:
                                                    switch 42 {
                                                    case 1:
                                                        v54 := []int{fi52, fi49, fi52}
                                                        _ = v54
                                                    default:
                                                        v1 += (v1 >> ((65535) & 63))
                                                        v55 := GenS1{F0: (0 <= 16), F1: fi49}
                                                        var v56 uint64 = 0
                                                        v57 := GenS1{F0: true, F1: (fi52 >> ((1) & 63))}
                                                        v58 := []int{}
                                                        v59 := GenS7{F0: v32, F1: 74}
                                                        var v60 uint32 = (^(v3))
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                case -635:
                                                    fmt.Println(v1, v1)
                                                default:
                                                    var v61 uint8 = 4
                                                    _ = v61
                                                }
                                                _ = v53
                                            }
                                            v62 := GenS8{F0: v1}
                                            switch 601 {
                                            case 2147483647:
                                                var v63 uint8 = (^(v1))
                                                _ = v63
                                            }
                                            _ = fi52
                                            _ = v62
                                        }
                                        for fi64 := 0; fi64 < 2; fi64++ {
                                            v65 := GenS4{F0: v1, F1: 128}
                                            v66 := GenS9{F0: 0}
                                            _ = fi64
                                            _ = v65
                                            _ = v66
                                        }
                                    } else {
                                        v1 = (v1 % ((v1) | 1))
                                    }
                                    var v67 uint8 = uint8(v3)
                                    _ = v67
                                }
                                _ = fi49
                            }
                            _ = v47
                            _ = v48
                        }
                        v68 := GenS2{F0: -3, F1: -9223372036854775808}
                        if (42 > 0) {
                            var v69 string = "hello"
                            v70 := GenS3{F0: 9223372036854775807, F1: 794}
                            _ = v69
                            _ = v70
                        } else {
                            v71 := GenS4{F0: (v1 % ((v1) | 1)), F1: -542}
                            v3 = (v3 ^ 4294967295)
                            for fi72 := 0; fi72 < 11; fi72++ {
                                v73 := GenS1{F0: (((((("foo bar" + (("hello" + "foo bar") + "hello")) + "a") + "hello") <= "a") || (453 == 804)) && (!(-2147483648 != 455))), F1: fi72}
                                _ = fi72
                                _ = v73
                            }
                            _ = v71
                        }
                        _ = v68
                    } else {
                        fmt.Println(v3, v1, v3)
                        switch ("foo bar" + ("hello" + "hello")) {
                        case "":
                            var v74 uint32 = 721
                            _ = v74
                        case "foo bar":
                            v75 := v4["hello"]
                            _ = v75
                        default:
                            v76 := []int{32767, -1, -9223372036854775808}
                            _ = v76
                        }
                    }
                case "a":
                    switch "hello" {
                    case "\n\t":
                        if false {
                            v77 := GenS10{F0: 231}
                            v3 = (v3 + (v3 & v3))
                            _ = v77
                        }
                    case "a":
                        v2 = float64(0.0)
                    case "foo bar":
                        fmt.Println(v2)
                    default:
                        if false {
                            if true {
                                v78 := v4["hello"]
                                switch "a" {
                                case "foo bar":
                                    for fi79 := 0; fi79 < 19; fi79++ {
                                        v4[((("hello" + "\n\t") + ("\n\t" + ("a" + ""))) + "")] = (^(v78))
                                        _ = fi79
                                    }
                                case "a":
                                    if ((("a" + "foo bar") + ("hello" + ("\n\t" + "a"))) != "a") {
                                        for fi80 := 0; fi80 < 6; fi80++ {
                                            v1 = 0
                                            _ = fi80
                                        }
                                    } else {
                                        v2 = v2
                                        v81 := []int{42, 84, 127}
                                        switch -743 {
                                        case 926:
                                            var v82 uint = 80
                                            _ = v82
                                        case 32767:
                                            fmt.Println(v78, v2)
                                        default:
                                            v83 := GenS8{F0: (v1 ^ (v1 * (^(v1))))}
                                            _ = v83
                                        }
                                        _ = v81
                                    }
                                default:
                                    if (false || (!(true && false))) {
                                        switch (("\n\t" + "foo bar") + "hello") {
                                        case "a":
                                            for fi84 := 0; fi84 < 16; fi84++ {
                                                switch (-(fi84)) {
                                                case 32768:
                                                    if ((fi84 & fi84) >= (^(fi84))) {
                                                        v85 := GenS5{F0: 1, F1: 1}
                                                        _ = v85
                                                    }
                                                case 32767:
                                                    fmt.Println(v2)
                                                default:
                                                    fmt.Println(v2, v2, v3)
                                                }
                                                if true {
                                                    v86 := v4["\n\t"]
                                                    switch "hello" {
                                                    case "":
                                                        if false {
                                                            v32 -= (^(v32))
                                                            if false {
                                                                for fi87 := 0; fi87 < 3; fi87++ {
                                                                    var v88 int8 = 0
                                                                    _ = fi87
                                                                    _ = v88
                                                                }
                                                                var v89 uint16 = 42
                                                                var v90 uint64 = 457
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                v91 := []int{32768, 2147483647, fi84, (fi84 >> ((65535) & 63))}
                                                                if len(v91) >= 2 {
                                                                    v91 = v91[1:len(v91)]
                                                                }
                                                                _ = len(v4)
                                                                _ = v91
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v92 uint8 = (v1 + v1)
                                                        _ = v92
                                                    case "foo bar":
                                                        var v93 int64 = 42
                                                        _ = v93
                                                    default:
                                                        for fi94 := 0; fi94 < 5; fi94++ {
                                                            fmt.Println(v78)
                                                            v95 := []int{(fi84 + (fi84 * (fi84 / ((42) | 1)))), fi84, 9223372036854775807, fi84}
                                                            _ = fi94
                                                            _ = v95
                                                        }
                                                    }
                                                    if false {
                                                        v32 = (^(v32))
                                                        _ = len(v4)
                                                        for fi96 := 0; fi96 < 16; fi96++ {
                                                            var v97 string = ("a" + "foo bar")
                                                            var v98 int8 = 127
                                                            fmt.Println(v3, v2, v86)
                                                            _ = fi96
                                                            _ = v97
                                                            _ = v98
                                                        }
                                                    } else {
                                                        v2 += ((v2 + ((v2 / v2) + (v2 / 1.0))) * v2)
                                                        switch fi84 {
                                                        case -2147483648:
                                                            var v99 int8 = -56
                                                            _ = v99
                                                        case 9223372036854775807:
                                                            if (42 < 32767) {
                                                                switch ("hello" + "a") {
                                                                case "hello":
                                                                    var v100 float64 = (0.0 / -2.25)
                                                                    _ = v100
                                                                case "\n\t":
                                                                    v101 := GenS10{F0: 42}
                                                                    _ = v101
                                                                default:
                                                                    for fi102 := 0; fi102 < 1; fi102++ {
                                                                        switch (fi102 - fi84) {
                                                                        case -81:
                                                                            var v103 bool = (127 <= -1)
                                                                            _ = v103
                                                                        default:
                                                                            var v104 uint16 = 2
                                                                            _ = v104
                                                                        }
                                                                        var v105 uint32 = v3
                                                                        _ = fi102
                                                                        _ = v105
                                                                    }
                                                                }
                                                                if (false && ((^(fi84)) != fi84)) {
                                                                    for fi106 := 0; fi106 < 18; fi106++ {
                                                                        var v107 int64 = -583
                                                                        _ = fi106
                                                                        _ = v107
                                                                    }
                                                                    if (!false) {
                                                                        var v108 uint16 = 1
                                                                        v109 := []int{512, fi84, 127}
                                                                        _ = v108
                                                                        _ = v109
                                                                    }
                                                                } else {
                                                                    for fi110 := 0; fi110 < 16; fi110++ {
                                                                        var v111 int = 42
                                                                        if ("" >= (("" + "foo bar") + "\n\t")) {
                                                                            var v112 int64 = 1
                                                                            var v113 float32 = ((-2.25 - ((3.14159 * 1.5) / 1.5)) + 1e10)
                                                                            var v114 uint16 = 2
                                                                            var v115 uint32 = uint32(v1)
                                                                            _ = v112
                                                                            _ = v113
                                                                            _ = v114
                                                                            _ = v115
                                                                        }
                                                                        fmt.Println(v2, v111, v2)
                                                                        v3 += (v3 * (v3 | 347))
                                                                        _ = fi110
                                                                        _ = v111
                                                                    }
                                                                    var v116 uint = 65535
                                                                    var v117 float64 = 1.5
                                                                    _ = v116
                                                                    _ = v117
                                                                }
                                                                var v118 uint = 2
                                                                _ = v118
                                                            } else {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    var v119 string = "\n\t"
                                                                    _ = v119
                                                                case "":
                                                                    var v120 uint32 = v86
                                                                    _ = v120
                                                                }
                                                                if false {
                                                                    v121 := []int{(fi84 - fi84), fi84, -628, fi84}
                                                                    v122 := GenS6{F0: v1}
                                                                    _ = v121
                                                                    _ = v122
                                                                }
                                                            }
                                                        case 386:
                                                            v123 := GenS5{F0: 823, F1: -1}
                                                            _ = v123
                                                        default:
                                                            v32 -= v32
                                                        }
                                                        for fi124 := 0; fi124 < 7; fi124++ {
                                                            v2 = v2
                                                            v125 := []int{fi84, 127, 32768, (fi84 ^ (fi124 | fi124))}
                                                            v3 -= v86
                                                            _ = fi124
                                                            _ = v125
                                                        }
                                                    }
                                                    _ = v86
                                                }
                                                _ = fi84
                                            }
                                        case "hello":
                                            for fi126 := 0; fi126 < 9; fi126++ {
                                                if false {
                                                    if (1 < -31) {
                                                        v3 = v3
                                                    } else {
                                                        if true {
                                                            var v127 int16 = v32
                                                            _ = v127
                                                        }
                                                    }
                                                    var v128 uint = 863
                                                    _ = len(v4)
                                                    _ = v128
                                                } else {
                                                    var v129 int8 = 42
                                                    var v130 uint64 = 18446744073709551615
                                                    var v131 int16 = (-(v32))
                                                    _ = v129
                                                    _ = v130
                                                    _ = v131
                                                }
                                                _ = fi126
                                            }
                                        default:
                                            v132 := GenS11{F0: 0, F1: -1}
                                            _ = v132
                                        }
                                        v133 := []int{-825, 0, 0, 1}
                                        v3 = v3
                                        _ = v133
                                    } else {
                                        var v134 uint = 1
                                        var v135 int32 = -1
                                        _ = v134
                                        _ = v135
                                    }
                                }
                                _ = v78
                            }
                            v32 -= (v32 / (((v32 & (v32 << ((65535) & 63)))) | 1))
                        } else {
                            var v136 string = "foo bar"
                            var v137 int = 0
                            _ = v136
                            _ = v137
                        }
                    }
                }
                fmt.Println(v1, v3)
            }
        } else {
            fmt.Println(v3, v2)
            v138 := GenS10{F0: 255}
            var v139 uint8 = (v1 + (v1 % (((v1 * v1)) | 1)))
            _ = v138
            _ = v139
        }
    case "foo bar":
        if (((0.0 / (-2.25 * (((1.5 + 1e10) - float32((float64(-2.25) - v2))) / (1.5 / 0.1)))) + 0.1) < (float32(v2) + 1e10)) {
            for fi140 := 0; fi140 < 13; fi140++ {
                v1 = v1
                if (10 > 100) {
                    v141 := []int{0, (fi140 / (((fi140 & fi140)) | 1)), (fi140 >> ((42) & 63)), (fi140 / (((fi140 >> ((2) & 63))) | 1))}
                    var v142 int16 = v32
                    _ = v141
                    _ = v142
                } else {
                    if true {
                        v143 := GenS12{F0: "hello", F1: 11, F2: 205}
                        _ = v143
                    }
                    v32 = v32
                    var v144 float32 = 3.14159
                    _ = v144
                }
                var v145 uint = 278
                _ = fi140
                _ = v145
            }
        }
    case "\n\t":
        fmt.Println(v32, v32)
    }
    v1 = v1
    v4["hello"] = v3
    var v146 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v146
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
