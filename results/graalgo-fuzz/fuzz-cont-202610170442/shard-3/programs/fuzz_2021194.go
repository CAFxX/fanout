package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    v1 := []string{("" + "foo bar"), ""}
    var v2 int8 = 0
    v3 := make(map[int]bool)
    var v4 float32 = ((0.1 + ((((float32(-2.25) * ((3.14159 - 0.1) - 0.1)) + 1e10) + (3.14159 * -2.25)) - 0.1)) * 0.0)
    v5 := make(map[string]string)
    v6 := make(map[uint32]int)
    v4 += v4
    var v7 float64 = (float64((1e10 * ((float32((((1.5 / 1.0) + 0.1) - 0.1)) + 0.1) + v4))) + (float64(0.1) / 1.5))
    v7 = v7
    if (127 < 1) {
        _ = len(v1)
        v4 -= (1e10 - v4)
    } else {
        fmt.Println(v4)
        v7 = (float64(v4) / float64(1e10))
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 float64 = (-2.25 / v7)
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(fi8, v7)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v11 int = 9223372036854775807
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = (v2 | (v2 % ((v2) | 1)))
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 int16 = -101
        if false {
            v14 := GenS1{F0: 2}
            fmt.Println(v7)
            switch "hello" {
            case "":
                v13 -= v13
            case "a":
                v13 = (v13 * -311)
            case "hello":
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 uint = 1
                    var v17 uint32 = 882
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = v14
        } else {
            switch 9223372036854775807 {
            case 2147483647:
                fmt.Println(v4, v4, fi12)
            case 32768:
                v18 := GenS2{F0: -1, F1: -753}
                _ = v18
            default:
                _ = len(v3)
            }
            fmt.Println(v11, fi12, v4)
            var v19 int64 = 1
            _ = v19
        }
        _ = fi12
        _ = v13
    }
    var v20 int8 = v2
    v2 -= (v2 ^ v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
