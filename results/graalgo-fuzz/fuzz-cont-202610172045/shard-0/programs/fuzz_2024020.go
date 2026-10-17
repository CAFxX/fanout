package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint, p2 int32) int32 {
    var v6 uint64 = 1
    v7 := GenS1{F0: 0, F1: 0}
    fmt.Println(p0, v6)
    v8 := []int{}
    v9 := GenS1{F0: int16(v6), F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p2
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 uint16 = 2
    v5 := make(map[int]int)
    var v10 uint64 = 0
    var v11 bool = (!(-829 <= 32768))
    if v11 {
        var v12 int = int(v4)
        var v13 int32 = 2147483647
        _ = v12
        _ = v13
    }
    if false {
        v4 -= v4
        var v14 float64 = (float64(3.14159) / 1.0)
        if v11 {
            if v11 {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v10 += 42
                    _ = fi15
                }
            } else {
                fmt.Println(v3)
                v4 += 255
            }
            var v16 uint32 = 2
            _ = v16
        }
        _ = v14
    } else {
        v10 = (v10 ^ 42)
    }
    v11 = (18446744073709551615 >= uint(v10))
    v10 += 0
    var v17 uint8 = 42
    if (-1 == 503) {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS1{F0: -32768, F1: 100}
            _ = fi18
            _ = v19
        }
        var v20 int16 = -32768
        _ = v20
    }
    var v21 int64 = 42
    switch 0 {
    case -2147483648:
        v22 := GenS1{F0: 1, F1: 1}
        _ = v22
    default:
        v23 := GenS2{F0: (("\n\t" + "\n\t") + "\n\t")}
        _ = v23
    }
    v24 := GenS1{F0: 920, F1: -101}
    fmt.Println(v21, v3, v17)
    fmt.Println(v3, v11, v17)
    v25 := v5[32768]
    var v26 int = v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
