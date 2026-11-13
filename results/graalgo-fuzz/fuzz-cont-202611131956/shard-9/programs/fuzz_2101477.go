package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int }
type GenS2 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 int64; F1 int16; F2 int16 }
func genFn0() int8 {
    v5 := []int{42, -2147483648}
    var v6 int16 = 0
    var v7 uint8 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 2147483647)
    _ = v5
    _ = v6
    _ = v7
    return -128
}
func genFn1() int8 {
    if (false && true) {
        var v8 int16 = 42
        _ = v8
    } else {
        var v9 uint16 = 42
        switch ("hello" + ("foo bar" + "")) {
        case "foo bar":
            v10 := GenS1{F0: 42, F1: (1 > uint(v9)), F2: 32768}
            _ = v10
        case "":
            var v11 uint8 = 2
            _ = v11
        case "hello":
            v12 := []int{-31, -262, 128}
            _ = v12
        default:
            v13 := GenS1{F0: 42, F1: false, F2: 320}
            _ = v13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(fi14, v9)
            _ = fi14
        }
        _ = v9
    }
    var v15 uint64 = 42
    _ = v15
    return 100
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 948
    var v3 uint16 = 1
    v4 := make(map[string]bool)
    fmt.Println(v2, v3)
    v4["a"] = (int8(v1) > -95)
    for fi16 := 0; fi16 < 5; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if false {
                delete(v4, "a")
                v1 = 2147483647
                v1 += (v1 * v1)
            } else {
                _ = len(v4)
            }
            _ = fi17
        }
        switch "\n\t" {
        case "foo bar":
            if true {
                v18 := GenS1{F0: (v2 / ((v2) | 1)), F1: (203 >= 1), F2: 42}
                if true {
                    v19 := v4["foo bar"]
                    var v20 float64 = 3.14159
                    var v21 uint64 = 868
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(v1, v3)
                }
                _ = v18
            } else {
                var v22 uint16 = (v3 / ((v3) | 1))
                var v23 uint64 = v2
                _ = v22
                _ = v23
            }
        case "hello":
            v2 -= (^(v2))
        case "\n\t":
            var v24 bool = (3.14159 < float64(((0.1 - 1e10) - 1.5)))
            _ = v24
        default:
            v3 = (v3 & v3)
        }
        v25 := []int{9223372036854775807, -577, -1}
        _ = fi16
        _ = v25
    }
    if (-9223372036854775808 >= 1) {
        var v26 int16 = 436
        switch 1 {
        case -319:
            v1 -= v1
        case -2147483648:
            var v27 int32 = int32(v2)
            _ = v27
        case 128:
            for fi28 := 0; fi28 < 18; fi28++ {
                switch "hello" {
                case "a":
                    v1 = -141
                }
                var v29 uint64 = 0
                for fi30 := 0; fi30 < 10; fi30++ {
                    v2 = 18446744073709551615
                    _ = fi30
                }
                _ = fi28
                _ = v29
            }
        default:
            _ = len(v4)
        }
        fmt.Println(v26, v2, v26)
        _ = v26
    } else {
        v1 += -1
        v3 = 42
        v3 -= v3
    }
    fmt.Println(v1, v1)
    var v31 int = -2147483648
    if false {
        var v32 int64 = 9223372036854775807
        _ = v32
    }
    v33 := GenS1{F0: (v2 + v2), F1: false, F2: v31}
    v34 := GenS2{F0: uint16(v2), F1: (-(v1)), F2: 127}
    if (!((("" + "\n\t") + ("\n\t" + "hello")) != ("foo bar" + "foo bar"))) {
        for fi35 := 0; fi35 < 13; fi35++ {
            if false {
                v36 := GenS1{F0: v2, F1: ((v31 & (^(fi35))) != v31), F2: (fi35 % (((fi35 & fi35)) | 1))}
                if (!true) {
                    fmt.Println(v3)
                    v37 := GenS3{F0: -9223372036854775808, F1: 168, F2: ((^(v3)) <= (^(v3)))}
                    _ = v37
                } else {
                    fmt.Println(fi35, v3, fi35)
                    for fi38 := 0; fi38 < 18; fi38++ {
                        if true {
                            var v39 int32 = 1
                            if (true && (65535 < 42)) {
                                var v40 bool = false
                                var v41 uint64 = v2
                                _ = v40
                                _ = v41
                            }
                            _ = v39
                        }
                        v3 = 65535
                        switch 32768 {
                        case 4294967295:
                            var v42 uint64 = 18446744073709551615
                            _ = v42
                        case 0:
                            v43 := []int{}
                            _ = v43
                        case 967:
                            var v44 int32 = v1
                            _ = v44
                        default:
                            v45 := GenS2{F0: (v3 << ((362) & 63)), F1: (v1 % (((v1 << ((2) & 63))) | 1)), F2: 2}
                            var v46 bool = false
                            var v47 bool = (v3 < v3)
                            if (-1 < (v1 / (((v1 | -1)) | 1))) {
                                switch "\n\t" {
                                case "":
                                    if v47 {
                                        var v48 int64 = 9223372036854775807
                                        v1 -= v1
                                        switch "hello" {
                                        case "hello":
                                            fmt.Println(fi38)
                                        case "a":
                                            _ = len(v4)
                                        default:
                                            var v49 uint32 = 65535
                                            _ = v49
                                        }
                                        _ = v48
                                    }
                                case "hello":
                                    if (true && v46) {
                                        v50 := GenS1{F0: v2, F1: ((v47 && ((0.1 / 1.5) > (1.5 / ((((0.1 * 0.1) - 3.14159) - (1e10 / -2.25)) - 0.0)))) && (!((v1 ^ (v1 & (v1 ^ -551))) != (v1 + 0)))), F2: v31}
                                        var v51 float32 = 1e10
                                        _ = v50
                                        _ = v51
                                    } else {
                                        if (!v47) {
                                            v3 += (v3 * (^(v3)))
                                        } else {
                                            if v46 {
                                                var v52 int8 = 127
                                                _ = v52
                                            } else {
                                                v53 := GenS4{F0: "", F1: uint(v3), F2: v2}
                                                if ((0 == 339) || (!((1.5 + (0.0 * 0.1)) != float32(1e10)))) {
                                                    v54 := []int{fi38, fi38, (v31 << ((42) & 63))}
                                                    var v55 float32 = 1.5
                                                    fmt.Println(v55, fi35, v46)
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if false {
                                                    if (!true) {
                                                        if false {
                                                            if (false && (0 <= -97)) {
                                                                switch ("hello" + (("\n\t" + "foo bar") + "foo bar")) {
                                                                case "hello":
                                                                    if v46 {
                                                                        v1 += v1
                                                                    } else {
                                                                        switch (v31 >> ((908) & 63)) {
                                                                        case 32767:
                                                                            var v56 uint = 2
                                                                            _ = v56
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        fmt.Println(v47, fi38, fi35)
                                                                    }
                                                                case "":
                                                                    if (-2.25 > 3.14159) {
                                                                        delete(v4, "\n\t")
                                                                        fmt.Println(v46)
                                                                        if v46 {
                                                                            var v57 uint32 = 2147483647
                                                                            _ = v57
                                                                        } else {
                                                                            v31 -= (fi35 | fi38)
                                                                        }
                                                                    }
                                                                case "a":
                                                                    _ = len(v4)
                                                                }
                                                                v31 += v31
                                                                switch "" {
                                                                case "":
                                                                    var v58 string = ("\n\t" + "foo bar")
                                                                    _ = v58
                                                                case "hello":
                                                                    v59 := GenS4{F0: "\n\t", F1: 2, F2: (v2 & v2)}
                                                                    _ = v59
                                                                case "\n\t":
                                                                    var v60 int32 = v1
                                                                    _ = v60
                                                                default:
                                                                    if v47 {
                                                                        v31 = (v31 * 4294967295)
                                                                        fmt.Println(v31, fi38)
                                                                        var v61 float64 = (3.14159 / 3.14159)
                                                                        _ = v61
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v31, v3, v2)
                                                        } else {
                                                            v62 := []int{(fi35 - v31), v31, 2147483647, fi38}
                                                            var v63 float32 = 0.1
                                                            var v64 int = fi38
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        var v65 uint = 2
                                                        _ = v65
                                                    }
                                                } else {
                                                    var v66 int8 = genFn1()
                                                    v4[("" + ("hello" + ""))] = v46
                                                    var v67 uint8 = 46
                                                    _ = v66
                                                    _ = v67
                                                }
                                                if v46 {
                                                    switch ("hello" + "hello") {
                                                    case "foo bar":
                                                        v68 := GenS2{F0: v3, F1: 42, F2: 42}
                                                        _ = v68
                                                    case "\n\t":
                                                        if true {
                                                            var v69 uint = 58
                                                            _ = v69
                                                        } else {
                                                            if (!true) {
                                                                var v70 int = v31
                                                                _ = v70
                                                            } else {
                                                                var v71 string = "hello"
                                                                switch 1 {
                                                                case 42:
                                                                    v72 := GenS5{F0: 0, F1: 1}
                                                                    _ = v72
                                                                }
                                                                switch (fi35 % ((fi38) | 1)) {
                                                                case 9223372036854775807:
                                                                    if (0 == 6) {
                                                                        _ = len(v4)
                                                                        var v73 uint16 = v3
                                                                        fmt.Println(v1, v71)
                                                                        _ = v73
                                                                    }
                                                                case 546:
                                                                    v47 = v46
                                                                case 32767:
                                                                    v47 = true
                                                                    v74 := GenS6{F0: genFn0()}
                                                                    v1 -= (v1 & v1)
                                                                    _ = v74
                                                                }
                                                                _ = v71
                                                            }
                                                            var v75 int64 = 9223372036854775807
                                                            _ = v75
                                                        }
                                                    case "hello":
                                                        v31 = v31
                                                    default:
                                                        var v76 bool = true
                                                        _ = v76
                                                    }
                                                    if ((v2 - 42) == v2) {
                                                        fmt.Println(v1, v3, v1)
                                                    }
                                                    switch fi35 {
                                                    case 459:
                                                        var v77 string = "\n\t"
                                                        _ = v77
                                                    default:
                                                        fmt.Println(v47, v47, v2)
                                                        var v78 uint = 18446744073709551615
                                                        _ = v78
                                                    }
                                                    v79 := []int{(^(fi35)), 0, (fi35 - (fi35 & 789)), (fi35 >> ((973) & 63))}
                                                    switch ("" + "a") {
                                                    case "\n\t":
                                                        var v80 float32 = 3.14159
                                                        _ = v80
                                                    case "hello":
                                                        v81 := GenS4{F0: "foo bar", F1: 671, F2: 0}
                                                        _ = v81
                                                    default:
                                                        _ = len(v79)
                                                    }
                                                    var v82 uint = 857
                                                    _ = v79
                                                    _ = v82
                                                } else {
                                                    v83 := GenS2{F0: v3, F1: -2147483648, F2: 128}
                                                    switch "a" {
                                                    case "\n\t":
                                                        var v84 int = (fi35 + (-(v31)))
                                                        _ = v84
                                                    case "hello":
                                                        v85 := GenS7{F0: 127, F1: 2, F2: v3}
                                                        _ = v85
                                                    case "a":
                                                        if (v3 >= (v3 - (^(v3)))) {
                                                            v86 := GenS2{F0: (^(v3)), F1: int32(v2), F2: 128}
                                                            _ = v86
                                                        }
                                                    }
                                                    _ = v83
                                                }
                                                v31 += -1
                                                switch fi38 {
                                                case 32768:
                                                    v87 := GenS8{F0: v3}
                                                    _ = v87
                                                case 0:
                                                    if (!true) {
                                                        if v46 {
                                                            v88 := []int{fi35}
                                                            var v89 uint8 = 42
                                                            var v90 int8 = genFn1()
                                                            v1 = 0
                                                            v91 := GenS2{F0: v3, F1: 843, F2: 250}
                                                            _ = v88
                                                            _ = v89
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                    }
                                                case -9223372036854775808:
                                                    var v92 float64 = 1e10
                                                    _ = v92
                                                }
                                                if (true && v47) {
                                                    var v93 uint64 = 42
                                                    _ = v93
                                                }
                                                _ = v53
                                            }
                                            switch "a" {
                                            case "hello":
                                                v31 += 128
                                            }
                                            if (65535 >= 1) {
                                                v1 = 745
                                                var v94 float32 = 3.14159
                                                v1 += (v1 & (v1 & -656))
                                                fmt.Println(v3)
                                                _ = v94
                                            } else {
                                                fmt.Println(fi35)
                                            }
                                            if v47 {
                                                if true {
                                                    v95 := []int{(fi35 << ((0) & 63))}
                                                    v31 = fi38
                                                    _ = v95
                                                }
                                            }
                                        }
                                        var v96 float32 = (float32(0.0) - (0.1 / 1.0))
                                        _ = v96
                                    }
                                case "\n\t":
                                    if v46 {
                                        switch "a" {
                                        case "foo bar":
                                            var v97 float32 = 1.5
                                            _ = v97
                                        case "":
                                            v98 := GenS5{F0: 2, F1: -1}
                                            v3 = v3
                                            v47 = (!v46)
                                            _ = v98
                                        case "\n\t":
                                            if v46 {
                                                var v99 int8 = -114
                                                v1 -= 1
                                                var v100 uint = 667
                                                _ = v100
                                                _ = v99
                                            } else {
                                                fmt.Println(v31, v1)
                                            }
                                        default:
                                            v101 := GenS9{F0: 42, F1: 201}
                                            _ = v101
                                        }
                                        switch "a" {
                                        case "foo bar":
                                            delete(v4, "a")
                                        case "hello":
                                            v2 = (v2 - 42)
                                        }
                                        v102 := GenS10{F0: v3, F1: -9223372036854775808}
                                        _ = v102
                                    } else {
                                        v1 += v1
                                    }
                                }
                                var v103 int8 = 0
                                _ = v103
                            } else {
                                fmt.Println(v47, v31, v47)
                            }
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        if (!(false && (18446744073709551615 != 65535))) {
                            var v104 string = "foo bar"
                            _ = v104
                        }
                        _ = fi38
                    }
                    var v105 int16 = 131
                    _ = v105
                }
                _ = v36
            }
            v106 := GenS3{F0: 9223372036854775807, F1: 255, F2: ("hello" == "a")}
            fmt.Println(v31, v31)
            _ = fi35
            _ = v106
        }
        if false {
            v107 := []int{(v31 / ((v31) | 1)), v31}
            _ = v107
        }
    }
    for fi108 := 0; fi108 < 18; fi108++ {
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "a":
            for fi109 := 0; fi109 < 14; fi109++ {
                var v110 uint16 = (v3 << ((4294967295) & 63))
                var v111 int8 = genFn0()
                _ = fi109
                _ = v110
                _ = v111
            }
        }
        for fi112 := 0; fi112 < 6; fi112++ {
            if (true && (!(true && ((v3 + 32767) > v3)))) {
                switch "a" {
                case "a":
                    fmt.Println(fi108, v3)
                case "foo bar":
                    var v113 float32 = float32(float64(1e10))
                    _ = v113
                }
                v31 = 1
                var v114 int8 = 127
                _ = v114
            }
            _ = fi112
        }
        _ = fi108
    }
    v115 := GenS11{F0: -701, F1: -45, F2: 856}
    var v116 int8 = 100
    fmt.Println(v3)
    if true {
        if true {
            var v117 int = v31
            if true {
                var v118 string = "\n\t"
                fmt.Println(v116, v1, v1)
                v31 = (v31 << ((757) & 63))
                _ = v118
            } else {
                for fi119 := 0; fi119 < 10; fi119++ {
                    switch fi119 {
                    case 0:
                        var v120 float64 = 1.5
                        _ = v120
                    case -1:
                        fmt.Println(v31, v2)
                    case 32767:
                        fmt.Println(v2, v117)
                    default:
                        fmt.Println(v3, v1, v31)
                    }
                    _ = fi119
                }
            }
            _ = v117
        }
        var v121 uint8 = 178
        _ = v121
    } else {
        for fi122 := 0; fi122 < 16; fi122++ {
            v3 = 42
            v123 := []int{(fi122 | (v31 / (((-(v31))) | 1))), fi122}
            fmt.Println(v3)
            _ = fi122
            _ = v123
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
}
