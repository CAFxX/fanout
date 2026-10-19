package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int16; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int8; F1 int8; F2 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 int32 }

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{764, 4294967295, 1}
    var v3 string = ("" + "a")
    v3 = v3
    var v4 uint64 = 0
    if len(v2) > 0 {
        v2[0] = 65535
    }
    for fi5 := 0; fi5 < 7; fi5++ {
        switch (fi5 + (^(fi5))) {
        case 9223372036854775807:
            v2 = append(v2, 2147483648)
        case 0:
            v6 := GenS1{F0: v3, F1: 1, F2: 1}
            _ = v6
        default:
            _ = len(v2)
        }
        v4 = 287
        _ = fi5
    }
    v1 = v1
    v4 = v4
    v4 += 4294967295
    var v7 int = -9223372036854775808
    var v8 bool = true
    if (1 < 343) {
        switch ((v3 + v3) + (v3 + "hello")) {
        case "foo bar":
            switch (v7 | 1) {
            case 32768:
                var v9 float64 = 1.5
                _ = v9
            }
        case "":
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
        }
        if v8 {
            if true {
                v8 = (!(v1 < v1))
                fmt.Println(v7, v7)
            } else {
                for fi10 := 0; fi10 < 12; fi10++ {
                    if (!(v1 == 3.14159)) {
                        fmt.Println(v1, v8)
                        v1 = 0.1
                        v11 := GenS1{F0: v3, F1: -982, F2: -1}
                        _ = v11
                    }
                    if len(v2) > 0 {
                        v2[0] = 2147483647
                    }
                    if false {
                        if (!v8) {
                            v12 := GenS2{F0: 1, F1: 65535}
                            switch v3 {
                            case "a":
                                for fi13 := 0; fi13 < 10; fi13++ {
                                    var v14 int64 = 42
                                    var v15 int = fi10
                                    var v16 uint64 = 291
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                            case "\n\t":
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi17
                                }
                            case "hello":
                                v8 = (!v8)
                            default:
                                if v8 {
                                    v18 := GenS1{F0: v3, F1: -9223372036854775808, F2: 100}
                                    _ = v18
                                } else {
                                    if ((^(v4)) >= 127) {
                                        var v19 uint = 4294967295
                                        switch (v3 + v3) {
                                        case "foo bar":
                                            if (v3 < ("\n\t" + (v3 + v3))) {
                                                _ = len(v2)
                                            } else {
                                                var v20 uint64 = 707
                                                _ = v20
                                            }
                                        case "\n\t":
                                            var v21 int16 = 32767
                                            _ = v21
                                        default:
                                            if v8 {
                                                v1 = ((0.0 / 1.0) - float32((1.5 - 0.1)))
                                                fmt.Println(v3, v7)
                                                var v22 int16 = 1000
                                                _ = v22
                                            } else {
                                                switch ((v3 + ("\n\t" + v3)) + ("\n\t" + "\n\t")) {
                                                case "hello":
                                                    var v23 int32 = 368
                                                    _ = v23
                                                case "a":
                                                    if len(v2) > 0 {
                                                        v2[0] = 789
                                                    }
                                                case "\n\t":
                                                    var v24 int16 = 1
                                                    _ = v24
                                                default:
                                                    if v8 {
                                                        v19 = (v19 | 625)
                                                        var v25 uint64 = v4
                                                        _ = v25
                                                    } else {
                                                        var v26 int = v7
                                                        if v8 {
                                                            for fi27 := 0; fi27 < 16; fi27++ {
                                                                _ = len(v2)
                                                                switch "a" {
                                                                case "\n\t":
                                                                    _ = len(v2)
                                                                case "a":
                                                                    v28 := GenS3{F0: 32767, F1: -1, F2: 65535}
                                                                    _ = v28
                                                                default:
                                                                    var v29 uint16 = 1
                                                                    _ = v29
                                                                }
                                                                v1 += v1
                                                                _ = fi27
                                                            }
                                                            switch "a" {
                                                            case "\n\t":
                                                                if ((0.1 * (1.5 * 3.14159)) <= (0.1 * 0.1)) {
                                                                    v19 += v19
                                                                    if len(v2) > 0 {
                                                                        v2[0] = 65535
                                                                    }
                                                                    for fi30 := 0; fi30 < 9; fi30++ {
                                                                        if true {
                                                                            v7 += 9223372036854775807
                                                                            switch ("foo bar" + v3) {
                                                                            case "\n\t":
                                                                                var v31 bool = true
                                                                                _ = v31
                                                                            case "a":
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = 42
                                                                                }
                                                                            default:
                                                                                var v32 uint16 = 187
                                                                                _ = v32
                                                                            }
                                                                        } else {
                                                                            var v33 uint16 = 386
                                                                            _ = v33
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v34 := GenS2{F0: 0, F1: v19}
                                                                            _ = v34
                                                                        case "":
                                                                            if ((v26 % ((fi30) | 1)) < v7) {
                                                                                var v35 uint32 = 627
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v3 = v3
                                                                                _ = v35
                                                                            }
                                                                        case "\n\t":
                                                                            fmt.Println(v3, fi30, v3)
                                                                        }
                                                                        _ = fi30
                                                                    }
                                                                } else {
                                                                    fmt.Println(v7, v19, v7)
                                                                }
                                                            case "":
                                                                v26 += (v26 << ((v19) & 63))
                                                            case "a":
                                                                var v36 uint16 = 1
                                                                _ = v36
                                                            }
                                                        } else {
                                                            v8 = v8
                                                            if ((^(v4)) != 42) {
                                                                for fi37 := 0; fi37 < 3; fi37++ {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = fi37
                                                                }
                                                            }
                                                        }
                                                        _ = v26
                                                    }
                                                }
                                                for fi38 := 0; fi38 < 2; fi38++ {
                                                    v3 = v3
                                                    if ((v7 - (fi10 + -179)) >= -2147483648) {
                                                        if v8 {
                                                            v39 := GenS1{F0: v3, F1: 0, F2: 42}
                                                            var v40 int64 = 390
                                                            var v41 int64 = v40
                                                            _ = v39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    } else {
                                                        var v42 float32 = v1
                                                        _ = v42
                                                    }
                                                    _ = fi38
                                                }
                                            }
                                        }
                                        _ = v19
                                    }
                                    if ((v8 && (v8 || v8)) || v8) {
                                        if (!v8) {
                                            v43 := GenS3{F0: 42, F1: -806, F2: 42}
                                            if (float64((0.1 / v1)) == 1e10) {
                                                var v44 bool = v8
                                                _ = v44
                                            }
                                            _ = v43
                                        } else {
                                            v1 -= float32((((3.14159 - 3.14159) + (float64(v1) * 1e10)) + -2.25))
                                            v8 = v8
                                        }
                                    } else {
                                        _ = len(v2)
                                    }
                                    v4 += 0
                                }
                            }
                            _ = v12
                        }
                    }
                    _ = fi10
                }
            }
            if false {
                v2 = append(v2, 1)
                if (911 >= 2147483647) {
                    fmt.Println(v3, v8, v1)
                    if (((-9223372036854775808 >= 9223372036854775807) || false) && (205 < (v4 * 0))) {
                        if (-2.25 != float64(((v1 * v1) - float32(0.1)))) {
                            for fi45 := 0; fi45 < 7; fi45++ {
                                fmt.Println(fi45)
                                if v8 {
                                    v1 = float32(((0.1 - float64(float32(1e10))) - float64(0.0)))
                                    fmt.Println(v1)
                                    for fi46 := 0; fi46 < 7; fi46++ {
                                        var v47 int64 = 566
                                        var v48 float32 = v1
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                for fi49 := 0; fi49 < 8; fi49++ {
                                    if v8 {
                                        v2 = append(v2, 2147483647)
                                    } else {
                                        var v50 bool = ((-2.25 >= -2.25) || v8)
                                        v51 := GenS4{F0: 18446744073709551615, F1: 0}
                                        var v52 uint64 = v4
                                        if v8 {
                                            switch v3 {
                                            case "hello":
                                                v1 = v1
                                            default:
                                                var v53 uint8 = 0
                                                _ = v53
                                            }
                                            v50 = v50
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        v8 = v50
                                        var v54 uint = 255
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                        _ = v54
                                    }
                                    _ = fi49
                                }
                                _ = fi45
                            }
                        }
                        for fi55 := 0; fi55 < 13; fi55++ {
                            v1 += 3.14159
                            switch fi55 {
                            case 32767:
                                if v8 {
                                    if (1 <= 1) {
                                        if v8 {
                                            fmt.Println(fi55, v1)
                                            v56 := GenS5{F0: 1}
                                            _ = v56
                                        }
                                        if (42 == -32768) {
                                            for fi57 := 0; fi57 < 19; fi57++ {
                                                v58 := GenS6{F0: (v4 ^ v4)}
                                                fmt.Println(fi55, fi55)
                                                v59 := GenS5{F0: -74}
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v4 -= v4
                                                if len(v2) > 0 {
                                                    v2[0] = 1
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483647
                                                }
                                                _ = fi57
                                                _ = v58
                                                _ = v59
                                            }
                                            v7 += (fi55 & -687)
                                        } else {
                                            fmt.Println(v8, v8, v7)
                                            for fi60 := 0; fi60 < 19; fi60++ {
                                                if (!(105 != 0)) {
                                                    v61 := GenS7{F0: 1}
                                                    _ = v61
                                                } else {
                                                    v62 := GenS6{F0: (v4 + v4)}
                                                    var v63 int64 = -9223372036854775808
                                                    _ = v62
                                                    _ = v63
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                fmt.Println(v1)
                                                _ = fi60
                                            }
                                        }
                                    }
                                    for fi64 := 0; fi64 < 3; fi64++ {
                                        v1 -= float32(1.5)
                                        var v65 uint8 = 42
                                        if v8 {
                                            var v66 uint32 = 53
                                            _ = v66
                                        }
                                        if (!true) {
                                            var v67 float64 = float64(v1)
                                            _ = v67
                                        } else {
                                            if v8 {
                                                v68 := GenS8{F0: 16, F1: v65, F2: 4294967295}
                                                if ((!v8) || (100 < 1)) {
                                                    v69 := GenS6{F0: (v4 % ((621) | 1))}
                                                    if v8 {
                                                        switch "" {
                                                        case "":
                                                            v2 = append(v2, 827)
                                                        default:
                                                            switch (fi64 ^ v7) {
                                                            case 42:
                                                                switch 4294967295 {
                                                                case -808:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case 9223372036854775807:
                                                                    var v70 int64 = 0
                                                                    _ = v70
                                                                default:
                                                                    var v71 string = (("" + "a") + v3)
                                                                    _ = v71
                                                                }
                                                            default:
                                                                var v72 int16 = 0
                                                                _ = v72
                                                            }
                                                        }
                                                    } else {
                                                        var v73 int = (fi55 ^ fi55)
                                                        _ = v73
                                                    }
                                                    if true {
                                                        var v74 int8 = 42
                                                        v8 = false
                                                        _ = v74
                                                    } else {
                                                        if v8 {
                                                            switch v3 {
                                                            case "":
                                                                if (v1 != (v1 - v1)) {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                v75 := GenS6{F0: (v4 | (v4 + v4))}
                                                                _ = v75
                                                            case "hello":
                                                                v1 += (v1 / v1)
                                                            }
                                                            if v8 {
                                                                if v8 {
                                                                    fmt.Println(v1, v8)
                                                                } else {
                                                                    var v76 float64 = float64(v1)
                                                                    _ = v76
                                                                }
                                                                var v77 string = v3
                                                                _ = v77
                                                            }
                                                        }
                                                    }
                                                    _ = v69
                                                } else {
                                                    var v78 int16 = 19
                                                    var v79 int8 = 42
                                                    _ = v78
                                                    _ = v79
                                                }
                                                _ = v68
                                            }
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi64
                                        _ = v65
                                    }
                                    v80 := GenS4{F0: 0, F1: 0}
                                    _ = v80
                                } else {
                                    var v81 uint16 = 0
                                    fmt.Println(v7, v7, v7)
                                    var v82 uint = 4294967295
                                    _ = v81
                                    _ = v82
                                }
                            case 128:
                                var v83 uint8 = 42
                                _ = v83
                            default:
                                for fi84 := 0; fi84 < 14; fi84++ {
                                    var v85 uint8 = 144
                                    _ = fi84
                                    _ = v85
                                }
                            }
                            _ = fi55
                        }
                    }
                }
                var v86 int32 = -2147483648
                _ = v86
            }
            if len(v2) > 0 {
                v2[0] = 506
            }
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (v4 < v4) {
                for fi87 := 0; fi87 < 20; fi87++ {
                    if (1 < -103) {
                        v1 -= float32(((1e10 - 1e10) + 3.14159))
                    } else {
                        fmt.Println(v1)
                    }
                    _ = fi87
                }
                var v88 bool = v8
                _ = v88
            } else {
                for fi89 := 0; fi89 < 10; fi89++ {
                    for fi90 := 0; fi90 < 5; fi90++ {
                        v91 := GenS6{F0: (v4 / ((v4) | 1))}
                        var v92 uint8 = 112
                        _ = len(v2)
                        _ = fi90
                        _ = v91
                        _ = v92
                    }
                    _ = fi89
                }
            }
            fmt.Println(v8)
        }
        var v93 float64 = 3.14159
        _ = v93
    }
    fmt.Println(v3)
    for fi94 := 0; fi94 < 5; fi94++ {
        for fi95 := 0; fi95 < 18; fi95++ {
            v8 = (fi95 != fi94)
            if false {
                var v96 uint16 = 391
                var v97 uint32 = 1
                if len(v2) > 0 {
                    v2[0] = uint32(v96)
                }
                if ((^(v4)) >= (^(v4))) {
                    v98 := GenS9{F0: 42, F1: -1, F2: 32767}
                    var v99 uint32 = v97
                    switch (fi95 ^ fi94) {
                    case -2147483648:
                        var v100 uint16 = 2
                        _ = v100
                    case -9223372036854775808:
                        if true {
                            v7 = v7
                        }
                    }
                    _ = v98
                    _ = v99
                }
                if v8 {
                    var v101 float32 = float32((-2.25 + (3.14159 - 3.14159)))
                    if true {
                        var v102 int64 = 7
                        _ = v102
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v101
                } else {
                    v103 := GenS10{F0: 635}
                    fmt.Println(v97)
                    v3 = ""
                    _ = v103
                }
                _ = v96
                _ = v97
            }
            _ = fi95
        }
        fmt.Println(v7)
        _ = fi94
    }
    if false {
        for fi104 := 0; fi104 < 7; fi104++ {
            v2 = append(v2, 65535)
            _ = fi104
        }
    } else {
        if len(v2) > 0 {
            v2[0] = 65535
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v105 uint32 = 162
        _ = v105
    }
    for fi106 := 0; fi106 < 17; fi106++ {
        if (804 > 0) {
            fmt.Println(fi106, v1, v3)
            switch v3 {
            case "hello":
                v107 := GenS6{F0: (v4 * (^(v4)))}
                _ = v107
            case "\n\t":
                var v108 int = -303
                _ = v108
            default:
                if (2 != 114) {
                    for fi109 := 0; fi109 < 14; fi109++ {
                        if v8 {
                            var v110 string = v3
                            _ = v110
                        } else {
                            switch -1 {
                            case -477:
                                var v111 int64 = -866
                                _ = v111
                            case 42:
                                v112 := GenS7{F0: 371}
                                _ = v112
                            default:
                                v3 = v3
                                if (403 < 2) {
                                    var v113 uint32 = 2147483648
                                    _ = v113
                                } else {
                                    if false {
                                        v4 = 749
                                        switch v3 {
                                        case "foo bar":
                                            fmt.Println(v3, v3)
                                        case "\n\t":
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        case "":
                                            var v114 uint32 = 369
                                            var v115 uint16 = 1
                                            switch (v3 + "") {
                                            case "\n\t":
                                                if (((((v114 + v114) == (v114 + v114)) || v8) || v8) && (v8 || (v8 || v8))) {
                                                    v1 += v1
                                                } else {
                                                    var v116 uint64 = 940
                                                    _ = v116
                                                }
                                            case "hello":
                                                var v117 int32 = 0
                                                _ = v117
                                            case "":
                                                v114 += (^(v114))
                                            default:
                                                v4 -= v4
                                            }
                                            _ = v114
                                            _ = v115
                                        }
                                    }
                                    v8 = false
                                    v118 := GenS5{F0: 127}
                                    var v119 uint8 = 183
                                    v119 += 42
                                    _ = v118
                                    _ = v119
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                        }
                        if (-310 < 198) {
                            var v120 uint64 = v4
                            _ = len(v2)
                            _ = v120
                        } else {
                            fmt.Println(fi106, v1)
                        }
                        v121 := GenS11{F0: 111, F1: -2147483648}
                        var v122 uint8 = 42
                        var v123 int64 = 42
                        _ = fi109
                        _ = v121
                        _ = v122
                        _ = v123
                    }
                    var v124 uint64 = (^(v4))
                    _ = v124
                }
            }
            if (-954 <= -9223372036854775808) {
                var v125 int32 = 42
                fmt.Println(v8, v4)
                _ = v125
            } else {
                v2 = append(v2, 0)
            }
        }
        _ = fi106
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
