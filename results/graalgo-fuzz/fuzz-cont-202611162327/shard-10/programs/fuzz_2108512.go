package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 0
    var v3 uint = 42
    v4 := []uint32{uint32(v3), 388, 2147483648, 0}
    var v5 float32 = (0.0 / (1.5 - float32((3.14159 - ((1e10 * 3.14159) * (float64(-2.25) / 3.14159))))))
    v6 := make(map[string]int)
    var v7 int8 = 100
    for fi8 := 0; fi8 < 10; fi8++ {
        switch "a" {
        case "a":
            v5 = v5
        case "":
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v3, v3)
                _ = fi9
            }
        default:
            var v10 int16 = 0
            _ = v10
        }
        v11 := v6[((("hello" + ("" + "")) + "hello") + "foo bar")]
        _ = fi8
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 2147483647, F2: 0}
    v7 = int8(v3)
    if (true && (0 >= 1)) {
        fmt.Println(v1, v3)
        if (602 >= -1) {
            _ = len(v6)
        }
        var v13 float32 = (v5 / v5)
        _ = v13
    }
    var v14 uint32 = 2147483647
    _ = len(v6)
    var v15 uint32 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
