package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 int16, p2 uint64) uint {
    p2 -= p2
    p1 = (p1 ^ 401)
    p2 = p2
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 float64 = float64((float32(1e10) * -2.25))
        var v6 float32 = float32(1e10)
        v6 -= (((((float32(v5) * -2.25) + 0.1) - v6) - ((v6 / 3.14159) / (3.14159 - -2.25))) * 3.14159)
        _ = fi4
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 int = 127
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    v1 = v1
    v1 = v1
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 42}
        var v10 int64 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 683}
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 uint32 = 65535
        switch "a" {
        case "\n\t":
            for fi14 := 0; fi14 < 1; fi14++ {
                v13 = 178
                fmt.Println(fi14, fi12, v13)
                _ = fi14
            }
        }
        _ = fi12
        _ = v13
    }
    var v15 int16 = 1000
    var v16 uint32 = 65535
    var v17 uint64 = 707
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
