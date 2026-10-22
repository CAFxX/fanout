package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8, p1 bool) int32 {
    p1 = p1
    p1 = (!(p1 && (-9223372036854775808 > 9223372036854775807)))
    p1 = (true || false)
    var v6 bool = p1
    _ = p0
    _ = p1
    _ = v6
    return -2147483648
}

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]int)
    v3 := []uint32{0, 42}
    var v4 int64 = -9223372036854775808
    var v5 string = "hello"
    for fi7 := 0; fi7 < 7; fi7++ {
        v4 += -566
        _ = fi7
    }
    if (v4 < 1) {
        fmt.Println(v5, v4, v4)
        var v8 int8 = 100
        _ = v8
    } else {
        var v9 uint8 = 0
        v1 = (!false)
        fmt.Println(v5, v5)
        _ = v9
    }
    fmt.Println(v1, v1, v1)
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(v5, fi10)
        _ = fi10
    }
    var v11 int16 = -306
    var v12 uint64 = 4294967295
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int = (fi13 / ((fi13) | 1))
        v15 := GenS1{F0: 196}
        fmt.Println(v14, v11, fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v4, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
