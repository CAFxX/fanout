package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int = 128
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    fmt.Println(v1)
    switch ("a" + "foo bar") {
    case "\n\t":
        for fi4 := 0; fi4 < 9; fi4++ {
            for fi5 := 0; fi5 < 5; fi5++ {
                var v6 uint16 = 251
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
    case "foo bar":
        _ = len(v2)
    default:
        v7 := GenS1{F0: 255, F1: -296}
        _ = v7
    }
    v8 := []int{v1, v1, v1, 32767}
    if ((-126 > 42) || false) {
        fmt.Println(v1)
    } else {
        var v9 int = -9223372036854775808
        if len(v8) > 0 {
            v8[0] = v1
        }
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS2{F0: 127}
        _ = fi10
        _ = v11
    }
    fmt.Println(v1)
    if true {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 float32 = float32(0.1)
            switch "\n\t" {
            case "hello":
                v1 = (v1 << ((255) & 63))
            case "\n\t":
                v13 += ((3.14159 - (0.0 * (0.0 / v13))) * 0.1)
            case "":
                var v14 uint8 = 66
                _ = v14
            default:
                var v15 uint8 = 255
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        var v16 bool = (!(!(!(1.5 != ((1.5 / 1.0) / (3.14159 - (1e10 * 0.0)))))))
        v1 += v1
        _ = v16
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
