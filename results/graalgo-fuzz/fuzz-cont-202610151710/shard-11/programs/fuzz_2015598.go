package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 1
    v3 := make(map[string]string)
    var v4 int16 = -1
    var v5 float32 = (-2.25 + ((0.1 + (1.5 - (((0.1 + 0.1) * ((3.14159 - ((-2.25 / ((1.5 + 1.5) + 3.14159)) * 3.14159)) / (((0.0 + ((1.5 + 0.1) + 0.0)) - 1.5) - (1.5 - (1.5 * 1.5))))) * (0.1 / 1.0)))) / ((-2.25 / 1.5) - ((0.0 / float32((1.5 + 3.14159))) + 3.14159))))
    var v6 int16 = 42
    fmt.Println(v2)
    v7 := GenS1{F0: 17, F1: (v4 << ((42) & 63))}
    v5 -= v5
    switch "a" {
    case "\n\t":
        if ((true && true) || (42 != 2)) {
            if true {
                v4 = 32767
            } else {
                for fi8 := 0; fi8 < 11; fi8++ {
                    for fi9 := 0; fi9 < 8; fi9++ {
                        v5 = 1.5
                        var v10 int16 = v6
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
                var v11 float64 = (((3.14159 * (((1e10 + 3.14159) + (1e10 / (0.1 * (0.1 / (-2.25 / float64(0.1)))))) - 3.14159)) + (-2.25 + ((1.5 / 3.14159) - 1.5))) - 1e10)
                v12 := v3["foo bar"]
                _ = v11
                _ = v12
            }
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: 2, F1: v6}
                fmt.Println(v4)
                v2 += v2
                _ = fi13
                _ = v14
            }
        }
    case "a":
        if true {
            if (v2 < v2) {
                v4 -= (v4 | -1)
                fmt.Println(v5, v5, v4)
            }
            var v15 uint = 959
            v2 = (v2 / (((^(v2))) | 1))
            _ = v15
        }
    }
    v16 := []int{-988, -763}
    var v17 float64 = (3.14159 / 3.14159)
    var v18 float32 = v5
    if (true || false) {
        delete(v3, "foo bar")
    }
    v18 += (((-2.25 * 0.0) * (v5 / float32(1e10))) / (-2.25 * (v5 - (v18 + -2.25))))
    fmt.Println(v17, v5, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
