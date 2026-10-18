package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint64; F2 uint }
func genFn0(p0 int64, p1 int8) uint8 {
    var v7 uint32 = 2
    fmt.Println(v7, v7, v7)
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(p0, p1, fi8)
        switch ("\n\t" + "foo bar") {
        case "":
            var v9 string = "\n\t"
            _ = v9
        default:
            v7 = 42
        }
        switch ("\n\t" + "\n\t") {
        case "a":
            if false {
                var v10 int8 = p1
                v11 := GenS1{F0: -46, F1: 438}
                if (!false) {
                    v12 := []int{fi8, (fi8 ^ fi8), (fi8 >> ((4294967295) & 63))}
                    _ = len(v12)
                    if ((^(p1)) != v10) {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            if (101 >= 255) {
                                v7 = (v7 + (v7 ^ v7))
                                var v14 float32 = (1e10 - (-2.25 * 0.0))
                                v15 := GenS1{F0: v10, F1: 42}
                                _ = v14
                                _ = v15
                            } else {
                                p0 -= p0
                                var v16 int32 = -458
                                _ = v16
                            }
                            _ = fi13
                        }
                        if len(v12) > 0 {
                            v12[0] = (fi8 | fi8)
                        }
                        if ((1 != -2147483648) && false) {
                            for fi17 := 0; fi17 < 1; fi17++ {
                                p1 = p1
                                var v18 int32 = 2147483647
                                if (!true) {
                                    v7 += 2147483648
                                }
                                _ = fi17
                                _ = v18
                            }
                            if true {
                                fmt.Println(p0, v10, v7)
                            } else {
                                v19 := GenS1{F0: int8(v7), F1: 923}
                                _ = v19
                            }
                        }
                    }
                    _ = v12
                }
                _ = v10
                _ = v11
            } else {
                for fi20 := 0; fi20 < 18; fi20++ {
                    if ((!(248 == 2)) || (!true)) {
                        v21 := []int{(fi8 + fi8)}
                        _ = v21
                    }
                    p1 = (-(p1))
                    if (false || (false || true)) {
                        var v22 int64 = -9223372036854775808
                        fmt.Println(fi20)
                        if (!true) {
                            var v23 uint32 = (v7 & (^(v7)))
                            _ = v23
                        } else {
                            var v24 uint = 42
                            v25 := GenS1{F0: (p1 % ((p1) | 1)), F1: 758}
                            _ = v24
                            _ = v25
                        }
                        p0 = p0
                        v26 := GenS2{F0: -32768}
                        v27 := []int{(^(fi20))}
                        _ = v22
                        _ = v26
                        _ = v27
                    }
                    _ = fi20
                }
            }
        case "":
            if (1 > 1) {
                if true {
                    for fi28 := 0; fi28 < 15; fi28++ {
                        if true {
                            if (!true) {
                                v7 = (v7 >> ((40) & 63))
                                v7 -= v7
                                v29 := []int{}
                                _ = v29
                            }
                        }
                        v30 := GenS2{F0: 0}
                        var v31 int16 = -32768
                        _ = fi28
                        _ = v30
                        _ = v31
                    }
                    v32 := []int{(-(fi8)), (fi8 >> ((255) & 63)), fi8}
                    _ = v32
                } else {
                    switch "hello" {
                    case "":
                        p1 -= (p1 ^ (p1 - p1))
                    case "foo bar":
                        switch ((("hello" + ("" + (("foo bar" + ("hello" + "a")) + ""))) + (("a" + "hello") + ("foo bar" + "a"))) + "") {
                        case "foo bar":
                            switch "hello" {
                            case "hello":
                                p1 += int8(fi8)
                            case "a":
                                for fi33 := 0; fi33 < 19; fi33++ {
                                    var v34 float32 = 1.5
                                    _ = fi33
                                    _ = v34
                                }
                            case "\n\t":
                                fmt.Println(p1)
                            }
                        case "a":
                            for fi35 := 0; fi35 < 14; fi35++ {
                                p0 -= p0
                                _ = fi35
                            }
                        case "":
                            for fi36 := 0; fi36 < 18; fi36++ {
                                p1 = (p1 * 1)
                                p0 -= p0
                                _ = fi36
                            }
                        }
                    case "\n\t":
                        switch fi8 {
                        case 785:
                            v37 := GenS3{F0: -813, F1: 805, F2: v7}
                            _ = v37
                        case 32767:
                            v38 := GenS3{F0: (p0 - (-(p0))), F1: 0, F2: v7}
                            _ = v38
                        }
                    default:
                        var v39 int64 = 9223372036854775807
                        _ = v39
                    }
                    fmt.Println(p1, v7)
                    p0 -= (p0 * p0)
                }
                p1 = -128
            } else {
                var v40 int = (fi8 << ((275) & 63))
                _ = v40
            }
        case "hello":
            if (true || false) {
                var v41 float32 = (((3.14159 - float32(0.1)) + (float32(-2.25) * -2.25)) / 1.5)
                _ = v41
            } else {
                v42 := GenS4{F0: 255, F1: true, F2: p0}
                _ = v42
            }
        }
        _ = fi8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = v7
    return 211
}
func genFn1() int64 {
    var v43 bool = (((((!(0 >= -9223372036854775808)) && false) && false) && (true || true)) || ((18446744073709551615 > 4294967295) || true))
    var v44 int64 = -1
    var v45 float64 = float64(0.0)
    switch "a" {
    case "a":
        switch ("hello" + ("hello" + ((("foo bar" + (("foo bar" + "\n\t") + "hello")) + "hello") + (("" + ("foo bar" + "\n\t")) + "\n\t")))) {
        case "":
            v44 -= -9223372036854775808
        case "a":
            for fi46 := 0; fi46 < 9; fi46++ {
                fmt.Println(v43, v45, fi46)
                _ = fi46
            }
        }
    default:
        if (42 >= 2) {
            v47 := []int{128}
            _ = v47
        } else {
            v48 := GenS1{F0: 42, F1: 1}
            v49 := GenS2{F0: 0}
            _ = v48
            _ = v49
        }
    }
    _ = v43
    _ = v44
    _ = v45
    return 42
}

