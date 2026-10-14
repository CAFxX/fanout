package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 bool; F2 int }
type GenS6 struct { F0 uint16; F1 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 int; F1 int16; F2 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint16; F1 string; F2 int64 }
type GenS13 struct { F0 int64; F1 int }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int32; F1 int64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 int8; F2 uint }
func genFn0() uint32 {
    v4 := GenS1{F0: (957 == 2), F1: ("foo bar" + ("a" + ("" + ("a" + "\n\t")))), F2: 1000}
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 string = ("" + "foo bar")
        if true {
            v6 = "hello"
        } else {
            v7 := GenS2{F0: -480, F1: -2147483648}
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(fi8, fi8, fi8)
                _ = fi8
            }
            switch "" {
            case "hello":
                v6 = (((v6 + (v6 + ("" + v6))) + "hello") + v6)
            case "\n\t":
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 bool = true
                    fmt.Println(v10, v6)
                    v6 = ("" + "")
                    _ = fi9
                    _ = v10
                }
            }
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    var v11 uint16 = 42
    _ = v11
    _ = v4
    return 555
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 127
    var v3 string = "a"
    var v12 uint = 440
    if true {
        fmt.Println(v2, v2, v3)
        if true {
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    v12 = (v12 | 255)
                    var v14 int32 = 463
                    _ = v14
                }
                v15 := GenS1{F0: (!false), F1: "\n\t", F2: 445}
                if true {
                    var v16 int64 = -1
                    fmt.Println(v16)
                    _ = v16
                } else {
                    if false {
                        if ("" < ("a" + ("hello" + "hello"))) {
                            v17 := GenS2{F0: 0, F1: 2147483647}
                            for fi18 := 0; fi18 < 2; fi18++ {
                                var v19 float32 = 1e10
                                var v20 bool = false
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            _ = v17
                        }
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v22 := GenS1{F0: false, F1: v3, F2: 740}
                            v12 += v12
                            _ = fi21
                            _ = v22
                        }
                        if (!(!(!(42 == 1)))) {
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    v3 = "a"
                                    v23 := GenS2{F0: 1, F1: 2147483647}
                                    _ = v23
                                } else {
                                    if (65535 != 2147483648) {
                                        v24 := []int{}
                                        _ = v24
                                    } else {
                                        var v25 int8 = 127
                                        if true {
                                            if true {
                                                fmt.Println(fi13, v3)
                                            }
                                        } else {
                                            fmt.Println(fi13, v3)
                                        }
                                        _ = v25
                                    }
                                    var v26 string = ("\n\t" + v3)
                                    v27 := []int{fi13}
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                    }
                    var v28 int8 = int8(v12)
                    switch "\n\t" {
                    case "foo bar":
                        if ((true || (!false)) && (-865 == -1)) {
                            v12 += (v12 & (v12 * v12))
                        }
                    case "":
                        fmt.Println(fi13, fi13, v12)
                    case "hello":
                        v28 += 127
                    default:
                        v29 := GenS3{F0: 661, F1: 0, F2: 613}
                        _ = v29
                    }
                    _ = v28
                }
                _ = fi13
                _ = v15
            }
        } else {
            switch v3 {
            case "a":
                for fi30 := 0; fi30 < 3; fi30++ {
                    if false {
                        var v31 int = (fi30 & 898)
                        v3 = "foo bar"
                        _ = v31
                    }
                    switch 42 {
                    case 4294967295:
                        fmt.Println(fi30, v3, v12)
                    case 0:
                        if (!(0.0 < 0.0)) {
                            if true {
                                v32 := GenS1{F0: (!(true && true)), F1: "\n\t", F2: int16(fi30)}
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    v3 = "\n\t"
                                    _ = fi33
                                }
                                _ = v32
                            }
                            var v34 float32 = float32(1.5)
                            if (((fi30 ^ (fi30 & fi30)) == fi30) || false) {
                                v2 = (v2 | (^(v2)))
                                if true {
                                    v34 = v34
                                    v34 -= (3.14159 - (float32((-2.25 - 3.14159)) + v34))
                                }
                            } else {
                                if true {
                                    var v35 uint32 = genFn0()
                                    _ = v35
                                } else {
                                    v34 += v34
                                }
                                if false {
                                    v36 := GenS4{F0: 79}
                                    var v37 float32 = v34
                                    if true {
                                        if (float64(v37) != 0.1) {
                                            if (false || (42 > 32767)) {
                                                switch -81 {
                                                case 42:
                                                    v3 = v3
                                                case -1:
                                                    if false {
                                                        for fi38 := 0; fi38 < 6; fi38++ {
                                                            v34 += v34
                                                            v39 := GenS2{F0: 32767, F1: -2147483648}
                                                            _ = fi38
                                                            _ = v39
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v3, v37, v2)
                                                }
                                                v40 := GenS1{F0: false, F1: v3, F2: 281}
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if (!false) {
                                                        fmt.Println(v12)
                                                        for fi41 := 0; fi41 < 4; fi41++ {
                                                            var v42 bool = (1 != fi30)
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        v34 = v34
                                                    }
                                                }
                                                _ = v40
                                            }
                                            for fi43 := 0; fi43 < 1; fi43++ {
                                                if (false && false) {
                                                    var v44 int = fi30
                                                    _ = v44
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        if true {
                                                            v45 := GenS3{F0: 42, F1: 350, F2: 0}
                                                            _ = v45
                                                        }
                                                        v12 = 65535
                                                        v46 := v1["a"]
                                                        var v47 uint64 = 42
                                                        fmt.Println(v37, v2)
                                                        _ = v46
                                                        _ = v47
                                                    case "hello":
                                                        v34 -= v34
                                                    case "":
                                                        switch (fi43 & (-(fi30))) {
                                                        case 2147483647:
                                                            var v48 uint16 = 0
                                                            _ = v48
                                                        }
                                                        if true {
                                                            if true {
                                                                v49 := []int{(fi30 % (((fi43 ^ fi43)) | 1)), 561, (fi43 ^ fi30), fi30}
                                                                _ = v49
                                                            }
                                                            var v50 float32 = (v37 / 1.5)
                                                            v12 -= (v12 * (v12 >> ((0) & 63)))
                                                            _ = v50
                                                        } else {
                                                            v51 := GenS5{F0: 943, F1: true, F2: 4294967295}
                                                            fmt.Println(v2)
                                                            v52 := GenS2{F0: -279, F1: -767}
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        if true {
                                                            v12 = v12
                                                        } else {
                                                            var v53 int = 531
                                                            v53 += (fi30 ^ fi43)
                                                            v37 = v37
                                                            v34 -= (((0.0 - v37) + (v37 / (3.14159 - 1.5))) / -2.25)
                                                            if (255 != 1) {
                                                                if (((fi43 + v53) >= (fi30 & fi30)) || (1 < -1)) {
                                                                    v54 := GenS6{F0: 65535, F1: -566}
                                                                    switch fi30 {
                                                                    case 0:
                                                                        var v55 uint16 = 580
                                                                        _ = v55
                                                                    }
                                                                    if (false && true) {
                                                                        switch fi43 {
                                                                        case 366:
                                                                            var v56 int32 = 42
                                                                            _ = v56
                                                                        case -1:
                                                                            var v57 bool = false
                                                                            _ = v57
                                                                        }
                                                                        var v58 uint = 4294967295
                                                                        v59 := GenS3{F0: 42, F1: 32767, F2: 255}
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                    v2 += v2
                                                                    v60 := GenS7{F0: 4294967295}
                                                                    _ = v54
                                                                    _ = v60
                                                                } else {
                                                                    v34 += (float32(-2.25) - (0.0 * 1.5))
                                                                    fmt.Println(v2, v3)
                                                                    v61 := GenS8{F0: (false || true), F1: 0}
                                                                    v62 := []int{295}
                                                                    if false {
                                                                        var v63 int8 = -128
                                                                        var v64 uint = v12
                                                                        v65 := GenS9{F0: 615, F1: 931}
                                                                        var v66 uint8 = v2
                                                                        var v67 int8 = v63
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                var v68 int32 = -2147483648
                                                                if true {
                                                                    var v69 int64 = 0
                                                                    _ = v69
                                                                } else {
                                                                    var v70 float64 = ((1.5 - ((1e10 / 1e10) + (((3.14159 - (1e10 / 1.0)) - 0.0) * ((float64(0.1) - (((0.1 + float64(1e10)) / (1e10 / 1.0)) + ((float64((v37 / (v34 * ((1.5 / v37) * v37)))) - 1e10) - 0.1))) + ((1.5 + 3.14159) + 1e10))))) - float64((v34 - (v34 * 3.14159))))
                                                                    _ = v70
                                                                }
                                                                _ = v68
                                                            } else {
                                                                v71 := GenS9{F0: (v53 << ((v12) & 63)), F1: 1}
                                                                _ = v71
                                                            }
                                                            _ = len(v1)
                                                            _ = v53
                                                        }
                                                    default:
                                                        switch fi30 {
                                                        case -712:
                                                            switch (^(fi43)) {
                                                            case 4294967295:
                                                                v72 := GenS10{F0: fi43, F1: 1, F2: 409}
                                                                _ = v72
                                                            }
                                                        case 648:
                                                            var v73 float32 = 3.14159
                                                            _ = v73
                                                        }
                                                    }
                                                    var v74 float32 = v34
                                                    _ = v74
                                                }
                                                var v75 float64 = ((3.14159 + 1.5) - 0.0)
                                                _ = fi43
                                                _ = v75
                                            }
                                        }
                                        var v76 int32 = 42
                                        for fi77 := 0; fi77 < 7; fi77++ {
                                            v78 := []int{}
                                            var v79 int8 = -1
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                        _ = v76
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                if (2147483647 != -9223372036854775808) {
                                    var v80 int64 = 827
                                    var v81 int32 = -2147483648
                                    for fi82 := 0; fi82 < 8; fi82++ {
                                        var v83 int32 = v81
                                        v84 := []int{fi82, (fi82 - fi82), (fi82 | 32767)}
                                        _ = fi82
                                        _ = v83
                                        _ = v84
                                    }
                                    _ = v80
                                    _ = v81
                                } else {
                                    if true {
                                        var v85 uint = 4294967295
                                        _ = v85
                                    } else {
                                        v86 := GenS11{F0: 1}
                                        _ = v86
                                    }
                                }
                            }
                            _ = v34
                        }
                    default:
                        fmt.Println(fi30)
                    }
                    _ = fi30
                }
            case "hello":
                if true {
                    if (32768 != 32768) {
                        switch -219 {
                        case 0:
                            for fi87 := 0; fi87 < 3; fi87++ {
                                for fi88 := 0; fi88 < 4; fi88++ {
                                    v1[""] = 65535
                                    v3 = ("hello" + "hello")
                                    _ = fi88
                                }
                                v2 = (v2 % ((v2) | 1))
                                v2 = 189
                                _ = fi87
                            }
                        case 2147483647:
                            for fi89 := 0; fi89 < 4; fi89++ {
                                var v90 float32 = (-2.25 + (1.5 + (1e10 - 0.0)))
                                _ = fi89
                                _ = v90
                            }
                        default:
                            v91 := GenS12{F0: 255, F1: (v3 + ""), F2: -1}
                            _ = v91
                        }
                        switch -9223372036854775808 {
                        case 716:
                            for fi92 := 0; fi92 < 12; fi92++ {
                                v93 := []int{}
                                v2 = (v2 ^ (v2 + (v2 ^ 128)))
                                _ = fi92
                                _ = v93
                            }
                        case 42:
                            v94 := GenS1{F0: true, F1: "a", F2: 1000}
                            _ = v94
                        default:
                            var v95 uint8 = 1
                            _ = v95
                        }
                        v96 := GenS6{F0: 2, F1: -159}
                        _ = v96
                    } else {
                        var v97 float32 = 1e10
                        v98 := GenS13{F0: -332, F1: 9223372036854775807}
                        for fi99 := 0; fi99 < 13; fi99++ {
                            var v100 uint8 = (v2 * v2)
                            _ = fi99
                            _ = v100
                        }
                        _ = v97
                        _ = v98
                    }
                } else {
                    for fi101 := 0; fi101 < 12; fi101++ {
                        v3 = v3
                        _ = fi101
                    }
                    for fi102 := 0; fi102 < 3; fi102++ {
                        v103 := []int{fi102, (fi102 % ((fi102) | 1)), fi102}
                        if (true && (0.0 >= 1.5)) {
                            if true {
                                v3 = "foo bar"
                            }
                            for fi104 := 0; fi104 < 20; fi104++ {
                                switch int(v2) {
                                case 2147483647:
                                    var v105 uint32 = uint32(v2)
                                    _ = v105
                                case 176:
                                    fmt.Println(fi102, fi104, fi104)
                                case -2147483648:
                                    v12 -= (v12 * 435)
                                }
                                _ = fi104
                            }
                            if false {
                                if true {
                                    v3 = "hello"
                                } else {
                                    fmt.Println(v3, fi102)
                                    for fi106 := 0; fi106 < 2; fi106++ {
                                        var v107 int8 = 0
                                        var v108 uint16 = 32767
                                        fmt.Println(fi102)
                                        _ = fi106
                                        _ = v107
                                        _ = v108
                                    }
                                    switch fi102 {
                                    case 9223372036854775807:
                                        switch (^(fi102)) {
                                        case -9223372036854775808:
                                            v109 := GenS14{F0: 454}
                                            _ = v109
                                        case -677:
                                            fmt.Println(fi102)
                                        }
                                    case 127:
                                        var v110 uint = 18446744073709551615
                                        _ = v110
                                    case -1:
                                        v111 := GenS1{F0: (!false), F1: (("hello" + (("\n\t" + v3) + v3)) + (("a" + v3) + v3)), F2: -256}
                                        _ = v111
                                    }
                                }
                                fmt.Println(v2)
                                var v112 int32 = 0
                                _ = v112
                            }
                        }
                        v113 := GenS13{F0: 721, F1: 128}
                        _ = fi102
                        _ = v103
                        _ = v113
                    }
                }
            default:
                if (553 == 915) {
                    v114 := GenS2{F0: -458, F1: 1}
                    var v115 int = 567
                    _ = v114
                    _ = v115
                }
            }
        }
        if (451 == 303) {
            v116 := GenS2{F0: 42, F1: -2147483648}
            v2 = v2
            v12 -= (v12 | 0)
            _ = v116
        }
    }
    if false {
        var v117 int = -759
        var v118 int8 = 35
        _ = v117
        _ = v118
    } else {
        v12 += v12
    }
    if ((0.1 - float32(1.5)) < (3.14159 + 0.1)) {
        fmt.Println(v2, v2)
        var v119 int16 = -32768
        var v120 float32 = float32((-2.25 - 1.5))
        _ = v119
        _ = v120
    }
    if false {
        for fi121 := 0; fi121 < 3; fi121++ {
            var v122 uint8 = 138
            fmt.Println(v122, fi121, v12)
            var v123 int16 = 32767
            _ = fi121
            _ = v122
            _ = v123
        }
        _ = len(v1)
        switch "" {
        case "a":
            if ((!(!(!true))) || false) {
                if (1 <= 42) {
                    _ = len(v1)
                    var v124 uint32 = genFn0()
                    _ = v124
                } else {
                    var v125 int32 = -1
                    for fi126 := 0; fi126 < 2; fi126++ {
                        fmt.Println(v3)
                        var v127 int64 = 42
                        v128 := GenS4{F0: 119}
                        _ = fi126
                        _ = v127
                        _ = v128
                    }
                    _ = v125
                }
                v129 := []int{}
                _ = v129
            }
        }
    }
    for fi130 := 0; fi130 < 19; fi130++ {
        fmt.Println(fi130)
        switch fi130 {
        case 127:
            fmt.Println(v12)
        case -623:
            v3 = v3
        case -2147483648:
            v3 = v3
        default:
            v131 := GenS15{F0: 0, F1: -9223372036854775808}
            _ = v131
        }
        _ = fi130
    }
    if true {
        var v132 uint16 = 732
        _ = v132
    } else {
        if (!false) {
            v3 = v3
            fmt.Println(v3, v3, v12)
        }
        if true {
            v133 := []int{760}
            switch 816 {
            case 0:
                fmt.Println(v2)
            case 32768:
                switch ("foo bar" + v3) {
                case "a":
                    if ((v12 << (((v12 << (((v12 | uint(v2))) & 63))) & 63)) < (v12 / ((v12) | 1))) {
                        v12 -= 42
                        _ = len(v133)
                        _ = len(v1)
                    } else {
                        v12 = 18446744073709551615
                        for fi134 := 0; fi134 < 11; fi134++ {
                            if (!false) {
                                switch v3 {
                                case "hello":
                                    v3 = v3
                                case "":
                                    fmt.Println(v2, v3, fi134)
                                default:
                                    var v135 string = "hello"
                                    _ = v135
                                }
                            } else {
                                var v136 int8 = 100
                                switch fi134 {
                                case 42:
                                    var v137 uint16 = 2
                                    _ = v137
                                default:
                                    v138 := GenS16{F0: 42}
                                    _ = v138
                                }
                                _ = v136
                            }
                            _ = fi134
                        }
                    }
                case "foo bar":
                    switch -2147483648 {
                    case 1:
                        v139 := v1[(v3 + "hello")]
                        _ = v139
                    case -1:
                        for fi140 := 0; fi140 < 2; fi140++ {
                            switch v3 {
                            case "a":
                                var v141 int16 = 403
                                _ = v141
                            }
                            if len(v133) > 0 {
                                v133[0] = (fi140 | (fi140 % (((fi140 + (-(fi140)))) | 1)))
                            }
                            _ = fi140
                        }
                    case 4294967295:
                        var v142 float32 = 3.14159
                        _ = v142
                    default:
                        v2 -= 2
                    }
                }
            case 127:
                var v143 uint8 = v2
                _ = v143
            }
            _ = v133
        } else {
            v144 := []int{9223372036854775807, 2147483647}
            _ = v144
        }
    }
    v145 := []int{}
    v12 += v12
    v146 := GenS17{F0: 1, F1: 127, F2: (v12 | v12)}
    var v147 int16 = 0
    v12 += 48
    for fi148 := 0; fi148 < 10; fi148++ {
        if len(v145) > 0 {
            v145[0] = -1
        }
        _ = fi148
    }
    var v149 float64 = (1.5 + 1e10)
    for fi150 := 0; fi150 < 1; fi150++ {
        v151 := GenS6{F0: 32767, F1: 9223372036854775807}
        var v152 uint64 = 18446744073709551615
        _ = fi150
        _ = v151
        _ = v152
    }
    v3 = v3
    _ = len(v1)
    switch v3 {
    case "hello":
        v147 += -107
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v145:", len(v145), v145)
    fmt.Println("v147:", v147)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v145
    _ = v146
    _ = v147
    _ = v149
    _ = v2
    _ = v3
}
