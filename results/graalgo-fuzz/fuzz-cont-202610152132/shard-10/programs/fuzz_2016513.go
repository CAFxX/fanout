package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }
type GenS3 struct { F0 int8; F1 int16 }
func genFn0(p0 int, p1 int8, p2 int8) string {
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int8 = p2
        var v10 uint8 = 127
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return ("" + "hello")
}

func main() {
    v1 := []uint32{155, 1, 171}
    var v2 float32 = (float32((0.0 * (0.1 * -2.25))) - -2.25)
    var v3 int64 = 577
    var v4 uint64 = 246
    var v5 uint8 = 2
    v6 := make(map[uint32]int)
    v3 -= 0
    var v11 float32 = float32(float64(-2.25))
    fmt.Println(v4)
    v11 = v11
    if (false || true) {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 14; fi13++ {
                if true {
                    var v14 uint16 = 2
                    if (1.5 < (1.5 / ((-2.25 + 1.5) * -2.25))) {
                        if true {
                            var v15 int32 = 2147483647
                            _ = v15
                        }
                        var v16 string = "hello"
                        _ = v16
                    }
                    if (false && false) {
                        switch 128 {
                        case 624:
                            if false {
                                v2 -= (((v2 - v2) / v2) / 1.5)
                                switch "" {
                                case "a":
                                    v2 = v2
                                default:
                                    v3 -= v3
                                }
                                switch fi12 {
                                case 0:
                                    if (825 < 140) {
                                        var v17 uint32 = 2147483648
                                        fmt.Println(v5, v2)
                                        if false {
                                            fmt.Println(v5, v14)
                                            _ = len(v1)
                                        }
                                        _ = v17
                                    } else {
                                        v1 = append(v1, 2147483648)
                                        var v18 uint8 = v5
                                        if true {
                                            fmt.Println(v14, fi12)
                                        }
                                        v19 := GenS1{F0: 1}
                                        _ = v18
                                        _ = v19
                                    }
                                default:
                                    v20 := GenS1{F0: (v14 >> ((uint(v5)) & 63))}
                                    _ = v20
                                }
                            }
                        }
                        switch fi12 {
                        case 32767:
                            v21 := GenS1{F0: v14}
                            _ = v21
                        default:
                            _ = len(v6)
                        }
                        fmt.Println(v14)
                    } else {
                        var v22 int32 = -2147483648
                        var v23 uint32 = 4294967295
                        v14 += v14
                        _ = v22
                        _ = v23
                    }
                    switch (fi12 + 2147483647) {
                    case 0:
                        var v24 uint16 = v14
                        _ = v24
                    case 1:
                        v3 = (v3 - -864)
                    case -267:
                        var v25 int64 = (v3 ^ -146)
                        _ = v25
                    default:
                        v6[42] = (fi12 >> ((2) & 63))
                    }
                    _ = v14
                }
                var v26 int16 = 217
                delete(v6, 115)
                _ = fi13
                _ = v26
            }
            _ = fi12
        }
    }
    if (!(false && (false && false))) {
        var v27 uint = 2
        v28 := GenS2{F0: 100, F1: -128, F2: v3}
        _ = v27
        _ = v28
    }
    var v29 uint16 = 608
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 int = (fi30 << ((18446744073709551615) & 63))
        fmt.Println(v2, v29)
        v32 := v6[42]
        _ = fi30
        _ = v31
        _ = v32
    }
    if ((4294967295 > 255) || (!(1 == 42))) {
        v33 := v6[248]
        v29 = 42
        v29 = 65535
        _ = v33
    }
    if (!false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println(v3, v3)
    for fi34 := 0; fi34 < 11; fi34++ {
        var v35 uint16 = v29
        v36 := GenS3{F0: 100, F1: int16(v35)}
        fmt.Println(fi34)
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 uint64 = (^(v4))
    var v38 string = "foo bar"
    var v39 string = v38
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
