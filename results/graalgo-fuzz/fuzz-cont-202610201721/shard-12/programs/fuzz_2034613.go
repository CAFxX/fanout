package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int; F2 string }
type GenS4 struct { F0 int; F1 string }
type GenS5 struct { F0 uint; F1 int32; F2 int32 }
type GenS6 struct { F0 int64; F1 uint16; F2 bool }
type GenS7 struct { F0 int32; F1 string; F2 uint }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 string; F2 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint; F1 int8 }
type GenS14 struct { F0 int8 }
func genFn0(p0 bool, p1 bool, p2 string) uint32 {
    p1 = (1000 > 1)
    for fi5 := 0; fi5 < 13; fi5++ {
        fmt.Println(fi5, fi5)
        switch p2 {
        case "a":
            v6 := []int{(fi5 * fi5), fi5, fi5, (fi5 & 9223372036854775807)}
            _ = v6
        case "":
            v7 := GenS1{F0: false, F1: fi5, F2: 969}
            _ = v7
        }
        p1 = p1
        _ = fi5
    }
    p1 = ((-531 < 0) || (p0 || (0.1 < 3.14159)))
    _ = p0
    _ = p1
    _ = p2
    return 366
}
func genFn1() int32 {
    v8 := []int{32767, -2147483648, -9223372036854775808, 0}
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(fi9)
        var v10 string = "a"
        _ = fi9
        _ = v10
    }
    var v11 float32 = 0.0
    if false {
        switch "hello" {
        case "\n\t":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "hello":
            v11 -= (v11 + v11)
        case "a":
            if false {
                if false {
                    switch 0 {
                    case -9223372036854775808:
                        var v12 int64 = 896
                        _ = v12
                    default:
                        v11 = 0.1
                    }
                    if (32767 <= 9223372036854775807) {
                        v13 := GenS2{F0: 0}
                        var v14 uint = 468
                        _ = v13
                        _ = v14
                    }
                    if (!true) {
                        var v15 float32 = 3.14159
                        if (!false) {
                            v11 = v15
                            if (!(1 <= 1)) {
                                var v16 uint8 = 2
                                _ = v16
                            }
                            var v17 int16 = -1
                            _ = v17
                        } else {
                            fmt.Println(v11)
                        }
                        if len(v8) > 0 {
                            v8[0] = 0
                        }
                        _ = v15
                    } else {
                        if true {
                            if false {
                                var v18 int8 = 42
                                _ = v18
                            } else {
                                if true {
                                    v19 := GenS3{F0: 0, F1: 127, F2: ("\n\t" + "\n\t")}
                                    var v20 int32 = 2147483647
                                    v20 = v20
                                    _ = v19
                                    _ = v20
                                } else {
                                    for fi21 := 0; fi21 < 20; fi21++ {
                                        if false {
                                            v22 := GenS4{F0: 127, F1: (("\n\t" + "foo bar") + (("hello" + "foo bar") + "\n\t"))}
                                            _ = v22
                                        }
                                        v23 := GenS5{F0: 255, F1: 39, F2: 1}
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            v11 = ((3.14159 + v11) * (3.14159 * 1e10))
                                            v8 = append(v8, fi21)
                                            v25 := GenS2{F0: 811}
                                            _ = fi24
                                            _ = v25
                                        }
                                        _ = fi21
                                        _ = v23
                                    }
                                    v11 -= 1.5
                                    v26 := GenS6{F0: 9223372036854775807, F1: 312, F2: (-1 > 42)}
                                    _ = v26
                                }
                                if len(v8) > 0 {
                                    v8[0] = 0
                                }
                                switch 42 {
                                case 128:
                                    for fi27 := 0; fi27 < 3; fi27++ {
                                        for fi28 := 0; fi28 < 2; fi28++ {
                                            var v29 uint8 = 1
                                            v29 = 42
                                            var v30 int64 = 9223372036854775807
                                            _ = fi28
                                            _ = v29
                                            _ = v30
                                        }
                                        if (65 > -1) {
                                            if true {
                                                var v31 uint = 4294967295
                                                _ = v31
                                            }
                                        }
                                        _ = fi27
                                    }
                                default:
                                    v11 -= ((float32(1.5) / float32((3.14159 + ((((0.1 + (((((float64(v11) + 3.14159) - 1.5) / -2.25) - float64(float32((((0.1 / (float64(v11) + (-2.25 * (0.1 + -2.25)))) * 3.14159) - ((-2.25 * (1e10 - 1.5)) - 0.1))))) - 1.5)) / (0.0 + (3.14159 / 1.0))) * 1e10) * float64((float32((-2.25 / (1.5 - float64((v11 / 1.0))))) / v11)))))) - 0.0)
                                }
                            }
                            var v32 int64 = 0
                            _ = v32
                        }
                    }
                }
                if false {
                    v11 = (((1.5 / (3.14159 * v11)) * 1.5) - v11)
                    fmt.Println(v11)
                    fmt.Println(v11)
                }
            }
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            v11 = ((3.14159 - v11) + v11)
            var v34 bool = true
            switch ("a" + ((("a" + "foo bar") + "") + "foo bar")) {
            case "foo bar":
                if ((-2.25 - (-2.25 / (0.0 + 0.1))) <= 3.14159) {
                    v34 = v34
                    v11 = (v11 * v11)
                    if len(v8) > 0 {
                        v8[0] = (^(fi33))
                    }
                }
            case "":
                v35 := GenS7{F0: 76, F1: "a", F2: 825}
                _ = v35
            case "hello":
                if len(v8) > 0 {
                    v8[0] = (fi33 ^ fi33)
                }
            }
            _ = fi33
            _ = v34
        }
    }
    v11 = (float32(1e10) * v11)
    _ = v11
    _ = v8
    return 551
}

