package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint64; F1 int16 }
type GenS11 struct { F0 int32; F1 uint; F2 uint8 }
type GenS12 struct { F0 uint32; F1 bool }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 int32; F1 uint; F2 uint8 }
type GenS15 struct { F0 bool; F1 int64 }
type GenS16 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]int)
    var v3 uint8 = 2
    v4 := []uint32{1, 2}
    var v5 float32 = (3.14159 + -2.25)
    v6 := make(map[string]string)
    v2[882] = 9223372036854775807
    switch 276 {
    case 128:
        var v7 uint16 = (^(v1))
        _ = v7
    }
    switch "a" {
    case "hello":
        for fi8 := 0; fi8 < 15; fi8++ {
            if (!((fi8 * (fi8 - 2147483647)) <= fi8)) {
                fmt.Println(v5, v5)
                for fi9 := 0; fi9 < 7; fi9++ {
                    var v10 int16 = -112
                    fmt.Println(v10)
                    v11 := GenS1{F0: v1, F1: 127, F2: "\n\t"}
                    delete(v6, (("" + ("" + ("hello" + ""))) + "foo bar"))
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if len(v4) > 0 {
                    v4[0] = uint32(v3)
                }
            }
            var v12 uint16 = v1
            v1 += v1
            _ = fi8
            _ = v12
        }
    default:
        if (-1 == 341) {
            if false {
                v1 -= (v1 >> ((65535) & 63))
                if (v3 != v3) {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if (0 == 32767) {
                            if true {
                                var v14 int32 = 813
                                _ = v14
                            }
                            v3 -= 1
                            fmt.Println(v5, v3, v1)
                        } else {
                            v15 := GenS1{F0: v1, F1: 37, F2: ""}
                            var v16 float64 = 1.5
                            v1 += v1
                            _ = v15
                            _ = v16
                        }
                        v4 = append(v4, uint32(v3))
                        var v17 int16 = 818
                        _ = fi13
                        _ = v17
                    }
                    _ = len(v4)
                    if (!true) {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            for fi19 := 0; fi19 < 10; fi19++ {
                                var v20 int32 = -2147483648
                                _ = fi19
                                _ = v20
                            }
                            v21 := GenS1{F0: (v1 / ((v1) | 1)), F1: 1, F2: ("a" + "hello")}
                            _ = fi18
                            _ = v21
                        }
                        v22 := GenS1{F0: v1, F1: -1, F2: "foo bar"}
                        _ = v22
                    } else {
                        for fi23 := 0; fi23 < 11; fi23++ {
                            var v24 int16 = 0
                            v5 = 0.0
                            for fi25 := 0; fi25 < 19; fi25++ {
                                v26 := GenS2{F0: 70, F1: "foo bar"}
                                switch (fi25 & 390) {
                                case 128:
                                    var v27 float64 = (3.14159 - 1e10)
                                    _ = v27
                                default:
                                    switch ("foo bar" + "foo bar") {
                                    case "\n\t":
                                        v28 := GenS2{F0: 695, F1: ("hello" + "foo bar")}
                                        _ = v28
                                    case "foo bar":
                                        v24 = v24
                                    }
                                    var v29 int64 = 1
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                            }
                            _ = fi23
                            _ = v24
                        }
                        _ = len(v4)
                    }
                }
            } else {
                var v30 int8 = -128
                v30 += (v30 ^ (v30 >> ((4294967295) & 63)))
                _ = v30
            }
            if true {
                var v31 float32 = (0.0 * 1.5)
                var v32 float64 = (0.1 + 1.5)
                _ = v31
                _ = v32
            }
            if false {
                if false {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v5 = (v5 * float32(1e10))
                        v34 := GenS2{F0: 0, F1: ("" + "\n\t")}
                        _ = fi33
                        _ = v34
                    }
                    var v35 int8 = 67
                    _ = v35
                } else {
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                    v1 = v1
                    switch "a" {
                    case "hello":
                        for fi36 := 0; fi36 < 7; fi36++ {
                            if (-2.25 != (float64(float32(1.5)) + -2.25)) {
                                switch "foo bar" {
                                case "foo bar":
                                    for fi37 := 0; fi37 < 11; fi37++ {
                                        switch "" {
                                        case "\n\t":
                                            var v38 float32 = (v5 - 0.1)
                                            _ = v38
                                        case "a":
                                            v39 := GenS3{F0: 255}
                                            if false {
                                                v1 += (v1 << ((4294967295) & 63))
                                            } else {
                                                v2[0] = 128
                                                var v40 uint = 852
                                                v41 := GenS4{F0: 2147483647, F1: -2147483648, F2: 260}
                                                v5 += ((0.1 * v5) + v5)
                                                _ = v40
                                                _ = v41
                                            }
                                            v5 = 0.0
                                            switch ("foo bar" + "") {
                                            case "":
                                                v42 := GenS3{F0: 18446744073709551615}
                                                _ = v42
                                            case "hello":
                                                switch ((("\n\t" + "hello") + ("hello" + "hello")) + "hello") {
                                                case "\n\t":
                                                    var v43 uint = 951
                                                    _ = v43
                                                }
                                            }
                                            switch ("\n\t" + "\n\t") {
                                            case "a":
                                                if (false && (526 > 32767)) {
                                                    v5 = v5
                                                }
                                            case "foo bar":
                                                v1 += v1
                                            default:
                                                delete(v2, 2)
                                            }
                                            if (!(9223372036854775807 <= 9223372036854775807)) {
                                                fmt.Println(v1)
                                            }
                                            switch fi36 {
                                            case 9223372036854775807:
                                                v1 = v1
                                            case 32767:
                                                var v44 uint16 = 790
                                                _ = v44
                                            default:
                                                delete(v6, ("a" + ("hello" + "a")))
                                            }
                                            _ = v39
                                        case "hello":
                                            fmt.Println(v5)
                                        default:
                                            if (true || ((fi37 > (fi36 % ((127) | 1))) || true)) {
                                                v45 := GenS2{F0: 2, F1: ("a" + "hello")}
                                                var v46 int = 32768
                                                _ = v45
                                                _ = v46
                                            } else {
                                                fmt.Println(v1)
                                                var v47 uint = 42
                                                _ = v47
                                            }
                                        }
                                        _ = fi37
                                    }
                                case "hello":
                                    v2[2147483647] = fi36
                                case "":
                                    switch fi36 {
                                    case 0:
                                        var v48 string = "a"
                                        _ = v48
                                    case 32767:
                                        if true {
                                            if true {
                                                v49 := GenS5{F0: 127}
                                                _ = v49
                                            } else {
                                                if false {
                                                    switch ("foo bar" + ("hello" + "foo bar")) {
                                                    case "hello":
                                                        var v50 int8 = -84
                                                        _ = v50
                                                    }
                                                } else {
                                                    for fi51 := 0; fi51 < 12; fi51++ {
                                                        if false {
                                                            if true {
                                                                var v52 bool = false
                                                                _ = v52
                                                            }
                                                            var v53 uint64 = 4294967295
                                                            v4 = append(v4, 2147483648)
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS6{F0: v3}
                                                            v1 = v1
                                                            var v55 uint32 = 2
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                        switch ("a" + "hello") {
                                                        case "":
                                                            v4 = append(v4, 65535)
                                                        case "hello":
                                                            v56 := GenS2{F0: 0, F1: "foo bar"}
                                                            _ = v56
                                                        }
                                                        _ = fi51
                                                    }
                                                    v57 := GenS4{F0: 823, F1: -1, F2: 2}
                                                    _ = v57
                                                }
                                                if true {
                                                    if ((("hello" == "hello") && ((!(-595 > (fi36 % (((^(fi36))) | 1)))) && ((!(int32(fi36) > 2147483647)) && (0 == 42)))) || (true && true)) {
                                                        v58 := GenS1{F0: v1, F1: -128, F2: (("hello" + ((("foo bar" + (("" + "foo bar") + "")) + "a") + (("\n\t" + ("a" + "hello")) + "hello"))) + (("foo bar" + (((("\n\t" + "a") + "a") + "hello") + ("" + ("foo bar" + "foo bar")))) + "\n\t"))}
                                                        if ("hello" == ("hello" + ("foo bar" + ""))) {
                                                            var v59 uint16 = v1
                                                            var v60 float64 = 1.5
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v5 += (3.14159 / v5)
                                                            if (39 >= 577) {
                                                                var v61 int = 32768
                                                                _ = len(v2)
                                                                _ = v61
                                                            } else {
                                                                v1 = v1
                                                                if (2147483647 == -2147483648) {
                                                                    switch "hello" {
                                                                    case "foo bar":
                                                                        var v62 uint64 = 346
                                                                        _ = v62
                                                                    case "\n\t":
                                                                        for fi63 := 0; fi63 < 11; fi63++ {
                                                                            v3 += (v3 + uint8(v1))
                                                                            v64 := GenS5{F0: fi63}
                                                                            var v65 int32 = -2147483648
                                                                            _ = fi63
                                                                            _ = v64
                                                                            _ = v65
                                                                        }
                                                                    case "":
                                                                        var v66 uint16 = v1
                                                                        _ = v66
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        _ = v58
                                                    } else {
                                                        var v67 bool = (!false)
                                                        for fi68 := 0; fi68 < 16; fi68++ {
                                                            v4 = append(v4, 1)
                                                            if (2147483647 >= 4294967295) {
                                                                if true {
                                                                    if true {
                                                                        v5 += ((float32(3.14159) / float32((1e10 - 0.0))) / 1.0)
                                                                        switch (("foo bar" + "\n\t") + "hello") {
                                                                        case "\n\t":
                                                                            fmt.Println(v5)
                                                                        case "foo bar":
                                                                            switch ("\n\t" + "foo bar") {
                                                                            case "\n\t":
                                                                                if (-661 == 1) {
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        var v69 uint8 = (v3 - 94)
                                                                                        _ = v69
                                                                                    case "\n\t":
                                                                                        fmt.Println(fi36, v5)
                                                                                    case "a":
                                                                                        var v70 string = ("\n\t" + (((("foo bar" + "a") + "") + "a") + "foo bar"))
                                                                                        _ = v70
                                                                                    default:
                                                                                        v71 := GenS5{F0: (^(fi36))}
                                                                                        _ = v71
                                                                                    }
                                                                                    v4 = append(v4, 929)
                                                                                    if v67 {
                                                                                        v5 -= 0.0
                                                                                        if (!(v67 || (true && true))) {
                                                                                            v72 := GenS7{F0: (1e10 > 1e10)}
                                                                                            var v73 uint = 2
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        } else {
                                                                                            var v74 uint = 753
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            v3 -= 247
                                                                                            v75 := GenS6{F0: (^(v3))}
                                                                                            if true {
                                                                                                switch fi36 {
                                                                                                case -2147483648:
                                                                                                    if v67 {
                                                                                                        v2[2] = (^(fi36))
                                                                                                        v67 = v67
                                                                                                    } else {
                                                                                                        v2[2147483648] = fi68
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            if v67 {
                                                                                                                if (!(!false)) {
                                                                                                                    var v76 uint64 = 0
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                            } else {
                                                                                                                v3 = (v3 / ((v3) | 1))
                                                                                                                var v77 uint = v74
                                                                                                                if false {
                                                                                                                    var v78 int = fi36
                                                                                                                    _ = v78
                                                                                                                } else {
                                                                                                                    var v79 uint32 = 4294967295
                                                                                                                    v3 = v3
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                                v74 -= v77
                                                                                                                switch (fi36 % (((fi36 / ((fi36) | 1))) | 1)) {
                                                                                                                case 4294967295:
                                                                                                                    fmt.Println(v5, fi68, v67)
                                                                                                                default:
                                                                                                                    var v80 uint64 = 590
                                                                                                                    _ = v80
                                                                                                                }
                                                                                                                _ = v77
                                                                                                            }
                                                                                                        case "foo bar":
                                                                                                            var v81 uint16 = (v1 + 1)
                                                                                                            _ = v81
                                                                                                        case "\n\t":
                                                                                                            var v82 uint16 = v1
                                                                                                            _ = v82
                                                                                                        }
                                                                                                    }
                                                                                                    var v83 int = (^(fi68))
                                                                                                    _ = v83
                                                                                                default:
                                                                                                    if false {
                                                                                                        var v84 int32 = 771
                                                                                                        fmt.Println(v5, fi68)
                                                                                                        var v85 int8 = -115
                                                                                                        var v86 uint64 = 18446744073709551615
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                        _ = v86
                                                                                                    } else {
                                                                                                        v1 = (v1 / ((v1) | 1))
                                                                                                        var v87 uint16 = v1
                                                                                                        _ = v87
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            var v88 int = (fi68 - int(v3))
                                                                                            var v89 int32 = 1
                                                                                            if false {
                                                                                                v90 := GenS8{F0: -32768}
                                                                                                var v91 uint16 = (^(v1))
                                                                                                _ = v90
                                                                                                _ = v91
                                                                                            } else {
                                                                                                v5 = v5
                                                                                                switch (v88 % ((fi68) | 1)) {
                                                                                                case 32768:
                                                                                                    switch ("a" + "\n\t") {
                                                                                                    case "hello":
                                                                                                        v3 -= (v3 % ((v3) | 1))
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    var v92 string = "foo bar"
                                                                                                    if (((v5 + 3.14159) != v5) && true) {
                                                                                                        var v93 int = fi68
                                                                                                        switch "a" {
                                                                                                        case "a":
                                                                                                            v89 += v89
                                                                                                        default:
                                                                                                            fmt.Println(v93, v74)
                                                                                                        }
                                                                                                        v94 := GenS6{F0: (v3 % ((v3) | 1))}
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    _ = v92
                                                                                                }
                                                                                                switch "" {
                                                                                                case "a":
                                                                                                    fmt.Println(v67, v67, fi36)
                                                                                                }
                                                                                            }
                                                                                            if true {
                                                                                                v1 = v1
                                                                                            } else {
                                                                                                if (!("\n\t" < ("\n\t" + "\n\t"))) {
                                                                                                    var v95 float32 = (v5 - ((v5 + v5) * 0.0))
                                                                                                    var v96 uint64 = 340
                                                                                                    _ = v95
                                                                                                    _ = v96
                                                                                                }
                                                                                                v97 := v6[""]
                                                                                                v5 += v5
                                                                                                _ = v97
                                                                                            }
                                                                                            v98 := GenS9{F0: 594, F1: v88}
                                                                                            _ = v74
                                                                                            _ = v75
                                                                                            _ = v88
                                                                                            _ = v89
                                                                                            _ = v98
                                                                                        }
                                                                                        v1 = (v1 << ((548) & 63))
                                                                                    }
                                                                                    v67 = true
                                                                                    v99 := GenS10{F0: 434, F1: 1}
                                                                                    v3 = v3
                                                                                    _ = v99
                                                                                }
                                                                                var v100 int8 = -93
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 2147483648
                                                                                }
                                                                                _ = v100
                                                                            case "hello":
                                                                                if (1 >= 9223372036854775807) {
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        var v101 int8 = -128
                                                                                        _ = v101
                                                                                    case "\n\t":
                                                                                        v102 := GenS11{F0: 688, F1: 89, F2: v3}
                                                                                        _ = v102
                                                                                    case "hello":
                                                                                        v4 = append(v4, 65535)
                                                                                    }
                                                                                    var v103 bool = v67
                                                                                    _ = v103
                                                                                }
                                                                            default:
                                                                                fmt.Println(v3)
                                                                            }
                                                                        case "":
                                                                            var v104 uint = 65535
                                                                            _ = v104
                                                                        }
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 2147483648
                                                                    }
                                                                    if v67 {
                                                                        var v105 string = ("" + ("foo bar" + "\n\t"))
                                                                        if v67 {
                                                                            var v106 string = v105
                                                                            switch v105 {
                                                                            case "foo bar":
                                                                                v1 -= v1
                                                                            case "a":
                                                                                v107 := v6[""]
                                                                                _ = v107
                                                                            case "\n\t":
                                                                                var v108 uint64 = 2
                                                                                _ = v108
                                                                            }
                                                                            _ = v106
                                                                        } else {
                                                                            v109 := GenS5{F0: -2147483648}
                                                                            if v67 {
                                                                                _ = len(v4)
                                                                                switch (v105 + v105) {
                                                                                case "a":
                                                                                    _ = len(v6)
                                                                                case "hello":
                                                                                    v5 -= (1e10 - v5)
                                                                                case "foo bar":
                                                                                    var v110 int8 = 0
                                                                                    _ = v110
                                                                                default:
                                                                                    var v111 int16 = 42
                                                                                    _ = v111
                                                                                }
                                                                                _ = len(v2)
                                                                            }
                                                                            if (100 < int8(v3)) {
                                                                                v1 = (v1 - (v1 & 0))
                                                                            }
                                                                            _ = v109
                                                                        }
                                                                        switch (((("\n\t" + "a") + v105) + "foo bar") + "\n\t") {
                                                                        case "hello":
                                                                            v112 := GenS10{F0: 2, F1: -1}
                                                                            _ = v112
                                                                        case "a":
                                                                            switch (fi36 | (fi36 & -675)) {
                                                                            case 127:
                                                                                v113 := GenS12{F0: 0, F1: v67}
                                                                                _ = v113
                                                                            default:
                                                                                v114 := GenS2{F0: 248, F1: v105}
                                                                                _ = v114
                                                                            }
                                                                        default:
                                                                            v3 = v3
                                                                        }
                                                                        _ = v105
                                                                    }
                                                                } else {
                                                                    fmt.Println(v5, v1, fi36)
                                                                }
                                                                var v115 uint = 846
                                                                delete(v2, 2147483648)
                                                                _ = v115
                                                            }
                                                            if v67 {
                                                                _ = len(v6)
                                                                v5 = float32(0.0)
                                                                _ = len(v4)
                                                            }
                                                            v116 := GenS6{F0: v3}
                                                            v3 = 131
                                                            _ = fi68
                                                            _ = v116
                                                        }
                                                        _ = v67
                                                    }
                                                    fmt.Println(v3, v5, fi36)
                                                }
                                                v117 := GenS9{F0: 4294967295, F1: fi36}
                                                _ = v117
                                            }
                                            v118 := GenS3{F0: 18446744073709551615}
                                            var v119 uint32 = 42
                                            _ = v118
                                            _ = v119
                                        }
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                }
                                var v120 float64 = (3.14159 * 1.5)
                                _ = v120
                            } else {
                                var v121 uint = 4294967295
                                _ = v121
                            }
                            for fi122 := 0; fi122 < 13; fi122++ {
                                v123 := GenS10{F0: 18446744073709551615, F1: 42}
                                if len(v4) > 0 {
                                    v4[0] = 805
                                }
                                var v124 uint32 = 937
                                _ = fi122
                                _ = v123
                                _ = v124
                            }
                            if false {
                                v125 := GenS1{F0: (v1 << ((500) & 63)), F1: -99, F2: "\n\t"}
                                v5 -= v5
                                v126 := GenS7{F0: true}
                                _ = v125
                                _ = v126
                            } else {
                                v5 = 3.14159
                                for fi127 := 0; fi127 < 18; fi127++ {
                                    if (((!((-814 <= 42) && (true || (1 > -63)))) && true) && true) {
                                        fmt.Println(v5, v1)
                                        var v128 float64 = -2.25
                                        _ = v128
                                    } else {
                                        var v129 bool = true
                                        _ = v129
                                    }
                                    _ = fi127
                                }
                                for fi130 := 0; fi130 < 9; fi130++ {
                                    var v131 uint64 = 0
                                    _ = fi130
                                    _ = v131
                                }
                            }
                            _ = fi36
                        }
                    case "\n\t":
                        switch "foo bar" {
                        case "a":
                            switch "foo bar" {
                            case "\n\t":
                                var v132 int16 = 88
                                _ = v132
                            default:
                                for fi133 := 0; fi133 < 3; fi133++ {
                                    var v134 uint32 = 2
                                    _ = fi133
                                    _ = v134
                                }
                            }
                        case "foo bar":
                            if ((true || false) || (!(false || false))) {
                                switch -530 {
                                case 24:
                                    v135 := GenS13{F0: uint32(v1)}
                                    _ = v135
                                case 298:
                                    if (-578 <= 999) {
                                        v136 := GenS14{F0: -570, F1: 2, F2: (v3 ^ (v3 % (((^(v3))) | 1)))}
                                        switch 4294967295 {
                                        case 128:
                                            var v137 float64 = 3.14159
                                            _ = v137
                                        case 2147483647:
                                            v4 = append(v4, 2147483648)
                                        default:
                                            var v138 string = "hello"
                                            _ = v138
                                        }
                                        var v139 uint64 = 200
                                        _ = v136
                                        _ = v139
                                    } else {
                                        fmt.Println(v5)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        var v140 uint64 = 2
                                        _ = v140
                                    }
                                }
                                v3 = v3
                                for fi141 := 0; fi141 < 11; fi141++ {
                                    if true {
                                        var v142 int32 = 42
                                        v143 := GenS6{F0: (^(v3))}
                                        switch 32767 {
                                        case -974:
                                            v5 = (v5 * v5)
                                        case 0:
                                            var v144 int32 = (v142 % ((v142) | 1))
                                            _ = v144
                                        default:
                                            v145 := v6["a"]
                                            _ = v145
                                        }
                                        _ = v142
                                        _ = v143
                                    }
                                    _ = fi141
                                }
                            } else {
                                var v146 uint64 = 0
                                v147 := GenS15{F0: (1 > -1), F1: 679}
                                _ = v146
                                _ = v147
                            }
                        default:
                            var v148 uint16 = 42
                            _ = v148
                        }
                    case "foo bar":
                        v149 := GenS16{F0: -32768, F1: -2147483648, F2: 2}
                        _ = v149
                    }
                }
                fmt.Println(v1, v3, v3)
                v3 = 33
            } else {
                v150 := GenS2{F0: 256, F1: "foo bar"}
                _ = v150
            }
        }
    }
    v1 = v1
    v1 = v1
    v1 = v1
    v151 := GenS11{F0: 2147483647, F1: 2, F2: v3}
    var v152 float32 = (float32(0.0) / v5)
    v152 += 3.14159
    v3 = 115
    v153 := GenS4{F0: 2, F1: -2147483648, F2: 265}
    fmt.Println("v1:", v1)
    fmt.Println("v152:", v152)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v151
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
