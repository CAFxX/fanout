package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint }

func main() {
    var v1 uint8 = 199
    var v2 float64 = ((3.14159 * 1.5) / 1.0)
    var v3 uint8 = v1
    v4 := []string{("\n\t" + "")}
    var v5 uint64 = 397
    v6 := []uint32{514, 923}
    v5 += (^(v5))
    if (976 == 2147483648) {
        if (2147483648 <= 65535) {
            for fi7 := 0; fi7 < 4; fi7++ {
                fmt.Println(v5, v3, v5)
                v1 = (v1 / ((v3) | 1))
                _ = fi7
            }
        } else {
            var v8 uint8 = (v1 / ((v3) | 1))
            v2 -= (1.5 + float64(float32((-2.25 * (v2 - ((float64(3.14159) + (((float64(float32(v2)) * v2) * v2) - 1.5)) / v2))))))
            if (0 != 1) {
                v8 += v3
                v9 := GenS1{F0: 0, F1: 42}
                _ = v9
            }
            _ = v8
        }
    }
    var v10 uint32 = 65535
    v11 := GenS2{F0: -2147483648, F1: v5, F2: 583}
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v10, v5)
            if (!((fi13 / ((-329) | 1)) <= (^(fi12)))) {
                v1 -= v3
                fmt.Println(v2, v1, v5)
                v14 := GenS3{F0: (^(v10))}
                _ = v14
            } else {
                v15 := GenS3{F0: 65}
                v4 = append(v4, ("\n\t" + ""))
                v3 += 127
                if (-603 == 1) {
                    v2 = ((v2 - v2) + ((1.5 / (((v2 + float64(3.14159)) * 0.0) / v2)) / 1.5))
                    var v16 int16 = int16(v5)
                    v17 := GenS4{F0: 313}
                    _ = v16
                    _ = v17
                } else {
                    var v18 uint8 = 0
                    var v19 uint8 = v18
                    switch fi13 {
                    case 32767:
                        v10 = (v10 * 65535)
                    case 1:
                        v5 -= (v5 >> ((65535) & 63))
                    case -2147483648:
                        v20 := GenS5{F0: -27, F1: 12}
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
                v3 = (v3 >> ((2) & 63))
                v2 -= v2
                _ = v15
            }
            if (!false) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v3 += 128
                v21 := GenS2{F0: 0, F1: (v5 | 0), F2: 46}
                _ = v21
            } else {
                var v22 uint = 255
                if (v2 != 1e10) {
                    fmt.Println(fi12, v22)
                    switch 0 {
                    case -380:
                        fmt.Println(v10, v3, fi12)
                    case 475:
                        v22 += v22
                    case -191:
                        var v23 uint32 = uint32(v3)
                        _ = v23
                    default:
                        if false {
                            v10 = (^(v10))
                            if true {
                                v24 := GenS5{F0: 1, F1: 42}
                                var v25 uint16 = 1
                                _ = v24
                                _ = v25
                            } else {
                                var v26 bool = ((true || ((3.14159 * 1e10) == 3.14159)) || true)
                                var v27 int32 = 2147483647
                                _ = v26
                                _ = v27
                            }
                            v10 -= uint32(v22)
                            var v28 int32 = 42
                            var v29 uint32 = (v10 >> (((^(v22))) & 63))
                            _ = v28
                            _ = v29
                        } else {
                            var v30 uint8 = (v3 * (v1 ^ 2))
                            switch "hello" {
                            case "foo bar":
                                if false {
                                    v31 := GenS2{F0: -2147483648, F1: (v5 << ((v22) & 63)), F2: v22}
                                    _ = v31
                                } else {
                                    v1 = (^(v3))
                                    v32 := GenS5{F0: 1, F1: -51}
                                    v22 = (v22 + (v22 * 42))
                                    if true {
                                        fmt.Println(v10, v3, v2)
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    v3 = 255
                                    _ = v32
                                }
                            }
                            _ = v30
                        }
                    }
                    _ = len(v6)
                }
                _ = v22
            }
            _ = fi13
        }
        _ = fi12
    }
    v1 = 95
    v2 -= 3.14159
    switch 32768 {
    case 4294967295:
        v33 := GenS6{F0: 255, F1: v1, F2: uint(v1)}
        _ = v33
    default:
        for fi34 := 0; fi34 < 6; fi34++ {
            v2 -= (v2 * float64((-2.25 * float32(v2))))
            _ = fi34
        }
    }
    if (false || (-2147483648 <= 544)) {
        var v35 bool = ((^(v10)) <= v10)
        v36 := GenS1{F0: 8, F1: 32}
        var v37 int = 4294967295
        _ = v35
        _ = v36
        _ = v37
    }
    if true {
        v5 += 18446744073709551615
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
