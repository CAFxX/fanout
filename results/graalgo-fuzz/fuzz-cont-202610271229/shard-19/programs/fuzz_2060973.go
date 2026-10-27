package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 bool; F2 int64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32; F1 int }
func genFn0(p0 int32) int64 {
    p0 += p0
    p0 -= p0
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := []int{}
    var v2 uint32 = 477
    var v3 float32 = 1.5
    if len(v1) > 0 {
        v1[0] = 0
    }
    if (true || true) {
        v3 -= 1e10
    }
    var v4 int32 = -283
    switch 59 {
    case 0:
        for fi5 := 0; fi5 < 13; fi5++ {
            v4 = v4
            _ = fi5
        }
    case 42:
        var v6 int16 = 923
        _ = v6
    case -2147483648:
        if true {
            v7 := GenS1{F0: 632, F1: 2}
            v8 := GenS1{F0: 0, F1: 0}
            var v9 bool = false
            _ = v7
            _ = v8
            _ = v9
        }
    default:
        v10 := GenS2{F0: -45}
        _ = v10
    }
    var v11 float32 = ((1e10 + (1e10 + v3)) / v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v12 uint = 2
        var v13 int16 = 1
        v12 += (^(v12))
        _ = v12
        _ = v13
    } else {
        if len(v1) > 0 {
            v1[0] = 32768
        }
        if len(v1) > 0 {
            v1[0] = -922
        }
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v14 := GenS2{F0: -88}
    v15 := GenS2{F0: 44}
    var v16 int64 = 965
    var v17 uint = 1
    v17 -= (v17 << ((v17) & 63))
    var v18 float32 = ((v3 * (0.0 + (v11 * v11))) + -2.25)
    if false {
        fmt.Println(v3, v17)
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                v11 = v3
                fmt.Println(v17, v18, fi20)
                _ = fi20
            }
            var v21 uint8 = 128
            var v22 int8 = -84
            _ = fi19
            _ = v21
            _ = v22
        }
    } else {
        if false {
            _ = len(v1)
            var v23 int32 = (v4 << ((v17) & 63))
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "hello":
                if (!((v2 * (^(v2))) > 729)) {
                    if false {
                        var v24 int = 658
                        fmt.Println(v3, v2, v16)
                        if len(v1) > 0 {
                            v1[0] = v24
                        }
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 8; fi25++ {
                            fmt.Println(v11, v18)
                            _ = fi25
                        }
                        v2 = (v2 >> ((v17) & 63))
                        if (v2 != (v2 << ((v17) & 63))) {
                            for fi26 := 0; fi26 < 7; fi26++ {
                                for fi27 := 0; fi27 < 7; fi27++ {
                                    v28 := GenS2{F0: 127}
                                    _ = fi27
                                    _ = v28
                                }
                                var v29 bool = false
                                _ = fi26
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 9; fi30++ {
                                var v31 uint64 = 246
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = (fi30 & (fi30 >> ((4294967295) & 63)))
                                    }
                                } else {
                                    switch (fi30 << ((v17) & 63)) {
                                    case 1:
                                        v23 += v23
                                    case -9223372036854775808:
                                        if false {
                                            v32 := GenS3{F0: "\n\t", F1: 42}
                                            _ = len(v1)
                                            _ = v32
                                        } else {
                                            var v33 float32 = -2.25
                                            fmt.Println(v17)
                                            for fi34 := 0; fi34 < 8; fi34++ {
                                                if (-9223372036854775808 > (v16 / ((v16) | 1))) {
                                                    var v35 float32 = 0.1
                                                    _ = v35
                                                }
                                                switch ("hello" + "") {
                                                case "":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v18)
                                                    fmt.Println(v23)
                                                    if (!(!(514 < v2))) {
                                                        switch fi34 {
                                                        case 32768:
                                                            if (255 != 302) {
                                                                var v36 int = fi30
                                                                _ = v36
                                                            } else {
                                                                v37 := GenS1{F0: 255, F1: v17}
                                                                _ = v37
                                                            }
                                                        case 0:
                                                            v1 = append(v1, -542)
                                                        default:
                                                            var v38 int8 = 42
                                                            _ = v38
                                                        }
                                                        switch fi30 {
                                                        case 0:
                                                            v16 += (v16 * (v16 & 362))
                                                        case 57:
                                                            switch fi34 {
                                                            case 42:
                                                                if true {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v17 += 4294967295
                                                                    case "\n\t":
                                                                        v39 := GenS2{F0: -128}
                                                                        _ = v39
                                                                    }
                                                                    v40 := GenS2{F0: -1}
                                                                    _ = v40
                                                                } else {
                                                                    fmt.Println(v23, v18, v2)
                                                                    v33 -= v33
                                                                    var v41 float64 = -2.25
                                                                    _ = v41
                                                                }
                                                            case 127:
                                                                var v42 uint16 = 32767
                                                                _ = v42
                                                            case 75:
                                                                v43 := GenS3{F0: "hello", F1: (v16 ^ genFn0(-253))}
                                                                _ = v43
                                                            default:
                                                                var v44 uint32 = v2
                                                                _ = v44
                                                            }
                                                        case 9223372036854775807:
                                                            v17 -= v17
                                                        }
                                                        if (1 > 1) {
                                                            v45 := GenS2{F0: 127}
                                                            var v46 string = "\n\t"
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    }
                                                case "hello":
                                                    v1 = append(v1, fi30)
                                                default:
                                                    v47 := GenS1{F0: 0, F1: 65535}
                                                    _ = v47
                                                }
                                                v17 += v17
                                                _ = fi34
                                            }
                                            _ = v33
                                        }
                                    }
                                    for fi48 := 0; fi48 < 5; fi48++ {
                                        var v49 int8 = 100
                                        _ = len(v1)
                                        var v50 uint64 = (v31 & 18446744073709551615)
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                    if true {
                                        if false {
                                            v51 := GenS1{F0: 783, F1: v17}
                                            fmt.Println(v4)
                                            _ = v51
                                        } else {
                                            var v52 int8 = 42
                                            switch 128 {
                                            case 9223372036854775807:
                                                _ = len(v1)
                                            default:
                                                for fi53 := 0; fi53 < 6; fi53++ {
                                                    v54 := GenS3{F0: ("a" + "hello"), F1: (v16 & v16)}
                                                    var v55 int64 = (v16 + (^(v16)))
                                                    _ = fi53
                                                    _ = v54
                                                    _ = v55
                                                }
                                            }
                                            _ = v52
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = (fi30 / ((fi30) | 1))
                                        }
                                    }
                                }
                                _ = fi30
                                _ = v31
                            }
                            var v56 float32 = v11
                            _ = v56
                        }
                    }
                    v57 := GenS1{F0: 2, F1: (v17 / ((v17) | 1))}
                    fmt.Println(v23, v4)
                    _ = v57
                }
            case "\n\t":
                switch "foo bar" {
                case "":
                    if (false && (42 >= 65535)) {
                        switch "a" {
                        case "hello":
                            if false {
                                switch int(v23) {
                                case 9223372036854775807:
                                    if (v18 != (v18 / v18)) {
                                        var v58 int = 128
                                        switch (v58 * v58) {
                                        case 0:
                                            switch (v58 + (v58 * (v58 & (^(v58))))) {
                                            case -459:
                                                v1 = append(v1, v58)
                                            case 32767:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 587:
                                                switch "hello" {
                                                case "hello":
                                                    for fi59 := 0; fi59 < 1; fi59++ {
                                                        var v60 int32 = v23
                                                        v2 -= v2
                                                        _ = fi59
                                                        _ = v60
                                                    }
                                                case "a":
                                                    v2 = (v2 + 1)
                                                case "\n\t":
                                                    if true {
                                                        v1 = append(v1, (v58 / ((v58) | 1)))
                                                        if ((v17 | (^(v17))) >= (v17 - v17)) {
                                                            if (32767 <= 255) {
                                                                for fi61 := 0; fi61 < 6; fi61++ {
                                                                    fmt.Println(fi61, v18, fi61)
                                                                    _ = fi61
                                                                }
                                                                fmt.Println(v18, v4, v4)
                                                                var v62 int16 = 42
                                                                _ = v62
                                                            }
                                                            for fi63 := 0; fi63 < 4; fi63++ {
                                                                switch 4294967295 {
                                                                case 0:
                                                                    for fi64 := 0; fi64 < 15; fi64++ {
                                                                        v65 := GenS4{F0: 239}
                                                                        _ = fi64
                                                                        _ = v65
                                                                    }
                                                                }
                                                                _ = fi63
                                                            }
                                                        } else {
                                                            var v66 uint8 = uint8(v17)
                                                            _ = v66
                                                        }
                                                    } else {
                                                        v67 := GenS1{F0: 390, F1: v17}
                                                        _ = v67
                                                    }
                                                default:
                                                    var v68 uint64 = 1
                                                    _ = v68
                                                }
                                            default:
                                                v69 := GenS5{F0: 231, F1: ((v17 << ((v17) & 63)) < v17), F2: (v16 - (v16 - v16))}
                                                _ = v69
                                            }
                                        case -9223372036854775808:
                                            for fi70 := 0; fi70 < 1; fi70++ {
                                                _ = len(v1)
                                                var v71 uint8 = 2
                                                _ = fi70
                                                _ = v71
                                            }
                                        }
                                        _ = v58
                                    } else {
                                        fmt.Println(v18, v23)
                                    }
                                case 32768:
                                    var v72 uint16 = 42
                                    _ = v72
                                case 32767:
                                    if (false && (v11 <= float32(0.0))) {
                                        for fi73 := 0; fi73 < 9; fi73++ {
                                            switch (("a" + ("hello" + "hello")) + ("" + ("a" + "a"))) {
                                            case "":
                                                fmt.Println(v16)
                                            case "a":
                                                if false {
                                                    fmt.Println(v16, fi73)
                                                } else {
                                                    if true {
                                                        v74 := GenS4{F0: fi73}
                                                        switch fi73 {
                                                        case 90:
                                                            var v75 int16 = 57
                                                            _ = v75
                                                        case -2147483648:
                                                            var v76 int8 = -29
                                                            _ = v76
                                                        }
                                                        v17 -= 255
                                                        _ = v74
                                                    } else {
                                                        v77 := GenS2{F0: -1}
                                                        var v78 int32 = 354
                                                        if len(v1) > 0 {
                                                            v1[0] = -2147483648
                                                        }
                                                        _ = v77
                                                        _ = v78
                                                    }
                                                }
                                            default:
                                                for fi79 := 0; fi79 < 2; fi79++ {
                                                    v17 = v17
                                                    _ = fi79
                                                }
                                            }
                                            _ = fi73
                                        }
                                        var v80 int64 = (v16 ^ v16)
                                        var v81 uint64 = uint64(v17)
                                        _ = v80
                                        _ = v81
                                    }
                                }
                                var v82 uint = v17
                                fmt.Println(v2, v18)
                                _ = v82
                            }
                        case "foo bar":
                            if false {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if (v16 >= v16) {
                                    fmt.Println(v4, v23, v2)
                                }
                            }
                        case "a":
                            if (!(!(2 >= 255))) {
                                v11 += v11
                            } else {
                                v11 += v18
                            }
                        default:
                            if true {
                                switch int(v4) {
                                case 0:
                                    if true {
                                        var v83 uint16 = 0
                                        for fi84 := 0; fi84 < 6; fi84++ {
                                            var v85 uint64 = 42
                                            var v86 uint = (v17 << ((v17) & 63))
                                            for fi87 := 0; fi87 < 20; fi87++ {
                                                var v88 float32 = v18
                                                var v89 uint16 = (v83 << (((^(v86))) & 63))
                                                v11 = v3
                                                v90 := GenS6{F0: v23}
                                                fmt.Println(v89)
                                                _ = fi87
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                            }
                                            _ = fi84
                                            _ = v85
                                            _ = v86
                                        }
                                        v4 -= v4
                                        _ = v83
                                    }
                                case 127:
                                    var v91 uint16 = 65535
                                    _ = v91
                                case 9223372036854775807:
                                    v3 = (v3 / v3)
                                }
                                var v92 int8 = 0
                                for fi93 := 0; fi93 < 14; fi93++ {
                                    var v94 uint8 = 0
                                    var v95 uint8 = 0
                                    v17 = v17
                                    _ = fi93
                                    _ = v94
                                    _ = v95
                                }
                                _ = v92
                            } else {
                                for fi96 := 0; fi96 < 8; fi96++ {
                                    if len(v1) > 0 {
                                        v1[0] = 32768
                                    }
                                    v97 := GenS1{F0: 0, F1: 18446744073709551615}
                                    _ = fi96
                                    _ = v97
                                }
                                fmt.Println(v18, v2, v23)
                            }
                        }
                    } else {
                        v98 := GenS7{F0: uint32(v17), F1: 32767}
                        if ((v23 ^ v4) == (v4 | (v23 | v23))) {
                            fmt.Println(v11)
                            var v99 int16 = 1000
                            _ = v99
                        }
                        _ = v98
                    }
                case "\n\t":
                    var v100 int16 = 0
                    _ = v100
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    if (13 != 916) {
                        var v101 float32 = (v18 / 0.1)
                        if true {
                            v2 = v2
                            var v102 bool = (!false)
                            var v103 int32 = 474
                            _ = v102
                            _ = v103
                        } else {
                            var v104 int = 32767
                            fmt.Println(v18, v101, v11)
                            for fi105 := 0; fi105 < 13; fi105++ {
                                var v106 uint = 0
                                if len(v1) > 0 {
                                    v1[0] = fi105
                                }
                                _ = fi105
                                _ = v106
                            }
                            _ = v104
                        }
                        _ = v101
                    } else {
                        switch ("hello" + ("foo bar" + "hello")) {
                        case "hello":
                            _ = len(v1)
                        case "a":
                            switch 175 {
                            case 128:
                                v107 := GenS1{F0: 42, F1: v17}
                                _ = v107
                            case 42:
                                var v108 uint32 = 2147483648
                                _ = v108
                            default:
                                var v109 float64 = 1e10
                                _ = v109
                            }
                        case "":
                            for fi110 := 0; fi110 < 1; fi110++ {
                                for fi111 := 0; fi111 < 19; fi111++ {
                                    var v112 uint8 = 2
                                    _ = fi111
                                    _ = v112
                                }
                                var v113 bool = true
                                v23 += v4
                                _ = fi110
                                _ = v113
                            }
                        }
                        fmt.Println(v17, v3)
                        v114 := GenS4{F0: 9223372036854775807}
                        _ = v114
                    }
                }
            case "foo bar":
                v16 = -948
            }
            _ = v23
        }
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    var v115 float32 = v3
    v116 := GenS5{F0: 32767, F1: ((1 > 1000) && false), F2: genFn0(int32(v2))}
    var v117 int = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v115:", v115)
    fmt.Println("v117:", v117)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v115
    _ = v116
    _ = v117
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
