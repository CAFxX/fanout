package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 uint }

func main() {
    var v1 bool = (0 <= -2147483648)
    var v2 float64 = 3.14159
    v3 := []uint32{65535, 2147483647, 2147483647}
    var v4 float64 = (v2 - (0.0 + v2))
    fmt.Println(v1)
    fmt.Println(v2, v4)
    v5 := GenS1{F0: v1, F1: 2, F2: -240}
    var v6 int64 = 610
    var v7 float64 = (v4 - 0.0)
    v8 := GenS1{F0: v1, F1: 4294967295, F2: v6}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS2{F0: 186, F1: 32767}
        if (0 < 4294967295) {
            v7 += float64(0.1)
        } else {
            fmt.Println(v6, v1)
            v4 = float64(1.5)
            for fi11 := 0; fi11 < 9; fi11++ {
                v7 = 0.0
                v12 := GenS3{F0: ((255 <= 128) || false), F1: 18446744073709551615, F2: 1}
                _ = fi11
                _ = v12
            }
        }
        v6 = -444
        _ = fi9
        _ = v10
    }
    var v13 float32 = (0.0 * ((((((-2.25 * 0.1) + 1.5) / -2.25) + (float32(v2) + 3.14159)) * (3.14159 + (-2.25 - 0.0))) * ((1e10 * -2.25) - -2.25)))
    for fi14 := 0; fi14 < 8; fi14++ {
        fmt.Println(v2, fi14)
        switch ((("hello" + "") + ("hello" + "foo bar")) + "") {
        case "hello":
            v4 += float64(v13)
        default:
            for fi15 := 0; fi15 < 13; fi15++ {
                v1 = ((-32768 != -32768) && (false || v1))
                v13 -= 0.0
                _ = fi15
            }
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
