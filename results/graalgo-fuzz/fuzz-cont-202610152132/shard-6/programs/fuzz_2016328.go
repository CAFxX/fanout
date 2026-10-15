package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16) int32 {
    v6 := []int{902, -2147483648, 32768, 128}
    v7 := GenS1{F0: 594}
    _ = p0
    _ = v6
    _ = v7
    return 646
}
func genFn1(p0 uint32) uint32 {
    fmt.Println(p0)
    var v8 int8 = 40
    _ = p0
    _ = v8
    return p0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (-2.25 - 3.14159)
    v3 := make(map[string]int)
    v4 := []uint32{100, 4294967295, 320, 412}
    var v5 bool = true
    v9 := GenS1{F0: genFn0(32767)}
    var v10 int = 1
    for fi11 := 0; fi11 < 20; fi11++ {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
        var v12 int8 = 112
        _ = fi11
        _ = v12
    }
    v10 += (v10 << ((65535) & 63))
    v13 := GenS1{F0: int32(v10)}
    fmt.Println(v2, v10)
    fmt.Println(v5)
    var v14 float32 = (0.1 - 1.5)
    fmt.Println(v14, v2, v14)
    if len(v4) > 0 {
        v4[0] = genFn1(2)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
