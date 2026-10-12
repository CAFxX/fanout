package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int; F1 int8; F2 int64 }
type GenS6 struct { F0 uint; F1 int32; F2 uint }
type GenS7 struct { F0 uint; F1 int; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS9 struct { F0 uint; F1 int8; F2 int }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 string; F1 int32 }
type GenS13 struct { F0 int; F1 uint16; F2 string }

func main() {
    v1 := []uint32{}
    var v2 int8 = -127
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    if (((((1e10 + (-2.25 * (1e10 + 3.14159))) + float32(-2.25)) * float32(((1.5 - 1.5) - float64(1.5)))) * 0.1) >= 1.5) {
        for fi5 := 0; fi5 < 15; fi5++ {
            v6 := GenS1{F0: (1 == 2147483647)}
            var v7 string = ""
            _ = fi5
            _ = v6
            _ = v7
        }
        switch "foo bar" {
        case "a":
            var v8 int = -931
            _ = v8
        case "foo bar":
            v2 = (v2 & 16)
        case "hello":
            var v9 bool = (-159 == 42)
            _ = v9
        }
    }
    v10 := GenS1{F0: (((179 <= 255) && true) && (!((127 > 0) && ((v2 << ((1) & 63)) != v2))))}
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            switch "\n\t" {
            case "a":
                v2 = (-(v2))
            case "\n\t":
                v13 := GenS2{F0: 1}
                _ = v13
            case "foo bar":
                if (!false) {
                    switch "a" {
                    case "foo bar":
                        if false {
                            v14 := GenS3{F0: 0, F1: -1, F2: fi11}
                            fmt.Println(v2, fi12)
                            var v15 uint16 = 65535
                            fmt.Println(v2, v2)
                            _ = v14
                            _ = v15
                        }
                    default:
                        var v16 uint = 2
                        _ = v16
                    }
                    fmt.Println(v2)
                    fmt.Println(fi11, v2, fi12)
                    v2 = -128
                    switch 2147483647 {
                    case 42:
                        var v17 int64 = 0
                        _ = v17
                    case 643:
                        switch "hello" {
                        case "hello":
                            v18 := GenS4{F0: 2147483647, F1: 817}
                            _ = v18
                        case "":
                            v19 := GenS4{F0: 42, F1: 483}
                            _ = v19
                        case "foo bar":
                            v1 = append(v1, 491)
                        }
                    case -2147483648:
                        if false {
                            v20 := GenS3{F0: 243, F1: -284, F2: 500}
                            v2 = v2
                            var v21 int8 = v2
                            v21 = v21
                            var v22 int64 = 42
                            v2 = (v21 % (((v21 - v21)) | 1))
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            var v23 int64 = 1
                            switch fi12 {
                            case -9223372036854775808:
                                var v24 int8 = (v2 / (((v2 % (((v2 * (v2 << ((2) & 63)))) | 1))) | 1))
                                _ = v24
                            }
                            v23 += (v23 | 9223372036854775807)
                            _ = v23
                        }
                    }
                    var v25 float32 = float32((1.5 * float64((3.14159 / 3.14159))))
                    _ = v25
                } else {
                    var v26 int16 = 1000
                    _ = v26
                }
                delete(v3, "a")
            default:
                switch ("foo bar" + "") {
                case "a":
                    if true {
                        switch fi11 {
                        case 42:
                            var v27 bool = (42 < 0)
                            _ = v27
                        case 1:
                            var v28 int16 = -32768
                            _ = v28
                        default:
                            v29 := GenS4{F0: 449, F1: 596}
                            _ = v29
                        }
                        var v30 uint64 = 18446744073709551615
                        v30 = v30
                        _ = v30
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 645
                        }
                        if len(v1) > 0 {
                            v1[0] = 1
                        }
                    }
                case "\n\t":
                    v31 := GenS3{F0: 42, F1: 1, F2: (fi12 & -9223372036854775808)}
                    _ = v31
                case "":
                    v32 := GenS1{F0: true}
                    _ = v32
                }
            }
            v33 := GenS3{F0: 43, F1: -1, F2: fi12}
            fmt.Println(v2, fi12, v2)
            if true {
                _ = len(v1)
                if (!true) {
                    if true {
                        var v34 uint32 = 385
                        _ = v34
                    } else {
                        switch "hello" {
                        case "foo bar":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            fmt.Println(fi12, fi12)
                        }
                        var v35 int16 = 697
                        v36 := GenS5{F0: fi11, F1: (v2 << ((0) & 63)), F2: 0}
                        _ = v35
                        _ = v36
                    }
                    fmt.Println(fi12, fi12, v2)
                    v37 := GenS2{F0: 1}
                    _ = v37
                } else {
                    switch fi12 {
                    case 32767:
                        var v38 int32 = -66
                        switch "a" {
                        case "\n\t":
                            if true {
                                var v39 uint32 = 1
                                _ = v39
                            } else {
                                switch "\n\t" {
                                case "":
                                    switch "" {
                                    case "hello":
                                        v40 := GenS6{F0: 2, F1: v38, F2: 1}
                                        v2 += 3
                                        var v41 float64 = 1.5
                                        _ = v40
                                        _ = v41
                                    case "foo bar":
                                        v42 := GenS5{F0: (fi12 | fi12), F1: -74, F2: 1}
                                        _ = v42
                                    default:
                                        switch fi11 {
                                        case 128:
                                            switch ("\n\t" + "\n\t") {
                                            case "\n\t":
                                                switch "\n\t" {
                                                case "":
                                                    fmt.Println(fi12, fi11, fi11)
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            case "":
                                                if len(v1) > 0 {
                                                    v1[0] = 339
                                                }
                                            case "a":
                                                switch (fi12 >> ((874) & 63)) {
                                                case 1:
                                                    var v43 uint8 = 2
                                                    _ = v43
                                                default:
                                                    delete(v4, 2147483648)
                                                }
                                            }
                                        case 365:
                                            delete(v4, 2147483647)
                                        }
                                    }
                                case "\n\t":
                                    if (false && false) {
                                        var v44 uint32 = 1
                                        var v45 float32 = (-2.25 / 3.14159)
                                        if (0 < -1) {
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                fmt.Println(fi11, fi11)
                                            case "hello":
                                                v4[(v44 ^ (v44 / ((v44) | 1)))] = 1
                                            case "a":
                                                delete(v4, 2147483648)
                                            }
                                            if true {
                                                v46 := GenS7{F0: 2, F1: -2147483648, F2: 0}
                                                delete(v4, v44)
                                                if true {
                                                    v45 = 1e10
                                                } else {
                                                    if false {
                                                        var v47 uint8 = 127
                                                        _ = v47
                                                    }
                                                    var v48 uint16 = 32767
                                                    _ = v48
                                                }
                                                _ = v46
                                            } else {
                                                v49 := GenS4{F0: 2147483647, F1: 65535}
                                                v50 := GenS1{F0: (226 > 65535)}
                                                v51 := GenS3{F0: 54, F1: 42, F2: fi11}
                                                _ = v49
                                                _ = v50
                                                _ = v51
                                            }
                                            if true {
                                                if false {
                                                    if true {
                                                        v44 -= 462
                                                    } else {
                                                        if (true && true) {
                                                            v2 = (v2 % ((1) | 1))
                                                            v52 := GenS7{F0: 414, F1: 32768, F2: 0}
                                                            _ = v52
                                                        }
                                                        fmt.Println(v38, v45)
                                                    }
                                                } else {
                                                    v53 := GenS8{F0: 0, F1: 945, F2: 268}
                                                    _ = v53
                                                }
                                            }
                                        } else {
                                            v45 = (v45 * ((v45 - (v45 / v45)) - (1e10 - v45)))
                                            switch 927 {
                                            case 200:
                                                if false {
                                                    _ = len(v3)
                                                    switch ("foo bar" + ("foo bar" + "a")) {
                                                    case "hello":
                                                        v54 := GenS9{F0: 0, F1: v2, F2: fi11}
                                                        _ = v54
                                                    case "":
                                                        var v55 uint = 345
                                                        _ = v55
                                                    }
                                                    if true {
                                                        var v56 int = fi11
                                                        v56 += (fi12 & (fi12 % (((v56 ^ (v56 + 649))) | 1)))
                                                        _ = v56
                                                    }
                                                }
                                            case 4294967295:
                                                if false {
                                                    v45 = v45
                                                    if false {
                                                        v57 := GenS9{F0: 695, F1: -128, F2: fi12}
                                                        if false {
                                                            v44 = 932
                                                            if true {
                                                                var v58 bool = ((v2 ^ v2) >= 100)
                                                                if (fi11 != 0) {
                                                                    var v59 string = "hello"
                                                                    v1 = append(v1, v44)
                                                                    _ = v59
                                                                }
                                                                _ = v58
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            var v60 uint64 = 18446744073709551615
                                                            _ = v60
                                                        }
                                                        if false {
                                                            v61 := GenS4{F0: 625, F1: 255}
                                                            _ = v61
                                                        }
                                                        _ = v57
                                                    }
                                                }
                                            }
                                        }
                                        _ = v44
                                        _ = v45
                                    }
                                    if (true || true) {
                                        switch (-(fi12)) {
                                        case 127:
                                            delete(v4, 65535)
                                        case 42:
                                            if (!(1.5 != (1e10 - 3.14159))) {
                                                var v62 uint = 2
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v62
                                            } else {
                                                var v63 uint32 = 42
                                                v63 += (v63 * 2147483647)
                                                _ = v63
                                            }
                                        }
                                        var v64 float32 = (0.0 / (3.14159 + (-2.25 / 1e10)))
                                        _ = v64
                                    }
                                    if (65535 <= 255) {
                                        if false {
                                            switch 2147483647 {
                                            case 32768:
                                                switch fi11 {
                                                case 128:
                                                    var v65 string = (("\n\t" + ((((("\n\t" + "a") + ("foo bar" + "")) + "foo bar") + "a") + "hello")) + ("\n\t" + ("\n\t" + "hello")))
                                                    _ = v65
                                                case 684:
                                                    v66 := GenS10{F0: (v38 % ((886) | 1))}
                                                    _ = v66
                                                default:
                                                    var v67 uint64 = 1
                                                    _ = v67
                                                }
                                                fmt.Println(v2)
                                            }
                                        }
                                        v38 = v38
                                    }
                                    if (fi12 == (fi12 >> ((948) & 63))) {
                                        fmt.Println(v38, fi11, fi12)
                                        v68 := GenS11{F0: (("\n\t" + "a") + "hello")}
                                        _ = v68
                                    }
                                }
                                if (2 <= 18446744073709551615) {
                                    if true {
                                        fmt.Println(fi11)
                                    } else {
                                        _ = len(v1)
                                        var v69 float64 = (1e10 - (3.14159 / 0.1))
                                        _ = v69
                                    }
                                } else {
                                    if (2147483647 <= 2147483647) {
                                        var v70 uint = 42
                                        _ = v70
                                    } else {
                                        delete(v3, "a")
                                        var v71 uint8 = 1
                                        _ = v71
                                    }
                                    v72 := GenS12{F0: "a", F1: v38}
                                    var v73 int8 = v2
                                    _ = v72
                                    _ = v73
                                }
                            }
                            v2 = -128
                            var v74 uint64 = 2
                            _ = v74
                        case "":
                            var v75 int8 = (^(v2))
                            _ = v75
                        default:
                            v76 := v4[2]
                            _ = v76
                        }
                        _ = v38
                    case 806:
                        if false {
                            switch 128 {
                            case 1:
                                switch (("\n\t" + "a") + ("" + "")) {
                                case "a":
                                    switch fi11 {
                                    case 32768:
                                        switch "\n\t" {
                                        case "a":
                                            v77 := GenS13{F0: 1, F1: 67, F2: ("hello" + "foo bar")}
                                            _ = v77
                                        case "foo bar":
                                            if (18446744073709551615 > 42) {
                                                if (0 <= 0) {
                                                    if (2 <= 16) {
                                                        v2 = (v2 / (((v2 * (v2 / ((v2) | 1)))) | 1))
                                                        v2 -= 42
                                                    }
                                                }
                                                var v78 string = ("a" + "a")
                                                _ = v78
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v79 int32 = 361
                                                switch ((("\n\t" + "\n\t") + "hello") + ("" + ((("\n\t" + "foo bar") + "a") + "\n\t"))) {
                                                case "\n\t":
                                                    var v80 string = ""
                                                    _ = v80
                                                case "foo bar":
                                                    v81 := v4[0]
                                                    _ = v81
                                                case "":
                                                    v79 += v79
                                                default:
                                                    if ((9223372036854775807 <= 1) || (!true)) {
                                                        if (v79 == (v79 | v79)) {
                                                            if false {
                                                                v82 := GenS9{F0: 1, F1: -1, F2: fi12}
                                                                if (v2 != v2) {
                                                                    switch ("" + ("a" + ((("a" + "a") + "a") + ""))) {
                                                                    case "foo bar":
                                                                        var v83 float64 = 1e10
                                                                        _ = v83
                                                                    }
                                                                    _ = len(v4)
                                                                } else {
                                                                    var v84 uint64 = 42
                                                                    v85 := GenS7{F0: 693, F1: fi12, F2: 65535}
                                                                    if (true || (true && ((float32(1e10) / (3.14159 / 1.5)) >= 1e10))) {
                                                                        if false {
                                                                            if (true && false) {
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v2 = (^(v2))
                                                                            }
                                                                            fmt.Println(v2, v79)
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                                _ = v82
                                                            }
                                                        }
                                                        if (!(true || false)) {
                                                            var v86 bool = (1e10 == (3.14159 - (-2.25 * (1.5 * 1.5))))
                                                            var v87 int16 = 0
                                                            if (!false) {
                                                                switch (fi12 << ((2) & 63)) {
                                                                case 2147483647:
                                                                    if v86 {
                                                                        var v88 uint = 255
                                                                        v86 = ((4294967295 > 0) || false)
                                                                        _ = v88
                                                                    }
                                                                case 32768:
                                                                    var v89 uint = 18446744073709551615
                                                                    _ = v89
                                                                case 1:
                                                                    var v90 int64 = 1
                                                                    _ = v90
                                                                default:
                                                                    var v91 int64 = -9223372036854775808
                                                                    _ = v91
                                                                }
                                                            }
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        switch "" {
                                                        case "\n\t":
                                                            var v92 int16 = 394
                                                            _ = v92
                                                        case "foo bar":
                                                            var v93 uint64 = 148
                                                            _ = v93
                                                        default:
                                                            delete(v4, 0)
                                                        }
                                                    }
                                                }
                                                _ = v79
                                            }
                                        case "hello":
                                            if (0 < 42) {
                                                var v94 int32 = 0
                                                v95 := GenS5{F0: -9223372036854775808, F1: (v2 - (v2 % ((0) | 1))), F2: 679}
                                                var v96 uint8 = 1
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            } else {
                                                v2 = v2
                                                v97 := GenS9{F0: 65535, F1: (v2 ^ (v2 | (^(v2)))), F2: -2147483648}
                                                _ = v97
                                            }
                                        }
                                    case 4294967295:
                                        _ = len(v1)
                                    case 1:
                                        var v98 int16 = 1
                                        _ = v98
                                    default:
                                        fmt.Println(fi11, fi11, fi11)
                                    }
                                }
                            }
                            if true {
                                v2 = (v2 % ((-1) | 1))
                                switch ("hello" + "\n\t") {
                                case "":
                                    _ = len(v3)
                                }
                            }
                        }
                    case 42:
                        var v99 uint16 = 588
                        if false {
                            var v100 float64 = -2.25
                            switch 32767 {
                            case 1:
                                fmt.Println(v2, fi12)
                            case 32768:
                                v2 += 42
                            default:
                                v100 += v100
                            }
                            _ = v100
                        }
                        _ = v99
                    default:
                        v101 := GenS1{F0: false}
                        _ = v101
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 200
                }
            }
            _ = fi12
            _ = v33
        }
        _ = fi11
    }
    for fi102 := 0; fi102 < 16; fi102++ {
        var v103 int32 = 0
        v1 = append(v1, 65535)
        _ = fi102
        _ = v103
    }
    v1 = append(v1, 2147483648)
    var v104 uint16 = 1
    var v105 bool = (!false)
    fmt.Println(v2)
    for fi106 := 0; fi106 < 12; fi106++ {
        v107 := GenS13{F0: fi106, F1: 65535, F2: ("a" + "foo bar")}
        _ = fi106
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
}
