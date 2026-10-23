package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int; F1 int; F2 int16 }
type GenS4 struct { F0 uint8; F1 int; F2 uint }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int16; F1 uint32; F2 int16 }
type GenS7 struct { F0 int64; F1 int32; F2 int16 }
type GenS8 struct { F0 string; F1 int8 }
type GenS9 struct { F0 bool; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 int64; F1 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int8; F1 bool; F2 uint16 }
type GenS15 struct { F0 int8; F1 uint64 }
type GenS16 struct { F0 uint8; F1 int8 }
type GenS17 struct { F0 int16; F1 int8 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 uint64 }

func main() {
    var v1 int64 = -1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = (1.5 + ((((0.0 / 1.0) + 3.14159) + (0.0 / 0.1)) * float32(1e10)))
    v5 := make(map[int]int)
    var v6 uint32 = 560
    var v7 uint = 1
    if false {
        v8 := GenS1{F0: (v1 ^ v1)}
        switch "hello" {
        case "a":
            if false {
                switch ("a" + "hello") {
                case "foo bar":
                    v7 = 18446744073709551615
                case "hello":
                    if true {
                        for fi9 := 0; fi9 < 5; fi9++ {
                            v10 := GenS2{F0: 42, F1: 2147483648, F2: v6}
                            for fi11 := 0; fi11 < 3; fi11++ {
                                fmt.Println(v7, v6, v1)
                                _ = fi11
                            }
                            _ = fi9
                            _ = v10
                        }
                        delete(v5, 866)
                        var v12 string = "foo bar"
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            v14 := GenS2{F0: 1, F1: (v6 / ((v6) | 1)), F2: (v6 % ((v6) | 1))}
                            _ = fi13
                            _ = v14
                        }
                        if (32767 > 255) {
                            if ((true && false) || false) {
                                if true {
                                    if true {
                                        v15 := GenS3{F0: -9223372036854775808, F1: -1, F2: 423}
                                        switch 397 {
                                        case 32768:
                                            var v16 int8 = 127
                                            _ = v16
                                        case 4294967295:
                                            switch 4294967295 {
                                            case 2147483647:
                                                v17 := []int{2147483647}
                                                _ = v17
                                            case 53:
                                                var v18 uint = v7
                                                _ = v18
                                            }
                                        default:
                                            var v19 uint8 = 154
                                            _ = v19
                                        }
                                        for fi20 := 0; fi20 < 17; fi20++ {
                                            v1 = v1
                                            if (0 == 267) {
                                                v6 += (v6 % (((v6 & (v6 / ((1) | 1)))) | 1))
                                                fmt.Println(v1, fi20, v1)
                                                switch (^(fi20)) {
                                                case 1:
                                                    v21 := GenS4{F0: 42, F1: (fi20 & fi20), F2: (v7 | v7)}
                                                    _ = v21
                                                case -1:
                                                    v7 = v7
                                                case 127:
                                                    v7 = v7
                                                default:
                                                    v22 := GenS4{F0: 81, F1: fi20, F2: 255}
                                                    _ = v22
                                                }
                                            } else {
                                                if true {
                                                    delete(v5, 0)
                                                    if ((v1 | v1) <= v1) {
                                                        v23 := GenS1{F0: (v1 - v1)}
                                                        for fi24 := 0; fi24 < 20; fi24++ {
                                                            var v25 int64 = v1
                                                            v1 = v1
                                                            _ = fi24
                                                            _ = v25
                                                        }
                                                        _ = v23
                                                    }
                                                } else {
                                                    switch "hello" {
                                                    case "hello":
                                                        for fi26 := 0; fi26 < 16; fi26++ {
                                                            switch ("a" + "hello") {
                                                            case "a":
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    if false {
                                                                        var v27 float32 = 3.14159
                                                                        _ = v27
                                                                    }
                                                                case "":
                                                                    if true {
                                                                        var v28 int8 = 127
                                                                        _ = v28
                                                                    }
                                                                case "foo bar":
                                                                    v29 := []int{}
                                                                    _ = v29
                                                                }
                                                            case "":
                                                                if false {
                                                                    v7 = (v7 - v7)
                                                                }
                                                                if false {
                                                                    v6 = v6
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        v1 = 42
                                                                    case "hello":
                                                                        v30 := GenS5{F0: 0, F1: -655}
                                                                        _ = v30
                                                                    case "":
                                                                        v31 := v3[38]
                                                                        _ = v31
                                                                    default:
                                                                        var v32 string = "foo bar"
                                                                        _ = v32
                                                                    }
                                                                    switch ("foo bar" + "hello") {
                                                                    case "a":
                                                                        switch fi26 {
                                                                        case -1:
                                                                            delete(v2, "foo bar")
                                                                        }
                                                                    case "":
                                                                        v1 += (^(v1))
                                                                    case "\n\t":
                                                                        v33 := GenS3{F0: 4294967295, F1: fi20, F2: -32768}
                                                                        _ = v33
                                                                    default:
                                                                        switch (-(fi20)) {
                                                                        case 693:
                                                                            v34 := GenS3{F0: 1, F1: fi20, F2: 0}
                                                                            _ = v34
                                                                        default:
                                                                            v35 := GenS2{F0: 42, F1: 2147483647, F2: (v6 | v6)}
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                }
                                                                var v36 bool = true
                                                                if ((-2.25 + ((((float64(v4) + 0.0) + 0.0) - 1.5) / 1e10)) <= (0.0 + ((0.0 + float64(v4)) + 1.5))) {
                                                                    v37 := []int{}
                                                                    _ = v37
                                                                }
                                                                fmt.Println(v7, v4)
                                                                _ = v36
                                                            default:
                                                                v4 = v4
                                                            }
                                                            fmt.Println(v6, v4, fi26)
                                                            v1 += (v1 % (((v1 - -9223372036854775808)) | 1))
                                                            _ = fi26
                                                        }
                                                    default:
                                                        v38 := GenS5{F0: 0, F1: 1}
                                                        _ = v38
                                                    }
                                                    v39 := GenS6{F0: 1000, F1: v6, F2: -1}
                                                    switch ("a" + (("\n\t" + "hello") + ("" + ((("a" + (("a" + "hello") + ("a" + ("foo bar" + "\n\t")))) + ("hello" + "\n\t")) + (("\n\t" + "a") + ("\n\t" + "\n\t")))))) {
                                                    case "":
                                                        fmt.Println(v1, v1)
                                                    default:
                                                        v40 := v5[(fi20 / ((fi20) | 1))]
                                                        _ = v40
                                                    }
                                                    _ = v39
                                                }
                                            }
                                            _ = fi20
                                        }
                                        _ = v15
                                    } else {
                                        if true {
                                            v7 -= v7
                                            var v41 uint64 = 42
                                            _ = v41
                                        }
                                    }
                                } else {
                                    v7 = v7
                                    fmt.Println(v4, v4)
                                }
                            } else {
                                v42 := []int{32767}
                                v1 = v1
                                _ = v42
                            }
                            var v43 uint = 4294967295
                            _ = v43
                        }
                    }
                default:
                    var v44 int8 = 127
                    _ = v44
                }
            } else {
                _ = len(v5)
                var v45 uint = (v7 - (v7 + v7))
                _ = v45
            }
        case "\n\t":
            for fi46 := 0; fi46 < 1; fi46++ {
                fmt.Println(fi46)
                _ = fi46
            }
        case "":
            switch ("\n\t" + ("" + "\n\t")) {
            case "a":
                v47 := GenS6{F0: 32767, F1: (v6 << (((v7 << ((v7) & 63))) & 63)), F2: 0}
                _ = v47
            case "hello":
                for fi48 := 0; fi48 < 18; fi48++ {
                    for fi49 := 0; fi49 < 5; fi49++ {
                        fmt.Println(v6, v7)
                        _ = fi49
                    }
                    var v50 uint16 = 658
                    v51 := GenS7{F0: (v1 ^ (v1 * v1)), F1: 42, F2: 42}
                    _ = fi48
                    _ = v50
                    _ = v51
                }
            case "\n\t":
                if (!true) {
                    if (float32(1.5) >= (float32(1.5) / ((v4 - v4) / v4))) {
                        v52 := GenS7{F0: v1, F1: 1, F2: 1000}
                        v6 = 869
                        _ = v52
                    }
                    if ((v6 * 2147483648) == v6) {
                        for fi53 := 0; fi53 < 5; fi53++ {
                            v4 = (v4 + v4)
                            if (false || ((-(fi53)) == fi53)) {
                                for fi54 := 0; fi54 < 3; fi54++ {
                                    if true {
                                        v1 = -9223372036854775808
                                    }
                                    if true {
                                        if (-32768 < 1000) {
                                            var v55 int16 = 1000
                                            if true {
                                                var v56 uint16 = 278
                                                var v57 float64 = 1e10
                                                _ = v56
                                                _ = v57
                                            }
                                            if (-2147483648 > 42) {
                                                v55 = int16(v1)
                                                var v58 int8 = 0
                                                v6 -= (v6 * (^(v6)))
                                                _ = v58
                                            } else {
                                                var v59 uint32 = (v6 << ((v7) & 63))
                                                var v60 int16 = (v55 + 32767)
                                                _ = v59
                                                _ = v60
                                            }
                                            _ = v55
                                        } else {
                                            v61 := []int{}
                                            _ = v61
                                        }
                                        if ("hello" == ("hello" + "")) {
                                            var v62 float64 = (-2.25 / 3.14159)
                                            _ = v62
                                        }
                                        var v63 uint = 408
                                        if ((v6 & 170) > 1) {
                                            if false {
                                                var v64 uint = v63
                                                v65 := GenS2{F0: 4294967295, F1: (v6 ^ (v6 % ((536) | 1))), F2: 716}
                                                var v66 uint16 = 0
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            } else {
                                                var v67 int32 = -35
                                                v68 := []int{4294967295, 1}
                                                v67 -= (v67 % (((v67 ^ (v67 * -1))) | 1))
                                                if (v7 != 255) {
                                                    var v69 int16 = 1000
                                                    v7 += 170
                                                    if len(v68) > 0 {
                                                        v68[0] = (-(fi53))
                                                    }
                                                    _ = v69
                                                } else {
                                                    switch ("hello" + "\n\t") {
                                                    case "hello":
                                                        var v70 int64 = -741
                                                        _ = v70
                                                    case "a":
                                                        v71 := GenS8{F0: "a", F1: -1}
                                                        _ = v71
                                                    case "":
                                                        fmt.Println(v63, v4, fi54)
                                                    default:
                                                        var v72 string = "hello"
                                                        _ = v72
                                                    }
                                                    var v73 string = "\n\t"
                                                    if true {
                                                        var v74 bool = false
                                                        fmt.Println(v7, v67, fi54)
                                                        fmt.Println(fi53, v1, v63)
                                                        _ = v74
                                                    }
                                                    _ = v73
                                                }
                                                if false {
                                                    fmt.Println(v6)
                                                    v63 -= (v63 >> ((220) & 63))
                                                } else {
                                                    fmt.Println(v1, v6, v6)
                                                    v6 = 65535
                                                    var v75 uint = v7
                                                    var v76 uint8 = 127
                                                    _ = v75
                                                    _ = v76
                                                }
                                                _ = v67
                                                _ = v68
                                            }
                                            var v77 uint = v7
                                            _ = v77
                                        }
                                        _ = v63
                                    } else {
                                        var v78 int32 = 42
                                        v78 = 2147483647
                                        if false {
                                            fmt.Println(fi54, v6)
                                            fmt.Println(fi53)
                                        }
                                        _ = v78
                                    }
                                    _ = fi54
                                }
                                v79 := GenS3{F0: fi53, F1: (fi53 << ((782) & 63)), F2: 0}
                                _ = v79
                            } else {
                                switch fi53 {
                                case 32768:
                                    for fi80 := 0; fi80 < 6; fi80++ {
                                        if ((2 != 42) && (65535 > 363)) {
                                            v81 := GenS9{F0: false, F1: 42, F2: 2}
                                            v82 := GenS10{F0: v6}
                                            if true {
                                                if false {
                                                    v83 := []int{(fi53 ^ (^(fi53)))}
                                                    switch fi53 {
                                                    case 0:
                                                        switch fi53 {
                                                        case -1:
                                                            if true {
                                                                var v84 uint8 = 0
                                                                _ = v84
                                                            }
                                                        }
                                                    case 32767:
                                                        if ("a" == "") {
                                                            v85 := GenS3{F0: 601, F1: fi80, F2: 0}
                                                            _ = v85
                                                        }
                                                    default:
                                                        if true {
                                                            fmt.Println(v4)
                                                            var v86 uint64 = 42
                                                            var v87 uint8 = 40
                                                            _ = v86
                                                            _ = v87
                                                        } else {
                                                            fmt.Println(fi53, v1, v7)
                                                            v7 -= 316
                                                            switch (fi80 + fi80) {
                                                            case -39:
                                                                if true {
                                                                    v4 = ((1.5 * v4) / v4)
                                                                    if (255 < 148) {
                                                                        fmt.Println(v7, v4, v4)
                                                                        v83 = append(v83, 9223372036854775807)
                                                                        if true {
                                                                            fmt.Println(v6, fi80, v4)
                                                                            if (!true) {
                                                                                v88 := v2[("foo bar" + "hello")]
                                                                                var v89 int64 = v1
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch ("hello" + ("foo bar" + ((("" + "\n\t") + "") + "foo bar"))) {
                                                                        case "":
                                                                            var v90 uint = v7
                                                                            _ = v90
                                                                        default:
                                                                            switch fi80 {
                                                                            case 1:
                                                                                if len(v83) >= 2 {
                                                                                    v83 = v83[1:len(v83)]
                                                                                }
                                                                            case -9223372036854775808:
                                                                                var v91 int = -645
                                                                                if ((v1 > (v1 | (-(v1)))) || (v91 > (fi53 + fi80))) {
                                                                                    fmt.Println(v7, fi80, v1)
                                                                                }
                                                                                _ = v91
                                                                            case 9223372036854775807:
                                                                                _ = len(v83)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            case -28:
                                                                v6 += v6
                                                            }
                                                        }
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v92 := GenS1{F0: (v1 << ((v7) & 63))}
                                                            _ = v92
                                                        case "hello":
                                                            _ = len(v3)
                                                        case "a":
                                                            v93 := GenS6{F0: 42, F1: v6, F2: 1000}
                                                            _ = v93
                                                        default:
                                                            var v94 uint8 = 1
                                                            _ = v94
                                                        }
                                                        var v95 int32 = 47
                                                        if false {
                                                            if (18446744073709551615 >= 4294967295) {
                                                                if false {
                                                                    var v96 uint = v7
                                                                    fmt.Println(fi53)
                                                                    _ = v96
                                                                } else {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(fi80)
                                                                    default:
                                                                        if (-26 <= 89) {
                                                                            if len(v83) > 0 {
                                                                                v83[0] = fi80
                                                                            }
                                                                        }
                                                                    }
                                                                    if (true && true) {
                                                                        fmt.Println(fi80)
                                                                        v97 := GenS4{F0: 133, F1: 9223372036854775807, F2: v7}
                                                                        _ = v97
                                                                    } else {
                                                                        v6 = v6
                                                                        if (!(true || true)) {
                                                                            fmt.Println(v4)
                                                                        }
                                                                        v7 = (v7 | (v7 / ((v7) | 1)))
                                                                    }
                                                                    v4 = 1.5
                                                                }
                                                            } else {
                                                                v83 = append(v83, 127)
                                                                if true {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        _ = len(v5)
                                                                    case "a":
                                                                        if false {
                                                                            var v98 float64 = float64((0.1 - (((1.5 * (((3.14159 * v4) + ((1.5 - -2.25) / float32(1.5))) - (float32((1e10 * (0.1 + float64(1.5)))) * -2.25))) + ((v4 * (float32(1.5) * (v4 * float32((3.14159 / (0.0 * float64(v4))))))) / v4)) * float32(1.5))))
                                                                            v4 += 0.0
                                                                            _ = v98
                                                                        }
                                                                    }
                                                                }
                                                                if ((v7 & v7) <= 0) {
                                                                    var v99 uint8 = 42
                                                                    if true {
                                                                        v100 := GenS1{F0: 872}
                                                                        v101 := GenS11{F0: ("\n\t" + ("\n\t" + ("a" + "foo bar"))), F1: 757}
                                                                        _ = v100
                                                                        _ = v101
                                                                    }
                                                                    v99 = uint8(v7)
                                                                    _ = v99
                                                                }
                                                            }
                                                            var v102 uint32 = 2147483648
                                                            delete(v5, fi80)
                                                            _ = v102
                                                        }
                                                        var v103 uint64 = 910
                                                        _ = v103
                                                        _ = v95
                                                    }
                                                    var v104 uint64 = 890
                                                    _ = v104
                                                    _ = v83
                                                } else {
                                                    v4 = 3.14159
                                                    v105 := GenS7{F0: 0, F1: 999, F2: 945}
                                                    _ = v105
                                                }
                                                v3[v6] = (fi80 << ((uint(v6)) & 63))
                                                v106 := GenS12{F0: (^(v1)), F1: int32(v1)}
                                                _ = v106
                                            } else {
                                                _ = len(v5)
                                                v107 := GenS13{F0: 126}
                                                fmt.Println(v1)
                                                _ = v107
                                            }
                                            _ = v81
                                            _ = v82
                                        }
                                        v7 = v7
                                        _ = fi80
                                    }
                                }
                                var v108 uint16 = 32767
                                fmt.Println(v108)
                                _ = v108
                            }
                            _ = fi53
                        }
                    } else {
                        for fi109 := 0; fi109 < 4; fi109++ {
                            v110 := GenS11{F0: ("foo bar" + ""), F1: 42}
                            _ = fi109
                            _ = v110
                        }
                    }
                    v111 := GenS8{F0: ("hello" + "foo bar"), F1: 127}
                    _ = v111
                } else {
                    for fi112 := 0; fi112 < 5; fi112++ {
                        var v113 float64 = 0.1
                        v113 += 0.1
                        _ = fi112
                        _ = v113
                    }
                    switch "a" {
                    case "hello":
                        fmt.Println(v7, v4)
                    case "\n\t":
                        if ((v6 & (v6 - 2147483648)) > v6) {
                            v7 += 1
                            v4 = (v4 / ((1e10 * v4) * v4))
                        } else {
                            v114 := GenS1{F0: -296}
                            for fi115 := 0; fi115 < 1; fi115++ {
                                v1 = (-(v1))
                                if (true || false) {
                                    fmt.Println(v4)
                                    for fi116 := 0; fi116 < 5; fi116++ {
                                        v6 = (v6 / ((v6) | 1))
                                        v6 -= 65535
                                        v7 -= 65535
                                        _ = fi116
                                    }
                                    switch fi115 {
                                    case -450:
                                        if false {
                                            v6 = (v6 + v6)
                                            v117 := GenS3{F0: (fi115 * int(v1)), F1: fi115, F2: -442}
                                            v1 -= v1
                                            _ = v117
                                        } else {
                                            if (("hello" + "foo bar") > "") {
                                                v6 = (v6 & v6)
                                                v118 := GenS14{F0: 40, F1: false, F2: 120}
                                                v119 := []int{fi115}
                                                _ = v118
                                                _ = v119
                                            }
                                        }
                                    case 128:
                                        if false {
                                            v120 := []int{}
                                            var v121 uint16 = 1
                                            _ = v120
                                            _ = v121
                                        }
                                    case 4294967295:
                                        delete(v5, fi115)
                                    }
                                } else {
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(fi115, v4)
                                    case "foo bar":
                                        if false {
                                            v6 -= v6
                                            v4 -= v4
                                            _ = len(v3)
                                        }
                                    default:
                                        fmt.Println(v6, fi115)
                                    }
                                    var v122 float64 = float64(1e10)
                                    _ = v122
                                }
                                var v123 int8 = 42
                                _ = fi115
                                _ = v123
                            }
                            var v124 float64 = 1e10
                            _ = v114
                            _ = v124
                        }
                    }
                }
            }
        default:
            v125 := GenS1{F0: 180}
            _ = v125
        }
        _ = v8
    }
    v126 := GenS10{F0: 2147483647}
    for fi127 := 0; fi127 < 20; fi127++ {
        switch (fi127 / (((fi127 << ((v7) & 63))) | 1)) {
        case 2147483647:
            if (fi127 >= (fi127 & fi127)) {
                var v128 int8 = 0
                _ = v128
            } else {
                v129 := GenS15{F0: 100, F1: 258}
                if (-348 > -221) {
                    v130 := []int{(fi127 - (fi127 | (^(fi127))))}
                    fmt.Println(fi127)
                    _ = v130
                } else {
                    _ = len(v2)
                }
                if false {
                    switch "hello" {
                    case "foo bar":
                        switch "\n\t" {
                        case "foo bar":
                            v131 := []int{}
                            _ = v131
                        case "\n\t":
                            for fi132 := 0; fi132 < 9; fi132++ {
                                if (false && (21 <= 32767)) {
                                    v133 := GenS4{F0: 1, F1: -170, F2: (v7 >> ((42) & 63))}
                                    v134 := v5[fi132]
                                    v1 -= v1
                                    _ = v133
                                    _ = v134
                                }
                                _ = fi132
                            }
                        case "a":
                            fmt.Println(v6)
                        }
                    case "hello":
                        fmt.Println(v7, fi127)
                    case "":
                        fmt.Println(v1, v1, v4)
                    default:
                        switch (-(fi127)) {
                        case 32768:
                            delete(v5, fi127)
                        default:
                            v135 := []int{fi127}
                            _ = v135
                        }
                    }
                }
                _ = v129
            }
        }
        v136 := []int{int(v1), (fi127 | (fi127 / (((fi127 + (^(fi127)))) | 1)))}
        _ = fi127
        _ = v136
    }
    for fi137 := 0; fi137 < 11; fi137++ {
        var v138 int64 = (v1 - (v1 % (((v1 << ((v7) & 63))) | 1)))
        var v139 int16 = 1000
        var v140 uint32 = v6
        _ = fi137
        _ = v138
        _ = v139
        _ = v140
    }
    if false {
        var v141 float64 = ((1e10 / float64(v4)) - 3.14159)
        v142 := GenS3{F0: -9223372036854775808, F1: 32768, F2: -511}
        _ = v141
        _ = v142
    } else {
        fmt.Println(v4)
        v7 = v7
    }
    switch 128 {
    case -1:
        if true {
            switch "" {
            case "":
                var v143 bool = false
                _ = v143
            default:
                if false {
                    v7 = v7
                    fmt.Println(v6, v6)
                    switch "\n\t" {
                    case "a":
                        v144 := GenS16{F0: 2, F1: -103}
                        _ = v144
                    case "\n\t":
                        for fi145 := 0; fi145 < 10; fi145++ {
                            var v146 uint64 = 0
                            for fi147 := 0; fi147 < 16; fi147++ {
                                var v148 uint = (v7 >> (((^(v7))) & 63))
                                _ = fi147
                                _ = v148
                            }
                            _ = fi145
                            _ = v146
                        }
                    default:
                        v149 := GenS17{F0: 32767, F1: -128}
                        _ = v149
                    }
                }
            }
            for fi150 := 0; fi150 < 10; fi150++ {
                switch fi150 {
                case -217:
                    var v151 int16 = -975
                    _ = v151
                case 32767:
                    v152 := []int{}
                    _ = v152
                case 9223372036854775807:
                    v153 := GenS1{F0: v1}
                    _ = v153
                default:
                    fmt.Println(v1, v1)
                }
                fmt.Println(fi150, v4, v7)
                var v154 int64 = (v1 + (v1 / ((706) | 1)))
                _ = fi150
                _ = v154
            }
        } else {
            var v155 uint64 = 1
            v6 = (^(v6))
            var v156 uint32 = 2147483648
            _ = v155
            _ = v156
        }
    default:
        v1 -= (v1 << ((255) & 63))
    }
    var v157 uint8 = 167
    for fi158 := 0; fi158 < 7; fi158++ {
        if false {
            v7 -= v7
            if true {
                var v159 uint16 = 2
                var v160 bool = (!(!false))
                v159 += (v159 & (v159 & v159))
                _ = v159
                _ = v160
            }
        } else {
            v1 += 0
            var v161 int = (fi158 / ((fi158) | 1))
            var v162 uint64 = 2
            _ = v161
            _ = v162
        }
        var v163 int16 = -32768
        v164 := GenS17{F0: (-(v163)), F1: 100}
        _ = fi158
        _ = v163
        _ = v164
    }
    v7 = uint(v157)
    for fi165 := 0; fi165 < 9; fi165++ {
        var v166 float32 = (float32(0.0) / (-2.25 - (v4 + float32(float64((v4 * ((3.14159 - float32(1e10)) + 0.0)))))))
        v166 = float32((0.0 + (1.5 + (((((-2.25 - 0.0) + float64(v166)) - float64((((-2.25 + (v4 - (-2.25 * v166))) + ((v166 + v166) / (0.1 * ((v4 - v4) + (-2.25 * (((v166 - 1.5) + 0.0) * v166)))))) / (-2.25 / ((0.1 / (3.14159 * (0.1 + float32(3.14159)))) * (1.5 - v166)))))) - 3.14159) - float64(float32(float64(v166)))))))
        _ = fi165
        _ = v166
    }
    var v167 float32 = v4
    switch "hello" {
    case "\n\t":
        if false {
            for fi168 := 0; fi168 < 14; fi168++ {
                switch "\n\t" {
                case "hello":
                    var v169 int32 = -1
                    _ = v169
                case "a":
                    switch "hello" {
                    case "":
                        if false {
                            var v170 int16 = 0
                            switch fi168 {
                            case 32768:
                                var v171 float32 = v4
                                _ = v171
                            case 2147483647:
                                v1 -= 390
                            case 0:
                                var v172 uint = (v7 >> ((v7) & 63))
                                _ = v172
                            default:
                                var v173 uint32 = 0
                                _ = v173
                            }
                            _ = v170
                        } else {
                            for fi174 := 0; fi174 < 5; fi174++ {
                                v175 := []int{}
                                _ = fi174
                                _ = v175
                            }
                            v176 := GenS15{F0: 58, F1: uint64(v6)}
                            _ = v176
                        }
                    default:
                        var v177 int16 = 42
                        _ = v177
                    }
                case "foo bar":
                    switch "" {
                    case "foo bar":
                        v6 += 65535
                    case "a":
                        for fi178 := 0; fi178 < 7; fi178++ {
                            switch "foo bar" {
                            case "\n\t":
                                var v179 float64 = -2.25
                                _ = v179
                            case "hello":
                                var v180 uint16 = 1
                                _ = v180
                            case "a":
                                var v181 uint32 = 1
                                _ = v181
                            }
                            v1 += (^(v1))
                            _ = fi178
                        }
                    default:
                        fmt.Println(v7, v1)
                    }
                }
                _ = fi168
            }
        } else {
            if (-54 <= 9223372036854775807) {
                switch -577 {
                case 128:
                    v6 -= (v6 * (v6 << (((^(v7))) & 63)))
                case -2147483648:
                    switch -612 {
                    case 0:
                        var v182 uint8 = (v157 + (v157 - (v157 * v157)))
                        _ = v182
                    }
                case 42:
                    for fi183 := 0; fi183 < 4; fi183++ {
                        if true {
                            v4 -= (v4 - 0.0)
                            if (!((v157 & v157) <= v157)) {
                                v1 = v1
                                fmt.Println(fi183)
                                for fi184 := 0; fi184 < 18; fi184++ {
                                    var v185 int = (fi184 | 1)
                                    _ = fi184
                                    _ = v185
                                }
                            }
                        } else {
                            v186 := GenS18{F0: -815}
                            _ = v186
                        }
                        _ = fi183
                    }
                }
                var v187 uint32 = (v6 >> ((255) & 63))
                _ = v187
            }
            v188 := GenS15{F0: 106, F1: 380}
            _ = v188
        }
    case "foo bar":
        for fi189 := 0; fi189 < 19; fi189++ {
            for fi190 := 0; fi190 < 16; fi190++ {
                v167 = v167
                _ = fi190
            }
            fmt.Println(v167, v1)
            v1 += v1
            _ = fi189
        }
    case "a":
        fmt.Println(v157, v7, v6)
    default:
        if false {
            switch "hello" {
            case "hello":
                var v191 int16 = 42
                _ = v191
            case "\n\t":
                v167 += (1e10 + -2.25)
            default:
                v192 := []int{32768}
                _ = v192
            }
        }
    }
    v7 = (v7 & 1)
    switch (("foo bar" + "foo bar") + ("foo bar" + "\n\t")) {
    case "hello":
        v193 := GenS19{F0: 1}
        _ = v193
    }
    fmt.Println("v1:", v1)
    fmt.Println("v157:", v157)
    fmt.Println("v167:", v167)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v126
    _ = v157
    _ = v167
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
