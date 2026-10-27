package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 4294967295
    var v3 int32 = 42
    if false {
        var v4 float64 = float64((0.0 + 0.1))
        _ = v4
    }
    v3 += -958
    v2 += (v2 | v2)
    for fi5 := 0; fi5 < 1; fi5++ {
        v6 := []int{fi5, fi5}
        _ = fi5
        _ = v6
    }
    var v7 bool = false
    v8 := GenS1{F0: 42, F1: -1}
    if v7 {
        var v9 float32 = 0.0
        for fi10 := 0; fi10 < 2; fi10++ {
            switch (^(fi10)) {
            case 1:
                var v11 uint32 = (^(v1))
                _ = v11
            case 127:
                var v12 float32 = 0.0
                _ = v12
            case -267:
                var v13 float32 = 1e10
                _ = v13
            default:
                v1 = 211
            }
            var v14 uint8 = 157
            _ = fi10
            _ = v14
        }
        _ = v9
    }
    fmt.Println(v2)
    v15 := GenS1{F0: 275, F1: 42}
    var v16 uint16 = 0
    v7 = ((!v7) || false)
    var v17 int32 = (^(v3))
    if v7 {
        v18 := GenS1{F0: v16, F1: 1}
        v19 := []int{-254, 853, 0}
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
