package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 1
    v3 := []uint32{598}
    v2 = (v2 * v2)
    fmt.Println(v2)
    _ = len(v3)
    v4 := GenS1{F0: 2, F1: 884, F2: (-28 != -1)}
    var v5 int32 = -2147483648
    var v6 uint16 = (v2 ^ v2)
    for fi7 := 0; fi7 < 12; fi7++ {
        if ((!true) && false) {
            var v8 uint = 2
            fmt.Println(fi7)
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(fi7, fi9)
                if len(v3) > 0 {
                    v3[0] = 42
                }
                _ = fi9
            }
            _ = v8
        }
        _ = fi7
    }
    if false {
        var v10 float32 = (1e10 / 3.14159)
        for fi11 := 0; fi11 < 9; fi11++ {
            v2 += (v6 >> ((153) & 63))
            _ = fi11
        }
        _ = v10
    }
    v6 -= v2
    v6 = (v6 | 146)
    var v12 uint32 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
