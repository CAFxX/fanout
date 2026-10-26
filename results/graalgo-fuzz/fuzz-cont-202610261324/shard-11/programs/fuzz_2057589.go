package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
func genFn0(p0 bool, p1 string) int {
    var v6 uint8 = 128
    fmt.Println(p0, p1, p1)
    var v7 int32 = 1
    if p0 {
        if p0 {
            if (-64 <= 127) {
                var v8 uint = 2
                var v9 int32 = v7
                if false {
                    var v10 float64 = -2.25
                    _ = v10
                } else {
                    v9 = (v9 | v9)
                    p0 = (65535 > 65535)
                    var v11 int8 = 75
                    _ = v11
                }
                _ = v8
                _ = v9
            }
            if p0 {
                var v12 int8 = 0
                _ = v12
            } else {
                for fi13 := 0; fi13 < 8; fi13++ {
                    for fi14 := 0; fi14 < 13; fi14++ {
                        v6 = (v6 / ((1) | 1))
                        _ = fi14
                    }
                    var v15 int16 = 960
                    _ = fi13
                    _ = v15
                }
                p1 = "a"
                var v16 int = -1
                _ = v16
            }
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 128
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[string]int)
    var v5 uint32 = 2147483648
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(v1, fi17)
        v1 = 784
        _ = fi17
    }
    fmt.Println(v1, v1)
    var v18 float32 = 0.0
    switch "a" {
    case "\n\t":
        v18 = v18
    case "a":
        var v19 int64 = 9223372036854775807
        _ = v19
    case "foo bar":
        v18 = v18
    }
    v20 := GenS1{F0: -348, F1: v1}
    var v21 float64 = (((-2.25 - 3.14159) + 0.0) / ((1e10 + (0.1 * float64(v18))) * 0.0))
    var v22 uint32 = (v5 >> ((815) & 63))
    var v23 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
