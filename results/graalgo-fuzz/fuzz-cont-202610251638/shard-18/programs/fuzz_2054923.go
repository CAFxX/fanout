package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = (float32((-2.25 - 0.1)) - 1e10)
    var v2 uint = 1
    var v3 uint = v2
    var v4 float64 = 1.5
    v5 := []uint32{4294967295, 4294967295}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 337)
    for fi6 := 0; fi6 < 18; fi6++ {
        switch "hello" {
        case "hello":
            fmt.Println(v1, v1)
        case "\n\t":
            v2 -= v2
        default:
            v7 := GenS1{F0: 32767}
            _ = v7
        }
        _ = fi6
    }
    if (false || false) {
        v4 = v4
        v8 := GenS1{F0: -1}
        _ = v8
    } else {
        v4 -= (((float64((((v1 - v1) + 1e10) + v1)) + 0.1) + ((-2.25 * v4) / v4)) / (-2.25 - (float64((float32((v4 * float64(v1))) + 1.5)) - (-2.25 + (v4 * float64((-2.25 + 1.5)))))))
        v4 = v4
        v2 = v2
    }
    v2 -= (v3 - v3)
    var v9 int8 = -26
    var v10 int = 127
    v10 -= 341
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
