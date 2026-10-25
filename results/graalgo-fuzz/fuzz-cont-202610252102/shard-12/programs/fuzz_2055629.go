package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 int32; F2 int16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int32; F1 uint64 }
type GenS13 struct { F0 uint; F1 int32; F2 int64 }
type GenS14 struct { F0 int; F1 uint }
type GenS15 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS16 struct { F0 uint32; F1 int64; F2 int }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int32; F1 int; F2 int }
type GenS19 struct { F0 string; F1 int; F2 int8 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int; F1 uint16; F2 uint16 }
type GenS22 struct { F0 uint64; F1 uint8 }
func genFn0(p0 bool) uint {
    fmt.Println(p0)
    v4 := GenS1{F0: 32767, F1: "hello"}
    _ = p0
    _ = v4
    return 562
}

func main() {
    v1 := []int{32767, -2147483648, 32767}
    var v2 bool = (((2 <= 2) && false) && ((!true) && true))
    var v3 uint8 = 2
    var v5 int8 = 42
    if v2 {
        switch 1 {
        case 127:
            v6 := GenS2{F0: 255, F1: v3}
            _ = v6
        }
    } else {
        v3 = (v3 << ((2) & 63))
    }
    var v7 int = 32767
    v3 = (v3 / ((v3) | 1))
    v2 = ((((!((255 <= 18446744073709551615) || (true || (1 > 1)))) && ((18446744073709551615 > 18446744073709551615) && v2)) || (("hello" <= ("\n\t" + "\n\t")) || v2)) || ("a" > ((("hello" + (("\n\t" + "foo bar") + "")) + (("" + "a") + "\n\t")) + "hello")))
    var v8 uint32 = 2
    v5 = 100
    var v9 float64 = (0.0 * 3.14159)
    var v10 int32 = 42
    if v2 {
        if v2 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v11 float32 = float32(v9)
            var v12 uint16 = 65535
            _ = v11
            _ = v12
        } else {
            v13 := GenS3{F0: v2}
            var v14 float64 = 3.14159
            fmt.Println(v8)
            _ = v13
            _ = v14
        }
        fmt.Println(v9, v2)
        switch 626 {
        case -2147483648:
            switch (^(v7)) {
            case 9223372036854775807:
                if false {
                    if true {
                        v2 = (-9223372036854775808 <= 9223372036854775807)
                    }
                    fmt.Println(v9, v8, v8)
                    if (true && (uint64(v3) == 18446744073709551615)) {
                        for fi15 := 0; fi15 < 15; fi15++ {
                            v16 := GenS3{F0: ("" <= "hello")}
                            _ = fi15
                            _ = v16
                        }
                        switch "hello" {
                        case "a":
                            for fi17 := 0; fi17 < 8; fi17++ {
                                var v18 int32 = (v10 - 2147483647)
                                v10 += (v10 / ((v18) | 1))
                                fmt.Println(v9, v9)
                                _ = fi17
                                _ = v18
                            }
                        case "":
                            var v19 float32 = float32((0.0 + float64(1e10)))
                            _ = v19
                        default:
                            for fi20 := 0; fi20 < 20; fi20++ {
                                if (!v2) {
                                    v21 := GenS4{F0: 1}
                                    if v2 {
                                        v5 -= v5
                                        switch (("\n\t" + "\n\t") + "") {
                                        case "foo bar":
                                            for fi22 := 0; fi22 < 7; fi22++ {
                                                fmt.Println(v10, v7, v5)
                                                v23 := GenS4{F0: -541}
                                                v2 = (true || v2)
                                                _ = fi22
                                                _ = v23
                                            }
                                        case "":
                                            v2 = true
                                        case "hello":
                                            if v2 {
                                                v24 := GenS1{F0: -2147483648, F1: ("\n\t" + "\n\t")}
                                                _ = v24
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = 32768
                                                }
                                            }
                                        }
                                    }
                                    var v25 int8 = (v5 << ((1) & 63))
                                    _ = v21
                                    _ = v25
                                } else {
                                    var v26 int64 = 213
                                    if v2 {
                                        for fi27 := 0; fi27 < 12; fi27++ {
                                            fmt.Println(v5, fi20)
                                            _ = fi27
                                        }
                                        fmt.Println(v2, v3, v3)
                                    } else {
                                        for fi28 := 0; fi28 < 4; fi28++ {
                                            if true {
                                                if (!(4294967295 < 392)) {
                                                    var v29 int16 = -1
                                                    if (18446744073709551615 == uint(v8)) {
                                                        v9 -= v9
                                                    } else {
                                                        if false {
                                                            v30 := GenS5{F0: (^(v8)), F1: 32768}
                                                            v31 := GenS6{F0: 239}
                                                            _ = v30
                                                            _ = v31
                                                        } else {
                                                            fmt.Println(v29)
                                                        }
                                                        v10 += v10
                                                    }
                                                    _ = v29
                                                } else {
                                                    if v2 {
                                                        switch (fi28 & v7) {
                                                        case 32768:
                                                            v5 += (v5 << ((1) & 63))
                                                        case 1:
                                                            _ = len(v1)
                                                        case 128:
                                                            v32 := GenS2{F0: (^(v3)), F1: v3}
                                                            _ = v32
                                                        }
                                                        if v2 {
                                                            v8 -= v8
                                                            _ = len(v1)
                                                            switch (fi28 - 168) {
                                                            case -1:
                                                                v33 := GenS1{F0: (fi28 & (v7 ^ (-(v7)))), F1: "a"}
                                                                var v34 uint64 = 0
                                                                if (v2 && ((v3 + 105) <= v3)) {
                                                                    if true {
                                                                        if v2 {
                                                                            if v2 {
                                                                                var v35 uint = 42
                                                                                _ = v35
                                                                            } else {
                                                                                v36 := GenS4{F0: v26}
                                                                                v2 = true
                                                                                fmt.Println(v8, v34)
                                                                                _ = v36
                                                                            }
                                                                        }
                                                                        switch ((((("" + "a") + "\n\t") + ("\n\t" + ("\n\t" + "a"))) + ((("\n\t" + "a") + ("a" + "")) + "")) + (("hello" + "a") + ("" + "foo bar"))) {
                                                                        case "a":
                                                                            v34 -= (v34 + v34)
                                                                            if (!v2) {
                                                                                fmt.Println(v7, v8, v9)
                                                                            } else {
                                                                                if (4294967295 == v34) {
                                                                                    var v37 uint32 = (v8 - (^(v8)))
                                                                                    _ = v37
                                                                                }
                                                                                if v2 {
                                                                                    v1 = append(v1, (v7 % (((fi28 - int(v3))) | 1)))
                                                                                } else {
                                                                                    var v38 uint32 = (v8 << ((437) & 63))
                                                                                    _ = v38
                                                                                }
                                                                            }
                                                                            v39 := GenS4{F0: -9223372036854775808}
                                                                            _ = v39
                                                                        case "foo bar":
                                                                            if len(v1) > 0 {
                                                                                v1[0] = fi28
                                                                            }
                                                                        case "\n\t":
                                                                            var v40 float32 = ((1e10 + 3.14159) + -2.25)
                                                                            _ = v40
                                                                        }
                                                                    } else {
                                                                        var v41 int = 2147483647
                                                                        var v42 int = (fi28 - fi28)
                                                                        _ = v41
                                                                        _ = v42
                                                                    }
                                                                }
                                                                _ = v33
                                                                _ = v34
                                                            default:
                                                                var v43 int32 = (v10 + v10)
                                                                _ = v43
                                                            }
                                                        } else {
                                                            v10 += -146
                                                            fmt.Println(v2)
                                                        }
                                                        v5 = (-(v5))
                                                    }
                                                    var v44 uint = 297
                                                    v3 = uint8(v44)
                                                    _ = v44
                                                }
                                                v10 += (v10 | v10)
                                                var v45 uint64 = 0
                                                _ = v45
                                            } else {
                                                switch -145 {
                                                case 0:
                                                    v7 -= (v7 / ((fi28) | 1))
                                                default:
                                                    if (1 != v3) {
                                                        v26 += v26
                                                        v46 := GenS7{F0: (v10 % ((v10) | 1))}
                                                        _ = v46
                                                    } else {
                                                        var v47 uint64 = 351
                                                        _ = v47
                                                    }
                                                }
                                            }
                                            v48 := GenS8{F0: uint(v3)}
                                            v49 := GenS6{F0: 1}
                                            var v50 uint32 = (v8 << ((uint(v8)) & 63))
                                            _ = fi28
                                            _ = v48
                                            _ = v49
                                            _ = v50
                                        }
                                        v51 := GenS9{F0: v5, F1: (-(v10)), F2: -1}
                                        _ = v51
                                    }
                                    for fi52 := 0; fi52 < 14; fi52++ {
                                        v53 := GenS2{F0: v3, F1: (v3 | (v3 & (^(v3))))}
                                        v1 = append(v1, fi52)
                                        _ = fi52
                                        _ = v53
                                    }
                                    _ = v26
                                }
                                if v2 {
                                    v54 := GenS1{F0: v7, F1: "foo bar"}
                                    _ = v54
                                } else {
                                    var v55 uint64 = 127
                                    switch fi20 {
                                    case 42:
                                        if (1 > v3) {
                                            _ = len(v1)
                                            v9 = 1.5
                                        } else {
                                            var v56 uint32 = v8
                                            v55 = 0
                                            for fi57 := 0; fi57 < 2; fi57++ {
                                                if (false || (v3 >= (v3 + (v3 | (^(v3)))))) {
                                                    fmt.Println(v9)
                                                    switch ("a" + ("hello" + (("a" + ("hello" + "foo bar")) + (("foo bar" + "") + "hello")))) {
                                                    case "":
                                                        fmt.Println(v9, v10)
                                                    case "hello":
                                                        _ = len(v1)
                                                    default:
                                                        if true {
                                                            var v58 uint8 = (v3 - 0)
                                                            if v2 {
                                                                v9 -= v9
                                                                v59 := GenS10{F0: v7}
                                                                _ = v59
                                                            }
                                                            v56 = v56
                                                            _ = v58
                                                        }
                                                    }
                                                } else {
                                                    v60 := GenS8{F0: 123}
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        v56 -= 2147483647
                                                    case "hello":
                                                        if ((v9 / (v9 / v9)) != 0.0) {
                                                            if (65535 < uint(v8)) {
                                                                fmt.Println(v3, fi20, v7)
                                                                v61 := GenS11{F0: (1 >= 42)}
                                                                _ = v61
                                                            }
                                                        }
                                                    default:
                                                        var v62 int = (v7 % ((567) | 1))
                                                        _ = v62
                                                    }
                                                    if v2 {
                                                        switch ("\n\t" + ("" + ((("" + "hello") + "a") + (("a" + "foo bar") + (((("\n\t" + ("\n\t" + "a")) + "\n\t") + "a") + "\n\t"))))) {
                                                        case "a":
                                                            var v63 int32 = v10
                                                            _ = v63
                                                        case "foo bar":
                                                            v64 := GenS12{F0: (v10 << ((1) & 63)), F1: v55}
                                                            _ = v64
                                                        }
                                                        v9 += 0.0
                                                    }
                                                    fmt.Println(v2, fi57)
                                                    v56 += uint32(v55)
                                                    _ = v60
                                                }
                                                var v65 uint16 = 893
                                                _ = fi57
                                                _ = v65
                                            }
                                            _ = v56
                                        }
                                    default:
                                        for fi66 := 0; fi66 < 18; fi66++ {
                                            if (v2 && (v5 >= v5)) {
                                                var v67 float64 = v9
                                                switch (("" + ("hello" + "foo bar")) + ("hello" + "")) {
                                                case "hello":
                                                    var v68 int64 = -9223372036854775808
                                                    _ = v68
                                                case "":
                                                    var v69 int = 1
                                                    _ = v69
                                                }
                                                _ = v67
                                            } else {
                                                fmt.Println(v8, fi66, v7)
                                                fmt.Println(fi66, v55, v2)
                                                if false {
                                                    fmt.Println(v9)
                                                    fmt.Println(fi20, v8)
                                                }
                                            }
                                            _ = fi66
                                        }
                                    }
                                    _ = v55
                                }
                                _ = fi20
                            }
                        }
                        var v70 int = (v7 << ((42) & 63))
                        _ = v70
                    }
                } else {
                    v2 = v2
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v71 int = (-(v7))
                    _ = v71
                }
            case 598:
                if false {
                    if v2 {
                        var v72 bool = (!true)
                        if true {
                            fmt.Println(v10)
                            var v73 int8 = (v5 >> ((0) & 63))
                            _ = v73
                        } else {
                            v74 := GenS13{F0: 2, F1: v10, F2: -286}
                            v75 := GenS3{F0: false}
                            switch (v7 ^ (v7 | (v7 & v7))) {
                            case 42:
                                if len(v1) > 0 {
                                    v1[0] = v7
                                }
                            case -801:
                                if v2 {
                                    _ = len(v1)
                                    fmt.Println(v8, v5, v3)
                                    v76 := GenS3{F0: (1 > 1)}
                                    _ = v76
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v8, v72, v72)
                                    fmt.Println(v8, v7)
                                }
                            default:
                                fmt.Println(v3, v8)
                            }
                            _ = v74
                            _ = v75
                        }
                        v1 = append(v1, v7)
                        _ = v72
                    }
                    if true {
                        for fi77 := 0; fi77 < 12; fi77++ {
                            for fi78 := 0; fi78 < 2; fi78++ {
                                fmt.Println(v2)
                                v9 = (v9 - 3.14159)
                                _ = fi78
                            }
                            var v79 uint16 = 65535
                            _ = fi77
                            _ = v79
                        }
                        var v80 uint32 = v8
                        v81 := GenS6{F0: 164}
                        _ = v80
                        _ = v81
                    } else {
                        v5 += (v5 << ((477) & 63))
                    }
                    var v82 uint32 = v8
                    _ = v82
                }
            default:
                var v83 uint16 = uint16(v8)
                _ = v83
            }
        default:
            v9 = 0.1
        }
    }
    fmt.Println(v2)
    v7 = (^(v7))
    var v84 string = ("" + "foo bar")
    var v85 int64 = -1
    switch ((v84 + v84) + v84) {
    case "hello":
        switch (v7 + 4294967295) {
        case 42:
            fmt.Println(v5)
        case -9223372036854775808:
            v86 := GenS3{F0: (!v2)}
            _ = v86
        case -373:
            switch "\n\t" {
            case "a":
                switch ("\n\t" + "") {
                case "":
                    var v87 uint32 = v8
                    _ = v87
                case "\n\t":
                    for fi88 := 0; fi88 < 19; fi88++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        for fi89 := 0; fi89 < 10; fi89++ {
                            v7 += fi89
                            var v90 int16 = 971
                            _ = fi89
                            _ = v90
                        }
                        v10 = (-(v10))
                        _ = fi88
                    }
                case "foo bar":
                    if (!false) {
                        for fi91 := 0; fi91 < 11; fi91++ {
                            if true {
                                var v92 int64 = (v85 / (((v85 % ((281) | 1))) | 1))
                                for fi93 := 0; fi93 < 13; fi93++ {
                                    switch v84 {
                                    case "":
                                        switch (("foo bar" + v84) + "\n\t") {
                                        case "foo bar":
                                            var v94 uint8 = v3
                                            _ = v94
                                        case "":
                                            switch -2147483648 {
                                            case 4294967295:
                                                var v95 bool = true
                                                switch (fi91 | v7) {
                                                case 42:
                                                    var v96 string = ("foo bar" + v84)
                                                    _ = v96
                                                }
                                                v8 = (v8 << ((uint(v8)) & 63))
                                                _ = v95
                                            case 32768:
                                                fmt.Println(v2, v10)
                                            case 2147483647:
                                                var v97 int = fi93
                                                _ = v97
                                            }
                                        default:
                                            var v98 float32 = (0.1 / float32(v9))
                                            _ = v98
                                        }
                                    }
                                    if ((^(v8)) != v8) {
                                        if (!v2) {
                                            if v2 {
                                                v1 = append(v1, (-(fi93)))
                                                if (-554 >= 1) {
                                                    if (((v8 * (v8 % ((0) | 1))) == (v8 * (v8 >> ((genFn0(((v5 << ((genFn0((uint16(v8) >= 255))) & 63)) <= v5))) & 63)))) && false) {
                                                        var v99 uint8 = (v3 - (v3 / ((v3) | 1)))
                                                        if (-785 != v10) {
                                                            v8 -= v8
                                                            switch (v7 << ((974) & 63)) {
                                                            case -145:
                                                                var v100 int64 = (v92 & v85)
                                                                _ = v100
                                                            }
                                                            switch (v84 + v84) {
                                                            case "":
                                                                _ = len(v1)
                                                            case "foo bar":
                                                                v101 := GenS14{F0: int(v99), F1: 1}
                                                                _ = v101
                                                            }
                                                        }
                                                        v2 = ((v10 & (-(v10))) > (v10 - (v10 | -2147483648)))
                                                        _ = v99
                                                    }
                                                    var v102 int = (^(fi93))
                                                    switch fi91 {
                                                    case 794:
                                                        switch (v7 - 128) {
                                                        case 128:
                                                            v92 += (v85 | (v85 | (^(v92))))
                                                        case 2147483647:
                                                            var v103 uint32 = v8
                                                            _ = v103
                                                        }
                                                    case 0:
                                                        v104 := GenS15{F0: 0, F1: v10, F2: v8}
                                                        _ = v104
                                                    case -1:
                                                        v105 := GenS2{F0: (v3 >> ((255) & 63)), F1: (v3 % ((v3) | 1))}
                                                        _ = v105
                                                    default:
                                                        v106 := GenS2{F0: (v3 - 127), F1: (^(v3))}
                                                        _ = v106
                                                    }
                                                    _ = v102
                                                } else {
                                                    if (uint16(v3) >= 334) {
                                                        var v107 float32 = 0.0
                                                        if (1000 == 0) {
                                                            switch "" {
                                                            case "a":
                                                                v5 = (^(v5))
                                                            default:
                                                                var v108 uint16 = 0
                                                                _ = v108
                                                            }
                                                        }
                                                        _ = v107
                                                    } else {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v1 = append(v1, (v7 % (((fi91 ^ (fi93 / ((fi93) | 1)))) | 1)))
                                                        case "foo bar":
                                                            _ = len(v1)
                                                        case "":
                                                            switch (((v84 + ((v84 + "") + v84)) + v84) + ((("hello" + v84) + "hello") + "")) {
                                                            case "hello":
                                                                if len(v1) > 0 {
                                                                    v1[0] = (fi93 / (((-(fi93))) | 1))
                                                                }
                                                            case "a":
                                                                v92 = v85
                                                            case "":
                                                                v109 := GenS16{F0: 465, F1: v92, F2: (fi93 - fi91)}
                                                                var v110 uint32 = 225
                                                                v10 -= (v10 >> ((0) & 63))
                                                                v9 += (float64(-2.25) * (v9 - float64(0.0)))
                                                                v7 = 32767
                                                                _ = v109
                                                                _ = v110
                                                            }
                                                        }
                                                    }
                                                    var v111 uint64 = 460
                                                    _ = v111
                                                }
                                            }
                                            if (!v2) {
                                                v2 = true
                                                v112 := GenS17{F0: v3}
                                                v8 -= v8
                                                if false {
                                                    v85 = (v85 / ((int64(v5)) | 1))
                                                }
                                                _ = v112
                                            }
                                            v9 = 3.14159
                                            v2 = ((92 != (v5 / (((v5 + 1)) | 1))) || ((!(255 > v3)) && (true && v2)))
                                            v113 := GenS2{F0: v3, F1: (^(v3))}
                                            v3 = v3
                                            _ = v113
                                        }
                                        if false {
                                            if v2 {
                                                var v114 bool = (v3 <= (^(v3)))
                                                var v115 int = 870
                                                fmt.Println(v10, v8, v84)
                                                if (v114 && (!v114)) {
                                                    var v116 string = ("\n\t" + ("a" + (v84 + "hello")))
                                                    _ = v116
                                                } else {
                                                    v117 := GenS18{F0: int32(v5), F1: (fi93 << ((genFn0((!true))) & 63)), F2: v115}
                                                    var v118 uint = genFn0(((v7 ^ v7) <= int(v10)))
                                                    switch (v7 << ((233) & 63)) {
                                                    case -2147483648:
                                                        var v119 uint8 = 42
                                                        _ = v119
                                                    }
                                                    _ = v117
                                                    _ = v118
                                                }
                                                _ = v114
                                                _ = v115
                                            } else {
                                                var v120 uint32 = (v8 ^ 886)
                                                if false {
                                                    v2 = v2
                                                } else {
                                                    v121 := GenS19{F0: v84, F1: (fi93 % (((fi93 % ((32768) | 1))) | 1)), F2: (v5 % (((v5 - 100)) | 1))}
                                                    _ = v121
                                                }
                                                _ = v120
                                            }
                                            if false {
                                                var v122 uint8 = v3
                                                if v2 {
                                                    fmt.Println(v10, v84)
                                                }
                                                if v2 {
                                                    var v123 uint32 = v8
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v124 float64 = 3.14159
                                                    _ = v123
                                                    _ = v124
                                                } else {
                                                    v125 := GenS14{F0: fi93, F1: 1}
                                                    _ = v125
                                                }
                                                _ = v122
                                            } else {
                                                var v126 float64 = v9
                                                fmt.Println(v8, v84, v92)
                                                _ = v126
                                            }
                                        }
                                        fmt.Println(fi91, v8)
                                    } else {
                                        v9 += 1e10
                                        var v127 int = (fi93 - -9223372036854775808)
                                        if (255 != 1) {
                                            var v128 uint = 2
                                            _ = v128
                                        } else {
                                            switch (v84 + (v84 + ((("" + v84) + "\n\t") + ((v84 + (v84 + v84)) + v84)))) {
                                            case "\n\t":
                                                v129 := GenS5{F0: (v8 << ((genFn0(((v9 / (((float64(3.14159) * v9) / v9) / ((((1e10 + 3.14159) - v9) * v9) / v9))) != v9))) & 63)), F1: v7}
                                                _ = v129
                                            case "hello":
                                                v8 += v8
                                            case "":
                                                v2 = false
                                            default:
                                                v10 = int32(v92)
                                            }
                                            fmt.Println(v8, v9)
                                            var v130 int32 = v10
                                            _ = v130
                                        }
                                        var v131 int32 = (v10 ^ (^(v10)))
                                        _ = v127
                                        _ = v131
                                    }
                                    var v132 uint = genFn0((((((false || ((v8 - (v8 - v8)) < 568)) || (4294967295 < 18446744073709551615)) || false) && (false || (!(!(v8 > v8))))) && (1 == 1)))
                                    v133 := GenS8{F0: v132}
                                    var v134 uint64 = 42
                                    v85 = v85
                                    _ = fi93
                                    _ = v132
                                    _ = v133
                                    _ = v134
                                }
                                v85 += (^(v92))
                                _ = v92
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi91
                        }
                    }
                default:
                    var v135 int32 = (v10 & v10)
                    _ = v135
                }
            case "":
                if len(v1) > 0 {
                    v1[0] = -9223372036854775808
                }
            case "hello":
                v136 := GenS13{F0: genFn0((2 == 1)), F1: (v10 << ((0) & 63)), F2: v85}
                _ = v136
            default:
                if ((255 <= v3) || (!v2)) {
                    fmt.Println(v7, v85, v9)
                } else {
                    if true {
                        v1 = append(v1, (v7 * v7))
                    }
                }
            }
        default:
            if false {
                if v2 {
                    v85 = v85
                    if (true && ((v5 - (v5 - (v5 >> ((114) & 63)))) != (v5 + v5))) {
                        switch (v7 / (((-(v7))) | 1)) {
                        case 9223372036854775807:
                            if (v2 && (v10 > 2147483647)) {
                                for fi137 := 0; fi137 < 6; fi137++ {
                                    v2 = (((v3 - (v3 % (((v3 * (^(v3)))) | 1))) == v3) && ((v2 || v2) || false))
                                    v138 := GenS20{F0: "\n\t"}
                                    _ = fi137
                                    _ = v138
                                }
                                for fi139 := 0; fi139 < 1; fi139++ {
                                    for fi140 := 0; fi140 < 18; fi140++ {
                                        var v141 int8 = 0
                                        _ = fi140
                                        _ = v141
                                    }
                                    v1 = append(v1, (fi139 + (-(fi139))))
                                    if len(v1) > 0 {
                                        v1[0] = v7
                                    }
                                    _ = fi139
                                }
                            } else {
                                if (int64(v10) < v85) {
                                    fmt.Println(v2, v10, v8)
                                    v7 = (v7 * (^(v7)))
                                }
                                for fi142 := 0; fi142 < 2; fi142++ {
                                    if false {
                                        var v143 int8 = (v5 ^ v5)
                                        if ((v85 / ((v85) | 1)) < (v85 / ((-105) | 1))) {
                                            if true {
                                                var v144 float32 = (0.0 + -2.25)
                                                var v145 int32 = v10
                                                _ = v144
                                                _ = v145
                                            } else {
                                                var v146 string = v84
                                                switch v7 {
                                                case 0:
                                                    if (!v2) {
                                                        switch ("hello" + v146) {
                                                        case "foo bar":
                                                            var v147 float64 = v9
                                                            _ = v147
                                                        case "hello":
                                                            switch v146 {
                                                            case "foo bar":
                                                                if v2 {
                                                                    v146 = ""
                                                                    v146 = v84
                                                                    if v2 {
                                                                        v85 = (v85 / ((int64(fi142)) | 1))
                                                                    } else {
                                                                        if (!false) {
                                                                            if v2 {
                                                                                switch (fi142 >> ((genFn0(v2)) & 63)) {
                                                                                case -549:
                                                                                    switch 9223372036854775807 {
                                                                                    case 698:
                                                                                        var v148 uint16 = 255
                                                                                        _ = v148
                                                                                    default:
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v5 = (^(v5))
                                                                                }
                                                                                var v149 float64 = v9
                                                                                _ = v149
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v150 uint64 = 953
                                                                    var v151 int16 = -199
                                                                    _ = v150
                                                                    _ = v151
                                                                }
                                                            case "\n\t":
                                                                for fi152 := 0; fi152 < 19; fi152++ {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    _ = fi152
                                                                }
                                                            case "a":
                                                                if (18446744073709551615 != 862) {
                                                                    v7 -= 127
                                                                    if false {
                                                                        var v153 uint32 = 111
                                                                        if ((v3 >> ((70) & 63)) > v3) {
                                                                            v154 := GenS20{F0: ""}
                                                                            fmt.Println(v10, v85, v10)
                                                                            _ = v154
                                                                        } else {
                                                                            switch int(v10) {
                                                                            case 32768:
                                                                                v8 += v8
                                                                            case 0:
                                                                                switch (v146 + ((v146 + ("\n\t" + v84)) + "")) {
                                                                                case "a":
                                                                                    if (42 == 0) {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = v7
                                                                                        }
                                                                                        var v155 float32 = 1.5
                                                                                        _ = v155
                                                                                    }
                                                                                default:
                                                                                    var v156 uint8 = (v3 >> ((23) & 63))
                                                                                    _ = v156
                                                                                }
                                                                            case 2147483647:
                                                                                v9 = (v9 / v9)
                                                                            default:
                                                                                var v157 int64 = (v85 / ((v85) | 1))
                                                                                _ = v157
                                                                            }
                                                                            v3 = 1
                                                                            if ((v3 + v3) >= 128) {
                                                                                fmt.Println(fi142)
                                                                                var v158 uint16 = 65535
                                                                                v9 = v9
                                                                                _ = v158
                                                                            }
                                                                        }
                                                                        if v2 {
                                                                            fmt.Println(v2, v153, v10)
                                                                            v159 := GenS21{F0: fi142, F1: 574, F2: 255}
                                                                            v160 := GenS22{F0: 0, F1: 0}
                                                                            _ = v159
                                                                            _ = v160
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(v9, v8, v3)
                                                                            }
                                                                        }
                                                                        _ = v153
                                                                    } else {
                                                                        if false {
                                                                            _ = len(v1)
                                                                            var v161 int16 = -1
                                                                            _ = v161
                                                                        }
                                                                        var v162 int32 = (v10 / (((v10 & v10)) | 1))
                                                                        _ = v162
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v2 = (!(695 != v85))
                                                    }
                                                default:
                                                    for fi163 := 0; fi163 < 11; fi163++ {
                                                        v2 = (v2 && v2)
                                                        _ = fi163
                                                    }
                                                }
                                                var v164 uint = 65535
                                                _ = v146
                                                _ = v164
                                            }
                                            if (669 >= -1) {
                                                v2 = v2
                                                _ = len(v1)
                                            }
                                        }
                                        _ = v143
                                    }
                                    _ = fi142
                                }
                                fmt.Println(v7)
                            }
                        case 1:
                            v85 = 804
                        case 2147483647:
                            _ = len(v1)
                        default:
                            var v165 float32 = 1.5
                            _ = v165
                        }
                        if len(v1) > 0 {
                            v1[0] = (v7 - -9223372036854775808)
                        }
                        if v2 {
                            fmt.Println(v85, v10)
                        } else {
                            _ = len(v1)
                            var v166 int = 32767
                            _ = v166
                        }
                    }
                }
                if (v3 == v3) {
                    fmt.Println(v5, v84, v2)
                    var v167 int = 9223372036854775807
                    if ((v7 % ((-1) | 1)) >= 440) {
                        v84 = (v84 + v84)
                        var v168 int32 = (-(v10))
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = v168
                    }
                    _ = v167
                }
                _ = len(v1)
            }
        }
    case "foo bar":
        var v169 int = v7
        _ = v169
    case "a":
        if (v84 < v84) {
            var v170 int16 = 1
            _ = v170
        } else {
            v171 := GenS13{F0: 532, F1: v10, F2: -9223372036854775808}
            v172 := GenS3{F0: v2}
            if v2 {
                var v173 uint32 = 2147483648
                var v174 int8 = v5
                _ = v173
                _ = v174
            }
            _ = v171
            _ = v172
        }
    }
    v175 := GenS8{F0: 0}
    var v176 float64 = v9
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v176:", v176)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v175
    _ = v176
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v84
    _ = v85
    _ = v9
}
