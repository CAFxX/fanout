package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0() int {
    var v4 int64 = -1
    v5 := GenS1{F0: 2, F1: 2}
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    v1 = v1
    var v6 string = v1
    v7 := []int{907, -383}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint16 = 255
            fmt.Println(v1, fi9)
            _ = fi9
            _ = v10
        }
        var v11 float64 = 1e10
        _ = fi8
        _ = v11
    }
    var v12 uint = 18446744073709551615
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 11; fi14++ {
            if (v12 >= 2) {
                var v15 bool = (((42 > -1) && (65535 < 2147483648)) && (((((0.0 / float64(-2.25)) + (((1.5 * 1e10) + 1e10) / ((float64((1e10 / 1.5)) - 3.14159) * (float64(0.1) * (1.5 / 1.0))))) * (((3.14159 - 3.14159) / 1.5) - float64(((1e10 + (((((1e10 - 3.14159) + 0.1) + (float32(1.5) + (0.1 / (-2.25 + 3.14159)))) - 3.14159) / 1.5)) / 3.14159)))) * 0.0) > -2.25))
                _ = len(v7)
                var v16 uint = 1
                _ = v15
                _ = v16
            } else {
                var v17 uint16 = 862
                v7 = append(v7, fi13)
                v18 := GenS1{F0: 2, F1: 4294967295}
                v12 -= (v12 % (((^(v12))) | 1))
                _ = v17
                _ = v18
            }
            _ = fi14
        }
        var v19 int8 = 81
        var v20 int8 = 123
        _ = fi13
        _ = v19
        _ = v20
    }
    var v21 int16 = -32768
    var v22 int16 = 791
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v6
    _ = v7
}
