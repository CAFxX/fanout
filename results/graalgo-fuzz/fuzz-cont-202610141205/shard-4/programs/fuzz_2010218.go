package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int32; F1 int32; F2 uint32 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 int = -724
    v5 := GenS1{F0: (425 == 65535), F1: 72}
    var v6 uint32 = 42
    var v7 float32 = (float32(-2.25) / 3.14159)
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (!false), F1: (-(v4))}
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 float64 = (-2.25 * (-2.25 - (1e10 * 1e10)))
            v7 += ((((1.5 / v7) / ((((3.14159 + -2.25) - v7) / 1.5) - v7)) + v7) - -2.25)
            fmt.Println(fi8, v11, v7)
            var v12 string = "\n\t"
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    v7 = v7
    v13 := v1[v4]
    fmt.Println(v13, v6, v4)
    v13 = v13
    v14 := v1[v4]
    switch 2147483647 {
    case 4294967295:
        v15 := GenS1{F0: (65535 <= 42), F1: 520}
        _ = v15
    case 739:
        v6 -= (^(v6))
    case 866:
        v16 := []int{2147483647}
        _ = v16
    }
    var v17 float32 = float32(0.1)
    v18 := v1[(^(v4))]
    if (377 == 65535) {
        var v19 int16 = 1
        v20 := GenS2{F0: 42, F1: v19}
        switch -2147483648 {
        case 615:
            v21 := []int{int(v6), v4}
            _ = v21
        case 9223372036854775807:
            if (!(v6 == v6)) {
                v22 := GenS3{F0: 2147483647, F1: -408, F2: (v6 ^ v6)}
                v23 := GenS1{F0: false, F1: v4}
                _ = v22
                _ = v23
            } else {
                delete(v1, (^(v4)))
                v24 := GenS2{F0: 609, F1: (^(v19))}
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(v17, v7)
                _ = fi25
            }
        }
        _ = v19
        _ = v20
    }
    v26 := []int{}
    v27 := GenS4{F0: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
