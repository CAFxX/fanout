package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 string; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 0
    v3 := make(map[int]string)
    fmt.Println(v2)
    for fi4 := 0; fi4 < 17; fi4++ {
        v2 += (v2 & 128)
        switch fi4 {
        case -2147483648:
            v2 += 2
        case 127:
            v5 := GenS1{F0: 940, F1: 2147483648}
            _ = v5
        default:
            for fi6 := 0; fi6 < 11; fi6++ {
                v2 += v2
                v2 -= v2
                _ = fi6
            }
        }
        if (((0.0 / 1e10) * 3.14159) <= 3.14159) {
            v7 := []int{fi4, fi4, (fi4 ^ fi4)}
            v8 := GenS1{F0: 255, F1: 330}
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    v2 = (v2 + 2)
    if ((false || (true || false)) || false) {
        delete(v3, 42)
    }
    v2 -= v2
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 15, F1: 2}
        v11 := GenS2{F0: (fi9 | 902)}
        var v12 int16 = -1
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    switch "\n\t" {
    case "\n\t":
        v13 := v1[-1]
        _ = v13
    }
    v14 := GenS3{F0: (v2 & 161), F1: "hello", F2: true}
    var v15 int = 42
    v16 := GenS4{F0: v2}
    v17 := GenS5{F0: 42}
    if true {
        var v18 uint32 = 95
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
