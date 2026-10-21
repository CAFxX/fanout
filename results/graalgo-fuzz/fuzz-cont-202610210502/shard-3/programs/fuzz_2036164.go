package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint32, p1 uint8, p2 uint) int {
    p2 = (^(p2))
    fmt.Println(p2, p0)
    var v6 uint = 437
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 32768
}
func genFn1(p0 int, p1 uint64, p2 int16) string {
    p2 = (p2 >> ((65535) & 63))
    v7 := []int{}
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 bool = (!false)
        var v11 bool = v10
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 157
    v3 := []int{32768, -718}
    var v4 uint64 = 458
    v5 := make(map[int]string)
    if (-128 >= -128) {
        v2 -= v2
        var v12 uint8 = 104
        _ = v12
    }
    if (!(v4 != v4)) {
        v2 = v2
    }
    v4 += v4
    var v13 float64 = float64(1e10)
    var v14 uint = 20
    var v15 uint = v2
    v16 := GenS1{F0: v15, F1: 42}
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint64 = v4
            _ = len(v3)
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    var v20 uint8 = 128
    fmt.Println(v2)
    v14 = 956
    _ = len(v3)
    var v21 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
