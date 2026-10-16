package main

import (
    "fmt"
)

func genFn0(p0 int8, p1 uint8) uint64 {
    var v5 int8 = (^(p0))
    if false {
        var v6 int8 = (v5 ^ 127)
        fmt.Println(v5, p0)
        v6 += p0
        _ = v6
    }
    var v7 int64 = -1
    switch "\n\t" {
    case "a":
        v8 := []int{-447, 295, 32767, 127}
        _ = v8
    default:
        var v9 int = 128
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v7
    return 1
}
func genFn1() int32 {
    var v10 int64 = 442
    v10 -= (v10 - (v10 | (v10 + v10)))
    fmt.Println(v10)
    v10 -= 1
    v10 = v10
    _ = v10
    return 1
}

func main() {
    var v1 int16 = 42
    var v2 float32 = -2.25
    var v3 int16 = (v1 << ((255) & 63))
    v4 := make(map[string]uint32)
    v1 += (v3 + (^(v1)))
    fmt.Println(v3, v2)
    var v11 int64 = 9223372036854775807
    v11 = (v11 % (((v11 | (v11 | (v11 - 9223372036854775807)))) | 1))
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 float64 = -2.25
        _ = fi12
        _ = v13
    }
    var v14 float32 = (v2 * (v2 - v2))
    delete(v4, ("a" + "a"))
    v3 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
