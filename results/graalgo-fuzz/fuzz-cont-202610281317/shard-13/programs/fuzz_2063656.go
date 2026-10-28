package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint; F1 bool; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS8 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint16 = 756
    var v2 int16 = 32767
    var v3 uint = 255
    v4 := []uint32{4294967295, 976, 300, 42}
    for fi5 := 0; fi5 < 3; fi5++ {
        switch (fi5 % ((fi5) | 1)) {
        case -81:
            var v6 uint32 = 495
            _ = v6
        case 42:
            v7 := GenS1{F0: (^(v3))}
            _ = v7
        case -1:
            fmt.Println(v2, v1, v1)
        }
        _ = fi5
    }
    var v8 int16 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v8 -= 1
    v9 := GenS1{F0: 1}
    v2 = v2
    switch 127 {
    case 2147483647:
        var v10 bool = ((127 != -116) && true)
        _ = v10
    case -492:
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                if true {
                    switch 2147483647 {
                    case 42:
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v3, v1, v3)
                        case "a":
                            switch ("hello" + "hello") {
                            case "hello":
                                fmt.Println(v1, v2)
                            case "a":
                                switch fi11 {
                                case -654:
                                    var v13 int64 = -1
                                    _ = v13
                                default:
                                    if (!false) {
                                        _ = len(v4)
                                    }
                                }
                            case "\n\t":
                                v14 := GenS1{F0: 0}
                                _ = v14
                            }
                        case "hello":
                            switch ("" + (("foo bar" + ("hello" + "")) + ("foo bar" + ""))) {
                            case "":
                                var v15 string = "\n\t"
                                _ = v15
                            case "hello":
                                v1 -= v1
                                if (4294967295 <= 0) {
                                    fmt.Println(fi11)
                                }
                            case "a":
                                v8 = (v2 >> ((v3) & 63))
                            }
                        default:
                            if len(v4) > 0 {
                                v4[0] = 42
                            }
                        }
                    case 32767:
                        var v16 uint8 = 42
                        _ = v16
                    }
                    v8 = v8
                    switch 2147483647 {
                    case 32767:
                        if (-1 <= 0) {
                            if false {
                                v8 = (v2 / ((v8) | 1))
                                v8 = v2
                            }
                        }
                    case 127:
                        if (((0.1 - ((1e10 / 0.1) / 1.0)) != (1e10 / 3.14159)) || true) {
                            fmt.Println(fi11, fi11)
                            v2 -= (v8 & (v2 + v8))
                        }
                    case 0:
                        if true {
                            var v17 uint16 = v1
                            _ = v17
                        } else {
                            fmt.Println(v3, v2, v1)
                            if len(v4) > 0 {
                                v4[0] = 2147483647
                            }
                            fmt.Println(v2, v3)
                        }
                    }
                } else {
                    if true {
                        var v18 float64 = float64(((3.14159 * ((1e10 + 0.1) * 3.14159)) / 1.0))
                        v8 += (v2 & (v2 / ((v8) | 1)))
                        v19 := GenS1{F0: (v3 << (((v3 >> (((v3 * (v3 % ((v3) | 1)))) & 63))) & 63))}
                        _ = v18
                        _ = v19
                    } else {
                        if (-9223372036854775808 == 387) {
                            if false {
                                v3 -= (v3 << (((v3 * 1)) & 63))
                                v1 = 65535
                                v2 = (v2 + (v2 >> (((v3 ^ (v3 << (((v3 - v3)) & 63)))) & 63)))
                                var v20 uint16 = (v1 | 2)
                                v1 += (v20 - (v20 ^ (^(v1))))
                                _ = v20
                            }
                        } else {
                            v8 += (^(v8))
                        }
                        v2 += v8
                        var v21 string = "hello"
                        _ = v21
                    }
                    if ("" != "hello") {
                        var v22 int = fi11
                        _ = len(v4)
                        var v23 int = fi12
                        _ = v22
                        _ = v23
                    } else {
                        v4 = append(v4, 601)
                        var v24 uint32 = 657
                        fmt.Println(fi12, v1, v3)
                        v1 = 42
                        _ = v24
                    }
                }
                v1 -= (^(v1))
                _ = fi12
            }
            _ = fi11
        }
    case -187:
        for fi25 := 0; fi25 < 6; fi25++ {
            if (("\n\t" + "a") <= "foo bar") {
                fmt.Println(v3, v8, v8)
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 int8 = 89
                    _ = fi26
                    _ = v27
                }
                var v28 int8 = int8(v1)
                _ = v28
            } else {
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 int32 = -1
                    _ = fi29
                    _ = v30
                }
                if (!true) {
                    v2 -= -434
                    v2 += (v8 | -1)
                    var v31 int8 = -1
                    _ = v31
                }
            }
            v4 = append(v4, 325)
            _ = fi25
        }
    }
    switch "" {
    case "a":
        if false {
            var v32 string = (("hello" + "foo bar") + "foo bar")
            _ = v32
        } else {
            switch "foo bar" {
            case "a":
                v8 = v8
            case "":
                v2 += (v2 - (v8 * v2))
            case "hello":
                v33 := GenS1{F0: (^(v3))}
                _ = v33
            default:
                v34 := GenS2{F0: 653, F1: "\n\t", F2: 184}
                _ = v34
            }
            if (float32(0.1) == (1e10 * float32((float64((0.1 - 1e10)) + 1.5)))) {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                switch 848 {
                case -650:
                    v1 = v1
                case 2147483647:
                    v3 = v3
                case 32767:
                    _ = len(v4)
                default:
                    fmt.Println(v8)
                }
            }
            for fi35 := 0; fi35 < 6; fi35++ {
                switch "" {
                case "a":
                    v2 += -1
                case "\n\t":
                    fmt.Println(v1, v8)
                case "":
                    for fi36 := 0; fi36 < 3; fi36++ {
                        if (375 > -2147483648) {
                            if len(v4) > 0 {
                                v4[0] = 98
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if ((!false) && false) {
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = 65535
                                    }
                                }
                            }
                        }
                        _ = fi36
                    }
                default:
                    var v37 int8 = 54
                    _ = v37
                }
                var v38 uint8 = 2
                var v39 uint64 = 1
                _ = fi35
                _ = v38
                _ = v39
            }
        }
    case "\n\t":
        switch 388 {
        case 587:
            var v40 int = 128
            _ = v40
        case 32768:
            v41 := GenS1{F0: (v3 | 0)}
            _ = v41
        case 32767:
            v2 = (v2 - v2)
        }
    case "foo bar":
        for fi42 := 0; fi42 < 3; fi42++ {
            v43 := GenS1{F0: (v3 & v3)}
            _ = fi42
            _ = v43
        }
    }
    if false {
        if false {
            v8 += 218
            for fi44 := 0; fi44 < 6; fi44++ {
                if false {
                    if (!false) {
                        for fi45 := 0; fi45 < 20; fi45++ {
                            v2 += (v2 - (v8 << (((v3 % ((v3) | 1))) & 63)))
                            if true {
                                switch fi44 {
                                case 1:
                                    v46 := GenS2{F0: 0, F1: "hello", F2: 220}
                                    _ = v46
                                case -89:
                                    v1 -= v1
                                case 4294967295:
                                    v47 := GenS3{F0: 128}
                                    _ = v47
                                }
                                var v48 uint64 = 589
                                _ = len(v4)
                                _ = v48
                            }
                            _ = fi45
                        }
                        var v49 uint64 = 4294967295
                        var v50 float32 = ((0.1 * (float32(((-2.25 / 1.5) / 1.0)) - 1.5)) + (3.14159 - 1.5))
                        _ = v49
                        _ = v50
                    } else {
                        var v51 bool = true
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 18; fi52++ {
                        if len(v4) > 0 {
                            v4[0] = 1
                        }
                        fmt.Println(fi52, v1)
                        _ = fi52
                    }
                } else {
                    if false {
                        v1 += v1
                        v1 = (v1 << (((v3 ^ v3)) & 63))
                        if (v8 < v8) {
                            var v53 int = (fi44 >> (((v3 << (((v3 >> ((v3) & 63))) & 63))) & 63))
                            _ = v53
                        } else {
                            v1 -= v1
                            switch (fi44 * (fi44 - int(v1))) {
                            case 229:
                                if (2 < 1) {
                                    if false {
                                        fmt.Println(v1, v2)
                                        switch "foo bar" {
                                        case "":
                                            if true {
                                                switch fi44 {
                                                case -25:
                                                    v54 := GenS1{F0: v3}
                                                    _ = v54
                                                case 2147483647:
                                                    fmt.Println(v2, v3)
                                                case 4294967295:
                                                    v1 = 255
                                                }
                                                v2 -= (^(v8))
                                                v8 += -274
                                            } else {
                                                v3 = 4294967295
                                            }
                                        case "a":
                                            fmt.Println(v3, fi44)
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            var v55 uint = v3
                                            _ = v55
                                        }
                                        var v56 uint = v3
                                        _ = v56
                                    }
                                    v1 += 0
                                    v8 = 725
                                } else {
                                    fmt.Println(v1, v1, v2)
                                    v3 = v3
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            }
                        }
                    } else {
                        for fi57 := 0; fi57 < 20; fi57++ {
                            v3 = (v3 | v3)
                            var v58 uint16 = 319
                            if false {
                                fmt.Println(v2, fi57, v3)
                            }
                            _ = fi57
                            _ = v58
                        }
                        var v59 int64 = 9223372036854775807
                        _ = v59
                    }
                    v3 -= (v3 & (v3 | 1))
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v60 int64 = 43
                _ = fi44
                _ = v60
            }
            if ((!true) && true) {
                v1 -= 753
                v1 += (v1 | (^(v1)))
                var v61 string = "\n\t"
                _ = v61
            }
        }
        var v62 int32 = 0
        _ = v62
    } else {
        switch 346 {
        case -1:
            switch ("\n\t" + ("a" + "")) {
            case "hello":
                v4 = append(v4, 0)
            case "foo bar":
                v3 = v3
            case "":
                var v63 uint16 = (v1 & (^(v1)))
                _ = v63
            }
        case 68:
            if (678 < 4294967295) {
                v1 += (v1 + 2)
            } else {
                switch int(v1) {
                case -926:
                    if (1.5 == (float32(-2.25) * float32((float64(1e10) * ((float64(3.14159) + 3.14159) + 1.5))))) {
                        _ = len(v4)
                    } else {
                        switch (("" + "a") + "hello") {
                        case "foo bar":
                            v64 := GenS2{F0: 2, F1: "\n\t", F2: 173}
                            _ = v64
                        case "\n\t":
                            for fi65 := 0; fi65 < 5; fi65++ {
                                switch ("foo bar" + ("\n\t" + "hello")) {
                                case "foo bar":
                                    v8 = v8
                                }
                                _ = fi65
                            }
                        default:
                            _ = len(v4)
                        }
                        if ((!((true && (false && ("hello" >= "hello"))) || false)) || true) {
                            if (!(188 == 42)) {
                                if true {
                                    v66 := GenS3{F0: 1}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v2 -= v8
                                    _ = v66
                                }
                            }
                        }
                        v67 := GenS4{F0: (v3 / ((v3) | 1)), F1: false, F2: int8(v1)}
                        _ = v67
                    }
                default:
                    v68 := GenS2{F0: 1, F1: ("\n\t" + "foo bar"), F2: 249}
                    _ = v68
                }
            }
        case 9223372036854775807:
            switch -412 {
            case 42:
                v2 -= (v2 | (v2 / ((0) | 1)))
            case 4294967295:
                _ = len(v4)
            default:
                v1 = v1
            }
        default:
            if true {
                for fi69 := 0; fi69 < 4; fi69++ {
                    v8 = v2
                    v1 = v1
                    if ((v2 | v2) > (v8 / ((v2) | 1))) {
                        v1 += (v1 ^ v1)
                    } else {
                        for fi70 := 0; fi70 < 8; fi70++ {
                            var v71 bool = (196 <= 32)
                            _ = fi70
                            _ = v71
                        }
                        if (2147483647 > 0) {
                            switch "\n\t" {
                            case "":
                                v72 := GenS5{F0: (v1 + v1)}
                                _ = v72
                            default:
                                if (fi69 >= -815) {
                                    if false {
                                        var v73 uint8 = 255
                                        _ = v73
                                    } else {
                                        v74 := GenS1{F0: (v3 | v3)}
                                        v75 := GenS3{F0: 102}
                                        _ = v74
                                        _ = v75
                                    }
                                } else {
                                    var v76 int16 = -561
                                    _ = v76
                                }
                            }
                        }
                        fmt.Println(v2)
                    }
                    _ = fi69
                }
            } else {
                v3 = 42
                for fi77 := 0; fi77 < 1; fi77++ {
                    v4 = append(v4, 2147483647)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if len(v4) > 0 {
                        v4[0] = 407
                    }
                    _ = fi77
                }
                if (!(uint64(v3) <= 18446744073709551615)) {
                    if false {
                        for fi78 := 0; fi78 < 11; fi78++ {
                            var v79 int64 = -9223372036854775808
                            fmt.Println(fi78)
                            _ = fi78
                            _ = v79
                        }
                        switch 628 {
                        case 2147483647:
                            v1 -= v1
                        default:
                            if (!false) {
                                v80 := GenS1{F0: 42}
                                v2 = (v8 % (((v2 | v8)) | 1))
                                _ = v80
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(v3, v3)
                            fmt.Println(v3, v3)
                        } else {
                            v81 := GenS6{F0: (47 != 128), F1: (789 > 1)}
                            _ = v81
                        }
                        v82 := GenS7{F0: 2, F1: 1, F2: 2}
                        for fi83 := 0; fi83 < 20; fi83++ {
                            var v84 int8 = 0
                            _ = fi83
                            _ = v84
                        }
                        _ = v82
                    }
                    for fi85 := 0; fi85 < 1; fi85++ {
                        for fi86 := 0; fi86 < 7; fi86++ {
                            v87 := GenS8{F0: (0 < 65535), F1: 0}
                            var v88 int32 = 1
                            _ = fi86
                            _ = v87
                            _ = v88
                        }
                        for fi89 := 0; fi89 < 13; fi89++ {
                            var v90 int8 = 0
                            var v91 uint8 = 2
                            _ = fi89
                            _ = v90
                            _ = v91
                        }
                        _ = fi85
                    }
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
