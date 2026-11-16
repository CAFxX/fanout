package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8; F1 string; F2 bool }
type GenS5 struct { F0 int8; F1 int8; F2 int }
type GenS6 struct { F0 int8; F1 uint32 }
func genFn0(p0 int8, p1 uint32, p2 int8) uint32 {
    v6 := GenS1{F0: 0, F1: 88}
    if true {
        var v7 int16 = 1
        _ = v7
    } else {
        if (((189 <= 18446744073709551615) || (1 >= -2147483648)) && false) {
            for fi8 := 0; fi8 < 5; fi8++ {
                p2 = p2
                p1 = 4294967295
                switch 2147483647 {
                case 826:
                    p0 -= 127
                case 42:
                    var v9 uint16 = 2
                    _ = v9
                case 32768:
                    p1 = 1
                default:
                    var v10 int64 = 9223372036854775807
                    _ = v10
                }
                _ = fi8
            }
            fmt.Println(p2, p1, p1)
            var v11 uint8 = 0
            _ = v11
        }
        fmt.Println(p0, p2, p2)
        var v12 int8 = (p0 ^ 42)
        _ = v12
    }
    if (1 <= 4294967295) {
        var v13 string = "\n\t"
        _ = v13
    } else {
        var v14 string = "foo bar"
        p2 = p0
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p1
}
func genFn1(p0 uint64, p1 uint16) uint8 {
    p1 = (p1 >> ((65535) & 63))
    v15 := GenS2{F0: 1000, F1: 65535}
    _ = p0
    _ = p1
    _ = v15
    return 21
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = (0.0 * 1.5)
    var v3 int = 824
    var v4 int = (v3 & (v3 >> ((255) & 63)))
    v5 := make(map[uint32]uint32)
    v5[genFn0(127, 1, 1)] = 2
    var v16 uint = 2
    for fi17 := 0; fi17 < 16; fi17++ {
        if (42 >= 1) {
            for fi18 := 0; fi18 < 5; fi18++ {
                var v19 int64 = 304
                var v20 bool = false
                if (v1 <= v1) {
                    var v21 uint = (v16 + v16)
                    _ = len(v5)
                    if (2147483647 == 2147483647) {
                        v16 = v21
                        v3 = (v3 >> ((v21) & 63))
                    }
                    _ = v21
                } else {
                    if false {
                        if true {
                            v4 = -881
                            var v22 int = v4
                            v3 = fi17
                            _ = v22
                        }
                    } else {
                        v4 += v4
                    }
                    if ((v2 <= v2) || false) {
                        var v23 int32 = -2147483648
                        _ = v23
                    }
                    v24 := GenS1{F0: v16, F1: 127}
                    _ = v24
                }
                _ = fi18
                _ = v19
                _ = v20
            }
        } else {
            for fi25 := 0; fi25 < 11; fi25++ {
                if false {
                    var v26 uint8 = 136
                    switch (v1 + (v1 + v1)) {
                    case "":
                        v26 += (v26 * (v26 >> (((v16 << ((v16) & 63))) & 63)))
                    case "foo bar":
                        fmt.Println(v2, v3)
                    case "\n\t":
                        var v27 int8 = -80
                        _ = v27
                    }
                    _ = v26
                } else {
                    if true {
                        if false {
                            v28 := GenS1{F0: (^(v16)), F1: 128}
                            _ = v28
                        } else {
                            fmt.Println(v2)
                        }
                        v1 = ("a" + ("\n\t" + v1))
                    }
                    switch v3 {
                    case 32768:
                        v29 := GenS2{F0: 1, F1: 42}
                        _ = v29
                    case 4294967295:
                        if false {
                            var v30 float32 = 1.5
                            v31 := []int{}
                            _ = v30
                            _ = v31
                        }
                    case 32767:
                        var v32 uint8 = 255
                        _ = v32
                    }
                }
                fmt.Println(v1, v3, fi17)
                _ = fi25
            }
            _ = len(v5)
        }
        delete(v5, 2)
        _ = fi17
    }
    for fi33 := 0; fi33 < 18; fi33++ {
        switch 32767 {
        case -747:
            if (128 < 127) {
                switch (v3 << (((v16 | (v16 ^ v16))) & 63)) {
                case 497:
                    if (false && true) {
                        v2 -= v2
                        v16 += (v16 % ((v16) | 1))
                    } else {
                        if false {
                            v34 := GenS3{F0: (^(v3))}
                            for fi35 := 0; fi35 < 9; fi35++ {
                                v36 := []int{fi35, v3, (v4 % ((v4) | 1))}
                                v37 := GenS1{F0: v16, F1: 117}
                                _ = len(v36)
                                _ = fi35
                                _ = v36
                                _ = v37
                            }
                            switch v1 {
                            case "foo bar":
                                var v38 uint32 = 2
                                _ = v38
                            case "":
                                v39 := GenS4{F0: 42, F1: v1, F2: true}
                                _ = v39
                            case "a":
                                for fi40 := 0; fi40 < 10; fi40++ {
                                    var v41 int = (fi33 ^ 127)
                                    _ = fi40
                                    _ = v41
                                }
                            default:
                                var v42 uint = (^(v16))
                                _ = v42
                            }
                            _ = v34
                        } else {
                            v4 -= (v3 / ((v3) | 1))
                        }
                        var v43 uint32 = 2147483647
                        var v44 int32 = 2147483647
                        _ = v43
                        _ = v44
                    }
                case -518:
                    fmt.Println(fi33, v3)
                case 32767:
                    v45 := []int{fi33, (^(v4))}
                    _ = v45
                default:
                    switch v1 {
                    case "hello":
                        switch (v4 << (((v16 % ((0) | 1))) & 63)) {
                        case 4294967295:
                            var v46 float64 = 0.1
                            _ = v46
                        case 42:
                            var v47 float32 = v2
                            _ = v47
                        case 2147483647:
                            for fi48 := 0; fi48 < 10; fi48++ {
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    delete(v5, uint32(v16))
                                }
                                var v49 string = "a"
                                _ = fi48
                                _ = v49
                            }
                        default:
                            v50 := GenS1{F0: (^(v16)), F1: 0}
                            _ = v50
                        }
                    default:
                        switch v1 {
                        case "\n\t":
                            v51 := GenS2{F0: 42, F1: v16}
                            _ = v51
                        }
                    }
                }
                var v52 uint = (v16 ^ (v16 << ((1) & 63)))
                _ = v52
            }
        case 4294967295:
            fmt.Println(v2, v4)
        case 128:
            if false {
                var v53 string = ((v1 + (v1 + v1)) + (v1 + v1))
                var v54 uint64 = 42
                v4 = (^(v3))
                _ = v53
                _ = v54
            } else {
                v2 -= (0.1 / 1.0)
                var v55 int = fi33
                _ = v55
            }
        }
        _ = fi33
    }
    switch (v3 * (v4 / ((v4) | 1))) {
    case 767:
        fmt.Println(v16, v16)
    case 1:
        var v56 uint64 = 42
        _ = v56
    case -9223372036854775808:
        v57 := []int{}
        _ = v57
    default:
        for fi58 := 0; fi58 < 3; fi58++ {
            fmt.Println(v3, v3)
            _ = fi58
        }
    }
    v3 += v3
    if (-226 >= -297) {
        fmt.Println(v3)
    }
    v3 = (v4 << (((^(v16))) & 63))
    switch v4 {
    case -942:
        v5[95] = 1
    case 351:
        v2 += 1.5
    default:
        if (!true) {
            v1 = v1
            switch v1 {
            case "\n\t":
                var v59 uint64 = 1
                _ = v59
            case "":
                switch 32767 {
                case 42:
                    switch v1 {
                    case "hello":
                        v60 := []int{}
                        _ = v60
                    case "foo bar":
                        switch "a" {
                        case "hello":
                            fmt.Println(v16, v2)
                        case "foo bar":
                            var v61 uint32 = 2147483647
                            _ = v61
                        case "a":
                            switch v1 {
                            case "foo bar":
                                v62 := GenS3{F0: v3}
                                _ = v62
                            case "a":
                                if (1 != 625) {
                                    v63 := GenS1{F0: v16, F1: genFn1(42, 2)}
                                    var v64 string = v1
                                    _ = v63
                                    _ = v64
                                } else {
                                    v65 := v5[2147483647]
                                    for fi66 := 0; fi66 < 10; fi66++ {
                                        v16 = v16
                                        _ = fi66
                                    }
                                    _ = v65
                                }
                            case "hello":
                                if false {
                                    v67 := GenS1{F0: (v16 | v16), F1: 128}
                                    if (187 >= v16) {
                                        v68 := GenS4{F0: 127, F1: "\n\t", F2: (1 > 0)}
                                        v3 += (v4 + v4)
                                        _ = v68
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            for fi69 := 0; fi69 < 9; fi69++ {
                                                var v70 uint = (v16 | (v16 + (v16 * 4294967295)))
                                                if true {
                                                    var v71 int64 = -254
                                                    if (true && true) {
                                                        switch (fi69 / ((v4) | 1)) {
                                                        case 127:
                                                            v72 := GenS3{F0: 32768}
                                                            _ = v72
                                                        case 584:
                                                            v2 -= 3.14159
                                                        default:
                                                            v73 := GenS2{F0: 0, F1: 65535}
                                                            _ = v73
                                                        }
                                                    } else {
                                                        v4 += (fi69 + v3)
                                                        for fi74 := 0; fi74 < 12; fi74++ {
                                                            fmt.Println(fi74)
                                                            _ = len(v5)
                                                            _ = fi74
                                                        }
                                                    }
                                                    fmt.Println(v70, v71, v3)
                                                    _ = v71
                                                } else {
                                                    var v75 int16 = -928
                                                    fmt.Println(v70, v75, v75)
                                                    _ = v75
                                                }
                                                _ = fi69
                                                _ = v70
                                            }
                                        }
                                    }
                                    _ = v67
                                } else {
                                    v76 := []int{-2147483648, v4, 2147483647}
                                    switch v4 {
                                    case 758:
                                        switch "a" {
                                        case "":
                                            v77 := GenS4{F0: genFn1(29, 534), F1: v1, F2: true}
                                            _ = v77
                                        }
                                    case 128:
                                        v4 = (v3 % ((32767) | 1))
                                    case 2147483647:
                                        switch v1 {
                                        case "\n\t":
                                            if (2147483647 > 2147483647) {
                                                switch 285 {
                                                case 127:
                                                    var v78 int8 = -128
                                                    _ = v78
                                                default:
                                                    v16 += v16
                                                }
                                                var v79 bool = true
                                                _ = v79
                                            } else {
                                                _ = len(v76)
                                                for fi80 := 0; fi80 < 15; fi80++ {
                                                    v81 := v5[0]
                                                    for fi82 := 0; fi82 < 15; fi82++ {
                                                        fmt.Println(v2, fi80, fi80)
                                                        var v83 string = (v1 + ("hello" + v1))
                                                        _ = fi82
                                                        _ = v83
                                                    }
                                                    if (-325 == 0) {
                                                        v16 += v16
                                                        v4 = (v4 / ((0) | 1))
                                                    }
                                                    _ = fi80
                                                    _ = v81
                                                }
                                                var v84 bool = false
                                                _ = v84
                                            }
                                        case "":
                                            fmt.Println(v16, v2, v3)
                                        case "foo bar":
                                            fmt.Println(v3, v2)
                                        }
                                    default:
                                        v16 += (v16 % (((v16 / ((65535) | 1))) | 1))
                                    }
                                    _ = v76
                                }
                            }
                        }
                    default:
                        v16 += 4294967295
                    }
                case 1:
                    fmt.Println(v16)
                default:
                    fmt.Println(v3, v2)
                }
            }
        } else {
            v4 = -997
        }
    }
    v85 := GenS1{F0: v16, F1: 128}
    fmt.Println(v4)
    v86 := []int{-1, 479}
    v87 := GenS5{F0: 42, F1: 66, F2: v4}
    var v88 int16 = -32768
    v89 := GenS5{F0: 0, F1: 100, F2: v3}
    v90 := GenS6{F0: -33, F1: 353}
    v91 := GenS6{F0: -107, F1: 65535}
    fmt.Println(v4, v88, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v86:", len(v86), v86)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v85
    _ = v86
    _ = v87
    _ = v88
    _ = v89
    _ = v90
    _ = v91
}
