package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    v1 := []string{("\n\t" + (("" + "hello") + "foo bar")), "", "", "\n\t"}
    var v2 int = 42
    var v3 int32 = -1
    switch (v2 | (v2 + v2)) {
    case 496:
        var v4 float32 = ((-2.25 + ((float32((float64(1e10) / -2.25)) * float32((1e10 * (float64((0.1 + (-2.25 - 0.0))) + float64((0.0 + float32((-2.25 / float64((0.1 - 1e10)))))))))) + 0.0)) / 3.14159)
        _ = v4
    case 2147483647:
        var v5 uint32 = 0
        _ = v5
    case -1:
        var v6 int8 = -28
        _ = v6
    }
    v2 = v2
    var v7 uint16 = 489
    v8 := GenS1{F0: -127, F1: (^(v7))}
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 = append(v1, "hello")
        var v10 float64 = (1.5 / float64(-2.25))
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi11
    }
    v7 = (v7 ^ v7)
    var v12 float64 = (0.1 * 1e10)
    var v13 float32 = ((((1.5 * (float32((((v12 - (float64(3.14159) + v12)) * v12) * v12)) - (-2.25 - 1e10))) / (0.1 * ((0.0 / (((((1e10 + (((1e10 / ((((0.0 - float32(3.14159)) * -2.25) * ((1e10 + 0.0) / 0.1)) - ((-2.25 * 3.14159) - 1e10))) + -2.25) / (1.5 + 3.14159))) * 0.1) / (float32(v12) + -2.25)) - float32((-2.25 * -2.25))) / 1.0)) / 1.5))) + -2.25) - 3.14159)
    var v14 float32 = 0.1
    if false {
        var v15 float32 = -2.25
        v3 = (v3 % (((v3 >> ((18446744073709551615) & 63))) | 1))
        v12 = v12
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
