package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }

func main() {
    var v1 int64 = 0
    v2 := make(map[uint32]bool)
    var v3 float64 = float64(float32(0.0))
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 int32 = 1
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        if true {
            fmt.Println(v3)
            fmt.Println(v3)
            fmt.Println(v1)
        }
        _ = fi6
    }
    if false {
        fmt.Println(v3)
        v3 = float64(1e10)
        v1 = v1
    }
    var v7 uint8 = 42
    fmt.Println(v3, v1, v7)
    v3 = v3
    v8 := GenS1{F0: -16, F1: 18446744073709551615, F2: 42}
    _ = len(v2)
    var v9 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
