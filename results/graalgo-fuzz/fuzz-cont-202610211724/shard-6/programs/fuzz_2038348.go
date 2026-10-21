package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint }
func genFn0(p0 string, p1 int16, p2 uint32) int8 {
    p1 = p1
    var v7 float32 = ((0.0 * float32((0.1 * float64(-2.25)))) * 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 53
}

func main() {
    var v1 float32 = float32(1e10)
    v2 := []string{"foo bar", "hello"}
    v3 := []string{"", "foo bar", "\n\t"}
    var v4 int32 = 805
    var v5 int16 = -20
    v6 := []uint32{264, 453, 2147483647, 42}
    var v8 uint16 = 2
    v9 := GenS1{F0: ("hello" + ""), F1: 1, F2: v8}
    v10 := GenS2{F0: 42}
    v2 = append(v2, "hello")
    var v11 uint = 2
    v12 := GenS3{F0: 798, F1: -1, F2: (^(v11))}
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 int32 = (v4 >> (((v11 + v11)) & 63))
        _ = fi13
        _ = v14
    }
    var v15 string = ("" + "foo bar")
    switch 810 {
    case -200:
        var v16 uint32 = 70
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v8, v1)
        _ = fi17
    }
    v11 -= v11
    v18 := GenS2{F0: 42}
    v1 = (float32(float64(3.14159)) / ((v1 / (v1 * ((v1 * v1) - v1))) * v1))
    v19 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: uint(v8)}
    for fi20 := 0; fi20 < 11; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            _ = len(v3)
            _ = fi21
        }
        if true {
            v5 = v5
            v1 = v1
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
