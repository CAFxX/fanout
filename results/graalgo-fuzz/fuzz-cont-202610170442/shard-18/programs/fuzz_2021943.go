package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = float64((float32((((-2.25 * 0.0) - 1.5) / float64(0.1))) * 1e10))
    var v3 uint = 776
    var v4 uint8 = 109
    var v5 float64 = v2
    var v6 int8 = 127
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int32 = 2147483647
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 304
    v2 = (v2 + (float64((3.14159 - (float32(v5) / (float32(v5) + float32(v5))))) * v5))
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: -1, F1: 32767, F2: 0}
        v3 = uint(v9)
        _ = fi10
        _ = v11
    }
    delete(v1, -246)
    v12 := []int{127, 4294967295, 32768, -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
