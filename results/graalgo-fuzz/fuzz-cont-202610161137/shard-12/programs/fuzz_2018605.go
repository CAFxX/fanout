package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint8; F2 int8 }
func genFn0(p0 int64, p1 int64, p2 uint64) uint8 {
    var v6 uint = 986
    if false {
        if (4294967295 < 288) {
            v7 := []int{}
            if len(v7) > 0 {
                v7[0] = -9223372036854775808
            }
            for fi8 := 0; fi8 < 20; fi8++ {
                fmt.Println(p1)
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = fi8
            }
            _ = v7
        } else {
            switch -291 {
            case 1:
                switch -1 {
                case 688:
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v6 -= v6
                        fmt.Println(v6, p0)
                        _ = fi9
                    }
                case 9223372036854775807:
                    p1 -= -416
                default:
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 uint16 = 944
                        _ = fi10
                        _ = v11
                    }
                }
            case 127:
                var v12 int8 = -128
                _ = v12
            case -9223372036854775808:
                fmt.Println(p1, p1)
            }
            var v13 int = 9223372036854775807
            _ = v13
        }
        var v14 int32 = 2147483647
        _ = v14
    } else {
        var v15 int = 21
        switch "foo bar" {
        case "hello":
            var v16 int = 32768
            _ = v16
        default:
            v17 := []int{-1, 0}
            _ = v17
        }
        _ = v15
    }
    switch ("hello" + "") {
    case "a":
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float32 = -2.25
            switch (fi18 * (fi18 >> (((v6 % (((v6 / ((v6) | 1))) | 1))) & 63))) {
            case -2147483648:
                if (1 > uint32(p2)) {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 uint8 = 254
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 3; fi22++ {
                        switch "" {
                        case "\n\t":
                            v23 := []int{(-(fi22)), (-(fi22)), fi22, 729}
                            _ = v23
                        case "a":
                            var v24 int8 = 76
                            _ = v24
                        case "foo bar":
                            p0 += p1
                        }
                        var v25 int32 = -1
                        var v26 uint16 = 255
                        _ = fi22
                        _ = v25
                        _ = v26
                    }
                    var v27 uint = (v6 >> (((v6 | (v6 | v6))) & 63))
                    _ = v27
                } else {
                    fmt.Println(p1, p1)
                }
            default:
                var v28 int8 = 127
                _ = v28
            }
            if (65535 <= 65535) {
                p1 = (p0 & (p1 + (p0 - 0)))
            } else {
                v29 := []int{fi18}
                if true {
                    var v30 string = ""
                    if (-1 < -1) {
                        var v31 uint32 = 2147483648
                        v30 = ("foo bar" + v30)
                        v6 = uint(p2)
                        _ = v31
                    } else {
                        var v32 uint = 352
                        if true {
                            var v33 bool = (1 == -1)
                            var v34 uint64 = p2
                            var v35 int64 = p0
                            _ = v33
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS1{F0: 932, F1: (v6 * 31)}
                            p2 += 0
                            for fi37 := 0; fi37 < 20; fi37++ {
                                v32 = v6
                                _ = fi37
                            }
                            _ = v36
                        }
                        switch v30 {
                        case "hello":
                            for fi38 := 0; fi38 < 5; fi38++ {
                                fmt.Println(v30, fi38)
                                var v39 uint8 = 128
                                _ = fi38
                                _ = v39
                            }
                        case "foo bar":
                            p0 += p1
                        default:
                            switch fi18 {
                            case 0:
                                p0 -= 0
                            }
                        }
                        _ = v32
                    }
                    _ = v30
                }
                switch (fi18 & 32768) {
                case -608:
                    p0 = p0
                case 361:
                    var v40 uint16 = 716
                    _ = v40
                default:
                    var v41 uint8 = 255
                    _ = v41
                }
                _ = v29
            }
            _ = fi18
            _ = v19
        }
    case "hello":
        v42 := GenS1{F0: 0, F1: uint(p2)}
        _ = v42
    case "":
        p1 += -9223372036854775808
    }
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v6, p1, p2)
    default:
        var v43 int = 0
        _ = v43
    }
    for fi44 := 0; fi44 < 17; fi44++ {
        for fi45 := 0; fi45 < 7; fi45++ {
            v6 = (v6 | v6)
            v6 += (v6 - v6)
            v46 := GenS2{F0: p2, F1: 1000, F2: 1}
            _ = fi45
            _ = v46
        }
        if (float64(0.0) > 1.5) {
            var v47 int64 = 0
            _ = v47
        } else {
            p2 = (p2 / (((^(p2))) | 1))
        }
        var v48 string = ""
        _ = fi44
        _ = v48
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    var v1 uint = 42
    var v2 int16 = -32768
    v3 := make(map[string]int)
    var v4 int64 = int64(v2)
    var v5 uint32 = 448
    if (-713 <= -2147483648) {
        if false {
            v49 := v3["hello"]
            switch v49 {
            case -1:
                v50 := GenS3{F0: ((188 > 42) && false)}
                _ = v50
            case -9223372036854775808:
                v2 -= (v2 + v2)
            case 32767:
                v51 := GenS2{F0: uint64(v1), F1: v2, F2: 0}
                _ = v51
            default:
                v52 := GenS2{F0: 0, F1: (v2 % ((v2) | 1)), F2: -734}
                _ = v52
            }
            _ = v49
        } else {
            if false {
                var v53 int8 = 0
                v54 := []int{95, 596, -244}
                var v55 int = -9223372036854775808
                _ = v53
                _ = v54
                _ = v55
            } else {
                if (0 <= 0) {
                    v1 += (v1 / ((v1) | 1))
                }
                for fi56 := 0; fi56 < 6; fi56++ {
                    var v57 int32 = 851
                    if true {
                        for fi58 := 0; fi58 < 19; fi58++ {
                            var v59 int8 = -128
                            _ = fi58
                            _ = v59
                        }
                        v60 := GenS1{F0: 1, F1: v1}
                        _ = v60
                    }
                    _ = fi56
                    _ = v57
                }
                if true {
                    switch "a" {
                    case "":
                        fmt.Println(v5, v2, v5)
                    default:
                        fmt.Println(v1, v5, v1)
                    }
                    v1 = v1
                    v2 += (v2 ^ v2)
                }
            }
            if true {
                v2 -= 42
            } else {
                v1 = (v1 * v1)
            }
        }
        var v61 int32 = 0
        _ = v61
    } else {
        delete(v3, "")
        v62 := GenS3{F0: ((!(3.14159 == 0.1)) || true)}
        v5 -= 501
        _ = v62
    }
    fmt.Println(v2)
    if false {
        v4 -= 226
    }
    var v63 int32 = int32(v4)
    switch 1 {
    case 128:
        v64 := GenS4{F0: "foo bar", F1: 42}
        _ = v64
    case 0:
        var v65 float64 = (((-2.25 * 1e10) + 0.1) / -2.25)
        _ = v65
    case 32767:
        for fi66 := 0; fi66 < 16; fi66++ {
            v4 = (v4 * 419)
            switch "foo bar" {
            case "hello":
                v63 += (v63 & v63)
            case "\n\t":
                v67 := GenS4{F0: "hello", F1: 110}
                _ = v67
            case "foo bar":
                for fi68 := 0; fi68 < 14; fi68++ {
                    v69 := []int{32767, (fi66 << ((v1) & 63)), fi66}
                    v63 -= (v63 ^ v63)
                    v5 = (v5 << ((255) & 63))
                    _ = fi68
                    _ = v69
                }
            default:
                v70 := []int{(fi66 << (((v1 | (v1 << ((v1) & 63)))) & 63)), fi66, (fi66 % ((fi66) | 1)), fi66}
                _ = v70
            }
            _ = fi66
        }
    default:
        v2 -= v2
    }
    if true {
        delete(v3, "foo bar")
        switch 574 {
        case 32767:
            for fi71 := 0; fi71 < 16; fi71++ {
                if (true && true) {
                    var v72 float64 = 0.0
                    _ = v72
                } else {
                    v5 += (v5 | 65535)
                    for fi73 := 0; fi73 < 15; fi73++ {
                        v2 += 32767
                        var v74 string = "a"
                        _ = fi73
                        _ = v74
                    }
                    var v75 bool = false
                    _ = v75
                }
                var v76 int = fi71
                var v77 float64 = 1e10
                _ = fi71
                _ = v76
                _ = v77
            }
        case 9223372036854775807:
            fmt.Println(v5)
        case 1:
            fmt.Println(v4, v4, v4)
        }
        v5 = v5
    } else {
        v78 := GenS5{F0: (388 <= 1), F1: 0, F2: 100}
        delete(v3, ("hello" + ("foo bar" + ("foo bar" + (("foo bar" + "a") + (("hello" + "") + "\n\t"))))))
        v63 += v63
        _ = v78
    }
    v2 = 66
    var v79 uint32 = (v5 << (((^(v1))) & 63))
    v2 = v2
    var v80 int64 = -965
    v79 += 1
    fmt.Println(v2, v80)
    var v81 float32 = (-2.25 / 1e10)
    var v82 uint8 = genFn0((v4 ^ 1), v80, 42)
    v83 := GenS5{F0: ("" <= "\n\t"), F1: 0, F2: 77}
    v63 += v63
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
