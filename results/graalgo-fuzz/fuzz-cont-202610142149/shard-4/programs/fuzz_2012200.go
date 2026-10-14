package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 string; F2 uint }

func main() {
    v1 := []string{}
    var v2 float64 = 0.0
    var v3 float32 = 3.14159
    v4 := []uint32{659, 2}
    v5 := GenS1{F0: 1}
    switch 211 {
    case 9223372036854775807:
        var v6 int16 = 820
        _ = v6
    }
    switch 612 {
    case -9223372036854775808:
        v1 = append(v1, "a")
    case 32768:
        if false {
            fmt.Println(v2)
            v2 += v2
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 float32 = 0.1
                switch "hello" {
                case "hello":
                    switch 115 {
                    case 42:
                        fmt.Println(v3, v3)
                    }
                case "\n\t":
                    var v9 uint64 = 18446744073709551615
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        }
    case 1:
        fmt.Println(v3, v2)
    }
    v2 -= v2
    if (42 >= -54) {
        var v10 uint16 = 1
        _ = v10
    }
    v11 := GenS2{F0: 1, F1: "a", F2: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += ((((-2.25 - -2.25) + 0.1) - 0.1) - (v3 * v3))
    v3 = (v3 * 0.1)
    fmt.Println(v3, v2)
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 int64 = -9223372036854775808
        _ = fi12
        _ = v13
    }
    var v14 string = "\n\t"
    v2 = (v2 / float64(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
