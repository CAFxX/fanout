package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 int16) bool {
    v5 := []int{128, -984, 32767}
    fmt.Println(p0)
    v6 := GenS1{F0: (p0 | 1000), F1: 566}
    v5 = append(v5, 441)
    _ = p0
    _ = v5
    _ = v6
    return false
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 int8 = -95
    v2[("a" + ("a" + (("hello" + "a") + (("" + "a") + "foo bar"))))] = 1
    var v7 int = 32767
    delete(v2, ("a" + ("a" + "")))
    switch (v7 ^ v7) {
    case 1:
        var v8 int32 = 962
        _ = v8
    default:
        v9 := v2[""]
        _ = v9
    }
    var v10 bool = (1 != -9223372036854775808)
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint32 = 4294967295
        v13 := GenS1{F0: 42, F1: 255}
        v14 := []int{(v7 / ((-1) | 1))}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v4 += v4
    v15 := GenS2{F0: 0}
    var v16 float64 = 0.1
    if (v1 == 3.14159) {
        var v17 int8 = 1
        v18 := GenS1{F0: 1000, F1: 580}
        _ = v17
        _ = v18
    } else {
        var v19 int16 = -32768
        _ = v19
    }
    v20 := GenS3{F0: 557, F1: 42, F2: 2}
    if (!(v7 != v7)) {
        var v21 float32 = -2.25
        var v22 bool = (v10 && true)
        switch "foo bar" {
        case "a":
            v23 := GenS1{F0: 1000, F1: 0}
            _ = v23
        }
        _ = v21
        _ = v22
    }
    v24 := GenS4{F0: 42, F1: true}
    v4 = v4
    v25 := []int{v7, v7, v7}
    for fi26 := 0; fi26 < 14; fi26++ {
        v10 = true
        _ = fi26
    }
    var v27 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
