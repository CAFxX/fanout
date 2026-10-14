package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64; F2 bool }
type GenS5 struct { F0 int; F1 int16; F2 string }
type GenS6 struct { F0 int32; F1 string; F2 int32 }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS10 struct { F0 int8; F1 int8 }
type GenS11 struct { F0 uint8; F1 uint8 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint; F1 uint32; F2 int32 }
type GenS14 struct { F0 uint64; F1 string }
type GenS15 struct { F0 int8 }
func genFn0(p0 uint32, p1 string, p2 uint16) uint64 {
    if false {
        var v4 string = "foo bar"
        _ = v4
    }
    p0 = (p0 ^ (p0 >> ((1) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return 36
}
func genFn1(p0 uint8) bool {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    if false {
        if (false || (63 <= 460)) {
            if true {
                v7 := []int{-9223372036854775808, 128}
                fmt.Println(p0)
                _ = v7
            } else {
                v8 := []int{}
                _ = v8
            }
            p0 = 185
        }
    }
    var v9 float64 = float64(3.14159)
    _ = p0
    _ = v9
    return true
}

func main() {
    v1 := []int{42}
    var v2 uint8 = 42
    v3 := make(map[int]string)
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v2, v2)
        v11 := GenS1{F0: 4294967295, F1: fi10, F2: (fi10 >> ((623) & 63))}
        _ = fi10
        _ = v11
    }
    switch "" {
    case "foo bar":
        v2 = v2
    case "":
        switch 531 {
        case 4294967295:
            fmt.Println(v2)
        case 0:
            var v12 float64 = 0.1
            _ = v12
        }
    default:
        if (!(!(4294967295 > 128))) {
            v2 -= (v2 << ((2) & 63))
        } else {
            v13 := GenS2{F0: 1}
            _ = v13
        }
    }
    var v14 uint16 = 762
    v2 += (^(v2))
    fmt.Println(v2, v14)
    v15 := v3[335]
    if false {
        _ = len(v3)
    }
    fmt.Println(v2)
    var v16 uint8 = (v2 + uint8(v14))
    if (127 < 0) {
        switch 277 {
        case 286:
            var v17 float32 = 3.14159
            _ = v17
        case 42:
            for fi18 := 0; fi18 < 13; fi18++ {
                v16 -= (v16 & 0)
                _ = fi18
            }
        case 1:
            var v19 uint = 4294967295
            _ = v19
        }
        switch "" {
        case "hello":
            v14 -= 2
        }
        var v20 uint = 255
        _ = v20
    } else {
        for fi21 := 0; fi21 < 1; fi21++ {
            v15 = (("a" + "\n\t") + "a")
            fmt.Println(v16, v2, v16)
            _ = fi21
        }
    }
    v22 := GenS3{F0: uint8(v14)}
    var v23 int = 42
    var v24 string = v15
    if (true || (128 > (v16 * v2))) {
        var v25 uint64 = 0
        if (true && ((((false && (int64(v23) > int64(v25))) && (!true)) || true) || (42 != -636))) {
            var v26 float32 = 3.14159
            _ = v26
        } else {
            switch 42 {
            case -827:
                switch (v23 * (v23 - v23)) {
                case 1:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case 32767:
                    for fi27 := 0; fi27 < 1; fi27++ {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            _ = len(v3)
                            _ = fi28
                        }
                        if false {
                            v15 = ("foo bar" + ("" + "a"))
                            for fi29 := 0; fi29 < 18; fi29++ {
                                switch "" {
                                case "hello":
                                    var v30 uint32 = 4294967295
                                    _ = v30
                                case "foo bar":
                                    var v31 uint8 = (v2 - (v2 << ((169) & 63)))
                                    _ = v31
                                default:
                                    v32 := GenS4{F0: v2, F1: v25, F2: true}
                                    _ = v32
                                }
                                _ = fi29
                            }
                            v33 := GenS1{F0: (v25 + v25), F1: 167, F2: (v23 - v23)}
                            _ = v33
                        } else {
                            if (uint(v14) == 314) {
                                fmt.Println(v16, v16)
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    fmt.Println(v23)
                                    var v35 uint64 = v25
                                    delete(v3, v23)
                                    if true {
                                        if ((v14 ^ (^(v14))) < v14) {
                                            if (-907 < -1) {
                                                v36 := GenS5{F0: -9223372036854775808, F1: -905, F2: v15}
                                                _ = len(v3)
                                                var v37 uint64 = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            fmt.Println(v14)
                                        }
                                    }
                                    _ = fi34
                                    _ = v35
                                }
                                if false {
                                    for fi38 := 0; fi38 < 10; fi38++ {
                                        v39 := GenS2{F0: (fi27 % (((v23 ^ 0)) | 1))}
                                        v40 := GenS6{F0: -2147483648, F1: (v15 + v24), F2: 441}
                                        v41 := v3[fi27]
                                        v2 = (v2 + (^(v16)))
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    }
                                    fmt.Println(v15)
                                    delete(v3, 2147483647)
                                }
                            }
                            var v42 int32 = 0
                            _ = v42
                        }
                        v43 := GenS7{F0: 324, F1: 1}
                        _ = fi27
                        _ = v43
                    }
                case 128:
                    var v44 float32 = float32(0.0)
                    _ = v44
                default:
                    if false {
                        if false {
                            switch ((v15 + "foo bar") + "") {
                            case "hello":
                                for fi45 := 0; fi45 < 12; fi45++ {
                                    var v46 uint32 = 4294967295
                                    for fi47 := 0; fi47 < 9; fi47++ {
                                        v1 = append(v1, (v23 & fi45))
                                        v23 = fi45
                                        _ = fi47
                                    }
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            var v48 int = (^(v23))
                            v49 := GenS4{F0: uint8(v14), F1: genFn0(2147483647, (v15 + ((v15 + ("a" + v15)) + ("foo bar" + "a"))), 0), F2: true}
                            _ = v48
                            _ = v49
                        } else {
                            if (v2 < v16) {
                                v15 = v24
                                v15 = (v15 + v15)
                            }
                        }
                        v23 = (v23 | v23)
                        if true {
                            v16 = (v2 - (v2 * v2))
                            if true {
                                for fi50 := 0; fi50 < 19; fi50++ {
                                    if (v25 < v25) {
                                        for fi51 := 0; fi51 < 19; fi51++ {
                                            fmt.Println(v25, v16, v24)
                                            if len(v1) > 0 {
                                                v1[0] = (v23 >> ((876) & 63))
                                            }
                                            var v52 bool = true
                                            _ = fi51
                                            _ = v52
                                        }
                                    } else {
                                        for fi53 := 0; fi53 < 10; fi53++ {
                                            v23 = (fi53 / ((fi53) | 1))
                                            v54 := GenS3{F0: (v2 | 127)}
                                            _ = fi53
                                            _ = v54
                                        }
                                        if (65535 >= 18446744073709551615) {
                                            if false {
                                                fmt.Println(v14, v15, v15)
                                            }
                                        } else {
                                            var v55 float32 = (1e10 * float32((1e10 - -2.25)))
                                            _ = v55
                                        }
                                    }
                                    _ = len(v3)
                                    _ = fi50
                                }
                                fmt.Println(v15, v2)
                            } else {
                                var v56 uint16 = (^(v14))
                                if true {
                                    if false {
                                        v57 := GenS3{F0: 128}
                                        v58 := GenS8{F0: int32(v25)}
                                        _ = v57
                                        _ = v58
                                    }
                                } else {
                                    if false {
                                        for fi59 := 0; fi59 < 18; fi59++ {
                                            v15 = ("\n\t" + ("" + v24))
                                            v60 := GenS5{F0: fi59, F1: 1, F2: v24}
                                            _ = fi59
                                            _ = v60
                                        }
                                    } else {
                                        for fi61 := 0; fi61 < 9; fi61++ {
                                            var v62 float32 = 0.1
                                            delete(v3, (v23 * v23))
                                            v63 := v3[(v23 + 128)]
                                            _ = fi61
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                }
                                v24 = (v15 + (v15 + v15))
                                _ = v56
                            }
                            if true {
                                v25 += 0
                            }
                        }
                    } else {
                        v64 := GenS9{F0: 646, F1: 755, F2: v25}
                        switch v24 {
                        case "\n\t":
                            v65 := GenS10{F0: 0, F1: 1}
                            _ = v65
                        case "foo bar":
                            v14 = v14
                        case "a":
                            var v66 uint32 = 1
                            _ = v66
                        default:
                            v15 = v24
                        }
                        var v67 uint = 255
                        _ = v64
                        _ = v67
                    }
                }
            case 1:
                v68 := GenS11{F0: (v2 << ((4294967295) & 63)), F1: v2}
                _ = v68
            case 4294967295:
                switch v24 {
                case "":
                    fmt.Println(v25, v2)
                case "a":
                    v16 = v16
                case "foo bar":
                    if false {
                        var v69 bool = ((!(1 != -2147483648)) || true)
                        for fi70 := 0; fi70 < 3; fi70++ {
                            var v71 int32 = -476
                            _ = fi70
                            _ = v71
                        }
                        v72 := GenS4{F0: v16, F1: v25, F2: v69}
                        _ = v69
                        _ = v72
                    } else {
                        switch v24 {
                        case "a":
                            switch 343 {
                            case -648:
                                switch "a" {
                                case "\n\t":
                                    switch (v23 ^ 830) {
                                    case 385:
                                        for fi73 := 0; fi73 < 13; fi73++ {
                                            v74 := GenS7{F0: 433, F1: -338}
                                            fmt.Println(v15, v23, v25)
                                            _ = fi73
                                            _ = v74
                                        }
                                    case 32767:
                                        var v75 int8 = int8(v16)
                                        _ = v75
                                    }
                                case "foo bar":
                                    for fi76 := 0; fi76 < 19; fi76++ {
                                        var v77 int32 = 0
                                        var v78 int32 = (v77 % ((v77) | 1))
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            default:
                                v23 -= v23
                            }
                        default:
                            v14 = v14
                        }
                    }
                default:
                    v23 = (v23 - 343)
                }
            default:
                var v79 string = "foo bar"
                _ = v79
            }
        }
        for fi80 := 0; fi80 < 15; fi80++ {
            fmt.Println(v2)
            _ = fi80
        }
        _ = v25
    } else {
        for fi81 := 0; fi81 < 3; fi81++ {
            if (712 < (v14 >> ((18446744073709551615) & 63))) {
                fmt.Println(v24, v24, v14)
                if (!false) {
                    var v82 uint = 1
                    _ = v82
                }
                switch (fi81 % (((fi81 + 42)) | 1)) {
                case -9223372036854775808:
                    switch v15 {
                    case "foo bar":
                        for fi83 := 0; fi83 < 2; fi83++ {
                            if false {
                                var v84 int8 = -1
                                v85 := GenS3{F0: (v2 & (v2 + (v2 >> ((10) & 63))))}
                                var v86 int32 = -556
                                _ = v84
                                _ = v85
                                _ = v86
                            } else {
                                var v87 int64 = 797
                                var v88 float64 = 1e10
                                _ = v87
                                _ = v88
                            }
                            var v89 uint16 = v14
                            fmt.Println(v16, v2)
                            _ = fi83
                            _ = v89
                        }
                    case "\n\t":
                        for fi90 := 0; fi90 < 9; fi90++ {
                            var v91 int64 = 1
                            _ = fi90
                            _ = v91
                        }
                    }
                case -148:
                    for fi92 := 0; fi92 < 7; fi92++ {
                        if (!false) {
                            if false {
                                switch ("\n\t" + v15) {
                                case "\n\t":
                                    _ = len(v3)
                                    if false {
                                        v23 = 559
                                    }
                                    fmt.Println(v2)
                                case "hello":
                                    v2 += 1
                                default:
                                    if false {
                                        v93 := GenS5{F0: (v23 * fi81), F1: 42, F2: ("" + (v15 + (v15 + v15)))}
                                        var v94 uint8 = v2
                                        v95 := GenS12{F0: ("foo bar" + v24)}
                                        _ = v93
                                        _ = v94
                                        _ = v95
                                    }
                                }
                                v23 -= (fi81 >> ((796) & 63))
                            }
                        } else {
                            v16 = (v2 & (v2 - v16))
                            var v96 int = 540
                            _ = v96
                        }
                        if false {
                            v24 = v24
                            if false {
                                var v97 int8 = 43
                                _ = v97
                            }
                            v14 -= v14
                            if (100 == 119) {
                                switch "foo bar" {
                                case "":
                                    switch v24 {
                                    case "hello":
                                        switch ((v24 + "a") + (v24 + v24)) {
                                        case "\n\t":
                                            v98 := GenS10{F0: 40, F1: -41}
                                            _ = v98
                                        case "hello":
                                            fmt.Println(v2, fi81)
                                        case "foo bar":
                                            v2 = v2
                                        }
                                    default:
                                        var v99 float64 = 1e10
                                        var v100 int16 = -110
                                        _ = v100
                                        _ = v99
                                    }
                                }
                                var v101 uint16 = v14
                                fmt.Println(v16)
                                _ = v101
                            } else {
                                switch v24 {
                                case "a":
                                    v3[(^(fi92))] = v15
                                case "hello":
                                    switch v15 {
                                    case "foo bar":
                                        fmt.Println(v2)
                                    case "":
                                        if len(v1) > 0 {
                                            v1[0] = (fi92 << ((42) & 63))
                                        }
                                    default:
                                        if true {
                                            v102 := GenS13{F0: 4294967295, F1: 482, F2: -64}
                                            _ = len(v3)
                                            v23 -= -782
                                            _ = v102
                                        } else {
                                            if true {
                                                switch (("" + v24) + v15) {
                                                case "\n\t":
                                                    var v103 int64 = 42
                                                    _ = v103
                                                case "foo bar":
                                                    if true {
                                                        delete(v3, -2147483648)
                                                    } else {
                                                        if true {
                                                            _ = len(v3)
                                                            fmt.Println(v2)
                                                        } else {
                                                            v23 -= -2147483648
                                                        }
                                                        fmt.Println(fi81)
                                                    }
                                                    if (4294967295 <= 4294967295) {
                                                        v104 := GenS4{F0: (v16 << ((0) & 63)), F1: 165, F2: ((-92 < -128) && false)}
                                                        var v105 int16 = 510
                                                        fmt.Println(v24)
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                default:
                                                    switch fi92 {
                                                    case 127:
                                                        var v106 int8 = -128
                                                        _ = v106
                                                    default:
                                                        v107 := GenS5{F0: 127, F1: 32767, F2: v24}
                                                        _ = v107
                                                    }
                                                }
                                            }
                                            v108 := v3[v23]
                                            _ = v108
                                        }
                                    }
                                case "":
                                    var v109 uint = 65535
                                    _ = v109
                                default:
                                    fmt.Println(fi92)
                                }
                                var v110 float32 = 1.5
                                if (1e10 < v110) {
                                    switch fi92 {
                                    case 32768:
                                        var v111 int16 = 42
                                        _ = v111
                                    }
                                    var v112 int16 = 42
                                    var v113 int8 = -128
                                    _ = v112
                                    _ = v113
                                } else {
                                    v114 := GenS14{F0: 2, F1: v15}
                                    if (v14 <= (v14 + v14)) {
                                        v23 -= (fi81 / (((v23 / (((fi81 >> ((uint(v2)) & 63))) | 1))) | 1))
                                    } else {
                                        v24 = "a"
                                        v115 := v3[(fi92 | fi81)]
                                        var v116 uint = 0
                                        if (!((!(2 != v116)) || (574 < -32768))) {
                                            v2 = v2
                                        } else {
                                            if len(v1) > 0 {
                                                v1[0] = v23
                                            }
                                            var v117 uint64 = genFn0(2147483647, (v24 + (v115 + v24)), (v14 / (((v14 ^ v14)) | 1)))
                                            var v118 int8 = -128
                                            _ = v117
                                            _ = v118
                                        }
                                        v119 := GenS2{F0: fi92}
                                        _ = v115
                                        _ = v116
                                        _ = v119
                                    }
                                    _ = v114
                                }
                                switch "a" {
                                case "\n\t":
                                    v110 += 1e10
                                case "hello":
                                    switch v15 {
                                    case "hello":
                                        fmt.Println(v16, fi92)
                                    case "foo bar":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                }
                                v15 = ("\n\t" + v15)
                                _ = v110
                            }
                        }
                        var v120 bool = true
                        _ = fi92
                        _ = v120
                    }
                case 704:
                    v121 := v3[(fi81 | (v23 << ((42) & 63)))]
                    _ = v121
                }
            } else {
                v24 = ""
                for fi122 := 0; fi122 < 6; fi122++ {
                    var v123 int64 = -1
                    _ = fi122
                    _ = v123
                }
                if true {
                    var v124 float32 = (0.0 * (3.14159 + -2.25))
                    fmt.Println(v14, v23, fi81)
                    v125 := GenS10{F0: 0, F1: 127}
                    _ = v124
                    _ = v125
                } else {
                    var v126 int = fi81
                    delete(v3, (^(v126)))
                    _ = v126
                }
            }
            v127 := GenS15{F0: -82}
            _ = fi81
            _ = v127
        }
    }
    v16 += v2
    if (true || (true || (440 >= 2))) {
        switch "" {
        case "\n\t":
            v24 = v24
        default:
            fmt.Println(v15, v24, v2)
        }
    } else {
        v23 = v23
        var v128 float32 = 3.14159
        _ = v128
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
