package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []int{423, -234, 127, 0}
    v2 := make(map[int]bool)
    var v3 float64 = 1.5
    v3 = ((0.0 / v3) * v3)
    v2[764] = (65535 >= 2)
    _ = len(v1)
    var v4 uint32 = 4294967295
    v4 = 65535
    v3 -= float64(-2.25)
    fmt.Println(v3, v4)
    v1 = append(v1, 127)
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 17; fi6++ {
        v3 = v3
        _ = fi6
    }
    if (653 < 272) {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int8 = -128
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 int16 = 42
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: (v4 | (v4 - (v4 * (v4 >> ((5) & 63)))))}
        var v13 uint32 = 2147483647
        for fi14 := 0; fi14 < 14; fi14++ {
            delete(v2, 127)
            _ = fi14
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
