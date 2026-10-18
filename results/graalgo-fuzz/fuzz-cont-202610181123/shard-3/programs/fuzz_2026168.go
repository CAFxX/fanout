package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 int16 }
type GenS3 struct { F0 bool; F1 uint64; F2 int64 }

func main() {
    v1 := []int{-919, -2147483648, 4294967295}
    v2 := []uint32{165}
    var v3 uint32 = 528
    var v4 int64 = 0
    var v5 uint16 = 670
    var v6 int8 = 1
    v1 = append(v1, int(v5))
    switch 127 {
    case 9223372036854775807:
        v7 := GenS1{F0: 2}
        _ = v7
    default:
        if (-63 >= (^(v6))) {
            for fi8 := 0; fi8 < 7; fi8++ {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 bool = true
                    _ = fi9
                    _ = v10
                }
                switch fi8 {
                case 8:
                    v11 := GenS1{F0: 127}
                    _ = v11
                case 128:
                    var v12 uint32 = (^(v3))
                    _ = v12
                }
                if len(v1) > 0 {
                    v1[0] = -9223372036854775808
                }
                _ = fi8
            }
            fmt.Println(v5, v5)
        }
    }
    if false {
        v6 = (v6 >> ((4294967295) & 63))
    } else {
        switch "hello" {
        case "hello":
            fmt.Println(v5, v4, v5)
        default:
            if ((127 > 255) && (v3 > (v3 - v3))) {
                if true {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        v5 -= 0
                        if true {
                            v6 = v6
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi13
                    }
                    fmt.Println(v3, v5)
                    if false {
                        fmt.Println(v6)
                        switch 128 {
                        case -2147483648:
                            var v14 float64 = -2.25
                            _ = v14
                        case 128:
                            v1 = append(v1, 633)
                        case 9223372036854775807:
                            var v15 float64 = 0.1
                            _ = v15
                        }
                    } else {
                        if (false || (v6 > (v6 % ((v6) | 1)))) {
                            v3 = (v3 + v3)
                            switch -9223372036854775808 {
                            case -9223372036854775808:
                                var v16 uint64 = 524
                                _ = v16
                            case 32768:
                                if true {
                                    var v17 uint8 = 42
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = v17
                                } else {
                                    fmt.Println(v6, v4)
                                    v2 = append(v2, v3)
                                }
                            case -1:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v3 = (^(v3))
                            }
                            if false {
                                v18 := GenS1{F0: 193}
                                _ = v18
                            } else {
                                v4 = (-(v4))
                            }
                        }
                        v5 = 0
                        v6 += v6
                    }
                } else {
                    v2 = append(v2, v3)
                }
                v3 = 65535
                if (1 >= (^(v4))) {
                    var v19 uint64 = uint64(v5)
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 int64 = (-(v4))
                        _ = fi20
                        _ = v21
                    }
                    v4 -= v4
                    _ = v19
                }
            } else {
                var v22 uint8 = 255
                v3 -= (v3 ^ v3)
                v22 = v22
                _ = v22
            }
        }
    }
    v3 += 607
    var v23 uint64 = 624
    if false {
        for fi24 := 0; fi24 < 19; fi24++ {
            switch (-(fi24)) {
            case 755:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 616:
                if (!(-128 >= 0)) {
                    var v25 uint8 = 128
                    _ = v25
                } else {
                    v26 := GenS2{F0: (fi24 / (((fi24 * fi24)) | 1)), F1: 42, F2: -32768}
                    _ = v26
                }
            default:
                v27 := GenS3{F0: (3.14159 <= -2.25), F1: v23, F2: (v4 % ((v4) | 1))}
                _ = v27
            }
            _ = fi24
        }
        v28 := GenS1{F0: 45}
        _ = v28
    } else {
        for fi29 := 0; fi29 < 20; fi29++ {
            v3 += v3
            for fi30 := 0; fi30 < 7; fi30++ {
                var v31 float64 = (1.5 + -2.25)
                fmt.Println(fi29, v5, fi29)
                var v32 uint16 = v5
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi29
        }
        fmt.Println(v5, v6, v6)
    }
    var v33 int16 = 1000
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v23 -= (v23 | (v23 | (v23 << ((0) & 63))))
    for fi34 := 0; fi34 < 20; fi34++ {
        v23 += v23
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
