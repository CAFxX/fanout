package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint8 }
func genFn0(p0 uint8) uint16 {
    fmt.Println(p0)
    var v6 uint32 = 42
    v7 := GenS1{F0: v6, F1: 2147483648, F2: p0}
    _ = p0
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 float64 = (1.5 - -2.25)
    var v2 uint16 = 42
    var v3 int64 = 32
    v4 := make(map[string]bool)
    var v5 float64 = v1
    v8 := GenS1{F0: 2147483647, F1: 0, F2: 128}
    v4["hello"] = false
    var v9 uint32 = 1
    v10 := GenS1{F0: (v9 | v9), F1: (v9 | v9), F2: 127}
    v11 := GenS1{F0: 42, F1: 2147483648, F2: 193}
    for fi12 := 0; fi12 < 8; fi12++ {
        v9 -= 2
        for fi13 := 0; fi13 < 18; fi13++ {
            v1 -= v5
            v14 := GenS1{F0: v9, F1: v9, F2: 2}
            delete(v4, "a")
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    fmt.Println(v2)
    var v15 uint32 = v9
    v15 -= v9
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
