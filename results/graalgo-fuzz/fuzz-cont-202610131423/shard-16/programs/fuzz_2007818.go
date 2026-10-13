package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16) uint64 {
    v6 := []int{-2147483648, -1, 0}
    p0 = 255
    var v7 uint16 = (p0 / ((p0) | 1))
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]uint32)
    var v3 float64 = float64(float32((((v1 - (float64((3.14159 * float32((((v1 - 0.1) - 3.14159) * (v1 * 1e10))))) - float64(0.0))) - (0.0 - (v1 + (0.0 * 3.14159)))) * v1)))
    v4 := []string{"hello", "foo bar"}
    var v5 uint32 = 861
    var v8 float64 = (v3 - v3)
    v5 += v5
    switch "a" {
    case "foo bar":
        v9 := GenS1{F0: false, F1: 127}
        _ = v9
    case "\n\t":
        v1 = v1
    case "hello":
        if false {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 float32 = ((0.1 / (0.1 - (float32(3.14159) / (1.5 * -2.25)))) / -2.25)
                v12 := GenS1{F0: false, F1: 0}
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 string = (("a" + "hello") + ("hello" + ""))
            _ = v13
        } else {
            v1 = 1e10
        }
    default:
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 uint16 = 1
            fmt.Println(v3, v8, fi14)
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v5)
    v3 = (1e10 - v8)
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            v4 = append(v4, "hello")
            v3 = v3
            _ = fi16
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
    }
    v2[822] = v5
    fmt.Println(v8)
    v8 -= (v8 * -2.25)
    _ = len(v2)
    var v17 int8 = 0
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v8 = (v8 / (v3 / (v8 / 1.0)))
    v2[-9223372036854775808] = v5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
