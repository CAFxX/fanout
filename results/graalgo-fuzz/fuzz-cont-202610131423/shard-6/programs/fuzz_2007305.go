package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16; F2 string }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int64; F1 int32; F2 bool }
type GenS11 struct { F0 int8; F1 int32; F2 int8 }
type GenS12 struct { F0 uint32; F1 int; F2 int16 }
type GenS13 struct { F0 string; F1 uint64; F2 int }
type GenS14 struct { F0 bool; F1 uint32 }
type GenS15 struct { F0 string; F1 bool }
type GenS16 struct { F0 uint8; F1 uint64 }
type GenS17 struct { F0 int16; F1 string }
func genFn0(p0 uint32, p1 uint32) int8 {
    v6 := []int{}
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = -2.25
        var v9 bool = true
        fmt.Println(p0)
        _ = fi7
        _ = v8
        _ = v9
    }
    switch 133 {
    case 179:
        var v10 float32 = float32(1e10)
        _ = v10
    case 251:
        var v11 float32 = 1e10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}
func genFn1(p0 bool, p1 uint) int64 {
    p1 += p1
    v12 := []int{9223372036854775807, 128}
    if len(v12) > 0 {
        v12[0] = -2147483648
    }
    _ = p0
    _ = p1
    _ = v12
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 bool = ((false || true) && ((1 <= 127) || true))
    v4 := []uint32{2, 0, 4294967295}
    var v5 int32 = -449
    v5 = v5
    v4 = append(v4, 79)
    if v3 {
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 255}
                fmt.Println(v3)
                v5 = (v5 % ((2147483647) | 1))
                _ = fi14
                _ = v15
            }
            v3 = v3
            for fi16 := 0; fi16 < 12; fi16++ {
                if v3 {
                    if v3 {
                        v17 := GenS1{F0: 0}
                        v5 += (v5 % ((int32(fi13)) | 1))
                        _ = v17
                    } else {
                        fmt.Println(v5)
                    }
                    var v18 uint = 282
                    _ = v18
                } else {
                    fmt.Println(v3, v5)
                }
                if (fi13 >= (-(fi16))) {
                    v5 = (v5 - (v5 % (((v5 >> ((255) & 63))) | 1)))
                    if true {
                        var v19 string = "a"
                        var v20 uint = 855
                        v20 += 423
                        _ = v19
                        _ = v20
                    } else {
                        var v21 int8 = 1
                        v4 = append(v4, 2147483648)
                        var v22 uint16 = 0
                        v5 = -1
                        fmt.Println(v22, fi16)
                        if v3 {
                            if v3 {
                                v3 = v3
                                delete(v1, 251)
                                fmt.Println(v22, v3, fi13)
                                var v23 int64 = -700
                                if v3 {
                                    v4 = append(v4, 37)
                                    if (true || v3) {
                                        v24 := GenS1{F0: 1}
                                        if (840 >= 18446744073709551615) {
                                            v5 -= (v5 >> ((0) & 63))
                                        } else {
                                            if (32768 > (fi13 - fi16)) {
                                                switch (fi16 % ((fi13) | 1)) {
                                                case -996:
                                                    fmt.Println(v22, v23)
                                                default:
                                                    var v25 int = -9223372036854775808
                                                    _ = v25
                                                }
                                                v22 += (v22 + v22)
                                                var v26 uint64 = 2
                                                v27 := GenS2{F0: 183}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 uint8 = 189
                                                v21 -= (v21 / ((v21) | 1))
                                                if ((42 == -128) || v3) {
                                                    v3 = false
                                                    var v29 uint8 = (v28 ^ v28)
                                                    _ = v29
                                                }
                                                _ = v28
                                            }
                                        }
                                        var v30 float32 = 1.5
                                        _ = v24
                                        _ = v30
                                    } else {
                                        v2["a"] = fi16
                                        var v31 bool = (!v3)
                                        _ = v31
                                    }
                                }
                                var v32 uint8 = 2
                                _ = v23
                                _ = v32
                            }
                            if ((fi16 | (-(fi16))) >= fi16) {
                                if v3 {
                                    if (-32768 < 1) {
                                        v33 := GenS3{F0: 1, F1: 984, F2: "hello"}
                                        fmt.Println(fi13)
                                        if len(v4) > 0 {
                                            v4[0] = 2
                                        }
                                        _ = v33
                                    }
                                    _ = len(v1)
                                }
                                if len(v4) > 0 {
                                    v4[0] = 473
                                }
                                var v34 uint16 = (v22 ^ v22)
                                _ = v34
                            } else {
                                if v3 {
                                    var v35 float32 = float32(-2.25)
                                    if false {
                                        if (!false) {
                                            switch "\n\t" {
                                            case "a":
                                                switch "\n\t" {
                                                case "hello":
                                                    if v3 {
                                                        v3 = (3.14159 > 0.0)
                                                        var v36 uint64 = 1
                                                        if (-474 >= -9223372036854775808) {
                                                            switch int(v22) {
                                                            case 32768:
                                                                switch "foo bar" {
                                                                case "":
                                                                    switch fi16 {
                                                                    case -58:
                                                                        if (0 >= 259) {
                                                                            if v3 {
                                                                                var v37 int32 = v5
                                                                                v22 -= 425
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    var v38 uint16 = (v22 + 32767)
                                                                                    v39 := GenS4{F0: v3, F1: (^(v22))}
                                                                                    v3 = (!((v21 | v21) == (-(v21))))
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                v36 = (v36 & v36)
                                                                                _ = v37
                                                                            } else {
                                                                                v5 += (v5 / ((974) | 1))
                                                                            }
                                                                            switch fi16 {
                                                                            case 128:
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            }
                                                                            var v40 uint = 18446744073709551615
                                                                            _ = v40
                                                                        }
                                                                    case 4294967295:
                                                                        var v41 uint16 = v22
                                                                        _ = v41
                                                                    case 701:
                                                                        _ = len(v2)
                                                                    default:
                                                                        var v42 uint8 = 42
                                                                        _ = v42
                                                                    }
                                                                case "hello":
                                                                    fmt.Println(v3, v22)
                                                                }
                                                                v21 += -1
                                                                var v43 int64 = 42
                                                                switch ("foo bar" + "foo bar") {
                                                                case "\n\t":
                                                                    var v44 int = 32767
                                                                    _ = v44
                                                                case "hello":
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v45 int = fi16
                                                                        _ = v45
                                                                    case "foo bar":
                                                                        var v46 uint = 149
                                                                        _ = v46
                                                                    default:
                                                                        var v47 int32 = (v5 - v5)
                                                                        _ = v47
                                                                    }
                                                                case "foo bar":
                                                                    var v48 float64 = float64((((float32((3.14159 / 3.14159)) / v35) / 1.0) - (v35 / v35)))
                                                                    _ = v48
                                                                }
                                                                v5 = -901
                                                                _ = v43
                                                            case 168:
                                                                v49 := GenS3{F0: 397, F1: 92, F2: (("foo bar" + "\n\t") + "a")}
                                                                _ = v49
                                                            case -2147483648:
                                                                v4 = append(v4, 65535)
                                                            }
                                                            switch fi13 {
                                                            case 9223372036854775807:
                                                                var v50 string = ("a" + ((("a" + "foo bar") + "hello") + "a"))
                                                                if false {
                                                                    v51 := GenS5{F0: 65535}
                                                                    switch "" {
                                                                    case "":
                                                                        delete(v1, 842)
                                                                    }
                                                                    _ = v51
                                                                } else {
                                                                    v36 -= (^(v36))
                                                                    delete(v1, 2147483647)
                                                                    var v52 uint8 = 128
                                                                    _ = v52
                                                                }
                                                                _ = v50
                                                            }
                                                            v21 = genFn0(2147483648, 42)
                                                        }
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            v53 := GenS6{F0: (v5 + v5), F1: (v36 >> ((0) & 63)), F2: 9223372036854775807}
                                                            _ = v53
                                                        case -49:
                                                            if true {
                                                                if v3 {
                                                                    var v54 uint64 = v36
                                                                    switch ("hello" + "foo bar") {
                                                                    case "\n\t":
                                                                        var v55 int64 = 466
                                                                        _ = v55
                                                                    case "foo bar":
                                                                        v4 = append(v4, 42)
                                                                    }
                                                                    _ = v54
                                                                } else {
                                                                    v36 = v36
                                                                }
                                                                v36 += v36
                                                                v21 = (v21 & v21)
                                                            }
                                                        case -1:
                                                            var v56 uint = 863
                                                            _ = v56
                                                        default:
                                                            if (-1 < -1) {
                                                                var v57 uint32 = 662
                                                                var v58 uint8 = 222
                                                                v59 := GenS7{F0: (^(v57))}
                                                                _ = v57
                                                                _ = v58
                                                                _ = v59
                                                            }
                                                        }
                                                        if len(v4) > 0 {
                                                            v4[0] = 1
                                                        }
                                                        _ = v36
                                                    } else {
                                                        fmt.Println(v5)
                                                        var v60 uint64 = 0
                                                        v22 -= v22
                                                        _ = v60
                                                    }
                                                case "":
                                                    if ((1e10 + 1.5) >= (float64(-2.25) * -2.25)) {
                                                        v61 := GenS3{F0: 0, F1: 829, F2: "\n\t"}
                                                        v62 := GenS8{F0: 2147483648, F1: "foo bar"}
                                                        fmt.Println(v21, v22, fi16)
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                default:
                                                    v3 = ((42 == 42) || (42 <= 18446744073709551615))
                                                }
                                            case "":
                                                v63 := v2[""]
                                                _ = v63
                                            }
                                            v3 = false
                                            switch 127 {
                                            case 132:
                                                v64 := GenS1{F0: 255}
                                                _ = v64
                                            default:
                                                fmt.Println(v22, v3)
                                            }
                                            var v65 bool = v3
                                            _ = v65
                                        }
                                        v3 = (2 > uint(v22))
                                        switch fi13 {
                                        case -941:
                                            v3 = (((((("a" + (("hello" + "\n\t") + ((((("hello" + ("\n\t" + ("a" + ("\n\t" + ("foo bar" + (("a" + (("" + "") + "hello")) + ("hello" + "hello"))))))) + "") + "") + "") + "\n\t"))) + "") + "foo bar") + ("" + (("\n\t" + ("\n\t" + ("" + "foo bar"))) + ("" + "a")))) != "") || (v3 || v3))
                                        default:
                                            var v66 uint16 = (v22 | (^(v22)))
                                            var v67 float64 = 1e10
                                            _ = v66
                                            _ = v67
                                        }
                                    }
                                    _ = v35
                                } else {
                                    var v68 float32 = 3.14159
                                    v3 = false
                                    v68 = (0.0 - (v68 / 0.1))
                                    v69 := GenS9{F0: uint(v22)}
                                    _ = v68
                                    _ = v69
                                }
                                if false {
                                    v21 += genFn0(2147483648, 1)
                                    if v3 {
                                        var v70 int8 = (v21 ^ (v21 << ((179) & 63)))
                                        v5 -= v5
                                        var v71 float32 = (((-2.25 * -2.25) / (3.14159 / 3.14159)) / 1.0)
                                        v72 := v1[uint32(v22)]
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    } else {
                                        fmt.Println(fi13, fi16)
                                        var v73 uint32 = 390
                                        if false {
                                            var v74 int16 = -32768
                                            _ = v74
                                        } else {
                                            v22 += (^(v22))
                                            var v75 int = (fi13 | (fi13 << ((878) & 63)))
                                            _ = v75
                                        }
                                        _ = v73
                                    }
                                } else {
                                    switch fi16 {
                                    case 0:
                                        v3 = (255 <= (^(v22)))
                                    case 535:
                                        switch "foo bar" {
                                        case "hello":
                                            switch int(v22) {
                                            case 127:
                                                v21 -= -90
                                            case 42:
                                                if v3 {
                                                    switch "\n\t" {
                                                    case "a":
                                                        fmt.Println(fi13)
                                                    default:
                                                        v4 = append(v4, 42)
                                                    }
                                                    if v3 {
                                                        v76 := GenS10{F0: -1, F1: 42, F2: v3}
                                                        _ = v76
                                                    } else {
                                                        if (1.5 != ((1e10 + (float32(float64(0.0)) * 0.1)) + -2.25)) {
                                                            v77 := GenS1{F0: 162}
                                                            v78 := GenS2{F0: 128}
                                                            if false {
                                                                fmt.Println(v22, v21)
                                                            } else {
                                                                v3 = ((127 != 0) || (-2147483648 != 0))
                                                                fmt.Println(v22)
                                                                v5 += 2147483647
                                                            }
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        switch (fi16 * (fi13 << ((132) & 63))) {
                                                        case 2147483647:
                                                            v79 := GenS5{F0: v22}
                                                            _ = v79
                                                        }
                                                        var v80 string = ("a" + "foo bar")
                                                        v5 = v5
                                                        _ = v80
                                                    }
                                                } else {
                                                    v81 := GenS11{F0: (v21 << ((0) & 63)), F1: v5, F2: v21}
                                                    switch fi16 {
                                                    case 9223372036854775807:
                                                        var v82 uint64 = 42
                                                        _ = v82
                                                    }
                                                    v5 -= (v5 | -561)
                                                    v83 := GenS12{F0: 663, F1: 9223372036854775807, F2: -32768}
                                                    _ = v81
                                                    _ = v83
                                                }
                                            }
                                        case "\n\t":
                                            v84 := GenS13{F0: "a", F1: 0, F2: fi13}
                                            var v85 uint8 = 255
                                            var v86 bool = v3
                                            fmt.Println(v86, v3, v5)
                                            _ = v84
                                            _ = v85
                                            _ = v86
                                        default:
                                            fmt.Println(v21)
                                        }
                                    case -856:
                                        v3 = (42 == -872)
                                    default:
                                        fmt.Println(fi16)
                                    }
                                    fmt.Println(v21)
                                }
                                v21 = (v21 / (((-(v21))) | 1))
                                var v87 uint32 = 0
                                if v3 {
                                    v88 := GenS14{F0: (v3 || (true && (v21 < v21))), F1: (v87 ^ (v87 >> ((4294967295) & 63)))}
                                    var v89 float32 = 1.5
                                    _ = v88
                                    _ = v89
                                }
                                var v90 uint8 = 0
                                var v91 uint32 = (v87 - v87)
                                var v92 int16 = 158
                                if (v92 <= v92) {
                                    var v93 float32 = 1e10
                                    _ = v93
                                } else {
                                    if (((-2.25 * 1.5) - (float64(-2.25) - (1.5 / 1.5))) < (3.14159 / (0.1 / (1.5 * (float64((1e10 / 1.5)) + -2.25))))) {
                                        fmt.Println(v87)
                                    } else {
                                        v4 = append(v4, v87)
                                        switch "" {
                                        case "a":
                                            v94 := GenS15{F0: ("\n\t" + ""), F1: (v3 || (float64(-2.25) > (-2.25 * 1e10)))}
                                            _ = v94
                                        default:
                                            switch "hello" {
                                            case "foo bar":
                                                v90 = v90
                                            case "\n\t":
                                                v92 += v92
                                                v4 = append(v4, 621)
                                                v21 = 36
                                            case "hello":
                                                var v95 uint32 = 0
                                                _ = v95
                                            }
                                        }
                                        var v96 uint64 = 151
                                        _ = v96
                                    }
                                }
                                _ = v87
                                _ = v90
                                _ = v91
                                _ = v92
                            }
                            var v97 int64 = 0
                            _ = v97
                        } else {
                            var v98 string = ("\n\t" + ("hello" + "a"))
                            var v99 uint64 = 191
                            var v100 int8 = (v21 + (v21 - (v21 / (((v21 + genFn0(2147483647, 383))) | 1))))
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                        var v101 float64 = 1.5
                        _ = v101
                        _ = v21
                        _ = v22
                    }
                    _ = len(v4)
                    if v3 {
                        var v102 uint8 = 176
                        v3 = v3
                        if true {
                            var v103 uint32 = 4294967295
                            v3 = v3
                            _ = v103
                        }
                        _ = v102
                    } else {
                        switch ("a" + "") {
                        case "foo bar":
                            var v104 float32 = 0.1
                            _ = v104
                        case "\n\t":
                            v105 := GenS16{F0: 1, F1: 18446744073709551615}
                            _ = v105
                        }
                        if (127 != 1) {
                            if v3 {
                                var v106 float64 = (0.1 * -2.25)
                                _ = v106
                            }
                            fmt.Println(v3)
                        } else {
                            var v107 string = "hello"
                            var v108 int16 = 604
                            v5 = v5
                            switch "\n\t" {
                            case "foo bar":
                                v109 := GenS17{F0: -32768, F1: "hello"}
                                _ = v109
                            default:
                                v3 = v3
                            }
                            fmt.Println(fi13, v5)
                            _ = v107
                            _ = v108
                        }
                    }
                    fmt.Println(fi16, v5)
                }
                v5 -= -691
                _ = fi16
            }
            _ = fi13
        }
        if v3 {
            var v110 float32 = -2.25
            _ = v110
        } else {
            fmt.Println(v3, v5)
        }
    }
    var v111 uint64 = 413
    var v112 int16 = 42
    var v113 float32 = (3.14159 / 1.0)
    v112 -= (^(v112))
    v114 := v1[4294967295]
    for fi115 := 0; fi115 < 7; fi115++ {
        if v3 {
            fmt.Println(v113, v114, v3)
            v111 += (v111 * (v111 * v111))
        }
        v3 = (!(v112 >= (v112 << ((42) & 63))))
        v5 -= (v5 / (((v5 + (v5 | v5))) | 1))
        _ = fi115
    }
    var v116 int64 = -9223372036854775808
    var v117 int32 = (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
