package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64; F1 int; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 bool; F1 int16; F2 uint32 }
type GenS11 struct { F0 int8; F1 uint64; F2 int64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 bool; F1 uint64; F2 int64 }
func genFn0(p0 int) bool {
    var v6 float64 = ((-2.25 - (0.1 + 1e10)) - 1e10)
    if true {
        fmt.Println(p0)
        v7 := GenS1{F0: 253, F1: 1}
        var v8 int8 = 85
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v6
    return true
}
func genFn1(p0 uint32, p1 int8, p2 string) bool {
    if true {
        v9 := []int{-2147483648, -9223372036854775808, -1, -1}
        p1 += p1
        _ = v9
    } else {
        v10 := GenS2{F0: -2147483648, F1: 2, F2: p1}
        v11 := GenS3{F0: -2147483648}
        var v12 int8 = -128
        _ = v10
        _ = v11
        _ = v12
    }
    if (42 >= 0) {
        switch p2 {
        case "foo bar":
            if (1 != 2) {
                v13 := GenS4{F0: -9223372036854775808}
                fmt.Println(p0, p1, p0)
                _ = v13
            }
        case "":
            var v14 int64 = 512
            _ = v14
        case "\n\t":
            v15 := GenS5{F0: 1, F1: -2147483648, F2: 42}
            _ = v15
        default:
            p2 = "a"
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                v18 := []int{}
                if (true || true) {
                    var v19 uint16 = 253
                    if (18446744073709551615 > 1) {
                        var v20 uint16 = (v19 / ((v19) | 1))
                        v18 = append(v18, fi16)
                        _ = v20
                    } else {
                        fmt.Println(v19, fi16, v19)
                    }
                    v21 := GenS6{F0: 0}
                    _ = v19
                    _ = v21
                }
                switch p2 {
                case "":
                    v22 := GenS4{F0: 9223372036854775807}
                    _ = v22
                case "hello":
                    if (42 <= 255) {
                        p1 = (p1 & -118)
                        var v23 uint = 441
                        _ = v23
                    } else {
                        if (false && (true && false)) {
                            p2 = ""
                            v24 := GenS5{F0: -1, F1: fi17, F2: 0}
                            p2 = p2
                            if ((42 < 2) || ((p0 | (p0 + p0)) <= 65535)) {
                                fmt.Println(fi17, p1)
                            }
                            if (38 < (p1 << ((65535) & 63))) {
                                p2 = ((p2 + ("" + p2)) + ((p2 + p2) + p2))
                                var v25 uint64 = 42
                                var v26 float64 = (0.1 - ((-2.25 * 0.0) * 0.0))
                                _ = v25
                                _ = v26
                            } else {
                                p2 = p2
                                _ = len(v18)
                                if (true || true) {
                                    switch fi16 {
                                    case 564:
                                        fmt.Println(fi16, fi17)
                                    case -9223372036854775808:
                                        var v27 int32 = 42
                                        _ = v27
                                    case 9223372036854775807:
                                        fmt.Println(fi17, fi16)
                                    default:
                                        v28 := GenS6{F0: -449}
                                        _ = v28
                                    }
                                    var v29 int16 = 1
                                    _ = v29
                                } else {
                                    v30 := GenS7{F0: 255}
                                    var v31 uint64 = 4294967295
                                    _ = v30
                                    _ = v31
                                }
                                if len(v18) > 0 {
                                    v18[0] = 45
                                }
                            }
                            _ = v24
                        } else {
                            var v32 float64 = (3.14159 * 1.5)
                            _ = v32
                        }
                        if false {
                            if true {
                                var v33 uint16 = 255
                                v34 := GenS8{F0: 42, F1: -820}
                                _ = v33
                                _ = v34
                            }
                            if true {
                                if (!(!true)) {
                                    p2 = ("a" + "hello")
                                    var v35 uint32 = p0
                                    _ = v35
                                } else {
                                    var v36 string = ("\n\t" + p2)
                                    var v37 int = fi16
                                    if false {
                                        p1 += -1
                                        v36 = ("\n\t" + v36)
                                    }
                                    _ = v36
                                    _ = v37
                                }
                                if false {
                                    var v38 string = ("a" + p2)
                                    v39 := GenS2{F0: 1, F1: 2, F2: (-(p1))}
                                    _ = v38
                                    _ = v39
                                }
                            } else {
                                p2 = p2
                                if (0 <= 128) {
                                    p0 = p0
                                    if false {
                                        var v40 uint16 = 588
                                        _ = v40
                                    }
                                } else {
                                    p1 += 0
                                    p1 -= (p1 / ((p1) | 1))
                                    if true {
                                        v18 = append(v18, 2147483647)
                                        if len(v18) > 0 {
                                            v18[0] = -726
                                        }
                                        p2 = "foo bar"
                                        var v41 int16 = 1000
                                        p2 = ((p2 + "foo bar") + "a")
                                        _ = v41
                                    }
                                }
                                var v42 int32 = 2147483647
                                var v43 uint = 255
                                fmt.Println(fi16, p1)
                                switch 234 {
                                case -659:
                                    if false {
                                        _ = len(v18)
                                        v18 = append(v18, 4294967295)
                                        v42 = v42
                                    } else {
                                        var v44 string = p2
                                        _ = v44
                                    }
                                case -1:
                                    if (!false) {
                                        p0 = (p0 & p0)
                                        var v45 int32 = (v42 << (((v43 / ((v43) | 1))) & 63))
                                        _ = v45
                                    } else {
                                        var v46 int32 = (v42 ^ v42)
                                        if (((-(fi16)) >= (fi17 >> (((v43 | 42)) & 63))) && ((521 > 42) && false)) {
                                            switch "hello" {
                                            case "hello":
                                                fmt.Println(v42, v42)
                                                if true {
                                                    var v47 bool = ((!(2 != 173)) || (!(255 != 32767)))
                                                    var v48 bool = (!(42 >= 127))
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                }
                                                if (0 >= 1) {
                                                    var v49 uint16 = 255
                                                    var v50 bool = true
                                                    _ = v49
                                                    _ = v50
                                                } else {
                                                    v42 += v42
                                                    var v51 float64 = (((0.1 * 1e10) * 0.0) + 0.1)
                                                    if (!true) {
                                                        v46 += 914
                                                        fmt.Println(v42)
                                                        v52 := GenS4{F0: -9223372036854775808}
                                                        _ = v52
                                                    } else {
                                                        if (p0 != p0) {
                                                            fmt.Println(v42)
                                                            if (((v51 - -2.25) - v51) >= v51) {
                                                                p1 = p1
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v53 int16 = 42
                                                                    _ = v53
                                                                case "hello":
                                                                    fmt.Println(v42, p0)
                                                                    var v54 int = 886
                                                                    _ = v54
                                                                }
                                                                if true {
                                                                    p0 = (p0 >> ((uint(p0)) & 63))
                                                                }
                                                            }
                                                            fmt.Println(p1)
                                                        } else {
                                                            fmt.Println(v46)
                                                            p2 = ((p2 + ("\n\t" + p2)) + (p2 + ("hello" + "")))
                                                            var v55 uint64 = 4294967295
                                                            _ = v55
                                                        }
                                                    }
                                                    _ = v51
                                                }
                                            case "":
                                                fmt.Println(fi16, p1, v43)
                                            }
                                        }
                                        var v56 int = fi16
                                        _ = v46
                                        _ = v56
                                    }
                                }
                                v42 -= (v42 << ((v43) & 63))
                                _ = v42
                                _ = v43
                            }
                            switch fi17 {
                            case 0:
                                p2 = p2
                            default:
                                var v57 float32 = -2.25
                                _ = v57
                            }
                        }
                        v58 := GenS5{F0: int64(p1), F1: fi16, F2: 0}
                        _ = v58
                    }
                case "\n\t":
                    p1 -= (p1 | p1)
                    p2 = p2
                }
                fmt.Println(fi16)
                _ = fi17
                _ = v18
            }
            fmt.Println(p0, p2)
            fmt.Println(p1, p1)
            _ = fi16
        }
    }
    var v59 int16 = -32768
    if false {
        v59 += -32768
    } else {
        var v60 uint8 = 31
        _ = v60
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v59
    return (!(true || (-2.25 == float64(1e10))))
}

