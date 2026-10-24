package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 string }
func genFn0(p0 uint8, p1 uint8) uint64 {
    v6 := []int{2147483647, 455}
    var v7 int8 = 100
    var v8 float32 = (-2.25 / (((-2.25 + 1e10) + -2.25) * 1e10))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    v9 := GenS1{F0: 1, F1: 0, F2: "\n\t"}
    switch 664 {
    case 0:
        if false {
            var v10 uint16 = 2
            v11 := []int{43}
            v11 = append(v11, 1)
            _ = v10
            _ = v11
        }
    case -9223372036854775808:
        v12 := []int{-398, 32768}
        _ = v12
    case 2147483647:
        v13 := []int{}
        _ = v13
    }
    var v14 uint = 2
    v14 -= v14
    _ = v14
    _ = v9
    return ""
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]uint32)
    var v3 float32 = float32(((((((1e10 - (-2.25 * 0.0)) * 0.1) / 1.0) - 0.1) + 0.0) + (0.0 + (float64((float32((1.5 + (float64(0.0) / 1.5))) * 1e10)) / 1.0))))
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 128
    fmt.Println(v5)
    if false {
        v3 = v3
        switch (("hello" + ("foo bar" + "")) + "") {
        case "a":
            v1 = 305
        case "foo bar":
            v15 := GenS1{F0: 1, F1: -9223372036854775808, F2: "\n\t"}
            _ = v15
        case "\n\t":
            fmt.Println(v5, v5, v1)
        }
        var v16 int32 = -611
        _ = v16
    } else {
        v5 += (v5 | (v5 | (v5 | v5)))
    }
    var v17 int16 = (v1 >> ((694) & 63))
    var v18 int8 = -11
    var v19 uint32 = 648
    fmt.Println(v1, v17)
    fmt.Println(v1, v3, v3)
    v19 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
