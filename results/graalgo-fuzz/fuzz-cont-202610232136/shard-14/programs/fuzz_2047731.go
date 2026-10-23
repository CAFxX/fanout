package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    v3 := make(map[int]uint32)
    var v4 int8 = -128
    v5 := make(map[string]int)
    var v6 int64 = 1
    fmt.Println(v2, v6)
    v7 := GenS1{F0: -32768, F1: 87}
    switch 672 {
    case 9223372036854775807:
        var v8 string = "hello"
        _ = v8
    default:
        for fi9 := 0; fi9 < 11; fi9++ {
            if (!(42 <= 262)) {
                delete(v1, fi9)
                fmt.Println(fi9, v2, v2)
                v2 = v2
            }
            v2 = (!v2)
            _ = fi9
        }
    }
    var v10 uint64 = 1
    v11 := []int{-9223372036854775808, 0, -1}
    var v12 int64 = v6
    for fi13 := 0; fi13 < 13; fi13++ {
        if v2 {
            var v14 int = (fi13 - (fi13 % ((fi13) | 1)))
            v15 := GenS2{F0: (!v2), F1: 65535, F2: 128}
            for fi16 := 0; fi16 < 1; fi16++ {
                v17 := GenS1{F0: 32767, F1: (fi13 + fi16)}
                v5["a"] = (fi13 - fi16)
                v18 := GenS2{F0: ((0 < 0) || (1000 > -725)), F1: 65535, F2: 0}
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = v14
            _ = v15
        } else {
            var v19 uint8 = 128
            v20 := GenS1{F0: -213, F1: (-(fi13))}
            _ = v19
            _ = v20
        }
        var v21 uint = 65535
        _ = fi13
        _ = v21
    }
    v6 -= -1
    var v22 bool = v2
    var v23 int16 = 32767
    for fi24 := 0; fi24 < 10; fi24++ {
        delete(v3, (fi24 ^ -184))
        v25 := GenS1{F0: -885, F1: 127}
        _ = fi24
        _ = v25
    }
    fmt.Println(v10)
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 uint32 = 2
        _ = fi26
        _ = v27
    }
    var v28 int32 = 1
    v10 -= 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
