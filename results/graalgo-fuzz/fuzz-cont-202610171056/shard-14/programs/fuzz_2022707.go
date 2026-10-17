package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    var v1 int16 = 16
    var v2 int8 = 42
    v3 := make(map[uint32]int)
    var v4 uint64 = 14
    var v5 float64 = 0.1
    var v6 uint8 = 42
    v7 := v3[1]
    switch "a" {
    case "":
        var v8 int32 = int32(v7)
        _ = v8
    case "foo bar":
        v2 -= (-(v2))
    }
    v9 := GenS1{F0: 2147483647, F1: 561}
    var v10 string = ("" + ("foo bar" + "foo bar"))
    var v11 bool = (true && true)
    var v12 float64 = v5
    for fi13 := 0; fi13 < 2; fi13++ {
        v11 = (v11 && (!true))
        _ = fi13
    }
    var v14 int64 = 9223372036854775807
    if ((-2.25 < 1e10) && false) {
        delete(v3, 157)
        v15 := v3[42]
        _ = v15
    } else {
        delete(v3, 2147483648)
        var v16 int = 9223372036854775807
        v1 += 1000
        _ = v16
    }
    if ((v6 / ((v6) | 1)) < v6) {
        var v17 uint64 = (^(v4))
        var v18 float64 = float64(1.5)
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        v12 = (v12 + (v5 + ((v5 * float64(0.0)) + 1e10)))
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
