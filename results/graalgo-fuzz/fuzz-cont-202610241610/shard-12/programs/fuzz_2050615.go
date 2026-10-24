package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int64; F1 int64; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint8, p1 uint) uint16 {
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = (^(p0))
        v7 := []int{1, -151}
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint = 1
            if false {
                switch "\n\t" {
                case "a":
                    v9 = v9
                case "hello":
                    p0 -= 42
                default:
                    if ((p0 == 0) || (true && (false && false))) {
                        v9 = (v9 % ((v9) | 1))
                    }
                }
                v10 := GenS1{F0: 0, F1: 929}
                var v11 uint16 = 42
                _ = v10
                _ = v11
            }
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0, p1)
    var v12 int = 2147483647
    v13 := GenS1{F0: 1, F1: 0}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 883
}
func genFn1(p0 uint64) uint64 {
    v14 := []int{1, -9223372036854775808, 42}
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint8 = 255
        for fi17 := 0; fi17 < 10; fi17++ {
            fmt.Println(fi15)
            var v18 int64 = 9223372036854775807
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = v14
    return p0
}

func main() {
    var v1 uint64 = 997
    var v2 int64 = -1
    v3 := make(map[uint32]bool)
    var v4 int16 = -1
    var v5 uint16 = 817
    v3[42] = (3.14159 <= 3.14159)
    var v19 float32 = -2.25
    for fi20 := 0; fi20 < 2; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := []int{1, fi21, fi21}
            if ((-178 <= 2147483647) || false) {
                v4 = (v4 * (v4 + int16(fi20)))
                if (false || false) {
                    fmt.Println(v4)
                    fmt.Println(v19, v2, v2)
                }
                v2 += v2
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 string = "a"
        if (false || (!true)) {
            var v25 int16 = 32767
            if true {
                for fi26 := 0; fi26 < 11; fi26++ {
                    var v27 int32 = -1
                    _ = fi26
                    _ = v27
                }
            } else {
                var v28 uint64 = genFn1((v1 | 524))
                if (!false) {
                    switch (fi23 ^ -665) {
                    case -800:
                        for fi29 := 0; fi29 < 11; fi29++ {
                            var v30 uint32 = 65535
                            v31 := GenS1{F0: 0, F1: (v30 >> ((0) & 63))}
                            var v32 int16 = (v25 / (((v4 - -32768)) | 1))
                            _ = fi29
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    case -956:
                        var v33 float32 = 0.1
                        _ = v33
                    default:
                        var v34 uint = 781
                        _ = v34
                    }
                }
                v19 = float32(0.0)
                _ = v28
            }
            v3[42] = (209 <= 127)
            _ = v25
        } else {
            var v35 int8 = 42
            if (false || (true || (v4 >= v4))) {
                for fi36 := 0; fi36 < 13; fi36++ {
                    var v37 float32 = (v19 * (0.0 + v19))
                    v37 -= v37
                    _ = fi36
                    _ = v37
                }
            }
            _ = v35
        }
        if false {
            if (!false) {
                switch fi23 {
                case 32768:
                    v4 = (v4 * v4)
                default:
                    var v38 int64 = v2
                    _ = v38
                }
            } else {
                var v39 string = "\n\t"
                delete(v3, 656)
                _ = v39
            }
            v1 += (v1 * v1)
        } else {
            if true {
                if false {
                    v19 = v19
                    v40 := GenS1{F0: 481, F1: 0}
                    _ = v40
                }
                for fi41 := 0; fi41 < 10; fi41++ {
                    if (v19 == v19) {
                        if (2147483648 > 240) {
                            v24 = v24
                        } else {
                            var v42 uint = 18446744073709551615
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi23)
                            }
                            if (false && false) {
                                var v43 uint64 = uint64(v42)
                                _ = v43
                            }
                            _ = v42
                        }
                        v44 := GenS2{F0: (-(v2)), F1: v2, F2: -548}
                        if (18446744073709551615 >= 835) {
                            v45 := []int{127, fi41, (-(fi23)), (-(fi23))}
                            _ = v45
                        } else {
                            v46 := GenS3{F0: (fi41 << ((0) & 63))}
                            _ = v46
                        }
                        _ = v44
                    }
                    v1 = v1
                    _ = fi41
                }
                v47 := GenS4{F0: (-(v4))}
                _ = v47
            } else {
                if (fi23 > (fi23 & (^(fi23)))) {
                    if (255 == v5) {
                        var v48 string = v24
                        _ = v48
                    } else {
                        var v49 int32 = -307
                        _ = v49
                    }
                    v24 = v24
                    for fi50 := 0; fi50 < 8; fi50++ {
                        var v51 uint8 = 255
                        _ = len(v3)
                        _ = fi50
                        _ = v51
                    }
                } else {
                    fmt.Println(v2)
                }
            }
        }
        _ = fi23
        _ = v24
    }
    if (!(true && (100 > 103))) {
        v52 := GenS1{F0: v5, F1: 65535}
        switch 4 {
        case 42:
            switch "\n\t" {
            case "\n\t":
                var v53 int32 = 0
                _ = v53
            case "foo bar":
                var v54 string = "\n\t"
                _ = v54
            case "hello":
                v19 = (v19 + (v19 / v19))
            }
        case 128:
            v2 = v2
        default:
            v5 = v5
        }
        var v55 int = 2147483647
        _ = v52
        _ = v55
    } else {
        for fi56 := 0; fi56 < 10; fi56++ {
            v2 = -705
            v57 := GenS2{F0: v2, F1: (v2 - (v2 - v2)), F2: (v2 & -9223372036854775808)}
            _ = fi56
            _ = v57
        }
        var v58 uint8 = 128
        _ = v58
    }
    switch "a" {
    case "a":
        switch 919 {
        case -1:
            for fi59 := 0; fi59 < 20; fi59++ {
                var v60 uint16 = v5
                var v61 uint32 = 460
                v61 += v61
                _ = fi59
                _ = v60
                _ = v61
            }
        case -396:
            v62 := GenS1{F0: 2, F1: 2}
            _ = v62
        case 32767:
            v2 = (v2 ^ v2)
        default:
            delete(v3, 42)
        }
    default:
        v63 := []int{-952, 977}
        _ = v63
    }
    v5 = genFn0(13, 42)
    v64 := []int{}
    var v65 uint32 = 0
    switch "\n\t" {
    case "foo bar":
        var v66 uint64 = uint64(v65)
        _ = v66
    case "hello":
        var v67 float32 = (v19 + ((1.5 * v19) / (((v19 * v19) / v19) + v19)))
        _ = v67
    case "":
        fmt.Println(v65)
    }
    if false {
        for fi68 := 0; fi68 < 2; fi68++ {
            v69 := GenS5{F0: 255}
            _ = len(v64)
            _ = fi68
            _ = v69
        }
        if len(v64) > 0 {
            v64[0] = -524
        }
    }
    var v70 string = "foo bar"
    for fi71 := 0; fi71 < 10; fi71++ {
        switch v70 {
        case "\n\t":
            v72 := GenS2{F0: -9223372036854775808, F1: v2, F2: -1}
            _ = v72
        }
        for fi73 := 0; fi73 < 16; fi73++ {
            _ = len(v64)
            _ = fi73
        }
        _ = fi71
    }
    fmt.Println(v65)
    if true {
        var v74 uint16 = genFn0(uint8(v5), 18446744073709551615)
        var v75 uint = 1
        _ = v74
        _ = v75
    } else {
        v76 := GenS2{F0: (v2 >> ((1) & 63)), F1: (v2 % ((v2) | 1)), F2: (v2 + (v2 ^ 0))}
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
    _ = v70
}
