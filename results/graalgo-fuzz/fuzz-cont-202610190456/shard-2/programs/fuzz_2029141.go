package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int8; F1 uint8; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (((-2.25 - ((float64(0.0) - (0.1 / 1.0)) + (3.14159 * ((-2.25 - -2.25) * (1e10 + 1.5))))) + ((((1e10 * 0.1) / float64(-2.25)) + ((3.14159 * 1e10) - 1.5)) / 0.1)) >= ((0.0 - 1e10) + (float64(0.0) * 3.14159)))
    var v3 int = 4294967295
    v4 := make(map[uint32]bool)
    var v5 int8 = 127
    var v6 float64 = 1.5
    fmt.Println(v3, v3, v5)
    v3 = (^(v3))
    var v7 float32 = 1e10
    fmt.Println(v2)
    v2 = (1 >= 344)
    var v8 float64 = (v6 + 0.1)
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int8 = (^(v5))
        fmt.Println(v3)
        v11 := GenS1{F0: (v2 || v2), F1: 0}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v4)
    fmt.Println(v7, v2, v5)
    if false {
        if (true || true) {
            v3 += (v3 * 32767)
            fmt.Println(v5)
            switch "\n\t" {
            case "hello":
                v5 = v5
            case "foo bar":
                v12 := GenS1{F0: ((255 == 504) || v2), F1: 2147483647}
                _ = v12
            }
        } else {
            v13 := []int{(v3 << ((993) & 63)), v3, (v3 ^ v3)}
            if true {
                fmt.Println(v5, v7)
            } else {
                v2 = ((316 < 65535) && (2147483647 >= 42))
                for fi14 := 0; fi14 < 20; fi14++ {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v15 int32 = 42
                    _ = fi14
                    _ = v15
                }
            }
            _ = v13
        }
        v16 := []int{(v3 % ((-475) | 1))}
        _ = v16
    } else {
        v7 = (v7 + (v7 * (float32(3.14159) + v7)))
    }
    v8 -= (((v6 * (float64((((1.5 / 1.0) + (v7 / v7)) / (v7 * (((float32(v6) * -2.25) / float32(v8)) + (v7 / v7))))) - v6)) / 1e10) + 0.0)
    fmt.Println(v3)
    v17 := GenS2{F0: v5, F1: 2, F2: ("" + ("foo bar" + "foo bar"))}
    var v18 uint64 = 2
    var v19 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
