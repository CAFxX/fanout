package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int16 }

func main() {
    var v1 string = ("" + "foo bar")
    v2 := make(map[int]uint32)
    var v3 uint64 = 536
    fmt.Println(v3, v1)
    v4 := []int{32768, -2147483648}
    v5 := GenS1{F0: 9223372036854775807, F1: true, F2: 1000}
    v1 = "\n\t"
    _ = len(v4)
    var v6 int32 = -499
    v7 := GenS1{F0: 330, F1: ((0.0 * ((1e10 * 1.5) * float32((0.0 * (float64(0.0) / float64((0.1 / -2.25))))))) == (0.0 + ((1e10 / -2.25) - ((1.5 - 0.0) + 1.5)))), F2: 42}
    var v8 uint32 = 783
    for fi9 := 0; fi9 < 7; fi9++ {
        switch "\n\t" {
        case "a":
            if (-32768 == -306) {
                v6 -= v6
                fmt.Println(v3, v8, v3)
                switch (fi9 ^ fi9) {
                case -995:
                    v6 += 145
                default:
                    v3 = (v3 + v3)
                }
            }
        }
        _ = fi9
    }
    v6 += 1
    _ = len(v2)
    switch ("a" + v1) {
    case "hello":
        var v10 string = v1
        _ = v10
    default:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
