package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint32 }
func genFn0(p0 uint64, p1 uint16) int64 {
    v6 := GenS1{F0: false, F1: -2147483648, F2: 2}
    v7 := GenS1{F0: false, F1: -1, F2: 65535}
    p1 += (p1 | uint16(p0))
    var v8 uint16 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 int8, p2 int8) int {
    p2 = p1
    v9 := GenS1{F0: (("\n\t" + "hello") != "\n\t"), F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -503
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 uint = 2
    v4 := []string{}
    var v5 int8 = 1
    var v10 int = 24
    var v11 float32 = 0.1
    if (-2.25 <= v11) {
        var v12 float32 = (1e10 - (1e10 / 0.1))
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(fi13, v10)
        _ = len(v4)
        v1 = v1
        _ = fi13
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint = v3
        if false {
            v11 = v11
            if (int64(v1) <= genFn0(0, 2)) {
                var v16 int32 = v1
                fmt.Println(v5, v11)
                _ = v16
            }
            var v17 int8 = 42
            _ = v17
        }
        fmt.Println(v11, v5, fi14)
        _ = fi14
        _ = v15
    }
    v18 := GenS1{F0: true, F1: 2147483647, F2: 65535}
    fmt.Println(v5)
    var v19 float32 = v11
    var v20 float64 = (0.0 + 0.0)
    v20 += 1.5
    v5 -= 42
    var v21 uint8 = 1
    v11 += v19
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
