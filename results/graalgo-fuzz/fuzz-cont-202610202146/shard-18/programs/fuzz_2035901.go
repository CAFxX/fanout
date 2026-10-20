package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int8; F2 int32 }
func genFn0(p0 int16, p1 uint64, p2 int32) int64 {
    v7 := GenS1{F0: 1}
    for fi8 := 0; fi8 < 10; fi8++ {
        fmt.Println(p1, p2, p1)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := []string{}
    var v2 string = "hello"
    v3 := make(map[int]bool)
    var v4 uint32 = 2147483648
    var v5 uint64 = 755
    var v6 int = -9223372036854775808
    _ = len(v3)
    var v9 int32 = 222
    var v10 bool = true
    v9 -= 1
    if (false || (v10 || true)) {
        for fi11 := 0; fi11 < 4; fi11++ {
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v4)
                fmt.Println(v9)
                v2 = v2
                _ = fi12
            }
            var v13 int8 = 0
            var v14 uint8 = 42
            _ = fi11
            _ = v13
            _ = v14
        }
    }
    v10 = (!(v10 && v10))
    var v15 int = (v6 + 2147483647)
    fmt.Println(v2, v5)
    if ((0.1 / (float32((0.0 - 1.5)) + (0.0 * 0.0))) > 1e10) {
        v16 := GenS2{F0: -1, F1: -100, F2: (v9 % (((v9 >> ((371) & 63))) | 1))}
        _ = v16
    }
    _ = len(v1)
    var v17 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
