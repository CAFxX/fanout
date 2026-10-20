package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint64, p1 string) uint16 {
    if (!(0 < 42)) {
        fmt.Println(p1)
        fmt.Println(p1, p0)
    }
    var v7 int16 = 1000
    switch 9223372036854775807 {
    case 42:
        fmt.Println(v7, p1, p0)
    case 127:
        v8 := GenS1{F0: -9223372036854775808, F1: p0}
        _ = v8
    }
    switch 2147483647 {
    case 690:
        var v9 uint8 = 1
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 7; fi10++ {
            p1 = "hello"
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = -404
    var v3 float32 = float32(1.5)
    var v4 string = "a"
    var v5 float32 = v3
    var v6 float64 = 0.1
    for fi11 := 0; fi11 < 12; fi11++ {
        v2 += v2
        _ = fi11
    }
    v12 := []int{4294967295}
    fmt.Println(v2, v6, v6)
    var v13 int8 = 127
    v14 := GenS1{F0: 1, F1: 2}
    v13 -= v13
    v13 = 127
    switch v4 {
    case "a":
        switch 1 {
        case 0:
            v6 = 0.0
        }
    case "":
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(v6)
            v1[65535] = true
            fmt.Println(v2, v5, v4)
            _ = fi15
        }
    case "foo bar":
        v4 = "hello"
    }
    var v16 int32 = 74
    var v17 int64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
