package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    v3 := []uint32{}
    var v4 int8 = 1
    v5 := GenS1{F0: -32768}
    fmt.Println(v1, v1)
    var v6 int = 2147483647
    if (1e10 > (1e10 / 1.0)) {
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int = (v6 - 602)
            _ = fi7
            _ = v8
        }
        v2[2] = v6
        v9 := v2[4294967295]
        _ = v9
    }
    v1 = v1
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint16 = 65535
            v12 += (v12 >> ((0) & 63))
            var v13 int = 32768
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v4, v1)
        _ = fi10
    }
    if true {
        var v14 uint = 42
        v15 := GenS1{F0: 0}
        v14 = (v14 >> (((v14 * v14)) & 63))
        _ = v14
        _ = v15
    } else {
        var v16 uint16 = 42
        var v17 uint = 4294967295
        _ = v16
        _ = v17
    }
    v4 += -70
    var v18 int8 = v4
    fmt.Println(v6, v4)
    v19 := GenS1{F0: 42}
    var v20 uint64 = 0
    var v21 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
