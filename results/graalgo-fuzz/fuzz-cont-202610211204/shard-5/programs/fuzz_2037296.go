package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 48
    var v4 uint64 = 1
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: true, F1: 65535}
        v8 := []int{fi6, (fi6 << ((65535) & 63))}
        v5 = v5
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 526, F1: (v4 % ((989) | 1))}
    var v10 uint32 = (v3 / ((981) | 1))
    var v11 float32 = -2.25
    v12 := GenS3{F0: -9223372036854775808, F1: -66, F2: 1}
    var v13 int16 = 928
    fmt.Println(v5, v5, v11)
    for fi14 := 0; fi14 < 13; fi14++ {
        v13 = v13
        _ = fi14
    }
    if true {
        var v15 bool = (true || (42 <= 61))
        fmt.Println(v5, v4, v10)
        _ = v15
    }
    if false {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int64 = 0
            _ = fi16
            _ = v17
        }
    }
    var v18 float64 = 0.0
    v11 += v11
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
