package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int8) int32 {
    switch "foo bar" {
    case "foo bar":
        switch 32768 {
        case -9223372036854775808:
            p0 = p0
        }
    case "":
        var v7 int = 42
        _ = v7
    }
    var v8 bool = (255 == 1)
    var v9 uint16 = 305
    var v10 int64 = -1
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return -206
}
func genFn1(p0 int32, p1 int) uint8 {
    var v11 float32 = 0.0
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS1{F0: 127, F1: 32767, F2: 4294967295}
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 42, F1: p1, F2: 148}
    v15 := []int{(p1 + p1)}
    var v16 string = "hello"
    _ = p0
    _ = p1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 int64 = -897
    var v2 int = -868
    v3 := []uint32{42, 2}
    var v4 uint16 = 42
    var v5 uint32 = 65535
    v6 := []int{(v2 << ((379) & 63)), v2}
    var v17 float32 = 0.0
    v4 += (v4 % (((^(v4))) | 1))
    if false {
        v18 := GenS1{F0: 1, F1: 4294967295, F2: (v5 + v5)}
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v4 = (v4 ^ v4)
        var v20 float64 = ((0.1 - -2.25) + 3.14159)
        var v21 uint64 = 0
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 61, F1: 127, F2: v5}
    if false {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 bool = false
            for fi25 := 0; fi25 < 19; fi25++ {
                v2 = -2147483648
                var v26 bool = (!(802 >= 18446744073709551615))
                var v27 uint32 = v5
                v2 = (-(fi25))
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi23
            _ = v24
        }
        if ((v2 ^ v2) == int(v4)) {
            var v28 uint16 = v4
            fmt.Println(v5)
            v4 = (^(v28))
            _ = v28
        }
        v4 = v4
    } else {
        for fi29 := 0; fi29 < 18; fi29++ {
            var v30 uint8 = 215
            _ = fi29
            _ = v30
        }
    }
    v31 := GenS1{F0: 1, F1: 9223372036854775807, F2: 2147483648}
    v32 := GenS2{F0: 40, F1: -594}
    v17 -= 0.0
    switch (v2 * 0) {
    case -2147483648:
        var v33 uint16 = 32767
        _ = v33
    case 128:
        var v34 uint32 = v5
        _ = v34
    case 42:
        switch v2 {
        case 547:
            v35 := GenS2{F0: 458, F1: 1}
            _ = v35
        case 42:
            v5 += v5
        default:
            for fi36 := 0; fi36 < 2; fi36++ {
                v2 += (v2 >> ((2) & 63))
                _ = fi36
            }
        }
    default:
        switch int(v4) {
        case 494:
            for fi37 := 0; fi37 < 6; fi37++ {
                _ = len(v3)
                v3 = append(v3, (^(v5)))
                var v38 uint32 = v5
                _ = fi37
                _ = v38
            }
        case 1:
            var v39 bool = (float64(v17) == 0.1)
            _ = v39
        case 380:
            var v40 uint = 4294967295
            _ = v40
        }
    }
    var v41 uint = 255
    switch "a" {
    case "\n\t":
        v4 = v4
    case "foo bar":
        switch v2 {
        case -9223372036854775808:
            for fi42 := 0; fi42 < 10; fi42++ {
                switch -1 {
                case 127:
                    if false {
                        var v43 bool = (!true)
                        _ = v43
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        for fi44 := 0; fi44 < 14; fi44++ {
                            if (-42 <= 42) {
                                switch 1 {
                                case -2147483648:
                                    switch "hello" {
                                    case "a":
                                        fmt.Println(v41, v17, v5)
                                    case "hello":
                                        switch v2 {
                                        case -9223372036854775808:
                                            fmt.Println(v1, v2, v17)
                                        case 566:
                                            v1 += int64(v5)
                                        }
                                        v1 = -1
                                        switch (fi44 & 4294967295) {
                                        case 1:
                                            v45 := GenS3{F0: ((0 < v5) && (true || ("\n\t" == (("a" + "\n\t") + "foo bar")))), F1: v5}
                                            _ = v45
                                        case 579:
                                            var v46 uint8 = 0
                                            _ = v46
                                        case 32767:
                                            v4 = v4
                                        default:
                                            v47 := GenS2{F0: 1, F1: genFn0(0)}
                                            _ = v47
                                        }
                                    default:
                                        fmt.Println(v17)
                                    }
                                case 973:
                                    if len(v3) > 0 {
                                        v3[0] = (^(v5))
                                    }
                                }
                                fmt.Println(v17)
                                v4 += 0
                                var v48 uint32 = v5
                                _ = v48
                            }
                            var v49 int = int(v5)
                            fmt.Println(v17)
                            _ = fi44
                            _ = v49
                        }
                    }
                case 576:
                    var v50 uint8 = 0
                    _ = v50
                default:
                    v51 := GenS2{F0: 32767, F1: 0}
                    _ = v51
                }
                v17 += float32((-2.25 + 1.5))
                _ = fi42
            }
        default:
            v2 += (v2 % ((v2) | 1))
        }
    default:
        var v52 uint16 = v4
        _ = v52
    }
    switch 32767 {
    case 688:
        v53 := GenS1{F0: 128, F1: v2, F2: v5}
        _ = v53
    case -219:
        v54 := GenS4{F0: 192}
        _ = v54
    case -2147483648:
        for fi55 := 0; fi55 < 13; fi55++ {
            var v56 int64 = (v1 - (^(v1)))
            for fi57 := 0; fi57 < 13; fi57++ {
                v41 = 181
                v17 = ((1.5 + ((float32((float64(3.14159) / 1.0)) + v17) - v17)) / v17)
                _ = fi57
            }
            fmt.Println(v2, v41)
            _ = fi55
            _ = v56
        }
    }
    for fi58 := 0; fi58 < 3; fi58++ {
        if (42 < (-(v2))) {
            fmt.Println(v41)
            _ = len(v6)
        } else {
            v1 = v1
            var v59 float64 = ((-2.25 - (0.0 + float64(v17))) + 3.14159)
            _ = v59
        }
        if false {
            var v60 uint16 = 65535
            switch "" {
            case "foo bar":
                _ = len(v3)
            case "\n\t":
                v61 := GenS2{F0: 32767, F1: 42}
                _ = v61
            case "a":
                if (!false) {
                    v5 = v5
                    for fi62 := 0; fi62 < 1; fi62++ {
                        if false {
                            var v63 uint8 = genFn1(2147483647, v2)
                            if false {
                                if true {
                                    var v64 uint64 = 1
                                    v1 = 0
                                    var v65 int32 = 0
                                    _ = v64
                                    _ = v65
                                } else {
                                    if (uint32(v41) <= (v5 >> (((^(v41))) & 63))) {
                                        v6 = append(v6, fi62)
                                        v66 := GenS1{F0: (v63 % (((v63 | 0)) | 1)), F1: v2, F2: (v5 - 229)}
                                        var v67 int64 = 1
                                        fmt.Println(v60, v2)
                                        _ = v66
                                        _ = v67
                                    }
                                    var v68 int16 = 42
                                    v69 := GenS4{F0: (v63 % (((v63 + (v63 / ((v63) | 1)))) | 1))}
                                    _ = v68
                                    _ = v69
                                }
                                if (!true) {
                                    var v70 int32 = -51
                                    var v71 float64 = 1e10
                                    _ = v70
                                    _ = v71
                                }
                            } else {
                                v17 -= ((0.1 + 1.5) * v17)
                                if false {
                                    v72 := GenS4{F0: 153}
                                    _ = v72
                                } else {
                                    fmt.Println(v63, v5)
                                    var v73 uint8 = v63
                                    _ = v73
                                }
                            }
                            _ = v63
                        }
                        _ = fi62
                    }
                }
            }
            for fi74 := 0; fi74 < 16; fi74++ {
                v41 = (v41 ^ (v41 ^ (v41 + v41)))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v75 int16 = -110
                _ = fi74
                _ = v75
            }
            _ = v60
        }
        _ = fi58
    }
    switch "foo bar" {
    case "hello":
        var v76 uint8 = 42
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
