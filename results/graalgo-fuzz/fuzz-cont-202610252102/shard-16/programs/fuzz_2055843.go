package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
func genFn0(p0 uint) uint32 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int8 = 100
        var v8 int = (fi6 / ((32767) | 1))
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint16 = 32767
    v9 -= 623
    var v10 float32 = (1e10 + ((1.5 - 3.14159) + 1e10))
    _ = p0
    _ = v10
    _ = v9
    return uint32(p0)
}

func main() {
    var v1 string = ("foo bar" + (("a" + "\n\t") + ""))
    var v2 int64 = 42
    v3 := []uint32{}
    var v4 uint32 = 0
    v5 := []uint32{2147483647, v4, v4, v4}
    var v11 uint16 = uint16(v4)
    v12 := GenS1{F0: 18446744073709551615, F1: 4294967295}
    for fi13 := 0; fi13 < 3; fi13++ {
        v2 = (-(v2))
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    if true {
        switch v1 {
        case "a":
            fmt.Println(v11, v1, v11)
        case "hello":
            fmt.Println(v2, v11, v1)
        case "\n\t":
            v5 = append(v5, 2)
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        if len(v5) > 0 {
            v5[0] = (v4 + v4)
        }
    }
    var v15 string = (("a" + "\n\t") + "hello")
    v4 += genFn0(18446744073709551615)
    v4 = (v4 % ((v4) | 1))
    if false {
        v1 = v15
        fmt.Println(v2, v15)
        v4 = 42
    }
    if (255 >= 0) {
        v15 = (v1 + v1)
        fmt.Println(v11)
    } else {
        _ = len(v3)
        var v16 int32 = -213
        var v17 bool = false
        _ = v16
        _ = v17
    }
    v11 = uint16(v4)
    v11 -= v11
    fmt.Println(v15, v2, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
