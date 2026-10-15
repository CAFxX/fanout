package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 bool; F2 int64 }
type GenS3 struct { F0 int; F1 int64; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int8; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint }
func genFn0() int16 {
    var v6 string = ""
    var v7 int16 = 1000
    switch ((v6 + v6) + "foo bar") {
    case "\n\t":
        v8 := GenS1{F0: 156}
        _ = v8
    }
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int32, p1 uint64, p2 uint64) string {
    v9 := []int{-2147483648}
    var v10 int16 = -32768
    var v11 int32 = (p0 << ((uint(p2)) & 63))
    fmt.Println(v11, p0, v11)
    v10 = (v10 | 1000)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return "hello"
}

func main() {
    v1 := []uint32{2, 2147483647}
    var v2 uint8 = 134
    v3 := make(map[uint32]int)
    var v4 uint = 0
    var v5 uint64 = 2
    _ = len(v1)
    if (!(((1.5 - 3.14159) >= 3.14159) && true)) {
        var v12 int = 983
        _ = v12
    } else {
        if true {
            var v13 uint32 = 216
            if len(v1) > 0 {
                v1[0] = (v13 + 2147483647)
            }
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: -392}
                _ = fi14
                _ = v15
            }
            _ = v13
        }
        var v16 int32 = -1
        _ = v16
    }
    v4 += (v4 | 18446744073709551615)
    v17 := v3[2147483648]
    for fi18 := 0; fi18 < 7; fi18++ {
        switch "a" {
        case "hello":
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 uint32 = 255
                _ = fi19
                _ = v20
            }
        case "\n\t":
            var v21 float32 = 1.5
            _ = v21
        }
        v2 = 16
        _ = fi18
    }
    var v22 int8 = -1
    if false {
        var v23 int32 = -2147483648
        v24 := GenS2{F0: 375, F1: true, F2: 0}
        v25 := GenS3{F0: 9223372036854775807, F1: -9223372036854775808, F2: 9223372036854775807}
        _ = v23
        _ = v24
        _ = v25
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            v27 := GenS1{F0: 935}
            if false {
                var v28 float64 = 1.5
                v22 += v22
                _ = v28
            }
            v29 := GenS4{F0: false}
            _ = fi26
            _ = v27
            _ = v29
        }
        switch "foo bar" {
        case "\n\t":
            v30 := GenS5{F0: 4294967295, F1: (v22 & 68), F2: 0}
            _ = v30
        case "foo bar":
            fmt.Println(v17, v4, v17)
        }
        fmt.Println(v5)
    }
    v31 := GenS6{F0: 32767}
    v17 += -1
    v32 := GenS7{F0: v4}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
