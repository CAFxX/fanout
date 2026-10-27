package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    var v1 uint16 = 32767
    v2 := []string{"\n\t"}
    var v3 uint32 = 4294967295
    v3 += (v3 / ((uint32(v1)) | 1))
    v1 -= 999
    v4 := GenS1{F0: 100}
    v5 := GenS1{F0: -108}
    v1 = (v1 + v1)
    v6 := GenS2{F0: v3, F1: 1}
    var v7 uint16 = v1
    switch 42 {
    case 0:
        var v8 bool = (1.5 <= 0.1)
        _ = v8
    default:
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(fi9)
            _ = fi9
        }
    }
    fmt.Println(v3)
    var v10 float64 = float64(0.0)
    switch ("foo bar" + (("foo bar" + "a") + "a")) {
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 uint64 = 1
            var v13 uint16 = (^(v1))
            _ = fi11
            _ = v12
            _ = v13
        }
    case "\n\t":
        switch 128 {
        case 240:
            v2 = append(v2, "a")
        case -1:
            v1 = 255
        default:
            var v14 uint8 = 1
            _ = v14
        }
    case "":
        v7 = (^(v7))
    default:
        fmt.Println(v7, v3)
    }
    fmt.Println(v1, v7)
    fmt.Println(v1, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
