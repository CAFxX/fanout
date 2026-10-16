package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint16; F1 uint64; F2 int }
type GenS4 struct { F0 int8; F1 bool; F2 uint64 }
type GenS5 struct { F0 int16; F1 int8; F2 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint32; F2 string }
type GenS9 struct { F0 bool; F1 int8; F2 int16 }
func genFn0(p0 bool, p1 int8) uint {
    if false {
        p1 -= 1
        v6 := GenS1{F0: 1, F1: 42}
        _ = v6
    }
    if p0 {
        var v7 uint16 = 32767
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := []int{32768, fi8, (fi8 + fi8)}
            _ = fi8
            _ = v9
        }
        if (!p0) {
            switch (("hello" + ("hello" + "hello")) + "\n\t") {
            case "a":
                var v10 int = -916
                _ = v10
            case "\n\t":
                if p0 {
                    var v11 uint64 = 42
                    _ = v11
                } else {
                    if p0 {
                        var v12 int64 = -661
                        _ = v12
                    }
                    v13 := []int{-194, 1, 32768}
                    _ = v13
                }
            }
            var v14 int32 = 0
            _ = v14
        } else {
            var v15 bool = true
            v7 += v7
            if (42 <= 893) {
                switch ("foo bar" + (("hello" + "\n\t") + "")) {
                case "foo bar":
                    v16 := GenS1{F0: 1, F1: (^(v7))}
                    _ = v16
                case "\n\t":
                    var v17 int = 32768
                    _ = v17
                case "hello":
                    if ((-2.25 + (1.5 * 0.1)) <= 1.5) {
                        switch -2147483648 {
                        case 32767:
                            var v18 int16 = 1000
                            _ = v18
                        case 1:
                            v15 = false
                        case 9223372036854775807:
                            if (false || false) {
                                v19 := GenS2{F0: 0, F1: 65535}
                                v20 := []int{1, 127}
                                _ = v19
                                _ = v20
                            } else {
                                var v21 uint16 = 0
                                switch 32767 {
                                case -2147483648:
                                    if (2147483647 > 641) {
                                        var v22 string = (("" + ("hello" + "foo bar")) + ("foo bar" + ("\n\t" + "")))
                                        _ = v22
                                    }
                                case 2147483647:
                                    switch "hello" {
                                    case "a":
                                        switch 42 {
                                        case -2147483648:
                                            p0 = ((42 > int32(p1)) && (255 <= v21))
                                        case 9223372036854775807:
                                            var v23 float64 = ((-2.25 / 1.5) + 3.14159)
                                            _ = v23
                                        }
                                    case "":
                                        p1 = 46
                                    case "hello":
                                        switch ("" + "") {
                                        case "hello":
                                            var v24 bool = (1.5 <= 3.14159)
                                            _ = v24
                                        case "a":
                                            for fi25 := 0; fi25 < 13; fi25++ {
                                                switch 417 {
                                                case -658:
                                                    fmt.Println(v7, v15, fi25)
                                                default:
                                                    for fi26 := 0; fi26 < 4; fi26++ {
                                                        fmt.Println(p1, v21, v21)
                                                        var v27 int32 = 1
                                                        _ = fi26
                                                        _ = v27
                                                    }
                                                }
                                                _ = fi25
                                            }
                                        case "\n\t":
                                            for fi28 := 0; fi28 < 15; fi28++ {
                                                v29 := []int{(^(fi28)), 2147483647, fi28, 42}
                                                var v30 uint16 = 381
                                                _ = fi28
                                                _ = v29
                                                _ = v30
                                            }
                                        default:
                                            switch 2147483647 {
                                            case 2147483647:
                                                for fi31 := 0; fi31 < 19; fi31++ {
                                                    v32 := GenS3{F0: v7, F1: 0, F2: 0}
                                                    if v15 {
                                                        fmt.Println(p1)
                                                    } else {
                                                        if false {
                                                            v33 := GenS4{F0: (-(p1)), F1: true, F2: 18446744073709551615}
                                                            v34 := GenS2{F0: 42, F1: 4294967295}
                                                            _ = v33
                                                            _ = v34
                                                        }
                                                        if ((^(v7)) <= 133) {
                                                            switch ("a" + "foo bar") {
                                                            case "foo bar":
                                                                if (p1 > p1) {
                                                                    var v35 uint8 = 128
                                                                    var v36 int8 = p1
                                                                    v35 += v35
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                            }
                                                        } else {
                                                            if (v15 && (421 >= 0)) {
                                                                var v37 bool = false
                                                                var v38 bool = true
                                                                v39 := GenS2{F0: 9223372036854775807, F1: 65535}
                                                                _ = v37
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                if v15 {
                                                                    v21 += v21
                                                                } else {
                                                                    fmt.Println(fi31, v21)
                                                                }
                                                            }
                                                            for fi40 := 0; fi40 < 12; fi40++ {
                                                                var v41 string = ((("" + "hello") + "foo bar") + "a")
                                                                var v42 uint64 = 18446744073709551615
                                                                _ = fi40
                                                                _ = v41
                                                                _ = v42
                                                            }
                                                        }
                                                    }
                                                    _ = fi31
                                                    _ = v32
                                                }
                                            case 1:
                                                var v43 int = 4294967295
                                                _ = v43
                                            default:
                                                for fi44 := 0; fi44 < 11; fi44++ {
                                                    v45 := []int{fi44, fi44, (fi44 | fi44)}
                                                    v7 = (v21 ^ 2)
                                                    v46 := GenS4{F0: (p1 * p1), F1: (!true), F2: 18446744073709551615}
                                                    _ = fi44
                                                    _ = v45
                                                    _ = v46
                                                }
                                            }
                                        }
                                    }
                                case -456:
                                    var v47 float64 = (3.14159 * 0.0)
                                    _ = v47
                                }
                                _ = v21
                            }
                        }
                        fmt.Println(v15, v15, v7)
                        switch -2147483648 {
                        case 9223372036854775807:
                            var v48 float64 = 1.5
                            _ = v48
                        }
                    } else {
                        for fi49 := 0; fi49 < 20; fi49++ {
                            var v50 uint16 = (^(v7))
                            var v51 float64 = float64(0.1)
                            v52 := GenS5{F0: 1000, F1: (p1 << ((799) & 63)), F2: 42}
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    }
                }
                var v53 uint = 1
                var v54 float64 = ((((float64(3.14159) - 1.5) - 0.0) / 1.5) / 1.5)
                _ = v53
                _ = v54
            }
            _ = v15
        }
        _ = v7
    } else {
        switch -525 {
        case 2147483647:
            v55 := GenS5{F0: -1, F1: (p1 / (((p1 | p1)) | 1)), F2: 840}
            _ = v55
        case -2147483648:
            var v56 uint64 = 1
            _ = v56
        default:
            var v57 bool = p0
            _ = v57
        }
        v58 := []int{}
        v58 = append(v58, -260)
        _ = v58
    }
    switch 242 {
    case 127:
        if p0 {
            p0 = (true && ((!p0) || (450 == 220)))
            if (-852 >= 0) {
                switch 4294967295 {
                case -9223372036854775808:
                    var v59 uint8 = 0
                    _ = v59
                default:
                    switch (("\n\t" + ("a" + ((("foo bar" + "hello") + "\n\t") + ""))) + (("\n\t" + ("" + "a")) + "hello")) {
                    case "":
                        if ((false || (p0 && true)) || ((!(600 == 9223372036854775807)) || p0)) {
                            var v60 uint64 = 42
                            _ = v60
                        } else {
                            fmt.Println(p1, p0)
                            switch 127 {
                            case 33:
                                var v61 float64 = 0.1
                                _ = v61
                            case 280:
                                p0 = (!(p0 || true))
                            default:
                                p1 = (p1 % ((p1) | 1))
                            }
                            var v62 uint32 = 2147483647
                            _ = v62
                        }
                    case "\n\t":
                        switch 4294967295 {
                        case 653:
                            if (42 > -202) {
                                for fi63 := 0; fi63 < 12; fi63++ {
                                    p1 = p1
                                    _ = fi63
                                }
                                if true {
                                    v64 := GenS1{F0: -1, F1: 255}
                                    v65 := []int{32768, 9223372036854775807, 9223372036854775807, 0}
                                    _ = v64
                                    _ = v65
                                } else {
                                    switch 127 {
                                    case 2147483647:
                                        if false {
                                            var v66 uint8 = 225
                                            _ = v66
                                        }
                                    case 384:
                                        p0 = (((p1 >> ((934) & 63)) <= (p1 | 0)) || p0)
                                    }
                                    v67 := []int{42}
                                    _ = v67
                                }
                                var v68 uint32 = 4294967295
                                _ = v68
                            } else {
                                p1 = p1
                                v69 := []int{-528, -2147483648}
                                var v70 bool = p0
                                _ = v69
                                _ = v70
                            }
                        case 43:
                            for fi71 := 0; fi71 < 10; fi71++ {
                                if p0 {
                                    p1 = 0
                                } else {
                                    v72 := GenS1{F0: 640, F1: 2}
                                    p1 += (p1 >> ((42) & 63))
                                    p0 = p0
                                    _ = v72
                                }
                                p1 -= 80
                                _ = fi71
                            }
                        case -302:
                            for fi73 := 0; fi73 < 19; fi73++ {
                                if true {
                                    switch (^(fi73)) {
                                    case 1:
                                        var v74 int64 = -500
                                        _ = v74
                                    case 42:
                                        v75 := []int{fi73}
                                        _ = v75
                                    case 2147483647:
                                        fmt.Println(p0)
                                    default:
                                        var v76 int16 = 1
                                        _ = v76
                                    }
                                    p1 = p1
                                }
                                p0 = p0
                                var v77 bool = (p0 || (0 > -9223372036854775808))
                                _ = fi73
                                _ = v77
                            }
                        default:
                            if (1 > 2147483648) {
                                p0 = ((1 != 42) || (p0 && false))
                                v78 := GenS6{F0: true}
                                _ = v78
                            } else {
                                p0 = p0
                            }
                        }
                    case "a":
                        v79 := GenS5{F0: 32767, F1: p1, F2: 18}
                        _ = v79
                    default:
                        p1 += p1
                    }
                }
                fmt.Println(p0)
            } else {
                for fi80 := 0; fi80 < 5; fi80++ {
                    v81 := GenS3{F0: 2, F1: 18446744073709551615, F2: fi80}
                    _ = fi80
                    _ = v81
                }
                p0 = p0
                v82 := []int{-553, 9223372036854775807, 0, -2147483648}
                _ = v82
            }
        }
    case -2147483648:
        for fi83 := 0; fi83 < 18; fi83++ {
            switch ((("" + ("\n\t" + "a")) + "hello") + "a") {
            case "\n\t":
                var v84 float32 = ((1.5 * -2.25) * 3.14159)
                _ = v84
            case "":
                p1 = p1
            }
            _ = fi83
        }
    default:
        var v85 uint = 148
        _ = v85
    }
    switch "\n\t" {
    case "hello":
        v86 := GenS5{F0: 32767, F1: p1, F2: 584}
        _ = v86
    case "\n\t":
        var v87 int64 = -1
        _ = v87
    case "":
        p0 = (p0 || ((float32(3.14159) * -2.25) != 3.14159))
    default:
        for fi88 := 0; fi88 < 15; fi88++ {
            p1 -= p1
            _ = fi88
        }
    }
    _ = p0
    _ = p1
    return 945
}
func genFn1(p0 uint) int64 {
    for fi89 := 0; fi89 < 17; fi89++ {
        for fi90 := 0; fi90 < 3; fi90++ {
            if true {
                v91 := []int{-2147483648, fi90, (fi90 * (fi89 | fi89))}
                p0 += (^(p0))
                _ = v91
            }
            if (((665 < 1) || true) && (!(genFn0(false, -1) < (p0 & (p0 >> (((p0 | 65535)) & 63)))))) {
                var v92 int64 = 1
                _ = v92
            }
            switch (("a" + "hello") + "") {
            case "\n\t":
                p0 -= (p0 & genFn0(true, 100))
            case "hello":
                if false {
                    p0 = genFn0(((float64(0.0) + -2.25) <= (((1.5 - 0.0) * (-2.25 * 1e10)) + float64(0.1))), 1)
                }
            case "foo bar":
                var v93 uint64 = 42
                _ = v93
            default:
                v94 := GenS7{F0: 42, F1: 255, F2: 2147483647}
                _ = v94
            }
            _ = fi90
        }
        _ = fi89
    }
    var v95 int32 = -1
    for fi96 := 0; fi96 < 3; fi96++ {
        if ((0 >= 1) || false) {
            v95 -= v95
            var v97 uint32 = 956
            _ = v97
        }
        fmt.Println(fi96, v95, v95)
        v95 = v95
        _ = fi96
    }
    v98 := GenS3{F0: 255, F1: 18446744073709551615, F2: -9223372036854775808}
    _ = p0
    _ = v95
    _ = v98
    return -574
}

func main() {
    var v1 uint = 0
    var v2 uint32 = 42
    var v3 uint8 = 9
    var v4 uint = (v1 ^ v1)
    v5 := make(map[string]bool)
    v2 = 858
    if false {
        v1 -= 1
        v2 = v2
        v2 = v2
    }
    v99 := GenS8{F0: 2147483647, F1: (v2 % (((v2 - v2)) | 1)), F2: ""}
    if (4 >= (v3 % ((v3) | 1))) {
        v3 = (v3 & (v3 ^ 2))
        var v100 uint = 0
        if ((^(v4)) >= genFn0(true, -128)) {
            for fi101 := 0; fi101 < 7; fi101++ {
                var v102 int8 = int8(fi101)
                _ = fi101
                _ = v102
            }
            var v103 int16 = 963
            var v104 float32 = 1.5
            _ = v103
            _ = v104
        }
        _ = v100
    }
    v4 = (^(v4))
    var v105 int = 4294967295
    for fi106 := 0; fi106 < 15; fi106++ {
        var v107 bool = ("hello" == ("\n\t" + "hello"))
        _ = fi106
        _ = v107
    }
    v3 += v3
    fmt.Println(v105)
    var v108 float32 = float32((0.1 / 1.0))
    v2 += 4294967295
    v109 := GenS9{F0: (0 >= v2), F1: 68, F2: 23}
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v105
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v99
}
