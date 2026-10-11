package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int32, p1 int8, p2 uint16) bool {
    var v7 uint = 255
    v7 = v7
    for fi8 := 0; fi8 < 15; fi8++ {
        if false {
            p0 -= (p0 + 2147483647)
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint16 = p2
                _ = fi9
                _ = v10
            }
            switch (fi8 + fi8) {
            case 32768:
                p1 = p1
            }
        }
        p1 += (p1 - (-(p1)))
        _ = fi8
    }
    fmt.Println(p2)
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}
func genFn1(p0 int) uint8 {
    if false {
        v11 := GenS1{F0: 1}
        v12 := []int{(-(p0))}
        var v13 int8 = 100
        _ = v11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        p0 -= (fi14 & -9223372036854775808)
        var v15 int32 = -2147483648
        if (-128 >= 1) {
            switch ("" + ("hello" + ("\n\t" + "a"))) {
            case "foo bar":
                v16 := GenS1{F0: 2}
                _ = v16
            case "":
                p0 += (p0 - (fi14 << ((255) & 63)))
            case "hello":
                fmt.Println(fi14)
            default:
                p0 += fi14
            }
            if false {
                var v17 int64 = 0
                _ = v17
            } else {
                v18 := GenS2{F0: 1}
                v15 += (v15 + (^(v15)))
                if ((false || (false && true)) && true) {
                    fmt.Println(p0, fi14, v15)
                }
                _ = v18
            }
        } else {
            v15 = (v15 >> ((0) & 63))
            var v19 int8 = 42
            v19 += v19
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    v20 := []int{(p0 + (p0 >> ((0) & 63))), p0, 4294967295}
    if (false && false) {
        var v21 string = ("a" + (("hello" + "hello") + "hello"))
        var v22 int = (p0 - -2147483648)
        var v23 uint32 = 103
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println(p0)
    _ = p0
    _ = v20
    return 216
}

func main() {
    var v1 uint16 = 42
    var v2 float32 = -2.25
    var v3 float64 = (1e10 * float64((((v2 + (v2 + 1e10)) + 1.5) + ((v2 / (((v2 / 0.1) / v2) - float32(3.14159))) - float32(1e10)))))
    var v4 uint16 = v1
    v5 := make(map[int]int)
    var v6 int = -246
    var v24 float32 = (v2 * 0.1)
    if true {
        for fi25 := 0; fi25 < 19; fi25++ {
            v24 += v2
            var v26 uint8 = 0
            if (129 > 2) {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 bool = (80 <= 100)
                    v4 = (v4 / (((v4 / ((v4) | 1))) | 1))
                    if v28 {
                        if (v28 || v28) {
                            if (!(42 == 156)) {
                                v2 = float32(((v3 - 3.14159) - v3))
                            }
                            v6 -= (^(fi27))
                            var v29 int32 = 0
                            _ = v29
                        }
                        switch fi27 {
                        case 720:
                            if v28 {
                                fmt.Println(v26)
                            }
                        case -2147483648:
                            var v30 uint16 = 32767
                            _ = v30
                        case 32768:
                            var v31 uint = 0
                            _ = v31
                        default:
                            v32 := GenS1{F0: 18446744073709551615}
                            _ = v32
                        }
                    }
                    _ = fi27
                    _ = v28
                }
            } else {
                switch 0 {
                case 705:
                    switch "\n\t" {
                    case "a":
                        for fi33 := 0; fi33 < 2; fi33++ {
                            v1 -= (^(v1))
                            _ = fi33
                        }
                    case "\n\t":
                        delete(v5, (fi25 & 127))
                    default:
                        v2 += ((((0.0 / 1.5) * v24) * 0.1) - (v24 / (v24 - 3.14159)))
                    }
                case 237:
                    v34 := GenS3{F0: 0, F1: 0, F2: v1}
                    _ = v34
                }
                switch "\n\t" {
                case "hello":
                    var v35 int = v6
                    _ = v35
                case "a":
                    for fi36 := 0; fi36 < 5; fi36++ {
                        var v37 uint64 = 4294967295
                        v4 -= 2
                        v37 += 4294967295
                        _ = fi36
                        _ = v37
                    }
                case "":
                    v38 := GenS4{F0: "foo bar", F1: v26}
                    _ = v38
                default:
                    switch 4294967295 {
                    case 0:
                        v39 := GenS1{F0: 18446744073709551615}
                        _ = v39
                    default:
                        if (!(100 <= 42)) {
                            v40 := GenS4{F0: "foo bar", F1: v26}
                            _ = v40
                        } else {
                            for fi41 := 0; fi41 < 19; fi41++ {
                                switch "\n\t" {
                                case "foo bar":
                                    v3 = (v3 + (v3 + 1e10))
                                default:
                                    _ = len(v5)
                                }
                                v42 := GenS4{F0: (("a" + "a") + ("\n\t" + ("hello" + ("\n\t" + ("a" + ("" + "hello")))))), F1: 128}
                                fmt.Println(v1, v4)
                                fmt.Println(v1, v4)
                                _ = fi41
                                _ = v42
                            }
                            var v43 int64 = 0
                            v24 -= v24
                            _ = v43
                        }
                    }
                }
            }
            _ = fi25
            _ = v26
        }
        v1 = v1
    } else {
        var v44 uint32 = 65535
        v3 += ((v3 / 0.1) - (v3 * v3))
        _ = v44
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        for fi46 := 0; fi46 < 10; fi46++ {
            v47 := GenS3{F0: -485, F1: 42, F2: (v4 | v4)}
            if true {
                var v48 int8 = 1
                var v49 float64 = v3
                v50 := []int{(fi46 % (((^(v6))) | 1)), 0, v6}
                _ = v48
                _ = v49
                _ = v50
            } else {
                var v51 uint32 = uint32(v4)
                v4 += (v1 * (v4 / (((v4 >> ((18446744073709551615) & 63))) | 1)))
                v52 := GenS2{F0: v1}
                _ = v51
                _ = v52
            }
            var v53 uint = 1
            _ = fi46
            _ = v47
            _ = v53
        }
        var v54 uint = 65
        v1 = 32767
        _ = fi45
        _ = v54
    }
    fmt.Println(v1)
    v4 -= 1
    if true {
        v55 := GenS3{F0: -516, F1: 18446744073709551615, F2: v4}
        for fi56 := 0; fi56 < 17; fi56++ {
            var v57 float32 = v24
            _ = fi56
            _ = v57
        }
        _ = v55
    } else {
        v58 := v5[(v6 * 127)]
        if (!true) {
            v59 := GenS2{F0: v4}
            switch v6 {
            case 42:
                v3 += (v3 * v3)
            default:
                v6 -= 32767
            }
            v1 = (v4 ^ 2)
            _ = v59
        } else {
            fmt.Println(v6, v2, v3)
        }
        if false {
            v60 := GenS1{F0: 875}
            _ = v60
        } else {
            var v61 int16 = 32767
            v4 += (v4 | (^(v4)))
            for fi62 := 0; fi62 < 7; fi62++ {
                v6 += v58
                v61 -= 1000
                _ = fi62
            }
            _ = v61
        }
        _ = v58
    }
    for fi63 := 0; fi63 < 3; fi63++ {
        var v64 uint = 65535
        v64 = 156
        switch (v6 ^ fi63) {
        case 127:
            var v65 uint16 = v1
            _ = v65
        default:
            fmt.Println(v3, v64, v2)
        }
        _ = fi63
        _ = v64
    }
    v2 = v24
    for fi66 := 0; fi66 < 12; fi66++ {
        v67 := []int{328, (v6 & 63), fi66, (fi66 * (-(fi66)))}
        _ = len(v67)
        _ = fi66
        _ = v67
    }
    fmt.Println(v2, v3, v24)
    var v68 int = (v6 | (v6 + (v6 - (v6 % ((323) | 1)))))
    for fi69 := 0; fi69 < 20; fi69++ {
        var v70 uint16 = (^(v4))
        var v71 uint8 = 241
        var v72 int64 = 0
        _ = fi69
        _ = v70
        _ = v71
        _ = v72
    }
    if (!(true || (true || false))) {
        v73 := GenS5{F0: 42, F1: -1, F2: 4294967295}
        for fi74 := 0; fi74 < 3; fi74++ {
            fmt.Println(v3, v68)
            _ = fi74
        }
        switch v68 {
        case 4294967295:
            v75 := GenS5{F0: 1, F1: -9223372036854775808, F2: 0}
            _ = v75
        default:
            for fi76 := 0; fi76 < 2; fi76++ {
                switch "a" {
                case "foo bar":
                    switch v6 {
                    case -255:
                        delete(v5, fi76)
                    case 32767:
                        for fi77 := 0; fi77 < 10; fi77++ {
                            if false {
                                v78 := GenS5{F0: 914, F1: 1, F2: 18446744073709551615}
                                if (1e10 >= v2) {
                                    v68 = 32768
                                }
                                _ = v78
                            } else {
                                var v79 int8 = -69
                                v1 = v4
                                fmt.Println(v3, v68)
                                _ = v79
                            }
                            v24 = (v24 + (1.5 * v24))
                            _ = fi77
                        }
                    case -674:
                        fmt.Println(v24, fi76)
                    default:
                        v80 := GenS2{F0: (v4 * (v4 << ((18446744073709551615) & 63)))}
                        _ = v80
                    }
                case "\n\t":
                    v81 := []int{(^(v68)), fi76, fi76, v6}
                    _ = v81
                case "hello":
                    v82 := GenS6{F0: 255}
                    _ = v82
                default:
                    v68 -= (fi76 + 119)
                }
                _ = fi76
            }
        }
        _ = v73
    } else {
        var v83 uint = 0
        v84 := GenS5{F0: 42, F1: 9223372036854775807, F2: v83}
        v85 := GenS6{F0: 255}
        _ = v83
        _ = v84
        _ = v85
    }
    v1 = (^(v1))
    v86 := []int{(v68 / (((v68 / (((^(v6))) | 1))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v86:", len(v86), v86)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v86
}
