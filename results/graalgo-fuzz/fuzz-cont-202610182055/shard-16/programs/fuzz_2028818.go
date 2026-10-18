package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
func genFn0(p0 uint32) int64 {
    var v4 int64 = 1
    v5 := GenS1{F0: v4, F1: (-(v4)), F2: true}
    var v6 uint16 = 65535
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (v4 * (v4 + (^(v4))))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 42
    var v3 string = ("" + "foo bar")
    fmt.Println(v2)
    var v7 int32 = 42
    v3 = "\n\t"
    var v8 int16 = -1
    if true {
        v8 -= (v8 >> ((18446744073709551615) & 63))
        var v9 bool = (float64(0.1) > (1.5 + (0.1 * float64(float32(3.14159)))))
        _ = v9
    }
    fmt.Println(v8, v2, v7)
    for fi10 := 0; fi10 < 16; fi10++ {
        v11 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: true}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 1, F1: -655, F2: (false || ((-(v7)) != (v7 * -236)))}
        var v14 uint64 = 4294967295
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 4294967295
    v16 := []int{4294967295, (v2 - 4294967295), v2, (v2 / ((v2) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
