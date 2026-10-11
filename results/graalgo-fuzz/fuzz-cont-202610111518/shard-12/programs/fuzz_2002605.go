package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 float64 = (0.1 * 1e10)
    var v2 uint8 = 0
    v3 := []int{2147483647, 767}
    v4 := []int{42, 32768}
    v5 := GenS1{F0: 2}
    v2 += 42
    fmt.Println(v1, v1)
    v6 := GenS1{F0: 255}
    for fi7 := 0; fi7 < 16; fi7++ {
        if true {
            switch fi7 {
            case 1:
                if true {
                    _ = len(v3)
                    for fi8 := 0; fi8 < 1; fi8++ {
                        if (float32(1e10) <= ((3.14159 + 1.5) / -2.25)) {
                            if true {
                                v9 := GenS2{F0: 1, F1: 2, F2: -9223372036854775808}
                                fmt.Println(v1, v2)
                                switch ("\n\t" + "foo bar") {
                                case "\n\t":
                                    v2 = 0
                                }
                                var v10 float64 = v1
                                _ = v10
                                _ = v9
                            } else {
                                var v11 int64 = -9223372036854775808
                                _ = v11
                            }
                            if (!false) {
                                _ = len(v3)
                            } else {
                                var v12 int16 = 1000
                                if (127 == (^(v2))) {
                                    v4 = append(v4, fi8)
                                    v12 = (v12 - int16(fi7))
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                } else {
                                    _ = len(v4)
                                    fmt.Println(fi8, v1)
                                }
                                if (65535 != 2) {
                                    var v13 bool = (!true)
                                    fmt.Println(v12, v2)
                                    fmt.Println(v1, v13, v2)
                                    _ = v13
                                } else {
                                    if false {
                                        v14 := GenS2{F0: 583, F1: 1, F2: -740}
                                        _ = v14
                                    } else {
                                        v15 := GenS1{F0: 42}
                                        v16 := GenS2{F0: 402, F1: 17, F2: 42}
                                        switch (fi8 / (((-(fi7))) | 1)) {
                                        case 0:
                                            v1 = (((v1 / v1) / v1) * (((-2.25 + 1.5) * 1.5) + (((v1 + (v1 * v1)) + (0.1 + v1)) + 1.5)))
                                        case 127:
                                            v12 = (v12 % (((v12 % ((0) | 1))) | 1))
                                        default:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        _ = v15
                                        _ = v16
                                    }
                                }
                                _ = v12
                            }
                            v3 = append(v3, 0)
                        } else {
                            if true {
                                switch ("" + "") {
                                case "\n\t":
                                    var v17 int64 = 9223372036854775807
                                    _ = v17
                                default:
                                    v2 += (v2 % ((v2) | 1))
                                }
                            } else {
                                var v18 int = (fi7 / ((-2147483648) | 1))
                                _ = v18
                            }
                            fmt.Println(fi8, fi7)
                            var v19 float32 = 0.0
                            v2 += (v2 << ((12) & 63))
                            _ = v19
                        }
                        _ = fi8
                    }
                }
            default:
                switch (fi7 * (fi7 + fi7)) {
                case 320:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v2 = (v2 >> ((528) & 63))
                        _ = fi20
                    }
                }
            }
            var v21 int32 = 42
            _ = v21
        } else {
            v1 += (1e10 * (3.14159 / (0.0 / v1)))
        }
        for fi22 := 0; fi22 < 15; fi22++ {
            v2 -= (v2 ^ v2)
            _ = fi22
        }
        v1 += v1
        _ = fi7
    }
    var v23 int8 = 127
    v4 = append(v4, 42)
    var v24 uint16 = 255
    v2 = v2
    for fi25 := 0; fi25 < 15; fi25++ {
        v1 += 1e10
        v23 = v23
        _ = fi25
    }
    v2 -= v2
    if (v2 == (v2 * v2)) {
        var v26 uint8 = (^(v2))
        _ = v26
    }
    v27 := GenS3{F0: -1}
    v24 += (v24 / ((255) | 1))
    _ = len(v3)
    v1 = v1
    if false {
        var v28 uint64 = 690
        _ = v28
    } else {
        for fi29 := 0; fi29 < 18; fi29++ {
            if len(v3) > 0 {
                v3[0] = fi29
            }
            switch "hello" {
            case "hello":
                fmt.Println(v24, v24, fi29)
            case "a":
                fmt.Println(v24, v24, v23)
            case "foo bar":
                var v30 uint = 42
                _ = v30
            }
            _ = fi29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
