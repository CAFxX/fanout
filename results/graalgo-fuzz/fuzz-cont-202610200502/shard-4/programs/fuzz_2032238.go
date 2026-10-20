package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint; F1 uint8; F2 int64 }
func genFn0(p0 uint64) uint {
    if true {
        var v4 int8 = 1
        v5 := GenS1{F0: false, F1: ""}
        _ = v4
        _ = v5
    }
    v6 := GenS1{F0: true, F1: "\n\t"}
    v7 := []int{}
    if false {
        fmt.Println(p0)
        p0 -= 1
        for fi8 := 0; fi8 < 15; fi8++ {
            var v9 float32 = (1.5 * 1.5)
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1() uint16 {
    if true {
        var v10 uint64 = 0
        _ = v10
    } else {
        var v11 uint16 = 962
        _ = v11
    }
    var v12 string = "\n\t"
    _ = v12
    return 65535
}

func main() {
    var v1 uint32 = 72
    v2 := []int{-743}
    v3 := []uint32{(^(v1)), v1, 456, v1}
    var v13 string = "foo bar"
    fmt.Println(v1, v1)
    if len(v3) > 0 {
        v3[0] = (v1 + 742)
    }
    fmt.Println(v13, v13)
    if (127 <= 0) {
        var v14 uint64 = 985
        var v15 float32 = (0.1 / ((3.14159 / (1.5 + 0.1)) * 3.14159))
        _ = v14
        _ = v15
    }
    v1 += (v1 >> ((genFn0(1)) & 63))
    v13 = "hello"
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint32 = v1
        var v18 uint8 = 1
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = append(v3, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v2)
    v13 = (v13 + (v13 + "foo bar"))
    var v19 uint32 = 2
    var v20 float32 = float32(0.0)
    v21 := GenS2{F0: 898, F1: 132, F2: -9223372036854775808}
    v22 := GenS2{F0: 0, F1: 127, F2: -308}
    for fi23 := 0; fi23 < 3; fi23++ {
        var v24 int64 = 0
        var v25 string = (v13 + v13)
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
