package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint8 }
func genFn0(p0 bool, p1 int) int16 {
    var v4 uint = 990
    switch (p1 * (p1 | (p1 * p1))) {
    case 930:
        var v5 int16 = 32767
        _ = v5
    case -549:
        var v6 int16 = 625
        _ = v6
    }
    fmt.Println(p1, p0, p0)
    if (862 > 9223372036854775807) {
        var v7 uint64 = 42
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return 409
}

func main() {
    var v1 int8 = -77
    var v2 uint16 = 1
    var v3 int = 32767
    v2 += v2
    v8 := []int{2147483647, v3, (v3 & -1)}
    if (1 < 46) {
        v9 := GenS1{F0: 42, F1: 9223372036854775807, F2: (v2 * 32767)}
        v10 := GenS1{F0: 1000, F1: 42, F2: (v2 + v2)}
        _ = v10
        _ = v9
    } else {
        if len(v8) > 0 {
            v8[0] = (-(v3))
        }
    }
    var v11 bool = (2147483648 >= 42)
    v11 = true
    switch "a" {
    case "foo bar":
        switch (v3 | -137) {
        case 37:
            v1 -= (v1 ^ (v1 ^ v1))
        }
    case "a":
        switch 42 {
        case -2147483648:
            switch ("hello" + ("a" + "\n\t")) {
            case "hello":
                _ = len(v8)
            case "a":
                v11 = (0 != 494)
            default:
                if v11 {
                    fmt.Println(v1, v3, v1)
                }
            }
        case 1:
            var v12 int8 = v1
            _ = v12
        case 127:
            v11 = v11
        }
    case "hello":
        v1 += (v1 - v1)
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch "hello" {
        case "a":
            v14 := GenS2{F0: 71, F1: 4294967295}
            _ = v14
        case "hello":
            v15 := GenS3{F0: v11, F1: 2}
            _ = v15
        }
        var v16 float64 = 1e10
        var v17 float64 = float64(3.14159)
        _ = fi13
        _ = v16
        _ = v17
    }
    if (v11 || v11) {
        for fi18 := 0; fi18 < 17; fi18++ {
            v8 = append(v8, fi18)
            v2 -= v2
            v19 := GenS4{F0: 2147483648, F1: (true && v11)}
            _ = fi18
            _ = v19
        }
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v2 = (v2 << ((2) & 63))
    v20 := GenS5{F0: 2}
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 uint32 = 42
        v22 = (v22 & 2147483647)
        var v23 uint8 = 235
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
