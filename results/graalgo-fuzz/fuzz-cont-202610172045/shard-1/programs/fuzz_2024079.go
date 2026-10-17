package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 int16 }
type GenS3 struct { F0 bool; F1 bool; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int32; F2 bool }
type GenS8 struct { F0 string; F1 int32; F2 int8 }

func main() {
    v1 := []uint32{}
    var v2 float32 = 3.14159
    var v3 int16 = 1
    v4 := make(map[int]bool)
    var v5 uint8 = 127
    if false {
        switch (("\n\t" + "") + ("\n\t" + "hello")) {
        case "":
            v5 = 127
        case "a":
            v5 += (^(v5))
        case "hello":
            var v6 string = "\n\t"
            _ = v6
        }
    }
    v5 = v5
    if true {
        if ((v5 + (v5 * v5)) < (v5 - (v5 & v5))) {
            v2 += 0.1
        } else {
            for fi7 := 0; fi7 < 11; fi7++ {
                if false {
                    if len(v1) > 0 {
                        v1[0] = 255
                    }
                }
                var v8 float64 = 3.14159
                _ = fi7
                _ = v8
            }
            v4[32767] = false
        }
    }
    v5 = (v5 - v5)
    if (false && (false || true)) {
        for fi9 := 0; fi9 < 13; fi9++ {
            delete(v4, fi9)
            for fi10 := 0; fi10 < 10; fi10++ {
                fmt.Println(v2)
                v3 = (v3 * 1000)
                var v11 int32 = 42
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        var v12 int8 = -51
        _ = v12
    } else {
        var v13 uint64 = 42
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 0, F1: 42, F2: 0}
            var v16 int8 = 42
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    }
    if false {
        fmt.Println(v2, v2, v2)
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            switch (fi17 << ((339) & 63)) {
            case -2147483648:
                var v18 int32 = 0
                _ = v18
            case 435:
                if true {
                    v19 := GenS2{F0: 65535, F1: (fi17 << ((146) & 63)), F2: -859}
                    _ = v19
                }
            case -131:
                if false {
                    var v20 int8 = 119
                    var v21 string = ("foo bar" + "\n\t")
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(v3)
                }
            }
        case "a":
            fmt.Println(v3)
        default:
            for fi22 := 0; fi22 < 14; fi22++ {
                switch (("foo bar" + "\n\t") + "") {
                case "a":
                    switch -327 {
                    case -756:
                        var v23 uint32 = 2
                        _ = v23
                    }
                case "":
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = len(v1)
                    } else {
                        switch "" {
                        case "foo bar":
                            switch fi22 {
                            case 4294967295:
                                if (32767 >= 32767) {
                                    if (9223372036854775807 != -1) {
                                        v2 = v2
                                    } else {
                                        v2 += ((1e10 * v2) + 3.14159)
                                        if false {
                                            delete(v4, (fi17 | 4294967295))
                                            if (!(0 > 83)) {
                                                v24 := GenS3{F0: (false && (11 == -2147483648)), F1: true, F2: (-128 >= 42)}
                                                _ = v24
                                            }
                                            var v25 int64 = -158
                                            if len(v1) > 0 {
                                                v1[0] = 42
                                            }
                                            _ = v25
                                        } else {
                                            var v26 int16 = v3
                                            var v27 int32 = -1
                                            fmt.Println(v2, v27)
                                            v3 = -463
                                            var v28 float32 = (3.14159 / v2)
                                            _ = v26
                                            _ = v27
                                            _ = v28
                                        }
                                        if ((fi22 * fi22) > -2147483648) {
                                            fmt.Println(v5, v2, fi22)
                                            fmt.Println(fi22, v5)
                                            var v29 int64 = -574
                                            _ = v29
                                        } else {
                                            v3 += (^(v3))
                                        }
                                    }
                                }
                            }
                        case "hello":
                            fmt.Println(v3, v3)
                        case "":
                            switch (fi17 & (fi22 ^ (fi17 | 2147483647))) {
                            case 127:
                                if false {
                                    switch 0 {
                                    case 9223372036854775807:
                                        var v30 int32 = 0
                                        _ = v30
                                    }
                                    switch (("" + "a") + "\n\t") {
                                    case "hello":
                                        fmt.Println(v3, v5, v3)
                                    default:
                                        var v31 uint64 = 2
                                        _ = v31
                                    }
                                }
                                v32 := GenS1{F0: 4294967295, F1: 128, F2: 0}
                                _ = v32
                            }
                        default:
                            var v33 int16 = (v3 + (-(v3)))
                            _ = v33
                        }
                        switch "" {
                        case "hello":
                            var v34 int16 = 42
                            _ = v34
                        case "":
                            v2 = v2
                        default:
                            v5 += v5
                        }
                        v35 := GenS1{F0: uint32(v5), F1: fi17, F2: 33}
                        _ = v35
                    }
                }
                switch fi17 {
                case -624:
                    switch (fi17 / ((fi22) | 1)) {
                    case 443:
                        v3 = int16(v5)
                    case 0:
                        var v36 int8 = 1
                        _ = v36
                    default:
                        var v37 int = int(v5)
                        _ = v37
                    }
                    var v38 uint = 424
                    _ = v38
                case 9223372036854775807:
                    v5 += (^(v5))
                case 128:
                    if true {
                        v5 += (v5 * (v5 - (v5 >> ((18446744073709551615) & 63))))
                        if false {
                            v3 = -32768
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v2 -= v2
                                v5 = 25
                            case "hello":
                                switch fi22 {
                                case 32767:
                                    v39 := GenS2{F0: 2, F1: (fi22 / ((fi17) | 1)), F2: v3}
                                    _ = v39
                                case -9223372036854775808:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 4294967295:
                                    v2 = float32(1e10)
                                default:
                                    v40 := GenS4{F0: 65535}
                                    _ = v40
                                }
                            }
                        }
                        v1 = append(v1, 1)
                    }
                }
                v3 += v3
                _ = fi22
            }
        }
        fmt.Println(fi17)
        _ = fi17
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        fmt.Println(v5)
        if ((23 >= 42) && (-128 >= 127)) {
            var v42 uint64 = 0
            fmt.Println(v2)
            _ = v42
        } else {
            if true {
                var v43 int = 128
                if (-128 == -2) {
                    v44 := GenS5{F0: 65535, F1: 144}
                    if true {
                        for fi45 := 0; fi45 < 10; fi45++ {
                            switch v43 {
                            case 4294967295:
                                if (307 != -2147483648) {
                                    v46 := GenS1{F0: 2147483647, F1: 4294967295, F2: 255}
                                    var v47 uint64 = 42
                                    v47 = v47
                                    _ = v46
                                    _ = v47
                                } else {
                                    v4[fi41] = true
                                    var v48 int = fi45
                                    _ = v48
                                }
                            case 1:
                                switch v43 {
                                case 674:
                                    v49 := GenS6{F0: 0, F1: 0}
                                    _ = v49
                                }
                            }
                            var v50 uint = 66
                            switch int(v50) {
                            case 9223372036854775807:
                                v51 := GenS2{F0: (^(v50)), F1: (v43 | (-(fi45))), F2: (v3 + (v3 << ((v50) & 63)))}
                                _ = v51
                            case -291:
                                v52 := GenS2{F0: (^(v50)), F1: (-(v43)), F2: (v3 << ((v50) & 63))}
                                _ = v52
                            }
                            _ = fi45
                            _ = v50
                        }
                        var v53 int64 = -1
                        for fi54 := 0; fi54 < 7; fi54++ {
                            v3 = 1000
                            _ = fi54
                        }
                        _ = v53
                    } else {
                        fmt.Println(v3, v2, fi41)
                        if (((0.1 * 1.5) * ((0.1 - (1e10 - 3.14159)) + float64(v2))) <= float64((v2 * (1e10 - float32(-2.25))))) {
                            v55 := GenS6{F0: 75, F1: 212}
                            _ = v55
                        } else {
                            v5 = 65
                            for fi56 := 0; fi56 < 18; fi56++ {
                                if len(v1) > 0 {
                                    v1[0] = 680
                                }
                                v1 = append(v1, 292)
                                _ = fi56
                            }
                        }
                        var v57 float32 = float32(1.5)
                        _ = v57
                    }
                    switch "a" {
                    case "\n\t":
                        var v58 int8 = -128
                        _ = v58
                    case "foo bar":
                        v59 := GenS1{F0: 4294967295, F1: 939, F2: v5}
                        _ = v59
                    default:
                        var v60 uint = 0
                        _ = v60
                    }
                    _ = v44
                }
                if false {
                    var v61 uint = 999
                    if ((0 >= 2) || (0 != 1)) {
                        delete(v4, -1)
                        v61 = v61
                    }
                    _ = v61
                } else {
                    v2 -= v2
                    v62 := GenS2{F0: 65535, F1: (-(fi41)), F2: v3}
                    _ = v62
                }
                _ = v43
            }
            for fi63 := 0; fi63 < 11; fi63++ {
                var v64 uint64 = 52
                var v65 uint = 537
                v66 := GenS2{F0: (v65 % ((v65) | 1)), F1: fi63, F2: -174}
                _ = fi63
                _ = v64
                _ = v65
                _ = v66
            }
        }
        fmt.Println(fi41, v2, fi41)
        _ = fi41
    }
    for fi67 := 0; fi67 < 10; fi67++ {
        v68 := GenS1{F0: 42, F1: int(v5), F2: (v5 % ((0) | 1))}
        v69 := GenS7{F0: 445, F1: 1, F2: false}
        _ = fi67
        _ = v68
        _ = v69
    }
    fmt.Println(v3, v2)
    v1 = append(v1, 65535)
    var v70 uint32 = 2147483647
    v70 += v70
    for fi71 := 0; fi71 < 9; fi71++ {
        if (v3 <= v3) {
            v70 = (v70 % (((v70 % ((48) | 1))) | 1))
            if (100 > 42) {
                var v72 int = 42
                v5 -= v5
                _ = v72
            }
            var v73 uint64 = 2
            _ = v73
        }
        v74 := GenS8{F0: "foo bar", F1: -46, F2: 42}
        v75 := GenS6{F0: 37, F1: (v5 ^ v5)}
        _ = fi71
        _ = v74
        _ = v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
}
