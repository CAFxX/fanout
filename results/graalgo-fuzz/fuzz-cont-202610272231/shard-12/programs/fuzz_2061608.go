package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
type GenS5 struct { F0 string; F1 int16; F2 int32 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint; F1 uint; F2 bool }
func genFn0() uint64 {
    var v5 int = 593
    for fi6 := 0; fi6 < 11; fi6++ {
        switch v5 {
        case 42:
            if true {
                for fi7 := 0; fi7 < 6; fi7++ {
                    var v8 int8 = -114
                    _ = fi7
                    _ = v8
                }
            }
        case 9223372036854775807:
            if (0 == 255) {
                v9 := []int{414, -925, fi6, -884}
                var v10 int64 = 568
                _ = v10
                _ = v9
            } else {
                var v11 uint = 1
                v12 := GenS1{F0: v11}
                var v13 int16 = -32768
                _ = v11
                _ = v12
                _ = v13
            }
        default:
            var v14 int32 = -1
            _ = v14
        }
        _ = fi6
    }
    v5 = (v5 ^ 127)
    for fi15 := 0; fi15 < 19; fi15++ {
        v16 := GenS2{F0: 376}
        v17 := GenS1{F0: 4294967295}
        if false {
            v5 = 9223372036854775807
        } else {
            if true {
                v18 := GenS3{F0: -1, F1: 2}
                var v19 string = "foo bar"
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 14; fi20++ {
                    v21 := GenS3{F0: 904, F1: 2}
                    v5 -= -848
                    v5 -= (fi20 << ((42) & 63))
                    if (!(48 <= 2)) {
                        var v22 uint8 = 128
                        if (!(32767 == -32768)) {
                            v5 += v5
                            v23 := GenS4{F0: "foo bar", F1: -119, F2: (v22 << ((42) & 63))}
                            _ = v23
                        }
                        _ = v22
                    } else {
                        v24 := GenS2{F0: 9223372036854775807}
                        v5 = fi15
                        var v25 uint64 = 192
                        v26 := GenS4{F0: ("a" + ""), F1: 42, F2: 42}
                        if false {
                            v27 := []int{-2147483648, 1}
                            var v28 uint = 18446744073709551615
                            if (!false) {
                                _ = len(v27)
                            }
                            _ = v27
                            _ = v28
                        } else {
                            v25 -= v25
                            if true {
                                v29 := GenS5{F0: "hello", F1: -32768, F2: 0}
                                _ = v29
                            }
                        }
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    _ = fi20
                    _ = v21
                }
            }
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v5 = 128
    _ = v5
    return 852
}
func genFn1(p0 bool, p1 int64) uint64 {
    if (2 >= 42) {
        v30 := []int{-9223372036854775808, 9223372036854775807, 127}
        v31 := GenS6{F0: 42, F1: 65535, F2: 1000}
        _ = v30
        _ = v31
    } else {
        if (!(4294967295 != 1)) {
            fmt.Println(p1, p0)
            p1 = (p1 + 601)
            for fi32 := 0; fi32 < 14; fi32++ {
                if (p0 && p0) {
                    if (true || ((-1 >= 0) || p0)) {
                        for fi33 := 0; fi33 < 16; fi33++ {
                            if (4294967295 < 18446744073709551615) {
                                var v34 int8 = 1
                                switch 127 {
                                case -373:
                                    var v35 int8 = int8(p1)
                                    _ = v35
                                case -49:
                                    p1 += p1
                                case 42:
                                    var v36 uint32 = 0
                                    var v37 uint = 158
                                    _ = v36
                                    _ = v37
                                default:
                                    v38 := GenS4{F0: "\n\t", F1: (fi33 + (fi32 & 0)), F2: 42}
                                    _ = v38
                                }
                                if true {
                                    fmt.Println(p0, p1)
                                }
                                if (((p1 + 9223372036854775807) == p1) || (826 < 2)) {
                                    v39 := []int{fi33}
                                    if len(v39) >= 2 {
                                        v39 = v39[1:len(v39)]
                                    }
                                    _ = v39
                                }
                                v40 := GenS7{F0: (65535 >= 2)}
                                _ = v34
                                _ = v40
                            } else {
                                v41 := GenS4{F0: ("hello" + ("hello" + (("foo bar" + "a") + "\n\t"))), F1: (fi33 * fi32), F2: 255}
                                _ = v41
                            }
                            v42 := GenS6{F0: 40, F1: 0, F2: 1}
                            var v43 bool = p0
                            _ = fi33
                            _ = v42
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v45 := GenS8{F0: 2}
                            var v46 string = "\n\t"
                            v47 := GenS9{F0: genFn0()}
                            v48 := GenS1{F0: 18446744073709551615}
                            var v49 uint8 = 2
                            v50 := []int{127, (fi44 - -2147483648), fi44}
                            var v51 uint64 = 697
                            var v52 uint8 = v49
                            var v53 int16 = -32768
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                            _ = v48
                            _ = v49
                            _ = v50
                            _ = v51
                            _ = v52
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 11; fi54++ {
                            switch "\n\t" {
                            case "":
                                v55 := []int{(-(fi32)), (fi32 + (fi32 + 32768)), (fi54 & fi54)}
                                _ = v55
                            case "\n\t":
                                p0 = ("a" != "")
                            case "foo bar":
                                var v56 string = "a"
                                _ = v56
                            default:
                                var v57 float64 = 1.5
                                _ = v57
                            }
                            p1 -= 9223372036854775807
                            v58 := []int{(fi32 + (^(fi32))), 0, 32768, (fi54 - fi32)}
                            _ = fi54
                            _ = v58
                        }
                    } else {
                        v59 := []int{128}
                        _ = v59
                    }
                    if p0 {
                        if p0 {
                            for fi60 := 0; fi60 < 20; fi60++ {
                                switch "foo bar" {
                                case "foo bar":
                                    p1 = -9223372036854775808
                                    v61 := []int{fi32}
                                    p1 = (-(p1))
                                    _ = v61
                                }
                                v62 := []int{(-(fi32))}
                                v63 := GenS1{F0: 582}
                                if p0 {
                                    p0 = (p0 || (!((1 <= -1) || true)))
                                    fmt.Println(p0, p1)
                                }
                                _ = fi60
                                _ = v62
                                _ = v63
                            }
                        } else {
                            p0 = true
                            v64 := GenS3{F0: 497, F1: 18446744073709551615}
                            var v65 int = (fi32 >> ((759) & 63))
                            _ = v64
                            _ = v65
                        }
                        var v66 uint = 161
                        v67 := GenS10{F0: 65535, F1: (v66 / ((v66) | 1)), F2: p0}
                        _ = v66
                        _ = v67
                    }
                } else {
                    for fi68 := 0; fi68 < 4; fi68++ {
                        v69 := GenS4{F0: ("a" + ""), F1: fi68, F2: 0}
                        _ = fi68
                        _ = v69
                    }
                }
                var v70 float64 = 1e10
                fmt.Println(p0, v70, fi32)
                _ = fi32
                _ = v70
            }
        } else {
            for fi71 := 0; fi71 < 12; fi71++ {
                fmt.Println(p1)
                v72 := []int{(^(fi71)), 24, fi71}
                _ = fi71
                _ = v72
            }
        }
        v73 := []int{9223372036854775807, 0, -958, 765}
        _ = v73
    }
    switch "hello" {
    case "foo bar":
        var v74 float32 = 0.1
        _ = v74
    case "a":
        fmt.Println(p1, p0)
    default:
        p1 += 160
    }
    v75 := GenS2{F0: p1}
    var v76 int32 = 1
    _ = p0
    _ = p1
    _ = v75
    _ = v76
    return 2
}

func main() {
    var v1 int32 = -10
    var v2 uint32 = 4294967295
    v3 := make(map[string]string)
    var v4 uint8 = 2
    var v77 int8 = 1
    for fi78 := 0; fi78 < 11; fi78++ {
        v4 = v4
        _ = fi78
    }
    v1 = (v1 | v1)
    v77 -= v77
    var v79 int32 = v1
    var v80 int32 = -2147483648
    v4 = 127
    v1 = -1
    var v81 uint32 = v2
    v80 += (v80 << ((42) & 63))
    var v82 float32 = ((((1e10 - (3.14159 / 1.5)) + 3.14159) * 3.14159) - float32(0.0))
    for fi83 := 0; fi83 < 6; fi83++ {
        v84 := GenS3{F0: 0, F1: 274}
        delete(v3, "foo bar")
        _ = fi83
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v77:", v77)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v79
    _ = v80
    _ = v81
    _ = v82
}
