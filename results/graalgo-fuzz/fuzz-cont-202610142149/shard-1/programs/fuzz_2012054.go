package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS10 struct { F0 uint; F1 uint }
type GenS11 struct { F0 uint }
func genFn0(p0 int16) int {
    v7 := GenS1{F0: p0, F1: -1, F2: 402}
    var v8 string = ("a" + "")
    var v9 int64 = 59
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint16, p1 uint16, p2 int8) uint {
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            v11 := []int{}
            p1 = p1
            _ = v11
        } else {
            for fi12 := 0; fi12 < 2; fi12++ {
                v13 := GenS2{F0: 741, F1: true}
                v14 := []int{}
                v15 := GenS2{F0: 2, F1: (!(32767 < p1))}
                p1 -= (p1 << ((255) & 63))
                p0 -= 2
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        _ = fi10
    }
    p2 = (^(p2))
    switch 32767 {
    case -1:
        p1 = (p1 % (((^(p1))) | 1))
    default:
        switch int(p0) {
        case -522:
            fmt.Println(p2, p0, p0)
        default:
            if (1 != 0) {
                if true {
                    var v16 int16 = 0
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v18 := []int{}
                        v19 := GenS1{F0: (v16 / ((v16) | 1)), F1: (fi17 >> ((4294967295) & 63)), F2: (v16 * (-(v16)))}
                        for fi20 := 0; fi20 < 10; fi20++ {
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            fmt.Println(fi20, fi20)
                            v18 = append(v18, 1)
                            _ = fi20
                        }
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                }
                switch int(p2) {
                case 967:
                    p2 += p2
                default:
                    switch -558 {
                    case 1:
                        fmt.Println(p1, p0, p2)
                    case -1:
                        p0 -= 178
                    case 4294967295:
                        var v21 string = (("\n\t" + "foo bar") + "")
                        _ = v21
                    default:
                        v22 := GenS2{F0: 527, F1: (3.14159 < -2.25)}
                        _ = v22
                    }
                }
                v23 := []int{128, 128, 111, genFn0(1000)}
                _ = v23
            }
        }
    }
    if (!(true && false)) {
        for fi24 := 0; fi24 < 8; fi24++ {
            p2 -= p2
            _ = fi24
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    var v3 int8 = -128
    v4 := []uint32{2, 840, 441, 359}
    var v5 uint16 = 430
    var v6 uint64 = v2
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := v1[42]
        var v27 int32 = 42
        v27 += v27
        _ = fi25
        _ = v26
        _ = v27
    }
    v2 = 188
    var v28 string = "\n\t"
    v28 = ("foo bar" + v28)
    if false {
        for fi29 := 0; fi29 < 7; fi29++ {
            v3 = v3
            v5 = 1
            v6 -= (v6 / (((^(v6))) | 1))
            _ = fi29
        }
        v2 = v6
        if (((-(v3)) > v3) || false) {
            if false {
                v4 = append(v4, 2147483647)
                for fi30 := 0; fi30 < 9; fi30++ {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        switch "hello" {
                        case "foo bar":
                            if true {
                                fmt.Println(v5, v2)
                                var v32 uint = 0
                                switch v28 {
                                case "foo bar":
                                    var v33 int16 = -1
                                    _ = v33
                                case "hello":
                                    fmt.Println(v5, v2)
                                }
                                _ = v32
                            }
                        }
                        _ = len(v4)
                        var v34 int16 = 1
                        v35 := GenS3{F0: (!false)}
                        var v36 string = v28
                        _ = fi31
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    v37 := GenS1{F0: int16(v5), F1: (fi30 % ((0) | 1)), F2: 42}
                    _ = fi30
                    _ = v37
                }
            } else {
                var v38 uint64 = v6
                v39 := GenS4{F0: 354, F1: 1}
                var v40 int32 = 42
                _ = v38
                _ = v39
                _ = v40
            }
            v41 := GenS4{F0: (v5 * (^(v5))), F1: 0}
            _ = v41
        } else {
            fmt.Println(v28)
            switch v28 {
            case "hello":
                switch ((((v28 + v28) + v28) + "foo bar") + (v28 + v28)) {
                case "foo bar":
                    v3 = (v3 >> ((genFn1(v5, v5, (v3 & -128))) & 63))
                case "hello":
                    var v42 uint8 = 2
                    _ = v42
                default:
                    var v43 int16 = 293
                    _ = v43
                }
            case "":
                fmt.Println(v5, v28)
            case "a":
                v6 = v6
            }
        }
    } else {
        fmt.Println(v3, v28, v5)
        switch 9223372036854775807 {
        case 32768:
            v2 = v6
        case -550:
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 bool = true
                for fi46 := 0; fi46 < 3; fi46++ {
                    switch 32768 {
                    case 560:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 32768:
                        v4 = append(v4, 138)
                    default:
                        switch 4294967295 {
                        case 127:
                            if ((0 != -2147483648) && true) {
                                v28 = v28
                            }
                        default:
                            var v47 uint16 = (v5 + (^(v5)))
                            _ = v47
                        }
                    }
                    if (true || false) {
                        var v48 uint8 = 113
                        var v49 int16 = 660
                        if (v28 <= (v28 + "")) {
                            switch -9223372036854775808 {
                            case 127:
                                fmt.Println(fi44)
                            case -809:
                                var v50 uint16 = (^(v5))
                                _ = v50
                            }
                            v48 -= uint8(v5)
                            v51 := GenS4{F0: 65535, F1: 4294967295}
                            _ = v51
                        }
                        _ = v48
                        _ = v49
                    }
                    var v52 int16 = 753
                    _ = fi46
                    _ = v52
                }
                _ = fi44
                _ = v45
            }
        case -605:
            v1[438] = 791
        }
    }
    if ((false && (!(!true))) || (!(true && ((-9223372036854775808 < 42) && true)))) {
        for fi53 := 0; fi53 < 19; fi53++ {
            v2 = 759
            _ = fi53
        }
        if len(v4) > 0 {
            v4[0] = 592
        }
    } else {
        fmt.Println(v5, v28, v28)
        var v54 uint8 = 0
        if (v6 != v6) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v55 string = v28
            _ = v55
        }
        _ = v54
    }
    v2 += uint64(v5)
    if (!false) {
        v2 -= (^(v2))
    }
    var v56 uint = 425
    var v57 int32 = int32(v3)
    switch ("a" + ((v28 + "") + v28)) {
    case "a":
        for fi58 := 0; fi58 < 8; fi58++ {
            v57 = (v57 & (-(v57)))
            fmt.Println(v3)
            _ = fi58
        }
    case "hello":
        if (false || false) {
            fmt.Println(v6, v5, v3)
            v57 = (v57 << (((v56 + (v56 ^ v56))) & 63))
            var v59 int32 = (v57 - v57)
            _ = v59
        }
    }
    fmt.Println(v2, v56, v6)
    var v60 uint64 = v6
    for fi61 := 0; fi61 < 15; fi61++ {
        v62 := GenS5{F0: v3, F1: -2147483648}
        var v63 float64 = 3.14159
        _ = fi61
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 4; fi64++ {
        v65 := GenS6{F0: v6, F1: 217}
        _ = fi64
        _ = v65
    }
    for fi66 := 0; fi66 < 13; fi66++ {
        switch int(v57) {
        case -457:
            fmt.Println(v6, fi66)
        }
        for fi67 := 0; fi67 < 14; fi67++ {
            v60 += (v2 >> ((genFn1(v5, 103, v3)) & 63))
            v68 := GenS4{F0: v5, F1: 42}
            if false {
                var v69 float32 = float32((0.0 / (0.0 - (0.1 + -2.25))))
                v70 := GenS4{F0: v5, F1: 2147483647}
                v28 = "\n\t"
                v71 := GenS5{F0: -45, F1: fi66}
                if true {
                    if false {
                        v72 := GenS2{F0: uint32(v6), F1: false}
                        delete(v1, 1)
                        _ = v72
                    } else {
                        var v73 uint32 = 319
                        var v74 string = v28
                        _ = v73
                        _ = v74
                    }
                    if (v3 >= (v3 & 1)) {
                        var v75 int = genFn0(-720)
                        _ = v75
                    } else {
                        if false {
                            fmt.Println(v57)
                            v3 = (v3 ^ (v3 ^ (v3 + (-(v3)))))
                            v76 := GenS7{F0: 285, F1: 127, F2: -217}
                            _ = v76
                        }
                        v77 := GenS8{F0: (v5 - (^(v5)))}
                        switch (((v28 + v28) + v28) + "") {
                        case "\n\t":
                            v78 := GenS9{F0: 42, F1: 110, F2: uint8(v56)}
                            _ = v78
                        case "":
                            switch (v28 + ((v28 + v28) + "foo bar")) {
                            case "a":
                                v5 += (v5 << ((v56) & 63))
                            case "hello":
                                var v79 int16 = 713
                                _ = v79
                            case "foo bar":
                                var v80 int16 = 1
                                if true {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v28, v56, v5)
                                        v80 += v80
                                        var v81 float32 = v69
                                        _ = v81
                                    default:
                                        v57 += (v57 / (((v57 - (v57 >> (((v56 << ((v56) & 63))) & 63)))) | 1))
                                    }
                                }
                                var v82 uint64 = (v2 * v6)
                                _ = v80
                                _ = v82
                            default:
                                v83 := GenS10{F0: 137, F1: (v56 | uint(v6))}
                                _ = v83
                            }
                        case "hello":
                            v1[0] = 2147483648
                            v3 = -100
                            switch 128 {
                            case 2147483647:
                                delete(v1, 196)
                            default:
                                switch 128 {
                                case -1:
                                    fmt.Println(fi66)
                                }
                            }
                        }
                        _ = v77
                    }
                    var v84 float32 = (v69 / ((v69 / ((0.0 / v69) * -2.25)) - 3.14159))
                    v57 += v57
                    _ = v84
                }
                _ = v69
                _ = v70
                _ = v71
            } else {
                if len(v4) > 0 {
                    v4[0] = 1
                }
                if false {
                    var v85 uint64 = (v2 - (v6 * v6))
                    v86 := GenS11{F0: v56}
                    v4 = append(v4, 894)
                    _ = v85
                    _ = v86
                } else {
                    if false {
                        v57 += (^(v57))
                        v5 += (v5 % ((v5) | 1))
                    } else {
                        if false {
                            v4 = append(v4, 2147483647)
                            v57 += (^(v57))
                        } else {
                            v60 = v2
                        }
                        v3 -= 1
                        fmt.Println(v56, fi66, v3)
                    }
                    fmt.Println(fi67, v2)
                }
            }
            _ = fi67
            _ = v68
        }
        _ = fi66
    }
    var v87 bool = (!false)
    for fi88 := 0; fi88 < 1; fi88++ {
        v3 += v3
        v87 = false
        _ = fi88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v6
    _ = v60
    _ = v87
}
