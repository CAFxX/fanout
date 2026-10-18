package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 bool = true
    var v2 float32 = (1.5 * (3.14159 / (1.5 - 3.14159)))
    v3 := make(map[int]uint32)
    v4 := []int{}
    _ = len(v3)
    fmt.Println(v2, v2)
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 int64 = 42
        v7 := GenS1{F0: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2, v1)
    _ = len(v4)
    fmt.Println(v2, v1)
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS2{F0: 0}
        v10 := GenS2{F0: 1}
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 int64 = 100
            if false {
                v12 += 9223372036854775807
            }
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if (true || (239 <= 2)) {
        fmt.Println(v1, v2)
        switch "" {
        case "\n\t":
            v2 = 0.0
        }
    } else {
        var v13 int32 = 1
        _ = v13
    }
    _ = len(v3)
    switch "foo bar" {
    case "\n\t":
        v1 = (!(946 < 516))
    case "a":
        v2 -= float32(1e10)
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v1 = v1
    var v14 int32 = 1
    fmt.Println(v2, v14, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
