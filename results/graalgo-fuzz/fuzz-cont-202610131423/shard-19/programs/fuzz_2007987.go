package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 string; F1 bool }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 uint16; F1 uint }
type GenS12 struct { F0 int8; F1 string }
type GenS13 struct { F0 int8; F1 uint32 }
type GenS14 struct { F0 int; F1 int16; F2 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 uint8; F1 int32 }
type GenS17 struct { F0 bool }
func genFn0(p0 int32, p1 int, p2 int8) uint64 {
    v6 := GenS1{F0: ("a" + (("a" + ("\n\t" + ("\n\t" + (("\n\t" + (((("foo bar" + "foo bar") + "hello") + "a") + "hello")) + "foo bar")))) + "\n\t")), F1: p1}
    var v7 uint64 = 116
    v7 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}
func genFn1() int {
    if true {
        var v8 int8 = 0
        _ = v8
    }
    switch 4294967295 {
    case -958:
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 float64 = 1e10
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 42}
                switch ("\n\t" + ("a" + "a")) {
                case "hello":
                    v13 := GenS2{F0: genFn0(1, fi9, 127), F1: 144, F2: 18446744073709551615}
                    v10 -= (float64(0.1) - (0.1 * v10))
                    if true {
                        var v14 uint64 = 42
                        _ = v14
                    }
                    _ = v13
                case "\n\t":
                    fmt.Println(fi11, fi11)
                    v10 += -2.25
                case "":
                    var v15 float64 = float64(3.14159)
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
            v16 := GenS3{F0: 18446744073709551615}
            _ = fi9
            _ = v10
            _ = v16
        }
    }
    return 127
}

