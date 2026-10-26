package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int64 = 42
    var v2 uint32 = 2147483647
    v3 := []int{2147483647, 127, 896}
    var v4 int = 32767
    switch (v4 / ((v4) | 1)) {
    case 9223372036854775807:
        var v5 uint8 = 135
        _ = v5
    default:
        var v6 uint16 = 1
        _ = v6
    }
    fmt.Println(v2, v1)
    if true {
        if len(v3) > 0 {
            v3[0] = v4
        }
        fmt.Println(v1, v4, v1)
        if ("\n\t" >= ("a" + "\n\t")) {
            var v7 float32 = (-2.25 * 0.0)
            if len(v3) > 0 {
                v3[0] = (v4 ^ 477)
            }
            v1 += v1
            _ = v7
        }
    } else {
        switch "" {
        case "\n\t":
            _ = len(v3)
        default:
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 610}
                v4 = 35
                if (2 < 4294967295) {
                    if false {
                        var v10 int32 = int32(v4)
                        var v11 string = ("hello" + "hello")
                        var v12 int8 = -128
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi8
                _ = v9
            }
        }
        v1 = 228
    }
    _ = len(v3)
    v1 += v1
    v13 := GenS1{F0: 42}
    switch ("hello" + "") {
    case "\n\t":
        var v14 string = "\n\t"
        _ = v14
    case "a":
        var v15 int32 = -263
        _ = v15
    default:
        if len(v3) > 0 {
            v3[0] = (v4 >> ((255) & 63))
        }
    }
    var v16 float64 = ((float64(1.5) - (1e10 * 1.5)) - 0.1)
    var v17 float32 = (3.14159 / ((1e10 * (0.0 * (1e10 - (1e10 + 1.5)))) * ((((1.5 * float32(v16)) + (0.1 * (-2.25 - 0.0))) - -2.25) * -2.25)))
    if (!true) {
        var v18 uint = 569
        v16 = ((1.5 * ((float64(-2.25) + (((v16 / v16) * -2.25) * float64(float32(v16)))) - v16)) * (v16 / 3.14159))
        _ = v18
    }
    v4 -= (-(v4))
    var v19 float32 = v17
    v19 = (float32(v16) - ((float32(v16) * 1.5) / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
