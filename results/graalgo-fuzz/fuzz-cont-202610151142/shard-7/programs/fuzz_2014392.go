package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint }

func main() {
    var v1 int16 = 1
    var v2 int64 = -565
    var v3 float64 = 1.5
    v4 := make(map[int]bool)
    v5 := GenS1{F0: -1}
    var v6 bool = true
    if (2147483647 <= -2147483648) {
        var v7 uint = 255
        _ = v7
    } else {
        if v6 {
            switch -9223372036854775808 {
            case 320:
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := GenS1{F0: 0}
                    v10 := GenS1{F0: -87}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
            v6 = (59 != 0)
        } else {
            if true {
                v4[32768] = ((3.14159 <= 1.5) && (!v6))
            }
            v11 := []int{-2147483648}
            _ = v11
        }
    }
    var v12 int = -411
    var v13 uint = 42
    for fi14 := 0; fi14 < 14; fi14++ {
        v13 = 50
        _ = fi14
    }
    var v15 int64 = (^(v2))
    switch v12 {
    case 0:
        if (42 != 2147483647) {
            fmt.Println(v3)
            switch -500 {
            case -1:
                var v16 string = "a"
                _ = v16
            case 127:
                v3 = v3
            default:
                v13 += v13
            }
            var v17 int64 = v2
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 16; fi18++ {
            v19 := GenS2{F0: 42, F1: 97, F2: 734}
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 18; fi21++ {
            if ((3.14159 + ((0.0 + 1.5) * 1.5)) < (1e10 * float32(((3.14159 * v3) * (0.0 + (((3.14159 - v3) - float64((1.5 + (3.14159 + (-2.25 / (float32(((((v3 + v3) * v3) * ((((v3 * -2.25) / v3) / (float64(1e10) * ((float64(3.14159) * v3) * (1.5 - 0.1)))) + 1.5)) + (1.5 * (v3 * v3)))) + 0.1)))))) - v3)))))) {
                if true {
                    v22 := GenS3{F0: 42}
                    switch ("\n\t" + "") {
                    case "":
                        v2 += -9223372036854775808
                    case "hello":
                        fmt.Println(v6)
                    default:
                        switch (fi21 + (fi21 ^ (fi21 & 42))) {
                        case 9223372036854775807:
                            var v23 uint32 = 0
                            _ = v23
                        case -194:
                            var v24 int64 = -1
                            _ = v24
                        case 127:
                            v4[(-(fi20))] = (0 != -2147483648)
                        default:
                            var v25 int = fi21
                            _ = v25
                        }
                        v12 += 128
                        var v26 int32 = 1
                        v13 += (v13 % (((v13 - (v13 - (^(v13))))) | 1))
                        _ = v26
                    }
                    switch 16 {
                    case -9223372036854775808:
                        fmt.Println(v12, v3)
                    }
                    _ = v22
                } else {
                    v15 = v15
                    v27 := []int{32768, (^(fi20)), fi21}
                    var v28 int8 = 1
                    _ = v27
                    _ = v28
                }
                v3 += v3
                var v29 int64 = int64(v13)
                _ = v29
            } else {
                var v30 uint8 = 241
                var v31 uint8 = v30
                _ = v30
                _ = v31
            }
            _ = fi21
        }
        v32 := GenS4{F0: 0, F1: 18446744073709551615}
        _ = fi20
        _ = v32
    }
    var v33 uint16 = 191
    v34 := []int{128, v12, 9223372036854775807}
    v35 := v4[v12]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
