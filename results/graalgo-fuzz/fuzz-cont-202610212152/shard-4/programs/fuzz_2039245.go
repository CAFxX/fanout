package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 string; F2 int32 }
type GenS7 struct { F0 int; F1 bool; F2 int16 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint, p1 int) int {
    v5 := GenS1{F0: 42, F1: p0}
    var v6 float32 = ((0.1 / (3.14159 * 1e10)) - 0.1)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p1
}
func genFn1(p0 int8, p1 uint64, p2 bool) int8 {
    v7 := GenS1{F0: 516, F1: 2}
    v8 := []int{2147483647}
    fmt.Println(p2, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 bool = (true || (42 <= 1))
    v2 := make(map[int]bool)
    var v3 uint64 = 897
    var v4 uint64 = (v3 << ((65535) & 63))
    switch ("hello" + "hello") {
    case "a":
        v3 += (v3 * 18446744073709551615)
    default:
        v1 = false
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(fi10)
            v3 -= (v3 << ((537) & 63))
            v11 := GenS2{F0: -1, F1: (^(v3))}
            _ = fi10
            _ = v11
        }
        v2[genFn0(183, -9223372036854775808)] = v1
        var v12 uint8 = 22
        _ = fi9
        _ = v12
    }
    switch 434 {
    case 32768:
        for fi13 := 0; fi13 < 1; fi13++ {
            if (1 > 118) {
                var v14 bool = (156 != 9)
                _ = v14
            }
            if true {
                var v15 int32 = -2147483648
                v16 := []int{9223372036854775807, fi13, (fi13 ^ genFn0(2, int(v15))), (^(fi13))}
                var v17 int32 = v15
                _ = v15
                _ = v16
                _ = v17
            } else {
                var v18 uint64 = v3
                v3 = 611
                if (-1 == 0) {
                    fmt.Println(v18, fi13, fi13)
                    v19 := []int{genFn0(2, -1), (^(fi13)), 2147483647}
                    _ = v19
                }
                _ = v18
            }
            _ = fi13
        }
    case 4294967295:
        if (false || v1) {
            fmt.Println(v4, v4)
        }
    }
    v1 = false
    var v20 uint16 = 987
    var v21 uint32 = 0
    for fi22 := 0; fi22 < 6; fi22++ {
        if (!false) {
            v1 = v1
            switch (^(fi22)) {
            case 9223372036854775807:
                var v23 string = "hello"
                _ = v23
            case 2147483647:
                v21 += (v21 >> ((uint(v20)) & 63))
            }
        }
        _ = fi22
    }
    if true {
        var v24 uint16 = 2
        var v25 int32 = 773
        var v26 string = ""
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: -32768, F1: v3}
    v28 := GenS3{F0: "\n\t", F1: -1}
    switch -2147483648 {
    case 925:
        if (-24 <= 0) {
            if v1 {
                switch 32767 {
                case 0:
                    if true {
                        for fi29 := 0; fi29 < 3; fi29++ {
                            if (255 < 127) {
                                v30 := []int{}
                                fmt.Println(fi29, v3)
                                _ = v30
                            } else {
                                fmt.Println(v21)
                                var v31 float32 = float32(float64(float32((1.5 + 1e10))))
                                if v1 {
                                    v20 -= (v20 >> ((0) & 63))
                                    for fi32 := 0; fi32 < 7; fi32++ {
                                        if (!(-1 > 1000)) {
                                            if true {
                                                v20 = v20
                                                v33 := v2[32768]
                                                if false {
                                                    v34 := GenS3{F0: ("a" + ("a" + "\n\t")), F1: -602}
                                                    var v35 int16 = 32767
                                                    v4 = (^(v4))
                                                    _ = v34
                                                    _ = v35
                                                }
                                                _ = v33
                                            } else {
                                                v4 += 2
                                            }
                                            v36 := v2[genFn0(57, genFn0(18446744073709551615, (fi29 / (((fi32 - (fi32 % ((genFn0(0, (fi32 & (fi29 % (((^(fi29))) | 1))))) | 1)))) | 1))))]
                                            v31 += ((v31 + v31) + v31)
                                            _ = v36
                                        }
                                        _ = fi32
                                    }
                                } else {
                                    if true {
                                        var v37 int64 = 9223372036854775807
                                        _ = v37
                                    }
                                    fmt.Println(v4, v4, v21)
                                    var v38 bool = (!((v20 <= v20) && (127 > -1)))
                                    _ = v38
                                }
                                _ = v31
                            }
                            v1 = ((v3 ^ (v3 >> ((1) & 63))) > 18446744073709551615)
                            v20 += 32767
                            _ = fi29
                        }
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 uint = 0
                            v21 = (v21 | (v21 - 1))
                            for fi41 := 0; fi41 < 6; fi41++ {
                                v42 := []int{9223372036854775807, fi39, (fi41 ^ genFn0((v40 & v40), (fi39 >> (((^(v40))) & 63)))), 0}
                                switch fi41 {
                                case -675:
                                    v42 = append(v42, 127)
                                case 4294967295:
                                    v3 = (v4 | 2)
                                case 2147483647:
                                    fmt.Println(v40, v3)
                                }
                                if (427 >= 1) {
                                    fmt.Println(fi41)
                                }
                                _ = fi41
                                _ = v42
                            }
                            _ = fi39
                            _ = v40
                        }
                    } else {
                        v43 := GenS3{F0: "foo bar", F1: 1}
                        _ = v43
                    }
                case 4294967295:
                    for fi44 := 0; fi44 < 13; fi44++ {
                        var v45 bool = ("foo bar" <= "hello")
                        _ = fi44
                        _ = v45
                    }
                default:
                    if (v1 && (v1 || (!true))) {
                        v46 := []int{-676}
                        v47 := GenS2{F0: -32768, F1: (v3 ^ 1)}
                        _ = v46
                        _ = v47
                    } else {
                        v20 -= (v20 | v20)
                        fmt.Println(v4, v21, v21)
                    }
                }
            } else {
                if v1 {
                    var v48 uint = 255
                    _ = v48
                } else {
                    var v49 int8 = 42
                    for fi50 := 0; fi50 < 12; fi50++ {
                        var v51 uint8 = 42
                        _ = fi50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 11; fi52++ {
                        if (!false) {
                            fmt.Println(fi52)
                            if v1 {
                                var v53 float32 = (3.14159 * float32(3.14159))
                                v21 = 4294967295
                                v54 := []int{(-(fi52)), fi52, -292, (-(fi52))}
                                _ = v53
                                _ = v54
                            } else {
                                if false {
                                    switch genFn0(992, -1) {
                                    case 32767:
                                        fmt.Println(v21)
                                    case -18:
                                        if (v21 == (v21 | v21)) {
                                            for fi55 := 0; fi55 < 9; fi55++ {
                                                v56 := GenS4{F0: v1}
                                                _ = fi55
                                                _ = v56
                                            }
                                            var v57 int8 = (^(v49))
                                            _ = v57
                                        }
                                    case 9223372036854775807:
                                        for fi58 := 0; fi58 < 7; fi58++ {
                                            switch (fi52 ^ 1) {
                                            case -716:
                                                v59 := v2[394]
                                                _ = v59
                                            case -1:
                                                v60 := GenS5{F0: "foo bar", F1: 2, F2: -1}
                                                _ = v60
                                            }
                                            var v61 int16 = -257
                                            _ = fi58
                                            _ = v61
                                        }
                                    default:
                                        v62 := []int{(fi52 * 128), -2147483648}
                                        _ = v62
                                    }
                                } else {
                                    if (true || (1000 <= -120)) {
                                        var v63 string = "a"
                                        v64 := []int{fi52}
                                        _ = v63
                                        _ = v64
                                    }
                                    v65 := GenS2{F0: 999, F1: 4294967295}
                                    v66 := []int{0, fi52, (fi52 - (fi52 % ((32767) | 1))), (fi52 << ((506) & 63))}
                                    _ = v65
                                    _ = v66
                                }
                            }
                        } else {
                            v49 = 100
                        }
                        _ = fi52
                    }
                    _ = v49
                }
            }
        } else {
            fmt.Println(v20, v20, v3)
            if (3.14159 <= 1.5) {
                var v67 bool = v1
                switch ("" + "") {
                case "foo bar":
                    var v68 uint = 65535
                    _ = v68
                case "hello":
                    switch genFn0(65535, -9223372036854775808) {
                    case 302:
                        _ = len(v2)
                    case 32767:
                        fmt.Println(v21, v67)
                    case -270:
                        var v69 uint8 = 2
                        _ = v69
                    default:
                        v70 := GenS5{F0: "\n\t", F1: v20, F2: -2147483648}
                        _ = v70
                    }
                case "\n\t":
                    v71 := GenS6{F0: 1, F1: (("hello" + "\n\t") + ("" + "\n\t")), F2: 42}
                    _ = v71
                }
                _ = v67
            }
        }
    case 42:
        v72 := GenS2{F0: -815, F1: v3}
        _ = v72
    case -2147483648:
        var v73 int = 367
        _ = v73
    default:
        v74 := []int{}
        _ = v74
    }
    fmt.Println(v3, v21, v3)
    v75 := GenS7{F0: -9223372036854775808, F1: (!false), F2: -32768}
    if false {
        var v76 int64 = 708
        switch "\n\t" {
        case "hello":
            switch -9223372036854775808 {
            case 42:
                v77 := GenS2{F0: -32768, F1: v3}
                _ = v77
            case 2147483647:
                var v78 uint = 42
                _ = v78
            }
        case "\n\t":
            switch "hello" {
            case "\n\t":
                if true {
                    if false {
                        fmt.Println(v3)
                        delete(v2, 9223372036854775807)
                    } else {
                        var v79 int32 = -1
                        fmt.Println(v79, v3, v21)
                        var v80 string = ""
                        _ = v79
                        _ = v80
                    }
                    var v81 uint8 = 127
                    switch -380 {
                    case 42:
                        v82 := GenS8{F0: ("hello" + "a"), F1: -371}
                        _ = v82
                    case -506:
                        for fi83 := 0; fi83 < 7; fi83++ {
                            fmt.Println(v81, v3, v3)
                            if (650 != (v20 % ((65535) | 1))) {
                                for fi84 := 0; fi84 < 20; fi84++ {
                                    var v85 uint64 = v3
                                    _ = fi84
                                    _ = v85
                                }
                            } else {
                                fmt.Println(fi83)
                                switch (("foo bar" + "") + "foo bar") {
                                case "a":
                                    v86 := GenS8{F0: "a", F1: -2147483648}
                                    _ = v86
                                default:
                                    v3 += v4
                                }
                            }
                            _ = fi83
                        }
                    case 1:
                        v81 += (^(v81))
                    default:
                        switch "" {
                        case "hello":
                            for fi87 := 0; fi87 < 3; fi87++ {
                                for fi88 := 0; fi88 < 10; fi88++ {
                                    if v1 {
                                        var v89 int = (fi88 + (^(fi87)))
                                        fmt.Println(v4, fi88)
                                        v20 = v20
                                        _ = v89
                                    }
                                    v21 = v21
                                    v90 := []int{127}
                                    v1 = v1
                                    _ = fi88
                                    _ = v90
                                }
                                v91 := GenS8{F0: "", F1: -2147483648}
                                v92 := []int{}
                                _ = fi87
                                _ = v91
                                _ = v92
                            }
                        case "":
                            if v1 {
                                v2[-778] = true
                            }
                        }
                    }
                    _ = v81
                } else {
                    v76 = v76
                    var v93 int8 = 61
                    if v1 {
                        var v94 float64 = -2.25
                        switch ("\n\t" + ("\n\t" + "")) {
                        case "\n\t":
                            var v95 bool = v1
                            _ = v95
                        case "hello":
                            v76 = (-(v76))
                        case "a":
                            var v96 int16 = 1
                            _ = v96
                        default:
                            if true {
                                v97 := GenS8{F0: ("" + ((("foo bar" + (("" + "a") + "")) + "a") + ("" + (("\n\t" + ("" + ("hello" + "\n\t"))) + "hello")))), F1: -561}
                                _ = v97
                            }
                        }
                        if v1 {
                            var v98 int32 = -1
                            _ = v98
                        }
                        _ = v94
                    } else {
                        fmt.Println(v76, v4)
                        if v1 {
                            if (v1 && false) {
                                var v99 uint32 = (v21 * (v21 - v21))
                                switch 1 {
                                case 2147483647:
                                    var v100 uint8 = 198
                                    _ = v100
                                case -162:
                                    _ = len(v2)
                                case 127:
                                    v93 -= (v93 - v93)
                                }
                                _ = v99
                            }
                        } else {
                            if (!true) {
                                if ((!(true && (v20 <= uint16(v3)))) || (v20 < (^(v20)))) {
                                    for fi101 := 0; fi101 < 19; fi101++ {
                                        v20 += (v20 ^ (v20 / ((v20) | 1)))
                                        _ = fi101
                                    }
                                    v102 := []int{}
                                    v76 = v76
                                    _ = v102
                                }
                                for fi103 := 0; fi103 < 16; fi103++ {
                                    _ = len(v2)
                                    var v104 string = "\n\t"
                                    v105 := []int{128, (-(fi103)), (^(fi103)), int(v21)}
                                    _ = fi103
                                    _ = v104
                                    _ = v105
                                }
                                switch "hello" {
                                case "a":
                                    switch 2147483647 {
                                    case 628:
                                        var v106 uint32 = (^(v21))
                                        _ = v106
                                    case -496:
                                        for fi107 := 0; fi107 < 1; fi107++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if ((!((v76 + (v76 << ((674) & 63))) > (^(v76)))) || (float32(1.5) == 1e10)) {
                                                    v108 := []int{}
                                                    _ = v108
                                                } else {
                                                    v21 += 0
                                                }
                                            }
                                            _ = fi107
                                        }
                                    case 1:
                                        var v109 uint64 = (^(v4))
                                        _ = v109
                                    default:
                                        v110 := []int{genFn0(42, -9223372036854775808)}
                                        _ = v110
                                    }
                                case "hello":
                                    v111 := v2[int(v76)]
                                    _ = v111
                                }
                            } else {
                                v20 -= 765
                                v3 = 42
                                v21 -= 485
                            }
                            v112 := GenS9{F0: 42}
                            delete(v2, int(v76))
                            _ = v112
                        }
                        for fi113 := 0; fi113 < 19; fi113++ {
                            switch fi113 {
                            case 127:
                                var v114 float64 = float64(float32(1e10))
                                _ = v114
                            default:
                                var v115 float32 = ((3.14159 + (-2.25 - (0.0 - ((3.14159 + (1e10 + ((1e10 - 0.0) * float32(1e10)))) * (0.0 + float32((0.0 * (((0.1 - (0.0 - 1e10)) / 1.5) + 0.0)))))))) / 1e10)
                                _ = v115
                            }
                            fmt.Println(fi113, v4, v1)
                            _ = fi113
                        }
                    }
                    _ = v93
                }
            }
        case "a":
            if v1 {
                if (!(float32((-2.25 * 0.1)) > 1.5)) {
                    if ((v21 ^ uint32(v4)) >= (^(v21))) {
                        delete(v2, 0)
                        v20 = 42
                    }
                    v20 -= (v20 + 1)
                    v116 := GenS10{F0: v21, F1: -128}
                    _ = v116
                }
                var v117 float32 = (1e10 + (((0.0 / (0.1 / 1.0)) + 1.5) - 1.5))
                _ = v117
            }
        }
        switch 4294967295 {
        case 4294967295:
            v118 := v2[646]
            _ = v118
        }
        _ = v76
    } else {
        if (0 != 4294967295) {
            v4 = (v3 << ((1) & 63))
        }
        var v119 uint32 = v21
        if true {
            fmt.Println(v4)
            var v120 int64 = 0
            _ = v120
        }
        _ = v119
    }
    fmt.Println(v1, v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v75
}
