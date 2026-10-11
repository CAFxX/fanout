package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 int64; F2 string }
type GenS5 struct { F0 uint; F1 int8; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483648, 2147483647}
    var v3 string = "\n\t"
    fmt.Println(v3)
    var v4 int64 = -1
    switch "foo bar" {
    case "hello":
        var v5 int16 = 32767
        _ = v5
    case "foo bar":
        fmt.Println(v4)
    case "a":
        for fi6 := 0; fi6 < 14; fi6++ {
            if (215 > 255) {
                var v7 float64 = float64((0.0 / 1.5))
                _ = v7
            }
            switch "" {
            case "\n\t":
                v4 = v4
            default:
                for fi8 := 0; fi8 < 6; fi8++ {
                    v3 = ((v3 + v3) + "a")
                    v4 += (v4 * v4)
                    var v9 uint = 2
                    if (1 < -28) {
                        var v10 int64 = int64(fi6)
                        v10 = (v10 >> ((v9) & 63))
                        _ = v10
                    }
                    _ = fi8
                    _ = v9
                }
            }
            _ = fi6
        }
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v3 = "a"
        _ = fi11
    }
    if false {
        var v12 float64 = -2.25
        if true {
            if (("\n\t" != v3) || false) {
                v12 -= v12
            }
            v13 := GenS1{F0: true}
            _ = v13
        }
        _ = v12
    }
    var v14 float64 = ((1.5 + (1.5 - 1e10)) + ((1.5 + 1e10) + -2.25))
    for fi15 := 0; fi15 < 10; fi15++ {
        v14 = (v14 * (v14 + v14))
        _ = fi15
    }
    if (42 < 1) {
        var v16 uint32 = 385
        _ = v16
    }
    switch v3 {
    case "hello":
        if (-48 < 34) {
            switch 32767 {
            case -2147483648:
                switch -2147483648 {
                case -2147483648:
                    if true {
                        fmt.Println(v3, v3, v4)
                        v4 = (-(v4))
                    } else {
                        var v17 int32 = 42
                        _ = v17
                    }
                }
            case 9223372036854775807:
                for fi18 := 0; fi18 < 16; fi18++ {
                    var v19 uint = 0
                    _ = fi18
                    _ = v19
                }
            case -1:
                v3 = v3
            default:
                switch 2147483647 {
                case 4294967295:
                    if true {
                        switch "" {
                        case "a":
                            v20 := GenS2{F0: 1, F1: 291, F2: int32(v4)}
                            _ = v20
                        case "hello":
                            v3 = v3
                        }
                        var v21 int = 127
                        v22 := GenS1{F0: ((v21 ^ v21) != (v21 >> ((18446744073709551615) & 63)))}
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(v14, v3)
                    }
                case 722:
                    var v23 float64 = 1.5
                    _ = v23
                case -1:
                    var v24 float64 = v14
                    _ = v24
                default:
                    fmt.Println(v3)
                }
            }
        } else {
            fmt.Println(v4, v4)
            v25 := GenS3{F0: v3, F1: 0}
            _ = v25
        }
    case "foo bar":
        delete(v1, 85)
    }
    if false {
        v26 := GenS2{F0: 4294967295, F1: 406, F2: -1}
        _ = v26
    } else {
        if true {
            v4 += v4
            v4 = -1
            for fi27 := 0; fi27 < 4; fi27++ {
                var v28 int64 = v4
                _ = fi27
                _ = v28
            }
        }
        for fi29 := 0; fi29 < 1; fi29++ {
            for fi30 := 0; fi30 < 5; fi30++ {
                v4 += v4
                if len(v2) > 0 {
                    v2[0] = 42
                }
                _ = fi30
            }
            v31 := GenS3{F0: ("foo bar" + v3), F1: 18446744073709551615}
            if len(v2) > 0 {
                v2[0] = 1
            }
            _ = fi29
            _ = v31
        }
        v4 = v4
    }
    switch v3 {
    case "foo bar":
        switch 0 {
        case 9223372036854775807:
            var v32 uint16 = 380
            _ = v32
        case 32768:
            v33 := GenS2{F0: 557, F1: 18446744073709551615, F2: 198}
            _ = v33
        default:
            for fi34 := 0; fi34 < 6; fi34++ {
                if ((v4 - v4) < v4) {
                    for fi35 := 0; fi35 < 5; fi35++ {
                        v4 += v4
                        v4 = 42
                        _ = fi35
                    }
                } else {
                    v36 := GenS1{F0: true}
                    _ = v36
                }
                v14 -= float64(3.14159)
                _ = fi34
            }
        }
    default:
        _ = len(v2)
    }
    if false {
        for fi37 := 0; fi37 < 11; fi37++ {
            v38 := v1[(fi37 - (fi37 / ((fi37) | 1)))]
            _ = fi37
            _ = v38
        }
        switch 167 {
        case 473:
            for fi39 := 0; fi39 < 12; fi39++ {
                delete(v1, fi39)
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 uint32 = 4294967295
                    delete(v1, fi40)
                    switch v3 {
                    case "\n\t":
                        var v42 float64 = (v14 / ((3.14159 + -2.25) + float64(((0.0 - 1.5) - 0.1))))
                        _ = v42
                    case "":
                        if (1.5 > -2.25) {
                            v43 := GenS4{F0: -32768, F1: (v4 / ((9223372036854775807) | 1)), F2: "hello"}
                            v14 -= v14
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v43
                        }
                    case "hello":
                        v14 -= 1.5
                    default:
                        v14 = 1.5
                    }
                    _ = fi40
                    _ = v41
                }
                _ = fi39
            }
        case 32767:
            v44 := GenS1{F0: false}
            _ = v44
        case -683:
            fmt.Println(v14, v14, v3)
        default:
            v14 -= v14
        }
        v14 = (1e10 / float64((0.1 / float32((1.5 * v14)))))
    } else {
        if (false && false) {
            for fi45 := 0; fi45 < 13; fi45++ {
                var v46 uint64 = 18446744073709551615
                _ = fi45
                _ = v46
            }
        } else {
            v1[-788] = v3
            v3 = v3
        }
        if ((v14 - (v14 - float64(0.0))) == (1.5 - float64(3.14159))) {
            if true {
                v47 := GenS2{F0: 1, F1: 1, F2: -163}
                v14 = ((v14 + v14) * (float64(3.14159) + v14))
                if (((float32(v14) / 3.14159) < (float32(float64(0.0)) + (0.1 * (0.1 * 0.1)))) && (42 == 909)) {
                    if (1 < 128) {
                        if (!((!false) || true)) {
                            var v48 string = (v3 + v3)
                            for fi49 := 0; fi49 < 6; fi49++ {
                                var v50 int16 = -32768
                                _ = fi49
                                _ = v50
                            }
                            _ = v48
                        } else {
                            for fi51 := 0; fi51 < 9; fi51++ {
                                v4 -= (-(v4))
                                var v52 int16 = 42
                                var v53 uint64 = 18446744073709551615
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                        }
                        v14 = v14
                    } else {
                        v14 = 0.1
                        v3 = v3
                        _ = len(v2)
                    }
                    switch (v3 + ("a" + (v3 + (v3 + (v3 + v3))))) {
                    case "a":
                        if false {
                            v4 -= 138
                        } else {
                            var v54 bool = false
                            if (!v54) {
                                var v55 float32 = ((3.14159 * 3.14159) * float32((v14 + (v14 / 0.1))))
                                for fi56 := 0; fi56 < 5; fi56++ {
                                    v14 -= 3.14159
                                    _ = fi56
                                }
                                _ = v55
                            }
                            _ = v54
                        }
                    case "foo bar":
                        v14 = v14
                    case "\n\t":
                        fmt.Println(v14, v14, v4)
                    }
                    v14 += -2.25
                } else {
                    for fi57 := 0; fi57 < 4; fi57++ {
                        if false {
                            for fi58 := 0; fi58 < 20; fi58++ {
                                if (false && (false || true)) {
                                    v14 -= v14
                                } else {
                                    fmt.Println(fi58)
                                }
                                _ = fi58
                            }
                        }
                        v3 = (v3 + "\n\t")
                        v59 := GenS1{F0: false}
                        _ = fi57
                        _ = v59
                    }
                    v4 -= (v4 >> ((0) & 63))
                    fmt.Println(v4, v3)
                }
                _ = v47
            } else {
                switch "" {
                case "foo bar":
                    var v60 uint8 = 128
                    _ = v60
                }
                var v61 int32 = -183
                var v62 uint8 = 225
                _ = v61
                _ = v62
            }
            if (!true) {
                v63 := GenS5{F0: 444, F1: 96, F2: 0}
                _ = v63
            } else {
                v4 = -659
            }
            for fi64 := 0; fi64 < 16; fi64++ {
                _ = len(v2)
                _ = fi64
            }
        } else {
            v65 := GenS3{F0: v3, F1: 65535}
            v66 := GenS6{F0: 178}
            fmt.Println(v14, v14, v4)
            _ = v65
            _ = v66
        }
    }
    var v67 int32 = 786
    v4 -= v4
    if false {
        v68 := GenS7{F0: 872, F1: v4}
        var v69 uint16 = 84
        _ = v68
        _ = v69
    }
    if len(v2) > 0 {
        v2[0] = 42
    }
    v14 = (((-2.25 / 0.1) / float64(0.1)) + -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v67
}
