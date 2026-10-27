package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 int16; F1 uint64; F2 bool }
type GenS5 struct { F0 uint16; F1 bool; F2 int64 }
func genFn0(p0 bool, p1 int, p2 uint16) uint16 {
    var v7 uint8 = 2
    var v8 int16 = 109
    var v9 uint16 = 45
    v10 := []int{(p1 % ((p1) | 1)), 839, int(p2)}
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := []uint32{778, 2147483647, 2147483647, 0}
    var v2 uint32 = 2147483648
    var v3 uint64 = 440
    var v4 int8 = 0
    var v5 int8 = v4
    var v6 uint64 = 18446744073709551615
    for fi11 := 0; fi11 < 20; fi11++ {
        fmt.Println(fi11, v2, v3)
        _ = fi11
    }
    var v12 uint16 = 255
    for fi13 := 0; fi13 < 19; fi13++ {
        v2 = v2
        _ = fi13
    }
    if false {
        if true {
            var v14 uint = 1
            v15 := GenS1{F0: genFn0(true, -2147483648, (^(v12))), F1: (v2 & 65535)}
            v6 = v3
            _ = v14
            _ = v15
        } else {
            _ = len(v1)
            fmt.Println(v4, v12, v2)
        }
        v2 = v2
        var v16 uint32 = (v2 - 121)
        _ = v16
    } else {
        v4 -= (v5 >> ((102) & 63))
        if true {
            v12 -= 1
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int64 = -561
                if false {
                    fmt.Println(v12)
                    v6 = (v3 * v6)
                } else {
                    switch fi17 {
                    case 904:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v6 = 535
                    }
                    var v19 int64 = -9223372036854775808
                    _ = v19
                }
                var v20 uint8 = 0
                _ = fi17
                _ = v18
                _ = v20
            }
        } else {
            var v21 int = 128
            switch (v21 & (v21 << ((16) & 63))) {
            case 534:
                if len(v1) > 0 {
                    v1[0] = 279
                }
            case -838:
                var v22 uint16 = (v12 | uint16(v2))
                _ = v22
            case -9223372036854775808:
                switch -9223372036854775808 {
                case 1:
                    var v23 int16 = 542
                    _ = v23
                case -495:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -636:
                    var v24 int32 = 2147483647
                    _ = v24
                default:
                    fmt.Println(v21)
                }
            default:
                for fi25 := 0; fi25 < 6; fi25++ {
                    var v26 float32 = 1.5
                    v27 := GenS2{F0: 127, F1: 0}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
            _ = v21
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi28 := 0; fi28 < 5; fi28++ {
            _ = len(v1)
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int32 = -375
                _ = fi29
                _ = v30
            }
            switch 996 {
            case 4294967295:
                for fi31 := 0; fi31 < 3; fi31++ {
                    if (1 < 0) {
                        var v32 int = (fi28 ^ fi28)
                        switch fi31 {
                        case -2147483648:
                            var v33 int8 = v5
                            _ = v33
                        case -1:
                            if (true && (true || ("a" != ("a" + ("foo bar" + ""))))) {
                                var v34 int8 = v4
                                _ = v34
                            } else {
                                v1 = append(v1, (^(v2)))
                                v4 = v4
                            }
                        case 222:
                            if true {
                                v6 += v3
                                switch "foo bar" {
                                case "a":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case "\n\t":
                                    if ("a" <= "foo bar") {
                                        v35 := GenS1{F0: 1, F1: (v2 | (v2 ^ 65535))}
                                        if false {
                                            switch "" {
                                            case "":
                                                if true {
                                                    v4 = (v5 % (((v5 / ((v5) | 1))) | 1))
                                                    if false {
                                                        v36 := GenS3{F0: true, F1: -1, F2: -32768}
                                                        _ = v36
                                                    }
                                                    v37 := GenS3{F0: (false && false), F1: -292, F2: int16(v3)}
                                                    _ = v37
                                                } else {
                                                    fmt.Println(v2)
                                                    var v38 float32 = 1.5
                                                    _ = v38
                                                }
                                            case "\n\t":
                                                v1 = append(v1, 108)
                                            default:
                                                _ = len(v1)
                                            }
                                        }
                                        _ = v35
                                    } else {
                                        v39 := GenS2{F0: (v4 | 23), F1: 542}
                                        _ = v39
                                    }
                                    switch ("a" + ("foo bar" + "hello")) {
                                    case "a":
                                        fmt.Println(v6, v5)
                                    case "\n\t":
                                        v40 := GenS2{F0: 0, F1: 0}
                                        _ = v40
                                    default:
                                        fmt.Println(fi28, v3, v6)
                                    }
                                case "hello":
                                    fmt.Println(fi28)
                                }
                            }
                        default:
                            v5 = v4
                        }
                        v2 -= (v2 / (((v2 + (v2 << ((2) & 63)))) | 1))
                        _ = v32
                    }
                    fmt.Println(v4, v2, v5)
                    if true {
                        switch 127 {
                        case 688:
                            _ = len(v1)
                        case 649:
                            switch (fi31 >> ((1) & 63)) {
                            case -2147483648:
                                v6 = (v6 ^ 1)
                            case -1:
                                if true {
                                    var v41 int = 588
                                    _ = v41
                                } else {
                                    v2 += (v2 | (v2 ^ 0))
                                    switch 2147483647 {
                                    case 32768:
                                        if ((true && (0.0 >= (0.0 + (3.14159 + 0.1)))) || (false && false)) {
                                            var v42 uint = 18446744073709551615
                                            _ = v42
                                        }
                                    case -1:
                                        fmt.Println(v4)
                                    default:
                                        v5 += int8(v3)
                                    }
                                    v5 = v4
                                    v2 -= (^(v2))
                                    switch 335 {
                                    case -2147483648:
                                        if len(v1) > 0 {
                                            v1[0] = 491
                                        }
                                    }
                                }
                            }
                        case 9223372036854775807:
                            var v43 int64 = 42
                            _ = v43
                        default:
                            var v44 uint64 = v3
                            _ = v44
                        }
                    } else {
                        if false {
                            fmt.Println(fi28, fi28)
                        }
                        v1 = append(v1, 65535)
                        var v45 uint64 = v3
                        _ = v45
                    }
                    v12 = (v12 + genFn0((v2 <= (v2 >> ((2) & 63))), 129, 255))
                    switch "" {
                    case "foo bar":
                        var v46 string = ""
                        _ = v46
                    default:
                        v47 := GenS2{F0: 118, F1: 42}
                        _ = v47
                    }
                    v6 = v6
                    _ = fi31
                }
            case -1:
                var v48 int = (fi28 * (fi28 / ((-2147483648) | 1)))
                _ = v48
            case 127:
                var v49 float32 = (-2.25 / (-2.25 / 3.14159))
                _ = v49
            }
            _ = fi28
        }
    default:
        for fi50 := 0; fi50 < 5; fi50++ {
            if ((!true) && true) {
                if (-1 > (^(v4))) {
                    _ = len(v1)
                    switch "\n\t" {
                    case "hello":
                        for fi51 := 0; fi51 < 20; fi51++ {
                            v2 -= v2
                            v52 := GenS4{F0: 0, F1: v3, F2: true}
                            v53 := GenS5{F0: (v12 % (((v12 - genFn0(true, fi50, (v12 + v12)))) | 1)), F1: (("hello" + (("" + "") + "\n\t")) >= ""), F2: int64(v5)}
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    default:
                        switch "" {
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        default:
                            var v54 uint32 = 4294967295
                            _ = v54
                        }
                    }
                    fmt.Println(v3)
                }
                v5 -= 100
                var v55 int32 = 1
                _ = v55
            }
            _ = fi50
        }
    }
    fmt.Println(v6, v5, v12)
    v1 = append(v1, (v2 & 42))
    var v56 int16 = -1
    var v57 uint32 = 2
    for fi58 := 0; fi58 < 2; fi58++ {
        var v59 float32 = 3.14159
        v1 = append(v1, v2)
        _ = fi58
        _ = v59
    }
    for fi60 := 0; fi60 < 16; fi60++ {
        if (110 < (^(v57))) {
            if (false || false) {
                for fi61 := 0; fi61 < 3; fi61++ {
                    var v62 int = 1
                    var v63 int32 = 42
                    if (0.0 > 3.14159) {
                        fmt.Println(v3, v63, fi61)
                    }
                    v63 -= (v63 / ((v63) | 1))
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                v64 := GenS4{F0: v56, F1: (v6 % ((0) | 1)), F2: (!false)}
                _ = v64
            }
            var v65 int8 = v4
            _ = v65
        }
        v66 := GenS4{F0: v56, F1: 778, F2: (2 > 255)}
        _ = fi60
        _ = v66
    }
    v6 += (v3 * v3)
    switch -9223372036854775808 {
    case -1:
        var v67 int = int(v4)
        _ = v67
    case 32767:
        if false {
            if false {
                v3 = v6
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
