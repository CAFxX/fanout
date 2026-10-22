package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint64, p1 uint16) int {
    v6 := []int{}
    v7 := GenS1{F0: -732, F1: 2}
    var v8 bool = true
    var v9 uint64 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return -992
}
func genFn1(p0 int32, p1 uint16, p2 int64) uint {
    p0 = 315
    for fi10 := 0; fi10 < 15; fi10++ {
        p2 = -9223372036854775808
        _ = fi10
    }
    if false {
        p0 = 1
        var v11 float64 = float64(((1.5 / ((0.1 - (1e10 / 1e10)) * (-2.25 / (-2.25 + (1.5 / 1.0))))) / (1.5 + (-2.25 * (((1e10 * 0.0) * float32(float64(1e10))) - 1e10)))))
        v12 := GenS1{F0: -768, F1: 103}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 18; fi13++ {
            p2 = p2
            var v14 string = "hello"
            v15 := GenS1{F0: -32768, F1: 100}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 272
}

func main() {
    v1 := []uint32{2147483647, 2147483648, 2, 1}
    var v2 uint16 = 32767
    v3 := []uint32{1, 42, 2147483648, 42}
    var v4 string = "foo bar"
    v5 := make(map[int]uint32)
    v4 = v4
    if false {
        v16 := GenS2{F0: 2}
        _ = v16
    }
    var v17 uint = 239
    var v18 uint8 = 0
    fmt.Println(v4)
    v4 = "hello"
    v4 = (v4 + "\n\t")
    v19 := GenS2{F0: (v17 % (((v17 >> ((genFn1(42, (v2 ^ v2), 0)) & 63))) | 1))}
    var v20 int16 = 42
    fmt.Println(v18)
    switch v4 {
    case "hello":
        v21 := GenS3{F0: v4, F1: v20}
        _ = v21
    case "foo bar":
        v17 = v17
    }
    if (((3.14159 <= 1e10) && true) && (!(float64(1.5) != 0.1))) {
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := v5[fi22]
            v24 := GenS4{F0: -40}
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    _ = len(v3)
    var v25 bool = (true || false)
    v20 = (v20 << (((v17 >> ((v17) & 63))) & 63))
    fmt.Println(v2)
    for fi26 := 0; fi26 < 6; fi26++ {
        switch v4 {
        case "\n\t":
            var v27 int64 = 517
            _ = v27
        case "hello":
            for fi28 := 0; fi28 < 6; fi28++ {
                switch (v4 + (v4 + v4)) {
                case "a":
                    fmt.Println(v20, v25, fi26)
                case "foo bar":
                    var v29 uint16 = 945
                    _ = v29
                default:
                    fmt.Println(v17, fi28, v2)
                    var v30 uint8 = (^(v18))
                    var v31 float64 = 1e10
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
        default:
            v17 = (v17 | (^(v17)))
        }
        if true {
            var v32 string = v4
            var v33 string = v32
            _ = v32
            _ = v33
        }
        fmt.Println(v2)
        _ = fi26
    }
    v2 = (v2 - v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
