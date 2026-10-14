package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int64 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]uint32)
    var v3 uint8 = uint8(v1)
    var v4 uint64 = 18446744073709551615
    var v5 float64 = -2.25
    v5 += (v5 - -2.25)
    v4 += (v4 | (v4 << ((v1) & 63)))
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS1{F0: 2147483647}
        fmt.Println(fi6)
        v2[("foo bar" + ("" + (("" + ("foo bar" + "")) + "\n\t")))] = 0
        _ = fi6
        _ = v7
    }
    v4 += (v4 & v4)
    if true {
        v8 := GenS2{F0: v1, F1: 1}
        fmt.Println(v4, v3, v1)
        v4 += (^(v4))
        _ = v8
    } else {
        _ = len(v2)
        v4 = 1
    }
    var v9 float32 = 0.1
    if true {
        v10 := GenS1{F0: 42}
        var v11 uint32 = 2
        switch "a" {
        case "foo bar":
            fmt.Println(v3, v9, v3)
        case "hello":
            switch 42 {
            case -210:
                v1 += v1
            case 576:
                v5 += ((1e10 + (v5 * v5)) * (((float64((((((3.14159 + (v9 + v9)) * v9) * (float32(v5) / v9)) / ((v9 / ((v9 - (0.1 * (1e10 * (1e10 / (v9 / v9))))) + v9)) / float32((v5 / v5)))) + (v9 / v9))) + v5) / v5) - (v5 * ((v5 - ((((0.0 + (v5 - ((v5 / (3.14159 - (float64(3.14159) - v5))) * v5))) * float64(((v9 - v9) - float32((v5 / v5))))) / v5) / v5)) + (((((-2.25 * v5) + v5) * v5) - 0.0) * (1e10 / 1.0))))))
            default:
                v12 := []int{2147483647, 777, 9223372036854775807, 667}
                _ = v12
            }
        }
        _ = v10
        _ = v11
    }
    v5 = 3.14159
    if true {
        v3 += v3
    }
    v13 := []int{-2147483648, 166, 264, -997}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
