package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -9223372036854775808
    var v3 int = -700
    v4 := GenS1{F0: 42}
    for fi5 := 0; fi5 < 6; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float32 = 0.0
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    if (18446744073709551615 <= 2) {
        switch 32768 {
        case 279:
            v2 = v2
        default:
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint32 = 0
                _ = fi8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{-600, (v3 % ((v3) | 1))}
            _ = fi10
            _ = v11
        }
        fmt.Println(v3)
    }
    v2 -= v2
    for fi12 := 0; fi12 < 15; fi12++ {
        v2 += v2
        _ = fi12
    }
    var v13 uint32 = 0
    var v14 uint = 766
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint64 = uint64(v13)
        _ = fi15
        _ = v16
    }
    v14 += 13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
