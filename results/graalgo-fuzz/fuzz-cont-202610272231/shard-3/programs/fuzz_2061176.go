package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int; F1 bool; F2 int }
type GenS8 struct { F0 int64; F1 uint; F2 uint }
func genFn0() int16 {
    for fi4 := 0; fi4 < 9; fi4++ {
        switch "foo bar" {
        case "hello":
            v5 := GenS1{F0: 912, F1: 9223372036854775807}
            _ = v5
        case "":
            var v6 int32 = 42
            _ = v6
        case "\n\t":
            v7 := []int{(fi4 & (fi4 * -811)), fi4}
            _ = v7
        }
        _ = fi4
    }
    switch "foo bar" {
    case "\n\t":
        var v8 uint64 = 2
        _ = v8
    case "foo bar":
        var v9 int8 = 73
        _ = v9
    case "":
        var v10 uint16 = 32767
        _ = v10
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 uint = 4294967295
            _ = fi11
            _ = v12
        }
    }
    return 1000
}

func main() {
    v1 := []uint32{}
    v2 := []string{}
    v3 := make(map[uint32]uint32)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint16 = 65535
        if (fi13 > fi13) {
            if (!false) {
                switch -9223372036854775808 {
                case 267:
                    v15 := GenS2{F0: false, F1: 127, F2: 679}
                    _ = v15
                case 308:
                    v14 = v14
                default:
                    switch (fi13 ^ (fi13 & (fi13 & 2147483647))) {
                    case -9223372036854775808:
                        if (false || true) {
                            fmt.Println(fi13, fi13)
                        }
                    default:
                        for fi16 := 0; fi16 < 19; fi16++ {
                            v14 -= (v14 ^ v14)
                            _ = fi16
                        }
                    }
                }
            } else {
                v14 -= v14
                v14 -= 42
            }
            var v17 uint8 = 44
            v18 := GenS3{F0: 857, F1: (false && true), F2: 902}
            _ = v17
            _ = v18
        }
        _ = fi13
        _ = v14
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 float64 = (3.14159 - 1e10)
        v20 = (v20 * v20)
        _ = fi19
        _ = v20
    }
    var v21 int = -63
    v22 := GenS4{F0: 304}
    for fi23 := 0; fi23 < 15; fi23++ {
        if true {
            v24 := GenS2{F0: false, F1: 127, F2: 42}
            var v25 uint16 = 32767
            _ = v24
            _ = v25
        }
        v21 = (-(v21))
        _ = fi23
    }
    v21 = v21
    v3[896] = 65535
    fmt.Println(v21)
    var v26 uint = 18446744073709551615
    switch ("\n\t" + "hello") {
    case "foo bar":
        delete(v3, 4294967295)
    case "a":
        v21 = (v21 ^ (-(v21)))
    }
    var v27 int16 = 435
    switch v21 {
    case 1:
        for fi28 := 0; fi28 < 16; fi28++ {
            v29 := GenS5{F0: 1}
            _ = fi28
            _ = v29
        }
    case 0:
        v30 := GenS1{F0: 2147483648, F1: 42}
        _ = v30
    case 736:
        fmt.Println(v21, v27, v21)
    default:
        var v31 uint = (v26 ^ 42)
        _ = v31
    }
    if (42 > 100) {
        v32 := GenS6{F0: v26, F1: 117}
        _ = v32
    }
    v33 := GenS7{F0: v21, F1: false, F2: -493}
    fmt.Println(v26, v21, v21)
    v34 := GenS8{F0: -9223372036854775808, F1: (v26 ^ v26), F2: v26}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
