package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int16 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{"foo bar"}
    var v2 float32 = 1e10
    var v3 uint16 = 1
    var v4 int64 = -308
    v5 := make(map[uint32]uint32)
    v6 := []int{-818, 32767}
    v3 = (v3 % ((42) | 1))
    fmt.Println(v4)
    switch "hello" {
    case "hello":
        for fi7 := 0; fi7 < 14; fi7++ {
            _ = len(v5)
            v4 -= (v4 - (-(v4)))
            _ = fi7
        }
    case "":
        if true {
            for fi8 := 0; fi8 < 7; fi8++ {
                v5[521] = 4294967295
                _ = fi8
            }
        } else {
            v2 -= v2
            for fi9 := 0; fi9 < 13; fi9++ {
                v4 = v4
                if (203 == 286) {
                    v10 := GenS1{F0: 0, F1: (fi9 ^ 32768), F2: 1000}
                    for fi11 := 0; fi11 < 5; fi11++ {
                        if false {
                            v2 -= v2
                            switch fi11 {
                            case 0:
                                switch "" {
                                case "":
                                    delete(v5, 65535)
                                case "\n\t":
                                    if (3.14159 < float64(float32(1e10))) {
                                        if (!(-2147483648 > -2147483648)) {
                                            fmt.Println(v4, v3, v4)
                                            v4 -= 0
                                            v12 := GenS2{F0: 1}
                                            _ = v12
                                        }
                                        var v13 uint = 65535
                                        _ = v13
                                    }
                                default:
                                    if false {
                                        switch "a" {
                                        case "a":
                                            var v14 int16 = 32767
                                            _ = v14
                                        case "\n\t":
                                            v3 += 255
                                        }
                                    } else {
                                        var v15 int8 = -114
                                        v4 = (v4 << ((65535) & 63))
                                        _ = v15
                                    }
                                }
                            case 128:
                                v16 := GenS1{F0: 2, F1: fi11, F2: 32767}
                                _ = v16
                            case -93:
                                fmt.Println(fi11, fi9, v2)
                            }
                        }
                        _ = fi11
                    }
                    var v17 int16 = 144
                    _ = v10
                    _ = v17
                }
                v18 := GenS2{F0: 94}
                _ = fi9
                _ = v18
            }
        }
    default:
        var v19 uint16 = (v3 / (((v3 ^ v3)) | 1))
        _ = v19
    }
    v4 -= (v4 + (v4 - 42))
    v20 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: -32768}
    v3 = v3
    var v21 uint32 = 2
    if true {
        if false {
            v21 -= v21
            var v22 int8 = 42
            _ = v22
        } else {
            delete(v5, (v21 / ((v21) | 1)))
            fmt.Println(v3, v4, v3)
        }
    }
    for fi23 := 0; fi23 < 11; fi23++ {
        switch "hello" {
        case "":
            v24 := GenS1{F0: 255, F1: 32768, F2: int16(v3)}
            _ = v24
        case "\n\t":
            if ((0 > 4294967295) && true) {
                var v25 int64 = v4
                _ = v25
            }
        default:
            if true {
                v3 = v3
                fmt.Println(fi23, v2)
            } else {
                v4 = -9223372036854775808
                _ = len(v5)
                var v26 uint = 255
                _ = v26
            }
        }
        var v27 int8 = -1
        _ = fi23
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
