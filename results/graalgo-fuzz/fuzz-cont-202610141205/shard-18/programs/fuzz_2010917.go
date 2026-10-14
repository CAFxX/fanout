package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 string; F1 int16; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 bool; F1 int8; F2 uint64 }
type GenS8 struct { F0 int32; F1 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64; F1 uint }

func main() {
    var v1 int = -2147483648
    var v2 uint16 = 65535
    v3 := make(map[uint32]string)
    v4 := []int{128, (-(v1)), v1}
    var v5 int32 = 2147483647
    var v6 uint = 18446744073709551615
    v7 := GenS1{F0: (v6 - v6), F1: ((0.1 * ((1e10 * ((1e10 * float32(-2.25)) * (0.0 - 1e10))) - 0.1)) < 0.0)}
    var v8 int = v1
    v6 += (^(v6))
    if (v8 == (v8 & (v8 * (-(v1))))) {
        fmt.Println(v8)
        for fi9 := 0; fi9 < 11; fi9++ {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS2{F0: (v5 + 42)}
                _ = fi10
                _ = v11
            }
            if (0 >= 264) {
                var v12 float32 = float32(3.14159)
                if len(v4) > 0 {
                    v4[0] = (fi9 ^ (v1 << (((v6 - v6)) & 63)))
                }
                if true {
                    var v13 int8 = 42
                    _ = v13
                } else {
                    v14 := GenS2{F0: -2147483648}
                    if (true || (32767 <= (v1 * v1))) {
                        v12 = ((float32(0.0) + v12) - float32(0.0))
                        var v15 uint64 = 65
                        _ = v15
                    } else {
                        fmt.Println(v5, fi9)
                        var v16 int64 = 0
                        _ = v16
                    }
                    fmt.Println(v8, v5)
                    _ = v14
                }
                _ = v12
            } else {
                v1 = 572
            }
            _ = fi9
        }
        v5 = 0
    } else {
        switch "hello" {
        case "a":
            v2 -= v2
        }
        var v17 uint = (^(v6))
        switch "hello" {
        case "\n\t":
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(v5, v17, v8)
                var v19 int16 = 1000
                _ = fi18
                _ = v19
            }
        case "hello":
            v20 := GenS1{F0: 18446744073709551615, F1: false}
            _ = v20
        case "a":
            for fi21 := 0; fi21 < 15; fi21++ {
                for fi22 := 0; fi22 < 13; fi22++ {
                    if true {
                        switch "\n\t" {
                        case "":
                            if false {
                                var v23 int8 = 127
                                v24 := GenS2{F0: v5}
                                fmt.Println(v2)
                                _ = v23
                                _ = v24
                            }
                        case "\n\t":
                            var v25 uint8 = 128
                            _ = v25
                        default:
                            var v26 string = "\n\t"
                            _ = v26
                        }
                        if (!(-9223372036854775808 < (v1 + (fi22 + -9223372036854775808)))) {
                            v2 += v2
                        } else {
                            if ((v2 + 0) <= 632) {
                                v5 = v5
                            } else {
                                v4 = append(v4, (^(v1)))
                                _ = len(v3)
                            }
                        }
                        if ((!false) && (!true)) {
                            switch "" {
                            case "":
                                v6 = (v17 | (v17 | (v17 + (v17 << (((^(v17))) & 63)))))
                            case "\n\t":
                                v1 = int(v5)
                            case "hello":
                                var v27 uint32 = 2
                                _ = v27
                            default:
                                v28 := GenS3{F0: 42, F1: "foo bar"}
                                _ = v28
                            }
                        } else {
                            v8 = (v1 - fi21)
                        }
                    } else {
                        v29 := GenS4{F0: "foo bar", F1: 0, F2: 1000}
                        switch "\n\t" {
                        case "":
                            v30 := GenS3{F0: 128, F1: "\n\t"}
                            _ = v30
                        default:
                            v31 := v3[4294967295]
                            _ = v31
                        }
                        _ = len(v4)
                        _ = v29
                    }
                    _ = fi22
                }
                v1 += -707
                _ = fi21
            }
        }
        _ = v17
    }
    if false {
        fmt.Println(v8, v6)
        v5 -= (v5 & (v5 ^ (^(v5))))
        switch (^(v1)) {
        case 128:
            if false {
                var v32 int16 = 42
                fmt.Println(v2)
                if ((0 == 42) || true) {
                    if ((float64(((0.1 + 0.1) - ((0.1 - float32(float64((0.0 + (0.1 + float32(3.14159)))))) / ((-2.25 * float32(0.1)) / 0.1)))) < -2.25) && true) {
                        fmt.Println(v32)
                        switch "" {
                        case "foo bar":
                            var v33 uint16 = 65535
                            _ = v33
                        case "":
                            if false {
                                v34 := GenS5{F0: 2147483648}
                                _ = v34
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v4 = append(v4, v1)
                                case "foo bar":
                                    var v35 float32 = 0.1
                                    _ = v35
                                case "":
                                    fmt.Println(v1)
                                }
                                fmt.Println(v2, v5, v32)
                            }
                        default:
                            for fi36 := 0; fi36 < 16; fi36++ {
                                var v37 uint64 = 18446744073709551615
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                }
                _ = v32
            }
        default:
            var v38 float64 = 1.5
            _ = v38
        }
    }
    switch (v8 % ((v1) | 1)) {
    case 32768:
        for fi39 := 0; fi39 < 19; fi39++ {
            switch "a" {
            case "hello":
                v40 := GenS2{F0: -2147483648}
                _ = v40
            case "a":
                v5 = (v5 & (v5 / ((248) | 1)))
            }
            if (((!((v2 + v2) <= (v2 / (((^(v2))) | 1)))) && (int32(v6) < 1)) && false) {
                var v41 int16 = 1000
                var v42 uint8 = 2
                _ = v41
                _ = v42
            } else {
                v5 = (-(v5))
                for fi43 := 0; fi43 < 15; fi43++ {
                    fmt.Println(fi39)
                    v4 = append(v4, (fi43 ^ fi39))
                    v44 := v3[178]
                    _ = fi43
                    _ = v44
                }
            }
            _ = fi39
        }
    case -417:
        var v45 float64 = 0.1
        _ = v45
    }
    fmt.Println(v2, v6)
    switch v8 {
    case -173:
        for fi46 := 0; fi46 < 3; fi46++ {
            var v47 uint8 = 128
            v8 = int(v47)
            _ = fi46
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        if false {
            var v49 uint = v6
            _ = v49
        } else {
            if ((0.1 * (0.0 - 1e10)) < ((1.5 + 0.1) - -2.25)) {
                if len(v4) > 0 {
                    v4[0] = (v1 + v8)
                }
                for fi50 := 0; fi50 < 18; fi50++ {
                    switch (("foo bar" + "") + ("hello" + "\n\t")) {
                    case "":
                        var v51 float32 = 1e10
                        _ = v51
                    }
                    v52 := GenS6{F0: 2, F1: 42}
                    if (!false) {
                        v2 -= 255
                        var v53 uint64 = 4
                        switch (fi50 % (((fi48 % ((fi48) | 1))) | 1)) {
                        case 9223372036854775807:
                            switch "foo bar" {
                            case "foo bar":
                                v3[91] = ("a" + "\n\t")
                                var v54 uint8 = 2
                                _ = v54
                            default:
                                switch "a" {
                                case "foo bar":
                                    var v55 uint16 = 1
                                    fmt.Println(v1, v8, fi50)
                                    _ = len(v4)
                                    _ = v55
                                }
                            }
                        }
                        var v56 int8 = 1
                        _ = v53
                        _ = v56
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = fi50
                    _ = v52
                }
                if false {
                    var v57 int = 32768
                    v2 = v2
                    var v58 uint16 = v2
                    _ = v57
                    _ = v58
                }
            }
            fmt.Println(v6, fi48, v6)
        }
        switch 127 {
        case 675:
            for fi59 := 0; fi59 < 20; fi59++ {
                if false {
                    v8 += int(v2)
                    var v60 bool = true
                    if true {
                        switch fi48 {
                        case 32768:
                            v60 = (9223372036854775807 != 1)
                        default:
                            var v61 int16 = -32768
                            _ = v61
                        }
                        v6 = v6
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2)
                        case "a":
                            if (!(!(42 >= 127))) {
                                switch (("a" + "a") + ("" + ("" + (("" + (("\n\t" + "foo bar") + "hello")) + "")))) {
                                case "\n\t":
                                    fmt.Println(v8, v1, v1)
                                default:
                                    switch "foo bar" {
                                    case "":
                                        var v62 uint16 = (v2 * v2)
                                        _ = v62
                                    case "a":
                                        fmt.Println(v6, fi59)
                                    default:
                                        _ = len(v3)
                                    }
                                }
                                v63 := GenS2{F0: (v5 % ((v5) | 1))}
                                _ = v63
                            }
                        default:
                            var v64 int8 = -1
                            var v65 uint16 = (v2 & (v2 | (v2 | v2)))
                            _ = v64
                            _ = v65
                        }
                    }
                    _ = v60
                }
                v66 := GenS3{F0: 14, F1: "\n\t"}
                var v67 float32 = float32(3.14159)
                switch "\n\t" {
                case "\n\t":
                    if ((3.14159 / 1.0) <= (-2.25 + (float64((-2.25 / v67)) + (0.0 * ((-2.25 * (float64(v67) - (1e10 * (0.1 - (0.0 - float64(v67)))))) + (float64(float32(1e10)) / float64(-2.25))))))) {
                        var v68 uint = v6
                        _ = v68
                    } else {
                        switch (^(fi48)) {
                        case -1:
                            v69 := GenS2{F0: 1}
                            _ = v69
                        }
                    }
                default:
                    v2 += (v2 & v2)
                    if true {
                        v1 = (-(fi48))
                        if true {
                            var v70 int8 = 0
                            _ = v70
                        } else {
                            var v71 int32 = (v5 + (v5 / (((v5 + v5)) | 1)))
                            _ = v71
                        }
                        var v72 float64 = 3.14159
                        _ = v72
                    } else {
                        fmt.Println(v67, v1)
                        if true {
                            v73 := GenS2{F0: v5}
                            _ = v73
                        } else {
                            var v74 int64 = -9223372036854775808
                            v74 = (-(v74))
                            _ = v74
                        }
                        v8 = -686
                    }
                    v6 -= 1
                }
                _ = fi59
                _ = v66
                _ = v67
            }
        case 128:
            var v75 uint16 = v2
            _ = v75
        case -856:
            if (true || (true || false)) {
                for fi76 := 0; fi76 < 18; fi76++ {
                    delete(v3, 2147483648)
                    var v77 int16 = 452
                    _ = fi76
                    _ = v77
                }
            }
        default:
            v78 := GenS2{F0: v5}
            _ = v78
        }
        for fi79 := 0; fi79 < 11; fi79++ {
            switch "a" {
            case "a":
                fmt.Println(v5)
            case "\n\t":
                switch fi79 {
                case 0:
                    var v80 int8 = 99
                    _ = v80
                case 1:
                    var v81 float32 = (1e10 * 1e10)
                    _ = v81
                }
            }
            v8 += (v1 | (v8 ^ 1))
            _ = fi79
        }
        _ = fi48
    }
    v82 := GenS1{F0: (v6 ^ (v6 * (v6 >> ((537) & 63)))), F1: false}
    var v83 uint16 = (v2 % ((255) | 1))
    switch -9223372036854775808 {
    case 0:
        if false {
            v5 += (v5 % ((-1) | 1))
            v84 := GenS2{F0: 734}
            _ = v84
        } else {
            switch (("a" + "hello") + (("hello" + "hello") + "a")) {
            case "":
                fmt.Println(v2, v2)
            case "a":
                v85 := GenS7{F0: false, F1: 1, F2: 1}
                _ = v85
            case "hello":
                v86 := GenS1{F0: v6, F1: false}
                _ = v86
            }
            var v87 uint = (v6 + (v6 + v6))
            _ = v87
        }
    case -527:
        v5 = (v5 ^ v5)
    case 660:
        var v88 uint = (v6 << ((v6) & 63))
        _ = v88
    default:
        v83 = v2
    }
    if (false && (!true)) {
        if ((((1.5 + (0.1 - -2.25)) == 1e10) && ((v8 & (v1 << ((v6) & 63))) != (v8 % ((42) | 1)))) || (-2.25 < 0.1)) {
            v89 := GenS8{F0: (v5 + 649), F1: 42}
            _ = v89
        }
        if ((1 == 42) || true) {
            v6 = v6
        }
        var v90 bool = (true && (531 >= 2147483647))
        _ = v90
    }
    fmt.Println(v1, v83, v6)
    fmt.Println(v1, v8, v8)
    for fi91 := 0; fi91 < 1; fi91++ {
        v5 = (-(v5))
        var v92 uint = v6
        v93 := GenS2{F0: -198}
        _ = fi91
        _ = v92
        _ = v93
    }
    if (18446744073709551615 >= 0) {
        for fi94 := 0; fi94 < 2; fi94++ {
            v8 += 902
            v95 := GenS9{F0: 4294967295}
            for fi96 := 0; fi96 < 14; fi96++ {
                switch ("" + "a") {
                case "a":
                    var v97 uint64 = 825
                    _ = v97
                default:
                    if true {
                        switch (v1 << ((65535) & 63)) {
                        case 32767:
                            if false {
                                if ((0.0 == 3.14159) || (true || (86 <= 42))) {
                                    if len(v4) > 0 {
                                        v4[0] = (fi96 % ((v1) | 1))
                                    }
                                }
                            }
                            v8 = v8
                            v6 -= v6
                        default:
                            v5 = (v5 + v5)
                        }
                        var v98 int = 1
                        _ = v98
                    } else {
                        _ = len(v3)
                        var v99 uint8 = 255
                        _ = v99
                    }
                }
                var v100 float64 = ((1e10 - (3.14159 * 3.14159)) + 1e10)
                v101 := GenS10{F0: 490, F1: (^(v6))}
                _ = fi96
                _ = v100
                _ = v101
            }
            _ = fi94
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v82
    _ = v83
}
