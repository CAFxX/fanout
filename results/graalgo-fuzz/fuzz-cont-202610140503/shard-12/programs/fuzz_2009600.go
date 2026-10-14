package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 string; F2 int64 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 218
    var v3 uint64 = 928
    var v4 float32 = 0.0
    var v5 int32 = 2147483647
    var v6 float32 = v4
    v5 += -1
    v7 := []int{9223372036854775807, -2147483648, 1}
    v5 += v5
    var v8 int64 = -9223372036854775808
    if false {
        if true {
            var v9 int = 127
            if true {
                if (!true) {
                    if (v5 >= v5) {
                        var v10 uint32 = 0
                        v3 = (v3 | v3)
                        v7 = append(v7, v9)
                        _ = v10
                    }
                } else {
                    v3 = 82
                    if false {
                        v3 -= 2
                    }
                }
            } else {
                if true {
                    v9 = (v9 | (v9 - -21))
                    _ = len(v1)
                } else {
                    v2 -= (^(v2))
                    v4 -= float32(float64(1e10))
                }
                v6 += (v4 * (1e10 / (v6 / v4)))
            }
            for fi11 := 0; fi11 < 14; fi11++ {
                fmt.Println(v8, v8)
                var v12 uint8 = 109
                var v13 uint32 = 722
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v9
        }
    } else {
        switch 9223372036854775807 {
        case 32767:
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 string = "foo bar"
                var v16 int8 = 0
                switch (fi14 + (-(fi14))) {
                case -558:
                    for fi17 := 0; fi17 < 2; fi17++ {
                        if len(v7) > 0 {
                            v7[0] = int(v2)
                        }
                        switch (fi17 % ((fi14) | 1)) {
                        case 9223372036854775807:
                            switch (v15 + v15) {
                            case "foo bar":
                                if (false || (42 <= 494)) {
                                    v18 := GenS1{F0: (v5 * v5), F1: (77 != 18446744073709551615)}
                                    var v19 int64 = v8
                                    v20 := GenS1{F0: -385, F1: (1 >= 2)}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    v2 += (^(v2))
                                    v3 = v3
                                    v1[fi17] = (("\n\t" + (v15 + v15)) + v15)
                                }
                            case "\n\t":
                                var v21 float64 = 1e10
                                _ = v21
                            case "hello":
                                var v22 uint = 2
                                _ = v22
                            default:
                                var v23 uint64 = v3
                                _ = v23
                            }
                        case -701:
                            switch v15 {
                            case "foo bar":
                                v4 = float32(-2.25)
                            default:
                                v8 = v8
                            }
                        case 128:
                            v6 -= ((v6 - v4) + v6)
                        default:
                            fmt.Println(fi14, v5)
                        }
                        fmt.Println(v4, v3, fi17)
                        _ = fi17
                    }
                case 837:
                    var v24 float32 = v6
                    _ = v24
                case -1:
                    var v25 bool = ((v8 ^ v8) >= v8)
                    _ = v25
                }
                _ = fi14
                _ = v15
                _ = v16
            }
        case 43:
            v26 := GenS2{F0: uint16(v3)}
            _ = v26
        case -327:
            v3 = v3
        default:
            delete(v1, int(v8))
        }
        v7 = append(v7, 0)
        v5 = (^(v5))
    }
    fmt.Println(v6, v8, v8)
    if true {
        _ = len(v1)
    }
    v27 := GenS3{F0: "a", F1: "\n\t", F2: v8}
    v28 := GenS4{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
