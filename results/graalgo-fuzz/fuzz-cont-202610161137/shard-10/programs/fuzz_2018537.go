package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 int16; F2 bool }
func genFn0(p0 int8, p1 uint32, p2 uint32) int64 {
    var v6 float32 = 0.0
    fmt.Println(p2, p0)
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 uint16, p1 uint8) bool {
    p1 = (p1 & 128)
    var v7 float64 = (float64(3.14159) + 0.0)
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = v7
    return true
}

func main() {
    var v1 int8 = -80
    var v2 bool = false
    v3 := []int{}
    v4 := make(map[int]uint32)
    var v5 uint64 = 481
    switch 2147483647 {
    case 388:
        v8 := GenS1{F0: true}
        _ = v8
    case 4294967295:
        fmt.Println(v5, v1)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 uint8 = 0
    _ = len(v4)
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: v2}
        _ = fi10
        _ = v11
    }
    var v12 uint = 18446744073709551615
    v1 += v1
    if (("hello" + ("\n\t" + "foo bar")) > "a") {
        v1 = v1
    }
    v5 = (v5 | v5)
    v13 := GenS1{F0: v2}
    v14 := GenS1{F0: (-994 <= -2147483648)}
    fmt.Println(v12, v1)
    fmt.Println(v9, v1, v5)
    if (true && v2) {
        v15 := GenS2{F0: (^(v12)), F1: 42, F2: v2}
        switch 128 {
        case 0:
            var v16 uint32 = 155
            _ = v16
        }
        var v17 float32 = 1.5
        _ = v15
        _ = v17
    }
    var v18 uint8 = uint8(v12)
    switch -108 {
    case 32767:
        fmt.Println(v1, v12)
    default:
        v18 = uint8(v5)
    }
    var v19 int8 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
