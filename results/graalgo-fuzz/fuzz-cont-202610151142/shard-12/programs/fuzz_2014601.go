package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 bool }
func genFn0() uint {
    if false {
        v7 := GenS1{F0: 42}
        var v8 int64 = 9223372036854775807
        _ = v7
        _ = v8
    } else {
        var v9 int8 = -102
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{(fi10 | fi10)}
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    var v12 uint64 = 0
    for fi13 := 0; fi13 < 17; fi13++ {
        switch 745 {
        case 588:
            var v14 int64 = 804
            _ = v14
        case -383:
            v12 += (v12 ^ v12)
        case -1:
            v15 := []int{(-(fi13)), fi13, -1}
            _ = v15
        default:
            var v16 int16 = 405
            _ = v16
        }
        if (100 == 112) {
            fmt.Println(fi13)
            var v17 uint16 = 32767
            _ = v17
        }
        v18 := GenS1{F0: -9223372036854775808}
        _ = fi13
        _ = v18
    }
    _ = v12
    return 455
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.1
    var v4 int32 = 1
    var v5 uint8 = 236
    var v6 int16 = -766
    if true {
        v19 := v2[uint32(v5)]
        _ = v19
    }
    var v20 float32 = v3
    for fi21 := 0; fi21 < 3; fi21++ {
        switch (("a" + "\n\t") + "foo bar") {
        case "":
            v22 := GenS2{F0: "a"}
            _ = v22
        }
        if false {
            v23 := GenS3{F0: true, F1: false}
            if ((v3 / v3) != ((v3 - float32(((0.1 / 1.5) - float64((v20 * (-2.25 / v3)))))) + (((1e10 - v3) - float32((0.0 * -2.25))) * 0.1))) {
                delete(v2, 2147483647)
            } else {
                var v24 int16 = v6
                _ = v24
            }
            _ = v23
        }
        var v25 float64 = 3.14159
        _ = fi21
        _ = v25
    }
    if false {
        fmt.Println(v5, v6)
    }
    var v26 uint8 = v5
    var v27 int = -1
    v5 = v5
    var v28 int = (v27 % (((v27 | (^(v27)))) | 1))
    switch "hello" {
    case "foo bar":
        v29 := []int{(v28 << ((255) & 63)), (v27 % ((753) | 1)), -138}
        _ = v29
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 uint8 = 63
            v4 = (v4 & 426)
            fmt.Println(v26, v6)
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
