package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[string]uint32)
    var v3 float64 = (float64(float32((0.0 / float64(3.14159)))) * 3.14159)
    if ((1e10 >= (float32(3.14159) * 0.1)) || false) {
        v3 = v3
        _ = len(v2)
    }
    var v4 int32 = 2147483647
    v1 = v1
    if (((0.1 + 0.0) - float64((float32((v3 + (v3 - v3))) - ((float32((float64(0.0) / (v3 + (v3 / (1.5 + v3))))) / -2.25) * 1.5)))) < (v3 - 0.0)) {
        var v5 uint16 = (^(v1))
        _ = len(v2)
        var v6 int64 = 9223372036854775807
        _ = v5
        _ = v6
    }
    v1 = v1
    v3 -= (v3 * v3)
    var v7 int8 = 1
    switch "\n\t" {
    case "hello":
        v4 -= (v4 + v4)
    default:
        var v8 int8 = v7
        _ = v8
    }
    if false {
        v9 := []int{42, 1, 2147483647, 4294967295}
        _ = v9
    }
    switch 32767 {
    case 32768:
        var v10 uint16 = v1
        _ = v10
    }
    _ = len(v2)
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS1{F0: 1, F1: 127, F2: 335}
        switch 127 {
        case 315:
            var v13 int8 = (^(v7))
            _ = v13
        }
        v2[(("a" + "a") + ("\n\t" + "a"))] = 309
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
