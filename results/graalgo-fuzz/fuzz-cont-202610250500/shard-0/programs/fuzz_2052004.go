package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int8 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int; F1 int64; F2 uint16 }
type GenS5 struct { F0 int8; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint64; F1 bool }

func main() {
    var v1 uint = 3
    var v2 uint32 = 4294967295
    var v3 int16 = 32767
    var v4 uint = v1
    v4 -= (^(v1))
    v3 = (v3 / ((v3) | 1))
    for fi5 := 0; fi5 < 3; fi5++ {
        if (false || (!(!(v2 < (v2 / ((42) | 1)))))) {
            v2 -= v2
            v6 := GenS1{F0: -43, F1: 2, F2: 0}
            v7 := GenS2{F0: true, F1: false}
            _ = v6
            _ = v7
        } else {
            switch "" {
            case "":
                v8 := []int{fi5, fi5, (fi5 >> ((v4) & 63))}
                _ = v8
            default:
                v9 := []int{(fi5 / ((fi5) | 1))}
                _ = v9
            }
        }
        var v10 int64 = -985
        _ = fi5
        _ = v10
    }
    if true {
        switch "hello" {
        case "a":
            v1 -= (^(v4))
        case "hello":
            var v11 uint32 = 42
            _ = v11
        case "":
            v1 = v4
        }
    }
    if true {
        v12 := GenS3{F0: 1, F1: 781}
        v4 -= v4
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint = v4
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint = v1
        _ = fi15
        _ = v16
    }
    v1 = (v1 + v1)
    if (205 < 1) {
        fmt.Println(v2)
        var v17 uint32 = 0
        switch (("a" + ("hello" + ("" + "hello"))) + ("hello" + "")) {
        case "a":
            var v18 int8 = 117
            _ = v18
        case "hello":
            v19 := []int{-9223372036854775808, 886, 9223372036854775807, 42}
            _ = v19
        case "":
            var v20 int32 = 0
            _ = v20
        }
        _ = v17
    } else {
        v21 := GenS4{F0: 2147483647, F1: -9223372036854775808, F2: 63}
        _ = v21
    }
    v1 = (v4 % (((^(v1))) | 1))
    v1 = (v1 >> (((^(v1))) & 63))
    v22 := GenS5{F0: -128, F1: (18 <= 128), F2: 42}
    v23 := GenS6{F0: 4294967295, F1: true}
    v24 := []int{0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
