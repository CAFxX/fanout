package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int8 = -1
    v4 := make(map[int]bool)
    var v5 uint16 = 1
    v3 = 1
    var v6 int = 128
    v7 := GenS1{F0: true, F1: 565}
    fmt.Println(v3, v3, v3)
    v6 += 630
    if (1000 > -1) {
        switch v6 {
        case 128:
            v6 -= (v6 << ((926) & 63))
        }
        var v8 float64 = 3.14159
        _ = v8
    } else {
        var v9 uint16 = v5
        _ = v9
    }
    var v10 float64 = ((((((float64(1.5) * 3.14159) / (1e10 / -2.25)) / 1e10) / (3.14159 - (0.0 - (3.14159 - 0.0)))) + 0.0) - float64(-2.25))
    var v11 uint = 65535
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: true, F1: 2}
        var v14 float64 = v10
        _ = fi12
        _ = v13
        _ = v14
    }
    switch "a" {
    case "foo bar":
        v3 = v3
    case "a":
        v6 = 2147483647
    case "hello":
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint64 = 4294967295
            fmt.Println(v11)
            var v17 int16 = -358
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v3 = (v3 + (v3 | (v3 * v3)))
    var v18 int32 = 502
    var v19 bool = (int64(v6) < -9223372036854775808)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
