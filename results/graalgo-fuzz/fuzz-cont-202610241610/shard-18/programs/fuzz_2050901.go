package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0() uint8 {
    var v6 float64 = 0.0
    var v7 int16 = -1
    var v8 uint16 = 42
    fmt.Println(v7, v6)
    if (!(!false)) {
        v9 := GenS1{F0: 1}
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    var v4 bool = true
    var v5 uint16 = 415
    fmt.Println(v3, v4)
    var v10 int16 = (-(v3))
    v5 = 65535
    switch "hello" {
    case "\n\t":
        v10 = (v10 + 0)
    }
    fmt.Println(v5)
    for fi11 := 0; fi11 < 10; fi11++ {
        v1[538] = v4
        var v12 uint16 = 2
        var v13 int16 = v10
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: v5}
    v15 := GenS1{F0: v5}
    for fi16 := 0; fi16 < 10; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            v10 -= 1
            _ = fi17
        }
        var v18 uint64 = 4294967295
        _ = fi16
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
