package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }

func main() {
    var v1 uint32 = 880
    var v2 uint32 = (v1 << ((0) & 63))
    var v3 int16 = 42
    v4 := make(map[uint32]bool)
    var v5 uint = 686
    var v6 int8 = -96
    v6 = v6
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint16 = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (v3 % ((200) | 1)), F1: -1, F2: 1}
    fmt.Println(v5, v3)
    v5 = 4294967295
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            _ = len(v4)
            fmt.Println(fi10, v5)
            _ = fi11
        }
        var v12 int16 = (v3 | v3)
        _ = fi10
        _ = v12
    }
    _ = len(v4)
    for fi13 := 0; fi13 < 3; fi13++ {
        v1 -= 1
        _ = fi13
    }
    v6 += 0
    var v14 uint8 = uint8(v2)
    var v15 int16 = int16(v1)
    v2 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
