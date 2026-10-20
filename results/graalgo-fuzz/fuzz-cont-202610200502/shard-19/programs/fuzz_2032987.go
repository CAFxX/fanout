package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }

func main() {
    var v1 int = -9223372036854775808
    v2 := []int{v1}
    var v3 int32 = -229
    var v4 float64 = (1e10 + 1.5)
    v5 := make(map[uint32]string)
    var v6 int32 = v3
    var v7 int8 = -128
    v6 = 2147483647
    v4 += v4
    v1 -= 9223372036854775807
    var v8 uint16 = 65535
    var v9 uint8 = 49
    var v10 float32 = (float32(0.0) / float32(float64(1.5)))
    if false {
        var v11 uint = 65535
        var v12 uint = (v11 ^ v11)
        var v13 uint32 = 235
        _ = v11
        _ = v12
        _ = v13
    }
    switch "" {
    case "a":
        v14 := GenS1{F0: 182, F1: ("a" + ((((("\n\t" + "\n\t") + ("" + "foo bar")) + "") + "") + "\n\t"))}
        _ = v14
    case "\n\t":
        var v15 uint = 255
        _ = v15
    case "":
        var v16 int16 = 151
        _ = v16
    }
    switch v1 {
    case -2147483648:
        v6 = (v3 + -892)
    default:
        var v17 int = v1
        _ = v17
    }
    var v18 float32 = (v10 - (1e10 / v10))
    fmt.Println(v10)
    for fi19 := 0; fi19 < 13; fi19++ {
        v10 += (v18 + v18)
        var v20 int8 = 13
        if false {
            var v21 bool = false
            var v22 bool = v21
            _ = v21
            _ = v22
        } else {
            v4 = ((0.0 * v4) / v4)
            for fi23 := 0; fi23 < 18; fi23++ {
                _ = len(v2)
                var v24 uint = 0
                var v25 float32 = -2.25
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        _ = fi19
        _ = v20
    }
    fmt.Println(v9, v6, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
