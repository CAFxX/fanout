package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 int8; F1 int16; F2 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int64; F2 string }
type GenS10 struct { F0 int8; F1 uint; F2 int }
type GenS11 struct { F0 int32; F1 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int; F1 uint; F2 int }
type GenS14 struct { F0 int16; F1 int32; F2 uint }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 bool; F1 int16 }
type GenS17 struct { F0 uint32; F1 uint8; F2 bool }
type GenS18 struct { F0 string }
type GenS19 struct { F0 uint; F1 int8; F2 bool }
type GenS20 struct { F0 uint64; F1 uint64 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS23 struct { F0 uint64; F1 int64 }
type GenS24 struct { F0 uint8; F1 bool }
type GenS25 struct { F0 int16; F1 bool; F2 bool }
type GenS26 struct { F0 uint8; F1 int8 }
type GenS27 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS28 struct { F0 int }
type GenS29 struct { F0 bool; F1 uint8; F2 int32 }
type GenS30 struct { F0 string }
type GenS31 struct { F0 int16; F1 uint64 }
type GenS32 struct { F0 string }
type GenS33 struct { F0 int16; F1 uint64 }
type GenS34 struct { F0 int32; F1 int8; F2 uint32 }

func main() {
    v1 := []uint32{179}
    var v2 float64 = 0.1
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 uint64 = 577
    if true {
        var v6 int = 128
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 int = -289
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: -2147483648}
                if (v2 <= -2.25) {
                    var v11 int16 = 435
                    _ = v11
                }
                fmt.Println(fi7, v2, fi9)
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
        var v12 bool = false
        _ = v12
        _ = v6
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch fi13 {
        case 4294967295:
            var v14 int64 = 0
            _ = v14
        case 191:
            v15 := GenS2{F0: 67, F1: ((^(v5)) < (^(v5)))}
            _ = v15
        }
        switch "hello" {
        case "hello":
            for fi16 := 0; fi16 < 2; fi16++ {
                v17 := GenS2{F0: 255, F1: ((1e10 / 0.1) == 3.14159)}
                _ = fi16
                _ = v17
            }
        case "":
            v2 = ((0.0 / float64(float32(v2))) * (v2 - (((3.14159 - 1e10) - v2) * float64(1.5))))
        }
        var v18 float64 = (v2 + v2)
        _ = fi13
        _ = v18
    }
    v2 = float64(-2.25)
    if false {
        v19 := v3[2147483647]
        v20 := GenS2{F0: 255, F1: ((v19 || (-595 < 1)) || (v19 && v19))}
        v19 = ((1e10 * v2) > ((0.1 * (v2 * ((1e10 / v2) - (v2 + v2)))) + v2))
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: -1, F1: 165}
    if (false && false) {
        switch "a" {
        case "":
            for fi22 := 0; fi22 < 12; fi22++ {
                if (42 >= 2) {
                    switch ((("a" + "a") + "\n\t") + (("foo bar" + "") + "\n\t")) {
                    case "\n\t":
                        if false {
                            v23 := GenS3{F0: -177, F1: 1}
                            if true {
                                v24 := GenS4{F0: -1}
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    if (!true) {
                                        v5 -= (v5 % ((v5) | 1))
                                        var v26 string = (("foo bar" + "") + "\n\t")
                                        _ = v26
                                    }
                                    var v27 uint16 = uint16(v5)
                                    if false {
                                        v5 = 42
                                    } else {
                                        fmt.Println(v5, fi22)
                                    }
                                    _ = fi25
                                    _ = v27
                                }
                                _ = v24
                            }
                            _ = v23
                        } else {
                            v28 := v3[(fi22 * (^(fi22)))]
                            _ = v28
                        }
                    case "a":
                        for fi29 := 0; fi29 < 7; fi29++ {
                            switch 128 {
                            case 0:
                                v5 += v5
                            }
                            _ = len(v3)
                            fmt.Println(v2, v5)
                            _ = fi29
                        }
                    case "hello":
                        v30 := GenS5{F0: -1, F1: ""}
                        _ = v30
                    }
                    v2 = 3.14159
                    if false {
                        switch (fi22 - (fi22 / ((fi22) | 1))) {
                        case -205:
                            if (false && true) {
                                if (false && true) {
                                    fmt.Println(fi22, v2, v5)
                                    fmt.Println(v2)
                                    switch "hello" {
                                    case "a":
                                        var v31 string = "hello"
                                        _ = v31
                                    default:
                                        v32 := GenS6{F0: 32, F1: 42, F2: false}
                                        _ = v32
                                    }
                                }
                                var v33 int64 = 104
                                _ = v33
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                var v34 uint64 = v5
                                for fi35 := 0; fi35 < 7; fi35++ {
                                    var v36 float64 = (v2 / ((v2 * 3.14159) * ((v2 + v2) + (v2 / v2))))
                                    switch "\n\t" {
                                    case "hello":
                                        if true {
                                            var v37 int = (fi22 * (fi22 | 127))
                                            v34 += 1
                                            fmt.Println(fi35)
                                            _ = v37
                                        }
                                    }
                                    fmt.Println(v2, v34, fi35)
                                    _ = fi35
                                    _ = v36
                                }
                                _ = v34
                            }
                        case 1:
                            for fi38 := 0; fi38 < 13; fi38++ {
                                var v39 uint16 = 0
                                if (((true && false) || (42 > int16(fi38))) || true) {
                                    switch "\n\t" {
                                    case "hello":
                                        var v40 uint16 = 42
                                        _ = v40
                                    case "":
                                        var v41 float32 = (0.1 - ((1.5 / -2.25) + (float32(v2) * float32((v2 + (v2 * (((v2 / float64(((1e10 - 1.5) - (-2.25 - (3.14159 / 0.1))))) + v2) + ((((v2 / v2) + 3.14159) - v2) / ((v2 * 0.1) - -2.25)))))))))
                                        _ = v41
                                    case "\n\t":
                                        switch "a" {
                                        case "\n\t":
                                            if false {
                                                delete(v3, fi22)
                                                if (true || (true || (0 == 18446744073709551615))) {
                                                    switch 446 {
                                                    case 183:
                                                        if (18446744073709551615 >= v5) {
                                                            var v42 uint16 = v39
                                                            if true {
                                                                fmt.Println(v39, v5)
                                                                var v43 float32 = 1e10
                                                                _ = v43
                                                            }
                                                            _ = v42
                                                        } else {
                                                            v5 += v5
                                                            var v44 uint32 = 2147483648
                                                            _ = v44
                                                        }
                                                    case -71:
                                                        var v45 int16 = 42
                                                        _ = v45
                                                    default:
                                                        v1 = append(v1, 4294967295)
                                                    }
                                                } else {
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v46 uint64 = v5
                                                        _ = v46
                                                    default:
                                                        switch fi22 {
                                                        case -733:
                                                            if false {
                                                                var v47 uint32 = 2
                                                                _ = v47
                                                            } else {
                                                                var v48 uint8 = 42
                                                                v2 = v2
                                                                _ = v48
                                                            }
                                                        case 29:
                                                            v2 -= float64((1e10 * 0.1))
                                                        case 32768:
                                                            var v49 uint8 = 42
                                                            _ = v49
                                                        default:
                                                            fmt.Println(v39)
                                                        }
                                                    }
                                                    var v50 int16 = 42
                                                    var v51 string = "\n\t"
                                                    _ = v50
                                                    _ = v51
                                                }
                                            } else {
                                                v5 += uint64(v39)
                                                var v52 string = "\n\t"
                                                v53 := GenS3{F0: -1, F1: 1}
                                                _ = v52
                                                _ = v53
                                            }
                                        case "a":
                                            var v54 int8 = 42
                                            _ = v54
                                        case "hello":
                                            if true {
                                                v55 := GenS7{F0: 4294967295}
                                                var v56 uint16 = 255
                                                fmt.Println(v56)
                                                if (v5 == (^(v5))) {
                                                    v39 += (v39 * v39)
                                                    v39 = v39
                                                } else {
                                                    switch "\n\t" {
                                                    case "hello":
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v57 := GenS2{F0: 42, F1: ("a" == "foo bar")}
                                                        if (float64(3.14159) >= ((v2 * v2) - 0.0)) {
                                                            v58 := GenS8{F0: 65535}
                                                            _ = v58
                                                        }
                                                        _ = v57
                                                    }
                                                    v56 = v56
                                                    fmt.Println(v39, v5, fi22)
                                                }
                                                _ = v55
                                                _ = v56
                                            } else {
                                                var v59 int = -1
                                                _ = v59
                                            }
                                        }
                                        if (58 == 127) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                }
                                var v60 int64 = 689
                                _ = fi38
                                _ = v39
                                _ = v60
                            }
                        case -717:
                            fmt.Println(v2, v2, v2)
                        default:
                            delete(v4, 959)
                        }
                    }
                }
                if (v2 != v2) {
                    if (!false) {
                        switch fi22 {
                        case 0:
                            if true {
                                fmt.Println(v2, v5, fi22)
                                if (-1 < 42) {
                                    switch ("foo bar" + ("hello" + "hello")) {
                                    case "\n\t":
                                        v5 -= (v5 << ((42) & 63))
                                    default:
                                        var v61 int = fi22
                                        _ = v61
                                    }
                                } else {
                                    v5 -= 2
                                    v2 -= (v2 / v2)
                                    for fi62 := 0; fi62 < 1; fi62++ {
                                        if (127 <= 128) {
                                            var v63 uint32 = 0
                                            var v64 int64 = 1
                                            _ = v63
                                            _ = v64
                                        } else {
                                            var v65 bool = (979 == 255)
                                            delete(v3, (fi22 / ((2147483647) | 1)))
                                            var v66 int = 32768
                                            _ = v65
                                            _ = v66
                                        }
                                        _ = fi62
                                    }
                                }
                                v5 -= 612
                            }
                        default:
                            v67 := GenS4{F0: int16(v5)}
                            _ = v67
                        }
                        var v68 string = (("a" + "foo bar") + (("\n\t" + "foo bar") + "a"))
                        var v69 int8 = 0
                        _ = v68
                        _ = v69
                    } else {
                        v2 = v2
                    }
                } else {
                    switch ("" + ("a" + ("foo bar" + ""))) {
                    case "foo bar":
                        var v70 float32 = float32(1.5)
                        _ = v70
                    case "a":
                        if (0 == 255) {
                            var v71 int32 = 2147483647
                            switch (^(fi22)) {
                            case -9223372036854775808:
                                switch "a" {
                                case "\n\t":
                                    if true {
                                        if (!false) {
                                            v72 := v3[(fi22 / ((201) | 1))]
                                            _ = v72
                                        } else {
                                            for fi73 := 0; fi73 < 16; fi73++ {
                                                var v74 int64 = 9223372036854775807
                                                v75 := GenS9{F0: (v74 * int64(fi73)), F1: v74, F2: ("\n\t" + (((("hello" + "a") + "hello") + "\n\t") + "foo bar"))}
                                                if true {
                                                    v74 = (v74 & v74)
                                                } else {
                                                    fmt.Println(fi73, v2)
                                                }
                                                _ = fi73
                                                _ = v74
                                                _ = v75
                                            }
                                            if false {
                                                var v76 int32 = 0
                                                fmt.Println(v5, fi22)
                                                fmt.Println(v76, v5, v5)
                                                _ = v76
                                            }
                                        }
                                        if true {
                                            if len(v1) > 0 {
                                                v1[0] = 42
                                            }
                                            v2 -= v2
                                        } else {
                                            for fi77 := 0; fi77 < 2; fi77++ {
                                                v2 -= v2
                                                var v78 uint8 = 183
                                                _ = fi77
                                                _ = v78
                                            }
                                            for fi79 := 0; fi79 < 5; fi79++ {
                                                if (!(!(501 != 42))) {
                                                    v80 := v4[65535]
                                                    var v81 uint32 = 0
                                                    if (0 <= 9223372036854775807) {
                                                        var v82 uint64 = (v5 / ((v5) | 1))
                                                        _ = v82
                                                    }
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    fmt.Println(fi22, v2, v71)
                                                    switch fi22 {
                                                    case -506:
                                                        fmt.Println(v5, fi79)
                                                    case 9223372036854775807:
                                                        v83 := GenS4{F0: 0}
                                                        v84 := GenS2{F0: 42, F1: false}
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                    if true {
                                                        fmt.Println(fi79, fi22, fi79)
                                                        if false {
                                                            var v85 int32 = v71
                                                            delete(v4, 65535)
                                                            v86 := v4[630]
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        _ = len(v1)
                                                    } else {
                                                        v87 := GenS1{F0: (v71 << ((42) & 63))}
                                                        _ = v87
                                                    }
                                                    v71 += (v71 * v71)
                                                }
                                                _ = fi79
                                            }
                                        }
                                        var v88 int64 = 929
                                        _ = v88
                                    } else {
                                        for fi89 := 0; fi89 < 10; fi89++ {
                                            if (2 < 2147483647) {
                                                if (127 > -1) {
                                                    var v90 float64 = v2
                                                    _ = v90
                                                } else {
                                                    v5 -= (v5 / ((97) | 1))
                                                    if true {
                                                        _ = len(v1)
                                                        if (-32768 <= 0) {
                                                            v71 += 2147483647
                                                            v91 := GenS3{F0: 908, F1: 42}
                                                            v2 += (1e10 - ((float64(3.14159) + 3.14159) + -2.25))
                                                            v5 = (v5 % ((4294967295) | 1))
                                                            switch (^(fi22)) {
                                                            case 128:
                                                                v92 := GenS5{F0: 9223372036854775807, F1: ""}
                                                                _ = v92
                                                            case 1:
                                                                var v93 uint8 = 128
                                                                _ = v93
                                                            case -667:
                                                                fmt.Println(v5, v2, v5)
                                                            }
                                                            switch fi22 {
                                                            case 32767:
                                                                var v94 int = (fi89 | -9223372036854775808)
                                                                _ = v94
                                                            default:
                                                                var v95 float64 = 0.1
                                                                _ = v95
                                                            }
                                                            _ = v91
                                                        } else {
                                                            switch (fi89 / ((fi22) | 1)) {
                                                            case 9223372036854775807:
                                                                var v96 int32 = (v71 ^ v71)
                                                                v1 = append(v1, 1)
                                                                v97 := GenS10{F0: -86, F1: 255, F2: fi89}
                                                                var v98 int32 = (v71 << ((408) & 63))
                                                                v71 = (^(v98))
                                                                _ = v96
                                                                _ = v97
                                                                _ = v98
                                                            case 0:
                                                                switch fi22 {
                                                                case 42:
                                                                    v2 = 0.0
                                                                case 1:
                                                                    fmt.Println(v2, fi22, fi89)
                                                                default:
                                                                    switch fi89 {
                                                                    case 127:
                                                                        v99 := GenS11{F0: 42, F1: -862}
                                                                        _ = v99
                                                                    case 32767:
                                                                        v5 -= v5
                                                                    default:
                                                                        switch "\n\t" {
                                                                        case "hello":
                                                                            switch fi89 {
                                                                            case 4294967295:
                                                                                if true {
                                                                                    v5 = v5
                                                                                }
                                                                            case 969:
                                                                                v5 = (v5 ^ v5)
                                                                            case 42:
                                                                                if true {
                                                                                    fmt.Println(fi22, v71, v71)
                                                                                    if true {
                                                                                        var v100 int8 = 127
                                                                                        _ = v100
                                                                                    } else {
                                                                                        var v101 uint32 = 2
                                                                                        _ = v101
                                                                                    }
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        v2 -= v2
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 135
                                                                                        }
                                                                                    case "a":
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v2 -= (1e10 - (float64(1e10) - v2))
                                                                                    }
                                                                                    var v102 bool = true
                                                                                    _ = v102
                                                                                } else {
                                                                                    _ = len(v1)
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                if true {
                                                                    if (("a" < "foo bar") || (true || false)) {
                                                                        fmt.Println(v5, fi22)
                                                                        switch fi22 {
                                                                        case -9223372036854775808:
                                                                            if ((v71 <= v71) || (!(!(false && ((!false) || false))))) {
                                                                                fmt.Println(fi89)
                                                                                v1 = append(v1, 0)
                                                                                if false {
                                                                                    if (!(v71 == (v71 << ((170) & 63)))) {
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v2)
                                                                                        case "a":
                                                                                            var v103 int8 = 1
                                                                                            _ = v103
                                                                                        case "foo bar":
                                                                                            v5 = 18446744073709551615
                                                                                        }
                                                                                        var v104 string = (((("" + ("hello" + "a")) + ("\n\t" + ("\n\t" + ("\n\t" + "")))) + "") + "\n\t")
                                                                                        _ = v104
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    if (42 != 0) {
                                                                                        v2 -= (1.5 * (0.1 - -2.25))
                                                                                        v5 = v5
                                                                                        if ((!(4294967295 > 1)) && true) {
                                                                                            if ((65535 > 732) || true) {
                                                                                                v105 := GenS12{F0: 1}
                                                                                                _ = v105
                                                                                            }
                                                                                            v106 := GenS13{F0: -2147483648, F1: 0, F2: fi22}
                                                                                            _ = v106
                                                                                        }
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "a":
                                                                                        if (!false) {
                                                                                            v4[2147483648] = "a"
                                                                                        } else {
                                                                                            v5 -= (v5 >> ((858) & 63))
                                                                                            if false {
                                                                                                delete(v4, 617)
                                                                                                v107 := GenS14{F0: 168, F1: (v71 / ((42) | 1)), F2: 18446744073709551615}
                                                                                                fmt.Println(fi22)
                                                                                                _ = v107
                                                                                            } else {
                                                                                                if (891 != 52) {
                                                                                                    v2 += ((v2 * (v2 / -2.25)) + (v2 * float64(-2.25)))
                                                                                                    var v108 string = ((("hello" + ("hello" + "foo bar")) + "\n\t") + "")
                                                                                                    _ = len(v1)
                                                                                                    _ = v108
                                                                                                } else {
                                                                                                    if (!(42 != int16(v71))) {
                                                                                                        if (-418 <= (fi89 ^ (-(fi89)))) {
                                                                                                            v109 := GenS15{F0: 65535}
                                                                                                            if ("a" < ("" + "foo bar")) {
                                                                                                                _ = len(v4)
                                                                                                                var v110 uint8 = 44
                                                                                                                _ = v110
                                                                                                            } else {
                                                                                                                v2 = 0.0
                                                                                                                v111 := GenS10{F0: 59, F1: 42, F2: fi22}
                                                                                                                var v112 uint = 65535
                                                                                                                _ = v111
                                                                                                                _ = v112
                                                                                                            }
                                                                                                            _ = v109
                                                                                                        } else {
                                                                                                            v113 := GenS16{F0: (0 == 128), F1: 42}
                                                                                                            if false {
                                                                                                                v1 = append(v1, 2)
                                                                                                                fmt.Println(v71)
                                                                                                            }
                                                                                                            if ((false && true) && true) {
                                                                                                                var v114 uint8 = 42
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            _ = v113
                                                                                                        }
                                                                                                        v115 := GenS17{F0: 856, F1: 2, F2: ((^(fi22)) != fi22)}
                                                                                                        switch (("" + "") + "a") {
                                                                                                        case "\n\t":
                                                                                                            v116 := v3[(fi89 << ((615) & 63))]
                                                                                                            switch (-(fi22)) {
                                                                                                            case 2147483647:
                                                                                                                if (-865 != 0) {
                                                                                                                    var v117 int16 = 1000
                                                                                                                    v118 := v4[164]
                                                                                                                    v117 += 22
                                                                                                                    _ = v117
                                                                                                                    _ = v118
                                                                                                                } else {
                                                                                                                    if (int32(v5) < v71) {
                                                                                                                        switch "hello" {
                                                                                                                        case "":
                                                                                                                            v5 = 843
                                                                                                                        case "\n\t":
                                                                                                                            v2 += -2.25
                                                                                                                        default:
                                                                                                                            v119 := GenS17{F0: 1, F1: 2, F2: (-9223372036854775808 != 9223372036854775807)}
                                                                                                                            _ = v119
                                                                                                                        }
                                                                                                                        switch "\n\t" {
                                                                                                                        case "\n\t":
                                                                                                                            if v116 {
                                                                                                                                var v120 bool = (749 != 65535)
                                                                                                                                v116 = true
                                                                                                                                _ = v120
                                                                                                                            } else {
                                                                                                                                v4[0] = (("a" + "\n\t") + (("" + "hello") + ("a" + ("hello" + "a"))))
                                                                                                                                v71 = 426
                                                                                                                                fmt.Println(fi89, v5, v116)
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v121 := GenS13{F0: 128, F1: 255, F2: fi22}
                                                                                                                            _ = v121
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case -9223372036854775808:
                                                                                                                v1 = append(v1, 609)
                                                                                                            default:
                                                                                                                var v122 uint8 = 127
                                                                                                                _ = v122
                                                                                                            }
                                                                                                            v123 := GenS18{F0: "\n\t"}
                                                                                                            switch "a" {
                                                                                                            case "a":
                                                                                                                v71 -= v71
                                                                                                            case "hello":
                                                                                                                var v124 int = (fi89 | int(v5))
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            var v125 int64 = -9223372036854775808
                                                                                                            _ = v116
                                                                                                            _ = v123
                                                                                                            _ = v125
                                                                                                        }
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    switch fi89 {
                                                                                                    case 9223372036854775807:
                                                                                                        fmt.Println(v71, fi89, fi22)
                                                                                                        if (2 <= 138) {
                                                                                                            v2 = (((v2 / v2) - v2) - v2)
                                                                                                            fmt.Println(fi22, fi89, fi89)
                                                                                                            if (false && (!false)) {
                                                                                                                if (!(-2.25 == (-2.25 + 0.1))) {
                                                                                                                    var v126 bool = false
                                                                                                                    _ = v126
                                                                                                                }
                                                                                                                fmt.Println(fi89)
                                                                                                            }
                                                                                                        } else {
                                                                                                            if (fi89 <= -514) {
                                                                                                                switch "foo bar" {
                                                                                                                case "foo bar":
                                                                                                                    switch fi22 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        var v127 uint8 = 42
                                                                                                                        _ = v127
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    switch "" {
                                                                                                                    case "\n\t":
                                                                                                                        v71 += -338
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    fmt.Println(fi89)
                                                                                                                default:
                                                                                                                    switch (("\n\t" + ("\n\t" + "hello")) + "foo bar") {
                                                                                                                    case "a":
                                                                                                                        var v128 uint16 = 255
                                                                                                                        _ = v128
                                                                                                                    case "foo bar":
                                                                                                                        if ((!false) || (((1.5 - -2.25) * (((1.5 * float32((0.0 / (v2 / v2)))) / ((3.14159 / 1.0) / 3.14159)) + float32(v2))) <= 3.14159)) {
                                                                                                                            v71 += v71
                                                                                                                            _ = len(v1)
                                                                                                                        } else {
                                                                                                                            var v129 uint16 = 65535
                                                                                                                            v129 -= v129
                                                                                                                            _ = v129
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v130 := GenS17{F0: 443, F1: 128, F2: (614 <= 2147483647)}
                                                                                                                        _ = v130
                                                                                                                    }
                                                                                                                }
                                                                                                                v131 := GenS19{F0: 255, F1: -113, F2: (-847 <= -32768)}
                                                                                                                switch ("" + ("\n\t" + "a")) {
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(v2)
                                                                                                                case "\n\t":
                                                                                                                    switch "hello" {
                                                                                                                    case "":
                                                                                                                        v71 += (v71 ^ 1)
                                                                                                                    case "\n\t":
                                                                                                                        switch "\n\t" {
                                                                                                                        case "\n\t":
                                                                                                                            var v132 bool = (4294967295 <= 2)
                                                                                                                            _ = v132
                                                                                                                        default:
                                                                                                                            v71 = v71
                                                                                                                            var v133 uint16 = 661
                                                                                                                            _ = v133
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v134 string = ("" + "")
                                                                                                                        _ = v134
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    fmt.Println(fi89, v5)
                                                                                                                }
                                                                                                                _ = v131
                                                                                                            } else {
                                                                                                                var v135 string = ""
                                                                                                                _ = v135
                                                                                                            }
                                                                                                            if false {
                                                                                                                v136 := GenS9{F0: 1, F1: -9223372036854775808, F2: "\n\t"}
                                                                                                                var v137 int16 = 1
                                                                                                                if (!(42 >= 42)) {
                                                                                                                    v3[fi22] = true
                                                                                                                } else {
                                                                                                                    v138 := GenS19{F0: 338, F1: 4, F2: false}
                                                                                                                    v139 := GenS20{F0: 42, F1: v5}
                                                                                                                    fmt.Println(v5, v5)
                                                                                                                    v140 := GenS6{F0: 127, F1: v137, F2: (124 != 12)}
                                                                                                                    switch "\n\t" {
                                                                                                                    case "a":
                                                                                                                        if len(v1) >= 2 {
                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v141 int8 = -43
                                                                                                                    _ = v138
                                                                                                                    _ = v139
                                                                                                                    _ = v140
                                                                                                                    _ = v141
                                                                                                                }
                                                                                                                _ = v136
                                                                                                                _ = v137
                                                                                                            }
                                                                                                        }
                                                                                                        v142 := v4[42]
                                                                                                        _ = v142
                                                                                                    default:
                                                                                                        var v143 bool = (625 >= 863)
                                                                                                        _ = v143
                                                                                                    }
                                                                                                    v71 -= (v71 / (((v71 & v71)) | 1))
                                                                                                }
                                                                                                switch "hello" {
                                                                                                case "hello":
                                                                                                    if (0 == (v5 & v5)) {
                                                                                                        v144 := GenS21{F0: "a"}
                                                                                                        v2 = float64(float32(-2.25))
                                                                                                        v71 += (-(v71))
                                                                                                        v2 = v2
                                                                                                        var v145 uint64 = (v5 * v5)
                                                                                                        _ = v144
                                                                                                        _ = v145
                                                                                                    } else {
                                                                                                        v2 = v2
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    if true {
                                                                                                        delete(v3, fi89)
                                                                                                        v2 += v2
                                                                                                    } else {
                                                                                                        switch (fi89 | 42) {
                                                                                                        case -1:
                                                                                                            switch fi22 {
                                                                                                            case 427:
                                                                                                                v71 += (v71 + v71)
                                                                                                            case 1:
                                                                                                                v71 = v71
                                                                                                                var v146 int16 = 0
                                                                                                                _ = v146
                                                                                                            case 9223372036854775807:
                                                                                                                v5 -= 52
                                                                                                            }
                                                                                                        }
                                                                                                        v147 := GenS17{F0: 2147483647, F1: 202, F2: false}
                                                                                                        _ = v147
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    var v148 bool = true
                                                                                                    _ = v148
                                                                                                }
                                                                                                v5 = v5
                                                                                            }
                                                                                            _ = len(v1)
                                                                                            switch "\n\t" {
                                                                                            case "":
                                                                                                fmt.Println(fi89, v71, fi89)
                                                                                            case "a":
                                                                                                var v149 uint16 = 65535
                                                                                                _ = v149
                                                                                            default:
                                                                                                v71 += (v71 - v71)
                                                                                            }
                                                                                            var v150 uint = 18446744073709551615
                                                                                            _ = v150
                                                                                        }
                                                                                    case "\n\t":
                                                                                        fmt.Println(v5, v2, fi22)
                                                                                    default:
                                                                                        v151 := GenS13{F0: fi22, F1: 4294967295, F2: (fi89 << ((1) & 63))}
                                                                                        _ = v151
                                                                                    }
                                                                                    v71 = -2147483648
                                                                                }
                                                                                var v152 uint64 = (v5 ^ (v5 | v5))
                                                                                if true {
                                                                                    var v153 bool = true
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                    fmt.Println(fi89, fi89)
                                                                                    fmt.Println(v2, v71)
                                                                                    v3[fi22] = ((fi89 % ((fi89) | 1)) > fi89)
                                                                                    var v154 int = int(v5)
                                                                                    _ = v153
                                                                                    _ = v154
                                                                                }
                                                                                v5 += (v152 >> ((2) & 63))
                                                                                switch "" {
                                                                                case "hello":
                                                                                    var v155 int16 = 1
                                                                                    v156 := GenS22{F0: (v71 - 156), F1: 0, F2: 791}
                                                                                    _ = v155
                                                                                    _ = v156
                                                                                case "foo bar":
                                                                                    v157 := GenS23{F0: v152, F1: 889}
                                                                                    _ = v157
                                                                                default:
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = 272
                                                                                    }
                                                                                }
                                                                                _ = v152
                                                                            } else {
                                                                                v71 -= v71
                                                                                v158 := GenS24{F0: 35, F1: true}
                                                                                v159 := GenS18{F0: "foo bar"}
                                                                                v160 := GenS25{F0: 1, F1: (true || true), F2: (1 > 42)}
                                                                                _ = v158
                                                                                _ = v159
                                                                                _ = v160
                                                                            }
                                                                        case 32767:
                                                                            v161 := v4[2147483647]
                                                                            _ = v161
                                                                        case 760:
                                                                            fmt.Println(fi22, v71)
                                                                        default:
                                                                            v5 = 628
                                                                        }
                                                                        v162 := GenS26{F0: 42, F1: 127}
                                                                        _ = v162
                                                                    }
                                                                    if ((!(!(1 >= 42))) && true) {
                                                                        v163 := GenS21{F0: ""}
                                                                        _ = v163
                                                                    }
                                                                    if false {
                                                                        v71 = -2147483648
                                                                        v164 := GenS27{F0: 90, F1: 1, F2: 127}
                                                                        _ = v164
                                                                    }
                                                                }
                                                            }
                                                            var v165 float64 = v2
                                                            v165 += (0.1 - float64(((0.1 / 3.14159) + (1e10 + 1e10))))
                                                            _ = v165
                                                        }
                                                        v71 += v71
                                                    }
                                                    fmt.Println(v2)
                                                    if ((v2 - -2.25) > 0.0) {
                                                        if (true || (1 < int8(v5))) {
                                                            v2 = (v2 / (0.1 + v2))
                                                            v5 += 42
                                                            v71 += 0
                                                        } else {
                                                            v166 := v3[(fi89 ^ (fi89 & (-(fi89))))]
                                                            var v167 uint64 = (v5 >> ((2) & 63))
                                                            var v168 string = ((((("" + ((("a" + "foo bar") + "") + ("foo bar" + "hello"))) + "a") + (((("\n\t" + "foo bar") + "a") + "a") + "")) + "") + ("hello" + ""))
                                                            _ = v166
                                                            _ = v167
                                                            _ = v168
                                                        }
                                                        if (564 < 42) {
                                                            fmt.Println(fi89, v5, v2)
                                                            if (((v2 / v2) - v2) < 3.14159) {
                                                                var v169 uint64 = v5
                                                                v170 := v3[32768]
                                                                v171 := GenS28{F0: (fi89 | (fi22 / ((fi89) | 1)))}
                                                                _ = v169
                                                                _ = v170
                                                                _ = v171
                                                            }
                                                            delete(v3, fi22)
                                                            v3[(fi89 << ((52) & 63))] = true
                                                            var v172 int8 = 1
                                                            _ = v172
                                                        }
                                                    } else {
                                                        v5 = v5
                                                        var v173 float64 = v2
                                                        _ = v173
                                                    }
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v174 := GenS29{F0: true, F1: 1, F2: (v71 ^ (v71 >> ((1) & 63)))}
                                                    _ = v174
                                                case "":
                                                    if true {
                                                        switch fi89 {
                                                        case 2147483647:
                                                            var v175 string = ""
                                                            _ = v175
                                                        case 32768:
                                                            v176 := GenS30{F0: ""}
                                                            _ = v176
                                                        default:
                                                            var v177 uint16 = 0
                                                            _ = v177
                                                        }
                                                        v178 := GenS31{F0: -967, F1: 690}
                                                        switch "a" {
                                                        case "hello":
                                                            var v179 bool = (true || (!true))
                                                            _ = v179
                                                        default:
                                                            switch fi89 {
                                                            case -182:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                        }
                                                        _ = v178
                                                    } else {
                                                        v180 := GenS30{F0: (("" + "hello") + ("\n\t" + ""))}
                                                        fmt.Println(v2, v5)
                                                        v71 -= (^(v71))
                                                        _ = v180
                                                    }
                                                }
                                            }
                                            _ = fi89
                                        }
                                    }
                                }
                            case -1:
                                switch "a" {
                                case "\n\t":
                                    var v181 uint8 = 2
                                    _ = v181
                                case "":
                                    v5 = 429
                                default:
                                    _ = len(v1)
                                }
                            default:
                                v2 = (float64(float32(((v2 - (v2 * 0.0)) * v2))) - (1.5 - 1.5))
                            }
                            _ = v71
                        } else {
                            for fi182 := 0; fi182 < 19; fi182++ {
                                var v183 bool = true
                                _ = fi182
                                _ = v183
                            }
                            var v184 uint16 = 0
                            v185 := GenS15{F0: 822}
                            _ = v184
                            _ = v185
                        }
                    case "hello":
                        v2 = v2
                    }
                }
                var v186 float32 = (3.14159 * 1.5)
                _ = fi22
                _ = v186
            }
        default:
            fmt.Println(v2, v5)
        }
        for fi187 := 0; fi187 < 3; fi187++ {
            var v188 uint32 = 1
            if (!true) {
                var v189 int64 = 0
                switch (("\n\t" + "") + "a") {
                case "":
                    for fi190 := 0; fi190 < 4; fi190++ {
                        if len(v1) > 0 {
                            v1[0] = v188
                        }
                        v189 -= v189
                        var v191 int64 = (v189 >> ((4294967295) & 63))
                        _ = fi190
                        _ = v191
                    }
                default:
                    var v192 float32 = 1.5
                    _ = v192
                }
                var v193 int32 = -506
                _ = v189
                _ = v193
            }
            _ = fi187
            _ = v188
        }
    } else {
        var v194 uint16 = 860
        v194 += (v194 | (v194 ^ 126))
        _ = v194
    }
    if false {
        v5 = v5
    } else {
        var v195 uint = 42
        if true {
            switch "foo bar" {
            case "a":
                if false {
                    for fi196 := 0; fi196 < 13; fi196++ {
                        v197 := v3[fi196]
                        _ = fi196
                        _ = v197
                    }
                    switch 4294967295 {
                    case -1:
                        if (!(0 == 1)) {
                            v5 = v5
                            if (false || true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                fmt.Println(v2)
                                var v198 string = "a"
                                _ = v198
                            }
                        } else {
                            v199 := GenS32{F0: ""}
                            var v200 int32 = -313
                            switch 9223372036854775807 {
                            case 4294967295:
                                v200 = (v200 & (^(v200)))
                            case 32767:
                                v200 -= -1
                            default:
                                v201 := GenS33{F0: 0, F1: v5}
                                _ = v201
                            }
                            _ = v199
                            _ = v200
                        }
                    default:
                        switch 76 {
                        case -1:
                            switch "" {
                            case "\n\t":
                                var v202 uint32 = 2147483648
                                _ = v202
                            default:
                                fmt.Println(v195)
                            }
                        case 128:
                            v2 = 1.5
                        default:
                            for fi203 := 0; fi203 < 7; fi203++ {
                                v204 := GenS1{F0: -1}
                                v4[1] = "a"
                                for fi205 := 0; fi205 < 4; fi205++ {
                                    fmt.Println(fi203, fi205, fi205)
                                    var v206 uint16 = uint16(v195)
                                    var v207 int16 = 1000
                                    _ = fi205
                                    _ = v206
                                    _ = v207
                                }
                                _ = fi203
                                _ = v204
                            }
                        }
                    }
                } else {
                    for fi208 := 0; fi208 < 7; fi208++ {
                        v5 = (v5 - v5)
                        v1 = append(v1, 4294967295)
                        _ = fi208
                    }
                    var v209 string = "\n\t"
                    _ = v209
                }
            }
            for fi210 := 0; fi210 < 9; fi210++ {
                var v211 bool = false
                _ = fi210
                _ = v211
            }
        }
        for fi212 := 0; fi212 < 19; fi212++ {
            var v213 uint8 = 244
            _ = fi212
            _ = v213
        }
        _ = v195
    }
    v5 = (v5 | 42)
    for fi214 := 0; fi214 < 9; fi214++ {
        v215 := GenS34{F0: 1, F1: 127, F2: 0}
        v2 += (v2 - (v2 / (1.5 - 0.0)))
        v2 += v2
        _ = fi214
        _ = v215
    }
    v1 = append(v1, 2147483647)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
