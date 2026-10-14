package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
func genFn0(p0 uint32) int64 {
    var v5 int16 = 0
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 2147483648
    var v3 int16 = 467
    v4 := make(map[int]bool)
    var v6 int8 = 83
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint16 = 0
        _ = fi7
        _ = v8
    }
    fmt.Println(v2, v1, v1)
    if true {
        var v9 float64 = 1.5
        var v10 uint16 = 758
        if true {
            v11 := []int{}
            v9 = v9
            _ = v11
        }
        _ = v10
        _ = v9
    }
    switch -1 {
    case 128:
        for fi12 := 0; fi12 < 12; fi12++ {
            v3 = (v3 >> ((65535) & 63))
            _ = fi12
        }
    default:
        v13 := GenS1{F0: -865, F1: 1, F2: 296}
        _ = v13
    }
    fmt.Println(v1, v3, v1)
    fmt.Println(v1, v1)
    var v14 float64 = 1.5
    v14 += v14
    fmt.Println(v3, v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
