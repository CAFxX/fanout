package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0() uint32 {
    var v6 uint32 = 977
    var v7 uint64 = 18446744073709551615
    switch ("foo bar" + ("a" + "\n\t")) {
    case "foo bar":
        v8 := []int{0, 269, -991, 0}
        _ = v8
    case "":
        var v9 int64 = 975
        _ = v9
    default:
        fmt.Println(v7, v6)
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{}
        _ = len(v11)
        _ = fi10
        _ = v11
    }
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int8, p1 uint16, p2 uint) int8 {
    v12 := GenS1{F0: 32768, F1: 1}
    var v13 float64 = float64(float32(1.5))
    var v14 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return (p0 & p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 4294967295
    var v3 uint64 = (v2 / ((v2) | 1))
    var v4 int32 = -622
    var v5 uint = 599
    v2 += 18446744073709551615
    if true {
        var v15 uint = (^(v5))
        _ = v15
    } else {
        v2 = 0
    }
    v3 = (v3 % (((^(v3))) | 1))
    var v16 int32 = v4
    switch ("hello" + (("foo bar" + "") + "a")) {
    case "\n\t":
        var v17 uint8 = 122
        _ = v17
    }
    v4 = v4
    if (v5 == 189) {
        v18 := []int{0}
        v18 = append(v18, 680)
        if (-368 >= (v4 / ((v4) | 1))) {
            v19 := GenS1{F0: 27, F1: v2}
            if (!false) {
                v20 := v1[0]
                _ = v20
            }
            v2 += v2
            _ = v19
        }
        _ = v18
    } else {
        if true {
            v3 = (v2 ^ (v3 % ((42) | 1)))
            fmt.Println(v2, v3)
        }
        var v21 int16 = 1000
        _ = v21
    }
    v4 = (v16 & 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
