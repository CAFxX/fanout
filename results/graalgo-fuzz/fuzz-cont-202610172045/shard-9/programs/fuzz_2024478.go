package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 int = 2147483647
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    v4 := []int{(v1 << ((4294967295) & 63)), 1}
    fmt.Println(v2)
    switch "\n\t" {
    case "hello":
        v2 = 4294967295
    default:
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 bool = (214 <= 128)
            if (((1e10 - 1.5) + ((1e10 - ((3.14159 / 3.14159) - 3.14159)) / 3.14159)) < float64(1e10)) {
                if v6 {
                    var v7 int16 = 1
                    v3[v1] = v6
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v7
                }
            }
            for fi8 := 0; fi8 < 14; fi8++ {
                v6 = (32767 > 255)
                _ = fi8
            }
            _ = fi5
            _ = v6
        }
    }
    fmt.Println(v2, v1)
    v1 -= 32768
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 float64 = 0.1
        switch "a" {
        case "foo bar":
            for fi11 := 0; fi11 < 3; fi11++ {
                v2 -= v2
                _ = fi11
            }
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            v10 = -2.25
            _ = fi12
        }
        _ = fi9
        _ = v10
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v4 = append(v4, 0)
        v1 += (fi13 * fi13)
        _ = fi13
    }
    _ = len(v4)
    v14 := GenS1{F0: -82, F1: "\n\t"}
    var v15 int8 = -80
    v15 -= (v15 / ((v15) | 1))
    v16 := GenS1{F0: 1000, F1: "a"}
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
