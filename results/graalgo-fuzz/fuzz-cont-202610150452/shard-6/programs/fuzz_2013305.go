package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint16) uint16 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return (p0 << ((65535) & 63))
}
func genFn1(p0 bool, p1 uint, p2 uint8) bool {
    v7 := GenS1{F0: p0}
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 0
    v3 := []uint32{uint32(v1)}
    var v4 float32 = float32(float64(((float32((0.0 + (-2.25 / 1.0))) + (1.5 - (float32(1e10) - 1.5))) / 0.1)))
    v5 := []int{}
    var v6 uint = 18446744073709551615
    fmt.Println(v6, v6)
    var v8 uint64 = 42
    v2 = v2
    v4 = v4
    v2 += 0
    fmt.Println(v8, v8)
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS2{F0: 2147483647, F1: 927}
        v6 = 27
        if false {
            var v11 int64 = 9223372036854775807
            _ = v11
        } else {
            v12 := GenS2{F0: 1, F1: 947}
            var v13 string = "hello"
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v14 := GenS2{F0: 2147483647, F1: 621}
    for fi15 := 0; fi15 < 15; fi15++ {
        v2 -= v2
        _ = fi15
    }
    var v16 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
