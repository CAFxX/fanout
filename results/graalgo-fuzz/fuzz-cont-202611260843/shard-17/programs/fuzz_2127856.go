package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 int; F1 uint8; F2 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 bool; F1 int16 }
type GenS11 struct { F0 bool; F1 uint }
type GenS12 struct { F0 uint; F1 bool; F2 uint8 }
type GenS13 struct { F0 string; F1 int16; F2 int }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 string; F1 int8 }
func genFn0(p0 int32, p1 string, p2 string) uint16 {
    v4 := GenS1{F0: 0, F1: 400, F2: (p0 % ((-1) | 1))}
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int16 = 1000
        v7 := []int{(fi5 / ((1) | 1)), fi5}
        fmt.Println(fi5, p2)
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1() string {
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(fi8)
        _ = fi8
    }
    v9 := GenS1{F0: 446, F1: 427, F2: 42}
    var v10 uint32 = 4294967295
    _ = v10
    _ = v9
    return "hello"
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[int]uint32)
    var v3 float64 = 1e10
    if (-2.25 <= (-2.25 / (1e10 - 3.14159))) {
        switch ("foo bar" + "a") {
        case "\n\t":
            var v11 int = 32767
            _ = v11
        default:
            var v12 int16 = 1000
            _ = v12
        }
        if true {
            fmt.Println(v3, v1)
            for fi13 := 0; fi13 < 4; fi13++ {
                var v14 float32 = (1e10 / 3.14159)
                _ = fi13
                _ = v14
            }
            var v15 int64 = 160
            _ = v15
        } else {
            fmt.Println(v1)
            var v16 bool = ((v3 + v3) < 1.5)
            _ = v16
        }
        v2[-2147483648] = 2
    }
    v17 := GenS2{F0: 1, F1: 0}
    v18 := []int{1, -842, 56}
    var v19 int16 = -433
    v20 := GenS2{F0: 655, F1: 285}
    var v21 float64 = (v3 * (0.1 - -2.25))
    for fi22 := 0; fi22 < 1; fi22++ {
        _ = len(v2)
        for fi23 := 0; fi23 < 10; fi23++ {
            v1 = (v1 | v1)
            v19 -= (v19 << ((465) & 63))
            v24 := v2[32768]
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    if true {
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(v1)
            _ = fi25
        }
        for fi26 := 0; fi26 < 19; fi26++ {
            if false {
                if false {
                    if true {
                        var v27 uint64 = 42
                        var v28 float32 = 0.1
                        _ = v27
                        _ = v28
                    } else {
                        delete(v2, (fi26 % ((fi26) | 1)))
                    }
                }
                v29 := GenS3{F0: -9223372036854775808}
                _ = v29
            } else {
                for fi30 := 0; fi30 < 4; fi30++ {
                    if (9223372036854775807 >= -667) {
                        v31 := GenS4{F0: fi26}
                        _ = v31
                    } else {
                        delete(v2, (-(fi30)))
                    }
                    if (0 != -128) {
                        if true {
                            switch "" {
                            case "foo bar":
                                if (1 == (fi30 - fi26)) {
                                    fmt.Println(fi26)
                                    var v32 uint8 = 0
                                    _ = v32
                                } else {
                                    v33 := GenS3{F0: 0}
                                    _ = v33
                                }
                            case "a":
                                v19 += 210
                            case "\n\t":
                                v21 -= ((0.1 / float64(0.1)) / float64(1e10))
                            default:
                                fmt.Println(v3, fi26, v19)
                            }
                            var v34 uint64 = v1
                            switch ("hello" + "a") {
                            case "foo bar":
                                switch (("\n\t" + "a") + "") {
                                case "hello":
                                    v19 = (v19 - v19)
                                case "a":
                                    v21 += float64(3.14159)
                                case "foo bar":
                                    var v35 float32 = float32(0.1)
                                    _ = v35
                                default:
                                    fmt.Println(fi30)
                                    var v36 int = (fi30 + (-(fi26)))
                                    v37 := GenS5{F0: 444, F1: (fi26 & (^(fi30))), F2: 1}
                                    v21 += v21
                                    _ = v36
                                    _ = v37
                                }
                            case "":
                                fmt.Println(v34)
                            }
                            if true {
                                fmt.Println(fi30, fi30, fi26)
                                _ = len(v2)
                                var v38 uint8 = 127
                                _ = v38
                            }
                            _ = v34
                        }
                    } else {
                        if true {
                            var v39 bool = (((!true) || false) || (!(155 > 4294967295)))
                            var v40 float64 = v21
                            _ = v39
                            _ = v40
                        } else {
                            fmt.Println(fi26)
                            v41 := GenS6{F0: 1, F1: v19}
                            if true {
                                if false {
                                    v21 -= (-2.25 / ((v21 * (((((((float64(1.5) * float64(0.1)) / (float64(3.14159) + float64(-2.25))) - v3) + ((1.5 - v3) - v3)) * float64(1e10)) / v21) - v3)) + float64(1e10)))
                                    v42 := GenS2{F0: 0, F1: 2}
                                    v2[(fi26 | fi26)] = 2
                                    _ = v42
                                } else {
                                    var v43 int8 = 42
                                    _ = v43
                                }
                            }
                            if ((^(v1)) <= 42) {
                                var v44 int32 = -2147483648
                                var v45 int32 = 2147483647
                                _ = v44
                                _ = v45
                            }
                            _ = v41
                        }
                        _ = len(v2)
                        if true {
                            if (2 > 0) {
                                if true {
                                    if (!(((((^(v19)) == (v19 - (v19 >> ((180) & 63)))) || ((v19 << ((uint(v1)) & 63)) <= 588)) && (!(-128 == 42))) || (972 < v19))) {
                                        v46 := GenS2{F0: -1, F1: 32767}
                                        switch "\n\t" {
                                        case "\n\t":
                                            v3 -= (float64(1.5) * (3.14159 - v3))
                                        case "foo bar":
                                            v1 = 1
                                        default:
                                            fmt.Println(v21, v19, v1)
                                        }
                                        _ = v46
                                    } else {
                                        if true {
                                            var v47 uint8 = 42
                                            _ = v47
                                        }
                                    }
                                } else {
                                    v48 := GenS7{F0: 2147483647, F1: "hello"}
                                    v21 = (v3 - v21)
                                    v49 := GenS5{F0: -862, F1: -996, F2: 1}
                                    v1 = 2
                                    _ = v48
                                    _ = v49
                                }
                                switch fi30 {
                                case -2147483648:
                                    var v50 uint8 = 8
                                    _ = v50
                                case -1:
                                    var v51 string = "a"
                                    _ = v51
                                }
                                v2[-9223372036854775808] = 506
                                v21 = ((0.1 * v3) - float64((0.1 * 1.5)))
                            } else {
                                var v52 int32 = 2147483647
                                switch "foo bar" {
                                case "":
                                    if true {
                                        _ = len(v2)
                                        if (-1 < (v19 | (v19 / ((1) | 1)))) {
                                            fmt.Println(v21)
                                            v53 := GenS6{F0: 0, F1: v19}
                                            _ = v53
                                        }
                                        var v54 int8 = 1
                                        switch (("a" + (("foo bar" + ("\n\t" + ("foo bar" + ((("foo bar" + "hello") + "foo bar") + "hello")))) + (("a" + ("" + ("a" + "a"))) + ""))) + ("a" + ("" + ("foo bar" + ("a" + "\n\t"))))) {
                                        case "\n\t":
                                            v55 := GenS8{F0: (fi26 >> ((42) & 63)), F1: 255, F2: -305}
                                            v56 := v2[(fi26 << ((839) & 63))]
                                            _ = v55
                                            _ = v56
                                        case "foo bar":
                                            v52 -= (v52 << ((255) & 63))
                                        case "":
                                            if false {
                                                v57 := GenS7{F0: v52, F1: (("foo bar" + "foo bar") + ((("hello" + "") + ("" + "")) + "hello"))}
                                                if ((v54 ^ v54) == (v54 * (v54 << ((1) & 63)))) {
                                                    if (!true) {
                                                        if false {
                                                            if false {
                                                                var v58 string = "foo bar"
                                                                _ = v58
                                                            } else {
                                                                var v59 int = -539
                                                                switch fi26 {
                                                                case 32767:
                                                                    if true {
                                                                        v52 -= -1
                                                                    } else {
                                                                        var v60 string = ""
                                                                        fmt.Println(v60, fi26, v52)
                                                                        var v61 uint32 = 1
                                                                        _ = v60
                                                                        _ = v61
                                                                    }
                                                                }
                                                                v62 := GenS8{F0: (fi26 & fi30), F1: 201, F2: 9223372036854775807}
                                                                fmt.Println(fi30, v59)
                                                                _ = v59
                                                                _ = v62
                                                            }
                                                            var v63 int = 977
                                                            if true {
                                                                switch (v63 & fi30) {
                                                                case 9223372036854775807:
                                                                    v2[-1] = 319
                                                                }
                                                                v21 = 1.5
                                                            } else {
                                                                v21 = (v21 * v3)
                                                                var v64 float32 = -2.25
                                                                v65 := GenS2{F0: v52, F1: 32767}
                                                                var v66 int16 = (v19 % ((658) | 1))
                                                                var v67 int32 = v52
                                                                v64 -= float32(0.1)
                                                                _ = v64
                                                                _ = v65
                                                                _ = v66
                                                                _ = v67
                                                            }
                                                            _ = v63
                                                        } else {
                                                            if true {
                                                                fmt.Println(v3, v52, v1)
                                                                var v68 float32 = ((1e10 + 0.1) + -2.25)
                                                                fmt.Println(v21, v3)
                                                                _ = v68
                                                            }
                                                            switch "\n\t" {
                                                            case "":
                                                                switch fi30 {
                                                                case 9223372036854775807:
                                                                    var v69 uint32 = 42
                                                                    _ = v69
                                                                case -9223372036854775808:
                                                                    v1 -= v1
                                                                case 128:
                                                                    if false {
                                                                        v54 += (v54 & v54)
                                                                        switch 128 {
                                                                        case -13:
                                                                            var v70 uint16 = 0
                                                                            _ = v70
                                                                        case -10:
                                                                            fmt.Println(fi30, v19, fi26)
                                                                        default:
                                                                            v71 := GenS5{F0: 9223372036854775807, F1: fi26, F2: v52}
                                                                            _ = v71
                                                                        }
                                                                        var v72 int = fi26
                                                                        var v73 float32 = (((3.14159 - 0.1) / 1.0) + 3.14159)
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                default:
                                                                    v3 = 1.5
                                                                }
                                                            default:
                                                                if (true && true) {
                                                                    if true {
                                                                        v52 = -918
                                                                        switch (("foo bar" + "hello") + ("" + ("foo bar" + "a"))) {
                                                                        case "\n\t":
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v74 float32 = 0.1
                                                                                _ = v74
                                                                            case "\n\t":
                                                                                if false {
                                                                                    v3 -= (v21 / ((v21 * v21) * v21))
                                                                                    var v75 uint32 = 15
                                                                                    _ = v75
                                                                                } else {
                                                                                    if false {
                                                                                        var v76 uint16 = 1
                                                                                        var v77 uint32 = 4294967295
                                                                                        _ = v76
                                                                                        _ = v77
                                                                                    } else {
                                                                                        switch (^(fi30)) {
                                                                                        case -1:
                                                                                            delete(v2, -87)
                                                                                        default:
                                                                                            if (65535 <= 1) {
                                                                                                switch "a" {
                                                                                                case "":
                                                                                                    switch (-(fi30)) {
                                                                                                    case 432:
                                                                                                        var v78 uint32 = 0
                                                                                                        _ = v78
                                                                                                    case 128:
                                                                                                        switch "hello" {
                                                                                                        case "foo bar":
                                                                                                            switch -2147483648 {
                                                                                                            case 2147483647:
                                                                                                                v2[fi26] = 42
                                                                                                            case -9223372036854775808:
                                                                                                                v21 += float64((0.0 * 3.14159))
                                                                                                            case 1:
                                                                                                                _ = len(v2)
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v79 := GenS9{F0: 42}
                                                                                                            _ = v79
                                                                                                        case "":
                                                                                                            fmt.Println(v52, fi30)
                                                                                                        default:
                                                                                                            v21 += 3.14159
                                                                                                        }
                                                                                                    case 32767:
                                                                                                        switch "a" {
                                                                                                        case "foo bar":
                                                                                                            _ = len(v18)
                                                                                                        default:
                                                                                                            var v80 int32 = 637
                                                                                                            _ = v80
                                                                                                        }
                                                                                                    default:
                                                                                                        switch fi30 {
                                                                                                        case -606:
                                                                                                            switch "" {
                                                                                                            case "":
                                                                                                                v52 -= 42
                                                                                                            case "foo bar":
                                                                                                                if false {
                                                                                                                    var v81 float32 = float32(v21)
                                                                                                                    _ = v81
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                fmt.Println(v3)
                                                                                                            }
                                                                                                        case 380:
                                                                                                            switch "\n\t" {
                                                                                                            case "\n\t":
                                                                                                                v82 := GenS10{F0: (((v54 >> ((18446744073709551615) & 63)) >= v54) || true), F1: (v19 << ((925) & 63))}
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            v83 := GenS11{F0: false, F1: 2}
                                                                                                            v3 -= (v3 / v3)
                                                                                                            _ = v83
                                                                                                        case -219:
                                                                                                            v2[fi30] = 4294967295
                                                                                                            v52 -= -2147483648
                                                                                                            if ((^(v54)) <= v54) {
                                                                                                                if false {
                                                                                                                    fmt.Println(v1, fi30, v54)
                                                                                                                } else {
                                                                                                                    var v84 bool = false
                                                                                                                    v85 := GenS3{F0: -446}
                                                                                                                    v86 := GenS12{F0: 601, F1: (v19 > int16(v52)), F2: 128}
                                                                                                                    _ = v84
                                                                                                                    _ = v85
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                                var v87 uint64 = v1
                                                                                                                _ = v87
                                                                                                            } else {
                                                                                                                switch "\n\t" {
                                                                                                                case "\n\t":
                                                                                                                    fmt.Println(v21, v52)
                                                                                                                case "a":
                                                                                                                    if len(v18) >= 2 {
                                                                                                                        v18 = v18[1:len(v18)]
                                                                                                                    }
                                                                                                                case "":
                                                                                                                    if ("foo bar" <= "a") {
                                                                                                                        v19 = -736
                                                                                                                    }
                                                                                                                }
                                                                                                                var v88 uint16 = 2
                                                                                                                var v89 uint32 = 42
                                                                                                                _ = v88
                                                                                                                _ = v89
                                                                                                            }
                                                                                                            var v90 uint16 = 235
                                                                                                            fmt.Println(fi30, v90)
                                                                                                            _ = v90
                                                                                                        default:
                                                                                                            var v91 int32 = v52
                                                                                                            _ = v91
                                                                                                        }
                                                                                                    }
                                                                                                case "foo bar":
                                                                                                    v92 := GenS13{F0: "foo bar", F1: v19, F2: 127}
                                                                                                    _ = v92
                                                                                                }
                                                                                                v93 := GenS1{F0: 2, F1: 18446744073709551615, F2: (v52 + 0)}
                                                                                                _ = v93
                                                                                            } else {
                                                                                                v54 -= (v54 + (v54 << ((42) & 63)))
                                                                                            }
                                                                                        }
                                                                                        if ((0.0 / 1.5) >= (1.5 + (0.1 * 1.5))) {
                                                                                            switch ("foo bar" + "\n\t") {
                                                                                            case "a":
                                                                                                var v94 uint8 = 2
                                                                                                _ = v94
                                                                                            case "hello":
                                                                                                if len(v18) >= 2 {
                                                                                                    v18 = v18[1:len(v18)]
                                                                                                }
                                                                                            default:
                                                                                                var v95 bool = (9223372036854775807 < 87)
                                                                                                _ = v95
                                                                                            }
                                                                                            fmt.Println(fi30)
                                                                                            switch "foo bar" {
                                                                                            case "foo bar":
                                                                                                v54 -= (^(v54))
                                                                                            case "hello":
                                                                                                var v96 int32 = -1
                                                                                                v97 := GenS8{F0: fi30, F1: 152, F2: -930}
                                                                                                v96 = (^(v96))
                                                                                                _ = v96
                                                                                                _ = v97
                                                                                            case "a":
                                                                                                var v98 int32 = v52
                                                                                                _ = v98
                                                                                            }
                                                                                        }
                                                                                        v99 := GenS9{F0: -9223372036854775808}
                                                                                        _ = v99
                                                                                    }
                                                                                    fmt.Println(v52, v52)
                                                                                    fmt.Println(v21)
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi30, v1)
                                                                        }
                                                                    } else {
                                                                        v19 -= -1
                                                                        if (false || true) {
                                                                            v100 := GenS14{F0: fi30}
                                                                            _ = v100
                                                                        }
                                                                        var v101 int64 = 100
                                                                        _ = v101
                                                                    }
                                                                    var v102 uint16 = 152
                                                                    _ = v102
                                                                } else {
                                                                    var v103 int = fi30
                                                                    _ = v103
                                                                }
                                                            }
                                                            v3 = v21
                                                        }
                                                        v18 = append(v18, (fi30 & -1))
                                                        _ = len(v18)
                                                    } else {
                                                        delete(v2, 32768)
                                                        var v104 int32 = (v52 - v52)
                                                        v105 := GenS15{F0: 12}
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    v106 := GenS16{F0: "foo bar", F1: (-(v54))}
                                                    fmt.Println(fi30, v21)
                                                    v107 := GenS12{F0: 1, F1: true, F2: 255}
                                                    fmt.Println(v19)
                                                    v108 := GenS2{F0: v52, F1: 207}
                                                    if len(v18) > 0 {
                                                        v18[0] = -570
                                                    }
                                                    delete(v2, fi26)
                                                    _ = v106
                                                    _ = v107
                                                    _ = v108
                                                }
                                                _ = v57
                                            } else {
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                                v109 := v2[fi30]
                                                _ = v109
                                            }
                                        default:
                                            v2[fi30] = 781
                                            if (true && true) {
                                                switch "a" {
                                                case "foo bar":
                                                    if len(v18) >= 2 {
                                                        v18 = v18[1:len(v18)]
                                                    }
                                                default:
                                                    var v110 uint16 = 1
                                                    _ = v110
                                                }
                                            }
                                            v19 = v19
                                        }
                                        _ = v54
                                    } else {
                                        var v111 uint = 216
                                        _ = v111
                                    }
                                case "foo bar":
                                    var v112 uint32 = 47
                                    _ = v112
                                case "\n\t":
                                    var v113 float32 = 0.1
                                    fmt.Println(v19)
                                    delete(v2, fi26)
                                    _ = v113
                                default:
                                    var v114 bool = true
                                    _ = v114
                                }
                                _ = v52
                            }
                        } else {
                            var v115 int16 = 898
                            _ = v115
                        }
                        v19 -= -32768
                        var v116 int16 = (v19 - -32768)
                        _ = v116
                    }
                    _ = fi30
                }
                v117 := GenS13{F0: "a", F1: (v19 & 183), F2: -494}
                _ = v117
            }
            var v118 int32 = -2147483648
            v1 = (v1 + 42)
            _ = fi26
            _ = v118
        }
    } else {
        var v119 uint64 = v1
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = v119
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
