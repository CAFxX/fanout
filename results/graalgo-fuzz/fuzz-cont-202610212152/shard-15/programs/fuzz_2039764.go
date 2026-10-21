package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    var v5 int32 = 846
    fmt.Println(v1)
    v3 += v3
    var v6 uint64 = 42
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: (v6 ^ v6)}
        v3 = (v3 * (v3 << ((4294967295) & 63)))
        _ = fi7
        _ = v8
    }
    v1 = float64(0.0)
    var v9 int16 = v3
    fmt.Println(v3)
    var v10 int = -122
    fmt.Println(v1)
    v1 -= (v1 * -2.25)
    if (!true) {
        var v11 int8 = -107
        v12 := []int{(v10 + 123), v10, (v10 | v10)}
        _ = v11
        _ = v12
    } else {
        var v13 int64 = -948
        v3 = v3
        if true {
            var v14 int16 = (v9 & v3)
            v3 -= v3
            var v15 uint64 = 42
            _ = v14
            _ = v15
        }
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
