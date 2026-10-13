package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int8; F1 int16 }
type GenS5 struct { F0 bool; F1 uint16; F2 int32 }
type GenS6 struct { F0 int; F1 uint8; F2 int8 }
type GenS7 struct { F0 bool; F1 uint16; F2 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int, p1 int64) string {
    v7 := []int{p0, (p0 / (((p0 - p0)) | 1)), p0}
    v8 := GenS1{F0: 127, F1: 2147483647, F2: p0}
    var v9 int = (^(p0))
    v10 := GenS1{F0: 42, F1: 914, F2: (p0 ^ v9)}
    var v11 bool = (!true)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}
func genFn1(p0 int64) int8 {
    p0 += -9223372036854775808
    var v12 uint8 = 255
    v13 := []int{-629, -1}
    for fi14 := 0; fi14 < 13; fi14++ {
        v12 = 128
        v15 := GenS1{F0: 75, F1: -414, F2: (fi14 | 9223372036854775807)}
        v16 := GenS2{F0: 2}
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!((19 != 2147483648) || false)) {
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int8 = -1
                fmt.Println(fi17)
                _ = fi18
                _ = v19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                if (-1 > 0) {
                    p0 = (p0 >> ((131) & 63))
                    _ = len(v13)
                    switch fi20 {
                    case 127:
                        v21 := GenS3{F0: v12, F1: 127}
                        _ = v21
                    }
                } else {
                    var v22 uint = 18446744073709551615
                    _ = v22
                }
                p0 = (p0 - p0)
                v23 := GenS4{F0: 127, F1: 32767}
                _ = fi20
                _ = v23
            }
            _ = fi17
        }
        fmt.Println(p0, v12)
        v12 = 255
    } else {
        v13 = append(v13, -218)
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    }
    _ = p0
    _ = v12
    _ = v13
    return 100
}

