package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint = 65535
    var v3 uint8 = 128
    var v4 int32 = 1
    v5 := []string{"\n\t", "a", "", "foo bar"}
    var v6 bool = (100 == -9223372036854775808)
    if v6 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi7 := 0; fi7 < 16; fi7++ {
            if v6 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v8 float32 = float32((0.1 / (((-2.25 * (1e10 * 3.14159)) + float64(((1e10 * 3.14159) / (1e10 + 3.14159)))) - 1e10)))
                v4 = v4
                _ = v8
            }
            v3 -= uint8(v1)
            _ = fi7
        }
    }
    var v9 float32 = 1.5
    if v6 {
        fmt.Println(v1)
        _ = len(v5)
    } else {
        v1 += (^(v1))
        v10 := GenS1{F0: v3, F1: 9223372036854775807}
        _ = v10
    }
    var v11 int8 = 1
    switch "foo bar" {
    case "\n\t":
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: v3, F1: (fi12 % (((-(fi12))) | 1))}
            for fi14 := 0; fi14 < 16; fi14++ {
                fmt.Println(v4, fi12, v11)
                v15 := GenS1{F0: (v3 - 127), F1: (fi12 % ((42) | 1))}
                _ = len(v5)
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
    case "a":
        v16 := GenS2{F0: 669}
        _ = v16
    }
    v9 = (((float32(3.14159) + (v9 - float32((0.1 + -2.25)))) - ((((((1.5 - float32(1.5)) / 1.0) + ((((v9 + v9) + (v9 + v9)) - (v9 + (((1.5 * v9) * float32(0.0)) + v9))) * -2.25)) - (v9 * 3.14159)) + float32(1.5)) + (v9 - ((0.1 / v9) * 1e10)))) + v9)
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
