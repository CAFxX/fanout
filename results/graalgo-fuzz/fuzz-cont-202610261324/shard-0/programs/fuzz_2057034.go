package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 65535
    v3 := make(map[string]bool)
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 int = -552
    fmt.Println(v2, v2)
    var v7 float64 = float64(float32(-2.25))
    v8 := GenS1{F0: -1, F1: 988}
    fmt.Println(v2)
    v2 -= (v2 - (^(v2)))
    v7 -= (1.5 + ((v7 / ((v7 * (v7 - v7)) - 1e10)) + 0.1))
    if (-1 > -251) {
        v9 := GenS1{F0: -128, F1: 300}
        v10 := GenS1{F0: 0, F1: 2147483647}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS1{F0: -1, F1: 1}
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                var v15 bool = true
                var v16 int = -1
                fmt.Println(v2)
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi13
        }
    }
    v2 -= (v2 / ((v2) | 1))
    fmt.Println(v2, v7, v6)
    var v17 uint32 = 4294967295
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS1{F0: -64, F1: (^(v17))}
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: 42, F1: v17}
    var v21 uint = v2
    v22 := []int{(v6 ^ v6)}
    if (1 == 2147483647) {
        v2 = v2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
