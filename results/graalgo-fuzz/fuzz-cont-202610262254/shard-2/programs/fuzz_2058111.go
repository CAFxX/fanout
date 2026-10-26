package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    var v3 int = 9223372036854775807
    var v4 int16 = -1
    fmt.Println(v1, v1, v4)
    v5 := GenS1{F0: -513, F1: ("\n\t" + ""), F2: 127}
    v1 -= (v1 ^ 199)
    if true {
        v1 -= 4294967295
        var v6 uint32 = 978
        _ = v6
    }
    var v7 float32 = (-2.25 - (0.1 - 0.1))
    v8 := GenS1{F0: -755, F1: ("" + ("\n\t" + "")), F2: 2}
    var v9 bool = (2 >= 0)
    var v10 float32 = (float32(1e10) + 0.1)
    if (0 >= 2) {
        var v11 bool = (true && (32768 >= -1))
        var v12 int16 = (v4 - (v4 << ((225) & 63)))
        if v9 {
            if (77 != 62) {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 float64 = ((3.14159 - 3.14159) + float64(1e10))
                    _ = fi13
                    _ = v14
                }
                var v15 uint32 = 4294967295
                _ = v15
            }
        }
        _ = v11
        _ = v12
    } else {
        var v16 float64 = (1e10 / 1e10)
        v17 := GenS2{F0: (v3 / (((v3 / (((^(v3))) | 1))) | 1))}
        _ = v16
        _ = v17
    }
    v10 += (((0.1 - float32(-2.25)) - v10) * v7)
    v18 := GenS2{F0: (-(v3))}
    var v19 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
