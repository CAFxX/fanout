package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS9 struct { F0 uint8; F1 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16; F1 uint16; F2 bool }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 bool; F1 int16 }
func genFn0() int8 {
    if true {
        var v7 bool = (1 < 796)
        _ = v7
    } else {
        if true {
            v8 := []int{42, 0}
            if false {
                for fi9 := 0; fi9 < 13; fi9++ {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        fmt.Println(fi10, fi9)
                        var v11 uint32 = 1
                        _ = fi10
                        _ = v11
                    }
                    v12 := GenS1{F0: "hello"}
                    var v13 float64 = ((1e10 - -2.25) / float64(-2.25))
                    _ = fi9
                    _ = v12
                    _ = v13
                }
                var v14 uint16 = 1
                fmt.Println(v14)
                _ = v14
            } else {
                var v15 int64 = -9223372036854775808
                var v16 uint8 = 1
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                for fi18 := 0; fi18 < 6; fi18++ {
                    var v19 uint8 = 231
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 7; fi20++ {
                    if true {
                        if true {
                            v8 = append(v8, fi17)
                        } else {
                            if (49 != 2) {
                                v21 := GenS2{F0: 4294967295, F1: 32767}
                                fmt.Println(fi20, fi17)
                                v8 = append(v8, fi17)
                                _ = v21
                            } else {
                                var v22 int8 = -128
                                _ = v22
                            }
                            v23 := GenS3{F0: ""}
                            _ = v23
                        }
                    } else {
                        v24 := GenS4{F0: -1, F1: false, F2: ((-32768 >= 42) && true)}
                        _ = v24
                    }
                    _ = fi20
                }
                _ = fi17
            }
            _ = v8
        }
        v25 := GenS2{F0: 1, F1: 0}
        _ = v25
    }
    v26 := []int{}
    v27 := GenS1{F0: ""}
    var v28 int = 0
    var v29 bool = false
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    return -77
}
func genFn1() uint16 {
    v30 := GenS5{F0: -568, F1: -118}
    var v31 string = ""
    for fi32 := 0; fi32 < 8; fi32++ {
        v33 := GenS3{F0: "a"}
        _ = fi32
        _ = v33
    }
    for fi34 := 0; fi34 < 8; fi34++ {
        if (!(771 >= 2)) {
            switch "foo bar" {
            case "foo bar":
                if true {
                    for fi35 := 0; fi35 < 20; fi35++ {
                        switch (fi34 - (fi35 >> ((4294967295) & 63))) {
                        case 143:
                            v31 = v31
                        case 127:
                            switch ("hello" + v31) {
                            case "a":
                                if false {
                                    v31 = "\n\t"
                                    fmt.Println(fi34, fi35, v31)
                                }
                            case "\n\t":
                                if true {
                                    if true {
                                        switch v31 {
                                        case "\n\t":
                                            v36 := GenS5{F0: 42, F1: 0}
                                            _ = v36
                                        case "":
                                            switch "\n\t" {
                                            case "hello":
                                                var v37 int8 = 100
                                                _ = v37
                                            case "a":
                                                var v38 float64 = ((3.14159 * 0.1) * (0.1 * 1e10))
                                                _ = v38
                                            case "":
                                                var v39 float64 = 1.5
                                                _ = v39
                                            default:
                                                var v40 int64 = 372
                                                _ = v40
                                            }
                                        default:
                                            v41 := []int{(fi35 ^ fi35), -612, fi34, 42}
                                            _ = v41
                                        }
                                        v31 = (("a" + v31) + "")
                                    }
                                }
                            }
                        default:
                            v31 = v31
                        }
                        fmt.Println(fi34, v31, fi34)
                        var v42 int64 = 42
                        _ = fi35
                        _ = v42
                    }
                    fmt.Println(v31, v31)
                    var v43 bool = true
                    _ = v43
                } else {
                    v44 := GenS2{F0: 563, F1: 858}
                    fmt.Println(fi34, fi34)
                    _ = v44
                }
            case "\n\t":
                var v45 int8 = -1
                _ = v45
            case "hello":
                v31 = "hello"
            default:
                fmt.Println(v31, v31)
            }
            for fi46 := 0; fi46 < 18; fi46++ {
                fmt.Println(v31, fi46, fi34)
                var v47 uint16 = 32767
                v48 := []int{(^(fi46)), (fi34 / ((777) | 1)), 32767}
                _ = fi46
                _ = v47
                _ = v48
            }
            v31 = v31
        }
        switch v31 {
        case "\n\t":
            v49 := GenS6{F0: 127, F1: fi34}
            _ = v49
        }
        v50 := []int{fi34, (fi34 + (fi34 & fi34)), -9223372036854775808, (fi34 % ((fi34) | 1))}
        _ = fi34
        _ = v50
    }
    switch -9223372036854775808 {
    case 128:
        var v51 int = -2147483648
        _ = v51
    }
    _ = v30
    _ = v31
    return 645
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (-168 == 1)
    v3 := []int{0, 32767, 127, -2147483648}
    var v4 int32 = -682
    var v5 int16 = 42
    var v6 int32 = v4
    for fi52 := 0; fi52 < 8; fi52++ {
        var v53 uint16 = genFn1()
        switch "a" {
        case "\n\t":
            fmt.Println(v53, fi52, v6)
        case "foo bar":
            var v54 int64 = 0
            _ = v54
        }
        var v55 uint16 = v53
        _ = fi52
        _ = v53
        _ = v55
    }
    for fi56 := 0; fi56 < 4; fi56++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch fi56 {
        case 1:
            var v57 uint32 = 2147483647
            _ = v57
        case -1:
            v2 = (v2 && v2)
        default:
            var v58 int8 = genFn0()
            _ = v58
        }
        _ = fi56
    }
    var v59 int32 = v4
    switch -9223372036854775808 {
    case 9223372036854775807:
        switch ("hello" + (("hello" + "a") + ("" + ("foo bar" + ("a" + ""))))) {
        case "foo bar":
            for fi60 := 0; fi60 < 8; fi60++ {
                v61 := v1[fi60]
                var v62 uint16 = 32767
                _ = fi60
                _ = v61
                _ = v62
            }
        case "hello":
            _ = len(v3)
        default:
            fmt.Println(v4)
        }
    case 1:
        switch "" {
        case "hello":
            var v63 float32 = (0.0 + 1e10)
            _ = v63
        case "":
            v2 = (!v2)
        case "foo bar":
            for fi64 := 0; fi64 < 8; fi64++ {
                var v65 uint8 = 14
                for fi66 := 0; fi66 < 14; fi66++ {
                    v4 -= (^(v6))
                    v4 -= (v59 - (v59 >> ((65535) & 63)))
                    if v2 {
                        if v2 {
                            if (true || false) {
                                v59 = 1
                                if v2 {
                                    v67 := GenS7{F0: 4294967295, F1: -52}
                                    v59 = (v6 & v6)
                                    _ = v67
                                } else {
                                    _ = len(v1)
                                    if ((((100 <= -127) && v2) || false) || v2) {
                                        var v68 int8 = genFn0()
                                        var v69 uint = 0
                                        var v70 uint32 = 4294967295
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = (fi66 & (fi66 >> ((42) & 63)))
                                        }
                                        if (100 != 1) {
                                            var v71 uint32 = 122
                                            if false {
                                                v72 := GenS7{F0: 0, F1: 42}
                                                switch "a" {
                                                case "hello":
                                                    fmt.Println(v59, v71, fi66)
                                                }
                                                _ = v72
                                            }
                                            _ = v71
                                        }
                                        _ = len(v1)
                                        var v73 bool = (genFn1() == 255)
                                        if (v2 && v73) {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            switch (("a" + ("\n\t" + ("hello" + ""))) + "hello") {
                                            case "":
                                                v1[fi64] = fi64
                                            default:
                                                var v74 bool = true
                                                _ = v74
                                            }
                                        } else {
                                            fmt.Println(v4, v4)
                                        }
                                        _ = v73
                                    }
                                }
                                fmt.Println(v5, v2, fi66)
                            } else {
                                switch (^(fi66)) {
                                case -2147483648:
                                    var v75 string = ("" + "hello")
                                    _ = v75
                                case -709:
                                    v4 -= (-(v6))
                                    v76 := GenS3{F0: "foo bar"}
                                    _ = v76
                                default:
                                    fmt.Println(fi64)
                                }
                                v4 = v4
                            }
                        } else {
                            var v77 int64 = 0
                            _ = v77
                        }
                        var v78 uint64 = 4294967295
                        _ = v78
                    } else {
                        v6 += v6
                        _ = len(v3)
                    }
                    _ = fi66
                }
                _ = fi64
                _ = v65
            }
        }
    default:
        for fi79 := 0; fi79 < 19; fi79++ {
            switch 1 {
            case 1:
                v80 := GenS6{F0: 127, F1: 747}
                _ = v80
            default:
                v81 := GenS5{F0: 0, F1: v5}
                _ = v81
            }
            _ = fi79
        }
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        for fi82 := 0; fi82 < 13; fi82++ {
            fmt.Println(v6, v59)
            v3 = append(v3, fi82)
            v83 := GenS8{F0: 0, F1: (v5 << ((2) & 63)), F2: 127}
            _ = fi82
            _ = v83
        }
    case "hello":
        if true {
            fmt.Println(v2)
            var v84 float64 = 1e10
            v6 += -1
            _ = v84
        } else {
            var v85 int16 = (v5 & v5)
            var v86 int64 = 543
            _ = v85
            _ = v86
        }
    case "a":
        if v2 {
            v5 = v5
        } else {
            if v2 {
                switch 2147483647 {
                case -558:
                    fmt.Println(v4, v5, v2)
                case 2147483647:
                    v5 -= v5
                }
                v5 += 612
                if v2 {
                    for fi87 := 0; fi87 < 11; fi87++ {
                        v1[42] = 127
                        _ = fi87
                    }
                }
            } else {
                for fi88 := 0; fi88 < 1; fi88++ {
                    var v89 int16 = (v5 ^ 1)
                    _ = fi88
                    _ = v89
                }
            }
            var v90 int64 = 0
            var v91 bool = (!v2)
            _ = v90
            _ = v91
        }
    }
    if v2 {
        v1[9223372036854775807] = 4294967295
        var v92 float32 = 0.0
        _ = v92
    } else {
        switch ("hello" + "foo bar") {
        case "":
            switch ("a" + "a") {
            case "a":
                if true {
                    for fi93 := 0; fi93 < 20; fi93++ {
                        v94 := GenS9{F0: 0, F1: 42}
                        _ = fi93
                        _ = v94
                    }
                } else {
                    fmt.Println(v5, v5, v6)
                }
            default:
                for fi95 := 0; fi95 < 6; fi95++ {
                    if (float32((3.14159 * float64(0.0))) >= 1e10) {
                        v96 := GenS3{F0: "foo bar"}
                        switch fi95 {
                        case 42:
                            var v97 uint16 = 32767
                            _ = v97
                        case 1:
                            if (genFn1() >= genFn1()) {
                                v3 = append(v3, (fi95 / (((fi95 << ((42) & 63))) | 1)))
                                for fi98 := 0; fi98 < 5; fi98++ {
                                    v2 = v2
                                    _ = fi98
                                }
                                v6 = v4
                            } else {
                                fmt.Println(v4, v5, fi95)
                                v4 += int32(fi95)
                                v99 := GenS5{F0: (v5 & 1), F1: v5}
                                _ = v99
                            }
                        default:
                            for fi100 := 0; fi100 < 16; fi100++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v101 uint64 = 1
                                _ = fi100
                                _ = v101
                            }
                        }
                        _ = len(v1)
                        _ = v96
                    } else {
                        v102 := GenS10{F0: 0}
                        v2 = ((v5 * 884) < -522)
                        for fi103 := 0; fi103 < 8; fi103++ {
                            if (((42 > 100) && v2) && false) {
                                v104 := GenS3{F0: ((("foo bar" + "hello") + "\n\t") + "\n\t")}
                                v105 := GenS11{F0: ((genFn0() > -119) || (!v2))}
                                _ = v104
                                _ = v105
                            } else {
                                v106 := GenS12{F0: (v5 << ((182) & 63)), F1: 2, F2: false}
                                v107 := GenS4{F0: 348, F1: false, F2: v2}
                                fmt.Println(v5, v59)
                                _ = v106
                                _ = v107
                            }
                            v6 += (v4 / ((-1) | 1))
                            _ = fi103
                        }
                        _ = v102
                    }
                    v59 = (v6 ^ (v59 / ((v6) | 1)))
                    for fi108 := 0; fi108 < 2; fi108++ {
                        fmt.Println(v4, v5, v59)
                        var v109 int8 = 100
                        _ = fi108
                        _ = v109
                    }
                    _ = fi95
                }
            }
        default:
            for fi110 := 0; fi110 < 17; fi110++ {
                fmt.Println(fi110)
                _ = fi110
            }
        }
    }
    v6 -= v59
    fmt.Println(v4)
    for fi111 := 0; fi111 < 16; fi111++ {
        var v112 bool = (!false)
        _ = fi111
        _ = v112
    }
    v59 = 42
    v113 := GenS13{F0: 1}
    if true {
        v2 = (true && false)
        for fi114 := 0; fi114 < 4; fi114++ {
            var v115 int64 = 9223372036854775807
            v116 := GenS13{F0: 255}
            v117 := GenS14{F0: true, F1: v5}
            _ = fi114
            _ = v115
            _ = v116
            _ = v117
        }
        v118 := GenS5{F0: 0, F1: v5}
        _ = v118
    } else {
        v119 := v1[0]
        v120 := v1[v119]
        _ = v119
        _ = v120
    }
    v59 = -1
    v5 = v5
    v121 := GenS1{F0: ""}
    var v122 uint8 = 130
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
