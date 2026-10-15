package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint32 }
func genFn0(p0 int8, p1 bool) uint {
    var v6 string = "a"
    if (!p1) {
        var v7 uint32 = 2147483648
        fmt.Println(p0, p0, v6)
        _ = v7
    } else {
        p0 = p0
    }
    v8 := GenS1{F0: 32767, F1: 0, F2: 423}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 string = "hello"
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 1000
    var v3 float64 = 0.1
    var v4 uint16 = 255
    var v5 int = -401
    v3 += v3
    v11 := GenS1{F0: v4, F1: -286, F2: 65535}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint32 = uint32(v4)
        v14 := []int{(v5 ^ 572)}
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 = v2
    v15 := []int{v5, v5, (v5 & 32767)}
    v15 = append(v15, v5)
    var v16 int = (v5 ^ 42)
    var v17 int8 = 6
    var v18 string = ("\n\t" + "a")
    var v19 float64 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
