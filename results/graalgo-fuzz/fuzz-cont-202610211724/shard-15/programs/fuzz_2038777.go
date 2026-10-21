package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 uint16, p1 uint16, p2 uint32) string {
    var v5 bool = (true && (((-32768 > -32768) || true) || false))
    var v6 string = ""
    var v7 int16 = 42
    v7 = 326
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 int16 = -1
    var v2 int8 = -1
    var v3 string = "foo bar"
    var v4 uint = 686
    var v8 uint16 = 32767
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v8, v1, v8)
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v4, v2)
            v11 := []int{fi9, (fi10 * fi10), (fi10 * 32767)}
            v12 := GenS1{F0: 42, F1: -662}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 uint = 42
    var v14 int32 = 2147483647
    var v15 bool = true
    var v16 string = v3
    v17 := GenS1{F0: 28, F1: (v1 << ((243) & 63))}
    v18 := []int{9223372036854775807, 128, -1}
    var v19 uint8 = 128
    _ = len(v18)
    switch v16 {
    case "a":
        fmt.Println(v14, v14)
    case "hello":
        if len(v18) > 0 {
            v18[0] = 1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
