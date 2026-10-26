package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    v3 := []uint32{2147483647}
    v4 := []uint32{243, 65535, 1}
    v2 = v2
    var v5 float64 = 0.1
    if (v5 == ((v2 - v2) * 1e10)) {
        var v6 int16 = -442
        var v7 int = -882
        v6 = (^(v6))
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 4294967295, F1: 42}
    v5 = v2
    v2 += 0.0
    v3 = append(v3, 42)
    v9 := GenS1{F0: 65535, F1: 0}
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v2 += (float64(float32(1.5)) - float64(float32(float64(-2.25))))
            _ = fi11
        }
        v3 = append(v3, 4294967295)
        var v12 uint16 = 32767
        _ = fi10
        _ = v12
    }
    v13 := GenS2{F0: 1, F1: 42}
    fmt.Println(v2, v2)
    if (686 != 65535) {
        fmt.Println(v5, v5)
    } else {
        delete(v1, 2147483648)
        v2 = (v2 + 1e10)
        var v14 uint8 = 128
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
