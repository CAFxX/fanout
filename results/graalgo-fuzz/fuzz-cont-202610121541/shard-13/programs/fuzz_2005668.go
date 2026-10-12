package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }
func genFn0(p0 string, p1 uint16) uint16 {
    var v7 float64 = 3.14159
    if true {
        var v8 int16 = 421
        fmt.Println(v8, v8, v7)
        fmt.Println(v7)
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            v10 := GenS1{F0: p0, F1: 486, F2: 112}
            _ = v10
        }
        v11 := []int{}
        var v12 int = fi9
        _ = fi9
        _ = v11
        _ = v12
    }
    var v13 uint32 = 2147483648
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return p1
}

func main() {
    var v1 uint16 = 2
    var v2 float32 = (0.0 - 0.0)
    var v3 uint = 18446744073709551615
    v4 := make(map[int]bool)
    v5 := []uint32{}
    v6 := make(map[uint32]uint32)
    var v14 int32 = 1
    fmt.Println(v2)
    v6[2147483647] = 0
    v1 = (v1 >> (((v3 + v3)) & 63))
    if len(v5) > 0 {
        v5[0] = 65535
    }
    if (true || (272 > 798)) {
        var v15 uint32 = 2147483647
        v3 -= v3
        _ = v15
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            v17 := GenS1{F0: ("foo bar" + "a"), F1: 42, F2: v14}
            _ = fi16
            _ = v17
        }
    }
    var v18 float64 = (3.14159 + (1e10 * 1e10))
    v19 := GenS2{F0: 9223372036854775807, F1: ("a" + "a")}
    v20 := GenS1{F0: "\n\t", F1: 0, F2: (v14 << (((^(v3))) & 63))}
    if ((v18 * v18) <= -2.25) {
        if (!(true && false)) {
            var v21 bool = true
            var v22 float64 = 0.0
            _ = v21
            _ = v22
        } else {
            var v23 int64 = -429
            _ = v23
        }
        v24 := GenS3{F0: (!true), F1: 0, F2: -659}
        _ = v24
    }
    if true {
        if true {
            v25 := GenS3{F0: true, F1: int8(v1), F2: 9223372036854775807}
            var v26 string = ""
            _ = v25
            _ = v26
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if len(v5) > 0 {
            v5[0] = 285
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
