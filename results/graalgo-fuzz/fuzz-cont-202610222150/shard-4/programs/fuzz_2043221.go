package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 int }
func genFn0(p0 uint32) int16 {
    if false {
        p0 -= p0
        if ((true && (0 < 32767)) || false) {
            var v7 int32 = -2147483648
            _ = v7
        }
        v8 := []int{677, 668, 32767}
        _ = v8
    } else {
        v9 := GenS1{F0: -904}
        fmt.Println(p0)
        _ = v9
    }
    v10 := []int{}
    p0 -= (p0 % (((p0 << ((65535) & 63))) | 1))
    _ = p0
    _ = v10
    return 671
}
func genFn1(p0 uint64, p1 string) int8 {
    var v11 int64 = -908
    p1 = ""
    fmt.Println(p1, p0, p1)
    fmt.Println(p0)
    var v12 bool = false
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 100
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    var v3 bool = (-9223372036854775808 > -631)
    v4 := make(map[string]bool)
    var v5 uint32 = (^(v1))
    v6 := []int{-1}
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 13; fi14++ {
        fmt.Println(v5, v13, v13)
        for fi15 := 0; fi15 < 11; fi15++ {
            v3 = (3.14159 <= 1.5)
            _ = len(v4)
            var v16 uint32 = v5
            _ = fi15
            _ = v16
        }
        if true {
            switch fi14 {
            case 32768:
                var v17 uint32 = 642
                _ = v17
            }
            for fi18 := 0; fi18 < 7; fi18++ {
                v19 := GenS2{F0: 858}
                _ = fi18
                _ = v19
            }
            v13 -= v13
        }
        _ = fi14
    }
    v3 = v3
    v20 := GenS3{F0: 32767, F1: 0, F2: -2147483648}
    if ("hello" > "") {
        _ = len(v6)
    }
    if v3 {
        var v21 float64 = v13
        _ = v21
    } else {
        var v22 int8 = 100
        _ = v22
    }
    var v23 uint64 = 4294967295
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v13)
        v25 := GenS1{F0: (fi24 >> ((42) & 63))}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
