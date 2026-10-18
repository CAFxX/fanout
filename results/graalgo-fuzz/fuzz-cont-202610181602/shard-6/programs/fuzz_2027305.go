package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 bool; F1 string; F2 int }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int8; F1 string; F2 int64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int; F1 uint32 }
type GenS11 struct { F0 uint32; F1 bool }
func genFn0(p0 uint32) uint64 {
    p0 = p0
    v7 := GenS1{F0: "foo bar"}
    var v8 uint64 = 4294967295
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{0, 42}
    var v3 int = 564
    var v4 int8 = 71
    var v5 uint16 = 65535
    var v6 uint32 = 42
    var v9 uint8 = 128
    if ("a" != "hello") {
        var v10 bool = (genFn0((v6 ^ v6)) != 4294967295)
        fmt.Println(v6, v10, v9)
        _ = v10
    }
    var v11 string = "\n\t"
    if false {
        v12 := GenS2{F0: 32767, F1: 42}
        switch (^(v3)) {
        case 798:
            var v13 int16 = 388
            _ = v13
        case -317:
            for fi14 := 0; fi14 < 14; fi14++ {
                v15 := GenS3{F0: 406, F1: -331}
                _ = fi14
                _ = v15
            }
        }
        _ = v12
    }
    v5 = v5
    if false {
        var v16 uint64 = 317
        _ = v16
    }
    if (true && false) {
        var v17 uint32 = (v6 % ((65535) | 1))
        switch v3 {
        case -271:
            for fi18 := 0; fi18 < 15; fi18++ {
                v9 -= (v9 / (((^(v9))) | 1))
                fmt.Println(v4, fi18, v9)
                v6 = v17
                _ = fi18
            }
        case 1:
            var v19 uint16 = (v5 / (((v5 + (v5 ^ 1))) | 1))
            _ = v19
        }
        fmt.Println(v4)
        _ = v17
    }
    v20 := GenS4{F0: 751, F1: v9}
    var v21 float32 = (((((3.14159 + 0.0) - 0.0) * 1e10) * 0.0) / 0.1)
    for fi22 := 0; fi22 < 2; fi22++ {
        v1[42] = (v3 / ((1) | 1))
        if true {
            v9 = (^(v9))
            if ((v9 & 255) < v9) {
                for fi23 := 0; fi23 < 13; fi23++ {
                    if (((v21 / 1.0) * 0.1) < float32(float64(float32(((1e10 * 0.0) * (1e10 / 0.1)))))) {
                        var v24 int = fi22
                        switch v11 {
                        case "\n\t":
                            v25 := GenS5{F0: (-1 > 1), F1: v11, F2: -2147483648}
                            _ = v25
                        case "foo bar":
                            v21 = (v21 * float32(0.0))
                            if (!true) {
                                if true {
                                    v4 += (v4 % ((0) | 1))
                                    var v26 int = 127
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(fi23)
                                    default:
                                        v27 := GenS5{F0: ((v6 / ((v6) | 1)) > 0), F1: (v11 + v11), F2: fi23}
                                        _ = v27
                                    }
                                    v24 -= v26
                                    if true {
                                        var v28 float64 = (float64(v21) - 0.0)
                                        v29 := GenS1{F0: v11}
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v28
                                        _ = v29
                                    } else {
                                        v30 := GenS6{F0: 0, F1: v6}
                                        _ = v30
                                    }
                                    _ = v26
                                } else {
                                    var v31 uint32 = v6
                                    v32 := GenS5{F0: false, F1: (v11 + (((v11 + v11) + (v11 + v11)) + v11)), F2: (v3 % ((32767) | 1))}
                                    _ = v31
                                    _ = v32
                                }
                                fmt.Println(v21)
                                v33 := GenS2{F0: 0, F1: 4294967295}
                                _ = v33
                            } else {
                                var v34 uint64 = 2
                                var v35 uint64 = 403
                                var v36 int = fi22
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = len(v2)
                        case "hello":
                            v24 = (v24 - v3)
                        }
                        _ = v24
                    } else {
                        if false {
                            v11 = ("" + v11)
                            fmt.Println(v21, fi23, v5)
                            v3 += fi23
                            v1[(fi22 & fi22)] = 731
                        }
                        var v37 int16 = 32767
                        v38 := GenS7{F0: (v4 >> ((2) & 63)), F1: (v5 >> ((0) & 63)), F2: v6}
                        _ = v37
                        _ = v38
                    }
                    switch fi23 {
                    case -589:
                        v39 := GenS5{F0: false, F1: "\n\t", F2: 208}
                        _ = v39
                    case 32767:
                        switch 2147483647 {
                        case 9223372036854775807:
                            fmt.Println(fi22, v6, v21)
                        case 513:
                            var v40 int = -472
                            _ = v40
                        case 42:
                            v5 = v5
                        default:
                            switch 2147483647 {
                            case -345:
                                v21 = v21
                            case 406:
                                v4 = 0
                            default:
                                if (0.0 < 0.0) {
                                    _ = len(v2)
                                    if false {
                                        if true {
                                            if false {
                                                v41 := GenS6{F0: 625, F1: (v6 + v6)}
                                                _ = v41
                                            } else {
                                                delete(v1, fi23)
                                            }
                                            var v42 int64 = 0
                                            var v43 float64 = 3.14159
                                            fmt.Println(v3, fi22)
                                            v44 := GenS7{F0: 100, F1: (^(v5)), F2: uint32(v5)}
                                            _ = v42
                                            _ = v43
                                            _ = v44
                                        }
                                    } else {
                                        var v45 uint = 4294967295
                                        var v46 string = (v11 + v11)
                                        v21 -= ((3.14159 + v21) - (-2.25 - 1.5))
                                        _ = v45
                                        _ = v46
                                    }
                                } else {
                                    switch fi23 {
                                    case 32767:
                                        switch fi23 {
                                        case -853:
                                            switch (-(fi22)) {
                                            case 9223372036854775807:
                                                if false {
                                                    switch (fi22 % ((1) | 1)) {
                                                    case 127:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    if (1 != genFn0(v6)) {
                                                        if len(v2) > 0 {
                                                            v2[0] = 2147483648
                                                        }
                                                        switch (-(fi22)) {
                                                        case -103:
                                                            var v47 bool = (!((v5 * 32767) != (v5 - (v5 >> ((0) & 63)))))
                                                            _ = v47
                                                        }
                                                    }
                                                } else {
                                                    var v48 uint64 = 2
                                                    var v49 uint8 = 1
                                                    if false {
                                                        v11 = ((((v11 + "") + "") + ("" + ("foo bar" + v11))) + (v11 + "foo bar"))
                                                        v5 += (v5 & 1)
                                                        var v50 int8 = (v4 >> ((1) & 63))
                                                        _ = v50
                                                    } else {
                                                        v51 := GenS8{F0: -1, F1: "", F2: -471}
                                                        v3 -= v3
                                                        _ = v51
                                                    }
                                                    _ = v48
                                                    _ = v49
                                                }
                                            }
                                        case 32767:
                                            v1[(fi22 - (fi22 | fi22))] = fi23
                                        case 9223372036854775807:
                                            v52 := GenS9{F0: fi22}
                                            _ = v52
                                        default:
                                            v11 = (((v11 + v11) + v11) + v11)
                                        }
                                    case 684:
                                        if false {
                                            v2 = append(v2, v6)
                                            v53 := GenS6{F0: 4294967295, F1: uint32(v5)}
                                            switch v11 {
                                            case "foo bar":
                                                _ = len(v1)
                                            case "\n\t":
                                                if (!(1 < 851)) {
                                                    v3 -= (fi23 + fi22)
                                                    if ((v9 << ((255) & 63)) != 1) {
                                                        var v54 int64 = 42
                                                        switch (v11 + v11) {
                                                        case "hello":
                                                            var v55 int = fi22
                                                            _ = v55
                                                        case "a":
                                                            v56 := GenS2{F0: -1, F1: 2}
                                                            _ = v56
                                                        case "":
                                                            var v57 int32 = -1
                                                            _ = v57
                                                        }
                                                        _ = v54
                                                    } else {
                                                        delete(v1, 32768)
                                                    }
                                                    v21 += float32(((3.14159 - 0.1) + 0.0))
                                                }
                                                if false {
                                                    var v58 int16 = 739
                                                    v3 += -378
                                                    var v59 int8 = (v4 * v4)
                                                    _ = v58
                                                    _ = v59
                                                }
                                            case "a":
                                                var v60 float64 = 1e10
                                                _ = v60
                                            }
                                            v1[42] = v3
                                            _ = v53
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v21, fi22)
                                    default:
                                        v4 = v4
                                    }
                                }
                            }
                        }
                    case -177:
                        v61 := GenS10{F0: 0, F1: (v6 | v6)}
                        _ = v61
                    }
                    var v62 uint8 = (^(v9))
                    v11 = (v11 + v11)
                    _ = fi23
                    _ = v62
                }
                switch (v3 >> ((18446744073709551615) & 63)) {
                case 32767:
                    if ((v6 << ((2) & 63)) > uint32(v5)) {
                        var v63 uint16 = (v5 * (v5 & v5))
                        v11 = v11
                        _ = v63
                    } else {
                        v4 = v4
                    }
                default:
                    v4 += v4
                }
                fmt.Println(v6, v4)
            } else {
                switch ("a" + "foo bar") {
                case "a":
                    switch -9223372036854775808 {
                    case 981:
                        var v64 uint = 4294967295
                        _ = v64
                    case -101:
                        v1[(fi22 ^ 1)] = (^(v3))
                    default:
                        for fi65 := 0; fi65 < 7; fi65++ {
                            if ((-32768 < 549) && (v6 <= 4294967295)) {
                                _ = len(v1)
                                v66 := GenS11{F0: 2, F1: ((42 >= -32768) && false)}
                                _ = v66
                            } else {
                                fmt.Println(v4, v11, fi22)
                            }
                            _ = fi65
                        }
                    }
                }
                var v67 uint64 = genFn0(v6)
                _ = v67
            }
        }
        _ = fi22
    }
    if true {
        v21 = float32(((3.14159 * ((-2.25 - 1e10) / 0.1)) * (-2.25 - (1e10 + float64((float32((-2.25 / (0.1 * 3.14159))) - -2.25))))))
    }
    v68 := GenS10{F0: v3, F1: v6}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if ((3.14159 * 0.1) < 1e10) {
        switch (v11 + "foo bar") {
        case "\n\t":
            var v69 uint8 = (v9 - v9)
            _ = v69
        default:
            fmt.Println(v21, v3)
        }
        for fi70 := 0; fi70 < 17; fi70++ {
            v71 := GenS4{F0: 0, F1: 1}
            _ = fi70
            _ = v71
        }
        var v72 uint32 = (v6 ^ v6)
        _ = v72
    } else {
        v21 = (v21 / v21)
        var v73 uint32 = (v6 / (((v6 & (v6 * v6))) | 1))
        _ = v73
    }
    for fi74 := 0; fi74 < 11; fi74++ {
        fmt.Println(v3)
        fmt.Println(v3, v11, v5)
        v11 = v11
        _ = fi74
    }
    for fi75 := 0; fi75 < 10; fi75++ {
        fmt.Println(v9, v4, v21)
        var v76 float64 = ((((-2.25 - float64(((v21 - v21) - float32((1.5 / (-2.25 / 0.1)))))) - float64(3.14159)) / float64(((-2.25 + (float32((3.14159 * 0.1)) + float32(float64(((3.14159 + 1e10) + v21))))) * (float32(((1e10 - 0.1) * 1e10)) * v21)))) + -2.25)
        v77 := GenS1{F0: (v11 + v11)}
        _ = fi75
        _ = v76
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v9
}
