package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint, p1 int) uint16 {
    var v4 bool = (170 > 944)
    p0 = (^(p0))
    p1 -= (p1 - p1)
    v5 := GenS1{F0: 127, F1: 0, F2: 70}
    var v6 int8 = -40
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 97
}

func main() {
    var v1 bool = (false && (630 != 838))
    v2 := []uint32{65535, 2}
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    v7 := GenS1{F0: 0, F1: 125, F2: 65535}
    switch "hello" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v8 := GenS1{F0: -59, F1: 127, F2: 42}
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS2{F0: 441, F1: 50}
        _ = fi9
        _ = v10
    }
    var v11 float32 = (-2.25 + 1.5)
    v12 := GenS2{F0: 0, F1: 127}
    fmt.Println(v11, v11)
    v11 = (1e10 - (-2.25 / v11))
    if v1 {
        v1 = (-1 == 127)
        v11 = 1e10
    }
    if (v1 || true) {
        v3["foo bar"] = 42
        for fi13 := 0; fi13 < 15; fi13++ {
            fmt.Println(v11, v1, fi13)
            for fi14 := 0; fi14 < 7; fi14++ {
                v15 := GenS2{F0: -1, F1: 1}
                v11 = v11
                _ = fi14
                _ = v15
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi13
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v1 = true
        v16 := v3["hello"]
        _ = v16
    }
    fmt.Println(v11, v1)
    v17 := GenS1{F0: 127, F1: 205, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
