package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 int32 }

func main() {
    var v1 int64 = -874
    var v2 bool = false
    var v3 uint32 = 158
    var v4 bool = (!v2)
    var v5 uint16 = 1
    var v6 string = "hello"
    var v7 float32 = -2.25
    v8 := GenS1{F0: 860, F1: 43, F2: 804}
    v9 := GenS2{F0: v5, F1: 640, F2: 1}
    if (v6 != v6) {
        v2 = (754 != 291)
        fmt.Println(v6, v4, v2)
    }
    fmt.Println(v6, v6)
    v10 := GenS1{F0: 4294967295, F1: 278, F2: 1}
    if v4 {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int8 = 0
            fmt.Println(v4, v4, v2)
            _ = fi11
            _ = v12
        }
        var v13 uint8 = 161
        switch 128 {
        case -706:
            v14 := []int{}
            _ = v14
        case -714:
            var v15 bool = v4
            _ = v15
        }
        _ = v13
    } else {
        if (!true) {
            v6 = (("a" + "foo bar") + v6)
        }
    }
    v4 = (((v1 ^ v1) <= -9223372036854775808) || false)
    var v16 string = (v6 + v6)
    v17 := GenS2{F0: v5, F1: 65535, F2: 0}
    v4 = (395 != 1)
    v18 := []int{-9223372036854775808, 618}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
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
