package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int; F1 int8; F2 int64 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int32; F1 int64; F2 int32 }
type GenS7 struct { F0 int; F1 int32; F2 int64 }
type GenS8 struct { F0 int16; F1 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int16; F1 uint16; F2 int8 }
type GenS13 struct { F0 int64; F1 int; F2 int16 }
type GenS14 struct { F0 int64; F1 int32 }
type GenS15 struct { F0 uint16 }
func genFn0(p0 uint, p1 int64, p2 string) uint64 {
    var v7 float64 = (((3.14159 + 1e10) * 1.5) + 0.1)
    v8 := []int{-576, 1}
    if false {
        v8 = append(v8, 32768)
        if ((!(127 == 100)) && true) {
            switch 32768 {
            case 0:
                p2 = p2
            case 2147483647:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            default:
                var v9 bool = true
                _ = v9
            }
            p2 = ("a" + "hello")
        } else {
            var v10 int16 = 42
            var v11 float32 = 0.1
            v10 = (v10 * -32768)
            _ = v10
            _ = v11
        }
    } else {
        fmt.Println(v7, p0)
        p2 = "a"
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int = fi12
        v8 = append(v8, 2147483647)
        var v14 uint8 = 127
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 int, p1 uint8, p2 string) uint8 {
    p2 = p2
    var v15 uint16 = 255
    if true {
        v16 := []int{}
        fmt.Println(p2)
        _ = v16
    } else {
        p2 = p2
        var v17 int = (-(p0))
        v18 := []int{(p0 % ((v17) | 1)), 219}
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        if false {
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 int64 = -9223372036854775808
                v21 -= v21
                switch p2 {
                case "":
                    v22 := GenS1{F0: 65535}
                    _ = v22
                case "foo bar":
                    v23 := GenS2{F0: false, F1: ((p2 + p2) + p2)}
                    _ = v23
                case "a":
                    fmt.Println(p1, p2, v15)
                default:
                    v24 := []int{(fi19 - 128), (p0 % ((fi20) | 1)), (fi20 ^ (fi19 % (((fi20 >> ((0) & 63))) | 1)))}
                    _ = v24
                }
                _ = fi20
                _ = v21
            }
            v25 := []int{p0, 128}
            _ = v25
        } else {
            if ((32767 < 1000) && true) {
                p2 = p2
            }
            var v26 uint8 = p1
            v27 := GenS3{F0: (fi19 >> ((479) & 63)), F1: 1, F2: -1}
            _ = v26
            _ = v27
        }
        var v28 int = (fi19 % (((^(fi19))) | 1))
        p2 = ((p2 + p2) + p2)
        _ = fi19
        _ = v28
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p1
}

func main() {
    var v1 int8 = -128
    var v2 uint = 584
    var v3 int32 = 1
    var v4 uint64 = 2
    var v5 int64 = -1
    v6 := []string{"\n\t", ("\n\t" + "hello"), ("foo bar" + "")}
    if true {
        if ((-(v5)) > 42) {
            switch 42 {
            case 2147483647:
                var v29 uint8 = 42
                _ = v29
            case 302:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case 42:
                v2 = (^(v2))
            }
            if false {
                v6 = append(v6, ("" + ((((("foo bar" + "") + "hello") + "a") + "foo bar") + ("\n\t" + "\n\t"))))
                if true {
                    v1 += v1
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v31 := GenS1{F0: 42}
                        v32 := GenS4{F0: true, F1: (^(v2)), F2: 65535}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                } else {
                    v2 = (v2 << (((^(v2))) & 63))
                }
                switch ((("" + "foo bar") + "hello") + ("" + "")) {
                case "":
                    var v33 uint = 379
                    _ = v33
                case "foo bar":
                    var v34 int64 = (-(v5))
                    _ = v34
                case "hello":
                    for fi35 := 0; fi35 < 4; fi35++ {
                        switch "" {
                        case "a":
                            v3 = -455
                        }
                        if false {
                            v36 := GenS5{F0: 149}
                            if true {
                                for fi37 := 0; fi37 < 14; fi37++ {
                                    switch "" {
                                    case "":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case "\n\t":
                                        var v38 uint8 = 128
                                        _ = v38
                                    }
                                    _ = fi37
                                }
                                switch "a" {
                                case "foo bar":
                                    if (65535 < 801) {
                                        switch 127 {
                                        case 4294967295:
                                            v4 = (v4 + (^(v4)))
                                        case 2147483647:
                                            if true {
                                                switch (fi35 & int(v4)) {
                                                case 32768:
                                                    if (-32768 >= 1000) {
                                                        v1 -= v1
                                                        v6 = append(v6, "a")
                                                        if true {
                                                            if (!true) {
                                                                var v39 int32 = (v3 >> ((1) & 63))
                                                                v39 -= (v3 % ((v39) | 1))
                                                                if (true || (0 >= 641)) {
                                                                    for fi40 := 0; fi40 < 1; fi40++ {
                                                                        v41 := GenS2{F0: (true && true), F1: "hello"}
                                                                        _ = fi40
                                                                        _ = v41
                                                                    }
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                    var v42 uint32 = 710
                                                                    _ = v42
                                                                }
                                                                _ = v39
                                                            }
                                                            if true {
                                                                var v43 bool = (134 > 497)
                                                                _ = v43
                                                            } else {
                                                                var v44 uint16 = 42
                                                                switch (("" + "\n\t") + "hello") {
                                                                case "hello":
                                                                    var v45 float32 = 0.1
                                                                    _ = v45
                                                                }
                                                                var v46 int32 = int32(v2)
                                                                _ = v44
                                                                _ = v46
                                                            }
                                                            for fi47 := 0; fi47 < 13; fi47++ {
                                                                v48 := GenS6{F0: 1, F1: -6, F2: (v3 << ((v2) & 63))}
                                                                fmt.Println(v3, fi35, fi35)
                                                                v4 += v4
                                                                _ = fi47
                                                                _ = v48
                                                            }
                                                        }
                                                    }
                                                case -914:
                                                    _ = len(v6)
                                                case 617:
                                                    v3 = v3
                                                }
                                            }
                                        default:
                                            v2 -= (^(v2))
                                        }
                                    }
                                case "a":
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        if true {
                                            v49 := GenS3{F0: (fi35 ^ fi35), F1: v1, F2: (-(v5))}
                                            if len(v6) > 0 {
                                                v6[0] = "a"
                                            }
                                            var v50 int64 = (v5 ^ v5)
                                            _ = v49
                                            _ = v50
                                        }
                                    case "foo bar":
                                        var v51 float64 = 1.5
                                        _ = v51
                                    case "hello":
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    default:
                                        var v52 bool = false
                                        _ = v52
                                    }
                                case "":
                                    v53 := GenS7{F0: (fi35 + fi35), F1: int32(v1), F2: (v5 | -1)}
                                    _ = v53
                                default:
                                    if true {
                                        switch "foo bar" {
                                        case "":
                                            var v54 int16 = 1000
                                            _ = v54
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = (("foo bar" + "\n\t") + (("hello" + "\n\t") + ("a" + "hello")))
                                            }
                                        }
                                        var v55 int32 = (v3 / (((v3 % ((v3) | 1))) | 1))
                                        _ = v55
                                    }
                                }
                            } else {
                                if len(v6) > 0 {
                                    v6[0] = ""
                                }
                                switch fi35 {
                                case 0:
                                    v5 = (v5 * (^(v5)))
                                default:
                                    var v56 string = "hello"
                                    _ = v56
                                }
                            }
                            _ = v36
                        } else {
                            for fi57 := 0; fi57 < 10; fi57++ {
                                var v58 uint8 = 156
                                if true {
                                    if (true || false) {
                                        v59 := GenS3{F0: (-(fi35)), F1: (^(v1)), F2: v5}
                                        var v60 uint8 = 2
                                        _ = v59
                                        _ = v60
                                    }
                                    v6 = append(v6, "a")
                                    var v61 float64 = 0.0
                                    _ = v61
                                } else {
                                    v2 = (v2 / (((^(v2))) | 1))
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                var v62 uint64 = uint64(v58)
                                _ = fi57
                                _ = v58
                                _ = v62
                            }
                            if true {
                                v5 -= v5
                                if true {
                                    fmt.Println(v3)
                                }
                                var v63 int16 = -32768
                                _ = v63
                            }
                        }
                        fmt.Println(v3, v4)
                        _ = fi35
                    }
                default:
                    if (56 > 1) {
                        var v64 float64 = (1.5 / 1e10)
                        _ = v64
                    }
                }
            }
            v1 = (-(v1))
        } else {
            fmt.Println(v1, v3, v5)
            for fi65 := 0; fi65 < 14; fi65++ {
                v2 = (v2 - 481)
                for fi66 := 0; fi66 < 4; fi66++ {
                    fmt.Println(v5, v5, fi66)
                    v67 := GenS7{F0: fi65, F1: 947, F2: v5}
                    var v68 int8 = (v1 + -1)
                    _ = len(v6)
                    switch (-(fi65)) {
                    case 2147483647:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case -898:
                        v69 := GenS8{F0: 32767, F1: v5}
                        _ = v69
                    case 4294967295:
                        switch (fi66 + 32768) {
                        case 1:
                            if ((v5 >> ((v2) & 63)) <= v5) {
                                fmt.Println(v2, v3, v1)
                                v70 := GenS6{F0: v3, F1: (v5 | 9223372036854775807), F2: int32(v68)}
                                _ = v70
                            } else {
                                v2 -= (v2 * 0)
                                fmt.Println(v2)
                                v2 -= v2
                            }
                        case -9223372036854775808:
                            v71 := GenS9{F0: (v2 % (((v2 ^ (^(v2)))) | 1))}
                            _ = v71
                        case 0:
                            v6 = append(v6, "\n\t")
                        }
                    }
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v1)
                    case "a":
                        switch (fi66 << (((^(v2))) & 63)) {
                        case 1:
                            switch fi66 {
                            case 32767:
                                fmt.Println(v68)
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (("" + "hello") + "hello")
                                }
                            }
                        }
                    }
                    _ = fi66
                    _ = v67
                    _ = v68
                }
                _ = fi65
            }
            var v72 bool = (!false)
            _ = v72
        }
        for fi73 := 0; fi73 < 7; fi73++ {
            var v74 int8 = -128
            switch ("hello" + "\n\t") {
            case "":
                if (!true) {
                    if (-9223372036854775808 == (v5 / ((int64(v3)) | 1))) {
                        if false {
                            if ((0.0 - float64((3.14159 / 1.5))) <= float64(-2.25)) {
                                for fi75 := 0; fi75 < 5; fi75++ {
                                    if (1 <= 358) {
                                        v1 -= (v1 ^ v74)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        if (1 > 0) {
                                            v2 += (^(v2))
                                        } else {
                                            var v76 int8 = (v1 / ((v1) | 1))
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            var v77 uint8 = 1
                                            v78 := GenS9{F0: (v2 - v2)}
                                            var v79 bool = true
                                            if true {
                                                switch fi75 {
                                                case 9223372036854775807:
                                                    var v80 uint = v2
                                                    v81 := GenS1{F0: 0}
                                                    _ = v80
                                                    _ = v81
                                                case -625:
                                                    var v82 int = fi73
                                                    _ = v82
                                                case 0:
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v79)
                                                v83 := GenS6{F0: 205, F1: v5, F2: (v3 >> (((^(v2))) & 63))}
                                                _ = v83
                                            }
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                            _ = v79
                                        }
                                    } else {
                                        switch (fi75 << ((v2) & 63)) {
                                        case 127:
                                            v5 = 0
                                        case 32768:
                                            var v84 uint16 = 694
                                            _ = v84
                                        default:
                                            if len(v6) > 0 {
                                                v6[0] = "a"
                                            }
                                        }
                                        fmt.Println(v3, v1)
                                        _ = len(v6)
                                    }
                                    _ = fi75
                                }
                                v74 = -128
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                            _ = len(v6)
                        } else {
                            _ = len(v6)
                            v74 = v1
                        }
                    } else {
                        for fi85 := 0; fi85 < 20; fi85++ {
                            var v86 uint = (^(v2))
                            _ = fi85
                            _ = v86
                        }
                        var v87 uint16 = 1
                        v4 = (v4 * genFn0((^(v2)), v5, "a"))
                        _ = v87
                    }
                    v88 := GenS10{F0: false}
                    _ = v88
                }
            }
            switch ((("\n\t" + ("\n\t" + "hello")) + "a") + "a") {
            case "hello":
                for fi89 := 0; fi89 < 4; fi89++ {
                    var v90 int16 = -32768
                    _ = fi89
                    _ = v90
                }
            case "\n\t":
                switch "a" {
                case "":
                    v6 = append(v6, "a")
                case "hello":
                    var v91 bool = (3.14159 > 0.0)
                    _ = v91
                case "foo bar":
                    for fi92 := 0; fi92 < 6; fi92++ {
                        var v93 uint32 = uint32(v2)
                        _ = fi92
                        _ = v93
                    }
                default:
                    fmt.Println(v3)
                }
            default:
                var v94 uint8 = 0
                _ = v94
            }
            _ = fi73
            _ = v74
        }
    } else {
        for fi95 := 0; fi95 < 5; fi95++ {
            if true {
                v5 = (v5 - (^(v5)))
            } else {
                switch "foo bar" {
                case "foo bar":
                    var v96 uint16 = 911
                    _ = v96
                default:
                    fmt.Println(fi95, v1)
                }
                v3 = 42
            }
            for fi97 := 0; fi97 < 9; fi97++ {
                v5 -= v5
                _ = fi97
            }
            switch (fi95 + fi95) {
            case 32768:
                v6 = append(v6, "\n\t")
            case 1:
                if (v5 >= 0) {
                    v98 := GenS8{F0: -64, F1: (v5 << ((v2) & 63))}
                    fmt.Println(v2)
                    _ = v98
                }
            case 9223372036854775807:
                if len(v6) > 0 {
                    v6[0] = "\n\t"
                }
            default:
                for fi99 := 0; fi99 < 8; fi99++ {
                    if true {
                        fmt.Println(v4)
                        var v100 string = "foo bar"
                        _ = v100
                    }
                    fmt.Println(fi99)
                    _ = fi99
                }
            }
            _ = fi95
        }
        v101 := GenS1{F0: 65535}
        var v102 int = 42
        _ = v101
        _ = v102
    }
    switch 9223372036854775807 {
    case -1:
        if (float32(-2.25) <= 1.5) {
            for fi103 := 0; fi103 < 18; fi103++ {
                fmt.Println(v1, v5, v1)
                var v104 int16 = -32768
                _ = fi103
                _ = v104
            }
            var v105 uint8 = 128
            switch "foo bar" {
            case "hello":
                var v106 uint8 = 146
                _ = v106
            case "foo bar":
                var v107 string = ("foo bar" + "")
                _ = v107
            case "\n\t":
                var v108 int8 = v1
                _ = v108
            default:
                v105 = (v105 & v105)
            }
            _ = v105
        } else {
            v4 = (v4 << ((v2) & 63))
            for fi109 := 0; fi109 < 2; fi109++ {
                v3 = v3
                v110 := GenS10{F0: (uint16(v4) < 255)}
                _ = fi109
                _ = v110
            }
        }
    case -530:
        v111 := GenS11{F0: 1}
        _ = v111
    default:
        fmt.Println(v1, v3, v2)
    }
    if (!((0.0 - 0.0) >= -2.25)) {
        switch "\n\t" {
        case "hello":
            if false {
                if false {
                    if ((v5 + v5) == v5) {
                        v1 = v1
                        for fi112 := 0; fi112 < 13; fi112++ {
                            for fi113 := 0; fi113 < 19; fi113++ {
                                if len(v6) > 0 {
                                    v6[0] = ("\n\t" + ("" + ((((((("foo bar" + ("" + ("foo bar" + ("a" + "hello")))) + "\n\t") + ("hello" + "")) + "foo bar") + "a") + "") + "\n\t")))
                                }
                                if len(v6) > 0 {
                                    v6[0] = ("" + ("foo bar" + "a"))
                                }
                                switch -2147483648 {
                                case 32768:
                                    if (!(3.14159 <= ((-2.25 - 3.14159) + 1e10))) {
                                        v114 := GenS2{F0: false, F1: ""}
                                        _ = v114
                                    } else {
                                        if false {
                                            v115 := GenS12{F0: -32768, F1: 42, F2: v1}
                                            fmt.Println(fi112)
                                            v5 -= v5
                                            _ = len(v6)
                                            _ = v115
                                        } else {
                                            if true {
                                                v5 = v5
                                                _ = len(v6)
                                                fmt.Println(v3, fi113)
                                                v4 = (v4 << (((v2 << ((2) & 63))) & 63))
                                                v116 := GenS8{F0: 1, F1: (v5 - 0)}
                                                _ = v116
                                            }
                                            switch "hello" {
                                            case "hello":
                                                if (32767 > 808) {
                                                    v1 = -1
                                                } else {
                                                    if (42 < v2) {
                                                        var v117 float32 = ((3.14159 - 0.1) * ((0.0 * -2.25) / 1.0))
                                                        v118 := GenS4{F0: false, F1: (v2 ^ 4294967295), F2: 18446744073709551615}
                                                        _ = v117
                                                        _ = v118
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = "foo bar"
                                                    }
                                                }
                                            case "foo bar":
                                                v119 := GenS8{F0: 0, F1: v5}
                                                _ = v119
                                            }
                                            if (((!(!(1 >= uint32(v4)))) && ((v2 - v2) > (v2 | (^(v2))))) || true) {
                                                v3 += 1
                                            } else {
                                                v5 = v5
                                            }
                                        }
                                        v120 := GenS4{F0: true, F1: v2, F2: (v2 | v2)}
                                        switch int(v4) {
                                        case -9223372036854775808:
                                            v121 := GenS13{F0: 42, F1: (fi112 / ((127) | 1)), F2: 1000}
                                            _ = v121
                                        }
                                        _ = v120
                                    }
                                    v122 := GenS14{F0: v5, F1: (^(v3))}
                                    _ = v122
                                default:
                                    v2 += (^(v2))
                                }
                                _ = fi113
                            }
                            var v123 float64 = float64(1.5)
                            switch 1 {
                            case 288:
                                var v124 int = int(v3)
                                _ = v124
                            }
                            _ = fi112
                            _ = v123
                        }
                    } else {
                        for fi125 := 0; fi125 < 17; fi125++ {
                            v5 = v5
                            var v126 int32 = v3
                            if len(v6) > 0 {
                                v6[0] = ("hello" + "foo bar")
                            }
                            _ = fi125
                            _ = v126
                        }
                        v127 := GenS11{F0: 697}
                        if false {
                            switch 127 {
                            case -1:
                                v4 -= (v4 << ((v2) & 63))
                            default:
                                v5 = -9223372036854775808
                            }
                            _ = len(v6)
                            v6 = append(v6, "\n\t")
                        }
                        _ = v127
                    }
                    v128 := GenS5{F0: genFn1(-170, 128, ("hello" + ("" + "")))}
                    if false {
                        for fi129 := 0; fi129 < 20; fi129++ {
                            v2 += (v2 / ((42) | 1))
                            v1 = v1
                            _ = fi129
                        }
                    }
                    _ = v128
                }
                for fi130 := 0; fi130 < 17; fi130++ {
                    var v131 float32 = 3.14159
                    var v132 int32 = (v3 & v3)
                    _ = fi130
                    _ = v131
                    _ = v132
                }
                if false {
                    fmt.Println(v3)
                    switch "hello" {
                    case "a":
                        switch "" {
                        case "\n\t":
                            v2 = v2
                        case "":
                            v133 := GenS10{F0: true}
                            _ = v133
                        default:
                            v3 += (v3 << ((v2) & 63))
                        }
                    case "hello":
                        v134 := GenS13{F0: (-(v5)), F1: 630, F2: int16(v4)}
                        _ = v134
                    case "":
                        fmt.Println(v2, v2)
                    default:
                        var v135 string = "foo bar"
                        _ = v135
                    }
                } else {
                    v136 := GenS8{F0: 42, F1: v5}
                    _ = v136
                }
            }
        case "\n\t":
            v137 := GenS4{F0: (0 == 128), F1: (^(v2)), F2: (v2 * 951)}
            _ = v137
        case "a":
            fmt.Println(v5, v1, v1)
        default:
            var v138 uint16 = 804
            _ = v138
        }
        _ = len(v6)
        switch 9223372036854775807 {
        case -2:
            v1 += -128
        default:
            for fi139 := 0; fi139 < 7; fi139++ {
                v6 = append(v6, "\n\t")
                switch 32767 {
                case 0:
                    v6 = append(v6, "foo bar")
                case 309:
                    v140 := GenS8{F0: -357, F1: (v5 ^ v5)}
                    _ = v140
                }
                for fi141 := 0; fi141 < 14; fi141++ {
                    fmt.Println(fi139, v2)
                    _ = fi141
                }
                _ = fi139
            }
        }
    }
    fmt.Println(v5)
    var v142 uint8 = 138
    fmt.Println(v3, v142, v4)
    for fi143 := 0; fi143 < 12; fi143++ {
        for fi144 := 0; fi144 < 6; fi144++ {
            switch (-(fi144)) {
            case -395:
                v2 -= v2
                v3 -= v3
                fmt.Println(fi144, fi144)
            case 32768:
                var v145 int64 = -9223372036854775808
                _ = v145
            default:
                var v146 uint16 = 495
                v1 = int8(v3)
                v147 := GenS15{F0: 2}
                var v148 int64 = 528
                _ = v146
                _ = v147
                _ = v148
            }
            switch -541 {
            case 9223372036854775807:
                if true {
                    fmt.Println(fi143)
                } else {
                    if len(v6) > 0 {
                        v6[0] = "hello"
                    }
                }
            default:
                if true {
                    v1 = (-(v1))
                    var v149 float32 = -2.25
                    _ = v149
                }
            }
            _ = fi144
        }
        var v150 uint64 = (v4 + (v4 >> ((v2) & 63)))
        _ = fi143
        _ = v150
    }
    var v151 uint64 = (v4 ^ v4)
    v142 = (^(v142))
    var v152 uint8 = v142
    var v153 int8 = (^(v1))
    if true {
        var v154 uint8 = v152
        _ = v154
    } else {
        var v155 float32 = -2.25
        if true {
            if len(v6) > 0 {
                v6[0] = (("\n\t" + "\n\t") + (("" + "") + "hello"))
            }
        } else {
            if (!("a" == "\n\t")) {
                for fi156 := 0; fi156 < 8; fi156++ {
                    if (!((!(657 < (fi156 << (((v2 * v2)) & 63)))) && false)) {
                        switch (fi156 / (((-(fi156))) | 1)) {
                        case 128:
                            if (v142 != v152) {
                                if len(v6) > 0 {
                                    v6[0] = "\n\t"
                                }
                                var v157 int = fi156
                                fmt.Println(v5, v1)
                                _ = v157
                            }
                        case 32768:
                            if ((uint32(v142) != 42) || (1e10 >= (3.14159 / 0.1))) {
                                if len(v6) > 0 {
                                    v6[0] = "\n\t"
                                }
                            } else {
                                var v158 uint32 = 2
                                v159 := GenS15{F0: 255}
                                _ = v158
                                _ = v159
                            }
                        }
                    }
                    var v160 uint64 = v4
                    _ = fi156
                    _ = v160
                }
            }
            for fi161 := 0; fi161 < 9; fi161++ {
                v4 = uint64(v2)
                v162 := GenS2{F0: (((fi161 % (((fi161 + fi161)) | 1)) != (fi161 >> ((795) & 63))) || (int64(fi161) != (^(v5)))), F1: "foo bar"}
                _ = fi161
                _ = v162
            }
            var v163 uint64 = (v151 * genFn0(1, (v5 >> ((302) & 63)), ("a" + "foo bar")))
            _ = v163
        }
        _ = v155
    }
    var v164 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v142:", v142)
    fmt.Println("v151:", v151)
    fmt.Println("v152:", v152)
    fmt.Println("v153:", v153)
    fmt.Println("v164:", v164)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v142
    _ = v151
    _ = v152
    _ = v153
    _ = v164
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