func main() {
    var v1 uint8 = 127
    var v2 uint8 = 2
    v3 := make(map[int]int)
    var v4 int8 = 1
    v5 := make(map[uint32]bool)
    v6 := make(map[int]bool)
    for fi24 := 0; fi24 < 4; fi24++ {
        if (42 != 32767) {
            var v25 float64 = (0.1 * 0.0)
            fmt.Println(v4, v4, v1)
            _ = v25
        } else {
            if (((1e10 * 3.14159) + (1e10 - 1.5)) < 1.5) {
                _ = len(v5)
                fmt.Println(v2, v2, v4)
                fmt.Println(fi24)
            } else {
                v26 := GenS3{F0: 29, F1: (fi24 % ((9223372036854775807) | 1))}
                for fi27 := 0; fi27 < 18; fi27++ {
                    v28 := GenS1{F0: v1, F1: 0, F2: 32767}
                    v29 := v6[411]
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                switch 42 {
                case 1:
                    switch "hello" {
                    case "foo bar":
                        v4 = (v4 - genFn1(0))
                    default:
                        v30 := v3[fi24]
                        _ = v30
                    }
                }
                _ = v26
            }
            v4 = (v4 / ((1) | 1))
        }
        fmt.Println(v2, v1)
        v2 = (v1 / (((v1 & (v2 * (v2 / (((v2 - (v2 / (((v1 / (((v2 % (((^(v1))) | 1))) | 1))) | 1)))) | 1))))) | 1))
        _ = fi24
    }
    if true {
        fmt.Println(v2, v1, v2)
    } else {
        v1 = 255
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        v32 := GenS2{F0: v2}
        v33 := GenS5{F0: true, F1: 42, F2: 894}
        _ = fi31
        _ = v32
        _ = v33
    }
    switch "hello" {
    case "hello":
        var v34 int16 = -32768
        _ = v34
    }
    v35 := []int{128, 0, -9223372036854775808, 231}
    switch 127 {
    case 4294967295:
        switch 128 {
        case 4294967295:
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 float64 = 1e10
                var v38 string = "a"
                _ = fi36
                _ = v37
                _ = v38
            }
        case -9223372036854775808:
            v39 := GenS6{F0: 711, F1: (v1 << ((4294967295) & 63)), F2: (v4 % ((genFn1(9223372036854775807)) | 1))}
            _ = v39
        default:
            for fi40 := 0; fi40 < 19; fi40++ {
                var v41 uint8 = (v1 + (^(v2)))
                var v42 uint16 = 702
                _ = fi40
                _ = v41
                _ = v42
            }
        }
    case 478:
        v1 = (v1 * v2)
    default:
        var v43 int64 = int64(v4)
        _ = v43
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        v1 = v2
        _ = fi44
    }
    switch "a" {
    case "foo bar":
        for fi45 := 0; fi45 < 9; fi45++ {
            v2 = 42
            _ = fi45
        }
    case "hello":
        var v46 bool = false
        _ = v46
    default:
        var v47 uint = 656
        _ = v47
    }
    v48 := GenS6{F0: -631, F1: 127, F2: genFn1(0)}
    if false {
        switch 0 {
        case 4294967295:
            if (2147483647 <= 2147483648) {
                v2 -= (^(v1))
                var v49 uint32 = 42
                _ = v49
            } else {
                for fi50 := 0; fi50 < 14; fi50++ {
                    var v51 int = (fi50 ^ fi50)
                    v1 = v2
                    switch -930 {
                    case 32768:
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        case "hello":
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                        }
                    case -152:
                        fmt.Println(fi50, v1, fi50)
                    }
                    _ = fi50
                    _ = v51
                }
            }
        case -9223372036854775808:
            var v52 uint32 = 245
            _ = v52
        case -829:
            if (127 < 0) {
                if (4294967295 >= 42) {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            if (!false) {
                                switch (fi54 & int(v2)) {
                                case 4294967295:
                                    var v55 string = "foo bar"
                                    _ = v55
                                default:
                                    var v56 int32 = 594
                                    _ = v56
                                }
                            } else {
                                fmt.Println(v2, v1, v2)
                                delete(v5, 1)
                                fmt.Println(fi53, v1)
                            }
                            v35 = append(v35, -773)
                            v57 := GenS1{F0: (v2 % ((v1) | 1)), F1: -2147483648, F2: fi54}
                            _ = fi54
                            _ = v57
                        }
                        v4 -= v4
                        _ = fi53
                    }
                    switch "foo bar" {
                    case "\n\t":
                        fmt.Println(v4, v4, v4)
                    case "":
                        for fi58 := 0; fi58 < 13; fi58++ {
                            for fi59 := 0; fi59 < 8; fi59++ {
                                delete(v3, 943)
                                _ = fi59
                            }
                            _ = fi58
                        }
                    case "a":
                        var v60 bool = true
                        _ = v60
                    default:
                        v61 := GenS7{F0: false, F1: 0, F2: 2147483647}
                        _ = v61
                    }
                }
                if true {
                    var v62 float32 = 3.14159
                    _ = v62
                }
                for fi63 := 0; fi63 < 5; fi63++ {
                    var v64 int32 = 1
                    fmt.Println(v64)
                    var v65 float64 = 1e10
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            } else {
                fmt.Println(v2, v1, v4)
                fmt.Println(v1)
            }
        }
        v66 := GenS3{F0: v1, F1: -866}
        v67 := GenS8{F0: 981}
        _ = v66
        _ = v67
    } else {
        _ = len(v6)
    }
    if (942 <= 2) {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v1, v1)
        case "hello":
            v1 = v2
        case "":
            v1 += (^(v1))
        }
        for fi68 := 0; fi68 < 12; fi68++ {
            var v69 int64 = 42
            if (false || false) {
                var v70 bool = (genFn1((v69 / (((v69 / (((v69 % ((v69) | 1))) | 1))) | 1))) != (v4 | v4))
                _ = v70
            }
            delete(v6, (fi68 << ((255) & 63)))
            _ = fi68
            _ = v69
        }
    }
    v2 = (v2 % (((v1 | (v2 ^ (v1 ^ v1)))) | 1))
    fmt.Println(v1, v1, v2)
    switch -622 {
    case -408:
        var v71 uint64 = 4294967295
        _ = v71
    case 32767:
        v2 += (v1 << ((124) & 63))
    case -9223372036854775808:
        v72 := GenS9{F0: v2, F1: 65535}
        _ = v72
    }
    v73 := GenS8{F0: 1}
    fmt.Println(v1, v2, v1)
    var v74 uint32 = 4294967295
    if true {
        if len(v35) >= 2 {
            v35 = v35[1:len(v35)]
        }
        for fi75 := 0; fi75 < 14; fi75++ {
            v1 = (^(v2))
            switch "\n\t" {
            case "hello":
                var v76 string = "a"
                _ = v76
            case "\n\t":
                v77 := GenS6{F0: 4294967295, F1: (v2 - (v2 / (((v2 % (((^(v2))) | 1))) | 1))), F2: 0}
                _ = v77
            }
            _ = fi75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v73
    _ = v74
}
