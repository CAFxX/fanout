package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32; F1 int8; F2 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool; F1 uint32; F2 int8 }
type GenS10 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int64 = -915
    v2 := make(map[int]int)
    v3 := []string{}
    var v4 int32 = -2147483648
    v5 := make(map[uint32]int)
    fmt.Println(v1, v4)
    var v6 uint64 = 501
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := v5[0]
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            if len(v3) > 0 {
                v3[0] = "a"
            }
            v1 += v1
            if true {
                v6 = v6
            }
            _ = fi10
        }
        _ = fi7
    }
    if false {
        v6 -= v6
        switch "" {
        case "foo bar":
            v1 -= 9223372036854775807
        case "\n\t":
            fmt.Println(v4, v4)
        default:
            var v11 int16 = 892
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 int16 = 1000
                if false {
                    var v15 int8 = 127
                    _ = v15
                } else {
                    var v16 int8 = 1
                    _ = v16
                }
                var v17 int16 = v14
                _ = fi13
                _ = v14
                _ = v17
            }
            _ = fi12
        }
    } else {
        switch "a" {
        case "a":
            v6 += v6
        case "\n\t":
            _ = len(v3)
        }
        v18 := GenS1{F0: (v1 >> ((255) & 63)), F1: 65535, F2: 1}
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        switch ("foo bar" + (("a" + "a") + "hello")) {
        case "hello":
            v3 = append(v3, "a")
        default:
            for fi20 := 0; fi20 < 16; fi20++ {
                v4 = v4
                _ = fi20
            }
        }
        if true {
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 float64 = 0.0
                _ = fi21
                _ = v22
            }
        }
        _ = fi19
    }
    switch "hello" {
    case "a":
        if false {
            v1 -= (v1 & v1)
            v4 = v4
            fmt.Println(v6, v1)
        }
    case "hello":
        switch "foo bar" {
        case "\n\t":
            if false {
                v6 = (v6 / ((v6) | 1))
                _ = len(v2)
                switch -9223372036854775808 {
                case -201:
                    _ = len(v3)
                case 128:
                    var v23 int64 = (v1 + -9223372036854775808)
                    _ = v23
                case 1:
                    if (749 <= 32767) {
                        fmt.Println(v6, v6)
                        var v24 int8 = 42
                        _ = v24
                    }
                }
            }
        case "foo bar":
            var v25 int32 = v4
            _ = v25
        default:
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS2{F0: 271, F1: 127, F2: 1000}
                _ = fi26
                _ = v27
            }
        }
    case "foo bar":
        v6 = (v6 << ((91) & 63))
    default:
        v1 -= -253
    }
    var v28 float32 = (-2.25 + 0.0)
    var v29 uint = 4294967295
    for fi30 := 0; fi30 < 17; fi30++ {
        var v31 bool = (!(!false))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi30
        _ = v31
    }
    if ((1.5 / 1.0) <= 0.0) {
        switch ("foo bar" + ("a" + "hello")) {
        case "\n\t":
            v32 := GenS2{F0: (v29 * 65535), F1: 128, F2: -346}
            _ = v32
        case "foo bar":
            if len(v3) > 0 {
                v3[0] = (("\n\t" + ("hello" + "")) + ((("hello" + "") + "foo bar") + ""))
            }
        }
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        v3 = append(v3, "hello")
    } else {
        var v33 uint64 = v6
        for fi34 := 0; fi34 < 3; fi34++ {
            v29 += (^(v29))
            for fi35 := 0; fi35 < 17; fi35++ {
                var v36 int32 = -2147483648
                var v37 uint64 = 42
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = fi34
        }
        _ = v33
    }
    if true {
        switch "foo bar" {
        case "foo bar":
            for fi38 := 0; fi38 < 10; fi38++ {
                v29 = 1
                v39 := GenS3{F0: (529 > 4294967295), F1: int8(v6)}
                _ = fi38
                _ = v39
            }
        case "\n\t":
            v29 += v29
        }
        if true {
            if (float32(0.1) != (3.14159 * (v28 / v28))) {
                var v40 float64 = (float64(float32(((1.5 / 3.14159) + 1e10))) + 1e10)
                _ = v40
            } else {
                if (false || false) {
                    v41 := GenS3{F0: (0 <= 42), F1: 42}
                    var v42 float64 = (((0.0 - (-2.25 + (1.5 / 0.1))) + 1.5) - ((float64(3.14159) - 1e10) / 3.14159))
                    if (uint32(v29) >= 2147483648) {
                        v43 := GenS3{F0: false, F1: 127}
                        var v44 uint8 = 1
                        _ = v43
                        _ = v44
                    }
                    _ = v41
                    _ = v42
                } else {
                    v45 := GenS1{F0: v1, F1: 390, F2: int16(v4)}
                    v29 = v29
                    switch 1 {
                    case 127:
                        v1 = (v1 % (((v1 - (v1 * v1))) | 1))
                    case -682:
                        for fi46 := 0; fi46 < 11; fi46++ {
                            switch "a" {
                            case "\n\t":
                                v47 := GenS4{F0: int32(fi46)}
                                _ = v47
                            case "a":
                                var v48 float32 = 0.1
                                _ = v48
                            case "":
                                v49 := v5[1]
                                _ = v49
                            default:
                                if true {
                                    switch "foo bar" {
                                    case "":
                                        if true {
                                            v29 += 949
                                        }
                                    case "hello":
                                        v50 := GenS4{F0: int32(v1)}
                                        _ = v50
                                    }
                                    if (((fi46 ^ -1) > (fi46 - 2147483647)) || (v28 == v28)) {
                                        _ = len(v3)
                                        v29 -= v29
                                        if (0.1 == float64(float32((1.5 + 1.5)))) {
                                            v51 := GenS5{F0: fi46, F1: v1}
                                            _ = v51
                                        }
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = ""
                                    }
                                }
                            }
                            _ = fi46
                        }
                    case -798:
                        var v52 int8 = -1
                        _ = v52
                    default:
                        v53 := GenS4{F0: (v4 / (((v4 % (((v4 + v4)) | 1))) | 1))}
                        _ = v53
                    }
                    _ = v45
                }
                fmt.Println(v1)
            }
            if (-314 == 2147483647) {
                var v54 int16 = -400
                fmt.Println(v6, v54)
                _ = v54
            }
            _ = len(v5)
        } else {
            if (!(-44 <= 65)) {
                for fi55 := 0; fi55 < 1; fi55++ {
                    switch fi55 {
                    case -2147483648:
                        if (v29 >= (v29 / ((42) | 1))) {
                            var v56 float64 = float64(float32(-2.25))
                            for fi57 := 0; fi57 < 13; fi57++ {
                                if false {
                                    var v58 float64 = v56
                                    var v59 float32 = (v28 - (v28 - v28))
                                    switch ("foo bar" + "foo bar") {
                                    case "":
                                        var v60 int = fi55
                                        _ = v60
                                    case "a":
                                        v59 = v28
                                    case "\n\t":
                                        v1 -= (v1 << ((v29) & 63))
                                    default:
                                        _ = len(v5)
                                    }
                                    _ = v58
                                    _ = v59
                                }
                                _ = fi57
                            }
                            switch "" {
                            case "foo bar":
                                v1 = (v1 ^ v1)
                            case "\n\t":
                                fmt.Println(v56)
                            case "a":
                                for fi61 := 0; fi61 < 17; fi61++ {
                                    var v62 uint16 = uint16(v29)
                                    fmt.Println(v1, fi61)
                                    _ = fi61
                                    _ = v62
                                }
                            default:
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            }
                            _ = v56
                        } else {
                            fmt.Println(v4)
                        }
                    default:
                        if ("hello" < ("hello" + "\n\t")) {
                            v28 = v28
                        }
                    }
                    _ = fi55
                }
                _ = len(v5)
                fmt.Println(v6)
            }
        }
        v63 := GenS1{F0: (v1 << (((v29 / (((v29 >> (((v29 * (v29 / (((^(v29))) | 1)))) & 63))) | 1))) & 63)), F1: 339, F2: 980}
        _ = v63
    }
    _ = len(v5)
    for fi64 := 0; fi64 < 6; fi64++ {
        for fi65 := 0; fi65 < 14; fi65++ {
            v66 := GenS4{F0: v4}
            _ = fi65
            _ = v66
        }
        for fi67 := 0; fi67 < 17; fi67++ {
            var v68 uint64 = v6
            v69 := GenS2{F0: v29, F1: 43, F2: 1}
            switch fi64 {
            case 1:
                switch ("a" + "\n\t") {
                case "a":
                    delete(v2, fi64)
                case "hello":
                    var v70 int = (fi67 - fi67)
                    _ = v70
                case "\n\t":
                    fmt.Println(v68, v1, v4)
                default:
                    if false {
                        switch 9223372036854775807 {
                        case 128:
                            var v71 int8 = -127
                            _ = v71
                        case 1:
                            fmt.Println(v28, v28)
                        }
                    }
                }
            }
            _ = fi67
            _ = v68
            _ = v69
        }
        _ = fi64
    }
    var v72 int64 = (v1 << (((v29 + (^(v29)))) & 63))
    if true {
        for fi73 := 0; fi73 < 14; fi73++ {
            v72 = -804
            switch 68 {
            case -1:
                for fi74 := 0; fi74 < 1; fi74++ {
                    v6 += 2
                    if (true && true) {
                        v28 -= 3.14159
                        var v75 int32 = 772
                        if ("foo bar" != "foo bar") {
                            v28 += (((((v28 / float32(3.14159)) / 1.0) + 0.1) - 3.14159) / v28)
                            v28 = v28
                        } else {
                            if (-976 >= 2147483647) {
                                if (false || true) {
                                    v29 = v29
                                    switch (fi74 + 2147483647) {
                                    case -374:
                                        switch fi73 {
                                        case -170:
                                            v76 := GenS6{F0: 560}
                                            _ = v76
                                        }
                                    default:
                                        if true {
                                            v4 = (v4 << ((v29) & 63))
                                            v77 := GenS7{F0: -2147483648, F1: 0, F2: v29}
                                            _ = v77
                                        } else {
                                            fmt.Println(fi74, v29)
                                            var v78 string = "a"
                                            v6 = 0
                                            _ = v78
                                        }
                                    }
                                    var v79 uint64 = 4294967295
                                    var v80 float32 = v28
                                    _ = v79
                                    _ = v80
                                } else {
                                    v29 = v29
                                    if true {
                                        switch fi73 {
                                        case 619:
                                            if true {
                                                fmt.Println(v1, v6)
                                                if false {
                                                    delete(v2, -594)
                                                    v81 := GenS2{F0: (v29 ^ (v29 ^ (v29 % (((^(v29))) | 1)))), F1: 42, F2: 1000}
                                                    _ = v81
                                                } else {
                                                    v82 := GenS1{F0: v1, F1: 42, F2: 1}
                                                    _ = v82
                                                }
                                                var v83 bool = false
                                                _ = v83
                                            } else {
                                                v4 = v4
                                            }
                                        case 32768:
                                            v84 := GenS2{F0: 255, F1: 127, F2: 832}
                                            _ = v84
                                        }
                                        if true {
                                            v3 = append(v3, "")
                                            if true {
                                                var v85 int16 = 32767
                                                fmt.Println(v1)
                                                if false {
                                                    fmt.Println(fi73, v28, v4)
                                                } else {
                                                    var v86 uint8 = 196
                                                    v86 = (v86 / ((47) | 1))
                                                    _ = v86
                                                }
                                                _ = v85
                                            }
                                        } else {
                                            var v87 uint8 = 128
                                            v28 = (1e10 * (float32(3.14159) * (v28 + 0.0)))
                                            v4 -= v75
                                            _ = v87
                                        }
                                        v28 = (((3.14159 - -2.25) / (v28 - float32(0.1))) / v28)
                                    }
                                    v3 = append(v3, "\n\t")
                                    v88 := GenS2{F0: v29, F1: 127, F2: 510}
                                    var v89 string = ("foo bar" + "hello")
                                    _ = v88
                                    _ = v89
                                }
                                var v90 uint8 = 2
                                v6 -= uint64(v29)
                                _ = v90
                            } else {
                                if false {
                                    fmt.Println(fi74, v29)
                                    v28 += v28
                                    switch ("a" + "\n\t") {
                                    case "":
                                        v3 = append(v3, "foo bar")
                                    case "hello":
                                        switch ("a" + "a") {
                                        case "\n\t":
                                            v91 := GenS8{F0: ""}
                                            _ = v91
                                        case "foo bar":
                                            switch fi73 {
                                            case 127:
                                                v72 += (v1 & -9223372036854775808)
                                            }
                                        }
                                    case "foo bar":
                                        v92 := GenS5{F0: fi74, F1: 42}
                                        v4 = (v4 % (((v75 ^ (v75 | -821))) | 1))
                                        v1 = v72
                                        _ = v92
                                    default:
                                        var v93 int32 = -2147483648
                                        _ = v93
                                    }
                                } else {
                                    var v94 uint8 = 255
                                    v95 := GenS7{F0: v4, F1: 0, F2: (v29 ^ (v29 / ((42) | 1)))}
                                    _ = v94
                                    _ = v95
                                }
                                fmt.Println(v6, v75, v72)
                                if (!false) {
                                    var v96 uint32 = 2147483647
                                    v4 = v4
                                    v28 = ((v28 * (1e10 / 0.1)) + v28)
                                    _ = v96
                                } else {
                                    v29 = v29
                                    var v97 uint16 = 65535
                                    _ = v97
                                }
                                var v98 string = "a"
                                var v99 int8 = 0
                                _ = v98
                                _ = v99
                            }
                            delete(v5, 2)
                            if false {
                                v100 := GenS1{F0: v1, F1: (v29 << ((v29) & 63)), F2: 42}
                                var v101 int16 = 1
                                var v102 uint16 = 2
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v100
                                _ = v101
                                _ = v102
                            }
                        }
                        _ = v75
                    } else {
                        switch "a" {
                        case "foo bar":
                            if true {
                                v4 = (v4 >> ((v29) & 63))
                                v72 -= v1
                                switch "a" {
                                case "":
                                    fmt.Println(v28, v28, v4)
                                }
                            }
                        case "a":
                            fmt.Println(v4, fi74)
                            var v103 int16 = -833
                            _ = v103
                        default:
                            delete(v5, 137)
                        }
                        v104 := GenS9{F0: (42 >= 100), F1: 0, F2: -49}
                        v6 -= v6
                        _ = v104
                    }
                    _ = fi74
                }
            }
            _ = fi73
        }
        _ = len(v2)
        var v105 uint64 = v6
        _ = v105
    } else {
        for fi106 := 0; fi106 < 3; fi106++ {
            if (true && true) {
                v107 := GenS10{F0: v6, F1: int64(v29)}
                _ = v107
            }
            fmt.Println(v6, v1)
            _ = fi106
        }
    }
    fmt.Println(v29)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v72
}
