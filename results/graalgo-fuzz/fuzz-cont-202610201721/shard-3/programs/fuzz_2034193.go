package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 uint8 = 42
    var v2 uint16 = 143
    var v3 float32 = ((((1.5 - (0.1 - float32(1.5))) * (1e10 * 1.5)) / (0.0 - 3.14159)) * -2.25)
    v4 := []string{"a"}
    v5 := []int{-364, 88, 42, 9223372036854775807}
    var v6 float32 = v3
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    var v8 string = ("a" + "a")
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println(v1)
    if (-1 > -9223372036854775808) {
        var v9 uint8 = (v1 >> ((uint(v1)) & 63))
        _ = v9
    } else {
        var v10 uint = 0
        for fi11 := 0; fi11 < 8; fi11++ {
            if len(v5) > 0 {
                v5[0] = fi11
            }
            _ = len(v4)
            _ = fi11
        }
        _ = v10
    }
    fmt.Println(v6)
    _ = len(v4)
    v2 -= v2
    v12 := GenS1{F0: v2, F1: int8(v1)}
    var v13 float64 = 0.1
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS1{F0: (v2 ^ (^(v2))), F1: int8(v1)}
        _ = fi14
        _ = v15
    }
    v5 = append(v5, 0)
    fmt.Println(v2, v2)
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
