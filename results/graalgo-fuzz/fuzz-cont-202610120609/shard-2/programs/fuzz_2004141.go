package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 bool = true
    v6 := GenS1{F0: 1, F1: 2147483648}
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 902, F1: 0}
        var v9 float64 = -2.25
        v10 := GenS1{F0: 32767, F1: 272}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: 478, F1: 0}
    var v12 int8 = 127
    fmt.Println(v5, v5)
    v13 := v4[4294967295]
    fmt.Println(v13, v13)
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            switch v13 {
            case "":
                v13 = "a"
            default:
                v12 = 127
            }
            var v16 float64 = (float64((1.5 - (3.14159 * ((3.14159 - -2.25) - (1.5 * float32((3.14159 + ((0.1 - 3.14159) - 1e10)))))))) - -2.25)
            var v17 uint = 18446744073709551615
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int64 = 9223372036854775807
        _ = fi14
        _ = v18
    }
    fmt.Println(v13, v13, v12)
    v13 = "hello"
    v19 := []int{-2147483648, 114, 128, 412}
    switch 9223372036854775807 {
    case 127:
        v13 = v13
    case 9223372036854775807:
        fmt.Println(v12)
    default:
        var v20 int64 = -1
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
