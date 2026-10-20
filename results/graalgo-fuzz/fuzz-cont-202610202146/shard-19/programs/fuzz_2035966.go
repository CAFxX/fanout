package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    var v3 int64 = -1
    var v4 int16 = -1
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := GenS1{F0: 2147483647}
        v1 += v1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 491}
    _ = len(v2)
    v1 -= (v1 + v1)
    v3 = v3
    v8 := v2[4294967295]
    var v9 uint16 = 731
    var v10 uint64 = 4294967295
    switch v8 {
    case "hello":
        v11 := GenS1{F0: 2147483648}
        _ = v11
    default:
        v12 := []int{2147483647}
        _ = v12
    }
    if (65535 >= 2147483648) {
        v1 += v1
        v13 := v2[42]
        switch "\n\t" {
        case "foo bar":
            var v14 int = -182
            _ = v14
        default:
            fmt.Println(v9, v9, v13)
        }
        _ = v13
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint16 = 1
            v10 = (v10 | (v10 | (^(v10))))
            v17 := GenS2{F0: 751, F1: v10}
            _ = fi15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 10; fi18++ {
            if (float64(((float32((0.1 * ((float64(0.0) - ((0.0 / 1.5) - 0.1)) + -2.25))) - 0.0) + float32(3.14159))) <= 3.14159) {
                var v19 uint = 18446744073709551615
                v3 -= v3
                _ = v19
            }
            _ = fi18
        }
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
