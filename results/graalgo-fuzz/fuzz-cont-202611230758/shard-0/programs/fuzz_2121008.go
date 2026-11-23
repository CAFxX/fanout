package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 uint16) uint64 {
    switch "hello" {
    case "hello":
        if ((2147483647 >= 2) || false) {
            for fi6 := 0; fi6 < 10; fi6++ {
                fmt.Println(fi6)
                var v7 int = (fi6 / ((fi6) | 1))
                _ = fi6
                _ = v7
            }
        } else {
            v8 := GenS1{F0: 42, F1: 4294967295}
            fmt.Println(p0)
            _ = v8
        }
    case "\n\t":
        fmt.Println(p0)
    case "":
        var v9 string = ("" + "a")
        _ = v9
    default:
        v10 := GenS2{F0: 100, F1: 252, F2: -1}
        _ = v10
    }
    var v11 int32 = -1
    for fi12 := 0; fi12 < 13; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v14 := GenS3{F0: 1000, F1: 42}
            _ = fi13
            _ = v14
        }
        var v15 int32 = 430
        switch "a" {
        case "foo bar":
            switch (fi12 & (fi12 >> ((1) & 63))) {
            case 0:
                var v16 int64 = -152
                _ = v16
            }
        default:
            switch ("foo bar" + "") {
            case "\n\t":
                var v17 uint32 = 120
                _ = v17
            case "hello":
                var v18 uint32 = 65535
                _ = v18
            }
        }
        _ = fi12
        _ = v15
    }
    _ = p0
    _ = v11
    return 42
}
func genFn1(p0 uint8) int64 {
    p0 = p0
    p0 -= (p0 / ((p0) | 1))
    _ = p0
    return 46
}

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 127
    var v3 uint = (v1 + 417)
    v4 := make(map[int]uint32)
    var v5 int16 = 1
    v19 := GenS4{F0: 2}
    var v20 bool = (float32(float64(((-2.25 - (3.14159 / 3.14159)) + 1e10))) <= (float32(1e10) * float32(1.5)))
    var v21 int16 = (v5 - 0)
    if v20 {
        var v22 uint32 = 2
        _ = v22
    }
    switch ("hello" + "") {
    case "a":
        fmt.Println(v2)
    }
    v23 := []int{-1, -840}
    v24 := GenS4{F0: 1}
    if len(v23) > 0 {
        v23[0] = 42
    }
    v2 = v2
    var v25 int = 416
    var v26 float32 = 1e10
    v23 = append(v23, -2147483648)
    switch v25 {
    case 366:
        for fi27 := 0; fi27 < 1; fi27++ {
            v2 = v2
            _ = fi27
        }
    case 42:
        var v28 int16 = v21
        _ = v28
    case -9223372036854775808:
        var v29 int8 = v2
        _ = v29
    }
    v4[9223372036854775807] = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
