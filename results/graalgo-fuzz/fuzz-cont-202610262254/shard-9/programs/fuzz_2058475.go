package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }

func main() {
    var v1 float64 = (1.5 / (float64(-2.25) + 1e10))
    var v2 uint32 = 0
    v3 := []int{2147483647, -493}
    v4 := []uint32{}
    v4 = append(v4, (v2 / ((771) | 1)))
    var v5 int16 = -50
    v6 := GenS1{F0: 886, F1: -2147483648}
    if true {
        fmt.Println(v1)
    }
    var v7 int = 127
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 = v2
        if len(v4) > 0 {
            v4[0] = 65535
        }
        _ = fi8
    }
    switch v7 {
    case -804:
        switch "a" {
        case "\n\t":
            var v9 int32 = 638
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 20; fi10++ {
                switch ("hello" + "") {
                case "\n\t":
                    if len(v4) > 0 {
                        v4[0] = 206
                    }
                default:
                    v2 = (v2 % ((2) | 1))
                }
                for fi11 := 0; fi11 < 19; fi11++ {
                    v5 = v5
                    _ = fi11
                }
                _ = fi10
            }
        case "foo bar":
            var v12 uint64 = 4294967295
            _ = v12
        }
    case 32767:
        var v13 float32 = 0.0
        _ = v13
    case 32768:
        v14 := GenS1{F0: 4294967295, F1: (v7 | (-(v7)))}
        _ = v14
    }
    v5 = int16(v7)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
