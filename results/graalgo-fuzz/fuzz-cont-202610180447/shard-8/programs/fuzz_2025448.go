package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 string; F2 uint64 }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint; F2 uint8 }
type GenS8 struct { F0 uint16; F1 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool; F1 int8 }
type GenS12 struct { F0 uint16; F1 int; F2 int64 }
type GenS13 struct { F0 int8; F1 uint }
func genFn0() uint32 {
    var v6 uint32 = 2147483648
    fmt.Println(v6)
    for fi7 := 0; fi7 < 20; fi7++ {
        if true {
            if false {
                if true {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        if false {
                            var v9 uint32 = 2147483647
                            if false {
                                var v10 int8 = -5
                                if false {
                                    v10 -= 1
                                }
                                _ = v10
                            }
                            var v11 int = fi8
                            _ = v11
                            _ = v9
                        }
                        v12 := GenS1{F0: "hello", F1: (!true), F2: v6}
                        v6 = 80
                        _ = fi8
                        _ = v12
                    }
                } else {
                    v13 := GenS1{F0: "a", F1: ((v6 >> ((875) & 63)) >= (v6 << ((65535) & 63))), F2: 0}
                    fmt.Println(v6)
                    var v14 uint32 = (v6 << ((4294967295) & 63))
                    _ = v13
                    _ = v14
                }
                var v15 int16 = 42
                v16 := GenS2{F0: (v6 & (v6 % (((v6 * (v6 + (v6 + (v6 >> ((42) & 63)))))) | 1))), F1: 65535}
                _ = v15
                _ = v16
            } else {
                fmt.Println(fi7, fi7)
            }
        }
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS2{F0: v6, F1: 4294967295}
            v6 = (v6 / ((v6) | 1))
            _ = fi17
            _ = v18
        }
        _ = fi7
    }
    fmt.Println(v6)
    _ = v6
    return v6
}
func genFn1(p0 string, p1 uint, p2 uint16) bool {
    v19 := []int{}
    p2 = p2
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            p0 = p0
            var v22 float64 = float64(3.14159)
            switch p0 {
            case "\n\t":
                var v23 int32 = 42
                _ = v23
            case "foo bar":
                v24 := GenS1{F0: p0, F1: true, F2: 42}
                _ = v24
            default:
                v25 := GenS3{F0: 567}
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    if ((false || true) && ((42 >= 127) && (genFn0() > 2147483647))) {
        v26 := GenS1{F0: "a", F1: (42 == 0), F2: 286}
        _ = v26
    } else {
        fmt.Println(p0, p0, p2)
        var v27 uint8 = 127
        _ = v27
    }
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return (false && (!((true || (42 > 15)) || false)))
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{42}
    v3 := []uint32{0, 42, 65535}
    v4 := make(map[int]uint32)
    v5 := []uint32{42, 775}
    var v28 int16 = 645
    for fi29 := 0; fi29 < 9; fi29++ {
        v28 = v28
        v28 = -32768
        for fi30 := 0; fi30 < 18; fi30++ {
            v1[(fi29 / (((-(fi30))) | 1))] = 9223372036854775807
            fmt.Println(v28)
            v28 -= (v28 & (v28 - 32767))
            _ = fi30
        }
        _ = fi29
    }
    if (575 != 79) {
        var v31 float32 = 0.1
        if (-1 < 42) {
            var v32 bool = (false || (1 != (v28 - (-(v28)))))
            var v33 int8 = 0
            _ = v32
            _ = v33
        }
        var v34 uint8 = 128
        _ = v31
        _ = v34
    } else {
        var v35 int16 = (-(v28))
        v28 = v35
        _ = v35
    }
    fmt.Println(v28)
    fmt.Println(v28)
    fmt.Println(v28)
    for fi36 := 0; fi36 < 2; fi36++ {
        var v37 int8 = 29
        if (!false) {
            switch "" {
            case "a":
                for fi38 := 0; fi38 < 14; fi38++ {
                    v37 = v37
                    v37 += 41
                    _ = fi38
                }
            case "":
                var v39 uint8 = 255
                _ = v39
            case "\n\t":
                fmt.Println(fi36)
            default:
                if false {
                    if (v37 > 42) {
                        v4[4294967295] = 506
                    } else {
                        for fi40 := 0; fi40 < 1; fi40++ {
                            switch fi40 {
                            case -445:
                                switch ("\n\t" + ((("\n\t" + "a") + ("foo bar" + "foo bar")) + "\n\t")) {
                                case "a":
                                    v28 -= (-(v28))
                                case "foo bar":
                                    var v41 int = fi40
                                    _ = v41
                                case "hello":
                                    var v42 string = ((("foo bar" + "hello") + "a") + "hello")
                                    _ = v42
                                default:
                                    v28 += 1
                                }
                            case 128:
                                var v43 uint16 = 313
                                if false {
                                    if false {
                                        v43 -= v43
                                        v28 = v28
                                        if true {
                                            if true {
                                                switch "foo bar" {
                                                case "":
                                                    v44 := GenS1{F0: "\n\t", F1: (!(v43 != v43)), F2: 2}
                                                    _ = v44
                                                case "\n\t":
                                                    var v45 int16 = v28
                                                    _ = v45
                                                case "hello":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            } else {
                                                switch 187 {
                                                case 2147483647:
                                                    if false {
                                                        v46 := v4[fi40]
                                                        var v47 int = 2147483647
                                                        v46 = v46
                                                        var v48 bool = (42 > v37)
                                                        _ = v46
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                case 4294967295:
                                                    v49 := GenS3{F0: (v43 / ((v43) | 1))}
                                                    _ = v49
                                                }
                                                switch (fi40 >> ((1) & 63)) {
                                                case 128:
                                                    v43 -= (v43 / (((v43 ^ 2)) | 1))
                                                case -2147483648:
                                                    switch "" {
                                                    case "":
                                                        v28 = 0
                                                    case "\n\t":
                                                        v50 := GenS3{F0: v43}
                                                        _ = v50
                                                    default:
                                                        v5 = append(v5, genFn0())
                                                    }
                                                }
                                                switch "" {
                                                case "":
                                                    v51 := v1[fi36]
                                                    _ = v51
                                                default:
                                                    switch 127 {
                                                    case 128:
                                                        v37 -= v37
                                                    }
                                                }
                                            }
                                            if true {
                                                switch "\n\t" {
                                                case "hello":
                                                    v37 = -1
                                                case "a":
                                                    var v52 int = -1
                                                    _ = v52
                                                }
                                                v53 := GenS3{F0: (v43 % ((0) | 1))}
                                                v1[(^(fi36))] = fi36
                                                _ = v53
                                            }
                                        } else {
                                            fmt.Println(fi40, v37, fi40)
                                        }
                                    }
                                }
                                var v54 float64 = (float64(-2.25) + ((float64(float32(0.0)) - (-2.25 * float64(1e10))) / (0.1 + 3.14159)))
                                _ = v43
                                _ = v54
                            case 896:
                                if false {
                                    switch "a" {
                                    case "foo bar":
                                        if true {
                                            v28 += (v28 << ((515) & 63))
                                            v28 = 1
                                        } else {
                                            if (v37 < 100) {
                                                _ = len(v5)
                                            }
                                            if true {
                                                v37 = (v37 / (((v37 << ((44) & 63))) | 1))
                                                v28 = v28
                                                var v55 float32 = 1e10
                                                _ = v55
                                            }
                                        }
                                    case "\n\t":
                                        v56 := GenS4{F0: v37, F1: ("a" + ("hello" + "hello")), F2: 2}
                                        _ = v56
                                    case "":
                                        var v57 int64 = -9223372036854775808
                                        _ = v57
                                    }
                                    var v58 int32 = -1
                                    v1[fi36] = fi40
                                    _ = v58
                                } else {
                                    v59 := GenS3{F0: 277}
                                    var v60 uint16 = 942
                                    _ = v59
                                    _ = v60
                                }
                            }
                            _ = fi40
                        }
                        fmt.Println(fi36)
                    }
                    var v61 float32 = 0.1
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 3; fi62++ {
                        v63 := GenS3{F0: 2}
                        _ = fi62
                        _ = v63
                    }
                }
            }
            if ((-599 != v28) || (!((v28 * (v28 / ((-1) | 1))) != (v28 - (v28 ^ (v28 ^ v28)))))) {
                v64 := GenS2{F0: 2, F1: 51}
                _ = v64
            }
            switch fi36 {
            case 32767:
                for fi65 := 0; fi65 < 18; fi65++ {
                    switch -217 {
                    case 42:
                        var v66 uint64 = 2
                        _ = len(v2)
                        _ = v66
                    case 128:
                        if ((^(v37)) <= v37) {
                            delete(v1, fi65)
                            v67 := GenS1{F0: ("" + "hello"), F1: (!(int64(v37) != -9223372036854775808)), F2: 2147483648}
                            _ = v67
                        } else {
                            if false {
                                if false {
                                    v68 := GenS2{F0: 614, F1: 249}
                                    v69 := GenS4{F0: (v37 & 7), F1: "", F2: 276}
                                    var v70 int16 = (v28 << ((0) & 63))
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                }
                                delete(v1, 32768)
                                _ = len(v3)
                                fmt.Println(v37)
                                if (2 >= 755) {
                                    v1[(fi36 | 4294967295)] = 365
                                    var v71 uint64 = 0
                                    _ = v71
                                }
                            }
                        }
                    }
                    _ = fi65
                }
            case -386:
                var v72 int64 = 262
                _ = v72
            case 0:
                for fi73 := 0; fi73 < 7; fi73++ {
                    _ = len(v3)
                    _ = fi73
                }
            }
        }
        fmt.Println(fi36, fi36)
        _ = fi36
        _ = v37
    }
    var v74 float32 = -2.25
    for fi75 := 0; fi75 < 10; fi75++ {
        v74 += float32(1e10)
        v76 := GenS5{F0: -598, F1: (-520 > 32767), F2: 32767}
        _ = fi75
        _ = v76
    }
    v3 = append(v3, 24)
    v77 := v1[-9223372036854775808]
    var v78 uint32 = 877
    v79 := GenS5{F0: -36, F1: (1 < 2147483647), F2: int(v28)}
    var v80 bool = false
    for fi81 := 0; fi81 < 19; fi81++ {
        if (v74 >= (v74 / v74)) {
            v82 := v4[32767]
            if ((!(!true)) || (0 < int32(v77))) {
                for fi83 := 0; fi83 < 16; fi83++ {
                    if v80 {
                        v84 := GenS5{F0: -260, F1: (!v80), F2: 352}
                        _ = v84
                    }
                    _ = fi83
                }
                v85 := GenS6{F0: fi81}
                v86 := GenS3{F0: 360}
                _ = v85
                _ = v86
            } else {
                var v87 float32 = float32((0.0 + 3.14159))
                var v88 int = 4294967295
                for fi89 := 0; fi89 < 20; fi89++ {
                    var v90 uint16 = 42
                    _ = len(v3)
                    _ = fi89
                    _ = v90
                }
                _ = v87
                _ = v88
            }
            _ = v82
        }
        v4[fi81] = (v78 - 0)
        if (((770 == uint16(v78)) || (v80 && (-474 < v28))) || v80) {
            switch (^(fi81)) {
            case -9223372036854775808:
                switch "foo bar" {
                case "\n\t":
                    v28 -= (v28 | v28)
                case "foo bar":
                    if v80 {
                        var v91 int32 = 0
                        _ = v91
                    } else {
                        v92 := GenS4{F0: 42, F1: ("hello" + "foo bar"), F2: 0}
                        switch "a" {
                        case "foo bar":
                            if (!v80) {
                                v74 = (1.5 - 1e10)
                            } else {
                                for fi93 := 0; fi93 < 1; fi93++ {
                                    if ((!false) || v80) {
                                        var v94 int8 = -128
                                        switch v77 {
                                        case 32767:
                                            var v95 float64 = 3.14159
                                            _ = v95
                                        case 1:
                                            v80 = (int32(v28) == -466)
                                        case 127:
                                            v96 := GenS7{F0: ("hello" + "\n\t"), F1: uint(v78), F2: uint8(v78)}
                                            _ = v96
                                        }
                                        _ = v94
                                    }
                                    v78 = v78
                                    _ = fi93
                                }
                            }
                        case "\n\t":
                            if false {
                                v97 := GenS8{F0: 255, F1: 0}
                                if (653 >= 32767) {
                                    switch "foo bar" {
                                    case "a":
                                        switch ("hello" + ("\n\t" + "\n\t")) {
                                        case "a":
                                            for fi98 := 0; fi98 < 18; fi98++ {
                                                v1[(fi98 & (fi81 << ((4294967295) & 63)))] = (fi81 + fi81)
                                                _ = len(v3)
                                                _ = fi98
                                            }
                                        case "hello":
                                            v99 := GenS9{F0: 127}
                                            _ = v99
                                        case "\n\t":
                                            var v100 bool = false
                                            _ = v100
                                        default:
                                            if len(v3) > 0 {
                                                v3[0] = v78
                                            }
                                        }
                                    case "\n\t":
                                        if true {
                                            for fi101 := 0; fi101 < 18; fi101++ {
                                                var v102 int = int(v28)
                                                _ = fi101
                                                _ = v102
                                            }
                                            _ = len(v3)
                                            fmt.Println(v74, v78)
                                        } else {
                                            v78 = genFn0()
                                            v103 := GenS10{F0: 65535}
                                            fmt.Println(v77)
                                            _ = v103
                                        }
                                    }
                                    v104 := GenS7{F0: (("foo bar" + (("" + ("" + "")) + "foo bar")) + ("a" + "hello")), F1: 18446744073709551615, F2: 255}
                                    var v105 uint16 = 255
                                    _ = v104
                                    _ = v105
                                } else {
                                    v106 := GenS5{F0: 391, F1: v80, F2: (v77 + 2147483647)}
                                    _ = v106
                                }
                                _ = v97
                            } else {
                                var v107 int16 = 1000
                                _ = v107
                            }
                        default:
                            v108 := GenS11{F0: v80, F1: 1}
                            _ = v108
                        }
                        _ = v92
                    }
                case "hello":
                    _ = len(v4)
                }
            }
            v109 := GenS12{F0: 634, F1: fi81, F2: -1}
            _ = v109
        } else {
            fmt.Println(v77)
            fmt.Println(v74, v78)
        }
        _ = fi81
    }
    v110 := GenS13{F0: 68, F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v74:", v74)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v110
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
