package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint; F2 bool }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint16, p1 int16, p2 int16) int8 {
    v4 := GenS1{F0: -1}
    fmt.Println(p1, p1, p0)
    var v5 int = 13
    v6 := GenS2{F0: 0}
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 127
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]string)
    var v3 int8 = 100
    v7 := GenS1{F0: 272}
    switch "hello" {
    case "foo bar":
        if (2147483647 != v1) {
            v3 = (v3 | genFn0(255, 437, 631))
            delete(v2, 32768)
        } else {
            v8 := GenS3{F0: v3, F1: 2, F2: (1 <= 824)}
            _ = v8
        }
    case "":
        v1 = (-(v1))
    }
    switch int(v1) {
    case 1:
        fmt.Println(v1, v3)
    case 42:
        var v9 uint32 = 819
        _ = v9
    }
    v10 := GenS1{F0: 885}
    var v11 uint = 18446744073709551615
    v12 := GenS3{F0: (v3 + v3), F1: v11, F2: true}
    var v13 float64 = (1e10 + 1.5)
    var v14 uint8 = 127
    v3 += int8(v14)
    switch "a" {
    case "a":
        v15 := GenS4{F0: 0, F1: 42}
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 6; fi16++ {
            v14 -= (v14 | (v14 % (((v14 * v14)) | 1)))
            _ = len(v2)
            fmt.Println(fi16, v11)
            _ = fi16
        }
    case "":
        v13 = v13
    default:
        if (!false) {
            v14 -= v14
            v11 += (v11 / (((v11 & v11)) | 1))
        }
    }
    if false {
        v17 := GenS5{F0: 2}
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v7
}
