package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 719
    var v3 int64 = 430
    v4 := GenS1{F0: 32767, F1: -32768}
    v5 := GenS1{F0: 1, F1: 238}
    if false {
        var v6 float32 = 1.5
        _ = v6
    }
    switch 88 {
    case 0:
        var v7 uint32 = v1
        _ = v7
    case 32767:
        for fi8 := 0; fi8 < 1; fi8++ {
            if (0 != 716) {
                var v9 int16 = 1
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 int = fi10
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                var v12 int16 = 1
                v3 -= int64(fi8)
                _ = v12
            }
            switch fi8 {
            case -1:
                switch "" {
                case "foo bar":
                    if false {
                        var v13 int = fi8
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 int64 = (v3 & (v3 / (((v3 * (v3 << ((4294967295) & 63)))) | 1)))
                            fmt.Println(v1, fi14)
                            var v16 float64 = 0.0
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 3; fi17++ {
                            fmt.Println(v13, v3)
                            fmt.Println(fi8, fi8)
                            var v18 int16 = 0
                            if true {
                                v19 := GenS2{F0: (false && ((v1 * v1) <= v1)), F1: (v1 / ((v1) | 1)), F2: fi17}
                                v20 := GenS3{F0: 84, F1: (v1 >> ((0) & 63))}
                                _ = v19
                                _ = v20
                            }
                            _ = fi17
                            _ = v18
                        }
                        _ = v13
                    } else {
                        if (int32(fi8) == -886) {
                            v21 := []int{fi8, (fi8 >> ((1) & 63))}
                            v2 = (v2 + (v2 + (v2 / ((v2) | 1))))
                            v2 = 2147483647
                            _ = v21
                        }
                        switch (fi8 - fi8) {
                        case 32768:
                            var v22 bool = false
                            _ = v22
                        case 42:
                            var v23 uint = 255
                            _ = v23
                        default:
                            v24 := GenS3{F0: 1, F1: v1}
                            _ = v24
                        }
                        v2 -= (v2 * -1)
                    }
                case "hello":
                    var v25 int64 = v3
                    _ = v25
                default:
                    v3 += -9223372036854775808
                }
            case -2147483648:
                v26 := []int{fi8}
                _ = v26
            case 127:
                if (255 != 252) {
                    var v27 uint64 = 4294967295
                    var v28 uint64 = (v27 / ((551) | 1))
                    _ = v27
                    _ = v28
                }
            }
            _ = fi8
        }
    default:
        switch -2147483648 {
        case -976:
            v2 -= (v2 - v2)
        }
    }
    v3 += v3
    fmt.Println(v2, v2)
    v2 += (v2 - v2)
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 uint = 18446744073709551615
        _ = fi29
        _ = v30
    }
    if (true && true) {
        v1 = (v1 + (v1 % ((v1) | 1)))
        v2 -= v2
        fmt.Println(v1, v3, v3)
    }
    v31 := GenS4{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
