package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 135
    v3 := make(map[uint32]int)
    var v4 int = -983
    var v5 int = (^(v4))
    var v6 uint16 = 843
    fmt.Println(v4)
    for fi7 := 0; fi7 < 20; fi7++ {
        v4 = 748
        _ = fi7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        switch v4 {
        case 32767:
            v9 := GenS1{F0: (^(v6)), F1: 255}
            _ = v9
        default:
            for fi10 := 0; fi10 < 18; fi10++ {
                _ = len(v3)
                v2 = (v2 * v2)
                var v11 int = fi10
                _ = fi10
                _ = v11
            }
        }
        _ = fi8
    }
    v12 := GenS1{F0: (v6 & v6), F1: 42}
    for fi13 := 0; fi13 < 7; fi13++ {
        v1[(v4 * v5)] = fi13
        if ((v5 < (^(v4))) && (82 < 100)) {
            if true {
                for fi14 := 0; fi14 < 13; fi14++ {
                    var v15 float32 = 0.1
                    _ = fi14
                    _ = v15
                }
                v5 += (fi13 - int(v2))
            }
            v16 := GenS2{F0: 1, F1: 9223372036854775807}
            if ((!true) || ((1.5 / -2.25) == (-2.25 * (float64(0.0) + 0.0)))) {
                var v17 int64 = 1
                _ = v17
            }
            _ = v16
        }
        v18 := GenS3{F0: 42, F1: -1}
        _ = fi13
        _ = v18
    }
    v6 -= (^(v6))
    fmt.Println(v6, v4, v6)
    switch int(v6) {
    case 127:
        var v19 int8 = -128
        _ = v19
    }
    v20 := []int{(-(v4))}
    var v21 int8 = 42
    if false {
        v22 := v1[int(v6)]
        for fi23 := 0; fi23 < 14; fi23++ {
            for fi24 := 0; fi24 < 7; fi24++ {
                v20 = append(v20, -742)
                if true {
                    var v25 uint32 = v2
                    if true {
                        var v26 uint = 249
                        if (2 > 0) {
                            var v27 uint32 = (v2 / ((v2) | 1))
                            var v28 int16 = 1000
                            if true {
                                _ = len(v3)
                                switch 0 {
                                case -431:
                                    fmt.Println(v5)
                                case 128:
                                    var v29 float64 = -2.25
                                    var v30 int = (v5 ^ fi24)
                                    v22 = v4
                                    _ = v29
                                    _ = v30
                                default:
                                    v21 -= v21
                                }
                                if (!true) {
                                    var v31 int64 = 278
                                    _ = v31
                                }
                                v32 := GenS1{F0: v6, F1: 127}
                                _ = v32
                            } else {
                                if len(v20) > 0 {
                                    v20[0] = fi23
                                }
                            }
                            _ = v27
                            _ = v28
                        }
                        fmt.Println(fi23, v5, v21)
                        _ = v26
                    } else {
                        v33 := GenS2{F0: 0, F1: -9223372036854775808}
                        v20 = append(v20, (v22 | fi24))
                        v34 := GenS3{F0: 1, F1: 42}
                        v35 := GenS4{F0: 18446744073709551615}
                        fmt.Println(v4, v2)
                        fmt.Println(v25)
                        var v36 float64 = (0.1 * 0.1)
                        var v37 float32 = 1e10
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    _ = v25
                }
                _ = fi24
            }
            _ = fi23
        }
        v4 = v5
        _ = v22
    } else {
        v5 += v5
        for fi38 := 0; fi38 < 15; fi38++ {
            v2 -= 65535
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            v4 += 610
            _ = fi38
        }
        var v39 uint32 = (v2 % ((2147483648) | 1))
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
