package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint64 = 4294967295
    var v3 int16 = -1
    var v4 uint8 = 128
    fmt.Println(v1, v2, v2)
    v5 := GenS1{F0: 42}
    v4 = (v4 / (((^(v4))) | 1))
    v2 += (v2 << ((0) & 63))
    var v6 uint16 = 0
    switch "hello" {
    case "hello":
        var v7 int16 = v3
        _ = v7
    }
    v8 := []int{32767, 42}
    v4 -= (v4 ^ v4)
    if (v6 != (^(v6))) {
        _ = len(v8)
        var v9 float32 = ((1e10 - 3.14159) - 0.1)
        _ = v9
    } else {
        fmt.Println(v2, v6, v6)
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v1 -= (v1 + v1)
        for fi11 := 0; fi11 < 13; fi11++ {
            v1 = uint32(v6)
            _ = fi11
        }
        if (42 != -2147483648) {
            v6 -= 32767
            var v12 uint8 = v4
            _ = v12
        } else {
            var v13 int16 = -32768
            _ = v13
        }
        _ = fi10
    }
    var v14 int64 = -902
    v4 += (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
