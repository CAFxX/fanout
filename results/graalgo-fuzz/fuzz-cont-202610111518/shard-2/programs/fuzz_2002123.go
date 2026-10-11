package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 string; F2 uint64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 string; F1 int64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 string; F1 bool; F2 uint32 }
type GenS11 struct { F0 string; F1 uint8; F2 uint }
type GenS12 struct { F0 int; F1 uint8 }
func genFn0(p0 int32, p1 string, p2 uint8) uint32 {
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            switch 128 {
            case 32768:
                var v8 bool = (false && true)
                _ = v8
            default:
                if false {
                    p2 += (p2 % ((p2) | 1))
                    switch (^(fi7)) {
                    case 128:
                        p0 -= (p0 * (p0 >> ((uint(p2)) & 63)))
                    }
                    if true {
                        p1 = "a"
                        var v9 uint = 42
                        _ = v9
                    }
                } else {
                    var v10 uint32 = 0
                    var v11 uint16 = 927
                    v12 := []int{(fi7 | (fi7 / ((-212) | 1))), -57, (fi7 % ((703) | 1))}
                    _ = v10
                    _ = v11
                    _ = v12
                }
            }
            v13 := GenS1{F0: (1.5 <= 1e10), F1: (p2 / (((p2 >> ((2) & 63))) | 1))}
            fmt.Println(p2)
            fmt.Println(fi6)
            var v14 string = p1
            _ = fi7
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    p2 = (p2 - (p2 * (p2 & 0)))
    for fi15 := 0; fi15 < 9; fi15++ {
        v16 := GenS1{F0: (p2 <= p2), F1: p2}
        var v17 int8 = 42
        _ = fi15
        _ = v16
        _ = v17
    }
    switch (("" + (p1 + "hello")) + p1) {
    case "\n\t":
        v18 := GenS2{F0: (p2 & (^(p2)))}
        _ = v18
    }
    v19 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return 2147483647
}
func genFn1() int {
    if false {
        v20 := GenS1{F0: true, F1: 30}
        _ = v20
    }
    var v21 int16 = 205
    v21 = (v21 / ((v21) | 1))
    switch "" {
    case "":
        v22 := GenS2{F0: 29}
        _ = v22
    case "foo bar":
        v23 := []int{489}
        _ = v23
    }
    var v24 float64 = 0.1
    _ = v21
    _ = v24
    return 1
}

