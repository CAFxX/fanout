package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 int, p2 int) int16 {
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := GenS1{F0: (fi5 % ((p1) | 1)), F1: 1000, F2: 127}
        var v7 bool = (410 >= 717)
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(p1, p2, p1)
    var v8 bool = (!false)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-966}
    var v3 uint8 = 195
    v4 := make(map[string]bool)
    if true {
        v2 = append(v2, 128)
    }
    v9 := GenS1{F0: 128, F1: 42, F2: 100}
    var v10 string = "foo bar"
    _ = len(v4)
    if false {
        if false {
            var v11 uint = 1
            v12 := GenS1{F0: 0, F1: -1, F2: 100}
            v2 = append(v2, 0)
            _ = v11
            _ = v12
        }
    } else {
        v3 += 42
        var v13 uint16 = 352
        _ = v13
    }
    v14 := GenS1{F0: 15, F1: 42, F2: 42}
    switch (v10 + v10) {
    case "foo bar":
        var v15 float32 = (1e10 - 0.1)
        _ = v15
    case "\n\t":
        v3 -= v3
    case "a":
        if true {
            _ = len(v1)
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 -= (v3 ^ 0)
                var v17 int32 = -810
                var v18 uint32 = 23
                _ = fi16
                _ = v17
                _ = v18
            }
        } else {
            switch "" {
            case "foo bar":
                v19 := GenS2{F0: -128}
                _ = v19
            default:
                var v20 uint = 1
                _ = v20
            }
        }
    }
    switch "hello" {
    case "":
        var v21 int32 = -1
        _ = v21
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 32767
        }
    default:
        var v22 uint8 = 225
        _ = v22
    }
    _ = len(v1)
    v3 -= (^(v3))
    var v23 float32 = 0.1
    var v24 float32 = 3.14159
    fmt.Println(v23, v24, v24)
    switch "hello" {
    case "hello":
        v24 -= ((v23 * v24) + v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
