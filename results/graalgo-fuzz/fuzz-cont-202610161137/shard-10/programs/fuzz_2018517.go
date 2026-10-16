package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint }
type GenS4 struct { F0 uint32; F1 int16; F2 bool }
type GenS5 struct { F0 uint32; F1 bool; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 uint32 }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int64; F1 string }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS12 struct { F0 int8; F1 string; F2 int64 }
func genFn0() string {
    var v7 uint64 = 0
    v7 -= (v7 << ((18446744073709551615) & 63))
    v7 -= (^(v7))
    var v8 uint = 2
    if true {
        var v9 uint64 = 2
        _ = v9
    } else {
        v10 := GenS1{F0: "", F1: -1, F2: -1}
        var v11 float32 = 0.0
        var v12 int = 1
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 int) uint8 {
    v13 := GenS1{F0: ("a" + ("hello" + "foo bar")), F1: 1, F2: -1}
    var v14 uint32 = 268
    for fi15 := 0; fi15 < 10; fi15++ {
        p0 = fi15
        var v16 uint = 4294967295
        var v17 uint16 = 741
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: -1, F1: -115}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = v13
    _ = v14
    return 255
}

func main() {
    var v1 int64 = 1
    var v2 float32 = float32(1e10)
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    v5 := []string{"\n\t"}
    var v6 bool = (2 > 589)
    v4["\n\t"] = 4294967295
    fmt.Println(v2, v6, v2)
    if v6 {
        if false {
            v2 -= ((v2 * v2) * 1e10)
            v2 = 3.14159
            for fi20 := 0; fi20 < 12; fi20++ {
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
                fmt.Println(v6, fi20)
                _ = fi20
            }
        }
        v1 = 0
        switch "a" {
        case "hello":
            if (!(v6 && (!v6))) {
                v3[257] = ""
                v1 += (v1 / ((v1) | 1))
                switch "foo bar" {
                case "hello":
                    switch (("a" + "hello") + "\n\t") {
                    case "a":
                        if v6 {
                            v2 -= v2
                            v2 = v2
                            _ = len(v5)
                        }
                    case "":
                        v2 = 3.14159
                    case "foo bar":
                        v21 := GenS2{F0: 999, F1: 1}
                        _ = v21
                    }
                case "":
                    switch "" {
                    case "":
                        v2 = ((v2 / -2.25) / 1.5)
                    case "hello":
                        switch -426 {
                        case 0:
                            if (128 <= genFn1(-715)) {
                                _ = len(v5)
                                var v22 int64 = (-(v1))
                                _ = v22
                            }
                        case 127:
                            _ = len(v5)
                        }
                    case "foo bar":
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if (!(0.0 > 3.14159)) {
                                if false {
                                    if (v1 < -9223372036854775808) {
                                        _ = len(v5)
                                    } else {
                                        var v24 uint16 = 65535
                                        switch "\n\t" {
                                        case "\n\t":
                                            for fi25 := 0; fi25 < 14; fi25++ {
                                                switch ("" + (("a" + "a") + "")) {
                                                case "":
                                                    if true {
                                                        fmt.Println(v1, v24)
                                                        var v26 string = ("hello" + "a")
                                                        _ = v26
                                                    }
                                                    switch ("hello" + "a") {
                                                    case "\n\t":
                                                        v6 = (("hello" <= "a") && true)
                                                    case "foo bar":
                                                        var v27 uint32 = 65535
                                                        v5 = append(v5, "foo bar")
                                                        v6 = (fi23 < fi25)
                                                        switch 3 {
                                                        case -289:
                                                            v28 := GenS3{F0: 2, F1: ((v27 + (v27 * (v27 << ((0) & 63)))) == (v27 ^ v27)), F2: 4294967295}
                                                            _ = v28
                                                        case 808:
                                                            v29 := GenS3{F0: 1, F1: false, F2: uint(v27)}
                                                            v1 = v1
                                                            _ = v29
                                                        case 32767:
                                                            switch 37 {
                                                            case 42:
                                                                if v6 {
                                                                    fmt.Println(v27)
                                                                } else {
                                                                    v5 = append(v5, "")
                                                                    var v30 bool = v6
                                                                    v31 := GenS4{F0: (v27 << ((719) & 63)), F1: -1, F2: v30}
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                            default:
                                                                v32 := GenS3{F0: 1, F1: (!true), F2: 4294967295}
                                                                _ = v32
                                                            }
                                                        }
                                                        _ = v27
                                                    case "a":
                                                        var v33 int16 = 1
                                                        _ = v33
                                                    }
                                                    switch int(v1) {
                                                    case -2147483648:
                                                        v34 := GenS5{F0: 1, F1: v6, F2: -2147483648}
                                                        _ = v34
                                                    case 852:
                                                        var v35 int16 = 616
                                                        _ = v35
                                                    }
                                                    if false {
                                                        v24 += (v24 & (v24 ^ 0))
                                                        v36 := GenS2{F0: 2147483647, F1: 100}
                                                        _ = v36
                                                    } else {
                                                        v6 = (746 > int16(v1))
                                                        var v37 int64 = (v1 << ((1) & 63))
                                                        if v6 {
                                                            var v38 float64 = (-2.25 - float64(v2))
                                                            v24 = v24
                                                            v24 = v24
                                                            var v39 uint16 = v24
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        _ = v37
                                                    }
                                                case "\n\t":
                                                    if false {
                                                        fmt.Println(fi23, v6, v1)
                                                        v5 = append(v5, "")
                                                    }
                                                case "a":
                                                    fmt.Println(v1)
                                                }
                                                v40 := GenS3{F0: 569, F1: true, F2: 4294967295}
                                                _ = fi25
                                                _ = v40
                                            }
                                        case "a":
                                            v41 := GenS6{F0: (v1 > (v1 & v1))}
                                            _ = v41
                                        case "hello":
                                            v6 = v6
                                        }
                                        _ = v24
                                    }
                                    var v42 uint32 = 42
                                    _ = v42
                                } else {
                                    if (!true) {
                                        fmt.Println(v6, v2, v6)
                                        fmt.Println(v2, v2, v1)
                                        switch 42 {
                                        case -2147483648:
                                            switch (fi23 << ((649) & 63)) {
                                            case 9223372036854775807:
                                                v6 = (!(145 != 750))
                                            }
                                        case 128:
                                            switch (^(fi23)) {
                                            case 127:
                                                var v43 int32 = 807
                                                _ = v43
                                            case 32767:
                                                _ = len(v4)
                                            case 2147483647:
                                                fmt.Println(fi23, v1)
                                            }
                                        case -775:
                                            var v44 float64 = 3.14159
                                            _ = v44
                                        }
                                    }
                                    var v45 uint8 = 255
                                    _ = v45
                                }
                            }
                            v6 = (1 >= -32768)
                            var v46 int = (fi23 * 0)
                            _ = fi23
                            _ = v46
                        }
                    default:
                        var v47 uint16 = 224
                        _ = v47
                    }
                case "foo bar":
                    if v6 {
                        if v6 {
                            switch -2147483648 {
                            case -857:
                                delete(v3, 1)
                            default:
                                if (0 != 32767) {
                                    for fi48 := 0; fi48 < 10; fi48++ {
                                        switch (-(fi48)) {
                                        case 42:
                                            for fi49 := 0; fi49 < 2; fi49++ {
                                                fmt.Println(fi49, v1, v1)
                                                v1 = (v1 * (v1 - v1))
                                                v50 := GenS3{F0: 1, F1: (42 != 801), F2: 4294967295}
                                                _ = fi49
                                                _ = v50
                                            }
                                        default:
                                            for fi51 := 0; fi51 < 11; fi51++ {
                                                var v52 uint32 = 338
                                                switch "a" {
                                                case "":
                                                    v53 := GenS2{F0: 42, F1: 0}
                                                    _ = v53
                                                case "\n\t":
                                                    v6 = v6
                                                case "hello":
                                                    fmt.Println(v1)
                                                default:
                                                    switch ("hello" + "a") {
                                                    case "a":
                                                        if ("a" != "foo bar") {
                                                            var v54 int8 = 42
                                                            if (!v6) {
                                                                v2 += (1e10 - -2.25)
                                                            } else {
                                                                _ = len(v5)
                                                                v55 := GenS2{F0: -384, F1: 1}
                                                                if v6 {
                                                                    if (!(v6 && v6)) {
                                                                        v56 := GenS6{F0: v6}
                                                                        v57 := GenS1{F0: ("hello" + "foo bar"), F1: -1, F2: 832}
                                                                        var v58 uint64 = 2
                                                                        switch (fi48 + (fi48 ^ 42)) {
                                                                        case 42:
                                                                            switch ("hello" + "foo bar") {
                                                                            case "hello":
                                                                                var v59 int32 = 2147483647
                                                                                _ = v59
                                                                            case "a":
                                                                                var v60 string = "\n\t"
                                                                                _ = v60
                                                                            }
                                                                        case 9223372036854775807:
                                                                            var v61 uint8 = 2
                                                                            _ = v61
                                                                        case 390:
                                                                            var v62 bool = v6
                                                                            _ = v62
                                                                        default:
                                                                            switch fi51 {
                                                                            case -576:
                                                                                v63 := GenS7{F0: int8(v1), F1: v52}
                                                                                _ = v63
                                                                            case 127:
                                                                                v6 = (-127 >= 49)
                                                                            default:
                                                                                var v64 int32 = 25
                                                                                _ = v64
                                                                            }
                                                                        }
                                                                        _ = v56
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                    v65 := GenS8{F0: 67, F1: (-(fi51))}
                                                                    _ = v65
                                                                } else {
                                                                    fmt.Println(v6)
                                                                }
                                                                _ = v55
                                                            }
                                                            v66 := GenS9{F0: (v1 / ((v1) | 1)), F1: "foo bar"}
                                                            _ = v54
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                                switch (fi48 >> ((65535) & 63)) {
                                                case 127:
                                                    _ = len(v4)
                                                case 128:
                                                    fmt.Println(v52, v6, v1)
                                                case -576:
                                                    switch fi51 {
                                                    case 780:
                                                        v67 := GenS10{F0: 346}
                                                        _ = v67
                                                    case 668:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    case -230:
                                                        v68 := GenS10{F0: 65535}
                                                        _ = v68
                                                    default:
                                                        var v69 uint8 = 173
                                                        _ = v69
                                                    }
                                                default:
                                                    v1 += v1
                                                }
                                                _ = fi51
                                                _ = v52
                                            }
                                        }
                                        v1 -= v1
                                        _ = fi48
                                    }
                                }
                            }
                            v1 = (v1 - 0)
                            var v70 uint32 = 2
                            _ = v70
                        }
                        if (42 >= 37) {
                            var v71 string = "foo bar"
                            for fi72 := 0; fi72 < 18; fi72++ {
                                v2 += ((v2 * v2) * 3.14159)
                                _ = fi72
                            }
                            v71 = v71
                            _ = v71
                        }
                    }
                }
            } else {
                delete(v3, 990)
                if (!v6) {
                    v73 := GenS10{F0: 32767}
                    _ = v73
                } else {
                    v4["foo bar"] = 954
                    v74 := GenS9{F0: (v1 - v1), F1: "hello"}
                    if ((v2 - ((v2 * 0.0) - float32((0.0 - 0.1)))) > (((v2 + 0.1) + (0.1 * (((v2 * v2) / ((((((0.0 - (3.14159 + 3.14159)) / -2.25) + v2) - float32((((1.5 - ((0.0 / float64(float32(0.1))) + 0.1)) + 0.1) + (-2.25 / (float64(0.0) + 0.1))))) * v2) * ((-2.25 * -2.25) - ((((v2 / float32(1e10)) * 1e10) / ((v2 + 1e10) / v2)) / 3.14159)))) + float32(1e10)))) + v2)) {
                        v6 = v6
                        v6 = (!((!((-(v1)) == (v1 + 42))) && v6))
                        v6 = v6
                    }
                    _ = v74
                }
                var v75 int64 = (v1 >> ((1) & 63))
                _ = v75
            }
        case "":
            v6 = true
        case "foo bar":
            v4[("\n\t" + ("a" + "\n\t"))] = 348
        }
    } else {
        v76 := GenS3{F0: 42, F1: (v6 || (!v6)), F2: 608}
        switch 127 {
        case 2147483647:
            var v77 int = -997
            _ = v77
        case 440:
            v1 += (v1 % ((1) | 1))
        case -617:
            v2 -= ((float32((0.1 + 0.0)) / v2) + (v2 / v2))
        }
        _ = v76
    }
    switch "" {
    case "foo bar":
        var v78 string = ""
        _ = v78
    case "hello":
        v79 := GenS11{F0: 573, F1: 102, F2: 32767}
        _ = v79
    case "":
        switch (("foo bar" + "\n\t") + "a") {
        case "hello":
            var v80 int8 = -1
            _ = v80
        case "foo bar":
            var v81 string = (("hello" + "a") + "hello")
            _ = v81
        case "":
            v2 = v2
        }
    }
    delete(v4, "")
    fmt.Println(v2)
    switch -2147483648 {
    case 32768:
        for fi82 := 0; fi82 < 7; fi82++ {
            fmt.Println(v1, fi82, fi82)
            var v83 int16 = 26
            _ = fi82
            _ = v83
        }
    case -1:
        v84 := GenS3{F0: 534, F1: (v6 && (v2 > (0.1 - v2))), F2: 331}
        _ = v84
    }
    if v6 {
        if false {
            if (435 != 2) {
                v85 := GenS1{F0: ("foo bar" + "foo bar"), F1: (v1 << ((65535) & 63)), F2: 0}
                _ = v85
            }
        } else {
            v86 := GenS7{F0: -1, F1: 4294967295}
            v2 -= (v2 + ((0.0 * float32(0.1)) / ((v2 - v2) / 1.5)))
            v1 = v1
            _ = v86
        }
        var v87 uint8 = 255
        delete(v3, 2147483647)
        _ = v87
    } else {
        for fi88 := 0; fi88 < 17; fi88++ {
            fmt.Println(fi88)
            fmt.Println(fi88, fi88)
            _ = fi88
        }
        for fi89 := 0; fi89 < 8; fi89++ {
            v90 := GenS11{F0: 42, F1: -13, F2: 679}
            switch ("a" + ("" + ("foo bar" + "foo bar"))) {
            case "a":
                v1 += -822
            default:
                if v6 {
                    if v6 {
                        v1 = v1
                        var v91 bool = (155 == 42)
                        switch fi89 {
                        case 442:
                            v2 -= (0.1 * v2)
                        case 42:
                            var v92 int = fi89
                            _ = v92
                        }
                        _ = v91
                    } else {
                        if (795 != 4294967295) {
                            if ((fi89 % (((fi89 << ((0) & 63))) | 1)) == (fi89 >> ((2) & 63))) {
                                var v93 float64 = (3.14159 + float64(-2.25))
                                switch "\n\t" {
                                case "hello":
                                    var v94 uint8 = 255
                                    _ = v94
                                case "\n\t":
                                    for fi95 := 0; fi95 < 5; fi95++ {
                                        var v96 bool = v6
                                        v6 = (((1e10 / v93) / v93) != v93)
                                        v97 := GenS12{F0: -1, F1: "foo bar", F2: v1}
                                        _ = fi95
                                        _ = v96
                                        _ = v97
                                    }
                                default:
                                    var v98 float64 = (0.0 / v93)
                                    _ = v98
                                }
                                _ = v93
                            }
                            fmt.Println(v6, fi89)
                        }
                    }
                    if ((v1 | v1) > (v1 | v1)) {
                        v2 -= ((v2 * 0.1) / v2)
                        if ((((v2 - 0.0) / float32((((3.14159 / (-2.25 * float64(1e10))) + (0.1 * 0.1)) - float64(v2)))) / ((v2 * (v2 + v2)) + v2)) >= ((v2 + ((v2 - -2.25) - (((((3.14159 * (-2.25 + v2)) + v2) / (((-2.25 * float32(3.14159)) * (v2 + v2)) / (v2 - v2))) * v2) * v2))) - v2)) {
                            v6 = false
                        } else {
                            v99 := GenS1{F0: "hello", F1: v1, F2: 2147483647}
                            var v100 float64 = (1.5 - 3.14159)
                            _ = v100
                            _ = v99
                        }
                    } else {
                        switch "" {
                        case "hello":
                            v1 = v1
                        case "a":
                            _ = len(v4)
                        case "foo bar":
                            _ = len(v4)
                        }
                        if v6 {
                            v2 = (v2 * ((((v2 + ((3.14159 - (float32(3.14159) * ((1e10 - float32(-2.25)) + 0.0))) + v2)) / 0.1) * ((((1.5 + v2) - v2) / (float32(-2.25) - ((v2 * (v2 / (float32(1e10) / ((v2 * v2) * v2)))) * v2))) / 0.1)) / v2))
                        } else {
                            v2 -= ((v2 + (v2 + v2)) * float32(-2.25))
                            _ = len(v5)
                            for fi101 := 0; fi101 < 15; fi101++ {
                                v6 = false
                                switch "foo bar" {
                                case "":
                                    var v102 float64 = 1e10
                                    _ = v102
                                case "\n\t":
                                    var v103 float32 = v2
                                    _ = v103
                                case "foo bar":
                                    fmt.Println(v2, v2)
                                }
                                _ = fi101
                            }
                        }
                    }
                    fmt.Println(v2, fi89, fi89)
                } else {
                    var v104 uint16 = 2
                    if (42 != v1) {
                        fmt.Println(v1, fi89, v6)
                        fmt.Println(v2, v1)
                        var v105 int = (fi89 + -9223372036854775808)
                        _ = v105
                    }
                    switch fi89 {
                    case 2147483647:
                        v106 := GenS2{F0: 2147483647, F1: 109}
                        _ = v106
                    }
                    _ = v104
                }
            }
            _ = fi89
            _ = v90
        }
        if false {
            v107 := GenS3{F0: 560, F1: ((935 < -1) || (v1 <= (v1 / (((^(v1))) | 1)))), F2: 1}
            fmt.Println(v6)
            _ = v107
        }
    }
    v1 += v1
    if len(v5) > 0 {
        v5[0] = ("" + (("hello" + "") + "foo bar"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
