package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -359
    v3 := make(map[uint32]string)
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(v4)
            v7 := GenS1{F0: v4, F1: (fi6 - fi6)}
            var v8 int16 = 1
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    switch ("" + "a") {
    case "hello":
        v9 := v1[748]
        _ = v9
    case "foo bar":
        v2 = 42
    case "":
        fmt.Println(v2, v4)
    default:
        v2 -= v2
    }
    var v10 int8 = 0
    if true {
        var v11 uint8 = v4
        delete(v3, 1)
        v4 = v4
        _ = v11
    }
    v4 -= v4
    var v12 float64 = (0.0 + ((-2.25 / ((3.14159 * 1e10) * float64(0.1))) + (0.0 / ((-2.25 / float64((1e10 + 0.1))) - 1.5))))
    var v13 uint8 = 0
    var v14 bool = (42 >= 677)
    var v15 string = ""
    v16 := GenS2{F0: (v10 ^ (v10 ^ v10)), F1: (v4 + v4)}
    switch -315 {
    case 128:
        v13 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
