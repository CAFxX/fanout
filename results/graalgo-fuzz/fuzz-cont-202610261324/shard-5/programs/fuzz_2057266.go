package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint }
type GenS7 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 uint; F1 uint32 }
type GenS10 struct { F0 uint32; F1 int32; F2 int8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []uint32{741, 42, 2147483647, 497}
    var v2 int32 = 42
    var v3 int16 = 42
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v4 uint8 = 15
    v1 = append(v1, 65535)
    if len(v1) > 0 {
        v1[0] = 143
    }
    v2 = (v2 << ((358) & 63))
    for fi5 := 0; fi5 < 18; fi5++ {
        if len(v1) > 0 {
            v1[0] = 2147483647
        }
        v3 = 42
        if false {
            var v6 uint32 = 0
            for fi7 := 0; fi7 < 14; fi7++ {
                if false {
                    if (0 != -59) {
                        v2 = v2
                        if true {
                            switch "\n\t" {
                            case "":
                                var v8 string = "a"
                                _ = v8
                            case "a":
                                fmt.Println(v3)
                            case "\n\t":
                                v3 = -1
                            default:
                                v9 := GenS1{F0: (fi7 + fi5)}
                                _ = v9
                            }
                            v4 = v4
                        } else {
                            if ("a" >= "\n\t") {
                                if ((!true) && (949 > 18446744073709551615)) {
                                    v6 -= 2147483648
                                    var v10 uint = 0
                                    _ = v10
                                }
                                if (true || false) {
                                    v6 = v6
                                } else {
                                    v4 = (^(v4))
                                }
                                fmt.Println(v2, fi5)
                            }
                        }
                        var v11 int = -9223372036854775808
                        _ = v11
                    } else {
                        switch fi5 {
                        case 215:
                            if false {
                                v12 := GenS1{F0: (fi7 % (((fi5 >> ((1) & 63))) | 1))}
                                _ = v12
                            } else {
                                var v13 uint8 = v4
                                _ = v13
                            }
                        default:
                            var v14 int8 = 1
                            _ = v14
                        }
                        v15 := GenS1{F0: (^(fi5))}
                        v6 = (^(v6))
                        _ = v15
                    }
                }
                v6 = (v6 / ((0) | 1))
                _ = fi7
            }
            _ = v6
        } else {
            v4 += v4
            fmt.Println(v3)
        }
        _ = fi5
    }
    var v16 float32 = 1.5
    fmt.Println(v16, v4, v16)
    var v17 bool = true
    if v17 {
        v18 := GenS1{F0: -613}
        var v19 int8 = 127
        if (v2 != (v2 + 1)) {
            v2 += 2147483647
        } else {
            var v20 uint64 = 447
            v3 += v3
            v21 := GenS1{F0: 42}
            _ = v20
            _ = v21
        }
        _ = v18
        _ = v19
    } else {
        var v22 string = ""
        _ = v22
    }
    if len(v1) > 0 {
        v1[0] = 541
    }
    if v17 {
        switch ("\n\t" + "") {
        case "a":
            v1 = append(v1, 4294967295)
        default:
            switch "a" {
            case "\n\t":
                v23 := GenS1{F0: int(v2)}
                _ = v23
            case "a":
                v24 := GenS1{F0: 0}
                _ = v24
            case "":
                switch 1 {
                case 841:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v2 = (v2 * -272)
                        _ = fi25
                    }
                case -2147483648:
                    v17 = (1000 >= v3)
                case 0:
                    var v26 int = 4294967295
                    _ = v26
                }
            default:
                var v27 string = ("a" + "")
                _ = v27
            }
        }
        if (true || v17) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            if ((42 > -92) || v17) {
                for fi29 := 0; fi29 < 8; fi29++ {
                    switch "a" {
                    case "hello":
                        fmt.Println(v17, fi28, v17)
                        if (((("\n\t" + "hello") + ("hello" + "foo bar")) + ("" + ("" + ("hello" + "\n\t")))) >= "hello") {
                            v17 = true
                            if false {
                                v4 = 42
                                if v17 {
                                    fmt.Println(v4, v2, v4)
                                    if v17 {
                                        fmt.Println(fi28, v4, v17)
                                    }
                                    v3 += v3
                                }
                            } else {
                                var v30 int16 = v3
                                _ = v30
                            }
                            v31 := GenS1{F0: fi28}
                            _ = v31
                        }
                        var v32 string = "hello"
                        _ = v32
                    case "\n\t":
                        var v33 uint = 18446744073709551615
                        _ = v33
                    case "":
                        v34 := GenS1{F0: (fi29 | (-(fi29)))}
                        _ = len(v1)
                        _ = v34
                    }
                    fmt.Println(fi28, v3, fi29)
                    _ = fi29
                }
            }
            switch "hello" {
            case "hello":
                v17 = false
            case "foo bar":
                switch "\n\t" {
                case "":
                    var v35 bool = true
                    _ = v35
                case "\n\t":
                    if (42 >= 0) {
                        for fi36 := 0; fi36 < 16; fi36++ {
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(v4, v3, v4)
                            case "\n\t":
                                if (116 > 127) {
                                    var v37 int16 = (v3 % (((v3 * (v3 >> ((255) & 63)))) | 1))
                                    v38 := GenS2{F0: v4}
                                    v16 += v16
                                    var v39 int = fi36
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            case "a":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = fi36
                        }
                        fmt.Println(v2, v3, v2)
                    }
                default:
                    v40 := GenS1{F0: (fi28 + 1)}
                    _ = v40
                }
            default:
                fmt.Println(v2, v17)
            }
            v41 := GenS3{F0: -1}
            _ = fi28
            _ = v41
        }
    }
    if (255 != 42) {
        for fi42 := 0; fi42 < 8; fi42++ {
            switch "a" {
            case "\n\t":
                v43 := GenS4{F0: 0, F1: v3}
                _ = v43
            case "hello":
                fmt.Println(v2, v17, v16)
            case "":
                for fi44 := 0; fi44 < 8; fi44++ {
                    switch "" {
                    case "a":
                        v45 := GenS3{F0: 100}
                        _ = v45
                    default:
                        if ((!v17) && false) {
                            if (!(!(false || false))) {
                                _ = len(v1)
                                v2 = int32(fi44)
                                v1 = append(v1, 2)
                                fmt.Println(v3, v4)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v46 uint8 = v4
                                _ = v46
                            }
                            v3 = v3
                        } else {
                            fmt.Println(v4, v2, v4)
                            var v47 uint8 = (^(v4))
                            switch fi44 {
                            case 678:
                                switch "" {
                                case "foo bar":
                                    v17 = true
                                }
                            }
                            _ = v47
                        }
                        if true {
                            v4 -= (v4 % ((v4) | 1))
                            v48 := GenS5{F0: (fi44 - (^(fi42))), F1: 42, F2: int16(v2)}
                            fmt.Println(v3, v2, v4)
                            _ = v48
                        }
                        var v49 int16 = (v3 << ((2) & 63))
                        var v50 int16 = (v49 % ((v49) | 1))
                        _ = v49
                        _ = v50
                    }
                    _ = fi44
                }
            default:
                if (1 >= 0) {
                    if v17 {
                        for fi51 := 0; fi51 < 8; fi51++ {
                            var v52 int32 = (v2 * (v2 >> ((2) & 63)))
                            _ = fi51
                            _ = v52
                        }
                    } else {
                        if ((v17 || v17) || v17) {
                            var v53 int64 = 0
                            v54 := GenS6{F0: 0, F1: (v2 >> ((uint(v4)) & 63)), F2: 65535}
                            _ = v53
                            _ = v54
                        } else {
                            var v55 uint64 = 1
                            for fi56 := 0; fi56 < 8; fi56++ {
                                switch ("" + ((("\n\t" + "foo bar") + "hello") + ("\n\t" + "hello"))) {
                                case "\n\t":
                                    v17 = ((v4 + (v4 | 255)) <= 184)
                                case "":
                                    var v57 string = "a"
                                    _ = v57
                                }
                                fmt.Println(fi42, v4)
                                _ = fi56
                            }
                            _ = v55
                        }
                    }
                } else {
                    fmt.Println(fi42, v17)
                    v3 = (v3 * v3)
                    var v58 int32 = 0
                    _ = v58
                }
            }
            if ((v3 | v3) >= v3) {
                v59 := GenS7{F0: 255, F1: (v4 & 2), F2: 669}
                var v60 uint64 = 2
                _ = v59
                _ = v60
            } else {
                if v17 {
                    v1 = append(v1, 199)
                    var v61 int16 = v3
                    v1 = append(v1, 2147483647)
                    _ = v61
                }
                switch int(v4) {
                case 42:
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v63 := GenS8{F0: "foo bar", F1: 9223372036854775807}
                        if v17 {
                            _ = len(v1)
                            if (65535 >= 4294967295) {
                                v2 = v2
                                v17 = true
                                switch fi62 {
                                case 127:
                                    v64 := GenS5{F0: fi62, F1: 1, F2: v3}
                                    _ = v64
                                case 4294967295:
                                    v65 := GenS3{F0: -5}
                                    fmt.Println(v3, v4)
                                    var v66 int8 = 100
                                    _ = v65
                                    _ = v66
                                case 1:
                                    if (v16 != 0.1) {
                                        v1 = append(v1, 4294967295)
                                        if (false && v17) {
                                            var v67 uint8 = (^(v4))
                                            _ = v67
                                        }
                                        v4 += 128
                                    } else {
                                        v68 := GenS6{F0: 42, F1: -823, F2: 42}
                                        fmt.Println(v3)
                                        if (v4 != v4) {
                                            var v69 int16 = v3
                                            if false {
                                                v17 = (((v3 - v69) > -134) || (fi62 != 1))
                                                switch "hello" {
                                                case "foo bar":
                                                    v1 = append(v1, 42)
                                                case "a":
                                                    fmt.Println(v17)
                                                }
                                                if v17 {
                                                    fmt.Println(v3)
                                                } else {
                                                    _ = len(v1)
                                                }
                                                if (!v17) {
                                                    fmt.Println(v4, v2)
                                                }
                                                if true {
                                                    v17 = v17
                                                }
                                                if (!(v17 && v17)) {
                                                    v2 -= 2147483647
                                                } else {
                                                    switch ("" + "foo bar") {
                                                    case "foo bar":
                                                        fmt.Println(fi62)
                                                    case "":
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483647
                                                        }
                                                        if (!(42 >= (v2 + v2))) {
                                                            var v70 int8 = 15
                                                            _ = v70
                                                        }
                                                        var v71 int32 = 1
                                                        _ = v71
                                                    case "\n\t":
                                                        if v17 {
                                                            v4 = v4
                                                            v4 = 42
                                                            v16 = v16
                                                            var v72 bool = v17
                                                            v73 := GenS7{F0: 151, F1: (v4 / ((2) | 1)), F2: 18446744073709551615}
                                                            _ = v72
                                                            _ = v73
                                                        }
                                                    default:
                                                        var v74 int32 = 401
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            _ = v69
                                        }
                                        v2 -= v2
                                        _ = v68
                                    }
                                default:
                                    v17 = true
                                }
                                v75 := GenS9{F0: 42, F1: 2}
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                                switch -122 {
                                case 2147483647:
                                    var v76 float32 = v16
                                    _ = v76
                                case -9223372036854775808:
                                    fmt.Println(v4, v16)
                                    if (!v17) {
                                        _ = len(v1)
                                    } else {
                                        switch ("foo bar" + "\n\t") {
                                        case "\n\t":
                                            switch ("foo bar" + "\n\t") {
                                            case "a":
                                                if false {
                                                    var v77 int = fi62
                                                    var v78 int32 = (v2 + -2147483648)
                                                    v77 = -385
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    v17 = v17
                                                }
                                                v79 := GenS10{F0: 60, F1: -1, F2: -70}
                                                v17 = ("" != "")
                                                _ = v79
                                            case "":
                                                switch "hello" {
                                                case "":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            case "hello":
                                                switch -1 {
                                                case 32767:
                                                    v1 = append(v1, 468)
                                                case 9223372036854775807:
                                                    var v80 uint64 = 207
                                                    _ = v80
                                                }
                                            }
                                        case "a":
                                            if true {
                                                if v17 {
                                                    fmt.Println(v16, v4)
                                                    var v81 int64 = -9223372036854775808
                                                    var v82 uint32 = 1
                                                    _ = v81
                                                    _ = v82
                                                } else {
                                                    switch 4294967295 {
                                                    case -969:
                                                        _ = len(v1)
                                                    case 4294967295:
                                                        var v83 int64 = -1
                                                        _ = v83
                                                    case 127:
                                                        v16 += v16
                                                        fmt.Println(fi62, v16)
                                                        if (v4 <= v4) {
                                                            var v84 uint8 = v4
                                                            if (627 != 468) {
                                                                v3 -= (v3 ^ v3)
                                                            }
                                                            if true {
                                                                v1 = append(v1, 0)
                                                                var v85 uint32 = 2147483648
                                                                v85 = uint32(v84)
                                                                _ = v85
                                                            }
                                                            _ = v84
                                                        }
                                                    }
                                                    fmt.Println(v2, v16, v2)
                                                }
                                                if (((!v17) || false) && (!(32767 == 1))) {
                                                    if true {
                                                        if (!true) {
                                                            v1 = append(v1, 360)
                                                            if v17 {
                                                                var v86 int64 = 330
                                                                if len(v1) > 0 {
                                                                    v1[0] = 439
                                                                }
                                                                v87 := GenS11{F0: (v4 | 42)}
                                                                _ = v86
                                                                _ = v87
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = 4294967295
                                                            }
                                                        }
                                                        if v17 {
                                                            if v17 {
                                                                var v88 int16 = 42
                                                                _ = v88
                                                            }
                                                            v2 = (v2 % ((1) | 1))
                                                            if (125 <= (^(v4))) {
                                                                v17 = (-9223372036854775808 >= 1)
                                                            }
                                                            var v89 int = (fi62 & 2147483647)
                                                            var v90 uint = 2
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v89
                                                            _ = v90
                                                        } else {
                                                            v91 := GenS12{F0: 0, F1: 173}
                                                            var v92 int64 = 9223372036854775807
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        var v93 uint = 65535
                                                        v94 := GenS12{F0: 42, F1: (v93 * (^(v93)))}
                                                        _ = v93
                                                        _ = v94
                                                    } else {
                                                        v95 := GenS9{F0: 4294967295, F1: 2147483648}
                                                        _ = v95
                                                    }
                                                }
                                            }
                                        case "":
                                            v96 := GenS9{F0: 2, F1: 4294967295}
                                            _ = v96
                                        default:
                                            v3 -= v3
                                        }
                                        v4 -= 149
                                    }
                                    v97 := GenS13{F0: v3, F1: 902}
                                    _ = v97
                                case 4294967295:
                                    var v98 int = (fi42 - 4294967295)
                                    _ = v98
                                }
                                _ = v75
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v4 = 2
                                v16 = v16
                            }
                            if (0 != -128) {
                                if (-271 >= 991) {
                                    v99 := GenS9{F0: 406, F1: 2}
                                    var v100 uint = 2
                                    _ = v100
                                    _ = v99
                                }
                                var v101 string = "a"
                                if false {
                                    v16 = v16
                                    var v102 int64 = 42
                                    _ = v102
                                }
                                _ = v101
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                                switch (fi62 % ((fi62) | 1)) {
                                case 471:
                                    var v103 int64 = 42
                                    _ = v103
                                default:
                                    var v104 uint64 = 751
                                    _ = v104
                                }
                            }
                            v1 = append(v1, 2147483647)
                        } else {
                            v16 += v16
                            var v105 bool = false
                            fmt.Println(v2, fi62)
                            _ = v105
                        }
                        if false {
                            var v106 uint32 = uint32(v4)
                            if v17 {
                                fmt.Println(v16, v3, fi62)
                            }
                            _ = v106
                        } else {
                            var v107 int16 = (v3 & (v3 << ((0) & 63)))
                            v108 := GenS14{F0: (v107 >> ((18446744073709551615) & 63))}
                            fmt.Println(v3, v17)
                            var v109 int64 = -9223372036854775808
                            _ = v107
                            _ = v108
                            _ = v109
                        }
                        _ = fi62
                        _ = v63
                    }
                case 4294967295:
                    v3 = 42
                default:
                    for fi110 := 0; fi110 < 10; fi110++ {
                        fmt.Println(v16)
                        _ = fi110
                    }
                }
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
            v111 := GenS15{F0: 526, F1: 42}
            _ = fi42
            _ = v111
        }
        var v112 float64 = 3.14159
        _ = v112
    } else {
        v17 = ((v4 - 31) == (v4 % ((1) | 1)))
    }
    for fi113 := 0; fi113 < 19; fi113++ {
        for fi114 := 0; fi114 < 2; fi114++ {
            fmt.Println(v3, v2, fi113)
            v3 = (v3 - (v3 << ((65535) & 63)))
            _ = fi114
        }
        _ = fi113
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
