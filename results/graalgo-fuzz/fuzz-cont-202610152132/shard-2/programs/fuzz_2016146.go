package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int64 }

func main() {
    var v1 int16 = -1
    v2 := make(map[int]string)
    var v3 int64 = 1
    var v4 uint8 = 234
    v3 += -9223372036854775808
    v5 := GenS1{F0: 32768}
    switch ("" + "a") {
    case "":
        var v6 float32 = float32(1e10)
        _ = v6
    case "a":
        v7 := GenS1{F0: 9223372036854775807}
        _ = v7
    }
    v4 += 255
    v1 += 0
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 369
        v10 := []int{fi8, fi8, (fi8 | (fi8 / ((1) | 1)))}
        _ = len(v10)
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        switch "a" {
        case "a":
            v4 -= (v4 - v4)
        case "\n\t":
            switch "foo bar" {
            case "":
                v4 += (v4 + v4)
            }
        case "":
            var v12 uint64 = 143
            _ = v12
        }
        _ = fi11
    }
    switch -2147483648 {
    case 721:
        fmt.Println(v4)
    }
    if true {
        if false {
            var v13 uint64 = 0
            var v14 float32 = 0.0
            v15 := GenS1{F0: 32768}
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 121, F1: v3}
        switch 634 {
        case 128:
            var v17 int16 = (^(v1))
            _ = v17
        case -1:
            var v18 float32 = 1.5
            _ = v18
        case 1:
            var v19 float32 = 3.14159
            _ = v19
        default:
            v1 = 1000
        }
        _ = v16
    }
    v3 += v3
    v3 -= -9223372036854775808
    for fi20 := 0; fi20 < 1; fi20++ {
        v3 += (v3 << ((770) & 63))
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
