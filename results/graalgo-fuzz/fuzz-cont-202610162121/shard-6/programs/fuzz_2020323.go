package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int64; F2 uint }
type GenS4 struct { F0 int64; F1 int8; F2 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 int8 }
type GenS6 struct { F0 int; F1 uint32; F2 bool }
func genFn0(p0 int) uint8 {
    p0 = p0
    var v4 int64 = -9223372036854775808
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 float32 = 0.0
    var v2 bool = (42 <= -312)
    var v3 uint = 255
    v5 := []int{32768}
    v3 = 0
    fmt.Println(v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(fi7, fi6, v1)
            v8 := GenS1{F0: 128, F1: 255, F2: 128}
            _ = len(v5)
            _ = fi7
            _ = v8
        }
        switch fi6 {
        case -2147483648:
            _ = len(v5)
        case -1:
            fmt.Println(v1, v3, v3)
        case -9223372036854775808:
            switch -508 {
            case -9223372036854775808:
                switch ("" + "foo bar") {
                case "\n\t":
                    v9 := GenS1{F0: 127, F1: v3, F2: 228}
                    _ = v9
                case "foo bar":
                    v2 = (("a" == (("hello" + "") + "")) && true)
                }
            case 128:
                v10 := GenS2{F0: 42}
                _ = v10
            }
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    v3 -= 255
    if (42 != -1) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v11 uint64 = 2
        _ = v11
    } else {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS2{F0: 533}
            _ = fi12
            _ = v13
        }
        var v14 bool = false
        _ = v14
    }
    if len(v5) > 0 {
        v5[0] = 2147483647
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v3 = v3
        v2 = v2
        _ = fi15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v16 := GenS2{F0: 0}
    if v2 {
        var v17 int16 = 42
        for fi18 := 0; fi18 < 19; fi18++ {
            v5 = append(v5, (fi18 | fi18))
            v19 := GenS3{F0: v3, F1: 9223372036854775807, F2: v3}
            for fi20 := 0; fi20 < 17; fi20++ {
                if len(v5) > 0 {
                    v5[0] = 128
                }
                if v2 {
                    _ = len(v5)
                    v2 = false
                    var v21 float64 = (0.0 * 0.0)
                    fmt.Println(v17, v3)
                    if len(v5) > 0 {
                        v5[0] = -27
                    }
                    v21 += v21
                    if len(v5) > 0 {
                        v5[0] = (fi18 >> (((v3 << ((65535) & 63))) & 63))
                    }
                    switch (fi18 & 4294967295) {
                    case 1:
                        v22 := GenS4{F0: -9223372036854775808, F1: -1, F2: 42}
                        _ = v22
                    case -2147483648:
                        fmt.Println(v1)
                    case 794:
                        if v2 {
                            switch (fi20 << ((v3) & 63)) {
                            case -993:
                                v5 = append(v5, -1)
                            case 326:
                                switch "a" {
                                case "foo bar":
                                    v21 += ((-2.25 / v21) + v21)
                                case "\n\t":
                                    if len(v5) > 0 {
                                        v5[0] = 47
                                    }
                                case "hello":
                                    var v23 uint64 = 4294967295
                                    _ = v23
                                default:
                                    v24 := GenS4{F0: -9223372036854775808, F1: -38, F2: 2147483647}
                                    _ = v24
                                }
                            case 32768:
                                var v25 float64 = 1e10
                                _ = v25
                            default:
                                switch 9223372036854775807 {
                                case 32768:
                                    v5 = append(v5, fi20)
                                case 9223372036854775807:
                                    v26 := GenS2{F0: v17}
                                    _ = v26
                                case 2147483647:
                                    var v27 bool = v2
                                    _ = v27
                                }
                            }
                            v28 := GenS5{F0: v17, F1: 42, F2: 0}
                            v29 := GenS4{F0: 0, F1: 62, F2: -2147483648}
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 uint8 = 127
                        _ = v30
                    }
                    _ = v21
                } else {
                    if len(v5) > 0 {
                        v5[0] = 42
                    }
                    if true {
                        var v31 int8 = 77
                        var v32 int32 = -301
                        _ = v31
                        _ = v32
                    }
                    var v33 uint = (^(v3))
                    _ = v33
                }
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
        _ = v17
    } else {
        for fi34 := 0; fi34 < 2; fi34++ {
            for fi35 := 0; fi35 < 13; fi35++ {
                var v36 int = fi35
                if ((470 < 1) && v2) {
                    if false {
                        if len(v5) > 0 {
                            v5[0] = (v36 ^ fi35)
                        }
                        if v2 {
                            v2 = v2
                            var v37 uint32 = 47
                            var v38 int64 = 0
                            _ = v37
                            _ = v38
                        } else {
                            v39 := GenS1{F0: 128, F1: v3, F2: 207}
                            _ = v39
                        }
                    } else {
                        v3 -= 108
                    }
                    v3 = (v3 - v3)
                } else {
                    v1 = v1
                    if v2 {
                        var v40 bool = v2
                        _ = v40
                    }
                }
                _ = fi35
                _ = v36
            }
            for fi41 := 0; fi41 < 14; fi41++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi41
            }
            v2 = v2
            _ = fi34
        }
        fmt.Println(v3, v1, v1)
    }
    if (v1 != float32(1e10)) {
        var v42 uint8 = 128
        _ = v42
    } else {
        if ((3.14159 - 1.5) != (float64(float32((1.5 - (0.1 * -2.25)))) + float64(v1))) {
            v43 := GenS1{F0: 134, F1: v3, F2: 89}
            _ = v43
        } else {
            if len(v5) > 0 {
                v5[0] = 9223372036854775807
            }
            for fi44 := 0; fi44 < 6; fi44++ {
                if v2 {
                    var v45 float64 = 3.14159
                    for fi46 := 0; fi46 < 1; fi46++ {
                        switch -1 {
                        case 2147483647:
                            if (fi44 == (fi44 * fi44)) {
                                var v47 int32 = 547
                                _ = v47
                            }
                        case 0:
                            if v2 {
                                fmt.Println(v2, v3)
                                switch fi44 {
                                case -9223372036854775808:
                                    fmt.Println(fi46, fi46)
                                }
                                var v48 int16 = -32768
                                _ = v48
                            }
                        default:
                            var v49 int64 = -1
                            _ = v49
                        }
                        _ = fi46
                    }
                    _ = v45
                }
                var v50 int64 = 0
                _ = fi44
                _ = v50
            }
        }
        v1 = ((-2.25 * ((((((v1 + (v1 * v1)) + ((v1 + float32(0.0)) + (v1 * -2.25))) / 1.5) * 3.14159) - v1) / (v1 - (1e10 / ((-2.25 + 0.0) * v1))))) + v1)
        v51 := GenS6{F0: -934, F1: 455, F2: ((v1 / -2.25) >= ((v1 * (v1 - (((1e10 + 1e10) / 3.14159) + 3.14159))) - v1))}
        _ = v51
    }
    var v52 int = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v5
    _ = v52
}
