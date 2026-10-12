package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int8 }
func genFn0() int64 {
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int = fi7
        fmt.Println(v6, v6)
        _ = fi7
        _ = v8
    }
    v6 = v6
    switch 9223372036854775807 {
    case 128:
        v6 = "a"
    case 127:
        v9 := GenS1{F0: 0, F1: v6, F2: 127}
        _ = v9
    default:
        var v10 uint64 = 4294967295
        _ = v10
    }
    _ = v6
    return 829
}
func genFn1(p0 int8, p1 uint32, p2 int8) uint8 {
    p0 = (p0 << ((18446744073709551615) & 63))
    var v11 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (v11 * v11)
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = -2.25
    v3 := make(map[int]uint32)
    var v4 float32 = (float32((((1.5 - v2) / (3.14159 * (1e10 - v2))) + (v2 / 1.0))) * 1.5)
    v5 := make(map[uint32]uint32)
    var v12 uint32 = 2147483647
    fmt.Println(v12)
    var v13 float32 = v4
    var v14 int16 = 1
    for fi15 := 0; fi15 < 8; fi15++ {
        delete(v3, (fi15 * 32767))
        fmt.Println(v12, v4, v2)
        for fi16 := 0; fi16 < 14; fi16++ {
            if false {
                if (!false) {
                    _ = len(v5)
                    v1[("\n\t" + "")] = false
                }
            }
            delete(v3, (fi15 + fi15))
            delete(v3, -998)
            _ = fi16
        }
        _ = fi15
    }
    var v17 float64 = v2
    var v18 uint32 = 65535
    v17 = v17
    var v19 int8 = 42
    var v20 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
