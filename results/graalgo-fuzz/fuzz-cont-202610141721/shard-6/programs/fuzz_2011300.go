package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 int }

func main() {
    var v1 uint = 2
    var v2 float64 = 3.14159
    v3 := make(map[string]uint32)
    var v4 bool = ((v1 ^ (v1 % ((645) | 1))) != v1)
    var v5 float32 = float32(v2)
    var v6 uint8 = 0
    switch ("" + "\n\t") {
    case "foo bar":
        v4 = v4
    case "\n\t":
        v7 := []int{954, -9223372036854775808}
        _ = v7
    }
    v4 = v4
    fmt.Println(v5)
    var v8 bool = (((0.0 * float32((v2 / -2.25))) * 0.0) > float32(((1.5 - v2) * ((v2 + ((((float64(v5) - (-2.25 / 1.0)) - 1.5) - v2) * v2)) + (0.0 - 3.14159)))))
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(v6)
        var v10 int = fi9
        if v8 {
            v11 := []int{fi9, v10}
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            if (int64(v6) >= -1) {
                v5 = float32(1.5)
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := []int{-169, (fi12 & (fi12 << ((0) & 63)))}
        v13 = append(v13, fi12)
        _ = fi12
        _ = v13
    }
    switch "hello" {
    case "":
        switch -98 {
        case 9223372036854775807:
            v14 := []int{127, -9223372036854775808, 1}
            _ = v14
        case 128:
            var v15 float64 = float64(v5)
            _ = v15
        default:
            v16 := GenS1{F0: 1}
            _ = v16
        }
    default:
        v5 -= 0.0
    }
    fmt.Println(v6, v6, v8)
    fmt.Println(v6)
    fmt.Println(v1, v1)
    v17 := GenS1{F0: (v1 / ((v1) | 1))}
    v18 := GenS2{F0: (^(v6)), F1: 0, F2: -9223372036854775808}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
