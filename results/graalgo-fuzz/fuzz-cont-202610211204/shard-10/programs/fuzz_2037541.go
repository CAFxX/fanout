package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint8 = 42
    var v2 int8 = -128
    v3 := make(map[string]string)
    var v4 uint8 = (^(v1))
    var v5 uint8 = 128
    fmt.Println(v1)
    switch "a" {
    case "foo bar":
        _ = len(v3)
    case "hello":
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(v1, v2, v5)
            _ = fi6
        }
    case "a":
        var v7 float64 = 0.0
        _ = v7
    default:
        var v8 float32 = 0.0
        _ = v8
    }
    v9 := GenS1{F0: 42}
    var v10 int16 = 569
    if (!false) {
        if true {
            switch 2147483647 {
            case 195:
                fmt.Println(v10, v4, v5)
            }
        } else {
            var v11 float32 = (0.0 - float32(-2.25))
            _ = v11
        }
        var v12 uint32 = 925
        _ = v12
    }
    switch 42 {
    case 42:
        var v13 bool = ((0.1 / (1.5 / 0.1)) >= (0.1 - 0.1))
        _ = v13
    case 4294967295:
        v3[("\n\t" + "hello")] = ("hello" + "")
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := []int{}
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
