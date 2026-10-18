package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int = 1
    var v2 float64 = 0.1
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 float32 = 0.0
    var v6 int16 = -32768
    v2 = float64(((float32((v2 - (1.5 / -2.25))) - (v5 + -2.25)) + v5))
    var v7 float64 = v2
    switch "a" {
    case "a":
        if true {
            var v8 uint16 = 176
            _ = v8
        } else {
            for fi9 := 0; fi9 < 3; fi9++ {
                switch "foo bar" {
                case "a":
                    delete(v4, 307)
                case "\n\t":
                    v1 += fi9
                }
                v5 = -2.25
                _ = fi9
            }
        }
    default:
        var v10 int = (v1 + (v1 & (v1 / ((v1) | 1))))
        _ = v10
    }
    var v11 int64 = 42
    fmt.Println(v1, v6)
    v12 := GenS1{F0: (v1 % (((v1 & (-(v1)))) | 1))}
    var v13 bool = (!true)
    v11 = v11
    v11 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
