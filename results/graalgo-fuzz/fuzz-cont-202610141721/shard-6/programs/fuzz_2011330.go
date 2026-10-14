package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = (((((float64((-2.25 * ((0.0 + ((1.5 + 0.0) / 1e10)) / 1.5))) + float64(0.0)) + 3.14159) / 0.1) - -2.25) * 1.5)
    var v3 uint8 = 2
    var v4 bool = false
    var v5 int = 0
    v6 := make(map[int]uint32)
    v3 -= (v3 >> ((1) & 63))
    for fi7 := 0; fi7 < 7; fi7++ {
        switch "a" {
        case "a":
            v2 = 0.0
        }
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: -668}
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 9223372036854775807}
        _ = fi7
        _ = v10
    }
    v4 = v4
    v11 := GenS2{F0: 83, F1: 0, F2: -1}
    var v12 uint16 = 1
    var v13 uint64 = 0
    var v14 string = ""
    var v15 uint32 = 0
    fmt.Println(v13, v14, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
