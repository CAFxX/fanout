package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 0.0
    var v3 uint = 333
    var v4 float32 = (v2 * (0.0 + float32(3.14159)))
    var v5 float64 = (0.0 - 3.14159)
    v3 += v3
    var v6 uint16 = 93
    fmt.Println(v4)
    fmt.Println(v5, v6, v3)
    v7 := GenS1{F0: 43}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int16 = 0
            v4 = (v4 / (float32(v5) / v4))
            var v11 float32 = v4
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{fi8, -436, fi8, (fi8 << ((v3) & 63))}
        _ = fi8
        _ = v12
    }
    v5 -= v5
    var v13 float64 = v5
    fmt.Println(v4, v13, v2)
    if true {
        if (0.1 == v5) {
            var v14 int32 = 1
            _ = v14
        } else {
            fmt.Println(v5)
            if false {
                v1[778] = true
            } else {
                v6 = 231
                v5 -= (1.5 * v13)
            }
            v15 := GenS2{F0: v6, F1: 127, F2: 42}
            _ = v15
        }
        var v16 uint32 = 291
        _ = v16
    }
    v13 += v5
    var v17 uint16 = 199
    v18 := GenS1{F0: -1}
    switch "foo bar" {
    case "foo bar":
        v17 = (v6 | v17)
    case "hello":
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 uint = 0
            for fi21 := 0; fi21 < 9; fi21++ {
                v5 += (((v5 + float64((float32((v5 * (v5 / (float64(v4) * ((v5 + v5) + 1e10))))) * (((v2 - (v2 - 0.0)) - (((0.1 + v4) / -2.25) - v2)) * (((-2.25 / (((float32(v5) / v4) / (float32(v13) / v2)) - (v2 * v4))) * v4) + (float32(v5) - (((((v2 + ((float32((1.5 * v5)) * -2.25) - v2)) - ((1e10 + v2) - 0.1)) * (float32((float64(-2.25) + float64((0.1 * (0.0 + v4))))) * v2)) - (v2 * 1e10)) * (v4 * -2.25)))))))) + (((0.0 + (3.14159 / v5)) / v13) - v13)) + (3.14159 - 1.5))
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    case "a":
        var v22 float64 = 1e10
        _ = v22
    }
    fmt.Println(v6, v3, v5)
    var v23 uint32 = 2147483647
    var v24 uint = (^(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
