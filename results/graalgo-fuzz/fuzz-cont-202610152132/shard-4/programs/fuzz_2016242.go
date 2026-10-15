package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint32; F2 uint32 }
func genFn0(p0 bool, p1 bool, p2 int16) uint64 {
    if p1 {
        v6 := GenS1{F0: (127 < 128)}
        _ = v6
    }
    var v7 bool = (float64((((0.0 + 0.1) * float32(1.5)) + 1.5)) >= (3.14159 + -2.25))
    switch ("" + "\n\t") {
    case "":
        if true {
            var v8 float64 = 3.14159
            var v9 uint64 = 18446744073709551615
            p1 = p0
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if p1 {
            p2 = p2
            var v11 int16 = 42
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}
func genFn1() int64 {
    var v12 int8 = 42
    v13 := GenS1{F0: (!(1e10 <= 0.1))}
    var v14 uint64 = 784
    fmt.Println(v14, v12)
    _ = v12
    _ = v13
    _ = v14
    return -1
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    v3 := []int{9223372036854775807, 1}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]uint32)
    v1 = (v1 + v1)
    var v15 uint8 = 255
    v16 := GenS2{F0: 100, F1: 1, F2: 59}
    var v17 uint8 = v15
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS2{F0: -1, F1: 2, F2: uint32(v1)}
        _ = fi18
        _ = v19
    }
    v2["foo bar"] = 0
    var v20 uint8 = (^(v15))
    switch 540 {
    case 127:
        delete(v5, 32768)
    case 738:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch ("a" + "foo bar") {
    case "":
        var v21 int32 = 2147483647
        _ = v21
    case "\n\t":
        var v22 int16 = 32767
        _ = v22
    case "hello":
        var v23 uint16 = 191
        _ = v23
    default:
        var v24 uint = 2
        _ = v24
    }
    var v25 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
