package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 uint64; F1 int64 }
type GenS9 struct { F0 int; F1 bool; F2 uint16 }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int16; F1 bool }
type GenS17 struct { F0 int32 }

func main() {
    var v1 int32 = -2147483648
    var v2 float64 = -2.25
    var v3 int64 = 1
    var v4 uint64 = 747
    v5 := []uint32{2147483647, 0, 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi6 := 0; fi6 < 18; fi6++ {
        if true {
            for fi7 := 0; fi7 < 5; fi7++ {
                fmt.Println(fi7, v1, v3)
                _ = fi7
            }
        } else {
            var v8 int64 = v3
            if (v2 > 0.0) {
                var v9 int16 = 1
                _ = v9
            } else {
                v8 = (v8 % ((v3) | 1))
                if true {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(fi10, v2, v3)
                        v4 = (v4 ^ (^(v4)))
                        _ = fi10
                    }
                }
                v11 := GenS1{F0: (!false), F1: 85}
                _ = v11
            }
            if true {
                if ((v4 < (v4 & v4)) || (("foo bar" + "") == ("a" + ("hello" + ("" + "a"))))) {
                    var v12 int16 = -615
                    var v13 uint32 = 350
                    var v14 int32 = (^(v1))
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 bool = true
                var v16 uint = 42
                _ = v15
                _ = v16
            }
            _ = v8
        }
        v4 = 616
        v2 += 1.5
        _ = fi6
    }
    switch -2147483648 {
    case 4294967295:
        fmt.Println(v4, v1, v1)
    case -369:
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                v18 := GenS1{F0: true, F1: int8(fi17)}
                for fi19 := 0; fi19 < 8; fi19++ {
                    if (!true) {
                        var v20 int64 = (v3 >> ((2) & 63))
                        v21 := GenS1{F0: true, F1: 72}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint64 = v4
                        var v23 int16 = -141
                        if false {
                            var v24 int = (fi19 >> ((18446744073709551615) & 63))
                            _ = v24
                        } else {
                            fmt.Println(fi19, v2)
                            switch (fi19 | 2147483647) {
                            case 127:
                                v25 := GenS2{F0: (fi19 & 41), F1: 77}
                                _ = v25
                            case -2147483648:
                                _ = len(v5)
                            case 984:
                                v26 := GenS2{F0: -9223372036854775808, F1: 42}
                                _ = v26
                            }
                        }
                        _ = v22
                        _ = v23
                    }
                    if false {
                        switch fi19 {
                        case -1:
                            v27 := GenS1{F0: true, F1: 55}
                            _ = v27
                        case 2147483647:
                            var v28 float32 = -2.25
                            _ = v28
                        }
                    }
                    var v29 uint8 = 42
                    _ = fi19
                    _ = v29
                }
                v1 = -149
                _ = fi17
                _ = v18
            }
            if true {
                if len(v5) > 0 {
                    v5[0] = 0
                }
            }
        }
    default:
        if (true && false) {
            var v30 int16 = 862
            fmt.Println(v30, v30)
            _ = v30
        } else {
            var v31 float32 = ((3.14159 + 3.14159) * 1.5)
            _ = v31
        }
    }
    switch 864 {
    case -466:
        v5 = append(v5, 827)
    case 128:
        if (0.0 >= 0.0) {
            if (222 >= 128) {
                switch ("a" + (("a" + "") + "foo bar")) {
                case "":
                    var v32 uint = 404
                    _ = v32
                case "hello":
                    v1 -= 1
                case "foo bar":
                    var v33 uint16 = 950
                    _ = v33
                }
            }
        } else {
            for fi34 := 0; fi34 < 13; fi34++ {
                v35 := GenS2{F0: (fi34 * -2147483648), F1: 1}
                if false {
                    switch (((("hello" + "") + ("\n\t" + "a")) + "hello") + "\n\t") {
                    case "foo bar":
                        var v36 uint64 = v4
                        _ = v36
                    case "":
                        var v37 uint8 = 0
                        _ = v37
                    case "a":
                        var v38 int8 = 100
                        _ = v38
                    }
                    var v39 uint32 = 0
                    _ = v39
                } else {
                    fmt.Println(v4)
                }
                _ = fi34
                _ = v35
            }
        }
    case 4294967295:
        var v40 int16 = 1
        _ = v40
    }
    var v41 bool = (!true)
    switch 9223372036854775807 {
    case 1:
        v5 = append(v5, 2147483647)
    case -9223372036854775808:
        v2 -= v2
    }
    if true {
        var v42 float32 = (3.14159 / (3.14159 - 0.0))
        var v43 int64 = (v3 * v3)
        fmt.Println(v4)
        _ = v42
        _ = v43
    } else {
        fmt.Println(v1, v41)
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 int32 = v1
        var v46 uint16 = uint16(v4)
        v3 += v3
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 int = 9223372036854775807
    if v41 {
        switch 42 {
        case 32767:
            var v48 uint32 = 42
            _ = v48
        case 945:
            var v49 uint64 = (v4 >> ((0) & 63))
            _ = v49
        }
        switch ("a" + "foo bar") {
        case "hello":
            if v41 {
                var v50 bool = true
                v51 := GenS1{F0: false, F1: 0}
                _ = v50
                _ = v51
            } else {
                fmt.Println(v2)
            }
        case "a":
            if v41 {
                for fi52 := 0; fi52 < 8; fi52++ {
                    var v53 int16 = -1
                    switch "" {
                    case "":
                        v54 := GenS2{F0: fi52, F1: 801}
                        _ = v54
                    case "\n\t":
                        fmt.Println(v53, v4)
                    case "hello":
                        switch "foo bar" {
                        case "":
                            v5 = append(v5, 2147483647)
                        case "\n\t":
                            v2 = v2
                        case "foo bar":
                            switch v47 {
                            case 339:
                                for fi55 := 0; fi55 < 11; fi55++ {
                                    var v56 int64 = (v3 << ((255) & 63))
                                    var v57 uint16 = 32767
                                    _ = fi55
                                    _ = v56
                                    _ = v57
                                }
                            case 502:
                                switch (v47 & (v47 << ((1) & 63))) {
                                case -960:
                                    var v58 bool = (v47 != (^(fi52)))
                                    _ = v58
                                }
                            default:
                                v59 := GenS2{F0: v47, F1: 2147483647}
                                _ = v59
                            }
                        default:
                            v3 = v3
                        }
                    }
                    _ = fi52
                    _ = v53
                }
                for fi60 := 0; fi60 < 9; fi60++ {
                    for fi61 := 0; fi61 < 4; fi61++ {
                        v41 = (1000 > -1)
                        v5 = append(v5, 2147483647)
                        if ((1.5 + 3.14159) <= 3.14159) {
                            v62 := GenS3{F0: 93, F1: ((v4 >= 4294967295) || v41), F2: 163}
                            _ = v62
                        } else {
                            if ((!false) && (0 < uint(v4))) {
                                var v63 uint64 = v4
                                if len(v5) > 0 {
                                    v5[0] = 293
                                }
                                if (!v41) {
                                    _ = len(v5)
                                }
                                _ = v63
                            } else {
                                fmt.Println(v1, v3)
                            }
                            v64 := GenS2{F0: (v47 / ((9223372036854775807) | 1)), F1: 987}
                            _ = v64
                        }
                        _ = fi61
                    }
                    _ = fi60
                }
                if len(v5) > 0 {
                    v5[0] = 1
                }
            } else {
                var v65 uint = 4294967295
                if v41 {
                    var v66 uint8 = 28
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 3; fi67++ {
                        v2 = 1.5
                        _ = fi67
                    }
                    if (((v2 * 1e10) - ((3.14159 + v2) / v2)) != v2) {
                        fmt.Println(v4, v65)
                        v47 = v47
                    } else {
                        var v68 uint64 = 459
                        switch "" {
                        case "hello":
                            var v69 int8 = 42
                            _ = v69
                        case "a":
                            v70 := GenS4{F0: (^(v68)), F1: 818, F2: (^(v4))}
                            _ = v70
                        default:
                            if v41 {
                                var v71 float64 = ((v2 - (1e10 + (v2 + (v2 + ((-2.25 + float64(1e10)) + v2))))) * 3.14159)
                                _ = v71
                            } else {
                                var v72 uint = 848
                                switch (v47 / ((0) | 1)) {
                                case 2147483647:
                                    switch "foo bar" {
                                    case "foo bar":
                                        v68 += (^(v4))
                                    case "hello":
                                        if (!(2147483647 < 1)) {
                                            v73 := GenS1{F0: false, F1: 0}
                                            fmt.Println(v65)
                                            if ("\n\t" == (("" + ((("foo bar" + "foo bar") + ("" + "hello")) + "")) + "foo bar")) {
                                                if (v4 == 0) {
                                                    fmt.Println(v47)
                                                    if true {
                                                        v5 = append(v5, 0)
                                                    }
                                                    for fi74 := 0; fi74 < 6; fi74++ {
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = fi74
                                                    }
                                                }
                                            }
                                            _ = v73
                                        } else {
                                            if v41 {
                                                if true {
                                                    for fi75 := 0; fi75 < 9; fi75++ {
                                                        for fi76 := 0; fi76 < 13; fi76++ {
                                                            fmt.Println(fi75)
                                                            switch "hello" {
                                                            case "a":
                                                                var v77 uint16 = 1
                                                                _ = v77
                                                            case "":
                                                                if false {
                                                                    var v78 int32 = v1
                                                                    _ = v78
                                                                }
                                                            case "\n\t":
                                                                if true {
                                                                    var v79 uint64 = v4
                                                                    var v80 int = fi75
                                                                    var v81 float64 = (v2 / v2)
                                                                    _ = v79
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    var v82 int32 = 42
                                                                    _ = v82
                                                                }
                                                            }
                                                            _ = len(v5)
                                                            _ = fi76
                                                        }
                                                        switch v47 {
                                                        case -1:
                                                            _ = len(v5)
                                                        case 601:
                                                            v5 = append(v5, 320)
                                                        case 4294967295:
                                                            if false {
                                                                var v83 uint64 = v4
                                                                _ = v83
                                                            }
                                                        }
                                                        var v84 uint8 = 131
                                                        _ = fi75
                                                        _ = v84
                                                    }
                                                    for fi85 := 0; fi85 < 9; fi85++ {
                                                        var v86 uint = 659
                                                        v2 += float64((0.0 * ((0.0 + (float32(float64(0.1)) * (0.1 - 0.0))) + (3.14159 * 0.1))))
                                                        if false {
                                                            if (v41 && false) {
                                                                if (false && v41) {
                                                                    var v87 string = "foo bar"
                                                                    _ = v87
                                                                } else {
                                                                    v88 := GenS5{F0: ("a" + "hello"), F1: v3}
                                                                    var v89 bool = ((127 > 188) || (((0.0 - float32(v2)) < 3.14159) && (!(1e10 < (float32(0.0) + -2.25)))))
                                                                    for fi90 := 0; fi90 < 18; fi90++ {
                                                                        fmt.Println(v1, v89, v4)
                                                                        switch ((("\n\t" + "foo bar") + (("" + (("foo bar" + "hello") + "\n\t")) + "hello")) + "") {
                                                                        case "":
                                                                            if true {
                                                                                v91 := GenS2{F0: (fi90 ^ v47), F1: 33}
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v92 uint32 = 110
                                                                                    _ = v92
                                                                                case "hello":
                                                                                    v89 = true
                                                                                case "\n\t":
                                                                                    var v93 uint32 = 2147483647
                                                                                    _ = v93
                                                                                }
                                                                                _ = len(v5)
                                                                                fmt.Println(v1, v1)
                                                                                fmt.Println(v1, v65)
                                                                                _ = v91
                                                                            } else {
                                                                                if ((v68 > 1) || ((v65 + 255) == 4294967295)) {
                                                                                    v86 -= v86
                                                                                    fmt.Println(v3, v72)
                                                                                }
                                                                                var v94 uint16 = 337
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = 4294967295
                                                                                }
                                                                                _ = v94
                                                                            }
                                                                        case "hello":
                                                                            switch ("a" + "") {
                                                                            case "foo bar":
                                                                                var v95 int64 = v3
                                                                                _ = v95
                                                                            case "a":
                                                                                if (v89 && v89) {
                                                                                    v4 -= 387
                                                                                    v2 += -2.25
                                                                                } else {
                                                                                    _ = len(v5)
                                                                                    v96 := GenS2{F0: v47, F1: 65535}
                                                                                    _ = v96
                                                                                }
                                                                            default:
                                                                                _ = len(v5)
                                                                            }
                                                                        default:
                                                                            v2 -= (v2 + (v2 * v2))
                                                                        }
                                                                        v4 = (v4 << ((v65) & 63))
                                                                        _ = fi90
                                                                    }
                                                                    _ = v88
                                                                    _ = v89
                                                                }
                                                                v5 = append(v5, 2147483648)
                                                                for fi97 := 0; fi97 < 13; fi97++ {
                                                                    var v98 uint = v72
                                                                    v47 = v47
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        var v99 int = (-(v47))
                                                                        _ = v99
                                                                    default:
                                                                        v1 -= (^(v1))
                                                                    }
                                                                    v1 = (v1 + (^(v1)))
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "\n\t":
                                                                        v41 = ((!false) || v41)
                                                                    default:
                                                                        var v100 uint = v72
                                                                        _ = v100
                                                                    }
                                                                    v2 = v2
                                                                    v101 := GenS6{F0: 605, F1: uint32(v86)}
                                                                    switch fi85 {
                                                                    case -1:
                                                                        v2 -= (v2 + v2)
                                                                    case 838:
                                                                        var v102 float32 = 0.0
                                                                        _ = v102
                                                                    }
                                                                    _ = fi97
                                                                    _ = v101
                                                                    _ = v98
                                                                }
                                                            } else {
                                                                var v103 bool = v41
                                                                if (!v41) {
                                                                    if (130 >= 42) {
                                                                        var v104 uint16 = 1
                                                                        for fi105 := 0; fi105 < 2; fi105++ {
                                                                            switch "" {
                                                                            case "":
                                                                                v106 := GenS7{F0: v4, F1: -2147483648}
                                                                                _ = v106
                                                                            case "\n\t":
                                                                                if v103 {
                                                                                    switch v47 {
                                                                                    case 128:
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    case 4294967295:
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = 844
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        switch fi85 {
                                                                                        case 128:
                                                                                            fmt.Println(v47, v2)
                                                                                        case 136:
                                                                                            var v107 int16 = -32768
                                                                                            _ = v107
                                                                                        case -277:
                                                                                            v41 = ((v41 && true) || (32768 >= 918))
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            v104 = (v104 & v104)
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = 1
                                                                                            }
                                                                                        } else {
                                                                                            var v108 uint32 = 2
                                                                                            var v109 int64 = -9223372036854775808
                                                                                            _ = v108
                                                                                            _ = v109
                                                                                        }
                                                                                    }
                                                                                    if (1 == 1) {
                                                                                        var v110 float32 = 0.1
                                                                                        _ = v110
                                                                                    } else {
                                                                                        var v111 bool = v41
                                                                                        if (!false) {
                                                                                            v112 := GenS8{F0: (v68 - v68), F1: v3}
                                                                                            switch ("" + "\n\t") {
                                                                                            case "foo bar":
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = uint32(v104)
                                                                                                }
                                                                                            case "a":
                                                                                                if (0 >= 42) {
                                                                                                    if (v103 || v103) {
                                                                                                        if v103 {
                                                                                                            v113 := GenS9{F0: (v47 - (fi105 % ((-1) | 1))), F1: false, F2: v104}
                                                                                                            v86 = (^(v86))
                                                                                                            _ = v113
                                                                                                        } else {
                                                                                                            v114 := GenS10{F0: "", F1: 1}
                                                                                                            fmt.Println(v86, v47, v86)
                                                                                                            if v103 {
                                                                                                                var v115 uint8 = 42
                                                                                                                v65 = 201
                                                                                                                v116 := GenS11{F0: (false || ("" > ""))}
                                                                                                                _ = v115
                                                                                                                _ = v116
                                                                                                            } else {
                                                                                                                v3 -= v3
                                                                                                                var v117 float32 = 3.14159
                                                                                                                _ = v117
                                                                                                            }
                                                                                                            switch -653 {
                                                                                                            case -963:
                                                                                                                switch "\n\t" {
                                                                                                                case "a":
                                                                                                                    v3 -= (v3 + (v3 | (v3 - v3)))
                                                                                                                case "foo bar":
                                                                                                                    var v118 uint8 = 128
                                                                                                                    _ = v118
                                                                                                                case "":
                                                                                                                    v4 -= (v68 % ((771) | 1))
                                                                                                                }
                                                                                                            case 127:
                                                                                                                var v119 string = ""
                                                                                                                _ = v119
                                                                                                            case -902:
                                                                                                                if ((v72 ^ v65) != (v86 >> ((v65) & 63))) {
                                                                                                                    v120 := GenS1{F0: false, F1: 14}
                                                                                                                    v1 -= 528
                                                                                                                    switch (^(v47)) {
                                                                                                                    case 42:
                                                                                                                        if (v1 > 42) {
                                                                                                                            if ((-(v1)) < (v1 & (^(v1)))) {
                                                                                                                                var v121 int32 = (v1 >> ((v86) & 63))
                                                                                                                                var v122 uint64 = v68
                                                                                                                                fmt.Println(fi105, v122)
                                                                                                                                _ = v121
                                                                                                                                _ = v122
                                                                                                                            }
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                            var v123 int = 9223372036854775807
                                                                                                                            _ = v123
                                                                                                                        }
                                                                                                                    case 127:
                                                                                                                        if false {
                                                                                                                            switch (fi105 << ((18446744073709551615) & 63)) {
                                                                                                                            case -1:
                                                                                                                                fmt.Println(v2)
                                                                                                                            }
                                                                                                                            fmt.Println(v3, v65)
                                                                                                                            var v124 float32 = 1.5
                                                                                                                            var v125 bool = ((((fi105 >> (((v72 + v72)) & 63)) >= 1) && (v41 && v111)) || ((1 < 2147483647) || (42 >= 160)))
                                                                                                                            _ = v124
                                                                                                                            _ = v125
                                                                                                                        } else {
                                                                                                                            fmt.Println(v65, v47)
                                                                                                                        }
                                                                                                                    case 32767:
                                                                                                                        if v111 {
                                                                                                                            fmt.Println(v103)
                                                                                                                            if (42 != 255) {
                                                                                                                                v126 := GenS10{F0: "a", F1: -8}
                                                                                                                                if len(v5) > 0 {
                                                                                                                                    v5[0] = 0
                                                                                                                                }
                                                                                                                                switch 9223372036854775807 {
                                                                                                                                case 0:
                                                                                                                                    var v127 int8 = 1
                                                                                                                                    _ = v127
                                                                                                                                default:
                                                                                                                                    _ = len(v5)
                                                                                                                                }
                                                                                                                                switch 9223372036854775807 {
                                                                                                                                case -2147483648:
                                                                                                                                    v111 = (!v41)
                                                                                                                                case 1:
                                                                                                                                    v128 := GenS12{F0: (("foo bar" + "foo bar") + "a")}
                                                                                                                                    _ = v128
                                                                                                                                }
                                                                                                                                fmt.Println(v4, v111)
                                                                                                                                _ = v126
                                                                                                                            } else {
                                                                                                                                if ((("hello" + ("hello" + "a")) + ("hello" + "")) != "hello") {
                                                                                                                                    if true {
                                                                                                                                        v47 = (v47 >> (((v72 ^ (v65 << (((v72 - (v65 & (v65 + v72)))) & 63)))) & 63))
                                                                                                                                    }
                                                                                                                                    fmt.Println(v41, v72)
                                                                                                                                    v129 := GenS13{F0: -446}
                                                                                                                                    v111 = true
                                                                                                                                    _ = v129
                                                                                                                                }
                                                                                                                                switch "hello" {
                                                                                                                                case "foo bar":
                                                                                                                                    v130 := GenS12{F0: "a"}
                                                                                                                                    _ = v130
                                                                                                                                case "\n\t":
                                                                                                                                    v2 += v2
                                                                                                                                case "":
                                                                                                                                    switch ("\n\t" + "foo bar") {
                                                                                                                                    case "foo bar":
                                                                                                                                        v5 = append(v5, 1)
                                                                                                                                    case "\n\t":
                                                                                                                                        fmt.Println(v104)
                                                                                                                                    default:
                                                                                                                                        var v131 uint8 = 39
                                                                                                                                        _ = v131
                                                                                                                                    }
                                                                                                                                    v68 += v4
                                                                                                                                default:
                                                                                                                                    fmt.Println(v41)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    _ = v120
                                                                                                                } else {
                                                                                                                    v5 = append(v5, 42)
                                                                                                                    var v132 int = fi105
                                                                                                                    v133 := GenS5{F0: ("hello" + (("a" + "\n\t") + ("hello" + "foo bar"))), F1: (v3 * -765)}
                                                                                                                    fmt.Println(fi105)
                                                                                                                    _ = v132
                                                                                                                    _ = v133
                                                                                                                }
                                                                                                            }
                                                                                                            v5 = append(v5, 65535)
                                                                                                            var v134 string = "\n\t"
                                                                                                            _ = v114
                                                                                                            _ = v134
                                                                                                        }
                                                                                                        var v135 uint16 = v104
                                                                                                        v47 = v47
                                                                                                        _ = v135
                                                                                                    } else {
                                                                                                        var v136 int = v47
                                                                                                        _ = len(v5)
                                                                                                        _ = v136
                                                                                                    }
                                                                                                    var v137 int64 = (^(v3))
                                                                                                    _ = v137
                                                                                                } else {
                                                                                                    v138 := GenS1{F0: (!true), F1: 42}
                                                                                                    if (v2 <= ((v2 - (float64(float32(v2)) + (1e10 + (v2 * v2)))) * v2)) {
                                                                                                        v65 = uint(v104)
                                                                                                    }
                                                                                                    _ = v138
                                                                                                }
                                                                                            }
                                                                                            _ = v112
                                                                                        } else {
                                                                                            v5 = append(v5, 42)
                                                                                        }
                                                                                        if v103 {
                                                                                            var v139 int64 = v3
                                                                                            fmt.Println(v47)
                                                                                            _ = v139
                                                                                        }
                                                                                        _ = v111
                                                                                    }
                                                                                    var v140 int16 = -32768
                                                                                    _ = v140
                                                                                } else {
                                                                                    v72 = v65
                                                                                }
                                                                            case "hello":
                                                                                var v141 uint = v86
                                                                                var v142 uint16 = (v104 ^ 673)
                                                                                if (2 == v68) {
                                                                                    if ((v68 % (((v68 ^ (^(v68)))) | 1)) >= v68) {
                                                                                        if true {
                                                                                            var v143 uint8 = 255
                                                                                            var v144 uint16 = (v104 * v104)
                                                                                            v41 = v103
                                                                                            _ = v143
                                                                                            _ = v144
                                                                                        }
                                                                                        v5 = append(v5, 65535)
                                                                                        var v145 int32 = (^(v1))
                                                                                        _ = v145
                                                                                    }
                                                                                    v141 = (v72 ^ v72)
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                } else {
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi85, v103)
                                                                                    case "":
                                                                                        var v146 int64 = (v3 / (((v3 & -1)) | 1))
                                                                                        _ = v146
                                                                                    default:
                                                                                        if ((-(v3)) >= (v3 % ((458) | 1))) {
                                                                                            v68 -= (^(v4))
                                                                                            v1 += 2147483647
                                                                                        } else {
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            v86 += (v72 * (v72 ^ v65))
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    var v147 int = 0
                                                                                    _ = v147
                                                                                }
                                                                                _ = v141
                                                                                _ = v142
                                                                            }
                                                                            _ = fi105
                                                                        }
                                                                        _ = v104
                                                                    } else {
                                                                        if false {
                                                                            fmt.Println(v65)
                                                                            if (v103 && false) {
                                                                                switch ("\n\t" + "foo bar") {
                                                                                case "hello":
                                                                                    if false {
                                                                                        for fi148 := 0; fi148 < 6; fi148++ {
                                                                                            var v149 int32 = (^(v1))
                                                                                            fmt.Println(v65, fi148, v86)
                                                                                            _ = fi148
                                                                                            _ = v149
                                                                                        }
                                                                                        for fi150 := 0; fi150 < 3; fi150++ {
                                                                                            var v151 int16 = -428
                                                                                            v65 -= (^(v65))
                                                                                            _ = fi150
                                                                                            _ = v151
                                                                                        }
                                                                                        var v152 int = 42
                                                                                        _ = v152
                                                                                    } else {
                                                                                        switch fi85 {
                                                                                        case 761:
                                                                                            fmt.Println(v1, fi85)
                                                                                        case 4294967295:
                                                                                            v5 = append(v5, 0)
                                                                                        case -86:
                                                                                            var v153 int8 = 100
                                                                                            _ = v153
                                                                                        default:
                                                                                            fmt.Println(fi85, v65)
                                                                                        }
                                                                                    }
                                                                                case "foo bar":
                                                                                    v154 := GenS14{F0: 0}
                                                                                    _ = v154
                                                                                default:
                                                                                    var v155 uint16 = 2
                                                                                    _ = v155
                                                                                }
                                                                                var v156 bool = v41
                                                                                var v157 float64 = v2
                                                                                _ = v156
                                                                                _ = v157
                                                                            } else {
                                                                                v158 := GenS15{F0: 0}
                                                                                _ = v158
                                                                            }
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        } else {
                                                                            var v159 uint = 255
                                                                            if ((v4 ^ 18446744073709551615) < (v68 % (((v4 + (^(v68)))) | 1))) {
                                                                                var v160 int8 = 0
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                _ = v160
                                                                            } else {
                                                                                v161 := GenS16{F0: 32767, F1: false}
                                                                                var v162 uint32 = 4294967295
                                                                                _ = v161
                                                                                _ = v162
                                                                            }
                                                                            _ = v159
                                                                        }
                                                                        fmt.Println(v3)
                                                                    }
                                                                    for fi163 := 0; fi163 < 3; fi163++ {
                                                                        var v164 int32 = (v1 ^ (^(v1)))
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            var v165 uint = (v86 * (v72 >> ((985) & 63)))
                                                                            _ = v165
                                                                        }
                                                                        if false {
                                                                            fmt.Println(v72, v3, v164)
                                                                            var v166 float32 = (((1e10 - (3.14159 * 1.5)) * (-2.25 + 1e10)) * float32(0.1))
                                                                            _ = v166
                                                                        }
                                                                        v1 -= v1
                                                                        _ = fi163
                                                                        _ = v164
                                                                    }
                                                                } else {
                                                                    v86 -= 18446744073709551615
                                                                    v86 += (^(v65))
                                                                }
                                                                switch "" {
                                                                case "hello":
                                                                    for fi167 := 0; fi167 < 1; fi167++ {
                                                                        if v103 {
                                                                            var v168 bool = ((v68 ^ (^(v4))) >= (v4 & v4))
                                                                            if v168 {
                                                                                v169 := GenS14{F0: 255}
                                                                                _ = v169
                                                                            }
                                                                            _ = v168
                                                                        } else {
                                                                            if v41 {
                                                                                if (65535 > 32767) {
                                                                                    fmt.Println(v86, v41, v2)
                                                                                } else {
                                                                                    var v170 bool = (42 != -269)
                                                                                    _ = v170
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v103, v103)
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                                switch fi167 {
                                                                                case -743:
                                                                                    v5 = append(v5, 65535)
                                                                                    var v171 uint64 = 486
                                                                                    _ = len(v5)
                                                                                    fmt.Println(fi167)
                                                                                    _ = v171
                                                                                case 127:
                                                                                    v65 = (v72 / (((v86 ^ (v86 & v86))) | 1))
                                                                                case -1:
                                                                                    v5 = append(v5, 0)
                                                                                }
                                                                            }
                                                                            v68 = v4
                                                                            if (32767 > 509) {
                                                                                var v172 int8 = 0
                                                                                _ = v172
                                                                            } else {
                                                                                var v173 float64 = v2
                                                                                _ = v173
                                                                            }
                                                                            v5 = append(v5, 2147483648)
                                                                        }
                                                                        fmt.Println(v103, v41)
                                                                        _ = fi167
                                                                    }
                                                                case "foo bar":
                                                                    _ = len(v5)
                                                                case "":
                                                                    v174 := GenS17{F0: int32(v3)}
                                                                    _ = v174
                                                                }
                                                                _ = v103
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            v72 -= 2
                                                        }
                                                        _ = fi85
                                                        _ = v86
                                                    }
                                                    if v41 {
                                                        if (("" > "hello") && (0.1 <= (0.0 * 0.0))) {
                                                            if (v41 && false) {
                                                                switch ("hello" + (("foo bar" + "a") + "")) {
                                                                case "":
                                                                    for fi175 := 0; fi175 < 18; fi175++ {
                                                                        v1 -= v1
                                                                        _ = fi175
                                                                    }
                                                                case "\n\t":
                                                                    v4 = (v4 << ((255) & 63))
                                                                case "hello":
                                                                    v176 := GenS13{F0: -777}
                                                                    _ = v176
                                                                }
                                                            }
                                                            for fi177 := 0; fi177 < 5; fi177++ {
                                                                var v178 uint64 = (v4 >> (((v72 % ((v72) | 1))) & 63))
                                                                _ = fi177
                                                                _ = v178
                                                            }
                                                            var v179 string = "a"
                                                            _ = v179
                                                        }
                                                        v1 = v1
                                                    }
                                                }
                                                v2 = 1e10
                                                var v180 float32 = float32(v2)
                                                _ = v180
                                            } else {
                                                v3 = v3
                                                var v181 float64 = v2
                                                fmt.Println(v65, v1)
                                                _ = v181
                                            }
                                        }
                                    }
                                case 118:
                                    switch "a" {
                                    case "a":
                                        switch "hello" {
                                        case "foo bar":
                                            var v182 uint64 = (v68 / ((321) | 1))
                                            _ = v182
                                        default:
                                            for fi183 := 0; fi183 < 19; fi183++ {
                                                v184 := GenS6{F0: 2147483647, F1: 1}
                                                for fi185 := 0; fi185 < 19; fi185++ {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = fi185
                                                }
                                                var v186 float32 = 0.1
                                                _ = fi183
                                                _ = v184
                                                _ = v186
                                            }
                                        }
                                    default:
                                        if (27 == 127) {
                                            switch 127 {
                                            case 2147483647:
                                                var v187 int = (^(v47))
                                                _ = v187
                                            }
                                            for fi188 := 0; fi188 < 7; fi188++ {
                                                v72 += (v65 >> ((v72) & 63))
                                                _ = fi188
                                            }
                                        }
                                    }
                                }
                                _ = v72
                            }
                        }
                        _ = len(v5)
                        _ = v68
                    }
                    fmt.Println(v65, v2, v41)
                }
                _ = v65
            }
        case "":
            switch v47 {
            case -17:
                if true {
                    var v189 uint16 = 508
                    for fi190 := 0; fi190 < 7; fi190++ {
                        var v191 uint32 = 958
                        v192 := GenS7{F0: v4, F1: (fi190 + fi190)}
                        for fi193 := 0; fi193 < 4; fi193++ {
                            v194 := GenS5{F0: "foo bar", F1: -9223372036854775808}
                            _ = fi193
                            _ = v194
                        }
                        _ = fi190
                        _ = v191
                        _ = v192
                    }
                    _ = v189
                }
            case 1:
                v4 -= v4
            case 32767:
                v5 = append(v5, 2147483647)
            }
        default:
            var v195 int = 42
            _ = v195
        }
    } else {
        _ = len(v5)
    }
    v196 := GenS8{F0: 831, F1: 0}
    v5 = append(v5, 65535)
    switch 127 {
    case -822:
        for fi197 := 0; fi197 < 19; fi197++ {
            var v198 bool = v41
            v47 = (-(fi197))
            _ = fi197
            _ = v198
        }
    case -854:
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    v4 = (^(v4))
    var v199 string = "a"
    if true {
        v47 = (v47 & (v47 >> ((509) & 63)))
    }
    var v200 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v199:", v199)
    fmt.Println("v2:", v2)
    fmt.Println("v200:", v200)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v196
    _ = v199
    _ = v2
    _ = v200
    _ = v3
    _ = v4
    _ = v41
    _ = v47
    _ = v5
}
