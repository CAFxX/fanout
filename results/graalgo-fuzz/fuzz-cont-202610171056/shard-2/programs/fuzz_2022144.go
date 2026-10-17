package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 uint64 = 86
    var v5 int64 = 9223372036854775807
    var v6 uint16 = 2
    var v7 float64 = 0.0
    v4 -= 405
    v8 := []int{9223372036854775807, -2147483648, 32767}
    for fi9 := 0; fi9 < 13; fi9++ {
        v7 -= v7
        var v10 uint32 = 1
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = (-(v5))
            var v14 uint8 = 42
            var v15 int = fi11
            switch "hello" {
            case "hello":
                var v16 uint8 = v14
                _ = v16
            default:
                v13 += v5
            }
            fmt.Println(v4, fi12)
            var v17 int64 = 0
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
            _ = v17
        }
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 bool = (fi18 == fi11)
            _ = fi18
            _ = v19
        }
        _ = fi11
    }
    var v20 uint64 = v4
    if false {
        v6 += 42
    } else {
        delete(v2, 42)
        var v21 uint8 = 1
        _ = v21
    }
    var v22 int16 = 0
    v23 := v2[2147483648]
    v1 -= int32(v23)
    v20 = (v20 | (v4 & v20))
    var v24 uint64 = v20
    var v25 float64 = float64((((float32(v7) - (float32(v7) - (float32(float64(float32(v7))) - float32(float64(float32(0.0)))))) / (float32((v7 - 1.5)) / (1.5 * (3.14159 * ((0.0 - ((((0.1 / -2.25) / 1e10) - 3.14159) - -2.25)) - 0.0))))) * -2.25))
    v26 := v3[v23]
    v22 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
