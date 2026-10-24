package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int16 }
type GenS5 struct { F0 uint64; F1 int64; F2 string }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 uint; F1 uint32; F2 bool }
type GenS8 struct { F0 uint32; F1 int8; F2 int8 }
type GenS9 struct { F0 int16; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint8; F1 int16 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool; F1 uint64 }
type GenS13 struct { F0 uint32 }

func main() {
    var v1 int32 = 1
    var v2 uint8 = 255
    var v3 uint = 717
    v1 = v1
    switch 42 {
    case 127:
        fmt.Println(v1, v3)
    case 9223372036854775807:
        for fi4 := 0; fi4 < 19; fi4++ {
            fmt.Println(v3, v2, v2)
            _ = fi4
        }
    case -562:
        v2 = 128
    }
    fmt.Println(v1)
    v5 := []int{42}
    var v6 int16 = 42
    fmt.Println(v3, v3)
    v7 := GenS1{F0: 4294967295}
    for fi8 := 0; fi8 < 12; fi8++ {
        switch 533 {
        case -934:
            var v9 uint32 = 859
            _ = v9
        case -182:
            v10 := GenS1{F0: 1}
            _ = v10
        case 32768:
            var v11 bool = true
            _ = v11
        default:
            switch (fi8 % ((127) | 1)) {
            case -1:
                for fi12 := 0; fi12 < 1; fi12++ {
                    fmt.Println(fi12, fi12)
                    v1 -= -224
                    _ = fi12
                }
            case 42:
                if true {
                    v6 += v6
                } else {
                    switch (fi8 % ((fi8) | 1)) {
                    case 9223372036854775807:
                        for fi13 := 0; fi13 < 14; fi13++ {
                            var v14 int16 = (v6 & v6)
                            v2 = (v2 ^ 255)
                            _ = fi13
                            _ = v14
                        }
                    case -582:
                        v2 -= v2
                    case 0:
                        if false {
                            var v15 uint32 = 1
                            if true {
                                v1 += 2147483647
                            } else {
                                for fi16 := 0; fi16 < 8; fi16++ {
                                    _ = len(v5)
                                    _ = fi16
                                }
                                v1 = (v1 >> ((645) & 63))
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                            }
                            _ = v15
                        } else {
                            switch int(v1) {
                            case -391:
                                v6 = (-(v6))
                            default:
                                for fi17 := 0; fi17 < 18; fi17++ {
                                    if false {
                                        var v18 bool = ((^(v6)) <= 1000)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        var v19 int16 = v6
                                        _ = v18
                                        _ = v19
                                    }
                                    fmt.Println(v1, fi8, v2)
                                    if (false && false) {
                                        var v20 uint = v3
                                        var v21 uint16 = 32767
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (false || false) {
                                            v2 += v2
                                        }
                                        v5 = append(v5, fi8)
                                    }
                                    _ = fi17
                                }
                            }
                        }
                    }
                }
            default:
                switch (fi8 * fi8) {
                case 32767:
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case -36:
                    if (v1 >= v1) {
                        fmt.Println(v1)
                        v2 = 0
                        v5 = append(v5, (fi8 / ((42) | 1)))
                    } else {
                        v1 = (v1 * 584)
                        var v22 uint8 = 2
                        v23 := GenS2{F0: (v22 << (((^(v3))) & 63)), F1: 42, F2: int32(v2)}
                        _ = v22
                        _ = v23
                    }
                }
            }
        }
        switch (("a" + "a") + (("\n\t" + "\n\t") + ("hello" + "\n\t"))) {
        case "a":
            var v24 uint = (v3 + 65535)
            _ = v24
        case "\n\t":
            v1 = (^(v1))
        case "hello":
            var v25 uint16 = 1
            _ = v25
        default:
            for fi26 := 0; fi26 < 1; fi26++ {
                v3 -= (v3 >> ((uint(v2)) & 63))
                v6 += int16(fi8)
                _ = fi26
            }
        }
        _ = fi8
    }
    switch "hello" {
    case "hello":
        if false {
            switch -947 {
            case -81:
                var v27 string = "hello"
                _ = v27
            default:
                var v28 int8 = 11
                _ = v28
            }
            var v29 uint64 = 4294967295
            var v30 float32 = (1e10 - 1.5)
            _ = v29
            _ = v30
        }
    case "":
        v1 = v1
    case "foo bar":
        v31 := GenS3{F0: 121}
        _ = v31
    }
    if ((v1 >= (v1 >> (((v3 | uint(v2))) & 63))) || (!true)) {
        for fi32 := 0; fi32 < 12; fi32++ {
            fmt.Println(v6, v2, v2)
            _ = fi32
        }
        v2 = 2
        if false {
            for fi33 := 0; fi33 < 11; fi33++ {
                switch fi33 {
                case 32767:
                    if len(v5) > 0 {
                        v5[0] = fi33
                    }
                case -9223372036854775808:
                    if (!(v6 > (v6 << (((^(v3))) & 63)))) {
                        v5 = append(v5, (^(fi33)))
                        var v34 uint64 = 4294967295
                        _ = v34
                    } else {
                        v2 = (^(v2))
                        var v35 int16 = (v6 % (((v6 + v6)) | 1))
                        v1 += (-(v1))
                        _ = v35
                    }
                default:
                    if (!(18446744073709551615 != 42)) {
                        switch int(v2) {
                        case -2147483648:
                            v36 := GenS1{F0: (-(fi33))}
                            _ = v36
                        case 4294967295:
                            switch "\n\t" {
                            case "foo bar":
                                for fi37 := 0; fi37 < 14; fi37++ {
                                    v6 += (v6 >> ((255) & 63))
                                    v1 = (v1 % (((v1 * (-(v1)))) | 1))
                                    _ = fi37
                                }
                            case "a":
                                v5 = append(v5, (fi33 ^ (fi33 * fi33)))
                            default:
                                for fi38 := 0; fi38 < 17; fi38++ {
                                    if (true || (false && false)) {
                                        fmt.Println(fi33, v6)
                                        fmt.Println(fi33, v1)
                                        if true {
                                            v1 = (v1 - -341)
                                            if (!((v3 < v3) || true)) {
                                                v1 -= (v1 + v1)
                                            }
                                            if true {
                                                v3 = v3
                                                if ((float64(-2.25) <= (((1.5 - -2.25) * 1e10) + 0.0)) && false) {
                                                    if false {
                                                        v1 -= 1
                                                    }
                                                } else {
                                                    v2 = v2
                                                    var v39 string = "a"
                                                    _ = v39
                                                }
                                            } else {
                                                var v40 uint8 = (v2 & v2)
                                                _ = v40
                                            }
                                        }
                                    } else {
                                        fmt.Println(v2)
                                        var v41 bool = true
                                        _ = len(v5)
                                        _ = v41
                                    }
                                    switch fi38 {
                                    case -1:
                                        if (!(!((v6 + v6) < v6))) {
                                            v42 := GenS4{F0: 1, F1: (v2 | v2), F2: -765}
                                            _ = len(v5)
                                            if (int8(v1) != 100) {
                                                v43 := GenS1{F0: fi33}
                                                var v44 uint8 = v2
                                                _ = v43
                                                _ = v44
                                            } else {
                                                fmt.Println(v6)
                                                if true {
                                                    v45 := GenS5{F0: 0, F1: 0, F2: ("\n\t" + "foo bar")}
                                                    var v46 uint32 = 65535
                                                    var v47 uint16 = 854
                                                    _ = v45
                                                    _ = v46
                                                    _ = v47
                                                }
                                                v3 = (v3 + v3)
                                            }
                                            _ = v42
                                        }
                                        var v48 float32 = float32(1e10)
                                        var v49 uint8 = v2
                                        _ = v48
                                        _ = v49
                                    case 128:
                                        if false {
                                            v5 = append(v5, (fi38 ^ (-(fi38))))
                                            var v50 string = "a"
                                            switch int(v1) {
                                            case -730:
                                                v3 = 0
                                            case 0:
                                                switch -1 {
                                                case 543:
                                                    switch v50 {
                                                    case "\n\t":
                                                        var v51 int16 = (v6 - 0)
                                                        var v52 bool = false
                                                        _ = v51
                                                        _ = v52
                                                    case "a":
                                                        if false {
                                                            var v53 int64 = 90
                                                            switch (fi33 % (((-(fi38))) | 1)) {
                                                            case 42:
                                                                var v54 uint = 1
                                                                _ = v54
                                                            case 9223372036854775807:
                                                                v5 = append(v5, (fi38 / ((fi33) | 1)))
                                                            default:
                                                                var v55 int16 = (v6 * (^(v6)))
                                                                _ = v55
                                                            }
                                                            v3 = v3
                                                            _ = v53
                                                        }
                                                    default:
                                                        var v56 uint16 = uint16(v2)
                                                        _ = v56
                                                    }
                                                case 32767:
                                                    switch (fi33 + 0) {
                                                    case 32768:
                                                        v50 = "foo bar"
                                                    default:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    v6 -= (v6 - v6)
                                                    switch fi38 {
                                                    case 1:
                                                        v5 = append(v5, fi33)
                                                    case 128:
                                                        if false {
                                                            fmt.Println(v50)
                                                        } else {
                                                            var v57 uint8 = 166
                                                            if true {
                                                                v50 = v50
                                                            } else {
                                                                var v58 string = "\n\t"
                                                                _ = v58
                                                            }
                                                            v2 -= (^(v2))
                                                            _ = v57
                                                        }
                                                    case 42:
                                                        switch ("hello" + (((v50 + (v50 + (("a" + ((v50 + v50) + (v50 + "hello"))) + "a"))) + v50) + v50)) {
                                                        case "\n\t":
                                                            var v59 int32 = 1
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                            case 1:
                                                var v60 float64 = 0.1
                                                _ = v60
                                            default:
                                                var v61 int8 = 1
                                                _ = v61
                                            }
                                            _ = v50
                                        }
                                        if len(v5) > 0 {
                                            v5[0] = (-(fi38))
                                        }
                                        if (v2 < v2) {
                                            v2 = 2
                                        } else {
                                            if ((v1 | v1) < v1) {
                                                fmt.Println(fi38)
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            var v62 uint32 = 2
                                            _ = v62
                                        }
                                    default:
                                        switch 486 {
                                        case 742:
                                            v5 = append(v5, fi33)
                                        case -408:
                                            if false {
                                                var v63 float64 = float64(1e10)
                                                _ = v63
                                            }
                                        }
                                    }
                                    _ = fi38
                                }
                            }
                        }
                        if false {
                            for fi64 := 0; fi64 < 7; fi64++ {
                                if (65535 >= v3) {
                                    var v65 float64 = 1.5
                                    fmt.Println(v6, v2, fi64)
                                    _ = v65
                                }
                                if false {
                                    var v66 bool = (!false)
                                    v67 := GenS6{F0: (v6 << (((v3 | v3)) & 63)), F1: 2147483648}
                                    v2 = (v2 % (((^(v2))) | 1))
                                    v1 -= (-(v1))
                                    _ = v66
                                    _ = v67
                                }
                                v2 += v2
                                _ = fi64
                            }
                            v6 = (v6 << ((v3) & 63))
                        }
                        v2 = (v2 << (((v3 + (^(v3)))) & 63))
                    }
                }
                var v68 float32 = 3.14159
                for fi69 := 0; fi69 < 10; fi69++ {
                    if (100 < -1) {
                        v5 = append(v5, fi69)
                        if (65535 >= 42) {
                            if true {
                                v3 += 4294967295
                                v68 -= v68
                            } else {
                                var v70 int8 = -25
                                switch (fi33 << ((v3) & 63)) {
                                case 287:
                                    v68 = 3.14159
                                default:
                                    fmt.Println(v3, v70)
                                }
                                _ = v70
                            }
                            var v71 int = fi33
                            _ = v71
                        } else {
                            var v72 uint16 = 42
                            _ = v72
                        }
                    }
                    v3 = v3
                    v68 = -2.25
                    if true {
                        v2 -= (^(v2))
                    }
                    _ = fi69
                }
                _ = fi33
                _ = v68
            }
            for fi73 := 0; fi73 < 2; fi73++ {
                v74 := GenS3{F0: 255}
                switch (fi73 + (fi73 ^ (fi73 + (fi73 / (((fi73 ^ (fi73 & 128))) | 1))))) {
                case 9223372036854775807:
                    v2 += (v2 | v2)
                case 4294967295:
                    switch fi73 {
                    case -2147483648:
                        v3 += (^(v3))
                    case 32768:
                        var v75 int8 = -108
                        _ = v75
                    case -1:
                        v76 := GenS7{F0: (v3 + (v3 ^ 42)), F1: 956, F2: ((fi73 >> (((v3 >> (((v3 & uint(v2))) & 63))) & 63)) < 1)}
                        _ = v76
                    }
                case 281:
                    v1 -= 0
                default:
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if true {
                            var v77 string = "foo bar"
                            if (773 == 961) {
                                var v78 bool = ((v6 / (((v6 * v6)) | 1)) == v6)
                                v6 = 0
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v78
                            } else {
                                v79 := GenS4{F0: 9223372036854775807, F1: v2, F2: v6}
                                _ = v79
                            }
                            _ = v77
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        switch ("hello" + "hello") {
                        case "a":
                            var v80 int64 = int64(v1)
                            _ = v80
                        case "":
                            v6 -= (^(v6))
                        }
                        v2 = (v2 & v2)
                    }
                }
                switch (((("a" + "a") + "a") + "foo bar") + (("a" + "foo bar") + "hello")) {
                case "a":
                    v5 = append(v5, fi73)
                case "":
                    var v81 uint = v3
                    _ = v81
                case "hello":
                    for fi82 := 0; fi82 < 16; fi82++ {
                        fmt.Println(fi82, fi73)
                        _ = fi82
                    }
                }
                _ = fi73
                _ = v74
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 -= v3
        }
    } else {
        var v83 int16 = -32768
        switch ((("hello" + ("a" + "hello")) + "a") + "a") {
        case "hello":
            switch "foo bar" {
            case "":
                switch -978 {
                case -913:
                    if (true || false) {
                        v84 := GenS8{F0: 702, F1: -1, F2: -123}
                        _ = v84
                    }
                case -195:
                    var v85 uint16 = 1
                    _ = v85
                case 2147483647:
                    fmt.Println(v2)
                }
            case "a":
                switch 1 {
                case -968:
                    v86 := GenS9{F0: (-(v6)), F1: 0, F2: 42}
                    _ = v86
                case 32768:
                    var v87 bool = ((true || true) || true)
                    _ = v87
                case 32767:
                    v6 = 1
                }
            default:
                for fi88 := 0; fi88 < 6; fi88++ {
                    var v89 uint8 = uint8(v3)
                    if true {
                        for fi90 := 0; fi90 < 3; fi90++ {
                            fmt.Println(v6, fi90)
                            _ = fi90
                        }
                    } else {
                        fmt.Println(v83)
                        v91 := GenS10{F0: v2, F1: (^(v6))}
                        _ = v91
                    }
                    for fi92 := 0; fi92 < 16; fi92++ {
                        v93 := GenS11{F0: (2 == 18446744073709551615)}
                        _ = fi92
                        _ = v93
                    }
                    _ = fi88
                    _ = v89
                }
            }
        default:
            v1 = (v1 ^ v1)
        }
        _ = v83
    }
    switch "hello" {
    case "hello":
        var v94 float32 = (0.1 + -2.25)
        _ = v94
    case "\n\t":
        var v95 uint32 = 355
        _ = v95
    case "a":
        v3 = v3
    }
    v96 := GenS12{F0: true, F1: uint64(v3)}
    v1 = (v1 >> (((v3 >> ((2) & 63))) & 63))
    _ = len(v5)
    if (false || (1 >= 4294967295)) {
        v97 := GenS13{F0: 560}
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v96
}