func main() {
    var v1 uint8 = 46
    var v2 uint32 = 2147483647
    var v3 uint16 = 108
    var v4 bool = false
    var v5 string = "a"
    var v6 uint8 = 1
    var v50 uint32 = (v2 - v2)
    var v51 int8 = 68
    var v52 float32 = (0.0 - -2.25)
    v2 = (v2 * v50)
    if true {
        if (-2147483648 <= -611) {
            for fi53 := 0; fi53 < 10; fi53++ {
                v2 = v2
                v54 := GenS2{F0: 502}
                _ = fi53
                _ = v54
            }
        }
        v55 := GenS5{F0: v51}
        _ = v55
    } else {
        var v56 int16 = -251
        v57 := []int{}
        for fi58 := 0; fi58 < 15; fi58++ {
            switch 976 {
            case -9223372036854775808:
                switch fi58 {
                case 127:
                    var v59 bool = v4
                    _ = v59
                default:
                    for fi60 := 0; fi60 < 19; fi60++ {
                        switch fi60 {
                        case 4294967295:
                            switch (fi58 & 32767) {
                            case -503:
                                if len(v57) > 0 {
                                    v57[0] = fi60
                                }
                            case 9223372036854775807:
                                v57 = append(v57, (fi58 % ((fi60) | 1)))
                            default:
                                v61 := GenS3{F0: -9223372036854775808, F1: (v3 << ((65535) & 63)), F2: (v2 << ((622) & 63))}
                                _ = v61
                            }
                        }
                        v4 = (false && v4)
                        v62 := GenS3{F0: 0, F1: 0, F2: v2}
                        var v63 uint = 18446744073709551615
                        _ = fi60
                        _ = v62
                        _ = v63
                    }
                }
            case 128:
                fmt.Println(v6)
            case 42:
                if (!((((((v5 + v5) + v5) + "hello") < v5) && v4) && false)) {
                    switch -2147483648 {
                    case 128:
                        var v64 float64 = 1e10
                        _ = v64
                    case 42:
                        var v65 uint32 = v2
                        _ = v65
                    case 169:
                        fmt.Println(fi58, v4)
                    default:
                        fmt.Println(v51, v2)
                    }
                    if len(v57) >= 2 {
                        v57 = v57[1:len(v57)]
                    }
                    var v66 int8 = v51
                    _ = v66
                }
            }
            if (!v4) {
                if v4 {
                    var v67 bool = false
                    _ = v67
                } else {
                    for fi68 := 0; fi68 < 4; fi68++ {
                        v57 = append(v57, 4294967295)
                        if len(v57) >= 2 {
                            v57 = v57[1:len(v57)]
                        }
                        if len(v57) >= 2 {
                            v57 = v57[1:len(v57)]
                        }
                        _ = fi68
                    }
                }
                if (148 == 18446744073709551615) {
                    var v69 uint64 = 74
                    _ = v69
                }
                var v70 float64 = (0.0 - (1e10 - 0.0))
                _ = v70
            }
            _ = fi58
        }
        _ = v56
        _ = v57
    }
    v71 := GenS5{F0: (v51 << ((18446744073709551615) & 63))}
    var v72 uint32 = 1
    var v73 int32 = -5
    v74 := []int{42, 32767, 9223372036854775807, 1}
    if true {
        fmt.Println(v2, v50)
    }
    var v75 int16 = 32767
    v76 := GenS1{F0: (v51 << ((2) & 63)), F1: 1}
    var v77 int64 = int64(v3)
    if (4294967295 < 32768) {
        var v78 int8 = int8(v75)
        if v4 {
            for fi79 := 0; fi79 < 8; fi79++ {
                v1 += genFn0((^(v77)), (v78 / (((v51 % (((-(v51))) | 1))) | 1)))
                _ = len(v74)
                _ = fi79
            }
        }
        _ = v78
    } else {
        if v4 {
            v80 := GenS2{F0: 0}
            switch "" {
            case "a":
                var v81 int16 = v75
                _ = v81
            case "":
                v72 = (^(v2))
            }
            _ = v80
        } else {
            v1 = (v6 / ((genFn0((v77 % (((v77 & (-(v77)))) | 1)), -29)) | 1))
            v74 = append(v74, -358)
        }
        if false {
            var v82 string = v5
            if v4 {
                var v83 string = v5
                _ = v83
            }
            _ = v82
        }
    }
    switch -2147483648 {
    case -216:
        for fi84 := 0; fi84 < 10; fi84++ {
            for fi85 := 0; fi85 < 18; fi85++ {
                fmt.Println(v6)
                switch (v5 + ("hello" + v5)) {
                case "a":
                    var v86 int = 4294967295
                    _ = v86
                case "\n\t":
                    if (18446744073709551615 >= 42) {
                        var v87 int32 = (v73 * v73)
                        v88 := GenS2{F0: (v75 * (v75 * int16(fi84)))}
                        _ = v87
                        _ = v88
                    }
                case "foo bar":
                    fmt.Println(v1, v5, v4)
                default:
                    v74 = append(v74, int(v3))
                }
                _ = fi85
            }
            v74 = append(v74, fi84)
            for fi89 := 0; fi89 < 14; fi89++ {
                fmt.Println(v73, v6, v6)
                var v90 float32 = (v52 - 3.14159)
                if len(v74) >= 2 {
                    v74 = v74[1:len(v74)]
                }
                fmt.Println(v1, v72)
                var v91 uint = uint(v72)
                v92 := GenS1{F0: (v51 | v51), F1: 753}
                _ = fi89
                _ = v90
                _ = v91
                _ = v92
            }
            _ = fi84
        }
    }
    v93 := GenS6{F0: (v3 - (v3 | v3)), F1: 0, F2: 2}
    var v94 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("slicelen v74:", len(v74), v74)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v93
    _ = v94
}
