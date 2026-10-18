package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint16 }
type GenS10 struct { F0 string; F1 int; F2 uint }
type GenS11 struct { F0 uint8; F1 int }
type GenS12 struct { F0 uint8; F1 int; F2 int16 }
func genFn0() uint32 {
    if (true || true) {
        v7 := []int{}
        var v8 uint64 = 395
        _ = v7
        _ = v8
    }
    var v9 uint8 = 2
    if true {
        v10 := []int{int(v9), 636}
        if true {
            v9 = v9
            v11 := GenS1{F0: (^(v9)), F1: 0, F2: 205}
            _ = v11
        } else {
            fmt.Println(v9)
            if true {
                if true {
                    var v12 int8 = 42
                    _ = v12
                } else {
                    var v13 uint16 = 677
                    switch "\n\t" {
                    case "foo bar":
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                    case "":
                        var v14 string = "foo bar"
                        _ = v14
                    }
                    _ = v13
                }
                fmt.Println(v9)
            } else {
                for fi15 := 0; fi15 < 11; fi15++ {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        v17 := GenS1{F0: (^(v9)), F1: -1, F2: int64(fi16)}
                        _ = fi16
                        _ = v17
                    }
                    v9 -= 1
                    _ = fi15
                }
                v9 = v9
                v18 := GenS1{F0: v9, F1: -840, F2: 42}
                _ = v18
            }
        }
        _ = v10
    } else {
        for fi19 := 0; fi19 < 8; fi19++ {
            switch ("hello" + "") {
            case "hello":
                fmt.Println(v9)
            case "\n\t":
                for fi20 := 0; fi20 < 15; fi20++ {
                    v9 = 93
                    if (true && (!(848 < 0))) {
                        var v21 int16 = -1
                        _ = v21
                    } else {
                        switch "a" {
                        case "foo bar":
                            if true {
                                v22 := GenS1{F0: (^(v9)), F1: 191, F2: int64(v9)}
                                v23 := []int{}
                                if false {
                                    v24 := GenS1{F0: v9, F1: 0, F2: -589}
                                    _ = v24
                                } else {
                                    var v25 bool = true
                                    _ = v25
                                }
                                v26 := GenS2{F0: 2, F1: -9223372036854775808, F2: -2147483648}
                                _ = v22
                                _ = v23
                                _ = v26
                            } else {
                                if (true && false) {
                                    v9 -= v9
                                    var v27 int = 127
                                    _ = v27
                                } else {
                                    if (1e10 != 1e10) {
                                        v9 -= (v9 / ((0) | 1))
                                        var v28 string = ("\n\t" + ("" + "hello"))
                                        if (((float64(1e10) == float64(1e10)) || true) && (127 < 44)) {
                                            v28 = v28
                                            var v29 int32 = -2147483648
                                            _ = v29
                                        }
                                        _ = v28
                                    } else {
                                        v9 = v9
                                        var v30 float32 = 1e10
                                        _ = v30
                                    }
                                }
                                v9 = v9
                                v9 -= v9
                            }
                        default:
                            v31 := []int{2147483647, fi20, (fi20 % ((4294967295) | 1)), fi20}
                            _ = v31
                        }
                        v32 := []int{(-(fi20)), 0, (fi19 / ((fi19) | 1))}
                        _ = v32
                    }
                    var v33 bool = (!true)
                    _ = fi20
                    _ = v33
                }
            case "foo bar":
                fmt.Println(v9, v9)
            }
            _ = fi19
        }
        var v34 uint32 = 2
        switch 4294967295 {
        case -982:
            v35 := []int{32767, 9223372036854775807, -9223372036854775808, 814}
            _ = v35
        case 4294967295:
            var v36 uint16 = 2
            _ = v36
        case 48:
            fmt.Println(v34, v34)
        default:
            switch "\n\t" {
            case "a":
                if (975 == 2) {
                    for fi37 := 0; fi37 < 18; fi37++ {
                        switch (fi37 / ((128) | 1)) {
                        case 0:
                            v34 += v34
                        default:
                            v34 = v34
                        }
                        var v38 string = "hello"
                        if false {
                            switch (-(fi37)) {
                            case 2147483647:
                                var v39 uint = 1
                                _ = v39
                            case 1:
                                var v40 uint = 42
                                _ = v40
                            }
                            var v41 string = v38
                            _ = v41
                        }
                        _ = fi37
                        _ = v38
                    }
                    v9 -= 0
                    v9 += uint8(v34)
                }
            case "":
                v34 -= (^(v34))
            }
        }
        _ = v34
    }
    if true {
        var v42 uint16 = 11
        switch "a" {
        case "a":
            if ((false || (((0 <= 11) && (32767 < -148)) && (1 > 283))) || true) {
                v43 := GenS1{F0: v9, F1: -251, F2: -86}
                v44 := GenS3{F0: 4294967295}
                v9 += (^(v9))
                _ = v43
                _ = v44
            }
        case "":
            v45 := GenS4{F0: (("hello" + "hello") + "\n\t")}
            _ = v45
        case "foo bar":
            var v46 float64 = 1.5
            _ = v46
        default:
            for fi47 := 0; fi47 < 20; fi47++ {
                switch 42 {
                case 2147483647:
                    v9 = v9
                }
                switch "foo bar" {
                case "a":
                    v42 = 32767
                default:
                    if (!(42 < 18446744073709551615)) {
                        switch (-(fi47)) {
                        case 0:
                            var v48 uint32 = 2
                            _ = v48
                        }
                    }
                }
                if (fi47 == -1) {
                    for fi49 := 0; fi49 < 20; fi49++ {
                        v50 := GenS1{F0: v9, F1: 2147483647, F2: -1}
                        v51 := GenS3{F0: 109}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                } else {
                    var v52 bool = (3.14159 <= 3.14159)
                    switch -57 {
                    case 9223372036854775807:
                        switch "a" {
                        case "a":
                            v53 := []int{fi47}
                            _ = v53
                        default:
                            v54 := []int{}
                            _ = v54
                        }
                    }
                    _ = v52
                }
                _ = fi47
            }
        }
        _ = v42
    } else {
        v55 := []int{42, 0, 42, 9223372036854775807}
        var v56 int = 761
        _ = v55
        _ = v56
    }
    fmt.Println(v9)
    _ = v9
    return 42
}

