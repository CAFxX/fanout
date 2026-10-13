package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int; F1 uint8; F2 int }
type GenS6 struct { F0 int64; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint; F1 uint16; F2 int64 }
type GenS8 struct { F0 int32; F1 uint8; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint8; F2 int16 }

func main() {
    v1 := []int{9223372036854775807, 42, -2147483648}
    var v2 int32 = 42
    var v3 int32 = v2
    v1 = append(v1, -641)
    var v4 int8 = 0
    if true {
        if (((0.1 != float64(0.1)) || (4294967295 >= 721)) || false) {
            switch "\n\t" {
            case "hello":
                v2 += -706
            case "\n\t":
                var v5 uint16 = 1
                _ = v5
            case "foo bar":
                if (2 > 2) {
                    var v6 bool = true
                    for fi7 := 0; fi7 < 19; fi7++ {
                        var v8 bool = v6
                        _ = fi7
                        _ = v8
                    }
                    _ = v6
                } else {
                    v4 += 1
                }
            default:
                v1 = append(v1, -451)
            }
        } else {
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            for fi9 := 0; fi9 < 11; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    if (!(1e10 <= (3.14159 + 3.14159))) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if len(v1) > 0 {
                            v1[0] = fi10
                        }
                    }
                    var v11 float32 = 1e10
                    _ = fi10
                    _ = v11
                }
                v4 = (v4 % ((11) | 1))
                var v12 uint8 = 1
                _ = fi9
                _ = v12
            }
            var v13 uint32 = 65535
            _ = v13
        }
        if false {
            var v14 int32 = (v2 | (v2 << ((42) & 63)))
            if true {
                if true {
                    if false {
                        v1 = append(v1, -414)
                        var v15 bool = false
                        _ = v15
                    } else {
                        var v16 int64 = 1
                        v4 -= -1
                        _ = v16
                    }
                    var v17 uint16 = 396
                    v18 := GenS1{F0: 1}
                    _ = v17
                    _ = v18
                }
                var v19 int8 = v4
                if (v3 < (^(v2))) {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        var v21 uint16 = 692
                        fmt.Println(v21, v3, v19)
                        v4 += v4
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 18; fi22++ {
                        for fi23 := 0; fi23 < 9; fi23++ {
                            var v24 uint64 = 4294967295
                            var v25 uint64 = (^(v24))
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        var v26 uint = 18446744073709551615
                        _ = fi22
                        _ = v26
                    }
                } else {
                    if (0 <= 255) {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            _ = len(v1)
                            _ = fi27
                        }
                    }
                    v4 = -1
                    if (!true) {
                        var v28 int8 = v19
                        fmt.Println(v28, v2, v28)
                        switch "\n\t" {
                        case "\n\t":
                            _ = len(v1)
                        case "":
                            v14 -= 2147483647
                        }
                        _ = v28
                    }
                }
                _ = v19
            } else {
                switch 276 {
                case 42:
                    fmt.Println(v4)
                case 1:
                    fmt.Println(v4, v2, v2)
                case -623:
                    v3 += (v14 | v14)
                }
            }
            _ = v14
        } else {
            var v29 uint64 = 2
            if true {
                if (4294967295 < 0) {
                    var v30 uint16 = 1
                    var v31 uint16 = 32767
                    if len(v1) > 0 {
                        v1[0] = 9223372036854775807
                    }
                    _ = v30
                    _ = v31
                }
                v29 -= v29
            } else {
                v2 += (v2 >> ((2) & 63))
                if (!true) {
                    v32 := GenS2{F0: 1}
                    for fi33 := 0; fi33 < 19; fi33++ {
                        v2 += 912
                        _ = fi33
                    }
                    _ = len(v1)
                    _ = v32
                } else {
                    v3 = -2147483648
                }
                for fi34 := 0; fi34 < 1; fi34++ {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        fmt.Println(fi34)
                        _ = fi35
                    }
                    v36 := GenS1{F0: (fi34 >> ((18446744073709551615) & 63))}
                    _ = fi34
                    _ = v36
                }
            }
            _ = v29
        }
    } else {
        var v37 int = 948
        switch (v37 << ((724) & 63)) {
        case 2147483647:
            var v38 float32 = -2.25
            _ = v38
        case 42:
            v2 = 0
        case 801:
            v39 := GenS2{F0: 1}
            _ = v39
        default:
            v4 += (-(v4))
        }
        if (((v37 | 32767) > (v37 % ((v37) | 1))) && ((-32768 != 209) || true)) {
            for fi40 := 0; fi40 < 13; fi40++ {
                v1 = append(v1, (fi40 % (((fi40 * fi40)) | 1)))
                _ = fi40
            }
            if true {
                for fi41 := 0; fi41 < 18; fi41++ {
                    var v42 int8 = -5
                    _ = fi41
                    _ = v42
                }
                if (0 >= 1) {
                    v43 := GenS1{F0: -786}
                    v4 = (v4 + (v4 - (^(v4))))
                    v4 -= (v4 + v4)
                    _ = v43
                } else {
                    if (!false) {
                        for fi44 := 0; fi44 < 8; fi44++ {
                            if len(v1) > 0 {
                                v1[0] = 128
                            }
                            v3 = 2147483647
                            _ = fi44
                        }
                        switch (v37 % ((v37) | 1)) {
                        case -2147483648:
                            switch "\n\t" {
                            case "":
                                for fi45 := 0; fi45 < 9; fi45++ {
                                    fmt.Println(v37, fi45, v3)
                                    var v46 string = ""
                                    _ = fi45
                                    _ = v46
                                }
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = 847
                                }
                            }
                        case 0:
                            var v47 int32 = (v2 / ((v3) | 1))
                            _ = v47
                        default:
                            v48 := GenS3{F0: -745, F1: 4294967295}
                            _ = v48
                        }
                        v3 += (v3 & 88)
                    }
                    switch "" {
                    case "":
                        if len(v1) > 0 {
                            v1[0] = v37
                        }
                    case "foo bar":
                        if (0.0 <= ((0.1 - (1.5 / 1.0)) * -2.25)) {
                            switch "\n\t" {
                            case "hello":
                                var v49 int32 = (v2 << ((526) & 63))
                                _ = v49
                            case "a":
                                if true {
                                    fmt.Println(v37, v3)
                                    var v50 string = "\n\t"
                                    _ = v50
                                }
                            case "\n\t":
                                var v51 int = (-(v37))
                                _ = v51
                            default:
                                v3 = v2
                            }
                        }
                    case "a":
                        v52 := GenS4{F0: 704, F1: 42}
                        _ = v52
                    }
                }
                var v53 int = (v37 - v37)
                _ = v53
            } else {
                for fi54 := 0; fi54 < 7; fi54++ {
                    switch "" {
                    case "hello":
                        switch "" {
                        case "":
                            var v55 bool = (-1 <= 9223372036854775807)
                            _ = v55
                        case "foo bar":
                            v2 -= v2
                        case "\n\t":
                            v37 = (-(v37))
                        default:
                            v2 = (-(v2))
                        }
                    case "foo bar":
                        v56 := GenS2{F0: 2}
                        _ = v56
                    case "\n\t":
                        switch ("\n\t" + ("a" + "a")) {
                        case "hello":
                            for fi57 := 0; fi57 < 1; fi57++ {
                                var v58 int16 = 0
                                v59 := GenS3{F0: v58, F1: 4294967295}
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                        case "a":
                            v60 := GenS5{F0: v37, F1: 2, F2: v37}
                            _ = v60
                        case "\n\t":
                            if ("\n\t" >= "hello") {
                                var v61 int = fi54
                                if true {
                                    fmt.Println(v61, v3)
                                    for fi62 := 0; fi62 < 17; fi62++ {
                                        fmt.Println(v4, v2)
                                        var v63 uint16 = 1
                                        var v64 uint = 363
                                        if true {
                                            v37 += fi62
                                        }
                                        var v65 uint16 = v63
                                        _ = fi62
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    }
                                } else {
                                    if true {
                                        v61 += (fi54 << ((164) & 63))
                                    } else {
                                        if true {
                                            v66 := GenS5{F0: fi54, F1: 128, F2: (fi54 + (v61 - -362))}
                                            v67 := GenS2{F0: 42}
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                }
                                _ = v61
                            } else {
                                var v68 float32 = (0.0 - ((1e10 * 1.5) + 0.0))
                                for fi69 := 0; fi69 < 10; fi69++ {
                                    if ("a" <= "a") {
                                        if len(v1) > 0 {
                                            v1[0] = v37
                                        }
                                        _ = len(v1)
                                    } else {
                                        var v70 int64 = int64(fi54)
                                        _ = v70
                                    }
                                    _ = fi69
                                }
                                fmt.Println(v68)
                                _ = v68
                            }
                        default:
                            for fi71 := 0; fi71 < 9; fi71++ {
                                var v72 int16 = -1
                                _ = fi71
                                _ = v72
                            }
                        }
                    default:
                        if true {
                            if false {
                                if false {
                                    fmt.Println(v37, v4, v4)
                                    v73 := GenS3{F0: -32768, F1: 0}
                                    for fi74 := 0; fi74 < 7; fi74++ {
                                        v2 = (v3 & v2)
                                        _ = fi74
                                    }
                                    _ = v73
                                }
                                v3 += v2
                                switch (fi54 * v37) {
                                case 880:
                                    for fi75 := 0; fi75 < 2; fi75++ {
                                        v37 = (^(fi75))
                                        _ = fi75
                                    }
                                }
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (fi54 + v37)
                                }
                                v76 := GenS5{F0: 42, F1: 255, F2: v37}
                                v37 -= 385
                                _ = v76
                            }
                            if (65535 <= 4294967295) {
                                v77 := GenS6{F0: 42, F1: 128, F2: 42}
                                v4 -= -40
                                v78 := GenS7{F0: 585, F1: 1, F2: 749}
                                _ = v77
                                _ = v78
                            }
                            for fi79 := 0; fi79 < 7; fi79++ {
                                v37 += 128
                                _ = fi79
                            }
                        } else {
                            switch "foo bar" {
                            case "a":
                                if true {
                                    v37 = v37
                                }
                            case "hello":
                                switch -924 {
                                case -2147483648:
                                    switch -9223372036854775808 {
                                    case -839:
                                        var v80 bool = false
                                        _ = v80
                                    case 9223372036854775807:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case -1:
                                        for fi81 := 0; fi81 < 1; fi81++ {
                                            v2 = v3
                                            _ = len(v1)
                                            if len(v1) > 0 {
                                                v1[0] = -516
                                            }
                                            v4 -= (^(v4))
                                            if ((2 >= 828) && false) {
                                                if true {
                                                    var v82 uint8 = 2
                                                    if ((uint64(v82) != 0) && false) {
                                                        v2 -= v3
                                                        v3 += (-(v2))
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            switch 32768 {
                                                            case 2147483647:
                                                                var v83 int16 = 1
                                                                _ = v83
                                                            case 42:
                                                                var v84 int16 = 1
                                                                _ = v84
                                                            }
                                                        case "\n\t":
                                                            _ = len(v1)
                                                        default:
                                                            v1 = append(v1, 2147483647)
                                                        }
                                                        if (true && (538 == 0)) {
                                                            fmt.Println(v4, v3)
                                                        } else {
                                                            v1 = append(v1, fi54)
                                                            var v85 string = ""
                                                            v37 += (fi54 - 127)
                                                            v37 = 128
                                                            _ = v85
                                                        }
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = 2147483647
                                                    }
                                                    _ = v82
                                                } else {
                                                    v3 += (^(v3))
                                                    _ = len(v1)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        if false {
                                                            if false {
                                                                if false {
                                                                    var v86 int = v37
                                                                    _ = len(v1)
                                                                    var v87 int64 = 9223372036854775807
                                                                    v87 -= 981
                                                                    v88 := GenS8{F0: v3, F1: 128, F2: 4294967295}
                                                                    _ = v86
                                                                    _ = v87
                                                                    _ = v88
                                                                } else {
                                                                    var v89 uint8 = 127
                                                                    var v90 int16 = 1
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                                var v91 int32 = (v2 / (((v3 ^ (-(v2)))) | 1))
                                                                _ = v91
                                                            }
                                                        }
                                                    }
                                                }
                                                switch v37 {
                                                case -1:
                                                    v1 = append(v1, fi54)
                                                case -232:
                                                    v1 = append(v1, (fi81 * fi81))
                                                case -2147483648:
                                                    v1 = append(v1, 2147483647)
                                                    switch "hello" {
                                                    case "a":
                                                        v92 := GenS9{F0: 18446744073709551615}
                                                        _ = v92
                                                    case "":
                                                        if (!(0.1 >= -2.25)) {
                                                            var v93 bool = false
                                                            var v94 uint = 85
                                                            _ = v93
                                                            _ = v94
                                                        } else {
                                                            if (0 > 4294967295) {
                                                                v95 := GenS9{F0: 0}
                                                                switch -1 {
                                                                case -2147483648:
                                                                    switch (^(fi81)) {
                                                                    case 32767:
                                                                        switch v37 {
                                                                        case -43:
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                v1 = append(v1, 4294967295)
                                                                                var v96 string = ("hello" + "\n\t")
                                                                                _ = v96
                                                                            default:
                                                                                v3 -= (-(v2))
                                                                            }
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v2, fi54, v4)
                                                                        }
                                                                    case 2147483647:
                                                                        v97 := GenS1{F0: 2147483647}
                                                                        _ = v97
                                                                    default:
                                                                        v4 += v4
                                                                    }
                                                                }
                                                                _ = v95
                                                            }
                                                        }
                                                    default:
                                                        v37 += 1
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v1)
                                                default:
                                                    fmt.Println(v37)
                                                }
                                            } else {
                                                v3 = (v2 * v3)
                                                fmt.Println(v4, v37, fi81)
                                            }
                                            _ = fi81
                                        }
                                    }
                                }
                            case "\n\t":
                                v2 = (^(v3))
                            }
                            v4 -= 46
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = (v37 % ((2147483647) | 1))
                    }
                    v98 := GenS2{F0: 127}
                    _ = fi54
                    _ = v98
                }
            }
            var v99 string = "hello"
            _ = v99
        }
        _ = v37
    }
    var v100 int = 42
    if true {
        for fi101 := 0; fi101 < 3; fi101++ {
            var v102 bool = false
            _ = fi101
            _ = v102
        }
        fmt.Println(v100)
    } else {
        var v103 int64 = 42
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        _ = v103
    }
    v100 += v100
    if (2 > 524) {
        v104 := GenS10{F0: 1000, F1: 42, F2: 756}
        _ = v104
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = len(v1)
        fmt.Println(v4, v2)
    }
    v3 += (-(v2))
    fmt.Println(v100, v100, v2)
    v105 := GenS5{F0: (v100 % ((-1) | 1)), F1: 0, F2: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v105
    _ = v2
    _ = v3
    _ = v4
}
