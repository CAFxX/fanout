package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 bool; F1 int; F2 uint32 }
type GenS11 struct { F0 int }
func genFn0(p0 uint32, p1 string, p2 int64) int8 {
    var v4 int32 = -1
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 350}
        if true {
            for fi7 := 0; fi7 < 18; fi7++ {
                fmt.Println(fi5)
                switch fi7 {
                case 1:
                    var v8 int64 = p2
                    _ = v8
                }
                v9 := []int{-2147483648}
                v10 := GenS1{F0: 18446744073709551615}
                _ = fi7
                _ = v10
                _ = v9
            }
            v11 := GenS1{F0: 42}
            if (!true) {
                v12 := []int{(fi5 & fi5), int(p0), fi5, (fi5 >> ((18446744073709551615) & 63))}
                _ = v12
            } else {
                if ((!false) && true) {
                    fmt.Println(p0, p0, p1)
                    p1 = (((p1 + p1) + p1) + ((p1 + (("foo bar" + ("a" + ("\n\t" + p1))) + "\n\t")) + "foo bar"))
                }
                if (272 >= 170) {
                    var v13 bool = (true && (((0.0 >= float64(((3.14159 - (1e10 + 0.1)) / 3.14159))) && false) || true))
                    for fi14 := 0; fi14 < 11; fi14++ {
                        fmt.Println(p1)
                        if v13 {
                            if (!false) {
                                if (0 != 509) {
                                    if (4294967295 == 65535) {
                                        switch fi14 {
                                        case 128:
                                            v13 = ((-2147483648 > fi5) || (!(!false)))
                                            var v15 int64 = -9223372036854775808
                                            fmt.Println(fi5, v13)
                                            _ = v15
                                        default:
                                            var v16 bool = v13
                                            _ = v16
                                        }
                                        p0 += 405
                                        v17 := []int{int(v4), (fi5 >> ((517) & 63)), (fi14 & (fi5 | fi14)), 4294967295}
                                        _ = v17
                                    } else {
                                        switch 1 {
                                        case 32768:
                                            v4 -= -146
                                        case 613:
                                            var v18 int64 = (^(p2))
                                            _ = v18
                                        default:
                                            v19 := GenS2{F0: 255, F1: 102}
                                            _ = v19
                                        }
                                        v20 := []int{}
                                        var v21 int8 = -128
                                        _ = v20
                                        _ = v21
                                    }
                                    if false {
                                        switch p1 {
                                        case "foo bar":
                                            v13 = (!true)
                                        case "":
                                            fmt.Println(fi14)
                                        default:
                                            fmt.Println(fi14, v4, fi14)
                                        }
                                    }
                                } else {
                                    v22 := GenS2{F0: 343, F1: 127}
                                    p0 = p0
                                    v23 := []int{fi5, fi5}
                                    p2 = (p2 | 1)
                                    _ = v22
                                    _ = v23
                                }
                                v24 := GenS1{F0: 1}
                                if v13 {
                                    v4 = v4
                                    switch (p1 + ("" + p1)) {
                                    case "foo bar":
                                        var v25 int32 = v4
                                        _ = v25
                                    }
                                    if v13 {
                                        v26 := []int{42, 42}
                                        p1 = (("hello" + "\n\t") + ("\n\t" + (p1 + p1)))
                                        v13 = (-128 >= -128)
                                        switch "\n\t" {
                                        case "":
                                            var v27 uint8 = 95
                                            _ = v27
                                        }
                                        _ = v26
                                    } else {
                                        var v28 int8 = 63
                                        p2 -= p2
                                        _ = v28
                                    }
                                }
                                _ = v24
                            } else {
                                p1 = "foo bar"
                                var v29 uint = 1
                                var v30 uint64 = 1
                                _ = v29
                                _ = v30
                            }
                            var v31 int64 = p2
                            v32 := GenS1{F0: 273}
                            _ = v31
                            _ = v32
                        }
                        _ = fi14
                    }
                    switch (fi5 | fi5) {
                    case 32767:
                        fmt.Println(v13, fi5, p2)
                    case 0:
                        p1 = p1
                    default:
                        v33 := GenS1{F0: 174}
                        _ = v33
                    }
                    _ = v13
                }
                var v34 uint32 = (p0 / ((2147483647) | 1))
                _ = v34
            }
            _ = v11
        } else {
            var v35 int = fi5
            _ = v35
        }
        fmt.Println(fi5, p0)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -128
}
func genFn1(p0 int, p1 string) uint {
    p1 = (p1 + p1)
    var v36 uint16 = 0
    var v37 int8 = 0
    var v38 float32 = 3.14159
    v37 = (v37 << ((304) & 63))
    _ = p0
    _ = p1
    _ = v36
    _ = v37
    _ = v38
    return 4294967295
}

