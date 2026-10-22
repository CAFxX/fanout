package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    var v3 int16 = -32768
    for fi4 := 0; fi4 < 7; fi4++ {
        v3 = 1000
        _ = fi4
    }
    switch 32768 {
    case 32768:
        v3 = v3
    case 128:
        fmt.Println(v3)
    case 542:
        v5 := GenS1{F0: 2147483647, F1: 486}
        _ = v5
    default:
        v6 := GenS1{F0: -1, F1: 77}
        _ = v6
    }
    v3 -= 1000
    v3 -= v3
    fmt.Println(v3)
    fmt.Println(v3)
    v1[292] = false
    switch 32768 {
    case 32768:
        fmt.Println(v3)
    case 976:
        v7 := GenS2{F0: (958 == 345), F1: "foo bar"}
        _ = v7
    case -2147483648:
        for fi8 := 0; fi8 < 9; fi8++ {
            v3 += v3
            fmt.Println(fi8, v3)
            _ = fi8
        }
    }
    if (!(42 == 33)) {
        v3 -= v3
        v1[239] = true
    } else {
        v9 := GenS3{F0: 0, F1: true}
        _ = v9
    }
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