func main() {
    v1 := []uint32{65535, 2147483648}
    var v2 uint16 = 2
    var v3 int64 = -810
    v4 := make(map[string]uint32)
    fmt.Println(v2, v3)
    var v36 string = ""
    v37 := GenS1{F0: (0 <= 2), F1: 0, F2: 57}
    var v38 string = "\n\t"
    v39 := GenS8{F0: -9223372036854775808, F1: v2}
    v1 = append(v1, 771)
    if (1 > 41) {
        fmt.Println(v2)
        v3 += (v3 - (v3 << ((991) & 63)))
        for fi40 := 0; fi40 < 1; fi40++ {
            switch fi40 {
            case 9223372036854775807:
                v4[v38] = 0
            default:
                if ((true || true) || false) {
                    var v41 uint16 = 2
                    _ = v41
                }
            }
            _ = fi40
        }
    } else {
        v36 = v36
        if (!false) {
            fmt.Println(v2)
            if (42 >= 0) {
                v4[(v36 + (("foo bar" + ("hello" + "a")) + v36))] = 1
            } else {
                v2 = (v2 * (v2 << ((1) & 63)))
                v3 += v3
                for fi42 := 0; fi42 < 4; fi42++ {
                    switch 1 {
                    case -344:
                        var v43 int64 = -9223372036854775808
                        _ = v43
                    case 127:
                        if (!false) {
                            switch (v38 + v36) {
                            case "":
                                var v44 uint64 = 42
                                _ = v44
                            case "hello":
                                if (1 <= 0) {
                                    var v45 uint64 = 2
                                    v45 = 0
                                    _ = v45
                                } else {
                                    if true {
                                        var v46 uint8 = 127
                                        if (1 > 0) {
                                            for fi47 := 0; fi47 < 9; fi47++ {
                                                var v48 string = v38
                                                _ = fi47
                                                _ = v48
                                            }
                                        } else {
                                            var v49 uint8 = v46
                                            fmt.Println(v46, v2, v38)
                                            if (127 != 77) {
                                                switch (v36 + "hello") {
                                                case "":
                                                    for fi50 := 0; fi50 < 10; fi50++ {
                                                        fmt.Println(fi50)
                                                        if len(v1) > 0 {
                                                            v1[0] = 2147483648
                                                        }
                                                        switch fi50 {
                                                        case -17:
                                                            v38 = (((v38 + (((v36 + "foo bar") + v36) + v36)) + v36) + "")
                                                        case 42:
                                                            if true {
                                                                v4[v38] = 2147483648
                                                            } else {
                                                                fmt.Println(fi50, v38, fi42)
                                                                v38 = "\n\t"
                                                                v1 = append(v1, 1)
                                                            }
                                                        }
                                                        _ = fi50
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v36)
                                                case "hello":
                                                    v2 = 273
                                                default:
                                                    var v51 string = v36
                                                    _ = v51
                                                }
                                                v38 = v36
                                                fmt.Println(v46, v3)
                                            }
                                            _ = v49
                                        }
                                        _ = v46
                                    } else {
                                        v52 := v4[v38]
                                        for fi53 := 0; fi53 < 18; fi53++ {
                                            var v54 int8 = 127
                                            _ = fi53
                                            _ = v54
                                        }
                                        fmt.Println(v2, v38)
                                        _ = v52
                                    }
                                    var v55 float32 = (0.0 - (0.0 * float32(-2.25)))
                                    _ = v55
                                }
                            default:
                                if len(v1) > 0 {
                                    v1[0] = 240
                                }
                            }
                            delete(v4, (v36 + "foo bar"))
                            _ = len(v4)
                        }
                    default:
                        v56 := GenS6{F0: 373, F1: (v2 / (((v2 / ((32767) | 1))) | 1)), F2: false}
                        _ = v56
                    }
                    var v57 uint32 = 0
                    v58 := GenS9{F0: false}
                    _ = fi42
                    _ = v57
                    _ = v58
                }
            }
        } else {
            for fi59 := 0; fi59 < 10; fi59++ {
                if (!(((false || (255 == v2)) && false) || (42 > 0))) {
                    switch v36 {
                    case "a":
                        for fi60 := 0; fi60 < 10; fi60++ {
                            fmt.Println(v2, v36, v38)
                            fmt.Println(fi60, v2)
                            var v61 string = v36
                            _ = fi60
                            _ = v61
                        }
                    case "hello":
                        for fi62 := 0; fi62 < 7; fi62++ {
                            v63 := GenS6{F0: -122, F1: (v2 ^ v2), F2: true}
                            fmt.Println(v38)
                            v3 += (v3 % ((v3) | 1))
                            _ = fi62
                            _ = v63
                        }
                    default:
                        var v64 float32 = float32((1.5 + (0.0 + 1.5)))
                        _ = v64
                    }
                }
                _ = fi59
            }
            v36 = v38
            v3 = v3
        }
        if len(v1) > 0 {
            v1[0] = 730
        }
    }
    if (9 != 0) {
        fmt.Println(v38, v38)
        switch 4294967295 {
        case -969:
            v38 = v38
        case -544:
            for fi65 := 0; fi65 < 10; fi65++ {
                var v66 uint8 = 0
                _ = fi65
                _ = v66
            }
        default:
            v67 := GenS8{F0: 42, F1: v2}
            _ = v67
        }
    } else {
        for fi68 := 0; fi68 < 20; fi68++ {
            var v69 uint16 = v2
            fmt.Println(v36, v38)
            _ = fi68
            _ = v69
        }
    }
    var v70 uint64 = 2
    v71 := GenS6{F0: v3, F1: 914, F2: (-2.25 > 0.0)}
    if ((-1 >= 0) && false) {
        for fi72 := 0; fi72 < 16; fi72++ {
            if false {
                for fi73 := 0; fi73 < 18; fi73++ {
                    if true {
                        v70 = 400
                    } else {
                        var v74 bool = (891 >= uint(v70))
                        var v75 uint32 = 0
                        if false {
                            var v76 float32 = 1.5
                            _ = v76
                        } else {
                            var v77 uint32 = 2
                            _ = v77
                        }
                        var v78 string = v36
                        var v79 uint32 = v75
                        var v80 int16 = 1000
                        if false {
                            fmt.Println(v80, v36)
                        } else {
                            v81 := GenS10{F0: 0, F1: "\n\t", F2: 272}
                            var v82 int32 = -1
                            _ = v81
                            _ = v82
                        }
                        var v83 int16 = v80
                        v84 := GenS11{F0: 9223372036854775807}
                        v3 += (v3 & 1)
                        _ = v74
                        _ = v75
                        _ = v78
                        _ = v79
                        _ = v80
                        _ = v83
                        _ = v84
                    }
                    switch fi72 {
                    case -1:
                        var v85 int8 = -100
                        var v86 uint32 = 1
                        v87 := GenS7{F0: 1, F1: v36, F2: 2}
                        var v88 int8 = (v85 ^ v85)
                        _ = v85
                        _ = v86
                        _ = v87
                        _ = v88
                    default:
                        if (385 >= -32768) {
                            if false {
                                if false {
                                    fmt.Println(v36, v38, v2)
                                    v1 = append(v1, 42)
                                }
                            }
                            v89 := GenS11{F0: (v3 + v3)}
                            var v90 int = (fi73 - fi73)
                            v91 := v4["\n\t"]
                            if true {
                                var v92 string = "a"
                                _ = v92
                            }
                            fmt.Println(v3)
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                        if true {
                            var v93 int8 = -89
                            if (v3 != (v3 + 42)) {
                                v94 := GenS10{F0: (v3 ^ (v3 & v3)), F1: v38, F2: 0}
                                var v95 float64 = 0.0
                                if false {
                                    var v96 float64 = (0.1 - -2.25)
                                    v4["\n\t"] = 564
                                    _ = v96
                                }
                                _ = v94
                                _ = v95
                            }
                            _ = v93
                        }
                        v4[((v38 + ("" + (((v36 + v36) + v38) + ("hello" + v38)))) + v36)] = 1
                    }
                    v97 := v4["a"]
                    _ = fi73
                    _ = v97
                }
            }
            if false {
                v4[("a" + v36)] = 65535
            } else {
                v38 = v36
                v3 = v3
                v3 = (v3 * -9223372036854775808)
            }
            for fi98 := 0; fi98 < 5; fi98++ {
                delete(v4, v38)
                _ = fi98
            }
            _ = fi72
        }
        v3 -= v3
        if true {
            if (1 >= 2147483647) {
                var v99 int16 = 32767
                var v100 bool = false
                v38 = v36
                _ = v100
                _ = v99
            } else {
                for fi101 := 0; fi101 < 6; fi101++ {
                    v38 = "hello"
                    v102 := GenS12{F0: (^(fi101))}
                    _ = fi101
                    _ = v102
                }
                for fi103 := 0; fi103 < 14; fi103++ {
                    switch -239 {
                    case -52:
                        fmt.Println(fi103, fi103, fi103)
                    case 172:
                        var v104 string = ((((v38 + v36) + ("" + v38)) + (v36 + ((v36 + "a") + "\n\t"))) + v38)
                        _ = v104
                    }
                    if (fi103 < (fi103 - fi103)) {
                        var v105 bool = false
                        _ = v105
                    } else {
                        if ((2147483647 < uint32(v2)) || false) {
                            if true {
                                v106 := GenS13{F0: 601, F1: -12}
                                _ = v106
                            }
                            var v107 uint16 = 0
                            var v108 uint32 = 1
                            _ = v107
                            _ = v108
                        } else {
                            v109 := GenS9{F0: (true || true)}
                            for fi110 := 0; fi110 < 19; fi110++ {
                                v111 := GenS10{F0: (-(v3)), F1: ("" + "\n\t"), F2: 65535}
                                _ = fi110
                                _ = v111
                            }
                            _ = v109
                        }
                    }
                    _ = fi103
                }
            }
        } else {
            delete(v4, (v38 + (v38 + "")))
            for fi112 := 0; fi112 < 17; fi112++ {
                v113 := GenS14{F0: 24}
                _ = fi112
                _ = v113
            }
            v114 := GenS8{F0: -9223372036854775808, F1: v2}
            _ = v114
        }
    }
    var v115 int16 = 0
    var v116 string = v36
    for fi117 := 0; fi117 < 7; fi117++ {
        _ = len(v1)
        _ = fi117
    }
    switch ((((v36 + "") + ("foo bar" + v38)) + (("a" + (("foo bar" + (("" + v38) + "a")) + "foo bar")) + (("a" + v116) + v116))) + v38) {
    case "":
        for fi118 := 0; fi118 < 7; fi118++ {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            _ = fi118
        }
    default:
        var v119 int8 = 0
        _ = v119
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v115:", v115)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v70
    _ = v71
}
