package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 0.0
    if (v3 != (1.5 - 1.5)) {
        for fi4 := 0; fi4 < 8; fi4++ {
            var v5 int16 = 42
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 int16 = v5
                _ = fi6
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    for fi8 := 0; fi8 < 15; fi8++ {
        v3 -= (v3 - 0.1)
        v3 = float64((0.0 + 0.0))
        _ = fi8
    }
    v9 := []int{}
    if (float64(3.14159) <= v3) {
        _ = len(v9)
        var v10 int16 = 42
        v3 += v3
        _ = v10
    } else {
        var v11 uint = 0
        v12 := GenS1{F0: 4294967295, F1: 316}
        _ = v11
        _ = v12
    }
    fmt.Println(v3)
    v13 := GenS1{F0: 780, F1: 42}
    if false {
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(fi14, fi14)
            switch "" {
            case "hello":
                switch "a" {
                case "":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v3 = 3.14159
                        v3 = 0.0
                        _ = fi15
                    }
                case "\n\t":
                    v16 := v1[("a" + "foo bar")]
                    _ = v16
                case "a":
                    var v17 string = "foo bar"
                    _ = v17
                }
            }
            var v18 bool = false
            _ = fi14
            _ = v18
        }
        fmt.Println(v3)
        switch 733 {
        case 1:
            _ = len(v9)
        default:
            v9 = append(v9, -889)
        }
    }
    fmt.Println(v3)
    var v19 int32 = -1
    for fi20 := 0; fi20 < 14; fi20++ {
        fmt.Println(v3, fi20, v19)
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
