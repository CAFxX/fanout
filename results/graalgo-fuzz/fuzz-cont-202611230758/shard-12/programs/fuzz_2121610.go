package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 9223372036854775807
    var v3 int64 = 1
    var v4 int8 = 84
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := GenS1{F0: 0, F1: (fi5 / ((v2) | 1))}
        for fi7 := 0; fi7 < 16; fi7++ {
            _ = len(v1)
            v8 := GenS1{F0: 280, F1: 1}
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 781, F1: -159}
        _ = fi5
        _ = v6
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(v2, v3, v2)
        _ = fi10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        if (2 == 32767) {
            switch fi11 {
            case -1:
                _ = len(v1)
            case 2147483647:
                v12 := []int{}
                _ = v12
            }
        } else {
            fmt.Println(v4, v3)
        }
        var v13 int = v2
        v3 = v3
        _ = fi11
        _ = v13
    }
    if (true && (2 <= 610)) {
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    v2 = v2
    v4 = 0
    v15 := GenS1{F0: 333, F1: v2}
    v16 := GenS1{F0: 0, F1: (v2 ^ 0)}
    v17 := GenS2{F0: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
