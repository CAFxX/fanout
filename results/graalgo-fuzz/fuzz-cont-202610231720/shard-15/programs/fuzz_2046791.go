package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint8; F1 int64; F2 int64 }
func genFn0() uint64 {
    var v4 int32 = 1
    var v5 int64 = 1
    v6 := []int{42, 0, -2147483648, -874}
    var v7 uint32 = 27
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 197}
        switch fi8 {
        case -655:
            var v10 int8 = 0
            _ = v10
        case 128:
            for fi11 := 0; fi11 < 3; fi11++ {
                switch "foo bar" {
                case "a":
                    v12 := GenS1{F0: 0}
                    _ = v12
                case "":
                    _ = len(v6)
                case "\n\t":
                    v6 = append(v6, 4294967295)
                }
                _ = fi11
            }
        case 9223372036854775807:
            for fi13 := 0; fi13 < 15; fi13++ {
                v4 -= (v4 + v4)
                _ = fi13
            }
        default:
            v7 = 624
        }
        _ = fi8
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 42
    var v2 float32 = (3.14159 * ((1.5 + 0.0) * 0.0))
    var v3 string = ""
    v14 := []int{}
    var v15 int16 = 1000
    if (!true) {
        switch v3 {
        case "a":
            v15 = v15
        default:
            _ = len(v14)
        }
    } else {
        v16 := GenS1{F0: (-(v15))}
        var v17 uint32 = 2
        if ((((v2 + (v2 * 3.14159)) > v2) || (uint8(v1) > uint8(v1))) && (0.1 > (0.0 + float64(v2)))) {
            for fi18 := 0; fi18 < 4; fi18++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = fi18
            }
            var v19 int = 32768
            fmt.Println(v17)
            _ = v19
        }
        _ = v16
        _ = v17
    }
    fmt.Println(v1)
    v1 = v1
    switch 32767 {
    case 9223372036854775807:
        if true {
            if true {
                if true {
                    v15 += (v15 % ((v15) | 1))
                    var v20 uint16 = 763
                    v1 -= v1
                    _ = v20
                }
                if (v3 <= "hello") {
                    switch 15 {
                    case 0:
                        for fi21 := 0; fi21 < 8; fi21++ {
                            for fi22 := 0; fi22 < 16; fi22++ {
                                v23 := GenS2{F0: -1, F1: 2}
                                _ = fi22
                                _ = v23
                            }
                            if (!true) {
                                if ((v15 << (((v1 / ((0) | 1))) & 63)) != v15) {
                                    if (0 == 0) {
                                        switch int(v1) {
                                        case -551:
                                            switch (fi21 & (fi21 / ((fi21) | 1))) {
                                            case -351:
                                                v1 = v1
                                            case -660:
                                                v24 := GenS1{F0: 1000}
                                                _ = v24
                                            }
                                        case 969:
                                            v2 -= 0.0
                                        case -626:
                                            v15 -= (v15 << (((v1 / ((v1) | 1))) & 63))
                                        }
                                        v1 = 416
                                    }
                                    _ = len(v14)
                                    var v25 uint64 = 1
                                    _ = v25
                                } else {
                                    if (true || (-2.25 < (((1e10 / 3.14159) - 0.0) + 0.0))) {
                                        var v26 int16 = 588
                                        if true {
                                            v26 = v26
                                            fmt.Println(v2)
                                            var v27 int32 = 2147483647
                                            _ = v27
                                        }
                                        _ = v26
                                    }
                                }
                            } else {
                                if len(v14) > 0 {
                                    v14[0] = 98
                                }
                                for fi28 := 0; fi28 < 5; fi28++ {
                                    fmt.Println(fi21)
                                    _ = fi28
                                }
                            }
                            v29 := GenS2{F0: 974, F1: 42}
                            _ = fi21
                            _ = v29
                        }
                    case -2147483648:
                        for fi30 := 0; fi30 < 18; fi30++ {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            var v31 int16 = (v15 + v15)
                            _ = fi30
                            _ = v31
                        }
                    case 1:
                        v2 = float32(float64(1.5))
                    }
                    if len(v14) > 0 {
                        v14[0] = 167
                    }
                    var v32 float32 = v2
                    _ = v32
                }
                for fi33 := 0; fi33 < 1; fi33++ {
                    fmt.Println(v15)
                    _ = fi33
                }
            }
        } else {
            if (-101 == 117) {
                v2 = (v2 * ((((v2 + v2) / v2) * v2) / v2))
                var v34 int = -1
                _ = v34
            } else {
                fmt.Println(v2, v1, v2)
                var v35 float32 = ((3.14159 * v2) * (1e10 - (v2 / v2)))
                v36 := GenS2{F0: 302, F1: 255}
                _ = v35
                _ = v36
            }
            var v37 uint32 = 393
            _ = v37
        }
    default:
        switch 433 {
        case 9223372036854775807:
            v14 = append(v14, -1)
        case 4294967295:
            fmt.Println(v15, v3)
        case -1:
            v38 := GenS1{F0: (v15 | v15)}
            _ = v38
        default:
            var v39 int64 = 42
            _ = v39
        }
    }
    var v40 bool = (!(false || true))
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(fi41)
        _ = fi41
    }
    if v40 {
        if v40 {
            var v42 int8 = 0
            v15 = v15
            _ = v42
        } else {
            if false {
                v43 := GenS3{F0: 2147483647, F1: 255}
                _ = v43
            } else {
                switch 4294967295 {
                case 42:
                    for fi44 := 0; fi44 < 14; fi44++ {
                        for fi45 := 0; fi45 < 11; fi45++ {
                            switch (fi45 % ((fi45) | 1)) {
                            case 4294967295:
                                if (int8(v15) != 118) {
                                    v1 = 4294967295
                                } else {
                                    var v46 int32 = 42
                                    _ = v46
                                }
                            case -866:
                                if (91 < 65535) {
                                    var v47 float64 = 0.1
                                    v1 += v1
                                    _ = v47
                                } else {
                                    var v48 uint32 = 154
                                    _ = v48
                                }
                            }
                            v2 -= v2
                            _ = fi45
                        }
                        _ = fi44
                    }
                case 127:
                    fmt.Println(v3, v40, v40)
                case 0:
                    v2 = 1e10
                default:
                    var v49 int32 = 109
                    _ = v49
                }
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                switch 0 {
                case 239:
                    for fi50 := 0; fi50 < 11; fi50++ {
                        v51 := GenS1{F0: v15}
                        fmt.Println(v2, v3, v1)
                        for fi52 := 0; fi52 < 16; fi52++ {
                            switch 1 {
                            case 32767:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            case 1:
                                v53 := GenS1{F0: -871}
                                _ = v53
                            case 32768:
                                var v54 uint32 = 65535
                                _ = v54
                            default:
                                _ = len(v14)
                            }
                            _ = fi52
                        }
                        _ = fi50
                        _ = v51
                    }
                case 42:
                    _ = len(v14)
                case -1:
                    var v55 int16 = (v15 / (((v15 % ((v15) | 1))) | 1))
                    _ = v55
                default:
                    switch "foo bar" {
                    case "\n\t":
                        if len(v14) > 0 {
                            v14[0] = 2147483647
                        }
                    case "foo bar":
                        v3 = v3
                    case "a":
                        if v40 {
                            var v56 int16 = (v15 + 32767)
                            v1 += v1
                            switch 9223372036854775807 {
                            case -2147483648:
                                fmt.Println(v2, v15)
                            case 252:
                                v3 = "a"
                            }
                            _ = v56
                        }
                    default:
                        var v57 int8 = 127
                        _ = v57
                    }
                }
            }
            if len(v14) > 0 {
                v14[0] = 42
            }
        }
    } else {
        v58 := GenS1{F0: v15}
        for fi59 := 0; fi59 < 3; fi59++ {
            for fi60 := 0; fi60 < 20; fi60++ {
                switch "hello" {
                case "a":
                    var v61 int16 = (v15 & v15)
                    _ = v61
                }
                _ = fi60
            }
            v1 -= v1
            v62 := GenS2{F0: -1, F1: 1}
            _ = fi59
            _ = v62
        }
        _ = v58
    }
    switch "a" {
    case "foo bar":
        if v40 {
            v40 = (!(9223372036854775807 == 1))
        }
    case "a":
        v63 := GenS2{F0: 0, F1: 104}
        _ = v63
    default:
        v64 := GenS4{F0: 127, F1: -97}
        _ = v64
    }
    for fi65 := 0; fi65 < 2; fi65++ {
        fmt.Println(v15, v15)
        _ = fi65
    }
    if (v3 < ("hello" + v3)) {
        switch 42 {
        case 912:
            switch (v3 + v3) {
            case "a":
                for fi66 := 0; fi66 < 15; fi66++ {
                    switch (fi66 | fi66) {
                    case 106:
                        v40 = (!v40)
                    case 42:
                        v67 := GenS2{F0: 2147483647, F1: 0}
                        _ = v67
                    default:
                        switch 2147483647 {
                        case -135:
                            v68 := GenS5{F0: 201}
                            _ = v68
                        case 32767:
                            var v69 uint32 = 0
                            _ = v69
                        }
                    }
                    var v70 int64 = int64(v1)
                    _ = fi66
                    _ = v70
                }
            case "\n\t":
                for fi71 := 0; fi71 < 4; fi71++ {
                    _ = len(v14)
                    var v72 string = v3
                    _ = len(v14)
                    _ = fi71
                    _ = v72
                }
            }
        case -9223372036854775808:
            var v73 int64 = 9223372036854775807
            _ = v73
        default:
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        var v74 int64 = -62
        _ = v74
    }
    if v40 {
        v40 = v40
        v14 = append(v14, 0)
        var v75 uint16 = 42
        _ = v75
    }
    for fi76 := 0; fi76 < 18; fi76++ {
        v3 = v3
        var v77 uint8 = 168
        switch v3 {
        case "a":
            var v78 int32 = 2147483647
            _ = v78
        case "":
            if (true && (((((((v2 - (v2 - 0.0)) / v2) * float32(((1.5 / (float64(v2) * (1e10 + (0.0 - 0.0)))) / (3.14159 - 1.5)))) * v2) * ((-2.25 * (float32(float64(-2.25)) - (v2 + v2))) / v2)) == v2) || v40)) {
                if true {
                    if (4294967295 == genFn0()) {
                        if len(v14) > 0 {
                            v14[0] = -1
                        }
                    }
                    v14 = append(v14, fi76)
                } else {
                    v79 := GenS3{F0: (fi76 ^ (fi76 / ((fi76) | 1))), F1: v77}
                    _ = v79
                }
            } else {
                if (((((v77 * (^(v77))) == 42) && v40) || false) || true) {
                    switch (-(fi76)) {
                    case 310:
                        v1 = (v1 >> (((v1 | (v1 * 0))) & 63))
                    }
                    v14 = append(v14, 128)
                    for fi80 := 0; fi80 < 12; fi80++ {
                        switch 9223372036854775807 {
                        case 127:
                            v1 = v1
                        case -25:
                            var v81 uint8 = v77
                            _ = v81
                        }
                        _ = len(v14)
                        var v82 uint8 = 127
                        _ = fi80
                        _ = v82
                    }
                }
                if (v2 == 1e10) {
                    _ = len(v14)
                    if v40 {
                        for fi83 := 0; fi83 < 14; fi83++ {
                            v84 := GenS3{F0: (fi83 | fi83), F1: 162}
                            _ = fi83
                            _ = v84
                        }
                        if true {
                            switch fi76 {
                            case 42:
                                v85 := GenS6{F0: (v77 & v77), F1: 80, F2: 274}
                                _ = v85
                            case -281:
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            default:
                                v1 = 177
                            }
                        } else {
                            v86 := GenS3{F0: 127, F1: (v77 % (((^(v77))) | 1))}
                            v87 := GenS5{F0: genFn0()}
                            var v88 int8 = -1
                            _ = v86
                            _ = v87
                            _ = v88
                        }
                    }
                }
            }
        default:
            if (v40 || (!v40)) {
                v15 -= (-(v15))
                fmt.Println(v40)
                fmt.Println(v15)
            }
        }
        _ = fi76
        _ = v77
    }
    if v40 {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
    }
    var v89 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v40
    _ = v89
}
