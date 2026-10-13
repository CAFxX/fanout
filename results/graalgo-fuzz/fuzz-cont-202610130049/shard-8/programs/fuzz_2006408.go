package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    v1 := []uint32{}
    var v2 uint64 = 4294967295
    var v3 int16 = 0
    v2 = (v2 << ((65535) & 63))
    v2 = v2
    v4 := GenS1{F0: 2, F1: -1}
    fmt.Println(v3, v2)
    var v5 uint8 = 127
    v1 = append(v1, 2147483647)
    v6 := GenS1{F0: 128, F1: int32(v3)}
    var v7 float64 = -2.25
    var v8 int = -9223372036854775808
    var v9 int16 = -32768
    if len(v1) > 0 {
        v1[0] = 2147483648
    }
    var v10 int = v8
    v1 = append(v1, 765)
    v2 += v2
    switch "" {
    case "hello":
        switch -967 {
        case 0:
            v10 = -2147483648
        case 128:
            var v11 int16 = (v3 >> ((65535) & 63))
            _ = v11
        }
    case "":
        for fi12 := 0; fi12 < 9; fi12++ {
            if true {
                v7 -= (((v7 * v7) * ((-2.25 + v7) + (0.1 * 1e10))) - float64((((0.1 - float32((v7 - (((0.1 / v7) + 1.5) * v7)))) / ((0.1 / 1.0) * ((1e10 / 1.5) / ((1.5 + 0.0) / ((float32(3.14159) / (3.14159 - float32(v7))) + 0.0))))) * ((((1e10 - 0.0) + 0.0) + 0.0) - (float32((float64(0.0) / float64(3.14159))) + (-2.25 * float32(float64(float32((((v7 + 1e10) + v7) - v7))))))))))
            }
            for fi13 := 0; fi13 < 9; fi13++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v14 string = "a"
                fmt.Println(v2)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 7; fi15++ {
                v3 += v3
                v3 += v3
                var v16 float64 = float64((1.5 * 3.14159))
                _ = fi15
                _ = v16
            }
            _ = fi12
        }
    case "\n\t":
        var v17 uint64 = 42
        _ = v17
    }
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
