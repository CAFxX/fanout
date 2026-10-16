package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int16; F1 bool; F2 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint8 = 255
    v3 := []int{-259, -2147483648, -2147483648}
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    if (32767 == 1) {
        v1 = 840
        switch 127 {
        case 1:
            v2 -= 128
        default:
            v2 -= (v2 << ((629) & 63))
        }
    } else {
        v2 = (v2 + v2)
        if (!(-2147483648 != 749)) {
            v6 := GenS1{F0: 42, F1: -9223372036854775808, F2: 42}
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 string = ("\n\t" + "a")
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            v1 += v1
        }
        fmt.Println(v2)
    }
    fmt.Println(v1, v1)
    var v9 bool = true
    if true {
        var v10 int = 80
        if (!v9) {
            v11 := GenS2{F0: (v10 ^ 643), F1: 127}
            v12 := GenS3{F0: 1000, F1: true, F2: (1 <= uint(v2))}
            _ = v11
            _ = v12
        }
        _ = v10
    }
    v13 := GenS1{F0: 826, F1: -819, F2: 255}
    v9 = (-1 <= 1)
    var v14 uint16 = 2
    fmt.Println(v1, v2)
    v15 := GenS2{F0: -783, F1: -12}
    v14 = 255
    v16 := GenS2{F0: 1, F1: 100}
    var v17 int8 = -40
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
