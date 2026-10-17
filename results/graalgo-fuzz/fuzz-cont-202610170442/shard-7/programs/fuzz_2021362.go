package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 bool }

func main() {
    var v1 int64 = -52
    var v2 uint16 = 42
    v3 := make(map[uint32]uint32)
    var v4 string = ("\n\t" + "")
    v5 := []int{42, -71}
    switch "\n\t" {
    case "a":
        if len(v5) > 0 {
            v5[0] = -734
        }
    case "\n\t":
        v4 = v4
    case "":
        v1 = 28
    }
    _ = len(v5)
    v6 := GenS1{F0: ((42 >= 940) || false), F1: 32767}
    v2 = 349
    var v7 int8 = 127
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            v2 -= (v2 & (v2 & v2))
            v2 = 65535
            _ = fi9
        }
        v3[0] = 0
        _ = fi8
    }
    var v10 uint32 = 2147483647
    var v11 bool = (((v10 & 2147483648) == v10) || true)
    var v12 float64 = 1.5
    if v11 {
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                v11 = ((int8(v1) >= (-(v7))) && (!(0 <= 886)))
                var v15 bool = v11
                _ = fi14
                _ = v15
            }
            v7 = -4
            _ = fi13
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: 685, F1: 171}
            _ = fi16
            _ = v17
        }
        fmt.Println(v12, v10, v7)
    } else {
        v18 := GenS1{F0: v11, F1: -1}
        _ = v18
    }
    var v19 int64 = (v1 - v1)
    var v20 uint8 = 220
    if false {
        v2 -= uint16(v20)
    } else {
        switch (v4 + (v4 + ("a" + v4))) {
        case "foo bar":
            if v11 {
                v12 = 1.5
                var v21 uint64 = 1
                switch -766 {
                case -169:
                    v22 := GenS3{F0: int16(v7), F1: 255}
                    _ = v22
                default:
                    var v23 bool = (v1 < v19)
                    _ = v23
                }
                _ = v21
            } else {
                v24 := GenS4{F0: (v7 >= 34)}
                var v25 uint = 18446744073709551615
                _ = v24
                _ = v25
            }
        case "":
            fmt.Println(v20, v1, v12)
        case "hello":
            v2 = (v2 ^ v2)
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            if ((0 > v19) && ((v7 << ((0) & 63)) <= (v7 >> ((0) & 63)))) {
                switch ("a" + v4) {
                case "hello":
                    v27 := v3[(v10 | uint32(v2))]
                    _ = v27
                case "a":
                    v28 := v3[825]
                    _ = v28
                case "\n\t":
                    var v29 uint8 = 127
                    _ = v29
                }
                v12 -= v12
                v2 -= (v2 & (^(v2)))
            } else {
                if v11 {
                    var v30 int64 = 0
                    _ = v30
                }
            }
            _ = fi26
        }
    }
    v11 = v11
    v19 = 463
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
