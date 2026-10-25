package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint64; F2 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 int; F1 string }
type GenS11 struct { F0 string }
func genFn0(p0 int8) uint {
    for fi6 := 0; fi6 < 17; fi6++ {
        p0 -= p0
        _ = fi6
    }
    p0 += (p0 % ((p0) | 1))
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint16 = 255
    v3 := []uint32{76, 0, 2147483647, 0}
    var v4 uint64 = 32
    var v5 int64 = 1
    switch 127 {
    case 830:
        v7 := GenS1{F0: 68}
        _ = v7
    case -1:
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 string = "\n\t"
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                var v11 uint16 = 0
                if (((0.1 == (1e10 - (1e10 - ((float32(((v1 * float64((0.1 + (-2.25 / float32(1.5))))) / 0.1)) + (-2.25 - -2.25)) + -2.25)))) && (-302 > -2147483648)) || ("hello" < "\n\t")) {
                    fmt.Println(v1, v1)
                    if ((0.0 / ((-2.25 * 0.0) + (float32(v1) - ((0.0 * ((0.1 * -2.25) / 1.5)) + 3.14159)))) != -2.25) {
                        fmt.Println(v9)
                    }
                    fmt.Println(v5, v4, v5)
                    v12 := GenS1{F0: 42}
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    case 0:
        _ = len(v3)
    default:
        var v13 float32 = 0.0
        _ = v13
    }
    if false {
        var v14 int16 = 1
        fmt.Println(v4, v4, v2)
        if len(v3) > 0 {
            v3[0] = 42
        }
        _ = v14
    } else {
        _ = len(v3)
    }
    if (true && false) {
        switch 32767 {
        case 2147483647:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 string = (("hello" + "foo bar") + ("" + "hello"))
                _ = fi15
                _ = v16
            }
        case 42:
            var v17 uint64 = (v4 * (v4 ^ 622))
            _ = v17
        case -2147483648:
            if false {
                if len(v3) > 0 {
                    v3[0] = 1
                }
                switch "hello" {
                case "\n\t":
                    if (true || false) {
                        v2 = v2
                        var v18 uint64 = 4294967295
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        _ = v18
                    } else {
                        var v19 int64 = v5
                        v5 += v19
                        _ = v19
                    }
                case "a":
                    if false {
                        var v20 int16 = 1
                        v21 := GenS2{F0: 92, F1: 2, F2: 42}
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v5 = v5
                            var v23 int = fi22
                            v24 := GenS1{F0: 127}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v20
                        _ = v21
                    } else {
                        v25 := GenS2{F0: 255, F1: v2, F2: 100}
                        if true {
                            _ = len(v3)
                            switch (("a" + "foo bar") + "hello") {
                            case "\n\t":
                                fmt.Println(v1)
                            default:
                                v1 = v1
                            }
                            switch 0 {
                            case 2147483647:
                                v5 -= (v5 % ((1) | 1))
                            }
                        }
                        _ = v25
                    }
                case "hello":
                    switch 668 {
                    case 548:
                        v26 := GenS1{F0: 2}
                        _ = v26
                    case -2147483648:
                        v27 := GenS1{F0: 111}
                        _ = v27
                    }
                default:
                    if (!((false && false) && (-46 < 0))) {
                        v4 = 2
                        var v28 int16 = 1000
                        _ = v28
                    }
                }
            }
        }
    }
    switch "hello" {
    case "a":
        v5 -= -1
    case "":
        if true {
            var v29 string = "hello"
            _ = v29
        } else {
            if (100 != 100) {
                switch "hello" {
                case "foo bar":
                    var v30 int8 = 0
                    _ = v30
                case "a":
                    v2 -= 0
                case "hello":
                    var v31 int64 = v5
                    _ = v31
                default:
                    var v32 bool = (1 < -128)
                    _ = v32
                }
            } else {
                v4 = 2
            }
            if true {
                switch 32768 {
                case 4294967295:
                    v3 = append(v3, 42)
                }
                switch -1 {
                case 127:
                    v3 = append(v3, 0)
                case -889:
                    fmt.Println(v5)
                case 2147483647:
                    v33 := GenS1{F0: 128}
                    _ = v33
                default:
                    for fi34 := 0; fi34 < 16; fi34++ {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            v36 := GenS2{F0: 0, F1: 451, F2: -96}
                            if false {
                                v4 -= (v4 >> ((2) & 63))
                                v1 = v1
                                var v37 float64 = v1
                                v38 := GenS1{F0: 128}
                                fmt.Println(v2)
                                _ = v37
                                _ = v38
                            }
                            _ = fi35
                            _ = v36
                        }
                        if (!true) {
                            var v39 int16 = 1
                            _ = len(v3)
                            v40 := GenS1{F0: 255}
                            _ = v39
                            _ = v40
                        } else {
                            v5 += v5
                        }
                        _ = fi34
                    }
                }
                if (0 != v4) {
                    switch "" {
                    case "hello":
                        switch 2147483647 {
                        case 128:
                            var v41 string = ("\n\t" + (("a" + "\n\t") + "hello"))
                            _ = v41
                        case 130:
                            var v42 uint8 = 128
                            _ = v42
                        default:
                            var v43 string = "foo bar"
                            _ = v43
                        }
                    case "\n\t":
                        var v44 int32 = 0
                        _ = v44
                    case "":
                        v3 = append(v3, 2)
                    default:
                        v2 -= (v2 - (v2 / ((v2) | 1)))
                    }
                    if false {
                        v45 := GenS3{F0: (542 == 215)}
                        switch (("a" + "\n\t") + "foo bar") {
                        case "\n\t":
                            for fi46 := 0; fi46 < 8; fi46++ {
                                if (((v5 << ((uint(v2)) & 63)) <= (v5 << ((718) & 63))) || false) {
                                    fmt.Println(v4, v1)
                                    var v47 float32 = (0.0 / (float32((0.0 - (v1 + v1))) - (3.14159 * (0.1 / 1.5))))
                                    _ = v47
                                } else {
                                    for fi48 := 0; fi48 < 10; fi48++ {
                                        if (((v5 / ((-920) | 1)) == -519) && ((!((("\n\t" + "a") != "foo bar") && true)) || (false && false))) {
                                            _ = len(v3)
                                            v1 += -2.25
                                            if false {
                                                fmt.Println(v2, v2, v2)
                                            }
                                        }
                                        var v49 int16 = 32767
                                        v49 = (-(v49))
                                        _ = fi48
                                        _ = v49
                                    }
                                }
                                if (586 <= -2147483648) {
                                    fmt.Println(v1)
                                    var v50 float32 = float32(1.5)
                                    var v51 uint64 = v4
                                    _ = v50
                                    _ = v51
                                }
                                v2 = 1
                                _ = fi46
                            }
                        case "foo bar":
                            v5 = v5
                        default:
                            for fi52 := 0; fi52 < 20; fi52++ {
                                fmt.Println(v4)
                                var v53 uint32 = 2147483648
                                for fi54 := 0; fi54 < 2; fi54++ {
                                    _ = len(v3)
                                    _ = fi54
                                }
                                _ = fi52
                                _ = v53
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        _ = v45
                    } else {
                        v55 := GenS4{F0: v5, F1: (^(v4))}
                        if true {
                            v4 += 0
                            var v56 uint = uint(v2)
                            v56 -= (v56 & (v56 - (v56 % (((v56 / ((42) | 1))) | 1))))
                            _ = v56
                        }
                        for fi57 := 0; fi57 < 8; fi57++ {
                            v58 := GenS1{F0: 127}
                            v59 := GenS3{F0: (-1 == v5)}
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                        _ = v55
                    }
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                }
            }
            var v60 int64 = (-(v5))
            _ = v60
        }
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v61 := GenS5{F0: 2}
    if false {
        _ = len(v3)
        var v62 int16 = int16(v2)
        _ = len(v3)
        _ = v62
    } else {
        fmt.Println(v1, v2)
        for fi63 := 0; fi63 < 2; fi63++ {
            v4 = v4
            v3 = append(v3, 42)
            if (9223372036854775807 >= (fi63 / ((128) | 1))) {
                fmt.Println(v2, v4, v1)
            }
            _ = fi63
        }
        var v64 uint16 = 255
        _ = v64
    }
    if (65535 > 569) {
        if false {
            var v65 int16 = 32767
            v66 := GenS4{F0: v5, F1: v4}
            _ = v65
            _ = v66
        } else {
            for fi67 := 0; fi67 < 4; fi67++ {
                fmt.Println(fi67, fi67)
                if len(v3) > 0 {
                    v3[0] = 0
                }
                _ = fi67
            }
            for fi68 := 0; fi68 < 6; fi68++ {
                v5 = v5
                var v69 uint = 4294967295
                _ = fi68
                _ = v69
            }
        }
        v2 += v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        for fi70 := 0; fi70 < 14; fi70++ {
            for fi71 := 0; fi71 < 13; fi71++ {
                v5 += (v5 - v5)
                _ = fi71
            }
            _ = fi70
        }
        var v72 uint8 = 255
        for fi73 := 0; fi73 < 6; fi73++ {
            if false {
                var v74 uint32 = 42
                var v75 bool = false
                _ = v74
                _ = v75
            }
            _ = fi73
        }
        _ = v72
    }
    if (true || ((!true) || (-32768 >= -32768))) {
        if ((0.0 + 1e10) != v1) {
            switch "hello" {
            case "":
                if false {
                    for fi76 := 0; fi76 < 18; fi76++ {
                        switch fi76 {
                        case 128:
                            for fi77 := 0; fi77 < 18; fi77++ {
                                v5 -= (v5 * -366)
                                switch -67 {
                                case -559:
                                    v78 := GenS5{F0: 2}
                                    _ = v78
                                }
                                _ = fi77
                            }
                        case 32768:
                            var v79 int8 = 42
                            _ = v79
                        }
                        switch ("a" + "a") {
                        case "\n\t":
                            v80 := GenS3{F0: (v2 <= v2)}
                            _ = v80
                        }
                        _ = fi76
                    }
                    _ = len(v3)
                    v2 += 42
                }
            case "a":
                v5 = (v5 - v5)
            case "foo bar":
                v4 = v4
            default:
                for fi81 := 0; fi81 < 14; fi81++ {
                    var v82 uint32 = 65535
                    if (1 == 0) {
                        switch -9223372036854775808 {
                        case -2147483648:
                            switch (-(fi81)) {
                            case 0:
                                var v83 float64 = (0.1 * (v1 - v1))
                                _ = v83
                            }
                        default:
                            fmt.Println(v5, v1, v4)
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = fi81
                    _ = v82
                }
            }
            switch 32767 {
            case 127:
                for fi84 := 0; fi84 < 14; fi84++ {
                    var v85 float32 = (0.0 + (float32(v1) * -2.25))
                    _ = fi84
                    _ = v85
                }
            case -9223372036854775808:
                for fi86 := 0; fi86 < 6; fi86++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    for fi87 := 0; fi87 < 11; fi87++ {
                        v88 := GenS1{F0: 37}
                        v3 = append(v3, 703)
                        _ = fi87
                        _ = v88
                    }
                    _ = fi86
                }
            case -369:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                v1 = v1
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            var v89 int16 = 540
            var v90 int8 = -1
            v5 += (v5 * v5)
            _ = v89
            _ = v90
        }
        for fi91 := 0; fi91 < 8; fi91++ {
            v5 -= v5
            var v92 uint = genFn0(0)
            var v93 uint8 = 1
            _ = fi91
            _ = v92
            _ = v93
        }
    } else {
        v94 := GenS6{F0: 135, F1: (v2 + 42)}
        if (!("\n\t" > "")) {
            if (int8(v5) == 0) {
                fmt.Println(v2, v5, v1)
                if false {
                    fmt.Println(v2, v5)
                } else {
                    v2 -= (v2 - (v2 - v2))
                    if true {
                        v95 := GenS2{F0: 128, F1: 255, F2: 100}
                        _ = v95
                    }
                    v4 = 18446744073709551615
                }
            } else {
                var v96 uint8 = 9
                _ = v96
            }
            if len(v3) > 0 {
                v3[0] = 42
            }
        } else {
            v97 := GenS7{F0: 128, F1: 42, F2: false}
            _ = v97
        }
        v98 := GenS4{F0: 737, F1: (v4 % (((v4 / ((v4) | 1))) | 1))}
        _ = v94
        _ = v98
    }
    v99 := GenS4{F0: v5, F1: v4}
    switch "hello" {
    case "hello":
        v100 := GenS8{F0: 0}
        _ = v100
    case "":
        var v101 uint64 = v4
        _ = v101
    default:
        fmt.Println(v2, v5, v1)
    }
    if ((v5 | 42) > v5) {
        if (false && false) {
            if false {
                var v102 uint8 = 11
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v103 := GenS5{F0: 2147483648}
                _ = v102
                _ = v103
            }
            fmt.Println(v2)
            v104 := GenS1{F0: 128}
            _ = v104
        } else {
            switch "foo bar" {
            case "\n\t":
                v4 += v4
            case "hello":
                switch -701 {
                case 32768:
                    if true {
                        var v105 int8 = 1
                        v1 = v1
                        _ = v105
                    } else {
                        v1 = ((0.1 / 1.5) + (((((((v1 * (v1 * 0.1)) / (v1 / (1e10 / (-2.25 + (v1 / (v1 / v1)))))) / 1.0) / v1) - ((0.0 / -2.25) / 0.1)) - v1) / v1))
                    }
                case 42:
                    v106 := GenS4{F0: v5, F1: v4}
                    _ = v106
                default:
                    _ = len(v3)
                }
            case "foo bar":
                if (-1 <= int(v4)) {
                    v107 := GenS9{F0: 18446744073709551615, F1: (v2 + v2)}
                    var v108 bool = false
                    if v108 {
                        for fi109 := 0; fi109 < 17; fi109++ {
                            v4 += 18446744073709551615
                            var v110 int64 = 9223372036854775807
                            _ = fi109
                            _ = v110
                        }
                        var v111 int16 = 0
                        _ = v111
                    } else {
                        v112 := GenS5{F0: 136}
                        fmt.Println(v1, v2, v5)
                        v4 -= (v4 * v4)
                        _ = v112
                    }
                    _ = v107
                    _ = v108
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v5, v4, v5)
                    switch (("" + "") + "a") {
                    case "\n\t":
                        for fi113 := 0; fi113 < 6; fi113++ {
                            v2 -= 245
                            var v114 float64 = v1
                            _ = fi113
                            _ = v114
                        }
                    default:
                        v115 := GenS10{F0: -381, F1: "foo bar"}
                        _ = v115
                    }
                }
            }
            switch "foo bar" {
            case "hello":
                v116 := GenS11{F0: ("a" + (("" + "") + "a"))}
                _ = v116
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            case "":
                if (-128 != -10) {
                    var v117 uint64 = v4
                    var v118 uint64 = (^(v117))
                    _ = v117
                    _ = v118
                } else {
                    fmt.Println(v4, v2)
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    var v119 int16 = 1000
                    _ = v119
                }
            }
        }
        _ = len(v3)
    } else {
        switch -180 {
        case 42:
            fmt.Println(v1)
        }
        fmt.Println(v2, v1)
        fmt.Println(v2, v4, v2)
    }
    switch (("" + "hello") + ("" + "\n\t")) {
    case "a":
        var v120 int32 = 0
        _ = v120
    }
    var v121 uint = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi122 := 0; fi122 < 5; fi122++ {
        if (true || false) {
            v1 -= v1
            var v123 int8 = 1
            _ = v123
        }
        v4 -= (v4 / ((4294967295) | 1))
        fmt.Println(fi122, v4, v4)
        _ = fi122
    }
    fmt.Println(v2, v1)
    v3 = append(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v99
}
