package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int = 551
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 int32 = 2147483647
        _ = len(v4)
        v8 := v5[4294967295]
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = (float32(0.0) - (-2.25 - (0.0 * 3.14159)))
    v10 := []int{}
    v11 := GenS1{F0: -32768}
    var v12 int32 = -564
    fmt.Println(v9, v9)
    switch v2 {
    case 32768:
        if (!(3.14159 > -2.25)) {
            for fi13 := 0; fi13 < 3; fi13++ {
                v14 := GenS1{F0: -1}
                _ = fi13
                _ = v14
            }
        } else {
            var v15 bool = false
            _ = v15
        }
    case 32767:
        v2 += v2
    case 2147483647:
        v9 = v9
    }
    v16 := GenS2{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
