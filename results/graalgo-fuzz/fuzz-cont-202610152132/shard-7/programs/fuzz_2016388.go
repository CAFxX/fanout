package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint8 = 42
    var v2 uint = 808
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    var v6 uint32 = 4294967295
    fmt.Println(v3, v6, v3)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int = (fi7 >> (((v2 * (v2 & (v2 + v2)))) & 63))
        var v9 int = (fi7 + (fi7 >> (((v2 / (((^(v2))) | 1))) & 63)))
        v5[(v8 | (v9 >> ((1) & 63)))] = ((v3 < (float32(0.1) - (1e10 * float32(-2.25)))) && false)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v4[int(v6)] = false
        for fi11 := 0; fi11 < 16; fi11++ {
            v3 = float32(0.1)
            _ = fi11
        }
        _ = len(v4)
        _ = fi10
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int32 = 42
        _ = fi12
        _ = v13
    }
    v14 := []int{314, -439, -843, 32768}
    v15 := GenS1{F0: (v6 | (v6 << (((^(v2))) & 63)))}
    v1 += v1
    v6 = (v6 / ((v6) | 1))
    v1 += (v1 ^ 243)
    fmt.Println(v3)
    v16 := v5[-9223372036854775808]
    if (v16 && v16) {
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    v18 := GenS2{F0: 42, F1: 372}
    v3 = (v3 - (v3 + float32(0.1)))
    var v19 int16 = 453
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
