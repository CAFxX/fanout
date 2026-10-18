package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    v5 := []string{}
    var v6 uint = 76
    v2 = v2
    if true {
        _ = len(v3)
        var v7 uint8 = 0
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint32 = 612
        _ = fi8
        _ = v9
    }
    v10 := v3["a"]
    v2 = float32((-2.25 + 1e10))
    _ = len(v5)
    v10 += 4294967295
    var v11 uint8 = 1
    if (true || ((^(v6)) < (^(v6)))) {
        v12 := GenS1{F0: 692}
        _ = v12
    } else {
        var v13 int64 = 9223372036854775807
        fmt.Println(v10, v6, v11)
        if (v13 <= v13) {
            fmt.Println(v10, v2, v10)
        } else {
            if (false || (817 > 1)) {
                if (217 <= 519) {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        fmt.Println(v13, v2, v13)
                        var v15 int = (-(fi14))
                        v4[""] = (v10 - (v10 / ((v10) | 1)))
                        _ = fi14
                        _ = v15
                    }
                }
            } else {
                v11 += v11
            }
            v16 := GenS1{F0: 255}
            v2 = 0.1
            _ = v16
        }
        _ = v13
    }
    var v17 uint8 = v11
    v6 -= (v6 >> (((v6 ^ (v6 * uint(v10)))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