func main() {
    var v1 bool = (false || false)
    v2 := make(map[int]int)
    var v3 int8 = -128
    v4 := make(map[uint32]bool)
    var v5 uint16 = 42
    _ = len(v2)
    switch 32768 {
    case 42:
        if ((v5 ^ v5) >= 462) {
            v3 = (v3 | -40)
            if (1 <= int64(v5)) {
                v3 = (v3 % (((-(v3))) | 1))
                switch (("a" + ("\n\t" + "")) + ("a" + "\n\t")) {
                case "\n\t":
                    for fi25 := 0; fi25 < 9; fi25++ {
                        var v26 int = fi25
                        v27 := []int{genFn1(), fi25, (fi25 * genFn1())}
                        for fi28 := 0; fi28 < 11; fi28++ {
                            var v29 uint32 = 121
                            _ = fi28
                            _ = v29
                        }
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                case "foo bar":
                    switch genFn1() {
                    case 882:
                        v30 := []int{32768, -9223372036854775808, 128, genFn1()}
                        _ = v30
                    case 32768:
                        var v31 bool = true
                        _ = v31
                    default:
                        if (0 >= uint8(v5)) {
                            if v1 {
                                var v32 float64 = (3.14159 + 1e10)
                                v33 := GenS1{F0: v1, F1: 47}
                                _ = v32
                                _ = v33
                            }
                            if v1 {
                                var v34 uint8 = 0
                                _ = v34
                            } else {
                                var v35 float64 = 3.14159
                                _ = v35
                            }
                            v36 := GenS2{F0: 127}
                            _ = v36
                        }
                    }
                }
                var v37 int = 127
                _ = v37
            }
            for fi38 := 0; fi38 < 4; fi38++ {
                switch fi38 {
                case -9223372036854775808:
                    var v39 int32 = 2147483647
                    _ = v39
                }
                var v40 uint8 = 128
                _ = fi38
                _ = v40
            }
        } else {
            v41 := GenS3{F0: (-(v3)), F1: "hello", F2: 0}
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := []int{}
                _ = fi42
                _ = v43
            }
            v5 = (v5 - (v5 >> ((0) & 63)))
            _ = v41
        }
    case 2147483647:
        fmt.Println(v3, v3, v5)
    case -2147483648:
        var v44 string = "\n\t"
        _ = v44
    }
    if (983 >= 42) {
        v45 := GenS1{F0: (42 != -93), F1: 199}
        v5 = v5
        _ = v45
    } else {
        var v46 uint8 = 128
        _ = v46
    }
    v47 := v4[4294967295]
    v1 = (0 > (^(v5)))
    for fi48 := 0; fi48 < 4; fi48++ {
        if v1 {
            if (v5 != v5) {
                for fi49 := 0; fi49 < 4; fi49++ {
                    var v50 uint64 = 1
                    v51 := v2[(^(fi49))]
                    v51 -= v51
                    _ = fi49
                    _ = v50
                    _ = v51
                }
                if (0 != 449) {
                    var v52 uint = 2
                    switch (fi48 >> (((v52 + (v52 / (((v52 + 2)) | 1)))) & 63)) {
                    case -209:
                        v53 := GenS4{F0: 613, F1: 42, F2: (fi48 >> ((v52) & 63))}
                        _ = v53
                    case 2147483647:
                        for fi54 := 0; fi54 < 14; fi54++ {
                            v1 = (!(v47 || (!v1)))
                            v55 := GenS1{F0: false, F1: 19}
                            _ = fi54
                            _ = v55
                        }
                    case 1:
                        v56 := GenS4{F0: 65535, F1: 2, F2: fi48}
                        _ = v56
                    default:
                        v57 := []int{}
                        _ = v57
                    }
                    fmt.Println(v5, v3, v5)
                    _ = v52
                }
            } else {
                for fi58 := 0; fi58 < 5; fi58++ {
                    fmt.Println(v5, v3, v3)
                    if (32767 <= -1) {
                        v3 = -117
                        v1 = v1
                        v59 := GenS1{F0: (-2.25 <= ((-2.25 - ((1.5 - 3.14159) / 1.0)) / (1.5 - (1.5 + float64(float32(-2.25)))))), F1: 250}
                        fmt.Println(fi48, v1, v1)
                        fmt.Println(v1, v5, v5)
                        var v60 uint16 = v5
                        v61 := GenS1{F0: (127 == 42), F1: 2}
                        fmt.Println(v60, v47)
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        if (74 == 359) {
                            if true {
                                if v47 {
                                    if v1 {
                                        switch 42 {
                                        case -2147483648:
                                            switch (fi48 << ((253) & 63)) {
                                            case 24:
                                                v62 := []int{-2147483648, fi48, (fi48 >> ((1) & 63))}
                                                _ = v62
                                            case 4294967295:
                                                v47 = false
                                            default:
                                                switch 9223372036854775807 {
                                                case 81:
                                                    v63 := GenS2{F0: 42}
                                                    _ = v63
                                                case -36:
                                                    v64 := GenS5{F0: "hello", F1: ("hello" + "\n\t"), F2: 28}
                                                    _ = v64
                                                case 42:
                                                    v65 := GenS6{F0: 0, F1: 32767}
                                                    _ = v65
                                                default:
                                                    var v66 float32 = 3.14159
                                                    _ = v66
                                                }
                                            }
                                        case -472:
                                            if v1 {
                                                var v67 bool = true
                                                switch -532 {
                                                case 2147483647:
                                                    if true {
                                                        var v68 uint64 = 2
                                                        var v69 int32 = -1
                                                        _ = v68
                                                        _ = v69
                                                    } else {
                                                        v70 := GenS7{F0: 824, F1: 476}
                                                        v71 := GenS4{F0: (v5 ^ v5), F1: 724, F2: -329}
                                                        _ = v70
                                                        _ = v71
                                                    }
                                                case 32767:
                                                    v5 += (v5 & v5)
                                                }
                                                _ = v67
                                            } else {
                                                if (!("foo bar" == "")) {
                                                    v47 = ((0.0 * 3.14159) <= ((0.1 - 0.0) * (1e10 / float64((3.14159 - (0.0 - 1e10))))))
                                                    var v72 uint32 = 2147483648
                                                    _ = v72
                                                } else {
                                                    var v73 float64 = 0.0
                                                    v73 = 1e10
                                                    v3 = (v3 ^ v3)
                                                    _ = v73
                                                }
                                                v74 := GenS4{F0: v5, F1: 0, F2: (fi58 * (fi48 ^ (fi58 << ((255) & 63))))}
                                                var v75 int16 = -758
                                                _ = v74
                                                _ = v75
                                            }
                                        }
                                    } else {
                                        v76 := GenS8{F0: ((("\n\t" + "foo bar") + "\n\t") + ("hello" + (((("foo bar" + "hello") + ("\n\t" + ("hello" + ("a" + "")))) + ("foo bar" + "")) + "a"))), F1: -85}
                                        v77 := GenS2{F0: 42}
                                        v78 := GenS8{F0: "a", F1: 1}
                                        if (genFn0(-1, (("\n\t" + "foo bar") + ("hello" + "\n\t")), 128) >= 1) {
                                            v79 := GenS6{F0: -575, F1: 32767}
                                            _ = len(v2)
                                            if (v47 || false) {
                                                fmt.Println(v3)
                                                fmt.Println(fi48)
                                                v1 = (!v1)
                                            } else {
                                                v4[18] = true
                                                var v80 float32 = 0.1
                                                switch "" {
                                                case "\n\t":
                                                    if v1 {
                                                        v81 := GenS1{F0: (!(2147483647 == 81)), F1: 255}
                                                        fmt.Println(fi58, v1, v1)
                                                        if v47 {
                                                            var v82 uint = 65535
                                                            _ = v82
                                                        } else {
                                                            v83 := GenS9{F0: 42}
                                                            v84 := GenS4{F0: 826, F1: 1, F2: fi48}
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                        _ = v81
                                                    }
                                                case "hello":
                                                    switch ("hello" + "\n\t") {
                                                    case "":
                                                        v5 += (^(v5))
                                                    case "a":
                                                        v3 = v3
                                                    case "foo bar":
                                                        var v85 int8 = (v3 / ((v3) | 1))
                                                        _ = v85
                                                    default:
                                                        v80 = v80
                                                    }
                                                default:
                                                    _ = len(v4)
                                                }
                                                _ = v80
                                            }
                                            var v86 int8 = (v3 - (v3 >> ((801) & 63)))
                                            _ = v79
                                            _ = v86
                                        }
                                        v87 := GenS10{F0: ("foo bar" + "\n\t"), F1: (18446744073709551615 != 2), F2: 42}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                        _ = v87
                                    }
                                    fmt.Println(v5, v5)
                                    switch (fi48 | (fi58 % (((fi48 | fi48)) | 1))) {
                                    case 128:
                                        v88 := GenS6{F0: 1, F1: 32767}
                                        _ = v88
                                    case 832:
                                        var v89 uint = 2
                                        v3 -= v3
                                        v47 = (-942 < -950)
                                        fmt.Println(v89, v5, v5)
                                        v89 = 18446744073709551615
                                        _ = v89
                                    default:
                                        var v90 int64 = 1
                                        _ = v90
                                    }
                                }
                                v3 -= (v3 & 0)
                                v91 := GenS3{F0: int8(v5), F1: "", F2: -1}
                                v92 := GenS3{F0: (v3 - -1), F1: "a", F2: 0}
                                _ = v91
                                _ = v92
                            } else {
                                var v93 float64 = -2.25
                                v94 := []int{fi58, 234, (fi58 % (((fi48 << ((49) & 63))) | 1))}
                                if true {
                                    var v95 bool = (0.1 > (3.14159 + 3.14159))
                                    _ = v95
                                } else {
                                    v96 := GenS8{F0: "\n\t", F1: 0}
                                    switch ("hello" + "") {
                                    case "foo bar":
                                        v97 := GenS1{F0: (42 >= 565), F1: 2}
                                        _ = v97
                                    default:
                                        if (9223372036854775807 > 4294967295) {
                                            var v98 uint8 = 252
                                            delete(v4, 45)
                                            v99 := GenS2{F0: v98}
                                            var v100 uint16 = 58
                                            v47 = (18446744073709551615 >= 2)
                                            _ = v100
                                            _ = v98
                                            _ = v99
                                        } else {
                                            v1 = (1 < -960)
                                            if (!(0 != 42)) {
                                                if false {
                                                    if (!(true || ((v93 >= v93) && (!(!true))))) {
                                                        switch 79 {
                                                        case 42:
                                                            if true {
                                                                _ = len(v94)
                                                                v93 += v93
                                                                _ = len(v94)
                                                                v94 = append(v94, 9223372036854775807)
                                                                switch (fi48 / (((fi48 + fi48)) | 1)) {
                                                                case 32768:
                                                                    fmt.Println(fi58, v93, fi58)
                                                                }
                                                            }
                                                        case -2147483648:
                                                            var v101 int64 = -687
                                                            _ = v101
                                                        case -664:
                                                            v47 = true
                                                        }
                                                        v93 = 1e10
                                                    } else {
                                                        var v102 uint64 = 494
                                                        fmt.Println(v3, v93, v93)
                                                        v102 -= v102
                                                        _ = v102
                                                    }
                                                    fmt.Println(v5, fi48, v93)
                                                }
                                                var v103 uint = 1
                                                _ = v103
                                            } else {
                                                v47 = (0 <= 65)
                                            }
                                            if len(v94) >= 2 {
                                                v94 = v94[1:len(v94)]
                                            }
                                            v93 = float64(0.1)
                                            var v104 uint = 42
                                            _ = v104
                                        }
                                    }
                                    _ = v96
                                }
                                _ = v93
                                _ = v94
                            }
                            var v105 float32 = 0.1
                            v1 = true
                            _ = v105
                        } else {
                            v5 = v5
                            if true {
                                v106 := GenS1{F0: true, F1: 2}
                                var v107 uint8 = 128
                                _ = v106
                                _ = v107
                            }
                        }
                        var v108 bool = (("" + "hello") <= (("\n\t" + "foo bar") + ("\n\t" + "hello")))
                        if v1 {
                            var v109 int = fi58
                            fmt.Println(v109, v109, v108)
                            v47 = (!true)
                            _ = v109
                        }
                        _ = v108
                    }
                    v3 = (v3 ^ (v3 % (((^(v3))) | 1)))
                    _ = fi58
                }
                v1 = v47
            }
        } else {
            v5 = 42
        }
        if (1 > -1) {
            fmt.Println(fi48, fi48)
        } else {
            delete(v2, fi48)
            var v110 uint8 = 0
            _ = v110
        }
        v111 := GenS8{F0: "\n\t", F1: 1}
        _ = fi48
        _ = v111
    }
    if v47 {
        v47 = false
        v112 := GenS11{F0: "", F1: 42, F2: 1}
        v113 := GenS12{F0: genFn1(), F1: 255}
        _ = v112
        _ = v113
    }
    fmt.Println(v1, v47, v3)
    fmt.Println(v5, v1, v3)
    switch 32767 {
    case 127:
        v114 := v2[-9223372036854775808]
        _ = v114
    case 32768:
        v3 = (-(v3))
    }
    for fi115 := 0; fi115 < 16; fi115++ {
        var v116 int16 = 243
        _ = fi115
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
