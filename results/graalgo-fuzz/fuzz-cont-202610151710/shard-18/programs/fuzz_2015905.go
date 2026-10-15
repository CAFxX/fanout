package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := []string{"foo bar", "a", "", "foo bar"}
    var v5 uint32 = 1
    v6 := GenS1{F0: v5, F1: -1, F2: 127}
    if ((1e10 <= 0.0) || true) {
        v5 -= (v5 - 1)
    }
    switch -487 {
    case 32768:
        var v7 int16 = 42
        _ = v7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v5)
    fmt.Println(v5)
    var v8 float32 = (0.0 * ((-2.25 - (0.1 + 0.1)) + 1.5))
    v8 = v8
    if (444 >= 18446744073709551615) {
        delete(v1, "a")
        var v9 uint64 = 0
        _ = v9
    } else {
        var v10 uint = 18446744073709551615
        _ = v10
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
