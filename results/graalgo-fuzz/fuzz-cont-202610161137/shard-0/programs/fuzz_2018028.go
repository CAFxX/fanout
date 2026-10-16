package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }

func main() {
    var v1 string = "a"
    var v2 uint8 = 127
    var v3 string = ""
    var v4 uint = 65535
    var v5 uint64 = 18446744073709551615
    v6 := GenS1{F0: 35, F1: true, F2: ("" + v3)}
    fmt.Println(v1)
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float64 = 1.5
        var v9 int = fi7
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 42
    fmt.Println(v5, v4)
    fmt.Println(v2, v2, v2)
    v11 := GenS1{F0: v2, F1: (65535 <= 42), F2: (("hello" + (v3 + v1)) + v3)}
    var v12 float32 = ((((0.1 / 3.14159) / float32(3.14159)) - (0.1 / 1.0)) / (3.14159 * (3.14159 / ((0.0 - 3.14159) - (1.5 * (((-2.25 / 1e10) / (1.5 - 0.0)) / (3.14159 - ((0.0 + 1e10) * (1.5 / 1.5)))))))))
    fmt.Println(v10, v2)
    fmt.Println(v12, v4, v10)
    for fi13 := 0; fi13 < 20; fi13++ {
        v10 -= (v10 >> (((v4 + (v4 >> (((^(v4))) & 63)))) & 63))
        _ = fi13
    }
    if (-763 <= 1) {
        v1 = v3
        v5 += (v5 + (v5 - (v5 % ((18446744073709551615) | 1))))
        var v14 int32 = -174
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 float32 = 1e10
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
