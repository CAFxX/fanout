package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 string }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    var v3 int32 = 305
    var v4 float64 = float64((1.5 + 3.14159))
    v5 := []string{"\n\t", "foo bar", "\n\t", "hello"}
    for fi6 := 0; fi6 < 6; fi6++ {
        v4 -= (((float64(1e10) - ((v4 - 1.5) / 1.5)) + v4) * v4)
        _ = fi6
    }
    v3 += (v3 >> ((42) & 63))
    v3 = v3
    var v7 int64 = -9223372036854775808
    fmt.Println(v4, v3)
    v7 = 1
    var v8 uint16 = 1
    if ((-2.25 + 0.1) < float32(float64(float32(v4)))) {
        for fi9 := 0; fi9 < 17; fi9++ {
            v4 += 0.1
            fmt.Println(v3, v8)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi9
        }
    } else {
        v4 = 1e10
    }
    if (!(false && ((v8 / ((v8) | 1)) >= (v8 | 1)))) {
        v10 := GenS1{F0: false, F1: 65535, F2: ("" + ("" + "hello"))}
        _ = v10
    }
    v11 := GenS2{F0: 347}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
