package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64; F1 uint64 }
type GenS11 struct { F0 uint; F1 bool }
type GenS12 struct { F0 uint8; F1 uint16 }
type GenS13 struct { F0 uint64; F1 uint }
func genFn0() uint {
    var v6 uint16 = 255
    fmt.Println(v6)
    _ = v6
    return 255
}
func genFn1(p0 uint64, p1 uint32) uint16 {
    var v7 int = 2147483647
    v8 := []int{673, (^(v7)), (v7 / ((v7) | 1)), (v7 << ((899) & 63))}
    for fi9 := 0; fi9 < 1; fi9++ {
        _ = len(v8)
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int = 655
            v12 := GenS1{F0: v7, F1: (fi9 >> ((4294967295) & 63))}
            var v13 uint16 = uint16(p1)
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    var v14 uint64 = 845
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    _ = v8
    return 659
}

func main() {
    v1 := []int{32767, 331, 768}
    v2 := make(map[uint32]string)
    v3 := []string{}
    v4 := []string{("\n\t" + ""), "hello"}
    var v5 bool = (1 == 272)
    var v15 uint = 100
    if true {
        switch "hello" {
        case "foo bar":
            switch ("" + ((("\n\t" + "foo bar") + "foo bar") + "a")) {
            case "hello":
                v16 := GenS2{F0: ((!true) && false), F1: ""}
                _ = v16
            default:
                v17 := GenS1{F0: 128, F1: -1}
                _ = v17
            }
        }
        if (!false) {
            var v18 int32 = -2147483648
            v18 -= (v18 / ((-361) | 1))
            v19 := GenS3{F0: 467}
            _ = v18
            _ = v19
        } else {
            for fi20 := 0; fi20 < 5; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    if (v5 && true) {
                        v5 = (65535 > 333)
                        v22 := GenS1{F0: fi21, F1: (fi20 * 0)}
                        _ = v22
                    } else {
                        var v23 int = (-(fi21))
                        _ = v23
                    }
                    v24 := GenS3{F0: 2147483648}
                    var v25 uint16 = 1
                    _ = fi21
                    _ = v24
                    _ = v25
                }
                switch ("hello" + "a") {
                case "foo bar":
                    var v26 uint32 = 432
                    _ = v26
                case "\n\t":
                    if v5 {
                        switch (fi20 / ((0) | 1)) {
                        case -2147483648:
                            fmt.Println(fi20)
                        case 968:
                            for fi27 := 0; fi27 < 9; fi27++ {
                                switch (^(fi27)) {
                                case 2147483647:
                                    v28 := GenS3{F0: 233}
                                    _ = v28
                                case 498:
                                    fmt.Println(fi27, v5)
                                case -1:
                                    if v5 {
                                        v3 = append(v3, ("a" + "a"))
                                        if v5 {
                                            if (892 >= 42) {
                                                _ = len(v2)
                                                if (2147483648 != 42) {
                                                    v29 := GenS2{F0: true, F1: "foo bar"}
                                                    _ = v29
                                                }
                                                v5 = (!((genFn0() > (v15 / ((v15) | 1))) || v5))
                                            } else {
                                                var v30 uint8 = 0
                                                v5 = (fi27 != fi20)
                                                _ = v30
                                            }
                                            v5 = v5
                                        }
                                    } else {
                                        if v5 {
                                            if v5 {
                                                fmt.Println(fi20, fi20, fi27)
                                            } else {
                                                var v31 bool = false
                                                v15 = 0
                                                fmt.Println(v31, v31)
                                                _ = v31
                                            }
                                        } else {
                                            if ((-2.25 - 0.0) < 1e10) {
                                                switch (((("\n\t" + "foo bar") + "hello") + "a") + (("hello" + ("" + ("hello" + ("\n\t" + "hello")))) + "hello")) {
                                                case "":
                                                    var v32 int64 = 9223372036854775807
                                                    _ = v32
                                                case "foo bar":
                                                    v5 = (-32768 > 549)
                                                }
                                                if v5 {
                                                    if (false && v5) {
                                                        var v33 int32 = -1
                                                        _ = v33
                                                    }
                                                }
                                            } else {
                                                v34 := GenS4{F0: fi20, F1: 1, F2: -651}
                                                v35 := GenS5{F0: (^(v15))}
                                                if ((v15 & v15) == v15) {
                                                    fmt.Println(v5, fi20, fi27)
                                                    var v36 uint32 = 645
                                                    _ = v36
                                                } else {
                                                    v37 := GenS1{F0: -1, F1: (fi27 + -289)}
                                                    if v5 {
                                                        switch ("foo bar" + "foo bar") {
                                                        case "a":
                                                            var v38 uint16 = 65535
                                                            _ = v38
                                                        default:
                                                            var v39 float32 = ((((0.1 / (float32((3.14159 + 0.0)) + (1.5 + ((3.14159 - 0.0) + (0.0 - ((0.0 - 1e10) * 3.14159)))))) * 0.1) - -2.25) - ((3.14159 - 1.5) * (0.0 - (3.14159 / (1e10 + ((1e10 / -2.25) + float32(0.1)))))))
                                                            _ = v39
                                                        }
                                                        if v5 {
                                                            delete(v2, 0)
                                                            if false {
                                                                switch "a" {
                                                                case "\n\t":
                                                                    v15 -= 2
                                                                default:
                                                                    var v40 string = "a"
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                switch fi27 {
                                                                case 128:
                                                                    fmt.Println(v15, fi20)
                                                                    v15 -= (v15 + genFn0())
                                                                    fmt.Println(fi27)
                                                                }
                                                            }
                                                            if v5 {
                                                                var v41 float32 = (-2.25 + 0.1)
                                                                v15 = v15
                                                                _ = v41
                                                            }
                                                        }
                                                        v42 := GenS3{F0: 4294967295}
                                                        var v43 uint32 = 2147483648
                                                        switch 0 {
                                                        case 1:
                                                            v44 := GenS5{F0: genFn0()}
                                                            _ = v44
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        var v45 bool = false
                                                        if (v45 && true) {
                                                            fmt.Println(v45, v45)
                                                        }
                                                        if (!true) {
                                                            fmt.Println(v45, v15, v45)
                                                            var v46 int32 = -951
                                                            v15 += genFn0()
                                                            _ = v46
                                                        }
                                                        _ = v45
                                                    }
                                                    fmt.Println(fi27)
                                                    _ = v37
                                                }
                                                v47 := GenS6{F0: 127}
                                                _ = v34
                                                _ = v35
                                                _ = v47
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                        }
                                    }
                                default:
                                    v5 = (2 >= 1)
                                }
                                var v48 float32 = 1.5
                                _ = fi27
                                _ = v48
                            }
                        }
                        switch "hello" {
                        case "":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            for fi49 := 0; fi49 < 16; fi49++ {
                                v50 := GenS2{F0: (1 > 42), F1: ("" + "foo bar")}
                                v51 := GenS5{F0: v15}
                                _ = fi49
                                _ = v50
                                _ = v51
                            }
                        }
                        v52 := GenS1{F0: (fi20 >> ((genFn0()) & 63)), F1: (fi20 / (((fi20 % ((1) | 1))) | 1))}
                        _ = v52
                    }
                case "hello":
                    var v53 string = "hello"
                    _ = v53
                default:
                    var v54 uint16 = genFn1(582, 6)
                    _ = v54
                }
                var v55 int16 = 400
                _ = fi20
                _ = v55
            }
        }
        var v56 uint64 = 1
        _ = v56
    } else {
        v5 = v5
        for fi57 := 0; fi57 < 6; fi57++ {
            if v5 {
                v15 = 0
            }
            _ = fi57
        }
        fmt.Println(v5, v15)
    }
    for fi58 := 0; fi58 < 8; fi58++ {
        v59 := GenS5{F0: (v15 << ((v15) & 63))}
        var v60 bool = v5
        if (v60 && (v60 && v5)) {
            v61 := GenS7{F0: 18446744073709551615}
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            v1 = append(v1, (fi58 ^ (fi58 & 4294967295)))
            _ = v61
        } else {
            if v5 {
                if ("foo bar" != ((("" + "hello") + "") + ("hello" + ("a" + ("" + "hello"))))) {
                    for fi62 := 0; fi62 < 11; fi62++ {
                        var v63 int64 = 0
                        var v64 string = "a"
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                    switch ("\n\t" + "") {
                    case "hello":
                        fmt.Println(fi58)
                    default:
                        if (!v60) {
                            v5 = ((127 > 2) || v60)
                            var v65 uint16 = 2
                            _ = v65
                        }
                    }
                    switch 363 {
                    case 391:
                        v66 := GenS8{F0: "a"}
                        _ = v66
                    case -1:
                        for fi67 := 0; fi67 < 2; fi67++ {
                            fmt.Println(v15, v15, fi58)
                            v60 = false
                            _ = fi67
                        }
                    }
                } else {
                    for fi68 := 0; fi68 < 20; fi68++ {
                        fmt.Println(fi58, fi68)
                        if (v5 || false) {
                            var v69 uint64 = 0
                            v70 := GenS9{F0: (!false)}
                            v71 := GenS6{F0: 0}
                            _ = v69
                            _ = v70
                            _ = v71
                        } else {
                            var v72 int = fi58
                            if ((1.5 >= 0.0) && v60) {
                                var v73 uint64 = 18446744073709551615
                                _ = v73
                            }
                            var v74 int16 = 1
                            _ = v72
                            _ = v74
                        }
                        v75 := GenS10{F0: 0, F1: 18446744073709551615}
                        var v76 uint = genFn0()
                        _ = fi68
                        _ = v75
                        _ = v76
                    }
                }
                if (-1 <= 82) {
                    fmt.Println(fi58)
                }
                v15 -= v15
            }
        }
        _ = fi58
        _ = v59
        _ = v60
    }
    v2[25] = "\n\t"
    v77 := GenS11{F0: (v15 << (((v15 * (v15 << (((v15 << (((v15 / ((65535) | 1))) & 63))) & 63)))) & 63)), F1: v5}
    if true {
        for fi78 := 0; fi78 < 9; fi78++ {
            var v79 float64 = 1.5
            _ = fi78
            _ = v79
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(v15, v5)
        case "hello":
            v80 := GenS1{F0: -2147483648, F1: 128}
            _ = v80
        default:
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v15, v5)
            case "":
                var v81 uint8 = 138
                _ = v81
            default:
                v82 := GenS3{F0: 2147483647}
                _ = v82
            }
        }
    }
    v15 += v15
    for fi83 := 0; fi83 < 9; fi83++ {
        v4 = append(v4, "")
        _ = fi83
    }
    var v84 float64 = -2.25
    v85 := GenS5{F0: (v15 + (v15 << ((255) & 63)))}
    v86 := v2[2147483647]
    fmt.Println(v84)
    if (-1 >= 0) {
        v5 = ((float32(v84) / float32((v84 / v84))) != 0.0)
    } else {
        if len(v4) > 0 {
            v4[0] = v86
        }
        var v87 string = "a"
        v86 = v87
        _ = v87
    }
    var v88 uint16 = 976
    fmt.Println(v84)
    switch v86 {
    case "hello":
        if v5 {
            if false {
                var v89 uint8 = 205
                if v5 {
                    v86 = v86
                } else {
                    v2[2147483648] = (v86 + v86)
                    if (-2147483648 == 0) {
                        var v90 string = ("foo bar" + "")
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        v89 += (v89 / ((128) | 1))
                        _ = v90
                    }
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                }
                _ = v89
            }
            if ((v5 || v5) || (v84 == v84)) {
                v91 := GenS12{F0: 0, F1: v88}
                switch (v86 + v86) {
                case "a":
                    if (v5 && (((v84 * v84) - (v84 / v84)) != 1e10)) {
                        fmt.Println(v5, v86)
                        v3 = append(v3, "a")
                    } else {
                        fmt.Println(v86)
                        var v92 bool = (!(!((v15 % ((1) | 1)) == v15)))
                        v5 = (v92 || false)
                        _ = v92
                    }
                case "hello":
                    v5 = false
                }
                var v93 string = ("hello" + v86)
                _ = v91
                _ = v93
            }
            fmt.Println(v5)
        }
    case "foo bar":
        if v5 {
            for fi94 := 0; fi94 < 11; fi94++ {
                v88 = genFn1(42, 381)
                v95 := GenS2{F0: v5, F1: ("a" + v86)}
                for fi96 := 0; fi96 < 2; fi96++ {
                    v1 = append(v1, (-(fi94)))
                    _ = fi96
                }
                _ = fi94
                _ = v95
            }
        }
    }
    if v5 {
        var v97 uint64 = 4294967295
        var v98 uint64 = (^(v97))
        _ = v97
        _ = v98
    }
    v99 := GenS13{F0: 4294967295, F1: genFn0()}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v86:", v86)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v84
    _ = v85
    _ = v86
    _ = v88
    _ = v99
}
