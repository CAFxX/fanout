package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 string; F1 string; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 int32; F2 int64 }
type GenS9 struct { F0 uint32 }
func genFn0(p0 uint32) string {
    if (9223372036854775807 >= 366) {
        v6 := GenS1{F0: 9223372036854775807}
        _ = v6
    } else {
        var v7 uint64 = 2
        v8 := []int{}
        v9 := GenS1{F0: 0}
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 18446744073709551615}
    var v11 uint32 = (p0 >> ((202) & 63))
    _ = p0
    _ = v10
    _ = v11
    return ""
}
func genFn1(p0 uint64, p1 uint16) int8 {
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 uint64 = 2
        if false {
            var v14 int8 = 42
            p0 = (p0 - 681)
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if (true || true) {
        switch -2147483648 {
        case 127:
            var v15 int = 129
            _ = v15
        case -1:
            p0 = 0
        }
    }
    switch "\n\t" {
    case "a":
        p1 -= p1
    case "\n\t":
        var v16 string = "foo bar"
        _ = v16
    case "foo bar":
        v17 := []int{-1, 32767}
        _ = v17
    default:
        p0 = (p0 | (^(p0)))
    }
    fmt.Println(p1, p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 127
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 int = 58
    var v4 uint64 = 18446744073709551615
    v5 := make(map[string]string)
    v4 = v4
    if true {
        switch v3 {
        case 9223372036854775807:
            var v18 uint = 622
            _ = v18
        case -384:
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint = 2
                v1 -= (v1 - v1)
                _ = fi19
                _ = v20
            }
        case 32768:
            v4 -= (v4 ^ (v4 - v4))
        }
        var v21 string = ""
        for fi22 := 0; fi22 < 2; fi22++ {
            for fi23 := 0; fi23 < 7; fi23++ {
                if true {
                    v24 := []int{(v3 + (fi23 + (-(fi23)))), v3, (v3 + fi22)}
                    fmt.Println(fi22)
                    var v25 uint = 938
                    _ = v24
                    _ = v25
                } else {
                    var v26 int8 = genFn1((v4 & (v4 | (v4 << ((585) & 63)))), 42)
                    _ = v26
                }
                _ = fi23
            }
            var v27 int32 = 2147483647
            _ = fi22
            _ = v27
        }
        _ = v21
    }
    if false {
        fmt.Println(v4, v4, v1)
    }
    if true {
        var v28 uint16 = 431
        v29 := GenS1{F0: -1}
        if ((!(-1 >= 32767)) && true) {
            if true {
                if false {
                    var v30 int32 = -1
                    _ = v30
                } else {
                    v4 += (^(v4))
                }
                v31 := GenS3{F0: -1, F1: 2}
                _ = v31
            }
            v3 += (v3 ^ 9223372036854775807)
        } else {
            v32 := GenS2{F0: 4294967295}
            fmt.Println(v3)
            for fi33 := 0; fi33 < 6; fi33++ {
                for fi34 := 0; fi34 < 2; fi34++ {
                    v35 := []int{128}
                    v36 := GenS3{F0: -1, F1: (v4 + 2)}
                    var v37 uint8 = 42
                    _ = fi34
                    _ = v35
                    _ = v36
                    _ = v37
                }
                if ((42 <= 0) && (false && (!(226 >= uint32(v28))))) {
                    v38 := v2[("a" + ("" + "\n\t"))]
                    fmt.Println(v1)
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 11; fi39++ {
                        _ = len(v5)
                        v40 := GenS4{F0: "\n\t", F1: "foo bar", F2: (v4 & (^(v4)))}
                        _ = fi39
                        _ = v40
                    }
                    v41 := GenS4{F0: ("" + ((("foo bar" + "hello") + ("\n\t" + "foo bar")) + "a")), F1: ("hello" + "a"), F2: 42}
                    _ = v41
                }
                _ = fi33
            }
            _ = v32
        }
        _ = v28
        _ = v29
    }
    for fi42 := 0; fi42 < 19; fi42++ {
        var v43 uint64 = 42
        v44 := GenS3{F0: 1000, F1: v4}
        v4 = 866
        _ = fi42
        _ = v43
        _ = v44
    }
    v45 := GenS3{F0: 1, F1: 4294967295}
    switch "a" {
    case "hello":
        v46 := []int{}
        _ = v46
    case "a":
        var v47 uint64 = (^(v4))
        _ = v47
    case "foo bar":
        fmt.Println(v4, v4, v4)
    default:
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 uint = 850
            _ = fi48
            _ = v49
        }
    }
    if true {
        if (!(4294967295 > 65535)) {
            if (true && false) {
                if (true && true) {
                    if (!(("foo bar" != "\n\t") && (genFn1(v4, 65535) > genFn1(0, 1)))) {
                        if true {
                            if false {
                                switch 2147483647 {
                                case -828:
                                    v1 = (v1 % ((-10) | 1))
                                default:
                                    for fi50 := 0; fi50 < 3; fi50++ {
                                        v51 := []int{v3, v3, fi50, (v3 << ((255) & 63))}
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                                delete(v5, "a")
                                var v52 uint = 1
                                _ = v52
                            } else {
                                switch v3 {
                                case 42:
                                    switch ("foo bar" + "") {
                                    case "a":
                                        var v53 int64 = 1
                                        _ = v53
                                    default:
                                        v54 := GenS1{F0: 995}
                                        _ = v54
                                    }
                                case 32767:
                                    v1 = v1
                                default:
                                    for fi55 := 0; fi55 < 7; fi55++ {
                                        if false {
                                            v1 -= v1
                                            var v56 float64 = ((0.1 / ((0.0 + 1e10) / 1.0)) + 0.1)
                                            _ = v56
                                        } else {
                                            var v57 bool = false
                                            v2["\n\t"] = (fi55 << ((4294967295) & 63))
                                            delete(v5, "foo bar")
                                            _ = v57
                                        }
                                        _ = fi55
                                    }
                                }
                            }
                            var v58 int16 = 42
                            _ = v58
                        } else {
                            var v59 int64 = 1
                            _ = v59
                        }
                        var v60 int = 0
                        if true {
                            switch "foo bar" {
                            case "\n\t":
                                var v61 uint16 = 255
                                _ = v61
                            case "foo bar":
                                v3 = 9223372036854775807
                            default:
                                v2["a"] = (v60 / ((v60) | 1))
                            }
                        } else {
                            var v62 string = (("foo bar" + "\n\t") + "hello")
                            var v63 uint32 = 883
                            fmt.Println(v63, v60)
                            _ = v62
                            _ = v63
                        }
                        _ = v60
                    }
                    switch (("\n\t" + (("foo bar" + ("\n\t" + "hello")) + "hello")) + "a") {
                    case "hello":
                        fmt.Println(v4, v4)
                    case "a":
                        switch "a" {
                        case "\n\t":
                            if true {
                                var v64 int16 = 32767
                                _ = v64
                            } else {
                                switch ((("\n\t" + "\n\t") + (((("\n\t" + "hello") + ("foo bar" + ("hello" + ("foo bar" + "hello")))) + ("foo bar" + "")) + "foo bar")) + (("foo bar" + ((("foo bar" + ("hello" + "")) + "a") + ("hello" + "foo bar"))) + ("hello" + "foo bar"))) {
                                case "hello":
                                    _ = len(v5)
                                case "a":
                                    var v65 float32 = 1e10
                                    _ = v65
                                }
                            }
                        }
                    }
                    for fi66 := 0; fi66 < 4; fi66++ {
                        if false {
                            for fi67 := 0; fi67 < 8; fi67++ {
                                if (0 > 1) {
                                    if (127 <= 127) {
                                        v68 := GenS2{F0: 1}
                                        _ = v68
                                    }
                                    fmt.Println(fi66, v3, v4)
                                    switch fi67 {
                                    case 1:
                                        v4 += v4
                                    case 741:
                                        if false {
                                            var v69 bool = true
                                            if (v69 && (135 != 0)) {
                                                fmt.Println(fi67)
                                            } else {
                                                fmt.Println(v69, fi66, v3)
                                                v70 := []int{586}
                                                _ = v70
                                            }
                                            _ = v69
                                        }
                                    default:
                                        if true {
                                            v71 := []int{9223372036854775807}
                                            switch 1 {
                                            case 32768:
                                                var v72 uint32 = 4294967295
                                                _ = v72
                                            }
                                            _ = v71
                                        }
                                        v1 += 100
                                        v73 := []int{}
                                        _ = v73
                                    }
                                } else {
                                    fmt.Println(v1, v3)
                                    v74 := []int{(fi67 % ((fi67) | 1)), (^(fi67))}
                                    _ = v74
                                }
                                fmt.Println(fi67, v3)
                                _ = fi67
                            }
                            _ = len(v5)
                            switch "a" {
                            case "\n\t":
                                v75 := GenS5{F0: 720}
                                _ = v75
                            case "a":
                                v76 := GenS6{F0: (^(v4))}
                                _ = v76
                            case "hello":
                                if false {
                                    v77 := GenS2{F0: (v4 - (v4 % ((664) | 1)))}
                                    for fi78 := 0; fi78 < 2; fi78++ {
                                        fmt.Println(v1, fi78, fi78)
                                        _ = fi78
                                    }
                                    _ = v77
                                } else {
                                    for fi79 := 0; fi79 < 1; fi79++ {
                                        v3 += fi79
                                        v80 := GenS6{F0: (v4 ^ (v4 % (((^(v4))) | 1)))}
                                        _ = fi79
                                        _ = v80
                                    }
                                    var v81 uint32 = 726
                                    _ = v81
                                }
                            }
                        }
                        _ = len(v2)
                        var v82 uint32 = 2147483648
                        _ = fi66
                        _ = v82
                    }
                } else {
                    var v83 int8 = (-(v1))
                    _ = v83
                }
                _ = len(v5)
                switch (((("hello" + "") + "foo bar") + "hello") + "a") {
                case "":
                    v4 -= v4
                case "hello":
                    v84 := []int{v3, (v3 % ((1) | 1)), v3}
                    _ = v84
                case "a":
                    v3 -= (^(v3))
                default:
                    if ((!(42 != 51)) && (0 >= 255)) {
                        if false {
                            var v85 int32 = 490
                            _ = v85
                        } else {
                            var v86 uint64 = v4
                            v87 := GenS2{F0: v4}
                            _ = v86
                            _ = v87
                        }
                        if false {
                            v88 := []int{(-(v3)), v3, 32768, 128}
                            _ = v88
                        }
                    } else {
                        fmt.Println(v1, v1)
                        var v89 uint32 = 0
                        _ = v89
                    }
                }
            }
            var v90 uint8 = 127
            _ = v90
        }
        fmt.Println(v4, v3)
    } else {
        if (499 < 4294967295) {
            v91 := v2["\n\t"]
            var v92 uint8 = 95
            v92 += (v92 | 128)
            _ = v91
            _ = v92
        } else {
            v93 := GenS3{F0: 1000, F1: (^(v4))}
            for fi94 := 0; fi94 < 8; fi94++ {
                _ = len(v5)
                _ = fi94
            }
            var v95 uint8 = 1
            _ = v93
            _ = v95
        }
        if false {
            switch ((("a" + (("" + (("a" + "a") + "a")) + "hello")) + ("a" + "")) + "foo bar") {
            case "foo bar":
                var v96 uint16 = 65535
                _ = v96
            case "":
                var v97 uint8 = 247
                _ = v97
            }
            for fi98 := 0; fi98 < 20; fi98++ {
                fmt.Println(fi98)
                for fi99 := 0; fi99 < 1; fi99++ {
                    v100 := []int{(fi98 << ((uint(v4)) & 63)), v3, 128}
                    if (1 == -1) {
                        var v101 int32 = 2147483647
                        if true {
                            v101 -= -2147483648
                            if (255 > 1) {
                                if len(v100) > 0 {
                                    v100[0] = (fi99 % ((v3) | 1))
                                }
                                v102 := GenS7{F0: 627}
                                v100 = append(v100, 127)
                                _ = v102
                            } else {
                                var v103 uint8 = 128
                                v101 -= 42
                                _ = v103
                            }
                            var v104 uint64 = (^(v4))
                            _ = v104
                        }
                        v105 := GenS3{F0: int16(fi98), F1: v4}
                        _ = v101
                        _ = v105
                    } else {
                        v106 := GenS8{F0: ("a" + "hello"), F1: -2147483648, F2: 1}
                        _ = v106
                    }
                    fmt.Println(v3)
                    _ = fi99
                    _ = v100
                }
                _ = fi98
            }
            v4 -= (^(v4))
        }
        v3 += 4294967295
    }
    for fi107 := 0; fi107 < 1; fi107++ {
        var v108 string = "a"
        if true {
            for fi109 := 0; fi109 < 7; fi109++ {
                if (genFn1(4294967295, 1) < -128) {
                    v110 := []int{282, fi109, (^(fi109))}
                    v111 := GenS7{F0: 42}
                    _ = v110
                    _ = v111
                }
                switch "" {
                case "foo bar":
                    if ((-2.25 + 1.5) != (((1.5 / 1.0) + 3.14159) + 1e10)) {
                        var v112 int = 127
                        fmt.Println(v4)
                        _ = v112
                    } else {
                        v113 := v2[""]
                        v113 = 851
                        var v114 float64 = -2.25
                        _ = v113
                        _ = v114
                    }
                case "":
                    v108 = ((v108 + "\n\t") + "\n\t")
                case "\n\t":
                    v3 = (v3 << ((65535) & 63))
                default:
                    var v115 uint = 1
                    _ = v115
                }
                _ = fi109
            }
            var v116 bool = false
            _ = v116
        } else {
            if (false || true) {
                var v117 uint64 = (v4 * v4)
                _ = v117
            } else {
                v3 += fi107
            }
        }
        var v118 uint = 18446744073709551615
        _ = fi107
        _ = v108
        _ = v118
    }
    v119 := GenS9{F0: 42}
    v3 = v3
    v4 = v4
    var v120 int = v3
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
