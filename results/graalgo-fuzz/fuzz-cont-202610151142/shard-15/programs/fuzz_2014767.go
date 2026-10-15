package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
func genFn0(p0 int, p1 int32, p2 int16) uint16 {
    p2 -= 718
    var v4 uint8 = 255
    fmt.Println(p1)
    var v5 uint64 = 540
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 85
}

func main() {
    v1 := []int{42, -2147483648, 32768}
    var v2 float64 = 0.0
    var v3 string = "hello"
    var v6 float64 = v2
    v3 = v3
    fmt.Println(v2)
    var v7 int64 = -1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint64 = 18446744073709551615
        _ = len(v1)
        _ = fi8
        _ = v9
    }
    switch v3 {
    case "\n\t":
        v7 = (v7 + v7)
    default:
        if len(v1) > 0 {
            v1[0] = 985
        }
    }
    if true {
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(fi10, v7)
            _ = fi10
        }
    } else {
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 float32 = float32(v2)
            var v13 int16 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 int32 = 42
    v2 -= (v6 - (((v2 + v6) * v6) + v6))
    for fi15 := 0; fi15 < 3; fi15++ {
        v1 = append(v1, fi15)
        v3 = (v3 + "foo bar")
        _ = fi15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if len(v1) > 0 {
            v1[0] = (fi16 & fi16)
        }
        v17 := GenS1{F0: (v7 & v7), F1: (v14 ^ (v14 | (v14 * -2147483648))), F2: 127}
        var v18 uint64 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
