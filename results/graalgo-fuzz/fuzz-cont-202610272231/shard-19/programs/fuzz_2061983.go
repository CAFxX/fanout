package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 string; F1 string; F2 uint32 }
type GenS6 struct { F0 int; F1 int64; F2 int16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint8; F1 int16; F2 int8 }
func genFn0(p0 int16, p1 int8) uint32 {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = 255
        v9 := GenS1{F0: (-(fi7)), F1: (p0 / (((p0 * 522)) | 1))}
        _ = fi7
        _ = v8
        _ = v9
    }
    if (42 == 4294967295) {
        if (-1 >= -399) {
            v10 := GenS1{F0: -129, F1: (p0 >> ((1) & 63))}
            p0 = (p0 * (p0 - 735))
            _ = v10
        }
    }
    _ = p0
    _ = p1
    return 507
}

func main() {
    var v1 float64 = 3.14159
    var v2 int32 = 42
    var v3 int8 = 1
    var v4 uint64 = 4294967295
    var v5 uint16 = 65535
    v6 := make(map[uint32]bool)
    var v11 bool = true
    v12 := GenS2{F0: -32768}
    v13 := []int{int(v5), 32767, 127, -1}
    if (0 < 64) {
        v14 := GenS3{F0: v3, F1: 1}
        _ = v14
    } else {
        v1 = v1
        fmt.Println(v4, v1)
    }
    _ = len(v13)
    v15 := GenS1{F0: 9223372036854775807, F1: -32768}
    v11 = (!false)
    var v16 uint32 = 2
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            switch int(v3) {
            case 57:
                if len(v13) > 0 {
                    v13[0] = 407
                }
            case 146:
                var v19 uint8 = 128
                _ = v19
            case 128:
                _ = len(v6)
            default:
                fmt.Println(v4, v3, v1)
            }
            _ = fi18
        }
        if true {
            var v20 float32 = 3.14159
            v21 := GenS2{F0: 6}
            v16 = genFn0(1000, (v3 & 127))
            _ = v20
            _ = v21
        } else {
            v16 = (v16 - (v16 >> ((18446744073709551615) & 63)))
        }
        if v11 {
            v2 = 207
            v6[uint32(v4)] = (0 <= int16(v3))
        } else {
            for fi22 := 0; fi22 < 4; fi22++ {
                if (!(125 >= 169)) {
                    v1 += -2.25
                }
                _ = fi22
            }
            if (float32(v1) >= (float32((v1 * 1.5)) / float32((v1 + v1)))) {
                v23 := v6[(v16 & (v16 % (((^(v16))) | 1)))]
                v3 = v3
                _ = v23
            }
        }
        _ = fi17
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 int = (fi25 ^ 9223372036854775807)
            _ = fi25
            _ = v26
        }
        _ = len(v13)
        v2 -= v2
        _ = fi24
    }
    v27 := GenS1{F0: int(v16), F1: -1}
    v16 = genFn0(-1, v3)
    v1 = (float64(1.5) / 3.14159)
    var v28 int = 128
    for fi29 := 0; fi29 < 8; fi29++ {
        switch 1 {
        case 1:
            var v30 string = "foo bar"
            _ = v30
        case 32767:
            var v31 uint8 = 2
            _ = v31
        case 591:
            switch ("\n\t" + "a") {
            case "":
                switch (v28 - (v28 * (fi29 / (((v28 ^ (fi29 - (v28 ^ 42)))) | 1)))) {
                case -1:
                    for fi32 := 0; fi32 < 2; fi32++ {
                        var v33 float32 = 0.0
                        fmt.Println(v4, v33)
                        switch "" {
                        case "\n\t":
                            v34 := v6[v16]
                            _ = v34
                        case "hello":
                            if v11 {
                                if false {
                                    v35 := GenS1{F0: 1, F1: -1}
                                    var v36 int = (^(fi32))
                                    _ = len(v13)
                                    _ = v35
                                    _ = v36
                                }
                            }
                        case "":
                            if false {
                                _ = len(v6)
                                v37 := GenS1{F0: (fi32 + 127), F1: 32767}
                                fmt.Println(fi29)
                                var v38 int64 = 482
                                _ = v37
                                _ = v38
                            }
                        }
                        _ = fi32
                        _ = v33
                    }
                case 1:
                    var v39 float64 = v1
                    _ = v39
                case 0:
                    switch 32767 {
                    case 1:
                        var v40 uint = 1
                        _ = v40
                    case 217:
                        var v41 bool = v11
                        _ = v41
                    case 9223372036854775807:
                        fmt.Println(v16, v2)
                    }
                default:
                    for fi42 := 0; fi42 < 8; fi42++ {
                        if false {
                            if ("foo bar" >= (("" + ("hello" + "hello")) + ("\n\t" + ("" + "\n\t")))) {
                                var v43 uint32 = v16
                                v43 += genFn0(-557, (v3 >> ((666) & 63)))
                                v2 -= v2
                                _ = v43
                            } else {
                                if v11 {
                                    if v11 {
                                        if len(v13) >= 2 {
                                            v13 = v13[1:len(v13)]
                                        }
                                        if (v11 || v11) {
                                            var v44 int32 = (v2 ^ int32(v16))
                                            var v45 int64 = 42
                                            if v11 {
                                                if true {
                                                    switch (v28 | 9223372036854775807) {
                                                    case 9223372036854775807:
                                                        var v46 bool = v11
                                                        _ = v46
                                                    case 0:
                                                        v28 -= fi42
                                                    }
                                                } else {
                                                    v44 -= (v44 & (v2 ^ 42))
                                                }
                                                v13 = append(v13, v28)
                                                v4 += (^(v4))
                                            }
                                            _ = v44
                                            _ = v45
                                        }
                                        v2 -= v2
                                    }
                                    fmt.Println(v4)
                                    v47 := GenS4{F0: (v28 % ((fi42) | 1)), F1: 1}
                                    _ = v47
                                }
                                if true {
                                    fmt.Println(v16, v16, fi29)
                                    v5 += (v5 + v5)
                                    var v48 int32 = int32(v5)
                                    var v49 string = ("\n\t" + ("" + ""))
                                    if (v5 > (v5 + 1)) {
                                        var v50 float64 = float64(((0.0 + 3.14159) + -2.25))
                                        if (127 == 170) {
                                            v51 := GenS5{F0: (v49 + v49), F1: v49, F2: (v16 | (v16 * 1))}
                                            v52 := GenS3{F0: (v3 << ((65535) & 63)), F1: (v16 / (((v16 % ((1) | 1))) | 1))}
                                            switch v49 {
                                            case "":
                                                v49 = v49
                                            default:
                                                if v11 {
                                                    var v53 uint = 1
                                                    _ = v53
                                                } else {
                                                    if v11 {
                                                        if true {
                                                            if (!v11) {
                                                                switch v28 {
                                                                case 4294967295:
                                                                    switch (fi42 % (((v28 >> ((356) & 63))) | 1)) {
                                                                    case -694:
                                                                        v49 = "\n\t"
                                                                    case -1:
                                                                        _ = len(v13)
                                                                    case 32768:
                                                                        v2 = (v2 + 0)
                                                                    default:
                                                                        var v54 uint8 = 42
                                                                        _ = v54
                                                                    }
                                                                default:
                                                                    var v55 bool = false
                                                                    _ = v55
                                                                }
                                                                v56 := GenS6{F0: (fi29 / (((v28 + 9223372036854775807)) | 1)), F1: -651, F2: -32768}
                                                                _ = v56
                                                            } else {
                                                                v1 -= float64(1e10)
                                                                v57 := GenS1{F0: 32768, F1: -1}
                                                                _ = v57
                                                            }
                                                            var v58 uint16 = 362
                                                            if (v11 && (!(!false))) {
                                                                if v11 {
                                                                    if (-112 != fi29) {
                                                                        v5 += 65535
                                                                        if (!(1 < (v16 % ((genFn0(-32768, (^(v3)))) | 1)))) {
                                                                            if (v49 >= "foo bar") {
                                                                                var v59 uint32 = v16
                                                                                fmt.Println(fi29)
                                                                                if (v11 && v11) {
                                                                                    v60 := GenS7{F0: (v2 % (((-(v2))) | 1))}
                                                                                    v61 := GenS7{F0: -1}
                                                                                    v6[v59] = v11
                                                                                    _ = v60
                                                                                    _ = v61
                                                                                }
                                                                                fmt.Println(v5, v4)
                                                                                var v62 int64 = 0
                                                                                v49 = "hello"
                                                                                _ = v59
                                                                                _ = v62
                                                                            } else {
                                                                                v3 -= int8(v48)
                                                                                var v63 uint = 255
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                    } else {
                                                                        fmt.Println(fi42, v16, v49)
                                                                        var v64 float32 = -2.25
                                                                        v65 := GenS8{F0: (v48 & (v2 - (^(v48)))), F1: uint64(v58)}
                                                                        switch ("hello" + (v49 + v49)) {
                                                                        case "\n\t":
                                                                            var v66 int8 = v3
                                                                            _ = v66
                                                                        case "":
                                                                            v16 = v16
                                                                            if ((!v11) && ((!true) && (((!(v11 && v11)) || ((v64 * v64) == ((v64 * v64) + 3.14159))) && (0 >= 996)))) {
                                                                                var v67 int16 = -1
                                                                                _ = v67
                                                                            }
                                                                        case "hello":
                                                                            var v68 string = "foo bar"
                                                                            _ = v68
                                                                        default:
                                                                            v2 = -2147483648
                                                                        }
                                                                        var v69 float32 = 1e10
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v69
                                                                    }
                                                                    var v70 float64 = 1e10
                                                                    _ = v70
                                                                } else {
                                                                    fmt.Println(v1)
                                                                    var v71 uint8 = 95
                                                                    fmt.Println(fi29, v16, v4)
                                                                    v28 = (-(fi29))
                                                                    v16 += (^(v16))
                                                                    _ = v71
                                                                }
                                                                v72 := GenS9{F0: 255, F1: 491, F2: -128}
                                                                v6[(v16 ^ 2147483648)] = false
                                                                v13 = append(v13, (v28 ^ (fi29 / ((2147483647) | 1))))
                                                                if (true && v11) {
                                                                    if v11 {
                                                                        var v73 int64 = 673
                                                                        _ = v73
                                                                    }
                                                                    var v74 bool = (1.5 != (float64(float32((v1 / v50))) * v1))
                                                                    if (false && (!((v11 && false) || false))) {
                                                                        v50 = v50
                                                                        var v75 int = 164
                                                                        var v76 uint = 1
                                                                        v75 += (-(fi29))
                                                                        _ = v75
                                                                        _ = v76
                                                                    }
                                                                    _ = v74
                                                                }
                                                                _ = v72
                                                            }
                                                            _ = v58
                                                        } else {
                                                            switch (((v49 + v49) + "") + ("hello" + ((v49 + v49) + ("a" + (v49 + ""))))) {
                                                            case "\n\t":
                                                                v11 = v11
                                                            }
                                                        }
                                                        v1 += (0.1 / v1)
                                                        var v77 uint16 = 42
                                                        _ = v77
                                                    } else {
                                                        v16 -= (v16 % ((genFn0(-328, v3)) | 1))
                                                        switch v49 {
                                                        case "\n\t":
                                                            if len(v13) >= 2 {
                                                                v13 = v13[1:len(v13)]
                                                            }
                                                        }
                                                        v1 = 1e10
                                                    }
                                                }
                                                var v78 int8 = v3
                                                if len(v13) >= 2 {
                                                    v13 = v13[1:len(v13)]
                                                }
                                                fmt.Println(v4)
                                                v11 = v11
                                                _ = v78
                                            }
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v50
                                    }
                                    if (2 < 0) {
                                        if (!(false || true)) {
                                            if len(v13) > 0 {
                                                v13[0] = (^(v28))
                                            }
                                            v28 = 480
                                        }
                                        v6[531] = v11
                                        delete(v6, uint32(v4))
                                    } else {
                                        v16 = genFn0(-413, v3)
                                    }
                                    var v79 float32 = 1e10
                                    _ = v48
                                    _ = v49
                                    _ = v79
                                }
                            }
                            if v11 {
                                var v80 int8 = (v3 ^ -128)
                                _ = v80
                            }
                            v81 := GenS7{F0: v2}
                            var v82 float32 = 0.1
                            v11 = true
                            _ = v81
                            _ = v82
                        } else {
                            var v83 string = ("" + ("\n\t" + "hello"))
                            if (0 != 1) {
                                v1 = v1
                                if ((v2 - v2) <= (v2 | 0)) {
                                    v3 += v3
                                    switch v83 {
                                    case "foo bar":
                                        var v84 uint32 = genFn0(1, (-(v3)))
                                        _ = v84
                                    case "":
                                        fmt.Println(v16, v1, v4)
                                    case "hello":
                                        v85 := v6[42]
                                        _ = v85
                                    default:
                                        if len(v13) > 0 {
                                            v13[0] = (fi42 | v28)
                                        }
                                    }
                                    var v86 uint = 42
                                    v13 = append(v13, 127)
                                    _ = v86
                                }
                            } else {
                                fmt.Println(v28)
                            }
                            _ = v83
                        }
                        var v87 uint32 = 2
                        delete(v6, 2147483647)
                        var v88 int64 = 1
                        v11 = (100 < 42)
                        v1 -= (v1 - 0.1)
                        _ = fi42
                        _ = v87
                        _ = v88
                    }
                }
            case "hello":
                switch "" {
                case "":
                    fmt.Println(v11, v2)
                case "\n\t":
                    for fi89 := 0; fi89 < 8; fi89++ {
                        var v90 uint16 = (v5 * 432)
                        _ = fi89
                        _ = v90
                    }
                default:
                    var v91 string = "a"
                    _ = v91
                }
            case "a":
                var v92 uint16 = v5
                _ = v92
            }
        default:
            v11 = v11
        }
        _ = fi29
    }
    fmt.Println(v28, v5, v28)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
