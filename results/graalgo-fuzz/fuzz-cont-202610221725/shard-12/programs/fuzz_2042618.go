package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 string }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int32; F1 int64; F2 uint16 }
type GenS9 struct { F0 uint16; F1 bool; F2 bool }
type GenS10 struct { F0 int8; F1 int32; F2 uint }
type GenS11 struct { F0 int; F1 bool }
type GenS12 struct { F0 int8; F1 string }
func genFn0() uint {
    v7 := []int{32768}
    var v8 float64 = -2.25
    var v9 uint16 = 65535
    var v10 float64 = v8
    if (26 < -1) {
        _ = len(v7)
        var v11 int16 = -600
        _ = v11
    } else {
        for fi12 := 0; fi12 < 8; fi12++ {
            v13 := GenS1{F0: (("a" + ("a" + ("\n\t" + "hello"))) + ("\n\t" + "\n\t")), F1: (^(v9)), F2: ("a" + "hello")}
            if (!false) {
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 int8 = -35
                    _ = fi14
                    _ = v15
                }
                fmt.Println(v8, fi12, fi12)
                v16 := GenS1{F0: "hello", F1: v9, F2: "hello"}
                _ = v16
            }
            v10 = float64(-2.25)
            _ = fi12
            _ = v13
        }
        var v17 int64 = 0
        var v18 float64 = v10
        _ = v17
        _ = v18
    }
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1(p0 int32, p1 int) uint8 {
    var v19 uint64 = 4294967295
    if false {
        switch p1 {
        case 9223372036854775807:
            if (!true) {
                if (!(v19 >= (^(v19)))) {
                    v20 := GenS1{F0: "", F1: 32767, F2: "\n\t"}
                    switch (p1 ^ (p1 << ((0) & 63))) {
                    case -1:
                        for fi21 := 0; fi21 < 13; fi21++ {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                fmt.Println(p1, p1, v19)
                                _ = fi22
                            }
                            _ = fi21
                        }
                    case 127:
                        p1 = 838
                    default:
                        switch ("\n\t" + "hello") {
                        case "":
                            var v23 string = "hello"
                            _ = v23
                        case "hello":
                            v19 = v19
                        }
                    }
                    v24 := []int{}
                    _ = v20
                    _ = v24
                } else {
                    if true {
                        v25 := []int{p1, p1}
                        _ = v25
                    }
                    v19 = v19
                    switch p1 {
                    case -2147483648:
                        if false {
                            if true {
                                v26 := []int{}
                                v19 = (v19 + (v19 >> ((1) & 63)))
                                _ = v26
                            } else {
                                if false {
                                    v27 := []int{}
                                    v27 = append(v27, p1)
                                    _ = v27
                                }
                            }
                            p0 = (p0 / (((p0 >> ((genFn0()) & 63))) | 1))
                            var v28 uint8 = 0
                            _ = v28
                        } else {
                            p1 = p1
                            var v29 uint = 42
                            v30 := []int{p1, (p1 - (p1 * p1))}
                            _ = v29
                            _ = v30
                        }
                    case 9223372036854775807:
                        v31 := GenS2{F0: 1, F1: 9223372036854775807}
                        _ = v31
                    default:
                        for fi32 := 0; fi32 < 14; fi32++ {
                            p0 += p0
                            _ = fi32
                        }
                    }
                }
                v33 := []int{-681}
                _ = v33
            }
        case -77:
            v34 := []int{int(v19), 0, -606, -592}
            _ = v34
        case -1:
            fmt.Println(p0, p1)
        }
        var v35 int8 = -54
        _ = v35
    } else {
        var v36 float64 = ((3.14159 - ((3.14159 + ((-2.25 / ((0.1 + -2.25) * -2.25)) - (0.0 * float64(0.0)))) / -2.25)) * (1e10 * ((((0.0 * 3.14159) - 1.5) * (((0.1 * -2.25) + float64((-2.25 * (0.1 * 0.0)))) - -2.25)) + 0.0)))
        var v37 uint64 = 2
        v38 := []int{173, (p1 ^ (p1 >> ((genFn0()) & 63))), p1}
        _ = v36
        _ = v37
        _ = v38
    }
    _ = p0
    _ = p1
    _ = v19
    return uint8(v19)
}

