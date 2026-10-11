package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 uint16; F1 string }
func genFn0() int16 {
    var v5 int64 = -26
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = v5
        var v7 int64 = v5
        _ = fi6
        _ = v7
    }
    _ = v5
    return -641
}

func main() {
    v1 := []uint32{42, 483, 516}
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32((float64(0.1) - (0.0 * ((0.1 / (float64(1.5) + 1.5)) * float64(float32(-2.25))))))
    var v8 bool = false
    v9 := GenS1{F0: 42, F1: "hello", F2: "foo bar"}
    fmt.Println(v4)
    fmt.Println(v4)
    v8 = true
    if ((800 <= 65535) && v8) {
        v4 = v4
    }
    var v10 uint32 = 740
    v11 := GenS2{F0: 32767, F1: "foo bar"}
    fmt.Println(v10, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
