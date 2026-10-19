package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int16 = 1000
    var v2 uint16 = 190
    var v3 int16 = v1
    var v4 float32 = (-2.25 / 1.0)
    var v5 string = "hello"
    var v6 int32 = -1
    v5 = (("foo bar" + v5) + "")
    v5 = "hello"
    v7 := GenS1{F0: (v3 % ((v1) | 1))}
    switch v5 {
    case "\n\t":
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = v5
            switch v5 {
            case "foo bar":
                v1 += v1
            case "":
                if (false || (!true)) {
                    var v9 int8 = 0
                    if (false || (!(!false))) {
                        v4 = 0.0
                        var v10 uint32 = 606
                        if false {
                            if true {
                                v11 := GenS2{F0: 1}
                                _ = v11
                            } else {
                                if false {
                                    for fi12 := 0; fi12 < 2; fi12++ {
                                        v10 += 65535
                                        _ = fi12
                                    }
                                }
                            }
                        } else {
                            var v13 uint64 = 4294967295
                            switch fi8 {
                            case 400:
                                var v14 uint64 = 245
                                _ = v14
                            case 577:
                                for fi15 := 0; fi15 < 16; fi15++ {
                                    v16 := GenS2{F0: 127}
                                    v1 += 1000
                                    fmt.Println(v13, v1, v1)
                                    _ = fi15
                                    _ = v16
                                }
                            case 0:
                                v17 := []int{-474, fi8, fi8, (fi8 << ((1) & 63))}
                                _ = v17
                            default:
                                var v18 float64 = float64(v4)
                                _ = v18
                            }
                            _ = v13
                        }
                        _ = v10
                    }
                    _ = v9
                } else {
                    v19 := []int{0, fi8}
                    var v20 int = 4294967295
                    switch -419 {
                    case 9223372036854775807:
                        _ = len(v19)
                    default:
                        v6 = v6
                    }
                    _ = v19
                    _ = v20
                }
            case "a":
                if (true && (v4 < (v4 / float32(3.14159)))) {
                    v21 := GenS3{F0: -787, F1: 346}
                    var v22 float32 = v4
                    for fi23 := 0; fi23 < 2; fi23++ {
                        v24 := GenS4{F0: "\n\t", F1: 1}
                        v25 := GenS4{F0: "a", F1: 1}
                        if (!true) {
                            v6 += (v6 / ((1) | 1))
                            var v26 float32 = v4
                            v22 = (v4 * v26)
                            var v27 int64 = 460
                            v28 := []int{}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                }
            }
            v4 = float32(((-2.25 - 1.5) - (((-2.25 + (3.14159 + 0.1)) + 3.14159) * 3.14159)))
            _ = fi8
        }
    }
    v2 = (v2 / ((0) | 1))
    for fi29 := 0; fi29 < 2; fi29++ {
        for fi30 := 0; fi30 < 13; fi30++ {
            v1 -= (v1 ^ (v1 ^ -1))
            v31 := GenS5{F0: 128, F1: 42}
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println(v2)
    v32 := []int{-9223372036854775808, 32768}
    v33 := GenS3{F0: 9223372036854775807, F1: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
