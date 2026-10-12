package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint32 }
type GenS7 struct { F0 int; F1 uint32; F2 bool }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS13 struct { F0 int; F1 string }
type GenS14 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS15 struct { F0 int64; F1 uint }
type GenS16 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS17 struct { F0 int64; F1 int32 }
func genFn0(p0 uint8, p1 uint32) string {
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 0}
        v6 := GenS1{F0: -9223372036854775808}
        v7 := []int{127}
        _ = fi4
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: "hello", F1: -655}
    v9 := []int{42, 32768, 2147483647, 32767}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "\n\t"
}
func genFn1(p0 uint) int {
    var v10 uint8 = 2
    if (-207 == -2147483648) {
        v11 := GenS2{F0: ("a" + "\n\t"), F1: 32767}
        _ = v11
    }
    v12 := []int{-9223372036854775808, 44, 0}
    v12 = append(v12, 755)
    if false {
        v13 := GenS3{F0: 127}
        _ = v13
    } else {
        v10 = 42
        for fi14 := 0; fi14 < 9; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint16 = 255
                v17 := GenS4{F0: (false && ((((!(!(false && ((!false) || true)))) || (42 < 913)) && (!(!(-128 > -128)))) || true)), F1: -425}
                var v18 int16 = -1
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            v12 = append(v12, (fi14 * fi14))
            var v19 int32 = 1
            _ = fi14
            _ = v19
        }
    }
    _ = p0
    _ = v10
    _ = v12
    return 0
}

