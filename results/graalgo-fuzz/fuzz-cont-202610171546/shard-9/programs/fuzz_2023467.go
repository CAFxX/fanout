package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint = 927
    var v2 uint8 = 215
    v3 := make(map[uint32]uint32)
    var v4 int64 = 0
    v5 := make(map[string]bool)
    v6 := []int{0, 32767, 1, 127}
    fmt.Println(v4)
    v1 += 1
    v4 = v4
    v7 := GenS1{F0: (448 <= 866), F1: 2, F2: (v1 | v1)}
    _ = len(v5)
    v8 := GenS2{F0: (v2 - v2)}
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: ((1e10 > 3.14159) && (v2 <= v2)), F1: 255, F2: 2}
        delete(v3, 2147483647)
        _ = fi9
        _ = v10
    }
    switch 32768 {
    case 127:
        for fi11 := 0; fi11 < 1; fi11++ {
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
    case -44:
        var v13 float64 = 3.14159
        _ = v13
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
