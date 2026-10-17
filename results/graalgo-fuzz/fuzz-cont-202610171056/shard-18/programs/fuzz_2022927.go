package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 bool; F1 string; F2 bool }
type GenS12 struct { F0 uint; F1 int64 }
type GenS13 struct { F0 uint; F1 string; F2 uint16 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint32 }

func main() {
    v1 := []int{127, 127, 32767, 0}
    var v2 int8 = 1
    v3 := []uint32{2147483648, 1}
    var v4 int32 = -2147483648
    fmt.Println(v2, v4)
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v2 -= (v2 | (-(v2)))
    v4 = v4
    v5 := GenS1{F0: v4, F1: (v4 | v4), F2: (true && false)}
    if true {
        for fi6 := 0; fi6 < 17; fi6++ {
            if false {
                for fi7 := 0; fi7 < 1; fi7++ {
                    switch ("" + "foo bar") {
                    case "a":
                        if (v4 != v4) {
                            fmt.Println(fi7, fi6)
                            var v8 float32 = 0.1
                            var v9 float32 = 1e10
                            _ = v8
                            _ = v9
                        } else {
                            var v10 uint = 1
                            switch fi6 {
                            case 4294967295:
                                if len(v1) > 0 {
                                    v1[0] = int(v2)
                                }
                            default:
                                v2 = -7
                            }
                            v4 = (^(v4))
                            _ = v10
                        }
                    case "\n\t":
                        switch -287 {
                        case -1:
                            v11 := GenS1{F0: v4, F1: 1, F2: false}
                            _ = v11
                        case 9223372036854775807:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            var v12 int64 = 207
                            _ = v12
                        }
                    case "hello":
                        v2 = (-(v2))
                    }
                    _ = fi7
                }
                fmt.Println(fi6, v4)
                var v13 uint64 = 2
                _ = v13
            } else {
                switch "\n\t" {
                case "\n\t":
                    v4 = (v4 & 0)
                case "foo bar":
                    v4 += (v4 >> ((103) & 63))
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 float32 = (-2.25 * (-2.25 - 1.5))
                    _ = fi14
                    _ = v15
                }
            }
            var v16 int64 = 9223372036854775807
            _ = len(v3)
            _ = fi6
            _ = v16
        }
        if false {
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 float64 = -2.25
                _ = fi17
                _ = v18
            }
        } else {
            switch ((((((("foo bar" + "hello") + ("foo bar" + ("hello" + "foo bar"))) + "a") + ("hello" + "")) + "\n\t") + "hello") + ("a" + "hello")) {
            case "":
                var v19 float32 = (0.1 + float32((-2.25 + -2.25)))
                _ = v19
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 0
                }
            default:
                if (1 > 1) {
                    v4 -= -1
                } else {
                    fmt.Println(v2, v2)
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if ((!(false || (false && (!(4294967295 != 2))))) && (!((true && (!((!false) && true))) && true))) {
                            fmt.Println(fi20)
                            _ = len(v1)
                        } else {
                            fmt.Println(v2, v2)
                            var v21 int8 = v2
                            _ = v21
                        }
                        _ = fi20
                    }
                }
            }
            var v22 int64 = 0
            _ = v22
        }
    }
    var v23 float64 = float64(((float32(float64(1.5)) + 1.5) + 3.14159))
    if (18446744073709551615 <= 346) {
        for fi24 := 0; fi24 < 16; fi24++ {
            if (("hello" != (("" + "foo bar") + ("foo bar" + ("a" + "")))) || (false || (128 > 128))) {
                switch 32767 {
                case -2147483648:
                    v25 := GenS1{F0: v4, F1: -1, F2: ((-(v2)) > 0)}
                    _ = v25
                case 32767:
                    if (0 <= (v4 | 2147483647)) {
                        v26 := GenS1{F0: 2147483647, F1: 42, F2: true}
                        v2 = v2
                        var v27 uint8 = 255
                        _ = v26
                        _ = v27
                    }
                default:
                    if true {
                        v28 := GenS1{F0: v4, F1: 2147483647, F2: (109 >= -718)}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v28
                    }
                }
                if true {
                    for fi29 := 0; fi29 < 12; fi29++ {
                        switch (("hello" + "") + "a") {
                        case "foo bar":
                            var v30 string = (("foo bar" + "\n\t") + "a")
                            _ = v30
                        case "a":
                            _ = len(v3)
                            fmt.Println(v23)
                        default:
                            v2 = (v2 - 42)
                        }
                        fmt.Println(fi24)
                        _ = fi29
                    }
                    for fi31 := 0; fi31 < 3; fi31++ {
                        switch fi31 {
                        case 41:
                            switch fi24 {
                            case 94:
                                v32 := GenS1{F0: 0, F1: v4, F2: (false && (!true))}
                                _ = v32
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                            }
                        case 127:
                            v33 := GenS2{F0: -1}
                            _ = v33
                        default:
                            var v34 uint16 = 776
                            _ = v34
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v35 float64 = float64(0.0)
                        _ = fi31
                        _ = v35
                    }
                    switch "" {
                    case "hello":
                        fmt.Println(v4, v23, fi24)
                    case "\n\t":
                        if true {
                            v2 = (^(v2))
                            v4 = v4
                            if ((137 > 0) || true) {
                                if (-2147483648 > (v4 / (((v4 * v4)) | 1))) {
                                    v36 := GenS1{F0: (v4 - v4), F1: v4, F2: false}
                                    v4 -= (v4 - v4)
                                    _ = v36
                                }
                                v37 := GenS3{F0: 42, F1: 255}
                                _ = v37
                            } else {
                                if false {
                                    if true {
                                        _ = len(v3)
                                        var v38 float64 = -2.25
                                        _ = v38
                                    } else {
                                        switch -2147483648 {
                                        case -816:
                                            switch "hello" {
                                            case "a":
                                                v39 := GenS4{F0: "a"}
                                                _ = v39
                                            case "foo bar":
                                                v1 = append(v1, (fi24 * fi24))
                                            case "":
                                                if false {
                                                    v40 := GenS4{F0: "a"}
                                                    _ = v40
                                                } else {
                                                    for fi41 := 0; fi41 < 19; fi41++ {
                                                        v4 += v4
                                                        var v42 int64 = -639
                                                        switch (fi41 >> ((96) & 63)) {
                                                        case 32767:
                                                            if ((v42 < 42) || ((18446744073709551615 != 570) || false)) {
                                                                v42 -= v42
                                                                var v43 int8 = (v2 | v2)
                                                                _ = v43
                                                            } else {
                                                                var v44 float64 = (0.1 + 1e10)
                                                                v4 = (v4 % (((v4 + -1)) | 1))
                                                                _ = len(v1)
                                                                _ = v44
                                                            }
                                                            switch fi41 {
                                                            case -2147483648:
                                                                switch (fi24 << ((65535) & 63)) {
                                                                case 42:
                                                                    fmt.Println(fi24)
                                                                case -2147483648:
                                                                    if (2147483647 <= (fi41 ^ (fi24 ^ fi24))) {
                                                                        switch "hello" {
                                                                        case "":
                                                                            if (!(2147483648 <= 997)) {
                                                                                var v45 int64 = 592
                                                                                fmt.Println(v2, fi24, v2)
                                                                                v46 := GenS5{F0: 167}
                                                                                _ = v45
                                                                                _ = v46
                                                                            } else {
                                                                                _ = len(v1)
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(fi24)
                                                                        case "\n\t":
                                                                            switch 4294967295 {
                                                                            case 122:
                                                                                var v47 bool = true
                                                                                switch (fi24 & fi24) {
                                                                                case 766:
                                                                                    var v48 int64 = v42
                                                                                    _ = v48
                                                                                case 128:
                                                                                    switch "" {
                                                                                    case "":
                                                                                        v49 := GenS5{F0: 788}
                                                                                        _ = v49
                                                                                    }
                                                                                case -738:
                                                                                    v2 -= (v2 ^ 0)
                                                                                default:
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    case "foo bar":
                                                                                        v42 = (v42 % ((-9223372036854775808) | 1))
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi41)
                                                                                _ = v47
                                                                            case 42:
                                                                                v50 := GenS4{F0: "\n\t"}
                                                                                _ = v50
                                                                            case 127:
                                                                                v3 = append(v3, 2)
                                                                            }
                                                                        }
                                                                        fmt.Println(v23, v2)
                                                                        v51 := GenS6{F0: v42}
                                                                        _ = v51
                                                                    } else {
                                                                        v52 := GenS7{F0: 699}
                                                                        v42 += 1
                                                                        var v53 uint64 = 1
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                }
                                                            case 0:
                                                                v23 = 1e10
                                                            }
                                                            _ = len(v3)
                                                        case -140:
                                                            var v54 int16 = 0
                                                            _ = v54
                                                        }
                                                        _ = fi41
                                                        _ = v42
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            if (-9223372036854775808 > -914) {
                                                if (!(!true)) {
                                                    switch (^(fi24)) {
                                                    case -9223372036854775808:
                                                        v55 := GenS8{F0: 2}
                                                        _ = v55
                                                    default:
                                                        if false {
                                                            for fi56 := 0; fi56 < 3; fi56++ {
                                                                switch ("" + ("\n\t" + "foo bar")) {
                                                                case "foo bar":
                                                                    switch ("hello" + (("" + "a") + "")) {
                                                                    case "foo bar":
                                                                        v23 = v23
                                                                    }
                                                                case "":
                                                                    if true {
                                                                        var v57 int16 = 32767
                                                                        _ = v57
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v58 float64 = v23
                                                                        _ = v58
                                                                    }
                                                                case "\n\t":
                                                                    _ = len(v1)
                                                                }
                                                                v23 -= (v23 * float64((1e10 - (-2.25 / 3.14159))))
                                                                _ = fi56
                                                            }
                                                            _ = len(v1)
                                                            var v59 float32 = float32((((3.14159 * (float64(0.1) / -2.25)) * float64(-2.25)) - v23))
                                                            _ = v59
                                                        }
                                                    }
                                                    v1 = append(v1, fi24)
                                                } else {
                                                    v4 += (v4 << ((241) & 63))
                                                    fmt.Println(v4, v2)
                                                }
                                                var v60 uint32 = 372
                                                if (!false) {
                                                    for fi61 := 0; fi61 < 1; fi61++ {
                                                        fmt.Println(fi61, v4, fi24)
                                                        switch ("\n\t" + "a") {
                                                        case "\n\t":
                                                            v3 = append(v3, 1)
                                                            fmt.Println(v23)
                                                            if ((v60 / (((v60 | (^(v60)))) | 1)) == v60) {
                                                                v1 = append(v1, fi61)
                                                                v4 = (v4 * v4)
                                                                if true {
                                                                    v2 = v2
                                                                    var v62 int = -9223372036854775808
                                                                    _ = v62
                                                                } else {
                                                                    v60 += (v60 + v60)
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v23 = v23
                                                        default:
                                                            var v63 float32 = (1e10 + 1.5)
                                                            _ = v63
                                                        }
                                                        _ = fi61
                                                    }
                                                    v2 += (v2 - v2)
                                                } else {
                                                    for fi64 := 0; fi64 < 5; fi64++ {
                                                        var v65 bool = false
                                                        v65 = (0.1 >= v23)
                                                        var v66 float32 = ((0.0 - 3.14159) - (3.14159 + (1.5 + -2.25)))
                                                        _ = fi64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                    v67 := GenS1{F0: 0, F1: 42, F2: (true && (255 < 996))}
                                                    var v68 uint16 = 65535
                                                    _ = v67
                                                    _ = v68
                                                }
                                                _ = v60
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if false {
                                                        v2 -= v2
                                                        v2 += (v2 ^ v2)
                                                        if (false && (-33 == int64(fi24))) {
                                                            switch "" {
                                                            case "\n\t":
                                                                v69 := GenS4{F0: ("" + "\n\t")}
                                                                _ = v69
                                                            case "hello":
                                                                var v70 int = fi24
                                                                _ = v70
                                                            default:
                                                                for fi71 := 0; fi71 < 20; fi71++ {
                                                                    if ("a" != (("" + "foo bar") + "hello")) {
                                                                        switch (fi24 % ((fi71) | 1)) {
                                                                        case 127:
                                                                            if false {
                                                                                fmt.Println(v4, fi24)
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 2147483647
                                                                                }
                                                                            }
                                                                        case 17:
                                                                            v2 += (v2 << ((0) & 63))
                                                                        case 0:
                                                                            v23 += v23
                                                                        default:
                                                                            v1 = append(v1, 819)
                                                                        }
                                                                    }
                                                                    var v72 int = -131
                                                                    switch (v72 * v72) {
                                                                    case -374:
                                                                        _ = len(v3)
                                                                    case 946:
                                                                        var v73 int = (fi71 + (fi24 - v72))
                                                                        _ = v73
                                                                    case 268:
                                                                        v1 = append(v1, (fi24 ^ (fi71 >> ((2) & 63))))
                                                                    default:
                                                                        if ((!(1 <= 921)) || (1 != 830)) {
                                                                            v72 = (fi71 << ((241) & 63))
                                                                            fmt.Println(fi71, v23)
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                v23 = float64(0.1)
                                                                                fmt.Println(v23, fi71, fi71)
                                                                                v72 = v72
                                                                            case "hello":
                                                                                switch ("\n\t" + "hello") {
                                                                                case "":
                                                                                    switch v72 {
                                                                                    case 2147483647:
                                                                                        if (v72 <= fi71) {
                                                                                            fmt.Println(v4)
                                                                                            v72 -= v72
                                                                                            v23 = v23
                                                                                        } else {
                                                                                            v74 := GenS9{F0: 255}
                                                                                            switch (v72 - (v72 * 938)) {
                                                                                            case -2147483648:
                                                                                                var v75 uint64 = 617
                                                                                                _ = v75
                                                                                            default:
                                                                                                v2 = (v2 - v2)
                                                                                            }
                                                                                            _ = len(v3)
                                                                                            _ = v74
                                                                                        }
                                                                                    case 127:
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        if false {
                                                                                            var v76 int64 = 1
                                                                                            if false {
                                                                                                var v77 int64 = (v76 ^ -437)
                                                                                                var v78 int64 = v76
                                                                                                v77 = v76
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = 69
                                                                                                }
                                                                                                fmt.Println(v78)
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            } else {
                                                                                                fmt.Println(fi24, fi71)
                                                                                                var v79 uint32 = 2147483648
                                                                                                _ = v79
                                                                                            }
                                                                                            fmt.Println(v23, v76, fi71)
                                                                                            fmt.Println(fi71, v72, v23)
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            v2 = v2
                                                                                            var v80 string = "a"
                                                                                            _ = v76
                                                                                            _ = v80
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = v72
                                                                                            }
                                                                                            var v81 uint = 42
                                                                                            _ = v81
                                                                                        }
                                                                                    case 42:
                                                                                        var v82 uint16 = 243
                                                                                        _ = v82
                                                                                    default:
                                                                                        if false {
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                fmt.Println(v2, fi24)
                                                                                            case "a":
                                                                                                if true {
                                                                                                    v2 += (v2 * 0)
                                                                                                    switch ("" + "\n\t") {
                                                                                                    case "":
                                                                                                        v83 := GenS3{F0: -32768, F1: 71}
                                                                                                        _ = v83
                                                                                                    case "a":
                                                                                                        v2 += (v2 / ((1) | 1))
                                                                                                    case "hello":
                                                                                                        var v84 uint16 = 65535
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    v85 := GenS10{F0: "\n\t", F1: 0}
                                                                                                    _ = v85
                                                                                                } else {
                                                                                                    v4 = (v4 >> ((18446744073709551615) & 63))
                                                                                                    var v86 uint = 18446744073709551615
                                                                                                    if (0 < 2) {
                                                                                                        v87 := GenS11{F0: (!(((3.14159 / 0.1) + (float32(1e10) / 0.1)) < 1.5)), F1: ("\n\t" + ""), F2: (!(!(false || ((0.1 / 1e10) <= (1e10 + 0.0)))))}
                                                                                                        var v88 int = fi24
                                                                                                        var v89 uint16 = 255
                                                                                                        _ = v87
                                                                                                        _ = v88
                                                                                                        _ = v89
                                                                                                    }
                                                                                                    v23 -= (v23 - (0.0 / float64(1e10)))
                                                                                                    var v90 int16 = 1
                                                                                                    var v91 int64 = 1
                                                                                                    _ = v86
                                                                                                    _ = v90
                                                                                                    _ = v91
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    v92 := GenS9{F0: 473}
                                                                                                    if len(v3) >= 2 {
                                                                                                        v3 = v3[1:len(v3)]
                                                                                                    }
                                                                                                    _ = v92
                                                                                                } else {
                                                                                                    v4 = v4
                                                                                                }
                                                                                            }
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v93 int64 = 0
                                                                                            _ = v93
                                                                                        } else {
                                                                                            v94 := GenS8{F0: 524}
                                                                                            v4 = v4
                                                                                            var v95 int16 = 1
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                        }
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v96 uint8 = 1
                                                                                    _ = v96
                                                                                }
                                                                            case "a":
                                                                                v23 = (-2.25 / 1e10)
                                                                            default:
                                                                                var v97 int32 = (v4 & 721)
                                                                                _ = v97
                                                                            }
                                                                        } else {
                                                                            var v98 int16 = 0
                                                                            _ = v98
                                                                        }
                                                                    }
                                                                    v99 := GenS8{F0: 42}
                                                                    if (!false) {
                                                                        _ = len(v3)
                                                                        v100 := GenS12{F0: 0, F1: -1}
                                                                        _ = v100
                                                                    }
                                                                    _ = fi71
                                                                    _ = v72
                                                                    _ = v99
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v23 = float64((0.0 + 1.5))
                                                }
                                                v2 -= 72
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        default:
                                            switch (fi24 / (((fi24 + fi24)) | 1)) {
                                            case -304:
                                                fmt.Println(fi24, v2)
                                            case 32767:
                                                if (!((false && false) || true)) {
                                                    v101 := GenS8{F0: 1}
                                                    _ = v101
                                                }
                                            case 955:
                                                v23 = v23
                                            default:
                                                var v102 uint32 = 65535
                                                _ = v102
                                            }
                                        }
                                    }
                                }
                                var v103 int32 = (^(v4))
                                v103 += 2147483647
                                _ = v103
                            }
                        } else {
                            var v104 uint16 = 2
                            switch (^(fi24)) {
                            case 310:
                                v2 = v2
                            }
                            _ = v104
                        }
                    }
                }
                if (((false || (false && ((3.14159 > v23) && true))) || true) || (!(0 < 42))) {
                    switch "a" {
                    case "":
                        v2 += int8(fi24)
                    case "hello":
                        v3 = append(v3, 2147483647)
                    }
                    v2 = 100
                }
            }
            var v105 uint16 = 206
            _ = fi24
            _ = v105
        }
        var v106 uint16 = 42
        v4 += v4
        _ = v106
    } else {
        v107 := GenS13{F0: 2, F1: "foo bar", F2: 896}
        if (!true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v108 uint16 = 126
            _ = v108
        } else {
            var v109 uint = 357
            v110 := GenS4{F0: ""}
            if (true || true) {
                v1 = append(v1, -558)
                var v111 float64 = float64(((1e10 * (0.0 - 3.14159)) - -2.25))
                _ = v111
            }
            _ = v109
            _ = v110
        }
        _ = v107
    }
    for fi112 := 0; fi112 < 3; fi112++ {
        var v113 int64 = 882
        var v114 uint8 = 0
        _ = fi112
        _ = v113
        _ = v114
    }
    var v115 uint64 = 887
    fmt.Println(v2)
    fmt.Println(v4, v23, v23)
    fmt.Println(v23, v115)
    v116 := GenS14{F0: 0}
    v23 += -2.25
    for fi117 := 0; fi117 < 1; fi117++ {
        var v118 uint64 = (v115 + v115)
        if ((false && true) || (true || (2 > 0))) {
            switch fi117 {
            case -227:
                var v119 int16 = int16(v115)
                _ = v119
            case 1:
                if true {
                    switch 9223372036854775807 {
                    case 127:
                        if len(v3) > 0 {
                            v3[0] = 941
                        }
                    case 459:
                        if ((42 > (v118 ^ (^(v115)))) || false) {
                            switch (fi117 % (((fi117 * fi117)) | 1)) {
                            case -9223372036854775808:
                                var v120 float64 = 3.14159
                                _ = v120
                            case 488:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case 2147483647:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            for fi121 := 0; fi121 < 20; fi121++ {
                                fmt.Println(v118)
                                _ = fi121
                            }
                            for fi122 := 0; fi122 < 1; fi122++ {
                                var v123 uint16 = 380
                                var v124 uint32 = 726
                                _ = fi122
                                _ = v123
                                _ = v124
                            }
                        } else {
                            v125 := GenS2{F0: 9223372036854775807}
                            var v126 bool = (0 > 42)
                            var v127 float64 = float64(1.5)
                            _ = v125
                            _ = v126
                            _ = v127
                        }
                    default:
                        switch "" {
                        case "a":
                            v23 += 0.1
                        case "hello":
                            var v128 uint = 42
                            _ = v128
                        case "":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            fmt.Println(v2)
                        }
                    }
                } else {
                    var v129 uint = 255
                    _ = v129
                }
            default:
                v130 := GenS6{F0: -9223372036854775808}
                _ = v130
            }
            switch fi117 {
            case -145:
                var v131 float64 = float64(1e10)
                _ = v131
            case 127:
                var v132 uint = 385
                _ = v132
            }
        } else {
            var v133 float32 = float32(v23)
            v134 := GenS12{F0: 255, F1: 9223372036854775807}
            for fi135 := 0; fi135 < 5; fi135++ {
                var v136 uint64 = v115
                if (!false) {
                    var v137 uint8 = 42
                    _ = v137
                } else {
                    v136 = v115
                    v136 = 2
                    fmt.Println(v136, v115, v23)
                }
                _ = fi135
                _ = v136
            }
            _ = v133
            _ = v134
        }
        _ = fi117
        _ = v118
    }
    v138 := GenS7{F0: v115}
    v139 := GenS15{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v115:", v115)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v115
    _ = v116
    _ = v138
    _ = v139
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
