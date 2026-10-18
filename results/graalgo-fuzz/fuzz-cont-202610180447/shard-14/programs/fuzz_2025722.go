package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 bool = (true || true)
    var v2 float32 = 3.14159
    v3 := []string{("hello" + "\n\t"), "hello", "\n\t"}
    var v4 string = ("" + ("" + ""))
    var v5 float32 = float32(((float64((v2 + ((float32(float64(v2)) + (v2 / ((1.5 / (float32((-2.25 / 1.5)) - (((3.14159 / v2) * ((float32(((1e10 / ((1e10 * 0.1) + float64(0.0))) / 1e10)) + v2) + v2)) + ((v2 * 0.0) + v2)))) / ((v2 - (1e10 + v2)) / (float32(float64(v2)) / (v2 / ((v2 / v2) * (((v2 * ((0.1 * (v2 - (1e10 - 3.14159))) / (v2 * ((1.5 + -2.25) - v2)))) - (v2 + float32(-2.25))) - v2)))))))) + v2))) * 3.14159) - float64(0.0)))
    var v6 int8 = -128
    var v7 uint16 = 0
    var v8 float64 = 3.14159
    for fi9 := 0; fi9 < 18; fi9++ {
        _ = len(v3)
        for fi10 := 0; fi10 < 11; fi10++ {
            v2 -= v5
            _ = fi10
        }
        v2 = 1.5
        _ = fi9
    }
    if (65535 < 65535) {
        v11 := GenS1{F0: 54}
        _ = v11
    } else {
        if (511 <= 18446744073709551615) {
            if len(v3) > 0 {
                v3[0] = v4
            }
            var v12 int16 = 1
            v2 -= v2
            _ = v12
        }
    }
    var v13 string = "foo bar"
    v7 -= v7
    v2 -= (float32(v8) / 3.14159)
    v7 = v7
    v2 = 1.5
    v8 += v8
    var v14 uint16 = 908
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
