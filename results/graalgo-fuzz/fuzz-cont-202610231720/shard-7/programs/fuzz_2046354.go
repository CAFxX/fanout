package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 uint16; F2 uint }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS9 struct { F0 int; F1 uint32; F2 uint8 }
type GenS10 struct { F0 uint32; F1 uint64; F2 bool }
type GenS11 struct { F0 int; F1 uint16; F2 int64 }
func genFn0(p0 bool, p1 int16, p2 int16) uint16 {
    if p0 {
        v7 := GenS1{F0: 100}
        for fi8 := 0; fi8 < 13; fi8++ {
            v9 := GenS2{F0: 120, F1: 2147483647}
            var v10 uint64 = 1
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 int8 = 42
                if (2147483647 < -965) {
                    if p0 {
                        if p0 {
                            if ((^(v10)) <= 2) {
                                var v13 uint8 = 1
                                var v14 uint64 = (v10 * v10)
                                if p0 {
                                    switch "hello" {
                                    case "a":
                                        p0 = true
                                    case "":
                                        var v15 uint32 = 55
                                        _ = v15
                                    default:
                                        switch "\n\t" {
                                        case "foo bar":
                                            v16 := GenS3{F0: -993, F1: ""}
                                            _ = v16
                                        case "a":
                                            v17 := []int{(fi8 % ((fi8) | 1))}
                                            _ = v17
                                        default:
                                            if p0 {
                                                var v18 int = fi11
                                                fmt.Println(v12)
                                                _ = v18
                                            } else {
                                                var v19 uint = 252
                                                _ = v19
                                            }
                                        }
                                    }
                                    if true {
                                        v20 := GenS2{F0: v13, F1: 2}
                                        _ = v20
                                    }
                                    if ((true || (v14 > (v14 % (((^(v14))) | 1)))) && ((v10 & v14) >= (^(v10)))) {
                                        var v21 uint8 = v13
                                        switch fi8 {
                                        case 128:
                                            v21 = (^(v13))
                                        case 32768:
                                            var v22 uint8 = v13
                                            _ = v22
                                        case 1:
                                            fmt.Println(fi8, v10)
                                        default:
                                            fmt.Println(v21, v10, fi11)
                                        }
                                        switch 0 {
                                        case -1:
                                            v23 := []int{(fi11 >> ((0) & 63)), (fi8 % ((-9223372036854775808) | 1))}
                                            _ = v23
                                        case 187:
                                            var v24 uint64 = (v14 << ((4294967295) & 63))
                                            _ = v24
                                        }
                                        _ = v21
                                    } else {
                                        var v25 bool = true
                                        switch "" {
                                        case "hello":
                                            v26 := []int{1}
                                            _ = v26
                                        case "\n\t":
                                            var v27 int64 = 409
                                            _ = v27
                                        }
                                        var v28 bool = v25
                                        p2 -= int16(v12)
                                        _ = v25
                                        _ = v28
                                    }
                                    v29 := GenS3{F0: fi8, F1: ("hello" + ("a" + "foo bar"))}
                                    v30 := []int{fi8, 980, fi11, fi8}
                                    _ = v29
                                    _ = v30
                                }
                                _ = v13
                                _ = v14
                            } else {
                                if (!true) {
                                    v31 := []int{32767}
                                    p1 += p2
                                    v32 := GenS2{F0: 42, F1: 2}
                                    _ = v31
                                    _ = v32
                                }
                                var v33 uint = 255
                                var v34 uint = v33
                                _ = v33
                                _ = v34
                            }
                            if (0.0 <= float64(3.14159)) {
                                v35 := GenS4{F0: 740, F1: 128, F2: 0}
                                fmt.Println(v10, p2, p1)
                                v36 := []int{510, (^(fi11)), fi11, fi8}
                                var v37 float64 = (((1.5 + -2.25) * 0.0) + 1e10)
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                fmt.Println(p0, fi11, fi11)
                                p2 = p1
                                if (-15 >= 0) {
                                    v38 := GenS4{F0: -797, F1: 2, F2: 2147483647}
                                    p2 = 1
                                    _ = v38
                                } else {
                                    p0 = (float32((1.5 / 3.14159)) <= -2.25)
                                    switch (("a" + "hello") + "a") {
                                    case "a":
                                        if false {
                                            var v39 string = "hello"
                                            v12 = 1
                                            switch (^(fi11)) {
                                            case -781:
                                                fmt.Println(v10)
                                            case 2147483647:
                                                if false {
                                                    switch ((v39 + "\n\t") + v39) {
                                                    case "a":
                                                        v39 = v39
                                                    case "hello":
                                                        var v40 float32 = 3.14159
                                                        _ = v40
                                                    default:
                                                        v10 += 133
                                                    }
                                                    p2 += p2
                                                    v41 := GenS5{F0: 2}
                                                    var v42 uint16 = 255
                                                    fmt.Println(p1)
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    if (v12 < (-(v12))) {
                                                        v43 := []int{(fi8 * (^(fi11))), int(v12), fi8, fi11}
                                                        fmt.Println(p0, v10)
                                                        _ = v43
                                                    }
                                                    v44 := []int{(fi11 << ((4294967295) & 63)), (fi11 - fi11), 128, (^(fi8))}
                                                    _ = v44
                                                }
                                            }
                                            _ = v39
                                        }
                                    case "hello":
                                        v45 := []int{(fi8 / (((fi8 & (fi8 ^ fi8))) | 1)), fi11, (fi11 / ((0) | 1))}
                                        _ = v45
                                    }
                                }
                            }
                            v46 := GenS4{F0: 42, F1: 1, F2: -2147483648}
                            _ = v46
                        }
                        var v47 bool = ((0 < 42) && (216 < 128))
                        var v48 int16 = (p1 >> ((597) & 63))
                        _ = v47
                        _ = v48
                    } else {
                        p2 += (p2 & (-(p1)))
                        if true {
                            v49 := GenS1{F0: v12}
                            var v50 int64 = 1
                            _ = v49
                            _ = v50
                        } else {
                            fmt.Println(p2, v10)
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                fmt.Println(p2)
                            default:
                                p1 = (p1 + p1)
                            }
                            p0 = p0
                        }
                        v51 := []int{}
                        _ = v51
                    }
                } else {
                    var v52 int8 = 0
                    v53 := GenS6{F0: (fi11 * 9223372036854775807), F1: 823, F2: 136}
                    fmt.Println(p0, v12, fi8)
                    v10 += v10
                    var v54 int = (fi11 ^ (fi11 / ((9223372036854775807) | 1)))
                    _ = v52
                    _ = v53
                    _ = v54
                }
                v12 -= (v12 >> ((0) & 63))
                v55 := []int{(fi8 + (fi8 - (-(fi11)))), (fi11 + fi11)}
                _ = fi11
                _ = v12
                _ = v55
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    p1 = (p1 & p1)
    if p0 {
        switch 127 {
        case 127:
            v56 := GenS6{F0: 32768, F1: 32767, F2: 1}
            _ = v56
        default:
            p2 = -272
        }
        for fi57 := 0; fi57 < 15; fi57++ {
            v58 := GenS7{F0: (p2 & 1000), F1: 42}
            _ = fi57
            _ = v58
        }
    } else {
        v59 := []int{938, 32768, -1}
        v60 := GenS8{F0: 4294967295, F1: -49, F2: 65535}
        _ = v59
        _ = v60
    }
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1(p0 int, p1 bool) int {
    fmt.Println(p1, p1)
    for fi61 := 0; fi61 < 3; fi61++ {
        p0 -= p0
        var v62 uint16 = 817
        _ = fi61
        _ = v62
    }
    _ = p0
    _ = p1
    return (p0 | (p0 % ((p0) | 1)))
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 int = 128
    var v4 uint = 54
    var v5 string = "foo bar"
    var v6 string = "\n\t"
    if v2 {
        if v2 {
            v63 := GenS6{F0: v3, F1: 640, F2: v4}
            fmt.Println(v5, v3)
            fmt.Println(v2, v5, v3)
            _ = v63
        } else {
            for fi64 := 0; fi64 < 1; fi64++ {
                if (250 < 65535) {
                    for fi65 := 0; fi65 < 20; fi65++ {
                        v66 := []int{}
                        var v67 int16 = int16(v4)
                        v68 := GenS9{F0: (fi64 ^ genFn1((fi65 % (((fi65 ^ int(v4))) | 1)), (65535 <= 391))), F1: 2147483647, F2: 193}
                        var v69 uint16 = 1
                        switch (v6 + v6) {
                        case "\n\t":
                            if len(v66) >= 2 {
                                v66 = v66[1:len(v66)]
                            }
                        }
                        var v70 uint8 = 1
                        _ = fi65
                        _ = v66
                        _ = v67
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    var v71 int64 = 0
                    _ = v71
                }
                for fi72 := 0; fi72 < 14; fi72++ {
                    v4 = (v4 - (v4 << ((v4) & 63)))
                    var v73 uint16 = 255
                    _ = fi72
                    _ = v73
                }
                switch (v5 + (("" + (v6 + v5)) + v6)) {
                case "\n\t":
                    if false {
                        v4 = (v4 - 900)
                    }
                case "foo bar":
                    v74 := GenS10{F0: 759, F1: 0, F2: (9223372036854775807 != -464)}
                    _ = v74
                default:
                    if (float64(0.0) <= ((-2.25 / 1.5) / 1.0)) {
                        switch v5 {
                        case "foo bar":
                            fmt.Println(v3, v6)
                        case "a":
                            v75 := []int{127, fi64, -2147483648, v3}
                            _ = v75
                        }
                        fmt.Println(v4)
                    } else {
                        var v76 bool = true
                        _ = len(v1)
                        _ = v76
                    }
                }
                _ = fi64
            }
        }
        for fi77 := 0; fi77 < 11; fi77++ {
            var v78 bool = v2
            if (0 != 1) {
                fmt.Println(v3)
                for fi79 := 0; fi79 < 17; fi79++ {
                    if false {
                        var v80 uint16 = 65535
                        v4 = v4
                        _ = v80
                    } else {
                        v1[fi79] = (v6 + (((v5 + ((("\n\t" + "") + v6) + v6)) + v6) + ("\n\t" + "foo bar")))
                        switch ("\n\t" + (v6 + "a")) {
                        case "foo bar":
                            var v81 string = (((("hello" + v5) + v5) + (v5 + ("" + (v6 + "")))) + ("foo bar" + "\n\t"))
                            v82 := GenS5{F0: 2}
                            _ = v81
                            _ = v82
                        }
                        var v83 uint64 = 42
                        var v84 int8 = 100
                        _ = v83
                        _ = v84
                    }
                    _ = fi79
                }
            } else {
                if (3.14159 >= -2.25) {
                    if (v78 && (!false)) {
                        v85 := []int{(fi77 | (v3 % ((genFn1(fi77, v2)) | 1))), (-(fi77)), v3, 361}
                        v86 := v1[2147483647]
                        var v87 int = (^(v3))
                        _ = v85
                        _ = v86
                        _ = v87
                    } else {
                        v78 = false
                    }
                    if v78 {
                        if v78 {
                            if (v4 > v4) {
                                v88 := GenS7{F0: 32767, F1: 127}
                                var v89 uint64 = 979
                                var v90 float32 = (float32(0.1) - 0.1)
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                            if (false || ((245 < (v3 ^ (^(fi77)))) || v78)) {
                                v78 = (genFn0(((!(1 >= 1)) && (!v78)), -32768, 42) != 255)
                            } else {
                                var v91 uint32 = 42
                                for fi92 := 0; fi92 < 15; fi92++ {
                                    v91 = (^(v91))
                                    var v93 bool = (v2 || (v3 <= (fi92 ^ fi92)))
                                    _ = fi92
                                    _ = v93
                                }
                                v94 := GenS6{F0: v3, F1: 631, F2: v4}
                                _ = v91
                                _ = v94
                            }
                            fmt.Println(v5, v4, v5)
                        } else {
                            v6 = v6
                            var v95 int8 = 49
                            if (!(0.1 > (-2.25 / float64(((0.0 / 1.0) + 1.5))))) {
                                v6 = "foo bar"
                                switch ("\n\t" + ("a" + v6)) {
                                case "a":
                                    v96 := []int{(-(v3)), -9223372036854775808, fi77, fi77}
                                    _ = v96
                                case "\n\t":
                                    var v97 uint16 = 42
                                    _ = v97
                                case "foo bar":
                                    if (v2 || v78) {
                                        var v98 int16 = 1000
                                        var v99 int32 = 42
                                        v98 = (v98 << (((v4 / ((v4) | 1))) & 63))
                                        _ = v98
                                        _ = v99
                                    }
                                }
                            }
                            _ = v95
                        }
                        var v100 bool = (0.1 == 0.1)
                        v101 := []int{(^(fi77)), v3}
                        _ = v100
                        _ = v101
                    }
                    v102 := GenS7{F0: 1, F1: 255}
                    _ = v102
                } else {
                    for fi103 := 0; fi103 < 1; fi103++ {
                        var v104 float64 = float64(((-2.25 - ((float32(0.0) * 1.5) - ((0.1 * 3.14159) * 1e10))) / 0.1))
                        fmt.Println(fi77)
                        v2 = v78
                        fmt.Println(v4, v4)
                        _ = fi103
                        _ = v104
                    }
                    switch v5 {
                    case "\n\t":
                        v105 := GenS10{F0: 42, F1: 42, F2: v2}
                        _ = v105
                    case "hello":
                        v3 = genFn1((fi77 | genFn1(fi77, v78)), (0 != 1))
                    }
                    var v106 int16 = 481
                    _ = v106
                }
                fmt.Println(v78)
            }
            _ = fi77
            _ = v78
        }
        v107 := GenS6{F0: (v3 << ((998) & 63)), F1: 223, F2: (v4 | v4)}
        _ = v107
    }
    fmt.Println(v3)
    var v108 float64 = 1.5
    for fi109 := 0; fi109 < 19; fi109++ {
        var v110 bool = (v2 && (!((-1 <= 42) && true)))
        var v111 uint8 = 71
        _ = fi109
        _ = v110
        _ = v111
    }
    v112 := GenS3{F0: -2147483648, F1: (v6 + "a")}
    v108 = v108
    for fi113 := 0; fi113 < 14; fi113++ {
        fmt.Println(v108)
        if (-2147483648 < 127) {
            fmt.Println(v4, v3, v108)
        }
        v1[v3] = v6
        _ = fi113
    }
    if v2 {
        var v114 int8 = 127
        _ = v114
    }
    v115 := GenS11{F0: v3, F1: genFn0((!true), -1, -32768), F2: 1}
    var v116 bool = (true || ((v2 && ((v6 + ((v6 + (v6 + (v5 + (v5 + "hello")))) + v5)) < (v5 + "hello"))) && true))
    v6 = "hello"
    fmt.Println(v108)
    fmt.Println(v108)
    v6 = ((v6 + "") + (v5 + "\n\t"))
    switch v6 {
    case "foo bar":
        var v117 int32 = 227
        _ = v117
    default:
        v118 := []int{440, (v3 * (^(v3)))}
        _ = v118
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v108:", v108)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v112
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
