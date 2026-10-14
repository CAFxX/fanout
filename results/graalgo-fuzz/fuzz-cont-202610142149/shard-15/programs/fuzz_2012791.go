package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0(p0 uint, p1 uint64, p2 uint32) string {
    var v5 float64 = 0.1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "hello"
}
func genFn1() int16 {
    var v6 uint32 = 2147483648
    var v7 float32 = float32(float64((0.0 - 0.0)))
    v8 := []int{-646}
    _ = len(v8)
    _ = v6
    _ = v7
    _ = v8
    return -742
}

func main() {
    var v1 int8 = 127
    var v2 bool = (129 > 1)
    var v3 uint16 = 32767
    v4 := make(map[string]uint32)
    _ = len(v4)
    v3 += v3
    var v9 uint = 4294967295
    v1 += (v1 << (((^(v9))) & 63))
    fmt.Println(v3)
    var v10 string = "foo bar"
    var v11 uint8 = 212
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 float64 = -2.25
        v13 = v13
        _ = fi12
        _ = v13
    }
    fmt.Println(v1)
    var v14 int16 = 423
    for fi15 := 0; fi15 < 15; fi15++ {
        v2 = v2
        v10 = ("" + v10)
        switch (-(fi15)) {
        case 1:
            var v16 float32 = -2.25
            _ = v16
        default:
            fmt.Println(v10, v3)
        }
        _ = fi15
    }
    var v17 string = "a"
    v18 := GenS1{F0: -2147483648, F1: "\n\t"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
