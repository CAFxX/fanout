package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0() bool {
    var v5 uint = 2
    var v6 float32 = (1e10 + 3.14159)
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 uint64, p1 uint16) uint8 {
    if false {
        var v7 bool = false
        _ = v7
    } else {
        var v8 uint16 = (^(p1))
        _ = v8
    }
    var v9 uint64 = (p0 & p0)
    p0 += 759
    _ = p0
    _ = p1
    _ = v9
    return 255
}

func main() {
    var v1 int16 = 1
    var v2 bool = (65535 <= 65535)
    v3 := make(map[uint32]uint32)
    v4 := []string{}
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 int32 = -1
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float64 = (1e10 * 0.0)
        var v14 uint8 = 127
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 = false
    var v15 int8 = 1
    if v2 {
        v16 := GenS1{F0: 32768, F1: "hello"}
        _ = v16
    }
    v17 := GenS1{F0: -298, F1: "\n\t"}
    v18 := v3[2]
    v2 = (59 >= (v18 >> ((65535) & 63)))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v19 int32 = 2147483647
    var v20 int16 = v1
    v18 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
