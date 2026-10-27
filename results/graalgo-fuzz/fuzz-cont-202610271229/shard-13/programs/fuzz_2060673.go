package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int8; F2 int }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 int64; F1 uint16; F2 string }
type GenS6 struct { F0 uint; F1 int32; F2 int64 }
type GenS7 struct { F0 int16; F1 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 uint32; F2 int16 }
type GenS10 struct { F0 uint32; F1 uint64 }
type GenS11 struct { F0 string; F1 int32; F2 int8 }
type GenS12 struct { F0 uint64; F1 uint64; F2 int }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 uint }
type GenS15 struct { F0 int32; F1 uint8 }
type GenS16 struct { F0 int32; F1 bool }
type GenS17 struct { F0 uint; F1 int8; F2 uint8 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint32; F1 int16; F2 string }
func genFn0(p0 int8) uint8 {
    fmt.Println(p0)
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{2147483647, (fi6 | 32767), (fi6 * fi6)}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    return 127
}
func genFn1(p0 int, p1 int16) uint8 {
    switch "" {
    case "foo bar":
        if ((2 < 18446744073709551615) && (!(32768 <= 1))) {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: true, F1: p1}
                p0 = (p0 & fi8)
                _ = fi8
                _ = v9
            }
        } else {
            if true {
                if true {
                    switch (p0 / ((p0) | 1)) {
                    case -9223372036854775808:
                        var v10 string = "a"
                        _ = v10
                    case 2147483647:
                        for fi11 := 0; fi11 < 1; fi11++ {
                            v12 := GenS1{F0: (602 != 528), F1: p1}
                            for fi13 := 0; fi13 < 19; fi13++ {
                                fmt.Println(fi11, fi13)
                                p0 = fi13
                                _ = fi13
                            }
                            _ = fi11
                            _ = v12
                        }
                    case 528:
                        var v14 uint = 65535
                        _ = v14
                    }
                    p0 = -357
                    v15 := GenS2{F0: 42, F1: 65535, F2: 660}
                    _ = v15
                } else {
                    fmt.Println(p1, p0)
                }
            } else {
                v16 := GenS2{F0: -1, F1: 555, F2: 373}
                var v17 int64 = 1
                _ = v16
                _ = v17
            }
        }
    case "\n\t":
        p0 += p0
    }
    v18 := GenS2{F0: -37, F1: 0, F2: 2147483648}
    v19 := []int{(-(p0))}
    if true {
        p0 += p0
    } else {
        v19 = append(v19, p0)
        p0 = 0
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if true {
            var v21 bool = false
            var v22 float64 = 1.5
            for fi23 := 0; fi23 < 14; fi23++ {
                if true {
                    v19 = append(v19, (p0 % ((p0) | 1)))
                    fmt.Println(p0, fi20)
                }
                var v24 int32 = 1
                fmt.Println(fi20, fi23)
                _ = fi23
                _ = v24
            }
            _ = v21
            _ = v22
        }
        v25 := GenS2{F0: -188, F1: 539, F2: 2}
        if false {
            var v26 int64 = 0
            _ = v26
        }
        _ = fi20
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 255
}

