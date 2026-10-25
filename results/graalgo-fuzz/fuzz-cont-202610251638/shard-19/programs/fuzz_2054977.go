package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 742
    v3 := []uint32{}
    v4 := []int{-254, 246, 128, 2147483647}
    v5 := make(map[string]uint32)
    v2 = (v2 | v2)
    var v6 uint16 = (v2 & 42)
    if (0 > 1) {
        var v7 int32 = 0
        v2 -= (v2 >> ((1) & 63))
        _ = v7
    }
    fmt.Println(v6)
    v8 := v5[(("\n\t" + "\n\t") + "hello")]
    v8 -= (v8 | v8)
    fmt.Println(v6, v8)
    var v9 uint32 = (v8 * 0)
    v8 = 65535
    var v10 uint8 = 2
    v6 -= (v6 | (v6 / ((v6) | 1)))
    switch ("a" + "") {
    case "a":
        v11 := GenS1{F0: true, F1: -1, F2: 881}
        _ = v11
    case "\n\t":
        if (int64(v2) < -573) {
            var v12 int16 = 42
            _ = v12
        } else {
            v4 = append(v4, -9223372036854775808)
            delete(v5, "a")
        }
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v2 = uint16(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
