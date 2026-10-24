package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int; F1 bool; F2 int8 }

func main() {
    var v1 uint64 = 924
    v2 := []int{}
    var v3 bool = (!(true && true))
    v4 := make(map[uint32]int)
    var v5 bool = (908 < 0)
    v3 = v3
    var v6 float64 = 0.1
    var v7 bool = false
    v8 := GenS1{F0: (v1 * 2), F1: ""}
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: 42, F1: "hello"}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 18; fi11++ {
        _ = len(v4)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int32 = 42
            var v14 float64 = v6
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    var v15 int32 = 2147483647
    var v16 float32 = 0.1
    v1 = (v1 - v1)
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v3, fi17, v16)
        fmt.Println(v16, v1, v3)
        var v18 uint64 = v1
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 1
    if (v16 < float32(0.1)) {
        v15 += (v15 * 42)
        _ = len(v2)
        var v20 bool = (-1 == 127)
        _ = v20
    } else {
        fmt.Println(v3)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v21 := GenS1{F0: v1, F1: ""}
        _ = v21
    }
    v22 := GenS2{F0: -150, F1: (768 <= 0), F2: 127}
    v7 = (v3 || true)
    v23 := GenS1{F0: uint64(v19), F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
