package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 int }
type GenS9 struct { F0 int32; F1 uint8; F2 int32 }
type GenS10 struct { F0 int16; F1 bool }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8; F1 int8; F2 int8 }
type GenS14 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS15 struct { F0 int8; F1 uint16; F2 int8 }
type GenS16 struct { F0 int; F1 uint64; F2 uint }
func genFn0(p0 string, p1 uint) int16 {
    var v7 int = 0
    v8 := GenS1{F0: -88, F1: 2}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 727
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-97, 127, 475, 42}
    var v3 uint64 = 0
    var v4 uint32 = 233
    var v5 uint64 = 0
    v6 := []string{"hello", "hello", "hello"}
    v3 += (v3 - v3)
    v5 = v5
    v5 = (^(v5))
    v3 -= v5
    if (-44 <= 1) {
        var v9 float32 = (float32((0.0 + 3.14159)) / ((1e10 + 3.14159) / -2.25))
        v9 = (v9 - ((3.14159 / 1.5) * (v9 / v9)))
        v10 := GenS2{F0: -691, F1: (true && true)}
        _ = v10
        _ = v9
    }
    if (!((-2147483648 >= 0) || false)) {
        v4 -= (^(v4))
        v5 += v3
        v4 = (v4 ^ (v4 + 980))
    }
    switch 32768 {
    case 1:
        fmt.Println(v4, v5)
    case 128:
        fmt.Println(v4, v3, v3)
    case 812:
        for fi11 := 0; fi11 < 9; fi11++ {
            v3 += v3
            _ = fi11
        }
    default:
        var v12 float32 = 1.5
        _ = v12
    }
    switch ("\n\t" + ("a" + "a")) {
    case "foo bar":
        if false {
            v4 = (v4 << ((uint(v4)) & 63))
        } else {
            v3 += v3
        }
    case "":
        var v13 int32 = 2147483647
        _ = v13
    case "a":
        _ = len(v2)
    default:
        if true {
            fmt.Println(v3, v5, v5)
            v5 = v5
            var v14 uint8 = 248
            _ = v14
        } else {
            var v15 int32 = 2147483647
            var v16 uint32 = v4
            for fi17 := 0; fi17 < 9; fi17++ {
                v16 = 2147483648
                for fi18 := 0; fi18 < 18; fi18++ {
                    v16 = v16
                    fmt.Println(v5)
                    fmt.Println(v16)
                    _ = fi18
                }
                v15 += v15
                _ = fi17
            }
            _ = v15
            _ = v16
        }
    }
    var v19 int16 = 32767
    if false {
        if false {
            fmt.Println(v3, v5, v5)
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 uint16 = 10
                if false {
                    _ = len(v1)
                }
                var v22 int8 = 127
                _ = fi20
                _ = v21
                _ = v22
            }
            var v23 int64 = 9223372036854775807
            _ = v23
        } else {
            if true {
                for fi24 := 0; fi24 < 15; fi24++ {
                    var v25 float64 = -2.25
                    _ = fi24
                    _ = v25
                }
                v26 := GenS2{F0: 9223372036854775807, F1: true}
                fmt.Println(v4)
                _ = v26
            }
            v19 = (v19 >> ((42) & 63))
            v27 := GenS1{F0: 1, F1: uint32(v3)}
            _ = v27
        }
        delete(v1, (v4 * v4))
        v3 = (v5 - (v5 & (v3 & (v3 + (v3 & uint64(v4))))))
    } else {
        _ = len(v1)
        fmt.Println(v19)
        if false {
            switch "a" {
            case "":
                v4 += (^(v4))
            case "foo bar":
                v5 = uint64(v4)
            default:
                if true {
                    var v28 int32 = -531
                    _ = v28
                }
            }
            var v29 int8 = 0
            var v30 int32 = 1
            _ = v29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        if false {
            v3 = (v3 << ((0) & 63))
            if ((v3 / ((v3) | 1)) != v5) {
                var v32 float32 = -2.25
                _ = v32
            }
        } else {
            for fi33 := 0; fi33 < 11; fi33++ {
                fmt.Println(v19, fi33, fi33)
                _ = fi33
            }
            if false {
                for fi34 := 0; fi34 < 7; fi34++ {
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            v19 = genFn0(((("hello" + "foo bar") + "\n\t") + ("a" + "\n\t")), 42)
                        case "\n\t":
                            v19 = (v19 * -32768)
                        case "":
                            var v35 int16 = (^(v19))
                            _ = v35
                        }
                    } else {
                        switch (fi34 & (fi34 >> ((624) & 63))) {
                        case 0:
                            if (126 >= -606) {
                                switch 0 {
                                case -377:
                                    var v36 string = "a"
                                    _ = v36
                                case 42:
                                    var v37 uint8 = 73
                                    _ = v37
                                case 984:
                                    var v38 uint16 = 1
                                    _ = v38
                                default:
                                    fmt.Println(v19)
                                }
                                v19 += (v19 ^ (v19 >> ((1) & 63)))
                            }
                        default:
                            switch int(v4) {
                            case 2147483647:
                                v39 := GenS3{F0: 0, F1: 0, F2: (^(v4))}
                                _ = v39
                            case 0:
                                fmt.Println(v5, fi34, v3)
                            default:
                                var v40 int64 = 807
                                _ = v40
                            }
                        }
                        v41 := GenS4{F0: fi34, F1: 100}
                        if ((true || (false && true)) || true) {
                            var v42 uint64 = v3
                            switch fi34 {
                            case 32768:
                                if (false && true) {
                                    v43 := GenS5{F0: 4, F1: 0, F2: genFn0("hello", 2)}
                                    var v44 int = fi31
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 uint64 = (v5 & v42)
                                    _ = v45
                                }
                            default:
                                var v46 uint = 42
                                _ = v46
                            }
                            _ = v42
                        } else {
                            v47 := GenS1{F0: -76, F1: (v4 - 2147483648)}
                            var v48 int = 311
                            if (-1 <= 8) {
                                var v49 int8 = 100
                                _ = v49
                            } else {
                                var v50 int64 = -1
                                _ = v50
                            }
                            switch (("foo bar" + "a") + "\n\t") {
                            case "foo bar":
                                v51 := GenS6{F0: 2147483647}
                                _ = v51
                            case "":
                                _ = len(v2)
                            }
                            var v52 uint = 255
                            _ = v47
                            _ = v48
                            _ = v52
                        }
                        _ = v41
                    }
                    if (127 <= -1) {
                        if true {
                            if false {
                                v3 -= v3
                                var v53 float64 = 3.14159
                                if true {
                                    var v54 uint32 = (v4 + v4)
                                    if true {
                                        v6 = append(v6, "\n\t")
                                        var v55 uint8 = uint8(v54)
                                        _ = v55
                                    } else {
                                        fmt.Println(v3)
                                        var v56 float64 = -2.25
                                        _ = len(v2)
                                        _ = v56
                                    }
                                    v57 := GenS7{F0: (fi31 << ((4294967295) & 63))}
                                    if false {
                                        if false {
                                            v58 := v1[1]
                                            _ = v58
                                        } else {
                                            if (((v4 & v4) < v54) && (2 >= 2)) {
                                                v1[0] = (-(fi31))
                                                var v59 string = "a"
                                                fmt.Println(fi34, v4)
                                                v59 = ((((v59 + v59) + v59) + v59) + v59)
                                                _ = len(v1)
                                                if (3.14159 == 0.0) {
                                                    v3 = (v3 & uint64(v4))
                                                } else {
                                                    var v60 uint64 = (^(v5))
                                                    _ = v60
                                                }
                                                _ = v59
                                            } else {
                                                v61 := GenS3{F0: 2, F1: 127, F2: (v4 & 2)}
                                                _ = v61
                                            }
                                            v62 := GenS2{F0: 682, F1: true}
                                            fmt.Println(v53)
                                            _ = v62
                                        }
                                    }
                                    _ = v54
                                    _ = v57
                                } else {
                                    var v63 int = fi34
                                    _ = len(v1)
                                    _ = v63
                                }
                                if len(v2) > 0 {
                                    v2[0] = fi31
                                }
                                fmt.Println(fi34, v4, v3)
                                _ = v53
                            } else {
                                v4 -= v4
                            }
                            var v64 uint64 = v3
                            v3 = 376
                            _ = v64
                        } else {
                            _ = len(v2)
                            v5 = v3
                            if len(v2) > 0 {
                                v2[0] = 431
                            }
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        delete(v1, v4)
                        v4 += (v4 + 49)
                    }
                    _ = fi34
                }
                var v65 int8 = 0
                v19 += genFn0("foo bar", 218)
                _ = v65
            }
            v5 = v5
        }
        for fi66 := 0; fi66 < 2; fi66++ {
            v5 += (^(v3))
            switch -845 {
            case 32768:
                var v67 float64 = (1e10 + 1.5)
                _ = v67
            }
            _ = fi66
        }
        _ = fi31
    }
    v3 = 372
    for fi68 := 0; fi68 < 15; fi68++ {
        switch "a" {
        case "foo bar":
            delete(v1, v4)
        default:
            var v69 int64 = -1
            _ = v69
        }
        if true {
            if (0.1 <= 3.14159) {
                if (true && true) {
                    v70 := GenS7{F0: (fi68 | (fi68 << ((42) & 63)))}
                    for fi71 := 0; fi71 < 12; fi71++ {
                        v72 := GenS8{F0: 18446744073709551615, F1: 966}
                        _ = fi71
                        _ = v72
                    }
                    _ = v70
                } else {
                    for fi73 := 0; fi73 < 20; fi73++ {
                        v19 -= v19
                        _ = fi73
                    }
                    v19 = (v19 | v19)
                    v74 := GenS2{F0: -1, F1: ((true && false) && (true && (!true)))}
                    _ = v74
                }
                for fi75 := 0; fi75 < 12; fi75++ {
                    v76 := GenS9{F0: -2147483648, F1: 2, F2: 0}
                    delete(v1, (v4 + (v4 - v4)))
                    var v77 float32 = -2.25
                    _ = len(v1)
                    _ = fi75
                    _ = v76
                    _ = v77
                }
                var v78 int32 = 1
                _ = v78
            } else {
                for fi79 := 0; fi79 < 4; fi79++ {
                    if (!true) {
                        v80 := GenS10{F0: (v19 - (v19 | v19)), F1: ("" < "foo bar")}
                        v5 += 4294967295
                        var v81 uint8 = 128
                        v82 := GenS7{F0: fi79}
                        _ = v80
                        _ = v81
                        _ = v82
                    }
                    _ = fi79
                }
                _ = len(v1)
                var v83 uint = 4294967295
                _ = v83
            }
            if false {
                var v84 int64 = 9223372036854775807
                var v85 int64 = v84
                v86 := v1[v4]
                _ = v84
                _ = v85
                _ = v86
            }
        } else {
            var v87 uint64 = v3
            _ = v87
        }
        _ = fi68
    }
    v1[556] = 826
    if (-2.25 >= 1e10) {
        for fi88 := 0; fi88 < 17; fi88++ {
            v89 := GenS11{F0: 42}
            switch ("a" + "a") {
            case "\n\t":
                if (-9223372036854775808 == fi88) {
                    fmt.Println(v3)
                    v19 -= (v19 * v19)
                    var v90 uint16 = 42
                    _ = v90
                }
            case "":
                if len(v6) > 0 {
                    v6[0] = ("\n\t" + ("hello" + ""))
                }
            case "foo bar":
                v91 := GenS1{F0: 127, F1: (v4 / ((42) | 1))}
                _ = v91
            default:
                v92 := GenS12{F0: (v3 * (v5 << ((18446744073709551615) & 63)))}
                _ = v92
            }
            switch "\n\t" {
            case "\n\t":
                for fi93 := 0; fi93 < 9; fi93++ {
                    switch "\n\t" {
                    case "":
                        if (false && true) {
                            var v94 int = fi88
                            var v95 int = -2147483648
                            _ = v94
                            _ = v95
                        } else {
                            var v96 int = int(v4)
                            v97 := GenS13{F0: 0, F1: 102, F2: -1}
                            var v98 string = ("foo bar" + "hello")
                            v96 = fi93
                            v19 += (v19 + (v19 / ((v19) | 1)))
                            _ = v96
                            _ = v97
                            _ = v98
                        }
                    case "\n\t":
                        if ((^(fi88)) != -40) {
                            v3 = v5
                            var v99 float32 = float32((1.5 * 1e10))
                            _ = v99
                        }
                    case "a":
                        if false {
                            switch "a" {
                            case "":
                                if true {
                                    if (1 > v3) {
                                        var v100 float32 = 0.1
                                        v4 += v4
                                        v101 := v1[uint32(v3)]
                                        _ = v100
                                        _ = v101
                                    } else {
                                        var v102 uint8 = 2
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v102
                                    }
                                    if false {
                                        v5 -= 18446744073709551615
                                        v103 := GenS14{F0: 2, F1: 127, F2: v19}
                                        _ = v103
                                    }
                                    var v104 int32 = -28
                                    _ = v104
                                }
                            default:
                                v4 = (v4 ^ 2147483648)
                            }
                        } else {
                            fmt.Println(v3, v4, fi93)
                            if true {
                                v105 := GenS2{F0: 216, F1: true}
                                var v106 uint16 = 32767
                                var v107 float32 = (0.0 / 1e10)
                                _ = v105
                                _ = v106
                                _ = v107
                            } else {
                                if false {
                                    v108 := GenS2{F0: 9223372036854775807, F1: false}
                                    _ = v108
                                } else {
                                    var v109 int64 = -1
                                    if (0 > int8(v3)) {
                                        var v110 int8 = -79
                                        if (v19 >= (^(v19))) {
                                            v111 := GenS15{F0: (^(v110)), F1: 32767, F2: v110}
                                            v112 := GenS16{F0: (fi93 ^ (fi93 << ((829) & 63))), F1: v3, F2: 65535}
                                            v3 += v3
                                            v113 := GenS1{F0: (v110 + (v110 / ((-1) | 1))), F1: 2147483648}
                                            delete(v1, (v4 | 0))
                                            _ = v111
                                            _ = v112
                                            _ = v113
                                        }
                                        var v114 uint16 = 0
                                        _ = v110
                                        _ = v114
                                    } else {
                                        v115 := GenS6{F0: 1}
                                        _ = v115
                                    }
                                    if (!false) {
                                        if (1 < v109) {
                                            var v116 int = fi93
                                            var v117 uint16 = 42
                                            var v118 int8 = -128
                                            var v119 float32 = 1.5
                                            _ = v116
                                            _ = v117
                                            _ = v118
                                            _ = v119
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(v5)
                                            v3 += (v5 * v3)
                                        } else {
                                            fmt.Println(v5)
                                        }
                                        if false {
                                            fmt.Println(v5)
                                            v4 = v4
                                            if ((v3 * (v3 * (v5 & 0))) <= (v5 ^ (v5 & v3))) {
                                                var v120 uint = 464
                                                v121 := GenS8{F0: (v120 | (v120 & v120)), F1: (fi93 * (fi88 | fi88))}
                                                _ = v120
                                                _ = v121
                                            }
                                        } else {
                                            var v122 float32 = 0.0
                                            var v123 int32 = -1
                                            _ = v122
                                            _ = v123
                                        }
                                    }
                                    v124 := GenS13{F0: 1, F1: 127, F2: 6}
                                    _ = v109
                                    _ = v124
                                }
                                v4 = 718
                                fmt.Println(v19, v4, fi88)
                            }
                            if (!(2147483647 == 2147483647)) {
                                v125 := GenS6{F0: 42}
                                _ = v125
                            }
                        }
                    default:
                        v4 = v4
                        if true {
                            if (-1 == 42) {
                                var v126 int8 = 100
                                _ = v126
                            }
                        }
                        var v127 string = "a"
                        _ = v127
                    }
                    _ = fi93
                }
            case "":
                if ((0.1 * float64(0.1)) > ((float64(0.0) + ((3.14159 - (0.1 * 0.0)) + 1e10)) / (float64(3.14159) / 1.0))) {
                    delete(v1, 226)
                    v3 -= 2
                    delete(v1, v4)
                }
            case "hello":
                if len(v2) > 0 {
                    v2[0] = fi88
                }
            }
            _ = fi88
            _ = v89
        }
    }
    v128 := GenS9{F0: -1, F1: 161, F2: -714}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v128
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
