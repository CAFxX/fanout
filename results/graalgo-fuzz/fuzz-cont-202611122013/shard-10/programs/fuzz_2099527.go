package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    var v1 bool = (((526 == -381) && (!((42 == 536) && true))) && false)
    var v2 int16 = -162
    v3 := make(map[string]int)
    var v4 int64 = 42
    var v5 int = -1
    var v6 uint16 = 65535
    var v7 int64 = (v4 / (((v4 << ((0) & 63))) | 1))
    var v8 float64 = float64(0.1)
    if (v2 != v2) {
        switch v5 {
        case -174:
            v9 := v3["a"]
            _ = v9
        }
        v10 := GenS1{F0: (v6 * (v6 % (((^(v6))) | 1))), F1: 1, F2: "a"}
        _ = v10
    } else {
        v11 := GenS2{F0: 255, F1: 127}
        for fi12 := 0; fi12 < 1; fi12++ {
            if (!(1 > v2)) {
                var v13 string = (("\n\t" + "\n\t") + "\n\t")
                v14 := GenS1{F0: v6, F1: v4, F2: "foo bar"}
                switch (v5 * fi12) {
                case -1:
                    for fi15 := 0; fi15 < 8; fi15++ {
                        fmt.Println(v2, v13, v5)
                        _ = fi15
                    }
                case 42:
                    switch 9223372036854775807 {
                    case 9223372036854775807:
                        var v16 float64 = 3.14159
                        _ = v16
                    case 128:
                        switch v13 {
                        case "hello":
                            if (v1 || (int64(v6) >= (v4 + (v4 - v7)))) {
                                if (1 > 98) {
                                    v1 = false
                                } else {
                                    v17 := []int{-105, fi12}
                                    var v18 float64 = (1.5 / ((v8 * ((((3.14159 - v8) / float64(float32((1.5 + 3.14159)))) * 1e10) + v8)) - (v8 * v8)))
                                    v17 = append(v17, fi12)
                                    _ = v17
                                    _ = v18
                                }
                            } else {
                                if false {
                                    v2 += 824
                                }
                                v19 := []int{(^(v5)), (^(v5)), (-(fi12))}
                                _ = v19
                            }
                        case "":
                            var v20 int16 = v2
                            _ = v20
                        case "\n\t":
                            v21 := GenS3{F0: 42, F1: (v7 / ((826) | 1)), F2: 0}
                            _ = v21
                        }
                    case 945:
                        if ((float32(((0.1 + (((v8 / v8) - 0.0) / v8)) - (((float64(1.5) * (1e10 - 0.1)) - -2.25) - float64((3.14159 + (float32(0.1) + ((0.1 / 1.0) / 1.0))))))) - (1e10 + 0.1)) > 0.0) {
                            _ = len(v3)
                            v7 += v4
                            v22 := GenS3{F0: 853, F1: (v4 * v4), F2: -36}
                            _ = v22
                        }
                    default:
                        var v23 int32 = -784
                        _ = v23
                    }
                case -138:
                    v1 = true
                default:
                    v13 = v13
                }
                _ = v13
                _ = v14
            } else {
                v24 := []int{fi12, (fi12 / ((-492) | 1))}
                _ = v24
            }
            _ = fi12
        }
        _ = v11
    }
    v8 = v8
    var v25 int = v5
    switch "hello" {
    case "":
        v7 -= (v4 << ((65535) & 63))
    case "foo bar":
        if ((995 == 2) && v1) {
            var v26 uint16 = 42
            _ = v26
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                _ = len(v3)
                _ = fi27
            }
        }
    }
    if (-2.25 != 0.0) {
        var v28 uint64 = 2
        var v29 bool = false
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
