package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 float32 = 1.5
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    v5 := []string{"", "a", "foo bar", ""}
    var v6 int64 = -94
    var v7 int32 = 0
    for fi8 := 0; fi8 < 20; fi8++ {
        if ((((((((((0.1 / (0.0 * float64(v1))) / 1.0) + float64(1.5)) + (-2.25 - -2.25)) + float64((-2.25 + (v1 + (3.14159 * float32(float64(-2.25))))))) - (-2.25 - -2.25)) * 1e10) + (-2.25 * -2.25)) - 0.0) <= 0.1) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v6 += v6
                _ = fi9
            }
        }
        _ = fi8
    }
    v6 += (v6 & 42)
    fmt.Println(v1, v7, v1)
    delete(v4, 2)
    v1 = v1
    v6 += v6
    if (4294967295 != 4294967295) {
        v10 := GenS1{F0: 0}
        v11 := GenS2{F0: v6, F1: "\n\t"}
        _ = v10
        _ = v11
    } else {
        if (100 <= -1) {
            var v12 string = "hello"
            _ = v12
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
