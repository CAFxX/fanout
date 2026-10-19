package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16; F1 uint8; F2 int }

func main() {
    var v1 int = -1
    v2 := make(map[string]int)
    var v3 int8 = 19
    v4 := make(map[string]int)
    var v5 string = "a"
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(v3, v1, v1)
        var v7 uint32 = 0
        _ = fi6
        _ = v7
    }
    _ = len(v4)
    v8 := []int{(-(v1)), v1, (v1 / ((-853) | 1))}
    switch v1 {
    case 720:
        v8 = append(v8, 32767)
    case 9223372036854775807:
        var v9 int64 = 9223372036854775807
        _ = v9
    }
    if (!(!(-927 >= -713))) {
        v5 = v5
        v1 += v1
        fmt.Println(v1, v3, v3)
    }
    switch v5 {
    case "\n\t":
        v10 := GenS1{F0: 1, F1: ("foo bar" + v5)}
        _ = v10
    case "a":
        v3 = (v3 - 127)
    case "hello":
        if true {
            v11 := GenS1{F0: 255, F1: (v5 + "")}
            fmt.Println(v1, v3)
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS2{F0: 0, F1: 127, F2: (fi12 * (v1 * 32767))}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            v15 := GenS1{F0: 65535, F1: v5}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS2{F0: 32767, F1: 0, F2: v1}
    v3 -= (v3 ^ v3)
    _ = len(v4)
    v3 -= v3
    var v17 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
