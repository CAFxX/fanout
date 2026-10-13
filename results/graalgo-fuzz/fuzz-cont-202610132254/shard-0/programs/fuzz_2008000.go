package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint, p1 int32, p2 bool) bool {
    v6 := []int{0, 9223372036854775807, 127, 2147483647}
    p0 += (p0 + 65535)
    p0 += (^(p0))
    var v7 uint32 = 4294967295
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (0 == 100)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 155
    var v3 uint64 = 4294967295
    var v4 uint8 = 50
    v5 := make(map[string]bool)
    v9 := GenS1{F0: 0}
    fmt.Println(v3, v2)
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := v1[fi11]
            _ = fi11
            _ = v12
        }
        v2 -= 4294967295
        _ = fi10
    }
    v2 += (v2 + v2)
    fmt.Println(v3)
    var v13 uint32 = 535
    v3 = (v3 ^ 2)
    v14 := GenS1{F0: -204}
    var v15 uint8 = (v4 & (v4 & 42))
    v16 := []int{-2147483648, 127, 2147483647, 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
