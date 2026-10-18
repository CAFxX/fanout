package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 bool }
type GenS2 struct { F0 uint; F1 int32; F2 int32 }

func main() {
    var v1 int8 = -70
    var v2 uint32 = 686
    var v3 string = "hello"
    v4 := make(map[string]uint32)
    v5 := []uint32{(v2 + 821), v2, (^(v2)), 65535}
    v6 := []string{((v3 + v3) + ((v3 + "a") + v3)), v3, ((v3 + "foo bar") + v3)}
    var v7 int8 = (v1 + v1)
    var v8 uint16 = 178
    v9 := GenS1{F0: 1, F1: -616, F2: false}
    fmt.Println(v8, v1)
    v10 := GenS1{F0: 128, F1: 9223372036854775807, F2: (true || (498 < -1))}
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    if false {
        v11 := GenS1{F0: 0, F1: 0, F2: (1 <= 1)}
        for fi12 := 0; fi12 < 18; fi12++ {
            var v13 int32 = 2147483647
            v14 := GenS2{F0: 10, F1: v13, F2: (-(v13))}
            var v15 int64 = 1
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        var v16 uint = 42
        _ = v11
        _ = v16
    } else {
        v4["hello"] = (v2 ^ v2)
        v6 = append(v6, ("" + v3))
        v1 -= 0
    }
    _ = len(v4)
    v17 := GenS2{F0: 2, F1: -780, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
