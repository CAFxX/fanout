package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint16 = 2
    var v2 int64 = -1
    v3 := make(map[uint32]string)
    var v4 float64 = 1e10
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 uint64 = 18446744073709551615
        var v7 int64 = 9223372036854775807
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int64 = -1
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: ((255 > 255) && true)}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 bool = true
        _ = fi12
        _ = v13
    }
    switch 127 {
    case 9223372036854775807:
        v4 = v4
    case -632:
        if false {
            v14 := []int{}
            fmt.Println(v4)
            v2 += v2
            _ = v14
        }
    case 637:
        var v15 float32 = 1.5
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            v18 := []int{9223372036854775807, 924, (fi17 - fi17)}
            _ = fi17
            _ = v18
        }
        v2 = (v2 << ((0) & 63))
        _ = fi16
    }
    v19 := []int{2147483647, 2147483647, 127}
    v20 := GenS1{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
