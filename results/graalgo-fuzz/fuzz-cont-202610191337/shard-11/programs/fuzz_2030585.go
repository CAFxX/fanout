package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16; F1 int64; F2 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 7; fi4++ {
        v2 += (v2 / (((^(v2))) | 1))
        v5 := GenS1{F0: 0, F1: 1, F2: 735}
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint8 = 216
        v8 := []int{-2147483648, -9223372036854775808}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = ("\n\t" + ("hello" + "a"))
    var v10 int16 = 0
    var v11 int64 = 1
    v11 = 42
    v12 := GenS2{F0: v9, F1: v11}
    v13 := GenS1{F0: v2, F1: 762, F2: 1}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 float64 = 3.14159
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 uint = 444
            v18 := GenS3{F0: v10, F1: 42, F2: v11}
            v2 -= (v2 & 127)
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v10)
        _ = fi14
        _ = v15
    }
    v11 = v11
    if true {
        fmt.Println(v9)
    }
    v10 += v10
    fmt.Println(v9, v9)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v9
}
