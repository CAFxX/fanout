package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int32 = -813
    v2 := []uint32{2}
    var v3 float64 = (float64(((1.5 - float32(0.1)) + 1e10)) - 0.0)
    var v4 int32 = (v1 << ((4294967295) & 63))
    v2 = append(v2, 2147483648)
    if len(v2) > 0 {
        v2[0] = 2
    }
    var v5 uint64 = 18446744073709551615
    v6 := GenS1{F0: (v5 | 570)}
    fmt.Println(v1)
    v1 += (v4 + (v1 - -767))
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "foo bar" {
        case "a":
            var v8 int = 42
            _ = v8
        default:
            v9 := GenS1{F0: 2}
            _ = v9
        }
        var v10 uint32 = 155
        _ = fi7
        _ = v10
    }
    if true {
        v4 -= (v1 << ((861) & 63))
    }
    v1 = (v1 >> ((842) & 63))
    if (3.14159 > 0.1) {
        v2 = append(v2, 42)
        v2 = append(v2, 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
