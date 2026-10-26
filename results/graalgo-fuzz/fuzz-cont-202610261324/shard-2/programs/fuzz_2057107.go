package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 bool }

func main() {
    var v1 int16 = 42
    v2 := make(map[string]string)
    var v3 uint = 65535
    var v4 float64 = 3.14159
    var v5 int8 = 92
    v4 -= 3.14159
    v6 := GenS1{F0: int16(v3), F1: -270, F2: 52}
    v7 := GenS1{F0: (v1 / ((v1) | 1)), F1: 0, F2: 453}
    fmt.Println(v3, v5)
    switch ("" + "hello") {
    case "\n\t":
        var v8 int32 = 0
        _ = v8
    default:
        v3 = v3
    }
    switch ("hello" + ("foo bar" + "hello")) {
    case "":
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                v1 += v1
                if false {
                    v11 := v2["hello"]
                    v12 := []int{32767, fi9, 1, fi9}
                    v3 = v3
                    _ = v11
                    _ = v12
                } else {
                    v3 = v3
                    if true {
                        var v13 int64 = 638
                        _ = v13
                    }
                    v14 := []int{fi10, fi9, fi9, fi10}
                    if false {
                        var v15 uint = (^(v3))
                        var v16 int32 = 42
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: v1, F1: (fi9 << (((v3 - (v3 % ((v3) | 1)))) & 63)), F2: 42}
                    _ = v14
                    _ = v17
                }
                v18 := GenS2{F0: -9223372036854775808, F1: 32767, F2: true}
                _ = fi10
                _ = v18
            }
            _ = fi9
        }
    default:
        v19 := []int{454, 32768, 127, 805}
        _ = v19
    }
    var v20 bool = ((!(1.5 != (0.1 / float32(float64((1e10 + (float32((0.0 - v4)) + 3.14159))))))) || true)
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 float64 = 0.1
        var v23 int = fi21
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: -472, F1: -16, F2: (-776 != -9223372036854775808)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
