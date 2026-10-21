package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0(p0 int64) int {
    p0 -= -1
    v6 := []int{-396, -9223372036854775808, 42, 3}
    var v7 int16 = -1
    _ = p0
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"hello"}
    v3 := make(map[string]int)
    var v4 uint32 = 1
    var v5 float64 = 0.1
    v5 = (v5 + -2.25)
    for fi8 := 0; fi8 < 3; fi8++ {
        v4 = v4
        _ = fi8
    }
    v5 = ((1.5 / ((v5 + (v5 - float64(float32(((1e10 / float64((-2.25 / (0.0 + (-2.25 + 1e10))))) + v5))))) / 1.0)) + v5)
    v9 := GenS1{F0: 42, F1: 2147483648}
    if true {
        for fi10 := 0; fi10 < 12; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                switch "" {
                case "\n\t":
                    switch "hello" {
                    case "":
                        v5 -= 3.14159
                    case "a":
                        var v12 uint8 = 42
                        _ = v12
                    case "hello":
                        var v13 int16 = -346
                        _ = v13
                    }
                case "hello":
                    v4 -= v4
                case "":
                    v5 -= v5
                default:
                    var v14 int16 = 32767
                    _ = v14
                }
                _ = fi11
            }
            _ = fi10
        }
        v1[""] = v4
    } else {
        fmt.Println(v4, v5)
    }
    v2 = append(v2, "a")
    if true {
        var v15 uint16 = 1
        _ = v15
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            v1["foo bar"] = 65535
            fmt.Println(v5, fi16)
            _ = fi16
        }
    }
    var v17 int64 = 42
    v5 -= v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
