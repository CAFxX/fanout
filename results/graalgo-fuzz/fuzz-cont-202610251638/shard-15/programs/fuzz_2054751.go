package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 uint8; F2 int8 }
type GenS8 struct { F0 int; F1 int; F2 int8 }
type GenS9 struct { F0 int16; F1 uint16 }
type GenS10 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 uint32; F1 int16; F2 bool }
type GenS13 struct { F0 uint16; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 int8) int8 {
    switch "hello" {
    case "foo bar":
        var v6 uint16 = 2
        _ = v6
    case "a":
        switch "hello" {
        case "hello":
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 uint8 = 127
                var v9 uint16 = 42
                switch 245 {
                case 32767:
                    v10 := GenS1{F0: ("foo bar" + ""), F1: 255}
                    _ = v10
                default:
                    fmt.Println(fi7, v9)
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        case "foo bar":
            v11 := []int{42, -9223372036854775808, -410}
            _ = v11
        }
    default:
        switch "foo bar" {
        case "\n\t":
            v12 := GenS1{F0: ("" + (("\n\t" + "\n\t") + "foo bar")), F1: 27}
            _ = v12
        }
    }
    if (2 != 0) {
        var v13 float64 = (3.14159 + 1.5)
        _ = v13
    }
    _ = p0
    _ = p1
    return (p1 << ((2) & 63))
}

func main() {
    v1 := []string{("hello" + "")}
    var v2 string = ("foo bar" + "\n\t")
    var v3 int = 32768
    v4 := []uint32{1}
    var v5 string = "hello"
    v14 := GenS2{F0: 127}
    if (126 >= 1) {
        if true {
            var v15 uint = 42
            switch (v3 & 2147483647) {
            case 128:
                v5 = "hello"
            case 0:
                if (42 > 125) {
                    v5 = (v5 + v5)
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if false {
                        v5 = ((v5 + (((v5 + "") + v2) + v2)) + v5)
                        v1 = append(v1, v2)
                        var v16 uint64 = 4294967295
                        _ = v16
                    }
                }
            case -9223372036854775808:
                var v17 uint = 0
                _ = v17
            }
            _ = v15
        }
        switch v5 {
        case "hello":
            if false {
                v18 := GenS2{F0: 255}
                _ = len(v1)
                for fi19 := 0; fi19 < 13; fi19++ {
                    v20 := GenS3{F0: -2147483648, F1: -797}
                    if true {
                        switch v3 {
                        case 4294967295:
                            var v21 uint16 = 723
                            _ = v21
                        case 653:
                            v3 += 128
                        }
                        var v22 uint32 = 185
                        if (false || (333 > 0)) {
                            var v23 float64 = 0.0
                            _ = len(v1)
                            var v24 bool = (!true)
                            _ = v23
                            _ = v24
                        } else {
                            v25 := GenS4{F0: 130}
                            _ = v25
                        }
                        _ = v22
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v18
            }
        case "a":
            if (!(174 > 1)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v26 int16 = 1000
                if true {
                    v1 = append(v1, v5)
                } else {
                    v27 := GenS1{F0: v5, F1: 1}
                    _ = v27
                }
                _ = v26
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v28 := GenS1{F0: v5, F1: 32767}
                v29 := GenS1{F0: (("" + "hello") + "hello"), F1: 42}
                _ = v28
                _ = v29
            }
        case "foo bar":
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int = -412
                _ = fi30
                _ = v31
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    } else {
        for fi32 := 0; fi32 < 12; fi32++ {
            for fi33 := 0; fi33 < 3; fi33++ {
                if ((18446744073709551615 == 0) || false) {
                    switch v2 {
                    case "\n\t":
                        switch (fi33 / (((fi33 & v3)) | 1)) {
                        case 0:
                            var v34 uint = 428
                            _ = v34
                        case -1:
                            fmt.Println(v5, v2, fi33)
                        case -291:
                            fmt.Println(fi32)
                        }
                        if false {
                            var v35 uint = 502
                            v35 += (v35 ^ (v35 % (((^(v35))) | 1)))
                            _ = v35
                        } else {
                            if (-79 >= 45) {
                                if true {
                                    v4 = append(v4, 796)
                                }
                                var v36 uint = 415
                                var v37 int32 = -838
                                _ = v36
                                _ = v37
                            } else {
                                var v38 int = 32768
                                var v39 uint16 = 2
                                switch v5 {
                                case "":
                                    v40 := GenS1{F0: "foo bar", F1: (v39 << ((42) & 63))}
                                    _ = v40
                                case "hello":
                                    fmt.Println(v38, v3)
                                }
                                _ = v38
                                _ = v39
                            }
                        }
                    }
                } else {
                    v3 -= v3
                    switch v5 {
                    case "":
                        var v41 int16 = 276
                        _ = v41
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = len(v1)
                }
                var v42 int = v3
                fmt.Println(fi33, fi32, v2)
                _ = fi33
                _ = v42
            }
            if (!false) {
                v43 := GenS3{F0: 2147483647, F1: 2147483647}
                var v44 float64 = 0.0
                var v45 int = fi32
                _ = v43
                _ = v44
                _ = v45
            }
            v3 -= v3
            _ = fi32
        }
        if false {
            for fi46 := 0; fi46 < 9; fi46++ {
                if len(v1) > 0 {
                    v1[0] = v5
                }
                v2 = v2
                fmt.Println(fi46, fi46)
                _ = fi46
            }
        }
        for fi47 := 0; fi47 < 19; fi47++ {
            var v48 int = (fi47 ^ -9223372036854775808)
            v49 := GenS5{F0: -1, F1: v5, F2: 0}
            for fi50 := 0; fi50 < 3; fi50++ {
                var v51 uint64 = 42
                fmt.Println(v2, v2)
                _ = len(v1)
                _ = fi50
                _ = v51
            }
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    for fi52 := 0; fi52 < 10; fi52++ {
        switch -9223372036854775808 {
        case -9223372036854775808:
            v1 = append(v1, ((v2 + "\n\t") + (v2 + v2)))
        case 127:
            var v53 uint = 4294967295
            _ = v53
        case 0:
            _ = len(v4)
        }
        var v54 uint64 = 18446744073709551615
        _ = fi52
        _ = v54
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        v56 := GenS6{F0: -269}
        v3 = v3
        for fi57 := 0; fi57 < 9; fi57++ {
            var v58 bool = false
            var v59 uint32 = 644
            v60 := GenS7{F0: v3, F1: 82, F2: 79}
            _ = fi57
            _ = v58
            _ = v59
            _ = v60
        }
        _ = fi55
        _ = v56
    }
    var v61 uint16 = 291
    var v62 int32 = 1
    for fi63 := 0; fi63 < 12; fi63++ {
        fmt.Println(v61, v62, v2)
        v64 := GenS8{F0: (fi63 ^ (v3 - v3)), F1: v3, F2: 42}
        _ = fi63
        _ = v64
    }
    switch (v5 + ((v2 + v2) + (v5 + v5))) {
    case "\n\t":
        for fi65 := 0; fi65 < 1; fi65++ {
            v66 := GenS1{F0: v5, F1: (^(v61))}
            for fi67 := 0; fi67 < 11; fi67++ {
                v68 := GenS8{F0: v3, F1: (^(v3)), F2: 0}
                var v69 float32 = 3.14159
                switch fi67 {
                case -1:
                    _ = len(v4)
                case 128:
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                case 692:
                    if (false && ((!((v62 + v62) < (v62 - v62))) || (!false))) {
                        if (!true) {
                            _ = len(v4)
                        } else {
                            if ((4294967295 != 663) && true) {
                                if true {
                                    var v70 bool = true
                                    _ = v70
                                } else {
                                    var v71 string = v5
                                    switch (v5 + ("hello" + v2)) {
                                    case "":
                                        v2 = (v5 + v5)
                                    case "foo bar":
                                        v71 = (v2 + "foo bar")
                                    default:
                                        var v72 float32 = (-2.25 - 0.1)
                                        _ = v72
                                    }
                                    v73 := GenS1{F0: v5, F1: (v61 | 2)}
                                    _ = v71
                                    _ = v73
                                }
                            } else {
                                var v74 uint8 = 127
                                _ = v74
                            }
                        }
                    } else {
                        if true {
                            var v75 int32 = v62
                            if true {
                                v3 = -251
                                if ((!((((((((0.0 - 1e10) * float64(float32(1e10))) / float64(((v69 * (float32((1e10 / 1.5)) / (v69 - v69))) - -2.25))) / 1e10) + 1e10) * 3.14159) > (((-2.25 - (1.5 + 0.0)) + 0.1) * float64((v69 * v69)))) && true)) && (false && (!true))) {
                                    if true {
                                        v76 := GenS3{F0: v75, F1: v75}
                                        v77 := GenS2{F0: 188}
                                        v78 := GenS5{F0: 0, F1: "foo bar", F2: 765}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        var v79 uint32 = 653
                                        _ = v79
                                    case "hello":
                                        var v80 float64 = float64(float32(0.0))
                                        _ = v80
                                    case "a":
                                        v81 := GenS6{F0: -9223372036854775808}
                                        _ = v81
                                    }
                                    v2 = ("a" + ((((v2 + v5) + "") + "hello") + v5))
                                    switch v3 {
                                    case -9223372036854775808:
                                        v5 = ((v2 + ("a" + v5)) + "foo bar")
                                    case 2147483647:
                                        fmt.Println(v75, v75, fi65)
                                    default:
                                        fmt.Println(v2, v75, fi67)
                                        var v82 float32 = v69
                                        var v83 uint = 527
                                        _ = v82
                                        _ = v83
                                    }
                                }
                            } else {
                                v84 := GenS1{F0: v5, F1: 255}
                                _ = v84
                            }
                            _ = v75
                        }
                        v5 = v5
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    if ((42 > 1) || false) {
                        if (!(1 >= 1000)) {
                            v85 := GenS5{F0: 1, F1: "hello", F2: 0}
                            if false {
                                var v86 int8 = -128
                                if (!true) {
                                    if false {
                                        v86 = v86
                                        var v87 int8 = v86
                                        v61 -= (^(v61))
                                        v62 -= (v62 ^ (v62 * v62))
                                        _ = v87
                                    } else {
                                        v2 = v2
                                        v62 -= v62
                                        v88 := GenS6{F0: -9223372036854775808}
                                        _ = v88
                                    }
                                } else {
                                    fmt.Println(v61, v62)
                                }
                                v3 -= v3
                                v5 = v5
                                _ = v86
                            } else {
                                v1 = append(v1, v2)
                                switch v5 {
                                case "foo bar":
                                    v89 := GenS4{F0: 2}
                                    _ = v89
                                case "\n\t":
                                    v90 := GenS1{F0: v5, F1: 516}
                                    _ = v90
                                case "a":
                                    v91 := GenS3{F0: v62, F1: (v62 + (v62 | (-(v62))))}
                                    _ = v91
                                default:
                                    v92 := GenS9{F0: -732, F1: (v61 << ((4294967295) & 63))}
                                    _ = v92
                                }
                            }
                            _ = v85
                        } else {
                            var v93 int16 = 42
                            v3 = (v3 * v3)
                            switch (fi67 * fi65) {
                            case -1:
                                var v94 uint8 = 1
                                _ = v94
                            case 32767:
                                fmt.Println(fi65, v5, fi65)
                            default:
                                v62 += (v62 % (((v62 >> ((493) & 63))) | 1))
                            }
                            _ = v93
                        }
                    } else {
                        v61 = 42
                        if ((false || (107 >= 141)) && (!(2147483648 >= 2147483647))) {
                            var v95 int32 = (v62 + (v62 ^ v62))
                            _ = v95
                        }
                        var v96 uint8 = 3
                        _ = v96
                    }
                    v97 := GenS5{F0: 0, F1: v2, F2: 255}
                    switch v5 {
                    case "":
                        if (float64(v69) < ((float64(3.14159) / 1.5) + 0.0)) {
                            var v98 uint = 65535
                            switch v5 {
                            case "hello":
                                _ = len(v1)
                            }
                            v99 := GenS10{F0: 1, F1: 1, F2: 2147483647}
                            var v100 uint16 = 1
                            _ = v100
                            _ = v98
                            _ = v99
                        } else {
                            v5 = v2
                            var v101 uint64 = 18446744073709551615
                            v69 = v69
                            _ = v101
                        }
                    }
                    var v102 float64 = (((0.1 + (float64(0.1) / ((0.1 + 0.0) - 0.0))) * (((((0.1 + -2.25) * (-2.25 / 0.1)) * float64(v69)) + ((((0.0 / -2.25) * 0.1) - (3.14159 / (0.1 * (1e10 * ((1e10 * (0.1 / ((0.1 * (1.5 / (-2.25 * float64(v69)))) + 1.5))) - 3.14159))))) * -2.25)) + 1e10)) + 3.14159)
                    _ = v102
                    _ = v97
                }
                _ = fi67
                _ = v68
                _ = v69
            }
            v62 = v62
            _ = fi65
            _ = v66
        }
    default:
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    fmt.Println(v61, v62)
    var v103 bool = (0.0 < (-2.25 * 0.0))
    for fi104 := 0; fi104 < 13; fi104++ {
        var v105 string = ("a" + ((v2 + v5) + "a"))
        v1 = append(v1, (v105 + v5))
        var v106 uint32 = 1
        _ = fi104
        _ = v105
        _ = v106
    }
    if true {
        var v107 float32 = 0.1
        var v108 uint8 = 128
        if true {
            if v103 {
                for fi109 := 0; fi109 < 5; fi109++ {
                    for fi110 := 0; fi110 < 13; fi110++ {
                        fmt.Println(v108, v107)
                        _ = fi110
                    }
                    _ = fi109
                }
            }
            for fi111 := 0; fi111 < 13; fi111++ {
                if false {
                    v62 -= (v62 / ((v62) | 1))
                    var v112 int8 = -1
                    _ = v112
                } else {
                    var v113 bool = v103
                    v114 := GenS11{F0: 283, F1: -32768}
                    _ = v113
                    _ = v114
                }
                v103 = ((v3 << ((0) & 63)) <= (v3 - (^(fi111))))
                _ = fi111
            }
            for fi115 := 0; fi115 < 8; fi115++ {
                var v116 uint16 = (v61 << ((2) & 63))
                var v117 uint32 = 1
                fmt.Println(v3, v116, v62)
                _ = fi115
                _ = v116
                _ = v117
            }
        } else {
            fmt.Println(v62, v5, v108)
        }
        _ = v107
        _ = v108
    } else {
        var v118 int8 = -128
        v3 -= v3
        var v119 int8 = (v118 % ((v118) | 1))
        _ = v118
        _ = v119
    }
    for fi120 := 0; fi120 < 18; fi120++ {
        _ = len(v4)
        switch ("a" + (v5 + v2)) {
        case "hello":
            switch fi120 {
            case 782:
                v121 := GenS12{F0: 100, F1: -422, F2: (v103 || true)}
                _ = v121
            case 872:
                v3 = (fi120 >> ((42) & 63))
            }
        default:
            v62 = (-(v62))
        }
        if len(v1) > 0 {
            v1[0] = v5
        }
        _ = fi120
    }
    fmt.Println(v62, v103)
    switch (v2 + v2) {
    case "foo bar":
        switch -779 {
        case -2147483648:
            var v122 uint16 = v61
            _ = v122
        case 42:
            v2 = (v5 + ("hello" + "foo bar"))
        }
    case "":
        if ((v61 << ((754) & 63)) < v61) {
            v123 := GenS4{F0: 127}
            _ = v123
        }
    case "a":
        var v124 int8 = genFn0(9223372036854775807, -46)
        _ = v124
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch -9223372036854775808 {
    case 127:
        switch (v5 + (v5 + (v2 + "foo bar"))) {
        case "hello":
            v3 -= (v3 ^ v3)
        }
    case 9223372036854775807:
        fmt.Println(v2, v103)
    case 4294967295:
        v61 = v61
    }
    for fi125 := 0; fi125 < 8; fi125++ {
        if v103 {
            switch (v3 | (v3 ^ (v3 % (((-(fi125))) | 1)))) {
            case 9223372036854775807:
                for fi126 := 0; fi126 < 15; fi126++ {
                    var v127 string = v2
                    if false {
                        _ = len(v1)
                    } else {
                        var v128 int64 = 0
                        _ = v128
                    }
                    _ = fi126
                    _ = v127
                }
            default:
                v62 += v62
            }
            v129 := GenS8{F0: fi125, F1: (^(fi125)), F2: 0}
            _ = v129
        }
        v103 = v103
        v130 := GenS13{F0: (v61 - 0), F1: v61, F2: v62}
        _ = fi125
        _ = v130
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v103
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
}
