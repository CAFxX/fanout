package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 int = -930
    v3 := make(map[int]int)
    v4 := []int{}
    var v5 uint8 = 127
    v6 := GenS1{F0: -666}
    var v7 uint64 = 42
    var v8 int64 = -9223372036854775808
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 bool = false
        v11 := GenS1{F0: 1}
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        _ = len(v3)
        _ = len(v1)
        var v13 uint64 = v7
        _ = fi12
        _ = v13
    }
    if (1.5 > 3.14159) {
        if true {
            for fi14 := 0; fi14 < 6; fi14++ {
                switch fi14 {
                case -482:
                    v5 = (v5 >> ((255) & 63))
                }
                _ = fi14
            }
        } else {
            v5 = 95
        }
        for fi15 := 0; fi15 < 6; fi15++ {
            delete(v3, (fi15 | (-(fi15))))
            var v16 float32 = (3.14159 / 1.0)
            _ = fi15
            _ = v16
        }
        var v17 uint = 1
        _ = v17
    } else {
        fmt.Println(v7, v7, v7)
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS2{F0: 488, F1: 498}
            _ = fi18
            _ = v19
        }
        v20 := v3[v2]
        _ = v20
    }
    v8 -= v8
    v21 := GenS1{F0: -968}
    _ = len(v1)
    var v22 uint16 = 81
    v23 := GenS1{F0: -2147483648}
    var v24 uint16 = v22
    var v25 uint32 = 2147483648
    delete(v1, 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
