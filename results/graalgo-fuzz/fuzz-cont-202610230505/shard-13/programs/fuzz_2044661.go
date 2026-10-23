package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }

func main() {
    var v1 uint32 = 245
    v2 := make(map[string]string)
    var v3 uint16 = 42
    var v4 uint16 = v3
    for fi5 := 0; fi5 < 2; fi5++ {
        v4 -= v3
        v6 := GenS1{F0: 202}
        _ = fi5
        _ = v6
    }
    v1 = (v1 & (v1 & 4294967295))
    fmt.Println(v4)
    v3 -= v3
    v1 -= (v1 >> ((0) & 63))
    switch "foo bar" {
    case "":
        var v7 int8 = 0
        _ = v7
    }
    if (0 != -1) {
        switch 32767 {
        case 537:
            for fi8 := 0; fi8 < 14; fi8++ {
                v9 := []int{fi8, fi8}
                var v10 float32 = 1.5
                _ = fi8
                _ = v10
                _ = v9
            }
        case 42:
            v11 := GenS2{F0: 894, F1: v1}
            _ = v11
        case 247:
            v12 := GenS1{F0: 181}
            _ = v12
        }
        v13 := []int{-671, 204}
        for fi14 := 0; fi14 < 6; fi14++ {
            fmt.Println(v4, v1)
            _ = fi14
        }
        _ = v13
    }
    _ = len(v2)
    v15 := GenS2{F0: 42, F1: v1}
    var v16 string = "a"
    switch -1 {
    case 32768:
        var v17 int8 = -1
        _ = v17
    case 128:
        v18 := []int{}
        _ = v18
    default:
        v19 := GenS3{F0: 127, F1: -376, F2: -1}
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
