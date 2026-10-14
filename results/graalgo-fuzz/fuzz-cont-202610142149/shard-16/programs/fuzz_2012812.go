package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 int64; F1 bool; F2 uint64 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint; F1 string; F2 string }
func genFn0(p0 uint64, p1 bool) uint32 {
    var v7 int32 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 uint = 110
            var v11 int64 = 1
            _ = fi9
            _ = v10
            _ = v11
        }
        v7 += int32(fi8)
        var v12 int8 = 127
        _ = fi8
        _ = v12
    }
    v13 := []int{1, 32767, -242}
    var v14 int64 = 42
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            if (!(p1 && p1)) {
                p1 = (1 == 0)
                var v17 uint32 = 2147483647
                v18 := GenS1{F0: (v7 * (v7 / ((2147483647) | 1))), F1: p0}
                _ = v17
                _ = v18
            } else {
                if (!(116 < 2)) {
                    v14 = 42
                } else {
                    p1 = (fi15 >= fi16)
                }
                if len(v13) > 0 {
                    v13[0] = (fi15 ^ (fi15 + fi15))
                }
            }
            fmt.Println(p1, p1)
            v13 = append(v13, fi16)
            _ = fi16
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v7
    return 864
}
func genFn1(p0 uint32) uint32 {
    v19 := []int{72, -904, 4294967295}
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int32 = 42
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint8 = 128
            if ((0.0 - 1e10) < (float64(1.5) + 1e10)) {
                switch fi20 {
                case -728:
                    var v24 bool = false
                    _ = v24
                case 0:
                    if len(v19) > 0 {
                        v19[0] = (fi20 >> ((956) & 63))
                    }
                    var v25 int = fi20
                    fmt.Println(fi22)
                    _ = v25
                }
                p0 += (^(p0))
                var v26 uint64 = 4294967295
                _ = v26
            } else {
                var v27 uint8 = v23
                if (-1 <= 42) {
                    v27 -= v23
                } else {
                    v28 := GenS1{F0: v21, F1: 0}
                    _ = len(v19)
                    _ = v28
                }
                fmt.Println(fi22, v21)
                if true {
                    var v29 uint8 = (v23 << ((255) & 63))
                    _ = v29
                }
                v27 = v27
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch "hello" {
    case "foo bar":
        var v30 int64 = 42
        _ = v30
    case "a":
        var v31 float64 = 3.14159
        _ = v31
    case "\n\t":
        v32 := GenS2{F0: 255, F1: -104}
        _ = v32
    default:
        for fi33 := 0; fi33 < 4; fi33++ {
            switch "a" {
            case "foo bar":
                var v34 uint8 = 42
                _ = v34
            case "a":
                var v35 uint64 = 42
                _ = v35
            case "":
                v36 := GenS3{F0: 9223372036854775807, F1: p0, F2: 65535}
                _ = v36
            default:
                switch (^(fi33)) {
                case 127:
                    _ = len(v19)
                case -570:
                    switch "hello" {
                    case "foo bar":
                        _ = len(v19)
                    case "\n\t":
                        var v37 uint8 = 114
                        _ = v37
                    case "":
                        v38 := GenS1{F0: 42, F1: 497}
                        _ = v38
                    }
                }
            }
            _ = fi33
        }
    }
    if (!(1 <= 42)) {
        v19 = append(v19, 8)
        if len(v19) > 0 {
            v19[0] = 2147483647
        }
    } else {
        p0 = 346
        p0 = p0
        fmt.Println(p0)
    }
    _ = p0
    _ = v19
    return (p0 << ((482) & 63))
}

func main() {
    var v1 float64 = 1.5
    var v2 uint64 = 918
    var v3 float32 = 0.1
    v4 := []int{22, 128}
    var v5 int8 = -45
    var v6 int = -914
    v6 = v6
    for fi39 := 0; fi39 < 11; fi39++ {
        fmt.Println(v2, v3)
        var v40 int = (v6 | (^(fi39)))
        for fi41 := 0; fi41 < 2; fi41++ {
            v1 -= (v1 * (0.0 - (v1 / v1)))
            fmt.Println(fi39, v2)
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    fmt.Println(v2, v1)
    for fi42 := 0; fi42 < 12; fi42++ {
        for fi43 := 0; fi43 < 5; fi43++ {
            switch -9223372036854775808 {
            case 32768:
                _ = len(v4)
            case 2147483647:
                var v44 int64 = -9223372036854775808
                _ = v44
            case -985:
                fmt.Println(v6)
            }
            _ = fi43
        }
        _ = fi42
    }
    v2 -= (^(v2))
    if false {
        v2 = v2
        for fi45 := 0; fi45 < 10; fi45++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v46 := GenS4{F0: false, F1: -9223372036854775808}
            var v47 uint32 = 0
            _ = fi45
            _ = v46
            _ = v47
        }
    } else {
        var v48 float32 = (v3 * v3)
        _ = v48
    }
    fmt.Println(v2, v2)
    if (2 <= (^(v2))) {
        for fi49 := 0; fi49 < 3; fi49++ {
            for fi50 := 0; fi50 < 14; fi50++ {
                v51 := GenS5{F0: 751, F1: v2, F2: 115}
                v1 = v1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi50
                _ = v51
            }
            _ = fi49
        }
        switch v6 {
        case 42:
            var v52 int64 = 9223372036854775807
            _ = v52
        case -2147483648:
            var v53 bool = (128 > 127)
            _ = v53
        case 2147483647:
            fmt.Println(v1, v1, v3)
        default:
            var v54 uint16 = 1
            _ = v54
        }
        v1 += 1.5
    }
    var v55 int32 = 2147483647
    switch int(v2) {
    case 9223372036854775807:
        for fi56 := 0; fi56 < 2; fi56++ {
            var v57 uint16 = 1
            switch ("hello" + (("" + "hello") + "foo bar")) {
            case "a":
                _ = len(v4)
            case "\n\t":
                v55 = (v55 | v55)
            case "":
                for fi58 := 0; fi58 < 15; fi58++ {
                    switch ("" + "a") {
                    case "\n\t":
                        if (!((v57 << ((42) & 63)) < (v57 ^ 255))) {
                            var v59 int32 = (-(v55))
                            _ = v59
                        } else {
                            v60 := GenS5{F0: 1, F1: v2, F2: 52}
                            _ = v60
                        }
                    default:
                        v61 := GenS6{F0: false}
                        _ = v61
                    }
                    _ = len(v4)
                    fmt.Println(fi56)
                    _ = fi58
                }
            default:
                if false {
                    for fi62 := 0; fi62 < 8; fi62++ {
                        v63 := GenS7{F0: v2, F1: v55}
                        _ = fi62
                        _ = v63
                    }
                } else {
                    fmt.Println(v3)
                }
            }
            _ = fi56
            _ = v57
        }
    case 32767:
        switch "hello" {
        case "hello":
            var v64 int8 = 42
            _ = v64
        case "":
            if true {
                var v65 uint32 = 738
                if true {
                    v4 = append(v4, v6)
                } else {
                    if false {
                        for fi66 := 0; fi66 < 10; fi66++ {
                            if (376 != uint(v65)) {
                                v2 -= v2
                                switch (-(fi66)) {
                                case -382:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case 4294967295:
                                    v67 := GenS1{F0: v55, F1: 0}
                                    _ = v67
                                }
                            } else {
                                var v68 uint = 255
                                if false {
                                    if (0 < 42) {
                                        v4 = append(v4, (-(fi66)))
                                        for fi69 := 0; fi69 < 20; fi69++ {
                                            v70 := GenS8{F0: int64(v68)}
                                            if len(v4) > 0 {
                                                v4[0] = v6
                                            }
                                            _ = fi69
                                            _ = v70
                                        }
                                    }
                                    if (("\n\t" + (("\n\t" + (("\n\t" + "\n\t") + "hello")) + "a")) >= "") {
                                        var v71 string = ("a" + ("" + (("hello" + "foo bar") + "")))
                                        _ = v71
                                    } else {
                                        fmt.Println(fi66, v6)
                                        for fi72 := 0; fi72 < 5; fi72++ {
                                            v5 -= v5
                                            v3 = (1e10 / v3)
                                            _ = fi72
                                        }
                                    }
                                }
                                if (1e10 < (0.0 - (v1 / v1))) {
                                    for fi73 := 0; fi73 < 7; fi73++ {
                                        v3 -= (((v3 - v3) / (3.14159 + v3)) + float32(1e10))
                                        _ = fi73
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = fi66
                                    }
                                } else {
                                    if true {
                                        v2 += 1000
                                        v65 = (v65 & 42)
                                    } else {
                                        var v74 int64 = -1
                                        var v75 int8 = 0
                                        _ = v74
                                        _ = v75
                                    }
                                    for fi76 := 0; fi76 < 13; fi76++ {
                                        if (42 == -928) {
                                            var v77 uint8 = 0
                                            fmt.Println(v55, fi76, v1)
                                            if len(v4) > 0 {
                                                v4[0] = -9223372036854775808
                                            }
                                            _ = v77
                                        } else {
                                            _ = len(v4)
                                            switch "\n\t" {
                                            case "a":
                                                var v78 int8 = (v5 >> ((18446744073709551615) & 63))
                                                _ = v78
                                            case "":
                                                var v79 uint32 = (^(v65))
                                                _ = v79
                                            case "\n\t":
                                                if (!false) {
                                                    v80 := GenS9{F0: 448, F1: (v2 * 988)}
                                                    _ = v80
                                                } else {
                                                    if true {
                                                        var v81 bool = (int16(fi66) > 32767)
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        if (-1 > -1) {
                                                            if true {
                                                                v5 = (^(v5))
                                                                var v82 bool = v81
                                                                if true {
                                                                    v83 := GenS10{F0: -234, F1: (v5 >= (v5 ^ (v5 | v5))), F2: uint64(v65)}
                                                                    _ = v83
                                                                }
                                                                _ = v82
                                                            } else {
                                                                var v84 bool = (v1 <= float64(0.0))
                                                                v4 = append(v4, v6)
                                                                v81 = v81
                                                                _ = v84
                                                            }
                                                            fmt.Println(v1)
                                                        } else {
                                                            v65 -= genFn1(uint32(v68))
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = 32767
                                                        }
                                                        v85 := GenS4{F0: v81, F1: 1}
                                                        fmt.Println(v68, v55, v6)
                                                        if (!false) {
                                                            v81 = (483 > -1)
                                                            if true {
                                                                v3 -= (3.14159 * v3)
                                                                switch "a" {
                                                                case "foo bar":
                                                                    _ = len(v4)
                                                                case "a":
                                                                    if v81 {
                                                                        switch (v6 + 32768) {
                                                                        case 44:
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                        case 4294967295:
                                                                            fmt.Println(v2)
                                                                        }
                                                                        v4 = append(v4, (fi76 & (-(fi76))))
                                                                        v86 := GenS5{F0: int16(v65), F1: 123, F2: 123}
                                                                        var v87 int8 = -1
                                                                        var v88 int32 = (v55 >> (((v68 | v68)) & 63))
                                                                        v1 = -2.25
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                case "\n\t":
                                                                    var v89 uint16 = 2
                                                                    _ = v89
                                                                }
                                                            } else {
                                                                fmt.Println(v65)
                                                                if true {
                                                                    switch (fi76 & -9223372036854775808) {
                                                                    case 0:
                                                                        _ = len(v4)
                                                                    case -704:
                                                                        v81 = v81
                                                                    case 4294967295:
                                                                        var v90 uint8 = 128
                                                                        _ = v90
                                                                    default:
                                                                        v6 = (v6 - v6)
                                                                    }
                                                                    v3 = v3
                                                                    var v91 int16 = -1
                                                                    var v92 uint32 = (v65 + v65)
                                                                    var v93 string = (("foo bar" + "hello") + "foo bar")
                                                                    switch fi76 {
                                                                    case 4294967295:
                                                                        var v94 bool = (32767 <= (v91 * int16(fi66)))
                                                                        _ = v94
                                                                    case 32767:
                                                                        v95 := GenS11{F0: (("\n\t" + "\n\t") + v93)}
                                                                        var v96 int16 = 1000
                                                                        _ = v95
                                                                        _ = v96
                                                                    case 2147483647:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = fi76
                                                                        }
                                                                    default:
                                                                        if ((v55 >> ((2) & 63)) == v55) {
                                                                            v93 = ((v93 + v93) + "")
                                                                            switch fi76 {
                                                                            case 785:
                                                                                v97 := GenS10{F0: 545, F1: v81, F2: (^(v2))}
                                                                                _ = v97
                                                                            default:
                                                                                _ = len(v4)
                                                                            }
                                                                        }
                                                                    }
                                                                    var v98 uint16 = 65535
                                                                    _ = v91
                                                                    _ = v92
                                                                    _ = v93
                                                                    _ = v98
                                                                }
                                                                v1 = v1
                                                            }
                                                        } else {
                                                            var v99 int8 = 42
                                                            _ = v99
                                                        }
                                                        _ = v81
                                                        _ = v85
                                                    } else {
                                                        switch "foo bar" {
                                                        case "":
                                                            if ((^(v55)) != int32(v5)) {
                                                                v65 -= (v65 | v65)
                                                                var v100 uint16 = 734
                                                                var v101 int16 = 0
                                                                _ = v100
                                                                _ = v101
                                                            } else {
                                                                var v102 int8 = v5
                                                                var v103 int8 = 42
                                                                var v104 float64 = (-2.25 * v1)
                                                                _ = v102
                                                                _ = v103
                                                                _ = v104
                                                            }
                                                        case "a":
                                                            v1 = (0.1 - ((0.1 / 1.0) + 1.5))
                                                        case "hello":
                                                            switch (((("hello" + (("foo bar" + "foo bar") + ("" + ("\n\t" + "")))) + "foo bar") + "hello") + "\n\t") {
                                                            case "hello":
                                                                fmt.Println(fi66)
                                                            case "foo bar":
                                                                var v105 int = (v6 << ((2) & 63))
                                                                _ = v105
                                                            case "\n\t":
                                                                if true {
                                                                    var v106 int64 = int64(fi66)
                                                                    var v107 string = "foo bar"
                                                                    if len(v4) > 0 {
                                                                        v4[0] = -1
                                                                    }
                                                                    _ = v106
                                                                    _ = v107
                                                                } else {
                                                                    if (856 == uint64(v68)) {
                                                                        if true {
                                                                            fmt.Println(v1, fi76, v2)
                                                                            v6 -= (fi76 - fi66)
                                                                        }
                                                                        var v108 uint32 = uint32(v2)
                                                                        _ = v108
                                                                    }
                                                                    if (((v6 == (v6 ^ (v6 >> (((v68 / ((42) | 1))) & 63)))) && true) || false) {
                                                                        var v109 uint8 = 45
                                                                        var v110 int32 = (v55 | -642)
                                                                        _ = v109
                                                                        _ = v110
                                                                    } else {
                                                                        v65 = genFn1((v65 * (v65 >> ((2) & 63))))
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v3)
                                                            var v111 bool = true
                                                            _ = v111
                                                        default:
                                                            if false {
                                                                _ = len(v4)
                                                                var v112 float64 = 0.0
                                                                _ = v112
                                                            } else {
                                                                var v113 int = -601
                                                                _ = v113
                                                            }
                                                        }
                                                        if false {
                                                            var v114 uint16 = 42
                                                            v4 = append(v4, (fi66 ^ (v6 >> ((680) & 63))))
                                                            _ = v114
                                                        }
                                                    }
                                                }
                                            default:
                                                v6 = 127
                                            }
                                            var v115 string = "hello"
                                            v2 += 0
                                            fmt.Println(v65, v1, v115)
                                            _ = v115
                                        }
                                        _ = fi76
                                    }
                                    v55 += (-(v55))
                                }
                                _ = v68
                            }
                            v116 := GenS12{F0: 78, F1: "hello", F2: "foo bar"}
                            _ = fi66
                            _ = v116
                        }
                        var v117 string = "\n\t"
                        _ = v117
                    }
                    if (v3 > v3) {
                        _ = len(v4)
                        v3 += 1e10
                        var v118 int8 = 42
                        _ = v118
                    }
                }
                _ = v65
            }
        case "\n\t":
            var v119 uint8 = 127
            _ = v119
        default:
            v3 -= float32((v1 * v1))
        }
    case -394:
        v55 -= (-(v55))
    }
    v4 = append(v4, v6)
    v2 = (v2 * 0)
    var v120 bool = true
    if v120 {
        v55 = (v55 * (v55 / ((v55) | 1)))
        var v121 int16 = 199
        fmt.Println(v2, v55)
        _ = v121
    } else {
        fmt.Println(v55, v120)
    }
    if (((v6 * -2147483648) >= v6) && v120) {
        for fi122 := 0; fi122 < 16; fi122++ {
            var v123 uint8 = 2
            var v124 int32 = (v55 | v55)
            _ = fi122
            _ = v123
            _ = v124
        }
        for fi125 := 0; fi125 < 17; fi125++ {
            fmt.Println(v5)
            if (!(v120 && (v120 || (((v1 - float64((1.5 / 0.1))) + 0.1) != v1)))) {
                v2 += (v2 >> ((255) & 63))
                v2 += (v2 * v2)
            } else {
                var v126 uint32 = 65535
                v3 += v3
                _ = v126
            }
            _ = fi125
        }
    }
    v127 := GenS6{F0: v120}
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v120
    _ = v127
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
