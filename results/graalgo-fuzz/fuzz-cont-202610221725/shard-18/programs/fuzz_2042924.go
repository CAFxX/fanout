package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int16; F1 uint8; F2 int32 }
func genFn0() int {
    var v6 float64 = 1.5
    var v7 uint = 430
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint16, p1 uint16) int16 {
    for fi8 := 0; fi8 < 14; fi8++ {
        v9 := GenS1{F0: 2, F1: 0}
        _ = fi8
        _ = v9
    }
    fmt.Println(p0, p0)
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            var v11 int32 = 2147483647
            v12 := GenS1{F0: 108, F1: 887}
            v11 -= v11
            _ = v11
            _ = v12
        }
        var v13 float64 = 3.14159
        _ = fi10
        _ = v13
    }
    _ = p0
    _ = p1
    return -965
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v4 uint64 = 1
    var v5 float64 = 0.0
    v14 := GenS2{F0: 248, F1: 443}
    v15 := []int{-582}
    v16 := GenS3{F0: (v1 - v1), F1: 128, F2: 42}
    var v17 int = -148
    v5 += v5
    v18 := GenS3{F0: (v1 ^ v1), F1: 1, F2: 616}
    for fi19 := 0; fi19 < 12; fi19++ {
        if true {
            v1 += (v1 | v1)
            v4 += v4
        }
        _ = fi19
    }
    v20 := GenS3{F0: (v1 >> ((129) & 63)), F1: 1, F2: 42}
    fmt.Println(v4, v4, v1)
    var v21 int = genFn0()
    switch (v21 % ((v17) | 1)) {
    case 42:
        fmt.Println(v17)
    case -9223372036854775808:
        switch "\n\t" {
        case "hello":
            v21 -= (^(v17))
        case "\n\t":
            fmt.Println(v21, v5)
        case "foo bar":
            v17 = 0
        }
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint64 = v4
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
