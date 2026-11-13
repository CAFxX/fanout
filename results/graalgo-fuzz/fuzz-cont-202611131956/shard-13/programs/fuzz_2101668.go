package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int32; F2 int64 }
type GenS4 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint16 }
type GenS10 struct { F0 int64; F1 int32; F2 uint }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32; F1 int32 }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 int8; F1 int32; F2 uint32 }
type GenS18 struct { F0 string }
type GenS19 struct { F0 int16 }
func genFn0(p0 uint32) bool {
    if false {
        if false {
            v5 := []int{-1, 210, 32768}
            v6 := GenS1{F0: 876, F1: 42}
            _ = v5
            _ = v6
        } else {
            switch "hello" {
            case "foo bar":
                for fi7 := 0; fi7 < 17; fi7++ {
                    v8 := []int{fi7, -39}
                    _ = fi7
                    _ = v8
                }
            case "\n\t":
                v9 := GenS1{F0: 527, F1: (p0 << ((1) & 63))}
                _ = v9
            }
            if ((0 < 1) && (!(768 > 65535))) {
                if (-2.25 == -2.25) {
                    for fi10 := 0; fi10 < 6; fi10++ {
                        v11 := GenS1{F0: 42, F1: 2}
                        var v12 uint64 = 18446744073709551615
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    if ((("a" + (("a" + "a") + "")) >= "hello") || true) {
                        var v13 uint64 = 0
                        var v14 uint64 = v13
                        var v15 bool = true
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if false {
                            p0 = p0
                            v16 := []int{-2147483648, -724}
                            fmt.Println(p0)
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 20; fi17++ {
                            for fi18 := 0; fi18 < 17; fi18++ {
                                v19 := []int{}
                                p0 -= (p0 / ((p0) | 1))
                                _ = fi18
                                _ = v19
                            }
                            switch "a" {
                            case "":
                                v20 := GenS1{F0: 966, F1: p0}
                                _ = v20
                            case "foo bar":
                                v21 := GenS1{F0: 42, F1: (p0 + (^(p0)))}
                                _ = v21
                            }
                            fmt.Println(p0, p0)
                            _ = fi17
                        }
                    }
                }
                v22 := []int{127, 0, 122}
                var v23 uint = 4294967295
                _ = v22
                _ = v23
            } else {
                var v24 uint16 = 255
                if true {
                    v25 := GenS2{F0: -131}
                    _ = v25
                } else {
                    var v26 uint8 = 147
                    _ = v26
                }
                _ = v24
            }
        }
    } else {
        var v27 string = "foo bar"
        v28 := []int{0, 9223372036854775807, 32768, 0}
        _ = v27
        _ = v28
    }
    var v29 uint32 = (p0 / ((85) | 1))
    var v30 int64 = -785
    v31 := GenS3{F0: 2147483647, F1: 2147483647, F2: 0}
    for fi32 := 0; fi32 < 17; fi32++ {
        switch "hello" {
        case "a":
            for fi33 := 0; fi33 < 8; fi33++ {
                v30 = (v30 + 42)
                p0 = (v29 << ((255) & 63))
                _ = fi33
            }
        case "":
            v30 = (-(v30))
        default:
            var v34 int8 = 1
            _ = v34
        }
        fmt.Println(v29, v29, v30)
        if false {
            if ((42 == 2) && false) {
                fmt.Println(v30, p0, fi32)
            }
        }
        _ = fi32
    }
    _ = p0
    _ = v29
    _ = v30
    _ = v31
    return false
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 uint = 0
    var v4 int8 = 32
    v35 := []int{1, 32768, -9223372036854775808, -573}
    v36 := GenS4{F0: 42, F1: 2147483647, F2: 2147483648}
    v3 -= v3
    v37 := GenS5{F0: int32(v3), F1: (("a" + "foo bar") + (((("foo bar" + "a") + "\n\t") + ("" + "a")) + "a")), F2: 860}
    var v38 int = 2147483647
    v4 = -128
    v38 += v38
    v39 := GenS6{F0: 0}
    fmt.Println(v4, v3)
    if false {
        switch v38 {
        case -777:
            v38 -= 127
        default:
            var v40 bool = (v4 > v4)
            _ = v40
        }
        v38 -= (v38 ^ (^(v38)))
        var v41 int64 = -1
        _ = v41
    } else {
        v42 := GenS3{F0: v38, F1: 0, F2: 689}
        var v43 uint16 = 281
        _ = v42
        _ = v43
    }
    if (v38 > -20) {
        v38 += (v38 ^ (v38 % (((v38 & v38)) | 1)))
    }
    _ = len(v35)
    v38 -= v38
    fmt.Println(v38, v3, v4)
    if true {
        if ((!true) || true) {
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 int32 = -1
                var v46 bool = false
                _ = fi44
                _ = v45
                _ = v46
            }
            switch "foo bar" {
            case "a":
                if (true || (true && true)) {
                    fmt.Println(v3)
                    for fi47 := 0; fi47 < 6; fi47++ {
                        fmt.Println(v3, v4, v38)
                        v35 = append(v35, 711)
                        if len(v35) > 0 {
                            v35[0] = (v38 ^ v38)
                        }
                        _ = fi47
                    }
                    if false {
                        if ((v3 | v3) > v3) {
                            v48 := GenS7{F0: uint64(v3), F1: ((1e10 * 0.0) < float64((((((3.14159 / 0.1) * 1.5) - 1.5) / ((0.1 * (-2.25 + -2.25)) + 1e10)) / -2.25))), F2: 65535}
                            var v49 float64 = 3.14159
                            _ = v48
                            _ = v49
                        } else {
                            v38 = 9223372036854775807
                            if ((207 >= 65535) && true) {
                                if (204 <= -1) {
                                    for fi50 := 0; fi50 < 17; fi50++ {
                                        v35 = append(v35, (fi50 % ((v38) | 1)))
                                        if ((-837 < -1) || true) {
                                            fmt.Println(v3)
                                            if true {
                                                for fi51 := 0; fi51 < 11; fi51++ {
                                                    if false {
                                                        switch "hello" {
                                                        case "":
                                                            v4 += v4
                                                        }
                                                        v3 += v3
                                                        switch ("foo bar" + "") {
                                                        case "\n\t":
                                                            if true {
                                                                if ((128 == (fi50 * fi51)) && false) {
                                                                    var v52 uint = v3
                                                                    delete(v1, "\n\t")
                                                                    _ = v52
                                                                } else {
                                                                    if (1e10 == -2.25) {
                                                                        var v53 string = (("\n\t" + "foo bar") + "hello")
                                                                        if (0 > 2) {
                                                                            var v54 uint8 = 215
                                                                            switch -9223372036854775808 {
                                                                            case 662:
                                                                                v53 = ("\n\t" + ("\n\t" + ""))
                                                                            default:
                                                                                if false {
                                                                                    fmt.Println(v54)
                                                                                    v55 := GenS1{F0: 2, F1: 1}
                                                                                    if true {
                                                                                        fmt.Println(v3)
                                                                                        var v56 uint = v3
                                                                                        var v57 int16 = 32767
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    _ = v55
                                                                                }
                                                                                v38 += fi51
                                                                                _ = len(v35)
                                                                                v53 = "a"
                                                                            }
                                                                            var v58 string = "foo bar"
                                                                            _ = v54
                                                                            _ = v58
                                                                        }
                                                                        _ = v53
                                                                    }
                                                                }
                                                                var v59 uint = v3
                                                                _ = v59
                                                            }
                                                        case "hello":
                                                            var v60 int16 = 32767
                                                            _ = v60
                                                        default:
                                                            v61 := GenS8{F0: "a"}
                                                            _ = v61
                                                        }
                                                    }
                                                    _ = fi51
                                                }
                                                var v62 uint8 = 128
                                                _ = v62
                                            } else {
                                                v38 = 128
                                                for fi63 := 0; fi63 < 11; fi63++ {
                                                    v64 := GenS9{F0: "\n\t", F1: 255}
                                                    v3 = (v3 & (v3 % (((^(v3))) | 1)))
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            }
                                            var v65 float32 = (1e10 * (((-2.25 / (float32(0.1) - -2.25)) + 1e10) + 0.1))
                                            _ = v65
                                        }
                                        for fi66 := 0; fi66 < 15; fi66++ {
                                            var v67 int = fi50
                                            v67 += fi66
                                            _ = fi66
                                            _ = v67
                                        }
                                        _ = fi50
                                    }
                                    var v68 float64 = -2.25
                                    _ = v68
                                } else {
                                    var v69 int8 = -1
                                    _ = v69
                                }
                                if (0 != -1) {
                                    v70 := GenS10{F0: -1, F1: -2147483648, F2: 1}
                                    _ = v70
                                }
                                if ((934 != 2) || false) {
                                    var v71 string = ("" + "")
                                    _ = v71
                                } else {
                                    var v72 float64 = (1.5 + 0.1)
                                    _ = v72
                                }
                            }
                            _ = len(v35)
                        }
                        v73 := GenS11{F0: "a"}
                        for fi74 := 0; fi74 < 16; fi74++ {
                            v75 := GenS10{F0: 0, F1: 0, F2: (^(v3))}
                            _ = fi74
                            _ = v75
                        }
                        _ = v73
                    } else {
                        _ = len(v35)
                        if (false || true) {
                            v38 += 715
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        } else {
                            var v76 int8 = (v4 + (v4 + int8(v3)))
                            if (995 != -163) {
                                v77 := GenS5{F0: 798, F1: "hello", F2: 255}
                                v4 = 100
                                v78 := v1[(("hello" + "foo bar") + "a")]
                                _ = v77
                                _ = v78
                            }
                            _ = v76
                        }
                    }
                } else {
                    if (3.14159 == 0.0) {
                        if ((32767 <= 65535) || ((0 > -9223372036854775808) && false)) {
                            _ = len(v35)
                            v79 := GenS10{F0: -441, F1: -2147483648, F2: (v3 | 0)}
                            _ = v79
                        } else {
                            fmt.Println(v38, v3, v38)
                            switch int(v3) {
                            case 0:
                                v80 := GenS12{F0: 2}
                                _ = v80
                            default:
                                if true {
                                    switch (^(v38)) {
                                    case -762:
                                        switch (v38 | v38) {
                                        case 32767:
                                            fmt.Println(v38)
                                        case -16:
                                            v38 -= v38
                                        default:
                                            var v81 float64 = (3.14159 + 0.1)
                                            _ = v81
                                        }
                                    }
                                    v82 := GenS2{F0: v38}
                                    _ = v82
                                }
                            }
                        }
                        if (1 != 0) {
                            if len(v35) > 0 {
                                v35[0] = -389
                            }
                            if false {
                                switch "" {
                                case "":
                                    for fi83 := 0; fi83 < 18; fi83++ {
                                        v84 := GenS13{F0: uint16(v3)}
                                        _ = fi83
                                        _ = v84
                                    }
                                case "hello":
                                    fmt.Println(v4, v3)
                                default:
                                    v85 := GenS14{F0: 1}
                                    _ = v85
                                }
                                if true {
                                    v35 = append(v35, v38)
                                    v86 := v2["a"]
                                    _ = v86
                                }
                            }
                        }
                        v87 := GenS4{F0: (v4 - v4), F1: 4294967295, F2: 65535}
                        _ = v87
                    }
                    if true {
                        fmt.Println(v38, v3, v3)
                        switch (v38 & v38) {
                        case 42:
                            v3 = 2
                        case -1:
                            _ = len(v2)
                        default:
                            for fi88 := 0; fi88 < 8; fi88++ {
                                var v89 int64 = int64(v3)
                                switch "foo bar" {
                                case "foo bar":
                                    switch 32767 {
                                    case 4294967295:
                                        if false {
                                            if (("" + "") != "a") {
                                                var v90 float64 = (0.1 - float64(0.1))
                                                switch 594 {
                                                case -208:
                                                    var v91 uint8 = 0
                                                    _ = v91
                                                default:
                                                    for fi92 := 0; fi92 < 6; fi92++ {
                                                        var v93 uint32 = 1
                                                        _ = fi92
                                                        _ = v93
                                                    }
                                                }
                                                _ = v90
                                            } else {
                                                delete(v2, "a")
                                                var v94 int32 = 874
                                                _ = v94
                                            }
                                        } else {
                                            v95 := GenS15{F0: 856, F1: -812}
                                            v89 += -154
                                            _ = v95
                                        }
                                    case -671:
                                        v3 += 2
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            switch (fi88 | fi88) {
                                            case 32768:
                                                v3 += v3
                                            case 909:
                                                var v96 uint64 = 4294967295
                                                _ = v96
                                            case 1:
                                                for fi97 := 0; fi97 < 13; fi97++ {
                                                    v98 := GenS16{F0: (true && false), F1: fi97}
                                                    if false {
                                                        var v99 int = fi88
                                                        var v100 uint32 = 998
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                    if true {
                                                        if (0 <= 32767) {
                                                            var v101 int32 = 196
                                                            _ = len(v35)
                                                            _ = v101
                                                        } else {
                                                            var v102 int64 = v89
                                                            var v103 uint = 65535
                                                            var v104 uint = (v3 ^ 1)
                                                            fmt.Println(fi88)
                                                            v103 = (v3 % ((v3) | 1))
                                                            _ = v102
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                    }
                                                    v89 = v89
                                                    _ = fi97
                                                    _ = v98
                                                }
                                            default:
                                                switch v38 {
                                                case -2147483648:
                                                    var v105 float32 = 1.5
                                                    _ = v105
                                                case -1:
                                                    if (!(!(-2147483648 <= 171))) {
                                                        var v106 int64 = 42
                                                        fmt.Println(v38, v89, v38)
                                                        _ = v106
                                                    }
                                                case 192:
                                                    var v107 uint = (v3 / ((v3) | 1))
                                                    _ = v107
                                                default:
                                                    v4 -= (v4 | 40)
                                                }
                                            }
                                        case "hello":
                                            v108 := v2[("\n\t" + "")]
                                            _ = v108
                                        }
                                    }
                                default:
                                    var v109 int32 = 869
                                    _ = v109
                                }
                                var v110 int32 = 235
                                _ = fi88
                                _ = v110
                                _ = v89
                            }
                        }
                        switch ("a" + "foo bar") {
                        case "":
                            v111 := GenS15{F0: 16, F1: -1}
                            _ = v111
                        case "a":
                            fmt.Println(v4, v38)
                        }
                    }
                }
            }
            if len(v35) >= 2 {
                v35 = v35[1:len(v35)]
            }
        } else {
            for fi112 := 0; fi112 < 11; fi112++ {
                switch 32768 {
                case 2147483647:
                    for fi113 := 0; fi113 < 8; fi113++ {
                        v114 := GenS17{F0: 42, F1: int32(v4), F2: 2147483648}
                        _ = fi113
                        _ = v114
                    }
                case 0:
                    v3 -= 255
                case -943:
                    if true {
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        case "hello":
                            fmt.Println(v4, v38)
                        default:
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                fmt.Println(v3, v4, v4)
                            case "hello":
                                var v115 float32 = -2.25
                                _ = v115
                            case "foo bar":
                                for fi116 := 0; fi116 < 8; fi116++ {
                                    var v117 int16 = 0
                                    if (true || false) {
                                        var v118 uint32 = 656
                                        var v119 string = "hello"
                                        _ = v118
                                        _ = v119
                                    }
                                    _ = fi116
                                    _ = v117
                                }
                            }
                        }
                        for fi120 := 0; fi120 < 10; fi120++ {
                            fmt.Println(fi120, fi120)
                            _ = fi120
                        }
                    } else {
                        v38 = (v38 / ((4294967295) | 1))
                    }
                }
                if len(v35) > 0 {
                    v35[0] = (fi112 + v38)
                }
                _ = fi112
            }
            switch "" {
            case "foo bar":
                for fi121 := 0; fi121 < 17; fi121++ {
                    delete(v1, (("hello" + ("foo bar" + ((((("hello" + "a") + "") + ("foo bar" + ("\n\t" + ""))) + "") + ("" + "")))) + "a"))
                    if true {
                        if len(v35) > 0 {
                            v35[0] = fi121
                        }
                        switch v38 {
                        case -790:
                            var v122 int8 = v4
                            _ = v122
                        default:
                            v4 -= v4
                        }
                        for fi123 := 0; fi123 < 1; fi123++ {
                            if len(v35) > 0 {
                                v35[0] = v38
                            }
                            v3 = 715
                            fmt.Println(fi123)
                            _ = fi123
                        }
                    }
                    _ = fi121
                }
            case "":
                v35 = append(v35, (v38 ^ (v38 % ((v38) | 1))))
            case "\n\t":
                switch "hello" {
                case "a":
                    fmt.Println(v4, v38)
                }
            }
            if false {
                v38 += v38
                v124 := GenS13{F0: 198}
                v3 = (v3 >> (((v3 | (^(v3)))) & 63))
                _ = v124
            }
        }
        v3 -= v3
        switch v38 {
        case 9223372036854775807:
            var v125 int64 = 9223372036854775807
            _ = v125
        case 128:
            if true {
                v3 -= (v3 * v3)
                v126 := GenS9{F0: "foo bar", F1: 385}
                _ = v126
            } else {
                _ = len(v1)
                switch (^(v38)) {
                case 9223372036854775807:
                    v38 -= v38
                default:
                    var v127 uint = v3
                    _ = v127
                }
                v2["\n\t"] = 4294967295
            }
        default:
            switch "a" {
            case "foo bar":
                var v128 bool = false
                _ = v128
            case "\n\t":
                if false {
                    v129 := GenS17{F0: (v4 | v4), F1: 1, F2: 0}
                    _ = len(v35)
                    switch "\n\t" {
                    case "hello":
                        v38 -= (^(v38))
                    case "\n\t":
                        if false {
                            for fi130 := 0; fi130 < 2; fi130++ {
                                if len(v35) > 0 {
                                    v35[0] = -285
                                }
                                v4 = 100
                                _ = fi130
                            }
                        } else {
                            v131 := GenS3{F0: 32768, F1: 2147483647, F2: 1}
                            var v132 int = v38
                            _ = v131
                            _ = v132
                        }
                    }
                    _ = v129
                }
            default:
                v133 := GenS18{F0: ("a" + "foo bar")}
                _ = v133
            }
        }
    } else {
        v134 := GenS19{F0: 1000}
        _ = v134
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
