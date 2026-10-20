package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint }
type GenS10 struct { F0 uint8; F1 int16; F2 int8 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int; F1 int64 }
type GenS13 struct { F0 int16; F1 int; F2 uint }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int8; F1 int64; F2 uint64 }
type GenS18 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int32, p1 int8, p2 int) uint64 {
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            if ((false || (18446744073709551615 < 0)) && true) {
                if (!true) {
                    v8 := []int{p2}
                    v9 := GenS1{F0: 1, F1: false}
                    _ = v8
                    _ = v9
                }
                switch ("a" + ("a" + "a")) {
                case "a":
                    v10 := []int{fi6, fi6, (fi6 % ((p2) | 1))}
                    _ = v10
                }
                var v11 float64 = -2.25
                _ = v11
            } else {
                v12 := []int{(fi7 | fi6), fi6, (p2 - (p2 % ((2147483647) | 1))), fi7}
                var v13 string = "a"
                _ = v12
                _ = v13
            }
            var v14 int32 = -2147483648
            _ = fi7
            _ = v14
        }
        switch "\n\t" {
        case "":
            v15 := GenS2{F0: 42, F1: p0}
            _ = v15
        case "hello":
            p2 += (fi6 % (((p2 ^ -2147483648)) | 1))
        case "foo bar":
            var v16 uint32 = 42
            _ = v16
        }
        _ = fi6
    }
    p1 += 42
    v17 := []int{(p2 ^ 4294967295), (-(p2)), (p2 - p2), 127}
    v18 := GenS3{F0: 501}
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 int = -2147483648
    v2 := []int{}
    v3 := []int{(v1 << ((984) & 63)), v1, 32768, (v1 - v1)}
    var v4 int = v1
    v5 := make(map[uint32]uint32)
    v1 = v4
    var v19 string = ""
    fmt.Println(v4, v4)
    v20 := GenS4{F0: (792 >= -1)}
    v19 = v19
    switch (("foo bar" + v19) + v19) {
    case "hello":
        if (!false) {
            v4 = (v4 ^ v4)
            v19 = (("\n\t" + v19) + (v19 + v19))
        }
    }
    var v21 float32 = (1.5 * 0.1)
    v5[4294967295] = 101
    var v22 uint8 = 197
    var v23 int16 = 1000
    if (984 <= 317) {
        if (false && true) {
            switch v19 {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "a":
                var v24 float64 = float64(v21)
                _ = v24
            default:
                v22 += 0
            }
            v4 += (v4 % ((v4) | 1))
        } else {
            switch "foo bar" {
            case "a":
                if (232 <= (v22 | 6)) {
                    if (((-2.25 * 0.1) - 1e10) <= 1.5) {
                        v22 = (v22 / (((v22 % (((v22 * v22)) | 1))) | 1))
                        var v25 uint64 = genFn0(-198, -125, (v1 >> ((42) & 63)))
                        if (v21 == 1e10) {
                            v26 := GenS5{F0: 65535}
                            switch v19 {
                            case "foo bar":
                                fmt.Println(v19, v22, v21)
                            }
                            _ = v26
                        } else {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v1, v4)
                            case "\n\t":
                                for fi27 := 0; fi27 < 4; fi27++ {
                                    _ = len(v3)
                                    v28 := GenS6{F0: -1}
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            fmt.Println(v19, v21)
                        }
                        _ = v25
                    } else {
                        v21 += v21
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v29 uint32 = 42
                    _ = v29
                } else {
                    switch "foo bar" {
                    case "a":
                        v30 := GenS2{F0: -739, F1: -2147483648}
                        _ = v30
                    case "":
                        v1 -= (v1 | v4)
                    case "foo bar":
                        switch v19 {
                        case "a":
                            fmt.Println(v1, v21)
                        default:
                            v31 := GenS7{F0: (v23 & v23)}
                            _ = v31
                        }
                    }
                }
            case "":
                switch v1 {
                case 42:
                    _ = len(v5)
                case 32767:
                    if (!(!false)) {
                        if (!true) {
                            v4 = v1
                            v21 = (v21 + ((float32(0.1) - (v21 - v21)) - v21))
                            var v32 uint8 = v22
                            _ = v32
                        }
                        v23 = -573
                    }
                case -529:
                    var v33 uint32 = 553
                    _ = v33
                }
            case "hello":
                for fi34 := 0; fi34 < 2; fi34++ {
                    if (2147483647 > -207) {
                        if false {
                            v21 += float32(3.14159)
                        } else {
                            v35 := GenS8{F0: v23, F1: 365}
                            if false {
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    var v37 int64 = 0
                                    _ = fi36
                                    _ = v37
                                }
                                v22 += 255
                            }
                            if (891 <= v23) {
                                var v38 float64 = 3.14159
                                fmt.Println(fi34, v23)
                                delete(v5, 35)
                                _ = v38
                            } else {
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    if (956 != 1) {
                                        var v40 string = v19
                                        v19 = (v19 + ("" + (v19 + v40)))
                                        if true {
                                            v41 := GenS4{F0: true}
                                            fmt.Println(v4, v4)
                                            var v42 uint8 = v22
                                            _ = v41
                                            _ = v42
                                        }
                                        _ = len(v5)
                                        v43 := GenS7{F0: (v23 ^ v23)}
                                        _ = v40
                                        _ = v43
                                    } else {
                                        if true {
                                            var v44 uint16 = 129
                                            _ = v44
                                        } else {
                                            fmt.Println(v19, fi34)
                                            if true {
                                                switch "a" {
                                                case "hello":
                                                    var v45 int16 = 42
                                                    _ = v45
                                                case "\n\t":
                                                    v46 := GenS9{F0: (v23 / (((v23 - 1)) | 1)), F1: v22, F2: 42}
                                                    _ = v46
                                                case "":
                                                    if (581 < 2147483647) {
                                                        var v47 uint = 475
                                                        var v48 uint8 = 42
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                }
                                                if true {
                                                    v4 -= fi39
                                                }
                                                if true {
                                                    _ = len(v2)
                                                    v22 += (v22 - (v22 | v22))
                                                }
                                            }
                                            fmt.Println(v22, v21, v23)
                                        }
                                    }
                                    var v49 bool = false
                                    var v50 uint16 = 277
                                    _ = fi39
                                    _ = v49
                                    _ = v50
                                }
                                v1 -= 42
                                if (1.5 > v21) {
                                    switch fi34 {
                                    case -1:
                                        for fi51 := 0; fi51 < 1; fi51++ {
                                            delete(v5, 42)
                                            _ = len(v3)
                                            if (v23 > (v23 & v23)) {
                                                var v52 uint8 = v22
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if true {
                                                    v52 = 0
                                                    fmt.Println(v19, v1, fi34)
                                                }
                                                _ = v52
                                            }
                                            var v53 int = 32768
                                            if ((v22 | v22) == v22) {
                                                v22 = v22
                                            }
                                            if false {
                                                switch 127 {
                                                case -462:
                                                    fmt.Println(v1, v22, v19)
                                                case 127:
                                                    var v54 int16 = v23
                                                    _ = v54
                                                case 9223372036854775807:
                                                    v21 = v21
                                                default:
                                                    if (int32(v23) >= -2147483648) {
                                                        var v55 int16 = -452
                                                        v21 += 0.0
                                                        var v56 uint16 = 1
                                                        if (v55 <= (v55 % (((v23 / ((v55) | 1))) | 1))) {
                                                            var v57 int16 = 268
                                                            v19 = ""
                                                            v22 += 128
                                                            _ = v57
                                                        } else {
                                                            v23 = (v23 & (v55 << ((255) & 63)))
                                                        }
                                                        if false {
                                                            switch (v53 - 307) {
                                                            case 127:
                                                                v58 := GenS10{F0: 127, F1: 42, F2: -128}
                                                                v59 := GenS9{F0: -32768, F1: (v22 ^ 128), F2: 4294967295}
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                            v19 = v19
                                                            v22 -= v22
                                                            v55 += (v55 >> ((1) & 63))
                                                        }
                                                        switch 1 {
                                                        case 32768:
                                                            switch (v19 + v19) {
                                                            case "\n\t":
                                                                switch (((v19 + (v19 + v19)) + "foo bar") + v19) {
                                                                case "\n\t":
                                                                    var v60 float64 = 3.14159
                                                                    _ = v60
                                                                }
                                                            case "hello":
                                                                v1 = v4
                                                            case "":
                                                                v61 := v5[65535]
                                                                _ = v61
                                                            default:
                                                                var v62 uint32 = 301
                                                                _ = v62
                                                            }
                                                        case 128:
                                                            var v63 bool = (false || false)
                                                            _ = v63
                                                        case 490:
                                                            fmt.Println(fi51, v4, v22)
                                                        }
                                                        _ = len(v3)
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(v22, v23, v23)
                                                        v4 = (v53 % ((v1) | 1))
                                                        var v64 uint32 = 0
                                                        _ = v64
                                                    }
                                                }
                                                fmt.Println(v1)
                                                fmt.Println(v22)
                                            } else {
                                                fmt.Println(v23, v22, v19)
                                                if true {
                                                    v21 -= 3.14159
                                                    fmt.Println(v22, v22)
                                                    var v65 int16 = (v23 << ((4294967295) & 63))
                                                    _ = v65
                                                } else {
                                                    if (true || (("" >= "a") && true)) {
                                                        switch v19 {
                                                        case "":
                                                            var v66 float64 = ((1e10 + 1e10) / -2.25)
                                                            _ = v66
                                                        }
                                                        var v67 uint = 42
                                                        _ = v67
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = v1
                                                        }
                                                        fmt.Println(fi34)
                                                    }
                                                    fmt.Println(v1, v53)
                                                    switch v19 {
                                                    case "hello":
                                                        if true {
                                                            v22 = 128
                                                            v1 -= 42
                                                        }
                                                    case "\n\t":
                                                        v53 = v53
                                                    case "a":
                                                        v21 = (float32(1e10) / (v21 * v21))
                                                    default:
                                                        if (v21 >= 0.1) {
                                                            fmt.Println(fi51, v19)
                                                            fmt.Println(v53, v21)
                                                            _ = len(v5)
                                                        } else {
                                                            v1 = (v1 >> ((0) & 63))
                                                            switch ("\n\t" + (((("foo bar" + ((v19 + (v19 + "")) + "\n\t")) + "a") + v19) + v19)) {
                                                            case "foo bar":
                                                                if (1 >= 0) {
                                                                    v19 = ((("foo bar" + v19) + v19) + "")
                                                                    if (873 == 164) {
                                                                        fmt.Println(v4)
                                                                        if true {
                                                                            if (!(!(!true))) {
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            v4 = fi34
                                                                            fmt.Println(v23, v53)
                                                                        } else {
                                                                            if (435 < 42) {
                                                                                _ = len(v2)
                                                                            } else {
                                                                                if (v22 > v22) {
                                                                                    var v68 int = (v1 + 128)
                                                                                    _ = v68
                                                                                } else {
                                                                                    v4 -= (v1 << ((18446744073709551615) & 63))
                                                                                    v69 := GenS9{F0: (v23 & (v23 % (((v23 >> ((18446744073709551615) & 63))) | 1))), F1: (v22 >> ((65535) & 63)), F2: 4294967295}
                                                                                    v70 := GenS8{F0: v23, F1: 255}
                                                                                    if ((v19 + v19) != v19) {
                                                                                        if true {
                                                                                            v22 = (v22 + (v22 >> ((42) & 63)))
                                                                                            var v71 int = (v4 - (fi51 | (^(v53))))
                                                                                            v23 += v23
                                                                                            _ = v71
                                                                                        } else {
                                                                                            v19 = v19
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            v72 := GenS8{F0: v23, F1: 42}
                                                                                            var v73 int8 = 100
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        }
                                                                                        switch 127 {
                                                                                        case 2147483647:
                                                                                            switch v19 {
                                                                                            case "foo bar":
                                                                                                var v74 float64 = 3.14159
                                                                                                _ = v74
                                                                                            default:
                                                                                                v4 += v1
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v4 -= 4294967295
                                                                                    }
                                                                                    _ = v69
                                                                                    _ = v70
                                                                                }
                                                                                if false {
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = (fi51 - -9223372036854775808)
                                                                                    }
                                                                                } else {
                                                                                    var v75 int8 = -128
                                                                                    v76 := GenS10{F0: v22, F1: 534, F2: (-(v75))}
                                                                                    var v77 string = ("foo bar" + v19)
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                var v78 int32 = 371
                                                                                v79 := GenS6{F0: (v23 ^ (v23 - (v23 & v23)))}
                                                                                _ = v78
                                                                                _ = v79
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (9223372036854775807 >= 9223372036854775807) {
                                                                            switch ("foo bar" + "foo bar") {
                                                                            case "\n\t":
                                                                                v80 := GenS4{F0: (int64(v23) <= -9223372036854775808)}
                                                                                _ = v80
                                                                            }
                                                                        } else {
                                                                            v81 := GenS11{F0: 0}
                                                                            var v82 int16 = (^(v23))
                                                                            v19 = ("hello" + "\n\t")
                                                                            _ = v81
                                                                            _ = v82
                                                                        }
                                                                        var v83 uint32 = 42
                                                                        var v84 int = (-(v53))
                                                                        _ = v83
                                                                        _ = v84
                                                                    }
                                                                } else {
                                                                    v19 = v19
                                                                }
                                                            case "\n\t":
                                                                var v85 uint32 = 0
                                                                _ = v85
                                                            case "":
                                                                var v86 string = (((v19 + v19) + v19) + (("a" + "hello") + "foo bar"))
                                                                _ = v86
                                                            default:
                                                                if true {
                                                                    v22 = (v22 % ((2) | 1))
                                                                    v4 += fi51
                                                                }
                                                            }
                                                            v87 := GenS5{F0: 8}
                                                            _ = v87
                                                        }
                                                    }
                                                    fmt.Println(fi51, fi51)
                                                }
                                                var v88 int64 = -49
                                                if false {
                                                    switch int(v22) {
                                                    case 127:
                                                        fmt.Println(v23, fi51, v21)
                                                    case -548:
                                                        v89 := GenS2{F0: (v88 << ((2) & 63)), F1: int32(v88)}
                                                        _ = v89
                                                    default:
                                                        if (v22 >= 117) {
                                                            v19 = (((v19 + "a") + v19) + "foo bar")
                                                        }
                                                    }
                                                    delete(v5, 1)
                                                    switch 0 {
                                                    case -929:
                                                        var v90 float32 = v21
                                                        _ = v90
                                                    case 127:
                                                        v22 = 255
                                                    case 4294967295:
                                                        if false {
                                                            v1 = v4
                                                            v23 -= (v23 / (((v23 ^ v23)) | 1))
                                                            if (1.5 > (0.0 / 0.1)) {
                                                                var v91 float32 = -2.25
                                                                _ = v91
                                                            } else {
                                                                switch (fi51 & int(v23)) {
                                                                case 350:
                                                                    v23 += (v23 + v23)
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                                v92 := v5[548]
                                                _ = v88
                                                _ = v92
                                            }
                                            v53 -= (fi34 | v53)
                                            var v93 string = ""
                                            _ = fi51
                                            _ = v53
                                            _ = v93
                                        }
                                    }
                                    v2 = append(v2, -2147483648)
                                    if false {
                                        v1 = v1
                                        if false {
                                            if ((((1e10 - v21) / v21) <= (v21 - (v21 / -2.25))) && true) {
                                                delete(v5, 467)
                                            }
                                            v22 += v22
                                        }
                                        fmt.Println(v19, fi34)
                                    }
                                }
                            }
                            _ = v35
                        }
                        v94 := GenS2{F0: 597, F1: -1}
                        for fi95 := 0; fi95 < 9; fi95++ {
                            if (v22 >= 255) {
                                var v96 int = -2147483648
                                v22 -= (v22 & 0)
                                v97 := GenS12{F0: (v1 << ((42) & 63)), F1: -1}
                                _ = v96
                                _ = v97
                            }
                            _ = fi95
                        }
                        _ = v94
                    } else {
                        v23 = v23
                        for fi98 := 0; fi98 < 7; fi98++ {
                            switch v19 {
                            case "\n\t":
                                if (102 <= 1) {
                                    var v99 int32 = -2147483648
                                    v21 -= v21
                                    _ = v99
                                }
                            case "foo bar":
                                v100 := GenS12{F0: v1, F1: 1}
                                _ = v100
                            case "a":
                                v101 := GenS12{F0: (fi34 / (((v1 / (((v4 | (fi98 + (fi98 + (v4 * (fi98 & (fi98 - 32768))))))) | 1))) | 1)), F1: 1}
                                _ = v101
                            default:
                                var v102 float32 = (float32(1.5) / v21)
                                _ = v102
                            }
                            v103 := GenS13{F0: (v23 - v23), F1: v4, F2: 42}
                            _ = fi98
                            _ = v103
                        }
                        v104 := GenS14{F0: v22}
                        _ = v104
                    }
                    var v105 uint = 118
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi34
                    _ = v105
                }
            default:
                if len(v3) > 0 {
                    v3[0] = v1
                }
            }
        }
        var v106 int = (-(v1))
        _ = v106
    } else {
        delete(v5, 4294967295)
        fmt.Println(v23, v23, v22)
        v107 := GenS13{F0: -1, F1: (v1 % ((v1) | 1)), F2: 2}
        _ = v107
    }
    v3 = append(v3, (v1 << ((0) & 63)))
    for fi108 := 0; fi108 < 3; fi108++ {
        v5[2] = 0
        for fi109 := 0; fi109 < 4; fi109++ {
            v110 := GenS15{F0: 127}
            _ = fi109
            _ = v110
        }
        _ = fi108
    }
    for fi111 := 0; fi111 < 13; fi111++ {
        switch (v1 * 887) {
        case 4294967295:
            if false {
                fmt.Println(v19, fi111)
            } else {
                fmt.Println(v22, v21)
                v112 := GenS7{F0: -32768}
                _ = v112
            }
        }
        for fi113 := 0; fi113 < 3; fi113++ {
            v23 -= 42
            _ = fi113
        }
        v1 = (-(fi111))
        _ = fi111
    }
    delete(v5, 2147483648)
    var v114 int16 = 0
    for fi115 := 0; fi115 < 8; fi115++ {
        switch "foo bar" {
        case "\n\t":
            v116 := GenS16{F0: 0}
            _ = v116
        case "hello":
            v114 = -1
        case "":
            v4 += v1
        default:
            if false {
                for fi117 := 0; fi117 < 9; fi117++ {
                    if (false && false) {
                        if ((2147483648 < 2147483648) && (!(!((v19 + (("" + ("hello" + "")) + "\n\t")) >= v19)))) {
                            v118 := GenS17{F0: 127, F1: -9223372036854775808, F2: 42}
                            var v119 string = "\n\t"
                            var v120 bool = ((true || (1 == 65535)) || false)
                            v4 = (-(fi117))
                            var v121 string = "\n\t"
                            fmt.Println(v23)
                            _ = v118
                            _ = v119
                            _ = v120
                            _ = v121
                        }
                    }
                    var v122 int32 = -1
                    _ = fi117
                    _ = v122
                }
                v5[2147483647] = 507
            } else {
                switch v1 {
                case -715:
                    v22 -= v22
                default:
                    switch "a" {
                    case "a":
                        var v123 float32 = ((1.5 / (v21 + ((((3.14159 - v21) - -2.25) * v21) / 1e10))) - ((v21 - (0.1 * v21)) - (float32(3.14159) / ((v21 - v21) - float32(float64(v21))))))
                        _ = v123
                    }
                }
                v124 := GenS7{F0: (v114 | (v23 | v114))}
                _ = v124
            }
        }
        _ = fi115
    }
    if false {
        if true {
            var v125 bool = (1 != -1)
            for fi126 := 0; fi126 < 4; fi126++ {
                if v125 {
                    v1 = v4
                }
                if len(v3) > 0 {
                    v3[0] = 128
                }
                _ = fi126
            }
            _ = v125
        }
        if false {
            if ((true || (v114 <= v23)) && (927 != 255)) {
                switch 112 {
                case -185:
                    var v127 uint8 = v22
                    _ = v127
                }
                for fi128 := 0; fi128 < 16; fi128++ {
                    v4 = fi128
                    _ = fi128
                }
                var v129 int64 = 1
                _ = v129
            } else {
                var v130 string = v19
                v22 = v22
                _ = v130
            }
            v131 := GenS18{F0: 4294967295, F1: 2147483647}
            _ = v131
        }
        v132 := GenS12{F0: v4, F1: -9223372036854775808}
        _ = v132
    }
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v114
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
