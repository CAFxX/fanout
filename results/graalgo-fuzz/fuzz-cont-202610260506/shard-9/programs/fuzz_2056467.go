package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{788, 32768}
    v2 := []string{"a"}
    var v3 int16 = 42
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 42
    _ = len(v4)
    v5 -= (v5 / ((2) | 1))
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(v5)
        if (649 > 42) {
            var v7 float64 = 1e10
            _ = v7
        } else {
            if false {
                if true {
                    var v8 uint64 = 18446744073709551615
                    v9 := v4[0]
                    _ = v8
                    _ = v9
                } else {
                    var v10 int = (fi6 + -2147483648)
                    _ = v10
                }
                fmt.Println(fi6, fi6)
                _ = len(v1)
            }
        }
        fmt.Println(v5, v5)
        _ = fi6
    }
    if false {
        var v11 bool = true
        v5 += 255
        var v12 int64 = -9223372036854775808
        _ = v11
        _ = v12
    }
    var v13 uint32 = 65535
    v5 -= 42
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 367}
        fmt.Println(v3)
        _ = fi14
        _ = v15
    }
    var v16 uint32 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
