package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 uint32; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint; F1 bool; F2 int16 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint16; F1 uint32 }
type GenS12 struct { F0 string; F1 int32; F2 uint8 }
func genFn0() int32 {
    v7 := []int{9223372036854775807, 4294967295, 9223372036854775807, 36}
    if false {
        if (-2.25 > 1.5) {
            v8 := GenS1{F0: (!true), F1: 443, F2: 271}
            v9 := GenS2{F0: 2147483648, F1: 9223372036854775807}
            if (("foo bar" > (("a" + "a") + "")) || true) {
                var v10 int32 = -1
                _ = v10
            } else {
                if len(v7) > 0 {
                    v7[0] = 1
                }
                var v11 uint = 18446744073709551615
                _ = v11
            }
            _ = v8
            _ = v9
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := GenS1{F0: true, F1: 65535, F2: 42}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        var v15 int8 = 127
        _ = v15
    } else {
        var v16 int16 = -32768
        fmt.Println(v16)
        _ = v16
    }
    switch 130 {
    case 659:
        var v17 int8 = -128
        _ = v17
    }
    var v18 int8 = 127
    _ = v18
    _ = v7
    return -684
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 1
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 float32 = 3.14159
    var v6 float32 = 0.0
    if true {
        if false {
            v19 := []int{}
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = v19
        } else {
            _ = len(v1)
        }
    }
    v6 = float32(1.5)
    switch "a" {
    case "a":
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 uint64 = 4294967295
            _ = fi20
            _ = v21
        }
    case "foo bar":
        if ("a" != "hello") {
            switch 32768 {
            case -9223372036854775808:
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(v5, v2)
                    _ = fi22
                }
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                v3[fi23] = false
                _ = fi23
            }
            _ = len(v4)
        } else {
            if true {
                if (4294967295 != 18446744073709551615) {
                    for fi24 := 0; fi24 < 17; fi24++ {
                        delete(v4, "")
                        fmt.Println(v5, v2)
                        if (4294967295 >= 1) {
                            var v25 float64 = 1.5
                            v2 -= 63
                            delete(v3, (fi24 ^ (fi24 / (((^(fi24))) | 1))))
                            _ = v25
                        } else {
                            v5 = -2.25
                        }
                        _ = fi24
                    }
                    v26 := GenS1{F0: (genFn0() == -504), F1: 501, F2: -1}
                    _ = v26
                }
            } else {
                var v27 int16 = -1
                for fi28 := 0; fi28 < 20; fi28++ {
                    v27 += v27
                    if true {
                        var v29 uint32 = 2
                        var v30 float32 = (v5 / 1.0)
                        for fi31 := 0; fi31 < 19; fi31++ {
                            v5 -= v5
                            v5 = v5
                            var v32 bool = true
                            _ = fi31
                            _ = v32
                        }
                        _ = v29
                        _ = v30
                    }
                    _ = fi28
                }
                switch 42 {
                case 654:
                    v6 = 3.14159
                default:
                    v2 = (v2 | (v2 + 251))
                }
                _ = v27
            }
            v33 := GenS1{F0: true, F1: 533, F2: 1}
            v5 = -2.25
            _ = v33
        }
    }
    v34 := GenS3{F0: 42, F1: -1}
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 5; fi36++ {
            v37 := GenS4{F0: 4294967295, F1: 4294967295, F2: -32768}
            if true {
                var v38 uint8 = 42
                v39 := []int{(fi36 << ((42) & 63))}
                _ = v38
                _ = v39
            }
            if (false || (false && false)) {
                v4["foo bar"] = "a"
                fmt.Println(v2, v6)
            }
            v2 = (v2 + 128)
            _ = fi36
            _ = v37
        }
        v40 := GenS5{F0: false}
        _ = fi35
        _ = v40
    }
    var v41 string = "a"
    var v42 int = -9
    fmt.Println(v5, v42, v2)
    fmt.Println(v5, v42, v2)
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := GenS1{F0: (2 < v2), F1: 1, F2: 234}
        _ = fi43
        _ = v44
    }
    var v45 uint16 = 574
    var v46 uint16 = uint16(v2)
    if true {
        v47 := GenS6{F0: 2, F1: false, F2: -249}
        for fi48 := 0; fi48 < 4; fi48++ {
            v45 = v45
            _ = fi48
        }
        _ = v47
    }
    if (!(!(127 >= 107))) {
        if (false || true) {
            if false {
                switch v42 {
                case -167:
                    if (788 == 1) {
                        if (!((255 != v2) || true)) {
                            var v49 uint64 = 2
                            _ = v49
                        }
                        if ((2 == (v2 * v2)) && false) {
                            for fi50 := 0; fi50 < 12; fi50++ {
                                v51 := GenS1{F0: (1 <= 0), F1: v45, F2: -256}
                                v52 := GenS3{F0: -710, F1: 705}
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                            _ = len(v1)
                            var v53 int = v42
                            _ = v53
                        } else {
                            for fi54 := 0; fi54 < 10; fi54++ {
                                for fi55 := 0; fi55 < 5; fi55++ {
                                    fmt.Println(fi54)
                                    v56 := GenS7{F0: 9223372036854775807, F1: 1}
                                    _ = fi55
                                    _ = v56
                                }
                                _ = fi54
                            }
                        }
                    } else {
                        v3[(v42 * 32768)] = (!(42 >= -351))
                    }
                case 4294967295:
                    fmt.Println(v5)
                case 9223372036854775807:
                    v57 := v1[(^(v42))]
                    _ = v57
                }
            }
            for fi58 := 0; fi58 < 15; fi58++ {
                for fi59 := 0; fi59 < 19; fi59++ {
                    v41 = "\n\t"
                    _ = fi59
                }
                _ = fi58
            }
        } else {
            switch (v41 + "hello") {
            case "\n\t":
                delete(v4, "foo bar")
            case "foo bar":
                v41 = v41
            case "":
                switch v42 {
                case -1:
                    if true {
                        switch (v41 + (v41 + (v41 + (v41 + v41)))) {
                        case "\n\t":
                            if true {
                                switch (v42 - v42) {
                                case 128:
                                    for fi60 := 0; fi60 < 3; fi60++ {
                                        v3[v42] = (v42 > (fi60 / ((-9223372036854775808) | 1)))
                                        v61 := GenS8{F0: uint16(v2)}
                                        _ = fi60
                                        _ = v61
                                    }
                                default:
                                    var v62 uint32 = 2147483647
                                    _ = v62
                                }
                                for fi63 := 0; fi63 < 7; fi63++ {
                                    fmt.Println(v2)
                                    fmt.Println(fi63, v6)
                                    fmt.Println(v41, v5, v42)
                                    _ = fi63
                                }
                                for fi64 := 0; fi64 < 9; fi64++ {
                                    v65 := []int{(v42 & fi64)}
                                    v1[-429] = 2147483647
                                    _ = fi64
                                    _ = v65
                                }
                            }
                        case "foo bar":
                            if (127 != -1) {
                                v2 -= (v2 - 0)
                            }
                        case "":
                            for fi66 := 0; fi66 < 19; fi66++ {
                                for fi67 := 0; fi67 < 1; fi67++ {
                                    v41 = ("\n\t" + ("foo bar" + v41))
                                    if (v2 < v2) {
                                        v68 := v3[(v42 | (fi67 + v42))]
                                        v46 -= 32767
                                        _ = v68
                                    }
                                    _ = fi67
                                }
                                for fi69 := 0; fi69 < 9; fi69++ {
                                    fmt.Println(v2)
                                    _ = fi69
                                }
                                _ = fi66
                            }
                        default:
                            for fi70 := 0; fi70 < 20; fi70++ {
                                var v71 int = 9223372036854775807
                                fmt.Println(v71)
                                var v72 int16 = 16
                                _ = fi70
                                _ = v71
                                _ = v72
                            }
                        }
                    } else {
                        for fi73 := 0; fi73 < 10; fi73++ {
                            v74 := v4[v41]
                            for fi75 := 0; fi75 < 11; fi75++ {
                                v5 = (float32(1.5) * ((((3.14159 - (v5 - 0.1)) - v5) - ((v6 * -2.25) - float32(1e10))) * (v6 + (v5 + (-2.25 / float32(1.5))))))
                                v76 := GenS3{F0: -94, F1: 42}
                                var v77 string = v74
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                            _ = fi73
                            _ = v74
                        }
                        var v78 float64 = 1.5
                        _ = v78
                    }
                default:
                    for fi79 := 0; fi79 < 7; fi79++ {
                        var v80 int32 = -2147483648
                        v45 -= v46
                        _ = fi79
                        _ = v80
                    }
                }
            }
            v45 = 258
            v6 -= (float32(3.14159) + (0.1 + 3.14159))
        }
        delete(v4, v41)
        fmt.Println(v45, v2, v2)
    } else {
        var v81 uint = 2
        _ = v81
    }
    if false {
        if ((v46 / (((^(v46))) | 1)) < (^(v45))) {
            if true {
                fmt.Println(v2)
                v82 := []int{(v42 / (((v42 & int(v2))) | 1)), 128}
                switch (("a" + (((v41 + "a") + ("a" + v41)) + ((v41 + v41) + v41))) + v41) {
                case "a":
                    for fi83 := 0; fi83 < 14; fi83++ {
                        fmt.Println(v42)
                        var v84 bool = true
                        if v84 {
                            fmt.Println(v84)
                        } else {
                            switch v42 {
                            case 646:
                                v1[(v42 * v42)] = (v42 ^ 128)
                            case -674:
                                var v85 uint = 65535
                                _ = v85
                            default:
                                v86 := GenS9{F0: 989}
                                _ = v86
                            }
                            v82 = append(v82, v42)
                            v5 += v6
                        }
                        _ = fi83
                        _ = v84
                    }
                default:
                    v42 = (v42 + (v42 ^ v42))
                }
                _ = v82
            }
            if true {
                var v87 int8 = 127
                _ = v87
            } else {
                v46 = (^(v45))
                if ((v2 ^ (v2 & (v2 * v2))) <= v2) {
                    var v88 string = ((v41 + v41) + v41)
                    fmt.Println(v41)
                    _ = v88
                }
                var v89 uint = 4294967295
                _ = v89
            }
        } else {
            v90 := []int{(v42 & (v42 + -409)), 128}
            _ = v90
        }
        for fi91 := 0; fi91 < 8; fi91++ {
            if true {
                var v92 int32 = -1
                v93 := GenS10{F0: ((!(4294967295 < 1)) && true)}
                v94 := GenS11{F0: 255, F1: 4294967295}
                _ = v92
                _ = v93
                _ = v94
            } else {
                var v95 float32 = v5
                for fi96 := 0; fi96 < 5; fi96++ {
                    v97 := GenS4{F0: uint(v45), F1: 65535, F2: -1}
                    switch v42 {
                    case -2147483648:
                        var v98 int8 = -128
                        _ = v98
                    default:
                        if (4294967295 < 589) {
                            v99 := []int{(-(fi91))}
                            _ = v99
                        }
                    }
                    _ = fi96
                    _ = v97
                }
                _ = v95
            }
            var v100 int16 = 69
            if true {
                if false {
                    if (genFn0() == 0) {
                        v101 := []int{}
                        var v102 float64 = (1.5 - ((0.0 - ((0.0 / 3.14159) * -2.25)) - float64(v5)))
                        switch v42 {
                        case 1:
                            if (127 >= 0) {
                                v100 -= (v100 - v100)
                            } else {
                                v103 := GenS12{F0: v41, F1: genFn0(), F2: (v2 & v2)}
                                v104 := GenS10{F0: (2147483647 >= 65535)}
                                _ = v103
                                _ = v104
                            }
                        case 32767:
                            for fi105 := 0; fi105 < 10; fi105++ {
                                var v106 float64 = v102
                                _ = fi105
                                _ = v106
                            }
                        case 150:
                            fmt.Println(v46, v42)
                        default:
                            for fi107 := 0; fi107 < 8; fi107++ {
                                var v108 int = v42
                                var v109 string = v41
                                _ = fi107
                                _ = v108
                                _ = v109
                            }
                        }
                        _ = v101
                        _ = v102
                    }
                } else {
                    var v110 int16 = int16(v46)
                    if (true || false) {
                        fmt.Println(v45)
                    }
                    fmt.Println(v110, v5)
                    _ = v110
                }
            } else {
                var v111 int = v42
                _ = v111
            }
            _ = fi91
            _ = v100
        }
        var v112 int32 = 1
        _ = v112
    } else {
        v113 := v4[v41]
        for fi114 := 0; fi114 < 1; fi114++ {
            var v115 int = v42
            v116 := GenS5{F0: false}
            _ = fi114
            _ = v115
            _ = v116
        }
        fmt.Println(v113, v6)
        _ = v113
    }
    v1[v42] = v42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
