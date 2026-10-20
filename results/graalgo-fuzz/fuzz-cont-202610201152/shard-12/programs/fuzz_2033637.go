package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{128, 32768, -979, 119}
    v3 := []int{128, -176, 2147483647}
    v4 := GenS1{F0: "\n\t", F1: 42, F2: -44}
    if false {
        if (!(!false)) {
            var v5 uint32 = 4294967295
            _ = v5
        } else {
            var v6 float32 = 0.0
            v1[2147483647] = "foo bar"
            for fi7 := 0; fi7 < 7; fi7++ {
                v8 := GenS1{F0: "a", F1: 128, F2: 100}
                _ = fi7
                _ = v8
            }
            _ = v6
        }
    }
    var v9 int16 = -32768
    for fi10 := 0; fi10 < 17; fi10++ {
        var v11 int64 = -1
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int = (^(fi10))
            v14 := GenS2{F0: 2147483647, F1: "foo bar"}
            if ("foo bar" <= "foo bar") {
                _ = len(v1)
                v3 = append(v3, 1)
                fmt.Println(v9)
            }
            var v15 uint8 = 251
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v1[1] = "foo bar"
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 float64 = float64(float32(1.5))
        v17 += 3.14159
        _ = fi16
        _ = v17
    }
    v9 -= v9
    switch 2147483647 {
    case -574:
        v9 -= (v9 % ((v9) | 1))
    case -1:
        var v18 uint64 = 76
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
