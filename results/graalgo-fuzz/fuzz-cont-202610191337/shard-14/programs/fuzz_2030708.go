package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint }

func main() {
    var v1 int = -445
    var v2 float32 = (float32(0.0) - 0.1)
    var v3 int = (v1 * (v1 << ((2) & 63)))
    var v4 int16 = 0
    var v5 float64 = 1.5
    var v6 float64 = (0.0 - float64(((v2 / v2) / v2)))
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 = 3.14159
        switch (v3 | (fi7 % (((v1 & fi7)) | 1))) {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := []int{v1, fi8, v3}
                fmt.Println(v1)
                _ = fi8
                _ = v9
            }
        }
        _ = fi7
    }
    v10 := GenS1{F0: 42}
    var v11 uint32 = 2147483648
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int32 = 377
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: "hello", F1: 463}
    fmt.Println(v3, v2)
    v15 := GenS1{F0: -9223372036854775808}
    v16 := GenS2{F0: "", F1: 18446744073709551615}
    fmt.Println(v3, v1)
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 14; fi18++ {
            switch (fi18 % (((v3 << ((811) & 63))) | 1)) {
            case -9223372036854775808:
                v19 := []int{-9223372036854775808, 0, (v1 + 32767), v3}
                _ = v19
            default:
                v6 = (v5 * 1e10)
            }
            _ = fi18
        }
        v20 := []int{v1, 9223372036854775807, 409, (v3 - (v1 ^ fi17))}
        _ = fi17
        _ = v20
    }
    fmt.Println(v11)
    var v21 uint16 = 32767
    var v22 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
