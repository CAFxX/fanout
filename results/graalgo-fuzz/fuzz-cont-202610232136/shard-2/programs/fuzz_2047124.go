package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint32; F2 int8 }
type GenS8 struct { F0 int8; F1 uint32; F2 int32 }
type GenS9 struct { F0 int8; F1 uint32; F2 int }
type GenS10 struct { F0 uint; F1 string; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint8; F2 uint16 }
func genFn0(p0 uint, p1 bool) uint16 {
    v6 := []int{-869, 128, 0, 4294967295}
    v7 := GenS1{F0: 0, F1: (p0 ^ p0)}
    v8 := GenS1{F0: -594, F1: (p0 | p0)}
    if p1 {
        for fi9 := 0; fi9 < 6; fi9++ {
            if true {
                var v10 float32 = 1e10
                switch "\n\t" {
                case "":
                    if (-648 == 0) {
                        var v11 int32 = -138
                        var v12 string = "hello"
                        for fi13 := 0; fi13 < 17; fi13++ {
                            v12 = v12
                            var v14 float64 = -2.25
                            _ = fi13
                            _ = v14
                        }
                        _ = v11
                        _ = v12
                    }
                default:
                    var v15 float64 = (float64((v10 + (0.1 + (v10 + v10)))) - 1e10)
                    _ = v15
                }
                _ = v10
            } else {
                p1 = (p1 || (-999 > 0))
                var v16 uint64 = 703
                var v17 uint16 = 65535
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: 1000, F1: (p0 / ((p0) | 1))}
            _ = fi9
            _ = v18
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            if false {
                for fi20 := 0; fi20 < 20; fi20++ {
                    if len(v6) > 0 {
                        v6[0] = fi20
                    }
                    _ = fi20
                }
            }
            v21 := GenS2{F0: 0, F1: 100, F2: -401}
            v6 = append(v6, fi19)
            _ = fi19
            _ = v21
        }
        p0 -= (p0 ^ (p0 / (((p0 ^ (p0 ^ p0))) | 1)))
    }
    var v22 uint = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v22
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int = 1
    var v2 int = v1
    var v3 float32 = float32(0.0)
    v4 := make(map[int]int)
    var v5 uint32 = 2147483647
    if true {
        v23 := []int{127, 32768, (v1 + 42)}
        for fi24 := 0; fi24 < 9; fi24++ {
            fmt.Println(fi24, v2, fi24)
            _ = fi24
        }
        v5 -= v5
        _ = v23
    }
    v25 := GenS3{F0: "foo bar", F1: 2147483647, F2: ""}
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS3{F0: ("\n\t" + "a"), F1: 2147483647, F2: ("hello" + "hello")}
        _ = fi26
        _ = v27
    }
    if (255 < uint8(v5)) {
        v28 := []int{v1, 0, (^(v1))}
        _ = v28
    }
    v5 -= (v5 | (^(v5)))
    var v29 uint64 = 949
    v30 := GenS4{F0: 42, F1: 4294967295}
    v31 := GenS3{F0: ("a" + "a"), F1: -2147483648, F2: ("\n\t" + "hello")}
    v32 := GenS5{F0: "foo bar"}
    switch (^(v1)) {
    case 127:
        v33 := GenS6{F0: 42, F1: "foo bar", F2: 1}
        _ = v33
    case 507:
        switch "\n\t" {
        case "foo bar":
            v2 -= (-(v1))
        case "":
            for fi34 := 0; fi34 < 15; fi34++ {
                var v35 int64 = 885
                _ = len(v4)
                _ = fi34
                _ = v35
            }
        case "\n\t":
            switch "" {
            case "\n\t":
                if true {
                    var v36 float64 = 0.1
                    _ = v36
                } else {
                    if false {
                        v37 := []int{(^(v1)), v2, 127}
                        if true {
                            v38 := GenS3{F0: "\n\t", F1: -815, F2: ("a" + "foo bar")}
                            switch -1 {
                            case -33:
                                if false {
                                    var v39 float64 = 1.5
                                    v29 = (v29 ^ (v29 + (v29 / ((v29) | 1))))
                                    _ = v39
                                } else {
                                    v40 := GenS7{F0: 0, F1: (v5 - (^(v5))), F2: 42}
                                    switch (v2 & v1) {
                                    case 0:
                                        var v41 uint16 = 42
                                        _ = v41
                                    case 32767:
                                        if (((127 < 2) && ((v5 % (((v5 << ((1) & 63))) | 1)) == v5)) || ((v5 + v5) >= (v5 / ((42) | 1)))) {
                                            var v42 int16 = -924
                                            var v43 int64 = -518
                                            for fi44 := 0; fi44 < 2; fi44++ {
                                                v43 = 1
                                                fmt.Println(v2, v43)
                                                _ = fi44
                                            }
                                            _ = v42
                                            _ = v43
                                        } else {
                                            var v45 int = 780
                                            _ = v45
                                        }
                                    }
                                    _ = v40
                                }
                            case 4294967295:
                                var v46 int16 = -590
                                _ = v46
                            default:
                                v47 := v4[(v1 / ((v1) | 1))]
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 19; fi48++ {
                                switch (v1 - (v1 * -9223372036854775808)) {
                                case -9223372036854775808:
                                    for fi49 := 0; fi49 < 9; fi49++ {
                                        v50 := GenS1{F0: 1000, F1: 942}
                                        v37 = append(v37, 1)
                                        _ = fi49
                                        _ = v50
                                    }
                                }
                                v37 = append(v37, v1)
                                var v51 string = ""
                                _ = fi48
                                _ = v51
                            }
                            _ = v38
                        }
                        var v52 bool = (!true)
                        _ = v37
                        _ = v52
                    } else {
                        v53 := GenS4{F0: 846, F1: 255}
                        v29 -= 533
                        for fi54 := 0; fi54 < 20; fi54++ {
                            v5 += (^(v5))
                            switch fi54 {
                            case 32768:
                                if false {
                                    v55 := []int{}
                                    var v56 int32 = -1
                                    _ = v55
                                    _ = v56
                                } else {
                                    switch "" {
                                    case "\n\t":
                                        for fi57 := 0; fi57 < 16; fi57++ {
                                            _ = len(v4)
                                            if true {
                                                var v58 int64 = -98
                                                if (!(42 == -32768)) {
                                                    v59 := GenS8{F0: -7, F1: 35, F2: 0}
                                                    var v60 float32 = 1.5
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    v61 := []int{(v1 + (v2 - 32768)), 597, (fi57 + (fi54 | -558))}
                                                    _ = v61
                                                }
                                                v62 := GenS5{F0: "hello"}
                                                _ = v58
                                                _ = v62
                                            }
                                            _ = fi57
                                        }
                                    case "hello":
                                        fmt.Println(v2, v5, fi54)
                                    case "a":
                                        switch 1 {
                                        case -2147483648:
                                            v3 += -2.25
                                        case -768:
                                            switch "foo bar" {
                                            case "a":
                                                if (float32(0.1) > float32(1e10)) {
                                                    var v63 string = "hello"
                                                    _ = v63
                                                } else {
                                                    for fi64 := 0; fi64 < 11; fi64++ {
                                                        switch v1 {
                                                        case -330:
                                                            v65 := []int{(fi64 / (((fi54 >> ((1) & 63))) | 1)), fi64, (fi54 + -383), fi64}
                                                            _ = v65
                                                        }
                                                        fmt.Println(v29)
                                                        var v66 uint64 = 449
                                                        _ = fi64
                                                        _ = v66
                                                    }
                                                    _ = len(v4)
                                                }
                                            case "hello":
                                                var v67 uint8 = 0
                                                _ = v67
                                            default:
                                                if false {
                                                    var v68 int64 = -1
                                                    var v69 int64 = (v68 + (-(v68)))
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                        default:
                                            if true {
                                                v3 = ((1e10 - 0.0) / 0.1)
                                            }
                                        }
                                    default:
                                        var v70 float64 = (0.0 + -2.25)
                                        _ = v70
                                    }
                                    var v71 uint64 = (^(v29))
                                    var v72 string = "hello"
                                    _ = v71
                                    _ = v72
                                }
                            default:
                                v3 = (float32(1e10) / ((float32((0.1 - 1.5)) - v3) - v3))
                            }
                            var v73 uint = 250
                            _ = fi54
                            _ = v73
                        }
                        _ = v53
                    }
                }
            case "foo bar":
                v2 = -2147483648
            case "a":
                var v74 int = -959
                _ = v74
            default:
                for fi75 := 0; fi75 < 14; fi75++ {
                    var v76 uint32 = v5
                    v77 := GenS9{F0: -87, F1: (v76 >> ((65535) & 63)), F2: 0}
                    var v78 int = 32767
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                }
            }
        default:
            if (42 != (v29 >> ((510) & 63))) {
                for fi79 := 0; fi79 < 10; fi79++ {
                    var v80 uint64 = 1
                    if true {
                        if false {
                            for fi81 := 0; fi81 < 12; fi81++ {
                                var v82 uint = 18446744073709551615
                                if true {
                                    if (-9223372036854775808 < -9223372036854775808) {
                                        var v83 int32 = 321
                                        var v84 int = v2
                                        _ = v83
                                        _ = v84
                                    }
                                    if false {
                                        fmt.Println(v29)
                                    }
                                }
                                switch "a" {
                                case "hello":
                                    if false {
                                        v3 = ((v3 + float32(1e10)) * (0.1 / float32(3.14159)))
                                    }
                                case "a":
                                    v29 -= uint64(v82)
                                default:
                                    if (!true) {
                                        delete(v4, (fi79 << (((v82 - v82)) & 63)))
                                    } else {
                                        v85 := []int{}
                                        _ = v85
                                    }
                                }
                                _ = fi81
                                _ = v82
                            }
                            var v86 uint = 255
                            for fi87 := 0; fi87 < 7; fi87++ {
                                v88 := []int{}
                                var v89 string = (("\n\t" + "a") + ("\n\t" + "a"))
                                _ = fi87
                                _ = v88
                                _ = v89
                            }
                            _ = v86
                        }
                        for fi90 := 0; fi90 < 6; fi90++ {
                            fmt.Println(v80, v80)
                            v1 = v2
                            _ = fi90
                        }
                        var v91 int32 = 953
                        _ = v91
                    }
                    if true {
                        if false {
                            v92 := GenS9{F0: 121, F1: 4294967295, F2: 443}
                            var v93 bool = (((true && true) && ((v29 / ((592) | 1)) >= 1)) && (-97 <= 81))
                            fmt.Println(v80, v93)
                            _ = v92
                            _ = v93
                        } else {
                            if false {
                                var v94 float32 = (v3 - v3)
                                var v95 float32 = v94
                                _ = v94
                                _ = v95
                            } else {
                                v3 = v3
                                v96 := GenS10{F0: 289, F1: "a", F2: v5}
                                _ = v96
                            }
                            for fi97 := 0; fi97 < 17; fi97++ {
                                v5 -= 4294967295
                                if false {
                                    var v98 float32 = v3
                                    _ = v98
                                } else {
                                    fmt.Println(v29, v3, v3)
                                    v2 = fi97
                                    fmt.Println(v80, fi79)
                                }
                                v3 -= 3.14159
                                _ = fi97
                            }
                            switch (v1 >> ((563) & 63)) {
                            case 917:
                                var v99 int32 = 42
                                _ = v99
                            }
                        }
                        v100 := GenS11{F0: true, F1: 127, F2: 255}
                        _ = v100
                    }
                    _ = fi79
                    _ = v80
                }
                var v101 uint64 = 4294967295
                if false {
                    var v102 int = (v2 | v2)
                    _ = len(v4)
                    _ = v102
                } else {
                    for fi103 := 0; fi103 < 3; fi103++ {
                        var v104 int8 = 7
                        _ = fi103
                        _ = v104
                    }
                    v105 := []int{v1, 2147483647, v2, -9223372036854775808}
                    var v106 uint16 = 255
                    _ = v105
                    _ = v106
                }
                _ = v101
            } else {
                if true {
                    for fi107 := 0; fi107 < 7; fi107++ {
                        v5 += (v5 + v5)
                        switch "\n\t" {
                        case "":
                            fmt.Println(v2, v29, v1)
                        }
                        _ = fi107
                    }
                } else {
                    var v108 uint64 = v29
                    v109 := []int{v2, 32768, v1}
                    fmt.Println(v5)
                    _ = v108
                    _ = v109
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
