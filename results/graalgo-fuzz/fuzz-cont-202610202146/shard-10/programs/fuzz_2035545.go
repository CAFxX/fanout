package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 bool }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    var v1 int64 = 0
    var v2 uint16 = 813
    v3 := make(map[string]uint32)
    var v4 int32 = 42
    var v5 uint16 = v2
    var v6 string = "\n\t"
    v2 = 515
    var v7 int8 = 0
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := []int{}
        _ = fi8
        _ = v9
    }
    switch v6 {
    case "":
        v2 = v2
    }
    if (!((^(v7)) < v7)) {
        v2 = v2
    }
    v10 := []int{}
    v10 = append(v10, -1)
    var v11 uint64 = 509
    if (42 != -1) {
        v2 -= 65535
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        if (false && (true || (-2.25 > 0.0))) {
            _ = len(v3)
            v13 := GenS1{F0: 78, F1: -67, F2: false}
            _ = v13
        } else {
            if (1 < 334) {
                v14 := GenS2{F0: int16(v1), F1: 2}
                var v15 int32 = v4
                var v16 int8 = v7
                _ = v14
                _ = v15
                _ = v16
            } else {
                _ = len(v3)
                v2 = (v5 / ((v5) | 1))
            }
        }
        v17 := GenS1{F0: 128, F1: -1, F2: (-1 < v1)}
        v6 = v6
        _ = fi12
        _ = v17
    }
    var v18 bool = true
    switch 259 {
    case 32767:
        var v19 int = -2147483648
        _ = v19
    }
    var v20 int = -437
    v1 = (v1 % (((v1 - (v1 + v1))) | 1))
    var v21 int64 = (v1 ^ (v1 * v1))
    switch v6 {
    case "":
        v7 += (v7 & v7)
    case "foo bar":
        v11 = 1
    default:
        v4 = (v4 % (((^(v4))) | 1))
    }
    v2 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
