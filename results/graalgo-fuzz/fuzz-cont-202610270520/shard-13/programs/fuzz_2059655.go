package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 int }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]int)
    var v3 bool = (0 < 0)
    var v4 uint16 = 32767
    v5 := make(map[int]string)
    var v6 float64 = 0.1
    var v7 int64 = 53
    v8 := GenS1{F0: 1, F1: 122}
    var v9 uint16 = v4
    switch (v1 + v1) {
    case "hello":
        v1 = (v1 + "a")
    case "":
        if v3 {
            if v3 {
                fmt.Println(v3, v4)
                v3 = (-1 >= -732)
            }
        } else {
            switch -382 {
            case 0:
                var v10 int16 = -32768
                _ = v10
            case 2147483647:
                for fi11 := 0; fi11 < 15; fi11++ {
                    var v12 uint8 = 42
                    fmt.Println(v12, v1, v12)
                    _ = fi11
                    _ = v12
                }
            }
        }
    case "a":
        var v13 int32 = 1
        _ = v13
    default:
        for fi14 := 0; fi14 < 3; fi14++ {
            v9 -= (v4 ^ v9)
            _ = fi14
        }
    }
    v7 = -9223372036854775808
    if (1 >= 1) {
        var v15 float64 = v6
        v16 := GenS2{F0: 223, F1: v9, F2: 127}
        _ = v15
        _ = v16
    } else {
        v17 := GenS2{F0: 0, F1: 255, F2: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 uint16 = 2
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
