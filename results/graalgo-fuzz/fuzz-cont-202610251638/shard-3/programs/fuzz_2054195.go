package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int8; F1 uint32; F2 int64 }
type GenS5 struct { F0 int8; F1 int64; F2 int8 }
type GenS6 struct { F0 string; F1 int16; F2 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 uint8; F2 int16 }
type GenS9 struct { F0 int8; F1 int64; F2 int32 }
type GenS10 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int64; F1 uint; F2 int16 }
type GenS13 struct { F0 uint16; F1 int8; F2 uint32 }
type GenS14 struct { F0 int16; F1 int8; F2 int32 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 int; F1 int8 }
type GenS17 struct { F0 int32; F1 int64 }
type GenS18 struct { F0 bool }
func genFn0(p0 uint32) bool {
    v7 := GenS1{F0: -2147483648, F1: -1}
    if false {
        var v8 int32 = 42
        var v9 uint16 = 2
        _ = v8
        _ = v9
    } else {
        if (42 >= 333) {
            var v10 int8 = 127
            fmt.Println(p0, v10)
            _ = v10
        }
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11, p0)
            for fi12 := 0; fi12 < 6; fi12++ {
                switch fi11 {
                case -573:
                    p0 += 4294967295
                case 9223372036854775807:
                    p0 -= p0
                case 755:
                    v13 := []int{(fi12 << ((347) & 63))}
                    var v14 uint8 = 136
                    switch "\n\t" {
                    case "":
                        var v15 int16 = 32767
                        _ = v15
                    default:
                        switch (fi11 / ((fi11) | 1)) {
                        case 32768:
                            v16 := GenS2{F0: -122}
                            _ = v16
                        default:
                            v17 := GenS3{F0: "hello", F1: (^(v14))}
                            _ = v17
                        }
                    }
                    p0 = p0
                    _ = v13
                    _ = v14
                }
                if false {
                    p0 -= 889
                } else {
                    switch 32767 {
                    case -558:
                        fmt.Println(fi11)
                    case 2147483647:
                        v18 := []int{(fi11 * fi11), 128, (fi12 ^ (fi12 ^ 32767))}
                        _ = v18
                    case 9223372036854775807:
                        switch "foo bar" {
                        case "hello":
                            v19 := GenS4{F0: 0, F1: 2147483647, F2: -1}
                            _ = v19
                        case "\n\t":
                            fmt.Println(p0, fi12, fi12)
                        }
                    default:
                        if true {
                            p0 += (p0 + (p0 % (((p0 + (p0 * p0))) | 1)))
                            var v20 float64 = 1.5
                            v21 := GenS1{F0: fi12, F1: -32768}
                            _ = v20
                            _ = v21
                        } else {
                            v22 := []int{}
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            _ = v22
                        }
                    }
                }
                _ = fi12
            }
            _ = fi11
        }
        switch "a" {
        case "hello":
            p0 -= (p0 | (p0 % (((p0 & p0)) | 1)))
        case "\n\t":
            v23 := GenS5{F0: 42, F1: 9223372036854775807, F2: 4}
            _ = v23
        }
    }
    var v24 uint = 435
    _ = p0
    _ = v24
    _ = v7
    return true
}
func genFn1() int8 {
    if true {
        switch ("\n\t" + "\n\t") {
        case "":
            v25 := GenS6{F0: ("hello" + "hello"), F1: -1, F2: ((32767 != 42) || (930 > 725))}
            _ = v25
        }
        if false {
            var v26 uint8 = 157
            v26 = v26
            v26 -= v26
            _ = v26
        }
        var v27 int8 = -128
        _ = v27
    }
    var v28 uint8 = 13
    v28 -= (v28 * v28)
    var v29 uint8 = v28
    for fi30 := 0; fi30 < 8; fi30++ {
        var v31 uint8 = v28
        _ = fi30
        _ = v31
    }
    _ = v28
    _ = v29
    return -128
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 255
    var v3 int16 = 1000
    var v4 bool = (378 >= 2)
    v5 := make(map[uint32]string)
    var v6 bool = v4
    if (65535 > 42) {
        fmt.Println(v4)
        var v32 int16 = v3
        _ = v32
    } else {
        var v33 int64 = -518
        var v34 uint8 = 229
        if true {
            var v35 int16 = (v3 ^ (v3 + v3))
            v4 = (331 >= 18446744073709551615)
            _ = v35
        }
        _ = v33
        _ = v34
    }
    fmt.Println(v3, v2, v4)
    fmt.Println(v4)
    if v6 {
        if v6 {
            v36 := GenS6{F0: "", F1: v3, F2: (!v6)}
            var v37 int8 = 100
            var v38 int = -9223372036854775808
            _ = v36
            _ = v37
            _ = v38
        } else {
            v39 := []int{-748, 127, 1}
            for fi40 := 0; fi40 < 4; fi40++ {
                for fi41 := 0; fi41 < 1; fi41++ {
                    v6 = true
                    _ = fi41
                }
                _ = fi40
            }
            if false {
                v42 := GenS6{F0: "\n\t", F1: (v3 * (v3 & 225)), F2: (!v6)}
                _ = v42
            }
            _ = v39
        }
        v3 = (v3 << ((18446744073709551615) & 63))
        var v43 bool = (v6 || (false && ((^(v3)) >= 1000)))
        _ = v43
    }
    var v44 float32 = 1.5
    v45 := []int{-288, -2147483648, 4294967295, 9223372036854775807}
    for fi46 := 0; fi46 < 16; fi46++ {
        var v47 bool = v6
        v48 := GenS7{F0: (v4 && (!v6))}
        v3 += v3
        _ = fi46
        _ = v47
        _ = v48
    }
    switch 4294967295 {
    case -923:
        delete(v5, 0)
    default:
        v4 = (893 != 2)
    }
    v44 = (v44 / v44)
    v49 := GenS7{F0: (897 == 4294967295)}
    for fi50 := 0; fi50 < 4; fi50++ {
        if v6 {
            v2 = (v2 / ((v2) | 1))
            fmt.Println(fi50)
        } else {
            fmt.Println(v3)
            if (!v6) {
                if true {
                    fmt.Println(fi50, v3)
                    var v51 uint = 65535
                    _ = v51
                } else {
                    var v52 uint32 = 0
                    _ = v52
                }
                _ = len(v1)
                var v53 uint32 = 2
                _ = v53
            } else {
                if v6 {
                    for fi54 := 0; fi54 < 20; fi54++ {
                        var v55 uint16 = 32767
                        v56 := GenS7{F0: (true && v4)}
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    v57 := GenS4{F0: 127, F1: 42, F2: -1}
                    _ = v57
                } else {
                    v44 += ((v44 / v44) - v44)
                    v2 = (v2 - 128)
                }
                switch 128 {
                case 128:
                    var v58 int = fi50
                    _ = v58
                case 127:
                    if (42 != 42) {
                        v59 := GenS5{F0: -128, F1: -1, F2: 1}
                        _ = v59
                    }
                case 0:
                    var v60 uint64 = 42
                    _ = v60
                default:
                    if v6 {
                        for fi61 := 0; fi61 < 14; fi61++ {
                            v2 -= (v2 >> ((103) & 63))
                            var v62 string = ""
                            switch (v62 + "\n\t") {
                            case "foo bar":
                                var v63 uint = 871
                                _ = v63
                            case "":
                                var v64 uint32 = uint32(v2)
                                v65 := GenS2{F0: fi61}
                                _ = v64
                                _ = v65
                            case "\n\t":
                                switch ((v62 + v62) + v62) {
                                case "":
                                    fmt.Println(v62, v6)
                                case "a":
                                    if false {
                                        fmt.Println(v44, v4, v62)
                                        if v4 {
                                            var v66 int16 = 32767
                                            fmt.Println(fi61, v4, fi61)
                                            fmt.Println(fi61)
                                            _ = v66
                                        }
                                        _ = len(v1)
                                    } else {
                                        v44 = float32(((0.1 + 1.5) * 3.14159))
                                        if v6 {
                                            switch fi61 {
                                            case 128:
                                                var v67 float32 = v44
                                                v68 := GenS8{F0: fi61, F1: (v2 / ((63) | 1)), F2: (-(v3))}
                                                if (!(v6 || v6)) {
                                                    v3 = -32768
                                                    if (v4 && (0 <= 2147483647)) {
                                                        var v69 int32 = 0
                                                        _ = v69
                                                    } else {
                                                        if (208 >= -1) {
                                                            v5[42] = ""
                                                            v2 += 127
                                                            fmt.Println(v6, v3)
                                                        } else {
                                                            switch v62 {
                                                            case "a":
                                                                fmt.Println(v62, v67, fi61)
                                                            case "hello":
                                                                fmt.Println(fi50, v4)
                                                            case "foo bar":
                                                                delete(v5, 2147483647)
                                                            default:
                                                                var v70 uint = 95
                                                                _ = v70
                                                            }
                                                            fmt.Println(fi50, v62, v6)
                                                        }
                                                        switch fi61 {
                                                        case -529:
                                                            if (!true) {
                                                                var v71 uint = 565
                                                                _ = v71
                                                            } else {
                                                                if v6 {
                                                                    delete(v5, 903)
                                                                    v72 := v1[fi50]
                                                                    var v73 int = 0
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                                v2 = v2
                                                                _ = len(v1)
                                                            }
                                                        case 642:
                                                            var v74 uint16 = 32767
                                                            _ = v74
                                                        }
                                                        v75 := GenS7{F0: false}
                                                        _ = v75
                                                    }
                                                }
                                                _ = v67
                                                _ = v68
                                            case 42:
                                                v76 := GenS6{F0: "hello", F1: (v3 + 42), F2: true}
                                                _ = v76
                                            }
                                            if len(v45) > 0 {
                                                v45[0] = (^(fi61))
                                            }
                                        } else {
                                            fmt.Println(v6, fi50)
                                            v3 += -1
                                            v77 := GenS4{F0: 100, F1: 2147483648, F2: 0}
                                            v62 = (v62 + ((("" + v62) + v62) + ((("a" + v62) + v62) + "foo bar")))
                                            _ = v77
                                        }
                                    }
                                case "\n\t":
                                    var v78 int16 = 32767
                                    _ = v78
                                }
                            }
                            v79 := GenS8{F0: 784, F1: v2, F2: (-(v3))}
                            if len(v45) > 0 {
                                v45[0] = fi50
                            }
                            v80 := GenS8{F0: -366, F1: v2, F2: 0}
                            _ = fi61
                            _ = v62
                            _ = v79
                            _ = v80
                        }
                    }
                }
            }
            if (42 >= 209) {
                var v81 uint64 = 18446744073709551615
                var v82 int = (-(fi50))
                _ = v81
                _ = v82
            }
        }
        _ = fi50
    }
    v2 = (v2 - 173)
    switch 2147483647 {
    case 4294967295:
        switch -798 {
        case -724:
            if true {
                if (42 < 384) {
                    switch "foo bar" {
                    case "hello":
                        v3 -= (v3 + v3)
                    case "\n\t":
                        if v4 {
                            _ = len(v5)
                            v83 := GenS9{F0: -128, F1: -9223372036854775808, F2: 42}
                            _ = v83
                        } else {
                            fmt.Println(v2, v2)
                        }
                    default:
                        v84 := GenS10{F0: (v2 << ((286) & 63)), F1: (v2 << ((691) & 63)), F2: 274}
                        _ = v84
                    }
                    v3 = v3
                    v85 := GenS11{F0: "hello"}
                    _ = v85
                }
            }
        case -9223372036854775808:
            for fi86 := 0; fi86 < 1; fi86++ {
                v87 := GenS12{F0: 1, F1: 65535, F2: (v3 ^ v3)}
                for fi88 := 0; fi88 < 11; fi88++ {
                    var v89 int = 32768
                    var v90 string = "hello"
                    _ = fi88
                    _ = v89
                    _ = v90
                }
                v91 := GenS12{F0: 974, F1: 255, F2: v3}
                _ = fi86
                _ = v87
                _ = v91
            }
        default:
            for fi92 := 0; fi92 < 14; fi92++ {
                v93 := GenS4{F0: 100, F1: 65535, F2: 203}
                _ = fi92
                _ = v93
            }
        }
    }
    delete(v5, 93)
    if (255 >= 875) {
        switch "a" {
        case "foo bar":
            for fi94 := 0; fi94 < 15; fi94++ {
                fmt.Println(v3)
                v6 = false
                switch (fi94 - (fi94 % (((fi94 | fi94)) | 1))) {
                case 9223372036854775807:
                    if v4 {
                        switch fi94 {
                        case 1:
                            v95 := GenS13{F0: 784, F1: 100, F2: 1}
                            _ = v95
                        case 127:
                            var v96 float64 = 3.14159
                            _ = v96
                        default:
                            for fi97 := 0; fi97 < 14; fi97++ {
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                if false {
                                    v98 := GenS14{F0: 1, F1: genFn1(), F2: -1}
                                    var v99 uint32 = 1
                                    _ = v98
                                    _ = v99
                                }
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                _ = fi97
                            }
                        }
                        var v100 string = ("a" + "\n\t")
                        if true {
                            v4 = v6
                        }
                        _ = v100
                    }
                case 42:
                    if v4 {
                        if v6 {
                            v2 += 128
                            v44 -= v44
                            _ = len(v45)
                        } else {
                            var v101 float64 = (1e10 - 1e10)
                            var v102 uint8 = (v2 & (v2 >> ((18446744073709551615) & 63)))
                            v103 := GenS11{F0: "foo bar"}
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                        if v4 {
                            switch (fi94 & fi94) {
                            case 862:
                                fmt.Println(v2, v44, fi94)
                            case 42:
                                v104 := GenS15{F0: "hello"}
                                _ = v104
                            default:
                                v3 = -32768
                            }
                        } else {
                            v105 := GenS16{F0: (fi94 << ((655) & 63)), F1: 42}
                            var v106 uint8 = v2
                            _ = v105
                            _ = v106
                        }
                    } else {
                        var v107 uint64 = 422
                        _ = v107
                    }
                case -1:
                    var v108 int = (fi94 << ((4294967295) & 63))
                    _ = v108
                default:
                    v109 := GenS17{F0: 2147483647, F1: -466}
                    _ = v109
                }
                _ = fi94
            }
        case "a":
            v6 = true
        default:
            v6 = (38 > v3)
        }
        switch 2147483647 {
        case -105:
            if (((255 >= 35) && false) || true) {
                if (!v4) {
                    v110 := GenS18{F0: true}
                    switch 4294967295 {
                    case 1:
                        v3 += -596
                    }
                    v5[4294967295] = ("hello" + "")
                    _ = v110
                }
                v111 := GenS5{F0: 1, F1: 9223372036854775807, F2: int8(v3)}
                _ = v111
            }
        default:
            if (v3 <= 1) {
                v4 = v6
                var v112 uint = 773
                _ = v112
            } else {
                if v6 {
                    if v4 {
                        var v113 uint64 = 823
                        var v114 int16 = (^(v3))
                        var v115 uint64 = (v113 * v113)
                        _ = v113
                        _ = v114
                        _ = v115
                    }
                    var v116 uint32 = 622
                    if (!false) {
                        if true {
                            for fi117 := 0; fi117 < 16; fi117++ {
                                if false {
                                    if (42 != 959) {
                                        for fi118 := 0; fi118 < 17; fi118++ {
                                            v2 -= v2
                                            v119 := GenS10{F0: (^(v2)), F1: 0, F2: v116}
                                            _ = fi118
                                            _ = v119
                                        }
                                        for fi120 := 0; fi120 < 20; fi120++ {
                                            v116 = v116
                                            _ = fi120
                                        }
                                    } else {
                                        if len(v45) > 0 {
                                            v45[0] = (fi117 / ((fi117) | 1))
                                        }
                                    }
                                } else {
                                    var v121 int32 = 2147483647
                                    _ = v121
                                }
                                v3 = (^(v3))
                                fmt.Println(v44, v2, v116)
                                _ = fi117
                            }
                            v4 = ((true && v4) && v4)
                        } else {
                            if (!(false && true)) {
                                v116 += (v116 << ((832) & 63))
                            }
                            v2 = (^(v2))
                            v4 = v6
                        }
                        var v122 uint16 = 255
                        fmt.Println(v2, v122)
                        _ = v122
                    } else {
                        var v123 uint8 = 42
                        for fi124 := 0; fi124 < 6; fi124++ {
                            var v125 int64 = 9223372036854775807
                            var v126 int16 = -1
                            if ("" <= "hello") {
                                fmt.Println(v4, v4)
                                for fi127 := 0; fi127 < 10; fi127++ {
                                    var v128 uint32 = (^(v116))
                                    var v129 string = "hello"
                                    _ = fi127
                                    _ = v128
                                    _ = v129
                                }
                            } else {
                                for fi130 := 0; fi130 < 12; fi130++ {
                                    var v131 int8 = -128
                                    v116 = (v116 & v116)
                                    _ = fi130
                                    _ = v131
                                }
                            }
                            _ = fi124
                            _ = v125
                            _ = v126
                        }
                        _ = v123
                    }
                    _ = v116
                } else {
                    if (("foo bar" + "a") >= ("foo bar" + ("hello" + "\n\t"))) {
                        if (2 >= (v2 % ((v2) | 1))) {
                            for fi132 := 0; fi132 < 7; fi132++ {
                                fmt.Println(v4, v3, v4)
                                _ = fi132
                            }
                        }
                    } else {
                        if true {
                            var v133 int = 2147483647
                            var v134 uint8 = v2
                            var v135 uint8 = (v2 / ((v2) | 1))
                            _ = v133
                            _ = v134
                            _ = v135
                        } else {
                            var v136 float64 = 0.0
                            if (599 != 4294967295) {
                                if true {
                                    if v6 {
                                        for fi137 := 0; fi137 < 7; fi137++ {
                                            for fi138 := 0; fi138 < 16; fi138++ {
                                                v139 := GenS18{F0: false}
                                                _ = fi138
                                                _ = v139
                                            }
                                            v6 = (v6 && (-128 == 1))
                                            _ = fi137
                                        }
                                    }
                                    var v140 bool = ((^(v2)) != (^(v2)))
                                    _ = v140
                                }
                                v6 = v6
                            } else {
                                v136 -= 3.14159
                                fmt.Println(v44, v44)
                                var v141 uint8 = v2
                                _ = v141
                            }
                            for fi142 := 0; fi142 < 5; fi142++ {
                                var v143 float32 = -2.25
                                _ = fi142
                                _ = v143
                            }
                            _ = v136
                        }
                    }
                }
            }
        }
    }
    var v144 int64 = 641
    v6 = (!v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v144:", v144)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v144
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v49
    _ = v5
    _ = v6
}
