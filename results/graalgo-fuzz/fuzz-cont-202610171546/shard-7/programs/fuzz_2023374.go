package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
func genFn0() uint {
    v6 := GenS1{F0: 255, F1: 42}
    var v7 int64 = -1
    v8 := []int{500, 0}
    v7 += v7
    _ = len(v8)
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 uint32 = 2
    var v2 uint64 = 0
    var v3 uint8 = 119
    var v4 uint16 = 323
    var v5 string = ("foo bar" + ("foo bar" + ("a" + "\n\t")))
    v3 = (v3 + 96)
    v1 += (v1 - v1)
    v9 := GenS1{F0: (^(v4)), F1: 1}
    v10 := []int{}
    v4 += v4
    v3 += (v3 + uint8(v1))
    if (0 == 42) {
        var v11 int16 = -510
        switch -17 {
        case 1:
            var v12 float64 = -2.25
            _ = v12
        default:
            if true {
                var v13 bool = false
                var v14 float64 = 3.14159
                var v15 int32 = -2147483648
                _ = v13
                _ = v14
                _ = v15
            }
        }
        v10 = append(v10, -9223372036854775808)
        _ = v11
    } else {
        var v16 float64 = 3.14159
        if (int16(v2) >= 0) {
            var v17 float32 = -2.25
            _ = v17
        } else {
            v3 = 255
            v16 = v16
            v5 = v5
        }
        var v18 bool = false
        _ = v16
        _ = v18
    }
    v5 = ("a" + "a")
    if ((2 < 255) || (1e10 != (3.14159 + float32((0.0 / 1.0))))) {
        var v19 bool = false
        v10 = append(v10, 127)
        _ = v19
    } else {
        v1 += v1
        v10 = append(v10, 0)
        fmt.Println(v3, v1, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
