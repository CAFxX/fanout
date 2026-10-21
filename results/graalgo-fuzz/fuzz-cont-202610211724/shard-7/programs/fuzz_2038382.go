package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int8 }

func main() {
    var v1 uint16 = 737
    var v2 uint16 = 42
    v3 := make(map[string]string)
    var v4 int64 = -9223372036854775808
    var v5 int = 32768
    if (-843 == 0) {
        var v6 uint16 = 788
        v7 := []int{-2147483648, (^(v5))}
        _ = v6
        _ = v7
    }
    switch (("" + "") + "") {
    case "hello":
        if (1e10 <= ((0.1 - 3.14159) + float32(((3.14159 - 3.14159) + (1.5 / ((0.1 - ((0.0 / -2.25) - 3.14159)) - 1.5)))))) {
            v8 := []int{(v5 + (v5 & -9223372036854775808)), v5}
            if true {
                var v9 float32 = 1e10
                _ = v9
            }
            v4 += v4
            _ = v8
        }
    default:
        v10 := GenS1{F0: 2, F1: 127, F2: 100}
        _ = v10
    }
    v11 := []int{(^(v5))}
    v1 += 0
    for fi12 := 0; fi12 < 20; fi12++ {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        if len(v11) > 0 {
            v11[0] = v5
        }
        v2 = (v1 * v2)
        _ = fi12
    }
    v5 -= (^(v5))
    v13 := GenS1{F0: 0, F1: -58, F2: -1}
    var v14 int64 = v4
    var v15 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
