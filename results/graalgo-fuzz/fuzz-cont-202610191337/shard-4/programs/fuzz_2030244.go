package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }

func main() {
    var v1 int16 = 42
    v2 := []int{0, 2147483647}
    var v3 float64 = (1.5 * 3.14159)
    v4 := make(map[uint32]uint32)
    v5 := GenS1{F0: -110, F1: 571, F2: 2}
    v6 := GenS2{F0: (v1 >> ((255) & 63))}
    fmt.Println(v1, v1)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = ("foo bar" + "foo bar")
        v9 := GenS2{F0: 0}
        var v10 uint8 = 2
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 string = "foo bar"
    fmt.Println(v11, v3)
    fmt.Println(v3, v3)
    fmt.Println(v1, v3, v11)
    switch 42 {
    case 968:
        var v12 int32 = 713
        _ = v12
    case 32767:
        for fi13 := 0; fi13 < 17; fi13++ {
            v3 += (0.1 * v3)
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS3{F0: 203, F1: 975, F2: 42}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    case 127:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
