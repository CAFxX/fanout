package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }

func main() {
    var v1 uint = 473
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    v4 := []string{(("foo bar" + ("hello" + "")) + ""), "foo bar", "a"}
    _ = len(v4)
    fmt.Println(v1)
    switch "\n\t" {
    case "foo bar":
        v1 = (v1 / ((v1) | 1))
    case "\n\t":
        v1 -= (v1 + v1)
    case "hello":
        var v5 float64 = 0.0
        _ = v5
    default:
        v1 = v1
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int32 = -378
        v7 = (v7 << ((v1) & 63))
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        switch fi8 {
        case 128:
            for fi9 := 0; fi9 < 4; fi9++ {
                v1 = 18446744073709551615
                _ = fi9
            }
        }
        _ = fi8
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := v3[""]
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: "a", F1: -1, F2: "foo bar"}
    v13 := GenS2{F0: 2, F1: uint8(v1), F2: (112 == 128)}
    v1 -= (v1 + v1)
    fmt.Println(v1)
    v1 += 872
    var v14 float32 = (-2.25 + (0.0 + ((3.14159 / float32((1.5 / (0.0 - 3.14159)))) * (-2.25 - (0.1 + (1.5 / ((-2.25 / (((3.14159 + -2.25) / 1.5) * 1e10)) - (-2.25 + (float32(-2.25) / 1e10)))))))))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
