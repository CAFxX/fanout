package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 string }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 42
    var v3 int8 = 1
    v4 := make(map[string]int)
    fmt.Println(v2)
    var v5 uint64 = 2
    var v6 uint32 = 896
    v7 := GenS1{F0: (v1 / ((v1) | 1)), F1: 204}
    v8 := GenS2{F0: 255}
    v9 := GenS3{F0: 65535, F1: ((((("\n\t" + "\n\t") + "\n\t") + "hello") + "hello") + "a")}
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 int64 = -9223372036854775808
        _ = fi10
        _ = v11
    }
    var v12 float64 = -2.25
    v2 = 127
    fmt.Println(v6, v6, v2)
    if (v12 <= (float64(-2.25) / v12)) {
        var v13 uint16 = 0
        var v14 int8 = v3
        var v15 int64 = int64(v6)
        _ = v13
        _ = v14
        _ = v15
    }
    v12 = (v12 * (float64((-2.25 * -2.25)) * v12))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
