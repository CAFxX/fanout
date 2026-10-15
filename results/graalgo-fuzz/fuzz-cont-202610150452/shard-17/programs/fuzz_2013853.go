package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    var v2 bool = (100 <= -128)
    v3 := []int{32768, 127}
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 1, F1: 65535, F2: 1}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        delete(v4, "a")
        var v9 string = "\n\t"
        v10 := GenS1{F0: 42, F1: 887, F2: 255}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 128, F1: 850, F2: 1}
    v2 = v2
    for fi12 := 0; fi12 < 4; fi12++ {
        v2 = true
        v1 += v1
        _ = fi12
    }
    switch "a" {
    case "a":
        switch -9223372036854775808 {
        case 0:
            v1 = (v1 / ((42) | 1))
        }
    }
    fmt.Println(v1, v1)
    for fi13 := 0; fi13 < 10; fi13++ {
        if (-2.25 != 0.1) {
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int = (fi13 - fi13)
                var v16 uint8 = 128
                if v2 {
                    v3 = append(v3, (fi14 - (-(fi13))))
                    v17 := GenS1{F0: 255, F1: 255, F2: 0}
                    _ = v17
                }
                var v18 uint32 = 72
                _ = fi14
                _ = v15
                _ = v16
                _ = v18
            }
        }
        var v19 int16 = 0
        _ = fi13
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
