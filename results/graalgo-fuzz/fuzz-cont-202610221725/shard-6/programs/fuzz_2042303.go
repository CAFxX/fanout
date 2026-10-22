package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    var v1 int8 = -54
    v2 := make(map[string]int)
    var v3 uint16 = 65535
    v4 := make(map[string]bool)
    var v5 int32 = 86
    switch 4294967295 {
    case 32767:
        v1 = (v1 + (v1 / (((-(v1))) | 1)))
    case 0:
        var v6 int64 = 761
        _ = v6
    case 127:
        v3 -= (v3 + (v3 % (((v3 % ((1) | 1))) | 1)))
    }
    v5 = v5
    var v7 uint32 = 2147483648
    v8 := GenS1{F0: 948}
    var v9 int64 = -559
    v10 := GenS2{F0: 128, F1: 2}
    for fi11 := 0; fi11 < 15; fi11++ {
        if false {
            fmt.Println(v9, v7)
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = v3
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                _ = len(v2)
                _ = fi14
            }
        }
        _ = fi11
    }
    switch -9223372036854775808 {
    case -2147483648:
        v1 -= (v1 ^ (v1 % ((42) | 1)))
    default:
        v2[("hello" + "\n\t")] = -611
    }
    var v15 int32 = -2147483648
    var v16 int16 = -491
    var v17 float32 = 3.14159
    v18 := v4[("\n\t" + "hello")]
    switch "foo bar" {
    case "hello":
        if (454 <= 2) {
            v17 = v17
            v18 = true
            var v19 float64 = 0.1
            _ = v19
        }
    case "foo bar":
        var v20 bool = (v18 || (v18 || v18))
        _ = v20
    }
    v21 := GenS3{F0: v1, F1: 1}
    if true {
        var v22 uint = 2
        _ = v22
    }
    if v18 {
        fmt.Println(v5, v1)
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
