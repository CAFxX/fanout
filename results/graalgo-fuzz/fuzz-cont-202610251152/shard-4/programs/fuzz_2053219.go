package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int32 }

func main() {
    v1 := []uint32{0, 2147483648}
    var v2 float64 = (-2.25 / 1.0)
    v3 := []int{128, -2147483648, 42}
    v4 := make(map[string]bool)
    v1 = append(v1, 65535)
    var v5 uint16 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = ((1.5 + ((v2 - (float64(-2.25) - ((1e10 - 0.1) + v2))) + float64(1e10))) * ((float64(1.5) * v2) / v2))
    _ = len(v4)
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            switch "\n\t" {
            case "foo bar":
                v5 += v5
            case "":
                var v8 int32 = -2147483648
                _ = v8
            case "a":
                v5 = v5
            }
            _ = fi7
        }
        v4["a"] = (2 != 2)
        _ = fi6
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        _ = len(v4)
        v5 = v5
        v1 = append(v1, 65535)
        _ = fi9
    }
    var v10 int = -9223372036854775808
    v3 = append(v3, v10)
    v11 := GenS1{F0: 0, F1: 962, F2: 315}
    v2 = float64(((0.0 - float32(v2)) + (-2.25 / float32(v2))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
