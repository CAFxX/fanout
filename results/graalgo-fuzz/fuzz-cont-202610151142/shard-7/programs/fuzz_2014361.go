package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    var v1 int64 = -438
    v2 := make(map[int]bool)
    var v3 float32 = (1.5 - -2.25)
    if true {
        v4 := GenS1{F0: 65535, F1: 0, F2: 1}
        _ = v4
    }
    fmt.Println(v1, v3)
    fmt.Println(v3, v1)
    switch "hello" {
    case "\n\t":
        v5 := GenS1{F0: 628, F1: -1, F2: 65535}
        _ = v5
    default:
        v6 := GenS2{F0: 0, F1: (128 != 2)}
        _ = v6
    }
    v3 = 1.5
    var v7 int32 = -2147483648
    for fi8 := 0; fi8 < 6; fi8++ {
        v3 = (float32(0.1) * 0.1)
        v1 = v1
        _ = fi8
    }
    switch 32767 {
    case -404:
        v9 := GenS2{F0: 18446744073709551615, F1: false}
        _ = v9
    case -9223372036854775808:
        v10 := []int{-958, 32767, -124, -898}
        _ = v10
    case 1:
        v1 = (v1 * v1)
    default:
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint16 = 32767
            v13 := []int{(-(fi11))}
            if true {
                v12 -= 0
                fmt.Println(v7)
                var v14 uint = 18446744073709551615
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
