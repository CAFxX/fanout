package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint8 = 255
    var v2 uint32 = 2147483647
    v3 := make(map[string]int)
    v4 := []int{-191, 128, 4294967295}
    for fi5 := 0; fi5 < 5; fi5++ {
        v1 += 127
        switch "foo bar" {
        case "hello":
            delete(v3, "")
        case "":
            v6 := GenS1{F0: (v2 & v2)}
            _ = v6
        default:
            var v7 uint32 = 479
            _ = v7
        }
        var v8 uint64 = 0
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := v3[("a" + "a")]
        _ = fi9
        _ = v10
    }
    var v11 float32 = (((0.0 - (-2.25 * 1e10)) + (float32(-2.25) / (1.5 + float32(float64(3.14159))))) - 1.5)
    switch "a" {
    case "a":
        fmt.Println(v11)
    case "foo bar":
        fmt.Println(v1, v1)
    case "":
        var v12 int16 = 520
        _ = v12
    default:
        fmt.Println(v2, v1)
    }
    var v13 int16 = 32767
    var v14 uint = 0
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 int16 = (v13 >> ((uint(v2)) & 63))
        if (float64((0.1 / ((v11 / (1e10 - float32((1e10 + ((3.14159 * float64(v11)) + -2.25))))) + 3.14159))) == 0.0) {
            var v17 int32 = -1
            var v18 uint16 = 195
            _ = v17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
