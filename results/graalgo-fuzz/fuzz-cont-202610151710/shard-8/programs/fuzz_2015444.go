package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int16 = 0
    var v2 float64 = 3.14159
    v3 := make(map[int]string)
    if false {
        var v4 bool = (!(false || (v1 < (v1 * v1))))
        var v5 uint32 = 2
        _ = v4
        _ = v5
    }
    switch "foo bar" {
    case "foo bar":
        var v6 int32 = 2147483647
        _ = v6
    case "hello":
        v2 += ((0.0 / 1.5) * v2)
    default:
        v7 := GenS1{F0: 42, F1: 0, F2: 0}
        _ = v7
    }
    fmt.Println(v2, v2)
    var v8 string = "foo bar"
    if true {
        v9 := GenS1{F0: -547, F1: 451, F2: 1}
        v2 += (0.0 * 1.5)
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 int16 = 42
            for fi12 := 0; fi12 < 11; fi12++ {
                v2 = float64((1.5 / (1e10 - ((((0.0 * 1e10) * 1.5) + 1.5) - (-2.25 + (1e10 / ((float32((v2 * (v2 * v2))) * float32(((3.14159 / ((((v2 - 0.0) - v2) / (v2 * 1e10)) * v2)) / (v2 - ((-2.25 / v2) + (v2 - v2)))))) * 1.5)))))))
                switch "foo bar" {
                case "foo bar":
                    v13 := GenS2{F0: 100}
                    _ = v13
                default:
                    var v14 string = v8
                    _ = v14
                }
                _ = fi12
            }
            _ = len(v3)
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        fmt.Println(v8, v8, v1)
    }
    switch -1 {
    case 4294967295:
        var v15 uint64 = 9
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS2{F0: 100}
        _ = fi16
        _ = v17
    }
    v3[4294967295] = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
