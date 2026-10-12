package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint = 2
    v2 := make(map[string]int)
    var v3 int8 = 92
    fmt.Println(v1, v3)
    v1 -= (v1 * 153)
    v3 += v3
    v3 = v3
    v4 := GenS1{F0: 255, F1: 255}
    switch -539 {
    case 102:
        v3 -= v3
    case -297:
        var v5 float64 = 1e10
        _ = v5
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        fmt.Println(v3)
        _ = fi6
    }
    var v7 string = "hello"
    if (42 < 1) {
        v7 = v7
        var v8 float64 = 1.5
        if false {
            var v9 uint16 = 0
            v10 := []int{32767, 32767, 326}
            _ = v10
            _ = v9
        } else {
            var v11 float32 = 1e10
            for fi12 := 0; fi12 < 6; fi12++ {
                for fi13 := 0; fi13 < 6; fi13++ {
                    if false {
                        v11 += 0.0
                        var v14 float64 = v8
                        _ = v14
                    }
                    _ = len(v2)
                    v3 -= v3
                    _ = fi13
                }
                _ = fi12
            }
            _ = v11
        }
        _ = v8
    }
    fmt.Println(v7)
    fmt.Println(v1, v7, v7)
    v15 := GenS1{F0: 1, F1: 42}
    var v16 int = -382
    v16 = v16
    fmt.Println(v16, v1, v3)
    var v17 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
