package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 376
    var v4 bool = (((((0.1 / -2.25) * 1e10) + 1e10) - 1.5) <= (1e10 * (-2.25 + 1.5)))
    fmt.Println(v3)
    v5 := GenS1{F0: 20, F1: 24}
    v6 := GenS1{F0: 131, F1: 134}
    v4 = true
    v4 = (42 < 127)
    switch 128 {
    case 128:
        var v7 uint64 = 740
        _ = v7
    case 268:
        var v8 bool = v4
        _ = v8
    default:
        var v9 uint16 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            if true {
                v4 = (1 <= 255)
                v12 := GenS2{F0: 1, F1: -1}
                _ = v12
            }
            fmt.Println(v4, v4, v4)
            fmt.Println(v3)
            _ = fi11
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        delete(v2, 893)
        _ = fi13
    }
    v4 = false
    fmt.Println(v4, v4)
    for fi14 := 0; fi14 < 18; fi14++ {
        v4 = true
        _ = fi14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        var v16 int64 = -46
        v16 = v16
        _ = fi15
        _ = v16
    }
    var v17 int64 = 1
    delete(v2, 4294967295)
    delete(v1, -478)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
