package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 string; F2 int8 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 string; F1 uint8; F2 int32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 int8, p2 uint16) uint8 {
    var v6 int8 = (^(p1))
    switch "foo bar" {
    case "a":
        p0 = 249
    case "foo bar":
        p2 = 311
    case "\n\t":
        fmt.Println(p0, p2, p1)
    }
    var v7 int8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 8
}
func genFn1(p0 int8) uint64 {
    v8 := GenS1{F0: 1000, F1: 131, F2: 1}
    if (4294967295 < 42) {
        p0 = 42
        if true {
            p0 = p0
            p0 = 0
        } else {
            fmt.Println(p0)
            for fi9 := 0; fi9 < 7; fi9++ {
                if true {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v11 := GenS2{F0: 2}
                        var v12 int64 = 1
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    if false {
                        p0 -= p0
                        var v13 string = "a"
                        _ = v13
                    } else {
                        fmt.Println(fi9, fi9)
                        fmt.Println(p0)
                    }
                    v14 := GenS3{F0: 0, F1: "", F2: p0}
                    _ = v14
                }
                v15 := GenS2{F0: 2}
                _ = fi9
                _ = v15
            }
        }
        var v16 int64 = 0
        _ = v16
    } else {
        v17 := GenS2{F0: 65535}
        var v18 uint32 = 2
        _ = v17
        _ = v18
    }
    switch 132 {
    case -390:
        fmt.Println(p0)
    }
    v19 := GenS3{F0: 0, F1: (("a" + "") + ((("foo bar" + "a") + "") + ("foo bar" + "hello"))), F2: 1}
    _ = p0
    _ = v19
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{127}
    var v2 float32 = 1.5
    v3 := make(map[string]int)
    var v4 bool = true
    var v5 int8 = 127
    fmt.Println(v4, v4)
    v2 = ((v2 / v2) * 0.1)
    v2 = v2
    var v20 uint8 = genFn0(32767, (^(v5)), 2)
    fmt.Println(v5)
    v20 = (v20 | 42)
    for fi21 := 0; fi21 < 4; fi21++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!v4) {
            var v22 string = ("\n\t" + "foo bar")
            var v23 int32 = 2147483647
            _ = v22
            _ = v23
        }
        if (true && (v4 && (v4 && ((!false) && true)))) {
            var v24 int64 = -9223372036854775808
            if (1 <= 2) {
                switch (fi21 - fi21) {
                case 81:
                    if v4 {
                        fmt.Println(v20)
                    } else {
                        switch "" {
                        case "hello":
                            v25 := GenS2{F0: 28}
                            _ = v25
                        case "foo bar":
                            v2 = 3.14159
                        case "\n\t":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            switch fi21 {
                            case -9223372036854775808:
                                var v26 uint64 = 346
                                _ = v26
                            case 1:
                                var v27 float32 = (((0.1 / v2) + float32(0.0)) + v2)
                                _ = v27
                            case 2147483647:
                                switch (fi21 << ((1) & 63)) {
                                case 42:
                                    v5 += -4
                                case 128:
                                    for fi28 := 0; fi28 < 11; fi28++ {
                                        v5 = 0
                                        if (v4 || (!v4)) {
                                            var v29 uint16 = 1
                                            v2 = ((v2 / 1.0) - 3.14159)
                                            _ = v29
                                        }
                                        fmt.Println(fi21, v2)
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v30 uint16 = 992
                                            _ = v30
                                        default:
                                            var v31 float32 = (0.0 * v2)
                                            _ = v31
                                        }
                                        _ = fi28
                                    }
                                case -718:
                                    for fi32 := 0; fi32 < 13; fi32++ {
                                        var v33 float64 = (-2.25 - (3.14159 - -2.25))
                                        _ = fi32
                                        _ = v33
                                    }
                                default:
                                    var v34 uint64 = 1
                                    _ = v34
                                }
                            }
                        }
                    }
                case 737:
                    switch (fi21 * -9223372036854775808) {
                    case -9223372036854775808:
                        if v4 {
                            switch "" {
                            case "foo bar":
                                for fi35 := 0; fi35 < 4; fi35++ {
                                    fmt.Println(fi35, v20, v5)
                                    v5 += v5
                                    if v4 {
                                        fmt.Println(v2)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v36 := GenS4{F0: (fi21 >> ((1) & 63)), F1: 2}
                                        v20 = (v20 - 153)
                                        _ = v36
                                    }
                                    _ = fi35
                                }
                            default:
                                switch (fi21 + (^(fi21))) {
                                case 232:
                                    if true {
                                        var v37 float32 = 1e10
                                        v24 -= (v24 ^ v24)
                                        var v38 int = 1
                                        _ = v37
                                        _ = v38
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = fi21
                                        }
                                        for fi39 := 0; fi39 < 8; fi39++ {
                                            var v40 uint32 = 551
                                            v24 += v24
                                            var v41 string = ("" + "")
                                            v42 := GenS5{F0: 0, F1: 1}
                                            if false {
                                                if false {
                                                    if false {
                                                        v4 = (((float32(1e10) + (((1e10 + v2) + v2) - float32(3.14159))) + (((3.14159 - 1.5) / 1.0) - ((3.14159 - 1.5) / (v2 / v2)))) > ((v2 - v2) / (float32((1.5 + 0.1)) / v2)))
                                                        v43 := GenS6{F0: ("hello" + v41), F1: v20, F2: int32(v40)}
                                                        if (fi21 == (fi21 & fi39)) {
                                                            var v44 string = v41
                                                            v24 = (v24 ^ (v24 >> ((255) & 63)))
                                                            switch int(v24) {
                                                            case -21:
                                                                switch "hello" {
                                                                case "hello":
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                default:
                                                                    var v45 int32 = 0
                                                                    _ = v45
                                                                }
                                                            case 549:
                                                                v46 := GenS4{F0: fi39, F1: 65535}
                                                                _ = v46
                                                            case 4294967295:
                                                                v40 -= v40
                                                                v24 += (v24 | -559)
                                                                var v47 int16 = 1000
                                                                _ = v47
                                                            default:
                                                                v48 := GenS7{F0: -103}
                                                                _ = v48
                                                            }
                                                            _ = v44
                                                        } else {
                                                            var v49 uint32 = (v40 + (v40 << ((642) & 63)))
                                                            _ = v49
                                                        }
                                                        var v50 uint = 65535
                                                        v50 -= v50
                                                        var v51 bool = v4
                                                        _ = v43
                                                        _ = v50
                                                        _ = v51
                                                    }
                                                    v41 = (v41 + v41)
                                                    v52 := v3["a"]
                                                    _ = v52
                                                }
                                            }
                                            _ = fi39
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                        }
                                        for fi53 := 0; fi53 < 18; fi53++ {
                                            var v54 int = (fi21 * 1)
                                            _ = fi53
                                            _ = v54
                                        }
                                    }
                                case -884:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 127:
                                    v3["\n\t"] = (fi21 - (fi21 / ((fi21) | 1)))
                                default:
                                    v2 = float32(1e10)
                                }
                            }
                            var v55 uint8 = v20
                            var v56 uint16 = 255
                            _ = v55
                            _ = v56
                        }
                    case 128:
                        var v57 int64 = v24
                        _ = v57
                    case 562:
                        var v58 uint8 = (v20 + (^(v20)))
                        _ = v58
                    default:
                        var v59 int64 = (v24 << ((0) & 63))
                        _ = v59
                    }
                default:
                    for fi60 := 0; fi60 < 6; fi60++ {
                        delete(v3, "")
                        _ = fi60
                    }
                }
            }
            v2 = (-2.25 - ((v2 + float32(1e10)) - (v2 * float32(3.14159))))
            _ = v24
        }
        _ = fi21
    }
    if (v4 && v4) {
        if (!(1.5 > 1.5)) {
            for fi61 := 0; fi61 < 9; fi61++ {
                v62 := GenS1{F0: int16(fi61), F1: 589, F2: 65535}
                _ = fi61
                _ = v62
            }
            fmt.Println(v2)
        } else {
            var v63 uint = 147
            var v64 float32 = float32(3.14159)
            if ((int64(v5) >= 42) || true) {
                fmt.Println(v64, v2)
            } else {
                if ((0 >= 42) || (false && v4)) {
                    v63 = uint(v20)
                } else {
                    v65 := GenS5{F0: 42, F1: 0}
                    fmt.Println(v20)
                    _ = v65
                }
                v2 = v2
            }
            _ = v63
            _ = v64
        }
    } else {
        var v66 string = "\n\t"
        _ = v66
    }
    var v67 string = ("hello" + (("hello" + "hello") + ((("a" + "foo bar") + (("hello" + "\n\t") + "hello")) + "\n\t")))
    v5 += -1
    if (v2 < v2) {
        for fi68 := 0; fi68 < 10; fi68++ {
            for fi69 := 0; fi69 < 10; fi69++ {
                switch (^(fi68)) {
                case -9223372036854775808:
                    if true {
                        switch (-(fi69)) {
                        case 0:
                            if (!(((v20 >> ((4294967295) & 63)) > (^(v20))) && v4)) {
                                if v4 {
                                    var v70 int8 = -128
                                    fmt.Println(v70)
                                    v20 = 200
                                    if true {
                                        if len(v1) > 0 {
                                            v1[0] = fi68
                                        }
                                        fmt.Println(fi68, v2)
                                        fmt.Println(v20, v4)
                                    }
                                    switch -9223372036854775808 {
                                    case 32767:
                                        switch fi69 {
                                        case 610:
                                            if (v70 >= (-(v70))) {
                                                var v71 int8 = 82
                                                var v72 int8 = -1
                                                if false {
                                                    v20 = genFn0(0, 1, 973)
                                                    switch (fi69 + fi68) {
                                                    case -2147483648:
                                                        v71 += (v72 & (v71 % (((v72 - v71)) | 1)))
                                                    case -885:
                                                        switch fi68 {
                                                        case -1:
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            if (0 != int16(v72)) {
                                                                v2 += float32(1e10)
                                                            } else {
                                                                v73 := GenS1{F0: 42, F1: 109, F2: 65535}
                                                                v70 = (-(v5))
                                                                _ = v73
                                                            }
                                                        }
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    v74 := GenS1{F0: -1, F1: -609, F2: 1}
                                                    v5 += v70
                                                    v4 = v4
                                                    switch (-(fi69)) {
                                                    case 242:
                                                        var v75 bool = v4
                                                        _ = v75
                                                    case 605:
                                                        _ = len(v1)
                                                    case 128:
                                                        var v76 uint16 = 1
                                                        _ = v76
                                                    }
                                                    _ = v74
                                                }
                                                _ = v71
                                                _ = v72
                                            }
                                        case -9223372036854775808:
                                            var v77 bool = v4
                                            _ = v77
                                        }
                                    default:
                                        fmt.Println(fi68, v2)
                                    }
                                    v4 = true
                                    _ = v70
                                }
                            }
                        case 1:
                            fmt.Println(v67, v67, v5)
                        default:
                            var v78 string = v67
                            _ = v78
                        }
                        _ = len(v1)
                    }
                case 32767:
                    v79 := GenS2{F0: 659}
                    _ = v79
                default:
                    if (fi68 > (-(fi69))) {
                        var v80 bool = (((!false) && false) && false)
                        _ = v80
                    }
                }
                _ = fi69
            }
            v67 = ("\n\t" + v67)
            _ = fi68
        }
        v20 += genFn0(1, (v5 | (v5 % (((v5 * v5)) | 1))), 272)
        var v81 uint = 526
        _ = v81
    } else {
        var v82 uint64 = 1
        v83 := GenS8{F0: 0, F1: v82, F2: 2147483647}
        _ = v82
        _ = v83
    }
    v84 := v3[v67]
    v5 -= -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v67
    _ = v84
}
