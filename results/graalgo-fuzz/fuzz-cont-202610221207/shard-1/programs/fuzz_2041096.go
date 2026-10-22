package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }

func main() {
    var v1 int16 = -44
    v2 := make(map[uint32]int)
    var v3 uint32 = 0
    var v4 bool = true
    var v5 uint32 = (v3 + 2147483647)
    var v6 int16 = (v1 << ((2) & 63))
    if v4 {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{-9223372036854775808}
            if true {
                v9 := GenS1{F0: "a", F1: "foo bar", F2: 496}
                _ = v9
            } else {
                v2[1] = fi7
            }
            _ = fi7
            _ = v8
        }
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS2{F0: -101, F1: 1, F2: fi10}
            v12 := []int{int(v6)}
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            v4 = v4
            var v14 uint = 42
            _ = fi13
            _ = v14
        }
        v1 = (v1 ^ (^(v6)))
    }
    var v15 float64 = 1.5
    v15 = (3.14159 * (((((1e10 / 0.1) - (v15 - v15)) * 1.5) * (1e10 + float64(1e10))) - (v15 + float64(3.14159))))
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 uint16 = 0
        _ = fi16
        _ = v17
    }
    var v18 float32 = (0.0 + 0.0)
    v6 += 1000
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
