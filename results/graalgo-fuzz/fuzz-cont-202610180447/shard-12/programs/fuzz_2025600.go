package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint32; F1 uint8; F2 bool }
type GenS7 struct { F0 uint64; F1 int; F2 uint32 }
type GenS8 struct { F0 uint64; F1 uint64; F2 int }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 int8; F1 uint32 }
type GenS11 struct { F0 int16; F1 int32 }
type GenS12 struct { F0 int64; F1 string }
type GenS13 struct { F0 bool }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint = 211
    v3 := []uint32{}
    v4 := make(map[uint32]int)
    var v5 uint8 = 2
    switch ("foo bar" + "") {
    case "foo bar":
        if (!(-1 >= -2147483648)) {
            var v6 uint = (v2 ^ 4294967295)
            _ = v6
        }
    case "hello":
        if false {
            v1 -= 1
        } else {
            delete(v4, v1)
            for fi7 := 0; fi7 < 17; fi7++ {
                v1 += (v1 ^ uint32(v2))
                fmt.Println(v2)
                var v8 uint8 = v5
                _ = fi7
                _ = v8
            }
        }
    case "a":
        for fi9 := 0; fi9 < 20; fi9++ {
            v4[65535] = fi9
            _ = fi9
        }
    default:
        v10 := GenS1{F0: -39, F1: -225}
        _ = v10
    }
    var v11 int = 42
    switch "\n\t" {
    case "\n\t":
        v3 = append(v3, v1)
    case "hello":
        if false {
            if (42 >= 1) {
                var v12 int = v11
                v5 = v5
                _ = v12
            }
        }
    case "":
        switch -735 {
        case -2147483648:
            for fi13 := 0; fi13 < 9; fi13++ {
                if true {
                    v11 += int(v2)
                }
                _ = fi13
            }
        case 42:
            if (((1.5 + -2.25) + (0.0 - 3.14159)) <= 1.5) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    v3 = append(v3, v1)
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v3 = append(v3, v1)
                        v5 += 0
                        _ = fi15
                    }
                    _ = fi14
                }
                var v16 int32 = 508
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 float32 = (0.1 - (1e10 / (0.1 + 0.0)))
                for fi19 := 0; fi19 < 8; fi19++ {
                    if (("hello" + "") < "") {
                        v4[v1] = 32767
                        v20 := GenS1{F0: 0, F1: -32768}
                        _ = v20
                    }
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
        }
    default:
        for fi21 := 0; fi21 < 6; fi21++ {
            switch ("a" + "a") {
            case "hello":
                var v22 uint32 = v1
                _ = v22
            default:
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(v1)
                    v24 := GenS1{F0: 0, F1: -32768}
                    v25 := GenS2{F0: (^(v5))}
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            var v26 string = "foo bar"
            _ = fi21
            _ = v26
        }
    }
    v2 += (v2 | 4294967295)
    var v27 uint64 = 18446744073709551615
    for fi28 := 0; fi28 < 8; fi28++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi29 := 0; fi29 < 18; fi29++ {
            v11 += v11
            if false {
                var v30 float32 = 1.5
                if (fi29 != fi29) {
                    var v31 float32 = v30
                    _ = v31
                } else {
                    if (917 >= uint16(v1)) {
                        v30 += float32(-2.25)
                        var v32 int16 = 1
                        _ = v32
                    } else {
                        if (v1 <= v1) {
                            fmt.Println(v11, v2, fi29)
                            v33 := GenS1{F0: -118, F1: 32767}
                            var v34 uint64 = 18446744073709551615
                            _ = v33
                            _ = v34
                        } else {
                            if len(v3) > 0 {
                                v3[0] = v1
                            }
                        }
                        delete(v4, (v1 + 187))
                        if ((v2 ^ uint(v5)) <= 71) {
                            v35 := GenS1{F0: 100, F1: 42}
                            v5 += v5
                            if true {
                                switch "foo bar" {
                                case "":
                                    switch fi28 {
                                    case 0:
                                        v30 = v30
                                    case 127:
                                        switch "a" {
                                        case "":
                                            if (("hello" + (("hello" + "a") + "a")) != "\n\t") {
                                                switch v11 {
                                                case 4294967295:
                                                    v5 -= (v5 & v5)
                                                case 127:
                                                    var v36 int8 = 127
                                                    _ = v36
                                                case -9223372036854775808:
                                                    switch v11 {
                                                    case 0:
                                                        v2 = v2
                                                    case 32768:
                                                        v2 -= v2
                                                    }
                                                default:
                                                    v37 := v4[(v1 * (v1 % ((331) | 1)))]
                                                    _ = v37
                                                }
                                                _ = len(v3)
                                            }
                                        case "a":
                                            var v38 int16 = -32768
                                            _ = v38
                                        case "\n\t":
                                            var v39 float64 = (3.14159 - 1e10)
                                            _ = v39
                                        }
                                    case -879:
                                        if false {
                                            var v40 bool = false
                                            _ = v40
                                        }
                                    default:
                                        if true {
                                            v2 -= 18446744073709551615
                                            var v41 int16 = -32768
                                            _ = v41
                                        }
                                    }
                                case "hello":
                                    switch v11 {
                                    case 127:
                                        var v42 float32 = ((v30 - (1.5 * v30)) / (((v30 * v30) - v30) - v30))
                                        _ = v42
                                    case 1:
                                        if true {
                                            fmt.Println(v5, v2)
                                        }
                                    case 84:
                                        if (v11 < (fi29 | (fi28 % (((-(fi28))) | 1)))) {
                                            if (false || (1.5 != 1e10)) {
                                                var v43 uint16 = 42
                                                v44 := GenS2{F0: v5}
                                                _ = v43
                                                _ = v44
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = (v1 % ((42) | 1))
                                            }
                                        } else {
                                            var v45 int8 = 0
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v46 uint = v2
                                            _ = v45
                                            _ = v46
                                        }
                                    default:
                                        v5 -= (v5 ^ v5)
                                    }
                                }
                                var v47 int32 = 2147483647
                                if ((int32(v5) < v47) || false) {
                                    v47 -= 1
                                    fmt.Println(fi28, v47)
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    var v48 int = (fi28 >> ((4294967295) & 63))
                                    fmt.Println(v11, v48, v5)
                                    _ = v48
                                }
                                _ = v47
                            } else {
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        switch (v11 + 2147483647) {
                                        case 128:
                                            var v49 int32 = 2147483647
                                            _ = v49
                                        }
                                    case "hello":
                                        switch v11 {
                                        case 2147483647:
                                            v50 := v4[2]
                                            _ = v50
                                        case -416:
                                            var v51 float32 = 1.5
                                            _ = v51
                                        }
                                    case "\n\t":
                                        v27 += 111
                                    default:
                                        switch "hello" {
                                        case "hello":
                                            v1 += v1
                                        case "":
                                            if true {
                                                v52 := v4[991]
                                                v53 := GenS3{F0: "a"}
                                                switch "a" {
                                                case "hello":
                                                    v11 = v11
                                                default:
                                                    v3 = append(v3, (^(v1)))
                                                }
                                                _ = len(v4)
                                                _ = v52
                                                _ = v53
                                            } else {
                                                v54 := GenS4{F0: (fi29 ^ (v11 + fi28)), F1: (42 < 1)}
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if (1e10 >= v30) {
                                                        var v55 float32 = v30
                                                        _ = v55
                                                    } else {
                                                        v27 = (^(v27))
                                                        switch (v11 << (((v2 * uint(v5))) & 63)) {
                                                        case 514:
                                                            v56 := GenS5{F0: v1, F1: ""}
                                                            _ = v56
                                                        case 128:
                                                            fmt.Println(v1, v5, v2)
                                                        }
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 4294967295
                                                    }
                                                }
                                                fmt.Println(fi28, v27, v27)
                                                _ = v54
                                            }
                                        default:
                                            v1 = (v1 | (v1 - v1))
                                        }
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v11 -= 9223372036854775807
                            }
                            switch "" {
                            case "a":
                                var v57 string = ("hello" + "\n\t")
                                _ = v57
                            case "foo bar":
                                var v58 bool = true
                                _ = v58
                            default:
                                _ = len(v3)
                            }
                            _ = v35
                        }
                    }
                    v2 = v2
                    v11 -= (fi29 | (fi29 | 128))
                }
                _ = v30
            } else {
                delete(v4, (v1 - 2147483648))
                v59 := GenS6{F0: (v1 << (((^(v2))) & 63)), F1: 1, F2: (696 != 9)}
                v5 = 42
                _ = v59
            }
            _ = fi29
        }
        _ = fi28
    }
    v11 += (v11 * v11)
    if len(v3) > 0 {
        v3[0] = (v1 ^ v1)
    }
    var v60 uint64 = v27
    switch "foo bar" {
    case "":
        var v61 bool = false
        _ = v61
    case "hello":
        if false {
            v4[v1] = v11
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v62 float64 = 0.1
            var v63 int64 = 34
            _ = v62
            _ = v63
        }
    default:
        var v64 int16 = -32768
        _ = v64
    }
    if false {
        v65 := GenS6{F0: (v1 / (((v1 * 197)) | 1)), F1: (^(v5)), F2: (-128 <= 24)}
        for fi66 := 0; fi66 < 19; fi66++ {
            switch "a" {
            case "hello":
                v5 -= v5
            default:
                v67 := GenS5{F0: uint32(v60), F1: "\n\t"}
                _ = v67
            }
            _ = fi66
        }
        switch ("hello" + "a") {
        case "foo bar":
            v68 := GenS7{F0: 0, F1: v11, F2: v1}
            _ = v68
        case "":
            for fi69 := 0; fi69 < 7; fi69++ {
                fmt.Println(v1)
                for fi70 := 0; fi70 < 10; fi70++ {
                    _ = len(v4)
                    var v71 bool = (-128 >= -98)
                    v4[(v1 - (v1 << (((v2 / ((v2) | 1))) & 63)))] = fi70
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    v72 := GenS8{F0: v27, F1: 2, F2: fi70}
                    v5 -= uint8(v2)
                    _ = fi70
                    _ = v71
                    _ = v72
                }
                _ = fi69
            }
        case "a":
            if (true || true) {
                var v73 bool = (3.14159 < 1.5)
                _ = v73
            } else {
                switch -837 {
                case 2147483647:
                    for fi74 := 0; fi74 < 4; fi74++ {
                        if false {
                            for fi75 := 0; fi75 < 14; fi75++ {
                                switch "hello" {
                                case "foo bar":
                                    if (!true) {
                                        v1 = v1
                                        v5 = (v5 & (v5 | uint8(v27)))
                                        v76 := GenS9{F0: (v27 >> ((v2) & 63)), F1: v2}
                                        _ = v76
                                    }
                                }
                                var v77 float64 = float64((float32(float64((0.1 / 1.0))) * 0.1))
                                v78 := GenS4{F0: -1, F1: (!false)}
                                _ = fi75
                                _ = v77
                                _ = v78
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            switch (fi74 ^ (v11 ^ (^(fi74)))) {
                            case -854:
                                fmt.Println(v11)
                            case 32767:
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    for fi79 := 0; fi79 < 15; fi79++ {
                                        if (v2 == (v2 >> ((v2) & 63))) {
                                            _ = len(v4)
                                        } else {
                                            v27 += v60
                                            v4[(v1 | v1)] = fi79
                                        }
                                        _ = len(v4)
                                        if ((^(v27)) >= v27) {
                                            v80 := GenS9{F0: v60, F1: v2}
                                            v81 := GenS4{F0: (fi74 | 657), F1: true}
                                            fmt.Println(v11, v60)
                                            _ = v80
                                            _ = v81
                                        } else {
                                            v82 := GenS6{F0: 2147483648, F1: (v5 - (^(v5))), F2: false}
                                            if (!(true || true)) {
                                                var v83 int8 = -105
                                                _ = v83
                                            } else {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if (!true) {
                                                    if (("\n\t" + "hello") == "a") {
                                                        v2 = v2
                                                        fmt.Println(v5, v11)
                                                        v60 = 1
                                                        fmt.Println(v5, fi79)
                                                        delete(v4, v1)
                                                    }
                                                    var v84 uint8 = 15
                                                    fmt.Println(fi79)
                                                    _ = v84
                                                }
                                                switch (v11 & int(v60)) {
                                                case 4294967295:
                                                    _ = len(v4)
                                                }
                                            }
                                            _ = v82
                                        }
                                        _ = fi79
                                    }
                                case "a":
                                    switch (fi74 & 42) {
                                    case -383:
                                        v85 := GenS10{F0: 107, F1: (v1 ^ (v1 & (v1 & v1)))}
                                        _ = v85
                                    default:
                                        switch fi74 {
                                        case -777:
                                            v4[uint32(v2)] = (^(fi74))
                                        case -853:
                                            if (!(true && true)) {
                                                var v86 int8 = 42
                                                _ = v86
                                            } else {
                                                v87 := GenS7{F0: (v27 << ((uint(v5)) & 63)), F1: (fi74 * fi74), F2: (v1 + (v1 & 44))}
                                                if false {
                                                    var v88 bool = false
                                                    var v89 uint16 = 916
                                                    for fi90 := 0; fi90 < 5; fi90++ {
                                                        v5 = (v5 & v5)
                                                        switch "" {
                                                        case "\n\t":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case "hello":
                                                            if false {
                                                                var v91 uint32 = 2147483647
                                                                var v92 int = v11
                                                                var v93 uint64 = v60
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                            } else {
                                                                switch 127 {
                                                                case 825:
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v94 := GenS11{F0: 42, F1: 2147483647}
                                                                        _ = v94
                                                                    }
                                                                case 378:
                                                                    if ((-2.25 / ((float64((-2.25 * (0.0 / ((1e10 + 3.14159) - 0.1)))) - 3.14159) * 1e10)) < 1e10) {
                                                                        v95 := GenS12{F0: -9223372036854775808, F1: ("" + "a")}
                                                                        var v96 int16 = 1
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        var v97 uint64 = (v60 ^ v27)
                                                                        _ = v97
                                                                    }
                                                                case -429:
                                                                    if v88 {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    }
                                                                default:
                                                                    var v98 int = (v11 & (fi90 ^ fi74))
                                                                    _ = v98
                                                                }
                                                                if v88 {
                                                                    switch (fi90 ^ 42) {
                                                                    case -198:
                                                                        v88 = ("" >= "\n\t")
                                                                    case 42:
                                                                        v60 = (^(v60))
                                                                    }
                                                                    var v99 uint = 4294967295
                                                                    _ = v99
                                                                } else {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = v1
                                                                    }
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2147483647
                                                                    }
                                                                    v60 = v60
                                                                }
                                                            }
                                                        case "":
                                                            fmt.Println(v2, v27, v5)
                                                        }
                                                        _ = fi90
                                                    }
                                                    _ = v88
                                                    _ = v89
                                                } else {
                                                    var v100 int64 = -720
                                                    _ = v100
                                                }
                                                v2 = v2
                                                _ = v87
                                            }
                                        case 128:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    }
                                }
                            case -578:
                                switch (fi74 << ((uint(v60)) & 63)) {
                                case 0:
                                    v101 := GenS1{F0: -128, F1: 32767}
                                    _ = v101
                                case 184:
                                    var v102 int = (v11 + 2147483647)
                                    _ = v102
                                }
                            }
                        default:
                            v27 += v27
                        }
                        if ((1e10 + 1e10) <= 1.5) {
                            v103 := GenS13{F0: false}
                            _ = v103
                        }
                        _ = fi74
                    }
                case 42:
                    var v104 int8 = 42
                    _ = v104
                case 0:
                    v105 := GenS2{F0: (v5 >> ((v2) & 63))}
                    _ = v105
                }
                delete(v4, v1)
            }
        default:
            v106 := GenS12{F0: 9223372036854775807, F1: "a"}
            _ = v106
        }
        _ = v65
    } else {
        v107 := GenS1{F0: -49, F1: -32768}
        _ = v107
    }
    var v108 bool = true
    fmt.Println(v27, v60)
    var v109 int64 = 9223372036854775807
    fmt.Println(v108)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v108
    _ = v109
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v60
}
