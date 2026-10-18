package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int = 4294967295
    var v4 uint64 = 42
    v5 := []int{(-(v3)), -891, 32767}
    v6 := GenS1{F0: "hello", F1: -1, F2: 40}
    v3 -= 127
    var v7 float64 = (1.5 + 0.0)
    switch "" {
    case "foo bar":
        _ = len(v2)
    }
    if (-423 > int64(v3)) {
        v3 = 127
        v8 := GenS2{F0: 1}
        v7 = v7
        _ = v8
    }
    v7 = (-2.25 * -2.25)
    fmt.Println(v7, v3, v7)
    fmt.Println(v3, v4)
    for fi9 := 0; fi9 < 3; fi9++ {
        _ = len(v1)
        _ = fi9
    }
    var v10 int = (v3 - 32767)
    var v11 int8 = 100
    _ = len(v2)
    if (!false) {
        var v12 int32 = -2147483648
        var v13 float32 = (0.0 * (1e10 * (1e10 - 1e10)))
        v12 -= (v12 - v12)
        _ = v12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
