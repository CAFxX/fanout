package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 int16; F1 bool; F2 int64 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 int; F2 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 uint32; F2 int }
type GenS10 struct { F0 int8 }

func main() {
    var v1 int32 = -163
    var v2 int = 0
    var v3 bool = (!(false || (!(false && false))))
    v4 := []uint32{2, 327, 399}
    if true {
        switch 128 {
        case 128:
            var v5 bool = (v1 > v1)
            _ = v5
        case 233:
            var v6 bool = (v3 && v3)
            _ = v6
        default:
            v7 := GenS1{F0: "a"}
            _ = v7
        }
        fmt.Println(v1, v1)
    }
    v2 = (v2 - v2)
    switch (v2 - 32767) {
    case -1:
        if (!v3) {
            if v3 {
                for fi8 := 0; fi8 < 13; fi8++ {
                    fmt.Println(v1)
                    _ = fi8
                }
                if v3 {
                    for fi9 := 0; fi9 < 10; fi9++ {
                        fmt.Println(fi9, fi9)
                        _ = fi9
                    }
                } else {
                    var v10 uint8 = 255
                    _ = v10
                }
                v1 -= v1
            }
            v11 := GenS1{F0: ("\n\t" + "\n\t")}
            _ = v11
        } else {
            var v12 uint8 = 203
            _ = v12
        }
    }
    v3 = (!v3)
    for fi13 := 0; fi13 < 5; fi13++ {
        if (((true || v3) || v3) && v3) {
            if (false && (1e10 >= 0.1)) {
                var v14 uint16 = 1
                for fi15 := 0; fi15 < 2; fi15++ {
                    switch (v2 << ((26) & 63)) {
                    case -2147483648:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 175:
                        v14 = 65535
                    default:
                        if (v14 <= 255) {
                            v1 = 870
                            if true {
                                v16 := GenS1{F0: ("a" + ("" + ("\n\t" + (("hello" + "hello") + "hello"))))}
                                var v17 float64 = 3.14159
                                v4 = append(v4, 403)
                                _ = v16
                                _ = v17
                            } else {
                                if v3 {
                                    switch "foo bar" {
                                    case "hello":
                                        if (v3 && v3) {
                                            if v3 {
                                                _ = len(v4)
                                                if len(v4) > 0 {
                                                    v4[0] = 42
                                                }
                                                switch fi15 {
                                                case 6:
                                                    v18 := GenS2{F0: (fi15 << ((512) & 63)), F1: 1}
                                                    _ = v18
                                                case -2147483648:
                                                    var v19 int8 = 100
                                                    _ = v19
                                                case 484:
                                                    v2 -= (fi15 * (fi15 + -233))
                                                }
                                            } else {
                                                v4 = append(v4, 812)
                                                if false {
                                                    v14 += (^(v14))
                                                    v1 = (v1 + v1)
                                                } else {
                                                    switch (fi13 << ((106) & 63)) {
                                                    case 20:
                                                        var v20 int16 = -343
                                                        _ = v20
                                                    case -1:
                                                        v2 -= (^(fi15))
                                                    default:
                                                        var v21 int64 = -9223372036854775808
                                                        _ = v21
                                                    }
                                                }
                                            }
                                        }
                                    default:
                                        var v22 uint = 42
                                        _ = v22
                                    }
                                    if (!v3) {
                                        v3 = true
                                    } else {
                                        v3 = v3
                                        fmt.Println(v14, v14, fi13)
                                    }
                                    v14 = 1
                                }
                            }
                        }
                    }
                    _ = fi15
                }
                if (v3 && (2 == 1)) {
                    switch fi13 {
                    case 32768:
                        v2 -= (fi13 * (v2 ^ (fi13 + v2)))
                    case -12:
                        v1 = (^(v1))
                    case 9223372036854775807:
                        v23 := GenS1{F0: ("hello" + "")}
                        _ = v23
                    }
                    switch "a" {
                    case "a":
                        for fi24 := 0; fi24 < 4; fi24++ {
                            fmt.Println(v1, v1)
                            _ = fi24
                        }
                    case "":
                        for fi25 := 0; fi25 < 5; fi25++ {
                            if v3 {
                                fmt.Println(v14, fi13, v14)
                                if ((true || (v3 || v3)) && (1 != 127)) {
                                    v26 := GenS2{F0: fi25, F1: 1}
                                    _ = v26
                                }
                                v27 := GenS3{F0: -3, F1: ((0 != -32768) || v3), F2: 42}
                                _ = v27
                            } else {
                                v4 = append(v4, 937)
                                fmt.Println(v2)
                                var v28 uint64 = 42
                                _ = v28
                            }
                            var v29 float64 = 1e10
                            v4 = append(v4, 37)
                            v14 = v14
                            _ = len(v4)
                            _ = fi25
                            _ = v29
                        }
                    default:
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v4 = append(v4, 4294967295)
                            _ = fi30
                        }
                    }
                }
                _ = v14
            }
        } else {
            fmt.Println(v2, v3, v2)
            fmt.Println(v1)
            var v31 float32 = 0.0
            _ = v31
        }
        var v32 uint8 = 41
        _ = fi13
        _ = v32
    }
    var v33 int8 = -128
    for fi34 := 0; fi34 < 18; fi34++ {
        v4 = append(v4, 2147483647)
        _ = fi34
    }
    if v3 {
        v1 += v1
        if ((v2 + (v2 | (v2 + (v2 + (^(v2)))))) != 733) {
            v4 = append(v4, 2)
            if (430 == 818) {
                if ("foo bar" == "") {
                    if true {
                        var v35 int8 = -1
                        switch v2 {
                        case 42:
                            fmt.Println(v3, v3, v35)
                        case 1:
                            v35 = v33
                        case 128:
                            var v36 bool = true
                            _ = v36
                        }
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                        _ = v35
                    } else {
                        v37 := GenS4{F0: 255, F1: 32767}
                        _ = v37
                    }
                    if (2 <= 0) {
                        var v38 int32 = v1
                        var v39 uint = 2
                        _ = v38
                        _ = v39
                    }
                    v2 -= (v2 + -644)
                }
                var v40 float32 = -2.25
                if (0 != 1) {
                    if v3 {
                        for fi41 := 0; fi41 < 11; fi41++ {
                            var v42 int32 = (v1 % (((v1 ^ v1)) | 1))
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi41
                            _ = v42
                        }
                        v3 = v3
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    v2 = v2
                } else {
                    if v3 {
                        if (0 >= 42) {
                            for fi43 := 0; fi43 < 10; fi43++ {
                                switch (fi43 << ((2) & 63)) {
                                case 128:
                                    v44 := GenS5{F0: (v33 / ((v33) | 1))}
                                    _ = v44
                                case 94:
                                    var v45 bool = (-32768 >= 32767)
                                    _ = v45
                                }
                                _ = fi43
                            }
                            var v46 uint8 = 1
                            var v47 uint16 = 0
                            _ = v46
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 9; fi48++ {
                            switch "" {
                            case "":
                                switch 9223372036854775807 {
                                case -762:
                                    if v3 {
                                        if v3 {
                                            for fi49 := 0; fi49 < 19; fi49++ {
                                                v50 := GenS6{F0: 1, F1: (fi48 | 559), F2: 155}
                                                _ = fi49
                                                _ = v50
                                            }
                                        }
                                    }
                                case 1:
                                    v1 += 369
                                default:
                                    v51 := GenS3{F0: 0, F1: v3, F2: 642}
                                    _ = v51
                                }
                            case "\n\t":
                                for fi52 := 0; fi52 < 6; fi52++ {
                                    v3 = v3
                                    if (222 == 255) {
                                        switch "a" {
                                        case "foo bar":
                                            v2 += fi48
                                        case "hello":
                                            if true {
                                                switch "\n\t" {
                                                case "hello":
                                                    v33 = -113
                                                case "\n\t":
                                                    var v53 int64 = -9223372036854775808
                                                    _ = v53
                                                case "a":
                                                    v3 = v3
                                                default:
                                                    v4 = append(v4, 42)
                                                }
                                                if v3 {
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v54 bool = (469 >= 255)
                                                        _ = v54
                                                    }
                                                    if v3 {
                                                        var v55 int64 = -753
                                                        if len(v4) > 0 {
                                                            v4[0] = 4294967295
                                                        }
                                                        _ = v55
                                                    }
                                                }
                                            }
                                        default:
                                            switch -423 {
                                            case 32768:
                                                switch "" {
                                                case "a":
                                                    if (!(1 < 2)) {
                                                        fmt.Println(v2, fi52)
                                                    } else {
                                                        if ((!(true || ((true || (255 >= 1)) && (true || (v33 > (v33 * 0)))))) && v3) {
                                                            var v56 uint32 = 0
                                                            var v57 bool = v3
                                                            v4 = append(v4, v56)
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                case "\n\t":
                                                    v1 = v1
                                                default:
                                                    if true {
                                                        v40 += 0.0
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            switch v2 {
                                                            case 42:
                                                                if v3 {
                                                                    fmt.Println(v33, v33, fi48)
                                                                }
                                                            case 226:
                                                                var v58 int32 = (v1 << ((1) & 63))
                                                                _ = v58
                                                            }
                                                            switch 32767 {
                                                            case -1:
                                                                v59 := GenS1{F0: "\n\t"}
                                                                _ = v59
                                                            case 127:
                                                                v3 = false
                                                            }
                                                            v33 += (v33 / ((v33) | 1))
                                                            v60 := GenS6{F0: 187, F1: 128, F2: 65535}
                                                            v61 := GenS5{F0: v33}
                                                            _ = v60
                                                            _ = v61
                                                        case "":
                                                            fmt.Println(fi48)
                                                        case "a":
                                                            v4 = append(v4, 2147483647)
                                                        }
                                                    }
                                                }
                                            default:
                                                v2 += v2
                                            }
                                        }
                                        v3 = v3
                                    }
                                    _ = fi52
                                }
                            case "foo bar":
                                var v62 uint32 = 42
                                _ = v62
                            default:
                                v63 := GenS4{F0: 42, F1: 796}
                                _ = v63
                            }
                            for fi64 := 0; fi64 < 16; fi64++ {
                                var v65 float32 = v40
                                var v66 uint8 = 0
                                switch "a" {
                                case "hello":
                                    v67 := GenS2{F0: fi64, F1: 18446744073709551615}
                                    _ = v67
                                }
                                _ = fi64
                                _ = v65
                                _ = v66
                            }
                            fmt.Println(v33, v33, v3)
                            _ = fi48
                        }
                    } else {
                        var v68 int16 = -1
                        _ = v68
                    }
                }
                _ = v40
            }
            var v69 int64 = -9223372036854775808
            _ = v69
        } else {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        }
    }
    v1 = (v1 / ((v1) | 1))
    var v70 int64 = 0
    v4 = append(v4, 1)
    for fi71 := 0; fi71 < 18; fi71++ {
        v72 := GenS4{F0: 18446744073709551615, F1: 96}
        for fi73 := 0; fi73 < 2; fi73++ {
            v74 := GenS4{F0: 2, F1: 255}
            v33 = v33
            v2 = (v2 ^ fi71)
            _ = fi73
            _ = v74
        }
        v33 = v33
        _ = fi71
        _ = v72
    }
    if ((v70 - (v70 * (-(v70)))) != (v70 ^ v70)) {
        fmt.Println(v3, v33, v70)
        v75 := GenS7{F0: 0}
        v70 -= (v70 / ((9223372036854775807) | 1))
        _ = v75
    }
    if (((v1 - (^(v1))) >= 42) || (255 <= 0)) {
        fmt.Println(v1, v3, v70)
        for fi76 := 0; fi76 < 7; fi76++ {
            fmt.Println(v1, fi76, v70)
            switch fi76 {
            case -991:
                fmt.Println(v2)
            case -1:
                v1 -= 767
            case 4294967295:
                for fi77 := 0; fi77 < 11; fi77++ {
                    v70 = v70
                    var v78 bool = (v3 || false)
                    v2 = fi76
                    _ = fi77
                    _ = v78
                }
            }
            _ = fi76
        }
    }
    switch ((("" + "hello") + "") + ("a" + "a")) {
    case "hello":
        var v79 bool = v3
        _ = v79
    }
    switch ("foo bar" + "") {
    case "a":
        v80 := GenS2{F0: (v2 % ((v2) | 1)), F1: 18446744073709551615}
        _ = v80
    case "foo bar":
        var v81 uint64 = 759
        _ = v81
    case "\n\t":
        v3 = v3
    default:
        switch -378 {
        case 1:
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        case -981:
            switch "foo bar" {
            case "a":
                v3 = v3
            }
        case 593:
            _ = len(v4)
        }
    }
    switch 4294967295 {
    case -985:
        if v3 {
            for fi82 := 0; fi82 < 20; fi82++ {
                var v83 int16 = int16(v33)
                if (4294967295 == 18446744073709551615) {
                    var v84 uint8 = 42
                    for fi85 := 0; fi85 < 6; fi85++ {
                        var v86 int64 = -883
                        fmt.Println(v2, fi85)
                        var v87 int16 = -491
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    switch "a" {
                    case "hello":
                        if (1 != v2) {
                            if false {
                                v83 = v83
                            } else {
                                v3 = (!(true && ((!v3) && (2 == 95))))
                                v88 := GenS8{F0: 32767}
                                v1 -= v1
                                _ = v88
                            }
                        }
                    case "":
                        if ((v1 - -2147483648) > (v1 - (v1 ^ v1))) {
                            fmt.Println(v84, v83)
                        } else {
                            for fi89 := 0; fi89 < 8; fi89++ {
                                if true {
                                    switch (fi89 / ((int(v70)) | 1)) {
                                    case 326:
                                        if (-2147483648 < v1) {
                                            var v90 int32 = int32(v84)
                                            _ = v90
                                        } else {
                                            v2 -= v2
                                        }
                                    case -509:
                                        v33 = int8(fi89)
                                    default:
                                        var v91 string = "hello"
                                        fmt.Println(fi82, v2, v3)
                                        switch fi82 {
                                        case -1:
                                            v84 -= v84
                                        case -9223372036854775808:
                                            v92 := GenS9{F0: fi82, F1: 775, F2: int(v1)}
                                            _ = v92
                                        case 32768:
                                            v91 = (("" + (v91 + ((v91 + v91) + v91))) + (v91 + v91))
                                        default:
                                            if v3 {
                                                v83 -= 0
                                                if (false || true) {
                                                    fmt.Println(v1, v3, v91)
                                                    if ((((((v91 + v91) + v91) + "\n\t") + "a") < v91) || (4294967295 < 255)) {
                                                        v2 = (fi89 >> ((uint(v84)) & 63))
                                                        v93 := GenS8{F0: 2}
                                                        _ = v93
                                                    }
                                                    var v94 float32 = (0.1 * 3.14159)
                                                    switch (v91 + ("foo bar" + "\n\t")) {
                                                    case "":
                                                        if (((v33 * (v33 * (v33 + v33))) >= (v33 + (v33 | -1))) || (!(v3 && v3))) {
                                                            switch 1 {
                                                            case 32768:
                                                                if len(v4) > 0 {
                                                                    v4[0] = 2147483648
                                                                }
                                                            case -2147483648:
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            default:
                                                                fmt.Println(v2, v1)
                                                            }
                                                        } else {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            var v95 uint64 = 511
                                                            _ = v95
                                                        }
                                                    }
                                                    _ = v94
                                                } else {
                                                    var v96 uint8 = v84
                                                    fmt.Println(v96, v33)
                                                    v2 = (fi82 & 9223372036854775807)
                                                    _ = v96
                                                }
                                            }
                                        }
                                        if ((!v3) && (906 == 18446744073709551615)) {
                                            var v97 bool = (v3 && ((fi89 ^ (fi82 >> ((1) & 63))) >= (fi82 >> ((4294967295) & 63))))
                                            _ = v97
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        var v98 uint16 = 255
                                        _ = v91
                                        _ = v98
                                    }
                                    v3 = (1.5 > 1.5)
                                } else {
                                    v99 := GenS10{F0: v33}
                                    v83 = (v83 % ((v83) | 1))
                                    _ = v99
                                }
                                _ = fi89
                            }
                        }
                    case "foo bar":
                        v2 = (^(fi82))
                    }
                    _ = v84
                }
                _ = fi82
                _ = v83
            }
            v1 -= -2147483648
            v3 = (-2.25 <= -2.25)
        }
    case 128:
        switch "a" {
        case "foo bar":
            var v100 string = ("" + "")
            _ = v100
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v70
}
