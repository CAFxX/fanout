package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int8; F2 int32 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64; F1 int16; F2 int8 }
type GenS11 struct { F0 uint64; F1 uint; F2 bool }
type GenS12 struct { F0 int; F1 int64 }
func genFn0() uint {
    switch "hello" {
    case "":
        v5 := []int{32768, 127}
        _ = v5
    case "hello":
        var v6 string = ""
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 11; fi7++ {
            if true {
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 uint = 655
                    _ = fi8
                    _ = v9
                }
                var v10 int8 = -39
                v11 := []int{1, fi7, fi7}
                _ = v10
                _ = v11
            }
            v12 := []int{(fi7 & 127), 2147483647, (fi7 / ((fi7) | 1)), (fi7 + 1)}
            var v13 uint = 2
            _ = fi7
            _ = v12
            _ = v13
        }
    }
    var v14 uint16 = 701
    _ = v14
    return 18446744073709551615
}
func genFn1(p0 string, p1 int16) int64 {
    v15 := []int{32768, -1}
    v15 = append(v15, -521)
    v16 := GenS1{F0: 2147483647, F1: 348, F2: 2}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -9223372036854775808
}

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 1
    v3 := make(map[uint32]int)
    var v4 uint = v1
    switch "a" {
    case "\n\t":
        var v17 float64 = ((-2.25 / 3.14159) - -2.25)
        _ = v17
    default:
        v18 := GenS1{F0: -1, F1: 1, F2: 41}
        _ = v18
    }
    switch -532 {
    case 0:
        var v19 uint = v4
        _ = v19
    case -995:
        switch 567 {
        case 42:
            v20 := GenS1{F0: -2147483648, F1: 2147483648, F2: uint8(v4)}
            _ = v20
        case 32768:
            fmt.Println(v1, v1)
        default:
            v21 := GenS2{F0: 255}
            _ = v21
        }
    case -2147483648:
        v22 := GenS1{F0: 2147483647, F1: 4294967295, F2: 42}
        _ = v22
    default:
        _ = len(v3)
    }
    if (42 != 42) {
        v23 := GenS3{F0: 42, F1: 51}
        for fi24 := 0; fi24 < 14; fi24++ {
            for fi25 := 0; fi25 < 15; fi25++ {
                v26 := GenS2{F0: (v4 >> (((v1 ^ v4)) & 63))}
                switch ((("" + "foo bar") + "") + (("hello" + "") + ("hello" + ("a" + ("" + "a"))))) {
                case "a":
                    fmt.Println(fi25)
                case "":
                    switch fi25 {
                    case 42:
                        v27 := GenS4{F0: 42}
                        _ = v27
                    default:
                        v28 := GenS1{F0: 717, F1: 356, F2: 240}
                        _ = v28
                    }
                default:
                    if (0.1 < 0.0) {
                        v29 := GenS5{F0: (((0.1 - 0.0) > 3.14159) || true), F1: 42}
                        _ = v29
                    } else {
                        fmt.Println(v1, v2)
                    }
                }
                _ = fi25
                _ = v26
            }
            v2 -= 2
            for fi30 := 0; fi30 < 5; fi30++ {
                if (((v1 >> ((800) & 63)) == (v1 * v1)) || true) {
                    switch fi24 {
                    case -1:
                        v4 = v4
                    case 32767:
                        delete(v3, 611)
                    default:
                        var v31 int8 = 42
                        _ = v31
                    }
                    switch ("a" + "hello") {
                    case "hello":
                        fmt.Println(fi24, v1)
                        var v32 uint16 = v2
                        if ((v4 << (((v1 - v4)) & 63)) >= (^(v1))) {
                            if false {
                                v33 := []int{fi30, fi30, fi30, fi30}
                                v4 -= v1
                                v3[2] = fi30
                                _ = v33
                            } else {
                                v34 := GenS3{F0: 1, F1: 65535}
                                var v35 float32 = 1.5
                                _ = v34
                                _ = v35
                            }
                            fmt.Println(v32, v2)
                        }
                        v36 := GenS6{F0: (42 == (v1 | v1))}
                        v4 += (^(v1))
                        _ = v32
                        _ = v36
                    }
                    switch ("" + "\n\t") {
                    case "foo bar":
                        v4 = v4
                    default:
                        var v37 string = ("hello" + "foo bar")
                        _ = v37
                    }
                    v38 := v3[131]
                    var v39 float32 = 3.14159
                    _ = v38
                    _ = v39
                } else {
                    v40 := []int{}
                    var v41 int32 = -680
                    _ = v40
                    _ = v41
                }
                _ = fi30
            }
            _ = fi24
        }
        _ = v23
    } else {
        v2 += v2
        var v42 uint8 = 128
        if false {
            for fi43 := 0; fi43 < 20; fi43++ {
                v1 += v4
                v44 := GenS7{F0: (2147483647 <= 831), F1: 1, F2: 42}
                _ = fi43
                _ = v44
            }
        } else {
            var v45 uint16 = 65535
            v2 -= (v2 & v2)
            var v46 float64 = 3.14159
            _ = v45
            _ = v46
        }
        _ = v42
    }
    var v47 uint16 = (^(v2))
    switch "\n\t" {
    case "hello":
        var v48 int16 = 557
        _ = v48
    default:
        switch ("foo bar" + (("\n\t" + "foo bar") + "hello")) {
        case "\n\t":
            for fi49 := 0; fi49 < 9; fi49++ {
                if true {
                    for fi50 := 0; fi50 < 4; fi50++ {
                        v1 = uint(v47)
                        if (42 <= 0) {
                            v51 := GenS6{F0: (!(false && false))}
                            var v52 float32 = 0.0
                            v4 -= v1
                            _ = v51
                            _ = v52
                        } else {
                            v53 := GenS2{F0: (^(v1))}
                            var v54 int64 = int64(v4)
                            _ = v53
                            _ = v54
                        }
                        var v55 uint8 = uint8(v47)
                        if (!(-767 != int16(fi50))) {
                            var v56 float32 = 0.1
                            _ = v56
                        } else {
                            var v57 float64 = 0.1
                            v4 -= genFn0()
                            _ = v57
                        }
                        _ = fi50
                        _ = v55
                    }
                    fmt.Println(v47)
                    var v58 uint64 = 4294967295
                    _ = v58
                } else {
                    var v59 uint = 42
                    switch 0 {
                    case -9223372036854775808:
                        switch "a" {
                        case "foo bar":
                            var v60 uint64 = 18446744073709551615
                            _ = v60
                        default:
                            v2 += 1
                        }
                    default:
                        v61 := []int{52, -2147483648, 32768}
                        _ = v61
                    }
                    _ = v59
                }
                var v62 int = (fi49 - fi49)
                if true {
                    v63 := GenS8{F0: "\n\t", F1: (2 >= 818)}
                    var v64 bool = (1 <= 65535)
                    _ = v63
                    _ = v64
                }
                _ = fi49
                _ = v62
            }
        case "foo bar":
            v2 -= (v2 % ((255) | 1))
        case "":
            for fi65 := 0; fi65 < 14; fi65++ {
                v47 = v47
                _ = fi65
            }
        default:
            v1 -= genFn0()
        }
    }
    if (!(9223372036854775807 == 42)) {
        var v66 uint8 = 128
        _ = v66
    }
    v67 := GenS8{F0: "hello", F1: false}
    fmt.Println(v1)
    v68 := GenS7{F0: ((v47 | (v47 * (v47 & (v2 & (v47 / ((v47) | 1)))))) != (^(v2))), F1: -1, F2: 925}
    v69 := GenS9{F0: 878}
    switch -528 {
    case 423:
        v70 := GenS10{F0: int64(v1), F1: -650, F2: -99}
        _ = v70
    default:
        if ((v2 + v2) != (v2 & 65535)) {
            if false {
                v71 := []int{1, 574, 1, -9223372036854775808}
                switch 32768 {
                case 42:
                    for fi72 := 0; fi72 < 13; fi72++ {
                        v71 = append(v71, (-(fi72)))
                        var v73 float32 = 1e10
                        var v74 uint = (v4 / ((v4) | 1))
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                case 128:
                    for fi75 := 0; fi75 < 15; fi75++ {
                        if (1 > genFn1("hello", -536)) {
                            switch fi75 {
                            case 127:
                                _ = len(v3)
                            default:
                                switch "" {
                                case "hello":
                                    v76 := GenS11{F0: 417, F1: (v4 & (^(v4))), F2: ((((1e10 >= 1e10) && ((^(v1)) == v1)) || true) || true)}
                                    _ = v76
                                }
                            }
                        }
                        fmt.Println(v2)
                        _ = fi75
                    }
                }
                _ = v71
            } else {
                v77 := []int{260, 2147483647, 1}
                if true {
                    switch 42 {
                    case 2147483647:
                        v47 += (v47 << ((uint(v2)) & 63))
                    default:
                        var v78 uint32 = 4294967295
                        _ = v78
                    }
                    if true {
                        v4 = genFn0()
                        v79 := GenS5{F0: false, F1: 422}
                        _ = v79
                    } else {
                        if true {
                            if false {
                                v4 = genFn0()
                            }
                            for fi80 := 0; fi80 < 5; fi80++ {
                                for fi81 := 0; fi81 < 18; fi81++ {
                                    var v82 uint8 = 2
                                    _ = fi81
                                    _ = v82
                                }
                                switch -223 {
                                case 127:
                                    switch (fi80 << ((v4) & 63)) {
                                    case 9223372036854775807:
                                        var v83 float32 = 1.5
                                        _ = v83
                                    case 1:
                                        v47 -= (^(v47))
                                    }
                                case -476:
                                    if (1 != 2) {
                                        for fi84 := 0; fi84 < 10; fi84++ {
                                            switch "a" {
                                            case "hello":
                                                var v85 uint8 = 1
                                                _ = v85
                                            case "\n\t":
                                                v47 -= v47
                                            }
                                            _ = fi84
                                        }
                                        var v86 int16 = 42
                                        if (!true) {
                                            for fi87 := 0; fi87 < 5; fi87++ {
                                                if true {
                                                    v86 += v86
                                                }
                                                v4 += (v1 ^ 1)
                                                _ = len(v77)
                                                delete(v3, 498)
                                                _ = fi87
                                            }
                                            for fi88 := 0; fi88 < 7; fi88++ {
                                                var v89 uint64 = 18446744073709551615
                                                v86 += (v86 & v86)
                                                var v90 bool = (0 != -643)
                                                _ = fi88
                                                _ = v89
                                                _ = v90
                                            }
                                            v2 += (^(v2))
                                        }
                                        _ = v86
                                    }
                                default:
                                    v91 := GenS12{F0: -271, F1: genFn1("a", 42)}
                                    _ = v91
                                }
                                fmt.Println(v47)
                                _ = fi80
                            }
                        }
                    }
                } else {
                    v92 := GenS3{F0: 1, F1: uint32(v4)}
                    var v93 uint64 = 42
                    _ = v92
                    _ = v93
                }
                v94 := v3[0]
                _ = v77
                _ = v94
            }
            switch ("foo bar" + "") {
            case "foo bar":
                fmt.Println(v47, v2, v2)
            default:
                for fi95 := 0; fi95 < 13; fi95++ {
                    if (!false) {
                        var v96 string = "hello"
                        fmt.Println(v1, v2, v96)
                        var v97 uint16 = (v47 / ((620) | 1))
                        _ = v96
                        _ = v97
                    }
                    _ = fi95
                }
            }
        } else {
            for fi98 := 0; fi98 < 3; fi98++ {
                v99 := []int{fi98, 128}
                _ = fi98
                _ = v99
            }
        }
    }
    var v100 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v67
    _ = v68
    _ = v69
}
