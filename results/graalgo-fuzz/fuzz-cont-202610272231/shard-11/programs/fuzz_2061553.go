package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint; F1 bool; F2 int }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 string }
type GenS6 struct { F0 bool; F1 int; F2 uint }
type GenS7 struct { F0 int16; F1 string; F2 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 int16; F1 int8 }
type GenS13 struct { F0 uint; F1 uint64; F2 bool }
type GenS14 struct { F0 int32; F1 uint8 }
type GenS15 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS16 struct { F0 uint32; F1 int16; F2 int16 }
func genFn0(p0 int8) int {
    v4 := []int{32768, 479, 128}
    p0 -= (p0 * p0)
    p0 = 42
    _ = p0
    _ = v4
    return 128
}
func genFn1(p0 int) string {
    v5 := GenS1{F0: 1, F1: 0}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return ("a" + "hello")
}

func main() {
    var v1 uint32 = 0
    v2 := []int{305, 1, 4294967295, 577}
    v3 := make(map[string]bool)
    v6 := GenS2{F0: 118, F1: true, F2: 32768}
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int32 = 1
            switch -1 {
            case 50:
                delete(v3, "\n\t")
                var v10 string = "hello"
                _ = v10
            case 9223372036854775807:
                if true {
                    if true {
                        var v11 uint = 42
                        if (v1 > v1) {
                            v9 = 1
                            v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        if true {
                            if ((3.14159 - (1.5 / -2.25)) <= 0.0) {
                                if true {
                                    v1 += v1
                                    fmt.Println(v9, fi7)
                                    v1 = v1
                                    if true {
                                        var v12 float64 = float64((0.1 + 1.5))
                                        if true {
                                            if false {
                                                switch (fi8 / ((fi7) | 1)) {
                                                case 32767:
                                                    fmt.Println(v9)
                                                case 0:
                                                    if len(v2) > 0 {
                                                        v2[0] = -9223372036854775808
                                                    }
                                                case 9223372036854775807:
                                                    v12 = (v12 / v12)
                                                }
                                                switch ("a" + "a") {
                                                case "hello":
                                                    var v13 uint = v11
                                                    _ = v13
                                                }
                                                var v14 uint32 = 24
                                                _ = v14
                                            } else {
                                                switch "" {
                                                case "hello":
                                                    if (65535 == v1) {
                                                        fmt.Println(v9)
                                                        var v15 int16 = -32768
                                                        var v16 uint8 = 165
                                                        _ = v15
                                                        _ = v16
                                                    }
                                                case "\n\t":
                                                    switch "foo bar" {
                                                    case "":
                                                        if false {
                                                            v17 := GenS3{F0: 507, F1: v11}
                                                            v11 -= v11
                                                            switch fi8 {
                                                            case 0:
                                                                v11 -= 255
                                                            }
                                                            _ = v17
                                                        } else {
                                                            v18 := GenS1{F0: -9223372036854775808, F1: 1}
                                                            var v19 int8 = 1
                                                            _ = v18
                                                            _ = v19
                                                        }
                                                    case "foo bar":
                                                        var v20 int64 = -9223372036854775808
                                                        _ = v20
                                                    }
                                                case "foo bar":
                                                    var v21 bool = ((("a" + "foo bar") <= "\n\t") && true)
                                                    _ = v21
                                                default:
                                                    var v22 uint64 = uint64(v1)
                                                    v11 += (^(v11))
                                                    _ = v22
                                                }
                                                if true {
                                                    if ((^(v11)) <= (v11 & 1)) {
                                                        fmt.Println(v1)
                                                        var v23 int16 = 42
                                                        v9 = (v9 ^ (v9 ^ (-(v9))))
                                                        _ = v23
                                                    } else {
                                                        if ((1 < 42) && (true && true)) {
                                                            var v24 float64 = v12
                                                            if true {
                                                                v3["a"] = (-1 >= -73)
                                                                v12 += (float64(float32((1.5 * v12))) - ((v24 - (0.1 * 1.5)) / 0.1))
                                                                var v25 uint = (v11 << ((0) & 63))
                                                                _ = v25
                                                            } else {
                                                                var v26 uint32 = 2
                                                                fmt.Println(fi8)
                                                                _ = v26
                                                            }
                                                            v27 := GenS4{F0: -80}
                                                            _ = v24
                                                            _ = v27
                                                        }
                                                        v3["foo bar"] = ((v1 + 314) <= 65)
                                                        var v28 bool = (genFn0(1) <= -576)
                                                        _ = v28
                                                    }
                                                    if (1.5 > 0.1) {
                                                        var v29 uint32 = (v1 - (^(v1)))
                                                        fmt.Println(v12, v29)
                                                        var v30 int8 = int8(v9)
                                                        v3[""] = false
                                                        v11 += v11
                                                        _ = v29
                                                        _ = v30
                                                    } else {
                                                        var v31 uint64 = 4294967295
                                                        _ = v31
                                                    }
                                                    var v32 uint = (v11 % ((v11) | 1))
                                                    _ = v32
                                                } else {
                                                    fmt.Println(v11, v11, v1)
                                                    var v33 int8 = -13
                                                    _ = v33
                                                }
                                                var v34 uint64 = 2
                                                _ = v34
                                            }
                                            if false {
                                                var v35 string = ""
                                                if ((!true) || (0 < 0)) {
                                                    var v36 uint16 = 255
                                                    _ = v36
                                                }
                                                _ = v35
                                            }
                                            fmt.Println(v12, v11)
                                        } else {
                                            var v37 float32 = 1.5
                                            switch int(v11) {
                                            case -1:
                                                v37 -= v37
                                            case 4294967295:
                                                var v38 uint64 = 2
                                                _ = v38
                                            case 9223372036854775807:
                                                switch "a" {
                                                case "hello":
                                                    _ = len(v2)
                                                case "\n\t":
                                                    if true {
                                                        v39 := GenS5{F0: 218, F1: 32767, F2: "\n\t"}
                                                        _ = v39
                                                    } else {
                                                        if true {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                }
                                            }
                                            if true {
                                                var v40 uint8 = 0
                                                var v41 int64 = 9223372036854775807
                                                _ = v40
                                                _ = v41
                                            }
                                            _ = v37
                                        }
                                        _ = v12
                                    } else {
                                        v42 := v3[(("hello" + "a") + ("" + "a"))]
                                        switch ("\n\t" + "") {
                                        case "hello":
                                            fmt.Println(v42)
                                        case "\n\t":
                                            v9 += (v9 ^ (^(v9)))
                                        case "foo bar":
                                            v42 = v42
                                        default:
                                            fmt.Println(fi8, v11)
                                            v43 := GenS6{F0: (v9 < v9), F1: genFn0(-1), F2: (v11 + (v11 / ((v11) | 1)))}
                                            _ = v43
                                        }
                                        _ = v42
                                    }
                                } else {
                                    v44 := GenS2{F0: (^(v11)), F1: (fi7 < int(v1)), F2: 9223372036854775807}
                                    if ((v9 << ((v11) & 63)) > (v9 & v9)) {
                                        var v45 float32 = 1e10
                                        fmt.Println(fi8)
                                        v46 := GenS7{F0: 1, F1: "foo bar", F2: (v11 >> ((v11) & 63))}
                                        v1 -= v1
                                        var v47 uint8 = 1
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    v1 = uint32(v11)
                                    _ = v44
                                }
                            } else {
                                v48 := GenS7{F0: 378, F1: "foo bar", F2: (v11 * v11)}
                                _ = v48
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v49 := GenS6{F0: false, F1: 12, F2: v11}
                            _ = v49
                        }
                        _ = v11
                    } else {
                        if ((1.5 / 0.1) <= (1e10 + 3.14159)) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if true {
                                if (true && (-1 == 0)) {
                                    v1 += v1
                                } else {
                                    var v50 uint32 = v1
                                    var v51 uint = 65535
                                    _ = v50
                                    _ = v51
                                }
                            }
                        }
                        if false {
                            v52 := GenS5{F0: 42, F1: 2, F2: ""}
                            v53 := GenS8{F0: v1}
                            v54 := GenS9{F0: 42}
                            _ = v52
                            _ = v53
                            _ = v54
                        } else {
                            switch "hello" {
                            case "hello":
                                if (-44 > -1) {
                                    var v55 int32 = 42
                                    var v56 uint = 2
                                    v9 -= (v55 % (((v55 * v55)) | 1))
                                    _ = v55
                                    _ = v56
                                } else {
                                    if false {
                                        if ((^(fi7)) > (fi8 + (fi7 << ((1) & 63)))) {
                                            v57 := v3[""]
                                            if (true || true) {
                                                fmt.Println(v1)
                                                var v58 uint = 1
                                                _ = v58
                                            } else {
                                                switch ("foo bar" + "a") {
                                                case "hello":
                                                    if v57 {
                                                        var v59 float32 = float32(-2.25)
                                                        _ = v59
                                                    } else {
                                                        var v60 int64 = 0
                                                        _ = v60
                                                    }
                                                case "\n\t":
                                                    if (v9 <= (v9 % ((0) | 1))) {
                                                        if (1e10 == 3.14159) {
                                                            v61 := GenS10{F0: false, F1: "\n\t"}
                                                            _ = v61
                                                        }
                                                        var v62 int16 = -420
                                                        var v63 int16 = 259
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                }
                                                var v64 int32 = v9
                                                _ = v64
                                            }
                                            _ = v57
                                        } else {
                                            v65 := v3[("hello" + "a")]
                                            var v66 int8 = 1
                                            _ = v65
                                            _ = v66
                                        }
                                        v67 := GenS11{F0: (-(fi8)), F1: v1}
                                        v68 := GenS4{F0: 116}
                                        _ = v67
                                        _ = v68
                                    }
                                    v69 := GenS12{F0: 32767, F1: 0}
                                    v9 = (v9 ^ 0)
                                    _ = v69
                                }
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v9 -= v9
                        }
                    }
                }
            case -9223372036854775808:
                if (uint16(v1) > 2) {
                    var v70 uint8 = 128
                    var v71 float32 = 3.14159
                    _ = v70
                    _ = v71
                } else {
                    var v72 int64 = 549
                    var v73 int8 = int8(v1)
                    _ = v72
                    _ = v73
                }
                switch ("a" + "") {
                case "a":
                    v74 := GenS4{F0: -76}
                    _ = v74
                }
            }
            v3["\n\t"] = false
            _ = len(v2)
            if (false && false) {
                fmt.Println(fi8, v9, fi7)
                _ = len(v3)
                switch fi8 {
                case 42:
                    fmt.Println(v1, fi7, fi7)
                case 0:
                    if true {
                        var v75 int16 = -32768
                        fmt.Println(v75, fi7, fi8)
                        _ = v75
                    } else {
                        v76 := GenS11{F0: (fi7 - 32768), F1: (v1 | 42)}
                        var v77 uint8 = 255
                        _ = v76
                        _ = v77
                    }
                case -1:
                    if len(v2) > 0 {
                        v2[0] = genFn0(42)
                    }
                }
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if (genFn0(-1) == 2147483647) {
        if (true || (2 != 18446744073709551615)) {
            _ = len(v3)
        }
        var v78 uint8 = 0
        _ = v78
    } else {
        v79 := GenS13{F0: 42, F1: 169, F2: (false || true)}
        _ = v79
    }
    v1 = (v1 | (v1 % ((0) | 1)))
    _ = len(v3)
    var v80 uint16 = 871
    if ((((1.5 / (-2.25 + 0.0)) / (((float64(-2.25) - (((float64(((1e10 / -2.25) * 1.5)) * 1.5) - 0.0) - 0.0)) / 1e10) + (3.14159 + 0.0))) != 0.1) || (true || true)) {
        var v81 int8 = -4
        switch "a" {
        case "":
            var v82 uint32 = 4294967295
            _ = v82
        case "a":
            var v83 uint16 = 385
            _ = v83
        default:
            for fi84 := 0; fi84 < 10; fi84++ {
                if true {
                    v85 := GenS14{F0: 719, F1: 42}
                    var v86 uint64 = 18446744073709551615
                    _ = v85
                    _ = v86
                } else {
                    for fi87 := 0; fi87 < 5; fi87++ {
                        switch "foo bar" {
                        case "":
                            if (!true) {
                                v80 -= (^(v80))
                                v81 -= (v81 / (((v81 / (((v81 | (v81 >> ((65535) & 63)))) | 1))) | 1))
                                fmt.Println(v1, v81, fi84)
                            } else {
                                fmt.Println(v1, v1)
                            }
                        default:
                            v88 := GenS15{F0: 157, F1: 1, F2: 1}
                            _ = v88
                        }
                        fmt.Println(fi87, fi87, v1)
                        _ = fi87
                    }
                    v3["hello"] = ((true && (false && false)) && true)
                }
                _ = fi84
            }
        }
        fmt.Println(v1, v1, v1)
        _ = v81
    }
    v89 := v3["foo bar"]
    switch 32767 {
    case -284:
        v80 = v80
    case 9223372036854775807:
        fmt.Println(v89, v1)
    }
    for fi90 := 0; fi90 < 15; fi90++ {
        v91 := GenS6{F0: (100 >= 127), F1: (fi90 * 32767), F2: 65535}
        switch "" {
        case "a":
            if (1 == int(v80)) {
                var v92 int64 = 652
                switch "\n\t" {
                case "a":
                    delete(v3, "hello")
                case "":
                    v89 = v89
                default:
                    for fi93 := 0; fi93 < 13; fi93++ {
                        var v94 int64 = 9223372036854775807
                        var v95 uint32 = v1
                        var v96 bool = v89
                        var v97 uint = uint(v95)
                        if true {
                            if (1 > (v92 & (v92 + (v92 & v94)))) {
                                switch 128 {
                                case -2147483648:
                                    fmt.Println(v80, v80)
                                    if true {
                                        v1 += v95
                                        v80 = (v80 + (v80 % ((v80) | 1)))
                                    }
                                    if false {
                                        switch (fi93 + (fi93 * fi93)) {
                                        case -449:
                                            var v98 bool = true
                                            _ = v98
                                        case 0:
                                            var v99 uint32 = (v95 / ((uint32(v97)) | 1))
                                            _ = v99
                                        case 127:
                                            var v100 int8 = 1
                                            _ = v100
                                        }
                                        v89 = ((((v97 << (((v97 % ((2) | 1))) & 63)) <= v97) || (2 > 2)) && v89)
                                        v101 := GenS16{F0: (v1 + v95), F1: 1, F2: 1}
                                        v92 += v94
                                        fmt.Println(v1)
                                        _ = v101
                                    } else {
                                        switch 0 {
                                        case 2147483647:
                                            v80 = 2
                                        case 32767:
                                            _ = len(v2)
                                        }
                                        if (false && v96) {
                                            var v102 uint32 = v1
                                            _ = v102
                                        }
                                        _ = len(v2)
                                    }
                                case 882:
                                    if (!v96) {
                                        v103 := GenS5{F0: 234, F1: (v80 & (v80 / ((v80) | 1))), F2: ""}
                                        var v104 string = "a"
                                        _ = v103
                                        _ = v104
                                    } else {
                                        v105 := v3["foo bar"]
                                        _ = v105
                                    }
                                case 0:
                                    _ = len(v2)
                                default:
                                    var v106 int16 = 32767
                                    v97 = 225
                                    v107 := GenS8{F0: 4294967295}
                                    _ = v106
                                    _ = v107
                                }
                                if (fi90 <= genFn0(-74)) {
                                    v108 := v3[("foo bar" + "a")]
                                    fmt.Println(v80)
                                    _ = v108
                                } else {
                                    fmt.Println(v96, v1, v96)
                                    v92 -= (v92 << (((v97 >> ((18446744073709551615) & 63))) & 63))
                                }
                                v80 = uint16(v1)
                            }
                            v109 := v3["foo bar"]
                            _ = v109
                        } else {
                            v97 -= 245
                            v80 -= (v80 * (v80 % (((v80 * (^(v80)))) | 1)))
                            fmt.Println(v97, v96, v94)
                        }
                        var v110 int8 = 0
                        _ = fi93
                        _ = v110
                        _ = v94
                        _ = v95
                        _ = v96
                        _ = v97
                    }
                }
                var v111 uint = 65535
                _ = v111
                _ = v92
            }
        case "foo bar":
            v1 -= 2147483648
        default:
            v80 = v80
        }
        _ = fi90
        _ = v91
    }
    var v112 int8 = -33
    var v113 bool = (("a" + "foo bar") >= "a")
    for fi114 := 0; fi114 < 17; fi114++ {
        if false {
            var v115 int = (fi114 ^ fi114)
            _ = v115
        } else {
            var v116 uint64 = 390
            v2 = append(v2, fi114)
            v117 := GenS14{F0: 479, F1: 69}
            _ = v116
            _ = v117
        }
        _ = fi114
    }
    v113 = v113
    switch 0 {
    case 42:
        v113 = (!(1e10 > 3.14159))
    }
    v118 := v3[((((("foo bar" + ("" + "foo bar")) + "foo bar") + ("foo bar" + "a")) + "a") + "foo bar")]
    v89 = false
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v118:", v118)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v80:", v80)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v112
    _ = v113
    _ = v118
    _ = v2
    _ = v3
    _ = v6
    _ = v80
    _ = v89
}
