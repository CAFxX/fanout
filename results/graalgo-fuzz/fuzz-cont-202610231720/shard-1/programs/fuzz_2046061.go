package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 42
    v3 := []string{("hello" + (("hello" + "a") + "a"))}
    var v4 uint32 = 185
    v2 += v2
    var v5 int64 = -581
    var v6 float64 = (((float64(((1.5 + ((1e10 / (1e10 / (0.1 + (1e10 * (1e10 / float32(0.1)))))) - 0.0)) + 0.1)) - 3.14159) + -2.25) / 3.14159)
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint16 = 2
            v10 := GenS1{F0: (fi8 % (((fi7 * (fi8 << ((65535) & 63)))) | 1)), F1: 629, F2: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: fi7, F1: 32767, F2: 53}
        var v12 int32 = -465
        _ = fi7
        _ = v11
        _ = v12
    }
    v1[2147483647] = "\n\t"
    v13 := v1[9223372036854775807]
    var v14 uint = 0
    v5 = v5
    for fi15 := 0; fi15 < 2; fi15++ {
        _ = len(v3)
        _ = fi15
    }
    v16 := GenS1{F0: 128, F1: 1, F2: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
