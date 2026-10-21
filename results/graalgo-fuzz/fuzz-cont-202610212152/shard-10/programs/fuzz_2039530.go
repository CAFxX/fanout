package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int64 }

func main() {
    var v1 uint16 = 2
    var v2 float64 = 0.0
    var v3 uint = 1
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    if false {
        v6 := GenS1{F0: 1}
        fmt.Println(v1)
        _ = v6
    } else {
        v7 := GenS2{F0: 937, F1: 128}
        v8 := []int{}
        v9 := GenS3{F0: 128, F1: 127}
        _ = v7
        _ = v8
        _ = v9
    }
    if (((-2.25 + ((((float32(((((v2 * (0.1 * float64((((1e10 / 1.5) - 1e10) * (((float32(-2.25) - 3.14159) / 1.5) - 0.0))))) - 0.0) + v2) + ((v2 * float64(0.1)) - (v2 / 1.5)))) / 3.14159) * 0.1) * 1e10) * 0.1)) - (0.0 + (float32(((float64(float32(v2)) * v2) + (0.1 / (float64((-2.25 / -2.25)) * ((((float64(-2.25) - 0.1) - (float64(1e10) - v2)) + -2.25) * ((v2 - (1e10 / v2)) - float64(-2.25))))))) + 1.5))) != 1.5) {
        for fi10 := 0; fi10 < 14; fi10++ {
            fmt.Println(v1, fi10, fi10)
            var v11 uint8 = 122
            _ = fi10
            _ = v11
        }
    }
    v1 += uint16(v3)
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int16 = 300
        var v14 int8 = 100
        v15 := GenS1{F0: int(v1)}
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 uint32 = 2147483647
    v17 := GenS4{F0: 0}
    fmt.Println(v3)
    v18 := v4[v16]
    var v19 uint32 = 208
    if (uint64(v19) == 304) {
        for fi20 := 0; fi20 < 3; fi20++ {
            fmt.Println(v1, v2, v2)
            _ = fi20
        }
        v2 += (v2 - v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
