package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }

func main() {
    v1 := []uint32{0, 2147483647, 1}
    var v2 int16 = 32767
    var v3 float32 = float32(3.14159)
    var v4 uint32 = 65535
    v5 := []uint32{}
    v2 = (v2 ^ v2)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 float64 = -2.25
    fmt.Println(v2, v4)
    v6 -= (v6 + v6)
    v3 -= float32(v6)
    v3 = (float32(0.0) + v3)
    v4 -= (v4 | v4)
    if false {
        if (-32768 >= (^(v2))) {
            v6 += v6
            var v7 int8 = 100
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                v2 -= -618
                _ = fi8
            }
            v2 = (v2 * v2)
        }
        v3 -= v3
        var v9 int64 = -837
        _ = v9
    }
    v2 -= v2
    v3 = float32(v6)
    var v10 int16 = v2
    switch ("a" + "a") {
    case "":
        var v11 uint8 = 1
        _ = v11
    case "a":
        switch "" {
        case "a":
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(v3, fi12, v4)
                var v13 int64 = 9223372036854775807
                _ = fi12
                _ = v13
            }
        }
    default:
        var v14 uint = 2
        _ = v14
    }
    v15 := GenS1{F0: 711, F1: 792, F2: "\n\t"}
    v1 = append(v1, (v4 | v4))
    fmt.Println(v3, v2, v2)
    var v16 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
