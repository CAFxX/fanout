package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    var v4 uint64 = 18446744073709551615
    var v5 string = "\n\t"
    v6 := []uint32{65535, 65535}
    var v7 uint = 114
    v8 := GenS1{F0: 107, F1: true}
    v5 = v5
    fmt.Println(v7, v4)
    switch -1 {
    case 127:
        if len(v6) > 0 {
            v6[0] = 93
        }
    case 2147483647:
        for fi9 := 0; fi9 < 3; fi9++ {
            v5 = "a"
            for fi10 := 0; fi10 < 4; fi10++ {
                v6 = append(v6, 2)
                v11 := GenS1{F0: 100, F1: (!((((0.0 + (3.14159 - (-2.25 - 1e10))) + 1e10) == (1.5 / (0.1 + float64(1.5)))) && (false && true)))}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    default:
        switch 128 {
        case 289:
            v7 += v7
        case -659:
            var v12 bool = (false && true)
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int = (-(fi13))
        if false {
            var v15 bool = false
            _ = v15
        } else {
            switch v14 {
            case 9223372036854775807:
                var v16 uint32 = 4294967295
                _ = v16
            }
            v17 := GenS2{F0: 32, F1: (v5 + (v5 + (v5 + "foo bar")))}
            fmt.Println(v4, v14)
            _ = v17
        }
        for fi18 := 0; fi18 < 15; fi18++ {
            v3[200] = 0
            var v19 int = (^(fi13))
            var v20 uint16 = 103
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi13
        _ = v14
    }
    v7 = v7
    var v21 uint32 = 920
    var v22 float32 = 0.1
    fmt.Println(v4, v4)
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 bool = (v21 > v21)
        _ = fi23
        _ = v24
    }
    if true {
        if true {
            var v25 string = "foo bar"
            _ = len(v6)
            var v26 float32 = (v22 * 0.0)
            _ = v25
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
