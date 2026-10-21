package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{-1, 127, -1}
    var v2 float64 = 1e10
    var v3 uint32 = 503
    v4 := make(map[int]uint32)
    if true {
        var v5 int32 = 0
        _ = v5
    } else {
        v1 = append(v1, 0)
        fmt.Println(v3, v3)
    }
    var v6 float64 = 1e10
    fmt.Println(v3, v3, v3)
    var v7 int8 = 100
    var v8 int32 = -1
    v9 := GenS1{F0: (v3 & (^(v3))), F1: ("foo bar" + ""), F2: ("foo bar" < "hello")}
    var v10 float32 = ((float32(v6) / 3.14159) * (1.5 / (0.1 + float32(v2))))
    v6 += v6
    var v11 int32 = v8
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v7, v10, fi12)
        v2 = 0.0
        _ = len(v4)
        _ = fi12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch (fi13 << ((255) & 63)) {
        case 128:
            var v14 int = fi13
            _ = v14
        case 1:
            fmt.Println(fi13, v7, v3)
        default:
            var v15 uint16 = 1
            _ = v15
        }
        v10 = v10
        var v16 int16 = 0
        _ = fi13
        _ = v16
    }
    var v17 int32 = 0
    var v18 int8 = (v7 - v7)
    switch "" {
    case "foo bar":
        var v19 int64 = 546
        _ = v19
    case "\n\t":
        if false {
            var v20 int = 4294967295
            fmt.Println(v10)
            _ = v20
        }
    case "":
        for fi21 := 0; fi21 < 18; fi21++ {
            for fi22 := 0; fi22 < 13; fi22++ {
                if (0 >= (fi22 * fi22)) {
                    var v23 bool = true
                    v11 = (v17 & v8)
                    v7 = v7
                    _ = v23
                }
                _ = fi22
            }
            _ = fi21
        }
    default:
        if true {
            v18 = 100
            var v24 int8 = -128
            _ = v24
        }
    }
    v25 := GenS1{F0: (^(v3)), F1: ((("foo bar" + "\n\t") + "a") + (("a" + ("a" + "")) + "")), F2: false}
    if true {
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS2{F0: 345}
            _ = fi26
            _ = v27
        }
    }
    v2 = ((3.14159 + (-2.25 * 1.5)) - (v2 / v6))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
