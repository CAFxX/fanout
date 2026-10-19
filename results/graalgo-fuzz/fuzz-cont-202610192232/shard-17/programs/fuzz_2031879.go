package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 uint }

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    v3 := []string{"foo bar", "a"}
    var v4 uint32 = 737
    var v5 uint8 = 255
    var v6 float64 = ((0.1 - 1e10) - (float64(-2.25) - 1.5))
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v6)
        v5 = v5
        fmt.Println(v5)
        _ = fi7
    }
    v4 += 2147483648
    v8 := GenS1{F0: "", F1: 4294967295, F2: ("hello" + "hello")}
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v6, v1)
        v10 := GenS2{F0: 0}
        _ = fi9
        _ = v10
    }
    var v11 int = 32768
    v1 = int32(v4)
    fmt.Println(v4, v6, v5)
    for fi12 := 0; fi12 < 20; fi12++ {
        if (((0 >= int8(v5)) || true) || (((("foo bar" + "hello") + "a") + ("" + "hello")) >= "hello")) {
            v6 = float64(float32(-2.25))
            fmt.Println(fi12, v6, v4)
        }
        _ = fi12
    }
    fmt.Println(v5)
    for fi13 := 0; fi13 < 17; fi13++ {
        v5 += (^(v5))
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int = (fi14 * (v11 * fi13))
            var v16 bool = false
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
