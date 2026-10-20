package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    var v3 int16 = 32767
    var v4 int8 = 100
    fmt.Println(v2)
    fmt.Println(v4)
    fmt.Println(v2, v2, v3)
    if true {
        v3 = v3
        var v5 int = 1
        var v6 uint = 253
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 1, F1: 65535, F2: (v3 % ((-32768) | 1))}
        v4 += int8(v3)
        _ = fi7
        _ = v8
    }
    switch 45 {
    case -9223372036854775808:
        v3 = v3
    default:
        v9 := []int{366, -634}
        _ = v9
    }
    switch ((v2 + "foo bar") + ("foo bar" + "")) {
    case "\n\t":
        var v10 int64 = -57
        _ = v10
    }
    v3 += (v3 / ((v3) | 1))
    var v11 float64 = 0.0
    var v12 float64 = float64(0.0)
    var v13 uint8 = 1
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v2)
        _ = fi14
    }
    v3 -= (v3 * v3)
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 float32 = 1e10
        _ = fi15
        _ = v16
    }
    var v17 int = 2147483647
    fmt.Println(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
