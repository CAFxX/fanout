package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 string; F1 uint32; F2 uint16 }

func main() {
    var v1 uint = 42
    var v2 float32 = 1.5
    v3 := []uint32{}
    var v4 bool = true
    var v5 float32 = (v2 / v2)
    var v6 int32 = 0
    if false {
        v7 := GenS1{F0: 255, F1: 0}
        var v8 string = (("hello" + "") + "")
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: "\n\t", F1: 312, F2: 32767}
    if v4 {
        var v10 int8 = -52
        for fi11 := 0; fi11 < 6; fi11++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi11
        }
        _ = v10
    } else {
        v12 := GenS1{F0: 1, F1: 54}
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 = v5
            v5 -= v2
            _ = fi13
        }
    }
    if (!v4) {
        var v14 int32 = 42
        _ = v14
    }
    var v15 uint8 = 124
    if v4 {
        var v16 int32 = v6
        fmt.Println(v5, v16, v16)
        _ = v16
    } else {
        v17 := GenS1{F0: (v15 ^ v15), F1: 36}
        var v18 int = 128
        _ = len(v3)
        _ = v17
        _ = v18
    }
    v1 -= v1
    v1 -= 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
