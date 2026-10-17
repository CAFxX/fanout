package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 string; F1 uint32; F2 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 string; F1 uint16; F2 uint32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 uint; F1 int32 }
func genFn0(p0 int8, p1 int32) bool {
    var v5 int32 = p1
    var v6 float64 = 3.14159
    v6 = v6
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 0, F1: (fi7 / (((fi7 & fi7)) | 1)), F2: 42}
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (!false)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 936
    v3 := make(map[uint32]bool)
    var v4 int = 183
    switch "foo bar" {
    case "":
        if true {
            if true {
                v9 := []int{(^(v4)), v4, v4}
                v4 -= 42
                _ = v9
            } else {
                switch "" {
                case "\n\t":
                    var v10 float32 = 0.1
                    _ = v10
                case "foo bar":
                    var v11 int8 = int8(v2)
                    _ = v11
                case "":
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(v4, fi12, fi12)
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v4 -= (fi13 % (((fi13 / ((fi13) | 1))) | 1))
                            _ = fi13
                        }
                        v4 = 42
                        _ = fi12
                    }
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v4 = fi15
                        if true {
                            v4 += fi15
                            v16 := GenS1{F0: 255, F1: fi15, F2: 32767}
                            _ = v16
                        } else {
                            v17 := GenS2{F0: 4294967295}
                            _ = v17
                        }
                        v4 += fi15
                        _ = fi15
                    }
                    _ = fi14
                }
                v18 := GenS2{F0: 298}
                _ = v18
            }
            var v19 int32 = 2147483647
            v4 -= 1
            _ = v19
        } else {
            if true {
                if ((4294967295 > 42) && (("hello" + (("a" + "a") + "a")) >= ("foo bar" + (((("hello" + "a") + "foo bar") + "foo bar") + (("\n\t" + ("hello" + "a")) + "a"))))) {
                    v2 = 348
                    v2 = 65535
                } else {
                    v20 := GenS3{F0: "foo bar", F1: 0}
                    _ = v20
                }
            } else {
                var v21 int32 = -1
                _ = v21
            }
            switch v4 {
            case -608:
                _ = len(v1)
            case 127:
                if ((652 <= 523) && true) {
                    if (v2 == 2147483648) {
                        if (true && (-306 >= -1)) {
                            if true {
                                v22 := GenS4{F0: 1, F1: 42, F2: 918}
                                v2 = (^(v2))
                                _ = v22
                            } else {
                                switch ("foo bar" + ("a" + "hello")) {
                                case "hello":
                                    fmt.Println(v2)
                                case "foo bar":
                                    if ((1e10 / 1.0) <= 3.14159) {
                                        switch v4 {
                                        case -158:
                                            var v23 int8 = 127
                                            _ = v23
                                        case 127:
                                            var v24 uint16 = 32767
                                            _ = v24
                                        default:
                                            switch (v4 | (v4 >> ((1) & 63))) {
                                            case 1:
                                                v2 -= v2
                                            case 42:
                                                v25 := GenS5{F0: 42, F1: 1}
                                                _ = v25
                                            case 32768:
                                                fmt.Println(v2, v4)
                                            default:
                                                for fi26 := 0; fi26 < 11; fi26++ {
                                                    v2 = v2
                                                    var v27 int8 = 38
                                                    _ = fi26
                                                    _ = v27
                                                }
                                            }
                                        }
                                        if true {
                                            v4 = -362
                                        } else {
                                            v28 := []int{127, v4}
                                            var v29 float32 = ((float32(3.14159) + (-2.25 * 0.1)) - -2.25)
                                            switch (("" + "foo bar") + ("hello" + "a")) {
                                            case "a":
                                                v2 -= v2
                                            case "\n\t":
                                                for fi30 := 0; fi30 < 13; fi30++ {
                                                    if true {
                                                        v4 -= -1
                                                    } else {
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v31 float64 = 0.1
                                                            _ = v31
                                                        case "hello":
                                                            switch fi30 {
                                                            case 42:
                                                                switch (v4 - (v4 >> ((787) & 63))) {
                                                                case -1:
                                                                    fmt.Println(v29, v2, fi30)
                                                                case -198:
                                                                    var v32 int16 = -179
                                                                    _ = v32
                                                                case 9223372036854775807:
                                                                    v4 = (v4 >> ((656) & 63))
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v33 float64 = 3.14159
                                                            _ = v33
                                                        }
                                                    }
                                                    var v34 bool = (!true)
                                                    _ = fi30
                                                    _ = v34
                                                }
                                            case "hello":
                                                switch (v4 * (v4 + (-(v4)))) {
                                                case 9223372036854775807:
                                                    v35 := GenS1{F0: 177, F1: (v4 << ((350) & 63)), F2: 42}
                                                    _ = v35
                                                case -638:
                                                    fmt.Println(v29)
                                                }
                                            }
                                            _ = v28
                                            _ = v29
                                        }
                                    }
                                case "\n\t":
                                    var v36 int8 = -128
                                    _ = v36
                                default:
                                    fmt.Println(v4, v4)
                                }
                            }
                            if ((v4 - v4) > (v4 % (((v4 << ((18446744073709551615) & 63))) | 1))) {
                                v2 = (v2 / ((287) | 1))
                            }
                            if false {
                                for fi37 := 0; fi37 < 16; fi37++ {
                                    switch "" {
                                    case "a":
                                        switch "hello" {
                                        case "hello":
                                            v2 = v2
                                        case "\n\t":
                                            if true {
                                                for fi38 := 0; fi38 < 7; fi38++ {
                                                    switch (-(v4)) {
                                                    case 127:
                                                        v39 := []int{fi37}
                                                        _ = v39
                                                    case -8:
                                                        fmt.Println(v4, v4)
                                                    case 32767:
                                                        var v40 string = "\n\t"
                                                        _ = v40
                                                    default:
                                                        var v41 bool = (1 != -1)
                                                        _ = v41
                                                    }
                                                    v42 := GenS5{F0: 65535, F1: 0}
                                                    _ = fi38
                                                    _ = v42
                                                }
                                            } else {
                                                v4 = v4
                                            }
                                        case "foo bar":
                                            if (false || (!(1 >= int32(v4)))) {
                                                v43 := []int{fi37, v4, 128, v4}
                                                if (-2.25 <= 0.1) {
                                                    for fi44 := 0; fi44 < 12; fi44++ {
                                                        fmt.Println(v4, fi44)
                                                        v4 -= 32768
                                                        var v45 float64 = 1.5
                                                        _ = fi44
                                                        _ = v45
                                                    }
                                                    fmt.Println(v2, v2)
                                                    var v46 int16 = 894
                                                    _ = v46
                                                } else {
                                                    v47 := GenS6{F0: -28, F1: -231}
                                                    if ("" > ("hello" + "\n\t")) {
                                                        for fi48 := 0; fi48 < 8; fi48++ {
                                                            if (!(127 <= 2)) {
                                                                v2 -= v2
                                                                v49 := GenS1{F0: uint16(v2), F1: (^(fi48)), F2: 255}
                                                                v4 += (fi37 % (((-(fi37))) | 1))
                                                                _ = v49
                                                            } else {
                                                                v4 += -369
                                                            }
                                                            _ = fi48
                                                        }
                                                    }
                                                    _ = v47
                                                }
                                                _ = v43
                                            }
                                        }
                                    case "\n\t":
                                        v50 := GenS7{F0: "\n\t", F1: (v2 << ((4294967295) & 63)), F2: (((!(-33 >= 42)) && true) || (128 >= 436))}
                                        _ = v50
                                    default:
                                        fmt.Println(fi37)
                                    }
                                    v51 := GenS8{F0: 550}
                                    v52 := GenS3{F0: "foo bar", F1: v2}
                                    _ = fi37
                                    _ = v51
                                    _ = v52
                                }
                                var v53 uint = 2
                                _ = v53
                            } else {
                                v4 -= (v4 % ((4294967295) | 1))
                                v54 := []int{}
                                _ = v54
                            }
                        } else {
                            fmt.Println(v4, v4)
                            v4 = (v4 ^ v4)
                            if false {
                                v4 = 32767
                            } else {
                                v2 = 2
                                v55 := GenS1{F0: 444, F1: v4, F2: 2}
                                _ = v55
                            }
                        }
                        fmt.Println(v2, v2)
                    }
                    v2 = (v2 & v2)
                } else {
                    fmt.Println(v4, v2)
                    v56 := GenS1{F0: 217, F1: v4, F2: 65535}
                    _ = v56
                }
            }
        }
    case "foo bar":
        v4 = (v4 / ((v4) | 1))
    case "hello":
        v2 = (v2 & (v2 >> ((0) & 63)))
    }
    v4 -= -587
    if (246 <= 18446744073709551615) {
        v4 -= (v4 - 469)
        v57 := []int{v4, v4, (v4 << ((859) & 63)), v4}
        _ = v57
    } else {
        switch v4 {
        case -549:
            v4 = v4
        case -9223372036854775808:
            switch "\n\t" {
            case "":
                var v58 int64 = -217
                _ = v58
            }
        }
        v4 = 613
    }
    v59 := GenS6{F0: -69, F1: 119}
    if false {
        switch ("hello" + "a") {
        case "\n\t":
            v60 := v3[(^(v2))]
            _ = v60
        case "":
            fmt.Println(v4)
        default:
            v61 := v3[v2]
            _ = v61
        }
    } else {
        for fi62 := 0; fi62 < 16; fi62++ {
            if false {
                switch 42 {
                case 937:
                    var v63 int = fi62
                    _ = v63
                case 182:
                    v4 -= v4
                case 127:
                    if false {
                        var v64 int8 = 29
                        _ = v64
                    } else {
                        if (267 >= uint16(v2)) {
                            fmt.Println(v4, fi62)
                            if true {
                                v2 = v2
                                v2 = (v2 >> ((0) & 63))
                            } else {
                                if false {
                                    if (127 == -1) {
                                        switch v4 {
                                        case 130:
                                            for fi65 := 0; fi65 < 11; fi65++ {
                                                fmt.Println(v4, v2)
                                                v66 := GenS9{F0: 0}
                                                v2 += 2147483647
                                                _ = fi65
                                                _ = v66
                                            }
                                        case -27:
                                            var v67 uint = 2
                                            _ = v67
                                        case 127:
                                            if (711 < -601) {
                                                var v68 float64 = 0.1
                                                _ = v68
                                            }
                                        default:
                                            fmt.Println(fi62, fi62)
                                        }
                                    } else {
                                        var v69 bool = (false && true)
                                        var v70 uint16 = 32767
                                        _ = v69
                                        _ = v70
                                    }
                                }
                                for fi71 := 0; fi71 < 16; fi71++ {
                                    switch (v4 & fi71) {
                                    case -542:
                                        fmt.Println(fi71)
                                    }
                                    switch fi62 {
                                    case 2147483647:
                                        v2 = v2
                                    case -303:
                                        v72 := GenS4{F0: 127, F1: 1, F2: 131}
                                        _ = v72
                                    case 4294967295:
                                        if (148 == 42) {
                                            if false {
                                                v73 := GenS8{F0: 255}
                                                _ = v73
                                            }
                                            var v74 uint16 = 2
                                            _ = v74
                                        } else {
                                            v4 -= (v4 | v4)
                                        }
                                    default:
                                        v4 = fi62
                                    }
                                    _ = fi71
                                }
                            }
                            switch v4 {
                            case 2147483647:
                                var v75 int16 = 42
                                _ = v75
                            case 32767:
                                v2 -= v2
                            case 4294967295:
                                v76 := GenS9{F0: 752}
                                _ = v76
                            default:
                                for fi77 := 0; fi77 < 14; fi77++ {
                                    var v78 uint16 = 255
                                    v79 := GenS1{F0: 42, F1: 9223372036854775807, F2: (^(v78))}
                                    var v80 uint16 = (v78 % (((v78 >> ((42) & 63))) | 1))
                                    v80 = v78
                                    v4 = fi77
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                            }
                        }
                        var v81 float32 = 3.14159
                        for fi82 := 0; fi82 < 9; fi82++ {
                            var v83 uint8 = 42
                            _ = fi82
                            _ = v83
                        }
                        _ = v81
                    }
                }
                fmt.Println(fi62)
                for fi84 := 0; fi84 < 14; fi84++ {
                    delete(v3, (v2 << ((18446744073709551615) & 63)))
                    fmt.Println(fi62, v4)
                    v4 += -9223372036854775808
                    v85 := GenS10{F0: ("" + (("foo bar" + "") + "hello")), F1: 0, F2: (v2 | 2)}
                    _ = fi84
                    _ = v85
                }
            }
            fmt.Println(v4)
            var v86 bool = (v2 == 42)
            _ = fi62
            _ = v86
        }
        var v87 uint = 719
        _ = v87
    }
    var v88 int = v4
    v88 += (^(v4))
    v89 := GenS11{F0: 42}
    if true {
        for fi90 := 0; fi90 < 10; fi90++ {
            v91 := []int{4294967295, v4, (fi90 % ((fi90) | 1))}
            _ = len(v91)
            var v92 int16 = 42
            _ = fi90
            _ = v91
            _ = v92
        }
        switch (^(v88)) {
        case 9223372036854775807:
            switch "a" {
            case "":
                v4 -= 4294967295
            case "\n\t":
                for fi93 := 0; fi93 < 16; fi93++ {
                    v88 = (v88 - fi93)
                    _ = fi93
                }
            case "a":
                v94 := GenS12{F0: (v2 % (((v2 ^ 1)) | 1)), F1: -9223372036854775808, F2: (^(v2))}
                _ = v94
            default:
                v4 += v88
            }
        case -278:
            fmt.Println(v4, v88)
        case 1:
            v95 := []int{v88, v88, (v88 / ((v4) | 1)), (v4 % ((4294967295) | 1))}
            _ = v95
        }
    } else {
        var v96 uint64 = 1
        v4 -= (v88 / (((-(v4))) | 1))
        _ = v96
    }
    if (!(false && (((375 <= 18446744073709551615) || true) || (!(false && (!false)))))) {
        v88 -= (v88 / ((v4) | 1))
        v97 := GenS13{F0: (!(255 >= 65535)), F1: 0}
        _ = v97
    }
    var v98 int32 = 42
    for fi99 := 0; fi99 < 1; fi99++ {
        for fi100 := 0; fi100 < 15; fi100++ {
            fmt.Println(fi99, v4)
            _ = fi100
        }
        v88 -= (v4 >> ((796) & 63))
        v88 += v4
        _ = fi99
    }
    v101 := GenS14{F0: 65535, F1: v98}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v88:", v88)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v88
    _ = v89
    _ = v98
}
