package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int16; F2 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 uint }
type GenS9 struct { F0 int8; F1 uint32; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16; F1 int; F2 uint8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 string; F1 int8 }
func genFn0(p0 bool, p1 string, p2 uint8) int32 {
    v6 := GenS1{F0: 255, F1: 0, F2: "foo bar"}
    var v7 int64 = 4
    var v8 int32 = -785
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v8
}
func genFn1(p0 uint64) uint16 {
    var v9 uint8 = 255
    if (-9223372036854775808 >= 0) {
        for fi10 := 0; fi10 < 2; fi10++ {
            p0 += p0
            v9 = v9
            _ = fi10
        }
        v9 += 42
    } else {
        p0 = 4294967295
        var v11 uint16 = 0
        v12 := GenS1{F0: 183, F1: -118, F2: ("" + "a")}
        _ = v11
        _ = v12
    }
    if (-9223372036854775808 >= 1) {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint32 = uint32(v9)
            _ = fi13
            _ = v14
        }
        v15 := []int{128}
        _ = v15
    }
    _ = p0
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    var v3 uint16 = 2
    var v4 uint16 = 2
    var v5 int64 = 9223372036854775807
    if (755 != 42) {
        if true {
            if (((-2.25 * (1.5 * ((-2.25 / 3.14159) * (1e10 * 1.5)))) - ((0.0 + 3.14159) / (0.1 + 1.5))) != 0.0) {
                var v16 int = -574
                v5 += (-(v5))
                v4 -= v4
                _ = v16
            } else {
                for fi17 := 0; fi17 < 16; fi17++ {
                    v4 = genFn1(0)
                    switch fi17 {
                    case -235:
                        v2 += (v2 - 65535)
                    default:
                        var v18 string = "\n\t"
                        _ = v18
                    }
                    _ = fi17
                }
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := GenS1{F0: 925, F1: 42, F2: (("foo bar" + ("a" + "\n\t")) + "a")}
                    v21 := []int{(fi19 - 42), fi19}
                    fmt.Println(v4)
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            var v22 bool = (0 > 2147483647)
            _ = v22
        }
    } else {
        if (1000 < 0) {
            v2 = v2
            v3 = (v4 + (v4 % ((425) | 1)))
        }
        v23 := v1[-1]
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 int64 = v5
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v5 = (^(v5))
    if true {
        if true {
            v1[-2147483648] = ("foo bar" + ("\n\t" + "\n\t"))
            if false {
                var v26 uint32 = 2147483648
                _ = v26
            }
            var v27 uint64 = 2
            _ = v27
        }
        v28 := []int{}
        _ = v28
    } else {
        if true {
            if (((v5 > (v5 | 0)) || false) || true) {
                if true {
                    v2 += 683
                }
                v2 += (v2 % (((v2 | (v2 / ((v2) | 1)))) | 1))
                v2 = (^(v2))
            }
            var v29 uint32 = 42
            v30 := []int{}
            _ = v29
            _ = v30
        } else {
            v1[1] = "hello"
            for fi31 := 0; fi31 < 20; fi31++ {
                switch "foo bar" {
                case "a":
                    fmt.Println(v5, v3)
                case "\n\t":
                    var v32 uint8 = 59
                    _ = v32
                case "":
                    if false {
                        switch "hello" {
                        case "a":
                            for fi33 := 0; fi33 < 17; fi33++ {
                                v34 := GenS1{F0: (v2 + (^(v2))), F1: 42, F2: "foo bar"}
                                _ = fi33
                                _ = v34
                            }
                        case "\n\t":
                            var v35 uint8 = 2
                            _ = v35
                        case "hello":
                            v2 = 65535
                        default:
                            delete(v1, int(v2))
                        }
                        fmt.Println(v3, fi31, v5)
                    }
                default:
                    fmt.Println(v5, v2, v5)
                }
                if true {
                    if (true && false) {
                        fmt.Println(v4)
                        if (int32(v5) <= genFn0((!false), "foo bar", 42)) {
                            var v36 uint8 = 190
                            _ = v36
                        } else {
                            for fi37 := 0; fi37 < 2; fi37++ {
                                var v38 int32 = -1
                                _ = fi37
                                _ = v38
                            }
                            if false {
                                fmt.Println(v4)
                            } else {
                                if true {
                                    fmt.Println(v3, v5, v3)
                                    if (true && (true && false)) {
                                        v39 := []int{930, -773, (fi31 & (fi31 | (fi31 / ((fi31) | 1)))), (fi31 + fi31)}
                                        _ = v39
                                    } else {
                                        if (42 > 887) {
                                            for fi40 := 0; fi40 < 12; fi40++ {
                                                switch fi40 {
                                                case 128:
                                                    v41 := v1[(fi40 * (fi40 - fi31))]
                                                    _ = v41
                                                default:
                                                    var v42 int = fi40
                                                    _ = v42
                                                }
                                                fmt.Println(fi40)
                                                if false {
                                                    v43 := []int{}
                                                    _ = v43
                                                }
                                                _ = fi40
                                            }
                                            for fi44 := 0; fi44 < 15; fi44++ {
                                                delete(v1, (^(fi44)))
                                                _ = fi44
                                            }
                                            var v45 uint32 = 0
                                            _ = v45
                                        }
                                        for fi46 := 0; fi46 < 17; fi46++ {
                                            var v47 int64 = v5
                                            var v48 string = "hello"
                                            v47 = (v47 + v5)
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                    fmt.Println(v4, v3, v4)
                                }
                                fmt.Println(fi31, v4, fi31)
                                switch (("foo bar" + "hello") + "hello") {
                                case "hello":
                                    switch (fi31 / (((fi31 | (fi31 / ((0) | 1)))) | 1)) {
                                    case 32767:
                                        if (("" + ("a" + "hello")) < "\n\t") {
                                            v5 = v5
                                            if false {
                                                v4 = (v3 >> ((1) & 63))
                                            }
                                            switch 205 {
                                            case 42:
                                                switch ("hello" + (("hello" + "a") + (("" + "hello") + (("hello" + "foo bar") + ("foo bar" + "foo bar"))))) {
                                                case "a":
                                                    var v49 uint32 = 588
                                                    _ = v49
                                                default:
                                                    v50 := GenS1{F0: (v2 >> ((979) & 63)), F1: int8(v5), F2: ""}
                                                    _ = v50
                                                }
                                            }
                                        } else {
                                            for fi51 := 0; fi51 < 9; fi51++ {
                                                _ = len(v1)
                                                _ = fi51
                                            }
                                            v52 := GenS2{F0: (((("\n\t" + "") + "hello") + ((("\n\t" + ("hello" + "a")) + ("" + (("a" + "foo bar") + (("a" + "\n\t") + "\n\t")))) + "foo bar")) + "")}
                                            _ = v52
                                        }
                                    case 143:
                                        for fi53 := 0; fi53 < 11; fi53++ {
                                            var v54 uint64 = 2
                                            var v55 uint32 = 313
                                            fmt.Println(fi53, fi53, v3)
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                    default:
                                        _ = len(v1)
                                    }
                                case "a":
                                    v5 = v5
                                case "foo bar":
                                    for fi56 := 0; fi56 < 18; fi56++ {
                                        v57 := []int{fi31, (fi56 << (((^(v2))) & 63)), 128, 32767}
                                        if (65535 != 4294967295) {
                                            var v58 int32 = -1
                                            _ = v58
                                        }
                                        _ = fi56
                                        _ = v57
                                    }
                                }
                            }
                        }
                        if true {
                            switch int(v5) {
                            case 869:
                                if false {
                                    v59 := []int{(fi31 % (((fi31 | -9223372036854775808)) | 1))}
                                    if false {
                                        for fi60 := 0; fi60 < 19; fi60++ {
                                            switch "\n\t" {
                                            case "foo bar":
                                                if false {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if true {
                                                            if (uint16(v2) == (v4 % ((v4) | 1))) {
                                                                v59 = append(v59, 9223372036854775807)
                                                                var v61 uint8 = 128
                                                                v62 := GenS3{F0: false, F1: 887, F2: 2147483648}
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                switch ("a" + "a") {
                                                                case "a":
                                                                    v63 := GenS2{F0: "a"}
                                                                    _ = v63
                                                                case "hello":
                                                                    v64 := GenS4{F0: 0}
                                                                    _ = v64
                                                                default:
                                                                    v4 = genFn1(300)
                                                                }
                                                                if false {
                                                                    v1[fi31] = "foo bar"
                                                                    _ = len(v59)
                                                                    switch ("\n\t" + "") {
                                                                    case "foo bar":
                                                                        fmt.Println(v5, v2)
                                                                    case "\n\t":
                                                                        switch 0 {
                                                                        case -9223372036854775808:
                                                                            fmt.Println(v3, v3, v5)
                                                                            fmt.Println(v3, v4)
                                                                        default:
                                                                            var v65 uint = v2
                                                                            _ = v65
                                                                        }
                                                                    case "":
                                                                        v3 = (v3 / (((v4 + v4)) | 1))
                                                                    default:
                                                                        if (127 >= 1) {
                                                                            var v66 int64 = (-(v5))
                                                                            var v67 uint = v2
                                                                            var v68 float32 = 0.0
                                                                            _ = v66
                                                                            _ = v67
                                                                            _ = v68
                                                                        } else {
                                                                            var v69 uint8 = 88
                                                                            var v70 int64 = (v5 - 770)
                                                                            var v71 uint16 = 377
                                                                            switch fi60 {
                                                                            case 757:
                                                                                switch 9223372036854775807 {
                                                                                case 42:
                                                                                    fmt.Println(fi60, v5, v71)
                                                                                }
                                                                            case -2147483648:
                                                                                v72 := GenS5{F0: 708}
                                                                                _ = v72
                                                                            default:
                                                                                v69 = v69
                                                                                var v73 string = ""
                                                                                var v74 uint64 = 558
                                                                                v75 := GenS5{F0: (^(v74))}
                                                                                _ = v73
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            _ = v69
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    fmt.Println(v2, fi31)
                                                                    var v76 int16 = 0
                                                                    var v77 int64 = (v5 | (v5 % ((42) | 1)))
                                                                    switch "" {
                                                                    case "hello":
                                                                        var v78 string = ("\n\t" + "\n\t")
                                                                        _ = v78
                                                                    }
                                                                    _ = v76
                                                                    _ = v77
                                                                }
                                                            }
                                                        } else {
                                                            var v79 uint32 = 1
                                                            _ = v79
                                                        }
                                                    default:
                                                        v80 := GenS6{F0: (858 > 1)}
                                                        _ = v80
                                                    }
                                                    if (!(696 <= 2)) {
                                                        if ((1 == 823) && (int8(v3) < 100)) {
                                                            var v81 bool = (127 == 32767)
                                                            v82 := GenS2{F0: "\n\t"}
                                                            _ = v81
                                                            _ = v82
                                                        } else {
                                                            var v83 bool = false
                                                            v4 = (^(v4))
                                                            _ = v83
                                                        }
                                                    } else {
                                                        fmt.Println(v2, v2, fi60)
                                                        if ((v2 > (v2 << ((942) & 63))) && false) {
                                                            switch ("hello" + "\n\t") {
                                                            case "":
                                                                v3 += genFn1(42)
                                                            case "a":
                                                                var v84 float32 = 3.14159
                                                                _ = v84
                                                            default:
                                                                v5 = v5
                                                            }
                                                            v1[fi60] = ("\n\t" + "foo bar")
                                                            var v85 int16 = -32768
                                                            _ = v85
                                                        } else {
                                                            v5 += -899
                                                            v3 += v4
                                                            v86 := GenS7{F0: (v2 * v2)}
                                                            v87 := GenS8{F0: 127, F1: (v2 % ((v2) | 1))}
                                                            _ = v86
                                                            _ = v87
                                                        }
                                                        v59 = append(v59, (fi31 / ((fi31) | 1)))
                                                    }
                                                    fmt.Println(v3, v4)
                                                } else {
                                                    v88 := GenS9{F0: 1, F1: 0, F2: 0}
                                                    var v89 uint32 = 2147483647
                                                    v2 -= (v2 - (v2 | (v2 << (((v2 << ((uint(v89)) & 63))) & 63))))
                                                    _ = v88
                                                    _ = v89
                                                }
                                            case "":
                                                switch fi31 {
                                                case -997:
                                                    v90 := GenS10{F0: (fi31 << (((v2 & v2)) & 63))}
                                                    _ = v90
                                                case 2147483647:
                                                    v4 += v3
                                                }
                                            default:
                                                var v91 uint = (v2 + v2)
                                                _ = v91
                                            }
                                            if false {
                                                v59 = append(v59, (fi31 + fi60))
                                                fmt.Println(v5, v3, fi60)
                                                var v92 uint32 = 439
                                                v1[fi31] = ""
                                                _ = v92
                                            } else {
                                                var v93 uint64 = uint64(v4)
                                                v94 := GenS9{F0: int8(v5), F1: 42, F2: 127}
                                                _ = v93
                                                _ = v94
                                            }
                                            switch "\n\t" {
                                            case "\n\t":
                                                if true {
                                                    if (false && (1e10 < 3.14159)) {
                                                        var v95 uint64 = 517
                                                        _ = v95
                                                    }
                                                    _ = len(v59)
                                                } else {
                                                    switch (fi60 + fi31) {
                                                    case 42:
                                                        var v96 uint64 = 42
                                                        _ = v96
                                                    case 803:
                                                        v5 += v5
                                                    }
                                                    v1[-2147483648] = "hello"
                                                    var v97 uint64 = 4294967295
                                                    _ = v97
                                                }
                                                var v98 bool = false
                                                _ = v98
                                            case "foo bar":
                                                var v99 uint8 = 23
                                                _ = v99
                                            }
                                            fmt.Println(v2, fi31, v4)
                                            _ = fi60
                                        }
                                    } else {
                                        if (!true) {
                                            var v100 uint8 = 42
                                            _ = v100
                                        } else {
                                            for fi101 := 0; fi101 < 1; fi101++ {
                                                var v102 uint64 = 18446744073709551615
                                                var v103 uint8 = 42
                                                var v104 int64 = (v5 | (v5 % (((v5 - 1)) | 1)))
                                                _ = fi101
                                                _ = v102
                                                _ = v103
                                                _ = v104
                                            }
                                            v105 := GenS6{F0: false}
                                            var v106 int8 = 42
                                            _ = v105
                                            _ = v106
                                        }
                                        var v107 string = ""
                                        _ = v107
                                    }
                                    _ = v59
                                }
                            default:
                                v5 -= 42
                            }
                        }
                    } else {
                        v2 -= 945
                        v5 = v5
                        if true {
                            v108 := GenS11{F0: 1, F1: -2147483648, F2: 128}
                            if true {
                                var v109 int64 = v5
                                switch ("a" + ("\n\t" + "\n\t")) {
                                case "":
                                    if (!false) {
                                        v110 := GenS7{F0: v2}
                                        if false {
                                            switch fi31 {
                                            case 128:
                                                var v111 uint8 = 160
                                                _ = v111
                                            case -2147483648:
                                                fmt.Println(v4)
                                            }
                                            if (false && (0 <= 100)) {
                                                v4 = (v3 + (v4 & (v3 ^ (^(v4)))))
                                                v2 += v2
                                            } else {
                                                var v112 uint32 = 1
                                                if (127 != fi31) {
                                                    fmt.Println(fi31, v112)
                                                    v112 = v112
                                                } else {
                                                    if (int8(fi31) > 82) {
                                                        for fi113 := 0; fi113 < 8; fi113++ {
                                                            switch fi113 {
                                                            case 2147483647:
                                                                v114 := GenS9{F0: 0, F1: v112, F2: 0}
                                                                _ = v114
                                                            case 16:
                                                                v115 := GenS3{F0: (!true), F1: 0, F2: v112}
                                                                _ = v115
                                                            default:
                                                                fmt.Println(v109)
                                                            }
                                                            if false {
                                                                v116 := GenS11{F0: 0, F1: (fi31 - fi113), F2: 226}
                                                                var v117 uint16 = (v4 - (v3 - genFn1(4294967295)))
                                                                var v118 uint32 = uint32(v4)
                                                                _ = v116
                                                                _ = v117
                                                                _ = v118
                                                            } else {
                                                                fmt.Println(v2, v5)
                                                                fmt.Println(v3, v112, v112)
                                                                var v119 uint = 561
                                                                var v120 uint16 = v4
                                                                _ = v119
                                                                _ = v120
                                                            }
                                                            var v121 float64 = 3.14159
                                                            _ = fi113
                                                            _ = v121
                                                        }
                                                        v122 := []int{fi31, (fi31 ^ fi31), 0, fi31}
                                                        _ = v122
                                                    }
                                                    for fi123 := 0; fi123 < 1; fi123++ {
                                                        if false {
                                                            if (27 == 127) {
                                                                switch ("hello" + "foo bar") {
                                                                case "\n\t":
                                                                    var v124 uint16 = (^(v3))
                                                                    _ = v124
                                                                }
                                                                var v125 uint32 = 65535
                                                                v126 := GenS12{F0: v3}
                                                                _ = v125
                                                                _ = v126
                                                            }
                                                        }
                                                        v3 = 1
                                                        _ = fi123
                                                    }
                                                }
                                                _ = v112
                                            }
                                            fmt.Println(v3)
                                        } else {
                                            fmt.Println(v109, v4, v109)
                                            v127 := GenS1{F0: v2, F1: int8(fi31), F2: ""}
                                            _ = v127
                                        }
                                        _ = v110
                                    } else {
                                        v5 = int64(fi31)
                                    }
                                default:
                                    if (1 <= 42) {
                                        var v128 bool = true
                                        v129 := GenS7{F0: v2}
                                        _ = v128
                                        _ = v129
                                    }
                                }
                                fmt.Println(v5, v4)
                                _ = v109
                            } else {
                                v5 = 550
                                if true {
                                    fmt.Println(v2, v4)
                                    v4 = (v3 % ((v3) | 1))
                                    delete(v1, 0)
                                } else {
                                    v4 = (v3 * 65535)
                                    v4 += (v3 / ((v3) | 1))
                                    var v130 float32 = (0.0 + 0.0)
                                    _ = v130
                                }
                            }
                            v4 = genFn1(2)
                            _ = v108
                        } else {
                            v2 += (v2 * (v2 & (v2 * v2)))
                            v2 = v2
                        }
                    }
                    fmt.Println(v2)
                    fmt.Println(v3, v4)
                } else {
                    if false {
                        v4 += v3
                    }
                    for fi131 := 0; fi131 < 13; fi131++ {
                        var v132 float64 = ((1.5 - 3.14159) * (3.14159 + -2.25))
                        _ = fi131
                        _ = v132
                    }
                    fmt.Println(v3, v3)
                }
                var v133 float64 = 3.14159
                _ = fi31
                _ = v133
            }
            v134 := GenS5{F0: 18446744073709551615}
            _ = v134
        }
    }
    var v135 bool = true
    v136 := GenS5{F0: 0}
    if (0.1 >= -2.25) {
        v137 := GenS13{F0: ("foo bar" + "a"), F1: 100}
        _ = v137
    }
    if v135 {
        var v138 int16 = 42
        for fi139 := 0; fi139 < 13; fi139++ {
            v135 = false
            v5 -= v5
            _ = fi139
        }
        v140 := GenS13{F0: "hello", F1: 127}
        _ = v138
        _ = v140
    } else {
        var v141 uint16 = v4
        fmt.Println(v5)
        var v142 int8 = 100
        _ = v141
        _ = v142
    }
    for fi143 := 0; fi143 < 15; fi143++ {
        v135 = v135
        for fi144 := 0; fi144 < 11; fi144++ {
            fmt.Println(v135, v3)
            fmt.Println(v135)
            _ = fi144
        }
        v4 = (v4 & genFn1(42))
        _ = fi143
    }
    var v145 bool = (!v135)
    for fi146 := 0; fi146 < 11; fi146++ {
        v147 := []int{(-(fi146)), fi146, (^(fi146)), (fi146 / (((fi146 | 42)) | 1))}
        v148 := GenS13{F0: ("" + ("" + ((("a" + "foo bar") + ("\n\t" + ((("\n\t" + "") + "hello") + ("\n\t" + "foo bar")))) + ""))), F1: 100}
        v5 -= -1
        _ = fi146
        _ = v147
        _ = v148
    }
    fmt.Println(v3, v5)
    switch -875 {
    case -1:
        switch 32767 {
        case 45:
            switch "" {
            case "a":
                v5 = 0
            }
        case -1:
            v135 = true
        default:
            delete(v1, -1)
        }
    case 4294967295:
        if false {
            var v149 uint64 = 414
            var v150 uint8 = uint8(v3)
            var v151 int16 = 0
            _ = v149
            _ = v150
            _ = v151
        }
    case -38:
        var v152 uint16 = (^(v3))
        _ = v152
    default:
        v145 = (true || ((v2 >> (((v2 * (v2 << (((^(v2))) & 63)))) & 63)) < v2))
    }
    if true {
        var v153 int = 75
        _ = v153
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v135:", v135)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v135
    _ = v136
    _ = v145
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
