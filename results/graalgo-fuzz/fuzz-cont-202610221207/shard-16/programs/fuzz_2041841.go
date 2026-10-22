package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 uint8 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint8; F1 bool; F2 int8 }
type GenS11 struct { F0 uint32; F1 uint }
func genFn0(p0 uint32, p1 string, p2 uint32) int {
    v6 := GenS1{F0: 18446744073709551615, F1: 32767, F2: (p0 >> ((995) & 63))}
    v7 := GenS1{F0: 767, F1: 42, F2: p2}
    p1 = p1
    v8 := GenS2{F0: 219}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    var v4 uint = (v1 ^ v1)
    v5 := make(map[uint32]string)
    v1 = v1
    var v9 int64 = -1
    v4 += 4294967295
    fmt.Println(v9, v9, v4)
    if ((v9 == -1) || true) {
        var v10 int32 = -289
        if (("\n\t" + (("" + "a") + ("\n\t" + ""))) != "\n\t") {
            for fi11 := 0; fi11 < 7; fi11++ {
                if ("hello" >= "foo bar") {
                    v12 := GenS1{F0: 0, F1: 42, F2: 2}
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v4 += 889
                        switch (fi13 % (((fi13 + 2147483647)) | 1)) {
                        case 0:
                            var v14 int8 = 82
                            _ = v14
                        case 127:
                            var v15 uint8 = 128
                            _ = v15
                        case -2147483648:
                            switch 128 {
                            case -9223372036854775808:
                                var v16 int64 = int64(v10)
                                _ = v16
                            case 2147483647:
                                if (v9 != v9) {
                                    v17 := []int{fi13, (^(fi13)), (fi11 >> ((v1) & 63))}
                                    _ = v17
                                } else {
                                    fmt.Println(v9, v1)
                                }
                                var v18 uint = v4
                                _ = v18
                            case 127:
                                switch "hello" {
                                case "hello":
                                    v19 := GenS2{F0: 2}
                                    _ = v19
                                case "":
                                    v20 := GenS3{F0: ("foo bar" + ""), F1: v10, F2: 4294967295}
                                    _ = v20
                                case "foo bar":
                                    var v21 int32 = -1
                                    _ = v21
                                default:
                                    switch fi11 {
                                    case 130:
                                        v10 += (v10 + (v10 / ((v10) | 1)))
                                    case 1:
                                        v22 := GenS4{F0: 672}
                                        _ = v22
                                    case -1:
                                        var v23 float64 = float64(0.0)
                                        _ = v23
                                    }
                                }
                            }
                        }
                        v24 := []int{}
                        var v25 uint16 = 748
                        _ = fi13
                        _ = v24
                        _ = v25
                    }
                    if (4294967295 >= 4294967295) {
                        v26 := v5[42]
                        var v27 uint32 = 735
                        _ = v26
                        _ = v27
                    } else {
                        v4 = 65535
                        if false {
                            v10 = (^(v10))
                            v28 := GenS3{F0: ("hello" + "a"), F1: (v10 + (v10 % (((v10 & v10)) | 1))), F2: 341}
                            _ = v28
                        } else {
                            fmt.Println(v9, v4)
                            _ = len(v3)
                        }
                        if true {
                            for fi29 := 0; fi29 < 18; fi29++ {
                                var v30 float64 = -2.25
                                fmt.Println(v9, v4)
                                v31 := []int{}
                                if (((v10 - 318) != v10) && true) {
                                    if ((1e10 * 1.5) >= 0.1) {
                                        if false {
                                            var v32 string = ("hello" + ("a" + "foo bar"))
                                            if (v10 < v10) {
                                                var v33 uint32 = 904
                                                fmt.Println(v10, v32, v4)
                                                v34 := GenS1{F0: (v1 & v4), F1: 1, F2: v33}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                var v35 string = (v32 + v32)
                                                if false {
                                                    if (fi29 == (fi29 % ((fi29) | 1))) {
                                                        v32 = (((v32 + v32) + (v32 + "\n\t")) + "\n\t")
                                                    } else {
                                                        v1 = v1
                                                        v30 = v30
                                                    }
                                                    var v36 uint = v1
                                                    v4 = v4
                                                    v37 := GenS5{F0: 2, F1: 2147483647}
                                                    if (42 < 21) {
                                                        var v38 int16 = 1000
                                                        _ = v38
                                                    }
                                                    switch ((v35 + "") + "a") {
                                                    case "\n\t":
                                                        var v39 uint8 = 128
                                                        _ = v39
                                                    case "a":
                                                        v40 := GenS3{F0: v32, F1: v10, F2: 42}
                                                        _ = v40
                                                    case "hello":
                                                        var v41 int = fi29
                                                        v41 -= genFn0(128, v32, 1)
                                                        var v42 uint16 = 1000
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                    v10 -= v10
                                                    v43 := GenS1{F0: v1, F1: 443, F2: 634}
                                                    _ = v36
                                                    _ = v37
                                                    _ = v43
                                                }
                                                if ((true && false) && (!false)) {
                                                    var v44 int64 = (v9 + -1)
                                                    switch fi11 {
                                                    case 127:
                                                        switch -201 {
                                                        case -363:
                                                            v45 := GenS1{F0: (v1 >> ((v4) & 63)), F1: 2, F2: 100}
                                                            _ = v45
                                                        }
                                                    }
                                                    v30 = (((((0.1 + float64((0.0 + (0.0 * 1e10)))) - float64(0.1)) + 0.1) / v30) / 0.1)
                                                    fmt.Println(v9)
                                                    var v46 uint16 = 74
                                                    v47 := GenS2{F0: 42}
                                                    _ = v44
                                                    _ = v46
                                                    _ = v47
                                                } else {
                                                    var v48 int64 = (-(v9))
                                                    _ = v48
                                                }
                                                _ = v35
                                            }
                                            _ = v32
                                        } else {
                                            fmt.Println(v4)
                                            var v49 uint64 = 567
                                            v30 = float64(1e10)
                                            _ = v49
                                        }
                                        if len(v31) > 0 {
                                            v31[0] = (-(fi29))
                                        }
                                    } else {
                                        if ((!false) && (-119 < 100)) {
                                            var v50 float32 = 1e10
                                            if (false || (42 <= 393)) {
                                                var v51 uint16 = 42
                                                _ = v51
                                            } else {
                                                v52 := GenS4{F0: 0}
                                                var v53 uint = v1
                                                var v54 uint64 = 4294967295
                                                if true {
                                                    var v55 uint64 = (v54 * 1)
                                                    _ = v55
                                                }
                                                v53 += v53
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                            }
                                            v56 := GenS6{F0: (v9 << ((v1) & 63))}
                                            _ = v50
                                            _ = v56
                                        }
                                        switch -1 {
                                        case 679:
                                            v57 := GenS4{F0: 255}
                                            _ = v57
                                        case 32767:
                                            var v58 int = (fi29 ^ fi29)
                                            _ = v58
                                        }
                                    }
                                    v9 += -1
                                    var v59 float64 = (-2.25 - 0.1)
                                    _ = v59
                                }
                                var v60 float64 = (v30 + 1.5)
                                var v61 uint8 = 42
                                var v62 int16 = -32768
                                v63 := GenS2{F0: v61}
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v60
                                _ = v61
                                _ = v62
                                _ = v63
                            }
                        }
                    }
                    _ = v12
                }
                _ = fi11
            }
        } else {
            switch "a" {
            case "\n\t":
                _ = len(v5)
            case "foo bar":
                v1 += v4
            }
            for fi64 := 0; fi64 < 15; fi64++ {
                if false {
                    for fi65 := 0; fi65 < 14; fi65++ {
                        var v66 int64 = (v9 >> ((v1) & 63))
                        _ = fi65
                        _ = v66
                    }
                    v1 = (v4 ^ 65535)
                }
                _ = fi64
            }
            var v67 int8 = 42
            _ = v67
        }
        for fi68 := 0; fi68 < 19; fi68++ {
            var v69 float64 = float64(float32(1e10))
            if false {
                v70 := GenS6{F0: v9}
                for fi71 := 0; fi71 < 8; fi71++ {
                    fmt.Println(v69, fi68, fi71)
                    if true {
                        v72 := GenS7{F0: 74, F1: 42}
                        v10 -= -627
                        _ = v72
                    } else {
                        fmt.Println(v10, fi71, v9)
                        if (42 == 255) {
                            if true {
                                fmt.Println(v4)
                            } else {
                                if true {
                                    if false {
                                        var v73 bool = true
                                        fmt.Println(v4)
                                        v1 = (v1 % ((42) | 1))
                                        _ = v73
                                    } else {
                                        var v74 int16 = 1000
                                        if ((1.5 < (0.1 - 0.0)) && (42 <= 923)) {
                                            v74 -= (v74 >> ((v1) & 63))
                                            v75 := GenS8{F0: v9, F1: (32767 == v74)}
                                            v76 := []int{-435, fi68}
                                            _ = v75
                                            _ = v76
                                        }
                                        var v77 bool = (false || (fi68 != fi68))
                                        _ = v74
                                        _ = v77
                                    }
                                    var v78 string = "foo bar"
                                    switch fi68 {
                                    case 2147483647:
                                        var v79 int16 = 0
                                        _ = v79
                                    case -578:
                                        fmt.Println(v10, v4)
                                    case 9223372036854775807:
                                        delete(v2, "\n\t")
                                    default:
                                        v80 := GenS9{F0: -1, F1: (v9 | v9)}
                                        _ = v80
                                    }
                                    _ = v78
                                }
                            }
                            var v81 uint16 = 0
                            _ = v81
                        } else {
                            var v82 int64 = (v9 | 9223372036854775807)
                            var v83 uint32 = 489
                            if false {
                                switch 209 {
                                case -1:
                                    v84 := []int{fi71, (-(fi68)), 9223372036854775807}
                                    _ = v84
                                case 128:
                                    var v85 int8 = 1
                                    v86 := []int{(fi71 * (fi68 + 9223372036854775807)), genFn0(v83, "", (v83 << (((v4 ^ (v1 | v1))) & 63)))}
                                    _ = v85
                                    _ = v86
                                default:
                                    switch "\n\t" {
                                    case "":
                                        v87 := GenS10{F0: 155, F1: ((0.0 - 1.5) <= ((v69 + v69) - ((v69 * -2.25) / v69))), F2: 39}
                                        _ = v87
                                    case "hello":
                                        if false {
                                            v3[int(v1)] = (fi71 >> ((v4) & 63))
                                            v9 = v82
                                            var v88 int8 = -128
                                            _ = v88
                                        }
                                    case "a":
                                        if true {
                                            var v89 int = -269
                                            if false {
                                                fmt.Println(v83, fi68, fi68)
                                                var v90 uint32 = v83
                                                var v91 uint64 = 0
                                                v5[0] = "\n\t"
                                                var v92 int32 = v10
                                                v93 := GenS4{F0: 32767}
                                                _ = v90
                                                _ = v91
                                                _ = v92
                                                _ = v93
                                            } else {
                                                var v94 int16 = 727
                                                if (-1 != v94) {
                                                    var v95 bool = true
                                                    _ = v95
                                                }
                                                if true {
                                                    v96 := GenS5{F0: 564, F1: v83}
                                                    var v97 float32 = float32(float64(((1e10 / (0.1 / 1.0)) * (0.1 / float32(1.5)))))
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    var v98 uint16 = 174
                                                    var v99 int = (fi68 + -487)
                                                    var v100 uint32 = v83
                                                    _ = v100
                                                    _ = v98
                                                    _ = v99
                                                }
                                                _ = v94
                                            }
                                            _ = v89
                                        }
                                    }
                                }
                            }
                            _ = v82
                            _ = v83
                        }
                    }
                    _ = fi71
                }
                _ = v70
            }
            _ = fi68
            _ = v69
        }
        _ = v10
    } else {
        switch ("\n\t" + "a") {
        case "foo bar":
            v101 := GenS8{F0: 42, F1: true}
            _ = v101
        case "hello":
            v102 := GenS11{F0: 2, F1: v1}
            _ = v102
        case "\n\t":
            v1 = (v1 << ((v1) & 63))
        }
        var v103 uint32 = 844
        fmt.Println(v1)
        _ = v103
    }
    for fi104 := 0; fi104 < 1; fi104++ {
        v4 = 524
        fmt.Println(v9)
        v1 -= (v1 + v1)
        _ = fi104
    }
    v105 := []int{-1}
    if (("foo bar" == "hello") && false) {
        if false {
            var v106 int16 = 1
            _ = v106
        } else {
            v9 += v9
            switch 555 {
            case 128:
                if ((127 >= 127) && (((1e10 / 3.14159) + 3.14159) <= float64((0.0 * (3.14159 * 1.5))))) {
                    v4 -= v4
                }
            case -1:
                var v107 uint16 = uint16(v1)
                _ = v107
            case 127:
                if true {
                    var v108 int16 = 1
                    var v109 float32 = 0.1
                    v9 -= (v9 + -795)
                    _ = v108
                    _ = v109
                }
            }
        }
    }
    v9 = (v9 & -9223372036854775808)
    switch "\n\t" {
    case "a":
        v110 := v3[2147483647]
        _ = v110
    }
    var v111 uint16 = 65535
    if true {
        var v112 float64 = 1.5
        switch "hello" {
        case "hello":
            for fi113 := 0; fi113 < 17; fi113++ {
                v1 -= 584
                fmt.Println(v111, v1, v4)
                fmt.Println(v1, v9)
                _ = fi113
            }
        case "a":
            if (32767 == -1) {
                v111 = v111
                v1 -= v4
            } else {
                v4 += (v1 & (v1 - (v1 * v4)))
                for fi114 := 0; fi114 < 11; fi114++ {
                    var v115 bool = (!true)
                    var v116 int16 = 1
                    _ = fi114
                    _ = v115
                    _ = v116
                }
            }
        case "foo bar":
            var v117 bool = false
            _ = v117
        default:
            fmt.Println(v9, v111)
        }
        _ = v112
    }
    if len(v105) >= 2 {
        v105 = v105[1:len(v105)]
    }
    v4 -= (v1 * 0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v111:", v111)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v105
    _ = v111
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
