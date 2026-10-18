package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0(p0 bool, p1 uint16, p2 uint64) uint {
    var v6 uint = 0
    var v7 float64 = ((0.0 - ((0.1 * 1.5) + 0.1)) - (0.1 / 3.14159))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int8 = 1
    var v2 int8 = -103
    v3 := make(map[uint32]bool)
    var v4 uint8 = 1
    v5 := []uint32{531}
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: fi8}
        v10 := GenS1{F0: 32767}
        fmt.Println(fi8, v4, v1)
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 = (v2 << ((354) & 63))
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint16 = 32767
        v13 := GenS1{F0: (fi11 << ((0) & 63))}
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 962, F1: 840}
    fmt.Println(v4)
    v15 := GenS2{F0: 2147483647, F1: 881}
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    if (!false) {
        var v16 string = "a"
        v17 := GenS2{F0: 1, F1: -734}
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 float32 = 1e10
            fmt.Println(v19, v19)
            _ = fi18
            _ = v19
        }
        var v20 uint64 = 42
        for fi21 := 0; fi21 < 17; fi21++ {
            fmt.Println(v1)
            v1 = (v1 & (v1 - (v1 & v1)))
            _ = fi21
        }
        _ = v20
    }
    var v22 string = "foo bar"
    var v23 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
