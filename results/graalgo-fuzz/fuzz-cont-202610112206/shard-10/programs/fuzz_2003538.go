package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint; F2 uint64 }
func genFn0(p0 uint32) int {
    v7 := []int{}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    var v10 int16 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        switch "a" {
        case "\n\t":
            fmt.Println(fi11, fi11, fi11)
        case "a":
            var v12 uint = 1
            _ = v12
        }
        _ = fi11
    }
    _ = p0
    _ = v10
    _ = v7
    return 32768
}
func genFn1(p0 int8) uint16 {
    var v13 float32 = 3.14159
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 uint = 18446744073709551615
        var v17 int32 = -2147483648
        if false {
            var v18 uint32 = 65535
            _ = v18
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    v19 := []int{genFn0(2147483647), 42, genFn0(2147483648)}
    for fi20 := 0; fi20 < 12; fi20++ {
        _ = len(v19)
        v21 := GenS2{F0: 38, F1: 4294967295, F2: 1}
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v13
    _ = v14
    _ = v19
    return 32767
}

func main() {
    var v1 int8 = -88
    v2 := []uint32{921, 65535}
    var v3 float32 = 0.0
    v4 := []int{}
    var v5 uint64 = 42
    var v6 bool = (2 > 1)
    var v22 bool = (!(648 != -944))
    var v23 int64 = 9223372036854775807
    v22 = v6
    var v24 int = 127
    v25 := GenS1{F0: v24}
    switch (v24 | -1) {
    case 359:
        fmt.Println(v3)
    case 869:
        var v26 bool = (v23 >= v23)
        _ = v26
    default:
        var v27 uint = 42
        _ = v27
    }
    fmt.Println(v5, v23)
    v6 = (v6 || ((1 < -2147483648) && (false && (0 >= 89))))
    _ = len(v2)
    var v28 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
