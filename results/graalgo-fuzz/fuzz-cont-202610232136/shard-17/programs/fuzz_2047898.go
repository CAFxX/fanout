package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int64; F2 int64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 string; F1 int32; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 bool; F1 uint8; F2 int32 }
type GenS11 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int16; F1 uint16; F2 int8 }
type GenS15 struct { F0 int16; F1 bool }
type GenS16 struct { F0 bool; F1 uint }
type GenS17 struct { F0 int8; F1 int8; F2 uint64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 uint32; F1 uint64 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 uint8; F1 uint32 }
type GenS22 struct { F0 int }
type GenS23 struct { F0 uint64; F1 uint; F2 int }
type GenS24 struct { F0 bool; F1 uint32; F2 int64 }
type GenS25 struct { F0 uint; F1 uint }
type GenS26 struct { F0 int32; F1 uint64; F2 uint64 }

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 1.5
    v3 := make(map[int]bool)
    var v4 int64 = 9223372036854775807
    var v5 float32 = v1
    var v6 uint16 = 42
    if (255 < 0) {
        if (((!true) || true) && (false && false)) {
            var v7 int64 = -919
            if (true || true) {
                _ = len(v3)
            } else {
                if (65535 <= 255) {
                    fmt.Println(v5, v1, v1)
                    if (v2 >= v2) {
                        if (1 == -979) {
                            var v8 uint16 = (v6 ^ (v6 << ((255) & 63)))
                            _ = v8
                        }
                        if true {
                            if false {
                                v4 = (v7 - 0)
                            } else {
                                fmt.Println(v4)
                            }
                        }
                    } else {
                        var v9 string = "a"
                        _ = v9
                    }
                } else {
                    delete(v3, 42)
                    if true {
                        var v10 bool = (!false)
                        switch 33 {
                        case 923:
                            if false {
                                v11 := GenS1{F0: 255, F1: true, F2: v10}
                                if v10 {
                                    var v12 uint16 = v6
                                    _ = v12
                                } else {
                                    for fi13 := 0; fi13 < 6; fi13++ {
                                        v14 := GenS1{F0: 128, F1: (v10 || true), F2: v10}
                                        _ = len(v3)
                                        _ = fi13
                                        _ = v14
                                    }
                                }
                                var v15 uint32 = 2
                                _ = v11
                                _ = v15
                            } else {
                                var v16 float64 = float64(3.14159)
                                _ = v16
                            }
                        case 32768:
                            for fi17 := 0; fi17 < 6; fi17++ {
                                v10 = false
                                if false {
                                    var v18 bool = (42 >= 42)
                                    _ = v18
                                }
                                _ = fi17
                            }
                        case 128:
                            var v19 uint8 = 127
                            _ = v19
                        }
                        _ = len(v3)
                        _ = v10
                    }
                }
            }
            var v20 int8 = -1
            _ = v20
            _ = v7
        } else {
            if false {
                fmt.Println(v4, v5)
                if (!true) {
                    if true {
                        var v21 string = ("a" + "hello")
                        v21 = v21
                        _ = v21
                    } else {
                        v5 += (v5 - v1)
                    }
                    var v22 bool = true
                    v23 := GenS2{F0: (-635 < -2147483648), F1: "hello"}
                    _ = v22
                    _ = v23
                }
                v3[0] = (true || false)
            }
            var v24 uint8 = 123
            _ = v24
        }
    } else {
        fmt.Println(v4)
        var v25 int8 = 4
        v26 := GenS3{F0: 4294967295}
        _ = v25
        _ = v26
    }
    switch "hello" {
    case "a":
        v27 := GenS4{F0: -1, F1: v4, F2: v4}
        _ = v27
    case "":
        if (2 <= 1) {
            switch 32768 {
            case -2147483648:
                if ((v4 + v4) != (v4 - v4)) {
                    v28 := []int{1}
                    v1 += v1
                    for fi29 := 0; fi29 < 12; fi29++ {
                        var v30 uint32 = 4294967295
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            case 2147483647:
                v4 = v4
            }
            var v31 int = 719
            _ = v31
        }
    default:
        var v32 int64 = v4
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v6 -= v6
        _ = fi33
    }
    if false {
        v5 = -2.25
        for fi34 := 0; fi34 < 18; fi34++ {
            v35 := GenS5{F0: fi34, F1: (^(fi34))}
            v36 := GenS1{F0: 255, F1: (true && true), F2: false}
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    if ("foo bar" <= "") {
        v4 -= (v4 % ((v4) | 1))
    } else {
        var v37 uint64 = 429
        _ = v37
    }
    fmt.Println(v5, v5, v6)
    v38 := GenS6{F0: "\n\t", F1: 2147483647, F2: 0}
    switch 127 {
    case 1:
        delete(v3, 128)
    }
    v4 = (^(v4))
    v39 := GenS7{F0: int16(v4)}
    if true {
        v4 = (v4 - (v4 >> ((255) & 63)))
        v40 := GenS2{F0: false, F1: "\n\t"}
        _ = v40
    } else {
        v41 := GenS2{F0: (!false), F1: (("foo bar" + "") + "hello")}
        v1 = (1.5 + v5)
        v6 -= v6
        _ = v41
    }
    v42 := GenS3{F0: 18446744073709551615}
    if false {
        v1 += ((1e10 * v1) * v5)
        switch 399 {
        case 344:
            if (843 > 254) {
                if (223 == 42) {
                    var v43 uint8 = 220
                    v44 := GenS1{F0: (v43 & (v43 % (((v43 / ((255) | 1))) | 1))), F1: true, F2: true}
                    _ = v43
                    _ = v44
                }
                v4 = 1
                v45 := []int{-955, 42}
                _ = v45
            }
        case 32768:
            var v46 bool = false
            _ = v46
        }
    } else {
        switch -2147483648 {
        case 32767:
            if true {
                if (68 > 1) {
                    fmt.Println(v4)
                    for fi47 := 0; fi47 < 5; fi47++ {
                        var v48 bool = false
                        var v49 uint32 = 585
                        switch (fi47 >> ((65535) & 63)) {
                        case 4294967295:
                            v3[-289] = v48
                        }
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                    v4 = (v4 | 1)
                }
            }
        case -2147483648:
            if false {
                fmt.Println(v6, v2, v4)
                if false {
                    _ = len(v3)
                    v50 := GenS4{F0: -9223372036854775808, F1: v4, F2: (v4 % (((v4 ^ -9223372036854775808)) | 1))}
                    switch "" {
                    case "":
                        switch 1 {
                        case 32767:
                            for fi51 := 0; fi51 < 20; fi51++ {
                                if false {
                                    for fi52 := 0; fi52 < 4; fi52++ {
                                        v53 := GenS8{F0: 4294967295}
                                        var v54 uint = 18446744073709551615
                                        if true {
                                            v55 := v3[fi52]
                                            v56 := []int{fi51}
                                            switch "hello" {
                                            case "hello":
                                                v57 := GenS7{F0: -1}
                                                _ = v57
                                            case "a":
                                                var v58 int32 = 2147483647
                                                _ = v58
                                            case "\n\t":
                                                fmt.Println(v4, v4)
                                            }
                                            v1 = 3.14159
                                            v5 += 3.14159
                                            _ = v55
                                            _ = v56
                                        }
                                        if true {
                                            var v59 int32 = 42
                                            _ = v59
                                        }
                                        var v60 int = 42
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                        _ = v60
                                    }
                                    switch fi51 {
                                    case 4294967295:
                                        v4 += v4
                                    }
                                    switch fi51 {
                                    case 128:
                                        v61 := GenS9{F0: 1, F1: 127}
                                        _ = v61
                                    case 2147483647:
                                        v5 = v5
                                    case 4294967295:
                                        v62 := GenS3{F0: 634}
                                        _ = v62
                                    default:
                                        if (true || ((fi51 >> ((255) & 63)) >= 4294967295)) {
                                            v1 -= -2.25
                                            switch 42 {
                                            case -840:
                                                if (true && (0 <= 1)) {
                                                    v1 += (v1 + v5)
                                                } else {
                                                    var v63 uint16 = 42
                                                    v1 = float32(v2)
                                                    var v64 uint8 = 128
                                                    _ = v63
                                                    _ = v64
                                                }
                                            case 128:
                                                if (1 < 0) {
                                                    var v65 int64 = (v4 & (v4 & 42))
                                                    _ = v65
                                                } else {
                                                    v2 += (0.0 + v2)
                                                }
                                            case 421:
                                                if (2 == 596) {
                                                    var v66 uint16 = v6
                                                    _ = v66
                                                } else {
                                                    var v67 uint64 = 113
                                                    switch -131 {
                                                    case 1:
                                                        switch (("foo bar" + "foo bar") + "foo bar") {
                                                        case "\n\t":
                                                            v68 := v3[(fi51 >> ((255) & 63))]
                                                            _ = v68
                                                        }
                                                    }
                                                    var v69 int8 = 1
                                                    _ = v67
                                                    _ = v69
                                                }
                                            default:
                                                v70 := GenS10{F0: true, F1: 2, F2: 1}
                                                _ = v70
                                            }
                                            v4 += (v4 % ((v4) | 1))
                                        }
                                    }
                                } else {
                                    if false {
                                        for fi71 := 0; fi71 < 15; fi71++ {
                                            v72 := []int{fi71, fi51, (fi51 & fi51)}
                                            switch "foo bar" {
                                            case "":
                                                switch ((("\n\t" + "foo bar") + "\n\t") + "hello") {
                                                case "a":
                                                    v1 = v5
                                                default:
                                                    if ("foo bar" > "foo bar") {
                                                        fmt.Println(v6)
                                                        if len(v72) > 0 {
                                                            v72[0] = (-(fi51))
                                                        }
                                                        var v73 float64 = v2
                                                        v1 -= 0.0
                                                        _ = v73
                                                    } else {
                                                        var v74 float32 = v5
                                                        fmt.Println(fi51, v74, v2)
                                                        fmt.Println(fi51, fi51, v4)
                                                        var v75 uint64 = 1
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                }
                                            case "\n\t":
                                                var v76 float32 = (1.5 * v1)
                                                _ = v76
                                            case "foo bar":
                                                v72 = append(v72, (^(fi51)))
                                            }
                                            if (v5 == float32(v2)) {
                                                v1 -= (v5 + v1)
                                                v77 := GenS7{F0: -32768}
                                                _ = len(v3)
                                                _ = v77
                                            }
                                            if false {
                                                v1 = v1
                                                v5 += float32(v2)
                                                if false {
                                                    v1 = (v5 - ((0.1 * v5) + 0.1))
                                                }
                                            } else {
                                                if false {
                                                    switch int(v6) {
                                                    case -9223372036854775808:
                                                        v5 += 3.14159
                                                    case 9223372036854775807:
                                                        v4 = v4
                                                    case -2147483648:
                                                        v5 -= v1
                                                    }
                                                    var v78 float32 = float32((v2 * v2))
                                                    v79 := GenS11{F0: -2147483648, F1: 0, F2: 128}
                                                    _ = v78
                                                    _ = v79
                                                }
                                                var v80 int16 = -859
                                                var v81 int8 = -107
                                                var v82 int32 = int32(fi51)
                                                switch ("\n\t" + ("a" + "\n\t")) {
                                                case "hello":
                                                    v82 = -712
                                                case "a":
                                                    v72 = append(v72, fi71)
                                                }
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            }
                                            _ = fi71
                                            _ = v72
                                        }
                                    }
                                }
                                fmt.Println(v4)
                                var v83 int = (fi51 >> ((18446744073709551615) & 63))
                                _ = fi51
                                _ = v83
                            }
                        }
                    case "a":
                        var v84 float32 = v5
                        _ = v84
                    case "\n\t":
                        if false {
                            for fi85 := 0; fi85 < 12; fi85++ {
                                for fi86 := 0; fi86 < 3; fi86++ {
                                    switch 0 {
                                    case 0:
                                        var v87 int64 = (v4 + 0)
                                        _ = v87
                                    case -832:
                                        if false {
                                            var v88 int16 = 42
                                            v89 := GenS7{F0: 1000}
                                            _ = v88
                                            _ = v89
                                        } else {
                                            switch "a" {
                                            case "hello":
                                                if true {
                                                    var v90 float32 = 0.1
                                                    _ = v90
                                                }
                                            case "":
                                                if true {
                                                    if ((false || (!(true && false))) && true) {
                                                        switch "" {
                                                        case "foo bar":
                                                            if ((42 == -93) && ((0 == 42) || false)) {
                                                                fmt.Println(v6)
                                                            } else {
                                                                v4 = (v4 << ((1) & 63))
                                                            }
                                                        default:
                                                            var v91 string = "foo bar"
                                                            _ = v91
                                                        }
                                                        switch fi85 {
                                                        case -2147483648:
                                                            switch (fi86 - (fi85 << ((18446744073709551615) & 63))) {
                                                            case 932:
                                                                v2 -= -2.25
                                                            case 816:
                                                                switch 4294967295 {
                                                                case 227:
                                                                    if (!(250 != 2)) {
                                                                        v92 := GenS8{F0: 65535}
                                                                        switch -1 {
                                                                        case 42:
                                                                            v93 := GenS5{F0: (fi85 - fi85), F1: (fi85 ^ (fi85 / ((fi85) | 1)))}
                                                                            _ = v93
                                                                        case 768:
                                                                            fmt.Println(fi86, v5)
                                                                        case -277:
                                                                            if false {
                                                                                fmt.Println(v4, fi86)
                                                                            }
                                                                        }
                                                                        _ = v92
                                                                    }
                                                                default:
                                                                    v94 := []int{(fi86 - (fi85 >> ((0) & 63))), fi86}
                                                                    var v95 bool = false
                                                                    switch fi85 {
                                                                    case -970:
                                                                        if true {
                                                                            v96 := GenS12{F0: (v4 % ((0) | 1))}
                                                                            _ = v96
                                                                        } else {
                                                                            v97 := v3[42]
                                                                            fmt.Println(v5, v4)
                                                                            _ = v97
                                                                        }
                                                                        v98 := GenS13{F0: fi85}
                                                                        _ = v98
                                                                    case -1:
                                                                        v99 := GenS9{F0: -128, F1: 199}
                                                                        _ = v99
                                                                    }
                                                                    fmt.Println(v5)
                                                                    _ = v94
                                                                    _ = v95
                                                                }
                                                            case 9223372036854775807:
                                                                var v100 float32 = (((1.5 - (v5 - 1.5)) - v5) + (0.0 - v1))
                                                                _ = v100
                                                            }
                                                        case 32768:
                                                            v4 = (-(v4))
                                                        }
                                                        if false {
                                                            switch 127 {
                                                            case -9223372036854775808:
                                                                var v101 int = (fi86 - (fi86 << ((255) & 63)))
                                                                _ = v101
                                                            case 921:
                                                                var v102 int64 = (v4 | v4)
                                                                _ = v102
                                                            case 766:
                                                                var v103 int8 = -1
                                                                _ = v103
                                                            default:
                                                                v5 += v1
                                                            }
                                                            var v104 string = ""
                                                            _ = v104
                                                        } else {
                                                            switch fi86 {
                                                            case 128:
                                                                var v105 int8 = -75
                                                                _ = v105
                                                            }
                                                        }
                                                        v2 = 3.14159
                                                        fmt.Println(v1, v4, v4)
                                                        v3[fi85] = true
                                                        v1 = v5
                                                    }
                                                    v106 := GenS6{F0: (("foo bar" + "a") + "\n\t"), F1: 0, F2: 255}
                                                    v107 := GenS8{F0: 2147483647}
                                                    _ = v106
                                                    _ = v107
                                                } else {
                                                    var v108 uint64 = 18446744073709551615
                                                    _ = v108
                                                }
                                            case "a":
                                                var v109 string = "\n\t"
                                                _ = v109
                                            }
                                        }
                                    case 734:
                                        switch ("\n\t" + ("a" + ("\n\t" + "foo bar"))) {
                                        case "foo bar":
                                            v110 := GenS13{F0: -2147483648}
                                            _ = v110
                                        default:
                                            var v111 uint8 = 127
                                            _ = v111
                                        }
                                    default:
                                        var v112 uint8 = 255
                                        _ = v112
                                    }
                                    v113 := []int{}
                                    fmt.Println(v1, v1, fi85)
                                    _ = len(v113)
                                    v114 := GenS11{F0: -2147483648, F1: 128, F2: 6}
                                    v1 = (v5 * v1)
                                    _ = fi86
                                    _ = v113
                                    _ = v114
                                }
                                var v115 int64 = v4
                                _ = fi85
                                _ = v115
                            }
                            v116 := GenS14{F0: 761, F1: (v6 & v6), F2: 60}
                            v2 -= (3.14159 - (1.5 + (v2 - -2.25)))
                            _ = v116
                        }
                    default:
                        fmt.Println(v5, v2)
                    }
                    _ = v50
                } else {
                    var v117 int64 = (v4 - -951)
                    _ = v117
                }
                if (((2147483647 != -2147483648) && (65535 <= 0)) && true) {
                    v5 = (v1 - (v5 + (-2.25 / (v1 * v5))))
                    v6 -= (v6 ^ v6)
                    v118 := GenS15{F0: 0, F1: true}
                    _ = v118
                }
            } else {
                var v119 float32 = 1.5
                v3[735] = true
                _ = v119
            }
        case -1:
            v120 := GenS11{F0: 899, F1: 0, F2: 128}
            _ = v120
        }
        v121 := GenS16{F0: false, F1: 4294967295}
        _ = v121
    }
    v122 := []int{4294967295, 127}
    for fi123 := 0; fi123 < 20; fi123++ {
        var v124 int16 = 1000
        if true {
            switch "foo bar" {
            case "\n\t":
                switch "a" {
                case "":
                    fmt.Println(fi123, v6)
                case "\n\t":
                    var v125 uint = 2
                    _ = v125
                }
            case "hello":
                delete(v3, fi123)
            case "":
                fmt.Println(v124, v6, fi123)
            default:
                for fi126 := 0; fi126 < 9; fi126++ {
                    if false {
                        _ = len(v122)
                        var v127 int8 = -57
                        var v128 int8 = -56
                        _ = v127
                        _ = v128
                    }
                    if false {
                        if false {
                            v129 := GenS17{F0: 0, F1: 1, F2: 715}
                            v2 -= v2
                            switch ((("" + "") + ("a" + "foo bar")) + "foo bar") {
                            case "hello":
                                v130 := GenS14{F0: (v124 & (v124 % (((-(v124))) | 1))), F1: 722, F2: -1}
                                _ = v130
                            case "a":
                                fmt.Println(v1)
                            case "\n\t":
                                switch "foo bar" {
                                case "\n\t":
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v1)
                                    case "":
                                        v131 := GenS1{F0: 2, F1: true, F2: true}
                                        _ = v131
                                    default:
                                        var v132 uint16 = 704
                                        var v133 int64 = (v4 << ((42) & 63))
                                        v134 := GenS5{F0: (fi126 % (((^(fi123))) | 1)), F1: -1}
                                        _ = v132
                                        _ = v133
                                        _ = v134
                                    }
                                case "":
                                    if len(v122) > 0 {
                                        v122[0] = (^(fi123))
                                    }
                                case "a":
                                    var v135 string = "a"
                                    _ = v135
                                }
                            default:
                                switch fi126 {
                                case 32768:
                                    if false {
                                        var v136 float32 = ((v5 / v5) + (float32(1.5) * v1))
                                        v6 = v6
                                        var v137 int64 = 9223372036854775807
                                        _ = v136
                                        _ = v137
                                    } else {
                                        switch "a" {
                                        case "a":
                                            v2 = v2
                                        default:
                                            v2 = (((1.5 / v2) + 0.0) - (((((((1e10 + v2) / ((v2 * (v2 - (-2.25 - v2))) - 1.5)) - float64((v1 / float32(1e10)))) / (v2 - 1e10)) - v2) + v2) + -2.25))
                                            if (!(v124 == 32767)) {
                                                if true {
                                                    _ = len(v122)
                                                } else {
                                                    fmt.Println(v1)
                                                    var v138 uint32 = 992
                                                    v124 -= -32768
                                                    _ = v138
                                                }
                                            }
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                var v139 uint = 0
                                                _ = v139
                                            case "":
                                                v1 = v5
                                            }
                                            switch (fi123 % ((939) | 1)) {
                                            case 4294967295:
                                                v140 := GenS2{F0: true, F1: (("hello" + "a") + (("" + (("foo bar" + "a") + "\n\t")) + "hello"))}
                                                _ = v140
                                            case 2147483647:
                                                var v141 int16 = 158
                                                _ = v141
                                            case 9223372036854775807:
                                                if false {
                                                    switch "hello" {
                                                    case "hello":
                                                        if (127 == uint8(v6)) {
                                                            var v142 int16 = v124
                                                            _ = v142
                                                        } else {
                                                            v143 := GenS9{F0: -128, F1: 128}
                                                            if false {
                                                                var v144 int64 = v4
                                                                _ = v144
                                                            } else {
                                                                var v145 int8 = -128
                                                                v4 = 744
                                                                _ = v145
                                                            }
                                                            if true {
                                                                v6 -= 0
                                                                v146 := GenS18{F0: v6}
                                                                _ = v146
                                                            } else {
                                                                v4 += -9223372036854775808
                                                                v147 := v3[fi126]
                                                                v148 := GenS19{F0: 2147483648, F1: 42}
                                                                if len(v122) >= 2 {
                                                                    v122 = v122[1:len(v122)]
                                                                }
                                                                _ = v147
                                                                _ = v148
                                                            }
                                                            _ = v143
                                                        }
                                                    case "\n\t":
                                                        _ = len(v3)
                                                    }
                                                } else {
                                                    if true {
                                                        if false {
                                                            switch fi123 {
                                                            case 0:
                                                                v6 -= v6
                                                            case -1:
                                                                v5 = float32((0.1 + ((float64(v1) * (0.0 + v2)) / 3.14159)))
                                                            case 128:
                                                                v6 = 0
                                                            default:
                                                                v124 = (v124 - v124)
                                                            }
                                                            delete(v3, fi123)
                                                            var v149 int16 = 32767
                                                            _ = v149
                                                        }
                                                        var v150 int8 = 0
                                                        v4 = -1
                                                        _ = v150
                                                    } else {
                                                        if false {
                                                            var v151 int32 = 0
                                                            _ = v151
                                                        } else {
                                                            switch (fi126 + fi123) {
                                                            case -2147483648:
                                                                switch (fi123 & (fi123 * -370)) {
                                                                case -704:
                                                                    var v152 int16 = (v124 | (-(v124)))
                                                                    _ = v152
                                                                default:
                                                                    switch (("foo bar" + (("a" + ("hello" + "hello")) + "hello")) + ("foo bar" + "a")) {
                                                                    case "\n\t":
                                                                        var v153 int16 = v124
                                                                        _ = v153
                                                                    case "hello":
                                                                        v124 = (v124 & v124)
                                                                    }
                                                                }
                                                                v154 := v3[fi123]
                                                                var v155 float64 = v2
                                                                switch 42 {
                                                                case -285:
                                                                    fmt.Println(fi123, fi123, v4)
                                                                case -2147483648:
                                                                    var v156 uint32 = 677
                                                                    _ = v156
                                                                case -1:
                                                                    var v157 uint = 2
                                                                    _ = v157
                                                                }
                                                                v122 = append(v122, fi126)
                                                                _ = v154
                                                                _ = v155
                                                            default:
                                                                if (fi123 != fi126) {
                                                                    if (v6 != (^(v6))) {
                                                                        var v158 uint32 = 888
                                                                        if ((v124 - (v124 * 1000)) != 0) {
                                                                            var v159 float32 = ((float32(v2) / v5) / 1.0)
                                                                            _ = len(v122)
                                                                            _ = v159
                                                                        }
                                                                        _ = v158
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch (fi123 << ((384) & 63)) {
                                                        case -9223372036854775808:
                                                            var v160 uint8 = 255
                                                            _ = v160
                                                        case 32767:
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                v161 := GenS6{F0: "hello", F1: -806, F2: 2}
                                                                _ = v161
                                                            case "a":
                                                                v1 -= v5
                                                            case "foo bar":
                                                                fmt.Println(v4, fi126)
                                                            }
                                                            v3[0] = false
                                                            var v162 uint16 = (v6 + (v6 ^ (v6 - v6)))
                                                            if true {
                                                                if (!true) {
                                                                    v6 -= 2
                                                                    v1 -= v1
                                                                    var v163 uint = 218
                                                                    if len(v122) >= 2 {
                                                                        v122 = v122[1:len(v122)]
                                                                    }
                                                                    _ = v163
                                                                }
                                                                var v164 int8 = 1
                                                                var v165 uint = 2
                                                                var v166 string = "hello"
                                                                _ = v164
                                                                _ = v165
                                                                _ = v166
                                                            }
                                                            delete(v3, 0)
                                                            _ = v162
                                                        case -1:
                                                            if (!((-(v4)) >= -1)) {
                                                                v2 -= (1.5 / v2)
                                                            }
                                                            fmt.Println(v2, fi123)
                                                            v6 = v6
                                                        }
                                                        switch (fi126 + (fi123 * int(v6))) {
                                                        case -705:
                                                            v167 := GenS14{F0: v124, F1: 618, F2: -128}
                                                            _ = v167
                                                        case 127:
                                                            v168 := GenS12{F0: (v4 << ((uint(v6)) & 63))}
                                                            _ = v168
                                                        case 2147483647:
                                                            v169 := GenS19{F0: 988, F1: 18446744073709551615}
                                                            if true {
                                                                var v170 int64 = v4
                                                                switch "" {
                                                                case "hello":
                                                                    _ = len(v122)
                                                                case "a":
                                                                    switch (("" + "") + ("hello" + "hello")) {
                                                                    case "foo bar":
                                                                        switch fi126 {
                                                                        case -279:
                                                                            v170 = v4
                                                                        case 890:
                                                                            if (2 == 451) {
                                                                                v5 -= ((-2.25 * ((v5 + 0.1) - v1)) * v5)
                                                                                if true {
                                                                                    v171 := GenS12{F0: v170}
                                                                                    switch fi123 {
                                                                                    case 32768:
                                                                                        fmt.Println(v1)
                                                                                        v6 = 255
                                                                                    case 128:
                                                                                        if false {
                                                                                            var v172 int64 = (^(v170))
                                                                                            var v173 uint16 = (v6 / (((v6 % ((v6) | 1))) | 1))
                                                                                            _ = v172
                                                                                            _ = v173
                                                                                        } else {
                                                                                            v170 += (v170 % ((v4) | 1))
                                                                                            _ = len(v3)
                                                                                            var v174 int32 = 42
                                                                                            _ = v174
                                                                                        }
                                                                                    case -62:
                                                                                        var v175 int64 = 9223372036854775807
                                                                                        _ = v175
                                                                                    }
                                                                                    fmt.Println(v2, v5, v2)
                                                                                    _ = v171
                                                                                } else {
                                                                                    var v176 int8 = 32
                                                                                    var v177 uint64 = 4294967295
                                                                                    switch ("foo bar" + "foo bar") {
                                                                                    case "":
                                                                                        if len(v122) >= 2 {
                                                                                            v122 = v122[1:len(v122)]
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v1 += v5
                                                                                    case "hello":
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v178 uint64 = 0
                                                                                            v6 = (v6 % (((v6 >> ((uint(v178)) & 63))) | 1))
                                                                                            if len(v122) >= 2 {
                                                                                                v122 = v122[1:len(v122)]
                                                                                            }
                                                                                            _ = v178
                                                                                        }
                                                                                    default:
                                                                                        var v179 int16 = v124
                                                                                        _ = v179
                                                                                    }
                                                                                    _ = v176
                                                                                    _ = v177
                                                                                }
                                                                                v2 += v2
                                                                            }
                                                                        }
                                                                    case "":
                                                                        v180 := GenS20{F0: 2147483647}
                                                                        _ = v180
                                                                    case "hello":
                                                                        var v181 int = 32768
                                                                        var v182 string = ""
                                                                        v2 = (1.5 / 3.14159)
                                                                        _ = v181
                                                                        _ = v182
                                                                    default:
                                                                        var v183 uint16 = 32767
                                                                        _ = v183
                                                                    }
                                                                }
                                                                _ = v170
                                                            }
                                                            _ = v169
                                                        }
                                                    }
                                                }
                                            }
                                            var v184 uint64 = 498
                                            v185 := GenS21{F0: 128, F1: 2147483647}
                                            if false {
                                                if ((false || (v184 <= 946)) && false) {
                                                    v5 = ((v1 * v1) + (3.14159 / 1.0))
                                                    v186 := GenS14{F0: v124, F1: v6, F2: 1}
                                                    var v187 uint32 = 365
                                                    _ = v186
                                                    _ = v187
                                                }
                                                v5 -= ((v1 + v5) / (v5 - v1))
                                            } else {
                                                if (963 > -458) {
                                                    var v188 bool = ((^(fi123)) <= 32767)
                                                    v4 -= (v4 << ((65535) & 63))
                                                    v184 = (v184 & (v184 | v184))
                                                    _ = v188
                                                } else {
                                                    switch (fi123 >> ((1) & 63)) {
                                                    case 1:
                                                        if (false && (!true)) {
                                                            v122 = append(v122, -9223372036854775808)
                                                            if (1 != 848) {
                                                                if false {
                                                                    fmt.Println(v5)
                                                                    var v189 uint8 = 127
                                                                    var v190 uint32 = 66
                                                                    v2 = float64(0.1)
                                                                    v191 := GenS4{F0: (v4 % (((v4 | v4)) | 1)), F1: (v4 - (^(v4))), F2: 9223372036854775807}
                                                                    v192 := GenS22{F0: (fi126 | (fi126 & fi123))}
                                                                    if len(v122) > 0 {
                                                                        v122[0] = fi126
                                                                    }
                                                                    v4 = (v4 ^ v4)
                                                                    switch "hello" {
                                                                    case "":
                                                                        var v193 uint64 = (v184 >> ((747) & 63))
                                                                        _ = v193
                                                                    case "\n\t":
                                                                        _ = len(v122)
                                                                        v184 += 1
                                                                        v5 = (v5 - 0.0)
                                                                        fmt.Println(v189, fi126, v5)
                                                                        v194 := GenS17{F0: 0, F1: 100, F2: v184}
                                                                        fmt.Println(v5)
                                                                        var v195 float32 = 1.5
                                                                        _ = v194
                                                                        _ = v195
                                                                    }
                                                                    _ = v189
                                                                    _ = v190
                                                                    _ = v191
                                                                    _ = v192
                                                                }
                                                                var v196 bool = (v2 == (0.0 / -2.25))
                                                                var v197 float32 = (v1 * 3.14159)
                                                                v122 = append(v122, fi126)
                                                                _ = v196
                                                                _ = v197
                                                            } else {
                                                                fmt.Println(v1)
                                                                var v198 uint8 = 10
                                                                v199 := GenS6{F0: "a", F1: 42, F2: (^(v198))}
                                                                _ = v198
                                                                _ = v199
                                                            }
                                                            var v200 uint32 = 0
                                                            _ = v200
                                                        }
                                                    case -34:
                                                        fmt.Println(fi126, v1)
                                                    case -9223372036854775808:
                                                        v201 := GenS12{F0: (^(v4))}
                                                        _ = v201
                                                    default:
                                                        v202 := GenS20{F0: -2147483648}
                                                        _ = v202
                                                    }
                                                    v203 := GenS23{F0: v184, F1: uint(v184), F2: (fi123 % ((int(v124)) | 1))}
                                                    _ = v203
                                                }
                                            }
                                            _ = v184
                                            _ = v185
                                        }
                                        v204 := GenS19{F0: 2147483647, F1: 1}
                                        v124 -= v124
                                        v205 := GenS24{F0: (!(uint32(v6) == 377)), F1: 2147483648, F2: v4}
                                        _ = v204
                                        _ = v205
                                    }
                                case 4294967295:
                                    v124 += 169
                                case 1:
                                    switch ("a" + "\n\t") {
                                    case "foo bar":
                                        if false {
                                            delete(v3, fi126)
                                            v206 := GenS8{F0: 2147483647}
                                            v207 := GenS25{F0: 564, F1: 65535}
                                            _ = v206
                                            _ = v207
                                        } else {
                                            v208 := GenS25{F0: 65535, F1: 600}
                                            _ = v208
                                        }
                                    case "\n\t":
                                        fmt.Println(v2, v1, v1)
                                        switch "foo bar" {
                                        case "hello":
                                            v124 = v124
                                        case "a":
                                            _ = len(v122)
                                        case "\n\t":
                                            var v209 bool = true
                                            _ = v209
                                        default:
                                            var v210 uint = 255
                                            _ = v210
                                        }
                                        v211 := GenS4{F0: (v4 >> ((1) & 63)), F1: v4, F2: 553}
                                        _ = v211
                                    default:
                                        if (985 >= 536) {
                                            fmt.Println(v124)
                                        } else {
                                            if len(v122) >= 2 {
                                                v122 = v122[1:len(v122)]
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v129
                        } else {
                            fmt.Println(v2)
                            var v212 float64 = (v2 - (v2 - 0.1))
                            _ = v212
                        }
                    }
                    var v213 string = ("a" + "foo bar")
                    _ = fi126
                    _ = v213
                }
            }
            var v214 uint8 = 2
            if (false || (!false)) {
                fmt.Println(v2, v214)
                switch 127 {
                case 742:
                    v2 = v2
                default:
                    if len(v122) > 0 {
                        v122[0] = fi123
                    }
                }
            }
            _ = v214
        }
        var v215 uint16 = 32767
        _ = fi123
        _ = v124
        _ = v215
    }
    if (0.1 == (v2 + (1e10 - 3.14159))) {
        switch "" {
        case "foo bar":
            v1 = (v5 * v1)
        case "\n\t":
            v216 := v3[-711]
            _ = v216
        case "":
            var v217 uint16 = v6
            _ = v217
        default:
            for fi218 := 0; fi218 < 17; fi218++ {
                var v219 uint32 = 4294967295
                v1 -= v5
                v1 -= 1.5
                _ = fi218
                _ = v219
            }
        }
        v220 := GenS26{F0: -2147483648, F1: 747, F2: 18446744073709551615}
        _ = v220
    } else {
        for fi221 := 0; fi221 < 16; fi221++ {
            v3[fi221] = true
            v6 = (^(v6))
            _ = fi221
        }
        var v222 int32 = 888
        for fi223 := 0; fi223 < 7; fi223++ {
            v224 := GenS25{F0: 133, F1: 4294967295}
            var v225 uint64 = 733
            var v226 bool = true
            _ = fi223
            _ = v224
            _ = v225
            _ = v226
        }
        _ = v222
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v122
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
