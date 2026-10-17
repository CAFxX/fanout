package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []int{}
    var v2 uint = 65535
    var v3 uint64 = 42
    var v4 float64 = (1.5 / 1.0)
    var v5 int = -566
    v6 := make(map[uint32]bool)
    if false {
        var v7 uint8 = 199
        _ = v7
    } else {
        v8 := GenS1{F0: 127}
        v3 = 42
        _ = v8
    }
    var v9 uint16 = 255
    fmt.Println(v4, v5)
    delete(v6, 163)
    var v10 float64 = v4
    var v11 bool = (true || (false && ((0.1 * -2.25) > ((1.5 * (0.1 - 0.1)) / ((0.1 - 0.1) - float32(v10))))))
    var v12 bool = ("a" >= "hello")
    switch "a" {
    case "":
        var v13 int32 = -318
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 7; fi14++ {
            _ = len(v6)
            _ = fi14
        }
    case "foo bar":
        fmt.Println(v4, v11, v12)
    default:
        v1 = append(v1, (^(v5)))
    }
    v6[790] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
