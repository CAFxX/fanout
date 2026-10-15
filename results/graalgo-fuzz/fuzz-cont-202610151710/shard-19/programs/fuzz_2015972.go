package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }

func main() {
    var v1 uint = 255
    var v2 float32 = float32(1e10)
    v3 := []uint32{42, 2147483648, 2}
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 int32 = 42
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        switch (("\n\t" + "hello") + ("" + "hello")) {
        case "":
            var v7 int64 = 2
            _ = v7
        case "hello":
            v8 := GenS1{F0: -157, F1: 1}
            _ = v8
        default:
            var v9 uint64 = 0
            _ = v9
        }
        _ = fi6
    }
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: -1, F1: -1}
        _ = fi10
        _ = v11
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v1 += 1
            var v14 int8 = 127
            _ = fi13
            _ = v14
        }
        v2 -= v2
        _ = fi12
    }
    var v15 int16 = 1000
    var v16 float64 = 0.0
    for fi17 := 0; fi17 < 8; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint = 0
            _ = fi18
            _ = v19
        }
        if true {
            var v20 int16 = (^(v15))
            v21 := GenS1{F0: v20, F1: 53}
            v22 := GenS1{F0: (v15 & v15), F1: 42}
            _ = v20
            _ = v21
            _ = v22
        }
        var v23 int8 = 107
        _ = fi17
        _ = v23
    }
    var v24 float32 = 0.1
    var v25 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
