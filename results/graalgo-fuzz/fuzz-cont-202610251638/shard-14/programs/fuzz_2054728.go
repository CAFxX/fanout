package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float32 = ((float32(float64(3.14159)) * float32((0.1 - 1e10))) - -2.25)
    v2 := make(map[uint32]string)
    var v3 uint = 255
    v4 := []uint32{2147483647}
    v5 := make(map[string]string)
    fmt.Println(v1, v1)
    switch 128 {
    case 9223372036854775807:
        var v6 uint16 = 1
        _ = v6
    }
    var v7 float32 = (v1 * 0.1)
    if (-1 >= 32767) {
        if true {
            v8 := GenS1{F0: -1, F1: 346}
            _ = v8
        }
        _ = len(v4)
        v7 += 0.0
    }
    var v9 int = -1
    var v10 bool = ((v9 >> ((v3) & 63)) < (v9 | v9))
    switch "\n\t" {
    case "a":
        v4 = append(v4, 42)
    }
    if v10 {
        fmt.Println(v9, v9)
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v7 = (float32(3.14159) + (float32((((0.0 - 1.5) + 0.1) - ((3.14159 + 0.0) + 1.5))) / 0.1))
            _ = len(v4)
            v9 = v9
            _ = fi11
        }
    }
    v12 := GenS1{F0: -94, F1: 435}
    v3 += 42
    v13 := GenS1{F0: 1000, F1: 1}
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := v5["foo bar"]
        if (((1e10 * 0.0) / ((1e10 + float64(0.0)) - -2.25)) == -2.25) {
            switch "a" {
            case "hello":
                v16 := v2[692]
                _ = v16
            case "\n\t":
                var v17 int16 = -738
                _ = v17
            default:
                v18 := GenS2{F0: 42}
                _ = v18
            }
        } else {
            for fi19 := 0; fi19 < 18; fi19++ {
                v1 = 1e10
                fmt.Println(v15)
                _ = fi19
            }
        }
        _ = fi14
        _ = v15
    }
    v10 = (255 <= (v3 + 2))
    var v20 int8 = -110
    if len(v4) > 0 {
        v4[0] = 266
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
