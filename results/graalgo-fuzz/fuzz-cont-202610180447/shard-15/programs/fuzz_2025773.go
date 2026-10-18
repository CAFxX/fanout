package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0() uint16 {
    var v7 uint16 = 2
    var v8 int = 227
    var v9 int16 = -473
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 129
    var v3 int64 = 9223372036854775807
    v4 := []uint32{4294967295, 952, 1, 511}
    var v5 float64 = ((1e10 - 3.14159) + (3.14159 + 0.1))
    var v6 uint32 = 42
    v10 := GenS1{F0: 42}
    v11 := GenS1{F0: 4294967295}
    switch 32767 {
    case 4294967295:
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 string = ("\n\t" + (("\n\t" + "a") + (("hello" + "\n\t") + "\n\t")))
            _ = fi12
            _ = v13
        }
    }
    v2 -= (^(v2))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 float64 = ((0.1 - (3.14159 * (1.5 * ((0.0 / ((v5 / 0.1) + v5)) / (((v5 / (float64(float32(v5)) + v5)) + 0.1) / v5))))) * 3.14159)
        v2 += (v2 - (^(v2)))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: (!true)}
    var v17 int = -2147483648
    v3 = v3
    if (false || false) {
        v2 -= v2
        v18 := GenS2{F0: (int32(v17) > -1)}
        _ = v18
    }
    fmt.Println(v2, v5, v5)
    var v19 float32 = ((0.1 / 1.5) + 0.0)
    if (false || true) {
        v6 += 0
    }
    if true {
        switch ("" + "hello") {
        case "hello":
            for fi20 := 0; fi20 < 6; fi20++ {
                var v21 bool = (32767 == -1)
                _ = fi20
                _ = v21
            }
        }
        v22 := GenS3{F0: v2, F1: 0}
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
