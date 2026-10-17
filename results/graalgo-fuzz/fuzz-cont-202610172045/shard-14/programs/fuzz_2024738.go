package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 uint16; F1 int64 }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint64 }
type GenS10 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS11 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS12 struct { F0 int; F1 uint; F2 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 uint; F1 int8 }
type GenS16 struct { F0 int; F1 int8; F2 int8 }
type GenS17 struct { F0 int64; F1 uint8; F2 int }
type GenS18 struct { F0 int8; F1 uint }
type GenS19 struct { F0 bool; F1 bool; F2 uint64 }
type GenS20 struct { F0 bool }
type GenS21 struct { F0 int8; F1 uint64 }
type GenS22 struct { F0 int8; F1 bool }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 int64; F1 int64; F2 string }
func genFn0(p0 string, p1 uint8, p2 int8) uint64 {
    v5 := []int{42, int(p2)}
    if false {
        fmt.Println(p1, p0)
    }
    var v6 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint64, p1 int16, p2 bool) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: -1, F1: 0}
        p2 = p2
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        for fi10 := 0; fi10 < 20; fi10++ {
            p1 = (p1 - (p1 + (p1 + p1)))
            p1 = (p1 ^ p1)
            var v11 int8 = 1
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := []int{9223372036854775807, 703}
    for fi13 := 0; fi13 < 10; fi13++ {
        if (true || (p1 >= (p1 - p1))) {
            fmt.Println(fi13, fi13)
            switch fi13 {
            case 128:
                if (((p0 >> ((159) & 63)) > genFn0("hello", 59, -1)) && p2) {
                    p2 = (p2 || p2)
                    var v14 int32 = 380
                    _ = v14
                }
            case -412:
                if ((1.5 == 0.0) || p2) {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if p2 {
                            if len(v12) > 0 {
                                v12[0] = 571
                            }
                        }
                        p0 = (p0 | 330)
                        _ = fi15
                    }
                    var v16 int64 = 0
                    _ = v16
                } else {
                    var v17 int64 = 42
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if p2 {
                            switch (fi18 ^ (fi18 % ((2147483647) | 1))) {
                            case -9223372036854775808:
                                switch (fi18 << ((212) & 63)) {
                                case 4294967295:
                                    var v19 float32 = (0.1 - -2.25)
                                    switch ("a" + ("foo bar" + "\n\t")) {
                                    case "":
                                        fmt.Println(v19, v17)
                                    case "\n\t":
                                        v17 = 325
                                    }
                                    var v20 int32 = 0
                                    _ = v19
                                    _ = v20
                                case -78:
                                    _ = len(v12)
                                default:
                                    var v21 uint32 = 2
                                    _ = v21
                                }
                            case 128:
                                switch (("hello" + "hello") + ((("foo bar" + "") + "\n\t") + (("hello" + (("" + "hello") + "a")) + ""))) {
                                case "foo bar":
                                    if (("foo bar" + "foo bar") <= "\n\t") {
                                        v22 := GenS1{F0: 100, F1: 255}
                                        var v23 int8 = 93
                                        if false {
                                            if (p2 || (!(0 > 1))) {
                                                var v24 float32 = 0.0
                                                fmt.Println(p0)
                                                if (!false) {
                                                    if p2 {
                                                        v25 := GenS1{F0: v23, F1: 747}
                                                        _ = v25
                                                    } else {
                                                        switch (fi18 / (((-(fi18))) | 1)) {
                                                        case 989:
                                                            p2 = (!(p2 || false))
                                                        }
                                                    }
                                                    if p2 {
                                                        fmt.Println(v17, p2)
                                                    } else {
                                                        var v26 string = ""
                                                        v17 = v17
                                                        _ = v26
                                                    }
                                                    var v27 uint32 = 4294967295
                                                    _ = v27
                                                }
                                                _ = v24
                                            }
                                            v28 := GenS2{F0: int16(v17), F1: 65535}
                                            _ = v28
                                        }
                                        _ = v22
                                        _ = v23
                                    }
                                case "hello":
                                    if p2 {
                                        _ = len(v12)
                                        if p2 {
                                            if p2 {
                                                v29 := GenS3{F0: genFn0("a", 255, 1), F1: fi13, F2: ("a" + "a")}
                                                var v30 uint8 = 1
                                                p0 = (^(p0))
                                                _ = v29
                                                _ = v30
                                            }
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                            switch fi13 {
                                            case 9223372036854775807:
                                                var v31 int = fi13
                                                _ = v31
                                            case -2147483648:
                                                p1 = 1
                                            default:
                                                if p2 {
                                                    v12 = append(v12, 0)
                                                } else {
                                                    var v32 bool = p2
                                                    v33 := GenS4{F0: p2, F1: ((1.5 + (1e10 / 1.0)) > (-2.25 / 1.5))}
                                                    v32 = p2
                                                    p2 = v32
                                                    _ = v32
                                                    _ = v33
                                                }
                                            }
                                            v34 := GenS1{F0: 100, F1: 577}
                                            if len(v12) > 0 {
                                                v12[0] = fi13
                                            }
                                            _ = v34
                                        }
                                        if len(v12) > 0 {
                                            v12[0] = fi13
                                        }
                                        switch fi13 {
                                        case 493:
                                            if len(v12) > 0 {
                                                v12[0] = fi18
                                            }
                                        }
                                        v35 := GenS2{F0: (p1 - (p1 / (((p1 >> ((4294967295) & 63))) | 1))), F1: 695}
                                        _ = v35
                                    }
                                case "":
                                    var v36 int = 0
                                    _ = v36
                                default:
                                    v17 += (-(v17))
                                }
                            }
                            p0 = p0
                            v37 := GenS5{F0: 255, F1: 42}
                            switch fi18 {
                            case -405:
                                var v38 uint64 = (p0 & p0)
                                _ = v38
                            case -514:
                                var v39 float32 = 0.1
                                _ = v39
                            case -754:
                                fmt.Println(fi18)
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    if false {
                                        var v40 uint64 = p0
                                        var v41 float64 = (float64(((float32(((-2.25 / (float64(3.14159) + (0.1 - (0.0 * float64(3.14159))))) - (1e10 + 1.5))) + float32(1e10)) / 1e10)) * (3.14159 + 3.14159))
                                        _ = v40
                                        _ = v41
                                    } else {
                                        p1 = (p1 / ((p1) | 1))
                                        var v42 int64 = v17
                                        _ = v42
                                    }
                                case "foo bar":
                                    var v43 uint = 18446744073709551615
                                    _ = v43
                                }
                            }
                            _ = v37
                        } else {
                            v12 = append(v12, -2147483648)
                            if (!((978 <= p1) || p2)) {
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            } else {
                                switch ("hello" + ("hello" + ("\n\t" + "a"))) {
                                case "hello":
                                    fmt.Println(fi13, p0, fi18)
                                case "\n\t":
                                    p1 += p1
                                case "foo bar":
                                    p1 = p1
                                default:
                                    var v44 int8 = 100
                                    _ = v44
                                }
                            }
                            p2 = p2
                            switch (fi13 >> ((64) & 63)) {
                            case 523:
                                v45 := GenS3{F0: (p0 / (((p0 << ((240) & 63))) | 1)), F1: fi18, F2: ("\n\t" + "\n\t")}
                                _ = v45
                            default:
                                if (((fi13 << ((811) & 63)) != 1) || ((fi13 + -475) < fi18)) {
                                    v46 := GenS6{F0: 0, F1: 52}
                                    var v47 uint32 = 0
                                    _ = v46
                                    _ = v47
                                } else {
                                    p1 += (^(p1))
                                    _ = len(v12)
                                }
                            }
                            var v48 string = (("foo bar" + ("hello" + "\n\t")) + ("\n\t" + ""))
                            p1 = (p1 + -1)
                            v49 := GenS5{F0: 0, F1: (v17 / ((-447) | 1))}
                            _ = v48
                            _ = v49
                        }
                        switch 4294967295 {
                        case -208:
                            p1 = p1
                        case -2147483648:
                            if ((2 > 42) || false) {
                                p1 -= p1
                            } else {
                                v50 := GenS7{F0: "a"}
                                _ = v50
                            }
                        }
                        switch (("a" + "") + "a") {
                        case "hello":
                            v12 = append(v12, fi13)
                        default:
                            v51 := GenS7{F0: ""}
                            _ = v51
                        }
                        p1 += 32767
                        p2 = (!((p0 > genFn0(("foo bar" + "\n\t"), 128, -97)) && p2))
                        _ = fi18
                    }
                    v52 := GenS8{F0: p2}
                    _ = v17
                    _ = v52
                }
            case 9223372036854775807:
                v53 := GenS9{F0: 9223372036854775807, F1: 18446744073709551615}
                _ = v53
            default:
                for fi54 := 0; fi54 < 13; fi54++ {
                    switch "hello" {
                    case "a":
                        if p2 {
                            if len(v12) > 0 {
                                v12[0] = fi13
                            }
                            v55 := GenS10{F0: 2, F1: genFn0("hello", 192, -127), F2: uint8(p0)}
                            p2 = (!p2)
                            _ = v55
                        } else {
                            _ = len(v12)
                            if (fi54 != (fi54 * -9223372036854775808)) {
                                p0 = genFn0(("foo bar" + "\n\t"), 2, -86)
                                var v56 string = ((("\n\t" + "hello") + "\n\t") + ((("foo bar" + "a") + ("" + "")) + "a"))
                                p2 = p2
                                p1 += (-(p1))
                                _ = v56
                            }
                            if true {
                                var v57 uint8 = 0
                                var v58 uint32 = 42
                                v59 := GenS6{F0: p0, F1: 42}
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                        }
                    case "hello":
                        p1 = (^(p1))
                    case "":
                        var v60 string = "a"
                        _ = v60
                    default:
                        switch ("\n\t" + "hello") {
                        case "\n\t":
                            if p2 {
                                var v61 float64 = 1.5
                                v62 := GenS11{F0: p1, F1: 710, F2: 407}
                                _ = v61
                                _ = v62
                            }
                        case "foo bar":
                            var v63 int64 = 698
                            _ = v63
                        case "":
                            v64 := GenS12{F0: (fi54 * (fi54 % (((fi13 * (fi54 * fi13))) | 1))), F1: 441, F2: -1}
                            _ = v64
                        }
                    }
                    _ = fi54
                }
            }
        }
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]string)
    v3 := []string{("hello" + ("" + "hello")), "\n\t", ""}
    var v4 int32 = -2147483648
    for fi65 := 0; fi65 < 14; fi65++ {
        if (330 >= 255) {
            if true {
                switch "" {
                case "\n\t":
                    var v66 int8 = -128
                    _ = v66
                case "hello":
                    var v67 int = (fi65 + (fi65 + (fi65 << ((451) & 63))))
                    _ = v67
                case "":
                    v1 = 42
                default:
                    if (!false) {
                        for fi68 := 0; fi68 < 9; fi68++ {
                            v1 += (v1 + 0)
                            var v69 uint = 168
                            switch "hello" {
                            case "":
                                switch "foo bar" {
                                case "\n\t":
                                    if true {
                                        if (618 != 0) {
                                            if (!(2147483648 == 65535)) {
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                                v70 := GenS13{F0: v4}
                                                fmt.Println(v4)
                                                _ = v70
                                            } else {
                                                v4 = v4
                                                if true {
                                                    v69 = (^(v69))
                                                    if false {
                                                        v71 := GenS12{F0: 1, F1: v69, F2: 298}
                                                        v72 := GenS14{F0: v4}
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v73 := GenS11{F0: -32768, F1: 255, F2: 2}
                                                            _ = v73
                                                        case "a":
                                                            if true {
                                                                if true {
                                                                    switch ("a" + "\n\t") {
                                                                    case "foo bar":
                                                                        _ = len(v3)
                                                                    }
                                                                    if true {
                                                                        v69 += 2
                                                                        if ((float64(((3.14159 / 0.1) + (-2.25 * (1.5 / 0.1)))) - 0.0) > (0.1 - 0.0)) {
                                                                            v74 := GenS11{F0: (^(v1)), F1: 255, F2: 65535}
                                                                            _ = v74
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            _ = len(v3)
                                                                            if (-213 > int(v4)) {
                                                                                var v75 float64 = (1e10 * -2.25)
                                                                                v69 = (^(v69))
                                                                                _ = v75
                                                                            }
                                                                        } else {
                                                                            switch fi68 {
                                                                            case 9223372036854775807:
                                                                                v4 -= -2147483648
                                                                                v4 += v4
                                                                                v69 -= 2
                                                                            default:
                                                                                v4 += (v4 + (v4 | (v4 & (v4 >> ((v69) & 63)))))
                                                                                if false {
                                                                                    if (128 == 101) {
                                                                                        if (42 >= 0) {
                                                                                            v2["foo bar"] = "hello"
                                                                                            v76 := GenS15{F0: 42, F1: 42}
                                                                                            v1 -= v1
                                                                                            var v77 uint32 = 4294967295
                                                                                            _ = v76
                                                                                            _ = v77
                                                                                        } else {
                                                                                            v78 := v2["hello"]
                                                                                            _ = v78
                                                                                        }
                                                                                        v1 = v1
                                                                                    }
                                                                                    v4 -= (v4 << ((18446744073709551615) & 63))
                                                                                }
                                                                            }
                                                                            v4 -= 1
                                                                            switch fi68 {
                                                                            case 32768:
                                                                                var v79 string = "foo bar"
                                                                                var v80 uint64 = 42
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = "hello"
                                                                                }
                                                                                _ = v79
                                                                                _ = v80
                                                                            case 42:
                                                                                _ = len(v2)
                                                                            case 255:
                                                                                v4 = v4
                                                                            }
                                                                        }
                                                                    }
                                                                    var v81 uint = v69
                                                                    _ = v81
                                                                } else {
                                                                    var v82 int64 = -402
                                                                    var v83 uint16 = 575
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            }
                                                        default:
                                                            v84 := GenS3{F0: 142, F1: 9223372036854775807, F2: "foo bar"}
                                                            _ = v84
                                                        }
                                                        v4 = (v4 ^ (-(v4)))
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                } else {
                                                    _ = len(v3)
                                                    if true {
                                                        fmt.Println(v69)
                                                    } else {
                                                        v85 := v2["a"]
                                                        v86 := GenS12{F0: fi68, F1: (v69 ^ (v69 * v69)), F2: (v4 % ((v4) | 1))}
                                                        var v87 string = v85
                                                        v2[(v87 + "")] = ("foo bar" + ("foo bar" + v87))
                                                        v88 := GenS15{F0: (v69 * v69), F1: -1}
                                                        _ = v85
                                                        _ = v86
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    var v89 float32 = float32(1.5)
                                                    _ = v89
                                                }
                                            }
                                            if ((108 == 55) && false) {
                                                var v90 float32 = float32(1e10)
                                                v91 := GenS6{F0: 493, F1: 127}
                                                _ = v90
                                                _ = v91
                                            } else {
                                                if true {
                                                    v1 = v1
                                                    v4 = v4
                                                    switch "hello" {
                                                    case "hello":
                                                        _ = len(v2)
                                                    default:
                                                        var v92 int32 = 2147483647
                                                        _ = v92
                                                    }
                                                } else {
                                                    var v93 int8 = -128
                                                    _ = v93
                                                }
                                                switch "" {
                                                case "a":
                                                    if true {
                                                        v94 := GenS16{F0: 4294967295, F1: -31, F2: -128}
                                                        _ = v94
                                                    } else {
                                                        fmt.Println(fi68)
                                                    }
                                                case "\n\t":
                                                    var v95 string = ((("a" + ("hello" + "foo bar")) + (("" + "a") + "")) + "a")
                                                    _ = v95
                                                case "foo bar":
                                                    if true {
                                                        var v96 bool = false
                                                        delete(v2, "")
                                                        var v97 int = fi65
                                                        _ = v96
                                                        _ = v97
                                                    }
                                                    fmt.Println(fi68, fi65)
                                                default:
                                                    v4 += v4
                                                }
                                                var v98 int8 = int8(v4)
                                                _ = v98
                                            }
                                            var v99 int64 = 740
                                            _ = v99
                                        } else {
                                            var v100 uint64 = 42
                                            _ = v100
                                        }
                                        v4 = 141
                                        v2[((("a" + ("a" + "")) + (((("foo bar" + ("hello" + ("hello" + "foo bar"))) + "\n\t") + "a") + "\n\t")) + (("a" + (((("" + "foo bar") + "a") + "foo bar") + "")) + ""))] = ("foo bar" + "\n\t")
                                    } else {
                                        switch (fi68 % ((fi68) | 1)) {
                                        case 772:
                                            switch (^(fi68)) {
                                            case -1:
                                                _ = len(v2)
                                            case 32767:
                                                var v101 int16 = v1
                                                _ = v101
                                            case 162:
                                                if false {
                                                    v1 = (v1 & v1)
                                                    v1 = (v1 & (^(v1)))
                                                    switch ("hello" + ("hello" + "hello")) {
                                                    case "":
                                                        var v102 int = fi68
                                                        _ = v102
                                                    }
                                                }
                                            default:
                                                if ((!(v69 >= 0)) || true) {
                                                    var v103 string = ("" + "")
                                                    v69 = 2
                                                    if true {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if (0 > -57) {
                                                            if (((!true) || (!true)) || false) {
                                                                switch fi68 {
                                                                case 790:
                                                                    v103 = (((v103 + "\n\t") + ((("foo bar" + "") + ("hello" + (v103 + v103))) + v103)) + "foo bar")
                                                                case 9223372036854775807:
                                                                    var v104 float64 = ((1.5 + 1e10) - 0.0)
                                                                    _ = v104
                                                                case -405:
                                                                    v105 := GenS16{F0: fi65, F1: 0, F2: 1}
                                                                    if (65535 > 974) {
                                                                        v103 = v103
                                                                        fmt.Println(fi65, v4, fi68)
                                                                        v106 := GenS9{F0: 42, F1: genFn0(v103, 219, 0)}
                                                                        v107 := GenS16{F0: fi68, F1: -96, F2: -11}
                                                                        var v108 int8 = -128
                                                                        _ = v106
                                                                        _ = v107
                                                                        _ = v108
                                                                    } else {
                                                                        v103 = "foo bar"
                                                                    }
                                                                    fmt.Println(v4)
                                                                    _ = v105
                                                                default:
                                                                    fmt.Println(v4, v103)
                                                                }
                                                                if (fi68 == (fi68 - fi65)) {
                                                                    var v109 bool = (fi65 < fi65)
                                                                    v109 = v109
                                                                    _ = v109
                                                                } else {
                                                                    var v110 uint64 = 0
                                                                    var v111 bool = (65535 < 0)
                                                                    var v112 string = v103
                                                                    _ = v110
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                                v103 = v103
                                                            }
                                                        } else {
                                                            var v113 uint8 = 124
                                                            var v114 int8 = 123
                                                            switch (fi68 / ((fi68) | 1)) {
                                                            case 127:
                                                                if false {
                                                                    v115 := GenS16{F0: fi68, F1: v114, F2: 1}
                                                                    switch "" {
                                                                    case "":
                                                                        v2[v103] = v103
                                                                        var v116 float32 = float32((float64(0.0) + (1e10 * float64(3.14159))))
                                                                        v117 := GenS17{F0: -8, F1: 86, F2: 4294967295}
                                                                        var v118 int64 = 1
                                                                        _ = v116
                                                                        _ = v117
                                                                        _ = v118
                                                                    default:
                                                                        v1 += (v1 % ((v1) | 1))
                                                                    }
                                                                    _ = v115
                                                                }
                                                                v113 = (v113 ^ v113)
                                                            case 1:
                                                                v103 = "foo bar"
                                                            case 0:
                                                                v69 = (v69 & v69)
                                                            }
                                                            v114 -= v114
                                                            if (true || ((0.0 != 0.0) && true)) {
                                                                v3 = append(v3, (((v103 + ("foo bar" + v103)) + v103) + ""))
                                                            } else {
                                                                var v119 int8 = v114
                                                                _ = v119
                                                            }
                                                            v120 := GenS9{F0: genFn1(uint64(v113), v1, true), F1: 18446744073709551615}
                                                            _ = v113
                                                            _ = v114
                                                            _ = v120
                                                        }
                                                        if (v69 != v69) {
                                                            var v121 string = ("\n\t" + "")
                                                            _ = v121
                                                        } else {
                                                            delete(v2, (("" + ((v103 + "a") + "hello")) + v103))
                                                        }
                                                    } else {
                                                        if ((!false) && false) {
                                                            var v122 uint8 = 247
                                                            _ = v122
                                                        }
                                                    }
                                                    _ = v103
                                                }
                                            }
                                        default:
                                            fmt.Println(fi68, v1, fi65)
                                        }
                                        v69 = v69
                                    }
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            case "a":
                                switch (fi68 * 955) {
                                case 128:
                                    var v123 int64 = 9223372036854775807
                                    _ = v123
                                default:
                                    fmt.Println(fi65, fi65, fi68)
                                }
                            default:
                                fmt.Println(fi68, fi68)
                            }
                            if false {
                                if ((-1 >= 0) && (v69 != v69)) {
                                    var v124 uint = 42
                                    if (32767 >= 2) {
                                        v69 = (v69 / (((v69 ^ (v124 - 18446744073709551615))) | 1))
                                        v125 := v2["foo bar"]
                                        if (1e10 < float64(3.14159)) {
                                            if true {
                                                v126 := GenS3{F0: genFn0(v125, 2, -70), F1: 1, F2: ("" + v125)}
                                                _ = v126
                                            } else {
                                                v127 := GenS18{F0: -1, F1: v69}
                                                switch fi68 {
                                                case -72:
                                                    v1 -= (v1 << ((v124) & 63))
                                                case 127:
                                                    v125 = "a"
                                                case -2147483648:
                                                    if (v1 == 892) {
                                                        fmt.Println(fi68, v4)
                                                        v128 := GenS3{F0: 42, F1: 32767, F2: v125}
                                                        _ = v128
                                                    }
                                                    delete(v2, (v125 + "foo bar"))
                                                default:
                                                    switch 0 {
                                                    case 0:
                                                        fmt.Println(v124, v69)
                                                    case -392:
                                                        var v129 int64 = 42
                                                        _ = v129
                                                    case 42:
                                                        if true {
                                                            v69 += 65535
                                                            fmt.Println(v124)
                                                            fmt.Println(v125)
                                                        }
                                                    default:
                                                        if ("hello" == (v125 + "\n\t")) {
                                                            fmt.Println(v125, fi68, v125)
                                                        } else {
                                                            v69 += v69
                                                            v130 := GenS14{F0: (v4 - v4)}
                                                            _ = v130
                                                        }
                                                    }
                                                }
                                                _ = v127
                                            }
                                            var v131 int16 = 0
                                            switch v125 {
                                            case "":
                                                var v132 int16 = v131
                                                _ = v132
                                            default:
                                                if true {
                                                    var v133 bool = ((v131 - (v1 * (v1 ^ v1))) >= (v1 | (v131 << (((v124 + v124)) & 63))))
                                                    switch fi68 {
                                                    case 4294967295:
                                                        if v133 {
                                                            v69 -= v124
                                                            var v134 uint32 = 900
                                                            var v135 uint16 = 42
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                    case 32767:
                                                        if (fi65 >= (fi65 << ((1) & 63))) {
                                                            if (true && false) {
                                                                v1 += 1
                                                            }
                                                            v136 := GenS16{F0: (fi65 ^ -2147483648), F1: 83, F2: -56}
                                                            _ = v136
                                                        } else {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        var v137 float64 = float64((-2.25 - 0.0))
                                                        _ = v137
                                                    }
                                                    _ = v133
                                                }
                                            }
                                            _ = v131
                                        } else {
                                            fmt.Println(fi68, v125, v4)
                                        }
                                        v138 := v2["foo bar"]
                                        _ = v125
                                        _ = v138
                                    }
                                    _ = v124
                                }
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                                switch 608 {
                                case -1:
                                    v139 := GenS2{F0: -304, F1: 2}
                                    _ = v139
                                case 2147483647:
                                    fmt.Println(v4, fi65, v4)
                                default:
                                    v140 := GenS19{F0: (-2.25 <= 0.1), F1: false, F2: uint64(v69)}
                                    _ = v140
                                }
                            } else {
                                switch (-(fi65)) {
                                case 0:
                                    var v141 int64 = -1
                                    _ = v141
                                }
                                var v142 bool = (1 >= 0)
                                _ = v142
                            }
                            v143 := GenS20{F0: (false || (false || ("" < "a")))}
                            _ = fi68
                            _ = v143
                            _ = v69
                        }
                        v4 += v4
                    }
                }
            }
            v144 := GenS21{F0: -55, F1: 4294967295}
            _ = v144
        } else {
            var v145 uint8 = 139
            _ = v145
        }
        _ = fi65
    }
    v1 -= 32767
    for fi146 := 0; fi146 < 8; fi146++ {
        var v147 uint16 = 915
        _ = fi146
        _ = v147
    }
    for fi148 := 0; fi148 < 10; fi148++ {
        if ((-125 <= (fi148 << ((93) & 63))) || (((!false) || (false && true)) || (float32(3.14159) <= (float32(1e10) * 0.1)))) {
            v2[("foo bar" + ("foo bar" + "hello"))] = ("a" + "a")
            fmt.Println(fi148)
            switch (-(fi148)) {
            case 635:
                if true {
                    var v149 bool = true
                    var v150 int16 = v1
                    if true {
                        _ = len(v2)
                        delete(v2, "hello")
                        v149 = (((65535 == 2) && (((2 <= 42) && (1.5 >= 0.1)) || false)) || ((v1 - 0) >= 481))
                    }
                    _ = v149
                    _ = v150
                }
            default:
                var v151 float32 = 1e10
                _ = v151
            }
        }
        switch fi148 {
        case 32768:
            var v152 bool = (!(2 <= 18446744073709551615))
            _ = v152
        default:
            v1 -= 1
        }
        _ = fi148
    }
    v4 = 42
    if true {
        for fi153 := 0; fi153 < 9; fi153++ {
            if (false || true) {
                v4 += (v4 >> ((0) & 63))
                v154 := GenS22{F0: int8(v4), F1: false}
                _ = v154
            }
            if (65535 < 1) {
                var v155 int8 = 127
                for fi156 := 0; fi156 < 6; fi156++ {
                    v155 -= -1
                    switch "\n\t" {
                    case "":
                        switch ((("foo bar" + "") + ("a" + "a")) + "") {
                        case "foo bar":
                            v157 := GenS23{F0: v155}
                            _ = v157
                        case "hello":
                            if true {
                                var v158 int32 = (v4 + 42)
                                if (v158 <= v4) {
                                    fmt.Println(v158)
                                } else {
                                    var v159 uint = 18446744073709551615
                                    var v160 string = (("foo bar" + ((("hello" + "foo bar") + "") + ((("foo bar" + "") + ("" + "a")) + ("\n\t" + "a")))) + "foo bar")
                                    v1 -= v1
                                    _ = v159
                                    _ = v160
                                }
                                _ = v158
                            } else {
                                if (v155 != v155) {
                                    var v161 int32 = 1
                                    _ = v161
                                } else {
                                    var v162 int16 = v1
                                    _ = v162
                                }
                            }
                        }
                        _ = len(v2)
                        if (fi153 <= fi153) {
                            var v163 bool = (true && (fi153 == fi153))
                            v163 = ((v4 > v4) || false)
                            if (!false) {
                                fmt.Println(v155, v1, v155)
                            } else {
                                v4 -= 0
                            }
                            v155 -= v155
                            _ = v163
                        } else {
                            var v164 uint8 = 128
                            _ = v164
                        }
                    }
                    _ = fi156
                }
                _ = v155
            } else {
                var v165 uint32 = 42
                var v166 string = ("a" + "a")
                _ = v165
                _ = v166
            }
            _ = fi153
        }
        v1 = (v1 * v1)
    } else {
        switch "a" {
        case "":
            switch "" {
            case "\n\t":
                for fi167 := 0; fi167 < 11; fi167++ {
                    v4 -= v4
                    _ = fi167
                }
            default:
                switch -9223372036854775808 {
                case 2147483647:
                    v168 := GenS24{F0: 42, F1: 0, F2: ("hello" + "hello")}
                    _ = v168
                case 42:
                    for fi169 := 0; fi169 < 10; fi169++ {
                        delete(v2, ("" + ("foo bar" + "\n\t")))
                        _ = fi169
                    }
                default:
                    if false {
                        var v170 int = 4294967295
                        _ = v170
                    }
                }
            }
        case "a":
            v4 += 241
        case "\n\t":
            v171 := GenS23{F0: 0}
            _ = v171
        default:
            var v172 int32 = (^(v4))
            _ = v172
        }
    }
    v4 = v4
    if (!(255 <= 255)) {
        for fi173 := 0; fi173 < 1; fi173++ {
            var v174 string = "\n\t"
            for fi175 := 0; fi175 < 10; fi175++ {
                if false {
                    v4 = (v4 + v4)
                }
                _ = fi175
            }
            _ = fi173
            _ = v174
        }
    }
    var v176 float64 = (0.0 / 1e10)
    for fi177 := 0; fi177 < 2; fi177++ {
        v1 -= (v1 & (v1 >> ((931) & 63)))
        var v178 int = fi177
        _ = fi177
        _ = v178
    }
    var v179 uint8 = 2
    v3 = append(v3, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v176:", v176)
    fmt.Println("v179:", v179)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v176
    _ = v179
    _ = v2
    _ = v3
    _ = v4
}
