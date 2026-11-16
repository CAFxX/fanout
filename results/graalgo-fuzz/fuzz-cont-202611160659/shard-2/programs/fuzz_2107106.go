package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0() uint32 {
    var v6 float32 = (3.14159 / 1.5)
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{fi7, fi7}
        _ = fi7
        _ = v8
    }
    fmt.Println(v6)
    _ = v6
    return 760
}

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 179
    var v3 uint8 = 128
    var v4 float32 = (0.0 * 0.0)
    var v5 int64 = 9223372036854775807
    for fi9 := 0; fi9 < 16; fi9++ {
        v1 += v1
        switch fi9 {
        case 127:
            for fi10 := 0; fi10 < 9; fi10++ {
                v11 := []int{-9223372036854775808, (fi10 * fi10)}
                _ = fi10
                _ = v11
            }
        }
        _ = fi9
    }
    var v12 int = 32768
    v4 = (v4 / (v4 * float32(float64(float32(v1)))))
    v13 := []int{}
    v14 := GenS1{F0: 128, F1: 82}
    fmt.Println(v5, v4, v12)
    if len(v13) > 0 {
        v13[0] = 779
    }
    v5 = v5
    var v15 uint8 = 246
    var v16 uint16 = 298
    if true {
        var v17 int8 = 100
        v12 += v12
        var v18 uint = 984
        _ = v17
        _ = v18
    } else {
        var v19 bool = (1 < 42)
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        _ = len(v13)
        fmt.Println(fi20, fi20)
        v15 += (v3 - 142)
        _ = fi20
    }
    var v21 uint8 = (v15 & v3)
    var v22 int = v12
    for fi23 := 0; fi23 < 4; fi23++ {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        v5 += -1
        _ = fi23
    }
    v24 := GenS1{F0: 255, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
