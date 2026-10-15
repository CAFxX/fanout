package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint32; F1 uint; F2 uint64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int; F2 uint8 }
type GenS6 struct { F0 int64; F1 int8; F2 int }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 uint16; F1 int32 }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 int; F1 uint; F2 int64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int32; F1 uint64; F2 string }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int16; F1 int64 }
type GenS16 struct { F0 uint; F1 bool }
type GenS17 struct { F0 bool; F1 int32; F2 uint64 }
type GenS18 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 int8; F1 uint; F2 bool }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 int16; F1 string }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint; F1 uint16; F2 int8 }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 uint32; F1 uint }
type GenS28 struct { F0 bool; F1 bool }
type GenS29 struct { F0 int64; F1 uint32 }
type GenS30 struct { F0 int64 }
type GenS31 struct { F0 int32; F1 int32; F2 bool }
type GenS32 struct { F0 int16; F1 int; F2 int }
func genFn0() int {
    var v5 int64 = -218
    v6 := []int{834, 2147483647, 9223372036854775807}
    var v7 int = -9223372036854775808
    _ = v5
    _ = v6
    _ = v7
    return (v7 + v7)
}
func genFn1(p0 int16, p1 uint16) int8 {
    p1 = 42
    p1 = p1
    v8 := []int{genFn0()}
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: fi9, F1: (fi9 % ((32767) | 1))}
        if (!false) {
            p1 = (^(p1))
            p1 = (p1 | (p1 % (((p1 | 32767)) | 1)))
            for fi11 := 0; fi11 < 13; fi11++ {
                v8 = append(v8, (fi11 % (((-(fi11))) | 1)))
                p1 = (p1 & 42)
                var v12 int = fi11
                _ = fi11
                _ = v12
            }
        }
        v13 := GenS2{F0: 2147483647, F1: 1, F2: 101}
        _ = fi9
        _ = v10
        _ = v13
    }
    if (0 != -929) {
        v14 := GenS3{F0: (("foo bar" + "\n\t") + "a"), F1: p1}
        for fi15 := 0; fi15 < 1; fi15++ {
            if true {
                if (832 >= 42) {
                    if false {
                        var v16 int16 = (p0 | p0)
                        if false {
                            fmt.Println(v16, p0, v16)
                        }
                        v16 = (p0 % (((v16 >> ((0) & 63))) | 1))
                        _ = v16
                    } else {
                        if false {
                            if true {
                                if true {
                                    p0 = int16(fi15)
                                } else {
                                    if (500 >= 42) {
                                        var v17 uint32 = 223
                                        var v18 int32 = -2147483648
                                        _ = len(v8)
                                        _ = v17
                                        _ = v18
                                    } else {
                                        p0 = 125
                                    }
                                }
                                _ = len(v8)
                            }
                            for fi19 := 0; fi19 < 9; fi19++ {
                                var v20 int16 = p0
                                _ = fi19
                                _ = v20
                            }
                            switch "" {
                            case "a":
                                if true {
                                    p1 = (p1 & (^(p1)))
                                    if false {
                                        if true {
                                            var v21 int32 = 1
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v21
                                        } else {
                                            p0 += (-(p0))
                                            fmt.Println(p0, p1, p1)
                                            var v22 int = -464
                                            _ = v22
                                        }
                                    }
                                    var v23 uint16 = 928
                                    _ = v23
                                } else {
                                    fmt.Println(p0, p0)
                                }
                            case "\n\t":
                                var v24 uint = 65535
                                _ = v24
                            case "hello":
                                fmt.Println(p1)
                            default:
                                switch (^(fi15)) {
                                case -2147483648:
                                    _ = len(v8)
                                case 127:
                                    for fi25 := 0; fi25 < 11; fi25++ {
                                        switch ("hello" + ("\n\t" + "hello")) {
                                        case "foo bar":
                                            if true {
                                                v26 := GenS4{F0: -71}
                                                _ = v26
                                            } else {
                                                p0 = (-(p0))
                                            }
                                        case "hello":
                                            p1 += (p1 | p1)
                                        default:
                                            if true {
                                                var v27 uint = 398
                                                v8 = append(v8, genFn0())
                                                _ = v27
                                            }
                                        }
                                        v8 = append(v8, 127)
                                        _ = fi25
                                    }
                                case -365:
                                    if (false && ((!(p1 <= (p1 / ((2) | 1)))) || (!false))) {
                                        switch (fi15 % ((fi15) | 1)) {
                                        case 666:
                                            if true {
                                                var v28 float32 = 3.14159
                                                v29 := GenS2{F0: 0, F1: 862, F2: 2}
                                                switch "a" {
                                                case "":
                                                    if true {
                                                        if (!(!(!(v28 >= (v28 - float32((3.14159 + 0.0))))))) {
                                                            v30 := GenS2{F0: 362, F1: 895, F2: 4294967295}
                                                            if (!((true || false) && false)) {
                                                                v28 = v28
                                                                switch fi15 {
                                                                case 0:
                                                                    if (458 >= 4294967295) {
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            var v31 int = (fi15 + (^(fi15)))
                                                                            _ = v31
                                                                        case "":
                                                                            var v32 int16 = (p0 >> ((255) & 63))
                                                                            _ = v32
                                                                        case "foo bar":
                                                                            if len(v8) >= 2 {
                                                                                v8 = v8[1:len(v8)]
                                                                            }
                                                                        }
                                                                        if false {
                                                                            var v33 int64 = -855
                                                                            _ = v33
                                                                        } else {
                                                                            for fi34 := 0; fi34 < 19; fi34++ {
                                                                                var v35 int8 = 42
                                                                                var v36 float32 = (float32((0.0 + 1.5)) - float32(1.5))
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = (fi15 - (fi34 | 15))
                                                                                }
                                                                                var v37 int = 4294967295
                                                                                _ = fi34
                                                                                _ = v35
                                                                                _ = v36
                                                                                _ = v37
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    for fi38 := 0; fi38 < 15; fi38++ {
                                                                        var v39 uint16 = p1
                                                                        _ = fi38
                                                                        _ = v39
                                                                    }
                                                                }
                                                                v40 := GenS5{F0: 1, F1: fi15, F2: 128}
                                                                _ = v40
                                                            }
                                                            _ = v30
                                                        }
                                                        v41 := GenS6{F0: 551, F1: -34, F2: (fi15 * (fi15 << ((18446744073709551615) & 63)))}
                                                        v42 := GenS4{F0: 58}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                default:
                                                    if (!false) {
                                                        v43 := GenS1{F0: (fi15 & (fi15 - 128)), F1: (fi15 + (-(fi15)))}
                                                        _ = v43
                                                    }
                                                }
                                                _ = v28
                                                _ = v29
                                            } else {
                                                p0 = 42
                                                v44 := GenS4{F0: -87}
                                                _ = v44
                                            }
                                        case -150:
                                            v45 := GenS5{F0: 2, F1: fi15, F2: 237}
                                            _ = v45
                                        default:
                                            v46 := GenS7{F0: 127, F1: -1}
                                            _ = v46
                                        }
                                        switch "a" {
                                        case "a":
                                            if true {
                                                for fi47 := 0; fi47 < 11; fi47++ {
                                                    switch "a" {
                                                    case "a":
                                                        v8 = append(v8, (fi47 / ((fi47) | 1)))
                                                    case "hello":
                                                        if false {
                                                            switch "hello" {
                                                            case "a":
                                                                var v48 int = (fi47 % ((fi15) | 1))
                                                                _ = v48
                                                            case "\n\t":
                                                                var v49 uint64 = 0
                                                                _ = v49
                                                            case "":
                                                                v8 = append(v8, genFn0())
                                                            }
                                                            p0 -= (p0 ^ p0)
                                                            p0 += (p0 + (p0 / ((42) | 1)))
                                                        }
                                                    }
                                                    _ = fi47
                                                }
                                            }
                                        case "\n\t":
                                            v50 := GenS2{F0: uint32(p1), F1: 4294967295, F2: uint64(p1)}
                                            _ = v50
                                        case "":
                                            if len(v8) > 0 {
                                                v8[0] = -771
                                            }
                                        }
                                    } else {
                                        if false {
                                            p1 -= (p1 | p1)
                                            fmt.Println(fi15, fi15)
                                            v51 := GenS3{F0: "\n\t", F1: p1}
                                            _ = v51
                                        } else {
                                            if false {
                                                for fi52 := 0; fi52 < 10; fi52++ {
                                                    switch (("a" + "a") + "a") {
                                                    case "":
                                                        if false {
                                                            p0 = 430
                                                            fmt.Println(fi15)
                                                        } else {
                                                            var v53 int32 = -413
                                                            _ = v53
                                                        }
                                                    case "foo bar":
                                                        var v54 uint64 = 18446744073709551615
                                                        _ = v54
                                                    default:
                                                        p1 = p1
                                                    }
                                                    _ = fi52
                                                }
                                                var v55 uint16 = (p1 & (p1 >> ((328) & 63)))
                                                _ = len(v8)
                                                _ = v55
                                            }
                                        }
                                        fmt.Println(fi15)
                                        p1 -= 0
                                    }
                                default:
                                    switch "" {
                                    case "":
                                        var v56 uint = 255
                                        _ = v56
                                    case "\n\t":
                                        if (1 != 2147483647) {
                                            switch (^(fi15)) {
                                            case 42:
                                                if true {
                                                    v57 := GenS8{F0: p1, F1: -1}
                                                    _ = v57
                                                }
                                            case 4294967295:
                                                var v58 uint32 = 65535
                                                _ = v58
                                            case 32768:
                                                fmt.Println(fi15, fi15)
                                            default:
                                                var v59 uint = 4294967295
                                                _ = v59
                                            }
                                            var v60 uint64 = 356
                                            for fi61 := 0; fi61 < 18; fi61++ {
                                                if ("" >= "a") {
                                                    if (127 > 71) {
                                                        fmt.Println(v60, p0)
                                                        p1 += uint16(v60)
                                                    } else {
                                                        if len(v8) >= 2 {
                                                            v8 = v8[1:len(v8)]
                                                        }
                                                        p0 = p0
                                                        fmt.Println(v60)
                                                    }
                                                }
                                                v60 += v60
                                                _ = fi61
                                            }
                                            _ = v60
                                        }
                                    case "foo bar":
                                        v62 := GenS1{F0: 128, F1: 4294967295}
                                        _ = v62
                                    default:
                                        if true {
                                            var v63 int64 = -9223372036854775808
                                            var v64 string = "\n\t"
                                            _ = v63
                                            _ = v64
                                        } else {
                                            switch ("a" + "a") {
                                            case "hello":
                                                if false {
                                                    fmt.Println(p1, p1)
                                                    fmt.Println(fi15, fi15, fi15)
                                                }
                                            case "":
                                                v65 := GenS9{F0: 1, F1: 2}
                                                _ = v65
                                            default:
                                                if false {
                                                    p0 = p0
                                                    v66 := GenS5{F0: 42, F1: fi15, F2: 42}
                                                    _ = v66
                                                }
                                            }
                                        }
                                    }
                                }
                            }
                        } else {
                            if len(v8) > 0 {
                                v8[0] = fi15
                            }
                            _ = len(v8)
                            if (2 >= 2147483647) {
                                for fi67 := 0; fi67 < 18; fi67++ {
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            if false {
                                                var v68 bool = false
                                                fmt.Println(p0, fi67, p1)
                                                _ = v68
                                            } else {
                                                _ = len(v8)
                                            }
                                        }
                                    }
                                    var v69 float64 = 1.5
                                    _ = fi67
                                    _ = v69
                                }
                                v70 := GenS2{F0: 42, F1: 42, F2: 42}
                                switch fi15 {
                                case 413:
                                    for fi71 := 0; fi71 < 20; fi71++ {
                                        p0 = p0
                                        v8 = append(v8, (fi71 - fi15))
                                        _ = fi71
                                    }
                                case 128:
                                    var v72 uint16 = 931
                                    _ = v72
                                case 9223372036854775807:
                                    if (184 < 64) {
                                        var v73 string = (("\n\t" + "a") + ("" + "hello"))
                                        p0 -= p0
                                        _ = v73
                                    } else {
                                        switch (("hello" + ((((("a" + "") + (("hello" + ("\n\t" + (("foo bar" + "a") + "\n\t"))) + ("\n\t" + "a"))) + (("foo bar" + ("" + "\n\t")) + "a")) + "a") + "")) + "a") {
                                        case "foo bar":
                                            v8 = append(v8, 128)
                                        case "hello":
                                            v74 := GenS2{F0: 52, F1: 255, F2: 42}
                                            _ = v74
                                        case "":
                                            if len(v8) > 0 {
                                                v8[0] = fi15
                                            }
                                        default:
                                            var v75 uint64 = 2
                                            _ = v75
                                        }
                                        if true {
                                            if true {
                                                fmt.Println(fi15, fi15)
                                            } else {
                                                for fi76 := 0; fi76 < 9; fi76++ {
                                                    fmt.Println(p1)
                                                    p1 = 255
                                                    var v77 uint8 = 92
                                                    v77 -= (v77 * v77)
                                                    _ = fi76
                                                    _ = v77
                                                }
                                                var v78 string = ""
                                                v79 := GenS3{F0: (v78 + v78), F1: p1}
                                                _ = v78
                                                _ = v79
                                            }
                                        } else {
                                            for fi80 := 0; fi80 < 5; fi80++ {
                                                if true {
                                                    p1 = p1
                                                    var v81 uint8 = 127
                                                    v8 = append(v8, fi15)
                                                    _ = v81
                                                } else {
                                                    var v82 int64 = -699
                                                    _ = v82
                                                }
                                                switch genFn0() {
                                                case -845:
                                                    if true {
                                                        fmt.Println(fi15, fi15)
                                                        p1 = 32767
                                                    }
                                                case -3:
                                                    fmt.Println(fi80, p1)
                                                }
                                                p1 = p1
                                                _ = fi80
                                            }
                                            v83 := GenS6{F0: 1, F1: 100, F2: fi15}
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            _ = v83
                                        }
                                        if (1e10 > 1e10) {
                                            fmt.Println(p0, p0, fi15)
                                            if false {
                                                v84 := GenS3{F0: ((((("foo bar" + ("\n\t" + "foo bar")) + "") + ("\n\t" + "\n\t")) + ("" + "\n\t")) + "hello"), F1: p1}
                                                _ = v84
                                            } else {
                                                fmt.Println(p0, p1, p1)
                                                v85 := GenS8{F0: 65535, F1: 1}
                                                _ = v85
                                            }
                                        } else {
                                            _ = len(v8)
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            v86 := GenS10{F0: (fi15 >> ((147) & 63)), F1: 255, F2: 9223372036854775807}
                                            _ = v86
                                        }
                                    }
                                default:
                                    switch "a" {
                                    case "a":
                                        var v87 uint8 = 128
                                        _ = v87
                                    case "\n\t":
                                        switch fi15 {
                                        case 530:
                                            v88 := GenS8{F0: p1, F1: 0}
                                            _ = v88
                                        case 0:
                                            switch fi15 {
                                            case 1:
                                                _ = len(v8)
                                            case 9223372036854775807:
                                                _ = len(v8)
                                            case -1:
                                                p1 = (p1 | p1)
                                            default:
                                                p0 += (p0 ^ p0)
                                            }
                                        case 42:
                                            var v89 float64 = (0.1 + -2.25)
                                            _ = v89
                                        default:
                                            v90 := GenS2{F0: 2147483648, F1: 4294967295, F2: 4294967295}
                                            _ = v90
                                        }
                                    default:
                                        for fi91 := 0; fi91 < 9; fi91++ {
                                            if len(v8) > 0 {
                                                v8[0] = (fi91 - fi91)
                                            }
                                            _ = fi91
                                        }
                                    }
                                }
                                _ = v70
                            } else {
                                fmt.Println(p0, fi15, fi15)
                                v92 := GenS2{F0: 2147483647, F1: 4294967295, F2: 283}
                                _ = v92
                            }
                        }
                        for fi93 := 0; fi93 < 14; fi93++ {
                            if true {
                                fmt.Println(fi93, fi93)
                                if false {
                                    if false {
                                        var v94 int64 = -1
                                        _ = v94
                                    } else {
                                        p1 -= 494
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = (^(fi15))
                                    }
                                } else {
                                    switch (fi15 - fi15) {
                                    case 220:
                                        var v95 int16 = (p0 / ((1000) | 1))
                                        _ = v95
                                    case 258:
                                        v8 = append(v8, fi15)
                                    case 127:
                                        _ = len(v8)
                                    default:
                                        fmt.Println(fi15, fi15, fi93)
                                    }
                                }
                                switch genFn0() {
                                case 2147483647:
                                    if (false || (!(true || false))) {
                                        var v96 int64 = int64(fi93)
                                        v97 := GenS8{F0: 1, F1: -1}
                                        var v98 string = ""
                                        _ = v96
                                        _ = v97
                                        _ = v98
                                    } else {
                                        var v99 string = "foo bar"
                                        switch v99 {
                                        case "":
                                            fmt.Println(p1)
                                        case "\n\t":
                                            var v100 uint32 = 1
                                            _ = v100
                                        case "foo bar":
                                            var v101 float64 = 0.0
                                            _ = v101
                                        default:
                                            if (325 != 18446744073709551615) {
                                                switch v99 {
                                                case "a":
                                                    var v102 uint16 = p1
                                                    _ = v102
                                                case "hello":
                                                    if (false || true) {
                                                        if (-9223372036854775808 == -580) {
                                                            var v103 float32 = (((0.0 / 1.0) - 0.0) / 3.14159)
                                                            p1 = p1
                                                            var v104 int = (fi15 | fi15)
                                                            p0 += p0
                                                            _ = len(v8)
                                                            switch v104 {
                                                            case 4294967295:
                                                                fmt.Println(v103)
                                                            default:
                                                                v105 := GenS11{F0: 42}
                                                                _ = v105
                                                            }
                                                            if len(v8) > 0 {
                                                                v8[0] = -2147483648
                                                            }
                                                            var v106 uint64 = 2
                                                            if ((v106 ^ v106) == (^(v106))) {
                                                                _ = len(v8)
                                                                v106 -= (v106 % ((0) | 1))
                                                                if (-1 <= 9223372036854775807) {
                                                                    v106 = v106
                                                                }
                                                            }
                                                            switch (((v99 + v99) + (v99 + v99)) + "a") {
                                                            case "\n\t":
                                                                if ((p1 << ((18446744073709551615) & 63)) <= p1) {
                                                                    v99 = (v99 + "foo bar")
                                                                    if ((65535 >= 18446744073709551615) || (false && (1000 == (p0 << ((1) & 63))))) {
                                                                        switch 127 {
                                                                        case 4294967295:
                                                                            if false {
                                                                                v107 := GenS6{F0: -1, F1: 0, F2: v104}
                                                                                switch ("hello" + (v99 + v99)) {
                                                                                case "":
                                                                                    if true {
                                                                                        if ((v104 << ((362) & 63)) >= (v104 & (^(fi93)))) {
                                                                                            var v108 float64 = 1.5
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                _ = len(v8)
                                                                                            case "":
                                                                                                if false {
                                                                                                    if true {
                                                                                                        switch genFn0() {
                                                                                                        case 128:
                                                                                                            fmt.Println(fi93, v103, fi93)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v103 = v103
                                                                                                        v108 = (3.14159 - float64(float32(v108)))
                                                                                                        fmt.Println(p0, v108)
                                                                                                        v109 := GenS9{F0: (v106 + 18446744073709551615), F1: 150}
                                                                                                        _ = v109
                                                                                                    }
                                                                                                }
                                                                                            case "a":
                                                                                                if len(v8) > 0 {
                                                                                                    v8[0] = (fi93 >> ((891) & 63))
                                                                                                }
                                                                                            default:
                                                                                                v103 = (v103 / v103)
                                                                                            }
                                                                                            if (!true) {
                                                                                                switch -504 {
                                                                                                case -2147483648:
                                                                                                    v106 += v106
                                                                                                default:
                                                                                                    p1 -= p1
                                                                                                }
                                                                                                fmt.Println(fi15, v103)
                                                                                                p0 -= (p0 ^ (-(p0)))
                                                                                                fmt.Println(v106, v103)
                                                                                                var v110 int64 = 1
                                                                                                _ = v110
                                                                                            }
                                                                                            _ = v108
                                                                                        } else {
                                                                                            switch fi15 {
                                                                                            case 4294967295:
                                                                                                var v111 uint16 = (p1 * (p1 << ((2) & 63)))
                                                                                                v112 := GenS12{F0: -600, F1: v106, F2: "foo bar"}
                                                                                                if false {
                                                                                                    if len(v8) > 0 {
                                                                                                        v8[0] = fi93
                                                                                                    }
                                                                                                    v103 = ((0.1 + v103) / v103)
                                                                                                }
                                                                                                if true {
                                                                                                    if false {
                                                                                                        var v113 uint64 = v106
                                                                                                        v103 -= v103
                                                                                                        if true {
                                                                                                            fmt.Println(p1, fi93, v106)
                                                                                                            v114 := GenS8{F0: p1, F1: 0}
                                                                                                            _ = v114
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v113 = v106
                                                                                                                v104 = (fi15 - (fi15 | (-(v104))))
                                                                                                                v115 := GenS13{F0: 42}
                                                                                                                var v116 float64 = (-2.25 - 3.14159)
                                                                                                                _ = v115
                                                                                                                _ = v116
                                                                                                            } else {
                                                                                                                v117 := GenS13{F0: -224}
                                                                                                                if true {
                                                                                                                    p1 = 2
                                                                                                                } else {
                                                                                                                    v118 := GenS14{F0: p0}
                                                                                                                    var v119 int32 = -2147483648
                                                                                                                    _ = v118
                                                                                                                    _ = v119
                                                                                                                }
                                                                                                                _ = v117
                                                                                                            }
                                                                                                            switch (fi93 - fi15) {
                                                                                                            case 148:
                                                                                                                if (-2.25 == 0.0) {
                                                                                                                    v120 := GenS15{F0: (p0 * -32768), F1: 1}
                                                                                                                    _ = v120
                                                                                                                }
                                                                                                            default:
                                                                                                                if len(v8) > 0 {
                                                                                                                    v8[0] = 127
                                                                                                                }
                                                                                                            }
                                                                                                            var v121 uint64 = v106
                                                                                                            _ = v121
                                                                                                        }
                                                                                                        _ = v113
                                                                                                    } else {
                                                                                                        if (0 >= 1) {
                                                                                                            fmt.Println(v99, v99)
                                                                                                            switch "a" {
                                                                                                            case "":
                                                                                                                switch v99 {
                                                                                                                case "\n\t":
                                                                                                                    v122 := GenS16{F0: uint(v111), F1: (-1 >= -63)}
                                                                                                                    fmt.Println(p0)
                                                                                                                    p1 = (^(p1))
                                                                                                                    _ = v122
                                                                                                                case "foo bar":
                                                                                                                    switch "" {
                                                                                                                    case "a":
                                                                                                                        if (((4294967295 > 67) || (v106 <= 1)) || true) {
                                                                                                                            var v123 int64 = -1
                                                                                                                            if (false && false) {
                                                                                                                                _ = len(v8)
                                                                                                                            }
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v99, p0)
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    if (false || false) {
                                                                                                                        switch genFn0() {
                                                                                                                        case 1:
                                                                                                                            var v124 string = ("" + v99)
                                                                                                                            _ = v124
                                                                                                                        case 977:
                                                                                                                            v125 := GenS17{F0: false, F1: 1, F2: v106}
                                                                                                                            _ = v125
                                                                                                                        case 127:
                                                                                                                            v126 := GenS14{F0: (^(p0))}
                                                                                                                            _ = v126
                                                                                                                        default:
                                                                                                                            if false {
                                                                                                                                v99 = ""
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v106 -= v106
                                                                                                                        if (65535 < 2) {
                                                                                                                            if ((!(0 <= 42)) || (-1 < -396)) {
                                                                                                                                switch 4294967295 {
                                                                                                                                case 128:
                                                                                                                                    if true {
                                                                                                                                        if false {
                                                                                                                                            v99 = ("\n\t" + v99)
                                                                                                                                            _ = len(v8)
                                                                                                                                        }
                                                                                                                                        _ = len(v8)
                                                                                                                                    } else {
                                                                                                                                        p1 += uint16(v106)
                                                                                                                                        var v127 int = fi15
                                                                                                                                        _ = v127
                                                                                                                                    }
                                                                                                                                case 2147483647:
                                                                                                                                    var v128 int64 = -9223372036854775808
                                                                                                                                    _ = v128
                                                                                                                                default:
                                                                                                                                    if false {
                                                                                                                                        p1 = 2
                                                                                                                                        if ("a" < (v99 + v99)) {
                                                                                                                                            v8 = append(v8, v104)
                                                                                                                                            p0 = p0
                                                                                                                                            if false {
                                                                                                                                                var v129 uint64 = 0
                                                                                                                                                _ = v129
                                                                                                                                            } else {
                                                                                                                                                if false {
                                                                                                                                                    var v130 uint32 = 42
                                                                                                                                                    var v131 uint32 = (v130 | 2)
                                                                                                                                                    _ = v130
                                                                                                                                                    _ = v131
                                                                                                                                                }
                                                                                                                                                v8 = append(v8, 95)
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v132 string = v99
                                                                                                                                        _ = v132
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = len(v8)
                                                                                                                            var v133 int = genFn0()
                                                                                                                            _ = v133
                                                                                                                        }
                                                                                                                        if (135 < (v111 >> ((65535) & 63))) {
                                                                                                                            if true {
                                                                                                                                if (true || false) {
                                                                                                                                    v8 = append(v8, (v104 / ((fi15) | 1)))
                                                                                                                                    v134 := GenS14{F0: -738}
                                                                                                                                    fmt.Println(p1, v111)
                                                                                                                                    _ = v134
                                                                                                                                }
                                                                                                                                var v135 float64 = (((-2.25 - 1.5) * 3.14159) / 3.14159)
                                                                                                                                _ = v135
                                                                                                                            } else {
                                                                                                                                var v136 int16 = p0
                                                                                                                                _ = v136
                                                                                                                            }
                                                                                                                            v137 := GenS18{F0: 255, F1: 1, F2: 1}
                                                                                                                            var v138 int8 = -98
                                                                                                                            var v139 uint16 = (v111 / (((p1 - (p1 * v111))) | 1))
                                                                                                                            if false {
                                                                                                                                var v140 uint8 = uint8(p1)
                                                                                                                                if len(v8) > 0 {
                                                                                                                                    v8[0] = fi93
                                                                                                                                }
                                                                                                                                v99 = v99
                                                                                                                                _ = v140
                                                                                                                            }
                                                                                                                            _ = v137
                                                                                                                            _ = v138
                                                                                                                            _ = v139
                                                                                                                        }
                                                                                                                        if len(v8) > 0 {
                                                                                                                            v8[0] = fi15
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v141 := GenS19{F0: 42}
                                                                                                                        if len(v8) >= 2 {
                                                                                                                            v8 = v8[1:len(v8)]
                                                                                                                        }
                                                                                                                        v99 = (((v99 + v99) + (v99 + ((v99 + "foo bar") + v99))) + (v99 + ""))
                                                                                                                        if false {
                                                                                                                            if (false && false) {
                                                                                                                                if (uint16(v106) < v111) {
                                                                                                                                    var v142 int = 9223372036854775807
                                                                                                                                    _ = v142
                                                                                                                                } else {
                                                                                                                                    switch (v104 / (((fi15 - 42)) | 1)) {
                                                                                                                                    case -1:
                                                                                                                                        fmt.Println(fi15)
                                                                                                                                    default:
                                                                                                                                        v8 = append(v8, genFn0())
                                                                                                                                    }
                                                                                                                                    v143 := GenS2{F0: 2147483647, F1: 4294967295, F2: (^(v106))}
                                                                                                                                    _ = v143
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    var v144 int16 = 404
                                                                                                                                    v104 = (v104 >> ((18446744073709551615) & 63))
                                                                                                                                    _ = v144
                                                                                                                                }
                                                                                                                                v103 = ((3.14159 + -2.25) + v103)
                                                                                                                            } else {
                                                                                                                                v99 = ("\n\t" + v99)
                                                                                                                                switch fi15 {
                                                                                                                                case 0:
                                                                                                                                    var v145 float64 = 0.0
                                                                                                                                    _ = v145
                                                                                                                                case 796:
                                                                                                                                    var v146 float32 = ((1e10 / v103) - float32((0.0 - 1.5)))
                                                                                                                                    _ = v146
                                                                                                                                default:
                                                                                                                                    v147 := GenS20{F0: 0, F1: 4294967295, F2: ((!true) && false)}
                                                                                                                                    _ = v147
                                                                                                                                }
                                                                                                                                var v148 float32 = float32(3.14159)
                                                                                                                                _ = v148
                                                                                                                            }
                                                                                                                            v106 -= (v106 >> ((18446744073709551615) & 63))
                                                                                                                            v149 := GenS10{F0: fi93, F1: 795, F2: -407}
                                                                                                                            _ = v149
                                                                                                                        } else {
                                                                                                                            fmt.Println(p1, fi93)
                                                                                                                        }
                                                                                                                        v104 += genFn0()
                                                                                                                        _ = v141
                                                                                                                    }
                                                                                                                    fmt.Println(v106, v106, fi15)
                                                                                                                    v8 = append(v8, genFn0())
                                                                                                                    if len(v8) >= 2 {
                                                                                                                        v8 = v8[1:len(v8)]
                                                                                                                    }
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                v8 = append(v8, fi15)
                                                                                                            default:
                                                                                                                switch (-(v104)) {
                                                                                                                case -2147483648:
                                                                                                                    p0 = -1
                                                                                                                    v150 := GenS21{F0: 113}
                                                                                                                    if true {
                                                                                                                        if (2 < (v106 & v106)) {
                                                                                                                            var v151 uint16 = uint16(v106)
                                                                                                                            p0 = (p0 * p0)
                                                                                                                            if true {
                                                                                                                                if false {
                                                                                                                                    if len(v8) > 0 {
                                                                                                                                        v8[0] = v104
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                v99 = (("" + v99) + "hello")
                                                                                                                                v106 = (v106 - (v106 * (v106 * (v106 << ((2) & 63)))))
                                                                                                                                var v152 uint64 = (^(v106))
                                                                                                                                if true {
                                                                                                                                    if true {
                                                                                                                                        v153 := GenS4{F0: 12}
                                                                                                                                        _ = v153
                                                                                                                                    } else {
                                                                                                                                        var v154 uint32 = 2147483648
                                                                                                                                        _ = v154
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v152
                                                                                                                            }
                                                                                                                            v103 += (v103 - v103)
                                                                                                                            _ = v151
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v104 = -808
                                                                                                                    v155 := GenS3{F0: v99, F1: (v111 * p1)}
                                                                                                                    _ = v150
                                                                                                                    _ = v155
                                                                                                                case -916:
                                                                                                                    switch v99 {
                                                                                                                    case "":
                                                                                                                        if len(v8) > 0 {
                                                                                                                            v8[0] = (fi93 - (v104 * genFn0()))
                                                                                                                        }
                                                                                                                        if true {
                                                                                                                            v104 += genFn0()
                                                                                                                            v8 = append(v8, genFn0())
                                                                                                                            var v156 uint8 = 1
                                                                                                                            _ = v156
                                                                                                                        } else {
                                                                                                                            var v157 uint64 = v106
                                                                                                                            _ = v157
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        var v158 uint8 = 255
                                                                                                                        _ = v158
                                                                                                                    case "a":
                                                                                                                        v99 = v99
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if false {
                                                                                                                        var v159 int = fi93
                                                                                                                        v8 = append(v8, v104)
                                                                                                                        p0 -= (p0 - 168)
                                                                                                                        _ = v159
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if false {
                                                                                                                if (((v106 / ((v106) | 1)) > 704) || (862 == -251)) {
                                                                                                                    if false {
                                                                                                                        v99 = v99
                                                                                                                        var v160 uint = 65535
                                                                                                                        _ = v160
                                                                                                                    } else {
                                                                                                                        var v161 int64 = -737
                                                                                                                        _ = v161
                                                                                                                    }
                                                                                                                    fmt.Println(fi15, p0)
                                                                                                                    fmt.Println(v103)
                                                                                                                } else {
                                                                                                                    v162 := GenS22{F0: (p0 - p0), F1: (v99 + "a")}
                                                                                                                    v163 := GenS23{F0: 0}
                                                                                                                    v164 := GenS16{F0: 18446744073709551615, F1: (((p0 * p0) != p0) && (!(813 >= 312)))}
                                                                                                                    if false {
                                                                                                                        if ((!(-2147483648 >= 42)) || true) {
                                                                                                                            fmt.Println(v106, p0, p1)
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        _ = len(v8)
                                                                                                                    }
                                                                                                                    _ = v162
                                                                                                                    _ = v163
                                                                                                                    _ = v164
                                                                                                                }
                                                                                                                v111 = 255
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v165 uint8 = 2
                                                                                                            v111 += v111
                                                                                                            v165 -= (^(v165))
                                                                                                            _ = v165
                                                                                                        }
                                                                                                    }
                                                                                                    var v166 uint32 = 808
                                                                                                    _ = len(v8)
                                                                                                    _ = v166
                                                                                                }
                                                                                                _ = v111
                                                                                                _ = v112
                                                                                            }
                                                                                            v103 += v103
                                                                                            var v167 int = -1
                                                                                            _ = v167
                                                                                        }
                                                                                        switch v104 {
                                                                                        case 128:
                                                                                            if len(v8) >= 2 {
                                                                                                v8 = v8[1:len(v8)]
                                                                                            }
                                                                                        case 127:
                                                                                            var v168 uint8 = 255
                                                                                            if (v104 >= 32768) {
                                                                                                var v169 int = (-(fi93))
                                                                                                _ = v169
                                                                                            }
                                                                                            v170 := GenS23{F0: 1}
                                                                                            var v171 float64 = (-2.25 - 0.0)
                                                                                            v103 = (1e10 + (v103 * float32(v171)))
                                                                                            _ = v168
                                                                                            _ = v170
                                                                                            _ = v171
                                                                                        case -9223372036854775808:
                                                                                            var v172 uint64 = v106
                                                                                            _ = v172
                                                                                        }
                                                                                        v173 := GenS9{F0: v106, F1: 183}
                                                                                        if (!false) {
                                                                                            switch "" {
                                                                                            case "a":
                                                                                                var v174 uint16 = p1
                                                                                                _ = v174
                                                                                            case "foo bar":
                                                                                                v106 += (v106 ^ (v106 & 2))
                                                                                            case "\n\t":
                                                                                                p1 += 132
                                                                                            default:
                                                                                                v104 += (fi15 + (^(fi15)))
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(p0, v103, p1)
                                                                                            var v175 int32 = 0
                                                                                            _ = v175
                                                                                        }
                                                                                        var v176 uint = 77
                                                                                        _ = v173
                                                                                        _ = v176
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v177 uint16 = (p1 * (p1 / ((29) | 1)))
                                                                                    _ = v177
                                                                                case "hello":
                                                                                    if (0 != 252) {
                                                                                        var v178 uint8 = 128
                                                                                        _ = v178
                                                                                    } else {
                                                                                        if false {
                                                                                            var v179 uint64 = 0
                                                                                            switch (v99 + v99) {
                                                                                            case "foo bar":
                                                                                                _ = len(v8)
                                                                                            }
                                                                                            _ = v179
                                                                                        }
                                                                                        v180 := GenS24{F0: 286, F1: 255, F2: -128}
                                                                                        v103 -= (1.5 * v103)
                                                                                        _ = v180
                                                                                    }
                                                                                default:
                                                                                    p0 += p0
                                                                                }
                                                                                var v181 uint = 341
                                                                                _ = v107
                                                                                _ = v181
                                                                            }
                                                                        }
                                                                        p0 += -851
                                                                        if len(v8) > 0 {
                                                                            v8[0] = (fi93 / ((fi15) | 1))
                                                                        }
                                                                        v106 = 2
                                                                    }
                                                                    if len(v8) > 0 {
                                                                        v8[0] = (v104 - (fi93 % ((fi15) | 1)))
                                                                    }
                                                                    v104 = -70
                                                                    var v182 int32 = -2147483648
                                                                    _ = v182
                                                                }
                                                            case "hello":
                                                                v183 := GenS1{F0: fi15, F1: 9223372036854775807}
                                                                _ = v183
                                                            }
                                                            _ = v103
                                                            _ = v104
                                                            _ = v106
                                                        }
                                                        var v184 float64 = (0.1 * ((-2.25 / float64((float32(-2.25) + (((0.0 - (0.0 * float32(float64(1e10)))) - 3.14159) * float32(float64(float32(float64(-2.25)))))))) - 3.14159))
                                                        var v185 int8 = -128
                                                        _ = len(v8)
                                                        _ = v184
                                                        _ = v185
                                                    }
                                                default:
                                                    fmt.Println(p0, p1, p1)
                                                }
                                                fmt.Println(p0)
                                            } else {
                                                var v186 uint16 = 42
                                                v187 := GenS25{F0: 0}
                                                _ = v186
                                                _ = v187
                                            }
                                        }
                                        v188 := GenS26{F0: int8(p0)}
                                        _ = v188
                                        _ = v99
                                    }
                                case 127:
                                    v189 := GenS27{F0: 477, F1: 909}
                                    _ = v189
                                case 42:
                                    fmt.Println(fi15, p0)
                                default:
                                    if (0.1 < (0.0 - 0.0)) {
                                        v190 := GenS2{F0: 2147483647, F1: 389, F2: 18446744073709551615}
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        p1 = (p1 + p1)
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        _ = v190
                                    } else {
                                        switch ("a" + "a") {
                                        case "hello":
                                            p0 += (p0 % ((p0) | 1))
                                        default:
                                            v191 := GenS10{F0: fi93, F1: 390, F2: 9223372036854775807}
                                            _ = v191
                                        }
                                        fmt.Println(fi93)
                                        if false {
                                            var v192 float32 = (3.14159 + 3.14159)
                                            _ = v192
                                        } else {
                                            v8 = append(v8, (-(fi15)))
                                        }
                                    }
                                }
                            } else {
                                if true {
                                    v8 = append(v8, fi15)
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                } else {
                                    fmt.Println(fi15)
                                    var v193 int8 = 100
                                    _ = v193
                                }
                                switch (fi93 ^ (-(fi93))) {
                                case 127:
                                    var v194 string = "foo bar"
                                    _ = v194
                                case 886:
                                    if false {
                                        p1 = 258
                                    } else {
                                        if true {
                                            p0 = (p0 ^ p0)
                                        } else {
                                            if true {
                                                p0 += -32768
                                                v195 := GenS28{F0: false, F1: false}
                                                _ = v195
                                            } else {
                                                if len(v8) >= 2 {
                                                    v8 = v8[1:len(v8)]
                                                }
                                            }
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            v196 := GenS2{F0: 2147483647, F1: 172, F2: 0}
                                            _ = v196
                                        }
                                        var v197 string = "\n\t"
                                        _ = v197
                                    }
                                }
                                fmt.Println(p0, p1, fi15)
                            }
                            var v198 float32 = (float32(((1.5 - 0.0) + 3.14159)) - ((1e10 + (1.5 - 3.14159)) / ((3.14159 / (-2.25 + 0.0)) * 1e10)))
                            v199 := GenS29{F0: -882, F1: 66}
                            _ = fi93
                            _ = v198
                            _ = v199
                        }
                    }
                    var v200 float32 = -2.25
                    _ = v200
                }
            }
            for fi201 := 0; fi201 < 10; fi201++ {
                p1 = 32767
                var v202 int = genFn0()
                if len(v8) > 0 {
                    v8[0] = 42
                }
                _ = fi201
                _ = v202
            }
            switch 128 {
            case -642:
                p1 += (p1 - (p1 * (p1 >> ((18446744073709551615) & 63))))
            case 4294967295:
                for fi203 := 0; fi203 < 8; fi203++ {
                    var v204 uint64 = 42
                    _ = fi203
                    _ = v204
                }
            default:
                for fi205 := 0; fi205 < 11; fi205++ {
                    fmt.Println(fi205, fi15, fi205)
                    v206 := GenS25{F0: 1}
                    fmt.Println(fi205, fi205, fi205)
                    _ = fi205
                    _ = v206
                }
            }
            _ = fi15
        }
        var v207 int = 32768
        _ = v14
        _ = v207
    }
    _ = p0
    _ = p1
    _ = v8
    return 42
}

func main() {
    var v1 uint16 = 2
    var v2 string = "foo bar"
    var v3 int32 = 2147483647
    var v4 uint32 = 0
    for fi208 := 0; fi208 < 5; fi208++ {
        var v209 string = v2
        _ = fi208
        _ = v209
    }
    fmt.Println(v3, v4, v1)
    for fi210 := 0; fi210 < 11; fi210++ {
        v3 = (v3 - v3)
        _ = fi210
    }
    fmt.Println(v4, v2)
    var v211 uint64 = 437
    if ((v3 + (-(v3))) <= v3) {
        fmt.Println(v2, v1, v211)
        if (!(false && ((true || (!false)) && (-39 == -115)))) {
            fmt.Println(v1, v4)
            v3 += (v3 + 42)
        } else {
            switch (v2 + "hello") {
            case "a":
                var v212 float32 = (0.1 / (0.0 + 0.1))
                _ = v212
            case "":
                fmt.Println(v2, v3)
            default:
                for fi213 := 0; fi213 < 3; fi213++ {
                    v214 := []int{}
                    switch (-(fi213)) {
                    case 128:
                        switch 128 {
                        case 1:
                            if (((!false) || ((v4 >= uint32(v1)) && true)) && (0.1 != 3.14159)) {
                                v215 := GenS30{F0: -1}
                                _ = v215
                            } else {
                                var v216 uint32 = (v4 ^ v4)
                                fmt.Println(v211, v211, v1)
                                _ = v216
                            }
                        default:
                            switch genFn0() {
                            case -666:
                                switch (v2 + v2) {
                                case "a":
                                    for fi217 := 0; fi217 < 13; fi217++ {
                                        switch (v2 + v2) {
                                        case "":
                                            var v218 int16 = 42
                                            _ = v218
                                        }
                                        if (-102 <= genFn1(1, (v1 ^ (v1 * v1)))) {
                                            if len(v214) >= 2 {
                                                v214 = v214[1:len(v214)]
                                            }
                                            var v219 uint64 = v211
                                            _ = v219
                                        }
                                        var v220 uint8 = 127
                                        _ = fi217
                                        _ = v220
                                    }
                                case "":
                                    var v221 uint8 = 1
                                    _ = v221
                                default:
                                    v222 := GenS6{F0: -650, F1: 89, F2: genFn0()}
                                    _ = v222
                                }
                            default:
                                v1 += (^(v1))
                            }
                        }
                    default:
                        for fi223 := 0; fi223 < 13; fi223++ {
                            v4 -= 2147483647
                            _ = fi223
                        }
                    }
                    switch "a" {
                    case "hello":
                        var v224 string = v2
                        _ = v224
                    default:
                        var v225 uint = 4294967295
                        _ = v225
                    }
                    _ = fi213
                    _ = v214
                }
            }
            var v226 float32 = 3.14159
            for fi227 := 0; fi227 < 8; fi227++ {
                fmt.Println(fi227)
                fmt.Println(v2)
                _ = fi227
            }
            _ = v226
        }
        v1 = uint16(v211)
    } else {
        v2 = v2
        v3 -= (-(v3))
    }
    v1 -= (^(v1))
    v2 = ""
    for fi228 := 0; fi228 < 15; fi228++ {
        var v229 uint8 = 0
        _ = fi228
        _ = v229
    }
    v230 := GenS31{F0: (v3 << ((42) & 63)), F1: v3, F2: true}
    if (100 >= 42) {
        v231 := []int{}
        v211 -= (v211 - 926)
        if false {
            v3 = v3
            var v232 uint = 42
            var v233 uint32 = (v4 / ((196) | 1))
            _ = v232
            _ = v233
        } else {
            var v234 uint = 0
            for fi235 := 0; fi235 < 12; fi235++ {
                fmt.Println(v2)
                if (!false) {
                    for fi236 := 0; fi236 < 6; fi236++ {
                        var v237 string = "\n\t"
                        v234 = v234
                        _ = fi236
                        _ = v237
                    }
                    for fi238 := 0; fi238 < 3; fi238++ {
                        v2 = v2
                        _ = fi238
                    }
                } else {
                    v234 -= 350
                    if true {
                        for fi239 := 0; fi239 < 15; fi239++ {
                            var v240 string = "a"
                            var v241 int64 = int64(v3)
                            var v242 int64 = int64(v1)
                            _ = fi239
                            _ = v240
                            _ = v241
                            _ = v242
                        }
                    }
                    var v243 uint32 = v4
                    _ = v243
                }
                v244 := GenS29{F0: int64(v4), F1: (v4 & (v4 << ((v234) & 63)))}
                _ = fi235
                _ = v244
            }
            fmt.Println(v234)
            _ = v234
        }
        _ = v231
    }
    var v245 int = 32767
    for fi246 := 0; fi246 < 5; fi246++ {
        var v247 int16 = 37
        fmt.Println(v2, fi246)
        switch (fi246 ^ v245) {
        case 4294967295:
            v247 += (-(v247))
        default:
            v211 -= 18446744073709551615
        }
        _ = fi246
        _ = v247
    }
    switch (v245 | (v245 << ((396) & 63))) {
    case 0:
        var v248 uint32 = (v4 - (^(v4)))
        _ = v248
    default:
        v245 = 32768
    }
    if false {
        v211 -= 2
        v249 := GenS32{F0: 499, F1: v245, F2: (v245 + 811)}
        _ = v249
    } else {
        if (1 < (^(v4))) {
            fmt.Println(v3, v4, v2)
        } else {
            v250 := []int{v245, v245, 176, -757}
            v250 = append(v250, v245)
            _ = v250
        }
        v251 := []int{genFn0(), v245}
        var v252 float64 = float64(0.0)
        _ = v251
        _ = v252
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v211:", v211)
    fmt.Println("v245:", v245)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v211
    _ = v230
    _ = v245
    _ = v3
    _ = v4
}
