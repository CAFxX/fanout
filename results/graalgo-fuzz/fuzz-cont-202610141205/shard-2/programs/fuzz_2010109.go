package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    v2 := []string{"", "\n\t", "\n\t"}
    var v3 uint = 65535
    v4 := make(map[int]uint32)
    v5 := []uint32{}
    var v6 uint32 = 2
    if (!false) {
        delete(v4, -2147483648)
        v6 = 293
    } else {
        switch "" {
        case "foo bar":
            var v7 uint64 = 782
            _ = v7
        default:
            var v8 float64 = 1e10
            _ = v8
        }
    }
    v9 := GenS1{F0: 42}
    v6 += (v6 | (v6 * v6))
    if true {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int8 = 42
            var v12 int = (^(fi10))
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS1{F0: 1}
            _ = fi13
            _ = v14
        }
        _ = len(v4)
    }
    v6 -= 2
    delete(v4, -437)
    v6 -= 0
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
