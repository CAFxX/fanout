package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 uint32; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS9 struct { F0 uint16; F1 string; F2 int64 }
type GenS10 struct { F0 int32; F1 int }
type GenS11 struct { F0 int32; F1 uint; F2 uint16 }
type GenS12 struct { F0 int; F1 int32; F2 string }
type GenS13 struct { F0 int; F1 string; F2 int }

func main() {
    var v1 uint64 = 186
    v2 := make(map[string]int)
    var v3 uint64 = 0
    var v4 int16 = -1
    var v5 string = "hello"
    v6 := []string{}
    switch v5 {
    case "foo bar":
        v7 := GenS1{F0: (!true), F1: v5}
        _ = v7
    case "a":
        v4 = (v4 + 1000)
    case "\n\t":
        _ = len(v6)
    }
    delete(v2, (v5 + v5))
    delete(v2, (v5 + v5))
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v1 = (v1 << ((65535) & 63))
            switch (fi8 >> ((2) & 63)) {
            case 2147483647:
                if true {
                    v4 -= (^(v4))
                } else {
                    delete(v2, ("hello" + ("" + (("hello" + v5) + ""))))
                    delete(v2, v5)
                    var v10 uint = 1
                    _ = len(v6)
                    _ = v10
                }
            case -1:
                v5 = v5
            case 127:
                v3 += 1
            default:
                switch v5 {
                case "hello":
                    v11 := GenS2{F0: (false || (true || true)), F1: 213, F2: v4}
                    _ = v11
                default:
                    var v12 int16 = (-(v4))
                    fmt.Println(fi9)
                    v1 = v1
                    _ = v12
                }
            }
            fmt.Println(fi8, v3)
            _ = fi9
        }
        var v13 int16 = -323
        if false {
            v1 = v1
        }
        _ = fi8
        _ = v13
    }
    switch 4294967295 {
    case 128:
        switch "hello" {
        case "":
            switch v5 {
            case "foo bar":
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch "hello" {
                    case "hello":
                        var v15 int = fi14
                        _ = v15
                    default:
                        var v16 uint = 255
                        _ = v16
                    }
                    _ = fi14
                }
            case "a":
                if ((42 < 42) && (0 >= -128)) {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        fmt.Println(fi17)
                        if false {
                            switch (-(fi17)) {
                            case -9223372036854775808:
                                v3 = v1
                            case 145:
                                for fi18 := 0; fi18 < 10; fi18++ {
                                    if (false && (false && ((false || false) && false))) {
                                        if false {
                                            var v19 int32 = -2147483648
                                            _ = v19
                                        } else {
                                            if true {
                                                var v20 float32 = float32(1e10)
                                                var v21 bool = (uint32(v1) > 2147483648)
                                                v22 := v2[v5]
                                                _ = v20
                                                _ = v21
                                                _ = v22
                                            } else {
                                                fmt.Println(fi18, v5, fi17)
                                            }
                                            switch ("a" + v5) {
                                            case "a":
                                                var v23 int32 = -2147483648
                                                _ = v23
                                            default:
                                                if ((^(v4)) > v4) {
                                                    v24 := GenS3{F0: uint32(v3), F1: 1, F2: 111}
                                                    if false {
                                                        switch 128 {
                                                        case 128:
                                                            switch (fi18 % ((-7) | 1)) {
                                                            case -2147483648:
                                                                v25 := GenS2{F0: true, F1: 127, F2: -1}
                                                                _ = v25
                                                            default:
                                                                fmt.Println(fi18)
                                                            }
                                                        default:
                                                            var v26 bool = (((-(v4)) >= (v4 * v4)) && false)
                                                            _ = v26
                                                        }
                                                        v27 := GenS3{F0: 0, F1: 0, F2: 222}
                                                        _ = v27
                                                    } else {
                                                        var v28 int = fi17
                                                        _ = v28
                                                    }
                                                    _ = v24
                                                }
                                            }
                                            var v29 uint32 = 2
                                            switch (v5 + (("\n\t" + ("" + v5)) + v5)) {
                                            case "a":
                                                v1 = v1
                                            case "hello":
                                                v30 := GenS2{F0: (!true), F1: 42, F2: -1}
                                                _ = v30
                                            }
                                            _ = v29
                                        }
                                        _ = len(v2)
                                    } else {
                                        if false {
                                            switch v5 {
                                            case "hello":
                                                var v31 uint = 0
                                                _ = v31
                                            }
                                            v32 := GenS3{F0: 2147483648, F1: 214, F2: 255}
                                            _ = v32
                                        } else {
                                            if (true && (float64((0.1 + 0.1)) < 1e10)) {
                                                var v33 string = "\n\t"
                                                var v34 int8 = 0
                                                v33 = ((v33 + "hello") + v33)
                                                _ = v33
                                                _ = v34
                                            }
                                            v35 := GenS3{F0: 933, F1: 42, F2: 0}
                                            v36 := GenS2{F0: false, F1: 1, F2: v4}
                                            var v37 int32 = 42
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                        }
                                    }
                                    v38 := GenS4{F0: 32767}
                                    switch (-(fi17)) {
                                    case -1:
                                        v1 = 433
                                    case 9223372036854775807:
                                        switch (^(fi18)) {
                                        case 0:
                                            if (((v4 ^ 1) >= v4) || true) {
                                                var v39 int32 = 0
                                                var v40 int64 = 882
                                                _ = v39
                                                _ = v40
                                            } else {
                                                var v41 bool = false
                                                var v42 string = v5
                                                _ = v41
                                                _ = v42
                                            }
                                        case -563:
                                            v2[(v5 + v5)] = (fi17 << ((2) & 63))
                                        case 4294967295:
                                            v5 = v5
                                        }
                                    case 42:
                                        if true {
                                            var v43 uint = 1
                                            if true {
                                                delete(v2, (v5 + "\n\t"))
                                                v44 := GenS4{F0: 65535}
                                                var v45 float32 = 1e10
                                                _ = v44
                                                _ = v45
                                            }
                                            _ = v43
                                        }
                                        if false {
                                            var v46 float32 = 1.5
                                            fmt.Println(v5)
                                            _ = v46
                                        } else {
                                            switch (fi17 & (fi17 ^ (fi17 + (fi18 << ((0) & 63))))) {
                                            case 32767:
                                                if false {
                                                    if (668 == 814) {
                                                        v3 += v1
                                                        var v47 bool = ("a" >= v5)
                                                        var v48 uint64 = v3
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    fmt.Println(v5, v1, fi18)
                                                }
                                            case -2:
                                                fmt.Println(v3, v1, fi17)
                                            }
                                        }
                                        v49 := GenS5{F0: 929, F1: 0, F2: 1}
                                        switch -9223372036854775808 {
                                        case -2147483648:
                                            if (4294967295 != 2) {
                                                switch (fi18 >> ((255) & 63)) {
                                                case -9223372036854775808:
                                                    var v50 string = v5
                                                    _ = v50
                                                default:
                                                    switch v5 {
                                                    case "":
                                                        v2["hello"] = 128
                                                    case "hello":
                                                        v5 = v5
                                                    case "a":
                                                        if (159 < 0) {
                                                            if true {
                                                                fmt.Println(v1)
                                                                v51 := GenS4{F0: 32767}
                                                                fmt.Println(fi18)
                                                                v4 = v4
                                                                if ((!(false || ("foo bar" >= v5))) || false) {
                                                                    if (67 > 127) {
                                                                        v6 = append(v6, v5)
                                                                        v3 += (v1 % ((42) | 1))
                                                                        var v52 int16 = (v4 % ((42) | 1))
                                                                        _ = v52
                                                                    }
                                                                    v3 = v3
                                                                    var v53 uint8 = 226
                                                                    _ = v53
                                                                }
                                                                _ = v51
                                                            }
                                                        }
                                                    default:
                                                        _ = len(v2)
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1, v5, fi17)
                                            }
                                        }
                                        if true {
                                            fmt.Println(fi18, v4, v4)
                                            var v54 int64 = -124
                                            if (false || false) {
                                                var v55 uint32 = 4294967295
                                                fmt.Println(fi17, v3)
                                                v56 := GenS5{F0: 18446744073709551615, F1: (v55 / ((2) | 1)), F2: 1}
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = v54
                                        }
                                        v1 = v1
                                        var v57 uint32 = 65535
                                        v5 = v5
                                        var v58 int32 = -175
                                        var v59 uint = 348
                                        v60 := GenS2{F0: false, F1: uint8(v1), F2: 1000}
                                        _ = v49
                                        _ = v57
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    switch fi17 {
                                    case 907:
                                        v3 = v1
                                    case 4294967295:
                                        v61 := GenS4{F0: 2}
                                        _ = v61
                                    case 32768:
                                        var v62 int64 = 1
                                        _ = v62
                                    }
                                    _ = fi18
                                    _ = v38
                                }
                            case 42:
                                var v63 string = (v5 + ("\n\t" + v5))
                                _ = v63
                            }
                        }
                        switch (^(fi17)) {
                        case 32768:
                            v64 := GenS6{F0: 42}
                            _ = v64
                        case -9223372036854775808:
                            if false {
                                delete(v2, "a")
                            }
                        case 25:
                            if true {
                                v3 = v3
                                switch (v5 + ((v5 + ("foo bar" + v5)) + "\n\t")) {
                                case "foo bar":
                                    v5 = v5
                                case "\n\t":
                                    if true {
                                        if true {
                                            fmt.Println(v5, v5, v4)
                                            switch fi17 {
                                            case -9223372036854775808:
                                                for fi65 := 0; fi65 < 11; fi65++ {
                                                    v4 = (^(v4))
                                                    v66 := GenS7{F0: 1}
                                                    _ = fi65
                                                    _ = v66
                                                }
                                            }
                                        }
                                        for fi67 := 0; fi67 < 1; fi67++ {
                                            var v68 int64 = -170
                                            _ = fi67
                                            _ = v68
                                        }
                                    }
                                case "":
                                    if (380 <= 42) {
                                        for fi69 := 0; fi69 < 20; fi69++ {
                                            if false {
                                                if len(v6) > 0 {
                                                    v6[0] = (v5 + v5)
                                                }
                                                v70 := GenS8{F0: 2, F1: false, F2: 800}
                                                if true {
                                                    if true {
                                                        v71 := GenS6{F0: 283}
                                                        var v72 string = (("\n\t" + ("a" + v5)) + v5)
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    v5 = "foo bar"
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v73 uint64 = v1
                                                        _ = v73
                                                    }
                                                }
                                                if (2147483647 < -435) {
                                                    var v74 uint32 = 2147483647
                                                    fmt.Println(fi69, fi69)
                                                    _ = len(v6)
                                                    _ = v74
                                                }
                                                fmt.Println(fi69, v4, v1)
                                                v5 = v5
                                                _ = v70
                                            }
                                            v75 := GenS9{F0: 499, F1: v5, F2: 104}
                                            _ = fi69
                                            _ = v75
                                        }
                                        v76 := GenS3{F0: 1, F1: 0, F2: 42}
                                        _ = v76
                                    } else {
                                        v4 = (v4 - 42)
                                        v3 = (^(v3))
                                        fmt.Println(v1)
                                    }
                                }
                            }
                        default:
                            v77 := GenS8{F0: 255, F1: false, F2: 32767}
                            _ = v77
                        }
                        _ = fi17
                    }
                    for fi78 := 0; fi78 < 15; fi78++ {
                        fmt.Println(fi78, v3, v5)
                        _ = fi78
                    }
                    v4 -= -32768
                } else {
                    if false {
                        for fi79 := 0; fi79 < 8; fi79++ {
                            fmt.Println(v1, v3)
                            _ = fi79
                        }
                    } else {
                        _ = len(v2)
                        if (-343 == 9223372036854775807) {
                            v80 := GenS3{F0: 4294967295, F1: 42, F2: 212}
                            _ = v80
                        }
                        if false {
                            var v81 float64 = (((0.1 * 1.5) - float64((float32((1e10 * 3.14159)) / ((0.1 + 1.5) / 1e10)))) + (-2.25 - float64(float32(float64(0.1)))))
                            var v82 int32 = -2147483648
                            _ = v81
                            _ = v82
                        }
                    }
                    v4 = (v4 >> ((255) & 63))
                }
            }
        }
    case -550:
        if (42 < -2147483648) {
            var v83 uint16 = 1
            if true {
                fmt.Println(v3)
                fmt.Println(v1, v5)
            } else {
                var v84 uint = 903
                _ = v84
            }
            v3 = 0
            _ = v83
        }
    case -1:
        switch -871 {
        case -2147483648:
            if (!false) {
                v3 = 2
                v4 = v4
                var v85 uint64 = 446
                _ = v85
            } else {
                for fi86 := 0; fi86 < 10; fi86++ {
                    v87 := GenS1{F0: ((!false) && ((false || (!true)) || true)), F1: v5}
                    v4 = v4
                    _ = fi86
                    _ = v87
                }
                v88 := GenS6{F0: 9223372036854775807}
                for fi89 := 0; fi89 < 19; fi89++ {
                    var v90 int = (fi89 ^ fi89)
                    v91 := GenS6{F0: 1}
                    var v92 int64 = 9223372036854775807
                    _ = fi89
                    _ = v90
                    _ = v91
                    _ = v92
                }
                _ = v88
            }
        case -977:
            v93 := GenS10{F0: -2147483648, F1: 9223372036854775807}
            _ = v93
        }
    }
    if (!false) {
        v94 := GenS10{F0: 1, F1: 32768}
        v5 = (v5 + "")
        fmt.Println(v4, v3, v1)
        _ = v94
    }
    var v95 int32 = 1
    for fi96 := 0; fi96 < 5; fi96++ {
        v97 := GenS11{F0: v95, F1: 2, F2: 65535}
        var v98 string = (("" + (v5 + v5)) + v5)
        var v99 int8 = -128
        _ = fi96
        _ = v97
        _ = v98
        _ = v99
    }
    for fi100 := 0; fi100 < 9; fi100++ {
        if false {
            v101 := GenS12{F0: fi100, F1: (v95 + v95), F2: v5}
            var v102 string = ("\n\t" + ("" + v5))
            if (!false) {
                fmt.Println(v1)
            } else {
                var v103 float64 = -2.25
                var v104 bool = false
                for fi105 := 0; fi105 < 8; fi105++ {
                    v106 := GenS5{F0: 0, F1: 2147483648, F2: v95}
                    _ = fi105
                    _ = v106
                }
                _ = v103
                _ = v104
            }
            _ = v101
            _ = v102
        } else {
            var v107 uint64 = (v3 << ((643) & 63))
            v3 -= 1
            var v108 float64 = (-2.25 + 0.1)
            _ = v107
            _ = v108
        }
        if true {
            for fi109 := 0; fi109 < 6; fi109++ {
                var v110 uint8 = 1
                fmt.Println(fi109)
                _ = fi109
                _ = v110
            }
            var v111 string = v5
            _ = v111
        }
        delete(v2, "\n\t")
        _ = fi100
    }
    v112 := GenS6{F0: 1}
    v5 = v5
    _ = len(v6)
    var v113 string = (v5 + v5)
    v4 = v4
    for fi114 := 0; fi114 < 4; fi114++ {
        for fi115 := 0; fi115 < 17; fi115++ {
            v116 := GenS13{F0: fi115, F1: ((((v113 + "a") + ("foo bar" + (v5 + v5))) + "\n\t") + v113), F2: 42}
            v3 += (v3 << ((255) & 63))
            _ = fi115
            _ = v116
        }
        _ = fi114
    }
    v117 := GenS6{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v112
    _ = v113
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v95
}
