package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 bool }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 string; F1 int8; F2 int64 }
type GenS8 struct { F0 uint64; F1 int64; F2 int64 }
type GenS9 struct { F0 uint8; F1 int16; F2 int }
type GenS10 struct { F0 int16; F1 string; F2 string }
type GenS11 struct { F0 int16 }

func main() {
    var v1 uint64 = 980
    v2 := make(map[uint32]uint32)
    var v3 float64 = (float64(3.14159) * 0.1)
    var v4 uint64 = 216
    v5 := []int{9223372036854775807, 0, 0}
    var v6 uint8 = 42
    switch -9223372036854775808 {
    case 30:
        var v7 uint = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            switch fi8 {
            case 412:
                var v10 uint8 = v6
                _ = v10
            case 1:
                v5 = append(v5, fi9)
                v11 := GenS1{F0: 57, F1: 1, F2: -1}
                switch (fi9 % ((32767) | 1)) {
                case -1:
                    if false {
                        v2[1] = 1
                        v2[4294967295] = 398
                    } else {
                        fmt.Println(v6, v4, v1)
                        v6 += v6
                    }
                case 285:
                    var v12 uint32 = 4294967295
                    _ = v12
                }
                if (-2.25 >= 1e10) {
                    var v13 int32 = 0
                    fmt.Println(v1, fi8)
                    var v14 uint32 = 0
                    v4 += 4294967295
                    _ = v13
                    _ = v14
                } else {
                    fmt.Println(fi9, v1)
                    v15 := v2[1]
                    _ = v15
                }
                _ = v11
            case 127:
                v16 := GenS1{F0: 10, F1: 1, F2: -29}
                _ = v16
            default:
                _ = len(v5)
            }
            var v17 float64 = v3
            _ = fi9
            _ = v17
        }
        _ = fi8
    }
    if false {
        for fi18 := 0; fi18 < 2; fi18++ {
            if true {
                v4 = (v4 - v4)
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi18)
                    case "a":
                        if (-9223372036854775808 == 9223372036854775807) {
                            if len(v5) > 0 {
                                v5[0] = (fi19 | (fi18 >> ((18446744073709551615) & 63)))
                            }
                            if false {
                                v20 := GenS1{F0: 71, F1: -14, F2: 22}
                                _ = v20
                            }
                        }
                    }
                    var v21 int8 = 20
                    _ = fi19
                    _ = v21
                }
                if true {
                    if (2 == 65535) {
                        for fi22 := 0; fi22 < 17; fi22++ {
                            var v23 float32 = 0.0
                            _ = fi22
                            _ = v23
                        }
                        v2[uint32(v6)] = 2147483647
                        var v24 float32 = 3.14159
                        _ = v24
                    }
                    var v25 uint16 = 65535
                    v25 -= v25
                    _ = v25
                }
            } else {
                v2[519] = 2
                switch int(v6) {
                case 32767:
                    var v26 uint = 816
                    _ = v26
                default:
                    v27 := GenS2{F0: 337, F1: "\n\t", F2: true}
                    _ = v27
                }
            }
            _ = fi18
        }
        var v28 int8 = 42
        if (true && (2 < 42)) {
            if (2 >= 2) {
                v29 := GenS1{F0: v28, F1: 100, F2: (v28 + 1)}
                _ = v29
            }
        } else {
            v1 += (v1 << ((0) & 63))
            v1 -= 374
        }
        _ = v28
    }
    var v30 uint = 1
    var v31 uint32 = 911
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if ((v3 + (0.0 / v3)) != float64(0.0)) {
        _ = len(v2)
        v32 := GenS1{F0: 100, F1: -1, F2: 42}
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 float32 = 0.1
            if false {
                if false {
                    var v35 int = fi33
                    v30 = (v30 & v30)
                    _ = v35
                } else {
                    if false {
                        v3 = (v3 - 1.5)
                        var v36 uint = v30
                        _ = v36
                    } else {
                        var v37 int8 = 100
                        var v38 bool = ((v31 * (v31 - (v31 & (^(v31))))) < 42)
                        _ = v37
                        _ = v38
                    }
                }
                var v39 uint8 = (v6 << (((v30 % (((v30 / ((v30) | 1))) | 1))) & 63))
                v39 = (v39 - v6)
                _ = v39
            } else {
                switch 127 {
                case 127:
                    if false {
                        v31 = (v31 % ((v31) | 1))
                    } else {
                        if false {
                            var v40 uint8 = v6
                            v30 -= 0
                            _ = v40
                        } else {
                            v34 = v34
                            if ((fi33 & int(v30)) <= (fi33 | -2147483648)) {
                                if (!true) {
                                    var v41 uint = (v30 % (((^(v30))) | 1))
                                    _ = v41
                                } else {
                                    _ = len(v5)
                                }
                                v30 = v30
                                v31 -= (v31 + (v31 - v31))
                            } else {
                                _ = len(v2)
                                if (v3 < 0.0) {
                                    v34 = (1e10 * (v34 / (1e10 * v34)))
                                    for fi42 := 0; fi42 < 2; fi42++ {
                                        if (false || ((^(v6)) != (^(v6)))) {
                                            if (!(65535 < (v30 % ((v30) | 1)))) {
                                                delete(v2, (^(v31)))
                                                fmt.Println(v31)
                                                v43 := GenS2{F0: v30, F1: ("foo bar" + ""), F2: ((543 > 0) && false)}
                                                _ = v43
                                            }
                                            fmt.Println(v30, v34, fi42)
                                            v44 := GenS1{F0: -128, F1: -1, F2: 0}
                                            var v45 int8 = 100
                                            v46 := GenS3{F0: (-(fi42)), F1: 1, F2: 42}
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        v5 = append(v5, (-(fi33)))
                                        var v47 float64 = 0.0
                                        _ = fi42
                                        _ = v47
                                    }
                                } else {
                                    switch (("hello" + (("a" + "foo bar") + "hello")) + "hello") {
                                    case "a":
                                        for fi48 := 0; fi48 < 18; fi48++ {
                                            var v49 int64 = -9223372036854775808
                                            var v50 int16 = int16(v1)
                                            v5 = append(v5, (fi33 % ((128) | 1)))
                                            if true {
                                                fmt.Println(v30)
                                                switch fi33 {
                                                case 9223372036854775807:
                                                    v51 := GenS4{F0: 128}
                                                    _ = v51
                                                default:
                                                    _ = len(v5)
                                                }
                                            } else {
                                                v52 := GenS2{F0: (^(v30)), F1: "hello", F2: (false && true)}
                                                _ = v52
                                            }
                                            _ = fi48
                                            _ = v49
                                            _ = v50
                                        }
                                    case "hello":
                                        var v53 bool = (!(549 < 42))
                                        _ = v53
                                    case "":
                                        var v54 uint8 = (^(v6))
                                        _ = v54
                                    default:
                                        var v55 uint32 = 233
                                        _ = v55
                                    }
                                }
                                v56 := GenS3{F0: (fi33 ^ (^(fi33))), F1: v1, F2: 32767}
                                _ = v56
                            }
                        }
                        if false {
                            v57 := GenS2{F0: (v30 >> (((v30 & 919)) & 63)), F1: "hello", F2: (true && true)}
                            switch ("hello" + "") {
                            case "hello":
                                switch fi33 {
                                case 128:
                                    var v58 int64 = 42
                                    _ = v58
                                case -449:
                                    v59 := GenS5{F0: 42}
                                    _ = v59
                                }
                            case "foo bar":
                                if true {
                                    switch (^(fi33)) {
                                    case -470:
                                        switch "\n\t" {
                                        case "hello":
                                            if (v6 == (v6 >> ((v30) & 63))) {
                                                var v60 int = int(v31)
                                                v61 := v2[42]
                                                var v62 int8 = 42
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            }
                                        }
                                    case 32767:
                                        if (42 <= 0) {
                                            fmt.Println(v4, v4)
                                            switch fi33 {
                                            case 9223372036854775807:
                                                v2[(v31 / (((v31 >> ((2) & 63))) | 1))] = 121
                                            case 0:
                                                for fi63 := 0; fi63 < 1; fi63++ {
                                                    v64 := GenS1{F0: 42, F1: 0, F2: 1}
                                                    v5 = append(v5, (fi63 ^ fi63))
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            case -2147483648:
                                                var v65 uint8 = 255
                                                _ = v65
                                            default:
                                                var v66 uint = 254
                                                _ = v66
                                            }
                                            var v67 string = "a"
                                            _ = v67
                                        }
                                    default:
                                        var v68 int8 = -4
                                        _ = v68
                                    }
                                }
                            case "\n\t":
                                v2[v31] = 723
                            default:
                                v69 := GenS2{F0: 4294967295, F1: ("foo bar" + ""), F2: (!(uint(v31) > (v30 + v30)))}
                                _ = v69
                            }
                            if false {
                                for fi70 := 0; fi70 < 15; fi70++ {
                                    var v71 int8 = -128
                                    if false {
                                        switch "\n\t" {
                                        case "hello":
                                            if (v6 > (v6 / (((^(v6))) | 1))) {
                                                v3 += -2.25
                                            }
                                        }
                                        v34 -= (v34 * -2.25)
                                        v72 := GenS6{F0: v1, F1: -9223372036854775808}
                                        _ = v72
                                    }
                                    v34 = v34
                                    _ = fi70
                                    _ = v71
                                }
                            }
                            _ = v57
                        }
                    }
                case 32768:
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        if ((false && false) && ((!true) || (((v31 * (^(v31))) <= (v31 >> ((v30) & 63))) && (1e10 <= float32((v3 / v3)))))) {
                            for fi73 := 0; fi73 < 7; fi73++ {
                                var v74 int8 = 127
                                var v75 uint32 = (v31 ^ (^(v31)))
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                        }
                    }
                }
                v3 = v3
                for fi76 := 0; fi76 < 15; fi76++ {
                    var v77 uint = v30
                    v78 := GenS4{F0: (v6 - v6)}
                    if true {
                        var v79 uint8 = 53
                        var v80 int16 = 451
                        v81 := GenS5{F0: v1}
                        var v82 float64 = v3
                        if (((v34 / v34) / -2.25) == v34) {
                            if true {
                                v83 := GenS7{F0: "hello", F1: -32, F2: 9223372036854775807}
                                if false {
                                    v84 := GenS2{F0: (v77 / ((v77) | 1)), F1: "foo bar", F2: ((0 >= 749) && (-42 >= -1))}
                                    fmt.Println(v34, fi33, v30)
                                    if (-793 != 9223372036854775807) {
                                        v85 := GenS3{F0: (fi33 / ((-882) | 1)), F1: v4, F2: uint16(v77)}
                                        v77 += 255
                                        _ = v85
                                    }
                                    _ = v84
                                } else {
                                    var v86 uint32 = (v31 + (v31 * v31))
                                    _ = v86
                                }
                                _ = v83
                            }
                            v87 := GenS8{F0: 42, F1: 0, F2: -81}
                            v88 := GenS9{F0: uint8(v31), F1: v80, F2: 9223372036854775807}
                            _ = v87
                            _ = v88
                        } else {
                            v4 -= 18446744073709551615
                            if false {
                                if true {
                                    var v89 int = (fi76 << (((v30 | (v30 & 42))) & 63))
                                    if (false || (v31 >= v31)) {
                                        fmt.Println(v80)
                                        v90 := GenS9{F0: (v79 * (v79 | v79)), F1: v80, F2: fi33}
                                        _ = v90
                                    }
                                    fmt.Println(v4)
                                    _ = v89
                                } else {
                                    if (int8(v80) > 42) {
                                        v91 := GenS1{F0: 127, F1: -1, F2: int8(v30)}
                                        var v92 int64 = -1
                                        var v93 uint8 = 127
                                        _ = v91
                                        _ = v92
                                        _ = v93
                                    }
                                    v80 -= (v80 ^ (v80 << (((v77 - v77)) & 63)))
                                    if (!(!false)) {
                                        v94 := v2[(v31 * 1)]
                                        v1 = v1
                                        switch "\n\t" {
                                        case "foo bar":
                                            if true {
                                                var v95 int16 = (v80 << ((v30) & 63))
                                                if true {
                                                    var v96 bool = (v1 <= uint64(v6))
                                                    v97 := GenS10{F0: v80, F1: ("" + "foo bar"), F2: "a"}
                                                    _ = v96
                                                    _ = v97
                                                }
                                                var v98 float32 = v34
                                                _ = v95
                                                _ = v98
                                            } else {
                                                v3 = v82
                                            }
                                        case "hello":
                                            v99 := GenS7{F0: "\n\t", F1: -12, F2: 9223372036854775807}
                                            _ = v99
                                        case "":
                                            if ((v80 >> ((v30) & 63)) > (-(v80))) {
                                                var v100 string = "hello"
                                                var v101 float32 = (((-2.25 - v34) - (v34 + (((float32(-2.25) + (v34 * 0.1)) - float32((((v3 - (float64((0.1 + ((1e10 * v34) / (v34 + -2.25)))) / (v82 / (v82 - (v82 / v3))))) + v3) + (((float64((0.0 / v34)) / ((0.0 / 0.1) * v3)) / v82) * ((v3 / v3) * (0.0 + v82)))))) + 1.5))) - 1e10)
                                                _ = v100
                                                _ = v101
                                            }
                                            if ((fi33 - -2147483648) != (fi33 << ((481) & 63))) {
                                                if len(v5) > 0 {
                                                    v5[0] = fi76
                                                }
                                                v102 := GenS10{F0: v80, F1: "a", F2: "hello"}
                                                var v103 int = (-(fi76))
                                                _ = v102
                                                _ = v103
                                            }
                                            v1 += 0
                                        }
                                        _ = v94
                                    } else {
                                        v104 := GenS10{F0: v80, F1: "\n\t", F2: "hello"}
                                        _ = v104
                                    }
                                }
                                v2[v31] = v31
                                switch "foo bar" {
                                case "foo bar":
                                    v4 = (v1 - (v4 + (v1 ^ (^(v4)))))
                                case "":
                                    var v105 float32 = float32(float64(3.14159))
                                    _ = v105
                                default:
                                    if true {
                                        fmt.Println(v3)
                                        var v106 float64 = (v3 / v3)
                                        var v107 uint32 = 4294967295
                                        _ = v106
                                        _ = v107
                                    }
                                }
                            } else {
                                fmt.Println(v3)
                                var v108 int16 = v80
                                var v109 float64 = v3
                                _ = v108
                                _ = v109
                            }
                        }
                        _ = v79
                        _ = v80
                        _ = v81
                        _ = v82
                    }
                    _ = fi76
                    _ = v77
                    _ = v78
                }
            }
            v34 += ((1.5 * 1e10) + (0.1 / 1.0))
            _ = fi33
            _ = v34
        }
        _ = v32
    } else {
        if (((uint8(v30) <= (^(v6))) && (true && (false && false))) && (!(!(51 <= 65535)))) {
            v110 := GenS8{F0: (v1 & v1), F1: 694, F2: -369}
            _ = v110
        }
        v111 := GenS4{F0: v6}
        _ = v111
    }
    v112 := GenS11{F0: -831}
    for fi113 := 0; fi113 < 15; fi113++ {
        var v114 uint64 = 284
        if false {
            var v115 int32 = 2147483647
            _ = v115
        } else {
            for fi116 := 0; fi116 < 2; fi116++ {
                var v117 uint8 = v6
                v3 -= v3
                _ = fi116
                _ = v117
            }
            var v118 uint8 = 255
            _ = v118
        }
        _ = fi113
        _ = v114
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
