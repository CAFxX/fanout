package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 bool }
func genFn0() uint8 {
    var v7 uint16 = 559
    switch -2147483648 {
    case 771:
        var v8 uint16 = v7
        _ = v8
    case 9223372036854775807:
        var v9 float64 = (1.5 * 3.14159)
        _ = v9
    }
    _ = v7
    return 0
}
func genFn1(p0 uint32, p1 string) uint16 {
    v10 := []int{-2147483648, 24, 9223372036854775807, 32768}
    p1 = p1
    var v11 float32 = (0.0 - ((float32(((float64(0.0) + ((0.0 / 1.0) / (1.5 / -2.25))) / 3.14159)) - (3.14159 / -2.25)) + (0.1 + 0.1)))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 857
}

func main() {
    var v1 bool = (-854 >= 32768)
    var v2 int8 = -1
    v3 := []uint32{2147483648, 998, 2147483647}
    v4 := make(map[string]bool)
    var v5 uint64 = 42
    v6 := make(map[int]bool)
    if len(v3) > 0 {
        v3[0] = 65535
    }
    var v12 uint8 = 128
    v13 := GenS1{F0: 2147483647, F1: 0, F2: true}
    var v14 string = (("a" + "foo bar") + "foo bar")
    v15 := v4[("a" + "hello")]
    v12 -= genFn0()
    v1 = true
    if v15 {
        switch "foo bar" {
        case "hello":
            switch "\n\t" {
            case "hello":
                v5 -= v5
            default:
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 bool = (103 == 2147483647)
                    _ = fi16
                    _ = v17
                }
            }
        case "foo bar":
            fmt.Println(v14, v15, v12)
        default:
            v3 = append(v3, 381)
        }
        v15 = (false && (-1 >= 2147483647))
    }
    v1 = v15
    var v18 int = 2147483647
    v4[v14] = v15
    fmt.Println(v2, v14, v18)
    v19 := GenS1{F0: 242, F1: 0, F2: (-970 > 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
