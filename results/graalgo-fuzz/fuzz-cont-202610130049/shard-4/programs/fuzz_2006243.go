package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int32 }
func genFn0(p0 uint8, p1 int32, p2 string) int64 {
    v6 := []int{}
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: fi7, F1: 0, F2: (-(p1))}
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    p2 = (("\n\t" + ("" + "hello")) + p2)
    p1 = 616
    var v9 int8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return -256
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 1e10
    var v3 uint64 = 172
    v4 := make(map[int]int)
    var v5 int64 = 0
    v3 += v3
    delete(v1, 42)
    fmt.Println(v2, v3, v5)
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := []int{593, fi10, 4294967295, (fi10 - int(v5))}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(fi12)
        _ = fi12
    }
    v3 = (^(v3))
    v3 -= v3
    var v13 uint32 = 65535
    v14 := GenS1{F0: -468, F1: 1, F2: 503}
    var v15 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
