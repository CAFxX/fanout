package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float32 = (float32(float64((1.5 + 1.5))) * float32(float64(3.14159)))
    var v2 int8 = 1
    v3 := []int{-586}
    v4 := make(map[int]string)
    var v5 uint8 = 127
    fmt.Println(v1, v5, v5)
    for fi6 := 0; fi6 < 5; fi6++ {
        v5 = (v5 & v5)
        _ = fi6
    }
    v7 := GenS1{F0: 42, F1: ("foo bar" + "a")}
    v4[9223372036854775807] = "a"
    var v8 float64 = ((float64(v1) + 1.5) / 0.1)
    if true {
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 float32 = (float32((v8 + v8)) * v1)
            v8 -= v8
            var v11 int8 = (v2 / ((int8(fi9)) | 1))
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int32 = 2147483647
    v13 := GenS2{F0: v2}
    switch -333 {
    case -2147483648:
        v12 = (v12 << ((18446744073709551615) & 63))
    }
    v8 = -2.25
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 float64 = ((float64(((float32(float64((float32(v8) * (-2.25 * v1)))) + 3.14159) - 1e10)) / 1.5) * -2.25)
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
