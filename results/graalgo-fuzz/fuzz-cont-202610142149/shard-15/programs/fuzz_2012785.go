package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483647}
    v1 = append(v1, "hello")
    var v4 float32 = 0.1
    v4 -= (v4 * (v4 / 0.1))
    v4 = 1e10
    var v5 uint32 = 0
    v4 += v4
    for fi6 := 0; fi6 < 2; fi6++ {
        if false {
            var v7 int64 = -1
            switch 32767 {
            case 107:
                var v8 uint32 = 2147483647
                _ = v8
            }
            _ = v7
        }
        if true {
            if (-1 >= -32768) {
                var v9 uint8 = 89
                fmt.Println(v9)
                v10 := GenS1{F0: v9}
                _ = v10
                _ = v9
            }
            v4 = ((0.0 / v4) * 1e10)
        }
        _ = fi6
    }
    v5 = 4294967295
    v11 := GenS1{F0: 127}
    v4 = (v4 - 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
