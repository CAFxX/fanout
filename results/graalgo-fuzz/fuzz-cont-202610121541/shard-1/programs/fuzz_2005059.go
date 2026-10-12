package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint8; F1 int; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 int64; F1 int64; F2 int32 }
type GenS5 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 int16 }
type GenS8 struct { F0 uint8; F1 string; F2 uint }
type GenS9 struct { F0 bool }
func genFn0(p0 uint8, p1 int16) int32 {
    var v5 int = 9223372036854775807
    if false {
        v6 := GenS1{F0: 405, F1: 68}
        _ = v6
    } else {
        switch v5 {
        case 4294967295:
            switch 32767 {
            case 21:
                var v7 float32 = (1e10 - float32((-2.25 / 1.5)))
                _ = v7
            case -900:
                p1 = (^(p1))
            }
        default:
            v8 := []int{}
            _ = v8
        }
        v9 := GenS1{F0: 2, F1: 1}
        p0 -= p0
        _ = v9
    }
    var v10 uint16 = 1
    var v11 int32 = 42
    if (false || true) {
        v12 := GenS2{F0: p0, F1: -969, F2: 214}
        v10 = (v10 / ((v10) | 1))
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v5
    return int32(p0)
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 int = 42
    var v3 uint8 = 128
    v4 := []string{(v1 + (((v1 + v1) + (v1 + "a")) + (v1 + v1))), v1}
    for fi13 := 0; fi13 < 10; fi13++ {
        v3 += (^(v3))
        v4 = append(v4, v1)
        switch v1 {
        case "a":
            switch (v1 + "foo bar") {
            case "hello":
                var v14 int16 = 0
                _ = v14
            case "foo bar":
                v3 = (v3 << ((65535) & 63))
            default:
                v4 = append(v4, "\n\t")
            }
        default:
            var v15 float32 = (3.14159 + ((0.1 + (1e10 * 3.14159)) + -2.25))
            _ = v15
        }
        _ = fi13
    }
    var v16 float32 = 1e10
    var v17 int64 = 9223372036854775807
    if true {
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 3; fi19++ {
                if (42 <= 2) {
                    var v20 int8 = 1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v20
                }
                _ = fi19
            }
            v16 = v16
            _ = fi18
        }
        v21 := GenS1{F0: 0, F1: 63}
        v22 := GenS3{F0: -32768, F1: 4294967295, F2: (v2 + (v2 - (v2 + v2)))}
        _ = v21
        _ = v22
    } else {
        v1 = v1
        var v23 uint16 = 1
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        if false {
            var v25 int32 = 0
            _ = v25
        } else {
            switch (fi24 >> ((923) & 63)) {
            case 42:
                if (false || false) {
                    v2 -= 0
                    v1 = v1
                } else {
                    v4 = append(v4, v1)
                    fmt.Println(fi24, v2, fi24)
                    var v26 uint = 18446744073709551615
                    _ = v26
                }
            case 267:
                v1 = (v1 + v1)
            }
            v2 = fi24
            if (((1e10 + v16) / 1e10) != 0.1) {
                v27 := GenS1{F0: 255, F1: -37}
                var v28 int64 = -1
                _ = v27
                _ = v28
            } else {
                var v29 uint = 1
                var v30 int16 = int16(fi24)
                _ = v29
                _ = v30
            }
        }
        if (false || (!((false || false) || false))) {
            var v31 uint32 = 4294967295
            var v32 uint16 = 1
            _ = v31
            _ = v32
        } else {
            var v33 uint8 = (v3 >> ((18446744073709551615) & 63))
            if ((1.5 + -2.25) >= (0.0 - (((0.1 * 0.1) + 1.5) - 3.14159))) {
                v34 := GenS3{F0: -161, F1: 627, F2: int(v33)}
                var v35 uint64 = 18446744073709551615
                var v36 uint = 2
                _ = v34
                _ = v35
                _ = v36
            }
            _ = len(v4)
            _ = v33
        }
        _ = fi24
    }
    var v37 int64 = (v17 & (v17 | 9223372036854775807))
    if (!true) {
        switch 4294967295 {
        case 127:
            v37 -= (v37 ^ v37)
        case 32767:
            v1 = (v1 + v1)
        }
        fmt.Println(v17, v3)
        if ((!((32767 > 527) || (false && true))) && true) {
            if false {
                v38 := GenS3{F0: 42, F1: 2, F2: v2}
                _ = v38
            }
            if true {
                switch 0 {
                case 4294967295:
                    v39 := GenS3{F0: 1, F1: 2, F2: (v2 - v2)}
                    _ = v39
                case 32767:
                    var v40 int16 = -121
                    _ = v40
                default:
                    fmt.Println(v2, v3, v2)
                }
                if false {
                    fmt.Println(v16)
                    switch (v2 / (((v2 >> ((65535) & 63))) | 1)) {
                    case 32768:
                        v2 = 9223372036854775807
                    case -884:
                        if (!false) {
                            for fi41 := 0; fi41 < 6; fi41++ {
                                var v42 int64 = 875
                                _ = fi41
                                _ = v42
                            }
                            if len(v4) > 0 {
                                v4[0] = v1
                            }
                        } else {
                            v3 = v3
                            for fi43 := 0; fi43 < 12; fi43++ {
                                if false {
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        var v45 bool = true
                                        switch (-(v2)) {
                                        case 410:
                                            v46 := GenS4{F0: v17, F1: (v17 << ((42) & 63)), F2: 0}
                                            _ = v46
                                        case -566:
                                            switch v2 {
                                            case 291:
                                                var v47 uint32 = 0
                                                _ = v47
                                            }
                                        case 887:
                                            switch 0 {
                                            case 604:
                                                v48 := GenS5{F0: 268, F1: 42, F2: 4294967295}
                                                _ = v48
                                            case 580:
                                                v49 := GenS4{F0: (v37 - -9223372036854775808), F1: v17, F2: 140}
                                                _ = v49
                                            case 178:
                                                v45 = v45
                                            default:
                                                v37 = (v37 / ((0) | 1))
                                            }
                                        default:
                                            v50 := GenS4{F0: v37, F1: v37, F2: 2147483647}
                                            _ = v50
                                        }
                                        v51 := GenS6{F0: 127}
                                        _ = fi44
                                        _ = v45
                                        _ = v51
                                    }
                                    if false {
                                        v17 = (v37 + (v37 >> ((2) & 63)))
                                    }
                                    switch v2 {
                                    case 544:
                                        v1 = ((v1 + "foo bar") + v1)
                                    case 2147483647:
                                        v52 := GenS7{F0: 307, F1: -32768}
                                        _ = v52
                                    case 42:
                                        var v53 string = ((("foo bar" + v1) + v1) + (v1 + v1))
                                        _ = v53
                                    }
                                } else {
                                    var v54 float32 = float32((3.14159 * 1.5))
                                    _ = v54
                                }
                                v55 := GenS5{F0: 880, F1: -128, F2: 18446744073709551615}
                                _ = fi43
                                _ = v55
                            }
                            fmt.Println(v2, v37)
                        }
                    case 2147483647:
                        v3 -= v3
                    }
                } else {
                    var v56 int = 1
                    if true {
                        v3 -= (v3 >> ((241) & 63))
                        var v57 int8 = 127
                        fmt.Println(v1, v57)
                        _ = v57
                    }
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    _ = v56
                }
                for fi58 := 0; fi58 < 15; fi58++ {
                    switch (fi58 * v2) {
                    case 128:
                        if false {
                            if len(v4) > 0 {
                                v4[0] = ((v1 + (v1 + v1)) + "hello")
                            }
                            var v59 bool = false
                            var v60 uint8 = v3
                            _ = v59
                            _ = v60
                        } else {
                            v17 += v17
                            switch (fi58 + (fi58 - v2)) {
                            case -1:
                                switch (v1 + ((v1 + "") + "")) {
                                case "\n\t":
                                    v61 := GenS6{F0: 127}
                                    _ = v61
                                case "":
                                    for fi62 := 0; fi62 < 4; fi62++ {
                                        var v63 uint32 = 983
                                        _ = fi62
                                        _ = v63
                                    }
                                default:
                                    if ((fi58 * 128) >= 0) {
                                        switch (v1 + "foo bar") {
                                        case "foo bar":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        switch (v1 + v1) {
                                        case "foo bar":
                                            for fi64 := 0; fi64 < 13; fi64++ {
                                                var v65 int16 = 42
                                                v4 = append(v4, v1)
                                                var v66 int32 = 135
                                                if false {
                                                    v4 = append(v4, ((v1 + (v1 + v1)) + ("hello" + v1)))
                                                }
                                                if true {
                                                    v67 := GenS3{F0: int16(v66), F1: 1, F2: fi64}
                                                    _ = v67
                                                }
                                                switch v1 {
                                                case "foo bar":
                                                    if (1 != 100) {
                                                        if false {
                                                            fmt.Println(v16)
                                                            fmt.Println(v65, v17, v3)
                                                            v37 += (v37 / ((v37) | 1))
                                                        } else {
                                                            var v68 uint32 = 304
                                                            var v69 int16 = (v65 / (((^(v65))) | 1))
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                    } else {
                                                        if ((42 != 614) || false) {
                                                            v70 := GenS4{F0: v37, F1: (-(v17)), F2: v66}
                                                            _ = v70
                                                        } else {
                                                            var v71 uint8 = v3
                                                            var v72 string = v1
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                        var v73 int8 = int8(v2)
                                                        fmt.Println(v1, v73)
                                                        _ = v73
                                                    }
                                                }
                                                if true {
                                                    var v74 int32 = v66
                                                    if ((false || true) || ((3.14159 / 1.5) > v16)) {
                                                        v3 = (v3 - 195)
                                                        v74 -= v66
                                                        if false {
                                                            v66 -= v74
                                                            var v75 int64 = v17
                                                            v76 := GenS8{F0: v3, F1: v1, F2: 195}
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                        var v77 float32 = ((((((0.0 * v16) / 0.1) * ((float32((-2.25 / (1.5 * -2.25))) + 0.0) * v16)) + 0.1) - v16) + 0.0)
                                                        _ = v77
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v78 float32 = 3.14159
                                                    var v79 uint = 383
                                                    _ = len(v4)
                                                    _ = v78
                                                    _ = v79
                                                }
                                                _ = fi64
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    }
                                }
                            case 9223372036854775807:
                                v2 = 553
                            case -884:
                                v1 = v1
                            default:
                                var v80 bool = ((v3 >> ((4294967295) & 63)) > v3)
                                _ = v80
                            }
                        }
                    }
                    v81 := GenS4{F0: -1, F1: (v17 * (v17 >> ((4294967295) & 63))), F2: genFn0((v3 << ((65535) & 63)), int16(fi58))}
                    _ = fi58
                    _ = v81
                }
            }
        } else {
            var v82 int32 = 0
            _ = v82
        }
    } else {
        var v83 uint16 = 903
        for fi84 := 0; fi84 < 12; fi84++ {
            for fi85 := 0; fi85 < 8; fi85++ {
                var v86 float64 = 0.1
                if len(v4) > 0 {
                    v4[0] = ((v1 + v1) + v1)
                }
                _ = fi85
                _ = v86
            }
            _ = fi84
        }
        _ = v83
    }
    var v87 int64 = (v17 + (v37 << ((1) & 63)))
    v88 := GenS9{F0: false}
    fmt.Println(v16, v87, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v87
    _ = v88
}
