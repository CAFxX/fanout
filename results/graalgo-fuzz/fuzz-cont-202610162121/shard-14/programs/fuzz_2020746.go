package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []uint32{549, 569}
    var v2 int64 = -9223372036854775808
    v3 := []uint32{0}
    var v4 uint32 = 2147483647
    v5 := []uint32{(v4 / (((v4 % ((1) | 1))) | 1))}
    var v6 int32 = -1
    v4 = 42
    var v7 uint16 = 1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = (^(v6))
    if (!(18446744073709551615 != 42)) {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int16 = -32768
            _ = fi8
            _ = v9
        }
    }
    var v10 uint16 = 42
    var v11 float64 = (1.5 * float64(-2.25))
    v12 := GenS1{F0: (true && (255 > 255))}
    v6 = (v6 * (v6 % ((v6) | 1)))
    v4 = (v4 / ((v4) | 1))
    _ = len(v1)
    v7 = v10
    v6 -= -1
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
