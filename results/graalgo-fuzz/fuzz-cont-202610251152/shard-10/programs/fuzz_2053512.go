package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int8; F1 string; F2 int }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]bool)
    v3 := []uint32{42}
    var v4 uint = 777
    var v5 float64 = ((((float64(float32(((float64(0.1) - 3.14159) - -2.25))) - (((1.5 - 0.0) / 1e10) / 1.5)) / 1e10) * (float64(3.14159) + (((0.0 / 1.0) + 1.5) - 1.5))) - 0.0)
    v6 := GenS1{F0: 127, F1: 32767, F2: 42}
    if len(v3) > 0 {
        v3[0] = 1
    }
    v7 := GenS1{F0: 160, F1: -32768, F2: 575}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint = (v4 + v4)
        var v10 float64 = float64(0.0)
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 += v1
    v11 := GenS2{F0: -749, F1: "\n\t"}
    fmt.Println(v5, v1)
    v1 = (v1 ^ v1)
    v12 := v2[int(v1)]
    var v13 uint16 = 2
    var v14 uint32 = 2
    v15 := GenS2{F0: 0, F1: "a"}
    v12 = v12
    v16 := GenS3{F0: v1, F1: ((("foo bar" + "hello") + ("hello" + ("hello" + "hello"))) + "hello"), F2: 0}
    switch -2147483648 {
    case 128:
        _ = len(v2)
    default:
        for fi17 := 0; fi17 < 11; fi17++ {
            switch "foo bar" {
            case "foo bar":
                v12 = false
            case "":
                v18 := v2[fi17]
                _ = v18
            case "\n\t":
                v14 += v14
            }
            _ = fi17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
