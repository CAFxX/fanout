package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
func genFn0(p0 bool) int64 {
    p0 = (721 < 9223372036854775807)
    for fi5 := 0; fi5 < 17; fi5++ {
        p0 = (true || p0)
        v6 := GenS1{F0: 42}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: -1}
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    v3 := make(map[uint32]bool)
    var v4 float64 = -2.25
    var v8 uint8 = 199
    v8 += (v8 << ((255) & 63))
    v8 += 221
    if (!true) {
        var v9 uint32 = uint32(v8)
        v8 = 255
        _ = v9
    }
    if (245 >= 969) {
        var v10 string = "a"
        _ = v10
    } else {
        v4 = 1e10
    }
    v4 += v4
    v11 := GenS1{F0: -1}
    v12 := GenS2{F0: 111, F1: false, F2: 1}
    var v13 uint64 = 18446744073709551615
    v4 = ((v4 + float64(0.0)) + 1e10)
    var v14 float32 = (3.14159 + (((-2.25 / 0.1) - 1e10) / -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
