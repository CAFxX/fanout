package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 uint32; F2 int8 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int16; F1 int; F2 uint32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 bool; F1 int8 }
type GenS15 struct { F0 uint16; F1 uint8 }
type GenS16 struct { F0 int32; F1 uint64 }
type GenS17 struct { F0 int16; F1 uint64 }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 int32; F1 string; F2 int32 }

func main() {
    var v1 uint = 129
    var v2 float32 = (1.5 + 0.1)
    var v3 int64 = 42
    v4 := make(map[int]string)
    var v5 uint32 = 0
    var v6 int8 = -72
    if false {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                fmt.Println(v5, v6, v1)
                v9 := GenS1{F0: 163}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 19; fi10++ {
                switch (fi7 | 1) {
                case -25:
                    v1 -= 255
                    var v11 uint32 = 2147483648
                    _ = v11
                }
                var v12 int8 = (v6 - (v6 & (v6 % ((1) | 1))))
                _ = fi10
                _ = v12
            }
            _ = fi7
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "" {
            case "foo bar":
                fmt.Println(v5)
            case "\n\t":
                v14 := []int{(fi13 | 4294967295), (-(fi13)), (fi13 / (((fi13 | (^(fi13)))) | 1)), fi13}
                _ = v14
            case "a":
                switch "a" {
                case "foo bar":
                    var v15 float64 = (1.5 / 1.5)
                    _ = v15
                }
            }
            v16 := GenS2{F0: 1, F1: "hello", F2: v1}
            switch (("" + "") + "") {
            case "":
                fmt.Println(v2, v5, fi13)
            default:
                fmt.Println(v5)
            }
            _ = fi13
            _ = v16
        }
        v17 := GenS3{F0: (true && true)}
        _ = v17
    } else {
        v6 -= v6
        v18 := []int{715}
        _ = v18
    }
    v19 := GenS2{F0: -32768, F1: "\n\t", F2: 65535}
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS1{F0: -1}
        _ = fi20
        _ = v21
    }
    if ((v3 < -163) && (69 >= int(v3))) {
        switch "" {
        case "\n\t":
            var v22 uint32 = (v5 | 170)
            _ = v22
        case "hello":
            if true {
                v1 = 42
                fmt.Println(v6, v3, v1)
            }
        case "a":
            switch 2147483647 {
            case 0:
                if false {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        switch fi23 {
                        case 9223372036854775807:
                            v1 -= (v1 ^ 410)
                        case -337:
                            var v24 int = (fi23 >> ((682) & 63))
                            _ = v24
                        case 1:
                            v25 := []int{fi23, -781}
                            _ = v25
                        }
                        fmt.Println(fi23, v1, v2)
                        _ = fi23
                    }
                    v26 := GenS4{F0: int32(v6)}
                    _ = len(v4)
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v28 := GenS1{F0: -533}
                        _ = fi27
                        _ = v28
                    }
                }
            case 1:
                for fi29 := 0; fi29 < 11; fi29++ {
                    v2 = (v2 * float32(float64(v2)))
                    var v30 string = "hello"
                    v31 := []int{fi29}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case 32767:
                v6 = (^(v6))
            }
        }
        var v32 string = ((((((("hello" + ("hello" + ("\n\t" + ((("a" + "") + "") + "a")))) + "hello") + "") + ("\n\t" + (("\n\t" + "\n\t") + "a"))) + "foo bar") + ("hello" + "foo bar")) + "a")
        var v33 string = "foo bar"
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 20; fi34++ {
            v1 = (v1 * uint(v5))
            _ = fi34
        }
        var v35 bool = (v5 != v5)
        var v36 uint16 = 0
        _ = v35
        _ = v36
    }
    v3 = int64(v6)
    var v37 uint32 = (v5 + 4294967295)
    for fi38 := 0; fi38 < 14; fi38++ {
        if false {
            v39 := GenS5{F0: 1000, F1: v37, F2: (v6 / ((v6) | 1))}
            _ = v39
        }
        _ = fi38
    }
    v40 := []int{128, 0, 2147483647}
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    if true {
        switch "hello" {
        case "\n\t":
            for fi41 := 0; fi41 < 6; fi41++ {
                fmt.Println(v2, v5)
                var v42 int8 = v6
                var v43 uint8 = 106
                _ = fi41
                _ = v42
                _ = v43
            }
        case "hello":
            v1 -= (v1 * v1)
        case "foo bar":
            delete(v4, 871)
        }
        for fi44 := 0; fi44 < 4; fi44++ {
            switch "hello" {
            case "a":
                v37 -= 2147483648
            default:
                if true {
                    for fi45 := 0; fi45 < 20; fi45++ {
                        switch ("" + "") {
                        case "a":
                            if false {
                                if len(v40) > 0 {
                                    v40[0] = (fi44 & (-(fi45)))
                                }
                                var v46 uint8 = 128
                                _ = v46
                            } else {
                                v4[(fi44 | (fi44 >> ((v1) & 63)))] = "a"
                            }
                        case "foo bar":
                            switch ("foo bar" + "foo bar") {
                            case "a":
                                v47 := v4[(fi45 - fi45)]
                                _ = v47
                            case "":
                                v1 = v1
                            case "foo bar":
                                v4[(fi44 * (fi44 + 9223372036854775807))] = (("foo bar" + "a") + (("a" + (("a" + "a") + "foo bar")) + ((("a" + "\n\t") + (("a" + "\n\t") + ("a" + ""))) + "hello")))
                            default:
                                if (!(4294967295 <= 183)) {
                                    var v48 int64 = -1
                                    _ = v48
                                } else {
                                    var v49 uint64 = 2
                                    _ = len(v40)
                                    if true {
                                        v3 = -9223372036854775808
                                        v3 -= v3
                                    }
                                    if (!(9223372036854775807 <= (fi44 ^ int(v49)))) {
                                        v4[42] = "\n\t"
                                        if false {
                                            switch int(v1) {
                                            case 42:
                                                var v50 uint64 = (v49 << (((v1 << ((v1) & 63))) & 63))
                                                _ = v50
                                            case -376:
                                                if ((0 > 293) && false) {
                                                    var v51 float64 = (0.0 + 0.0)
                                                    v49 = (v49 + (v49 ^ (v49 / (((v49 | uint64(v1))) | 1))))
                                                    _ = v51
                                                }
                                            case 117:
                                                var v52 uint = (v1 / (((^(v1))) | 1))
                                                _ = v52
                                            default:
                                                var v53 int16 = -866
                                                _ = v53
                                            }
                                        }
                                        v49 -= (v49 << ((v1) & 63))
                                        if (-630 > 1) {
                                            if ((v1 - v1) > (v1 / ((v1) | 1))) {
                                                _ = len(v4)
                                                var v54 int32 = int32(v37)
                                                _ = v54
                                            } else {
                                                fmt.Println(fi45)
                                            }
                                            v3 -= -1
                                            var v55 uint8 = 128
                                            v5 += (v5 ^ (v5 - (v5 >> ((uint(v55)) & 63))))
                                            var v56 bool = (true || true)
                                            if v56 {
                                                if false {
                                                    v37 += uint32(v49)
                                                    if v56 {
                                                        if (!true) {
                                                            var v57 int32 = 0
                                                            _ = v57
                                                        } else {
                                                            v5 -= 832
                                                        }
                                                        var v58 uint16 = 65535
                                                        v56 = false
                                                        _ = v58
                                                    }
                                                    if v56 {
                                                        v56 = (42 == -2147483648)
                                                        var v59 int32 = -2147483648
                                                        switch ("a" + "hello") {
                                                        case "\n\t":
                                                            var v60 uint = v1
                                                            _ = v60
                                                        case "":
                                                            var v61 uint32 = v37
                                                            _ = v61
                                                        case "foo bar":
                                                            fmt.Println(v55, v49)
                                                        default:
                                                            switch (fi44 % (((^(fi45))) | 1)) {
                                                            case -701:
                                                                var v62 int64 = v3
                                                                _ = v62
                                                            case -600:
                                                                v63 := v4[(fi44 + int(v59))]
                                                                _ = v63
                                                            case -525:
                                                                var v64 uint16 = 32767
                                                                _ = v64
                                                            }
                                                            var v65 uint32 = v5
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v66 float64 = 1.5
                                                                _ = v66
                                                            case "a":
                                                                v67 := GenS2{F0: -768, F1: ("\n\t" + ("hello" + "")), F2: (v1 << ((v1) & 63))}
                                                                _ = v67
                                                            case "hello":
                                                                switch fi44 {
                                                                case 1:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        v56 = (!v56)
                                                                    case "a":
                                                                        var v68 uint64 = v49
                                                                        _ = v68
                                                                    default:
                                                                        switch ("\n\t" + ("hello" + "a")) {
                                                                        case "":
                                                                            v69 := GenS5{F0: 845, F1: 2, F2: 127}
                                                                            _ = v69
                                                                        case "foo bar":
                                                                            switch ("hello" + ("\n\t" + "")) {
                                                                            case "":
                                                                                v49 = uint64(v1)
                                                                            case "hello":
                                                                                var v70 int32 = int32(v1)
                                                                                _ = v70
                                                                            }
                                                                        default:
                                                                            v3 += v3
                                                                        }
                                                                    }
                                                                case 128:
                                                                    var v71 int16 = 1000
                                                                    _ = v71
                                                                case 678:
                                                                    _ = len(v4)
                                                                default:
                                                                    switch ("a" + "\n\t") {
                                                                    case "foo bar":
                                                                        v56 = v56
                                                                        v65 = (v37 * v37)
                                                                        var v72 int64 = v3
                                                                        _ = v72
                                                                    }
                                                                }
                                                            }
                                                            _ = v65
                                                        }
                                                        if (710 > -422) {
                                                            if v56 {
                                                                v73 := GenS5{F0: 0, F1: (^(v37)), F2: 127}
                                                                _ = v73
                                                            } else {
                                                                fmt.Println(fi44)
                                                                fmt.Println(v3)
                                                            }
                                                        } else {
                                                            var v74 uint64 = v49
                                                            if v56 {
                                                                v75 := GenS6{F0: "", F1: false}
                                                                v37 -= 102
                                                                if (32767 < 42) {
                                                                    if v56 {
                                                                        if v56 {
                                                                            v49 -= (v49 >> ((v1) & 63))
                                                                            var v76 float64 = -2.25
                                                                            v77 := v4[fi45]
                                                                            fmt.Println(fi44, v6, fi45)
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            if true {
                                                                                if false {
                                                                                    v1 -= (^(v1))
                                                                                    var v78 bool = true
                                                                                    var v79 int8 = (-(v6))
                                                                                    _ = v78
                                                                                    _ = v79
                                                                                }
                                                                                v74 -= v74
                                                                                if (v6 != (v6 / ((100) | 1))) {
                                                                                    if (!(int32(v3) != -2147483648)) {
                                                                                        var v80 int32 = v59
                                                                                        v74 += (v74 - 42)
                                                                                        _ = v80
                                                                                    } else {
                                                                                        v81 := GenS3{F0: v56}
                                                                                        v6 = v6
                                                                                        v59 -= (v59 | (v59 - 976))
                                                                                        if (v5 == 57) {
                                                                                            switch (fi45 + (fi44 - (fi45 % ((fi44) | 1)))) {
                                                                                            case -419:
                                                                                                var v82 uint8 = 127
                                                                                                _ = v82
                                                                                            default:
                                                                                                delete(v4, fi45)
                                                                                            }
                                                                                            v83 := GenS2{F0: 32767, F1: "a", F2: (v1 * (v1 / (((v1 << ((v1) & 63))) | 1)))}
                                                                                            var v84 bool = (fi45 >= fi45)
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                        }
                                                                                        _ = v81
                                                                                    }
                                                                                    _ = len(v40)
                                                                                    v85 := v4[fi45]
                                                                                    _ = len(v40)
                                                                                    _ = v85
                                                                                }
                                                                            } else {
                                                                                v37 = (v37 / (((^(v37))) | 1))
                                                                            }
                                                                            var v86 uint64 = (v49 ^ 42)
                                                                            var v87 float32 = (-2.25 * v2)
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                    } else {
                                                                        var v88 uint64 = 4294967295
                                                                        fmt.Println(v59, v49, v5)
                                                                        if (!v56) {
                                                                            if true {
                                                                                if v56 {
                                                                                    if false {
                                                                                        var v89 uint64 = 0
                                                                                        v90 := GenS7{F0: (!v56), F1: uint16(v88)}
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                    }
                                                                                    var v91 int32 = 2147483647
                                                                                    if false {
                                                                                        if (!v56) {
                                                                                            switch 42 {
                                                                                            case 2147483647:
                                                                                                var v92 uint8 = (v55 * v55)
                                                                                                _ = v92
                                                                                            case 128:
                                                                                                if len(v40) > 0 {
                                                                                                    v40[0] = (fi45 | fi45)
                                                                                                }
                                                                                            case 4294967295:
                                                                                                switch fi45 {
                                                                                                case 32768:
                                                                                                    var v93 bool = v56
                                                                                                    _ = v93
                                                                                                case 2147483647:
                                                                                                    var v94 uint32 = v5
                                                                                                    _ = v94
                                                                                                default:
                                                                                                    v95 := GenS6{F0: "foo bar", F1: (575 != 2)}
                                                                                                    _ = v95
                                                                                                }
                                                                                            default:
                                                                                                var v96 int = -772
                                                                                                _ = v96
                                                                                            }
                                                                                        }
                                                                                        v59 = 2147483647
                                                                                        var v97 uint32 = 4294967295
                                                                                        _ = v97
                                                                                    }
                                                                                    _ = v91
                                                                                }
                                                                            } else {
                                                                                var v98 int16 = 314
                                                                                v99 := GenS6{F0: "foo bar", F1: v56}
                                                                                v74 += (v49 - v74)
                                                                                _ = v98
                                                                                _ = v99
                                                                            }
                                                                            var v100 uint8 = v55
                                                                            fmt.Println(v59, v1, v49)
                                                                            _ = v100
                                                                        } else {
                                                                            if true {
                                                                                v101 := GenS1{F0: 32767}
                                                                                _ = v101
                                                                            } else {
                                                                                delete(v4, (^(fi44)))
                                                                                var v102 int8 = 0
                                                                                switch fi44 {
                                                                                case -2147483648:
                                                                                    var v103 float64 = -2.25
                                                                                    _ = v103
                                                                                case 42:
                                                                                    v1 = v1
                                                                                case 2147483647:
                                                                                    fmt.Println(v6, v3, v6)
                                                                                default:
                                                                                    var v104 int = fi45
                                                                                    _ = v104
                                                                                }
                                                                                _ = v102
                                                                            }
                                                                            switch ("a" + (("a" + "hello") + "")) {
                                                                            case "\n\t":
                                                                                fmt.Println(v5, v37)
                                                                            case "foo bar":
                                                                                v105 := GenS8{F0: 128, F1: v56, F2: (v49 ^ (v74 / ((v74) | 1)))}
                                                                                _ = v105
                                                                            }
                                                                        }
                                                                        var v106 float64 = float64(v2)
                                                                        v1 = v1
                                                                        _ = v106
                                                                        _ = v88
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        delete(v4, int(v5))
                                                                    case "foo bar":
                                                                        v37 -= 65535
                                                                    case "hello":
                                                                        v74 += v74
                                                                    }
                                                                } else {
                                                                    var v107 uint16 = 942
                                                                    if false {
                                                                        _ = len(v4)
                                                                    } else {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v59 += (v59 * -2147483648)
                                                                            if v56 {
                                                                                var v108 uint32 = v5
                                                                                _ = v108
                                                                            } else {
                                                                                v6 = v6
                                                                                var v109 uint32 = uint32(v107)
                                                                                _ = len(v40)
                                                                                _ = v109
                                                                            }
                                                                            fmt.Println(v107)
                                                                        case "\n\t":
                                                                            if false {
                                                                                v59 = (v59 | v59)
                                                                            }
                                                                        default:
                                                                            switch "a" {
                                                                            case "":
                                                                                v2 -= (1e10 / (v2 * (v2 * (v2 / (((v2 * 0.0) - v2) / ((((v2 + float32(1e10)) / 3.14159) / (1.5 * (v2 / 1.0))) - (((1.5 * -2.25) - float32(1.5)) / 1.5)))))))
                                                                            case "a":
                                                                                switch (("hello" + ("hello" + "foo bar")) + "a") {
                                                                                case "hello":
                                                                                    if (1000 <= -1) {
                                                                                        switch "foo bar" {
                                                                                        case "":
                                                                                            fmt.Println(v6, fi44, fi45)
                                                                                        default:
                                                                                            if len(v40) > 0 {
                                                                                                v40[0] = fi45
                                                                                            }
                                                                                            var v110 uint64 = 155
                                                                                            _ = v110
                                                                                        }
                                                                                    } else {
                                                                                        if v56 {
                                                                                            var v111 bool = ((4294967295 > v74) && (int16(v3) >= -1))
                                                                                            v56 = (1e10 == float64((0.1 + (((v2 * (v2 - 0.0)) * (float32((3.14159 + 0.1)) + (v2 + ((v2 * (v2 / -2.25)) - v2)))) / ((v2 + (float32((-2.25 + 1e10)) - 1e10)) / 1.0)))))
                                                                                            var v112 uint = v1
                                                                                            if ((v59 + v59) == v59) {
                                                                                                fmt.Println(v6, v107)
                                                                                                fmt.Println(fi45)
                                                                                                if (0 <= int32(v5)) {
                                                                                                    fmt.Println(v6)
                                                                                                    fmt.Println(fi44, v111)
                                                                                                    var v113 int16 = -161
                                                                                                    _ = v113
                                                                                                } else {
                                                                                                    if v56 {
                                                                                                        v107 = (v107 / ((v107) | 1))
                                                                                                        v40 = append(v40, int(v5))
                                                                                                    }
                                                                                                    if (!v56) {
                                                                                                        switch ("\n\t" + "hello") {
                                                                                                        case "\n\t":
                                                                                                            var v114 uint64 = (v49 & v49)
                                                                                                            _ = v114
                                                                                                        case "":
                                                                                                            v115 := GenS4{F0: int32(v6)}
                                                                                                            _ = v115
                                                                                                        }
                                                                                                        var v116 float64 = float64(3.14159)
                                                                                                        switch ("\n\t" + "") {
                                                                                                        case "a":
                                                                                                            var v117 int = int(v112)
                                                                                                            _ = v117
                                                                                                        case "\n\t":
                                                                                                            var v118 float64 = ((v116 / (v116 - float64(((v2 + ((((-2.25 * float32(((v116 / float64(0.1)) * 0.1))) / (3.14159 + 1e10)) - v2) + ((float32(v116) + v2) * -2.25))) - v2)))) + (v116 / (v116 - 1.5)))
                                                                                                            _ = v118
                                                                                                        case "":
                                                                                                            switch (fi44 | fi45) {
                                                                                                            case 1:
                                                                                                                var v119 int32 = 2147483647
                                                                                                                _ = v119
                                                                                                            case 32768:
                                                                                                                if false {
                                                                                                                    v40 = append(v40, fi44)
                                                                                                                    var v120 uint64 = v49
                                                                                                                    var v121 string = "foo bar"
                                                                                                                    _ = v120
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v116
                                                                                                    }
                                                                                                    v122 := GenS8{F0: 255, F1: v111, F2: uint64(v112)}
                                                                                                    _ = v122
                                                                                                }
                                                                                            }
                                                                                            if len(v40) > 0 {
                                                                                                v40[0] = (fi44 ^ fi44)
                                                                                            }
                                                                                            v123 := v4[-1]
                                                                                            if true {
                                                                                                v56 = v111
                                                                                                switch "" {
                                                                                                case "\n\t":
                                                                                                    if v56 {
                                                                                                        v1 = (v1 * 486)
                                                                                                        if v111 {
                                                                                                            v1 = (v1 + v1)
                                                                                                            var v124 int16 = 1
                                                                                                            _ = v124
                                                                                                        }
                                                                                                        var v125 uint = v1
                                                                                                        _ = v125
                                                                                                    }
                                                                                                default:
                                                                                                    v2 -= -2.25
                                                                                                }
                                                                                            } else {
                                                                                                v126 := GenS9{F0: v123}
                                                                                                v127 := GenS9{F0: (v123 + v123)}
                                                                                                v107 -= (^(v107))
                                                                                                _ = v126
                                                                                                _ = v127
                                                                                            }
                                                                                            _ = v111
                                                                                            _ = v112
                                                                                            _ = v123
                                                                                        } else {
                                                                                            var v128 uint8 = v55
                                                                                            switch 2147483647 {
                                                                                            case -1:
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    if (float64((float32(1.5) + 1e10)) > 1e10) {
                                                                                                        var v129 int = int(v59)
                                                                                                        v107 -= v107
                                                                                                        v130 := GenS4{F0: 0}
                                                                                                        var v131 uint = v1
                                                                                                        v59 -= (v59 + (v59 & (v59 << ((v131) & 63))))
                                                                                                        v132 := GenS10{F0: 833, F1: 128, F2: (v5 | (v37 | v37))}
                                                                                                        var v133 uint8 = (v128 ^ v55)
                                                                                                        _ = v129
                                                                                                        _ = v130
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                        _ = v133
                                                                                                    }
                                                                                                }
                                                                                                var v134 int8 = -1
                                                                                                _ = v134
                                                                                            case -2147483648:
                                                                                                var v135 uint16 = (v107 | (v107 >> ((18446744073709551615) & 63)))
                                                                                                _ = v135
                                                                                            }
                                                                                            _ = v128
                                                                                        }
                                                                                        v136 := GenS11{F0: int8(v3)}
                                                                                        _ = v136
                                                                                    }
                                                                                case "":
                                                                                    v137 := GenS12{F0: v6}
                                                                                    _ = v137
                                                                                }
                                                                                var v138 string = "a"
                                                                                _ = v138
                                                                            case "\n\t":
                                                                                v139 := GenS13{F0: 2}
                                                                                _ = v139
                                                                            }
                                                                        }
                                                                        v6 += 51
                                                                        var v140 int = 32768
                                                                        _ = v140
                                                                    }
                                                                    fmt.Println(v59)
                                                                    _ = v107
                                                                }
                                                                _ = v75
                                                            } else {
                                                                if false {
                                                                    if true {
                                                                        v141 := GenS14{F0: (v56 && false), F1: 0}
                                                                        v142 := GenS9{F0: "a"}
                                                                        v56 = ((0.1 - 0.0) >= 1.5)
                                                                        _ = v141
                                                                        _ = v142
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v56, v2)
                                                                        v74 += v74
                                                                        v49 += (v49 ^ (v74 & (v74 / (((v49 ^ 42)) | 1))))
                                                                    } else {
                                                                        var v143 int64 = (v3 << ((uint(v55)) & 63))
                                                                        _ = v143
                                                                    }
                                                                    v6 = 127
                                                                } else {
                                                                    if true {
                                                                        if (v2 > ((0.0 - (v2 - ((((0.0 + (((float32((1e10 - 0.1)) + v2) + -2.25) / v2)) + 0.0) * v2) - ((v2 * (1.5 * v2)) + (0.0 * (1e10 / float32(0.1))))))) - v2)) {
                                                                            if v56 {
                                                                                fmt.Println(v74, v49)
                                                                                v144 := v4[-972]
                                                                                v3 = (-(v3))
                                                                                _ = v144
                                                                            }
                                                                            if false {
                                                                                if ((v56 && (!(v56 || v56))) && (uint16(v1) <= 0)) {
                                                                                    if false {
                                                                                        v145 := GenS4{F0: -284}
                                                                                        var v146 uint = 401
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                    }
                                                                                }
                                                                                var v147 bool = true
                                                                                v148 := v4[int(v5)]
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                            var v149 int = -606
                                                                            _ = v149
                                                                        } else {
                                                                            v150 := GenS15{F0: 1, F1: v55}
                                                                            _ = v150
                                                                        }
                                                                        fmt.Println(v5, v55)
                                                                    } else {
                                                                        v37 = v5
                                                                        fmt.Println(v59, v56)
                                                                        switch fi44 {
                                                                        case 0:
                                                                            _ = len(v40)
                                                                            switch ("a" + ("" + "a")) {
                                                                            case "a":
                                                                                v151 := GenS5{F0: int16(v3), F1: v5, F2: v6}
                                                                                _ = v151
                                                                            case "\n\t":
                                                                                if len(v40) >= 2 {
                                                                                    v40 = v40[1:len(v40)]
                                                                                }
                                                                            case "":
                                                                                var v152 uint8 = (v55 + (v55 - (v55 ^ 42)))
                                                                                _ = v152
                                                                            }
                                                                            v153 := GenS10{F0: -369, F1: (fi45 * (^(fi44))), F2: v5}
                                                                            _ = v153
                                                                        case 4294967295:
                                                                            v56 = ((v56 || v56) || ((0.1 >= -2.25) && false))
                                                                        case 1:
                                                                            v154 := GenS16{F0: (v59 << (((^(v1))) & 63)), F1: (^(v74))}
                                                                            _ = v154
                                                                        }
                                                                    }
                                                                    var v155 string = "foo bar"
                                                                    _ = v155
                                                                }
                                                                fmt.Println(v1, v59, fi44)
                                                            }
                                                            v56 = (false && v56)
                                                            _ = v74
                                                        }
                                                        _ = v59
                                                    } else {
                                                        fmt.Println(v55, v37)
                                                        v5 += (v5 & (v5 % (((v37 % ((v5) | 1))) | 1)))
                                                    }
                                                } else {
                                                    if false {
                                                        v1 = v1
                                                        if (1 >= v1) {
                                                            var v156 uint16 = 996
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if v56 {
                                                                    var v157 uint16 = (v156 | 210)
                                                                    _ = v157
                                                                } else {
                                                                    fmt.Println(v37, v37, v6)
                                                                    var v158 uint64 = 18446744073709551615
                                                                    var v159 float64 = 0.0
                                                                    _ = v158
                                                                    _ = v159
                                                                }
                                                            }
                                                            v2 -= v2
                                                            _ = v156
                                                        }
                                                        v160 := GenS14{F0: v56, F1: v6}
                                                        _ = v160
                                                    }
                                                    v161 := GenS17{F0: -944, F1: (v49 % ((v49) | 1))}
                                                    _ = v161
                                                }
                                            } else {
                                                v162 := GenS5{F0: -377, F1: (^(v37)), F2: (v6 | v6)}
                                                var v163 int8 = v6
                                                var v164 float64 = -2.25
                                                if ((!(1 <= -946)) && (42 != -32768)) {
                                                    var v165 int8 = -128
                                                    _ = v165
                                                } else {
                                                    v56 = true
                                                    if true {
                                                        if (float64(v2) <= ((1e10 + v164) * v164)) {
                                                            fmt.Println(fi45)
                                                            if v56 {
                                                                var v166 int64 = v3
                                                                if len(v40) >= 2 {
                                                                    v40 = v40[1:len(v40)]
                                                                }
                                                                v55 = (v55 << (((v1 | (v1 + (^(v1))))) & 63))
                                                                _ = v166
                                                            } else {
                                                                switch "\n\t" {
                                                                case "":
                                                                    v167 := GenS14{F0: v56, F1: v163}
                                                                    _ = v167
                                                                case "a":
                                                                    var v168 int32 = 0
                                                                    _ = v168
                                                                }
                                                                var v169 string = ("hello" + "")
                                                                _ = v169
                                                            }
                                                        }
                                                    } else {
                                                        fmt.Println(v5, fi45, v2)
                                                        v170 := GenS14{F0: (2147483647 > -168), F1: v6}
                                                        fmt.Println(v6, v5, v5)
                                                        _ = v170
                                                    }
                                                }
                                                switch (^(fi44)) {
                                                case -732:
                                                    var v171 int32 = 0
                                                    _ = v171
                                                case 427:
                                                    if false {
                                                        var v172 int = fi44
                                                        v173 := GenS16{F0: int32(v49), F1: 30}
                                                        _ = v172
                                                        _ = v173
                                                    }
                                                }
                                                _ = v162
                                                _ = v163
                                                _ = v164
                                            }
                                            var v174 uint16 = 0
                                            _ = v174
                                            _ = v55
                                            _ = v56
                                        }
                                        var v175 uint64 = 2
                                        v176 := GenS14{F0: false, F1: v6}
                                        _ = v175
                                        _ = v176
                                    } else {
                                        var v177 bool = (!(32767 > -1))
                                        if v177 {
                                            var v178 uint = v1
                                            v6 = int8(v49)
                                            v4[-1] = "foo bar"
                                            var v179 float32 = ((v2 / v2) + v2)
                                            if (!v177) {
                                                v180 := GenS13{F0: 2}
                                                if false {
                                                    _ = len(v40)
                                                    if len(v40) >= 2 {
                                                        v40 = v40[1:len(v40)]
                                                    }
                                                }
                                                switch fi44 {
                                                case 142:
                                                    v179 += ((v179 / v2) * 0.1)
                                                case -131:
                                                    v181 := GenS17{F0: 1000, F1: uint64(v1)}
                                                    _ = v181
                                                case -2147483648:
                                                    v1 += v1
                                                }
                                                _ = v180
                                            }
                                            var v182 uint = uint(v49)
                                            if v177 {
                                                if v177 {
                                                    v2 += -2.25
                                                    var v183 uint64 = 1
                                                    var v184 string = ((("hello" + "a") + "\n\t") + ("\n\t" + ("hello" + "\n\t")))
                                                    _ = v183
                                                    _ = v184
                                                }
                                                fmt.Println(v49)
                                            } else {
                                                v5 += v5
                                                var v185 float32 = v2
                                                _ = v185
                                            }
                                            fmt.Println(v37, v5)
                                            _ = v178
                                            _ = v179
                                            _ = v182
                                        }
                                        _ = v177
                                    }
                                    _ = len(v4)
                                    _ = v49
                                }
                            }
                        }
                        v1 += 547
                        _ = len(v4)
                        _ = fi45
                    }
                    if (!(((!true) && (1 > 127)) && ((false || (!true)) && (v6 <= v6)))) {
                        for fi186 := 0; fi186 < 2; fi186++ {
                            if (3.14159 == 0.1) {
                                v187 := GenS17{F0: 1, F1: 569}
                                v188 := GenS18{F0: ("foo bar" + ""), F1: 2}
                                _ = v187
                                _ = v188
                            }
                            _ = fi186
                        }
                    }
                    v189 := GenS8{F0: 2, F1: (42 < 42), F2: 631}
                    _ = v189
                }
            }
            var v190 uint8 = 255
            switch fi44 {
            case 127:
                fmt.Println(v6, v190, v1)
            case -2147483648:
                v191 := GenS7{F0: true, F1: 79}
                _ = v191
            }
            _ = fi44
            _ = v190
        }
    } else {
        switch (("a" + "hello") + "") {
        case "hello":
            fmt.Println(v5)
        default:
            v1 = v1
        }
        v192 := GenS19{F0: 754, F1: (("" + "") + "hello"), F2: -955}
        switch "\n\t" {
        case "\n\t":
            var v193 string = (("\n\t" + "\n\t") + "\n\t")
            _ = v193
        case "hello":
            var v194 uint = (v1 / ((18446744073709551615) | 1))
            _ = v194
        }
        _ = v192
    }
    if false {
        var v195 bool = (-2147483648 != -769)
        var v196 uint = (v1 * v1)
        if (248 > 42) {
            v197 := GenS2{F0: 1000, F1: ("a" + "a"), F2: (v196 * v196)}
            _ = v197
        }
        _ = v195
        _ = v196
    } else {
        v198 := GenS15{F0: 65535, F1: 128}
        fmt.Println(v6)
        _ = v198
    }
    v6 += v6
    v199 := GenS19{F0: 393, F1: "hello", F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v199
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
