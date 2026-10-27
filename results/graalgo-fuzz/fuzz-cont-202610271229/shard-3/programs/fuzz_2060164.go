package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint64; F2 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int64; F2 int }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint32; F1 uint8 }
type GenS11 struct { F0 uint32 }
func genFn0(p0 uint8, p1 int) int {
    v7 := []int{(p1 << ((0) & 63)), -1, 9223372036854775807, 573}
    for fi8 := 0; fi8 < 6; fi8++ {
        p0 -= 214
        fmt.Println(p1)
        _ = fi8
    }
    fmt.Println(p1, p1)
    if false {
        fmt.Println(p0, p1)
    } else {
        p1 += (p1 ^ 32767)
    }
    v7 = append(v7, 42)
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    var v3 float32 = (float32(1.5) + 0.0)
    v4 := make(map[string]bool)
    v5 := []uint32{65535}
    v6 := make(map[string]bool)
    var v9 uint = 4294967295
    fmt.Println(v9)
    switch "" {
    case "\n\t":
        var v10 int = 799
        _ = v10
    default:
        _ = len(v6)
    }
    var v11 uint = v9
    var v12 int64 = -9223372036854775808
    switch "hello" {
    case "foo bar":
        v13 := GenS1{F0: 528}
        _ = v13
    case "\n\t":
        delete(v4, (("foo bar" + ("" + "")) + ""))
    case "":
        switch -1 {
        case 1:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case 2147483647:
            switch -913 {
            case -474:
                v2 = true
            case 420:
                var v14 uint32 = 0
                _ = v14
            default:
                v15 := GenS1{F0: 2147483648}
                _ = v15
            }
        default:
            switch 128 {
            case -678:
                fmt.Println(v11, v9)
            }
        }
    default:
        v16 := GenS2{F0: 1, F1: 1, F2: -1}
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        switch 468 {
        case -134:
            v3 = -2.25
        case 342:
            if (false || v2) {
                switch "a" {
                case "":
                    fmt.Println(fi17)
                }
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 float64 = 1.5
                    switch fi17 {
                    case 450:
                        var v20 bool = (((65535 == 255) || v2) || false)
                        _ = v20
                    case -9223372036854775808:
                        fmt.Println(v9)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi18
                    _ = v19
                }
                var v21 int16 = 0
                _ = v21
            } else {
                var v22 float32 = 3.14159
                if (42 > 0) {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        fmt.Println(v22, fi23, v3)
                        var v24 int32 = -2147483648
                        _ = fi23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 7; fi25++ {
                        v12 = 757
                        if false {
                            v11 = v11
                            if (155 <= 184) {
                                v2 = (true || ((v2 || v2) || (!true)))
                            } else {
                                var v26 uint64 = 0
                                _ = v26
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v11 -= (^(v11))
                            case "a":
                                var v27 int64 = -734
                                _ = v27
                            }
                        }
                        var v28 uint = (^(v11))
                        _ = fi25
                        _ = v28
                    }
                } else {
                    v11 = (v11 + 0)
                }
                for fi29 := 0; fi29 < 3; fi29++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if v2 {
                        switch ("" + "foo bar") {
                        case "":
                            if (v2 || false) {
                                switch "a" {
                                case "foo bar":
                                    var v30 int8 = 42
                                    fmt.Println(v12, v11, v22)
                                    fmt.Println(fi17, fi29)
                                    _ = v30
                                case "a":
                                    switch (("" + ("\n\t" + "a")) + (("a" + ("foo bar" + "")) + "")) {
                                    case "a":
                                        v9 -= v11
                                    case "\n\t":
                                        var v31 uint32 = 0
                                        _ = v31
                                    case "foo bar":
                                        var v32 int64 = (v12 - v12)
                                        _ = v32
                                    default:
                                        var v33 bool = false
                                        _ = v33
                                    }
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                default:
                                    var v34 uint32 = 65535
                                    _ = v34
                                }
                                v22 = v3
                            }
                        case "hello":
                            if v2 {
                                fmt.Println(fi29, v9)
                            } else {
                                var v35 uint64 = 1
                                _ = v35
                            }
                        case "\n\t":
                            var v36 uint = (v9 >> ((v9) & 63))
                            _ = v36
                        }
                    }
                    _ = fi29
                }
                _ = v22
            }
        case 2147483647:
            for fi37 := 0; fi37 < 5; fi37++ {
                fmt.Println(v11, v9)
                if v2 {
                    var v38 uint32 = 42
                    _ = v38
                } else {
                    v39 := GenS3{F0: 42, F1: 42, F2: ((((("\n\t" + "a") + (("" + ("hello" + "")) + ("hello" + ""))) + "foo bar") + "foo bar") + ("\n\t" + "hello"))}
                    var v40 int64 = 905
                    v12 = v40
                    var v41 bool = v2
                    var v42 string = "foo bar"
                    switch 32768 {
                    case 2147483647:
                        fmt.Println(v40, v41, v9)
                    case -1:
                        switch v42 {
                        case "hello":
                            v43 := GenS4{F0: 366}
                            _ = v43
                        case "\n\t":
                            if (v2 && (!(!(((1e10 - 0.0) < (1e10 + float64(1e10))) || false)))) {
                                v44 := GenS5{F0: v42, F1: (v40 % ((551) | 1)), F2: int(v40)}
                                v9 += (v9 % ((v9) | 1))
                                var v45 float64 = 0.0
                                var v46 uint32 = 65535
                                if (!false) {
                                    if (!false) {
                                        _ = len(v6)
                                    } else {
                                        var v47 float32 = v3
                                        v12 = v12
                                        var v48 uint = v11
                                        _ = v47
                                        _ = v48
                                    }
                                    v11 = (v9 - 18446744073709551615)
                                } else {
                                    v9 = v11
                                    var v49 int64 = int64(fi37)
                                    _ = v49
                                }
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                        case "":
                            switch v42 {
                            case "hello":
                                var v50 uint32 = 65535
                                _ = v50
                            case "a":
                                switch v42 {
                                case "":
                                    v12 = (-(v40))
                                default:
                                    var v51 uint = v9
                                    _ = v51
                                }
                                v9 -= 4294967295
                            }
                        default:
                            if true {
                                fmt.Println(v2)
                                switch fi37 {
                                case 127:
                                    fmt.Println(v42, v3, v12)
                                    switch (((v42 + ("foo bar" + "\n\t")) + (v42 + v42)) + v42) {
                                    case "\n\t":
                                        var v52 uint16 = 65535
                                        _ = v52
                                    case "hello":
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                            }
                        }
                    case -9223372036854775808:
                        var v53 uint16 = 65535
                        switch "foo bar" {
                        case "":
                            switch v42 {
                            case "foo bar":
                                var v54 int = (fi37 / ((genFn0(1, (fi37 + fi37))) | 1))
                                var v55 int32 = 42
                                v9 += (v9 | (v11 | (^(v9))))
                                var v56 bool = (int8(v11) == -40)
                                _ = v54
                                _ = v55
                                _ = v56
                            case "\n\t":
                                var v57 int64 = v12
                                _ = v57
                            case "":
                                if true {
                                    if ((float64((v3 * float32((float64(1e10) - (-2.25 / 1.0))))) * (1e10 * 0.1)) > 3.14159) {
                                        var v58 int8 = 26
                                        var v59 float32 = v3
                                        fmt.Println(v2, v58, v41)
                                        _ = v58
                                        _ = v59
                                    }
                                    v60 := GenS4{F0: 0}
                                    var v61 uint16 = v53
                                    _ = v60
                                    _ = v61
                                }
                            default:
                                var v62 uint16 = 255
                                _ = v62
                            }
                        default:
                            v42 = v42
                        }
                        var v63 int64 = (v12 >> ((197) & 63))
                        _ = v53
                        _ = v63
                    }
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
                if (42 == -1) {
                    v64 := v6["\n\t"]
                    var v65 string = ("" + "\n\t")
                    _ = v64
                    _ = v65
                } else {
                    if v2 {
                        var v66 uint8 = 1
                        switch genFn0((v66 + v66), genFn0(v66, (fi17 - fi37))) {
                        case 0:
                            fmt.Println(v3, fi17, v11)
                        case 32767:
                            if v2 {
                                if len(v5) > 0 {
                                    v5[0] = 326
                                }
                            }
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                            }
                        }
                        v2 = (float64((float32((3.14159 / (((float64(((v3 + float32(1e10)) - (v3 + (v3 - v3)))) * 0.0) * ((0.0 / 1.0) * 1e10)) + 1e10))) - 1e10)) <= 1e10)
                        fmt.Println(fi17, fi37)
                        _ = len(v5)
                        _ = v66
                    } else {
                        var v67 int = int(v9)
                        _ = v67
                    }
                    v2 = v2
                    var v68 uint16 = 688
                    _ = v68
                }
                var v69 int32 = 2147483647
                v69 = 1
                var v70 uint64 = 664
                v3 -= (v3 - (((v3 * ((((v3 * v3) / 0.1) * float32(1e10)) - 1e10)) + v3) / (1e10 / v3)))
                _ = fi37
                _ = v69
                _ = v70
            }
        }
        switch fi17 {
        case 537:
            v3 = ((0.0 + 3.14159) - 3.14159)
        case -1:
            var v71 int = (fi17 / (((fi17 * 2147483647)) | 1))
            _ = v71
        }
        if v2 {
            v72 := GenS5{F0: "\n\t", F1: -1, F2: 792}
            v73 := v6["a"]
            switch ("\n\t" + ("" + ("a" + "foo bar"))) {
            case "\n\t":
                for fi74 := 0; fi74 < 11; fi74++ {
                    v3 -= float32((1e10 * 1e10))
                    switch (fi17 * (^(fi74))) {
                    case -9223372036854775808:
                        v75 := GenS3{F0: 255, F1: 0, F2: "foo bar"}
                        _ = len(v5)
                        v2 = v2
                        _ = v75
                    case -2147483648:
                        var v76 float64 = ((1e10 * 1e10) - (0.1 - ((3.14159 - -2.25) + 0.1)))
                        _ = v76
                    default:
                        v5 = append(v5, 35)
                    }
                    _ = fi74
                }
            case "hello":
                for fi77 := 0; fi77 < 3; fi77++ {
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                    if true {
                        v78 := GenS6{F0: 4294967295, F1: 4294967295}
                        if (fi77 == 391) {
                            v79 := GenS7{F0: 2, F1: 157, F2: 100}
                            var v80 uint32 = 962
                            v81 := GenS6{F0: 18446744073709551615, F1: 183}
                            _ = v79
                            _ = v80
                            _ = v81
                        } else {
                            var v82 int8 = -1
                            v83 := GenS8{F0: v82, F1: 255}
                            var v84 int32 = -2147483648
                            if (!v73) {
                                var v85 uint8 = 0
                                v86 := GenS9{F0: int16(v82), F1: ("foo bar" + "foo bar")}
                                switch int(v82) {
                                case 526:
                                    if v2 {
                                        var v87 float64 = 1.5
                                        switch fi17 {
                                        case 128:
                                            v88 := GenS4{F0: 741}
                                            _ = v88
                                        case 127:
                                            v85 = v85
                                        case -9223372036854775808:
                                            v82 -= 127
                                            v82 -= -109
                                            delete(v4, ("a" + "\n\t"))
                                        default:
                                            var v89 int16 = -1
                                            _ = v89
                                        }
                                        if ((v85 <= (v85 & (v85 & (^(v85))))) && v2) {
                                            if ((!true) && (!(0 > 2))) {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                _ = len(v1)
                                            }
                                            v90 := v1[508]
                                            _ = v90
                                        } else {
                                            var v91 string = ((("a" + (("foo bar" + "") + (("foo bar" + "foo bar") + "hello"))) + ("hello" + "hello")) + "a")
                                            var v92 bool = false
                                            v9 = 523
                                            _ = v91
                                            _ = v92
                                        }
                                        _ = v87
                                    } else {
                                        if (true || v2) {
                                            var v93 int64 = (v12 % ((v12) | 1))
                                            var v94 uint = 920
                                            v12 += 311
                                            _ = v93
                                            _ = v94
                                        }
                                        var v95 int8 = v82
                                        var v96 uint16 = 65535
                                        v1[int(v12)] = 65535
                                        fmt.Println(v82)
                                        _ = v95
                                        _ = v96
                                    }
                                }
                                _ = v85
                                _ = v86
                            } else {
                                v2 = false
                                v2 = (!(!false))
                            }
                            v97 := GenS2{F0: 4294967295, F1: 42, F2: (v84 << ((v9) & 63))}
                            _ = v82
                            _ = v83
                            _ = v84
                            _ = v97
                        }
                        v98 := v6["\n\t"]
                        _ = v78
                        _ = v98
                    } else {
                        fmt.Println(v11, v12)
                        v12 = (v12 ^ (v12 | 0))
                    }
                    if v2 {
                        var v99 bool = (!v2)
                        if true {
                            v100 := v1[(fi17 % ((fi17) | 1))]
                            _ = v100
                        } else {
                            var v101 float64 = -2.25
                            var v102 uint16 = 255
                            v101 -= 0.0
                            delete(v4, "")
                            var v103 bool = (v2 && (((v12 + (v12 / (((v12 - v12)) | 1))) != v12) || (89 > int8(v102))))
                            _ = v101
                            _ = v102
                            _ = v103
                        }
                        var v104 int = fi17
                        _ = v104
                        _ = v99
                    }
                    _ = fi77
                }
            }
            _ = v72
            _ = v73
        }
        _ = fi17
    }
    v3 += float32((3.14159 / 3.14159))
    v3 -= float32((-2.25 - 0.1))
    _ = len(v6)
    v12 = 0
    var v105 float32 = float32((1e10 * 1e10))
    for fi106 := 0; fi106 < 6; fi106++ {
        switch "hello" {
        case "hello":
            if (4294967295 != 896) {
                for fi107 := 0; fi107 < 5; fi107++ {
                    _ = len(v5)
                    if (693 != 2147483647) {
                        v9 = v9
                        switch fi107 {
                        case -9223372036854775808:
                            var v108 uint = (v11 % ((458) | 1))
                            _ = v108
                        case -974:
                            delete(v4, ("" + "hello"))
                        case 1:
                            v109 := GenS7{F0: 232, F1: 65535, F2: 127}
                            _ = v109
                        }
                    }
                    _ = fi107
                }
                for fi110 := 0; fi110 < 6; fi110++ {
                    if ((161 < 255) || ((!v2) || true)) {
                        v5 = append(v5, 629)
                        v111 := GenS2{F0: 4294967295, F1: -32768, F2: -499}
                        v9 -= v11
                        _ = v111
                    } else {
                        v3 = 3.14159
                    }
                    v6[("\n\t" + ("foo bar" + (("hello" + ("" + "foo bar")) + "\n\t")))] = v2
                    _ = fi110
                }
            }
        default:
            switch (fi106 & fi106) {
            case 4294967295:
                fmt.Println(v11, v11)
            case 548:
                delete(v4, "hello")
            case 32768:
                fmt.Println(v2)
            }
        }
        switch "\n\t" {
        case "":
            var v112 uint64 = 18446744073709551615
            _ = v112
        default:
            if (!v2) {
                fmt.Println(v105)
                if v2 {
                    var v113 float32 = (-2.25 + 0.0)
                    switch ("foo bar" + ("foo bar" + "a")) {
                    case "":
                        v114 := GenS10{F0: uint32(v9), F1: 2}
                        _ = v114
                    default:
                        var v115 float64 = 1e10
                        _ = v115
                    }
                    switch fi106 {
                    case 45:
                        delete(v6, (((((("a" + "") + ("" + "foo bar")) + "foo bar") + "hello") + "") + (("foo bar" + "") + "hello")))
                    }
                    _ = v113
                }
            }
        }
        _ = fi106
    }
    var v116 int = 0
    fmt.Println(v2, v9)
    if true {
        var v117 uint8 = 166
        var v118 uint = (^(v11))
        _ = v117
        _ = v118
    } else {
        var v119 int64 = v12
        if (!(3.14159 >= ((0.0 * (-2.25 / 1.0)) - 0.0))) {
            v120 := GenS11{F0: 4294967295}
            _ = v120
        }
        v119 = v12
        _ = v119
    }
    v121 := v6[("" + "a")]
    if v121 {
        _ = len(v5)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v11:", v11)
    fmt.Println("v116:", v116)
    fmt.Println("v12:", v12)
    fmt.Println("v121:", v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v105
    _ = v11
    _ = v116
    _ = v12
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