func main() {
    var v1 int16 = 0
    var v2 float64 = 0.1
    var v3 int = 4294967295
    var v4 uint16 = 42
    v5 := make(map[uint32]string)
    if false {
        var v61 uint16 = 255
        _ = v61
    }
    switch (v3 ^ 1) {
    case 127:
        v62 := GenS9{F0: 0}
        _ = v62
    case 32767:
        fmt.Println(v2, v4, v1)
    case 32768:
        var v63 uint16 = v4
        _ = v63
    default:
        v64 := GenS10{F0: ((false && true) || true), F1: int16(v4), F2: 2147483647}
        _ = v64
    }
    var v65 string = "a"
    var v66 int64 = 1
    if (!false) {
        fmt.Println(v1, v65)
        v66 = (v66 / ((v66) | 1))
    } else {
        var v67 float32 = (3.14159 * 0.1)
        fmt.Println(v65, v67, v4)
        _ = v67
    }
    for fi68 := 0; fi68 < 17; fi68++ {
        v69 := GenS4{F0: (v66 << ((42) & 63))}
        if (127 == -128) {
            fmt.Println(v65, v3, v3)
            if (-82 <= 127) {
                var v70 float32 = 1.5
                _ = v70
            } else {
                if (v3 < (fi68 >> ((134) & 63))) {
                    var v71 uint = 1
                    switch (v3 >> (((v71 & (^(v71)))) & 63)) {
                    case -660:
                        v72 := v5[0]
                        _ = v72
                    case 458:
                        v2 = 1.5
                    case -1:
                        switch -195 {
                        case 0:
                            v73 := []int{(fi68 / (((fi68 ^ (fi68 * v3))) | 1)), v3}
                            _ = v73
                        case 2147483647:
                            fmt.Println(v71, fi68)
                        default:
                            v65 = ("foo bar" + "foo bar")
                        }
                    default:
                        v66 += v66
                    }
                    v74 := []int{int(v66), fi68, 42, 32768}
                    _ = v71
                    _ = v74
                }
                v3 += (-(fi68))
                var v75 int = v3
                _ = v75
            }
            if true {
                v76 := GenS5{F0: v66, F1: (fi68 ^ v3), F2: 42}
                _ = v76
            }
        }
        v77 := GenS3{F0: -180}
        _ = fi68
        _ = v69
        _ = v77
    }
    for fi78 := 0; fi78 < 9; fi78++ {
        var v79 uint = 4294967295
        _ = fi78
        _ = v79
    }
    v80 := GenS11{F0: 0, F1: 2, F2: (v66 << ((1) & 63))}
    var v81 uint16 = 1
    switch v65 {
    case "a":
        if false {
            var v82 uint = 4294967295
            var v83 float32 = 3.14159
            v84 := GenS1{F0: 494, F1: 1}
            _ = v82
            _ = v83
            _ = v84
        } else {
            v66 += (v66 % ((9223372036854775807) | 1))
        }
    case "foo bar":
        v85 := []int{v3}
        _ = v85
    default:
        switch ("" + v65) {
        case "\n\t":
            v86 := []int{(v3 * int(v81)), (v3 << ((627) & 63)), (-(v3)), 0}
            _ = v86
        case "":
            if false {
                v1 += (v1 & (^(v1)))
                fmt.Println(v81)
                for fi87 := 0; fi87 < 15; fi87++ {
                    v88 := GenS12{F0: (v81 % (((v4 + v81)) | 1))}
                    _ = fi87
                    _ = v88
                }
            }
        default:
            for fi89 := 0; fi89 < 6; fi89++ {
                v90 := []int{v3, v3}
                v66 = v66
                _ = fi89
                _ = v90
            }
        }
    }
    var v91 int64 = (v66 - v66)
    for fi92 := 0; fi92 < 16; fi92++ {
        v4 -= v81
        _ = fi92
    }
    for fi93 := 0; fi93 < 18; fi93++ {
        v1 += (v1 * v1)
        _ = fi93
    }
    delete(v5, 2147483647)
    v94 := GenS9{F0: 0}
    v91 -= (v91 << ((18446744073709551615) & 63))
    for fi95 := 0; fi95 < 13; fi95++ {
        if (uint8(v81) >= 194) {
            if (!(529 < 2147483647)) {
                var v96 int8 = 1
                switch v65 {
                case "":
                    switch ((v65 + v65) + v65) {
                    case "foo bar":
                        var v97 int = fi95
                        _ = v97
                    default:
                        if true {
                            v98 := GenS1{F0: 994, F1: 1}
                            v99 := GenS2{F0: -2147483648, F1: 1, F2: 42}
                            _ = v98
                            _ = v99
                        }
                    }
                case "hello":
                    var v100 bool = (!(true || (88 <= 84)))
                    _ = v100
                default:
                    fmt.Println(v91, v81, v2)
                }
                _ = v96
            } else {
                v101 := []int{v3, 32768, fi95, (fi95 | v3)}
                v81 -= v81
                switch -2147483648 {
                case 174:
                    var v102 uint16 = 255
                    _ = v102
                case 0:
                    switch -2147483648 {
                    case -1:
                        for fi103 := 0; fi103 < 16; fi103++ {
                            if len(v101) > 0 {
                                v101[0] = fi95
                            }
                            _ = fi103
                        }
                    }
                case 42:
                    switch (v3 - fi95) {
                    case -645:
                        for fi104 := 0; fi104 < 15; fi104++ {
                            var v105 uint8 = uint8(v81)
                            var v106 uint64 = 1
                            var v107 float64 = (v2 - v2)
                            _ = fi104
                            _ = v105
                            _ = v106
                            _ = v107
                        }
                    default:
                        v1 = -32768
                    }
                default:
                    v3 -= (v3 << ((18446744073709551615) & 63))
                }
                _ = v101
            }
            fmt.Println(v91)
        } else {
            for fi108 := 0; fi108 < 7; fi108++ {
                var v109 int = 2147483647
                _ = fi108
                _ = v109
            }
            switch (v3 / (((fi95 ^ v3)) | 1)) {
            case 1:
                fmt.Println(v91, fi95)
            case 9223372036854775807:
                v4 = (v4 - 607)
            case 2147483647:
                switch v65 {
                case "\n\t":
                    var v110 int8 = 42
                    _ = v110
                case "foo bar":
                    switch 9223372036854775807 {
                    case -1:
                        var v111 int64 = -342
                        _ = v111
                    case 709:
                        fmt.Println(v81, v81)
                    case 50:
                        for fi112 := 0; fi112 < 5; fi112++ {
                            v3 = fi112
                            v113 := GenS5{F0: (-(v91)), F1: int(v4), F2: 175}
                            v114 := GenS13{F0: true, F1: 635, F2: (v66 + (v91 | (v66 >> ((18446744073709551615) & 63))))}
                            var v115 uint8 = 71
                            v116 := []int{-2147483648, (fi112 << ((1) & 63)), (fi95 & v3), (fi112 << ((70) & 63))}
                            v117 := GenS2{F0: 42, F1: 4294967295, F2: -6}
                            var v118 uint32 = 2147483647
                            var v119 uint64 = 18446744073709551615
                            _ = fi112
                            _ = v113
                            _ = v114
                            _ = v115
                            _ = v116
                            _ = v117
                            _ = v118
                            _ = v119
                        }
                    }
                }
            default:
                v120 := GenS12{F0: 892}
                _ = v120
            }
        }
        v121 := []int{(fi95 / (((fi95 & fi95)) | 1))}
        delete(v5, 600)
        _ = fi95
        _ = v121
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v81:", v81)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v65
    _ = v66
    _ = v80
    _ = v81
    _ = v91
    _ = v94
}
