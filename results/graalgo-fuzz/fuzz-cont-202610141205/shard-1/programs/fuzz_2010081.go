package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 int; F1 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64 }
func genFn0() uint64 {
    var v4 uint32 = 0
    v5 := GenS1{F0: "foo bar", F1: 42}
    _ = v4
    _ = v5
    return 18446744073709551615
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 42
    v3 := []uint32{}
    v1 = (v1 | (v1 ^ 4294967295))
    switch "\n\t" {
    case "foo bar":
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint8 = 231
            if true {
                v7 -= (v7 >> ((42) & 63))
                v2 = 0
            }
            if ((fi6 >> (((^(v1))) & 63)) >= (fi6 * fi6)) {
                fmt.Println(v2)
                var v8 uint = v1
                for fi9 := 0; fi9 < 6; fi9++ {
                    v2 -= (v2 / ((v2) | 1))
                    v1 += 255
                    var v10 uint = (v1 ^ 96)
                    var v11 uint16 = 1
                    var v12 bool = false
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                _ = v8
            } else {
                if false {
                    v13 := GenS2{F0: (v2 & int8(fi6))}
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v15 int64 = 1
                        _ = fi14
                        _ = v15
                    }
                    var v16 uint32 = 2
                    _ = v13
                    _ = v16
                } else {
                    v2 -= (v2 % (((v2 >> (((v1 * 2)) & 63))) | 1))
                    var v17 uint32 = 42
                    _ = v17
                }
            }
            _ = fi6
            _ = v7
        }
    case "a":
        switch "hello" {
        case "hello":
            var v18 uint = v1
            _ = v18
        case "foo bar":
            if (!false) {
                v2 -= v2
                if ((v1 ^ 4294967295) <= 896) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 += (v2 * v2)
                } else {
                    if (0 < 0) {
                        for fi19 := 0; fi19 < 9; fi19++ {
                            _ = len(v3)
                            _ = fi19
                        }
                        var v20 uint = 227
                        _ = v20
                    }
                }
            }
        case "\n\t":
            if ((v1 << ((v1) & 63)) < 42) {
                if len(v3) > 0 {
                    v3[0] = 173
                }
            } else {
                var v21 bool = (255 == 523)
                v21 = v21
                v3 = append(v3, 4294967295)
                _ = v21
            }
        default:
            v1 -= 700
        }
    default:
        v22 := GenS3{F0: -504, F1: v1}
        _ = v22
    }
    var v23 int8 = 100
    switch "\n\t" {
    case "hello":
        fmt.Println(v2)
    case "a":
        fmt.Println(v2, v1)
    }
    var v24 float32 = 1.5
    if (!true) {
        v25 := GenS4{F0: 1000, F1: -1}
        fmt.Println(v24)
        var v26 int = -2147483648
        _ = v25
        _ = v26
    }
    if false {
        var v27 uint = v1
        switch "\n\t" {
        case "a":
            if false {
                fmt.Println(v23, v1, v27)
                v2 -= v23
                var v28 int = 9223372036854775807
                _ = v28
            } else {
                _ = len(v3)
                var v29 int32 = 42
                _ = v29
            }
        case "foo bar":
            var v30 int32 = 42
            _ = v30
        case "":
            switch ("hello" + "\n\t") {
            case "a":
                if (42 <= 0) {
                    switch -654 {
                    case -1:
                        switch -208 {
                        case -39:
                            var v31 string = "a"
                            _ = v31
                        }
                    }
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 int32 = -604
                        var v34 uint64 = 18446744073709551615
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v36 := GenS5{F0: 42, F1: -652}
                        _ = fi35
                        _ = v36
                    }
                } else {
                    for fi37 := 0; fi37 < 20; fi37++ {
                        if (!true) {
                            var v38 int64 = -9223372036854775808
                            _ = v38
                        }
                        var v39 uint16 = 386
                        for fi40 := 0; fi40 < 18; fi40++ {
                            fmt.Println(v27, v24)
                            v39 = v39
                            fmt.Println(v23, v2)
                            _ = fi40
                        }
                        _ = fi37
                        _ = v39
                    }
                }
            default:
                if true {
                    switch 4294967295 {
                    case 2147483647:
                        fmt.Println(v1)
                    case -85:
                        var v41 bool = ((((float32(1.5) / v24) - 1e10) / v24) > 1e10)
                        _ = v41
                    case -884:
                        v42 := GenS2{F0: 1}
                        _ = v42
                    default:
                        if true {
                            var v43 float32 = (1.5 / float32(0.1))
                            if true {
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    for fi45 := 0; fi45 < 18; fi45++ {
                                        var v46 string = ""
                                        if (1e10 != 0.1) {
                                            v2 += v2
                                            var v47 int16 = -32768
                                            _ = v47
                                        }
                                        _ = fi45
                                        _ = v46
                                    }
                                    _ = fi44
                                }
                                var v48 int16 = 0
                                v49 := GenS1{F0: "foo bar", F1: (^(v27))}
                                _ = v48
                                _ = v49
                            } else {
                                if (2147483647 != -2147483648) {
                                    var v50 uint64 = 18446744073709551615
                                    var v51 uint32 = 42
                                    if (uint32(v50) > 0) {
                                        v52 := GenS6{F0: 1}
                                        for fi53 := 0; fi53 < 9; fi53++ {
                                            switch "a" {
                                            case "":
                                                var v54 string = "hello"
                                                _ = v54
                                            case "\n\t":
                                                var v55 bool = (uint32(v50) < (v51 | (v51 & (v51 - (v51 >> ((18446744073709551615) & 63))))))
                                                _ = v55
                                            }
                                            _ = fi53
                                        }
                                        _ = v52
                                    }
                                    _ = v50
                                    _ = v51
                                } else {
                                    v23 = (v23 >> ((v1) & 63))
                                }
                            }
                            _ = v43
                        }
                    }
                    v56 := GenS7{F0: 2147483648, F1: 0}
                    v57 := GenS8{F0: -990, F1: 127}
                    _ = v56
                    _ = v57
                }
            }
        }
        _ = v27
    } else {
        switch ("\n\t" + "hello") {
        case "foo bar":
            for fi58 := 0; fi58 < 4; fi58++ {
                v59 := GenS9{F0: (fi58 % ((9223372036854775807) | 1)), F1: 65535}
                if (genFn0() < 890) {
                    v2 -= -128
                    for fi60 := 0; fi60 < 5; fi60++ {
                        fmt.Println(v2, v23, v1)
                        if (v1 == v1) {
                            v23 = (-(v23))
                            if len(v3) > 0 {
                                v3[0] = 830
                            }
                        } else {
                            var v61 uint16 = 0
                            v61 -= (v61 - 940)
                            var v62 float32 = v24
                            _ = v61
                            _ = v62
                        }
                        var v63 int32 = 1
                        _ = fi60
                        _ = v63
                    }
                    for fi64 := 0; fi64 < 3; fi64++ {
                        if true {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v65 := GenS5{F0: 42, F1: 42}
                            if (-2.25 != (3.14159 / 1.5)) {
                                v2 = (v2 & v2)
                            } else {
                                v23 = v2
                                switch (fi58 * (fi64 >> ((v1) & 63))) {
                                case 420:
                                    if ((!((!false) && (128 >= 127))) && true) {
                                        var v66 int64 = 42
                                        _ = v66
                                    }
                                case 799:
                                    v1 += v1
                                case -9223372036854775808:
                                    var v67 string = "hello"
                                    _ = v67
                                default:
                                    v1 += (v1 - 0)
                                }
                                if ((fi64 - (fi64 * (fi58 - -9223372036854775808))) <= 15) {
                                    var v68 string = "\n\t"
                                    v69 := GenS4{F0: 796, F1: 42}
                                    _ = v68
                                    _ = v69
                                }
                            }
                            v23 = (v23 ^ v2)
                            _ = v65
                        }
                        v70 := GenS9{F0: int(v1), F1: 42}
                        _ = fi64
                        _ = v70
                    }
                } else {
                    for fi71 := 0; fi71 < 2; fi71++ {
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        _ = fi71
                    }
                }
                v72 := GenS10{F0: 0}
                _ = fi58
                _ = v59
                _ = v72
            }
        default:
            v23 = v2
        }
        var v73 int16 = 364
        var v74 bool = (false || (!true))
        _ = v73
        _ = v74
    }
    v75 := GenS11{F0: 1}
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v75
}
