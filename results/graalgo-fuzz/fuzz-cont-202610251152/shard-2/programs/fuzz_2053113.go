package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint64 }
func genFn0(p0 int64, p1 bool, p2 int16) int64 {
    if (282 > 2) {
        v6 := []int{-9223372036854775808}
        fmt.Println(p2)
        _ = v6
    }
    switch 127 {
    case -1:
        fmt.Println(p2, p2, p2)
    case -9223372036854775808:
        p0 += p0
    case 9223372036854775807:
        fmt.Println(p2, p2)
    default:
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 bool = p1
            v9 := []int{int(p2), (^(fi7)), (fi7 & fi7)}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = 42
    v3 := []int{42, 373, 921}
    var v4 uint32 = 2147483647
    v5 := []int{-2147483648, 2147483647, -1, -2147483648}
    v10 := GenS1{F0: 42, F1: 4294967295}
    v11 := GenS2{F0: 396}
    switch 32768 {
    case 91:
        v12 := GenS3{F0: false, F1: 2}
        _ = v12
    case -398:
        var v13 int16 = v2
        _ = v13
    }
    var v14 uint64 = 254
    if ((0.1 / 3.14159) <= (1e10 * (1.5 / 0.1))) {
        v1 += v1
    }
    if (!(!false)) {
        v2 = (v2 >> ((65535) & 63))
        for fi15 := 0; fi15 < 1; fi15++ {
            if len(v5) > 0 {
                v5[0] = (fi15 / ((fi15) | 1))
            }
            if false {
                if false {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v4 -= (v4 | (v4 >> ((42) & 63)))
                        _ = fi16
                    }
                    for fi17 := 0; fi17 < 3; fi17++ {
                        v2 -= -32768
                        v18 := GenS4{F0: -89, F1: ((v2 + v2) >= 42), F2: genFn0(9223372036854775807, true, int16(v14))}
                        switch "" {
                        case "foo bar":
                            v19 := GenS1{F0: 1, F1: 2}
                            _ = v19
                        case "\n\t":
                            v20 := GenS1{F0: 42, F1: 18446744073709551615}
                            _ = v20
                        default:
                            var v21 int8 = -58
                            var v22 uint8 = 204
                            _ = v21
                            _ = v22
                        }
                        _ = fi17
                        _ = v18
                    }
                    if (127 <= 127) {
                        fmt.Println(v1, v1)
                    }
                } else {
                    if ((-2147483648 <= 1) || (!true)) {
                        for fi23 := 0; fi23 < 17; fi23++ {
                            var v24 float32 = 3.14159
                            _ = fi23
                            _ = v24
                        }
                        fmt.Println(v14, v2, v2)
                    } else {
                        var v25 uint32 = v4
                        _ = v25
                    }
                    v26 := GenS5{F0: 0, F1: v14}
                    for fi27 := 0; fi27 < 4; fi27++ {
                        v28 := GenS4{F0: 127, F1: (9223372036854775807 != -9223372036854775808), F2: 0}
                        v2 -= -888
                        _ = len(v5)
                        _ = fi27
                        _ = v28
                    }
                    _ = v26
                }
            } else {
                v29 := GenS5{F0: -9223372036854775808, F1: v14}
                var v30 int32 = -327
                _ = v29
                _ = v30
            }
            v31 := GenS2{F0: v2}
            _ = fi15
            _ = v31
        }
        fmt.Println(v14, v2)
    } else {
        if (true || false) {
            v32 := GenS1{F0: -128, F1: 75}
            _ = v32
        } else {
            v1 = v1
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = (v4 & v4)
    v33 := GenS4{F0: -1, F1: (1 <= -128), F2: -1}
    switch "hello" {
    case "hello":
        if false {
            var v34 int64 = 9223372036854775807
            var v35 float64 = ((1.5 / (3.14159 - 0.1)) + 1.5)
            switch -535 {
            case 0:
                var v36 float32 = (v1 - float32(0.0))
                _ = v36
            case -817:
                for fi37 := 0; fi37 < 14; fi37++ {
                    fmt.Println(v14)
                    v38 := GenS6{F0: 42, F1: -32768}
                    v34 += v34
                    _ = fi37
                    _ = v38
                }
            }
            _ = v34
            _ = v35
        }
    case "foo bar":
        var v39 int32 = 0
        _ = v39
    case "\n\t":
        if (v4 != 72) {
            v4 = (v4 >> ((2) & 63))
            if (int(v4) >= 2147483647) {
                var v40 int = 32767
                for fi41 := 0; fi41 < 12; fi41++ {
                    if (!(!false)) {
                        var v42 uint16 = 58
                        var v43 uint8 = 128
                        switch "" {
                        case "":
                            v5 = append(v5, int(v14))
                        case "\n\t":
                            for fi44 := 0; fi44 < 5; fi44++ {
                                v42 -= (v42 * (v42 & v42))
                                _ = fi44
                            }
                        default:
                            if true {
                                if ((v1 > v1) && true) {
                                    fmt.Println(v42, fi41, v43)
                                    var v45 bool = false
                                    v45 = v45
                                    _ = v45
                                }
                                if (42 <= 0) {
                                    var v46 uint = 18446744073709551615
                                    var v47 uint16 = v42
                                    _ = v46
                                    _ = v47
                                }
                                v48 := GenS7{F0: 419, F1: 9223372036854775807}
                                _ = v48
                            } else {
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    v14 += v14
                                    v50 := GenS8{F0: (v14 >> ((2) & 63))}
                                    _ = fi49
                                    _ = v50
                                }
                                var v51 uint8 = (^(v43))
                                _ = v51
                            }
                        }
                        _ = v42
                        _ = v43
                    }
                    for fi52 := 0; fi52 < 19; fi52++ {
                        v53 := GenS4{F0: -9, F1: (255 == uint16(v14)), F2: 983}
                        _ = fi52
                        _ = v53
                    }
                    _ = fi41
                }
                v40 = (v40 ^ v40)
                _ = v40
            }
            switch 32768 {
            case 9223372036854775807:
                var v54 float32 = ((v1 - v1) - v1)
                _ = v54
            default:
                v55 := GenS7{F0: 4294967295, F1: -1}
                _ = v55
            }
        }
    default:
        if (!true) {
            if (!((381 >= 359) || ("" != "\n\t"))) {
                v3 = append(v3, 4294967295)
                v14 = v14
            } else {
                switch 0 {
                case 258:
                    if len(v5) > 0 {
                        v5[0] = 455
                    }
                case 2147483647:
                    v14 = (v14 << ((691) & 63))
                case 4294967295:
                    var v56 int16 = (v2 ^ int16(v4))
                    _ = v56
                default:
                    var v57 bool = (4294967295 != 2)
                    _ = v57
                }
                var v58 uint = 59
                var v59 int64 = 857
                _ = v58
                _ = v59
            }
            _ = len(v3)
            v60 := GenS4{F0: -128, F1: (!(255 > 344)), F2: -9223372036854775808}
            _ = v60
        }
    }
    var v61 uint64 = v14
    if false {
        var v62 uint8 = 145
        for fi63 := 0; fi63 < 14; fi63++ {
            v64 := GenS5{F0: -266, F1: (v61 * v61)}
            var v65 uint64 = (v61 >> ((18446744073709551615) & 63))
            _ = fi63
            _ = v64
            _ = v65
        }
        _ = v62
    }
    if len(v5) > 0 {
        v5[0] = -309
    }
    fmt.Println(v1, v14, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v61
}
