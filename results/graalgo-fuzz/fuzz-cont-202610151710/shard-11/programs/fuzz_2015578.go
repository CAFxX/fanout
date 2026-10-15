package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int16; F2 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 uint16; F2 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint }
func genFn0(p0 uint64, p1 uint64, p2 int16) uint16 {
    p1 = p0
    var v7 int32 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 bool = true
    v2 := []uint32{674, 2147483648, 898}
    var v3 int = 4294967295
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = (v4 | (v4 ^ v4))
    var v6 uint8 = uint8(v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v8 := GenS1{F0: 0, F1: 519}
    fmt.Println(v1, v1)
    var v9 uint64 = (v4 & (v5 | v4))
    var v10 float64 = (3.14159 + -2.25)
    var v11 uint16 = 845
    var v12 uint64 = 18446744073709551615
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v10, v1)
    case "foo bar":
        v3 -= int(v9)
    case "hello":
        if true {
            var v13 uint16 = v11
            switch (v3 / (((^(v3))) | 1)) {
            case 128:
                switch 42 {
                case -1:
                    if (!false) {
                        v1 = false
                        var v14 string = ("foo bar" + (("a" + "\n\t") + ("a" + ("hello" + (("" + ("\n\t" + (("hello" + "") + "a"))) + "\n\t")))))
                        _ = v14
                    }
                case 42:
                    var v15 float64 = v10
                    _ = v15
                default:
                    v16 := GenS2{F0: 2147483648}
                    _ = v16
                }
            case 127:
                if (18446744073709551615 <= 404) {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v1 = true
                        var v18 int32 = 1
                        v12 -= 42
                        _ = fi17
                        _ = v18
                    }
                    switch v3 {
                    case 2147483647:
                        var v19 bool = (2 == 42)
                        _ = v19
                    case 0:
                        if (!(-815 > 1)) {
                            v3 = -236
                            var v20 uint64 = (^(v12))
                            if true {
                                fmt.Println(v4, v11, v11)
                                v21 := GenS3{F0: 0, F1: 100}
                                _ = v21
                            }
                            _ = v20
                        }
                    case 143:
                        fmt.Println(v13, v12)
                    default:
                        var v22 int8 = -1
                        _ = v22
                    }
                } else {
                    var v23 float64 = (0.0 * (float64(float32(v10)) - v10))
                    _ = v23
                }
            case -9223372036854775808:
                for fi24 := 0; fi24 < 9; fi24++ {
                    v4 += (v9 / ((2) | 1))
                    _ = fi24
                }
            }
            _ = v13
        } else {
            fmt.Println(v1, v11)
            if (!v1) {
                fmt.Println(v9, v3, v5)
                v2 = append(v2, 1)
                v25 := GenS2{F0: 4294967295}
                _ = v25
            }
            if (((v1 || (18446744073709551615 > (v12 << ((2) & 63)))) && v1) || true) {
                switch ("foo bar" + "") {
                case "a":
                    var v26 int32 = int32(v6)
                    _ = v26
                case "hello":
                    v6 = v6
                case "":
                    for fi27 := 0; fi27 < 18; fi27++ {
                        var v28 float32 = ((0.1 + (-2.25 * (0.1 * -2.25))) * (1e10 * 0.1))
                        var v29 int8 = 100
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v1 = (((0.0 + 0.0) / -2.25) == 1e10)
                            fmt.Println(v9)
                            v2 = append(v2, 42)
                            _ = fi30
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    v31 := GenS4{F0: 127}
                    _ = v31
                }
                v11 = (v11 - 353)
                if false {
                    if v1 {
                        if v1 {
                            if ((v6 <= 31) && (-1 != -21)) {
                                var v32 uint32 = 65535
                                _ = v32
                            } else {
                                var v33 uint16 = (v11 - v11)
                                fmt.Println(v3, v33, v10)
                                v33 += genFn0((v9 - (^(v4))), 2, 1000)
                                _ = v33
                            }
                        } else {
                            var v34 uint8 = v6
                            var v35 uint16 = (v11 * v11)
                            v36 := GenS3{F0: 65535, F1: -1}
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = len(v2)
                        if (v1 && (v1 && (-24 >= 127))) {
                            var v37 int32 = -1
                            if (803 >= 954) {
                                v1 = (v11 < genFn0((v5 & v9), (v9 / ((v12) | 1)), int16(v3)))
                                if false {
                                    v2 = append(v2, 306)
                                } else {
                                    var v38 float64 = v10
                                    var v39 string = "\n\t"
                                    _ = v38
                                    _ = v39
                                }
                                var v40 int8 = 100
                                _ = v40
                            } else {
                                switch (v3 | (^(v3))) {
                                case -36:
                                    v12 = v4
                                }
                            }
                            _ = v37
                        }
                    } else {
                        v3 = (v3 * v3)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi41 := 0; fi41 < 7; fi41++ {
                            v42 := GenS1{F0: 0, F1: genFn0(18446744073709551615, 162, 383)}
                            _ = fi41
                            _ = v42
                        }
                    }
                    _ = len(v2)
                    var v43 int8 = 100
                    _ = v43
                } else {
                    var v44 uint64 = (^(v12))
                    v12 += 1
                    switch ("\n\t" + "foo bar") {
                    case "":
                        var v45 int16 = -257
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 9; fi46++ {
                            _ = len(v2)
                            _ = fi46
                        }
                    }
                    _ = v44
                }
            } else {
                for fi47 := 0; fi47 < 17; fi47++ {
                    for fi48 := 0; fi48 < 9; fi48++ {
                        fmt.Println(v3, fi48, v3)
                        var v49 float32 = -2.25
                        v2 = append(v2, 0)
                        _ = fi48
                        _ = v49
                    }
                    v2 = append(v2, 0)
                    if (!(!v1)) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi47
                }
            }
        }
    }
    v3 = -677
    var v50 int64 = 9223372036854775807
    if ((int16(v5) < 42) && ((v6 | 162) <= (v6 << ((0) & 63)))) {
        v6 = (v6 >> ((424) & 63))
        v51 := GenS5{F0: uint64(v11), F1: 1000, F2: 366}
        if len(v2) > 0 {
            v2[0] = uint32(v11)
        }
        _ = v51
    } else {
        var v52 float32 = (0.1 / 1e10)
        v53 := GenS6{F0: v6}
        _ = v52
        _ = v53
    }
    _ = len(v2)
    fmt.Println(v1, v1)
    for fi54 := 0; fi54 < 15; fi54++ {
        v55 := GenS7{F0: -671, F1: genFn0(v12, v9, -32768), F2: v1}
        switch "hello" {
        case "":
            _ = len(v2)
        default:
            v56 := GenS8{F0: 117}
            _ = v56
        }
        if len(v2) > 0 {
            v2[0] = 42
        }
        _ = fi54
        _ = v55
    }
    for fi57 := 0; fi57 < 4; fi57++ {
        v10 += 0.0
        for fi58 := 0; fi58 < 9; fi58++ {
            if (!v1) {
                switch "a" {
                case "foo bar":
                    fmt.Println(v4)
                default:
                    switch ("\n\t" + "hello") {
                    case "":
                        var v59 float32 = 3.14159
                        v60 := GenS2{F0: 2147483648}
                        _ = v59
                        _ = v60
                    }
                }
                if false {
                    switch fi57 {
                    case 525:
                        v9 += v12
                    case 1:
                        var v61 uint = 2
                        _ = v61
                    case 2147483647:
                        v62 := GenS1{F0: 127, F1: 959}
                        _ = v62
                    default:
                        if (float64(0.0) > v10) {
                            v11 = (^(v11))
                            fmt.Println(fi57)
                        } else {
                            v63 := GenS6{F0: v6}
                            _ = v63
                        }
                    }
                }
                if len(v2) > 0 {
                    v2[0] = 42
                }
            } else {
                if (!false) {
                    var v64 float32 = (((0.0 + 3.14159) / (float32((v10 + float64((0.0 / 3.14159)))) + float32(v10))) / ((0.1 / (-2.25 - 1.5)) + -2.25))
                    _ = v64
                }
            }
            if v1 {
                v6 = (v6 ^ (v6 | (v6 >> ((0) & 63))))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                switch "foo bar" {
                case "":
                    v2 = append(v2, 2)
                case "a":
                    fmt.Println(v3, fi58)
                case "hello":
                    v4 = (v5 * (v5 - (^(v12))))
                    if (840 <= -1) {
                        fmt.Println(v6)
                        var v65 int16 = 88
                        var v66 string = "hello"
                        var v67 float32 = 1e10
                        _ = v65
                        _ = v66
                        _ = v67
                    } else {
                        var v68 bool = true
                        v50 += -145
                        v10 = (-2.25 - v10)
                        _ = v68
                    }
                    _ = len(v2)
                default:
                    var v69 float64 = v10
                    _ = v69
                }
                if len(v2) > 0 {
                    v2[0] = 71
                }
                switch "\n\t" {
                case "\n\t":
                    var v70 uint = 4294967295
                    _ = v70
                case "hello":
                    var v71 float32 = (1e10 + 0.1)
                    _ = v71
                }
                if len(v2) > 0 {
                    v2[0] = 232
                }
            } else {
                v4 = v5
            }
            var v72 string = ""
            v5 = v9
            _ = fi58
            _ = v72
        }
        v10 = 0.0
        _ = fi57
    }
    fmt.Println(v4, v5)
    for fi73 := 0; fi73 < 4; fi73++ {
        if v1 {
            if (v6 >= 128) {
                var v74 int8 = 0
                var v75 uint8 = v6
                v12 = v12
                _ = v74
                _ = v75
            }
            if (2 == (v6 * v6)) {
                if ((v11 != (v11 ^ v11)) || (0.1 != float32((float64((float32((v10 * (0.0 - ((v10 * 0.1) / v10)))) * 3.14159)) + 0.1)))) {
                    var v76 uint8 = 2
                    fmt.Println(v10)
                    switch v3 {
                    case 32767:
                        v5 += v5
                    case 219:
                        for fi77 := 0; fi77 < 6; fi77++ {
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            switch "hello" {
                            case "hello":
                                switch "foo bar" {
                                case "hello":
                                    v78 := GenS6{F0: v6}
                                    _ = v78
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        switch "hello" {
                                        case "foo bar":
                                            var v79 uint8 = (v6 << ((393) & 63))
                                            _ = v79
                                        case "hello":
                                            fmt.Println(v11)
                                        case "":
                                            v2 = append(v2, 5)
                                        }
                                        switch "\n\t" {
                                        case "hello":
                                            if (1000 != 42) {
                                                v50 = (-(v50))
                                            } else {
                                                var v80 string = "foo bar"
                                                if (!false) {
                                                    fmt.Println(v9, fi73, v3)
                                                } else {
                                                    var v81 uint8 = v76
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = 1
                                            }
                                        case "\n\t":
                                            var v82 uint = 255
                                            _ = v82
                                        case "":
                                            v5 = v12
                                        default:
                                            v83 := GenS7{F0: 42, F1: 412, F2: v1}
                                            _ = v83
                                        }
                                        v12 += (^(v9))
                                    case "a":
                                        if len(v2) > 0 {
                                            v2[0] = 160
                                        }
                                    case "":
                                        fmt.Println(v6, v3, fi73)
                                    }
                                }
                            case "a":
                                v6 = 2
                            case "foo bar":
                                v84 := GenS6{F0: v76}
                                _ = v84
                            }
                            v2 = append(v2, 42)
                            _ = fi77
                        }
                    default:
                        v85 := GenS5{F0: v9, F1: 0, F2: 4294967295}
                        _ = v85
                    }
                    _ = v76
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v4 += (^(v9))
                for fi86 := 0; fi86 < 15; fi86++ {
                    v1 = v1
                    switch 32767 {
                    case 1:
                        v87 := GenS6{F0: (v6 - v6)}
                        _ = v87
                    case 127:
                        if len(v2) > 0 {
                            v2[0] = 241
                        }
                    }
                    fmt.Println(v5)
                    _ = fi86
                }
            } else {
                var v88 float32 = -2.25
                _ = v88
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v89 := GenS9{F0: 4294967295}
        _ = fi73
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v8
    _ = v9
}
