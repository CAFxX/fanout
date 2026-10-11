package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }

func main() {
    var v1 bool = (1e10 < (((-2.25 - 3.14159) - 1e10) / -2.25))
    var v2 uint8 = 0
    var v3 uint = 255
    v4 := make(map[uint32]string)
    v5 := make(map[string]uint32)
    v6 := []int{-9223372036854775808, -2147483648, 9223372036854775807, 239}
    if false {
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 string = "a"
                _ = fi8
                _ = v9
            }
            v1 = (40 <= 42)
            _ = fi7
        }
    }
    v3 -= 584
    v4[752] = "a"
    for fi10 := 0; fi10 < 18; fi10++ {
        v6 = append(v6, 9223372036854775807)
        fmt.Println(v1, v1)
        if false {
            _ = len(v4)
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 int = 32767
                _ = fi11
                _ = v12
            }
        }
        _ = fi10
    }
    v1 = (0.1 <= -2.25)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v13 int16 = -415
    v2 -= v2
    v14 := GenS1{F0: 771, F1: "", F2: 94}
    v3 += (^(v3))
    var v15 float64 = float64(float32(float64(3.14159)))
    if (((v2 * v2) < 114) || (4294967295 <= 2)) {
        fmt.Println(v15)
        var v16 uint64 = 18446744073709551615
        _ = v16
    }
    var v17 float32 = 3.14159
    for fi18 := 0; fi18 < 1; fi18++ {
        v15 = v15
        v19 := GenS1{F0: 2147483647, F1: "a", F2: 899}
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
