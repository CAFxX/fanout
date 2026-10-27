package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int64; F1 int; F2 uint32 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 string; F1 int; F2 uint64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 uint64; F2 int32 }
type GenS10 struct { F0 string; F1 string; F2 int }
func genFn0(p0 int32, p1 int) int16 {
    p1 -= 32767
    p0 += p0
    v6 := []int{p1}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 float32 = 3.14159
    var v5 int16 = 42
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(fi7, v4)
        _ = fi7
    }
    v5 = 42
    switch "a" {
    case "foo bar":
        switch -98 {
        case 32768:
            delete(v1, -84)
        default:
            v3[1] = "foo bar"
        }
    case "":
        v8 := GenS1{F0: 675}
        _ = v8
    }
    switch 42 {
    case 998:
        var v9 float64 = (float64(v4) - 1e10)
        _ = v9
    default:
        v5 += genFn0(0, 2147483647)
    }
    v10 := GenS1{F0: 4294967295}
    if (128 <= 423) {
        v11 := GenS1{F0: 0}
        var v12 int64 = -1
        var v13 uint8 = 2
        _ = v11
        _ = v12
        _ = v13
    } else {
        switch "foo bar" {
        case "foo bar":
            var v14 int8 = -128
            _ = v14
        case "":
            v15 := []int{int(v5)}
            _ = v15
        case "\n\t":
            for fi16 := 0; fi16 < 15; fi16++ {
                for fi17 := 0; fi17 < 17; fi17++ {
                    v4 = (0.1 - (v4 - (((0.0 / 1.0) * v4) - -2.25)))
                    v3[0] = "a"
                    if ((0.1 + 0.1) >= 0.1) {
                        v4 = v4
                        if false {
                            switch "\n\t" {
                            case "":
                                var v18 float64 = (1e10 * 1e10)
                                _ = v18
                            case "foo bar":
                                v5 -= genFn0(0, -1)
                                v5 = int16(fi17)
                                fmt.Println(fi16)
                            case "a":
                                v5 = v5
                            default:
                                v5 += 42
                            }
                            var v19 int32 = 42
                            switch fi16 {
                            case -675:
                                var v20 int64 = 42
                                _ = v20
                            }
                            _ = v19
                        }
                    } else {
                        if ((fi16 | fi16) < fi16) {
                            v5 = v5
                        }
                        var v21 string = ""
                        _ = v21
                    }
                    _ = fi17
                }
                _ = fi16
            }
        default:
            fmt.Println(v4, v5)
        }
        var v22 uint64 = 684
        if true {
            fmt.Println(v22, v5, v22)
        } else {
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 uint16 = 340
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 17; fi25++ {
                switch (fi25 & (fi25 >> ((2) & 63))) {
                case 32768:
                    _ = len(v2)
                case 42:
                    if ((fi25 <= fi25) || false) {
                        var v26 uint16 = 1
                        fmt.Println(v4, fi25)
                        v27 := []int{32768, fi25, 0}
                        _ = v26
                        _ = v27
                    } else {
                        v4 += v4
                        v5 = (v5 % ((v5) | 1))
                        var v28 float32 = v4
                        _ = v28
                    }
                case 476:
                    for fi29 := 0; fi29 < 10; fi29++ {
                        fmt.Println(v4, v5)
                        v30 := []int{}
                        _ = fi29
                        _ = v30
                    }
                default:
                    switch fi25 {
                    case -1:
                        v4 = ((((1.5 / v4) * float32((float64(3.14159) - 0.1))) * ((0.1 * 0.0) / (v4 * (v4 - v4)))) + (v4 + (v4 - -2.25)))
                    case 128:
                        fmt.Println(fi25, v4, v4)
                    case 127:
                        if false {
                            for fi31 := 0; fi31 < 14; fi31++ {
                                if false {
                                    var v32 bool = (((197 != 127) && (!(false || false))) && (genFn0(1, fi25) > (v5 % ((int16(v22)) | 1))))
                                    if v32 {
                                        delete(v2, "foo bar")
                                        v5 = -1
                                    }
                                    v33 := GenS2{F0: -81, F1: 2, F2: v32}
                                    v4 += v4
                                    _ = v32
                                    _ = v33
                                } else {
                                    if false {
                                        v34 := GenS2{F0: 1, F1: 0, F2: true}
                                        var v35 uint64 = v22
                                        _ = v34
                                        _ = v35
                                    } else {
                                        v5 += v5
                                        var v36 string = ""
                                        if false {
                                            var v37 int32 = 2147483647
                                            switch fi25 {
                                            case 0:
                                                if false {
                                                    v36 = v36
                                                    if (!(!true)) {
                                                        v36 = v36
                                                    }
                                                    v5 = v5
                                                }
                                            case 2147483647:
                                                _ = len(v1)
                                            }
                                            if (!(!true)) {
                                                v36 = (("foo bar" + "") + (v36 + "\n\t"))
                                                v38 := GenS1{F0: (v22 | (v22 & (v22 % (((v22 / (((^(v22))) | 1))) | 1))))}
                                                var v39 uint32 = 42
                                                v3[v39] = "foo bar"
                                                _ = v38
                                                _ = v39
                                            } else {
                                                v5 = 42
                                            }
                                            if (v22 != 722) {
                                                if ((-2.25 * 0.0) == v4) {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v22 = 1
                                                    default:
                                                        switch ("a" + (v36 + ("a" + (v36 + "\n\t")))) {
                                                        case "hello":
                                                            var v40 int32 = (v37 ^ v37)
                                                            _ = v40
                                                        case "a":
                                                            v41 := GenS2{F0: -30, F1: 42, F2: false}
                                                            _ = v41
                                                        case "foo bar":
                                                            if (824 != 0) {
                                                                v37 = v37
                                                                var v42 int8 = 105
                                                                _ = v42
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v36, v37)
                                                    var v43 int16 = v5
                                                    _ = v43
                                                } else {
                                                    if (true || ((v22 + v22) < (^(v22)))) {
                                                        v44 := GenS2{F0: -1, F1: 65535, F2: (v36 != (v36 + "\n\t"))}
                                                        fmt.Println(v5)
                                                        fmt.Println(v22)
                                                        _ = v44
                                                    }
                                                    v45 := GenS1{F0: v22}
                                                    v22 = v22
                                                    v46 := GenS3{F0: 590, F1: 32768}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                if true {
                                                    fmt.Println(fi25, v37)
                                                }
                                            } else {
                                                if false {
                                                    var v47 uint32 = 1
                                                    v48 := GenS4{F0: 302, F1: 18, F2: 1}
                                                    if true {
                                                        switch fi25 {
                                                        case -9223372036854775808:
                                                            var v49 int64 = -1
                                                            _ = v49
                                                        case 784:
                                                            if true {
                                                                fmt.Println(v4, v4, fi31)
                                                                fmt.Println(v5)
                                                                v1[fi25] = ("\n\t" + v36)
                                                            } else {
                                                                v50 := v1[9223372036854775807]
                                                                fmt.Println(fi31)
                                                                switch fi25 {
                                                                case 605:
                                                                    if (v37 >= (v37 % ((0) | 1))) {
                                                                        v22 = v22
                                                                        v5 += genFn0((v37 * v37), 42)
                                                                        delete(v2, v36)
                                                                        switch int(v22) {
                                                                        case -773:
                                                                            fmt.Println(v50, v36)
                                                                        case 128:
                                                                            fmt.Println(fi31, v36)
                                                                        case 32768:
                                                                            fmt.Println(fi25, fi31)
                                                                            var v51 uint16 = 52
                                                                            _ = v51
                                                                        }
                                                                    } else {
                                                                        v47 += v47
                                                                        if true {
                                                                            switch fi31 {
                                                                            case 32768:
                                                                                if false {
                                                                                    switch fi25 {
                                                                                    case 0:
                                                                                        v37 -= 0
                                                                                    case -9223372036854775808:
                                                                                        v5 -= (v5 % (((-(v5))) | 1))
                                                                                    case -1:
                                                                                        v52 := GenS5{F0: 1, F1: fi25, F2: 0}
                                                                                        _ = v52
                                                                                    default:
                                                                                        fmt.Println(v5, v47)
                                                                                        v53 := GenS6{F0: v47, F1: 32767}
                                                                                        fmt.Println(v5, v37, v50)
                                                                                        var v54 int8 = 1
                                                                                        v55 := GenS5{F0: 363, F1: 63, F2: (v47 ^ v47)}
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                        _ = v55
                                                                                    }
                                                                                }
                                                                            case 4294967295:
                                                                                var v56 string = v50
                                                                                _ = v56
                                                                            case 32767:
                                                                                v57 := []int{}
                                                                                _ = v57
                                                                            default:
                                                                                v5 = v5
                                                                            }
                                                                            var v58 int8 = 1
                                                                            var v59 uint64 = v22
                                                                            _ = v58
                                                                            _ = v59
                                                                        } else {
                                                                            var v60 int32 = 42
                                                                            v1[(fi25 | fi31)] = (("" + v36) + "foo bar")
                                                                            _ = v60
                                                                        }
                                                                        v36 = "hello"
                                                                    }
                                                                case 4294967295:
                                                                    v61 := []int{(fi25 - -1), (fi25 ^ (^(fi31))), (fi25 & 128), fi25}
                                                                    _ = v61
                                                                case 42:
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        var v62 int32 = 1
                                                                        _ = v62
                                                                    default:
                                                                        var v63 int64 = -1
                                                                        _ = v63
                                                                    }
                                                                default:
                                                                    v64 := GenS5{F0: 1, F1: (fi31 * -1), F2: v47}
                                                                    _ = v64
                                                                }
                                                                _ = v50
                                                            }
                                                            v65 := []int{fi31, fi31}
                                                            _ = v65
                                                        default:
                                                            if (true || (!false)) {
                                                                v66 := []int{int(v37), (-(fi25)), fi31, (fi25 * fi31)}
                                                                v4 -= (3.14159 / v4)
                                                                _ = v66
                                                            }
                                                        }
                                                        switch v36 {
                                                        case "":
                                                            if (3.14159 <= 1.5) {
                                                                switch -9223372036854775808 {
                                                                case 439:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v67 int32 = v37
                                                                        _ = v67
                                                                    case "\n\t":
                                                                        v68 := []int{(fi25 * (fi31 + fi25)), 0, fi31}
                                                                        _ = v68
                                                                    case "hello":
                                                                        v69 := []int{(fi31 & -2147483648), fi25, fi31}
                                                                        _ = v69
                                                                    default:
                                                                        v70 := v2[v36]
                                                                        var v71 uint32 = v47
                                                                        delete(v3, 65535)
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                case 0:
                                                                    v37 += 1
                                                                case 4294967295:
                                                                    var v72 uint32 = v47
                                                                    _ = v72
                                                                default:
                                                                    if true {
                                                                        v73 := GenS7{F0: "", F1: fi31, F2: (^(v22))}
                                                                        _ = v73
                                                                    } else {
                                                                        v74 := GenS2{F0: 0, F1: 651, F2: false}
                                                                        _ = v74
                                                                    }
                                                                }
                                                                if ((!true) || (v5 < v5)) {
                                                                    fmt.Println(fi25)
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v47 = v47
                                                        if (636 < 1) {
                                                            if (!((-2.25 <= (((1.5 * 1.5) * v4) - (v4 + v4))) && false)) {
                                                                v75 := GenS8{F0: -9223372036854775808}
                                                                fmt.Println(fi25, v47, v5)
                                                                _ = v75
                                                            }
                                                        }
                                                    }
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    var v76 float32 = (v4 - v4)
                                                    _ = v76
                                                }
                                            }
                                            var v77 uint64 = 4294967295
                                            _ = v37
                                            _ = v77
                                        } else {
                                            switch fi25 {
                                            case 4294967295:
                                                var v78 float64 = 1.5
                                                _ = v78
                                            case 287:
                                                v22 = (v22 * (^(v22)))
                                            case 1:
                                                var v79 bool = false
                                                _ = v79
                                            default:
                                                var v80 int32 = 1
                                                _ = v80
                                            }
                                            v81 := v1[1]
                                            _ = v81
                                        }
                                        v3[0] = "a"
                                        _ = v36
                                    }
                                    var v82 int = (fi25 % (((fi31 >> ((255) & 63))) | 1))
                                    v83 := []int{-2147483648, (fi31 >> ((377) & 63))}
                                    _ = v82
                                    _ = v83
                                }
                                switch ("\n\t" + "\n\t") {
                                case "foo bar":
                                    switch fi25 {
                                    case -9223372036854775808:
                                        v84 := GenS9{F0: -32768, F1: (v22 ^ (^(v22))), F2: 42}
                                        _ = v84
                                    case 32768:
                                        v85 := GenS3{F0: 1, F1: (fi25 << ((255) & 63))}
                                        _ = v85
                                    }
                                }
                                _ = fi31
                            }
                            v86 := []int{}
                            v4 += v4
                            _ = v86
                        } else {
                            v87 := GenS9{F0: (v5 + genFn0(-247, 0)), F1: (v22 % (((v22 >> ((0) & 63))) | 1)), F2: 1}
                            delete(v2, ("\n\t" + "\n\t"))
                            if (117 == 42) {
                                if true {
                                    var v88 int16 = (v5 >> ((752) & 63))
                                    for fi89 := 0; fi89 < 11; fi89++ {
                                        switch fi89 {
                                        case 552:
                                            fmt.Println(v88, v88, v4)
                                        case -9223372036854775808:
                                            v90 := []int{(-(fi89))}
                                            _ = v90
                                        case 4294967295:
                                            v5 -= 42
                                        default:
                                            if false {
                                                var v91 int64 = int64(v22)
                                                v92 := GenS1{F0: v22}
                                                _ = v91
                                                _ = v92
                                            } else {
                                                if true {
                                                    switch (fi89 / ((fi89) | 1)) {
                                                    case -375:
                                                        var v93 float64 = 0.0
                                                        _ = v93
                                                    }
                                                } else {
                                                    if (430 < 2147483647) {
                                                        v22 = 433
                                                        v94 := GenS1{F0: v22}
                                                        _ = v94
                                                    } else {
                                                        v22 += v22
                                                        var v95 int32 = int32(v88)
                                                        fmt.Println(v22, fi25, v95)
                                                        _ = v95
                                                    }
                                                    var v96 int32 = -1
                                                    if true {
                                                        v97 := GenS3{F0: 1, F1: (fi25 - (fi25 / ((881) | 1)))}
                                                        v96 = v96
                                                        var v98 float32 = (0.1 * v4)
                                                        switch ("a" + "foo bar") {
                                                        case "\n\t":
                                                            var v99 float64 = 0.1
                                                            _ = v99
                                                        case "a":
                                                            v100 := v2["\n\t"]
                                                            _ = v100
                                                        default:
                                                            fmt.Println(fi25)
                                                        }
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    _ = v96
                                                }
                                                v22 += v22
                                                switch "hello" {
                                                case "hello":
                                                    v101 := GenS7{F0: ("a" + "hello"), F1: (fi89 << ((0) & 63)), F2: 965}
                                                    _ = v101
                                                case "foo bar":
                                                    var v102 float32 = 0.1
                                                    _ = v102
                                                default:
                                                    var v103 int32 = -520
                                                    _ = v103
                                                }
                                            }
                                        }
                                        v5 += (^(v88))
                                        _ = fi89
                                    }
                                    v5 = v5
                                    _ = v88
                                }
                                delete(v3, 2)
                                var v104 int8 = -1
                                _ = v104
                            } else {
                                v105 := GenS10{F0: "hello", F1: ("foo bar" + ""), F2: (fi25 + (fi25 << ((1) & 63)))}
                                _ = v105
                            }
                            _ = v87
                        }
                    default:
                        var v106 uint16 = 1
                        _ = v106
                    }
                }
                _ = fi25
            }
            fmt.Println(v4, v4)
        }
        _ = v22
    }
    for fi107 := 0; fi107 < 1; fi107++ {
        var v108 int32 = -2147483648
        var v109 float64 = (1.5 / (0.0 - -2.25))
        v109 = v109
        _ = fi107
        _ = v108
        _ = v109
    }
    if false {
        v4 = v4
    } else {
        if true {
            v4 += v4
            var v110 int = 542
            fmt.Println(v110, v110)
            _ = v110
        } else {
            if true {
                var v111 string = "a"
                v112 := GenS7{F0: v111, F1: -271, F2: 4294967295}
                _ = v111
                _ = v112
            }
            _ = len(v3)
        }
        fmt.Println(v5)
        if true {
            v4 += 0.1
        } else {
            for fi113 := 0; fi113 < 15; fi113++ {
                var v114 float64 = ((3.14159 + (float64(-2.25) * -2.25)) * -2.25)
                v114 = -2.25
                if (65535 >= 308) {
                    fmt.Println(v5)
                } else {
                    switch -576 {
                    case 1:
                        var v115 uint16 = 0
                        _ = v115
                    case -2147483648:
                        v4 += 0.1
                    case 9223372036854775807:
                        v5 -= (v5 % ((genFn0(777, (-(fi113)))) | 1))
                    }
                    v5 -= (-(v5))
                }
                _ = fi113
                _ = v114
            }
        }
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
