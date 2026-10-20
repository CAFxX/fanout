package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 uint16 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint64 = 2
    var v2 uint = 425
    var v3 int16 = 1
    var v4 float64 = (1e10 / 0.1)
    var v5 float32 = -2.25
    var v6 bool = false
    var v7 float32 = (3.14159 * 1e10)
    fmt.Println(v7, v5, v1)
    if false {
        var v8 float64 = 1.5
        _ = v8
    } else {
        v5 += v5
        v9 := []int{278, 32767}
        _ = v9
    }
    v10 := GenS1{F0: 2, F1: 1, F2: 2147483648}
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            switch fi11 {
            case 2147483647:
                v13 := GenS2{F0: 699}
                _ = v13
            default:
                v1 = 4294967295
            }
            if v6 {
                if (v3 <= (v3 ^ v3)) {
                    switch "foo bar" {
                    case "a":
                        v7 -= (3.14159 / (v5 - v7))
                    }
                }
            } else {
                v14 := GenS1{F0: 127, F1: -383, F2: 1}
                var v15 uint64 = v1
                var v16 bool = ((!(!(!(65535 <= 2147483647)))) || v6)
                _ = v14
                _ = v15
                _ = v16
            }
            v17 := GenS3{F0: (^(v2))}
            if (-1 > 1) {
                if (2147483648 <= 4294967295) {
                    v18 := GenS4{F0: 1000, F1: 262}
                    v19 := GenS5{F0: 65535}
                    fmt.Println(v5)
                    _ = v18
                    _ = v19
                }
                v3 = v3
            }
            _ = fi12
            _ = v17
        }
        _ = fi11
    }
    switch "a" {
    case "\n\t":
        v4 = v4
    default:
        var v20 int64 = -623
        _ = v20
    }
    v5 = 1.5
    var v21 float32 = (3.14159 / v7)
    switch 14 {
    case -1:
        v2 += (v2 / (((^(v2))) | 1))
    }
    var v22 int16 = (v3 * v3)
    v2 -= 4294967295
    if v6 {
        v23 := []int{-482}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
