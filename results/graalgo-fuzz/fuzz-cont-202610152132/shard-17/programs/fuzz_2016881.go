package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = (0.1 >= float32((3.14159 * 1.5)))
    v2 := make(map[int]string)
    v3 := make(map[uint32]int)
    v4 := make(map[string]string)
    var v5 int = 32767
    var v6 uint16 = 65535
    if (-1 > -878) {
        v7 := GenS1{F0: 255}
        _ = v7
    } else {
        var v8 uint16 = (v6 | v6)
        v5 += v5
        v9 := GenS1{F0: 42}
        _ = v8
        _ = v9
    }
    var v10 uint64 = 934
    v5 = (^(v5))
    for fi11 := 0; fi11 < 19; fi11++ {
        switch "\n\t" {
        case "\n\t":
            v1 = (!v1)
        }
        _ = fi11
    }
    var v12 uint32 = 1
    _ = len(v3)
    v12 += (v12 / ((v12) | 1))
    fmt.Println(v10, v12)
    v10 -= (v10 - 42)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
