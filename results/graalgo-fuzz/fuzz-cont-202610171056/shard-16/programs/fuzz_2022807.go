package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 uint16; F2 string }
type GenS4 struct { F0 int8 }
func genFn0() uint32 {
    var v4 uint16 = 0
    v4 = (v4 * (^(v4)))
    v5 := GenS1{F0: 0}
    var v6 float32 = 1.5
    _ = v4
    _ = v5
    _ = v6
    return 2147483647
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    var v3 float64 = (1.5 / ((1e10 * (0.0 - 0.1)) * (1.5 - 3.14159)))
    v7 := GenS2{F0: "a", F1: 1}
    var v8 int16 = 42
    v9 := GenS2{F0: (("hello" + "hello") + "a"), F1: 9223372036854775807}
    var v10 bool = ((255 < 42) && (65535 != 0))
    var v11 string = "foo bar"
    var v12 uint = 255
    switch (v11 + v11) {
    case "a":
        v13 := []int{4294967295, -423}
        _ = v13
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            switch v11 {
            case "a":
                v1 = (v3 <= v3)
            case "foo bar":
                v15 := GenS3{F0: (v8 * (v8 ^ -85)), F1: 184, F2: v11}
                _ = v15
            default:
                _ = len(v2)
            }
            var v16 bool = (3.14159 < (1.5 + float32(0.1)))
            _ = fi14
            _ = v16
        }
    }
    fmt.Println(v11, v8)
    var v17 float64 = v3
    var v18 int8 = -128
    v19 := GenS4{F0: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
