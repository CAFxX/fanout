package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint; F1 string; F2 int8 }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 int; F1 uint; F2 uint64 }
type GenS6 struct { F0 string; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint8; F1 int }
type GenS10 struct { F0 uint16; F1 int32 }
type GenS11 struct { F0 int }
func genFn0(p0 uint8, p1 uint16) uint32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        if true {
            p0 = (p0 << ((2) & 63))
        }
        var v7 int16 = -357
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 891, F1: "\n\t"}
    p1 -= (p1 & (p1 - 65535))
    _ = p0
    _ = p1
    _ = v8
    return 2147483647
}
func genFn1(p0 uint, p1 uint32, p2 string) int32 {
    var v9 uint16 = 2
    v10 := GenS2{F0: 248, F1: 356, F2: 192}
    var v11 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 1
}

func main() {
    v1 := []uint32{483, 65535}
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]int)
    if false {
        var v12 uint64 = 18446744073709551615
        if (false || false) {
            var v13 bool = (true || (!(9223372036854775807 >= 0)))
            if ((42 != -46) || (!v13)) {
                v14 := GenS3{F0: 2, F1: "hello", F2: -128}
                var v15 int = -535
                _ = v14
                _ = v15
            } else {
                v1 = append(v1, 0)
                fmt.Println(v12, v12)
            }
            _ = v13
        }
        _ = v12
    } else {
        if true {
            var v16 float32 = ((((3.14159 * (1.5 + 0.1)) - float32(0.0)) - ((1e10 + 1e10) - 1.5)) * ((1e10 / 3.14159) - -2.25))
            var v17 string = ("" + ("" + ("hello" + "")))
            switch -9223372036854775808 {
            case -395:
                fmt.Println(v17, v16)
            default:
                var v18 int = 1
                _ = v18
            }
            _ = v16
            _ = v17
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 float64 = 1.5
            _ = fi19
            _ = v20
        }
        if true {
            if len(v1) > 0 {
                v1[0] = 120
            }
            switch 9223372036854775807 {
            case 42:
                var v21 float32 = (((-2.25 + (0.0 * (1.5 - float32(3.14159)))) - -2.25) / 0.1)
                _ = v21
            case 127:
                var v22 int32 = 0
                _ = v22
            case -276:
                switch "" {
                case "foo bar":
                    switch 127 {
                    case -864:
                        var v23 int16 = -32768
                        _ = v23
                    case 2147483647:
                        var v24 float32 = 3.14159
                        _ = v24
                    case -686:
                        var v25 float64 = 3.14159
                        _ = v25
                    default:
                        var v26 bool = (1 > 57)
                        _ = v26
                    }
                case "hello":
                    var v27 int16 = 32767
                    _ = v27
                default:
                    v28 := GenS4{F0: 1, F1: 1, F2: 42}
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 7; fi29++ {
                    var v30 int16 = -32768
                    _ = fi29
                    _ = v30
                }
            }
            if (float64(0.0) > 1.5) {
                for fi31 := 0; fi31 < 14; fi31++ {
                    for fi32 := 0; fi32 < 3; fi32++ {
                        var v33 uint8 = 1
                        v33 = (v33 >> ((65535) & 63))
                        v5[1] = 462
                        _ = fi32
                        _ = v33
                    }
                    var v34 uint64 = 42
                    _ = fi31
                    _ = v34
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 int16 = -493
                    var v37 uint32 = 2147483648
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            } else {
                if false {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        var v39 uint = 879
                        _ = fi38
                        _ = v39
                    }
                }
                if true {
                    switch 736 {
                    case 2147483647:
                        v40 := GenS5{F0: 2147483647, F1: 546, F2: 4294967295}
                        _ = v40
                    }
                    v41 := GenS6{F0: ("a" + ""), F1: 128}
                    _ = v41
                } else {
                    switch ("foo bar" + "\n\t") {
                    case "foo bar":
                        if (-830 == -2147483648) {
                            v42 := GenS7{F0: 2}
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 10; fi43++ {
                                v44 := v3[42]
                                v44 = ((3.14159 != 1e10) && false)
                                switch "foo bar" {
                                case "":
                                    var v45 uint8 = 168
                                    _ = v45
                                case "foo bar":
                                    v46 := GenS4{F0: 0, F1: -1, F2: 0}
                                    _ = v46
                                case "hello":
                                    fmt.Println(fi43)
                                }
                                _ = fi43
                                _ = v44
                            }
                            _ = len(v4)
                            var v47 string = "foo bar"
                            _ = v47
                        }
                    case "\n\t":
                        for fi48 := 0; fi48 < 14; fi48++ {
                            for fi49 := 0; fi49 < 9; fi49++ {
                                var v50 uint64 = 0
                                v51 := GenS7{F0: genFn0(128, 0)}
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                            _ = fi48
                        }
                    case "":
                        if (genFn1(65535, 1, "hello") == -2147483648) {
                            if (-54 <= -1) {
                                var v52 string = "foo bar"
                                var v53 int16 = 1
                                for fi54 := 0; fi54 < 15; fi54++ {
                                    v4[v52] = genFn0(111, 42)
                                    v52 = v52
                                    fmt.Println(v52, fi54, fi54)
                                    _ = fi54
                                }
                                _ = v52
                                _ = v53
                            }
                            v55 := GenS1{F0: 68, F1: (("\n\t" + ("hello" + "")) + ("\n\t" + ("hello" + "foo bar")))}
                            var v56 uint8 = 1
                            _ = v55
                            _ = v56
                        }
                    default:
                        for fi57 := 0; fi57 < 12; fi57++ {
                            if (!false) {
                                v1 = append(v1, 214)
                                v58 := GenS3{F0: 409, F1: "", F2: -1}
                                _ = v58
                            }
                            v59 := GenS6{F0: "foo bar", F1: 0}
                            _ = fi57
                            _ = v59
                        }
                    }
                    var v60 bool = true
                    v61 := GenS6{F0: "hello", F1: 184}
                    _ = v60
                    _ = v61
                }
            }
        }
    }
    var v62 uint64 = 395
    if len(v1) > 0 {
        v1[0] = 868
    }
    var v63 uint16 = 372
    fmt.Println(v62, v62)
    _ = len(v3)
    for fi64 := 0; fi64 < 15; fi64++ {
        v65 := GenS1{F0: genFn0(1, v63), F1: (("hello" + "foo bar") + "foo bar")}
        _ = fi64
        _ = v65
    }
    v62 += (v62 ^ v62)
    for fi66 := 0; fi66 < 5; fi66++ {
        if true {
            if false {
                var v67 int8 = 42
                _ = v67
            } else {
                var v68 int16 = 1
                v69 := v2[4294967295]
                for fi70 := 0; fi70 < 2; fi70++ {
                    var v71 uint8 = 190
                    _ = fi70
                    _ = v71
                }
                _ = v68
                _ = v69
            }
            var v72 bool = false
            _ = v72
        } else {
            v63 = 2
            for fi73 := 0; fi73 < 5; fi73++ {
                switch ((("a" + ("foo bar" + ("" + ""))) + ("\n\t" + (("a" + "a") + ("foo bar" + "a")))) + "\n\t") {
                case "hello":
                    v62 += (v62 << ((4294967295) & 63))
                case "a":
                    if true {
                        v5[690] = 834
                        delete(v4, ("\n\t" + "hello"))
                        v62 = 18446744073709551615
                    } else {
                        v62 -= v62
                    }
                case "foo bar":
                    fmt.Println(fi66)
                }
                if (!((fi66 > (fi66 % ((fi73) | 1))) && true)) {
                    var v74 uint32 = 594
                    v4[""] = 1
                    if (true || true) {
                        var v75 int16 = 42
                        _ = v75
                    } else {
                        if false {
                            if false {
                                var v76 float32 = -2.25
                                _ = v76
                            }
                            v77 := GenS1{F0: v74, F1: "a"}
                            switch "hello" {
                            case "hello":
                                switch (fi73 % ((2147483647) | 1)) {
                                case 4294967295:
                                    if false {
                                        if (v63 < v63) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if false {
                                                if false {
                                                    switch -828 {
                                                    case 128:
                                                        v63 = v63
                                                    case 42:
                                                        v78 := GenS8{F0: 255}
                                                        _ = v78
                                                    }
                                                }
                                                v79 := GenS5{F0: 2147483647, F1: 65535, F2: v62}
                                                fmt.Println(v63, v62)
                                                switch "" {
                                                case "\n\t":
                                                    v74 += v74
                                                default:
                                                    v62 = 2
                                                }
                                                v4[""] = v74
                                                _ = v79
                                            }
                                        } else {
                                            _ = len(v4)
                                            if true {
                                                var v80 int = 32768
                                                v80 = -470
                                                _ = v80
                                            }
                                            v1 = append(v1, v74)
                                        }
                                        if true {
                                            var v81 uint16 = v63
                                            v63 -= (v63 - (v81 - 0))
                                            var v82 uint = 65535
                                            if false {
                                                fmt.Println(v63, v81)
                                                switch ("hello" + "") {
                                                case "":
                                                    v63 -= 2
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    var v83 int = fi73
                                                    _ = v83
                                                }
                                            } else {
                                                v84 := GenS9{F0: 1, F1: fi73}
                                                var v85 uint32 = v74
                                                var v86 int8 = 127
                                                _ = v84
                                                _ = v85
                                                _ = v86
                                            }
                                            fmt.Println(v62)
                                            v82 = (v82 % ((v82) | 1))
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v87 string = "foo bar"
                                            var v88 uint64 = (v62 >> ((626) & 63))
                                            v62 += 860
                                            v63 = 0
                                            _ = v87
                                            _ = v88
                                        }
                                    } else {
                                        fmt.Println(v63)
                                        if (((1 != (^(fi73))) && true) || (0.0 != ((float64(((((3.14159 * -2.25) * 0.0) - 1.5) + 0.1)) * (-2.25 * (((0.1 / (3.14159 - 0.1)) * (((-2.25 - 3.14159) - -2.25) / 3.14159)) + 1.5))) + (1e10 / (3.14159 * ((1e10 + 0.0) * (1e10 + (1.5 - float64(((0.0 / 0.1) / 1.0)))))))))) {
                                            _ = len(v1)
                                            v62 -= (v62 % ((1) | 1))
                                        }
                                        var v89 string = ""
                                        _ = v89
                                    }
                                default:
                                    fmt.Println(v63, fi73)
                                }
                            case "a":
                                if ("\n\t" > "\n\t") {
                                    fmt.Println(fi66, v74)
                                    v90 := GenS4{F0: 0, F1: 42, F2: -9223372036854775808}
                                    if (921 < 255) {
                                        v63 = 42
                                    }
                                    _ = v90
                                } else {
                                    fmt.Println(v74, fi66, v63)
                                    if (false && true) {
                                        fmt.Println(v74, fi73, v74)
                                        v63 = (v63 >> ((65535) & 63))
                                    }
                                    var v91 float32 = 0.0
                                    _ = v91
                                }
                            default:
                                if (65535 <= 109) {
                                    var v92 int64 = 42
                                    if (fi73 <= fi66) {
                                        var v93 float32 = 1.5
                                        v1 = append(v1, v74)
                                        _ = v93
                                    } else {
                                        if false {
                                            v63 += (^(v63))
                                            switch -9223372036854775808 {
                                            case 32767:
                                                switch ("a" + ("\n\t" + "\n\t")) {
                                                case "hello":
                                                    v94 := v3[genFn0(128, (v63 << ((1) & 63)))]
                                                    _ = v94
                                                case "\n\t":
                                                    v62 = v62
                                                }
                                                v95 := GenS5{F0: fi66, F1: 255, F2: v62}
                                                if (0.1 <= float64(3.14159)) {
                                                    delete(v2, v74)
                                                    v63 += (v63 | v63)
                                                    _ = len(v4)
                                                    v96 := GenS8{F0: 567}
                                                    switch "a" {
                                                    case "\n\t":
                                                        v97 := v2[v74]
                                                        _ = v97
                                                    }
                                                    v62 -= v62
                                                    _ = v96
                                                }
                                                _ = v95
                                            }
                                            if ((1e10 * float64(1.5)) == (1e10 - (3.14159 - 0.1))) {
                                                if true {
                                                    if ((0.0 - (1e10 * 0.0)) >= 0.0) {
                                                        v62 -= v62
                                                        var v98 int = (^(fi73))
                                                        _ = v98
                                                    } else {
                                                        fmt.Println(v74, v74)
                                                        fmt.Println(fi66, fi73)
                                                    }
                                                }
                                                v74 -= 42
                                            }
                                        } else {
                                            var v99 int32 = 843
                                            v4[""] = v74
                                            var v100 string = "a"
                                            v5[genFn0(1, v63)] = fi66
                                            var v101 bool = true
                                            _ = v100
                                            _ = v101
                                            _ = v99
                                        }
                                        if false {
                                            fmt.Println(fi66, v62)
                                            var v102 uint64 = (v62 % (((^(v62))) | 1))
                                            var v103 uint64 = 496
                                            _ = v102
                                            _ = v103
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v104 int16 = 0
                                                _ = v104
                                            case "":
                                                var v105 int64 = v92
                                                _ = v105
                                            case "hello":
                                                var v106 bool = (true || false)
                                                _ = v106
                                            }
                                            var v107 uint32 = (v74 ^ (v74 * (v74 ^ (v74 ^ genFn0(2, v63)))))
                                            var v108 int = 887
                                            _ = v107
                                            _ = v108
                                        }
                                    }
                                    var v109 int = 0
                                    _ = v109
                                    _ = v92
                                } else {
                                    fmt.Println(fi66, v62, fi66)
                                }
                            }
                            _ = v77
                        } else {
                            fmt.Println(v63)
                        }
                        v74 = genFn0(207, 65535)
                        v63 = 32767
                    }
                    _ = v74
                } else {
                    var v110 int32 = -2147483648
                    v111 := GenS3{F0: 2, F1: "foo bar", F2: int8(v110)}
                    switch "hello" {
                    case "hello":
                        var v112 bool = false
                        _ = v112
                    case "foo bar":
                        v113 := v2[1]
                        _ = v113
                    default:
                        if ((v110 >> ((2) & 63)) < (v110 - 42)) {
                            v114 := GenS6{F0: ("" + "a"), F1: 249}
                            _ = v114
                        }
                        v115 := GenS10{F0: v63, F1: (v110 * 0)}
                        fmt.Println(fi73, v63, fi73)
                        _ = v115
                    }
                    _ = v110
                    _ = v111
                }
                v62 += (v62 >> ((824) & 63))
                v116 := GenS11{F0: (fi73 | fi73)}
                _ = fi73
                _ = v116
            }
        }
        fmt.Println(fi66, v63)
        _ = fi66
    }
    var v117 int64 = 827
    var v118 int16 = 283
    v63 = v63
    v119 := GenS11{F0: 0}
    if true {
        var v120 uint16 = (v63 - v63)
        var v121 int = 4294967295
        _ = v120
        _ = v121
    }
    var v122 uint64 = 463
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v122:", v122)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
}
