package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() uint16 {
    if false {
        v4 := GenS1{F0: 42}
        var v5 string = "foo bar"
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: -362}
    var v7 uint32 = 198
    switch "" {
    case "hello":
        fmt.Println(v7)
    case "foo bar":
        var v8 int = 9223372036854775807
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        if ("" > "a") {
            var v10 uint8 = 1
            var v11 bool = ((true && ((v10 << ((uint(v10)) & 63)) != v10)) && (!(!false)))
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint16 {
    var v12 uint32 = 4294967295
    switch "foo bar" {
    case "hello":
        fmt.Println(v12)
    default:
        fmt.Println(v12)
    }
    _ = v12
    return 255
}

func main() {
    var v1 uint8 = 255
    var v2 bool = false
    var v3 int16 = 1000
    v13 := GenS1{F0: 1}
    var v14 int16 = v3
    v15 := GenS2{F0: 1, F1: true}
    v14 = (v3 + v3)
    var v16 float64 = 1e10
    var v17 int = 32767
    v18 := []int{}
    v17 = 100
    var v19 float32 = -2.25
    var v20 int = (v17 - v17)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
