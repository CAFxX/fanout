package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint; F2 bool }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8; F1 int16; F2 int }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 uint8; F2 int }
type GenS13 struct { F0 uint64; F1 int; F2 bool }
type GenS14 struct { F0 uint }
func genFn0(p0 int32, p1 uint32, p2 int16) string {
    var v5 int32 = p0
    v6 := []int{4294967295, -305, 222, 776}
    if (0 != 4294967295) {
        v7 := GenS1{F0: 65535}
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    v10 := GenS1{F0: 519}
                    _ = v10
                }
                v11 := GenS2{F0: (fi9 & (fi9 & (fi9 >> ((2) & 63))))}
                _ = fi9
                _ = v11
            }
            var v12 bool = ((117 >= 9223372036854775807) && true)
            switch "\n\t" {
            case "foo bar":
                p0 = p0
            case "hello":
                switch ("foo bar" + "a") {
                case "hello":
                    v6 = append(v6, fi8)
                case "":
                    if (!v12) {
                        v5 += (v5 << ((0) & 63))
                        switch ("a" + "a") {
                        case "\n\t":
                            p2 = (^(p2))
                        case "foo bar":
                            switch fi8 {
                            case 9223372036854775807:
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    p1 = p1
                                    var v14 uint = 65535
                                    v15 := GenS2{F0: fi8}
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            case 4294967295:
                                v12 = (0 <= -27)
                            case -2147483648:
                                v16 := GenS3{F0: 0, F1: 18446744073709551615, F2: false}
                                _ = v16
                            default:
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        case "a":
                            switch "" {
                            case "hello":
                                fmt.Println(v12)
                            case "foo bar":
                                switch fi8 {
                                case 32767:
                                    fmt.Println(v12, p0, p0)
                                }
                            default:
                                p0 -= (v5 * p0)
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if (255 >= 128) {
                                p2 += p2
                            }
                        }
                    }
                case "foo bar":
                    if v12 {
                        p2 += p2
                    } else {
                        fmt.Println(v12, p1, fi8)
                        for fi17 := 0; fi17 < 15; fi17++ {
                            fmt.Println(v5)
                            switch (("hello" + "") + "a") {
                            case "a":
                                fmt.Println(p2, p2, p0)
                                fmt.Println(v12, p1)
                                var v18 int32 = p0
                                var v19 uint64 = 938
                                v20 := GenS4{F0: int8(fi8), F1: 9223372036854775807, F2: 704}
                                _ = v18
                                _ = v19
                                _ = v20
                            case "\n\t":
                                if false {
                                    switch (^(fi17)) {
                                    case 292:
                                        if len(v6) > 0 {
                                            v6[0] = (-(fi17))
                                        }
                                    default:
                                        switch fi8 {
                                        case -1:
                                            _ = len(v6)
                                        }
                                    }
                                } else {
                                    if (("foo bar" + ("\n\t" + (("hello" + ("hello" + "hello")) + "foo bar"))) > (("" + (("foo bar" + ("" + "foo bar")) + "foo bar")) + "")) {
                                        v12 = v12
                                        var v21 int = 1
                                        v5 = (p0 | (^(p0)))
                                        _ = v21
                                    }
                                    var v22 float32 = 0.0
                                    var v23 int16 = -32768
                                    _ = v22
                                    _ = v23
                                }
                            default:
                                if ((false || false) && false) {
                                    var v24 uint64 = 1
                                    _ = v24
                                } else {
                                    var v25 int = (fi17 & fi17)
                                    _ = v25
                                }
                            }
                            _ = fi17
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            v5 = p0
                            v27 := GenS3{F0: 42, F1: 273, F2: false}
                            fmt.Println(fi26)
                            switch "a" {
                            case "hello":
                                p1 = 65535
                            case "a":
                                v6 = append(v6, 9223372036854775807)
                            }
                            p2 = (p2 & p2)
                            _ = fi26
                            _ = v27
                        }
                    }
                }
            default:
                for fi28 := 0; fi28 < 6; fi28++ {
                    fmt.Println(v5)
                    _ = fi28
                }
            }
            _ = fi8
            _ = v12
        }
        _ = v7
    } else {
        fmt.Println(p2, p0, p2)
        p2 = 1
    }
    if true {
        for fi29 := 0; fi29 < 20; fi29++ {
            if len(v6) > 0 {
                v6[0] = fi29
            }
            v5 = (-(v5))
            _ = fi29
        }
        switch ("foo bar" + "") {
        case "hello":
            fmt.Println(p2, v5)
        }
    } else {
        var v30 uint32 = p1
        if (!false) {
            p0 = 1
            switch 0 {
            case 632:
                p2 += p2
            }
            v6 = append(v6, -742)
        } else {
            p2 = -343
        }
        if false {
            fmt.Println(v30, v5)
            if true {
                var v31 int8 = 42
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    if ((p1 < (v30 % ((v30) | 1))) && true) {
                        var v32 int64 = -66
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v30)
                        case "a":
                            if false {
                                if len(v6) > 0 {
                                    v6[0] = 1
                                }
                                switch ("a" + ("hello" + "hello")) {
                                case "hello":
                                    var v33 uint32 = v30
                                    _ = v33
                                default:
                                    v31 = 100
                                }
                                _ = len(v6)
                            }
                        }
                        _ = v32
                    } else {
                        v6 = append(v6, 127)
                    }
                default:
                    v34 := GenS4{F0: v31, F1: -961, F2: 1}
                    _ = v34
                }
                p1 -= p1
                _ = v31
            } else {
                v35 := GenS3{F0: 18446744073709551615, F1: 255, F2: (1 == 952)}
                if ((-(p2)) < 42) {
                    var v36 int32 = p0
                    _ = v36
                } else {
                    for fi37 := 0; fi37 < 2; fi37++ {
                        fmt.Println(v30)
                        var v38 int = fi37
                        var v39 uint = 358
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    fmt.Println(v5, p2, v5)
                }
                var v40 int8 = int8(p1)
                _ = v35
                _ = v40
            }
            for fi41 := 0; fi41 < 12; fi41++ {
                if false {
                    _ = len(v6)
                } else {
                    var v42 uint16 = 32767
                    _ = v42
                }
                p2 += p2
                _ = fi41
            }
        }
        _ = v30
    }
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return ("a" + (("\n\t" + (("" + "\n\t") + "foo bar")) + ""))
}
func genFn1() int {
    v43 := GenS2{F0: 313}
    var v44 string = "hello"
    if (!true) {
        v45 := []int{-874, -627, -397, 1}
        _ = v45
    }
    v44 = ((v44 + ("foo bar" + v44)) + v44)
    _ = v43
    _ = v44
    return 128
}

