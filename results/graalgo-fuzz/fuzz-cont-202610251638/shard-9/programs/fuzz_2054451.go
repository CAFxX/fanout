package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint64 = 624
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 uint16 = 65535
    v6 := v2[161]
    var v7 string = "hello"
    if false {
        v6 = v6
        switch -9223372036854775808 {
        case -990:
            v8 := GenS1{F0: 1, F1: -1, F2: uint8(v5)}
            _ = v8
        case 1:
            var v9 int = 0
            _ = v9
        }
    } else {
        var v10 uint8 = 128
        v11 := GenS1{F0: (v10 & v10), F1: -250, F2: (^(v10))}
        _ = v10
        _ = v11
    }
    v12 := []int{338, 2147483647}
    v6 = (2147483647 != 0)
    fmt.Println(v5)
    fmt.Println(v1)
    v13 := GenS1{F0: 1, F1: 42, F2: uint8(v5)}
    var v14 uint64 = (v1 / (((^(v1))) | 1))
    fmt.Println(v1, v5, v7)
    if ((!(42 == 42)) || (-289 > 484)) {
        if (18446744073709551615 > 255) {
            for fi15 := 0; fi15 < 8; fi15++ {
                if (0.0 == 1e10) {
                    fmt.Println(v7, v6)
                }
                v16 := GenS1{F0: 1, F1: -1, F2: 255}
                _ = fi15
                _ = v16
            }
        } else {
            v14 = (v14 ^ 153)
        }
        v5 -= (v5 % (((v5 >> ((255) & 63))) | 1))
        fmt.Println(v7, v7)
    } else {
        v17 := GenS1{F0: 42, F1: -873, F2: 104}
        fmt.Println(v7, v14, v7)
        switch -1 {
        case 42:
            v4[-1] = v6
        case 1:
            if ((!(1 == 0)) && v6) {
                v18 := GenS2{F0: (v1 ^ (v14 - v1)), F1: -2147483648}
                _ = v18
            }
        case -2147483648:
            for fi19 := 0; fi19 < 7; fi19++ {
                switch v7 {
                case "":
                    if v6 {
                        if (-14 != 2147483647) {
                            if v6 {
                                if v6 {
                                    v7 = "a"
                                    for fi20 := 0; fi20 < 5; fi20++ {
                                        if len(v12) > 0 {
                                            v12[0] = fi19
                                        }
                                        fmt.Println(v7, fi19)
                                        _ = fi20
                                    }
                                    delete(v3, (v7 + ((v7 + "") + "hello")))
                                }
                                var v21 bool = true
                                _ = v21
                            }
                        }
                        var v22 float32 = float32(0.0)
                        _ = v22
                    } else {
                        v7 = v7
                    }
                case "a":
                    var v23 int8 = -70
                    _ = v23
                default:
                    v14 = v14
                }
                v24 := GenS2{F0: 18446744073709551615, F1: -2147483648}
                _ = fi19
                _ = v24
            }
        }
        _ = v17
    }
    v25 := GenS2{F0: 4294967295, F1: 1}
    v1 += v14
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
