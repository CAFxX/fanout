package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 int16; F1 int32; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint16; F1 int32; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64; F1 int; F2 uint8 }
type GenS11 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS12 struct { F0 uint8; F1 int }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int; F1 uint32; F2 uint16 }
type GenS16 struct { F0 bool; F1 uint8 }

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 285
    var v3 uint16 = 42
    var v4 int8 = -1
    v5 := []int{128}
    var v6 int64 = -9223372036854775808
    v7 := GenS1{F0: -2147483648}
    var v8 float32 = ((3.14159 + 1e10) * 1e10)
    for fi9 := 0; fi9 < 10; fi9++ {
        if (-700 == 1) {
            var v10 float64 = 0.1
            _ = v10
        } else {
            v11 := GenS1{F0: (fi9 | 127)}
            v12 := GenS2{F0: 127}
            if true {
                var v13 int64 = v6
                _ = len(v5)
                v14 := GenS2{F0: (v1 + v1)}
                _ = v13
                _ = v14
            } else {
                if ("a" > ("foo bar" + (("foo bar" + "\n\t") + "a"))) {
                    v15 := GenS3{F0: ((("\n\t" + "") + (("foo bar" + "") + ("hello" + "\n\t"))) + ""), F1: 255}
                    _ = v15
                }
                v16 := GenS3{F0: "a", F1: uint(v1)}
                for fi17 := 0; fi17 < 16; fi17++ {
                    if (false && (!true)) {
                        v18 := GenS1{F0: int(v3)}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v18
                    } else {
                        var v19 uint8 = v1
                        _ = v19
                    }
                    _ = fi17
                }
                _ = v16
            }
            _ = v11
            _ = v12
        }
        for fi20 := 0; fi20 < 4; fi20++ {
            if (898 != 255) {
                switch fi9 {
                case 899:
                    var v21 uint32 = 308
                    _ = v21
                case -1:
                    if false {
                        if (2147483647 > -2147483648) {
                            var v22 int32 = 0
                            _ = v22
                        }
                        if false {
                            var v23 int = (fi9 ^ -2147483648)
                            fmt.Println(fi9)
                            _ = v23
                        }
                        v8 = -2.25
                    } else {
                        var v24 float32 = (v8 - v8)
                        if ((0.0 > ((3.14159 / v8) / 1.0)) || false) {
                            var v25 string = "\n\t"
                            if true {
                                _ = len(v5)
                                v5 = append(v5, 32768)
                                v5 = append(v5, fi9)
                                var v26 uint16 = 2
                                v27 := GenS1{F0: (^(fi20))}
                                _ = v26
                                _ = v27
                            }
                            v1 = 42
                            _ = v25
                        }
                        var v28 string = "a"
                        _ = v24
                        _ = v28
                    }
                case -85:
                    v6 += v6
                }
            }
            if false {
                var v29 uint64 = v2
                var v30 string = ("hello" + "")
                var v31 int = (fi20 ^ fi9)
                _ = v29
                _ = v30
                _ = v31
            } else {
                v32 := GenS2{F0: (^(v1))}
                if ((true && (!true)) || (true && (2147483647 <= -2147483648))) {
                    var v33 uint = 18446744073709551615
                    if (false || true) {
                        v34 := GenS1{F0: 947}
                        v8 -= (float32((0.1 - 0.0)) * (v8 + v8))
                        switch ("foo bar" + ("foo bar" + "foo bar")) {
                        case "\n\t":
                            _ = len(v5)
                        case "a":
                            switch (fi20 >> (((v33 + v33)) & 63)) {
                            case 128:
                                switch fi9 {
                                case -9223372036854775808:
                                    if (-2.25 == ((((float64(v8) + 0.1) - 3.14159) - 1.5) + ((1.5 / 1e10) + (0.0 * 3.14159)))) {
                                        v2 -= v2
                                    } else {
                                        if ((3.14159 - float32(0.0)) <= v8) {
                                            var v35 int64 = v6
                                            if len(v5) > 0 {
                                                v5[0] = -118
                                            }
                                            v36 := GenS4{F0: v2, F1: 725}
                                            _ = v35
                                            _ = v36
                                        }
                                        var v37 float64 = float64(v8)
                                        if (42 <= 0) {
                                            v38 := GenS5{F0: -32768, F1: 447, F2: -1}
                                            if len(v5) > 0 {
                                                v5[0] = 942
                                            }
                                            fmt.Println(v37, v1, v37)
                                            _ = v38
                                        } else {
                                            v39 := GenS5{F0: 42, F1: 1, F2: (v6 & v6)}
                                            var v40 int64 = v6
                                            _ = v39
                                            _ = v40
                                        }
                                        if (0.1 > v8) {
                                            switch "\n\t" {
                                            case "a":
                                                v6 = (v6 + v6)
                                            case "\n\t":
                                                v41 := GenS6{F0: uint32(v2)}
                                                _ = v41
                                            default:
                                                var v42 uint16 = (v3 | v3)
                                                _ = v42
                                            }
                                        }
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v37
                                    }
                                case 128:
                                    var v43 uint8 = (v1 % (((^(v1))) | 1))
                                    _ = v43
                                case 0:
                                    var v44 uint32 = 2147483648
                                    _ = v44
                                default:
                                    if (!(true || (v2 > (v2 % ((1) | 1))))) {
                                        fmt.Println(v1, v8)
                                        var v45 uint64 = v2
                                        var v46 bool = true
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v47 float64 = float64(3.14159)
                                        fmt.Println(v2, v1)
                                        _ = v47
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = fi20
                            }
                        default:
                            fmt.Println(v4)
                        }
                        if false {
                            switch (("" + "\n\t") + "") {
                            case "a":
                                fmt.Println(v1)
                            case "\n\t":
                                var v48 float32 = 1.5
                                _ = v48
                            case "foo bar":
                                var v49 bool = (false && ((("\n\t" >= "\n\t") && false) && true))
                                _ = v49
                            }
                            fmt.Println(v3, v33)
                            var v50 int64 = 0
                            _ = v50
                        } else {
                            if (false && false) {
                                fmt.Println(v33, v3)
                                v8 += (v8 - v8)
                                fmt.Println(v4)
                            }
                        }
                        _ = v34
                    }
                    switch fi9 {
                    case -578:
                        v5 = append(v5, int(v33))
                    case 128:
                        v51 := GenS4{F0: (v2 & 42), F1: 740}
                        _ = v51
                    }
                    _ = v33
                } else {
                    var v52 uint32 = 0
                    v8 = (v8 - 1e10)
                    var v53 uint8 = (v1 << ((4294967295) & 63))
                    _ = v52
                    _ = v53
                }
                v54 := GenS2{F0: (v1 ^ v1)}
                var v55 int8 = (v4 + v4)
                switch (fi9 >> ((1) & 63)) {
                case 42:
                    v56 := GenS5{F0: 446, F1: 447, F2: v6}
                    _ = v56
                default:
                    var v57 int = (fi20 % ((4294967295) | 1))
                    _ = v57
                }
                var v58 uint64 = 18446744073709551615
                var v59 uint64 = v58
                switch int(v3) {
                case 416:
                    if false {
                        switch fi9 {
                        case 127:
                            switch "hello" {
                            case "foo bar":
                                if len(v5) > 0 {
                                    v5[0] = (fi9 % ((fi20) | 1))
                                }
                            case "":
                                _ = len(v5)
                            }
                        case 0:
                            if len(v5) > 0 {
                                v5[0] = fi9
                            }
                        case -984:
                            if true {
                                v60 := GenS5{F0: 1, F1: 42, F2: v6}
                                v61 := GenS7{F0: v59, F1: ""}
                                _ = v60
                                _ = v61
                            }
                        }
                        v2 = (^(v58))
                        fmt.Println(v2, v3, v4)
                        fmt.Println(fi9, v2)
                        var v62 float32 = v8
                        _ = v62
                    }
                case 2147483647:
                    var v63 uint8 = v1
                    _ = v63
                default:
                    v4 = v4
                }
                if false {
                    switch "\n\t" {
                    case "hello":
                        v5 = append(v5, (fi20 - fi9))
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    switch 98 {
                    case 9223372036854775807:
                        if len(v5) > 0 {
                            v5[0] = (fi9 >> ((2) & 63))
                        }
                    default:
                        switch 9223372036854775807 {
                        case -141:
                            v5 = append(v5, 536)
                        case -198:
                            _ = len(v5)
                        case 9223372036854775807:
                            v4 = -128
                        }
                    }
                    if ((!(1 != 1)) && false) {
                        v1 = (v1 - v1)
                        fmt.Println(fi9, v2)
                        v64 := GenS8{F0: 44, F1: -1, F2: 32767}
                        v65 := GenS9{F0: (v6 / ((v6) | 1))}
                        switch "hello" {
                        case "hello":
                            v55 -= v4
                        default:
                            var v66 int = (^(fi20))
                            _ = v66
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v58 = (v59 << ((540) & 63))
                        _ = v64
                        _ = v65
                    }
                }
                _ = v32
                _ = v54
                _ = v55
                _ = v58
                _ = v59
            }
            var v67 float32 = 1.5
            _ = fi20
            _ = v67
        }
        if false {
            if true {
                switch "hello" {
                case "":
                    if (120 <= 2147483648) {
                        v8 = v8
                    }
                case "foo bar":
                    fmt.Println(v3)
                case "hello":
                    v68 := GenS10{F0: v2, F1: fi9, F2: (v1 % (((^(v1))) | 1))}
                    _ = v68
                }
                var v69 string = ""
                v70 := GenS2{F0: 128}
                _ = v69
                _ = v70
            } else {
                if false {
                    switch fi9 {
                    case -9223372036854775808:
                        v71 := GenS11{F0: 42, F1: (v2 >> ((255) & 63)), F2: v3}
                        _ = v71
                    case -1:
                        v72 := GenS12{F0: (v1 / ((v1) | 1)), F1: 42}
                        _ = v72
                    case 42:
                        for fi73 := 0; fi73 < 14; fi73++ {
                            var v74 float32 = float32(1e10)
                            if (false || false) {
                                fmt.Println(v4, v74)
                            }
                            if (false && true) {
                                var v75 int64 = (v6 / (((-(v6))) | 1))
                                _ = v75
                            }
                            _ = fi73
                            _ = v74
                        }
                    default:
                        var v76 bool = false
                        _ = v76
                    }
                    v5 = append(v5, (fi9 * (fi9 % ((-548) | 1))))
                    for fi77 := 0; fi77 < 19; fi77++ {
                        fmt.Println(fi77, fi77)
                        if false {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        _ = fi77
                    }
                } else {
                    for fi78 := 0; fi78 < 2; fi78++ {
                        _ = len(v5)
                        _ = len(v5)
                        var v79 uint8 = v1
                        _ = fi78
                        _ = v79
                    }
                    if (v2 == (^(v2))) {
                        _ = len(v5)
                        switch (fi9 + -1) {
                        case 127:
                            var v80 uint64 = (v2 % ((v2) | 1))
                            _ = v80
                        case -2147483648:
                            fmt.Println(v3, fi9)
                        default:
                            v8 -= v8
                        }
                    }
                }
                v5 = append(v5, 0)
                v1 += (^(v1))
            }
        } else {
            if false {
                if (0 > (-(fi9))) {
                    if (1 != 1) {
                        for fi81 := 0; fi81 < 16; fi81++ {
                            var v82 int32 = 2147483647
                            _ = fi81
                            _ = v82
                        }
                        v83 := GenS5{F0: 32767, F1: -1, F2: (v6 << ((255) & 63))}
                        _ = v83
                    }
                    if (true && (v2 == (v2 & v2))) {
                        if false {
                            v84 := GenS13{F0: 73}
                            switch ("hello" + "foo bar") {
                            case "":
                                fmt.Println(v8)
                            case "foo bar":
                                v5 = append(v5, fi9)
                            case "hello":
                                var v85 uint64 = (v2 | (v2 / (((v2 | (v2 ^ v2))) | 1)))
                                _ = v85
                            }
                            _ = v84
                        }
                        for fi86 := 0; fi86 < 2; fi86++ {
                            _ = len(v5)
                            v4 -= 127
                            _ = fi86
                        }
                    }
                } else {
                    v2 += 951
                    v4 += (v4 & (v4 - v4))
                }
            }
            var v87 float32 = v8
            _ = v87
        }
        _ = fi9
    }
    v5 = append(v5, 1)
    for fi88 := 0; fi88 < 5; fi88++ {
        switch "foo bar" {
        case "\n\t":
            for fi89 := 0; fi89 < 1; fi89++ {
                v6 -= (v6 >> ((42) & 63))
                _ = fi89
            }
        case "a":
            v6 = v6
        case "hello":
            v4 -= (v4 >> ((18446744073709551615) & 63))
        default:
            for fi90 := 0; fi90 < 14; fi90++ {
                switch "a" {
                case "foo bar":
                    switch (fi88 + (-(fi88))) {
                    case 32768:
                        var v91 int64 = -643
                        _ = v91
                    case 127:
                        var v92 string = "a"
                        _ = v92
                    case -55:
                        switch ("hello" + "a") {
                        case "hello":
                            var v93 string = "\n\t"
                            _ = v93
                        case "foo bar":
                            var v94 int16 = -32768
                            _ = v94
                        }
                    }
                case "":
                    v6 += (v6 << ((547) & 63))
                case "hello":
                    _ = len(v5)
                }
                v1 = (v1 + (^(v1)))
                var v95 float32 = float32(1.5)
                _ = fi90
                _ = v95
            }
        }
        _ = fi88
    }
    fmt.Println(v8, v6)
    v1 -= (v1 % ((128) | 1))
    for fi96 := 0; fi96 < 2; fi96++ {
        switch int(v6) {
        case 0:
            if true {
                var v97 uint32 = 0
                _ = v97
            }
        case -1:
            for fi98 := 0; fi98 < 2; fi98++ {
                fmt.Println(v3)
                if ("a" < "a") {
                    var v99 int8 = (v4 % ((v4) | 1))
                    _ = v99
                } else {
                    v3 = 42
                    var v100 int16 = -32768
                    var v101 bool = false
                    var v102 uint16 = 0
                    _ = v100
                    _ = v101
                    _ = v102
                }
                v1 -= 42
                _ = fi98
            }
        case 127:
            if false {
                if (!((!(int16(fi96) != -32768)) || false)) {
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v8, v8, v3)
                    case "a":
                        v3 += 255
                    case "foo bar":
                        v103 := GenS8{F0: (v3 % ((uint16(v1)) | 1)), F1: 42, F2: 455}
                        _ = v103
                    default:
                        var v104 uint16 = (^(v3))
                        _ = v104
                    }
                    fmt.Println(v8)
                    fmt.Println(v1, v2)
                } else {
                    switch (fi96 % ((fi96) | 1)) {
                    case -137:
                        if (float64(((((((v8 * v8) * 0.1) / v8) / (v8 - (float32(3.14159) + 0.1))) + v8) + 0.1)) < float64(v8)) {
                            v8 += v8
                            v5 = append(v5, 2147483647)
                        } else {
                            switch "" {
                            case "foo bar":
                                var v105 uint64 = (v2 << ((427) & 63))
                                _ = v105
                            }
                        }
                    case 0:
                        v5 = append(v5, fi96)
                    case -1:
                        if false {
                            switch (fi96 / (((fi96 * fi96)) | 1)) {
                            case 4294967295:
                                var v106 float32 = v8
                                _ = v106
                            case -697:
                                v107 := GenS14{F0: 1}
                                _ = v107
                            case 780:
                                var v108 int64 = v6
                                _ = v108
                            }
                            switch 824 {
                            case 2147483647:
                                v3 = v3
                            default:
                                v109 := GenS15{F0: fi96, F1: 99, F2: 1}
                                _ = v109
                            }
                            switch "a" {
                            case "foo bar":
                                v1 = (v1 % (((v1 + (v1 * 191))) | 1))
                            case "a":
                                var v110 uint16 = v3
                                _ = v110
                            case "\n\t":
                                for fi111 := 0; fi111 < 14; fi111++ {
                                    var v112 string = "\n\t"
                                    if ((!false) || false) {
                                        var v113 int32 = -1
                                        if len(v5) > 0 {
                                            v5[0] = (fi96 - int(v6))
                                        }
                                        var v114 float64 = float64(v8)
                                        v6 = v6
                                        v8 -= v8
                                        _ = v113
                                        _ = v114
                                    } else {
                                        v115 := GenS6{F0: 0}
                                        var v116 uint64 = v2
                                        _ = v115
                                        _ = v116
                                    }
                                    _ = fi111
                                    _ = v112
                                }
                            default:
                                v117 := GenS16{F0: true, F1: (v1 - uint8(v3))}
                                _ = v117
                            }
                        }
                    }
                    var v118 int16 = 1000
                    switch 4294967295 {
                    case 67:
                        switch "a" {
                        case "a":
                            if len(v5) > 0 {
                                v5[0] = (fi96 >> ((1) & 63))
                            }
                        case "hello":
                            v8 -= (v8 + (0.0 * 1e10))
                        default:
                            v4 = v4
                        }
                    case 511:
                        switch fi96 {
                        case 641:
                            v8 += v8
                        case -2147483648:
                            fmt.Println(v3)
                        case 2147483647:
                            fmt.Println(v3, v2, v6)
                        }
                    case -2147483648:
                        var v119 int64 = -856
                        _ = v119
                    }
                    _ = v118
                }
            } else {
                var v120 float32 = v8
                if false {
                    var v121 int16 = -1
                    if ("\n\t" == "a") {
                        v122 := GenS13{F0: (-(v4))}
                        _ = v122
                    }
                    v3 = v3
                    _ = v121
                }
                _ = v120
            }
        default:
            var v123 int8 = v4
            _ = v123
        }
        _ = fi96
    }
    var v124 int32 = int32(v6)
    var v125 int32 = (v124 << ((977) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v124:", v124)
    fmt.Println("v125:", v125)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