func main() {
    var v1 int64 = -774
    var v2 bool = (false || true)
    v3 := make(map[int]string)
    if false {
        var v39 int32 = 0
        _ = v39
    } else {
        for fi40 := 0; fi40 < 10; fi40++ {
            v1 += (^(v1))
            _ = fi40
        }
        var v41 uint8 = 223
        _ = v41
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        var v43 string = "a"
        if v2 {
            if (0 != 821) {
                var v44 int32 = -2147483648
                _ = v44
            } else {
                v43 = (v43 + (v43 + ""))
                v45 := GenS3{F0: v1, F1: (fi42 % (((fi42 >> ((255) & 63))) | 1)), F2: -177}
                _ = v45
            }
        } else {
            var v46 int16 = -32768
            v43 = ((v43 + v43) + ("hello" + ""))
            var v47 int32 = int32(v46)
            _ = v46
            _ = v47
        }
        _ = fi42
        _ = v43
    }
    v48 := []int{-340, 42, 42}
    v2 = v2
    v49 := v3[244]
    v50 := GenS1{F0: 764}
    v49 = v49
    switch 0 {
    case 42:
        fmt.Println(v49, v49, v1)
    case 4294967295:
        v1 = (-(v1))
    case 81:
        var v51 bool = (127 == 0)
        _ = v51
    default:
        v52 := GenS4{F0: -453}
        _ = v52
    }
    if (v2 && v2) {
        var v53 float32 = 0.1
        for fi54 := 0; fi54 < 2; fi54++ {
            for fi55 := 0; fi55 < 19; fi55++ {
                v56 := GenS5{F0: 0, F1: 2}
                v57 := GenS3{F0: v1, F1: (fi54 * fi54), F2: -1}
                if true {
                    if (("hello" == (v49 + (v49 + v49))) && v2) {
                        switch (fi55 | (fi54 + fi54)) {
                        case -1:
                            delete(v3, (fi54 << ((1) & 63)))
                        }
                        _ = len(v3)
                        switch v49 {
                        case "foo bar":
                            if (-1 <= 42) {
                                var v58 bool = v2
                                var v59 string = (v49 + v49)
                                _ = v58
                                _ = v59
                            }
                        default:
                            var v60 int16 = -1
                            _ = v60
                        }
                        fmt.Println(v49, v49)
                        if false {
                            if true {
                                fmt.Println(v49, v2)
                                v53 = (v53 + 3.14159)
                            }
                            var v61 int32 = 42
                            _ = v61
                        }
                        v53 -= ((-2.25 / (v53 + v53)) / 1.5)
                    } else {
                        var v62 uint = 18446744073709551615
                        if ((v2 || false) || v2) {
                            switch 128 {
                            case -2147483648:
                                v62 += (v62 << ((v62) & 63))
                            case 32768:
                                var v63 string = "hello"
                                _ = v63
                            case -999:
                                var v64 int64 = (v1 / ((-9223372036854775808) | 1))
                                _ = v64
                            }
                            _ = len(v3)
                        } else {
                            var v65 int32 = -2147483648
                            v2 = (!v2)
                            _ = v65
                        }
                        if ((fi54 ^ (-(fi54))) >= (-(fi54))) {
                            var v66 float32 = (((v53 * float32(1.5)) - (((v53 / 0.1) / 1.0) + (v53 - 0.1))) + 0.1)
                            _ = v66
                        } else {
                            var v67 bool = v2
                            if len(v48) > 0 {
                                v48[0] = (^(fi55))
                            }
                            var v68 bool = v2
                            var v69 uint32 = 1
                            v70 := GenS4{F0: fi54}
                            var v71 float64 = (float64(((v53 / v53) + float32((-2.25 - -2.25)))) * 1e10)
                            _ = v67
                            _ = v68
                            _ = v69
                            _ = v70
                            _ = v71
                        }
                        fmt.Println(v1, fi55)
                        _ = v62
                    }
                    fmt.Println(v1, v2, v49)
                    if (true || v2) {
                        v72 := GenS5{F0: 104, F1: 0}
                        if (v2 || true) {
                            fmt.Println(fi54, v2)
                            if v2 {
                                _ = len(v3)
                                switch fi55 {
                                case 127:
                                    switch v49 {
                                    case "a":
                                        var v73 int64 = (v1 & (v1 - -895))
                                        _ = v73
                                    case "hello":
                                        v74 := GenS1{F0: 4294967295}
                                        _ = v74
                                    default:
                                        switch -554 {
                                        case 9223372036854775807:
                                            var v75 int16 = 42
                                            _ = v75
                                        case 780:
                                            if v2 {
                                                var v76 uint8 = 255
                                                v53 = 0.0
                                                _ = v76
                                            } else {
                                                fmt.Println(v53, v1)
                                                v49 = v49
                                            }
                                        case 32768:
                                            v77 := GenS4{F0: fi55}
                                            _ = v77
                                        default:
                                            switch "foo bar" {
                                            case "":
                                                v53 = float32(1e10)
                                            default:
                                                v78 := v3[-647]
                                                _ = v78
                                            }
                                        }
                                    }
                                case 4294967295:
                                    if false {
                                        var v79 uint = 18
                                        var v80 string = ("hello" + "")
                                        _ = v79
                                        _ = v80
                                    } else {
                                        v49 = v49
                                        v81 := GenS6{F0: 682, F1: v2}
                                        _ = v81
                                    }
                                }
                            }
                            v1 = (v1 | int64(fi54))
                        }
                        v53 -= (1.5 + v53)
                        _ = v72
                    } else {
                        switch 1 {
                        case 4294967295:
                            var v82 uint32 = 0
                            _ = v82
                        case 32767:
                            switch v49 {
                            case "\n\t":
                                var v83 uint32 = 361
                                _ = v83
                            default:
                                fmt.Println(v49)
                            }
                        }
                    }
                } else {
                    var v84 float64 = 0.0
                    v85 := GenS7{F0: 170, F1: v1, F2: 18446744073709551615}
                    _ = len(v48)
                    _ = v84
                    _ = v85
                }
                _ = fi55
                _ = v56
                _ = v57
            }
            if false {
                delete(v3, fi54)
                v49 = v49
                var v86 string = v49
                _ = v86
            } else {
                var v87 float64 = 0.1
                _ = len(v3)
                _ = v87
            }
            _ = fi54
        }
        switch v49 {
        case "hello":
            for fi88 := 0; fi88 < 9; fi88++ {
                if false {
                    fmt.Println(v53, v49)
                    switch -682 {
                    case 32768:
                        for fi89 := 0; fi89 < 4; fi89++ {
                            var v90 uint32 = 0
                            _ = fi89
                            _ = v90
                        }
                    default:
                        v2 = true
                    }
                    switch v49 {
                    case "":
                        var v91 uint64 = 18446744073709551615
                        _ = v91
                    }
                } else {
                    v92 := GenS1{F0: 18446744073709551615}
                    var v93 uint16 = 255
                    _ = v92
                    _ = v93
                }
                _ = fi88
            }
        case "":
            for fi94 := 0; fi94 < 3; fi94++ {
                var v95 uint16 = 1
                if v2 {
                    v49 = ""
                } else {
                    v96 := GenS1{F0: 308}
                    _ = v96
                }
                switch ("hello" + v49) {
                case "hello":
                    v3[int(v95)] = v49
                case "\n\t":
                    v97 := GenS4{F0: fi94}
                    _ = v97
                case "":
                    v98 := GenS3{F0: v1, F1: fi94, F2: 480}
                    _ = v98
                default:
                    switch ("a" + v49) {
                    case "hello":
                        if len(v48) >= 2 {
                            v48 = v48[1:len(v48)]
                        }
                    case "foo bar":
                        v99 := GenS1{F0: 991}
                        _ = v99
                    }
                }
                _ = fi94
                _ = v95
            }
        default:
            v100 := GenS6{F0: 93, F1: v2}
            _ = v100
        }
        _ = v53
    } else {
        var v101 uint = 4294967295
        _ = v101
    }
    v102 := GenS2{F0: 32767, F1: genFn0(0, v49, (v1 >> ((57) & 63)))}
    v49 = v49
    v1 += v1
    if (93 > 1) {
        switch 4294967295 {
        case 127:
            var v103 int8 = 0
            _ = v103
        case 4294967295:
            for fi104 := 0; fi104 < 7; fi104++ {
                v105 := GenS3{F0: v1, F1: (fi104 >> ((4294967295) & 63)), F2: 1}
                _ = fi104
                _ = v105
            }
        }
    } else {
        switch 32768 {
        case -9223372036854775808:
            v1 = 42
        case 0:
            if (100 > 42) {
                for fi106 := 0; fi106 < 11; fi106++ {
                    v1 += (v1 % (((v1 & v1)) | 1))
                    var v107 int16 = 579
                    v1 = v1
                    _ = fi106
                    _ = v107
                }
                v1 -= v1
                fmt.Println(v49, v1)
            }
        case 83:
            for fi108 := 0; fi108 < 12; fi108++ {
                v3[-9223372036854775808] = "\n\t"
                if len(v48) > 0 {
                    v48[0] = fi108
                }
                _ = fi108
            }
        default:
            v109 := GenS7{F0: 0, F1: (-(v1)), F2: 0}
            _ = v109
        }
        for fi110 := 0; fi110 < 18; fi110++ {
            var v111 float64 = float64(-2.25)
            v2 = true
            var v112 uint32 = 2147483647
            _ = fi110
            _ = v111
            _ = v112
        }
        if (1 > 255) {
            fmt.Println(v1)
            v49 = v49
        } else {
            switch v49 {
            case "foo bar":
                if (2147483647 > -2147483648) {
                    switch "hello" {
                    case "hello":
                        v2 = v2
                    case "":
                        v113 := GenS8{F0: 2}
                        _ = v113
                    }
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        for fi114 := 0; fi114 < 20; fi114++ {
                            v1 = v1
                            _ = len(v3)
                            var v115 float64 = (0.0 * 0.1)
                            _ = fi114
                            _ = v115
                        }
                    case 0:
                        var v116 float64 = 1.5
                        _ = v116
                    }
                    fmt.Println(v1, v2, v2)
                } else {
                    v117 := GenS1{F0: 42}
                    _ = v117
                }
            case "a":
                v118 := v3[-422]
                _ = v118
            }
            v119 := GenS9{F0: 1}
            if v2 {
                if (float32((0.1 * 0.0)) >= (1.5 - 1e10)) {
                    if len(v48) >= 2 {
                        v48 = v48[1:len(v48)]
                    }
                    var v120 string = "\n\t"
                    var v121 uint16 = 32767
                    _ = v120
                    _ = v121
                }
            } else {
                v122 := GenS9{F0: 2}
                v1 -= v1
                _ = v122
            }
            _ = v119
        }
    }
    var v123 string = v49
    v124 := GenS10{F0: (0 > 1), F1: 0, F2: 0}
    for fi125 := 0; fi125 < 10; fi125++ {
        v126 := GenS11{F0: fi125}
        v123 = ((v49 + (v123 + (v123 + "hello"))) + v123)
        v1 -= v1
        _ = fi125
        _ = v126
    }
    switch 9223372036854775807 {
    case -170:
        v49 = "\n\t"
    case -2147483648:
        for fi127 := 0; fi127 < 14; fi127++ {
            var v128 int64 = (v1 - v1)
            fmt.Println(v49)
            var v129 int64 = v128
            _ = fi127
            _ = v128
            _ = v129
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v102
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
}
