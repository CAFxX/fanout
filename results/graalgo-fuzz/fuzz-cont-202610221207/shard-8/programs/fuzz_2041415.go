package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 uint8; F1 uint8 }
type GenS11 struct { F0 int8; F1 bool; F2 bool }
type GenS12 struct { F0 uint8 }

func main() {
    var v1 int8 = 0
    var v2 bool = true
    v3 := []int{42, 4294967295, 1}
    v4 := make(map[int]bool)
    var v5 uint64 = 4294967295
    var v6 int32 = 68
    _ = len(v3)
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v1 = v1
            switch fi8 {
            case 9223372036854775807:
                var v9 uint8 = 127
                _ = v9
            case -531:
                v10 := GenS1{F0: 1, F1: (!v2)}
                _ = v10
            case 301:
                var v11 int16 = 1
                _ = v11
            default:
                switch "a" {
                case "hello":
                    if v2 {
                        if (-38 <= fi7) {
                            v5 = v5
                            switch (-(fi7)) {
                            case 4294967295:
                                var v12 uint8 = 204
                                _ = v12
                            case -1:
                                var v13 int16 = 1
                                if false {
                                    _ = len(v4)
                                    fmt.Println(v5)
                                    var v14 int16 = (v13 + v13)
                                    v14 += v13
                                    _ = v14
                                } else {
                                    v2 = false
                                    v6 += v6
                                }
                                var v15 string = (("\n\t" + ("hello" + ("hello" + "hello"))) + "foo bar")
                                _ = v13
                                _ = v15
                            default:
                                var v16 int64 = -1
                                _ = v16
                            }
                        } else {
                            if (-846 < 0) {
                                fmt.Println(v6, fi8)
                                fmt.Println(fi8)
                            }
                        }
                        var v17 string = ""
                        _ = v17
                    } else {
                        var v18 uint32 = 851
                        v18 = 15
                        v19 := GenS1{F0: 0, F1: ((float32(0.0) - (0.0 / (1.5 + float32((1e10 + (0.1 * (float64((3.14159 - 1e10)) * 0.0))))))) == 0.0)}
                        _ = v18
                        _ = v19
                    }
                default:
                    v20 := GenS1{F0: 65535, F1: true}
                    _ = v20
                }
                var v21 uint32 = 4294967295
                var v22 int = (fi7 ^ 920)
                v23 := GenS1{F0: 2, F1: (v1 == (v1 | (v1 ^ v1)))}
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi8
        }
        _ = fi7
    }
    fmt.Println(v5)
    delete(v4, 127)
    var v24 uint64 = 2
    switch 2147483647 {
    case 0:
        _ = len(v3)
    case -248:
        fmt.Println(v1)
    case 128:
        if v2 {
            if (!v2) {
                switch 127 {
                case -1:
                    v1 = v1
                case 4294967295:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v4[-1] = (29 == 65535)
                if (("" + "") == "\n\t") {
                    v25 := GenS1{F0: 156, F1: v2}
                    var v26 uint8 = 1
                    _ = v25
                    _ = v26
                }
            }
            if v2 {
                v24 -= v24
            } else {
                v3 = append(v3, -608)
                fmt.Println(v5)
                v4[32768] = v2
            }
        } else {
            fmt.Println(v24, v24, v24)
            if ("hello" < "foo bar") {
                v5 = v24
                for fi27 := 0; fi27 < 16; fi27++ {
                    if v2 {
                        switch (fi27 & fi27) {
                        case 4294967295:
                            v3 = append(v3, fi27)
                        case -536:
                            v2 = v2
                        case -198:
                            for fi28 := 0; fi28 < 13; fi28++ {
                                if (971 <= 1) {
                                    switch (("foo bar" + "foo bar") + ("\n\t" + "foo bar")) {
                                    case "":
                                        var v29 uint = 255
                                        _ = v29
                                    }
                                    v2 = true
                                }
                                var v30 uint = 2
                                _ = fi28
                                _ = v30
                            }
                        }
                    }
                    _ = fi27
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 128
                }
                var v31 int64 = 9223372036854775807
                _ = v31
            }
        }
    default:
        var v32 uint16 = 42
        _ = v32
    }
    fmt.Println(v24)
    if (!(!(-1 >= -32768))) {
        if (42 != 255) {
            for fi33 := 0; fi33 < 19; fi33++ {
                v34 := GenS1{F0: 880, F1: false}
                _ = fi33
                _ = v34
            }
        } else {
            if v2 {
                var v35 int = 127
                var v36 int8 = (-(v1))
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 12; fi37++ {
                    switch (("foo bar" + "\n\t") + "foo bar") {
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = -1
                        }
                    case "a":
                        switch "\n\t" {
                        case "hello":
                            if (v2 || (!v2)) {
                                for fi38 := 0; fi38 < 9; fi38++ {
                                    if (v2 && true) {
                                        var v39 uint = 1
                                        _ = v39
                                    } else {
                                        if (!(2 >= 42)) {
                                            v40 := GenS2{F0: -1, F1: "hello", F2: 42}
                                            delete(v4, (^(fi37)))
                                            if (2 == 42) {
                                                var v41 float32 = -2.25
                                                _ = v41
                                            }
                                            _ = v40
                                        } else {
                                            fmt.Println(v24, v2)
                                            if v2 {
                                                v6 -= v6
                                                v24 = (v24 ^ 673)
                                                v42 := v4[(^(fi37))]
                                                if (32767 >= 1) {
                                                    v43 := GenS2{F0: 42, F1: "hello", F2: 42}
                                                    if true {
                                                        v24 += v24
                                                        v42 = v42
                                                    }
                                                    if true {
                                                        var v44 bool = true
                                                        fmt.Println(v42, v24)
                                                        v24 -= (v5 << ((1) & 63))
                                                        _ = v44
                                                    }
                                                    switch (fi38 * fi38) {
                                                    case -928:
                                                        if ((1 == 18446744073709551615) || true) {
                                                            var v45 uint16 = 273
                                                            var v46 int8 = int8(fi37)
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                    case -9223372036854775808:
                                                        if false {
                                                            var v47 uint = 18446744073709551615
                                                            v48 := GenS3{F0: v6}
                                                            v47 = v47
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    case 566:
                                                        var v49 int64 = 1
                                                        v50 := v4[(-(fi37))]
                                                        _ = v49
                                                        _ = v50
                                                    default:
                                                        v3 = append(v3, 0)
                                                    }
                                                    v2 = (!true)
                                                    _ = v43
                                                } else {
                                                    switch "hello" {
                                                    case "a":
                                                        if v2 {
                                                            var v51 uint32 = 1
                                                            v1 += (-(v1))
                                                            v24 -= v5
                                                            _ = v51
                                                        } else {
                                                            var v52 uint32 = 4294967295
                                                            fmt.Println(fi38, v24, fi37)
                                                            v52 = v52
                                                            _ = v52
                                                        }
                                                    case "foo bar":
                                                        fmt.Println(v5)
                                                    case "\n\t":
                                                        fmt.Println(v24, v2)
                                                    default:
                                                        switch (fi37 - fi38) {
                                                        case 127:
                                                            v2 = (false && false)
                                                        case 42:
                                                            v42 = (!true)
                                                        default:
                                                            v3 = append(v3, fi37)
                                                        }
                                                    }
                                                    v53 := GenS4{F0: 42, F1: -315}
                                                    _ = v53
                                                }
                                                var v54 uint64 = v5
                                                _ = v42
                                                _ = v54
                                            }
                                            switch (fi38 % (((^(fi37))) | 1)) {
                                            case 32768:
                                                switch "\n\t" {
                                                case "":
                                                    v1 = (v1 | 0)
                                                case "hello":
                                                    v1 -= (-(v1))
                                                }
                                            default:
                                                v55 := GenS5{F0: 42}
                                                _ = v55
                                            }
                                        }
                                        v56 := GenS6{F0: (v6 ^ (-(v6))), F1: 255, F2: 0}
                                        _ = v56
                                    }
                                    _ = fi38
                                }
                                switch "hello" {
                                case "\n\t":
                                    for fi57 := 0; fi57 < 1; fi57++ {
                                        v58 := GenS7{F0: 255, F1: "a"}
                                        _ = fi57
                                        _ = v58
                                    }
                                default:
                                    var v59 int8 = (-(v1))
                                    _ = v59
                                }
                            } else {
                                var v60 uint64 = (v5 ^ 4294967295)
                                v61 := GenS8{F0: v1, F1: (v60 % ((2) | 1))}
                                _ = v60
                                _ = v61
                            }
                        case "\n\t":
                            v1 = v1
                        default:
                            switch fi37 {
                            case 42:
                                if true {
                                    switch (("" + ("foo bar" + "")) + "hello") {
                                    case "hello":
                                        for fi62 := 0; fi62 < 10; fi62++ {
                                            fmt.Println(v24, v24, fi37)
                                            if len(v3) > 0 {
                                                v3[0] = fi37
                                            }
                                            _ = len(v3)
                                            var v63 int = (fi37 | fi62)
                                            _ = fi62
                                            _ = v63
                                        }
                                    case "\n\t":
                                        var v64 uint16 = 774
                                        _ = v64
                                    case "":
                                        var v65 float64 = (((0.1 - (3.14159 - -2.25)) - (1e10 + ((1.5 + 0.0) - (3.14159 * (1.5 - (float64(float32(0.1)) - 0.0)))))) - (0.0 - (((float64(3.14159) * 3.14159) * 0.0) * 1e10)))
                                        _ = v65
                                    default:
                                        for fi66 := 0; fi66 < 19; fi66++ {
                                            v24 += v24
                                            if ((v1 | v1) < int8(v6)) {
                                                switch "" {
                                                case "hello":
                                                    switch fi37 {
                                                    case 2147483647:
                                                        v1 -= (v1 / (((v1 | 1)) | 1))
                                                    case -1:
                                                        if v2 {
                                                            v24 = 567
                                                        } else {
                                                            switch ("\n\t" + "\n\t") {
                                                            case "":
                                                                v67 := GenS9{F0: v24, F1: 373}
                                                                _ = v67
                                                            default:
                                                                v1 = v1
                                                            }
                                                            v68 := GenS4{F0: 42, F1: -915}
                                                            _ = v68
                                                        }
                                                    case 0:
                                                        if ((true || false) && ((("a" < ("" + ((("" + ("\n\t" + "\n\t")) + "") + "foo bar"))) || (854 <= -688)) && (!(true && (!v2))))) {
                                                            if (42 == -32768) {
                                                                v5 += 743
                                                                v1 += 34
                                                                var v69 uint64 = v24
                                                                _ = v69
                                                            }
                                                        }
                                                    default:
                                                        var v70 float64 = 1.5
                                                        _ = v70
                                                    }
                                                }
                                                v4[9223372036854775807] = ((2147483648 >= 2) && (((!((fi66 & fi37) >= fi66)) && (v1 == (v1 % ((42) | 1)))) && v2))
                                            } else {
                                                var v71 bool = (!(!((1 == 2) && false)))
                                                if (1e10 == 0.1) {
                                                    v72 := GenS7{F0: 2, F1: ("hello" + "\n\t")}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        var v73 float64 = 1e10
                                                        if ((v2 || (2 != 508)) || (2 > 930)) {
                                                            var v74 bool = false
                                                            var v75 uint64 = v24
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        v76 := GenS9{F0: (v5 % ((132) | 1)), F1: 0}
                                                        _ = v73
                                                        _ = v76
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v72
                                                }
                                                var v77 uint16 = 861
                                                _ = v71
                                                _ = v77
                                            }
                                            _ = fi66
                                        }
                                    }
                                    var v78 uint16 = 2
                                    _ = v78
                                } else {
                                    if false {
                                        v1 = (v1 & 1)
                                        v2 = v2
                                    }
                                    var v79 uint16 = 639
                                    _ = v79
                                }
                            case 258:
                                if true {
                                    for fi80 := 0; fi80 < 1; fi80++ {
                                        var v81 int = fi37
                                        var v82 uint32 = 2
                                        var v83 int8 = 0
                                        _ = fi80
                                        _ = v81
                                        _ = v82
                                        _ = v83
                                    }
                                }
                            }
                        }
                    default:
                        var v84 int64 = -9223372036854775808
                        _ = v84
                    }
                    var v85 int32 = -1
                    _ = fi37
                    _ = v85
                }
                v86 := GenS10{F0: 10, F1: 0}
                var v87 uint32 = 4294967295
                _ = v86
                _ = v87
            }
            var v88 int16 = -548
            for fi89 := 0; fi89 < 11; fi89++ {
                v90 := v4[(fi89 % ((fi89) | 1))]
                for fi91 := 0; fi91 < 10; fi91++ {
                    var v92 int8 = (v1 - (v1 ^ v1))
                    if false {
                        v88 = v88
                        var v93 int8 = 127
                        v94 := GenS11{F0: 127, F1: (v2 || v90), F2: v90}
                        _ = v93
                        _ = v94
                    } else {
                        v95 := GenS12{F0: 1}
                        fmt.Println(v90)
                        var v96 string = "a"
                        _ = v95
                        _ = v96
                    }
                    v92 += v92
                    _ = fi91
                    _ = v92
                }
                var v97 string = "hello"
                _ = fi89
                _ = v90
                _ = v97
            }
            _ = v88
        }
        delete(v4, 4294967295)
    } else {
        v4[2147483647] = (true && (42 > (v1 << ((2) & 63))))
        v6 = v6
        var v98 int16 = 42
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
