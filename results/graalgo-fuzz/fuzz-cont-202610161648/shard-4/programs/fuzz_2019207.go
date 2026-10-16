package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 bool }
type GenS5 struct { F0 uint64; F1 int; F2 uint32 }
type GenS6 struct { F0 bool; F1 int8 }
func genFn0(p0 uint, p1 int16) int32 {
    var v7 int32 = 42
    v8 := GenS1{F0: -509}
    if false {
        if (852 < 2147483648) {
            var v9 int32 = (v7 / ((42) | 1))
            var v10 uint8 = 128
            _ = v10
            _ = v9
        } else {
            v11 := []int{-62, 1}
            v12 := GenS2{F0: 127, F1: "\n\t", F2: 53}
            _ = v11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13, v7)
            if (!false) {
                if false {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        v15 := GenS3{F0: -309}
                        v16 := []int{fi13, fi14}
                        fmt.Println(fi13, v7, p0)
                        var v17 int8 = 85
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    v7 = -1
                    if (128 > 191) {
                        if (!true) {
                            v18 := GenS3{F0: 0}
                            var v19 int8 = -1
                            _ = v18
                            _ = v19
                        }
                    }
                }
            }
            _ = fi13
        }
        v20 := []int{2147483647}
        var v21 uint16 = 32767
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(v7))
}
func genFn1() int8 {
    v22 := []int{-9223372036854775808, 32767, -9223372036854775808, 42}
    var v23 uint8 = 128
    if ((1.5 != 0.1) || false) {
        if false {
            if (0.1 != 1.5) {
                fmt.Println(v23)
            }
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 uint32 = 4294967295
                _ = fi24
                _ = v25
            }
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                v23 += 6
                var v27 uint64 = 0
                _ = fi26
                _ = v27
            }
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            if (9223372036854775807 != 672) {
                for fi28 := 0; fi28 < 19; fi28++ {
                    _ = len(v22)
                    _ = fi28
                }
                for fi29 := 0; fi29 < 19; fi29++ {
                    if (-931 != 0) {
                        var v30 uint64 = 2
                        var v31 uint8 = 83
                        _ = v30
                        _ = v31
                    } else {
                        fmt.Println(fi29, fi29)
                    }
                    v32 := GenS3{F0: 9223372036854775807}
                    _ = fi29
                    _ = v32
                }
            } else {
                if len(v22) > 0 {
                    v22[0] = 503
                }
            }
        }
        v23 = (v23 % ((127) | 1))
        if (529 != 4294967295) {
            if (42 == 9223372036854775807) {
                v23 = 1
                switch -648 {
                case 2147483647:
                    var v33 int16 = 1
                    _ = v33
                case -345:
                    if false {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            for fi35 := 0; fi35 < 4; fi35++ {
                                v36 := GenS3{F0: -1}
                                if false {
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    fmt.Println(v23)
                                    v23 -= v23
                                }
                                if ((647 < -171) && true) {
                                    v23 = v23
                                } else {
                                    var v37 uint8 = (^(v23))
                                    if (-1 == -2147483648) {
                                        if ((false || (!(18446744073709551615 <= 0))) && true) {
                                            var v38 string = ""
                                            switch fi34 {
                                            case 0:
                                                v23 += v37
                                            case 128:
                                                v38 = "foo bar"
                                            }
                                            _ = v38
                                        }
                                        v39 := GenS2{F0: (^(v23)), F1: "a", F2: 42}
                                        var v40 string = ("" + "")
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v41 := GenS1{F0: 9223372036854775807}
                                        v37 += (^(v37))
                                        switch "hello" {
                                        case "":
                                            var v42 uint64 = 0
                                            _ = v42
                                        case "foo bar":
                                            v22 = append(v22, fi35)
                                        case "hello":
                                            if false {
                                                v23 = v23
                                            }
                                        }
                                        _ = v41
                                    }
                                    v37 = v37
                                    _ = v37
                                }
                                _ = len(v22)
                                _ = fi35
                                _ = v36
                            }
                            if false {
                                var v43 int32 = 2147483647
                                _ = len(v22)
                                _ = v43
                            }
                            for fi44 := 0; fi44 < 11; fi44++ {
                                v45 := GenS1{F0: 35}
                                v23 -= (^(v23))
                                _ = fi44
                                _ = v45
                            }
                            _ = fi34
                        }
                        v22 = append(v22, 32768)
                    }
                default:
                    v46 := GenS4{F0: 9223372036854775807, F1: 255, F2: false}
                    _ = v46
                }
            }
            var v47 int32 = genFn0(2, -170)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = v47
        } else {
            for fi48 := 0; fi48 < 9; fi48++ {
                fmt.Println(v23, v23)
                for fi49 := 0; fi49 < 6; fi49++ {
                    v23 = 128
                    if (132 > 42) {
                        v23 += v23
                        if (0 <= 697) {
                            fmt.Println(v23, v23)
                        }
                    }
                    _ = fi49
                }
                v50 := GenS4{F0: (fi48 | (fi48 * 0)), F1: 0, F2: (false || false)}
                _ = fi48
                _ = v50
            }
        }
    }
    _ = v22
    _ = v23
    return 127
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    var v4 bool = true
    v5 := make(map[uint32]bool)
    v6 := []uint32{0, 644}
    var v51 int = 32768
    v52 := GenS5{F0: 2, F1: 0, F2: 65535}
    switch "hello" {
    case "\n\t":
        v1 = v1
    case "hello":
        v1 = (1e10 + ((v1 / float64(1.5)) + (float64(3.14159) - v1)))
    case "":
        switch ("" + ((("\n\t" + "\n\t") + ("\n\t" + "hello")) + ("hello" + ("foo bar" + "a")))) {
        case "\n\t":
            v53 := GenS2{F0: 175, F1: "foo bar", F2: 1}
            _ = v53
        default:
            v1 = v1
        }
    default:
        for fi54 := 0; fi54 < 20; fi54++ {
            var v55 uint16 = 898
            if (!(42 != 1)) {
                v56 := GenS6{F0: (2 != 42), F1: 42}
                v3 = v3
                _ = v56
            } else {
                fmt.Println(v1, v1)
                v4 = v4
            }
            for fi57 := 0; fi57 < 16; fi57++ {
                var v58 int8 = 42
                _ = fi57
                _ = v58
            }
            _ = fi54
            _ = v55
        }
    }
    var v59 int64 = -887
    v60 := GenS3{F0: v59}
    if (!false) {
        v4 = v4
        var v61 uint16 = 0
        v4 = true
        _ = v61
    } else {
        for fi62 := 0; fi62 < 12; fi62++ {
            var v63 uint64 = 45
            switch 42 {
            case -884:
                fmt.Println(v63, v1, v4)
            case 2147483647:
                v4 = v4
            case 127:
                fmt.Println(v1, v3)
            default:
                fmt.Println(v63, v59)
            }
            v64 := GenS2{F0: 51, F1: "foo bar", F2: -35}
            _ = fi62
            _ = v63
            _ = v64
        }
        v51 -= v51
    }
    var v65 uint8 = 208
    var v66 uint64 = 1
    if (0 == (v65 * 157)) {
        var v67 int32 = 775
        var v68 int8 = genFn1()
        for fi69 := 0; fi69 < 18; fi69++ {
            v68 -= -128
            fmt.Println(v51, v59, v66)
            fmt.Println(v68, v66)
            _ = fi69
        }
        _ = v67
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v59
    _ = v6
    _ = v60
    _ = v65
    _ = v66
}
