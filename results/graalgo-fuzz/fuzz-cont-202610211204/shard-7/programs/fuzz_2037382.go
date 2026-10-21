package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 bool; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int64; F1 int16; F2 int16 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 bool; F1 uint8; F2 int8 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 bool; F1 int64; F2 int16 }
type GenS18 struct { F0 uint8 }
type GenS19 struct { F0 string; F1 int64 }
type GenS20 struct { F0 int; F1 uint8; F2 int }
type GenS21 struct { F0 bool; F1 string; F2 int16 }
type GenS22 struct { F0 uint8; F1 uint }
type GenS23 struct { F0 int8 }
type GenS24 struct { F0 bool; F1 bool; F2 uint32 }
type GenS25 struct { F0 int32; F1 int16 }
type GenS26 struct { F0 int64 }
type GenS27 struct { F0 string; F1 uint32 }
type GenS28 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS29 struct { F0 int; F1 uint32; F2 uint16 }
func genFn0(p0 int32, p1 int) uint {
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := []int{(fi7 % ((2147483647) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 int8 = -1
    var v10 int64 = -848
    var v11 int64 = (^(v10))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 883
}
func genFn1() int8 {
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            v13 := GenS1{F0: -685, F1: (true || true)}
            var v14 uint16 = 32767
            v15 := GenS2{F0: "foo bar", F1: v14}
            _ = v13
            _ = v14
            _ = v15
        }
        fmt.Println(fi12)
        var v16 uint8 = 255
        _ = fi12
        _ = v16
    }
    var v17 uint8 = 222
    _ = v17
    return 127
}

func main() {
    var v1 uint = 785
    v2 := make(map[int]int)
    v3 := []uint32{0, 2147483648}
    v4 := []int{32768, 360, 32768, 2147483647}
    v5 := make(map[string]int)
    var v6 float32 = 3.14159
    switch 1 {
    case -698:
        for fi18 := 0; fi18 < 19; fi18++ {
            switch fi18 {
            case 32768:
                switch "hello" {
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        v19 := GenS3{F0: 127, F1: 0, F2: 18446744073709551615}
                        _ = v19
                    }
                case "foo bar":
                    for fi20 := 0; fi20 < 19; fi20++ {
                        var v21 int16 = -1
                        fmt.Println(v1, v21)
                        fmt.Println(fi20, v21, fi18)
                        _ = fi20
                        _ = v21
                    }
                case "hello":
                    fmt.Println(fi18)
                default:
                    if true {
                        v6 += v6
                        v6 = v6
                    } else {
                        if (-2147483648 >= -220) {
                            for fi22 := 0; fi22 < 7; fi22++ {
                                var v23 uint16 = 42
                                _ = len(v3)
                                var v24 float32 = (v6 / (0.1 * 1e10))
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                            v1 = v1
                        }
                        fmt.Println(v6)
                    }
                }
            }
            _ = len(v3)
            _ = fi18
        }
    }
    _ = len(v2)
    fmt.Println(v6)
    switch 84 {
    case 121:
        var v25 float32 = (3.14159 + v6)
        _ = v25
    case 831:
        v1 += 18446744073709551615
    }
    var v26 uint16 = 42
    switch 4294967295 {
    case 432:
        v27 := GenS4{F0: 255, F1: -121}
        _ = v27
    case 42:
        v28 := GenS2{F0: (("hello" + "a") + ((("" + "foo bar") + "foo bar") + "\n\t")), F1: (^(v26))}
        _ = v28
    case 128:
        if (3.14159 == 1.5) {
            switch "hello" {
            case "":
                var v29 uint32 = 0
                _ = v29
            case "a":
                fmt.Println(v1, v6)
            default:
                v26 += (v26 / (((v26 + 65535)) | 1))
            }
            switch ("foo bar" + ("" + "foo bar")) {
            case "":
                switch ((("a" + "hello") + "\n\t") + ("\n\t" + ("foo bar" + ("" + ("hello" + "a"))))) {
                case "hello":
                    var v30 bool = false
                    _ = v30
                case "a":
                    switch 32768 {
                    case -700:
                        if false {
                            var v31 uint8 = 127
                            var v32 float64 = 1e10
                            _ = v31
                            _ = v32
                        }
                    case 32768:
                        var v33 uint = 255
                        _ = v33
                    default:
                        switch 1 {
                        case 127:
                            var v34 int = 2147483647
                            _ = v34
                        case 128:
                            v2[4294967295] = 128
                        case 4294967295:
                            switch "a" {
                            case "":
                                if (0.1 == -2.25) {
                                    if true {
                                        var v35 uint8 = 128
                                        if (128 != 128) {
                                            fmt.Println(v26, v35)
                                            _ = len(v5)
                                            if true {
                                                v26 -= v26
                                                if false {
                                                    if true {
                                                        for fi36 := 0; fi36 < 3; fi36++ {
                                                            if false {
                                                                if ((fi36 > (fi36 + fi36)) || true) {
                                                                    _ = len(v4)
                                                                }
                                                            }
                                                            _ = fi36
                                                        }
                                                    } else {
                                                        switch "hello" {
                                                        case "hello":
                                                            for fi37 := 0; fi37 < 4; fi37++ {
                                                                v38 := GenS5{F0: 388, F1: 0}
                                                                if false {
                                                                    for fi39 := 0; fi39 < 10; fi39++ {
                                                                        if true {
                                                                            var v40 uint16 = v26
                                                                            _ = v40
                                                                        } else {
                                                                            var v41 int32 = 42
                                                                            var v42 float32 = (v6 * v6)
                                                                            var v43 uint32 = 42
                                                                            _ = v41
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        _ = fi39
                                                                    }
                                                                    for fi44 := 0; fi44 < 4; fi44++ {
                                                                        if (!true) {
                                                                            v45 := GenS6{F0: false, F1: 32767, F2: 60}
                                                                            _ = v45
                                                                        }
                                                                        _ = fi44
                                                                    }
                                                                    for fi46 := 0; fi46 < 7; fi46++ {
                                                                        v5[""] = fi37
                                                                        _ = len(v3)
                                                                        _ = fi46
                                                                    }
                                                                } else {
                                                                    var v47 int64 = -1
                                                                    if (uint64(v26) >= 276) {
                                                                        if (!(((v35 % ((v35) | 1)) == 255) || (false || true))) {
                                                                            _ = len(v2)
                                                                        }
                                                                    }
                                                                    _ = v47
                                                                }
                                                                v3 = append(v3, 4294967295)
                                                                _ = fi37
                                                                _ = v38
                                                            }
                                                        default:
                                                            if true {
                                                                if (985 >= -1) {
                                                                    v26 += 671
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v48 := GenS2{F0: ("hello" + ("hello" + ("\n\t" + ("" + "\n\t")))), F1: (^(v26))}
                                                    if true {
                                                        var v49 uint32 = 668
                                                        _ = v49
                                                    } else {
                                                        if (false && true) {
                                                            var v50 int32 = 42
                                                            for fi51 := 0; fi51 < 10; fi51++ {
                                                                fmt.Println(v6, fi51)
                                                                _ = fi51
                                                            }
                                                            for fi52 := 0; fi52 < 16; fi52++ {
                                                                var v53 uint32 = 1
                                                                if (v50 == v50) {
                                                                    v6 -= v6
                                                                } else {
                                                                    v50 = (-(v50))
                                                                    v54 := GenS7{F0: (v53 | 4294967295), F1: 18446744073709551615}
                                                                    _ = v54
                                                                }
                                                                _ = fi52
                                                                _ = v53
                                                            }
                                                            _ = v50
                                                        }
                                                        v6 = float32(float64((-2.25 + 0.1)))
                                                    }
                                                    _ = v48
                                                }
                                            }
                                        } else {
                                            switch ("a" + "\n\t") {
                                            case "foo bar":
                                                var v55 int16 = -32768
                                                _ = v55
                                            case "a":
                                                for fi56 := 0; fi56 < 4; fi56++ {
                                                    v35 += v35
                                                    _ = fi56
                                                }
                                            default:
                                                switch 562 {
                                                case -1:
                                                    v26 = 207
                                                default:
                                                    v26 = 553
                                                }
                                            }
                                        }
                                        v1 -= (v1 + v1)
                                        _ = v35
                                    } else {
                                        v6 = (1.5 / (1e10 / v6))
                                        fmt.Println(v1, v26, v26)
                                        v26 += 435
                                    }
                                    for fi57 := 0; fi57 < 16; fi57++ {
                                        var v58 int64 = -1
                                        _ = fi57
                                        _ = v58
                                    }
                                } else {
                                    v1 = (v1 - (^(v1)))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            }
                        }
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = 32768
                    }
                }
            }
            fmt.Println(v1, v26)
        }
    default:
        for fi59 := 0; fi59 < 7; fi59++ {
            switch (("foo bar" + "a") + "foo bar") {
            case "foo bar":
                switch (("foo bar" + ((("foo bar" + ("\n\t" + "foo bar")) + ("foo bar" + "foo bar")) + "a")) + ("hello" + "\n\t")) {
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = (fi59 / (((fi59 | (fi59 & (^(fi59))))) | 1))
                    }
                case "hello":
                    var v60 string = "foo bar"
                    _ = v60
                default:
                    var v61 int32 = -2147483648
                    _ = v61
                }
            default:
                if (!true) {
                    v2[(fi59 & fi59)] = fi59
                    if (float64(float32((1.5 + (1e10 / (3.14159 + float64(v6)))))) > 1.5) {
                        fmt.Println(v6)
                        for fi62 := 0; fi62 < 3; fi62++ {
                            switch fi62 {
                            case 4294967295:
                                var v63 string = ("foo bar" + "foo bar")
                                _ = v63
                            case 2147483647:
                                v64 := GenS8{F0: 0, F1: -1, F2: -539}
                                _ = v64
                            default:
                                fmt.Println(fi62, v6, fi59)
                            }
                            if true {
                                fmt.Println(v6)
                            }
                            _ = fi62
                        }
                    } else {
                        for fi65 := 0; fi65 < 9; fi65++ {
                            v66 := GenS5{F0: 0, F1: 793}
                            fmt.Println(v1, fi65)
                            _ = fi65
                            _ = v66
                        }
                        switch ("a" + ("" + "hello")) {
                        case "hello":
                            if (fi59 < fi59) {
                                if false {
                                    var v67 int = (fi59 ^ (fi59 - (fi59 | fi59)))
                                    fmt.Println(v6, v26, fi59)
                                    _ = v67
                                }
                                v68 := GenS9{F0: "hello", F1: 42}
                                _ = v68
                            }
                        case "foo bar":
                            var v69 float64 = 0.0
                            _ = v69
                        }
                    }
                    if (true && (-634 <= -1)) {
                        var v70 int8 = 42
                        if true {
                            if ((3.14159 + 3.14159) != (3.14159 - float64((v6 + v6)))) {
                                if false {
                                    switch 790 {
                                    case 47:
                                        for fi71 := 0; fi71 < 19; fi71++ {
                                            fmt.Println(v70)
                                            fmt.Println(v6, v1)
                                            if len(v4) > 0 {
                                                v4[0] = (fi71 + (fi71 * 42))
                                            }
                                            _ = fi71
                                        }
                                    default:
                                        v6 -= v6
                                    }
                                    fmt.Println(v1, v6, v26)
                                }
                                fmt.Println(v26)
                            }
                        }
                        _ = len(v3)
                        _ = v70
                    } else {
                        for fi72 := 0; fi72 < 14; fi72++ {
                            v73 := GenS7{F0: 0, F1: 18446744073709551615}
                            fmt.Println(fi72, fi72, fi59)
                            var v74 float64 = 1e10
                            v1 = (v1 % ((380) | 1))
                            v1 = (v1 | (v1 / ((genFn0(-2147483648, fi59)) | 1)))
                            _ = len(v5)
                            _ = fi72
                            _ = v73
                            _ = v74
                        }
                    }
                }
            }
            v3 = append(v3, 849)
            _ = fi59
        }
    }
    for fi75 := 0; fi75 < 6; fi75++ {
        for fi76 := 0; fi76 < 3; fi76++ {
            _ = len(v3)
            var v77 int32 = -1
            switch "foo bar" {
            case "a":
                var v78 uint64 = 42
                _ = v78
            case "\n\t":
                switch (fi76 / (((fi75 / ((787) | 1))) | 1)) {
                case -285:
                    delete(v5, "\n\t")
                    v26 -= v26
                    var v79 uint16 = v26
                    _ = v79
                default:
                    fmt.Println(v26, fi76, v77)
                }
            case "hello":
                v80 := GenS4{F0: v1, F1: 1}
                _ = v80
            default:
                if ((fi75 != fi75) || false) {
                    var v81 float64 = -2.25
                    var v82 string = ("" + "a")
                    var v83 string = v82
                    _ = v81
                    _ = v82
                    _ = v83
                } else {
                    var v84 int = fi76
                    var v85 float64 = 3.14159
                    _ = len(v5)
                    switch "foo bar" {
                    case "":
                        var v86 uint64 = 1
                        _ = v86
                    case "\n\t":
                        var v87 uint64 = 18446744073709551615
                        _ = v87
                    case "hello":
                        var v88 uint16 = (^(v26))
                        _ = v88
                    }
                    switch ("\n\t" + "") {
                    case "":
                        fmt.Println(v85)
                    default:
                        fmt.Println(v26)
                    }
                    _ = v84
                    _ = v85
                }
            }
            fmt.Println(v77, fi75, fi75)
            v89 := GenS10{F0: 219}
            _ = fi76
            _ = v77
            _ = v89
        }
        v90 := GenS11{F0: -9223372036854775808, F1: 42}
        for fi91 := 0; fi91 < 13; fi91++ {
            v1 -= (^(v1))
            _ = fi91
        }
        _ = fi75
        _ = v90
    }
    v26 -= 1
    for fi92 := 0; fi92 < 9; fi92++ {
        for fi93 := 0; fi93 < 5; fi93++ {
            v3 = append(v3, 1)
            _ = fi93
        }
        _ = fi92
    }
    var v94 int16 = -523
    v95 := GenS12{F0: -9223372036854775808}
    for fi96 := 0; fi96 < 11; fi96++ {
        switch -9223372036854775808 {
        case -9223372036854775808:
            for fi97 := 0; fi97 < 5; fi97++ {
                if true {
                    switch (fi96 * fi96) {
                    case 32767:
                        v6 += v6
                    case 128:
                        if (!(599 != int32(v1))) {
                            v26 = 311
                            if ((true || false) && false) {
                                v98 := GenS13{F0: (-2147483648 < 1), F1: 255, F2: 66}
                                _ = v98
                            }
                        }
                    case 164:
                        v99 := v2[fi96]
                        _ = v99
                    default:
                        if (!true) {
                            var v100 int = fi96
                            if (genFn1() == 0) {
                                _ = len(v2)
                                v26 -= (v26 / (((v26 * v26)) | 1))
                                switch 1 {
                                case 0:
                                    fmt.Println(v100, v1, v94)
                                case 127:
                                    fmt.Println(v26)
                                }
                            } else {
                                var v101 float32 = 3.14159
                                fmt.Println(v26, v100, v6)
                                v101 += (v101 * -2.25)
                                _ = v101
                            }
                            var v102 int32 = -2147483648
                            _ = v100
                            _ = v102
                        } else {
                            var v103 bool = true
                            if (42 != genFn1()) {
                                switch "foo bar" {
                                case "foo bar":
                                    if (!((genFn1() == 100) && v103)) {
                                        switch ("\n\t" + "a") {
                                        case "foo bar":
                                            var v104 int32 = -504
                                            _ = v104
                                        case "hello":
                                            var v105 uint = v1
                                            _ = v105
                                        }
                                    }
                                case "":
                                    if v103 {
                                        v26 -= (v26 ^ 0)
                                    } else {
                                        var v106 float32 = v6
                                        switch fi96 {
                                        case 2147483647:
                                            v107 := GenS6{F0: (0 != 203), F1: (^(v26)), F2: 100}
                                            _ = v107
                                        case 32767:
                                            v108 := GenS14{F0: "a"}
                                            _ = v108
                                        default:
                                            v94 -= (v94 & (v94 + (v94 % (((v94 - v94)) | 1))))
                                            v109 := GenS10{F0: uint8(v26)}
                                            v6 = v6
                                            if ((fi96 << (((v1 / (((v1 + v1)) | 1))) & 63)) <= fi96) {
                                                v4 = append(v4, fi96)
                                                v1 = v1
                                            }
                                            var v110 int16 = (-(v94))
                                            var v111 uint = 42
                                            _ = v109
                                            _ = v110
                                            _ = v111
                                        }
                                        _ = v106
                                    }
                                }
                                switch ("hello" + "") {
                                case "a":
                                    v6 = (v6 + (v6 + ((0.0 + v6) - float32(1.5))))
                                case "foo bar":
                                    v1 = (^(v1))
                                }
                            } else {
                                if (32767 < (v94 ^ v94)) {
                                    if (false || v103) {
                                        v6 += v6
                                        var v112 int64 = 1
                                        _ = v112
                                    }
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        v113 := GenS15{F0: -676}
                                        v1 = (v1 ^ v1)
                                        if v103 {
                                            v6 = v6
                                        }
                                        _ = v113
                                    default:
                                        if (1 < 4294967295) {
                                            if len(v3) > 0 {
                                                v3[0] = 2147483648
                                            }
                                            var v114 string = ("" + "\n\t")
                                            if (!(42 < -273)) {
                                                var v115 uint16 = 32767
                                                var v116 uint16 = v26
                                                var v117 uint8 = 172
                                                if v103 {
                                                    var v118 uint32 = uint32(v117)
                                                    v119 := v5["\n\t"]
                                                    v4 = append(v4, 42)
                                                    _ = v118
                                                    _ = v119
                                                }
                                                fmt.Println(fi97, v6)
                                                v114 = "a"
                                                v120 := GenS1{F0: 0, F1: ((1e10 * 1.5) == 1.5)}
                                                var v121 bool = true
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                                _ = v120
                                                _ = v121
                                            }
                                            v122 := GenS16{F0: 140}
                                            switch "a" {
                                            case "":
                                                v123 := v5[(("\n\t" + (v114 + v114)) + ((v114 + "") + ((v114 + v114) + (v114 + "\n\t"))))]
                                                _ = v123
                                            case "hello":
                                                v114 = v114
                                            case "\n\t":
                                                var v124 uint16 = 991
                                                _ = v124
                                            }
                                            _ = v114
                                            _ = v122
                                        } else {
                                            switch (fi96 - (^(fi97))) {
                                            case -2147483648:
                                                if true {
                                                    v125 := GenS17{F0: false, F1: -885, F2: (v94 & v94)}
                                                    var v126 uint = v1
                                                    v127 := GenS18{F0: uint8(v26)}
                                                    _ = v125
                                                    _ = v126
                                                    _ = v127
                                                }
                                            case 53:
                                                if (!v103) {
                                                    var v128 uint16 = v26
                                                    v94 -= (^(v94))
                                                    _ = v128
                                                } else {
                                                    var v129 int32 = 152
                                                    if v103 {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if v103 {
                                                                var v130 uint = (^(v1))
                                                                delete(v5, ("hello" + (((("" + ("foo bar" + "\n\t")) + ("\n\t" + "foo bar")) + ("foo bar" + "\n\t")) + "a")))
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = v130
                                                            } else {
                                                                var v131 uint16 = (v26 * v26)
                                                                if v103 {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v6 += ((1e10 / (v6 + (-2.25 - (3.14159 * 0.1)))) + v6)
                                                                    case "hello":
                                                                        v103 = true
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    default:
                                                                        if (int32(v131) != 541) {
                                                                            v26 -= v131
                                                                            if v103 {
                                                                                v103 = (v103 || v103)
                                                                            } else {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v94 += v94
                                                                    if ((!(0 >= 4294967295)) || (42 != (v131 ^ 65535))) {
                                                                        v132 := GenS15{F0: 9223372036854775807}
                                                                        _ = v132
                                                                    }
                                                                } else {
                                                                    if (false || true) {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            fmt.Println(v94, v94, fi97)
                                                                        case "hello":
                                                                            if ((v6 * float32(3.14159)) < (1.5 / (v6 + (v6 * (v6 / float32(float64((1e10 * v6)))))))) {
                                                                                fmt.Println(v94, v6)
                                                                                var v133 uint64 = 42
                                                                                var v134 int16 = -32768
                                                                                _ = v133
                                                                                _ = v134
                                                                            }
                                                                        case "a":
                                                                            var v135 int64 = -103
                                                                            _ = v135
                                                                        default:
                                                                            v129 += (^(v129))
                                                                        }
                                                                        v6 = 3.14159
                                                                        v136 := GenS12{F0: 1}
                                                                        _ = v136
                                                                    } else {
                                                                        switch (fi97 & fi97) {
                                                                        case 42:
                                                                            var v137 uint8 = 133
                                                                            _ = v137
                                                                        }
                                                                        v6 -= v6
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 65535
                                                                            }
                                                                        case "a":
                                                                            var v138 uint = (^(v1))
                                                                            _ = v138
                                                                        case "":
                                                                            v2[(fi96 | (fi96 | fi97))] = 559
                                                                        default:
                                                                            v1 -= v1
                                                                        }
                                                                    }
                                                                    fmt.Println(v26, v26)
                                                                }
                                                                _ = v131
                                                            }
                                                            v1 = v1
                                                            var v139 uint64 = 663
                                                            fmt.Println(v26, v94, fi96)
                                                            if (uint8(v139) != 0) {
                                                                var v140 string = "a"
                                                                _ = v140
                                                            } else {
                                                                v141 := GenS7{F0: 4294967295, F1: (v139 & 2)}
                                                                v142 := GenS2{F0: "hello", F1: v26}
                                                                fmt.Println(v26, v129, v94)
                                                                _ = v141
                                                                _ = v142
                                                            }
                                                            v6 -= (float32(0.1) / 3.14159)
                                                            _ = v139
                                                        }
                                                    } else {
                                                        fmt.Println(v94, fi96)
                                                        var v143 float64 = (((float64(v6) + ((((((((-2.25 + 1.5) + ((0.1 / 1.5) * 0.0)) - 3.14159) - (1.5 - -2.25)) - (3.14159 * 0.1)) - 3.14159) / (((0.1 / ((1e10 * (1e10 * 0.1)) + ((1.5 - (-2.25 - ((((0.0 / 1.5) - (-2.25 + 0.1)) + 1.5) + ((-2.25 * 1.5) - (((0.0 * 1e10) + 0.0) - float64(float32(((0.0 / 1.0) - 0.0)))))))) - 0.0))) - (0.1 - (0.0 * 3.14159))) * -2.25)) / 1.0)) - (0.1 / 0.1)) + (-2.25 / 3.14159))
                                                        _ = v143
                                                    }
                                                    _ = v129
                                                }
                                            case 2147483647:
                                                v4 = append(v4, 128)
                                            }
                                        }
                                    }
                                    v4 = append(v4, (fi97 % ((fi96) | 1)))
                                    var v144 uint32 = 4294967295
                                    _ = v144
                                }
                                switch fi97 {
                                case 2147483647:
                                    var v145 float64 = ((3.14159 * 1e10) + 1e10)
                                    _ = v145
                                case 1:
                                    if (genFn0(-1, 42) < genFn0(-1, fi97)) {
                                        v94 = (v94 % ((v94) | 1))
                                    } else {
                                        v2[(fi96 / ((128) | 1))] = -9223372036854775808
                                        fmt.Println(v103, fi97)
                                    }
                                case 127:
                                    v5["\n\t"] = fi96
                                }
                                v6 = 3.14159
                            }
                            _ = v103
                        }
                    }
                    if (128 >= 127) {
                        v146 := GenS19{F0: "foo bar", F1: 0}
                        v147 := GenS17{F0: false, F1: 0, F2: v94}
                        switch 128 {
                        case 682:
                            if true {
                                var v148 string = "foo bar"
                                var v149 string = v148
                                if false {
                                    var v150 uint32 = 65535
                                    if (v150 < v150) {
                                        v151 := GenS20{F0: 65, F1: 37, F2: (fi97 ^ fi97)}
                                        var v152 int64 = 983
                                        if false {
                                            var v153 int = fi96
                                            var v154 int32 = 1
                                            if false {
                                                v155 := GenS21{F0: ((false && (fi96 != (fi97 + 42))) && true), F1: v148, F2: -367}
                                                _ = v155
                                            } else {
                                                fmt.Println(v154, fi96, v150)
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = fi96
                                            }
                                            v156 := GenS22{F0: 2, F1: 398}
                                            _ = v153
                                            _ = v154
                                            _ = v156
                                        } else {
                                            v94 += (v94 * v94)
                                        }
                                        _ = v151
                                        _ = v152
                                    }
                                    v157 := GenS16{F0: 1}
                                    _ = v150
                                    _ = v157
                                }
                                _ = v148
                                _ = v149
                            }
                        default:
                            v158 := GenS23{F0: 66}
                            _ = v158
                        }
                        _ = v146
                        _ = v147
                    }
                } else {
                    if false {
                        v159 := GenS24{F0: true, F1: true, F2: 2147483647}
                        delete(v5, ("\n\t" + "foo bar"))
                        switch 4294967295 {
                        case 128:
                            var v160 uint32 = 65535
                            _ = v160
                        case -798:
                            v26 = (v26 ^ v26)
                        }
                        _ = v159
                    } else {
                        if (!true) {
                            switch (fi97 >> ((v1) & 63)) {
                            case -1:
                                var v161 int8 = 1
                                _ = v161
                            }
                            if false {
                                var v162 uint32 = 65535
                                var v163 uint64 = 18446744073709551615
                                v6 -= 3.14159
                                _ = v162
                                _ = v163
                            } else {
                                v5["foo bar"] = fi97
                                v1 = (v1 | v1)
                                v164 := GenS25{F0: -1, F1: int16(v1)}
                                if true {
                                    var v165 uint8 = 188
                                    _ = v165
                                } else {
                                    if false {
                                        if (false || (112 <= -222)) {
                                            var v166 bool = ((true && true) || true)
                                            _ = v166
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                var v167 bool = true
                                                _ = v167
                                            default:
                                                switch "a" {
                                                case "foo bar":
                                                    v26 = v26
                                                case "a":
                                                    var v168 int16 = 0
                                                    v169 := GenS20{F0: fi96, F1: 127, F2: (fi96 * fi97)}
                                                    if false {
                                                        if len(v3) > 0 {
                                                            v3[0] = 2147483648
                                                        }
                                                    } else {
                                                        if false {
                                                            v170 := GenS26{F0: 418}
                                                            _ = v170
                                                        }
                                                        if ((v168 >> ((v1) & 63)) != 42) {
                                                            delete(v5, (("" + "hello") + ""))
                                                            var v171 int16 = (v94 << (((v1 & v1)) & 63))
                                                            if ("foo bar" == "a") {
                                                                if (false && (2147483647 == 2147483648)) {
                                                                    var v172 int32 = 0
                                                                    v173 := GenS27{F0: "", F1: 42}
                                                                    var v174 int64 = -9223372036854775808
                                                                    _ = v172
                                                                    _ = v173
                                                                    _ = v174
                                                                } else {
                                                                    v6 = (v6 + ((((3.14159 + float32(3.14159)) + ((((3.14159 / (float32(0.0) + (v6 + float32(0.1)))) - ((v6 * v6) / 1.5)) - (v6 * (v6 - v6))) / (((0.1 / v6) * v6) + v6))) - v6) / v6))
                                                                    var v175 float64 = (1e10 / (3.14159 * 1.5))
                                                                    v6 += float32(v175)
                                                                    if false {
                                                                        fmt.Println(v94)
                                                                        v1 = 255
                                                                    }
                                                                    _ = v175
                                                                }
                                                            }
                                                            _ = v171
                                                        }
                                                        v176 := GenS28{F0: 42, F1: 42, F2: v26}
                                                        _ = v176
                                                    }
                                                    var v177 uint16 = (v26 | v26)
                                                    _ = v168
                                                    _ = v169
                                                    _ = v177
                                                default:
                                                    switch fi96 {
                                                    case 4294967295:
                                                        v2[fi97] = (fi96 & fi97)
                                                    }
                                                    if ((fi97 >> ((v1) & 63)) >= fi97) {
                                                        v1 = (v1 / ((18446744073709551615) | 1))
                                                        fmt.Println(v6)
                                                    } else {
                                                        var v178 float64 = 3.14159
                                                        _ = v178
                                                    }
                                                }
                                            }
                                        }
                                        if ((v26 ^ v26) < 2) {
                                            fmt.Println(v1)
                                            fmt.Println(fi96, v1, v94)
                                            v179 := GenS29{F0: 2147483647, F1: 0, F2: (v26 ^ v26)}
                                            v5[("\n\t" + ("" + ("foo bar" + ("foo bar" + (("hello" + "hello") + (("a" + "foo bar") + (("hello" + "hello") + "hello")))))))] = fi96
                                            var v180 string = "hello"
                                            _ = v179
                                            _ = v180
                                        } else {
                                            if true {
                                                v181 := GenS3{F0: uint8(v26), F1: 128, F2: 42}
                                                _ = v181
                                            } else {
                                                switch (^(fi96)) {
                                                case -1:
                                                    if false {
                                                        if false {
                                                            var v182 uint8 = 128
                                                            _ = v182
                                                        } else {
                                                            var v183 int64 = 1
                                                            _ = v183
                                                        }
                                                        var v184 uint8 = 42
                                                        _ = v184
                                                    }
                                                case -612:
                                                    v185 := GenS28{F0: 1, F1: 120, F2: v26}
                                                    _ = v185
                                                default:
                                                    if (65535 < (v1 * genFn0(1, (fi96 >> (((^(v1))) & 63))))) {
                                                        if (1.5 != -2.25) {
                                                            switch (^(fi96)) {
                                                            case -9223372036854775808:
                                                                var v186 int64 = 42
                                                                _ = v186
                                                            case -1:
                                                                if false {
                                                                    var v187 uint32 = 639
                                                                    _ = v187
                                                                } else {
                                                                    var v188 uint32 = 4294967295
                                                                    v94 = 1
                                                                    v94 = v94
                                                                    _ = v188
                                                                }
                                                            default:
                                                                var v189 int64 = -1
                                                                _ = v189
                                                            }
                                                        } else {
                                                            _ = len(v3)
                                                            if (!("\n\t" >= "")) {
                                                                fmt.Println(v26, v94)
                                                                fmt.Println(v26, v1)
                                                                v190 := GenS28{F0: 42, F1: 2, F2: 0}
                                                                _ = v190
                                                            }
                                                        }
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v191 := GenS24{F0: true, F1: true, F2: 4294967295}
                                                            _ = v191
                                                        }
                                                        fmt.Println(v6, v6)
                                                    }
                                                }
                                            }
                                            var v192 int64 = 942
                                            switch (("hello" + (("\n\t" + ("" + "foo bar")) + "hello")) + ("foo bar" + "\n\t")) {
                                            case "hello":
                                                delete(v2, fi96)
                                            }
                                            _ = v192
                                        }
                                        if (!((v94 ^ (v94 | v94)) > v94)) {
                                            fmt.Println(v1)
                                        } else {
                                            switch fi96 {
                                            case 380:
                                                v193 := v2[(fi97 ^ fi97)]
                                                _ = v193
                                            }
                                        }
                                    }
                                }
                                fmt.Println(v26)
                                _ = v164
                            }
                        }
                    }
                }
                if (!(!(!(!(("" + "\n\t") < ("a" + "\n\t")))))) {
                    var v194 int = fi97
                    _ = v194
                }
                v5["a"] = fi97
                _ = fi97
            }
        case 346:
            if false {
                v26 += (^(v26))
                for fi195 := 0; fi195 < 2; fi195++ {
                    var v196 int = 971
                    _ = fi195
                    _ = v196
                }
            } else {
                v1 += (v1 ^ 255)
                v197 := v2[fi96]
                _ = v197
            }
        default:
            for fi198 := 0; fi198 < 15; fi198++ {
                fmt.Println(v26, fi96)
                var v199 int32 = 67
                _ = fi198
                _ = v199
            }
        }
        for fi200 := 0; fi200 < 4; fi200++ {
            var v201 float64 = 0.0
            v94 += 0
            _ = fi200
            _ = v201
        }
        _ = fi96
    }
    v202 := GenS7{F0: 2, F1: 687}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v202
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v94
    _ = v95
}