func main() {
    v1 := []int{-104, 32768, 128}
    var v2 int32 = 42
    var v3 uint64 = 18446744073709551615
    v4 := make(map[string]uint32)
    var v5 uint32 = 0
    _ = len(v4)
    v27 := GenS1{F0: ((-(v2)) == 878), F1: -736}
    _ = len(v4)
    v2 -= v2
    for fi28 := 0; fi28 < 6; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 float64 = (3.14159 / 3.14159)
            _ = fi29
            _ = v30
        }
        v5 -= (v5 * (v5 - (v5 >> ((0) & 63))))
        switch "\n\t" {
        case "a":
            v31 := v4[(("" + ("hello" + "a")) + "foo bar")]
            _ = v31
        }
        _ = fi28
    }
    switch ("hello" + ("" + "foo bar")) {
    case "hello":
        if (true || (681 == 0)) {
            if (1 <= (^(v3))) {
                if ((((1.5 / (1e10 + 3.14159)) - 0.0) + (((1.5 - -2.25) * float64((float32(float64(float32(float64((1.5 * 0.1))))) + (((0.1 + 0.0) / (1e10 / (0.1 - (((((0.1 * 0.0) * 3.14159) / 1.5) / ((float32(0.0) - float32(-2.25)) - (-2.25 - float32(0.1)))) + 1.5)))) + -2.25)))) * 3.14159)) != (-2.25 - 1.5)) {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v33 := v4[("hello" + "hello")]
                        v5 = v33
                        if false {
                            v34 := v4[(("hello" + "foo bar") + "hello")]
                            var v35 int = (fi32 - 32767)
                            _ = v34
                            _ = v35
                        } else {
                            switch fi32 {
                            case 9223372036854775807:
                                v3 = (v3 >> ((143) & 63))
                            case 127:
                                var v36 uint32 = (^(v5))
                                _ = v36
                            case 0:
                                var v37 float32 = 1e10
                                _ = v37
                            default:
                                v38 := GenS2{F0: 2147483647, F1: 255, F2: 293}
                                _ = v38
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                    v39 := GenS1{F0: true, F1: 42}
                    var v40 int8 = 87
                    _ = v39
                    _ = v40
                }
                for fi41 := 0; fi41 < 11; fi41++ {
                    if ((("a" + "a") != "foo bar") && true) {
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 128
                            }
                            var v42 int8 = -36
                            _ = v42
                        }
                    } else {
                        v43 := GenS3{F0: (v5 % ((2147483647) | 1)), F1: 100, F2: (fi41 / (((fi41 / ((127) | 1))) | 1))}
                        _ = v43
                    }
                    fmt.Println(v2)
                    _ = fi41
                }
            } else {
                for fi44 := 0; fi44 < 20; fi44++ {
                    v45 := v4["a"]
                    _ = fi44
                    _ = v45
                }
            }
            for fi46 := 0; fi46 < 16; fi46++ {
                _ = len(v4)
                v2 -= (v2 % ((v2) | 1))
                _ = fi46
            }
            v47 := GenS3{F0: v5, F1: 1, F2: 1}
            _ = v47
        }
    case "":
        delete(v4, "\n\t")
    default:
        v5 = (v5 >> ((255) & 63))
    }
    v48 := GenS1{F0: ((882 == 803) || true), F1: 1}
    if false {
        v2 = 669
        v49 := GenS3{F0: v5, F1: 0, F2: -9223372036854775808}
        _ = v49
    } else {
        v2 += v2
    }
    switch ("foo bar" + "hello") {
    case "":
        for fi50 := 0; fi50 < 17; fi50++ {
            if false {
                if len(v1) > 0 {
                    v1[0] = (fi50 - (fi50 % ((fi50) | 1)))
                }
            }
            if len(v1) > 0 {
                v1[0] = (fi50 << ((42) & 63))
            }
            _ = fi50
        }
    case "hello":
        v1 = append(v1, 943)
    case "foo bar":
        var v51 string = "foo bar"
        _ = v51
    default:
        var v52 uint64 = (^(v3))
        _ = v52
    }
    if (!(v5 == (v5 - 4294967295))) {
        switch 898 {
        case 32768:
            var v53 uint16 = 781
            _ = v53
        case 1:
            v3 = v3
        case -952:
            v1 = append(v1, 127)
        default:
            var v54 float32 = -2.25
            _ = v54
        }
        for fi55 := 0; fi55 < 8; fi55++ {
            v3 = (v3 + (v3 >> ((374) & 63)))
            v2 = -665
            if true {
                var v56 uint16 = 255
                v5 -= v5
                _ = v56
            } else {
                if (v3 != v3) {
                    for fi57 := 0; fi57 < 17; fi57++ {
                        if false {
                            v58 := v4["\n\t"]
                            _ = v58
                        }
                        _ = fi57
                    }
                } else {
                    var v59 float64 = 1.5
                    if (false || (42 > 1)) {
                        if false {
                            v3 -= 42
                            v4[("a" + "foo bar")] = (^(v5))
                            switch "hello" {
                            case "":
                                v2 += v2
                            case "foo bar":
                                v3 -= (v3 << ((uint(v5)) & 63))
                            case "\n\t":
                                v5 += v5
                            }
                        }
                        var v60 int8 = 127
                        var v61 uint = uint(v3)
                        _ = v60
                        _ = v61
                    }
                    for fi62 := 0; fi62 < 3; fi62++ {
                        v4[("hello" + (("" + (("\n\t" + "a") + "")) + ""))] = v5
                        v5 -= 4294967295
                        switch "\n\t" {
                        case "a":
                            v2 = (v2 << ((65535) & 63))
                        case "foo bar":
                            var v63 string = "\n\t"
                            _ = v63
                        default:
                            v4[("a" + "hello")] = 550
                        }
                        _ = fi62
                    }
                    _ = v59
                }
            }
            _ = fi55
        }
        if (1 <= uint8(v5)) {
            var v64 uint32 = v5
            fmt.Println(v5, v2, v5)
            _ = v64
        } else {
            switch ("foo bar" + (("" + "a") + "\n\t")) {
            case "\n\t":
                if true {
                    if ("a" >= (("foo bar" + ("a" + ("hello" + ("\n\t" + (((("" + "\n\t") + "") + ("foo bar" + "a")) + "\n\t"))))) + "hello")) {
                        switch 1 {
                        case 128:
                            v5 = (v5 << ((1) & 63))
                        case 162:
                            v5 -= v5
                        default:
                            if ("hello" == "\n\t") {
                                v2 = v2
                                fmt.Println(v3, v5, v3)
                                if true {
                                    for fi65 := 0; fi65 < 13; fi65++ {
                                        for fi66 := 0; fi66 < 20; fi66++ {
                                            v3 -= (v3 / (((v3 * 4294967295)) | 1))
                                            if ((((0.0 + (0.0 / 3.14159)) * 1.5) / 1.5) >= 1e10) {
                                                v4[("\n\t" + "")] = 1
                                            }
                                            _ = fi66
                                        }
                                        fmt.Println(v3)
                                        _ = fi65
                                    }
                                    for fi67 := 0; fi67 < 18; fi67++ {
                                        v5 += (v5 ^ 1)
                                        var v68 float32 = (1e10 + (0.1 * -2.25))
                                        _ = fi67
                                        _ = v68
                                    }
                                }
                            } else {
                                switch 116 {
                                case 873:
                                    switch "foo bar" {
                                    case "a":
                                        v69 := GenS4{F0: 127, F1: "\n\t"}
                                        _ = v69
                                    case "hello":
                                        fmt.Println(v5, v5)
                                    }
                                case 32768:
                                    fmt.Println(v3, v2)
                                default:
                                    if len(v1) > 0 {
                                        v1[0] = -9223372036854775808
                                    }
                                }
                                if (42 >= -523) {
                                    v5 = (v5 | (v5 & v5))
                                    for fi70 := 0; fi70 < 16; fi70++ {
                                        if ((-2.25 > 1.5) && true) {
                                            v5 = v5
                                            v5 = (v5 / ((0) | 1))
                                        }
                                        _ = len(v4)
                                        _ = fi70
                                    }
                                } else {
                                    v2 -= v2
                                    if (true && true) {
                                        switch ("" + (("a" + (("\n\t" + "hello") + "hello")) + "hello")) {
                                        case "foo bar":
                                            v1 = append(v1, 42)
                                        default:
                                            if true {
                                                v71 := GenS2{F0: -810, F1: 42, F2: (^(v5))}
                                                _ = v71
                                            }
                                        }
                                        if true {
                                            delete(v4, (("hello" + "hello") + (("" + "foo bar") + "foo bar")))
                                            var v72 float64 = (0.1 / 1.0)
                                            _ = v72
                                        } else {
                                            switch 2147483647 {
                                            case 2147483647:
                                                if (-872 > 42) {
                                                    for fi73 := 0; fi73 < 9; fi73++ {
                                                        for fi74 := 0; fi74 < 9; fi74++ {
                                                            if true {
                                                                fmt.Println(fi73, fi74, fi74)
                                                                var v75 uint8 = 73
                                                                v76 := GenS2{F0: v2, F1: 439, F2: v5}
                                                                v77 := GenS5{F0: 0, F1: 255, F2: "foo bar"}
                                                                fmt.Println(v3, v75, v75)
                                                                v78 := GenS3{F0: 279, F1: 99, F2: -1}
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                                _ = v78
                                                            } else {
                                                                v1 = append(v1, 127)
                                                                if (genFn1((fi74 / ((fi74) | 1)), 932) < 128) {
                                                                    switch (fi74 & (fi74 % ((-201) | 1))) {
                                                                    case 2147483647:
                                                                        v3 -= v3
                                                                    case 9223372036854775807:
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (fi73 / (((fi73 * 250)) | 1))
                                                                        }
                                                                    case -2147483648:
                                                                        v1 = append(v1, -266)
                                                                    default:
                                                                        var v79 uint = 4294967295
                                                                        _ = v79
                                                                    }
                                                                } else {
                                                                    v4["foo bar"] = (v5 | (v5 + (v5 - (^(v5)))))
                                                                    fmt.Println(v2, v5)
                                                                    if true {
                                                                        if true {
                                                                            v2 = -2147483648
                                                                            v3 += v3
                                                                            v4["\n\t"] = (v5 / ((v5) | 1))
                                                                        } else {
                                                                            if (562 < 4294967295) {
                                                                                v3 -= (^(v3))
                                                                                _ = len(v1)
                                                                                var v80 bool = (v3 > (v3 / ((v3) | 1)))
                                                                                _ = v80
                                                                            }
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v81 int8 = 104
                                                                            v81 -= (-(v81))
                                                                            _ = v81
                                                                        }
                                                                        switch fi73 {
                                                                        case 32768:
                                                                            switch fi74 {
                                                                            case -9223372036854775808:
                                                                                v82 := GenS2{F0: v2, F1: 188, F2: v5}
                                                                                _ = v82
                                                                            }
                                                                            var v83 string = ("" + (("foo bar" + "") + ("a" + "foo bar")))
                                                                            _ = v83
                                                                        case 0:
                                                                            v2 += (v2 / (((v2 - -2147483648)) | 1))
                                                                        }
                                                                    }
                                                                    switch ("hello" + "") {
                                                                    case "a":
                                                                        v5 = v5
                                                                    case "foo bar":
                                                                        if (v2 > v2) {
                                                                            fmt.Println(fi74, v5, v2)
                                                                        } else {
                                                                            switch fi73 {
                                                                            case 2147483647:
                                                                                var v84 int16 = 229
                                                                                _ = v84
                                                                            }
                                                                            if (false || (98 > 128)) {
                                                                                if (!false) {
                                                                                    v5 = (v5 % ((v5) | 1))
                                                                                    v2 += v2
                                                                                } else {
                                                                                    v85 := GenS1{F0: (!false), F1: 611}
                                                                                    _ = v85
                                                                                }
                                                                                v4["hello"] = v5
                                                                            }
                                                                        }
                                                                    case "":
                                                                        delete(v4, ("\n\t" + (("a" + (("" + (("hello" + "") + "\n\t")) + "a")) + ("" + ""))))
                                                                    default:
                                                                        v86 := v4[("hello" + "hello")]
                                                                        _ = v86
                                                                    }
                                                                    var v87 int16 = 515
                                                                    v88 := GenS5{F0: 1, F1: 255, F2: ""}
                                                                    switch fi74 {
                                                                    case 1:
                                                                        if true {
                                                                            v89 := GenS3{F0: (v5 + v5), F1: 74, F2: (fi74 >> ((651) & 63))}
                                                                            var v90 int64 = 523
                                                                            _ = v89
                                                                            _ = v90
                                                                        }
                                                                    default:
                                                                        var v91 uint32 = (v5 ^ (^(v5)))
                                                                        _ = v91
                                                                    }
                                                                    var v92 uint64 = (v3 * (^(v3)))
                                                                    if (((v5 & v5) == v5) && ((v5 << ((uint(v92)) & 63)) < (v5 % ((42) | 1)))) {
                                                                        v93 := GenS5{F0: -9223372036854775808, F1: 2, F2: "\n\t"}
                                                                        _ = v93
                                                                    } else {
                                                                        v87 = (v87 / (((-(v87))) | 1))
                                                                        switch ("a" + "foo bar") {
                                                                        case "hello":
                                                                            v5 = v5
                                                                            var v94 float64 = (((1e10 + 3.14159) * (((((0.0 / -2.25) * (3.14159 / (1.5 / ((-2.25 * (0.1 / 1e10)) / 1e10)))) * 0.0) + 1e10) / 1.0)) / ((1e10 * (-2.25 + (0.0 * 1e10))) / 0.1))
                                                                            _ = v94
                                                                        case "\n\t":
                                                                            if false {
                                                                                _ = len(v1)
                                                                            } else {
                                                                                v87 = v87
                                                                            }
                                                                        case "foo bar":
                                                                            if false {
                                                                                if (false || (127 < 1)) {
                                                                                    if (!(-2.25 < -2.25)) {
                                                                                        v95 := v4["hello"]
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            fmt.Println(v87, fi73, v2)
                                                                                        case "":
                                                                                            var v96 uint64 = v92
                                                                                            v3 -= v96
                                                                                            if false {
                                                                                                if true {
                                                                                                    var v97 uint64 = (v3 | v92)
                                                                                                    _ = v97
                                                                                                }
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v98 string = "a"
                                                                                                    _ = v98
                                                                                                case "\n\t":
                                                                                                    v3 = 18446744073709551615
                                                                                                default:
                                                                                                    var v99 uint32 = (v95 * (v95 | v5))
                                                                                                    _ = v99
                                                                                                }
                                                                                            }
                                                                                            _ = v96
                                                                                        case "foo bar":
                                                                                            if (374 > 2) {
                                                                                                v2 = v2
                                                                                                var v100 float64 = 0.1
                                                                                                var v101 uint64 = v92
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            }
                                                                                        default:
                                                                                            var v102 uint64 = 4294967295
                                                                                            if false {
                                                                                                v103 := GenS2{F0: 909, F1: 0, F2: (v95 + (v5 % ((v95) | 1)))}
                                                                                                v104 := GenS6{F0: 4294967295, F1: 1, F2: 1}
                                                                                                _ = v103
                                                                                                _ = v104
                                                                                            }
                                                                                            _ = v102
                                                                                        }
                                                                                        var v105 uint8 = 60
                                                                                        var v106 uint64 = (v3 & (v92 * 4294967295))
                                                                                        fmt.Println(v3, v3, v87)
                                                                                        _ = v105
                                                                                        _ = v106
                                                                                        _ = v95
                                                                                    } else {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        if true {
                                                                                            switch ("foo bar" + "hello") {
                                                                                            case "a":
                                                                                                var v107 uint64 = v3
                                                                                                _ = v107
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v92, v2)
                                                                                        }
                                                                                        switch 2147483647 {
                                                                                        case 9223372036854775807:
                                                                                            v92 = 1
                                                                                        case 127:
                                                                                            var v108 int64 = 42
                                                                                            _ = v108
                                                                                        default:
                                                                                            v109 := GenS3{F0: 1, F1: 127, F2: 0}
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    _ = len(v1)
                                                                                    v110 := GenS7{F0: v87, F1: 111}
                                                                                    _ = v110
                                                                                }
                                                                                v87 = (v87 + (v87 % ((v87) | 1)))
                                                                                v87 -= v87
                                                                            } else {
                                                                                var v111 bool = false
                                                                                if false {
                                                                                    v2 += (v2 + (v2 << ((255) & 63)))
                                                                                } else {
                                                                                    v111 = (v5 >= v5)
                                                                                    switch (fi74 >> ((1) & 63)) {
                                                                                    case 0:
                                                                                        var v112 int8 = 0
                                                                                        _ = v112
                                                                                    }
                                                                                    var v113 uint16 = 0
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            v87 = -15
                                                                                            if v111 {
                                                                                                v114 := GenS8{F0: v87}
                                                                                                v3 = (^(v92))
                                                                                                switch ("foo bar" + (("a" + "foo bar") + "hello")) {
                                                                                                case "foo bar":
                                                                                                    var v115 float32 = -2.25
                                                                                                    _ = v115
                                                                                                case "hello":
                                                                                                    var v116 string = "\n\t"
                                                                                                    _ = v116
                                                                                                case "\n\t":
                                                                                                    v2 += v2
                                                                                                default:
                                                                                                    v117 := GenS8{F0: -1}
                                                                                                    _ = v117
                                                                                                }
                                                                                                _ = v114
                                                                                            } else {
                                                                                                v113 -= v113
                                                                                                v5 = (v5 & v5)
                                                                                                if (!false) {
                                                                                                    v87 += (v87 << ((65535) & 63))
                                                                                                    var v118 int16 = v87
                                                                                                    switch fi73 {
                                                                                                    case 795:
                                                                                                        switch (fi73 * -927) {
                                                                                                        case 4294967295:
                                                                                                            v1 = append(v1, fi74)
                                                                                                        case 9223372036854775807:
                                                                                                            var v119 uint16 = v113
                                                                                                            _ = v119
                                                                                                        default:
                                                                                                            v120 := GenS2{F0: (^(v2)), F1: v113, F2: (v5 & v5)}
                                                                                                            _ = v120
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v118
                                                                                                }
                                                                                            }
                                                                                            v87 -= (v87 ^ (v87 % (((v87 * -32768)) | 1)))
                                                                                        } else {
                                                                                            var v121 int8 = -1
                                                                                            v122 := GenS9{F0: (917 == 135), F1: 65535, F2: (v87 + (v87 + -949))}
                                                                                            _ = v121
                                                                                            _ = v122
                                                                                        }
                                                                                    case "hello":
                                                                                        var v123 int32 = (v2 ^ 203)
                                                                                        _ = v123
                                                                                    case "a":
                                                                                        if (18446744073709551615 > v3) {
                                                                                            _ = len(v4)
                                                                                            var v124 bool = v111
                                                                                            v92 += (v3 - v3)
                                                                                            _ = v124
                                                                                        } else {
                                                                                            v125 := GenS5{F0: -9223372036854775808, F1: v113, F2: "\n\t"}
                                                                                            v1 = append(v1, (fi74 >> ((18446744073709551615) & 63)))
                                                                                            if v111 {
                                                                                                v111 = (948 < 168)
                                                                                                if (!v111) {
                                                                                                    if v111 {
                                                                                                        v111 = true
                                                                                                        v126 := v4["\n\t"]
                                                                                                        _ = v126
                                                                                                    }
                                                                                                    var v127 uint16 = (v113 * v113)
                                                                                                    _ = v127
                                                                                                } else {
                                                                                                    if v111 {
                                                                                                        v1 = append(v1, fi74)
                                                                                                        var v128 bool = v111
                                                                                                        v2 -= (v2 % ((-1) | 1))
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            v3 = 0
                                                                                                        case "hello":
                                                                                                            var v129 int32 = (v2 - 0)
                                                                                                            _ = v129
                                                                                                        }
                                                                                                        switch fi74 {
                                                                                                        case 42:
                                                                                                            if v111 {
                                                                                                                switch "foo bar" {
                                                                                                                case "":
                                                                                                                    var v130 int8 = -1
                                                                                                                    _ = v130
                                                                                                                default:
                                                                                                                    var v131 float64 = 0.1
                                                                                                                    _ = v131
                                                                                                                }
                                                                                                                var v132 uint32 = (v5 >> ((86) & 63))
                                                                                                                fmt.Println(fi74)
                                                                                                                _ = v132
                                                                                                            }
                                                                                                        }
                                                                                                        if v128 {
                                                                                                            var v133 int16 = -65
                                                                                                            v3 -= 2
                                                                                                            _ = v133
                                                                                                        }
                                                                                                        _ = len(v4)
                                                                                                        _ = v128
                                                                                                    } else {
                                                                                                        var v134 uint32 = (^(v5))
                                                                                                        v5 = v5
                                                                                                        if (1 > 928) {
                                                                                                            v113 = v113
                                                                                                        }
                                                                                                        var v135 float64 = 0.1
                                                                                                        _ = v134
                                                                                                        _ = v135
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    if true {
                                                                                                        v2 = v2
                                                                                                        fmt.Println(v2, v87, v111)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                if (v113 > v113) {
                                                                                                    v136 := GenS10{F0: (v5 ^ v5), F1: v92}
                                                                                                    _ = v136
                                                                                                } else {
                                                                                                    if (v111 || v111) {
                                                                                                        var v137 uint8 = genFn0(1)
                                                                                                        v92 += uint64(v113)
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            if v111 {
                                                                                                                var v138 string = "hello"
                                                                                                                var v139 string = "foo bar"
                                                                                                                _ = v138
                                                                                                                _ = v139
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v140 := GenS11{F0: "foo bar", F1: (v2 * (v2 - v2)), F2: 127}
                                                                                                            _ = v140
                                                                                                        case "hello":
                                                                                                            var v141 float64 = float64((3.14159 / ((-2.25 + -2.25) * 3.14159)))
                                                                                                            _ = v141
                                                                                                        }
                                                                                                        v142 := GenS9{F0: ((v113 >> ((1) & 63)) >= 585), F1: 649, F2: v87}
                                                                                                        _ = v137
                                                                                                        _ = v142
                                                                                                    }
                                                                                                    switch (fi74 & 42) {
                                                                                                    case 1:
                                                                                                        var v143 bool = (v111 && (4294967295 < (^(v3))))
                                                                                                        _ = v143
                                                                                                    default:
                                                                                                        if (v5 != 2) {
                                                                                                            if (!(-923 != -9223372036854775808)) {
                                                                                                                var v144 uint32 = 65535
                                                                                                                v3 = (v92 - uint64(v5))
                                                                                                                _ = v144
                                                                                                            }
                                                                                                            fmt.Println(fi74, v113)
                                                                                                        }
                                                                                                    }
                                                                                                    var v145 bool = ((float32((0.0 - ((float64(0.1) * 1e10) + 1.5))) <= -2.25) && (!true))
                                                                                                    _ = v145
                                                                                                }
                                                                                            }
                                                                                            _ = v125
                                                                                        }
                                                                                    }
                                                                                    _ = v113
                                                                                }
                                                                                if (!(v111 && (v87 < 748))) {
                                                                                    var v146 int64 = -795
                                                                                    var v147 uint32 = 2
                                                                                    v87 -= 1000
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                }
                                                                                _ = v111
                                                                            }
                                                                        default:
                                                                            if (v5 >= 2) {
                                                                                var v148 int32 = v2
                                                                                _ = v148
                                                                            } else {
                                                                                switch fi74 {
                                                                                case 128:
                                                                                    delete(v4, "a")
                                                                                case -2147483648:
                                                                                    if false {
                                                                                        var v149 uint8 = 128
                                                                                        _ = v149
                                                                                    }
                                                                                case 738:
                                                                                    v150 := GenS3{F0: v5, F1: -1, F2: fi74}
                                                                                    _ = v150
                                                                                }
                                                                                v3 -= uint64(v5)
                                                                            }
                                                                            v4[("hello" + "")] = v5
                                                                            v92 = (v92 & (^(v92)))
                                                                        }
                                                                    }
                                                                    var v151 uint16 = 2
                                                                    _ = v151
                                                                    _ = v87
                                                                    _ = v88
                                                                    _ = v92
                                                                }
                                                                var v152 bool = false
                                                                _ = v152
                                                            }
                                                            var v153 uint8 = 255
                                                            if (((100 <= 5) || false) || false) {
                                                                v154 := GenS12{F0: 1, F1: v3, F2: 2147483647}
                                                                _ = v154
                                                            }
                                                            _ = fi74
                                                            _ = v153
                                                        }
                                                        _ = fi73
                                                    }
                                                    var v155 int8 = 100
                                                    _ = v155
                                                } else {
                                                    _ = len(v1)
                                                }
                                            case 32768:
                                                v5 += 2
                                            case -985:
                                                var v156 int32 = 2147483647
                                                _ = v156
                                            }
                                            if true {
                                                switch 1 {
                                                case 9223372036854775807:
                                                    v2 = v2
                                                case -139:
                                                    v157 := v4[("foo bar" + ("\n\t" + ("hello" + "\n\t")))]
                                                    _ = v157
                                                case 994:
                                                    if len(v1) > 0 {
                                                        v1[0] = -791
                                                    }
                                                default:
                                                    v158 := GenS13{F0: 166}
                                                    _ = v158
                                                }
                                                if (394 <= 32767) {
                                                    var v159 bool = ((((("a" + "foo bar") + "hello") + ("\n\t" + "hello")) + "") != ((("" + ("\n\t" + "")) + "a") + "hello"))
                                                    v159 = ("\n\t" >= "a")
                                                    v3 = v3
                                                    _ = v159
                                                }
                                            }
                                        }
                                    }
                                    var v160 string = "a"
                                    _ = v160
                                }
                                v2 -= (v2 << ((4294967295) & 63))
                            }
                        }
                        v161 := GenS7{F0: 32767, F1: -128}
                        _ = v161
                    } else {
                        v3 = v3
                        if true {
                            var v162 float32 = (float32((0.0 + (float64(1e10) * 1e10))) / (1.5 - -2.25))
                            v2 += (-(v2))
                            v163 := GenS13{F0: 65535}
                            _ = v162
                            _ = v163
                        } else {
                            for fi164 := 0; fi164 < 7; fi164++ {
                                for fi165 := 0; fi165 < 16; fi165++ {
                                    var v166 uint = 901
                                    v2 += (^(v2))
                                    if false {
                                        if (true && (-9223372036854775808 <= -13)) {
                                            if true {
                                                var v167 bool = (((!true) && ((^(fi165)) <= fi164)) || (v2 == 2147483647))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v168 uint64 = 4294967295
                                                _ = v167
                                                _ = v168
                                            } else {
                                                v3 = 500
                                                v166 = v166
                                                var v169 int = fi164
                                                _ = v169
                                            }
                                        }
                                        v1 = append(v1, (fi164 | fi165))
                                        var v170 bool = true
                                        _ = v170
                                    }
                                    _ = len(v1)
                                    v166 = v166
                                    v171 := GenS8{F0: -899}
                                    v172 := v4[("hello" + "\n\t")]
                                    _ = fi165
                                    _ = v166
                                    _ = v171
                                    _ = v172
                                }
                                var v173 uint = 655
                                v173 += 4294967295
                                _ = fi164
                                _ = v173
                            }
                            var v174 uint64 = v3
                            v175 := v4[("hello" + ("a" + (("\n\t" + "\n\t") + (("" + ("foo bar" + "\n\t")) + "hello"))))]
                            _ = v174
                            _ = v175
                        }
                        for fi176 := 0; fi176 < 8; fi176++ {
                            v3 -= v3
                            for fi177 := 0; fi177 < 5; fi177++ {
                                v178 := GenS7{F0: 1, F1: 0}
                                if (true && false) {
                                    if (0 > 29) {
                                        var v179 uint32 = (v5 / ((0) | 1))
                                        _ = v179
                                    }
                                    switch "" {
                                    case "a":
                                        switch "hello" {
                                        case "hello":
                                            if false {
                                                v3 += v3
                                                v3 -= v3
                                            }
                                        }
                                    case "":
                                        var v180 int64 = -9223372036854775808
                                        _ = v180
                                    case "\n\t":
                                        v5 += 516
                                    default:
                                        if (1 <= 42) {
                                            var v181 int32 = (^(v2))
                                            _ = v181
                                        } else {
                                            switch ("hello" + "") {
                                            case "\n\t":
                                                if true {
                                                    fmt.Println(v3, v5, v3)
                                                    var v182 uint32 = (v5 / (((v5 + uint32(v3))) | 1))
                                                    _ = v182
                                                }
                                            }
                                            v5 = (v5 >> ((991) & 63))
                                            v2 = 42
                                        }
                                        if true {
                                            var v183 uint64 = (v3 * uint64(v5))
                                            if true {
                                                v5 += v5
                                                if false {
                                                    _ = len(v4)
                                                    if (0.1 <= (float64(-2.25) * ((3.14159 * (1e10 - 1e10)) / 0.1))) {
                                                        if (1e10 == -2.25) {
                                                            var v184 int32 = 725
                                                            var v185 int32 = v184
                                                            _ = v184
                                                            _ = v185
                                                        } else {
                                                            if false {
                                                                var v186 int = (fi176 << ((474) & 63))
                                                                fmt.Println(v186, v183)
                                                                _ = v186
                                                            } else {
                                                                v2 = v2
                                                            }
                                                            v5 = (v5 ^ (v5 * 201))
                                                            var v187 uint16 = 2
                                                            _ = v187
                                                        }
                                                        v188 := GenS14{F0: (true && (-1 > -32768)), F1: 0}
                                                        _ = v188
                                                    }
                                                }
                                                if ((^(v5)) >= (v5 ^ (^(v5)))) {
                                                    _ = len(v1)
                                                    var v189 uint8 = 29
                                                    _ = v189
                                                }
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    v2 = (^(v2))
                                                default:
                                                    var v190 int = (-(fi177))
                                                    _ = v190
                                                }
                                                v1 = append(v1, fi177)
                                                v183 += (v183 / (((v183 / ((v183) | 1))) | 1))
                                            }
                                            _ = v183
                                        }
                                    }
                                    v3 += (v3 >> ((255) & 63))
                                }
                                switch "hello" {
                                case "a":
                                    v5 = v5
                                case "foo bar":
                                    var v191 float32 = float32(float64(0.0))
                                    _ = v191
                                case "hello":
                                    if (!true) {
                                        v3 += v3
                                        v5 = (v5 / (((^(v5))) | 1))
                                    }
                                default:
                                    var v192 int16 = 1000
                                    v3 += v3
                                    _ = v192
                                }
                                _ = fi177
                                _ = v178
                            }
                            switch -44 {
                            case 1:
                                var v193 int32 = v2
                                _ = v193
                            case 32767:
                                switch ("\n\t" + "foo bar") {
                                case "foo bar":
                                    switch 128 {
                                    case 127:
                                        v194 := GenS1{F0: (!false), F1: -542}
                                        _ = v194
                                    case 42:
                                        if true {
                                            for fi195 := 0; fi195 < 5; fi195++ {
                                                v196 := GenS11{F0: "", F1: v2, F2: int8(v3)}
                                                if ("hello" > "a") {
                                                    if false {
                                                        v3 = (v3 ^ 2)
                                                        if false {
                                                            if (0 != -391) {
                                                                switch "hello" {
                                                                case "hello":
                                                                    var v197 int64 = 466
                                                                    _ = v197
                                                                }
                                                                fmt.Println(fi176)
                                                            } else {
                                                                switch fi195 {
                                                                case -9223372036854775808:
                                                                    v3 = 18446744073709551615
                                                                case 974:
                                                                    var v198 uint32 = v5
                                                                    _ = v198
                                                                case 829:
                                                                    delete(v4, "hello")
                                                                    var v199 string = "hello"
                                                                    var v200 int8 = -128
                                                                    _ = v199
                                                                    _ = v200
                                                                default:
                                                                    if ((fi195 | (fi176 ^ (fi195 % ((fi195) | 1)))) > fi195) {
                                                                        switch (^(fi195)) {
                                                                        case 127:
                                                                            if (((uint(v3) > 2) || (-1 == 1000)) || (1 > 255)) {
                                                                                v1 = append(v1, fi195)
                                                                                v5 = (v5 * 559)
                                                                                v3 += v3
                                                                                if false {
                                                                                    var v201 int32 = v2
                                                                                    _ = v201
                                                                                }
                                                                            }
                                                                        case 451:
                                                                            if false {
                                                                                v202 := GenS15{F0: v2, F1: 15}
                                                                                var v203 int64 = 1
                                                                                _ = v202
                                                                                _ = v203
                                                                            }
                                                                        case -931:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (fi195 * fi195)
                                                                            }
                                                                        }
                                                                        v204 := GenS12{F0: (v3 ^ 718), F1: (v3 / ((v3) | 1)), F2: (fi195 + 9223372036854775807)}
                                                                        _ = v204
                                                                    }
                                                                }
                                                            }
                                                            var v205 uint64 = 456
                                                            v3 -= 155
                                                            _ = v205
                                                        } else {
                                                            var v206 uint32 = (v5 % ((65535) | 1))
                                                            var v207 uint8 = 1
                                                            fmt.Println(v2, v207)
                                                            var v208 int8 = 100
                                                            switch ("hello" + "\n\t") {
                                                            case "":
                                                                if (!true) {
                                                                    v207 -= v207
                                                                    v209 := GenS11{F0: "a", F1: 0, F2: 0}
                                                                    var v210 uint64 = (v3 & (v3 / ((v3) | 1)))
                                                                    _ = v209
                                                                    _ = v210
                                                                }
                                                            case "foo bar":
                                                                v3 = (v3 * 367)
                                                            case "\n\t":
                                                                var v211 uint16 = 65535
                                                                _ = v211
                                                            default:
                                                                _ = len(v4)
                                                            }
                                                            _ = v206
                                                            _ = v207
                                                            _ = v208
                                                        }
                                                    }
                                                } else {
                                                    v2 = (-(v2))
                                                }
                                                var v212 uint32 = (v5 & v5)
                                                _ = fi195
                                                _ = v196
                                                _ = v212
                                            }
                                        }
                                    default:
                                        switch "" {
                                        case "hello":
                                            for fi213 := 0; fi213 < 9; fi213++ {
                                                var v214 float64 = -2.25
                                                if (578 < 874) {
                                                    if true {
                                                        v5 = (v5 & v5)
                                                        var v215 int16 = -1
                                                        _ = v215
                                                    } else {
                                                        var v216 int32 = v2
                                                        v2 -= (v216 - v216)
                                                        if (!(!true)) {
                                                            v217 := GenS16{F0: v2, F1: (27 <= 1)}
                                                            v2 = 755
                                                            var v218 int32 = (-(v216))
                                                            if false {
                                                                var v219 uint = 0
                                                                _ = v219
                                                            }
                                                            var v220 uint16 = 42
                                                            _ = v217
                                                            _ = v218
                                                            _ = v220
                                                        } else {
                                                            v1 = append(v1, (fi213 << ((788) & 63)))
                                                            v221 := v4["foo bar"]
                                                            var v222 int = fi176
                                                            _ = v221
                                                            _ = v222
                                                        }
                                                        _ = v216
                                                    }
                                                }
                                                if false {
                                                    var v223 int64 = -1
                                                    if false {
                                                        v214 = (1.5 - float64(-2.25))
                                                        var v224 int32 = (v2 % (((v2 >> ((2) & 63))) | 1))
                                                        v224 = v224
                                                        v225 := v4[("foo bar" + ((("foo bar" + "") + "a") + "\n\t"))]
                                                        _ = v224
                                                        _ = v225
                                                    }
                                                    v214 += float64((0.0 - 3.14159))
                                                    _ = v223
                                                }
                                                _ = fi213
                                                _ = v214
                                            }
                                        }
                                    }
                                case "\n\t":
                                    switch "" {
                                    case "\n\t":
                                        fmt.Println(v2, v5, v2)
                                    }
                                case "":
                                    for fi226 := 0; fi226 < 14; fi226++ {
                                        if len(v1) > 0 {
                                            v1[0] = (fi176 / ((int(v3)) | 1))
                                        }
                                        _ = fi226
                                    }
                                default:
                                    if (fi176 > (fi176 % ((fi176) | 1))) {
                                        v227 := v4["\n\t"]
                                        if false {
                                            v227 += v5
                                            if (!((v5 - v227) == 1)) {
                                                var v228 uint32 = (v5 / (((v227 & 42)) | 1))
                                                _ = v228
                                            } else {
                                                var v229 float64 = (-2.25 * float64(1.5))
                                                _ = v229
                                            }
                                            for fi230 := 0; fi230 < 2; fi230++ {
                                                if (true && false) {
                                                    v2 = v2
                                                }
                                                var v231 int16 = 42
                                                _ = fi230
                                                _ = v231
                                            }
                                        } else {
                                            var v232 uint8 = 0
                                            _ = v232
                                        }
                                        v233 := GenS6{F0: 784, F1: v2, F2: 0}
                                        _ = v227
                                        _ = v233
                                    }
                                }
                            default:
                                _ = len(v4)
                            }
                            _ = fi176
                        }
                    }
                }
            default:
                v5 -= v5
            }
            v2 -= 2147483647
            for fi234 := 0; fi234 < 6; fi234++ {
                v235 := GenS17{F0: 18446744073709551615, F1: 42, F2: genFn1(42, -32768)}
                for fi236 := 0; fi236 < 5; fi236++ {
                    if ((v3 > 421) || (-9223372036854775808 < 42)) {
                        var v237 bool = (-32768 != 1000)
                        v238 := GenS18{F0: v5}
                        _ = v237
                        _ = v238
                    } else {
                        v5 = v5
                    }
                    _ = fi236
                }
                switch 32767 {
                case 50:
                    v3 = (v3 % ((v3) | 1))
                default:
                    for fi239 := 0; fi239 < 9; fi239++ {
                        fmt.Println(fi239, fi239, v3)
                        v240 := GenS10{F0: v5, F1: 370}
                        _ = fi239
                        _ = v240
                    }
                }
                _ = fi234
                _ = v235
            }
        }
    } else {
        fmt.Println(v3, v5, v5)
        if (0 == -1) {
            if (-32768 <= 1000) {
                var v241 uint = 4294967295
                _ = v241
            }
            var v242 int32 = (^(v2))
            v5 = (^(v5))
            _ = v242
        } else {
            fmt.Println(v2, v2, v5)
            for fi243 := 0; fi243 < 10; fi243++ {
                switch int(v5) {
                case -484:
                    v244 := GenS11{F0: "", F1: 42, F2: 100}
                    _ = v244
                default:
                    v245 := GenS19{F0: 42, F1: 32767, F2: "hello"}
                    _ = v245
                }
                var v246 float64 = 0.1
                for fi247 := 0; fi247 < 19; fi247++ {
                    var v248 uint8 = 1
                    v2 -= (v2 - v2)
                    _ = len(v4)
                    _ = fi247
                    _ = v248
                }
                _ = fi243
                _ = v246
            }
        }
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v48
    _ = v5
}
