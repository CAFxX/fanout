package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int8 }
type GenS2 struct { F0 int8; F1 uint64; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint; F1 int64; F2 uint64 }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint16; F1 uint; F2 int8 }
type GenS9 struct { F0 int64; F1 uint; F2 uint8 }
type GenS10 struct { F0 int; F1 int64; F2 uint64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 uint8; F1 uint16; F2 int }
type GenS13 struct { F0 uint8; F1 bool }
type GenS14 struct { F0 string; F1 uint16; F2 int64 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint64; F1 string; F2 uint8 }
type GenS18 struct { F0 bool; F1 int32; F2 int16 }
type GenS19 struct { F0 uint; F1 uint; F2 uint32 }
type GenS20 struct { F0 int }

func main() {
    var v1 uint16 = 1
    var v2 float64 = (-2.25 + float64(1.5))
    var v3 uint32 = 42
    var v4 int32 = 0
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            v2 = v2
            for fi6 := 0; fi6 < 12; fi6++ {
                v3 -= v3
                var v7 int32 = (v4 | v4)
                var v8 int64 = 785
                fmt.Println(v1, v8)
                var v9 int = fi6
                _ = fi6
                _ = v7
                _ = v8
                _ = v9
            }
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                if true {
                    var v11 uint16 = v1
                    if (v4 <= v4) {
                        var v12 int32 = v4
                        v13 := []int{fi10, -9223372036854775808, fi5}
                        switch ((("foo bar" + "hello") + "") + ("a" + "")) {
                        case "foo bar":
                            v3 += (v3 | v3)
                        default:
                            if false {
                                switch "a" {
                                case "hello":
                                    v12 = v12
                                default:
                                    switch fi5 {
                                    case 9223372036854775807:
                                        switch (fi5 | fi5) {
                                        case 55:
                                            var v14 int8 = 0
                                            _ = v14
                                        }
                                    case 642:
                                        v15 := GenS1{F0: 0, F1: v12, F2: -1}
                                        _ = v15
                                    default:
                                        if (v2 <= float64(((0.0 / 1.5) * (0.0 * 0.1)))) {
                                            if false {
                                                switch ("foo bar" + "a") {
                                                case "\n\t":
                                                    _ = len(v13)
                                                }
                                                var v16 uint64 = 4294967295
                                                switch 0 {
                                                case 0:
                                                    if (42 > 42) {
                                                        var v17 int64 = 42
                                                        _ = v17
                                                    } else {
                                                        fmt.Println(fi10, v4, v16)
                                                    }
                                                default:
                                                    v18 := GenS2{F0: 120, F1: v16, F2: false}
                                                    _ = v18
                                                }
                                                _ = v16
                                            }
                                            var v19 uint = 255
                                            if (0.1 < 1e10) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if len(v13) > 0 {
                                                        v13[0] = 128
                                                    }
                                                }
                                                switch (^(fi10)) {
                                                case 1:
                                                    if (((fi10 >= 32767) && ((1e10 - (0.0 + 3.14159)) >= (0.1 + 1.5))) && false) {
                                                        v11 = 0
                                                        if false {
                                                            fmt.Println(v3, v12)
                                                            if true {
                                                                var v20 uint64 = 1
                                                                if len(v13) >= 2 {
                                                                    v13 = v13[1:len(v13)]
                                                                }
                                                                v4 = (v4 | 42)
                                                                _ = v20
                                                            } else {
                                                                v2 -= 3.14159
                                                            }
                                                        }
                                                        switch fi10 {
                                                        case -1:
                                                            v21 := GenS3{F0: -9223372036854775808}
                                                            _ = v21
                                                        default:
                                                            v1 = v1
                                                            v4 = (v4 + v4)
                                                        }
                                                        _ = len(v13)
                                                    } else {
                                                        v2 -= v2
                                                        var v22 uint8 = 105
                                                        v23 := GenS3{F0: 1}
                                                        switch -1 {
                                                        case 42:
                                                            var v24 uint32 = v3
                                                            _ = v24
                                                        case 128:
                                                            if (42 < 1) {
                                                                v25 := GenS3{F0: 9223372036854775807}
                                                                _ = v25
                                                            }
                                                        }
                                                        if (false || true) {
                                                            var v26 uint = (v19 % ((42) | 1))
                                                            var v27 uint8 = 171
                                                            _ = v26
                                                            _ = v27
                                                        } else {
                                                            v19 = (v19 >> (((v19 / ((v19) | 1))) & 63))
                                                            var v28 int32 = v12
                                                            _ = v28
                                                        }
                                                        if len(v13) > 0 {
                                                            v13[0] = fi5
                                                        }
                                                        v29 := GenS2{F0: 42, F1: 4294967295, F2: true}
                                                        if (32767 >= -215) {
                                                            _ = len(v13)
                                                            v19 += v19
                                                            if (!true) {
                                                                v1 = 255
                                                                if (v19 > v19) {
                                                                    if len(v13) > 0 {
                                                                        v13[0] = (fi5 & fi5)
                                                                    }
                                                                }
                                                                var v30 bool = true
                                                                _ = v30
                                                            }
                                                        }
                                                        _ = v22
                                                        _ = v23
                                                        _ = v29
                                                    }
                                                case 9223372036854775807:
                                                    var v31 float32 = ((1e10 - ((((3.14159 / (((float32(float64(0.1)) + 1.5) * 0.1) + 3.14159)) - 1.5) / (-2.25 / (1.5 - ((1e10 * (0.0 * 1e10)) * 3.14159)))) / 1.5)) - 3.14159)
                                                    _ = v31
                                                case 32768:
                                                    fmt.Println(fi10, v12)
                                                default:
                                                    if false {
                                                        if (true || (v1 >= (v1 / (((v11 / (((v1 / (((v1 | 1)) | 1))) | 1))) | 1)))) {
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                v3 -= uint32(v19)
                                                            }
                                                        }
                                                        if true {
                                                            v12 = v4
                                                            var v32 int = int(v19)
                                                            _ = v32
                                                        } else {
                                                            var v33 uint64 = 2
                                                            v2 = v2
                                                            _ = v33
                                                        }
                                                        var v34 float32 = 0.0
                                                        _ = v34
                                                    }
                                                }
                                                fmt.Println(fi10)
                                            }
                                            _ = v19
                                        } else {
                                            switch (fi5 ^ (^(fi5))) {
                                            case 32768:
                                                v3 += (v3 << ((4294967295) & 63))
                                            case 127:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v35 int32 = (v12 ^ v12)
                                                    _ = v35
                                                case "a":
                                                    v36 := GenS2{F0: 98, F1: 18446744073709551615, F2: ((2147483647 >= fi10) || (-185 == 42))}
                                                    _ = v36
                                                default:
                                                    v1 = (v11 << ((1) & 63))
                                                }
                                            case 0:
                                                v37 := GenS1{F0: 255, F1: int32(v3), F2: 127}
                                                _ = v37
                                            default:
                                                var v38 uint64 = 42
                                                v39 := GenS3{F0: -313}
                                                var v40 uint = 2
                                                v41 := GenS1{F0: (v40 % (((v40 << ((v40) & 63))) | 1)), F1: v12, F2: -18}
                                                var v42 uint32 = (^(v3))
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                                _ = v42
                                            }
                                            v3 = (v3 << ((364) & 63))
                                            if (0 != 651) {
                                                var v43 int64 = int64(v4)
                                                v13 = append(v13, fi10)
                                                v44 := GenS4{F0: 2, F1: 2}
                                                if len(v13) > 0 {
                                                    v13[0] = -353
                                                }
                                                _ = v43
                                                _ = v44
                                            }
                                            var v45 uint16 = v1
                                            _ = v45
                                        }
                                        if (!true) {
                                            var v46 uint32 = v3
                                            if false {
                                                if true {
                                                    var v47 int = fi5
                                                    v1 = 255
                                                    _ = v47
                                                }
                                                switch fi5 {
                                                case 836:
                                                    v13 = append(v13, (-(fi5)))
                                                default:
                                                    v48 := GenS5{F0: -255, F1: 42}
                                                    _ = v48
                                                }
                                                _ = len(v13)
                                            }
                                            if (v3 < v46) {
                                                if false {
                                                    switch "a" {
                                                    case "a":
                                                        v46 += 2147483647
                                                    case "hello":
                                                        if ("\n\t" > (("hello" + "\n\t") + ("" + "a"))) {
                                                            var v49 string = ("hello" + ("a" + "hello"))
                                                            _ = v49
                                                        }
                                                    }
                                                    switch (fi10 << ((255) & 63)) {
                                                    case 128:
                                                        v2 += v2
                                                    case 32768:
                                                        if len(v13) > 0 {
                                                            v13[0] = (fi5 * fi10)
                                                        }
                                                    }
                                                    if false {
                                                        fmt.Println(v46, v11, v11)
                                                    } else {
                                                        var v50 uint32 = 654
                                                        fmt.Println(v2)
                                                        if len(v13) >= 2 {
                                                            v13 = v13[1:len(v13)]
                                                        }
                                                        _ = v50
                                                    }
                                                } else {
                                                    fmt.Println(v4, v46, v11)
                                                    if false {
                                                        v51 := GenS1{F0: 538, F1: v12, F2: 42}
                                                        _ = v51
                                                    }
                                                    var v52 int16 = -32768
                                                    _ = v52
                                                }
                                                v13 = append(v13, fi10)
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        switch (-(fi5)) {
                                                        case 4294967295:
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        }
                                                    }
                                                    v13 = append(v13, fi10)
                                                    var v53 int64 = -9223372036854775808
                                                    _ = v53
                                                }
                                            }
                                            switch "" {
                                            case "a":
                                                switch (fi10 / ((32768) | 1)) {
                                                case -1:
                                                    var v54 float64 = (((v2 + -2.25) * float64(((0.0 + 0.0) - 0.0))) - (float64(3.14159) + 0.1))
                                                    _ = v54
                                                }
                                            default:
                                                _ = len(v13)
                                            }
                                            v55 := GenS6{F0: 405, F1: -240, F2: 18446744073709551615}
                                            _ = v46
                                            _ = v55
                                        }
                                        v2 = float64(-2.25)
                                    }
                                }
                                if len(v13) > 0 {
                                    v13[0] = (fi5 | fi10)
                                }
                            } else {
                                if len(v13) > 0 {
                                    v13[0] = (fi5 * fi5)
                                }
                                v12 -= v12
                            }
                        }
                        if (!(42 <= 127)) {
                            _ = len(v13)
                        } else {
                            v56 := GenS7{F0: "", F1: v1, F2: 2}
                            var v57 float64 = -2.25
                            fmt.Println(v4)
                            _ = v56
                            _ = v57
                        }
                        _ = v12
                        _ = v13
                    }
                    switch "hello" {
                    case "":
                        if (!((fi5 & -880) >= fi10)) {
                            if ((true || (!true)) || true) {
                                v58 := GenS5{F0: -1, F1: 0}
                                var v59 int = (-(fi5))
                                _ = v58
                                _ = v59
                            } else {
                                if (127 <= fi5) {
                                    var v60 int16 = 42
                                    if true {
                                        v61 := GenS1{F0: 255, F1: (v4 - (v4 + (v4 ^ 42))), F2: 0}
                                        var v62 uint8 = 255
                                        _ = v61
                                        _ = v62
                                    } else {
                                        v63 := GenS4{F0: 255, F1: 128}
                                        _ = v63
                                    }
                                    var v64 uint8 = 128
                                    _ = v60
                                    _ = v64
                                } else {
                                    v65 := []int{(fi5 ^ fi5), (-(fi10))}
                                    if ((-128 > 1) && true) {
                                        var v66 uint = 42
                                        switch (("foo bar" + "\n\t") + ("\n\t" + "foo bar")) {
                                        case "foo bar":
                                            var v67 uint32 = (v3 >> ((v66) & 63))
                                            _ = v67
                                        default:
                                            v4 = v4
                                        }
                                        v68 := GenS7{F0: "a", F1: (v1 % ((v1) | 1)), F2: 0}
                                        v69 := GenS8{F0: 215, F1: 4294967295, F2: 100}
                                        _ = v66
                                        _ = v68
                                        _ = v69
                                    }
                                    var v70 string = "a"
                                    v70 = "foo bar"
                                    _ = v65
                                    _ = v70
                                }
                                v71 := GenS9{F0: 42, F1: 255, F2: 206}
                                _ = v71
                            }
                            fmt.Println(fi10, v2, v1)
                            v72 := GenS8{F0: (v1 & 255), F1: 2, F2: -128}
                            if true {
                                if (false || false) {
                                    if (!(42 > 708)) {
                                        v3 = (v3 & 2)
                                        var v73 string = "foo bar"
                                        _ = v73
                                    } else {
                                        v74 := GenS4{F0: 0, F1: 1}
                                        var v75 bool = (false || false)
                                        v3 = 1
                                        _ = v74
                                        _ = v75
                                    }
                                    v76 := []int{}
                                    switch "foo bar" {
                                    case "foo bar":
                                        if len(v76) > 0 {
                                            v76[0] = (fi5 << ((42) & 63))
                                        }
                                    case "\n\t":
                                        var v77 int = fi5
                                        _ = v77
                                    default:
                                        v78 := GenS2{F0: -88, F1: 1, F2: (("a" == "") && false)}
                                        _ = v78
                                    }
                                    _ = v76
                                }
                            }
                            var v79 float32 = (1e10 + 0.1)
                            _ = v72
                            _ = v79
                        } else {
                            v80 := []int{(fi10 - (fi5 / (((fi10 ^ fi10)) | 1)))}
                            var v81 int32 = (v4 / ((v4) | 1))
                            var v82 uint = 784
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                    case "a":
                        if false {
                            v3 = v3
                            fmt.Println(v2, fi10)
                            v2 = v2
                            if (true && false) {
                                fmt.Println(v1, v4, v2)
                                var v83 int8 = 0
                                _ = v83
                            } else {
                                v84 := GenS7{F0: "\n\t", F1: 32767, F2: 9}
                                var v85 int32 = (-(v4))
                                if true {
                                    var v86 bool = true
                                    var v87 bool = v86
                                    _ = v86
                                    _ = v87
                                }
                                _ = v84
                                _ = v85
                            }
                        } else {
                            switch (^(fi5)) {
                            case 1:
                                v88 := []int{127}
                                _ = v88
                            }
                            switch "foo bar" {
                            case "foo bar":
                                var v89 uint = 561
                                _ = v89
                            default:
                                v3 += (v3 - v3)
                            }
                        }
                    default:
                        if (false || ((fi10 ^ (fi10 | 2147483647)) <= (fi5 >> ((225) & 63)))) {
                            if false {
                                v90 := []int{}
                                switch (fi5 * fi10) {
                                case 843:
                                    v90 = append(v90, (-(fi10)))
                                case 498:
                                    if (true && (!false)) {
                                        if (0.0 == float32(float64(3.14159))) {
                                            v1 -= v11
                                            switch fi10 {
                                            case 0:
                                                if (true || false) {
                                                    if (0 >= 1) {
                                                        var v91 bool = (217 > v1)
                                                        v1 = (v11 * v1)
                                                        _ = v91
                                                    } else {
                                                        v90 = append(v90, fi5)
                                                        if (false && false) {
                                                            v3 -= 0
                                                        }
                                                        v2 = (v2 * float64((((1.5 / 3.14159) * 3.14159) / (-2.25 - ((((1e10 * 0.1) * (3.14159 * ((-2.25 * 1e10) - (0.0 - float32(-2.25))))) - (((0.0 + (1.5 * (0.1 + float32(v2)))) / -2.25) + (float32((v2 / (v2 + ((0.1 * v2) + float64(((0.1 + 1e10) + 0.1)))))) + (0.0 + 1.5)))) * 1e10)))))
                                                        v92 := GenS1{F0: 255, F1: v4, F2: 1}
                                                        v93 := GenS10{F0: (^(fi10)), F1: 42, F2: 42}
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    v94 := GenS1{F0: 538, F1: 745, F2: 127}
                                                    if len(v90) > 0 {
                                                        v90[0] = fi5
                                                    }
                                                    v2 += float64((0.0 * (0.1 * (-2.25 + (1e10 - -2.25)))))
                                                    _ = v94
                                                }
                                            }
                                        } else {
                                            v95 := GenS11{F0: "", F1: "hello"}
                                            _ = v95
                                        }
                                        switch ("hello" + "hello") {
                                        case "foo bar":
                                            if (v2 == (-2.25 / v2)) {
                                                if len(v90) > 0 {
                                                    v90[0] = fi10
                                                }
                                                var v96 uint8 = 2
                                                var v97 int8 = int8(v96)
                                                _ = v96
                                                _ = v97
                                            }
                                        default:
                                            _ = len(v90)
                                        }
                                        v2 = float64(-2.25)
                                    } else {
                                        var v98 uint64 = 42
                                        if len(v90) >= 2 {
                                            v90 = v90[1:len(v90)]
                                        }
                                        var v99 uint16 = v11
                                        switch "hello" {
                                        case "foo bar":
                                            switch (("a" + ("foo bar" + ("foo bar" + "foo bar"))) + "hello") {
                                            case "hello":
                                                v100 := GenS12{F0: 119, F1: (v99 << ((0) & 63)), F2: 127}
                                                _ = v100
                                            case "a":
                                                fmt.Println(v98, v2, fi5)
                                            default:
                                                var v101 int = fi5
                                                _ = v101
                                            }
                                        case "a":
                                            fmt.Println(v99, v3, v99)
                                        case "":
                                            switch "a" {
                                            case "hello":
                                                v102 := GenS13{F0: uint8(v98), F1: true}
                                                _ = v102
                                            case "foo bar":
                                                v3 = (^(v3))
                                            default:
                                                var v103 uint = 4294967295
                                                _ = v103
                                            }
                                        }
                                        _ = v98
                                        _ = v99
                                    }
                                }
                                _ = v90
                            } else {
                                v104 := GenS8{F0: v11, F1: 4294967295, F2: -93}
                                _ = v104
                            }
                            v105 := GenS2{F0: -1, F1: 2, F2: false}
                            v11 = 442
                            fmt.Println(v1, v2)
                            _ = v105
                        }
                    }
                    var v106 float32 = 0.0
                    switch "hello" {
                    case "":
                        fmt.Println(v3)
                    case "a":
                        var v107 bool = ((-9223372036854775808 < 400) && (((true || ((false && true) && true)) || true) && true))
                        _ = v107
                    }
                    _ = v106
                    _ = v11
                } else {
                    v108 := GenS12{F0: 0, F1: (v1 & v1), F2: 0}
                    var v109 uint = 18446744073709551615
                    _ = v108
                    _ = v109
                }
                _ = fi10
            }
        }
        var v110 float64 = v2
        if false {
            v1 -= v1
            var v111 int8 = 100
            if true {
                var v112 string = "hello"
                _ = v112
            } else {
                v113 := []int{(fi5 ^ (fi5 >> ((345) & 63)))}
                for fi114 := 0; fi114 < 16; fi114++ {
                    v115 := GenS14{F0: "hello", F1: v1, F2: -1}
                    _ = fi114
                    _ = v115
                }
                fmt.Println(v3, v1)
                _ = v113
            }
            _ = v111
        }
        _ = fi5
        _ = v110
    }
    v2 -= v2
    if (!false) {
        v3 = (v3 + (^(v3)))
        v116 := GenS7{F0: "a", F1: (v1 << ((18446744073709551615) & 63)), F2: 127}
        v3 -= 65535
        _ = v116
    } else {
        if (!false) {
            switch ("foo bar" + ("a" + "hello")) {
            case "a":
                switch 981 {
                case -413:
                    switch 1 {
                    case 0:
                        if (v4 <= v4) {
                            switch 2147483647 {
                            case 791:
                                v117 := GenS15{F0: "foo bar"}
                                _ = v117
                            case 32767:
                                v3 -= (^(v3))
                            default:
                                for fi118 := 0; fi118 < 12; fi118++ {
                                    v2 = 3.14159
                                    for fi119 := 0; fi119 < 10; fi119++ {
                                        if (!((false || (int16(v1) > -32768)) || true)) {
                                            v120 := GenS4{F0: uint8(v1), F1: 42}
                                            v4 = (v4 >> ((166) & 63))
                                            _ = v120
                                        } else {
                                            v121 := []int{}
                                            var v122 int64 = 42
                                            _ = v121
                                            _ = v122
                                        }
                                        var v123 string = "a"
                                        _ = fi119
                                        _ = v123
                                    }
                                    _ = fi118
                                }
                            }
                        } else {
                            var v124 string = "a"
                            _ = v124
                        }
                    default:
                        v2 = 1e10
                    }
                case 42:
                    var v125 uint = 1
                    _ = v125
                case 32768:
                    for fi126 := 0; fi126 < 4; fi126++ {
                        switch "a" {
                        case "hello":
                            if (4294967295 >= 255) {
                                if true {
                                    v127 := []int{(fi126 * -9223372036854775808), 730, (fi126 + 127)}
                                    v128 := GenS16{F0: (fi126 ^ (fi126 & int(v1)))}
                                    _ = v127
                                    _ = v128
                                } else {
                                    var v129 int8 = -128
                                    v4 += 916
                                    _ = v129
                                }
                                var v130 int = (fi126 % ((fi126) | 1))
                                if ((true || (!(0.1 >= 1e10))) || (true || true)) {
                                    fmt.Println(fi126, fi126)
                                }
                                _ = v130
                            } else {
                                v3 = v3
                            }
                        default:
                            var v131 uint64 = 18446744073709551615
                            _ = v131
                        }
                        v2 += v2
                        _ = fi126
                    }
                }
            case "":
                if false {
                    v132 := []int{}
                    v1 += v1
                    v2 -= 3.14159
                    _ = v132
                }
            case "\n\t":
                var v133 uint32 = (^(v3))
                _ = v133
            }
        }
        if ((((("\n\t" + "\n\t") + "hello") < "") && (!((true || false) || ("\n\t" <= "a")))) || (1 > -57)) {
            for fi134 := 0; fi134 < 9; fi134++ {
                fmt.Println(v4)
                if (false || (v3 < 1)) {
                    fmt.Println(v1, v2, v4)
                    if (!false) {
                        v135 := GenS17{F0: 4294967295, F1: "hello", F2: 181}
                        _ = v135
                    }
                }
                switch "foo bar" {
                case "hello":
                    if (!false) {
                        if ((int16(fi134) <= 735) || true) {
                            switch "foo bar" {
                            case "\n\t":
                                v136 := []int{4294967295}
                                _ = v136
                            case "hello":
                                var v137 float64 = float64(1e10)
                                _ = v137
                            case "foo bar":
                                fmt.Println(v1, v4, v2)
                            }
                            fmt.Println(v2, v2, v1)
                        } else {
                            v2 = -2.25
                            for fi138 := 0; fi138 < 9; fi138++ {
                                if true {
                                    switch ("\n\t" + "a") {
                                    case "\n\t":
                                        if true {
                                            fmt.Println(v3, v2, v1)
                                        }
                                    case "a":
                                        v139 := []int{}
                                        var v140 uint32 = v3
                                        _ = v139
                                        _ = v140
                                    default:
                                        switch ((("hello" + "hello") + "foo bar") + "") {
                                        case "hello":
                                            if (!false) {
                                                v141 := []int{(fi134 * fi138), (-(fi138)), (^(fi134)), 428}
                                                if len(v141) > 0 {
                                                    v141[0] = fi138
                                                }
                                                v142 := GenS14{F0: "a", F1: (v1 >> ((42) & 63)), F2: -9223372036854775808}
                                                _ = v141
                                                _ = v142
                                            } else {
                                                var v143 int32 = (^(v4))
                                                _ = v143
                                            }
                                        }
                                    }
                                    if false {
                                        v144 := []int{}
                                        var v145 float64 = (v2 / 0.1)
                                        var v146 uint8 = uint8(v3)
                                        _ = v144
                                        _ = v145
                                        _ = v146
                                    }
                                    fmt.Println(fi134, v3, fi134)
                                    v1 -= 50
                                    v147 := GenS12{F0: uint8(v1), F1: 42, F2: (fi138 - fi134)}
                                    _ = v147
                                } else {
                                    v148 := []int{(fi138 & fi138), (^(fi134))}
                                    v2 -= (float64(0.1) - (((v2 - ((v2 - v2) - float64(-2.25))) * v2) / (v2 - ((1.5 + (1.5 + v2)) * (v2 + (v2 + ((v2 - ((v2 / v2) - (-2.25 * v2))) * v2)))))))
                                    v149 := GenS7{F0: "\n\t", F1: 0, F2: 1}
                                    _ = v148
                                    _ = v149
                                }
                                v4 -= (v4 * (-(v4)))
                                var v150 float64 = (((v2 - v2) / (v2 + v2)) + (3.14159 + v2))
                                _ = fi138
                                _ = v150
                            }
                        }
                        var v151 int64 = 42
                        _ = v151
                    } else {
                        if false {
                            fmt.Println(v1)
                        } else {
                            v152 := []int{}
                            if ((1.5 * float64(((1.5 + 1e10) * 1.5))) < float64((float32((v2 / 1e10)) - 1e10))) {
                                if (!false) {
                                    switch ("foo bar" + (("" + ("a" + "hello")) + ("a" + ("a" + "\n\t")))) {
                                    case "":
                                        v4 = (v4 ^ (v4 & 2147483647))
                                    case "\n\t":
                                        var v153 int64 = -623
                                        _ = v153
                                    case "foo bar":
                                        var v154 int8 = 94
                                        _ = v154
                                    }
                                } else {
                                    for fi155 := 0; fi155 < 20; fi155++ {
                                        v3 = 42
                                        _ = fi155
                                    }
                                    var v156 int32 = (-(v4))
                                    _ = v156
                                }
                                var v157 int = (fi134 - fi134)
                                fmt.Println(v1)
                                _ = v157
                            }
                            _ = v152
                        }
                        if (2 < 255) {
                            v158 := GenS18{F0: ((fi134 | (fi134 & (fi134 / (((^(fi134))) | 1)))) >= (fi134 >> ((65535) & 63))), F1: v4, F2: 1}
                            v3 = v3
                            var v159 int64 = int64(v1)
                            _ = v158
                            _ = v159
                        }
                        var v160 int32 = v4
                        _ = v160
                    }
                case "foo bar":
                    for fi161 := 0; fi161 < 9; fi161++ {
                        var v162 bool = false
                        var v163 uint = 4294967295
                        _ = fi161
                        _ = v162
                        _ = v163
                    }
                }
                _ = fi134
            }
            v164 := []int{}
            _ = v164
        }
    }
    v165 := GenS19{F0: 0, F1: 394, F2: (^(v3))}
    fmt.Println(v2, v4, v3)
    v2 = v2
    switch -2147483648 {
    case -263:
        v166 := GenS15{F0: "a"}
        _ = v166
    case 1:
        v167 := []int{4294967295, 42, 4294967295, 164}
        _ = v167
    case -9223372036854775808:
        v2 += (v2 * ((1.5 - 3.14159) - ((0.0 + 1e10) * v2)))
    }
    for fi168 := 0; fi168 < 1; fi168++ {
        v169 := GenS10{F0: fi168, F1: 9223372036854775807, F2: 1}
        _ = fi168
        _ = v169
    }
    var v170 int16 = 32767
    switch 1 {
    case 42:
        if false {
            v171 := GenS5{F0: 1, F1: 2}
            v172 := []int{739}
            for fi173 := 0; fi173 < 3; fi173++ {
                if (2 >= 18446744073709551615) {
                    for fi174 := 0; fi174 < 20; fi174++ {
                        _ = len(v172)
                        fmt.Println(fi173, v1)
                        _ = fi174
                    }
                    v175 := GenS3{F0: -1}
                    v172 = append(v172, fi173)
                    _ = v175
                }
                _ = fi173
            }
            _ = v171
            _ = v172
        } else {
            for fi176 := 0; fi176 < 5; fi176++ {
                v2 = (0.1 + (v2 - 1.5))
                v177 := []int{}
                _ = fi176
                _ = v177
            }
            v2 += v2
        }
    default:
        var v178 bool = false
        _ = v178
    }
    if (false || true) {
        for fi179 := 0; fi179 < 6; fi179++ {
            for fi180 := 0; fi180 < 3; fi180++ {
                var v181 uint16 = v1
                _ = fi180
                _ = v181
            }
            var v182 uint = 109
            _ = fi179
            _ = v182
        }
    }
    for fi183 := 0; fi183 < 4; fi183++ {
        switch (("foo bar" + ("foo bar" + (("a" + "\n\t") + "hello"))) + ("a" + "foo bar")) {
        case "a":
            for fi184 := 0; fi184 < 17; fi184++ {
                v185 := []int{fi183}
                v186 := GenS20{F0: (fi184 | fi183)}
                v187 := GenS10{F0: 128, F1: 9223372036854775807, F2: 323}
                _ = fi184
                _ = v185
                _ = v186
                _ = v187
            }
        case "":
            for fi188 := 0; fi188 < 17; fi188++ {
                v189 := []int{(fi188 & fi188), -974}
                fmt.Println(v170)
                v1 += (v1 * v1)
                _ = fi188
                _ = v189
            }
        default:
            if (!(true || true)) {
                for fi190 := 0; fi190 < 8; fi190++ {
                    var v191 uint8 = 128
                    _ = fi190
                    _ = v191
                }
                fmt.Println(fi183)
            } else {
                v192 := []int{}
                v2 += ((v2 + ((v2 * float64((float32(((1e10 * float64(1.5)) - v2)) - (1e10 + 1e10)))) * v2)) / v2)
                if (int8(v4) == -128) {
                    _ = len(v192)
                }
                _ = v192
            }
        }
        for fi193 := 0; fi193 < 2; fi193++ {
            v2 -= (v2 * (v2 / (v2 * v2)))
            fmt.Println(fi183, v2, v1)
            switch "a" {
            case "a":
                v1 += 660
            }
            _ = fi193
        }
        _ = fi183
    }
    var v194 uint = 255
    v195 := GenS6{F0: uint(v1), F1: 42, F2: uint64(v1)}
    v1 -= uint16(v194)
    v194 = v194
    fmt.Println("v1:", v1)
    fmt.Println("v170:", v170)
    fmt.Println("v194:", v194)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v165
    _ = v170
    _ = v194
    _ = v195
    _ = v2
    _ = v3
    _ = v4
}
