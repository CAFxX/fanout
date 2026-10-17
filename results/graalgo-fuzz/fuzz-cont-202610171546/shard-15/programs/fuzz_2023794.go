package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 int64 = -1
    var v2 float32 = 1.5
    var v3 int32 = 42
    if ((255 == 42) && (((127 == 771) && false) && true)) {
        if (!true) {
            v2 += (float32(3.14159) + 1.5)
            v2 += v2
            fmt.Println(v3)
        }
        if (1 != 42) {
            v2 += (0.0 - v2)
            if false {
                fmt.Println(v3, v3, v1)
                if ((6 <= -1) && false) {
                    v4 := []int{-9223372036854775808, 42}
                    switch 0 {
                    case -1:
                        v3 += (v3 ^ -2147483648)
                    case 128:
                        if len(v4) > 0 {
                            v4[0] = -409
                        }
                    default:
                        v1 -= v1
                    }
                    var v5 uint8 = 169
                    _ = v4
                    _ = v5
                } else {
                    var v6 int64 = (^(v1))
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 uint = 359
                        var v9 float32 = 1e10
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    fmt.Println(v6)
                    _ = v6
                }
            }
            v10 := GenS1{F0: (("hello" == "a") || (false && (true && (!((false || false) || (-2147483648 < -829))))))}
            _ = v10
        }
        fmt.Println(v1)
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            if ((128 != 1) && true) {
                var v12 uint = 52
                fmt.Println(fi11)
                v2 = float32(((float64(((3.14159 - v2) + (0.1 + float32(-2.25)))) - (1e10 - (3.14159 * -2.25))) / (float64(1e10) - (0.0 + (1.5 * 1.5)))))
                _ = v12
            }
            var v13 int = 128
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(fi11, v2)
                _ = fi14
            }
            _ = fi11
            _ = v13
        }
        for fi15 := 0; fi15 < 12; fi15++ {
            var v16 uint32 = 563
            v1 = -9223372036854775808
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: 4294967295, F1: 18446744073709551615}
        _ = v17
    }
    fmt.Println(v2)
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint32 = 1
        v20 := GenS2{F0: 0, F1: 1}
        var v21 int8 = 100
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    v22 := []int{0, int(v1), -1}
    fmt.Println(v2)
    fmt.Println(v2, v1)
    v3 += (^(v3))
    if len(v22) > 0 {
        v22[0] = -910
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
}
