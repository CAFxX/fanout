package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 42
    v3 := make(map[int]int)
    var v4 float32 = 0.0
    var v5 string = ((("" + "foo bar") + "hello") + ("hello" + (("hello" + "\n\t") + ((("hello" + "\n\t") + ((("a" + "\n\t") + ("" + "")) + "foo bar")) + ("foo bar" + (("hello" + "hello") + "foo bar"))))))
    v6 := []int{}
    var v7 int8 = 127
    v5 = v5
    v2 += v2
    var v8 int16 = -1
    v9 := GenS1{F0: (0 > (v2 - 9223372036854775807)), F1: (((!(0 > 0)) && (598 <= 18446744073709551615)) && (((9223372036854775807 < 313) && (true || ((v2 << ((18446744073709551615) & 63)) < (v2 & -1)))) || ((v7 == (v7 & v7)) && (!(760 > 297))))), F2: 834}
    v3[32768] = 32767
    if false {
        var v10 uint32 = 302
        var v11 int8 = (^(v7))
        _ = v10
        _ = v11
    } else {
        switch 128 {
        case 127:
            v7 = (^(v7))
        default:
            var v12 int64 = (v2 / (((v2 | (v2 * (v2 & v2)))) | 1))
            _ = v12
        }
        if (true || ((v2 << ((18446744073709551615) & 63)) < (v2 | (^(v2))))) {
            v13 := GenS2{F0: 624, F1: (v5 + v5), F2: 810}
            switch 42 {
            case 263:
                for fi14 := 0; fi14 < 14; fi14++ {
                    v7 += (v7 - (^(v7)))
                    for fi15 := 0; fi15 < 1; fi15++ {
                        _ = len(v3)
                        v16 := GenS1{F0: false, F1: ((!(42 == 2147483647)) || true), F2: 721}
                        _ = fi15
                        _ = v16
                    }
                    if len(v6) > 0 {
                        v6[0] = fi14
                    }
                    _ = fi14
                }
            case -1:
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint64 = 18446744073709551615
                    var v19 int64 = (v2 % ((v2) | 1))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            default:
                v20 := GenS3{F0: (v2 - v2), F1: (v7 % (((v7 ^ (v7 - v7))) | 1)), F2: 42}
                _ = v20
            }
            if false {
                fmt.Println(v5, v4, v7)
                switch v5 {
                case "\n\t":
                    _ = len(v6)
                case "foo bar":
                    v5 = v5
                case "a":
                    if true {
                        var v21 uint = 426
                        _ = v21
                    } else {
                        if ((v2 + (-(v2))) >= (v2 + v2)) {
                            if false {
                                var v22 int16 = (v8 << ((255) & 63))
                                switch (v5 + (v5 + ("a" + ("\n\t" + (v5 + v5))))) {
                                case "hello":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                case "":
                                    var v23 uint32 = 4294967295
                                    _ = v23
                                default:
                                    var v24 uint8 = 200
                                    _ = v24
                                }
                                _ = v22
                            }
                            var v25 uint = 2
                            v26 := GenS4{F0: 282, F1: 2, F2: -626}
                            _ = v25
                            _ = v26
                        } else {
                            fmt.Println(v7, v2)
                            var v27 uint32 = 25
                            _ = v27
                        }
                        var v28 float64 = 1e10
                        _ = v28
                    }
                }
                v2 += (v2 | (-(v2)))
            } else {
                if (0 == v8) {
                    var v29 int8 = v7
                    if (true && false) {
                        if true {
                            v6 = append(v6, 1)
                        }
                    }
                    _ = v29
                }
                fmt.Println(v8)
            }
            _ = v13
        } else {
            v30 := GenS5{F0: (v2 | 0)}
            _ = v30
        }
        v7 -= -1
    }
    switch "foo bar" {
    case "\n\t":
        for fi31 := 0; fi31 < 7; fi31++ {
            if (false && true) {
                v5 = "\n\t"
                if (v2 > v2) {
                    v4 += ((float32(1e10) / 3.14159) * ((3.14159 + -2.25) * v4))
                    var v32 float32 = v4
                    var v33 string = ""
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint32 = 1
                    if (4294967295 < 565) {
                        var v35 float32 = float32(float64(0.0))
                        fmt.Println(fi31, v2, v7)
                        v36 := GenS5{F0: 1}
                        _ = v35
                        _ = v36
                    }
                    _ = v34
                }
                if false {
                    fmt.Println(v7, v2, v2)
                } else {
                    var v37 uint = 280
                    _ = v37
                }
            }
            v38 := GenS3{F0: v2, F1: (v7 ^ (-(v7))), F2: 128}
            _ = fi31
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 11; fi39++ {
            v40 := v1[65535]
            fmt.Println(v40, v7)
            var v41 int32 = 1
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    switch 605 {
    case -2147483648:
        v42 := GenS3{F0: v2, F1: v7, F2: 157}
        _ = v42
    case -826:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v43 := GenS4{F0: 807, F1: 29, F2: 1}
    v3[9223372036854775807] = 32768
    if false {
        for fi44 := 0; fi44 < 1; fi44++ {
            var v45 uint8 = 42
            var v46 string = "a"
            _ = fi44
            _ = v45
            _ = v46
        }
        var v47 float64 = 0.1
        _ = v47
    } else {
        for fi48 := 0; fi48 < 2; fi48++ {
            v49 := GenS6{F0: v7, F1: 18446744073709551615, F2: (fi48 >> ((65535) & 63))}
            v8 += 447
            _ = fi48
            _ = v49
        }
        for fi50 := 0; fi50 < 14; fi50++ {
            for fi51 := 0; fi51 < 18; fi51++ {
                switch v5 {
                case "a":
                    switch ("hello" + v5) {
                    case "foo bar":
                        fmt.Println(fi51, v8, v4)
                    case "a":
                        var v52 uint8 = 0
                        _ = v52
                    case "":
                        var v53 uint32 = 2147483648
                        _ = v53
                    }
                case "hello":
                    var v54 int16 = 1000
                    _ = v54
                }
                _ = fi51
            }
            v5 = ("" + v5)
            for fi55 := 0; fi55 < 16; fi55++ {
                var v56 int = (fi55 >> ((4294967295) & 63))
                if (1 < 199) {
                    var v57 uint16 = 266
                    if len(v6) > 0 {
                        v6[0] = 2147483647
                    }
                    _ = v57
                }
                _ = fi55
                _ = v56
            }
            _ = fi50
        }
        if false {
            for fi58 := 0; fi58 < 18; fi58++ {
                var v59 int8 = v7
                var v60 int = (fi58 >> ((42) & 63))
                var v61 float64 = ((1e10 / 1.0) + -2.25)
                _ = fi58
                _ = v59
                _ = v60
                _ = v61
            }
            fmt.Println(v7, v8, v7)
            for fi62 := 0; fi62 < 18; fi62++ {
                if len(v6) > 0 {
                    v6[0] = -9223372036854775808
                }
                _ = fi62
            }
        }
    }
    v8 = -551
    switch 32767 {
    case 1:
        v7 = 42
    case 9223372036854775807:
        v2 += (v2 * v2)
    default:
        fmt.Println(v8, v5, v4)
    }
    var v63 uint8 = 245
    v64 := GenS4{F0: 163, F1: 0, F2: 2147483647}
    v63 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v7
    _ = v8
    _ = v9
}
