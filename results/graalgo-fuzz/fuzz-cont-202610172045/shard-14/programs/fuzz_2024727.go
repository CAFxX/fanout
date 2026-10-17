package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }

func main() {
    var v1 float64 = (float64(3.14159) / 1e10)
    var v2 string = ""
    var v3 int16 = 42
    v4 := make(map[int]int)
    v5 := make(map[int]bool)
    if true {
        v1 = v1
    }
    delete(v4, -893)
    fmt.Println(v3, v3, v1)
    if false {
        _ = len(v4)
    }
    switch 762 {
    case -9223372036854775808:
        var v6 float32 = ((0.0 * (1e10 * 3.14159)) * float32(((0.0 / v1) + (v1 + (v1 + (v1 - v1))))))
        _ = v6
    case 128:
        v7 := GenS1{F0: -67, F1: v3, F2: 0}
        _ = v7
    case -703:
        v2 = ((v2 + "\n\t") + (("a" + v2) + ((v2 + (((v2 + ("hello" + v2)) + v2) + "\n\t")) + "hello")))
    }
    v8 := []int{2147483647, -1, 128}
    v9 := GenS1{F0: 66, F1: 0, F2: 256}
    var v10 int32 = 0
    var v11 float64 = 0.1
    v12 := GenS1{F0: 0, F1: v3, F2: 65535}
    v3 = 689
    for fi13 := 0; fi13 < 13; fi13++ {
        v11 = (v11 * (float64(3.14159) - float64(float32(v11))))
        fmt.Println(v1, v2, v10)
        _ = fi13
    }
    v14 := GenS1{F0: -128, F1: v3, F2: 2147483648}
    v3 -= (v3 << ((65535) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
