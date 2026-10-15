package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint8; F2 int }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 uint; F1 bool; F2 uint8 }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 uint32; F1 int32 }
type GenS14 struct { F0 int64; F1 int; F2 uint8 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int8; F1 int16; F2 uint }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 int64 = -801
    var v4 uint = 42
    var v5 int8 = int8(v3)
    v6 := GenS1{F0: -2147483648, F1: "", F2: -290}
    var v7 float64 = (-2.25 - (0.1 * 3.14159))
    var v8 int = 1
    if true {
        switch "\n\t" {
        case "a":
            v9 := GenS2{F0: -2147483648, F1: (^(v4))}
            _ = v9
        case "\n\t":
            switch v8 {
            case 822:
                if true {
                    v10 := GenS1{F0: 42, F1: ("hello" + "a"), F2: 2147483647}
                    _ = v10
                } else {
                    var v11 float64 = v7
                    fmt.Println(v8, v1)
                    if ((false && false) || true) {
                        switch (v8 >> ((270) & 63)) {
                        case -1:
                            fmt.Println(v4, v11)
                        default:
                            var v12 uint = (v4 / ((4294967295) | 1))
                            _ = v12
                        }
                        var v13 uint16 = 48
                        v14 := []int{}
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                }
            case -1:
                var v15 uint8 = 75
                _ = v15
            default:
                v16 := GenS3{F0: (v3 >> ((v4) & 63)), F1: 222}
                _ = v16
            }
        default:
            if (982 >= 321) {
                fmt.Println(v5, v5)
                v7 = v7
                fmt.Println(v1, v5, v5)
            }
        }
    } else {
        switch -9223372036854775808 {
        case 32767:
            v1 = (-(v5))
        case 4294967295:
            for fi17 := 0; fi17 < 3; fi17++ {
                v5 += (v1 & v5)
                _ = fi17
            }
        }
    }
    switch (v8 & (^(v8))) {
    case 32767:
        var v18 uint32 = 0
        _ = v18
    case 288:
        var v19 uint16 = 409
        _ = v19
    case -1:
        if (65535 <= 32767) {
            for fi20 := 0; fi20 < 9; fi20++ {
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 int32 = 42
                    var v23 uint = 65535
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                fmt.Println(v1, v1, v5)
                if (true && true) {
                    v2[(("\n\t" + "\n\t") + "\n\t")] = "\n\t"
                    switch "a" {
                    case "foo bar":
                        if (!true) {
                            fmt.Println(v5, v7, fi20)
                            var v24 float32 = ((((-2.25 / 1.5) - 1.5) * (1.5 / 1.5)) + 1.5)
                            fmt.Println(v1, v8)
                            _ = v24
                        }
                    case "\n\t":
                        v2[("" + ("" + (("a" + "foo bar") + ("foo bar" + "a"))))] = "foo bar"
                    case "":
                        var v25 string = ((((((("hello" + "foo bar") + "foo bar") + "") + "foo bar") + "hello") + "a") + (("\n\t" + "") + "foo bar"))
                        _ = v25
                    }
                    var v26 uint64 = 523
                    _ = v26
                } else {
                    if false {
                        var v27 int32 = -911
                        var v28 uint = 255
                        _ = v27
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v7 = v7
                            v4 += 255
                            switch "a" {
                            case "\n\t":
                                _ = len(v2)
                            case "hello":
                                switch ("" + "a") {
                                case "foo bar":
                                    v2[("hello" + "hello")] = "foo bar"
                                case "":
                                    var v30 int64 = v3
                                    _ = v30
                                }
                                if false {
                                    _ = len(v2)
                                    v31 := GenS2{F0: -155, F1: 2}
                                    switch "foo bar" {
                                    case "\n\t":
                                        if true {
                                            v7 = (v7 + -2.25)
                                            switch ("" + "hello") {
                                            case "foo bar":
                                                v5 -= v5
                                            case "hello":
                                                var v32 int32 = 0
                                                _ = v32
                                            case "\n\t":
                                                v33 := []int{}
                                                _ = v33
                                            }
                                        } else {
                                            var v34 float32 = ((3.14159 + 0.0) + (1e10 / 1e10))
                                            switch (-(v8)) {
                                            case 127:
                                                v35 := GenS3{F0: v3, F1: 128}
                                                _ = v35
                                            case 128:
                                                switch "\n\t" {
                                                case "hello":
                                                    if (0 != 1) {
                                                        switch fi20 {
                                                        case 127:
                                                            v7 = (v7 - -2.25)
                                                        default:
                                                            v34 = (float32((v7 + -2.25)) - (v34 - (v34 + v34)))
                                                        }
                                                        fmt.Println(v3, v1)
                                                        v2["a"] = (("" + "\n\t") + "")
                                                    } else {
                                                        if true {
                                                            switch v8 {
                                                            case 127:
                                                                v36 := []int{(fi20 % ((fi29) | 1)), -401}
                                                                _ = v36
                                                            }
                                                            v37 := []int{v8, v8, (fi20 + fi29), (v8 % ((fi20) | 1))}
                                                            _ = v37
                                                        }
                                                        v38 := GenS3{F0: (v3 * v3), F1: 0}
                                                        _ = v38
                                                    }
                                                case "foo bar":
                                                    switch v8 {
                                                    case -1:
                                                        switch v8 {
                                                        case -1:
                                                            v4 -= (^(v4))
                                                        case 32768:
                                                            switch ("" + (("foo bar" + "\n\t") + "foo bar")) {
                                                            case "a":
                                                                fmt.Println(v1)
                                                            case "":
                                                                if true {
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        switch -789 {
                                                                        case -2147483648:
                                                                            var v39 bool = false
                                                                            _ = v39
                                                                        case 719:
                                                                            var v40 uint64 = 42
                                                                            _ = v40
                                                                        case -1:
                                                                            if false {
                                                                                if false {
                                                                                    var v41 uint8 = 2
                                                                                    _ = v41
                                                                                } else {
                                                                                    var v42 float64 = (0.0 - ((v7 + (-2.25 - v7)) - ((((3.14159 / v7) - float64((3.14159 - v34))) - v7) / 1.5)))
                                                                                    var v43 uint64 = 2
                                                                                    _ = v42
                                                                                    _ = v43
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    if true {
                                                                        if false {
                                                                            v44 := []int{(fi29 + 0)}
                                                                            v3 = v3
                                                                            var v45 float32 = v34
                                                                            _ = v44
                                                                            _ = v45
                                                                        } else {
                                                                            var v46 uint8 = 2
                                                                            var v47 uint64 = 42
                                                                            _ = v46
                                                                            _ = v47
                                                                        }
                                                                        v1 = v5
                                                                        v2["\n\t"] = ("hello" + "a")
                                                                        if false {
                                                                            v48 := GenS2{F0: 42, F1: v4}
                                                                            if (42 < 4294967295) {
                                                                                var v49 float32 = v34
                                                                                fmt.Println(v4, v3, v1)
                                                                                switch ("foo bar" + "\n\t") {
                                                                                case "a":
                                                                                    if (true && (127 > 196)) {
                                                                                        switch 4294967295 {
                                                                                        case 574:
                                                                                            v50 := GenS2{F0: int32(v4), F1: 464}
                                                                                            _ = v50
                                                                                        case 536:
                                                                                            v7 = (v7 / float64(v49))
                                                                                        }
                                                                                        if (v1 >= int8(fi20)) {
                                                                                            v3 = -809
                                                                                            v51 := v2[("" + (("\n\t" + "\n\t") + ((("" + ("\n\t" + "foo bar")) + "hello") + ("hello" + "foo bar"))))]
                                                                                            _ = v51
                                                                                        }
                                                                                    } else {
                                                                                        v5 += v1
                                                                                        v7 = ((1.5 * (-2.25 - (1e10 - 1.5))) + (v7 / v7))
                                                                                        v52 := []int{v8}
                                                                                        _ = v52
                                                                                    }
                                                                                case "":
                                                                                    if true {
                                                                                        v53 := []int{(-(fi29))}
                                                                                        var v54 bool = (512 >= 239)
                                                                                        v5 += (-(v1))
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                    }
                                                                                case "hello":
                                                                                    v7 = ((float64(v34) + 0.1) / v7)
                                                                                default:
                                                                                    fmt.Println(v5, v8, v3)
                                                                                }
                                                                                v55 := GenS2{F0: -647, F1: (v4 - v4)}
                                                                                v2[""] = ("\n\t" + "hello")
                                                                                _ = v49
                                                                                _ = v55
                                                                            }
                                                                            fmt.Println(v5)
                                                                            _ = v48
                                                                        }
                                                                        fmt.Println(v1, v5, fi29)
                                                                    } else {
                                                                        var v56 int8 = (v1 + (v5 ^ 100))
                                                                        if false {
                                                                            switch (((("a" + "hello") + "hello") + "foo bar") + ("hello" + "hello")) {
                                                                            case "foo bar":
                                                                                if (uint32(v4) >= 2147483647) {
                                                                                    v57 := []int{736, fi29, fi29}
                                                                                    _ = v57
                                                                                }
                                                                            case "\n\t":
                                                                                if false {
                                                                                    switch ((("" + "\n\t") + "\n\t") + "a") {
                                                                                    case "foo bar":
                                                                                        v7 = float64(v34)
                                                                                    default:
                                                                                        var v58 uint32 = 891
                                                                                        _ = v58
                                                                                    }
                                                                                    v59 := []int{-9223372036854775808, fi20, -473}
                                                                                    v34 -= float32(3.14159)
                                                                                    _ = v59
                                                                                } else {
                                                                                    var v60 int32 = 42
                                                                                    v61 := GenS4{F0: (v60 - 693)}
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v1)
                                                                            default:
                                                                                if false {
                                                                                    var v62 uint = v4
                                                                                    v62 -= v62
                                                                                    v2["\n\t"] = "\n\t"
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                        } else {
                                                                            if (42 <= 2147483648) {
                                                                                v7 += (0.0 + v7)
                                                                                var v63 int64 = (v3 ^ 1)
                                                                                v5 += (^(v5))
                                                                                v64 := []int{(fi29 | 36), (-(fi29))}
                                                                                _ = v63
                                                                                _ = v64
                                                                            } else {
                                                                                v56 = 42
                                                                                _ = len(v2)
                                                                            }
                                                                            if true {
                                                                                if true {
                                                                                    switch fi20 {
                                                                                    case 0:
                                                                                        var v65 int32 = 216
                                                                                        _ = v65
                                                                                    default:
                                                                                        if (9223372036854775807 != 42) {
                                                                                            var v66 float32 = v34
                                                                                            _ = v66
                                                                                        }
                                                                                    }
                                                                                    v67 := GenS5{F0: -480}
                                                                                    var v68 bool = false
                                                                                    _ = v67
                                                                                    _ = v68
                                                                                } else {
                                                                                    v69 := []int{1, v8, 127}
                                                                                    _ = v69
                                                                                }
                                                                            } else {
                                                                                v34 = (1.5 - (v34 * (v34 / v34)))
                                                                                v4 = v4
                                                                            }
                                                                            if true {
                                                                                var v70 int16 = -774
                                                                                var v71 float64 = v7
                                                                                _ = v70
                                                                                _ = v71
                                                                            }
                                                                            var v72 float64 = 0.0
                                                                            v7 -= 1.5
                                                                            _ = v72
                                                                        }
                                                                        switch ("" + "a") {
                                                                        case "\n\t":
                                                                            var v73 uint16 = 1
                                                                            _ = v73
                                                                        case "a":
                                                                            v1 = v5
                                                                        }
                                                                        _ = v56
                                                                    }
                                                                    var v74 float64 = 1.5
                                                                    _ = v74
                                                                }
                                                            case "foo bar":
                                                                v34 += 1.5
                                                            default:
                                                                if (("hello" + "hello") > "hello") {
                                                                    _ = len(v2)
                                                                } else {
                                                                    if ((!(false && ((!true) && false))) || false) {
                                                                        if true {
                                                                            if false {
                                                                                if false {
                                                                                    if false {
                                                                                        switch "\n\t" {
                                                                                        case "hello":
                                                                                            fmt.Println(fi20, v34)
                                                                                        case "a":
                                                                                            v8 += fi20
                                                                                        case "":
                                                                                            if true {
                                                                                                v75 := []int{(v8 - v8)}
                                                                                                _ = v75
                                                                                            }
                                                                                        }
                                                                                        var v76 int32 = int32(fi29)
                                                                                        _ = v76
                                                                                    } else {
                                                                                        var v77 uint = v4
                                                                                        var v78 float64 = (0.1 / (((1.5 * v7) + 0.1) + v7))
                                                                                        _ = v77
                                                                                        _ = v78
                                                                                    }
                                                                                    fmt.Println(v5)
                                                                                }
                                                                                if (1.5 >= v7) {
                                                                                    v79 := GenS6{F0: v5, F1: 2, F2: v8}
                                                                                    if true {
                                                                                        switch 2147483647 {
                                                                                        case 708:
                                                                                            v5 = -24
                                                                                        case -3:
                                                                                            switch (v8 + fi29) {
                                                                                            case 32768:
                                                                                                switch (fi29 + v8) {
                                                                                                case 1:
                                                                                                    if ((0 == 255) || (v7 >= v7)) {
                                                                                                        switch "hello" {
                                                                                                        case "":
                                                                                                            var v80 string = ("a" + ((("\n\t" + ("\n\t" + "")) + (("" + "a") + ("a" + "foo bar"))) + "a"))
                                                                                                            _ = v80
                                                                                                        case "\n\t":
                                                                                                            var v81 float64 = v7
                                                                                                            fmt.Println(v81)
                                                                                                            var v82 int64 = (v3 % ((v3) | 1))
                                                                                                            _ = v81
                                                                                                            _ = v82
                                                                                                        case "hello":
                                                                                                            v4 -= v4
                                                                                                        }
                                                                                                        v4 = (v4 / ((v4) | 1))
                                                                                                    } else {
                                                                                                        v1 = 0
                                                                                                    }
                                                                                                }
                                                                                                var v83 uint64 = 18446744073709551615
                                                                                                _ = v83
                                                                                            }
                                                                                        default:
                                                                                            switch fi29 {
                                                                                            case -9223372036854775808:
                                                                                                v1 = v5
                                                                                            case 32767:
                                                                                                if (!false) {
                                                                                                    _ = len(v2)
                                                                                                    switch fi29 {
                                                                                                    case 42:
                                                                                                        var v84 float32 = ((1.5 + (-2.25 + 0.1)) / v34)
                                                                                                        var v85 uint64 = 4294967295
                                                                                                        v86 := GenS5{F0: v3}
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    case 490:
                                                                                                        if (v3 == 42) {
                                                                                                            switch fi20 {
                                                                                                            case -522:
                                                                                                                v7 += v7
                                                                                                            case -310:
                                                                                                                v87 := []int{v8, (^(fi20)), fi20}
                                                                                                                _ = v87
                                                                                                            case -9223372036854775808:
                                                                                                                fmt.Println(v34, v7)
                                                                                                            default:
                                                                                                                switch ("hello" + ("\n\t" + "\n\t")) {
                                                                                                                case "":
                                                                                                                    v4 = 229
                                                                                                                default:
                                                                                                                    var v88 uint8 = 37
                                                                                                                    _ = v88
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v89 := []int{(v8 % ((fi29) | 1)), (^(v8)), fi20, 32768}
                                                                                                    _ = v89
                                                                                                }
                                                                                            case 647:
                                                                                                var v90 int32 = 1
                                                                                                _ = v90
                                                                                            }
                                                                                            switch ("\n\t" + ("hello" + ("foo bar" + (("a" + "") + "\n\t")))) {
                                                                                            case "foo bar":
                                                                                                if true {
                                                                                                    v91 := GenS2{F0: 2147483647, F1: v4}
                                                                                                    _ = v91
                                                                                                }
                                                                                            case "\n\t":
                                                                                                switch v8 {
                                                                                                case 42:
                                                                                                    var v92 int8 = (v1 % ((31) | 1))
                                                                                                    _ = v92
                                                                                                default:
                                                                                                    var v93 uint = (v4 + (v4 % ((v4) | 1)))
                                                                                                    _ = v93
                                                                                                }
                                                                                            case "":
                                                                                                if true {
                                                                                                    _ = len(v2)
                                                                                                    var v94 int8 = -11
                                                                                                    if false {
                                                                                                        v95 := GenS1{F0: (v8 / (((v8 * fi20)) | 1)), F1: "hello", F2: 2147483647}
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        v96 := []int{}
                                                                                                        switch "" {
                                                                                                        case "":
                                                                                                            v97 := GenS3{F0: (^(v3)), F1: 48}
                                                                                                            _ = v97
                                                                                                        case "foo bar":
                                                                                                            if true {
                                                                                                                var v98 int8 = (v94 - (v5 | v94))
                                                                                                                var v99 int32 = -2147483648
                                                                                                                _ = v98
                                                                                                                _ = v99
                                                                                                            }
                                                                                                        case "\n\t":
                                                                                                            v3 = 869
                                                                                                        default:
                                                                                                            if true {
                                                                                                                _ = len(v2)
                                                                                                                v2["foo bar"] = "hello"
                                                                                                                fmt.Println(v34, v94, v7)
                                                                                                            } else {
                                                                                                                var v100 float64 = (-2.25 * 1.5)
                                                                                                                v1 -= (v1 / ((v94) | 1))
                                                                                                                fmt.Println(v4, v34)
                                                                                                                _ = v100
                                                                                                            }
                                                                                                        }
                                                                                                        var v101 int = (v8 | (fi20 * (v8 << ((v4) & 63))))
                                                                                                        switch 904 {
                                                                                                        case -1:
                                                                                                            if len(v96) >= 2 {
                                                                                                                v96 = v96[1:len(v96)]
                                                                                                            }
                                                                                                        case 520:
                                                                                                            var v102 uint32 = 830
                                                                                                            _ = v102
                                                                                                        default:
                                                                                                            fmt.Println(v7, v4, v34)
                                                                                                        }
                                                                                                        v103 := GenS4{F0: int32(v3)}
                                                                                                        _ = v101
                                                                                                        _ = v103
                                                                                                        _ = v96
                                                                                                    }
                                                                                                    var v104 int = 128
                                                                                                    _ = v104
                                                                                                    _ = v94
                                                                                                } else {
                                                                                                    switch "hello" {
                                                                                                    case "\n\t":
                                                                                                        fmt.Println(v1, v3, v7)
                                                                                                    case "a":
                                                                                                        v105 := []int{fi20, fi29, (fi20 - v8)}
                                                                                                        _ = v105
                                                                                                    case "":
                                                                                                        v106 := GenS6{F0: v1, F1: 128, F2: -2147483648}
                                                                                                        _ = v106
                                                                                                    }
                                                                                                }
                                                                                            default:
                                                                                                switch "\n\t" {
                                                                                                case "":
                                                                                                    v107 := GenS7{F0: -1, F1: "a"}
                                                                                                    var v108 uint8 = 127
                                                                                                    v109 := []int{(fi29 << ((710) & 63)), (fi20 % ((32768) | 1))}
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                default:
                                                                                                    fmt.Println(v4, v1, v34)
                                                                                                }
                                                                                            }
                                                                                            v1 = (v5 / (((-(v5))) | 1))
                                                                                        }
                                                                                        v110 := GenS8{F0: 42}
                                                                                        v7 = float64(-2.25)
                                                                                        v4 += 0
                                                                                        switch fi29 {
                                                                                        case 952:
                                                                                            v4 += (v4 * (v4 % (((v4 << (((v4 ^ (v4 | (v4 >> ((2) & 63))))) & 63))) | 1)))
                                                                                        default:
                                                                                            switch -197 {
                                                                                            case -2147483648:
                                                                                                switch "" {
                                                                                                case "hello":
                                                                                                    v111 := GenS1{F0: (fi20 << ((18446744073709551615) & 63)), F1: "a", F2: 2147483647}
                                                                                                    _ = v111
                                                                                                case "":
                                                                                                    _ = len(v2)
                                                                                                case "foo bar":
                                                                                                    if true {
                                                                                                        fmt.Println(v7, v3)
                                                                                                    }
                                                                                                }
                                                                                            case -709:
                                                                                                var v112 int32 = -1
                                                                                                _ = v112
                                                                                            case -80:
                                                                                                fmt.Println(fi29, v4, fi29)
                                                                                            default:
                                                                                                fmt.Println(v4, v8, v5)
                                                                                            }
                                                                                        }
                                                                                        _ = v110
                                                                                    }
                                                                                    var v113 int32 = -1
                                                                                    _ = v113
                                                                                    _ = v79
                                                                                } else {
                                                                                    delete(v2, "foo bar")
                                                                                }
                                                                                v3 += v3
                                                                                v114 := []int{-197}
                                                                                _ = v114
                                                                            }
                                                                            var v115 uint = (v4 ^ 0)
                                                                            _ = v115
                                                                        }
                                                                        v116 := []int{}
                                                                        _ = v116
                                                                    } else {
                                                                        fmt.Println(v8, fi29, v7)
                                                                    }
                                                                    v7 += (v7 + v7)
                                                                    v4 = (^(v4))
                                                                    var v117 int16 = 638
                                                                    var v118 int64 = v3
                                                                    _ = v117
                                                                    _ = v118
                                                                }
                                                            }
                                                        default:
                                                            v1 += (^(v5))
                                                        }
                                                    default:
                                                        switch ((("\n\t" + ("" + ("foo bar" + ("\n\t" + ("a" + (("\n\t" + "\n\t") + "")))))) + "") + "\n\t") {
                                                        case "hello":
                                                            var v119 int = (v8 + (^(v8)))
                                                            _ = v119
                                                        default:
                                                            var v120 bool = (-81 < v5)
                                                            _ = v120
                                                        }
                                                    }
                                                case "a":
                                                    var v121 uint64 = 2
                                                    _ = v121
                                                default:
                                                    v122 := []int{(fi29 & (fi20 & (-(fi20)))), fi20, fi20, v8}
                                                    _ = v122
                                                }
                                            case 2147483647:
                                                if false {
                                                    fmt.Println(v34, v7, v5)
                                                }
                                                v123 := GenS8{F0: 1}
                                                _ = v123
                                            }
                                            v4 += 1
                                            v8 = 128
                                            _ = v34
                                        }
                                    case "hello":
                                        var v124 int = v8
                                        _ = v124
                                    }
                                    v3 = (-(v3))
                                    if (!true) {
                                        v125 := []int{32767, -213}
                                        if (!false) {
                                            if (0 < 1) {
                                                v125 = append(v125, (v8 & fi29))
                                                var v126 int8 = v1
                                                _ = v126
                                            } else {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v127 := GenS1{F0: v8, F1: "a", F2: 0}
                                                    _ = v127
                                                case "hello":
                                                    fmt.Println(v5, v1)
                                                case "a":
                                                    var v128 string = ("\n\t" + ("a" + "a"))
                                                    _ = v128
                                                }
                                            }
                                            fmt.Println(fi20)
                                            if false {
                                                switch (v8 + (^(fi29))) {
                                                case 105:
                                                    var v129 int64 = (^(v3))
                                                    _ = v129
                                                case 2147483647:
                                                    switch "a" {
                                                    case "a":
                                                        var v130 int = (v8 / ((v8) | 1))
                                                        _ = v130
                                                    case "hello":
                                                        var v131 int8 = 0
                                                        _ = v131
                                                    default:
                                                        var v132 bool = true
                                                        _ = v132
                                                    }
                                                    switch (fi29 ^ 528) {
                                                    case 877:
                                                        if false {
                                                            var v133 bool = false
                                                            var v134 uint32 = 1
                                                            _ = v133
                                                            _ = v134
                                                        } else {
                                                            fmt.Println(v1, fi20)
                                                            if len(v125) >= 2 {
                                                                v125 = v125[1:len(v125)]
                                                            }
                                                        }
                                                    case 4294967295:
                                                        v4 = v4
                                                    case 42:
                                                        var v135 int16 = 32767
                                                        _ = v135
                                                    }
                                                }
                                                switch "" {
                                                case "a":
                                                    var v136 uint64 = 13
                                                    _ = v136
                                                case "hello":
                                                    var v137 uint8 = 128
                                                    _ = v137
                                                }
                                                switch -1 {
                                                case 4294967295:
                                                    switch (("foo bar" + "a") + "\n\t") {
                                                    case "":
                                                        if len(v125) >= 2 {
                                                            v125 = v125[1:len(v125)]
                                                        }
                                                    case "a":
                                                        var v138 int32 = -1
                                                        _ = v138
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                }
                                            }
                                            var v139 int16 = -683
                                            _ = v139
                                        }
                                        v5 -= v1
                                        _ = v125
                                    }
                                    _ = v31
                                } else {
                                    v4 += (v4 & (v4 >> ((v4) & 63)))
                                }
                                var v140 int16 = 42
                                _ = v140
                            case "foo bar":
                                var v141 int = 2147483647
                                _ = v141
                            }
                            _ = fi29
                        }
                        if true {
                            v142 := GenS9{F0: (v4 / (((v4 & v4)) | 1))}
                            _ = v142
                        }
                    }
                    for fi143 := 0; fi143 < 16; fi143++ {
                        var v144 uint32 = 2147483648
                        _ = fi143
                        _ = v144
                    }
                    switch fi20 {
                    case 1:
                        switch (fi20 % ((32768) | 1)) {
                        case 952:
                            v145 := GenS10{F0: uint8(v4), F1: 435}
                            _ = v145
                        case 127:
                            v146 := v2["hello"]
                            _ = v146
                        case 848:
                            var v147 bool = false
                            _ = v147
                        }
                    }
                }
                _ = fi20
            }
            for fi148 := 0; fi148 < 15; fi148++ {
                fmt.Println(v4, v4)
                var v149 int64 = -9223372036854775808
                _ = fi148
                _ = v149
            }
        } else {
            var v150 int = v8
            v151 := GenS6{F0: (v1 >> (((^(v4))) & 63)), F1: 42, F2: (-(v150))}
            _ = v150
            _ = v151
        }
    default:
        switch ("\n\t" + "") {
        case "foo bar":
            for fi152 := 0; fi152 < 11; fi152++ {
                if true {
                    v5 = (^(v5))
                } else {
                    if false {
                        v153 := GenS11{F0: v4, F1: (!(42 <= -1)), F2: 128}
                        _ = v153
                    } else {
                        v154 := GenS12{F0: v8, F1: -1}
                        v155 := GenS2{F0: -1, F1: 405}
                        _ = v154
                        _ = v155
                    }
                    var v156 int = fi152
                    _ = v156
                }
                _ = fi152
            }
        case "hello":
            for fi157 := 0; fi157 < 20; fi157++ {
                v158 := GenS13{F0: 2147483647, F1: -2147483648}
                var v159 int8 = -26
                for fi160 := 0; fi160 < 15; fi160++ {
                    v8 += (fi160 >> (((v4 & v4)) & 63))
                    _ = fi160
                }
                _ = fi157
                _ = v158
                _ = v159
            }
        default:
            var v161 int16 = 1000
            _ = v161
        }
    }
    v162 := GenS13{F0: 856, F1: -1}
    var v163 uint8 = 128
    var v164 int = v8
    v163 = (v163 << (((v4 & (v4 * (^(v4))))) & 63))
    var v165 float32 = (3.14159 / float32(v7))
    switch "\n\t" {
    case "a":
        fmt.Println(v163, v1, v1)
    default:
        for fi166 := 0; fi166 < 7; fi166++ {
            switch ("\n\t" + (("\n\t" + "a") + "hello")) {
            case "":
                switch -538 {
                case 32768:
                    if false {
                        var v167 int64 = v3
                        v7 = float64(float32((0.0 - -2.25)))
                        if false {
                            v2["foo bar"] = "\n\t"
                            v168 := GenS14{F0: v3, F1: 42, F2: (v163 << ((2) & 63))}
                            v169 := GenS15{F0: v163}
                            _ = v168
                            _ = v169
                        } else {
                            fmt.Println(v8, v3)
                            fmt.Println(v165, v4)
                            var v170 uint = (v4 >> ((v4) & 63))
                            _ = v170
                        }
                        _ = v167
                    }
                case 32767:
                    if (0 == 559) {
                        v171 := []int{(v164 & fi166)}
                        var v172 uint64 = 0
                        v173 := GenS14{F0: (^(v3)), F1: 2147483647, F2: 2}
                        _ = v171
                        _ = v172
                        _ = v173
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v174 string = "hello"
                            _ = v174
                        default:
                            v3 = v3
                        }
                        var v175 string = (("hello" + ("hello" + ("foo bar" + "hello"))) + "")
                        _ = v175
                    }
                }
            case "foo bar":
                v176 := GenS7{F0: -1, F1: (("hello" + "") + "\n\t")}
                _ = v176
            case "hello":
                for fi177 := 0; fi177 < 9; fi177++ {
                    var v178 uint64 = 0
                    if false {
                        var v179 bool = (!false)
                        var v180 float32 = 1.5
                        v4 = (^(v4))
                        if (!v179) {
                            v181 := GenS3{F0: v3, F1: (v163 >> ((v4) & 63))}
                            v2[""] = ("foo bar" + "hello")
                            v163 -= (v163 * (v163 * (v163 - v163)))
                            _ = v181
                        } else {
                            fmt.Println(fi177, v164)
                            switch "foo bar" {
                            case "hello":
                                if v179 {
                                    v178 = 42
                                    fmt.Println(v180, v165)
                                }
                            case "":
                                var v182 uint8 = (^(v163))
                                _ = v182
                            default:
                                v183 := GenS16{F0: v1, F1: int16(v163), F2: (v4 >> ((1) & 63))}
                                _ = v183
                            }
                            v5 -= 0
                        }
                        _ = v179
                        _ = v180
                    }
                    fmt.Println(v165, v8, v1)
                    v184 := v2["hello"]
                    if (984 != 32767) {
                        var v185 int32 = 1
                        v186 := GenS2{F0: (v185 ^ v185), F1: uint(v163)}
                        _ = v185
                        _ = v186
                    }
                    _ = fi177
                    _ = v178
                    _ = v184
                }
            default:
                for fi187 := 0; fi187 < 9; fi187++ {
                    if false {
                        var v188 uint = v4
                        _ = v188
                    } else {
                        if true {
                            v7 = 1e10
                            switch "a" {
                            case "\n\t":
                                var v189 uint64 = 0
                                v5 += v5
                                _ = v189
                            case "":
                                v190 := GenS7{F0: -85, F1: ((("foo bar" + "") + ("" + "")) + "\n\t")}
                                v191 := GenS3{F0: v3, F1: (v163 << (((v4 << (((^(v4))) & 63))) & 63))}
                                switch 0 {
                                case -267:
                                    v192 := GenS15{F0: v163}
                                    _ = v192
                                }
                                _ = v190
                                _ = v191
                            }
                            var v193 uint32 = 2147483647
                            var v194 bool = true
                            _ = v193
                            _ = v194
                        }
                    }
                    var v195 int8 = (v5 | (v1 * v1))
                    if (true || false) {
                        fmt.Println(v4)
                    } else {
                        v196 := GenS16{F0: int8(v163), F1: 1000, F2: (v4 | 65535)}
                        var v197 uint = (v4 ^ v4)
                        v198 := v2["foo bar"]
                        _ = v196
                        _ = v197
                        _ = v198
                    }
                    _ = fi187
                    _ = v195
                }
            }
            fmt.Println(fi166, v8)
            v164 = -467
            _ = fi166
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v163:", v163)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v162
    _ = v163
    _ = v164
    _ = v165
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
