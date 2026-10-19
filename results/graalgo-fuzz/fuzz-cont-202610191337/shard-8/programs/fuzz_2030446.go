package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 127
    var v3 int32 = -1
    var v4 bool = false
    var v5 int64 = 1
    v2 -= 17
    var v6 uint32 = 23
    fmt.Println(v1, v2, v3)
    v2 = (v2 + (^(v2)))
    var v7 uint16 = v1
    v3 += (v3 ^ (v3 - (v3 | v3)))
    if v4 {
        if v4 {
            v3 -= (^(v3))
            var v8 float64 = 3.14159
            if v4 {
                v9 := GenS1{F0: 32767, F1: (-118 != int16(v3))}
                if (0.1 == 3.14159) {
                    if (18446744073709551615 <= 4294967295) {
                        var v10 int32 = 2147483647
                        var v11 uint8 = (v2 & v2)
                        v7 -= v1
                        _ = v10
                        _ = v11
                    } else {
                        v12 := []int{127}
                        var v13 uint64 = 42
                        _ = v12
                        _ = v13
                    }
                } else {
                    v8 += 3.14159
                    fmt.Println(v7, v6, v7)
                    var v14 bool = (v8 > 1.5)
                    _ = v14
                }
                _ = v9
            }
            _ = v8
        } else {
            v4 = v4
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{-101, 32768, -9223372036854775808, 127}
            _ = fi15
            _ = v16
        }
    }
    var v17 int16 = 1
    v18 := GenS2{F0: v3, F1: (0.0 == 0.0)}
    switch "foo bar" {
    case "\n\t":
        for fi19 := 0; fi19 < 19; fi19++ {
            v17 = -32768
            _ = fi19
        }
    case "foo bar":
        v20 := []int{}
        _ = v20
    default:
        v1 = v7
    }
    v3 = v3
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 int32 = (v3 - v3)
        if false {
            v23 := GenS1{F0: -32768, F1: true}
            v5 -= (v5 & 9223372036854775807)
            v7 = 88
            _ = v23
        } else {
            switch "\n\t" {
            case "foo bar":
                v3 = 42
            default:
                var v24 uint32 = 65535
                _ = v24
            }
            var v25 float32 = 1.5
            _ = v25
        }
        switch (-(fi21)) {
        case -2147483648:
            v26 := []int{42, fi21}
            _ = v26
        default:
            for fi27 := 0; fi27 < 11; fi27++ {
                fmt.Println(v5, fi21, v17)
                v2 += (v2 / ((1) | 1))
                _ = fi27
            }
        }
        _ = fi21
        _ = v22
    }
    fmt.Println(v4, v2)
    fmt.Println(v5)
    var v28 int8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
