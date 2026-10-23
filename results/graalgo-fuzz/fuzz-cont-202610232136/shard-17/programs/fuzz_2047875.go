package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 255
    v3 := make(map[uint32]string)
    v2 = v2
    if (1e10 != (1.5 - float64(float32(1e10)))) {
        v4 := []int{-450}
        if true {
            v1 = (v1 / ((v1) | 1))
            var v5 uint16 = 2
            _ = v5
        }
        delete(v3, 2147483647)
        _ = v4
    } else {
        fmt.Println(v1, v2)
        for fi6 := 0; fi6 < 15; fi6++ {
            v7 := []int{fi6}
            var v8 int8 = -128
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = len(v3)
    }
    var v9 float32 = (float32((3.14159 + 3.14159)) + 0.1)
    var v10 int8 = 42
    v11 := v3[42]
    switch 9223372036854775807 {
    case -114:
        v12 := GenS1{F0: 0}
        _ = v12
    case 32767:
        v13 := GenS2{F0: 106}
        _ = v13
    case 127:
        v14 := v3[157]
        _ = v14
    default:
        v10 = v10
    }
    var v15 int32 = 2147483647
    switch v11 {
    case "\n\t":
        v2 = v2
    case "hello":
        v11 = ((v11 + (v11 + ("hello" + ((v11 + "\n\t") + "\n\t")))) + "")
    default:
        v16 := v3[4294967295]
        _ = v16
    }
    v10 -= v10
    var v17 uint = v2
    var v18 uint = 18446744073709551615
    if false {
        var v19 int16 = 0
        _ = v19
    } else {
        v20 := []int{-2147483648}
        _ = v20
    }
    v21 := []int{9223372036854775807}
    switch v11 {
    case "hello":
        var v22 int64 = int64(v17)
        _ = v22
    case "a":
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int = (fi23 | fi23)
            v15 = (v15 * v15)
            _ = fi23
            _ = v24
        }
    case "foo bar":
        var v25 uint64 = 1
        _ = v25
    default:
        var v26 uint32 = 2
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
