package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 int64, p1 uint) uint32 {
    var v4 int64 = (p0 - (-(p0)))
    for fi5 := 0; fi5 < 1; fi5++ {
        v4 -= (p0 << ((p1) & 63))
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{480, 4294967295}
    var v3 int64 = -1
    fmt.Println(v1, v3)
    var v6 float32 = (3.14159 + 0.0)
    switch -9223372036854775808 {
    case 4294967295:
        var v7 int16 = 1
        _ = v7
    default:
        var v8 uint16 = 0
        _ = v8
    }
    var v9 int = -889
    var v10 int32 = 2147483647
    var v11 float64 = 0.0
    var v12 float64 = (float64(0.1) * (((v11 * 3.14159) / 1.0) - v11))
    v11 = (v12 + (v11 + v12))
    v9 += v9
    v13 := GenS1{F0: 315, F1: (v10 * (v10 / (((v10 | v10)) | 1)))}
    fmt.Println(v6)
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 bool = true
            var v17 int64 = (v3 / ((v3) | 1))
            v18 := GenS1{F0: 18446744073709551615, F1: v10}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            switch "hello" {
            case "\n\t":
                var v20 int64 = 42
                _ = v20
            default:
                var v21 int = 127
                _ = v21
            }
            fmt.Println(v3, fi19)
            _ = fi19
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
