package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint32; F1 int8; F2 uint }
type GenS4 struct { F0 int; F1 uint64; F2 uint32 }
func genFn0(p0 int16, p1 uint) int8 {
    var v7 int64 = -1
    for fi8 := 0; fi8 < 13; fi8++ {
        fmt.Println(v7)
        var v9 int32 = -534
        _ = fi8
        _ = v9
    }
    switch "foo bar" {
    case "a":
        var v10 float32 = 0.1
        _ = v10
    case "hello":
        p0 -= p0
    case "":
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float64 = 0.1
        _ = fi12
        _ = v13
    }
    v14 := []int{-2147483648, 0, 1}
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 73
}

func main() {
    var v1 string = ""
    v2 := []int{-2147483648, 2147483647}
    var v3 int8 = 42
    v4 := make(map[int]bool)
    v5 := []uint32{839, 2147483647}
    var v6 int = 2147483647
    v6 = 388
    v15 := GenS1{F0: -42, F1: 42, F2: 2}
    for fi16 := 0; fi16 < 13; fi16++ {
        v6 = -119
        _ = fi16
    }
    fmt.Println(v1)
    v17 := GenS2{F0: (v1 + ("\n\t" + "")), F1: 1}
    var v18 float64 = 0.0
    if true {
        v19 := GenS3{F0: 851, F1: 1, F2: 42}
        v20 := GenS4{F0: (v6 << ((65535) & 63)), F1: 4294967295, F2: 0}
        fmt.Println(v18, v1)
        _ = v19
        _ = v20
    } else {
        v1 = v1
        v21 := GenS1{F0: 453, F1: -1, F2: 1}
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = v21
    }
    fmt.Println(v18, v1)
    var v22 string = v1
    var v23 int8 = v3
    v24 := v4[(v6 & (v6 >> ((2) & 63)))]
    var v25 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