func main() {
    var v1 uint32 = 869
    var v2 string = ("\n\t" + "hello")
    var v3 float64 = (((0.0 - -2.25) + 1e10) / 0.1)
    var v4 int8 = 42
    v46 := GenS5{F0: v1, F1: (-(v4))}
    if (1 != 416) {
        v47 := GenS3{F0: 2, F1: 704, F2: true}
        _ = v47
    }
    switch -1 {
    case 32767:
        if (!true) {
            v4 += (v4 ^ (-(v4)))
            fmt.Println(v2)
            var v48 float64 = (v3 * (v3 * 1.5))
            _ = v48
        }
    case -118:
        fmt.Println(v2)
    case 128:
        v49 := GenS6{F0: 42}
        _ = v49
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        v51 := GenS5{F0: (v1 ^ (v1 & (v1 | (v1 + (v1 - (v1 / (((v1 * (v1 | (v1 >> ((190) & 63))))) | 1))))))), F1: v4}
        v1 = v1
        _ = fi50
        _ = v51
    }
    var v52 string = ""
    if true {
        if (int16(v1) < 0) {
            v3 += 0.1
            var v53 int = -201
            _ = v53
        }
        var v54 int64 = -1
        _ = v54
    } else {
        v55 := []int{}
        _ = v55
    }
    if ((!((!((true && (0.1 >= 0.1)) && true)) || (0 != 40))) && false) {
        switch "foo bar" {
        case "\n\t":
            if (false && false) {
                var v56 int64 = 0
                v57 := []int{32768, -2147483648, 32767}
                _ = v56
                _ = v57
            } else {
                for fi58 := 0; fi58 < 2; fi58++ {
                    for fi59 := 0; fi59 < 20; fi59++ {
                        if (true || true) {
                            v52 = v2
                            v3 = (0.1 + (3.14159 / -2.25))
                            v60 := []int{fi59}
                            _ = v60
                        }
                        switch "hello" {
                        case "\n\t":
                            v61 := GenS7{F0: 19}
                            _ = v61
                        case "foo bar":
                            var v62 uint8 = 0
                            _ = v62
                        default:
                            fmt.Println(v3, v3, fi58)
                            if true {
                                v63 := GenS6{F0: v1}
                                var v64 int16 = 0
                                _ = v63
                                _ = v64
                            }
                        }
                        var v65 int32 = 1
                        _ = fi59
                        _ = v65
                    }
                    _ = fi58
                }
                v1 = (v1 * v1)
                v66 := []int{93, -900, genFn1(), 278}
                _ = v66
            }
        case "hello":
            fmt.Println(v1, v3)
        case "":
            v67 := GenS8{F0: 1}
            _ = v67
        default:
            v68 := GenS9{F0: -903}
            _ = v68
        }
    }
    v52 = v2
    var v69 uint64 = 4294967295
    v70 := []int{9223372036854775807}
    v71 := GenS10{F0: (v4 / ((1) | 1)), F1: 520, F2: 128}
    var v72 int = 1
    for fi73 := 0; fi73 < 12; fi73++ {
        fmt.Println(v3, v3)
        _ = fi73
    }
    for fi74 := 0; fi74 < 11; fi74++ {
        var v75 uint16 = 65535
        _ = fi74
        _ = v75
    }
    v72 = (v72 / ((42) | 1))
    switch (^(v72)) {
    case -697:
        if len(v70) >= 2 {
            v70 = v70[1:len(v70)]
        }
    }
    for fi76 := 0; fi76 < 19; fi76++ {
        for fi77 := 0; fi77 < 20; fi77++ {
            var v78 int16 = -1
            _ = fi77
            _ = v78
        }
        for fi79 := 0; fi79 < 1; fi79++ {
            _ = len(v70)
            var v80 uint64 = v69
            switch (fi79 / ((2147483647) | 1)) {
            case -182:
                v81 := GenS11{F0: 42}
                v80 -= (v80 ^ v69)
                _ = v81
            case -2147483648:
                var v82 string = "hello"
                _ = v82
            default:
                var v83 float64 = 3.14159
                _ = v83
            }
            if ((!true) || true) {
                switch v52 {
                case "\n\t":
                    var v84 float32 = (1e10 - -2.25)
                    _ = v84
                case "hello":
                    v69 = (v69 / ((v80) | 1))
                }
                var v85 uint64 = v80
                _ = v85
            }
            _ = fi79
            _ = v80
        }
        if true {
            var v86 int32 = -193
            _ = v86
        } else {
            fmt.Println(v1)
            if true {
                v87 := GenS12{F0: false, F1: 42, F2: fi76}
                v3 = v3
                if true {
                    v88 := GenS3{F0: 670, F1: 623, F2: (((float32((v3 + v3)) > (1e10 * 0.1)) && ((-1 < int16(v69)) && true)) || false)}
                    for fi89 := 0; fi89 < 17; fi89++ {
                        var v90 float64 = (0.0 * 1e10)
                        v4 = v4
                        v91 := GenS13{F0: v69, F1: fi76, F2: (false && false)}
                        v70 = append(v70, fi89)
                        fmt.Println(v52)
                        fmt.Println(v52, v2, fi89)
                        _ = fi89
                        _ = v90
                        _ = v91
                    }
                    _ = v88
                }
                _ = v87
            } else {
                var v92 int64 = -472
                for fi93 := 0; fi93 < 8; fi93++ {
                    v3 += (v3 * v3)
                    v94 := GenS4{F0: v4, F1: -9223372036854775808, F2: -32768}
                    _ = len(v70)
                    if (1 >= v1) {
                        v95 := GenS5{F0: 0, F1: (v4 & 91)}
                        v96 := GenS13{F0: v69, F1: (^(fi93)), F2: (!((0.1 > (0.1 - (0.1 / 1.5))) && true))}
                        switch fi76 {
                        case 17:
                            var v97 uint = 65535
                            _ = v97
                        case -186:
                            v98 := GenS14{F0: 745}
                            _ = v98
                        }
                        _ = v95
                        _ = v96
                    }
                    if ((!((v69 / (((v69 / ((817) | 1))) | 1)) <= (v69 / (((^(v69))) | 1)))) || false) {
                        switch v52 {
                        case "":
                            var v99 float64 = v3
                            _ = v99
                        }
                        var v100 bool = false
                        if v100 {
                            switch (v72 | fi93) {
                            case 9223372036854775807:
                                if (!v100) {
                                    switch 38 {
                                    case 4294967295:
                                        fmt.Println(v4, fi76, v3)
                                    case -2147483648:
                                        if (v100 && (-2.25 < v3)) {
                                            v70 = append(v70, genFn1())
                                            v70 = append(v70, (v72 + fi93))
                                        }
                                    case -248:
                                        if true {
                                            var v101 string = ("hello" + (v52 + v52))
                                            var v102 bool = v100
                                            if (!false) {
                                                switch v101 {
                                                case "hello":
                                                    var v103 int64 = 9223372036854775807
                                                    _ = v103
                                                case "":
                                                    _ = len(v70)
                                                default:
                                                    v70 = append(v70, genFn1())
                                                }
                                                v104 := GenS5{F0: v1, F1: -128}
                                                var v105 uint16 = 349
                                                _ = v104
                                                _ = v105
                                            } else {
                                                if len(v70) > 0 {
                                                    v70[0] = 558
                                                }
                                                v106 := GenS7{F0: 1}
                                                _ = v106
                                            }
                                            if len(v70) >= 2 {
                                                v70 = v70[1:len(v70)]
                                            }
                                            _ = v101
                                            _ = v102
                                        }
                                    default:
                                        var v107 uint8 = 2
                                        _ = v107
                                    }
                                    var v108 uint = 1
                                    if ((((("a" + v2) + v2) + "\n\t") + ("foo bar" + v52)) == v52) {
                                        var v109 float64 = v3
                                        fmt.Println(v1, fi76)
                                        _ = v109
                                    }
                                    _ = v108
                                }
                            default:
                                switch v2 {
                                case "":
                                    if (361 != 1) {
                                        if true {
                                            fmt.Println(v100, fi93)
                                            if ((v4 & v4) < v4) {
                                                v1 = 1
                                                switch ("hello" + v2) {
                                                case "hello":
                                                    v110 := GenS2{F0: (fi76 & int(v69))}
                                                    _ = v110
                                                case "":
                                                    v92 = (-(v92))
                                                default:
                                                    if len(v70) >= 2 {
                                                        v70 = v70[1:len(v70)]
                                                    }
                                                }
                                                if (v100 && (v100 || ((-1 <= int8(v1)) && true))) {
                                                    switch 32767 {
                                                    case -688:
                                                        if len(v70) >= 2 {
                                                            v70 = v70[1:len(v70)]
                                                        }
                                                    case 127:
                                                        var v111 bool = false
                                                        _ = v111
                                                    case 411:
                                                        v72 += 9223372036854775807
                                                    }
                                                }
                                            }
                                        } else {
                                            if len(v70) > 0 {
                                                v70[0] = -94
                                            }
                                            v3 += 0.0
                                        }
                                    }
                                case "foo bar":
                                    v112 := GenS8{F0: 761}
                                    _ = v112
                                }
                            }
                            v113 := GenS11{F0: 127}
                            v114 := GenS8{F0: v1}
                            _ = v113
                            _ = v114
                        }
                        _ = v100
                    }
                    _ = fi93
                    _ = v94
                }
                _ = v92
            }
            var v115 float64 = (-2.25 / ((0.0 / 1.0) - (0.1 * v3)))
            _ = v115
        }
        _ = fi76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v52
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
