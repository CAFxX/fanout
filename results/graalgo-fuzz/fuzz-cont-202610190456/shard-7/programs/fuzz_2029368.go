package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 uint32 = 2147483647
    v2 := []int{}
    v3 := []string{"", ""}
    v4 := make(map[string]bool)
    var v5 uint = 18446744073709551615
    v5 = 227
    v6 := GenS1{F0: 1, F1: -32768, F2: 9223372036854775807}
    if false {
        for fi7 := 0; fi7 < 13; fi7++ {
            fmt.Println(v5)
            for fi8 := 0; fi8 < 12; fi8++ {
                if len(v2) > 0 {
                    v2[0] = 9223372036854775807
                }
                var v9 uint16 = 412
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    var v10 float32 = 0.1
    v11 := GenS2{F0: 0}
    fmt.Println(v1, v1, v5)
    var v12 uint32 = v1
    var v13 string = "\n\t"
    var v14 float32 = ((v10 + v10) * 1.5)
    v1 -= 742
    v10 += 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
