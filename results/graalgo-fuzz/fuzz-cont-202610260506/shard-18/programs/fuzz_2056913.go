package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint32 }

func main() {
    var v1 string = "hello"
    var v2 bool = (-256 >= -1)
    var v3 uint32 = 65535
    v4 := make(map[string]bool)
    for fi5 := 0; fi5 < 2; fi5++ {
        if v2 {
            var v6 uint64 = 18446744073709551615
            var v7 float32 = 1.5
            _ = v6
            _ = v7
        }
        var v8 string = (v1 + (("" + v1) + "a"))
        var v9 int = 9223372036854775807
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 uint64 = 476
    var v11 uint64 = (v10 * (^(v10)))
    var v12 bool = (1000 != -1)
    for fi13 := 0; fi13 < 6; fi13++ {
        if true {
            v1 = ("" + "hello")
        } else {
            if v2 {
                v10 = 4294967295
            } else {
                v2 = (v12 || false)
                delete(v4, (v1 + v1))
            }
            v14 := GenS1{F0: 42, F1: v11, F2: v3}
            var v15 int8 = 19
            _ = v14
            _ = v15
        }
        var v16 int16 = 208
        var v17 float32 = 1.5
        _ = fi13
        _ = v16
        _ = v17
    }
    var v18 int16 = -818
    if v2 {
        v19 := []int{1, 9223372036854775807, -420, -1}
        v2 = (32767 >= 0)
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v3, v2)
        var v21 int16 = v18
        var v22 uint64 = v11
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = (^(v3))
    v18 = (-(v18))
    var v24 uint32 = (v23 * (v3 >> ((65535) & 63)))
    var v25 float64 = (1e10 * (-2.25 - 0.0))
    delete(v4, v1)
    fmt.Println(v12, v2)
    v26 := GenS1{F0: v18, F1: v10, F2: (v24 << ((512) & 63))}
    switch 127 {
    case 32768:
        _ = len(v4)
    case 42:
        v12 = false
    }
    delete(v4, "\n\t")
    fmt.Println(v23, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
