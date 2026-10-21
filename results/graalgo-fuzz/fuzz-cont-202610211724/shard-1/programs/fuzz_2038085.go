package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 uint8; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS9 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS10 struct { F0 int32 }
func genFn0(p0 uint64) bool {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int64 = -9223372036854775808
        v8 := GenS1{F0: 832}
        if false {
            v9 := []int{(fi6 | 9223372036854775807)}
            v10 := GenS2{F0: (("\n\t" + (("\n\t" + ("" + "foo bar")) + "a")) + "hello")}
            _ = v10
            _ = v9
        } else {
            v7 = v7
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    p0 -= 18446744073709551615
    _ = p0
    return (42 >= 1)
}
func genFn1() int {
    switch "foo bar" {
    case "a":
        var v11 int32 = -2147483648
        _ = v11
    default:
        for fi12 := 0; fi12 < 17; fi12++ {
            fmt.Println(fi12)
            v13 := GenS1{F0: 326}
            switch (-(fi12)) {
            case -9223372036854775808:
                v14 := []int{fi12, fi12}
                _ = v14
            case 128:
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(fi15, fi15)
                    var v16 int = 128
                    v17 := GenS1{F0: 1}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case -402:
                v18 := GenS1{F0: 989}
                _ = v18
            }
            _ = fi12
            _ = v13
        }
    }
    switch "foo bar" {
    case "":
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int = 482
                v21 = (fi19 + fi19)
                var v22 int32 = 477
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := []int{fi19}
            v24 := GenS3{F0: 255, F1: 1, F2: 113}
            _ = fi19
            _ = v23
            _ = v24
        }
    case "\n\t":
        var v25 int16 = 967
        _ = v25
    }
    if (!(1 != 466)) {
        var v26 uint32 = 0
        _ = v26
    } else {
        if true {
            var v27 float64 = ((0.1 + float64((1e10 - float32(0.1)))) / float64(float32(1.5)))
            if (false || true) {
                switch 2147483647 {
                case -1:
                    var v28 uint64 = 42
                    _ = v28
                case 42:
                    fmt.Println(v27)
                case -205:
                    v29 := GenS1{F0: 792}
                    _ = v29
                default:
                    var v30 uint32 = 4294967295
                    _ = v30
                }
            }
            v31 := []int{}
            _ = v27
            _ = v31
        } else {
            var v32 float64 = (3.14159 / (((1.5 / 3.14159) + float64((0.0 + 3.14159))) / (((0.0 - 3.14159) + float64(1.5)) - 0.0)))
            v33 := []int{0}
            var v34 uint64 = 919
            _ = v32
            _ = v33
            _ = v34
        }
    }
    var v35 string = ("" + "hello")
    var v36 int64 = -460
    _ = v35
    _ = v36
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 58
    var v3 string = ("" + "")
    v4 := []int{-497, -9223372036854775808, -86}
    var v5 uint8 = 42
    if false {
        v37 := v1[2147483648]
        var v38 uint64 = 18446744073709551615
        v37 -= -916
        _ = v37
        _ = v38
    }
    v4 = append(v4, 4294967295)
    if len(v4) > 0 {
        v4[0] = 9223372036854775807
    }
    var v39 uint32 = 848
    var v40 int32 = -780
    v2 -= -128
    fmt.Println(v2, v2)
    var v41 int16 = -1
    v2 += 42
    if true {
        var v42 int16 = -32768
        v5 = v5
        if (422 >= v42) {
            if ((float32(1e10) - 0.0) < (1e10 * -2.25)) {
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 int32 = 0
                    _ = fi43
                    _ = v44
                }
                v45 := GenS1{F0: 898}
                v46 := GenS4{F0: 2}
                _ = v45
                _ = v46
            } else {
                v42 += 0
                v1[v39] = 2147483647
            }
            v4 = append(v4, genFn1())
            var v47 uint = 1
            _ = v47
        } else {
            v40 += 1
            if (32767 > int16(v39)) {
                fmt.Println(v5, v42)
                v48 := GenS5{F0: 4294967295, F1: (v5 & 127), F2: (!true)}
                fmt.Println(v2, v3, v40)
                _ = v48
            } else {
                if true {
                    v49 := GenS6{F0: 18446744073709551615}
                    _ = v49
                } else {
                    fmt.Println(v3, v42)
                    var v50 uint = 416
                    _ = v50
                }
            }
        }
        _ = v42
    }
    if (348 >= 1) {
        var v51 int8 = v2
        v52 := GenS7{F0: (v40 >> ((1) & 63))}
        _ = v51
        _ = v52
    } else {
        v39 = v39
        for fi53 := 0; fi53 < 7; fi53++ {
            if false {
                if false {
                    v2 += 91
                    v2 = (v2 >> ((2) & 63))
                    for fi54 := 0; fi54 < 12; fi54++ {
                        var v55 int16 = 1
                        if (!(false || (!false))) {
                            _ = len(v4)
                        }
                        _ = fi54
                        _ = v55
                    }
                }
            }
            switch ((v3 + v3) + v3) {
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "a":
                var v56 float32 = 0.1
                _ = v56
            case "hello":
                switch fi53 {
                case 127:
                    for fi57 := 0; fi57 < 2; fi57++ {
                        fmt.Println(fi57)
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v58 := v1[v39]
                        _ = fi57
                        _ = v58
                    }
                case -577:
                    for fi59 := 0; fi59 < 16; fi59++ {
                        var v60 uint8 = (v5 & 104)
                        if (true || true) {
                            switch fi53 {
                            case -656:
                                v60 = 2
                            }
                        } else {
                            if ((true && (255 <= 1)) || false) {
                                var v61 int64 = 242
                                v39 -= 2
                                var v62 int16 = v41
                                _ = v61
                                _ = v62
                            } else {
                                v2 += v2
                                if (false || ((^(v60)) == (^(v60)))) {
                                    var v63 int16 = (v41 / ((v41) | 1))
                                    switch fi59 {
                                    case 976:
                                        v2 += v2
                                        var v64 uint16 = 550
                                        v4 = append(v4, 114)
                                        _ = v64
                                    case -1:
                                        var v65 bool = true
                                        var v66 uint8 = (v5 | v60)
                                        fmt.Println(v3, v2, v2)
                                        if len(v4) > 0 {
                                            v4[0] = int(v5)
                                        }
                                        switch fi59 {
                                        case -80:
                                            if (v65 && (180 < 369)) {
                                                fmt.Println(v60, v66, v63)
                                                fmt.Println(v66)
                                                var v67 int8 = (v2 / (((^(v2))) | 1))
                                                _ = v67
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        fmt.Println(fi59, v66, v66)
                                        var v68 uint = 0
                                        v5 = v60
                                        if false {
                                            v69 := GenS4{F0: 734}
                                            var v70 uint8 = (v5 * (v66 + v60))
                                            switch fi53 {
                                            case 128:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v71 float32 = float32(1.5)
                                                v1[(v39 ^ v39)] = 4294967295
                                                _ = v71
                                            case -1:
                                                var v72 bool = (v65 || v65)
                                                _ = v72
                                            case 1:
                                                if (32767 != v41) {
                                                    v60 = v70
                                                    v2 -= v2
                                                } else {
                                                    switch ((v3 + ((v3 + ("hello" + ("\n\t" + (v3 + v3)))) + v3)) + "foo bar") {
                                                    case "hello":
                                                        v73 := GenS4{F0: 0}
                                                        _ = v73
                                                    default:
                                                        var v74 uint64 = 42
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            _ = v69
                                            _ = v70
                                        }
                                        _ = v65
                                        _ = v66
                                        _ = v68
                                    case 0:
                                        v75 := v1[1]
                                        _ = v75
                                    default:
                                        v3 = (v3 + v3)
                                        var v76 uint32 = 470
                                        if (!true) {
                                            fmt.Println(v39)
                                            var v77 uint32 = 899
                                            fmt.Println(fi53)
                                            _ = v77
                                        }
                                        fmt.Println(v76, v63)
                                        v78 := GenS2{F0: "\n\t"}
                                        _ = v76
                                        _ = v78
                                    }
                                    var v79 uint = 0
                                    v40 = v40
                                    _ = v63
                                    _ = v79
                                } else {
                                    switch fi59 {
                                    case 42:
                                        var v80 int16 = 0
                                        var v81 float32 = 3.14159
                                        _ = len(v4)
                                        var v82 bool = (65535 > 35)
                                        switch fi59 {
                                        case -582:
                                            if (!true) {
                                                var v83 string = v3
                                                var v84 int = genFn1()
                                                fmt.Println(v82, v83)
                                                _ = v83
                                                _ = v84
                                            } else {
                                                v60 -= 255
                                            }
                                        case 660:
                                            switch fi53 {
                                            case -524:
                                                v85 := GenS8{F0: v60, F1: v80, F2: (^(v39))}
                                                _ = v85
                                            }
                                        default:
                                            var v86 uint16 = 2
                                            _ = v86
                                        }
                                        if v82 {
                                            var v87 uint16 = 683
                                            _ = v87
                                        }
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    default:
                                        if true {
                                            fmt.Println(fi59, v41, v60)
                                        } else {
                                            var v88 int = (fi59 % ((fi53) | 1))
                                            switch v3 {
                                            case "foo bar":
                                                v3 = v3
                                            default:
                                                switch (^(fi59)) {
                                                case -2147483648:
                                                    var v89 bool = (true || (false || false))
                                                    if len(v4) > 0 {
                                                        v4[0] = (-(fi59))
                                                    }
                                                    _ = v89
                                                case 128:
                                                    v90 := GenS9{F0: (v2 % ((v2) | 1)), F1: 65535, F2: (v39 >> ((241) & 63))}
                                                    _ = v90
                                                }
                                            }
                                            var v91 int = v88
                                            switch (^(fi59)) {
                                            case 42:
                                                _ = len(v4)
                                            }
                                            delete(v1, 0)
                                            _ = v88
                                            _ = v91
                                        }
                                    }
                                    v60 = v5
                                }
                                v40 = -2147483648
                            }
                            var v92 uint8 = (v60 >> ((uint(v60)) & 63))
                            var v93 uint = 4294967295
                            _ = v92
                            _ = v93
                        }
                        v40 = (^(v40))
                        if true {
                            var v94 uint64 = 1
                            fmt.Println(v60, v41, v41)
                            _ = v94
                        }
                        switch fi53 {
                        case 128:
                            var v95 int64 = -1
                            _ = v95
                        }
                        var v96 int = (fi59 << ((18446744073709551615) & 63))
                        var v97 int8 = (v2 >> ((48) & 63))
                        _ = fi59
                        _ = v60
                        _ = v96
                        _ = v97
                    }
                case 872:
                    if false {
                        delete(v1, v39)
                        var v98 int8 = (^(v2))
                        var v99 uint = 937
                        _ = v98
                        _ = v99
                    } else {
                        var v100 float64 = 0.1
                        var v101 uint64 = 776
                        if (true && false) {
                            var v102 int8 = (v2 & v2)
                            _ = v102
                        } else {
                            if false {
                                _ = len(v1)
                                fmt.Println(v101, v3, v39)
                            } else {
                                _ = len(v1)
                                var v103 float64 = -2.25
                                _ = v103
                            }
                            var v104 int32 = v40
                            _ = v104
                        }
                        _ = v100
                        _ = v101
                    }
                default:
                    v5 -= v5
                }
            }
            v39 = (^(v39))
            _ = fi53
        }
        for fi105 := 0; fi105 < 3; fi105++ {
            v1[(v39 ^ 42)] = (fi105 + int(v2))
            _ = fi105
        }
    }
    v5 = (^(v5))
    for fi106 := 0; fi106 < 14; fi106++ {
        v39 = v39
        _ = fi106
    }
    for fi107 := 0; fi107 < 18; fi107++ {
        if (255 > 536) {
            v4 = append(v4, (fi107 * (-(fi107))))
            v40 = (-(v40))
            v108 := v1[4294967295]
            _ = v108
        } else {
            for fi109 := 0; fi109 < 7; fi109++ {
                if (42 != 244) {
                    fmt.Println(v39, v39)
                }
                _ = fi109
            }
        }
        var v110 uint64 = 0
        v111 := GenS10{F0: (v40 | (v40 | v40))}
        _ = fi107
        _ = v110
        _ = v111
    }
    switch "" {
    case "foo bar":
        var v112 uint64 = 2
        _ = v112
    case "a":
        v41 -= v41
    default:
        var v113 float32 = 0.0
        _ = v113
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
