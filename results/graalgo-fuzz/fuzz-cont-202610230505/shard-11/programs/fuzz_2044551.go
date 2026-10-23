package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"\n\t", "a"}
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 int8 = -1
    var v6 uint16 = 183
    var v7 uint16 = v6
    v5 = (v5 / ((v5) | 1))
    var v8 string = "\n\t"
    v9 := GenS1{F0: 0, F1: 1}
    v10 := v4[2147483647]
    fmt.Println(v6)
    var v11 bool = true
    fmt.Println(v6, v8)
    for fi12 := 0; fi12 < 17; fi12++ {
        switch fi12 {
        case 766:
            v7 = 552
        case -2147483648:
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 string = ("foo bar" + v8)
                _ = fi13
                _ = v14
            }
        case 9223372036854775807:
            var v15 uint = 1
            _ = v15
        default:
            v7 = 65535
        }
        if ((2147483647 > 65535) && ((18446744073709551615 != 255) && v11)) {
            var v16 float64 = 1e10
            var v17 uint64 = 18446744073709551615
            fmt.Println(v17)
            _ = v16
            _ = v17
        }
        switch "\n\t" {
        case "\n\t":
            v10 = (((v10 + v10) + v10) + "")
        }
        _ = fi12
    }
    var v18 int = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
