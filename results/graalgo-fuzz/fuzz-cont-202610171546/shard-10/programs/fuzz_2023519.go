package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 string; F1 string; F2 uint16 }
type GenS3 struct { F0 int8; F1 int16; F2 uint }
type GenS4 struct { F0 uint32; F1 bool; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint16; F1 int16; F2 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int64; F1 int64; F2 bool }
type GenS11 struct { F0 uint32; F1 string; F2 uint8 }
type GenS12 struct { F0 string; F1 int16; F2 int8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    v3 := []string{"foo bar", "foo bar"}
    var v4 uint16 = 1
    switch 127 {
    case 165:
        v4 += v4
    }
    delete(v2, (("hello" + "foo bar") + (("\n\t" + "\n\t") + ("\n\t" + "a"))))
    if (false || true) {
        if false {
            if (42 <= 42) {
                fmt.Println(v4)
                var v5 uint8 = 128
                for fi6 := 0; fi6 < 15; fi6++ {
                    var v7 uint32 = 42
                    for fi8 := 0; fi8 < 18; fi8++ {
                        _ = len(v3)
                        var v9 int = (fi6 - (-(fi8)))
                        _ = fi8
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
                _ = v5
            } else {
                if len(v3) > 0 {
                    v3[0] = (("foo bar" + "a") + "hello")
                }
            }
            if len(v3) > 0 {
                v3[0] = ""
            }
        } else {
            var v10 uint64 = 4294967295
            if (false || true) {
                if true {
                    if (4294967295 < 4294967295) {
                        fmt.Println(v4, v4)
                        var v11 float64 = 1e10
                        _ = v11
                    } else {
                        var v12 int8 = 42
                        _ = v12
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v10 = v10
            }
            _ = v10
        }
        v4 = v4
        var v13 int32 = 0
        _ = v13
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS1{F0: 0, F1: 326, F2: 65535}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS2{F0: "\n\t", F1: "", F2: (v4 | v4)}
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint16 = v4
        for fi19 := 0; fi19 < 16; fi19++ {
            if len(v3) > 0 {
                v3[0] = ""
            }
            v20 := GenS3{F0: -128, F1: 1000, F2: 4294967295}
            if (true || (("\n\t" + ("" + ((("hello" + "") + "a") + "hello"))) > "\n\t")) {
                if false {
                    v4 = v18
                } else {
                    v1[(fi19 >> ((4294967295) & 63))] = fi19
                    switch fi17 {
                    case 1:
                        fmt.Println(fi19, fi17, fi17)
                    case -2147483648:
                        if (!true) {
                            v21 := GenS2{F0: ("\n\t" + "hello"), F1: "foo bar", F2: (v4 % ((1) | 1))}
                            var v22 float32 = (1e10 / 3.14159)
                            fmt.Println(v18)
                            fmt.Println(v22)
                            _ = v21
                            _ = v22
                        } else {
                            v4 -= 0
                            var v23 uint = 673
                            v24 := GenS3{F0: -47, F1: 0, F2: v23}
                            _ = v23
                            _ = v24
                        }
                    case 0:
                        if false {
                            v18 = (v18 / ((v18) | 1))
                        }
                    }
                    var v25 int = (^(fi17))
                    _ = v25
                }
            }
            if true {
                if (33 <= 42) {
                    v26 := GenS4{F0: 2, F1: false, F2: -2147483648}
                    _ = v26
                } else {
                    v27 := GenS5{F0: 2}
                    var v28 uint64 = 1
                    v18 += 255
                    _ = v27
                    _ = v28
                }
                v29 := GenS4{F0: 65535, F1: (float64(((0.0 - (-2.25 * (0.1 / 1e10))) / -2.25)) < 0.0), F2: -1}
                _ = v29
            } else {
                if ((0.0 * (0.0 * 0.1)) < (3.14159 - ((float64(1.5) + ((0.0 + 0.1) * (3.14159 * (3.14159 + -2.25)))) * 1.5))) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = len(v3)
                    if true {
                        fmt.Println(fi19, fi17)
                    } else {
                        delete(v1, (fi17 ^ (fi17 % ((9223372036854775807) | 1))))
                    }
                }
                fmt.Println(v4, v18, v4)
            }
            _ = fi19
            _ = v20
        }
        switch -666 {
        case 32768:
            if (!(42 >= 2147483648)) {
                if true {
                    var v30 string = ""
                    v4 = (v4 / (((v4 * v4)) | 1))
                    var v31 float64 = 0.0
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 4; fi32++ {
                        var v33 int8 = 35
                        v2["a"] = (fi17 | fi32)
                        _ = fi32
                        _ = v33
                    }
                }
            }
        case -1:
            v34 := GenS6{F0: 18446744073709551615, F1: 939}
            _ = v34
        }
        _ = fi17
        _ = v18
    }
    v4 = v4
    v4 = 32767
    v4 = (v4 | (v4 * v4))
    if true {
        if (0 == -1) {
            if false {
                if true {
                    v4 = v4
                    v35 := GenS3{F0: 0, F1: int16(v4), F2: 2}
                    _ = v35
                } else {
                    fmt.Println(v4)
                    fmt.Println(v4)
                    if false {
                        switch (("" + "a") + "hello") {
                        case "\n\t":
                            switch 32767 {
                            case 2147483647:
                                fmt.Println(v4)
                            default:
                                if false {
                                    var v36 uint32 = 2
                                    _ = v36
                                } else {
                                    var v37 int = -894
                                    for fi38 := 0; fi38 < 12; fi38++ {
                                        if (0 != 246) {
                                            fmt.Println(v4, fi38)
                                            if (255 == 0) {
                                                v4 = (v4 ^ v4)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v39 int8 = 127
                                                    _ = v39
                                                case "hello":
                                                    fmt.Println(v37, fi38, fi38)
                                                default:
                                                    for fi40 := 0; fi40 < 9; fi40++ {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v4 = (v4 - (^(v4)))
                                                        case "a":
                                                            v37 = (v37 + fi38)
                                                        case "":
                                                            var v41 int32 = 2147483647
                                                            v42 := GenS2{F0: ("hello" + ("a" + (("\n\t" + "a") + "a"))), F1: "a", F2: 42}
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = fi40
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "hello":
                                                    v37 = 109
                                                case "":
                                                    for fi43 := 0; fi43 < 6; fi43++ {
                                                        switch v37 {
                                                        case -54:
                                                            var v44 bool = (423 == -1)
                                                            _ = v44
                                                        case 4294967295:
                                                            if false {
                                                                if true {
                                                                    v45 := GenS2{F0: "\n\t", F1: ("foo bar" + "a"), F2: 32767}
                                                                    _ = v45
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(fi38, v4, fi43)
                                                                        var v46 int64 = 1
                                                                        v47 := GenS1{F0: -128, F1: 42, F2: 914}
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                    var v48 int8 = -1
                                                                    var v49 uint32 = 600
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                var v50 float32 = 1e10
                                                                var v51 uint64 = 4294967295
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        default:
                                                            v3 = append(v3, ("" + (("" + "hello") + "")))
                                                        }
                                                        _ = fi43
                                                    }
                                                default:
                                                    _ = len(v3)
                                                }
                                            } else {
                                                fmt.Println(fi38, v4, v37)
                                            }
                                            v52 := v1[v37]
                                            _ = v52
                                        }
                                        _ = fi38
                                    }
                                    _ = v37
                                }
                            }
                        case "":
                            switch "\n\t" {
                            case "hello":
                                if false {
                                    v53 := GenS3{F0: 1, F1: 1000, F2: 255}
                                    var v54 uint16 = (v4 * v4)
                                    _ = v53
                                    _ = v54
                                }
                            case "a":
                                v4 = v4
                            case "\n\t":
                                _ = len(v3)
                            default:
                                var v55 int = -303
                                _ = v55
                            }
                        case "foo bar":
                            v4 = v4
                        }
                        v4 = v4
                        if (v4 != (v4 % (((v4 % ((707) | 1))) | 1))) {
                            if false {
                                if true {
                                    v4 += 255
                                    if (!false) {
                                        v56 := GenS7{F0: 393, F1: 802, F2: 100}
                                        switch 128 {
                                        case 400:
                                            _ = len(v1)
                                        case 2147483647:
                                            var v57 uint8 = 1
                                            _ = v57
                                        default:
                                            var v58 bool = false
                                            _ = v58
                                        }
                                        v4 -= (^(v4))
                                        _ = v56
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            switch "hello" {
                                            case "hello":
                                                v4 -= (v4 | v4)
                                            case "\n\t":
                                                v59 := GenS8{F0: 42, F1: 398}
                                                _ = v59
                                            case "a":
                                                v60 := GenS6{F0: 18446744073709551615, F1: (v4 ^ (v4 % ((1) | 1)))}
                                                _ = v60
                                            default:
                                                var v61 int16 = 1000
                                                _ = v61
                                            }
                                        case "hello":
                                            var v62 string = ""
                                            _ = v62
                                        default:
                                            fmt.Println(v4)
                                        }
                                        v63 := GenS2{F0: "\n\t", F1: "\n\t", F2: 65535}
                                        _ = v63
                                    }
                                }
                            } else {
                                var v64 uint32 = 0
                                v64 -= (v64 * (^(v64)))
                                _ = v64
                            }
                            v4 += (v4 / ((428) | 1))
                            var v65 int32 = -200
                            _ = v65
                        } else {
                            v4 = v4
                        }
                    } else {
                        var v66 uint64 = 4294967295
                        v66 -= v66
                        _ = v66
                    }
                }
                v4 = (^(v4))
            } else {
                if false {
                    v2[("a" + "")] = -666
                    for fi67 := 0; fi67 < 3; fi67++ {
                        fmt.Println(v4)
                        v3 = append(v3, "foo bar")
                        fmt.Println(fi67, v4)
                        _ = fi67
                    }
                    for fi68 := 0; fi68 < 12; fi68++ {
                        fmt.Println(v4)
                        var v69 uint = 2
                        switch (("" + "hello") + "a") {
                        case "\n\t":
                            v4 += (v4 ^ v4)
                        case "a":
                            for fi70 := 0; fi70 < 13; fi70++ {
                                v3 = append(v3, ("foo bar" + "hello"))
                                v2[("foo bar" + "a")] = fi70
                                _ = fi70
                            }
                        case "hello":
                            v69 = (^(v69))
                        default:
                            var v71 float32 = float32(0.1)
                            _ = v71
                        }
                        _ = fi68
                        _ = v69
                    }
                }
                fmt.Println(v4)
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v4 += v4
            v72 := GenS9{F0: (!true), F1: (v4 >> ((18446744073709551615) & 63))}
            _ = v72
        }
    }
    if true {
        v4 = v4
    } else {
        var v73 int8 = 0
        if true {
            fmt.Println(v4, v73)
            var v74 int32 = 808
            if (500 >= 0) {
                if (127 > 1) {
                    for fi75 := 0; fi75 < 15; fi75++ {
                        if (!((float32(1e10) * (float32(0.1) + 0.0)) != (((1e10 + ((1.5 * (1.5 / float32(0.1))) / 1e10)) + (0.0 * 3.14159)) / (0.1 + (float32(1.5) - 1.5))))) {
                            v73 = v73
                        }
                        _ = fi75
                    }
                    _ = len(v2)
                    if true {
                        for fi76 := 0; fi76 < 17; fi76++ {
                            if false {
                                fmt.Println(v73, v4, v73)
                            }
                            switch (fi76 >> ((987) & 63)) {
                            case -628:
                                v74 += v74
                            case 2147483647:
                                _ = len(v3)
                            default:
                                if (0 < (fi76 / (((fi76 - fi76)) | 1))) {
                                    v4 -= v4
                                    fmt.Println(v74, v74)
                                    switch "hello" {
                                    case "foo bar":
                                        v74 = -1
                                    }
                                }
                            }
                            _ = fi76
                        }
                        v73 = int8(v74)
                        var v77 bool = false
                        _ = v77
                    }
                } else {
                    switch 32767 {
                    case 9223372036854775807:
                        var v78 uint64 = 18446744073709551615
                        _ = v78
                    case 42:
                        v79 := GenS4{F0: 4294967295, F1: false, F2: v74}
                        _ = v79
                    }
                    _ = len(v3)
                    v73 += v73
                }
                fmt.Println(v74, v73, v4)
                switch -685 {
                case -938:
                    fmt.Println(v73, v74)
                case -2147483648:
                    v1[-9223372036854775808] = 32768
                case -680:
                    v80 := GenS10{F0: -30, F1: 963, F2: false}
                    _ = v80
                default:
                    switch "foo bar" {
                    case "foo bar":
                        switch "foo bar" {
                        case "":
                            var v81 int8 = 1
                            _ = v81
                        default:
                            v82 := GenS10{F0: 9223372036854775807, F1: 42, F2: false}
                            _ = v82
                        }
                    case "\n\t":
                        v74 = 42
                    case "":
                        switch 2147483647 {
                        case 42:
                            for fi83 := 0; fi83 < 2; fi83++ {
                                v84 := GenS11{F0: 1, F1: ((("" + (("a" + ("hello" + "a")) + ("" + "a"))) + "foo bar") + (("\n\t" + "foo bar") + ("a" + ("foo bar" + "foo bar")))), F2: 0}
                                v73 -= 42
                                fmt.Println(fi83)
                                _ = fi83
                                _ = v84
                            }
                        case -9223372036854775808:
                            switch "a" {
                            case "":
                                if (true || false) {
                                    var v85 string = "foo bar"
                                    _ = v85
                                }
                            case "hello":
                                if true {
                                    switch ("" + ("a" + "\n\t")) {
                                    case "":
                                        v74 += v74
                                    default:
                                        v73 -= v73
                                    }
                                    _ = len(v3)
                                    var v86 int32 = (v74 & (v74 * (v74 - 42)))
                                    _ = v86
                                }
                            case "foo bar":
                                if (141 >= 35) {
                                    for fi87 := 0; fi87 < 11; fi87++ {
                                        switch ((("a" + "a") + "a") + ("a" + "a")) {
                                        case "hello":
                                            var v88 uint8 = 255
                                            _ = v88
                                        case "":
                                            fmt.Println(fi87)
                                        }
                                        _ = fi87
                                    }
                                }
                            }
                        }
                    }
                }
            }
            _ = v74
        }
        if true {
            if true {
                var v89 float32 = float32(((3.14159 + 0.0) * 3.14159))
                _ = v89
            } else {
                fmt.Println(v73, v73)
            }
        }
        _ = v73
    }
    if true {
        switch "a" {
        case "foo bar":
            var v90 string = "foo bar"
            _ = v90
        case "a":
            var v91 float32 = (-2.25 - 3.14159)
            _ = v91
        default:
            v92 := GenS3{F0: 127, F1: 811, F2: 4294967295}
            _ = v92
        }
    }
    v4 -= v4
    for fi93 := 0; fi93 < 1; fi93++ {
        v4 += (v4 | v4)
        for fi94 := 0; fi94 < 3; fi94++ {
            var v95 uint = 34
            var v96 uint32 = 65535
            _ = fi94
            _ = v95
            _ = v96
        }
        if ("foo bar" != "") {
            v4 = 0
            for fi97 := 0; fi97 < 11; fi97++ {
                switch (fi97 | (fi93 / (((fi93 * (fi97 - fi97))) | 1))) {
                case 128:
                    v4 = v4
                case 911:
                    var v98 float32 = (3.14159 / -2.25)
                    _ = v98
                default:
                    v4 = v4
                }
                var v99 int8 = 127
                v99 -= 0
                _ = fi97
                _ = v99
            }
            v4 += (^(v4))
        }
        _ = fi93
    }
    switch 32767 {
    case -696:
        for fi100 := 0; fi100 < 15; fi100++ {
            switch "" {
            case "\n\t":
                v4 -= (v4 & v4)
            default:
                switch "" {
                case "hello":
                    v4 = (v4 << ((589) & 63))
                }
            }
            v101 := GenS12{F0: ("hello" + "\n\t"), F1: 32767, F2: 100}
            _ = fi100
            _ = v101
        }
    default:
        v2["\n\t"] = 32768
    }
    v102 := GenS2{F0: "hello", F1: "a", F2: 1}
    var v103 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v103
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
