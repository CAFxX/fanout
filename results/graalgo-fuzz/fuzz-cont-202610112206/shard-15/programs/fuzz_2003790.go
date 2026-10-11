package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
func genFn0(p0 uint16) uint8 {
    var v5 bool = (-2.25 == 1e10)
    v5 = ((-1 <= -9223372036854775808) || false)
    v6 := GenS1{F0: -73, F1: "hello"}
    switch 32768 {
    case 9223372036854775807:
        v5 = (-527 != 1)
    default:
        var v7 int8 = 42
        _ = v7
    }
    var v8 uint64 = 798
    _ = p0
    _ = v5
    _ = v6
    _ = v8
    return 128
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 115
    var v3 uint = 255
    var v4 int64 = -9223372036854775808
    v9 := []int{2147483647, int(v2)}
    var v10 float64 = 1.5
    var v11 uint16 = 840
    fmt.Println(v1, v11, v4)
    v10 += (((0.1 / (v10 * (1e10 + v10))) - float64(-2.25)) / 1.5)
    v2 += genFn0(255)
    _ = len(v9)
    switch "" {
    case "hello":
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    v4 += v4
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 float32 = (1.5 * (0.0 + 3.14159))
        var v14 int16 = 32767
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = len(v9)
    if false {
        if (-1 != -964) {
            for fi15 := 0; fi15 < 15; fi15++ {
                v10 += (1e10 * float64(1.5))
                var v16 uint32 = v1
                var v17 int64 = 1
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 float32 = float32((((float64(float32(((float64(1.5) * v10) - float64(0.1)))) - 0.0) * float64(0.0)) + ((v10 * (v10 - (v10 * -2.25))) / v10)))
            _ = v18
        }
        var v19 uint64 = 216
        var v20 uint16 = uint16(v3)
        _ = v19
        _ = v20
    }
    v11 += v11
    fmt.Println(v11, v3)
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int8 = 127
        var v23 string = "\n\t"
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
