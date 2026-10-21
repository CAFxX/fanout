package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
func genFn0(p0 int16, p1 uint) uint {
    fmt.Println(p1)
    var v7 int64 = -9223372036854775808
    var v8 int8 = 117
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (^(p1))
}

func main() {
    var v1 bool = true
    var v2 float64 = -2.25
    v3 := make(map[uint32]string)
    var v4 uint8 = 44
    v5 := []int{294, -698, 0, -754}
    var v6 uint32 = 2147483647
    if v1 {
        if len(v5) > 0 {
            v5[0] = -984
        }
    }
    v6 += 4294967295
    var v9 uint64 = 428
    v10 := GenS1{F0: -2147483648}
    if false {
        v6 -= (v6 << ((2) & 63))
        v6 = v6
        if (337 >= -9223372036854775808) {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: (fi11 / ((4294967295) | 1))}
                v6 = v6
                _ = fi11
                _ = v12
            }
            var v13 int = -9223372036854775808
            _ = v13
        }
    }
    if false {
        v14 := GenS1{F0: -9223372036854775808}
        fmt.Println(v2)
        v1 = true
        _ = v14
    }
    v15 := GenS2{F0: "foo bar", F1: 2, F2: 42}
    v6 += (^(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
