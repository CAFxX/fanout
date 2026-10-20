package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int16 }
type GenS2 struct { F0 bool; F1 int64 }
func genFn0() int {
    var v4 bool = false
    v5 := []int{4294967295, -134, 150, 32767}
    var v6 float32 = 0.0
    fmt.Println(v6, v6)
    var v7 uint16 = 2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() string {
    if false {
        var v8 bool = ((2 <= 2) && false)
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 string = "\n\t"
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        var v11 uint8 = 2
        v12 := []int{int(v11), 127, 1}
        _ = v11
        _ = v12
    }
    var v13 uint8 = 42
    v13 -= 33
    v13 -= v13
    _ = v13
    return "hello"
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 int = 2147483647
    v14 := GenS1{F0: 950, F1: v3, F2: 1000}
    var v15 int = 2147483647
    fmt.Println(v3)
    v16 := GenS1{F0: -2147483648, F1: v3, F2: -1}
    switch "a" {
    case "a":
        v15 -= genFn0()
    case "hello":
        v15 -= v3
    case "foo bar":
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint32 = 2
            fmt.Println(v3)
            v19 := GenS2{F0: true, F1: 9223372036854775807}
            _ = fi17
            _ = v18
            _ = v19
        }
    default:
        fmt.Println(v15)
    }
    v3 = v15
    fmt.Println(v15, v15)
    var v20 int = v3
    var v21 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
