package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 int16; F2 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 int32; F2 bool }

func main() {
    var v1 uint8 = 184
    var v2 uint8 = v1
    v3 := []string{"\n\t", (("a" + "a") + ("\n\t" + "foo bar")), "foo bar", (("foo bar" + "a") + "foo bar")}
    var v4 uint32 = 2
    fmt.Println(v1, v1, v4)
    var v5 int64 = 9223372036854775807
    switch "\n\t" {
    case "hello":
        v1 += 128
    }
    if false {
        if (191 != 42) {
            for fi6 := 0; fi6 < 5; fi6++ {
                if false {
                    v3 = append(v3, "a")
                    switch fi6 {
                    case 1:
                        var v7 int64 = v5
                        _ = v7
                    case 0:
                        if false {
                            v1 = (v1 + 127)
                            for fi8 := 0; fi8 < 12; fi8++ {
                                v4 -= v4
                                var v9 int64 = v5
                                _ = fi8
                                _ = v9
                            }
                            v3 = append(v3, "foo bar")
                        } else {
                            v10 := GenS1{F0: v2, F1: 0, F2: "a"}
                            if (837 < 42) {
                                var v11 float64 = float64(1.5)
                                _ = v11
                            } else {
                                if true {
                                    for fi12 := 0; fi12 < 17; fi12++ {
                                        switch (fi12 | fi6) {
                                        case 170:
                                            v2 = 38
                                        }
                                        if ((fi12 | fi12) == (^(fi6))) {
                                            var v13 int64 = (^(v5))
                                            fmt.Println(v5, v5, v13)
                                            switch "\n\t" {
                                            case "foo bar":
                                                v13 -= v13
                                                v1 = (v1 % ((v1) | 1))
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            _ = v13
                                        } else {
                                            fmt.Println(fi12)
                                            var v14 bool = (false || false)
                                            v2 = v1
                                            _ = v14
                                        }
                                        _ = fi12
                                    }
                                    switch (fi6 / ((-817) | 1)) {
                                    case 9223372036854775807:
                                        var v15 int32 = 0
                                        _ = v15
                                    case -9223372036854775808:
                                        for fi16 := 0; fi16 < 4; fi16++ {
                                            if len(v3) > 0 {
                                                v3[0] = ("\n\t" + ((((("foo bar" + "foo bar") + "foo bar") + "") + "\n\t") + "a"))
                                            }
                                            if false {
                                                _ = len(v3)
                                                v17 := GenS2{F0: (v2 | v2), F1: int16(v2), F2: 101}
                                                _ = v17
                                            } else {
                                                v5 -= v5
                                                v18 := GenS2{F0: 1, F1: 731, F2: -532}
                                                if (1e10 < 0.0) {
                                                    var v19 uint32 = (^(v4))
                                                    _ = v19
                                                }
                                                var v20 uint64 = 1
                                                if ((("" + ("a" + (("hello" + "") + "hello"))) > "\n\t") && false) {
                                                    var v21 int16 = -982
                                                    _ = v21
                                                }
                                                _ = v18
                                                _ = v20
                                            }
                                            _ = fi16
                                        }
                                    case 32768:
                                        switch ("" + ("foo bar" + "")) {
                                        case "\n\t":
                                            for fi22 := 0; fi22 < 4; fi22++ {
                                                v23 := GenS3{F0: (!(-449 == 358)), F1: 0, F2: (fi6 | fi22)}
                                                _ = fi22
                                                _ = v23
                                            }
                                        case "":
                                            if false {
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                                var v24 int = -249
                                                var v25 int64 = 440
                                                _ = v24
                                                _ = v25
                                            }
                                        case "hello":
                                            for fi26 := 0; fi26 < 18; fi26++ {
                                                var v27 int16 = -735
                                                var v28 int64 = (v5 + v5)
                                                v4 = (^(v4))
                                                fmt.Println(v4, v5)
                                                _ = fi26
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                    }
                                }
                            }
                            v5 -= v5
                            _ = v10
                        }
                    }
                } else {
                    fmt.Println(v5)
                }
                _ = fi6
            }
            v3 = append(v3, ("" + "foo bar"))
        }
        switch 4294967295 {
        case 32767:
            var v29 uint = 712
            _ = v29
        }
        fmt.Println(v4, v2)
    } else {
        var v30 int8 = 0
        fmt.Println(v30, v5)
        v31 := GenS1{F0: (v2 ^ v2), F1: int32(v1), F2: "\n\t"}
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        var v33 int = (fi32 / (((fi32 + 128)) | 1))
        _ = fi32
        _ = v33
    }
    if ((((float64(((((((0.0 - (-2.25 / float32(1e10))) * 1e10) / 1.5) / 1.0) * (float32(((-2.25 * float64(1e10)) * -2.25)) + (0.1 / 1.0))) / (0.0 - 1e10))) + (-2.25 + (1e10 + 0.1))) / (3.14159 / float64(3.14159))) + 1.5) >= (1e10 / 1.0)) {
        if false {
            v1 = (v1 % ((142) | 1))
        }
    }
    v34 := GenS4{F0: (^(v4))}
    switch "a" {
    case "a":
        fmt.Println(v4, v1, v4)
    case "foo bar":
        for fi35 := 0; fi35 < 16; fi35++ {
            var v36 uint32 = uint32(v1)
            if (v2 <= (v2 << ((2) & 63))) {
                v36 = (v36 % (((v4 << ((2) & 63))) | 1))
                v1 -= (v2 / (((v2 << ((0) & 63))) | 1))
            } else {
                if false {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v36)
                    case "hello":
                        if true {
                            fmt.Println(v1)
                            if ((((("hello" + "hello") + "a") + "a") >= "a") || (true && ((v5 >> ((255) & 63)) <= (v5 * v5)))) {
                                switch (((("a" + "") + "") + ("hello" + "")) + "\n\t") {
                                case "a":
                                    v36 = v4
                                case "\n\t":
                                    v36 = 727
                                default:
                                    if false {
                                        var v37 string = ("hello" + ("" + "hello"))
                                        if (-77 == -128) {
                                            fmt.Println(v5, v2, v2)
                                            for fi38 := 0; fi38 < 12; fi38++ {
                                                var v39 uint32 = v4
                                                v39 -= 0
                                                v40 := GenS5{F0: -1, F1: fi35}
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                            }
                                            _ = len(v3)
                                        }
                                        if (false || (v5 > (v5 ^ 369))) {
                                            switch fi35 {
                                            case -846:
                                                _ = len(v3)
                                            }
                                        } else {
                                            fmt.Println(v4, v4)
                                            if (!((0.1 - -2.25) == 1.5)) {
                                                if true {
                                                    for fi41 := 0; fi41 < 7; fi41++ {
                                                        var v42 uint = 2
                                                        _ = fi41
                                                        _ = v42
                                                    }
                                                    v43 := GenS4{F0: v36}
                                                    if len(v3) > 0 {
                                                        v3[0] = v37
                                                    }
                                                    _ = v43
                                                }
                                                var v44 int = (fi35 >> ((65535) & 63))
                                                _ = v44
                                            }
                                        }
                                        _ = v37
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = ""
                                }
                            } else {
                                if false {
                                    v45 := GenS4{F0: (v36 / ((100) | 1))}
                                    fmt.Println(v36, v4, v4)
                                    var v46 int64 = (v5 - (v5 ^ (v5 + v5)))
                                    _ = v45
                                    _ = v46
                                } else {
                                    var v47 int8 = -1
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        switch "a" {
                                        case "a":
                                            var v49 uint = 42
                                            _ = v49
                                        case "\n\t":
                                            var v50 int8 = v47
                                            _ = v50
                                        case "":
                                            v4 += v36
                                        default:
                                            switch "\n\t" {
                                            case "\n\t":
                                                _ = len(v3)
                                            }
                                        }
                                        _ = fi48
                                    }
                                    _ = v47
                                }
                                if (550 <= 1) {
                                    var v51 int = 0
                                    _ = v51
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        switch "" {
                                        case "foo bar":
                                            for fi52 := 0; fi52 < 12; fi52++ {
                                                v3 = append(v3, "foo bar")
                                                _ = fi52
                                            }
                                        case "a":
                                            v3 = append(v3, ("foo bar" + "\n\t"))
                                        default:
                                            var v53 uint16 = 65535
                                            _ = v53
                                        }
                                    case "":
                                        if true {
                                            v2 -= v2
                                            if (false || false) {
                                                switch (fi35 - fi35) {
                                                case 4294967295:
                                                    switch "hello" {
                                                    case "":
                                                        v5 -= (v5 + (^(v5)))
                                                    case "\n\t":
                                                        for fi54 := 0; fi54 < 18; fi54++ {
                                                            var v55 uint8 = v1
                                                            fmt.Println(v1)
                                                            v56 := GenS6{F0: 100}
                                                            _ = fi54
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                    default:
                                                        v2 += v1
                                                    }
                                                case 0:
                                                    fmt.Println(v2)
                                                }
                                            }
                                            v36 += 2
                                        } else {
                                            var v57 int16 = -32768
                                            switch fi35 {
                                            case 128:
                                                for fi58 := 0; fi58 < 8; fi58++ {
                                                    if ((v2 ^ 127) <= (v2 & v2)) {
                                                        if (((v4 * (v36 % ((0) | 1))) < 736) && (518 <= 4294967295)) {
                                                            if false {
                                                                v59 := GenS1{F0: v1, F1: 2147483647, F2: "\n\t"}
                                                                _ = v59
                                                            }
                                                            var v60 int16 = (v57 * (v57 % ((v57) | 1)))
                                                            var v61 uint8 = (v2 ^ (v1 * v1))
                                                            v62 := GenS2{F0: v1, F1: 1000, F2: 2147483647}
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            v5 -= (-(v5))
                                                            var v63 int8 = -92
                                                            var v64 bool = true
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            v65 := GenS7{F0: v5, F1: "hello"}
                                                            _ = v65
                                                        }
                                                        v2 -= v1
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                    var v66 uint16 = 871
                                                    _ = fi58
                                                    _ = v66
                                                }
                                            case -1:
                                                switch (("\n\t" + ("a" + (("a" + "foo bar") + ""))) + (("" + ("foo bar" + "")) + "\n\t")) {
                                                case "a":
                                                    var v67 bool = (!true)
                                                    _ = v67
                                                case "hello":
                                                    for fi68 := 0; fi68 < 6; fi68++ {
                                                        v3 = append(v3, "foo bar")
                                                        v69 := GenS8{F0: v5}
                                                        fmt.Println(v4)
                                                        _ = fi68
                                                        _ = v69
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v36, v2, v57)
                                                default:
                                                    v36 += v4
                                                }
                                            }
                                            _ = v57
                                        }
                                    }
                                }
                                v2 = v1
                            }
                        }
                    }
                    var v70 uint = 760
                    _ = v70
                } else {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                }
                v5 = (v5 >> ((42) & 63))
            }
            _ = fi35
            _ = v36
        }
    }
    var v71 bool = (2 >= (v4 % ((v4) | 1)))
    fmt.Println(v1)
    for fi72 := 0; fi72 < 9; fi72++ {
        if true {
            v3 = append(v3, "foo bar")
            if v71 {
                switch fi72 {
                case 42:
                    for fi73 := 0; fi73 < 6; fi73++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi73
                    }
                case 4294967295:
                    v74 := GenS1{F0: 255, F1: -2147483648, F2: "foo bar"}
                    _ = v74
                }
                fmt.Println(v71)
                v4 = (v4 % (((v4 % ((uint32(v1)) | 1))) | 1))
            } else {
                fmt.Println(v2)
                var v75 uint8 = 2
                switch "" {
                case "a":
                    var v76 int8 = 42
                    _ = v76
                case "":
                    if (-391 == 0) {
                        if ((3.14159 * 1e10) >= (((-2.25 / 1e10) - (0.1 + 1.5)) / 1.0)) {
                            switch "hello" {
                            case "":
                                v77 := GenS8{F0: (v5 + (v5 + -1))}
                                _ = v77
                            case "hello":
                                var v78 float32 = (0.1 + 3.14159)
                                _ = v78
                            case "\n\t":
                                fmt.Println(v71)
                            default:
                                v75 = 0
                            }
                            var v79 uint16 = 895
                            _ = v79
                        }
                        fmt.Println(v71, v1, v5)
                    } else {
                        var v80 int = (fi72 >> ((472) & 63))
                        fmt.Println(v4, v75)
                        _ = v80
                    }
                }
                _ = v75
            }
        }
        for fi81 := 0; fi81 < 6; fi81++ {
            v82 := GenS4{F0: (^(v4))}
            _ = fi81
            _ = v82
        }
        _ = fi72
    }
    v83 := GenS9{F0: 4294967295, F1: -2147483648, F2: true}
    var v84 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v71
    _ = v83
    _ = v84
}
