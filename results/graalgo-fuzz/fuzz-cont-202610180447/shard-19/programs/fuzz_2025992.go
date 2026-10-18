package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 string; F1 uint8; F2 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 int16; F1 bool; F2 int8 }
type GenS12 struct { F0 uint; F1 string; F2 int32 }
type GenS13 struct { F0 int32; F1 bool; F2 string }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint8; F1 int64; F2 int32 }
type GenS16 struct { F0 int64; F1 uint16; F2 bool }
type GenS17 struct { F0 uint; F1 uint; F2 uint64 }
func genFn0(p0 int64, p1 uint, p2 string) uint64 {
    v7 := []int{127, -849, 2147483647, 0}
    switch p2 {
    case "a":
        _ = len(v7)
    default:
        v8 := GenS1{F0: -1}
        _ = v8
    }
    v9 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return 0
}

func main() {
    var v1 int8 = 127
    v2 := []int{-368}
    v3 := []string{"foo bar"}
    v4 := make(map[int]int)
    var v5 uint16 = 255
    var v6 uint = 18446744073709551615
    fmt.Println(v6, v1, v6)
    v10 := GenS1{F0: v1}
    var v11 uint64 = 4294967295
    if false {
        if ((((0.1 * -2.25) - 1.5) / -2.25) == 1e10) {
            for fi12 := 0; fi12 < 20; fi12++ {
                if true {
                    if (!(!true)) {
                        if ((v1 * v1) < 102) {
                            for fi13 := 0; fi13 < 12; fi13++ {
                                v14 := GenS1{F0: (v1 | v1)}
                                fmt.Println(fi12)
                                if false {
                                    var v15 uint32 = 0
                                    if (!false) {
                                        if true {
                                            v11 -= v11
                                            var v16 bool = (653 < v15)
                                            if ((v16 && false) && (((-2.25 * ((-2.25 / 1.0) / 1.0)) - 1.5) != 1e10)) {
                                                var v17 int32 = -1
                                                fmt.Println(v6, fi13, fi12)
                                                v16 = (-2.25 != 1e10)
                                                v6 += v6
                                                var v18 int64 = -1
                                                _ = v17
                                                _ = v18
                                            }
                                            if (true && false) {
                                                v19 := GenS1{F0: 12}
                                                _ = v19
                                            }
                                            _ = v16
                                        } else {
                                            if (0.0 >= 3.14159) {
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                                if false {
                                                    var v20 uint16 = 737
                                                    _ = v20
                                                }
                                            }
                                            var v21 float64 = ((1.5 + (0.1 / float64(0.1))) - 3.14159)
                                            if ((uint16(v11) > v5) && true) {
                                                switch (("foo bar" + "hello") + ((("foo bar" + "a") + "a") + "foo bar")) {
                                                case "":
                                                    if ("" != "a") {
                                                        v1 = (v1 / ((v1) | 1))
                                                    } else {
                                                        v21 += ((v21 / ((v21 + v21) / v21)) - v21)
                                                    }
                                                case "\n\t":
                                                    v1 = v1
                                                }
                                                _ = len(v4)
                                                v22 := GenS1{F0: v1}
                                                v4[fi13] = -2147483648
                                                var v23 int = fi12
                                                if (false && ((v5 > 2) || false)) {
                                                    v15 = v15
                                                    v11 = 1
                                                } else {
                                                    if (!(-1 > -2147483648)) {
                                                        v5 += (^(v5))
                                                    } else {
                                                        if ((int8(v11) > 100) && false) {
                                                            if false {
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    switch "" {
                                                                    case "hello":
                                                                        switch 0 {
                                                                        case 127:
                                                                            if (!((v6 ^ 65535) >= (v6 & v6))) {
                                                                                var v24 bool = (!((!(("" != "hello") || true)) || true))
                                                                                _ = v24
                                                                            }
                                                                        case 1:
                                                                            v1 = int8(v5)
                                                                        case 32767:
                                                                            switch (fi13 % ((v23) | 1)) {
                                                                            case -979:
                                                                                v4[(fi12 | 2147483647)] = 127
                                                                            case 151:
                                                                                switch fi12 {
                                                                                case 0:
                                                                                    fmt.Println(v21)
                                                                                case 128:
                                                                                    v25 := GenS2{F0: 1}
                                                                                    _ = v25
                                                                                default:
                                                                                    fmt.Println(fi13, fi13, fi13)
                                                                                }
                                                                            case 1:
                                                                                switch fi12 {
                                                                                case 459:
                                                                                    fmt.Println(v15, v11, v11)
                                                                                case 128:
                                                                                    var v26 bool = (42 <= -1)
                                                                                    _ = v26
                                                                                case 4294967295:
                                                                                    _ = len(v4)
                                                                                default:
                                                                                    v21 += (1e10 - -2.25)
                                                                                    v6 = 2
                                                                                    v27 := GenS1{F0: (v1 ^ -1)}
                                                                                    _ = v27
                                                                                }
                                                                            default:
                                                                                v23 += (fi13 - (fi12 - v23))
                                                                            }
                                                                        default:
                                                                            v5 -= (v5 & uint16(v11))
                                                                        }
                                                                    case "":
                                                                        fmt.Println(v6)
                                                                    }
                                                                default:
                                                                    v28 := v4[int(v6)]
                                                                    _ = v28
                                                                }
                                                                if (!false) {
                                                                    v11 = v11
                                                                    if true {
                                                                        var v29 int64 = 0
                                                                        _ = v29
                                                                    }
                                                                    v30 := GenS3{F0: (v5 * 767), F1: v15}
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    v31 := v4[-1]
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    v32 := GenS4{F0: (((0.0 - ((v21 / (float64(-2.25) / v21)) + (v21 + v21))) / v21) != ((float64((0.1 + 1.5)) + 1.5) + float64(float32(0.1)))), F1: int(v6)}
                                                                    _ = v32
                                                                }
                                                                v21 = v21
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v5, v5)
                                                                }
                                                                var v33 float32 = (0.1 / 0.1)
                                                                if true {
                                                                    if (!(73 != int16(v1))) {
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        v1 += (v1 >> ((65535) & 63))
                                                                        var v34 int32 = int32(v1)
                                                                        _ = v34
                                                                    } else {
                                                                        v35 := v4[v23]
                                                                        _ = v35
                                                                    }
                                                                    _ = len(v3)
                                                                    if (true && (v1 >= -1)) {
                                                                        _ = len(v4)
                                                                        fmt.Println(v6)
                                                                    } else {
                                                                        if ((v15 / ((v15) | 1)) <= (v15 - v15)) {
                                                                            var v36 uint8 = 142
                                                                            v37 := GenS5{F0: -288}
                                                                            v15 += v15
                                                                            _ = v36
                                                                            _ = v37
                                                                        } else {
                                                                            fmt.Println(v5, v33)
                                                                            v21 = ((v21 - v21) * v21)
                                                                        }
                                                                        fmt.Println(v23)
                                                                    }
                                                                } else {
                                                                    var v38 string = ""
                                                                    if false {
                                                                        v11 -= (v11 | genFn0(-1, (v6 * (v6 * (v6 | (^(v6))))), v38))
                                                                        v5 -= 32767
                                                                        switch (-(fi12)) {
                                                                        case 569:
                                                                            var v39 int16 = -1
                                                                            _ = v39
                                                                        }
                                                                    } else {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                    _ = v38
                                                                }
                                                                _ = v33
                                                            }
                                                            v4[(fi12 + (-(fi13)))] = (^(fi13))
                                                            var v40 uint = 65535
                                                            _ = v40
                                                        }
                                                        v11 = genFn0(-801, v6, ("" + (("foo bar" + "foo bar") + ("hello" + "a"))))
                                                        switch ("a" + "\n\t") {
                                                        case "hello":
                                                            switch "" {
                                                            case "foo bar":
                                                                var v41 string = "foo bar"
                                                                _ = v41
                                                            default:
                                                                switch (fi12 ^ 655) {
                                                                case 4294967295:
                                                                    var v42 int32 = 0
                                                                    _ = v42
                                                                case -190:
                                                                    var v43 int32 = int32(v1)
                                                                    _ = v43
                                                                case 0:
                                                                    v44 := GenS6{F0: "foo bar", F1: 120, F2: 42}
                                                                    _ = v44
                                                                default:
                                                                    var v45 int32 = 42
                                                                    _ = v45
                                                                }
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(fi12, v1, v11)
                                                        }
                                                    }
                                                }
                                                _ = v22
                                                _ = v23
                                            } else {
                                                v46 := GenS6{F0: "\n\t", F1: 0, F2: -1}
                                                var v47 uint64 = v11
                                                fmt.Println(v21, v21)
                                                _ = v46
                                                _ = v47
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            fmt.Println(v5, v1, v1)
                                            _ = v21
                                        }
                                    }
                                    v48 := GenS7{F0: v5}
                                    _ = v15
                                    _ = v48
                                }
                                var v49 uint = v6
                                _ = fi13
                                _ = v14
                                _ = v49
                            }
                            if (1.5 > 1.5) {
                                var v50 int8 = -1
                                v5 = uint16(v11)
                                _ = v50
                            }
                            v6 = v6
                        }
                    }
                }
                var v51 uint = (v6 + (v6 ^ v6))
                v4[(^(fi12))] = fi12
                _ = fi12
                _ = v51
            }
        }
        v11 -= genFn0(-262, 2, "")
        if ((!(!true)) || (63 >= 1)) {
            if true {
                fmt.Println(v11)
            } else {
                delete(v4, 9223372036854775807)
                v1 -= v1
            }
        } else {
            var v52 uint16 = 221
            switch "a" {
            case "\n\t":
                if true {
                    var v53 int = int(v5)
                    _ = v53
                }
            case "hello":
                v3 = append(v3, "hello")
            default:
                v52 -= (v52 * (v5 << (((v6 % ((v6) | 1))) & 63)))
            }
            if true {
                fmt.Println(v11, v52, v1)
            }
            _ = v52
        }
    }
    for fi54 := 0; fi54 < 20; fi54++ {
        if true {
            if true {
                if true {
                    switch fi54 {
                    case 0:
                        v5 = (^(v5))
                    case 913:
                        v5 = v5
                    case -847:
                        v1 = (v1 - int8(v6))
                    default:
                        switch "foo bar" {
                        case "hello":
                            for fi55 := 0; fi55 < 12; fi55++ {
                                v56 := GenS5{F0: 42}
                                v1 = v1
                                switch (fi55 + (-(fi54))) {
                                case 127:
                                    var v57 bool = true
                                    _ = v57
                                case 224:
                                    v3 = append(v3, "a")
                                case 619:
                                    v1 += v1
                                }
                                _ = fi55
                                _ = v56
                            }
                        }
                    }
                } else {
                    v3 = append(v3, (("\n\t" + "hello") + "foo bar"))
                    var v58 uint32 = 42
                    v4[int(v1)] = (fi54 - 4294967295)
                    _ = v58
                }
                switch (("foo bar" + ("a" + "a")) + "a") {
                case "hello":
                    var v59 uint16 = (^(v5))
                    _ = v59
                case "":
                    v11 -= v11
                case "a":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        } else {
            fmt.Println(fi54)
            var v60 string = ""
            _ = v60
        }
        if (22 > 533) {
            v6 += (^(v6))
        } else {
            v6 += 2
            if (v6 <= (v6 - (^(v6)))) {
                for fi61 := 0; fi61 < 16; fi61++ {
                    fmt.Println(v5, v11, v11)
                    _ = fi61
                }
                v62 := v4[(fi54 % (((fi54 & -1)) | 1))]
                _ = v62
            } else {
                switch "foo bar" {
                case "hello":
                    var v63 uint32 = 4294967295
                    _ = v63
                case "a":
                    if false {
                        switch "hello" {
                        case "hello":
                            if true {
                                if false {
                                    var v64 uint32 = 1
                                    _ = v64
                                }
                                if (int(v1) > (fi54 ^ int(v1))) {
                                    var v65 int16 = -1
                                    _ = v65
                                }
                                _ = len(v4)
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = fi54
                                }
                                fmt.Println(fi54, v1, v1)
                                if false {
                                    v11 -= genFn0(568, 0, "a")
                                    var v66 uint16 = v5
                                    _ = v66
                                } else {
                                    if ("hello" <= "") {
                                        for fi67 := 0; fi67 < 10; fi67++ {
                                            var v68 int64 = -497
                                            if false {
                                                fmt.Println(v11)
                                                switch "foo bar" {
                                                case "":
                                                    var v69 int = fi67
                                                    switch (("" + "\n\t") + ("foo bar" + ("a" + ("" + (("\n\t" + ("foo bar" + "foo bar")) + ""))))) {
                                                    case "a":
                                                        v70 := GenS8{F0: (v11 | genFn0(v68, v6, (("" + "foo bar") + "a"))), F1: 2}
                                                        _ = v70
                                                    case "":
                                                        var v71 int16 = 42
                                                        _ = v71
                                                    default:
                                                        if true {
                                                            v72 := GenS6{F0: ("a" + (("" + "foo bar") + ("" + "hello"))), F1: 42, F2: -32768}
                                                            fmt.Println(fi67, v11)
                                                            v1 += v1
                                                            _ = v72
                                                        }
                                                    }
                                                    _ = v69
                                                case "\n\t":
                                                    var v73 int = fi67
                                                    _ = v73
                                                default:
                                                    fmt.Println(v5)
                                                }
                                            }
                                            _ = fi67
                                            _ = v68
                                        }
                                        if false {
                                            var v74 float64 = (-2.25 / float64(-2.25))
                                            if (-2.25 < -2.25) {
                                                for fi75 := 0; fi75 < 19; fi75++ {
                                                    var v76 int64 = -9223372036854775808
                                                    _ = fi75
                                                    _ = v76
                                                }
                                                delete(v4, (fi54 >> ((v6) & 63)))
                                                v77 := GenS9{F0: -9223372036854775808}
                                                _ = v77
                                            }
                                            _ = v74
                                        } else {
                                            for fi78 := 0; fi78 < 6; fi78++ {
                                                v3 = append(v3, ((("foo bar" + "\n\t") + "") + ((("hello" + "a") + "") + "foo bar")))
                                                _ = fi78
                                            }
                                            v5 = (^(v5))
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v79 := GenS10{F0: (^(v5)), F1: (v5 - 773)}
                                        _ = v79
                                    }
                                    var v80 int = int(v5)
                                    v81 := GenS1{F0: v1}
                                    _ = v80
                                    _ = v81
                                }
                            }
                        default:
                            for fi82 := 0; fi82 < 9; fi82++ {
                                fmt.Println(v1)
                                _ = fi82
                            }
                        }
                        v11 = v11
                        switch fi54 {
                        case 127:
                            if (2147483647 != 0) {
                                switch -384 {
                                case -2147483648:
                                    v5 += (v5 - v5)
                                case 42:
                                    if ((!(!true)) || true) {
                                        var v83 int = 42
                                        v84 := GenS11{F0: -1, F1: true, F2: (v1 / ((-1) | 1))}
                                        var v85 uint8 = 0
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    } else {
                                        v1 += (v1 >> (((v6 | v6)) & 63))
                                    }
                                }
                                for fi86 := 0; fi86 < 5; fi86++ {
                                    fmt.Println(fi54, v5)
                                    _ = fi86
                                }
                                v87 := GenS12{F0: (v6 + (v6 * v6)), F1: "\n\t", F2: 1}
                                _ = v87
                            } else {
                                v11 += 375
                                v1 += -1
                            }
                        case 566:
                            if len(v2) > 0 {
                                v2[0] = 197
                            }
                        default:
                            for fi88 := 0; fi88 < 11; fi88++ {
                                if (0 != 0) {
                                    var v89 int64 = -1
                                    _ = v89
                                } else {
                                    v90 := GenS7{F0: (v5 & (^(v5)))}
                                    fmt.Println(v6, v11)
                                    v5 -= v5
                                    _ = v90
                                }
                                _ = fi88
                            }
                        }
                    } else {
                        v11 = v11
                        switch fi54 {
                        case 880:
                            if true {
                                _ = len(v2)
                                v5 += (v5 + 42)
                                if true {
                                    switch fi54 {
                                    case 0:
                                        var v91 uint8 = 0
                                        _ = v91
                                    }
                                    if true {
                                        v6 += v6
                                        for fi92 := 0; fi92 < 16; fi92++ {
                                            v5 -= (v5 << (((v6 ^ v6)) & 63))
                                            _ = fi92
                                        }
                                    }
                                }
                            } else {
                                _ = len(v3)
                                var v93 int16 = 0
                                _ = v93
                            }
                        case -346:
                            v94 := GenS13{F0: 0, F1: (!((!true) || (v5 == (v5 + (v5 / (((v5 >> ((v6) & 63))) | 1)))))), F2: "foo bar"}
                            _ = v94
                        }
                        v95 := GenS14{F0: (!(!true))}
                        _ = v95
                    }
                case "foo bar":
                    v96 := GenS15{F0: 2, F1: -9223372036854775808, F2: 460}
                    _ = v96
                }
                var v97 bool = (!false)
                _ = v97
            }
        }
        if len(v3) > 0 {
            v3[0] = (("hello" + "a") + "")
        }
        _ = fi54
    }
    v6 = v6
    v98 := GenS14{F0: false}
    switch -1 {
    case 32768:
        v99 := GenS16{F0: 1, F1: (v5 | 851), F2: true}
        _ = v99
    case 0:
        for fi100 := 0; fi100 < 13; fi100++ {
            v11 = (v11 * genFn0(9223372036854775807, (^(v6)), (("foo bar" + "hello") + "")))
            switch (^(fi100)) {
            case 143:
                v1 = (v1 ^ (v1 >> ((1) & 63)))
            case 32768:
                var v101 int = (-(fi100))
                _ = v101
            case 127:
                fmt.Println(v6, v11, v6)
            }
            var v102 int = int(v5)
            _ = fi100
            _ = v102
        }
    }
    for fi103 := 0; fi103 < 18; fi103++ {
        for fi104 := 0; fi104 < 1; fi104++ {
            _ = len(v4)
            var v105 int64 = 681
            var v106 uint64 = (v11 & 2)
            _ = fi104
            _ = v105
            _ = v106
        }
        v5 = v5
        v107 := GenS17{F0: v6, F1: 255, F2: genFn0(-665, (v6 ^ (v6 / ((18446744073709551615) | 1))), "\n\t")}
        _ = fi103
        _ = v107
    }
    for fi108 := 0; fi108 < 2; fi108++ {
        for fi109 := 0; fi109 < 14; fi109++ {
            _ = len(v3)
            fmt.Println(fi108, v11, v5)
            switch 4294967295 {
            case -132:
                v5 += (v5 % ((1) | 1))
                v6 = v6
                fmt.Println(v6, v6)
            case 128:
                var v110 int = (fi109 / ((128) | 1))
                _ = v110
            case -2147483648:
                fmt.Println(v6)
            default:
                var v111 uint = v6
                _ = v111
            }
            _ = fi109
        }
        _ = fi108
    }
    if (true && (!(false || (!(0.0 <= 1e10))))) {
        v1 = 127
    } else {
        for fi112 := 0; fi112 < 7; fi112++ {
            var v113 uint8 = 1
            v114 := GenS11{F0: -32768, F1: (584 != -821), F2: v1}
            var v115 uint8 = v113
            _ = fi112
            _ = v113
            _ = v114
            _ = v115
        }
    }
    fmt.Println(v11, v1)
    if (0 >= 1) {
        for fi116 := 0; fi116 < 16; fi116++ {
            if true {
                fmt.Println(v11, v11)
            }
            _ = len(v4)
            _ = fi116
        }
        var v117 uint8 = 127
        _ = v117
    }
    fmt.Println(v11, v1)
    var v118 uint16 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v118:", v118)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
