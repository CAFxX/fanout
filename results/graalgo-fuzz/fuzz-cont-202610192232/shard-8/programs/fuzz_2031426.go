package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint32; F1 uint64; F2 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 string; F2 uint8 }
type GenS7 struct { F0 int64 }
func genFn0(p0 int64) uint {
    var v4 int64 = p0
    for fi5 := 0; fi5 < 18; fi5++ {
        p0 -= v4
        switch "hello" {
        case "foo bar":
            var v6 float64 = 0.1
            _ = v6
        case "a":
            fmt.Println(fi5)
        default:
            v7 := []int{}
            _ = v7
        }
        _ = fi5
    }
    var v8 int64 = -728
    v8 -= -346
    _ = p0
    _ = v4
    _ = v8
    return 255
}
func genFn1() int {
    var v9 uint16 = 65535
    v10 := GenS1{F0: v9, F1: 65535}
    v11 := GenS2{F0: 1, F1: 42, F2: -1}
    var v12 uint = 254
    for fi13 := 0; fi13 < 20; fi13++ {
        if true {
            switch -1 {
            case 405:
                fmt.Println(v12)
            case -101:
                var v14 int8 = -47
                _ = v14
            case 0:
                v15 := []int{1, (fi13 * 42), fi13, fi13}
                _ = v15
            }
        }
        _ = fi13
    }
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 449
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -651
    v3 := make(map[uint32]string)
    if true {
        fmt.Println(v2)
    }
    switch "foo bar" {
    case "\n\t":
        var v16 float64 = 0.0
        _ = v16
    case "hello":
        if false {
            delete(v1, 104)
            if false {
                switch "hello" {
                case "foo bar":
                    if (37 != 65535) {
                        var v17 float64 = float64(float32(0.1))
                        var v18 uint32 = 1
                        switch "a" {
                        case "a":
                            fmt.Println(v2, v2, v2)
                        case "\n\t":
                            switch genFn1() {
                            case 128:
                                if false {
                                    v2 = genFn1()
                                    for fi19 := 0; fi19 < 19; fi19++ {
                                        for fi20 := 0; fi20 < 18; fi20++ {
                                            var v21 int32 = -1
                                            if (false && (-1 > v21)) {
                                                switch "a" {
                                                case "":
                                                    var v22 string = ((("\n\t" + "a") + (("hello" + "") + "hello")) + (("" + "hello") + "foo bar"))
                                                    _ = v22
                                                case "\n\t":
                                                    var v23 uint8 = 2
                                                    _ = v23
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    v24 := GenS2{F0: (v18 / ((791) | 1)), F1: 889, F2: (fi20 ^ fi19)}
                                                    v25 := GenS1{F0: 980, F1: 42}
                                                    fmt.Println(v21)
                                                    _ = v24
                                                    _ = v25
                                                default:
                                                    var v26 uint32 = (^(v18))
                                                    _ = v26
                                                }
                                                v27 := GenS3{F0: 100}
                                                _ = v27
                                            } else {
                                                v18 = (v18 ^ v18)
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                        _ = fi19
                                    }
                                    switch v2 {
                                    case 622:
                                        v28 := []int{}
                                        _ = v28
                                    }
                                }
                            case -1:
                                for fi29 := 0; fi29 < 9; fi29++ {
                                    var v30 int64 = -67
                                    fmt.Println(v17, v18)
                                    v31 := []int{(^(fi29)), (-(fi29))}
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                            case -2147483648:
                                switch "" {
                                case "":
                                    v2 += v2
                                }
                            default:
                                v18 = v18
                            }
                        default:
                            for fi32 := 0; fi32 < 15; fi32++ {
                                v33 := GenS1{F0: 632, F1: genFn0(42)}
                                for fi34 := 0; fi34 < 18; fi34++ {
                                    fmt.Println(fi34, v17)
                                    _ = fi34
                                }
                                v35 := GenS4{F0: 7, F1: 42, F2: 18446744073709551615}
                                _ = fi32
                                _ = v33
                                _ = v35
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v2)
                        fmt.Println(v2)
                        var v36 uint32 = 1
                        _ = v36
                    }
                default:
                    for fi37 := 0; fi37 < 7; fi37++ {
                        var v38 uint64 = 727
                        var v39 uint8 = 79
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                }
            } else {
                if false {
                    fmt.Println(v2)
                    for fi40 := 0; fi40 < 9; fi40++ {
                        v41 := []int{(fi40 ^ genFn1())}
                        fmt.Println(v2, v2)
                        _ = fi40
                        _ = v41
                    }
                } else {
                    var v42 uint16 = 582
                    var v43 uint8 = 42
                    _ = v42
                    _ = v43
                }
            }
            for fi44 := 0; fi44 < 11; fi44++ {
                fmt.Println(v2, v2)
                _ = fi44
            }
        } else {
            for fi45 := 0; fi45 < 7; fi45++ {
                v2 = fi45
                if ("hello" == "a") {
                    fmt.Println(v2, fi45)
                    v46 := v1[2147483647]
                    _ = v46
                } else {
                    var v47 uint16 = 255
                    v2 += v2
                    _ = v47
                }
                v48 := GenS5{F0: (((1.5 * 0.0) - (3.14159 * (1e10 * (-2.25 * 3.14159)))) != 3.14159)}
                _ = fi45
                _ = v48
            }
            v2 = 234
            v49 := GenS3{F0: -128}
            _ = v49
        }
    }
    v50 := GenS3{F0: 0}
    v51 := GenS2{F0: 197, F1: 893, F2: -2147483648}
    v2 -= genFn1()
    var v52 int16 = -32768
    v53 := GenS1{F0: 32767, F1: 4294967295}
    v52 = (v52 << ((4294967295) & 63))
    if false {
        v52 += (^(v52))
        var v54 uint = 255
        _ = v54
    }
    v55 := GenS1{F0: 65535, F1: 255}
    if true {
        v52 = v52
    } else {
        for fi56 := 0; fi56 < 16; fi56++ {
            var v57 uint64 = 4294967295
            for fi58 := 0; fi58 < 6; fi58++ {
                if (false || (1 > 138)) {
                    v59 := GenS1{F0: 314, F1: genFn0(42)}
                    _ = v59
                }
                _ = fi58
            }
            switch "foo bar" {
            case "":
                var v60 float64 = 1.5
                _ = v60
            }
            _ = fi56
            _ = v57
        }
        fmt.Println(v52)
        var v61 int32 = -1
        _ = v61
    }
    var v62 int64 = 42
    var v63 int8 = 42
    v64 := GenS6{F0: 829, F1: "hello", F2: 127}
    for fi65 := 0; fi65 < 20; fi65++ {
        var v66 bool = true
        _ = fi65
        _ = v66
    }
    if (42 < 466) {
        if true {
            for fi67 := 0; fi67 < 9; fi67++ {
                v68 := []int{(v2 << ((2) & 63)), (v2 - v2), fi67, genFn1()}
                var v69 int16 = 331
                _ = fi67
                _ = v68
                _ = v69
            }
            v70 := GenS3{F0: (v63 % (((^(v63))) | 1))}
            _ = v70
        }
        switch 9223372036854775807 {
        case 1:
            v63 += v63
        case -2147483648:
            for fi71 := 0; fi71 < 2; fi71++ {
                _ = len(v3)
                for fi72 := 0; fi72 < 3; fi72++ {
                    var v73 uint8 = 42
                    _ = len(v3)
                    var v74 uint8 = 128
                    _ = fi72
                    _ = v73
                    _ = v74
                }
                switch "foo bar" {
                case "foo bar":
                    v75 := []int{(fi71 + fi71)}
                    _ = v75
                case "\n\t":
                    v62 = v62
                case "":
                    v76 := GenS7{F0: v62}
                    _ = v76
                }
                _ = fi71
            }
        case 127:
            var v77 uint32 = 2147483647
            _ = v77
        }
        if true {
            for fi78 := 0; fi78 < 9; fi78++ {
                v3[42] = (("foo bar" + ("a" + ("hello" + ""))) + "\n\t")
                if false {
                    for fi79 := 0; fi79 < 7; fi79++ {
                        var v80 int16 = (v52 / (((^(v52))) | 1))
                        _ = fi79
                        _ = v80
                    }
                    var v81 uint32 = 0
                    _ = v81
                }
                _ = fi78
            }
        }
    } else {
        var v82 float64 = 0.1
        v83 := GenS4{F0: 100, F1: -901, F2: 1}
        _ = v82
        _ = v83
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v55
    _ = v62
    _ = v63
    _ = v64
}
