package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 string; F1 uint32; F2 uint16 }
func genFn0() uint8 {
    v5 := GenS1{F0: 1, F1: 128, F2: 1}
    if true {
        var v6 int64 = -9223372036854775808
        fmt.Println(v6)
        _ = v6
    } else {
        var v7 uint32 = 644
        var v8 int32 = 42
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 32768, F1: 4294967295, F2: 969}
    _ = v5
    _ = v9
    return 128
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    var v3 float64 = 0.1
    v4 := []uint32{572, 4294967295, 1, 2}
    v4 = append(v4, 2)
    v10 := GenS2{F0: 42, F1: 1, F2: -32768}
    v11 := GenS2{F0: 255, F1: 1, F2: -711}
    fmt.Println(v3)
    switch 1 {
    case -308:
        var v12 uint = 4294967295
        _ = v12
    default:
        v3 += float64(1e10)
    }
    if (-5 > 42) {
        var v13 int = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        v2[fi14] = true
        switch (fi14 << ((18446744073709551615) & 63)) {
        case 128:
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS3{F0: ""}
                if false {
                    fmt.Println(v1)
                    var v17 uint8 = 30
                    if false {
                        fmt.Println(v3)
                    }
                    _ = v17
                } else {
                    v18 := GenS4{F0: false}
                    if true {
                        if (false || (false || (!(v3 >= v3)))) {
                            if (!(!true)) {
                                switch fi14 {
                                case 32768:
                                    var v19 bool = true
                                    _ = v19
                                case 42:
                                    _ = len(v4)
                                case 0:
                                    fmt.Println(v1, v1, fi15)
                                    v1 += (v1 * v1)
                                    v1 = v1
                                default:
                                    v1 = v1
                                }
                                v1 += 732
                                if (18446744073709551615 != 0) {
                                    if false {
                                        v1 += (^(v1))
                                        v1 -= 42
                                    } else {
                                        v4 = append(v4, 2)
                                        if true {
                                            v4 = append(v4, 65535)
                                            var v20 uint = 416
                                            var v21 bool = (false || (-1 > 98))
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 int64 = -9223372036854775808
                                            if len(v4) > 0 {
                                                v4[0] = 801
                                            }
                                            _ = v22
                                        }
                                    }
                                    var v23 uint = 42
                                    v23 = 42
                                    _ = v23
                                }
                            } else {
                                v3 += ((-2.25 * (v3 / 3.14159)) + ((((1.5 - (v3 / v3)) * -2.25) * v3) + 1e10))
                                var v24 bool = (149 == 212)
                                var v25 uint64 = 876
                                switch 42 {
                                case 4294967295:
                                    if v24 {
                                        v25 = 4294967295
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if v24 {
                                            v26 := GenS5{F0: v25, F1: -128}
                                            fmt.Println(v1, fi14)
                                            _ = v26
                                        }
                                    }
                                case 636:
                                    var v27 bool = (-851 == 42)
                                    _ = v27
                                case 127:
                                    v28 := GenS1{F0: 1, F1: fi14, F2: 553}
                                    _ = v28
                                default:
                                    var v29 uint8 = genFn0()
                                    _ = v29
                                }
                                v25 = 42
                                _ = v24
                                _ = v25
                            }
                            v30 := v2[fi14]
                            v30 = v30
                            _ = v30
                        }
                    }
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
        default:
            var v31 int16 = 1000
            _ = v31
        }
        _ = fi14
    }
    v32 := GenS6{F0: "\n\t", F1: 42, F2: 247}
    v33 := GenS4{F0: (0.0 <= (1.5 / float32(1e10)))}
    v3 -= v3
    var v34 uint16 = 32767
    switch 1 {
    case 9223372036854775807:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 759:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
