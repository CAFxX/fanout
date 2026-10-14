package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0(p0 bool) string {
    p0 = true
    fmt.Println(p0)
    p0 = (!true)
    _ = p0
    return "a"
}
func genFn1(p0 int8, p1 uint64, p2 uint64) int64 {
    var v7 uint64 = 30
    fmt.Println(p2, p0, p2)
    switch "a" {
    case "\n\t":
        if true {
            var v8 int32 = 2147483647
            var v9 int64 = -366
            fmt.Println(v8)
            _ = v8
            _ = v9
        }
    }
    var v10 int32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 bool = (!(0 == -1))
    v3 := []uint32{464, 2147483648, 2147483648, 65535}
    v4 := []uint32{2147483648, 1}
    var v5 string = ""
    var v6 uint32 = 0
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 int8 = -1
        if true {
            if (42 < 4294967295) {
                v12 += v12
            }
            v12 = (v12 % (((v12 >> ((4294967295) & 63))) | 1))
            v5 = (v5 + (((("hello" + "") + "") + v5) + "hello"))
        }
        v13 := GenS1{F0: 902, F1: -32768}
        _ = fi11
        _ = v12
        _ = v13
    }
    v6 = v6
    var v14 float64 = (1.5 / -2.25)
    if len(v3) > 0 {
        v3[0] = v6
    }
    var v15 uint8 = 2
    v16 := GenS1{F0: uint(v15), F1: -32768}
    var v17 int64 = 1
    delete(v1, "hello")
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS1{F0: 575, F1: 1000}
        var v20 float32 = (float32(-2.25) + (float32((0.0 + 0.1)) + (3.14159 / 1.5)))
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
