package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 int64, p1 bool, p2 bool) uint16 {
    p0 = (p0 << ((592) & 63))
    var v5 float32 = float32(0.0)
    var v6 int16 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 209
}

func main() {
    var v1 uint8 = 128
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 int64 = 0
    var v7 float64 = (float64(((((0.1 - 0.1) - (0.1 * 1e10)) / 0.1) - (1e10 / float32(-2.25)))) / 1.5)
    v2 = (42 == 42)
    v8 := v3[9223372036854775807]
    var v9 uint64 = 18446744073709551615
    fmt.Println(v4, v4)
    v10 := v3[-9223372036854775808]
    for fi11 := 0; fi11 < 12; fi11++ {
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v9, v10)
        case "a":
            if (!true) {
                for fi12 := 0; fi12 < 8; fi12++ {
                    var v13 int8 = 100
                    _ = fi12
                    _ = v13
                }
                var v14 float32 = 0.0
                v15 := GenS1{F0: (v9 | 784), F1: 9223372036854775807}
                _ = v14
                _ = v15
            }
        case "hello":
            v1 = v1
        }
        _ = fi11
    }
    fmt.Println(v2, v10)
    var v16 uint64 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
