package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 int16; F2 uint64 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint32; F2 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 uint64 }
type GenS10 struct { F0 int32; F1 uint64 }
type GenS11 struct { F0 uint32; F1 uint64 }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint64; F1 int32 }
type GenS15 struct { F0 int8; F1 int; F2 int }
type GenS16 struct { F0 int64; F1 bool; F2 int }
type GenS17 struct { F0 uint32 }

func main() {
    var v1 uint = 81
    v2 := []int{-2147483648, 0, 42, 9223372036854775807}
    v3 := []int{86, 2147483647, 32767}
    v4 := make(map[int]int)
    if false {
        v1 += (v1 << ((42) & 63))
        switch 0 {
        case 42:
            for fi5 := 0; fi5 < 11; fi5++ {
                fmt.Println(v1)
                v1 = (v1 / ((2) | 1))
                fmt.Println(fi5, fi5)
                _ = fi5
            }
        case -1:
            v1 += 2
        case 9223372036854775807:
            if true {
                fmt.Println(v1)
                var v6 int8 = 120
                _ = v6
            } else {
                var v7 int8 = 1
                if true {
                    v7 = v7
                    if true {
                        if (false || (720 < 908)) {
                            fmt.Println(v7, v7)
                            var v8 float32 = 0.1
                            _ = v8
                        } else {
                            fmt.Println(v7)
                            if (-352 >= -987) {
                                for fi9 := 0; fi9 < 1; fi9++ {
                                    v2 = append(v2, (fi9 / ((fi9) | 1)))
                                    v7 -= v7
                                    _ = fi9
                                }
                                delete(v4, -9223372036854775808)
                                v1 = (v1 ^ (v1 >> ((v1) & 63)))
                            } else {
                                v1 += v1
                            }
                            switch 42 {
                            case 127:
                                fmt.Println(v1)
                            case -723:
                                fmt.Println(v1)
                            case -97:
                                v7 = (^(v7))
                            }
                        }
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                } else {
                    v10 := GenS1{F0: 0, F1: 1, F2: ("foo bar" + "hello")}
                    switch -55 {
                    case -2147483648:
                        v11 := GenS1{F0: 255, F1: 0, F2: ("" + (("a" + (("hello" + "foo bar") + "hello")) + "a"))}
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 9; fi12++ {
                            if false {
                                if true {
                                    switch (fi12 / ((fi12) | 1)) {
                                    case 154:
                                        fmt.Println(v1, fi12, v7)
                                    default:
                                        v7 += -1
                                    }
                                    if false {
                                        if (2147483647 >= 65535) {
                                            if true {
                                                v7 = -1
                                                if true {
                                                    v7 -= (v7 / ((v7) | 1))
                                                } else {
                                                    v1 = 2
                                                }
                                                if false {
                                                    v7 = (v7 << ((v1) & 63))
                                                } else {
                                                    v7 += v7
                                                    _ = len(v4)
                                                }
                                            } else {
                                                if (v1 >= v1) {
                                                    v13 := GenS1{F0: 65535, F1: 104, F2: "hello"}
                                                    if true {
                                                        switch ("" + "") {
                                                        case "\n\t":
                                                            if len(v3) > 0 {
                                                                v3[0] = (fi12 << ((v1) & 63))
                                                            }
                                                        case "":
                                                            v1 -= (v1 << ((v1) & 63))
                                                        case "foo bar":
                                                            v14 := GenS1{F0: 122, F1: 0, F2: ("\n\t" + "")}
                                                            _ = v14
                                                        default:
                                                            var v15 float32 = float32(1e10)
                                                            _ = v15
                                                        }
                                                    }
                                                    if true {
                                                        fmt.Println(fi12, fi12)
                                                        v7 += (^(v7))
                                                        for fi16 := 0; fi16 < 9; fi16++ {
                                                            v17 := GenS1{F0: 255, F1: 0, F2: ("hello" + ("hello" + "a"))}
                                                            v18 := GenS2{F0: 42}
                                                            _ = fi16
                                                            _ = v17
                                                            _ = v18
                                                        }
                                                    }
                                                    _ = v13
                                                }
                                                for fi19 := 0; fi19 < 12; fi19++ {
                                                    fmt.Println(v7, fi12)
                                                    _ = fi19
                                                }
                                            }
                                            for fi20 := 0; fi20 < 13; fi20++ {
                                                switch "hello" {
                                                case "\n\t":
                                                    if false {
                                                        var v21 int8 = (v7 << (((^(v1))) & 63))
                                                        var v22 uint64 = 4294967295
                                                        if false {
                                                            v23 := GenS2{F0: 0}
                                                            _ = v23
                                                        }
                                                        fmt.Println(v21)
                                                        fmt.Println(fi12)
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                }
                                                _ = fi20
                                            }
                                        }
                                        if (!false) {
                                            fmt.Println(v7, v7, v7)
                                            var v24 int8 = v7
                                            _ = v24
                                        } else {
                                            var v25 int8 = -1
                                            switch ("foo bar" + ("hello" + "foo bar")) {
                                            case "a":
                                                switch "\n\t" {
                                                case "":
                                                    if (!(808 < 65535)) {
                                                        if (519 < 65535) {
                                                            var v26 float32 = (3.14159 * 3.14159)
                                                            for fi27 := 0; fi27 < 3; fi27++ {
                                                                if false {
                                                                    var v28 bool = true
                                                                    _ = v28
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v1 += v1
                                                                    case "":
                                                                        var v29 float64 = (3.14159 * 0.0)
                                                                        _ = v29
                                                                    default:
                                                                        if (0 == 503) {
                                                                            var v30 float64 = 1.5
                                                                            v31 := GenS3{F0: "foo bar", F1: 42, F2: 0}
                                                                            v1 += (v1 % ((v1) | 1))
                                                                            v1 = (v1 / (((v1 ^ (v1 / ((v1) | 1)))) | 1))
                                                                            if (255 > uint8(v1)) {
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = fi12
                                                                                }
                                                                                fmt.Println(v1)
                                                                                var v32 uint8 = 131
                                                                                _ = v32
                                                                            }
                                                                            v25 += (-(v25))
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                        fmt.Println(v26, v7)
                                                                        switch (^(fi27)) {
                                                                        case 4294967295:
                                                                            if (0.0 != 0.0) {
                                                                                fmt.Println(v25)
                                                                            } else {
                                                                                var v33 int32 = -1
                                                                                if false {
                                                                                    fmt.Println(v26, v1, fi12)
                                                                                }
                                                                                _ = v33
                                                                            }
                                                                            var v34 int16 = 32767
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v34
                                                                        case 640:
                                                                            var v35 uint32 = 0
                                                                            _ = v35
                                                                        case 42:
                                                                            v26 = v26
                                                                        }
                                                                    }
                                                                    switch ("\n\t" + "\n\t") {
                                                                    case "a":
                                                                        v36 := GenS4{F0: -512, F1: 1}
                                                                        _ = v36
                                                                    case "foo bar":
                                                                        var v37 float32 = (float32(3.14159) - (1e10 / v26))
                                                                        _ = v37
                                                                    case "":
                                                                        fmt.Println(v1)
                                                                    default:
                                                                        fmt.Println(v1, v25)
                                                                    }
                                                                }
                                                                var v38 int8 = (^(v7))
                                                                var v39 bool = true
                                                                _ = fi27
                                                                _ = v38
                                                                _ = v39
                                                            }
                                                            for fi40 := 0; fi40 < 7; fi40++ {
                                                                if true {
                                                                    fmt.Println(v7, v25)
                                                                }
                                                                _ = fi40
                                                            }
                                                            _ = v26
                                                        } else {
                                                            if ((0 >= 959) && (18446744073709551615 <= 513)) {
                                                                fmt.Println(fi12, v7, v7)
                                                            }
                                                            if (331 == 42) {
                                                                var v41 uint32 = 4294967295
                                                                if false {
                                                                    switch ((((("a" + "a") + ("foo bar" + ("" + "foo bar"))) + ("" + "\n\t")) + "\n\t") + "a") {
                                                                    case "":
                                                                        var v42 int32 = 2147483647
                                                                        _ = v42
                                                                    case "\n\t":
                                                                        delete(v4, -952)
                                                                    case "foo bar":
                                                                        v7 += (v25 & v7)
                                                                    default:
                                                                        v7 = (v7 << ((211) & 63))
                                                                    }
                                                                    if (((v41 | v41) != v41) || false) {
                                                                        if (false || false) {
                                                                            var v43 uint = (v1 * (v1 / ((255) | 1)))
                                                                            _ = v43
                                                                        }
                                                                        v7 = (v7 % ((v7) | 1))
                                                                        switch ("hello" + "hello") {
                                                                        case "hello":
                                                                            if ((("hello" + "\n\t") + (("hello" + ("a" + "\n\t")) + "\n\t")) < ("a" + "hello")) {
                                                                                fmt.Println(fi12, v41, v7)
                                                                                v7 = int8(fi12)
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                            } else {
                                                                                _ = len(v2)
                                                                                v1 = (v1 << (((v1 >> (((v1 >> (((^(v1))) & 63))) & 63))) & 63))
                                                                                v7 = -1
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi12, v41)
                                                                    }
                                                                    v25 -= (v7 << ((uint(v41)) & 63))
                                                                } else {
                                                                    var v44 uint = v1
                                                                    _ = v44
                                                                }
                                                                v45 := GenS5{F0: 4294967295}
                                                                _ = v41
                                                                _ = v45
                                                            }
                                                        }
                                                        for fi46 := 0; fi46 < 1; fi46++ {
                                                            _ = len(v4)
                                                            _ = fi46
                                                        }
                                                    } else {
                                                        for fi47 := 0; fi47 < 7; fi47++ {
                                                            if false {
                                                                v48 := GenS6{F0: 42, F1: 65535, F2: (0 > (v1 << ((v1) & 63)))}
                                                                v49 := GenS7{F0: 144}
                                                                switch (("foo bar" + ("foo bar" + "a")) + "hello") {
                                                                case "hello":
                                                                    v7 = 127
                                                                case "\n\t":
                                                                    v25 -= (v25 & 127)
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                            } else {
                                                                if (0 <= 903) {
                                                                    var v50 uint = v1
                                                                    if true {
                                                                        v25 -= v25
                                                                        v51 := v4[(^(fi12))]
                                                                        if ((685 >= -287) || ((v1 ^ (v1 | (v50 + (v50 % ((810) | 1))))) > 42)) {
                                                                            v50 = 255
                                                                            fmt.Println(v51, fi12)
                                                                        }
                                                                        _ = v51
                                                                    }
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        switch fi12 {
                                                                        case 4294967295:
                                                                            v7 += (v7 ^ v7)
                                                                        default:
                                                                            v25 = -23
                                                                        }
                                                                    case "hello":
                                                                        v2 = append(v2, fi12)
                                                                    default:
                                                                        var v52 int = fi12
                                                                        _ = v52
                                                                    }
                                                                    _ = v50
                                                                }
                                                                v7 -= (v7 * (v25 >> ((v1) & 63)))
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            if false {
                                                                if ((!true) && false) {
                                                                    var v53 uint = v1
                                                                    var v54 uint32 = 65535
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                if (0 == 213) {
                                                                    v1 = 18446744073709551615
                                                                    v25 = 42
                                                                    if true {
                                                                        v2 = append(v2, (fi12 >> ((2) & 63)))
                                                                    }
                                                                    var v55 int = (fi12 ^ (-(fi47)))
                                                                    _ = v55
                                                                } else {
                                                                    v56 := GenS2{F0: -9223372036854775808}
                                                                    if false {
                                                                        var v57 uint16 = 0
                                                                        switch (fi47 + int(v57)) {
                                                                        case -240:
                                                                            v58 := GenS8{F0: -2147483648}
                                                                            var v59 string = ("hello" + "")
                                                                            _ = v58
                                                                            _ = v59
                                                                        case 32767:
                                                                            switch -1 {
                                                                            case 9223372036854775807:
                                                                                v7 = int8(fi47)
                                                                            case 250:
                                                                                fmt.Println(v7)
                                                                            case 1:
                                                                                v60 := GenS5{F0: 2}
                                                                                _ = v60
                                                                            }
                                                                        default:
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                var v61 float64 = 1.5
                                                                                _ = v61
                                                                            case "foo bar":
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    switch 32767 {
                                                                                    case 32768:
                                                                                        v62 := GenS1{F0: 917, F1: int64(v1), F2: "\n\t"}
                                                                                        _ = v62
                                                                                    case 127:
                                                                                        v63 := GenS6{F0: 2, F1: 1, F2: (0.1 >= (3.14159 - 0.0))}
                                                                                        _ = v63
                                                                                    case -867:
                                                                                        v1 = v1
                                                                                    default:
                                                                                        v64 := GenS2{F0: 9223372036854775807}
                                                                                        _ = v64
                                                                                    }
                                                                                case "foo bar":
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = 463
                                                                                    }
                                                                                default:
                                                                                    v65 := GenS3{F0: "foo bar", F1: 32767, F2: 2}
                                                                                    _ = v65
                                                                                }
                                                                            default:
                                                                                switch ((((("" + "") + "hello") + ((("\n\t" + "hello") + "hello") + "\n\t")) + ("foo bar" + "\n\t")) + "a") {
                                                                                case "foo bar":
                                                                                    fmt.Println(fi47, v7, fi12)
                                                                                case "\n\t":
                                                                                    v25 = int8(fi47)
                                                                                case "":
                                                                                    v7 -= -90
                                                                                }
                                                                            }
                                                                        }
                                                                        if (1 < 9223372036854775807) {
                                                                            v66 := GenS9{F0: (fi47 << ((v1) & 63)), F1: 116}
                                                                            fmt.Println(v57)
                                                                            delete(v4, fi47)
                                                                            _ = v66
                                                                        } else {
                                                                            fmt.Println(v1, v7, v57)
                                                                            v67 := GenS4{F0: -9223372036854775808, F1: 226}
                                                                            _ = v67
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                    if ((-(v25)) >= 42) {
                                                                        v1 = v1
                                                                    }
                                                                    v68 := GenS10{F0: 354, F1: 18446744073709551615}
                                                                    _ = v56
                                                                    _ = v68
                                                                }
                                                            } else {
                                                                fmt.Println(v7, fi47, fi12)
                                                                var v69 uint16 = 77
                                                                _ = v69
                                                            }
                                                            _ = fi47
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        if (true || (!false)) {
                                                            for fi70 := 0; fi70 < 18; fi70++ {
                                                                v7 = v25
                                                                if (841 > 18446744073709551615) {
                                                                    v1 += 1
                                                                    switch 32768 {
                                                                    case 0:
                                                                        v71 := GenS11{F0: 977, F1: 2}
                                                                        _ = v71
                                                                    default:
                                                                        if false {
                                                                            v72 := GenS12{F0: -32768, F1: 1}
                                                                            v73 := GenS13{F0: 0}
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                    }
                                                                } else {
                                                                    v1 += v1
                                                                    switch fi12 {
                                                                    case -615:
                                                                        var v74 int = fi12
                                                                        _ = v74
                                                                    default:
                                                                        if true {
                                                                            var v75 uint32 = 4294967295
                                                                            var v76 int16 = -32768
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 879
                                                                            }
                                                                            _ = v75
                                                                            _ = v76
                                                                        }
                                                                    }
                                                                }
                                                                fmt.Println(v1)
                                                                switch ("a" + "foo bar") {
                                                                case "\n\t":
                                                                    v77 := GenS13{F0: 255}
                                                                    _ = v77
                                                                case "foo bar":
                                                                    if ((3.14159 - 0.0) > 1.5) {
                                                                        v78 := GenS14{F0: 2, F1: 2147483647}
                                                                        _ = v78
                                                                    } else {
                                                                        v25 -= v25
                                                                    }
                                                                }
                                                                _ = fi70
                                                            }
                                                        } else {
                                                            if true {
                                                                for fi79 := 0; fi79 < 5; fi79++ {
                                                                    v25 -= v7
                                                                    var v80 int64 = -321
                                                                    fmt.Println(v80, fi79, v1)
                                                                    _ = fi79
                                                                    _ = v80
                                                                }
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi12
                                                                }
                                                            }
                                                            var v81 string = ""
                                                            _ = v81
                                                        }
                                                    } else {
                                                        v82 := GenS11{F0: 65535, F1: 4294967295}
                                                        _ = v82
                                                    }
                                                }
                                            case "hello":
                                                v83 := GenS15{F0: (-(v7)), F1: (fi12 << (((^(v1))) & 63)), F2: fi12}
                                                _ = v83
                                            default:
                                                v25 += 0
                                            }
                                            _ = v25
                                        }
                                        var v84 uint64 = 474
                                        _ = v84
                                    } else {
                                        var v85 string = "hello"
                                        v86 := GenS16{F0: -9223372036854775808, F1: (-841 > 1000), F2: fi12}
                                        if (65535 <= 1) {
                                            var v87 int32 = 1
                                            fmt.Println(v87)
                                            v87 = (v87 / ((739) | 1))
                                            _ = v87
                                        }
                                        _ = v85
                                        _ = v86
                                    }
                                    fmt.Println(v1)
                                } else {
                                    var v88 uint = (v1 % ((0) | 1))
                                    fmt.Println(v7, fi12, fi12)
                                    v88 = v88
                                    _ = v88
                                }
                                v89 := GenS11{F0: 2, F1: 707}
                                _ = v89
                            } else {
                                v7 -= v7
                                fmt.Println(v1, v1)
                            }
                            _ = fi12
                        }
                    }
                    for fi90 := 0; fi90 < 2; fi90++ {
                        v1 = (v1 | (v1 - (v1 - v1)))
                        _ = fi90
                    }
                    _ = v10
                }
                if (42 >= 0) {
                    v4[556] = 32767
                } else {
                    if ((1 < 1) && (4294967295 < (^(v1)))) {
                        var v91 int8 = (v7 | v7)
                        switch 1 {
                        case 1:
                            for fi92 := 0; fi92 < 14; fi92++ {
                                v7 -= v7
                                _ = fi92
                            }
                        case 492:
                            _ = len(v3)
                        case 32768:
                            fmt.Println(v1, v91, v7)
                        default:
                            switch "" {
                            case "\n\t":
                                for fi93 := 0; fi93 < 16; fi93++ {
                                    _ = len(v4)
                                    for fi94 := 0; fi94 < 15; fi94++ {
                                        fmt.Println(fi93, v7, v1)
                                        v1 = (v1 % (((^(v1))) | 1))
                                        _ = fi94
                                    }
                                    _ = fi93
                                }
                            }
                        }
                        var v95 int32 = 42
                        _ = v91
                        _ = v95
                    } else {
                        if true {
                            fmt.Println(v1)
                            if false {
                                var v96 bool = (int64(v1) <= 0)
                                _ = v96
                            }
                        }
                        v7 = (v7 | (v7 * (v7 + (-(v7)))))
                    }
                    fmt.Println(v7, v1)
                    v97 := GenS17{F0: 2147483647}
                    _ = v97
                }
                _ = v7
            }
        default:
            for fi98 := 0; fi98 < 10; fi98++ {
                v99 := GenS11{F0: 2147483647, F1: 0}
                fmt.Println(v1, v1)
                _ = fi98
                _ = v99
            }
        }
        for fi100 := 0; fi100 < 16; fi100++ {
            var v101 int64 = -1
            fmt.Println(v1, v1)
            var v102 int8 = 0
            _ = fi100
            _ = v101
            _ = v102
        }
    }
    v103 := GenS17{F0: 42}
    v104 := GenS15{F0: 100, F1: 32768, F2: -270}
    v3 = append(v3, 128)
    if (false && true) {
        var v105 uint16 = 42
        var v106 uint64 = 18446744073709551615
        var v107 uint32 = 2147483648
        _ = v105
        _ = v106
        _ = v107
    } else {
        v108 := GenS6{F0: 255, F1: 4294967295, F2: true}
        var v109 float32 = float32((1.5 - 1.5))
        _ = v108
        _ = v109
    }
    v1 -= v1
    for fi110 := 0; fi110 < 5; fi110++ {
        delete(v4, fi110)
        _ = fi110
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "hello":
        if false {
            if true {
                switch "\n\t" {
                case "":
                    var v111 int8 = 85
                    _ = v111
                case "hello":
                    var v112 uint32 = 601
                    _ = v112
                default:
                    v1 += 4294967295
                }
                var v113 int16 = 49
                v2 = append(v2, 32768)
                _ = v113
            } else {
                fmt.Println(v1)
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "hello":
                        v3 = append(v3, 103)
                    }
                case "\n\t":
                    v1 -= (v1 + v1)
                }
                v1 -= v1
            }
            _ = len(v3)
        }
    default:
        v114 := GenS13{F0: 32767}
        _ = v114
    }
    v1 += v1
    switch 505 {
    case 0:
        var v115 uint64 = 838
        _ = v115
    case 128:
        if (!false) {
            var v116 string = "a"
            _ = v116
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v103
    _ = v104
    _ = v2
    _ = v3
    _ = v4
}
