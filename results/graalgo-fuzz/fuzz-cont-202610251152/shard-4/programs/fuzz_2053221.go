package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int8; F1 uint32; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64; F1 int16 }
type GenS15 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint8) bool {
    p0 = (p0 ^ 65535)
    var v7 uint64 = 2
    var v8 uint = uint(p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint = 725
    var v2 string = ("" + "foo bar")
    var v3 uint64 = 42
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v6 string = ("a" + v2)
    var v9 uint16 = 0
    var v10 int = -2147483648
    v11 := GenS1{F0: 42}
    switch v2 {
    case "foo bar":
        switch v2 {
        case "\n\t":
            v12 := []int{(v10 / ((v10) | 1)), v10, (v10 & 128), 32768}
            _ = v12
        case "foo bar":
            fmt.Println(v9, v9, v1)
        default:
            fmt.Println(v3, v1, v10)
        }
    default:
        fmt.Println(v9)
    }
    switch v2 {
    case "foo bar":
        for fi13 := 0; fi13 < 11; fi13++ {
            v3 = (^(v3))
            _ = len(v5)
            if (2147483648 <= 0) {
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := []int{(v10 >> ((v1) & 63)), 9223372036854775807, v10, (fi13 | (fi14 / ((127) | 1)))}
                    fmt.Println(fi13, v6)
                    var v16 uint16 = v9
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                switch 32767 {
                case -544:
                    v17 := GenS1{F0: -32768}
                    _ = v17
                case -2147483648:
                    v18 := GenS2{F0: -32768, F1: int32(v1), F2: v9}
                    _ = v18
                case -476:
                    v19 := GenS1{F0: 0}
                    _ = v19
                }
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    _ = len(v4)
                    switch ((v2 + "a") + v6) {
                    case "hello":
                        if true {
                            v21 := GenS3{F0: -854}
                            fmt.Println(v1)
                            v6 = v2
                            _ = v21
                        } else {
                            var v22 int16 = 959
                            _ = v22
                        }
                    case "\n\t":
                        v23 := GenS4{F0: 100, F1: 157}
                        _ = v23
                    case "foo bar":
                        switch fi20 {
                        case -338:
                            v6 = "foo bar"
                        }
                    }
                    var v24 int64 = 1
                    _ = fi20
                    _ = v24
                }
                if (-1 != 0) {
                    for fi25 := 0; fi25 < 18; fi25++ {
                        var v26 int16 = 1
                        v27 := []int{}
                        v28 := GenS4{F0: 127, F1: (^(v9))}
                        fmt.Println(v3)
                        if true {
                            fmt.Println(v1, v10)
                            if false {
                                v9 -= v9
                                var v29 float32 = 1.5
                                v2 = "foo bar"
                                _ = v29
                            } else {
                                switch "" {
                                case "a":
                                    v30 := GenS2{F0: (^(v26)), F1: int32(v1), F2: (v9 & 1)}
                                    _ = v30
                                case "foo bar":
                                    fmt.Println(fi25, fi13, v10)
                                case "hello":
                                    var v31 int = (v10 - -9223372036854775808)
                                    _ = v31
                                default:
                                    v10 = fi13
                                }
                                var v32 uint16 = (v9 | (v9 % ((65535) | 1)))
                                var v33 uint8 = uint8(v3)
                                _ = v32
                                _ = v33
                            }
                            var v34 int32 = 42
                            _ = v34
                        } else {
                            if (-1 == v26) {
                                switch (("" + (v2 + (v6 + "a"))) + "foo bar") {
                                case "\n\t":
                                    _ = len(v27)
                                case "foo bar":
                                    if false {
                                        var v35 float32 = 0.1
                                        delete(v4, v6)
                                        v36 := GenS2{F0: (v26 ^ v26), F1: -206, F2: (v9 | (v9 + v9))}
                                        fmt.Println(v35)
                                        _ = v35
                                        _ = v36
                                    }
                                }
                                var v37 int = fi25
                                _ = v37
                            }
                            switch v2 {
                            case "hello":
                                var v38 int16 = 45
                                _ = v38
                            case "foo bar":
                                v6 = "foo bar"
                            }
                        }
                        var v39 uint8 = 42
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v39
                    }
                    v40 := GenS2{F0: -984, F1: 170, F2: (v9 - v9)}
                    switch (v2 + (("" + v2) + ("foo bar" + v2))) {
                    case "a":
                        for fi41 := 0; fi41 < 3; fi41++ {
                            v5[-741] = (fi13 / ((-782) | 1))
                            v1 += v1
                            _ = fi41
                        }
                    case "foo bar":
                        v42 := GenS5{F0: 4294967295, F1: 128, F2: uint32(v9)}
                        _ = v42
                    }
                    _ = v40
                } else {
                    switch v2 {
                    case "hello":
                        for fi43 := 0; fi43 < 5; fi43++ {
                            switch "\n\t" {
                            case "a":
                                if false {
                                    v3 += 4294967295
                                    v44 := GenS4{F0: 127, F1: (^(v9))}
                                    _ = v44
                                } else {
                                    if (180 >= v10) {
                                        v3 += 2
                                        if true {
                                            fmt.Println(v6, v1, v9)
                                            if true {
                                                if (-1 == 9223372036854775807) {
                                                    var v45 uint8 = 255
                                                    if true {
                                                        var v46 uint8 = (v45 ^ 1)
                                                        _ = v46
                                                    } else {
                                                        v5[int(v9)] = (fi13 - fi13)
                                                    }
                                                    _ = v45
                                                }
                                                v47 := []int{}
                                                if (false && true) {
                                                    v48 := GenS1{F0: 1000}
                                                    var v49 uint16 = (v9 + v9)
                                                    _ = v48
                                                    _ = v49
                                                }
                                                if false {
                                                    if false {
                                                        switch (fi43 ^ 0) {
                                                        case -2147483648:
                                                            fmt.Println(fi43)
                                                        case 779:
                                                            if true {
                                                                fmt.Println(v3, v9, v6)
                                                                var v50 uint16 = v9
                                                                _ = v50
                                                            }
                                                            var v51 string = v6
                                                            v2 = v2
                                                            v52 := GenS6{F0: 0, F1: 2, F2: -591}
                                                            v2 = v6
                                                            _ = v51
                                                            _ = v52
                                                        default:
                                                            v53 := GenS6{F0: -128, F1: 892, F2: -503}
                                                            _ = v53
                                                        }
                                                        var v54 uint64 = 18446744073709551615
                                                        v55 := GenS5{F0: v1, F1: 2, F2: 374}
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        var v56 bool = true
                                                        _ = v56
                                                    }
                                                }
                                                _ = v47
                                            } else {
                                                switch "hello" {
                                                case "foo bar":
                                                    delete(v4, v2)
                                                case "a":
                                                    fmt.Println(v10)
                                                default:
                                                    if false {
                                                        switch int(v1) {
                                                        case 32767:
                                                            v57 := GenS4{F0: 127, F1: v9}
                                                            _ = v57
                                                        case 42:
                                                            var v58 int = v10
                                                            _ = v58
                                                        default:
                                                            if true {
                                                                var v59 uint = (^(v1))
                                                                v1 = (v59 & (v59 - (v59 | (^(v59)))))
                                                                v6 = v6
                                                                if ((-2.25 + 0.0) >= float32((0.0 + -2.25))) {
                                                                    var v60 uint64 = 4294967295
                                                                    v60 = (v60 ^ (v60 + (^(v60))))
                                                                    v61 := v5[fi13]
                                                                    _ = len(v4)
                                                                    v6 = v6
                                                                    _ = v60
                                                                    _ = v61
                                                                }
                                                                fmt.Println(v6, v59, v3)
                                                                if (-90 >= 42) {
                                                                    delete(v5, (fi13 * 552))
                                                                    v62 := GenS7{F0: ((v3 << ((v59) & 63)) <= v3)}
                                                                    if (1 == -2147483648) {
                                                                        _ = len(v4)
                                                                        var v63 int8 = -60
                                                                        _ = v63
                                                                    } else {
                                                                        v64 := GenS5{F0: (v1 & 521), F1: 0, F2: 150}
                                                                        v59 = v1
                                                                        v65 := []int{fi13, fi13}
                                                                        _ = v64
                                                                        _ = v65
                                                                    }
                                                                    _ = v62
                                                                }
                                                                v66 := []int{(-(fi43))}
                                                                _ = v59
                                                                _ = v66
                                                            } else {
                                                                var v67 int64 = -1
                                                                if true {
                                                                    if (v67 == v67) {
                                                                        v1 = (v1 >> (((^(v1))) & 63))
                                                                        v68 := GenS8{F0: ((65535 >= v9) || false)}
                                                                        _ = v68
                                                                    }
                                                                }
                                                                var v69 string = v6
                                                                _ = v67
                                                                _ = v69
                                                            }
                                                        }
                                                        v70 := []int{v10, (v10 + fi13), (fi13 - v10)}
                                                        v71 := v4[v6]
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                }
                                                if (v1 == (v1 + (v1 % ((v1) | 1)))) {
                                                    var v72 uint16 = (v9 / ((v9) | 1))
                                                    _ = v72
                                                }
                                                v73 := GenS9{F0: -1, F1: false}
                                                _ = v73
                                            }
                                            switch (v6 + ("a" + v6)) {
                                            case "hello":
                                                var v74 int = 756
                                                _ = v74
                                            default:
                                                if (v2 >= "hello") {
                                                    var v75 float64 = -2.25
                                                    _ = v75
                                                }
                                            }
                                        } else {
                                            v2 = ("a" + v6)
                                        }
                                        v76 := GenS6{F0: -1, F1: 2, F2: 2147483647}
                                        _ = v76
                                    } else {
                                        var v77 float32 = ((float32((((0.1 / ((-2.25 - (float64((float32((1e10 / ((float64(-2.25) - float64(3.14159)) / 1.0))) - (0.0 - 1e10))) * ((((0.1 / 1.0) + 1.5) / 1.0) - 3.14159))) + 0.1)) * 1.5) + (float64(1.5) + 0.0))) + 0.1) + (((1.5 * 3.14159) - 0.1) * float32(float64((-2.25 + 0.1)))))
                                        _ = v77
                                    }
                                    v78 := GenS10{F0: v9}
                                    v79 := GenS11{F0: v3}
                                    _ = v78
                                    _ = v79
                                }
                            default:
                                var v80 float64 = 0.1
                                _ = v80
                            }
                            var v81 int16 = 1
                            _ = fi43
                            _ = v81
                        }
                    }
                }
                switch 481 {
                case 694:
                    if ((float64((0.1 * (0.1 - float32(3.14159)))) + (0.0 * 3.14159)) <= (0.0 + 1e10)) {
                        for fi82 := 0; fi82 < 15; fi82++ {
                            _ = len(v5)
                            if true {
                                v1 = uint(v9)
                                delete(v4, "a")
                            } else {
                                v10 = (-(v10))
                                v2 = v6
                            }
                            if false {
                                if true {
                                    switch (fi82 | fi13) {
                                    case -2147483648:
                                        v83 := GenS9{F0: 1, F1: (v9 != v9)}
                                        var v84 float32 = 3.14159
                                        _ = v83
                                        _ = v84
                                    case 32768:
                                        switch v2 {
                                        case "foo bar":
                                            v85 := GenS12{F0: 65535, F1: 42, F2: 273}
                                            _ = v85
                                        }
                                    case 4294967295:
                                        var v86 int = v10
                                        _ = v86
                                    default:
                                        switch (v2 + ("hello" + "")) {
                                        case "\n\t":
                                            v2 = "\n\t"
                                        default:
                                            v5[127] = 32767
                                        }
                                    }
                                } else {
                                    v6 = v2
                                }
                                v6 = (("hello" + "") + "foo bar")
                                if (true || false) {
                                    var v87 int8 = -128
                                    v6 = v6
                                    _ = v87
                                } else {
                                    if true {
                                        v5[(fi13 & (-(v10)))] = (v10 - v10)
                                        var v88 int = int(v9)
                                        var v89 uint = (v1 << (((v1 * v1)) & 63))
                                        if (!true) {
                                            if (false || (18446744073709551615 < 606)) {
                                                if (2147483647 != 2) {
                                                    v1 += 0
                                                }
                                            } else {
                                                if false {
                                                    v6 = "a"
                                                    if (false || true) {
                                                        var v90 uint = v1
                                                        _ = v90
                                                    } else {
                                                        delete(v5, 126)
                                                        v91 := GenS13{F0: -1}
                                                        if (1.5 > 0.1) {
                                                            v92 := GenS2{F0: 536, F1: int32(v3), F2: (v9 + 417)}
                                                            fmt.Println(fi13)
                                                            v89 = v89
                                                            _ = v92
                                                        } else {
                                                            var v93 uint16 = v9
                                                            var v94 bool = ((-195 != 0) && (v93 != v9))
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        _ = v91
                                                    }
                                                }
                                                v88 += int(v1)
                                                var v95 int = (fi13 / ((0) | 1))
                                                _ = v95
                                            }
                                        } else {
                                            switch "hello" {
                                            case "":
                                                var v96 float32 = 3.14159
                                                _ = v96
                                            case "a":
                                                if true {
                                                    switch -290 {
                                                    case 2147483647:
                                                        v97 := GenS6{F0: -128, F1: 42, F2: 1}
                                                        _ = v97
                                                    }
                                                    v98 := []int{}
                                                    if (!((v6 + "") <= v6)) {
                                                        v1 -= 249
                                                        v9 += 2
                                                        v2 = ("" + v6)
                                                    }
                                                    var v99 int = fi82
                                                    _ = v98
                                                    _ = v99
                                                } else {
                                                    if (1 != 17) {
                                                        fmt.Println(v3, v2)
                                                        v9 = v9
                                                        v9 = v9
                                                    } else {
                                                        var v100 uint32 = 752
                                                        v9 += 0
                                                        var v101 float64 = float64(0.0)
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                    v6 = "a"
                                                    v102 := GenS9{F0: 0, F1: true}
                                                    _ = v102
                                                }
                                            case "hello":
                                                fmt.Println(v88)
                                            default:
                                                v10 -= (fi82 ^ 32767)
                                            }
                                        }
                                        var v103 uint8 = 0
                                        v9 += v9
                                        _ = v103
                                        _ = v88
                                        _ = v89
                                    } else {
                                        v1 = v1
                                    }
                                }
                            } else {
                                var v104 float32 = -2.25
                                switch v2 {
                                case "a":
                                    var v105 uint8 = 17
                                    _ = v105
                                case "foo bar":
                                    fmt.Println(v2)
                                case "":
                                    v5[fi13] = fi13
                                }
                                _ = v104
                            }
                            v2 = ((v6 + "a") + v6)
                            _ = fi82
                        }
                        var v106 bool = false
                        _ = v106
                    }
                default:
                    var v107 int8 = 67
                    _ = v107
                }
            }
            _ = fi13
        }
    }
    for fi108 := 0; fi108 < 18; fi108++ {
        if true {
            var v109 uint16 = v9
            var v110 uint = 191
            v110 = (v1 ^ v110)
            _ = v109
            _ = v110
        } else {
            v111 := []int{(v10 + -886)}
            _ = v111
        }
        var v112 int = fi108
        v113 := GenS12{F0: 870, F1: 30, F2: 418}
        _ = fi108
        _ = v112
        _ = v113
    }
    v114 := GenS14{F0: -278, F1: -1}
    v115 := []int{}
    v3 -= v3
    for fi116 := 0; fi116 < 5; fi116++ {
        var v117 int8 = 127
        _ = fi116
        _ = v117
    }
    var v118 string = "hello"
    v9 = v9
    if true {
        v119 := GenS12{F0: 4294967295, F1: -32768, F2: uint32(v9)}
        _ = v119
    } else {
        var v120 string = (v6 + v118)
        var v121 int64 = 9223372036854775807
        v122 := GenS10{F0: uint16(v1)}
        _ = v120
        _ = v121
        _ = v122
    }
    switch -1 {
    case 244:
        if true {
            v123 := v5[242]
            _ = v123
        } else {
            v6 = (((v6 + v6) + (v2 + "a")) + (v118 + "hello"))
            for fi124 := 0; fi124 < 9; fi124++ {
                var v125 bool = (65535 < 2)
                var v126 uint64 = (^(v3))
                var v127 uint8 = 249
                _ = fi124
                _ = v125
                _ = v126
                _ = v127
            }
        }
    case 32767:
        v128 := GenS6{F0: 1, F1: 1, F2: 42}
        _ = v128
    case 616:
        v3 = v3
    }
    switch v118 {
    case "\n\t":
        if (true || true) {
            if false {
                v6 = (v6 + "hello")
            } else {
                v1 = 1
                v118 = ""
                var v129 int32 = -167
                _ = v129
            }
            var v130 int32 = -2147483648
            fmt.Println(v118)
            _ = v130
        } else {
            v131 := GenS15{F0: -1}
            var v132 int16 = 0
            _ = v131
            _ = v132
        }
    case "foo bar":
        var v133 string = "foo bar"
        _ = v133
    default:
        var v134 uint64 = (v3 | (^(v3)))
        _ = v134
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v115:", len(v115), v115)
    fmt.Println("v118:", v118)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v114
    _ = v115
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
