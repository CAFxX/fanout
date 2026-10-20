package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]int)
    var v3 uint64 = 1
    var v4 uint64 = (v3 + 0)
    fmt.Println(v4, v3)
    var v5 int16 = -190
    if false {
        fmt.Println(v5)
        var v6 int = 0
        switch v6 {
        case -1:
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint8 = 128
                v9 := []int{}
                _ = fi7
                _ = v8
                _ = v9
            }
        case 127:
            var v10 float32 = -2.25
            _ = v10
        default:
            if true {
                var v11 float64 = 0.0
                v5 -= v5
                v1[(("a" + "foo bar") + ((("" + "") + "hello") + "hello"))] = (false && (389 <= 2))
                _ = v11
            }
        }
        _ = v6
    }
    v1[""] = (1 <= -2147483648)
    fmt.Println(v4, v3, v3)
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 float32 = 1e10
        var v14 uint64 = 2
        _ = fi12
        _ = v13
        _ = v14
    }
    if (9223372036854775807 < 967) {
        var v15 uint32 = 470
        v5 += 566
        _ = v15
    } else {
        var v16 float64 = 0.0
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v5 -= (v5 ^ (v5 ^ v5))
    if (2147483647 >= 2147483647) {
        for fi19 := 0; fi19 < 16; fi19++ {
            fmt.Println(fi19)
            v3 -= (^(v3))
            v5 += v5
            _ = fi19
        }
        v1["a"] = false
    } else {
        v5 = 42
    }
    v20 := GenS1{F0: 0, F1: 696, F2: 65535}
    var v21 float32 = ((((float32((float64((3.14159 / 3.14159)) * float64(float32(3.14159)))) / 1.0) - 1.5) * (float32(((0.0 / 0.1) + float64((-2.25 - ((0.1 * 3.14159) - (0.1 + (float32((3.14159 * (1.5 * (((((1.5 * 0.1) * 3.14159) / (1.5 / (float64((3.14159 + ((0.0 - 1.5) + (0.0 / (1e10 + (((-2.25 - ((-2.25 + (float32(3.14159) / -2.25)) / ((0.0 + (((1e10 * 3.14159) + 3.14159) * (((((-2.25 - (1e10 * ((1.5 * 0.1) - (1.5 * (0.1 / (1e10 - 1.5)))))) * 1e10) + 1e10) + (3.14159 * 0.1)) - float32(1.5)))) * (float32(-2.25) - 0.1)))) + -2.25) - 1e10)))))) * ((0.0 * -2.25) + ((1.5 * 1.5) / -2.25))))) * (((3.14159 - 3.14159) / 1.5) / (float64(0.0) - float64(0.1)))) * 1e10)))) - 1e10))))))) - 0.0)) - float32((0.1 + 0.1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
