package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = (1.5 + (-2.25 * (float64(1.5) + 0.1)))
    var v3 float32 = 1.5
    var v4 int16 = 1000
    for fi5 := 0; fi5 < 20; fi5++ {
        if true {
            var v6 uint32 = 2147483647
            var v7 float64 = 3.14159
            _ = v6
            _ = v7
        }
        var v8 int = (-(fi5))
        var v9 float64 = (((v2 * (1e10 - (((0.1 / 0.1) * v2) * ((float64(v3) * ((float64((v3 + 1.5)) / (((v2 + ((0.0 - v2) - (1.5 * ((0.1 - ((v2 + v2) * 1.5)) / (float64(-2.25) / v2))))) - ((v2 + (0.0 + float64(v3))) * v2)) - v2)) * v2)) - v2)))) / float64((float32(0.0) / (1.5 - (v3 - float32(v2)))))) - 1e10)
        _ = fi5
        _ = v8
        _ = v9
    }
    switch 2147483647 {
    case 233:
        v3 -= -2.25
    case 32767:
        switch 4294967295 {
        case 32767:
            var v10 int = 0
            _ = v10
        case -2147483648:
            v11 := []int{}
            _ = v11
        }
    case 4294967295:
        var v12 uint64 = 48
        _ = v12
    default:
        v13 := v1[9223372036854775807]
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 uint32 = 1
        var v16 int = 0
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint32 = 2147483648
    v1[4294967295] = v17
    fmt.Println(v3, v17, v17)
    v3 = (float32(((float64(v3) - v2) + v2)) + v3)
    v4 = 42
    var v18 float32 = 1.5
    var v19 int = 128
    fmt.Println(v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