func main() {
    var v1 int = 127
    v2 := []string{"\n\t", ("hello" + (("a" + ("hello" + "\n\t")) + ""))}
    var v3 int64 = 156
    var v4 uint64 = 0
    var v5 float32 = 3.14159
    var v6 int16 = 0
    if (!((4294967295 != 255) || false)) {
        v1 += v1
        var v39 int = (v1 - 127)
        _ = v39
    }
    var v40 float64 = ((0.0 * (((-2.25 + 0.0) + 0.1) / 1.5)) / 1.0)
    var v41 float64 = 1.5
    v42 := GenS3{F0: -2147483648, F1: v4, F2: true}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v43 := GenS4{F0: 18446744073709551615}
    switch v1 {
    case 128:
        var v44 string = "foo bar"
        _ = v44
    case -1:
        switch v1 {
        case 2147483647:
            v45 := GenS1{F0: "", F1: 332, F2: "hello"}
            _ = v45
        case -1:
            for fi46 := 0; fi46 < 15; fi46++ {
                fmt.Println(v5, v1, v5)
                if true {
                    var v47 int32 = -280
                    switch "a" {
                    case "a":
                        var v48 string = "hello"
                        _ = v48
                    case "hello":
                        for fi49 := 0; fi49 < 8; fi49++ {
                            v50 := GenS5{F0: (v4 >> ((160) & 63))}
                            if true {
                                var v51 string = ""
                                _ = v51
                            }
                            _ = fi49
                            _ = v50
                        }
                    }
                    _ = v47
                }
                _ = fi46
            }
        case 4294967295:
            v3 = v3
        }
    default:
        if true {
            for fi52 := 0; fi52 < 15; fi52++ {
                if true {
                    if (v5 <= (((v5 + ((v5 + ((v5 - v5) + 1.5)) - v5)) / v5) + 0.0)) {
                        v1 = (v1 ^ (fi52 * 411))
                        _ = len(v2)
                        var v53 int = (v1 >> ((0) & 63))
                        _ = v53
                    }
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v41, v6, v3)
                    }
                    v5 -= v5
                } else {
                    fmt.Println(v41, v6)
                    if (!false) {
                        for fi54 := 0; fi54 < 12; fi54++ {
                            _ = len(v2)
                            _ = fi54
                        }
                        if ((((-2.25 / (((v5 - 3.14159) - ((v5 * (3.14159 / v5)) + ((v5 - (-2.25 / (v5 * v5))) * v5))) * (0.1 - 3.14159))) - (v5 - -2.25)) * 1.5) >= ((v5 * 1e10) / v5)) {
                            v55 := GenS4{F0: 150}
                            var v56 uint = 749
                            v6 = (-(v6))
                            _ = v55
                            _ = v56
                        } else {
                            if false {
                                v4 += v4
                                switch ("\n\t" + "") {
                                case "":
                                    var v57 float32 = 1.5
                                    _ = v57
                                default:
                                    var v58 string = "foo bar"
                                    _ = v58
                                }
                            }
                            fmt.Println(v4)
                            v4 += 2
                        }
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                    } else {
                        var v59 int8 = -128
                        v3 = (v3 + v3)
                        _ = v59
                    }
                    fmt.Println(v5, v6, v1)
                }
                fmt.Println(v41, v3, v41)
                _ = fi52
            }
            if ((v3 + 1) == 42) {
                var v60 uint8 = 2
                var v61 bool = ("\n\t" >= "\n\t")
                _ = v60
                _ = v61
            } else {
                if true {
                    v62 := GenS6{F0: 1, F1: 789}
                    v63 := GenS2{F0: 2, F1: v3}
                    _ = v62
                    _ = v63
                } else {
                    var v64 string = ("" + ("" + "hello"))
                    _ = v64
                }
                for fi65 := 0; fi65 < 20; fi65++ {
                    if (false || (!((!true) && false))) {
                        if (v41 < v41) {
                            var v66 int16 = v6
                            if len(v2) > 0 {
                                v2[0] = (("" + ("hello" + "\n\t")) + "")
                            }
                            _ = v66
                        }
                        if ((v6 | (v6 << ((65535) & 63))) == v6) {
                            v2 = append(v2, "a")
                        } else {
                            for fi67 := 0; fi67 < 7; fi67++ {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                _ = fi67
                            }
                            var v68 uint32 = 2147483647
                            if (!(1 < (v6 + (v6 / ((v6) | 1))))) {
                                var v69 int8 = -128
                                v68 += 42
                                if (true && true) {
                                    if false {
                                        switch (v1 + fi65) {
                                        case 727:
                                            switch v1 {
                                            case 128:
                                                var v70 string = "\n\t"
                                                _ = v70
                                            default:
                                                v4 = 42
                                            }
                                        case 198:
                                            for fi71 := 0; fi71 < 15; fi71++ {
                                                v6 += -875
                                                var v72 int = (v1 * int(v3))
                                                if false {
                                                    v6 = 0
                                                }
                                                _ = fi71
                                                _ = v72
                                            }
                                        case 1:
                                            if (!(490 > (^(fi65)))) {
                                                if (((v4 / ((v4) | 1)) != v4) || (!false)) {
                                                    for fi73 := 0; fi73 < 4; fi73++ {
                                                        v3 -= 204
                                                        v74 := GenS6{F0: 2, F1: 255}
                                                        if (true && (0 > 295)) {
                                                            var v75 uint64 = (v4 & (v4 * (v4 / ((v4) | 1))))
                                                            _ = v75
                                                        }
                                                        var v76 int = int(v69)
                                                        v76 += fi73
                                                        if (128 <= 0) {
                                                            if (557 == -2147483648) {
                                                                switch fi73 {
                                                                case -9223372036854775808:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        v77 := GenS4{F0: 255}
                                                                        _ = v77
                                                                    case "foo bar":
                                                                        v40 = v40
                                                                    }
                                                                case 794:
                                                                    _ = len(v2)
                                                                case 4294967295:
                                                                    v5 = v5
                                                                }
                                                                v6 = (v6 * 826)
                                                                fmt.Println(v69)
                                                            }
                                                            var v78 uint16 = 122
                                                            _ = v78
                                                        }
                                                        var v79 int16 = v6
                                                        _ = fi73
                                                        _ = v74
                                                        _ = v76
                                                        _ = v79
                                                    }
                                                    var v80 bool = (!false)
                                                    fmt.Println(v1)
                                                    _ = v80
                                                } else {
                                                    var v81 uint32 = (v68 - (^(v68)))
                                                    _ = v81
                                                }
                                                if true {
                                                    v82 := GenS7{F0: 4294967295, F1: ""}
                                                    _ = v82
                                                } else {
                                                    var v83 string = "\n\t"
                                                    _ = v83
                                                }
                                            }
                                        default:
                                            var v84 int32 = 1
                                            _ = v84
                                        }
                                        for fi85 := 0; fi85 < 2; fi85++ {
                                            v4 += (v4 * 409)
                                            var v86 uint = 200
                                            var v87 uint8 = 255
                                            _ = fi85
                                            _ = v86
                                            _ = v87
                                        }
                                        v88 := GenS8{F0: 2147483647, F1: v3, F2: 911}
                                        _ = v88
                                    } else {
                                        for fi89 := 0; fi89 < 11; fi89++ {
                                            var v90 bool = true
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            _ = fi89
                                            _ = v90
                                        }
                                        switch int(v3) {
                                        case -230:
                                            var v91 float64 = 1.5
                                            _ = v91
                                        case -593:
                                            var v92 uint = 18446744073709551615
                                            _ = v92
                                        case -1:
                                            v93 := GenS5{F0: (v4 % (((v4 & (v4 % ((2) | 1)))) | 1))}
                                            _ = v93
                                        }
                                        v94 := GenS9{F0: 2, F1: (!false), F2: false}
                                        _ = v94
                                    }
                                    for fi95 := 0; fi95 < 20; fi95++ {
                                        var v96 uint64 = v4
                                        var v97 int16 = (-(v6))
                                        _ = fi95
                                        _ = v96
                                        _ = v97
                                    }
                                }
                                _ = v69
                            } else {
                                v5 += ((((v5 - float32((v41 + v40))) / v5) / 1.0) * v5)
                            }
                            _ = v68
                        }
                    } else {
                        v1 = fi65
                        var v98 uint32 = 2147483647
                        _ = v98
                    }
                    if true {
                        var v99 string = ("a" + "hello")
                        _ = v99
                    }
                    v100 := GenS10{F0: 0, F1: 1, F2: 817}
                    _ = fi65
                    _ = v100
                }
                var v101 int8 = 1
                _ = v101
            }
            for fi102 := 0; fi102 < 7; fi102++ {
                v41 = -2.25
                _ = fi102
            }
        } else {
            v103 := GenS11{F0: 8, F1: false}
            var v104 uint64 = (v4 + 2)
            _ = v103
            _ = v104
        }
    }
    v105 := GenS12{F0: 100, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
