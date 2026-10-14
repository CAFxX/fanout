package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint16 = 874
    var v2 uint = 255
    v3 := make(map[string]bool)
    var v4 int = -2147483648
    var v5 int8 = -108
    v4 = (-(v4))
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 float64 = ((((((3.14159 - (1e10 / 1.0)) / 1.0) + (0.0 / -2.25)) * 0.1) * (-2.25 - -2.25)) * -2.25)
        fmt.Println(v4)
        v7 = 0.0
        _ = fi6
        _ = v7
    }
    v2 -= 964
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int16 = 32767
        _ = fi8
        _ = v9
    }
    if ((-534 <= int32(v2)) && (3.14159 != (-2.25 + (((0.1 / (1e10 / ((3.14159 / 1.0) - 0.1))) - ((float64((0.1 + -2.25)) / 1e10) - 1e10)) * ((0.1 + ((1e10 / 1.5) + 3.14159)) * -2.25))))) {
        v10 := GenS1{F0: (v5 / ((42) | 1)), F1: (!false)}
        v2 += 65535
        _ = v10
    }
    v5 = v5
    for fi11 := 0; fi11 < 4; fi11++ {
        v2 = 18446744073709551615
        _ = len(v3)
        _ = fi11
    }
    v12 := GenS2{F0: -729, F1: -32768}
    v1 = v1
    v13 := GenS2{F0: 9223372036854775807, F1: 813}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
