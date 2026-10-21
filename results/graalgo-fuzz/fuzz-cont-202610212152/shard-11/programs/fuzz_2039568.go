package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = ((0.0 / 0.1) - ((1e10 * 0.0) + 3.14159))
    var v3 uint16 = 255
    v4 := []string{"", "", ""}
    v5 := make(map[int]uint32)
    v2 += (v2 - 0.1)
    v2 = 3.14159
    v3 += 65535
    if false {
        v2 = float32((((-2.25 + (0.0 / 3.14159)) / -2.25) - (((0.1 / ((0.0 / (3.14159 * (3.14159 / (((float64(v2) * 3.14159) / 1.5) * -2.25)))) - 0.1)) / (-2.25 / ((((1.5 - (0.1 / float64((v2 * (v2 - 0.0))))) / 1.0) - float64((0.0 + (((1.5 - (float32((3.14159 + -2.25)) - v2)) * ((v2 / 1e10) + v2)) - v2)))) - (0.0 + (1.5 + 3.14159))))) + (-2.25 * float64(0.1)))))
        var v6 uint8 = 255
        _ = v6
    } else {
        var v7 float64 = ((3.14159 + 1e10) - ((float64(3.14159) + -2.25) / -2.25))
        v3 += 26
        for fi8 := 0; fi8 < 16; fi8++ {
            v3 = v3
            fmt.Println(v2)
            var v9 uint32 = 0
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v3 = 32767
    v3 -= (^(v3))
    _ = len(v5)
    v10 := GenS1{F0: (42 < -128), F1: 128}
    v2 = (v2 + (0.1 * (v2 + v2)))
    v11 := GenS1{F0: (!(-32768 >= -32768)), F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
