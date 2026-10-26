package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8; F1 int16; F2 int16 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 uint16; F2 bool }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 int32 }
func genFn0(p0 int32) bool {
    var v5 uint8 = 2
    v5 += (v5 & (v5 << ((18446744073709551615) & 63)))
    var v6 float64 = 1e10
    _ = p0
    _ = v5
    _ = v6
    return (127 < -91)
}
func genFn1(p0 string, p1 uint, p2 uint16) uint32 {
    p2 += 42
    var v7 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int16 = 0
    var v2 uint16 = 32767
    v3 := make(map[string]uint32)
    var v4 bool = false
    var v8 int16 = v1
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch "hello" {
            case "foo bar":
                var v11 uint = 0
                _ = v11
            case "\n\t":
                v12 := GenS1{F0: 213}
                _ = v12
            case "":
                v13 := GenS1{F0: 0}
                _ = v13
            }
            fmt.Println(v1)
            _ = fi10
        }
        _ = fi9
    }
    if v4 {
        switch "a" {
        case "":
            v3["foo bar"] = 4294967295
        case "foo bar":
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 int = fi14
                switch "a" {
                case "foo bar":
                    v16 := GenS2{F0: 0}
                    _ = v16
                case "\n\t":
                    var v17 uint16 = (v2 | 623)
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        case "\n\t":
            switch ("" + "foo bar") {
            case "\n\t":
                var v18 string = "\n\t"
                _ = v18
            case "a":
                var v19 bool = (42 > 1)
                _ = v19
            }
        }
        switch 222 {
        case 128:
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 uint64 = 30
                for fi22 := 0; fi22 < 1; fi22++ {
                    if ((fi22 - fi20) != (fi22 << ((2) & 63))) {
                        v23 := GenS3{F0: fi20, F1: -32768}
                        v24 := []int{-2147483648, (fi22 ^ fi20)}
                        switch "a" {
                        case "":
                            switch "foo bar" {
                            case "hello":
                                switch "hello" {
                                case "foo bar":
                                    v2 = v2
                                case "hello":
                                    switch fi22 {
                                    case -1:
                                        var v25 bool = false
                                        _ = v25
                                    case 488:
                                        v8 -= v8
                                    }
                                default:
                                    if (42 >= -2147483648) {
                                        delete(v3, ("" + ("hello" + "hello")))
                                    } else {
                                        if (("" + "hello") == "\n\t") {
                                            fmt.Println(v2, v1, v1)
                                            fmt.Println(v4)
                                            var v26 uint = 302
                                            _ = v26
                                        } else {
                                            var v27 uint8 = 128
                                            _ = v27
                                        }
                                        v21 -= (^(v21))
                                        if (((float32(0.1) + 1.5) * 0.0) < -2.25) {
                                            v28 := GenS4{F0: int16(fi20)}
                                            var v29 int32 = 524
                                            var v30 int16 = v8
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        } else {
                                            var v31 string = "a"
                                            _ = v31
                                        }
                                    }
                                }
                            case "\n\t":
                                switch (fi20 & -9223372036854775808) {
                                case -1:
                                    v2 += 0
                                }
                            case "foo bar":
                                if (1 < v21) {
                                    var v32 uint8 = 42
                                    v32 = (v32 + (v32 & 255))
                                    _ = v32
                                }
                            default:
                                if len(v24) > 0 {
                                    v24[0] = (fi20 ^ fi20)
                                }
                            }
                        case "\n\t":
                            var v33 uint = 42
                            _ = v33
                        case "foo bar":
                            var v34 uint64 = (v21 - 2)
                            _ = v34
                        }
                        _ = v23
                        _ = v24
                    }
                    var v35 int64 = 0
                    v21 = 0
                    v36 := GenS5{F0: 128, F1: 715, F2: v8}
                    var v37 uint16 = (v2 | v2)
                    _ = fi22
                    _ = v35
                    _ = v36
                    _ = v37
                }
                _ = fi20
                _ = v21
            }
        case 0:
            v1 = -1
        case 4294967295:
            var v38 uint16 = 32767
            _ = v38
        }
        switch (("\n\t" + "foo bar") + "hello") {
        case "a":
            var v39 uint = 1
            _ = v39
        case "\n\t":
            switch 9223372036854775807 {
            case 32767:
                for fi40 := 0; fi40 < 11; fi40++ {
                    if v4 {
                        var v41 int64 = -9223372036854775808
                        for fi42 := 0; fi42 < 15; fi42++ {
                            v43 := v3["foo bar"]
                            var v44 uint8 = 1
                            v45 := GenS2{F0: v1}
                            if (v4 && false) {
                                fmt.Println(fi42, fi40)
                            } else {
                                v46 := GenS3{F0: (^(fi42)), F1: v8}
                                if (v4 && (0.0 == (float32((3.14159 + (float64(3.14159) + 1e10))) / 1.5))) {
                                    switch (fi40 ^ fi42) {
                                    case 128:
                                        v47 := []int{(fi40 ^ 32768)}
                                        _ = v47
                                    case -1:
                                        v8 += v1
                                    case 0:
                                        var v48 int = (fi40 << ((65535) & 63))
                                        _ = v48
                                    default:
                                        fmt.Println(v43, v1)
                                    }
                                    var v49 int = -9223372036854775808
                                    v4 = (!v4)
                                    _ = v49
                                } else {
                                    fmt.Println(fi40)
                                }
                                v41 = v41
                                _ = v46
                            }
                            _ = fi42
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        v50 := v3[(("" + "") + (((("\n\t" + "hello") + ("foo bar" + ("" + "hello"))) + "a") + "\n\t"))]
                        _ = v41
                        _ = v50
                    }
                    switch (fi40 | (fi40 - fi40)) {
                    case 4294967295:
                        var v51 bool = (!v4)
                        _ = v51
                    case 32768:
                        fmt.Println(v4, v4)
                    case -9223372036854775808:
                        switch "a" {
                        case "a":
                            v52 := GenS3{F0: fi40, F1: v8}
                            _ = v52
                        case "foo bar":
                            v2 = 42
                        case "":
                            if v4 {
                                var v53 int64 = 1
                                v53 += v53
                                var v54 string = (("foo bar" + "foo bar") + "hello")
                                _ = v53
                                _ = v54
                            } else {
                                var v55 bool = false
                                _ = v55
                            }
                        }
                    default:
                        v56 := GenS6{F0: v8}
                        _ = v56
                    }
                    switch "\n\t" {
                    case "":
                        v1 = 1
                    case "hello":
                        var v57 int8 = 99
                        _ = v57
                    default:
                        fmt.Println(v8, fi40, v1)
                    }
                    _ = fi40
                }
            case -2147483648:
                if false {
                    for fi58 := 0; fi58 < 18; fi58++ {
                        var v59 int32 = 2147483647
                        fmt.Println(v2, v4, v4)
                        switch ("\n\t" + "") {
                        case "foo bar":
                            for fi60 := 0; fi60 < 16; fi60++ {
                                if (!v4) {
                                    v61 := GenS7{F0: (v2 ^ v2), F1: (v2 | 42), F2: (1 < 616)}
                                    _ = v61
                                } else {
                                    fmt.Println(v8)
                                }
                                v1 = v1
                                if v4 {
                                    fmt.Println(v8)
                                }
                                fmt.Println(fi58, v4)
                                _ = fi60
                            }
                        case "hello":
                            if true {
                                v62 := GenS8{F0: -9223372036854775808, F1: 255}
                                _ = v62
                            }
                        case "a":
                            v4 = ((!true) && v4)
                        default:
                            for fi63 := 0; fi63 < 3; fi63++ {
                                v64 := []int{fi63, -436, -122}
                                v8 -= (v8 * v8)
                                _ = fi63
                                _ = v64
                            }
                        }
                        _ = fi58
                        _ = v59
                    }
                    if (((1.5 + 1e10) * ((-2.25 + ((0.0 + -2.25) + (-2.25 * float32(3.14159)))) / (1e10 * 3.14159))) != (1e10 * 1e10)) {
                        v65 := []int{-1}
                        v66 := GenS7{F0: (v2 * 42), F1: 32767, F2: v4}
                        _ = v65
                        _ = v66
                    }
                    v67 := GenS9{F0: -87}
                    _ = v67
                }
            }
        case "hello":
            var v68 float32 = ((0.1 * (((float32(3.14159) / float32(1.5)) - 3.14159) / (1e10 - -2.25))) + ((0.1 / 3.14159) + 0.0))
            _ = v68
        default:
            var v69 uint16 = (v2 + (v2 >> ((255) & 63)))
            _ = v69
        }
    }
    fmt.Println(v8, v2)
    fmt.Println(v2, v1, v1)
    var v70 uint64 = 42
    fmt.Println(v70, v70, v1)
    for fi71 := 0; fi71 < 3; fi71++ {
        var v72 string = ("\n\t" + "\n\t")
        v73 := GenS6{F0: (v1 << ((118) & 63))}
        _ = fi71
        _ = v72
        _ = v73
    }
    v4 = (!(0.0 < -2.25))
    var v74 int32 = 0
    if (1 == 42) {
        switch "" {
        case "foo bar":
            v70 += (^(v70))
        }
    } else {
        var v75 uint8 = 1
        _ = v75
    }
    v76 := v3[""]
    v3[""] = 604
    if (v76 >= (v76 / ((v76) | 1))) {
        v1 += (^(v1))
    }
    if false {
        for fi77 := 0; fi77 < 6; fi77++ {
            v78 := GenS8{F0: 1, F1: 754}
            var v79 uint = 0
            _ = fi77
            _ = v78
            _ = v79
        }
        v2 -= 255
    }
    var v80 uint8 = 0
    v81 := []int{-1, 882}
    for fi82 := 0; fi82 < 8; fi82++ {
        var v83 uint = 2
        _ = fi82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    fmt.Println("v76:", v76)
    fmt.Println("v8:", v8)
    fmt.Println("v80:", v80)
    fmt.Println("slicelen v81:", len(v81), v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v74
    _ = v76
    _ = v8
    _ = v80
    _ = v81
}
