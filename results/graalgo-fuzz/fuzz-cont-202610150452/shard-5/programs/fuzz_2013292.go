package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    if true {
        v4 := v3[(v1 + "\n\t")]
        v4 = v4
        _ = v4
    } else {
        v5 := []int{-359}
        v6 := GenS1{F0: 127, F1: 42}
        _ = v5
        _ = v6
    }
    v3["\n\t"] = "hello"
    if (false && true) {
        v7 := GenS1{F0: 127, F1: 0}
        v1 = "hello"
        var v8 string = v1
        _ = v7
        _ = v8
    }
    var v9 uint16 = 516
    if true {
        for fi10 := 0; fi10 < 6; fi10++ {
            v9 -= (v9 | v9)
            _ = fi10
        }
    } else {
        var v11 int = 128
        v12 := GenS1{F0: 26, F1: 1}
        _ = v11
        _ = v12
    }
    var v13 int = 9223372036854775807
    if (!true) {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 0, F1: 127}
            _ = fi14
            _ = v15
        }
        delete(v2, 0)
    } else {
        fmt.Println(v1, v13, v13)
        var v16 float64 = 1.5
        _ = v16
    }
    v17 := GenS2{F0: 42, F1: 364, F2: 87}
    fmt.Println(v1, v1, v9)
    switch v13 {
    case -2147483648:
        v1 = (((v1 + "") + v1) + v1)
    case 728:
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := v3[v1]
            _ = fi18
            _ = v19
        }
    case 32768:
        v20 := []int{(v13 & v13)}
        _ = v20
    default:
        switch v13 {
        case -818:
            v21 := []int{}
            _ = v21
        default:
            fmt.Println(v9, v9)
        }
    }
    fmt.Println(v9, v1, v13)
    var v22 int = 927
    v13 -= v22
    if false {
        v23 := GenS3{F0: 42}
        v24 := GenS4{F0: (-(v13)), F1: -646, F2: -128}
        v25 := GenS3{F0: 9223372036854775807}
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v9
}
