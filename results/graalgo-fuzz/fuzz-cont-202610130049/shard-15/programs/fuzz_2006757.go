package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 bool; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint }

func main() {
    v1 := []int{-756, 2147483647, -2147483648, 1}
    var v2 uint64 = 4294967295
    var v3 int32 = 42
    var v4 uint8 = 214
    v5 := []int{-1, -347}
    var v6 string = "a"
    fmt.Println(v6, v6)
    var v7 uint = 255
    fmt.Println(v2)
    if (!true) {
        v3 = v3
        var v8 int8 = -1
        var v9 string = (v6 + v6)
        _ = v8
        _ = v9
    } else {
        var v10 int8 = 100
        var v11 string = v6
        if (1.5 <= 1.5) {
            switch (v6 + v6) {
            case "foo bar":
                fmt.Println(v10)
            }
        }
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS1{F0: 1, F1: 826}
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v4, v6)
            _ = fi14
        }
        if (1 >= 42) {
            v15 := GenS1{F0: 4294967295, F1: 255}
            _ = v15
        } else {
            var v16 int16 = 1
            for fi17 := 0; fi17 < 8; fi17++ {
                switch (^(fi12)) {
                case -984:
                    v18 := GenS2{F0: 273}
                    _ = v18
                }
                switch fi12 {
                case -9223372036854775808:
                    fmt.Println(v3, v6, v3)
                case 2147483647:
                    v5 = append(v5, (fi12 | (fi17 - (fi12 - 2147483647))))
                default:
                    switch fi12 {
                    case -350:
                        var v19 bool = ((v3 | (^(v3))) >= -1)
                        _ = v19
                    case 4294967295:
                        var v20 uint = 42
                        _ = v20
                    }
                }
                _ = fi17
            }
            v6 = "\n\t"
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v4)
    switch (v6 + ((("hello" + v6) + "hello") + v6)) {
    case "hello":
        var v21 int64 = 361
        _ = v21
    case "\n\t":
        if ((v2 - 1) != v2) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (!(103 <= 93)) {
                if false {
                    var v22 int64 = -9223372036854775808
                    _ = v22
                } else {
                    if ((v3 >> ((42) & 63)) <= v3) {
                        var v23 string = v6
                        var v24 uint8 = 122
                        for fi25 := 0; fi25 < 5; fi25++ {
                            v3 = -1
                            v3 += v3
                            _ = fi25
                        }
                        _ = v23
                        _ = v24
                    }
                    v1 = append(v1, 42)
                    v26 := GenS3{F0: v3, F1: v3, F2: 65535}
                    _ = v26
                }
            } else {
                if true {
                    var v27 uint16 = 255
                    _ = v27
                }
            }
        } else {
            if true {
                var v28 uint = 18446744073709551615
                if ((v3 >> ((v7) & 63)) == 1) {
                    var v29 uint16 = 255
                    for fi30 := 0; fi30 < 11; fi30++ {
                        if true {
                            v28 = uint(v4)
                        }
                        var v31 uint8 = v4
                        if true {
                            switch v6 {
                            case "a":
                                v7 = uint(v4)
                            case "":
                                v7 = 4294967295
                            case "foo bar":
                                _ = len(v1)
                            default:
                                if true {
                                    var v32 int = 128
                                    _ = v32
                                } else {
                                    var v33 uint16 = (v29 ^ (v29 << ((v28) & 63)))
                                    switch (-(fi30)) {
                                    case 42:
                                        fmt.Println(v29, v28)
                                    case 0:
                                        _ = len(v5)
                                    }
                                    for fi34 := 0; fi34 < 12; fi34++ {
                                        v35 := GenS1{F0: (v2 % ((v2) | 1)), F1: uint16(v31)}
                                        v36 := GenS3{F0: 0, F1: v3, F2: v33}
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                    _ = v33
                                }
                            }
                            v37 := GenS3{F0: v3, F1: -1, F2: 708}
                            if true {
                                var v38 int64 = -9223372036854775808
                                _ = v38
                            } else {
                                v6 = (((("hello" + (v6 + v6)) + v6) + v6) + v6)
                                var v39 float32 = 1.5
                                _ = v39
                            }
                            _ = v37
                        }
                        _ = fi30
                        _ = v31
                    }
                    _ = v29
                }
                v28 += v28
                _ = v28
            } else {
                var v40 uint64 = (v2 * v2)
                for fi41 := 0; fi41 < 16; fi41++ {
                    switch v6 {
                    case "\n\t":
                        var v42 int = (fi41 | 59)
                        _ = v42
                    case "foo bar":
                        v6 = (v6 + v6)
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi41
                }
                var v43 float32 = 1e10
                _ = v40
                _ = v43
            }
            if (true && true) {
                for fi44 := 0; fi44 < 12; fi44++ {
                    v7 = uint(v2)
                    v45 := GenS2{F0: 2147483647}
                    if (false && ((fi44 >> ((255) & 63)) >= (-(fi44)))) {
                        for fi46 := 0; fi46 < 19; fi46++ {
                            v47 := GenS4{F0: 1, F1: (fi46 & fi44)}
                            _ = fi46
                            _ = v47
                        }
                        var v48 int8 = 42
                        _ = v48
                    } else {
                        switch (fi44 + fi44) {
                        case -400:
                            if (!true) {
                                if ((v2 * v2) <= (v2 + (v2 << ((v7) & 63)))) {
                                    for fi49 := 0; fi49 < 12; fi49++ {
                                        var v50 int = -9223372036854775808
                                        switch fi44 {
                                        case 128:
                                            if false {
                                                var v51 string = v6
                                                var v52 bool = (!(!false))
                                                fmt.Println(fi44)
                                                _ = v51
                                                _ = v52
                                            } else {
                                                var v53 uint8 = 247
                                                _ = v53
                                            }
                                        case 42:
                                            var v54 uint8 = (v4 >> ((uint(v2)) & 63))
                                            _ = v54
                                        default:
                                            fmt.Println(v6, v6)
                                        }
                                        _ = fi49
                                        _ = v50
                                    }
                                    switch v6 {
                                    case "a":
                                        v3 = (v3 * v3)
                                    }
                                    var v55 float64 = (3.14159 - 0.1)
                                    _ = v55
                                } else {
                                    v1 = append(v1, (fi44 / (((fi44 | (fi44 & fi44))) | 1)))
                                    v5 = append(v5, (fi44 & fi44))
                                }
                                v56 := GenS3{F0: v3, F1: v3, F2: 255}
                                var v57 int16 = -32768
                                _ = v56
                                _ = v57
                            }
                        case 181:
                            v58 := GenS4{F0: 583, F1: (fi44 >> (((v7 ^ 4294967295)) & 63))}
                            _ = v58
                        case 152:
                            var v59 uint8 = v4
                            _ = v59
                        }
                    }
                    _ = fi44
                    _ = v45
                }
                if len(v5) > 0 {
                    v5[0] = 42
                }
            } else {
                v60 := GenS5{F0: (!false), F1: -9223372036854775808, F2: 42}
                if len(v5) > 0 {
                    v5[0] = 4294967295
                }
                v2 -= (v2 | (^(v2)))
                _ = v60
            }
            fmt.Println(v4, v2, v7)
        }
    }
    if (32767 <= 32767) {
        v61 := GenS1{F0: (v2 / ((v2) | 1)), F1: 65535}
        v62 := GenS5{F0: false, F1: -1, F2: 977}
        _ = v61
        _ = v62
    } else {
        v1 = append(v1, -1)
        fmt.Println(v3)
        if (!(2147483648 >= 2147483648)) {
            for fi63 := 0; fi63 < 7; fi63++ {
                v7 = v7
                v3 -= -1
                var v64 uint = (v7 ^ uint(v4))
                _ = fi63
                _ = v64
            }
            for fi65 := 0; fi65 < 16; fi65++ {
                if ((0.1 * ((1e10 + 0.1) * 0.1)) >= (float32(1.5) - -2.25)) {
                    var v66 uint64 = v2
                    v66 = v66
                    var v67 int8 = 42
                    _ = v66
                    _ = v67
                } else {
                    fmt.Println(fi65)
                    fmt.Println(v6)
                }
                _ = fi65
            }
        } else {
            _ = len(v1)
            var v68 int = -324
            if true {
                if (0.0 >= ((float64(3.14159) / (-2.25 * (float64(((1.5 * 3.14159) + -2.25)) - 1e10))) / -2.25)) {
                    var v69 float32 = 1.5
                    _ = v69
                } else {
                    var v70 uint16 = 698
                    if len(v1) > 0 {
                        v1[0] = -727
                    }
                    for fi71 := 0; fi71 < 17; fi71++ {
                        v5 = append(v5, 32768)
                        if (false || true) {
                            v3 = v3
                        } else {
                            v2 -= (v2 % (((v2 + (v2 << ((v7) & 63)))) | 1))
                            v2 += 396
                        }
                        var v72 uint16 = v70
                        _ = fi71
                        _ = v72
                    }
                    _ = v70
                }
            } else {
                v73 := GenS4{F0: 42, F1: int(v7)}
                _ = v73
            }
            _ = v68
        }
    }
    v74 := GenS6{F0: (v7 & (v7 % ((4294967295) | 1)))}
    var v75 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v75
}
