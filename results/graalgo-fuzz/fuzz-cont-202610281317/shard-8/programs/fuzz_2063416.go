package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint; F1 int8; F2 bool }
type GenS10 struct { F0 uint64; F1 int; F2 int32 }
func genFn0() uint16 {
    switch 0 {
    case 9223372036854775807:
        if false {
            v7 := []int{}
            switch 128 {
            case -9223372036854775808:
                var v8 uint = 854
                _ = v8
            case -925:
                var v9 uint = 1
                _ = v9
            default:
                var v10 uint32 = 807
                _ = v10
            }
            _ = v7
        }
    }
    switch 0 {
    case 80:
        v11 := []int{-915, 127}
        _ = v11
    }
    return 0
}
func genFn1(p0 uint) int32 {
    v12 := []int{-1, 42, 2147483647}
    for fi13 := 0; fi13 < 7; fi13++ {
        if (-2.25 < -2.25) {
            if (false && (-418 < 2147483647)) {
                for fi14 := 0; fi14 < 18; fi14++ {
                    var v15 float64 = -2.25
                    _ = fi14
                    _ = v15
                }
                fmt.Println(fi13, p0)
            } else {
                v16 := GenS1{F0: fi13}
                _ = v16
            }
        } else {
            v17 := GenS1{F0: (fi13 * fi13)}
            var v18 string = ("foo bar" + "hello")
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            p0 += (^(p0))
            var v21 int8 = int8(fi19)
            _ = fi20
            _ = v21
        }
        p0 = (p0 | p0)
        _ = fi19
    }
    if (true || false) {
        var v22 uint32 = 0
        if ((-449 >= -9223372036854775808) && (true && (false || false))) {
            if len(v12) > 0 {
                v12[0] = 583
            }
            for fi23 := 0; fi23 < 3; fi23++ {
                for fi24 := 0; fi24 < 11; fi24++ {
                    if (1 >= 42) {
                        var v25 uint8 = 255
                        if true {
                            v22 = 2147483647
                        } else {
                            v26 := GenS1{F0: fi24}
                            fmt.Println(v22, v25, v22)
                            _ = v26
                        }
                        v27 := GenS2{F0: uint8(v22)}
                        _ = v25
                        _ = v27
                    }
                    _ = fi24
                }
                for fi28 := 0; fi28 < 13; fi28++ {
                    if true {
                        v29 := GenS3{F0: 2, F1: 65535}
                        if ("a" == "\n\t") {
                            fmt.Println(fi23)
                            v30 := GenS3{F0: 260, F1: 65535}
                            if true {
                                var v31 uint32 = 2
                                _ = v31
                            } else {
                                var v32 uint64 = 18446744073709551615
                                _ = v32
                            }
                            var v33 int32 = 1
                            if false {
                                if false {
                                    if (false && true) {
                                        var v34 uint16 = 0
                                        switch "\n\t" {
                                        case "hello":
                                            v34 = (^(v34))
                                        case "":
                                            v35 := GenS4{F0: v33, F1: 550}
                                            _ = v35
                                        case "a":
                                            var v36 string = "\n\t"
                                            _ = v36
                                        default:
                                            var v37 uint64 = 990
                                            _ = v37
                                        }
                                        v22 = v22
                                        _ = v34
                                    }
                                    v38 := GenS2{F0: 42}
                                    _ = v38
                                } else {
                                    if ((v22 & 979) <= (v22 | 164)) {
                                        if len(v12) > 0 {
                                            v12[0] = 582
                                        }
                                        v39 := GenS3{F0: 2, F1: 255}
                                        _ = v39
                                    }
                                    v33 = (-(v33))
                                }
                                fmt.Println(fi28, fi28, v22)
                                if ((!true) && false) {
                                    p0 = p0
                                } else {
                                    var v40 uint = (p0 << ((p0) & 63))
                                    _ = v40
                                }
                            } else {
                                v22 = uint32(p0)
                                p0 += p0
                                var v41 int16 = 1
                                _ = v41
                            }
                            _ = v30
                            _ = v33
                        } else {
                            var v42 uint64 = 0
                            v12 = append(v12, (fi28 * (fi23 % ((fi23) | 1))))
                            v43 := GenS1{F0: (fi28 & (-(fi28)))}
                            _ = v42
                            _ = v43
                        }
                        if (("hello" <= ("\n\t" + ("foo bar" + (("\n\t" + "\n\t") + "a")))) && (65535 < (^(p0)))) {
                            var v44 int32 = -1
                            fmt.Println(v22)
                            fmt.Println(p0, p0)
                            _ = v44
                        } else {
                            fmt.Println(p0, fi28, fi28)
                            fmt.Println(p0, fi23)
                            if true {
                                if false {
                                    p0 = (p0 % ((1) | 1))
                                } else {
                                    v45 := GenS1{F0: 100}
                                    _ = v45
                                }
                                p0 = p0
                                if len(v12) > 0 {
                                    v12[0] = fi23
                                }
                            } else {
                                v22 = v22
                                fmt.Println(fi23)
                            }
                        }
                        if (!false) {
                            switch "" {
                            case "hello":
                                switch (^(fi23)) {
                                case 127:
                                    if true {
                                        _ = len(v12)
                                        v12 = append(v12, 1)
                                        v22 += v22
                                        switch (fi23 + fi23) {
                                        case 396:
                                            v46 := GenS1{F0: (fi28 * 1)}
                                            _ = v46
                                        case 447:
                                            var v47 uint8 = 127
                                            _ = v47
                                        case -399:
                                            v48 := GenS4{F0: int32(p0), F1: 498}
                                            _ = v48
                                        }
                                        var v49 uint = (^(p0))
                                        var v50 uint32 = (^(v22))
                                        _ = v49
                                        _ = v50
                                    } else {
                                        _ = len(v12)
                                        var v51 int64 = 612
                                        _ = v51
                                    }
                                default:
                                    if len(v12) > 0 {
                                        v12[0] = (fi23 ^ fi28)
                                    }
                                }
                            case "a":
                                v22 -= (v22 - v22)
                                var v52 int64 = 42
                                _ = v52
                            case "":
                                switch fi23 {
                                case 2147483647:
                                    fmt.Println(p0, v22)
                                }
                            default:
                                p0 = p0
                            }
                            var v53 int32 = 0
                            var v54 int64 = 0
                            _ = v53
                            _ = v54
                        } else {
                            var v55 uint64 = 1
                            fmt.Println(v55, fi28, p0)
                            v55 -= 4294967295
                            _ = v55
                        }
                        switch "\n\t" {
                        case "foo bar":
                            if len(v12) > 0 {
                                v12[0] = fi23
                            }
                        case "a":
                            if (187 != 480) {
                                fmt.Println(fi28, fi23)
                                switch ((("hello" + "\n\t") + "foo bar") + "\n\t") {
                                case "hello":
                                    switch (fi23 << (((p0 >> ((p0) & 63))) & 63)) {
                                    case 959:
                                        p0 = (p0 * p0)
                                    case -993:
                                        var v56 int = (fi23 | 127)
                                        _ = v56
                                    case 32767:
                                        var v57 uint16 = 42
                                        _ = v57
                                    default:
                                        v22 = uint32(p0)
                                    }
                                default:
                                    var v58 int16 = -1
                                    _ = v58
                                }
                                var v59 int32 = -954
                                fmt.Println(p0, fi28, fi28)
                                _ = v59
                            } else {
                                p0 -= p0
                            }
                        case "\n\t":
                            if (false || false) {
                                p0 += p0
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    if (-13 < -1) {
                                        switch "hello" {
                                        case "a":
                                            p0 = p0
                                        }
                                        var v60 int32 = 2147483647
                                        if ((42 >= 32768) || false) {
                                            v60 = (v60 / ((int32(fi23)) | 1))
                                            if len(v12) > 0 {
                                                v12[0] = fi23
                                            }
                                        }
                                        _ = v60
                                    } else {
                                        if len(v12) > 0 {
                                            v12[0] = fi23
                                        }
                                        var v61 uint = (p0 & p0)
                                        _ = v61
                                    }
                                    v62 := GenS3{F0: 945, F1: 2}
                                    _ = v62
                                default:
                                    p0 = p0
                                }
                                v63 := GenS5{F0: 42, F1: ((false || true) && (true || (true || (((254 >= 0) && false) || (true || (true && false))))))}
                                p0 = (p0 + 18446744073709551615)
                                var v64 int64 = 9223372036854775807
                                if true {
                                    var v65 uint8 = 128
                                    var v66 string = "\n\t"
                                    if ((fi23 / ((fi28) | 1)) <= (fi28 & (fi28 & fi23))) {
                                        var v67 uint16 = 2
                                        v67 = (v67 ^ genFn0())
                                        _ = v67
                                    }
                                    _ = v65
                                    _ = v66
                                } else {
                                    var v68 int8 = -128
                                    switch fi28 {
                                    case 128:
                                        v64 += (v64 % ((-1) | 1))
                                    default:
                                        switch (fi28 / ((-638) | 1)) {
                                        case -850:
                                            v64 = 915
                                        }
                                    }
                                    _ = v68
                                }
                                _ = v63
                                _ = v64
                            }
                        default:
                            if (1.5 != 3.14159) {
                                if len(v12) > 0 {
                                    v12[0] = 1
                                }
                                switch (^(fi28)) {
                                case 128:
                                    switch (fi23 - fi28) {
                                    case 128:
                                        v69 := GenS3{F0: 4294967295, F1: 42}
                                        _ = v69
                                    case 0:
                                        if true {
                                            fmt.Println(fi28, p0, p0)
                                        } else {
                                            if true {
                                                p0 += 255
                                                if false {
                                                    v22 += (^(v22))
                                                    switch (fi23 >> (((p0 + (p0 | p0))) & 63)) {
                                                    case -520:
                                                        var v70 float64 = (1.5 / 1e10)
                                                        _ = v70
                                                    case 128:
                                                        var v71 uint8 = 161
                                                        if len(v12) > 0 {
                                                            v12[0] = fi28
                                                        }
                                                        var v72 uint8 = v71
                                                        _ = v71
                                                        _ = v72
                                                    case 2147483647:
                                                        v73 := GenS6{F0: ("hello" + "foo bar")}
                                                        _ = v73
                                                    default:
                                                        if false {
                                                            var v74 int = fi23
                                                            fmt.Println(fi23, v74, p0)
                                                            p0 = 18446744073709551615
                                                            _ = v74
                                                        }
                                                        var v75 uint = 42
                                                        _ = v75
                                                    }
                                                } else {
                                                    v22 = v22
                                                    fmt.Println(fi23, fi23, p0)
                                                }
                                                p0 = (p0 & p0)
                                            }
                                            p0 -= p0
                                        }
                                    }
                                case 127:
                                    if len(v12) > 0 {
                                        v12[0] = (-(fi23))
                                    }
                                case 9223372036854775807:
                                    _ = len(v12)
                                default:
                                    var v76 string = "\n\t"
                                    _ = v76
                                }
                                var v77 uint64 = 472
                                _ = v77
                            }
                        }
                        v78 := GenS1{F0: fi23}
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                        v22 += v22
                        _ = v29
                        _ = v78
                    } else {
                        _ = len(v12)
                        v22 = 2
                    }
                    _ = fi28
                }
                v79 := GenS7{F0: -116}
                _ = fi23
                _ = v79
            }
        }
        var v80 float32 = 3.14159
        _ = v22
        _ = v80
    }
    _ = p0
    _ = v12
    return -809
}

