package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 590
    var v3 uint16 = (v2 + 0)
    var v4 string = "a"
    var v5 float32 = float32(1e10)
    var v6 string = v4
    v5 -= 3.14159
    v7 := GenS1{F0: 127, F1: v3, F2: 128}
    v8 := GenS2{F0: 2147483647}
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v9, v2, v3)
        _ = fi10
    }
    if (0 < 35) {
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 int8 = 100
            var v13 int32 = int32(v3)
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(v4)
    } else {
        fmt.Println(v2)
        switch 497 {
        case 32768:
            switch "hello" {
            case "a":
                v6 = v4
            case "foo bar":
                var v14 float64 = float64(-2.25)
                _ = v14
            case "\n\t":
                v15 := v1[int(v3)]
                _ = v15
            default:
                switch 0 {
                case 32767:
                    v5 = (0.0 - float32(((3.14159 * (v9 / float64(1.5))) * v9)))
                case 32768:
                    fmt.Println(v4, v9, v9)
                case 0:
                    v5 = v5
                }
            }
        }
    }
    v4 = v6
    if true {
        v9 -= (float64((1.5 / v5)) / 1e10)
        v16 := GenS1{F0: -9223372036854775808, F1: v2, F2: 127}
        _ = v16
    }
    if true {
        v17 := []int{}
        if true {
            _ = len(v1)
        }
        _ = v17
    } else {
        v4 = v6
    }
    v18 := GenS3{F0: -32768}
    v19 := []int{}
    v5 -= ((((float32(((((v9 / (v9 + 1e10)) - float64(v5)) + (0.1 * 0.1)) / v9)) / -2.25) * ((((0.1 / 3.14159) - ((v5 + v5) + 1.5)) * v5) - 1e10)) / (((1e10 * float32(v9)) - v5) / 1e10)) * v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
