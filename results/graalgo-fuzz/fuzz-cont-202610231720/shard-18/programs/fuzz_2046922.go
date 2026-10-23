package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 int32, p2 int16) int {
    p2 = 42
    var v6 uint = 18446744073709551615
    fmt.Println(p2, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{800, 2147483647, 2}
    var v3 uint64 = 551
    v4 := []uint32{194}
    var v5 string = "hello"
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 float32 = (1.5 / -2.25)
        if true {
            v9 := GenS1{F0: 128, F1: 42, F2: -1}
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 bool = (true && true)
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 42}
    if false {
        switch -9223372036854775808 {
        case 127:
            v1[4294967295] = false
        case -9223372036854775808:
            v1[2] = (-481 == 1000)
        case -1:
            var v13 bool = false
            _ = v13
        default:
            var v14 int8 = -112
            _ = v14
        }
        var v15 int8 = -116
        var v16 int16 = 545
        _ = v15
        _ = v16
    } else {
        switch -9223372036854775808 {
        case 652:
            v5 = "hello"
        }
        if true {
            fmt.Println(v3, v3)
            v1[uint32(v3)] = (0.0 != 0.0)
        } else {
            var v17 uint32 = 1
            _ = v17
        }
        var v18 int16 = 1
        _ = v18
    }
    var v19 uint64 = (^(v3))
    var v20 float32 = -2.25
    v20 = v20
    if (false && ((0 <= 381) && true)) {
        delete(v1, 303)
    } else {
        v21 := GenS1{F0: 255, F1: 42, F2: 2147483647}
        v4 = append(v4, 1)
        _ = v21
    }
    var v22 int8 = 0
    fmt.Println(v20, v19, v22)
    v19 = v3
    for fi23 := 0; fi23 < 11; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            v19 -= v3
            var v25 float64 = (0.1 * 3.14159)
            _ = fi24
            _ = v25
        }
        var v26 int64 = -1
        _ = fi23
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
