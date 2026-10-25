package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 int16 }
type GenS3 struct { F0 string; F1 uint32 }

func main() {
    var v1 float64 = (0.0 + 0.0)
    var v2 uint8 = 155
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 4294967295
    var v5 string = "foo bar"
    var v6 int16 = -84
    v6 = -1
    for fi7 := 0; fi7 < 18; fi7++ {
        v5 = "a"
        _ = fi7
    }
    if (-128 >= -21) {
        fmt.Println(v4, v5, v4)
        v8 := GenS1{F0: 2147483648, F1: -166}
        switch "\n\t" {
        case "a":
            v9 := []int{2147483647, 127}
            _ = v9
        default:
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v4 -= (v4 % ((1) | 1))
                    _ = fi11
                }
                if true {
                    var v12 string = v5
                    _ = v12
                }
                _ = fi10
            }
        }
        _ = v8
    }
    if false {
        fmt.Println(v4, v1)
        v13 := GenS1{F0: uint32(v2), F1: -9223372036854775808}
        _ = v13
    } else {
        v14 := GenS2{F0: v5, F1: int64(v2), F2: (v6 % ((v6) | 1))}
        v15 := []int{128, -476}
        _ = v14
        _ = v15
    }
    switch -219 {
    case -1:
        var v16 uint64 = v4
        _ = v16
    }
    var v17 int = 9223372036854775807
    var v18 int16 = (^(v6))
    v4 += (v4 % (((v4 / ((2) | 1))) | 1))
    v19 := GenS1{F0: 42, F1: v17}
    var v20 bool = ((0 > 0) || true)
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int8 = 127
        v23 := GenS2{F0: "", F1: -1, F2: -711}
        v24 := GenS1{F0: 257, F1: v17}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    v6 = (v6 % (((v6 + 1)) | 1))
    fmt.Println(v5)
    v25 := GenS3{F0: ((v5 + v5) + ""), F1: 22}
    switch v17 {
    case 2147483647:
        var v26 uint16 = 23
        _ = v26
    case 127:
        if ((v4 ^ (v4 % (((v4 >> ((4294967295) & 63))) | 1))) >= 541) {
            v27 := GenS1{F0: 65535, F1: (v17 >> ((2) & 63))}
            fmt.Println(v6, v3, v2)
            _ = v27
        }
    default:
        fmt.Println(v2, v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
