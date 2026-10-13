package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }

func main() {
    var v1 uint32 = 417
    var v2 uint16 = 1
    var v3 int16 = -1
    v4 := []int{588, 42, 163}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    var v7 uint8 = 42
    v1 = v1
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := GenS1{F0: v2, F1: 128, F2: 182}
        _ = fi8
        _ = v9
    }
    var v10 int = 127
    if ((!true) || (-2.25 != (((3.14159 * 1.5) - 3.14159) + (3.14159 * ((-2.25 - 3.14159) * 0.0))))) {
        var v11 int32 = -1
        switch "\n\t" {
        case "\n\t":
            var v12 uint = 18446744073709551615
            _ = v12
        case "foo bar":
            _ = len(v6)
        }
        fmt.Println(v3)
        _ = v11
    } else {
        v13 := GenS1{F0: (v2 ^ (v2 + v2)), F1: (v7 - (v7 / ((207) | 1))), F2: 9223372036854775807}
        if (2147483647 < 0) {
            var v14 string = "a"
            fmt.Println(v2, v7)
            _ = v14
        } else {
            var v15 float64 = (1.5 / 3.14159)
            if len(v4) > 0 {
                v4[0] = (v10 / ((v10) | 1))
            }
            var v16 int = v10
            _ = v15
            _ = v16
        }
        _ = v13
    }
    var v17 uint64 = 2
    v2 = (v2 / (((v2 >> ((2) & 63))) | 1))
    v18 := GenS1{F0: (v2 >> ((502) & 63)), F1: 117, F2: 539}
    v2 = (v2 ^ (v2 | v2))
    if len(v4) > 0 {
        v4[0] = (v10 % ((v10) | 1))
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        fmt.Println(v3, v10, v2)
        delete(v6, 0)
        if true {
            for fi20 := 0; fi20 < 20; fi20++ {
                v10 = fi20
                _ = fi20
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        _ = fi19
    }
    delete(v6, (v1 - (v1 ^ (v1 * uint32(v7)))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
