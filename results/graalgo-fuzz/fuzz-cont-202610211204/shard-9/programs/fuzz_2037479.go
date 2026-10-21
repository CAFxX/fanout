package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int8; F1 bool; F2 uint16 }
type GenS12 struct { F0 uint8; F1 uint16 }
type GenS13 struct { F0 int64; F1 int16; F2 string }
func genFn0(p0 int8, p1 uint16, p2 uint64) bool {
    var v5 int = -644
    fmt.Println(p2, p0, p2)
    fmt.Println(p2, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}
func genFn1() uint {
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            fmt.Println(fi7, fi6)
            fmt.Println(fi7, fi6)
            var v8 uint = 255
            v9 := GenS1{F0: (fi6 << ((0) & 63)), F1: 661, F2: 2147483647}
            _ = fi7
            _ = v8
            _ = v9
        }
        if false {
            v10 := []int{}
            var v11 float64 = (3.14159 / 1.5)
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    v12 := GenS2{F0: 924}
    v13 := []int{473}
    v14 := GenS1{F0: 128, F1: 2, F2: 42}
    _ = v12
    _ = v13
    _ = v14
    return 2
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 255
    var v3 uint = 255
    var v4 int32 = -928
    var v15 int64 = 42
    var v16 bool = (32767 >= 575)
    var v17 uint64 = 2
    switch 507 {
    case 0:
        if (2 == 2) {
            if v16 {
                for fi18 := 0; fi18 < 17; fi18++ {
                    fmt.Println(v17)
                    fmt.Println(v17)
                    switch ("foo bar" + "") {
                    case "a":
                        switch fi18 {
                        case 32767:
                            switch "" {
                            case "a":
                                v19 := GenS2{F0: 65535}
                                _ = v19
                            case "hello":
                                switch 42 {
                                case 1:
                                    for fi20 := 0; fi20 < 20; fi20++ {
                                        fmt.Println(v15, v4)
                                        v4 = (v4 ^ v4)
                                        _ = fi20
                                    }
                                default:
                                    v3 = genFn1()
                                }
                            case "\n\t":
                                if (true || (v16 || v16)) {
                                    for fi21 := 0; fi21 < 7; fi21++ {
                                        v22 := GenS1{F0: (fi18 - -1), F1: 1, F2: v4}
                                        _ = fi21
                                        _ = v22
                                    }
                                } else {
                                    v15 = (v15 % ((v15) | 1))
                                    if v16 {
                                        v23 := GenS2{F0: 355}
                                        switch 18 {
                                        case 1:
                                            fmt.Println(v17, v2, v16)
                                        case -380:
                                            v24 := GenS3{F0: 32767, F1: 9223372036854775807}
                                            _ = v24
                                        case 32768:
                                            for fi25 := 0; fi25 < 16; fi25++ {
                                                fmt.Println(v16)
                                                if (136 == 639) {
                                                    var v26 bool = v16
                                                    _ = v26
                                                }
                                                var v27 uint32 = 528
                                                _ = fi25
                                                _ = v27
                                            }
                                        }
                                        if (-1 < v4) {
                                            if (((3.14159 + 0.1) / -2.25) == (1e10 + (1.5 * ((((1.5 - (3.14159 / ((0.0 + (1.5 + -2.25)) * -2.25))) / 0.1) / 1.5) / 1e10)))) {
                                                v28 := GenS4{F0: uint(v2), F1: 2147483648, F2: -9223372036854775808}
                                                _ = v28
                                            }
                                            for fi29 := 0; fi29 < 1; fi29++ {
                                                var v30 string = "a"
                                                _ = fi29
                                                _ = v30
                                            }
                                        } else {
                                            var v31 bool = v16
                                            var v32 uint16 = 1
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v23
                                    } else {
                                        _ = len(v1)
                                    }
                                    switch (fi18 % ((fi18) | 1)) {
                                    case -464:
                                        v16 = false
                                    case 128:
                                        v15 -= (v15 & -9223372036854775808)
                                    }
                                }
                            default:
                                if v16 {
                                    if (18 == 2) {
                                        v33 := []int{127, fi18, (-(fi18)), fi18}
                                        var v34 uint16 = 2
                                        _ = v33
                                        _ = v34
                                    }
                                    v35 := []int{(fi18 % (((-(fi18))) | 1)), fi18}
                                    for fi36 := 0; fi36 < 9; fi36++ {
                                        if (!v16) {
                                            var v37 uint16 = uint16(v2)
                                            var v38 float32 = (0.0 * (0.0 - 1.5))
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v17 -= (^(v17))
                                            v39 := GenS5{F0: (v16 && (255 <= (v3 / ((18446744073709551615) | 1)))), F1: (v16 && (!v16))}
                                            if v16 {
                                                if v16 {
                                                    var v40 uint16 = 255
                                                    fmt.Println(fi36)
                                                    _ = v40
                                                } else {
                                                    if v16 {
                                                        v41 := GenS6{F0: -1}
                                                        switch "hello" {
                                                        case "":
                                                            if false {
                                                                if v16 {
                                                                    var v42 uint64 = 226
                                                                    if (-533 >= 32767) {
                                                                        v42 += (v17 & (v42 ^ 870))
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v43 := GenS2{F0: 751}
                                                                            _ = v43
                                                                        case "":
                                                                            if (float32(1e10) <= 1e10) {
                                                                                delete(v1, 9223372036854775807)
                                                                                fmt.Println(fi18, v17)
                                                                                v16 = (int32(v15) > -611)
                                                                            }
                                                                        case "hello":
                                                                            v16 = v16
                                                                        default:
                                                                            if (!((^(fi18)) <= 4294967295)) {
                                                                                v42 = 18446744073709551615
                                                                                var v44 uint = 255
                                                                                v45 := GenS2{F0: 157}
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                        var v46 uint64 = (^(v42))
                                                                        _ = v46
                                                                    } else {
                                                                        var v47 uint64 = (v42 >> (((v3 | v3)) & 63))
                                                                        v35 = append(v35, (fi18 << ((genFn1()) & 63)))
                                                                        _ = v47
                                                                    }
                                                                    v48 := GenS4{F0: v3, F1: 66, F2: (v15 & v15)}
                                                                    _ = v42
                                                                    _ = v48
                                                                }
                                                            }
                                                        case "foo bar":
                                                            v3 += 42
                                                        }
                                                        _ = v41
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v2, fi18, fi18)
                                                    _ = len(v1)
                                                } else {
                                                    v3 = (v3 >> ((uint(v2)) & 63))
                                                    var v49 int16 = 352
                                                    v50 := GenS7{F0: uint32(v2)}
                                                    _ = v49
                                                    _ = v50
                                                }
                                            } else {
                                                v51 := GenS3{F0: 255, F1: (v15 * v15)}
                                                if ((-128 > int8(fi36)) && (!((v4 / ((v4) | 1)) < v4))) {
                                                    v16 = v16
                                                    v52 := GenS7{F0: uint32(v3)}
                                                    switch ("" + ("hello" + "hello")) {
                                                    case "hello":
                                                        v53 := GenS8{F0: (((v16 || (1 < 10)) && true) || v16)}
                                                        _ = v53
                                                    }
                                                    _ = v52
                                                }
                                                fmt.Println(v4, v2)
                                                _ = v51
                                            }
                                            _ = v39
                                        }
                                        var v54 int64 = -507
                                        _ = fi36
                                        _ = v54
                                    }
                                    _ = v35
                                }
                            }
                        default:
                            if (32767 != 132) {
                                for fi55 := 0; fi55 < 13; fi55++ {
                                    fmt.Println(v2)
                                    v15 = (v15 + v15)
                                    var v56 uint32 = 2147483648
                                    v56 += (v56 ^ v56)
                                    if (324 == (^(v56))) {
                                        delete(v1, fi55)
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                v57 := []int{494, (fi18 ^ fi18), (fi18 & (fi18 | fi18))}
                                v58 := GenS8{F0: false}
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                    _ = fi18
                }
                fmt.Println(v15, v2)
            }
            v3 -= 0
        } else {
            var v59 float32 = (0.0 + float32(-2.25))
            switch 32768 {
            case 128:
                for fi60 := 0; fi60 < 4; fi60++ {
                    fmt.Println(v16, v17, v15)
                    v61 := GenS5{F0: true, F1: v16}
                    _ = fi60
                    _ = v61
                }
            case 4294967295:
                for fi62 := 0; fi62 < 3; fi62++ {
                    for fi63 := 0; fi63 < 17; fi63++ {
                        v64 := GenS9{F0: 1, F1: v2, F2: v17}
                        v65 := GenS10{F0: "foo bar"}
                        v66 := GenS11{F0: -1, F1: (1.5 < v59), F2: 696}
                        _ = fi63
                        _ = v64
                        _ = v65
                        _ = v66
                    }
                    _ = fi62
                }
            default:
                fmt.Println(v2, v15)
            }
            _ = v59
        }
    case 127:
        for fi67 := 0; fi67 < 5; fi67++ {
            v68 := GenS12{F0: uint8(v3), F1: 0}
            _ = fi67
            _ = v68
        }
    }
    v1[0] = ("" + ("\n\t" + ("a" + "foo bar")))
    fmt.Println(v16, v15)
    v17 += (^(v17))
    v3 = 18446744073709551615
    v69 := []int{128, -9223372036854775808, 32767, 42}
    var v70 float64 = 1.5
    v71 := GenS7{F0: 1}
    switch (("a" + "foo bar") + "") {
    case "\n\t":
        var v72 uint32 = 30
        _ = v72
    case "foo bar":
        if len(v69) >= 2 {
            v69 = v69[1:len(v69)]
        }
    case "":
        v3 = (v3 / ((65535) | 1))
    default:
        switch 128 {
        case 128:
            v1[int(v17)] = ((("\n\t" + "\n\t") + "") + ("foo bar" + "\n\t"))
        }
    }
    v73 := GenS13{F0: v15, F1: 42, F2: ""}
    v15 += (^(v15))
    for fi74 := 0; fi74 < 14; fi74++ {
        var v75 uint = uint(v17)
        _ = fi74
        _ = v75
    }
    switch "\n\t" {
    case "foo bar":
        if (4294967295 != 42) {
            _ = len(v69)
            v69 = append(v69, 9223372036854775807)
        } else {
            var v76 uint64 = v17
            _ = v76
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v69:", len(v69), v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v73
}