func main() {
    var v1 uint = 4294967295
    var v2 uint32 = 42
    v3 := []string{(("a" + "") + ""), ("a" + "hello"), "foo bar"}
    v4 := make(map[string]uint32)
    v5 := make(map[int]uint32)
    v1 = v1
    switch ("foo bar" + ("\n\t" + "hello")) {
    case "hello":
        v4[""] = v2
    case "foo bar":
        if (true || (true && (true || (-2147483648 != 4294967295)))) {
            v17 := GenS2{F0: 390, F1: genFn0(-1, 128, -39), F2: (v1 >> ((v1) & 63))}
            var v18 string = "foo bar"
            _ = v17
            _ = v18
        }
    }
    v19 := GenS4{F0: uint8(v2), F1: 0}
    for fi20 := 0; fi20 < 17; fi20++ {
        v1 -= v1
        var v21 int8 = -35
        _ = fi20
        _ = v21
    }
    v1 = (v1 ^ v1)
    if false {
        v2 = (v2 | v2)
        switch "foo bar" {
        case "foo bar":
            v22 := GenS2{F0: 18446744073709551615, F1: 4294967295, F2: 357}
            _ = v22
        case "":
            var v23 float64 = -2.25
            _ = v23
        case "a":
            if true {
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v1)
                case "hello":
                    var v24 float64 = 3.14159
                    _ = v24
                default:
                    _ = len(v3)
                }
                switch 2147483647 {
                case 395:
                    v25 := GenS5{F0: 1, F1: 9223372036854775807}
                    _ = v25
                default:
                    v26 := GenS3{F0: 42}
                    _ = v26
                }
                v27 := v4[("\n\t" + "foo bar")]
                _ = v27
            }
        }
        v2 = 65535
    } else {
        v2 = v2
        if (!(!(true || true))) {
            if true {
                v3 = append(v3, "\n\t")
            }
        } else {
            var v28 uint64 = 1
            v2 += v2
            for fi29 := 0; fi29 < 1; fi29++ {
                if true {
                    switch "hello" {
                    case "hello":
                        v30 := GenS6{F0: "a", F1: (169 >= 127)}
                        _ = v30
                    case "a":
                        for fi31 := 0; fi31 < 20; fi31++ {
                            v32 := v4["a"]
                            if true {
                                v33 := GenS7{F0: 0, F1: 0}
                                _ = v33
                            } else {
                                v2 = 916
                                v1 = (v1 % (((v1 / (((v1 / ((2) | 1))) | 1))) | 1))
                            }
                            _ = fi31
                            _ = v32
                        }
                    case "foo bar":
                        v34 := GenS3{F0: 65535}
                        _ = v34
                    }
                    v1 = (v1 >> ((v1) & 63))
                }
                switch (fi29 << ((0) & 63)) {
                case 32768:
                    v2 = uint32(v1)
                case -1:
                    if false {
                        for fi35 := 0; fi35 < 12; fi35++ {
                            if ((true || false) || (!((3.14159 * (1.5 - float64(0.0))) != 0.1))) {
                                switch "\n\t" {
                                case "\n\t":
                                    if (!(-687 != 1)) {
                                        var v36 uint16 = 0
                                        var v37 string = "\n\t"
                                        _ = v36
                                        _ = v37
                                    }
                                case "foo bar":
                                    v28 = (^(v28))
                                case "hello":
                                    switch 32768 {
                                    case 0:
                                        v38 := v5[fi35]
                                        _ = v38
                                    case 2147483647:
                                        switch (fi35 | (fi29 * (^(fi35)))) {
                                        case -9223372036854775808:
                                            v5[(fi29 & (^(fi29)))] = (v2 | v2)
                                        case 1:
                                            v39 := GenS8{F0: uint8(v2)}
                                            _ = v39
                                        case -697:
                                            var v40 float32 = 1e10
                                            _ = v40
                                        default:
                                            fmt.Println(v1)
                                        }
                                    }
                                }
                                v41 := GenS8{F0: 1}
                                switch -14 {
                                case 0:
                                    var v42 bool = true
                                    _ = v42
                                default:
                                    var v43 string = ""
                                    _ = v43
                                }
                                var v44 uint = uint(v2)
                                v2 -= (v2 % ((v2) | 1))
                                fmt.Println(v44, v28)
                                v44 -= 710
                                switch ("foo bar" + "hello") {
                                case "\n\t":
                                    var v45 int8 = -1
                                    _ = v45
                                case "":
                                    var v46 int64 = -9223372036854775808
                                    _ = v46
                                case "a":
                                    var v47 uint = uint(v2)
                                    _ = v47
                                }
                                _ = v41
                                _ = v44
                            }
                            _ = fi35
                        }
                        if (!false) {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            switch 127 {
                            case 32768:
                                var v48 uint32 = (v2 & 449)
                                _ = v48
                            case 128:
                                if (-2147483648 >= fi29) {
                                    v1 += (v1 << ((v1) & 63))
                                    var v49 int64 = -9223372036854775808
                                    _ = v49
                                } else {
                                    fmt.Println(v28)
                                }
                            case -1:
                                if (float32(-2.25) >= 1.5) {
                                    v50 := GenS7{F0: 215, F1: -3}
                                    var v51 uint8 = 255
                                    if (-32768 == -32768) {
                                        var v52 int16 = 32767
                                        for fi53 := 0; fi53 < 20; fi53++ {
                                            v1 = 0
                                            fmt.Println(v28, v2, v51)
                                            _ = fi53
                                        }
                                        v54 := GenS5{F0: (v52 | (v52 ^ v52)), F1: 1}
                                        _ = v52
                                        _ = v54
                                    }
                                    _ = v50
                                    _ = v51
                                } else {
                                    var v55 float32 = 1.5
                                    for fi56 := 0; fi56 < 10; fi56++ {
                                        v55 = -2.25
                                        _ = fi56
                                    }
                                    _ = v55
                                }
                            }
                        }
                    } else {
                        v1 += 42
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 = v1
                    }
                }
                v3 = append(v3, "a")
                _ = fi29
            }
            _ = v28
        }
    }
    switch 445 {
    case 9223372036854775807:
        v2 += (^(v2))
    case 127:
        for fi57 := 0; fi57 < 10; fi57++ {
            v2 = v2
            _ = fi57
        }
    }
    fmt.Println(v1)
    for fi58 := 0; fi58 < 15; fi58++ {
        if (!true) {
            v59 := GenS9{F0: (("hello" + "foo bar") + (("hello" + "foo bar") + "")), F1: 0}
            _ = v59
        }
        _ = fi58
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (9223372036854775807 != 42) {
        switch 512 {
        case 127:
            if (false || ((float32(1e10) == (-2.25 * 1.5)) && true)) {
                v1 = 668
                if true {
                    var v60 uint64 = 18446744073709551615
                    fmt.Println(v2)
                    _ = v60
                }
                for fi61 := 0; fi61 < 19; fi61++ {
                    var v62 int32 = 42
                    _ = len(v5)
                    _ = fi61
                    _ = v62
                }
            } else {
                v5[4294967295] = (v2 >> ((v1) & 63))
            }
        case 32768:
            fmt.Println(v1)
        default:
            v1 += 0
        }
        for fi63 := 0; fi63 < 20; fi63++ {
            if false {
                fmt.Println(v2)
                v2 = (v2 ^ (v2 << (((v1 + (^(v1)))) & 63)))
            }
            for fi64 := 0; fi64 < 4; fi64++ {
                var v65 float64 = -2.25
                if true {
                    switch fi64 {
                    case 32768:
                        if false {
                            v66 := v5[(fi63 % (((-(fi64))) | 1))]
                            var v67 int16 = -599
                            var v68 uint64 = 690
                            _ = v66
                            _ = v67
                            _ = v68
                        } else {
                            if true {
                                switch fi64 {
                                case 32767:
                                    v69 := GenS10{F0: "", F1: "a"}
                                    _ = v69
                                case -703:
                                    var v70 uint8 = 23
                                    _ = v70
                                case -1:
                                    v65 += v65
                                    var v71 uint8 = 127
                                    _ = v71
                                }
                                v1 = (v1 + 0)
                                var v72 string = ""
                                if true {
                                    if (-1 != -831) {
                                        v1 = 42
                                        var v73 string = v72
                                        v74 := GenS11{F0: 255, F1: v1}
                                        _ = v73
                                        _ = v74
                                    }
                                    if false {
                                        switch (v72 + (v72 + v72)) {
                                        case "":
                                            var v75 uint8 = 0
                                            _ = v75
                                        default:
                                            if (!false) {
                                                v76 := GenS12{F0: 100, F1: v72}
                                                v1 += uint(v2)
                                                var v77 int = int(v1)
                                                var v78 string = "a"
                                                switch v77 {
                                                case -820:
                                                    fmt.Println(v1, v72, v77)
                                                default:
                                                    if true {
                                                        if (!(0.0 <= ((1e10 * (0.0 / 3.14159)) / 1e10))) {
                                                            var v79 int32 = -1
                                                            _ = v79
                                                        } else {
                                                            var v80 uint32 = 334
                                                            fmt.Println(v1)
                                                            v77 -= 42
                                                            _ = v80
                                                        }
                                                        v81 := GenS5{F0: 0, F1: 464}
                                                        v82 := v4[("" + v78)]
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                                _ = v76
                                                _ = v77
                                                _ = v78
                                            }
                                        }
                                        var v83 int16 = 764
                                        if false {
                                            var v84 int64 = int64(v2)
                                            switch -1 {
                                            case 42:
                                                switch v72 {
                                                case "hello":
                                                    v65 = 1.5
                                                }
                                            }
                                            fmt.Println(fi63)
                                            _ = v84
                                        } else {
                                            var v85 uint64 = 18446744073709551615
                                            _ = v85
                                        }
                                        var v86 bool = ((v83 + (v83 >> (((^(v1))) & 63))) > v83)
                                        _ = v83
                                        _ = v86
                                    } else {
                                        if (!true) {
                                            v87 := GenS13{F0: int8(v2), F1: 0}
                                            var v88 uint16 = 2
                                            v2 -= (v2 ^ uint32(v1))
                                            _ = v87
                                            _ = v88
                                        } else {
                                            if (-2147483648 <= 0) {
                                                v1 = (^(v1))
                                            }
                                            v65 = ((1.5 - -2.25) * 0.0)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = v72
                                        }
                                        v1 -= (v1 + 65535)
                                    }
                                } else {
                                    v89 := GenS14{F0: genFn1(), F1: 0, F2: 731}
                                    v65 -= (1.5 * 0.1)
                                    _ = v89
                                }
                                _ = v72
                            } else {
                                if (!true) {
                                    if (v2 <= (^(v2))) {
                                        if true {
                                            v2 = v2
                                            fmt.Println(v1, v2)
                                            v90 := GenS3{F0: (v1 << (((v1 - v1)) & 63))}
                                            if ((255 != 972) && true) {
                                                v65 -= (((v65 + (v65 * v65)) + 1.5) + ((((v65 - v65) * v65) * 3.14159) * 3.14159))
                                                v91 := GenS15{F0: 255}
                                                _ = v91
                                            } else {
                                                fmt.Println(v65, v2)
                                            }
                                            _ = v90
                                        } else {
                                            v3 = append(v3, "foo bar")
                                        }
                                        switch ("foo bar" + ("" + "")) {
                                        case "hello":
                                            v1 = (v1 + v1)
                                        }
                                        switch ("foo bar" + "a") {
                                        case "\n\t":
                                            v65 = v65
                                            if false {
                                                v65 -= ((v65 / (v65 / (((0.1 * v65) / v65) - 0.0))) + (v65 / v65))
                                                v1 -= 401
                                                var v92 uint = v1
                                                _ = v92
                                            }
                                            var v93 int16 = -426
                                            _ = v93
                                        case "a":
                                            if ((true || (127 > 127)) || true) {
                                                v65 += (((v65 + v65) - (0.1 - float64(-2.25))) / v65)
                                                fmt.Println(fi63, v2)
                                                var v94 int = (fi64 >> ((v1) & 63))
                                                _ = v94
                                            } else {
                                                v1 += v1
                                                var v95 int8 = 1
                                                _ = v95
                                            }
                                        case "":
                                            var v96 int32 = 42
                                            _ = v96
                                        }
                                        var v97 int64 = -896
                                        v98 := GenS1{F0: "a", F1: fi63}
                                        v97 = (v97 | (^(v97)))
                                        var v99 int = 715
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                    fmt.Println(v1, fi64, v2)
                                }
                                v2 += v2
                            }
                            delete(v4, "a")
                            var v100 int32 = -2147483648
                            _ = v100
                        }
                    }
                } else {
                    fmt.Println(v1, fi63)
                    v101 := GenS16{F0: 0, F1: -2147483648}
                    _ = v101
                }
                _ = fi64
                _ = v65
            }
            _ = fi63
        }
    } else {
        v102 := GenS17{F0: (((true && false) || true) || false)}
        _ = v102
    }
    v2 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
