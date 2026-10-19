package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 uint64, p1 uint64) int64 {
    v7 := []int{-817, -9223372036854775808, 9223372036854775807}
    for fi8 := 0; fi8 < 1; fi8++ {
        switch 665 {
        case 9223372036854775807:
            var v9 int16 = 582
            _ = v9
        }
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: 42, F1: int8(fi8), F2: 2147483647}
            _ = v10
        }
        fmt.Println(fi8, p1, p1)
        _ = fi8
    }
    p0 = (p1 / ((p0) | 1))
    v11 := GenS1{F0: 9223372036854775807, F1: 127, F2: 376}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return -812
}

func main() {
    var v1 bool = (false || (4294967295 >= 4294967295))
    v2 := []string{"hello", ""}
    v3 := []uint32{2}
    var v4 int8 = 0
    var v5 int64 = 9223372036854775807
    var v6 uint16 = 178
    _ = len(v2)
    if (v1 && (-852 > 32767)) {
        var v12 int = 638
        var v13 uint64 = 18446744073709551615
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v12
        _ = v13
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v14 uint16 = (v6 % ((65535) | 1))
        _ = v14
    }
    var v15 uint32 = 1
    var v16 float32 = (1e10 + (3.14159 - 0.1))
    if (!v1) {
        v3 = append(v3, (v15 & v15))
        var v17 bool = (0 == genFn0(2, 880))
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        fmt.Println(v6, v15)
        _ = fi18
    }
    v19 := GenS2{F0: v6, F1: "", F2: 127}
    fmt.Println(v4, v5, v6)
    var v20 uint = 4294967295
    fmt.Println(v4, v4)
    fmt.Println(v1, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6, v20, v1)
    v4 -= (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
