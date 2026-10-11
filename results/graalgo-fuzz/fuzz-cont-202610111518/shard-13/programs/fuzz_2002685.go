package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int64; F1 string; F2 uint }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint64; F1 int; F2 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 int; F1 uint32 }
type GenS13 struct { F0 int; F1 string }
type GenS14 struct { F0 uint32; F1 string }
type GenS15 struct { F0 int32; F1 string }
func genFn0() uint64 {
    var v6 int = -523
    v6 -= (v6 | -840)
    v6 = v6
    v6 = 873
    _ = v6
    return 4294967295
}
func genFn1(p0 uint8, p1 string) int16 {
    v7 := []int{1, 127, -646, 32767}
    p1 = "a"
    switch 32768 {
    case -1:
        v8 := GenS1{F0: 0}
        _ = v8
    case 109:
        var v9 uint64 = genFn0()
        _ = v9
    }
    var v10 bool = (1 >= 0)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return -1
}

func main() {
    var v1 uint8 = 100
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    var v4 int8 = -128
    var v5 uint64 = 871
    for fi11 := 0; fi11 < 1; fi11++ {
        var v12 float64 = -2.25
        _ = fi11
        _ = v12
    }
    switch ("hello" + "hello") {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        if ((v1 != (v1 << ((2) & 63))) && false) {
            v3 += (v3 % (((v3 | v3)) | 1))
        }
    case "":
        if (-1 == 42) {
            v13 := GenS2{F0: 231, F1: 2, F2: (true || true)}
            if true {
                var v14 uint64 = v5
                v15 := GenS2{F0: -2147483648, F1: 2, F2: (false || false)}
                fmt.Println(v5)
                _ = v14
                _ = v15
            }
            _ = v13
        }
    default:
        switch ("" + "foo bar") {
        case "a":
            v16 := v2[("hello" + ((("a" + "a") + "foo bar") + ("hello" + (("foo bar" + "\n\t") + ("" + ("\n\t" + "\n\t"))))))]
            _ = v16
        }
    }
    v17 := []int{-9223372036854775808, -9223372036854775808, 9223372036854775807}
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    if ((^(v1)) <= v1) {
        var v18 int32 = -1
        _ = v18
    }
    v19 := GenS1{F0: (v5 - (^(v5)))}
    var v20 int = 9223372036854775807
    var v21 float32 = (float32(3.14159) / ((0.0 + 1.5) + float32(0.1)))
    if true {
        var v22 int8 = v4
        for fi23 := 0; fi23 < 5; fi23++ {
            v24 := GenS2{F0: fi23, F1: 42, F2: (0.0 != (0.0 - 1e10))}
            _ = fi23
            _ = v24
        }
        _ = v22
    } else {
        switch 1 {
        case 127:
            if true {
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS2{F0: -494, F1: 291, F2: ((false || true) || ((v20 * 32768) > 32768))}
                    _ = fi25
                    _ = v26
                }
            }
        case -9223372036854775808:
            v27 := GenS2{F0: -302, F1: 392, F2: true}
            _ = v27
        default:
            for fi28 := 0; fi28 < 12; fi28++ {
                _ = len(v17)
                switch "foo bar" {
                case "\n\t":
                    var v29 float32 = v21
                    _ = v29
                }
                if len(v17) > 0 {
                    v17[0] = fi28
                }
                _ = fi28
            }
        }
        fmt.Println(v3, v20, v1)
    }
    switch (v20 | v20) {
    case 42:
        var v30 int16 = (v3 & v3)
        _ = v30
    case 181:
        switch "" {
        case "":
            switch (v20 - v20) {
            case 32767:
                fmt.Println(v4, v21, v5)
            case 2147483647:
                switch v20 {
                case 626:
                    v4 += -1
                case 32768:
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v20, v20)
                    default:
                        v31 := GenS2{F0: -2147483648, F1: 124, F2: false}
                        _ = v31
                    }
                case 9223372036854775807:
                    for fi32 := 0; fi32 < 5; fi32++ {
                        if true {
                            fmt.Println(v4, v20, v3)
                        } else {
                            v21 -= (v21 - (3.14159 - v21))
                            v20 -= (v20 << ((4294967295) & 63))
                        }
                        for fi33 := 0; fi33 < 9; fi33++ {
                            if ((false && (-9223372036854775808 < (fi32 << ((1) & 63)))) && true) {
                                if false {
                                    switch (fi32 | 32767) {
                                    case -2147483648:
                                        if false {
                                            var v34 int = (fi32 ^ 4294967295)
                                            fmt.Println(v1, v21)
                                            fmt.Println(v34)
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            delete(v2, "foo bar")
                                            switch fi32 {
                                            case 9223372036854775807:
                                                var v35 int = fi32
                                                _ = v35
                                            case -2147483648:
                                                v36 := GenS1{F0: v5}
                                                _ = v36
                                            case 32767:
                                                if false {
                                                    v37 := GenS1{F0: v5}
                                                    v4 = v4
                                                    _ = v37
                                                } else {
                                                    v38 := GenS3{F0: 1, F1: 42, F2: 1}
                                                    _ = v38
                                                }
                                            }
                                            _ = v34
                                        }
                                    }
                                    _ = len(v2)
                                    v1 = (v1 | v1)
                                } else {
                                    if true {
                                        v20 = 0
                                    } else {
                                        v39 := GenS2{F0: fi32, F1: uint16(v1), F2: (1 > -142)}
                                        if false {
                                            fmt.Println(v20)
                                            if true {
                                                var v40 int = (-(fi33))
                                                var v41 uint16 = 429
                                                _ = v40
                                                _ = v41
                                            }
                                            v21 += v21
                                        }
                                        _ = v39
                                    }
                                    v20 = fi33
                                    if false {
                                        v42 := GenS1{F0: v5}
                                        _ = v42
                                    } else {
                                        v1 = v1
                                        if (!((^(v1)) <= 1)) {
                                            var v43 int64 = 9223372036854775807
                                            v44 := v2[""]
                                            _ = len(v2)
                                            _ = v43
                                            _ = v44
                                        } else {
                                            if ((("hello" + ("" + "\n\t")) < "hello") && ((v4 >> ((18446744073709551615) & 63)) < v4)) {
                                                if false {
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v3)
                                                        case "hello":
                                                            if true {
                                                                var v45 uint16 = 255
                                                                v46 := GenS4{F0: 2}
                                                                if (0 == int16(v20)) {
                                                                    if true {
                                                                        var v47 int64 = 1
                                                                        var v48 int64 = v47
                                                                        _ = v47
                                                                        _ = v48
                                                                    }
                                                                    var v49 int16 = (v3 << ((uint(v5)) & 63))
                                                                    v49 = 681
                                                                    v4 = 122
                                                                    _ = v49
                                                                } else {
                                                                    var v50 int8 = (v4 - v4)
                                                                    _ = v50
                                                                }
                                                                _ = v45
                                                                _ = v46
                                                            }
                                                        default:
                                                            if ("hello" != ("hello" + "foo bar")) {
                                                                v51 := GenS2{F0: -9223372036854775808, F1: 577, F2: true}
                                                                _ = v51
                                                            } else {
                                                                if true {
                                                                    var v52 int8 = (v4 / (((v4 << ((18446744073709551615) & 63))) | 1))
                                                                    v53 := v2[(("a" + "\n\t") + "a")]
                                                                    _ = v52
                                                                    _ = v53
                                                                } else {
                                                                    if len(v17) >= 2 {
                                                                        v17 = v17[1:len(v17)]
                                                                    }
                                                                    switch ("a" + "\n\t") {
                                                                    case "":
                                                                        fmt.Println(v21, fi32, v1)
                                                                    case "\n\t":
                                                                        v4 = (v4 / (((v4 - (v4 / ((v4) | 1)))) | 1))
                                                                    case "a":
                                                                        var v54 uint8 = v1
                                                                        _ = v54
                                                                    default:
                                                                        if (((v21 - v21) / (v21 * 1e10)) != (v21 / v21)) {
                                                                            var v55 string = "a"
                                                                            v56 := v2[v55]
                                                                            v57 := GenS5{F0: v5, F1: (v3 >> ((1) & 63))}
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v57
                                                                        } else {
                                                                            var v58 uint32 = 65535
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                }
                                                                v3 = (v3 % ((634) | 1))
                                                            }
                                                        }
                                                    case "a":
                                                        var v59 float64 = float64(float32(((1e10 * (-2.25 - 1e10)) / 1e10)))
                                                        _ = v59
                                                    case "hello":
                                                        var v60 int8 = v4
                                                        _ = v60
                                                    default:
                                                        if (v1 >= (v1 * 1)) {
                                                            var v61 uint8 = 157
                                                            var v62 string = ""
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            var v63 int = -1
                                                            fmt.Println(fi32, v3)
                                                            v20 = (-(fi33))
                                                            _ = v63
                                                        }
                                                    }
                                                    v4 = v4
                                                    if (!((!(18446744073709551615 == 255)) || false)) {
                                                        if (("foo bar" + "\n\t") <= "foo bar") {
                                                            switch fi33 {
                                                            case 128:
                                                                v64 := GenS6{F0: -1, F1: "\n\t", F2: uint(v1)}
                                                                _ = v64
                                                            }
                                                            v4 = v4
                                                            v4 = -128
                                                        }
                                                    } else {
                                                        if true {
                                                            var v65 uint8 = v1
                                                            v5 = v5
                                                            _ = v65
                                                        }
                                                        v3 = (^(v3))
                                                        if len(v17) > 0 {
                                                            v17[0] = (-(fi32))
                                                        }
                                                        if false {
                                                            if (!false) {
                                                                if false {
                                                                    v17 = append(v17, -1)
                                                                    v2["foo bar"] = false
                                                                } else {
                                                                    v3 = 0
                                                                    switch (("\n\t" + "hello") + ("\n\t" + "foo bar")) {
                                                                    case "hello":
                                                                        v66 := GenS7{F0: fi32, F1: "a"}
                                                                        _ = v66
                                                                    case "":
                                                                        var v67 int16 = (v3 / (((v3 & genFn1(127, ("\n\t" + "")))) | 1))
                                                                        _ = v67
                                                                    case "a":
                                                                        v68 := GenS8{F0: 834, F1: 127}
                                                                        _ = v68
                                                                    default:
                                                                        switch v20 {
                                                                        case -9223372036854775808:
                                                                            var v69 int64 = 0
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                var v70 int = fi33
                                                                v71 := GenS2{F0: (^(fi33)), F1: 42, F2: (((false && true) && false) && false)}
                                                                _ = v70
                                                                _ = v71
                                                            }
                                                            var v72 int8 = (v4 % (((v4 ^ v4)) | 1))
                                                            _ = v72
                                                        }
                                                        fmt.Println(fi32, v5, fi33)
                                                        switch -1 {
                                                        case 32768:
                                                            var v73 int64 = 348
                                                            _ = v73
                                                        case -634:
                                                            v2[""] = (v3 != v3)
                                                        default:
                                                            var v74 int = v20
                                                            _ = v74
                                                        }
                                                        switch "\n\t" {
                                                        case "a":
                                                            var v75 int8 = -1
                                                            _ = v75
                                                        case "":
                                                            v76 := GenS9{F0: (v5 >> ((18446744073709551615) & 63)), F1: (fi32 % ((v20) | 1)), F2: 42}
                                                            _ = v76
                                                        case "\n\t":
                                                            var v77 float32 = v21
                                                            _ = v77
                                                        default:
                                                            if len(v17) > 0 {
                                                                v17[0] = v20
                                                            }
                                                        }
                                                        v78 := GenS10{F0: 488}
                                                        _ = v78
                                                    }
                                                } else {
                                                    fmt.Println(v20)
                                                    v79 := GenS11{F0: 0, F1: 165}
                                                    _ = v79
                                                }
                                                if (genFn1(147, "foo bar") > v3) {
                                                    if (!true) {
                                                        v80 := GenS12{F0: (^(fi32)), F1: 2147483648}
                                                        v81 := GenS11{F0: 2, F1: 255}
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    v82 := GenS5{F0: (v5 * 715), F1: -1}
                                                    v17 = append(v17, (fi32 - (fi32 | fi32)))
                                                    _ = v82
                                                } else {
                                                    _ = len(v17)
                                                }
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                            } else {
                                                if false {
                                                    var v83 uint8 = 42
                                                    _ = v83
                                                }
                                                var v84 int64 = 0
                                                fmt.Println(v84, v4)
                                                _ = v84
                                            }
                                            v20 = (^(fi32))
                                        }
                                        var v85 uint = 622
                                        _ = v85
                                    }
                                }
                                var v86 int8 = v4
                                _ = v86
                            } else {
                                if ((v5 % ((1) | 1)) < 18446744073709551615) {
                                    v4 += (v4 | -128)
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v4 -= (-(v4))
                                case "foo bar":
                                    v5 += 4294967295
                                default:
                                    fmt.Println(fi32)
                                }
                            }
                            switch (fi32 >> ((255) & 63)) {
                            case -442:
                                var v87 uint32 = 4294967295
                                _ = v87
                            case 4294967295:
                                switch (("\n\t" + "foo bar") + (((("hello" + "") + "") + "\n\t") + "a")) {
                                case "\n\t":
                                    var v88 int8 = 42
                                    _ = v88
                                case "":
                                    var v89 float32 = 0.1
                                    _ = v89
                                case "foo bar":
                                    var v90 int64 = 1
                                    _ = v90
                                }
                            case 128:
                                switch v20 {
                                case 32767:
                                    v4 = v4
                                case 20:
                                    fmt.Println(fi32, v4, fi33)
                                case -2147483648:
                                    var v91 uint64 = v5
                                    _ = v91
                                }
                            default:
                                var v92 int16 = 0
                                _ = v92
                            }
                            _ = fi33
                        }
                        if (65535 != 2) {
                            if (true && true) {
                                v3 = int16(v4)
                                v93 := GenS13{F0: fi32, F1: "foo bar"}
                                v21 -= (v21 - 1.5)
                                _ = v93
                            } else {
                                for fi94 := 0; fi94 < 12; fi94++ {
                                    var v95 int32 = -2147483648
                                    fmt.Println(v20, v4, fi32)
                                    _ = fi94
                                    _ = v95
                                }
                                for fi96 := 0; fi96 < 12; fi96++ {
                                    if (!(42 == 873)) {
                                        var v97 uint16 = 32767
                                        var v98 uint8 = (v1 << ((82) & 63))
                                        switch (v20 * fi32) {
                                        case 0:
                                            var v99 bool = true
                                            _ = v99
                                        case -841:
                                            v100 := GenS12{F0: v20, F1: 42}
                                            _ = v100
                                        case 128:
                                            switch (("a" + (("" + "\n\t") + "")) + "\n\t") {
                                            case "foo bar":
                                                v3 = 32767
                                                v97 -= v97
                                            case "":
                                                var v101 uint = 18446744073709551615
                                                _ = v101
                                            default:
                                                var v102 float64 = 3.14159
                                                _ = v102
                                            }
                                        default:
                                            v103 := GenS14{F0: 2, F1: ""}
                                            _ = v103
                                        }
                                        var v104 int64 = 0
                                        fmt.Println(v21, v97)
                                        _ = v104
                                        _ = v97
                                        _ = v98
                                    } else {
                                        fmt.Println(v20)
                                        switch (^(fi32)) {
                                        case 561:
                                            var v105 uint8 = 128
                                            if (!(!(true && (419 != 0)))) {
                                                fmt.Println(v105, v1)
                                                switch fi32 {
                                                case 0:
                                                    switch (fi32 * 32767) {
                                                    case 32768:
                                                        v21 = v21
                                                    case 9223372036854775807:
                                                        _ = len(v2)
                                                    }
                                                default:
                                                    v106 := GenS9{F0: v5, F1: fi96, F2: 827}
                                                    _ = v106
                                                }
                                                var v107 bool = true
                                                v2["a"] = (!v107)
                                                switch ("hello" + ((("\n\t" + "foo bar") + "hello") + "foo bar")) {
                                                case "hello":
                                                    var v108 int = (-(fi32))
                                                    _ = v108
                                                case "a":
                                                    var v109 uint16 = 405
                                                    _ = v109
                                                default:
                                                    v21 += float32((1e10 + (3.14159 / 3.14159)))
                                                }
                                                var v110 int32 = -1
                                                var v111 bool = ((fi32 % (((fi96 ^ -665)) | 1)) < fi32)
                                                v107 = (-1 == int64(v3))
                                                _ = v107
                                                _ = v110
                                                _ = v111
                                            } else {
                                                fmt.Println(v4, v20, v21)
                                                var v112 bool = false
                                                v113 := GenS15{F0: -2147483648, F1: (("\n\t" + ("" + ("" + ""))) + "foo bar")}
                                                _ = v112
                                                _ = v113
                                            }
                                            var v114 int16 = (^(v3))
                                            _ = v105
                                            _ = v114
                                        }
                                    }
                                    if (!true) {
                                        var v115 float32 = v21
                                        var v116 int = (^(v20))
                                        v17 = append(v17, (v20 * 1))
                                        _ = v115
                                        _ = v116
                                    } else {
                                        var v117 int64 = 91
                                        if ((!(0 < (-(v3)))) && false) {
                                            var v118 float64 = float64(0.1)
                                            v21 += ((v21 * v21) * v21)
                                            var v119 uint8 = v1
                                            _ = v118
                                            _ = v119
                                        } else {
                                            v20 = -2147483648
                                        }
                                        _ = v117
                                    }
                                    _ = fi96
                                }
                                v17 = append(v17, -838)
                            }
                            v3 = (v3 >> ((4294967295) & 63))
                        }
                        _ = fi32
                    }
                }
            default:
                v20 = v20
            }
        case "foo bar":
            var v120 float32 = v21
            _ = v120
        }
    case 32767:
        var v121 float32 = (v21 - (((3.14159 / v21) + float32(float64(float32(1.5)))) * v21))
        _ = v121
    }
    v20 = (v20 - v20)
    v5 -= 806
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
