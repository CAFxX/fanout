package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint8, p1 uint8, p2 int) uint16 {
    v7 := []int{}
    if (-2.25 < ((((-2.25 * (0.1 * 1.5)) - 3.14159) - (float64((float32(0.1) * 1e10)) + 1.5)) / (1.5 + float64((float32(float64(0.0)) * float32(1.5)))))) {
        var v8 uint32 = 289
        for fi9 := 0; fi9 < 4; fi9++ {
            v7 = append(v7, (fi9 ^ -963))
            if true {
                v10 := GenS1{F0: 4294967295}
                switch "\n\t" {
                case "foo bar":
                    if false {
                        var v11 uint8 = p1
                        for fi12 := 0; fi12 < 1; fi12++ {
                            var v13 int64 = -9223372036854775808
                            _ = fi12
                            _ = v13
                        }
                        v8 -= (v8 << ((2) & 63))
                        _ = v11
                    }
                default:
                    var v14 float64 = (1.5 / 0.1)
                    _ = v14
                }
                _ = v10
            }
            _ = fi9
        }
        _ = v8
    }
    p1 -= (p1 >> ((2) & 63))
    var v15 float64 = (3.14159 * 0.1)
    if (((!(-2.25 != (3.14159 - 3.14159))) || (!true)) || (450 >= -1)) {
        fmt.Println(p1, v15, v15)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 0
}

func main() {
    v1 := []string{}
    v2 := []uint32{539}
    var v3 uint32 = 2147483647
    var v4 uint64 = 0
    v5 := make(map[uint32]int)
    var v6 int32 = 217
    v6 += v6
    delete(v5, v3)
    v5[(v3 % (((^(v3))) | 1))] = 0
    var v16 uint64 = 18446744073709551615
    fmt.Println(v16)
    var v17 bool = true
    v3 = v3
    v6 -= int32(v3)
    if false {
        v6 += (v6 + v6)
    } else {
        var v18 uint16 = 65535
        _ = v18
    }
    fmt.Println(v17)
    switch "foo bar" {
    case "\n\t":
        _ = len(v5)
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        switch fi19 {
        case -2147483648:
            var v20 uint16 = 437
            _ = v20
        default:
            if true {
                var v21 int32 = (v6 >> ((539) & 63))
                _ = v21
            }
        }
        _ = fi19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
