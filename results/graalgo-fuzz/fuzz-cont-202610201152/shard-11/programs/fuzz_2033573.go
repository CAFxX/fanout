package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 string = ("hello" + "hello")
    var v2 int8 = 67
    v3 := []uint32{}
    var v4 int16 = 89
    v5 := make(map[string]uint32)
    var v6 uint32 = 42
    v4 = 55
    switch "hello" {
    case "a":
        v2 = (v2 * (-(v2)))
    }
    v7 := GenS1{F0: v1}
    v8 := GenS2{F0: 48, F1: 117, F2: v1}
    v4 = (v4 & v4)
    if true {
        var v9 float64 = 3.14159
        _ = v9
    } else {
        var v10 uint = 937
        switch v1 {
        case "hello":
            switch 2147483647 {
            case 245:
                var v11 float64 = (-2.25 / 1e10)
                _ = v11
            }
        case "foo bar":
            v1 = v1
        default:
            var v12 float32 = 3.14159
            _ = v12
        }
        _ = v10
    }
    v3 = append(v3, 2147483648)
    v2 = 1
    v13 := GenS3{F0: uint8(v6)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
