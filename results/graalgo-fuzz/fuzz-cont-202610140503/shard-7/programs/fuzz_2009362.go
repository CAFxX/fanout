package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
type GenS4 struct { F0 int; F1 uint64; F2 int8 }

func main() {
    var v1 bool = (853 == 42)
    var v2 uint = 42
    var v3 int = 32768
    v4 := []int{v3}
    v5 := []uint32{443, 133, 2147483647}
    var v6 string = ("\n\t" + (("a" + (((((("hello" + "foo bar") + "") + "a") + ((("hello" + "") + ("\n\t" + "")) + "foo bar")) + "\n\t") + "hello")) + ((("hello" + "\n\t") + "foo bar") + "\n\t")))
    if (!false) {
        fmt.Println(v6, v2, v2)
    }
    v6 = v6
    v7 := GenS1{F0: true, F1: 255, F2: v6}
    if (39 > 732) {
        v6 = v6
        var v8 int64 = 1
        var v9 float64 = 1.5
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            v6 = ("\n\t" + v6)
            fmt.Println(v1, v2)
            _ = fi10
        }
        v11 := GenS2{F0: (!false), F1: 42}
        _ = v11
    }
    v12 := GenS1{F0: false, F1: 42, F2: v6}
    switch v3 {
    case 198:
        v5 = append(v5, 0)
    default:
        if v1 {
            v13 := GenS3{F0: v6, F1: 0, F2: (false || v1)}
            _ = v13
        }
    }
    fmt.Println(v6)
    v14 := GenS4{F0: v3, F1: 1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
