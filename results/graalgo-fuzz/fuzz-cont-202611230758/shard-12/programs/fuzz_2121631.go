package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483647, 490, 549}
    var v2 int8 = -21
    v3 := make(map[string]int)
    var v4 uint8 = 52
    var v5 int = 605
    v6 := make(map[uint32]bool)
    var v7 string = ("hello" + "hello")
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = float32(0.1)
        if (false || true) {
            v10 := GenS1{F0: -1}
            fmt.Println(v4, fi8, v9)
            var v11 float64 = float64(v9)
            _ = v10
            _ = v11
        } else {
            v2 += v2
        }
        _ = fi8
        _ = v9
    }
    var v12 int16 = 1000
    _ = len(v6)
    var v13 uint32 = 2147483648
    v5 += 500
    fmt.Println(v13, v2)
    v2 -= v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
