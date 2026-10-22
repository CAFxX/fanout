package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int16; F1 int16; F2 int64 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 bool }
func genFn0(p0 uint32, p1 uint8) uint16 {
    p0 -= (p0 | p0)
    v7 := GenS1{F0: 100, F1: -445, F2: p1}
    switch 42 {
    case 128:
        if ((p1 % ((p1) | 1)) > 171) {
            switch 9223372036854775807 {
            case 0:
                for fi8 := 0; fi8 < 5; fi8++ {
                    var v9 int8 = 104
                    p0 = p0
                    v10 := GenS1{F0: v9, F1: (fi8 >> ((22) & 63)), F2: (^(p1))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case 1:
                v11 := GenS1{F0: 42, F1: 127, F2: (p1 >> ((886) & 63))}
                _ = v11
            default:
                v12 := GenS1{F0: 0, F1: -90, F2: p1}
                _ = v12
            }
            p1 += 127
            if ("a" > (("hello" + "\n\t") + "foo bar")) {
                p1 = 182
            } else {
                v13 := []int{-355, 4294967295}
                fmt.Println(p1, p0)
                p0 -= 90
                _ = v13
            }
        }
    }
    v14 := GenS2{F0: (^(p0))}
    var v15 uint32 = (p0 << ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v7
    return 32767
}
func genFn1(p0 uint16, p1 uint) uint16 {
    var v16 int8 = -128
    switch 4294967295 {
    case 4294967295:
        if false {
            v17 := GenS1{F0: v16, F1: 4294967295, F2: 127}
            _ = v17
        }
    default:
        v18 := GenS3{F0: v16, F1: v16, F2: 852}
        _ = v18
    }
    v19 := GenS2{F0: 496}
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v16
    _ = v19
    return genFn0(217, 42)
}

func main() {
    var v1 string = (("foo bar" + "\n\t") + ("a" + ""))
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 bool = true
    v5 := []uint32{4294967295, 695, 342}
    v6 := make(map[string]string)
    var v20 uint64 = 2
    delete(v3, 705)
    if ((("foo bar" + v1) >= (v1 + v1)) && (!(1.5 <= -2.25))) {
        v21 := GenS3{F0: 0, F1: 100, F2: 2147483647}
        v5 = append(v5, 277)
        var v22 int32 = 812
        _ = v21
        _ = v22
    }
    var v23 uint = 255
    v6[v1] = "a"
    if (100 < 24) {
        delete(v3, 2147483647)
        var v24 uint64 = 18446744073709551615
        _ = v24
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v25 := GenS1{F0: 100, F1: 32767, F2: 42}
        _ = v25
    }
    var v26 uint32 = 4294967295
    var v27 uint64 = v20
    v4 = true
    var v28 int16 = 266
    v20 -= v27
    v29 := GenS4{F0: v28, F1: 468}
    for fi30 := 0; fi30 < 11; fi30++ {
        if ((1 != uint64(v26)) || ((!((1e10 * (((float64(0.0) / 1.0) - 0.1) + 0.1)) > 3.14159)) && v2)) {
            if len(v5) > 0 {
                v5[0] = v26
            }
            if (2 > 128) {
                _ = len(v6)
                v31 := GenS3{F0: 100, F1: 42, F2: 412}
                _ = v31
            } else {
                v2 = (-1 <= -1)
                if (!(((v28 / ((-627) | 1)) <= v28) || v4)) {
                    for fi32 := 0; fi32 < 12; fi32++ {
                        v33 := GenS5{F0: v28, F1: (v28 << (((v23 << (((v23 << ((0) & 63))) & 63))) & 63)), F2: 9223372036854775807}
                        var v34 uint8 = 127
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    switch (v1 + "hello") {
                    case "a":
                        fmt.Println(v2, v27)
                    default:
                        switch v1 {
                        case "hello":
                            v35 := GenS6{F0: 42, F1: (v2 || (-852 >= 2147483647))}
                            _ = v35
                        case "\n\t":
                            var v36 float32 = float32(0.1)
                            _ = v36
                        }
                    }
                    if (v2 || (484 > 32767)) {
                        v26 -= v26
                        fmt.Println(v20)
                    } else {
                        if true {
                            if (1 == 1) {
                                var v37 uint = v23
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v37
                            } else {
                                var v38 float64 = 3.14159
                                v39 := GenS1{F0: -128, F1: int(v26), F2: 127}
                                v40 := GenS3{F0: 42, F1: -128, F2: 65535}
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            fmt.Println(v28)
                            switch (((v1 + (("foo bar" + v1) + ("foo bar" + ("\n\t" + ((v1 + v1) + ((("foo bar" + v1) + "hello") + v1)))))) + "hello") + "\n\t") {
                            case "hello":
                                _ = len(v3)
                            }
                        } else {
                            var v41 uint16 = 255
                            switch "foo bar" {
                            case "\n\t":
                                switch fi30 {
                                case 32768:
                                    var v42 int16 = (v28 - v28)
                                    _ = v42
                                case 0:
                                    if len(v5) > 0 {
                                        v5[0] = v26
                                    }
                                default:
                                    switch "" {
                                    case "":
                                        var v43 uint = uint(v26)
                                        _ = v43
                                    case "foo bar":
                                        switch (fi30 + fi30) {
                                        case 1:
                                            for fi44 := 0; fi44 < 17; fi44++ {
                                                if v2 {
                                                    fmt.Println(v27)
                                                    var v45 float64 = 0.0
                                                    v46 := GenS7{F0: v27, F1: 0}
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    if (!(127 >= 106)) {
                                                        var v47 uint32 = (v26 & (^(v26)))
                                                        var v48 int64 = 42
                                                        v41 += genFn0((v47 * 1), 127)
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                    v27 += 4294967295
                                                }
                                                if false {
                                                    if (!(3.14159 >= 1.5)) {
                                                        fmt.Println(v20)
                                                        switch fi44 {
                                                        case 32767:
                                                            v5 = append(v5, v26)
                                                        case -1:
                                                            v49 := GenS2{F0: (v26 / ((2147483648) | 1))}
                                                            _ = v49
                                                        case -2147483648:
                                                            if len(v5) > 0 {
                                                                v5[0] = 65535
                                                            }
                                                        default:
                                                            v41 = uint16(v26)
                                                        }
                                                        var v50 float64 = (0.0 + 3.14159)
                                                        _ = v50
                                                    } else {
                                                        var v51 uint64 = uint64(v26)
                                                        v41 -= 2
                                                        _ = v51
                                                    }
                                                    var v52 float64 = (((0.1 - (0.1 * (0.0 - (-2.25 + -2.25)))) * 1e10) / 0.1)
                                                    var v53 float64 = v52
                                                    switch (v1 + (v1 + v1)) {
                                                    case "a":
                                                        var v54 int64 = 9223372036854775807
                                                        _ = v54
                                                    }
                                                    v26 -= v26
                                                    v20 = (v27 >> ((v23) & 63))
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    switch (fi44 ^ fi44) {
                                                    case 4294967295:
                                                        if v2 {
                                                            delete(v3, 65535)
                                                            if true {
                                                                v1 = ("\n\t" + (("a" + "a") + ""))
                                                            }
                                                            fmt.Println(fi44, v41)
                                                        } else {
                                                            if v4 {
                                                                v4 = (!(1.5 < 3.14159))
                                                                var v55 int16 = v28
                                                                _ = v55
                                                            } else {
                                                                v27 = v27
                                                                v56 := GenS8{F0: v26, F1: v26}
                                                                var v57 bool = false
                                                                if (!(67 > 127)) {
                                                                    v41 = v41
                                                                    var v58 float64 = (1e10 - 1e10)
                                                                    var v59 bool = true
                                                                    _ = v58
                                                                    _ = v59
                                                                } else {
                                                                    v28 = -865
                                                                    switch (("foo bar" + (v1 + "hello")) + (((("hello" + ("\n\t" + (v1 + (v1 + (v1 + "\n\t"))))) + v1) + v1) + "a")) {
                                                                    case "":
                                                                        fmt.Println(v27, v57)
                                                                    case "foo bar":
                                                                        v2 = v57
                                                                    case "hello":
                                                                        v23 += v23
                                                                        var v60 int16 = v28
                                                                        if (0.1 >= 3.14159) {
                                                                            var v61 string = v1
                                                                            var v62 bool = true
                                                                            var v63 int32 = -1
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        _ = v60
                                                                    default:
                                                                        fmt.Println(v23)
                                                                    }
                                                                }
                                                                var v64 int64 = 42
                                                                _ = v56
                                                                _ = v57
                                                                _ = v64
                                                            }
                                                            var v65 int64 = 1
                                                            var v66 int = (fi30 ^ fi30)
                                                            _ = v65
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                                if (v28 == v28) {
                                                    v4 = (-72 <= 127)
                                                }
                                                _ = fi44
                                            }
                                        case 2147483647:
                                            var v67 int64 = 42
                                            _ = v67
                                        case 692:
                                            fmt.Println(v1)
                                        }
                                    default:
                                        for fi68 := 0; fi68 < 4; fi68++ {
                                            v69 := v6[(v1 + (v1 + "hello"))]
                                            fmt.Println(v27)
                                            var v70 uint = 2
                                            _ = fi68
                                            _ = v69
                                            _ = v70
                                        }
                                    }
                                }
                            default:
                                if v2 {
                                    _ = len(v6)
                                    if v2 {
                                        v71 := GenS5{F0: v28, F1: int16(v41), F2: 9223372036854775807}
                                        _ = v71
                                    } else {
                                        var v72 uint64 = (v27 << ((v23) & 63))
                                        v73 := v6[v1]
                                        _ = v72
                                        _ = v73
                                    }
                                    if true {
                                        for fi74 := 0; fi74 < 3; fi74++ {
                                            var v75 int64 = 106
                                            v2 = true
                                            switch "a" {
                                            case "foo bar":
                                                var v76 string = "foo bar"
                                                _ = v76
                                            }
                                            _ = fi74
                                            _ = v75
                                        }
                                        for fi77 := 0; fi77 < 6; fi77++ {
                                            v78 := GenS9{F0: 42}
                                            var v79 int = (fi30 ^ fi77)
                                            _ = fi77
                                            _ = v78
                                            _ = v79
                                        }
                                    }
                                }
                            }
                            _ = v41
                        }
                    }
                }
            }
            fmt.Println(fi30, v28)
        }
        v80 := GenS10{F0: v2}
        for fi81 := 0; fi81 < 6; fi81++ {
            v28 += v28
            _ = fi81
        }
        _ = fi30
        _ = v80
    }
    for fi82 := 0; fi82 < 7; fi82++ {
        v83 := GenS5{F0: v28, F1: (v28 >> ((uint(v27)) & 63)), F2: -890}
        for fi84 := 0; fi84 < 1; fi84++ {
            if false {
                var v85 uint32 = 2
                var v86 int16 = (v28 + v28)
                _ = v85
                _ = v86
            } else {
                v26 = v26
                var v87 int64 = -426
                fmt.Println(fi84, v28, v23)
                _ = v87
            }
            v2 = v2
            var v88 int = fi82
            _ = fi84
            _ = v88
        }
        _ = fi82
        _ = v83
    }
    var v89 string = "hello"
    if (-2.25 > 0.0) {
        if (!(v27 != (v20 >> ((748) & 63)))) {
            fmt.Println(v23)
            switch ("a" + ("\n\t" + (("foo bar" + "a") + v1))) {
            case "foo bar":
                for fi90 := 0; fi90 < 14; fi90++ {
                    fmt.Println(v4, v23)
                    v91 := GenS8{F0: (v26 - (^(v26))), F1: 4294967295}
                    _ = fi90
                    _ = v91
                }
            default:
                v1 = "hello"
            }
            switch "foo bar" {
            case "\n\t":
                v1 = (((v1 + v89) + v89) + (v89 + v89))
            case "":
                delete(v3, v26)
            case "foo bar":
                for fi92 := 0; fi92 < 10; fi92++ {
                    if len(v5) > 0 {
                        v5[0] = (v26 + v26)
                    }
                    var v93 int32 = 2147483647
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi92
                    _ = v93
                }
            }
        }
        if true {
            var v94 uint32 = (^(v26))
            _ = v94
        }
    }
    v27 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
}