func main() {
    var v1 int = 457
    var v2 int8 = 42
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 int32 = 0
    var v6 uint16 = 1
    if false {
        var v57 int32 = (^(v5))
        if false {
            fmt.Println(v5, v5, v5)
            switch "\n\t" {
            case "\n\t":
                if true {
                    if (true || false) {
                        fmt.Println(v2, v2, v2)
                    } else {
                        v57 = (v5 >> ((480) & 63))
                        v2 = (v2 | v2)
                        v58 := GenS5{F0: (v6 & 2), F1: 1}
                        _ = v58
                    }
                    v2 = v2
                }
            case "foo bar":
                var v59 uint64 = 0
                _ = v59
            case "a":
                for fi60 := 0; fi60 < 3; fi60++ {
                    v61 := []int{fi60, (fi60 ^ 1), int(v2), fi60}
                    if len(v61) > 0 {
                        v61[0] = fi60
                    }
                    if len(v61) > 0 {
                        v61[0] = fi60
                    }
                    _ = fi60
                    _ = v61
                }
            default:
                fmt.Println(v5, v57)
            }
            var v62 uint16 = (v6 ^ (v6 >> ((241) & 63)))
            _ = v62
        } else {
            var v63 int8 = 1
            _ = v63
        }
        for fi64 := 0; fi64 < 2; fi64++ {
            switch v1 {
            case 32768:
                var v65 bool = (9223372036854775807 != (-(fi64)))
                _ = v65
            default:
                v66 := GenS6{F0: ""}
                _ = v66
            }
            switch v1 {
            case 678:
                if (748 != 1) {
                    v67 := []int{v1, fi64}
                    v57 = (v57 / ((v5) | 1))
                    v2 = -1
                    _ = v67
                } else {
                    var v68 int64 = 107
                    delete(v4, (fi64 ^ (v1 | 127)))
                    _ = v68
                }
            case 4294967295:
                var v69 float64 = 1.5
                _ = v69
            case 9223372036854775807:
                var v70 int32 = v57
                _ = v70
            default:
                if true {
                    v1 = -98
                    var v71 string = "a"
                    for fi72 := 0; fi72 < 10; fi72++ {
                        v73 := []int{(-(fi72)), (-(v1)), (v1 ^ fi64)}
                        if (((981 < 0) || true) || false) {
                            if ((!false) && (v1 < (fi72 + 2147483647))) {
                                if len(v73) > 0 {
                                    v73[0] = v1
                                }
                            } else {
                                if false {
                                    v2 += (v2 - (v2 << ((323) & 63)))
                                    if (222 < -32768) {
                                        if (!(1 != 4294967295)) {
                                            v57 = (v5 + int32(fi64))
                                        }
                                    }
                                } else {
                                    _ = len(v73)
                                    if false {
                                        fmt.Println(fi72, v71)
                                    }
                                }
                            }
                        } else {
                            v57 -= (v5 >> ((2) & 63))
                            v2 -= v2
                        }
                        v57 = (v57 * v57)
                        if true {
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    v5 -= v5
                                case "foo bar":
                                    _ = len(v3)
                                default:
                                    var v74 string = (v71 + "\n\t")
                                    _ = v74
                                }
                                fmt.Println(v2, fi72)
                                v75 := GenS7{F0: (false && (false || false))}
                                _ = v75
                            } else {
                                if false {
                                    var v76 uint = 18446744073709551615
                                    v77 := v4[(fi64 + (fi64 % ((32768) | 1)))]
                                    _ = v76
                                    _ = v77
                                } else {
                                    if ((v71 <= (v71 + v71)) || (127 != 8)) {
                                        v71 = ((v71 + "") + "")
                                        var v78 uint8 = 2
                                        v57 += v57
                                        var v79 uint64 = 42
                                        _ = v78
                                        _ = v79
                                    } else {
                                        v80 := GenS8{F0: v6}
                                        var v81 uint64 = 672
                                        if (1 <= int16(v6)) {
                                            switch fi72 {
                                            case 42:
                                                _ = len(v4)
                                            default:
                                                v2 -= 100
                                            }
                                        } else {
                                            switch ("\n\t" + (v71 + "foo bar")) {
                                            case "":
                                                v6 = (v6 >> ((42) & 63))
                                            case "\n\t":
                                                v4[(fi64 / (((fi72 / (((fi64 * (fi64 + 2147483647))) | 1))) | 1))] = "a"
                                            }
                                            if (true || (-32768 != -32768)) {
                                                var v82 float64 = (0.1 / (0.1 * (1e10 + float64((1.5 - -2.25)))))
                                                fmt.Println(fi72, v82, fi72)
                                                if false {
                                                    v6 -= uint16(v81)
                                                    v83 := GenS7{F0: ((v82 - v82) > float64(((3.14159 + (0.0 * 0.0)) - (0.1 - 1e10))))}
                                                    if (!false) {
                                                        fmt.Println(v57)
                                                        fmt.Println(v81, v71, fi64)
                                                        v84 := GenS5{F0: 0, F1: (^(v81))}
                                                        _ = v84
                                                    } else {
                                                        v85 := GenS9{F0: 4294967295, F1: (^(v6))}
                                                        if (!true) {
                                                            v86 := GenS6{F0: "hello"}
                                                            var v87 int32 = -1
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        _ = v85
                                                    }
                                                    _ = v83
                                                } else {
                                                    var v88 uint8 = 210
                                                    var v89 uint8 = (^(v88))
                                                    _ = v88
                                                    _ = v89
                                                }
                                                if len(v73) >= 2 {
                                                    v73 = v73[1:len(v73)]
                                                }
                                                _ = v82
                                            } else {
                                                switch v1 {
                                                case -9223372036854775808:
                                                    var v90 int16 = 0
                                                    _ = v90
                                                default:
                                                    var v91 uint = 177
                                                    _ = v91
                                                }
                                                var v92 int64 = 168
                                                _ = v92
                                            }
                                        }
                                        _ = v80
                                        _ = v81
                                    }
                                }
                                var v93 float32 = 0.1
                                _ = v93
                            }
                            v1 -= (fi72 ^ (fi64 / (((^(fi72))) | 1)))
                            v94 := v3[v1]
                            _ = v94
                        }
                        var v95 bool = ((32767 < fi72) && true)
                        _ = fi72
                        _ = v73
                        _ = v95
                    }
                    _ = v71
                } else {
                    var v96 uint = 0
                    _ = v96
                }
            }
            _ = fi64
        }
        _ = v57
    } else {
        v97 := GenS10{F0: ("\n\t" + (("foo bar" + ((("\n\t" + "foo bar") + "foo bar") + (("\n\t" + ("" + ("hello" + "foo bar"))) + "a"))) + "hello")), F1: (-(v1)), F2: 0}
        fmt.Println(v6)
        v98 := GenS11{F0: 203, F1: v1}
        _ = v97
        _ = v98
    }
    v99 := GenS11{F0: 2, F1: 531}
    v100 := GenS12{F0: 97, F1: v1, F2: -584}
    for fi101 := 0; fi101 < 8; fi101++ {
        switch (v1 % ((42) | 1)) {
        case -220:
            v102 := GenS1{F0: 128, F1: (^(v5)), F2: -9223372036854775808}
            _ = v102
        }
        v103 := []int{}
        for fi104 := 0; fi104 < 1; fi104++ {
            var v105 int16 = -223
            v106 := v3[1]
            _ = fi104
            _ = v105
            _ = v106
        }
        _ = fi101
        _ = v103
    }
    _ = len(v3)
    var v107 int = v1
    var v108 float32 = ((float32((0.1 - -2.25)) - ((0.0 / 3.14159) + 3.14159)) / float32(-2.25))
    v6 -= (v6 ^ 920)
    for fi109 := 0; fi109 < 11; fi109++ {
        for fi110 := 0; fi110 < 20; fi110++ {
            var v111 int8 = v2
            _ = len(v4)
            var v112 uint8 = 128
            delete(v4, 468)
            fmt.Println(v5, v1, v112)
            var v113 float32 = v108
            _ = fi110
            _ = v111
            _ = v112
            _ = v113
        }
        _ = fi109
    }
    v107 = int(v5)
    for fi114 := 0; fi114 < 12; fi114++ {
        v5 += v5
        switch (fi114 & (v107 * fi114)) {
        case -1:
            if true {
                v115 := []int{1, (v1 * fi114)}
                v116 := GenS2{F0: 887, F1: 262, F2: (v5 / ((v5) | 1))}
                _ = v115
                _ = v116
            } else {
                v6 = 0
                var v117 uint16 = v6
                var v118 int8 = (-(v2))
                _ = v117
                _ = v118
            }
        default:
            v6 = 0
        }
        fmt.Println(v107)
        _ = fi114
    }
    fmt.Println(v5, v5, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v100
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v99
}
