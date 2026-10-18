package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 65535
    var v3 int16 = 586
    var v4 int = -576
    v5 := []int{v4, (v4 + v4), (v4 / ((531) | 1)), (v4 >> ((255) & 63))}
    var v6 uint = 255
    v2 += v2
    if true {
        if true {
            if false {
                var v7 uint32 = 492
                _ = v7
            } else {
                for fi8 := 0; fi8 < 1; fi8++ {
                    if (92 == 251) {
                        v9 := GenS1{F0: false, F1: 54, F2: 18446744073709551615}
                        var v10 uint16 = 367
                        v11 := GenS2{F0: v10}
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                    _ = len(v1)
                    _ = fi8
                }
            }
            v6 += v6
            if (1.5 <= ((3.14159 + 3.14159) * (float32((0.0 + -2.25)) - (((0.0 - (-2.25 * 1.5)) + (3.14159 - 1.5)) * float32(((0.0 / -2.25) - 1e10)))))) {
                delete(v1, 2147483648)
                var v12 int32 = 0
                _ = v12
            } else {
                fmt.Println(v4, v3, v3)
            }
        }
        v5 = append(v5, 4294967295)
    } else {
        v13 := GenS1{F0: (false && true), F1: 32, F2: 324}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v13
    }
    if false {
        v3 -= (v3 + -32768)
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 float64 = -2.25
                v17 := v1[2147483647]
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        fmt.Println(v4)
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                if false {
                    delete(v1, 2)
                    var v20 int64 = int64(fi19)
                    fmt.Println(v4, v6, v3)
                    var v21 uint8 = 255
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS2{F0: 32767}
                    fmt.Println(v2, fi18)
                    var v23 uint8 = 127
                    _ = v22
                    _ = v23
                }
                var v24 int64 = -1
                _ = fi19
                _ = v24
            }
            if true {
                var v25 float64 = (float64(0.0) / 1.0)
                var v26 float32 = (-2.25 + float32((v25 * ((float64(float32((v25 + (0.0 - (v25 * 1e10))))) / -2.25) * -2.25))))
                _ = v25
                _ = v26
            } else {
                for fi27 := 0; fi27 < 11; fi27++ {
                    switch 1 {
                    case -9223372036854775808:
                        v6 = v6
                        var v28 uint = (^(v6))
                        _ = v28
                    case 153:
                        v6 = v6
                    case 128:
                        var v29 float64 = (((0.0 - 0.0) * 0.1) * 0.0)
                        _ = v29
                    }
                    fmt.Println(v2, fi27, v3)
                    _ = fi27
                }
            }
            _ = fi18
        }
        var v30 uint = v6
        var v31 uint = (v30 - v30)
        _ = v30
        _ = v31
    }
    _ = len(v5)
    var v32 int16 = (v3 * v3)
    v4 = (v4 / ((v4) | 1))
    v6 = 0
    switch v4 {
    case 0:
        var v33 float64 = float64(0.0)
        _ = v33
    case 327:
        delete(v1, 42)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v1[2147483648] = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
