package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []uint32{2147483648, 42, 65535, 2}
    var v2 bool = false
    var v3 int16 = -32768
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]bool)
    v6 := []string{"", "\n\t"}
    v5[2147483648] = false
    v7 := GenS1{F0: 2}
    v8 := GenS1{F0: 2}
    var v9 int8 = -128
    fmt.Println(v3)
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := GenS2{F0: 116}
        v12 := GenS2{F0: 18446744073709551615}
        _ = fi10
        _ = v11
        _ = v12
    }
    v9 += (v9 + (v9 * v9))
    v9 = -35
    _ = len(v1)
    var v13 int64 = 1
    for fi14 := 0; fi14 < 17; fi14++ {
        if v2 {
            for fi15 := 0; fi15 < 8; fi15++ {
                v6 = append(v6, "a")
                _ = fi15
            }
        }
        v13 = (v13 >> ((42) & 63))
        _ = fi14
    }
    fmt.Println(v2, v9, v2)
    var v16 uint64 = 18446744073709551615
    v17 := v4["foo bar"]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
