package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 bool; F1 string }
func genFn0() uint32 {
    var v7 uint16 = 0
    v7 -= (v7 / (((v7 | 2)) | 1))
    switch "a" {
    case "\n\t":
        v8 := []int{1, 893}
        _ = v8
    case "foo bar":
        fmt.Println(v7)
    case "":
        v9 := []int{128, 2147483647, 32768, 32767}
        _ = v9
    default:
        v7 = (^(v7))
    }
    v10 := GenS1{F0: v7, F1: 2, F2: 238}
    _ = v10
    _ = v7
    return 419
}

func main() {
    var v1 int64 = 1
    v2 := []string{"hello", "hello"}
    var v3 float32 = (((1e10 + (0.0 * 1e10)) + 0.0) / 1.0)
    var v4 int64 = v1
    var v5 int8 = 100
    v6 := []uint32{16}
    fmt.Println(v3)
    var v11 uint8 = 128
    fmt.Println(v3, v3, v1)
    for fi12 := 0; fi12 < 4; fi12++ {
        if false {
            v5 -= 0
            var v13 int = fi12
            _ = v13
        }
        _ = len(v2)
        v1 += (v4 ^ v1)
        _ = fi12
    }
    if true {
        v14 := GenS2{F0: true, F1: "foo bar"}
        var v15 bool = false
        _ = v14
        _ = v15
    }
    switch 128 {
    case 32767:
        for fi16 := 0; fi16 < 20; fi16++ {
            v1 += (v1 % ((v1) | 1))
            if (v5 != (-(v5))) {
                var v17 int16 = -444
                _ = v17
            }
            _ = fi16
        }
    case 2147483647:
        var v18 uint16 = uint16(v11)
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: true, F1: "a"}
        switch "a" {
        case "\n\t":
            v21 := GenS1{F0: 0, F1: 572, F2: (v4 + 0)}
            _ = v21
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 string = "a"
            var v24 uint16 = 0
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi19
        _ = v20
    }
    v25 := GenS1{F0: 32767, F1: 2147483647, F2: (v4 | v1)}
    for fi26 := 0; fi26 < 20; fi26++ {
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
        _ = fi26
    }
    var v27 uint = 2
    fmt.Println(v11, v4, v1)
    var v28 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