func main() {
    var v1 bool = (true && true)
    var v2 int64 = 160
    var v3 int16 = 353
    v20 := []int{}
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 19; fi22++ {
            if (((("foo bar" + ("a" + "\n\t")) + "") + "\n\t") == ((("foo bar" + "\n\t") + ("a" + (("\n\t" + ("a" + "a")) + "\n\t"))) + ("\n\t" + ""))) {
                if v1 {
                    if v1 {
                        var v23 uint16 = 2
                        v24 := GenS5{F0: genFn1(75)}
                        v3 = v3
                        _ = v23
                        _ = v24
                    }
                    _ = len(v20)
                } else {
                    v2 -= (v2 + v2)
                    var v25 uint32 = 803
                    v26 := GenS4{F0: v1, F1: v3}
                    _ = v25
                    _ = v26
                }
                switch "hello" {
                case "\n\t":
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v20)
                    case "foo bar":
                        fmt.Println(fi21, fi22, fi22)
                    case "a":
                        var v27 bool = true
                        fmt.Println(v1, v1)
                        switch fi21 {
                        case -199:
                            var v28 int64 = v2
                            _ = v28
                        case 0:
                            var v29 int16 = -562
                            _ = v29
                        default:
                            _ = len(v20)
                        }
                        v30 := GenS2{F0: "\n\t", F1: -1}
                        _ = v27
                        _ = v30
                    }
                    fmt.Println(v1, fi21)
                }
            } else {
                var v31 uint16 = 0
                v20 = append(v20, 128)
                v2 += -1
                if (v1 && (127 == 29)) {
                    if v1 {
                        var v32 int = fi22
                        if len(v20) > 0 {
                            v20[0] = (-(v32))
                        }
                        _ = v32
                    } else {
                        v20 = append(v20, fi22)
                        var v33 float32 = 1.5
                        _ = v33
                    }
                    var v34 bool = v1
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = v34
                }
                v31 -= (^(v31))
                _ = v31
            }
            _ = fi22
        }
        for fi35 := 0; fi35 < 3; fi35++ {
            var v36 int16 = (v3 + v3)
            if (v1 || v1) {
                var v37 bool = false
                _ = v37
            } else {
                var v38 int16 = (v3 >> ((0) & 63))
                _ = v38
            }
            if (!true) {
                v39 := GenS2{F0: ("\n\t" + ""), F1: (v3 + 42)}
                if (fi21 <= genFn1(348)) {
                    v36 = v36
                    switch ("foo bar" + "foo bar") {
                    case "":
                        switch (fi21 * (fi35 - genFn1(18446744073709551615))) {
                        case 0:
                            if (v1 || (!((0.1 != (float32(((3.14159 + (0.1 / -2.25)) - (float64(float32(1.5)) + 1e10))) * (-2.25 - float32(1e10)))) || (!v1)))) {
                                v2 = (^(v2))
                            } else {
                                var v40 uint8 = 91
                                var v41 int8 = 127
                                var v42 int = (-(fi35))
                                _ = len(v20)
                                v20 = append(v20, (fi21 ^ -9223372036854775808))
                                v36 -= v36
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        case -2147483648:
                            var v43 uint32 = 463
                            _ = v43
                        default:
                            if v1 {
                                if v1 {
                                    v44 := GenS5{F0: (fi21 | fi35)}
                                    v45 := GenS6{F0: 2, F1: 42}
                                    if len(v20) > 0 {
                                        v20[0] = 127
                                    }
                                    _ = v44
                                    _ = v45
                                }
                                if ((0 > 65535) && (9223372036854775807 <= v2)) {
                                    var v46 uint16 = 255
                                    _ = v46
                                } else {
                                    var v47 int16 = 32767
                                    _ = v47
                                }
                                v48 := GenS7{F0: genFn1(655), F1: 0, F2: v1}
                                _ = v48
                            }
                        }
                    case "\n\t":
                        var v49 uint64 = 1
                        _ = v49
                    case "a":
                        var v50 int8 = 1
                        _ = v50
                    }
                } else {
                    if len(v20) > 0 {
                        v20[0] = int(v2)
                    }
                }
                var v51 uint64 = 1
                _ = v39
                _ = v51
            } else {
                v3 = -32768
                v20 = append(v20, fi21)
                if v1 {
                    if (213 > 42) {
                        if len(v20) > 0 {
                            v20[0] = fi21
                        }
                        var v52 float32 = 3.14159
                        _ = v52
                    } else {
                        v53 := GenS2{F0: "a", F1: (v3 % (((v36 - v36)) | 1))}
                        if v1 {
                            switch 54 {
                            case -158:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            case 32767:
                                if true {
                                    var v54 string = "\n\t"
                                    _ = v54
                                }
                            case -896:
                                var v55 int = (fi35 ^ genFn1(42))
                                _ = v55
                            }
                        }
                        v3 += v3
                        _ = v53
                    }
                    var v56 int = (fi35 & genFn1(4294967295))
                    var v57 string = "a"
                    switch (v56 * fi35) {
                    case 1:
                        _ = len(v20)
                    case 128:
                        v56 += fi21
                    default:
                        v20 = append(v20, v56)
                    }
                    fmt.Println(v1)
                    _ = v56
                    _ = v57
                } else {
                    var v58 bool = (37 <= 0)
                    _ = v58
                }
                v59 := GenS1{F0: 1}
                v60 := GenS8{F0: 4294967295, F1: 282}
                _ = v59
                _ = v60
            }
            _ = fi35
            _ = v36
        }
        _ = fi21
    }
    v3 -= 865
    if v1 {
        fmt.Println(v1, v2, v3)
        fmt.Println(v3)
    } else {
        v2 -= -917
    }
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    v3 += 32767
    switch "a" {
    case "\n\t":
        var v61 int16 = -731
        _ = v61
    case "":
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v3, v3, v2)
        default:
            _ = len(v20)
        }
    }
    v62 := GenS5{F0: -2147483648}
    for fi63 := 0; fi63 < 9; fi63++ {
        v1 = ((v3 / (((v3 - (v3 ^ -1))) | 1)) == (v3 ^ v3))
        for fi64 := 0; fi64 < 20; fi64++ {
            if v1 {
                var v65 int = 42
                v66 := GenS9{F0: 1}
                _ = v65
                _ = v66
            } else {
                var v67 int32 = -209
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                v67 = (-(v67))
                _ = v67
            }
            v68 := GenS1{F0: (v2 & v2)}
            var v69 bool = (-542 <= (v2 & v2))
            _ = fi64
            _ = v68
            _ = v69
        }
        for fi70 := 0; fi70 < 13; fi70++ {
            if false {
                v71 := GenS1{F0: v2}
                _ = v71
            } else {
                if (1 <= 18446744073709551615) {
                    if (131 < 1) {
                        if (0.0 < 0.1) {
                            if (!((false || true) && v1)) {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                if (4294967295 < 65535) {
                                    var v72 uint8 = 0
                                    fmt.Println(v72, v1)
                                    _ = v72
                                } else {
                                    var v73 uint = 107
                                    v20 = append(v20, (fi70 - fi70))
                                    _ = v73
                                }
                                var v74 float32 = ((1.5 - ((0.1 - 3.14159) / 1.5)) + -2.25)
                                _ = v74
                            }
                            v1 = (1 <= -128)
                            var v75 int64 = (v2 / ((v2) | 1))
                            _ = v75
                        }
                        v76 := GenS3{F0: -1}
                        _ = v76
                    }
                    if (!v1) {
                        v3 -= (^(v3))
                    } else {
                        v77 := GenS10{F0: 18446744073709551615}
                        if (!((!v1) && v1)) {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            v2 -= (v2 ^ 42)
                            var v78 uint32 = 4294967295
                            _ = v78
                        }
                        _ = v77
                    }
                }
                if v1 {
                    if len(v20) > 0 {
                        v20[0] = fi63
                    }
                    v2 -= 1
                    var v79 float32 = 1.5
                    _ = v79
                } else {
                    var v80 uint = 785
                    v1 = false
                    _ = v80
                }
                _ = len(v20)
            }
            switch ("a" + "hello") {
            case "hello":
                switch "hello" {
                case "hello":
                    v81 := GenS11{F0: 1}
                    _ = v81
                default:
                    if (!false) {
                        var v82 int = 0
                        _ = v82
                    } else {
                        switch fi63 {
                        case 32768:
                            if (!v1) {
                                var v83 int = fi70
                                if (!true) {
                                    switch "a" {
                                    case "\n\t":
                                        v84 := GenS12{F0: 1, F1: 488, F2: 42}
                                        _ = v84
                                    case "":
                                        if true {
                                            var v85 int8 = 42
                                            _ = v85
                                        }
                                    case "hello":
                                        v83 += v83
                                    }
                                } else {
                                    v2 = -1
                                    if ((float64((0.0 * 0.1)) <= 0.1) && (v1 && true)) {
                                        v3 -= (v3 + (v3 << ((42) & 63)))
                                        v86 := GenS10{F0: 2}
                                        v87 := GenS13{F0: (-(fi63)), F1: "hello"}
                                        if (42 <= 255) {
                                            v88 := GenS14{F0: 234, F1: 0, F2: 2}
                                            switch "" {
                                            case "a":
                                                v83 += (-(v83))
                                            default:
                                                if v1 {
                                                    if len(v20) > 0 {
                                                        v20[0] = (fi70 % ((456) | 1))
                                                    }
                                                } else {
                                                    if v1 {
                                                        v89 := GenS15{F0: v2, F1: 2}
                                                        v3 -= (v3 / ((1) | 1))
                                                        if (v1 && ("" < "a")) {
                                                            var v90 int64 = (-(v2))
                                                            _ = v90
                                                        }
                                                        fmt.Println(fi63, v1, fi70)
                                                        v83 = v83
                                                        if v1 {
                                                            v91 := GenS6{F0: 32767, F1: 4294967295}
                                                            v92 := GenS1{F0: (v2 - v2)}
                                                            fmt.Println(v3, v1)
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        _ = v89
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v93 float64 = 3.14159
                                                            _ = v93
                                                        case "\n\t":
                                                            fmt.Println(v2, fi70)
                                                        case "":
                                                            v3 = (v3 + v3)
                                                        default:
                                                            fmt.Println(fi70, v2)
                                                        }
                                                        v94 := GenS12{F0: 531, F1: 133, F2: 110}
                                                        _ = v94
                                                    }
                                                    var v95 int16 = (v3 % (((v3 ^ (^(v3)))) | 1))
                                                    _ = v95
                                                }
                                            }
                                            _ = v88
                                        }
                                        _ = v86
                                        _ = v87
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        var v96 int16 = (v3 | (v3 >> ((1) & 63)))
                                        _ = v96
                                    default:
                                        v1 = (2 != 21)
                                    }
                                    if v1 {
                                        _ = len(v20)
                                        var v97 bool = (!(!v1))
                                        _ = v97
                                    }
                                }
                                fmt.Println(v2)
                                v1 = (1 == 733)
                                _ = v83
                            }
                        case 4294967295:
                            v1 = (fi63 == 128)
                        case 127:
                            v98 := GenS11{F0: 499}
                            _ = v98
                        }
                        var v99 uint64 = 486
                        _ = v99
                    }
                    var v100 int16 = (v3 / (((v3 * (v3 ^ (-(v3))))) | 1))
                    _ = v100
                }
            default:
                var v101 int = (-(fi63))
                _ = v101
            }
            _ = fi70
        }
        _ = fi63
    }
    if v1 {
        var v102 float64 = 1.5
        _ = v102
    } else {
        v20 = append(v20, 32768)
    }
    v20 = append(v20, -9223372036854775808)
    v103 := GenS15{F0: (-(v2)), F1: 42}
    v104 := GenS16{F0: 4294967295, F1: (v3 | v3), F2: 22}
    if (42 > 2147483647) {
        var v105 uint64 = 324
        var v106 float32 = (1e10 - float32(float64(float32(0.1))))
        _ = v105
        _ = v106
    }
    v107 := GenS17{F0: (v2 ^ (v2 - v2)), F1: 42}
    var v108 uint64 = 42
    v3 = (v3 % ((0) | 1))
    fmt.Println(v2, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v103
    _ = v104
    _ = v107
    _ = v108
    _ = v2
    _ = v20
    _ = v3
    _ = v62
}
