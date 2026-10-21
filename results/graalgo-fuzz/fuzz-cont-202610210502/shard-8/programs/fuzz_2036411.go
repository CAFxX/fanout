package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
func genFn0(p0 int32, p1 int16) uint32 {
    p1 += (^(p1))
    var v5 int = 128
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    var v1 uint32 = 427
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 float32 = 0.1
    var v6 int = 32767
    v7 := GenS1{F0: (false && false), F1: 4294967295, F2: 0}
    v1 = v1
    switch (v6 ^ (v6 << ((2) & 63))) {
    case 731:
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    case 127:
        var v10 int8 = -4
        _ = v10
    case 1:
        v4 = v4
    }
    v11 := GenS1{F0: false, F1: 42, F2: 70}
    var v12 uint16 = 2
    var v13 int8 = -1
    var v14 uint16 = (v12 - v12)
    v15 := []int{-2147483648, v6}
    v13 += (v13 | v13)
    var v16 float64 = 1e10
    switch v6 {
    case -191:
        fmt.Println(v16)
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v15 = append(v15, fi17)
        if false {
            v12 = 65535
        }
        var v18 int64 = 886
        _ = fi17
        _ = v18
    }
    var v19 int16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
