package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{2147483648, 1, 1, 2}
    v3 := []string{"hello", "foo bar", "foo bar", (("\n\t" + ("foo bar" + (("\n\t" + "foo bar") + "hello"))) + "foo bar")}
    var v4 uint16 = 42
    v5 := make(map[string]bool)
    v6 := []string{""}
    for fi7 := 0; fi7 < 7; fi7++ {
        v1 = (float32((-2.25 - float64((1.5 * ((v1 / 1.5) / float32(((float64(1.5) - 0.0) + 1.5))))))) - v1)
        _ = fi7
    }
    v2 = append(v2, 2147483647)
    for fi8 := 0; fi8 < 5; fi8++ {
        if (65535 >= (v4 ^ 42)) {
            v9 := GenS1{F0: 101, F1: -1, F2: 18446744073709551615}
            _ = v9
        } else {
            v10 := GenS2{F0: 128, F1: 1}
            _ = v10
        }
        v6 = append(v6, "foo bar")
        _ = fi8
    }
    v2 = append(v2, 987)
    for fi11 := 0; fi11 < 7; fi11++ {
        fmt.Println(fi11, fi11, v1)
        for fi12 := 0; fi12 < 20; fi12++ {
            v13 := GenS1{F0: 42, F1: -390, F2: 4294967295}
            _ = fi12
            _ = v13
        }
        v1 -= 0.0
        _ = fi11
    }
    var v14 uint32 = 2
    fmt.Println(v14, v4)
    v15 := GenS1{F0: 0, F1: int16(v4), F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