func main() {
    var v1 uint64 = 565
    v2 := make(map[int]int)
    var v3 int32 = 42
    v4 := make(map[int]string)
    var v5 int8 = 1
    var v6 int32 = 133
    v3 += (v6 + genFn1(643))
    fmt.Println(v5, v1, v1)
    fmt.Println(v5, v1, v1)
    v81 := GenS8{F0: 772, F1: 18446744073709551615}
    v82 := GenS7{F0: -1}
    v83 := GenS9{F0: 537, F1: 127, F2: (!false)}
    switch "hello" {
    case "hello":
        for fi84 := 0; fi84 < 2; fi84++ {
            switch fi84 {
            case 1:
                var v85 int8 = (v5 << ((82) & 63))
                _ = v85
            case 32767:
                var v86 int32 = -432
                _ = v86
            default:
                v87 := v2[0]
                _ = v87
            }
            if (1 < 4294967295) {
                v5 += 100
            } else {
                fmt.Println(v1, v5, v6)
                v88 := []int{(fi84 * 365)}
                if ((v1 << ((737) & 63)) > (v1 << ((1) & 63))) {
                    switch int(v3) {
                    case 727:
                        switch fi84 {
                        case 32768:
                            for fi89 := 0; fi89 < 10; fi89++ {
                                if (1 <= 597) {
                                    switch fi89 {
                                    case -2147483648:
                                        v5 += (v5 + (v5 / ((-128) | 1)))
                                    case 2147483647:
                                        if true {
                                            v90 := GenS1{F0: fi89}
                                            _ = v90
                                        }
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v91 := GenS2{F0: 128}
                                        _ = v91
                                    case "":
                                        if false {
                                            v4[(^(fi89))] = (("foo bar" + (("" + "a") + "\n\t")) + ("\n\t" + "hello"))
                                        }
                                    default:
                                        var v92 int = fi84
                                        _ = v92
                                    }
                                }
                                _ = len(v2)
                                _ = fi89
                            }
                        case -2147483648:
                            for fi93 := 0; fi93 < 11; fi93++ {
                                switch fi93 {
                                case -283:
                                    switch ("foo bar" + "a") {
                                    case "foo bar":
                                        if true {
                                            var v94 bool = false
                                            _ = v94
                                        }
                                    default:
                                        v3 += (v3 - 1)
                                    }
                                }
                                v3 += v3
                                _ = fi93
                            }
                        }
                    default:
                        fmt.Println(v1)
                    }
                } else {
                    var v95 int8 = v5
                    v1 -= (v1 + (v1 * 18446744073709551615))
                    _ = v95
                }
                _ = v88
            }
            _ = fi84
        }
    case "foo bar":
        v96 := []int{128, int(v3), 2147483647, -775}
        _ = v96
    }
    var v97 string = ("a" + "\n\t")
    for fi98 := 0; fi98 < 20; fi98++ {
        var v99 uint64 = v1
        _ = fi98
        _ = v99
    }
    v6 -= (-(v6))
    if true {
        var v100 int16 = 545
        _ = v100
    } else {
        v3 -= genFn1(0)
        if false {
            v101 := GenS4{F0: (v3 - -2147483648), F1: v1}
            _ = v101
        } else {
            switch -559 {
            case 9223372036854775807:
                if (-613 > (v6 | v3)) {
                    delete(v2, 2147483647)
                    fmt.Println(v5)
                }
            case 128:
                v102 := GenS10{F0: 495, F1: -2147483648, F2: -2147483648}
                _ = v102
            case 0:
                for fi103 := 0; fi103 < 20; fi103++ {
                    fmt.Println(v6, v5, v5)
                    fmt.Println(fi103)
                    _ = fi103
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
    _ = v83
    _ = v97
}
