package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 uint32; F1 int64 }
func genFn0(p0 string, p1 uint, p2 uint64) int {
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int16 = 634
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: (false && false), F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483648
    v3 := make(map[int]string)
    var v4 int64 = -1
    v8 := GenS1{F0: v2, F1: (489 != 42), F2: 18446744073709551615}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v4 = v4
            var v11 uint64 = 0
            v12 := v3[(fi9 / ((fi10) | 1))]
            v13 := GenS1{F0: v2, F1: (("foo bar" + "") == ("" + (((((v12 + ("\n\t" + "\n\t")) + (("hello" + "hello") + ("" + ((v12 + v12) + (("a" + (v12 + "a")) + ""))))) + v12) + (v12 + "\n\t")) + (v12 + v12)))), F2: uint64(v2)}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi9
    }
    v14 := GenS1{F0: 42, F1: false, F2: 0}
    switch "\n\t" {
    case "a":
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint16 = 111
            var v17 float64 = float64((1e10 - float32(1.5)))
            var v18 int64 = (v4 ^ v4)
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    case "\n\t":
        for fi19 := 0; fi19 < 10; fi19++ {
            if (4294967295 > 0) {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 int64 = v4
                    v21 = 42
                    v4 = v4
                    if true {
                        v21 = v4
                        v2 = v2
                        v22 := GenS2{F0: int16(fi19)}
                        var v23 bool = (!true)
                        _ = v22
                        _ = v23
                    } else {
                        if true {
                            delete(v3, fi20)
                        }
                        v2 = (v2 << ((121) & 63))
                    }
                    if (3.14159 > 0.0) {
                        switch (fi19 + genFn0(("\n\t" + ((("a" + ("\n\t" + "\n\t")) + "foo bar") + (("foo bar" + "\n\t") + "a"))), 42, 4294967295)) {
                        case 4294967295:
                            fmt.Println(v2, v21, v2)
                        case -2147483648:
                            var v24 float32 = float32(((0.1 * 0.1) + 1.5))
                            _ = v24
                        default:
                            if (true && true) {
                                fmt.Println(fi20)
                            }
                        }
                        v21 = (^(v4))
                        v2 = v2
                    } else {
                        var v25 int8 = 100
                        if len(v1) > 0 {
                            v1[0] = v2
                        }
                        fmt.Println(v25, v21, v25)
                        v21 -= (v21 + 1)
                        if (v2 == v2) {
                            switch (fi19 & (fi20 - (-(fi19)))) {
                            case 127:
                                var v26 uint16 = 65535
                                _ = v26
                            case 42:
                                v27 := GenS1{F0: v2, F1: true, F2: 199}
                                _ = v27
                            }
                            var v28 uint16 = 1
                            var v29 uint16 = v28
                            _ = v28
                            _ = v29
                        } else {
                            var v30 uint32 = v2
                            _ = v30
                        }
                        v31 := GenS3{F0: (!true), F1: int16(v4), F2: -814}
                        v32 := GenS1{F0: v2, F1: true, F2: 230}
                        var v33 float64 = (0.1 + 1.5)
                        _ = v25
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    fmt.Println(fi20)
                    v34 := v3[(fi19 & (fi20 >> ((42) & 63)))]
                    switch -2147483648 {
                    case 1:
                        var v35 int64 = int64(v2)
                        _ = v35
                    }
                    var v36 int16 = -109
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                    if len(v1) > 0 {
                        v1[0] = (^(v2))
                    }
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    switch v34 {
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        switch v34 {
                        case "a":
                            if false {
                                if false {
                                    var v37 int64 = (v21 % ((v21) | 1))
                                    _ = v37
                                }
                                v38 := GenS1{F0: v2, F1: true, F2: 42}
                                v34 = v34
                                _ = v38
                            }
                        case "":
                            if len(v1) > 0 {
                                v1[0] = v2
                            }
                        case "foo bar":
                            if (2 >= 703) {
                                switch 175 {
                                case 9223372036854775807:
                                    v36 = (^(v36))
                                }
                            } else {
                                v39 := GenS4{F0: 113, F1: 42, F2: 0}
                                _ = v39
                            }
                        default:
                            v40 := GenS2{F0: v36}
                            _ = v40
                        }
                    default:
                        v36 = (v36 << ((49) & 63))
                    }
                    _ = fi20
                    _ = v21
                    _ = v34
                    _ = v36
                }
                for fi41 := 0; fi41 < 17; fi41++ {
                    fmt.Println(fi41)
                    v42 := v3[fi41]
                    _ = fi41
                    _ = v42
                }
                for fi43 := 0; fi43 < 6; fi43++ {
                    var v44 int8 = 0
                    var v45 int32 = -791
                    switch genFn0("", 739, 2) {
                    case 423:
                        switch "" {
                        case "":
                            var v46 float32 = (3.14159 - 0.1)
                            _ = v46
                        default:
                            v45 += (v45 / ((42) | 1))
                        }
                    case 4294967295:
                        v47 := GenS5{F0: 0, F1: ((v4 | v4) != v4)}
                        _ = v47
                    default:
                        v4 = (v4 / ((42) | 1))
                    }
                    var v48 int8 = int8(fi19)
                    _ = fi43
                    _ = v44
                    _ = v45
                    _ = v48
                }
            } else {
                v4 += (^(v4))
                if ((1.5 * 1e10) == (((3.14159 * (0.1 / 0.1)) * 0.1) + (0.1 - 3.14159))) {
                    switch (fi19 * (^(fi19))) {
                    case 42:
                        v4 -= -167
                    case -307:
                        v1 = append(v1, v2)
                    default:
                        v4 -= v4
                    }
                    if (18446744073709551615 > 239) {
                        var v49 bool = ((368 >= 0) && false)
                        v4 += (v4 | (v4 + v4))
                        v50 := GenS3{F0: v49, F1: 1000, F2: -279}
                        _ = v49
                        _ = v50
                    }
                }
                var v51 uint16 = 1
                _ = v51
            }
            _ = len(v1)
            _ = fi19
        }
    case "foo bar":
        if (65535 < uint(v2)) {
            _ = len(v1)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = 42
    for fi52 := 0; fi52 < 20; fi52++ {
        if true {
            var v53 uint64 = 1
            for fi54 := 0; fi54 < 10; fi54++ {
                v53 = 0
                _ = fi54
            }
            _ = v53
        } else {
            _ = len(v3)
        }
        _ = fi52
    }
    var v55 uint8 = 255
    v56 := GenS3{F0: (804 == 42), F1: 754, F2: 32767}
    for fi57 := 0; fi57 < 7; fi57++ {
        for fi58 := 0; fi58 < 5; fi58++ {
            v55 = (v55 - v55)
            _ = fi58
        }
        if false {
            v2 += (v2 | 947)
            if (!false) {
                for fi59 := 0; fi59 < 1; fi59++ {
                    v4 = 0
                    fmt.Println(fi57, v55, fi57)
                    v60 := GenS4{F0: 1, F1: uint64(v2), F2: uint16(v2)}
                    _ = fi59
                    _ = v60
                }
                v2 -= v2
                var v61 uint8 = v55
                _ = v61
            }
            var v62 int8 = 34
            _ = v62
        } else {
            if false {
                v2 = (v2 + uint32(v55))
            }
        }
        v4 -= 9223372036854775807
        _ = fi57
    }
    for fi63 := 0; fi63 < 12; fi63++ {
        if len(v1) > 0 {
            v1[0] = (v2 % ((v2) | 1))
        }
        for fi64 := 0; fi64 < 7; fi64++ {
            var v65 float32 = ((float32(((1e10 + float64(1e10)) / 3.14159)) / (1.5 + (1e10 - float32(-2.25)))) / (float32((3.14159 + 3.14159)) * (0.0 - 0.1)))
            _ = fi64
            _ = v65
        }
        if true {
            fmt.Println(fi63)
            var v66 bool = (!((127 > 0) || true))
            _ = v66
        } else {
            for fi67 := 0; fi67 < 7; fi67++ {
                v2 = v2
                v1 = append(v1, 2)
                if false {
                    v4 = v4
                }
                v2 -= uint32(v55)
                var v68 int8 = 120
                _ = fi67
                _ = v68
            }
        }
        _ = fi63
    }
    v69 := GenS6{F0: (v2 | (v2 >> ((730) & 63))), F1: v4}
    v1 = append(v1, v2)
    var v70 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v69
    _ = v70
    _ = v8
}
