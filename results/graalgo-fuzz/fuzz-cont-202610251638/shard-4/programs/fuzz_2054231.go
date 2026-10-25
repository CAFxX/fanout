package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint8, p1 uint16) uint64 {
    if (!false) {
        var v6 int16 = 0
        p1 -= p1
        _ = v6
    } else {
        fmt.Println(p0, p1)
        v7 := GenS1{F0: p1}
        p0 = (^(p0))
        _ = v7
    }
    var v8 uint16 = 1
    var v9 bool = (v8 > p1)
    p0 = (^(p0))
    v10 := GenS1{F0: 281}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    var v3 uint8 = 255
    var v4 float64 = 3.14159
    var v5 float32 = 1e10
    if false {
        fmt.Println(v5, v4, v3)
    }
    v5 -= float32((v4 * v4))
    v2[9223372036854775807] = -591
    for fi11 := 0; fi11 < 10; fi11++ {
        v5 += ((0.0 - v5) / 1e10)
        var v12 int8 = -19
        _ = fi11
        _ = v12
    }
    var v13 float32 = (3.14159 - (-2.25 + 0.1))
    v14 := GenS1{F0: 1}
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 float32 = v1
        _ = fi15
        _ = v16
    }
    fmt.Println(v5, v5)
    fmt.Println(v4, v3, v5)
    switch (("\n\t" + "foo bar") + "\n\t") {
    case "hello":
        switch -347 {
        case 0:
            var v17 string = "\n\t"
            _ = v17
        default:
            var v18 int = 233
            _ = v18
        }
    case "a":
        v19 := []int{2147483647, -9223372036854775808, 1, -2147483648}
        _ = v19
    }
    v3 -= v3
    fmt.Println(v3, v3, v3)
    switch 32767 {
    case -1:
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := []int{(fi20 << ((65535) & 63)), fi20, (fi20 + 127)}
            fmt.Println(v1, fi20)
            _ = fi20
            _ = v21
        }
    case 32767:
        v22 := GenS1{F0: 1}
        _ = v22
    default:
        var v23 uint = 375
        _ = v23
    }
    v5 += v1
    v24 := []int{32767, -2147483648, -565}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
