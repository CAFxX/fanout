package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 int64; F1 string }
func genFn0() uint8 {
    if (461 < 2) {
        var v5 uint = 751
        switch "foo bar" {
        case "":
            v5 = v5
        case "a":
            var v6 uint = (^(v5))
            _ = v6
        case "hello":
            fmt.Println(v5)
        }
        var v7 uint16 = 0
        _ = v5
        _ = v7
    }
    v8 := []int{9223372036854775807}
    _ = v8
    return 66
}

func main() {
    var v1 int32 = 1
    v2 := []uint32{135, 2147483647, 2, 2147483648}
    var v3 int = -9223372036854775808
    v4 := make(map[string]uint32)
    v9 := GenS1{F0: 94}
    v10 := GenS2{F0: 9223372036854775807, F1: 42, F2: "hello"}
    for fi11 := 0; fi11 < 3; fi11++ {
        v1 -= (^(v1))
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    v1 = 521
    fmt.Println(v3, v3)
    var v13 float64 = float64((1.5 + ((((3.14159 / float32(0.1)) + 1.5) + 1.5) + 1.5)))
    fmt.Println(v3, v13)
    v3 = 370
    v13 -= -2.25
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := GenS3{F0: 400, F1: ("" + "hello")}
        v16 := GenS1{F0: 15}
        v13 = v13
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
