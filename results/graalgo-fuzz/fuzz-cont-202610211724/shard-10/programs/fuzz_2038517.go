package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8; F2 bool }
type GenS6 struct { F0 string; F1 uint64; F2 bool }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 bool; F1 int64 }
func genFn0(p0 int) int8 {
    v4 := GenS1{F0: 444, F1: (162 <= 55), F2: 640}
    var v5 int = p0
    _ = p0
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 int64, p1 uint16, p2 int32) uint8 {
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: p1, F1: ((32767 >= 1000) && false), F2: 42}
            v9 := []int{fi7}
            var v10 uint32 = 0
            var v11 int64 = 25
            _ = fi7
            _ = v10
            _ = v11
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v12 uint64 = 981
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 128
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 602
    v3 := []uint32{2147483647, 0, 0}
    _ = len(v3)
    if (!(false && (-2147483648 < 9223372036854775807))) {
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 int8 = 127
            _ = fi13
            _ = v14
        }
    } else {
        if len(v3) > 0 {
            v3[0] = 0
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 4294967295 {
    case -1:
        v15 := GenS2{F0: 42, F1: 0}
        _ = v15
    case 459:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -2147483648:
        v1 += 4294967295
    default:
        switch -2147483648 {
        case 1:
            v16 := GenS3{F0: 1, F1: 0}
            _ = v16
        case -9223372036854775808:
            if false {
                var v17 uint8 = 2
                v1 -= v1
                _ = len(v3)
                _ = v17
            } else {
                if (false && false) {
                    if false {
                        if false {
                            v18 := GenS2{F0: genFn0(32767), F1: -128}
                            _ = v18
                        }
                        v1 = v1
                        var v19 float32 = float32(0.0)
                        _ = v19
                    }
                    var v20 float64 = 3.14159
                    _ = v20
                }
                if ((1 < 9223372036854775807) && (127 != -1)) {
                    var v21 float32 = (0.1 - 1e10)
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    _ = v21
                } else {
                    if true {
                        var v22 int32 = -285
                        _ = v22
                    } else {
                        var v23 uint8 = 127
                        v24 := GenS2{F0: 15, F1: 127}
                        _ = v23
                        _ = v24
                    }
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    var v25 uint16 = 255
                    _ = v25
                }
            }
        case 999:
            if false {
                v1 = 18446744073709551615
            }
        }
    }
    if false {
        v26 := GenS3{F0: 2147483647, F1: -32768}
        _ = v26
    } else {
        if (!false) {
            _ = len(v3)
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 uint = v2
                _ = fi27
                _ = v28
            }
            var v29 uint64 = v1
            _ = v29
        }
        var v30 int16 = -32768
        switch -9223372036854775808 {
        case -872:
            var v31 uint64 = (v1 * 1)
            _ = v31
        case 682:
            switch ("" + "a") {
            case "a":
                for fi32 := 0; fi32 < 11; fi32++ {
                    switch -2147483648 {
                    case -318:
                        switch "foo bar" {
                        case "a":
                            var v33 bool = (true && (3.14159 < 3.14159))
                            _ = v33
                        case "\n\t":
                            var v34 string = "a"
                            _ = v34
                        }
                    case 128:
                        v35 := GenS4{F0: genFn1(297, 0, 42)}
                        _ = v35
                    case 2147483647:
                        v36 := GenS5{F0: 65535, F1: -121, F2: ((853 > -2147483648) || (!true))}
                        _ = v36
                    }
                    var v37 int = fi32
                    fmt.Println(v30, fi32, v2)
                    _ = fi32
                    _ = v37
                }
            case "":
                fmt.Println(v2, v30, v30)
            default:
                switch 2147483647 {
                case 32768:
                    var v38 uint64 = 1
                    _ = v38
                default:
                    v39 := GenS3{F0: -538, F1: (^(v30))}
                    _ = v39
                }
            }
        case 42:
            if (!(v1 <= 4294967295)) {
                switch "" {
                case "foo bar":
                    fmt.Println(v1, v1, v30)
                case "":
                    for fi40 := 0; fi40 < 7; fi40++ {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            v3 = append(v3, 65535)
                            _ = fi41
                        }
                        var v42 int32 = -225
                        _ = fi40
                        _ = v42
                    }
                case "\n\t":
                    v43 := GenS3{F0: int32(v2), F1: -425}
                    _ = v43
                }
            } else {
                v30 = v30
                fmt.Println(v2)
            }
        }
        _ = v30
    }
    _ = len(v3)
    if ((2 <= 346) || true) {
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 float32 = (1.5 * 1.5)
            _ = fi44
            _ = v45
        }
        switch 0 {
        case 865:
            var v46 string = ""
            _ = v46
        default:
            for fi47 := 0; fi47 < 7; fi47++ {
                v48 := GenS5{F0: (v2 & v2), F1: 42, F2: true}
                if (1 != 527) {
                    for fi49 := 0; fi49 < 3; fi49++ {
                        if false {
                            _ = len(v3)
                            v50 := GenS6{F0: ("hello" + "hello"), F1: v1, F2: true}
                            v51 := GenS3{F0: 2147483647, F1: 1000}
                            _ = v50
                            _ = v51
                        }
                        _ = fi49
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    fmt.Println(v1, v1)
                }
                _ = fi47
                _ = v48
            }
        }
        var v52 bool = (431 > 42)
        _ = v52
    } else {
        if (42 == 235) {
            v2 -= v2
        } else {
            fmt.Println(v2)
            _ = len(v3)
        }
        v3 = append(v3, 2)
        for fi53 := 0; fi53 < 6; fi53++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi53
        }
    }
    var v54 int32 = 745
    v2 += v2
    switch -1 {
    case 127:
        for fi55 := 0; fi55 < 4; fi55++ {
            switch "hello" {
            case "hello":
                v54 -= (v54 ^ v54)
            case "foo bar":
                var v56 int8 = 0
                _ = v56
            case "a":
                if true {
                    if true {
                        v54 = (v54 - 42)
                    }
                    v54 = v54
                } else {
                    if false {
                        var v57 int16 = 1
                        _ = v57
                    }
                    fmt.Println(v2)
                    v2 -= v2
                }
            default:
                v1 = 18446744073709551615
            }
            _ = fi55
        }
    default:
        if true {
            if ((genFn1(-891, 42, (v54 << ((356) & 63))) > 25) || ("hello" >= "a")) {
                var v58 float32 = ((float32(-2.25) + 1e10) + (0.0 - float32(1.5)))
                _ = v58
            }
            if (!((v54 ^ v54) > v54)) {
                var v59 uint32 = 416
                var v60 int32 = v54
                if false {
                    var v61 int = -9223372036854775808
                    for fi62 := 0; fi62 < 6; fi62++ {
                        switch "hello" {
                        case "\n\t":
                            if true {
                                var v63 bool = true
                                for fi64 := 0; fi64 < 4; fi64++ {
                                    switch (fi64 - fi62) {
                                    case 2147483647:
                                        v65 := GenS4{F0: 128}
                                        _ = v65
                                    case 4294967295:
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                    }
                                    var v66 float32 = -2.25
                                    _ = fi64
                                    _ = v66
                                }
                                v67 := GenS2{F0: 125, F1: genFn0(v61)}
                                _ = v63
                                _ = v67
                            } else {
                                fmt.Println(v61, v60)
                                switch v61 {
                                case 4294967295:
                                    switch "" {
                                    case "hello":
                                        var v68 float64 = float64(((1.5 - 0.1) + 0.0))
                                        _ = v68
                                    case "":
                                        v69 := GenS5{F0: 270, F1: int8(v59), F2: true}
                                        _ = v69
                                    }
                                }
                            }
                        case "hello":
                            var v70 uint16 = 1
                            _ = v70
                        }
                        var v71 uint = (v2 % ((2) | 1))
                        var v72 int = 469
                        _ = fi62
                        _ = v71
                        _ = v72
                    }
                    _ = v61
                } else {
                    if len(v3) > 0 {
                        v3[0] = v59
                    }
                    if true {
                        fmt.Println(v59, v60)
                        for fi73 := 0; fi73 < 18; fi73++ {
                            switch (("" + (("" + "") + "a")) + (("" + "") + "foo bar")) {
                            case "foo bar":
                                v54 += (v54 * v54)
                            case "\n\t":
                                var v74 int64 = 0
                                _ = v74
                            case "a":
                                switch int(v1) {
                                case 128:
                                    var v75 string = ("hello" + "hello")
                                    _ = v75
                                case 4294967295:
                                    switch fi73 {
                                    case 32768:
                                        v76 := GenS7{F0: 182, F1: false}
                                        _ = v76
                                    case 32767:
                                        v77 := GenS2{F0: genFn0(-586), F1: -1}
                                        _ = v77
                                    default:
                                        var v78 int8 = -32
                                        _ = v78
                                    }
                                default:
                                    for fi79 := 0; fi79 < 12; fi79++ {
                                        var v80 uint32 = v59
                                        _ = fi79
                                        _ = v80
                                    }
                                }
                            }
                            var v81 int16 = 1000
                            _ = fi73
                            _ = v81
                        }
                        if false {
                            for fi82 := 0; fi82 < 14; fi82++ {
                                var v83 uint8 = 128
                                _ = fi82
                                _ = v83
                            }
                        } else {
                            var v84 float32 = ((1e10 * -2.25) - 1.5)
                            _ = v84
                        }
                    }
                    var v85 float32 = 0.0
                    _ = v85
                }
                _ = v59
                _ = v60
            }
        }
    }
    if true {
        v1 += 624
        v1 -= 2
        var v86 int16 = -1
        _ = v86
    } else {
        switch "" {
        case "a":
            for fi87 := 0; fi87 < 7; fi87++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v54 -= (v54 & (-(v54)))
                fmt.Println(v2, v54, fi87)
                _ = fi87
            }
        case "\n\t":
            v3 = append(v3, 1)
        case "hello":
            var v88 int8 = 0
            _ = v88
        }
        for fi89 := 0; fi89 < 18; fi89++ {
            for fi90 := 0; fi90 < 17; fi90++ {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
                var v91 int16 = 1
                _ = fi90
                _ = v91
            }
            for fi92 := 0; fi92 < 17; fi92++ {
                var v93 int16 = 1
                _ = fi92
                _ = v93
            }
            v54 += v54
            _ = fi89
        }
    }
    v54 = (v54 ^ 42)
    var v94 uint = (v2 - 560)
    for fi95 := 0; fi95 < 5; fi95++ {
        v54 = -36
        if (1 > 42) {
            v96 := GenS8{F0: (!false), F1: -1}
            var v97 int8 = 42
            v98 := GenS2{F0: (v97 * genFn0((fi95 % ((fi95) | 1)))), F1: genFn0(fi95)}
            _ = v96
            _ = v97
            _ = v98
        }
        _ = fi95
    }
    switch "hello" {
    case "foo bar":
        if (("foo bar" >= "hello") || false) {
            var v99 uint8 = uint8(v1)
            if (!true) {
                v99 = (v99 << ((743) & 63))
                if true {
                    var v100 string = "foo bar"
                    _ = v100
                }
            } else {
                var v101 uint8 = v99
                var v102 bool = false
                _ = v101
                _ = v102
            }
            _ = v99
        }
    case "\n\t":
        var v103 int = -1
        _ = v103
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v104 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v104
    _ = v2
    _ = v3
    _ = v54
    _ = v94
}
