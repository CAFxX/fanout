package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint32; F1 bool }

func main() {
    var v1 float32 = (0.1 - ((3.14159 / float32((0.0 + 3.14159))) * -2.25))
    v2 := []string{"a", "foo bar"}
    v3 := make(map[int]int)
    var v4 uint32 = 1
    v5 := GenS1{F0: 567, F1: 9223372036854775807, F2: true}
    v1 = ((0.0 / (v1 * v1)) / v1)
    v6 := GenS2{F0: 230, F1: (v4 - (v4 ^ v4))}
    if (v4 < v4) {
        v7 := GenS2{F0: 1, F1: (v4 ^ (v4 * (v4 >> ((18446744073709551615) & 63))))}
        _ = v7
    }
    v4 += (^(v4))
    v1 = (v1 - (v1 / v1))
    v4 = v4
    if false {
        var v8 int = -210
        _ = v8
    }
    fmt.Println(v1, v1)
    v3[998] = -9223372036854775808
    switch "" {
    case "a":
        switch 962 {
        case 42:
            v9 := GenS2{F0: -32768, F1: (v4 % ((v4) | 1))}
            _ = v9
        }
    }
    var v10 int16 = 32767
    v11 := GenS3{F0: (v4 | v4), F1: false}
    var v12 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
