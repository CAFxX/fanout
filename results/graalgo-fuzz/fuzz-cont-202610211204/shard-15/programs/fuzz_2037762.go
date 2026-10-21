package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string; F1 uint8; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint64; F1 int16; F2 string }
type GenS8 struct { F0 uint; F1 uint8; F2 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 bool; F2 int8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS14 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 uint32 = 2147483647
    v3 += v3
    var v5 int16 = -1
    v6 := []int{-9223372036854775808, 1}
    v7 := GenS1{F0: 936, F1: "", F2: 128}
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        v8 := GenS2{F0: 65535}
        var v9 uint8 = uint8(v3)
        _ = v8
        _ = v9
    }
    if false {
        if true {
            if true {
                if (3.14159 > 0.0) {
                    var v10 int64 = -587
                    for fi11 := 0; fi11 < 8; fi11++ {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            if (128 == 128) {
                                var v13 uint = 588
                                if true {
                                    v4 += (v4 | (^(v4)))
                                }
                                _ = v13
                            }
                            if true {
                                v14 := GenS3{F0: 1, F1: (fi11 >> ((548) & 63)), F2: 224}
                                v5 = 1000
                                _ = v14
                            } else {
                                v15 := GenS3{F0: 40, F1: 128, F2: uint8(v4)}
                                _ = v15
                            }
                            v16 := GenS4{F0: v5}
                            _ = fi12
                            _ = v16
                        }
                        v17 := GenS1{F0: (v3 + v3), F1: "", F2: fi11}
                        _ = fi11
                        _ = v17
                    }
                    _ = v10
                }
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                var v19 uint16 = 42
                v4 = v4
                var v20 uint64 = v3
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            _ = len(v2)
            fmt.Println(v5, v3, v4)
        }
        if (0 >= 42) {
            switch "\n\t" {
            case "\n\t":
                v21 := GenS4{F0: 32767}
                _ = v21
            case "":
                _ = len(v6)
            case "hello":
                for fi22 := 0; fi22 < 17; fi22++ {
                    v3 = 4294967295
                    if (fi22 <= int(v5)) {
                        switch -2147483648 {
                        case 9223372036854775807:
                            v1[-893] = (^(v4))
                        default:
                            var v23 float32 = (-2.25 * 1e10)
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v4 += 1
                            var v25 float64 = (0.1 / 1.0)
                            v5 -= -975
                            _ = fi24
                            _ = v25
                        }
                        fmt.Println(v4)
                    } else {
                        v6 = append(v6, fi22)
                        if true {
                            for fi26 := 0; fi26 < 17; fi26++ {
                                fmt.Println(v5, fi26, v4)
                                v27 := GenS4{F0: v5}
                                _ = fi26
                                _ = v27
                            }
                        }
                        fmt.Println(v4, fi22)
                    }
                    _ = fi22
                }
            default:
                if true {
                    v5 = int16(v4)
                    v3 -= 4294967295
                    v3 = v3
                }
            }
            v4 = v4
        } else {
            fmt.Println(v3)
        }
        for fi28 := 0; fi28 < 15; fi28++ {
            fmt.Println(v5, fi28, v3)
            var v29 int8 = 127
            _ = fi28
            _ = v29
        }
    }
    v30 := GenS5{F0: (("\n\t" + "") + (("a" + "a") + "\n\t")), F1: 2, F2: 32767}
    switch -109 {
    case 633:
        v4 = 4294967295
    }
    if false {
        if false {
            if (!false) {
                for fi31 := 0; fi31 < 8; fi31++ {
                    v32 := GenS6{F0: 65535, F1: -32768}
                    _ = fi31
                    _ = v32
                }
                if true {
                    v33 := GenS7{F0: v3, F1: v5, F2: "a"}
                    if (!(-239 <= 2147483647)) {
                        if true {
                            switch -1 {
                            case 0:
                                v4 -= v4
                            case -413:
                                v4 = v4
                            case 127:
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = -464
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = -304
                                    }
                                    var v34 uint = 4294967295
                                    _ = v34
                                }
                            default:
                                v35 := GenS4{F0: 32767}
                                _ = v35
                            }
                            switch (("hello" + ("hello" + "a")) + (("\n\t" + ("hello" + (("hello" + ("foo bar" + "foo bar")) + ""))) + "\n\t")) {
                            case "foo bar":
                                v36 := GenS6{F0: 648, F1: -823}
                                _ = v36
                            case "":
                                switch -9223372036854775808 {
                                case 4294967295:
                                    _ = len(v6)
                                case 834:
                                    v2[(("hello" + "\n\t") + "")] = "hello"
                                default:
                                    v37 := GenS5{F0: "hello", F1: 42, F2: 2}
                                    _ = v37
                                }
                            case "hello":
                                for fi38 := 0; fi38 < 13; fi38++ {
                                    v6 = append(v6, fi38)
                                    _ = fi38
                                }
                            default:
                                var v39 uint = 4294967295
                                _ = v39
                            }
                        }
                    }
                    _ = v33
                } else {
                    switch 342 {
                    case 42:
                        switch 32768 {
                        case 42:
                            var v40 int64 = 0
                            _ = v40
                        case 567:
                            var v41 int16 = v5
                            _ = v41
                        default:
                            v5 = v5
                        }
                    case -849:
                        for fi42 := 0; fi42 < 16; fi42++ {
                            var v43 float32 = 1e10
                            if (42 <= 1000) {
                                for fi44 := 0; fi44 < 18; fi44++ {
                                    var v45 uint = 1
                                    var v46 int64 = -9223372036854775808
                                    v47 := GenS5{F0: "\n\t", F1: 2, F2: 255}
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                                if (1 == 255) {
                                    for fi48 := 0; fi48 < 19; fi48++ {
                                        switch "foo bar" {
                                        case "a":
                                            var v49 float32 = 3.14159
                                            _ = v49
                                        case "":
                                            v50 := v2["\n\t"]
                                            _ = v50
                                        case "hello":
                                            var v51 int16 = v5
                                            _ = v51
                                        default:
                                            var v52 uint16 = 255
                                            _ = v52
                                        }
                                        _ = fi48
                                    }
                                    var v53 int32 = -2147483648
                                    if len(v6) > 0 {
                                        v6[0] = fi42
                                    }
                                    _ = v53
                                }
                            }
                            var v54 int32 = 2147483647
                            _ = fi42
                            _ = v43
                            _ = v54
                        }
                    default:
                        v5 -= v5
                    }
                    var v55 bool = (9223372036854775807 >= -2147483648)
                    _ = v55
                }
                v3 += v3
            } else {
                v56 := GenS6{F0: 703, F1: (v5 & v5)}
                _ = v56
            }
            v5 = (v5 ^ 0)
            if true {
                var v57 float64 = (1e10 - 3.14159)
                v57 = v57
                v58 := GenS3{F0: 42, F1: 32767, F2: 200}
                _ = v57
                _ = v58
            }
        } else {
            v59 := GenS8{F0: 255, F1: 42, F2: 0}
            v5 += -173
            v60 := GenS9{F0: (1 == 0)}
            _ = v59
            _ = v60
        }
        v61 := GenS1{F0: (v3 | v3), F1: "", F2: -554}
        _ = v61
    } else {
        switch "hello" {
        case "\n\t":
            if (false || true) {
                for fi62 := 0; fi62 < 8; fi62++ {
                    fmt.Println(fi62, v3, v5)
                    _ = fi62
                }
                switch 128 {
                case 128:
                    v4 -= v4
                case 755:
                    switch 521 {
                    case -687:
                        delete(v2, "\n\t")
                    case 42:
                        for fi63 := 0; fi63 < 8; fi63++ {
                            switch "foo bar" {
                            case "a":
                                v5 -= (-(v5))
                            default:
                                var v64 int16 = (v5 << ((474) & 63))
                                _ = v64
                            }
                            if (42 <= -1) {
                                var v65 int64 = 253
                                _ = v65
                            } else {
                                v3 = (v3 >> ((4294967295) & 63))
                                switch "\n\t" {
                                case "":
                                    if true {
                                        var v66 int32 = 0
                                        if (65535 == 0) {
                                            var v67 int16 = int16(fi63)
                                            v68 := GenS10{F0: 127, F1: (false && false), F2: 127}
                                            for fi69 := 0; fi69 < 15; fi69++ {
                                                delete(v1, (fi69 >> ((65535) & 63)))
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi69, v66, fi63)
                                                }
                                                _ = fi69
                                            }
                                            _ = v67
                                            _ = v68
                                        } else {
                                            if true {
                                                var v70 uint64 = (v3 - v3)
                                                v71 := GenS4{F0: (v5 ^ (^(v5)))}
                                                _ = v70
                                                _ = v71
                                            } else {
                                                var v72 bool = (!true)
                                                v73 := GenS11{F0: 32767}
                                                _ = v72
                                                _ = v73
                                            }
                                            var v74 uint16 = 42
                                            _ = v74
                                        }
                                        v75 := v1[(fi63 & 1)]
                                        _ = v66
                                        _ = v75
                                    } else {
                                        var v76 uint16 = 255
                                        v76 = (^(v76))
                                        _ = v76
                                    }
                                case "foo bar":
                                    switch fi63 {
                                    case -9223372036854775808:
                                        if true {
                                            if true {
                                                v77 := GenS9{F0: true}
                                                fmt.Println(fi63, v3)
                                                _ = v77
                                            } else {
                                                v78 := GenS12{F0: 9223372036854775807, F1: 1}
                                                v4 = 2147483648
                                                _ = v78
                                            }
                                            v4 = (v4 + v4)
                                            switch "a" {
                                            case "hello":
                                                v4 = (^(v4))
                                            case "\n\t":
                                                v1[(fi63 << ((0) & 63))] = (v4 << ((2) & 63))
                                            default:
                                                if false {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        for fi79 := 0; fi79 < 20; fi79++ {
                                                            if false {
                                                                fmt.Println(v4, fi79)
                                                                var v80 float64 = 1.5
                                                                var v81 uint32 = (^(v4))
                                                                fmt.Println(fi63, v4, v4)
                                                                v82 := GenS8{F0: 4294967295, F1: 1, F2: 1}
                                                                _ = v80
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            switch -1 {
                                                            case 1:
                                                                v5 = (v5 + v5)
                                                            }
                                                            _ = fi79
                                                        }
                                                    default:
                                                        for fi83 := 0; fi83 < 12; fi83++ {
                                                            var v84 int = fi83
                                                            v3 = uint64(v4)
                                                            _ = fi83
                                                            _ = v84
                                                        }
                                                    }
                                                    var v85 uint = 0
                                                    v86 := GenS12{F0: 0, F1: 255}
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    for fi87 := 0; fi87 < 18; fi87++ {
                                                        if true {
                                                            switch fi63 {
                                                            case 2147483647:
                                                                var v88 float32 = 3.14159
                                                                _ = v88
                                                            case -126:
                                                                var v89 int16 = (v5 + (-(v5)))
                                                                _ = v89
                                                            case 1:
                                                                v4 += 315
                                                            default:
                                                                if (42 <= 127) {
                                                                    if true {
                                                                        fmt.Println(v5, fi87)
                                                                        delete(v2, "")
                                                                        if false {
                                                                            if true {
                                                                                var v90 float32 = (-2.25 - (3.14159 + 1.5))
                                                                                _ = v90
                                                                            } else {
                                                                                var v91 uint = 255
                                                                                var v92 string = ((("a" + ("foo bar" + "hello")) + "") + "foo bar")
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                            var v93 int32 = 2147483647
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    if (!true) {
                                                                        v94 := GenS3{F0: 42, F1: fi87, F2: 241}
                                                                        _ = v94
                                                                    } else {
                                                                        v4 = 2147483647
                                                                        fmt.Println(v3, v3)
                                                                    }
                                                                    v95 := GenS13{F0: (v4 | v4), F1: false, F2: 255}
                                                                    _ = v95
                                                                }
                                                            }
                                                        }
                                                        _ = fi87
                                                    }
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v96 float64 = (0.1 - (-2.25 * -2.25))
                                                        _ = v96
                                                    case "hello":
                                                        for fi97 := 0; fi97 < 6; fi97++ {
                                                            var v98 int32 = 2147483647
                                                            _ = fi97
                                                            _ = v98
                                                        }
                                                    }
                                                    v6 = append(v6, 32767)
                                                }
                                            }
                                        } else {
                                            if true {
                                                v1[-9223372036854775808] = (^(v4))
                                            }
                                            v3 -= (^(v3))
                                            v99 := GenS10{F0: -34, F1: true, F2: 127}
                                            _ = v99
                                        }
                                    case 4294967295:
                                        v100 := GenS14{F0: 252}
                                        _ = v100
                                    case 32768:
                                        v4 += v4
                                    default:
                                        var v101 int16 = 1
                                        _ = v101
                                    }
                                }
                            }
                            v5 += v5
                            _ = fi63
                        }
                    default:
                        for fi102 := 0; fi102 < 11; fi102++ {
                            var v103 uint16 = 65535
                            v104 := GenS2{F0: v4}
                            switch ("\n\t" + "") {
                            case "\n\t":
                                var v105 uint32 = v4
                                _ = v105
                            case "foo bar":
                                var v106 uint16 = uint16(v4)
                                _ = v106
                            }
                            _ = fi102
                            _ = v103
                            _ = v104
                        }
                    }
                }
            }
        }
    }
    v4 = (v4 * (v4 ^ v4))
    for fi107 := 0; fi107 < 7; fi107++ {
        v4 += (v4 << ((18446744073709551615) & 63))
        _ = fi107
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
