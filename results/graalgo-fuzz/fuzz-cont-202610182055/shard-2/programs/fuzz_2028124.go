package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint }
type GenS2 struct { F0 int64; F1 int32; F2 int }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 bool) uint16 {
    var v5 float32 = 1.5
    v6 := []int{-2147483648, 32768, 127}
    v7 := GenS1{F0: 42, F1: 127, F2: 18446744073709551615}
    var v8 uint = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{0}
    var v3 uint16 = 0
    var v4 int32 = 42
    if len(v2) > 0 {
        v2[0] = 92
    }
    var v9 float32 = 0.1
    fmt.Println(v1, v4, v3)
    v2 = append(v2, 247)
    switch "\n\t" {
    case "foo bar":
        _ = len(v2)
    case "\n\t":
        var v10 uint = 1
        _ = v10
    case "a":
        var v11 int = 32768
        _ = v11
    }
    v1 += (v1 % (((v1 | 1)) | 1))
    if len(v2) > 0 {
        v2[0] = 211
    }
    v12 := GenS1{F0: 0, F1: 15, F2: 42}
    if (65535 >= genFn0(v1, (-1 < 1))) {
        switch 128 {
        case -549:
            if false {
                switch "\n\t" {
                case "\n\t":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        v14 := GenS1{F0: 897, F1: 2, F2: 255}
                        for fi15 := 0; fi15 < 14; fi15++ {
                            switch (fi13 >> ((2) & 63)) {
                            case -548:
                                v9 += v9
                            case 994:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v16 uint = 2
                                _ = v16
                            }
                            fmt.Println(v4, v3)
                            v17 := GenS2{F0: -603, F1: (^(v4)), F2: (fi15 + int(v1))}
                            _ = fi15
                            _ = v17
                        }
                        if (!false) {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 bool = (true || true)
                                var v20 float64 = 1.5
                                switch fi13 {
                                case 32767:
                                    var v21 float64 = ((v20 * (v20 - v20)) / 3.14159)
                                    _ = v21
                                default:
                                    switch 1 {
                                    case 32768:
                                        v1 += v1
                                    case 127:
                                        v4 = (v4 * (v4 & v4))
                                    case -9223372036854775808:
                                        var v22 uint = 923
                                        _ = v22
                                    }
                                }
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        } else {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                var v24 bool = true
                                fmt.Println(v3)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v24 = false
                                _ = fi23
                                _ = v24
                            }
                            var v25 int32 = (v4 | v4)
                            v26 := GenS3{F0: (v4 << ((642) & 63))}
                            _ = v25
                            _ = v26
                        }
                        _ = fi13
                        _ = v14
                    }
                case "foo bar":
                    if (1 == 0) {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            var v28 uint = 65535
                            var v29 uint32 = 0
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    } else {
                        var v30 float64 = -2.25
                        var v31 int8 = 112
                        _ = v30
                        _ = v31
                    }
                case "":
                    v1 -= (v1 % (((v1 / (((v1 | v1)) | 1))) | 1))
                }
            }
        case 32767:
            var v32 uint8 = 128
            _ = v32
        default:
            if false {
                v33 := GenS1{F0: 42, F1: 127, F2: 929}
                _ = v33
            }
        }
        var v34 int16 = v1
        _ = v34
    }
    if len(v2) > 0 {
        v2[0] = 0
    }
    v1 = v1
    fmt.Println(v9, v9, v4)
    _ = len(v2)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
